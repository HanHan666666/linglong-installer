#!/usr/bin/env bash
# META: repo_name=Linglong CI Release (Fedora 43)
# META: command=dnf copr enable mozixun/OpenAtom-Linyaps -y
# META: command=sudo dnf install linglong linglong-bin linyaps-web-store-installer -y
# META: next_steps=Add Linglong DNF repo and install linglong-bin and linyaps-web-store-installer.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if ! linglong_needs_install; then
    exit 0
fi

dnf copr enable mozixun/OpenAtom-Linyaps -y
sudo dnf install linglong linglong-bin linyaps-web-store-installer -y

check_linglong_installed
