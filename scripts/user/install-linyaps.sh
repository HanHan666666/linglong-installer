#!/usr/bin/env bash
set -e



if ll-cli repo show | grep -q 'cdn-linglong.odata.cc'; then
    echo "Repository 'testing' already exists. skipping addition."
else
    ll-cli repo add --alias=testing stable https://cdn-linglong.odata.cc
fi

if ll-cli list | grep -q 'com.dongpl.linglong-store.v2'; then
    echo "已安装，无需重复安装，覆盖安装请手动卸载后重新安装。"
else
    echo "安装玲珑应用商店，请稍等..."
    ll-cli install com.dongpl.linglong-store.v2 --repo testing
fi

