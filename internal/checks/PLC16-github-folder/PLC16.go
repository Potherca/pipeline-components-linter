package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC16(files map[string]string) []message.Message {
	var messages []message.Message

	messages = append(messages, checkFiles(files)...)
	messages = append(messages, checkFolders(files)...)

	return messages
}

func checkFiles(files map[string]string) []message.Message {
	var fileCodes = map[string]string{
		".github/FUNDING.yml": "PLC16001",
	}

	return asserts.FileExists(files, fileCodes)
}

func checkFolders(files map[string]string) []message.Message {
	var fileCodes = map[string]string{
		".github/workflows/": "PLC16002",
	}

	return asserts.FolderExists(files, fileCodes)
}
