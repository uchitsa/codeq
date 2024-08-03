package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	githubToken = "my_github_token"
	lang        = "Go"
	order       = "desc"
	pageCnt     = 10
	perPage     = 100
	sort        = "stars"
)

func main() {
	ctx := context.Background()
	tokenSrc := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	ocl := oauth2.NewClient(ctx, tokenSrc)
	cl := github.NewClient(ocl)

	for i := 1; i <= pageCnt; i++ {
		_, _, err := cl.Search.Repositories(ctx, fmt.Sprintf("language:%s", lang), &github.SearchOptions{
			Sort:      sort,
			Order:     order,
			TextMatch: false,
			ListOptions: github.ListOptions{
				Page:    i,
				PerPage: perPage,
			},
		})
		if err != nil {
			fmt.Printf("get repositories error:%v\n", err)
			continue
		}

	}
}
