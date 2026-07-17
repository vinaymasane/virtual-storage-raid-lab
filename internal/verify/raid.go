package verify

import (
	"fmt"
	"strings"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func VerifyRAID(cfg *config.Config) error {

	out, err := common.Output("cat", "/proc/mdstat")
	if err != nil {
		return err
	}

	if !strings.Contains(string(out), cfg.Storage.RAIDDevice) {
		return fmt.Errorf("%s not active", cfg.Storage.RAIDDevice)
	}

	return nil
}
