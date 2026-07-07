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

APP=raidlab
LOGFILE=/var/log/${APP}-bootstrap.log

exec > >(tee -a "$LOGFILE")
exec 2>&1

### Function to install or upgrade APT packages
install_or_upgrade_apt_pkg() {

    local pkg="$1"

    if dpkg -s "${pkg}" >/dev/null 2>&1; then

        info "${pkg} already installed"

        if apt list --upgradable 2>/dev/null \
           | grep -q "^${pkg}/"; then

            info "Upgrading ${pkg}"

            apt install --only-upgrade -y "${pkg}"

        else
            info "${pkg} already current"
        fi

    else

        info "Installing ${pkg}"

        apt install -y "${pkg}"
    fi
}

### Function to install or upgrade Go
install_or_upgrade_go() {

    local GO_VERSION="1.26.4"

    if binary_exists go; then

        CURRENT=$(go version | awk '{print $3}')

        info "Go already installed (${CURRENT})"

        if [[ "${CURRENT}" != "go${GO_VERSION}" ]]; then

            info "Upgrading Go to ${GO_VERSION}"

        else

            info "Go already current (${CURRENT})"
            return 0

        fi

    else

        info "Installing Go ${GO_VERSION}"

    fi

    # Download and install Go
    wget -q \
        "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
        -O /tmp/go.tar.gz

    rm -rf /usr/local/go

    tar -C /usr/local \
        -xzf /tmp/go.tar.gz

    info "Go ${GO_VERSION} installed Successfully"
}

### Function to install or upgrade Packer
install_or_upgrade_packer() {

    local PACKER_VERSION="1.15.4"

    if binary_exists packer; then

        CURRENT=$(packer version | awk '{print $2}')

        info "Packer already installed (${CURRENT})"

        if [[ "${CURRENT}" != "v${PACKER_VERSION}" ]]; then

            info "Upgrading Packer to ${PACKER_VERSION}"

        else

            info "Packer already current (${CURRENT})"
            return 0

        fi

    else

        info "Installing Packer ${PACKER_VERSION}"
    fi

    # Download and install Packer
    wget -q \
        "https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_amd64.zip" \
        -O /tmp/packer.zip

    unzip -o \
        /tmp/packer.zip \
        -d /usr/local/bin

    info "Packer ${PACKER_VERSION} installed Successfully"
}

### Pre-installation steps for the host machine
info "Updating package repository"

apt update -y
apt upgrade -y

info "Done."

info "Installing required packages"
BASE_PKGS=(
    curl
    wget
    git
    vim
    jq
    unzip
    openssh-server
    sudo
)

VIRT_PKGS=(
    qemu-utils
    qemu-system-x86
    qemu-kvm
    libvirt-daemon-system
    libvirt-clients
    virtinst
    virt-manager
    dnsmasq-base
    libguestfs-tools
    bridge-utils
)

STORAGE_PKGS=(
    mdadm
    util-linux
    parted
    fdisk
    e2fsprogs
    kpartx
)

AUTOMATION_PKGS=(
    ansible
    python3
    python3-pip
    python3-venv
)

GO_DEV_PKGS=(
    gcc
    g++
    make
)

TEST_PKGS=(
    fio
    sysstat
    smartmontools
    lsof
    psmisc
    netcat-openbsd
)

SUPPORT_PKGS=(
    rsync
    sshpass
    ncdu
)

### Install all required packages
for pkg in \
    "${BASE_PKGS[@]}" \
    "${VIRT_PKGS[@]}" \
    "${STORAGE_PKGS[@]}" \
    "${CONFIG_PKGS[@]}" \
    "${SUPPORT_PKGS[@]}" \
    "${GO_DEV_PKGS[@]}" \
    "${TEST_PKGS[@]}"
do
    install_or_upgrade_apt_pkg "${pkg}"
done

### Install or upgrade Go and Packer
install_or_upgrade_go

### Install or upgrade Packerd
install_or_upgrade_packer

### Ensure that the SSH service is running
ensure_service_running ssh
