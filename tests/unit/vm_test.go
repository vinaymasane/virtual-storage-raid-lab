package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestVMConfiguration(t *testing.T) {

	common.AssertExists(t, "configs/config.yaml")
}
