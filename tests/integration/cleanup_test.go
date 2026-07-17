package integration

import (
	"testing"

	"github.com/vinaymasane/virtual-storage-raid-lab/tests/common"
)

func TestCleanup(t *testing.T) {

	common.Run(t, "make", "clean")
}
