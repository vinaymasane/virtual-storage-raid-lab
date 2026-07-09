package ansible

import (
	"fmt"
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/vm"
)

func Run() error {

	cfg := common.DefaultConfig()

	if err := vm.WaitForSSH(); err != nil {
		return err
	}

	if err := GenerateInventory(); err != nil {
		return err
	}

	if _, err := os.Stat(cfg.AnsibleDir + "/site.yml"); err != nil {
		return fmt.Errorf("missing site.yml")
	}

	return common.RunDir(
		cfg.AnsibleDir,
		"ansible-playbook",
		"-i",
		"inventory.ini",
		"site.yml",
	)
}