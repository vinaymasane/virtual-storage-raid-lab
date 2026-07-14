package integration

import (
	"net"
	"testing"
	"time"
)

func TestSSHPort(t *testing.T) {

	c, err := net.DialTimeout(
		"tcp",
		"127.0.0.1:2222",
		5*time.Second,
	)

	if err != nil {
		t.Fatal(err)
	}

	c.Close()
}
