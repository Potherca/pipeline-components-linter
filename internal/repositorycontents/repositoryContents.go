package repositorycontents

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"strings"
)

var gitClone = git.Clone
var gitPlainOpen = git.PlainOpen

func GetContent(repo string) (map[string]string, error) {
	var (
		buffer     []byte
		commit     *object.Commit
		err        error
		files      map[string]string
		reader     io.ReadCloser
		ref        *plumbing.Reference
		repository *git.Repository
		tree       *object.Tree
	)

	files = make(map[string]string)

	repository, err = gitClone(memory.NewStorage(), nil, &git.CloneOptions{URL: repo})

	if err == nil && repository != nil {
		ref, err = repository.Head()

		if err == nil && ref != nil {
			commit, err = repository.CommitObject(ref.Hash())

			if err == nil && commit != nil {
				tree, err = commit.Tree()

				if err == nil && tree != nil {
					fileIter := tree.Files()

					err = fileIter.ForEach(func(file *object.File) error {
						reader, err = file.Blob.Reader()

						if err == nil {
							buffer, err = io.ReadAll(reader)

							if err == nil {
								contents := string(buffer)
								files[file.Name] = contents
							}
						}

						return err
					})
				}
			}
		}
	}

	return files, err
}

func GetDetails(path string) (Details, error) {
	var (
		err     error
		details Details
	)

	details = Details{}

	repository, err := gitPlainOpen(path)

	if err == nil && repository != nil {
		remotes, err := repository.Remotes()

		if err == nil && remotes != nil {
			for _, remote := range remotes {
				var (
					branches []string
				)

				remoteName := remote.Config().Name
				refs, _ := repository.References()

				err = refs.ForEach(func(ref *plumbing.Reference) error {
					if ref.Name().IsRemote() && strings.HasPrefix(ref.Name().Short(), remoteName+"/") {
						after, _ := strings.CutPrefix(ref.Name().Short(), remoteName+"/")
						branches = append(branches, after)
					}
					return nil
				})

				details[remoteName] = RepoDetails{
					Branches: branches,
					Remotes:  remote.Config().URLs,
				}
			}
		}
	}

	if err == nil && repository != nil {
		/*		log, err = repository.Log(&git.LogOptions{})
				if err == nil && log != nil {
					err = log.ForEach(func(commit *object.Commit) error {
						details = append(details, LogEntry{
							Timestamp: commit.Author.When,
						})

						return nil
					})
				}
		*/
	}

	return details, err
}

func GetLogs(path string) ([]LogEntry, error) {
	var (
		err  error
		log  object.CommitIter
		logs []LogEntry
	)

	repository, err := gitPlainOpen(path)

	if err == nil && repository != nil {
		log, err = repository.Log(&git.LogOptions{})
		if err == nil && log != nil {
			err = log.ForEach(func(commit *object.Commit) error {
				logs = append(logs, LogEntry{
					Timestamp: commit.Author.When,
				})

				return nil
			})
		}
	}

	return logs, err
}
