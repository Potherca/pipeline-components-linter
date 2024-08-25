package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"testing"
)

func TestPLC15(t *testing.T) {
	tests := map[string]struct {
		files  map[string]string
		status map[string]check.Status
	}{
		"app folder does not exist": {
			files:  nil,
			status: map[string]check.Status{"PLC15001": check.Skip, "PLC15002": check.Skip, "PLC15003": check.Skip},
		},
		"app folder exist but is empty": {
			files:  map[string]string{"app/": "__DIR__"},
			status: map[string]check.Status{"PLC15001": check.Fail, "PLC15002": check.Skip, "PLC15003": check.Skip},
		},
		"app folder exist and has non-gitkeep file": {
			files:  map[string]string{"app/": "__DIR__", "app/main.go": "__FILE__"},
			status: map[string]check.Status{"PLC15001": check.Pass, "PLC15002": check.Skip, "PLC15003": check.Skip},
		},
		"app folder exist and has non-empty gitkeep file": {
			files:  map[string]string{"app/": "__DIR__", "app/.gitkeep": "content"},
			status: map[string]check.Status{"PLC15001": check.Pass, "PLC15002": check.Pass, "PLC15003": check.Fail},
		},
		"app folder exist and has empty gitkeep file": {
			files:  map[string]string{"app/": "__DIR__", "app/.gitkeep": ""},
			status: map[string]check.Status{"PLC15001": check.Pass, "PLC15002": check.Pass, "PLC15003": check.Pass},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC15(test.files)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status)
			}
		})
	}
}
