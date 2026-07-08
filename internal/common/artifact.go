package common

import (
	"io"
	"os"
	"path/filepath"
)

func Copy(src, dst string) error {

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	os.MkdirAll(filepath.Dir(dst), 0755)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)

	return err
}

func Ensure(cfg *Config) {

	os.MkdirAll(cfg.OutputDir, 0755)

	os.MkdirAll(cfg.ArtifactDir, 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "logs"), 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "reports"), 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "console"), 0755)

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "vm"), 0755)
}
