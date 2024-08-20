package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC4001"] = "The repository MUST contain an `app/` folder"
	Codes["PLC4002"] = "The repository MUST contain a `.github/` folder"

	return Codes
}

func PLC4(files map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes["app/"] = "PLC4001"
	fileCodes[".github/"] = "PLC4002"

	return asserts.FolderExists(files, fileCodes)
}
