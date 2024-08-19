package checks

import (
	"fmt"
	"internal/check"
	"internal/message"
	"maps"
	"slices"
)

func PLC4(files map[string]string) []message.Message {
	var messages []message.Message

	var fileCodes = make(map[string]string)

	fileCodes["app/"] = "PLC4001"
	fileCodes[".github/"] = "PLC4002"

	requiredFiles := [2]string{
		".github/",
		"app/",
	}

	fileNames := slices.Collect(maps.Keys(files))

	for _, file := range requiredFiles {
		status := check.Fail
		if slices.Contains(fileNames, file) {
			status = check.Pass
		}

		checkMessage := fmt.Sprintf("The repository MUST contain a %s directory", file)
		code := fileCodes[file]

		messages = append(messages, message.CreateMessage(status, code, checkMessage))
	}

	return messages
}
