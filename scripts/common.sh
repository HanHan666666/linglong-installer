#!/usr/bin/env bash
set -e

SCRIPT_ARGS=("$@")
MIN_LINGLONG_VERSION="${LLI_MIN_LINGLONG_VERSION:-1.9.0}"

info() { echo "[信息] $*"; }
warn() { echo "[警告] $*"; }
error() { echo "[错误] $*" >&2; }

has_desktop_session() {
    [ -n "${DISPLAY:-}" ] || [ -n "${WAYLAND_DISPLAY:-}" ] || [ -n "${DBUS_SESSION_BUS_ADDRESS:-}" ] || \
    [ -n "${XDG_SESSION_TYPE:-}" ] || [ -n "${XDG_SESSION_ID:-}" ]
}

should_use_sudo_for_ll_cli() {
    if [ "${LLI_FORCE_SUDO:-}" = "1" ]; then
        return 0
    fi
    if [ -n "${LIMA_INSTANCE:-}" ] || [ -n "${LIMA_CIDATA:-}" ]; then
        return 0
    fi
    if ! has_desktop_session; then
        return 0
    fi
    return 1
}

run_ll_cli() {
    if should_use_sudo_for_ll_cli && command -v sudo >/dev/null 2>&1; then
        sudo -E ll-cli "$@"
        return $?
    fi
    ll-cli "$@"
}

require_root() {
    if [ "$(id -u)" -eq 0 ]; then
        return 0
    fi

    if [ "${LLI_ELEVATED:-}" = "1" ]; then
        error "需要 root 权限，但提权失败"
        exit 1
    fi

    local script_path="${BASH_SOURCE[1]:-$0}"
    if [ -z "${script_path}" ]; then
        error "无法解析用于提权的脚本路径"
        exit 1
    fi
    script_path="$(cd "$(dirname "${script_path}")" && pwd)/$(basename "${script_path}")"

    # Prefer pkexec on desktop sessions (graphical password dialog, better UX)
    if command -v pkexec >/dev/null 2>&1 && has_desktop_session; then
        exec pkexec env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
    fi

    # No graphical auth available, fall back to sudo (requires terminal input)
    if command -v sudo >/dev/null 2>&1; then
        if [ -n "${LLI_SUDO_HINT:-}" ]; then
            info "${LLI_SUDO_HINT}"
        fi
        exec sudo -E env LLI_ELEVATED=1 bash "${script_path}" "${SCRIPT_ARGS[@]}"
    fi

    error "需要 root 权限，但 pkexec/sudo 不可用"
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
            version="未知"
        fi
        info "检测到玲珑运行时：${version}"
        return 0
    fi

    warn "未安装玲珑运行时"
    return 1
}

linglong_needs_install() {
    if ! command -v ll-cli >/dev/null 2>&1; then
        return 0
    fi

    local current_version=""
    current_version=$(get_linglong_version || true)
    if [ -z "${current_version}" ]; then
        warn "玲珑版本未知，继续执行升级"
        return 0
    fi

    if version_lt "${current_version}" "${MIN_LINGLONG_VERSION}"; then
        info "玲珑版本 ${current_version} 低于 ${MIN_LINGLONG_VERSION}，正在升级"
        return 0
    fi

    info "玲珑运行时已安装（版本 ${current_version}），无需升级"
    return 1
}

add_apt_repo() {
    local repo_path="$1"
    local repo_file="/etc/apt/sources.list.d/linglong.list"
    local repo_url="https://ci.deepin.com/repo/obs/linglong:/CI:/release/${repo_path}/"

    if [ -f "${repo_file}" ]; then
        info "玲珑 APT 仓库已存在，跳过"
    else
        info "正在添加 APT 仓库：${repo_url}"
        echo "deb [trusted=yes] ${repo_url} ./" > "${repo_file}"
    fi

    info "正在刷新 APT 元数据..."
    apt update
}

add_obs_apt_repo() {
    local repo_path="$1"
    local keyring_name="$2"
    local list_name="$3"
    local base_url="https://obs-ci.odata.cc/obs-mirror"
    local keyring_dir="/usr/share/keyrings"
    local keyring_path="${keyring_dir}/${keyring_name}"
    local list_path="/etc/apt/sources.list.d/${list_name}"
    local repo_url="${base_url}/${repo_path}/"
    local arch

    arch=$(dpkg --print-architecture)

    install -d -m 0755 "${keyring_dir}"

    if [ -f "${keyring_path}" ]; then
        info "APT 仓库密钥已存在，跳过"
    else
        info "正在获取 APT 仓库密钥：${repo_url}Release.key"
        curl -fsSL "${repo_url}Release.key" | gpg --dearmor | tee "${keyring_path}" >/dev/null
    fi

    if [ -f "${list_path}" ]; then
        info "玲珑 APT 仓库已存在，跳过"
    else
        info "正在添加 APT 仓库：${repo_url}"
        echo "deb [arch=${arch} signed-by=${keyring_path}] ${repo_url} ./" > "${list_path}"
    fi

    info "正在刷新 APT 元数据..."
    apt update
}

add_dnf_repo() {
    local repo_url="$1"

    if ls /etc/yum.repos.d/linglong*.repo >/dev/null 2>&1; then
        info "玲珑 DNF 仓库已存在，跳过"
    else
        info "正在添加 DNF 仓库：${repo_url}"
        dnf config-manager addrepo --from-repofile "${repo_url}" || \
        dnf config-manager --add-repo "${repo_url}"
    fi

    info "正在刷新 DNF 元数据..."
    dnf update -y --refresh
}

set_dnf_gpgcheck_off() {
    if ls /etc/yum.repos.d/linglong*.repo >/dev/null 2>&1; then
        info "正在为玲珑仓库禁用 gpgcheck"
        sh -c "echo gpgcheck=0 >> /etc/yum.repos.d/linglong*.repo"
    fi
}
