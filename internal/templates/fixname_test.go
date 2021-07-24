package templates

import (
	"fmt"
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"testing"
)

func TestFixName(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	cases := []struct {
		original string
		expected string
	}{
		{"a.plush", "a.plush.html"},
		{"b.plush.html", "b.plush.html"},
		{"c.html", "c.plush.html"},
		{"c.something.plush.html", "c.something.plush.html"},
		{"d.something.html", "d.something.plush.html"},
		{".doc", ".doc"},
		{".something", ".something"},
		{"img.png", "img.png"},
		{"plush-html", "plush-html"},
		{"_partial.plush.html", "partial.plush.html"},
		{"_partial.html", "partial.plush.html"},
		{filepath.Join("folder", "a.plush"), filepath.Join("folder", "a.plush.html")},
		{filepath.Join("folder", "sub", "a.plush"), filepath.Join("folder", "sub", "a.plush.html")},
	}

	//Create those files originally
	for _, c := range cases {
		path := filepath.Join("app", "templates", filepath.Dir(c.original))
		err := os.MkdirAll(path, 0777)
		if err != nil {
			t.Fatal("could not create the path")
		}

		err = ioutil.WriteFile(filepath.Join("app", "templates", c.original), []byte("<h1><h2>"), 0777)
		if err != nil {
			fmt.Println(err)
			t.Fatalf("could not create the file: %v : %v", c.original, err)
		}
	}

	FixName(meta.Application{})

	for _, c := range cases {
		t.Run(c.original, func(t *testing.T) {
			content, err := ioutil.ReadFile(filepath.Join("app", "templates", c.expected))
			if err != nil {
				t.Errorf("did not find file: %v", c.expected)
			}

			if string(content) != "<h1><h2>" {
				t.Errorf("content is not correct: %v", string(content))
			}
		})
	}
}
