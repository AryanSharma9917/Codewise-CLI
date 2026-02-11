package env

import (
	"os"
)

// ListEnvs returns all available environments
func ListEnvs() ([]Env, error) {

	base, err := baseEnvPath()
	if err != nil {
		return nil, err
	}

	// If no environments exist yet,
	// return empty slice instead of error.
	if _, err := os.Stat(base); os.IsNotExist(err) {
		return []Env{}, nil
	}

	entries, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}

	var envs []Env

	for _, entry := range entries {

		if !entry.IsDir() {
			continue
		}

		name := entry.Name()

		e, err := LoadEnv(name)
		if err != nil {
			// skip corrupted envs safely
			continue
		}

		envs = append(envs, *e)
	}

	return envs, nil
}
