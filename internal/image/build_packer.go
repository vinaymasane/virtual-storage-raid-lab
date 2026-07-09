package image

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const packerDir = "packer"

// BuildWithPacker builds the image using packer
func BuildWithPacker() error {

	if err := os.RemoveAll(filepath.Join(packerDir, "output-image")); err != nil {
		return err
	}

	steps := [][]string{
		{"fmt", "."},
		{"init", "."},
		{"validate", "."},
		{"build", "."},
	}

	for _, s := range steps {

		log.Printf("Running: packer %s", s[0])

		if err := runPacker(s...); err != nil {
			return err
		}
	}

	img := filepath.Join(
		packerDir,
		"output-image",
		"disk_proto.qcow2",
	)

	if _, err := os.Stat(img); err != nil {
		return fmt.Errorf("%s not generated", img)
	}

	log.Println("QCOW2 image created.")

	return nil
}

func runPacker(args ...string) error {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Hour,
	)

	defer cancel()

	cmd := exec.CommandContext(ctx, "packer", args...)
	cmd.Dir = packerDir

	cmd.Env = append(os.Environ(),
		"PACKER_LOG=1",
	)

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	log.Print(out.String())

	return err
}
