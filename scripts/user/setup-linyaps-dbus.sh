#!/usr/bin/env bash
set -euo pipefail

BUS_NAME="org.linglong_store.LinyapsManager"
BIN_DIR="$HOME/.linglong-store-v2"
BIN="$BIN_DIR/linyaps-dbus-server"
SERV_DIR="$HOME/.config/systemd/user"

info() { echo "[INFO] $*"; }
warn() { echo "[WARN] $*"; }

ensure_runtime_dir() {
  if [ -n "${XDG_RUNTIME_DIR:-}" ] && [ -d "${XDG_RUNTIME_DIR}" ]; then
    return 0
  fi

  local uid
  uid="$(id -u)"
  local candidate="/run/user/${uid}"
  if [ -d "${candidate}" ]; then
    export XDG_RUNTIME_DIR="${candidate}"
    return 0
  fi

  return 1
}

user_bus_available() {
  ensure_runtime_dir || return 1
  local bus="${XDG_RUNTIME_DIR}/bus"
  [ -S "${bus}" ]
}

mkdir -p "$BIN_DIR"
if [ ! -f "$BIN" ]; then
  ll-cli run com.dongpl.linglong-store.v2 cp /opt/apps/com.dongpl.linglong-store.v2/files/bin/linyaps-dbus-server "$BIN" || true
fi

if [ ! -x "$BIN" ]; then
  echo "Missing executable: $BIN"
  exit 1
fi

mkdir -p "$SERV_DIR"

cat > "$SERV_DIR/${BUS_NAME}.service" <<EOF_SERVICE
[Unit]
Description=Linyaps Manager Service (User)
After=dbus.socket

[Service]
Type=dbus
BusName=${BUS_NAME}
Environment=DBUS_SESSION_BUS_ADDRESS=unix:path=%t/bus
ExecStart=${BIN}

[Install]
WantedBy=default.target
EOF_SERVICE

if ! command -v systemctl >/dev/null 2>&1; then
  warn "systemctl not found; skip enabling user service"
  exit 0
fi

if user_bus_available; then
  systemctl --user daemon-reload
  systemctl --user enable --now "${BUS_NAME}.service"
else
  warn "User systemd bus unavailable; skip enabling service"
  info "You can run later: systemctl --user enable --now ${BUS_NAME}.service"
fi
systemctl --user start org.linglong_store.LinyapsManager.service
