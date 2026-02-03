#!/usr/bin/env bash
# META: repo_name=Linglong OBS Mirror (Ubuntu 25.04)
# META: repo_url=https://obs-ci.odata.cc/obs-mirror/xUbuntu_25.04/
# META: command=apt update
# META: command=apt install -y linglong-bin linglong-box
# META: next_steps=Add Linglong OBS APT repo and install linglong-bin and linglong-box.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if ! linglong_needs_install; then
    exit 0
fi

add_obs_apt_repo "xUbuntu_25.04" "obs-ubuntu2504.gpg" "obs-ubuntu2504.list"
apt install -y linglong-bin linglong-box

check_linglong_installed
