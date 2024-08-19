package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC4(files map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes["app/"] = "PLC4001"
	fileCodes[".github/"] = "PLC4002"

	return asserts.FolderExists(files, fileCodes)
}
