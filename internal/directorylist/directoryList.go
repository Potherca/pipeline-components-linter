package directorylist

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListContent(root string, path string) ([]string, error) {
	var (
		dirEntries []os.DirEntry
		err        error
		files      []string
	)

	path = filepath.Clean(path)
	dirPath := filepath.Join(root, path)
	dirEntries, err = os.ReadDir(dirPath)

	if err == nil {
		for _, file := range dirEntries {
			var pathName string

			dirMarker := ""
			if file.IsDir() {
				dirMarker = "/"
			}

			if path == "" || path == "." {
				pathName = fmt.Sprintf("%s%s", file.Name(), dirMarker)
			} else {
				pathName = fmt.Sprintf("%s/%s%s", path, file.Name(), dirMarker)
			}

			if file.IsDir() {
				var fileNames []string

				fileNames, err = ListContent(root, pathName)

				if err == nil {
					files = append(files, fileNames...)
				}
			}

			files = append(files, pathName)
		}
	}

	return files, err
}
