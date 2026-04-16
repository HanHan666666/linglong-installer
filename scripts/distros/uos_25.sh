#!/usr/bin/env bash
# META: repo_name=UOS 25 system repository
# META: repo_url=(system default APT repository; no third-party source)
# META: command=apt update
# META: command=apt install -y linglong-bin linglong-box
# META: next_steps=Install linglong-bin and linglong-box from the system repository, then add the Linglong testing repo and install the store.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

deb_package_installed() {
    dpkg-query -W -f='${Status}' "$1" 2>/dev/null | grep -q "install ok installed"
}

require_root

needs_runtime_install=0
if linglong_needs_install; then
    needs_runtime_install=1
fi

# UOS 25 ships the runtime in its system repo, so avoid adding any Linglong CI
# source here and still keep linglong-box present for the later store setup.
needs_box_install=0
if ! deb_package_installed "linglong-box"; then
    needs_box_install=1
fi

if [ "${needs_runtime_install}" -eq 0 ] && [ "${needs_box_install}" -eq 0 ]; then
    info "UOS 25 runtime and linglong-box are already installed, skipping"
    exit 0
fi

apt update
apt install -y linglong-bin linglong-box

# Verify the extra runtime companion package explicitly because the generic
# version check only tells us whether ll-cli is already usable.
if ! deb_package_installed "linglong-box"; then
    error "linglong-box installation verification failed"
    exit 1
fi

check_linglong_installed
