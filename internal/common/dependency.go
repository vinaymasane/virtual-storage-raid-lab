package common

import (
	"fmt"
	"os/exec"
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
			return fmt.Errorf("Dependency Missing: %s", t)
		}

	}

	return nil
}

func CheckDependency(name string) error {

	_, err := exec.LookPath(name)

	return err
}