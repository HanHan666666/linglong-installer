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

if ! linglong_needs_install; then
    exit 0
fi

add_apt_repo "xUbuntu_24.04"
apt install -y linglong-bin linglong-installer policykit-1

check_linglong_installed
