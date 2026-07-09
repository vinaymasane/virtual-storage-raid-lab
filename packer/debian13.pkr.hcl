source "qemu" "debian" {

  vm_name = var.vm_name

  output_directory="output-image"

  disk_image=false

  format="qcow2"

  disk_size=var.disk_size

  accelerator=var.accelerator

  headless=true

  cpus=var.cpus

  memory=var.memory

  iso_url=var.iso_url

  iso_checksum=var.iso_checksum

  boot_wait="5s"

  http_directory="http"

  communicator="ssh"

  ssh_username=var.ssh_username

  ssh_password=var.ssh_password

  ssh_timeout="45m"

  shutdown_command="shutdown -P now"

  boot_command=[
    "<esc><wait>",
    "auto ",
    "console=ttyS0 ",
    "url=http://{{ .HTTPIP }}:{{ .HTTPPort }}/preseed.cfg",
    "<enter>"
  ]
}

build {

  sources=[
    "source.qemu.debian"
  ]

  provisioner "shell" {

    script="scripts/provision.sh"

  }

  provisioner "shell" {

    script="scripts/cleanup.sh"

  }

}