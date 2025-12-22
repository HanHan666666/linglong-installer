#!/usr/bin/env bash
set -euo pipefail

BUS_NAME="org.linglong_store.LinyapsManager"
BIN_DIR="$HOME/.linglong-store-v2"
BIN="$BIN_DIR/linyaps-dbus-server"
SERV_DIR="$HOME/.config/systemd/user"

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

systemctl --user daemon-reload
