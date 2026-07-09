package common

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
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

	err := cmd.Run()

	log.Print(out.String())

	if err != nil {
		return fmt.Errorf("%s failed : %w", name, err)
	}

	return nil
}

func StartBackground(cmd string, args ...string) error {

	c := exec.Command(cmd, args...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Start()
}
