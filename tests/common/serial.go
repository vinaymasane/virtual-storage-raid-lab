package common

import (
	"os"
	"testing"
)

func CheckSerialLog(t *testing.T, logfile string) {

	t.Helper()

	info, err := os.Stat(logfile)

	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {
		t.Fatal("serial log is empty")
	}
}
