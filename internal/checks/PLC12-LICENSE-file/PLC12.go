package checks

import (
	"internal/check"
	"internal/message"
	"internal/repositoryContents"
	"regexp"
	"strconv"
	"strings"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC12001"] = "The `LICENSE` file MUST be an MIT License"
	codes["PLC12002"] = "The `LICENSE` file MUST contain an attribution line"
	codes["PLC12003"] = "The attribution line MUST contain the year the component was created"
	codes["PLC12004"] = "The copyright year MAY contain a range of years"
	codes["PLC12005"] = "The copyright range of years, when present, MUST be the same as the latest active year"
	codes["PLC12006"] = "The attribution line MUST contain the copyright holder"
	codes["PLC12007"] = "The copyright holder MUST be `pipeline-components` or `Robbert Müller`"

	return codes
}

func PLC12(files map[string]string, repo map[string]string, logs []repositoryContents.LogEntry) []message.Message {
	var (
		messages []message.Message
		ok       bool
	)

	status := map[string]check.Status{}
	codes := listCodes()

	for code := range codes {
		status[code] = check.Skip
	}

	targetFile := "LICENSE"

	if _, ok = files[targetFile]; ok {
		for code := range codes {
			status[code] = check.Fail
		}

		fileContent := files[targetFile]

		if strings.Contains(fileContent, "MIT License") {
			status["PLC12001"] = check.Pass
			lines := strings.Split(fileContent, "\n")
			var attributionLine string

			for _, line := range lines {
				if strings.Contains(line, "(C)") || strings.Contains(line, "(c)") || strings.Contains(line, "Copyright") || strings.ContainsAny(line, "©Ⓒⓒ") {
					attributionLine = line
					break
				}
			}

			if attributionLine == "" {
				status["PLC12002"] = check.Fail
			} else {
				status["PLC12002"] = check.Pass
				status["PLC12003"] = check.Fail

				yearsPattern := regexp.MustCompile("(?P<Year>[0-9]{4})(-(?P<Range>[0-9]{4}))?")

				if yearsPattern.MatchString(attributionLine) {
					yearsMatch := yearsPattern.FindStringSubmatch(attributionLine)

					oldestYear, _ := strconv.Atoi(yearsMatch[yearsPattern.SubexpIndex("Year")])
					newestYear, _ := strconv.Atoi(yearsMatch[yearsPattern.SubexpIndex("Range")])

					firstCommit := logs[0].Timestamp.Year()

					if oldestYear == firstCommit {
						status["PLC12003"] = check.Pass
					}

					lastCommit := logs[len(logs)-1].Timestamp.Year()

					if newestYear != 0 {
						status["PLC12004"] = check.Pass
						status["PLC12005"] = check.Fail

						if newestYear == lastCommit {
							status["PLC12005"] = check.Pass
						}
					}
				}

				copyrightHolderPattern := regexp.MustCompile("[0-9]{4}(:?-[0-9]{4})?(?P<Holder>[^\n]+)")

				if copyrightHolderPattern.MatchString(attributionLine) {

					status["PLC12006"] = check.Pass

					holderMatch := copyrightHolderPattern.FindStringSubmatch(attributionLine)

					holder := holderMatch[copyrightHolderPattern.SubexpIndex("Holder")]

					if strings.Contains(holder, "pipeline-components") || strings.Contains(holder, "Pipeline Components") || strings.Contains(holder, "Robbert Müller") {
						status["PLC12007"] = check.Pass
					}
				}
			}
		}
	}

	for code, checkStatus := range status {
		messages = append(messages, message.CreateMessage(checkStatus, code, codes[code]))
	}

	return messages
}
