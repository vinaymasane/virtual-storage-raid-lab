package vm

import "os/exec"

func StartVM() error {

	return exec.Command(
		"qemu-system-x86_64",
		"-enable-kvm",
		"-m", "2048",
		"-drive", "file=disk_proto.qcow2,if=virtio",
		"-netdev", "user,id=n1,hostfwd=tcp::2222-:22",
		"-device", "virtio-net-pci,netdev=n1",
	).Start()
}
