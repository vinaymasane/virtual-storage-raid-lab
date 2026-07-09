package verify

import (
	"bytes"
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func CheckSerial() error {

	cfg := common.DefaultConfig()

	b, err := os.ReadFile(
		cfg.ArtifactDir + "/console/ttyS0.log",
	)

	if err != nil {
		return err
	}

	if !bytes.Contains(b, []byte("login")) {

		return os.ErrInvalid

	}

	return nil
}