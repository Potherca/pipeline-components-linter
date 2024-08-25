package asserts

import "strings"

func DirectoryIsEmpty(files map[string]string, targetFile string) bool {
	empty := true

	for key := range files {
		path := strings.Split(key, "/")[0] + "/"

		if path == targetFile && key != targetFile {
			empty = false
			break
		}
	}
	return empty
}
