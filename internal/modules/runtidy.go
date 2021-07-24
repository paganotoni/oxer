package modules

import (
	"os/exec"
	"oxer/internal/meta"
)

func RunTidy(app meta.Application) error {
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}
