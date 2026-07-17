package ansible

import (
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func Verify() error {

	if err := common.CheckDependency("ansible-playbook"); err != nil {
		common.Error("ansible-playbook not installed")
		return err
	}

	return nil
}
