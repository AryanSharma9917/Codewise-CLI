package env

import (
	"fmt"
	"os"
)

type DeleteOptions struct {
	Force bool // mapped from --yes
}

func DeleteEnv(name string, opts DeleteOptions) error {
	if !envExists(name) {
		return fmt.Errorf("environment %q does not exist", name)
	}

	dir, err := envDir(name)
	if err != nil {
		return err
	}

	return os.RemoveAll(dir)
}
