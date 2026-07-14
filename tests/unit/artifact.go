package unit

import (
	"testing"

	"virtual-storage-raid-lab/tests/common"
)

func TestArtifactDirectory(t *testing.T) {

	common.AssertExists(t, "artifacts")
}