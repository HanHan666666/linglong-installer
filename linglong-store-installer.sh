#!/usr/bin/env bash
set -euo pipefail

# ========= 前置检查 =========
# 检测到 银河麒麟v10 和 NixOS 时，给出特殊提示并退出
is_kylin_v10="false"
is_nixos="false"

if [[ -f /etc/os-release ]]; then
  # shellcheck disable=SC1091
  . /etc/os-release

  if [[ "${ID:-}" == "nixos" ]] || [[ "${NAME:-}" == *"NixOS"* ]] || [[ "${PRETTY_NAME:-}" == *"NixOS"* ]]; then
    is_nixos="true"
  fi

  if [[ "${ID:-}" == "kylin" ]] || [[ "${NAME:-}" == *"Kylin"* ]] || [[ "${NAME:-}" == *"银河麒麟"* ]] || \
     [[ "${PRETTY_NAME:-}" == *"Kylin"* ]] || [[ "${PRETTY_NAME:-}" == *"银河麒麟"* ]]; then
    if [[ "${VERSION_ID:-}" == "V10" ]] || [[ "${VERSION_ID:-}" == "10" ]] || [[ "${VERSION:-}" == *"V10"* ]] || \
       [[ "${PRETTY_NAME:-}" == *"V10"* ]]; then
      is_kylin_v10="true"
    fi
  fi
fi

if [[ -f /etc/NIXOS ]]; then
  is_nixos="true"
fi

if [[ -f /etc/kylin-version ]] && grep -qi "V10" /etc/kylin-version; then
  is_kylin_v10="true"
fi
if [[ -f /etc/kylin-release ]] && grep -qi "V10" /etc/kylin-release; then
  is_kylin_v10="true"
fi

if [[ "$is_kylin_v10" == "true" ]]; then
  echo -e "\033[1;33m[提示]\033[0m 银河麒麟 V10 系统因为有应用保护，请手动安装玲珑环境。"
  echo "安装方法见：https://bbs.deepin.org.cn/phone/zh/post/289061"
  echo "如果已安装，或者已经安装完成，可手动安装商店："
  echo "  ll-cli repo add --alias=testing stable https://cdn-linglong.odata.cc"
  echo "  ll-cli install com.dongpl.linglong-store.v2 --repo testing"
  exit 0
fi

if [[ "$is_nixos" == "true" ]]; then
  echo "[INFO] NixOS detected, manual installation is required."
  echo "[INFO] Please enable the linyaps service in /etc/nixos/configuration.nix and rebuild your system."
  echo ""
  echo "Example configuration:"
  echo "  services.linyaps.enable = true;"
  echo ""
  echo "Then run:"
  echo "  sudo nixos-rebuild switch"
  echo ""
  echo "  ll-cli repo add --alias=testing stable https://cdn-linglong.odata.cc"
  echo "  ll-cli install com.dongpl.linglong-store.v2 --repo testing"
  exit 0
fi
# ========= 前置检查结束 =========


# ========= 基本配置 =========
REPO_OWNER="hanplus"
REPO_NAME="linglong-installer"
VERSION="latest"

# 二进制基础名（不带架构）
BIN_NAME="linglong-store-installer"

# 下载地址前缀
BASE_URL="https://gitee.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${VERSION}"

err() {
  echo -e "\033[1;31m[ERROR]\033[0m $*" >&2
  exit 1
}

# ========= 架构检测 =========
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64 | amd64)
    ARCH_SUFFIX="amd64"
    ;;
  aarch64 | arm64)
    ARCH_SUFFIX="arm64"
    ;;
  *)
    err "不支持的架构: $ARCH"
    ;;
esac

BIN_FILE="${BIN_NAME}-${ARCH_SUFFIX}"
DOWNLOAD_URL="${BASE_URL}/${BIN_FILE}"

# ========= 下载目录 =========
WORKDIR="$(mktemp -d)"
BIN_PATH="${WORKDIR}/${BIN_FILE}"

cleanup() {
  rm -rf "$WORKDIR"
}
trap cleanup EXIT

curl -fL "$DOWNLOAD_URL" -o "$BIN_PATH"


chmod +x "$BIN_PATH"


"$BIN_PATH"
