package refresh

import (
	"oxer/internal/meta"
	"path/filepath"

	"github.com/markbates/refresh/refresh"
)

func TweakDevYml(app meta.Application) error {
	config := refresh.Configuration{}
	err := config.Load(".buffalo.dev.yml")
	if err != nil {
		return err
	}

	config.BuildTargetPath = "." + string(filepath.Separator) + filepath.Join("cmd", app.ModuleShort)
	config.LogName = "ox"

	return config.Dump(".buffalo.dev.yml")
}
