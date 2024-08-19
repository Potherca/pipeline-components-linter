package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC14(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes["renovate.json"] = "PLC14001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
