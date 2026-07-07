#!/usr/bin/env bash

### Set error handling options
set -Eeuo pipefail

### Export environment variables for non-interactive APT operations and PATH
export DEBIAN_FRONTEND=noninteractive
export DEBCONF_NONINTERACTIVE_SEEN=true
export PATH=$PATH:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

export SCRIPT_DIR
export REPO_ROOT


### Logging
LOG_PREFIX="[RAIDLAB]"

info() {
    echo "${LOG_PREFIX} $*"
}

warn() {
    echo "${LOG_PREFIX} WARNING: $*"
}

fail() {
    echo "${LOG_PREFIX} ERROR: $*"
    exit 1
}

### Function to check if a binary exists in the system
binary_exists() {

    command -v "$1" >/dev/null 2>&1
}

### Function to ensure a service is running
ensure_service_running() {

    local service="$1"

    info "Enabling ${service}"
    systemctl enable "${service}"

    if systemctl is-active --quiet "${service}"; then
        info "${service} already running"
    else
        info "Starting service : ${service}"
        systemctl start "${service}"

        if ! systemctl is-active --quiet "${service}"; then
            fail "Failed to start service: ${service}"
            return 1
        fi

        info "Started service: ${service}"
    fi
}
