package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC4001"] = "The repository MUST contain an `app/` folder"
	codes["PLC4002"] = "The repository MUST contain a `.github/` folder"

	return codes
}

func PLC4(files map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes["app/"] = "PLC4001"
	fileCodes[".github/"] = "PLC4002"

	return asserts.FolderExists(files, fileCodes)
}
