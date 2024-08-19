package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC17(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/FUNDING.yml"] = "PLC17001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
