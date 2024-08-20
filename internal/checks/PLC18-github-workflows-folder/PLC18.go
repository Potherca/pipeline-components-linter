package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC18001"] = "The `workflows/` folder MUST contain a `release.yml` file"

	return Codes
}

func PLC18(files map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/workflows/release.yml"] = "PLC18001"

	return asserts.FileExists(files, fileCodes)
}
