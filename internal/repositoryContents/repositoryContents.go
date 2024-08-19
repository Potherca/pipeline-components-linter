package repositoryContents

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
)

func GetContent(repo string) (map[string]string, error) {
	files := make(map[string]string)

	repository, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repo})

	if err == nil && repository != nil {
		ref, err := repository.Head()

		if err == nil && ref != nil {
			commit, err := repository.CommitObject(ref.Hash())

			if err == nil && commit != nil {
				tree, err := commit.Tree()

				if err == nil && tree != nil {
					fileIter := tree.Files()

					err = fileIter.ForEach(func(file *object.File) error {
						reader, err := file.Blob.Reader()

						if err == nil {
							buffer, err := io.ReadAll(reader)

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
