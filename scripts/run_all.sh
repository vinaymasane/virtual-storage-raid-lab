#!/usr/bin/env bash
set -Eeuo pipefail

make clean

make bootstrap

make build

make image

make mirror

make raid

make launch

make configure

make verify

make collect

go test ./... -v