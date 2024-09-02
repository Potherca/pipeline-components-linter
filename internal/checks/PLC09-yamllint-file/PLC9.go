package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC9001": "The `.yamllint` file MUST be identical to `.yamllint` file in the skeleton repository",
	}
}

func PLC9(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = map[string]string{
		".yamllint": "PLC9001",
	}

	return asserts.CompareFiles(files, repo, fileCodes)
}
