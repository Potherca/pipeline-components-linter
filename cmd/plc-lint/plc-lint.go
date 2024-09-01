package main

import (
	"fmt"
	"internal/check"
	"regexp"
	"sort"

	// plc1 "internal/checks/PLC1-component"
	// plc2 "internal/checks/PLC2-repository"
	// plc3 "internal/checks/PLC3-commits"
	plc4 "internal/checks/PLC04-folders"
	plc5 "internal/checks/PLC05-files"
	// plc6 "internal/checks/PLC6-gitignore-file
	// plc7 "internal/checks/PLC7-gitlab-ci.yml-file"
	plc8 "internal/checks/PLC08-mdlrc-file"
	plc9 "internal/checks/PLC09-yamllint-file"
	// plc10 "internal/checks/PLC10-action.yml-file"
	// plc11 "internal/checks/PLC11-Dockerfile"
	plc12 "internal/checks/PLC12-LICENSE-file"
	// plc13 "internal/checks/PLC13-README.md-file"
	plc14 "internal/checks/PLC14-renovate.json-file"
	plc15 "internal/checks/PLC15-app-folder"
	plc16 "internal/checks/PLC16-github-folder"
	plc17 "internal/checks/PLC17-FUNDING.yml-file"
	plc18 "internal/checks/PLC18-github-workflows-folder"
	plc19 "internal/checks/PLC19-release.yml-file"
	// plc20 "internal/checks/PLC20-examples-folder"
	"internal/directoryList"
	"internal/exitcodes"
	"internal/message"
	repo "internal/repositoryContents"
	"os"
	"path/filepath"
	"strings"
)

type CommandError struct {
	code    int
	message string
}

func CreateCommandError(
	code int,
	message string,
) CommandError {
	return CommandError{
		code:    code,
		message: message,
	}
}

func getFileList(projectPath string) map[string]string {
	files, fileListError := loadFiles(projectPath)

	if fileListError.code != exitcodes.Ok {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", fileListError.message)
		os.Exit(fileListError.code)
	}

	return files
}

func getMarkerForStatus(messageStatus check.Status, messageMarker message.Marker) string {
	var marker string

	switch messageStatus {
	case check.Pass:
		marker = messageMarker.Pass
	case check.Fail:
		marker = messageMarker.Fail
	case check.Skip:
		marker = messageMarker.Skip
	case check.Incomplete:
		marker = messageMarker.Incomplete
	default:
		errorMessage := fmt.Sprintf("Unknown or unsupported CheckStatus '%v'", messageStatus)
		panic(errorMessage)
	}

	return marker
}

func getMessageMarkers() message.Marker {
	// @TODO: Markers should be overridable from a configuration file
	messageMarker := message.Marker{
		Pass:       "✅",
		Fail:       "❌",
		Skip:       "⏭ ",
		Incomplete: "⚠️",
	}

	return messageMarker
}

func getPath(projectPath string) (string, CommandError) {
	var err error

	commandError := CreateCommandError(exitcodes.Ok, "")

	projectPath, err = filepath.Abs(projectPath)

	if err != nil {
		commandError = CreateCommandError(
			exitcodes.UnknownErrorOccurred,
			fmt.Sprintf("could not get absolute path for '%s': %v", projectPath, err))
	} else {
		var fileInfo os.FileInfo

		fileInfo, err = os.Stat(projectPath)

		if err != nil {
			if os.IsNotExist(err) {
				commandError = CreateCommandError(
					exitcodes.CouldNotFindDirectory,
					fmt.Sprintf("provided path '%s' does not exist", projectPath))
			} else {
				commandError = CreateCommandError(
					exitcodes.UnknownErrorOccurred,
					fmt.Sprintf("could not stat path '%s': %v", projectPath, err))
			}
		} else if !fileInfo.IsDir() {
			commandError = CreateCommandError(
				exitcodes.InvalidParameter,
				fmt.Sprintf("provided path '%s' is not a directory", projectPath))
		} // else: projectPath is an existing directory
	}

	return projectPath, commandError
}

func getProjectPath() string {
	projectPath := "."

	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	}

	projectPath, pathError := getPath(projectPath)

	if pathError.code != exitcodes.Ok {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", pathError.message)
		os.Exit(pathError.code)
	}

	return projectPath
}

func loadFiles(path string) (map[string]string, CommandError) {
	var fileMap = make(map[string]string)

	commandError := CreateCommandError(exitcodes.Ok, "")

	files, err := directoryList.ListContent(path, "")

	if err != nil {
		commandError = CreateCommandError(
			exitcodes.CouldNotRead,
			fmt.Sprintf("could not read files from '%s': %v", path, err))
	} else {
		for _, file := range files {
			if strings.HasSuffix(file, "/") {
				fileMap[file] = "__DIR__"
			} else {
				contentPath := filepath.Join(path, file)
				contents, err := os.ReadFile(contentPath)

				if err != nil {
					commandError = CreateCommandError(
						exitcodes.CouldNotRead,
						fmt.Sprintf("could not read file '%s': %v", file, err))

					break
				}

				fileMap[file] = string(contents)
			}
		}
	}

	return fileMap, commandError
}

func loadRepoLogs(path string) []repo.LogEntry {
	repoLogs, err := repo.GetLogs(path)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitcodes.UnknownErrorOccurred)
	}

	return repoLogs
}

func loadSkeletonFileList() map[string]string {
	var (
		fileListError   CommandError
		repoError       CommandError
		skeletonContent map[string]string
	)

	if len(os.Args) > 2 {
		skeletonPath := os.Args[2]
		skeletonPath, pathError := getPath(skeletonPath)

		if pathError.code != exitcodes.Ok {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", pathError.message)
			os.Exit(pathError.code)
		}

		skeletonContent, fileListError = loadFiles(skeletonPath)

		if fileListError.code != exitcodes.Ok {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", fileListError.message)
			os.Exit(fileListError.code)
		}
	} else {
		skeletonContent, repoError = loadSkeletonRepoContent("https://gitlab.com/pipeline-components/org/skeleton.git")

		if repoError.code != exitcodes.Ok {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", repoError.message)
			os.Exit(repoError.code)
		}
	}

	return skeletonContent
}

func loadSkeletonRepoContent(repoPath string) (map[string]string, CommandError) {
	commandError := CreateCommandError(exitcodes.Ok, "")

	files, err := repo.GetContent(repoPath)

	if err != nil {
		commandError = CreateCommandError(
			exitcodes.UnknownErrorOccurred,
			fmt.Sprintf("could not get content from '%s': %v", repoPath, err),
		)
	}

	return files, commandError
}

func printMessages(checks []message.Message) {
	var checkMessages []string

	messageMarkers := getMessageMarkers()

	for _, checkMessage := range checks {
		checkMessages = append(checkMessages, fmt.Sprintf("%s %s %s\n", checkMessage.Code, getMarkerForStatus(checkMessage.Status, messageMarkers), checkMessage.Message))
	}

	sortMessages(checkMessages)

	_, err := fmt.Fprint(os.Stdout, strings.Join(checkMessages, ""))

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exitcodes.CouldNotUpdate)
	}
}

func runChecks(files map[string]string, skeletonContent map[string]string, repoLogs []repo.LogEntry) []message.Message {
	var checks []message.Message

	checks = append(checks, plc4.PLC4(files)...)
	checks = append(checks, plc5.PLC5(files)...)
	checks = append(checks, plc8.PLC8(files, skeletonContent)...)
	checks = append(checks, plc9.PLC9(files, skeletonContent)...)
	checks = append(checks, plc12.PLC12(files, skeletonContent, repoLogs)...)
	checks = append(checks, plc14.PLC14(files, skeletonContent)...)
	checks = append(checks, plc15.PLC15(files)...)
	checks = append(checks, plc16.PLC16(files)...)
	checks = append(checks, plc17.PLC17(files, skeletonContent)...)
	checks = append(checks, plc18.PLC18(files)...)
	checks = append(checks, plc19.PLC19(files, skeletonContent)...)

	return checks
}

func sortMessages(checkMessages []string) []string {
	// Change all instance of PLC\d{1}\d{3} to PLC0$1$2
	re := regexp.MustCompile(`PLC0?(\d{1})(\d{3}) `)

	for i, checkMessage := range checkMessages {
		checkMessages[i] = re.ReplaceAllString(checkMessage, "PLC0$1$2 ")
	}

	sort.Strings(checkMessages)

	// Change all instance of PLC0\d{1}\d{2} back to PLC$1$2
	for i, checkMessage := range checkMessages {
		checkMessages[i] = re.ReplaceAllString(checkMessage, "PLC$1$2 ")
	}

	return checkMessages
}

func main() {
	projectPath := getProjectPath()

	files := getFileList(projectPath)
	skeletonContent := loadSkeletonFileList()
	repoLogs := loadRepoLogs(projectPath)

	checks := runChecks(files, skeletonContent, repoLogs)

	printMessages(checks)
}
