#!/usr/bin/env bash
# META: repo_name=Linglong CI Release (Fedora 41)
# META: repo_url=https://ci.deepin.com/repo/obs/linglong:/CI:/release/Fedora_41/linglong%3ACI%3Arelease.repo
# META: command=dnf update -y --refresh
# META: command=dnf install -y linglong-bin linyaps-web-store-installer
# META: next_steps=Add Linglong DNF repo and install linglong-bin and linyaps-web-store-installer.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if check_linglong_installed; then
    info "Linglong runtime already installed"
    exit 0
fi

add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/Fedora_41/linglong%3ACI%3Arelease.repo"
dnf install -y linglong-bin linyaps-web-store-installer

check_linglong_installed
