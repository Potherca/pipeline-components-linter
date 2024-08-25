package checks

import (
	"internal/asserts"
	"internal/check"
	"internal/message"
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

	status := map[string]check.Status{}
	codes := listCodes()

	status["PLC15001"] = check.Fail
	status["PLC15002"] = check.Skip
	status["PLC15003"] = check.Skip

	targetFile := "app/"

	if _, ok = files[targetFile]; !ok {
		status["PLC15001"] = check.Skip
	} else {
		empty := asserts.DirectoryIsEmpty(files, targetFile)

		if !empty {
			status["PLC15001"] = check.Pass
			if _, ok = files["app/.gitkeep"]; ok {
				status["PLC15002"] = check.Pass

				if len(files["app/.gitkeep"]) == 0 {
					status["PLC15003"] = check.Pass
				} else {
					status["PLC15003"] = check.Fail
				}
			}
		}
	}

	for code, checkStatus := range status {
		messages = append(messages, message.CreateMessage(checkStatus, code, codes[code]))
	}

	return messages
}
