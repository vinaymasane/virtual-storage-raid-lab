#!/usr/bin/env bash
set -e

mkdir -p artifacts/final

cp -r artifacts/* artifacts/final/ 2>/dev/null || true

journalctl -xe > artifacts/final/journal.log || true

dmesg > artifacts/final/dmesg.log || true

virsh list --all > artifacts/final/libvirt.txt || true