package templates

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"testing"
)

func TestFixPartialSentence(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	cases := []struct {
		original string
		expected string
	}{
		{`<%= partial("goals/form.html") %>`, `<%= partial("goals/form.plush.html") %>`},
		{`<%= partial("goals/form.html", {some: var, and: "value"}) %>`, `<%= partial("goals/form.plush.html", {some: var, and: "value"}) %>`},
		{`<%= partial("goals/form.html") %><%= partial("goals/form.html", {some: var, and: "value"}) %>`, `<%= partial("goals/form.plush.html") %><%= partial("goals/form.plush.html", {some: var, and: "value"}) %>`},
	}

	folder := filepath.Join("app", "templates")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		t.Fatalf("Failed creating the folder: %v", err)
	}

	for _, tcase := range cases {
		t.Run(tcase.original, func(t *testing.T) {
			err := ioutil.WriteFile(filepath.Join(folder, "file.plush.html"), []byte(tcase.original), 0777)
			if err != nil {
				t.Errorf("error creating file for %v: %v", tcase.original, err)
			}

			FixPartialSentence(meta.Application{})
			output, err := ioutil.ReadFile(filepath.Join(folder, "file.plush.html"))
			if err != nil {
				t.Fatal(err)
			}

			if string(output) != tcase.expected {
				t.Errorf("expected %v, got %v", tcase.expected, string(output))
			}
		})
	}
}
