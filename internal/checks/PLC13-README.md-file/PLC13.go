package checks

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/md"
	"github.com/gomarkdown/markdown/parser"
	"internal/check"
	"internal/message"
	"net/http"
	"reflect"
	"regexp"
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
		"PLC13010": "The 'Authors & contributors' section in the `README.md` file MAY contain a link for the initial author",
		"PLC13011": "The link for the initial author in the 'Authors & contributors' section in the `README.md` file, if present, MUST resolve",
		"PLC13012": "The 'Authors & contributors' section in the `README.md` file MUST link to the contributor's page",
		"PLC13013": "The contributor's page link in the 'Authors & contributors' section in the `README.md` file MUST resolve",
		"PLC13014": "The 'License' section in the `README.md` file MUST attribute Robbert Müller as creator",
		"PLC13015": "The attribution in the 'License' section in the `README.md` file MUST resolve to https://gitlab.com/mjrider",
		"PLC13016": "The 'License' section in the `README.md` file MUST state the license type as MIT",
		"PLC13017": "The 'License' section in the `README.md` file MUST link to the license file in the repository",
		"PLC13018": "The license link in the 'License' section in the `README.md` file MUST resolve",
	}
}

func appendNodeToDocument(parent *ast.Document, child ast.Node) {
	child.SetParent(parent)
	newChildren := append(parent.GetChildren(), child)
	parent.SetChildren(newChildren)
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

func getSections(document ast.Node) map[string]string {
	var (
		currentSectionName string
		sections           map[string]string
	)

	currentSection := &ast.Document{}
	markdownRenderer := md.NewRenderer()
	sections = make(map[string]string)

	currentSectionName = "__ROOT__"

	ast.WalkFunc(document, func(node ast.Node, entering bool) ast.WalkStatus {
		if entering {
			switch node := node.(type) {
			case *ast.Text:
			// Ignore

			case *ast.Heading:
				if node.Level == 1 {
					currentSection = &ast.Document{}
				} else if node.Level == 2 {
					sections[currentSectionName] = string(markdown.Render(currentSection, markdownRenderer))
					currentSection = &ast.Document{}
					currentSectionName = ""

					for _, child := range node.AsContainer().Children {
						currentSectionName += getContentFromNode(child)
					}
				} else {
					appendNodeToDocument(currentSection, node)
				}

			default:
				appendNodeToDocument(currentSection, node)
			}

		} else if _, ok := node.(*ast.Document); ok {
			sections[currentSectionName] = string(markdown.Render(currentSection, markdownRenderer))
		}

		return ast.GoToNext
	})

	return sections
}

func urlResolves(url string) bool {
	resolves := false

	response, err := http.Get(url)

	if err == nil {
		if response.StatusCode >= 200 && response.StatusCode <= 399 {
			resolves = true
		}
	}

	return resolves
}

func PLC13(componentName string, files map[string]string, repo map[string]string) []message.Message {
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
			for _, code := range []string{"PLC13002", "PLC13003", "PLC13004"} {
				status[code] = check.Fail
			}

			subjectParser := parser.NewWithExtensions(parser.CommonExtensions)
			subjectMarkdown := []byte(files[targetFile])
			subjectDocument := subjectParser.Parse(subjectMarkdown)
			subjectSections := getSections(subjectDocument)

			subjectHeadings := getHeadings(subjectDocument, 1, 1)

			if len(subjectHeadings) > 0 && strings.HasPrefix(subjectHeadings[0].content, "Pipeline Components: ") {
				status["PLC13002"] = check.Pass
			}

			skeletonParser := parser.NewWithExtensions(parser.CommonExtensions)
			skeletonMarkdown := []byte(repo[targetFile])
			skeletonDocument := skeletonParser.Parse(skeletonMarkdown)
			skeletonHeadings := getHeadings(skeletonDocument, 2, 2)
			skeletonSections := getSections(skeletonDocument)

			if _, ok := subjectSections["__ROOT__"]; ok && subjectSections["__ROOT__"] == skeletonSections["__ROOT__"] {
				status["PLC13003"] = check.Pass
			}

			subjectHeadings = getHeadings(subjectDocument, 2, 2)

			if reflect.DeepEqual(subjectHeadings, skeletonHeadings) {
				status["PLC13004"] = check.Pass
			} // else { @TODO: Show the differences between the headings }

			for checkCode, sectionName := range map[string]string{
				"PLC13005": "Versioning",
				"PLC13006": "Support",
				"PLC13007": "Contributing",
			} {
				if _, ok := subjectSections[sectionName]; ok {
					if subjectSections[sectionName] == skeletonSections[sectionName] {
						status[checkCode] = check.Pass
					} else {
						status[checkCode] = check.Fail
					}
				}
			}

			linkPattern := regexp.MustCompile(`\[(?P<Subject>[^]]+)\]\((?P<URL>[^)]+)\)`)

			if _, ok := subjectSections["Authors & contributors"]; ok {
				for _, code := range []string{"PLC13009", "PLC13012", "PLC13013"} {
					status[code] = check.Fail
				}

				authorLines := strings.Split(subjectSections["Authors & contributors"], "\n")
				for _, line := range authorLines {

					if strings.Contains(line, "The original setup of this repository is by ") {
						split := strings.Split(line, "by ")

						if len(split) > 1 && split[1] != "" {
							// @TODO: Use git history to validate user? If so, the check wording MUST be changed to reflect this!
							status["PLC13009"] = check.Pass

							if linkPattern.MatchString(split[1]) {
								status["PLC13010"] = check.Pass
								status["PLC13011"] = check.Fail

								matches := linkPattern.FindStringSubmatch(split[1])
								url := matches[linkPattern.SubexpIndex("URL")]

								if urlResolves(url) {
									status["PLC13011"] = check.Pass
								}
							}
						}
					} else if strings.Contains(line, "contributor's page") {
						matches := linkPattern.FindStringSubmatch(line)
						url := matches[linkPattern.SubexpIndex("URL")]

						if linkPattern.MatchString(line) {
							pattern := `https://gitlab.com/pipeline-components/(?:[^/]+/)?(?P<Component>[^/]+)/-/graphs/main`
							contributorsLinkPattern := regexp.MustCompile(pattern)

							if contributorsLinkPattern.MatchString(line) {
								status["PLC13012"] = check.Pass
								status["PLC13013"] = check.Fail

								if urlResolves(url) {
									status["PLC13013"] = check.Pass
								}
							}
						}
					}
				}
			}

			if _, ok := subjectSections["License"]; ok {
				for _, code := range []string{"PLC13014", "PLC13015", "PLC13016", "PLC13017", "PLC13018"} {
					status[code] = check.Fail
				}

				licenseLines := strings.Split(subjectSections["License"], "\n")
				for _, line := range licenseLines {
					if strings.Contains(line, "Created by [Robbert Müller]") {
						status["PLC13014"] = check.Pass
						status["PLC13015"] = check.Fail

						split := strings.Split(line, "by ")

						if len(split) > 1 && split[1] != "" {
							if linkPattern.MatchString(split[1]) {
								matches := linkPattern.FindStringSubmatch(split[1])
								url := matches[linkPattern.SubexpIndex("URL")]

								if urlResolves(url) {
									status["PLC13015"] = check.Pass
								}
							}
						}
					}
					if strings.Contains(line, "licensed under a ") {
						split := strings.Split(line, "under a ")

						if len(split) > 1 && split[1] != "" {
							if linkPattern.MatchString(split[1]) {
								matches := linkPattern.FindStringSubmatch(split[1])
								url := matches[linkPattern.SubexpIndex("URL")]
								subject := matches[linkPattern.SubexpIndex("Subject")]

								if subject == "MIT license" || subject == "MIT License" {
									status["PLC13016"] = check.Pass
								}

								if url == "./LICENSE" {
									status["PLC13017"] = check.Pass

									url = fmt.Sprintf(
										"https://gitlab.com/pipeline-components/%s/-/blob/HEAD/%s",
										componentName,
										url,
									)

									if urlResolves(url) {
										status["PLC13018"] = check.Pass
									}
								}
							}
						}
					}
				}
			}
		}
	}

	for code, checkStatus := range status {
		messages = append(messages, message.CreateMessage(checkStatus, code, codes[code]))
	}

	return messages
}
