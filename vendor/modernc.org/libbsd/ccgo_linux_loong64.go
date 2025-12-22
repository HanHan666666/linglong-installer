// Code generated for linux/loong64 by 'gcc -no-main-minimize --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-link-errors -ignore-unsupported-alignment -ignore-link-errors -I /home/cznic/src/modernc.org/builder/.exclude/modernc.org/libc/include/linux/loong64 -I /home/cznic/src/modernc.org/builder/.exclude/modernc.org/limd/include/linux/loong64 -shared -DPIC .libs/setmode.o.go .libs/arc4random.o.go .libs/arc4random_uniform.o.go .libs/bsd_getopt.o.go .libs/closefrom.o.go .libs/errc.o.go .libs/expand_number.o.go .libs/explicit_bzero.o.go .libs/fgetln.o.go .libs/fgetwln.o.go .libs/fparseln.o.go .libs/flopen.o.go .libs/fmtcheck.o.go .libs/fpurge.o.go .libs/freezero.o.go .libs/funopen.o.go .libs/getbsize.o.go .libs/getpeereid.o.go .libs/dehumanize_number.o.go .libs/humanize_number.o.go .libs/inet_net_pton.o.go .libs/md5.o.go .libs/nlist.o.go .libs/pidfile.o.go .libs/setproctitle.o.go .libs/progname.o.go .libs/pwcache.o.go .libs/readpassphrase.o.go .libs/reallocarray.o.go .libs/reallocf.o.go .libs/recallocarray.o.go .libs/heapsort.o.go .libs/merge.o.go .libs/radixsort.o.go .libs/stringlist.o.go .libs/strlcat.o.go .libs/strlcpy.o.go .libs/strmode.o.go .libs/strnstr.o.go .libs/strtonum.o.go .libs/strtoi.o.go .libs/strtou.o.go .libs/timeconv.o.go .libs/unvis.o.go .libs/vis.o.go .libs/wcslcat.o.go .libs/wcslcpy.o.go -lmd -o .libs/libbsd.so.0.12.0.go', DO NOT EDIT.

//go:build linux && loong64

package libbsd

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libmd"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_ARG_MAX = 131072
const m_BC_BASE_MAX = 99
const m_BC_DIM_MAX = 2048
const m_BC_SCALE_MAX = 99
const m_BC_STRING_MAX = 1000
const m_BIG_ENDIAN = "__BIG_ENDIAN"
const m_BUFSIZ = 1024
const m_BUS_ADRALN = 1
const m_BUS_ADRERR = 2
const m_BUS_MCEERR_AO = 5
const m_BUS_MCEERR_AR = 4
const m_BUS_OBJERR = 3
const m_BYTE_ORDER = "__BYTE_ORDER"
const m_CHARCLASS_NAME_MAX = 14
const m_CHAR_BIT = 8
const m_CHAR_MAX = 255
const m_CHAR_MIN = 0
const m_CLD_CONTINUED = 6
const m_CLD_DUMPED = 3
const m_CLD_EXITED = 1
const m_CLD_KILLED = 2
const m_CLD_STOPPED = 5
const m_CLD_TRAPPED = 4
const m_CMD2_CLR = 1
const m_CMD2_GBITS = 4
const m_CMD2_OBITS = 8
const m_CMD2_SET = 2
const m_CMD2_UBITS = 16
const m_COLL_WEIGHTS_MAX = 2
const m_DELAYTIMER_MAX = 0x7fffffff
const m_E2BIG = 7
const m_EACCES = 13
const m_EADDRINUSE = 98
const m_EADDRNOTAVAIL = 99
const m_EADV = 68
const m_EAFNOSUPPORT = 97
const m_EAGAIN = 11
const m_EALREADY = 114
const m_EBADE = 52
const m_EBADF = 9
const m_EBADFD = 77
const m_EBADMSG = 74
const m_EBADR = 53
const m_EBADRQC = 56
const m_EBADSLT = 57
const m_EBFONT = 59
const m_EBUSY = 16
const m_ECANCELED = 125
const m_ECHILD = 10
const m_ECHRNG = 44
const m_ECOMM = 70
const m_ECONNABORTED = 103
const m_ECONNREFUSED = 111
const m_ECONNRESET = 104
const m_EDEADLK = 35
const m_EDEADLOCK = "EDEADLK"
const m_EDESTADDRREQ = 89
const m_EDOM = 33
const m_EDOTDOT = 73
const m_EDQUOT = 122
const m_EEXIST = 17
const m_EFAULT = 14
const m_EFBIG = 27
const m_EHOSTDOWN = 112
const m_EHOSTUNREACH = 113
const m_EHWPOISON = 133
const m_EIDRM = 43
const m_EILSEQ = 84
const m_EINPROGRESS = 115
const m_EINTR = 4
const m_EINVAL = 22
const m_EIO = 5
const m_EISCONN = 106
const m_EISDIR = 21
const m_EISNAM = 120
const m_EKEYEXPIRED = 127
const m_EKEYREJECTED = 129
const m_EKEYREVOKED = 128
const m_EL2HLT = 51
const m_EL2NSYNC = 45
const m_EL3HLT = 46
const m_EL3RST = 47
const m_ELIBACC = 79
const m_ELIBBAD = 80
const m_ELIBEXEC = 83
const m_ELIBMAX = 82
const m_ELIBSCN = 81
const m_ELNRNG = 48
const m_ELOOP = 40
const m_EMEDIUMTYPE = 124
const m_EMFILE = 24
const m_EMLINK = 31
const m_EMSGSIZE = 90
const m_EMULTIHOP = 72
const m_ENAMETOOLONG = 36
const m_ENAVAIL = 119
const m_ENETDOWN = 100
const m_ENETRESET = 102
const m_ENETUNREACH = 101
const m_ENFILE = 23
const m_ENOANO = 55
const m_ENOBUFS = 105
const m_ENOCSI = 50
const m_ENODATA = 61
const m_ENODEV = 19
const m_ENOENT = 2
const m_ENOEXEC = 8
const m_ENOKEY = 126
const m_ENOLCK = 37
const m_ENOLINK = 67
const m_ENOMEDIUM = 123
const m_ENOMEM = 12
const m_ENOMSG = 42
const m_ENONET = 64
const m_ENOPKG = 65
const m_ENOPROTOOPT = 92
const m_ENOSPC = 28
const m_ENOSR = 63
const m_ENOSTR = 60
const m_ENOSYS = 38
const m_ENOTBLK = 15
const m_ENOTCONN = 107
const m_ENOTDIR = 20
const m_ENOTEMPTY = 39
const m_ENOTNAM = 118
const m_ENOTRECOVERABLE = 131
const m_ENOTSOCK = 88
const m_ENOTSUP = "EOPNOTSUPP"
const m_ENOTTY = 25
const m_ENOTUNIQ = 76
const m_ENXIO = 6
const m_EOPNOTSUPP = 95
const m_EOVERFLOW = 75
const m_EOWNERDEAD = 130
const m_EPERM = 1
const m_EPFNOSUPPORT = 96
const m_EPIPE = 32
const m_EPROTO = 71
const m_EPROTONOSUPPORT = 93
const m_EPROTOTYPE = 91
const m_ERANGE = 34
const m_EREMCHG = 78
const m_EREMOTE = 66
const m_EREMOTEIO = 121
const m_ERESTART = 85
const m_ERFKILL = 132
const m_EROFS = 30
const m_ESHUTDOWN = 108
const m_ESOCKTNOSUPPORT = 94
const m_ESPIPE = 29
const m_ESRCH = 3
const m_ESRMNT = 69
const m_ESTALE = 116
const m_ESTRPIPE = 86
const m_ETIME = 62
const m_ETIMEDOUT = 110
const m_ETOOMANYREFS = 109
const m_ETXTBSY = 26
const m_EUCLEAN = 117
const m_EUNATCH = 49
const m_EUSERS = 87
const m_EWOULDBLOCK = "EAGAIN"
const m_EXDEV = 18
const m_EXFULL = 54
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_EXPR_NEST_MAX = 32
const m_FD_SETSIZE = 1024
const m_FILENAME_MAX = 4096
const m_FILESIZEBITS = 64
const m_FOPEN_MAX = 1000
const m_FPARSELN_UNESCALL = 0x0f
const m_FPARSELN_UNESCCOMM = 0x04
const m_FPARSELN_UNESCCONT = 0x02
const m_FPARSELN_UNESCESC = 0x01
const m_FPARSELN_UNESCREST = 0x08
const m_FPE_FLTDIV = 3
const m_FPE_FLTINV = 7
const m_FPE_FLTOVF = 4
const m_FPE_FLTRES = 6
const m_FPE_FLTSUB = 8
const m_FPE_FLTUND = 5
const m_FPE_INTDIV = 1
const m_FPE_INTOVF = 2
const m_F_LOCK = 1
const m_F_OK = 0
const m_F_TEST = 3
const m_F_TLOCK = 2
const m_F_ULOCK = 0
const m_HAVE_ASPRINTF = 1
const m_HAVE_CLEARENV = 1
const m_HAVE_CONFIG_H = 1
const m_HAVE_DECL_ENVIRON = 1
const m_HAVE_DIRENT_H = 1
const m_HAVE_DIRFD = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_FLOCK = 1
const m_HAVE_FOPENCOOKIE = 1
const m_HAVE_GETAUXVAL = 1
const m_HAVE_GETENTROPY = 1
const m_HAVE_GETLINE = 1
const m_HAVE_GRP_H = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_OPEN_MEMSTREAM = 1
const m_HAVE_PROGRAM_INVOCATION_SHORT_NAME = 1
const m_HAVE_PWD_H = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDIO_EXT_H = 1
const m_HAVE_STDIO_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYSCONF = 1
const m_HAVE_SYS_DIR_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_TYPEOF = 1
const m_HAVE_UNISTD_H = 1
const m_HAVE_VASPRINTF = 1
const m_HAVE_WCHAR_H = 1
const m_HAVE___FPURGE = 1
const m_HAVE___PROGNAME = 1
const m_HN_AUTOSCALE = 0x20
const m_HN_B = 0x04
const m_HN_DECIMAL = 0x01
const m_HN_DIVISOR_1000 = 0x08
const m_HN_GETSCALE = 0x10
const m_HN_IEC_PREFIXES = 0x10
const m_HN_NOSPACE = 0x02
const m_HOST_NAME_MAX = 255
const m_ILL_BADSTK = 8
const m_ILL_COPROC = 7
const m_ILL_ILLADR = 3
const m_ILL_ILLOPC = 1
const m_ILL_ILLOPN = 2
const m_ILL_ILLTRP = 4
const m_ILL_PRVOPC = 5
const m_ILL_PRVREG = 6
const m_INT16_MAX = 0x7fff
const m_INT32_MAX = 0x7fffffff
const m_INT64_MAX = 0x7fffffffffffffff
const m_INT8_MAX = 0x7f
const m_INTMAX_MAX = "INT64_MAX"
const m_INTMAX_MIN = "INT64_MIN"
const m_INTPTR_MAX = "INT64_MAX"
const m_INTPTR_MIN = "INT64_MIN"
const m_INT_FAST16_MAX = "INT32_MAX"
const m_INT_FAST16_MIN = "INT32_MIN"
const m_INT_FAST32_MAX = "INT32_MAX"
const m_INT_FAST32_MIN = "INT32_MIN"
const m_INT_FAST64_MAX = "INT64_MAX"
const m_INT_FAST64_MIN = "INT64_MIN"
const m_INT_FAST8_MAX = "INT8_MAX"
const m_INT_FAST8_MIN = "INT8_MIN"
const m_INT_LEAST16_MAX = "INT16_MAX"
const m_INT_LEAST16_MIN = "INT16_MIN"
const m_INT_LEAST32_MAX = "INT32_MAX"
const m_INT_LEAST32_MIN = "INT32_MIN"
const m_INT_LEAST64_MAX = "INT64_MAX"
const m_INT_LEAST64_MIN = "INT64_MIN"
const m_INT_LEAST8_MAX = "INT8_MAX"
const m_INT_LEAST8_MIN = "INT8_MIN"
const m_INT_MAX = 0x7fffffff
const m_IOV_MAX = 1024
const m_LARCH_NGREG = 32
const m_LARCH_REG_A0 = 4
const m_LARCH_REG_NARGS = 8
const m_LARCH_REG_RA = 1
const m_LARCH_REG_S0 = 23
const m_LARCH_REG_S1 = 24
const m_LARCH_REG_S2 = 25
const m_LARCH_REG_SP = 3
const m_LIBBSD_ABI_ACCMODE = 1
const m_LIBBSD_ABI_ARC4RANDOM = 1
const m_LIBBSD_ABI_ARC4RANDOM_STIR = 1
const m_LIBBSD_ABI_ASPRINTF = 0
const m_LIBBSD_ABI_BSD_GETOPT = 1
const m_LIBBSD_ABI_CLOSEFROM = 1
const m_LIBBSD_ABI_ERR = 0
const m_LIBBSD_ABI_ERRC = 1
const m_LIBBSD_ABI_EXPAND_NUMBER = 1
const m_LIBBSD_ABI_EXPLICIT_BZERO = 1
const m_LIBBSD_ABI_FGETLN = 1
const m_LIBBSD_ABI_FLOPEN = 1
const m_LIBBSD_ABI_FMTCHECK = 1
const m_LIBBSD_ABI_FPURGE = 1
const m_LIBBSD_ABI_FREEZERO = 1
const m_LIBBSD_ABI_FUNOPEN = 1
const m_LIBBSD_ABI_GETBSIZE = 1
const m_LIBBSD_ABI_GETPEEREID = 1
const m_LIBBSD_ABI_HUMANIZE_NUMBER = 1
const m_LIBBSD_ABI_ID_FROM_NAME = 1
const m_LIBBSD_ABI_INET_NET_PTON = 1
const m_LIBBSD_ABI_MD5 = 1
const m_LIBBSD_ABI_NAME_FROM_ID = 1
const m_LIBBSD_ABI_NLIST = 1
const m_LIBBSD_ABI_PIDFILE = 1
const m_LIBBSD_ABI_PROCTITLE = 1
const m_LIBBSD_ABI_PROGNAME = 1
const m_LIBBSD_ABI_PWCACHE = 1
const m_LIBBSD_ABI_READPASSPHRASE = 1
const m_LIBBSD_ABI_REALLOCARRAY = 1
const m_LIBBSD_ABI_REALLOCF = 1
const m_LIBBSD_ABI_RECALLOCARRAY = 1
const m_LIBBSD_ABI_SORT = 1
const m_LIBBSD_ABI_STRINGLIST = 1
const m_LIBBSD_ABI_STRL = 1
const m_LIBBSD_ABI_STRMODE = 1
const m_LIBBSD_ABI_STRNSTR = 1
const m_LIBBSD_ABI_STRTONUM = 1
const m_LIBBSD_ABI_STRTOX = 1
const m_LIBBSD_ABI_TIME64 = 0
const m_LIBBSD_ABI_TIMECONV = 1
const m_LIBBSD_ABI_TRANSPARENT_LIBMD = 1
const m_LIBBSD_ABI_VIS = 1
const m_LIBBSD_ABI_WCSL = 1
const m_LIBBSD_DISABLE_DEPRECATED = 1
const m_LIBBSD_OVERLAY = 1
const m_LIBBSD_SYS_HAS_TIME64 = 1
const m_LIBBSD_SYS_TIME_BITS = 64
const m_LINE_MAX = 4096
const m_LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const m_LLONG_MAX = 0x7fffffffffffffff
const m_LOGIN_NAME_MAX = 256
const m_LONG_BIT = 64
const m_LONG_MAX = "__LONG_MAX"
const m_LT_OBJDIR = ".libs/"
const m_L_INCR = 1
const m_L_SET = 0
const m_L_XTND = 2
const m_L_ctermid = 20
const m_L_cuserid = 20
const m_L_tmpnam = 20
const m_MB_LEN_MAX = 4
const m_MINSIGSTKSZ = 4096
const m_MQ_PRIO_MAX = 32768
const m_NAME_MAX = 255
const m_NDEBUG = 1
const m_NGROUPS_MAX = 32
const m_NL_ARGMAX = 9
const m_NL_LANGMAX = 32
const m_NL_MSGMAX = 32767
const m_NL_NMAX = 16
const m_NL_SETMAX = 255
const m_NL_TEXTMAX = 2048
const m_NSIG = "_NSIG"
const m_NZERO = 20
const m_PACKAGE = "libbsd"
const m_PACKAGE_BUGREPORT = "libbsd@lists.freedesktop.org"
const m_PACKAGE_NAME = "libbsd"
const m_PACKAGE_STRING = "libbsd 0.12.0"
const m_PACKAGE_TARNAME = "libbsd"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "0.12.0"
const m_PATH_MAX = 4096
const m_PDP_ENDIAN = "__PDP_ENDIAN"
const m_PIC = 1
const m_PIPE_BUF = 4096
const m_POLL_ERR = 4
const m_POLL_HUP = 6
const m_POLL_IN = 1
const m_POLL_MSG = 3
const m_POLL_OUT = 2
const m_POLL_PRI = 5
const m_POSIX_CLOSE_RESTART = 0
const m_PTHREAD_DESTRUCTOR_ITERATIONS = 4
const m_PTHREAD_KEYS_MAX = 128
const m_PTHREAD_STACK_MIN = 2048
const m_PTRDIFF_MAX = "INT64_MAX"
const m_PTRDIFF_MIN = "INT64_MIN"
const m_P_tmpdir = "/tmp"
const m_RAND_MAX = 0x7fffffff
const m_RE_DUP_MAX = 255
const m_R_OK = 4
const m_SA_EXPOSE_TAGBITS = 0x00000800
const m_SA_NOCLDSTOP = 1
const m_SA_NOCLDWAIT = 2
const m_SA_NODEFER = 0x40000000
const m_SA_NOMASK = "SA_NODEFER"
const m_SA_ONESHOT = "SA_RESETHAND"
const m_SA_ONSTACK = 0x08000000
const m_SA_RESETHAND = 0x80000000
const m_SA_RESTART = 0x10000000
const m_SA_SIGINFO = 4
const m_SA_UNSUPPORTED = 0x00000400
const m_SCHAR_MAX = 127
const m_SEEK_DATA = 3
const m_SEEK_HOLE = 4
const m_SEGV_ACCERR = 2
const m_SEGV_BNDERR = 3
const m_SEGV_MAPERR = 1
const m_SEGV_MTEAERR = 8
const m_SEGV_MTESERR = 9
const m_SEGV_PKUERR = 4
const m_SEM_NSEMS_MAX = 256
const m_SEM_VALUE_MAX = 0x7fffffff
const m_SET_LEN = 6
const m_SET_LEN_INCR = 4
const m_SHRT_MAX = 0x7fff
const m_SIGABRT = 6
const m_SIGALRM = 14
const m_SIGBUS = 7
const m_SIGCHLD = 17
const m_SIGCONT = 18
const m_SIGEV_NONE = 1
const m_SIGEV_SIGNAL = 0
const m_SIGEV_THREAD = 2
const m_SIGEV_THREAD_ID = 4
const m_SIGFPE = 8
const m_SIGHUP = 1
const m_SIGILL = 4
const m_SIGINT = 2
const m_SIGIO = 29
const m_SIGIOT = "SIGABRT"
const m_SIGKILL = 9
const m_SIGPIPE = 13
const m_SIGPOLL = "SIGIO"
const m_SIGPROF = 27
const m_SIGPWR = 30
const m_SIGQUIT = 3
const m_SIGSEGV = 11
const m_SIGSTKFLT = 16
const m_SIGSTKSZ = 16384
const m_SIGSTOP = 19
const m_SIGSYS = 31
const m_SIGTERM = 15
const m_SIGTRAP = 5
const m_SIGTSTP = 20
const m_SIGTTIN = 21
const m_SIGTTOU = 22
const m_SIGUNUSED = "SIGSYS"
const m_SIGURG = 23
const m_SIGUSR1 = 10
const m_SIGUSR2 = 12
const m_SIGVTALRM = 26
const m_SIGWINCH = 28
const m_SIGXCPU = 24
const m_SIGXFSZ = 25
const m_SIG_ATOMIC_MAX = "INT32_MAX"
const m_SIG_ATOMIC_MIN = "INT32_MIN"
const m_SIG_BLOCK = 0
const m_SIG_SETMASK = 2
const m_SIG_UNBLOCK = 1
const m_SIZEOF_TIME_T = 8
const m_SIZE_MAX = "UINT64_MAX"
const m_SI_KERNEL = 128
const m_SI_USER = 0
const m_SSIZE_MAX = "LONG_MAX"
const m_SS_DISABLE = 2
const m_SS_FLAG_BITS = "SS_AUTODISARM"
const m_SS_ONSTACK = 1
const m_STATX_ALL = 0xfff
const m_STATX_ATIME = 0x20
const m_STATX_BASIC_STATS = 0x7ff
const m_STATX_BLOCKS = 0x400
const m_STATX_BTIME = 0x800
const m_STATX_CTIME = 0x80
const m_STATX_GID = 0x10
const m_STATX_INO = 0x100
const m_STATX_MODE = 2
const m_STATX_MTIME = 0x40
const m_STATX_NLINK = 4
const m_STATX_SIZE = 0x200
const m_STATX_TYPE = 1
const m_STATX_UID = 8
const m_STDC_HEADERS = 1
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_SYMLOOP_MAX = 40
const m_SYS_SECCOMP = 1
const m_SYS_USER_DISPATCH = 2
const m_S_IEXEC = "S_IXUSR"
const m_S_IFBLK = 0060000
const m_S_IFCHR = 0020000
const m_S_IFDIR = 16384
const m_S_IFIFO = 0010000
const m_S_IFLNK = 0120000
const m_S_IFMT = 0170000
const m_S_IFREG = 0100000
const m_S_IFSOCK = 0140000
const m_S_IREAD = "S_IRUSR"
const m_S_IRGRP = 32
const m_S_IROTH = 4
const m_S_IRUSR = 256
const m_S_IRWXG = 56
const m_S_IRWXO = 7
const m_S_IRWXU = 448
const m_S_ISGID = 1024
const m_S_ISTXT = "S_ISVTX"
const m_S_ISUID = 2048
const m_S_ISVTX = 512
const m_S_IWGRP = 16
const m_S_IWOTH = 2
const m_S_IWRITE = "S_IWUSR"
const m_S_IWUSR = 128
const m_S_IXGRP = 8
const m_S_IXOTH = 1
const m_S_IXUSR = 64
const m_TMP_MAX = 10000
const m_TRAP_BRANCH = 3
const m_TRAP_BRKPT = 1
const m_TRAP_HWBKPT = 4
const m_TRAP_TRACE = 2
const m_TRAP_UNK = 5
const m_TTY_NAME_MAX = 32
const m_TZNAME_MAX = 6
const m_UCHAR_MAX = 255
const m_UINT16_MAX = 0xffff
const m_UINT32_MAX = "0xffffffffu"
const m_UINT64_MAX = "0xffffffffffffffffu"
const m_UINT8_MAX = 0xff
const m_UINTMAX_MAX = "UINT64_MAX"
const m_UINTPTR_MAX = "UINT64_MAX"
const m_UINT_FAST16_MAX = "UINT32_MAX"
const m_UINT_FAST32_MAX = "UINT32_MAX"
const m_UINT_FAST64_MAX = "UINT64_MAX"
const m_UINT_FAST8_MAX = "UINT8_MAX"
const m_UINT_LEAST16_MAX = "UINT16_MAX"
const m_UINT_LEAST32_MAX = "UINT32_MAX"
const m_UINT_LEAST64_MAX = "UINT64_MAX"
const m_UINT_LEAST8_MAX = "UINT8_MAX"
const m_UINT_MAX = 0xffffffff
const m_USHRT_MAX = 0xffff
const m_UTIME_NOW = 0x3fffffff
const m_UTIME_OMIT = 0x3ffffffe
const m_VERSION = "0.12.0"
const m_WINT_MAX = "UINT32_MAX"
const m_WINT_MIN = 0
const m_WNOHANG = 1
const m_WORD_BIT = 32
const m_WUNTRACED = 2
const m_W_OK = 2
const m_X_OK = 1
const m__ABILP64 = 3
const m__ALL_SOURCE = 1
const m__CS_GNU_LIBC_VERSION = 2
const m__CS_GNU_LIBPTHREAD_VERSION = 3
const m__CS_PATH = 0
const m__CS_POSIX_V5_WIDTH_RESTRICTED_ENVS = 4
const m__CS_POSIX_V6_ILP32_OFF32_CFLAGS = 1116
const m__CS_POSIX_V6_ILP32_OFF32_LDFLAGS = 1117
const m__CS_POSIX_V6_ILP32_OFF32_LIBS = 1118
const m__CS_POSIX_V6_ILP32_OFF32_LINTFLAGS = 1119
const m__CS_POSIX_V6_ILP32_OFFBIG_CFLAGS = 1120
const m__CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS = 1121
const m__CS_POSIX_V6_ILP32_OFFBIG_LIBS = 1122
const m__CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS = 1123
const m__CS_POSIX_V6_LP64_OFF64_CFLAGS = 1124
const m__CS_POSIX_V6_LP64_OFF64_LDFLAGS = 1125
const m__CS_POSIX_V6_LP64_OFF64_LIBS = 1126
const m__CS_POSIX_V6_LP64_OFF64_LINTFLAGS = 1127
const m__CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS = 1128
const m__CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS = 1129
const m__CS_POSIX_V6_LPBIG_OFFBIG_LIBS = 1130
const m__CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS = 1131
const m__CS_POSIX_V6_WIDTH_RESTRICTED_ENVS = 1
const m__CS_POSIX_V7_ILP32_OFF32_CFLAGS = 1132
const m__CS_POSIX_V7_ILP32_OFF32_LDFLAGS = 1133
const m__CS_POSIX_V7_ILP32_OFF32_LIBS = 1134
const m__CS_POSIX_V7_ILP32_OFF32_LINTFLAGS = 1135
const m__CS_POSIX_V7_ILP32_OFFBIG_CFLAGS = 1136
const m__CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS = 1137
const m__CS_POSIX_V7_ILP32_OFFBIG_LIBS = 1138
const m__CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS = 1139
const m__CS_POSIX_V7_LP64_OFF64_CFLAGS = 1140
const m__CS_POSIX_V7_LP64_OFF64_LDFLAGS = 1141
const m__CS_POSIX_V7_LP64_OFF64_LIBS = 1142
const m__CS_POSIX_V7_LP64_OFF64_LINTFLAGS = 1143
const m__CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS = 1144
const m__CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS = 1145
const m__CS_POSIX_V7_LPBIG_OFFBIG_LIBS = 1146
const m__CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS = 1147
const m__CS_POSIX_V7_THREADS_CFLAGS = 1150
const m__CS_POSIX_V7_THREADS_LDFLAGS = 1151
const m__CS_POSIX_V7_WIDTH_RESTRICTED_ENVS = 5
const m__CS_V6_ENV = 1148
const m__CS_V7_ENV = 1149
const m__DARWIN_C_SOURCE = 1
const m__GNU_SOURCE = 1
const m__HPUX_ALT_XOPEN_SOCKET_API = 1
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
const m__LOONGARCH_ARCH = "la64v1.0"
const m__LOONGARCH_FPSET = 32
const m__LOONGARCH_SIM = "_ABILP64"
const m__LOONGARCH_SPFPSET = 32
const m__LOONGARCH_SZINT = 32
const m__LOONGARCH_SZLONG = 64
const m__LOONGARCH_SZPTR = 64
const m__LOONGARCH_TUNE = "generic"
const m__LP64 = 1
const m__NETBSD_SOURCE = 1
const m__NSIG = 65
const m__OPENBSD_SOURCE = 1
const m__PC_2_SYMLINKS = 20
const m__PC_ALLOC_SIZE_MIN = 18
const m__PC_ASYNC_IO = 10
const m__PC_CHOWN_RESTRICTED = 6
const m__PC_FILESIZEBITS = 13
const m__PC_LINK_MAX = 0
const m__PC_MAX_CANON = 1
const m__PC_MAX_INPUT = 2
const m__PC_NAME_MAX = 3
const m__PC_NO_TRUNC = 7
const m__PC_PATH_MAX = 4
const m__PC_PIPE_BUF = 5
const m__PC_PRIO_IO = 11
const m__PC_REC_INCR_XFER_SIZE = 14
const m__PC_REC_MAX_XFER_SIZE = 15
const m__PC_REC_MIN_XFER_SIZE = 16
const m__PC_REC_XFER_ALIGN = 17
const m__PC_SOCK_MAXBUF = 12
const m__PC_SYMLINK_MAX = 19
const m__PC_SYNC_IO = 9
const m__PC_VDISABLE = 8
const m__POSIX2_BC_BASE_MAX = 99
const m__POSIX2_BC_DIM_MAX = 2048
const m__POSIX2_BC_SCALE_MAX = 99
const m__POSIX2_BC_STRING_MAX = 1000
const m__POSIX2_CHARCLASS_NAME_MAX = 14
const m__POSIX2_COLL_WEIGHTS_MAX = 2
const m__POSIX2_C_BIND = "_POSIX_VERSION"
const m__POSIX2_EXPR_NEST_MAX = 32
const m__POSIX2_LINE_MAX = 2048
const m__POSIX2_RE_DUP_MAX = 255
const m__POSIX2_VERSION = "_POSIX_VERSION"
const m__POSIX_ADVISORY_INFO = "_POSIX_VERSION"
const m__POSIX_AIO_LISTIO_MAX = 2
const m__POSIX_AIO_MAX = 1
const m__POSIX_ARG_MAX = 4096
const m__POSIX_ASYNCHRONOUS_IO = "_POSIX_VERSION"
const m__POSIX_BARRIERS = "_POSIX_VERSION"
const m__POSIX_CHILD_MAX = 25
const m__POSIX_CHOWN_RESTRICTED = 1
const m__POSIX_CLOCKRES_MIN = 20000000
const m__POSIX_CLOCK_SELECTION = "_POSIX_VERSION"
const m__POSIX_CPUTIME = "_POSIX_VERSION"
const m__POSIX_DELAYTIMER_MAX = 32
const m__POSIX_FSYNC = "_POSIX_VERSION"
const m__POSIX_HOST_NAME_MAX = 255
const m__POSIX_IPV6 = "_POSIX_VERSION"
const m__POSIX_JOB_CONTROL = 1
const m__POSIX_LINK_MAX = 8
const m__POSIX_LOGIN_NAME_MAX = 9
const m__POSIX_MAPPED_FILES = "_POSIX_VERSION"
const m__POSIX_MAX_CANON = 255
const m__POSIX_MAX_INPUT = 255
const m__POSIX_MEMLOCK = "_POSIX_VERSION"
const m__POSIX_MEMLOCK_RANGE = "_POSIX_VERSION"
const m__POSIX_MEMORY_PROTECTION = "_POSIX_VERSION"
const m__POSIX_MESSAGE_PASSING = "_POSIX_VERSION"
const m__POSIX_MONOTONIC_CLOCK = "_POSIX_VERSION"
const m__POSIX_MQ_OPEN_MAX = 8
const m__POSIX_MQ_PRIO_MAX = 32
const m__POSIX_NAME_MAX = 14
const m__POSIX_NGROUPS_MAX = 8
const m__POSIX_NO_TRUNC = 1
const m__POSIX_OPEN_MAX = 20
const m__POSIX_PATH_MAX = 256
const m__POSIX_PIPE_BUF = 512
const m__POSIX_PTHREAD_SEMANTICS = 1
const m__POSIX_RAW_SOCKETS = "_POSIX_VERSION"
const m__POSIX_READER_WRITER_LOCKS = "_POSIX_VERSION"
const m__POSIX_REALTIME_SIGNALS = "_POSIX_VERSION"
const m__POSIX_REGEXP = 1
const m__POSIX_RE_DUP_MAX = 255
const m__POSIX_RTSIG_MAX = 8
const m__POSIX_SAVED_IDS = 1
const m__POSIX_SEMAPHORES = "_POSIX_VERSION"
const m__POSIX_SEM_NSEMS_MAX = 256
const m__POSIX_SEM_VALUE_MAX = 32767
const m__POSIX_SHARED_MEMORY_OBJECTS = "_POSIX_VERSION"
const m__POSIX_SHELL = 1
const m__POSIX_SIGQUEUE_MAX = 32
const m__POSIX_SPAWN = "_POSIX_VERSION"
const m__POSIX_SPIN_LOCKS = "_POSIX_VERSION"
const m__POSIX_SSIZE_MAX = 32767
const m__POSIX_SS_REPL_MAX = 4
const m__POSIX_STREAM_MAX = 8
const m__POSIX_SYMLINK_MAX = 255
const m__POSIX_SYMLOOP_MAX = 8
const m__POSIX_THREADS = "_POSIX_VERSION"
const m__POSIX_THREAD_ATTR_STACKADDR = "_POSIX_VERSION"
const m__POSIX_THREAD_ATTR_STACKSIZE = "_POSIX_VERSION"
const m__POSIX_THREAD_CPUTIME = "_POSIX_VERSION"
const m__POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const m__POSIX_THREAD_KEYS_MAX = 128
const m__POSIX_THREAD_PRIORITY_SCHEDULING = "_POSIX_VERSION"
const m__POSIX_THREAD_PROCESS_SHARED = "_POSIX_VERSION"
const m__POSIX_THREAD_SAFE_FUNCTIONS = "_POSIX_VERSION"
const m__POSIX_THREAD_THREADS_MAX = 64
const m__POSIX_TIMEOUTS = "_POSIX_VERSION"
const m__POSIX_TIMERS = "_POSIX_VERSION"
const m__POSIX_TIMER_MAX = 32
const m__POSIX_TRACE_EVENT_NAME_MAX = 30
const m__POSIX_TRACE_NAME_MAX = 8
const m__POSIX_TRACE_SYS_MAX = 8
const m__POSIX_TRACE_USER_EVENT_MAX = 32
const m__POSIX_TTY_NAME_MAX = 9
const m__POSIX_TZNAME_MAX = 6
const m__POSIX_V6_LP64_OFF64 = 1
const m__POSIX_V7_LP64_OFF64 = 1
const m__POSIX_VDISABLE = 0
const m__POSIX_VERSION = 200809
const m__SC_2_CHAR_TERM = 95
const m__SC_2_C_BIND = 47
const m__SC_2_C_DEV = 48
const m__SC_2_FORT_DEV = 49
const m__SC_2_FORT_RUN = 50
const m__SC_2_LOCALEDEF = 52
const m__SC_2_PBS = 168
const m__SC_2_PBS_ACCOUNTING = 169
const m__SC_2_PBS_CHECKPOINT = 175
const m__SC_2_PBS_LOCATE = 170
const m__SC_2_PBS_MESSAGE = 171
const m__SC_2_PBS_TRACK = 172
const m__SC_2_SW_DEV = 51
const m__SC_2_UPE = 97
const m__SC_2_VERSION = 46
const m__SC_ADVISORY_INFO = 132
const m__SC_AIO_LISTIO_MAX = 23
const m__SC_AIO_MAX = 24
const m__SC_AIO_PRIO_DELTA_MAX = 25
const m__SC_ARG_MAX = 0
const m__SC_ASYNCHRONOUS_IO = 12
const m__SC_ATEXIT_MAX = 87
const m__SC_AVPHYS_PAGES = 86
const m__SC_BARRIERS = 133
const m__SC_BC_BASE_MAX = 36
const m__SC_BC_DIM_MAX = 37
const m__SC_BC_SCALE_MAX = 38
const m__SC_BC_STRING_MAX = 39
const m__SC_CHILD_MAX = 1
const m__SC_CLK_TCK = 2
const m__SC_CLOCK_SELECTION = 137
const m__SC_COLL_WEIGHTS_MAX = 40
const m__SC_CPUTIME = 138
const m__SC_DELAYTIMER_MAX = 26
const m__SC_EXPR_NEST_MAX = 42
const m__SC_FSYNC = 15
const m__SC_GETGR_R_SIZE_MAX = 69
const m__SC_GETPW_R_SIZE_MAX = 70
const m__SC_HOST_NAME_MAX = 180
const m__SC_IOV_MAX = 60
const m__SC_IPV6 = 235
const m__SC_JOB_CONTROL = 7
const m__SC_LINE_MAX = 43
const m__SC_LOGIN_NAME_MAX = 71
const m__SC_MAPPED_FILES = 16
const m__SC_MEMLOCK = 17
const m__SC_MEMLOCK_RANGE = 18
const m__SC_MEMORY_PROTECTION = 19
const m__SC_MESSAGE_PASSING = 20
const m__SC_MINSIGSTKSZ = 249
const m__SC_MONOTONIC_CLOCK = 149
const m__SC_MQ_OPEN_MAX = 27
const m__SC_MQ_PRIO_MAX = 28
const m__SC_NGROUPS_MAX = 3
const m__SC_NPROCESSORS_CONF = 83
const m__SC_NPROCESSORS_ONLN = 84
const m__SC_NZERO = 109
const m__SC_OPEN_MAX = 4
const m__SC_PAGESIZE = 30
const m__SC_PAGE_SIZE = 30
const m__SC_PASS_MAX = 88
const m__SC_PHYS_PAGES = 85
const m__SC_PRIORITIZED_IO = 13
const m__SC_PRIORITY_SCHEDULING = 10
const m__SC_RAW_SOCKETS = 236
const m__SC_READER_WRITER_LOCKS = 153
const m__SC_REALTIME_SIGNALS = 9
const m__SC_REGEXP = 155
const m__SC_RE_DUP_MAX = 44
const m__SC_RTSIG_MAX = 31
const m__SC_SAVED_IDS = 8
const m__SC_SEMAPHORES = 21
const m__SC_SEM_NSEMS_MAX = 32
const m__SC_SEM_VALUE_MAX = 33
const m__SC_SHARED_MEMORY_OBJECTS = 22
const m__SC_SHELL = 157
const m__SC_SIGQUEUE_MAX = 34
const m__SC_SIGSTKSZ = 250
const m__SC_SPAWN = 159
const m__SC_SPIN_LOCKS = 154
const m__SC_SPORADIC_SERVER = 160
const m__SC_SS_REPL_MAX = 241
const m__SC_STREAMS = 174
const m__SC_STREAM_MAX = 5
const m__SC_SYMLOOP_MAX = 173
const m__SC_SYNCHRONIZED_IO = 14
const m__SC_THREADS = 67
const m__SC_THREAD_ATTR_STACKADDR = 77
const m__SC_THREAD_ATTR_STACKSIZE = 78
const m__SC_THREAD_CPUTIME = 139
const m__SC_THREAD_DESTRUCTOR_ITERATIONS = 73
const m__SC_THREAD_KEYS_MAX = 74
const m__SC_THREAD_PRIORITY_SCHEDULING = 79
const m__SC_THREAD_PRIO_INHERIT = 80
const m__SC_THREAD_PRIO_PROTECT = 81
const m__SC_THREAD_PROCESS_SHARED = 82
const m__SC_THREAD_ROBUST_PRIO_INHERIT = 247
const m__SC_THREAD_ROBUST_PRIO_PROTECT = 248
const m__SC_THREAD_SAFE_FUNCTIONS = 68
const m__SC_THREAD_SPORADIC_SERVER = 161
const m__SC_THREAD_STACK_MIN = 75
const m__SC_THREAD_THREADS_MAX = 76
const m__SC_TIMEOUTS = 164
const m__SC_TIMERS = 11
const m__SC_TIMER_MAX = 35
const m__SC_TRACE = 181
const m__SC_TRACE_EVENT_FILTER = 182
const m__SC_TRACE_EVENT_NAME_MAX = 242
const m__SC_TRACE_INHERIT = 183
const m__SC_TRACE_LOG = 184
const m__SC_TRACE_NAME_MAX = 243
const m__SC_TRACE_SYS_MAX = 244
const m__SC_TRACE_USER_EVENT_MAX = 245
const m__SC_TTY_NAME_MAX = 72
const m__SC_TYPED_MEMORY_OBJECTS = 165
const m__SC_TZNAME_MAX = 6
const m__SC_UIO_MAXIOV = 60
const m__SC_V6_ILP32_OFF32 = 176
const m__SC_V6_ILP32_OFFBIG = 177
const m__SC_V6_LP64_OFF64 = 178
const m__SC_V6_LPBIG_OFFBIG = 179
const m__SC_V7_ILP32_OFF32 = 237
const m__SC_V7_ILP32_OFFBIG = 238
const m__SC_V7_LP64_OFF64 = 239
const m__SC_V7_LPBIG_OFFBIG = 240
const m__SC_VERSION = 29
const m__SC_XBS5_ILP32_OFF32 = 125
const m__SC_XBS5_ILP32_OFFBIG = 126
const m__SC_XBS5_LP64_OFF64 = 127
const m__SC_XBS5_LPBIG_OFFBIG = 128
const m__SC_XOPEN_CRYPT = 92
const m__SC_XOPEN_ENH_I18N = 93
const m__SC_XOPEN_LEGACY = 129
const m__SC_XOPEN_REALTIME = 130
const m__SC_XOPEN_REALTIME_THREADS = 131
const m__SC_XOPEN_SHM = 94
const m__SC_XOPEN_STREAMS = 246
const m__SC_XOPEN_UNIX = 91
const m__SC_XOPEN_VERSION = 89
const m__SC_XOPEN_XCU_VERSION = 90
const m__SC_XOPEN_XPG2 = 98
const m__SC_XOPEN_XPG3 = 99
const m__SC_XOPEN_XPG4 = 100
const m__STDC_PREDEF_H = 1
const m__SYS_CDEFS_H = 1
const m__TANDEM_SOURCE = 1
const m__XOPEN_ENH_I18N = 1
const m__XOPEN_IOV_MAX = 16
const m__XOPEN_NAME_MAX = 255
const m__XOPEN_PATH_MAX = 1024
const m__XOPEN_UNIX = 1
const m__XOPEN_VERSION = 700
const m___ACCUM_EPSILON__ = "0x1P-15K"
const m___ACCUM_FBIT__ = 15
const m___ACCUM_IBIT__ = 16
const m___ACCUM_MAX__ = "0X7FFFFFFFP-15K"
const m___ATOMIC_ACQUIRE = 2
const m___ATOMIC_ACQ_REL = 4
const m___ATOMIC_CONSUME = 1
const m___ATOMIC_RELAXED = 0
const m___ATOMIC_RELEASE = 3
const m___ATOMIC_SEQ_CST = 5
const m___BIGGEST_ALIGNMENT__ = 16
const m___BIG_ENDIAN = 4321
const m___BYTE_ORDER = 1234
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___CCGO__ = 1
const m___CHAR_BIT__ = 8
const m___DA_FBIT__ = 31
const m___DA_IBIT__ = 32
const m___DBL_DECIMAL_DIG__ = 17
const m___DBL_DIG__ = 15
const m___DBL_HAS_DENORM__ = 1
const m___DBL_HAS_INFINITY__ = 1
const m___DBL_HAS_QUIET_NAN__ = 1
const m___DBL_IS_IEC_60559__ = 1
const m___DBL_MANT_DIG__ = 53
const m___DBL_MAX_10_EXP__ = 308
const m___DBL_MAX_EXP__ = 1024
const m___DECIMAL_DIG__ = 36
const m___DEC_EVAL_METHOD__ = 2
const m___DQ_FBIT__ = 63
const m___DQ_IBIT__ = 0
const m___ELF__ = 1
const m___EXTENSIONS__ = 1
const m___FINITE_MATH_ONLY__ = 0
const m___FLOAT128_TYPE__ = 1
const m___FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___FLT128_DECIMAL_DIG__ = 36
const m___FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___FLT128_DIG__ = 33
const m___FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___FLT128_HAS_DENORM__ = 1
const m___FLT128_HAS_INFINITY__ = 1
const m___FLT128_HAS_QUIET_NAN__ = 1
const m___FLT128_IS_IEC_60559__ = 1
const m___FLT128_MANT_DIG__ = 113
const m___FLT128_MAX_10_EXP__ = 4932
const m___FLT128_MAX_EXP__ = 16384
const m___FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT32X_DECIMAL_DIG__ = 17
const m___FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___FLT32X_DIG__ = 15
const m___FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___FLT32X_HAS_DENORM__ = 1
const m___FLT32X_HAS_INFINITY__ = 1
const m___FLT32X_HAS_QUIET_NAN__ = 1
const m___FLT32X_IS_IEC_60559__ = 1
const m___FLT32X_MANT_DIG__ = 53
const m___FLT32X_MAX_10_EXP__ = 308
const m___FLT32X_MAX_EXP__ = 1024
const m___FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const m___FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT32_DECIMAL_DIG__ = 9
const m___FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const m___FLT32_DIG__ = 6
const m___FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const m___FLT32_HAS_DENORM__ = 1
const m___FLT32_HAS_INFINITY__ = 1
const m___FLT32_HAS_QUIET_NAN__ = 1
const m___FLT32_IS_IEC_60559__ = 1
const m___FLT32_MANT_DIG__ = 24
const m___FLT32_MAX_10_EXP__ = 38
const m___FLT32_MAX_EXP__ = 128
const m___FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const m___FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT64X_DECIMAL_DIG__ = 36
const m___FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___FLT64X_DIG__ = 33
const m___FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___FLT64X_HAS_DENORM__ = 1
const m___FLT64X_HAS_INFINITY__ = 1
const m___FLT64X_HAS_QUIET_NAN__ = 1
const m___FLT64X_IS_IEC_60559__ = 1
const m___FLT64X_MANT_DIG__ = 113
const m___FLT64X_MAX_10_EXP__ = 4932
const m___FLT64X_MAX_EXP__ = 16384
const m___FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT64_DECIMAL_DIG__ = 17
const m___FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___FLT64_DIG__ = 15
const m___FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___FLT64_HAS_DENORM__ = 1
const m___FLT64_HAS_INFINITY__ = 1
const m___FLT64_HAS_QUIET_NAN__ = 1
const m___FLT64_IS_IEC_60559__ = 1
const m___FLT64_MANT_DIG__ = 53
const m___FLT64_MAX_10_EXP__ = 308
const m___FLT64_MAX_EXP__ = 1024
const m___FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const m___FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT_DECIMAL_DIG__ = 9
const m___FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const m___FLT_DIG__ = 6
const m___FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const m___FLT_EVAL_METHOD_TS_18661_3__ = 0
const m___FLT_EVAL_METHOD__ = 0
const m___FLT_HAS_DENORM__ = 1
const m___FLT_HAS_INFINITY__ = 1
const m___FLT_HAS_QUIET_NAN__ = 1
const m___FLT_IS_IEC_60559__ = 1
const m___FLT_MANT_DIG__ = 24
const m___FLT_MAX_10_EXP__ = 38
const m___FLT_MAX_EXP__ = 128
const m___FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const m___FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT_RADIX__ = 2
const m___FP_FAST_FMA = 1
const m___FP_FAST_FMAF = 1
const m___FP_FAST_FMAF32 = 1
const m___FP_FAST_FMAF32x = 1
const m___FP_FAST_FMAF64 = 1
const m___FRACT_EPSILON__ = "0x1P-15R"
const m___FRACT_FBIT__ = 15
const m___FRACT_IBIT__ = 0
const m___FRACT_MAX__ = "0X7FFFP-15R"
const m___FUNCTION__ = "__func__"
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
const m___GCC_HAVE_DWARF2_CFI_ASM = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const m___GCC_IEC_559 = 2
const m___GCC_IEC_559_COMPLEX = 2
const m___GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const m___GNUC_MINOR__ = 2
const m___GNUC_PATCHLEVEL__ = 1
const m___GNUC_RH_RELEASE__ = 6
const m___GNUC_STDC_INLINE__ = 1
const m___GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const m___GNUC__ = 14
const m___GXX_ABI_VERSION = 1019
const m___HAVE_GENERIC_SELECTION = 0
const m___HAVE_SPECULATION_SAFE_VALUE = 1
const m___HA_FBIT__ = 7
const m___HA_IBIT__ = 8
const m___HQ_FBIT__ = 15
const m___HQ_IBIT__ = 0
const m___INT16_MAX__ = 0x7fff
const m___INT32_MAX__ = 0x7fffffff
const m___INT32_TYPE__ = "int"
const m___INT64_MAX__ = 0x7fffffffffffffff
const m___INT8_MAX__ = 0x7f
const m___INTMAX_MAX__ = 0x7fffffffffffffff
const m___INTMAX_WIDTH__ = 64
const m___INTPTR_MAX__ = 0x7fffffffffffffff
const m___INTPTR_WIDTH__ = 64
const m___INT_FAST16_MAX__ = 0x7fffffffffffffff
const m___INT_FAST16_WIDTH__ = 64
const m___INT_FAST32_MAX__ = 0x7fffffffffffffff
const m___INT_FAST32_WIDTH__ = 64
const m___INT_FAST64_MAX__ = 0x7fffffffffffffff
const m___INT_FAST64_WIDTH__ = 64
const m___INT_FAST8_MAX__ = 0x7f
const m___INT_FAST8_WIDTH__ = 8
const m___INT_LEAST16_MAX__ = 0x7fff
const m___INT_LEAST16_WIDTH__ = 16
const m___INT_LEAST32_MAX__ = 0x7fffffff
const m___INT_LEAST32_TYPE__ = "int"
const m___INT_LEAST32_WIDTH__ = 32
const m___INT_LEAST64_MAX__ = 0x7fffffffffffffff
const m___INT_LEAST64_WIDTH__ = 64
const m___INT_LEAST8_MAX__ = 0x7f
const m___INT_LEAST8_WIDTH__ = 8
const m___INT_MAX__ = 0x7fffffff
const m___INT_WIDTH__ = 32
const m___LACCUM_EPSILON__ = "0x1P-31LK"
const m___LACCUM_FBIT__ = 31
const m___LACCUM_IBIT__ = 32
const m___LACCUM_MAX__ = "0X7FFFFFFFFFFFFFFFP-31LK"
const m___LDBL_DECIMAL_DIG__ = 36
const m___LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___LDBL_DIG__ = 33
const m___LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___LDBL_HAS_DENORM__ = 1
const m___LDBL_HAS_INFINITY__ = 1
const m___LDBL_HAS_QUIET_NAN__ = 1
const m___LDBL_IS_IEC_60559__ = 1
const m___LDBL_MANT_DIG__ = 113
const m___LDBL_MAX_10_EXP__ = 4932
const m___LDBL_MAX_EXP__ = 16384
const m___LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___LDBL_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___LDOUBLE_REDIRECTS_TO_FLOAT128_ABI = 0
const m___LFRACT_EPSILON__ = "0x1P-31LR"
const m___LFRACT_FBIT__ = 31
const m___LFRACT_IBIT__ = 0
const m___LFRACT_MAX__ = "0X7FFFFFFFP-31LR"
const m___LITTLE_ENDIAN = 1234
const m___LLACCUM_EPSILON__ = "0x1P-63LLK"
const m___LLACCUM_FBIT__ = 63
const m___LLACCUM_IBIT__ = 64
const m___LLACCUM_MAX__ = "0X7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFP-63LLK"
const m___LLFRACT_EPSILON__ = "0x1P-63LLR"
const m___LLFRACT_FBIT__ = 63
const m___LLFRACT_IBIT__ = 0
const m___LLFRACT_MAX__ = "0X7FFFFFFFFFFFFFFFP-63LLR"
const m___LONG_LONG_MAX__ = 0x7fffffffffffffff
const m___LONG_LONG_WIDTH__ = 64
const m___LONG_MAX = 9223372036854775807
const m___LONG_MAX__ = 0x7fffffffffffffff
const m___LONG_WIDTH__ = 64
const m___LP64__ = 1
const m___NO_INLINE__ = 1
const m___ORDER_BIG_ENDIAN__ = 4321
const m___ORDER_LITTLE_ENDIAN__ = 1234
const m___ORDER_PDP_ENDIAN__ = 3412
const m___PDP_ENDIAN = 3412
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTRDIFF_MAX__ = 0x7fffffffffffffff
const m___PTRDIFF_WIDTH__ = 64
const m___QQ_FBIT__ = 7
const m___QQ_IBIT__ = 0
const m___REDIRECT_FORTIFY = "__REDIRECT"
const m___REDIRECT_FORTIFY_NTH = "__REDIRECT_NTH"
const m___REENTRANT = 1
const m___REGISTER_PREFIX__ = "$"
const m___SACCUM_EPSILON__ = "0x1P-7HK"
const m___SACCUM_FBIT__ = 7
const m___SACCUM_IBIT__ = 8
const m___SACCUM_MAX__ = "0X7FFFP-7HK"
const m___SA_FBIT__ = 15
const m___SA_IBIT__ = 16
const m___SCHAR_MAX__ = 0x7f
const m___SCHAR_WIDTH__ = 8
const m___SFRACT_EPSILON__ = "0x1P-7HR"
const m___SFRACT_FBIT__ = 7
const m___SFRACT_IBIT__ = 0
const m___SFRACT_MAX__ = "0X7FP-7HR"
const m___SHRT_MAX__ = 0x7fff
const m___SHRT_WIDTH__ = 16
const m___SIG_ATOMIC_MAX__ = 0x7fffffff
const m___SIG_ATOMIC_TYPE__ = "int"
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
const m___SIZE_MAX__ = 0xffffffffffffffff
const m___SIZE_WIDTH__ = 64
const m___SQ_FBIT__ = 31
const m___SQ_IBIT__ = 0
const m___STDC_HOSTED__ = 1
const m___STDC_IEC_559_COMPLEX__ = 1
const m___STDC_IEC_559__ = 1
const m___STDC_IEC_60559_BFP__ = 201404
const m___STDC_IEC_60559_COMPLEX__ = 201404
const m___STDC_ISO_10646__ = 201706
const m___STDC_UTF_16__ = 1
const m___STDC_UTF_32__ = 1
const m___STDC_VERSION__ = 201710
const m___STDC_WANT_IEC_60559_ATTRIBS_EXT__ = 1
const m___STDC_WANT_IEC_60559_BFP_EXT__ = 1
const m___STDC_WANT_IEC_60559_DFP_EXT__ = 1
const m___STDC_WANT_IEC_60559_FUNCS_EXT__ = 1
const m___STDC_WANT_IEC_60559_TYPES_EXT__ = 1
const m___STDC_WANT_LIB_EXT2__ = 1
const m___STDC_WANT_MATH_SPEC_FUNCS__ = 1
const m___STDC__ = 1
const m___TA_FBIT__ = 63
const m___TA_IBIT__ = 64
const m___TQ_FBIT__ = 127
const m___TQ_IBIT__ = 0
const m___UACCUM_EPSILON__ = "0x1P-16UK"
const m___UACCUM_FBIT__ = 16
const m___UACCUM_IBIT__ = 16
const m___UACCUM_MAX__ = "0XFFFFFFFFP-16UK"
const m___UACCUM_MIN__ = "0.0UK"
const m___UDA_FBIT__ = 32
const m___UDA_IBIT__ = 32
const m___UDQ_FBIT__ = 64
const m___UDQ_IBIT__ = 0
const m___UFRACT_EPSILON__ = "0x1P-16UR"
const m___UFRACT_FBIT__ = 16
const m___UFRACT_IBIT__ = 0
const m___UFRACT_MAX__ = "0XFFFFP-16UR"
const m___UFRACT_MIN__ = "0.0UR"
const m___UHA_FBIT__ = 8
const m___UHA_IBIT__ = 8
const m___UHQ_FBIT__ = 16
const m___UHQ_IBIT__ = 0
const m___UINT16_MAX__ = 0xffff
const m___UINT32_MAX__ = 0xffffffff
const m___UINT64_MAX__ = 0xffffffffffffffff
const m___UINT8_MAX__ = 0xff
const m___UINTMAX_MAX__ = 0xffffffffffffffff
const m___UINTPTR_MAX__ = 0xffffffffffffffff
const m___UINT_FAST16_MAX__ = 0xffffffffffffffff
const m___UINT_FAST32_MAX__ = 0xffffffffffffffff
const m___UINT_FAST64_MAX__ = 0xffffffffffffffff
const m___UINT_FAST8_MAX__ = 0xff
const m___UINT_LEAST16_MAX__ = 0xffff
const m___UINT_LEAST32_MAX__ = 0xffffffff
const m___UINT_LEAST64_MAX__ = 0xffffffffffffffff
const m___UINT_LEAST8_MAX__ = 0xff
const m___ULACCUM_EPSILON__ = "0x1P-32ULK"
const m___ULACCUM_FBIT__ = 32
const m___ULACCUM_IBIT__ = 32
const m___ULACCUM_MAX__ = "0XFFFFFFFFFFFFFFFFP-32ULK"
const m___ULACCUM_MIN__ = "0.0ULK"
const m___ULFRACT_EPSILON__ = "0x1P-32ULR"
const m___ULFRACT_FBIT__ = 32
const m___ULFRACT_IBIT__ = 0
const m___ULFRACT_MAX__ = "0XFFFFFFFFP-32ULR"
const m___ULFRACT_MIN__ = "0.0ULR"
const m___ULLACCUM_EPSILON__ = "0x1P-64ULLK"
const m___ULLACCUM_FBIT__ = 64
const m___ULLACCUM_IBIT__ = 64
const m___ULLACCUM_MAX__ = "0XFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFP-64ULLK"
const m___ULLACCUM_MIN__ = "0.0ULLK"
const m___ULLFRACT_EPSILON__ = "0x1P-64ULLR"
const m___ULLFRACT_FBIT__ = 64
const m___ULLFRACT_IBIT__ = 0
const m___ULLFRACT_MAX__ = "0XFFFFFFFFFFFFFFFFP-64ULLR"
const m___ULLFRACT_MIN__ = "0.0ULLR"
const m___UQQ_FBIT__ = 8
const m___UQQ_IBIT__ = 0
const m___USACCUM_EPSILON__ = "0x1P-8UHK"
const m___USACCUM_FBIT__ = 8
const m___USACCUM_IBIT__ = 8
const m___USACCUM_MAX__ = "0XFFFFP-8UHK"
const m___USACCUM_MIN__ = "0.0UHK"
const m___USA_FBIT__ = 16
const m___USA_IBIT__ = 16
const m___USE_TIME_BITS64 = 1
const m___USFRACT_EPSILON__ = "0x1P-8UHR"
const m___USFRACT_FBIT__ = 8
const m___USFRACT_IBIT__ = 0
const m___USFRACT_MAX__ = "0XFFP-8UHR"
const m___USFRACT_MIN__ = "0.0UHR"
const m___USQ_FBIT__ = 32
const m___USQ_IBIT__ = 0
const m___UTA_FBIT__ = 64
const m___UTA_IBIT__ = 64
const m___UTQ_FBIT__ = 128
const m___UTQ_IBIT__ = 0
const m___VERSION__ = "14.2.1 20241104 (Red Hat 14.2.1-6)"
const m___WCHAR_MAX__ = 0x7fffffff
const m___WCHAR_TYPE__ = "int"
const m___WCHAR_WIDTH__ = 32
const m___WINT_MAX__ = 0xffffffff
const m___WINT_MIN__ = 0
const m___WINT_WIDTH__ = 32
const m___WORDSIZE = 64
const m___WORDSIZE_TIME64_COMPAT32 = 0
const m___builtin_copysignq = "__builtin_copysignf128"
const m___builtin_fabsq = "__builtin_fabsf128"
const m___builtin_huge_valq = "__builtin_huge_valf128"
const m___builtin_infq = "__builtin_inff128"
const m___builtin_nanq = "__builtin_nanf128"
const m___builtin_nansq = "__builtin_nansf128"
const m___glibc_c99_flexarr_available = 1
const m___gnu_linux__ = 1
const m___inline = "inline"
const m___linux = 1
const m___linux__ = 1
const m___loongarch64 = 1
const m___loongarch__ = 1
const m___loongarch_arch = "la64v1.0"
const m___loongarch_double_float = 1
const m___loongarch_frlen = 64
const m___loongarch_grlen = 64
const m___loongarch_hard_float = 1
const m___loongarch_lp64 = 1
const m___loongarch_simd = 1
const m___loongarch_simd_width = 128
const m___loongarch_sx = 1
const m___loongarch_tune = "generic"
const m___loongarch_version_major = 1
const m___loongarch_version_minor = 0
const m___restrict = "restrict"
const m___uc_flags = "uc_flags"
const m___ucontext = "ucontext"
const m___unix = 1
const m___unix__ = 1
const m_alloca = "__builtin_alloca"
const m_linux = 1
const m_unix = 1

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint64

type t__predefined_wchar_t = int32

type t__predefined_ptrdiff_t = int64

type Tnlink_t = uint32

type Tblksize_t = int32

type Tsize_t = uint64

type Tssize_t = int64

type Tregister_t = int64

type Ttime_t = int64

type Tsuseconds_t = int64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tu_int64_t = uint64

type Tmode_t = uint32

type Toff_t = int64

type Tino_t = uint64

type Tdev_t = uint64

type Tblkcnt_t = int64

type Tfsblkcnt_t = uint64

type Tfsfilcnt_t = uint64

type Ttimer_t = uintptr

type Tclockid_t = int32

type Tclock_t = int64

type Tpid_t = int32

type Tid_t = uint32

type Tuid_t = uint32

type Tgid_t = uint32

type Tkey_t = int32

type Tuseconds_t = uint32

type Tpthread_t = uintptr

type Tpthread_once_t = int32

type Tpthread_key_t = uint32

type Tpthread_spinlock_t = int32

type Tpthread_mutexattr_t = struct {
	F__attr uint32
}

type Tpthread_condattr_t = struct {
	F__attr uint32
}

type Tpthread_barrierattr_t = struct {
	F__attr uint32
}

type Tpthread_rwlockattr_t = struct {
	F__attr [2]uint32
}

type Tpthread_attr_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__s  [0][7]uint64
		F__i  [14]int32
	}
}

type Tpthread_mutex_t = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

type Tpthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][6]uintptr
		F__i  [12]int32
	}
}

type Tpthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__p  [0][7]uintptr
		F__i  [14]int32
	}
}

type Tpthread_barrier_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][4]uintptr
		F__i  [8]int32
	}
}

type Tu_int8_t = uint8

type Tu_int16_t = uint16

type Tu_int32_t = uint32

type Tcaddr_t = uintptr

type Tu_char = uint8

type Tu_short = uint16

type Tushort = uint16

type Tu_int = uint32

type Tuint = uint32

type Tu_long = uint64

type Tulong = uint64

type Tquad_t = int64

type Tu_quad_t = uint64

type Tuint16_t = uint16

type Tuint32_t = uint32

type Tuint64_t = uint64

type Ttimeval = struct {
	Ftv_sec  Ttime_t
	Ftv_usec Tsuseconds_t
}

type Ttimespec = struct {
	Ftv_sec  Ttime_t
	Ftv_nsec int64
}

type Tsigset_t = struct {
	F__bits [16]uint64
}

type t__sigset_t = Tsigset_t

type Tfd_mask = uint64

type Tfd_set = struct {
	Ffds_bits [16]uint64
}

type Tstat = struct {
	Fst_dev     Tdev_t
	Fst_ino     Tino_t
	Fst_mode    Tmode_t
	Fst_nlink   Tnlink_t
	Fst_uid     Tuid_t
	Fst_gid     Tgid_t
	Fst_rdev    Tdev_t
	F__pad      uint64
	Fst_size    Toff_t
	Fst_blksize Tblksize_t
	F__pad2     int32
	Fst_blocks  Tblkcnt_t
	Fst_atim    Ttimespec
	Fst_mtim    Ttimespec
	Fst_ctim    Ttimespec
	F__unused   [2]uint32
}

type Tstatx_timestamp = struct {
	Ftv_sec  Tint64_t
	Ftv_nsec Tuint32_t
	F__pad   Tuint32_t
}

type Tstatx = struct {
	Fstx_mask            Tuint32_t
	Fstx_blksize         Tuint32_t
	Fstx_attributes      Tuint64_t
	Fstx_nlink           Tuint32_t
	Fstx_uid             Tuint32_t
	Fstx_gid             Tuint32_t
	Fstx_mode            Tuint16_t
	F__pad0              [1]Tuint16_t
	Fstx_ino             Tuint64_t
	Fstx_size            Tuint64_t
	Fstx_blocks          Tuint64_t
	Fstx_attributes_mask Tuint64_t
	Fstx_atime           Tstatx_timestamp
	Fstx_btime           Tstatx_timestamp
	Fstx_ctime           Tstatx_timestamp
	Fstx_mtime           Tstatx_timestamp
	Fstx_rdev_major      Tuint32_t
	Fstx_rdev_minor      Tuint32_t
	Fstx_dev_major       Tuint32_t
	Fstx_dev_minor       Tuint32_t
	F__pad1              [14]Tuint64_t
}

type Twchar_t = int32

type Tmax_align_t = struct {
	F__ll int64
	F__ld float64
}

type Tptrdiff_t = int64

type Tlocale_t = uintptr

type Tstack_t = struct {
	Fss_sp    uintptr
	Fss_flags int32
	Fss_size  Tsize_t
}

type Tsigaltstack = Tstack_t

type Tgreg_t = uint64

type Tgregset_t = [32]uint64

type Tsigcontext = struct {
	Fsc_pc    uint64
	Fsc_regs  [32]uint64
	Fsc_flags uint32
}

type Tmcontext_t = struct {
	F__pc    uint64
	F__gregs [32]uint64
	F__flags uint32
}

type Tucontext_t = struct {
	Fuc_flags    uint64
	Fuc_link     uintptr
	Fuc_stack    Tstack_t
	Fuc_sigmask  Tsigset_t
	F__uc_pad    int64
	Fuc_mcontext Tmcontext_t
}

type Tucontext = Tucontext_t

type Tsigval = struct {
	Fsival_ptr   [0]uintptr
	Fsival_int   int32
	F__ccgo_pad2 [4]byte
}

type Tsiginfo_t = struct {
	Fsi_signo    int32
	Fsi_errno    int32
	Fsi_code     int32
	F__si_fields struct {
		F__si_common [0]struct {
			F__first struct {
				F__timer [0]struct {
					Fsi_timerid int32
					Fsi_overrun int32
				}
				F__piduid struct {
					Fsi_pid Tpid_t
					Fsi_uid Tuid_t
				}
			}
			F__second struct {
				F__sigchld [0]struct {
					Fsi_status int32
					Fsi_utime  Tclock_t
					Fsi_stime  Tclock_t
				}
				Fsi_value    Tsigval
				F__ccgo_pad2 [16]byte
			}
		}
		F__sigfault [0]struct {
			Fsi_addr     uintptr
			Fsi_addr_lsb int16
			F__first     struct {
				Fsi_pkey    [0]uint32
				F__addr_bnd struct {
					Fsi_lower uintptr
					Fsi_upper uintptr
				}
			}
		}
		F__sigpoll [0]struct {
			Fsi_band int64
			Fsi_fd   int32
		}
		F__sigsys [0]struct {
			Fsi_call_addr uintptr
			Fsi_syscall   int32
			Fsi_arch      uint32
		}
		F__pad [112]int8
	}
}

type Tsigaction = struct {
	F__sa_handler struct {
		Fsa_sigaction [0]uintptr
		Fsa_handler   uintptr
	}
	Fsa_mask     Tsigset_t
	Fsa_flags    int32
	Fsa_restorer uintptr
}

type Tsigevent = struct {
	Fsigev_value  Tsigval
	Fsigev_signo  int32
	Fsigev_notify int32
	F__sev_fields struct {
		Fsigev_notify_thread_id [0]Tpid_t
		F__sev_thread           [0]struct {
			Fsigev_notify_function   uintptr
			Fsigev_notify_attributes uintptr
		}
		F__pad [48]int8
	}
}

type Tsig_t = uintptr

type Tsighandler_t = uintptr

type Tsig_atomic_t = int32

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

type Tuintptr_t = uint64

type Tintptr_t = int64

type Tintmax_t = int64

type Tuint8_t = uint8

type Tuintmax_t = uint64

type Tint_fast8_t = int8

type Tint_fast64_t = int64

type Tint_least8_t = int8

type Tint_least16_t = int16

type Tint_least32_t = int32

type Tint_least64_t = int64

type Tuint_fast8_t = uint8

type Tuint_fast64_t = uint64

type Tuint_least8_t = uint8

type Tuint_least16_t = uint16

type Tuint_least32_t = uint32

type Tuint_least64_t = uint64

type Tint_fast16_t = int32

type Tint_fast32_t = int32

type Tuint_fast16_t = uint32

type Tuint_fast32_t = uint32

type Tva_list = uintptr

type t__isoc_va_list = uintptr

type Tfpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

type T_G_fpos64_t = Tfpos_t

type Tcookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type T_IO_cookie_io_functions_t = Tcookie_io_functions_t

type TBITCMD = struct {
	Fcmd  int8
	Fcmd2 int8
	Fbits Tmode_t
}

type Tbitcmd = TBITCMD

// C documentation
//
//	/*
//	 * Given the old mode and an array of bitcmd structures, apply the operations
//	 * described in the bitcmd structures to the old mode, and return the new mode.
//	 * Note that there is no '=' command; a strict assignment is just a '-' (clear
//	 * bits) followed by a '+' (set bits).
//	 */
func Xgetmode(tls *libc.TLS, bbox uintptr, omode Tmode_t) (r Tmode_t) {
	var clrval, newmode, value Tmode_t
	var set uintptr
	var v11 uint32
	_, _, _, _, _ = clrval, newmode, set, value, v11
	set = bbox
	newmode = omode
	value = uint32(0)
	for {
		switch int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd) {
		case int32('u'):
			goto _2
		case int32('g'):
			goto _3
		case int32('o'):
			goto _4
		case int32('+'):
			goto _5
		case int32('-'):
			goto _6
		case int32('X'):
			goto _7
		default:
			goto _8
		case int32('\000'):
			goto _9
		}
		goto _10
		/*
		 * When copying the user, group or other bits around, we "know"
		 * where the bits are in the mode so that we can do shifts to
		 * copy them around.  If we don't use shifts, it gets real
		 * grundgy with lots of single bit checks and bit sets.
		 */
	_2:
		;
		value = newmode & uint32(m_S_IRWXU) >> int32(6)
		goto common
	_3:
		;
		value = newmode & uint32(m_S_IRWXG) >> int32(3)
		goto common
	_4:
		;
		value = newmode & uint32(m_S_IRWXO)
		goto common
	common:
		;
		if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_CLR) != 0 {
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_SET) != 0 {
				v11 = uint32(m_S_IRWXO)
			} else {
				v11 = value
			}
			clrval = v11
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_UBITS) != 0 {
				newmode &= ^(clrval << libc.Int32FromInt32(6) & (*TBITCMD)(unsafe.Pointer(set)).Fbits)
			}
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_GBITS) != 0 {
				newmode &= ^(clrval << libc.Int32FromInt32(3) & (*TBITCMD)(unsafe.Pointer(set)).Fbits)
			}
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_OBITS) != 0 {
				newmode &= ^(clrval & (*TBITCMD)(unsafe.Pointer(set)).Fbits)
			}
		}
		if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_SET) != 0 {
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_UBITS) != 0 {
				newmode |= value << libc.Int32FromInt32(6) & (*TBITCMD)(unsafe.Pointer(set)).Fbits
			}
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_GBITS) != 0 {
				newmode |= value << libc.Int32FromInt32(3) & (*TBITCMD)(unsafe.Pointer(set)).Fbits
			}
			if int32((*TBITCMD)(unsafe.Pointer(set)).Fcmd2)&int32(m_CMD2_OBITS) != 0 {
				newmode |= value & (*TBITCMD)(unsafe.Pointer(set)).Fbits
			}
		}
		goto _10
	_5:
		;
		newmode |= (*TBITCMD)(unsafe.Pointer(set)).Fbits
		goto _10
	_6:
		;
		newmode &= ^(*TBITCMD)(unsafe.Pointer(set)).Fbits
		goto _10
	_7:
		;
		if omode&libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IFDIR)|libc.Int32FromInt32(m_S_IXUSR)|libc.Int32FromInt32(m_S_IXGRP)|libc.Int32FromInt32(m_S_IXOTH)) != 0 {
			newmode |= (*TBITCMD)(unsafe.Pointer(set)).Fbits
		}
		goto _10
	_9:
		;
	_8:
		;
		return newmode
	_10:
		;
		goto _1
	_1:
		;
		set += 8
	}
	return r
}

func Xsetmode(tls *libc.TLS, p uintptr) (r uintptr) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var endset, newset, newset1, newset2, newset3, newset4, newset5, saveset, set, v5 uintptr
	var equalopdone, serrno, setlen int32
	var lval int64
	var mask, perm, permXbits, who, v1 Tmode_t
	var op, v4 int8
	var setdiff, setdiff1, setdiff2, setdiff3, setdiff4, setdiff5 Tptrdiff_t
	var _ /* ep at bp+0 */ uintptr
	var _ /* signset at bp+8 */ Tsigset_t
	var _ /* sigoset at bp+136 */ Tsigset_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = endset, equalopdone, lval, mask, newset, newset1, newset2, newset3, newset4, newset5, op, perm, permXbits, saveset, serrno, set, setdiff, setdiff1, setdiff2, setdiff3, setdiff4, setdiff5, setlen, who, v1, v4, v5
	equalopdone = 0
	if !(*(*int8)(unsafe.Pointer(p)) != 0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return libc.UintptrFromInt32(0)
	}
	/*
	 * Get a copy of the mask for the permissions that are mask relative.
	 * Flip the bits, we want what's not set.  Since it's possible that
	 * the caller is opening files inside a signal handler, protect them
	 * as best we can.
	 */
	libc.Xsigfillset(tls, bp+8)
	libc.Xsigprocmask(tls, m_SIG_BLOCK, bp+8, bp+136)
	v1 = libc.Xumask(tls, uint32(0))
	mask = v1
	libc.Xumask(tls, v1)
	mask = ^mask
	libc.Xsigprocmask(tls, int32(m_SIG_SETMASK), bp+136, libc.UintptrFromInt32(0))
	setlen = libc.Int32FromInt32(m_SET_LEN) + libc.Int32FromInt32(2)
	set = Xreallocarray(tls, libc.UintptrFromInt32(0), libc.Uint64FromInt32(setlen), uint64(8))
	if set == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	saveset = set
	endset = set + uintptr(setlen-libc.Int32FromInt32(2))*8
	/*
	 * If an absolute number, get it and return; disallow non-octal digits
	 * or illegal bits.
	 */
	if libc.BoolInt32(uint32(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(p))))-uint32('0') < uint32(10)) != 0 {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		lval = libc.Xstrtol(tls, p, bp, int32(8))
		if *(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
			goto out
		}
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_ERANGE) && (lval == int64(0x7fffffffffffffff) || lval == -libc.Int64FromInt64(0x7fffffffffffffff)-libc.Int64FromInt32(1)) {
			goto out
		}
		if lval&int64(^(libc.Int32FromInt32(m_S_ISUID)|libc.Int32FromInt32(m_S_ISGID)|libc.Int32FromInt32(m_S_IRWXU)|libc.Int32FromInt32(m_S_IRWXG)|libc.Int32FromInt32(m_S_IRWXO)|libc.Int32FromInt32(m_S_ISVTX))) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
			goto out
		}
		perm = libc.Uint32FromInt64(lval)
		if set >= endset {
			setdiff = (int64(set) - int64(saveset)) / 8
			setlen += int32(m_SET_LEN_INCR)
			newset = Xreallocarray(tls, saveset, libc.Uint64FromInt32(setlen), uint64(8))
			if newset == libc.UintptrFromInt32(0) {
				goto out
			}
			set = newset + uintptr(setdiff)*8
			saveset = newset
			endset = newset + uintptr(setlen-libc.Int32FromInt32(2))*8
		}
		set = _addcmd(tls, set, libc.Uint32FromInt32(libc.Int32FromUint8('=')), libc.Uint32FromInt32(libc.Int32FromInt32(m_S_ISUID)|libc.Int32FromInt32(m_S_ISGID)|libc.Int32FromInt32(m_S_IRWXU)|libc.Int32FromInt32(m_S_IRWXG)|libc.Int32FromInt32(m_S_IRWXO)|libc.Int32FromInt32(m_S_ISVTX)), perm, mask)
		(*TBITCMD)(unsafe.Pointer(set)).Fcmd = 0
		return saveset
	}
	/*
	 * Build list of structures to set/clear/copy bits as described by
	 * each clause of the symbolic mode.
	 */
	for {
		/* First, find out which bits might be modified. */
		who = uint32(0)
		for {
			switch int32(*(*int8)(unsafe.Pointer(p))) {
			case int32('a'):
				who |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_ISUID) | libc.Int32FromInt32(m_S_ISGID) | libc.Int32FromInt32(m_S_IRWXU) | libc.Int32FromInt32(m_S_IRWXG) | libc.Int32FromInt32(m_S_IRWXO))
			case int32('u'):
				who |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_ISUID) | libc.Int32FromInt32(m_S_IRWXU))
			case int32('g'):
				who |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_ISGID) | libc.Int32FromInt32(m_S_IRWXG))
			case int32('o'):
				who |= uint32(m_S_IRWXO)
			default:
				goto getop
			}
			goto _3
		_3:
			;
			p++
		}
		goto getop
	getop:
		;
		v5 = p
		p++
		v4 = *(*int8)(unsafe.Pointer(v5))
		op = v4
		if int32(v4) != int32('+') && int32(op) != int32('-') && int32(op) != int32('=') {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
			goto out
		}
		if int32(op) == int32('=') {
			equalopdone = 0
		}
		who &= libc.Uint32FromInt32(^libc.Int32FromInt32(m_S_ISVTX))
		perm = uint32(0)
		permXbits = libc.Uint32FromInt32(0)
		for {
			switch int32(*(*int8)(unsafe.Pointer(p))) {
			case int32('r'):
				perm |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IRUSR) | libc.Int32FromInt32(m_S_IRGRP) | libc.Int32FromInt32(m_S_IROTH))
			case int32('s'):
				/*
				 * If specific bits where requested and
				 * only "other" bits ignore set-id.
				 */
				if who == uint32(0) || who&libc.Uint32FromInt32(^libc.Int32FromInt32(m_S_IRWXO)) != 0 {
					perm |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_ISUID) | libc.Int32FromInt32(m_S_ISGID))
				}
			case int32('t'):
				/*
				 * If specific bits where requested and
				 * only "other" bits ignore set-id.
				 */
				if who == uint32(0) || who&libc.Uint32FromInt32(^libc.Int32FromInt32(m_S_IRWXO)) != 0 {
					who |= uint32(m_S_ISVTX)
					perm |= uint32(m_S_ISVTX)
				}
			case int32('w'):
				perm |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IWUSR) | libc.Int32FromInt32(m_S_IWGRP) | libc.Int32FromInt32(m_S_IWOTH))
			case int32('X'):
				permXbits = libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXUSR) | libc.Int32FromInt32(m_S_IXGRP) | libc.Int32FromInt32(m_S_IXOTH))
			case int32('x'):
				perm |= libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXUSR) | libc.Int32FromInt32(m_S_IXGRP) | libc.Int32FromInt32(m_S_IXOTH))
			case int32('u'):
				fallthrough
			case int32('g'):
				fallthrough
			case int32('o'):
				/*
				 * When ever we hit 'u', 'g', or 'o', we have
				 * to flush out any partial mode that we have,
				 * and then do the copying of the mode bits.
				 */
				if perm != 0 {
					if set >= endset {
						setdiff1 = (int64(set) - int64(saveset)) / 8
						setlen += int32(m_SET_LEN_INCR)
						newset1 = Xreallocarray(tls, saveset, libc.Uint64FromInt32(setlen), uint64(8))
						if newset1 == libc.UintptrFromInt32(0) {
							goto out
						}
						set = newset1 + uintptr(setdiff1)*8
						saveset = newset1
						endset = newset1 + uintptr(setlen-libc.Int32FromInt32(2))*8
					}
					set = _addcmd(tls, set, libc.Uint32FromInt8(op), who, perm, mask)
					perm = uint32(0)
				}
				if int32(op) == int32('=') {
					equalopdone = int32(1)
				}
				if int32(op) == int32('+') && permXbits != 0 {
					if set >= endset {
						setdiff2 = (int64(set) - int64(saveset)) / 8
						setlen += int32(m_SET_LEN_INCR)
						newset2 = Xreallocarray(tls, saveset, libc.Uint64FromInt32(setlen), uint64(8))
						if newset2 == libc.UintptrFromInt32(0) {
							goto out
						}
						set = newset2 + uintptr(setdiff2)*8
						saveset = newset2
						endset = newset2 + uintptr(setlen-libc.Int32FromInt32(2))*8
					}
					set = _addcmd(tls, set, libc.Uint32FromInt32(libc.Int32FromUint8('X')), who, permXbits, mask)
					permXbits = uint32(0)
				}
				if set >= endset {
					setdiff3 = (int64(set) - int64(saveset)) / 8
					setlen += int32(m_SET_LEN_INCR)
					newset3 = Xreallocarray(tls, saveset, libc.Uint64FromInt32(setlen), uint64(8))
					if newset3 == libc.UintptrFromInt32(0) {
						goto out
					}
					set = newset3 + uintptr(setdiff3)*8
					saveset = newset3
					endset = newset3 + uintptr(setlen-libc.Int32FromInt32(2))*8
				}
				set = _addcmd(tls, set, libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(p))), who, libc.Uint32FromInt8(op), mask)
			default:
				/*
				 * Add any permissions that we haven't already
				 * done.
				 */
				if perm != 0 || int32(op) == int32('=') && !(equalopdone != 0) {
					if int32(op) == int32('=') {
						equalopdone = int32(1)
					}
					if set >= endset {
						setdiff4 = (int64(set) - int64(saveset)) / 8
						setlen += int32(m_SET_LEN_INCR)
						newset4 = Xreallocarray(tls, saveset, libc.Uint64FromInt32(setlen), uint64(8))
						if newset4 == libc.UintptrFromInt32(0) {
							goto out
						}
						set = newset4 + uintptr(setdiff4)*8
						saveset = newset4
						endset = newset4 + uintptr(setlen-libc.Int32FromInt32(2))*8
					}
					set = _addcmd(tls, set, libc.Uint32FromInt8(op), who, perm, mask)
					perm = uint32(0)
				}
				if permXbits != 0 {
					if set >= endset {
						setdiff5 = (int64(set) - int64(saveset)) / 8
						setlen += int32(m_SET_LEN_INCR)
						newset5 = Xreallocarray(tls, saveset, libc.Uint64FromInt32(setlen), uint64(8))
						if newset5 == libc.UintptrFromInt32(0) {
							goto out
						}
						set = newset5 + uintptr(setdiff5)*8
						saveset = newset5
						endset = newset5 + uintptr(setlen-libc.Int32FromInt32(2))*8
					}
					set = _addcmd(tls, set, libc.Uint32FromInt32(libc.Int32FromUint8('X')), who, permXbits, mask)
					permXbits = uint32(0)
				}
				goto apply
			}
			goto _6
		_6:
			;
			p++
		}
		goto apply
	apply:
		;
		if !(*(*int8)(unsafe.Pointer(p)) != 0) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(p))) != int32(',') {
			goto getop
		}
		p++
		goto _2
	_2:
	}
	(*TBITCMD)(unsafe.Pointer(set)).Fcmd = 0
	_compress_mode(tls, saveset)
	return saveset
	goto out
out:
	;
	serrno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xfree(tls, saveset)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = serrno
	return libc.UintptrFromInt32(0)
}

func _addcmd(tls *libc.TLS, set uintptr, op Tmode_t, who Tmode_t, oparg Tmode_t, mask Tmode_t) (r uintptr) {
	var v1, v2 uint32
	var v3, v4, v5 int32
	var p6, p7, p8 uintptr
	_, _, _, _, _, _, _, _ = v1, v2, v3, v4, v5, p6, p7, p8
	switch op {
	case uint32('='):
		(*TBITCMD)(unsafe.Pointer(set)).Fcmd = int8('-')
		if who != 0 {
			v1 = who
		} else {
			v1 = libc.Uint32FromInt32(libc.Int32FromInt32(m_S_ISUID) | libc.Int32FromInt32(m_S_ISGID) | libc.Int32FromInt32(m_S_IRWXU) | libc.Int32FromInt32(m_S_IRWXG) | libc.Int32FromInt32(m_S_IRWXO))
		}
		(*TBITCMD)(unsafe.Pointer(set)).Fbits = v1
		set += 8
		op = uint32('+')
		/* FALLTHROUGH */
		fallthrough
	case uint32('+'):
		fallthrough
	case uint32('-'):
		fallthrough
	case uint32('X'):
		(*TBITCMD)(unsafe.Pointer(set)).Fcmd = libc.Int8FromUint32(op)
		if who != 0 {
			v2 = who
		} else {
			v2 = mask
		}
		(*TBITCMD)(unsafe.Pointer(set)).Fbits = v2 & oparg
	case uint32('u'):
		fallthrough
	case uint32('g'):
		fallthrough
	case uint32('o'):
		(*TBITCMD)(unsafe.Pointer(set)).Fcmd = libc.Int8FromUint32(op)
		if who != 0 {
			if who&uint32(m_S_IRUSR) != 0 {
				v3 = int32(m_CMD2_UBITS)
			} else {
				v3 = 0
			}
			if who&uint32(m_S_IRGRP) != 0 {
				v4 = int32(m_CMD2_GBITS)
			} else {
				v4 = 0
			}
			if who&uint32(m_S_IROTH) != 0 {
				v5 = int32(m_CMD2_OBITS)
			} else {
				v5 = 0
			}
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd2 = int8(v3 | v4 | v5)
			(*TBITCMD)(unsafe.Pointer(set)).Fbits = libc.Uint32FromInt32(^libc.Int32FromInt32(0))
		} else {
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd2 = int8(libc.Int32FromInt32(m_CMD2_UBITS) | libc.Int32FromInt32(m_CMD2_GBITS) | libc.Int32FromInt32(m_CMD2_OBITS))
			(*TBITCMD)(unsafe.Pointer(set)).Fbits = mask
		}
		if oparg == uint32('+') {
			p6 = set + 1
			*(*int8)(unsafe.Pointer(p6)) = int8(int32(*(*int8)(unsafe.Pointer(p6))) | libc.Int32FromInt32(m_CMD2_SET))
		} else {
			if oparg == uint32('-') {
				p7 = set + 1
				*(*int8)(unsafe.Pointer(p7)) = int8(int32(*(*int8)(unsafe.Pointer(p7))) | libc.Int32FromInt32(m_CMD2_CLR))
			} else {
				if oparg == uint32('=') {
					p8 = set + 1
					*(*int8)(unsafe.Pointer(p8)) = int8(int32(*(*int8)(unsafe.Pointer(p8))) | (libc.Int32FromInt32(m_CMD2_SET) | libc.Int32FromInt32(m_CMD2_CLR)))
				}
			}
		}
		break
	}
	return set + libc.UintptrFromInt32(1)*8
}

// C documentation
//
//	/*
//	 * Given an array of bitcmd structures, compress by compacting consecutive
//	 * '+', '-' and 'X' commands into at most 3 commands, one of each.  The 'u',
//	 * 'g' and 'o' commands continue to be separate.  They could probably be
//	 * compacted, but it's not worth the effort.
//	 */
func _compress_mode(tls *libc.TLS, set uintptr) {
	var Xbits, clrbits, op, setbits, v2, v6, v7, v8 int32
	var nset, v3, v4 uintptr
	_, _, _, _, _, _, _, _, _, _, _ = Xbits, clrbits, nset, op, setbits, v2, v3, v4, v6, v7, v8
	nset = set
	for {
		/* Copy over any 'u', 'g' and 'o' commands. */
		for {
			v2 = int32((*TBITCMD)(unsafe.Pointer(nset)).Fcmd)
			op = v2
			if !(v2 != int32('+') && op != int32('-') && op != int32('X')) {
				break
			}
			v3 = set
			set += 8
			v4 = nset
			nset += 8
			*(*TBITCMD)(unsafe.Pointer(v3)) = *(*TBITCMD)(unsafe.Pointer(v4))
			if !(op != 0) {
				return
			}
		}
		v7 = libc.Int32FromInt32(0)
		Xbits = v7
		v6 = v7
		clrbits = v6
		setbits = v6
		for {
			v8 = int32((*TBITCMD)(unsafe.Pointer(nset)).Fcmd)
			op = v8
			if v8 == int32('-') {
				clrbits = int32(uint32(clrbits) | (*TBITCMD)(unsafe.Pointer(nset)).Fbits)
				setbits = int32(uint32(setbits) & ^(*TBITCMD)(unsafe.Pointer(nset)).Fbits)
				Xbits = int32(uint32(Xbits) & ^(*TBITCMD)(unsafe.Pointer(nset)).Fbits)
			} else {
				if op == int32('+') {
					setbits = int32(uint32(setbits) | (*TBITCMD)(unsafe.Pointer(nset)).Fbits)
					clrbits = int32(uint32(clrbits) & ^(*TBITCMD)(unsafe.Pointer(nset)).Fbits)
					Xbits = int32(uint32(Xbits) & ^(*TBITCMD)(unsafe.Pointer(nset)).Fbits)
				} else {
					if op == int32('X') {
						Xbits = int32(uint32(Xbits) | (*TBITCMD)(unsafe.Pointer(nset)).Fbits&libc.Uint32FromInt32(^setbits))
					} else {
						break
					}
				}
			}
			goto _5
		_5:
			;
			nset += 8
		}
		if clrbits != 0 {
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd = int8('-')
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd2 = 0
			(*TBITCMD)(unsafe.Pointer(set)).Fbits = libc.Uint32FromInt32(clrbits)
			set += 8
		}
		if setbits != 0 {
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd = int8('+')
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd2 = 0
			(*TBITCMD)(unsafe.Pointer(set)).Fbits = libc.Uint32FromInt32(setbits)
			set += 8
		}
		if Xbits != 0 {
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd = int8('X')
			(*TBITCMD)(unsafe.Pointer(set)).Fcmd2 = 0
			(*TBITCMD)(unsafe.Pointer(set)).Fbits = libc.Uint32FromInt32(Xbits)
			set += 8
		}
		goto _1
	_1:
	}
}

const m_AT_EACCESS = 0x200
const m_AT_EMPTY_PATH = 0x1000
const m_AT_NO_AUTOMOUNT = 0x800
const m_AT_RECURSIVE = 0x8000
const m_AT_REMOVEDIR = 0x200
const m_AT_STATX_DONT_SYNC = 0x4000
const m_AT_STATX_FORCE_SYNC = 0x2000
const m_AT_STATX_SYNC_AS_STAT = 0x0000
const m_AT_STATX_SYNC_TYPE = 0x6000
const m_AT_SYMLINK_FOLLOW = 0x400
const m_AT_SYMLINK_NOFOLLOW = 0x100
const m_BLOCKSZ = 64
const m_CANBSIZ = 255
const m_CLOCKS_PER_SEC = 1000000
const m_CLOCK_BOOTTIME = 7
const m_CLOCK_BOOTTIME_ALARM = 9
const m_CLOCK_MONOTONIC = 1
const m_CLOCK_MONOTONIC_COARSE = 6
const m_CLOCK_MONOTONIC_RAW = 4
const m_CLOCK_PROCESS_CPUTIME_ID = 2
const m_CLOCK_REALTIME = 0
const m_CLOCK_REALTIME_ALARM = 8
const m_CLOCK_REALTIME_COARSE = 5
const m_CLOCK_SGI_CYCLE = 10
const m_CLOCK_TAI = 11
const m_CLOCK_THREAD_CPUTIME_ID = 3
const m_CLONE_CHILD_CLEARTID = 0x00200000
const m_CLONE_CHILD_SETTID = 0x01000000
const m_CLONE_DETACHED = 0x00400000
const m_CLONE_FILES = 0x00000400
const m_CLONE_FS = 0x00000200
const m_CLONE_IO = 0x80000000
const m_CLONE_NEWCGROUP = 0x02000000
const m_CLONE_NEWIPC = 0x08000000
const m_CLONE_NEWNET = 0x40000000
const m_CLONE_NEWNS = 0x00020000
const m_CLONE_NEWPID = 0x20000000
const m_CLONE_NEWTIME = 0x00000080
const m_CLONE_NEWUSER = 0x10000000
const m_CLONE_NEWUTS = 0x04000000
const m_CLONE_PARENT = 0x00008000
const m_CLONE_PARENT_SETTID = 0x00100000
const m_CLONE_PIDFD = 0x00001000
const m_CLONE_PTRACE = 0x00002000
const m_CLONE_SETTLS = 0x00080000
const m_CLONE_SIGHAND = 0x00000800
const m_CLONE_SYSVSEM = 0x00040000
const m_CLONE_THREAD = 0x00010000
const m_CLONE_UNTRACED = 0x00800000
const m_CLONE_VFORK = 0x00004000
const m_CLONE_VM = 0x00000100
const m_CPU_SETSIZE = 1024
const m_CSIGNAL = 0x000000ff
const m_DEV_BSIZE = 512
const m_DN_ACCESS = 0x00000001
const m_DN_ATTRIB = 0x00000020
const m_DN_CREATE = 0x00000004
const m_DN_DELETE = 0x00000008
const m_DN_MODIFY = 0x00000002
const m_DN_MULTISHOT = 0x80000000
const m_DN_RENAME = 0x00000010
const m_FALLOC_FL_KEEP_SIZE = 1
const m_FALLOC_FL_PUNCH_HOLE = 2
const m_FAPPEND = "O_APPEND"
const m_FASYNC = "O_ASYNC"
const m_FD_CLOEXEC = 1
const m_FFSYNC = "O_SYNC"
const m_FNDELAY = "O_NDELAY"
const m_FNONBLOCK = "O_NONBLOCK"
const m_F_ADD_SEALS = 1033
const m_F_CANCELLK = 1029
const m_F_DUPFD = 0
const m_F_DUPFD_CLOEXEC = 1030
const m_F_GETFD = 1
const m_F_GETFL = 3
const m_F_GETLEASE = 1025
const m_F_GETLK = 5
const m_F_GETOWN = 9
const m_F_GETOWNER_UIDS = 17
const m_F_GETOWN_EX = 16
const m_F_GETPIPE_SZ = 1032
const m_F_GETSIG = 11
const m_F_GET_FILE_RW_HINT = 1037
const m_F_GET_RW_HINT = 1035
const m_F_GET_SEALS = 1034
const m_F_NOTIFY = 1026
const m_F_OFD_GETLK = 36
const m_F_OFD_SETLK = 37
const m_F_OFD_SETLKW = 38
const m_F_OWNER_GID = 2
const m_F_OWNER_PGRP = 2
const m_F_OWNER_PID = 1
const m_F_OWNER_TID = 0
const m_F_RDLCK = 0
const m_F_SEAL_FUTURE_WRITE = 0x0010
const m_F_SEAL_GROW = 0x0004
const m_F_SEAL_SEAL = 0x0001
const m_F_SEAL_SHRINK = 0x0002
const m_F_SEAL_WRITE = 0x0008
const m_F_SETFD = 2
const m_F_SETFL = 4
const m_F_SETLEASE = 1024
const m_F_SETLK = 6
const m_F_SETLKW = 7
const m_F_SETOWN = 8
const m_F_SETOWN_EX = 15
const m_F_SETPIPE_SZ = 1031
const m_F_SETSIG = 10
const m_F_SET_FILE_RW_HINT = 1038
const m_F_SET_RW_HINT = 1036
const m_F_UNLCK = 2
const m_F_WRLCK = 1
const m_ITIMER_PROF = 2
const m_ITIMER_REAL = 0
const m_ITIMER_VIRTUAL = 1
const m_IVSZ = 8
const m_KEYSZ = 32
const m_MADV_COLD = 20
const m_MADV_DODUMP = 17
const m_MADV_DOFORK = 11
const m_MADV_DONTDUMP = 16
const m_MADV_DONTFORK = 10
const m_MADV_DONTNEED = 4
const m_MADV_FREE = 8
const m_MADV_HUGEPAGE = 14
const m_MADV_HWPOISON = 100
const m_MADV_KEEPONFORK = 19
const m_MADV_MERGEABLE = 12
const m_MADV_NOHUGEPAGE = 15
const m_MADV_NORMAL = 0
const m_MADV_PAGEOUT = 21
const m_MADV_RANDOM = 1
const m_MADV_REMOVE = 9
const m_MADV_SEQUENTIAL = 2
const m_MADV_SOFT_OFFLINE = 101
const m_MADV_UNMERGEABLE = 13
const m_MADV_WILLNEED = 3
const m_MADV_WIPEONFORK = 18
const m_MAP_ANON = 32
const m_MAP_ANONYMOUS = "MAP_ANON"
const m_MAP_DENYWRITE = 0x0800
const m_MAP_EXECUTABLE = 0x1000
const m_MAP_FILE = 0
const m_MAP_FIXED = 0x10
const m_MAP_FIXED_NOREPLACE = 0x100000
const m_MAP_GROWSDOWN = 0x0100
const m_MAP_HUGETLB = 0x40000
const m_MAP_HUGE_MASK = 0x3f
const m_MAP_HUGE_SHIFT = 26
const m_MAP_LOCKED = 0x2000
const m_MAP_NONBLOCK = 0x10000
const m_MAP_NORESERVE = 0x4000
const m_MAP_POPULATE = 0x8000
const m_MAP_PRIVATE = 2
const m_MAP_SHARED = 0x01
const m_MAP_SHARED_VALIDATE = 0x03
const m_MAP_STACK = 0x20000
const m_MAP_SYNC = 0x80000
const m_MAP_TYPE = 0x0f
const m_MAXHOSTNAMELEN = 64
const m_MAXNAMLEN = 255
const m_MAXPATHLEN = 4096
const m_MAXSYMLINKS = 20
const m_MAX_HANDLE_SZ = 128
const m_MCL_CURRENT = 1
const m_MCL_FUTURE = 2
const m_MCL_ONFAULT = 4
const m_MFD_ALLOW_SEALING = 0x0002
const m_MFD_CLOEXEC = 0x0001
const m_MFD_HUGETLB = 0x0004
const m_MLOCK_ONFAULT = 0x01
const m_MREMAP_DONTUNMAP = 4
const m_MREMAP_FIXED = 2
const m_MREMAP_MAYMOVE = 1
const m_MS_ASYNC = 1
const m_MS_INVALIDATE = 2
const m_MS_SYNC = 4
const m_NBBY = 8
const m_NCARGS = 131072
const m_NGROUPS = 32
const m_NOFILE = 256
const m_O_APPEND = 02000
const m_O_ASYNC = 020000
const m_O_CLOEXEC = 02000000
const m_O_CREAT = 0100
const m_O_DIRECT = 040000
const m_O_DIRECTORY = 0200000
const m_O_DSYNC = 010000
const m_O_EXCL = 0200
const m_O_EXEC = "O_PATH"
const m_O_LARGEFILE = 0100000
const m_O_NDELAY = "O_NONBLOCK"
const m_O_NOATIME = 01000000
const m_O_NOCTTY = 0400
const m_O_NOFOLLOW = 0400000
const m_O_NONBLOCK = 04000
const m_O_PATH = 010000000
const m_O_RDONLY = 00
const m_O_RDWR = 02
const m_O_RSYNC = 04010000
const m_O_SEARCH = "O_PATH"
const m_O_SYNC = 04010000
const m_O_TMPFILE = 020200000
const m_O_TRUNC = 01000
const m_O_TTY_INIT = 0
const m_O_WRONLY = 01
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
const m_POSIX_MADV_DONTNEED = 4
const m_POSIX_MADV_NORMAL = 0
const m_POSIX_MADV_RANDOM = 1
const m_POSIX_MADV_SEQUENTIAL = 2
const m_POSIX_MADV_WILLNEED = 3
const m_PRIO_MAX = 20
const m_PRIO_PGRP = 1
const m_PRIO_PROCESS = 0
const m_PRIO_USER = 2
const m_PROT_EXEC = 4
const m_PROT_GROWSDOWN = 0x01000000
const m_PROT_GROWSUP = 0x02000000
const m_PROT_NONE = 0
const m_PROT_READ = 1
const m_PROT_WRITE = 2
const m_PTHREAD_CANCEL_ASYNCHRONOUS = 1
const m_PTHREAD_CANCEL_DEFERRED = 0
const m_PTHREAD_CANCEL_DISABLE = 1
const m_PTHREAD_CANCEL_ENABLE = 0
const m_PTHREAD_CANCEL_MASKED = 2
const m_PTHREAD_CREATE_DETACHED = 1
const m_PTHREAD_CREATE_JOINABLE = 0
const m_PTHREAD_EXPLICIT_SCHED = 1
const m_PTHREAD_INHERIT_SCHED = 0
const m_PTHREAD_MUTEX_DEFAULT = 0
const m_PTHREAD_MUTEX_ERRORCHECK = 2
const m_PTHREAD_MUTEX_NORMAL = 0
const m_PTHREAD_MUTEX_RECURSIVE = 1
const m_PTHREAD_MUTEX_ROBUST = 1
const m_PTHREAD_MUTEX_STALLED = 0
const m_PTHREAD_ONCE_INIT = 0
const m_PTHREAD_PRIO_INHERIT = 1
const m_PTHREAD_PRIO_NONE = 0
const m_PTHREAD_PRIO_PROTECT = 2
const m_PTHREAD_PROCESS_PRIVATE = 0
const m_PTHREAD_PROCESS_SHARED = 1
const m_PTHREAD_SCOPE_PROCESS = 1
const m_PTHREAD_SCOPE_SYSTEM = 0
const m_RLIMIT_AS = 9
const m_RLIMIT_CORE = 4
const m_RLIMIT_CPU = 0
const m_RLIMIT_DATA = 2
const m_RLIMIT_FSIZE = 1
const m_RLIMIT_LOCKS = 10
const m_RLIMIT_MEMLOCK = 8
const m_RLIMIT_MSGQUEUE = 12
const m_RLIMIT_NICE = 13
const m_RLIMIT_NLIMITS = 16
const m_RLIMIT_NOFILE = 7
const m_RLIMIT_NPROC = 6
const m_RLIMIT_RSS = 5
const m_RLIMIT_RTPRIO = 14
const m_RLIMIT_RTTIME = 15
const m_RLIMIT_SIGPENDING = 11
const m_RLIMIT_STACK = 3
const m_RLIM_NLIMITS = "RLIMIT_NLIMITS"
const m_RLIM_SAVED_CUR = "RLIM_INFINITY"
const m_RLIM_SAVED_MAX = "RLIM_INFINITY"
const m_RUSAGE_SELF = 0
const m_RUSAGE_THREAD = 1
const m_RWF_WRITE_LIFE_NOT_SET = 0
const m_RWH_WRITE_LIFE_EXTREME = 5
const m_RWH_WRITE_LIFE_LONG = 4
const m_RWH_WRITE_LIFE_MEDIUM = 3
const m_RWH_WRITE_LIFE_NONE = 1
const m_RWH_WRITE_LIFE_SHORT = 2
const m_SCHED_BATCH = 3
const m_SCHED_DEADLINE = 6
const m_SCHED_FIFO = 1
const m_SCHED_IDLE = 5
const m_SCHED_OTHER = 0
const m_SCHED_RESET_ON_FORK = 0x40000000
const m_SCHED_RR = 2
const m_SPLICE_F_GIFT = 8
const m_SPLICE_F_MORE = 4
const m_SPLICE_F_MOVE = 1
const m_SPLICE_F_NONBLOCK = 2
const m_SYNC_FILE_RANGE_WAIT_AFTER = 4
const m_SYNC_FILE_RANGE_WAIT_BEFORE = 1
const m_SYNC_FILE_RANGE_WRITE = 2
const m_S_IFDIR1 = 0040000
const m_S_IRGRP1 = 0040
const m_S_IROTH1 = 0004
const m_S_IRUSR1 = 0400
const m_S_IRWXG1 = 0070
const m_S_IRWXO1 = 0007
const m_S_IRWXU1 = 0700
const m_S_ISGID1 = 02000
const m_S_ISUID1 = 04000
const m_S_ISVTX1 = 01000
const m_S_IWGRP1 = 0020
const m_S_IWOTH1 = 0002
const m_S_IWUSR1 = 0200
const m_S_IXGRP1 = 0010
const m_S_IXOTH1 = 0001
const m_S_IXUSR1 = 0100
const m_TIMER_ABSTIME = 1
const m_TIME_UTC = 1
const m___LONG_MAX1 = 0x7fffffffffffffff
const m___tm_gmtoff = "tm_gmtoff"
const m___tm_zone = "tm_zone"
const m_inline = "__inline"
const m_loff_t = "off_t"
const m_prlimit64 = "prlimit"

type Tiovec = struct {
	Fiov_base uintptr
	Fiov_len  Tsize_t
}

type Tflock = struct {
	Fl_type   int16
	Fl_whence int16
	Fl_start  Toff_t
	Fl_len    Toff_t
	Fl_pid    Tpid_t
}

type Tfile_handle = struct {
	Fhandle_bytes uint32
	Fhandle_type  int32
}

type Tf_owner_ex = struct {
	Ftype1 int32
	Fpid   Tpid_t
}

type Titimerval = struct {
	Fit_interval Ttimeval
	Fit_value    Ttimeval
}

type Ttimezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Tu8 = uint8

type Tu32 = uint32

type Tchacha_ctx = struct {
	Finput [16]Tu32
}

var _sigma = [16]int8{'e', 'x', 'p', 'a', 'n', 'd', ' ', '3', '2', '-', 'b', 'y', 't', 'e', ' ', 'k'}
var _tau = [16]int8{'e', 'x', 'p', 'a', 'n', 'd', ' ', '1', '6', '-', 'b', 'y', 't', 'e', ' ', 'k'}

func _chacha_keysetup(tls *libc.TLS, x uintptr, k uintptr, kbits Tu32) {
	var constants uintptr
	_ = constants
	*(*Tu32)(unsafe.Pointer(x + 4*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 5*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 6*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 7*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12) + 3)))<<libc.Int32FromInt32(24)
	if kbits == uint32(256) {
		k += uintptr(16)
		constants = uintptr(unsafe.Pointer(&_sigma))
	} else {
		constants = uintptr(unsafe.Pointer(&_tau))
	}
	*(*Tu32)(unsafe.Pointer(x + 8*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(0) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 9*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(4) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 10*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(8) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 11*4)) = uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12)))) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(k + libc.UintptrFromInt32(12) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x)) = libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(0)))) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(0) + 1)))<<libc.Int32FromInt32(8) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(0) + 2)))<<libc.Int32FromInt32(16) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(0) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 1*4)) = libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(4)))) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(4) + 1)))<<libc.Int32FromInt32(8) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(4) + 2)))<<libc.Int32FromInt32(16) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(4) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 2*4)) = libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(8)))) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(8) + 1)))<<libc.Int32FromInt32(8) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(8) + 2)))<<libc.Int32FromInt32(16) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(8) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 3*4)) = libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(12)))) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(12) + 1)))<<libc.Int32FromInt32(8) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(12) + 2)))<<libc.Int32FromInt32(16) | libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(constants + libc.UintptrFromInt32(12) + 3)))<<libc.Int32FromInt32(24)
}

func _chacha_ivsetup(tls *libc.TLS, x uintptr, iv uintptr) {
	*(*Tu32)(unsafe.Pointer(x + 12*4)) = uint32(0)
	*(*Tu32)(unsafe.Pointer(x + 13*4)) = uint32(0)
	*(*Tu32)(unsafe.Pointer(x + 14*4)) = uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(0)))) | uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(0) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(0) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(0) + 3)))<<libc.Int32FromInt32(24)
	*(*Tu32)(unsafe.Pointer(x + 15*4)) = uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(4)))) | uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(4) + 1)))<<libc.Int32FromInt32(8) | uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(4) + 2)))<<libc.Int32FromInt32(16) | uint32(*(*Tu8)(unsafe.Pointer(iv + libc.UintptrFromInt32(4) + 3)))<<libc.Int32FromInt32(24)
}

func _chacha_encrypt_bytes(tls *libc.TLS, x uintptr, m uintptr, c uintptr, bytes Tu32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var ctarget uintptr
	var i uint32
	var j0, j1, j10, j11, j12, j13, j14, j15, j2, j3, j4, j5, j6, j7, j8, j9, x0, x1, x10, x11, x12, x13, x14, x15, x2, x3, x4, x5, x6, x7, x8, x9 Tu32
	var _ /* tmp at bp+0 */ [64]Tu8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = ctarget, i, j0, j1, j10, j11, j12, j13, j14, j15, j2, j3, j4, j5, j6, j7, j8, j9, x0, x1, x10, x11, x12, x13, x14, x15, x2, x3, x4, x5, x6, x7, x8, x9
	ctarget = libc.UintptrFromInt32(0)
	if !(bytes != 0) {
		return
	}
	j0 = *(*Tu32)(unsafe.Pointer(x))
	j1 = *(*Tu32)(unsafe.Pointer(x + 1*4))
	j2 = *(*Tu32)(unsafe.Pointer(x + 2*4))
	j3 = *(*Tu32)(unsafe.Pointer(x + 3*4))
	j4 = *(*Tu32)(unsafe.Pointer(x + 4*4))
	j5 = *(*Tu32)(unsafe.Pointer(x + 5*4))
	j6 = *(*Tu32)(unsafe.Pointer(x + 6*4))
	j7 = *(*Tu32)(unsafe.Pointer(x + 7*4))
	j8 = *(*Tu32)(unsafe.Pointer(x + 8*4))
	j9 = *(*Tu32)(unsafe.Pointer(x + 9*4))
	j10 = *(*Tu32)(unsafe.Pointer(x + 10*4))
	j11 = *(*Tu32)(unsafe.Pointer(x + 11*4))
	j12 = *(*Tu32)(unsafe.Pointer(x + 12*4))
	j13 = *(*Tu32)(unsafe.Pointer(x + 13*4))
	j14 = *(*Tu32)(unsafe.Pointer(x + 14*4))
	j15 = *(*Tu32)(unsafe.Pointer(x + 15*4))
	for {
		if bytes < uint32(64) {
			i = uint32(0)
			for {
				if !(i < bytes) {
					break
				}
				(*(*[64]Tu8)(unsafe.Pointer(bp)))[i] = *(*Tu8)(unsafe.Pointer(m + uintptr(i)))
				goto _2
			_2:
				;
				i++
			}
			m = bp
			ctarget = c
			c = bp
		}
		x0 = j0
		x1 = j1
		x2 = j2
		x3 = j3
		x4 = j4
		x5 = j5
		x6 = j6
		x7 = j7
		x8 = j8
		x9 = j9
		x10 = j10
		x11 = j11
		x12 = j12
		x13 = j13
		x14 = j14
		x15 = j15
		i = uint32(20)
		for {
			if !(i > uint32(0)) {
				break
			}
			x0 = (x0 + x4) & libc.Uint32FromUint32(0xFFFFFFFF)
			x12 = (x12^x0)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x12^x0)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x8 = (x8 + x12) & libc.Uint32FromUint32(0xFFFFFFFF)
			x4 = (x4^x8)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x4^x8)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x0 = (x0 + x4) & libc.Uint32FromUint32(0xFFFFFFFF)
			x12 = (x12^x0)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x12^x0)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x8 = (x8 + x12) & libc.Uint32FromUint32(0xFFFFFFFF)
			x4 = (x4^x8)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x4^x8)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x1 = (x1 + x5) & libc.Uint32FromUint32(0xFFFFFFFF)
			x13 = (x13^x1)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x13^x1)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x9 = (x9 + x13) & libc.Uint32FromUint32(0xFFFFFFFF)
			x5 = (x5^x9)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x5^x9)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x1 = (x1 + x5) & libc.Uint32FromUint32(0xFFFFFFFF)
			x13 = (x13^x1)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x13^x1)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x9 = (x9 + x13) & libc.Uint32FromUint32(0xFFFFFFFF)
			x5 = (x5^x9)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x5^x9)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x2 = (x2 + x6) & libc.Uint32FromUint32(0xFFFFFFFF)
			x14 = (x14^x2)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x14^x2)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x10 = (x10 + x14) & libc.Uint32FromUint32(0xFFFFFFFF)
			x6 = (x6^x10)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x6^x10)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x2 = (x2 + x6) & libc.Uint32FromUint32(0xFFFFFFFF)
			x14 = (x14^x2)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x14^x2)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x10 = (x10 + x14) & libc.Uint32FromUint32(0xFFFFFFFF)
			x6 = (x6^x10)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x6^x10)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x3 = (x3 + x7) & libc.Uint32FromUint32(0xFFFFFFFF)
			x15 = (x15^x3)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x15^x3)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x11 = (x11 + x15) & libc.Uint32FromUint32(0xFFFFFFFF)
			x7 = (x7^x11)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x7^x11)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x3 = (x3 + x7) & libc.Uint32FromUint32(0xFFFFFFFF)
			x15 = (x15^x3)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x15^x3)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x11 = (x11 + x15) & libc.Uint32FromUint32(0xFFFFFFFF)
			x7 = (x7^x11)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x7^x11)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x0 = (x0 + x5) & libc.Uint32FromUint32(0xFFFFFFFF)
			x15 = (x15^x0)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x15^x0)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x10 = (x10 + x15) & libc.Uint32FromUint32(0xFFFFFFFF)
			x5 = (x5^x10)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x5^x10)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x0 = (x0 + x5) & libc.Uint32FromUint32(0xFFFFFFFF)
			x15 = (x15^x0)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x15^x0)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x10 = (x10 + x15) & libc.Uint32FromUint32(0xFFFFFFFF)
			x5 = (x5^x10)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x5^x10)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x1 = (x1 + x6) & libc.Uint32FromUint32(0xFFFFFFFF)
			x12 = (x12^x1)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x12^x1)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x11 = (x11 + x12) & libc.Uint32FromUint32(0xFFFFFFFF)
			x6 = (x6^x11)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x6^x11)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x1 = (x1 + x6) & libc.Uint32FromUint32(0xFFFFFFFF)
			x12 = (x12^x1)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x12^x1)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x11 = (x11 + x12) & libc.Uint32FromUint32(0xFFFFFFFF)
			x6 = (x6^x11)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x6^x11)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x2 = (x2 + x7) & libc.Uint32FromUint32(0xFFFFFFFF)
			x13 = (x13^x2)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x13^x2)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x8 = (x8 + x13) & libc.Uint32FromUint32(0xFFFFFFFF)
			x7 = (x7^x8)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x7^x8)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x2 = (x2 + x7) & libc.Uint32FromUint32(0xFFFFFFFF)
			x13 = (x13^x2)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x13^x2)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x8 = (x8 + x13) & libc.Uint32FromUint32(0xFFFFFFFF)
			x7 = (x7^x8)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x7^x8)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			x3 = (x3 + x4) & libc.Uint32FromUint32(0xFFFFFFFF)
			x14 = (x14^x3)<<libc.Int32FromInt32(16)&libc.Uint32FromUint32(0xFFFFFFFF) | (x14^x3)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
			x9 = (x9 + x14) & libc.Uint32FromUint32(0xFFFFFFFF)
			x4 = (x4^x9)<<libc.Int32FromInt32(12)&libc.Uint32FromUint32(0xFFFFFFFF) | (x4^x9)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
			x3 = (x3 + x4) & libc.Uint32FromUint32(0xFFFFFFFF)
			x14 = (x14^x3)<<libc.Int32FromInt32(8)&libc.Uint32FromUint32(0xFFFFFFFF) | (x14^x3)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8))
			x9 = (x9 + x14) & libc.Uint32FromUint32(0xFFFFFFFF)
			x4 = (x4^x9)<<libc.Int32FromInt32(7)&libc.Uint32FromUint32(0xFFFFFFFF) | (x4^x9)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
			goto _3
		_3:
			;
			i -= uint32(2)
		}
		x0 = (x0 + j0) & libc.Uint32FromUint32(0xFFFFFFFF)
		x1 = (x1 + j1) & libc.Uint32FromUint32(0xFFFFFFFF)
		x2 = (x2 + j2) & libc.Uint32FromUint32(0xFFFFFFFF)
		x3 = (x3 + j3) & libc.Uint32FromUint32(0xFFFFFFFF)
		x4 = (x4 + j4) & libc.Uint32FromUint32(0xFFFFFFFF)
		x5 = (x5 + j5) & libc.Uint32FromUint32(0xFFFFFFFF)
		x6 = (x6 + j6) & libc.Uint32FromUint32(0xFFFFFFFF)
		x7 = (x7 + j7) & libc.Uint32FromUint32(0xFFFFFFFF)
		x8 = (x8 + j8) & libc.Uint32FromUint32(0xFFFFFFFF)
		x9 = (x9 + j9) & libc.Uint32FromUint32(0xFFFFFFFF)
		x10 = (x10 + j10) & libc.Uint32FromUint32(0xFFFFFFFF)
		x11 = (x11 + j11) & libc.Uint32FromUint32(0xFFFFFFFF)
		x12 = (x12 + j12) & libc.Uint32FromUint32(0xFFFFFFFF)
		x13 = (x13 + j13) & libc.Uint32FromUint32(0xFFFFFFFF)
		x14 = (x14 + j14) & libc.Uint32FromUint32(0xFFFFFFFF)
		x15 = (x15 + j15) & libc.Uint32FromUint32(0xFFFFFFFF)
		j12 = (j12 + libc.Uint32FromInt32(libc.Int32FromInt32(1))) & libc.Uint32FromUint32(0xFFFFFFFF)
		if !(j12 != 0) {
			j13 = (j13 + libc.Uint32FromInt32(libc.Int32FromInt32(1))) & libc.Uint32FromUint32(0xFFFFFFFF)
		}
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(0))) = uint8(uint32(uint8(x0)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(0) + 1)) = uint8(uint32(uint8(x0>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(0) + 2)) = uint8(uint32(uint8(x0>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(0) + 3)) = uint8(uint32(uint8(x0>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(4))) = uint8(uint32(uint8(x1)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(4) + 1)) = uint8(uint32(uint8(x1>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(4) + 2)) = uint8(uint32(uint8(x1>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(4) + 3)) = uint8(uint32(uint8(x1>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(8))) = uint8(uint32(uint8(x2)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(8) + 1)) = uint8(uint32(uint8(x2>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(8) + 2)) = uint8(uint32(uint8(x2>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(8) + 3)) = uint8(uint32(uint8(x2>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(12))) = uint8(uint32(uint8(x3)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(12) + 1)) = uint8(uint32(uint8(x3>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(12) + 2)) = uint8(uint32(uint8(x3>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(12) + 3)) = uint8(uint32(uint8(x3>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(16))) = uint8(uint32(uint8(x4)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(16) + 1)) = uint8(uint32(uint8(x4>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(16) + 2)) = uint8(uint32(uint8(x4>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(16) + 3)) = uint8(uint32(uint8(x4>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(20))) = uint8(uint32(uint8(x5)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(20) + 1)) = uint8(uint32(uint8(x5>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(20) + 2)) = uint8(uint32(uint8(x5>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(20) + 3)) = uint8(uint32(uint8(x5>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(24))) = uint8(uint32(uint8(x6)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(24) + 1)) = uint8(uint32(uint8(x6>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(24) + 2)) = uint8(uint32(uint8(x6>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(24) + 3)) = uint8(uint32(uint8(x6>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(28))) = uint8(uint32(uint8(x7)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(28) + 1)) = uint8(uint32(uint8(x7>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(28) + 2)) = uint8(uint32(uint8(x7>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(28) + 3)) = uint8(uint32(uint8(x7>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(32))) = uint8(uint32(uint8(x8)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(32) + 1)) = uint8(uint32(uint8(x8>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(32) + 2)) = uint8(uint32(uint8(x8>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(32) + 3)) = uint8(uint32(uint8(x8>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(36))) = uint8(uint32(uint8(x9)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(36) + 1)) = uint8(uint32(uint8(x9>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(36) + 2)) = uint8(uint32(uint8(x9>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(36) + 3)) = uint8(uint32(uint8(x9>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(40))) = uint8(uint32(uint8(x10)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(40) + 1)) = uint8(uint32(uint8(x10>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(40) + 2)) = uint8(uint32(uint8(x10>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(40) + 3)) = uint8(uint32(uint8(x10>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(44))) = uint8(uint32(uint8(x11)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(44) + 1)) = uint8(uint32(uint8(x11>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(44) + 2)) = uint8(uint32(uint8(x11>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(44) + 3)) = uint8(uint32(uint8(x11>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(48))) = uint8(uint32(uint8(x12)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(48) + 1)) = uint8(uint32(uint8(x12>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(48) + 2)) = uint8(uint32(uint8(x12>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(48) + 3)) = uint8(uint32(uint8(x12>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(52))) = uint8(uint32(uint8(x13)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(52) + 1)) = uint8(uint32(uint8(x13>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(52) + 2)) = uint8(uint32(uint8(x13>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(52) + 3)) = uint8(uint32(uint8(x13>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(56))) = uint8(uint32(uint8(x14)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(56) + 1)) = uint8(uint32(uint8(x14>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(56) + 2)) = uint8(uint32(uint8(x14>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(56) + 3)) = uint8(uint32(uint8(x14>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(60))) = uint8(uint32(uint8(x15)) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(60) + 1)) = uint8(uint32(uint8(x15>>libc.Int32FromInt32(8))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(60) + 2)) = uint8(uint32(uint8(x15>>libc.Int32FromInt32(16))) & libc.Uint32FromUint32(0xFF))
		*(*Tu8)(unsafe.Pointer(c + libc.UintptrFromInt32(60) + 3)) = uint8(uint32(uint8(x15>>libc.Int32FromInt32(24))) & libc.Uint32FromUint32(0xFF))
		if bytes <= uint32(64) {
			if bytes < uint32(64) {
				i = uint32(0)
				for {
					if !(i < bytes) {
						break
					}
					*(*Tu8)(unsafe.Pointer(ctarget + uintptr(i))) = *(*Tu8)(unsafe.Pointer(c + uintptr(i)))
					goto _4
				_4:
					;
					i++
				}
			}
			*(*Tu32)(unsafe.Pointer(x + 12*4)) = j12
			*(*Tu32)(unsafe.Pointer(x + 13*4)) = j13
			return
		}
		bytes -= uint32(64)
		c += uintptr(64)
		goto _1
	_1:
	}
}

// C documentation
//
//	/* Marked MAP_INHERIT_ZERO, so zero'd out in fork children. */
type T_rs = struct {
	Frs_have  Tsize_t
	Frs_count Tsize_t
}

// C documentation
//
//	/* Marked MAP_INHERIT_ZERO, so zero'd out in fork children. */
var _rs uintptr

// C documentation
//
//	/* Maybe be preserved in fork children, if _rs_allocate() decides. */
type T_rsx = struct {
	Frs_chacha Tchacha_ctx
	Frs_buf    [1024]uint8
}

// C documentation
//
//	/* Maybe be preserved in fork children, if _rs_allocate() decides. */
var _rsx uintptr

type Trlim_t = uint64

type Trlimit = struct {
	Frlim_cur Trlim_t
	Frlim_max Trlim_t
}

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
	F__reserved  [16]int64
}

type Tsched_param = struct {
	Fsched_priority int32
	F__reserved1    int32
	F__reserved2    [2]struct {
		F__reserved1 Ttime_t
		F__reserved2 int64
	}
	F__reserved3 int32
}

type Tcpu_set_t = struct {
	F__bits [16]uint64
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
}

type Titimerspec = struct {
	Fit_interval Ttimespec
	Fit_value    Ttimespec
}

type t__ptcb = struct {
	F__f    uintptr
	F__x    uintptr
	F__next uintptr
}

type Tcpu_set_t1 = struct {
	F__bits [16]uint64
}

var _arc4random_mtx = Tpthread_mutex_t{}
var __rs_forked Tsig_atomic_t

func __rs_init(tls *libc.TLS, buf uintptr, n Tsize_t) {
	var v1, v2, v3, v6 uintptr
	var v4 int32
	_, _, _, _, _ = v1, v2, v3, v4, v6
	if n < libc.Uint64FromInt32(libc.Int32FromInt32(m_KEYSZ)+libc.Int32FromInt32(m_IVSZ)) {
		return
	}
	if _rs == libc.UintptrFromInt32(0) {
		v1 = uintptr(unsafe.Pointer(&_rs))
		v2 = uintptr(unsafe.Pointer(&_rsx))
		v3 = libc.Xmmap(tls, libc.UintptrFromInt32(0), uint64(16), libc.Int32FromInt32(m_PROT_READ)|libc.Int32FromInt32(m_PROT_WRITE), libc.Int32FromInt32(m_MAP_ANON)|libc.Int32FromInt32(m_MAP_PRIVATE), -int32(1), 0)
		*(*uintptr)(unsafe.Pointer(v1)) = v3
		if v3 == uintptr(-libc.Int32FromInt32(1)) {
			v4 = -int32(1)
			goto _5
		}
		v6 = libc.Xmmap(tls, libc.UintptrFromInt32(0), uint64(1088), libc.Int32FromInt32(m_PROT_READ)|libc.Int32FromInt32(m_PROT_WRITE), libc.Int32FromInt32(m_MAP_ANON)|libc.Int32FromInt32(m_MAP_PRIVATE), -int32(1), 0)
		*(*uintptr)(unsafe.Pointer(v2)) = v6
		if v6 == uintptr(-libc.Int32FromInt32(1)) {
			libc.Xmunmap(tls, *(*uintptr)(unsafe.Pointer(v1)), uint64(16))
			*(*uintptr)(unsafe.Pointer(v1)) = libc.UintptrFromInt32(0)
			v4 = -int32(1)
			goto _5
		}
		libc.Xpthread_atfork(tls, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), __ccgo_fp(__rs_forkhandler))
		v4 = 0
		goto _5
	_5:
		if v4 == -int32(1) {
			libc.X_exit(tls, int32(1))
		}
	}
	_chacha_keysetup(tls, _rsx, buf, libc.Uint32FromInt32(libc.Int32FromInt32(m_KEYSZ)*libc.Int32FromInt32(8)))
	_chacha_ivsetup(tls, _rsx, buf+uintptr(m_KEYSZ))
}

func __rs_rekey(tls *libc.TLS, dat uintptr, datlen Tsize_t) {
	var i, m Tsize_t
	var v1 uint64
	var p3 uintptr
	_, _, _, _ = i, m, v1, p3
	/* fill rs_buf with the keystream */
	_chacha_encrypt_bytes(tls, _rsx, _rsx+64, _rsx+64, uint32(1024))
	/* mix in optional user provided data */
	if dat != 0 {
		if datlen < libc.Uint64FromInt32(libc.Int32FromInt32(m_KEYSZ)+libc.Int32FromInt32(m_IVSZ)) {
			v1 = datlen
		} else {
			v1 = libc.Uint64FromInt32(libc.Int32FromInt32(m_KEYSZ) + libc.Int32FromInt32(m_IVSZ))
		}
		m = v1
		i = uint64(0)
		for {
			if !(i < m) {
				break
			}
			p3 = _rsx + 64 + uintptr(i)
			*(*uint8)(unsafe.Pointer(p3)) = uint8(int32(*(*uint8)(unsafe.Pointer(p3))) ^ libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(dat + uintptr(i)))))
			goto _2
		_2:
			;
			i++
		}
	}
	/* immediately reinit for backtracking resistance */
	__rs_init(tls, _rsx+64, libc.Uint64FromInt32(libc.Int32FromInt32(m_KEYSZ)+libc.Int32FromInt32(m_IVSZ)))
	libc.Xmemset(tls, _rsx+64, 0, libc.Uint64FromInt32(libc.Int32FromInt32(m_KEYSZ)+libc.Int32FromInt32(m_IVSZ)))
	(*T_rs)(unsafe.Pointer(_rs)).Frs_have = libc.Uint64FromInt64(1024) - libc.Uint64FromInt32(m_KEYSZ) - libc.Uint64FromInt32(m_IVSZ)
}

func __rs_stir(tls *libc.TLS) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var _ /* rekey_fuzz at bp+40 */ Tuint32_t
	var _ /* rnd at bp+0 */ [40]uint8
	*(*Tuint32_t)(unsafe.Pointer(bp + 40)) = uint32(0)
	if libc.Xgetentropy(tls, bp, uint64(40)) == -int32(1) {
		libc.Xraise(tls, int32(m_SIGKILL))
	}
	if !(_rs != 0) {
		__rs_init(tls, bp, uint64(40))
	} else {
		__rs_rekey(tls, bp, uint64(40))
	}
	Xexplicit_bzero(tls, bp, uint64(40)) /* discard source seed */
	/* invalidate rs_buf */
	(*T_rs)(unsafe.Pointer(_rs)).Frs_have = uint64(0)
	libc.Xmemset(tls, _rsx+64, 0, uint64(1024))
	/* rekey interval should not be predictable */
	_chacha_encrypt_bytes(tls, _rsx, bp+40, bp+40, uint32(4))
	(*T_rs)(unsafe.Pointer(_rs)).Frs_count = uint64(libc.Uint32FromInt32(libc.Int32FromInt32(1024)*libc.Int32FromInt32(1024)) + *(*Tuint32_t)(unsafe.Pointer(bp + 40))%libc.Uint32FromInt32(libc.Int32FromInt32(1024)*libc.Int32FromInt32(1024)))
}

func __rs_stir_if_needed(tls *libc.TLS, len1 Tsize_t) {
	__rs_forkdetect(tls)
	if !(_rs != 0) || (*T_rs)(unsafe.Pointer(_rs)).Frs_count <= len1 {
		__rs_stir(tls)
	}
	if (*T_rs)(unsafe.Pointer(_rs)).Frs_count <= len1 {
		(*T_rs)(unsafe.Pointer(_rs)).Frs_count = uint64(0)
	} else {
		*(*Tsize_t)(unsafe.Pointer(_rs + 8)) -= len1
	}
}

func __rs_random_buf(tls *libc.TLS, _buf uintptr, n Tsize_t) {
	var buf, keystream uintptr
	var m Tsize_t
	var v1 uint64
	_, _, _, _ = buf, keystream, m, v1
	buf = _buf
	__rs_stir_if_needed(tls, n)
	for n > uint64(0) {
		if (*T_rs)(unsafe.Pointer(_rs)).Frs_have > uint64(0) {
			if n < (*T_rs)(unsafe.Pointer(_rs)).Frs_have {
				v1 = n
			} else {
				v1 = (*T_rs)(unsafe.Pointer(_rs)).Frs_have
			}
			m = v1
			keystream = _rsx + 64 + uintptr(1024) - uintptr((*T_rs)(unsafe.Pointer(_rs)).Frs_have)
			libc.Xmemcpy(tls, buf, keystream, m)
			libc.Xmemset(tls, keystream, 0, m)
			buf += uintptr(m)
			n -= m
			*(*Tsize_t)(unsafe.Pointer(_rs)) -= m
		}
		if (*T_rs)(unsafe.Pointer(_rs)).Frs_have == uint64(0) {
			__rs_rekey(tls, libc.UintptrFromInt32(0), uint64(0))
		}
	}
}

func __rs_random_u32(tls *libc.TLS, val uintptr) {
	var keystream uintptr
	_ = keystream
	__rs_stir_if_needed(tls, uint64(4))
	if (*T_rs)(unsafe.Pointer(_rs)).Frs_have < uint64(4) {
		__rs_rekey(tls, libc.UintptrFromInt32(0), uint64(0))
	}
	keystream = _rsx + 64 + uintptr(1024) - uintptr((*T_rs)(unsafe.Pointer(_rs)).Frs_have)
	libc.Xmemcpy(tls, val, keystream, uint64(4))
	libc.Xmemset(tls, keystream, 0, uint64(4))
	*(*Tsize_t)(unsafe.Pointer(_rs)) -= uint64(4)
}

func Xarc4random_stir(tls *libc.TLS) {
	libc.Xpthread_mutex_lock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
	__rs_stir(tls)
	libc.Xpthread_mutex_unlock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
}

func Xarc4random_addrandom(tls *libc.TLS, dat uintptr, datlen int32) {
	libc.Xpthread_mutex_lock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
	__rs_stir_if_needed(tls, libc.Uint64FromInt32(datlen))
	__rs_rekey(tls, dat, libc.Uint64FromInt32(datlen))
	libc.Xpthread_mutex_unlock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
}

func Xarc4random(tls *libc.TLS) (r Tuint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* val at bp+0 */ Tuint32_t
	libc.Xpthread_mutex_lock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
	__rs_random_u32(tls, bp)
	libc.Xpthread_mutex_unlock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
	return *(*Tuint32_t)(unsafe.Pointer(bp))
}

func Xarc4random_buf(tls *libc.TLS, buf uintptr, n Tsize_t) {
	libc.Xpthread_mutex_lock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
	__rs_random_buf(tls, buf, n)
	libc.Xpthread_mutex_unlock(tls, uintptr(unsafe.Pointer(&_arc4random_mtx)))
}

// C documentation
//
//	/*
//	 * Calculate a uniformly distributed random number less than upper_bound
//	 * avoiding "modulo bias".
//	 *
//	 * Uniformity is achieved by generating new random numbers until the one
//	 * returned is outside the range [0, 2**32 % upper_bound).  This
//	 * guarantees the selected random number will be inside
//	 * [2**32 % upper_bound, 2**32) which maps back to [0, upper_bound)
//	 * after reduction modulo upper_bound.
//	 */
func Xarc4random_uniform(tls *libc.TLS, upper_bound Tuint32_t) (r1 Tuint32_t) {
	var min, r Tuint32_t
	_, _ = min, r
	if upper_bound < uint32(2) {
		return uint32(0)
	}
	/* 2**32 % x == (2**32 - x) % x */
	min = -upper_bound % upper_bound
	/*
	 * This could theoretically loop forever but each retry has
	 * p > 0.5 (worst case, usually far better) of selecting a
	 * number inside the range we need, so it should rarely need
	 * to re-roll.
	 */
	for {
		r = Xarc4random(tls)
		if r >= min {
			break
		}
		goto _1
	_1:
	}
	return r % upper_bound
}

const m_SEEK_CUR = 1
const m_SEEK_END = 2
const m_SEEK_SET = 0
const m_no_argument = 0
const m_optional_argument = 2
const m_required_argument = 1

type Toption = struct {
	Fname    uintptr
	Fhas_arg int32
	Fflag    uintptr
	Fval     int32
}

func Xbsd_getopt(tls *libc.TLS, argc int32, argv uintptr, shortopts uintptr) (r int32) {
	if Xoptreset == int32(1) {
		Xoptreset = 0
		libc.Xoptind = 0
	}
	/*
	 * Make sure we are using the system getopt() and not a possible
	 * overlay macro.
	 */
	return libc.Xgetopt(tls, argc, argv, shortopts)
}

const m_DT_BLK = 6
const m_DT_CHR = 2
const m_DT_DIR = 4
const m_DT_FIFO = 1
const m_DT_LNK = 10
const m_DT_REG = 8
const m_DT_SOCK = 12
const m_DT_UNKNOWN = 0
const m_DT_WHT = 14
const m_INT_MAX1 = 2147483647
const m_OPEN_MAX = 256
const m_SYS_accept = 202
const m_SYS_accept4 = 242
const m_SYS_acct = 89
const m_SYS_add_key = 217
const m_SYS_adjtimex = 171
const m_SYS_arch_specific_syscall = 244
const m_SYS_bind = 200
const m_SYS_bpf = 280
const m_SYS_brk = 214
const m_SYS_cachestat = 451
const m_SYS_capget = 90
const m_SYS_capset = 91
const m_SYS_chdir = 49
const m_SYS_chroot = 51
const m_SYS_clock_adjtime = 266
const m_SYS_clock_getres = 114
const m_SYS_clock_gettime = 113
const m_SYS_clock_nanosleep = 115
const m_SYS_clock_settime = 112
const m_SYS_clone = 220
const m_SYS_clone3 = 435
const m_SYS_close = 57
const m_SYS_close_range = 436
const m_SYS_connect = 203
const m_SYS_copy_file_range = 285
const m_SYS_delete_module = 106
const m_SYS_dup = 23
const m_SYS_dup3 = 24
const m_SYS_epoll_create1 = 20
const m_SYS_epoll_ctl = 21
const m_SYS_epoll_pwait = 22
const m_SYS_epoll_pwait2 = 441
const m_SYS_eventfd2 = 19
const m_SYS_execve = 221
const m_SYS_execveat = 281
const m_SYS_exit = 93
const m_SYS_exit_group = 94
const m_SYS_faccessat = 48
const m_SYS_faccessat2 = 439
const m_SYS_fadvise64 = "__NR3264_fadvise64"
const m_SYS_fallocate = 47
const m_SYS_fanotify_init = 262
const m_SYS_fanotify_mark = 263
const m_SYS_fchdir = 50
const m_SYS_fchmod = 52
const m_SYS_fchmodat = 53
const m_SYS_fchmodat2 = 452
const m_SYS_fchown = 55
const m_SYS_fchownat = 54
const m_SYS_fcntl = "__NR3264_fcntl"
const m_SYS_fdatasync = 83
const m_SYS_fgetxattr = 10
const m_SYS_finit_module = 273
const m_SYS_flistxattr = 13
const m_SYS_flock = 32
const m_SYS_fremovexattr = 16
const m_SYS_fsconfig = 431
const m_SYS_fsetxattr = 7
const m_SYS_fsmount = 432
const m_SYS_fsopen = 430
const m_SYS_fspick = 433
const m_SYS_fstatfs = "__NR3264_fstatfs"
const m_SYS_fsync = 82
const m_SYS_ftruncate = "__NR3264_ftruncate"
const m_SYS_futex = 98
const m_SYS_futex_requeue = 456
const m_SYS_futex_wait = 455
const m_SYS_futex_waitv = 449
const m_SYS_futex_wake = 454
const m_SYS_get_mempolicy = 236
const m_SYS_get_robust_list = 100
const m_SYS_getcpu = 168
const m_SYS_getcwd = 17
const m_SYS_getdents64 = 61
const m_SYS_getegid = 177
const m_SYS_geteuid = 175
const m_SYS_getgid = 176
const m_SYS_getgroups = 158
const m_SYS_getitimer = 102
const m_SYS_getpeername = 205
const m_SYS_getpgid = 155
const m_SYS_getpid = 172
const m_SYS_getppid = 173
const m_SYS_getpriority = 141
const m_SYS_getrandom = 278
const m_SYS_getresgid = 150
const m_SYS_getresuid = 148
const m_SYS_getrusage = 165
const m_SYS_getsid = 156
const m_SYS_getsockname = 204
const m_SYS_getsockopt = 209
const m_SYS_gettid = 178
const m_SYS_gettimeofday = 169
const m_SYS_getuid = 174
const m_SYS_getxattr = 8
const m_SYS_init_module = 105
const m_SYS_inotify_add_watch = 27
const m_SYS_inotify_init1 = 26
const m_SYS_inotify_rm_watch = 28
const m_SYS_io_cancel = 3
const m_SYS_io_destroy = 1
const m_SYS_io_getevents = 4
const m_SYS_io_pgetevents = 292
const m_SYS_io_setup = 0
const m_SYS_io_submit = 2
const m_SYS_io_uring_enter = 426
const m_SYS_io_uring_register = 427
const m_SYS_io_uring_setup = 425
const m_SYS_ioctl = 29
const m_SYS_ioprio_get = 31
const m_SYS_ioprio_set = 30
const m_SYS_kcmp = 272
const m_SYS_kexec_file_load = 294
const m_SYS_kexec_load = 104
const m_SYS_keyctl = 219
const m_SYS_kill = 129
const m_SYS_landlock_add_rule = 445
const m_SYS_landlock_create_ruleset = 444
const m_SYS_landlock_restrict_self = 446
const m_SYS_lgetxattr = 9
const m_SYS_linkat = 37
const m_SYS_listen = 201
const m_SYS_listxattr = 11
const m_SYS_llistxattr = 12
const m_SYS_lookup_dcookie = 18
const m_SYS_lremovexattr = 15
const m_SYS_lseek = "__NR3264_lseek"
const m_SYS_lsetxattr = 6
const m_SYS_madvise = 233
const m_SYS_map_shadow_stack = 453
const m_SYS_mbind = 235
const m_SYS_membarrier = 283
const m_SYS_memfd_create = 279
const m_SYS_migrate_pages = 238
const m_SYS_mincore = 232
const m_SYS_mkdirat = 34
const m_SYS_mknodat = 33
const m_SYS_mlock = 228
const m_SYS_mlock2 = 284
const m_SYS_mlockall = 230
const m_SYS_mmap = "__NR3264_mmap"
const m_SYS_mount = 40
const m_SYS_mount_setattr = 442
const m_SYS_move_mount = 429
const m_SYS_move_pages = 239
const m_SYS_mprotect = 226
const m_SYS_mq_getsetattr = 185
const m_SYS_mq_notify = 184
const m_SYS_mq_open = 180
const m_SYS_mq_timedreceive = 183
const m_SYS_mq_timedsend = 182
const m_SYS_mq_unlink = 181
const m_SYS_mremap = 216
const m_SYS_msgctl = 187
const m_SYS_msgget = 186
const m_SYS_msgrcv = 188
const m_SYS_msgsnd = 189
const m_SYS_msync = 227
const m_SYS_munlock = 229
const m_SYS_munlockall = 231
const m_SYS_munmap = 215
const m_SYS_name_to_handle_at = 264
const m_SYS_nanosleep = 101
const m_SYS_nfsservctl = 42
const m_SYS_open_by_handle_at = 265
const m_SYS_open_tree = 428
const m_SYS_openat = 56
const m_SYS_openat2 = 437
const m_SYS_perf_event_open = 241
const m_SYS_personality = 92
const m_SYS_pidfd_getfd = 438
const m_SYS_pidfd_open = 434
const m_SYS_pidfd_send_signal = 424
const m_SYS_pipe2 = 59
const m_SYS_pivot_root = 41
const m_SYS_pkey_alloc = 289
const m_SYS_pkey_free = 290
const m_SYS_pkey_mprotect = 288
const m_SYS_ppoll = 73
const m_SYS_prctl = 167
const m_SYS_pread64 = 67
const m_SYS_preadv = 69
const m_SYS_preadv2 = 286
const m_SYS_prlimit64 = 261
const m_SYS_process_madvise = 440
const m_SYS_process_mrelease = 448
const m_SYS_process_vm_readv = 270
const m_SYS_process_vm_writev = 271
const m_SYS_pselect6 = 72
const m_SYS_ptrace = 117
const m_SYS_pwrite64 = 68
const m_SYS_pwritev = 70
const m_SYS_pwritev2 = 287
const m_SYS_quotactl = 60
const m_SYS_quotactl_fd = 443
const m_SYS_read = 63
const m_SYS_readahead = 213
const m_SYS_readlinkat = 78
const m_SYS_readv = 65
const m_SYS_reboot = 142
const m_SYS_recvfrom = 207
const m_SYS_recvmmsg = 243
const m_SYS_recvmsg = 212
const m_SYS_remap_file_pages = 234
const m_SYS_removexattr = 14
const m_SYS_renameat2 = 276
const m_SYS_request_key = 218
const m_SYS_restart_syscall = 128
const m_SYS_rseq = 293
const m_SYS_rt_sigaction = 134
const m_SYS_rt_sigpending = 136
const m_SYS_rt_sigprocmask = 135
const m_SYS_rt_sigqueueinfo = 138
const m_SYS_rt_sigreturn = 139
const m_SYS_rt_sigsuspend = 133
const m_SYS_rt_sigtimedwait = 137
const m_SYS_rt_tgsigqueueinfo = 240
const m_SYS_sched_get_priority_max = 125
const m_SYS_sched_get_priority_min = 126
const m_SYS_sched_getaffinity = 123
const m_SYS_sched_getattr = 275
const m_SYS_sched_getparam = 121
const m_SYS_sched_getscheduler = 120
const m_SYS_sched_rr_get_interval = 127
const m_SYS_sched_setaffinity = 122
const m_SYS_sched_setattr = 274
const m_SYS_sched_setparam = 118
const m_SYS_sched_setscheduler = 119
const m_SYS_sched_yield = 124
const m_SYS_seccomp = 277
const m_SYS_semctl = 191
const m_SYS_semget = 190
const m_SYS_semop = 193
const m_SYS_semtimedop = 192
const m_SYS_sendfile = "__NR3264_sendfile"
const m_SYS_sendmmsg = 269
const m_SYS_sendmsg = 211
const m_SYS_sendto = 206
const m_SYS_set_mempolicy = 237
const m_SYS_set_mempolicy_home_node = 450
const m_SYS_set_robust_list = 99
const m_SYS_set_tid_address = 96
const m_SYS_setdomainname = 162
const m_SYS_setfsgid = 152
const m_SYS_setfsuid = 151
const m_SYS_setgid = 144
const m_SYS_setgroups = 159
const m_SYS_sethostname = 161
const m_SYS_setitimer = 103
const m_SYS_setns = 268
const m_SYS_setpgid = 154
const m_SYS_setpriority = 140
const m_SYS_setregid = 143
const m_SYS_setresgid = 149
const m_SYS_setresuid = 147
const m_SYS_setreuid = 145
const m_SYS_setsid = 157
const m_SYS_setsockopt = 208
const m_SYS_settimeofday = 170
const m_SYS_setuid = 146
const m_SYS_setxattr = 5
const m_SYS_shmat = 196
const m_SYS_shmctl = 195
const m_SYS_shmdt = 197
const m_SYS_shmget = 194
const m_SYS_shutdown = 210
const m_SYS_sigaltstack = 132
const m_SYS_signalfd4 = 74
const m_SYS_socket = 198
const m_SYS_socketpair = 199
const m_SYS_splice = 76
const m_SYS_statfs = "__NR3264_statfs"
const m_SYS_statx = 291
const m_SYS_swapoff = 225
const m_SYS_swapon = 224
const m_SYS_symlinkat = 36
const m_SYS_sync = 81
const m_SYS_sync_file_range = 84
const m_SYS_syncfs = 267
const m_SYS_sysinfo = 179
const m_SYS_syslog = 116
const m_SYS_tee = 77
const m_SYS_tgkill = 131
const m_SYS_timer_create = 107
const m_SYS_timer_delete = 111
const m_SYS_timer_getoverrun = 109
const m_SYS_timer_gettime = 108
const m_SYS_timer_settime = 110
const m_SYS_timerfd_create = 85
const m_SYS_timerfd_gettime = 87
const m_SYS_timerfd_settime = 86
const m_SYS_times = 153
const m_SYS_tkill = 130
const m_SYS_truncate = "__NR3264_truncate"
const m_SYS_umask = 166
const m_SYS_umount2 = 39
const m_SYS_uname = 160
const m_SYS_unlinkat = 35
const m_SYS_unshare = 97
const m_SYS_userfaultfd = 282
const m_SYS_utimensat = 88
const m_SYS_vhangup = 58
const m_SYS_vmsplice = 75
const m_SYS_wait4 = 260
const m_SYS_waitid = 95
const m_SYS_write = 64
const m_SYS_writev = 66
const m_UINT_MAX1 = 4294967295
const m___NR3264_fadvise64 = 223
const m___NR3264_fcntl = 25
const m___NR3264_fstatfs = 44
const m___NR3264_ftruncate = 46
const m___NR3264_lseek = 62
const m___NR3264_mmap = 222
const m___NR3264_sendfile = 71
const m___NR3264_statfs = 43
const m___NR3264_truncate = 45
const m___NR_accept = 202
const m___NR_accept4 = 242
const m___NR_acct = 89
const m___NR_add_key = 217
const m___NR_adjtimex = 171
const m___NR_arch_specific_syscall = 244
const m___NR_bind = 200
const m___NR_bpf = 280
const m___NR_brk = 214
const m___NR_cachestat = 451
const m___NR_capget = 90
const m___NR_capset = 91
const m___NR_chdir = 49
const m___NR_chroot = 51
const m___NR_clock_adjtime = 266
const m___NR_clock_getres = 114
const m___NR_clock_gettime = 113
const m___NR_clock_nanosleep = 115
const m___NR_clock_settime = 112
const m___NR_clone = 220
const m___NR_clone3 = 435
const m___NR_close = 57
const m___NR_close_range = 436
const m___NR_connect = 203
const m___NR_copy_file_range = 285
const m___NR_delete_module = 106
const m___NR_dup = 23
const m___NR_dup3 = 24
const m___NR_epoll_create1 = 20
const m___NR_epoll_ctl = 21
const m___NR_epoll_pwait = 22
const m___NR_epoll_pwait2 = 441
const m___NR_eventfd2 = 19
const m___NR_execve = 221
const m___NR_execveat = 281
const m___NR_exit = 93
const m___NR_exit_group = 94
const m___NR_faccessat = 48
const m___NR_faccessat2 = 439
const m___NR_fadvise64 = "__NR3264_fadvise64"
const m___NR_fallocate = 47
const m___NR_fanotify_init = 262
const m___NR_fanotify_mark = 263
const m___NR_fchdir = 50
const m___NR_fchmod = 52
const m___NR_fchmodat = 53
const m___NR_fchmodat2 = 452
const m___NR_fchown = 55
const m___NR_fchownat = 54
const m___NR_fcntl = "__NR3264_fcntl"
const m___NR_fdatasync = 83
const m___NR_fgetxattr = 10
const m___NR_finit_module = 273
const m___NR_flistxattr = 13
const m___NR_flock = 32
const m___NR_fremovexattr = 16
const m___NR_fsconfig = 431
const m___NR_fsetxattr = 7
const m___NR_fsmount = 432
const m___NR_fsopen = 430
const m___NR_fspick = 433
const m___NR_fstatfs = "__NR3264_fstatfs"
const m___NR_fsync = 82
const m___NR_ftruncate = "__NR3264_ftruncate"
const m___NR_futex = 98
const m___NR_futex_requeue = 456
const m___NR_futex_wait = 455
const m___NR_futex_waitv = 449
const m___NR_futex_wake = 454
const m___NR_get_mempolicy = 236
const m___NR_get_robust_list = 100
const m___NR_getcpu = 168
const m___NR_getcwd = 17
const m___NR_getdents64 = 61
const m___NR_getegid = 177
const m___NR_geteuid = 175
const m___NR_getgid = 176
const m___NR_getgroups = 158
const m___NR_getitimer = 102
const m___NR_getpeername = 205
const m___NR_getpgid = 155
const m___NR_getpid = 172
const m___NR_getppid = 173
const m___NR_getpriority = 141
const m___NR_getrandom = 278
const m___NR_getresgid = 150
const m___NR_getresuid = 148
const m___NR_getrusage = 165
const m___NR_getsid = 156
const m___NR_getsockname = 204
const m___NR_getsockopt = 209
const m___NR_gettid = 178
const m___NR_gettimeofday = 169
const m___NR_getuid = 174
const m___NR_getxattr = 8
const m___NR_init_module = 105
const m___NR_inotify_add_watch = 27
const m___NR_inotify_init1 = 26
const m___NR_inotify_rm_watch = 28
const m___NR_io_cancel = 3
const m___NR_io_destroy = 1
const m___NR_io_getevents = 4
const m___NR_io_pgetevents = 292
const m___NR_io_setup = 0
const m___NR_io_submit = 2
const m___NR_io_uring_enter = 426
const m___NR_io_uring_register = 427
const m___NR_io_uring_setup = 425
const m___NR_ioctl = 29
const m___NR_ioprio_get = 31
const m___NR_ioprio_set = 30
const m___NR_kcmp = 272
const m___NR_kexec_file_load = 294
const m___NR_kexec_load = 104
const m___NR_keyctl = 219
const m___NR_kill = 129
const m___NR_landlock_add_rule = 445
const m___NR_landlock_create_ruleset = 444
const m___NR_landlock_restrict_self = 446
const m___NR_lgetxattr = 9
const m___NR_linkat = 37
const m___NR_listen = 201
const m___NR_listxattr = 11
const m___NR_llistxattr = 12
const m___NR_lookup_dcookie = 18
const m___NR_lremovexattr = 15
const m___NR_lseek = "__NR3264_lseek"
const m___NR_lsetxattr = 6
const m___NR_madvise = 233
const m___NR_map_shadow_stack = 453
const m___NR_mbind = 235
const m___NR_membarrier = 283
const m___NR_memfd_create = 279
const m___NR_migrate_pages = 238
const m___NR_mincore = 232
const m___NR_mkdirat = 34
const m___NR_mknodat = 33
const m___NR_mlock = 228
const m___NR_mlock2 = 284
const m___NR_mlockall = 230
const m___NR_mmap = "__NR3264_mmap"
const m___NR_mount = 40
const m___NR_mount_setattr = 442
const m___NR_move_mount = 429
const m___NR_move_pages = 239
const m___NR_mprotect = 226
const m___NR_mq_getsetattr = 185
const m___NR_mq_notify = 184
const m___NR_mq_open = 180
const m___NR_mq_timedreceive = 183
const m___NR_mq_timedsend = 182
const m___NR_mq_unlink = 181
const m___NR_mremap = 216
const m___NR_msgctl = 187
const m___NR_msgget = 186
const m___NR_msgrcv = 188
const m___NR_msgsnd = 189
const m___NR_msync = 227
const m___NR_munlock = 229
const m___NR_munlockall = 231
const m___NR_munmap = 215
const m___NR_name_to_handle_at = 264
const m___NR_nanosleep = 101
const m___NR_nfsservctl = 42
const m___NR_open_by_handle_at = 265
const m___NR_open_tree = 428
const m___NR_openat = 56
const m___NR_openat2 = 437
const m___NR_perf_event_open = 241
const m___NR_personality = 92
const m___NR_pidfd_getfd = 438
const m___NR_pidfd_open = 434
const m___NR_pidfd_send_signal = 424
const m___NR_pipe2 = 59
const m___NR_pivot_root = 41
const m___NR_pkey_alloc = 289
const m___NR_pkey_free = 290
const m___NR_pkey_mprotect = 288
const m___NR_ppoll = 73
const m___NR_prctl = 167
const m___NR_pread64 = 67
const m___NR_preadv = 69
const m___NR_preadv2 = 286
const m___NR_prlimit64 = 261
const m___NR_process_madvise = 440
const m___NR_process_mrelease = 448
const m___NR_process_vm_readv = 270
const m___NR_process_vm_writev = 271
const m___NR_pselect6 = 72
const m___NR_ptrace = 117
const m___NR_pwrite64 = 68
const m___NR_pwritev = 70
const m___NR_pwritev2 = 287
const m___NR_quotactl = 60
const m___NR_quotactl_fd = 443
const m___NR_read = 63
const m___NR_readahead = 213
const m___NR_readlinkat = 78
const m___NR_readv = 65
const m___NR_reboot = 142
const m___NR_recvfrom = 207
const m___NR_recvmmsg = 243
const m___NR_recvmsg = 212
const m___NR_remap_file_pages = 234
const m___NR_removexattr = 14
const m___NR_renameat2 = 276
const m___NR_request_key = 218
const m___NR_restart_syscall = 128
const m___NR_rseq = 293
const m___NR_rt_sigaction = 134
const m___NR_rt_sigpending = 136
const m___NR_rt_sigprocmask = 135
const m___NR_rt_sigqueueinfo = 138
const m___NR_rt_sigreturn = 139
const m___NR_rt_sigsuspend = 133
const m___NR_rt_sigtimedwait = 137
const m___NR_rt_tgsigqueueinfo = 240
const m___NR_sched_get_priority_max = 125
const m___NR_sched_get_priority_min = 126
const m___NR_sched_getaffinity = 123
const m___NR_sched_getattr = 275
const m___NR_sched_getparam = 121
const m___NR_sched_getscheduler = 120
const m___NR_sched_rr_get_interval = 127
const m___NR_sched_setaffinity = 122
const m___NR_sched_setattr = 274
const m___NR_sched_setparam = 118
const m___NR_sched_setscheduler = 119
const m___NR_sched_yield = 124
const m___NR_seccomp = 277
const m___NR_semctl = 191
const m___NR_semget = 190
const m___NR_semop = 193
const m___NR_semtimedop = 192
const m___NR_sendfile = "__NR3264_sendfile"
const m___NR_sendmmsg = 269
const m___NR_sendmsg = 211
const m___NR_sendto = 206
const m___NR_set_mempolicy = 237
const m___NR_set_mempolicy_home_node = 450
const m___NR_set_robust_list = 99
const m___NR_set_tid_address = 96
const m___NR_setdomainname = 162
const m___NR_setfsgid = 152
const m___NR_setfsuid = 151
const m___NR_setgid = 144
const m___NR_setgroups = 159
const m___NR_sethostname = 161
const m___NR_setitimer = 103
const m___NR_setns = 268
const m___NR_setpgid = 154
const m___NR_setpriority = 140
const m___NR_setregid = 143
const m___NR_setresgid = 149
const m___NR_setresuid = 147
const m___NR_setreuid = 145
const m___NR_setsid = 157
const m___NR_setsockopt = 208
const m___NR_settimeofday = 170
const m___NR_setuid = 146
const m___NR_setxattr = 5
const m___NR_shmat = 196
const m___NR_shmctl = 195
const m___NR_shmdt = 197
const m___NR_shmget = 194
const m___NR_shutdown = 210
const m___NR_sigaltstack = 132
const m___NR_signalfd4 = 74
const m___NR_socket = 198
const m___NR_socketpair = 199
const m___NR_splice = 76
const m___NR_statfs = "__NR3264_statfs"
const m___NR_statx = 291
const m___NR_swapoff = 225
const m___NR_swapon = 224
const m___NR_symlinkat = 36
const m___NR_sync = 81
const m___NR_sync_file_range = 84
const m___NR_syncfs = 267
const m___NR_sysinfo = 179
const m___NR_syslog = 116
const m___NR_tee = 77
const m___NR_tgkill = 131
const m___NR_timer_create = 107
const m___NR_timer_delete = 111
const m___NR_timer_getoverrun = 109
const m___NR_timer_gettime = 108
const m___NR_timer_settime = 110
const m___NR_timerfd_create = 85
const m___NR_timerfd_gettime = 87
const m___NR_timerfd_settime = 86
const m___NR_times = 153
const m___NR_tkill = 130
const m___NR_truncate = "__NR3264_truncate"
const m___NR_umask = 166
const m___NR_umount2 = 39
const m___NR_uname = 160
const m___NR_unlinkat = 35
const m___NR_unshare = 97
const m___NR_userfaultfd = 282
const m___NR_utimensat = 88
const m___NR_vhangup = 58
const m___NR_vmsplice = 75
const m___NR_wait4 = 260
const m___NR_waitid = 95
const m___NR_write = 64
const m___NR_writev = 66
const m_d_fileno = "d_ino"

type Tdirent = struct {
	Fd_ino    Tino_t
	Fd_off    Toff_t
	Fd_reclen uint16
	Fd_type   uint8
	Fd_name   [256]int8
}

func _closefrom_close(tls *libc.TLS, fd int32) {
	libc.Xclose(tls, fd)
}

func _sys_close_range(tls *libc.TLS, fd uint32, max_fd uint32, flags uint32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	return int32(libc.Xsyscall(tls, int64(m_SYS_close_range), libc.VaList(bp+8, fd, max_fd, flags)))
}

// C documentation
//
//	/*
//	 * Close all file descriptors greater than or equal to lowfd.
//	 * This is the expensive (fallback) method.
//	 */
func _closefrom_fallback(tls *libc.TLS, lowfd int32) {
	var fd, maxfd int64
	_, _ = fd, maxfd
	/*
	 * Fall back on sysconf(_SC_OPEN_MAX) or getdtablesize(). This is
	 * equivalent to checking the RLIMIT_NOFILE soft limit. It is
	 * possible for there to be open file descriptors past this limit
	 * but there is not much we can do about that since the hard limit
	 * may be RLIM_INFINITY (LLONG_MAX or ULLONG_MAX on modern systems).
	 */
	maxfd = libc.Xsysconf(tls, int32(m__SC_OPEN_MAX))
	if maxfd < int64(m_OPEN_MAX) {
		maxfd = int64(m_OPEN_MAX)
	}
	/* Make sure we did not get RLIM_INFINITY as the upper limit. */
	if maxfd > int64(m_INT_MAX1) {
		maxfd = int64(m_INT_MAX1)
	}
	fd = int64(lowfd)
	for {
		if !(fd < maxfd) {
			break
		}
		_closefrom_close(tls, int32(fd))
		goto _1
	_1:
		;
		fd++
	}
}

func _closefrom_procfs(tls *libc.TLS, lowfd int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var dent, dirp, fd_array, path, ptr, v1 uintptr
	var fd, fd_array_size, fd_array_used, i, ret, v2 int32
	var _ /* errstr at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _ = dent, dirp, fd, fd_array, fd_array_size, fd_array_used, i, path, ptr, ret, v1, v2
	fd_array = libc.UintptrFromInt32(0)
	fd_array_used = 0
	fd_array_size = 0
	ret = 0
	/* Use /proc/self/fd (or /dev/fd on macOS) if it exists. */
	path = __ccgo_ts
	dirp = libc.Xopendir(tls, path)
	if dirp == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	for {
		v1 = libc.Xreaddir(tls, dirp)
		dent = v1
		if !(v1 != libc.UintptrFromInt32(0)) {
			break
		}
		fd = int32(Xstrtonum(tls, dent+19, int64(lowfd), int64(m_INT_MAX1), bp))
		if *(*uintptr)(unsafe.Pointer(bp)) != libc.UintptrFromInt32(0) || fd == libc.Xdirfd(tls, dirp) {
			continue
		}
		if fd_array_used >= fd_array_size {
			if fd_array_size > 0 {
				fd_array_size *= int32(2)
			} else {
				fd_array_size = int32(32)
			}
			ptr = Xreallocarray(tls, fd_array, libc.Uint64FromInt32(fd_array_size), uint64(4))
			if ptr == libc.UintptrFromInt32(0) {
				ret = -int32(1)
				break
			}
			fd_array = ptr
		}
		v2 = fd_array_used
		fd_array_used++
		*(*int32)(unsafe.Pointer(fd_array + uintptr(v2)*4)) = fd
	}
	i = 0
	for {
		if !(i < fd_array_used) {
			break
		}
		_closefrom_close(tls, *(*int32)(unsafe.Pointer(fd_array + uintptr(i)*4)))
		goto _3
	_3:
		;
		i++
	}
	libc.Xfree(tls, fd_array)
	libc.Xclosedir(tls, dirp)
	return ret
}

// C documentation
//
//	/*
//	 * Close all file descriptors greater than or equal to lowfd.
//	 * We try the fast way first, falling back on the slow method.
//	 */
func Xclosefrom(tls *libc.TLS, lowfd int32) {
	if lowfd < 0 {
		lowfd = 0
	}
	/* Try the fast methods first, if possible. */
	if _sys_close_range(tls, libc.Uint32FromInt32(lowfd), uint32(0xffffffff), uint32(0)) == 0 {
		return
	}
	if _closefrom_procfs(tls, lowfd) != -int32(1) {
		return
	}
	/* Do things the slow way. */
	_closefrom_fallback(tls, lowfd)
}

/*
 * Copyright © 2004-2024 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

func Xvwarnc(tls *libc.TLS, code int32, format uintptr, ap Tva_list) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+14, libc.VaList(bp+8, Xgetprogname(tls)))
	if format != 0 {
		libc.Xvfprintf(tls, libc.Xstderr, format, ap)
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+19, 0)
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+22, libc.VaList(bp+8, libc.Xstrerror(tls, code)))
}

func Xwarnc(tls *libc.TLS, code int32, format uintptr, va uintptr) {
	var ap Tva_list
	_ = ap
	ap = va
	Xvwarnc(tls, code, format, ap)
	_ = ap
}

func Xverrc(tls *libc.TLS, status int32, code int32, format uintptr, ap Tva_list) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+14, libc.VaList(bp+8, Xgetprogname(tls)))
	if format != 0 {
		libc.Xvfprintf(tls, libc.Xstderr, format, ap)
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+19, 0)
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+22, libc.VaList(bp+8, libc.Xstrerror(tls, code)))
	libc.Xexit(tls, status)
}

func Xerrc(tls *libc.TLS, status int32, code int32, format uintptr, va uintptr) {
	var ap Tva_list
	_ = ap
	ap = va
	Xverrc(tls, status, code, format, ap)
	_ = ap
}

const m_PRIX16 = "X"
const m_PRIX32 = "X"
const m_PRIX8 = "X"
const m_PRIXFAST16 = "X"
const m_PRIXFAST32 = "X"
const m_PRIXFAST8 = "X"
const m_PRIXLEAST16 = "X"
const m_PRIXLEAST32 = "X"
const m_PRIXLEAST8 = "X"
const m_PRId16 = "d"
const m_PRId32 = "d"
const m_PRId8 = "d"
const m_PRIdFAST16 = "d"
const m_PRIdFAST32 = "d"
const m_PRIdFAST8 = "d"
const m_PRIdLEAST16 = "d"
const m_PRIdLEAST32 = "d"
const m_PRIdLEAST8 = "d"
const m_PRIi16 = "i"
const m_PRIi32 = "i"
const m_PRIi8 = "i"
const m_PRIiFAST16 = "i"
const m_PRIiFAST32 = "i"
const m_PRIiFAST8 = "i"
const m_PRIiLEAST16 = "i"
const m_PRIiLEAST32 = "i"
const m_PRIiLEAST8 = "i"
const m_PRIo16 = "o"
const m_PRIo32 = "o"
const m_PRIo8 = "o"
const m_PRIoFAST16 = "o"
const m_PRIoFAST32 = "o"
const m_PRIoFAST8 = "o"
const m_PRIoLEAST16 = "o"
const m_PRIoLEAST32 = "o"
const m_PRIoLEAST8 = "o"
const m_PRIu16 = "u"
const m_PRIu32 = "u"
const m_PRIu8 = "u"
const m_PRIuFAST16 = "u"
const m_PRIuFAST32 = "u"
const m_PRIuFAST8 = "u"
const m_PRIuLEAST16 = "u"
const m_PRIuLEAST32 = "u"
const m_PRIuLEAST8 = "u"
const m_PRIx16 = "x"
const m_PRIx32 = "x"
const m_PRIx8 = "x"
const m_PRIxFAST16 = "x"
const m_PRIxFAST32 = "x"
const m_PRIxFAST8 = "x"
const m_PRIxLEAST16 = "x"
const m_PRIxLEAST32 = "x"
const m_PRIxLEAST8 = "x"
const m_SCNd16 = "hd"
const m_SCNd32 = "d"
const m_SCNd8 = "hhd"
const m_SCNdFAST16 = "d"
const m_SCNdFAST32 = "d"
const m_SCNdFAST8 = "hhd"
const m_SCNdLEAST16 = "hd"
const m_SCNdLEAST32 = "d"
const m_SCNdLEAST8 = "hhd"
const m_SCNi16 = "hi"
const m_SCNi32 = "i"
const m_SCNi8 = "hhi"
const m_SCNiFAST16 = "i"
const m_SCNiFAST32 = "i"
const m_SCNiFAST8 = "hhi"
const m_SCNiLEAST16 = "hi"
const m_SCNiLEAST32 = "i"
const m_SCNiLEAST8 = "hhi"
const m_SCNo16 = "ho"
const m_SCNo32 = "o"
const m_SCNo8 = "hho"
const m_SCNoFAST16 = "o"
const m_SCNoFAST32 = "o"
const m_SCNoFAST8 = "hho"
const m_SCNoLEAST16 = "ho"
const m_SCNoLEAST32 = "o"
const m_SCNoLEAST8 = "hho"
const m_SCNu16 = "hu"
const m_SCNu32 = "u"
const m_SCNu8 = "hhu"
const m_SCNuFAST16 = "u"
const m_SCNuFAST32 = "u"
const m_SCNuFAST8 = "hhu"
const m_SCNuLEAST16 = "hu"
const m_SCNuLEAST32 = "u"
const m_SCNuLEAST8 = "hhu"
const m_SCNx16 = "hx"
const m_SCNx32 = "x"
const m_SCNx8 = "hhx"
const m_SCNxFAST16 = "x"
const m_SCNxFAST32 = "x"
const m_SCNxFAST8 = "hhx"
const m_SCNxLEAST16 = "hx"
const m_SCNxLEAST32 = "x"
const m_SCNxLEAST8 = "hhx"
const m___PRI64 = "l"
const m___PRIPTR = "l"

type Timaxdiv_t = struct {
	Fquot Tintmax_t
	Frem  Tintmax_t
}

/* Values for humanize_number(3)'s flags parameter. */

/* Values for humanize_number(3)'s scale parameter. */

/*
 * fparseln() specific operation flags.
 */

// C documentation
//
//	/*
//	 * Convert an expression of the following forms to a uint64_t.
//	 *	1) A positive decimal number.
//	 *	2) A positive decimal number followed by a 'b' or 'B' (mult by 1).
//	 *	3) A positive decimal number followed by a 'k' or 'K' (mult by 1 << 10).
//	 *	4) A positive decimal number followed by a 'm' or 'M' (mult by 1 << 20).
//	 *	5) A positive decimal number followed by a 'g' or 'G' (mult by 1 << 30).
//	 *	6) A positive decimal number followed by a 't' or 'T' (mult by 1 << 40).
//	 *	7) A positive decimal number followed by a 'p' or 'P' (mult by 1 << 50).
//	 *	8) A positive decimal number followed by a 'e' or 'E' (mult by 1 << 60).
//	 */
func Xexpand_number(tls *libc.TLS, buf uintptr, num uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var number Tuint64_t
	var shift uint32
	var _ /* endptr at bp+0 */ uintptr
	_, _ = number, shift
	number = libc.Xstrtoumax(tls, buf, bp, 0)
	if *(*uintptr)(unsafe.Pointer(bp)) == buf {
		/* No valid digits. */
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	}
	switch libc.Xtolower(tls, libc.Int32FromUint8(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))))))) {
	case int32('e'):
		shift = uint32(60)
	case int32('p'):
		shift = uint32(50)
	case int32('t'):
		shift = uint32(40)
	case int32('g'):
		shift = uint32(30)
	case int32('m'):
		shift = uint32(20)
	case int32('k'):
		shift = uint32(10)
	case int32('b'):
		fallthrough
	case int32('\000'): /* No unit. */
		*(*Tuint64_t)(unsafe.Pointer(num)) = number
		return 0
	default:
		/* Unrecognized unit. */
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	}
	if number<<shift>>shift != number {
		/* Overflow */
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ERANGE)
		return -int32(1)
	}
	*(*Tuint64_t)(unsafe.Pointer(num)) = number << shift
	return 0
}

func X__explicit_bzero_hook(tls *libc.TLS, buf uintptr, len1 Tsize_t) {
}

func Xexplicit_bzero(tls *libc.TLS, buf uintptr, len1 Tsize_t) {
	libc.Xmemset(tls, buf, 0, len1)
	X__explicit_bzero_hook(tls, buf, len1)
}

const m_FILEBUF_POOL_ITEMS = 32

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

type Tfilebuf = struct {
	Ffp   uintptr
	Fbuf  uintptr
	Flen1 Tsize_t
}

var _fb_pool [32]Tfilebuf
var _fb_pool_cur int32

func Xfgetln(tls *libc.TLS, stream uintptr, len1 uintptr) (r uintptr) {
	var fb uintptr
	var nread Tssize_t
	_, _ = fb, nread
	libc.Xflockfile(tls, stream)
	/* Try to diminish the possibility of several fgetln() calls being
	 * used on different streams, by using a pool of buffers per file. */
	fb = uintptr(unsafe.Pointer(&_fb_pool)) + uintptr(_fb_pool_cur)*24
	if (*Tfilebuf)(unsafe.Pointer(fb)).Ffp != stream && (*Tfilebuf)(unsafe.Pointer(fb)).Ffp != libc.UintptrFromInt32(0) {
		_fb_pool_cur++
		_fb_pool_cur %= int32(m_FILEBUF_POOL_ITEMS)
		fb = uintptr(unsafe.Pointer(&_fb_pool)) + uintptr(_fb_pool_cur)*24
	}
	(*Tfilebuf)(unsafe.Pointer(fb)).Ffp = stream
	nread = libc.Xgetline(tls, fb+8, fb+16, stream)
	libc.Xfunlockfile(tls, stream)
	/* Note: the getdelim/getline API ensures nread != 0. */
	if nread == int64(-int32(1)) {
		*(*Tsize_t)(unsafe.Pointer(len1)) = uint64(0)
		return libc.UintptrFromInt32(0)
	} else {
		*(*Tsize_t)(unsafe.Pointer(len1)) = libc.Uint64FromInt64(nread)
		return (*Tfilebuf)(unsafe.Pointer(fb)).Fbuf
	}
	return r
}

var _libbsd_emit_link_warning_fgetln = [115]int8{'T', 'h', 'e', ' ', 'f', 'g', 'e', 't', 'l', 'n', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'c', 'a', 'n', 'n', 'o', 't', ' ', 'b', 'e', ' ', 's', 'a', 'f', 'e', 'l', 'y', ' ', 'p', 'o', 'r', 't', 'e', 'd', ',', ' ', 'u', 's', 'e', ' ', 'g', 'e', 't', 'l', 'i', 'n', 'e', '(', '3', ')', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', ',', ' ', 'a', 's', ' ', 'i', 't', ' ', 'i', 's', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', ' ', 'b', 'y', ' ', 'G', 'N', 'U', ' ', 'a', 'n', 'd', ' ', 'P', 'O', 'S', 'I', 'X', '.', '1', '-', '2', '0', '0', '8', '.'}

const m_FILEWBUF_INIT_LEN = 128
const m_FILEWBUF_POOL_ITEMS = 32

type Twint_t = uint32

type Twctype_t = uint64

type Tmbstate_t = struct {
	F__opaque1 uint32
	F__opaque2 uint32
}

type t__mbstate_t = Tmbstate_t

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

type Tfilewbuf = struct {
	Ffp   uintptr
	Fwbuf uintptr
	Flen1 Tsize_t
}

var _fb_pool1 [32]Tfilewbuf
var _fb_pool_cur1 int32

func Xfgetwln(tls *libc.TLS, stream uintptr, lenp uintptr) (r uintptr) {
	var fb, wp, v3 uintptr
	var wc, v1 Twint_t
	var wused, v2 Tsize_t
	_, _, _, _, _, _, _ = fb, wc, wp, wused, v1, v2, v3
	wused = uint64(0)
	/* Try to diminish the possibility of several fgetwln() calls being
	 * used on different streams, by using a pool of buffers per file. */
	fb = uintptr(unsafe.Pointer(&_fb_pool1)) + uintptr(_fb_pool_cur1)*24
	if (*Tfilewbuf)(unsafe.Pointer(fb)).Ffp != stream && (*Tfilewbuf)(unsafe.Pointer(fb)).Ffp != libc.UintptrFromInt32(0) {
		_fb_pool_cur1++
		_fb_pool_cur1 %= int32(m_FILEWBUF_POOL_ITEMS)
		fb = uintptr(unsafe.Pointer(&_fb_pool1)) + uintptr(_fb_pool_cur1)*24
	}
	(*Tfilewbuf)(unsafe.Pointer(fb)).Ffp = stream
	for {
		v1 = libc.Xfgetwc(tls, stream)
		wc = v1
		if !(v1 != uint32(0xffffffff)) {
			break
		}
		if !((*Tfilewbuf)(unsafe.Pointer(fb)).Flen1 != 0) || wused >= (*Tfilewbuf)(unsafe.Pointer(fb)).Flen1 {
			if (*Tfilewbuf)(unsafe.Pointer(fb)).Flen1 != 0 {
				*(*Tsize_t)(unsafe.Pointer(fb + 16)) *= uint64(2)
			} else {
				(*Tfilewbuf)(unsafe.Pointer(fb)).Flen1 = uint64(m_FILEWBUF_INIT_LEN)
			}
			wp = Xreallocarray(tls, (*Tfilewbuf)(unsafe.Pointer(fb)).Fwbuf, (*Tfilewbuf)(unsafe.Pointer(fb)).Flen1, uint64(4))
			if wp == libc.UintptrFromInt32(0) {
				wused = uint64(0)
				break
			}
			(*Tfilewbuf)(unsafe.Pointer(fb)).Fwbuf = wp
		}
		v2 = wused
		wused++
		*(*Twchar_t)(unsafe.Pointer((*Tfilewbuf)(unsafe.Pointer(fb)).Fwbuf + uintptr(v2)*4)) = libc.Int32FromUint32(wc)
		if wc == uint32('\n') {
			break
		}
	}
	*(*Tsize_t)(unsafe.Pointer(lenp)) = wused
	if wused != 0 {
		v3 = (*Tfilewbuf)(unsafe.Pointer(fb)).Fwbuf
	} else {
		v3 = libc.UintptrFromInt32(0)
	}
	return v3
}

// C documentation
//
//	/* XXX: Ideally we'd recommend getwline(3), but unfortunately even though it
//	 * was part of the ISO/IEC TR 24731-2:2010 draft, it did not make it into C11
//	 * and is not widely implemented. */
var _libbsd_emit_link_warning_fgetwln = [115]int8{'T', 'h', 'e', ' ', 'f', 'g', 'e', 't', 'w', 'l', 'n', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'c', 'a', 'n', 'n', 'o', 't', ' ', 'b', 'e', ' ', 's', 'a', 'f', 'e', 'l', 'y', ' ', 'p', 'o', 'r', 't', 'e', 'd', ',', ' ', 'u', 's', 'e', ' ', 'f', 'g', 'e', 't', 'w', 'c', '(', '3', ')', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', ',', ' ', 'a', 's', ' ', 'i', 't', ' ', 'i', 's', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', ' ', 'b', 'y', ' ', 'C', '9', '9', ' ', 'a', 'n', 'd', ' ', 'P', 'O', 'S', 'I', 'X', '.', '1', '-', '2', '0', '0', '1', '.'}

const m_FPARSELN_UNESCALL1 = 15
const m_FPARSELN_UNESCCOMM1 = 4
const m_FPARSELN_UNESCCONT1 = 2
const m_FPARSELN_UNESCESC1 = 1
const m_FPARSELN_UNESCREST1 = 8
const m_static_assert = "_Static_assert"

// C documentation
//
//	/* isescaped():
//	 *	Return true if the character in *p that belongs to a string
//	 *	that starts in *sp, is escaped by the escape character esc.
//	 */
func _isescaped(tls *libc.TLS, sp uintptr, p uintptr, esc int32) (r int32) {
	var cp, v2 uintptr
	var ne Tsize_t
	_, _, _ = cp, ne, v2
	/* No escape character */
	if esc == int32('\000') {
		return 0
	}
	/* Count the number of escape characters that precede ours */
	ne = uint64(0)
	cp = p
	for {
		cp--
		v2 = cp
		if !(v2 >= sp && int32(*(*int8)(unsafe.Pointer(cp))) == esc) {
			break
		}
		goto _1
		goto _1
	_1:
		;
		ne++
	}
	/* Return true if odd number of escape characters */
	return libc.BoolInt32(ne&uint64(1) != uint64(0))
}

// C documentation
//
//	/* fparseln():
//	 *	Read a line from a file parsing continuations ending in  *	and eliminating trailing newlines, or comments starting with
//	 *	the comment char.
//	 */
func Xfparseln(tls *libc.TLS, fp uintptr, size uintptr, lineno uintptr, str uintptr, flags int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var buf, cp, v2, v3, v4, v5, v6, v7, v8, v9 uintptr
	var cnt, skipesc int32
	var com, con, esc, nl int8
	var len1 Tsize_t
	var s Tssize_t
	var _ /* ptr at bp+8 */ uintptr
	var _ /* ptrlen at bp+0 */ Tsize_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = buf, cnt, com, con, cp, esc, len1, nl, s, skipesc, v2, v3, v4, v5, v6, v7, v8, v9
	len1 = uint64(0)
	buf = libc.UintptrFromInt32(0)
	*(*Tsize_t)(unsafe.Pointer(bp)) = uint64(0)
	*(*uintptr)(unsafe.Pointer(bp + 8)) = libc.UintptrFromInt32(0)
	cnt = int32(1)
	if str == libc.UintptrFromInt32(0) {
		str = uintptr(unsafe.Pointer(&_dstr))
	}
	esc = *(*int8)(unsafe.Pointer(str))
	con = *(*int8)(unsafe.Pointer(str + 1))
	com = *(*int8)(unsafe.Pointer(str + 2))
	/*
	 * XXX: it would be cool to be able to specify the newline character,
	 * getdelim(3) does let us, but supporting it would diverge from BSDs.
	 */
	nl = int8('\n')
	libc.Xflockfile(tls, fp)
	for cnt != 0 {
		cnt = 0
		if lineno != 0 {
			*(*Tsize_t)(unsafe.Pointer(lineno))++
		}
		s = libc.Xgetline(tls, bp+8, bp, fp)
		if s < 0 {
			break
		}
		if s != 0 && com != 0 { /* Check and eliminate comments */
			cp = *(*uintptr)(unsafe.Pointer(bp + 8))
			for {
				if !(cp < *(*uintptr)(unsafe.Pointer(bp + 8))+uintptr(s)) {
					break
				}
				if int32(*(*int8)(unsafe.Pointer(cp))) == int32(com) && !(_isescaped(tls, *(*uintptr)(unsafe.Pointer(bp + 8)), cp, int32(esc)) != 0) {
					s = int64(cp) - int64(*(*uintptr)(unsafe.Pointer(bp + 8)))
					cnt = libc.BoolInt32(s == 0 && buf == libc.UintptrFromInt32(0))
					break
				}
				goto _1
			_1:
				;
				cp++
			}
		}
		if s != 0 && nl != 0 { /* Check and eliminate newlines */
			cp = *(*uintptr)(unsafe.Pointer(bp + 8)) + uintptr(s-int64(1))
			if int32(*(*int8)(unsafe.Pointer(cp))) == int32(nl) {
				s--
			} /* forget newline */
		}
		if s != 0 && con != 0 { /* Check and eliminate continuations */
			cp = *(*uintptr)(unsafe.Pointer(bp + 8)) + uintptr(s-int64(1))
			if int32(*(*int8)(unsafe.Pointer(cp))) == int32(con) && !(_isescaped(tls, *(*uintptr)(unsafe.Pointer(bp + 8)), cp, int32(esc)) != 0) {
				s-- /* forget continuation char */
				cnt = int32(1)
			}
		}
		if s == 0 {
			/*
			 * nothing to add, skip realloc except in case
			 * we need a minimal buf to return an empty line
			 */
			if cnt != 0 || buf != libc.UintptrFromInt32(0) {
				continue
			}
		}
		v2 = libc.Xrealloc(tls, buf, len1+libc.Uint64FromInt64(s)+uint64(1))
		cp = v2
		if v2 == libc.UintptrFromInt32(0) {
			libc.Xfunlockfile(tls, fp)
			libc.Xfree(tls, buf)
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
			return libc.UintptrFromInt32(0)
		}
		buf = cp
		libc.Xmemcpy(tls, buf+uintptr(len1), *(*uintptr)(unsafe.Pointer(bp + 8)), libc.Uint64FromInt64(s))
		len1 += libc.Uint64FromInt64(s)
		*(*int8)(unsafe.Pointer(buf + uintptr(len1))) = int8('\000')
	}
	libc.Xfunlockfile(tls, fp)
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
	if flags&int32(m_FPARSELN_UNESCALL1) != 0 && esc != 0 && buf != libc.UintptrFromInt32(0) && libc.Xstrchr(tls, buf, int32(esc)) != libc.UintptrFromInt32(0) {
		v3 = buf
		cp = v3
		*(*uintptr)(unsafe.Pointer(bp + 8)) = v3
		for int32(*(*int8)(unsafe.Pointer(cp))) != int32('\000') {
			for int32(*(*int8)(unsafe.Pointer(cp))) != int32('\000') && int32(*(*int8)(unsafe.Pointer(cp))) != int32(esc) {
				v4 = *(*uintptr)(unsafe.Pointer(bp + 8))
				*(*uintptr)(unsafe.Pointer(bp + 8))++
				v5 = cp
				cp++
				*(*int8)(unsafe.Pointer(v4)) = *(*int8)(unsafe.Pointer(v5))
			}
			if int32(*(*int8)(unsafe.Pointer(cp))) == int32('\000') || int32(*(*int8)(unsafe.Pointer(cp + 1))) == int32('\000') {
				break
			}
			skipesc = 0
			if int32(*(*int8)(unsafe.Pointer(cp + 1))) == int32(com) {
				skipesc += flags & int32(m_FPARSELN_UNESCCOMM1)
			}
			if int32(*(*int8)(unsafe.Pointer(cp + 1))) == int32(con) {
				skipesc += flags & int32(m_FPARSELN_UNESCCONT1)
			}
			if int32(*(*int8)(unsafe.Pointer(cp + 1))) == int32(esc) {
				skipesc += flags & int32(m_FPARSELN_UNESCESC1)
			}
			if int32(*(*int8)(unsafe.Pointer(cp + 1))) != int32(com) && int32(*(*int8)(unsafe.Pointer(cp + 1))) != int32(con) && int32(*(*int8)(unsafe.Pointer(cp + 1))) != int32(esc) {
				skipesc = flags & int32(m_FPARSELN_UNESCREST1)
			}
			if skipesc != 0 {
				cp++
			} else {
				v6 = *(*uintptr)(unsafe.Pointer(bp + 8))
				*(*uintptr)(unsafe.Pointer(bp + 8))++
				v7 = cp
				cp++
				*(*int8)(unsafe.Pointer(v6)) = *(*int8)(unsafe.Pointer(v7))
			}
			v8 = *(*uintptr)(unsafe.Pointer(bp + 8))
			*(*uintptr)(unsafe.Pointer(bp + 8))++
			v9 = cp
			cp++
			*(*int8)(unsafe.Pointer(v8)) = *(*int8)(unsafe.Pointer(v9))
		}
		*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))) = int8('\000')
		len1 = libc.Xstrlen(tls, buf)
	}
	if size != 0 {
		*(*Tsize_t)(unsafe.Pointer(size)) = len1
	}
	return buf
}

var _dstr = [3]int8{
	0: int8('\\'),
	1: int8('\\'),
	2: int8('#'),
}

const m_FPARSELN_UNESCALL2 = 0x0f
const m_FPARSELN_UNESCCOMM2 = 0x04
const m_FPARSELN_UNESCCONT2 = 0x02
const m_FPARSELN_UNESCESC2 = 0x01
const m_FPARSELN_UNESCREST2 = 0x08
const m_LOCK_EX = 2
const m_LOCK_NB = 4
const m_LOCK_SH = 1
const m_LOCK_UN = 8
const m_O_CREAT1 = 64
const m_O_NONBLOCK1 = 2048
const m_O_TRUNC1 = 512

/* Values for humanize_number(3)'s flags parameter. */

/* Values for humanize_number(3)'s scale parameter. */

/*
 * fparseln() specific operation flags.
 */

func _lock_file(tls *libc.TLS, fd int32, flags int32) (r int32) {
	var operation int32
	_ = operation
	operation = int32(m_LOCK_EX)
	if flags&int32(m_O_NONBLOCK1) != 0 {
		operation |= int32(m_LOCK_NB)
	}
	return libc.Xflock(tls, fd, operation)
}

// C documentation
//
//	/*
//	 * Reliably open and lock a file.
//	 *
//	 * Please do not modify this code without first reading the revision history
//	 * and discussing your changes with <des@freebsd.org>.  Don't be fooled by the
//	 * code's apparent simplicity; there would be no need for this function if it
//	 * was easy to get right.
//	 */
func _vflopenat(tls *libc.TLS, dirfd int32, path uintptr, flags int32, ap Tva_list) (r int32) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var fd, serrno, trunc, v2 int32
	var mode Tmode_t
	var _ /* fsb at bp+128 */ Tstat
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _ = fd, mode, serrno, trunc, v2
	mode = uint32(0)
	if flags&int32(m_O_CREAT1) != 0 {
		mode = libc.Uint32FromInt32(libc.VaInt32(&ap)) /* mode_t promoted to int */
	}
	trunc = flags & int32(m_O_TRUNC1)
	flags &= ^libc.Int32FromInt32(m_O_TRUNC1)
	for {
		v2 = libc.Xopenat(tls, dirfd, path, flags, libc.VaList(bp+264, mode))
		fd = v2
		if v2 == -int32(1) {
			/* non-existent or no access */
			return -int32(1)
		}
		if _lock_file(tls, fd, flags) == -int32(1) {
			/* unsupported or interrupted */
			serrno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			libc.Xclose(tls, fd)
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = serrno
			return -int32(1)
		}
		if libc.Xfstatat(tls, dirfd, path, bp, 0) == -int32(1) {
			/* disappeared from under our feet */
			libc.Xclose(tls, fd)
			goto _1
		}
		if libc.Xfstat(tls, fd, bp+128) == -int32(1) {
			/* can't happen [tm] */
			serrno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			libc.Xclose(tls, fd)
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = serrno
			return -int32(1)
		}
		if (*(*Tstat)(unsafe.Pointer(bp))).Fst_dev != (*(*Tstat)(unsafe.Pointer(bp + 128))).Fst_dev || (*(*Tstat)(unsafe.Pointer(bp))).Fst_ino != (*(*Tstat)(unsafe.Pointer(bp + 128))).Fst_ino {
			/* changed under our feet */
			libc.Xclose(tls, fd)
			goto _1
		}
		if trunc != 0 && libc.Xftruncate(tls, fd, 0) != 0 {
			/* can't happen [tm] */
			serrno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			libc.Xclose(tls, fd)
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = serrno
			return -int32(1)
		}
		/*
		 * The following change is provided as a specific example to
		 * avoid.
		 */
		return fd
		goto _1
	_1:
	}
	return r
}

func Xflopen(tls *libc.TLS, path uintptr, flags int32, va uintptr) (r int32) {
	var ap Tva_list
	var ret int32
	_, _ = ap, ret
	ap = va
	ret = _vflopenat(tls, -int32(100), path, flags, ap)
	_ = ap
	return ret
}

func Xflopenat(tls *libc.TLS, dirfd int32, path uintptr, flags int32, va uintptr) (r int32) {
	var ap Tva_list
	var ret int32
	_, _ = ap, ret
	ap = va
	ret = _vflopenat(tls, dirfd, path, flags, ap)
	_ = ap
	return ret
}

type ___e_fmtcheck_types = int32

const _FMTCHECK_START = 0
const _FMTCHECK_SHORT = 1
const _FMTCHECK_INT = 2
const _FMTCHECK_WINTT = 3
const _FMTCHECK_LONG = 4
const _FMTCHECK_QUAD = 5
const _FMTCHECK_INTMAXT = 6
const _FMTCHECK_PTRDIFFT = 7
const _FMTCHECK_SIZET = 8
const _FMTCHECK_POINTER = 9
const _FMTCHECK_CHARPOINTER = 10
const _FMTCHECK_SHORTPOINTER = 11
const _FMTCHECK_INTPOINTER = 12
const _FMTCHECK_LONGPOINTER = 13
const _FMTCHECK_QUADPOINTER = 14
const _FMTCHECK_INTMAXTPOINTER = 15
const _FMTCHECK_PTRDIFFTPOINTER = 16
const _FMTCHECK_SIZETPOINTER = 17
const _FMTCHECK_DOUBLE = 18
const _FMTCHECK_LONGDOUBLE = 19
const _FMTCHECK_STRING = 20
const _FMTCHECK_WSTRING = 21
const _FMTCHECK_WIDTH = 22
const _FMTCHECK_PRECISION = 23
const _FMTCHECK_DONE = 24
const _FMTCHECK_UNKNOWN = 25

type TEFT = int32

type _e_modifier = int32

const _MOD_NONE = 0
const _MOD_CHAR = 1
const _MOD_SHORT = 2
const _MOD_LONG = 3
const _MOD_QUAD = 4
const _MOD_INTMAXT = 5
const _MOD_LONGDOUBLE = 6
const _MOD_PTRDIFFT = 7
const _MOD_SIZET = 8

func _get_next_format_from_precision(tls *libc.TLS, pf uintptr) (r TEFT) {
	var f uintptr
	var modifier _e_modifier
	_, _ = f, modifier
	f = *(*uintptr)(unsafe.Pointer(pf))
	switch int32(*(*int8)(unsafe.Pointer(f))) {
	case int32('h'):
		f++
		if !(*(*int8)(unsafe.Pointer(f)) != 0) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		if int32(*(*int8)(unsafe.Pointer(f))) == int32('h') {
			f++
			modifier = int32(_MOD_CHAR)
		} else {
			modifier = int32(_MOD_SHORT)
		}
	case int32('j'):
		f++
		modifier = int32(_MOD_INTMAXT)
	case int32('l'):
		f++
		if !(*(*int8)(unsafe.Pointer(f)) != 0) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		if int32(*(*int8)(unsafe.Pointer(f))) == int32('l') {
			f++
			modifier = int32(_MOD_QUAD)
		} else {
			modifier = int32(_MOD_LONG)
		}
	case int32('q'):
		f++
		modifier = int32(_MOD_QUAD)
	case int32('t'):
		f++
		modifier = int32(_MOD_PTRDIFFT)
	case int32('z'):
		f++
		modifier = int32(_MOD_SIZET)
	case int32('L'):
		f++
		modifier = int32(_MOD_LONGDOUBLE)
	default:
		modifier = int32(_MOD_NONE)
		break
	}
	if !(*(*int8)(unsafe.Pointer(f)) != 0) {
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	}
	if libc.Xstrchr(tls, __ccgo_ts+26, int32(*(*int8)(unsafe.Pointer(f)))) != 0 {
		switch modifier {
		case int32(_MOD_LONG):
			goto _1
		case int32(_MOD_QUAD):
			goto _2
		case int32(_MOD_INTMAXT):
			goto _3
		case int32(_MOD_PTRDIFFT):
			goto _4
		case int32(_MOD_SIZET):
			goto _5
		case int32(_MOD_NONE):
			goto _6
		case int32(_MOD_SHORT):
			goto _7
		case int32(_MOD_CHAR):
			goto _8
		default:
			goto _9
		}
		goto _10
	_1:
		;
	_13:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_LONG)
		goto _12
	_12:
		;
		if 0 != 0 {
			goto _13
		}
		goto _11
	_11:
		;
	_2:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_QUAD)
	_3:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_INTMAXT)
	_4:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_PTRDIFFT)
	_5:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_SIZET)
	_8:
		;
	_7:
		;
	_6:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_INT)
	_9:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	_10:
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('n') {
		switch modifier {
		case int32(_MOD_CHAR):
			goto _14
		case int32(_MOD_SHORT):
			goto _15
		case int32(_MOD_LONG):
			goto _16
		case int32(_MOD_QUAD):
			goto _17
		case int32(_MOD_INTMAXT):
			goto _18
		case int32(_MOD_PTRDIFFT):
			goto _19
		case int32(_MOD_SIZET):
			goto _20
		case int32(_MOD_NONE):
			goto _21
		default:
			goto _22
		}
		goto _23
	_14:
		;
	_26:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_CHARPOINTER)
		goto _25
	_25:
		;
		if 0 != 0 {
			goto _26
		}
		goto _24
	_24:
		;
	_15:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_SHORTPOINTER)
	_16:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_LONGPOINTER)
	_17:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_QUADPOINTER)
	_18:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_INTMAXTPOINTER)
	_19:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_PTRDIFFTPOINTER)
	_20:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_SIZETPOINTER)
	_21:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_INTPOINTER)
	_22:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	_23:
	}
	if libc.Xstrchr(tls, __ccgo_ts+33, int32(*(*int8)(unsafe.Pointer(f)))) != 0 {
		if modifier != int32(_MOD_NONE) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_LONG)
	}
	if libc.Xstrchr(tls, __ccgo_ts+37, int32(*(*int8)(unsafe.Pointer(f)))) != 0 {
		switch modifier {
		case int32(_MOD_LONGDOUBLE):
			goto _27
		case int32(_MOD_NONE):
			goto _28
		case int32(_MOD_LONG):
			goto _29
		default:
			goto _30
		}
		goto _31
	_27:
		;
	_34:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_LONGDOUBLE)
		goto _33
	_33:
		;
		if 0 != 0 {
			goto _34
		}
		goto _32
	_32:
		;
	_29:
		;
	_28:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_DOUBLE)
	_30:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	_31:
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('c') {
		switch modifier {
		case int32(_MOD_LONG):
			goto _35
		case int32(_MOD_NONE):
			goto _36
		default:
			goto _37
		}
		goto _38
	_35:
		;
	_41:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_WINTT)
		goto _40
	_40:
		;
		if 0 != 0 {
			goto _41
		}
		goto _39
	_39:
		;
	_36:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_INT)
	_37:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	_38:
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('C') {
		if modifier != int32(_MOD_NONE) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_WINTT)
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('s') {
		switch modifier {
		case int32(_MOD_LONG):
			goto _42
		case int32(_MOD_NONE):
			goto _43
		default:
			goto _44
		}
		goto _45
	_42:
		;
	_48:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_WSTRING)
		goto _47
	_47:
		;
		if 0 != 0 {
			goto _48
		}
		goto _46
	_46:
		;
	_43:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_STRING)
	_44:
		;
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	_45:
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('S') {
		if modifier != int32(_MOD_NONE) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_WSTRING)
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('p') {
		if modifier != int32(_MOD_NONE) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_POINTER)
	}
	*(*uintptr)(unsafe.Pointer(pf)) = f
	return int32(_FMTCHECK_UNKNOWN)
	/*NOTREACHED*/
	return r
}

func _get_next_format_from_width(tls *libc.TLS, pf uintptr) (r TEFT) {
	var f uintptr
	_ = f
	f = *(*uintptr)(unsafe.Pointer(pf))
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('.') {
		f++
		if int32(*(*int8)(unsafe.Pointer(f))) == int32('*') {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_PRECISION)
		}
		/* eat any precision (empty is allowed) */
		for libc.BoolInt32(uint32(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(f))))-uint32('0') < uint32(10)) != 0 {
			f++
		}
		if !(*(*int8)(unsafe.Pointer(f)) != 0) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
	}
	*(*uintptr)(unsafe.Pointer(pf)) = f
	return _get_next_format_from_precision(tls, pf)
	/*NOTREACHED*/
	return r
}

func _get_next_format(tls *libc.TLS, pf uintptr, eft TEFT) (r TEFT) {
	var f uintptr
	var infmt int32
	_, _ = f, infmt
	if eft == int32(_FMTCHECK_WIDTH) {
		*(*uintptr)(unsafe.Pointer(pf))++
		return _get_next_format_from_width(tls, pf)
	} else {
		if eft == int32(_FMTCHECK_PRECISION) {
			*(*uintptr)(unsafe.Pointer(pf))++
			return _get_next_format_from_precision(tls, pf)
		}
	}
	f = *(*uintptr)(unsafe.Pointer(pf))
	infmt = 0
	for !(infmt != 0) {
		f = libc.Xstrchr(tls, f, int32('%'))
		if f == libc.UintptrFromInt32(0) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_DONE)
		}
		f++
		if !(*(*int8)(unsafe.Pointer(f)) != 0) {
			*(*uintptr)(unsafe.Pointer(pf)) = f
			return int32(_FMTCHECK_UNKNOWN)
		}
		if int32(*(*int8)(unsafe.Pointer(f))) != int32('%') {
			infmt = int32(1)
		} else {
			f++
		}
	}
	/* Eat any of the flags */
	for *(*int8)(unsafe.Pointer(f)) != 0 && libc.Xstrchr(tls, __ccgo_ts+46, int32(*(*int8)(unsafe.Pointer(f)))) != 0 {
		f++
	}
	if int32(*(*int8)(unsafe.Pointer(f))) == int32('*') {
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_WIDTH)
	}
	/* eat any width */
	for libc.BoolInt32(uint32(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(f))))-uint32('0') < uint32(10)) != 0 {
		f++
	}
	if !(*(*int8)(unsafe.Pointer(f)) != 0) {
		*(*uintptr)(unsafe.Pointer(pf)) = f
		return int32(_FMTCHECK_UNKNOWN)
	}
	*(*uintptr)(unsafe.Pointer(pf)) = f
	return _get_next_format_from_width(tls, pf)
	/*NOTREACHED*/
	return r
}

func Xfmtcheck(tls *libc.TLS, f1 uintptr, f2 uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var f1t, f2t, v1 TEFT
	var _ /* f1p at bp+0 */ uintptr
	var _ /* f2p at bp+8 */ uintptr
	_, _, _ = f1t, f2t, v1
	if !(f1 != 0) {
		return f2
	}
	*(*uintptr)(unsafe.Pointer(bp)) = f1
	f1t = int32(_FMTCHECK_START)
	*(*uintptr)(unsafe.Pointer(bp + 8)) = f2
	f2t = int32(_FMTCHECK_START)
	for {
		v1 = _get_next_format(tls, bp, f1t)
		f1t = v1
		if !(v1 != int32(_FMTCHECK_DONE)) {
			break
		}
		if f1t == int32(_FMTCHECK_UNKNOWN) {
			return f2
		}
		f2t = _get_next_format(tls, bp+8, f2t)
		if f1t != f2t {
			return f2
		}
	}
	return f1
}

const m_FSETLOCKING_BYCALLER = 2
const m_FSETLOCKING_INTERNAL = 1
const m_FSETLOCKING_QUERY = 0

func Xfpurge(tls *libc.TLS, fp uintptr) (r int32) {
	if fp == libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EBADF)
		return -int32(1)
	}
	libc.X__fpurge(tls, fp)
	return 0
}

func Xfreezero(tls *libc.TLS, ptr uintptr, sz Tsize_t) {
	/* This is legal. */
	if ptr == libc.UintptrFromInt32(0) {
		return
	}
	Xexplicit_bzero(tls, ptr, sz)
	libc.Xfree(tls, ptr)
}

/*
 * Copyright © 2004-2024 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

type Tfunopen_cookie = struct {
	Forig_cookie uintptr
	Freadfn      uintptr
	Fwritefn     uintptr
	Fseekfn      uintptr
	Fclosefn     uintptr
}

func _funopen_read(tls *libc.TLS, cookie uintptr, buf uintptr, size Tsize_t) (r Tssize_t) {
	var cookiewrap uintptr
	_ = cookiewrap
	cookiewrap = cookie
	if (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Freadfn == libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EBADF)
		return int64(-int32(1))
	}
	return int64((*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Freadfn})))(tls, (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Forig_cookie, buf, libc.Int32FromUint64(size)))
}

func _funopen_write(tls *libc.TLS, cookie uintptr, buf uintptr, size Tsize_t) (r Tssize_t) {
	var cookiewrap uintptr
	_ = cookiewrap
	cookiewrap = cookie
	if (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fwritefn == libc.UintptrFromInt32(0) {
		return int64(-libc.Int32FromInt32(1))
	}
	return int64((*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fwritefn})))(tls, (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Forig_cookie, buf, libc.Int32FromUint64(size)))
}

func _funopen_seek(tls *libc.TLS, cookie uintptr, offset uintptr, whence int32) (r int32) {
	var cookiewrap uintptr
	var soff Toff_t
	_, _ = cookiewrap, soff
	cookiewrap = cookie
	soff = *(*Toff_t)(unsafe.Pointer(offset))
	if (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fseekfn == libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ESPIPE)
		return -int32(1)
	}
	soff = (*(*func(*libc.TLS, uintptr, Toff_t, int32) Toff_t)(unsafe.Pointer(&struct{ uintptr }{(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fseekfn})))(tls, (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Forig_cookie, soff, whence)
	*(*Toff_t)(unsafe.Pointer(offset)) = soff
	return int32(*(*Toff_t)(unsafe.Pointer(offset)))
}

func _funopen_close(tls *libc.TLS, cookie uintptr) (r int32) {
	var cookiewrap uintptr
	var rc int32
	_, _ = cookiewrap, rc
	cookiewrap = cookie
	if (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fclosefn != 0 {
		rc = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fclosefn})))(tls, (*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Forig_cookie)
	} else {
		rc = 0
	}
	libc.Xfree(tls, cookiewrap)
	return rc
}

func Xfunopen(tls *libc.TLS, cookie uintptr, readfn uintptr, writefn uintptr, seekfn uintptr, closefn uintptr) (r uintptr) {
	var cookiewrap, mode uintptr
	var funcswrap Tcookie_io_functions_t
	_, _, _ = cookiewrap, funcswrap, mode
	funcswrap = Tcookie_io_functions_t{
		Fread:   __ccgo_fp(_funopen_read),
		Fwrite:  __ccgo_fp(_funopen_write),
		Fseek:   __ccgo_fp(_funopen_seek),
		Fclose1: __ccgo_fp(_funopen_close),
	}
	if readfn != 0 {
		if writefn == libc.UintptrFromInt32(0) {
			mode = __ccgo_ts + 53
		} else {
			mode = __ccgo_ts + 55
		}
	} else {
		if writefn != 0 {
			mode = __ccgo_ts + 58
		} else {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
			return libc.UintptrFromInt32(0)
		}
	}
	cookiewrap = libc.Xmalloc(tls, uint64(40))
	if cookiewrap == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Forig_cookie = cookie
	(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Freadfn = readfn
	(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fwritefn = writefn
	(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fseekfn = seekfn
	(*Tfunopen_cookie)(unsafe.Pointer(cookiewrap)).Fclosefn = closefn
	return libc.Xfopencookie(tls, cookiewrap, mode, funcswrap)
}

const m_KB = 1024
const m_MAXB = "GB"

func Xgetbsize(tls *libc.TLS, headerlenp uintptr, blocksizep uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var blocksize, max, mul, n, v15, v16, v17, v4 int64
	var form, p, v3 uintptr
	var _ /* ep at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _ = blocksize, form, max, mul, n, p, v15, v16, v17, v3, v4
	form = __ccgo_ts + 60
	v3 = libc.Xgetenv(tls, __ccgo_ts+61)
	p = v3
	if !(v3 != libc.UintptrFromInt32(0) && int32(*(*int8)(unsafe.Pointer(p))) != int32('\000')) {
		goto _1
	}
	v4 = libc.Xstrtol(tls, p, bp, int32(10))
	n = v4
	if v4 < 0 {
		goto underflow
	}
	if n == 0 {
		n = int64(1)
	}
	if *(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))) != 0 && *(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 1)) != 0 {
		goto fmterr
	}
	switch int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))))) {
	case int32('g'):
		goto _5
	case int32('G'):
		goto _6
	case int32('k'):
		goto _7
	case int32('K'):
		goto _8
	case int32('m'):
		goto _9
	case int32('M'):
		goto _10
	case int32('\000'):
		goto _11
	default:
		goto _12
	}
	goto _13
_6:
	;
_5:
	;
	form = __ccgo_ts + 71
	max = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) / (libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024))
	mul = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024)
	goto _13
_8:
	;
_7:
	;
	form = __ccgo_ts + 73
	max = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) / libc.Int64FromInt64(1024)
	mul = int64(1024)
	goto _13
_10:
	;
_9:
	;
	form = __ccgo_ts + 75
	max = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) / (libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024))
	mul = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024)
	goto _13
_11:
	;
	max = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024)
	mul = int64(1)
	goto _13
_12:
	;
	goto fmterr
fmterr:
	;
	libc.Xwarnx(tls, __ccgo_ts+77, libc.VaList(bp+16, p))
	n = int64(512)
	max = libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024) * libc.Int64FromInt64(1024)
	mul = int64(1)
	goto _13
_13:
	;
	if n > max {
		libc.Xwarnx(tls, __ccgo_ts+99, libc.VaList(bp+16, libc.Int64FromInt64(1024)*libc.Int64FromInt64(1024)*libc.Int64FromInt64(1024)/(libc.Int64FromInt64(1024)*libc.Int64FromInt64(1024)*libc.Int64FromInt64(1024))))
		n = max
	}
	v15 = n * mul
	blocksize = v15
	if !(v15 < int64(512)) {
		goto _14
	}
	goto underflow
underflow:
	;
	libc.Xwarnx(tls, __ccgo_ts+125, 0)
	form = __ccgo_ts + 60
	v16 = libc.Int64FromInt32(512)
	n = v16
	blocksize = v16
_14:
	;
	goto _2
_1:
	;
	v17 = libc.Int64FromInt32(512)
	n = v17
	blocksize = v17
_2:
	;
	libc.X__builtin_snprintf(tls, uintptr(unsafe.Pointer(&_header)), uint64(20), __ccgo_ts+150, libc.VaList(bp+16, n, form))
	*(*int32)(unsafe.Pointer(headerlenp)) = libc.Int32FromUint64(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_header))))
	*(*int64)(unsafe.Pointer(blocksizep)) = blocksize
	return uintptr(unsafe.Pointer(&_header))
}

var _header [20]int8

const m_AF_ALG = "PF_ALG"
const m_AF_APPLETALK = "PF_APPLETALK"
const m_AF_ASH = "PF_ASH"
const m_AF_ATMPVC = "PF_ATMPVC"
const m_AF_ATMSVC = "PF_ATMSVC"
const m_AF_AX25 = "PF_AX25"
const m_AF_BLUETOOTH = "PF_BLUETOOTH"
const m_AF_BRIDGE = "PF_BRIDGE"
const m_AF_CAIF = "PF_CAIF"
const m_AF_CAN = "PF_CAN"
const m_AF_DECnet = "PF_DECnet"
const m_AF_ECONET = "PF_ECONET"
const m_AF_FILE = "AF_LOCAL"
const m_AF_IB = "PF_IB"
const m_AF_IEEE802154 = "PF_IEEE802154"
const m_AF_INET = "PF_INET"
const m_AF_INET6 = "PF_INET6"
const m_AF_IPX = "PF_IPX"
const m_AF_IRDA = "PF_IRDA"
const m_AF_ISDN = "PF_ISDN"
const m_AF_IUCV = "PF_IUCV"
const m_AF_KCM = "PF_KCM"
const m_AF_KEY = "PF_KEY"
const m_AF_LLC = "PF_LLC"
const m_AF_LOCAL = "PF_LOCAL"
const m_AF_MAX = "PF_MAX"
const m_AF_MPLS = "PF_MPLS"
const m_AF_NETBEUI = "PF_NETBEUI"
const m_AF_NETLINK = "PF_NETLINK"
const m_AF_NETROM = "PF_NETROM"
const m_AF_NFC = "PF_NFC"
const m_AF_PACKET = "PF_PACKET"
const m_AF_PHONET = "PF_PHONET"
const m_AF_PPPOX = "PF_PPPOX"
const m_AF_QIPCRTR = "PF_QIPCRTR"
const m_AF_RDS = "PF_RDS"
const m_AF_ROSE = "PF_ROSE"
const m_AF_ROUTE = "PF_ROUTE"
const m_AF_RXRPC = "PF_RXRPC"
const m_AF_SECURITY = "PF_SECURITY"
const m_AF_SMC = "PF_SMC"
const m_AF_SNA = "PF_SNA"
const m_AF_TIPC = "PF_TIPC"
const m_AF_UNIX = "AF_LOCAL"
const m_AF_UNSPEC = "PF_UNSPEC"
const m_AF_VSOCK = "PF_VSOCK"
const m_AF_WANPIPE = "PF_WANPIPE"
const m_AF_X25 = "PF_X25"
const m_AF_XDP = "PF_XDP"
const m_INT_MAX2 = 0x7fffffff
const m_MSG_BATCH = 0x40000
const m_MSG_CMSG_CLOEXEC = 0x40000000
const m_MSG_CONFIRM = 0x0800
const m_MSG_CTRUNC = 0x0008
const m_MSG_DONTROUTE = 0x0004
const m_MSG_DONTWAIT = 0x0040
const m_MSG_EOR = 0x0080
const m_MSG_ERRQUEUE = 0x2000
const m_MSG_FASTOPEN = 0x20000000
const m_MSG_FIN = 0x0200
const m_MSG_MORE = 0x8000
const m_MSG_NOSIGNAL = 0x4000
const m_MSG_OOB = 0x0001
const m_MSG_PEEK = 0x0002
const m_MSG_PROXY = 0x0010
const m_MSG_RST = 0x1000
const m_MSG_SYN = 0x0400
const m_MSG_TRUNC = 0x0020
const m_MSG_WAITALL = 0x0100
const m_MSG_WAITFORONE = 0x10000
const m_MSG_ZEROCOPY = 0x4000000
const m_PF_ALG = 38
const m_PF_APPLETALK = 5
const m_PF_ASH = 18
const m_PF_ATMPVC = 8
const m_PF_ATMSVC = 20
const m_PF_AX25 = 3
const m_PF_BLUETOOTH = 31
const m_PF_BRIDGE = 7
const m_PF_CAIF = 37
const m_PF_CAN = 29
const m_PF_DECnet = 12
const m_PF_ECONET = 19
const m_PF_FILE = "PF_LOCAL"
const m_PF_IB = 27
const m_PF_IEEE802154 = 36
const m_PF_INET = 2
const m_PF_INET6 = 10
const m_PF_IPX = 4
const m_PF_IRDA = 23
const m_PF_ISDN = 34
const m_PF_IUCV = 32
const m_PF_KCM = 41
const m_PF_KEY = 15
const m_PF_LLC = 26
const m_PF_LOCAL = 1
const m_PF_MAX = 45
const m_PF_MPLS = 28
const m_PF_NETBEUI = 13
const m_PF_NETLINK = 16
const m_PF_NETROM = 6
const m_PF_NFC = 39
const m_PF_PACKET = 17
const m_PF_PHONET = 35
const m_PF_PPPOX = 24
const m_PF_QIPCRTR = 42
const m_PF_RDS = 21
const m_PF_ROSE = 11
const m_PF_ROUTE = "PF_NETLINK"
const m_PF_RXRPC = 33
const m_PF_SECURITY = 14
const m_PF_SMC = 43
const m_PF_SNA = 22
const m_PF_TIPC = 30
const m_PF_UNIX = "PF_LOCAL"
const m_PF_UNSPEC = 0
const m_PF_VSOCK = 40
const m_PF_WANPIPE = 25
const m_PF_X25 = 9
const m_PF_XDP = 44
const m_SCM_CREDENTIALS = 0x02
const m_SCM_RIGHTS = 0x01
const m_SCM_TIMESTAMP = "SO_TIMESTAMP"
const m_SCM_TIMESTAMPING = "SO_TIMESTAMPING"
const m_SCM_TIMESTAMPING_OPT_STATS = 54
const m_SCM_TIMESTAMPING_PKTINFO = 58
const m_SCM_TIMESTAMPNS = "SO_TIMESTAMPNS"
const m_SCM_TXTIME = "SO_TXTIME"
const m_SCM_WIFI_STATUS = "SO_WIFI_STATUS"
const m_SHUT_RD = 0
const m_SHUT_RDWR = 2
const m_SHUT_WR = 1
const m_SOCK_CLOEXEC = 02000000
const m_SOCK_DCCP = 6
const m_SOCK_DGRAM = 2
const m_SOCK_NONBLOCK = 04000
const m_SOCK_PACKET = 10
const m_SOCK_RAW = 3
const m_SOCK_RDM = 4
const m_SOCK_SEQPACKET = 5
const m_SOCK_STREAM = 1
const m_SOL_AAL = 265
const m_SOL_ALG = 279
const m_SOL_ATM = 264
const m_SOL_BLUETOOTH = 274
const m_SOL_CAIF = 278
const m_SOL_DCCP = 269
const m_SOL_DECNET = 261
const m_SOL_ICMPV6 = 58
const m_SOL_IP = 0
const m_SOL_IPV6 = 41
const m_SOL_IRDA = 266
const m_SOL_IUCV = 277
const m_SOL_KCM = 281
const m_SOL_LLC = 268
const m_SOL_NETBEUI = 267
const m_SOL_NETLINK = 270
const m_SOL_NFC = 280
const m_SOL_PACKET = 263
const m_SOL_PNPIPE = 275
const m_SOL_PPPOL2TP = 273
const m_SOL_RAW = 255
const m_SOL_RDS = 276
const m_SOL_RXRPC = 272
const m_SOL_SOCKET = 1
const m_SOL_TIPC = 271
const m_SOL_TLS = 282
const m_SOL_X25 = 262
const m_SOL_XDP = 283
const m_SOMAXCONN = 128
const m_SO_ACCEPTCONN = 30
const m_SO_ATTACH_BPF = 50
const m_SO_ATTACH_FILTER = 26
const m_SO_ATTACH_REUSEPORT_CBPF = 51
const m_SO_ATTACH_REUSEPORT_EBPF = 52
const m_SO_BINDTODEVICE = 25
const m_SO_BINDTOIFINDEX = 62
const m_SO_BPF_EXTENSIONS = 48
const m_SO_BROADCAST = 6
const m_SO_BSDCOMPAT = 14
const m_SO_BUSY_POLL = 46
const m_SO_BUSY_POLL_BUDGET = 70
const m_SO_CNX_ADVICE = 53
const m_SO_COOKIE = 57
const m_SO_DEBUG = 1
const m_SO_DETACH_BPF = "SO_DETACH_FILTER"
const m_SO_DETACH_FILTER = 27
const m_SO_DETACH_REUSEPORT_BPF = 68
const m_SO_DOMAIN = 39
const m_SO_DONTROUTE = 5
const m_SO_ERROR = 4
const m_SO_GET_FILTER = "SO_ATTACH_FILTER"
const m_SO_INCOMING_CPU = 49
const m_SO_INCOMING_NAPI_ID = 56
const m_SO_KEEPALIVE = 9
const m_SO_LINGER = 13
const m_SO_LOCK_FILTER = 44
const m_SO_MARK = 36
const m_SO_MAX_PACING_RATE = 47
const m_SO_MEMINFO = 55
const m_SO_NOFCS = 43
const m_SO_NO_CHECK = 11
const m_SO_OOBINLINE = 10
const m_SO_PASSCRED = 16
const m_SO_PASSSEC = 34
const m_SO_PEEK_OFF = 42
const m_SO_PEERCRED = 17
const m_SO_PEERGROUPS = 59
const m_SO_PEERNAME = 28
const m_SO_PEERSEC = 31
const m_SO_PREFER_BUSY_POLL = 69
const m_SO_PRIORITY = 12
const m_SO_PROTOCOL = 38
const m_SO_RCVBUF = 8
const m_SO_RCVBUFFORCE = 33
const m_SO_RCVLOWAT = 18
const m_SO_RCVTIMEO = 20
const m_SO_REUSEADDR = 2
const m_SO_REUSEPORT = 15
const m_SO_RXQ_OVFL = 40
const m_SO_SECURITY_AUTHENTICATION = 22
const m_SO_SECURITY_ENCRYPTION_NETWORK = 24
const m_SO_SECURITY_ENCRYPTION_TRANSPORT = 23
const m_SO_SELECT_ERR_QUEUE = 45
const m_SO_SNDBUF = 7
const m_SO_SNDBUFFORCE = 32
const m_SO_SNDLOWAT = 19
const m_SO_SNDTIMEO = 21
const m_SO_TIMESTAMP = 29
const m_SO_TIMESTAMPING = 37
const m_SO_TIMESTAMPNS = 35
const m_SO_TXTIME = 61
const m_SO_TYPE = 3
const m_SO_WIFI_STATUS = 41
const m_SO_ZEROCOPY = 60
const m_UINT_MAX2 = 0xffffffff

type Tsocklen_t = uint32

type Tsa_family_t = uint16

type Tmsghdr = struct {
	Fmsg_name       uintptr
	Fmsg_namelen    Tsocklen_t
	Fmsg_iov        uintptr
	Fmsg_iovlen     int32
	F__pad1         int32
	Fmsg_control    uintptr
	Fmsg_controllen Tsocklen_t
	F__pad2         int32
	Fmsg_flags      int32
}

type Tcmsghdr = struct {
	Fcmsg_len   Tsocklen_t
	F__pad1     int32
	Fcmsg_level int32
	Fcmsg_type  int32
}

type Tucred = struct {
	Fpid Tpid_t
	Fuid Tuid_t
	Fgid Tgid_t
}

type Tmmsghdr = struct {
	Fmsg_hdr Tmsghdr
	Fmsg_len uint32
}

type Tlinger = struct {
	Fl_onoff  int32
	Fl_linger int32
}

type Tsockaddr = struct {
	Fsa_family Tsa_family_t
	Fsa_data   [14]int8
}

type Tsockaddr_storage = struct {
	Fss_family    Tsa_family_t
	F__ss_padding [118]int8
	F__ss_align   uint64
}

type Tsockaddr_un = struct {
	Fsun_family Tsa_family_t
	Fsun_path   [108]int8
}

// C documentation
//
//	/* Linux and OpenBSD */
func Xgetpeereid(tls *libc.TLS, s int32, euid uintptr, egid uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ret int32
	var _ /* cred at bp+0 */ Tucred
	var _ /* credlen at bp+12 */ Tsocklen_t
	_ = ret
	*(*Tsocklen_t)(unsafe.Pointer(bp + 12)) = uint32(12)
	ret = libc.Xgetsockopt(tls, s, int32(m_SOL_SOCKET), int32(m_SO_PEERCRED), bp, bp+12)
	if ret != 0 {
		return ret
	}
	*(*Tuid_t)(unsafe.Pointer(euid)) = (*(*Tucred)(unsafe.Pointer(bp))).Fuid
	*(*Tgid_t)(unsafe.Pointer(egid)) = (*(*Tucred)(unsafe.Pointer(bp))).Fgid
	return 0
}

const m_INT64_MAX1 = 9223372036854775807

func Xdehumanize_number(tls *libc.TLS, buf uintptr, num uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var rc, sign, v1, v2 int32
	var rmax Tuint64_t
	var _ /* rval at bp+0 */ Tuint64_t
	_, _, _, _, _ = rc, rmax, sign, v1, v2
	sign = +libc.Int32FromInt32(1)
	/* The current expand_number() implementation uses bit shifts, so
	 * we cannot pass negative numbers, preserve the sign and apply it
	 * later. */
	for {
		v1 = int32(*(*int8)(unsafe.Pointer(buf)))
		v2 = libc.BoolInt32(v1 == int32(' ') || libc.Uint32FromInt32(v1)-uint32('\t') < uint32(5))
		goto _3
	_3:
		if !(v2 != 0) {
			break
		}
		buf++
	}
	if int32(*(*int8)(unsafe.Pointer(buf))) == int32('-') {
		sign = -int32(1)
		buf++
	}
	rc = Xexpand_number(tls, buf, bp)
	if rc < 0 {
		return rc
	}
	/* The sign has been stripped, so rval has the absolute value.
	 * Error out, regardless of the sign, if rval is greater than
	 * abs(INT64_MIN) (== INT64_MAX + 1), or if the sign is positive
	 * and the value has overflown by one (INT64_MAX + 1). */
	rmax = uint64(libc.Uint64FromInt64(libc.Int64FromInt64(m_INT64_MAX1)) + libc.Uint64FromUint64(1))
	if *(*Tuint64_t)(unsafe.Pointer(bp)) > rmax || *(*Tuint64_t)(unsafe.Pointer(bp)) == rmax && sign == +libc.Int32FromInt32(1) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ERANGE)
		return -int32(1)
	}
	*(*Tint64_t)(unsafe.Pointer(num)) = libc.Int64FromUint64(*(*Tuint64_t)(unsafe.Pointer(bp)) * libc.Uint64FromInt32(sign))
	return 0
}

const m_HN_AUTOSCALE1 = 32
const m_HN_B1 = 4
const m_HN_DECIMAL1 = 1
const m_HN_DIVISOR_10001 = 8
const m_HN_GETSCALE1 = 16
const m_HN_IEC_PREFIXES1 = 16
const m_HN_NOSPACE1 = 2
const m_INT64_MAX2 = 0x7fffffffffffffff
const m_LC_ALL = 6
const m_LC_ALL_MASK = 0x7fffffff
const m_LC_COLLATE = 3
const m_LC_CTYPE = 0
const m_LC_MESSAGES = 5
const m_LC_MONETARY = 4
const m_LC_NUMERIC = 1
const m_LC_TIME = 2

type Tlconv = struct {
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    int8
	Ffrac_digits        int8
	Fp_cs_precedes      int8
	Fp_sep_by_space     int8
	Fn_cs_precedes      int8
	Fn_sep_by_space     int8
	Fp_sign_posn        int8
	Fn_sign_posn        int8
	Fint_p_cs_precedes  int8
	Fint_p_sep_by_space int8
	Fint_n_cs_precedes  int8
	Fint_n_sep_by_space int8
	Fint_p_sign_posn    int8
	Fint_n_sign_posn    int8
}

/*
 * Copyright (c) 1996  Peter Wemm <peter@FreeBSD.org>.
 * All rights reserved.
 * Copyright (c) 2002 Networks Associates Technology, Inc.
 * All rights reserved.
 *
 * Portions of this software were developed for the FreeBSD Project by
 * ThinkSec AS and NAI Labs, the Security Research Division of Network
 * Associates, Inc.  under DARPA/SPAWAR contract N66001-01-C-8035
 * ("CBOSS"), as part of the DARPA CHATS research program.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, is permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote
 *    products derived from this software without specific prior written
 *    permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 * $FreeBSD: src/lib/libutil/libutil.h,v 1.47 2008/04/23 00:49:12 scf Exp $
 */

var _maxscale = int32(6)

func Xhumanize_number(tls *libc.TLS, buf uintptr, len1 Tsize_t, quotient Tint64_t, suffix uintptr, scale int32, flags int32) (r1 int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var baselen Tsize_t
	var divisor, max Tint64_t
	var divisordeccut, i, r, remainder, s1, s2, sign, v2 int32
	var prefixes, sep uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _ = baselen, divisor, divisordeccut, i, max, prefixes, r, remainder, s1, s2, sep, sign, v2
	/* Since so many callers don't check -1, NUL terminate the buffer */
	if len1 > uint64(0) {
		*(*int8)(unsafe.Pointer(buf)) = int8('\000')
	}
	/* validate args */
	if buf == libc.UintptrFromInt32(0) || suffix == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	if scale < 0 {
		return -int32(1)
	} else {
		if scale > _maxscale && scale & ^(libc.Int32FromInt32(m_HN_AUTOSCALE1)|libc.Int32FromInt32(m_HN_GETSCALE1)) != 0 {
			return -int32(1)
		}
	}
	if flags&int32(m_HN_DIVISOR_10001) != 0 && flags&int32(m_HN_IEC_PREFIXES1) != 0 {
		return -int32(1)
	}
	/* setup parameters */
	remainder = 0
	if flags&int32(m_HN_IEC_PREFIXES1) != 0 {
		baselen = uint64(2)
		/*
		 * Use the prefixes for power of two recommended by
		 * the International Electrotechnical Commission
		 * (IEC) in IEC 80000-3 (i.e. Ki, Mi, Gi...).
		 *
		 * HN_IEC_PREFIXES implies a divisor of 1024 here
		 * (use of HN_DIVISOR_1000 would have triggered
		 * an assertion earlier).
		 */
		divisor = int64(1024)
		divisordeccut = int32(973) /* ceil(.95 * 1024) */
		if flags&int32(m_HN_B1) != 0 {
			prefixes = __ccgo_ts + 163
		} else {
			prefixes = __ccgo_ts + 184
		}
	} else {
		baselen = uint64(1)
		if flags&int32(m_HN_DIVISOR_10001) != 0 {
			divisor = int64(1000)
			divisordeccut = int32(950)
			if flags&int32(m_HN_B1) != 0 {
				prefixes = __ccgo_ts + 205
			} else {
				prefixes = __ccgo_ts + 225
			}
		} else {
			divisor = int64(1024)
			divisordeccut = int32(973) /* ceil(.95 * 1024) */
			if flags&int32(m_HN_B1) != 0 {
				prefixes = __ccgo_ts + 245
			} else {
				prefixes = __ccgo_ts + 265
			}
		}
	}
	if quotient < 0 {
		sign = -int32(1)
		quotient = -quotient
		baselen += uint64(2) /* sign, digit */
	} else {
		sign = int32(1)
		baselen += uint64(1) /* digit */
	}
	if flags&int32(m_HN_NOSPACE1) != 0 {
		sep = __ccgo_ts + 60
	} else {
		sep = __ccgo_ts + 285
		baselen++
	}
	baselen += libc.Xstrlen(tls, suffix)
	/* Check if enough room for `x y' + suffix + `\0' */
	if len1 < baselen+uint64(1) {
		return -int32(1)
	}
	if scale&(libc.Int32FromInt32(m_HN_AUTOSCALE1)|libc.Int32FromInt32(m_HN_GETSCALE1)) != 0 {
		/* See if there is additional columns can be used. */
		max = int64(1)
		i = libc.Int32FromUint64(len1 - baselen)
		for {
			v2 = i
			i--
			if !(v2 > 0) {
				break
			}
			max *= int64(10)
			goto _1
		_1:
		}
		/*
		 * Divide the number until it fits the given column.
		 * If there will be an overflow by the rounding below,
		 * divide once more.
		 */
		i = 0
		for {
			if !((quotient >= max || quotient == max-int64(1) && (remainder >= divisordeccut || int64(remainder) >= divisor/int64(2))) && i < _maxscale) {
				break
			}
			remainder = int32(quotient % divisor)
			quotient /= divisor
			goto _3
		_3:
			;
			i++
		}
		if scale&int32(m_HN_GETSCALE1) != 0 {
			return i
		}
	} else {
		i = 0
		for {
			if !(i < scale && i < _maxscale) {
				break
			}
			remainder = int32(quotient % divisor)
			quotient /= divisor
			goto _4
		_4:
			;
			i++
		}
	}
	/* If a value <= 9.9 after rounding and ... */
	/*
	 * XXX - should we make sure there is enough space for the decimal
	 * place and if not, don't do HN_DECIMAL?
	 */
	if (quotient == int64(9) && remainder < divisordeccut || quotient < int64(9)) && i > 0 && flags&int32(m_HN_DECIMAL1) != 0 {
		s1 = int32(int64(int32(quotient)) + (int64(remainder*int32(10))+divisor/int64(2))/divisor/int64(10))
		s2 = int32((int64(remainder*int32(10)) + divisor/int64(2)) / divisor % int64(10))
		r = libc.X__builtin_snprintf(tls, buf, len1, __ccgo_ts+287, libc.VaList(bp+8, sign*s1, (*Tlconv)(unsafe.Pointer(libc.Xlocaleconv(tls))).Fdecimal_point, s2, sep, prefixes+uintptr(i*int32(3)), suffix))
	} else {
		r = libc.X__builtin_snprintf(tls, buf, len1, __ccgo_ts+300, libc.VaList(bp+8, int64(sign)*(quotient+(int64(remainder)+divisor/int64(2))/divisor), sep, prefixes+uintptr(i*int32(3)), suffix))
	}
	return r
}

const m_HN_AUTOSCALE2 = 0x20
const m_HN_B2 = 0x04
const m_HN_DECIMAL2 = 0x01
const m_HN_DIVISOR_10002 = 0x08
const m_HN_GETSCALE2 = 0x10
const m_HN_IEC_PREFIXES2 = 0x10
const m_HN_NOSPACE2 = 0x02
const m_INET6_ADDRSTRLEN = 46
const m_INET_ADDRSTRLEN = 16
const m_IN_CLASSA_MAX = 128
const m_IN_CLASSA_NET = 0xff000000
const m_IN_CLASSA_NSHIFT = 24
const m_IN_CLASSB_MAX = 65536
const m_IN_CLASSB_NET = 0xffff0000
const m_IN_CLASSB_NSHIFT = 16
const m_IN_CLASSC_NET = 0xffffff00
const m_IN_CLASSC_NSHIFT = 8
const m_IN_LOOPBACKNET = 127
const m_IPPORT_RESERVED = 1024
const m_IPPROTO_AH = 51
const m_IPPROTO_BEETPH = 94
const m_IPPROTO_COMP = 108
const m_IPPROTO_DCCP = 33
const m_IPPROTO_DSTOPTS = 60
const m_IPPROTO_EGP = 8
const m_IPPROTO_ENCAP = 98
const m_IPPROTO_ESP = 50
const m_IPPROTO_ETHERNET = 143
const m_IPPROTO_FRAGMENT = 44
const m_IPPROTO_GRE = 47
const m_IPPROTO_HOPOPTS = 0
const m_IPPROTO_ICMP = 1
const m_IPPROTO_ICMPV6 = 58
const m_IPPROTO_IDP = 22
const m_IPPROTO_IGMP = 2
const m_IPPROTO_IP = 0
const m_IPPROTO_IPIP = 4
const m_IPPROTO_IPV6 = 41
const m_IPPROTO_MAX = 263
const m_IPPROTO_MH = 135
const m_IPPROTO_MPLS = 137
const m_IPPROTO_MPTCP = 262
const m_IPPROTO_MTP = 92
const m_IPPROTO_NONE = 59
const m_IPPROTO_PIM = 103
const m_IPPROTO_PUP = 12
const m_IPPROTO_RAW = 255
const m_IPPROTO_ROUTING = 43
const m_IPPROTO_RSVP = 46
const m_IPPROTO_SCTP = 132
const m_IPPROTO_TCP = 6
const m_IPPROTO_TP = 29
const m_IPPROTO_UDP = 17
const m_IPPROTO_UDPLITE = 136
const m_IPV6_2292DSTOPTS = 4
const m_IPV6_2292HOPLIMIT = 8
const m_IPV6_2292HOPOPTS = 3
const m_IPV6_2292PKTINFO = 2
const m_IPV6_2292PKTOPTIONS = 6
const m_IPV6_2292RTHDR = 5
const m_IPV6_ADDRFORM = 1
const m_IPV6_ADDR_PREFERENCES = 72
const m_IPV6_ADD_MEMBERSHIP = "IPV6_JOIN_GROUP"
const m_IPV6_AUTHHDR = 10
const m_IPV6_AUTOFLOWLABEL = 70
const m_IPV6_CHECKSUM = 7
const m_IPV6_DONTFRAG = 62
const m_IPV6_DROP_MEMBERSHIP = "IPV6_LEAVE_GROUP"
const m_IPV6_DSTOPTS = 59
const m_IPV6_FREEBIND = 78
const m_IPV6_HDRINCL = 36
const m_IPV6_HOPLIMIT = 52
const m_IPV6_HOPOPTS = 54
const m_IPV6_IPSEC_POLICY = 34
const m_IPV6_JOIN_ANYCAST = 27
const m_IPV6_JOIN_GROUP = 20
const m_IPV6_LEAVE_ANYCAST = 28
const m_IPV6_LEAVE_GROUP = 21
const m_IPV6_MINHOPCOUNT = 73
const m_IPV6_MTU = 24
const m_IPV6_MTU_DISCOVER = 23
const m_IPV6_MULTICAST_ALL = 29
const m_IPV6_MULTICAST_HOPS = 18
const m_IPV6_MULTICAST_IF = 17
const m_IPV6_MULTICAST_LOOP = 19
const m_IPV6_NEXTHOP = 9
const m_IPV6_ORIGDSTADDR = 74
const m_IPV6_PATHMTU = 61
const m_IPV6_PKTINFO = 50
const m_IPV6_PMTUDISC_DO = 2
const m_IPV6_PMTUDISC_DONT = 0
const m_IPV6_PMTUDISC_INTERFACE = 4
const m_IPV6_PMTUDISC_OMIT = 5
const m_IPV6_PMTUDISC_PROBE = 3
const m_IPV6_PMTUDISC_WANT = 1
const m_IPV6_PREFER_SRC_CGA = 0x0008
const m_IPV6_PREFER_SRC_COA = 0x0004
const m_IPV6_PREFER_SRC_HOME = 0x0400
const m_IPV6_PREFER_SRC_NONCGA = 0x0800
const m_IPV6_PREFER_SRC_PUBLIC = 0x0002
const m_IPV6_PREFER_SRC_PUBTMP_DEFAULT = 0x0100
const m_IPV6_PREFER_SRC_TMP = 0x0001
const m_IPV6_RECVDSTOPTS = 58
const m_IPV6_RECVERR = 25
const m_IPV6_RECVFRAGSIZE = 77
const m_IPV6_RECVHOPLIMIT = 51
const m_IPV6_RECVHOPOPTS = 53
const m_IPV6_RECVORIGDSTADDR = "IPV6_ORIGDSTADDR"
const m_IPV6_RECVPATHMTU = 60
const m_IPV6_RECVPKTINFO = 49
const m_IPV6_RECVRTHDR = 56
const m_IPV6_RECVTCLASS = 66
const m_IPV6_ROUTER_ALERT = 22
const m_IPV6_ROUTER_ALERT_ISOLATE = 30
const m_IPV6_RTHDR = 57
const m_IPV6_RTHDRDSTOPTS = 55
const m_IPV6_RTHDR_LOOSE = 0
const m_IPV6_RTHDR_STRICT = 1
const m_IPV6_RTHDR_TYPE_0 = 0
const m_IPV6_RXDSTOPTS = "IPV6_DSTOPTS"
const m_IPV6_RXHOPOPTS = "IPV6_HOPOPTS"
const m_IPV6_TCLASS = 67
const m_IPV6_TRANSPARENT = 75
const m_IPV6_UNICAST_HOPS = 16
const m_IPV6_UNICAST_IF = 76
const m_IPV6_V6ONLY = 26
const m_IPV6_XFRM_POLICY = 35
const m_IP_ADD_MEMBERSHIP = 35
const m_IP_ADD_SOURCE_MEMBERSHIP = 39
const m_IP_BIND_ADDRESS_NO_PORT = 24
const m_IP_BLOCK_SOURCE = 38
const m_IP_CHECKSUM = 23
const m_IP_DEFAULT_MULTICAST_LOOP = 1
const m_IP_DEFAULT_MULTICAST_TTL = 1
const m_IP_DROP_MEMBERSHIP = 36
const m_IP_DROP_SOURCE_MEMBERSHIP = 40
const m_IP_FREEBIND = 15
const m_IP_HDRINCL = 3
const m_IP_IPSEC_POLICY = 16
const m_IP_MAX_MEMBERSHIPS = 20
const m_IP_MINTTL = 21
const m_IP_MSFILTER = 41
const m_IP_MTU = 14
const m_IP_MTU_DISCOVER = 10
const m_IP_MULTICAST_ALL = 49
const m_IP_MULTICAST_IF = 32
const m_IP_MULTICAST_LOOP = 34
const m_IP_MULTICAST_TTL = 33
const m_IP_NODEFRAG = 22
const m_IP_OPTIONS = 4
const m_IP_ORIGDSTADDR = 20
const m_IP_PASSSEC = 18
const m_IP_PKTINFO = 8
const m_IP_PKTOPTIONS = 9
const m_IP_PMTUDISC = 10
const m_IP_PMTUDISC_DO = 2
const m_IP_PMTUDISC_DONT = 0
const m_IP_PMTUDISC_INTERFACE = 4
const m_IP_PMTUDISC_OMIT = 5
const m_IP_PMTUDISC_PROBE = 3
const m_IP_PMTUDISC_WANT = 1
const m_IP_RECVERR = 11
const m_IP_RECVERR_RFC4884 = 26
const m_IP_RECVFRAGSIZE = 25
const m_IP_RECVOPTS = 6
const m_IP_RECVORIGDSTADDR = "IP_ORIGDSTADDR"
const m_IP_RECVRETOPTS = "IP_RETOPTS"
const m_IP_RECVTOS = 13
const m_IP_RECVTTL = 12
const m_IP_RETOPTS = 7
const m_IP_ROUTER_ALERT = 5
const m_IP_TOS = 1
const m_IP_TRANSPARENT = 19
const m_IP_TTL = 2
const m_IP_UNBLOCK_SOURCE = 37
const m_IP_UNICAST_IF = 50
const m_IP_XFRM_POLICY = 17
const m_MCAST_BLOCK_SOURCE = 43
const m_MCAST_EXCLUDE = 0
const m_MCAST_INCLUDE = 1
const m_MCAST_JOIN_GROUP = 42
const m_MCAST_JOIN_SOURCE_GROUP = 46
const m_MCAST_LEAVE_GROUP = 45
const m_MCAST_LEAVE_SOURCE_GROUP = 47
const m_MCAST_MSFILTER = 48
const m_MCAST_UNBLOCK_SOURCE = 44
const m___UAPI_DEF_IN6_ADDR = 0
const m___UAPI_DEF_IN6_ADDR_ALT = 0
const m___UAPI_DEF_IN6_PKTINFO = 0
const m___UAPI_DEF_IN_ADDR = 0
const m___UAPI_DEF_IN_CLASS = 0
const m___UAPI_DEF_IN_IPPROTO = 0
const m___UAPI_DEF_IN_PKTINFO = 0
const m___UAPI_DEF_IP6_MTUINFO = 0
const m___UAPI_DEF_IPPROTO_V6 = 0
const m___UAPI_DEF_IPV6_MREQ = 0
const m___UAPI_DEF_IPV6_OPTIONS = 0
const m___UAPI_DEF_IP_MREQ = 0
const m___UAPI_DEF_SOCKADDR_IN = 0
const m___UAPI_DEF_SOCKADDR_IN6 = 0

type Tin_port_t = uint16

type Tin_addr_t = uint32

type Tin_addr = struct {
	Fs_addr Tin_addr_t
}

type Tsockaddr_in = struct {
	Fsin_family Tsa_family_t
	Fsin_port   Tin_port_t
	Fsin_addr   Tin_addr
	Fsin_zero   [8]Tuint8_t
}

type Tin6_addr = struct {
	F__in6_union struct {
		F__s6_addr16 [0][8]Tuint16_t
		F__s6_addr32 [0][4]Tuint32_t
		F__s6_addr   [16]Tuint8_t
	}
}

type Tsockaddr_in6 = struct {
	Fsin6_family   Tsa_family_t
	Fsin6_port     Tin_port_t
	Fsin6_flowinfo Tuint32_t
	Fsin6_addr     Tin6_addr
	Fsin6_scope_id Tuint32_t
}

type Tipv6_mreq = struct {
	Fipv6mr_multiaddr Tin6_addr
	Fipv6mr_interface uint32
}

type Tip_opts = struct {
	Fip_dst  Tin_addr
	Fip_opts [40]int8
}

type Tip_mreq = struct {
	Fimr_multiaddr Tin_addr
	Fimr_interface Tin_addr
}

type Tip_mreqn = struct {
	Fimr_multiaddr Tin_addr
	Fimr_address   Tin_addr
	Fimr_ifindex   int32
}

type Tip_mreq_source = struct {
	Fimr_multiaddr  Tin_addr
	Fimr_interface  Tin_addr
	Fimr_sourceaddr Tin_addr
}

type Tip_msfilter = struct {
	Fimsf_multiaddr Tin_addr
	Fimsf_interface Tin_addr
	Fimsf_fmode     Tuint32_t
	Fimsf_numsrc    Tuint32_t
	Fimsf_slist     [1]Tin_addr
}

type Tgroup_req = struct {
	Fgr_interface Tuint32_t
	Fgr_group     Tsockaddr_storage
}

type Tgroup_source_req = struct {
	Fgsr_interface Tuint32_t
	Fgsr_group     Tsockaddr_storage
	Fgsr_source    Tsockaddr_storage
}

type Tgroup_filter = struct {
	Fgf_interface Tuint32_t
	Fgf_group     Tsockaddr_storage
	Fgf_fmode     Tuint32_t
	Fgf_numsrc    Tuint32_t
	Fgf_slist     [1]Tsockaddr_storage
}

type Tin_pktinfo = struct {
	Fipi_ifindex  int32
	Fipi_spec_dst Tin_addr
	Fipi_addr     Tin_addr
}

type Tin6_pktinfo = struct {
	Fipi6_addr    Tin6_addr
	Fipi6_ifindex uint32
}

type Tip6_mtuinfo = struct {
	Fip6m_addr Tsockaddr_in6
	Fip6m_mtu  Tuint32_t
}

// C documentation
//
//	/*
//	 * static int
//	 * inet_net_pton(af, src, dst, size)
//	 *	convert network number from presentation to network format.
//	 *	accepts hex octets, hex strings, decimal octets, and /CIDR.
//	 *	"size" is in bytes and describes "dst".
//	 * return:
//	 *	number of bits, either imputed classfully or specified with /CIDR,
//	 *	or -1 if some failure occurred (check errno).  ENOENT means it was
//	 *	not a valid network specification.
//	 * author:
//	 *	Paul Vixie (ISC), June 1996
//	 */
func Xinet_net_pton(tls *libc.TLS, af int32, src uintptr, dst uintptr, size Tsize_t) (r int32) {
	switch af {
	case int32(m_PF_INET):
		return _inet_net_pton_ipv4(tls, src, dst, size)
	default:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EAFNOSUPPORT)
		return -int32(1)
	}
	return r
}

// C documentation
//
//	/*
//	 * static int
//	 * inet_net_pton_ipv4(src, dst, size)
//	 *	convert IPv4 network number from presentation to network format.
//	 *	accepts hex octets, hex strings, decimal octets, and /CIDR.
//	 *	"size" is in bytes and describes "dst".
//	 * return:
//	 *	number of bits, either imputed classfully or specified with /CIDR,
//	 *	or -1 if some failure occurred (check errno).  ENOENT means it was
//	 *	not an IPv4 network specification.
//	 * note:
//	 *	network byte order assumed.  this means 192.5.5.240/28 has
//	 *	0x11110000 in its fourth octet.
//	 * author:
//	 *	Paul Vixie (ISC), June 1996
//	 */
func _inet_net_pton_ipv4(tls *libc.TLS, src uintptr, dst uintptr, size Tsize_t) (r int32) {
	var bits, ch, dirty, n, tmp, v10, v17, v2, v5 int32
	var odst, v1, v11, v14, v15, v16, v18, v21, v3, v8, p4, p6 uintptr
	var v13, v20, v7 Tsize_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bits, ch, dirty, n, odst, tmp, v1, v10, v11, v13, v14, v15, v16, v17, v18, v2, v20, v21, v3, v5, v7, v8, p4, p6
	odst = dst
	v1 = src
	src++
	ch = int32(*(*int8)(unsafe.Pointer(v1)))
	if ch == int32('0') && (int32(*(*int8)(unsafe.Pointer(src))) == int32('x') || int32(*(*int8)(unsafe.Pointer(src))) == int32('X')) && libc.BoolInt32(libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(src + 1))) < uint32(128)) != 0 && libc.Xisxdigit(tls, int32(*(*int8)(unsafe.Pointer(src + 1)))) != 0 {
		/* Hexadecimal: Eat nybble string. */
		if size <= uint64(0) {
			goto emsgsize
		}
		*(*uint8)(unsafe.Pointer(dst)) = uint8(0)
		dirty = libc.Int32FromInt32(0)
		src++ /* skip x or X. */
		for {
			v3 = src
			src++
			v2 = int32(*(*int8)(unsafe.Pointer(v3)))
			ch = v2
			if !(v2 != int32('\000') && libc.BoolInt32(libc.Uint32FromInt32(ch) < uint32(128)) != 0 && libc.Xisxdigit(tls, ch) != 0) {
				break
			}
			if libc.BoolInt32(libc.Uint32FromInt32(ch)-uint32('A') < uint32(26)) != 0 {
				ch = libc.Xtolower(tls, ch)
			}
			n = int32(int64(libc.Xstrchr(tls, uintptr(unsafe.Pointer(&_xdigits)), ch)) - t__predefined_ptrdiff_t(uintptr(unsafe.Pointer(&_xdigits))))
			p4 = dst
			*(*uint8)(unsafe.Pointer(p4)) = uint8(int32(*(*uint8)(unsafe.Pointer(p4))) | n)
			v5 = dirty
			dirty++
			if !(v5 != 0) {
				p6 = dst
				*(*uint8)(unsafe.Pointer(p6)) = uint8(int32(*(*uint8)(unsafe.Pointer(p6))) << libc.Int32FromInt32(4))
			} else {
				v7 = size
				size--
				if v7 > uint64(0) {
					dst++
					v8 = dst
					*(*uint8)(unsafe.Pointer(v8)) = uint8(0)
					dirty = libc.Int32FromInt32(0)
				} else {
					goto emsgsize
				}
			}
		}
		if dirty != 0 {
			size--
		}
	} else {
		if libc.BoolInt32(libc.Uint32FromInt32(ch) < uint32(128)) != 0 && libc.BoolInt32(libc.Uint32FromInt32(ch)-uint32('0') < uint32(10)) != 0 {
			/* Decimal: eat dotted digit string. */
			for {
				tmp = 0
				for {
					n = int32(int64(libc.Xstrchr(tls, uintptr(unsafe.Pointer(&_digits)), ch)) - t__predefined_ptrdiff_t(uintptr(unsafe.Pointer(&_digits))))
					tmp *= int32(10)
					tmp += n
					if tmp > int32(255) {
						goto enoent
					}
					goto _12
				_12:
					;
					v11 = src
					src++
					v10 = int32(*(*int8)(unsafe.Pointer(v11)))
					ch = v10
					if !(v10 != int32('\000') && libc.BoolInt32(libc.Uint32FromInt32(ch) < uint32(128)) != 0 && libc.BoolInt32(libc.Uint32FromInt32(ch)-uint32('0') < uint32(10)) != 0) {
						break
					}
				}
				v13 = size
				size--
				if v13 <= uint64(0) {
					goto emsgsize
				}
				v14 = dst
				dst++
				*(*uint8)(unsafe.Pointer(v14)) = libc.Uint8FromInt32(tmp)
				if ch == int32('\000') || ch == int32('/') {
					break
				}
				if ch != int32('.') {
					goto enoent
				}
				v15 = src
				src++
				ch = int32(*(*int8)(unsafe.Pointer(v15)))
				if !(libc.BoolInt32(libc.Uint32FromInt32(ch) < libc.Uint32FromInt32(128)) != 0) || !(libc.BoolInt32(libc.Uint32FromInt32(ch)-libc.Uint32FromUint8('0') < libc.Uint32FromInt32(10)) != 0) {
					goto enoent
				}
				goto _9
			_9:
			}
		} else {
			goto enoent
		}
	}
	bits = -int32(1)
	if ch == int32('/') && libc.BoolInt32(libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(src))) < uint32(128)) != 0 && libc.BoolInt32(libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(src)))-uint32('0') < uint32(10)) != 0 && dst > odst {
		/* CIDR width specifier.  Nothing can follow it. */
		v16 = src
		src++
		ch = int32(*(*int8)(unsafe.Pointer(v16))) /* Skip over the /. */
		bits = 0
		for {
			n = int32(int64(libc.Xstrchr(tls, uintptr(unsafe.Pointer(&_digits)), ch)) - t__predefined_ptrdiff_t(uintptr(unsafe.Pointer(&_digits))))
			bits *= int32(10)
			bits += n
			goto _19
		_19:
			;
			v18 = src
			src++
			v17 = int32(*(*int8)(unsafe.Pointer(v18)))
			ch = v17
			if !(v17 != int32('\000') && libc.BoolInt32(libc.Uint32FromInt32(ch) < uint32(128)) != 0 && libc.BoolInt32(libc.Uint32FromInt32(ch)-uint32('0') < uint32(10)) != 0) {
				break
			}
		}
		if ch != int32('\000') {
			goto enoent
		}
		if bits > int32(32) {
			goto emsgsize
		}
	}
	/* Firey death and destruction unless we prefetched EOS. */
	if ch != int32('\000') {
		goto enoent
	}
	/* If nothing was written to the destination, we found no address. */
	if dst == odst {
		goto enoent
	}
	/* If no CIDR spec was given, infer width from net class. */
	if bits == -int32(1) {
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(odst))) >= int32(240) { /* Class E */
			bits = int32(32)
		} else {
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(odst))) >= int32(224) { /* Class D */
				bits = int32(4)
			} else {
				if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(odst))) >= int32(192) { /* Class C */
					bits = int32(24)
				} else {
					if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(odst))) >= int32(128) { /* Class B */
						bits = int32(16)
					} else { /* Class A */
						bits = int32(8)
					}
				}
			}
		}
		/* If imputed mask is narrower than specified octets, widen. */
		if int64(bits) < (int64(dst)-int64(odst))*int64(8) {
			bits = int32((int64(dst) - int64(odst)) * int64(8))
		}
	}
	/* Extend network to cover the actual mask. */
	for int64(bits) > (int64(dst)-int64(odst))*int64(8) {
		v20 = size
		size--
		if v20 <= uint64(0) {
			goto emsgsize
		}
		v21 = dst
		dst++
		*(*uint8)(unsafe.Pointer(v21)) = uint8('\000')
	}
	return bits
	goto enoent
enoent:
	;
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOENT)
	return -int32(1)
	goto emsgsize
emsgsize:
	;
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EMSGSIZE)
	return -int32(1)
}

var _xdigits = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

var _digits = [11]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

const m_MD5_BLOCK_LENGTH = 64
const m_MD5_DIGEST_LENGTH = 16

type TMD5_CTX = struct {
	Fstate  [4]Tuint32_t
	Fcount  Tuint64_t
	Fbuffer [64]Tuint8_t
}

type TMD5Context = TMD5_CTX

/* Avoid polluting the namespace. Even though this makes this usage
 * implementation-specific, defining it unconditionally should not be
 * a problem, and better than possibly breaking unexpecting code. */

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/* We are almost always included from features.h. */

/* The GNU libc does not support any K&R compilers or the traditional mode
   of ISO C compilers anymore.  Check for some of the combinations not
   supported anymore.  */

/* Some user header file might have defined this before.  */

/* Compilers that lack __has_attribute may object to
       #if defined __has_attribute && __has_attribute (...)
   even though they do not need to evaluate the right-hand side of the &&.
   Similarly for __has_builtin, etc.  */

/* All functions, except those with callbacks or those that
   synchronize memory, are leaf functions.  */

/* GCC can always grok prototypes.  For C++ programs we add throw()
   to help it optimize the function calls.  But this only works with
   gcc 2.8.x and egcs.  For gcc 3.4 and up we even mark C functions
   as non-throwing using a function attribute since programs can use
   the -fexceptions options for C code as well.  */

/* These two macros are not used in glibc anymore.  They are kept here
   only because some other projects expect the macros to be defined.  */

/* For these things, GCC behaves the ANSI way normally,
   and the non-ANSI way under -traditional.  */

/* This is not a typedef so `const __ptr_t' does the right thing.  */

/* C++ needs to know that types and declarations are C, not C++.  */

/* Fortify support.  */

/* Use __builtin_dynamic_object_size at _FORTIFY_SOURCE=3 when available.  */

/* Support for flexible arrays.
   Headers that should use flexible arrays only if they're "real"
   (e.g. only if they won't affect sizeof()) should test
   #if __glibc_c99_flexarr_available.  */

/* __asm__ ("xyz") is used throughout the headers to rename functions
   at the assembly language level.  This is wrapped by the __REDIRECT
   macro, in order to support compilers that can do this some other
   way.  When compilers don't support asm-names at all, we have to do
   preprocessor tricks instead (which don't have exactly the right
   semantics, but it's the best we can do).

   Example:
   int __REDIRECT(setpgrp, (__pid_t pid, __pid_t pgrp), setpgid); */

/*
#elif __SOME_OTHER_COMPILER__

# define __REDIRECT(name, proto, alias) name proto; 	_Pragma("let " #name " = " #alias)
*/

/* GCC and clang have various useful declarations that can be made with
   the '__attribute__' syntax.  All of the ways we use this do fine if
   they are omitted for compilers that don't understand it.  */

/* At some point during the gcc 2.96 development the `malloc' attribute
   for functions was introduced.  We don't want to use it unconditionally
   (although this would be possible) since it generates warnings.  */

/* Tell the compiler which arguments to an allocation function
   indicate the size of the allocation.  */

/* Tell the compiler which argument to an allocation function
   indicates the alignment of the allocation.  */

/* At some point during the gcc 2.96 development the `pure' attribute
   for functions was introduced.  We don't want to use it unconditionally
   (although this would be possible) since it generates warnings.  */

/* This declaration tells the compiler that the value is constant.  */

/* At some point during the gcc 3.1 development the `used' attribute
   for functions was introduced.  We don't want to use it unconditionally
   (although this would be possible) since it generates warnings.  */

/* Since version 3.2, gcc allows marking deprecated functions.  */

/* Since version 4.5, gcc also allows one to specify the message printed
   when a deprecated function is used.  clang claims to be gcc 4.2, but
   may also support this feature.  */

/* At some point during the gcc 2.8 development the `format_arg' attribute
   for functions was introduced.  We don't want to use it unconditionally
   (although this would be possible) since it generates warnings.
   If several `format_arg' attributes are given for the same function, in
   gcc-3.0 and older, all but the last one are ignored.  In newer gccs,
   all designated arguments are considered.  */

/* At some point during the gcc 2.97 development the `strfmon' format
   attribute for functions was introduced.  We don't want to use it
   unconditionally (although this would be possible) since it
   generates warnings.  */

/* The nonnull function attribute marks pointer parameters that
   must not be NULL.  This has the name __nonnull in glibc,
   and __attribute_nonnull__ in files shared with Gnulib to avoid
   collision with a different __nonnull in DragonFlyBSD 5.9.  */

/* The returns_nonnull function attribute marks the return type of the function
   as always being non-null.  */

/* If fortification mode, we warn about unused results of certain
   function calls which can lead to problems.  */

/* Forces a function to be always inlined.  */

/* Associate error messages with the source location of the call site rather
   than with the source location inside the function.  */

/* GCC 4.3 and above with -std=c99 or -std=gnu99 implements ISO C99
   inline semantics, unless -fgnu89-inline is used.  Using __GNUC_STDC_INLINE__
   or __GNUC_GNU_INLINE is not a good enough check for gcc because gcc versions
   older than 4.3 may define these macros and still not guarantee GNU inlining
   semantics.

   clang++ identifies itself as gcc-4.2, but has support for GNU inlining
   semantics, that can be checked for by using the __GNUC_STDC_INLINE_ and
   __GNUC_GNU_INLINE__ macro definitions.  */

/* GCC 4.3 and above allow passing all anonymous arguments of an
   __extern_always_inline function to some other vararg function.  */

/* It is possible to compile containing GCC extensions even if GCC is
   run in pedantic mode if the uses are carefully marked using the
   `__extension__' keyword.  But this is not generally available before
   version 2.8.  */

/* __restrict is known in EGCS 1.2 and above, and in clang.
   It works also in C++ mode (outside of arrays), but only when spelled
   as '__restrict', not 'restrict'.  */

/* ISO C99 also allows to declare arrays as non-overlapping.  The syntax is
     array_name[restrict]
   GCC 3.1 and clang support this.
   This syntax is not usable in C++ mode.  */

/* Undefine (also defined in libc-symbols.h).  */

/* Gnulib avoids including these, as they don't work on non-glibc or
   older glibc platforms.  */
/* Copyright (C) 1999-2023 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/* Properties of long double type.  ldbl-128 version.
   Copyright (C) 2016-2023 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License  published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/* long double is distinct from double, so there is nothing to
   define here.  */

/* __glibc_macro_warning (MESSAGE) issues warning MESSAGE.  This is
   intended for use in preprocessor macros.

   Note: MESSAGE must be a _single_ string; concatenation of string
   literals is not supported.  */

/* Generic selection (ISO C11) is a C-only feature, available in GCC
   since version 4.9.  Previous versions do not provide generic
   selection, even though they might set __STDC_VERSION__ to 201112L,
   when in -std=c11 mode.  Thus, we must check for !defined __GNUC__
   when testing __STDC_VERSION__ for generic selection support.
   On the other hand, Clang also defines __GNUC__, so a clang-specific
   check is required to enable the use of generic selection.  */

/* Specify that a function such as setjmp or vfork may return
   twice.  */

/*
 * On non-glibc based systems, we cannot unconditionally use the
 * __GLIBC_PREREQ macro as it gets expanded before evaluation.
 */

/*
 * Some kFreeBSD headers expect those macros to be set for sanity checks.
 */

/* Define the ABI for the current system. */
//#define LIBBSD_SYS_TIME_BITS 0
//#define LIBBSD_SYS_HAS_TIME64 0

/* Linux headers define a struct with a member names __unused.
 * Debian bugs: #522773 (linux), #522774 (libc).
 * Disable for now. */

/*
 * Return the number of elements in a statically-allocated array,
 * __x.
 */

/*
 * We define this here since <stddef.h>, <sys/queue.h>, and <sys/types.h>
 * require it.
 */

/*
 * Given the pointer x to the member m of the struct s, return
 * a pointer to the containing structure.  When using GCC, we first
 * assign pointer x to a local variable, to check that its type is
 * compatible with member m.
 */

func Xlibbsd_MD5Init(tls *libc.TLS, context uintptr) {
	libmd.XMD5Init(tls, context)
}

var _libbsd_emit_link_warning_MD5Init = [77]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'I', 'n', 'i', 't', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5Update(tls *libc.TLS, context uintptr, data uintptr, len1 Tsize_t) {
	libmd.XMD5Update(tls, context, data, len1)
}

var _libbsd_emit_link_warning_MD5Update = [79]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'U', 'p', 'd', 'a', 't', 'e', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5Pad(tls *libc.TLS, context uintptr) {
	libmd.XMD5Pad(tls, context)
}

var _libbsd_emit_link_warning_MD5Pad = [76]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'P', 'a', 'd', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5Final(tls *libc.TLS, digest uintptr, context uintptr) {
	libmd.XMD5Final(tls, digest, context)
}

var _libbsd_emit_link_warning_MD5Final = [78]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'F', 'i', 'n', 'a', 'l', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5Transform(tls *libc.TLS, state uintptr, block uintptr) {
	libmd.XMD5Transform(tls, state, block)
}

var _libbsd_emit_link_warning_MD5Transform = [82]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'T', 'r', 'a', 'n', 's', 'f', 'o', 'r', 'm', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5End(tls *libc.TLS, context uintptr, buf uintptr) (r uintptr) {
	return libmd.XMD5End(tls, context, buf)
}

var _libbsd_emit_link_warning_MD5End = [76]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'E', 'n', 'd', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return libmd.XMD5File(tls, filename, buf)
}

var _libbsd_emit_link_warning_MD5File = [77]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'F', 'i', 'l', 'e', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, offset Toff_t, length Toff_t) (r uintptr) {
	return libmd.XMD5FileChunk(tls, filename, buf, offset, length)
}

var _libbsd_emit_link_warning_MD5FileChunk = [82]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'F', 'i', 'l', 'e', 'C', 'h', 'u', 'n', 'k', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

func Xlibbsd_MD5Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	return libmd.XMD5Data(tls, data, len1, buf)
}

var _libbsd_emit_link_warning_MD5Data = [77]int8{'T', 'h', 'e', ' ', 'M', 'D', '5', 'D', 'a', 't', 'a', '(', ')', ' ', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'l', 'i', 'b', 'b', 's', 'd', ' ', 'i', 's', ' ', 'a', ' ', 'd', 'e', 'p', 'r', 'e', 'c', 'a', 't', 'e', 'd', ' ', 'w', 'r', 'a', 'p', 'p', 'e', 'r', ',', ' ', 'u', 's', 'e', ' ', 'l', 'i', 'b', 'm', 'd', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', '.'}

const m_AT_BASE = 7
const m_AT_BASE_PLATFORM = 24
const m_AT_CLKTCK = 17
const m_AT_DCACHEBSIZE = 19
const m_AT_EGID = 14
const m_AT_ENTRY = 9
const m_AT_EUID = 12
const m_AT_EXECFD = 2
const m_AT_EXECFN = 31
const m_AT_FLAGS = 8
const m_AT_FPUCW = 18
const m_AT_GID = 13
const m_AT_HWCAP = 16
const m_AT_HWCAP2 = 26
const m_AT_ICACHEBSIZE = 20
const m_AT_IGNORE = 1
const m_AT_IGNOREPPC = 22
const m_AT_L1D_CACHEGEOMETRY = 43
const m_AT_L1D_CACHESHAPE = 35
const m_AT_L1D_CACHESIZE = 42
const m_AT_L1I_CACHEGEOMETRY = 41
const m_AT_L1I_CACHESHAPE = 34
const m_AT_L1I_CACHESIZE = 40
const m_AT_L2_CACHEGEOMETRY = 45
const m_AT_L2_CACHESHAPE = 36
const m_AT_L2_CACHESIZE = 44
const m_AT_L3_CACHEGEOMETRY = 47
const m_AT_L3_CACHESHAPE = 37
const m_AT_L3_CACHESIZE = 46
const m_AT_MINSIGSTKSZ = 51
const m_AT_NOTELF = 10
const m_AT_NULL = 0
const m_AT_PAGESZ = 6
const m_AT_PHDR = 3
const m_AT_PHENT = 4
const m_AT_PHNUM = 5
const m_AT_PLATFORM = 15
const m_AT_RANDOM = 25
const m_AT_SECURE = 23
const m_AT_SYSINFO = 32
const m_AT_SYSINFO_EHDR = 33
const m_AT_UCACHEBSIZE = 21
const m_AT_UID = 11
const m_AUX_FUNC = 2
const m_AUX_OBJECT = 1
const m_BIND_WEAK = 2
const m_DF_1_CONFALT = 0x00002000
const m_DF_1_DIRECT = 0x00000100
const m_DF_1_DISPRELDNE = 0x00008000
const m_DF_1_DISPRELPND = 0x00010000
const m_DF_1_EDITED = 0x00200000
const m_DF_1_ENDFILTEE = 0x00004000
const m_DF_1_GLOBAL = 0x00000002
const m_DF_1_GLOBAUDIT = 0x01000000
const m_DF_1_GROUP = 0x00000004
const m_DF_1_IGNMULDEF = 0x00040000
const m_DF_1_INITFIRST = 0x00000020
const m_DF_1_INTERPOSE = 0x00000400
const m_DF_1_LOADFLTR = 0x00000010
const m_DF_1_NODEFLIB = 0x00000800
const m_DF_1_NODELETE = 0x00000008
const m_DF_1_NODIRECT = 0x00020000
const m_DF_1_NODUMP = 0x00001000
const m_DF_1_NOHDR = 0x00100000
const m_DF_1_NOKSYMS = 0x00080000
const m_DF_1_NOOPEN = 0x00000040
const m_DF_1_NORELOC = 0x00400000
const m_DF_1_NOW = 0x00000001
const m_DF_1_ORIGIN = 0x00000080
const m_DF_1_PIE = 0x08000000
const m_DF_1_SINGLETON = 0x02000000
const m_DF_1_STUB = 0x04000000
const m_DF_1_SYMINTPOSE = 0x00800000
const m_DF_1_TRANS = 0x00000200
const m_DF_BIND_NOW = 0x00000008
const m_DF_ORIGIN = 0x00000001
const m_DF_P1_GROUPPERM = 0x00000002
const m_DF_P1_LAZYLOAD = 0x00000001
const m_DF_STATIC_TLS = 0x00000010
const m_DF_SYMBOLIC = 0x00000002
const m_DF_TEXTREL = 0x00000004
const m_DTF_1_CONFEXP = 0x00000002
const m_DTF_1_PARINIT = 0x00000001
const m_DT_ADDRNUM = 11
const m_DT_ADDRRNGHI = 0x6ffffeff
const m_DT_ADDRRNGLO = 0x6ffffe00
const m_DT_ALPHA_NUM = 1
const m_DT_AUDIT = 0x6ffffefc
const m_DT_AUXILIARY = 0x7ffffffd
const m_DT_BIND_NOW = 24
const m_DT_CHECKSUM = 0x6ffffdf8
const m_DT_CONFIG = 0x6ffffefa
const m_DT_DEBUG = 21
const m_DT_DEPAUDIT = 0x6ffffefb
const m_DT_ENCODING = 32
const m_DT_EXTRANUM = 3
const m_DT_FEATURE_1 = 0x6ffffdfc
const m_DT_FILTER = 0x7fffffff
const m_DT_FINI = 13
const m_DT_FINI_ARRAY = 26
const m_DT_FINI_ARRAYSZ = 28
const m_DT_FLAGS = 30
const m_DT_FLAGS_1 = 0x6ffffffb
const m_DT_GNU_CONFLICT = 0x6ffffef8
const m_DT_GNU_CONFLICTSZ = 0x6ffffdf6
const m_DT_GNU_HASH = 0x6ffffef5
const m_DT_GNU_LIBLIST = 0x6ffffef9
const m_DT_GNU_LIBLISTSZ = 0x6ffffdf7
const m_DT_GNU_PRELINKED = 0x6ffffdf5
const m_DT_HASH = 4
const m_DT_HIOS = 0x6ffff000
const m_DT_HIPROC = 0x7fffffff
const m_DT_IA_64_NUM = 1
const m_DT_INIT = 12
const m_DT_INIT_ARRAY = 25
const m_DT_INIT_ARRAYSZ = 27
const m_DT_JMPREL = 23
const m_DT_LOOS = 0x6000000d
const m_DT_LOPROC = 0x70000000
const m_DT_MIPS_AUX_DYNAMIC = 0x70000031
const m_DT_MIPS_BASE_ADDRESS = 0x70000006
const m_DT_MIPS_COMPACT_SIZE = 0x7000002f
const m_DT_MIPS_CONFLICT = 0x70000008
const m_DT_MIPS_CONFLICTNO = 0x7000000b
const m_DT_MIPS_CXX_FLAGS = 0x70000022
const m_DT_MIPS_DELTA_CLASS = 0x70000017
const m_DT_MIPS_DELTA_CLASSSYM = 0x70000020
const m_DT_MIPS_DELTA_CLASSSYM_NO = 0x70000021
const m_DT_MIPS_DELTA_CLASS_NO = 0x70000018
const m_DT_MIPS_DELTA_INSTANCE = 0x70000019
const m_DT_MIPS_DELTA_INSTANCE_NO = 0x7000001a
const m_DT_MIPS_DELTA_RELOC = 0x7000001b
const m_DT_MIPS_DELTA_RELOC_NO = 0x7000001c
const m_DT_MIPS_DELTA_SYM = 0x7000001d
const m_DT_MIPS_DELTA_SYM_NO = 0x7000001e
const m_DT_MIPS_DYNSTR_ALIGN = 0x7000002b
const m_DT_MIPS_FLAGS = 0x70000005
const m_DT_MIPS_GOTSYM = 0x70000013
const m_DT_MIPS_GP_VALUE = 0x70000030
const m_DT_MIPS_HIDDEN_GOTIDX = 0x70000027
const m_DT_MIPS_HIPAGENO = 0x70000014
const m_DT_MIPS_ICHECKSUM = 0x70000003
const m_DT_MIPS_INTERFACE = 0x7000002a
const m_DT_MIPS_INTERFACE_SIZE = 0x7000002c
const m_DT_MIPS_IVERSION = 0x70000004
const m_DT_MIPS_LIBLIST = 0x70000009
const m_DT_MIPS_LIBLISTNO = 0x70000010
const m_DT_MIPS_LOCALPAGE_GOTIDX = 0x70000025
const m_DT_MIPS_LOCAL_GOTIDX = 0x70000026
const m_DT_MIPS_LOCAL_GOTNO = 0x7000000a
const m_DT_MIPS_MSYM = 0x70000007
const m_DT_MIPS_NUM = 0x36
const m_DT_MIPS_OPTIONS = 0x70000029
const m_DT_MIPS_PERF_SUFFIX = 0x7000002e
const m_DT_MIPS_PIXIE_INIT = 0x70000023
const m_DT_MIPS_PLTGOT = 0x70000032
const m_DT_MIPS_PROTECTED_GOTIDX = 0x70000028
const m_DT_MIPS_RLD_MAP = 0x70000016
const m_DT_MIPS_RLD_MAP_REL = 0x70000035
const m_DT_MIPS_RLD_TEXT_RESOLVE_ADDR = 0x7000002d
const m_DT_MIPS_RLD_VERSION = 0x70000001
const m_DT_MIPS_RWPLT = 0x70000034
const m_DT_MIPS_SYMBOL_LIB = 0x70000024
const m_DT_MIPS_SYMTABNO = 0x70000011
const m_DT_MIPS_TIME_STAMP = 0x70000002
const m_DT_MIPS_UNREFEXTNO = 0x70000012
const m_DT_MOVEENT = 0x6ffffdfa
const m_DT_MOVESZ = 0x6ffffdfb
const m_DT_MOVETAB = 0x6ffffefe
const m_DT_NEEDED = 1
const m_DT_NIOS2_GP = 0x70000002
const m_DT_NULL = 0
const m_DT_NUM = 38
const m_DT_PLTGOT = 3
const m_DT_PLTPAD = 0x6ffffefd
const m_DT_PLTPADSZ = 0x6ffffdf9
const m_DT_PLTREL = 20
const m_DT_PLTRELSZ = 2
const m_DT_POSFLAG_1 = 0x6ffffdfd
const m_DT_PPC64_NUM = 4
const m_DT_PPC_NUM = 2
const m_DT_PREINIT_ARRAY = 32
const m_DT_PREINIT_ARRAYSZ = 33
const m_DT_PROCNUM = "DT_MIPS_NUM"
const m_DT_REL = 17
const m_DT_RELA = 7
const m_DT_RELACOUNT = 0x6ffffff9
const m_DT_RELAENT = 9
const m_DT_RELASZ = 8
const m_DT_RELCOUNT = 0x6ffffffa
const m_DT_RELENT = 19
const m_DT_RELR = 36
const m_DT_RELRENT = 37
const m_DT_RELRSZ = 35
const m_DT_RELSZ = 18
const m_DT_RPATH = 15
const m_DT_RUNPATH = 29
const m_DT_SONAME = 14
const m_DT_SPARC_NUM = 2
const m_DT_SPARC_REGISTER = 0x70000001
const m_DT_STRSZ = 10
const m_DT_STRTAB = 5
const m_DT_SYMBOLIC = 16
const m_DT_SYMENT = 11
const m_DT_SYMINENT = 0x6ffffdff
const m_DT_SYMINFO = 0x6ffffeff
const m_DT_SYMINSZ = 0x6ffffdfe
const m_DT_SYMTAB = 6
const m_DT_SYMTAB_SHNDX = 34
const m_DT_TEXTREL = 22
const m_DT_TLSDESC_GOT = 0x6ffffef7
const m_DT_TLSDESC_PLT = 0x6ffffef6
const m_DT_VALNUM = 12
const m_DT_VALRNGHI = 0x6ffffdff
const m_DT_VALRNGLO = 0x6ffffd00
const m_DT_VERDEF = 0x6ffffffc
const m_DT_VERDEFNUM = 0x6ffffffd
const m_DT_VERNEED = 0x6ffffffe
const m_DT_VERNEEDNUM = 0x6fffffff
const m_DT_VERSIONTAGNUM = 16
const m_DT_VERSYM = 0x6ffffff0
const m_EFA_PARISC_1_0 = 0x020b
const m_EFA_PARISC_1_1 = 0x0210
const m_EFA_PARISC_2_0 = 0x0214
const m_EF_ALPHA_32BIT = 1
const m_EF_ALPHA_CANRELAX = 2
const m_EF_ARM_ABI_FLOAT_HARD = 0x400
const m_EF_ARM_ABI_FLOAT_SOFT = 0x200
const m_EF_ARM_ALIGN8 = 0x40
const m_EF_ARM_APCS_26 = 0x08
const m_EF_ARM_APCS_FLOAT = 0x10
const m_EF_ARM_BE8 = 0x00800000
const m_EF_ARM_DYNSYMSUSESEGIDX = 0x08
const m_EF_ARM_EABIMASK = 0xFF000000
const m_EF_ARM_EABI_UNKNOWN = 0x00000000
const m_EF_ARM_EABI_VER1 = 0x01000000
const m_EF_ARM_EABI_VER2 = 0x02000000
const m_EF_ARM_EABI_VER3 = 0x03000000
const m_EF_ARM_EABI_VER4 = 0x04000000
const m_EF_ARM_EABI_VER5 = 0x05000000
const m_EF_ARM_HASENTRY = 0x02
const m_EF_ARM_INTERWORK = 0x04
const m_EF_ARM_LE8 = 0x00400000
const m_EF_ARM_MAPSYMSFIRST = 0x10
const m_EF_ARM_MAVERICK_FLOAT = 0x800
const m_EF_ARM_NEW_ABI = 0x80
const m_EF_ARM_OLD_ABI = 0x100
const m_EF_ARM_PIC = 0x20
const m_EF_ARM_RELEXEC = 0x01
const m_EF_ARM_SOFT_FLOAT = 0x200
const m_EF_ARM_SYMSARESORTED = 0x04
const m_EF_ARM_VFP_FLOAT = 0x400
const m_EF_CPU32 = 0x00810000
const m_EF_IA_64_ABI64 = 0x00000010
const m_EF_IA_64_ARCH = 0xff000000
const m_EF_IA_64_MASKOS = 0x0000000f
const m_EF_LARCH_ABI_DOUBLE_FLOAT = 0x03
const m_EF_LARCH_ABI_MODIFIER_MASK = 0x07
const m_EF_LARCH_ABI_SINGLE_FLOAT = 0x02
const m_EF_LARCH_ABI_SOFT_FLOAT = 0x01
const m_EF_LARCH_OBJABI_V1 = 0x40
const m_EF_MIPS_64BIT_WHIRL = 16
const m_EF_MIPS_ABI2 = 32
const m_EF_MIPS_ABI_ON32 = 64
const m_EF_MIPS_ARCH = 0xf0000000
const m_EF_MIPS_ARCH_1 = 0x00000000
const m_EF_MIPS_ARCH_2 = 0x10000000
const m_EF_MIPS_ARCH_3 = 0x20000000
const m_EF_MIPS_ARCH_32 = 0x50000000
const m_EF_MIPS_ARCH_32R2 = 0x70000000
const m_EF_MIPS_ARCH_4 = 0x30000000
const m_EF_MIPS_ARCH_5 = 0x40000000
const m_EF_MIPS_ARCH_64 = 0x60000000
const m_EF_MIPS_ARCH_64R2 = 0x80000000
const m_EF_MIPS_CPIC = 4
const m_EF_MIPS_FP64 = 512
const m_EF_MIPS_NAN2008 = 1024
const m_EF_MIPS_NOREORDER = 1
const m_EF_MIPS_PIC = 2
const m_EF_MIPS_XGOT = 8
const m_EF_PARISC_ARCH = 0x0000ffff
const m_EF_PARISC_EXT = 0x00020000
const m_EF_PARISC_LAZYSWAP = 0x00400000
const m_EF_PARISC_LSB = 0x00040000
const m_EF_PARISC_NO_KABP = 0x00100000
const m_EF_PARISC_TRAPNIL = 0x00010000
const m_EF_PARISC_WIDE = 0x00080000
const m_EF_PPC64_ABI = 3
const m_EF_PPC_EMB = 0x80000000
const m_EF_PPC_RELOCATABLE = 0x00010000
const m_EF_PPC_RELOCATABLE_LIB = 0x00008000
const m_EF_SH1 = 0x1
const m_EF_SH2 = 0x2
const m_EF_SH2A = 0xd
const m_EF_SH2A_NOFPU = 0x13
const m_EF_SH2A_SH3E = 0x18
const m_EF_SH2A_SH3_NOFPU = 0x16
const m_EF_SH2A_SH4 = 0x17
const m_EF_SH2A_SH4_NOFPU = 0x15
const m_EF_SH2E = 0xb
const m_EF_SH3 = 0x3
const m_EF_SH3E = 0x8
const m_EF_SH3_DSP = 0x5
const m_EF_SH3_NOMMU = 0x14
const m_EF_SH4 = 0x9
const m_EF_SH4A = 0xc
const m_EF_SH4AL_DSP = 0x6
const m_EF_SH4A_NOFPU = 0x11
const m_EF_SH4_NOFPU = 0x10
const m_EF_SH4_NOMMU_NOFPU = 0x12
const m_EF_SH_DSP = 0x4
const m_EF_SH_MACH_MASK = 0x1f
const m_EF_SH_UNKNOWN = 0x0
const m_EF_SPARCV9_MM = 3
const m_EF_SPARCV9_PSO = 1
const m_EF_SPARCV9_RMO = 2
const m_EF_SPARCV9_TSO = 0
const m_EF_SPARC_32PLUS = 0x000100
const m_EF_SPARC_EXT_MASK = 0xFFFF00
const m_EF_SPARC_HAL_R1 = 0x000400
const m_EF_SPARC_LEDATA = 0x800000
const m_EF_SPARC_SUN_US1 = 0x000200
const m_EF_SPARC_SUN_US3 = 0x000800
const m_EI_ABIVERSION = 8
const m_EI_CLASS = 4
const m_EI_DATA = 5
const m_EI_MAG0 = 0
const m_EI_MAG1 = 1
const m_EI_MAG2 = 2
const m_EI_MAG3 = 3
const m_EI_NIDENT = 16
const m_EI_OSABI = 7
const m_EI_PAD = 9
const m_EI_VERSION = 6
const m_ELFCLASS32 = 1
const m_ELFCLASS64 = 2
const m_ELFCLASSNONE = 0
const m_ELFCLASSNUM = 3
const m_ELFCOMPRESS_HIOS = 0x6fffffff
const m_ELFCOMPRESS_HIPROC = 0x7fffffff
const m_ELFCOMPRESS_LOOS = 0x60000000
const m_ELFCOMPRESS_LOPROC = 0x70000000
const m_ELFCOMPRESS_ZLIB = 1
const m_ELFCOMPRESS_ZSTD = 2
const m_ELFDATA2LSB = 1
const m_ELFDATA2MSB = 2
const m_ELFDATANONE = 0
const m_ELFDATANUM = 3
const m_ELFMAG = "\\177ELF"
const m_ELFMAG0 = 127
const m_ELFMAG1 = 69
const m_ELFMAG2 = 76
const m_ELFMAG3 = 70
const m_ELFOSABI_AIX = 7
const m_ELFOSABI_ARM = 97
const m_ELFOSABI_FREEBSD = 9
const m_ELFOSABI_GNU = 3
const m_ELFOSABI_HPUX = 1
const m_ELFOSABI_IRIX = 8
const m_ELFOSABI_LINUX = 3
const m_ELFOSABI_MODESTO = 11
const m_ELFOSABI_NETBSD = 2
const m_ELFOSABI_NONE = 0
const m_ELFOSABI_OPENBSD = 12
const m_ELFOSABI_SOLARIS = 6
const m_ELFOSABI_STANDALONE = 255
const m_ELFOSABI_SYSV = 0
const m_ELFOSABI_TRU64 = 10
const m_ELF_NOTE_ABI = "NT_GNU_ABI_TAG"
const m_ELF_NOTE_GNU = "GNU"
const m_ELF_NOTE_OS_FREEBSD = 3
const m_ELF_NOTE_OS_GNU = 1
const m_ELF_NOTE_OS_LINUX = 0
const m_ELF_NOTE_OS_SOLARIS2 = 2
const m_ELF_NOTE_PAGESIZE_HINT = 1
const m_ELF_NOTE_SOLARIS = "SUNW Solaris"
const m_ELF_ST_BIND = "ELF64_ST_BIND"
const m_ELF_ST_TYPE = "ELF64_ST_TYPE"
const m_ELF_TARG_CLASS = "ELFCLASS64"
const m_ELF_TARG_DATA = "ELFDATA2LSB"
const m_ELF_TARG_MACH = "EM_LOONGARCH"
const m_ELF_TARG_VER = "EV_CURRENT"
const m_EM_386 = 3
const m_EM_56800EX = 200
const m_EM_68HC05 = 72
const m_EM_68HC08 = 71
const m_EM_68HC11 = 70
const m_EM_68HC12 = 53
const m_EM_68HC16 = 69
const m_EM_68K = 4
const m_EM_78KOR = 199
const m_EM_8051 = 165
const m_EM_860 = 7
const m_EM_88K = 5
const m_EM_960 = 19
const m_EM_AARCH64 = 183
const m_EM_ALPHA = 0x9026
const m_EM_ALTERA_NIOS2 = 113
const m_EM_AMDGPU = 224
const m_EM_ARC = 45
const m_EM_ARCA = 109
const m_EM_ARC_A5 = 93
const m_EM_ARC_COMPACT = 93
const m_EM_ARC_COMPACT2 = 195
const m_EM_ARM = 40
const m_EM_AVR = 83
const m_EM_AVR32 = 185
const m_EM_BA1 = 201
const m_EM_BA2 = 202
const m_EM_BLACKFIN = 106
const m_EM_BPF = 247
const m_EM_C166 = 116
const m_EM_CDP = 215
const m_EM_CE = 119
const m_EM_CLOUDSHIELD = 192
const m_EM_COGE = 216
const m_EM_COLDFIRE = 52
const m_EM_COOL = 217
const m_EM_COREA_1ST = 193
const m_EM_COREA_2ND = 194
const m_EM_CR = 103
const m_EM_CR16 = 177
const m_EM_CRAYNV2 = 172
const m_EM_CRIS = 76
const m_EM_CRX = 114
const m_EM_CSKY = 252
const m_EM_CSR_KALIMBA = 219
const m_EM_CUDA = 190
const m_EM_CYPRESS_M8C = 161
const m_EM_D10V = 85
const m_EM_D30V = 86
const m_EM_DSP24 = 136
const m_EM_DSPIC30F = 118
const m_EM_DXP = 112
const m_EM_ECOG16 = 176
const m_EM_ECOG1X = 168
const m_EM_ECOG2 = 134
const m_EM_EMX16 = 212
const m_EM_EMX8 = 213
const m_EM_ETPU = 178
const m_EM_EXCESS = 111
const m_EM_F2MC16 = 104
const m_EM_FAKE_ALPHA = 41
const m_EM_FIREPATH = 78
const m_EM_FR20 = 37
const m_EM_FR30 = 84
const m_EM_FT32 = 222
const m_EM_FX66 = 66
const m_EM_H8S = 48
const m_EM_H8_300 = 46
const m_EM_H8_300H = 47
const m_EM_H8_500 = 49
const m_EM_HUANY = 81
const m_EM_IA_64 = 50
const m_EM_IP2K = 101
const m_EM_JAVELIN = 77
const m_EM_K10M = 181
const m_EM_KM32 = 210
const m_EM_KMX32 = 211
const m_EM_KVARC = 214
const m_EM_L10M = 180
const m_EM_LATTICEMICO32 = 138
const m_EM_LOONGARCH = 258
const m_EM_M16C = 117
const m_EM_M32 = 1
const m_EM_M32C = 120
const m_EM_M32R = 88
const m_EM_MANIK = 171
const m_EM_MAX = 102
const m_EM_MAXQ30 = 169
const m_EM_MCHP_PIC = 204
const m_EM_MCST_ELBRUS = 175
const m_EM_ME16 = 59
const m_EM_METAG = 174
const m_EM_MICROBLAZE = 189
const m_EM_MIPS = 8
const m_EM_MIPS_RS3_LE = 10
const m_EM_MIPS_X = 51
const m_EM_MMA = 54
const m_EM_MMDSP_PLUS = 160
const m_EM_MMIX = 80
const m_EM_MN10200 = 90
const m_EM_MN10300 = 89
const m_EM_MOXIE = 223
const m_EM_MSP430 = 105
const m_EM_NCPU = 56
const m_EM_NDR1 = 57
const m_EM_NDS32 = 167
const m_EM_NONE = 0
const m_EM_NORC = 218
const m_EM_NS32K = 97
const m_EM_NUM = 259
const m_EM_OPEN8 = 196
const m_EM_OPENRISC = 92
const m_EM_OR1K = 92
const m_EM_PARISC = 15
const m_EM_PCP = 55
const m_EM_PDSP = 63
const m_EM_PJ = 91
const m_EM_PPC = 20
const m_EM_PPC64 = 21
const m_EM_PRISM = 82
const m_EM_QDSP6 = 164
const m_EM_R32C = 162
const m_EM_RCE = 39
const m_EM_RH32 = 38
const m_EM_RISCV = 243
const m_EM_RL78 = 197
const m_EM_RS08 = 132
const m_EM_RX = 173
const m_EM_S370 = 9
const m_EM_S390 = 22
const m_EM_SCORE7 = 135
const m_EM_SEP = 108
const m_EM_SE_C17 = 139
const m_EM_SE_C33 = 107
const m_EM_SH = 42
const m_EM_SHARC = 133
const m_EM_SLE9X = 179
const m_EM_SNP1K = 99
const m_EM_SPARC = 2
const m_EM_SPARC32PLUS = 18
const m_EM_SPARCV9 = 43
const m_EM_ST100 = 60
const m_EM_ST19 = 74
const m_EM_ST200 = 100
const m_EM_ST7 = 68
const m_EM_ST9PLUS = 67
const m_EM_STARCORE = 58
const m_EM_STM8 = 186
const m_EM_STXP7X = 166
const m_EM_SVX = 73
const m_EM_TILE64 = 187
const m_EM_TILEGX = 191
const m_EM_TILEPRO = 188
const m_EM_TINYJ = 61
const m_EM_TI_ARP32 = 143
const m_EM_TI_C2000 = 141
const m_EM_TI_C5500 = 142
const m_EM_TI_C6000 = 140
const m_EM_TI_PRU = 144
const m_EM_TMM_GPP = 96
const m_EM_TPC = 98
const m_EM_TRICORE = 44
const m_EM_TRIMEDIA = 163
const m_EM_TSK3000 = 131
const m_EM_UNICORE = 110
const m_EM_V800 = 36
const m_EM_V850 = 87
const m_EM_VAX = 75
const m_EM_VIDEOCORE = 95
const m_EM_VIDEOCORE3 = 137
const m_EM_VIDEOCORE5 = 198
const m_EM_VISIUM = 221
const m_EM_VPP500 = 17
const m_EM_X86_64 = 62
const m_EM_XCORE = 203
const m_EM_XGATE = 115
const m_EM_XIMO16 = 170
const m_EM_XTENSA = 94
const m_EM_Z80 = 220
const m_EM_ZSP = 79
const m_ET_CORE = 4
const m_ET_DYN = 3
const m_ET_EXEC = 2
const m_ET_HIOS = 0xfeff
const m_ET_HIPROC = 0xffff
const m_ET_LOOS = 0xfe00
const m_ET_LOPROC = 0xff00
const m_ET_NONE = 0
const m_ET_NUM = 5
const m_ET_REL = 1
const m_EV_CURRENT = 1
const m_EV_NONE = 0
const m_EV_NUM = 2
const m_E_MIPS_ARCH_1 = 0x00000000
const m_E_MIPS_ARCH_2 = 0x10000000
const m_E_MIPS_ARCH_3 = 0x20000000
const m_E_MIPS_ARCH_32 = 0x50000000
const m_E_MIPS_ARCH_4 = 0x30000000
const m_E_MIPS_ARCH_5 = 0x40000000
const m_E_MIPS_ARCH_64 = 0x60000000
const m_Elf_Ehdr = "Elf64_Ehdr"
const m_Elf_Off = "Elf64_Off"
const m_Elf_Shdr = "Elf64_Shdr"
const m_Elf_Sword = "Elf64_Sword"
const m_Elf_Sym = "Elf64_Sym"
const m_Elf_Word = "Elf64_Word"
const m_GRP_COMDAT = 0x1
const m_LITUSE_ALPHA_ADDR = 0
const m_LITUSE_ALPHA_BASE = 1
const m_LITUSE_ALPHA_BYTOFF = 2
const m_LITUSE_ALPHA_JSR = 3
const m_LITUSE_ALPHA_TLS_GD = 4
const m_LITUSE_ALPHA_TLS_LDM = 5
const m_LL_NONE = 0
const m_MAP_ANON1 = 0x20
const m_MAP_PRIVATE1 = 0x02
const m_MIPS_AFL_ASE_DSP = 0x00000001
const m_MIPS_AFL_ASE_DSPR2 = 0x00000002
const m_MIPS_AFL_ASE_EVA = 0x00000004
const m_MIPS_AFL_ASE_MASK = 0x00001fff
const m_MIPS_AFL_ASE_MCU = 0x00000008
const m_MIPS_AFL_ASE_MDMX = 0x00000010
const m_MIPS_AFL_ASE_MICROMIPS = 0x00000800
const m_MIPS_AFL_ASE_MIPS16 = 0x00000400
const m_MIPS_AFL_ASE_MIPS3D = 0x00000020
const m_MIPS_AFL_ASE_MSA = 0x00000200
const m_MIPS_AFL_ASE_MT = 0x00000040
const m_MIPS_AFL_ASE_SMARTMIPS = 0x00000080
const m_MIPS_AFL_ASE_VIRT = 0x00000100
const m_MIPS_AFL_ASE_XPA = 0x00001000
const m_MIPS_AFL_EXT_10000 = 11
const m_MIPS_AFL_EXT_3900 = 10
const m_MIPS_AFL_EXT_4010 = 8
const m_MIPS_AFL_EXT_4100 = 9
const m_MIPS_AFL_EXT_4111 = 13
const m_MIPS_AFL_EXT_4120 = 14
const m_MIPS_AFL_EXT_4650 = 7
const m_MIPS_AFL_EXT_5400 = 15
const m_MIPS_AFL_EXT_5500 = 16
const m_MIPS_AFL_EXT_5900 = 6
const m_MIPS_AFL_EXT_LOONGSON_2E = 17
const m_MIPS_AFL_EXT_LOONGSON_2F = 18
const m_MIPS_AFL_EXT_LOONGSON_3A = 4
const m_MIPS_AFL_EXT_OCTEON = 5
const m_MIPS_AFL_EXT_OCTEON2 = 2
const m_MIPS_AFL_EXT_OCTEONP = 3
const m_MIPS_AFL_EXT_SB1 = 12
const m_MIPS_AFL_EXT_XLR = 1
const m_MIPS_AFL_FLAGS1_ODDSPREG = 1
const m_MIPS_AFL_REG_128 = 0x03
const m_MIPS_AFL_REG_32 = 0x01
const m_MIPS_AFL_REG_64 = 0x02
const m_MIPS_AFL_REG_NONE = 0x00
const m_NT_386_IOPERM = 0x201
const m_NT_386_TLS = 0x200
const m_NT_ARC_V2 = 0x600
const m_NT_ARM_HW_BREAK = 0x402
const m_NT_ARM_HW_WATCH = 0x403
const m_NT_ARM_PACA_KEYS = 0x407
const m_NT_ARM_PACG_KEYS = 0x408
const m_NT_ARM_PAC_ENABLED_KEYS = 0x40a
const m_NT_ARM_PAC_MASK = 0x406
const m_NT_ARM_SVE = 0x405
const m_NT_ARM_SYSTEM_CALL = 0x404
const m_NT_ARM_TAGGED_ADDR_CTRL = 0x409
const m_NT_ARM_TLS = 0x401
const m_NT_ARM_VFP = 0x400
const m_NT_ASRS = 8
const m_NT_AUXV = 6
const m_NT_FILE = 0x46494c45
const m_NT_FPREGSET = 2
const m_NT_GNU_ABI_TAG = 1
const m_NT_GNU_BUILD_ID = 3
const m_NT_GNU_GOLD_VERSION = 4
const m_NT_GNU_PROPERTY_TYPE_0 = 5
const m_NT_GWINDOWS = 7
const m_NT_LOONGARCH_CPUCFG = 0xa00
const m_NT_LOONGARCH_CSR = 0xa01
const m_NT_LOONGARCH_LASX = 0xa03
const m_NT_LOONGARCH_LBT = 0xa04
const m_NT_LOONGARCH_LSX = 0xa02
const m_NT_LWPSINFO = 17
const m_NT_LWPSTATUS = 16
const m_NT_METAG_CBUF = 0x500
const m_NT_METAG_RPIPE = 0x501
const m_NT_METAG_TLS = 0x502
const m_NT_MIPS_DSP = 0x800
const m_NT_MIPS_FP_MODE = 0x801
const m_NT_MIPS_MSA = 0x802
const m_NT_PLATFORM = 5
const m_NT_PPC_DSCR = 0x105
const m_NT_PPC_EBB = 0x106
const m_NT_PPC_PMU = 0x107
const m_NT_PPC_PPR = 0x104
const m_NT_PPC_SPE = 0x101
const m_NT_PPC_TAR = 0x103
const m_NT_PPC_TM_CDSCR = 0x10f
const m_NT_PPC_TM_CFPR = 0x109
const m_NT_PPC_TM_CGPR = 0x108
const m_NT_PPC_TM_CPPR = 0x10e
const m_NT_PPC_TM_CTAR = 0x10d
const m_NT_PPC_TM_CVMX = 0x10a
const m_NT_PPC_TM_CVSX = 0x10b
const m_NT_PPC_TM_SPR = 0x10c
const m_NT_PPC_VMX = 0x100
const m_NT_PPC_VSX = 0x102
const m_NT_PRCRED = 14
const m_NT_PRFPREG = 2
const m_NT_PRFPXREG = 20
const m_NT_PRPSINFO = 3
const m_NT_PRSTATUS = 1
const m_NT_PRXFPREG = 0x46e62b7f
const m_NT_PRXREG = 4
const m_NT_PSINFO = 13
const m_NT_PSTATUS = 10
const m_NT_RISCV_CSR = 0x900
const m_NT_RISCV_VECTOR = 0x901
const m_NT_S390_CTRS = 0x304
const m_NT_S390_GS_BC = 0x30c
const m_NT_S390_GS_CB = 0x30b
const m_NT_S390_HIGH_GPRS = 0x300
const m_NT_S390_LAST_BREAK = 0x306
const m_NT_S390_PREFIX = 0x305
const m_NT_S390_RI_CB = 0x30d
const m_NT_S390_SYSTEM_CALL = 0x307
const m_NT_S390_TDB = 0x308
const m_NT_S390_TIMER = 0x301
const m_NT_S390_TODCMP = 0x302
const m_NT_S390_TODPREG = 0x303
const m_NT_S390_VXRS_HIGH = 0x30a
const m_NT_S390_VXRS_LOW = 0x309
const m_NT_SIGINFO = 0x53494749
const m_NT_TASKSTRUCT = 4
const m_NT_UTSNAME = 15
const m_NT_VERSION = 1
const m_NT_VMCOREDD = 0x700
const m_NT_X86_XSTATE = 0x202
const m_N_ABS = 2
const m_N_BSS = 8
const m_N_COMM = 0x12
const m_N_DATA = 6
const m_N_EXT = 1
const m_N_FN = 30
const m_N_FORMAT = "%08x"
const m_N_INDR = 0x0a
const m_N_SETA = 0x14
const m_N_SETB = 0x1a
const m_N_SETD = 0x18
const m_N_SETT = 0x16
const m_N_SETV = 0x1c
const m_N_SIZE = 0x0c
const m_N_STAB = 0xe0
const m_N_TEXT = 4
const m_N_TYPE = 0x1e
const m_N_UNDF = 0
const m_N_WARN = 0x1e
const m_ODK_EXCEPTIONS = 2
const m_ODK_FILL = 5
const m_ODK_HWAND = 7
const m_ODK_HWOR = 8
const m_ODK_HWPATCH = 4
const m_ODK_NULL = 0
const m_ODK_PAD = 3
const m_ODK_REGINFO = 1
const m_ODK_TAGS = 6
const m_OEX_DISMISS = 0x80000
const m_OEX_FPDBUG = 0x40000
const m_OEX_FPU_DIV0 = 0x08
const m_OEX_FPU_INEX = 0x01
const m_OEX_FPU_INVAL = 0x10
const m_OEX_FPU_MAX = 0x1f00
const m_OEX_FPU_MIN = 0x1f
const m_OEX_FPU_OFLO = 0x04
const m_OEX_FPU_UFLO = 0x02
const m_OEX_PAGE0 = 0x10000
const m_OEX_PRECISEFP = "OEX_FPDBUG"
const m_OEX_SMM = 0x20000
const m_OHWA0_R4KEOP_CHECKED = 0x00000001
const m_OHWA1_R4KEOP_CLEAN = 0x00000002
const m_OHW_R4KEOP = 0x1
const m_OHW_R5KCVTL = 0x8
const m_OHW_R5KEOP = 0x4
const m_OHW_R8KPFETCH = 0x2
const m_OPAD_POSTFIX = 0x2
const m_OPAD_PREFIX = 0x1
const m_OPAD_SYMBOL = 0x4
const m_O_CREAT2 = 0100
const m_O_NONBLOCK2 = 04000
const m_O_RDONLY1 = 0
const m_O_TRUNC2 = 01000
const m_PF_ARM_ABS = 0x40000000
const m_PF_ARM_PI = 0x20000000
const m_PF_ARM_SB = 0x10000000
const m_PF_HP_CODE = 0x01000000
const m_PF_HP_FAR_SHARED = 0x00200000
const m_PF_HP_LAZYSWAP = 0x04000000
const m_PF_HP_MODIFY = 0x02000000
const m_PF_HP_NEAR_SHARED = 0x00400000
const m_PF_HP_PAGE_SIZE = 0x00100000
const m_PF_HP_SBP = 0x08000000
const m_PF_IA_64_NORECOV = 0x80000000
const m_PF_MASKOS = 0x0ff00000
const m_PF_MASKPROC = 0xf0000000
const m_PF_MIPS_LOCAL = 0x10000000
const m_PF_PARISC_SBP = 0x08000000
const m_PN_XNUM = 0xffff
const m_PPC64_OPT_LOCALENTRY = 4
const m_PPC64_OPT_MULTI_TOC = 2
const m_PPC64_OPT_TLS = 1
const m_PPC_OPT_TLS = 1
const m_PT_DYNAMIC = 2
const m_PT_GNU_EH_FRAME = 0x6474e550
const m_PT_GNU_PROPERTY = 0x6474e553
const m_PT_GNU_RELRO = 0x6474e552
const m_PT_GNU_STACK = 0x6474e551
const m_PT_HIOS = 0x6fffffff
const m_PT_HIPROC = 0x7fffffff
const m_PT_HISUNW = 0x6fffffff
const m_PT_INTERP = 3
const m_PT_LOAD = 1
const m_PT_LOOS = 0x60000000
const m_PT_LOPROC = 0x70000000
const m_PT_LOSUNW = 0x6ffffffa
const m_PT_MIPS_ABIFLAGS = 0x70000003
const m_PT_MIPS_OPTIONS = 0x70000002
const m_PT_MIPS_REGINFO = 0x70000000
const m_PT_MIPS_RTPROC = 0x70000001
const m_PT_NOTE = 4
const m_PT_NULL = 0
const m_PT_NUM = 8
const m_PT_PARISC_ARCHEXT = 0x70000000
const m_PT_PARISC_UNWIND = 0x70000001
const m_PT_PHDR = 6
const m_PT_SHLIB = 5
const m_PT_SUNWBSS = 0x6ffffffa
const m_PT_SUNWSTACK = 0x6ffffffb
const m_PT_TLS = 7
const m_RHF_NONE = 0
const m_R_386_16 = 20
const m_R_386_32 = 1
const m_R_386_32PLT = 11
const m_R_386_8 = 22
const m_R_386_COPY = 5
const m_R_386_GLOB_DAT = 6
const m_R_386_GOT32 = 3
const m_R_386_GOT32X = 43
const m_R_386_GOTOFF = 9
const m_R_386_GOTPC = 10
const m_R_386_IRELATIVE = 42
const m_R_386_JMP_SLOT = 7
const m_R_386_NONE = 0
const m_R_386_NUM = 44
const m_R_386_PC16 = 21
const m_R_386_PC32 = 2
const m_R_386_PC8 = 23
const m_R_386_PLT32 = 4
const m_R_386_RELATIVE = 8
const m_R_386_SIZE32 = 38
const m_R_386_TLS_DESC = 41
const m_R_386_TLS_DESC_CALL = 40
const m_R_386_TLS_DTPMOD32 = 35
const m_R_386_TLS_DTPOFF32 = 36
const m_R_386_TLS_GD = 18
const m_R_386_TLS_GD_32 = 24
const m_R_386_TLS_GD_CALL = 26
const m_R_386_TLS_GD_POP = 27
const m_R_386_TLS_GD_PUSH = 25
const m_R_386_TLS_GOTDESC = 39
const m_R_386_TLS_GOTIE = 16
const m_R_386_TLS_IE = 15
const m_R_386_TLS_IE_32 = 33
const m_R_386_TLS_LDM = 19
const m_R_386_TLS_LDM_32 = 28
const m_R_386_TLS_LDM_CALL = 30
const m_R_386_TLS_LDM_POP = 31
const m_R_386_TLS_LDM_PUSH = 29
const m_R_386_TLS_LDO_32 = 32
const m_R_386_TLS_LE = 17
const m_R_386_TLS_LE_32 = 34
const m_R_386_TLS_TPOFF = 14
const m_R_386_TLS_TPOFF32 = 37
const m_R_390_12 = 2
const m_R_390_16 = 3
const m_R_390_20 = 57
const m_R_390_32 = 4
const m_R_390_64 = 22
const m_R_390_8 = 1
const m_R_390_COPY = 9
const m_R_390_GLOB_DAT = 10
const m_R_390_GOT12 = 6
const m_R_390_GOT16 = 15
const m_R_390_GOT20 = 58
const m_R_390_GOT32 = 7
const m_R_390_GOT64 = 24
const m_R_390_GOTENT = 26
const m_R_390_GOTOFF16 = 27
const m_R_390_GOTOFF32 = 13
const m_R_390_GOTOFF64 = 28
const m_R_390_GOTPC = 14
const m_R_390_GOTPCDBL = 21
const m_R_390_GOTPLT12 = 29
const m_R_390_GOTPLT16 = 30
const m_R_390_GOTPLT20 = 59
const m_R_390_GOTPLT32 = 31
const m_R_390_GOTPLT64 = 32
const m_R_390_GOTPLTENT = 33
const m_R_390_JMP_SLOT = 11
const m_R_390_NONE = 0
const m_R_390_NUM = 61
const m_R_390_PC16 = 16
const m_R_390_PC16DBL = 17
const m_R_390_PC32 = 5
const m_R_390_PC32DBL = 19
const m_R_390_PC64 = 23
const m_R_390_PLT16DBL = 18
const m_R_390_PLT32 = 8
const m_R_390_PLT32DBL = 20
const m_R_390_PLT64 = 25
const m_R_390_PLTOFF16 = 34
const m_R_390_PLTOFF32 = 35
const m_R_390_PLTOFF64 = 36
const m_R_390_RELATIVE = 12
const m_R_390_TLS_DTPMOD = 54
const m_R_390_TLS_DTPOFF = 55
const m_R_390_TLS_GD32 = 40
const m_R_390_TLS_GD64 = 41
const m_R_390_TLS_GDCALL = 38
const m_R_390_TLS_GOTIE12 = 42
const m_R_390_TLS_GOTIE20 = 60
const m_R_390_TLS_GOTIE32 = 43
const m_R_390_TLS_GOTIE64 = 44
const m_R_390_TLS_IE32 = 47
const m_R_390_TLS_IE64 = 48
const m_R_390_TLS_IEENT = 49
const m_R_390_TLS_LDCALL = 39
const m_R_390_TLS_LDM32 = 45
const m_R_390_TLS_LDM64 = 46
const m_R_390_TLS_LDO32 = 52
const m_R_390_TLS_LDO64 = 53
const m_R_390_TLS_LE32 = 50
const m_R_390_TLS_LE64 = 51
const m_R_390_TLS_LOAD = 37
const m_R_390_TLS_TPOFF = 56
const m_R_68K_16 = 2
const m_R_68K_32 = 1
const m_R_68K_8 = 3
const m_R_68K_COPY = 19
const m_R_68K_GLOB_DAT = 20
const m_R_68K_GOT16 = 8
const m_R_68K_GOT16O = 11
const m_R_68K_GOT32 = 7
const m_R_68K_GOT32O = 10
const m_R_68K_GOT8 = 9
const m_R_68K_GOT8O = 12
const m_R_68K_JMP_SLOT = 21
const m_R_68K_NONE = 0
const m_R_68K_NUM = 43
const m_R_68K_PC16 = 5
const m_R_68K_PC32 = 4
const m_R_68K_PC8 = 6
const m_R_68K_PLT16 = 14
const m_R_68K_PLT16O = 17
const m_R_68K_PLT32 = 13
const m_R_68K_PLT32O = 16
const m_R_68K_PLT8 = 15
const m_R_68K_PLT8O = 18
const m_R_68K_RELATIVE = 22
const m_R_68K_TLS_DTPMOD32 = 40
const m_R_68K_TLS_DTPREL32 = 41
const m_R_68K_TLS_GD16 = 26
const m_R_68K_TLS_GD32 = 25
const m_R_68K_TLS_GD8 = 27
const m_R_68K_TLS_IE16 = 35
const m_R_68K_TLS_IE32 = 34
const m_R_68K_TLS_IE8 = 36
const m_R_68K_TLS_LDM16 = 29
const m_R_68K_TLS_LDM32 = 28
const m_R_68K_TLS_LDM8 = 30
const m_R_68K_TLS_LDO16 = 32
const m_R_68K_TLS_LDO32 = 31
const m_R_68K_TLS_LDO8 = 33
const m_R_68K_TLS_LE16 = 38
const m_R_68K_TLS_LE32 = 37
const m_R_68K_TLS_LE8 = 39
const m_R_68K_TLS_TPREL32 = 42
const m_R_AARCH64_ABS16 = 259
const m_R_AARCH64_ABS32 = 258
const m_R_AARCH64_ABS64 = 257
const m_R_AARCH64_ADD_ABS_LO12_NC = 277
const m_R_AARCH64_ADR_GOT_PAGE = 311
const m_R_AARCH64_ADR_PREL_LO21 = 274
const m_R_AARCH64_ADR_PREL_PG_HI21 = 275
const m_R_AARCH64_ADR_PREL_PG_HI21_NC = 276
const m_R_AARCH64_CALL26 = 283
const m_R_AARCH64_CONDBR19 = 280
const m_R_AARCH64_COPY = 1024
const m_R_AARCH64_GLOB_DAT = 1025
const m_R_AARCH64_GOTREL32 = 308
const m_R_AARCH64_GOTREL64 = 307
const m_R_AARCH64_GOT_LD_PREL19 = 309
const m_R_AARCH64_JUMP26 = 282
const m_R_AARCH64_JUMP_SLOT = 1026
const m_R_AARCH64_LD64_GOTOFF_LO15 = 310
const m_R_AARCH64_LD64_GOTPAGE_LO15 = 313
const m_R_AARCH64_LD64_GOT_LO12_NC = 312
const m_R_AARCH64_LDST128_ABS_LO12_NC = 299
const m_R_AARCH64_LDST16_ABS_LO12_NC = 284
const m_R_AARCH64_LDST32_ABS_LO12_NC = 285
const m_R_AARCH64_LDST64_ABS_LO12_NC = 286
const m_R_AARCH64_LDST8_ABS_LO12_NC = 278
const m_R_AARCH64_LD_PREL_LO19 = 273
const m_R_AARCH64_MOVW_GOTOFF_G0 = 300
const m_R_AARCH64_MOVW_GOTOFF_G0_NC = 301
const m_R_AARCH64_MOVW_GOTOFF_G1 = 302
const m_R_AARCH64_MOVW_GOTOFF_G1_NC = 303
const m_R_AARCH64_MOVW_GOTOFF_G2 = 304
const m_R_AARCH64_MOVW_GOTOFF_G2_NC = 305
const m_R_AARCH64_MOVW_GOTOFF_G3 = 306
const m_R_AARCH64_MOVW_PREL_G0 = 287
const m_R_AARCH64_MOVW_PREL_G0_NC = 288
const m_R_AARCH64_MOVW_PREL_G1 = 289
const m_R_AARCH64_MOVW_PREL_G1_NC = 290
const m_R_AARCH64_MOVW_PREL_G2 = 291
const m_R_AARCH64_MOVW_PREL_G2_NC = 292
const m_R_AARCH64_MOVW_PREL_G3 = 293
const m_R_AARCH64_MOVW_SABS_G0 = 270
const m_R_AARCH64_MOVW_SABS_G1 = 271
const m_R_AARCH64_MOVW_SABS_G2 = 272
const m_R_AARCH64_MOVW_UABS_G0 = 263
const m_R_AARCH64_MOVW_UABS_G0_NC = 264
const m_R_AARCH64_MOVW_UABS_G1 = 265
const m_R_AARCH64_MOVW_UABS_G1_NC = 266
const m_R_AARCH64_MOVW_UABS_G2 = 267
const m_R_AARCH64_MOVW_UABS_G2_NC = 268
const m_R_AARCH64_MOVW_UABS_G3 = 269
const m_R_AARCH64_NONE = 0
const m_R_AARCH64_P32_ABS32 = 1
const m_R_AARCH64_P32_COPY = 180
const m_R_AARCH64_P32_GLOB_DAT = 181
const m_R_AARCH64_P32_IRELATIVE = 188
const m_R_AARCH64_P32_JUMP_SLOT = 182
const m_R_AARCH64_P32_RELATIVE = 183
const m_R_AARCH64_P32_TLSDESC = 187
const m_R_AARCH64_P32_TLS_DTPMOD = 184
const m_R_AARCH64_P32_TLS_DTPREL = 185
const m_R_AARCH64_P32_TLS_TPREL = 186
const m_R_AARCH64_PREL16 = 262
const m_R_AARCH64_PREL32 = 261
const m_R_AARCH64_PREL64 = 260
const m_R_AARCH64_RELATIVE = 1027
const m_R_AARCH64_TLSDESC = 1031
const m_R_AARCH64_TLSDESC_ADD = 568
const m_R_AARCH64_TLSDESC_ADD_LO12 = 564
const m_R_AARCH64_TLSDESC_ADR_PAGE21 = 562
const m_R_AARCH64_TLSDESC_ADR_PREL21 = 561
const m_R_AARCH64_TLSDESC_CALL = 569
const m_R_AARCH64_TLSDESC_LD64_LO12 = 563
const m_R_AARCH64_TLSDESC_LDR = 567
const m_R_AARCH64_TLSDESC_LD_PREL19 = 560
const m_R_AARCH64_TLSDESC_OFF_G0_NC = 566
const m_R_AARCH64_TLSDESC_OFF_G1 = 565
const m_R_AARCH64_TLSGD_ADD_LO12_NC = 514
const m_R_AARCH64_TLSGD_ADR_PAGE21 = 513
const m_R_AARCH64_TLSGD_ADR_PREL21 = 512
const m_R_AARCH64_TLSGD_MOVW_G0_NC = 516
const m_R_AARCH64_TLSGD_MOVW_G1 = 515
const m_R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21 = 541
const m_R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC = 542
const m_R_AARCH64_TLSIE_LD_GOTTPREL_PREL19 = 543
const m_R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC = 540
const m_R_AARCH64_TLSIE_MOVW_GOTTPREL_G1 = 539
const m_R_AARCH64_TLSLD_ADD_DTPREL_HI12 = 528
const m_R_AARCH64_TLSLD_ADD_DTPREL_LO12 = 529
const m_R_AARCH64_TLSLD_ADD_DTPREL_LO12_NC = 530
const m_R_AARCH64_TLSLD_ADD_LO12_NC = 519
const m_R_AARCH64_TLSLD_ADR_PAGE21 = 518
const m_R_AARCH64_TLSLD_ADR_PREL21 = 517
const m_R_AARCH64_TLSLD_LDST128_DTPREL_LO12 = 572
const m_R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC = 573
const m_R_AARCH64_TLSLD_LDST16_DTPREL_LO12 = 533
const m_R_AARCH64_TLSLD_LDST16_DTPREL_LO12_NC = 534
const m_R_AARCH64_TLSLD_LDST32_DTPREL_LO12 = 535
const m_R_AARCH64_TLSLD_LDST32_DTPREL_LO12_NC = 536
const m_R_AARCH64_TLSLD_LDST64_DTPREL_LO12 = 537
const m_R_AARCH64_TLSLD_LDST64_DTPREL_LO12_NC = 538
const m_R_AARCH64_TLSLD_LDST8_DTPREL_LO12 = 531
const m_R_AARCH64_TLSLD_LDST8_DTPREL_LO12_NC = 532
const m_R_AARCH64_TLSLD_LD_PREL19 = 522
const m_R_AARCH64_TLSLD_MOVW_DTPREL_G0 = 526
const m_R_AARCH64_TLSLD_MOVW_DTPREL_G0_NC = 527
const m_R_AARCH64_TLSLD_MOVW_DTPREL_G1 = 524
const m_R_AARCH64_TLSLD_MOVW_DTPREL_G1_NC = 525
const m_R_AARCH64_TLSLD_MOVW_DTPREL_G2 = 523
const m_R_AARCH64_TLSLD_MOVW_G0_NC = 521
const m_R_AARCH64_TLSLD_MOVW_G1 = 520
const m_R_AARCH64_TLSLE_ADD_TPREL_HI12 = 549
const m_R_AARCH64_TLSLE_ADD_TPREL_LO12 = 550
const m_R_AARCH64_TLSLE_ADD_TPREL_LO12_NC = 551
const m_R_AARCH64_TLSLE_LDST128_TPREL_LO12 = 570
const m_R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC = 571
const m_R_AARCH64_TLSLE_LDST16_TPREL_LO12 = 554
const m_R_AARCH64_TLSLE_LDST16_TPREL_LO12_NC = 555
const m_R_AARCH64_TLSLE_LDST32_TPREL_LO12 = 556
const m_R_AARCH64_TLSLE_LDST32_TPREL_LO12_NC = 557
const m_R_AARCH64_TLSLE_LDST64_TPREL_LO12 = 558
const m_R_AARCH64_TLSLE_LDST64_TPREL_LO12_NC = 559
const m_R_AARCH64_TLSLE_LDST8_TPREL_LO12 = 552
const m_R_AARCH64_TLSLE_LDST8_TPREL_LO12_NC = 553
const m_R_AARCH64_TLSLE_MOVW_TPREL_G0 = 547
const m_R_AARCH64_TLSLE_MOVW_TPREL_G0_NC = 548
const m_R_AARCH64_TLSLE_MOVW_TPREL_G1 = 545
const m_R_AARCH64_TLSLE_MOVW_TPREL_G1_NC = 546
const m_R_AARCH64_TLSLE_MOVW_TPREL_G2 = 544
const m_R_AARCH64_TLS_DTPMOD = 1028
const m_R_AARCH64_TLS_DTPMOD64 = 1028
const m_R_AARCH64_TLS_DTPREL = 1029
const m_R_AARCH64_TLS_DTPREL64 = 1029
const m_R_AARCH64_TLS_TPREL = 1030
const m_R_AARCH64_TLS_TPREL64 = 1030
const m_R_AARCH64_TSTBR14 = 279
const m_R_ALPHA_BRADDR = 7
const m_R_ALPHA_COPY = 24
const m_R_ALPHA_DTPMOD64 = 31
const m_R_ALPHA_DTPREL16 = 36
const m_R_ALPHA_DTPREL64 = 33
const m_R_ALPHA_DTPRELHI = 34
const m_R_ALPHA_DTPRELLO = 35
const m_R_ALPHA_GLOB_DAT = 25
const m_R_ALPHA_GOTDTPREL = 32
const m_R_ALPHA_GOTTPREL = 37
const m_R_ALPHA_GPDISP = 6
const m_R_ALPHA_GPREL16 = 19
const m_R_ALPHA_GPREL32 = 3
const m_R_ALPHA_GPRELHIGH = 17
const m_R_ALPHA_GPRELLOW = 18
const m_R_ALPHA_HINT = 8
const m_R_ALPHA_JMP_SLOT = 26
const m_R_ALPHA_LITERAL = 4
const m_R_ALPHA_LITUSE = 5
const m_R_ALPHA_NONE = 0
const m_R_ALPHA_NUM = 46
const m_R_ALPHA_REFLONG = 1
const m_R_ALPHA_REFQUAD = 2
const m_R_ALPHA_RELATIVE = 27
const m_R_ALPHA_SREL16 = 9
const m_R_ALPHA_SREL32 = 10
const m_R_ALPHA_SREL64 = 11
const m_R_ALPHA_TLSGD = 29
const m_R_ALPHA_TLS_GD_HI = 28
const m_R_ALPHA_TLS_LDM = 30
const m_R_ALPHA_TPREL16 = 41
const m_R_ALPHA_TPREL64 = 38
const m_R_ALPHA_TPRELHI = 39
const m_R_ALPHA_TPRELLO = 40
const m_R_ARM_ABS12 = 6
const m_R_ARM_ABS16 = 5
const m_R_ARM_ABS32 = 2
const m_R_ARM_ABS32_NOI = 55
const m_R_ARM_ABS8 = 8
const m_R_ARM_ALU_PCREL_15_8 = 33
const m_R_ARM_ALU_PCREL_23_15 = 34
const m_R_ARM_ALU_PCREL_7_0 = 32
const m_R_ARM_ALU_PC_G0 = 58
const m_R_ARM_ALU_PC_G0_NC = 57
const m_R_ARM_ALU_PC_G1 = 60
const m_R_ARM_ALU_PC_G1_NC = 59
const m_R_ARM_ALU_PC_G2 = 61
const m_R_ARM_ALU_SBREL_19_12 = 36
const m_R_ARM_ALU_SBREL_27_20 = 37
const m_R_ARM_ALU_SB_G0 = 71
const m_R_ARM_ALU_SB_G0_NC = 70
const m_R_ARM_ALU_SB_G1 = 73
const m_R_ARM_ALU_SB_G1_NC = 72
const m_R_ARM_ALU_SB_G2 = 74
const m_R_ARM_AMP_VCALL9 = 12
const m_R_ARM_BASE_ABS = 31
const m_R_ARM_CALL = 28
const m_R_ARM_COPY = 20
const m_R_ARM_GLOB_DAT = 21
const m_R_ARM_GNU_VTENTRY = 100
const m_R_ARM_GNU_VTINHERIT = 101
const m_R_ARM_GOT32 = 26
const m_R_ARM_GOTOFF = 24
const m_R_ARM_GOTOFF12 = 98
const m_R_ARM_GOTPC = 25
const m_R_ARM_GOTRELAX = 99
const m_R_ARM_GOT_ABS = 95
const m_R_ARM_GOT_BREL12 = 97
const m_R_ARM_GOT_PREL = 96
const m_R_ARM_IRELATIVE = 160
const m_R_ARM_JUMP24 = 29
const m_R_ARM_JUMP_SLOT = 22
const m_R_ARM_LDC_PC_G0 = 67
const m_R_ARM_LDC_PC_G1 = 68
const m_R_ARM_LDC_PC_G2 = 69
const m_R_ARM_LDC_SB_G0 = 81
const m_R_ARM_LDC_SB_G1 = 82
const m_R_ARM_LDC_SB_G2 = 83
const m_R_ARM_LDRS_PC_G0 = 64
const m_R_ARM_LDRS_PC_G1 = 65
const m_R_ARM_LDRS_PC_G2 = 66
const m_R_ARM_LDRS_SB_G0 = 78
const m_R_ARM_LDRS_SB_G1 = 79
const m_R_ARM_LDRS_SB_G2 = 80
const m_R_ARM_LDR_PC_G1 = 62
const m_R_ARM_LDR_PC_G2 = 63
const m_R_ARM_LDR_SBREL_11_0 = 35
const m_R_ARM_LDR_SB_G0 = 75
const m_R_ARM_LDR_SB_G1 = 76
const m_R_ARM_LDR_SB_G2 = 77
const m_R_ARM_ME_TOO = 128
const m_R_ARM_MOVT_ABS = 44
const m_R_ARM_MOVT_BREL = 85
const m_R_ARM_MOVT_PREL = 46
const m_R_ARM_MOVW_ABS_NC = 43
const m_R_ARM_MOVW_BREL = 86
const m_R_ARM_MOVW_BREL_NC = 84
const m_R_ARM_MOVW_PREL_NC = 45
const m_R_ARM_NONE = 0
const m_R_ARM_NUM = 256
const m_R_ARM_PC13 = 4
const m_R_ARM_PC24 = 1
const m_R_ARM_PLT32 = 27
const m_R_ARM_PLT32_ABS = 94
const m_R_ARM_PREL31 = 42
const m_R_ARM_RABS22 = 253
const m_R_ARM_RBASE = 255
const m_R_ARM_REL32 = 3
const m_R_ARM_REL32_NOI = 56
const m_R_ARM_RELATIVE = 23
const m_R_ARM_RPC24 = 254
const m_R_ARM_RREL32 = 252
const m_R_ARM_RSBREL32 = 250
const m_R_ARM_RXPC25 = 249
const m_R_ARM_SBREL31 = 39
const m_R_ARM_SBREL32 = 9
const m_R_ARM_TARGET1 = 38
const m_R_ARM_TARGET2 = 41
const m_R_ARM_THM_ABS5 = 7
const m_R_ARM_THM_ALU_PREL_11_0 = 53
const m_R_ARM_THM_GOT_BREL12 = 131
const m_R_ARM_THM_JUMP19 = 51
const m_R_ARM_THM_JUMP24 = 30
const m_R_ARM_THM_JUMP6 = 52
const m_R_ARM_THM_MOVT_ABS = 48
const m_R_ARM_THM_MOVT_BREL = 88
const m_R_ARM_THM_MOVT_PREL = 50
const m_R_ARM_THM_MOVW_ABS_NC = 47
const m_R_ARM_THM_MOVW_BREL = 89
const m_R_ARM_THM_MOVW_BREL_NC = 87
const m_R_ARM_THM_MOVW_PREL_NC = 49
const m_R_ARM_THM_PC11 = 102
const m_R_ARM_THM_PC12 = 54
const m_R_ARM_THM_PC22 = 10
const m_R_ARM_THM_PC8 = 11
const m_R_ARM_THM_PC9 = 103
const m_R_ARM_THM_RPC22 = 251
const m_R_ARM_THM_SWI8 = 14
const m_R_ARM_THM_TLS_CALL = 93
const m_R_ARM_THM_TLS_DESCSEQ = 129
const m_R_ARM_THM_TLS_DESCSEQ16 = 129
const m_R_ARM_THM_TLS_DESCSEQ32 = 130
const m_R_ARM_THM_XPC22 = 16
const m_R_ARM_TLS_CALL = 91
const m_R_ARM_TLS_DESC = 13
const m_R_ARM_TLS_DESCSEQ = 92
const m_R_ARM_TLS_DTPMOD32 = 17
const m_R_ARM_TLS_DTPOFF32 = 18
const m_R_ARM_TLS_GD32 = 104
const m_R_ARM_TLS_GOTDESC = 90
const m_R_ARM_TLS_IE12GP = 111
const m_R_ARM_TLS_IE32 = 107
const m_R_ARM_TLS_LDM32 = 105
const m_R_ARM_TLS_LDO12 = 109
const m_R_ARM_TLS_LDO32 = 106
const m_R_ARM_TLS_LE12 = 110
const m_R_ARM_TLS_LE32 = 108
const m_R_ARM_TLS_TPOFF32 = 19
const m_R_ARM_V4BX = 40
const m_R_ARM_XPC25 = 15
const m_R_BPF_MAP_FD = 1
const m_R_BPF_NONE = 0
const m_R_CKCORE_ADDR32 = 1
const m_R_CKCORE_ADDRGOT = 17
const m_R_CKCORE_ADDRGOT_HI16 = 36
const m_R_CKCORE_ADDRGOT_LO16 = 37
const m_R_CKCORE_ADDRPLT = 18
const m_R_CKCORE_ADDRPLT_HI16 = 38
const m_R_CKCORE_ADDRPLT_LO16 = 39
const m_R_CKCORE_ADDR_HI16 = 24
const m_R_CKCORE_ADDR_LO16 = 25
const m_R_CKCORE_COPY = 10
const m_R_CKCORE_DOFFSET_IMM18 = 44
const m_R_CKCORE_DOFFSET_IMM18BY2 = 45
const m_R_CKCORE_DOFFSET_IMM18BY4 = 46
const m_R_CKCORE_DOFFSET_LO16 = 42
const m_R_CKCORE_GLOB_DAT = 11
const m_R_CKCORE_GOT12 = 30
const m_R_CKCORE_GOT32 = 15
const m_R_CKCORE_GOTOFF = 13
const m_R_CKCORE_GOTOFF_HI16 = 28
const m_R_CKCORE_GOTOFF_LO16 = 29
const m_R_CKCORE_GOTPC = 14
const m_R_CKCORE_GOTPC_HI16 = 26
const m_R_CKCORE_GOTPC_LO16 = 27
const m_R_CKCORE_GOT_HI16 = 31
const m_R_CKCORE_GOT_IMM18BY4 = 48
const m_R_CKCORE_GOT_LO16 = 32
const m_R_CKCORE_JUMP_SLOT = 12
const m_R_CKCORE_NONE = 0
const m_R_CKCORE_PCREL32 = 5
const m_R_CKCORE_PCRELIMM11BY2 = 3
const m_R_CKCORE_PCRELIMM8BY4 = 2
const m_R_CKCORE_PCRELJSR_IMM11BY2 = 6
const m_R_CKCORE_PCREL_IMM10BY2 = 22
const m_R_CKCORE_PCREL_IMM10BY4 = 23
const m_R_CKCORE_PCREL_IMM16BY2 = 20
const m_R_CKCORE_PCREL_IMM16BY4 = 21
const m_R_CKCORE_PCREL_IMM18BY2 = 43
const m_R_CKCORE_PCREL_IMM26BY2 = 19
const m_R_CKCORE_PCREL_IMM7BY4 = 50
const m_R_CKCORE_PCREL_JSR_IMM26BY2 = 40
const m_R_CKCORE_PLT12 = 33
const m_R_CKCORE_PLT32 = 16
const m_R_CKCORE_PLT_HI16 = 34
const m_R_CKCORE_PLT_IMM18BY4 = 49
const m_R_CKCORE_PLT_LO16 = 35
const m_R_CKCORE_RELATIVE = 9
const m_R_CKCORE_TLS_DTPMOD32 = 56
const m_R_CKCORE_TLS_DTPOFF32 = 57
const m_R_CKCORE_TLS_GD32 = 53
const m_R_CKCORE_TLS_IE32 = 52
const m_R_CKCORE_TLS_LDM32 = 54
const m_R_CKCORE_TLS_LDO32 = 55
const m_R_CKCORE_TLS_LE32 = 51
const m_R_CKCORE_TLS_TPOFF32 = 58
const m_R_CKCORE_TOFFSET_LO16 = 41
const m_R_CRIS_16 = 2
const m_R_CRIS_16_GOT = 13
const m_R_CRIS_16_GOTPLT = 15
const m_R_CRIS_16_PCREL = 5
const m_R_CRIS_32 = 3
const m_R_CRIS_32_GOT = 14
const m_R_CRIS_32_GOTPLT = 16
const m_R_CRIS_32_GOTREL = 17
const m_R_CRIS_32_PCREL = 6
const m_R_CRIS_32_PLT_GOTREL = 18
const m_R_CRIS_32_PLT_PCREL = 19
const m_R_CRIS_8 = 1
const m_R_CRIS_8_PCREL = 4
const m_R_CRIS_COPY = 9
const m_R_CRIS_GLOB_DAT = 10
const m_R_CRIS_GNU_VTENTRY = 8
const m_R_CRIS_GNU_VTINHERIT = 7
const m_R_CRIS_JUMP_SLOT = 11
const m_R_CRIS_NONE = 0
const m_R_CRIS_NUM = 20
const m_R_CRIS_RELATIVE = 12
const m_R_IA64_COPY = 0x84
const m_R_IA64_DIR32LSB = 0x25
const m_R_IA64_DIR32MSB = 0x24
const m_R_IA64_DIR64LSB = 0x27
const m_R_IA64_DIR64MSB = 0x26
const m_R_IA64_DTPMOD64LSB = 0xa7
const m_R_IA64_DTPMOD64MSB = 0xa6
const m_R_IA64_DTPREL14 = 0xb1
const m_R_IA64_DTPREL22 = 0xb2
const m_R_IA64_DTPREL32LSB = 0xb5
const m_R_IA64_DTPREL32MSB = 0xb4
const m_R_IA64_DTPREL64I = 0xb3
const m_R_IA64_DTPREL64LSB = 0xb7
const m_R_IA64_DTPREL64MSB = 0xb6
const m_R_IA64_FPTR32LSB = 0x45
const m_R_IA64_FPTR32MSB = 0x44
const m_R_IA64_FPTR64I = 0x43
const m_R_IA64_FPTR64LSB = 0x47
const m_R_IA64_FPTR64MSB = 0x46
const m_R_IA64_GPREL22 = 0x2a
const m_R_IA64_GPREL32LSB = 0x2d
const m_R_IA64_GPREL32MSB = 0x2c
const m_R_IA64_GPREL64I = 0x2b
const m_R_IA64_GPREL64LSB = 0x2f
const m_R_IA64_GPREL64MSB = 0x2e
const m_R_IA64_IMM14 = 0x21
const m_R_IA64_IMM22 = 0x22
const m_R_IA64_IMM64 = 0x23
const m_R_IA64_IPLTLSB = 0x81
const m_R_IA64_IPLTMSB = 0x80
const m_R_IA64_LDXMOV = 0x87
const m_R_IA64_LTOFF22 = 0x32
const m_R_IA64_LTOFF22X = 0x86
const m_R_IA64_LTOFF64I = 0x33
const m_R_IA64_LTOFF_DTPMOD22 = 0xaa
const m_R_IA64_LTOFF_DTPREL22 = 0xba
const m_R_IA64_LTOFF_FPTR22 = 0x52
const m_R_IA64_LTOFF_FPTR32LSB = 0x55
const m_R_IA64_LTOFF_FPTR32MSB = 0x54
const m_R_IA64_LTOFF_FPTR64I = 0x53
const m_R_IA64_LTOFF_FPTR64LSB = 0x57
const m_R_IA64_LTOFF_FPTR64MSB = 0x56
const m_R_IA64_LTOFF_TPREL22 = 0x9a
const m_R_IA64_LTV32LSB = 0x75
const m_R_IA64_LTV32MSB = 0x74
const m_R_IA64_LTV64LSB = 0x77
const m_R_IA64_LTV64MSB = 0x76
const m_R_IA64_NONE = 0x00
const m_R_IA64_PCREL21B = 0x49
const m_R_IA64_PCREL21BI = 0x79
const m_R_IA64_PCREL21F = 0x4b
const m_R_IA64_PCREL21M = 0x4a
const m_R_IA64_PCREL22 = 0x7a
const m_R_IA64_PCREL32LSB = 0x4d
const m_R_IA64_PCREL32MSB = 0x4c
const m_R_IA64_PCREL60B = 0x48
const m_R_IA64_PCREL64I = 0x7b
const m_R_IA64_PCREL64LSB = 0x4f
const m_R_IA64_PCREL64MSB = 0x4e
const m_R_IA64_PLTOFF22 = 0x3a
const m_R_IA64_PLTOFF64I = 0x3b
const m_R_IA64_PLTOFF64LSB = 0x3f
const m_R_IA64_PLTOFF64MSB = 0x3e
const m_R_IA64_REL32LSB = 0x6d
const m_R_IA64_REL32MSB = 0x6c
const m_R_IA64_REL64LSB = 0x6f
const m_R_IA64_REL64MSB = 0x6e
const m_R_IA64_SECREL32LSB = 0x65
const m_R_IA64_SECREL32MSB = 0x64
const m_R_IA64_SECREL64LSB = 0x67
const m_R_IA64_SECREL64MSB = 0x66
const m_R_IA64_SEGREL32LSB = 0x5d
const m_R_IA64_SEGREL32MSB = 0x5c
const m_R_IA64_SEGREL64LSB = 0x5f
const m_R_IA64_SEGREL64MSB = 0x5e
const m_R_IA64_SUB = 0x85
const m_R_IA64_TPREL14 = 0x91
const m_R_IA64_TPREL22 = 0x92
const m_R_IA64_TPREL64I = 0x93
const m_R_IA64_TPREL64LSB = 0x97
const m_R_IA64_TPREL64MSB = 0x96
const m_R_LARCH_32 = 1
const m_R_LARCH_32_PCREL = 99
const m_R_LARCH_64 = 2
const m_R_LARCH_ABS64_HI12 = 70
const m_R_LARCH_ABS64_LO20 = 69
const m_R_LARCH_ABS_HI20 = 67
const m_R_LARCH_ABS_LO12 = 68
const m_R_LARCH_ADD16 = 48
const m_R_LARCH_ADD24 = 49
const m_R_LARCH_ADD32 = 50
const m_R_LARCH_ADD64 = 51
const m_R_LARCH_ADD8 = 47
const m_R_LARCH_B16 = 64
const m_R_LARCH_B21 = 65
const m_R_LARCH_B26 = 66
const m_R_LARCH_COPY = 4
const m_R_LARCH_GNU_VTENTRY = 58
const m_R_LARCH_GNU_VTINHERIT = 57
const m_R_LARCH_GOT64_HI12 = 82
const m_R_LARCH_GOT64_LO20 = 81
const m_R_LARCH_GOT64_PC_HI12 = 78
const m_R_LARCH_GOT64_PC_LO20 = 77
const m_R_LARCH_GOT_HI20 = 79
const m_R_LARCH_GOT_LO12 = 80
const m_R_LARCH_GOT_PC_HI20 = 75
const m_R_LARCH_GOT_PC_LO12 = 76
const m_R_LARCH_IRELATIVE = 12
const m_R_LARCH_JUMP_SLOT = 5
const m_R_LARCH_MARK_LA = 20
const m_R_LARCH_MARK_PCREL = 21
const m_R_LARCH_NONE = 0
const m_R_LARCH_PCALA64_HI12 = 74
const m_R_LARCH_PCALA64_LO20 = 73
const m_R_LARCH_PCALA_HI20 = 71
const m_R_LARCH_PCALA_LO12 = 72
const m_R_LARCH_RELATIVE = 3
const m_R_LARCH_RELAX = 100
const m_R_LARCH_SOP_ADD = 35
const m_R_LARCH_SOP_AND = 36
const m_R_LARCH_SOP_ASSERT = 30
const m_R_LARCH_SOP_IF_ELSE = 37
const m_R_LARCH_SOP_NOT = 31
const m_R_LARCH_SOP_POP_32_S_0_10_10_16_S2 = 45
const m_R_LARCH_SOP_POP_32_S_0_5_10_16_S2 = 44
const m_R_LARCH_SOP_POP_32_S_10_12 = 40
const m_R_LARCH_SOP_POP_32_S_10_16 = 41
const m_R_LARCH_SOP_POP_32_S_10_16_S2 = 42
const m_R_LARCH_SOP_POP_32_S_10_5 = 38
const m_R_LARCH_SOP_POP_32_S_5_20 = 43
const m_R_LARCH_SOP_POP_32_U = 46
const m_R_LARCH_SOP_POP_32_U_10_12 = 39
const m_R_LARCH_SOP_PUSH_ABSOLUTE = 23
const m_R_LARCH_SOP_PUSH_DUP = 24
const m_R_LARCH_SOP_PUSH_GPREL = 25
const m_R_LARCH_SOP_PUSH_PCREL = 22
const m_R_LARCH_SOP_PUSH_PLT_PCREL = 29
const m_R_LARCH_SOP_PUSH_TLS_GD = 28
const m_R_LARCH_SOP_PUSH_TLS_GOT = 27
const m_R_LARCH_SOP_PUSH_TLS_TPREL = 26
const m_R_LARCH_SOP_SL = 33
const m_R_LARCH_SOP_SR = 34
const m_R_LARCH_SOP_SUB = 32
const m_R_LARCH_SUB16 = 53
const m_R_LARCH_SUB24 = 54
const m_R_LARCH_SUB32 = 55
const m_R_LARCH_SUB64 = 56
const m_R_LARCH_SUB8 = 52
const m_R_LARCH_TLS_DTPMOD32 = 6
const m_R_LARCH_TLS_DTPMOD64 = 7
const m_R_LARCH_TLS_DTPREL32 = 8
const m_R_LARCH_TLS_DTPREL64 = 9
const m_R_LARCH_TLS_GD_HI20 = 98
const m_R_LARCH_TLS_GD_PC_HI20 = 97
const m_R_LARCH_TLS_IE64_HI12 = 94
const m_R_LARCH_TLS_IE64_LO20 = 93
const m_R_LARCH_TLS_IE64_PC_HI12 = 90
const m_R_LARCH_TLS_IE64_PC_LO20 = 89
const m_R_LARCH_TLS_IE_HI20 = 91
const m_R_LARCH_TLS_IE_LO12 = 92
const m_R_LARCH_TLS_IE_PC_HI20 = 87
const m_R_LARCH_TLS_IE_PC_LO12 = 88
const m_R_LARCH_TLS_LD_HI20 = 96
const m_R_LARCH_TLS_LD_PC_HI20 = 95
const m_R_LARCH_TLS_LE64_HI12 = 86
const m_R_LARCH_TLS_LE64_LO20 = 85
const m_R_LARCH_TLS_LE_HI20 = 83
const m_R_LARCH_TLS_LE_LO12 = 84
const m_R_LARCH_TLS_TPREL32 = 10
const m_R_LARCH_TLS_TPREL64 = 11
const m_R_M32R_10_PCREL = 4
const m_R_M32R_10_PCREL_RELA = 36
const m_R_M32R_16 = 1
const m_R_M32R_16_RELA = 33
const m_R_M32R_18_PCREL = 5
const m_R_M32R_18_PCREL_RELA = 37
const m_R_M32R_24 = 3
const m_R_M32R_24_RELA = 35
const m_R_M32R_26_PCREL = 6
const m_R_M32R_26_PCREL_RELA = 38
const m_R_M32R_26_PLTREL = 49
const m_R_M32R_32 = 2
const m_R_M32R_32_RELA = 34
const m_R_M32R_COPY = 50
const m_R_M32R_GLOB_DAT = 51
const m_R_M32R_GNU_VTENTRY = 12
const m_R_M32R_GNU_VTINHERIT = 11
const m_R_M32R_GOT16_HI_SLO = 57
const m_R_M32R_GOT16_HI_ULO = 56
const m_R_M32R_GOT16_LO = 58
const m_R_M32R_GOT24 = 48
const m_R_M32R_GOTOFF = 54
const m_R_M32R_GOTOFF_HI_SLO = 63
const m_R_M32R_GOTOFF_HI_ULO = 62
const m_R_M32R_GOTOFF_LO = 64
const m_R_M32R_GOTPC24 = 55
const m_R_M32R_GOTPC_HI_SLO = 60
const m_R_M32R_GOTPC_HI_ULO = 59
const m_R_M32R_GOTPC_LO = 61
const m_R_M32R_HI16_SLO = 8
const m_R_M32R_HI16_SLO_RELA = 40
const m_R_M32R_HI16_ULO = 7
const m_R_M32R_HI16_ULO_RELA = 39
const m_R_M32R_JMP_SLOT = 52
const m_R_M32R_LO16 = 9
const m_R_M32R_LO16_RELA = 41
const m_R_M32R_NONE = 0
const m_R_M32R_NUM = 256
const m_R_M32R_REL32 = 45
const m_R_M32R_RELATIVE = 53
const m_R_M32R_RELA_GNU_VTENTRY = 44
const m_R_M32R_RELA_GNU_VTINHERIT = 43
const m_R_M32R_SDA16 = 10
const m_R_M32R_SDA16_RELA = 42
const m_R_MICROBLAZE_32 = 1
const m_R_MICROBLAZE_32_LO = 6
const m_R_MICROBLAZE_32_PCREL = 2
const m_R_MICROBLAZE_32_PCREL_LO = 4
const m_R_MICROBLAZE_32_SYM_OP_SYM = 10
const m_R_MICROBLAZE_64 = 5
const m_R_MICROBLAZE_64_NONE = 9
const m_R_MICROBLAZE_64_PCREL = 3
const m_R_MICROBLAZE_COPY = 21
const m_R_MICROBLAZE_GLOB_DAT = 18
const m_R_MICROBLAZE_GNU_VTENTRY = 12
const m_R_MICROBLAZE_GNU_VTINHERIT = 11
const m_R_MICROBLAZE_GOTOFF_32 = 20
const m_R_MICROBLAZE_GOTOFF_64 = 19
const m_R_MICROBLAZE_GOTPC_64 = 13
const m_R_MICROBLAZE_GOT_64 = 14
const m_R_MICROBLAZE_JUMP_SLOT = 17
const m_R_MICROBLAZE_NONE = 0
const m_R_MICROBLAZE_PLT_64 = 15
const m_R_MICROBLAZE_REL = 16
const m_R_MICROBLAZE_SRO32 = 7
const m_R_MICROBLAZE_SRW32 = 8
const m_R_MICROBLAZE_TLS = 22
const m_R_MICROBLAZE_TLSDTPMOD32 = 25
const m_R_MICROBLAZE_TLSDTPREL32 = 26
const m_R_MICROBLAZE_TLSDTPREL64 = 27
const m_R_MICROBLAZE_TLSGD = 23
const m_R_MICROBLAZE_TLSGOTTPREL32 = 28
const m_R_MICROBLAZE_TLSLD = 24
const m_R_MICROBLAZE_TLSTPREL32 = 29
const m_R_MIPS_16 = 1
const m_R_MIPS_26 = 4
const m_R_MIPS_32 = 2
const m_R_MIPS_64 = 18
const m_R_MIPS_ADD_IMMEDIATE = 34
const m_R_MIPS_CALL16 = 11
const m_R_MIPS_CALL_HI16 = 30
const m_R_MIPS_CALL_LO16 = 31
const m_R_MIPS_COPY = 126
const m_R_MIPS_DELETE = 27
const m_R_MIPS_GLOB_DAT = 51
const m_R_MIPS_GOT16 = 9
const m_R_MIPS_GOT_DISP = 19
const m_R_MIPS_GOT_HI16 = 22
const m_R_MIPS_GOT_LO16 = 23
const m_R_MIPS_GOT_OFST = 21
const m_R_MIPS_GOT_PAGE = 20
const m_R_MIPS_GPREL16 = 7
const m_R_MIPS_GPREL32 = 12
const m_R_MIPS_HI16 = 5
const m_R_MIPS_HIGHER = 28
const m_R_MIPS_HIGHEST = 29
const m_R_MIPS_INSERT_A = 25
const m_R_MIPS_INSERT_B = 26
const m_R_MIPS_JALR = 37
const m_R_MIPS_JUMP_SLOT = 127
const m_R_MIPS_LITERAL = 8
const m_R_MIPS_LO16 = 6
const m_R_MIPS_NONE = 0
const m_R_MIPS_NUM = 128
const m_R_MIPS_PC16 = 10
const m_R_MIPS_PJUMP = 35
const m_R_MIPS_REL16 = 33
const m_R_MIPS_REL32 = 3
const m_R_MIPS_RELGOT = 36
const m_R_MIPS_SCN_DISP = 32
const m_R_MIPS_SHIFT5 = 16
const m_R_MIPS_SHIFT6 = 17
const m_R_MIPS_SUB = 24
const m_R_MIPS_TLS_DTPMOD32 = 38
const m_R_MIPS_TLS_DTPMOD64 = 40
const m_R_MIPS_TLS_DTPREL32 = 39
const m_R_MIPS_TLS_DTPREL64 = 41
const m_R_MIPS_TLS_DTPREL_HI16 = 44
const m_R_MIPS_TLS_DTPREL_LO16 = 45
const m_R_MIPS_TLS_GD = 42
const m_R_MIPS_TLS_GOTTPREL = 46
const m_R_MIPS_TLS_LDM = 43
const m_R_MIPS_TLS_TPREL32 = 47
const m_R_MIPS_TLS_TPREL64 = 48
const m_R_MIPS_TLS_TPREL_HI16 = 49
const m_R_MIPS_TLS_TPREL_LO16 = 50
const m_R_MN10300_16 = 2
const m_R_MN10300_24 = 9
const m_R_MN10300_32 = 1
const m_R_MN10300_8 = 3
const m_R_MN10300_COPY = 20
const m_R_MN10300_GLOB_DAT = 21
const m_R_MN10300_GNU_VTENTRY = 8
const m_R_MN10300_GNU_VTINHERIT = 7
const m_R_MN10300_GOT16 = 19
const m_R_MN10300_GOT24 = 18
const m_R_MN10300_GOT32 = 17
const m_R_MN10300_GOTOFF16 = 14
const m_R_MN10300_GOTOFF24 = 13
const m_R_MN10300_GOTOFF32 = 12
const m_R_MN10300_GOTPC16 = 11
const m_R_MN10300_GOTPC32 = 10
const m_R_MN10300_JMP_SLOT = 22
const m_R_MN10300_NONE = 0
const m_R_MN10300_NUM = 24
const m_R_MN10300_PCREL16 = 5
const m_R_MN10300_PCREL32 = 4
const m_R_MN10300_PCREL8 = 6
const m_R_MN10300_PLT16 = 16
const m_R_MN10300_PLT32 = 15
const m_R_MN10300_RELATIVE = 23
const m_R_NIOS2_ALIGN = 21
const m_R_NIOS2_BFD_RELOC_16 = 13
const m_R_NIOS2_BFD_RELOC_32 = 12
const m_R_NIOS2_BFD_RELOC_8 = 14
const m_R_NIOS2_CACHE_OPX = 6
const m_R_NIOS2_CALL16 = 23
const m_R_NIOS2_CALL26 = 4
const m_R_NIOS2_CALL26_NOAT = 41
const m_R_NIOS2_CALLR = 20
const m_R_NIOS2_CALL_HA = 45
const m_R_NIOS2_CALL_LO = 44
const m_R_NIOS2_CJMP = 19
const m_R_NIOS2_COPY = 36
const m_R_NIOS2_GLOB_DAT = 37
const m_R_NIOS2_GNU_VTENTRY = 17
const m_R_NIOS2_GNU_VTINHERIT = 16
const m_R_NIOS2_GOT16 = 22
const m_R_NIOS2_GOTOFF = 40
const m_R_NIOS2_GOTOFF_HA = 25
const m_R_NIOS2_GOTOFF_LO = 24
const m_R_NIOS2_GOT_HA = 43
const m_R_NIOS2_GOT_LO = 42
const m_R_NIOS2_GPREL = 15
const m_R_NIOS2_HI16 = 9
const m_R_NIOS2_HIADJ16 = 11
const m_R_NIOS2_IMM5 = 5
const m_R_NIOS2_IMM6 = 7
const m_R_NIOS2_IMM8 = 8
const m_R_NIOS2_JUMP_SLOT = 38
const m_R_NIOS2_LO16 = 10
const m_R_NIOS2_NONE = 0
const m_R_NIOS2_PCREL16 = 3
const m_R_NIOS2_PCREL_HA = 27
const m_R_NIOS2_PCREL_LO = 26
const m_R_NIOS2_RELATIVE = 39
const m_R_NIOS2_S16 = 1
const m_R_NIOS2_TLS_DTPMOD = 33
const m_R_NIOS2_TLS_DTPREL = 34
const m_R_NIOS2_TLS_GD16 = 28
const m_R_NIOS2_TLS_IE16 = 31
const m_R_NIOS2_TLS_LDM16 = 29
const m_R_NIOS2_TLS_LDO16 = 30
const m_R_NIOS2_TLS_LE16 = 32
const m_R_NIOS2_TLS_TPREL = 35
const m_R_NIOS2_U16 = 2
const m_R_NIOS2_UJMP = 18
const m_R_OR1K_16 = 2
const m_R_OR1K_16_PCREL = 10
const m_R_OR1K_32 = 1
const m_R_OR1K_32_PCREL = 9
const m_R_OR1K_8 = 3
const m_R_OR1K_8_PCREL = 11
const m_R_OR1K_COPY = 18
const m_R_OR1K_GLOB_DAT = 19
const m_R_OR1K_GNU_VTENTRY = 7
const m_R_OR1K_GNU_VTINHERIT = 8
const m_R_OR1K_GOT16 = 14
const m_R_OR1K_GOTOFF_HI16 = 16
const m_R_OR1K_GOTOFF_LO16 = 17
const m_R_OR1K_GOTPC_HI16 = 12
const m_R_OR1K_GOTPC_LO16 = 13
const m_R_OR1K_HI_16_IN_INSN = 5
const m_R_OR1K_INSN_REL_26 = 6
const m_R_OR1K_JMP_SLOT = 20
const m_R_OR1K_LO_16_IN_INSN = 4
const m_R_OR1K_NONE = 0
const m_R_OR1K_PLT26 = 15
const m_R_OR1K_RELATIVE = 21
const m_R_OR1K_TLS_DTPMOD = 34
const m_R_OR1K_TLS_DTPOFF = 33
const m_R_OR1K_TLS_GD_HI16 = 22
const m_R_OR1K_TLS_GD_LO16 = 23
const m_R_OR1K_TLS_IE_HI16 = 28
const m_R_OR1K_TLS_IE_LO16 = 29
const m_R_OR1K_TLS_LDM_HI16 = 24
const m_R_OR1K_TLS_LDM_LO16 = 25
const m_R_OR1K_TLS_LDO_HI16 = 26
const m_R_OR1K_TLS_LDO_LO16 = 27
const m_R_OR1K_TLS_LE_HI16 = 30
const m_R_OR1K_TLS_LE_LO16 = 31
const m_R_OR1K_TLS_TPOFF = 32
const m_R_PARISC_COPY = 128
const m_R_PARISC_DIR14DR = 84
const m_R_PARISC_DIR14R = 6
const m_R_PARISC_DIR14WR = 83
const m_R_PARISC_DIR16DF = 87
const m_R_PARISC_DIR16F = 85
const m_R_PARISC_DIR16WF = 86
const m_R_PARISC_DIR17F = 4
const m_R_PARISC_DIR17R = 3
const m_R_PARISC_DIR21L = 2
const m_R_PARISC_DIR32 = 1
const m_R_PARISC_DIR64 = 80
const m_R_PARISC_DPREL14R = 22
const m_R_PARISC_DPREL21L = 18
const m_R_PARISC_EPLT = 130
const m_R_PARISC_FPTR64 = 64
const m_R_PARISC_GNU_VTENTRY = 232
const m_R_PARISC_GNU_VTINHERIT = 233
const m_R_PARISC_GPREL14DR = 92
const m_R_PARISC_GPREL14R = 30
const m_R_PARISC_GPREL14WR = 91
const m_R_PARISC_GPREL16DF = 95
const m_R_PARISC_GPREL16F = 93
const m_R_PARISC_GPREL16WF = 94
const m_R_PARISC_GPREL21L = 26
const m_R_PARISC_GPREL64 = 88
const m_R_PARISC_HIRESERVE = 255
const m_R_PARISC_IPLT = 129
const m_R_PARISC_LORESERVE = 128
const m_R_PARISC_LTOFF14DR = 100
const m_R_PARISC_LTOFF14R = 38
const m_R_PARISC_LTOFF14WR = 99
const m_R_PARISC_LTOFF16DF = 103
const m_R_PARISC_LTOFF16F = 101
const m_R_PARISC_LTOFF16WF = 102
const m_R_PARISC_LTOFF21L = 34
const m_R_PARISC_LTOFF64 = 96
const m_R_PARISC_LTOFF_FPTR14DR = 124
const m_R_PARISC_LTOFF_FPTR14R = 62
const m_R_PARISC_LTOFF_FPTR14WR = 123
const m_R_PARISC_LTOFF_FPTR16DF = 127
const m_R_PARISC_LTOFF_FPTR16F = 125
const m_R_PARISC_LTOFF_FPTR16WF = 126
const m_R_PARISC_LTOFF_FPTR21L = 58
const m_R_PARISC_LTOFF_FPTR32 = 57
const m_R_PARISC_LTOFF_FPTR64 = 120
const m_R_PARISC_LTOFF_TP14DR = 228
const m_R_PARISC_LTOFF_TP14F = 167
const m_R_PARISC_LTOFF_TP14R = 166
const m_R_PARISC_LTOFF_TP14WR = 227
const m_R_PARISC_LTOFF_TP16DF = 231
const m_R_PARISC_LTOFF_TP16F = 229
const m_R_PARISC_LTOFF_TP16WF = 230
const m_R_PARISC_LTOFF_TP21L = 162
const m_R_PARISC_LTOFF_TP64 = 224
const m_R_PARISC_NONE = 0
const m_R_PARISC_PCREL14DR = 76
const m_R_PARISC_PCREL14R = 14
const m_R_PARISC_PCREL14WR = 75
const m_R_PARISC_PCREL16DF = 79
const m_R_PARISC_PCREL16F = 77
const m_R_PARISC_PCREL16WF = 78
const m_R_PARISC_PCREL17F = 12
const m_R_PARISC_PCREL17R = 11
const m_R_PARISC_PCREL21L = 10
const m_R_PARISC_PCREL22F = 74
const m_R_PARISC_PCREL32 = 9
const m_R_PARISC_PCREL64 = 72
const m_R_PARISC_PLABEL14R = 70
const m_R_PARISC_PLABEL21L = 66
const m_R_PARISC_PLABEL32 = 65
const m_R_PARISC_PLTOFF14DR = 116
const m_R_PARISC_PLTOFF14R = 54
const m_R_PARISC_PLTOFF14WR = 115
const m_R_PARISC_PLTOFF16DF = 119
const m_R_PARISC_PLTOFF16F = 117
const m_R_PARISC_PLTOFF16WF = 118
const m_R_PARISC_PLTOFF21L = 50
const m_R_PARISC_SECREL32 = 41
const m_R_PARISC_SECREL64 = 104
const m_R_PARISC_SEGBASE = 48
const m_R_PARISC_SEGREL32 = 49
const m_R_PARISC_SEGREL64 = 112
const m_R_PARISC_TLS_DTPMOD32 = 242
const m_R_PARISC_TLS_DTPMOD64 = 243
const m_R_PARISC_TLS_DTPOFF32 = 244
const m_R_PARISC_TLS_DTPOFF64 = 245
const m_R_PARISC_TLS_GD14R = 235
const m_R_PARISC_TLS_GD21L = 234
const m_R_PARISC_TLS_GDCALL = 236
const m_R_PARISC_TLS_IE14R = "R_PARISC_LTOFF_TP14R"
const m_R_PARISC_TLS_IE21L = "R_PARISC_LTOFF_TP21L"
const m_R_PARISC_TLS_LDM14R = 238
const m_R_PARISC_TLS_LDM21L = 237
const m_R_PARISC_TLS_LDMCALL = 239
const m_R_PARISC_TLS_LDO14R = 241
const m_R_PARISC_TLS_LDO21L = 240
const m_R_PARISC_TLS_LE14R = "R_PARISC_TPREL14R"
const m_R_PARISC_TLS_LE21L = "R_PARISC_TPREL21L"
const m_R_PARISC_TLS_TPREL32 = "R_PARISC_TPREL32"
const m_R_PARISC_TLS_TPREL64 = "R_PARISC_TPREL64"
const m_R_PARISC_TPREL14DR = 220
const m_R_PARISC_TPREL14R = 158
const m_R_PARISC_TPREL14WR = 219
const m_R_PARISC_TPREL16DF = 223
const m_R_PARISC_TPREL16F = 221
const m_R_PARISC_TPREL16WF = 222
const m_R_PARISC_TPREL21L = 154
const m_R_PARISC_TPREL32 = 153
const m_R_PARISC_TPREL64 = 216
const m_R_PPC64_ADDR14 = "R_PPC_ADDR14"
const m_R_PPC64_ADDR14_BRNTAKEN = "R_PPC_ADDR14_BRNTAKEN"
const m_R_PPC64_ADDR14_BRTAKEN = "R_PPC_ADDR14_BRTAKEN"
const m_R_PPC64_ADDR16 = "R_PPC_ADDR16"
const m_R_PPC64_ADDR16_DS = 56
const m_R_PPC64_ADDR16_HA = "R_PPC_ADDR16_HA"
const m_R_PPC64_ADDR16_HI = "R_PPC_ADDR16_HI"
const m_R_PPC64_ADDR16_HIGH = 110
const m_R_PPC64_ADDR16_HIGHA = 111
const m_R_PPC64_ADDR16_HIGHER = 39
const m_R_PPC64_ADDR16_HIGHERA = 40
const m_R_PPC64_ADDR16_HIGHEST = 41
const m_R_PPC64_ADDR16_HIGHESTA = 42
const m_R_PPC64_ADDR16_LO = "R_PPC_ADDR16_LO"
const m_R_PPC64_ADDR16_LO_DS = 57
const m_R_PPC64_ADDR24 = "R_PPC_ADDR24"
const m_R_PPC64_ADDR30 = 37
const m_R_PPC64_ADDR32 = "R_PPC_ADDR32"
const m_R_PPC64_ADDR64 = 38
const m_R_PPC64_COPY = "R_PPC_COPY"
const m_R_PPC64_DTPMOD64 = 68
const m_R_PPC64_DTPREL16 = 74
const m_R_PPC64_DTPREL16_DS = 101
const m_R_PPC64_DTPREL16_HA = 77
const m_R_PPC64_DTPREL16_HI = 76
const m_R_PPC64_DTPREL16_HIGH = 114
const m_R_PPC64_DTPREL16_HIGHA = 115
const m_R_PPC64_DTPREL16_HIGHER = 103
const m_R_PPC64_DTPREL16_HIGHERA = 104
const m_R_PPC64_DTPREL16_HIGHEST = 105
const m_R_PPC64_DTPREL16_HIGHESTA = 106
const m_R_PPC64_DTPREL16_LO = 75
const m_R_PPC64_DTPREL16_LO_DS = 102
const m_R_PPC64_DTPREL64 = 78
const m_R_PPC64_GLOB_DAT = "R_PPC_GLOB_DAT"
const m_R_PPC64_GOT16 = "R_PPC_GOT16"
const m_R_PPC64_GOT16_DS = 58
const m_R_PPC64_GOT16_HA = "R_PPC_GOT16_HA"
const m_R_PPC64_GOT16_HI = "R_PPC_GOT16_HI"
const m_R_PPC64_GOT16_LO = "R_PPC_GOT16_LO"
const m_R_PPC64_GOT16_LO_DS = 59
const m_R_PPC64_GOT_DTPREL16_DS = 91
const m_R_PPC64_GOT_DTPREL16_HA = 94
const m_R_PPC64_GOT_DTPREL16_HI = 93
const m_R_PPC64_GOT_DTPREL16_LO_DS = 92
const m_R_PPC64_GOT_TLSGD16 = 79
const m_R_PPC64_GOT_TLSGD16_HA = 82
const m_R_PPC64_GOT_TLSGD16_HI = 81
const m_R_PPC64_GOT_TLSGD16_LO = 80
const m_R_PPC64_GOT_TLSLD16 = 83
const m_R_PPC64_GOT_TLSLD16_HA = 86
const m_R_PPC64_GOT_TLSLD16_HI = 85
const m_R_PPC64_GOT_TLSLD16_LO = 84
const m_R_PPC64_GOT_TPREL16_DS = 87
const m_R_PPC64_GOT_TPREL16_HA = 90
const m_R_PPC64_GOT_TPREL16_HI = 89
const m_R_PPC64_GOT_TPREL16_LO_DS = 88
const m_R_PPC64_IRELATIVE = 248
const m_R_PPC64_JMP_IREL = 247
const m_R_PPC64_JMP_SLOT = "R_PPC_JMP_SLOT"
const m_R_PPC64_NONE = "R_PPC_NONE"
const m_R_PPC64_PLT16_HA = "R_PPC_PLT16_HA"
const m_R_PPC64_PLT16_HI = "R_PPC_PLT16_HI"
const m_R_PPC64_PLT16_LO = "R_PPC_PLT16_LO"
const m_R_PPC64_PLT16_LO_DS = 60
const m_R_PPC64_PLT32 = "R_PPC_PLT32"
const m_R_PPC64_PLT64 = 45
const m_R_PPC64_PLTGOT16 = 52
const m_R_PPC64_PLTGOT16_DS = 65
const m_R_PPC64_PLTGOT16_HA = 55
const m_R_PPC64_PLTGOT16_HI = 54
const m_R_PPC64_PLTGOT16_LO = 53
const m_R_PPC64_PLTGOT16_LO_DS = 66
const m_R_PPC64_PLTREL32 = "R_PPC_PLTREL32"
const m_R_PPC64_PLTREL64 = 46
const m_R_PPC64_REL14 = "R_PPC_REL14"
const m_R_PPC64_REL14_BRNTAKEN = "R_PPC_REL14_BRNTAKEN"
const m_R_PPC64_REL14_BRTAKEN = "R_PPC_REL14_BRTAKEN"
const m_R_PPC64_REL16 = 249
const m_R_PPC64_REL16_HA = 252
const m_R_PPC64_REL16_HI = 251
const m_R_PPC64_REL16_LO = 250
const m_R_PPC64_REL24 = "R_PPC_REL24"
const m_R_PPC64_REL32 = "R_PPC_REL32"
const m_R_PPC64_REL64 = 44
const m_R_PPC64_RELATIVE = "R_PPC_RELATIVE"
const m_R_PPC64_SECTOFF = "R_PPC_SECTOFF"
const m_R_PPC64_SECTOFF_DS = 61
const m_R_PPC64_SECTOFF_HA = "R_PPC_SECTOFF_HA"
const m_R_PPC64_SECTOFF_HI = "R_PPC_SECTOFF_HI"
const m_R_PPC64_SECTOFF_LO = "R_PPC_SECTOFF_LO"
const m_R_PPC64_SECTOFF_LO_DS = 62
const m_R_PPC64_TLS = 67
const m_R_PPC64_TLSGD = 107
const m_R_PPC64_TLSLD = 108
const m_R_PPC64_TOC = 51
const m_R_PPC64_TOC16 = 47
const m_R_PPC64_TOC16_DS = 63
const m_R_PPC64_TOC16_HA = 50
const m_R_PPC64_TOC16_HI = 49
const m_R_PPC64_TOC16_LO = 48
const m_R_PPC64_TOC16_LO_DS = 64
const m_R_PPC64_TOCSAVE = 109
const m_R_PPC64_TPREL16 = 69
const m_R_PPC64_TPREL16_DS = 95
const m_R_PPC64_TPREL16_HA = 72
const m_R_PPC64_TPREL16_HI = 71
const m_R_PPC64_TPREL16_HIGH = 112
const m_R_PPC64_TPREL16_HIGHA = 113
const m_R_PPC64_TPREL16_HIGHER = 97
const m_R_PPC64_TPREL16_HIGHERA = 98
const m_R_PPC64_TPREL16_HIGHEST = 99
const m_R_PPC64_TPREL16_HIGHESTA = 100
const m_R_PPC64_TPREL16_LO = 70
const m_R_PPC64_TPREL16_LO_DS = 96
const m_R_PPC64_TPREL64 = 73
const m_R_PPC64_UADDR16 = "R_PPC_UADDR16"
const m_R_PPC64_UADDR32 = "R_PPC_UADDR32"
const m_R_PPC64_UADDR64 = 43
const m_R_PPC_ADDR14 = 7
const m_R_PPC_ADDR14_BRNTAKEN = 9
const m_R_PPC_ADDR14_BRTAKEN = 8
const m_R_PPC_ADDR16 = 3
const m_R_PPC_ADDR16_HA = 6
const m_R_PPC_ADDR16_HI = 5
const m_R_PPC_ADDR16_LO = 4
const m_R_PPC_ADDR24 = 2
const m_R_PPC_ADDR32 = 1
const m_R_PPC_COPY = 19
const m_R_PPC_DIAB_RELSDA_HA = 185
const m_R_PPC_DIAB_RELSDA_HI = 184
const m_R_PPC_DIAB_RELSDA_LO = 183
const m_R_PPC_DIAB_SDA21_HA = 182
const m_R_PPC_DIAB_SDA21_HI = 181
const m_R_PPC_DIAB_SDA21_LO = 180
const m_R_PPC_DTPMOD32 = 68
const m_R_PPC_DTPREL16 = 74
const m_R_PPC_DTPREL16_HA = 77
const m_R_PPC_DTPREL16_HI = 76
const m_R_PPC_DTPREL16_LO = 75
const m_R_PPC_DTPREL32 = 78
const m_R_PPC_EMB_BIT_FLD = 115
const m_R_PPC_EMB_MRKREF = 110
const m_R_PPC_EMB_NADDR16 = 102
const m_R_PPC_EMB_NADDR16_HA = 105
const m_R_PPC_EMB_NADDR16_HI = 104
const m_R_PPC_EMB_NADDR16_LO = 103
const m_R_PPC_EMB_NADDR32 = 101
const m_R_PPC_EMB_RELSDA = 116
const m_R_PPC_EMB_RELSEC16 = 111
const m_R_PPC_EMB_RELST_HA = 114
const m_R_PPC_EMB_RELST_HI = 113
const m_R_PPC_EMB_RELST_LO = 112
const m_R_PPC_EMB_SDA21 = 109
const m_R_PPC_EMB_SDA2I16 = 107
const m_R_PPC_EMB_SDA2REL = 108
const m_R_PPC_EMB_SDAI16 = 106
const m_R_PPC_GLOB_DAT = 20
const m_R_PPC_GOT16 = 14
const m_R_PPC_GOT16_HA = 17
const m_R_PPC_GOT16_HI = 16
const m_R_PPC_GOT16_LO = 15
const m_R_PPC_GOT_DTPREL16 = 91
const m_R_PPC_GOT_DTPREL16_HA = 94
const m_R_PPC_GOT_DTPREL16_HI = 93
const m_R_PPC_GOT_DTPREL16_LO = 92
const m_R_PPC_GOT_TLSGD16 = 79
const m_R_PPC_GOT_TLSGD16_HA = 82
const m_R_PPC_GOT_TLSGD16_HI = 81
const m_R_PPC_GOT_TLSGD16_LO = 80
const m_R_PPC_GOT_TLSLD16 = 83
const m_R_PPC_GOT_TLSLD16_HA = 86
const m_R_PPC_GOT_TLSLD16_HI = 85
const m_R_PPC_GOT_TLSLD16_LO = 84
const m_R_PPC_GOT_TPREL16 = 87
const m_R_PPC_GOT_TPREL16_HA = 90
const m_R_PPC_GOT_TPREL16_HI = 89
const m_R_PPC_GOT_TPREL16_LO = 88
const m_R_PPC_IRELATIVE = 248
const m_R_PPC_JMP_SLOT = 21
const m_R_PPC_LOCAL24PC = 23
const m_R_PPC_NONE = 0
const m_R_PPC_PLT16_HA = 31
const m_R_PPC_PLT16_HI = 30
const m_R_PPC_PLT16_LO = 29
const m_R_PPC_PLT32 = 27
const m_R_PPC_PLTREL24 = 18
const m_R_PPC_PLTREL32 = 28
const m_R_PPC_REL14 = 11
const m_R_PPC_REL14_BRNTAKEN = 13
const m_R_PPC_REL14_BRTAKEN = 12
const m_R_PPC_REL16 = 249
const m_R_PPC_REL16_HA = 252
const m_R_PPC_REL16_HI = 251
const m_R_PPC_REL16_LO = 250
const m_R_PPC_REL24 = 10
const m_R_PPC_REL32 = 26
const m_R_PPC_RELATIVE = 22
const m_R_PPC_SDAREL16 = 32
const m_R_PPC_SECTOFF = 33
const m_R_PPC_SECTOFF_HA = 36
const m_R_PPC_SECTOFF_HI = 35
const m_R_PPC_SECTOFF_LO = 34
const m_R_PPC_TLS = 67
const m_R_PPC_TLSGD = 95
const m_R_PPC_TLSLD = 96
const m_R_PPC_TOC16 = 255
const m_R_PPC_TPREL16 = 69
const m_R_PPC_TPREL16_HA = 72
const m_R_PPC_TPREL16_HI = 71
const m_R_PPC_TPREL16_LO = 70
const m_R_PPC_TPREL32 = 73
const m_R_PPC_UADDR16 = 25
const m_R_PPC_UADDR32 = 24
const m_R_RISCV_32 = 1
const m_R_RISCV_32_PCREL = 57
const m_R_RISCV_64 = 2
const m_R_RISCV_ADD16 = 34
const m_R_RISCV_ADD32 = 35
const m_R_RISCV_ADD64 = 36
const m_R_RISCV_ADD8 = 33
const m_R_RISCV_ALIGN = 43
const m_R_RISCV_BRANCH = 16
const m_R_RISCV_CALL = 18
const m_R_RISCV_CALL_PLT = 19
const m_R_RISCV_COPY = 4
const m_R_RISCV_GOT32_PCREL = 41
const m_R_RISCV_GOT_HI20 = 20
const m_R_RISCV_HI20 = 26
const m_R_RISCV_IRELATIVE = 58
const m_R_RISCV_JAL = 17
const m_R_RISCV_JUMP_SLOT = 5
const m_R_RISCV_LO12_I = 27
const m_R_RISCV_LO12_S = 28
const m_R_RISCV_NONE = 0
const m_R_RISCV_PCREL_HI20 = 23
const m_R_RISCV_PCREL_LO12_I = 24
const m_R_RISCV_PCREL_LO12_S = 25
const m_R_RISCV_PLT32 = 59
const m_R_RISCV_RELATIVE = 3
const m_R_RISCV_RELAX = 51
const m_R_RISCV_RVC_BRANCH = 44
const m_R_RISCV_RVC_JUMP = 45
const m_R_RISCV_RVC_LUI = 46
const m_R_RISCV_SET16 = 55
const m_R_RISCV_SET32 = 56
const m_R_RISCV_SET6 = 53
const m_R_RISCV_SET8 = 54
const m_R_RISCV_SET_ULEB128 = 60
const m_R_RISCV_SUB16 = 38
const m_R_RISCV_SUB32 = 39
const m_R_RISCV_SUB6 = 52
const m_R_RISCV_SUB64 = 40
const m_R_RISCV_SUB8 = 37
const m_R_RISCV_SUB_ULEB128 = 61
const m_R_RISCV_TLSDESC = 12
const m_R_RISCV_TLSDESC_ADD_LO12 = 64
const m_R_RISCV_TLSDESC_CALL = 65
const m_R_RISCV_TLSDESC_HI20 = 62
const m_R_RISCV_TLSDESC_LOAD_LO12 = 63
const m_R_RISCV_TLS_DTPMOD32 = 6
const m_R_RISCV_TLS_DTPMOD64 = 7
const m_R_RISCV_TLS_DTPREL32 = 8
const m_R_RISCV_TLS_DTPREL64 = 9
const m_R_RISCV_TLS_GD_HI20 = 22
const m_R_RISCV_TLS_GOT_HI20 = 21
const m_R_RISCV_TLS_TPREL32 = 10
const m_R_RISCV_TLS_TPREL64 = 11
const m_R_RISCV_TPREL_ADD = 32
const m_R_RISCV_TPREL_HI20 = 29
const m_R_RISCV_TPREL_LO12_I = 30
const m_R_RISCV_TPREL_LO12_S = 31
const m_R_SH_ALIGN = 29
const m_R_SH_CODE = 30
const m_R_SH_COPY = 162
const m_R_SH_COUNT = 28
const m_R_SH_DATA = 31
const m_R_SH_DIR32 = 1
const m_R_SH_DIR8BP = 7
const m_R_SH_DIR8L = 9
const m_R_SH_DIR8W = 8
const m_R_SH_DIR8WPL = 5
const m_R_SH_DIR8WPN = 3
const m_R_SH_DIR8WPZ = 6
const m_R_SH_FUNCDESC = 207
const m_R_SH_FUNCDESC_VALUE = 208
const m_R_SH_GLOB_DAT = 163
const m_R_SH_GNU_VTENTRY = 35
const m_R_SH_GNU_VTINHERIT = 34
const m_R_SH_GOT20 = 201
const m_R_SH_GOT32 = 160
const m_R_SH_GOTFUNCDESC = 203
const m_R_SH_GOTFUNCDEST20 = 204
const m_R_SH_GOTOFF = 166
const m_R_SH_GOTOFF20 = 202
const m_R_SH_GOTOFFFUNCDESC = 205
const m_R_SH_GOTOFFFUNCDEST20 = 206
const m_R_SH_GOTPC = 167
const m_R_SH_IND12W = 4
const m_R_SH_JMP_SLOT = 164
const m_R_SH_LABEL = 32
const m_R_SH_NONE = 0
const m_R_SH_NUM = 256
const m_R_SH_PLT32 = 161
const m_R_SH_REL32 = 2
const m_R_SH_RELATIVE = 165
const m_R_SH_SWITCH16 = 25
const m_R_SH_SWITCH32 = 26
const m_R_SH_SWITCH8 = 33
const m_R_SH_TLS_DTPMOD32 = 149
const m_R_SH_TLS_DTPOFF32 = 150
const m_R_SH_TLS_GD_32 = 144
const m_R_SH_TLS_IE_32 = 147
const m_R_SH_TLS_LDO_32 = 146
const m_R_SH_TLS_LD_32 = 145
const m_R_SH_TLS_LE_32 = 148
const m_R_SH_TLS_TPOFF32 = 151
const m_R_SH_USES = 27
const m_R_SPARC_10 = 30
const m_R_SPARC_11 = 31
const m_R_SPARC_13 = 11
const m_R_SPARC_16 = 2
const m_R_SPARC_22 = 10
const m_R_SPARC_32 = 3
const m_R_SPARC_5 = 44
const m_R_SPARC_6 = 45
const m_R_SPARC_64 = 32
const m_R_SPARC_7 = 43
const m_R_SPARC_8 = 1
const m_R_SPARC_COPY = 19
const m_R_SPARC_DISP16 = 5
const m_R_SPARC_DISP32 = 6
const m_R_SPARC_DISP64 = 46
const m_R_SPARC_DISP8 = 4
const m_R_SPARC_GLOB_DAT = 20
const m_R_SPARC_GLOB_JMP = 42
const m_R_SPARC_GNU_VTENTRY = 251
const m_R_SPARC_GNU_VTINHERIT = 250
const m_R_SPARC_GOT10 = 13
const m_R_SPARC_GOT13 = 14
const m_R_SPARC_GOT22 = 15
const m_R_SPARC_GOTDATA_HIX22 = 80
const m_R_SPARC_GOTDATA_LOX10 = 81
const m_R_SPARC_GOTDATA_OP = 84
const m_R_SPARC_GOTDATA_OP_HIX22 = 82
const m_R_SPARC_GOTDATA_OP_LOX10 = 83
const m_R_SPARC_H34 = 85
const m_R_SPARC_H44 = 50
const m_R_SPARC_HH22 = 34
const m_R_SPARC_HI22 = 9
const m_R_SPARC_HIPLT22 = 25
const m_R_SPARC_HIX22 = 48
const m_R_SPARC_HM10 = 35
const m_R_SPARC_JMP_SLOT = 21
const m_R_SPARC_L44 = 52
const m_R_SPARC_LM22 = 36
const m_R_SPARC_LO10 = 12
const m_R_SPARC_LOPLT10 = 26
const m_R_SPARC_LOX10 = 49
const m_R_SPARC_M44 = 51
const m_R_SPARC_NONE = 0
const m_R_SPARC_NUM = 253
const m_R_SPARC_OLO10 = 33
const m_R_SPARC_PC10 = 16
const m_R_SPARC_PC22 = 17
const m_R_SPARC_PCPLT10 = 29
const m_R_SPARC_PCPLT22 = 28
const m_R_SPARC_PCPLT32 = 27
const m_R_SPARC_PC_HH22 = 37
const m_R_SPARC_PC_HM10 = 38
const m_R_SPARC_PC_LM22 = 39
const m_R_SPARC_PLT32 = 24
const m_R_SPARC_PLT64 = 47
const m_R_SPARC_REGISTER = 53
const m_R_SPARC_RELATIVE = 22
const m_R_SPARC_REV32 = 252
const m_R_SPARC_SIZE32 = 86
const m_R_SPARC_SIZE64 = 87
const m_R_SPARC_TLS_DTPMOD32 = 74
const m_R_SPARC_TLS_DTPMOD64 = 75
const m_R_SPARC_TLS_DTPOFF32 = 76
const m_R_SPARC_TLS_DTPOFF64 = 77
const m_R_SPARC_TLS_GD_ADD = 58
const m_R_SPARC_TLS_GD_CALL = 59
const m_R_SPARC_TLS_GD_HI22 = 56
const m_R_SPARC_TLS_GD_LO10 = 57
const m_R_SPARC_TLS_IE_ADD = 71
const m_R_SPARC_TLS_IE_HI22 = 67
const m_R_SPARC_TLS_IE_LD = 69
const m_R_SPARC_TLS_IE_LDX = 70
const m_R_SPARC_TLS_IE_LO10 = 68
const m_R_SPARC_TLS_LDM_ADD = 62
const m_R_SPARC_TLS_LDM_CALL = 63
const m_R_SPARC_TLS_LDM_HI22 = 60
const m_R_SPARC_TLS_LDM_LO10 = 61
const m_R_SPARC_TLS_LDO_ADD = 66
const m_R_SPARC_TLS_LDO_HIX22 = 64
const m_R_SPARC_TLS_LDO_LOX10 = 65
const m_R_SPARC_TLS_LE_HIX22 = 72
const m_R_SPARC_TLS_LE_LOX10 = 73
const m_R_SPARC_TLS_TPOFF32 = 78
const m_R_SPARC_TLS_TPOFF64 = 79
const m_R_SPARC_UA16 = 55
const m_R_SPARC_UA32 = 23
const m_R_SPARC_UA64 = 54
const m_R_SPARC_WDISP16 = 40
const m_R_SPARC_WDISP19 = 41
const m_R_SPARC_WDISP22 = 8
const m_R_SPARC_WDISP30 = 7
const m_R_SPARC_WPLT30 = 18
const m_R_X86_64_16 = 12
const m_R_X86_64_32 = 10
const m_R_X86_64_32S = 11
const m_R_X86_64_64 = 1
const m_R_X86_64_8 = 14
const m_R_X86_64_COPY = 5
const m_R_X86_64_DTPMOD64 = 16
const m_R_X86_64_DTPOFF32 = 21
const m_R_X86_64_DTPOFF64 = 17
const m_R_X86_64_GLOB_DAT = 6
const m_R_X86_64_GOT32 = 3
const m_R_X86_64_GOT64 = 27
const m_R_X86_64_GOTOFF64 = 25
const m_R_X86_64_GOTPC32 = 26
const m_R_X86_64_GOTPC32_TLSDESC = 34
const m_R_X86_64_GOTPC64 = 29
const m_R_X86_64_GOTPCREL = 9
const m_R_X86_64_GOTPCREL64 = 28
const m_R_X86_64_GOTPCRELX = 41
const m_R_X86_64_GOTPLT64 = 30
const m_R_X86_64_GOTTPOFF = 22
const m_R_X86_64_IRELATIVE = 37
const m_R_X86_64_JUMP_SLOT = 7
const m_R_X86_64_NONE = 0
const m_R_X86_64_NUM = 43
const m_R_X86_64_PC16 = 13
const m_R_X86_64_PC32 = 2
const m_R_X86_64_PC64 = 24
const m_R_X86_64_PC8 = 15
const m_R_X86_64_PLT32 = 4
const m_R_X86_64_PLTOFF64 = 31
const m_R_X86_64_RELATIVE = 8
const m_R_X86_64_RELATIVE64 = 38
const m_R_X86_64_REX_GOTPCRELX = 42
const m_R_X86_64_SIZE32 = 32
const m_R_X86_64_SIZE64 = 33
const m_R_X86_64_TLSDESC = 36
const m_R_X86_64_TLSDESC_CALL = 35
const m_R_X86_64_TLSGD = 19
const m_R_X86_64_TLSLD = 20
const m_R_X86_64_TPOFF32 = 23
const m_R_X86_64_TPOFF64 = 18
const m_SELFMAG = 4
const m_SHF_ALPHA_GPREL = 0x10000000
const m_SHF_ARM_COMDEF = 0x80000000
const m_SHF_ARM_ENTRYSECT = 0x10000000
const m_SHF_IA_64_NORECOV = 0x20000000
const m_SHF_IA_64_SHORT = 0x10000000
const m_SHF_MASKOS = 0x0ff00000
const m_SHF_MASKPROC = 0xf0000000
const m_SHF_MIPS_ADDR = 0x40000000
const m_SHF_MIPS_GPREL = 0x10000000
const m_SHF_MIPS_LOCAL = 0x04000000
const m_SHF_MIPS_MERGE = 0x20000000
const m_SHF_MIPS_NAMES = 0x02000000
const m_SHF_MIPS_NODUPE = 0x01000000
const m_SHF_MIPS_NOSTRIP = 0x08000000
const m_SHF_MIPS_STRINGS = 0x80000000
const m_SHF_PARISC_HUGE = 0x40000000
const m_SHF_PARISC_SBP = 0x80000000
const m_SHF_PARISC_SHORT = 0x20000000
const m_SHN_ABS = 65521
const m_SHN_AFTER = 0xff01
const m_SHN_BEFORE = 0xff00
const m_SHN_COMMON = 65522
const m_SHN_HIOS = 0xff3f
const m_SHN_HIPROC = 0xff1f
const m_SHN_HIRESERVE = 0xffff
const m_SHN_LOOS = 0xff20
const m_SHN_LOPROC = 0xff00
const m_SHN_LORESERVE = 0xff00
const m_SHN_MIPS_ACOMMON = 0xff00
const m_SHN_MIPS_DATA = 0xff02
const m_SHN_MIPS_SCOMMON = 0xff03
const m_SHN_MIPS_SUNDEFINED = 0xff04
const m_SHN_MIPS_TEXT = 0xff01
const m_SHN_PARISC_ANSI_COMMON = 0xff00
const m_SHN_PARISC_HUGE_COMMON = 0xff01
const m_SHN_UNDEF = 0
const m_SHN_XINDEX = 0xffff
const m_SHT_ALPHA_DEBUG = 0x70000001
const m_SHT_ALPHA_REGINFO = 0x70000002
const m_SHT_CHECKSUM = 0x6ffffff8
const m_SHT_DYNAMIC = 6
const m_SHT_DYNSYM = 11
const m_SHT_FINI_ARRAY = 15
const m_SHT_GNU_ATTRIBUTES = 0x6ffffff5
const m_SHT_GNU_HASH = 0x6ffffff6
const m_SHT_GNU_LIBLIST = 0x6ffffff7
const m_SHT_GNU_verdef = 0x6ffffffd
const m_SHT_GNU_verneed = 0x6ffffffe
const m_SHT_GNU_versym = 0x6fffffff
const m_SHT_GROUP = 17
const m_SHT_HASH = 5
const m_SHT_HIOS = 0x6fffffff
const m_SHT_HIPROC = 0x7fffffff
const m_SHT_HISUNW = 0x6fffffff
const m_SHT_HIUSER = 0x8fffffff
const m_SHT_INIT_ARRAY = 14
const m_SHT_LOOS = 0x60000000
const m_SHT_LOPROC = 0x70000000
const m_SHT_LOSUNW = 0x6ffffffa
const m_SHT_LOUSER = 0x80000000
const m_SHT_MIPS_AUXSYM = 0x70000016
const m_SHT_MIPS_CONFLICT = 0x70000002
const m_SHT_MIPS_CONTENT = 0x7000000c
const m_SHT_MIPS_DEBUG = 0x70000005
const m_SHT_MIPS_DELTACLASS = 0x7000001d
const m_SHT_MIPS_DELTADECL = 0x7000001f
const m_SHT_MIPS_DELTAINST = 0x7000001c
const m_SHT_MIPS_DELTASYM = 0x7000001b
const m_SHT_MIPS_DENSE = 0x70000013
const m_SHT_MIPS_DWARF = 0x7000001e
const m_SHT_MIPS_EH_REGION = 0x70000027
const m_SHT_MIPS_EVENTS = 0x70000021
const m_SHT_MIPS_EXTSYM = 0x70000012
const m_SHT_MIPS_FDESC = 0x70000011
const m_SHT_MIPS_GPTAB = 0x70000003
const m_SHT_MIPS_IFACE = 0x7000000b
const m_SHT_MIPS_LIBLIST = 0x70000000
const m_SHT_MIPS_LINE = 0x70000019
const m_SHT_MIPS_LOCSTR = 0x70000018
const m_SHT_MIPS_LOCSYM = 0x70000015
const m_SHT_MIPS_MSYM = 0x70000001
const m_SHT_MIPS_OPTIONS = 0x7000000d
const m_SHT_MIPS_OPTSYM = 0x70000017
const m_SHT_MIPS_PACKAGE = 0x70000007
const m_SHT_MIPS_PACKSYM = 0x70000008
const m_SHT_MIPS_PDESC = 0x70000014
const m_SHT_MIPS_PDR_EXCEPTION = 0x70000029
const m_SHT_MIPS_PIXIE = 0x70000023
const m_SHT_MIPS_REGINFO = 0x70000006
const m_SHT_MIPS_RELD = 0x70000009
const m_SHT_MIPS_RFDESC = 0x7000001a
const m_SHT_MIPS_SHDR = 0x70000010
const m_SHT_MIPS_SYMBOL_LIB = 0x70000020
const m_SHT_MIPS_TRANSLATE = 0x70000022
const m_SHT_MIPS_UCODE = 0x70000004
const m_SHT_MIPS_WHIRL = 0x70000026
const m_SHT_MIPS_XLATE = 0x70000024
const m_SHT_MIPS_XLATE_DEBUG = 0x70000025
const m_SHT_MIPS_XLATE_OLD = 0x70000028
const m_SHT_NOBITS = 8
const m_SHT_NOTE = 7
const m_SHT_NULL = 0
const m_SHT_NUM = 20
const m_SHT_PARISC_DOC = 0x70000002
const m_SHT_PARISC_EXT = 0x70000000
const m_SHT_PARISC_UNWIND = 0x70000001
const m_SHT_PREINIT_ARRAY = 16
const m_SHT_PROGBITS = 1
const m_SHT_REL = 9
const m_SHT_RELA = 4
const m_SHT_RELR = 19
const m_SHT_SHLIB = 10
const m_SHT_STRTAB = 3
const m_SHT_SUNW_COMDAT = 0x6ffffffb
const m_SHT_SUNW_move = 0x6ffffffa
const m_SHT_SUNW_syminfo = 0x6ffffffc
const m_SHT_SYMTAB = 2
const m_SHT_SYMTAB_SHNDX = 18
const m_STB_GLOBAL = 1
const m_STB_GNU_UNIQUE = 10
const m_STB_HIOS = 12
const m_STB_HIPROC = 15
const m_STB_LOCAL = 0
const m_STB_LOOS = 10
const m_STB_LOPROC = 13
const m_STB_MIPS_SPLIT_COMMON = 13
const m_STB_NUM = 3
const m_STB_WEAK = 2
const m_STN_UNDEF = 0
const m_STO_ALPHA_NOPV = 0x80
const m_STO_ALPHA_STD_GPLOAD = 0x88
const m_STO_MIPS_DEFAULT = 0x0
const m_STO_MIPS_HIDDEN = 0x2
const m_STO_MIPS_INTERNAL = 0x1
const m_STO_MIPS_PLT = 0x8
const m_STO_MIPS_PROTECTED = 0x3
const m_STO_MIPS_SC_ALIGN_UNUSED = 0xff
const m_STO_PPC64_LOCAL_BIT = 5
const m_STO_PPC64_LOCAL_MASK = 0xe0
const m_STT_ARM_16BIT = "STT_HIPROC"
const m_STT_ARM_TFUNC = "STT_LOPROC"
const m_STT_COMMON = 5
const m_STT_FILE = 4
const m_STT_FUNC = 2
const m_STT_GNU_IFUNC = 10
const m_STT_HIOS = 12
const m_STT_HIPROC = 15
const m_STT_LOOS = 10
const m_STT_LOPROC = 13
const m_STT_NOTYPE = 0
const m_STT_NUM = 7
const m_STT_OBJECT = 1
const m_STT_PARISC_MILLICODE = 13
const m_STT_SECTION = 3
const m_STT_SPARC_REGISTER = 13
const m_STT_TLS = 6
const m_STV_DEFAULT = 0
const m_STV_HIDDEN = 2
const m_STV_INTERNAL = 1
const m_STV_PROTECTED = 3
const m_SYMINFO_BT_LOWRESERVE = 0xff00
const m_SYMINFO_BT_PARENT = 0xfffe
const m_SYMINFO_BT_SELF = 0xffff
const m_SYMINFO_CURRENT = 1
const m_SYMINFO_FLG_COPY = 0x0004
const m_SYMINFO_FLG_DIRECT = 0x0001
const m_SYMINFO_FLG_LAZYLOAD = 0x0008
const m_SYMINFO_FLG_PASSTHRU = 0x0002
const m_SYMINFO_NONE = 0
const m_SYMINFO_NUM = 2
const m_VER_DEF_CURRENT = 1
const m_VER_DEF_NONE = 0
const m_VER_DEF_NUM = 2
const m_VER_FLG_BASE = 0x1
const m_VER_FLG_WEAK = 0x2
const m_VER_NDX_ELIMINATE = 0xff01
const m_VER_NDX_GLOBAL = 1
const m_VER_NDX_LOCAL = 0
const m_VER_NDX_LORESERVE = 0xff00
const m_VER_NEED_CURRENT = 1
const m_VER_NEED_NONE = 0
const m_VER_NEED_NUM = 2
const m_n_hash = "n_desc"

type Tnlist = struct {
	F__ccgo0_0 struct {
		Fn_un [0]struct {
			Fn_next [0]uintptr
			Fn_strx [0]int64
			Fn_name uintptr
		}
		Fn_name uintptr
	}
	Fn_type  uint8
	Fn_other int8
	Fn_desc  int16
	Fn_value uint64
}

type TElf32_Half = uint16

type TElf64_Half = uint16

type TElf32_Word = uint32

type TElf32_Sword = int32

type TElf64_Word = uint32

type TElf64_Sword = int32

type TElf32_Xword = uint64

type TElf32_Sxword = int64

type TElf64_Xword = uint64

type TElf64_Sxword = int64

type TElf32_Addr = uint32

type TElf64_Addr = uint64

type TElf32_Off = uint32

type TElf64_Off = uint64

type TElf32_Section = uint16

type TElf64_Section = uint16

type TElf32_Versym = uint16

type TElf64_Versym = uint16

type TElf32_Ehdr = struct {
	Fe_ident     [16]uint8
	Fe_type      TElf32_Half
	Fe_machine   TElf32_Half
	Fe_version   TElf32_Word
	Fe_entry     TElf32_Addr
	Fe_phoff     TElf32_Off
	Fe_shoff     TElf32_Off
	Fe_flags     TElf32_Word
	Fe_ehsize    TElf32_Half
	Fe_phentsize TElf32_Half
	Fe_phnum     TElf32_Half
	Fe_shentsize TElf32_Half
	Fe_shnum     TElf32_Half
	Fe_shstrndx  TElf32_Half
}

type TElf64_Ehdr = struct {
	Fe_ident     [16]uint8
	Fe_type      TElf64_Half
	Fe_machine   TElf64_Half
	Fe_version   TElf64_Word
	Fe_entry     TElf64_Addr
	Fe_phoff     TElf64_Off
	Fe_shoff     TElf64_Off
	Fe_flags     TElf64_Word
	Fe_ehsize    TElf64_Half
	Fe_phentsize TElf64_Half
	Fe_phnum     TElf64_Half
	Fe_shentsize TElf64_Half
	Fe_shnum     TElf64_Half
	Fe_shstrndx  TElf64_Half
}

type TElf32_Shdr = struct {
	Fsh_name      TElf32_Word
	Fsh_type      TElf32_Word
	Fsh_flags     TElf32_Word
	Fsh_addr      TElf32_Addr
	Fsh_offset    TElf32_Off
	Fsh_size      TElf32_Word
	Fsh_link      TElf32_Word
	Fsh_info      TElf32_Word
	Fsh_addralign TElf32_Word
	Fsh_entsize   TElf32_Word
}

type TElf64_Shdr = struct {
	Fsh_name      TElf64_Word
	Fsh_type      TElf64_Word
	Fsh_flags     TElf64_Xword
	Fsh_addr      TElf64_Addr
	Fsh_offset    TElf64_Off
	Fsh_size      TElf64_Xword
	Fsh_link      TElf64_Word
	Fsh_info      TElf64_Word
	Fsh_addralign TElf64_Xword
	Fsh_entsize   TElf64_Xword
}

type TElf32_Chdr = struct {
	Fch_type      TElf32_Word
	Fch_size      TElf32_Word
	Fch_addralign TElf32_Word
}

type TElf64_Chdr = struct {
	Fch_type      TElf64_Word
	Fch_reserved  TElf64_Word
	Fch_size      TElf64_Xword
	Fch_addralign TElf64_Xword
}

type TElf32_Sym = struct {
	Fst_name  TElf32_Word
	Fst_value TElf32_Addr
	Fst_size  TElf32_Word
	Fst_info  uint8
	Fst_other uint8
	Fst_shndx TElf32_Section
}

type TElf64_Sym = struct {
	Fst_name  TElf64_Word
	Fst_info  uint8
	Fst_other uint8
	Fst_shndx TElf64_Section
	Fst_value TElf64_Addr
	Fst_size  TElf64_Xword
}

type TElf32_Syminfo = struct {
	Fsi_boundto TElf32_Half
	Fsi_flags   TElf32_Half
}

type TElf64_Syminfo = struct {
	Fsi_boundto TElf64_Half
	Fsi_flags   TElf64_Half
}

type TElf32_Rel = struct {
	Fr_offset TElf32_Addr
	Fr_info   TElf32_Word
}

type TElf64_Rel = struct {
	Fr_offset TElf64_Addr
	Fr_info   TElf64_Xword
}

type TElf32_Rela = struct {
	Fr_offset TElf32_Addr
	Fr_info   TElf32_Word
	Fr_addend TElf32_Sword
}

type TElf64_Rela = struct {
	Fr_offset TElf64_Addr
	Fr_info   TElf64_Xword
	Fr_addend TElf64_Sxword
}

type TElf32_Relr = uint32

type TElf64_Relr = uint64

type TElf32_Phdr = struct {
	Fp_type   TElf32_Word
	Fp_offset TElf32_Off
	Fp_vaddr  TElf32_Addr
	Fp_paddr  TElf32_Addr
	Fp_filesz TElf32_Word
	Fp_memsz  TElf32_Word
	Fp_flags  TElf32_Word
	Fp_align  TElf32_Word
}

type TElf64_Phdr = struct {
	Fp_type   TElf64_Word
	Fp_flags  TElf64_Word
	Fp_offset TElf64_Off
	Fp_vaddr  TElf64_Addr
	Fp_paddr  TElf64_Addr
	Fp_filesz TElf64_Xword
	Fp_memsz  TElf64_Xword
	Fp_align  TElf64_Xword
}

type TElf32_Dyn = struct {
	Fd_tag TElf32_Sword
	Fd_un  struct {
		Fd_ptr [0]TElf32_Addr
		Fd_val TElf32_Word
	}
}

type TElf64_Dyn = struct {
	Fd_tag TElf64_Sxword
	Fd_un  struct {
		Fd_ptr [0]TElf64_Addr
		Fd_val TElf64_Xword
	}
}

type TElf32_Verdef = struct {
	Fvd_version TElf32_Half
	Fvd_flags   TElf32_Half
	Fvd_ndx     TElf32_Half
	Fvd_cnt     TElf32_Half
	Fvd_hash    TElf32_Word
	Fvd_aux     TElf32_Word
	Fvd_next    TElf32_Word
}

type TElf64_Verdef = struct {
	Fvd_version TElf64_Half
	Fvd_flags   TElf64_Half
	Fvd_ndx     TElf64_Half
	Fvd_cnt     TElf64_Half
	Fvd_hash    TElf64_Word
	Fvd_aux     TElf64_Word
	Fvd_next    TElf64_Word
}

type TElf32_Verdaux = struct {
	Fvda_name TElf32_Word
	Fvda_next TElf32_Word
}

type TElf64_Verdaux = struct {
	Fvda_name TElf64_Word
	Fvda_next TElf64_Word
}

type TElf32_Verneed = struct {
	Fvn_version TElf32_Half
	Fvn_cnt     TElf32_Half
	Fvn_file    TElf32_Word
	Fvn_aux     TElf32_Word
	Fvn_next    TElf32_Word
}

type TElf64_Verneed = struct {
	Fvn_version TElf64_Half
	Fvn_cnt     TElf64_Half
	Fvn_file    TElf64_Word
	Fvn_aux     TElf64_Word
	Fvn_next    TElf64_Word
}

type TElf32_Vernaux = struct {
	Fvna_hash  TElf32_Word
	Fvna_flags TElf32_Half
	Fvna_other TElf32_Half
	Fvna_name  TElf32_Word
	Fvna_next  TElf32_Word
}

type TElf64_Vernaux = struct {
	Fvna_hash  TElf64_Word
	Fvna_flags TElf64_Half
	Fvna_other TElf64_Half
	Fvna_name  TElf64_Word
	Fvna_next  TElf64_Word
}

type TElf32_auxv_t = struct {
	Fa_type Tuint32_t
	Fa_un   struct {
		Fa_val Tuint32_t
	}
}

type TElf64_auxv_t = struct {
	Fa_type Tuint64_t
	Fa_un   struct {
		Fa_val Tuint64_t
	}
}

type TElf32_Nhdr = struct {
	Fn_namesz TElf32_Word
	Fn_descsz TElf32_Word
	Fn_type   TElf32_Word
}

type TElf64_Nhdr = struct {
	Fn_namesz TElf64_Word
	Fn_descsz TElf64_Word
	Fn_type   TElf64_Word
}

type TElf32_Move = struct {
	Fm_value   TElf32_Xword
	Fm_info    TElf32_Word
	Fm_poffset TElf32_Word
	Fm_repeat  TElf32_Half
	Fm_stride  TElf32_Half
}

type TElf64_Move = struct {
	Fm_value   TElf64_Xword
	Fm_info    TElf64_Xword
	Fm_poffset TElf64_Xword
	Fm_repeat  TElf64_Half
	Fm_stride  TElf64_Half
}

type TElf32_gptab = struct {
	Fgt_entry [0]struct {
		Fgt_g_value TElf32_Word
		Fgt_bytes   TElf32_Word
	}
	Fgt_header struct {
		Fgt_current_g_value TElf32_Word
		Fgt_unused          TElf32_Word
	}
}

type TElf32_RegInfo = struct {
	Fri_gprmask  TElf32_Word
	Fri_cprmask  [4]TElf32_Word
	Fri_gp_value TElf32_Sword
}

type TElf_Options = struct {
	Fkind    uint8
	Fsize    uint8
	Fsection TElf32_Section
	Finfo    TElf32_Word
}

type TElf_Options_Hw = struct {
	Fhwp_flags1 TElf32_Word
	Fhwp_flags2 TElf32_Word
}

type TElf32_Lib = struct {
	Fl_name       TElf32_Word
	Fl_time_stamp TElf32_Word
	Fl_checksum   TElf32_Word
	Fl_version    TElf32_Word
	Fl_flags      TElf32_Word
}

type TElf64_Lib = struct {
	Fl_name       TElf64_Word
	Fl_time_stamp TElf64_Word
	Fl_checksum   TElf64_Word
	Fl_version    TElf64_Word
	Fl_flags      TElf64_Word
}

type TElf32_Conflict = uint32

type TElf_MIPS_ABIFlags_v0 = struct {
	Fversion   TElf32_Half
	Fisa_level uint8
	Fisa_rev   uint8
	Fgpr_size  uint8
	Fcpr1_size uint8
	Fcpr2_size uint8
	Ffp_abi    uint8
	Fisa_ext   TElf32_Word
	Fases      TElf32_Word
	Fflags1    TElf32_Word
	Fflags2    TElf32_Word
}

const _Val_GNU_MIPS_ABI_FP_ANY = 0
const _Val_GNU_MIPS_ABI_FP_DOUBLE = 1
const _Val_GNU_MIPS_ABI_FP_SINGLE = 2
const _Val_GNU_MIPS_ABI_FP_SOFT = 3
const _Val_GNU_MIPS_ABI_FP_OLD_64 = 4
const _Val_GNU_MIPS_ABI_FP_XX = 5
const _Val_GNU_MIPS_ABI_FP_64 = 6
const _Val_GNU_MIPS_ABI_FP_64A = 7
const _Val_GNU_MIPS_ABI_FP_MAX = 7

// C documentation
//
//	/*
//	 * __elf_is_okay__ - Determine if ehdr really
//	 * is ELF and valid for the target platform.
//	 *
//	 * WARNING:  This is NOT an ELF ABI function and
//	 * as such its use should be restricted.
//	 */
func ___elf_is_okay__(tls *libc.TLS, ehdr uintptr) (r int32) {
	var retval int32
	_ = retval
	retval = 0
	/*
	 * We need to check magic, class size, endianess,
	 * and version before we look at the rest of the
	 * Elf_Ehdr structure.  These few elements are
	 * represented in a machine independent fashion.
	 */
	if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr))) == int32(m_ELFMAG0) && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr + 1))) == int32('E') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr + 2))) == int32('L') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr + 3))) == int32('F') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr + 4))) == int32(m_ELFCLASS64) && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr + 5))) == int32(m_ELFDATA2LSB) && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ehdr + 6))) == int32(m_EV_CURRENT) {
		/* Now check the machine dependent header */
		if libc.Int32FromUint16((*TElf64_Ehdr)(unsafe.Pointer(ehdr)).Fe_machine) == int32(m_EM_LOONGARCH) && (*TElf64_Ehdr)(unsafe.Pointer(ehdr)).Fe_version == uint32(m_EV_CURRENT) {
			retval = int32(1)
		}
	}
	return retval
}

// C documentation
//
//	/*
//	 * Convert an Elf_Sym into an nlist structure.  This fills in only the
//	 * n_value and n_type members.
//	 */
func _elf_sym_to_nlist(tls *libc.TLS, nl uintptr, s uintptr, shdr uintptr, shnum int32) {
	var sh, p5 uintptr
	var v1, v2, v3, v4 int32
	_, _, _, _, _, _ = sh, v1, v2, v3, v4, p5
	(*Tnlist)(unsafe.Pointer(nl)).Fn_value = (*TElf64_Sym)(unsafe.Pointer(s)).Fst_value
	switch libc.Int32FromUint16((*TElf64_Sym)(unsafe.Pointer(s)).Fst_shndx) {
	case m_SHN_UNDEF:
		fallthrough
	case int32(m_SHN_COMMON):
		(*Tnlist)(unsafe.Pointer(nl)).Fn_type = uint8(m_N_UNDF)
	case int32(m_SHN_ABS):
		if libc.Int32FromUint8((*TElf64_Sym)(unsafe.Pointer(s)).Fst_info)&int32(0xf) == int32(m_STT_FILE) {
			v1 = int32(m_N_FN)
		} else {
			v1 = int32(m_N_ABS)
		}
		(*Tnlist)(unsafe.Pointer(nl)).Fn_type = libc.Uint8FromInt32(v1)
	default:
		if libc.Int32FromUint16((*TElf64_Sym)(unsafe.Pointer(s)).Fst_shndx) >= shnum {
			(*Tnlist)(unsafe.Pointer(nl)).Fn_type = uint8(m_N_UNDF)
		} else {
			sh = shdr + uintptr((*TElf64_Sym)(unsafe.Pointer(s)).Fst_shndx)*64
			if (*TElf64_Shdr)(unsafe.Pointer(sh)).Fsh_type == uint32(m_SHT_PROGBITS) {
				if (*TElf64_Shdr)(unsafe.Pointer(sh)).Fsh_flags&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(0)) != 0 {
					v3 = int32(m_N_DATA)
				} else {
					v3 = int32(m_N_TEXT)
				}
				v2 = v3
			} else {
				if (*TElf64_Shdr)(unsafe.Pointer(sh)).Fsh_type == uint32(m_SHT_NOBITS) {
					v4 = int32(m_N_BSS)
				} else {
					v4 = m_N_UNDF
				}
				v2 = v4
			}
			(*Tnlist)(unsafe.Pointer(nl)).Fn_type = libc.Uint8FromInt32(v2)
		}
		break
	}
	if libc.Int32FromUint8((*TElf64_Sym)(unsafe.Pointer(s)).Fst_info)>>int32(4) == int32(m_STB_GLOBAL) || libc.Int32FromUint8((*TElf64_Sym)(unsafe.Pointer(s)).Fst_info)>>int32(4) == int32(m_STB_WEAK) {
		p5 = nl + 8
		*(*uint8)(unsafe.Pointer(p5)) = uint8(int32(*(*uint8)(unsafe.Pointer(p5))) | libc.Int32FromInt32(m_N_EXT))
	}
}

func X__fdnlist(tls *libc.TLS, fd int32, list uintptr) (r int32) {
	bp := tls.Alloc(24768)
	defer tls.Free(24768)
	var cc, i TElf64_Sword
	var errsave, nent, v6 int32
	var name, p, s, shdr, strtab uintptr
	var shdr_size, size, symsize, symstrsize TElf64_Word
	var symoff, symstroff TElf64_Off
	var v3 uint64
	var _ /* ehdr at bp+24576 */ TElf64_Ehdr
	var _ /* sbuf at bp+0 */ [1024]TElf64_Sym
	var _ /* st at bp+24640 */ Tstat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cc, errsave, i, name, nent, p, s, shdr, shdr_size, size, strtab, symoff, symsize, symstroff, symstrsize, v3, v6
	symoff = uint64(0)
	symstroff = uint64(0)
	symsize = uint32(0)
	symstrsize = uint32(0)
	nent = -int32(1)
	strtab = libc.UintptrFromInt32(0)
	shdr = libc.UintptrFromInt32(0)
	/* Make sure obj is OK */
	if libc.Xlseek(tls, fd, libc.Int64FromInt32(0), 0) == int64(-int32(1)) || libc.Uint64FromInt64(libc.Xread(tls, fd, bp+24576, uint64(64))) != uint64(64) || !(___elf_is_okay__(tls, bp+24576) != 0) || libc.Xfstat(tls, fd, bp+24640) < 0 {
		return -int32(1)
	}
	if libc.Int32FromUint16((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shnum) == 0 || uint64((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shentsize) != uint64(64) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ERANGE)
		return -int32(1)
	}
	/* calculate section header table size */
	shdr_size = libc.Uint32FromInt32(libc.Int32FromUint16((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shentsize) * libc.Int32FromUint16((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shnum))
	/* Make sure it's not too big to mmap */
	if libc.Int64FromUint32(shdr_size) > (*(*Tstat)(unsafe.Pointer(bp + 24640))).Fst_size {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EFBIG)
		return -int32(1)
	}
	shdr = libc.Xmalloc(tls, uint64(shdr_size))
	if shdr == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	/* Load section header table. */
	if libc.Xpread(tls, fd, shdr, uint64(shdr_size), libc.Int64FromUint64((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shoff)) != libc.Int64FromUint32(shdr_size) {
		goto done
	}
	/*
	 * Find the symbol table entry and it's corresponding
	 * string table entry.	Version 1.1 of the ABI states
	 * that there is only one symbol table but that this
	 * could change in the future.
	 */
	i = 0
	for {
		if !(i < libc.Int32FromUint16((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shnum)) {
			break
		}
		if (*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr(i)*64))).Fsh_type == uint32(m_SHT_SYMTAB) {
			if (*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr(i)*64))).Fsh_link >= uint32((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shnum) {
				goto done
			}
			symoff = (*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr(i)*64))).Fsh_offset
			symsize = uint32((*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr(i)*64))).Fsh_size)
			symstroff = (*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr((*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr(i)*64))).Fsh_link)*64))).Fsh_offset
			symstrsize = uint32((*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr((*(*TElf64_Shdr)(unsafe.Pointer(shdr + uintptr(i)*64))).Fsh_link)*64))).Fsh_size)
			break
		}
		goto _1
	_1:
		;
		i++
	}
	/* Check for files too large to mmap. */
	if libc.Int64FromUint32(symstrsize) > (*(*Tstat)(unsafe.Pointer(bp + 24640))).Fst_size {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EFBIG)
		goto done
	}
	/*
	 * Load string table into our address space.  This gives us
	 * an easy way to randomly access all the strings, without
	 * making the memory allocation permanent as with malloc/free
	 * (i.e., munmap will return it to the system).
	 */
	strtab = libc.Xmalloc(tls, uint64(symstrsize))
	if strtab == libc.UintptrFromInt32(0) {
		goto done
	}
	if libc.Xpread(tls, fd, strtab, uint64(symstrsize), libc.Int64FromUint64(symstroff)) != libc.Int64FromUint32(symstrsize) {
		goto done
	}
	/*
	 * clean out any left-over information for all valid entries.
	 * Type and value defined to be 0 if not found; historical
	 * versions cleared other and desc as well.  Also figure out
	 * the largest string length so don't read any more of the
	 * string table than we have to.
	 *
	 * XXX clearing anything other than n_type and n_value violates
	 * the semantics given in the man page.
	 */
	nent = 0
	p = list
	for {
		if !!(*(*uintptr)(unsafe.Pointer(p)) == uintptr(0) || int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(p))))) == 0) {
			break
		}
		(*Tnlist)(unsafe.Pointer(p)).Fn_type = uint8(0)
		(*Tnlist)(unsafe.Pointer(p)).Fn_other = 0
		(*Tnlist)(unsafe.Pointer(p)).Fn_desc = 0
		(*Tnlist)(unsafe.Pointer(p)).Fn_value = uint64(0)
		nent++
		goto _2
	_2:
		;
		p += 24
	}
	/* Don't process any further if object is stripped. */
	if symoff == uint64(0) {
		goto done
	}
	if libc.Xlseek(tls, fd, libc.Int64FromUint64(symoff), 0) == int64(-int32(1)) {
		nent = -int32(1)
		goto done
	}
	for symsize > uint32(0) && nent > 0 {
		if uint64(symsize) < libc.Uint64FromInt64(24576) {
			v3 = uint64(symsize)
		} else {
			v3 = libc.Uint64FromInt64(24576)
		}
		cc = libc.Int32FromUint64(v3)
		if libc.Xread(tls, fd, bp, libc.Uint64FromInt32(cc)) != int64(cc) {
			break
		}
		symsize -= libc.Uint32FromInt32(cc)
		s = bp
		for {
			if !(cc > 0 && nent > 0) {
				break
			}
			name = strtab + uintptr((*TElf64_Sym)(unsafe.Pointer(s)).Fst_name)
			if int32(*(*int8)(unsafe.Pointer(name))) == int32('\000') {
				goto _4
			}
			size = symstrsize - (*TElf64_Sym)(unsafe.Pointer(s)).Fst_name
			p = list
			for {
				if !!(*(*uintptr)(unsafe.Pointer(p)) == uintptr(0) || int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(p))))) == 0) {
					break
				}
				if int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(p))))) == int32('_') && libc.Xstrncmp(tls, name, *(*uintptr)(unsafe.Pointer(p))+uintptr(1), uint64(size)) == 0 || libc.Xstrncmp(tls, name, *(*uintptr)(unsafe.Pointer(p)), uint64(size)) == 0 {
					_elf_sym_to_nlist(tls, p, s, shdr, libc.Int32FromUint16((*(*TElf64_Ehdr)(unsafe.Pointer(bp + 24576))).Fe_shnum))
					nent--
					v6 = nent
					if v6 <= 0 {
						break
					}
				}
				goto _5
			_5:
				;
				p += 24
			}
			goto _4
		_4:
			;
			s += 24
			cc = TElf64_Sword(uint64(cc) - libc.Uint64FromInt64(24))
		}
	}
	goto done
done:
	;
	errsave = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xfree(tls, strtab)
	libc.Xfree(tls, shdr)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = errsave
	return nent
}

func Xnlist(tls *libc.TLS, name uintptr, list uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var fd, n int32
	_, _ = fd, n
	if list == libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	}
	fd = libc.Xopen(tls, name, m_O_RDONLY1, libc.VaList(bp+8, 0))
	if fd < 0 {
		return -int32(1)
	}
	n = X__fdnlist(tls, fd, list)
	libc.Xclose(tls, fd)
	return n
}

const m_O_CREAT3 = 64
const m_O_NONBLOCK3 = 2048
const m_O_TRUNC3 = 512
const m_O_WRONLY1 = 1

type Tpidfh = struct {
	Fpf_fd   int32
	Fpf_path uintptr
	Fpf_dev  Tdev_t
	Fpf_ino  Tino_t
}

func _pidfile_verify(tls *libc.TLS, pfh uintptr) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _ /* sb at bp+0 */ Tstat
	if pfh == libc.UintptrFromInt32(0) || (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd == -int32(1) {
		return int32(m_EINVAL)
	}
	/*
	 * Check remembered descriptor.
	 */
	if libc.Xfstat(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd, bp) == -int32(1) {
		return *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	}
	if (*(*Tstat)(unsafe.Pointer(bp))).Fst_dev != (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_dev || (*(*Tstat)(unsafe.Pointer(bp))).Fst_ino != (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_ino {
		return int32(m_EINVAL)
	}
	return 0
}

func _pidfile_read(tls *libc.TLS, path uintptr, pidptr uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var error1, fd, i int32
	var _ /* buf at bp+0 */ [16]int8
	var _ /* endptr at bp+16 */ uintptr
	_, _, _ = error1, fd, i
	fd = libc.Xopen(tls, path, m_O_RDONLY1, 0)
	if fd == -int32(1) {
		return *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	}
	i = int32(libc.Xread(tls, fd, bp, libc.Uint64FromInt64(16)-libc.Uint64FromInt32(1)))
	error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) /* Remember errno in case close() wants to change it. */
	libc.Xclose(tls, fd)
	if i == -int32(1) {
		return error1
	} else {
		if i == 0 {
			return int32(m_EAGAIN)
		}
	}
	(*(*[16]int8)(unsafe.Pointer(bp)))[i] = int8('\000')
	*(*Tpid_t)(unsafe.Pointer(pidptr)) = int32(libc.Xstrtol(tls, bp, bp+16, int32(10)))
	if *(*uintptr)(unsafe.Pointer(bp + 16)) != bp+uintptr(i) {
		return int32(m_EINVAL)
	}
	return 0
}

func Xpidfile_open(tls *libc.TLS, path uintptr, mode Tmode_t, pidptr uintptr) (r uintptr) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var count, error1, fd, len1, v2 int32
	var pfh uintptr
	var v3 bool
	var _ /* rqtp at bp+128 */ Ttimespec
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _ = count, error1, fd, len1, pfh, v2, v3
	pfh = libc.Xmalloc(tls, uint64(32))
	if pfh == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if path == libc.UintptrFromInt32(0) {
		len1 = libc.Xasprintf(tls, pfh+8, __ccgo_ts+310, libc.VaList(bp+152, Xgetprogname(tls)))
		if len1 < 0 {
			libc.Xfree(tls, pfh)
			return libc.UintptrFromInt32(0)
		}
	} else {
		(*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path = libc.Xstrdup(tls, path)
	}
	/*
	 * Open the PID file and obtain exclusive lock.
	 * We truncate PID file here only to remove old PID immediately,
	 * PID file will be truncated again in pidfile_write(), so
	 * pidfile_write() can be called multiple times.
	 */
	fd = Xflopen(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path, libc.Int32FromInt32(m_O_WRONLY1)|libc.Int32FromInt32(m_O_CREAT3)|libc.Int32FromInt32(m_O_TRUNC3)|libc.Int32FromInt32(m_O_NONBLOCK3), libc.VaList(bp+152, mode))
	if fd == -int32(1) {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) {
			if pidptr == libc.UintptrFromInt32(0) {
				*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EEXIST)
			} else {
				count = int32(20)
				(*(*Ttimespec)(unsafe.Pointer(bp + 128))).Ftv_sec = 0
				(*(*Ttimespec)(unsafe.Pointer(bp + 128))).Ftv_nsec = int64(5000000)
				for {
					*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = _pidfile_read(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path, pidptr)
					if v3 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(m_EAGAIN); !v3 {
						count--
						v2 = count
					}
					if v3 || v2 == 0 {
						break
					}
					libc.Xnanosleep(tls, bp+128, uintptr(0))
					goto _1
				_1:
				}
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) {
					*(*Tpid_t)(unsafe.Pointer(pidptr)) = -int32(1)
				}
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == 0 || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) {
					*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EEXIST)
				}
			}
		}
		error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		libc.Xfree(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path)
		libc.Xfree(tls, pfh)
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return libc.UintptrFromInt32(0)
	}
	/*
	 * Remember file information, so in pidfile_write() we are sure we write
	 * to the proper descriptor.
	 */
	if libc.Xfstat(tls, fd, bp) == -int32(1) {
		error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		libc.Xunlink(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path)
		libc.Xfree(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path)
		libc.Xclose(tls, fd)
		libc.Xfree(tls, pfh)
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return libc.UintptrFromInt32(0)
	}
	(*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd = fd
	(*Tpidfh)(unsafe.Pointer(pfh)).Fpf_dev = (*(*Tstat)(unsafe.Pointer(bp))).Fst_dev
	(*Tpidfh)(unsafe.Pointer(pfh)).Fpf_ino = (*(*Tstat)(unsafe.Pointer(bp))).Fst_ino
	return pfh
}

func Xpidfile_write(tls *libc.TLS, pfh uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var error1, fd int32
	var _ /* pidstr at bp+0 */ [16]int8
	_, _ = error1, fd
	/*
	 * Check remembered descriptor, so we don't overwrite some other
	 * file if pidfile was closed and descriptor reused.
	 */
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = _pidfile_verify(tls, pfh)
	if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
		/*
		 * Don't close descriptor, because we are not sure if it's ours.
		 */
		return -int32(1)
	}
	fd = (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd
	/*
	 * Truncate PID file, so multiple calls of pidfile_write() are allowed.
	 */
	if libc.Xftruncate(tls, fd, 0) == -int32(1) {
		error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		__pidfile_remove(tls, pfh, 0)
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return -int32(1)
	}
	libc.X__builtin_snprintf(tls, bp, uint64(16), __ccgo_ts+326, libc.VaList(bp+24, libc.Xgetpid(tls)))
	if libc.Xpwrite(tls, fd, bp, libc.Xstrlen(tls, bp), 0) != libc.Int64FromUint64(libc.Xstrlen(tls, bp)) {
		error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		__pidfile_remove(tls, pfh, 0)
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return -int32(1)
	}
	return 0
}

func Xpidfile_close(tls *libc.TLS, pfh uintptr) (r int32) {
	var error1 int32
	_ = error1
	error1 = _pidfile_verify(tls, pfh)
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return -int32(1)
	}
	if libc.Xclose(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd) == -int32(1) {
		error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	}
	libc.Xfree(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path)
	libc.Xfree(tls, pfh)
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return -int32(1)
	}
	return 0
}

func __pidfile_remove(tls *libc.TLS, pfh uintptr, freeit int32) (r int32) {
	var error1 int32
	_ = error1
	error1 = _pidfile_verify(tls, pfh)
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return -int32(1)
	}
	if libc.Xunlink(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path) == -int32(1) {
		error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	}
	if libc.Xclose(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd) == -int32(1) {
		if error1 == 0 {
			error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		}
	}
	if freeit != 0 {
		libc.Xfree(tls, (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_path)
		libc.Xfree(tls, pfh)
	} else {
		(*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd = -int32(1)
	}
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = error1
		return -int32(1)
	}
	return 0
}

func Xpidfile_remove(tls *libc.TLS, pfh uintptr) (r int32) {
	return __pidfile_remove(tls, pfh, int32(1))
}

func Xpidfile_fileno(tls *libc.TLS, pfh uintptr) (r int32) {
	if pfh == libc.UintptrFromInt32(0) || (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd == -int32(1) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	}
	return (*Tpidfh)(unsafe.Pointer(pfh)).Fpf_fd
}

const m_SPT_MAXTITLE = 255
const m___bool_true_false_are_defined = 1
const m_bool = "_Bool"
const m_false = 0
const m_true = 1

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

var _SPT struct {
	Farg0   uintptr
	Fbase   uintptr
	Fend    uintptr
	Fnul    uintptr
	Fwarned uint8
	Freset  uint8
	Ferror1 int32
}

func _spt_min(tls *libc.TLS, a Tsize_t, b Tsize_t) (r Tsize_t) {
	var v1 uint64
	_ = v1
	if a < b {
		v1 = a
	} else {
		v1 = b
	}
	return v1
}

// C documentation
//
//	/*
//	 * For discussion on the portability of the various methods, see
//	 * https://lists.freebsd.org/pipermail/freebsd-stable/2008-June/043136.html
//	 */
func _spt_clearenv(tls *libc.TLS) (r int32) {
	return libc.Xclearenv(tls)
}

func _spt_copyenv(tls *libc.TLS, envc int32, envp uintptr) (r int32) {
	var envcopy, eq uintptr
	var envsize, error1, i int32
	_, _, _, _, _ = envcopy, envsize, eq, error1, i
	if libc.Xenviron != envp {
		return 0
	}
	/* Make a copy of the old environ array of pointers, in case
	 * clearenv() or setenv() is implemented to free the internal
	 * environ array, because we will need to access the old environ
	 * contents to make the new copy. */
	envsize = libc.Int32FromUint64(libc.Uint64FromInt32(envc+libc.Int32FromInt32(1)) * uint64(8))
	envcopy = libc.Xmalloc(tls, libc.Uint64FromInt32(envsize))
	if envcopy == libc.UintptrFromInt32(0) {
		return *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	}
	libc.Xmemcpy(tls, envcopy, envp, libc.Uint64FromInt32(envsize))
	error1 = _spt_clearenv(tls)
	if error1 != 0 {
		libc.Xenviron = envp
		libc.Xfree(tls, envcopy)
		return error1
	}
	i = 0
	for {
		if !(*(*uintptr)(unsafe.Pointer(envcopy + uintptr(i)*8)) != 0) {
			break
		}
		eq = libc.Xstrchr(tls, *(*uintptr)(unsafe.Pointer(envcopy + uintptr(i)*8)), int32('='))
		if eq == libc.UintptrFromInt32(0) {
			goto _1
		}
		*(*int8)(unsafe.Pointer(eq)) = int8('\000')
		if libc.Xsetenv(tls, *(*uintptr)(unsafe.Pointer(envcopy + uintptr(i)*8)), eq+uintptr(1), int32(1)) < 0 {
			error1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		}
		*(*int8)(unsafe.Pointer(eq)) = int8('=')
		if error1 != 0 {
			/* Because the old environ might not be available
			 * anymore we will make do with the shallow copy. */
			libc.Xenviron = envcopy
			return error1
		}
		goto _1
	_1:
		;
		i++
	}
	/* Dispose of the shallow copy, now that we've finished transfering
	 * the old environment. */
	libc.Xfree(tls, envcopy)
	return 0
}

func _spt_copyargs(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	var i int32
	var tmp uintptr
	_, _ = i, tmp
	i = int32(1)
	for {
		if !(i < argc || i >= argc && *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) != 0) {
			break
		}
		if *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) == libc.UintptrFromInt32(0) {
			goto _1
		}
		tmp = libc.Xstrdup(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)))
		if tmp == libc.UintptrFromInt32(0) {
			return *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		}
		*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) = tmp
		goto _1
	_1:
		;
		i++
	}
	return 0
}

func Xsetproctitle_init(tls *libc.TLS, argc int32, argv uintptr, envp uintptr) {
	var base, end, nul, tmp uintptr
	var envc, error1, i int32
	_, _, _, _, _, _, _ = base, end, envc, error1, i, nul, tmp
	/* Try to make sure we got called with main() arguments. */
	if argc < 0 {
		return
	}
	base = *(*uintptr)(unsafe.Pointer(argv))
	if base == libc.UintptrFromInt32(0) {
		return
	}
	nul = base + uintptr(libc.Xstrlen(tls, base))
	end = nul + uintptr(1)
	i = 0
	for {
		if !(i < argc || i >= argc && *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) != 0) {
			break
		}
		if *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) == libc.UintptrFromInt32(0) || *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) != end {
			goto _1
		}
		end = *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) + uintptr(libc.Xstrlen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)))) + uintptr(1)
		goto _1
	_1:
		;
		i++
	}
	i = 0
	for {
		if !(*(*uintptr)(unsafe.Pointer(envp + uintptr(i)*8)) != 0) {
			break
		}
		if *(*uintptr)(unsafe.Pointer(envp + uintptr(i)*8)) != end {
			goto _2
		}
		end = *(*uintptr)(unsafe.Pointer(envp + uintptr(i)*8)) + uintptr(libc.Xstrlen(tls, *(*uintptr)(unsafe.Pointer(envp + uintptr(i)*8)))) + uintptr(1)
		goto _2
	_2:
		;
		i++
	}
	envc = i
	_SPT.Farg0 = libc.Xstrdup(tls, *(*uintptr)(unsafe.Pointer(argv)))
	if _SPT.Farg0 == libc.UintptrFromInt32(0) {
		_SPT.Ferror1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		return
	}
	tmp = libc.Xstrdup(tls, Xgetprogname(tls))
	if tmp == libc.UintptrFromInt32(0) {
		_SPT.Ferror1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		return
	}
	Xsetprogname(tls, tmp)
	error1 = _spt_copyenv(tls, envc, envp)
	if error1 != 0 {
		_SPT.Ferror1 = error1
		return
	}
	error1 = _spt_copyargs(tls, argc, argv)
	if error1 != 0 {
		_SPT.Ferror1 = error1
		return
	}
	_SPT.Fnul = nul
	_SPT.Fbase = base
	_SPT.Fend = end
}

func Xsetproctitle_impl(tls *libc.TLS, fmt uintptr, va uintptr) {
	bp := tls.Alloc(272)
	defer tls.Free(272)
	var ap Tva_list
	var len1 int32
	var nul, v1 uintptr
	var _ /* buf at bp+0 */ [256]int8
	_, _, _, _ = ap, len1, nul, v1
	if _SPT.Fbase == libc.UintptrFromInt32(0) {
		if !(_SPT.Fwarned != 0) {
			libc.Xwarnx(tls, __ccgo_ts+329, 0)
			_SPT.Fwarned = uint8(m_true)
		}
		return
	}
	if fmt != 0 {
		if int32(*(*int8)(unsafe.Pointer(fmt))) == int32('-') {
			/* Skip program name prefix. */
			fmt++
			len1 = 0
		} else {
			/* Print program name heading for grep. */
			libc.X__builtin_snprintf(tls, bp, uint64(256), __ccgo_ts+14, libc.VaList(bp+264, Xgetprogname(tls)))
			len1 = libc.Int32FromUint64(libc.Xstrlen(tls, bp))
		}
		ap = va
		len1 += libc.X__builtin_vsnprintf(tls, bp+uintptr(len1), uint64(256)-libc.Uint64FromInt32(len1), fmt, ap)
		_ = ap
	} else {
		len1 = libc.X__builtin_snprintf(tls, bp, uint64(256), __ccgo_ts+427, libc.VaList(bp+264, _SPT.Farg0))
	}
	if len1 <= 0 {
		_SPT.Ferror1 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		return
	}
	if !(_SPT.Freset != 0) {
		libc.Xmemset(tls, _SPT.Fbase, 0, libc.Uint64FromInt64(int64(_SPT.Fend)-int64(_SPT.Fbase)))
		_SPT.Freset = uint8(m_true)
	} else {
		libc.Xmemset(tls, _SPT.Fbase, 0, _spt_min(tls, uint64(256), libc.Uint64FromInt64(int64(_SPT.Fend)-int64(_SPT.Fbase))))
	}
	len1 = libc.Int32FromUint64(_spt_min(tls, libc.Uint64FromInt32(len1), _spt_min(tls, uint64(256), libc.Uint64FromInt64(int64(_SPT.Fend)-int64(_SPT.Fbase)))-uint64(1)))
	libc.Xmemcpy(tls, _SPT.Fbase, bp, libc.Uint64FromInt32(len1))
	nul = _SPT.Fbase + uintptr(len1)
	if nul < _SPT.Fnul {
		*(*int8)(unsafe.Pointer(_SPT.Fnul)) = int8('.')
	} else {
		if nul == _SPT.Fnul && nul+uintptr(1) < _SPT.Fend {
			*(*int8)(unsafe.Pointer(_SPT.Fnul)) = int8(' ')
			nul++
			v1 = nul
			*(*int8)(unsafe.Pointer(v1)) = int8('\000')
		}
	}
}

func Xgetprogname(tls *libc.TLS) (r uintptr) {
	if libc.Xprogram_invocation_short_name == libc.UintptrFromInt32(0) {
		libc.Xprogram_invocation_short_name = libc.Xprogram_invocation_short_name
	}
	return libc.Xprogram_invocation_short_name
}

func Xsetprogname(tls *libc.TLS, progname uintptr) {
	var i Tsize_t
	_ = i
	i = libc.Xstrlen(tls, progname)
	for {
		if !(i > uint64(0)) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(progname + uintptr(i-uint64(1))))) == int32('/') {
			libc.Xprogram_invocation_short_name = progname + uintptr(i)
			return
		}
		goto _1
	_1:
		;
		i--
	}
	libc.Xprogram_invocation_short_name = progname
}

const m_GID_SZ = 251
const m_GNMLEN = 32
const m_GNM_SZ = 251
const m_INVALID = 2
const m_UID_SZ = 317
const m_UNMLEN = 32
const m_UNM_SZ = 317
const m_VALID = 1
const m__PW_BUF_LEN = 1024

type Tgroup = struct {
	Fgr_name   uintptr
	Fgr_passwd uintptr
	Fgr_gid    Tgid_t
	Fgr_mem    uintptr
}

type Tpasswd = struct {
	Fpw_name   uintptr
	Fpw_passwd uintptr
	Fpw_uid    Tuid_t
	Fpw_gid    Tgid_t
	Fpw_gecos  uintptr
	Fpw_dir    uintptr
	Fpw_shell  uintptr
}

/*
 * Constants and data structures used to implement group and password file
 * caches.  Name lengths have been chosen to be as large as those supported
 * by the passwd and group files as well as the standard archive formats.
 * CACHE SIZES MUST BE PRIME
 */

/*
 * Node structures used in the user, group, uid, and gid caches.
 */

type TUIDC = struct {
	Fvalid int32
	Fname  [32]int8
	Fuid   Tuid_t
}

/*
 * Constants and data structures used to implement group and password file
 * caches.  Name lengths have been chosen to be as large as those supported
 * by the passwd and group files as well as the standard archive formats.
 * CACHE SIZES MUST BE PRIME
 */

/*
 * Node structures used in the user, group, uid, and gid caches.
 */

type Tuidc = TUIDC

type TGIDC = struct {
	Fvalid int32
	Fname  [32]int8
	Fgid   Tgid_t
}

type Tgidc = TGIDC

/*
 * Routines that control user, group, uid and gid caches.
 * Traditional passwd/group cache routines perform quite poorly with
 * archives. The chances of hitting a valid lookup with an archive is quite a
 * bit worse than with files already resident on the file system. These misses
 * create a MAJOR performance cost. To adress this problem, these routines
 * cache both hits and misses.
 */

var _uidtb uintptr /* uid to name cache */
var _gidtb uintptr /* gid to name cache */
var _usrtb uintptr /* user name to uid cache */
var _grptb uintptr /* group name to gid cache */

func _st_hash(tls *libc.TLS, name uintptr, len1 Tsize_t, tabsz int32) (r Tu_int) {
	var key Tu_int
	var v1 Tsize_t
	var v2 uintptr
	_, _, _ = key, v1, v2
	key = uint32(0)
	for {
		v1 = len1
		len1--
		if !(v1 != 0) {
			break
		}
		v2 = name
		name++
		key += libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(v2)))
		key = key<<libc.Int32FromInt32(8) | key>>libc.Int32FromInt32(24)
	}
	return key % libc.Uint32FromInt32(tabsz)
}

// C documentation
//
//	/*
//	 * uidtb_start
//	 *	creates an an empty uidtb
//	 * Return:
//	 *	0 if ok, -1 otherwise
//	 */
func _uidtb_start(tls *libc.TLS) (r int32) {
	var v1 uintptr
	_ = v1
	if _uidtb != libc.UintptrFromInt32(0) {
		return 0
	}
	if _fail != 0 {
		return -int32(1)
	}
	v1 = libc.Xcalloc(tls, uint64(m_UID_SZ), uint64(8))
	_uidtb = v1
	if v1 == libc.UintptrFromInt32(0) {
		_fail++
		return -int32(1)
	}
	return 0
}

var _fail int32

// C documentation
//
//	/*
//	 * gidtb_start
//	 *	creates an an empty gidtb
//	 * Return:
//	 *	0 if ok, -1 otherwise
//	 */
func _gidtb_start(tls *libc.TLS) (r int32) {
	var v1 uintptr
	_ = v1
	if _gidtb != libc.UintptrFromInt32(0) {
		return 0
	}
	if _fail1 != 0 {
		return -int32(1)
	}
	v1 = libc.Xcalloc(tls, uint64(m_GID_SZ), uint64(8))
	_gidtb = v1
	if v1 == libc.UintptrFromInt32(0) {
		_fail1++
		return -int32(1)
	}
	return 0
}

var _fail1 int32

// C documentation
//
//	/*
//	 * usrtb_start
//	 *	creates an an empty usrtb
//	 * Return:
//	 *	0 if ok, -1 otherwise
//	 */
func _usrtb_start(tls *libc.TLS) (r int32) {
	var v1 uintptr
	_ = v1
	if _usrtb != libc.UintptrFromInt32(0) {
		return 0
	}
	if _fail2 != 0 {
		return -int32(1)
	}
	v1 = libc.Xcalloc(tls, uint64(m_UNM_SZ), uint64(8))
	_usrtb = v1
	if v1 == libc.UintptrFromInt32(0) {
		_fail2++
		return -int32(1)
	}
	return 0
}

var _fail2 int32

// C documentation
//
//	/*
//	 * grptb_start
//	 *	creates an an empty grptb
//	 * Return:
//	 *	0 if ok, -1 otherwise
//	 */
func _grptb_start(tls *libc.TLS) (r int32) {
	var v1 uintptr
	_ = v1
	if _grptb != libc.UintptrFromInt32(0) {
		return 0
	}
	if _fail3 != 0 {
		return -int32(1)
	}
	v1 = libc.Xcalloc(tls, uint64(m_GNM_SZ), uint64(8))
	_grptb = v1
	if v1 == libc.UintptrFromInt32(0) {
		_fail3++
		return -int32(1)
	}
	return 0
}

var _fail3 int32

// C documentation
//
//	/*
//	 * user_from_uid()
//	 *	caches the name (if any) for the uid. If noname clear, we always
//	 *	return the stored name (if valid or invalid match).
//	 *	We use a simple hash table.
//	 * Return:
//	 *	Pointer to stored name (or a empty string)
//	 */
func Xuser_from_uid(tls *libc.TLS, uid Tuid_t, noname int32) (r uintptr) {
	bp := tls.Alloc(1104)
	defer tls.Free(1104)
	var pptr, ptr, v1 uintptr
	var _ /* pw at bp+48 */ uintptr
	var _ /* pwbuf at bp+56 */ [1024]int8
	var _ /* pwstore at bp+0 */ Tpasswd
	_, _, _ = pptr, ptr, v1
	*(*uintptr)(unsafe.Pointer(bp + 48)) = libc.UintptrFromInt32(0)
	ptr = libc.UintptrFromInt32(0)
	if _uidtb != libc.UintptrFromInt32(0) || _uidtb_start(tls) == 0 {
		/*
		 * see if we have this uid cached
		 */
		pptr = _uidtb + uintptr(uid%libc.Uint32FromInt32(m_UID_SZ))*8
		ptr = *(*uintptr)(unsafe.Pointer(pptr))
		if ptr != libc.UintptrFromInt32(0) && (*TUIDC)(unsafe.Pointer(ptr)).Fvalid > 0 && (*TUIDC)(unsafe.Pointer(ptr)).Fuid == uid {
			/*
			 * have an entry for this uid
			 */
			if !(noname != 0) || (*TUIDC)(unsafe.Pointer(ptr)).Fvalid == int32(m_VALID) {
				return ptr + 4
			}
			return libc.UintptrFromInt32(0)
		}
		if ptr == libc.UintptrFromInt32(0) {
			v1 = libc.Xmalloc(tls, uint64(40))
			ptr = v1
			*(*uintptr)(unsafe.Pointer(pptr)) = v1
		}
	}
	libc.Xgetpwuid_r(tls, uid, bp, bp+56, uint64(1024), bp+48)
	if *(*uintptr)(unsafe.Pointer(bp + 48)) == libc.UintptrFromInt32(0) {
		/*
		 * no match for this uid in the local password file
		 * a string that is the uid in numeric format
		 */
		if ptr == libc.UintptrFromInt32(0) {
			return libc.UintptrFromInt32(0)
		}
		(*TUIDC)(unsafe.Pointer(ptr)).Fuid = uid
		libc.X__builtin_snprintf(tls, ptr+4, uint64(m_UNMLEN), __ccgo_ts+326, libc.VaList(bp+1088, uid))
		(*TUIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_INVALID)
		if noname != 0 {
			return libc.UintptrFromInt32(0)
		}
	} else {
		/*
		 * there is an entry for this uid in the password file
		 */
		if ptr == libc.UintptrFromInt32(0) {
			return (*Tpasswd)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 48)))).Fpw_name
		}
		(*TUIDC)(unsafe.Pointer(ptr)).Fuid = uid
		Xstrlcpy(tls, ptr+4, (*Tpasswd)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 48)))).Fpw_name, uint64(32))
		(*TUIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_VALID)
	}
	return ptr + 4
}

// C documentation
//
//	/*
//	 * group_from_gid()
//	 *	caches the name (if any) for the gid. If noname clear, we always
//	 *	return the stored name (if valid or invalid match).
//	 *	We use a simple hash table.
//	 * Return:
//	 *	Pointer to stored name (or a empty string)
//	 */
func Xgroup_from_gid(tls *libc.TLS, gid Tgid_t, noname int32) (r uintptr) {
	bp := tls.Alloc(2688)
	defer tls.Free(2688)
	var pptr, ptr, v1 uintptr
	var _ /* gr at bp+32 */ uintptr
	var _ /* grbuf at bp+40 */ [2624]int8
	var _ /* grstore at bp+0 */ Tgroup
	_, _, _ = pptr, ptr, v1
	*(*uintptr)(unsafe.Pointer(bp + 32)) = libc.UintptrFromInt32(0)
	ptr = libc.UintptrFromInt32(0)
	if _gidtb != libc.UintptrFromInt32(0) || _gidtb_start(tls) == 0 {
		/*
		 * see if we have this gid cached
		 */
		pptr = _gidtb + uintptr(gid%libc.Uint32FromInt32(m_GID_SZ))*8
		ptr = *(*uintptr)(unsafe.Pointer(pptr))
		if ptr != libc.UintptrFromInt32(0) && (*TGIDC)(unsafe.Pointer(ptr)).Fvalid > 0 && (*TGIDC)(unsafe.Pointer(ptr)).Fgid == gid {
			/*
			 * have an entry for this gid
			 */
			if !(noname != 0) || (*TGIDC)(unsafe.Pointer(ptr)).Fvalid == int32(m_VALID) {
				return ptr + 4
			}
			return libc.UintptrFromInt32(0)
		}
		if ptr == libc.UintptrFromInt32(0) {
			v1 = libc.Xmalloc(tls, uint64(40))
			ptr = v1
			*(*uintptr)(unsafe.Pointer(pptr)) = v1
		}
	}
	libc.Xgetgrgid_r(tls, gid, bp, bp+40, uint64(2624), bp+32)
	if *(*uintptr)(unsafe.Pointer(bp + 32)) == libc.UintptrFromInt32(0) {
		/*
		 * no match for this gid in the local group file, put in
		 * a string that is the gid in numeric format
		 */
		if ptr == libc.UintptrFromInt32(0) {
			return libc.UintptrFromInt32(0)
		}
		(*TGIDC)(unsafe.Pointer(ptr)).Fgid = gid
		libc.X__builtin_snprintf(tls, ptr+4, uint64(m_GNMLEN), __ccgo_ts+326, libc.VaList(bp+2672, gid))
		(*TGIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_INVALID)
		if noname != 0 {
			return libc.UintptrFromInt32(0)
		}
	} else {
		/*
		 * there is an entry for this group in the group file
		 */
		if ptr == libc.UintptrFromInt32(0) {
			return (*Tgroup)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Fgr_name
		}
		(*TGIDC)(unsafe.Pointer(ptr)).Fgid = gid
		Xstrlcpy(tls, ptr+4, (*Tgroup)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Fgr_name, uint64(32))
		(*TGIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_VALID)
	}
	return ptr + 4
}

// C documentation
//
//	/*
//	 * uid_from_user()
//	 *	caches the uid for a given user name. We use a simple hash table.
//	 * Return:
//	 *	0 if the user name is found (filling in uid), -1 otherwise
//	 */
func Xuid_from_user(tls *libc.TLS, name uintptr, uid uintptr) (r int32) {
	bp := tls.Alloc(1088)
	defer tls.Free(1088)
	var namelen, v1 Tsize_t
	var pptr, ptr, v3 uintptr
	var v2 bool
	var v4 Tuid_t
	var _ /* pw at bp+48 */ uintptr
	var _ /* pwbuf at bp+56 */ [1024]int8
	var _ /* pwstore at bp+0 */ Tpasswd
	_, _, _, _, _, _, _ = namelen, pptr, ptr, v1, v2, v3, v4
	*(*uintptr)(unsafe.Pointer(bp + 48)) = libc.UintptrFromInt32(0)
	ptr = libc.UintptrFromInt32(0)
	/*
	 * return -1 for mangled names
	 */
	if v2 = name == libc.UintptrFromInt32(0); !v2 {
		v1 = libc.Xstrlen(tls, name)
		namelen = v1
	}
	if v2 || v1 == uint64(0) {
		return -int32(1)
	}
	if _usrtb != libc.UintptrFromInt32(0) || _usrtb_start(tls) == 0 {
		/*
		 * look up in hash table, if found and valid return the uid,
		 * if found and invalid, return a -1
		 */
		pptr = _usrtb + uintptr(_st_hash(tls, name, namelen, int32(m_UNM_SZ)))*8
		ptr = *(*uintptr)(unsafe.Pointer(pptr))
		if ptr != libc.UintptrFromInt32(0) && (*TUIDC)(unsafe.Pointer(ptr)).Fvalid > 0 && libc.Xstrcmp(tls, name, ptr+4) == 0 {
			if (*TUIDC)(unsafe.Pointer(ptr)).Fvalid == int32(m_INVALID) {
				return -int32(1)
			}
			*(*Tuid_t)(unsafe.Pointer(uid)) = (*TUIDC)(unsafe.Pointer(ptr)).Fuid
			return 0
		}
		if ptr == libc.UintptrFromInt32(0) {
			v3 = libc.Xmalloc(tls, uint64(40))
			ptr = v3
			*(*uintptr)(unsafe.Pointer(pptr)) = v3
		}
	}
	/*
	 * no match, look it up, if no match store it as an invalid entry,
	 * or store the matching uid
	 */
	libc.Xgetpwnam_r(tls, name, bp, bp+56, uint64(1024), bp+48)
	if ptr == libc.UintptrFromInt32(0) {
		if *(*uintptr)(unsafe.Pointer(bp + 48)) == libc.UintptrFromInt32(0) {
			return -int32(1)
		}
		*(*Tuid_t)(unsafe.Pointer(uid)) = (*Tpasswd)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 48)))).Fpw_uid
		return 0
	}
	Xstrlcpy(tls, ptr+4, name, uint64(32))
	if *(*uintptr)(unsafe.Pointer(bp + 48)) == libc.UintptrFromInt32(0) {
		(*TUIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_INVALID)
		return -int32(1)
	}
	(*TUIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_VALID)
	v4 = (*Tpasswd)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 48)))).Fpw_uid
	(*TUIDC)(unsafe.Pointer(ptr)).Fuid = v4
	*(*Tuid_t)(unsafe.Pointer(uid)) = v4
	return 0
}

// C documentation
//
//	/*
//	 * gid_from_group()
//	 *	caches the gid for a given group name. We use a simple hash table.
//	 * Return:
//	 *	0 if the group name is found (filling in gid), -1 otherwise
//	 */
func Xgid_from_group(tls *libc.TLS, name uintptr, gid uintptr) (r int32) {
	bp := tls.Alloc(2672)
	defer tls.Free(2672)
	var namelen, v1 Tsize_t
	var pptr, ptr, v3 uintptr
	var v2 bool
	var v4 Tgid_t
	var _ /* gr at bp+32 */ uintptr
	var _ /* grbuf at bp+40 */ [2624]int8
	var _ /* grstore at bp+0 */ Tgroup
	_, _, _, _, _, _, _ = namelen, pptr, ptr, v1, v2, v3, v4
	*(*uintptr)(unsafe.Pointer(bp + 32)) = libc.UintptrFromInt32(0)
	ptr = libc.UintptrFromInt32(0)
	/*
	 * return -1 for mangled names
	 */
	if v2 = name == libc.UintptrFromInt32(0); !v2 {
		v1 = libc.Xstrlen(tls, name)
		namelen = v1
	}
	if v2 || v1 == uint64(0) {
		return -int32(1)
	}
	if _grptb != libc.UintptrFromInt32(0) || _grptb_start(tls) == 0 {
		/*
		 * look up in hash table, if found and valid return the uid,
		 * if found and invalid, return a -1
		 */
		pptr = _grptb + uintptr(_st_hash(tls, name, namelen, int32(m_GID_SZ)))*8
		ptr = *(*uintptr)(unsafe.Pointer(pptr))
		if ptr != libc.UintptrFromInt32(0) && (*TGIDC)(unsafe.Pointer(ptr)).Fvalid > 0 && libc.Xstrcmp(tls, name, ptr+4) == 0 {
			if (*TGIDC)(unsafe.Pointer(ptr)).Fvalid == int32(m_INVALID) {
				return -int32(1)
			}
			*(*Tgid_t)(unsafe.Pointer(gid)) = (*TGIDC)(unsafe.Pointer(ptr)).Fgid
			return 0
		}
		if ptr == libc.UintptrFromInt32(0) {
			v3 = libc.Xmalloc(tls, uint64(40))
			ptr = v3
			*(*uintptr)(unsafe.Pointer(pptr)) = v3
		}
	}
	/*
	 * no match, look it up, if no match store it as an invalid entry,
	 * or store the matching gid
	 */
	libc.Xgetgrnam_r(tls, name, bp, bp+40, uint64(2624), bp+32)
	if ptr == libc.UintptrFromInt32(0) {
		if *(*uintptr)(unsafe.Pointer(bp + 32)) == libc.UintptrFromInt32(0) {
			return -int32(1)
		}
		*(*Tgid_t)(unsafe.Pointer(gid)) = (*Tgroup)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Fgr_gid
		return 0
	}
	Xstrlcpy(tls, ptr+4, name, uint64(32))
	if *(*uintptr)(unsafe.Pointer(bp + 32)) == libc.UintptrFromInt32(0) {
		(*TGIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_INVALID)
		return -int32(1)
	}
	(*TGIDC)(unsafe.Pointer(ptr)).Fvalid = int32(m_VALID)
	v4 = (*Tgroup)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Fgr_gid
	(*TGIDC)(unsafe.Pointer(ptr)).Fgid = v4
	*(*Tgid_t)(unsafe.Pointer(gid)) = v4
	return 0
}

const m_B0 = 0000000
const m_B1000000 = 0010010
const m_B110 = 0000003
const m_B115200 = 0010002
const m_B1152000 = 0010011
const m_B1200 = 0000011
const m_B134 = 0000004
const m_B150 = 0000005
const m_B1500000 = 0010012
const m_B1800 = 0000012
const m_B19200 = 0000016
const m_B200 = 0000006
const m_B2000000 = 0010013
const m_B230400 = 0010003
const m_B2400 = 0000013
const m_B2500000 = 0010014
const m_B300 = 0000007
const m_B3000000 = 0010015
const m_B3500000 = 0010016
const m_B38400 = 0000017
const m_B4000000 = 0010017
const m_B460800 = 0010004
const m_B4800 = 0000014
const m_B50 = 0000001
const m_B500000 = 0010005
const m_B57600 = 0010001
const m_B576000 = 0010006
const m_B600 = 0000010
const m_B75 = 0000002
const m_B921600 = 0010007
const m_B9600 = 0000015
const m_BRKINT = 0000002
const m_BS0 = 0000000
const m_BS1 = 0020000
const m_BSDLY = 0020000
const m_CBAUD = 0010017
const m_CBAUDEX = 0010000
const m_CIBAUD = 002003600000
const m_CLOCAL = 0004000
const m_CMSPAR = 010000000000
const m_CR0 = 0000000
const m_CR1 = 0001000
const m_CR2 = 0002000
const m_CR3 = 0003000
const m_CRDLY = 0003000
const m_CREAD = 0000200
const m_CRTSCTS = 020000000000
const m_CS5 = 0000000
const m_CS6 = 0000020
const m_CS7 = 0000040
const m_CS8 = 0000060
const m_CSIZE = 0000060
const m_CSTOPB = 0000100
const m_ECHO = 8
const m_ECHOCTL = 0001000
const m_ECHOE = 0000020
const m_ECHOK = 0000040
const m_ECHOKE = 0004000
const m_ECHONL = 64
const m_ECHOPRT = 0002000
const m_EXTA = 0000016
const m_EXTB = 0000017
const m_EXTPROC = 0200000
const m_FF0 = 0000000
const m_FF1 = 0100000
const m_FFDLY = 0100000
const m_FLUSHO = 0010000
const m_HUPCL = 0002000
const m_ICANON = 0000002
const m_ICRNL = 0000400
const m_IEXTEN = 0100000
const m_IGNBRK = 0000001
const m_IGNCR = 0000200
const m_IGNPAR = 0000004
const m_IMAXBEL = 0020000
const m_INLCR = 0000100
const m_INPCK = 0000020
const m_ISIG = 0000001
const m_ISTRIP = 0000040
const m_IUCLC = 0001000
const m_IUTF8 = 0040000
const m_IXANY = 0004000
const m_IXOFF = 0010000
const m_IXON = 0002000
const m_NCCS = 32
const m_NL0 = 0000000
const m_NL1 = 0000400
const m_NLDLY = 0000400
const m_NOFLSH = 0000200
const m_OCRNL = 0000010
const m_OFDEL = 0000200
const m_OFILL = 0000100
const m_OLCUC = 0000002
const m_ONLCR = 0000004
const m_ONLRET = 0000040
const m_ONOCR = 0000020
const m_OPOST = 0000001
const m_O_CREAT4 = 0100
const m_O_NONBLOCK4 = 04000
const m_O_RDONLY2 = 00
const m_O_RDWR1 = 2
const m_O_TRUNC4 = 01000
const m_O_WRONLY2 = 01
const m_PARENB = 0000400
const m_PARMRK = 0000010
const m_PARODD = 0001000
const m_PENDIN = 0040000
const m_RPP_ECHO_OFF = 0x00
const m_RPP_ECHO_ON = 1
const m_RPP_FORCELOWER = 4
const m_RPP_FORCEUPPER = 8
const m_RPP_REQUIRE_TTY = 2
const m_RPP_SEVENBIT = 16
const m_RPP_STDIN = 32
const m_TAB0 = 0000000
const m_TAB1 = 0004000
const m_TAB2 = 0010000
const m_TAB3 = 0014000
const m_TABDLY = 0014000
const m_TCIFLUSH = 0
const m_TCIOFF = 2
const m_TCIOFLUSH = 2
const m_TCION = 3
const m_TCOFLUSH = 1
const m_TCOOFF = 0
const m_TCOON = 1
const m_TCSADRAIN = 1
const m_TCSAFLUSH = 2
const m_TCSANOW = 0
const m_TCSASOFT = 0
const m_TOSTOP = 0000400
const m_VDISCARD = 13
const m_VEOF = 4
const m_VEOL = 11
const m_VEOL2 = 16
const m_VERASE = 2
const m_VINTR = 0
const m_VKILL = 3
const m_VLNEXT = 15
const m_VMIN = 6
const m_VQUIT = 1
const m_VREPRINT = 12
const m_VSTART = 8
const m_VSTOP = 9
const m_VSUSP = 10
const m_VSWTC = 7
const m_VT0 = 0000000
const m_VT1 = 0040000
const m_VTDLY = 0040000
const m_VTIME = 5
const m_VWERASE = 14
const m_XCASE = 0000004
const m_XTABS = 0014000
const m__PATH_BSHELL = "/bin/sh"
const m__PATH_CONSOLE = "/dev/console"
const m__PATH_DEFPATH = "/usr/local/bin:/bin:/usr/bin"
const m__PATH_DEV = "/dev/"
const m__PATH_DEVNULL = "/dev/null"
const m__PATH_KLOG = "/proc/kmsg"
const m__PATH_LASTLOG = "/var/log/lastlog"
const m__PATH_MAILDIR = "/var/mail"
const m__PATH_MAN = "/usr/share/man"
const m__PATH_MNTTAB = "/etc/fstab"
const m__PATH_MOUNTED = "/etc/mtab"
const m__PATH_NOLOGIN = "/etc/nologin"
const m__PATH_SENDMAIL = "/usr/sbin/sendmail"
const m__PATH_SHADOW = "/etc/shadow"
const m__PATH_SHELLS = "/etc/shells"
const m__PATH_STDPATH = "/bin:/usr/bin:/sbin:/usr/sbin"
const m__PATH_TMP = "/tmp/"
const m__PATH_TTY = "/dev/tty"
const m__PATH_UTMP = "/dev/null/utmp"
const m__PATH_VARDB = "/var/lib/misc/"
const m__PATH_VARRUN = "/var/run/"
const m__PATH_VARTMP = "/var/tmp/"
const m__PATH_VI = "/usr/bin/vi"
const m__PATH_WTMP = "/dev/null/wtmp"

type Twinsize = struct {
	Fws_row    uint16
	Fws_col    uint16
	Fws_xpixel uint16
	Fws_ypixel uint16
}

type Tcc_t = uint8

type Tspeed_t = uint32

type Ttcflag_t = uint32

type Ttermios = struct {
	Fc_iflag    Ttcflag_t
	Fc_oflag    Ttcflag_t
	Fc_cflag    Ttcflag_t
	Fc_lflag    Ttcflag_t
	Fc_line     Tcc_t
	Fc_cc       [32]Tcc_t
	F__c_ispeed Tspeed_t
	F__c_ospeed Tspeed_t
}

var _signo [65]Tsig_atomic_t

func Xreadpassphrase(tls *libc.TLS, prompt uintptr, buf uintptr, bufsiz Tsize_t, flags int32) (r uintptr) {
	bp := tls.Alloc(1648)
	defer tls.Free(1648)
	var end, p, v6, v8 uintptr
	var i, input, need_restart, output, save_errno, sigttou, v2, v3 int32
	var nr, v5 Tssize_t
	var v4 bool
	var _ /* ch at bp+0 */ int8
	var _ /* oterm at bp+64 */ Ttermios
	var _ /* sa at bp+128 */ Tsigaction
	var _ /* savealrm at bp+280 */ Tsigaction
	var _ /* savehup at bp+584 */ Tsigaction
	var _ /* saveint at bp+432 */ Tsigaction
	var _ /* savepipe at bp+1496 */ Tsigaction
	var _ /* savequit at bp+736 */ Tsigaction
	var _ /* saveterm at bp+888 */ Tsigaction
	var _ /* savetstp at bp+1040 */ Tsigaction
	var _ /* savettin at bp+1192 */ Tsigaction
	var _ /* savettou at bp+1344 */ Tsigaction
	var _ /* term at bp+4 */ Ttermios
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = end, i, input, need_restart, nr, output, p, save_errno, sigttou, v2, v3, v4, v5, v6, v8
	/* I suppose we could alloc on demand in this case (XXX). */
	if bufsiz == uint64(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return libc.UintptrFromInt32(0)
	}
	goto restart
restart:
	;
	i = 0
	for {
		if !(i < int32(m__NSIG)) {
			break
		}
		libc.AtomicStorePInt32(uintptr(unsafe.Pointer(&_signo))+uintptr(i)*4, 0)
		goto _1
	_1:
		;
		i++
	}
	nr = int64(-int32(1))
	save_errno = 0
	need_restart = 0
	/*
	 * Read and write to /dev/tty if available.  If not, read from
	 * stdin and write to stderr unless a tty is required.
	 */
	if v4 = flags&int32(m_RPP_STDIN) != 0; !v4 {
		v3 = libc.Xopen(tls, __ccgo_ts+430, int32(m_O_RDWR1), 0)
		output = v3
		v2 = v3
		input = v2
	}
	if v4 || v2 == -int32(1) {
		if flags&int32(m_RPP_REQUIRE_TTY) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOTTY)
			return libc.UintptrFromInt32(0)
		}
		input = m_STDIN_FILENO
		output = int32(m_STDERR_FILENO)
	}
	/*
	 * Turn off echo if possible.
	 * If we are using a tty but are not the foreground pgrp this will
	 * generate SIGTTOU, so do it *before* installing the signal handlers.
	 */
	if input != m_STDIN_FILENO && libc.Xtcgetattr(tls, input, bp+64) == 0 {
		libc.Xmemcpy(tls, bp+4, bp+64, uint64(60))
		if !(flags&libc.Int32FromInt32(m_RPP_ECHO_ON) != 0) {
			(*(*Ttermios)(unsafe.Pointer(bp + 4))).Fc_lflag &= libc.Uint32FromInt32(^(libc.Int32FromInt32(m_ECHO) | libc.Int32FromInt32(m_ECHONL)))
		}
		libc.Xtcsetattr(tls, input, libc.Int32FromInt32(m_TCSAFLUSH)|libc.Int32FromInt32(m_TCSASOFT), bp+4)
	} else {
		libc.Xmemset(tls, bp+4, 0, uint64(60))
		(*(*Ttermios)(unsafe.Pointer(bp + 4))).Fc_lflag |= uint32(m_ECHO)
		libc.Xmemset(tls, bp+64, 0, uint64(60))
		(*(*Ttermios)(unsafe.Pointer(bp + 64))).Fc_lflag |= uint32(m_ECHO)
	}
	/*
	 * Catch signals that would otherwise cause the user to end
	 * up with echo turned off in the shell.  Don't worry about
	 * things like SIGXCPU and SIGVTALRM for now.
	 */
	libc.Xsigemptyset(tls, bp+128+8)
	(*(*Tsigaction)(unsafe.Pointer(bp + 128))).Fsa_flags = 0 /* don't restart system calls */
	*(*uintptr)(unsafe.Pointer(bp + 128)) = __ccgo_fp(_handler)
	libc.Xsigaction(tls, int32(m_SIGALRM), bp+128, bp+280)
	libc.Xsigaction(tls, int32(m_SIGHUP), bp+128, bp+584)
	libc.Xsigaction(tls, int32(m_SIGINT), bp+128, bp+432)
	libc.Xsigaction(tls, int32(m_SIGPIPE), bp+128, bp+1496)
	libc.Xsigaction(tls, int32(m_SIGQUIT), bp+128, bp+736)
	libc.Xsigaction(tls, int32(m_SIGTERM), bp+128, bp+888)
	libc.Xsigaction(tls, int32(m_SIGTSTP), bp+128, bp+1040)
	libc.Xsigaction(tls, int32(m_SIGTTIN), bp+128, bp+1192)
	libc.Xsigaction(tls, int32(m_SIGTTOU), bp+128, bp+1344)
	if !(flags&libc.Int32FromInt32(m_RPP_STDIN) != 0) {
		libc.Xwrite(tls, output, prompt, libc.Xstrlen(tls, prompt))
	}
	end = buf + uintptr(bufsiz) - uintptr(1)
	p = buf
	for {
		v5 = libc.Xread(tls, input, bp, uint64(1))
		nr = v5
		if !(v5 == int64(1) && int32(*(*int8)(unsafe.Pointer(bp))) != int32('\n') && int32(*(*int8)(unsafe.Pointer(bp))) != int32('\r')) {
			break
		}
		if p < end {
			if flags&int32(m_RPP_SEVENBIT) != 0 {
				*(*int8)(unsafe.Pointer(bp)) = int8(int32(*(*int8)(unsafe.Pointer(bp))) & libc.Int32FromInt32(0x7f))
			}
			if libc.BoolInt32(uint32(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(bp))))|uint32(32)-uint32('a') < uint32(26)) != 0 {
				if flags&int32(m_RPP_FORCELOWER) != 0 {
					*(*int8)(unsafe.Pointer(bp)) = int8(libc.Xtolower(tls, libc.Int32FromUint8(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(bp))))))
				}
				if flags&int32(m_RPP_FORCEUPPER) != 0 {
					*(*int8)(unsafe.Pointer(bp)) = int8(libc.Xtoupper(tls, libc.Int32FromUint8(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(bp))))))
				}
			}
			v6 = p
			p++
			*(*int8)(unsafe.Pointer(v6)) = *(*int8)(unsafe.Pointer(bp))
		}
	}
	*(*int8)(unsafe.Pointer(p)) = int8('\000')
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	if !((*(*Ttermios)(unsafe.Pointer(bp + 4))).Fc_lflag&libc.Uint32FromInt32(m_ECHO) != 0) {
		libc.Xwrite(tls, output, __ccgo_ts+439, uint64(1))
	}
	/* Restore old terminal settings and signals. */
	if libc.Xmemcmp(tls, bp+4, bp+64, uint64(60)) != 0 {
		sigttou = libc.AtomicLoadPInt32(uintptr(unsafe.Pointer(&_signo)) + libc.UintptrFromInt32(m_SIGTTOU)*4)
		/* Ignore SIGTTOU generated when we are not the fg pgrp. */
		for libc.Xtcsetattr(tls, input, libc.Int32FromInt32(m_TCSAFLUSH)|libc.Int32FromInt32(m_TCSASOFT), bp+64) == -int32(1) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EINTR) && !(libc.AtomicLoadPInt32(uintptr(unsafe.Pointer(&_signo))+libc.UintptrFromInt32(m_SIGTTOU)*4) != 0) {
			continue
		}
		libc.AtomicStorePInt32(uintptr(unsafe.Pointer(&_signo))+22*4, sigttou)
	}
	libc.Xsigaction(tls, int32(m_SIGALRM), bp+280, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGHUP), bp+584, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGINT), bp+432, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGQUIT), bp+736, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGPIPE), bp+1496, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGTERM), bp+888, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGTSTP), bp+1040, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGTTIN), bp+1192, libc.UintptrFromInt32(0))
	libc.Xsigaction(tls, int32(m_SIGTTOU), bp+1344, libc.UintptrFromInt32(0))
	if input != m_STDIN_FILENO {
		libc.Xclose(tls, input)
	}
	/*
	 * If we were interrupted by a signal, resend it to ourselves
	 * now that we have restored the signal handlers.
	 */
	i = 0
	for {
		if !(i < int32(m__NSIG)) {
			break
		}
		if libc.AtomicLoadPInt32(uintptr(unsafe.Pointer(&_signo))+uintptr(i)*4) != 0 {
			libc.Xkill(tls, libc.Xgetpid(tls), i)
			switch i {
			case int32(m_SIGTSTP):
				fallthrough
			case int32(m_SIGTTIN):
				fallthrough
			case int32(m_SIGTTOU):
				need_restart = int32(1)
			}
		}
		goto _7
	_7:
		;
		i++
	}
	if need_restart != 0 {
		goto restart
	}
	if save_errno != 0 {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	}
	if nr == int64(-int32(1)) {
		v8 = libc.UintptrFromInt32(0)
	} else {
		v8 = buf
	}
	return v8
}

func _handler(tls *libc.TLS, s int32) {
	libc.AtomicStorePInt32(uintptr(unsafe.Pointer(&_signo))+uintptr(s)*4, int32(1))
}

const m_UINT64_MAX1 = 18446744073709551615

/*
 * This is sqrt(SIZE_MAX+1), as s1*s2 <= SIZE_MAX
 * if both s1 < MUL_NO_OVERFLOW and s2 < MUL_NO_OVERFLOW
 */

func Xreallocarray(tls *libc.TLS, optr uintptr, nmemb Tsize_t, size Tsize_t) (r uintptr) {
	if (nmemb >= libc.Uint64FromInt32(1)<<(libc.Uint64FromInt64(8)*libc.Uint64FromInt32(4)) || size >= libc.Uint64FromInt32(1)<<(libc.Uint64FromInt64(8)*libc.Uint64FromInt32(4))) && nmemb > uint64(0) && uint64(0xffffffffffffffff)/nmemb < size {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOMEM)
		return libc.UintptrFromInt32(0)
	}
	return libc.Xrealloc(tls, optr, size*nmemb)
}

const m_UINT64_MAX2 = "0xffffffffffffffffu"

func Xreallocf(tls *libc.TLS, ptr uintptr, size Tsize_t) (r uintptr) {
	var nptr uintptr
	_ = nptr
	nptr = libc.Xrealloc(tls, ptr, size)
	/*
	 * When the System V compatibility option (malloc "V" flag) is
	 * in effect, realloc(ptr, 0) frees the memory and returns NULL.
	 * So, to avoid double free, call free() only when size != 0.
	 * realloc(ptr, 0) can't fail when ptr != NULL.
	 */
	if !(nptr != 0) && ptr != 0 && size != uint64(0) {
		libc.Xfree(tls, ptr)
	}
	return nptr
}

const m_UINT64_MAX3 = 18446744073709551615

/*
 * This is sqrt(SIZE_MAX+1), as s1*s2 <= SIZE_MAX
 * if both s1 < MUL_NO_OVERFLOW and s2 < MUL_NO_OVERFLOW
 */

func Xrecallocarray(tls *libc.TLS, ptr uintptr, oldnmemb Tsize_t, newnmemb Tsize_t, size Tsize_t) (r uintptr) {
	var d, newsize, oldsize Tsize_t
	var newptr uintptr
	_, _, _, _ = d, newptr, newsize, oldsize
	if ptr == libc.UintptrFromInt32(0) {
		return libc.Xcalloc(tls, newnmemb, size)
	}
	if (newnmemb >= libc.Uint64FromInt32(1)<<(libc.Uint64FromInt64(8)*libc.Uint64FromInt32(4)) || size >= libc.Uint64FromInt32(1)<<(libc.Uint64FromInt64(8)*libc.Uint64FromInt32(4))) && newnmemb > uint64(0) && uint64(0xffffffffffffffff)/newnmemb < size {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOMEM)
		return libc.UintptrFromInt32(0)
	}
	newsize = newnmemb * size
	if (oldnmemb >= libc.Uint64FromInt32(1)<<(libc.Uint64FromInt64(8)*libc.Uint64FromInt32(4)) || size >= libc.Uint64FromInt32(1)<<(libc.Uint64FromInt64(8)*libc.Uint64FromInt32(4))) && oldnmemb > uint64(0) && uint64(0xffffffffffffffff)/oldnmemb < size {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return libc.UintptrFromInt32(0)
	}
	oldsize = oldnmemb * size
	/*
	 * Don't bother too much if we're shrinking just a bit,
	 * we do not shrink for series of small steps, oh well.
	 */
	if newsize <= oldsize {
		d = oldsize - newsize
		if d < oldsize/uint64(2) && d < libc.Uint64FromInt32(libc.Xgetpagesize(tls)) {
			libc.Xmemset(tls, ptr+uintptr(newsize), 0, d)
			return ptr
		}
	}
	newptr = libc.Xmalloc(tls, newsize)
	if newptr == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if newsize > oldsize {
		libc.Xmemcpy(tls, newptr, ptr, oldsize)
		libc.Xmemset(tls, newptr+uintptr(oldsize), 0, newsize-oldsize)
	} else {
		libc.Xmemcpy(tls, newptr, ptr, newsize)
	}
	Xexplicit_bzero(tls, ptr, oldsize)
	libc.Xfree(tls, ptr)
	return newptr
}

const m_UINT64_MAX4 = "0xffffffffffffffffu"

/*
 * Swap two areas of size number of bytes.  Although qsort(3) permits random
 * blocks of memory to be sorted, sorting pointers is almost certainly the
 * common case (and, were it not, could easily be made so).  Regardless, it
 * isn't worth optimizing; the SWAP's get sped up by the cache, and pointer
 * arithmetic gets lost in the time required for comparison function calls.
 */

/* Copy one block of size size to another. */

/*
 * Build the list into a heap, where a heap is defined such that for
 * the records K1 ... KN, Kj/2 >= Kj for 1 <= j/2 <= j <= N.
 *
 * There are two cases.  If j == nmemb, select largest of Ki and Kj.  If
 * j < nmemb, select largest of Ki, Kj and Kj+1.
 */

/*
 * Select the top of the heap and 'heapify'.  Since by far the most expensive
 * action is the call to the compar function, a considerable optimization
 * in the average case can be achieved due to the fact that k, the displaced
 * element, is usually quite small, so it would be preferable to first
 * heapify, always maintaining the invariant that the larger child is copied
 * over its parent's record.
 *
 * Then, starting from the *bottom* of the heap, finding k's correct place,
 * again maintaining the invariant.  As a result of the invariant no element
 * is 'lost' when k is assigned its correct place in the heap.
 *
 * The time savings from this optimization are on the order of 15-20% for the
 * average case. See Knuth, Vol. 3, page 158, problem 18.
 *
 * XXX Don't break the #define SELECT line, below.  Reiser cpp gets upset.
 */

// C documentation
//
//	/*
//	 * Heapsort -- Knuth, Vol. 3, page 145.  Runs in O (N lg N), both average
//	 * and worst.  While heapsort is faster than the worst case of quicksort,
//	 * the BSD quicksort does median selection so that the chance of finding
//	 * a data set that will trigger the worst case is nonexistent.  Heapsort's
//	 * only advantage over quicksort is that it requires little additional memory.
//	 */
func Xheapsort(tls *libc.TLS, vbase uintptr, nmemb Tsize_t, size Tsize_t, compar uintptr) (r int32) {
	var base, k, p, t, tmp1, tmp2, v1, v12, v13, v16, v17, v22, v23, v27, v28, v31, v32, v8, v9 uintptr
	var cnt, i, j, l, v10, v14, v19, v20, v25, v29, v3, v5, v6 Tsize_t
	var tmp int8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, cnt, i, j, k, l, p, t, tmp, tmp1, tmp2, v1, v10, v12, v13, v14, v16, v17, v19, v20, v22, v23, v25, v27, v28, v29, v3, v31, v32, v5, v6, v8, v9
	if nmemb <= uint64(1) {
		return 0
	}
	if !(size != 0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	}
	v1 = libc.Xmalloc(tls, size)
	k = v1
	if v1 == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	/*
	 * Items are numbered from 1 to nmemb, so offset from size bytes
	 * below the starting address.
	 */
	base = vbase - uintptr(size)
	l = nmemb/uint64(2) + uint64(1)
	for {
		l--
		v3 = l
		if !(v3 != 0) {
			break
		}
		i = l
		for {
			v5 = i * libc.Uint64FromInt32(2)
			j = v5
			if !(v5 <= nmemb) {
				break
			}
			p = base + uintptr(j*size)
			if j < nmemb && (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{compar})))(tls, p, p+uintptr(size)) < 0 {
				p += uintptr(size)
				j++
			}
			t = base + uintptr(i*size)
			if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{compar})))(tls, p, t) <= 0 {
				break
			}
			cnt = size
			for {
				tmp = *(*int8)(unsafe.Pointer(t))
				v8 = t
				t++
				*(*int8)(unsafe.Pointer(v8)) = *(*int8)(unsafe.Pointer(p))
				v9 = p
				p++
				*(*int8)(unsafe.Pointer(v9)) = tmp
				goto _7
			_7:
				;
				cnt--
				v6 = cnt
				if !(v6 != 0) {
					break
				}
			}
			goto _4
		_4:
			;
			i = j
		}
		goto _2
	_2:
	}
	/*
	 * For each element of the heap, save the largest element into its
	 * final slot, save the displaced element (k), then recreate the
	 * heap.
	 */
	for nmemb > uint64(1) {
		cnt = size
		tmp1 = k
		tmp2 = base + uintptr(nmemb*size)
		for {
			v12 = tmp1
			tmp1++
			v13 = tmp2
			tmp2++
			*(*int8)(unsafe.Pointer(v12)) = *(*int8)(unsafe.Pointer(v13))
			goto _11
		_11:
			;
			cnt--
			v10 = cnt
			if !(v10 != 0) {
				break
			}
		}
		cnt = size
		tmp1 = base + uintptr(nmemb*size)
		tmp2 = base + uintptr(size)
		for {
			v16 = tmp1
			tmp1++
			v17 = tmp2
			tmp2++
			*(*int8)(unsafe.Pointer(v16)) = *(*int8)(unsafe.Pointer(v17))
			goto _15
		_15:
			;
			cnt--
			v14 = cnt
			if !(v14 != 0) {
				break
			}
		}
		nmemb--
		i = uint64(1)
		for {
			v19 = i * libc.Uint64FromInt32(2)
			j = v19
			if !(v19 <= nmemb) {
				break
			}
			p = base + uintptr(j*size)
			if j < nmemb && (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{compar})))(tls, p, p+uintptr(size)) < 0 {
				p += uintptr(size)
				j++
			}
			t = base + uintptr(i*size)
			cnt = size
			tmp1 = t
			tmp2 = p
			for {
				v22 = tmp1
				tmp1++
				v23 = tmp2
				tmp2++
				*(*int8)(unsafe.Pointer(v22)) = *(*int8)(unsafe.Pointer(v23))
				goto _21
			_21:
				;
				cnt--
				v20 = cnt
				if !(v20 != 0) {
					break
				}
			}
			goto _18
		_18:
			;
			i = j
		}
		for {
			j = i
			i = j / uint64(2)
			p = base + uintptr(j*size)
			t = base + uintptr(i*size)
			if j == uint64(1) || (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{compar})))(tls, k, t) < 0 {
				cnt = size
				tmp1 = p
				tmp2 = k
				for {
					v27 = tmp1
					tmp1++
					v28 = tmp2
					tmp2++
					*(*int8)(unsafe.Pointer(v27)) = *(*int8)(unsafe.Pointer(v28))
					goto _26
				_26:
					;
					cnt--
					v25 = cnt
					if !(v25 != 0) {
						break
					}
				}
				break
			}
			cnt = size
			tmp1 = p
			tmp2 = t
			for {
				v31 = tmp1
				tmp1++
				v32 = tmp2
				tmp2++
				*(*int8)(unsafe.Pointer(v31)) = *(*int8)(unsafe.Pointer(v32))
				goto _30
			_30:
				;
				cnt--
				v29 = cnt
				if !(v29 != 0) {
					break
				}
			}
			goto _24
		_24:
		}
	}
	libc.Xfree(tls, k)
	return 0
}

const m_THRESHOLD = 16

/*
 * Find the next possible pointer head.  (Trickery for forcing an array
 * to do double duty as a linked list when objects do not align with word
 * boundaries.
 */
/* Assumption: PSIZE is a power of 2. */

// C documentation
//
//	/*
//	 * Arguments are as for qsort.
//	 */
func Xmergesort(tls *libc.TLS, base uintptr, nmemb Tsize_t, size Tsize_t, cmp uintptr) (r int32) {
	var b, f1, f2, l1, l2, last, list1, list2, p, p1, p2, q, t, tp2, v1, v10, v11, v12, v13, v15, v16, v18, v19, v21, v22, v24, v25, v26, v27, v28, v29, v4, v5 uintptr
	var big, iflag, sense, v2 int32
	var i, v8 Tsize_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b, big, f1, f2, i, iflag, l1, l2, last, list1, list2, p, p1, p2, q, sense, t, tp2, v1, v10, v11, v12, v13, v15, v16, v18, v19, v2, v21, v22, v24, v25, v26, v27, v28, v29, v4, v5, v8
	if size < libc.Uint64FromInt64(8)/libc.Uint64FromInt32(2) { /* Pointers must fit into 2 * size. */
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	}
	if nmemb == uint64(0) {
		return 0
	}
	/*
	 * XXX
	 * Stupid subtraction for the Cray.
	 */
	iflag = 0
	if !(size%libc.Uint64FromInt64(4) != 0) && !(libc.Uint64FromInt64(int64(base)-int64(libc.UintptrFromInt32(0)))%libc.Uint64FromInt64(4) != 0) {
		iflag = int32(1)
	}
	v1 = libc.Xmalloc(tls, nmemb*size+uint64(8))
	list2 = v1
	if v1 == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	list1 = base
	_setup(tls, list1, list2, nmemb, size, cmp)
	last = list2 + uintptr(nmemb*size)
	v2 = libc.Int32FromInt32(0)
	big = v2
	i = libc.Uint64FromInt32(v2)
	for *(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(list2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) != last {
		l2 = list1
		p1 = uintptr(libc.Uint64FromInt64(int64(list1+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1)))
		v4 = list2
		p2 = v4
		tp2 = v4
		for {
			if !(p2 != last) {
				break
			}
			p2 = *(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(p2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1)))))
			f1 = l2
			v5 = list1 + uintptr(int64(p2)-int64(list2))
			l1 = v5
			f2 = v5
			if p2 != last {
				p2 = *(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(p2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1)))))
			}
			l2 = list1 + uintptr(int64(p2)-int64(list2))
			for f1 < l1 && f2 < l2 {
				if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, f1, f2) <= 0 {
					q = f2
					b = f1
					t = l1
					sense = -int32(1)
				} else {
					q = f1
					b = f2
					t = l2
					sense = 0
				}
				if !!(big != 0) {
					goto _6
				} /* here i = 0 */
				for {
					b += uintptr(size)
					if !(b < t && (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, q, b) > sense) {
						break
					}
					i++
					v8 = i
					if v8 == uint64(6) {
						big = int32(1)
						goto EXPONENTIAL
					}
				}
				goto _7
			_6:
				;
				goto EXPONENTIAL
			EXPONENTIAL:
				;
				i = size
				for {
					v10 = b + uintptr(i)
					p = v10
					if v10 >= t {
						v11 = t - uintptr(size)
						p = v11
						if v11 > b && (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, q, p) <= sense {
							t = p
						} else {
							b = p
						}
						break
					} else {
						if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, q, p) <= sense {
							t = p
							if i == size {
								big = 0
							}
							goto FASTCASE
						} else {
							b = p
						}
					}
					goto _9
				_9:
					;
					i <<= uint64(1)
				}
				for t > b+uintptr(size) {
					i = libc.Uint64FromInt64(int64(t)-int64(b)) / size >> int32(1) * size
					v12 = b + uintptr(i)
					p = v12
					if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, q, v12) <= sense {
						t = p
					} else {
						b = p
					}
				}
				goto COPY
				goto FASTCASE
			FASTCASE:
				;
				for i > size {
					i >>= uint64(1)
					v13 = b + uintptr(i)
					p = v13
					if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, q, v13) <= sense {
						t = p
					} else {
						b = p
					}
				}
				goto COPY
			COPY:
				;
				b = t
			_7:
				;
				i = size
				if q == f1 {
					if iflag != 0 {
						for cond := true; cond; cond = f2 < b {
							*(*int32)(unsafe.Pointer(tp2)) = *(*int32)(unsafe.Pointer(f2))
							f2 += uintptr(4)
							tp2 += uintptr(4)
						}
						for {
							*(*int32)(unsafe.Pointer(tp2)) = *(*int32)(unsafe.Pointer(f1))
							f1 += uintptr(4)
							tp2 += uintptr(4)
							goto _14
						_14:
							;
							i -= uint64(4)
							if !(i != 0) {
								break
							}
						}
					} else {
						for cond := true; cond; cond = f2 < b {
							v15 = tp2
							tp2++
							v16 = f2
							f2++
							*(*uint8)(unsafe.Pointer(v15)) = *(*uint8)(unsafe.Pointer(v16))
						}
						for {
							v18 = tp2
							tp2++
							v19 = f1
							f1++
							*(*uint8)(unsafe.Pointer(v18)) = *(*uint8)(unsafe.Pointer(v19))
							goto _17
						_17:
							;
							i -= uint64(1)
							if !(i != 0) {
								break
							}
						}
					}
				} else {
					if iflag != 0 {
						for cond := true; cond; cond = f1 < b {
							*(*int32)(unsafe.Pointer(tp2)) = *(*int32)(unsafe.Pointer(f1))
							f1 += uintptr(4)
							tp2 += uintptr(4)
						}
						for {
							*(*int32)(unsafe.Pointer(tp2)) = *(*int32)(unsafe.Pointer(f2))
							f2 += uintptr(4)
							tp2 += uintptr(4)
							goto _20
						_20:
							;
							i -= uint64(4)
							if !(i != 0) {
								break
							}
						}
					} else {
						for cond := true; cond; cond = f1 < b {
							v21 = tp2
							tp2++
							v22 = f1
							f1++
							*(*uint8)(unsafe.Pointer(v21)) = *(*uint8)(unsafe.Pointer(v22))
						}
						for {
							v24 = tp2
							tp2++
							v25 = f2
							f2++
							*(*uint8)(unsafe.Pointer(v24)) = *(*uint8)(unsafe.Pointer(v25))
							goto _23
						_23:
							;
							i -= uint64(1)
							if !(i != 0) {
								break
							}
						}
					}
				}
			}
			if f2 < l2 {
				if iflag != 0 {
					for cond := true; cond; cond = f2 < l2 {
						*(*int32)(unsafe.Pointer(tp2)) = *(*int32)(unsafe.Pointer(f2))
						f2 += uintptr(4)
						tp2 += uintptr(4)
					}
				} else {
					for cond := true; cond; cond = f2 < l2 {
						v26 = tp2
						tp2++
						v27 = f2
						f2++
						*(*uint8)(unsafe.Pointer(v26)) = *(*uint8)(unsafe.Pointer(v27))
					}
				}
			} else {
				if f1 < l1 {
					if iflag != 0 {
						for cond := true; cond; cond = f1 < l1 {
							*(*int32)(unsafe.Pointer(tp2)) = *(*int32)(unsafe.Pointer(f1))
							f1 += uintptr(4)
							tp2 += uintptr(4)
						}
					} else {
						for cond := true; cond; cond = f1 < l1 {
							v28 = tp2
							tp2++
							v29 = f1
							f1++
							*(*uint8)(unsafe.Pointer(v28)) = *(*uint8)(unsafe.Pointer(v29))
						}
					}
				}
			}
			*(*uintptr)(unsafe.Pointer(p1)) = l2
			goto _3
		_3:
			;
			p1 = uintptr(libc.Uint64FromInt64(int64(l2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1)))
		}
		tp2 = list1 /* swap list1, list2 */
		list1 = list2
		list2 = tp2
		last = list2 + uintptr(nmemb*size)
	}
	if base == list2 {
		libc.Xmemmove(tls, list2, list1, nmemb*size)
		list2 = list1
	}
	libc.Xfree(tls, list2)
	return 0
}

// C documentation
//
//	/*
//	 * Optional hybrid natural/pairwise first pass.  Eats up list1 in runs of
//	 * increasing order, list2 in a corresponding linked list.  Checks for runs
//	 * when THRESHOLD/2 pairs compare with same sense.  (Only used when NATURAL
//	 * is defined.  Otherwise simple pairwise merging is used.)
//	 */
func _setup(tls *libc.TLS, list1 uintptr, list2 uintptr, n Tsize_t, size Tsize_t, cmp uintptr) {
	var f1, f2, l2, last, p2, s, v10, v13, v14, v17, v18, v19, v20, v4, v7, v8 uintptr
	var i, length, sense, size2, tmp, v11, v15, v5 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = f1, f2, i, l2, last, length, p2, s, sense, size2, tmp, v10, v11, v13, v14, v15, v17, v18, v19, v20, v4, v5, v7, v8
	size2 = libc.Int32FromUint64(size * uint64(2))
	if n <= uint64(5) {
		_insertionsort(tls, list1, n, size, cmp)
		*(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(list2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) = list2 + uintptr(n*size)
		return
	}
	/*
	 * Avoid running pointers out of bounds; limit n to evens
	 * for simplicity.
	 */
	i = libc.Int32FromUint64(uint64(4) + n&uint64(1))
	_insertionsort(tls, list1+uintptr((n-libc.Uint64FromInt32(i))*size), libc.Uint64FromInt32(i), size, cmp)
	last = list1 + uintptr(size*(n-libc.Uint64FromInt32(i)))
	*(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(list2+uintptr(int64(last)-int64(list1))+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) = list2 + uintptr(n*size)
	p2 = list2
	f1 = list1
	sense = libc.BoolInt32((*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, f1, f1+uintptr(size)) > 0)
	for {
		if !(f1 < last) {
			break
		}
		length = int32(2)
		/* Find pairs with same sense. */
		f2 = f1 + uintptr(size2)
		for {
			if !(f2 < last) {
				break
			}
			if libc.BoolInt32((*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, f2, f2+uintptr(size)) > 0) != sense {
				break
			}
			length += int32(2)
			goto _2
		_2:
			;
			f2 += uintptr(size2)
		}
		if length < int32(m_THRESHOLD) { /* Pairwise merge */
			for {
				v4 = uintptr(int64(f1+uintptr(size2))-int64(list1)) + list2
				*(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(p2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) = v4
				p2 = v4
				if sense > 0 {
					s = f1 + uintptr(size)
					i = libc.Int32FromUint64(size)
					for {
						tmp = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(f1)))
						v7 = f1
						f1++
						*(*uint8)(unsafe.Pointer(v7)) = *(*uint8)(unsafe.Pointer(s))
						v8 = s
						s++
						*(*uint8)(unsafe.Pointer(v8)) = libc.Uint8FromInt32(tmp)
						goto _6
					_6:
						;
						i--
						v5 = i
						if !(v5 != 0) {
							break
						}
					}
					f1 -= uintptr(size)
				}
				goto _3
			_3:
				;
				f1 += uintptr(size2)
				if !(f1 < f2) {
					break
				}
			}
		} else { /* Natural merge */
			l2 = f2
			f2 = f1 + uintptr(size2)
			for {
				if !(f2 < l2) {
					break
				}
				if libc.BoolInt32((*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, f2-uintptr(size), f2) > 0) != sense {
					v10 = uintptr(int64(f2)-int64(list1)) + list2
					*(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(p2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) = v10
					p2 = v10
					if sense > 0 {
						s = f2 - uintptr(size)
						for cond := true; cond; cond = f1 < s {
							i = libc.Int32FromUint64(size)
							for {
								tmp = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(f1)))
								v13 = f1
								f1++
								*(*uint8)(unsafe.Pointer(v13)) = *(*uint8)(unsafe.Pointer(s))
								v14 = s
								s++
								*(*uint8)(unsafe.Pointer(v14)) = libc.Uint8FromInt32(tmp)
								goto _12
							_12:
								;
								i--
								v11 = i
								if !(v11 != 0) {
									break
								}
							}
							s -= uintptr(size2)
						}
					}
					f1 = f2
				}
				goto _9
			_9:
				;
				f2 += uintptr(size2)
			}
			if sense > 0 {
				s = f2 - uintptr(size)
				for cond := true; cond; cond = f1 < s {
					i = libc.Int32FromUint64(size)
					for {
						tmp = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(f1)))
						v17 = f1
						f1++
						*(*uint8)(unsafe.Pointer(v17)) = *(*uint8)(unsafe.Pointer(s))
						v18 = s
						s++
						*(*uint8)(unsafe.Pointer(v18)) = libc.Uint8FromInt32(tmp)
						goto _16
					_16:
						;
						i--
						v15 = i
						if !(v15 != 0) {
							break
						}
					}
					s -= uintptr(size2)
				}
			}
			f1 = f2
			if f2 < last || (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, f2-uintptr(size), f2) > 0 {
				v19 = uintptr(int64(f2)-int64(list1)) + list2
				*(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(p2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) = v19
				p2 = v19
			} else {
				v20 = list2 + uintptr(n*size)
				*(*uintptr)(unsafe.Pointer(uintptr(libc.Uint64FromInt64(int64(p2+libc.UintptrFromInt64(8)-libc.UintptrFromInt32(1))-int64(libc.UintptrFromInt32(0))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))))) = v20
				p2 = v20
			}
		}
		goto _1
	_1:
		;
		sense = libc.BoolInt32(!(sense != 0))
	}
}

// C documentation
//
//	/*
//	 * This is to avoid out-of-bounds addresses in sorting the
//	 * last 4 elements.
//	 */
func _insertionsort(tls *libc.TLS, a uintptr, n Tsize_t, size Tsize_t, cmp uintptr) {
	var ai, s, t, u, v6, v7 uintptr
	var i, v4 int32
	var tmp uint8
	var v2 Tsize_t
	_, _, _, _, _, _, _, _, _, _ = ai, i, s, t, tmp, u, v2, v4, v6, v7
	ai = a + uintptr(size)
	for {
		n--
		v2 = n
		if !(v2 >= uint64(1)) {
			break
		}
		t = ai
		for {
			if !(t > a) {
				break
			}
			u = t - uintptr(size)
			if (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{cmp})))(tls, u, t) <= 0 {
				break
			}
			s = t
			i = libc.Int32FromUint64(size)
			for {
				tmp = *(*uint8)(unsafe.Pointer(u))
				v6 = u
				u++
				*(*uint8)(unsafe.Pointer(v6)) = *(*uint8)(unsafe.Pointer(s))
				v7 = s
				s++
				*(*uint8)(unsafe.Pointer(v7)) = tmp
				goto _5
			_5:
				;
				i--
				v4 = i
				if !(v4 != 0) {
					break
				}
			}
			u -= uintptr(size)
			goto _3
		_3:
			;
			t -= uintptr(size)
		}
		goto _1
	_1:
		;
		ai += uintptr(size)
	}
}

const m_SIZE = 512
const m_THRESHOLD1 = 20

type Tstack = struct {
	Fsa uintptr
	Fsn int32
	Fsi int32
}

func Xradixsort(tls *libc.TLS, a uintptr, n int32, tab uintptr, endch uint32) (r int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var c uint32
	var tr uintptr
	var _ /* tr0 at bp+0 */ [256]uint8
	_, _ = c, tr
	if tab == libc.UintptrFromInt32(0) {
		tr = bp
		c = uint32(0)
		for {
			if !(c < endch) {
				break
			}
			(*(*[256]uint8)(unsafe.Pointer(bp)))[c] = uint8(c + uint32(1))
			goto _1
		_1:
			;
			c++
		}
		(*(*[256]uint8)(unsafe.Pointer(bp)))[c] = uint8(0)
		c++
		for {
			if !(c < uint32(256)) {
				break
			}
			(*(*[256]uint8)(unsafe.Pointer(bp)))[c] = uint8(c)
			goto _2
		_2:
			;
			c++
		}
		endch = uint32(0)
	} else {
		endch = uint32(*(*uint8)(unsafe.Pointer(tab + uintptr(endch))))
		tr = tab
		if endch != uint32(0) && endch != uint32(255) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
			return -int32(1)
		}
	}
	_r_sort_a(tls, a, n, 0, tr, endch)
	return 0
}

func Xsradixsort(tls *libc.TLS, a uintptr, n int32, tab uintptr, endch uint32) (r int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var c uint32
	var ta, tr uintptr
	var _ /* tr0 at bp+0 */ [256]uint8
	_, _, _ = c, ta, tr
	if a == libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EFAULT)
		return -int32(1)
	}
	if tab == libc.UintptrFromInt32(0) {
		tr = bp
		c = uint32(0)
		for {
			if !(c < endch) {
				break
			}
			(*(*[256]uint8)(unsafe.Pointer(bp)))[c] = uint8(c + uint32(1))
			goto _1
		_1:
			;
			c++
		}
		(*(*[256]uint8)(unsafe.Pointer(bp)))[c] = uint8(0)
		c++
		for {
			if !(c < uint32(256)) {
				break
			}
			(*(*[256]uint8)(unsafe.Pointer(bp)))[c] = uint8(c)
			goto _2
		_2:
			;
			c++
		}
		endch = uint32(0)
	} else {
		endch = uint32(*(*uint8)(unsafe.Pointer(tab + uintptr(endch))))
		tr = tab
		if endch != uint32(0) && endch != uint32(255) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
			return -int32(1)
		}
	}
	if n < int32(m_THRESHOLD1) {
		_simplesort(tls, a, n, 0, tr, endch)
	} else {
		ta = Xreallocarray(tls, libc.UintptrFromInt32(0), libc.Uint64FromInt32(n), uint64(8))
		if ta == libc.UintptrFromInt32(0) {
			return -int32(1)
		}
		_r_sort_b(tls, a, ta, n, 0, tr, endch)
		libc.Xfree(tls, ta)
	}
	return 0
}

// C documentation
//
//	/* Unstable, in-place sort. */
func _r_sort_a(tls *libc.TLS, a uintptr, n int32, i int32, tr uintptr, endch uint32) {
	bp := tls.Alloc(10240)
	defer tls.Free(10240)
	var aj, ak, an, cp, r, sp, sp0, sp1, t, v1, v10, v12, v15, v16, v17, v2, v4, v6, v7, v9 uintptr
	var bigc, c, v18, v5, v8 uint32
	var temp Tstack
	var _ /* s at bp+0 */ [512]Tstack
	var _ /* top at bp+8192 */ [256]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aj, ak, an, bigc, c, cp, r, sp, sp0, sp1, t, temp, v1, v10, v12, v15, v16, v17, v18, v2, v4, v5, v6, v7, v8, v9
	/* Set up stack. */
	sp = bp
	(*Tstack)(unsafe.Pointer(sp)).Fsa = a
	(*Tstack)(unsafe.Pointer(sp)).Fsn = n
	v1 = sp
	sp += 16
	(*Tstack)(unsafe.Pointer(v1)).Fsi = i
	for !(bp >= sp) {
		sp -= 16
		v2 = sp
		a = (*Tstack)(unsafe.Pointer(v2)).Fsa
		n = (*Tstack)(unsafe.Pointer(sp)).Fsn
		i = (*Tstack)(unsafe.Pointer(sp)).Fsi
		if n < int32(m_THRESHOLD1) {
			_simplesort(tls, a, n, i, tr, endch)
			continue
		}
		an = a + uintptr(n)*8
		/* Make character histogram. */
		if _nc == uint32(0) {
			_bmin = uint32(255) /* First occupied bin, excluding eos. */
			ak = a
			for {
				if !(ak < an) {
					break
				}
				v4 = ak
				ak += 8
				c = uint32(*(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(v4)) + uintptr(i)))))))
				v6 = uintptr(unsafe.Pointer(&_count)) + uintptr(c)*4
				*(*uint32)(unsafe.Pointer(v6))++
				v5 = *(*uint32)(unsafe.Pointer(v6))
				if v5 == uint32(1) && c != endch {
					if c < _bmin {
						_bmin = c
					}
					_nc++
				}
				goto _3
			_3:
			}
			if sp+uintptr(_nc)*16 > bp+uintptr(m_SIZE)*16 { /* Get more stack. */
				_r_sort_a(tls, a, n, i, tr, endch)
				continue
			}
		}
		/*
		 * Special case: if all strings have the same
		 * character at position i, move on to the next
		 * character.
		 */
		if _nc == uint32(1) && _count[_bmin] == libc.Uint32FromInt32(n) {
			(*Tstack)(unsafe.Pointer(sp)).Fsa = a
			(*Tstack)(unsafe.Pointer(sp)).Fsn = n
			v7 = sp
			sp += 16
			(*Tstack)(unsafe.Pointer(v7)).Fsi = i + libc.Int32FromInt32(1)
			v8 = libc.Uint32FromInt32(0)
			_count[_bmin] = v8
			_nc = v8
			continue
		}
		/*
		 * Set top[]; push incompletely sorted bins onto stack.
		 * top[] = pointers to last out-of-place element in bins.
		 * count[] = counts of elements in bins.
		 * Before permuting: top[c-1] + count[c] = top[c];
		 * during deal: top[c] counts down to top[c-1].
		 */
		v9 = sp
		sp1 = v9
		sp0 = v9                /* Stack position of biggest bin. */
		bigc = uint32(2)        /* Size of biggest bin. */
		if endch == uint32(0) { /* Special case: set top[eos]. */
			v10 = a + uintptr(_count[0])*8
			ak = v10
			(*(*[256]uintptr)(unsafe.Pointer(bp + 8192)))[0] = v10
		} else {
			ak = a
			(*(*[256]uintptr)(unsafe.Pointer(bp + 8192)))[int32(255)] = an
		}
		cp = uintptr(unsafe.Pointer(&_count)) + uintptr(_bmin)*4
		for {
			if !(_nc > uint32(0)) {
				break
			}
			for *(*uint32)(unsafe.Pointer(cp)) == uint32(0) { /* Find next non-empty pile. */
				cp += 4
			}
			if *(*uint32)(unsafe.Pointer(cp)) > uint32(1) {
				if *(*uint32)(unsafe.Pointer(cp)) > bigc {
					bigc = *(*uint32)(unsafe.Pointer(cp))
					sp1 = sp
				}
				(*Tstack)(unsafe.Pointer(sp)).Fsa = ak
				(*Tstack)(unsafe.Pointer(sp)).Fsn = libc.Int32FromUint32(*(*uint32)(unsafe.Pointer(cp)))
				v12 = sp
				sp += 16
				(*Tstack)(unsafe.Pointer(v12)).Fsi = i + libc.Int32FromInt32(1)
			}
			ak += uintptr(*(*uint32)(unsafe.Pointer(cp))) * 8
			(*(*[256]uintptr)(unsafe.Pointer(bp + 8192)))[(int64(cp)-t__predefined_ptrdiff_t(uintptr(unsafe.Pointer(&_count))))/4] = ak
			_nc--
			goto _11
		_11:
			;
			cp += 4
		}
		temp = *(*Tstack)(unsafe.Pointer(sp0))
		*(*Tstack)(unsafe.Pointer(sp0)) = *(*Tstack)(unsafe.Pointer(sp1))
		*(*Tstack)(unsafe.Pointer(sp1)) = temp /* Play it safe -- biggest bin last. */
		/*
		 * Permute misplacements home.  Already home: everything
		 * before aj, and in bin[c], items from top[c] on.
		 * Inner loop:
		 *	r = next element to put in place;
		 *	ak = top[r[i]] = location to put the next element.
		 *	aj = bottom of 1st disordered bin.
		 * Outer loop:
		 *	Once the 1st disordered bin is done, ie. aj >= ak,
		 *	aj<-aj + count[c] connects the bins in a linked list;
		 *	reset count[c].
		 */
		aj = a
		for {
			if !(aj < an) {
				break
			}
			r = *(*uintptr)(unsafe.Pointer(aj))
			for {
				v18 = uint32(*(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(r + uintptr(i)))))))
				c = v18
				v17 = bp + 8192 + uintptr(v18)*8
				*(*uintptr)(unsafe.Pointer(v17)) -= 8
				v16 = *(*uintptr)(unsafe.Pointer(v17))
				v15 = v16
				ak = v15
				if !(aj < v15) {
					break
				}
				t = *(*uintptr)(unsafe.Pointer(ak))
				*(*uintptr)(unsafe.Pointer(ak)) = r
				r = t
				goto _14
			_14:
			}
			goto _13
		_13:
			;
			*(*uintptr)(unsafe.Pointer(aj)) = r
			aj += uintptr(_count[c]) * 8
			_count[c] = libc.Uint32FromInt32(0)
		}
	}
}

var _count [256]uint32

var _nc uint32

var _bmin uint32

// C documentation
//
//	/* Stable sort, requiring additional memory. */
func _r_sort_b(tls *libc.TLS, a uintptr, ta uintptr, n int32, i int32, tr uintptr, endch uint32) {
	bp := tls.Alloc(10240)
	defer tls.Free(10240)
	var ai, ak, cp, sp, sp0, sp1, v1, v11, v13, v14, v16, v17, v18, v2, v4, v6, v7, v8 uintptr
	var bigc, c, v10, v5 uint32
	var temp Tstack
	var _ /* s at bp+0 */ [512]Tstack
	var _ /* top at bp+8192 */ [256]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = ai, ak, bigc, c, cp, sp, sp0, sp1, temp, v1, v10, v11, v13, v14, v16, v17, v18, v2, v4, v5, v6, v7, v8
	sp = bp
	(*Tstack)(unsafe.Pointer(sp)).Fsa = a
	(*Tstack)(unsafe.Pointer(sp)).Fsn = n
	v1 = sp
	sp += 16
	(*Tstack)(unsafe.Pointer(v1)).Fsi = i
	for !(bp >= sp) {
		sp -= 16
		v2 = sp
		a = (*Tstack)(unsafe.Pointer(v2)).Fsa
		n = (*Tstack)(unsafe.Pointer(sp)).Fsn
		i = (*Tstack)(unsafe.Pointer(sp)).Fsi
		if n < int32(m_THRESHOLD1) {
			_simplesort(tls, a, n, i, tr, endch)
			continue
		}
		if _nc1 == uint32(0) {
			_bmin1 = uint32(255)
			ak = a + uintptr(n)*8
			for {
				ak -= 8
				v4 = ak
				if !(v4 >= a) {
					break
				}
				c = uint32(*(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ak)) + uintptr(i)))))))
				v6 = uintptr(unsafe.Pointer(&_count1)) + uintptr(c)*4
				*(*uint32)(unsafe.Pointer(v6))++
				v5 = *(*uint32)(unsafe.Pointer(v6))
				if v5 == uint32(1) && c != endch {
					if c < _bmin1 {
						_bmin1 = c
					}
					_nc1++
				}
				goto _3
			_3:
			}
			if sp+uintptr(_nc1)*16 > bp+uintptr(m_SIZE)*16 {
				_r_sort_b(tls, a, ta, n, i, tr, endch)
				continue
			}
		}
		v7 = sp
		sp1 = v7
		sp0 = v7
		bigc = uint32(2)
		if endch == uint32(0) {
			v8 = a + uintptr(_count1[0])*8
			ak = v8
			(*(*[256]uintptr)(unsafe.Pointer(bp + 8192)))[0] = v8
			_count1[0] = uint32(0)
		} else {
			ak = a
			(*(*[256]uintptr)(unsafe.Pointer(bp + 8192)))[int32(255)] = a + uintptr(n)*8
			_count1[int32(255)] = uint32(0)
		}
		cp = uintptr(unsafe.Pointer(&_count1)) + uintptr(_bmin1)*4
		for {
			if !(_nc1 > uint32(0)) {
				break
			}
			for *(*uint32)(unsafe.Pointer(cp)) == uint32(0) {
				cp += 4
			}
			v10 = *(*uint32)(unsafe.Pointer(cp))
			c = v10
			if v10 > uint32(1) {
				if c > bigc {
					bigc = c
					sp1 = sp
				}
				(*Tstack)(unsafe.Pointer(sp)).Fsa = ak
				(*Tstack)(unsafe.Pointer(sp)).Fsn = libc.Int32FromUint32(c)
				v11 = sp
				sp += 16
				(*Tstack)(unsafe.Pointer(v11)).Fsi = i + libc.Int32FromInt32(1)
			}
			ak += uintptr(c) * 8
			(*(*[256]uintptr)(unsafe.Pointer(bp + 8192)))[(int64(cp)-t__predefined_ptrdiff_t(uintptr(unsafe.Pointer(&_count1))))/4] = ak
			*(*uint32)(unsafe.Pointer(cp)) = uint32(0) /* Reset count[]. */
			_nc1--
			goto _9
		_9:
			;
			cp += 4
		}
		temp = *(*Tstack)(unsafe.Pointer(sp0))
		*(*Tstack)(unsafe.Pointer(sp0)) = *(*Tstack)(unsafe.Pointer(sp1))
		*(*Tstack)(unsafe.Pointer(sp1)) = temp
		ak = ta + uintptr(n)*8
		ai = a + uintptr(n)*8
		for {
			if !(ak > ta) {
				break
			} /* Copy to temp. */
			ak -= 8
			v13 = ak
			ai -= 8
			v14 = ai
			*(*uintptr)(unsafe.Pointer(v13)) = *(*uintptr)(unsafe.Pointer(v14))
			goto _12
		_12:
		}
		ak = ta + uintptr(n)*8
		for {
			ak -= 8
			v16 = ak
			if !(v16 >= ta) {
				break
			} /* Deal to piles. */
			v18 = bp + 8192 + uintptr(*(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ak)) + uintptr(i)))))))*8
			*(*uintptr)(unsafe.Pointer(v18)) -= 8
			v17 = *(*uintptr)(unsafe.Pointer(v18))
			*(*uintptr)(unsafe.Pointer(v17)) = *(*uintptr)(unsafe.Pointer(ak))
			goto _15
		_15:
		}
	}
}

var _count1 [256]uint32

var _nc1 uint32

var _bmin1 uint32

// C documentation
//
//	/* insertion sort */
func _simplesort(tls *libc.TLS, a uintptr, n int32, b int32, tr uintptr, endch uint32) {
	var ai, ak, s, t uintptr
	var ch, v5 uint8
	var v2 int32
	_, _, _, _, _, _, _ = ai, ak, ch, s, t, v2, v5
	ak = a + uintptr(1)*8
	for {
		n--
		v2 = n
		if !(v2 >= int32(1)) {
			break
		}
		ai = ak
		for {
			if !(ai > a) {
				break
			}
			s = *(*uintptr)(unsafe.Pointer(ai)) + uintptr(b)
			t = *(*uintptr)(unsafe.Pointer(ai + uintptr(-libc.Int32FromInt32(1))*8)) + uintptr(b)
			for {
				v5 = *(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(s)))))
				ch = v5
				if !(uint32(v5) != endch) {
					break
				}
				if libc.Int32FromUint8(ch) != libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(t)))))) {
					break
				}
				goto _4
			_4:
				;
				s++
				t++
			}
			if libc.Int32FromUint8(ch) >= libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(tr + uintptr(*(*uint8)(unsafe.Pointer(t)))))) {
				break
			}
			s = *(*uintptr)(unsafe.Pointer(ai))
			*(*uintptr)(unsafe.Pointer(ai)) = *(*uintptr)(unsafe.Pointer(ai + uintptr(-libc.Int32FromInt32(1))*8))
			*(*uintptr)(unsafe.Pointer(ai + uintptr(-libc.Int32FromInt32(1))*8)) = s
			goto _3
		_3:
			;
			ai -= 8
		}
		goto _1
	_1:
		;
		ak += 8
	}
}

const m__SL_CHUNKSIZE = 20

type TStringList = struct {
	Fsl_str uintptr
	Fsl_max Tsize_t
	Fsl_cur Tsize_t
}

type T_stringlist = TStringList

// C documentation
//
//	/*
//	 * sl_init(): Initialize a string list
//	 */
func Xsl_init(tls *libc.TLS) (r uintptr) {
	var sl uintptr
	_ = sl
	sl = libc.Xmalloc(tls, uint64(24))
	if sl == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	(*TStringList)(unsafe.Pointer(sl)).Fsl_cur = uint64(0)
	(*TStringList)(unsafe.Pointer(sl)).Fsl_max = uint64(m__SL_CHUNKSIZE)
	(*TStringList)(unsafe.Pointer(sl)).Fsl_str = Xreallocarray(tls, libc.UintptrFromInt32(0), (*TStringList)(unsafe.Pointer(sl)).Fsl_max, uint64(8))
	if (*TStringList)(unsafe.Pointer(sl)).Fsl_str == libc.UintptrFromInt32(0) {
		libc.Xfree(tls, sl)
		sl = libc.UintptrFromInt32(0)
	}
	return sl
}

// C documentation
//
//	/*
//	 * sl_add(): Add an item to the string list
//	 */
func Xsl_add(tls *libc.TLS, sl uintptr, name uintptr) (r int32) {
	var new1, v2 uintptr
	var v1 Tsize_t
	_, _, _ = new1, v1, v2
	if (*TStringList)(unsafe.Pointer(sl)).Fsl_cur == (*TStringList)(unsafe.Pointer(sl)).Fsl_max-uint64(1) {
		new1 = Xreallocarray(tls, (*TStringList)(unsafe.Pointer(sl)).Fsl_str, (*TStringList)(unsafe.Pointer(sl)).Fsl_max+libc.Uint64FromInt32(m__SL_CHUNKSIZE), uint64(8))
		if new1 == libc.UintptrFromInt32(0) {
			return -int32(1)
		}
		*(*Tsize_t)(unsafe.Pointer(sl + 8)) += uint64(m__SL_CHUNKSIZE)
		(*TStringList)(unsafe.Pointer(sl)).Fsl_str = new1
	}
	v2 = sl + 16
	v1 = *(*Tsize_t)(unsafe.Pointer(v2))
	*(*Tsize_t)(unsafe.Pointer(v2))++
	*(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(v1)*8)) = name
	return 0
}

// C documentation
//
//	/*
//	 * sl_free(): Free a stringlist
//	 */
func Xsl_free(tls *libc.TLS, sl uintptr, all int32) {
	var i Tsize_t
	_ = i
	if sl == libc.UintptrFromInt32(0) {
		return
	}
	if (*TStringList)(unsafe.Pointer(sl)).Fsl_str != 0 {
		if all != 0 {
			i = uint64(0)
			for {
				if !(i < (*TStringList)(unsafe.Pointer(sl)).Fsl_cur) {
					break
				}
				libc.Xfree(tls, *(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(i)*8)))
				goto _1
			_1:
				;
				i++
			}
		}
		libc.Xfree(tls, (*TStringList)(unsafe.Pointer(sl)).Fsl_str)
	}
	libc.Xfree(tls, sl)
}

// C documentation
//
//	/*
//	 * sl_find(): Find a name in the string list
//	 */
func Xsl_find(tls *libc.TLS, sl uintptr, name uintptr) (r uintptr) {
	var i Tsize_t
	_ = i
	i = uint64(0)
	for {
		if !(i < (*TStringList)(unsafe.Pointer(sl)).Fsl_cur) {
			break
		}
		if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(i)*8)), name) == 0 {
			return *(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(i)*8))
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.UintptrFromInt32(0)
}

func Xsl_delete(tls *libc.TLS, sl uintptr, name uintptr, all int32) (r int32) {
	var i, j, v3 Tsize_t
	var v4 uintptr
	_, _, _, _ = i, j, v3, v4
	i = uint64(0)
	for {
		if !(i < (*TStringList)(unsafe.Pointer(sl)).Fsl_cur) {
			break
		}
		if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(i)*8)), name) == 0 {
			if all != 0 {
				libc.Xfree(tls, *(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(i)*8)))
			}
			j = i + uint64(1)
			for {
				if !(j < (*TStringList)(unsafe.Pointer(sl)).Fsl_cur) {
					break
				}
				*(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(j-uint64(1))*8)) = *(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(j)*8))
				goto _2
			_2:
				;
				j++
			}
			v4 = sl + 16
			*(*Tsize_t)(unsafe.Pointer(v4))--
			v3 = *(*Tsize_t)(unsafe.Pointer(v4))
			*(*uintptr)(unsafe.Pointer((*TStringList)(unsafe.Pointer(sl)).Fsl_str + uintptr(v3)*8)) = libc.UintptrFromInt32(0)
			return 0
		}
		goto _1
	_1:
		;
		i++
	}
	return -int32(1)
}

// C documentation
//
//	/*
//	 * Appends src to string dst of size dsize (unlike strncat, dsize is the
//	 * full size of dst, not space left).  At most dsize-1 characters
//	 * will be copied.  Always NUL terminates (unless dsize <= strlen(dst)).
//	 * Returns strlen(src) + MIN(dsize, strlen(initial dst)).
//	 * If retval >= dsize, truncation occurred.
//	 */
func Xstrlcat(tls *libc.TLS, dst uintptr, src uintptr, dsize Tsize_t) (r Tsize_t) {
	var dlen, n, v1, v2 Tsize_t
	var odst, osrc, v3 uintptr
	_, _, _, _, _, _, _ = dlen, n, odst, osrc, v1, v2, v3
	odst = dst
	osrc = src
	n = dsize
	/* Find the end of dst and adjust bytes left but don't go past end. */
	for {
		v1 = n
		n--
		if !(v1 != uint64(0) && int32(*(*int8)(unsafe.Pointer(dst))) != int32('\000')) {
			break
		}
		dst++
	}
	dlen = libc.Uint64FromInt64(int64(dst) - int64(odst))
	n = dsize - dlen
	v2 = n
	n--
	if v2 == uint64(0) {
		return dlen + libc.Xstrlen(tls, src)
	}
	for int32(*(*int8)(unsafe.Pointer(src))) != int32('\000') {
		if n != uint64(0) {
			v3 = dst
			dst++
			*(*int8)(unsafe.Pointer(v3)) = *(*int8)(unsafe.Pointer(src))
			n--
		}
		src++
	}
	*(*int8)(unsafe.Pointer(dst)) = int8('\000')
	return dlen + libc.Uint64FromInt64(int64(src)-int64(osrc)) /* count does not include NUL */
}

// C documentation
//
//	/*
//	 * Copy string src to buffer dst of size dsize.  At most dsize-1
//	 * chars will be copied.  Always NUL terminates (unless dsize == 0).
//	 * Returns strlen(src); if retval >= dsize, truncation occurred.
//	 */
func Xstrlcpy(tls *libc.TLS, dst uintptr, src uintptr, dsize Tsize_t) (r Tsize_t) {
	var nleft, v1 Tsize_t
	var osrc, v3, v4, v5 uintptr
	var v2 int8
	_, _, _, _, _, _, _ = nleft, osrc, v1, v2, v3, v4, v5
	osrc = src
	nleft = dsize
	/* Copy as many bytes as will fit. */
	if nleft != uint64(0) {
		for {
			nleft--
			v1 = nleft
			if !(v1 != uint64(0)) {
				break
			}
			v3 = src
			src++
			v2 = *(*int8)(unsafe.Pointer(v3))
			v4 = dst
			dst++
			*(*int8)(unsafe.Pointer(v4)) = v2
			if int32(v2) == int32('\000') {
				break
			}
		}
	}
	/* Not enough room in dst, add NUL and traverse rest of src. */
	if nleft == uint64(0) {
		if dsize != uint64(0) {
			*(*int8)(unsafe.Pointer(dst)) = int8('\000')
		} /* NUL-terminate dst */
		for {
			v5 = src
			src++
			if !(*(*int8)(unsafe.Pointer(v5)) != 0) {
				break
			}
		}
	}
	return libc.Uint64FromInt64(int64(src) - int64(osrc) - libc.Int64FromInt32(1)) /* count does not include NUL */
}

const m_S_IFBLK1 = 24576
const m_S_IFCHR1 = 8192
const m_S_IFDIR2 = 16384
const m_S_IFIFO1 = 4096
const m_S_IFLNK1 = 40960
const m_S_IFMT1 = 61440
const m_S_IFREG1 = 32768
const m_S_IFSOCK1 = 49152
const m_S_IRGRP2 = 32
const m_S_IROTH2 = 4
const m_S_IRUSR2 = 256
const m_S_ISGID2 = 1024
const m_S_ISUID2 = 2048
const m_S_ISVTX2 = 512
const m_S_IWGRP2 = 16
const m_S_IWOTH2 = 2
const m_S_IWUSR2 = 128
const m_S_IXGRP2 = 8
const m_S_IXOTH2 = 1
const m_S_IXUSR2 = 64

func Xstrmode(tls *libc.TLS, mode Tmode_t, p uintptr) {
	var v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v4, v5, v6, v7, v8, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v4, v5, v6, v7, v8, v9
	/* print type */
	switch mode & libc.Uint32FromInt32(m_S_IFMT1) {
	case uint32(m_S_IFDIR2): /* directory */
		v1 = p
		p++
		*(*int8)(unsafe.Pointer(v1)) = int8('d')
	case uint32(m_S_IFCHR1): /* character special */
		v2 = p
		p++
		*(*int8)(unsafe.Pointer(v2)) = int8('c')
	case uint32(m_S_IFBLK1): /* block special */
		v3 = p
		p++
		*(*int8)(unsafe.Pointer(v3)) = int8('b')
	case uint32(m_S_IFREG1): /* regular */
		v4 = p
		p++
		*(*int8)(unsafe.Pointer(v4)) = int8('-')
	case uint32(m_S_IFLNK1): /* symbolic link */
		v5 = p
		p++
		*(*int8)(unsafe.Pointer(v5)) = int8('l')
	case uint32(m_S_IFSOCK1): /* socket */
		v6 = p
		p++
		*(*int8)(unsafe.Pointer(v6)) = int8('s')
	case uint32(m_S_IFIFO1): /* fifo */
		v7 = p
		p++
		*(*int8)(unsafe.Pointer(v7)) = int8('p')
	default: /* unknown */
		v8 = p
		p++
		*(*int8)(unsafe.Pointer(v8)) = int8('?')
		break
	}
	/* usr */
	if mode&uint32(m_S_IRUSR2) != 0 {
		v9 = p
		p++
		*(*int8)(unsafe.Pointer(v9)) = int8('r')
	} else {
		v10 = p
		p++
		*(*int8)(unsafe.Pointer(v10)) = int8('-')
	}
	if mode&uint32(m_S_IWUSR2) != 0 {
		v11 = p
		p++
		*(*int8)(unsafe.Pointer(v11)) = int8('w')
	} else {
		v12 = p
		p++
		*(*int8)(unsafe.Pointer(v12)) = int8('-')
	}
	switch mode & libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXUSR2)|libc.Int32FromInt32(m_S_ISUID2)) {
	case uint32(0):
		v13 = p
		p++
		*(*int8)(unsafe.Pointer(v13)) = int8('-')
	case uint32(m_S_IXUSR2):
		v14 = p
		p++
		*(*int8)(unsafe.Pointer(v14)) = int8('x')
	case uint32(m_S_ISUID2):
		v15 = p
		p++
		*(*int8)(unsafe.Pointer(v15)) = int8('S')
	case libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXUSR2) | libc.Int32FromInt32(m_S_ISUID2)):
		v16 = p
		p++
		*(*int8)(unsafe.Pointer(v16)) = int8('s')
		break
	}
	/* group */
	if mode&uint32(m_S_IRGRP2) != 0 {
		v17 = p
		p++
		*(*int8)(unsafe.Pointer(v17)) = int8('r')
	} else {
		v18 = p
		p++
		*(*int8)(unsafe.Pointer(v18)) = int8('-')
	}
	if mode&uint32(m_S_IWGRP2) != 0 {
		v19 = p
		p++
		*(*int8)(unsafe.Pointer(v19)) = int8('w')
	} else {
		v20 = p
		p++
		*(*int8)(unsafe.Pointer(v20)) = int8('-')
	}
	switch mode & libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXGRP2)|libc.Int32FromInt32(m_S_ISGID2)) {
	case uint32(0):
		v21 = p
		p++
		*(*int8)(unsafe.Pointer(v21)) = int8('-')
	case uint32(m_S_IXGRP2):
		v22 = p
		p++
		*(*int8)(unsafe.Pointer(v22)) = int8('x')
	case uint32(m_S_ISGID2):
		v23 = p
		p++
		*(*int8)(unsafe.Pointer(v23)) = int8('S')
	case libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXGRP2) | libc.Int32FromInt32(m_S_ISGID2)):
		v24 = p
		p++
		*(*int8)(unsafe.Pointer(v24)) = int8('s')
		break
	}
	/* other */
	if mode&uint32(m_S_IROTH2) != 0 {
		v25 = p
		p++
		*(*int8)(unsafe.Pointer(v25)) = int8('r')
	} else {
		v26 = p
		p++
		*(*int8)(unsafe.Pointer(v26)) = int8('-')
	}
	if mode&uint32(m_S_IWOTH2) != 0 {
		v27 = p
		p++
		*(*int8)(unsafe.Pointer(v27)) = int8('w')
	} else {
		v28 = p
		p++
		*(*int8)(unsafe.Pointer(v28)) = int8('-')
	}
	switch mode & libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXOTH2)|libc.Int32FromInt32(m_S_ISVTX2)) {
	case uint32(0):
		v29 = p
		p++
		*(*int8)(unsafe.Pointer(v29)) = int8('-')
	case uint32(m_S_IXOTH2):
		v30 = p
		p++
		*(*int8)(unsafe.Pointer(v30)) = int8('x')
	case uint32(m_S_ISVTX2):
		v31 = p
		p++
		*(*int8)(unsafe.Pointer(v31)) = int8('T')
	case libc.Uint32FromInt32(libc.Int32FromInt32(m_S_IXOTH2) | libc.Int32FromInt32(m_S_ISVTX2)):
		v32 = p
		p++
		*(*int8)(unsafe.Pointer(v32)) = int8('t')
		break
	}
	v33 = p
	p++
	*(*int8)(unsafe.Pointer(v33)) = int8(' ') /* will be a '+' if ACL's implemented */
	*(*int8)(unsafe.Pointer(p)) = int8('\000')
}

// C documentation
//
//	/*
//	 * Find the first occurrence of find in s, where the search is limited to the
//	 * first slen characters of s.
//	 */
func Xstrnstr(tls *libc.TLS, s uintptr, find uintptr, slen Tsize_t) (r uintptr) {
	var c, sc, v1, v4 int8
	var len1, v3 Tsize_t
	var v2, v5 uintptr
	var v6 bool
	_, _, _, _, _, _, _, _, _ = c, len1, sc, v1, v2, v3, v4, v5, v6
	v2 = find
	find++
	v1 = *(*int8)(unsafe.Pointer(v2))
	c = v1
	if int32(v1) != int32('\000') {
		len1 = libc.Xstrlen(tls, find)
		for cond := true; cond; cond = libc.Xstrncmp(tls, s, find, len1) != 0 {
			for cond := true; cond; cond = int32(sc) != int32(c) {
				v3 = slen
				slen--
				if v6 = v3 < uint64(1); !v6 {
					v5 = s
					s++
					v4 = *(*int8)(unsafe.Pointer(v5))
					sc = v4
				}
				if v6 || int32(v4) == int32('\000') {
					return libc.UintptrFromInt32(0)
				}
			}
			if len1 > slen {
				return libc.UintptrFromInt32(0)
			}
		}
		s--
	}
	return s
}

const m_S_IFBLK2 = 0060000
const m_S_IFCHR2 = 0020000
const m_S_IFDIR3 = 0040000
const m_S_IFIFO2 = 0010000
const m_S_IFLNK2 = 0120000
const m_S_IFMT2 = 0170000
const m_S_IFREG2 = 0100000
const m_S_IFSOCK2 = 0140000
const m_S_IRGRP3 = 0040
const m_S_IROTH3 = 0004
const m_S_IRUSR3 = 0400
const m_S_ISGID3 = 02000
const m_S_ISUID3 = 04000
const m_S_ISVTX3 = 01000
const m_S_IWGRP3 = 0020
const m_S_IWOTH3 = 0002
const m_S_IWUSR3 = 0200
const m_S_IXGRP3 = 0010
const m_S_IXOTH3 = 0001
const m_S_IXUSR3 = 0100

func Xstrtonum(tls *libc.TLS, nptr uintptr, minval int64, maxval int64, errstr uintptr) (r int64) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var rv int64
	var v1 uintptr
	var _ /* e at bp+0 */ int32
	var _ /* eptr at bp+16 */ uintptr
	var _ /* resp at bp+8 */ uintptr
	_, _ = rv, v1
	if errstr == libc.UintptrFromInt32(0) {
		errstr = bp + 8
	}
	if minval > maxval {
		goto out
	}
	rv = Xstrtoi(tls, nptr, bp+16, int32(10), minval, maxval, bp)
	switch *(*int32)(unsafe.Pointer(bp)) {
	case 0:
		*(*uintptr)(unsafe.Pointer(errstr)) = libc.UintptrFromInt32(0)
		return rv
	case int32(m_ECANCELED):
		fallthrough
	case int32(m_EOPNOTSUPP):
		goto out
	case int32(m_ERANGE):
		if *(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))) != 0 {
			goto out
		}
		if rv == maxval {
			v1 = __ccgo_ts + 441
		} else {
			v1 = __ccgo_ts + 451
		}
		*(*uintptr)(unsafe.Pointer(errstr)) = v1
		return 0
	default:
		libc.Xabort(tls)
	}
	goto out
out:
	;
	*(*uintptr)(unsafe.Pointer(errstr)) = __ccgo_ts + 461
	return 0
}

func Xstrtoi(tls *libc.TLS, nptr uintptr, endptr uintptr, base int32, lo Tintmax_t, hi Tintmax_t, rstatus uintptr) (r Tintmax_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var im Tintmax_t
	var serrno int32
	var _ /* ep at bp+0 */ uintptr
	var _ /* rep at bp+8 */ int32
	_, _ = im, serrno
	/* endptr may be NULL */
	if endptr == libc.UintptrFromInt32(0) {
		endptr = bp
	}
	if rstatus == libc.UintptrFromInt32(0) {
		rstatus = bp + 8
	}
	serrno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	im = libc.Xstrtoimax(tls, nptr, endptr, base)
	*(*int32)(unsafe.Pointer(rstatus)) = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = serrno
	/* No digits were found */
	if *(*int32)(unsafe.Pointer(rstatus)) == 0 && nptr == *(*uintptr)(unsafe.Pointer(endptr)) {
		*(*int32)(unsafe.Pointer(rstatus)) = int32(m_ECANCELED)
	}
	if im < lo {
		if *(*int32)(unsafe.Pointer(rstatus)) == 0 {
			*(*int32)(unsafe.Pointer(rstatus)) = int32(m_ERANGE)
		}
		return lo
	}
	if im > hi {
		if *(*int32)(unsafe.Pointer(rstatus)) == 0 {
			*(*int32)(unsafe.Pointer(rstatus)) = int32(m_ERANGE)
		}
		return hi
	}
	/* There are further characters after number */
	if *(*int32)(unsafe.Pointer(rstatus)) == 0 && int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(endptr))))) != int32('\000') {
		*(*int32)(unsafe.Pointer(rstatus)) = int32(m_EOPNOTSUPP)
	}
	return im
}

func Xstrtou(tls *libc.TLS, nptr uintptr, endptr uintptr, base int32, lo Tuintmax_t, hi Tuintmax_t, rstatus uintptr) (r Tuintmax_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var im Tuintmax_t
	var serrno int32
	var _ /* ep at bp+0 */ uintptr
	var _ /* rep at bp+8 */ int32
	_, _ = im, serrno
	/* endptr may be NULL */
	if endptr == libc.UintptrFromInt32(0) {
		endptr = bp
	}
	if rstatus == libc.UintptrFromInt32(0) {
		rstatus = bp + 8
	}
	serrno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	im = libc.Xstrtoumax(tls, nptr, endptr, base)
	*(*int32)(unsafe.Pointer(rstatus)) = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = serrno
	/* No digits were found */
	if *(*int32)(unsafe.Pointer(rstatus)) == 0 && nptr == *(*uintptr)(unsafe.Pointer(endptr)) {
		*(*int32)(unsafe.Pointer(rstatus)) = int32(m_ECANCELED)
	}
	if im < lo {
		if *(*int32)(unsafe.Pointer(rstatus)) == 0 {
			*(*int32)(unsafe.Pointer(rstatus)) = int32(m_ERANGE)
		}
		return lo
	}
	if im > hi {
		if *(*int32)(unsafe.Pointer(rstatus)) == 0 {
			*(*int32)(unsafe.Pointer(rstatus)) = int32(m_ERANGE)
		}
		return hi
	}
	/* There are further characters after number */
	if *(*int32)(unsafe.Pointer(rstatus)) == 0 && int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(endptr))))) != int32('\000') {
		*(*int32)(unsafe.Pointer(rstatus)) = int32(m_EOPNOTSUPP)
	}
	return im
}

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/*
 * - enable time64 functions
 * - symver libbsd<func>_time64 -> <func> 0.7
 */

type Tlibbsd_time64_t = int64

func Xlibbsd_time32_to_time_time64(tls *libc.TLS, t32 Tint32_t) (r Tlibbsd_time64_t) {
	return int64(t32)
}

func Xlibbsd_time_to_time32_time64(tls *libc.TLS, t Tlibbsd_time64_t) (r Tint32_t) {
	return int32(t)
}

func Xlibbsd_time64_to_time_time64(tls *libc.TLS, t64 Tint64_t) (r Tlibbsd_time64_t) {
	return t64
}

func Xlibbsd_time_to_time64_time64(tls *libc.TLS, t Tlibbsd_time64_t) (r Tint64_t) {
	return t
}

func Xlibbsd_time_to_long_time64(tls *libc.TLS, t Tlibbsd_time64_t) (r int64) {
	if uint64(8) == uint64(8) {
		return Xlibbsd_time_to_time64_time64(tls, t)
	}
	return t
}

func Xlibbsd_long_to_time_time64(tls *libc.TLS, tlong int64) (r Tlibbsd_time64_t) {
	if uint64(8) == uint64(4) {
		return Xlibbsd_time32_to_time_time64(tls, int32(tlong))
	}
	return tlong
}

func Xlibbsd_time_to_int_time64(tls *libc.TLS, t Ttime_t) (r int32) {
	if uint64(4) == uint64(8) {
		return int32(Xlibbsd_time_to_time64_time64(tls, t))
	}
	return int32(t)
}

func Xlibbsd_int_to_time_time64(tls *libc.TLS, tint int32) (r Tlibbsd_time64_t) {
	if uint64(4) == uint64(4) {
		return Xlibbsd_time32_to_time_time64(tls, tint)
	}
	return int64(tint)
}

const m_S_AMP = 13
const m_S_CTRL = 4
const m_S_EATCRNL = 12
const m_S_GROUND = 0
const m_S_HEX = 7
const m_S_HEX1 = 8
const m_S_HEX2 = 9
const m_S_META = 2
const m_S_META1 = 3
const m_S_MIME1 = 10
const m_S_MIME2 = 11
const m_S_NUMBER = 14
const m_S_OCTAL2 = 5
const m_S_OCTAL3 = 6
const m_S_START = 1
const m_S_STRING = 15
const m_UNVIS_END = "_VIS_END"
const m_UNVIS_NOCHAR = 3
const m_UNVIS_VALID = 1
const m_UNVIS_VALIDPUSH = 2
const m_VIS_CSTYLE = 0x0002
const m_VIS_DQ = 0x8000
const m_VIS_GLOB = 0x1000
const m_VIS_HTTP1808 = 128
const m_VIS_HTTP1866 = 512
const m_VIS_HTTPSTYLE = 0x0080
const m_VIS_MIMESTYLE = 256
const m_VIS_NL = 0x0010
const m_VIS_NOESCAPE = 1024
const m_VIS_NOLOCALE = 0x4000
const m_VIS_NOSLASH = 0x0040
const m_VIS_OCTAL = 0x0001
const m_VIS_SAFE = 0x0020
const m_VIS_SHELL = 0x2000
const m_VIS_SP = 0x0004
const m_VIS_TAB = 0x0008
const m__VIS_END = 2048

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/*
 * decode driven by state machine
 */

// C documentation
//
//	/*
//	 * RFC 1866
//	 */
type Tnv = struct {
	Fname  [7]int8
	Fvalue Tuint8_t
}

/*
 * Copyright © 2015 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/*
 * Copyright © 2004-2006, 2009-2011 Guillem Jover <guillem@hadrons.org>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. The name of the author may not be used to endorse or promote products
 *    derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL
 * THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
 * OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 * OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 * ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

/* Clang expands this to 1 if an identifier is *not* reserved. */

/*
 * Some libc implementations do not have a <sys/cdefs.h>, in particular
 * musl, try to handle this gracefully.
 */
/* Copyright (C) 1992-2023 Free Software Foundation, Inc.
   Copyright The GNU Toolchain Authors.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/*
 * decode driven by state machine
 */

// C documentation
//
//	/*
//	 * RFC 1866
//	 */
var _nv = [100]Tnv{
	0: {
		Fname:  [7]int8{'A', 'E', 'l', 'i', 'g'},
		Fvalue: uint8(198),
	},
	1: {
		Fname:  [7]int8{'A', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(193),
	},
	2: {
		Fname:  [7]int8{'A', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(194),
	},
	3: {
		Fname:  [7]int8{'A', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(192),
	},
	4: {
		Fname:  [7]int8{'A', 'r', 'i', 'n', 'g'},
		Fvalue: uint8(197),
	},
	5: {
		Fname:  [7]int8{'A', 't', 'i', 'l', 'd', 'e'},
		Fvalue: uint8(195),
	},
	6: {
		Fname:  [7]int8{'A', 'u', 'm', 'l'},
		Fvalue: uint8(196),
	},
	7: {
		Fname:  [7]int8{'C', 'c', 'e', 'd', 'i', 'l'},
		Fvalue: uint8(199),
	},
	8: {
		Fname:  [7]int8{'E', 'T', 'H'},
		Fvalue: uint8(208),
	},
	9: {
		Fname:  [7]int8{'E', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(201),
	},
	10: {
		Fname:  [7]int8{'E', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(202),
	},
	11: {
		Fname:  [7]int8{'E', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(200),
	},
	12: {
		Fname:  [7]int8{'E', 'u', 'm', 'l'},
		Fvalue: uint8(203),
	},
	13: {
		Fname:  [7]int8{'I', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(205),
	},
	14: {
		Fname:  [7]int8{'I', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(206),
	},
	15: {
		Fname:  [7]int8{'I', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(204),
	},
	16: {
		Fname:  [7]int8{'I', 'u', 'm', 'l'},
		Fvalue: uint8(207),
	},
	17: {
		Fname:  [7]int8{'N', 't', 'i', 'l', 'd', 'e'},
		Fvalue: uint8(209),
	},
	18: {
		Fname:  [7]int8{'O', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(211),
	},
	19: {
		Fname:  [7]int8{'O', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(212),
	},
	20: {
		Fname:  [7]int8{'O', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(210),
	},
	21: {
		Fname:  [7]int8{'O', 's', 'l', 'a', 's', 'h'},
		Fvalue: uint8(216),
	},
	22: {
		Fname:  [7]int8{'O', 't', 'i', 'l', 'd', 'e'},
		Fvalue: uint8(213),
	},
	23: {
		Fname:  [7]int8{'O', 'u', 'm', 'l'},
		Fvalue: uint8(214),
	},
	24: {
		Fname:  [7]int8{'T', 'H', 'O', 'R', 'N'},
		Fvalue: uint8(222),
	},
	25: {
		Fname:  [7]int8{'U', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(218),
	},
	26: {
		Fname:  [7]int8{'U', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(219),
	},
	27: {
		Fname:  [7]int8{'U', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(217),
	},
	28: {
		Fname:  [7]int8{'U', 'u', 'm', 'l'},
		Fvalue: uint8(220),
	},
	29: {
		Fname:  [7]int8{'Y', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(221),
	},
	30: {
		Fname:  [7]int8{'a', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(225),
	},
	31: {
		Fname:  [7]int8{'a', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(226),
	},
	32: {
		Fname:  [7]int8{'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(180),
	},
	33: {
		Fname:  [7]int8{'a', 'e', 'l', 'i', 'g'},
		Fvalue: uint8(230),
	},
	34: {
		Fname:  [7]int8{'a', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(224),
	},
	35: {
		Fname:  [7]int8{'a', 'm', 'p'},
		Fvalue: uint8(38),
	},
	36: {
		Fname:  [7]int8{'a', 'r', 'i', 'n', 'g'},
		Fvalue: uint8(229),
	},
	37: {
		Fname:  [7]int8{'a', 't', 'i', 'l', 'd', 'e'},
		Fvalue: uint8(227),
	},
	38: {
		Fname:  [7]int8{'a', 'u', 'm', 'l'},
		Fvalue: uint8(228),
	},
	39: {
		Fname:  [7]int8{'b', 'r', 'v', 'b', 'a', 'r'},
		Fvalue: uint8(166),
	},
	40: {
		Fname:  [7]int8{'c', 'c', 'e', 'd', 'i', 'l'},
		Fvalue: uint8(231),
	},
	41: {
		Fname:  [7]int8{'c', 'e', 'd', 'i', 'l'},
		Fvalue: uint8(184),
	},
	42: {
		Fname:  [7]int8{'c', 'e', 'n', 't'},
		Fvalue: uint8(162),
	},
	43: {
		Fname:  [7]int8{'c', 'o', 'p', 'y'},
		Fvalue: uint8(169),
	},
	44: {
		Fname:  [7]int8{'c', 'u', 'r', 'r', 'e', 'n'},
		Fvalue: uint8(164),
	},
	45: {
		Fname:  [7]int8{'d', 'e', 'g'},
		Fvalue: uint8(176),
	},
	46: {
		Fname:  [7]int8{'d', 'i', 'v', 'i', 'd', 'e'},
		Fvalue: uint8(247),
	},
	47: {
		Fname:  [7]int8{'e', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(233),
	},
	48: {
		Fname:  [7]int8{'e', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(234),
	},
	49: {
		Fname:  [7]int8{'e', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(232),
	},
	50: {
		Fname:  [7]int8{'e', 't', 'h'},
		Fvalue: uint8(240),
	},
	51: {
		Fname:  [7]int8{'e', 'u', 'm', 'l'},
		Fvalue: uint8(235),
	},
	52: {
		Fname:  [7]int8{'f', 'r', 'a', 'c', '1', '2'},
		Fvalue: uint8(189),
	},
	53: {
		Fname:  [7]int8{'f', 'r', 'a', 'c', '1', '4'},
		Fvalue: uint8(188),
	},
	54: {
		Fname:  [7]int8{'f', 'r', 'a', 'c', '3', '4'},
		Fvalue: uint8(190),
	},
	55: {
		Fname:  [7]int8{'g', 't'},
		Fvalue: uint8(62),
	},
	56: {
		Fname:  [7]int8{'i', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(237),
	},
	57: {
		Fname:  [7]int8{'i', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(238),
	},
	58: {
		Fname:  [7]int8{'i', 'e', 'x', 'c', 'l'},
		Fvalue: uint8(161),
	},
	59: {
		Fname:  [7]int8{'i', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(236),
	},
	60: {
		Fname:  [7]int8{'i', 'q', 'u', 'e', 's', 't'},
		Fvalue: uint8(191),
	},
	61: {
		Fname:  [7]int8{'i', 'u', 'm', 'l'},
		Fvalue: uint8(239),
	},
	62: {
		Fname:  [7]int8{'l', 'a', 'q', 'u', 'o'},
		Fvalue: uint8(171),
	},
	63: {
		Fname:  [7]int8{'l', 't'},
		Fvalue: uint8(60),
	},
	64: {
		Fname:  [7]int8{'m', 'a', 'c', 'r'},
		Fvalue: uint8(175),
	},
	65: {
		Fname:  [7]int8{'m', 'i', 'c', 'r', 'o'},
		Fvalue: uint8(181),
	},
	66: {
		Fname:  [7]int8{'m', 'i', 'd', 'd', 'o', 't'},
		Fvalue: uint8(183),
	},
	67: {
		Fname:  [7]int8{'n', 'b', 's', 'p'},
		Fvalue: uint8(160),
	},
	68: {
		Fname:  [7]int8{'n', 'o', 't'},
		Fvalue: uint8(172),
	},
	69: {
		Fname:  [7]int8{'n', 't', 'i', 'l', 'd', 'e'},
		Fvalue: uint8(241),
	},
	70: {
		Fname:  [7]int8{'o', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(243),
	},
	71: {
		Fname:  [7]int8{'o', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(244),
	},
	72: {
		Fname:  [7]int8{'o', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(242),
	},
	73: {
		Fname:  [7]int8{'o', 'r', 'd', 'f'},
		Fvalue: uint8(170),
	},
	74: {
		Fname:  [7]int8{'o', 'r', 'd', 'm'},
		Fvalue: uint8(186),
	},
	75: {
		Fname:  [7]int8{'o', 's', 'l', 'a', 's', 'h'},
		Fvalue: uint8(248),
	},
	76: {
		Fname:  [7]int8{'o', 't', 'i', 'l', 'd', 'e'},
		Fvalue: uint8(245),
	},
	77: {
		Fname:  [7]int8{'o', 'u', 'm', 'l'},
		Fvalue: uint8(246),
	},
	78: {
		Fname:  [7]int8{'p', 'a', 'r', 'a'},
		Fvalue: uint8(182),
	},
	79: {
		Fname:  [7]int8{'p', 'l', 'u', 's', 'm', 'n'},
		Fvalue: uint8(177),
	},
	80: {
		Fname:  [7]int8{'p', 'o', 'u', 'n', 'd'},
		Fvalue: uint8(163),
	},
	81: {
		Fname:  [7]int8{'q', 'u', 'o', 't'},
		Fvalue: uint8(34),
	},
	82: {
		Fname:  [7]int8{'r', 'a', 'q', 'u', 'o'},
		Fvalue: uint8(187),
	},
	83: {
		Fname:  [7]int8{'r', 'e', 'g'},
		Fvalue: uint8(174),
	},
	84: {
		Fname:  [7]int8{'s', 'e', 'c', 't'},
		Fvalue: uint8(167),
	},
	85: {
		Fname:  [7]int8{'s', 'h', 'y'},
		Fvalue: uint8(173),
	},
	86: {
		Fname:  [7]int8{'s', 'u', 'p', '1'},
		Fvalue: uint8(185),
	},
	87: {
		Fname:  [7]int8{'s', 'u', 'p', '2'},
		Fvalue: uint8(178),
	},
	88: {
		Fname:  [7]int8{'s', 'u', 'p', '3'},
		Fvalue: uint8(179),
	},
	89: {
		Fname:  [7]int8{'s', 'z', 'l', 'i', 'g'},
		Fvalue: uint8(223),
	},
	90: {
		Fname:  [7]int8{'t', 'h', 'o', 'r', 'n'},
		Fvalue: uint8(254),
	},
	91: {
		Fname:  [7]int8{'t', 'i', 'm', 'e', 's'},
		Fvalue: uint8(215),
	},
	92: {
		Fname:  [7]int8{'u', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(250),
	},
	93: {
		Fname:  [7]int8{'u', 'c', 'i', 'r', 'c'},
		Fvalue: uint8(251),
	},
	94: {
		Fname:  [7]int8{'u', 'g', 'r', 'a', 'v', 'e'},
		Fvalue: uint8(249),
	},
	95: {
		Fname:  [7]int8{'u', 'm', 'l'},
		Fvalue: uint8(168),
	},
	96: {
		Fname:  [7]int8{'u', 'u', 'm', 'l'},
		Fvalue: uint8(252),
	},
	97: {
		Fname:  [7]int8{'y', 'a', 'c', 'u', 't', 'e'},
		Fvalue: uint8(253),
	},
	98: {
		Fname:  [7]int8{'y', 'e', 'n'},
		Fvalue: uint8(165),
	},
	99: {
		Fname:  [7]int8{'y', 'u', 'm', 'l'},
		Fvalue: uint8(255),
	},
}

// C documentation
//
//	/*
//	 * unvis - decode characters previously encoded by vis
//	 */
func Xunvis(tls *libc.TLS, cp uintptr, c int32, astate uintptr, flag int32) (r int32) {
	var ia, is, lc, st, uc uint8
	var v22, v23, v24, v25, v26 int32
	var p19, p20, p21, p28 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = ia, is, lc, st, uc, v22, v23, v24, v25, v26, p19, p20, p21, p28
	uc = libc.Uint8FromInt32(c)
	/*
	 * Bottom 8 bits of astate hold the state machine state.
	 * Top 8 bits hold the current character in the http 1866 nv string decoding
	 */
	st = libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(astate)) & libc.Int32FromInt32(0xff))
	if flag&int32(m__VIS_END) != 0 {
		switch libc.Int32FromUint8(st) {
		case int32(m_S_OCTAL2):
			fallthrough
		case int32(m_S_OCTAL3):
			fallthrough
		case int32(m_S_HEX2):
			*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
			return int32(m_UNVIS_VALID)
		case m_S_GROUND:
			return int32(m_UNVIS_NOCHAR)
		default:
			return -int32(1)
		}
	}
	switch libc.Int32FromUint8(st) {
	case m_S_GROUND:
		goto _1
	case int32(m_S_START):
		goto _2
	case int32(m_S_META):
		goto _3
	case int32(m_S_META1):
		goto _4
	case int32(m_S_CTRL):
		goto _5
	case int32(m_S_OCTAL2):
		goto _6
	case int32(m_S_OCTAL3):
		goto _7
	case int32(m_S_HEX):
		goto _8
	case int32(m_S_HEX1):
		goto _9
	case int32(m_S_HEX2):
		goto _10
	case int32(m_S_MIME1):
		goto _11
	case int32(m_S_MIME2):
		goto _12
	case int32(m_S_EATCRNL):
		goto _13
	case int32(m_S_AMP):
		goto _14
	case int32(m_S_STRING):
		goto _15
	case int32(m_S_NUMBER):
		goto _16
	default:
		goto _17
	}
	goto _18
_1:
	;
	*(*int8)(unsafe.Pointer(cp)) = 0
	if flag&int32(m_VIS_NOESCAPE) == 0 && c == int32('\\') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_START)))
		return int32(m_UNVIS_NOCHAR)
	}
	if flag&int32(m_VIS_HTTP1808) != 0 && c == int32('%') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_HEX1)))
		return int32(m_UNVIS_NOCHAR)
	}
	if flag&int32(m_VIS_HTTP1866) != 0 && c == int32('&') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_AMP)))
		return int32(m_UNVIS_NOCHAR)
	}
	if flag&int32(m_VIS_MIMESTYLE) != 0 && c == int32('=') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_MIME1)))
		return int32(m_UNVIS_NOCHAR)
	}
	*(*int8)(unsafe.Pointer(cp)) = int8(c)
	return int32(m_UNVIS_VALID)
_2:
	;
	switch c {
	case int32('\\'):
		*(*int8)(unsafe.Pointer(cp)) = int8(c)
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('0'):
		fallthrough
	case int32('1'):
		fallthrough
	case int32('2'):
		fallthrough
	case int32('3'):
		fallthrough
	case int32('4'):
		fallthrough
	case int32('5'):
		fallthrough
	case int32('6'):
		fallthrough
	case int32('7'):
		*(*int8)(unsafe.Pointer(cp)) = int8(c - libc.Int32FromUint8('0'))
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_OCTAL2)))
		return int32(m_UNVIS_NOCHAR)
	case int32('M'):
		*(*int8)(unsafe.Pointer(cp)) = libc.Int8FromInt32(0200)
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_META)))
		return int32(m_UNVIS_NOCHAR)
	case int32('^'):
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_CTRL)))
		return int32(m_UNVIS_NOCHAR)
	case int32('n'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\n')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('r'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\r')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('b'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\b')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('a'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\007')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('v'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\v')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('t'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\t')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('f'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\f')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('s'):
		*(*int8)(unsafe.Pointer(cp)) = int8(' ')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('E'):
		*(*int8)(unsafe.Pointer(cp)) = int8('\033')
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	case int32('x'):
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_HEX)))
		return int32(m_UNVIS_NOCHAR)
	case int32('\n'):
		/*
		 * hidden newline
		 */
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_NOCHAR)
	case int32('$'):
		/*
		 * hidden marker
		 */
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_NOCHAR)
	default:
		if libc.BoolInt32(libc.Uint32FromInt32(c)-uint32(0x21) < uint32(0x5e)) != 0 {
			*(*int8)(unsafe.Pointer(cp)) = int8(c)
			*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
			return int32(m_UNVIS_VALID)
		}
	}
	goto bad
_3:
	;
	if c == int32('-') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_META1)))
	} else {
		if c == int32('^') {
			*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_CTRL)))
		} else {
			goto bad
		}
	}
	return int32(m_UNVIS_NOCHAR)
_4:
	;
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	p19 = cp
	*(*int8)(unsafe.Pointer(p19)) = int8(int32(*(*int8)(unsafe.Pointer(p19))) | c)
	return int32(m_UNVIS_VALID)
_5:
	;
	if c == int32('?') {
		p20 = cp
		*(*int8)(unsafe.Pointer(p20)) = int8(int32(*(*int8)(unsafe.Pointer(p20))) | libc.Int32FromInt32(0177))
	} else {
		p21 = cp
		*(*int8)(unsafe.Pointer(p21)) = int8(int32(*(*int8)(unsafe.Pointer(p21))) | c&libc.Int32FromInt32(037))
	}
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	return int32(m_UNVIS_VALID)
_6:
	; /* second possible octal digit */
	if libc.Int32FromUint8(uc) >= int32('0') && libc.Int32FromUint8(uc) <= int32('7') {
		/*
		 * yes - and maybe a third
		 */
		*(*int8)(unsafe.Pointer(cp)) = int8(int32(*(*int8)(unsafe.Pointer(cp)))<<int32(3) + (c - int32('0')))
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_OCTAL3)))
		return int32(m_UNVIS_NOCHAR)
	}
	/*
	 * no - done with current sequence, push back passed char
	 */
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	return int32(m_UNVIS_VALIDPUSH)
_7:
	; /* third possible octal digit */
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	if libc.Int32FromUint8(uc) >= int32('0') && libc.Int32FromUint8(uc) <= int32('7') {
		*(*int8)(unsafe.Pointer(cp)) = int8(int32(*(*int8)(unsafe.Pointer(cp)))<<int32(3) + (c - int32('0')))
		return int32(m_UNVIS_VALID)
	}
	/*
	 * we were done, push back passed char
	 */
	return int32(m_UNVIS_VALIDPUSH)
_8:
	;
	if !(libc.Xisxdigit(tls, libc.Int32FromUint8(uc)) != 0) {
		goto bad
	}
	/*FALLTHROUGH*/
_9:
	;
	if libc.Xisxdigit(tls, libc.Int32FromUint8(uc)) != 0 {
		if libc.BoolInt32(uint32(uc)-uint32('0') < uint32(10)) != 0 {
			v22 = libc.Int32FromUint8(uc) - int32('0')
		} else {
			v22 = libc.Xtolower(tls, libc.Int32FromUint8(uc)) - int32('a') + int32(10)
		}
		*(*int8)(unsafe.Pointer(cp)) = int8(v22)
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_HEX2)))
		return int32(m_UNVIS_NOCHAR)
	}
	/*
	 * no - done with current sequence, push back passed char
	 */
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	return int32(m_UNVIS_VALIDPUSH)
_10:
	;
	*(*int32)(unsafe.Pointer(astate)) = m_S_GROUND
	if libc.Xisxdigit(tls, libc.Int32FromUint8(uc)) != 0 {
		if libc.BoolInt32(uint32(uc)-uint32('0') < uint32(10)) != 0 {
			v23 = libc.Int32FromUint8(uc) - int32('0')
		} else {
			v23 = libc.Xtolower(tls, libc.Int32FromUint8(uc)) - int32('a') + int32(10)
		}
		*(*int8)(unsafe.Pointer(cp)) = int8(v23 | int32(*(*int8)(unsafe.Pointer(cp)))<<int32(4))
		return int32(m_UNVIS_VALID)
	}
	return int32(m_UNVIS_VALIDPUSH)
_11:
	;
	if libc.Int32FromUint8(uc) == int32('\n') || libc.Int32FromUint8(uc) == int32('\r') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_EATCRNL)))
		return int32(m_UNVIS_NOCHAR)
	}
	if libc.Xisxdigit(tls, libc.Int32FromUint8(uc)) != 0 && (libc.BoolInt32(uint32(uc)-uint32('0') < uint32(10)) != 0 || libc.BoolInt32(uint32(uc)-uint32('A') < uint32(26)) != 0) {
		if libc.BoolInt32(uint32(uc)-uint32('0') < uint32(10)) != 0 {
			v24 = libc.Int32FromUint8(uc) - int32('0')
		} else {
			v24 = libc.Int32FromUint8(uc) - int32('A') + int32(10)
		}
		*(*int8)(unsafe.Pointer(cp)) = int8(v24)
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_MIME2)))
		return int32(m_UNVIS_NOCHAR)
	}
	goto bad
_12:
	;
	if libc.Xisxdigit(tls, libc.Int32FromUint8(uc)) != 0 && (libc.BoolInt32(uint32(uc)-uint32('0') < uint32(10)) != 0 || libc.BoolInt32(uint32(uc)-uint32('A') < uint32(26)) != 0) {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		if libc.BoolInt32(uint32(uc)-uint32('0') < uint32(10)) != 0 {
			v25 = libc.Int32FromUint8(uc) - int32('0')
		} else {
			v25 = libc.Int32FromUint8(uc) - int32('A') + int32(10)
		}
		*(*int8)(unsafe.Pointer(cp)) = int8(v25 | int32(*(*int8)(unsafe.Pointer(cp)))<<int32(4))
		return int32(m_UNVIS_VALID)
	}
	goto bad
_13:
	;
	switch libc.Int32FromUint8(uc) {
	case int32('\r'):
		fallthrough
	case int32('\n'):
		return int32(m_UNVIS_NOCHAR)
	case int32('='):
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_MIME1)))
		return int32(m_UNVIS_NOCHAR)
	default:
		*(*int8)(unsafe.Pointer(cp)) = libc.Int8FromUint8(uc)
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
		return int32(m_UNVIS_VALID)
	}
_14:
	;
	*(*int8)(unsafe.Pointer(cp)) = 0
	if libc.Int32FromUint8(uc) == int32('#') {
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_NUMBER)))
		return int32(m_UNVIS_NOCHAR)
	}
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_STRING)))
	/*FALLTHROUGH*/
_15:
	;
	ia = libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(cp)))                                          /* index in the array */
	is = uint8(libc.Uint32FromInt32(*(*int32)(unsafe.Pointer(astate))) >> libc.Int32FromInt32(24)) /* index in the string */
	if libc.Int32FromUint8(is) == 0 {
		v26 = 0
	} else {
		v26 = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&_nv)) + uintptr(ia)*8 + uintptr(libc.Int32FromUint8(is)-int32(1)))))
	}
	lc = libc.Uint8FromInt32(v26) /* last character */
	if libc.Int32FromUint8(uc) == int32(';') {
		uc = uint8('\000')
	}
	for {
		if !(uint64(ia) < libc.Uint64FromInt64(800)/libc.Uint64FromInt64(8)) {
			break
		}
		if libc.Int32FromUint8(is) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&_nv)) + uintptr(ia)*8 + uintptr(libc.Int32FromUint8(is)-int32(1))))) != libc.Int32FromUint8(lc) {
			goto bad
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&_nv)) + uintptr(ia)*8 + uintptr(is)))) == libc.Int32FromUint8(uc) {
			break
		}
		goto _27
	_27:
		;
		ia++
	}
	if uint64(ia) == libc.Uint64FromInt64(800)/libc.Uint64FromInt64(8) {
		goto bad
	}
	if libc.Int32FromUint8(uc) != 0 {
		*(*int8)(unsafe.Pointer(cp)) = libc.Int8FromUint8(ia)
		*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromUint8(is)+libc.Int32FromInt32(1))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_STRING)))
		return int32(m_UNVIS_NOCHAR)
	}
	*(*int8)(unsafe.Pointer(cp)) = libc.Int8FromUint8(_nv[ia].Fvalue)
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	return int32(m_UNVIS_VALID)
_16:
	;
	if libc.Int32FromUint8(uc) == int32(';') {
		return int32(m_UNVIS_VALID)
	}
	if !(libc.BoolInt32(uint32(uc)-libc.Uint32FromUint8('0') < libc.Uint32FromInt32(10)) != 0) {
		goto bad
	}
	p28 = cp
	*(*int8)(unsafe.Pointer(p28)) = int8(int32(*(*int8)(unsafe.Pointer(p28))) + (int32(*(*int8)(unsafe.Pointer(cp)))*libc.Int32FromInt32(10) + libc.Int32FromUint8(uc) - libc.Int32FromUint8('0')))
	return int32(m_UNVIS_NOCHAR)
_17:
	;
	goto bad
bad:
	;
	/*
	 * decoder in unknown state - (probably uninitialized)
	 */
	*(*int32)(unsafe.Pointer(astate)) = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(0))<<libc.Int32FromInt32(24) | libc.Uint32FromInt32(libc.Int32FromInt32(m_S_GROUND)))
	return -int32(1)
_18:
	;
	return r
}

/*
 * strnunvisx - decode src into dst
 *
 *	Number of chars decoded into dst is returned, -1 on error.
 *	Dst is null terminated.
 */

func Xstrnunvisx(tls *libc.TLS, dst uintptr, dlen Tsize_t, src uintptr, flag int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var c, v1 int8
	var start, v14, v16, v18, v2 uintptr
	var v13, v15, v17, v19 Tsize_t
	var _ /* state at bp+4 */ int32
	var _ /* t at bp+0 */ int8
	_, _, _, _, _, _, _, _, _, _, _ = c, start, v1, v13, v14, v15, v16, v17, v18, v19, v2
	*(*int8)(unsafe.Pointer(bp)) = int8('\000')
	start = dst
	*(*int32)(unsafe.Pointer(bp + 4)) = 0
	for {
		v2 = src
		src++
		v1 = *(*int8)(unsafe.Pointer(v2))
		c = v1
		if !(int32(v1) != int32('\000')) {
			break
		}
		goto again
	again:
		;
		switch Xunvis(tls, bp, int32(c), bp+4, flag) {
		case int32(m_UNVIS_VALID):
			goto _3
		case int32(m_UNVIS_VALIDPUSH):
			goto _4
		case int32(m_UNVIS_NOCHAR):
			goto _5
		case 0:
			goto _6
		case -int32(1):
			goto _7
		default:
			goto _8
		}
		goto _9
	_3:
		;
	_12:
		;
		v13 = dlen
		dlen--
		if v13 == uint64(0) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOSPC)
			return -int32(1)
		}
		goto _11
	_11:
		;
		if 0 != 0 {
			goto _12
		}
		goto _10
	_10:
		;
		v14 = dst
		dst++
		*(*int8)(unsafe.Pointer(v14)) = *(*int8)(unsafe.Pointer(bp))
		goto _9
	_4:
		;
		v15 = dlen
		dlen--
		if v15 == uint64(0) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOSPC)
			return -int32(1)
		}
		v16 = dst
		dst++
		*(*int8)(unsafe.Pointer(v16)) = *(*int8)(unsafe.Pointer(bp))
		goto again
	_6:
		;
	_5:
		;
		goto _9
	_7:
		;
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	_8:
		;
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_EINVAL)
		return -int32(1)
	_9:
	}
	if Xunvis(tls, bp, int32(c), bp+4, int32(m__VIS_END)) == int32(m_UNVIS_VALID) {
		v17 = dlen
		dlen--
		if v17 == uint64(0) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOSPC)
			return -int32(1)
		}
		v18 = dst
		dst++
		*(*int8)(unsafe.Pointer(v18)) = *(*int8)(unsafe.Pointer(bp))
	}
	v19 = dlen
	dlen--
	if v19 == uint64(0) {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOSPC)
		return -int32(1)
	}
	*(*int8)(unsafe.Pointer(dst)) = int8('\000')
	return int32(int64(dst) - int64(start))
}

func Xstrunvisx(tls *libc.TLS, dst uintptr, src uintptr, flag int32) (r int32) {
	return Xstrnunvisx(tls, dst, libc.Uint64FromInt32(^libc.Int32FromInt32(0)), src, flag)
}

func Xstrunvis(tls *libc.TLS, dst uintptr, src uintptr) (r int32) {
	return Xstrnunvisx(tls, dst, libc.Uint64FromInt32(^libc.Int32FromInt32(0)), src, 0)
}

func Xstrnunvis_openbsd(tls *libc.TLS, dst uintptr, src uintptr, dlen Tsize_t) (r Tssize_t) {
	return int64(Xstrnunvisx(tls, dst, dlen, src, 0))
}

func Xstrnunvis_netbsd(tls *libc.TLS, dst uintptr, dlen Tsize_t, src uintptr) (r int32) {
	return Xstrnunvisx(tls, dst, dlen, src, 0)
}

const m_MAXEXTRAS = 30
const m_VIS_CSTYLE1 = 2
const m_VIS_DQ1 = 32768
const m_VIS_GLOB1 = 4096
const m_VIS_HTTP18081 = 0x0080
const m_VIS_HTTP18661 = 0x0200
const m_VIS_HTTPSTYLE1 = 128
const m_VIS_NL1 = 16
const m_VIS_NOESCAPE1 = 0x0400
const m_VIS_NOLOCALE1 = 16384
const m_VIS_NOSLASH1 = 64
const m_VIS_OCTAL1 = 1
const m_VIS_SAFE1 = 32
const m_VIS_SHELL1 = 8192
const m_VIS_SP1 = 4
const m_VIS_TAB1 = 8
const m__VIS_END1 = 0x0800

type Twctrans_t = uintptr

/* Keep it simple for now, no locale stuff */

var _char_shell = [19]Twchar_t{'\'', '`', '"', ';', '&', '<', '>', '(', ')', '|', '{', '}', ']', '\\', '$', '!', '^', '~'}
var _char_glob = [5]Twchar_t{'*', '?', '[', '#'}

/*
 * On NetBSD and glibc MB_LEN_MAX is currently > 8 which does not fit on any
 * integer integral type and it is probably wrong, since currently the maximum
 * number of bytes and character needs is 6. Until this is fixed, the
 * loops below are using sizeof(uint64_t) - 1 instead of MB_LEN_MAX, and
 * the assertion is commented out.
 */

// C documentation
//
//	/*
//	 * This is do_hvis, for HTTP style (RFC 1808)
//	 */
func _do_hvis(tls *libc.TLS, dst uintptr, c Twint_t, flags int32, nextc Twint_t, extra uintptr) (r uintptr) {
	var v1, v2, v3 uintptr
	_, _, _ = v1, v2, v3
	if libc.Xiswalnum(tls, c) != 0 || c == uint32('$') || c == uint32('-') || c == uint32('_') || c == uint32('.') || c == uint32('+') || c == uint32('!') || c == uint32('*') || c == uint32('\'') || c == uint32('(') || c == uint32(')') || c == uint32(',') {
		dst = _do_svis(tls, dst, c, flags, nextc, extra)
	} else {
		v1 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v1)) = int32('%')
		v2 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v2)) = *(*t__predefined_wchar_t)(unsafe.Pointer(__ccgo_ts + 469 + uintptr(c>>int32(4)&uint32(0xf))*4))
		v3 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v3)) = *(*t__predefined_wchar_t)(unsafe.Pointer(__ccgo_ts + 469 + uintptr(c&uint32(0xf))*4))
	}
	return dst
}

// C documentation
//
//	/*
//	 * This is do_mvis, for Quoted-Printable MIME (RFC 2045)
//	 * NB: No handling of long lines or CRLF.
//	 */
func _do_mvis(tls *libc.TLS, dst uintptr, c Twint_t, flags int32, nextc Twint_t, extra uintptr) (r uintptr) {
	var v1, v2, v3 uintptr
	_, _, _ = v1, v2, v3
	if c != uint32('\n') && (libc.Xiswspace(tls, c) != 0 && (nextc == uint32('\r') || nextc == uint32('\n')) || !(libc.Xiswspace(tls, c) != 0) && (c < uint32(33) || c > uint32(60) && c < uint32(62) || c > uint32(126)) || libc.Xwcschr(tls, __ccgo_ts+537, libc.Int32FromUint32(c)) != libc.UintptrFromInt32(0)) {
		v1 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v1)) = int32('=')
		v2 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v2)) = *(*t__predefined_wchar_t)(unsafe.Pointer(__ccgo_ts + 589 + uintptr(c>>int32(4)&uint32(0xf))*4))
		v3 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v3)) = *(*t__predefined_wchar_t)(unsafe.Pointer(__ccgo_ts + 589 + uintptr(c&uint32(0xf))*4))
	} else {
		dst = _do_svis(tls, dst, c, flags, nextc, extra)
	}
	return dst
}

// C documentation
//
//	/*
//	 * Output single byte of multibyte character.
//	 */
func _do_mbyte(tls *libc.TLS, dst uintptr, c Twint_t, flags int32, nextc Twint_t, iswextra int32) (r uintptr) {
	var v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v4, v5, v6, v7, v8, v9 uintptr
	var v21 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v4, v5, v6, v7, v8, v9
	if flags&int32(m_VIS_CSTYLE1) != 0 {
		switch c {
		case uint32('\n'):
			v1 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v1)) = int32('\\')
			v2 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v2)) = int32('n')
			return dst
		case uint32('\r'):
			v3 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v3)) = int32('\\')
			v4 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v4)) = int32('r')
			return dst
		case uint32('\b'):
			v5 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v5)) = int32('\\')
			v6 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v6)) = int32('b')
			return dst
		case uint32('\a'):
			v7 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v7)) = int32('\\')
			v8 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v8)) = int32('a')
			return dst
		case uint32('\v'):
			v9 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v9)) = int32('\\')
			v10 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v10)) = int32('v')
			return dst
		case uint32('\t'):
			v11 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v11)) = int32('\\')
			v12 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v12)) = int32('t')
			return dst
		case uint32('\f'):
			v13 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v13)) = int32('\\')
			v14 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v14)) = int32('f')
			return dst
		case uint32(' '):
			v15 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v15)) = int32('\\')
			v16 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v16)) = int32('s')
			return dst
		case uint32('\000'):
			v17 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v17)) = int32('\\')
			v18 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v18)) = int32('0')
			if libc.Int32FromUint8(uint8(nextc)) >= int32('0') && libc.Int32FromUint8(uint8(nextc)) <= int32('7') {
				v19 = dst
				dst += 4
				*(*Twchar_t)(unsafe.Pointer(v19)) = int32('0')
				v20 = dst
				dst += 4
				*(*Twchar_t)(unsafe.Pointer(v20)) = int32('0')
			}
			return dst
			/* We cannot encode these characters in VIS_CSTYLE
			 * because they special meaning */
			fallthrough
		case uint32('n'):
			fallthrough
		case uint32('r'):
			fallthrough
		case uint32('b'):
			fallthrough
		case uint32('a'):
			fallthrough
		case uint32('v'):
			fallthrough
		case uint32('t'):
			fallthrough
		case uint32('f'):
			fallthrough
		case uint32('s'):
			fallthrough
		case uint32('0'):
			fallthrough
		case uint32('M'):
			fallthrough
		case uint32('^'):
			fallthrough
		case uint32('$'): /* vis(1) -l */
		default:
			if flags&int32(m_VIS_NOLOCALE1) != 0 {
				v21 = libc.BoolInt32(c-uint32(0x21) < uint32(0x5e))
			} else {
				v21 = libc.Xiswgraph(tls, c)
			}
			if v21 != 0 && !(libc.Int32FromUint8(uint8(c)) >= int32('0') && libc.Int32FromUint8(uint8(c)) <= int32('7')) {
				v22 = dst
				dst += 4
				*(*Twchar_t)(unsafe.Pointer(v22)) = int32('\\')
				v23 = dst
				dst += 4
				*(*Twchar_t)(unsafe.Pointer(v23)) = libc.Int32FromUint32(c)
				return dst
			}
		}
	}
	if iswextra != 0 || c&uint32(0177) == uint32(' ') || flags&int32(m_VIS_OCTAL1) != 0 {
		v24 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v24)) = int32('\\')
		v25 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v25)) = libc.Int32FromUint8(uint8(uint32(uint8(c))>>libc.Int32FromInt32(6)&libc.Uint32FromInt32(03))) + int32('0')
		v26 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v26)) = libc.Int32FromUint8(uint8(uint32(uint8(c))>>libc.Int32FromInt32(3)&libc.Uint32FromInt32(07))) + int32('0')
		v27 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v27)) = libc.Int32FromUint32(c&uint32(07) + uint32('0'))
	} else {
		if flags&int32(m_VIS_NOSLASH1) == 0 {
			v28 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v28)) = int32('\\')
		}
		if c&uint32(0200) != 0 {
			c &= uint32(0177)
			v29 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v29)) = int32('M')
		}
		if libc.Xiswcntrl(tls, c) != 0 {
			v30 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v30)) = int32('^')
			if c == uint32(0177) {
				v31 = dst
				dst += 4
				*(*Twchar_t)(unsafe.Pointer(v31)) = int32('?')
			} else {
				v32 = dst
				dst += 4
				*(*Twchar_t)(unsafe.Pointer(v32)) = libc.Int32FromUint32(c + uint32('@'))
			}
		} else {
			v33 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v33)) = int32('-')
			v34 = dst
			dst += 4
			*(*Twchar_t)(unsafe.Pointer(v34)) = libc.Int32FromUint32(c)
		}
	}
	return dst
}

// C documentation
//
//	/*
//	 * This is do_vis, the central code of vis.
//	 * dst:	      Pointer to the destination buffer
//	 * c:	      Character to encode
//	 * flags:     Flags word
//	 * nextc:     The character following 'c'
//	 * extra:     Pointer to the list of extra characters to be
//	 *	      backslash-protected.
//	 */
func _do_svis(tls *libc.TLS, dst uintptr, c Twint_t, flags int32, nextc Twint_t, extra uintptr) (r uintptr) {
	var bmsk, wmsk Tuint64_t
	var i, iswextra, shft, v1 int32
	var v2 bool
	var v3 uintptr
	_, _, _, _, _, _, _, _ = bmsk, i, iswextra, shft, wmsk, v1, v2, v3
	iswextra = libc.BoolInt32(libc.Xwcschr(tls, extra, libc.Int32FromUint32(c)) != libc.UintptrFromInt32(0))
	if v2 = !(iswextra != 0); v2 {
		if flags&int32(m_VIS_NOLOCALE1) != 0 {
			v1 = libc.BoolInt32(c-uint32(0x21) < uint32(0x5e))
		} else {
			v1 = libc.Xiswgraph(tls, c)
		}
	}
	if v2 && (v1 != 0 || (c == uint32(' ') || c == uint32('\t') || c == uint32('\n')) || flags&int32(m_VIS_SAFE1) != 0 && (c == uint32('\b') || c == uint32('\a') || c == uint32('\r'))) {
		v3 = dst
		dst += 4
		*(*Twchar_t)(unsafe.Pointer(v3)) = libc.Int32FromUint32(c)
		return dst
	}
	/* See comment in istrsenvisx() output loop, below. */
	wmsk = uint64(0)
	i = libc.Int32FromUint64(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))
	for {
		if !(i >= 0) {
			break
		}
		shft = i * int32(m_CHAR_BIT)
		bmsk = libc.Uint64FromInt64(0xff) << shft
		wmsk |= bmsk
		if uint64(c)&wmsk != 0 || i == 0 {
			dst = _do_mbyte(tls, dst, uint32(uint64(c)&bmsk>>shft), flags, nextc, iswextra)
		}
		goto _4
	_4:
		;
		i--
	}
	return dst
}

type Tvisfun_t = uintptr

// C documentation
//
//	/*
//	 * Return the appropriate encoding function depending on the flags given.
//	 */
func _getvisfun(tls *libc.TLS, flags int32) (r Tvisfun_t) {
	if flags&int32(m_VIS_HTTPSTYLE1) != 0 {
		return __ccgo_fp(_do_hvis)
	}
	if flags&int32(m_VIS_MIMESTYLE) != 0 {
		return __ccgo_fp(_do_mvis)
	}
	return __ccgo_fp(_do_svis)
}

// C documentation
//
//	/*
//	 * Expand list of extra characters to not visually encode.
//	 */
func _makeextralist(tls *libc.TLS, flags int32, src uintptr) (r uintptr) {
	var d, dst, s, v1, v10, v11, v12, v13, v4, v5, v7, v8, v9 uintptr
	var i, len1 Tsize_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d, dst, i, len1, s, v1, v10, v11, v12, v13, v4, v5, v7, v8, v9
	len1 = libc.Xstrlen(tls, src)
	v1 = libc.Xcalloc(tls, len1+uint64(m_MAXEXTRAS), uint64(4))
	dst = v1
	if v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	if flags&int32(m_VIS_NOLOCALE1) != 0 || libc.Xmbstowcs(tls, dst, src, len1) == libc.Uint64FromInt32(-libc.Int32FromInt32(1)) {
		i = uint64(0)
		for {
			if !(i < len1) {
				break
			}
			*(*Twchar_t)(unsafe.Pointer(dst + uintptr(i)*4)) = libc.Int32FromUint8(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(src + uintptr(i)))))
			goto _2
		_2:
			;
			i++
		}
		d = dst + uintptr(len1)*4
	} else {
		d = dst + uintptr(libc.Xwcslen(tls, dst))*4
	}
	if flags&int32(m_VIS_GLOB1) != 0 {
		s = uintptr(unsafe.Pointer(&_char_glob))
		for {
			if !(*(*Twchar_t)(unsafe.Pointer(s)) != 0) {
				break
			}
			goto _3
			goto _3
		_3:
			;
			v4 = d
			d += 4
			v5 = s
			s += 4
			*(*Twchar_t)(unsafe.Pointer(v4)) = *(*Twchar_t)(unsafe.Pointer(v5))
		}
	}
	if flags&int32(m_VIS_SHELL1) != 0 {
		s = uintptr(unsafe.Pointer(&_char_shell))
		for {
			if !(*(*Twchar_t)(unsafe.Pointer(s)) != 0) {
				break
			}
			goto _6
			goto _6
		_6:
			;
			v7 = d
			d += 4
			v8 = s
			s += 4
			*(*Twchar_t)(unsafe.Pointer(v7)) = *(*Twchar_t)(unsafe.Pointer(v8))
		}
	}
	if flags&int32(m_VIS_SP1) != 0 {
		v9 = d
		d += 4
		*(*Twchar_t)(unsafe.Pointer(v9)) = int32(' ')
	}
	if flags&int32(m_VIS_TAB1) != 0 {
		v10 = d
		d += 4
		*(*Twchar_t)(unsafe.Pointer(v10)) = int32('\t')
	}
	if flags&int32(m_VIS_NL1) != 0 {
		v11 = d
		d += 4
		*(*Twchar_t)(unsafe.Pointer(v11)) = int32('\n')
	}
	if flags&int32(m_VIS_DQ1) != 0 {
		v12 = d
		d += 4
		*(*Twchar_t)(unsafe.Pointer(v12)) = int32('"')
	}
	if flags&int32(m_VIS_NOSLASH1) == 0 {
		v13 = d
		d += 4
		*(*Twchar_t)(unsafe.Pointer(v13)) = int32('\\')
	}
	*(*Twchar_t)(unsafe.Pointer(d)) = int32('\000')
	return dst
}

// C documentation
//
//	/*
//	 * istrsenvisx()
//	 *	The main internal function.
//	 *	All user-visible functions call this one.
//	 */
func _istrsenvisx(tls *libc.TLS, mbdstp uintptr, dlen uintptr, mbsrc uintptr, mblength Tsize_t, flags int32, mbextra uintptr, cerr_ptr uintptr) (r int32) {
	var bmsk, wmsk Tuint64_t
	var c Twint_t
	var cerr, clen, error1, i, shft, v12, v6 int32
	var dst, extra, mbdst, mdst, pdst, psrc, src, start, v1, v2, v3, v4, v5, v8 uintptr
	var f Tvisfun_t
	var len1, olen Tsize_t
	var maxolen, mbslength Tssize_t
	var v9 uint64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bmsk, c, cerr, clen, dst, error1, extra, f, i, len1, maxolen, mbdst, mbslength, mdst, olen, pdst, psrc, shft, src, start, wmsk, v1, v12, v2, v3, v4, v5, v6, v8, v9
	clen = 0
	error1 = -int32(1)
	mbslength = libc.Int64FromUint64(mblength)
	/*
	 * When inputing a single character, must also read in the
	 * next character for nextc, the look-ahead character.
	 */
	if mbslength == int64(1) {
		mbslength++
	}
	/*
	 * Input (mbsrc) is a char string considered to be multibyte
	 * characters.  The input loop will read this string pulling
	 * one character, possibly multiple bytes, from mbsrc and
	 * converting each to wchar_t in src.
	 *
	 * The vis conversion will be done using the wide char
	 * wchar_t string.
	 *
	 * This will then be converted back to a multibyte string to
	 * return to the caller.
	 */
	/* Allocate space for the wide char strings */
	v2 = libc.UintptrFromInt32(0)
	extra = v2
	v1 = v2
	pdst = v1
	psrc = v1
	mdst = libc.UintptrFromInt32(0)
	v3 = libc.Xcalloc(tls, libc.Uint64FromInt64(mbslength+int64(1)), uint64(4))
	psrc = v3
	if v3 == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	v4 = libc.Xcalloc(tls, libc.Uint64FromInt64(int64(16)*mbslength+int64(1)), uint64(4))
	pdst = v4
	if v4 == libc.UintptrFromInt32(0) {
		goto out
	}
	if *(*uintptr)(unsafe.Pointer(mbdstp)) == libc.UintptrFromInt32(0) {
		v5 = libc.Xcalloc(tls, libc.Uint64FromInt64(int64(16)*mbslength+int64(1)), uint64(1))
		mdst = v5
		if v5 == libc.UintptrFromInt32(0) {
			goto out
		}
		*(*uintptr)(unsafe.Pointer(mbdstp)) = mdst
	}
	mbdst = *(*uintptr)(unsafe.Pointer(mbdstp))
	dst = pdst
	src = psrc
	if flags&int32(m_VIS_NOLOCALE1) != 0 {
		/* Do one byte at a time conversion */
		cerr = int32(1)
	} else {
		/* Use caller's multibyte conversion error flag. */
		if cerr_ptr != 0 {
			v6 = *(*int32)(unsafe.Pointer(cerr_ptr))
		} else {
			v6 = 0
		}
		cerr = v6
	}
	/*
	 * Input loop.
	 * Handle up to mblength characters (not bytes).  We do not
	 * stop at NULs because we may be processing a block of data
	 * that includes NULs.
	 */
	for mbslength > 0 {
		/* Convert one multibyte character to wchar_t. */
		if !(cerr != 0) {
			clen = libc.Xmbtowc(tls, src, mbsrc, uint64(m_MB_LEN_MAX))
		}
		if cerr != 0 || clen < 0 {
			/* Conversion error, process as a byte instead. */
			*(*Twchar_t)(unsafe.Pointer(src)) = libc.Int32FromUint32(uint32(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(mbsrc)))))
			clen = int32(1)
			cerr = int32(1)
		}
		if clen == 0 {
			/*
			 * NUL in input gives 0 return value. process
			 * as single NUL byte and keep going.
			 */
			clen = int32(1)
		}
		/* Advance buffer character pointer. */
		src += 4
		/* Advance input pointer by number of bytes read. */
		mbsrc += uintptr(clen)
		/* Decrement input byte count. */
		mbslength -= int64(clen)
	}
	len1 = libc.Uint64FromInt64((int64(src) - int64(psrc)) / 4)
	src = psrc
	/*
	 * In the single character input case, we will have actually
	 * processed two characters, c and nextc.  Reset len back to
	 * just a single character.
	 */
	if mblength < len1 {
		len1 = mblength
	}
	/* Convert extra argument to list of characters for this mode. */
	extra = _makeextralist(tls, flags, mbextra)
	if !(extra != 0) {
		if dlen != 0 && *(*Tsize_t)(unsafe.Pointer(dlen)) == uint64(0) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOSPC)
			goto out
		}
		*(*int8)(unsafe.Pointer(mbdst)) = int8('\000') /* can't create extra, return "" */
		error1 = 0
		goto out
	}
	/* Look up which processing function to call. */
	f = _getvisfun(tls, flags)
	/*
	 * Main processing loop.
	 * Call do_Xvis processing function one character at a time
	 * with next character available for look-ahead.
	 */
	start = dst
	for {
		if !(len1 > uint64(0)) {
			break
		}
		v8 = src
		src += 4
		c = libc.Uint32FromInt32(*(*Twchar_t)(unsafe.Pointer(v8)))
		dst = (*(*func(*libc.TLS, uintptr, Twint_t, int32, Twint_t, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{f})))(tls, dst, c, flags, libc.Uint32FromInt32(*(*Twchar_t)(unsafe.Pointer(src))), extra)
		if dst == libc.UintptrFromInt32(0) {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = int32(m_ENOSPC)
			goto out
		}
		goto _7
	_7:
		;
		len1--
	}
	/* Terminate the string in the buffer. */
	*(*Twchar_t)(unsafe.Pointer(dst)) = int32('\000')
	/*
	 * Output loop.
	 * Convert wchar_t string back to multibyte output string.
	 * If we have hit a multi-byte conversion error on input,
	 * output byte-by-byte here.  Else use wctomb().
	 */
	len1 = libc.Xwcslen(tls, start)
	if dlen != 0 {
		v9 = *(*Tsize_t)(unsafe.Pointer(dlen))
	} else {
		v9 = libc.Xwcslen(tls, start)*uint64(m_MB_LEN_MAX) + uint64(1)
	}
	maxolen = libc.Int64FromUint64(v9)
	olen = uint64(0)
	dst = start
	for {
		if !(len1 > uint64(0)) {
			break
		}
		if !(cerr != 0) {
			clen = libc.Xwctomb(tls, mbdst, *(*Twchar_t)(unsafe.Pointer(dst)))
		}
		if cerr != 0 || clen < 0 {
			/*
			 * Conversion error, process as a byte(s) instead.
			 * Examine each byte and higher-order bytes for
			 * data.  E.g.,
			 *	0x000000000000a264 -> a2 64
			 *	0x000000001f00a264 -> 1f 00 a2 64
			 */
			clen = 0
			wmsk = uint64(0)
			i = libc.Int32FromUint64(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))
			for {
				if !(i >= 0) {
					break
				}
				shft = i * int32(m_CHAR_BIT)
				bmsk = libc.Uint64FromInt64(0xff) << shft
				wmsk |= bmsk
				if libc.Uint64FromInt32(*(*Twchar_t)(unsafe.Pointer(dst)))&wmsk != 0 || i == 0 {
					v12 = clen
					clen++
					*(*int8)(unsafe.Pointer(mbdst + uintptr(v12))) = libc.Int8FromUint64(libc.Uint64FromInt32(*(*Twchar_t)(unsafe.Pointer(dst))) & bmsk >> shft)
				}
				goto _11
			_11:
				;
				i--
			}
			cerr = int32(1)
		}
		/* If this character would exceed our output limit, stop. */
		if olen+libc.Uint64FromInt32(clen) > libc.Uint64FromInt64(maxolen) {
			break
		}
		/* Advance output pointer by number of bytes written. */
		mbdst += uintptr(clen)
		/* Advance buffer character pointer. */
		dst += 4
		/* Incrment output character count. */
		olen += libc.Uint64FromInt32(clen)
		goto _10
	_10:
		;
		len1--
	}
	/* Terminate the output string. */
	*(*int8)(unsafe.Pointer(mbdst)) = int8('\000')
	if flags&int32(m_VIS_NOLOCALE1) != 0 {
		/* Pass conversion error flag out. */
		if cerr_ptr != 0 {
			*(*int32)(unsafe.Pointer(cerr_ptr)) = cerr
		}
	}
	libc.Xfree(tls, extra)
	libc.Xfree(tls, pdst)
	libc.Xfree(tls, psrc)
	return libc.Int32FromUint64(olen)
	goto out
out:
	;
	libc.Xfree(tls, extra)
	libc.Xfree(tls, pdst)
	libc.Xfree(tls, psrc)
	libc.Xfree(tls, mdst)
	return error1
}

func _istrsenvisxna(tls *libc.TLS, _mbdst uintptr, dlen uintptr, mbsrc uintptr, mblength Tsize_t, flags int32, mbextra uintptr, cerr_ptr uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uintptr)(unsafe.Pointer(bp)) = _mbdst
	return _istrsenvisx(tls, bp, dlen, mbsrc, mblength, flags, mbextra, cerr_ptr)
}

func _istrsenvisxl(tls *libc.TLS, mbdst uintptr, dlen uintptr, mbsrc uintptr, flags int32, mbextra uintptr, cerr_ptr uintptr) (r int32) {
	var v1 uint64
	_ = v1
	if mbsrc != libc.UintptrFromInt32(0) {
		v1 = libc.Xstrlen(tls, mbsrc)
	} else {
		v1 = uint64(0)
	}
	return _istrsenvisxna(tls, mbdst, dlen, mbsrc, v1, flags, mbextra, cerr_ptr)
}

/*
 *	The "svis" variants all take an "extra" arg that is a pointer
 *	to a NUL-terminated list of characters to be encoded, too.
 *	These functions are useful e. g. to encode strings in such a
 *	way so that they are not interpreted by a shell.
 */

func Xsvis(tls *libc.TLS, _mbdst uintptr, c int32, flags int32, nextc int32, mbextra uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uintptr)(unsafe.Pointer(bp)) = _mbdst
	var ret int32
	var _ /* cc at bp+8 */ [2]int8
	_ = ret
	(*(*[2]int8)(unsafe.Pointer(bp + 8)))[0] = int8(c)
	(*(*[2]int8)(unsafe.Pointer(bp + 8)))[int32(1)] = int8(nextc)
	ret = _istrsenvisx(tls, bp, libc.UintptrFromInt32(0), bp+8, uint64(1), flags, mbextra, libc.UintptrFromInt32(0))
	if ret < 0 {
		return libc.UintptrFromInt32(0)
	}
	return *(*uintptr)(unsafe.Pointer(bp)) + uintptr(ret)
}

func Xsnvis(tls *libc.TLS, _mbdst uintptr, _dlen Tsize_t, c int32, flags int32, nextc int32, mbextra uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*uintptr)(unsafe.Pointer(bp)) = _mbdst
	*(*Tsize_t)(unsafe.Pointer(bp + 8)) = _dlen
	var ret int32
	var _ /* cc at bp+16 */ [2]int8
	_ = ret
	(*(*[2]int8)(unsafe.Pointer(bp + 16)))[0] = int8(c)
	(*(*[2]int8)(unsafe.Pointer(bp + 16)))[int32(1)] = int8(nextc)
	ret = _istrsenvisx(tls, bp, bp+8, bp+16, uint64(1), flags, mbextra, libc.UintptrFromInt32(0))
	if ret < 0 {
		return libc.UintptrFromInt32(0)
	}
	return *(*uintptr)(unsafe.Pointer(bp)) + uintptr(ret)
}

func Xstrsvis(tls *libc.TLS, mbdst uintptr, mbsrc uintptr, flags int32, mbextra uintptr) (r int32) {
	return _istrsenvisxl(tls, mbdst, libc.UintptrFromInt32(0), mbsrc, flags, mbextra, libc.UintptrFromInt32(0))
}

func Xstrsnvis(tls *libc.TLS, mbdst uintptr, _dlen Tsize_t, mbsrc uintptr, flags int32, mbextra uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxl(tls, mbdst, bp, mbsrc, flags, mbextra, libc.UintptrFromInt32(0))
}

func Xstrsvisx(tls *libc.TLS, mbdst uintptr, mbsrc uintptr, len1 Tsize_t, flags int32, mbextra uintptr) (r int32) {
	return _istrsenvisxna(tls, mbdst, libc.UintptrFromInt32(0), mbsrc, len1, flags, mbextra, libc.UintptrFromInt32(0))
}

func Xstrsnvisx(tls *libc.TLS, mbdst uintptr, _dlen Tsize_t, mbsrc uintptr, len1 Tsize_t, flags int32, mbextra uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxna(tls, mbdst, bp, mbsrc, len1, flags, mbextra, libc.UintptrFromInt32(0))
}

func Xstrsenvisx(tls *libc.TLS, mbdst uintptr, _dlen Tsize_t, mbsrc uintptr, len1 Tsize_t, flags int32, mbextra uintptr, cerr_ptr uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxna(tls, mbdst, bp, mbsrc, len1, flags, mbextra, cerr_ptr)
}

// C documentation
//
//	/*
//	 * vis - visually encode characters
//	 */
func Xvis(tls *libc.TLS, _mbdst uintptr, c int32, flags int32, nextc int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uintptr)(unsafe.Pointer(bp)) = _mbdst
	var ret int32
	var _ /* cc at bp+8 */ [2]int8
	_ = ret
	(*(*[2]int8)(unsafe.Pointer(bp + 8)))[0] = int8(c)
	(*(*[2]int8)(unsafe.Pointer(bp + 8)))[int32(1)] = int8(nextc)
	ret = _istrsenvisx(tls, bp, libc.UintptrFromInt32(0), bp+8, uint64(1), flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
	if ret < 0 {
		return libc.UintptrFromInt32(0)
	}
	return *(*uintptr)(unsafe.Pointer(bp)) + uintptr(ret)
}

func Xnvis(tls *libc.TLS, _mbdst uintptr, _dlen Tsize_t, c int32, flags int32, nextc int32) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*uintptr)(unsafe.Pointer(bp)) = _mbdst
	*(*Tsize_t)(unsafe.Pointer(bp + 8)) = _dlen
	var ret int32
	var _ /* cc at bp+16 */ [2]int8
	_ = ret
	(*(*[2]int8)(unsafe.Pointer(bp + 16)))[0] = int8(c)
	(*(*[2]int8)(unsafe.Pointer(bp + 16)))[int32(1)] = int8(nextc)
	ret = _istrsenvisx(tls, bp, bp+8, bp+16, uint64(1), flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
	if ret < 0 {
		return libc.UintptrFromInt32(0)
	}
	return *(*uintptr)(unsafe.Pointer(bp)) + uintptr(ret)
}

/*
 * strvis - visually encode characters from src into dst
 *
 *	Dst must be 4 times the size of src to account for possible
 *	expansion.  The length of dst, not including the trailing NULL,
 *	is returned.
 */

func Xstrvis(tls *libc.TLS, mbdst uintptr, mbsrc uintptr, flags int32) (r int32) {
	return _istrsenvisxl(tls, mbdst, libc.UintptrFromInt32(0), mbsrc, flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
}

func Xstrnvis_openbsd(tls *libc.TLS, mbdst uintptr, mbsrc uintptr, _dlen Tsize_t, flags int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxl(tls, mbdst, bp, mbsrc, flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
}

func Xstrnvis_netbsd(tls *libc.TLS, mbdst uintptr, _dlen Tsize_t, mbsrc uintptr, flags int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxl(tls, mbdst, bp, mbsrc, flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
}

func Xstravis(tls *libc.TLS, mbdstp uintptr, mbsrc uintptr, flags int32) (r int32) {
	var v1 uint64
	_ = v1
	*(*uintptr)(unsafe.Pointer(mbdstp)) = libc.UintptrFromInt32(0)
	if mbsrc != libc.UintptrFromInt32(0) {
		v1 = libc.Xstrlen(tls, mbsrc)
	} else {
		v1 = uint64(0)
	}
	return _istrsenvisx(tls, mbdstp, libc.UintptrFromInt32(0), mbsrc, v1, flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
}

/*
 * strvisx - visually encode characters from src into dst
 *
 *	Dst must be 4 times the size of src to account for possible
 *	expansion.  The length of dst, not including the trailing NULL,
 *	is returned.
 *
 *	Strvisx encodes exactly len characters from src into dst.
 *	This is useful for encoding a block of data.
 */

func Xstrvisx(tls *libc.TLS, mbdst uintptr, mbsrc uintptr, len1 Tsize_t, flags int32) (r int32) {
	return _istrsenvisxna(tls, mbdst, libc.UintptrFromInt32(0), mbsrc, len1, flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
}

func Xstrnvisx(tls *libc.TLS, mbdst uintptr, _dlen Tsize_t, mbsrc uintptr, len1 Tsize_t, flags int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxna(tls, mbdst, bp, mbsrc, len1, flags, __ccgo_ts+60, libc.UintptrFromInt32(0))
}

func Xstrenvisx(tls *libc.TLS, mbdst uintptr, _dlen Tsize_t, mbsrc uintptr, len1 Tsize_t, flags int32, cerr_ptr uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Tsize_t)(unsafe.Pointer(bp)) = _dlen
	return _istrsenvisxna(tls, mbdst, bp, mbsrc, len1, flags, __ccgo_ts+60, cerr_ptr)
}

// C documentation
//
//	/*
//	 * Appends src to string dst of size siz (unlike wcsncat, siz is the
//	 * full size of dst, not space left).  At most siz-1 characters
//	 * will be copied.  Always NUL terminates (unless siz == 0).
//	 * Returns wcslen(initial dst) + wcslen(src); if retval >= siz,
//	 * truncation occurred.
//	 */
func Xwcslcat(tls *libc.TLS, dst uintptr, src uintptr, siz Tsize_t) (r Tsize_t) {
	var d, s, v3 uintptr
	var dlen, n, v1 Tsize_t
	var v2 bool
	_, _, _, _, _, _, _ = d, dlen, n, s, v1, v2, v3
	d = dst
	s = src
	n = siz
	/* Find the end of dst and adjust bytes left but don't go past end */
	for {
		if v2 = *(*Twchar_t)(unsafe.Pointer(d)) != int32('\000'); v2 {
			v1 = n
			n--
		}
		if !(v2 && v1 != uint64(0)) {
			break
		}
		d += 4
	}
	dlen = libc.Uint64FromInt64((int64(d) - int64(dst)) / 4)
	n = siz - dlen
	if n == uint64(0) {
		return dlen + libc.Xwcslen(tls, s)
	}
	for *(*Twchar_t)(unsafe.Pointer(s)) != int32('\000') {
		if n != uint64(1) {
			v3 = d
			d += 4
			*(*Twchar_t)(unsafe.Pointer(v3)) = *(*Twchar_t)(unsafe.Pointer(s))
			n--
		}
		s += 4
	}
	*(*Twchar_t)(unsafe.Pointer(d)) = int32('\000')
	return dlen + libc.Uint64FromInt64((int64(s)-int64(src))/4) /* count does not include NUL */
}

// C documentation
//
//	/*
//	 * Copy src to string dst of size siz.  At most siz-1 characters
//	 * will be copied.  Always NUL terminates (unless siz == 0).
//	 * Returns wcslen(src); if retval >= siz, truncation occurred.
//	 */
func Xwcslcpy(tls *libc.TLS, dst uintptr, src uintptr, siz Tsize_t) (r Tsize_t) {
	var d, s, v6, v7, v8 uintptr
	var n, v1, v3 Tsize_t
	var v2 bool
	var v5 Twchar_t
	_, _, _, _, _, _, _, _, _, _ = d, n, s, v1, v2, v3, v5, v6, v7, v8
	d = dst
	s = src
	n = siz
	/* Copy as many bytes as will fit */
	if v2 = n != uint64(0); v2 {
		n--
		v1 = n
	}
	if v2 && v1 != uint64(0) {
		for {
			v6 = s
			s += 4
			v5 = *(*Twchar_t)(unsafe.Pointer(v6))
			v7 = d
			d += 4
			*(*Twchar_t)(unsafe.Pointer(v7)) = v5
			if v5 == 0 {
				break
			}
			goto _4
		_4:
			;
			n--
			v3 = n
			if !(v3 != uint64(0)) {
				break
			}
		}
	}
	/* Not enough room in dst, add NUL and traverse rest of src */
	if n == uint64(0) {
		if siz != uint64(0) {
			*(*Twchar_t)(unsafe.Pointer(d)) = int32('\000')
		} /* NUL-terminate dst */
		for {
			v8 = s
			s += 4
			if !(*(*Twchar_t)(unsafe.Pointer(v8)) != 0) {
				break
			}
		}
	}
	return libc.Uint64FromInt64((int64(s)-int64(src))/4 - libc.Int64FromInt32(1)) /* count does not include NUL */
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var Xoptreset int32

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "/proc/self/fd\x00%s: \x00: \x00%s\n\x00diouxX\x00DOU\x00aAeEfFgG\x00#'0- +\x00r\x00r+\x00w\x00\x00BLOCKSIZE\x00G\x00K\x00M\x00%s: unknown blocksize\x00maximum blocksize is %ldG\x00minimum blocksize is 512\x00%ld%s-blocks\x00B\x00\x00Ki\x00Mi\x00Gi\x00Ti\x00Pi\x00Ei\x00\x00\x00\x00Ki\x00Mi\x00Gi\x00Ti\x00Pi\x00Ei\x00B\x00\x00k\x00\x00M\x00\x00G\x00\x00T\x00\x00P\x00\x00E\x00\x00\x00\x00k\x00\x00M\x00\x00G\x00\x00T\x00\x00P\x00\x00E\x00B\x00\x00K\x00\x00M\x00\x00G\x00\x00T\x00\x00P\x00\x00E\x00\x00\x00\x00K\x00\x00M\x00\x00G\x00\x00T\x00\x00P\x00\x00E\x00 \x00%d%s%d%s%s%s\x00%ld%s%s%s\x00/var/run/%s.pid\x00%u\x00setproctitle not initialized, please either call setproctitle_init() or link against libbsd-ctor.\x00%s\x00/dev/tty\x00\n\x00too large\x00too small\x00invalid\x000\x00\x00\x001\x00\x00\x002\x00\x00\x003\x00\x00\x004\x00\x00\x005\x00\x00\x006\x00\x00\x007\x00\x00\x008\x00\x00\x009\x00\x00\x00a\x00\x00\x00b\x00\x00\x00c\x00\x00\x00d\x00\x00\x00e\x00\x00\x00f\x00\x00\x00\x00\x00\x00\x00#\x00\x00\x00$\x00\x00\x00@\x00\x00\x00[\x00\x00\x00\\\x00\x00\x00]\x00\x00\x00^\x00\x00\x00`\x00\x00\x00{\x00\x00\x00|\x00\x00\x00}\x00\x00\x00~\x00\x00\x00\x00\x00\x00\x000\x00\x00\x001\x00\x00\x002\x00\x00\x003\x00\x00\x004\x00\x00\x005\x00\x00\x006\x00\x00\x007\x00\x00\x008\x00\x00\x009\x00\x00\x00A\x00\x00\x00B\x00\x00\x00C\x00\x00\x00D\x00\x00\x00E\x00\x00\x00F\x00\x00\x00\x00\x00\x00\x00"
