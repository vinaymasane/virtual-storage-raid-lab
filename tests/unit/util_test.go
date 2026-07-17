package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestScriptsExist(t *testing.T) {

	common.AssertExists(t, "scripts/test.sh")
	common.AssertExists(t, "scripts/lint.sh")
	common.AssertExists(t, "scripts/coverage.sh")
}
