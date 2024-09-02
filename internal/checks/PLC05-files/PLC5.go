package checks

import (
	"internal/asserts"
	"internal/message"
)

func listCodes() map[string]string {
	return map[string]string{
		"PLC5001": "The repository MUST contain a `.gitignore` file",
		"PLC5002": "The repository MUST contain a `.gitlab-ci.yml` file",
		"PLC5003": "The repository MUST contain a `.mdlrc` file",
		"PLC5004": "The repository MUST contain a `.yamllint` file",
		"PLC5005": "The repository MUST contain a `action.yml` file",
		"PLC5006": "The repository MUST contain a `Dockerfile` file",
		"PLC5007": "The repository MUST contain a `LICENSE` file",
		"PLC5008": "The repository MUST contain a `README.md` file",
		"PLC5009": "The repository MUST contain a `renovate.json` file",
	}
}

func PLC5(files map[string]string) []message.Message {
	fileCodes := map[string]string{
		".gitignore":     "PLC5001",
		".gitlab-ci.yml": "PLC5002",
		".mdlrc":         "PLC5003",
		".yamllint":      "PLC5004",
		"action.yml":     "PLC5005",
		"Dockerfile":     "PLC5006",
		"LICENSE":        "PLC5007",
		"README.md":      "PLC5008",
		"renovate.json":  "PLC5009",
	}

	return asserts.FileExists(files, fileCodes)
}
