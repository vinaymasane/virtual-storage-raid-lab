package integration

import (
	"os"
	"testing"
)

func TestImageCreated(t *testing.T) {

	if _, err := os.Stat("output/disk_proto.qcow2"); err != nil {
		t.Fatal(err)
	}
}
