package main

import (
	"fmt"
	"internal/check"
	checks "internal/checks/PLC5"
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

	if err == nil {
		requiredFiles := [2]string{
			".github/",
			"app/",
		}

		for _, file := range requiredFiles {
			status := check.Fail
			if slices.Contains(fileNames, file) {
				status = check.Pass
			}

			checkMessage := fmt.Sprintf("The repository MUST contain a %s directory", file)
			code := fileCodes[file]

			var messageMarker = message.Marker{
				Pass: "✅",
				Fail: "❌",
			}

			fmt.Print(message.CreateMessage(status, code, checkMessage, messageMarker))
		}

		messages := checks.PLC5(fileNames)
		for _, checkMessage := range messages {
			fmt.Print(checkMessage)
		}
	}
}
