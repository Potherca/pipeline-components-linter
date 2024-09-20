package checks

import (
	"fmt"
	"internal/check"
	"internal/message"
	"internal/repositorycontents"
	"regexp"
	"strconv"
	"strings"
)

const seekLicenseName = "MIT License"
const seekLicenseText = "Permission is hereby granted"
const targetFile = "LICENSE"

func listCodes() map[string]string {
	return map[string]string{
		"PLC12001": "The `LICENSE` file MUST be an MIT License",
		"PLC12002": "The `LICENSE` file MUST contain an attribution line",
		"PLC12003": "The attribution line MUST contain the year the component was created",
		"PLC12004": "The copyright year MAY contain a range of years",
		"PLC12005": "The copyright range of years, when present, MUST be the same as the latest active year",
		"PLC12006": "The attribution line MUST contain the copyright holder",
		"PLC12007": "The copyright holder MUST be `pipeline-components` or `Robbert Müller`",
	}
}

func getAttributionLine(lines []string) string {
	var attributionLine string

	for _, line := range lines {
		if strings.Contains(line, "(C)") || strings.Contains(line, "(c)") || strings.Contains(line, "Copyright") || strings.ContainsAny(line, "©Ⓒⓒ") {
			attributionLine = line
			break
		}
	}
	return attributionLine
}

func getLicenseHolder(attributionLine string) string {
	var holder string

	copyrightHolderPattern := regexp.MustCompile(`(?:\([cC]\)|Copyright|[©Ⓒⓒ])(?:[0-9-\s]+)?(?P<Holder>[^\n]+)?`)

	if copyrightHolderPattern.MatchString(attributionLine) {

		holderMatch := copyrightHolderPattern.FindStringSubmatch(attributionLine)

		holder = holderMatch[copyrightHolderPattern.SubexpIndex("Holder")]
	}

	return holder
}

func getLicenseTexts(fileContent string, repo map[string]string) (string, string) {
	skeletonContent := strings.Join(strings.Split(repo[targetFile], "\n"), " ")

	index := strings.Index(fileContent, seekLicenseText)
	fileLicenseText := fileContent[index:]

	skeletonIndex := strings.Index(skeletonContent, seekLicenseText)
	skeletonLicenseText := skeletonContent[skeletonIndex:]

	return fileLicenseText, skeletonLicenseText
}

func getLicenseYears(attributionLine string) (int, int) {
	var oldestYear, newestYear int

	yearsPattern := regexp.MustCompile("(?P<Year>[0-9]{4})(-(?P<Range>[0-9]{4}))?")

	if yearsPattern.MatchString(attributionLine) {
		yearsMatch := yearsPattern.FindStringSubmatch(attributionLine)

		oldestYear, _ = strconv.Atoi(yearsMatch[yearsPattern.SubexpIndex("Year")])
		newestYear, _ = strconv.Atoi(yearsMatch[yearsPattern.SubexpIndex("Range")])
	}

	return oldestYear, newestYear
}

func PLC12(files map[string]string, repo map[string]string, logs []repositorycontents.LogEntry) []message.Message {
	var (
		messages []message.Message
		ok       bool
	)

	status := map[string]check.Status{}
	codes := listCodes()

	for code := range codes {
		status[code] = check.Skip
	}

	if _, ok = files[targetFile]; ok {
		if _, repoFileExists := repo[targetFile]; !repoFileExists {
			codes["PLC12001"] = fmt.Sprintf("The required `%s` file is missing from the skeleton repository", targetFile)
			status["PLC12001"] = check.Error
		} else {
			for code := range codes {
				if code != "PLC12004" && code != "PLC12005" {
					status[code] = check.Fail
				}
			}

			fileContent := files[targetFile]

			lines := strings.Split(fileContent, "\n")

			if strings.Contains(fileContent, seekLicenseName) {
				if strings.Contains(fileContent, seekLicenseText) {
					fileLicenseText, skeletonLicenseText := getLicenseTexts(fileContent, repo)

					if fileLicenseText == skeletonLicenseText {
						status["PLC12001"] = check.Pass
					} else {
						// TODO: Add a message that shows the difference between the two files
					}
				}
			}

			attributionLine := getAttributionLine(lines)

			if attributionLine != "" {
				status["PLC12002"] = check.Pass
				status["PLC12003"] = check.Fail

				oldestYear, newestYear := getLicenseYears(attributionLine)

				if oldestYear != 0 {
					if len(logs) == 0 {
						codes["PLC12003"] = fmt.Sprintf("No log entries found for the repository")
						status["PLC12003"] = check.Error
					} else {
						firstCommit := logs[0].Timestamp.Year()

						if oldestYear == firstCommit {
							status["PLC12003"] = check.Pass
						}

						if newestYear != 0 {
							status["PLC12004"] = check.Pass

							lastCommit := logs[len(logs)-1].Timestamp.Year()

							if newestYear == lastCommit {
								status["PLC12005"] = check.Pass
							}
						}
					}
				}

				holder := getLicenseHolder(attributionLine)

				if holder != "" {
					status["PLC12006"] = check.Pass

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
