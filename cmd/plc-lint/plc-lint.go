package main

import (
	"fmt"
	"internal/check"
	plc4 "internal/checks/PLC04-folders"
	plc5 "internal/checks/PLC05-files"
	"internal/exitcodes"
	"internal/message"
	"os"
	"path/filepath"
)

type commandError struct {
	code    int
	message string
}

func getPath(projectPath string) (string, commandError) {
	userError := commandError{
		code:    0,
		message: "",
	}

	projectPath, _ = filepath.Abs(projectPath)
	fileInfo, err := os.Stat(projectPath)

	if err != nil {
		if os.IsNotExist(err) {
			userError = commandError{
				code:    exitcodes.CouldNotFindDirectory,
				message: fmt.Sprintf("provided path '%s' does not exist", projectPath),
			}
		} else {
			userError = commandError{
				code:    exitcodes.UnknownErrorOccurred,
				message: fmt.Sprintf("could not stat path '%s': %v", projectPath, err),
			}
		}
	} else {
		if !fileInfo.IsDir() {
			userError = commandError{
				code:    exitcodes.InvalidParameter,
				message: fmt.Sprintf("provided path '%s' is not a directory", projectPath),
			}
		}
	}

	return projectPath, userError
}

func listFiles(path string) ([]string, error) {
	var fileNames []string

	files, err := os.ReadDir(path)

	if err != nil {
		err = fmt.Errorf("could not read files from '%s': %w", path, err)
	} else {
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
	projectPath := "."

	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	}

	projectPath, pathError := getPath(projectPath)

	if pathError.code != 0 {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", pathError.message)
		os.Exit(pathError.code)
	}

	var checks []message.Message

	// @TODO: Markers should be overridable from a configuration file
	messageMarker := message.Marker{
		Pass:       "✅",
		Fail:       "❌",
		Skip:       "⚠️",
		Incomplete: "❓",
	}

	fileNames, err := listFiles(projectPath)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	checks = append(checks, plc4.PLC4(fileNames)...)
	checks = append(checks, plc5.PLC5(fileNames)...)

	for _, checkMessage := range checks {
		var marker string

		switch checkMessage.Status {
		case check.Pass:
			marker = messageMarker.Pass
		case check.Fail:
			marker = messageMarker.Fail
		case check.Skip:
			marker = messageMarker.Skip
		case check.Incomplete:
			marker = messageMarker.Incomplete
		default:
			errorMessage := fmt.Sprintf("Unknown or unsupported CheckStatus '%v'", checkMessage.Status)
			panic(errorMessage)
		}

		_, _ = fmt.Fprintf(os.Stdout, "%s %s %s\n", checkMessage.Code, marker, checkMessage.Message)
	}
}
