package common

import "testing"

func VerifyArtifacts(t *testing.T) {

	t.Helper()

	files := []string{
		"artifacts/serial.log",
		"artifacts/mdstat.txt",
		"artifacts/lsblk.txt",
	}

	for _, f := range files {
		AssertExists(t, f)
	}
}