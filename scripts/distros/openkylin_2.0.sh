#!/usr/bin/env bash
# META: repo_name=Linglong CI Release (openKylin 2.0)
# META: repo_url=https://ci.deepin.com/repo/obs/linglong:/CI:/release/openkylin_2.0/
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

add_apt_repo "openkylin_2.0"
apt install -y linglong-bin linglong-installer policykit-1

check_linglong_installed
