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

		if err := CheckDependency(t); err != nil {
			Error("Dependency check failed for required tool: " + t + " " + err.Error())
			return err
		}
	}

	return nil
}

func CheckDependency(name string) error {

	if _, err := exec.LookPath(name); err != nil {
		Error("Missing required tool : " + name + " " + err.Error())
		return err
	}

	return nil
}
