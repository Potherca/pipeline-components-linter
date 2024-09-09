package checks

import (
	"fmt"
	"internal/check"
	"internal/message"
)

const targetFile = "README.md"

func listCodes() map[string]string {
	return map[string]string{
		"PLC13001": "🤖 The `README.md` file MUST pass the linting rules defined in `.mdlrc`",
		"PLC13002": "The `README.md` file MUST contain the same sections, in the same order, as the `README.md` file in the skeleton repository",
		"PLC13003": "The `README.md` file MUST contain `# Pipeline Components: <component-name>` heading as the first line",
		"PLC13004": "The lines directly after the heading MUST contain the same badges/shields as the `README.md` file in the skeleton repository",
		"PLC13005": "The 'Versioning', 'Support', and 'Contributing' sections in the `README.md` file MUST be identical to their counterparts in the `README.md` file in the skeleton repository",
		"PLC13006": "⁉ The 'Examples' section in the `README.md` file MUST be auto-generated from a separate example file in the repository",
		"PLC13007": "The 'Authors & contributors' section in the `README.md` file MUST state the author who initially set up the repository",
		"PLC13008": "The 'Authors & contributors' section in the `README.md` file MUST link to the contributor's page",
		"PLC13009": "The 'License' section in the `README.md` file MUST state the author who initially set up the repository",
		"PLC13010": "The 'License' section in the `README.md` file MUST state the license type as MIT",
		"PLC13011": "The 'License' section in the `README.md` file MUST link to the license file in the repository",
	}
}

func PLC13(files map[string]string, repo map[string]string) []message.Message {
	var (
		messages []message.Message
		ok       bool
	)

	codes := listCodes()
	status := map[string]check.Status{}

	for code := range codes {
		status[code] = check.Skip
	}

	if _, ok = files[targetFile]; ok {
		if _, repoFileExists := repo[targetFile]; !repoFileExists {
			codes["PLC13002"] = fmt.Sprintf("The required `%s` file is missing from the skeleton repository", targetFile)
			status["PLC13002"] = check.Error
		} else {
			for code := range codes {
				if code != "PLC13001" {
					status[code] = check.Fail
				}
			}
		}
	}

	for code, checkStatus := range status {
		messages = append(messages, message.CreateMessage(checkStatus, code, codes[code]))
	}

	return messages
}
