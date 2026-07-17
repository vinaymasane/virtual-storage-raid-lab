package integration

import (
	"testing"

	"github.com/vinaymasane/virtual-storage-raid-lab/tests/common"
)

func TestVerification(t *testing.T) {

	common.Run(t, "make", "verify")

	common.VerifyArtifacts(t)

	common.CheckSerialLog(t, "artifacts/serial.log")

	common.WaitSSH(t, "127.0.0.1:2222")
}
