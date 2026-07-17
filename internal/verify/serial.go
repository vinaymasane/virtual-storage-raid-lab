package verify

import (
	"fmt"

	"github.com/vinamasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinamasane/virtual-storage-raid-lab/internal/config"
)

func VerifySerial(cfg *config.Config) error {

	const serialLog = "artifacts/serial.log"

	if !common.Exists(serialLog) {
		return fmt.Errorf("serial log missing")
	}

	return nil
}
