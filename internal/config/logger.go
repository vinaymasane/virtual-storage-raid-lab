package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(file string) (*Config, error) {

	cfg := DefaultConfig()

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, Validate(cfg)
}

func DefaultConfig() *Config {

	return &Config{

		Image: ImageConfig{
			Format: "qcow2",
		},

		Artifact: ArtifactConfig{
			OutputDir: "artifacts",
		},

		SSH: SSHConfig{
			Port: 2222,
		},

		Ansible: AnsibleConfig{
			Inventory: "ansible/inventory.ini",
			Playbook:  "ansible/site.yml",
		},
	}
}
