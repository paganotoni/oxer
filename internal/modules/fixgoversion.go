package modules

import (
	"io/ioutil"
	"oxer/internal/meta"
	"regexp"
)

func FixGoVersion(app meta.Application) error {
	cnt, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return err
	}

	var re = regexp.MustCompile(`go [\d.]+`)
	s := re.ReplaceAllString(string(cnt), `go 1.16`)

	return ioutil.WriteFile("go.mod", []byte(s), 0777)
}
