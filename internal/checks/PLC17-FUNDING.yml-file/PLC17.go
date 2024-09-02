package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC17001": "The `FUNDING.yml` file MUST be identical to `FUNDING.yml` file in the skeleton repository",
	}
}

func PLC17(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = map[string]string{
		".github/FUNDING.yml": "PLC17001",
	}

	return asserts.CompareFiles(files, repo, fileCodes)
}
