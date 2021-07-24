package folders

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
)

// Move grifts into app/tasks
func MoveGrifts(app meta.Application) error {
	err := os.MkdirAll(filepath.Join("app"), 0755)
	if err != nil {
		return err
	}
	dest := filepath.Join("app", "tasks")
	err = os.Rename("grifts", dest)
	if err != nil {
		return fmt.Errorf("error moving %s into app: %v", "grifts", err)
	}

	// Rename the package in there.
	err = filepath.Walk(dest, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		s := strings.ReplaceAll(string(content), "package grifts", "package tasks")
		err = ioutil.WriteFile(path, []byte(s), info.Mode())
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
