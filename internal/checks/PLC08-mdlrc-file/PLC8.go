package checks

import (
	"bytes"
	"fmt"
	"internal/check"
	"internal/message"
)

func PLC8(files map[string]string, repo map[string]string) []message.Message {
	var (
		checkMessage string
		messages     []message.Message
		status       check.Status
	)

	code := "PLC8001"
	targetFile := ".mdlrc"

	for subjectFile, contents := range repo {
		if subjectFile == targetFile {
			checkMessage = fmt.Sprintf("The `%[1]s` file MUST be identical to `%[1]s` file in the skeleton repository", targetFile)

			if _, ok := files[targetFile]; !ok {
				status = check.Skip
			} else if bytes.Equal([]byte(files[targetFile]), []byte(contents)) {
				status = check.Pass
			} else {
				status = check.Fail
			}

			break
		}
	}

	messages = append(messages, message.CreateMessage(status, code, checkMessage))

	return messages
}
