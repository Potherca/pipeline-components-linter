package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC19001"] = "The `release.yml` file MUST be identical to `release.yml` file in the skeleton repository"

	return Codes
}

func PLC19(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/workflows/release.yml"] = "PLC19001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
