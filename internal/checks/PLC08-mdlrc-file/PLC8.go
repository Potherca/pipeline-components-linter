package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC8(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".mdlrc"] = "PLC8001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
