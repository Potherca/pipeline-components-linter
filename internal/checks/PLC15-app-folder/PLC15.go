package checks

import (
	"internal/check"
	"internal/message"
	"strings"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC15001"] = "The `app/` folder MUST have content"
	codes["PLC15002"] = "The `app/` folder content MAY be a `.gitkeep` file"
	codes["PLC15003"] = "The `app/.gitkeep` file, when present, MUST be empty"

	return codes
}

func PLC15(files map[string]string) []message.Message {
	var (
		messages []message.Message
		ok       bool
	)

	codes := listCodes()

	targetFile := "app/"

	if _, ok = files[targetFile]; !ok {
		messages = append(messages, message.CreateMessage(check.Skip, "PLC15001", codes["PLC15001"]))
		messages = append(messages, message.CreateMessage(check.Skip, "PLC15002", codes["PLC15002"]))
		messages = append(messages, message.CreateMessage(check.Skip, "PLC15003", codes["PLC15003"]))
	} else {
		empty := true

		for key := range files {
			path := strings.Split(key, "/")[0] + "/"

			if path == targetFile && key != targetFile {
				empty = false
				break
			}
		}

		if empty {
			// If the folder is empty, we can skip the last check
			messages = append(messages, message.CreateMessage(check.Fail, "PLC15001", codes["PLC15001"]))
			messages = append(messages, message.CreateMessage(check.Skip, "PLC15002", codes["PLC15002"]))
			messages = append(messages, message.CreateMessage(check.Skip, "PLC15003", codes["PLC15003"]))
		} else {
			messages = append(messages, message.CreateMessage(check.Pass, "PLC15001", codes["PLC15001"]))

			if _, ok = files["app/.gitkeep"]; ok {
				messages = append(messages, message.CreateMessage(check.Pass, "PLC15002", codes["PLC15002"]))

				if len(files["app/.gitkeep"]) == 0 {
					messages = append(messages, message.CreateMessage(check.Pass, "PLC15003", codes["PLC15003"]))
				} else {
					messages = append(messages, message.CreateMessage(check.Fail, "PLC15003", codes["PLC15003"]))
				}
			} else {
				messages = append(messages, message.CreateMessage(check.Skip, "PLC15002", codes["PLC15002"]))
				messages = append(messages, message.CreateMessage(check.Skip, "PLC15003", codes["PLC15003"]))
			}
		}
	}

	return messages
}
