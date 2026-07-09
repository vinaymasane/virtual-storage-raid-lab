package vm

import (
	"fmt"
	"time"
)

func WaitForVM() error {

	if err := WaitForSerialConsole(); err != nil {
		return err
	}

	if err := WaitForSSH(); err != nil {
		return err
	}

	return nil
}

func Retry(timeout time.Duration, fn func() bool) error {

	deadline := time.Now().Add(timeout)

	for {

		if fn() {
			return nil
		}

		if time.Now().After(deadline) {
			break
		}

		time.Sleep(2 * time.Second)

	}

	return fmt.Errorf("timeout")
}
