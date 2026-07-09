package verify

import (
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

// CollectArtifacts collects various system artifacts and saves them to the specified artifact directory.
func CollectArtifacts() error {

	cfg := common.DefaultConfig()

	os.MkdirAll(
		cfg.ArtifactDir+"/reports",
		0755,
	)

	common.Run(
		"bash",
		"-c",
		"cat /proc/mdstat > artifacts/reports/mdstat.txt",
	)

	common.Run(
		"bash",
		"-c",
		"lsblk -f > artifacts/reports/lsblk.txt",
	)

	common.Run(
		"bash",
		"-c",
		"ip addr > artifacts/reports/ipaddr.txt",
	)

	return nil
}
