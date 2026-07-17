package config

type Config struct {
	VM       VMConfig       `yaml:"vm"`
	Image    ImageConfig    `yaml:"image"`
	Storage  StorageConfig  `yaml:"storage"`
	SSH      SSHConfig      `yaml:"ssh"`
	Ansible  AnsibleConfig  `yaml:"ansible"`
	Artifact ArtifactConfig `yaml:"artifact"`
}

type VMConfig struct {
	Name   string `yaml:"name"`
	Memory int    `yaml:"memory"`
	CPUs   int    `yaml:"cpus"`
}

type ImageConfig struct {
	Output string `yaml:"output"`
	Format string `yaml:"format"`
}

type StorageConfig struct {
	PrimaryDisk   string `yaml:"primary_disk"`
	SecondaryDisk string `yaml:"secondary_disk"`
	RAIDDevice    string `yaml:"raid_device"`
}

type SSHConfig struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Key  string `yaml:"private_key"`
}

type AnsibleConfig struct {
	Inventory string `yaml:"inventory"`
	Playbook  string `yaml:"playbook"`
}

type ArtifactConfig struct {
	OutputDir string `yaml:"output_dir"`
}
