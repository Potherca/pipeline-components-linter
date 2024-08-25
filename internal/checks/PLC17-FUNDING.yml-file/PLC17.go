package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC17001"] = "The `FUNDING.yml` file MUST be identical to `FUNDING.yml` file in the skeleton repository"

	return codes
}

func PLC17(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/FUNDING.yml"] = "PLC17001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
