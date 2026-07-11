package common

import (
	"net"
	"testing"
	"time"
)

func WaitSSH(t *testing.T, host string) {

	t.Helper()

	timeout := time.After(60 * time.Second)

	for {

		select {

		case <-timeout:
			t.Fatal("SSH timeout")

		default:

			c, err := net.DialTimeout("tcp", host, 2*time.Second)

			if err == nil {
				c.Close()
				return
			}

			time.Sleep(2 * time.Second)
		}
	}
}