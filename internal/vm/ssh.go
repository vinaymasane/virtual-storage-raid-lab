package vm

import (
	"net"
	"time"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func WaitForSSH() error {

	cfg := common.DefaultConfig()

	return Retry(5*time.Minute, func() bool {

		conn, err := net.DialTimeout(
			"tcp",
			cfg.SSHHost+":"+cfg.SSHPort,
			2*time.Second,
		)

		if err != nil {
			return false
		}

		conn.Close()

		return true

	})
}
