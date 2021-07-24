package folders

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
	"testing"
)

func TestMoveGrifts(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	// create grifts folder
	err := os.Mkdir(filepath.Join(d, "grifts"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(d, "grifts", "task.go"), []byte("package grifts"), 0777)
	if err != nil {
		t.Error(err)
	}

	err = MoveGrifts(meta.Application{})
	if err != nil {
		t.Fatal(err)
	}

	// check if grifts folder is moved"
	_, err = os.Stat(filepath.Join("app", "tasks"))
	if err != nil {
		t.Error(err)
	}

	// check if task.go is moved
	_, err = os.Stat(filepath.Join("app", "tasks", "task.go"))
	if err != nil {
		t.Error(err)
	}

	// check task.go content
	b, err := ioutil.ReadFile(filepath.Join("app", "tasks", "task.go"))

	if err != nil {
		t.Error(err)
	}

	if strings.Contains(string(b), "grifts") {
		t.Error("grifts package was not changed")
	}
}
