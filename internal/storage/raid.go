package storage

import (
	"os"
	"time"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func CreateRaid() error {

	cfg := common.DefaultConfig()

	if _, err := os.Stat("/dev/nbd0"); err != nil {
		return err
	}

	if err := common.Run(
		"mdadm",
		"--create",
		cfg.RaidDevice,
		"--level=1",
		"--raid-devices=2",
		"/dev/nbd0",
		"/dev/loop10",
		"--force",
		"--run",
	); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	if err := common.Run(
		"partprobe",
		cfg.RaidDevice,
	); err != nil {
		return err
	}

	if err := common.Run(
		"udevadm",
		"settle",
	); err != nil {
		return err
	}

	if err := common.Run(
		"e2fsck",
		"-fy",
		cfg.RaidDevice,
	); err != nil {
		return err
	}

	if err := common.Run(
		"resize2fs",
		cfg.RaidDevice,
	); err != nil {
		return err
	}

	if err := common.Run(
		"mdadm",
		"--detail",
		cfg.RaidDevice,
	); err != nil {
		return err
	}

	return nil
}
