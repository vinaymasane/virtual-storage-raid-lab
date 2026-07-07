package test

import (
	"net"
	"testing"
	"time"
)

func TestSSHReachable(t *testing.T) {

	timeout := time.After(2 * time.Minute)

	for {
		select {

		case <-timeout:
			t.Fatal("SSH not reachable within timeout")

		default:
			conn, err := net.DialTimeout(
				"tcp",
				"127.0.0.1:2222",
				2*time.Second,
			)

			if err == nil {
				conn.Close()
				return
			}

			time.Sleep(2 * time.Second)
		}
	}
}