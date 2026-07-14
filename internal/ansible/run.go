package ansible

import (
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func Run(cfg *config.Config) error {

	return common.Run(
		"ansible-playbook",
		"-i",
		cfg.Ansible.Inventory,
		cfg.Ansible.Playbook,
	)
}