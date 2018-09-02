package main

import (
	"fmt"
	"log"

	"github.com/golangci/golangci-api/app/models"
	"github.com/golangci/golangci-api/app/utils"
	"github.com/golangci/golangci-api/pkg/todo/db"
	"github.com/golangci/golangci-worker/app/lib/github"
	"github.com/golangci/golib/server/context"
	_ "github.com/mattes/migrate/database/postgres" // pg
)

func main() {
	if err := fillUserIDs(); err != nil {
		panic(err)
	}
}

func fillUserIDs() error {
	ctx := utils.NewBackgroundContext()

	var auths []models.GithubAuth
	if err := models.NewGithubAuthQuerySet(db.Get(ctx)).All(&auths); err != nil {
		return fmt.Errorf("can't fetch all github auths: %s", err)
	}

	log.Printf("Got %d github auths to update", len(auths))

	for i, ga := range auths {
		if err := updateAuth(ctx, &ga); err != nil {
			log.Printf("Can't update auth %#v: %s", ga, err)
		}
		log.Printf("#%d/%d: successfully updated auth", i+1, len(auths))
	}

	return nil
}

func updateAuth(ctx *context.C, ga *models.GithubAuth) error {
	gc := github.Context{GithubAccessToken: ga.AccessToken}.GetClient(ctx.Ctx)
	u, _, err := gc.Users.Get(ctx.Ctx, "")
	if err != nil {
		return fmt.Errorf("can't get user: %s", err)
	}

	err = models.NewGithubAuthQuerySet(db.Get(ctx)).IDEq(ga.ID).GetUpdater().
		SetGithubUserID(uint64(u.GetID())).Update()
	if err != nil {
		return fmt.Errorf("can't update github user id to %d: %s", u.ID, err)
	}

	return nil
}