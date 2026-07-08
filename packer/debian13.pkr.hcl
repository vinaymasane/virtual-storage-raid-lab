source "qemu" "debian" {
  iso_url      = "debian-13-netinst.iso"
  disk_size    = "10G"
  format       = "qcow2"
  accelerator  = "kvm"
  output_directory = "output-image"
}

build {
  sources = ["source.qemu.debian"]

  provisioner "shell" {
    script = "packer/scripts/provision.sh"
  }
}