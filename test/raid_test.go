package test

import (
	"os"
	"strings"
	"testing"
)

func TestRAIDHealthy(t *testing.T) {

	data, err := os.ReadFile("/proc/mdstat")
	if err != nil {
		t.Fatalf("failed to read mdstat: %v", err)
	}

	if !strings.Contains(string(data), "[UU]") {
		t.Fatalf("RAID not healthy: %s", string(data))
	}
}