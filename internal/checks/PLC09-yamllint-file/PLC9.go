package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC9(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".yamllint"] = "PLC9001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
