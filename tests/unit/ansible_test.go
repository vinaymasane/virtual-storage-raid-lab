package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestAnsibleFiles(t *testing.T) {

	common.AssertExists(t, "ansible/site.yml")
	common.AssertExists(t, "ansible/inventory.ini")
}