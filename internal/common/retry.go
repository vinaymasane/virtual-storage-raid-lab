package common

import "time"

func Retry(count int, delay time.Duration, fn func() error) error {

	var err error

	for i := 0; i < count; i++ {

		err = fn()
		if err == nil {
			return nil
		}

		time.Sleep(delay)
	}

	return err
}