package main

import (
	"fmt"
	"internal/check"
	plc4 "internal/checks/PLC04-folders"
	plc5 "internal/checks/PLC05-files"
	"internal/message"
	"os"
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

	var messages []message.Message

	var messageMarker = message.Marker{
		Pass: "✅",
		Fail: "❌",
	}

	fileNames, err := listFiles(path)

	if err == nil {

		messages = append(messages, plc4.PLC4(fileNames)...)
		messages = append(messages, plc5.PLC5(fileNames)...)

		for _, checkMessage := range messages {
			var marker string

			switch checkMessage.Status {
			case check.Pass:
				marker = messageMarker.Pass
			case check.Fail:
				marker = messageMarker.Fail
			default:
				errorMessage := fmt.Sprintf("Unknown or unsupported CheckStatus '%v'", checkMessage.Status)
				panic(errorMessage)
			}

			fmt.Printf("%s %s %s\n", checkMessage.Code, marker, checkMessage.Message)
		}
	}
}
