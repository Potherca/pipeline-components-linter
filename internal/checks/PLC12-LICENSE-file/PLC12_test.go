package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"internal/repositorycontents"
	"testing"
	"time"
)

func TestPLC12(t *testing.T) {
	targetFile := "LICENSE"

	notMatchingTimestamp := time.Time{}
	matchingFirstTimestamp := time.Date(1234, 1, 1, 0, 0, 0, 0, time.UTC)
	matchingSecondTimestamp := time.Date(5678, 1, 1, 0, 0, 0, 0, time.UTC)

	mockLicense := map[string]string{targetFile: "mock content"}

	tests := map[string]struct {
		files  map[string]string
		logs   []repositorycontents.LogEntry
		repo   map[string]string
		status map[string]check.Status
	}{
		targetFile + " file absent": {
			files: nil,
			logs:  nil,
			repo:  nil,
			status: map[string]check.Status{
				"PLC12001": check.Skip,
				"PLC12002": check.Skip,
				"PLC12003": check.Skip,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Skip,
				"PLC12007": check.Skip,
			},
		},
		targetFile + " file absent from skeleton repo": {
			files: mockLicense,
			logs:  nil,
			repo:  nil,
			status: map[string]check.Status{
				"PLC12001": check.Error,
				"PLC12002": check.Skip,
				"PLC12003": check.Skip,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Skip,
				"PLC12007": check.Skip,
			},
		},
		targetFile + " file present, not MIT License": {
			files: mockLicense,
			logs:  nil,
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Fail,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present and MIT License, but not matching skeleton, without attribution or copyright year": {
			files: map[string]string{targetFile: "MIT License\nPermission is hereby granted foo"},
			logs:  nil,
			repo:  map[string]string{targetFile: "MIT License\nPermission is hereby granted bar"},
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Fail,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present and MIT License, matching skeleton, without attribution or copyright year": {
			files: map[string]string{targetFile: "MIT License\nPermission is hereby granted"},
			logs:  nil,
			repo:  map[string]string{targetFile: "MIT License\nPermission is hereby granted"},
			status: map[string]check.Status{
				"PLC12001": check.Pass,
				"PLC12002": check.Fail,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, with incorrect attribution name, without copyright year": {
			files: map[string]string{targetFile: "(c) Nomen Nescio"},
			logs:  nil,
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Pass,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright year, log absent": {
			files: map[string]string{targetFile: "(c) 1234"},
			logs:  nil,
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Error,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright year, not matching log": {
			files: map[string]string{targetFile: "(c) 1234"},
			logs:  []repositorycontents.LogEntry{{Timestamp: notMatchingTimestamp}},
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright year, matching log": {
			files: map[string]string{targetFile: "(c) 1234"},
			logs:  []repositorycontents.LogEntry{{Timestamp: matchingFirstTimestamp}},
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Pass,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright range, both not matching log": {
			files: map[string]string{targetFile: "(c) 1234-5678"},
			logs:  []repositorycontents.LogEntry{{Timestamp: notMatchingTimestamp}},
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Fail,
				"PLC12004": check.Pass,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright range, first matching log": {
			files: map[string]string{targetFile: "(c) 1234-5678"},
			logs: []repositorycontents.LogEntry{
				{Timestamp: matchingFirstTimestamp},
			},
			repo: mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Pass,
				"PLC12004": check.Pass,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright range, last matching log": {
			files: map[string]string{targetFile: "(c) 1234-5678"},
			logs: []repositorycontents.LogEntry{
				{Timestamp: notMatchingTimestamp},
				{Timestamp: matchingSecondTimestamp},
			},
			repo: mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Fail,
				"PLC12004": check.Pass,
				"PLC12005": check.Pass,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
		targetFile + " file present, not MIT License, with attribution, without attribution name, with copyright range, both matching log": {
			files: map[string]string{targetFile: "(c) 1234-5678"},
			logs: []repositorycontents.LogEntry{
				{Timestamp: matchingFirstTimestamp},
				{Timestamp: matchingSecondTimestamp},
			},
			repo: mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Pass,
				"PLC12004": check.Pass,
				"PLC12005": check.Pass,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		},
	}

	for _, symbol := range []string{"(C)", "(c)", "Copyright", "©", "Ⓒ", "ⓒ"} {
		tests[targetFile+" file present, not MIT License, attribution with '"+symbol+"', without copyright year"] = struct {
			files  map[string]string
			logs   []repositorycontents.LogEntry
			repo   map[string]string
			status map[string]check.Status
		}{
			files: map[string]string{targetFile: symbol},
			logs:  nil,
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Fail,
				"PLC12007": check.Fail,
			},
		}
	}

	for _, name := range []string{"pipeline-components", "Pipeline Components", "Robbert Müller"} {
		tests[targetFile+" file present, not MIT License, with correct attribution '"+name+"', without copyright year"] = struct {
			files  map[string]string
			logs   []repositorycontents.LogEntry
			repo   map[string]string
			status map[string]check.Status
		}{
			files: map[string]string{targetFile: "(c) " + name},
			logs:  nil,
			repo:  mockLicense,
			status: map[string]check.Status{
				"PLC12001": check.Fail,
				"PLC12002": check.Pass,
				"PLC12003": check.Fail,
				"PLC12004": check.Skip,
				"PLC12005": check.Skip,
				"PLC12006": check.Pass,
				"PLC12007": check.Pass,
			},
		}
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC12(test.files, test.repo, test.logs)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status, "%s expected status %v, got %v", message.Code, test.status[message.Code], message.Status)
			}
		})
	}
}
