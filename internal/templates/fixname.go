package templates

import (
	"io/fs"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
)

// FixesName of the templates
// - ensures .plush.html
// - removes underscore from partials
func FixName(app meta.Application) error {
	folder := filepath.Join("app", "templates")
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".plush.html") {
			return nil
		}

		if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".plush") {
			np := strings.Replace(filepath.Base(path), ".plush", "", -1)
			np = strings.Replace(np, ".html", "", -1)

			if strings.HasPrefix(np, "_") {
				np = np[1:]
			}

			np = np + ".plush.html"

			np = filepath.Join(filepath.Dir(path), np)
			err := os.Rename(path, np)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
