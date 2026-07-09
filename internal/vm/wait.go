package vm

import (
	"fmt"
	"time"
)

// WaitForVM waits for the virtual machine to become available by checking the serial console and SSH service.
func WaitForVM() error {

	if err := WaitForSerialConsole(); err != nil {
		return err
	}

	if err := WaitForSSH(); err != nil {
		return err
	}

	return nil
}

/* Retry repeatedly executes the provided function
 * until it returns true or the specified timeout is reached.
 * It sleeps for 2 seconds between attempts.
 */
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
