package main

import (
	"fmt"
	"internal/check"
	"internal/message"
	"os"
	"slices"
)

func listFiles(path string) (fileNames []string, err error) {
	files, err := os.ReadDir(path)

	if err == nil {
		for _, file := range files {
			name := file.Name()
			if file.IsDir() {
				name += "/"
			}
			fileNames = append(fileNames, name)
		}

	}

	return fileNames, err
}

func main() {
	path := "."

	fileNames, err := listFiles(path)

	var fileCodes = make(map[string]string)

	fileCodes["app/"] = "PLC4001"
	fileCodes[".github/"] = "PLC4002"
	fileCodes[".gitignore"] = "PLC5001"
	fileCodes[".gitlab-ci.yml"] = "PLC5002"
	fileCodes[".mdlrc"] = "PLC5003"
	fileCodes[".yamllint"] = "PLC5004"
	fileCodes["action.yml"] = "PLC5005"
	fileCodes["Dockerfile"] = "PLC5006"
	fileCodes["LICENSE"] = "PLC5007"
	fileCodes["README.md"] = "PLC5008"
	fileCodes["renovate.json"] = "PLC5009"

	if err == nil {
		requiredFiles := [11]string{
			".github/",
			"app/",
			".gitignore",
			".gitlab-ci.yml",
			".mdlrc",
			".yamllint",
			"action.yml",
			"Dockerfile",
			"LICENSE",
			"README.md",
			"renovate.json",
		}

		for _, file := range requiredFiles {
			fileType := "file"

			if file[len(file)-1:] == "/" {
				fileType = "directory"
			}

			status := check.Fail
			if slices.Contains(fileNames, file) {
				status = check.Pass
			}

			checkMessage := fmt.Sprintf("The repository MUST contain a %s %s", file, fileType)
			code := fileCodes[file]

			var messageMarker = message.Marker{
				Pass: "✅",
				Fail: "❌",
			}

			fmt.Print(message.CreateMessage(status, code, checkMessage, messageMarker))
		}
	}
}
