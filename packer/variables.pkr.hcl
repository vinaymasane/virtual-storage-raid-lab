variable "vm_name" {
  default="disk_proto"
}

variable "disk_size" {
  default=10240
}

variable "memory" {
  default=4096
}

variable "cpus" {
  default=2
}

variable "accelerator" {
  default="kvm"
}

variable "iso_url" {
  default="/home/adminboss/Downloads/debian-13.5.0-amd64-netinst.iso"
}

variable "iso_checksum" {
  default="sha256:95838884f5ea6c82421dfe6baaa5a639dbbe6756c1e380f9fe7a7cb0c1949d2a"
}

variable "ssh_username" {
  default="root"
}

variable "ssh_password" {
  default="root"
}