package folders

import (
	"io/fs"
	"io/ioutil"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
)

// Replace packages once those are moved into app
func ReplacePackages(app meta.Application) error {
	folder := filepath.Join(".")
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		packages := []string{
			"actions",
			"mailers",
			"middleware",
			"models",
			"tasks",
		}
		s := string(content)
		for _, pkg := range packages {
			orig := strings.Join([]string{app.Module, pkg}, "/")
			dest := strings.Join([]string{app.Module, "app", pkg}, "/")

			s = strings.ReplaceAll(s, orig, dest)
		}

		err = ioutil.WriteFile(path, []byte(s), info.Mode())
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
