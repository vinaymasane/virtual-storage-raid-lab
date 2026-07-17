package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/ansible"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/artifact"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/image"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/storage"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/verify"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/vm"
)

var configPath = flag.String("config", config.DefaultConfigPath, "path to raidlab configuration file")

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

	switch flag.Arg(0) {

	case "version":
		printVersion()
		return nil

	case "bootstrap":
		return a.bootstrap()

	case "build", "image":
		return a.image()

	case "mirror":
		return a.mirror()

	case "raid":
		return a.raid()

	case "launch":
		return a.launch()

	case "stop":
		return vm.StopVM()

	case "reboot":
		return vm.RebootVM()

	case "status":
		return vm.Status()

	case "configure":
		return a.configure()

	case "verify":
		return a.verify()

	case "collect":
		return a.collect()

	case "cleanup":
		return a.cleanup()

	case "test":
		return runTests()

	case "unit":
		return runUnitTests()

	case "integration":
		return runIntegrationTests()

	default:
		return fmt.Errorf("unknown command: %s", flag.Arg(0))
	}
}

func usage() {

	fmt.Printf(`
%s %s

Virtual Storage RAID Lab

Usage:

    raidlab <command>

Commands

    bootstrap

    image

    mirror

    raid

    launch

    stop

    reboot

    status

    configure

    verify

    collect

    test

    unit

    integration

    version


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

func (a *Application) bootstrap() (*config.Config, error) {

	cfg, err := config.Load(*configPath)
	if err != nil {
		return nil, err
	}

	if err := common.PreFlight(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (a *Application) image() error {

	_, err := a.bootstrap()
	if err != nil {
		return err
	}

	return image.BuildWithPacker()
}
func (a *Application) mirror() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	return storage.CreateMirror(cfg)
}

func (a *Application) raid() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	return storage.CreateRaid(cfg)
}

func (a *Application) launch() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	return vm.LaunchVM(cfg)
}

func (a *Application) configure() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	if err := ansible.Verify(); err != nil {
		return err
	}

	if err := ansible.GenerateInventory(cfg); err != nil {
		return err
	}

	return ansible.Run(cfg)
}

func (a *Application) verify() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	return verify.Verify(cfg)
}

func (a *Application) collect() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	c := artifact.New(cfg)

	return c.Collect()
}

func (a *Application) cleanup() error {

	cfg, err := a.bootstrap()
	if err != nil {
		return err
	}

	if err := vm.StopVM(); err != nil {
		log.Println("VM already stopped or unavailable:", err)
	}

	if err := storage.Cleanup(cfg); err != nil {
		return err
	}

	if err := common.Run("./bootstrap/cleanup_host.sh"); err != nil {
		return err
	}

	common.Info("Cleanup completed successfully")

	return nil
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

func verifyRepo() error {

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
