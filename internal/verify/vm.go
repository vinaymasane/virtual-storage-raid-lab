package verify

import (
	"fmt"

	"github.com/ekagra/virtual-storage-raid-lab/internal/common"
	"github.com/ekagra/virtual-storage-raid-lab/internal/config"
)

func VerifyVM(cfg *config.Config) error {

	out, err := common.Output(
		"virsh",
		"list",
		"--all",
	)

	if err != nil {
		return err
	}

	if !common.Contains(string(out), cfg.VM.Name) {
		return fmt.Errorf("VM not found")
	}

	return nil
}
