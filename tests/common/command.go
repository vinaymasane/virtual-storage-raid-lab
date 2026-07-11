package common

import (
	"os"
	"os/exec"
	"testing"
)

func Run(t *testing.T, name string, args ...string) {

	t.Helper()

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
}