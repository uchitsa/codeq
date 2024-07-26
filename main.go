package main

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const githubToken = "my_github_token"

func main() {
	ctx := context.Background()
	tokenSrc := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	ocl := oauth2.NewClient(ctx, tokenSrc)
	_ = github.NewClient(ocl)
}
