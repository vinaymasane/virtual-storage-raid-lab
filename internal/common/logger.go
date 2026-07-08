package common

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func InitLogger(cfg *Config) {

	os.MkdirAll(filepath.Join(cfg.ArtifactDir, "logs"), 0755)

	f, err := os.OpenFile(
		filepath.Join(cfg.ArtifactDir, "logs", "raidlab.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, f))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
