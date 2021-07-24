package executables

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"testing"
)

func Test_MoveMain(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	err := ioutil.WriteFile(filepath.Join(d, "main.go"), []byte(`package main`), 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = MoveMain(meta.Application{
		Module:      "something/test",
		ModuleShort: "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(filepath.Join(d, "cmd", "test", "main.go"))
	if err != nil {
		t.Errorf("Should have moved main.go to cmd/test/main.go")
	}
}
