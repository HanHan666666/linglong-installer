#!/usr/bin/env bash
set -e

SCRIPT_ARGS=("$@")

info() { echo "[INFO] $*"; }
warn() { echo "[WARN] $*"; }
error() { echo "[ERROR] $*" >&2; }

require_root() {
    if [ "$(id -u)" -eq 0 ]; then
        return 0
    fi

    if [ "${LLI_ELEVATED:-}" = "1" ]; then
        error "Root privileges are required but elevation failed"
        exit 1
    fi

    local script_path="${BASH_SOURCE[1]:-$0}"
    if [ -z "${script_path}" ]; then
        error "Unable to resolve script path for elevation"
        exit 1
    fi
    script_path="$(cd "$(dirname "${script_path}")" && pwd)/$(basename "${script_path}")"

    if command -v pkexec >/dev/null 2>&1; then
        exec pkexec env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
    fi

    if command -v sudo >/dev/null 2>&1; then
        exec sudo -E env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
    fi

    error "Root privileges are required, but pkexec/sudo is not available"
    exit 1
}

check_linglong_installed() {
    if command -v ll-cli >/dev/null 2>&1; then
        local version
        version=$(ll-cli --version 2>/dev/null | head -n1 || echo "unknown")
        info "Linglong runtime detected: ${version}"
        return 0
    fi

    warn "Linglong runtime not installed"
    return 1
}

add_apt_repo() {
    local repo_path="$1"
    local repo_file="/etc/apt/sources.list.d/linglong.list"
    local repo_url="https://ci.deepin.com/repo/obs/linglong:/CI:/release/${repo_path}/"

    if [ -f "${repo_file}" ]; then
        info "Linglong APT repo already exists, skipping"
    else
        info "Adding APT repo: ${repo_url}"
        echo "deb [trusted=yes] ${repo_url} ./" > "${repo_file}"
    fi

    info "Refreshing APT metadata..."
    apt update
}

add_dnf_repo() {
    local repo_url="$1"

    if ls /etc/yum.repos.d/linglong*.repo >/dev/null 2>&1; then
        info "Linglong DNF repo already exists, skipping"
    else
        info "Adding DNF repo: ${repo_url}"
        dnf config-manager addrepo --from-repofile "${repo_url}" || \
        dnf config-manager --add-repo "${repo_url}"
    fi

    info "Refreshing DNF metadata..."
    dnf update -y --refresh
}

set_dnf_gpgcheck_off() {
    if ls /etc/yum.repos.d/linglong*.repo >/dev/null 2>&1; then
        info "Disabling gpgcheck for Linglong repo"
        sh -c "echo gpgcheck=0 >> /etc/yum.repos.d/linglong*.repo"
    fi
}
