package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"testing"
)

func TestPLC13(t *testing.T) {
	tests := map[string]struct {
		files  map[string]string
		repo   map[string]string
		status map[string]check.Status
	}{
		targetFile + " file absent": {
			files: nil,
			repo:  nil,
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Skip,
				"PLC13003": check.Skip,
				"PLC13004": check.Skip,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
			},
		},
		targetFile + " file absent from skeleton repo": {
			files: map[string]string{targetFile: "mock content"},
			repo:  nil,
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Error,
				"PLC13003": check.Skip,
				"PLC13004": check.Skip,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC13(test.files, test.repo)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status, "%s expected status %v, got %v", message.Code, test.status[message.Code], message.Status)
			}
		})
	}
}
