package deploy

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func LoadEnvironment(name string) (*env.Env, error) {

	e, err := env.LoadEnv(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment: %w", err)
	}

	return e, nil
}
