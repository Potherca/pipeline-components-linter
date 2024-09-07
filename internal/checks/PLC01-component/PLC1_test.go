package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	repo "internal/repositorycontents"
	"testing"
)

func TestPLC1(t *testing.T) {
	tests := map[string]struct {
		projectPath string
		files       map[string]string
		repoLogs    []repo.LogEntry
		status      map[string]check.Status
	}{
		"Path is empty": {
			projectPath: "",
			files:       map[string]string{},
			repoLogs:    nil,
			status: map[string]check.Status{
				"PLC1001": check.Skip,
				"PLC1002": check.Skip,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided but files is empty": {
			projectPath: "some/path",
			files:       map[string]string{},
			repoLogs:    nil,
			status: map[string]check.Status{
				"PLC1001": check.Skip,
				"PLC1002": check.Skip,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided and files exist, but no Dockerfile": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "mock content",
			},
			repoLogs: nil,
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Skip,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided, Dockerfile exist, but Dockerfile has no ENV DEFAULTCMD": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "mock content",
			},
			repoLogs: nil,
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Skip,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided, Dockerfile exist, and Dockerfile has an invalid ENV DEFAULTCMD": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "ENV DEFAULTCMD",
			},
			repoLogs: nil,
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Skip,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided, Dockerfile exist, Dockerfile has valid ENV DEFAULTCMD that does not match the folder name": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "ENV DEFAULTCMD foo",
			},
			repoLogs: nil,
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Fail,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided, Dockerfile exist, Dockerfile has valid ENV DEFAULTCMD that matches the folder name": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "ENV DEFAULTCMD path",
			},
			repoLogs: nil,
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Pass,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided, Dockerfile exist, Dockerfile has matching ENV DEFAULTCMD, but no repo logs": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "ENV DEFAULTCMD path",
			},
			repoLogs: nil,
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Pass,
				"PLC1003": check.Fail,
			},
		},
		"Path is provided, Dockerfile exist, Dockerfile has matching ENV DEFAULTCMD, and repo logs exist": {
			projectPath: "some/path",
			files: map[string]string{
				"Dockerfile": "ENV DEFAULTCMD path",
			},
			repoLogs: []repo.LogEntry{},
			status: map[string]check.Status{
				"PLC1001": check.Pass,
				"PLC1002": check.Pass,
				"PLC1003": check.Pass,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC1(test.projectPath, test.files, test.repoLogs)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status)
			}
		})
	}
}
