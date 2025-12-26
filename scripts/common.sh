#!/usr/bin/env bash
set -e

SCRIPT_ARGS=("$@")
MIN_LINGLONG_VERSION="${LLI_MIN_LINGLONG_VERSION:-1.9.0}"

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

    local has_session=""
    if [ -n "${DISPLAY:-}" ] || [ -n "${WAYLAND_DISPLAY:-}" ] || [ -n "${DBUS_SESSION_BUS_ADDRESS:-}" ] || [ -n "${XDG_SESSION_TYPE:-}" ] || [ -n "${XDG_SESSION_ID:-}" ]; then
        has_session="1"
    fi

    if [ "${LLI_PREFER_PKEXEC:-}" = "1" ] && command -v pkexec >/dev/null 2>&1; then
        if [ -n "${has_session}" ]; then
            exec pkexec env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
        fi
        warn "pkexec is preferred but no desktop session detected; falling back to sudo"
    fi

    if command -v sudo >/dev/null 2>&1; then
        exec sudo -E env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
    fi

    if command -v pkexec >/dev/null 2>&1; then
        if [ -n "${has_session}" ]; then
            exec pkexec env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
        fi
        warn "pkexec is available but no desktop session detected"
    fi

    error "Root privileges are required, but pkexec/sudo is not available"
    exit 1
}

get_linglong_version() {
    local raw=""
    if command -v ll-cli >/dev/null 2>&1; then
        raw=$(ll-cli --version 2>/dev/null | head -n1 || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi
    if command -v dpkg-query >/dev/null 2>&1; then
        raw=$(dpkg-query -W -f='${Version}' linglong-bin 2>/dev/null || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi
    if command -v rpm >/dev/null 2>&1; then
        raw=$(rpm -q --qf '%{VERSION}' linglong-bin 2>/dev/null || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi
    if command -v pacman >/dev/null 2>&1; then
        raw=$(pacman -Qi linyaps 2>/dev/null | awk -F ': ' '/^Version/ {print $2; exit}' || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
        raw=$(pacman -Qi linglong-bin 2>/dev/null | awk -F ': ' '/^Version/ {print $2; exit}' || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi
    return 1
}

version_lt() {
    local a="$1"
    local b="$2"
    if command -v dpkg >/dev/null 2>&1; then
        dpkg --compare-versions "${a}" lt "${b}"
        return $?
    fi
    local first
    first=$(printf '%s\n' "${a}" "${b}" | sort -V | head -n1)
    [ "${first}" = "${a}" ] && [ "${a}" != "${b}" ]
}

check_linglong_installed() {
    if command -v ll-cli >/dev/null 2>&1; then
        local version
        version=$(get_linglong_version || true)
        if [ -z "${version}" ]; then
            version="unknown"
        fi
        info "Linglong runtime detected: ${version}"
        return 0
    fi

    warn "Linglong runtime not installed"
    return 1
}

linglong_needs_install() {
    if ! command -v ll-cli >/dev/null 2>&1; then
        return 0
    fi

    local current_version=""
    current_version=$(get_linglong_version || true)
    if [ -z "${current_version}" ]; then
        warn "Linglong version unknown; proceeding with upgrade"
        return 0
    fi

    if version_lt "${current_version}" "${MIN_LINGLONG_VERSION}"; then
        info "Linglong version ${current_version} is below ${MIN_LINGLONG_VERSION}, upgrading"
        return 0
    fi

    info "Linglong runtime already installed (version ${current_version}), no upgrade needed"
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
