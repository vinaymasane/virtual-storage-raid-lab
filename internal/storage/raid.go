package storage

import "os/exec"

func CreateRAID() error {

	return exec.Command(
		"mdadm",
		"--create",
		"/dev/md0",
		"--level=1",
		"--raid-devices=2",
		"/dev/nbd0p1",
		"/dev/nbd1p1",
		"--metadata=1.0",
	).Run()
}