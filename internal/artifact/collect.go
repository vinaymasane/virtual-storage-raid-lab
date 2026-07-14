package artifact

import (
	"fmt"
	"os"
	"path/filepath"

	"virtual-storage-raid-lab/internal/common"
)

const (
	ArtifactDir = "artifacts"
	LogDir      = "artifacts/logs"
)

type Collector struct {
	OutputDir string
}

func New() *Collector {
	return &Collector{
		OutputDir: ArtifactDir,
	}
}

func (c *Collector) Prepare() error {

	dirs := []string{
		ArtifactDir,
		LogDir,
	}

	for _, d := range dirs {

		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collector) Collect() error {

	if err := c.Prepare(); err != nil {
		return err
	}

	steps := []struct {
		src string
		dst string
	}{
		{
			"/proc/mdstat",
			filepath.Join(ArtifactDir, "mdstat.txt"),
		},
		{
			"/var/log/syslog",
			filepath.Join(LogDir, "syslog.log"),
		},
	}

	for _, s := range steps {

		if err := common.CopyFile(s.src, s.dst); err != nil {
			common.Warn(
				fmt.Sprintf("unable to collect %s", s.src),
			)
		}
	}

	return nil
}