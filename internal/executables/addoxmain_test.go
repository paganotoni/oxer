package executables

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
	"testing"
)

func Test_AddOxMain(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	err := AddOxMain(meta.Application{
		Module:      "something/test",
		ModuleShort: "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(filepath.Join(d, "cmd", "ox", "main.go"))
	if err != nil {
		t.Errorf("Should have created main.go in cmd/ox/main.go")
	}

	data, err := ioutil.ReadFile(filepath.Join(d, "cmd", "ox", "main.go"))
	if err != nil {
		t.Errorf("Could not read file created main.go in cmd/ox/main.go")
	}

	if !strings.Contains(string(data), `cli.Use(soda.Plugins(test.Migrations)...)`) {
		t.Errorf("not generated correctly")
	}

}
