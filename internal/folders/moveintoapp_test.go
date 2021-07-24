package folders

import (
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"testing"
)

func TestMoveIntoApp(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	// Create folders
	folders := []string{
		"actions",
		"assets",
		"mailers",
		"middleware",
		"models",
		"templates",
	}

	for _, fld := range folders {
		os.MkdirAll(filepath.Join(d, fld), 0777)
	}

	extrafolders := []string{
		// Should not be copied
		"extraone",
		"notcopied",
	}

	for _, fld := range extrafolders {
		os.MkdirAll(filepath.Join(d, fld), 0777)
	}

	err := MoveIntoApp(meta.Application{
		Module:      "something/test",
		ModuleShort: "test",
	})

	if err != nil {
		t.Error(err)
	}

	for _, fld := range folders {
		_, err := os.Stat(filepath.Join(d, "app", fld))
		if err == nil {
			continue
		}

		t.Error(err)
	}
}
