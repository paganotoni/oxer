package docker

import (
	"errors"
	"oxer/internal/meta"
)

// Replace existing dockerfile with something that does not depend on
// buffalo binary or docker image.
func ReplaceDockerfile(app meta.Application) error {
	return errors.New("not implemented")
}
