package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC19(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/workflows/release.yml"] = "PLC19001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
