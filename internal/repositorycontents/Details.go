package repositorycontents

type RepoDetails struct {
	Remotes  []string
	Branches []string
}

type Details map[string]RepoDetails
