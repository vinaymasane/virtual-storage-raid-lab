package vm

import "github.com/vinaymasane/virtual-storage-raid-lab/internal/common"

// StopVM stops the virtual machine by terminating the QEMU process.
func StopVM() error {

	return common.Run(
		"pkill",
		"-f",
		"qemu-system-x86_64",
	)
}
