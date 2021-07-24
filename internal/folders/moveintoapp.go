package folders

import (
	"fmt"
	"os"
	"oxer/internal/meta"
	"path/filepath"
)

// Move app folders into app/
func MoveIntoApp(app meta.Application) error {
	folders := []string{
		"actions",
		"assets",
		"mailers",
		"middleware",
		"models",
		"tasks",
		"locales",
		"templates",
	}

	err := os.MkdirAll("app", 0777)
	if err != nil {
		return fmt.Errorf("error creating app folder:%w", err)
	}

	for _, folder := range folders {
		_, err := os.Stat(filepath.Join(folder))
		if err != nil {
			//In case the folder doesn't exist
			//continue with the next folder
			continue
		}

		err = os.Rename(folder, filepath.Join("app", folder))
		if err != nil {
			return fmt.Errorf("error moving %s into app:%w", folder, err)
		}
	}

	return nil
}
