package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Load(file string) (*Config, error) {

	cfg := &Config{}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	applyDefaults(cfg)

	return cfg, Validate(cfg)
}

func applyDefaults(cfg *Config) {

	if cfg.Image.Format == "" {
		cfg.Image.Format = "qcow2"
	}

	if cfg.Artifact.OutputDir == "" {
		cfg.Artifact.OutputDir = "artifacts"
	}

	if cfg.SSH.Port == 0 {
		cfg.SSH.Port = 2222
	}

	if cfg.Ansible.Inventory == "" {
		cfg.Ansible.Inventory = "ansible/inventory.ini"
	}

	if cfg.Ansible.Playbook == "" {
		cfg.Ansible.Playbook = "ansible/site.yml"
	}
}
