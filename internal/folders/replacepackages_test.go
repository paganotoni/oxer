package folders

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"testing"
)

func TestReplacePackageNames(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	os.Remove("app/file.go")

	var original = `
		package main

		package main

		import (
			"log"

			"github.com/maddesa/hnib/actions"
			"github.com/maddesa/hnib/models"
			"github.com/maddesa/hnib/models/other"
			"github.com/maddesa/hnib/actions/something"
			_ "github.com/maddesa/hnib/internal/notouch"
		)

		func main() {
			app := actions.App()
			if err := app.Serve(); err != nil {
				log.Fatal(err)
			}
		}
	`

	err := os.MkdirAll("app", 0777)
	if err != nil {
		t.Fatalf("Failed creating the folder: %v", err)
	}

	err = ioutil.WriteFile(filepath.Join("app", "file.go"), []byte(original), 0777)
	if err != nil {
		t.Errorf("error creating file for %v: %v", original, err)
	}

	ReplacePackages(meta.Application{})

	var expected = `
		package main

		package main

		import (
			"log"

			"github.com/maddesa/hnib/app/actions"
			"github.com/maddesa/hnib/app/models"
			"github.com/maddesa/hnib/app/models/other"
			"github.com/maddesa/hnib/app/actions/something"
			_ "github.com/maddesa/hnib/internal/notouch"
		)

		func main() {
			app := actions.App()
			if err := app.Serve(); err != nil {
				log.Fatal(err)
			}
		}
	`

	output, err := ioutil.ReadFile(filepath.Join("app", "file.go"))
	if err != nil {
		t.Fatal(err)
	}

	if string(output) != expected {
		t.Errorf("expected %v, got %v", expected, string(output))
	}
}
