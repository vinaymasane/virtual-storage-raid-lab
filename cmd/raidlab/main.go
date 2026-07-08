package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/ansible"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/image"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/storage"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/verify"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/vm"
)

func usage() {
	fmt.Println("raidlab")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  build")
	fmt.Println("  mirror")
	fmt.Println("  raid")
	fmt.Println("  launch")
	fmt.Println("  configure")
	fmt.Println("  verify")
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	var err error

	switch os.Args[1] {

	case "build":
		err = image.BuildWithPacker()

	case "mirror":
		err = storage.CreateMirror()

	case "raid":
		err = storage.CreateRaid()

	case "launch":
		err = vm.LaunchVM()

	case "configure":
		err = ansible.Run()

	case "verify":
		err = verify.Run()

	default:
		usage()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
}