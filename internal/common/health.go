package common

import "fmt"

func PreFlight() error {
	if err := CheckDependencies(); err != nil {
		return err
	}

	dirs := []string{
		"artifacts",
		"output",
	}

	for _, d := range dirs {
		if err := EnsureDir(d); err != nil {
			return err
		}
	}

	return nil
}