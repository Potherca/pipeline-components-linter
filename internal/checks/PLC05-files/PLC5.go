package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() (Codes map[string]string) {
	Codes["PLC5001"] = "The repository MUST contain a `.gitignore` file"
	Codes["PLC5002"] = "The repository MUST contain a `.gitlab-ci.yml` file"
	Codes["PLC5003"] = "The repository MUST contain a `.mdlrc` file"
	Codes["PLC5004"] = "The repository MUST contain a `.yamllint` file"
	Codes["PLC5005"] = "The repository MUST contain a `action.yml` file"
	Codes["PLC5006"] = "The repository MUST contain a `Dockerfile` file"
	Codes["PLC5007"] = "The repository MUST contain a `LICENSE` file"
	Codes["PLC5008"] = "The repository MUST contain a `README.md` file"
	Codes["PLC5009"] = "The repository MUST contain a `renovate.json` file"

	return Codes
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
