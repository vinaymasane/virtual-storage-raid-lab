#!/bin/bash
set -eux

apt-get update

apt-get install -y \
openssh-server \
cloud-init \
sudo \
curl \
vim

systemctl enable ssh

grep -q ttyS0 /etc/default/grub || \
sed -i \
's/GRUB_CMDLINE_LINUX="/GRUB_CMDLINE_LINUX="console=ttyS0 /' \
/etc/default/grub

update-grub

systemctl enable ssh

systemctl start ssh

passwd -d root

mkdir -p /root/.ssh