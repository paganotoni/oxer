package node

import (
	"encoding/json"
	"io/ioutil"
	"oxer/internal/meta"
)

// TweakPackageJSON ensures it has the node scripts.
func TweakPackageJSON(app meta.Application) error {
	cnt, err := ioutil.ReadFile("package.json")
	if err != nil {
		return err
	}

	cnf := Config{}
	err = json.Unmarshal(cnt, &cnf)
	if err != nil {
		return err
	}

	if cnf.Scripts == nil {
		cnf.Scripts = make(map[string]interface{})
	}

	cnf.Scripts["build"] = "webpack --mode production --no-stats"
	cnf.Scripts["dev"] = "webpack --watch"

	b, err := json.MarshalIndent(cnf, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("package.json", b, 0777)
	if err != nil {
		return err
	}

	return nil
}
