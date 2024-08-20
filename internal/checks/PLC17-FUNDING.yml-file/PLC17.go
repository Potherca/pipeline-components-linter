package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC17001"] = "The `FUNDING.yml` file MUST be identical to `FUNDING.yml` file in the skeleton repository"

	return Codes
}

func PLC17(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/FUNDING.yml"] = "PLC17001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
