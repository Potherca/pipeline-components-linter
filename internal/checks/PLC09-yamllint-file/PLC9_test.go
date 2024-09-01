package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"testing"
)

func TestPLC9(t *testing.T) {
	targetFile := ".yamllint"

	tests := map[string]struct {
		files  map[string]string
		repo   map[string]string
		status map[string]check.Status
	}{
		targetFile + " file absent": {
			files:  nil,
			repo:   map[string]string{targetFile: "mock content"},
			status: map[string]check.Status{"PLC9001": check.Skip},
		},
		targetFile + " file present but not identical": {
			files:  map[string]string{targetFile: ""},
			repo:   map[string]string{targetFile: "mock content"},
			status: map[string]check.Status{"PLC9001": check.Fail},
		},
		targetFile + " file present but missing in repo": {
			files:  map[string]string{targetFile: "mock content"},
			repo:   nil,
			status: map[string]check.Status{"PLC9001": check.Error},
		},
		targetFile + " file present and identical": {
			files:  map[string]string{targetFile: "mock content"},
			repo:   map[string]string{targetFile: "mock content"},
			status: map[string]check.Status{"PLC9001": check.Pass},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC9(test.files, test.repo)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status)
			}
		})
	}
}
