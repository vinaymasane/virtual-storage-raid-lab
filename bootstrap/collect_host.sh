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

info "Collecting artifacts..."

OUT="${REPO_ROOT}/artifacts"

mkdir -p \
"${OUT}/logs" \
"${OUT}/reports" \
"${OUT}/console"

cp -f \
"${LOGFILE}" \
"${OUT}/logs/" \
2>/dev/null || true

cp -rf \
"${REPO_ROOT}/packer/output-image" \
"${OUT}/" \
2>/dev/null || true

mdadm --detail /dev/md0 \
>"${OUT}/reports/mdadm.txt" \
2>/dev/null || true

cat /proc/mdstat \
>"${OUT}/reports/mdstat.txt" \
2>/dev/null || true

lsblk -f \
>"${OUT}/reports/lsblk.txt"

ip addr \
>"${OUT}/reports/ipaddr.txt"

journalctl -u ssh \
>"${OUT}/reports/ssh.log" \
2>/dev/null || true

info "Artifacts Collected."
