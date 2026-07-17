package ansible

import (
	"fmt"
	"os"

	"github.com/vinamasane/virtual-storage-raid-lab/internal/config"
)

func GenerateInventory(cfg *config.Config) error {

	content := fmt.Sprintf(`[raidlab]
%s ansible_user=%s ansible_port=%d ansible_ssh_private_key_file=%s`,
		cfg.SSH.Host,
		cfg.SSH.User,
		cfg.SSH.Port,
		cfg.SSH.Key,
	)

	return os.WriteFile(
		cfg.Ansible.Inventory,
		[]byte(content),
		0644,
	)
}
