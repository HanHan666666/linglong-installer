#!/usr/bin/env bash
# META: repo_name=Linglong CI Release (Ubuntu 24.04)
# META: repo_url=https://ci.deepin.com/repo/obs/linglong:/CI:/release/xUbuntu_24.04/
# META: command=apt update
# META: command=apt install -y linglong-bin linglong-installer
# META: next_steps=Add Linglong APT repo and install linglong-bin and linglong-installer.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if check_linglong_installed; then
    info "Linglong runtime already installed"
    exit 0
fi

add_apt_repo "xUbuntu_24.04"
apt install -y linglong-bin linglong-installer

check_linglong_installed
