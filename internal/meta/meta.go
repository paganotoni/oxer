package meta

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

var (
	ErrModuleNameNotFound = errors.New("module name not found")
)

// Application holds information thats important for the
// transition of the app from buffalo to ox.
type Application struct {
	Module      string
	ModuleShort string
}

func New() (Application, error) {
	app := Application{}
	content, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return app, err
	}

	path := modfile.ModulePath(content)
	name := filepath.Base(path)

	if name == "." {
		return app, ErrModuleNameNotFound
	}

	parts := strings.Split(name, "/")

	app.Module = path
	app.ModuleShort = parts[len(parts)-1]

	return app, nil
}
