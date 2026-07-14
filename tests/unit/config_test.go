package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestConfigurationFiles(t *testing.T) {

	common.AssertExists(t, "configs/defaults.yaml")
	common.AssertExists(t, "configs/config.yaml")
}