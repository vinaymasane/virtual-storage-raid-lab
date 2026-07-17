#!/usr/bin/env bash

### Function to install or upgrade APT packages
install_or_upgrade_apt_pkg() {

    local pkg="$1"

    if pkg_installed "${pkg}"; then

        info "${pkg} already installed"

        if apt list --upgradable 2>/dev/null \
           | grep -q "^${pkg}/"; then

            info "Upgrading ${pkg}"

            run apt install --only-upgrade -y "${pkg}" || true

        else
            info "${pkg} already current"
        fi

    else

        info "Installing ${pkg}"

        run apt install -y "${pkg}"
    fi
}

### Function to install or upgrade Go
install_or_upgrade_go() {

    local GO_VERSION="1.26.4"

    if binary_exists go; then

        CURRENT=$(go version | awk '{print $3}')

        info "Go already installed: (${CURRENT})"

        if [[ "${CURRENT}" != "go${GO_VERSION}" ]]; then

            info "Upgrading Go to ${GO_VERSION}"

        else

            info "Go already current: (${CURRENT})"
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

    # Ensure current shell can locate it
    GO_ROOT="/usr/local/go"
    GO_BIN="${GO_ROOT}/bin"

    if ! echo ":${PATH}:" | grep -q ":${GO_BIN}:"; then
        export PATH="${GO_BIN}:${PATH}"
    fi

    hash -r

    # Clean up
    rm -f /tmp/go.tar.gz

    info "Go ${GO_VERSION} installed Successfully"
}

### Function to install or upgrade golangci-lint
install_or_upgrade_golangci_lint() {

    local VERSION="2.12.2"
    local INSTALL_DIR="/usr/local/bin"

    local ARCH
    ARCH="$(uname -m)"

    case "${ARCH}" in
        x86_64) ARCH="amd64" ;;
        aarch64) ARCH="arm64" ;;
        *) die "Unsupported architecture: ${ARCH}" ;;
    esac

    local PLATFORM="linux"
    local FILE="golangci-lint-${VERSION}-${PLATFORM}-${ARCH}.tar.gz"
    local URL="https://github.com/golangci/golangci-lint/releases/download/v${VERSION}/${FILE}"

    if binary_exists golangci-lint; then

        local CURRENT
        CURRENT=$(golangci-lint version | sed -n 's/.*version \([^ ]*\).*/\1/p')

        info "golangci-lint already installed (${CURRENT})"

        if [[ "${CURRENT}" == "v${VERSION}" || "${CURRENT}" == "${VERSION}" ]]; then
            info "golangci-lint already current (v${VERSION})"
            return 0
        fi

        info "Upgrading golangci-lint ${CURRENT} -> v${VERSION}"
    fi

    info "Downloading golangci-lint v${VERSION}..."

    local TMP
    TMP=$(mktemp -d)

    curl -fsSL "${URL}" -o "${TMP}/${FILE}" \
        || die "Failed downloading ${URL}"

    tar -xzf "${TMP}/${FILE}" -C "${TMP}" \
        || die "Failed extracting golangci-lint"

    install -m 0755 \
        "${TMP}/golangci-lint-${VERSION}-${PLATFORM}-${ARCH}/golangci-lint" \
        "${INSTALL_DIR}/golangci-lint" \
        || die "Installation failed"

    rm -rf "${TMP}"

    #
    # Ensure current shell can locate it
    #
    if ! echo ":${PATH}:" | grep -q ":${INSTALL_DIR}:"; then
        export PATH="${INSTALL_DIR}:${PATH}"
    fi

    hash -r

    binary_exists golangci-lint || die "golangci-lint not found after installation"

    info "golangci-lint location : $(command -v golangci-lint)"
    info "golangci-lint version  : $(golangci-lint version)"
}

### Function to install Go dependencies
install_go_dependencies() {

    info "Downloading Go modules..."

    go mod tidy

    info "Done."
}

### Function to install or upgrade Packer
install_or_upgrade_packer() {

    local PACKER_VERSION="1.15.4"

    if binary_exists packer; then

        CURRENT=$(packer version | awk '{print $2}')

        info "Packer already installed: (${CURRENT})"

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

    # Ensure current shell can locate it
    if ! echo ":${PATH}:" | grep -q ":/usr/local/bin:"; then
        export PATH="/usr/local/bin:${PATH}"
    fi
    hash -r

    # Clean up
    rm -f /tmp/packer.zip

    info "Packer ${PACKER_VERSION} installed Successfully"
}


### Function to generate SSH keys for testing
generate_ssh_keys() {

    local KEY_DIR="tests/testdata/ssh_keys"
    local KEY_FILE="${KEY_DIR}/id_ed25519"

    mkdir -p "${KEY_DIR}"

    if [[ -f "${KEY_FILE}" && -f "${KEY_FILE}.pub" ]]; then
        info "SSH key already exists"
        return 0
    fi

    info "Generating SSH key..."

    ssh-keygen \
        -q \
        -t ed25519 \
        -N "" \
        -f "${KEY_FILE}"

    info "SSH key created"
}

### Common functions for the host machine
export SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export COMMON_SH="${SCRIPT_DIR}/common_host.sh"

if [[ ! -f "${COMMON_SH}" ]]; then
    echo "ERROR: Cannot locate ${COMMON_SH}"
    exit 1
fi

# shellcheck source=bootstrap/common_host.sh
echo "Sourcing ${COMMON_SH}"
source "${COMMON_SH}"
info "Done."

[[ $EUID -eq 0 ]] || die "Run using sudo."


### Pre-installation steps for the host machine
info "Updating package repository"

apt update -y

info "Done."

info "Installing required packages"
BASE_PKGS=(
    curl
    wget
    git
    vim
    jq
    unzip
    tar
    gzip
    bzip2
    openssh-server
    openssh-client
    ca-certificates
    apt-transport-https
    sudo
)

VIRT_PKGS=(
    qemu-utils
    qemu-system-x86
    qemu-kvm
    qemu-system-gui
    qemu-block-extra
    libvirt-daemon-system
    libvirt-clients
    virtinst
    virt-manager
    dnsmasq-base
    libguestfs-tools
    bridge-utils
    ovmf
    seabios
    cloud-image-utils
)

STORAGE_PKGS=(
    mdadm
    util-linux
    parted
    gdisk
    e2fsprogs
    kpartx
    dosfstools
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
    "${AUTOMATION_PKGS[@]}" \
    "${SUPPORT_PKGS[@]}" \
    "${GO_DEV_PKGS[@]}" \
    "${TEST_PKGS[@]}"
do
    install_or_upgrade_apt_pkg "${pkg}"
done

### Install or upgrade Go and Packer
install_or_upgrade_go

### Install or upgrade golangci-lint
install_or_upgrade_golangci_lint

### Install Go dependencies
install_go_dependencies

### Install or upgrade Packer
install_or_upgrade_packer

### Generate SSH keys for testing
generate_ssh_keys

### Ensure that the SSH service is running
ensure_service_running ssh

export GOROOT="/usr/local/go"
export GOPATH="${HOME}/go"
export GOBIN="${GOPATH}/bin"

export PATH="/usr/local/go/bin:/usr/local/bin:${GOBIN}:${PATH}"
export PATH="${PATH}:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

hash -r
