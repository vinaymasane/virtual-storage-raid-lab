package common

import (
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func PreFlight(cfg *config.Config) error {

	if err := EnsureDirectories(cfg); err != nil {
		Error("Ensure directories check failed: " + err.Error())
		return err
	}

	if err := InitLogger(cfg); err != nil {
		Error("Logger initialization failed: " + err.Error())
		return err
	}

	if err := CheckDependencies(); err != nil {
		Error("Dependency check failed: " + err.Error())
		return err
	}

	return nil
}
