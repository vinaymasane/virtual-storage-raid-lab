package common

import (
	"io"
	"os"
	"path/filepath"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
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
		common.Error("Failed to open source file: " + err.Error())
		return err
	}
	defer in.Close()

	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		common.Error("Failed to create destination directory: " + err.Error())
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		common.Error("Failed to create destination file: " + err.Error())
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		common.Error("Failed to copy file: " + err.Error())
		return err
	}

	err = out.Sync()
	if err != nil {
		common.Error("Failed to sync destination file: " + err.Error())
		return err
	}

	return nil
}

// createRequiredDirectories creates the necessary directories for the artifact and output.
func createRequiredDirectories(cfg *Config) {

	os.MkdirAll(cfg.OutputDir, 0755)

	os.MkdirAll(cfg.ArtifactDir, 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "logs"), 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "reports"), 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "console"), 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "vm"), 0755)
}
