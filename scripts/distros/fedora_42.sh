#!/usr/bin/env bash
# META: repo_name=Linglong CI Release (Fedora 42)
# META: repo_url=https://ci.deepin.com/repo/obs/linglong:/CI:/release/Fedora_42/linglong%3ACI%3Arelease.repo
# META: command=dnf update -y --refresh
# META: command=dnf install -y linglong-bin linyaps-web-store-installer
# META: next_steps=Add Linglong DNF repo and install linglong-bin and linyaps-web-store-installer.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if ! linglong_needs_install; then
    exit 0
fi

add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/Fedora_42/linglong%3ACI%3Arelease.repo"
dnf install -y linglong-bin linyaps-web-store-installer policykit-1

check_linglong_installed
