package common

import (
	"io"
	"os"
	"path/filepath"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
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
        cfg.Artifact.OutputDir,
        filepath.Join(cfg.Artifact.OutputDir, "logs"),
        filepath.Join(cfg.Artifact.OutputDir, "reports"),
        filepath.Join(cfg.Artifact.OutputDir, "console"),
        filepath.Join(cfg.Artifact.OutputDir, "vm"),
        cfg.OutputDir,
        cfg.Coverage.OutputDir,
    }

    for _, d := range dirs {

        if err := os.MkdirAll(d, 0755); err != nil {
			Error("Failed to create directory " + d + ": " + err.Error())
            return err
        }
    }

    return nil
}
