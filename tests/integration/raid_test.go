package integration

import (
	"os/exec"
	"testing"
)

func TestMDADM(t *testing.T) {

	if err := exec.Command(
		"bash",
		"-c",
		"grep md0 /proc/mdstat",
	).Run(); err != nil {

		t.Fatal(err)

	}
}
