package verify

import (
	"net"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

// CheckSSH checks if the SSH service is available on the virtual machine by attempting to establish a TCP connection.
func CheckSSH() error {

	cfg := common.DefaultConfig()

	c, err := net.Dial(
		"tcp",
		cfg.SSHHost+":"+cfg.SSHPort,
	)

	if err != nil {

		return err

	}

	c.Close()

	return nil
}
