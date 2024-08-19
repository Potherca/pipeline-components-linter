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

type CommandError struct {
	code    int
	message string
}

func getPath(projectPath string) (string, CommandError) {
	var err error

	commandError := CommandError{
		code:    exitcodes.Ok,
		message: "",
	}

	projectPath, err = filepath.Abs(projectPath)

	if err != nil {
		commandError = CommandError{
			code:    exitcodes.UnknownErrorOccurred,
			message: fmt.Sprintf("could not get absolute path for '%s': %v", projectPath, err),
		}
	} else {
		var fileInfo os.FileInfo

		fileInfo, err = os.Stat(projectPath)

		if err != nil {
			if os.IsNotExist(err) {
				commandError = CommandError{
					code:    exitcodes.CouldNotFindDirectory,
					message: fmt.Sprintf("provided path '%s' does not exist", projectPath),
				}
			} else {
				commandError = CommandError{
					code:    exitcodes.UnknownErrorOccurred,
					message: fmt.Sprintf("could not stat path '%s': %v", projectPath, err),
				}
			}
		} else if !fileInfo.IsDir() {
			commandError = CommandError{
				code:    exitcodes.InvalidParameter,
				message: fmt.Sprintf("provided path '%s' is not a directory", projectPath),
			}
		} // else: projectPath is an existing directory
	}

	return projectPath, commandError
}

func listFiles(path string) ([]string, CommandError) {
	var fileNames []string

	commandError := CommandError{
		code:    exitcodes.Ok,
		message: "",
	}

	files, err := os.ReadDir(path)

	if err != nil {
		commandError = CommandError{
			code:    exitcodes.CouldNotRead,
			message: fmt.Sprintf("could not read files from '%s': %w", path, err),
		}
	} else {
		for _, file := range files {
			name := file.Name()
			if file.IsDir() {
				name += "/"
			}

			fileNames = append(fileNames, name)
		}
	}

	return fileNames, commandError
}

func main() {
	projectPath := "."
	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	}
	projectPath, pathError := getPath(projectPath)

	if pathError.code != exitcodes.Ok {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", pathError.message)
		os.Exit(pathError.code)
	}

	fileNames, fileListError := listFiles(projectPath)
	if fileListError.code != exitcodes.Ok {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", fileListError.message)
		os.Exit(fileListError.code)
	}

	var checks []message.Message

	// @TODO: Markers should be overridable from a configuration file
	messageMarker := message.Marker{
		Pass:       "✅",
		Fail:       "❌",
		Skip:       "⚠️",
		Incomplete: "❓",
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
