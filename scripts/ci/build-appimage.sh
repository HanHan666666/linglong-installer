#!/bin/sh
set -eu

if [ $# -ne 1 ]; then
  echo "usage: $0 <release_tag>" >&2
  exit 2
fi

RELEASE_TAG="$1"

need_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "missing required command: $1" >&2
    exit 1
  fi
}

need_cmd curl
need_cmd chmod
need_cmd base64
need_cmd install
need_cmd ln
need_cmd mkdir
need_cmd rm

APPIMAGETOOL_APPIMAGE="appimagetool-x86_64.AppImage"
APPIMAGETOOL_URL="https://github.com/AppImage/appimagetool/releases/download/continuous/appimagetool-x86_64.AppImage"

curl -fsSL -o "$APPIMAGETOOL_APPIMAGE" "$APPIMAGETOOL_URL"
chmod +x "$APPIMAGETOOL_APPIMAGE"
./"$APPIMAGETOOL_APPIMAGE" --appimage-extract
APPIMAGETOOL="$PWD/squashfs-root/AppRun"

make_appdir() {
  appdir="$1"
  src_bin="$2"
  arch="$3"

  rm -rf "$appdir"
  mkdir -p "$appdir/usr/bin"
  install -m 0755 "$src_bin" "$appdir/usr/bin/linglong-store-installer"

  printf '%s\n' \
    '#!/bin/sh' \
    'set -e' \
    'exec "${APPDIR}/usr/bin/linglong-store-installer" "$@"' \
    > "$appdir/AppRun"
  chmod +x "$appdir/AppRun"

  printf '%s\n' \
    '[Desktop Entry]' \
    'Type=Application' \
    'Name=Linglong Store Installer' \
    'NoDisplay=true' \
    'Exec=linglong-store-installer' \
    'Icon=linglong-store-installer' \
    'Terminal=true' \
    'Categories=Utility;' \
    'X-AppImage-Name=linglong-store-installer' \
    "X-AppImage-Version=${RELEASE_TAG}" \
    "X-AppImage-Arch=${arch}" \
    > "$appdir/linglong-store-installer.desktop"

  # 1x1 transparent PNG (keeps desktop integration happy without adding new assets)
  printf '%s' 'iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg==' | base64 -d > "$appdir/linglong-store-installer.png"
  ln -sf "linglong-store-installer.png" "$appdir/.DirIcon"
}

# amd64
make_appdir "AppDir-amd64" "linglong-store-installer-amd64" "x86_64"
ARCH=x86_64 "$APPIMAGETOOL" --no-appstream "AppDir-amd64" "linglong-store-installer-${RELEASE_TAG}-x86_64.AppImage"

# arm64
make_appdir "AppDir-arm64" "linglong-store-installer-arm64" "aarch64"
ARCH=aarch64 "$APPIMAGETOOL" --no-appstream "AppDir-arm64" "linglong-store-installer-${RELEASE_TAG}-aarch64.AppImage"
