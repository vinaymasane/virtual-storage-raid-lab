package ansible

import (
	"fmt"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func Verify() error {

	if err := common.CheckDependency("ansible-playbook"); err != nil {
		return fmt.Errorf("ansible-playbook not installed")
	}

	return nil
}