package integration

import (
	"testing"

	"github.com/vinamasane/virtual-storage-raid-lab/tests/common"
)

func TestAssignmentWorkflow(t *testing.T) {

	steps := [][]string{
		{"make", "build"},
		{"make", "image"},
		{"make", "mirror"},
		{"make", "raid"},
		{"make", "launch"},
		{"make", "configure"},
	}

	for _, step := range steps {
		common.Run(t, step[0], step[1:]...)
	}
}