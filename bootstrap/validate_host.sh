#!/usr/bin/env bash

### Common functions for the host machine
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

COMMON_SH="${SCRIPT_DIR}/common_host.sh"

if [[ ! -f "${COMMON_SH}" ]]; then
    echo "ERROR: Cannot locate ${COMMON_SH}"
    exit 1
fi

# shellcheck source=bootstrap/common_host.sh
source "${COMMON_SH}"

### Validate the host machine for RAID Lab
REQUIRED_TOOLS=(
    git
    curl
    go
    packer
    ansible
    mdadm
    qemu-img
    qemu-nbd
)

for tool in "${REQUIRED_TOOLS[@]}"
do
    if ! binary_exists "${tool}"; then
        fail "Missing required tool: ${tool}"
    fi
done

info "Required tools are installed on the host machine"

info "Installed versions:"

git --version
curl --version | head -1
go version
packer version
ansible --version | head -1
mdadm --version
qemu-img --version | head -1

info "Bootstrap completed Successfully"