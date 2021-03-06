package implementations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/golangci/golangci-api/internal/api/apierrors"
	"github.com/golangci/golangci-api/internal/shared/logutil"
	"github.com/golangci/golangci-api/internal/shared/providers/provider"
	"github.com/golangci/golangci-api/pkg/api/models"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Check the struct is implementing the Provider interface.
var _ provider.Provider = &Github{}

const GithubProviderName = "github.com"

type Github struct {
	accessToken string
	baseURL     *url.URL
	log         logutil.Log
}

func NewGithub(log logutil.Log, accessToken string) *Github {
	return &Github{
		log:         log,
		accessToken: accessToken,
	}
}

func (p Github) Name() string {
	return GithubProviderName
}

func (p Github) LinkToPullRequest(repo *models.Repo, num int) string {
	return fmt.Sprintf("https://github.com/%s/pull/%d", repo.DisplayFullName, num)
}

func (p *Github) SetBaseURL(s string) error {
	baseURL, err := url.Parse(s)
	if err != nil {
		return errors.Wrap(err, "failed to parse url")
	}

	p.baseURL = baseURL
	return nil
}

func (p Github) client(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: p.accessToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)
	if p.baseURL != nil {
		c.BaseURL = p.baseURL
	}

	return c
}

func (p Github) unwrapError(err error) error {
	if er, ok := err.(*github.ErrorResponse); ok {
		respCode := er.Response.StatusCode
		if respCode == http.StatusNotFound {
			return provider.ErrNotFound
		}
		if respCode == http.StatusUnauthorized {
			return provider.ErrUnauthorized
		}
		if respCode == http.StatusForbidden {
			if strings.Contains(er.Message, "Repository was archived") {
				// usually it contains "Repository was archived so is read-only."
				return provider.ErrRepoWasArchived
			}
		}
		if respCode == http.StatusUnprocessableEntity {
			// full msg is "You must purchase at least one more seat to add this user as a collaborator."
			if strings.Contains(er.Error(), "must purchase at least one more seat") {
				return provider.ErrNoFreeOrgSeats
			}
		}
	}

	return err
}

func parseGithubRepository(r *github.Repository, root bool) *provider.Repo {
	var source *provider.Repo
	if root && r.GetSource() != nil { // repository is a fork, select source
		source = parseGithubRepository(r.GetSource(), false)
	}

	var orgName string
	if r.GetOrganization() != nil {
		orgName = r.GetOrganization().GetLogin()
	}

	return &provider.Repo{
		ID:              r.GetID(),
		FullName:        r.GetFullName(),
		IsAdmin:         r.GetPermissions()["admin"],
		IsPrivate:       r.GetPrivate(),
		DefaultBranch:   r.GetDefaultBranch(),
		Source:          source,
		StargazersCount: r.GetStargazersCount(),
		Language:        r.GetLanguage(),
		Organization:    orgName,
		OwnerID:         r.GetOwner().GetID(),
	}
}

func parseGithubBranch(b *github.Branch) *provider.Branch {
	return &provider.Branch{
		CommitSHA: b.GetCommit().GetSHA(),
	}
}

func parseGithubPullRequestBranch(b *github.PullRequestBranch) *provider.Branch {
	return &provider.Branch{
		CommitSHA: b.GetSHA(),
	}
}

func parseGithubOrgMembership(m *github.Membership) *provider.OrgMembership {
	return &provider.OrgMembership{
		ID:      m.GetOrganization().GetID(),
		Name:    m.GetOrganization().GetLogin(),
		IsAdmin: m.GetRole() == "admin" && m.GetState() == "active",
	}
}

func (p Github) GetRepoByName(ctx context.Context, owner, repo string) (*provider.Repo, error) {
	r, _, err := p.client(ctx).Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, p.unwrapError(err)
	}

	return parseGithubRepository(r, true), nil
}

func (p Github) GetOrgMembershipByName(ctx context.Context, org string) (*provider.OrgMembership, error) {
	m, _, err := p.client(ctx).Organizations.GetOrgMembership(ctx, "", org)
	if err != nil {
		return nil, p.unwrapError(err)
	}

	return parseGithubOrgMembership(m), nil
}

func (p Github) parseHook(h *github.Hook) *provider.Hook {
	ctObj := h.Config["content_type"]
	ct := ""
	if ctObj != nil {
		if ctStr, ok := ctObj.(string); ok {
			ct = ctStr
		}
	}

	return &provider.Hook{
		HookConfig: provider.HookConfig{
			Name:        h.GetName(),
			Events:      h.Events,
			URL:         h.GetURL(),
			ContentType: ct,
		},
		ID: h.GetID(),
	}
}

func (p Github) CreateRepoHook(ctx context.Context, owner, repo string,
	hook *provider.HookConfig) (*provider.Hook, error) {

	githubHookCfg := github.Hook{
		Name:   &hook.Name,
		Events: hook.Events,
		Config: map[string]interface{}{
			"url":          hook.URL,
			"content_type": hook.ContentType,
		},
	}
	rh, _, err := p.client(ctx).Repositories.CreateHook(ctx, owner, repo, &githubHookCfg)
	if err != nil {
		return nil, p.unwrapError(err)
	}

	return p.parseHook(rh), nil
}

func (p Github) ListRepoHooks(ctx context.Context, owner, repo string) ([]provider.Hook, error) {
	listOptions := github.ListOptions{
		PerPage: 100,
	}
	hooks, _, err := p.client(ctx).Repositories.ListHooks(ctx, owner, repo, &listOptions)
	if err != nil {
		return nil, p.unwrapError(err)
	}

	if len(hooks) == listOptions.PerPage {
		return nil, fmt.Errorf("repo %s/%s has >%d hooks, need to support pagination",
			owner, repo, len(hooks))
	}

	var retHooks []provider.Hook
	for _, h := range hooks {
		retHooks = append(retHooks, *p.parseHook(h))
	}
	return retHooks, nil
}

func (p Github) GetBranch(ctx context.Context, owner, repo, branch string) (*provider.Branch, error) {
	grb, _, err := p.client(ctx).Repositories.GetBranch(ctx, owner, repo, branch)
	if err != nil {
		return nil, p.unwrapError(err)
	}

	return parseGithubBranch(grb), nil
}

func (p Github) DeleteRepoHook(ctx context.Context, owner, repo string, hookID int) error {
	_, err := p.client(ctx).Repositories.DeleteHook(ctx, owner, repo, hookID)
	if err != nil {
		return p.unwrapError(err)
	}

	return nil
}

func parseGithubCommitAuthor(ca *github.CommitAuthor) *provider.CommitAuthor {
	return &provider.CommitAuthor{
		Email: ca.GetEmail(),
	}
}

func parseGithubCommit(c *github.RepositoryCommit) *provider.Commit {
	gc := c.GetCommit()
	return &provider.Commit{
		SHA:       c.GetSHA(),
		Author:    parseGithubCommitAuthor(gc.GetAuthor()),
		Committer: parseGithubCommitAuthor(gc.GetCommitter()),
	}
}

func (p Github) ListPullRequestCommits(ctx context.Context, owner, repo string, number int) ([]*provider.Commit, error) {
	commits, _, err := p.client(ctx).PullRequests.ListCommits(ctx, owner, repo, number, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, p.unwrapError(err)
	}

	var ret []*provider.Commit
	for _, c := range commits {
		ret = append(ret, parseGithubCommit(c))
	}

	return ret, nil
}

func (p Github) SetCommitStatus(ctx context.Context, owner, repo, ref string, status *provider.CommitStatus) error {
	githubStatus := github.RepoStatus{
		State:       github.String(status.State),
		Description: github.String(status.Description),
		Context:     github.String(status.Context),
	}
	if status.TargetURL != "" {
		githubStatus.TargetURL = github.String(status.TargetURL)
	}

	_, _, err := p.client(ctx).Repositories.CreateStatus(ctx, owner, repo, ref, &githubStatus)
	if err != nil {
		return p.unwrapError(err)
	}

	return nil
}

func (p Github) ListRepos(ctx context.Context, cfg *provider.ListReposConfig) ([]provider.Repo, error) {
	opts := github.RepositoryListOptions{
		Visibility: cfg.Visibility,
		Sort:       cfg.Sort,
		ListOptions: github.ListOptions{
			PerPage: 100, // 100 is a max allowed value
		},
	}

	var ret []provider.Repo
	for {
		pageRepos, resp, err := p.client(ctx).Repositories.List(ctx, "", &opts)
		if err != nil {
			return nil, p.unwrapError(err)
		}

		for _, r := range pageRepos {
			ret = append(ret, *parseGithubRepository(r, true))
		}

		if resp.NextPage == 0 { // it's the last page
			break
		}

		if opts.Page == cfg.MaxPages { // TODO: fetch all, now we limit it
			p.log.Warnf("Limited repo list to %d entries (%d pages)", len(ret), cfg.MaxPages)
			break
		}

		opts.Page = resp.NextPage
	}

	return ret, nil
}

func (p Github) ListOrgMemberships(ctx context.Context, cfg *provider.ListOrgsConfig) ([]provider.OrgMembership, error) {
	opts := github.ListOrgMembershipsOptions{
		State: cfg.MembershipState,
		ListOptions: github.ListOptions{
			PerPage: 100, // 100 is a max allowed value
		},
	}

	var ret []provider.OrgMembership
	for {
		pageMemberships, resp, err := p.client(ctx).Organizations.ListOrgMemberships(ctx, &opts)
		if err != nil {
			return nil, p.unwrapError(err)
		}

		for _, m := range pageMemberships {
			ret = append(ret, *parseGithubOrgMembership(m))
		}

		if resp.NextPage == 0 { // it's the last page
			break
		}

		if opts.Page == cfg.MaxPages { // TODO: fetch all, now we limit it
			p.log.Warnf("Limited repo list to %d entries (%d pages)", len(ret), cfg.MaxPages)
			break
		}

		opts.Page = resp.NextPage
	}

	return ret, nil
}

func (p Github) GetPullRequest(ctx context.Context, owner, repo string, number int) (*provider.PullRequest, error) {
	pr, _, err := p.client(ctx).PullRequests.Get(ctx, owner, repo, number)
	if err != nil {
		return nil, p.unwrapError(err)
	}

	return &provider.PullRequest{
		Head:  parseGithubPullRequestBranch(pr.GetHead()),
		State: pr.GetState(),
	}, nil
}

func (p Github) ParsePullRequestEvent(ctx context.Context, payload []byte) (*provider.PullRequestEvent, error) {
	var ghEvent github.PullRequestEvent
	if err := json.Unmarshal(payload, &ghEvent); err != nil {
		return nil, errors.Wrapf(apierrors.ErrBadRequest, "invalid payload json: %s", err)
	}

	if ghEvent.PullRequest == nil {
		return nil, errors.New("got github webhook event without pull request field")
	}

	pr := ghEvent.GetPullRequest()
	repo := ghEvent.GetRepo()

	owner := repo.GetOwner().GetLogin()
	name := repo.GetName()
	if owner == "" {
		return nil, errors.New("no repo owner in event")
	}
	if name == "" {
		return nil, errors.New("no repo name in event")
	}

	// don't extract repo just as parseGithubRepository(ghEvent.GetRepo(), true)
	// because ghEvent.GetRepo() doesn't contain organization info
	fetchedRepo, err := p.GetRepoByName(ctx, repo.GetOwner().GetLogin(), repo.GetName())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch repo %s/%s", repo.GetOwner().GetLogin(), repo.GetName())
	}

	return &provider.PullRequestEvent{
		Repo:              fetchedRepo,
		Head:              parseGithubPullRequestBranch(pr.GetHead()),
		PullRequestNumber: pr.GetNumber(),
		Action:            provider.PullRequestAction(ghEvent.GetAction()),
	}, nil
}

func (p Github) AddCollaborator(ctx context.Context, owner, repo, username string) (*provider.RepoInvitation, error) {
	opts := github.RepositoryAddCollaboratorOptions{}
	resp, err := p.client(ctx).Repositories.AddCollaborator(ctx, owner, repo, username, &opts)
	if err != nil {
		// err goes to user if err == ErrNoFreeOrgSeats, don't write much here
		return nil, p.unwrapError(err)
	}

	if resp.StatusCode == http.StatusNoContent {
		return &provider.RepoInvitation{IsAlreadyCollaborator: true}, nil
	}

	// resp.Body is already closed, go-github needs changes in a upstream to properly support
	// response for AddCollaborator

	invitations, _, err := p.client(ctx).Repositories.ListInvitations(ctx, owner, repo, nil)
	if err != nil {
		return nil, errors.Wrap(p.unwrapError(err), "failed to list invitations")
	}

	for _, invitation := range invitations {
		if invitation.Invitee.GetLogin() == username {
			return &provider.RepoInvitation{ID: invitation.GetID()}, nil
		}
	}

	return nil, fmt.Errorf("not found needed invitation within invitations %#v", invitations)
}

func (p Github) RemoveCollaborator(ctx context.Context, owner, repo, username string) error {
	_, err := p.client(ctx).Repositories.RemoveCollaborator(ctx, owner, repo, username)
	if err != nil {
		return p.unwrapError(err)
	}

	return nil
}

func (p Github) AcceptRepoInvitation(ctx context.Context, invitationID int) error {
	_, err := p.client(ctx).Users.AcceptInvitation(ctx, invitationID)
	if err != nil {
		return p.unwrapError(err)
	}

	return nil
}
