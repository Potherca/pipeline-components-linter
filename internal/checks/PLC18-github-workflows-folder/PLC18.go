package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC18001": "The `workflows/` folder MUST contain a `release.yml` file",
	}
}

func PLC18(files map[string]string) []message.Message {
	var fileCodes = map[string]string{
		".github/workflows/release.yml": "PLC18001",
	}

	return asserts.FileExists(files, fileCodes)
}
