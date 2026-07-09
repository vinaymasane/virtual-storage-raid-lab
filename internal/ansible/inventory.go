package ansible

import (
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func GenerateInventory() error {

	cfg := common.DefaultConfig()

	data := `[vm]
127.0.0.1 ansible_port=2222 ansible_user=root ansible_password=root ansible_connection=ssh ansible_ssh_common_args='-o StrictHostKeyChecking=no'
`

	return os.WriteFile(
		cfg.AnsibleDir+"/inventory.ini",
		[]byte(data),
		0644,
	)
}
