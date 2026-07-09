package config

import (
	"log"
	"os/exec"
)

// RunAnsible executes the Ansible playbook to configure the virtual machine over SSH.
func RunAnsible() error {

	log.Println("Running Ansible configuration over SSH...")

	cmd := exec.Command(
		"ansible-playbook",
		"-i",
		"ansible/inventory.ini",
		"ansible/playbook.yml",
	)

	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	return cmd.Run()
}
