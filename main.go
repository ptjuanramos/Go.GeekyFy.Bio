package main

import (
	"context"

	"github.com/google/go-github/github"

	"bio.randomgithub/biohelper"
)

func main() {
	authenticationClient := biohelper.GetAuthenticationClient()
	client := github.NewClient(authenticationClient)

	newUserBio := biohelper.GetNewUserBio()
	userInput := &github.User{Bio: &newUserBio}
	client.Users.Edit(context.Background(), userInput)
}
