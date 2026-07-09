#!/usr/bin/env bash
set -e

VERSION=$(git describe --tags --always)

mkdir -p release

tar czf release/raidlab-${VERSION}.tar.gz \
bin \
bootstrap \
packer \
ansible \
README.md \
Makefile