package main

import (
	"leakGuard/compare"
	"leakGuard/parse"
	"leakGuard/storage"
)

func main() {
	repo_list := storage.LoadList()
	github_list := parse.SearchGithub()
	additions := compare.CompareLists(github_list, repo_list)

	if len(additions) > 0 {
		sendToTelegram(additions)
		sendToDiscord(additions)
	}
}
