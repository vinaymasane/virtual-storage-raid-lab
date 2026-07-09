package vm

import "github.com/vinaymasane/virtual-storage-raid-lab/internal/common"

func StopVM() error {

	return common.Run(
		"pkill",
		"-f",
		"qemu-system-x86_64",
	)
}