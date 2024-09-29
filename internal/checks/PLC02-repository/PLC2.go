package checks

import (
	"encoding/json"
	"internal/check"
	"internal/message"
	"internal/repositorycontents"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type branch struct {
	Default   bool   `json:"default"`
	Name      string `json:"name"`
	Protected bool   `json:"protected"`
	WebUrl    string `json:"web_url"`
}

var httpGet = http.Get

func listCodes() map[string]string {
	return map[string]string{
		"PLC2001": "The repository MUST be hosted under https://gitlab.com/pipeline-components/",
		"PLC2002": "The repository MUST be public",
		"PLC2003": "The repository MUST have a default branch",
		"PLC2004": "The default branch MUST be named `main`",
		"PLC2005": "The `main` branch MUST be protected",
	}
}

func normalizeGitUrl(url string) string {
	if strings.HasPrefix(url, "git@") {
		url = strings.Replace(url, ":", "/", 1)
		url = strings.Replace(url, "git@", "https://", 1)
	}

	return url
}

func PLC2(repoDetails repositorycontents.Details) []message.Message {
	var (
		messages []message.Message
	)

	status := map[string]check.Status{}
	codes := listCodes()

	for code := range codes {
		status[code] = check.Skip
	}

	for _, details := range repoDetails {
		for _, remoteUrl := range details.Remotes {
			status["PLC2001"] = check.Fail
			remoteUrl = normalizeGitUrl(remoteUrl)

			if strings.HasPrefix(remoteUrl, "https://gitlab.com/pipeline-components/") {
				project, _ := strings.CutPrefix(remoteUrl, "https://gitlab.com/pipeline-components/")
				status["PLC2001"] = check.Pass
				status["PLC2002"] = check.Fail

				apiUrl := "https://gitlab.com/api/v4/projects/" + url.QueryEscape(project) + "/repository/branches"

				response, err := httpGet(apiUrl)

				if err == nil {
					if response.StatusCode >= 200 && response.StatusCode <= 399 {
						status["PLC2002"] = check.Pass
						status["PLC2003"] = check.Fail
						status["PLC2004"] = check.Fail
						status["PLC2005"] = check.Fail

						var bodyBytes []byte

						if response.Body != nil {
							bodyBytes, err = io.ReadAll(response.Body)
						}

						var branches []branch

						err = json.Unmarshal(bodyBytes, &branches)

						if err == nil {
							for _, branch := range branches {
								if branch.Default == true {
									status["PLC2003"] = check.Pass

									if branch.Name == "main" {
										status["PLC2004"] = check.Pass
									}
									if branch.Protected == true {
										status["PLC2005"] = check.Pass
									}
								}
							}
						}
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
