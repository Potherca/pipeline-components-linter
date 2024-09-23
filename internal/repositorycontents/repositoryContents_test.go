package repositorycontents

import (
	"errors"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockError = errors.New("mock error")

func TestGetContent(t *testing.T) {
	tests := map[string]struct {
		mockFunction func(storage.Storer, billy.Filesystem, *git.CloneOptions) (*git.Repository, error)
		assertions   func(map[string]string, error)
	}{
		"test": {
			mockFunction: func(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
				return nil, mockError
			},
			assertions: func(logs map[string]string, err error) {
				assert.Equal(t, mockError, err)
				assert.Len(t, logs, 0)
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
