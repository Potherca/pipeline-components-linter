package asserts

import (
	"fmt"
	"internal/check"
	"internal/message"
	"maps"
	"slices"
)

func FolderExists(
	files map[string]string,
	fileCodes map[string]string,
) []message.Message {
	var (
		checkMessage string
		messages     []message.Message
		status       check.Status
	)

	fileNames := slices.Collect(maps.Keys(files))
	requiredFiles := slices.Collect(maps.Keys(fileCodes))

	for _, file := range requiredFiles {
		checkMessage = fmt.Sprintf("The repository MUST contain a `%s` directory", file)

		status = check.Fail

		if slices.Contains(fileNames, file) {
			status = check.Pass
		}

		messages = append(messages, message.CreateMessage(status, fileCodes[file], checkMessage))
	}

	return messages
}
