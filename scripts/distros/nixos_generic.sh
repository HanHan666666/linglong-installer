#!/usr/bin/env bash
# META: repo_name=NixOS module (services.linyaps)
# META: repo_url=https://github.com/NixOS/nixpkgs
# META: command=edit /etc/nixos/configuration.nix and enable services.linyaps
# META: command=nixos-rebuild switch
# META: next_steps=Enable the NixOS linyaps service and rebuild the system configuration.
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

cat <<'MSG'
[INFO] Please enable the linyaps service in /etc/nixos/configuration.nix and rebuild your system.

Example configuration:
  services.linyaps.enable = true;

Then run:
  sudo nixos-rebuild switch

  ll-cli repo add --alias=testing stable https://cdn-linglong.odata.cc
  ll-cli install com.dongpl.linglong-store.v2 --repo testing
MSG

exit 1
