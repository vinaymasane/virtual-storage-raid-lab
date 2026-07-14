package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/ansible"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/artifact"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/image"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/storage"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/vm"
)

const (
	appName    = "raidlab"
	appVersion = "1.0.0"
)

type Application struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger *log.Logger
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := &Application{
		ctx:    ctx,
		cancel: cancel,
		logger: log.Default(),
	}

	app.registerSignals()

	if err := app.run(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func (a *Application) registerSignals() {

	c := make(chan os.Signal, 2)

	signal.Notify(
		c,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {

		<-c

		a.logger.Println("Received termination signal...")

		a.cancel()

		os.Exit(1)
	}()
}

func (a *Application) run() error {

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
		return nil
	}

	cmd := flag.Arg(0)

	switch cmd {

	case "version":
		printVersion()

	case "bootstrap":
		_, err := config.Load("configs/config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		return common.PreFlight()

	case "build":
		return image.BuildWithPacker()

	case "image":
		return image.BuildWithPacker()

	case "mirror":
		return storage.CreateMirror()

	case "raid":
		return storage.CreateRaid()

	case "launch":
		return vm.LaunchVM()

	case "stop":
		return vm.StopVM()

	case "reboot":
		return vm.RebootVM()

	case "status":
		return vm.Status()

	case "configure":
		return ansible.Run()

	case "collect":
	    c := artifact.New()
		return c.Collect()

	case "test":
		return runTests()

	case "unit":
		return runUnitTests()

	case "integration":
		return runIntegrationTests()

	case "verify":
		return verify()

	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}

	return nil
}

func usage() {

	fmt.Printf(`
%s %s

Usage:

    raidlab <command>

Commands

    version

    bootstrap

    build

    image

    mirror

    raid

    launch

    stop

    reboot

    status

    configure

    test

    unit

    integration

    verify

    collect


Examples

    raidlab bootstrap

    raidlab build

    raidlab mirror

    raidlab raid

    raidlab launch

    raidlab configure

    raidlab test

`, appName, appVersion)
}

func printVersion() {

	fmt.Println("-----------------------------------")
	fmt.Printf("Application : %s\n", appName)
	fmt.Printf("Version     : %s\n", appVersion)
	fmt.Printf("Go Version  : %s\n", runtime.Version())
	fmt.Printf("OS          : %s\n", runtime.GOOS)
	fmt.Printf("Arch        : %s\n", runtime.GOARCH)
	fmt.Println("-----------------------------------")
}

func runTests() error {

	start := time.Now()

	log.Println("Running complete test suite...")

	if err := runUnitTests(); err != nil {
		return err
	}

	if err := runIntegrationTests(); err != nil {
		return err
	}

	log.Printf("All tests completed in %v\n", time.Since(start))

	return nil
}

func runUnitTests() error {

	log.Println("Running unit tests...")

	return common.Run(
		"go",
		"test",
		"./tests/unit/...",
		"-race",
		"-coverprofile=coverage/unit.out",
	)
}

func runIntegrationTests() error {

	log.Println("Running integration tests...")

	return common.Run(
		"go",
		"test",
		"./tests/integration/...",
		"-v",
		"-timeout=45m",
		"-coverprofile=coverage/integration.out",
	)
}

func verify() error {

	common.Info("Running repository verification...")

	collector := artifact.New()

	steps := []func() error{

		common.PreFlight,

		vm.WaitForVM,

		vm.WaitForSerialConsole,

		vm.WaitForSSH,

		collector.Collect,
	}

	for _, s := range steps {

		if err := s(); err != nil {
			return err
		}
	}

	return nil
}

func init() {

	required := []string{

		"artifacts",

		"coverage",

		"output",
	}

	for _, dir := range required {

		if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {

			log.Fatalf("cannot create %s : %v", dir, err)
		}
	}
}
