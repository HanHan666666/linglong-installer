#!/usr/bin/env bash
# =============================================================================
# 玲珑(Linyaps)环境一键安装脚本 - 纯 Shell 版本
# =============================================================================
#
# 用途：自动检测 Linux 发行版，添加软件源并安装玲珑运行环境(linglong-bin)
# 使用：sudo bash install-linyaps-env.sh
#
# 支持的发行版：
#   Debian 12/13/Testing/Sid, Ubuntu 24.04/25.04/25.10,
#   Deepin 23/23.1/25, UOS 20, openKylin 2.0,
#   Fedora 41/42/43/44/Rawhide, Evernight Vista 44（映射到 Fedora）, AnolisOS 8,
#   openEuler 23.09/24.03, openSUSE 15.6,
#   Arch Linux, Manjaro, Parabola,
#   Linux Mint (映射到 Ubuntu), MX Linux (映射到 Debian)
#
# 环境变量（可选）：
#   LLI_MIN_LINGLONG_VERSION  最低版本要求，低于此版本将触发升级（默认 1.9.0）
#   LLI_SKIP_VERSION_CHECK    设为 1 跳过版本检测，强制重新安装
#
# 注意：NixOS 不支持自动安装，请手动配置 services.linyaps.enable = true
# =============================================================================
set -euo pipefail

# --------------------------------- 配置 --------------------------------------
MIN_LINGLONG_VERSION="${LLI_MIN_LINGLONG_VERSION:-1.9.0}"
SKIP_VERSION_CHECK="${LLI_SKIP_VERSION_CHECK:-0}"

# --------------------------------- 日志工具 -----------------------------------
info()  { echo "[INFO]  $*"; }
warn()  { echo "[WARN]  $*"; }
error() { echo "[ERROR] $*" >&2; }

# --------------------------------- 权限检查 -----------------------------------
# 必须以 root 权限运行
check_root() {
    if [ "$(id -u)" -ne 0 ]; then
        error "此脚本需要 root 权限运行，请使用 sudo 或以 root 身份执行"
        exit 1
    fi
}

# --------------------------------- 版本工具 -----------------------------------

# 获取已安装的玲珑版本号
# 依次尝试 ll-cli、dpkg、rpm、pacman
get_linglong_version() {
    local raw=""

    # 方式1: 通过 ll-cli 命令获取
    if command -v ll-cli >/dev/null 2>&1; then
        raw=$(ll-cli --version 2>/dev/null | head -n1 || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi

    # 方式2: 通过 dpkg 查询（Debian/Ubuntu 系）
    if command -v dpkg-query >/dev/null 2>&1; then
        raw=$(dpkg-query -W -f='${Version}' linglong-bin 2>/dev/null || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi

    # 方式3: 通过 rpm 查询（Fedora/openEuler/AnolisOS 系）
    if command -v rpm >/dev/null 2>&1; then
        raw=$(rpm -q --qf '%{VERSION}' linglong-bin 2>/dev/null || true)
        if [[ "${raw}" =~ ([0-9]+\.[0-9]+\.[0-9]+) ]]; then
            echo "${BASH_REMATCH[1]}"
            return 0
        fi
    fi

    # 方式4: 通过 pacman 查询（Arch/Manjaro/Parabola 系）
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

# 版本比较：version_lt A B → A < B 时返回 0
version_lt() {
    local a="$1" b="$2"
    if command -v dpkg >/dev/null 2>&1; then
        dpkg --compare-versions "${a}" lt "${b}"
        return $?
    fi
    local first
    first=$(printf '%s\n' "${a}" "${b}" | sort -V | head -n1)
    [ "${first}" = "${a}" ] && [ "${a}" != "${b}" ]
}

# 判断是否需要安装/升级玲珑
# 返回 0 表示需要安装，返回 1 表示已满足要求
linglong_needs_install() {
    if [ "${SKIP_VERSION_CHECK}" = "1" ]; then
        info "跳过版本检测，强制安装"
        return 0
    fi

    if ! command -v ll-cli >/dev/null 2>&1; then
        info "未检测到玲珑运行环境，需要安装"
        return 0
    fi

    local current_version=""
    current_version=$(get_linglong_version || true)
    if [ -z "${current_version}" ]; then
        warn "无法获取玲珑版本号，将继续安装"
        return 0
    fi

    if version_lt "${current_version}" "${MIN_LINGLONG_VERSION}"; then
        info "当前版本 ${current_version} 低于最低要求 ${MIN_LINGLONG_VERSION}，需要升级"
        return 0
    fi

    info "玲珑环境已安装（版本 ${current_version}），无需操作"
    return 1
}

# --------------------------------- 系统识别 -----------------------------------

# 读取 /etc/os-release 并提取指定字段
# 用法: os_release_field "ID"
os_release_field() {
    local key="$1"
    if [ -f /etc/os-release ]; then
        # 匹配 KEY="value" 或 KEY=value，去除引号
        sed -n "s/^${key}=//p" /etc/os-release | tr -d '"' | head -n1
    fi
}

# 检测当前发行版 ID 和版本
detect_distro() {
    DISTRO_ID=$(os_release_field "ID")
    DISTRO_VERSION=$(os_release_field "VERSION_ID")
    DISTRO_ID_LIKE=$(os_release_field "ID_LIKE")
    DISTRO_CODENAME=$(os_release_field "VERSION_CODENAME")
    UBUNTU_CODENAME=$(os_release_field "UBUNTU_CODENAME")
    DEBIAN_CODENAME=$(os_release_field "DEBIAN_CODENAME")
    PRETTY_NAME=$(os_release_field "PRETTY_NAME")

    # 转小写
    DISTRO_ID=$(echo "${DISTRO_ID}" | tr '[:upper:]' '[:lower:]')
    DISTRO_ID_LIKE=$(echo "${DISTRO_ID_LIKE}" | tr '[:upper:]' '[:lower:]')

    info "检测到系统: ${PRETTY_NAME:-${DISTRO_ID} ${DISTRO_VERSION}}"
    info "发行版 ID: ${DISTRO_ID}, 版本: ${DISTRO_VERSION:-未知}"
}

# --------------------------------- 软件源管理 ---------------------------------

# 添加 APT 软件源（ci.deepin.com）
# 参数: $1 = 仓库路径（如 Debian_12, xUbuntu_24.04）
add_apt_repo() {
    local repo_path="$1"
    local repo_file="/etc/apt/sources.list.d/linglong.list"
    local repo_url="https://ci.deepin.com/repo/obs/linglong:/CI:/release/${repo_path}/"

    if [ -f "${repo_file}" ]; then
        info "APT 软件源已存在，跳过添加"
    else
        info "添加 APT 软件源: ${repo_url}"
        echo "deb [trusted=yes] ${repo_url} ./" > "${repo_file}"
    fi

    info "刷新 APT 元数据..."
    apt update
}

# 添加 OBS 镜像 APT 软件源（obs-ci.odata.cc）
# 参数: $1 = 仓库路径, $2 = GPG keyring 文件名, $3 = sources.list 文件名
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
        info "APT 仓库签名密钥已存在，跳过"
    else
        info "获取 APT 仓库签名密钥: ${repo_url}Release.key"
        curl -fsSL "${repo_url}Release.key" | gpg --dearmor | tee "${keyring_path}" >/dev/null
    fi

    if [ -f "${list_path}" ]; then
        info "APT 软件源已存在，跳过添加"
    else
        info "添加 APT 软件源: ${repo_url}"
        echo "deb [arch=${arch} signed-by=${keyring_path}] ${repo_url} ./" > "${list_path}"
    fi

    info "刷新 APT 元数据..."
    apt update
}

# 添加 DNF 软件源
# 参数: $1 = repo 文件 URL
add_dnf_repo() {
    local repo_url="$1"

    if ls /etc/yum.repos.d/linglong*.repo >/dev/null 2>&1; then
        info "DNF 软件源已存在，跳过添加"
    else
        info "添加 DNF 软件源: ${repo_url}"
        dnf config-manager addrepo --from-repofile "${repo_url}" || \
        dnf config-manager --add-repo "${repo_url}"
    fi

    info "刷新 DNF 元数据..."
    dnf update -y --refresh
}

# 关闭 linglong 仓库的 GPG 校验（openEuler 需要）
set_dnf_gpgcheck_off() {
    if ls /etc/yum.repos.d/linglong*.repo >/dev/null 2>&1; then
        info "关闭玲珑仓库的 gpgcheck"
        sh -c "echo gpgcheck=0 >> /etc/yum.repos.d/linglong*.repo"
    fi
}

# --------------------------------- 安装验证 -----------------------------------

# 验证玲珑是否安装成功
check_linglong_installed() {
    if command -v ll-cli >/dev/null 2>&1; then
        local version
        version=$(get_linglong_version || true)
        info "玲珑环境安装成功，版本: ${version:-未知}"
        return 0
    fi

    error "安装后未检测到 ll-cli 命令，安装可能失败"
    return 1
}

# --------------------------------- 衍生版映射 ---------------------------------

# Linux Mint → Ubuntu 版本映射
mint_to_ubuntu() {
    local version="$1"
    local major="${version%%.*}"
    case "${major}" in
        22) echo "24.04" ;;
        21) echo "22.04" ;;
        20) echo "20.04" ;;
        19) echo "18.04" ;;
        *)  return 1 ;;
    esac
}

# MX Linux → Debian 版本映射
mx_to_debian() {
    local version="$1"
    local major="${version%%.*}"
    case "${major}" in
        23) echo "12" ;;
        22) echo "11" ;;
        21) echo "11" ;;
        19) echo "10" ;;
        *)  return 1 ;;
    esac
}

# Ubuntu codename → 版本号映射
ubuntu_codename_to_version() {
    local codename
    codename=$(echo "$1" | tr '[:upper:]' '[:lower:]')
    case "${codename}" in
        noble)   echo "24.04" ;;
        mantic)  echo "23.10" ;;
        lunar)   echo "23.04" ;;
        kinetic) echo "22.10" ;;
        jammy)   echo "22.04" ;;
        focal)   echo "20.04" ;;
        bionic)  echo "18.04" ;;
        xenial)  echo "16.04" ;;
        *)       return 1 ;;
    esac
}

# Debian codename → 版本号映射
debian_codename_to_version() {
    local codename
    codename=$(echo "$1" | tr '[:upper:]' '[:lower:]')
    case "${codename}" in
        trixie)   echo "13" ;;
        bookworm) echo "12" ;;
        bullseye) echo "11" ;;
        buster)   echo "10" ;;
        sid)      echo "sid" ;;
        testing)  echo "testing" ;;
        *)        return 1 ;;
    esac
}

# Fedora 版本标签规范化：空版本统一按 Rawhide 处理。
normalize_fedora_version() {
    local version
    version=$(echo "$1" | tr '[:upper:]' '[:lower:]')
    if [ -z "${version}" ] || [ "${version}" = "rawhide" ]; then
        echo "rawhide"
        return 0
    fi
    echo "${version}"
}

# --------------------------------- 发行版安装逻辑 -----------------------------

is_arch_like_distro() {
    case "${DISTRO_ID}" in
        arch|archlinux|manjaro|parabola|garuda)
            return 0
            ;;
    esac

    local token
    for token in ${DISTRO_ID_LIKE:-}; do
        case "${token}" in
            arch|archlinux|manjaro|parabola)
                return 0
                ;;
        esac
    done

    return 1
}

# Debian 系安装
install_debian() {
    local version="$1"
    case "${version}" in
        12)
            add_apt_repo "Debian_12"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        13)
            add_apt_repo "Debian_13"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        sid)
            add_obs_apt_repo "Debian_Unstable" "obs-debian-unstable.gpg" "obs-debian-unstable.list"
            apt install -y linglong-bin linglong-box
            ;;
        testing)
            add_obs_apt_repo "Debian_Testing" "obs-debian-testing.gpg" "obs-debian-testing.list"
            apt install -y linglong-bin linglong-box
            ;;
        *)
            error "不支持的 Debian 版本: ${version}"
            return 1
            ;;
    esac
}

# Ubuntu 系安装
install_ubuntu() {
    local version="$1"
    case "${version}" in
        24.04)
            add_apt_repo "xUbuntu_24.04"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        25.04)
            add_obs_apt_repo "xUbuntu_25.04" "obs-ubuntu2504.gpg" "obs-ubuntu2504.list"
            apt install -y linglong-bin linglong-box
            ;;
        25.10)
            add_obs_apt_repo "Ubuntu_25.10_standard" "obs-ubuntu2510.gpg" "obs-ubuntu2510.list"
            apt install -y linglong-bin linglong-box
            ;;
        *)
            error "不支持的 Ubuntu 版本: ${version}"
            return 1
            ;;
    esac
}

# Deepin 系安装
install_deepin() {
    local version="$1"
    case "${version}" in
        23|23.1)
            add_apt_repo "Deepin_23"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        25)
            add_apt_repo "Deepin_25"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        *)
            error "不支持的 Deepin 版本: ${version}"
            return 1
            ;;
    esac
}

# UOS 安装
install_uos() {
    local version="$1"
    case "${version}" in
        20)
            add_apt_repo "uos_1070"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        *)
            error "不支持的 UOS 版本: ${version}"
            return 1
            ;;
    esac
}

# openKylin 安装
install_openkylin() {
    local version="$1"
    case "${version}" in
        2.0)
            add_apt_repo "openkylin_2.0"
            apt install -y linglong-bin linglong-installer linglong-box policykit-1
            ;;
        *)
            error "不支持的 openKylin 版本: ${version}"
            return 1
            ;;
    esac
}

# Fedora 系安装
install_fedora() {
    local version="$1"
    case "${version}" in
        41)
            add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/Fedora_41/linglong%3ACI%3Arelease.repo"
            dnf install -y linglong-bin linyaps-web-store-installer policykit-1
            ;;
        42)
            add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/Fedora_42/linglong%3ACI%3Arelease.repo"
            dnf install -y linglong-bin linyaps-web-store-installer policykit-1
            ;;
        43|44)
            # Fedora 44 与 Evernight Vista 44 当前复用 Fedora 43 的 COPR 安装路径。
            dnf copr enable mozixun/OpenAtom-Linyaps -y
            dnf install -y linglong linglong-bin linyaps-web-store-installer
            ;;
        rawhide)
            dnf copr enable mozixun/OpenAtom-Linyaps -y
            dnf install -y linglong linglong-bin linyaps-web-store-installer
            ;;
        *)
            error "不支持的 Fedora 版本: ${version}"
            return 1
            ;;
    esac
}

# AnolisOS 安装
install_anolis() {
    local version="$1"
    case "${version}" in
        8|8.*)
            add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/AnolisOS_8/linglong%3ACI%3Arelease.repo"
            dnf install -y linglong-bin linyaps-web-store-installer policykit-1
            ;;
        *)
            error "不支持的 AnolisOS 版本: ${version}"
            return 1
            ;;
    esac
}

# openEuler 安装
install_openeuler() {
    local version="$1"
    case "${version}" in
        23.09)
            add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/openEuler_23.09/linglong%3ACI%3Arelease.repo"
            set_dnf_gpgcheck_off
            dnf install -y linglong-bin linyaps-web-store-installer policykit-1
            ;;
        24.03)
            add_dnf_repo "https://ci.deepin.com/repo/obs/linglong:/CI:/release/openEuler_24.03/linglong%3ACI%3Arelease.repo"
            set_dnf_gpgcheck_off
            dnf install -y linglong-bin linyaps-web-store-installer policykit-1
            ;;
        *)
            error "不支持的 openEuler 版本: ${version}"
            return 1
            ;;
    esac
}

# openSUSE 安装
install_opensuse() {
    local version="$1"
    case "${version}" in
        15.6)
            zypper addrepo https://download.opensuse.org/repositories/home:/guanzi:/Fedora/15.6/home:guanzi:Fedora.repo || true
            zypper refresh
            zypper install -y linglong-bin linglong-box linglong-builder
            ;;
        *)
            error "不支持的 openSUSE 版本: ${version}"
            return 1
            ;;
    esac
}

# Arch 系安装（Arch / Manjaro / Parabola）
install_arch_based() {
    pacman -Syu --noconfirm linyaps
}

# --------------------------------- 主入口 -------------------------------------

# 根据检测到的发行版分发到对应的安装函数
dispatch_install() {
    local id="${DISTRO_ID}"
    local version="${DISTRO_VERSION}"

    case "${id}" in
        # ---------- Debian 及其衍生 ----------
        debian)
            # 如果 VERSION_ID 为空，尝试通过 codename 推断
            if [ -z "${version}" ] && [ -n "${DISTRO_CODENAME}" ]; then
                version=$(debian_codename_to_version "${DISTRO_CODENAME}" || true)
            fi
            install_debian "${version}"
            ;;

        ubuntu)
            install_ubuntu "${version}"
            ;;

        deepin)
            install_deepin "${version}"
            ;;

        uos)
            install_uos "${version}"
            ;;

        openkylin)
            install_openkylin "${version}"
            ;;

        # ---------- RPM 系 ----------
        fedora)
            # Fedora Rawhide 通常 VERSION_ID 为空或为 rawhide
            version=$(normalize_fedora_version "${version}")
            install_fedora "${version}"
            ;;

        evernight)
            # Evernight Vista 44 起改用独立 ID，但当前仅确认复用 Fedora 44 的安装链路。
            case "${version}" in
                44|44.*)
                    info "Evernight Vista ${version} → 使用 Fedora 44 安装源"
                    install_fedora "44"
                    ;;
                *)
                    error "当前仅支持 Evernight Vista 44，检测到版本: ${version:-未知}"
                    exit 1
                    ;;
            esac
            ;;

        anolis)
            install_anolis "${version}"
            ;;

        openeuler)
            install_openeuler "${version}"
            ;;

        # ---------- SUSE 系 ----------
        opensuse-leap|opensuse)
            install_opensuse "${version}"
            ;;

        # ---------- Arch 系 ----------
        arch|archlinux)
            install_arch_based
            ;;

        manjaro)
            install_arch_based
            ;;

        parabola)
            install_arch_based
            ;;

        garuda)
            install_arch_based
            ;;

        # ---------- 衍生版映射 ----------
        linuxmint)
            # Linux Mint → 映射到对应 Ubuntu 版本
            local ubuntu_ver=""
            ubuntu_ver=$(mint_to_ubuntu "${version}" || true)
            if [ -z "${ubuntu_ver}" ] && [ -n "${UBUNTU_CODENAME}" ]; then
                ubuntu_ver=$(ubuntu_codename_to_version "${UBUNTU_CODENAME}" || true)
            fi
            if [ -z "${ubuntu_ver}" ]; then
                error "无法将 Linux Mint ${version} 映射到对应的 Ubuntu 版本"
                exit 1
            fi
            info "Linux Mint ${version} → 使用 Ubuntu ${ubuntu_ver} 安装源"
            install_ubuntu "${ubuntu_ver}"
            ;;

        mx|mxlinux)
            # MX Linux → 映射到对应 Debian 版本
            local debian_ver=""
            debian_ver=$(mx_to_debian "${version}" || true)
            if [ -z "${debian_ver}" ] && [ -n "${DEBIAN_CODENAME:-${DISTRO_CODENAME}}" ]; then
                debian_ver=$(debian_codename_to_version "${DEBIAN_CODENAME:-${DISTRO_CODENAME}}" || true)
            fi
            if [ -z "${debian_ver}" ]; then
                error "无法将 MX Linux ${version} 映射到对应的 Debian 版本"
                exit 1
            fi
            info "MX Linux ${version} → 使用 Debian ${debian_ver} 安装源"
            install_debian "${debian_ver}"
            ;;

        # ---------- NixOS 不支持自动安装 ----------
        nixos)
            error "NixOS 不支持自动安装，请手动配置："
            cat <<'EOF'
  1. 编辑 /etc/nixos/configuration.nix，添加:
       services.linyaps.enable = true;
  2. 执行: sudo nixos-rebuild switch
EOF
            exit 1
            ;;

        # ---------- 未知发行版 ----------
        *)
            if is_arch_like_distro; then
                info "检测到 Arch 系衍生版 (${id})，按 Arch Linux 方式安装"
                install_arch_based
                return 0
            fi
            error "不支持的发行版: ${id} ${version}"
            error "支持列表: Debian, Ubuntu, Deepin, UOS, openKylin, Fedora, Evernight Vista 44, AnolisOS, openEuler, openSUSE, Arch, Manjaro, Parabola, Linux Mint, MX Linux"
            exit 1
            ;;
    esac
}

# --------------------------------- 主流程 -------------------------------------
main() {
    info "=========================================="
    info " 玲珑(Linyaps)环境安装脚本"
    info "=========================================="

    # 1. 检查 root 权限
    check_root

    # 2. 检测发行版
    detect_distro

    # 3. 判断是否需要安装
    if ! linglong_needs_install; then
        info "安装完成，无需操作"
        exit 0
    fi

    # 4. 执行对应发行版的安装流程
    info "开始安装玲珑环境..."
    dispatch_install

    # 5. 验证安装结果
    check_linglong_installed

    info "=========================================="
    info " 安装完成！"
    info "=========================================="
}

main "$@"
