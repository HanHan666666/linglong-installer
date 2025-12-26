#!/usr/bin/env bash
# META: repo_name=Pacman official repository (Parabola)
# META: repo_url=
# META: command=pacman -Syu --noconfirm linyaps
# META: next_steps=Install linyaps from the official repository.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

require_root

if ! linglong_needs_install; then
    exit 0
fi

pacman -Syu --noconfirm linyaps

check_linglong_installed
