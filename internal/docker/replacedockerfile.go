package docker

import (
	_ "embed"

	"fmt"
	"html/template"
	"os"
	"oxer/internal/meta"
)

var (
	//go:embed templates/Dockerfile.tmpl
	dockerfile string
)

// Replace existing dockerfile with something that does not depend on
// buffalo binary or docker image.
func ReplaceDockerfile(app meta.Application) error {
	tmpl := template.New("dockerfile")
	tmpl, err := tmpl.Parse(dockerfile)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	f, err := os.Create("Dockerfile")
	if err != nil {
		return fmt.Errorf("error opening Dockerfile: %w", err)
	}

	err = tmpl.Execute(f, app)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
