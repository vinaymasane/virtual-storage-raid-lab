package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func CreateMirror() error {

	cfg := common.DefaultConfig()

	if err := common.Run(
		"modprobe",
		"nbd",
		"max_part=8",
	); err != nil {
		return err
	}

	if _, err := os.Stat(cfg.ImageFile); err != nil {
		return fmt.Errorf("missing image %s", cfg.ImageFile)
	}

	if err := common.Run(
		"qemu-nbd",
		"--disconnect",
		"/dev/nbd0",
	); err != nil {
		// Ignore if not connected
	}

	if err := common.Run(
		"qemu-nbd",
		"--connect=/dev/nbd0",
		cfg.ImageFile,
	); err != nil {
		return err
	}

	if _, err := os.Stat(cfg.SecondaryDisk); os.IsNotExist(err) {

		if err := common.Run(
			"qemu-img",
			"create",
			"-f",
			"raw",
			cfg.SecondaryDisk,
			"10G",
		); err != nil {
			return err
		}
	}

	loop := "/dev/loop10"

	if err := common.Run(
		"losetup",
		loop,
		cfg.SecondaryDisk,
	); err != nil {
		return err
	}

	return nil
}