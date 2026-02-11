package deploy

import (
	"fmt"

	"github.com/aryansharma9917/codewise-cli/pkg/env"
)

func LoadEnvironment(name string) (*env.Env, error) {

	e, err := env.LoadEnv(name)
	if err != nil || e == nil {
		return nil, fmt.Errorf(
			"environment %q not found. run 'codewise env list' to see available environments",
			name,
		)
	}

	return e, nil
}
