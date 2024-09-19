package checks

import (
	"github.com/stretchr/testify/assert"
	"internal/check"
	"testing"
)

const mockCorrectHeader = "# Pipeline Components: Mock\n"
const mockIncorrectHeader = "# Mock Header\n"

const mockCorrectSections = `
## Usage
## Examples
## Versioning
## Support
## Contributing
## Authors & contributors
## License
`

const mockIncorrectSections = `
## First Section
## Second Section
`

const mockBadges = "[![A](B)](C)\n"
const mockOtherBadges = "[![D](E)](F)\n"

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
			},
		},
		targetFile + " file present, incorrect header, different badges, different sections": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockIncorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockCorrectSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Fail,
				"PLC13004": check.Fail,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Fail,
				"PLC13011": check.Fail,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
			},
		},
		targetFile + " file present, correct header, different badges, different sections": {
			files: map[string]string{targetFile: mockCorrectHeader + mockBadges + mockIncorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockCorrectSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Pass,
				"PLC13003": check.Fail,
				"PLC13004": check.Fail,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Fail,
				"PLC13011": check.Fail,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
			},
		},
		targetFile + " file present, incorrect header, different badges, same sections": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockCorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockCorrectSections},
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
				"PLC13010": check.Fail,
				"PLC13011": check.Fail,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
			},
		},
		targetFile + " file present, correct header, different badges, same sections": {
			files: map[string]string{targetFile: mockCorrectHeader + mockBadges + mockCorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockOtherBadges + mockCorrectSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Pass,
				"PLC13003": check.Fail,
				"PLC13004": check.Pass,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Fail,
				"PLC13011": check.Fail,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
			},
		},
		targetFile + " file present, incorrect header, matching badges, different sections": {
			files: map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockIncorrectSections},
			repo:  map[string]string{targetFile: mockIncorrectHeader + mockBadges + mockCorrectSections},
			status: map[string]check.Status{
				"PLC13001": check.Skip,
				"PLC13002": check.Fail,
				"PLC13003": check.Pass,
				"PLC13004": check.Fail,
				"PLC13005": check.Fail,
				"PLC13006": check.Fail,
				"PLC13007": check.Fail,
				"PLC13008": check.Skip,
				"PLC13009": check.Fail,
				"PLC13010": check.Fail,
				"PLC13011": check.Fail,
				"PLC13012": check.Fail,
				"PLC13013": check.Fail,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			messages := PLC13(test.files, test.repo)

			for _, message := range messages {
				assert.Equal(t, test.status[message.Code], message.Status, "%s expected status %v, got %v", message.Code, test.status[message.Code], message.Status)
			}
		})
	}
}
