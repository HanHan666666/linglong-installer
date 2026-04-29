#!/usr/bin/env bash
set -euo pipefail

# Smoke-test install-linyaps-env.sh in multiple distro containers.
# PASS criterion: the script reaches the runtime install command, not just a
# prerequisite package install for repository bootstrap.

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TARGET_SCRIPT="${TARGET_SCRIPT:-${SCRIPT_DIR}/install-linyaps-env.sh}"
TIMEOUT_SEC="${TIMEOUT_SEC:-240}"
RESULT_ROOT="${RESULT_ROOT:-${SCRIPT_DIR}/smoke-results}"
RUN_ID="$(date +%Y%m%d-%H%M%S)"
OUT_DIR="${RESULT_ROOT}/${RUN_ID}"
LOG_DIR="${OUT_DIR}/logs"
SUMMARY_CSV="${OUT_DIR}/summary.csv"
SUMMARY_TXT="${OUT_DIR}/summary.txt"

INSTALL_TRIGGER_REGEX='^\+ (apt install -y .*linglong-bin|dnf install -y .*linglong-bin|zypper install -y .*linglong-bin|pacman -Syu --noconfirm linyaps)'

CASE_MATRIX="$(cat <<'EOF'
debian12|debian:12|apt-get update; apt-get install -y ca-certificates
debian-sid|debian:sid|apt-get update; apt-get install -y ca-certificates curl gnupg
ubuntu2404|ubuntu:24.04|apt-get update; apt-get install -y ca-certificates
ubuntu2604|ubuntu:26.04|apt-get update; apt-get install -y ca-certificates
fedora42|fedora:42|dnf -y install ca-certificates dnf-plugins-core
opensuse156|opensuse/leap:15.6|zypper --non-interactive refresh
archlatest|archlinux:latest|
EOF
)"

usage() {
  cat <<'EOF'
Usage:
  ./docker-smoke.sh [--list] [case1 case2 ...]

Examples:
  ./docker-smoke.sh
  ./docker-smoke.sh --list
  ./docker-smoke.sh debian12 fedora42

Environment variables:
  TARGET_SCRIPT   Script under test (default: ./install-linyaps-env.sh)
  TIMEOUT_SEC     Per-case timeout seconds (default: 240)
  RESULT_ROOT     Output root directory (default: ./smoke-results)
EOF
}

csv_escape() {
  local s="${1:-}"
  s="${s//\"/\"\"}"
  printf '"%s"' "${s}"
}

is_selected_case() {
  local name="$1"
  if [ "${#SELECTED_CASES[@]}" -eq 0 ]; then
    return 0
  fi
  local c
  for c in "${SELECTED_CASES[@]}"; do
    if [ "${c}" = "${name}" ]; then
      return 0
    fi
  done
  return 1
}

run_case() {
  local case_name="$1"
  local image="$2"
  local pre_cmd="$3"

  local log_file="${LOG_DIR}/${case_name}.log"
  local run_rc=0
  local status="FAIL"
  local trigger_line=""
  local short_trigger="N/A"

  echo "[$(date +%H:%M:%S)] CASE=${case_name} IMAGE=${image}"

  set +e
  {
    echo "[INFO] Pulling image ${image}"
    docker pull "${image}"
    echo "[INFO] Running smoke test in ${image}"
    if command -v timeout >/dev/null 2>&1; then
      timeout "${TIMEOUT_SEC}"s docker run --rm \
        -e PRE_CMD="${pre_cmd}" \
        -e LLI_SKIP_VERSION_CHECK=1 \
        -e DEBIAN_FRONTEND=noninteractive \
        -v "${TARGET_SCRIPT}:/work/install-linyaps-env.sh:ro" \
        "${image}" \
        bash -lc '
          set +e
          if [ -n "${PRE_CMD:-}" ]; then
            echo "[PREPARE] ${PRE_CMD}"
            bash -lc "${PRE_CMD}"
            prep_rc=$?
            echo "[PREPARE] exit=${prep_rc} (ignored)"
          fi
          set -e
          bash -x /work/install-linyaps-env.sh
        '
    else
      docker run --rm \
        -e PRE_CMD="${pre_cmd}" \
        -e LLI_SKIP_VERSION_CHECK=1 \
        -e DEBIAN_FRONTEND=noninteractive \
        -v "${TARGET_SCRIPT}:/work/install-linyaps-env.sh:ro" \
        "${image}" \
        bash -lc '
          set +e
          if [ -n "${PRE_CMD:-}" ]; then
            echo "[PREPARE] ${PRE_CMD}"
            bash -lc "${PRE_CMD}"
            prep_rc=$?
            echo "[PREPARE] exit=${prep_rc} (ignored)"
          fi
          set -e
          bash -x /work/install-linyaps-env.sh
        '
    fi
  } >"${log_file}" 2>&1
  run_rc=$?
  set -e

  trigger_line="$(grep -E -m1 "${INSTALL_TRIGGER_REGEX}" "${log_file}" || true)"
  if [ -n "${trigger_line}" ]; then
    status="PASS"
    short_trigger="${trigger_line#'+ '}"
  fi

  printf '%s,%s,%s,%s,%s,%s\n' \
    "$(csv_escape "${case_name}")" \
    "$(csv_escape "${image}")" \
    "$(csv_escape "${status}")" \
    "$(csv_escape "${run_rc}")" \
    "$(csv_escape "${short_trigger}")" \
    "$(csv_escape "${log_file}")" \
    >> "${SUMMARY_CSV}"

  printf '%-12s %-24s %-4s %-4s %s\n' \
    "${case_name}" "${image}" "${status}" "${run_rc}" "${short_trigger}" \
    >> "${SUMMARY_TXT}"

  if [ "${status}" = "PASS" ]; then
    PASS_COUNT=$((PASS_COUNT + 1))
  else
    FAIL_COUNT=$((FAIL_COUNT + 1))
  fi
}

if [ "${1:-}" = "--help" ] || [ "${1:-}" = "-h" ]; then
  usage
  exit 0
fi

if [ "${1:-}" = "--list" ]; then
  while IFS='|' read -r name image _; do
    [ -z "${name}" ] && continue
    printf '%-12s %s\n' "${name}" "${image}"
  done <<< "${CASE_MATRIX}"
  exit 0
fi

TARGET_SCRIPT="$(readlink -f "${TARGET_SCRIPT}")"
if [ ! -f "${TARGET_SCRIPT}" ]; then
  echo "[ERROR] TARGET_SCRIPT not found: ${TARGET_SCRIPT}" >&2
  exit 1
fi

if ! command -v docker >/dev/null 2>&1; then
  echo "[ERROR] docker not found in PATH" >&2
  exit 1
fi

if ! docker info >/dev/null 2>&1; then
  echo "[ERROR] docker daemon is not available for current user" >&2
  exit 1
fi

declare -a SELECTED_CASES=()
if [ "$#" -gt 0 ]; then
  SELECTED_CASES=("$@")
fi

mkdir -p "${LOG_DIR}"
printf 'case,image,status,exit_code,trigger,log\n' > "${SUMMARY_CSV}"
printf '%-12s %-24s %-4s %-4s %s\n' "CASE" "IMAGE" "STAT" "RC" "TRIGGER" > "${SUMMARY_TXT}"

PASS_COUNT=0
FAIL_COUNT=0
TOTAL_COUNT=0

while IFS='|' read -r case_name image pre_cmd; do
  [ -z "${case_name}" ] && continue
  if ! is_selected_case "${case_name}"; then
    continue
  fi
  TOTAL_COUNT=$((TOTAL_COUNT + 1))
  run_case "${case_name}" "${image}" "${pre_cmd}"
done <<< "${CASE_MATRIX}"

if [ "${TOTAL_COUNT}" -eq 0 ]; then
  echo "[ERROR] No matching cases selected."
  echo "Use --list to view available cases."
  exit 1
fi

{
  echo
  echo "Results:"
  cat "${SUMMARY_TXT}"
  echo
  echo "Summary: total=${TOTAL_COUNT} pass=${PASS_COUNT} fail=${FAIL_COUNT}"
  echo "Artifacts:"
  echo "  ${SUMMARY_CSV}"
  echo "  ${SUMMARY_TXT}"
  echo "  ${LOG_DIR}/"
} >&2

if [ "${FAIL_COUNT}" -gt 0 ]; then
  exit 1
fi

