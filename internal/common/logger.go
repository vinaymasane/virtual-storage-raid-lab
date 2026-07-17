package common

import (
	"log"
	"os"
	"path/filepath"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

func Info(msg string) {
	log.Printf("[INFO ] %s", msg)
}

func Warn(msg string) {
	log.Printf("[WARN ] %s", msg)
}

func Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}

func InitLogger(cfg *config.Config) error {

    logDir := filepath.Join(cfg.Artifact.OutputDir, "logs")

    if err := os.MkdirAll(logDir, 0755); err != nil {
        return err
    }

    logfile := filepath.Join(logDir, "raidlab.log")

    f, err := os.OpenFile(
        logfile,
        os.O_CREATE|os.O_APPEND|os.O_WRONLY,
        0644,
    )
    if err != nil {
        return err
    }

    log.SetOutput(f)

    return nil
}
