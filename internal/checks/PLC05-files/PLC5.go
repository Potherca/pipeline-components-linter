package checks

import (
	"internal/asserts"
	"internal/message"
)

func PLC5(files map[string]string) []message.Message {
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

	return asserts.FileExists(files, fileCodes)
}
