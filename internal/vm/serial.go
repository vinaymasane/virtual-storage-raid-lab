package vm

import (
	"bytes"
	"os"
	"time"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

// WaitForSerialConsole waits for the serial console to be ready by checking the log file for specific login prompts.
func WaitForSerialConsole() error {

	cfg := common.DefaultConfig()

	logFile := cfg.ArtifactDir + "/console/ttyS0.log"

	return Retry(10*time.Minute, func() bool {

		b, err := os.ReadFile(logFile)

		if err != nil {
			return false
		}

		return bytes.Contains(b, []byte("login:")) ||
			bytes.Contains(b, []byte("Debian GNU/Linux"))

	})
}
