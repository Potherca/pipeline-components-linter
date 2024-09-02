package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC14001": "The `renovate.json` file MUST be identical to `renovate.json` file in the skeleton repository",
	}
}

func PLC14(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = map[string]string{
		"renovate.json": "PLC14001",
	}

	return asserts.CompareFiles(files, repo, fileCodes)
}
