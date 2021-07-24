package templates

import (
	"io/fs"
	"io/ioutil"
	"oxer/internal/meta"
	"path/filepath"
	"regexp"
	"strings"
)

// FixPartialSentence the name of the templates
// - ensures .plush.html
// - removes underscore from partials
func FixPartialSentence(app meta.Application) error {
	folder := filepath.Join("app", "templates")
	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".plush.html") {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var re = regexp.MustCompile(`<%= partial\("([^\.]*)(\.plush|\.html)`)
		s := re.ReplaceAllString(string(content), `<%= partial("$1.plush.html`)

		err = ioutil.WriteFile(path, []byte(s), info.Mode())
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
