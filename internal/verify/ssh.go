package verify

import (
	"fmt"

	"github.com/vinamasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinamasane/virtual-storage-raid-lab/internal/config"
)

func VerifySSH(cfg *config.Config) error {

	_, err := common.Output(
		"ssh",
		"-i", cfg.SSH.Key,
		"-p", fmt.Sprint(cfg.SSH.Port),
		cfg.SSH.User+"@"+cfg.SSH.Host,
		"echo ok",
	)

	return err
}
