package vm

import (
	"fmt"
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

// LaunchVM launches the virtual machine using qemu-system-x86_64
func LaunchVM() error {

	cfg := common.DefaultConfig()

	if _, err := os.Stat(cfg.ImageFile); err != nil {
		return fmt.Errorf("image not found: %s", cfg.ImageFile)
	}

	consoleLog := cfg.ArtifactDir + "/console/ttyS0.log"

	args := []string{
		"-enable-kvm",
		"-machine", "q35",
		"-cpu", "host",
		"-smp", "2",
		"-m", "4096",

		"-drive",
		fmt.Sprintf("file=%s,format=qcow2,if=virtio", cfg.ImageFile),

		"-nographic",

		"-serial",
		fmt.Sprintf("file:%s", consoleLog),

		"-netdev",
		"user,id=n1,hostfwd=tcp::2222-:22",

		"-device",
		"virtio-net-pci,netdev=n1",
	}

	return common.StartBackground(
		"qemu-system-x86_64",
		args...,
	)
}
