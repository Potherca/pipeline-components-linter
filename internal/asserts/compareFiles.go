package asserts

import (
	"bytes"
	"fmt"
	"internal/check"
	"internal/message"
)

func CompareFiles(
	files map[string]string,
	repo map[string]string,
	fileCodes map[string]string,
) []message.Message {
	var (
		checkMessage string
		messages     []message.Message
		status       check.Status
	)

	for targetFile := range fileCodes {
		if _, repoFileExists := repo[targetFile]; !repoFileExists {
			status = check.Error
			checkMessage = fmt.Sprintf("The required `%[1]s` file is missing from the skeleton repository", targetFile)
		} else {
			for subjectFile, contents := range repo {
				if subjectFile == targetFile {
					checkMessage = fmt.Sprintf("The `%[1]s` file MUST be identical to `%[1]s` file in the skeleton repository", targetFile)

					if _, targetFileExists := files[targetFile]; !targetFileExists {
						status = check.Skip
					} else if bytes.Equal([]byte(files[targetFile]), []byte(contents)) {
						status = check.Pass
					} else {
						status = check.Fail
					}

					break
				}
			}
		}

		messages = append(messages, message.CreateMessage(status, fileCodes[targetFile], checkMessage))
	}

	return messages
}
