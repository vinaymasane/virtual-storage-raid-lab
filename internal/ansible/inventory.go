package ansible

import (
	"os"
	"path/filepath"
)

func GenerateInventory() error {

	// Allow override via ANSIBLE_DIR env var, fallback to ./ansible
	ansibleDir := os.Getenv("ANSIBLE_DIR")
	if ansibleDir == "" {
		ansibleDir = "./ansible"
	}

	if err := os.MkdirAll(ansibleDir, 0o755); err != nil {
		return err
	}

	data := `[vm]
127.0.0.1 ansible_port=2222 ansible_user=root ansible_password=root ansible_connection=ssh ansible_ssh_common_args='-o StrictHostKeyChecking=no'
`

	return os.WriteFile(filepath.Join(ansibleDir, "inventory.ini"), []byte(data), 0o644)
}
