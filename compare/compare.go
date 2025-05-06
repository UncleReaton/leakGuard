package compare

import (
	"leakGuard/parse"
	"leakGuard/storage"
)

func CompareLists(new_list, old_list []parse.RepoList) []parse.RepoList {
	existing := make(map[string]struct{})
	for _, repo := range old_list {
		existing[repo.Name] = struct{}{}
	}

	var additions []parse.RepoList
	for _, repo := range new_list {
		if _, found := existing[repo.Name]; !found {
			additions = append(additions, repo)
		}
	}

	if len(additions) > 0 {
		storage.SaveToFile(new_list)
	}

	return additions
}
