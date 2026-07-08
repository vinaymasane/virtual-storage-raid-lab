package common

import (
	"fmt"
	"os/exec"
)

var RequiredTools = []string{
	"go",
	"packer",
	"ansible-playbook",
	"git",
	"qemu-img",
	"qemu-system-x86_64",
	"qemu-nbd",
	"mdadm",
	"ssh",
}

func CheckDependencies() error {

	for _, t := range RequiredTools {

		if _, err := exec.LookPath(t); err != nil {
			return fmt.Errorf("dependency missing : %s", t)
		}

	}

	return nil
}
