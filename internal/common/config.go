package common

import (
	"os"
	"path/filepath"
)

type Config struct {
	RepoRoot      string
	OutputDir     string
	ArtifactDir   string
	PackerDir     string
	AnsibleDir    string
	ImageName     string
	ImageFile     string
	SecondaryDisk string
	RaidDevice    string
	SSHUser       string
	SSHHost       string
	SSHPort       string
}

func DefaultConfig() *Config {

	root, _ := os.Getwd()

	return &Config{
		RepoRoot:      root,
		OutputDir:     filepath.Join(root, "output"),
		ArtifactDir:   filepath.Join(root, "artifacts"),
		PackerDir:     filepath.Join(root, "packer"),
		AnsibleDir:    filepath.Join(root, "ansible"),
		ImageName:     "disk_proto",
		ImageFile:     filepath.Join(root, "output", "disk_proto.qcow2"),
		SecondaryDisk: filepath.Join(root, "output", "secondary.raw"),
		RaidDevice:    "/dev/md0",
		SSHUser:       "root",
		SSHHost:       "127.0.0.1",
		SSHPort:       "2222",
	}
}
