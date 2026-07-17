package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestPackerFiles(t *testing.T) {

	common.AssertExists(t, "packer/debian13.pkr.hcl")
	common.AssertExists(t, "packer/variables.pkr.hcl")
}
