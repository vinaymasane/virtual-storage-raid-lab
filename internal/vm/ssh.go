package vm

import (
	"net"
	"time"
)

func SSHReady() bool {

	for i := 0; i < 60; i++ {

		conn, err := net.DialTimeout("tcp", "127.0.0.1:2222", 2*time.Second)

		if err == nil {
			conn.Close()
			return true
		}

		time.Sleep(2 * time.Second)
	}

	return false
}