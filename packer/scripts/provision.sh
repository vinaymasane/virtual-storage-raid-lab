#!/usr/bin/env bash

### Set error handling options
set -Eeuo pipefail

apt update
apt install -y openssh-server sudo mdadm qemu-guest-agent
systemctl enable ssh