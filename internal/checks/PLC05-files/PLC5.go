package checks

import (
	"fmt"
	"internal/check"
	"internal/message"
	"slices"
)

func PLC5(fileNames []string) []message.Message {
	var messages []message.Message

	var fileCodes = make(map[string]string)

	fileCodes[".gitignore"] = "PLC5001"
	fileCodes[".gitlab-ci.yml"] = "PLC5002"
	fileCodes[".mdlrc"] = "PLC5003"
	fileCodes[".yamllint"] = "PLC5004"
	fileCodes["action.yml"] = "PLC5005"
	fileCodes["Dockerfile"] = "PLC5006"
	fileCodes["LICENSE"] = "PLC5007"
	fileCodes["README.md"] = "PLC5008"
	fileCodes["renovate.json"] = "PLC5009"

	requiredFiles := [9]string{
		".gitignore",
		".gitlab-ci.yml",
		".mdlrc",
		".yamllint",
		"action.yml",
		"Dockerfile",
		"LICENSE",
		"README.md",
		"renovate.json",
	}

	for _, file := range requiredFiles {
		status := check.Fail
		if slices.Contains(fileNames, file) {
			status = check.Pass
		}

		messages = append(messages, message.CreateMessage(
			status,
			fileCodes[file],
			fmt.Sprintf("The repository MUST contain a %s file", file),
		))
	}

	return messages
}
