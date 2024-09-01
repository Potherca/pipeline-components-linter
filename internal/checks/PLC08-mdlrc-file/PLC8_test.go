package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"testing"
)

func TestPLC8(t *testing.T) {
	tests := map[string]struct {
		files  map[string]string
		repo   map[string]string
		status map[string]check.Status
	}{
		".mdlrc file absent": {
			files:  nil,
			repo:   map[string]string{".mdlrc": "mock content"},
			status: map[string]check.Status{"PLC8001": check.Skip},
		},
		".mdlrc file present but not identical": {
			files:  map[string]string{".mdlrc": ""},
			repo:   map[string]string{".mdlrc": "mock content"},
			status: map[string]check.Status{"PLC8001": check.Fail},
		},
		".mdlrc file present but missing in repo": {
			files:  map[string]string{".mdlrc": "mock content"},
			repo:   nil,
			status: map[string]check.Status{"PLC8001": check.Error},
		},
		".mdlrc file present and identical": {
			files:  map[string]string{".mdlrc": "mock content"},
			repo:   map[string]string{".mdlrc": "mock content"},
			status: map[string]check.Status{"PLC8001": check.Pass},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC8(test.files, test.repo)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status)
			}
		})
	}
}
