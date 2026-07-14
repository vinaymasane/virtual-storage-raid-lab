package common

import (
	"fmt"
	"os/exec"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

var RequiredTools = []string{
	"git",
	"go",
	"packer",
	"ansible-playbook",
	"qemu-img",
	"qemu-system-x86_64",
	"qemu-nbd",
	"mdadm",
	"ssh",
}

func CheckDependencies() error {

	for _, t := range RequiredTools {

		if _, err := exec.LookPath(t); err != nil {
			return common.Error("Dependency Missing: %s", t)
		}

	}

	return nil
}
