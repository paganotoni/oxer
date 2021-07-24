package folders

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
)

// Move actions/app.go to live inside /app folder, not inside actions
func MoveApp(app meta.Application) error {
	destination := filepath.Join("app", "app.go")
	err := os.Rename(
		filepath.Join("app", "actions", "app.go"),
		destination,
	)

	if err != nil {
		return fmt.Errorf("error moving app/actions/app.go:%w", err)
	}

	// Replace the actions package with "app"
	content, err := ioutil.ReadFile(destination)
	if err != nil {
		return fmt.Errorf("error reading app/app.go:%w", err)
	}

	content = bytes.Replace(content, []byte("package actions"), []byte("package app"), 1)
	err = ioutil.WriteFile(destination, content, 0777)
	if err != nil {
		return fmt.Errorf("error writing app/app.go:%w", err)
	}

	return nil
}
