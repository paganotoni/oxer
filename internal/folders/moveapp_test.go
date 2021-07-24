package folders

import (
	"fmt"
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
	"testing"
)

func TestMoveApp(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	err := os.MkdirAll(filepath.Join(d, "app", "actions"), 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(d, "app", "actions", "app.go"), []byte("package actions\nfunc main() {}"), 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = MoveApp(meta.Application{
		Module:      "something/test",
		ModuleShort: "test",
	})

	if err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadFile(filepath.Join(d, "app", "app.go"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(content))

	if !strings.Contains(string(content), "package app") {
		t.Errorf("app.go should contain 'package app'")
	}

}
