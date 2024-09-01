package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"testing"
)

func TestPLC4(t *testing.T) {
	tests := map[string]struct {
		files  map[string]string
		status map[string]check.Status
	}{
		"app/ folder absent, .github/ folder absent": {
			files:  nil,
			status: map[string]check.Status{"PLC4001": check.Fail, "PLC4002": check.Fail},
		},
		"app/ folder present, .github/ folder absent": {
			files:  map[string]string{"app/": "__DIR__"},
			status: map[string]check.Status{"PLC4001": check.Pass, "PLC4002": check.Fail},
		},
		"app/ folder does absent, .github/ folder present": {
			files:  map[string]string{".github/": "__DIR__"},
			status: map[string]check.Status{"PLC4001": check.Fail, "PLC4002": check.Pass},
		},
		"app/ folder present, .github/ folder present": {
			files:  map[string]string{"app/": "__DIR__", ".github/": "__DIR__"},
			status: map[string]check.Status{"PLC4001": check.Pass, "PLC4002": check.Pass},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC4(test.files)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status)
			}
		})
	}
}
