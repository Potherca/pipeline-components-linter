package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"strconv"
	"testing"
)

type testEntry struct {
	files  map[string]string
	status map[string]check.Status
}

func createTestEntry(files map[string]string, status map[string]check.Status) testEntry {
	return testEntry{
		files:  files,
		status: status,
	}
}

func TestPLC5(t *testing.T) {
	fileCodes := map[string]string{
		".gitignore":     "PLC5001",
		".gitlab-ci.yml": "PLC5002",
		".mdlrc":         "PLC5003",
		".yamllint":      "PLC5004",
		"action.yml":     "PLC5005",
		"Dockerfile":     "PLC5006",
		"LICENSE":        "PLC5007",
		"README.md":      "PLC5008",
		"renovate.json":  "PLC5009",
	}

	tests := make(map[string]testEntry)

	files := []string{
		".gitignore",
		".gitlab-ci.yml",
		".mdlrc",
		".yamllint",
		"action.yml",
		"Dockerfile",
		"LICENSE",
		"README.md",
		"renovate.json",
	}

	numFiles := len(files)
	numCombinations := 1 << numFiles

	for i := 0; i < numCombinations; i++ {
		testFiles := make(map[string]string)
		expectedStatus := make(map[string]check.Status)
		var testName string
		fileCount := 0

		if i == 0 {
			testName = "All " + strconv.Itoa(len(files)) + " files absent"
		} else if i == numCombinations-1 {
			testName = "All " + strconv.Itoa(len(files)) + " files present"
		} else {
			testName = "Combination: "
		}

		for j := 0; j < numFiles; j++ {
			if i&(1<<j) != 0 {
				testFiles[files[j]] = "mock content"
				expectedStatus[fileCodes[files[j]]] = check.Pass
				fileCount++
				testName += files[j] + " "
			} else {
				expectedStatus[fileCodes[files[j]]] = check.Fail
			}
		}

		if fileCount == 1 {
			for file := range testFiles {
				testName = "Only " + file + " present"
			}
		}

		if fileCount == numFiles-1 {
			for file := range testFiles {
				testName = "Only " + file + " absent"
			}
		}

		tests[testName] = createTestEntry(testFiles, expectedStatus)
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC5(test.files)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status)
			}
		})
	}
}
