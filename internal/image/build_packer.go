package image

import (
	"log"
	"os/exec"
)

func BuildWithPacker() error {

	log.Println("Starting Packer build...")

	cmd := exec.Command(
		"packer",
		"build",
		"-var", "disk_size=10G",
		"packer/debian13.pkr.hcl",
	)

	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	return cmd.Run()
}
