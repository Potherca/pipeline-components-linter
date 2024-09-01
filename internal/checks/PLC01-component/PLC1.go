package checks

import (
	"internal/check"
	"internal/message"
	repo "internal/repositoryContents"
	"path/filepath"
	"regexp"
	"strings"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC1001"] = "The Pipeline Component MUST live in its own folder"
	codes["PLC1002"] = "The Pipeline Component folder MUST be named after the main component it exposes"
	codes["PLC1003"] = "The Pipeline Component folder MUST be a git repository"

	return codes
}

func PLC1(projectPath string, files map[string]string, repoLogs []repo.LogEntry) []message.Message {
	var (
		messages []message.Message
	)

	status := map[string]check.Status{}
	codes := listCodes()

	for code := range codes {
		status[code] = check.Skip
	}

	if projectPath != "" && len(files) > 0 {
		status["PLC1001"] = check.Pass
	}

	if _, fileExists := files["Dockerfile"]; fileExists {
		dockerfileLines := strings.Split(files["Dockerfile"], "\n")
		for _, line := range dockerfileLines {
			if strings.Contains(line, "ENV DEFAULTCMD") {
				mainCommandPattern := regexp.MustCompile(`^\s*ENV DEFAULTCMD\s*=?\s*["']?(?P<Command>.+?)["']?$`)
				if mainCommandPattern.MatchString(line) {
					status["PLC1002"] = check.Fail

					matches := mainCommandPattern.FindStringSubmatch(line)
					mainCommand := matches[mainCommandPattern.SubexpIndex("Command")]

					if mainCommand == filepath.Base(projectPath) {
						status["PLC1002"] = check.Pass
					}
				}
			}
		}

	}

	if repoLogs == nil {
		status["PLC1003"] = check.Fail
	} else {
		status["PLC1003"] = check.Pass
	}

	for code, checkStatus := range status {
		messages = append(messages, message.CreateMessage(checkStatus, code, codes[code]))
	}

	return messages
}
