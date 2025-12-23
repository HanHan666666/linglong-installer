// Code generated for darwin/arm64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -ignore-link-errors -extended-errors -o libxau.go --package-name libxau .libs/libXau.a', DO NOT EDIT.

//go:build darwin && arm64

package libxau

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_BADSIG = "SIG_ERR"
const m_BIG_ENDIAN = "__DARWIN_BIG_ENDIAN"
const m_BUFSIZ = 1024
const m_BUS_ADRALN = 1
const m_BUS_ADRERR = 2
const m_BUS_NOOP = 0
const m_BUS_OBJERR = 3
const m_BYTE_ORDER = "__DARWIN_BYTE_ORDER"
const m_CLD_CONTINUED = 6
const m_CLD_DUMPED = 3
const m_CLD_EXITED = 1
const m_CLD_KILLED = 2
const m_CLD_NOOP = 0
const m_CLD_STOPPED = 5
const m_CLD_TRAPPED = 4
const m_CPUMON_MAKE_FATAL = 0x1000
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_FILENAME_MAX = 1024
const m_FOOTPRINT_INTERVAL_RESET = 0x1
const m_FOPEN_MAX = 20
const m_FPE_FLTDIV = 1
const m_FPE_FLTINV = 5
const m_FPE_FLTOVF = 2
const m_FPE_FLTRES = 4
const m_FPE_FLTSUB = 6
const m_FPE_FLTUND = 3
const m_FPE_INTDIV = 7
const m_FPE_INTOVF = 8
const m_FPE_NOOP = 0
const m_FUNCPROTO = 15
const m_FamilyKrb5Principal = 253
const m_FamilyLocal = 256
const m_FamilyLocalHost = 252
const m_FamilyNetname = 254
const m_FamilyWild = 65535
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_PATHCONF = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDIO_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_HAVE_WCHAR_H = 1
const m_ILL_BADSTK = 8
const m_ILL_COPROC = 7
const m_ILL_ILLADR = 5
const m_ILL_ILLOPC = 1
const m_ILL_ILLOPN = 4
const m_ILL_ILLTRP = 2
const m_ILL_NOOP = 0
const m_ILL_PRVOPC = 3
const m_ILL_PRVREG = 6
const m_INTMAX_MAX = "__INTMAX_MAX__"
const m_INTPTR_MAX = "__INTPTR_MAX__"
const m_INT_FAST16_MAX = "__INT_LEAST16_MAX"
const m_INT_FAST16_MIN = "__INT_LEAST16_MIN"
const m_INT_FAST32_MAX = "__INT_LEAST32_MAX"
const m_INT_FAST32_MIN = "__INT_LEAST32_MIN"
const m_INT_FAST64_MAX = "__INT_LEAST64_MAX"
const m_INT_FAST64_MIN = "__INT_LEAST64_MIN"
const m_INT_FAST8_MAX = "__INT_LEAST8_MAX"
const m_INT_FAST8_MIN = "__INT_LEAST8_MIN"
const m_INT_LEAST16_MAX = "__INT_LEAST16_MAX"
const m_INT_LEAST16_MIN = "__INT_LEAST16_MIN"
const m_INT_LEAST32_MAX = "__INT_LEAST32_MAX"
const m_INT_LEAST32_MIN = "__INT_LEAST32_MIN"
const m_INT_LEAST64_MAX = "__INT_LEAST64_MAX"
const m_INT_LEAST64_MIN = "__INT_LEAST64_MIN"
const m_INT_LEAST8_MAX = "__INT_LEAST8_MAX"
const m_INT_LEAST8_MIN = "__INT_LEAST8_MIN"
const m_IOPOL_APPLICATION = "IOPOL_STANDARD"
const m_IOPOL_ATIME_UPDATES_DEFAULT = 0
const m_IOPOL_ATIME_UPDATES_OFF = 1
const m_IOPOL_DEFAULT = 0
const m_IOPOL_IMPORTANT = 1
const m_IOPOL_MATERIALIZE_DATALESS_FILES_DEFAULT = 0
const m_IOPOL_MATERIALIZE_DATALESS_FILES_OFF = 1
const m_IOPOL_MATERIALIZE_DATALESS_FILES_ON = 2
const m_IOPOL_NORMAL = "IOPOL_IMPORTANT"
const m_IOPOL_PASSIVE = 2
const m_IOPOL_SCOPE_DARWIN_BG = 2
const m_IOPOL_SCOPE_PROCESS = 0
const m_IOPOL_SCOPE_THREAD = 1
const m_IOPOL_STANDARD = 5
const m_IOPOL_THROTTLE = 3
const m_IOPOL_TYPE_DISK = 0
const m_IOPOL_TYPE_VFS_ALLOW_LOW_SPACE_WRITES = 9
const m_IOPOL_TYPE_VFS_ATIME_UPDATES = 2
const m_IOPOL_TYPE_VFS_DISALLOW_RW_FOR_O_EVTONLY = 10
const m_IOPOL_TYPE_VFS_IGNORE_CONTENT_PROTECTION = 6
const m_IOPOL_TYPE_VFS_IGNORE_PERMISSIONS = 7
const m_IOPOL_TYPE_VFS_MATERIALIZE_DATALESS_FILES = 3
const m_IOPOL_TYPE_VFS_SKIP_MTIME_UPDATE = 8
const m_IOPOL_TYPE_VFS_STATFS_NO_DATA_VOLUME = 4
const m_IOPOL_TYPE_VFS_TRIGGER_RESOLVE = 5
const m_IOPOL_UTILITY = 4
const m_IOPOL_VFS_ALLOW_LOW_SPACE_WRITES_OFF = 0
const m_IOPOL_VFS_ALLOW_LOW_SPACE_WRITES_ON = 1
const m_IOPOL_VFS_CONTENT_PROTECTION_DEFAULT = 0
const m_IOPOL_VFS_CONTENT_PROTECTION_IGNORE = 1
const m_IOPOL_VFS_DISALLOW_RW_FOR_O_EVTONLY_DEFAULT = 0
const m_IOPOL_VFS_DISALLOW_RW_FOR_O_EVTONLY_ON = 1
const m_IOPOL_VFS_IGNORE_PERMISSIONS_OFF = 0
const m_IOPOL_VFS_IGNORE_PERMISSIONS_ON = 1
const m_IOPOL_VFS_NOCACHE_WRITE_FS_BLKSIZE_DEFAULT = 0
const m_IOPOL_VFS_NOCACHE_WRITE_FS_BLKSIZE_ON = 1
const m_IOPOL_VFS_SKIP_MTIME_UPDATE_IGNORE = 2
const m_IOPOL_VFS_SKIP_MTIME_UPDATE_OFF = 0
const m_IOPOL_VFS_SKIP_MTIME_UPDATE_ON = 1
const m_IOPOL_VFS_STATFS_FORCE_NO_DATA_VOLUME = 1
const m_IOPOL_VFS_STATFS_NO_DATA_VOLUME_DEFAULT = 0
const m_IOPOL_VFS_TRIGGER_RESOLVE_DEFAULT = 0
const m_IOPOL_VFS_TRIGGER_RESOLVE_OFF = 1
const m_LITTLE_ENDIAN = "__DARWIN_LITTLE_ENDIAN"
const m_LOCK_ERROR = 1
const m_LOCK_SUCCESS = 0
const m_LOCK_TIMEOUT = 2
const m_LT_OBJDIR = ".libs/"
const m_L_ctermid = 1024
const m_L_tmpnam = 1024
const m_MAC_OS_VERSION_11_0 = "__MAC_11_0"
const m_MAC_OS_VERSION_11_1 = "__MAC_11_1"
const m_MAC_OS_VERSION_11_3 = "__MAC_11_3"
const m_MAC_OS_VERSION_11_4 = "__MAC_11_4"
const m_MAC_OS_VERSION_11_5 = "__MAC_11_5"
const m_MAC_OS_VERSION_11_6 = "__MAC_11_6"
const m_MAC_OS_VERSION_12_0 = "__MAC_12_0"
const m_MAC_OS_VERSION_12_1 = "__MAC_12_1"
const m_MAC_OS_VERSION_12_2 = "__MAC_12_2"
const m_MAC_OS_VERSION_12_3 = "__MAC_12_3"
const m_MAC_OS_VERSION_12_4 = "__MAC_12_4"
const m_MAC_OS_VERSION_12_5 = "__MAC_12_5"
const m_MAC_OS_VERSION_12_6 = "__MAC_12_6"
const m_MAC_OS_VERSION_12_7 = "__MAC_12_7"
const m_MAC_OS_VERSION_13_0 = "__MAC_13_0"
const m_MAC_OS_VERSION_13_1 = "__MAC_13_1"
const m_MAC_OS_VERSION_13_2 = "__MAC_13_2"
const m_MAC_OS_VERSION_13_3 = "__MAC_13_3"
const m_MAC_OS_VERSION_13_4 = "__MAC_13_4"
const m_MAC_OS_VERSION_13_5 = "__MAC_13_5"
const m_MAC_OS_VERSION_13_6 = "__MAC_13_6"
const m_MAC_OS_VERSION_13_7 = "__MAC_13_7"
const m_MAC_OS_VERSION_14_0 = "__MAC_14_0"
const m_MAC_OS_VERSION_14_1 = "__MAC_14_1"
const m_MAC_OS_VERSION_14_2 = "__MAC_14_2"
const m_MAC_OS_VERSION_14_3 = "__MAC_14_3"
const m_MAC_OS_VERSION_14_4 = "__MAC_14_4"
const m_MAC_OS_VERSION_14_5 = "__MAC_14_5"
const m_MAC_OS_VERSION_14_6 = "__MAC_14_6"
const m_MAC_OS_VERSION_14_7 = "__MAC_14_7"
const m_MAC_OS_VERSION_15_0 = "__MAC_15_0"
const m_MAC_OS_VERSION_15_1 = "__MAC_15_1"
const m_MAC_OS_VERSION_15_2 = "__MAC_15_2"
const m_MAC_OS_VERSION_15_3 = "__MAC_15_3"
const m_MAC_OS_VERSION_15_4 = "__MAC_15_4"
const m_MAC_OS_X_VERSION_10_0 = "__MAC_10_0"
const m_MAC_OS_X_VERSION_10_1 = "__MAC_10_1"
const m_MAC_OS_X_VERSION_10_10 = "__MAC_10_10"
const m_MAC_OS_X_VERSION_10_10_2 = "__MAC_10_10_2"
const m_MAC_OS_X_VERSION_10_10_3 = "__MAC_10_10_3"
const m_MAC_OS_X_VERSION_10_11 = "__MAC_10_11"
const m_MAC_OS_X_VERSION_10_11_2 = "__MAC_10_11_2"
const m_MAC_OS_X_VERSION_10_11_3 = "__MAC_10_11_3"
const m_MAC_OS_X_VERSION_10_11_4 = "__MAC_10_11_4"
const m_MAC_OS_X_VERSION_10_12 = "__MAC_10_12"
const m_MAC_OS_X_VERSION_10_12_1 = "__MAC_10_12_1"
const m_MAC_OS_X_VERSION_10_12_2 = "__MAC_10_12_2"
const m_MAC_OS_X_VERSION_10_12_4 = "__MAC_10_12_4"
const m_MAC_OS_X_VERSION_10_13 = "__MAC_10_13"
const m_MAC_OS_X_VERSION_10_13_1 = "__MAC_10_13_1"
const m_MAC_OS_X_VERSION_10_13_2 = "__MAC_10_13_2"
const m_MAC_OS_X_VERSION_10_13_4 = "__MAC_10_13_4"
const m_MAC_OS_X_VERSION_10_14 = "__MAC_10_14"
const m_MAC_OS_X_VERSION_10_14_1 = "__MAC_10_14_1"
const m_MAC_OS_X_VERSION_10_14_4 = "__MAC_10_14_4"
const m_MAC_OS_X_VERSION_10_14_5 = "__MAC_10_14_5"
const m_MAC_OS_X_VERSION_10_14_6 = "__MAC_10_14_6"
const m_MAC_OS_X_VERSION_10_15 = "__MAC_10_15"
const m_MAC_OS_X_VERSION_10_15_1 = "__MAC_10_15_1"
const m_MAC_OS_X_VERSION_10_15_4 = "__MAC_10_15_4"
const m_MAC_OS_X_VERSION_10_16 = "__MAC_10_16"
const m_MAC_OS_X_VERSION_10_2 = "__MAC_10_2"
const m_MAC_OS_X_VERSION_10_3 = "__MAC_10_3"
const m_MAC_OS_X_VERSION_10_4 = "__MAC_10_4"
const m_MAC_OS_X_VERSION_10_5 = "__MAC_10_5"
const m_MAC_OS_X_VERSION_10_6 = "__MAC_10_6"
const m_MAC_OS_X_VERSION_10_7 = "__MAC_10_7"
const m_MAC_OS_X_VERSION_10_8 = "__MAC_10_8"
const m_MAC_OS_X_VERSION_10_9 = "__MAC_10_9"
const m_MB_CUR_MAX = "__mb_cur_max"
const m_MINSIGSTKSZ = 32768
const m_NDEBUG = 1
const m_NSIG = "__DARWIN_NSIG"
const m_NULL = "__DARWIN_NULL"
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 1
const m_PACKAGE = "libXau"
const m_PACKAGE_BUGREPORT = "https://gitlab.freedesktop.org/xorg/lib/libxau/-/issues"
const m_PACKAGE_NAME = "libXau"
const m_PACKAGE_STRING = "libXau 1.0.11"
const m_PACKAGE_TARNAME = "libXau"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "1.0.11"
const m_PACKAGE_VERSION_MAJOR = 1
const m_PACKAGE_VERSION_MINOR = 0
const m_PACKAGE_VERSION_PATCHLEVEL = 11
const m_PDP_ENDIAN = "__DARWIN_PDP_ENDIAN"
const m_POLL_ERR = 4
const m_POLL_HUP = 6
const m_POLL_IN = 1
const m_POLL_MSG = 3
const m_POLL_OUT = 2
const m_POLL_PRI = 5
const m_PRIO_DARWIN_BG = 0x1000
const m_PRIO_DARWIN_NONUI = 0x1001
const m_PRIO_DARWIN_PROCESS = 4
const m_PRIO_DARWIN_THREAD = 3
const m_PRIO_MAX = 20
const m_PRIO_PGRP = 1
const m_PRIO_PROCESS = 0
const m_PRIO_USER = 2
const m_PTRDIFF_MAX = "__PTRDIFF_MAX__"
const m_P_tmpdir = "/var/tmp/"
const m_RAND_MAX = 0x7fffffff
const m_RENAME_EXCL = 0x00000004
const m_RENAME_NOFOLLOW_ANY = 0x00000010
const m_RENAME_RESERVED1 = 0x00000008
const m_RENAME_SECLUDE = 0x00000001
const m_RENAME_SWAP = 0x00000002
const m_RLIMIT_AS = 5
const m_RLIMIT_CORE = 4
const m_RLIMIT_CPU = 0
const m_RLIMIT_CPU_USAGE_MONITOR = 0x2
const m_RLIMIT_DATA = 2
const m_RLIMIT_FOOTPRINT_INTERVAL = 0x4
const m_RLIMIT_FSIZE = 1
const m_RLIMIT_MEMLOCK = 6
const m_RLIMIT_NOFILE = 8
const m_RLIMIT_NPROC = 7
const m_RLIMIT_RSS = "RLIMIT_AS"
const m_RLIMIT_STACK = 3
const m_RLIMIT_THREAD_CPULIMITS = 0x3
const m_RLIMIT_WAKEUPS_MONITOR = 0x1
const m_RLIM_NLIMITS = 9
const m_RLIM_SAVED_CUR = "RLIM_INFINITY"
const m_RLIM_SAVED_MAX = "RLIM_INFINITY"
const m_RUSAGE_INFO_CURRENT = "RUSAGE_INFO_V6"
const m_RUSAGE_INFO_V0 = 0
const m_RUSAGE_INFO_V1 = 1
const m_RUSAGE_INFO_V2 = 2
const m_RUSAGE_INFO_V3 = 3
const m_RUSAGE_INFO_V4 = 4
const m_RUSAGE_INFO_V5 = 5
const m_RUSAGE_INFO_V6 = 6
const m_RUSAGE_SELF = 0
const m_RU_PROC_RUNS_RESLIDE = 0x00000001
const m_SA_64REGSET = 0x0200
const m_SA_NOCLDSTOP = 0x0008
const m_SA_NOCLDWAIT = 0x0020
const m_SA_NODEFER = 0x0010
const m_SA_ONSTACK = 0x0001
const m_SA_RESETHAND = 0x0004
const m_SA_RESTART = 0x0002
const m_SA_SIGINFO = 0x0040
const m_SA_USERTRAMP = 0x0100
const m_SEEK_CUR = 1
const m_SEEK_DATA = 4
const m_SEEK_END = 2
const m_SEEK_HOLE = 3
const m_SEEK_SET = 0
const m_SEGV_ACCERR = 2
const m_SEGV_MAPERR = 1
const m_SEGV_NOOP = 0
const m_SIGABRT = 6
const m_SIGALRM = 14
const m_SIGBUS = 10
const m_SIGCHLD = 20
const m_SIGCONT = 19
const m_SIGEMT = 7
const m_SIGEV_NONE = 0
const m_SIGEV_SIGNAL = 1
const m_SIGEV_THREAD = 3
const m_SIGFPE = 8
const m_SIGHUP = 1
const m_SIGILL = 4
const m_SIGINFO = 29
const m_SIGINT = 2
const m_SIGIO = 23
const m_SIGIOT = "SIGABRT"
const m_SIGKILL = 9
const m_SIGPIPE = 13
const m_SIGPROF = 27
const m_SIGQUIT = 3
const m_SIGSEGV = 11
const m_SIGSTKSZ = 131072
const m_SIGSTOP = 17
const m_SIGSYS = 12
const m_SIGTERM = 15
const m_SIGTRAP = 5
const m_SIGTSTP = 18
const m_SIGTTIN = 21
const m_SIGTTOU = 22
const m_SIGURG = 16
const m_SIGUSR1 = 30
const m_SIGUSR2 = 31
const m_SIGVTALRM = 26
const m_SIGWINCH = 28
const m_SIGXCPU = 24
const m_SIGXFSZ = 25
const m_SIG_BLOCK = 1
const m_SIG_SETMASK = 3
const m_SIG_UNBLOCK = 2
const m_SIZE_MAX = "__SIZE_MAX__"
const m_SI_ASYNCIO = 0x10004
const m_SI_MESGQ = 0x10005
const m_SI_QUEUE = 0x10002
const m_SI_TIMER = 0x10003
const m_SI_USER = 0x10001
const m_SS_DISABLE = 0x0004
const m_SS_ONSTACK = 0x0001
const m_STDC_HEADERS = 1
const m_SV_INTERRUPT = "SA_RESTART"
const m_SV_NOCLDSTOP = "SA_NOCLDSTOP"
const m_SV_NODEFER = "SA_NODEFER"
const m_SV_ONSTACK = "SA_ONSTACK"
const m_SV_RESETHAND = "SA_RESETHAND"
const m_SV_SIGINFO = "SA_SIGINFO"
const m_TARGET_IPHONE_SIMULATOR = 0
const m_TARGET_OS_ARROW = 1
const m_TARGET_OS_BRIDGE = 0
const m_TARGET_OS_DRIVERKIT = 0
const m_TARGET_OS_EMBEDDED = 0
const m_TARGET_OS_IOS = 0
const m_TARGET_OS_IOSMAC = 0
const m_TARGET_OS_IPHONE = 0
const m_TARGET_OS_LINUX = 0
const m_TARGET_OS_MAC = 1
const m_TARGET_OS_MACCATALYST = 0
const m_TARGET_OS_NANO = 0
const m_TARGET_OS_OSX = 1
const m_TARGET_OS_SIMULATOR = 0
const m_TARGET_OS_TV = 0
const m_TARGET_OS_UIKITFORMAC = 0
const m_TARGET_OS_UNIX = 0
const m_TARGET_OS_VISION = 0
const m_TARGET_OS_WATCH = 0
const m_TARGET_OS_WIN32 = 0
const m_TARGET_OS_WINDOWS = 0
const m_TARGET_OS_XR = 0
const m_TMP_MAX = 308915776
const m_TRAP_BRKPT = 1
const m_TRAP_TRACE = 2
const m_UINTMAX_MAX = "__UINTMAX_MAX__"
const m_UINTPTR_MAX = "__UINTPTR_MAX__"
const m_UINT_FAST16_MAX = "__UINT_LEAST16_MAX"
const m_UINT_FAST32_MAX = "__UINT_LEAST32_MAX"
const m_UINT_FAST64_MAX = "__UINT_LEAST64_MAX"
const m_UINT_FAST8_MAX = "__UINT_LEAST8_MAX"
const m_UINT_LEAST16_MAX = "__UINT_LEAST16_MAX"
const m_UINT_LEAST32_MAX = "__UINT_LEAST32_MAX"
const m_UINT_LEAST64_MAX = "__UINT_LEAST64_MAX"
const m_UINT_LEAST8_MAX = "__UINT_LEAST8_MAX"
const m_VERSION = "1.0.11"
const m_WAIT_MYPGRP = 0
const m_WAKEMON_DISABLE = 0x02
const m_WAKEMON_ENABLE = 0x01
const m_WAKEMON_GET_PARAMS = 0x04
const m_WAKEMON_MAKE_FATAL = 0x10
const m_WAKEMON_SET_DEFAULTS = 0x08
const m_WCHAR_MAX = "__WCHAR_MAX__"
const m_WCONTINUED = 0x00000010
const m_WCOREFLAG = 0200
const m_WEXITED = 0x00000004
const m_WNOHANG = 0x00000001
const m_WNOWAIT = 0x00000020
const m_WSTOPPED = 0x00000008
const m_WUNTRACED = 0x00000002
const m_XTHREADS = 1
const m__ALL_SOURCE = 1
const m__ARM_SIGNAL_ = 1
const m__DARWIN_C_SOURCE = 1
const m__DARWIN_FEATURE_64_BIT_INODE = 1
const m__DARWIN_FEATURE_ONLY_64_BIT_INODE = 1
const m__DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
const m__DARWIN_FEATURE_ONLY_VERS_1050 = 1
const m__DARWIN_FEATURE_UNIX_CONFORMANCE = 3
const m__FORTIFY_SOURCE = 2
const m__GNU_SOURCE = 1
const m__HPUX_ALT_XOPEN_SOCKET_API = 1
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
const m__LIBC_COUNT__MB_LEN_MAX = "_LIBC_UNSAFE_INDEXABLE"
const m__LIBC_COUNT__PATH_MAX = "_LIBC_UNSAFE_INDEXABLE"
const m__LP64 = 1
const m__NETBSD_SOURCE = 1
const m__OPENBSD_SOURCE = 1
const m__POSIX_PTHREAD_SEMANTICS = 1
const m__QUAD_HIGHWORD = 1
const m__QUAD_LOWWORD = 0
const m__RLIMIT_POSIX_FLAG = 0x1000
const m__STRUCT_MCONTEXT = "_STRUCT_MCONTEXT64"
const m__TANDEM_SOURCE = 1
const m__WSTOPPED = 0177
const m__X_INLINE = "inline"
const m__X_RESTRICT_KYWD = "restrict"
const m__Xconst = "const"
const m___AARCH64EL__ = 1
const m___AARCH64_CMODEL_SMALL__ = 1
const m___AARCH64_SIMD__ = 1
const m___API_TO_BE_DEPRECATED = 100000
const m___API_TO_BE_DEPRECATED_DRIVERKIT = 100000
const m___API_TO_BE_DEPRECATED_IOS = 100000
const m___API_TO_BE_DEPRECATED_IOSAPPLICATIONEXTENSION = 100000
const m___API_TO_BE_DEPRECATED_KERNELKIT = 100000
const m___API_TO_BE_DEPRECATED_MACCATALYST = 100000
const m___API_TO_BE_DEPRECATED_MACCATALYSTAPPLICATIONEXTENSION = 100000
const m___API_TO_BE_DEPRECATED_MACOS = 100000
const m___API_TO_BE_DEPRECATED_MACOSAPPLICATIONEXTENSION = 100000
const m___API_TO_BE_DEPRECATED_TVOS = 100000
const m___API_TO_BE_DEPRECATED_TVOSAPPLICATIONEXTENSION = 100000
const m___API_TO_BE_DEPRECATED_VISIONOS = 100000
const m___API_TO_BE_DEPRECATED_VISIONOSAPPLICATIONEXTENSION = 100000
const m___API_TO_BE_DEPRECATED_WATCHOS = 100000
const m___API_TO_BE_DEPRECATED_WATCHOSAPPLICATIONEXTENSION = 100000
const m___APPLE_CC__ = 6000
const m___APPLE__ = 1
const m___ARM64_ARCH_8__ = 1
const m___ARM_64BIT_STATE = 1
const m___ARM_ACLE = 200
const m___ARM_ALIGN_MAX_STACK_PWR = 4
const m___ARM_ARCH = 8
const m___ARM_ARCH_8_3__ = 1
const m___ARM_ARCH_8_4__ = 1
const m___ARM_ARCH_8_5__ = 1
const m___ARM_ARCH_ISA_A64 = 1
const m___ARM_ARCH_PROFILE = 'A'
const m___ARM_FEATURE_AES = 1
const m___ARM_FEATURE_ATOMICS = 1
const m___ARM_FEATURE_BTI = 1
const m___ARM_FEATURE_CLZ = 1
const m___ARM_FEATURE_COMPLEX = 1
const m___ARM_FEATURE_CRC32 = 1
const m___ARM_FEATURE_CRYPTO = 1
const m___ARM_FEATURE_DIRECTED_ROUNDING = 1
const m___ARM_FEATURE_DIV = 1
const m___ARM_FEATURE_DOTPROD = 1
const m___ARM_FEATURE_FMA = 1
const m___ARM_FEATURE_FP16_FML = 1
const m___ARM_FEATURE_FP16_SCALAR_ARITHMETIC = 1
const m___ARM_FEATURE_FP16_VECTOR_ARITHMETIC = 1
const m___ARM_FEATURE_FRINT = 1
const m___ARM_FEATURE_IDIV = 1
const m___ARM_FEATURE_JCVT = 1
const m___ARM_FEATURE_LDREX = 0xF
const m___ARM_FEATURE_NUMERIC_MAXMIN = 1
const m___ARM_FEATURE_PAUTH = 1
const m___ARM_FEATURE_QRDMX = 1
const m___ARM_FEATURE_RCPC = 1
const m___ARM_FEATURE_SHA2 = 1
const m___ARM_FEATURE_SHA3 = 1
const m___ARM_FEATURE_SHA512 = 1
const m___ARM_FEATURE_UNALIGNED = 1
const m___ARM_FP = 0xE
const m___ARM_FP16_ARGS = 1
const m___ARM_FP16_FORMAT_IEEE = 1
const m___ARM_NEON = 1
const m___ARM_NEON_FP = 0xE
const m___ARM_NEON__ = 1
const m___ARM_PCS_AAPCS64 = 1
const m___ARM_SIZEOF_MINIMAL_ENUM = 4
const m___ARM_SIZEOF_WCHAR_T = 4
const m___ARM_STATE_ZA = 1
const m___ARM_STATE_ZT0 = 1
const m___ATOMIC_ACQUIRE = 2
const m___ATOMIC_ACQ_REL = 4
const m___ATOMIC_CONSUME = 1
const m___ATOMIC_RELAXED = 0
const m___ATOMIC_RELEASE = 3
const m___ATOMIC_SEQ_CST = 5
const m___AVAILABILITY_FILE = "AvailabilityVersions.h"
const m___AVAILABILITY_VERSIONS_VERSION_HASH = 93585900
const m___AVAILABILITY_VERSIONS_VERSION_STRING = "Local"
const m___BIGGEST_ALIGNMENT__ = 8
const m___BITINT_MAXWIDTH__ = 128
const m___BLOCKS__ = 1
const m___BOOL_WIDTH__ = 8
const m___BRIDGEOS_2_0 = 20000
const m___BRIDGEOS_3_0 = 30000
const m___BRIDGEOS_3_1 = 30100
const m___BRIDGEOS_3_4 = 30400
const m___BRIDGEOS_4_0 = 40000
const m___BRIDGEOS_4_1 = 40100
const m___BRIDGEOS_5_0 = 50000
const m___BRIDGEOS_5_1 = 50100
const m___BRIDGEOS_5_3 = 50300
const m___BRIDGEOS_6_0 = 60000
const m___BRIDGEOS_6_2 = 60200
const m___BRIDGEOS_6_4 = 60400
const m___BRIDGEOS_6_5 = 60500
const m___BRIDGEOS_6_6 = 60600
const m___BRIDGEOS_7_0 = 70000
const m___BRIDGEOS_7_1 = 70100
const m___BRIDGEOS_7_2 = 70200
const m___BRIDGEOS_7_3 = 70300
const m___BRIDGEOS_7_4 = 70400
const m___BRIDGEOS_7_6 = 70600
const m___BRIDGEOS_8_0 = 80000
const m___BRIDGEOS_8_1 = 80100
const m___BRIDGEOS_8_2 = 80200
const m___BRIDGEOS_8_3 = 80300
const m___BRIDGEOS_8_4 = 80400
const m___BRIDGEOS_8_5 = 80500
const m___BRIDGEOS_8_6 = 80600
const m___BRIDGEOS_9_0 = 90000
const m___BRIDGEOS_9_1 = 90100
const m___BRIDGEOS_9_2 = 90200
const m___BRIDGEOS_9_3 = 90300
const m___BRIDGEOS_9_4 = 90400
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___CCGO__ = 1
const m___CHAR_BIT__ = 8
const m___CLANG_ATOMIC_BOOL_LOCK_FREE = 2
const m___CLANG_ATOMIC_CHAR16_T_LOCK_FREE = 2
const m___CLANG_ATOMIC_CHAR32_T_LOCK_FREE = 2
const m___CLANG_ATOMIC_CHAR_LOCK_FREE = 2
const m___CLANG_ATOMIC_INT_LOCK_FREE = 2
const m___CLANG_ATOMIC_LLONG_LOCK_FREE = 2
const m___CLANG_ATOMIC_LONG_LOCK_FREE = 2
const m___CLANG_ATOMIC_POINTER_LOCK_FREE = 2
const m___CLANG_ATOMIC_SHORT_LOCK_FREE = 2
const m___CLANG_ATOMIC_WCHAR_T_LOCK_FREE = 2
const m___CONSTANT_CFSTRINGS__ = 1
const m___DARWIN_64_BIT_INO_T = 1
const m___DARWIN_BIG_ENDIAN = 4321
const m___DARWIN_BYTE_ORDER = "__DARWIN_LITTLE_ENDIAN"
const m___DARWIN_C_ANSI = 010000
const m___DARWIN_C_FULL = 900000
const m___DARWIN_C_LEVEL = "__DARWIN_C_FULL"
const m___DARWIN_LITTLE_ENDIAN = 1234
const m___DARWIN_NON_CANCELABLE = 0
const m___DARWIN_NO_LONG_LONG = 0
const m___DARWIN_NSIG = 32
const m___DARWIN_ONLY_64_BIT_INO_T = 1
const m___DARWIN_ONLY_UNIX_CONFORMANCE = 1
const m___DARWIN_ONLY_VERS_1050 = 1
const m___DARWIN_OPAQUE_ARM_THREAD_STATE64 = 0
const m___DARWIN_PDP_ENDIAN = 3412
const m___DARWIN_SUF_EXTSN = "$DARWIN_EXTSN"
const m___DARWIN_UNIX03 = 1
const m___DARWIN_VERS_1050 = 1
const m___DARWIN_WCHAR_MAX = "__WCHAR_MAX__"
const m___DBL_DECIMAL_DIG__ = 17
const m___DBL_DENORM_MIN__ = 4.9406564584124654e-324
const m___DBL_DIG__ = 15
const m___DBL_EPSILON__ = 2.2204460492503131e-16
const m___DBL_HAS_DENORM__ = 1
const m___DBL_HAS_INFINITY__ = 1
const m___DBL_HAS_QUIET_NAN__ = 1
const m___DBL_MANT_DIG__ = 53
const m___DBL_MAX_10_EXP__ = 308
const m___DBL_MAX_EXP__ = 1024
const m___DBL_MAX__ = 1.7976931348623157e+308
const m___DBL_MIN__ = 2.2250738585072014e-308
const m___DBL_NORM_MAX__ = 1.7976931348623157e+308
const m___DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const m___DRIVERKIT_19_0 = 190000
const m___DRIVERKIT_20_0 = 200000
const m___DRIVERKIT_21_0 = 210000
const m___DRIVERKIT_22_0 = 220000
const m___DRIVERKIT_22_4 = 220400
const m___DRIVERKIT_22_5 = 220500
const m___DRIVERKIT_22_6 = 220600
const m___DRIVERKIT_23_0 = 230000
const m___DRIVERKIT_23_1 = 230100
const m___DRIVERKIT_23_2 = 230200
const m___DRIVERKIT_23_3 = 230300
const m___DRIVERKIT_23_4 = 230400
const m___DRIVERKIT_23_5 = 230500
const m___DRIVERKIT_23_6 = 230600
const m___DRIVERKIT_24_0 = 240000
const m___DRIVERKIT_24_1 = 240100
const m___DRIVERKIT_24_2 = 240200
const m___DRIVERKIT_24_3 = 240300
const m___DRIVERKIT_24_4 = 240400
const m___DYNAMIC__ = 1
const m___ENABLE_LEGACY_MAC_AVAILABILITY = 1
const m___ENVIRONMENT_MAC_OS_X_VERSION_MIN_REQUIRED__ = 150000
const m___ENVIRONMENT_OS_VERSION_MIN_REQUIRED__ = 150000
const m___EXTENSIONS__ = 1
const m___FINITE_MATH_ONLY__ = 0
const m___FLT16_DECIMAL_DIG__ = 5
const m___FLT16_DENORM_MIN__ = 5.9604644775390625e-8
const m___FLT16_DIG__ = 3
const m___FLT16_EPSILON__ = 9.765625e-4
const m___FLT16_HAS_DENORM__ = 1
const m___FLT16_HAS_INFINITY__ = 1
const m___FLT16_HAS_QUIET_NAN__ = 1
const m___FLT16_MANT_DIG__ = 11
const m___FLT16_MAX_10_EXP__ = 4
const m___FLT16_MAX_EXP__ = 16
const m___FLT16_MAX__ = 6.5504e+4
const m___FLT16_MIN__ = 6.103515625e-5
const m___FLT16_NORM_MAX__ = 6.5504e+4
const m___FLT_DECIMAL_DIG__ = 9
const m___FLT_DENORM_MIN__ = 1.40129846e-45
const m___FLT_DIG__ = 6
const m___FLT_EPSILON__ = 1.19209290e-7
const m___FLT_HAS_DENORM__ = 1
const m___FLT_HAS_INFINITY__ = 1
const m___FLT_HAS_QUIET_NAN__ = 1
const m___FLT_MANT_DIG__ = 24
const m___FLT_MAX_10_EXP__ = 38
const m___FLT_MAX_EXP__ = 128
const m___FLT_MAX__ = 3.40282347e+38
const m___FLT_MIN__ = 1.17549435e-38
const m___FLT_NORM_MAX__ = 3.40282347e+38
const m___FLT_RADIX__ = 2
const m___FPCLASS_NEGINF = 0x0004
const m___FPCLASS_NEGNORMAL = 0x0008
const m___FPCLASS_NEGSUBNORMAL = 0x0010
const m___FPCLASS_NEGZERO = 0x0020
const m___FPCLASS_POSINF = 0x0200
const m___FPCLASS_POSNORMAL = 0x0100
const m___FPCLASS_POSSUBNORMAL = 0x0080
const m___FPCLASS_POSZERO = 0x0040
const m___FPCLASS_QNAN = 0x0002
const m___FPCLASS_SNAN = 0x0001
const m___FP_FAST_FMA = 1
const m___FP_FAST_FMAF = 1
const m___FUNCTION__ = "__func__"
const m___GCC_ASM_FLAG_OUTPUTS__ = 1
const m___GCC_ATOMIC_BOOL_LOCK_FREE = 2
const m___GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const m___GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const m___GCC_ATOMIC_CHAR_LOCK_FREE = 2
const m___GCC_ATOMIC_INT_LOCK_FREE = 2
const m___GCC_ATOMIC_LLONG_LOCK_FREE = 2
const m___GCC_ATOMIC_LONG_LOCK_FREE = 2
const m___GCC_ATOMIC_POINTER_LOCK_FREE = 2
const m___GCC_ATOMIC_SHORT_LOCK_FREE = 2
const m___GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const m___GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const m___GCC_CONSTRUCTIVE_SIZE = 64
const m___GCC_DESTRUCTIVE_SIZE = 64
const m___GCC_HAVE_DWARF2_CFI_ASM = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const m___GNUC_MINOR__ = 2
const m___GNUC_PATCHLEVEL__ = 1
const m___GNUC_STDC_INLINE__ = 1
const m___GNUC__ = 4
const m___GXX_ABI_VERSION = 1002
const m___HAVE_FUNCTION_MULTI_VERSIONING = 1
const m___INT16_FMTd__ = "hd"
const m___INT16_FMTi__ = "hi"
const m___INT16_MAX__ = 32767
const m___INT16_TYPE__ = "short"
const m___INT32_FMTd__ = "d"
const m___INT32_FMTi__ = "i"
const m___INT32_MAX__ = 2147483647
const m___INT32_TYPE__ = "int"
const m___INT64_C_SUFFIX__ = "LL"
const m___INT64_FMTd__ = "lld"
const m___INT64_FMTi__ = "lli"
const m___INT64_MAX__ = 9223372036854775807
const m___INT8_FMTd__ = "hhd"
const m___INT8_FMTi__ = "hhi"
const m___INT8_MAX__ = 127
const m___INTMAX_C_SUFFIX__ = "L"
const m___INTMAX_FMTd__ = "ld"
const m___INTMAX_FMTi__ = "li"
const m___INTMAX_MAX__ = 9223372036854775807
const m___INTMAX_WIDTH__ = 64
const m___INTPTR_FMTd__ = "ld"
const m___INTPTR_FMTi__ = "li"
const m___INTPTR_MAX__ = 9223372036854775807
const m___INTPTR_WIDTH__ = 64
const m___INT_FAST16_FMTd__ = "hd"
const m___INT_FAST16_FMTi__ = "hi"
const m___INT_FAST16_MAX__ = 32767
const m___INT_FAST16_TYPE__ = "short"
const m___INT_FAST16_WIDTH__ = 16
const m___INT_FAST32_FMTd__ = "d"
const m___INT_FAST32_FMTi__ = "i"
const m___INT_FAST32_MAX__ = 2147483647
const m___INT_FAST32_TYPE__ = "int"
const m___INT_FAST32_WIDTH__ = 32
const m___INT_FAST64_FMTd__ = "lld"
const m___INT_FAST64_FMTi__ = "lli"
const m___INT_FAST64_MAX__ = 9223372036854775807
const m___INT_FAST64_WIDTH__ = 64
const m___INT_FAST8_FMTd__ = "hhd"
const m___INT_FAST8_FMTi__ = "hhi"
const m___INT_FAST8_MAX__ = 127
const m___INT_FAST8_WIDTH__ = 8
const m___INT_LEAST16_FMTd__ = "hd"
const m___INT_LEAST16_FMTi__ = "hi"
const m___INT_LEAST16_MAX__ = 32767
const m___INT_LEAST16_TYPE__ = "short"
const m___INT_LEAST16_WIDTH__ = 16
const m___INT_LEAST32_FMTd__ = "d"
const m___INT_LEAST32_FMTi__ = "i"
const m___INT_LEAST32_MAX__ = 2147483647
const m___INT_LEAST32_TYPE__ = "int"
const m___INT_LEAST32_WIDTH__ = 32
const m___INT_LEAST64_FMTd__ = "lld"
const m___INT_LEAST64_FMTi__ = "lli"
const m___INT_LEAST64_MAX = "INT64_MAX"
const m___INT_LEAST64_MAX__ = 9223372036854775807
const m___INT_LEAST64_MIN = "INT64_MIN"
const m___INT_LEAST64_WIDTH__ = 64
const m___INT_LEAST8_FMTd__ = "hhd"
const m___INT_LEAST8_FMTi__ = "hhi"
const m___INT_LEAST8_MAX__ = 127
const m___INT_LEAST8_WIDTH__ = 8
const m___INT_MAX__ = 2147483647
const m___INT_WIDTH__ = 32
const m___IPHONE_10_0 = 100000
const m___IPHONE_10_1 = 100100
const m___IPHONE_10_2 = 100200
const m___IPHONE_10_3 = 100300
const m___IPHONE_11_0 = 110000
const m___IPHONE_11_1 = 110100
const m___IPHONE_11_2 = 110200
const m___IPHONE_11_3 = 110300
const m___IPHONE_11_4 = 110400
const m___IPHONE_12_0 = 120000
const m___IPHONE_12_1 = 120100
const m___IPHONE_12_2 = 120200
const m___IPHONE_12_3 = 120300
const m___IPHONE_12_4 = 120400
const m___IPHONE_13_0 = 130000
const m___IPHONE_13_1 = 130100
const m___IPHONE_13_2 = 130200
const m___IPHONE_13_3 = 130300
const m___IPHONE_13_4 = 130400
const m___IPHONE_13_5 = 130500
const m___IPHONE_13_6 = 130600
const m___IPHONE_13_7 = 130700
const m___IPHONE_14_0 = 140000
const m___IPHONE_14_1 = 140100
const m___IPHONE_14_2 = 140200
const m___IPHONE_14_3 = 140300
const m___IPHONE_14_4 = 140400
const m___IPHONE_14_5 = 140500
const m___IPHONE_14_6 = 140600
const m___IPHONE_14_7 = 140700
const m___IPHONE_14_8 = 140800
const m___IPHONE_15_0 = 150000
const m___IPHONE_15_1 = 150100
const m___IPHONE_15_2 = 150200
const m___IPHONE_15_3 = 150300
const m___IPHONE_15_4 = 150400
const m___IPHONE_15_5 = 150500
const m___IPHONE_15_6 = 150600
const m___IPHONE_15_7 = 150700
const m___IPHONE_15_8 = 150800
const m___IPHONE_16_0 = 160000
const m___IPHONE_16_1 = 160100
const m___IPHONE_16_2 = 160200
const m___IPHONE_16_3 = 160300
const m___IPHONE_16_4 = 160400
const m___IPHONE_16_5 = 160500
const m___IPHONE_16_6 = 160600
const m___IPHONE_16_7 = 160700
const m___IPHONE_17_0 = 170000
const m___IPHONE_17_1 = 170100
const m___IPHONE_17_2 = 170200
const m___IPHONE_17_3 = 170300
const m___IPHONE_17_4 = 170400
const m___IPHONE_17_5 = 170500
const m___IPHONE_17_6 = 170600
const m___IPHONE_17_7 = 170700
const m___IPHONE_18_0 = 180000
const m___IPHONE_18_1 = 180100
const m___IPHONE_18_2 = 180200
const m___IPHONE_18_3 = 180300
const m___IPHONE_18_4 = 180400
const m___IPHONE_2_0 = 20000
const m___IPHONE_2_1 = 20100
const m___IPHONE_2_2 = 20200
const m___IPHONE_3_0 = 30000
const m___IPHONE_3_1 = 30100
const m___IPHONE_3_2 = 30200
const m___IPHONE_4_0 = 40000
const m___IPHONE_4_1 = 40100
const m___IPHONE_4_2 = 40200
const m___IPHONE_4_3 = 40300
const m___IPHONE_5_0 = 50000
const m___IPHONE_5_1 = 50100
const m___IPHONE_6_0 = 60000
const m___IPHONE_6_1 = 60100
const m___IPHONE_7_0 = 70000
const m___IPHONE_7_1 = 70100
const m___IPHONE_8_0 = 80000
const m___IPHONE_8_1 = 80100
const m___IPHONE_8_2 = 80200
const m___IPHONE_8_3 = 80300
const m___IPHONE_8_4 = 80400
const m___IPHONE_9_0 = 90000
const m___IPHONE_9_1 = 90100
const m___IPHONE_9_2 = 90200
const m___IPHONE_9_3 = 90300
const m___LDBL_DECIMAL_DIG__ = 17
const m___LDBL_DENORM_MIN__ = 4.9406564584124654e-324
const m___LDBL_DIG__ = 15
const m___LDBL_EPSILON__ = 2.2204460492503131e-16
const m___LDBL_HAS_DENORM__ = 1
const m___LDBL_HAS_INFINITY__ = 1
const m___LDBL_HAS_QUIET_NAN__ = 1
const m___LDBL_MANT_DIG__ = 53
const m___LDBL_MAX_10_EXP__ = 308
const m___LDBL_MAX_EXP__ = 1024
const m___LDBL_MAX__ = 1.7976931348623157e+308
const m___LDBL_MIN__ = 2.2250738585072014e-308
const m___LDBL_NORM_MAX__ = 1.7976931348623157e+308
const m___LITTLE_ENDIAN__ = 1
const m___LLONG_WIDTH__ = 64
const m___LONG_LONG_MAX__ = 9223372036854775807
const m___LONG_MAX__ = 9223372036854775807
const m___LONG_WIDTH__ = 64
const m___LP64__ = 1
const m___MACH__ = 1
const m___MAC_10_0 = 1000
const m___MAC_10_1 = 1010
const m___MAC_10_10 = 101000
const m___MAC_10_10_2 = 101002
const m___MAC_10_10_3 = 101003
const m___MAC_10_11 = 101100
const m___MAC_10_11_2 = 101102
const m___MAC_10_11_3 = 101103
const m___MAC_10_11_4 = 101104
const m___MAC_10_12 = 101200
const m___MAC_10_12_1 = 101201
const m___MAC_10_12_2 = 101202
const m___MAC_10_12_4 = 101204
const m___MAC_10_13 = 101300
const m___MAC_10_13_1 = 101301
const m___MAC_10_13_2 = 101302
const m___MAC_10_13_4 = 101304
const m___MAC_10_14 = 101400
const m___MAC_10_14_1 = 101401
const m___MAC_10_14_4 = 101404
const m___MAC_10_14_5 = 101405
const m___MAC_10_14_6 = 101406
const m___MAC_10_15 = 101500
const m___MAC_10_15_1 = 101501
const m___MAC_10_15_4 = 101504
const m___MAC_10_16 = 101600
const m___MAC_10_2 = 1020
const m___MAC_10_3 = 1030
const m___MAC_10_4 = 1040
const m___MAC_10_5 = 1050
const m___MAC_10_6 = 1060
const m___MAC_10_7 = 1070
const m___MAC_10_8 = 1080
const m___MAC_10_9 = 1090
const m___MAC_11_0 = 110000
const m___MAC_11_1 = 110100
const m___MAC_11_3 = 110300
const m___MAC_11_4 = 110400
const m___MAC_11_5 = 110500
const m___MAC_11_6 = 110600
const m___MAC_12_0 = 120000
const m___MAC_12_1 = 120100
const m___MAC_12_2 = 120200
const m___MAC_12_3 = 120300
const m___MAC_12_4 = 120400
const m___MAC_12_5 = 120500
const m___MAC_12_6 = 120600
const m___MAC_12_7 = 120700
const m___MAC_13_0 = 130000
const m___MAC_13_1 = 130100
const m___MAC_13_2 = 130200
const m___MAC_13_3 = 130300
const m___MAC_13_4 = 130400
const m___MAC_13_5 = 130500
const m___MAC_13_6 = 130600
const m___MAC_13_7 = 130700
const m___MAC_14_0 = 140000
const m___MAC_14_1 = 140100
const m___MAC_14_2 = 140200
const m___MAC_14_3 = 140300
const m___MAC_14_4 = 140400
const m___MAC_14_5 = 140500
const m___MAC_14_6 = 140600
const m___MAC_14_7 = 140700
const m___MAC_15_0 = 150000
const m___MAC_15_1 = 150100
const m___MAC_15_2 = 150200
const m___MAC_15_3 = 150300
const m___MAC_15_4 = 150400
const m___MAC_OS_X_VERSION_MAX_ALLOWED = "__MAC_15_4"
const m___MAC_OS_X_VERSION_MIN_REQUIRED = "__ENVIRONMENT_MAC_OS_X_VERSION_MIN_REQUIRED__"
const m___MEMORY_SCOPE_DEVICE = 1
const m___MEMORY_SCOPE_SINGLE = 4
const m___MEMORY_SCOPE_SYSTEM = 0
const m___MEMORY_SCOPE_WRKGRP = 2
const m___MEMORY_SCOPE_WVFRNT = 3
const m___NO_INLINE__ = 1
const m___NO_MATH_ERRNO__ = 1
const m___OBJC_BOOL_IS_BOOL = 1
const m___OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const m___OPENCL_MEMORY_SCOPE_DEVICE = 2
const m___OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const m___OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const m___OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const m___ORDER_BIG_ENDIAN__ = 4321
const m___ORDER_LITTLE_ENDIAN__ = 1234
const m___ORDER_PDP_ENDIAN__ = 3412
const m___PIC__ = 2
const m___POINTER_WIDTH__ = 64
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTHREAD_ATTR_SIZE__ = 56
const m___PTHREAD_CONDATTR_SIZE__ = 8
const m___PTHREAD_COND_SIZE__ = 40
const m___PTHREAD_MUTEXATTR_SIZE__ = 8
const m___PTHREAD_MUTEX_SIZE__ = 56
const m___PTHREAD_ONCE_SIZE__ = 8
const m___PTHREAD_RWLOCKATTR_SIZE__ = 16
const m___PTHREAD_RWLOCK_SIZE__ = 192
const m___PTHREAD_SIZE__ = 8176
const m___PTRDIFF_FMTd__ = "ld"
const m___PTRDIFF_FMTi__ = "li"
const m___PTRDIFF_MAX__ = 9223372036854775807
const m___PTRDIFF_WIDTH__ = 64
const m___SALC = 0x4000
const m___SAPP = 0x0100
const m___SCHAR_MAX__ = 127
const m___SEOF = 0x0020
const m___SERR = 0x0040
const m___SHRT_MAX__ = 32767
const m___SHRT_WIDTH__ = 16
const m___SIGN = 0x8000
const m___SIG_ATOMIC_MAX__ = 2147483647
const m___SIG_ATOMIC_WIDTH__ = 32
const m___SIZEOF_DOUBLE__ = 8
const m___SIZEOF_FLOAT__ = 4
const m___SIZEOF_INT128__ = 16
const m___SIZEOF_INT__ = 4
const m___SIZEOF_LONG_DOUBLE__ = 8
const m___SIZEOF_LONG_LONG__ = 8
const m___SIZEOF_LONG__ = 8
const m___SIZEOF_POINTER__ = 8
const m___SIZEOF_PTRDIFF_T__ = 8
const m___SIZEOF_SHORT__ = 2
const m___SIZEOF_SIZE_T__ = 8
const m___SIZEOF_WCHAR_T__ = 4
const m___SIZEOF_WINT_T__ = 4
const m___SIZE_FMTX__ = "lX"
const m___SIZE_FMTo__ = "lo"
const m___SIZE_FMTu__ = "lu"
const m___SIZE_FMTx__ = "lx"
const m___SIZE_MAX__ = 18446744073709551615
const m___SIZE_WIDTH__ = 64
const m___SLBF = 0x0001
const m___SMBF = 0x0080
const m___SMOD = 0x2000
const m___SNBF = 0x0002
const m___SNPT = 0x0800
const m___SOFF = 0x1000
const m___SOPT = 0x0400
const m___SRD = 0x0004
const m___SRW = 0x0010
const m___SSP__ = 1
const m___SSTR = 0x0200
const m___STDC_EMBED_EMPTY__ = 2
const m___STDC_EMBED_FOUND__ = 1
const m___STDC_EMBED_NOT_FOUND__ = 0
const m___STDC_HOSTED__ = 1
const m___STDC_NO_THREADS__ = 1
const m___STDC_UTF_16__ = 1
const m___STDC_UTF_32__ = 1
const m___STDC_VERSION__ = 201710
const m___STDC_WANT_IEC_60559_ATTRIBS_EXT__ = 1
const m___STDC_WANT_IEC_60559_BFP_EXT__ = 1
const m___STDC_WANT_IEC_60559_DFP_EXT__ = 1
const m___STDC_WANT_IEC_60559_FUNCS_EXT__ = 1
const m___STDC_WANT_IEC_60559_TYPES_EXT__ = 1
const m___STDC_WANT_LIB_EXT1__ = 1
const m___STDC_WANT_LIB_EXT2__ = 1
const m___STDC_WANT_MATH_SPEC_FUNCS__ = 1
const m___STDC__ = 1
const m___SWR = 0x0008
const m___TVOS_10_0 = 100000
const m___TVOS_10_0_1 = 100001
const m___TVOS_10_1 = 100100
const m___TVOS_10_2 = 100200
const m___TVOS_11_0 = 110000
const m___TVOS_11_1 = 110100
const m___TVOS_11_2 = 110200
const m___TVOS_11_3 = 110300
const m___TVOS_11_4 = 110400
const m___TVOS_12_0 = 120000
const m___TVOS_12_1 = 120100
const m___TVOS_12_2 = 120200
const m___TVOS_12_3 = 120300
const m___TVOS_12_4 = 120400
const m___TVOS_13_0 = 130000
const m___TVOS_13_2 = 130200
const m___TVOS_13_3 = 130300
const m___TVOS_13_4 = 130400
const m___TVOS_14_0 = 140000
const m___TVOS_14_1 = 140100
const m___TVOS_14_2 = 140200
const m___TVOS_14_3 = 140300
const m___TVOS_14_5 = 140500
const m___TVOS_14_6 = 140600
const m___TVOS_14_7 = 140700
const m___TVOS_15_0 = 150000
const m___TVOS_15_1 = 150100
const m___TVOS_15_2 = 150200
const m___TVOS_15_3 = 150300
const m___TVOS_15_4 = 150400
const m___TVOS_15_5 = 150500
const m___TVOS_15_6 = 150600
const m___TVOS_16_0 = 160000
const m___TVOS_16_1 = 160100
const m___TVOS_16_2 = 160200
const m___TVOS_16_3 = 160300
const m___TVOS_16_4 = 160400
const m___TVOS_16_5 = 160500
const m___TVOS_16_6 = 160600
const m___TVOS_17_0 = 170000
const m___TVOS_17_1 = 170100
const m___TVOS_17_2 = 170200
const m___TVOS_17_3 = 170300
const m___TVOS_17_4 = 170400
const m___TVOS_17_5 = 170500
const m___TVOS_17_6 = 170600
const m___TVOS_18_0 = 180000
const m___TVOS_18_1 = 180100
const m___TVOS_18_2 = 180200
const m___TVOS_18_3 = 180300
const m___TVOS_18_4 = 180400
const m___TVOS_9_0 = 90000
const m___TVOS_9_1 = 90100
const m___TVOS_9_2 = 90200
const m___UINT16_FMTX__ = "hX"
const m___UINT16_FMTo__ = "ho"
const m___UINT16_FMTu__ = "hu"
const m___UINT16_FMTx__ = "hx"
const m___UINT16_MAX__ = 65535
const m___UINT32_C_SUFFIX__ = "U"
const m___UINT32_FMTX__ = "X"
const m___UINT32_FMTo__ = "o"
const m___UINT32_FMTu__ = "u"
const m___UINT32_FMTx__ = "x"
const m___UINT32_MAX__ = 4294967295
const m___UINT64_C_SUFFIX__ = "ULL"
const m___UINT64_FMTX__ = "llX"
const m___UINT64_FMTo__ = "llo"
const m___UINT64_FMTu__ = "llu"
const m___UINT64_FMTx__ = "llx"
const m___UINT64_MAX__ = "18446744073709551615U"
const m___UINT8_FMTX__ = "hhX"
const m___UINT8_FMTo__ = "hho"
const m___UINT8_FMTu__ = "hhu"
const m___UINT8_FMTx__ = "hhx"
const m___UINT8_MAX__ = 255
const m___UINTMAX_C_SUFFIX__ = "UL"
const m___UINTMAX_FMTX__ = "lX"
const m___UINTMAX_FMTo__ = "lo"
const m___UINTMAX_FMTu__ = "lu"
const m___UINTMAX_FMTx__ = "lx"
const m___UINTMAX_MAX__ = 18446744073709551615
const m___UINTMAX_WIDTH__ = 64
const m___UINTPTR_FMTX__ = "lX"
const m___UINTPTR_FMTo__ = "lo"
const m___UINTPTR_FMTu__ = "lu"
const m___UINTPTR_FMTx__ = "lx"
const m___UINTPTR_MAX__ = 18446744073709551615
const m___UINTPTR_WIDTH__ = 64
const m___UINT_FAST16_FMTX__ = "hX"
const m___UINT_FAST16_FMTo__ = "ho"
const m___UINT_FAST16_FMTu__ = "hu"
const m___UINT_FAST16_FMTx__ = "hx"
const m___UINT_FAST16_MAX__ = 65535
const m___UINT_FAST32_FMTX__ = "X"
const m___UINT_FAST32_FMTo__ = "o"
const m___UINT_FAST32_FMTu__ = "u"
const m___UINT_FAST32_FMTx__ = "x"
const m___UINT_FAST32_MAX__ = 4294967295
const m___UINT_FAST64_FMTX__ = "llX"
const m___UINT_FAST64_FMTo__ = "llo"
const m___UINT_FAST64_FMTu__ = "llu"
const m___UINT_FAST64_FMTx__ = "llx"
const m___UINT_FAST64_MAX__ = "18446744073709551615U"
const m___UINT_FAST8_FMTX__ = "hhX"
const m___UINT_FAST8_FMTo__ = "hho"
const m___UINT_FAST8_FMTu__ = "hhu"
const m___UINT_FAST8_FMTx__ = "hhx"
const m___UINT_FAST8_MAX__ = 255
const m___UINT_LEAST16_FMTX__ = "hX"
const m___UINT_LEAST16_FMTo__ = "ho"
const m___UINT_LEAST16_FMTu__ = "hu"
const m___UINT_LEAST16_FMTx__ = "hx"
const m___UINT_LEAST16_MAX__ = 65535
const m___UINT_LEAST32_FMTX__ = "X"
const m___UINT_LEAST32_FMTo__ = "o"
const m___UINT_LEAST32_FMTu__ = "u"
const m___UINT_LEAST32_FMTx__ = "x"
const m___UINT_LEAST32_MAX__ = 4294967295
const m___UINT_LEAST64_FMTX__ = "llX"
const m___UINT_LEAST64_FMTo__ = "llo"
const m___UINT_LEAST64_FMTu__ = "llu"
const m___UINT_LEAST64_FMTx__ = "llx"
const m___UINT_LEAST64_MAX = "UINT64_MAX"
const m___UINT_LEAST64_MAX__ = "18446744073709551615U"
const m___UINT_LEAST8_FMTX__ = "hhX"
const m___UINT_LEAST8_FMTo__ = "hho"
const m___UINT_LEAST8_FMTu__ = "hhu"
const m___UINT_LEAST8_FMTx__ = "hhx"
const m___UINT_LEAST8_MAX__ = 255
const m___USER_LABEL_PREFIX__ = "_"
const m___VERSION__ = "Apple LLVM 17.0.0 (clang-1700.0.13.3)"
const m___VISIONOS_1_0 = 10000
const m___VISIONOS_1_1 = 10100
const m___VISIONOS_1_2 = 10200
const m___VISIONOS_1_3 = 10300
const m___VISIONOS_2_0 = 20000
const m___VISIONOS_2_1 = 20100
const m___VISIONOS_2_2 = 20200
const m___VISIONOS_2_3 = 20300
const m___VISIONOS_2_4 = 20400
const m___WATCHOS_10_0 = 100000
const m___WATCHOS_10_1 = 100100
const m___WATCHOS_10_2 = 100200
const m___WATCHOS_10_3 = 100300
const m___WATCHOS_10_4 = 100400
const m___WATCHOS_10_5 = 100500
const m___WATCHOS_10_6 = 100600
const m___WATCHOS_10_7 = 100700
const m___WATCHOS_11_0 = 110000
const m___WATCHOS_11_1 = 110100
const m___WATCHOS_11_2 = 110200
const m___WATCHOS_11_3 = 110300
const m___WATCHOS_11_4 = 110400
const m___WATCHOS_1_0 = 10000
const m___WATCHOS_2_0 = 20000
const m___WATCHOS_2_1 = 20100
const m___WATCHOS_2_2 = 20200
const m___WATCHOS_3_0 = 30000
const m___WATCHOS_3_1 = 30100
const m___WATCHOS_3_1_1 = 30101
const m___WATCHOS_3_2 = 30200
const m___WATCHOS_4_0 = 40000
const m___WATCHOS_4_1 = 40100
const m___WATCHOS_4_2 = 40200
const m___WATCHOS_4_3 = 40300
const m___WATCHOS_5_0 = 50000
const m___WATCHOS_5_1 = 50100
const m___WATCHOS_5_2 = 50200
const m___WATCHOS_5_3 = 50300
const m___WATCHOS_6_0 = 60000
const m___WATCHOS_6_1 = 60100
const m___WATCHOS_6_2 = 60200
const m___WATCHOS_7_0 = 70000
const m___WATCHOS_7_1 = 70100
const m___WATCHOS_7_2 = 70200
const m___WATCHOS_7_3 = 70300
const m___WATCHOS_7_4 = 70400
const m___WATCHOS_7_5 = 70500
const m___WATCHOS_7_6 = 70600
const m___WATCHOS_8_0 = 80000
const m___WATCHOS_8_1 = 80100
const m___WATCHOS_8_3 = 80300
const m___WATCHOS_8_4 = 80400
const m___WATCHOS_8_5 = 80500
const m___WATCHOS_8_6 = 80600
const m___WATCHOS_8_7 = 80700
const m___WATCHOS_8_8 = 80800
const m___WATCHOS_9_0 = 90000
const m___WATCHOS_9_1 = 90100
const m___WATCHOS_9_2 = 90200
const m___WATCHOS_9_3 = 90300
const m___WATCHOS_9_4 = 90400
const m___WATCHOS_9_5 = 90500
const m___WATCHOS_9_6 = 90600
const m___WCHAR_MAX__ = 2147483647
const m___WCHAR_TYPE__ = "int"
const m___WCHAR_WIDTH__ = 32
const m___WINT_MAX__ = 2147483647
const m___WINT_TYPE__ = "int"
const m___WINT_WIDTH__ = 32
const m___aarch64__ = 1
const m___apple_build_version__ = 17000013
const m___arm64 = 1
const m___arm64__ = 1
const m___clang__ = 1
const m___clang_literal_encoding__ = "UTF-8"
const m___clang_major__ = 17
const m___clang_minor__ = 0
const m___clang_patchlevel__ = 0
const m___clang_version__ = "17.0.0 (clang-1700.0.13.3)"
const m___clang_wide_literal_encoding__ = "UTF-32"
const m___const = "const"
const m___has_bounds_safety_attributes = 0
const m___has_ptrcheck = 0
const m___has_safe_buffers = 0
const m___header_inline = "inline"
const m___llvm__ = 1
const m___nonnull = "_Nonnull"
const m___null_unspecified = "_Null_unspecified"
const m___nullable = "_Nullable"
const m___pic__ = 2
const m___restrict = "restrict"
const m___restrict_arr = "restrict"
const m___signed = "signed"
const m___volatile = "volatile"
const m_ru_first = "ru_ixrss"
const m_ru_last = "ru_nivcsw"
const m_stderr = "__stderrp"
const m_stdin = "__stdinp"
const m_stdout = "__stdoutp"
const m_sv_onstack = "sv_flags"

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint64

type t__predefined_wchar_t = int32

type t__predefined_ptrdiff_t = int64

type TXauth = struct {
	Ffamily         uint16
	Faddress_length uint16
	Faddress        uintptr
	Fnumber_length  uint16
	Fnumber         uintptr
	Fname_length    uint16
	Fname           uintptr
	Fdata_length    uint16
	Fdata           uintptr
}

type Txauth = TXauth

type t__int8_t = int8

type t__uint8_t = uint8

type t__int16_t = int16

type t__uint16_t = uint16

type t__int32_t = int32

type t__uint32_t = uint32

type t__int64_t = int64

type t__uint64_t = uint64

type t__darwin_intptr_t = int64

type t__darwin_natural_t = uint32

type t__darwin_ct_rune_t = int32

type t__mbstate_t = struct {
	F_mbstateL  [0]int64
	F__mbstate8 [128]int8
}

type t__darwin_mbstate_t = struct {
	F_mbstateL  [0]int64
	F__mbstate8 [128]int8
}

type t__darwin_ptrdiff_t = int64

type t__darwin_size_t = uint64

type t__darwin_va_list = uintptr

type t__darwin_wchar_t = int32

type t__darwin_rune_t = int32

type t__darwin_wint_t = int32

type t__darwin_clock_t = uint64

type t__darwin_socklen_t = uint32

type t__darwin_ssize_t = int64

type t__darwin_time_t = int64

type t__darwin_blkcnt_t = int64

type t__darwin_blksize_t = int32

type t__darwin_dev_t = int32

type t__darwin_fsblkcnt_t = uint32

type t__darwin_fsfilcnt_t = uint32

type t__darwin_gid_t = uint32

type t__darwin_id_t = uint32

type t__darwin_ino64_t = uint64

type t__darwin_ino_t = uint64

type t__darwin_mach_port_name_t = uint32

type t__darwin_mach_port_t = uint32

type t__darwin_mode_t = uint16

type t__darwin_off_t = int64

type t__darwin_pid_t = int32

type t__darwin_sigset_t = uint32

type t__darwin_suseconds_t = int32

type t__darwin_uid_t = uint32

type t__darwin_useconds_t = uint32

type t__darwin_uuid_t = [16]uint8

type t__darwin_uuid_string_t = [37]int8

type t__darwin_pthread_handler_rec = struct {
	F__routine uintptr
	F__arg     uintptr
	F__next    uintptr
}

type T_opaque_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type T_opaque_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type T_opaque_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type T_opaque_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type T_opaque_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type T_opaque_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type T_opaque_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type T_opaque_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type T_opaque_pthread_t = struct {
	F__sig           int64
	F__cleanup_stack uintptr
	F__opaque        [8176]int8
}

type t__darwin_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type t__darwin_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type t__darwin_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type t__darwin_pthread_key_t = uint64

type t__darwin_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type t__darwin_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type t__darwin_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type t__darwin_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type t__darwin_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type t__darwin_pthread_t = uintptr

type t__darwin_nl_item = int32

type t__darwin_wctrans_t = int32

type t__darwin_wctype_t = uint32

type Tsize_t = uint64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tu_int8_t = uint8

type Tu_int16_t = uint16

type Tu_int32_t = uint32

type Tu_int64_t = uint64

type Tregister_t = int64

type Tintptr_t = int64

type Tuintptr_t = uint64

type Tuser_addr_t = uint64

type Tuser_size_t = uint64

type Tuser_ssize_t = int64

type Tuser_long_t = int64

type Tuser_ulong_t = uint64

type Tuser_time_t = int64

type Tuser_off_t = int64

type Tsyscall_arg_t = uint64

type Trsize_t = uint64

type Terrno_t = int32

type Tssize_t = int64

type Tva_list = uintptr

type Tfpos_t = int64

type t__sbuf = struct {
	F_base uintptr
	F_size int32
}

type TFILE = struct {
	F_p       uintptr
	F_r       int32
	F_w       int32
	F_flags   int16
	F_file    int16
	F_bf      t__sbuf
	F_lbfsize int32
	F_cookie  uintptr
	F_close   uintptr
	F_read    uintptr
	F_seek    uintptr
	F_write   uintptr
	F_ub      t__sbuf
	F_extra   uintptr
	F_ur      int32
	F_ubuf    [3]uint8
	F_nbuf    [1]uint8
	F_lb      t__sbuf
	F_blksize int32
	F_offset  Tfpos_t
}

type t__sFILE = TFILE

type Toff_t = int64

type Tidtype_t = int32

const _P_ALL = 0
const _P_PID = 1
const _P_PGID = 2

type Tpid_t = int32

type Tid_t = uint32

type Tsig_atomic_t = int32

type t__darwin_arm_exception_state = struct {
	F__exception t__uint32_t
	F__fsr       t__uint32_t
	F__far       t__uint32_t
}

type t__darwin_arm_exception_state64 = struct {
	F__far       t__uint64_t
	F__esr       t__uint32_t
	F__exception t__uint32_t
}

type t__darwin_arm_exception_state64_v2 = struct {
	F__far t__uint64_t
	F__esr t__uint64_t
}

type t__darwin_arm_thread_state = struct {
	F__r    [13]t__uint32_t
	F__sp   t__uint32_t
	F__lr   t__uint32_t
	F__pc   t__uint32_t
	F__cpsr t__uint32_t
}

type t__darwin_arm_thread_state64 = struct {
	F__x    [29]t__uint64_t
	F__fp   t__uint64_t
	F__lr   t__uint64_t
	F__sp   t__uint64_t
	F__pc   t__uint64_t
	F__cpsr t__uint32_t
	F__pad  t__uint32_t
}

type t__darwin_arm_vfp_state = struct {
	F__r     [64]t__uint32_t
	F__fpscr t__uint32_t
}

type t__darwin_arm_neon_state64 = int32

type t__darwin_arm_neon_state = int32

type t__arm_pagein_state = struct {
	F__pagein_error int32
}

type t__darwin_arm_sme_state = struct {
	F__svcr       t__uint64_t
	F__tpidr2_el0 t__uint64_t
	F__svl_b      t__uint16_t
}

type t__darwin_arm_sve_z_state = struct {
	F__ccgo_align [0]uint32
	F__z          [16][256]int8
}

type t__darwin_arm_sve_p_state = struct {
	F__ccgo_align [0]uint32
	F__p          [16][32]int8
}

type t__darwin_arm_sme_za_state = struct {
	F__ccgo_align [0]uint32
	F__za         [4096]int8
}

type t__darwin_arm_sme2_state = struct {
	F__ccgo_align [0]uint32
	F__zt0        [64]int8
}

type t__arm_legacy_debug_state = struct {
	F__bvr [16]t__uint32_t
	F__bcr [16]t__uint32_t
	F__wvr [16]t__uint32_t
	F__wcr [16]t__uint32_t
}

type t__darwin_arm_debug_state32 = struct {
	F__bvr       [16]t__uint32_t
	F__bcr       [16]t__uint32_t
	F__wvr       [16]t__uint32_t
	F__wcr       [16]t__uint32_t
	F__mdscr_el1 t__uint64_t
}

type t__darwin_arm_debug_state64 = struct {
	F__bvr       [16]t__uint64_t
	F__bcr       [16]t__uint64_t
	F__wvr       [16]t__uint64_t
	F__wcr       [16]t__uint64_t
	F__mdscr_el1 t__uint64_t
}

type t__darwin_arm_cpmu_state64 = struct {
	F__ctrs [16]t__uint64_t
}

type t__darwin_mcontext32 = struct {
	F__es t__darwin_arm_exception_state
	F__ss t__darwin_arm_thread_state
	F__fs t__darwin_arm_vfp_state
}

type t__darwin_mcontext64 = int32

type Tmcontext_t = uintptr

type Tpthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type t__darwin_sigaltstack = struct {
	Fss_sp    uintptr
	Fss_size  t__darwin_size_t
	Fss_flags int32
}

type Tstack_t = struct {
	Fss_sp    uintptr
	Fss_size  t__darwin_size_t
	Fss_flags int32
}

type t__darwin_ucontext = struct {
	Fuc_onstack  int32
	Fuc_sigmask  t__darwin_sigset_t
	Fuc_stack    t__darwin_sigaltstack
	Fuc_link     uintptr
	Fuc_mcsize   t__darwin_size_t
	Fuc_mcontext uintptr
}

type Tucontext_t = struct {
	Fuc_onstack  int32
	Fuc_sigmask  t__darwin_sigset_t
	Fuc_stack    t__darwin_sigaltstack
	Fuc_link     uintptr
	Fuc_mcsize   t__darwin_size_t
	Fuc_mcontext uintptr
}

type Tsigset_t = uint32

type Tuid_t = uint32

type Tsigval = struct {
	Fsival_ptr   [0]uintptr
	Fsival_int   int32
	F__ccgo_pad2 [4]byte
}

type Tsigevent = struct {
	Fsigev_notify            int32
	Fsigev_signo             int32
	Fsigev_value             Tsigval
	Fsigev_notify_function   uintptr
	Fsigev_notify_attributes uintptr
}

type Tsiginfo_t = struct {
	Fsi_signo  int32
	Fsi_errno  int32
	Fsi_code   int32
	Fsi_pid    Tpid_t
	Fsi_uid    Tuid_t
	Fsi_status int32
	Fsi_addr   uintptr
	Fsi_value  Tsigval
	Fsi_band   int64
	F__pad     [7]uint64
}

type t__siginfo = Tsiginfo_t

type t__sigaction_u = struct {
	F__sa_sigaction [0]uintptr
	F__sa_handler   uintptr
}

type t__sigaction = struct {
	F__sigaction_u t__sigaction_u
	Fsa_tramp      uintptr
	Fsa_mask       Tsigset_t
	Fsa_flags      int32
}

type Tsigaction = struct {
	F__sigaction_u t__sigaction_u
	Fsa_mask       Tsigset_t
	Fsa_flags      int32
}

type Tsig_t = uintptr

type Tsigvec = struct {
	Fsv_handler uintptr
	Fsv_mask    int32
	Fsv_flags   int32
}

type Tsigstack = struct {
	Fss_sp      uintptr
	Fss_onstack int32
}

type Tuint64_t = uint64

type Tint_least64_t = int64

type Tuint_least64_t = uint64

type Tint_fast64_t = int64

type Tuint_fast64_t = uint64

type Tuint32_t = uint32

type Tint_least32_t = int32

type Tuint_least32_t = uint32

type Tint_fast32_t = int32

type Tuint_fast32_t = uint32

type Tuint16_t = uint16

type Tint_least16_t = int16

type Tuint_least16_t = uint16

type Tint_fast16_t = int16

type Tuint_fast16_t = uint16

type Tuint8_t = uint8

type Tint_least8_t = int8

type Tuint_least8_t = uint8

type Tint_fast8_t = int8

type Tuint_fast8_t = uint8

type Tintmax_t = int64

type Tuintmax_t = uint64

type Ttimeval = struct {
	Ftv_sec  t__darwin_time_t
	Ftv_usec t__darwin_suseconds_t
}

type Trlim_t = uint64

type Trusage = struct {
	Fru_utime    Ttimeval
	Fru_stime    Ttimeval
	Fru_maxrss   int64
	Fru_ixrss    int64
	Fru_idrss    int64
	Fru_isrss    int64
	Fru_minflt   int64
	Fru_majflt   int64
	Fru_nswap    int64
	Fru_inblock  int64
	Fru_oublock  int64
	Fru_msgsnd   int64
	Fru_msgrcv   int64
	Fru_nsignals int64
	Fru_nvcsw    int64
	Fru_nivcsw   int64
}

type Trusage_info_t = uintptr

type Trusage_info_v0 = struct {
	Fri_uuid               [16]Tuint8_t
	Fri_user_time          Tuint64_t
	Fri_system_time        Tuint64_t
	Fri_pkg_idle_wkups     Tuint64_t
	Fri_interrupt_wkups    Tuint64_t
	Fri_pageins            Tuint64_t
	Fri_wired_size         Tuint64_t
	Fri_resident_size      Tuint64_t
	Fri_phys_footprint     Tuint64_t
	Fri_proc_start_abstime Tuint64_t
	Fri_proc_exit_abstime  Tuint64_t
}

type Trusage_info_v1 = struct {
	Fri_uuid                  [16]Tuint8_t
	Fri_user_time             Tuint64_t
	Fri_system_time           Tuint64_t
	Fri_pkg_idle_wkups        Tuint64_t
	Fri_interrupt_wkups       Tuint64_t
	Fri_pageins               Tuint64_t
	Fri_wired_size            Tuint64_t
	Fri_resident_size         Tuint64_t
	Fri_phys_footprint        Tuint64_t
	Fri_proc_start_abstime    Tuint64_t
	Fri_proc_exit_abstime     Tuint64_t
	Fri_child_user_time       Tuint64_t
	Fri_child_system_time     Tuint64_t
	Fri_child_pkg_idle_wkups  Tuint64_t
	Fri_child_interrupt_wkups Tuint64_t
	Fri_child_pageins         Tuint64_t
	Fri_child_elapsed_abstime Tuint64_t
}

type Trusage_info_v2 = struct {
	Fri_uuid                  [16]Tuint8_t
	Fri_user_time             Tuint64_t
	Fri_system_time           Tuint64_t
	Fri_pkg_idle_wkups        Tuint64_t
	Fri_interrupt_wkups       Tuint64_t
	Fri_pageins               Tuint64_t
	Fri_wired_size            Tuint64_t
	Fri_resident_size         Tuint64_t
	Fri_phys_footprint        Tuint64_t
	Fri_proc_start_abstime    Tuint64_t
	Fri_proc_exit_abstime     Tuint64_t
	Fri_child_user_time       Tuint64_t
	Fri_child_system_time     Tuint64_t
	Fri_child_pkg_idle_wkups  Tuint64_t
	Fri_child_interrupt_wkups Tuint64_t
	Fri_child_pageins         Tuint64_t
	Fri_child_elapsed_abstime Tuint64_t
	Fri_diskio_bytesread      Tuint64_t
	Fri_diskio_byteswritten   Tuint64_t
}

type Trusage_info_v3 = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
}

type Trusage_info_v4 = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
	Fri_logical_writes                Tuint64_t
	Fri_lifetime_max_phys_footprint   Tuint64_t
	Fri_instructions                  Tuint64_t
	Fri_cycles                        Tuint64_t
	Fri_billed_energy                 Tuint64_t
	Fri_serviced_energy               Tuint64_t
	Fri_interval_max_phys_footprint   Tuint64_t
	Fri_runnable_time                 Tuint64_t
}

type Trusage_info_v5 = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
	Fri_logical_writes                Tuint64_t
	Fri_lifetime_max_phys_footprint   Tuint64_t
	Fri_instructions                  Tuint64_t
	Fri_cycles                        Tuint64_t
	Fri_billed_energy                 Tuint64_t
	Fri_serviced_energy               Tuint64_t
	Fri_interval_max_phys_footprint   Tuint64_t
	Fri_runnable_time                 Tuint64_t
	Fri_flags                         Tuint64_t
}

type Trusage_info_v6 = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
	Fri_logical_writes                Tuint64_t
	Fri_lifetime_max_phys_footprint   Tuint64_t
	Fri_instructions                  Tuint64_t
	Fri_cycles                        Tuint64_t
	Fri_billed_energy                 Tuint64_t
	Fri_serviced_energy               Tuint64_t
	Fri_interval_max_phys_footprint   Tuint64_t
	Fri_runnable_time                 Tuint64_t
	Fri_flags                         Tuint64_t
	Fri_user_ptime                    Tuint64_t
	Fri_system_ptime                  Tuint64_t
	Fri_pinstructions                 Tuint64_t
	Fri_pcycles                       Tuint64_t
	Fri_energy_nj                     Tuint64_t
	Fri_penergy_nj                    Tuint64_t
	Fri_secure_time_in_system         Tuint64_t
	Fri_secure_ptime_in_system        Tuint64_t
	Fri_neural_footprint              Tuint64_t
	Fri_lifetime_max_neural_footprint Tuint64_t
	Fri_interval_max_neural_footprint Tuint64_t
	Fri_reserved                      [9]Tuint64_t
}

type Trusage_info_current = struct {
	Fri_uuid                          [16]Tuint8_t
	Fri_user_time                     Tuint64_t
	Fri_system_time                   Tuint64_t
	Fri_pkg_idle_wkups                Tuint64_t
	Fri_interrupt_wkups               Tuint64_t
	Fri_pageins                       Tuint64_t
	Fri_wired_size                    Tuint64_t
	Fri_resident_size                 Tuint64_t
	Fri_phys_footprint                Tuint64_t
	Fri_proc_start_abstime            Tuint64_t
	Fri_proc_exit_abstime             Tuint64_t
	Fri_child_user_time               Tuint64_t
	Fri_child_system_time             Tuint64_t
	Fri_child_pkg_idle_wkups          Tuint64_t
	Fri_child_interrupt_wkups         Tuint64_t
	Fri_child_pageins                 Tuint64_t
	Fri_child_elapsed_abstime         Tuint64_t
	Fri_diskio_bytesread              Tuint64_t
	Fri_diskio_byteswritten           Tuint64_t
	Fri_cpu_time_qos_default          Tuint64_t
	Fri_cpu_time_qos_maintenance      Tuint64_t
	Fri_cpu_time_qos_background       Tuint64_t
	Fri_cpu_time_qos_utility          Tuint64_t
	Fri_cpu_time_qos_legacy           Tuint64_t
	Fri_cpu_time_qos_user_initiated   Tuint64_t
	Fri_cpu_time_qos_user_interactive Tuint64_t
	Fri_billed_system_time            Tuint64_t
	Fri_serviced_system_time          Tuint64_t
	Fri_logical_writes                Tuint64_t
	Fri_lifetime_max_phys_footprint   Tuint64_t
	Fri_instructions                  Tuint64_t
	Fri_cycles                        Tuint64_t
	Fri_billed_energy                 Tuint64_t
	Fri_serviced_energy               Tuint64_t
	Fri_interval_max_phys_footprint   Tuint64_t
	Fri_runnable_time                 Tuint64_t
	Fri_flags                         Tuint64_t
	Fri_user_ptime                    Tuint64_t
	Fri_system_ptime                  Tuint64_t
	Fri_pinstructions                 Tuint64_t
	Fri_pcycles                       Tuint64_t
	Fri_energy_nj                     Tuint64_t
	Fri_penergy_nj                    Tuint64_t
	Fri_secure_time_in_system         Tuint64_t
	Fri_secure_ptime_in_system        Tuint64_t
	Fri_neural_footprint              Tuint64_t
	Fri_lifetime_max_neural_footprint Tuint64_t
	Fri_interval_max_neural_footprint Tuint64_t
	Fri_reserved                      [9]Tuint64_t
}

type Trlimit = struct {
	Frlim_cur Trlim_t
	Frlim_max Trlim_t
}

type Tproc_rlimit_control_wakeupmon = struct {
	Fwm_flags Tuint32_t
	Fwm_rate  Tint32_t
}

type Twait = struct {
	Fw_T [0]struct {
		F__ccgo0 uint32
	}
	Fw_S [0]struct {
		F__ccgo0 uint32
	}
	Fw_status int32
}

type Tct_rune_t = int32

type Trune_t = int32

type Twchar_t = int32

type Tdiv_t = struct {
	Fquot int32
	Frem  int32
}

type Tldiv_t = struct {
	Fquot int64
	Frem  int64
}

type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

type Tmalloc_type_id_t = uint64

type Tdev_t = int32

type Tmode_t = uint16 /* getsubopt(3) external variable */
/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

func XXauDisposeAuth(tls *libc.TLS, auth uintptr) {
	if auth != 0 {
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Faddress)
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fnumber)
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fname)
		if (*TXauth)(unsafe.Pointer(auth)).Fdata != 0 {
			libc.X__builtin___memset_chk(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata, 0, uint64((*TXauth)(unsafe.Pointer(auth)).Fdata_length), ^t__predefined_size_t(0))
			libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata)
		}
		libc.Xfree(tls, auth)
	}
	return
}

const m_ACCESSX_MAX_DESCRIPTORS = 100
const m_ATTRIBUTION_NAME_MAX = 255
const m_AT_EACCESS = 0x0010
const m_AT_FDONLY = 0x0400
const m_AT_REALDEV = 0x0200
const m_AT_REMOVEDIR = 0x0080
const m_AT_SYMLINK_FOLLOW = 0x0040
const m_AT_SYMLINK_NOFOLLOW = 0x0020
const m_AT_SYMLINK_NOFOLLOW_ANY = 0x0800
const m_BC_BASE_MAX = 99
const m_BC_DIM_MAX = 2048
const m_BC_SCALE_MAX = 99
const m_BC_STRING_MAX = 1000
const m_CHARCLASS_NAME_MAX = 14
const m_CHILD_MAX = 266
const m_CLOCK_MONOTONIC = "_CLOCK_MONOTONIC"
const m_CLOCK_MONOTONIC_RAW = "_CLOCK_MONOTONIC_RAW"
const m_CLOCK_MONOTONIC_RAW_APPROX = "_CLOCK_MONOTONIC_RAW_APPROX"
const m_CLOCK_PROCESS_CPUTIME_ID = "_CLOCK_PROCESS_CPUTIME_ID"
const m_CLOCK_REALTIME = "_CLOCK_REALTIME"
const m_CLOCK_THREAD_CPUTIME_ID = "_CLOCK_THREAD_CPUTIME_ID"
const m_CLOCK_UPTIME_RAW = "_CLOCK_UPTIME_RAW"
const m_CLOCK_UPTIME_RAW_APPROX = "_CLOCK_UPTIME_RAW_APPROX"
const m_COLL_WEIGHTS_MAX = 2
const m_CPF_IGNORE_MODE = 0x0002
const m_CPF_OVERWRITE = 0x0001
const m_DST_AUST = 2
const m_DST_CAN = 6
const m_DST_EET = 5
const m_DST_MET = 4
const m_DST_NONE = 0
const m_DST_USA = 1
const m_DST_WET = 3
const m_EQUIV_CLASS_MAX = 2
const m_EXPR_NEST_MAX = 32
const m_FAPPEND = "O_APPEND"
const m_FASYNC = "O_ASYNC"
const m_FCNTL_FS_SPECIFIC_BASE = 0x00010000
const m_FD_CLOEXEC = 1
const m_FD_SETSIZE = "__DARWIN_FD_SETSIZE"
const m_FFDSYNC = "O_DSYNC"
const m_FFSYNC = "O_FSYNC"
const m_FILESEC_GUID = "FILESEC_UUID"
const m_FNDELAY = "O_NONBLOCK"
const m_FNONBLOCK = "O_NONBLOCK"
const m_FREAD = 0x00000001
const m_FWRITE = 0x00000002
const m_F_ADDFILESIGS = 61
const m_F_ADDFILESIGS_FOR_DYLD_SIM = 83
const m_F_ADDFILESIGS_INFO = 103
const m_F_ADDFILESIGS_RETURN = 97
const m_F_ADDFILESUPPL = 104
const m_F_ADDSIGS = 59
const m_F_ADDSIGS_MAIN_BINARY = 113
const m_F_ALLOCATEALL = 0x00000004
const m_F_ALLOCATECONTIG = 0x00000002
const m_F_ALLOCATEPERSIST = 0x00000008
const m_F_ATTRIBUTION_TAG = 111
const m_F_BARRIERFSYNC = 85
const m_F_CHECK_LV = 98
const m_F_CHKCLEAN = 41
const m_F_CREATE_TAG = 0x00000001
const m_F_DELETE_TAG = 0x00000002
const m_F_DUPFD = 0
const m_F_DUPFD_CLOEXEC = 67
const m_F_FINDSIGS = 78
const m_F_FLUSH_DATA = 40
const m_F_FREEZE_FS = 53
const m_F_FULLFSYNC = 51
const m_F_GETCODEDIR = 72
const m_F_GETFD = 1
const m_F_GETFL = 3
const m_F_GETLEASE = 107
const m_F_GETLK = 7
const m_F_GETLKPID = 66
const m_F_GETNOSIGPIPE = 74
const m_F_GETOWN = 5
const m_F_GETPATH = 50
const m_F_GETPATH_MTMINFO = 71
const m_F_GETPATH_NOFIRMLINK = 102
const m_F_GETPROTECTIONCLASS = 63
const m_F_GETPROTECTIONLEVEL = 77
const m_F_GETSIGSINFO = 105
const m_F_GLOBAL_NOCACHE = 55
const m_F_LOCK = 1
const m_F_LOG2PHYS = 49
const m_F_LOG2PHYS_EXT = 65
const m_F_NOCACHE = 48
const m_F_NODIRECT = 62
const m_F_OFD_GETLK = 92
const m_F_OFD_SETLK = 90
const m_F_OFD_SETLKW = 91
const m_F_OFD_SETLKWTIMEOUT = 93
const m_F_OK = 0
const m_F_PATHPKG_CHECK = 52
const m_F_PEOFPOSMODE = 3
const m_F_PREALLOCATE = 42
const m_F_PUNCHHOLE = 99
const m_F_QUERY_TAG = 0x00000004
const m_F_RDADVISE = 44
const m_F_RDAHEAD = 45
const m_F_RDLCK = 1
const m_F_SETBACKINGSTORE = 70
const m_F_SETFD = 2
const m_F_SETFL = 4
const m_F_SETLEASE = 106
const m_F_SETLK = 8
const m_F_SETLKW = 9
const m_F_SETLKWTIMEOUT = 10
const m_F_SETNOSIGPIPE = 73
const m_F_SETOWN = 6
const m_F_SETPROTECTIONCLASS = 64
const m_F_SETSIZE = 43
const m_F_SINGLE_WRITER = 76
const m_F_SPECULATIVE_READ = 101
const m_F_TEST = 3
const m_F_THAW_FS = 54
const m_F_TLOCK = 2
const m_F_TRANSCODEKEY = 75
const m_F_TRANSFEREXTENTS = 110
const m_F_TRIM_ACTIVE_FILE = 100
const m_F_ULOCK = 0
const m_F_UNLCK = 2
const m_F_VOLPOSMODE = 4
const m_F_WRLCK = 3
const m_GETSIGSINFO_PLATFORM_BINARY = 1
const m_GID_MAX = 2147483647
const m_ITIMER_PROF = 2
const m_ITIMER_REAL = 0
const m_ITIMER_VIRTUAL = 1
const m_LINE_MAX = 2048
const m_LINK_MAX = 32767
const m_LOCK_EX = 0x02
const m_LOCK_NB = 0x04
const m_LOCK_SH = 0x01
const m_LOCK_UN = 0x08
const m_L_INCR = "SEEK_CUR"
const m_L_SET = "SEEK_SET"
const m_L_XTND = "SEEK_END"
const m_MAX_CANON = 1024
const m_MAX_INPUT = 1024
const m_NAME_MAX = 255
const m_NBBY = "__DARWIN_NBBY"
const m_NFDBITS = "__DARWIN_NFDBITS"
const m_NGROUPS_MAX = 16
const m_NZERO = 20
const m_OPEN_MAX = 10240
const m_O_ACCMODE = 0x0003
const m_O_ALERT = 0x20000000
const m_O_APPEND = 0x00000008
const m_O_ASYNC = 0x00000040
const m_O_CLOEXEC = 0x01000000
const m_O_CREAT = 0x00000200
const m_O_DIRECTORY = 0x00100000
const m_O_DP_AUTHENTICATE = 0x0004
const m_O_DP_GETRAWENCRYPTED = 0x0001
const m_O_DP_GETRAWUNENCRYPTED = 0x0002
const m_O_DSYNC = 0x400000
const m_O_EVTONLY = 0x00008000
const m_O_EXCL = 0x00000800
const m_O_EXEC = 0x40000000
const m_O_EXLOCK = 0x00000020
const m_O_FSYNC = "O_SYNC"
const m_O_NDELAY = "O_NONBLOCK"
const m_O_NOCTTY = 0x00020000
const m_O_NOFOLLOW = 0x00000100
const m_O_NOFOLLOW_ANY = 0x20000000
const m_O_NONBLOCK = 0x00000004
const m_O_POPUP = 0x80000000
const m_O_RDONLY = 0x0000
const m_O_RDWR = 0x0002
const m_O_RESOLVE_BENEATH = 0x00001000
const m_O_SHLOCK = 0x00000010
const m_O_SYMLINK = 0x00200000
const m_O_SYNC = 0x0080
const m_O_TRUNC = 0x00000400
const m_O_WRONLY = 0x0001
const m_PATH_MAX = 1024
const m_PIPE_BUF = 512
const m_RE_DUP_MAX = 255
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_SYNC_VOLUME_FULLSYNC = 0x01
const m_SYNC_VOLUME_WAIT = 0x02
const m_S_IEXEC = "S_IXUSR"
const m_S_IFBLK = 0060000
const m_S_IFCHR = 0020000
const m_S_IFDIR = 0040000
const m_S_IFIFO = 0010000
const m_S_IFLNK = 0120000
const m_S_IFMT = 0170000
const m_S_IFREG = 0100000
const m_S_IFSOCK = 0140000
const m_S_IFWHT = 0160000
const m_S_IREAD = "S_IRUSR"
const m_S_IRGRP = 0000040
const m_S_IROTH = 0000004
const m_S_IRUSR = 0000400
const m_S_IRWXG = 0000070
const m_S_IRWXO = 0000007
const m_S_IRWXU = 0000700
const m_S_ISGID = 0002000
const m_S_ISTXT = "S_ISVTX"
const m_S_ISUID = 0004000
const m_S_ISVTX = 0001000
const m_S_IWGRP = 0000020
const m_S_IWOTH = 0000002
const m_S_IWRITE = "S_IWUSR"
const m_S_IWUSR = 0000200
const m_S_IXGRP = 0000010
const m_S_IXOTH = 0000001
const m_S_IXUSR = 0000100
const m_TIME_UTC = 1
const m_UID_MAX = 2147483647
const m_USER_FSIGNATURES_CDHASH_LEN = 20
const m_X_BIG_ENDIAN = "BIG_ENDIAN"
const m_X_BYTE_ORDER = "BYTE_ORDER"
const m_X_LITTLE_ENDIAN = "LITTLE_ENDIAN"
const m__CS_DARWIN_USER_CACHE_DIR = 65538
const m__CS_DARWIN_USER_DIR = 65536
const m__CS_DARWIN_USER_TEMP_DIR = 65537
const m__CS_PATH = 1
const m__CS_POSIX_V6_ILP32_OFF32_CFLAGS = 2
const m__CS_POSIX_V6_ILP32_OFF32_LDFLAGS = 3
const m__CS_POSIX_V6_ILP32_OFF32_LIBS = 4
const m__CS_POSIX_V6_ILP32_OFFBIG_CFLAGS = 5
const m__CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS = 6
const m__CS_POSIX_V6_ILP32_OFFBIG_LIBS = 7
const m__CS_POSIX_V6_LP64_OFF64_CFLAGS = 8
const m__CS_POSIX_V6_LP64_OFF64_LDFLAGS = 9
const m__CS_POSIX_V6_LP64_OFF64_LIBS = 10
const m__CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS = 11
const m__CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS = 12
const m__CS_POSIX_V6_LPBIG_OFFBIG_LIBS = 13
const m__CS_POSIX_V6_WIDTH_RESTRICTED_ENVS = 14
const m__CS_XBS5_ILP32_OFF32_CFLAGS = 20
const m__CS_XBS5_ILP32_OFF32_LDFLAGS = 21
const m__CS_XBS5_ILP32_OFF32_LIBS = 22
const m__CS_XBS5_ILP32_OFF32_LINTFLAGS = 23
const m__CS_XBS5_ILP32_OFFBIG_CFLAGS = 24
const m__CS_XBS5_ILP32_OFFBIG_LDFLAGS = 25
const m__CS_XBS5_ILP32_OFFBIG_LIBS = 26
const m__CS_XBS5_ILP32_OFFBIG_LINTFLAGS = 27
const m__CS_XBS5_LP64_OFF64_CFLAGS = 28
const m__CS_XBS5_LP64_OFF64_LDFLAGS = 29
const m__CS_XBS5_LP64_OFF64_LIBS = 30
const m__CS_XBS5_LP64_OFF64_LINTFLAGS = 31
const m__CS_XBS5_LPBIG_OFFBIG_CFLAGS = 32
const m__CS_XBS5_LPBIG_OFFBIG_LDFLAGS = 33
const m__CS_XBS5_LPBIG_OFFBIG_LIBS = 34
const m__CS_XBS5_LPBIG_OFFBIG_LINTFLAGS = 35
const m__PC_2_SYMLINKS = 15
const m__PC_ALLOC_SIZE_MIN = 16
const m__PC_ASYNC_IO = 17
const m__PC_AUTH_OPAQUE_NP = 14
const m__PC_CASE_PRESERVING = 12
const m__PC_CASE_SENSITIVE = 11
const m__PC_CHOWN_RESTRICTED = 7
const m__PC_EXTENDED_SECURITY_NP = 13
const m__PC_FILESIZEBITS = 18
const m__PC_LINK_MAX = 1
const m__PC_MAX_CANON = 2
const m__PC_MAX_INPUT = 3
const m__PC_MIN_HOLE_SIZE = 27
const m__PC_NAME_CHARS_MAX = 10
const m__PC_NAME_MAX = 4
const m__PC_NO_TRUNC = 8
const m__PC_PATH_MAX = 5
const m__PC_PIPE_BUF = 6
const m__PC_PRIO_IO = 19
const m__PC_REC_INCR_XFER_SIZE = 20
const m__PC_REC_MAX_XFER_SIZE = 21
const m__PC_REC_MIN_XFER_SIZE = 22
const m__PC_REC_XFER_ALIGN = 23
const m__PC_SYMLINK_MAX = 24
const m__PC_SYNC_IO = 25
const m__PC_VDISABLE = 9
const m__PC_XATTR_SIZE_BITS = 26
const m__POSIX2_CHAR_TERM = 200112
const m__POSIX2_C_BIND = 200112
const m__POSIX2_C_DEV = 200112
const m__POSIX2_FORT_RUN = 200112
const m__POSIX2_LOCALEDEF = 200112
const m__POSIX2_SW_DEV = 200112
const m__POSIX2_UPE = 200112
const m__POSIX2_VERSION = 200112
const m__POSIX_CHOWN_RESTRICTED = 200112
const m__POSIX_FSYNC = 200112
const m__POSIX_IPV6 = 200112
const m__POSIX_JOB_CONTROL = 200112
const m__POSIX_MAPPED_FILES = 200112
const m__POSIX_MEMORY_PROTECTION = 200112
const m__POSIX_NO_TRUNC = 200112
const m__POSIX_READER_WRITER_LOCKS = 200112
const m__POSIX_REGEXP = 200112
const m__POSIX_SAVED_IDS = 200112
const m__POSIX_SHELL = 200112
const m__POSIX_SPAWN = 200112
const m__POSIX_THREADS = 200112
const m__POSIX_THREAD_ATTR_STACKADDR = 200112
const m__POSIX_THREAD_ATTR_STACKSIZE = 200112
const m__POSIX_THREAD_KEYS_MAX = 128
const m__POSIX_THREAD_PROCESS_SHARED = 200112
const m__POSIX_THREAD_SAFE_FUNCTIONS = 200112
const m__POSIX_V6_ILP32_OFF32 = "__ILP32_OFF32"
const m__POSIX_V6_ILP32_OFFBIG = "__ILP32_OFFBIG"
const m__POSIX_V6_LP64_OFF64 = "__LP64_OFF64"
const m__POSIX_V6_LPBIG_OFFBIG = "__LPBIG_OFFBIG"
const m__POSIX_V7_ILP32_OFF32 = "__ILP32_OFF32"
const m__POSIX_V7_ILP32_OFFBIG = "__ILP32_OFFBIG"
const m__POSIX_V7_LP64_OFF64 = "__LP64_OFF64"
const m__POSIX_V7_LPBIG_OFFBIG = "__LPBIG_OFFBIG"
const m__POSIX_VERSION = 200112
const m__SC_2_CHAR_TERM = 20
const m__SC_2_C_BIND = 18
const m__SC_2_C_DEV = 19
const m__SC_2_FORT_DEV = 21
const m__SC_2_FORT_RUN = 22
const m__SC_2_LOCALEDEF = 23
const m__SC_2_PBS = 59
const m__SC_2_PBS_ACCOUNTING = 60
const m__SC_2_PBS_CHECKPOINT = 61
const m__SC_2_PBS_LOCATE = 62
const m__SC_2_PBS_MESSAGE = 63
const m__SC_2_PBS_TRACK = 64
const m__SC_2_SW_DEV = 24
const m__SC_2_UPE = 25
const m__SC_2_VERSION = 17
const m__SC_ADVISORY_INFO = 65
const m__SC_AIO_LISTIO_MAX = 42
const m__SC_AIO_MAX = 43
const m__SC_AIO_PRIO_DELTA_MAX = 44
const m__SC_ARG_MAX = 1
const m__SC_ASYNCHRONOUS_IO = 28
const m__SC_ATEXIT_MAX = 107
const m__SC_BARRIERS = 66
const m__SC_BC_BASE_MAX = 9
const m__SC_BC_DIM_MAX = 10
const m__SC_BC_SCALE_MAX = 11
const m__SC_BC_STRING_MAX = 12
const m__SC_CHILD_MAX = 2
const m__SC_CLK_TCK = 3
const m__SC_CLOCK_SELECTION = 67
const m__SC_COLL_WEIGHTS_MAX = 13
const m__SC_CPUTIME = 68
const m__SC_DELAYTIMER_MAX = 45
const m__SC_EXPR_NEST_MAX = 14
const m__SC_FILE_LOCKING = 69
const m__SC_FSYNC = 38
const m__SC_GETGR_R_SIZE_MAX = 70
const m__SC_GETPW_R_SIZE_MAX = 71
const m__SC_HOST_NAME_MAX = 72
const m__SC_IOV_MAX = 56
const m__SC_IPV6 = 118
const m__SC_JOB_CONTROL = 6
const m__SC_LINE_MAX = 15
const m__SC_LOGIN_NAME_MAX = 73
const m__SC_MAPPED_FILES = 47
const m__SC_MEMLOCK = 30
const m__SC_MEMLOCK_RANGE = 31
const m__SC_MEMORY_PROTECTION = 32
const m__SC_MESSAGE_PASSING = 33
const m__SC_MONOTONIC_CLOCK = 74
const m__SC_MQ_OPEN_MAX = 46
const m__SC_MQ_PRIO_MAX = 75
const m__SC_NGROUPS_MAX = 4
const m__SC_NPROCESSORS_CONF = 57
const m__SC_NPROCESSORS_ONLN = 58
const m__SC_OPEN_MAX = 5
const m__SC_PAGESIZE = 29
const m__SC_PAGE_SIZE = "_SC_PAGESIZE"
const m__SC_PASS_MAX = 131
const m__SC_PHYS_PAGES = 200
const m__SC_PRIORITIZED_IO = 34
const m__SC_PRIORITY_SCHEDULING = 35
const m__SC_RAW_SOCKETS = 119
const m__SC_READER_WRITER_LOCKS = 76
const m__SC_REALTIME_SIGNALS = 36
const m__SC_REGEXP = 77
const m__SC_RE_DUP_MAX = 16
const m__SC_RTSIG_MAX = 48
const m__SC_SAVED_IDS = 7
const m__SC_SEMAPHORES = 37
const m__SC_SEM_NSEMS_MAX = 49
const m__SC_SEM_VALUE_MAX = 50
const m__SC_SHARED_MEMORY_OBJECTS = 39
const m__SC_SHELL = 78
const m__SC_SIGQUEUE_MAX = 51
const m__SC_SPAWN = 79
const m__SC_SPIN_LOCKS = 80
const m__SC_SPORADIC_SERVER = 81
const m__SC_SS_REPL_MAX = 126
const m__SC_STREAM_MAX = 26
const m__SC_SYMLOOP_MAX = 120
const m__SC_SYNCHRONIZED_IO = 40
const m__SC_THREADS = 96
const m__SC_THREAD_ATTR_STACKADDR = 82
const m__SC_THREAD_ATTR_STACKSIZE = 83
const m__SC_THREAD_CPUTIME = 84
const m__SC_THREAD_DESTRUCTOR_ITERATIONS = 85
const m__SC_THREAD_KEYS_MAX = 86
const m__SC_THREAD_PRIORITY_SCHEDULING = 89
const m__SC_THREAD_PRIO_INHERIT = 87
const m__SC_THREAD_PRIO_PROTECT = 88
const m__SC_THREAD_PROCESS_SHARED = 90
const m__SC_THREAD_SAFE_FUNCTIONS = 91
const m__SC_THREAD_SPORADIC_SERVER = 92
const m__SC_THREAD_STACK_MIN = 93
const m__SC_THREAD_THREADS_MAX = 94
const m__SC_TIMEOUTS = 95
const m__SC_TIMERS = 41
const m__SC_TIMER_MAX = 52
const m__SC_TRACE = 97
const m__SC_TRACE_EVENT_FILTER = 98
const m__SC_TRACE_EVENT_NAME_MAX = 127
const m__SC_TRACE_INHERIT = 99
const m__SC_TRACE_LOG = 100
const m__SC_TRACE_NAME_MAX = 128
const m__SC_TRACE_SYS_MAX = 129
const m__SC_TRACE_USER_EVENT_MAX = 130
const m__SC_TTY_NAME_MAX = 101
const m__SC_TYPED_MEMORY_OBJECTS = 102
const m__SC_TZNAME_MAX = 27
const m__SC_V6_ILP32_OFF32 = 103
const m__SC_V6_ILP32_OFFBIG = 104
const m__SC_V6_LP64_OFF64 = 105
const m__SC_V6_LPBIG_OFFBIG = 106
const m__SC_VERSION = 8
const m__SC_XBS5_ILP32_OFF32 = 122
const m__SC_XBS5_ILP32_OFFBIG = 123
const m__SC_XBS5_LP64_OFF64 = 124
const m__SC_XBS5_LPBIG_OFFBIG = 125
const m__SC_XOPEN_CRYPT = 108
const m__SC_XOPEN_ENH_I18N = 109
const m__SC_XOPEN_LEGACY = 110
const m__SC_XOPEN_REALTIME = 111
const m__SC_XOPEN_REALTIME_THREADS = 112
const m__SC_XOPEN_SHM = 113
const m__SC_XOPEN_STREAMS = 114
const m__SC_XOPEN_UNIX = 115
const m__SC_XOPEN_VERSION = 116
const m__SC_XOPEN_XCU_VERSION = 121
const m__V6_ILP32_OFF32 = "__ILP32_OFF32"
const m__V6_ILP32_OFFBIG = "__ILP32_OFFBIG"
const m__V6_LP64_OFF64 = "__LP64_OFF64"
const m__V6_LPBIG_OFFBIG = "__LPBIG_OFFBIG"
const m__XBS5_ILP32_OFF32 = "__ILP32_OFF32"
const m__XBS5_ILP32_OFFBIG = "__ILP32_OFFBIG"
const m__XBS5_LP64_OFF64 = "__LP64_OFF64"
const m__XBS5_LPBIG_OFFBIG = "__LPBIG_OFFBIG"
const m__XOPEN_CRYPT = 1
const m__XOPEN_ENH_I18N = 1
const m__XOPEN_SHM = 1
const m__XOPEN_UNIX = 1
const m__XOPEN_VERSION = 600
const m__XOPEN_XCU_VERSION = 4
const m___DARWIN_FD_SETSIZE = 1024
const m___DARWIN_NBBY = 8
const m___LP64_OFF64 = 1
const m___LPBIG_OFFBIG = 1
const m_static_assert = "_Static_assert"

type Tu_char = uint8

type Tu_short = uint16

type Tu_int = uint32

type Tu_long = uint64

type Tushort = uint16

type Tuint = uint32

type Tu_quad_t = uint64

type Tquad_t = int64

type Tqaddr_t = uintptr

type Tcaddr_t = uintptr

type Tdaddr_t = int32

type Tfixpt_t = uint32

type Tblkcnt_t = int64

type Tblksize_t = int32

type Tgid_t = uint32

type Tin_addr_t = uint32

type Tin_port_t = uint16

type Tino_t = uint64

type Tino64_t = uint64

type Tkey_t = int32

type Tnlink_t = uint16

type Tsegsz_t = int32

type Tswblk_t = int32

type Tclock_t = uint64

type Ttime_t = int64

type Tuseconds_t = uint32

type Tsuseconds_t = int32

type Tfd_set = struct {
	Ffds_bits [32]t__int32_t
}

type Tfd_mask = int32

type Tpthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type Tpthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type Tpthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type Tpthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type Tpthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type Tpthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type Tpthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type Tpthread_t = uintptr

type Tpthread_key_t = uint64

type Tfsblkcnt_t = uint32

type Tfsfilcnt_t = uint32

type Tflock = struct {
	Fl_start  Toff_t
	Fl_len    Toff_t
	Fl_pid    Tpid_t
	Fl_type   int16
	Fl_whence int16
}

type Ttimespec = struct {
	Ftv_sec  t__darwin_time_t
	Ftv_nsec int64
}

type Tflocktimeout = struct {
	Ffl      Tflock
	Ftimeout Ttimespec
}

type Tradvisory = struct {
	Fra_offset Toff_t
	Fra_count  int32
}

type Tfsignatures_t = struct {
	Ffs_file_start       Toff_t
	Ffs_blob_start       uintptr
	Ffs_blob_size        Tsize_t
	Ffs_fsignatures_size Tsize_t
	Ffs_cdhash           [20]int8
	Ffs_hash_type        int32
}

type Tfsignatures = Tfsignatures_t

type Tfsupplement_t = struct {
	Ffs_file_start Toff_t
	Ffs_blob_start Toff_t
	Ffs_blob_size  Tsize_t
	Ffs_orig_fd    int32
}

type Tfsupplement = Tfsupplement_t

type Tfchecklv_t = struct {
	Flv_file_start         Toff_t
	Flv_error_message_size Tsize_t
	Flv_error_message      uintptr
}

type Tfchecklv = Tfchecklv_t

type Tfgetsigsinfo_t = struct {
	Ffg_file_start      Toff_t
	Ffg_info_request    int32
	Ffg_sig_is_platform int32
}

type Tfgetsigsinfo = Tfgetsigsinfo_t

type Tfstore_t = struct {
	Ffst_flags      uint32
	Ffst_posmode    int32
	Ffst_offset     Toff_t
	Ffst_length     Toff_t
	Ffst_bytesalloc Toff_t
}

type Tfstore = Tfstore_t

type Tfpunchhole_t = struct {
	Ffp_flags  uint32
	Freserved  uint32
	Ffp_offset Toff_t
	Ffp_length Toff_t
}

type Tfpunchhole = Tfpunchhole_t

type Tftrimactivefile_t = struct {
	Ffta_offset Toff_t
	Ffta_length Toff_t
}

type Tftrimactivefile = Tftrimactivefile_t

type Tfspecread_t = struct {
	Ffsr_flags  uint32
	Freserved   uint32
	Ffsr_offset Toff_t
	Ffsr_length Toff_t
}

type Tfspecread = Tfspecread_t

type Tfattributiontag_t = struct {
	Fft_flags            uint32
	Fft_hash             uint64
	Fft_attribution_name [255]int8
}

type Tfattributiontag = Tfattributiontag_t

type Tlog2phys = struct {
	Fl2p_flags       uint32
	Fl2p_contigbytes Toff_t
	Fl2p_devoffset   Toff_t
}

type Tfilesec_t = uintptr

type Tfilesec_property_t = int32

const _FILESEC_OWNER = 1
const _FILESEC_GROUP = 2
const _FILESEC_UUID = 3
const _FILESEC_MODE = 4
const _FILESEC_ACL = 5
const _FILESEC_GRPUUID = 6
const _FILESEC_ACL_RAW = 100
const _FILESEC_ACL_ALLOCSIZE = 101

type Taccessx_descriptor = struct {
	Fad_name_offset uint32
	Fad_flags       int32
	Fad_pad         [2]int32
}

type Tuuid_t = [16]uint8

type Ttimeval64 = struct {
	Ftv_sec  t__int64_t
	Ftv_usec t__int64_t
}

type Titimerval = struct {
	Fit_interval Ttimeval
	Fit_value    Ttimeval
}

type Ttimezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Tclockinfo = struct {
	Fhz      int32
	Ftick    int32
	Ftickadj int32
	Fstathz  int32
	Fprofhz  int32
}

type Ttm = struct {
	Ftm_sec    int32
	Ftm_min    int32
	Ftm_hour   int32
	Ftm_mday   int32
	Ftm_mon    int32
	Ftm_year   int32
	Ftm_wday   int32
	Ftm_yday   int32
	Ftm_isdst  int32
	Ftm_gmtoff int64
	Ftm_zone   uintptr
} //TODO "timezone" // _time.h:106:13:

type Tclockid_t = int32

const __CLOCK_REALTIME = 0
const __CLOCK_MONOTONIC = 6
const __CLOCK_MONOTONIC_RAW = 4
const __CLOCK_MONOTONIC_RAW_APPROX = 5
const __CLOCK_UPTIME_RAW = 8
const __CLOCK_UPTIME_RAW_APPROX = 9
const __CLOCK_PROCESS_CPUTIME_ID = 12
const __CLOCK_THREAD_CPUTIME_ID = 16 /* getsubopt(3) external variable */
/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

var _buf = libc.UintptrFromInt32(0)

func _free_filename_buffer(tls *libc.TLS) {
	libc.Xfree(tls, _buf)
	_buf = libc.UintptrFromInt32(0)
}

func XXauFileName(tls *libc.TLS) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var name, slashDotXauthority, v1 uintptr
	var size Tsize_t
	var v2 int32
	_, _, _, _, _ = name, size, slashDotXauthority, v1, v2
	slashDotXauthority = __ccgo_ts
	v1 = libc.Xgetenv(tls, __ccgo_ts+13)
	name = v1
	if v1 != 0 {
		return name
	}
	name = libc.Xgetenv(tls, __ccgo_ts+24)
	if !(name != 0) {
		return libc.UintptrFromInt32(0)
	}
	size = libc.Xstrlen(tls, name) + libc.Xstrlen(tls, slashDotXauthority+1) + uint64(2)
	if size > _bsize || _buf == libc.UintptrFromInt32(0) {
		libc.Xfree(tls, _buf)
		_buf = libc.Xmalloc(tls, size)
		if !(_buf != 0) {
			_bsize = uint64(0)
			return libc.UintptrFromInt32(0)
		}
		if !(_atexit_registered != 0) {
			libc.Xatexit(tls, __ccgo_fp(_free_filename_buffer))
			_atexit_registered = int32(1)
		}
		_bsize = size
	}
	if int32(*(*int8)(unsafe.Pointer(name))) == int32('/') && int32(*(*int8)(unsafe.Pointer(name + 1))) == int32('\000') {
		v2 = int32(1)
	} else {
		v2 = 0
	}
	libc.X__builtin___snprintf_chk(tls, _buf, _bsize, 0, ^t__predefined_size_t(0), __ccgo_ts+29, libc.VaList(bp+8, name, slashDotXauthority+uintptr(v2)))
	return _buf
}

var _bsize Tsize_t

var _atexit_registered int32

/*
 * Copyright (c) 2000, 2023 Apple Computer, Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*
 * Copyright (c) 1989, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
 * Copyright (c) 2023 Apple Computer, Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*
 * Copyright (c) 1989, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/* define X_GETTIMEOFDAY macro, a portable gettimeofday() */

/* use POSIX name for signal */

/*
 * Copyright 1997 Metro Link Incorporated
 *
 *                           All Rights Reserved
 *
 * Permission to use, copy, modify, distribute, and sell this software and its
 * documentation for any purpose is hereby granted without fee, provided that
 * the above copyright notice appear in all copies and that both that
 * copyright notice and this permission notice appear in supporting
 * documentation, and that the names of the above listed copyright holder(s)
 * not be used in advertising or publicity pertaining to distribution of
 * the software without specific, written prior permission.  The above listed
 * copyright holder(s) make(s) no representations about the suitability of
 * this software for any purpose.  It is provided "as is" without express or
 * implied warranty.
 *
 * THE ABOVE LISTED COPYRIGHT HOLDER(S) DISCLAIM(S) ALL WARRANTIES WITH REGARD
 * TO THIS SOFTWARE, INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS, IN NO EVENT SHALL THE ABOVE LISTED COPYRIGHT HOLDER(S) BE
 * LIABLE FOR ANY SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR ANY
 * DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER
 * IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING
 * OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
 * Determine the machine's byte order.
 */

/* See if it is set in the imake config first */

/*
 * Copyright (c) 2000-2007 Apple Inc. All rights reserved.
 *
 * @APPLE_OSREFERENCE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. The rights granted to you under the License
 * may not be used to create, or enable the creation or redistribution of,
 * unlawful or unlicensed copies of an Apple operating system, or to
 * circumvent, violate, or enable the circumvention or violation of, any
 * terms of an Apple operating system software license agreement.
 *
 * Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_OSREFERENCE_LICENSE_HEADER_END@
 */
/*
 * Copyright 1995 NeXT Computer, Inc. All rights reserved.
 */

func XXauGetAuthByAddr(tls *libc.TLS, family uint32, address_length uint32, address uintptr, number_length uint32, number uintptr, name_length uint32, name uintptr) (r uintptr) {
	var auth_file, auth_name, entry uintptr
	_, _, _ = auth_file, auth_name, entry
	auth_name = XXauFileName(tls)
	if !(auth_name != 0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xaccess(tls, auth_name, libc.Int32FromInt32(1)<<libc.Int32FromInt32(2)) != 0 { /* checks REAL id */
		return libc.UintptrFromInt32(0)
	}
	auth_file = libc.Xfopen(tls, auth_name, __ccgo_ts+34)
	if !(auth_file != 0) {
		return libc.UintptrFromInt32(0)
	}
	for {
		entry = XXauReadAuth(tls, auth_file)
		if !(entry != 0) {
			break
		}
		/*
		 * Match when:
		 *   either family or entry->family are FamilyWild or
		 *    family and entry->family are the same and
		 *     address and entry->address are the same
		 *  and
		 *   either number or entry->number are empty or
		 *    number and entry->number are the same
		 *  and
		 *   either name or entry->name are empty or
		 *    name and entry->name are the same
		 */
		if (family == libc.Uint32FromInt32(libc.Int32FromInt32(m_FamilyWild)) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Ffamily) == int32(m_FamilyWild) || uint32((*TXauth)(unsafe.Pointer(entry)).Ffamily) == family && address_length == uint32((*TXauth)(unsafe.Pointer(entry)).Faddress_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Faddress, address, uint64(address_length)) == 0) && (number_length == uint32(0) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) == 0 || number_length == uint32((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Fnumber, number, uint64(number_length)) == 0) && (name_length == uint32(0) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fname_length) == 0 || uint32((*TXauth)(unsafe.Pointer(entry)).Fname_length) == name_length && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Fname, name, uint64(name_length)) == 0) {
			break
		}
		XXauDisposeAuth(tls, entry)
		goto _1
	_1:
	}
	libc.Xfclose(tls, auth_file)
	return entry
}

const m_PTHREAD_CANCEL_ASYNCHRONOUS = 0x00
const m_PTHREAD_CANCEL_DEFERRED = 0x02
const m_PTHREAD_CANCEL_DISABLE = 0x00
const m_PTHREAD_CANCEL_ENABLE = 0x01
const m_PTHREAD_CREATE_DETACHED = 2
const m_PTHREAD_CREATE_JOINABLE = 1
const m_PTHREAD_EXPLICIT_SCHED = 2
const m_PTHREAD_INHERIT_SCHED = 1
const m_PTHREAD_MUTEX_DEFAULT = "PTHREAD_MUTEX_NORMAL"
const m_PTHREAD_MUTEX_ERRORCHECK = 1
const m_PTHREAD_MUTEX_NORMAL = 0
const m_PTHREAD_MUTEX_POLICY_FAIRSHARE_NP = 1
const m_PTHREAD_MUTEX_POLICY_FIRSTFIT_NP = 3
const m_PTHREAD_MUTEX_RECURSIVE = 2
const m_PTHREAD_PRIO_INHERIT = 1
const m_PTHREAD_PRIO_NONE = 0
const m_PTHREAD_PRIO_PROTECT = 2
const m_PTHREAD_PROCESS_PRIVATE = 2
const m_PTHREAD_PROCESS_SHARED = 1
const m_PTHREAD_SCOPE_PROCESS = 2
const m_PTHREAD_SCOPE_SYSTEM = 1
const m_SCHED_FIFO = 4
const m_SCHED_OTHER = 1
const m_SCHED_RR = 2
const m_XMUTEX_INITIALIZER = "PTHREAD_MUTEX_INITIALIZER"
const m__PTHREAD_COND_SIG_init = 0x3CB0B1BB
const m__PTHREAD_ERRORCHECK_MUTEX_SIG_init = 0x32AAABA1
const m__PTHREAD_FIRSTFIT_MUTEX_SIG_init = 0x32AAABA3
const m__PTHREAD_MUTEX_SIG_init = 0x32AAABA7
const m__PTHREAD_ONCE_SIG_init = 0x30B1BCBA
const m__PTHREAD_RECURSIVE_MUTEX_SIG_init = 0x32AAABA2
const m__PTHREAD_RWLOCK_SIG_init = 0x2DA8B3B4
const m___SCHED_PARAM_SIZE__ = 4
const m_xfree = "free"
const m_xmalloc = "malloc"
const m_xthread_self = "pthread_self"

type Tsched_param = struct {
	Fsched_priority int32
	F__opaque       [4]int8
}

const _QOS_CLASS_USER_INTERACTIVE = 33
const _QOS_CLASS_USER_INITIATED = 25
const _QOS_CLASS_DEFAULT = 21
const _QOS_CLASS_UTILITY = 17
const _QOS_CLASS_BACKGROUND = 9
const _QOS_CLASS_UNSPECIFIED = 0

type Tqos_class_t = uint32

type Tpthread_override_t = uintptr

type Tmach_port_t = uint32

type Tpthread_jit_write_callback_t = uintptr

type Txthread_t = uintptr

type Txthread_key_t = uint64

type Txcondition_rec = struct {
	F__sig    int64
	F__opaque [40]int8
}

type Txmutex_rec = struct {
	F__sig    int64
	F__opaque [56]int8
}

type Txcondition_t = uintptr

type Txmutex_t = uintptr

/* aids understood by some debuggers */

func XXauGetBestAuthByAddr(tls *libc.TLS, family uint32, address_length uint32, address uintptr, number_length uint32, number uintptr, types_length int32, types uintptr, type_lengths uintptr) (r uintptr) {
	var auth_file, auth_name, best, entry uintptr
	var best_type, type1 int32
	_, _, _, _, _, _ = auth_file, auth_name, best, best_type, entry, type1
	auth_name = XXauFileName(tls)
	if !(auth_name != 0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xaccess(tls, auth_name, libc.Int32FromInt32(1)<<libc.Int32FromInt32(2)) != 0 { /* checks REAL id */
		return libc.UintptrFromInt32(0)
	}
	auth_file = libc.Xfopen(tls, auth_name, __ccgo_ts+34)
	if !(auth_file != 0) {
		return libc.UintptrFromInt32(0)
	}
	best = libc.UintptrFromInt32(0)
	best_type = types_length
	for {
		entry = XXauReadAuth(tls, auth_file)
		if !(entry != 0) {
			break
		}
		/*
		 * Match when:
		 *   either family or entry->family are FamilyWild or
		 *    family and entry->family are the same and
		 *     address and entry->address are the same
		 *  and
		 *   either number or entry->number are empty or
		 *    number and entry->number are the same
		 *  and
		 *   either name or entry->name are empty or
		 *    name and entry->name are the same
		 */
		if (family == libc.Uint32FromInt32(libc.Int32FromInt32(m_FamilyWild)) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Ffamily) == int32(m_FamilyWild) || uint32((*TXauth)(unsafe.Pointer(entry)).Ffamily) == family && (address_length == uint32((*TXauth)(unsafe.Pointer(entry)).Faddress_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Faddress, address, uint64(address_length)) == 0)) && (number_length == uint32(0) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) == 0 || number_length == uint32((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Fnumber, number, uint64(number_length)) == 0) {
			if best_type == 0 {
				best = entry
				break
			}
			type1 = 0
			for {
				if !(type1 < best_type) {
					break
				}
				if *(*int32)(unsafe.Pointer(type_lengths + uintptr(type1)*4)) == libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fname_length) && !(libc.Xstrncmp(tls, *(*uintptr)(unsafe.Pointer(types + uintptr(type1)*8)), (*TXauth)(unsafe.Pointer(entry)).Fname, uint64((*TXauth)(unsafe.Pointer(entry)).Fname_length)) != 0) {
					break
				}
				goto _2
			_2:
				;
				type1++
			}
			if type1 < best_type {
				if best != 0 {
					XXauDisposeAuth(tls, best)
				}
				best = entry
				best_type = type1
				if type1 == 0 {
					break
				}
				goto _1
			}
		}
		XXauDisposeAuth(tls, entry)
		goto _1
	_1:
	}
	libc.Xfclose(tls, auth_file)
	return best
}

const m_E2BIG = 7
const m_EACCES = 13
const m_EADDRINUSE = 48
const m_EADDRNOTAVAIL = 49
const m_EAFNOSUPPORT = 47
const m_EAGAIN = 35
const m_EALREADY = 37
const m_EAUTH = 80
const m_EBADARCH = 86
const m_EBADEXEC = 85
const m_EBADF = 9
const m_EBADMACHO = 88
const m_EBADMSG = 94
const m_EBADRPC = 72
const m_EBUSY = 16
const m_ECANCELED = 89
const m_ECHILD = 10
const m_ECONNABORTED = 53
const m_ECONNREFUSED = 61
const m_ECONNRESET = 54
const m_EDEADLK = 11
const m_EDESTADDRREQ = 39
const m_EDEVERR = 83
const m_EDOM = 33
const m_EDQUOT = 69
const m_EEXIST = 17
const m_EFAULT = 14
const m_EFBIG = 27
const m_EFTYPE = 79
const m_EF_IS_PURGEABLE = 0x00000008
const m_EF_IS_SPARSE = 0x00000010
const m_EF_IS_SYNC_ROOT = 0x00000004
const m_EF_IS_SYNTHETIC = 0x00000020
const m_EF_MAY_SHARE_BLOCKS = 0x00000001
const m_EF_NO_XATTRS = 0x00000002
const m_EF_SHARES_ALL_BLOCKS = 0x00000040
const m_EHOSTDOWN = 64
const m_EHOSTUNREACH = 65
const m_EIDRM = 90
const m_EILSEQ = 92
const m_EINPROGRESS = 36
const m_EINTR = 4
const m_EINVAL = 22
const m_EIO = 5
const m_EISCONN = 56
const m_EISDIR = 21
const m_ELAST = 106
const m_ELOOP = 62
const m_EMFILE = 24
const m_EMLINK = 31
const m_EMSGSIZE = 40
const m_EMULTIHOP = 95
const m_ENAMETOOLONG = 63
const m_ENEEDAUTH = 81
const m_ENETDOWN = 50
const m_ENETRESET = 52
const m_ENETUNREACH = 51
const m_ENFILE = 23
const m_ENOATTR = 93
const m_ENOBUFS = 55
const m_ENODATA = 96
const m_ENODEV = 19
const m_ENOENT = 2
const m_ENOEXEC = 8
const m_ENOLCK = 77
const m_ENOLINK = 97
const m_ENOMEM = 12
const m_ENOMSG = 91
const m_ENOPOLICY = 103
const m_ENOPROTOOPT = 42
const m_ENOSPC = 28
const m_ENOSR = 98
const m_ENOSTR = 99
const m_ENOSYS = 78
const m_ENOTBLK = 15
const m_ENOTCONN = 57
const m_ENOTDIR = 20
const m_ENOTEMPTY = 66
const m_ENOTRECOVERABLE = 104
const m_ENOTSOCK = 38
const m_ENOTSUP = 45
const m_ENOTTY = 25
const m_ENXIO = 6
const m_EOPNOTSUPP = 102
const m_EOVERFLOW = 84
const m_EOWNERDEAD = 105
const m_EPERM = 1
const m_EPFNOSUPPORT = 46
const m_EPIPE = 32
const m_EPROCLIM = 67
const m_EPROCUNAVAIL = 76
const m_EPROGMISMATCH = 75
const m_EPROGUNAVAIL = 74
const m_EPROTO = 100
const m_EPROTONOSUPPORT = 43
const m_EPROTOTYPE = 41
const m_EPWROFF = 82
const m_EQFULL = 106
const m_ERANGE = 34
const m_EREMOTE = 71
const m_EROFS = 30
const m_ERPCMISMATCH = 73
const m_ESHLIBVERS = 87
const m_ESHUTDOWN = 58
const m_ESOCKTNOSUPPORT = 44
const m_ESPIPE = 29
const m_ESRCH = 3
const m_ESTALE = 70
const m_ETIME = 101
const m_ETIMEDOUT = 60
const m_ETOOMANYREFS = 59
const m_ETXTBSY = 26
const m_EUSERS = 68
const m_EWOULDBLOCK = "EAGAIN"
const m_EXDEV = 18
const m_O_CREAT1 = 512
const m_O_EXCL1 = 2048
const m_O_WRONLY1 = 1
const m_SF_APPEND = 0x00040000
const m_SF_ARCHIVED = 0x00010000
const m_SF_DATALESS = 0x40000000
const m_SF_FIRMLINK = 0x00800000
const m_SF_IMMUTABLE = 0x00020000
const m_SF_NOUNLINK = 0x00100000
const m_SF_RESTRICTED = 0x00080000
const m_SF_SETTABLE = 0x3fff0000
const m_SF_SUPPORTED = 0x009f0000
const m_SF_SYNTHETIC = 0xc0000000
const m_S_BLKSIZE = 512
const m_Time_t = "time_t"
const m_UF_APPEND = 0x00000004
const m_UF_COMPRESSED = 0x00000020
const m_UF_DATAVAULT = 0x00000080
const m_UF_HIDDEN = 0x00008000
const m_UF_IMMUTABLE = 0x00000002
const m_UF_NODUMP = 0x00000001
const m_UF_OPAQUE = 0x00000008
const m_UF_SETTABLE = 0x0000ffff
const m_UF_TRACKED = 0x00000040

type Tostat = struct {
	Fst_dev       t__uint16_t
	Fst_ino       Tino_t
	Fst_mode      Tmode_t
	Fst_nlink     Tnlink_t
	Fst_uid       t__uint16_t
	Fst_gid       t__uint16_t
	Fst_rdev      t__uint16_t
	Fst_size      t__int32_t
	Fst_atimespec Ttimespec
	Fst_mtimespec Ttimespec
	Fst_ctimespec Ttimespec
	Fst_blksize   t__int32_t
	Fst_blocks    t__int32_t
	Fst_flags     t__uint32_t
	Fst_gen       t__uint32_t
}

type Tstat = struct {
	Fst_dev           Tdev_t
	Fst_mode          Tmode_t
	Fst_nlink         Tnlink_t
	Fst_ino           t__darwin_ino64_t
	Fst_uid           Tuid_t
	Fst_gid           Tgid_t
	Fst_rdev          Tdev_t
	Fst_atimespec     Ttimespec
	Fst_mtimespec     Ttimespec
	Fst_ctimespec     Ttimespec
	Fst_birthtimespec Ttimespec
	Fst_size          Toff_t
	Fst_blocks        Tblkcnt_t
	Fst_blksize       Tblksize_t
	Fst_flags         t__uint32_t
	Fst_gen           t__uint32_t
	Fst_lspare        t__int32_t
	Fst_qspare        [2]t__int64_t
}

/*
 * Error codes
 */

/* 11 was EAGAIN */

/* math software */

/* non-blocking and interrupt i/o */

/* ipc/network software -- argument errors */

/* ipc/network software -- operational errors */

/* should be rearranged */

/* quotas & mush */

/* Network File System */

/* Intelligent device errors */

/* Program loading errors */

/* This value is only discrete when compiling __DARWIN_UNIX03, or KERNEL */

/*
 * Copyright (c) 2000, 2023 Apple Computer, Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*
 * Copyright (c) 1989, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
 * Copyright (c) 2023 Apple Computer, Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*
 * Copyright (c) 1989, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
 * Copyright (c) 2000, 2002-2006, 2008-2010, 2012 Apple Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*-
 * Copyright (c) 1998-1999 Apple Computer, Inc. All Rights Reserved
 * Copyright (c) 1991, 1993, 1994
 *	The Regents of the University of California.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)unistd.h	8.12 (Berkeley) 4/27/95
 *
 *  Copyright (c)  1998 Apple Compter, Inc.
 *  All Rights Reserved
 */

/* History:
      7/14/99 EKN at Apple fixed getdirentriesattr from getdirentryattr
      3/26/98 CHW at Apple added real interface to searchfs call
	3/5/98  CHW at Apple added hfs semantic system calls headers
*/

func XXauLockAuth(tls *libc.TLS, file_name uintptr, retries int32, timeout int32, dead int64) (r int32) {
	bp := tls.Alloc(2224)
	defer tls.Free(2224)
	var creat_fd int32
	var now Ttime_t
	var _ /* creat_name at bp+0 */ [1025]int8
	var _ /* link_name at bp+1025 */ [1025]int8
	var _ /* statb at bp+2056 */ Tstat
	_, _ = creat_fd, now
	creat_fd = -int32(1)
	if libc.Xstrlen(tls, file_name) > uint64(1022) {
		return int32(m_LOCK_ERROR)
	}
	libc.X__builtin___snprintf_chk(tls, bp, uint64(1025), 0, ^t__predefined_size_t(0), __ccgo_ts+37, libc.VaList(bp+2208, file_name))
	libc.X__builtin___snprintf_chk(tls, bp+1025, uint64(1025), 0, ^t__predefined_size_t(0), __ccgo_ts+42, libc.VaList(bp+2208, file_name))
	if libc.Xstat(tls, bp, bp+2056) != -int32(1) {
		now = libc.Xtime(tls, libc.UintptrFromInt32(0))
		/*
		 * NFS may cause ctime to be before now, special
		 * case a 0 deadtime to force lock removal
		 */
		if dead == 0 || now-(*(*Tstat)(unsafe.Pointer(bp + 2056))).Fst_ctimespec.Ftv_sec > dead {
			libc.Xremove(tls, bp)
			libc.Xremove(tls, bp+1025)
		}
	}
	for retries > 0 {
		if creat_fd == -int32(1) {
			creat_fd = libc.Xopen(tls, bp, libc.Int32FromInt32(m_O_WRONLY1)|libc.Int32FromInt32(m_O_CREAT1)|libc.Int32FromInt32(m_O_EXCL1), libc.VaList(bp+2208, int32(0600)))
			if creat_fd == -int32(1) {
				if *(*int32)(unsafe.Pointer(libc.X__error(tls))) != int32(m_EACCES) && *(*int32)(unsafe.Pointer(libc.X__error(tls))) != int32(m_EEXIST) {
					return int32(m_LOCK_ERROR)
				}
			} else {
				libc.Xclose(tls, creat_fd)
			}
		}
		if creat_fd != -int32(1) {
			/* The file system may not support hard links, and pathconf should tell us that. */
			if int64(1) == libc.Xpathconf(tls, bp, int32(m__PC_LINK_MAX)) {
				if -int32(1) == libc.Xrename(tls, bp, bp+1025) {
					/* Is this good enough?  Perhaps we should retry.  TEST */
					return int32(m_LOCK_ERROR)
				} else {
					return m_LOCK_SUCCESS
				}
			} else {
				if libc.Xlink(tls, bp, bp+1025) != -int32(1) {
					return m_LOCK_SUCCESS
				}
				if *(*int32)(unsafe.Pointer(libc.X__error(tls))) == int32(m_ENOENT) {
					creat_fd = -int32(1) /* force re-creat next time around */
					continue
				}
				if *(*int32)(unsafe.Pointer(libc.X__error(tls))) != int32(m_EEXIST) {
					return int32(m_LOCK_ERROR)
				}
			}
		}
		libc.Xsleep(tls, libc.Uint32FromInt32(timeout))
		retries--
	}
	return int32(m_LOCK_TIMEOUT)
}

/* getsubopt(3) external variable */
/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

func _read_short(tls *libc.TLS, shortp uintptr, file uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* file_short at bp+0 */ [2]uint8
	if libc.Xfread(tls, bp, uint64(2), uint64(1), file) != uint64(1) {
		return 0
	}
	*(*uint16)(unsafe.Pointer(shortp)) = libc.Uint16FromInt32(libc.Int32FromUint8((*(*[2]uint8)(unsafe.Pointer(bp)))[0])*int32(256) + libc.Int32FromUint8((*(*[2]uint8)(unsafe.Pointer(bp)))[int32(1)]))
	return int32(1)
}

func _read_counted_string(tls *libc.TLS, countp uintptr, stringp uintptr, file uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var data uintptr
	var _ /* len at bp+0 */ uint16
	_ = data
	if _read_short(tls, bp, file) == 0 {
		return 0
	}
	if libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(bp))) == 0 {
		data = libc.UintptrFromInt32(0)
	} else {
		data = libc.Xmalloc(tls, uint64(uint32(*(*uint16)(unsafe.Pointer(bp)))))
		if !(data != 0) {
			return 0
		}
		if libc.Xfread(tls, data, uint64(1), uint64(*(*uint16)(unsafe.Pointer(bp))), file) != uint64(*(*uint16)(unsafe.Pointer(bp))) {
			libc.X__builtin___memset_chk(tls, data, 0, uint64(*(*uint16)(unsafe.Pointer(bp))), ^t__predefined_size_t(0))
			libc.Xfree(tls, data)
			return 0
		}
	}
	*(*uintptr)(unsafe.Pointer(stringp)) = data
	*(*uint16)(unsafe.Pointer(countp)) = *(*uint16)(unsafe.Pointer(bp))
	return int32(1)
}

func XXauReadAuth(tls *libc.TLS, auth_file uintptr) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var ret uintptr
	var _ /* local at bp+0 */ TXauth
	_ = ret
	*(*TXauth)(unsafe.Pointer(bp)) = TXauth{}
	if _read_short(tls, bp, auth_file) == 0 {
		goto fail
	}
	if _read_counted_string(tls, bp+2, bp+8, auth_file) == 0 {
		goto fail
	}
	if _read_counted_string(tls, bp+16, bp+24, auth_file) == 0 {
		goto fail
	}
	if _read_counted_string(tls, bp+32, bp+40, auth_file) == 0 {
		goto fail
	}
	if _read_counted_string(tls, bp+48, bp+56, auth_file) == 0 {
		goto fail
	}
	ret = libc.Xmalloc(tls, uint64(64))
	if ret == libc.UintptrFromInt32(0) {
		goto fail
	}
	*(*TXauth)(unsafe.Pointer(ret)) = *(*TXauth)(unsafe.Pointer(bp))
	return ret
	goto fail
fail:
	;
	libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Faddress)
	libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fnumber)
	libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fname)
	if (*(*TXauth)(unsafe.Pointer(bp))).Fdata != 0 {
		libc.X__builtin___memset_chk(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fdata, 0, uint64((*(*TXauth)(unsafe.Pointer(bp))).Fdata_length), ^t__predefined_size_t(0))
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fdata)
	}
	return libc.UintptrFromInt32(0)
}

const m_O_CREAT2 = 0x00000200
const m_O_EXCL2 = 0x00000800
const m_O_WRONLY2 = 0x0001

/*
 * Copyright (c) 2000, 2023 Apple Computer, Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*
 * Copyright (c) 1989, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
 * Copyright (c) 2023 Apple Computer, Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*
 * Copyright (c) 1989, 1993
 *	The Regents of the University of California.  All rights reserved.
 * (c) UNIX System Laboratories, Inc.
 * All or some portions of this file are derived from material licensed
 * to the University of California by American Telephone and Telegraph
 * Co. or Unix System Laboratories, Inc. and are reproduced herein with
 * the permission of UNIX System Laboratories, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/* define X_GETTIMEOFDAY macro, a portable gettimeofday() */

/* use POSIX name for signal */

/*
 * Copyright 1997 Metro Link Incorporated
 *
 *                           All Rights Reserved
 *
 * Permission to use, copy, modify, distribute, and sell this software and its
 * documentation for any purpose is hereby granted without fee, provided that
 * the above copyright notice appear in all copies and that both that
 * copyright notice and this permission notice appear in supporting
 * documentation, and that the names of the above listed copyright holder(s)
 * not be used in advertising or publicity pertaining to distribution of
 * the software without specific, written prior permission.  The above listed
 * copyright holder(s) make(s) no representations about the suitability of
 * this software for any purpose.  It is provided "as is" without express or
 * implied warranty.
 *
 * THE ABOVE LISTED COPYRIGHT HOLDER(S) DISCLAIM(S) ALL WARRANTIES WITH REGARD
 * TO THIS SOFTWARE, INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS, IN NO EVENT SHALL THE ABOVE LISTED COPYRIGHT HOLDER(S) BE
 * LIABLE FOR ANY SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR ANY
 * DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER
 * IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING
 * OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
 * Determine the machine's byte order.
 */

/* See if it is set in the imake config first */

/*
 * Copyright (c) 2000-2007 Apple Inc. All rights reserved.
 *
 * @APPLE_OSREFERENCE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. The rights granted to you under the License
 * may not be used to create, or enable the creation or redistribution of,
 * unlawful or unlicensed copies of an Apple operating system, or to
 * circumvent, violate, or enable the circumvention or violation of, any
 * terms of an Apple operating system software license agreement.
 *
 * Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_OSREFERENCE_LICENSE_HEADER_END@
 */
/*
 * Copyright 1995 NeXT Computer, Inc. All rights reserved.
 */

func XXauUnlockAuth(tls *libc.TLS, file_name uintptr) (r int32) {
	bp := tls.Alloc(2080)
	defer tls.Free(2080)
	var _ /* creat_name at bp+0 */ [1025]int8
	var _ /* link_name at bp+1025 */ [1025]int8
	if libc.Xstrlen(tls, file_name) > uint64(1022) {
		return 0
	}
	libc.X__builtin___snprintf_chk(tls, bp, uint64(1025), 0, ^t__predefined_size_t(0), __ccgo_ts+37, libc.VaList(bp+2064, file_name))
	libc.X__builtin___snprintf_chk(tls, bp+1025, uint64(1025), 0, ^t__predefined_size_t(0), __ccgo_ts+42, libc.VaList(bp+2064, file_name))
	/*
	 * I think this is the correct order
	 */
	libc.Xremove(tls, bp)
	libc.Xremove(tls, bp+1025)
	return int32(1)
}

/* Return values from XauLockAuth */

func _write_short(tls *libc.TLS, s uint16, file uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* file_short at bp+0 */ [2]uint8
	(*(*[2]uint8)(unsafe.Pointer(bp)))[0] = uint8(uint32(s) & libc.Uint32FromInt32(0xff00) >> int32(8))
	(*(*[2]uint8)(unsafe.Pointer(bp)))[int32(1)] = libc.Uint8FromInt32(libc.Int32FromUint16(s) & int32(0xff))
	if libc.Xfwrite(tls, bp, uint64(2), uint64(1), file) != uint64(1) {
		return 0
	}
	return int32(1)
}

func _write_counted_string(tls *libc.TLS, count uint16, string1 uintptr, file uintptr) (r int32) {
	if _write_short(tls, count, file) == 0 {
		return 0
	}
	if libc.Xfwrite(tls, string1, uint64(1), uint64(count), file) != uint64(count) {
		return 0
	}
	return int32(1)
}

func XXauWriteAuth(tls *libc.TLS, auth_file uintptr, auth uintptr) (r int32) {
	if _write_short(tls, (*TXauth)(unsafe.Pointer(auth)).Ffamily, auth_file) == 0 {
		return 0
	}
	if _write_counted_string(tls, (*TXauth)(unsafe.Pointer(auth)).Faddress_length, (*TXauth)(unsafe.Pointer(auth)).Faddress, auth_file) == 0 {
		return 0
	}
	if _write_counted_string(tls, (*TXauth)(unsafe.Pointer(auth)).Fnumber_length, (*TXauth)(unsafe.Pointer(auth)).Fnumber, auth_file) == 0 {
		return 0
	}
	if _write_counted_string(tls, (*TXauth)(unsafe.Pointer(auth)).Fname_length, (*TXauth)(unsafe.Pointer(auth)).Fname, auth_file) == 0 {
		return 0
	}
	if _write_counted_string(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata_length, (*TXauth)(unsafe.Pointer(auth)).Fdata, auth_file) == 0 {
		return 0
	}
	return int32(1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "/.Xauthority\x00XAUTHORITY\x00HOME\x00%s%s\x00rb\x00%s-c\x00%s-l\x00"
