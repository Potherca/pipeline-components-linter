package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	codes := make(map[string]string)

	codes["PLC5001"] = "The repository MUST contain a `.gitignore` file"
	codes["PLC5002"] = "The repository MUST contain a `.gitlab-ci.yml` file"
	codes["PLC5003"] = "The repository MUST contain a `.mdlrc` file"
	codes["PLC5004"] = "The repository MUST contain a `.yamllint` file"
	codes["PLC5005"] = "The repository MUST contain a `action.yml` file"
	codes["PLC5006"] = "The repository MUST contain a `Dockerfile` file"
	codes["PLC5007"] = "The repository MUST contain a `LICENSE` file"
	codes["PLC5008"] = "The repository MUST contain a `README.md` file"
	codes["PLC5009"] = "The repository MUST contain a `renovate.json` file"

	return codes
}

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
