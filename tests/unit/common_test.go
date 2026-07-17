package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestRepositoryLayout(t *testing.T) {

	common.AssertExists(t, "Makefile")
	common.AssertExists(t, "README.md")
	common.AssertExists(t, "go.mod")
}
