package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC14001"] = "The `renovate.json` file MUST be identical to `renovate.json` file in the skeleton repository"

	return Codes
}

func PLC14(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes["renovate.json"] = "PLC14001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
