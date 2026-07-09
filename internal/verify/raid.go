package verify

import "github.com/vinaymasane/virtual-storage-raid-lab/internal/common"

func CheckRaid() error {

	if err := common.Run(
		"bash",
		"-c",
		"cat /proc/mdstat",
	); err != nil {

		return err

	}

	return common.Run(
		"lsblk",
		"-f",
	)
}
