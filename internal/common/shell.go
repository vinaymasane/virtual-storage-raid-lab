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

	Info("Running: " + name + " " + fmt.Sprintf("%v", args))

	err := cmd.Run()

	Info(out.String())

	if err != nil {
		Error("Failed: " + name + " " + fmt.Sprintf("%v", args) + " : " + err.Error())
		return err
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

	Info("Working Directory: " + dir)
	Info("Running: " + name + " " + fmt.Sprintf("%v", args))

	err := cmd.Run()

	Info(out.String())

	if err != nil {
		Error("Failed: " + name + " " + fmt.Sprintf("%v", args) + " :  + err.Error())")
		return err
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
