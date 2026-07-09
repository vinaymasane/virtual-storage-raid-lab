package integration

import (
	"os/exec"
	"testing"
)

func TestSSHLogin(t *testing.T) {

	err := exec.Command(
		"ssh",
		"-o",
		"StrictHostKeyChecking=no",
		"-p",
		"2222",
		"root@127.0.0.1",
		"hostname",
	).Run()

	if err != nil {
		t.Fatal(err)
	}
}
