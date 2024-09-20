package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"strings"
	"testing"
)

const mockCorrectHeader = "# Pipeline Components: Mock\n"
const mockIncorrectHeader = "# Mock Header\n"

const mockCorrectSections = `
## Usage
## Examples
## Versioning
{{ .versioning_section }}
## Support
{{ .support_section }}
## Contributing
{{ .contributing_section }}
## Authors & contributors
{{ .author_section }}
## License
{{ .license_section }}
`

const mockIncorrectSections = `
## First Section
### Sub-section
## Second Section
`

const mockBadges = "[![A](B)](C)\n"
const mockOtherBadges = "[![D](E)](F)\n"

func populateTemplate(templateContent string, replace map[string]string) string {
	for A, B := range replace {
		templateContent = strings.Replace(templateContent, "{{ ."+A+" }}", B, -1)
	}

	return templateContent
}

func TestPLC13(t *testing.T) {
	tests := map[string]struct {
		files  map[string]string
		repo   map[string]string
		status map[string]check.Status
	}{
		targetFile + " file absent": {
			files: nil,
			repo:  nil,
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Skip,
				"PLC13003": check.Skip,
				"PLC13004": check.Skip,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Skip,
				"PLC13013": check.Skip,
				"PLC13014": check.Skip,
				"PLC13015": check.Skip,
				"PLC13016": check.Skip,
				"PLC13017": check.Skip,
				"PLC13018": check.Skip,
			},
		},
		targetFile + " file absent from skeleton repo": {
			files: map[string]string{targetFile: "# Mock File content"},
			repo:  nil,
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Error,
				"PLC13003": check.Skip,
				"PLC13004": check.Skip,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Skip,
				"PLC13013": check.Skip,
				"PLC13014": check.Skip,
				"PLC13015": check.Skip,
				"PLC13016": check.Skip,
				"PLC13017": check.Skip,
				"PLC13018": check.Skip,
			},
		},
		targetFile + " with different badges, different sections": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockIncorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Fail,
				"PLC13004": check.Fail,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Skip,
				"PLC13013": check.Skip,
				"PLC13014": check.Skip,
				"PLC13015": check.Skip,
				"PLC13016": check.Skip,
				"PLC13017": check.Skip,
				"PLC13018": check.Skip,
			},
		},
		targetFile + " with correct header": {
			files: map[string]string{targetFile: mockCorrectHeader + mockBadges + mockIncorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Pass,
				"PLC13003": check.Fail,
				"PLC13004": check.Fail,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Skip,
				"PLC13013": check.Skip,
				"PLC13014": check.Skip,
				"PLC13015": check.Skip,
				"PLC13016": check.Skip,
				"PLC13017": check.Skip,
				"PLC13018": check.Skip,
			},
		},
		targetFile + " with incorrect section contents": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "",
					"contributing_section": "",
					"license_section":      "",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Fail,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with matching badges, different sections": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockIncorrectSections},

			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Fail,
				"PLC13005": check.Skip,
				"PLC13006": check.Skip,
				"PLC13007": check.Skip,
				"PLC13008": check.Skip,
				"PLC13009": check.Skip,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Skip,
				"PLC13013": check.Skip,
				"PLC13014": check.Skip,
				"PLC13015": check.Skip,
				"PLC13016": check.Skip,
				"PLC13017": check.Skip,
				"PLC13018": check.Skip,
			},
		},
		targetFile + " with correct content": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Pass,
				"PLC13006": check.Pass,
				"PLC13007": check.Pass,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with author without link": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "The original setup of this repository is by Mock Author",
					"contributing_section": "",
					"license_section":      "",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Pass,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with author link unresolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "The original setup of this repository is by [Mock Author](https://httpbin.org/status/500)",
					"contributing_section": "",
					"license_section":      "",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Pass,
				"PLC13010": check.Pass,
				"PLC13011": check.Fail,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with author link resolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "The original setup of this repository is by [Mock Author](https://httpbin.org/status/200)",
					"contributing_section": "",
					"license_section":      "",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Pass,
				"PLC13010": check.Pass,
				"PLC13011": check.Pass,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with contributor link unresolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "For a full list of all authors and contributors, check [the contributor's page][contributors].\n[contributors]: https://gitlab.com/pipeline-components/_template_/-/graphs/main",
					"contributing_section": "",
					"license_section":      "",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Pass,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with contributor link resolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "For a full list of all authors and contributors, check [the contributor's page][contributors].\n[contributors]: https://gitlab.com/pipeline-components/org/skeleton/-/graphs/main",
					"contributing_section": "",
					"license_section":      "",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Pass,
				"PLC13013": check.Pass,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with attribution link unresolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "",
					"contributing_section": "",
					"license_section":      "Created by [Robbert Müller][mjrider]",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Pass,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with attribution link resolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "",
					"contributing_section": "",
					"license_section":      "Created by [Robbert Müller](https://gitlab.com/mjrider)",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Pass,
				"PLC13015": check.Pass,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with wrong license": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "",
					"contributing_section": "",
					"license_section":      "licensed under a [Mozilla Public License][license-link] (MPL).",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Fail,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with correct license unresolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "",
					"contributing_section": "",
					"license_section":      "licensed under a [MIT license](license-link)",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Pass,
				"PLC13017": check.Fail,
				"PLC13018": check.Fail,
			},
		},
		targetFile + " with correct license resolved": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":       "",
					"contributing_section": "",
					"license_section":      "licensed under a [MIT license][license-link] (MIT).\n[license-link]: ./LICENSE",
					"support_section":      "",
					"versioning_section":   "",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Skip,
				"PLC13011": check.Skip,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
				"PLC13014": check.Fail,
				"PLC13015": check.Fail,
				"PLC13016": check.Pass,
				"PLC13017": check.Pass,
				"PLC13018": check.Pass,
			},
		},
		targetFile + " with correct header, matching badges, correct content, author link resolved, contributor link resolved, attribution link resolved, correct license resolved": {
			files: map[string]string{targetFile: mockCorrectHeader + mockBadges +
				populateTemplate(mockSections, map[string]string{
					"author_section":  "The original setup of this repository is by [Mock Author](https://httpbin.org/status/200)\nFor a full list of all authors and contributors, check [the contributor's page][contributors].\n[contributors]: https://gitlab.com/pipeline-components/_template_/-/graphs/main",
					"license_section": "Created by [Robbert Müller][mjrider], licensed under a [MIT license][license-link]\n[mjrider]: https://gitlab.com/mjrider\n[license-link]: ./LICENSE\n",
				}),
			},
			repo: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Pass,
				"PLC13003": check.Pass,
				"PLC13004": check.Pass,
				"PLC13005": check.Pass,
				"PLC13006": check.Pass,
				"PLC13007": check.Pass,
				"PLC13008": check.Skip,
				"PLC13009": check.Pass,
				"PLC13010": check.Pass,
				"PLC13011": check.Pass,
				"PLC13012": check.Pass,
				"PLC13013": check.Pass,
				"PLC13014": check.Pass,
				"PLC13015": check.Pass,
				"PLC13016": check.Pass,
				"PLC13017": check.Pass,
				"PLC13018": check.Pass,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC13("org/skeleton", test.files, test.repo)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status, "%s expected status %v, got %v", message.Code, test.status[message.Code], message.Status)
			}
		})
	}
}
