package repositorycontents

import (
	"errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createCommit(t *testing.T, repo *git.Repository, files map[string]string) string {
	t.Helper()

	worktree, _ := repo.Worktree()

	for fileName, content := range files {
		file, _ := worktree.Filesystem.Create(fileName)
		file.Write([]byte(content))
		worktree.Add(fileName)
	}

	hash, _ := worktree.Commit("Commit message", &git.CommitOptions{
		AllowEmptyCommits: len(files) == 0,
	})

	return hash.String()
}

func TestGetContent(t *testing.T) {
	mockError := errors.New("mock error")
	mockFiles := map[string]string{"foo.txt": "foo content"}

	tests := map[string]struct {
		mockFunction func(storage.Storer, billy.Filesystem, *git.CloneOptions) (*git.Repository, error)
		assertions   func(map[string]string, error)
	}{
		"GetContent should complain when repo could not be cloned": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				return nil, mockError
			},
			assertions: func(content map[string]string, err error) {
				assert.Equal(t, mockError, err)
				assert.Len(t, content, 0)
			},
		},
		"GetContent should complain when cloned repo contains errors": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), memfs.New())

				return repository, mockError
			},
			assertions: func(content map[string]string, err error) {
				assert.Equal(t, mockError, err)
				assert.Len(t, content, 0)
			},
		},
		"GetContent should complain when cloned repo does not contain commits": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), memfs.New())

				return repository, nil
			},
			assertions: func(content map[string]string, err error) {
				assert.Equal(t, errors.New("reference not found"), err)
				assert.Len(t, content, 0)
			},
		},
		"GetContent should return an empty list when no files are present": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), memfs.New())

				createCommit(t, repository, nil)

				return repository, nil
			},
			assertions: func(content map[string]string, err error) {
				assert.Len(t, content, 0)
			},
		},
		"GetContent should return a list of commits when files are present": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), memfs.New())

				createCommit(t, repository, mockFiles)

				return repository, nil
			},
			assertions: func(content map[string]string, err error) {
				assert.Equal(t, mockFiles, content)
				assert.Len(t, content, 1)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			originalFunction := gitClone
			gitClone = test.mockFunction

			// Act
			content, err := GetContent("/mock/path")

			// Assert
			test.assertions(content, err)

			// After
			gitClone = originalFunction
		})
	}
}
