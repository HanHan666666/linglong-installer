#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
. "${SCRIPT_DIR}/../common.sh"

if run_ll_cli repo show | grep -q 'cdn-linglong.odata.cc'; then
    echo "Repository 'testing' already exists. skipping addition."
else
    run_ll_cli repo add --alias=testing stable https://cdn-linglong.odata.cc
fi

if run_ll_cli list | grep -q 'com.dongpl.linglong-store.v2'; then
    echo "已安装，开始检查systemd服务配置..."
else
    echo "安装玲珑应用商店，请稍等..."
    run_ll_cli install com.dongpl.linglong-store.v2 --repo testing
fi
