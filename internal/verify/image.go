package verify

import (
	"fmt"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func VerifyImage(cfg *config.Config) error {

	if !common.Exists(cfg.Image.Output) {
		return fmt.Errorf("missing image %s", cfg.Image.Output)
	}

	return nil
}
