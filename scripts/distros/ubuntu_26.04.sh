#!/usr/bin/env bash
# META: repo_name=Linglong OBS Mirror (Ubuntu 26.04)
# META: repo_url=https://obs-ci.odata.cc/obs-mirror/xUbuntu_26.04/
# META: command=apt install -y curl gpg
# META: command=apt update
# META: command=apt install -y linglong-bin linglong-box
# META: next_steps=Install curl and gpg for OBS key import, add Linglong OBS APT repo, and install linglong-bin and linglong-box.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if ! linglong_needs_install; then
	exit 0
fi

# Ubuntu 26.04 currently ships the Linglong packages through the OBS mirror,
# so we bootstrap curl plus gpg first to match the tested repository setup flow.
apt install -y curl gpg
add_obs_apt_repo "xUbuntu_26.04" "obs-ubuntu2604.gpg" "obs-ubuntu2604.list"
apt install -y linglong-bin linglong-box

check_linglong_installed
