package repositorycontents

import (
	"errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
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

var mockError = errors.New("mock error")

func TestGetContent(t *testing.T) {
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
				repository, _ := git.Init(memory.NewStorage(), nil)

				return repository, mockError
			},
			assertions: func(content map[string]string, err error) {
				assert.Equal(t, mockError, err)
				assert.Len(t, content, 0)
			},
		},
		"GetContent should complain when cloned repo does not contain commits": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), nil)

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

	originalFunction := gitClone
	defer func() { gitClone = originalFunction }()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			gitClone = test.mockFunction

			// Act
			content, err := GetContent("/mock/path")

			// Assert
			test.assertions(content, err)
		})
	}
}

func TestGetLogs(t *testing.T) {
	tests := map[string]struct {
		mockFunction func(string) (*git.Repository, error)
		assertions   func([]LogEntry, error)
	}{
		"GetLogs should complain when repo could not be cloned": {
			mockFunction: func(path string) (*git.Repository, error) {
				return nil, mockError
			},
			assertions: func(logs []LogEntry, err error) {
				assert.Equal(t, mockError, err)
				assert.Nil(t, logs)
			},
		},
		"GetLogs should complain when cloned repo contains errors": {
			mockFunction: func(path string) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), nil)

				return repository, mockError
			},
			assertions: func(logs []LogEntry, err error) {
				assert.Equal(t, mockError, err)
				assert.Nil(t, logs)
			},
		},
		"GetLogs should complain when repo does not contain commits": {
			mockFunction: func(path string) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), nil)

				return repository, nil
			},
			assertions: func(logs []LogEntry, err error) {
				assert.Equal(t, errors.New("reference not found"), err)
				assert.Nil(t, logs)
			},
		},
		"GetLogs should return logs when repo contains commits": {
			mockFunction: func(path string) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), memfs.New())

				createCommit(t, repository, nil)

				return repository, nil
			},
			assertions: func(logs []LogEntry, err error) {
				assert.Nil(t, err)
				assert.Len(t, logs, 1)
				assert.IsType(t, LogEntry{}, logs[0])
			},
		},
	}

	originalFunction := gitPlainOpen
	defer func() { gitPlainOpen = originalFunction }()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			gitPlainOpen = test.mockFunction

			// Act
			logs, err := GetLogs("/mock/path")

			// Assert
			test.assertions(logs, err)
		})
	}
}

func TestGetDetails(t *testing.T) {
	tests := map[string]struct {
		mockFunction func(string) (*git.Repository, error)
		assertions   func(Details, error)
	}{
		"GetDetails should complain when repo could not be cloned": {
			mockFunction: func(path string) (*git.Repository, error) {
				return nil, mockError
			},
			assertions: func(details Details, err error) {
				assert.Equal(t, mockError, err)
				assert.Len(t, details, 0)
			},
		},
		"GetDetails should complain when cloned repo contains errors": {
			mockFunction: func(path string) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), nil)

				return repository, mockError
			},
			assertions: func(details Details, err error) {
				assert.Equal(t, mockError, err)
				assert.Len(t, details, 0)
			},
		},
		"GetDetails should return an empty list when repo does not contain remote": {
			mockFunction: func(path string) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), memfs.New())

				createCommit(t, repository, nil)

				return repository, nil
			},
			assertions: func(details Details, err error) {
				assert.Nil(t, err)
				assert.Len(t, details, 0)
			},
		},
		"GetDetails should return details when repo contains remote": {
			mockFunction: func(path string) (*git.Repository, error) {
				repository, _ := git.Init(memory.NewStorage(), nil) // memfs.New()

				repository.CreateRemote(&config.RemoteConfig{
					Name: "origin",
					URLs: []string{"http://foo/foo.git"},
				})

				repository.CreateBranch(&config.Branch{
					Name:   "master",
					Remote: "origin",
					Merge:  "refs/remotes/origin/master",
				})

				return repository, nil
			},
			assertions: func(details Details, err error) {
				assert.Nil(t, err)
				assert.Len(t, details, 1)
				assert.NotEmpty(t, details["origin"])
				assert.IsType(t, RepoDetails{}, details["origin"])

				actual := details["origin"]
				expected := RepoDetails{Remotes: []string{"http://foo/foo.git"}, Branches: []string(nil)}

				assert.Equal(t, expected, actual)
			},
		},
	}

	originalFunction := gitPlainOpen
	defer func() { gitPlainOpen = originalFunction }()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			gitPlainOpen = test.mockFunction

			// Act
			details, err := GetDetails("/mock/path")

			// Assert
			test.assertions(details, err)
		})
	}
}
