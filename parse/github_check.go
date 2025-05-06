package parse

import (
	"context"
	"log"

	"github.com/google/go-github/v71/github"
)

type RepoList struct {
	Name string
	URL  string
}

func SearchGithub() []RepoList {
	var list_repos []RepoList
	ctx := context.Background()

	client := github.NewClient(nil)

	query := "42CTF in:name,description"

	options := &github.SearchOptions{
		Sort:        "stars",
		Order:       "desc",
		ListOptions: github.ListOptions{PerPage: 50},
	}

	for {
		repos, resp, err := client.Search.Repositories(ctx, query, options)
		var found_repo RepoList

		if err != nil {
			log.Fatal()
		}

		for _, repo := range repos.Repositories {
			found_repo.Name = repo.GetFullName()
			found_repo.URL = repo.GetHTMLURL()
			list_repos = append(list_repos, found_repo)
		}

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}
	return list_repos
}
