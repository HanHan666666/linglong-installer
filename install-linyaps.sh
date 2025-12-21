#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DISTRO_DIR="${SCRIPT_DIR}/scripts/distros"

. "${SCRIPT_DIR}/scripts/common.sh"

detect_distro() {
    if [ ! -f /etc/os-release ]; then
        error "Missing /etc/os-release"
        exit 1
    fi

    . /etc/os-release
    DISTRO_ID="${ID}"
    DISTRO_VERSION="${VERSION_ID}"
    DISTRO_NAME="${PRETTY_NAME:-${ID} ${VERSION_ID}}"

    info "Detected distro: ${DISTRO_NAME}"
}

resolve_script() {
    local script_name="${DISTRO_ID}_${DISTRO_VERSION}.sh"
    local script_path="${DISTRO_DIR}/${script_name}"

    if [ -f "${script_path}" ]; then
        echo "${script_path}"
        return 0
    fi

    local script_lower="${DISTRO_DIR}/$(echo "${DISTRO_ID}" | tr 'A-Z' 'a-z')_${DISTRO_VERSION}.sh"
    if [ -f "${script_lower}" ]; then
        echo "${script_lower}"
        return 0
    fi

    return 1
}

main() {
    echo "========================================"
    echo "  Linglong Installer Script"
    echo "========================================"
    echo ""

    detect_distro

    if check_linglong_installed; then
        info "Linglong runtime already installed"
        exit 0
    fi

    script_path="$(resolve_script || true)"
    if [ -z "${script_path}" ]; then
        error "Unsupported distro: ${DISTRO_NAME}"
        echo ""
        echo "Supported scripts:"
        if [ -d "${DISTRO_DIR}" ]; then
            ls -1 "${DISTRO_DIR}" | sed 's/\\.sh$//' | sort
        else
            echo "  (no scripts directory found)"
        fi
        exit 1
    fi

    require_root
    bash "${script_path}"

    echo ""
    info "Done"
}

main "$@"
