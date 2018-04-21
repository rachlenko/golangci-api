package oauth

import (
	"os"

	"github.com/markbates/goth/providers/github"
)

func GetPublicReposGithubAuthorizer(cbURL string) *Authorizer {
	gh := github.New(
		os.Getenv("GITHUB_KEY"),
		os.Getenv("GITHUB_SECRET"),
		os.Getenv("GITHUB_CALLBACK_HOST")+cbURL,
		"user:email",
		"public_repo",
	)

	const providerName = "github"
	gh.SetName(providerName)

	return &Authorizer{
		providerName: providerName,
		provider:     gh,
	}
}

func GetPrivateReposGithubAuthorizer(cbURL string) *Authorizer {
	gh := github.New(
		os.Getenv("GITHUB_KEY"),
		os.Getenv("GITHUB_SECRET"),
		os.Getenv("GITHUB_CALLBACK_HOST")+cbURL,
		"user:email",
		"repo",
	)

	const providerName = "github/private"
	gh.SetName(providerName)

	return &Authorizer{
		providerName: providerName,
		provider:     gh,
	}
}
