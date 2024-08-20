package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC9001"] = "The `.yamllint` file MUST be identical to `.yamllint` file in the skeleton repository"

	return Codes
}

func PLC9(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".yamllint"] = "PLC9001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
