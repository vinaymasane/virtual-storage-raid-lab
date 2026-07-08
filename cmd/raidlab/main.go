package main

import (
	"fmt"
	"os"

	"virtual-storage-raid-lab/internal/image"
	"virtual-storage-raid-lab/internal/storage"
	"virtual-storage-raid-lab/internal/vm"
	"virtual-storage-raid-lab/internal/config"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: raidlab <build|mirror|raid|launch|configure|validate>")
		return
	}

	switch os.Args[1] {

	case "build":
		image.BuildWithPacker()

	case "mirror":
		image.CreateMirror()

	case "raid":
		storage.CreateRAID()

	case "launch":
		vm.StartVM()

	case "configure":
		config.RunAnsible()


	default:
		fmt.Println("unknown command")
	}
}