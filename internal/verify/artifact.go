package verify

import (
	"fmt"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func VerifyArtifacts(cfg *config.Config) error {

	files := []string{
		"artifacts/mdstat.txt",
		"artifacts/lsblk.txt",
		"artifacts/serial.log",
	}

	for _, f := range files {

		if !common.Exists(f) {
			return fmt.Errorf("missing artifact %s", f)
		}
	}

	return nil
}
