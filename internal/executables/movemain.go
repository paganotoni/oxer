package executables

import (
	"fmt"
	"os"
	"oxer/internal/meta"
	"path/filepath"
)

// Moves the main.go file into cmd/[module]/main.go
func MoveMain(app meta.Application) error {
	binaryFolder := filepath.Join("cmd", app.ModuleShort)
	err := os.MkdirAll(binaryFolder, 0777)
	if err != nil {
		return fmt.Errorf("could not create folder:%w", err)
	}

	err = os.Rename("main.go", filepath.Join(binaryFolder, "main.go"))
	if err != nil {
		return fmt.Errorf("could not move main.go:%w", err)
	}

	return nil
}
