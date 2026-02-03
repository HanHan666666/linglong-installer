#!/usr/bin/env bash
# META: repo_name=Linglong OBS Repo (openSUSE 15.6)
# META: repo_url=https://download.opensuse.org/repositories/home:/guanzi:/Fedora/15.6/
# META: command=zypper addrepo https://download.opensuse.org/repositories/home:/guanzi:/Fedora/15.6/home:guanzi:Fedora.repo
# META: command=zypper refresh
# META: command=zypper install -y linglong-bin linglong-box linglong-builder
# META: next_steps=Add repo and install linglong-bin, linglong-box, and linglong-builder.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if ! linglong_needs_install; then
    exit 0
fi

zypper addrepo https://download.opensuse.org/repositories/home:/guanzi:/Fedora/15.6/home:guanzi:Fedora.repo
zypper refresh
zypper install -y linglong-bin linglong-box linglong-builder

check_linglong_installed
