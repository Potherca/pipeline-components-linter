package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC8001"] = "The `.mdlrc` file MUST be identical to `.mdlrc` file in the skeleton repository"

	return codes
}

func PLC8(files map[string]string, repo map[string]string) []message.Message {
	var fileCodes = make(map[string]string)

	fileCodes[".mdlrc"] = "PLC8001"

	return asserts.CompareFiles(files, repo, fileCodes)
}
