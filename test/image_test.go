package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestQCOW2Valid(t *testing.T) {

	out, err := exec.Command(
		"qemu-img",
		"info",
		"disk_proto.qcow2",
	).CombinedOutput()

	if err != nil {
		t.Fatalf("qemu-img failed: %v", err)
	}

	if !strings.Contains(string(out), "file format: qcow2") {
		t.Fatalf("invalid qcow2 image: %s", string(out))
	}
}
