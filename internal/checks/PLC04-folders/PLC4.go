package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC4001": "The repository MUST contain an `app/` folder",
		"PLC4002": "The repository MUST contain a `.github/` folder",
	}
}

func PLC4(files map[string]string) []message.Message {
	var fileCodes = map[string]string{
		"app/":     "PLC4001",
		".github/": "PLC4002",
	}

	return asserts.FolderExists(files, fileCodes)
}
