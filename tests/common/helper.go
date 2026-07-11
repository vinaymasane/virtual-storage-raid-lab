package common

import (
	"os"
	"testing"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func AssertExists(t *testing.T, path string) {
	t.Helper()
	if !Exists(path) {
		t.Fatalf("%s not found", path)
	}
}