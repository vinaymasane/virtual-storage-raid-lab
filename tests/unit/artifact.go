package unit

import (
	"testing"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func TestArtifactDirectory(t *testing.T) {

	common.AssertExists(t, "artifacts")
}
