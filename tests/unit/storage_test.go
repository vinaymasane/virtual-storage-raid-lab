package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestStorageScripts(t *testing.T) {

	common.AssertExists(t, "bootstrap/preinstall_host.sh")
	common.AssertExists(t, "bootstrap/common_host.sh")
}
