package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC8001": "The `.mdlrc` file MUST be identical to `.mdlrc` file in the skeleton repository",
	}
}

func PLC8(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = map[string]string{
		".mdlrc": "PLC8001",
	}

	return asserts.CompareFiles(files, repo, fileCodes)
}
