package repositorycontents

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
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
