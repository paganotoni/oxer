package embed

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
	"testing"
)

func Test_AddEmbed(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	err := AddEmbedFile(meta.Application{
		Module:      "something/test",
		ModuleShort: "test",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(filepath.Join(d, "embed.go"))
	if err != nil {
		t.Errorf("Should have created main.go in embed.go")
	}

	data, err := ioutil.ReadFile(filepath.Join(d, "embed.go"))
	if err != nil {
		t.Errorf("Could not read file created main.go in embed.go")
	}

	if !strings.Contains(string(data), `//go:embed app/templates public migrations config`) {
		t.Errorf("not generated correctly")
	}

}
