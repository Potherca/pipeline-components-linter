package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC19001"] = "The `release.yml` file MUST be identical to `release.yml` file in the skeleton repository"

	return codes
}

func PLC19(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".github/workflows/release.yml"] = "PLC19001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
