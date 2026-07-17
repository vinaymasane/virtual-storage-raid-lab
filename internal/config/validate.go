package config

import (
	"errors"
)

func Validate(cfg *Config) error {

	switch {

	case cfg.VM.Name == "":
		return errors.New("vm.name missing")

	case cfg.Image.Output == "":
		return errors.New("image.output missing")

	case cfg.Storage.RAIDDevice == "":
		return errors.New("storage.raid_device missing")

	case cfg.SSH.User == "":
		return errors.New("ssh.user missing")
	}

	return nil
}
