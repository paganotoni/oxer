package main

import (
	"oxer/internal/embed"
	"oxer/internal/executables"
	"oxer/internal/folders"
	"oxer/internal/meta"
	"oxer/internal/modules"
	"oxer/internal/refresh"
	"oxer/internal/templates"
	"oxer/internal/webpack"
)

// steps that should be taken to tranform the project
// source into ox expected format.
type step func(meta.Application) error

func main() {
	app, err := meta.New()
	if err != nil {
		panic(err)
	}

	steps := []step{
		executables.MoveMain,
		executables.AddOxMain,
		embed.AddEmbedFile,

		// Move folders into app moving app and
		// render to be outside the actions folder
		folders.MoveIntoApp,
		folders.MoveApp,
		folders.ReplacePackages,
		folders.MoveGrifts,

		//folders.MoveRender,

		// Fix templates names and partial() sentences
		// since fsbox does not auto-append .html.plush
		// And removing underscore because it is not supported
		// by embed.
		templates.FixName,
		templates.FixPartialSentence,

		// Replacing Packr sentences
		// embed.ReplacePackr,

		// Replace routes in webpack.config.js
		webpack.ReplaceAssetsRoutes,

		// Replace things in .buffalo.dev.yml
		refresh.TweakDevYml,

		// Replacing docker file
		// docker.ReplaceDockerfile,

		modules.FixGoVersion,
		modules.RunTidy,
	}

	for _, task := range steps {
		err := task(app)
		if err != nil {
			panic(err)
		}
	}
}
