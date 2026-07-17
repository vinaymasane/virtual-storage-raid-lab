package common

import (
	"io"
	"os"
	"path/filepath"

	"github.com/ekagra-ranjan/virtual-storage-raid-lab/internal/common/config"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		Error("Failed to open source file: " + err.Error())
		return err
	}
	defer in.Close()

	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		Error("Failed to create destination directory: " + err.Error())
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		Error("Failed to create destination file: " + err.Error())
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		Error("Failed to copy file: " + err.Error())
		return err
	}

	err = out.Sync()
	if err != nil {
		Error("Failed to sync destination file: " + err.Error())
		return err
	}

	return nil
}


func EnsureDirectories(cfg *config.Config) error {

	dirs := []string{
		cfg.image.output_directory,
		cfg.artifacts.output_dir,
		filepath.Join(cfg.artifacts.output_dir, "logs"),
		filepath.Join(cfg.artifacts.output_dir, "reports"),
		filepath.Join(cfg.artifacts.output_dir, "console"),
		filepath.Join(cfg.artifacts.output_dir, "vm"),
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	return nil
}