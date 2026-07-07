package test

import (
	"os"
	"testing"
)

func TestMissingImageHandled(t *testing.T) {

	_, err := os.Stat("nonexistent.qcow2")

	if err == nil {
		t.Fatal("expected failure but file exists")
	}
}
