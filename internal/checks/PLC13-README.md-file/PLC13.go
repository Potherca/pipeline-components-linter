package checks

import (
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"internal/check"
	"internal/message"
	"reflect"
	"strings"
)

const targetFile = "README.md"

func listCodes() map[string]string {
	return map[string]string{
		"PLC13001": "🤖 The `README.md` file MUST pass the linting rules defined in `.mdlrc`",
		"PLC13002": "The `README.md` file MUST contain `# Pipeline Components: <component-name>` heading as the first line",
		"PLC13003": "The lines directly after the heading MUST contain the same badges/shields as the `README.md` file in the skeleton repository",
		"PLC13004": "The `README.md` file MUST contain the same main sections, in the same order, as the `README.md` file in the skeleton repository",
		"PLC13005": "The 'Versioning' section in the `README.md` file MUST be identical to their counterparts in the `README.md` file in the skeleton repository",
		"PLC13006": "The 'Support' section in the `README.md` file MUST be identical to their counterparts in the `README.md` file in the skeleton repository",
		"PLC13007": "The 'Contributing' section in the `README.md` file MUST be identical to their counterparts in the `README.md` file in the skeleton repository",
		"PLC13008": "⁉ The 'Examples' section in the `README.md` file MUST be auto-generated from a separate example file in the repository",
		"PLC13009": "The 'Authors & contributors' section in the `README.md` file MUST state the author who initially set up the repository",
		"PLC13010": "The 'Authors & contributors' section in the `README.md` file MUST link to the contributor's page",
		"PLC13011": "The 'License' section in the `README.md` file MUST state the author who initially set up the repository",
		"PLC13012": "The 'License' section in the `README.md` file MUST state the license type as MIT",
		"PLC13013": "The 'License' section in the `README.md` file MUST link to the license file in the repository",
	}
}

func contentToString(literal []byte, content []byte) string {
	result := ""

	if literal != nil {
		result = string(literal)
	} else if content != nil {
		result = string(content)
	}

	return result
}

func getContentFromNode(node ast.Node) string {
	var content string

	if node.AsContainer() != nil {
		content = contentToString(node.AsContainer().Literal, node.AsContainer().Content)
	} else {
		content = contentToString(node.AsLeaf().Literal, node.AsLeaf().Content)
	}

	return content
}

func getHeadings(document ast.Node, minLevel int, maxLevel int) []struct {
	content string
	index   int
	level   int
} {
	var (
		headings []struct {
			content string
			index   int
			level   int
		}
	)

	headingCount := 0

	ast.WalkFunc(document, func(node ast.Node, entering bool) ast.WalkStatus {
		if entering {
			switch node := node.(type) {
			case *ast.Heading:
				headingCount++
				content := ""

				for _, child := range node.AsContainer().Children {
					content += getContentFromNode(child)
				}

				if node.Level <= maxLevel && node.Level >= minLevel {
					headings = append(headings, struct {
						content string
						index   int
						level   int
					}{
						content: content,
						index:   headingCount,
						level:   node.Level,
					})
				}
			}
		}

		return ast.GoToNext
	})

	return headings
}

func PLC13(files map[string]string, repo map[string]string) []message.Message {
	var (
		messages []message.Message
		ok       bool
	)

	codes := listCodes()
	status := map[string]check.Status{}

	for code := range codes {
		status[code] = check.Skip
	}

	if _, ok = files[targetFile]; ok {
		if _, repoFileExists := repo[targetFile]; !repoFileExists {
			codes["PLC13002"] = fmt.Sprintf("The required `%s` file is missing from the skeleton repository", targetFile)
			status["PLC13002"] = check.Error
		} else {
			for code := range codes {
				if code != "PLC13001" && code != "PLC13008" {
					status[code] = check.Fail
				}
			}

			markdownParser := parser.NewWithExtensions(parser.CommonExtensions)
			markdown := []byte(files[targetFile])
			document := markdownParser.Parse(markdown)

			subjectHeadings := getHeadings(document, 1, 1)

			if len(subjectHeadings) > 0 && strings.HasPrefix(subjectHeadings[0].content, "Pipeline Components: ") {
				status["PLC13002"] = check.Pass
			}

			skeletonParser := parser.NewWithExtensions(parser.CommonExtensions)
			skeletonMarkdown := []byte(repo[targetFile])
			skeletonDocument := skeletonParser.Parse(skeletonMarkdown)
			skeletonHeadings := getHeadings(skeletonDocument, 2, 2)

			subjectHeadings = getHeadings(document, 2, 2)

			if reflect.DeepEqual(subjectHeadings, skeletonHeadings) {
				status["PLC13004"] = check.Pass
			} // else { @TODO: Show the differences between the headings }
		}
	}

	for code, checkStatus := range status {
		messages = append(messages, message.CreateMessage(checkStatus, code, codes[code]))
	}

	return messages
}
