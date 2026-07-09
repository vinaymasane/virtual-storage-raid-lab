package verify

import (
	"net"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

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