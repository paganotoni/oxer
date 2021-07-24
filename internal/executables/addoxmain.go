package executables

import (
	_ "embed"
	"fmt"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"text/template"
)

var (
	//go:embed templates/ox/main.go.tmpl
	oxmain string
)

// AddOxMain adds cmd/ox/main.go for further customization
func AddOxMain(app meta.Application) error {
	tmpl := template.New("oxmain")
	tmpl, err := tmpl.Parse(oxmain)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	err = os.MkdirAll(filepath.Join("cmd", "ox"), 0777)
	if err != nil {
		return fmt.Errorf("error creating folder: %w", err)
	}

	f, err := os.Create(filepath.Join("cmd", "ox", "main.go"))
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	err = tmpl.Execute(f, app)
	if err != nil {
		return fmt.Errorf("error exec: %w", err)
	}

	return nil
}
