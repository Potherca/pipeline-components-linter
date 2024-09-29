package checks

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"internal/check"
	"internal/repositorycontents"
	"io"
	"net/http"
	"testing"
)

var mockDetails = repositorycontents.Details{
	"origin": repositorycontents.RepoDetails{
		Remotes:  []string{"https://gitlab.com/pipeline-components/foo"},
		Branches: []string{"master"},
	},
}

var mockError = errors.New("mock error")

var mockHttpGet = func(url string) (*http.Response, error) {
	return nil, nil
}

func TestPLC2(t *testing.T) {
	tests := map[string]struct {
		details      repositorycontents.Details
		mockFunction func(url string) (*http.Response, error)
		status       map[string]check.Status
	}{
		"Repository does not have remote(s)": {
			details:      nil,
			mockFunction: mockHttpGet,
			status: map[string]check.Status{
				"PLC2001": check.Skip,
				"PLC2002": check.Skip,
				"PLC2003": check.Skip,
				"PLC2004": check.Skip,
				"PLC2005": check.Skip,
			},
		},
		"Repository has remote not under gitlab.com/pipeline-components": {
			details: repositorycontents.Details{
				"origin": repositorycontents.RepoDetails{
					Remotes:  []string{"http://foo/foo.git"},
					Branches: []string{"master"},
				},
			},
			mockFunction: mockHttpGet,
			status: map[string]check.Status{
				"PLC2001": check.Fail,
				"PLC2002": check.Skip,
				"PLC2003": check.Skip,
				"PLC2004": check.Skip,
				"PLC2005": check.Skip,
			},
		},
		"Repository has remote under git@gitlab.com:pipeline-components": {
			details: repositorycontents.Details{
				"origin": repositorycontents.RepoDetails{
					Remotes:  []string{"git@gitlab.com:pipeline-components/foo"},
					Branches: []string{"master"},
				},
			},
			mockFunction: func(url string) (*http.Response, error) {
				return nil, mockError
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Fail,
				"PLC2003": check.Skip,
				"PLC2004": check.Skip,
				"PLC2005": check.Skip,
			},
		},
		"Repository has remote under https://gitlab.com/pipeline-components": {
			details: mockDetails,
			mockFunction: func(url string) (*http.Response, error) {
				return nil, mockError
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Fail,
				"PLC2003": check.Skip,
				"PLC2004": check.Skip,
				"PLC2005": check.Skip,
			},
		},
		"Repository under gitlab.com/pipeline-components is publicly accessible": {
			details: mockDetails,
			mockFunction: func(url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
				}, nil
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Pass,
				"PLC2003": check.Fail,
				"PLC2004": check.Fail,
				"PLC2005": check.Fail,
			},
		},
		"Publicly accessible repository under gitlab.com/pipeline-components without default branch": {
			details: mockDetails,
			mockFunction: func(url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString("[]")),
				}, nil
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Pass,
				"PLC2003": check.Fail,
				"PLC2004": check.Fail,
				"PLC2005": check.Fail,
			},
		},
		"Publicly accessible repository under gitlab.com/pipeline-components with unprotected non-'main' default branch": {
			details: mockDetails,
			mockFunction: func(url string) (*http.Response, error) {
				resp := &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(
						`[{"default": true, "name": "mock-branch", "protected": false, "web_url": "https://gitlab.com/pipeline-components/foo/-/tree/mock-branch"}]`),
					),
				}

				return resp, nil
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Pass,
				"PLC2003": check.Pass,
				"PLC2004": check.Fail,
				"PLC2005": check.Fail,
			},
		},
		"Publicly accessible repository under gitlab.com/pipeline-components with unprotected 'main' default branch": {
			details: mockDetails,
			mockFunction: func(url string) (*http.Response, error) {
				resp := &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(
						`[{"default": true, "name": "main", "protected": false, "web_url": "https://gitlab.com/pipeline-components/foo/-/tree/mock-branch"}]`),
					),
				}

				return resp, nil
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Pass,
				"PLC2003": check.Pass,
				"PLC2004": check.Pass,
				"PLC2005": check.Fail,
			},
		},
		"Publicly accessible repository under gitlab.com/pipeline-components with protected 'main' default branch": {
			details: mockDetails,
			mockFunction: func(url string) (*http.Response, error) {
				resp := &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(
						`[{"default": true, "name": "main", "protected": false, "web_url": "https://gitlab.com/pipeline-components/foo/-/tree/mock-branch"}]`),
					),
				}

				return resp, nil
			},
			status: map[string]check.Status{
				"PLC2001": check.Pass,
				"PLC2002": check.Pass,
				"PLC2003": check.Pass,
				"PLC2004": check.Pass,
				"PLC2005": check.Fail,
			},
		},
	}

	originalFunction := httpGet
	defer func() { httpGet = originalFunction }()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			httpGet = test.mockFunction

			// Act
			messages := PLC2(test.details)

			// Assert
			for _, message := range messages {
				assert.Equal(
					t,
					test.status[message.Code],
					message.Status,
					"%s expected status %v, got %v", message.Code, test.status[message.Code], message.Status,
				)
			}
		})
	}
}
