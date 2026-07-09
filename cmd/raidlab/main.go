package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/image"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/storage"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
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

	cfg := common.DefaultConfig()

	common.Ensure(cfg)

	common.InitLogger(cfg)

	if err := common.CheckDependencies(); err != nil {
		log.Fatal(err)
	}

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

		if err == nil {
			err = vm.WaitForVM()
		}

	case "stop":
		err = vm.StopVM()

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
