package embed

import (
	_ "embed"

	"fmt"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"text/template"
)

var (
	//go:embed templates/embed.go.tmpl
	embed string
)

// AddEmbedFile adds embed.go on the root folder so it holds
// the embedded files in the project for the build process.
func AddEmbedFile(app meta.Application) error {
	tmpl := template.New("embed")
	tmpl, err := tmpl.Parse(embed)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	f, err := os.Create(filepath.Join("embed.go"))
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	err = tmpl.Execute(f, app)
	if err != nil {
		return fmt.Errorf("error exec: %w", err)
	}

	return nil
}
