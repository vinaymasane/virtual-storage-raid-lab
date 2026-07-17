package artifact

import (
	"fmt"
	"path/filepath"

	"github.com/vinaymasane/virtual-storage-raid-lab/internal/common"
	"github.com/vinaymasane/virtual-storage-raid-lab/internal/config"
)

type Collector struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Collector {
	return &Collector{
		cfg: cfg,
	}
}

func (c *Collector) Collect() error {

	artifactDir := c.cfg.Artifact.OutputDir
	logDir := filepath.Join(artifactDir, "logs")
	reportDir := filepath.Join(artifactDir, "reports")
	consoleDir := filepath.Join(artifactDir, "console")
	vmDir := filepath.Join(artifactDir, "vm")

	steps := []struct {
		src string
		dst string
	}{
		{
			src: "/proc/mdstat",
			dst: filepath.Join(reportDir, "mdstat.txt"),
		},
		{
			src: "/proc/partitions",
			dst: filepath.Join(reportDir, "partitions.txt"),
		},
		{
			src: "/var/log/syslog",
			dst: filepath.Join(logDir, "syslog.log"),
		},
		{
			src: "/var/log/dmesg",
			dst: filepath.Join(logDir, "dmesg.log"),
		},
	}

	for _, s := range steps {
		if err := common.CopyFile(s.src, s.dst); err != nil {
			common.Warn(fmt.Sprintf("unable to collect %s: %v", s.src, err))
		}
	}

	// Optional runtime captures. Ignore failures; these are best-effort.
	_ = common.RunToFile(
		filepath.Join(consoleDir, "lsblk.txt"),
		"lsblk",
		"-f",
	)

	_ = common.RunToFile(
		filepath.Join(vmDir, "virsh-list.txt"),
		"virsh",
		"list",
		"--all",
	)

	return nil
}