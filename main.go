package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"os/exec"
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
		repos, _, err := cl.Search.Repositories(ctx, fmt.Sprintf("language:%s", lang), &github.SearchOptions{
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

		for _, repo := range repos.Repositories {
			cloneRepo(repo)
		}
	}
}

func cloneRepo(repo github.Repository) {
	fmt.Printf("cloning repo: [%s]", repo.GetName())

	cmd := exec.Command("git", "clone", repo.GetCloneURL(), repo.GetName())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("cloning repo [%s] error: %v\n", repo.GetName(), err)
	}
}
