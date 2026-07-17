package verify

import (
	"fmt"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func Verify(cfg *config.Config) error {
	checks := []struct {
		name string
		fn   func(*config.Config) error
	}{
		{"Image", VerifyImage},
		{"RAID", VerifyRAID},
		{"VM", VerifyVM},
		{"SSH", VerifySSH},
		{"Serial", VerifySerial},
		{"Artifacts", VerifyArtifacts},
	}

	for _, c := range checks {

		if err := c.fn(cfg); err != nil {
			return fmt.Errorf("%s verification failed: %w", c.name, err)
		}
	}

	common.Info("Verification completed Successfully.")
	return nil
}
