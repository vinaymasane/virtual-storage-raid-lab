package image

import "os/exec"

func CreateMirror() error {

	return exec.Command(
		"qemu-img",
		"create",
		"-f",
		"qcow2",
		"disk_mirror.qcow2",
		"10G",
	).Run()
}
