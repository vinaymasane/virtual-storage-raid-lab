package common

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
)

func Run(name string, args ...string) error {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		60*time.Minute,
	)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)

	cmd.Env = os.Environ()

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	log.Printf("Running: %s %v", name, args)

	err := cmd.Run()

	log.Print(out.String())

	if err != nil {
		return fmt.Errorf("%s failed : %w", name, err)
	}

	return nil
}

func RunDir(dir string, name string, args ...string) error {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		60*time.Minute,
	)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	cmd.Env = os.Environ()

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	common.Info(fmt.Sprintf("Running: %s %v", name, args))

	err := cmd.Run()

	common.Info(out.String())

	if err != nil {
		return common.Error(fmt.Sprintf("%s failed : %v", name, err))
	}

	return nil
}

func StartBackground(cmd string, args ...string) error {

	c := exec.Command(cmd, args...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Start()
}

func Output(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).CombinedOutput()
}