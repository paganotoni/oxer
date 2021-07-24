package webpack

import (
	"io/ioutil"
	"oxer/internal/meta"
	"strings"
)

// Replaces assets routes from the webpack file
// given now we use a different folder structure where we
// have assets inside app.
func ReplaceAssetsRoutes(app meta.Application) error {
	cnt, err := ioutil.ReadFile("webpack.config.js")
	if err != nil {
		return err
	}

	s := strings.ReplaceAll(string(cnt), "./assets", "./app/assets")
	s = strings.ReplaceAll(s, ".\\/assets", ".\\/app\\/assets")

	return ioutil.WriteFile("webpack.config.js", []byte(s), 0777)
}
