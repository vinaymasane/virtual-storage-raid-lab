#!/usr/bin/env bash
set -Eeuo pipefail

### Common functions for the host machine
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

COMMON_SH="${SCRIPT_DIR}/common_host.sh"

if [[ ! -f "${COMMON_SH}" ]]; then
    echo "ERROR: Cannot locate ${COMMON_SH}"
    exit 1
fi

# shellcheck source=bootstrap/common_host.sh
source "${COMMON_SH}"

[[ $EUID -eq 0 ]] || die "Run using sudo."

info "Cleaning workspace..."

rm -rf "${REPO_ROOT}/output"

rm -rf "${REPO_ROOT}/packer/output-image"

find "${REPO_ROOT}" \
-name "*.qcow2" \
-delete || true

find "${REPO_ROOT}" \
-name "*.raw" \
-delete || true

if mount | grep -q "/mnt/raid"; then
    umount /mnt/raid || true
fi

if [[ -e /dev/md0 ]]; then
    mdadm --stop /dev/md0 || true
fi

qemu-nbd --disconnect /dev/nbd0 2>/dev/null || true

info "Cleanup Complete."