// Code generated for linux/arm64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors --libc modernc.org/libc --package-name=libz -o libz.a.go libz.a', DO NOT EDIT.

//go:build linux && arm64

package libz

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_ARG_MAX = 131072
const m_BASE = 65521
const m_BC_BASE_MAX = 99
const m_BC_DIM_MAX = 2048
const m_BC_SCALE_MAX = 99
const m_BC_STRING_MAX = 1000
const m_BIG_ENDIAN = "__BIG_ENDIAN"
const m_BYTE_ORDER = "__BYTE_ORDER"
const m_CHARCLASS_NAME_MAX = 14
const m_CHAR_BIT = 8
const m_CHAR_MAX = 255
const m_CHAR_MIN = 0
const m_COLL_WEIGHTS_MAX = 2
const m_DEF_MEM_LEVEL = 8
const m_DEF_WBITS = "MAX_WBITS"
const m_DELAYTIMER_MAX = 0x7fffffff
const m_DYN_TREES = 2
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_EXPR_NEST_MAX = 32
const m_FD_SETSIZE = 1024
const m_FILESIZEBITS = 64
const m_F_LOCK = 1
const m_F_OK = 0
const m_F_TEST = 3
const m_F_TLOCK = 2
const m_F_ULOCK = 0
const m_HAVE_HIDDEN = 1
const m_HOST_NAME_MAX = 255
const m_INT_MAX = 0x7fffffff
const m_IOV_MAX = 1024
const m_LINE_MAX = 4096
const m_LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const m_LLONG_MAX = 0x7fffffffffffffff
const m_LOGIN_NAME_MAX = 256
const m_LONG_BIT = 64
const m_LONG_MAX = "__LONG_MAX"
const m_L_INCR = 1
const m_L_SET = 0
const m_L_XTND = 2
const m_MAX_MATCH = 258
const m_MAX_MEM_LEVEL = 9
const m_MAX_WBITS = 15
const m_MB_LEN_MAX = 4
const m_MIN_MATCH = 3
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
const m_NMAX = 5552
const m_NZERO = 20
const m_OS_CODE = 3
const m_PATH_MAX = 4096
const m_PDP_ENDIAN = "__PDP_ENDIAN"
const m_PIPE_BUF = 4096
const m_POSIX_CLOSE_RESTART = 0
const m_PRESET_DICT = 0x20
const m_PTHREAD_DESTRUCTOR_ITERATIONS = 4
const m_PTHREAD_KEYS_MAX = 128
const m_PTHREAD_STACK_MIN = 2048
const m_RAND_MAX = 0x7fffffff
const m_RE_DUP_MAX = 255
const m_R_OK = 4
const m_SCHAR_MAX = 127
const m_SEEK_CUR = 1
const m_SEEK_DATA = 3
const m_SEEK_END = 2
const m_SEEK_HOLE = 4
const m_SEEK_SET = 0
const m_SEM_NSEMS_MAX = 256
const m_SEM_VALUE_MAX = 0x7fffffff
const m_SHRT_MAX = 0x7fff
const m_SSIZE_MAX = "LONG_MAX"
const m_STATIC_TREES = 1
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_STORED_BLOCK = 0
const m_SYMLOOP_MAX = 40
const m_TTY_NAME_MAX = 32
const m_TZNAME_MAX = 6
const m_UCHAR_MAX = 255
const m_UINT_MAX = 0xffffffff
const m_USHRT_MAX = 0xffff
const m_WNOHANG = 1
const m_WORD_BIT = 32
const m_WUNTRACED = 2
const m_W_OK = 2
const m_X_OK = 1
const m_ZEXTERN = "extern"
const m_ZLIB_VERNUM = 0x1310
const m_ZLIB_VERSION = "1.3.1"
const m_ZLIB_VER_MAJOR = 1
const m_ZLIB_VER_MINOR = 3
const m_ZLIB_VER_REVISION = 1
const m_ZLIB_VER_SUBREVISION = 0
const m_Z_ASCII = "Z_TEXT"
const m_Z_BEST_COMPRESSION = 9
const m_Z_BEST_SPEED = 1
const m_Z_BINARY = 0
const m_Z_BLOCK = 5
const m_Z_DEFAULT_STRATEGY = 0
const m_Z_DEFLATED = 8
const m_Z_FILTERED = 1
const m_Z_FINISH = 4
const m_Z_FIXED = 4
const m_Z_FULL_FLUSH = 3
const m_Z_HUFFMAN_ONLY = 2
const m_Z_NEED_DICT = 2
const m_Z_NO_COMPRESSION = 0
const m_Z_NO_FLUSH = 0
const m_Z_NULL = 0
const m_Z_OK = 0
const m_Z_PARTIAL_FLUSH = 1
const m_Z_RLE = 3
const m_Z_STREAM_END = 1
const m_Z_SYNC_FLUSH = 2
const m_Z_TEXT = 1
const m_Z_TREES = 6
const m_Z_U4 = "unsigned"
const m_Z_UNKNOWN = 2
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
const m__GNU_SOURCE = 1
const m__LARGEFILE64_SOURCE = 1
const m__LP64 = 1
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
const m__XOPEN_ENH_I18N = 1
const m__XOPEN_IOV_MAX = 16
const m__XOPEN_NAME_MAX = 255
const m__XOPEN_PATH_MAX = 1024
const m__XOPEN_UNIX = 1
const m__XOPEN_VERSION = 700
const m___AARCH64EL__ = 1
const m___AARCH64_CMODEL_SMALL__ = 1
const m___ARM_64BIT_STATE = 1
const m___ARM_ALIGN_MAX_PWR = 28
const m___ARM_ALIGN_MAX_STACK_PWR = 16
const m___ARM_ARCH = 8
const m___ARM_ARCH_8A = 1
const m___ARM_ARCH_ISA_A64 = 1
const m___ARM_ARCH_PROFILE = 65
const m___ARM_FEATURE_CLZ = 1
const m___ARM_FEATURE_FMA = 1
const m___ARM_FEATURE_IDIV = 1
const m___ARM_FEATURE_NUMERIC_MAXMIN = 1
const m___ARM_FEATURE_UNALIGNED = 1
const m___ARM_FP = 14
const m___ARM_FP16_ARGS = 1
const m___ARM_FP16_FORMAT_IEEE = 1
const m___ARM_NEON = 1
const m___ARM_PCS_AAPCS64 = 1
const m___ARM_SIZEOF_MINIMAL_ENUM = 4
const m___ARM_SIZEOF_WCHAR_T = 4
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
const m___CHAR_UNSIGNED__ = 1
const m___DBL_DECIMAL_DIG__ = 17
const m___DBL_DIG__ = 15
const m___DBL_HAS_DENORM__ = 1
const m___DBL_HAS_INFINITY__ = 1
const m___DBL_HAS_QUIET_NAN__ = 1
const m___DBL_IS_IEC_60559__ = 2
const m___DBL_MANT_DIG__ = 53
const m___DBL_MAX_10_EXP__ = 308
const m___DBL_MAX_EXP__ = 1024
const m___DECIMAL_DIG__ = 36
const m___DEC_EVAL_METHOD__ = 2
const m___ELF__ = 1
const m___FINITE_MATH_ONLY__ = 0
const m___FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___FLT128_DECIMAL_DIG__ = 36
const m___FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___FLT128_DIG__ = 33
const m___FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___FLT128_HAS_DENORM__ = 1
const m___FLT128_HAS_INFINITY__ = 1
const m___FLT128_HAS_QUIET_NAN__ = 1
const m___FLT128_IS_IEC_60559__ = 2
const m___FLT128_MANT_DIG__ = 113
const m___FLT128_MAX_10_EXP__ = 4932
const m___FLT128_MAX_EXP__ = 16384
const m___FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT16_DECIMAL_DIG__ = 5
const m___FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const m___FLT16_DIG__ = 3
const m___FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const m___FLT16_HAS_DENORM__ = 1
const m___FLT16_HAS_INFINITY__ = 1
const m___FLT16_HAS_QUIET_NAN__ = 1
const m___FLT16_IS_IEC_60559__ = 2
const m___FLT16_MANT_DIG__ = 11
const m___FLT16_MAX_10_EXP__ = 4
const m___FLT16_MAX_EXP__ = 16
const m___FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const m___FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const m___FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const m___FLT32X_DECIMAL_DIG__ = 17
const m___FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___FLT32X_DIG__ = 15
const m___FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___FLT32X_HAS_DENORM__ = 1
const m___FLT32X_HAS_INFINITY__ = 1
const m___FLT32X_HAS_QUIET_NAN__ = 1
const m___FLT32X_IS_IEC_60559__ = 2
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
const m___FLT32_IS_IEC_60559__ = 2
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
const m___FLT64X_IS_IEC_60559__ = 2
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
const m___FLT64_IS_IEC_60559__ = 2
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
const m___FLT_EVAL_METHOD_C99__ = 0
const m___FLT_EVAL_METHOD_TS_18661_3__ = 0
const m___FLT_EVAL_METHOD__ = 0
const m___FLT_HAS_DENORM__ = 1
const m___FLT_HAS_INFINITY__ = 1
const m___FLT_HAS_QUIET_NAN__ = 1
const m___FLT_IS_IEC_60559__ = 2
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
const m___GCC_DESTRUCTIVE_SIZE = 256
const m___GCC_HAVE_DWARF2_CFI_ASM = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const m___GCC_IEC_559 = 2
const m___GCC_IEC_559_COMPLEX = 2
const m___GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const m___GNUC_MINOR__ = 2
const m___GNUC_PATCHLEVEL__ = 0
const m___GNUC_STDC_INLINE__ = 1
const m___GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const m___GNUC__ = 12
const m___GXX_ABI_VERSION = 1017
const m___HAVE_SPECULATION_SAFE_VALUE = 1
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
const m___LDBL_DECIMAL_DIG__ = 36
const m___LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___LDBL_DIG__ = 33
const m___LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___LDBL_HAS_DENORM__ = 1
const m___LDBL_HAS_INFINITY__ = 1
const m___LDBL_HAS_QUIET_NAN__ = 1
const m___LDBL_IS_IEC_60559__ = 2
const m___LDBL_MANT_DIG__ = 113
const m___LDBL_MAX_10_EXP__ = 4932
const m___LDBL_MAX_EXP__ = 16384
const m___LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___LDBL_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___LITTLE_ENDIAN = 1234
const m___LONG_LONG_MAX__ = 0x7fffffffffffffff
const m___LONG_LONG_WIDTH__ = 64
const m___LONG_MAX = 0x7fffffffffffffff
const m___LONG_MAX__ = 0x7fffffffffffffff
const m___LONG_WIDTH__ = 64
const m___LP64__ = 1
const m___NO_INLINE__ = 1
const m___ORDER_BIG_ENDIAN__ = 4321
const m___ORDER_LITTLE_ENDIAN__ = 1234
const m___ORDER_PDP_ENDIAN__ = 3412
const m___PDP_ENDIAN = 3412
const m___PIC__ = 2
const m___PIE__ = 2
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTRDIFF_MAX__ = 0x7fffffffffffffff
const m___PTRDIFF_WIDTH__ = 64
const m___SCHAR_MAX__ = 0x7f
const m___SCHAR_WIDTH__ = 8
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
const m___STDC_HOSTED__ = 1
const m___STDC_IEC_559_COMPLEX__ = 1
const m___STDC_IEC_559__ = 1
const m___STDC_IEC_60559_BFP__ = 201404
const m___STDC_IEC_60559_COMPLEX__ = 201404
const m___STDC_ISO_10646__ = 201706
const m___STDC_UTF_16__ = 1
const m___STDC_UTF_32__ = 1
const m___STDC_VERSION__ = 201710
const m___STDC__ = 1
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
const m___USE_TIME_BITS64 = 1
const m___VERSION__ = "12.2.0"
const m___WCHAR_MAX__ = 0xffffffff
const m___WCHAR_MIN__ = 0
const m___WCHAR_WIDTH__ = 32
const m___WINT_MAX__ = 0xffffffff
const m___WINT_MIN__ = 0
const m___WINT_WIDTH__ = 32
const m___aarch64__ = 1
const m___gnu_linux__ = 1
const m___inline = "inline"
const m___linux = 1
const m___linux__ = 1
const m___pic__ = 2
const m___pie__ = 2
const m___restrict = "restrict"
const m___restrict_arr = "restrict"
const m___unix = 1
const m___unix__ = 1
const m_alloca = "__builtin_alloca"
const m_blkcnt64_t = "blkcnt_t"
const m_fsblkcnt64_t = "fsblkcnt_t"
const m_fsfilcnt64_t = "fsfilcnt_t"
const m_ftruncate64 = "ftruncate"
const m_ino64_t = "ino_t"
const m_linux = 1
const m_local = "static"
const m_lockf64 = "lockf"
const m_lseek64 = "lseek"
const m_mkostemp64 = "mkostemp"
const m_mkostemps64 = "mkostemps"
const m_mkstemp64 = "mkstemp"
const m_mkstemps64 = "mkstemps"
const m_off64_t = "off_t"
const m_pread64 = "pread"
const m_pwrite64 = "pwrite"
const m_truncate64 = "truncate"
const m_unix = 1
const m_z_off64_t = "z_off_t"
const m_z_off_t = "off_t"
const m_zmemcmp = "memcmp"
const m_zmemcpy = "memcpy"

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint64

type t__predefined_wchar_t = uint32

type t__predefined_ptrdiff_t = int64

type Twchar_t = uint32

type Tmax_align_t = struct {
	F__ll int64
	F__ld float64
}

type Tsize_t = uint64

type Tptrdiff_t = int64

type Tz_size_t = uint64

type TByte = uint8

type TuInt = uint32

type TuLong = uint64

type TBytef = uint8

type Tcharf = uint8

type Tintf = int32

type TuIntf = uint32

type TuLongf = uint64

type Tvoidpc = uintptr

type Tvoidpf = uintptr

type Tvoidp = uintptr

type Tz_crc_t = uint32

type Tblksize_t = int32

type Tnlink_t = uint32

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

type Tva_list = uintptr

type Tintptr_t = int64

type Talloc_func = uintptr

type Tfree_func = uintptr

type Tz_stream = struct {
	Fnext_in   uintptr
	Favail_in  TuInt
	Ftotal_in  TuLong
	Fnext_out  uintptr
	Favail_out TuInt
	Ftotal_out TuLong
	Fmsg       uintptr
	Fstate     uintptr
	Fzalloc    Talloc_func
	Fzfree     Tfree_func
	Fopaque    Tvoidpf
	Fdata_type int32
	Fadler     TuLong
	Freserved  TuLong
}

type Tz_stream_s = Tz_stream

type Tz_streamp = uintptr

type Tgz_header = struct {
	Ftext      int32
	Ftime      TuLong
	Fxflags    int32
	Fos        int32
	Fextra     uintptr
	Fextra_len TuInt
	Fextra_max TuInt
	Fname      uintptr
	Fname_max  TuInt
	Fcomment   uintptr
	Fcomm_max  TuInt
	Fhcrc      int32
	Fdone      int32
}

type Tgz_header_s = Tgz_header

type Tgz_headerp = uintptr

type Tin_func = uintptr

type Tout_func = uintptr

type TgzFile = uintptr

type TgzFile_s = struct {
	Fhave uint32
	Fnext uintptr
	Fpos  Toff_t
}

type Tlocale_t = uintptr

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

type Tuch = uint8

type Tuchf = uint8

type Tush = uint16

type Tushf = uint16

type Tulg = uint64

/* Reverse the bytes in a 32-bit value */

/* NMAX is the largest n such that 255n(n+1)/2 + (n+1)(BASE-1) <= 2^32-1 */

/* use NO_DIVIDE if your processor does not do division in hardware --
   try it both ways to see which is faster */

// C documentation
//
//	/* ========================================================================= */
func Xadler32_z(tls *libc.TLS, adler TuLong, buf uintptr, len1 Tz_size_t) (r TuLong) {
	var n, v3 uint32
	var sum2 uint64
	var v1, v5 Tz_size_t
	var v2, v6 uintptr
	_, _, _, _, _, _, _ = n, sum2, v1, v2, v3, v5, v6
	/* split Adler-32 into component sums */
	sum2 = adler >> libc.Int32FromInt32(16) & uint64(0xffff)
	adler &= uint64(0xffff)
	/* in case user likes doing a byte at a time, keep it fast */
	if len1 == uint64(1) {
		adler += uint64(*(*TBytef)(unsafe.Pointer(buf)))
		if adler >= uint64(65521) {
			adler -= uint64(65521)
		}
		sum2 += adler
		if sum2 >= uint64(65521) {
			sum2 -= uint64(65521)
		}
		return adler | sum2<<int32(16)
	}
	/* initial Adler-32 value (deferred check for len == 1 speed) */
	if buf == uintptr(m_Z_NULL) {
		return uint64(1)
	}
	/* in case short lengths are provided, keep it somewhat fast */
	if len1 < uint64(16) {
		for {
			v1 = len1
			len1--
			if !(v1 != 0) {
				break
			}
			v2 = buf
			buf++
			adler += uint64(*(*TBytef)(unsafe.Pointer(v2)))
			sum2 += adler
		}
		if adler >= uint64(65521) {
			adler -= uint64(65521)
		}
		sum2 %= uint64(65521) /* only added so many BASE's */
		return adler | sum2<<int32(16)
	}
	/* do length NMAX blocks -- requires just one modulo operation */
	for len1 >= uint64(m_NMAX) {
		len1 -= uint64(m_NMAX)
		n = libc.Uint32FromInt32(libc.Int32FromInt32(m_NMAX) / libc.Int32FromInt32(16)) /* NMAX is divisible by 16 */
		for {
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf)))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + 8)))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler /* 16 sums unrolled */
			buf += uintptr(16)
			goto _4
		_4:
			;
			n--
			v3 = n
			if !(v3 != 0) {
				break
			}
		}
		adler %= uint64(65521)
		sum2 %= uint64(65521)
	}
	/* do remaining bytes (less than NMAX, still just one modulo) */
	if len1 != 0 { /* avoid modulos if none remaining */
		for len1 >= uint64(16) {
			len1 -= uint64(16)
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf)))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + 8)))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint64(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			buf += uintptr(16)
		}
		for {
			v5 = len1
			len1--
			if !(v5 != 0) {
				break
			}
			v6 = buf
			buf++
			adler += uint64(*(*TBytef)(unsafe.Pointer(v6)))
			sum2 += adler
		}
		adler %= uint64(65521)
		sum2 %= uint64(65521)
	}
	/* return recombined sums */
	return adler | sum2<<int32(16)
}

// C documentation
//
//	/* ========================================================================= */
func Xadler32(tls *libc.TLS, adler TuLong, buf uintptr, len1 TuInt) (r TuLong) {
	return Xadler32_z(tls, adler, buf, uint64(len1))
}

// C documentation
//
//	/* ========================================================================= */
func _adler32_combine_(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	var rem uint32
	var sum1, sum2 uint64
	_, _, _ = rem, sum1, sum2
	/* for negative len, return invalid adler32 as a clue for debugging */
	if len2 < 0 {
		return uint64(0xffffffff)
	}
	/* the derivation of this formula is left as an exercise for the reader */
	len2 %= libc.Int64FromUint32(65521) /* assumes len2 >= 0 */
	rem = libc.Uint32FromInt64(len2)
	sum1 = adler1 & uint64(0xffff)
	sum2 = uint64(rem) * sum1
	sum2 %= uint64(65521)
	sum1 += adler2&uint64(0xffff) + uint64(65521) - uint64(1)
	sum2 += adler1>>libc.Int32FromInt32(16)&uint64(0xffff) + adler2>>libc.Int32FromInt32(16)&uint64(0xffff) + uint64(65521) - uint64(rem)
	if sum1 >= uint64(65521) {
		sum1 -= uint64(65521)
	}
	if sum1 >= uint64(65521) {
		sum1 -= uint64(65521)
	}
	if sum2 >= libc.Uint64FromUint32(65521)<<libc.Int32FromInt32(1) {
		sum2 -= libc.Uint64FromUint32(65521) << libc.Int32FromInt32(1)
	}
	if sum2 >= uint64(65521) {
		sum2 -= uint64(65521)
	}
	return sum1 | sum2<<int32(16)
}

// C documentation
//
//	/* ========================================================================= */
func Xadler32_combine(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	return _adler32_combine_(tls, adler1, adler2, len2)
}

func Xadler32_combine64(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	return _adler32_combine_(tls, adler1, adler2, len2)
}

const m_N = 5
const m_POLY = 3988292384
const m_W = 8

/* Reverse the bytes in a 32-bit value */

/*
  A CRC of a message is computed on N braids of words in the message, where
  each word consists of W bytes (4 or 8). If N is 3, for example, then three
  running sparse CRCs are calculated respectively on each braid, at these
  indices in the array of words: 0, 3, 6, ..., 1, 4, 7, ..., and 2, 5, 8, ...
  This is done starting at a word boundary, and continues until as many blocks
  of N * W bytes as are available have been processed. The results are combined
  into a single CRC at the end. For this code, N must be in the range 1..6 and
  W must be 4 or 8. The upper limit on N can be increased if desired by adding
  more #if blocks, extending the patterns apparent in the code. In addition,
  crc32.h would need to be regenerated, if the maximum N value is increased.

  N and W are chosen empirically by benchmarking the execution time on a given
  processor. The choices for N and W below were based on testing on Intel Kaby
  Lake i7, AMD Ryzen 7, ARM Cortex-A57, Sparc64-VII, PowerPC POWER9, and MIPS64
  Octeon II processors. The Intel, AMD, and ARM processors were all fastest
  with N=5, W=8. The Sparc, PowerPC, and MIPS64 were all fastest at N=5, W=4.
  They were all tested with either gcc or clang, all using the -O3 optimization
  level. Your mileage may vary.
*/

/* Define N */

/*
  z_crc_t must be at least 32 bits. z_word_t must be at least as long as
  z_crc_t. It is assumed here that z_word_t is either 32 bits or 64 bits, and
  that bytes are eight bits.
*/

// C documentation
//
//	/*
//	  Define W and the associated z_word_t type. If W is not defined, then a
//	  braided calculation is not used, and the associated tables and code are not
//	  compiled.
//	 */
type Tz_word_t = uint64

/* If available, use the ARM processor CRC32 instruction. */

// C documentation
//
//	/*
//	  Swap the bytes in a z_word_t to convert between little and big endian. Any
//	  self-respecting compiler will optimize this to a single machine byte-swap
//	  instruction, if one is available. This assumes that word_t is either 32 bits
//	  or 64 bits.
//	 */
func _byte_swap(tls *libc.TLS, word Tz_word_t) (r Tz_word_t) {
	return word&uint64(0xff00000000000000)>>int32(56) | word&uint64(0xff000000000000)>>int32(40) | word&uint64(0xff0000000000)>>int32(24) | word&uint64(0xff00000000)>>int32(8) | word&uint64(0xff000000)<<int32(8) | word&uint64(0xff0000)<<int32(24) | word&uint64(0xff00)<<int32(40) | word&uint64(0xff)<<int32(56)
}

var _crc_table = [256]Tz_crc_t{
	1:   uint32(0x77073096),
	2:   uint32(0xee0e612c),
	3:   uint32(0x990951ba),
	4:   uint32(0x076dc419),
	5:   uint32(0x706af48f),
	6:   uint32(0xe963a535),
	7:   uint32(0x9e6495a3),
	8:   uint32(0x0edb8832),
	9:   uint32(0x79dcb8a4),
	10:  uint32(0xe0d5e91e),
	11:  uint32(0x97d2d988),
	12:  uint32(0x09b64c2b),
	13:  uint32(0x7eb17cbd),
	14:  uint32(0xe7b82d07),
	15:  uint32(0x90bf1d91),
	16:  uint32(0x1db71064),
	17:  uint32(0x6ab020f2),
	18:  uint32(0xf3b97148),
	19:  uint32(0x84be41de),
	20:  uint32(0x1adad47d),
	21:  uint32(0x6ddde4eb),
	22:  uint32(0xf4d4b551),
	23:  uint32(0x83d385c7),
	24:  uint32(0x136c9856),
	25:  uint32(0x646ba8c0),
	26:  uint32(0xfd62f97a),
	27:  uint32(0x8a65c9ec),
	28:  uint32(0x14015c4f),
	29:  uint32(0x63066cd9),
	30:  uint32(0xfa0f3d63),
	31:  uint32(0x8d080df5),
	32:  uint32(0x3b6e20c8),
	33:  uint32(0x4c69105e),
	34:  uint32(0xd56041e4),
	35:  uint32(0xa2677172),
	36:  uint32(0x3c03e4d1),
	37:  uint32(0x4b04d447),
	38:  uint32(0xd20d85fd),
	39:  uint32(0xa50ab56b),
	40:  uint32(0x35b5a8fa),
	41:  uint32(0x42b2986c),
	42:  uint32(0xdbbbc9d6),
	43:  uint32(0xacbcf940),
	44:  uint32(0x32d86ce3),
	45:  uint32(0x45df5c75),
	46:  uint32(0xdcd60dcf),
	47:  uint32(0xabd13d59),
	48:  uint32(0x26d930ac),
	49:  uint32(0x51de003a),
	50:  uint32(0xc8d75180),
	51:  uint32(0xbfd06116),
	52:  uint32(0x21b4f4b5),
	53:  uint32(0x56b3c423),
	54:  uint32(0xcfba9599),
	55:  uint32(0xb8bda50f),
	56:  uint32(0x2802b89e),
	57:  uint32(0x5f058808),
	58:  uint32(0xc60cd9b2),
	59:  uint32(0xb10be924),
	60:  uint32(0x2f6f7c87),
	61:  uint32(0x58684c11),
	62:  uint32(0xc1611dab),
	63:  uint32(0xb6662d3d),
	64:  uint32(0x76dc4190),
	65:  uint32(0x01db7106),
	66:  uint32(0x98d220bc),
	67:  uint32(0xefd5102a),
	68:  uint32(0x71b18589),
	69:  uint32(0x06b6b51f),
	70:  uint32(0x9fbfe4a5),
	71:  uint32(0xe8b8d433),
	72:  uint32(0x7807c9a2),
	73:  uint32(0x0f00f934),
	74:  uint32(0x9609a88e),
	75:  uint32(0xe10e9818),
	76:  uint32(0x7f6a0dbb),
	77:  uint32(0x086d3d2d),
	78:  uint32(0x91646c97),
	79:  uint32(0xe6635c01),
	80:  uint32(0x6b6b51f4),
	81:  uint32(0x1c6c6162),
	82:  uint32(0x856530d8),
	83:  uint32(0xf262004e),
	84:  uint32(0x6c0695ed),
	85:  uint32(0x1b01a57b),
	86:  uint32(0x8208f4c1),
	87:  uint32(0xf50fc457),
	88:  uint32(0x65b0d9c6),
	89:  uint32(0x12b7e950),
	90:  uint32(0x8bbeb8ea),
	91:  uint32(0xfcb9887c),
	92:  uint32(0x62dd1ddf),
	93:  uint32(0x15da2d49),
	94:  uint32(0x8cd37cf3),
	95:  uint32(0xfbd44c65),
	96:  uint32(0x4db26158),
	97:  uint32(0x3ab551ce),
	98:  uint32(0xa3bc0074),
	99:  uint32(0xd4bb30e2),
	100: uint32(0x4adfa541),
	101: uint32(0x3dd895d7),
	102: uint32(0xa4d1c46d),
	103: uint32(0xd3d6f4fb),
	104: uint32(0x4369e96a),
	105: uint32(0x346ed9fc),
	106: uint32(0xad678846),
	107: uint32(0xda60b8d0),
	108: uint32(0x44042d73),
	109: uint32(0x33031de5),
	110: uint32(0xaa0a4c5f),
	111: uint32(0xdd0d7cc9),
	112: uint32(0x5005713c),
	113: uint32(0x270241aa),
	114: uint32(0xbe0b1010),
	115: uint32(0xc90c2086),
	116: uint32(0x5768b525),
	117: uint32(0x206f85b3),
	118: uint32(0xb966d409),
	119: uint32(0xce61e49f),
	120: uint32(0x5edef90e),
	121: uint32(0x29d9c998),
	122: uint32(0xb0d09822),
	123: uint32(0xc7d7a8b4),
	124: uint32(0x59b33d17),
	125: uint32(0x2eb40d81),
	126: uint32(0xb7bd5c3b),
	127: uint32(0xc0ba6cad),
	128: uint32(0xedb88320),
	129: uint32(0x9abfb3b6),
	130: uint32(0x03b6e20c),
	131: uint32(0x74b1d29a),
	132: uint32(0xead54739),
	133: uint32(0x9dd277af),
	134: uint32(0x04db2615),
	135: uint32(0x73dc1683),
	136: uint32(0xe3630b12),
	137: uint32(0x94643b84),
	138: uint32(0x0d6d6a3e),
	139: uint32(0x7a6a5aa8),
	140: uint32(0xe40ecf0b),
	141: uint32(0x9309ff9d),
	142: uint32(0x0a00ae27),
	143: uint32(0x7d079eb1),
	144: uint32(0xf00f9344),
	145: uint32(0x8708a3d2),
	146: uint32(0x1e01f268),
	147: uint32(0x6906c2fe),
	148: uint32(0xf762575d),
	149: uint32(0x806567cb),
	150: uint32(0x196c3671),
	151: uint32(0x6e6b06e7),
	152: uint32(0xfed41b76),
	153: uint32(0x89d32be0),
	154: uint32(0x10da7a5a),
	155: uint32(0x67dd4acc),
	156: uint32(0xf9b9df6f),
	157: uint32(0x8ebeeff9),
	158: uint32(0x17b7be43),
	159: uint32(0x60b08ed5),
	160: uint32(0xd6d6a3e8),
	161: uint32(0xa1d1937e),
	162: uint32(0x38d8c2c4),
	163: uint32(0x4fdff252),
	164: uint32(0xd1bb67f1),
	165: uint32(0xa6bc5767),
	166: uint32(0x3fb506dd),
	167: uint32(0x48b2364b),
	168: uint32(0xd80d2bda),
	169: uint32(0xaf0a1b4c),
	170: uint32(0x36034af6),
	171: uint32(0x41047a60),
	172: uint32(0xdf60efc3),
	173: uint32(0xa867df55),
	174: uint32(0x316e8eef),
	175: uint32(0x4669be79),
	176: uint32(0xcb61b38c),
	177: uint32(0xbc66831a),
	178: uint32(0x256fd2a0),
	179: uint32(0x5268e236),
	180: uint32(0xcc0c7795),
	181: uint32(0xbb0b4703),
	182: uint32(0x220216b9),
	183: uint32(0x5505262f),
	184: uint32(0xc5ba3bbe),
	185: uint32(0xb2bd0b28),
	186: uint32(0x2bb45a92),
	187: uint32(0x5cb36a04),
	188: uint32(0xc2d7ffa7),
	189: uint32(0xb5d0cf31),
	190: uint32(0x2cd99e8b),
	191: uint32(0x5bdeae1d),
	192: uint32(0x9b64c2b0),
	193: uint32(0xec63f226),
	194: uint32(0x756aa39c),
	195: uint32(0x026d930a),
	196: uint32(0x9c0906a9),
	197: uint32(0xeb0e363f),
	198: uint32(0x72076785),
	199: uint32(0x05005713),
	200: uint32(0x95bf4a82),
	201: uint32(0xe2b87a14),
	202: uint32(0x7bb12bae),
	203: uint32(0x0cb61b38),
	204: uint32(0x92d28e9b),
	205: uint32(0xe5d5be0d),
	206: uint32(0x7cdcefb7),
	207: uint32(0x0bdbdf21),
	208: uint32(0x86d3d2d4),
	209: uint32(0xf1d4e242),
	210: uint32(0x68ddb3f8),
	211: uint32(0x1fda836e),
	212: uint32(0x81be16cd),
	213: uint32(0xf6b9265b),
	214: uint32(0x6fb077e1),
	215: uint32(0x18b74777),
	216: uint32(0x88085ae6),
	217: uint32(0xff0f6a70),
	218: uint32(0x66063bca),
	219: uint32(0x11010b5c),
	220: uint32(0x8f659eff),
	221: uint32(0xf862ae69),
	222: uint32(0x616bffd3),
	223: uint32(0x166ccf45),
	224: uint32(0xa00ae278),
	225: uint32(0xd70dd2ee),
	226: uint32(0x4e048354),
	227: uint32(0x3903b3c2),
	228: uint32(0xa7672661),
	229: uint32(0xd06016f7),
	230: uint32(0x4969474d),
	231: uint32(0x3e6e77db),
	232: uint32(0xaed16a4a),
	233: uint32(0xd9d65adc),
	234: uint32(0x40df0b66),
	235: uint32(0x37d83bf0),
	236: uint32(0xa9bcae53),
	237: uint32(0xdebb9ec5),
	238: uint32(0x47b2cf7f),
	239: uint32(0x30b5ffe9),
	240: uint32(0xbdbdf21c),
	241: uint32(0xcabac28a),
	242: uint32(0x53b39330),
	243: uint32(0x24b4a3a6),
	244: uint32(0xbad03605),
	245: uint32(0xcdd70693),
	246: uint32(0x54de5729),
	247: uint32(0x23d967bf),
	248: uint32(0xb3667a2e),
	249: uint32(0xc4614ab8),
	250: uint32(0x5d681b02),
	251: uint32(0x2a6f2b94),
	252: uint32(0xb40bbe37),
	253: uint32(0xc30c8ea1),
	254: uint32(0x5a05df1b),
	255: uint32(0x2d02ef8d),
}
var _crc_big_table = [256]Tz_word_t{
	1:   uint64(0x9630077700000000),
	2:   uint64(0x2c610eee00000000),
	3:   uint64(0xba51099900000000),
	4:   uint64(0x19c46d0700000000),
	5:   uint64(0x8ff46a7000000000),
	6:   uint64(0x35a563e900000000),
	7:   uint64(0xa395649e00000000),
	8:   uint64(0x3288db0e00000000),
	9:   uint64(0xa4b8dc7900000000),
	10:  uint64(0x1ee9d5e000000000),
	11:  uint64(0x88d9d29700000000),
	12:  uint64(0x2b4cb60900000000),
	13:  uint64(0xbd7cb17e00000000),
	14:  uint64(0x072db8e700000000),
	15:  uint64(0x911dbf9000000000),
	16:  uint64(0x6410b71d00000000),
	17:  uint64(0xf220b06a00000000),
	18:  uint64(0x4871b9f300000000),
	19:  uint64(0xde41be8400000000),
	20:  uint64(0x7dd4da1a00000000),
	21:  uint64(0xebe4dd6d00000000),
	22:  uint64(0x51b5d4f400000000),
	23:  uint64(0xc785d38300000000),
	24:  uint64(0x56986c1300000000),
	25:  uint64(0xc0a86b6400000000),
	26:  uint64(0x7af962fd00000000),
	27:  uint64(0xecc9658a00000000),
	28:  uint64(0x4f5c011400000000),
	29:  uint64(0xd96c066300000000),
	30:  uint64(0x633d0ffa00000000),
	31:  uint64(0xf50d088d00000000),
	32:  uint64(0xc8206e3b00000000),
	33:  uint64(0x5e10694c00000000),
	34:  uint64(0xe44160d500000000),
	35:  uint64(0x727167a200000000),
	36:  uint64(0xd1e4033c00000000),
	37:  uint64(0x47d4044b00000000),
	38:  uint64(0xfd850dd200000000),
	39:  uint64(0x6bb50aa500000000),
	40:  uint64(0xfaa8b53500000000),
	41:  uint64(0x6c98b24200000000),
	42:  uint64(0xd6c9bbdb00000000),
	43:  uint64(0x40f9bcac00000000),
	44:  uint64(0xe36cd83200000000),
	45:  uint64(0x755cdf4500000000),
	46:  uint64(0xcf0dd6dc00000000),
	47:  uint64(0x593dd1ab00000000),
	48:  uint64(0xac30d92600000000),
	49:  uint64(0x3a00de5100000000),
	50:  uint64(0x8051d7c800000000),
	51:  uint64(0x1661d0bf00000000),
	52:  uint64(0xb5f4b42100000000),
	53:  uint64(0x23c4b35600000000),
	54:  uint64(0x9995bacf00000000),
	55:  uint64(0x0fa5bdb800000000),
	56:  uint64(0x9eb8022800000000),
	57:  uint64(0x0888055f00000000),
	58:  uint64(0xb2d90cc600000000),
	59:  uint64(0x24e90bb100000000),
	60:  uint64(0x877c6f2f00000000),
	61:  uint64(0x114c685800000000),
	62:  uint64(0xab1d61c100000000),
	63:  uint64(0x3d2d66b600000000),
	64:  uint64(0x9041dc7600000000),
	65:  uint64(0x0671db0100000000),
	66:  uint64(0xbc20d29800000000),
	67:  uint64(0x2a10d5ef00000000),
	68:  uint64(0x8985b17100000000),
	69:  uint64(0x1fb5b60600000000),
	70:  uint64(0xa5e4bf9f00000000),
	71:  uint64(0x33d4b8e800000000),
	72:  uint64(0xa2c9077800000000),
	73:  uint64(0x34f9000f00000000),
	74:  uint64(0x8ea8099600000000),
	75:  uint64(0x18980ee100000000),
	76:  uint64(0xbb0d6a7f00000000),
	77:  uint64(0x2d3d6d0800000000),
	78:  uint64(0x976c649100000000),
	79:  uint64(0x015c63e600000000),
	80:  uint64(0xf4516b6b00000000),
	81:  uint64(0x62616c1c00000000),
	82:  uint64(0xd830658500000000),
	83:  uint64(0x4e0062f200000000),
	84:  uint64(0xed95066c00000000),
	85:  uint64(0x7ba5011b00000000),
	86:  uint64(0xc1f4088200000000),
	87:  uint64(0x57c40ff500000000),
	88:  uint64(0xc6d9b06500000000),
	89:  uint64(0x50e9b71200000000),
	90:  uint64(0xeab8be8b00000000),
	91:  uint64(0x7c88b9fc00000000),
	92:  uint64(0xdf1ddd6200000000),
	93:  uint64(0x492dda1500000000),
	94:  uint64(0xf37cd38c00000000),
	95:  uint64(0x654cd4fb00000000),
	96:  uint64(0x5861b24d00000000),
	97:  uint64(0xce51b53a00000000),
	98:  uint64(0x7400bca300000000),
	99:  uint64(0xe230bbd400000000),
	100: uint64(0x41a5df4a00000000),
	101: uint64(0xd795d83d00000000),
	102: uint64(0x6dc4d1a400000000),
	103: uint64(0xfbf4d6d300000000),
	104: uint64(0x6ae9694300000000),
	105: uint64(0xfcd96e3400000000),
	106: uint64(0x468867ad00000000),
	107: uint64(0xd0b860da00000000),
	108: uint64(0x732d044400000000),
	109: uint64(0xe51d033300000000),
	110: uint64(0x5f4c0aaa00000000),
	111: uint64(0xc97c0ddd00000000),
	112: uint64(0x3c71055000000000),
	113: uint64(0xaa41022700000000),
	114: uint64(0x10100bbe00000000),
	115: uint64(0x86200cc900000000),
	116: uint64(0x25b5685700000000),
	117: uint64(0xb3856f2000000000),
	118: uint64(0x09d466b900000000),
	119: uint64(0x9fe461ce00000000),
	120: uint64(0x0ef9de5e00000000),
	121: uint64(0x98c9d92900000000),
	122: uint64(0x2298d0b000000000),
	123: uint64(0xb4a8d7c700000000),
	124: uint64(0x173db35900000000),
	125: uint64(0x810db42e00000000),
	126: uint64(0x3b5cbdb700000000),
	127: uint64(0xad6cbac000000000),
	128: uint64(0x2083b8ed00000000),
	129: uint64(0xb6b3bf9a00000000),
	130: uint64(0x0ce2b60300000000),
	131: uint64(0x9ad2b17400000000),
	132: uint64(0x3947d5ea00000000),
	133: uint64(0xaf77d29d00000000),
	134: uint64(0x1526db0400000000),
	135: uint64(0x8316dc7300000000),
	136: uint64(0x120b63e300000000),
	137: uint64(0x843b649400000000),
	138: uint64(0x3e6a6d0d00000000),
	139: uint64(0xa85a6a7a00000000),
	140: uint64(0x0bcf0ee400000000),
	141: uint64(0x9dff099300000000),
	142: uint64(0x27ae000a00000000),
	143: uint64(0xb19e077d00000000),
	144: uint64(0x44930ff000000000),
	145: uint64(0xd2a3088700000000),
	146: uint64(0x68f2011e00000000),
	147: uint64(0xfec2066900000000),
	148: uint64(0x5d5762f700000000),
	149: uint64(0xcb67658000000000),
	150: uint64(0x71366c1900000000),
	151: uint64(0xe7066b6e00000000),
	152: uint64(0x761bd4fe00000000),
	153: uint64(0xe02bd38900000000),
	154: uint64(0x5a7ada1000000000),
	155: uint64(0xcc4add6700000000),
	156: uint64(0x6fdfb9f900000000),
	157: uint64(0xf9efbe8e00000000),
	158: uint64(0x43beb71700000000),
	159: uint64(0xd58eb06000000000),
	160: uint64(0xe8a3d6d600000000),
	161: uint64(0x7e93d1a100000000),
	162: uint64(0xc4c2d83800000000),
	163: uint64(0x52f2df4f00000000),
	164: uint64(0xf167bbd100000000),
	165: uint64(0x6757bca600000000),
	166: uint64(0xdd06b53f00000000),
	167: uint64(0x4b36b24800000000),
	168: uint64(0xda2b0dd800000000),
	169: uint64(0x4c1b0aaf00000000),
	170: uint64(0xf64a033600000000),
	171: uint64(0x607a044100000000),
	172: uint64(0xc3ef60df00000000),
	173: uint64(0x55df67a800000000),
	174: uint64(0xef8e6e3100000000),
	175: uint64(0x79be694600000000),
	176: uint64(0x8cb361cb00000000),
	177: uint64(0x1a8366bc00000000),
	178: uint64(0xa0d26f2500000000),
	179: uint64(0x36e2685200000000),
	180: uint64(0x95770ccc00000000),
	181: uint64(0x03470bbb00000000),
	182: uint64(0xb916022200000000),
	183: uint64(0x2f26055500000000),
	184: uint64(0xbe3bbac500000000),
	185: uint64(0x280bbdb200000000),
	186: uint64(0x925ab42b00000000),
	187: uint64(0x046ab35c00000000),
	188: uint64(0xa7ffd7c200000000),
	189: uint64(0x31cfd0b500000000),
	190: uint64(0x8b9ed92c00000000),
	191: uint64(0x1daede5b00000000),
	192: uint64(0xb0c2649b00000000),
	193: uint64(0x26f263ec00000000),
	194: uint64(0x9ca36a7500000000),
	195: uint64(0x0a936d0200000000),
	196: uint64(0xa906099c00000000),
	197: uint64(0x3f360eeb00000000),
	198: uint64(0x8567077200000000),
	199: uint64(0x1357000500000000),
	200: uint64(0x824abf9500000000),
	201: uint64(0x147ab8e200000000),
	202: uint64(0xae2bb17b00000000),
	203: uint64(0x381bb60c00000000),
	204: uint64(0x9b8ed29200000000),
	205: uint64(0x0dbed5e500000000),
	206: uint64(0xb7efdc7c00000000),
	207: uint64(0x21dfdb0b00000000),
	208: uint64(0xd4d2d38600000000),
	209: uint64(0x42e2d4f100000000),
	210: uint64(0xf8b3dd6800000000),
	211: uint64(0x6e83da1f00000000),
	212: uint64(0xcd16be8100000000),
	213: uint64(0x5b26b9f600000000),
	214: uint64(0xe177b06f00000000),
	215: uint64(0x7747b71800000000),
	216: uint64(0xe65a088800000000),
	217: uint64(0x706a0fff00000000),
	218: uint64(0xca3b066600000000),
	219: uint64(0x5c0b011100000000),
	220: uint64(0xff9e658f00000000),
	221: uint64(0x69ae62f800000000),
	222: uint64(0xd3ff6b6100000000),
	223: uint64(0x45cf6c1600000000),
	224: uint64(0x78e20aa000000000),
	225: uint64(0xeed20dd700000000),
	226: uint64(0x5483044e00000000),
	227: uint64(0xc2b3033900000000),
	228: uint64(0x612667a700000000),
	229: uint64(0xf71660d000000000),
	230: uint64(0x4d47694900000000),
	231: uint64(0xdb776e3e00000000),
	232: uint64(0x4a6ad1ae00000000),
	233: uint64(0xdc5ad6d900000000),
	234: uint64(0x660bdf4000000000),
	235: uint64(0xf03bd83700000000),
	236: uint64(0x53aebca900000000),
	237: uint64(0xc59ebbde00000000),
	238: uint64(0x7fcfb24700000000),
	239: uint64(0xe9ffb53000000000),
	240: uint64(0x1cf2bdbd00000000),
	241: uint64(0x8ac2baca00000000),
	242: uint64(0x3093b35300000000),
	243: uint64(0xa6a3b42400000000),
	244: uint64(0x0536d0ba00000000),
	245: uint64(0x9306d7cd00000000),
	246: uint64(0x2957de5400000000),
	247: uint64(0xbf67d92300000000),
	248: uint64(0x2e7a66b300000000),
	249: uint64(0xb84a61c400000000),
	250: uint64(0x021b685d00000000),
	251: uint64(0x942b6f2a00000000),
	252: uint64(0x37be0bb400000000),
	253: uint64(0xa18e0cc300000000),
	254: uint64(0x1bdf055a00000000),
	255: uint64(0x8def022d00000000),
}
var _crc_braid_table = [8][256]Tz_crc_t{
	0: {
		1:   uint32(0xaf449247),
		2:   uint32(0x85f822cf),
		3:   uint32(0x2abcb088),
		4:   uint32(0xd08143df),
		5:   uint32(0x7fc5d198),
		6:   uint32(0x55796110),
		7:   uint32(0xfa3df357),
		8:   uint32(0x7a7381ff),
		9:   uint32(0xd53713b8),
		10:  uint32(0xff8ba330),
		11:  uint32(0x50cf3177),
		12:  uint32(0xaaf2c220),
		13:  uint32(0x05b65067),
		14:  uint32(0x2f0ae0ef),
		15:  uint32(0x804e72a8),
		16:  uint32(0xf4e703fe),
		17:  uint32(0x5ba391b9),
		18:  uint32(0x711f2131),
		19:  uint32(0xde5bb376),
		20:  uint32(0x24664021),
		21:  uint32(0x8b22d266),
		22:  uint32(0xa19e62ee),
		23:  uint32(0x0edaf0a9),
		24:  uint32(0x8e948201),
		25:  uint32(0x21d01046),
		26:  uint32(0x0b6ca0ce),
		27:  uint32(0xa4283289),
		28:  uint32(0x5e15c1de),
		29:  uint32(0xf1515399),
		30:  uint32(0xdbede311),
		31:  uint32(0x74a97156),
		32:  uint32(0x32bf01bd),
		33:  uint32(0x9dfb93fa),
		34:  uint32(0xb7472372),
		35:  uint32(0x1803b135),
		36:  uint32(0xe23e4262),
		37:  uint32(0x4d7ad025),
		38:  uint32(0x67c660ad),
		39:  uint32(0xc882f2ea),
		40:  uint32(0x48cc8042),
		41:  uint32(0xe7881205),
		42:  uint32(0xcd34a28d),
		43:  uint32(0x627030ca),
		44:  uint32(0x984dc39d),
		45:  uint32(0x370951da),
		46:  uint32(0x1db5e152),
		47:  uint32(0xb2f17315),
		48:  uint32(0xc6580243),
		49:  uint32(0x691c9004),
		50:  uint32(0x43a0208c),
		51:  uint32(0xece4b2cb),
		52:  uint32(0x16d9419c),
		53:  uint32(0xb99dd3db),
		54:  uint32(0x93216353),
		55:  uint32(0x3c65f114),
		56:  uint32(0xbc2b83bc),
		57:  uint32(0x136f11fb),
		58:  uint32(0x39d3a173),
		59:  uint32(0x96973334),
		60:  uint32(0x6caac063),
		61:  uint32(0xc3ee5224),
		62:  uint32(0xe952e2ac),
		63:  uint32(0x461670eb),
		64:  uint32(0x657e037a),
		65:  uint32(0xca3a913d),
		66:  uint32(0xe08621b5),
		67:  uint32(0x4fc2b3f2),
		68:  uint32(0xb5ff40a5),
		69:  uint32(0x1abbd2e2),
		70:  uint32(0x3007626a),
		71:  uint32(0x9f43f02d),
		72:  uint32(0x1f0d8285),
		73:  uint32(0xb04910c2),
		74:  uint32(0x9af5a04a),
		75:  uint32(0x35b1320d),
		76:  uint32(0xcf8cc15a),
		77:  uint32(0x60c8531d),
		78:  uint32(0x4a74e395),
		79:  uint32(0xe53071d2),
		80:  uint32(0x91990084),
		81:  uint32(0x3edd92c3),
		82:  uint32(0x1461224b),
		83:  uint32(0xbb25b00c),
		84:  uint32(0x4118435b),
		85:  uint32(0xee5cd11c),
		86:  uint32(0xc4e06194),
		87:  uint32(0x6ba4f3d3),
		88:  uint32(0xebea817b),
		89:  uint32(0x44ae133c),
		90:  uint32(0x6e12a3b4),
		91:  uint32(0xc15631f3),
		92:  uint32(0x3b6bc2a4),
		93:  uint32(0x942f50e3),
		94:  uint32(0xbe93e06b),
		95:  uint32(0x11d7722c),
		96:  uint32(0x57c102c7),
		97:  uint32(0xf8859080),
		98:  uint32(0xd2392008),
		99:  uint32(0x7d7db24f),
		100: uint32(0x87404118),
		101: uint32(0x2804d35f),
		102: uint32(0x02b863d7),
		103: uint32(0xadfcf190),
		104: uint32(0x2db28338),
		105: uint32(0x82f6117f),
		106: uint32(0xa84aa1f7),
		107: uint32(0x070e33b0),
		108: uint32(0xfd33c0e7),
		109: uint32(0x527752a0),
		110: uint32(0x78cbe228),
		111: uint32(0xd78f706f),
		112: uint32(0xa3260139),
		113: uint32(0x0c62937e),
		114: uint32(0x26de23f6),
		115: uint32(0x899ab1b1),
		116: uint32(0x73a742e6),
		117: uint32(0xdce3d0a1),
		118: uint32(0xf65f6029),
		119: uint32(0x591bf26e),
		120: uint32(0xd95580c6),
		121: uint32(0x76111281),
		122: uint32(0x5cada209),
		123: uint32(0xf3e9304e),
		124: uint32(0x09d4c319),
		125: uint32(0xa690515e),
		126: uint32(0x8c2ce1d6),
		127: uint32(0x23687391),
		128: uint32(0xcafc06f4),
		129: uint32(0x65b894b3),
		130: uint32(0x4f04243b),
		131: uint32(0xe040b67c),
		132: uint32(0x1a7d452b),
		133: uint32(0xb539d76c),
		134: uint32(0x9f8567e4),
		135: uint32(0x30c1f5a3),
		136: uint32(0xb08f870b),
		137: uint32(0x1fcb154c),
		138: uint32(0x3577a5c4),
		139: uint32(0x9a333783),
		140: uint32(0x600ec4d4),
		141: uint32(0xcf4a5693),
		142: uint32(0xe5f6e61b),
		143: uint32(0x4ab2745c),
		144: uint32(0x3e1b050a),
		145: uint32(0x915f974d),
		146: uint32(0xbbe327c5),
		147: uint32(0x14a7b582),
		148: uint32(0xee9a46d5),
		149: uint32(0x41ded492),
		150: uint32(0x6b62641a),
		151: uint32(0xc426f65d),
		152: uint32(0x446884f5),
		153: uint32(0xeb2c16b2),
		154: uint32(0xc190a63a),
		155: uint32(0x6ed4347d),
		156: uint32(0x94e9c72a),
		157: uint32(0x3bad556d),
		158: uint32(0x1111e5e5),
		159: uint32(0xbe5577a2),
		160: uint32(0xf8430749),
		161: uint32(0x5707950e),
		162: uint32(0x7dbb2586),
		163: uint32(0xd2ffb7c1),
		164: uint32(0x28c24496),
		165: uint32(0x8786d6d1),
		166: uint32(0xad3a6659),
		167: uint32(0x027ef41e),
		168: uint32(0x823086b6),
		169: uint32(0x2d7414f1),
		170: uint32(0x07c8a479),
		171: uint32(0xa88c363e),
		172: uint32(0x52b1c569),
		173: uint32(0xfdf5572e),
		174: uint32(0xd749e7a6),
		175: uint32(0x780d75e1),
		176: uint32(0x0ca404b7),
		177: uint32(0xa3e096f0),
		178: uint32(0x895c2678),
		179: uint32(0x2618b43f),
		180: uint32(0xdc254768),
		181: uint32(0x7361d52f),
		182: uint32(0x59dd65a7),
		183: uint32(0xf699f7e0),
		184: uint32(0x76d78548),
		185: uint32(0xd993170f),
		186: uint32(0xf32fa787),
		187: uint32(0x5c6b35c0),
		188: uint32(0xa656c697),
		189: uint32(0x091254d0),
		190: uint32(0x23aee458),
		191: uint32(0x8cea761f),
		192: uint32(0xaf82058e),
		193: uint32(0x00c697c9),
		194: uint32(0x2a7a2741),
		195: uint32(0x853eb506),
		196: uint32(0x7f034651),
		197: uint32(0xd047d416),
		198: uint32(0xfafb649e),
		199: uint32(0x55bff6d9),
		200: uint32(0xd5f18471),
		201: uint32(0x7ab51636),
		202: uint32(0x5009a6be),
		203: uint32(0xff4d34f9),
		204: uint32(0x0570c7ae),
		205: uint32(0xaa3455e9),
		206: uint32(0x8088e561),
		207: uint32(0x2fcc7726),
		208: uint32(0x5b650670),
		209: uint32(0xf4219437),
		210: uint32(0xde9d24bf),
		211: uint32(0x71d9b6f8),
		212: uint32(0x8be445af),
		213: uint32(0x24a0d7e8),
		214: uint32(0x0e1c6760),
		215: uint32(0xa158f527),
		216: uint32(0x2116878f),
		217: uint32(0x8e5215c8),
		218: uint32(0xa4eea540),
		219: uint32(0x0baa3707),
		220: uint32(0xf197c450),
		221: uint32(0x5ed35617),
		222: uint32(0x746fe69f),
		223: uint32(0xdb2b74d8),
		224: uint32(0x9d3d0433),
		225: uint32(0x32799674),
		226: uint32(0x18c526fc),
		227: uint32(0xb781b4bb),
		228: uint32(0x4dbc47ec),
		229: uint32(0xe2f8d5ab),
		230: uint32(0xc8446523),
		231: uint32(0x6700f764),
		232: uint32(0xe74e85cc),
		233: uint32(0x480a178b),
		234: uint32(0x62b6a703),
		235: uint32(0xcdf23544),
		236: uint32(0x37cfc613),
		237: uint32(0x988b5454),
		238: uint32(0xb237e4dc),
		239: uint32(0x1d73769b),
		240: uint32(0x69da07cd),
		241: uint32(0xc69e958a),
		242: uint32(0xec222502),
		243: uint32(0x4366b745),
		244: uint32(0xb95b4412),
		245: uint32(0x161fd655),
		246: uint32(0x3ca366dd),
		247: uint32(0x93e7f49a),
		248: uint32(0x13a98632),
		249: uint32(0xbced1475),
		250: uint32(0x9651a4fd),
		251: uint32(0x391536ba),
		252: uint32(0xc328c5ed),
		253: uint32(0x6c6c57aa),
		254: uint32(0x46d0e722),
		255: uint32(0xe9947565),
	},
	1: {
		1:   uint32(0x4e890ba9),
		2:   uint32(0x9d121752),
		3:   uint32(0xd39b1cfb),
		4:   uint32(0xe15528e5),
		5:   uint32(0xafdc234c),
		6:   uint32(0x7c473fb7),
		7:   uint32(0x32ce341e),
		8:   uint32(0x19db578b),
		9:   uint32(0x57525c22),
		10:  uint32(0x84c940d9),
		11:  uint32(0xca404b70),
		12:  uint32(0xf88e7f6e),
		13:  uint32(0xb60774c7),
		14:  uint32(0x659c683c),
		15:  uint32(0x2b156395),
		16:  uint32(0x33b6af16),
		17:  uint32(0x7d3fa4bf),
		18:  uint32(0xaea4b844),
		19:  uint32(0xe02db3ed),
		20:  uint32(0xd2e387f3),
		21:  uint32(0x9c6a8c5a),
		22:  uint32(0x4ff190a1),
		23:  uint32(0x01789b08),
		24:  uint32(0x2a6df89d),
		25:  uint32(0x64e4f334),
		26:  uint32(0xb77fefcf),
		27:  uint32(0xf9f6e466),
		28:  uint32(0xcb38d078),
		29:  uint32(0x85b1dbd1),
		30:  uint32(0x562ac72a),
		31:  uint32(0x18a3cc83),
		32:  uint32(0x676d5e2c),
		33:  uint32(0x29e45585),
		34:  uint32(0xfa7f497e),
		35:  uint32(0xb4f642d7),
		36:  uint32(0x863876c9),
		37:  uint32(0xc8b17d60),
		38:  uint32(0x1b2a619b),
		39:  uint32(0x55a36a32),
		40:  uint32(0x7eb609a7),
		41:  uint32(0x303f020e),
		42:  uint32(0xe3a41ef5),
		43:  uint32(0xad2d155c),
		44:  uint32(0x9fe32142),
		45:  uint32(0xd16a2aeb),
		46:  uint32(0x02f13610),
		47:  uint32(0x4c783db9),
		48:  uint32(0x54dbf13a),
		49:  uint32(0x1a52fa93),
		50:  uint32(0xc9c9e668),
		51:  uint32(0x8740edc1),
		52:  uint32(0xb58ed9df),
		53:  uint32(0xfb07d276),
		54:  uint32(0x289cce8d),
		55:  uint32(0x6615c524),
		56:  uint32(0x4d00a6b1),
		57:  uint32(0x0389ad18),
		58:  uint32(0xd012b1e3),
		59:  uint32(0x9e9bba4a),
		60:  uint32(0xac558e54),
		61:  uint32(0xe2dc85fd),
		62:  uint32(0x31479906),
		63:  uint32(0x7fce92af),
		64:  uint32(0xcedabc58),
		65:  uint32(0x8053b7f1),
		66:  uint32(0x53c8ab0a),
		67:  uint32(0x1d41a0a3),
		68:  uint32(0x2f8f94bd),
		69:  uint32(0x61069f14),
		70:  uint32(0xb29d83ef),
		71:  uint32(0xfc148846),
		72:  uint32(0xd701ebd3),
		73:  uint32(0x9988e07a),
		74:  uint32(0x4a13fc81),
		75:  uint32(0x049af728),
		76:  uint32(0x3654c336),
		77:  uint32(0x78ddc89f),
		78:  uint32(0xab46d464),
		79:  uint32(0xe5cfdfcd),
		80:  uint32(0xfd6c134e),
		81:  uint32(0xb3e518e7),
		82:  uint32(0x607e041c),
		83:  uint32(0x2ef70fb5),
		84:  uint32(0x1c393bab),
		85:  uint32(0x52b03002),
		86:  uint32(0x812b2cf9),
		87:  uint32(0xcfa22750),
		88:  uint32(0xe4b744c5),
		89:  uint32(0xaa3e4f6c),
		90:  uint32(0x79a55397),
		91:  uint32(0x372c583e),
		92:  uint32(0x05e26c20),
		93:  uint32(0x4b6b6789),
		94:  uint32(0x98f07b72),
		95:  uint32(0xd67970db),
		96:  uint32(0xa9b7e274),
		97:  uint32(0xe73ee9dd),
		98:  uint32(0x34a5f526),
		99:  uint32(0x7a2cfe8f),
		100: uint32(0x48e2ca91),
		101: uint32(0x066bc138),
		102: uint32(0xd5f0ddc3),
		103: uint32(0x9b79d66a),
		104: uint32(0xb06cb5ff),
		105: uint32(0xfee5be56),
		106: uint32(0x2d7ea2ad),
		107: uint32(0x63f7a904),
		108: uint32(0x51399d1a),
		109: uint32(0x1fb096b3),
		110: uint32(0xcc2b8a48),
		111: uint32(0x82a281e1),
		112: uint32(0x9a014d62),
		113: uint32(0xd48846cb),
		114: uint32(0x07135a30),
		115: uint32(0x499a5199),
		116: uint32(0x7b546587),
		117: uint32(0x35dd6e2e),
		118: uint32(0xe64672d5),
		119: uint32(0xa8cf797c),
		120: uint32(0x83da1ae9),
		121: uint32(0xcd531140),
		122: uint32(0x1ec80dbb),
		123: uint32(0x50410612),
		124: uint32(0x628f320c),
		125: uint32(0x2c0639a5),
		126: uint32(0xff9d255e),
		127: uint32(0xb1142ef7),
		128: uint32(0x46c47ef1),
		129: uint32(0x084d7558),
		130: uint32(0xdbd669a3),
		131: uint32(0x955f620a),
		132: uint32(0xa7915614),
		133: uint32(0xe9185dbd),
		134: uint32(0x3a834146),
		135: uint32(0x740a4aef),
		136: uint32(0x5f1f297a),
		137: uint32(0x119622d3),
		138: uint32(0xc20d3e28),
		139: uint32(0x8c843581),
		140: uint32(0xbe4a019f),
		141: uint32(0xf0c30a36),
		142: uint32(0x235816cd),
		143: uint32(0x6dd11d64),
		144: uint32(0x7572d1e7),
		145: uint32(0x3bfbda4e),
		146: uint32(0xe860c6b5),
		147: uint32(0xa6e9cd1c),
		148: uint32(0x9427f902),
		149: uint32(0xdaaef2ab),
		150: uint32(0x0935ee50),
		151: uint32(0x47bce5f9),
		152: uint32(0x6ca9866c),
		153: uint32(0x22208dc5),
		154: uint32(0xf1bb913e),
		155: uint32(0xbf329a97),
		156: uint32(0x8dfcae89),
		157: uint32(0xc375a520),
		158: uint32(0x10eeb9db),
		159: uint32(0x5e67b272),
		160: uint32(0x21a920dd),
		161: uint32(0x6f202b74),
		162: uint32(0xbcbb378f),
		163: uint32(0xf2323c26),
		164: uint32(0xc0fc0838),
		165: uint32(0x8e750391),
		166: uint32(0x5dee1f6a),
		167: uint32(0x136714c3),
		168: uint32(0x38727756),
		169: uint32(0x76fb7cff),
		170: uint32(0xa5606004),
		171: uint32(0xebe96bad),
		172: uint32(0xd9275fb3),
		173: uint32(0x97ae541a),
		174: uint32(0x443548e1),
		175: uint32(0x0abc4348),
		176: uint32(0x121f8fcb),
		177: uint32(0x5c968462),
		178: uint32(0x8f0d9899),
		179: uint32(0xc1849330),
		180: uint32(0xf34aa72e),
		181: uint32(0xbdc3ac87),
		182: uint32(0x6e58b07c),
		183: uint32(0x20d1bbd5),
		184: uint32(0x0bc4d840),
		185: uint32(0x454dd3e9),
		186: uint32(0x96d6cf12),
		187: uint32(0xd85fc4bb),
		188: uint32(0xea91f0a5),
		189: uint32(0xa418fb0c),
		190: uint32(0x7783e7f7),
		191: uint32(0x390aec5e),
		192: uint32(0x881ec2a9),
		193: uint32(0xc697c900),
		194: uint32(0x150cd5fb),
		195: uint32(0x5b85de52),
		196: uint32(0x694bea4c),
		197: uint32(0x27c2e1e5),
		198: uint32(0xf459fd1e),
		199: uint32(0xbad0f6b7),
		200: uint32(0x91c59522),
		201: uint32(0xdf4c9e8b),
		202: uint32(0x0cd78270),
		203: uint32(0x425e89d9),
		204: uint32(0x7090bdc7),
		205: uint32(0x3e19b66e),
		206: uint32(0xed82aa95),
		207: uint32(0xa30ba13c),
		208: uint32(0xbba86dbf),
		209: uint32(0xf5216616),
		210: uint32(0x26ba7aed),
		211: uint32(0x68337144),
		212: uint32(0x5afd455a),
		213: uint32(0x14744ef3),
		214: uint32(0xc7ef5208),
		215: uint32(0x896659a1),
		216: uint32(0xa2733a34),
		217: uint32(0xecfa319d),
		218: uint32(0x3f612d66),
		219: uint32(0x71e826cf),
		220: uint32(0x432612d1),
		221: uint32(0x0daf1978),
		222: uint32(0xde340583),
		223: uint32(0x90bd0e2a),
		224: uint32(0xef739c85),
		225: uint32(0xa1fa972c),
		226: uint32(0x72618bd7),
		227: uint32(0x3ce8807e),
		228: uint32(0x0e26b460),
		229: uint32(0x40afbfc9),
		230: uint32(0x9334a332),
		231: uint32(0xddbda89b),
		232: uint32(0xf6a8cb0e),
		233: uint32(0xb821c0a7),
		234: uint32(0x6bbadc5c),
		235: uint32(0x2533d7f5),
		236: uint32(0x17fde3eb),
		237: uint32(0x5974e842),
		238: uint32(0x8aeff4b9),
		239: uint32(0xc466ff10),
		240: uint32(0xdcc53393),
		241: uint32(0x924c383a),
		242: uint32(0x41d724c1),
		243: uint32(0x0f5e2f68),
		244: uint32(0x3d901b76),
		245: uint32(0x731910df),
		246: uint32(0xa0820c24),
		247: uint32(0xee0b078d),
		248: uint32(0xc51e6418),
		249: uint32(0x8b976fb1),
		250: uint32(0x580c734a),
		251: uint32(0x168578e3),
		252: uint32(0x244b4cfd),
		253: uint32(0x6ac24754),
		254: uint32(0xb9595baf),
		255: uint32(0xf7d05006),
	},
	2: {
		1:   uint32(0x8d88fde2),
		2:   uint32(0xc060fd85),
		3:   uint32(0x4de80067),
		4:   uint32(0x5bb0fd4b),
		5:   uint32(0xd63800a9),
		6:   uint32(0x9bd000ce),
		7:   uint32(0x1658fd2c),
		8:   uint32(0xb761fa96),
		9:   uint32(0x3ae90774),
		10:  uint32(0x77010713),
		11:  uint32(0xfa89faf1),
		12:  uint32(0xecd107dd),
		13:  uint32(0x6159fa3f),
		14:  uint32(0x2cb1fa58),
		15:  uint32(0xa13907ba),
		16:  uint32(0xb5b2f36d),
		17:  uint32(0x383a0e8f),
		18:  uint32(0x75d20ee8),
		19:  uint32(0xf85af30a),
		20:  uint32(0xee020e26),
		21:  uint32(0x638af3c4),
		22:  uint32(0x2e62f3a3),
		23:  uint32(0xa3ea0e41),
		24:  uint32(0x02d309fb),
		25:  uint32(0x8f5bf419),
		26:  uint32(0xc2b3f47e),
		27:  uint32(0x4f3b099c),
		28:  uint32(0x5963f4b0),
		29:  uint32(0xd4eb0952),
		30:  uint32(0x99030935),
		31:  uint32(0x148bf4d7),
		32:  uint32(0xb014e09b),
		33:  uint32(0x3d9c1d79),
		34:  uint32(0x70741d1e),
		35:  uint32(0xfdfce0fc),
		36:  uint32(0xeba41dd0),
		37:  uint32(0x662ce032),
		38:  uint32(0x2bc4e055),
		39:  uint32(0xa64c1db7),
		40:  uint32(0x07751a0d),
		41:  uint32(0x8afde7ef),
		42:  uint32(0xc715e788),
		43:  uint32(0x4a9d1a6a),
		44:  uint32(0x5cc5e746),
		45:  uint32(0xd14d1aa4),
		46:  uint32(0x9ca51ac3),
		47:  uint32(0x112de721),
		48:  uint32(0x05a613f6),
		49:  uint32(0x882eee14),
		50:  uint32(0xc5c6ee73),
		51:  uint32(0x484e1391),
		52:  uint32(0x5e16eebd),
		53:  uint32(0xd39e135f),
		54:  uint32(0x9e761338),
		55:  uint32(0x13feeeda),
		56:  uint32(0xb2c7e960),
		57:  uint32(0x3f4f1482),
		58:  uint32(0x72a714e5),
		59:  uint32(0xff2fe907),
		60:  uint32(0xe977142b),
		61:  uint32(0x64ffe9c9),
		62:  uint32(0x2917e9ae),
		63:  uint32(0xa49f144c),
		64:  uint32(0xbb58c777),
		65:  uint32(0x36d03a95),
		66:  uint32(0x7b383af2),
		67:  uint32(0xf6b0c710),
		68:  uint32(0xe0e83a3c),
		69:  uint32(0x6d60c7de),
		70:  uint32(0x2088c7b9),
		71:  uint32(0xad003a5b),
		72:  uint32(0x0c393de1),
		73:  uint32(0x81b1c003),
		74:  uint32(0xcc59c064),
		75:  uint32(0x41d13d86),
		76:  uint32(0x5789c0aa),
		77:  uint32(0xda013d48),
		78:  uint32(0x97e93d2f),
		79:  uint32(0x1a61c0cd),
		80:  uint32(0x0eea341a),
		81:  uint32(0x8362c9f8),
		82:  uint32(0xce8ac99f),
		83:  uint32(0x4302347d),
		84:  uint32(0x555ac951),
		85:  uint32(0xd8d234b3),
		86:  uint32(0x953a34d4),
		87:  uint32(0x18b2c936),
		88:  uint32(0xb98bce8c),
		89:  uint32(0x3403336e),
		90:  uint32(0x79eb3309),
		91:  uint32(0xf463ceeb),
		92:  uint32(0xe23b33c7),
		93:  uint32(0x6fb3ce25),
		94:  uint32(0x225bce42),
		95:  uint32(0xafd333a0),
		96:  uint32(0x0b4c27ec),
		97:  uint32(0x86c4da0e),
		98:  uint32(0xcb2cda69),
		99:  uint32(0x46a4278b),
		100: uint32(0x50fcdaa7),
		101: uint32(0xdd742745),
		102: uint32(0x909c2722),
		103: uint32(0x1d14dac0),
		104: uint32(0xbc2ddd7a),
		105: uint32(0x31a52098),
		106: uint32(0x7c4d20ff),
		107: uint32(0xf1c5dd1d),
		108: uint32(0xe79d2031),
		109: uint32(0x6a15ddd3),
		110: uint32(0x27fdddb4),
		111: uint32(0xaa752056),
		112: uint32(0xbefed481),
		113: uint32(0x33762963),
		114: uint32(0x7e9e2904),
		115: uint32(0xf316d4e6),
		116: uint32(0xe54e29ca),
		117: uint32(0x68c6d428),
		118: uint32(0x252ed44f),
		119: uint32(0xa8a629ad),
		120: uint32(0x099f2e17),
		121: uint32(0x8417d3f5),
		122: uint32(0xc9ffd392),
		123: uint32(0x44772e70),
		124: uint32(0x522fd35c),
		125: uint32(0xdfa72ebe),
		126: uint32(0x924f2ed9),
		127: uint32(0x1fc7d33b),
		128: uint32(0xadc088af),
		129: uint32(0x2048754d),
		130: uint32(0x6da0752a),
		131: uint32(0xe02888c8),
		132: uint32(0xf67075e4),
		133: uint32(0x7bf88806),
		134: uint32(0x36108861),
		135: uint32(0xbb987583),
		136: uint32(0x1aa17239),
		137: uint32(0x97298fdb),
		138: uint32(0xdac18fbc),
		139: uint32(0x5749725e),
		140: uint32(0x41118f72),
		141: uint32(0xcc997290),
		142: uint32(0x817172f7),
		143: uint32(0x0cf98f15),
		144: uint32(0x18727bc2),
		145: uint32(0x95fa8620),
		146: uint32(0xd8128647),
		147: uint32(0x559a7ba5),
		148: uint32(0x43c28689),
		149: uint32(0xce4a7b6b),
		150: uint32(0x83a27b0c),
		151: uint32(0x0e2a86ee),
		152: uint32(0xaf138154),
		153: uint32(0x229b7cb6),
		154: uint32(0x6f737cd1),
		155: uint32(0xe2fb8133),
		156: uint32(0xf4a37c1f),
		157: uint32(0x792b81fd),
		158: uint32(0x34c3819a),
		159: uint32(0xb94b7c78),
		160: uint32(0x1dd46834),
		161: uint32(0x905c95d6),
		162: uint32(0xddb495b1),
		163: uint32(0x503c6853),
		164: uint32(0x4664957f),
		165: uint32(0xcbec689d),
		166: uint32(0x860468fa),
		167: uint32(0x0b8c9518),
		168: uint32(0xaab592a2),
		169: uint32(0x273d6f40),
		170: uint32(0x6ad56f27),
		171: uint32(0xe75d92c5),
		172: uint32(0xf1056fe9),
		173: uint32(0x7c8d920b),
		174: uint32(0x3165926c),
		175: uint32(0xbced6f8e),
		176: uint32(0xa8669b59),
		177: uint32(0x25ee66bb),
		178: uint32(0x680666dc),
		179: uint32(0xe58e9b3e),
		180: uint32(0xf3d66612),
		181: uint32(0x7e5e9bf0),
		182: uint32(0x33b69b97),
		183: uint32(0xbe3e6675),
		184: uint32(0x1f0761cf),
		185: uint32(0x928f9c2d),
		186: uint32(0xdf679c4a),
		187: uint32(0x52ef61a8),
		188: uint32(0x44b79c84),
		189: uint32(0xc93f6166),
		190: uint32(0x84d76101),
		191: uint32(0x095f9ce3),
		192: uint32(0x16984fd8),
		193: uint32(0x9b10b23a),
		194: uint32(0xd6f8b25d),
		195: uint32(0x5b704fbf),
		196: uint32(0x4d28b293),
		197: uint32(0xc0a04f71),
		198: uint32(0x8d484f16),
		199: uint32(0x00c0b2f4),
		200: uint32(0xa1f9b54e),
		201: uint32(0x2c7148ac),
		202: uint32(0x619948cb),
		203: uint32(0xec11b529),
		204: uint32(0xfa494805),
		205: uint32(0x77c1b5e7),
		206: uint32(0x3a29b580),
		207: uint32(0xb7a14862),
		208: uint32(0xa32abcb5),
		209: uint32(0x2ea24157),
		210: uint32(0x634a4130),
		211: uint32(0xeec2bcd2),
		212: uint32(0xf89a41fe),
		213: uint32(0x7512bc1c),
		214: uint32(0x38fabc7b),
		215: uint32(0xb5724199),
		216: uint32(0x144b4623),
		217: uint32(0x99c3bbc1),
		218: uint32(0xd42bbba6),
		219: uint32(0x59a34644),
		220: uint32(0x4ffbbb68),
		221: uint32(0xc273468a),
		222: uint32(0x8f9b46ed),
		223: uint32(0x0213bb0f),
		224: uint32(0xa68caf43),
		225: uint32(0x2b0452a1),
		226: uint32(0x66ec52c6),
		227: uint32(0xeb64af24),
		228: uint32(0xfd3c5208),
		229: uint32(0x70b4afea),
		230: uint32(0x3d5caf8d),
		231: uint32(0xb0d4526f),
		232: uint32(0x11ed55d5),
		233: uint32(0x9c65a837),
		234: uint32(0xd18da850),
		235: uint32(0x5c0555b2),
		236: uint32(0x4a5da89e),
		237: uint32(0xc7d5557c),
		238: uint32(0x8a3d551b),
		239: uint32(0x07b5a8f9),
		240: uint32(0x133e5c2e),
		241: uint32(0x9eb6a1cc),
		242: uint32(0xd35ea1ab),
		243: uint32(0x5ed65c49),
		244: uint32(0x488ea165),
		245: uint32(0xc5065c87),
		246: uint32(0x88ee5ce0),
		247: uint32(0x0566a102),
		248: uint32(0xa45fa6b8),
		249: uint32(0x29d75b5a),
		250: uint32(0x643f5b3d),
		251: uint32(0xe9b7a6df),
		252: uint32(0xffef5bf3),
		253: uint32(0x7267a611),
		254: uint32(0x3f8fa676),
		255: uint32(0xb2075b94),
	},
	3: {
		1:   uint32(0x80f0171f),
		2:   uint32(0xda91287f),
		3:   uint32(0x5a613f60),
		4:   uint32(0x6e5356bf),
		5:   uint32(0xeea341a0),
		6:   uint32(0xb4c27ec0),
		7:   uint32(0x343269df),
		8:   uint32(0xdca6ad7e),
		9:   uint32(0x5c56ba61),
		10:  uint32(0x06378501),
		11:  uint32(0x86c7921e),
		12:  uint32(0xb2f5fbc1),
		13:  uint32(0x3205ecde),
		14:  uint32(0x6864d3be),
		15:  uint32(0xe894c4a1),
		16:  uint32(0x623c5cbd),
		17:  uint32(0xe2cc4ba2),
		18:  uint32(0xb8ad74c2),
		19:  uint32(0x385d63dd),
		20:  uint32(0x0c6f0a02),
		21:  uint32(0x8c9f1d1d),
		22:  uint32(0xd6fe227d),
		23:  uint32(0x560e3562),
		24:  uint32(0xbe9af1c3),
		25:  uint32(0x3e6ae6dc),
		26:  uint32(0x640bd9bc),
		27:  uint32(0xe4fbcea3),
		28:  uint32(0xd0c9a77c),
		29:  uint32(0x5039b063),
		30:  uint32(0x0a588f03),
		31:  uint32(0x8aa8981c),
		32:  uint32(0xc478b97a),
		33:  uint32(0x4488ae65),
		34:  uint32(0x1ee99105),
		35:  uint32(0x9e19861a),
		36:  uint32(0xaa2befc5),
		37:  uint32(0x2adbf8da),
		38:  uint32(0x70bac7ba),
		39:  uint32(0xf04ad0a5),
		40:  uint32(0x18de1404),
		41:  uint32(0x982e031b),
		42:  uint32(0xc24f3c7b),
		43:  uint32(0x42bf2b64),
		44:  uint32(0x768d42bb),
		45:  uint32(0xf67d55a4),
		46:  uint32(0xac1c6ac4),
		47:  uint32(0x2cec7ddb),
		48:  uint32(0xa644e5c7),
		49:  uint32(0x26b4f2d8),
		50:  uint32(0x7cd5cdb8),
		51:  uint32(0xfc25daa7),
		52:  uint32(0xc817b378),
		53:  uint32(0x48e7a467),
		54:  uint32(0x12869b07),
		55:  uint32(0x92768c18),
		56:  uint32(0x7ae248b9),
		57:  uint32(0xfa125fa6),
		58:  uint32(0xa07360c6),
		59:  uint32(0x208377d9),
		60:  uint32(0x14b11e06),
		61:  uint32(0x94410919),
		62:  uint32(0xce203679),
		63:  uint32(0x4ed02166),
		64:  uint32(0x538074b5),
		65:  uint32(0xd37063aa),
		66:  uint32(0x89115cca),
		67:  uint32(0x09e14bd5),
		68:  uint32(0x3dd3220a),
		69:  uint32(0xbd233515),
		70:  uint32(0xe7420a75),
		71:  uint32(0x67b21d6a),
		72:  uint32(0x8f26d9cb),
		73:  uint32(0x0fd6ced4),
		74:  uint32(0x55b7f1b4),
		75:  uint32(0xd547e6ab),
		76:  uint32(0xe1758f74),
		77:  uint32(0x6185986b),
		78:  uint32(0x3be4a70b),
		79:  uint32(0xbb14b014),
		80:  uint32(0x31bc2808),
		81:  uint32(0xb14c3f17),
		82:  uint32(0xeb2d0077),
		83:  uint32(0x6bdd1768),
		84:  uint32(0x5fef7eb7),
		85:  uint32(0xdf1f69a8),
		86:  uint32(0x857e56c8),
		87:  uint32(0x058e41d7),
		88:  uint32(0xed1a8576),
		89:  uint32(0x6dea9269),
		90:  uint32(0x378bad09),
		91:  uint32(0xb77bba16),
		92:  uint32(0x8349d3c9),
		93:  uint32(0x03b9c4d6),
		94:  uint32(0x59d8fbb6),
		95:  uint32(0xd928eca9),
		96:  uint32(0x97f8cdcf),
		97:  uint32(0x1708dad0),
		98:  uint32(0x4d69e5b0),
		99:  uint32(0xcd99f2af),
		100: uint32(0xf9ab9b70),
		101: uint32(0x795b8c6f),
		102: uint32(0x233ab30f),
		103: uint32(0xa3caa410),
		104: uint32(0x4b5e60b1),
		105: uint32(0xcbae77ae),
		106: uint32(0x91cf48ce),
		107: uint32(0x113f5fd1),
		108: uint32(0x250d360e),
		109: uint32(0xa5fd2111),
		110: uint32(0xff9c1e71),
		111: uint32(0x7f6c096e),
		112: uint32(0xf5c49172),
		113: uint32(0x7534866d),
		114: uint32(0x2f55b90d),
		115: uint32(0xafa5ae12),
		116: uint32(0x9b97c7cd),
		117: uint32(0x1b67d0d2),
		118: uint32(0x4106efb2),
		119: uint32(0xc1f6f8ad),
		120: uint32(0x29623c0c),
		121: uint32(0xa9922b13),
		122: uint32(0xf3f31473),
		123: uint32(0x7303036c),
		124: uint32(0x47316ab3),
		125: uint32(0xc7c17dac),
		126: uint32(0x9da042cc),
		127: uint32(0x1d5055d3),
		128: uint32(0xa700e96a),
		129: uint32(0x27f0fe75),
		130: uint32(0x7d91c115),
		131: uint32(0xfd61d60a),
		132: uint32(0xc953bfd5),
		133: uint32(0x49a3a8ca),
		134: uint32(0x13c297aa),
		135: uint32(0x933280b5),
		136: uint32(0x7ba64414),
		137: uint32(0xfb56530b),
		138: uint32(0xa1376c6b),
		139: uint32(0x21c77b74),
		140: uint32(0x15f512ab),
		141: uint32(0x950505b4),
		142: uint32(0xcf643ad4),
		143: uint32(0x4f942dcb),
		144: uint32(0xc53cb5d7),
		145: uint32(0x45cca2c8),
		146: uint32(0x1fad9da8),
		147: uint32(0x9f5d8ab7),
		148: uint32(0xab6fe368),
		149: uint32(0x2b9ff477),
		150: uint32(0x71fecb17),
		151: uint32(0xf10edc08),
		152: uint32(0x199a18a9),
		153: uint32(0x996a0fb6),
		154: uint32(0xc30b30d6),
		155: uint32(0x43fb27c9),
		156: uint32(0x77c94e16),
		157: uint32(0xf7395909),
		158: uint32(0xad586669),
		159: uint32(0x2da87176),
		160: uint32(0x63785010),
		161: uint32(0xe388470f),
		162: uint32(0xb9e9786f),
		163: uint32(0x39196f70),
		164: uint32(0x0d2b06af),
		165: uint32(0x8ddb11b0),
		166: uint32(0xd7ba2ed0),
		167: uint32(0x574a39cf),
		168: uint32(0xbfdefd6e),
		169: uint32(0x3f2eea71),
		170: uint32(0x654fd511),
		171: uint32(0xe5bfc20e),
		172: uint32(0xd18dabd1),
		173: uint32(0x517dbcce),
		174: uint32(0x0b1c83ae),
		175: uint32(0x8bec94b1),
		176: uint32(0x01440cad),
		177: uint32(0x81b41bb2),
		178: uint32(0xdbd524d2),
		179: uint32(0x5b2533cd),
		180: uint32(0x6f175a12),
		181: uint32(0xefe74d0d),
		182: uint32(0xb586726d),
		183: uint32(0x35766572),
		184: uint32(0xdde2a1d3),
		185: uint32(0x5d12b6cc),
		186: uint32(0x077389ac),
		187: uint32(0x87839eb3),
		188: uint32(0xb3b1f76c),
		189: uint32(0x3341e073),
		190: uint32(0x6920df13),
		191: uint32(0xe9d0c80c),
		192: uint32(0xf4809ddf),
		193: uint32(0x74708ac0),
		194: uint32(0x2e11b5a0),
		195: uint32(0xaee1a2bf),
		196: uint32(0x9ad3cb60),
		197: uint32(0x1a23dc7f),
		198: uint32(0x4042e31f),
		199: uint32(0xc0b2f400),
		200: uint32(0x282630a1),
		201: uint32(0xa8d627be),
		202: uint32(0xf2b718de),
		203: uint32(0x72470fc1),
		204: uint32(0x4675661e),
		205: uint32(0xc6857101),
		206: uint32(0x9ce44e61),
		207: uint32(0x1c14597e),
		208: uint32(0x96bcc162),
		209: uint32(0x164cd67d),
		210: uint32(0x4c2de91d),
		211: uint32(0xccddfe02),
		212: uint32(0xf8ef97dd),
		213: uint32(0x781f80c2),
		214: uint32(0x227ebfa2),
		215: uint32(0xa28ea8bd),
		216: uint32(0x4a1a6c1c),
		217: uint32(0xcaea7b03),
		218: uint32(0x908b4463),
		219: uint32(0x107b537c),
		220: uint32(0x24493aa3),
		221: uint32(0xa4b92dbc),
		222: uint32(0xfed812dc),
		223: uint32(0x7e2805c3),
		224: uint32(0x30f824a5),
		225: uint32(0xb00833ba),
		226: uint32(0xea690cda),
		227: uint32(0x6a991bc5),
		228: uint32(0x5eab721a),
		229: uint32(0xde5b6505),
		230: uint32(0x843a5a65),
		231: uint32(0x04ca4d7a),
		232: uint32(0xec5e89db),
		233: uint32(0x6cae9ec4),
		234: uint32(0x36cfa1a4),
		235: uint32(0xb63fb6bb),
		236: uint32(0x820ddf64),
		237: uint32(0x02fdc87b),
		238: uint32(0x589cf71b),
		239: uint32(0xd86ce004),
		240: uint32(0x52c47818),
		241: uint32(0xd2346f07),
		242: uint32(0x88555067),
		243: uint32(0x08a54778),
		244: uint32(0x3c972ea7),
		245: uint32(0xbc6739b8),
		246: uint32(0xe60606d8),
		247: uint32(0x66f611c7),
		248: uint32(0x8e62d566),
		249: uint32(0x0e92c279),
		250: uint32(0x54f3fd19),
		251: uint32(0xd403ea06),
		252: uint32(0xe03183d9),
		253: uint32(0x60c194c6),
		254: uint32(0x3aa0aba6),
		255: uint32(0xba50bcb9),
	},
	4: {
		1:   uint32(0x9570d495),
		2:   uint32(0xf190af6b),
		3:   uint32(0x64e07bfe),
		4:   uint32(0x38505897),
		5:   uint32(0xad208c02),
		6:   uint32(0xc9c0f7fc),
		7:   uint32(0x5cb02369),
		8:   uint32(0x70a0b12e),
		9:   uint32(0xe5d065bb),
		10:  uint32(0x81301e45),
		11:  uint32(0x1440cad0),
		12:  uint32(0x48f0e9b9),
		13:  uint32(0xdd803d2c),
		14:  uint32(0xb96046d2),
		15:  uint32(0x2c109247),
		16:  uint32(0xe141625c),
		17:  uint32(0x7431b6c9),
		18:  uint32(0x10d1cd37),
		19:  uint32(0x85a119a2),
		20:  uint32(0xd9113acb),
		21:  uint32(0x4c61ee5e),
		22:  uint32(0x288195a0),
		23:  uint32(0xbdf14135),
		24:  uint32(0x91e1d372),
		25:  uint32(0x049107e7),
		26:  uint32(0x60717c19),
		27:  uint32(0xf501a88c),
		28:  uint32(0xa9b18be5),
		29:  uint32(0x3cc15f70),
		30:  uint32(0x5821248e),
		31:  uint32(0xcd51f01b),
		32:  uint32(0x19f3c2f9),
		33:  uint32(0x8c83166c),
		34:  uint32(0xe8636d92),
		35:  uint32(0x7d13b907),
		36:  uint32(0x21a39a6e),
		37:  uint32(0xb4d34efb),
		38:  uint32(0xd0333505),
		39:  uint32(0x4543e190),
		40:  uint32(0x695373d7),
		41:  uint32(0xfc23a742),
		42:  uint32(0x98c3dcbc),
		43:  uint32(0x0db30829),
		44:  uint32(0x51032b40),
		45:  uint32(0xc473ffd5),
		46:  uint32(0xa093842b),
		47:  uint32(0x35e350be),
		48:  uint32(0xf8b2a0a5),
		49:  uint32(0x6dc27430),
		50:  uint32(0x09220fce),
		51:  uint32(0x9c52db5b),
		52:  uint32(0xc0e2f832),
		53:  uint32(0x55922ca7),
		54:  uint32(0x31725759),
		55:  uint32(0xa40283cc),
		56:  uint32(0x8812118b),
		57:  uint32(0x1d62c51e),
		58:  uint32(0x7982bee0),
		59:  uint32(0xecf26a75),
		60:  uint32(0xb042491c),
		61:  uint32(0x25329d89),
		62:  uint32(0x41d2e677),
		63:  uint32(0xd4a232e2),
		64:  uint32(0x33e785f2),
		65:  uint32(0xa6975167),
		66:  uint32(0xc2772a99),
		67:  uint32(0x5707fe0c),
		68:  uint32(0x0bb7dd65),
		69:  uint32(0x9ec709f0),
		70:  uint32(0xfa27720e),
		71:  uint32(0x6f57a69b),
		72:  uint32(0x434734dc),
		73:  uint32(0xd637e049),
		74:  uint32(0xb2d79bb7),
		75:  uint32(0x27a74f22),
		76:  uint32(0x7b176c4b),
		77:  uint32(0xee67b8de),
		78:  uint32(0x8a87c320),
		79:  uint32(0x1ff717b5),
		80:  uint32(0xd2a6e7ae),
		81:  uint32(0x47d6333b),
		82:  uint32(0x233648c5),
		83:  uint32(0xb6469c50),
		84:  uint32(0xeaf6bf39),
		85:  uint32(0x7f866bac),
		86:  uint32(0x1b661052),
		87:  uint32(0x8e16c4c7),
		88:  uint32(0xa2065680),
		89:  uint32(0x37768215),
		90:  uint32(0x5396f9eb),
		91:  uint32(0xc6e62d7e),
		92:  uint32(0x9a560e17),
		93:  uint32(0x0f26da82),
		94:  uint32(0x6bc6a17c),
		95:  uint32(0xfeb675e9),
		96:  uint32(0x2a14470b),
		97:  uint32(0xbf64939e),
		98:  uint32(0xdb84e860),
		99:  uint32(0x4ef43cf5),
		100: uint32(0x12441f9c),
		101: uint32(0x8734cb09),
		102: uint32(0xe3d4b0f7),
		103: uint32(0x76a46462),
		104: uint32(0x5ab4f625),
		105: uint32(0xcfc422b0),
		106: uint32(0xab24594e),
		107: uint32(0x3e548ddb),
		108: uint32(0x62e4aeb2),
		109: uint32(0xf7947a27),
		110: uint32(0x937401d9),
		111: uint32(0x0604d54c),
		112: uint32(0xcb552557),
		113: uint32(0x5e25f1c2),
		114: uint32(0x3ac58a3c),
		115: uint32(0xafb55ea9),
		116: uint32(0xf3057dc0),
		117: uint32(0x6675a955),
		118: uint32(0x0295d2ab),
		119: uint32(0x97e5063e),
		120: uint32(0xbbf59479),
		121: uint32(0x2e8540ec),
		122: uint32(0x4a653b12),
		123: uint32(0xdf15ef87),
		124: uint32(0x83a5ccee),
		125: uint32(0x16d5187b),
		126: uint32(0x72356385),
		127: uint32(0xe745b710),
		128: uint32(0x67cf0be4),
		129: uint32(0xf2bfdf71),
		130: uint32(0x965fa48f),
		131: uint32(0x032f701a),
		132: uint32(0x5f9f5373),
		133: uint32(0xcaef87e6),
		134: uint32(0xae0ffc18),
		135: uint32(0x3b7f288d),
		136: uint32(0x176fbaca),
		137: uint32(0x821f6e5f),
		138: uint32(0xe6ff15a1),
		139: uint32(0x738fc134),
		140: uint32(0x2f3fe25d),
		141: uint32(0xba4f36c8),
		142: uint32(0xdeaf4d36),
		143: uint32(0x4bdf99a3),
		144: uint32(0x868e69b8),
		145: uint32(0x13febd2d),
		146: uint32(0x771ec6d3),
		147: uint32(0xe26e1246),
		148: uint32(0xbede312f),
		149: uint32(0x2baee5ba),
		150: uint32(0x4f4e9e44),
		151: uint32(0xda3e4ad1),
		152: uint32(0xf62ed896),
		153: uint32(0x635e0c03),
		154: uint32(0x07be77fd),
		155: uint32(0x92cea368),
		156: uint32(0xce7e8001),
		157: uint32(0x5b0e5494),
		158: uint32(0x3fee2f6a),
		159: uint32(0xaa9efbff),
		160: uint32(0x7e3cc91d),
		161: uint32(0xeb4c1d88),
		162: uint32(0x8fac6676),
		163: uint32(0x1adcb2e3),
		164: uint32(0x466c918a),
		165: uint32(0xd31c451f),
		166: uint32(0xb7fc3ee1),
		167: uint32(0x228cea74),
		168: uint32(0x0e9c7833),
		169: uint32(0x9becaca6),
		170: uint32(0xff0cd758),
		171: uint32(0x6a7c03cd),
		172: uint32(0x36cc20a4),
		173: uint32(0xa3bcf431),
		174: uint32(0xc75c8fcf),
		175: uint32(0x522c5b5a),
		176: uint32(0x9f7dab41),
		177: uint32(0x0a0d7fd4),
		178: uint32(0x6eed042a),
		179: uint32(0xfb9dd0bf),
		180: uint32(0xa72df3d6),
		181: uint32(0x325d2743),
		182: uint32(0x56bd5cbd),
		183: uint32(0xc3cd8828),
		184: uint32(0xefdd1a6f),
		185: uint32(0x7aadcefa),
		186: uint32(0x1e4db504),
		187: uint32(0x8b3d6191),
		188: uint32(0xd78d42f8),
		189: uint32(0x42fd966d),
		190: uint32(0x261ded93),
		191: uint32(0xb36d3906),
		192: uint32(0x54288e16),
		193: uint32(0xc1585a83),
		194: uint32(0xa5b8217d),
		195: uint32(0x30c8f5e8),
		196: uint32(0x6c78d681),
		197: uint32(0xf9080214),
		198: uint32(0x9de879ea),
		199: uint32(0x0898ad7f),
		200: uint32(0x24883f38),
		201: uint32(0xb1f8ebad),
		202: uint32(0xd5189053),
		203: uint32(0x406844c6),
		204: uint32(0x1cd867af),
		205: uint32(0x89a8b33a),
		206: uint32(0xed48c8c4),
		207: uint32(0x78381c51),
		208: uint32(0xb569ec4a),
		209: uint32(0x201938df),
		210: uint32(0x44f94321),
		211: uint32(0xd18997b4),
		212: uint32(0x8d39b4dd),
		213: uint32(0x18496048),
		214: uint32(0x7ca91bb6),
		215: uint32(0xe9d9cf23),
		216: uint32(0xc5c95d64),
		217: uint32(0x50b989f1),
		218: uint32(0x3459f20f),
		219: uint32(0xa129269a),
		220: uint32(0xfd9905f3),
		221: uint32(0x68e9d166),
		222: uint32(0x0c09aa98),
		223: uint32(0x99797e0d),
		224: uint32(0x4ddb4cef),
		225: uint32(0xd8ab987a),
		226: uint32(0xbc4be384),
		227: uint32(0x293b3711),
		228: uint32(0x758b1478),
		229: uint32(0xe0fbc0ed),
		230: uint32(0x841bbb13),
		231: uint32(0x116b6f86),
		232: uint32(0x3d7bfdc1),
		233: uint32(0xa80b2954),
		234: uint32(0xcceb52aa),
		235: uint32(0x599b863f),
		236: uint32(0x052ba556),
		237: uint32(0x905b71c3),
		238: uint32(0xf4bb0a3d),
		239: uint32(0x61cbdea8),
		240: uint32(0xac9a2eb3),
		241: uint32(0x39eafa26),
		242: uint32(0x5d0a81d8),
		243: uint32(0xc87a554d),
		244: uint32(0x94ca7624),
		245: uint32(0x01baa2b1),
		246: uint32(0x655ad94f),
		247: uint32(0xf02a0dda),
		248: uint32(0xdc3a9f9d),
		249: uint32(0x494a4b08),
		250: uint32(0x2daa30f6),
		251: uint32(0xb8dae463),
		252: uint32(0xe46ac70a),
		253: uint32(0x711a139f),
		254: uint32(0x15fa6861),
		255: uint32(0x808abcf4),
	},
	5: {
		1:   uint32(0xcf9e17c8),
		2:   uint32(0x444d29d1),
		3:   uint32(0x8bd33e19),
		4:   uint32(0x889a53a2),
		5:   uint32(0x4704446a),
		6:   uint32(0xccd77a73),
		7:   uint32(0x03496dbb),
		8:   uint32(0xca45a105),
		9:   uint32(0x05dbb6cd),
		10:  uint32(0x8e0888d4),
		11:  uint32(0x41969f1c),
		12:  uint32(0x42dff2a7),
		13:  uint32(0x8d41e56f),
		14:  uint32(0x0692db76),
		15:  uint32(0xc90cccbe),
		16:  uint32(0x4ffa444b),
		17:  uint32(0x80645383),
		18:  uint32(0x0bb76d9a),
		19:  uint32(0xc4297a52),
		20:  uint32(0xc76017e9),
		21:  uint32(0x08fe0021),
		22:  uint32(0x832d3e38),
		23:  uint32(0x4cb329f0),
		24:  uint32(0x85bfe54e),
		25:  uint32(0x4a21f286),
		26:  uint32(0xc1f2cc9f),
		27:  uint32(0x0e6cdb57),
		28:  uint32(0x0d25b6ec),
		29:  uint32(0xc2bba124),
		30:  uint32(0x49689f3d),
		31:  uint32(0x86f688f5),
		32:  uint32(0x9ff48896),
		33:  uint32(0x506a9f5e),
		34:  uint32(0xdbb9a147),
		35:  uint32(0x1427b68f),
		36:  uint32(0x176edb34),
		37:  uint32(0xd8f0ccfc),
		38:  uint32(0x5323f2e5),
		39:  uint32(0x9cbde52d),
		40:  uint32(0x55b12993),
		41:  uint32(0x9a2f3e5b),
		42:  uint32(0x11fc0042),
		43:  uint32(0xde62178a),
		44:  uint32(0xdd2b7a31),
		45:  uint32(0x12b56df9),
		46:  uint32(0x996653e0),
		47:  uint32(0x56f84428),
		48:  uint32(0xd00eccdd),
		49:  uint32(0x1f90db15),
		50:  uint32(0x9443e50c),
		51:  uint32(0x5bddf2c4),
		52:  uint32(0x58949f7f),
		53:  uint32(0x970a88b7),
		54:  uint32(0x1cd9b6ae),
		55:  uint32(0xd347a166),
		56:  uint32(0x1a4b6dd8),
		57:  uint32(0xd5d57a10),
		58:  uint32(0x5e064409),
		59:  uint32(0x919853c1),
		60:  uint32(0x92d13e7a),
		61:  uint32(0x5d4f29b2),
		62:  uint32(0xd69c17ab),
		63:  uint32(0x19020063),
		64:  uint32(0xe498176d),
		65:  uint32(0x2b0600a5),
		66:  uint32(0xa0d53ebc),
		67:  uint32(0x6f4b2974),
		68:  uint32(0x6c0244cf),
		69:  uint32(0xa39c5307),
		70:  uint32(0x284f6d1e),
		71:  uint32(0xe7d17ad6),
		72:  uint32(0x2eddb668),
		73:  uint32(0xe143a1a0),
		74:  uint32(0x6a909fb9),
		75:  uint32(0xa50e8871),
		76:  uint32(0xa647e5ca),
		77:  uint32(0x69d9f202),
		78:  uint32(0xe20acc1b),
		79:  uint32(0x2d94dbd3),
		80:  uint32(0xab625326),
		81:  uint32(0x64fc44ee),
		82:  uint32(0xef2f7af7),
		83:  uint32(0x20b16d3f),
		84:  uint32(0x23f80084),
		85:  uint32(0xec66174c),
		86:  uint32(0x67b52955),
		87:  uint32(0xa82b3e9d),
		88:  uint32(0x6127f223),
		89:  uint32(0xaeb9e5eb),
		90:  uint32(0x256adbf2),
		91:  uint32(0xeaf4cc3a),
		92:  uint32(0xe9bda181),
		93:  uint32(0x2623b649),
		94:  uint32(0xadf08850),
		95:  uint32(0x626e9f98),
		96:  uint32(0x7b6c9ffb),
		97:  uint32(0xb4f28833),
		98:  uint32(0x3f21b62a),
		99:  uint32(0xf0bfa1e2),
		100: uint32(0xf3f6cc59),
		101: uint32(0x3c68db91),
		102: uint32(0xb7bbe588),
		103: uint32(0x7825f240),
		104: uint32(0xb1293efe),
		105: uint32(0x7eb72936),
		106: uint32(0xf564172f),
		107: uint32(0x3afa00e7),
		108: uint32(0x39b36d5c),
		109: uint32(0xf62d7a94),
		110: uint32(0x7dfe448d),
		111: uint32(0xb2605345),
		112: uint32(0x3496dbb0),
		113: uint32(0xfb08cc78),
		114: uint32(0x70dbf261),
		115: uint32(0xbf45e5a9),
		116: uint32(0xbc0c8812),
		117: uint32(0x73929fda),
		118: uint32(0xf841a1c3),
		119: uint32(0x37dfb60b),
		120: uint32(0xfed37ab5),
		121: uint32(0x314d6d7d),
		122: uint32(0xba9e5364),
		123: uint32(0x750044ac),
		124: uint32(0x76492917),
		125: uint32(0xb9d73edf),
		126: uint32(0x320400c6),
		127: uint32(0xfd9a170e),
		128: uint32(0x1241289b),
		129: uint32(0xdddf3f53),
		130: uint32(0x560c014a),
		131: uint32(0x99921682),
		132: uint32(0x9adb7b39),
		133: uint32(0x55456cf1),
		134: uint32(0xde9652e8),
		135: uint32(0x11084520),
		136: uint32(0xd804899e),
		137: uint32(0x179a9e56),
		138: uint32(0x9c49a04f),
		139: uint32(0x53d7b787),
		140: uint32(0x509eda3c),
		141: uint32(0x9f00cdf4),
		142: uint32(0x14d3f3ed),
		143: uint32(0xdb4de425),
		144: uint32(0x5dbb6cd0),
		145: uint32(0x92257b18),
		146: uint32(0x19f64501),
		147: uint32(0xd66852c9),
		148: uint32(0xd5213f72),
		149: uint32(0x1abf28ba),
		150: uint32(0x916c16a3),
		151: uint32(0x5ef2016b),
		152: uint32(0x97fecdd5),
		153: uint32(0x5860da1d),
		154: uint32(0xd3b3e404),
		155: uint32(0x1c2df3cc),
		156: uint32(0x1f649e77),
		157: uint32(0xd0fa89bf),
		158: uint32(0x5b29b7a6),
		159: uint32(0x94b7a06e),
		160: uint32(0x8db5a00d),
		161: uint32(0x422bb7c5),
		162: uint32(0xc9f889dc),
		163: uint32(0x06669e14),
		164: uint32(0x052ff3af),
		165: uint32(0xcab1e467),
		166: uint32(0x4162da7e),
		167: uint32(0x8efccdb6),
		168: uint32(0x47f00108),
		169: uint32(0x886e16c0),
		170: uint32(0x03bd28d9),
		171: uint32(0xcc233f11),
		172: uint32(0xcf6a52aa),
		173: uint32(0x00f44562),
		174: uint32(0x8b277b7b),
		175: uint32(0x44b96cb3),
		176: uint32(0xc24fe446),
		177: uint32(0x0dd1f38e),
		178: uint32(0x8602cd97),
		179: uint32(0x499cda5f),
		180: uint32(0x4ad5b7e4),
		181: uint32(0x854ba02c),
		182: uint32(0x0e989e35),
		183: uint32(0xc10689fd),
		184: uint32(0x080a4543),
		185: uint32(0xc794528b),
		186: uint32(0x4c476c92),
		187: uint32(0x83d97b5a),
		188: uint32(0x809016e1),
		189: uint32(0x4f0e0129),
		190: uint32(0xc4dd3f30),
		191: uint32(0x0b4328f8),
		192: uint32(0xf6d93ff6),
		193: uint32(0x3947283e),
		194: uint32(0xb2941627),
		195: uint32(0x7d0a01ef),
		196: uint32(0x7e436c54),
		197: uint32(0xb1dd7b9c),
		198: uint32(0x3a0e4585),
		199: uint32(0xf590524d),
		200: uint32(0x3c9c9ef3),
		201: uint32(0xf302893b),
		202: uint32(0x78d1b722),
		203: uint32(0xb74fa0ea),
		204: uint32(0xb406cd51),
		205: uint32(0x7b98da99),
		206: uint32(0xf04be480),
		207: uint32(0x3fd5f348),
		208: uint32(0xb9237bbd),
		209: uint32(0x76bd6c75),
		210: uint32(0xfd6e526c),
		211: uint32(0x32f045a4),
		212: uint32(0x31b9281f),
		213: uint32(0xfe273fd7),
		214: uint32(0x75f401ce),
		215: uint32(0xba6a1606),
		216: uint32(0x7366dab8),
		217: uint32(0xbcf8cd70),
		218: uint32(0x372bf369),
		219: uint32(0xf8b5e4a1),
		220: uint32(0xfbfc891a),
		221: uint32(0x34629ed2),
		222: uint32(0xbfb1a0cb),
		223: uint32(0x702fb703),
		224: uint32(0x692db760),
		225: uint32(0xa6b3a0a8),
		226: uint32(0x2d609eb1),
		227: uint32(0xe2fe8979),
		228: uint32(0xe1b7e4c2),
		229: uint32(0x2e29f30a),
		230: uint32(0xa5facd13),
		231: uint32(0x6a64dadb),
		232: uint32(0xa3681665),
		233: uint32(0x6cf601ad),
		234: uint32(0xe7253fb4),
		235: uint32(0x28bb287c),
		236: uint32(0x2bf245c7),
		237: uint32(0xe46c520f),
		238: uint32(0x6fbf6c16),
		239: uint32(0xa0217bde),
		240: uint32(0x26d7f32b),
		241: uint32(0xe949e4e3),
		242: uint32(0x629adafa),
		243: uint32(0xad04cd32),
		244: uint32(0xae4da089),
		245: uint32(0x61d3b741),
		246: uint32(0xea008958),
		247: uint32(0x259e9e90),
		248: uint32(0xec92522e),
		249: uint32(0x230c45e6),
		250: uint32(0xa8df7bff),
		251: uint32(0x67416c37),
		252: uint32(0x6408018c),
		253: uint32(0xab961644),
		254: uint32(0x2045285d),
		255: uint32(0xefdb3f95),
	},
	6: {
		1:   uint32(0x24825136),
		2:   uint32(0x4904a26c),
		3:   uint32(0x6d86f35a),
		4:   uint32(0x920944d8),
		5:   uint32(0xb68b15ee),
		6:   uint32(0xdb0de6b4),
		7:   uint32(0xff8fb782),
		8:   uint32(0xff638ff1),
		9:   uint32(0xdbe1dec7),
		10:  uint32(0xb6672d9d),
		11:  uint32(0x92e57cab),
		12:  uint32(0x6d6acb29),
		13:  uint32(0x49e89a1f),
		14:  uint32(0x246e6945),
		15:  uint32(0x00ec3873),
		16:  uint32(0x25b619a3),
		17:  uint32(0x01344895),
		18:  uint32(0x6cb2bbcf),
		19:  uint32(0x4830eaf9),
		20:  uint32(0xb7bf5d7b),
		21:  uint32(0x933d0c4d),
		22:  uint32(0xfebbff17),
		23:  uint32(0xda39ae21),
		24:  uint32(0xdad59652),
		25:  uint32(0xfe57c764),
		26:  uint32(0x93d1343e),
		27:  uint32(0xb7536508),
		28:  uint32(0x48dcd28a),
		29:  uint32(0x6c5e83bc),
		30:  uint32(0x01d870e6),
		31:  uint32(0x255a21d0),
		32:  uint32(0x4b6c3346),
		33:  uint32(0x6fee6270),
		34:  uint32(0x0268912a),
		35:  uint32(0x26eac01c),
		36:  uint32(0xd965779e),
		37:  uint32(0xfde726a8),
		38:  uint32(0x9061d5f2),
		39:  uint32(0xb4e384c4),
		40:  uint32(0xb40fbcb7),
		41:  uint32(0x908ded81),
		42:  uint32(0xfd0b1edb),
		43:  uint32(0xd9894fed),
		44:  uint32(0x2606f86f),
		45:  uint32(0x0284a959),
		46:  uint32(0x6f025a03),
		47:  uint32(0x4b800b35),
		48:  uint32(0x6eda2ae5),
		49:  uint32(0x4a587bd3),
		50:  uint32(0x27de8889),
		51:  uint32(0x035cd9bf),
		52:  uint32(0xfcd36e3d),
		53:  uint32(0xd8513f0b),
		54:  uint32(0xb5d7cc51),
		55:  uint32(0x91559d67),
		56:  uint32(0x91b9a514),
		57:  uint32(0xb53bf422),
		58:  uint32(0xd8bd0778),
		59:  uint32(0xfc3f564e),
		60:  uint32(0x03b0e1cc),
		61:  uint32(0x2732b0fa),
		62:  uint32(0x4ab443a0),
		63:  uint32(0x6e361296),
		64:  uint32(0x96d8668c),
		65:  uint32(0xb25a37ba),
		66:  uint32(0xdfdcc4e0),
		67:  uint32(0xfb5e95d6),
		68:  uint32(0x04d12254),
		69:  uint32(0x20537362),
		70:  uint32(0x4dd58038),
		71:  uint32(0x6957d10e),
		72:  uint32(0x69bbe97d),
		73:  uint32(0x4d39b84b),
		74:  uint32(0x20bf4b11),
		75:  uint32(0x043d1a27),
		76:  uint32(0xfbb2ada5),
		77:  uint32(0xdf30fc93),
		78:  uint32(0xb2b60fc9),
		79:  uint32(0x96345eff),
		80:  uint32(0xb36e7f2f),
		81:  uint32(0x97ec2e19),
		82:  uint32(0xfa6add43),
		83:  uint32(0xdee88c75),
		84:  uint32(0x21673bf7),
		85:  uint32(0x05e56ac1),
		86:  uint32(0x6863999b),
		87:  uint32(0x4ce1c8ad),
		88:  uint32(0x4c0df0de),
		89:  uint32(0x688fa1e8),
		90:  uint32(0x050952b2),
		91:  uint32(0x218b0384),
		92:  uint32(0xde04b406),
		93:  uint32(0xfa86e530),
		94:  uint32(0x9700166a),
		95:  uint32(0xb382475c),
		96:  uint32(0xddb455ca),
		97:  uint32(0xf93604fc),
		98:  uint32(0x94b0f7a6),
		99:  uint32(0xb032a690),
		100: uint32(0x4fbd1112),
		101: uint32(0x6b3f4024),
		102: uint32(0x06b9b37e),
		103: uint32(0x223be248),
		104: uint32(0x22d7da3b),
		105: uint32(0x06558b0d),
		106: uint32(0x6bd37857),
		107: uint32(0x4f512961),
		108: uint32(0xb0de9ee3),
		109: uint32(0x945ccfd5),
		110: uint32(0xf9da3c8f),
		111: uint32(0xdd586db9),
		112: uint32(0xf8024c69),
		113: uint32(0xdc801d5f),
		114: uint32(0xb106ee05),
		115: uint32(0x9584bf33),
		116: uint32(0x6a0b08b1),
		117: uint32(0x4e895987),
		118: uint32(0x230faadd),
		119: uint32(0x078dfbeb),
		120: uint32(0x0761c398),
		121: uint32(0x23e392ae),
		122: uint32(0x4e6561f4),
		123: uint32(0x6ae730c2),
		124: uint32(0x95688740),
		125: uint32(0xb1ead676),
		126: uint32(0xdc6c252c),
		127: uint32(0xf8ee741a),
		128: uint32(0xf6c1cb59),
		129: uint32(0xd2439a6f),
		130: uint32(0xbfc56935),
		131: uint32(0x9b473803),
		132: uint32(0x64c88f81),
		133: uint32(0x404adeb7),
		134: uint32(0x2dcc2ded),
		135: uint32(0x094e7cdb),
		136: uint32(0x09a244a8),
		137: uint32(0x2d20159e),
		138: uint32(0x40a6e6c4),
		139: uint32(0x6424b7f2),
		140: uint32(0x9bab0070),
		141: uint32(0xbf295146),
		142: uint32(0xd2afa21c),
		143: uint32(0xf62df32a),
		144: uint32(0xd377d2fa),
		145: uint32(0xf7f583cc),
		146: uint32(0x9a737096),
		147: uint32(0xbef121a0),
		148: uint32(0x417e9622),
		149: uint32(0x65fcc714),
		150: uint32(0x087a344e),
		151: uint32(0x2cf86578),
		152: uint32(0x2c145d0b),
		153: uint32(0x08960c3d),
		154: uint32(0x6510ff67),
		155: uint32(0x4192ae51),
		156: uint32(0xbe1d19d3),
		157: uint32(0x9a9f48e5),
		158: uint32(0xf719bbbf),
		159: uint32(0xd39bea89),
		160: uint32(0xbdadf81f),
		161: uint32(0x992fa929),
		162: uint32(0xf4a95a73),
		163: uint32(0xd02b0b45),
		164: uint32(0x2fa4bcc7),
		165: uint32(0x0b26edf1),
		166: uint32(0x66a01eab),
		167: uint32(0x42224f9d),
		168: uint32(0x42ce77ee),
		169: uint32(0x664c26d8),
		170: uint32(0x0bcad582),
		171: uint32(0x2f4884b4),
		172: uint32(0xd0c73336),
		173: uint32(0xf4456200),
		174: uint32(0x99c3915a),
		175: uint32(0xbd41c06c),
		176: uint32(0x981be1bc),
		177: uint32(0xbc99b08a),
		178: uint32(0xd11f43d0),
		179: uint32(0xf59d12e6),
		180: uint32(0x0a12a564),
		181: uint32(0x2e90f452),
		182: uint32(0x43160708),
		183: uint32(0x6794563e),
		184: uint32(0x67786e4d),
		185: uint32(0x43fa3f7b),
		186: uint32(0x2e7ccc21),
		187: uint32(0x0afe9d17),
		188: uint32(0xf5712a95),
		189: uint32(0xd1f37ba3),
		190: uint32(0xbc7588f9),
		191: uint32(0x98f7d9cf),
		192: uint32(0x6019add5),
		193: uint32(0x449bfce3),
		194: uint32(0x291d0fb9),
		195: uint32(0x0d9f5e8f),
		196: uint32(0xf210e90d),
		197: uint32(0xd692b83b),
		198: uint32(0xbb144b61),
		199: uint32(0x9f961a57),
		200: uint32(0x9f7a2224),
		201: uint32(0xbbf87312),
		202: uint32(0xd67e8048),
		203: uint32(0xf2fcd17e),
		204: uint32(0x0d7366fc),
		205: uint32(0x29f137ca),
		206: uint32(0x4477c490),
		207: uint32(0x60f595a6),
		208: uint32(0x45afb476),
		209: uint32(0x612de540),
		210: uint32(0x0cab161a),
		211: uint32(0x2829472c),
		212: uint32(0xd7a6f0ae),
		213: uint32(0xf324a198),
		214: uint32(0x9ea252c2),
		215: uint32(0xba2003f4),
		216: uint32(0xbacc3b87),
		217: uint32(0x9e4e6ab1),
		218: uint32(0xf3c899eb),
		219: uint32(0xd74ac8dd),
		220: uint32(0x28c57f5f),
		221: uint32(0x0c472e69),
		222: uint32(0x61c1dd33),
		223: uint32(0x45438c05),
		224: uint32(0x2b759e93),
		225: uint32(0x0ff7cfa5),
		226: uint32(0x62713cff),
		227: uint32(0x46f36dc9),
		228: uint32(0xb97cda4b),
		229: uint32(0x9dfe8b7d),
		230: uint32(0xf0787827),
		231: uint32(0xd4fa2911),
		232: uint32(0xd4161162),
		233: uint32(0xf0944054),
		234: uint32(0x9d12b30e),
		235: uint32(0xb990e238),
		236: uint32(0x461f55ba),
		237: uint32(0x629d048c),
		238: uint32(0x0f1bf7d6),
		239: uint32(0x2b99a6e0),
		240: uint32(0x0ec38730),
		241: uint32(0x2a41d606),
		242: uint32(0x47c7255c),
		243: uint32(0x6345746a),
		244: uint32(0x9ccac3e8),
		245: uint32(0xb84892de),
		246: uint32(0xd5ce6184),
		247: uint32(0xf14c30b2),
		248: uint32(0xf1a008c1),
		249: uint32(0xd52259f7),
		250: uint32(0xb8a4aaad),
		251: uint32(0x9c26fb9b),
		252: uint32(0x63a94c19),
		253: uint32(0x472b1d2f),
		254: uint32(0x2aadee75),
		255: uint32(0x0e2fbf43),
	},
	7: {
		1:   uint32(0x36f290f3),
		2:   uint32(0x6de521e6),
		3:   uint32(0x5b17b115),
		4:   uint32(0xdbca43cc),
		5:   uint32(0xed38d33f),
		6:   uint32(0xb62f622a),
		7:   uint32(0x80ddf2d9),
		8:   uint32(0x6ce581d9),
		9:   uint32(0x5a17112a),
		10:  uint32(0x0100a03f),
		11:  uint32(0x37f230cc),
		12:  uint32(0xb72fc215),
		13:  uint32(0x81dd52e6),
		14:  uint32(0xdacae3f3),
		15:  uint32(0xec387300),
		16:  uint32(0xd9cb03b2),
		17:  uint32(0xef399341),
		18:  uint32(0xb42e2254),
		19:  uint32(0x82dcb2a7),
		20:  uint32(0x0201407e),
		21:  uint32(0x34f3d08d),
		22:  uint32(0x6fe46198),
		23:  uint32(0x5916f16b),
		24:  uint32(0xb52e826b),
		25:  uint32(0x83dc1298),
		26:  uint32(0xd8cba38d),
		27:  uint32(0xee39337e),
		28:  uint32(0x6ee4c1a7),
		29:  uint32(0x58165154),
		30:  uint32(0x0301e041),
		31:  uint32(0x35f370b2),
		32:  uint32(0x68e70125),
		33:  uint32(0x5e1591d6),
		34:  uint32(0x050220c3),
		35:  uint32(0x33f0b030),
		36:  uint32(0xb32d42e9),
		37:  uint32(0x85dfd21a),
		38:  uint32(0xdec8630f),
		39:  uint32(0xe83af3fc),
		40:  uint32(0x040280fc),
		41:  uint32(0x32f0100f),
		42:  uint32(0x69e7a11a),
		43:  uint32(0x5f1531e9),
		44:  uint32(0xdfc8c330),
		45:  uint32(0xe93a53c3),
		46:  uint32(0xb22de2d6),
		47:  uint32(0x84df7225),
		48:  uint32(0xb12c0297),
		49:  uint32(0x87de9264),
		50:  uint32(0xdcc92371),
		51:  uint32(0xea3bb382),
		52:  uint32(0x6ae6415b),
		53:  uint32(0x5c14d1a8),
		54:  uint32(0x070360bd),
		55:  uint32(0x31f1f04e),
		56:  uint32(0xddc9834e),
		57:  uint32(0xeb3b13bd),
		58:  uint32(0xb02ca2a8),
		59:  uint32(0x86de325b),
		60:  uint32(0x0603c082),
		61:  uint32(0x30f15071),
		62:  uint32(0x6be6e164),
		63:  uint32(0x5d147197),
		64:  uint32(0xd1ce024a),
		65:  uint32(0xe73c92b9),
		66:  uint32(0xbc2b23ac),
		67:  uint32(0x8ad9b35f),
		68:  uint32(0x0a044186),
		69:  uint32(0x3cf6d175),
		70:  uint32(0x67e16060),
		71:  uint32(0x5113f093),
		72:  uint32(0xbd2b8393),
		73:  uint32(0x8bd91360),
		74:  uint32(0xd0cea275),
		75:  uint32(0xe63c3286),
		76:  uint32(0x66e1c05f),
		77:  uint32(0x501350ac),
		78:  uint32(0x0b04e1b9),
		79:  uint32(0x3df6714a),
		80:  uint32(0x080501f8),
		81:  uint32(0x3ef7910b),
		82:  uint32(0x65e0201e),
		83:  uint32(0x5312b0ed),
		84:  uint32(0xd3cf4234),
		85:  uint32(0xe53dd2c7),
		86:  uint32(0xbe2a63d2),
		87:  uint32(0x88d8f321),
		88:  uint32(0x64e08021),
		89:  uint32(0x521210d2),
		90:  uint32(0x0905a1c7),
		91:  uint32(0x3ff73134),
		92:  uint32(0xbf2ac3ed),
		93:  uint32(0x89d8531e),
		94:  uint32(0xd2cfe20b),
		95:  uint32(0xe43d72f8),
		96:  uint32(0xb929036f),
		97:  uint32(0x8fdb939c),
		98:  uint32(0xd4cc2289),
		99:  uint32(0xe23eb27a),
		100: uint32(0x62e340a3),
		101: uint32(0x5411d050),
		102: uint32(0x0f066145),
		103: uint32(0x39f4f1b6),
		104: uint32(0xd5cc82b6),
		105: uint32(0xe33e1245),
		106: uint32(0xb829a350),
		107: uint32(0x8edb33a3),
		108: uint32(0x0e06c17a),
		109: uint32(0x38f45189),
		110: uint32(0x63e3e09c),
		111: uint32(0x5511706f),
		112: uint32(0x60e200dd),
		113: uint32(0x5610902e),
		114: uint32(0x0d07213b),
		115: uint32(0x3bf5b1c8),
		116: uint32(0xbb284311),
		117: uint32(0x8ddad3e2),
		118: uint32(0xd6cd62f7),
		119: uint32(0xe03ff204),
		120: uint32(0x0c078104),
		121: uint32(0x3af511f7),
		122: uint32(0x61e2a0e2),
		123: uint32(0x57103011),
		124: uint32(0xd7cdc2c8),
		125: uint32(0xe13f523b),
		126: uint32(0xba28e32e),
		127: uint32(0x8cda73dd),
		128: uint32(0x78ed02d5),
		129: uint32(0x4e1f9226),
		130: uint32(0x15082333),
		131: uint32(0x23fab3c0),
		132: uint32(0xa3274119),
		133: uint32(0x95d5d1ea),
		134: uint32(0xcec260ff),
		135: uint32(0xf830f00c),
		136: uint32(0x1408830c),
		137: uint32(0x22fa13ff),
		138: uint32(0x79eda2ea),
		139: uint32(0x4f1f3219),
		140: uint32(0xcfc2c0c0),
		141: uint32(0xf9305033),
		142: uint32(0xa227e126),
		143: uint32(0x94d571d5),
		144: uint32(0xa1260167),
		145: uint32(0x97d49194),
		146: uint32(0xccc32081),
		147: uint32(0xfa31b072),
		148: uint32(0x7aec42ab),
		149: uint32(0x4c1ed258),
		150: uint32(0x1709634d),
		151: uint32(0x21fbf3be),
		152: uint32(0xcdc380be),
		153: uint32(0xfb31104d),
		154: uint32(0xa026a158),
		155: uint32(0x96d431ab),
		156: uint32(0x1609c372),
		157: uint32(0x20fb5381),
		158: uint32(0x7bece294),
		159: uint32(0x4d1e7267),
		160: uint32(0x100a03f0),
		161: uint32(0x26f89303),
		162: uint32(0x7def2216),
		163: uint32(0x4b1db2e5),
		164: uint32(0xcbc0403c),
		165: uint32(0xfd32d0cf),
		166: uint32(0xa62561da),
		167: uint32(0x90d7f129),
		168: uint32(0x7cef8229),
		169: uint32(0x4a1d12da),
		170: uint32(0x110aa3cf),
		171: uint32(0x27f8333c),
		172: uint32(0xa725c1e5),
		173: uint32(0x91d75116),
		174: uint32(0xcac0e003),
		175: uint32(0xfc3270f0),
		176: uint32(0xc9c10042),
		177: uint32(0xff3390b1),
		178: uint32(0xa42421a4),
		179: uint32(0x92d6b157),
		180: uint32(0x120b438e),
		181: uint32(0x24f9d37d),
		182: uint32(0x7fee6268),
		183: uint32(0x491cf29b),
		184: uint32(0xa524819b),
		185: uint32(0x93d61168),
		186: uint32(0xc8c1a07d),
		187: uint32(0xfe33308e),
		188: uint32(0x7eeec257),
		189: uint32(0x481c52a4),
		190: uint32(0x130be3b1),
		191: uint32(0x25f97342),
		192: uint32(0xa923009f),
		193: uint32(0x9fd1906c),
		194: uint32(0xc4c62179),
		195: uint32(0xf234b18a),
		196: uint32(0x72e94353),
		197: uint32(0x441bd3a0),
		198: uint32(0x1f0c62b5),
		199: uint32(0x29fef246),
		200: uint32(0xc5c68146),
		201: uint32(0xf33411b5),
		202: uint32(0xa823a0a0),
		203: uint32(0x9ed13053),
		204: uint32(0x1e0cc28a),
		205: uint32(0x28fe5279),
		206: uint32(0x73e9e36c),
		207: uint32(0x451b739f),
		208: uint32(0x70e8032d),
		209: uint32(0x461a93de),
		210: uint32(0x1d0d22cb),
		211: uint32(0x2bffb238),
		212: uint32(0xab2240e1),
		213: uint32(0x9dd0d012),
		214: uint32(0xc6c76107),
		215: uint32(0xf035f1f4),
		216: uint32(0x1c0d82f4),
		217: uint32(0x2aff1207),
		218: uint32(0x71e8a312),
		219: uint32(0x471a33e1),
		220: uint32(0xc7c7c138),
		221: uint32(0xf13551cb),
		222: uint32(0xaa22e0de),
		223: uint32(0x9cd0702d),
		224: uint32(0xc1c401ba),
		225: uint32(0xf7369149),
		226: uint32(0xac21205c),
		227: uint32(0x9ad3b0af),
		228: uint32(0x1a0e4276),
		229: uint32(0x2cfcd285),
		230: uint32(0x77eb6390),
		231: uint32(0x4119f363),
		232: uint32(0xad218063),
		233: uint32(0x9bd31090),
		234: uint32(0xc0c4a185),
		235: uint32(0xf6363176),
		236: uint32(0x76ebc3af),
		237: uint32(0x4019535c),
		238: uint32(0x1b0ee249),
		239: uint32(0x2dfc72ba),
		240: uint32(0x180f0208),
		241: uint32(0x2efd92fb),
		242: uint32(0x75ea23ee),
		243: uint32(0x4318b31d),
		244: uint32(0xc3c541c4),
		245: uint32(0xf537d137),
		246: uint32(0xae206022),
		247: uint32(0x98d2f0d1),
		248: uint32(0x74ea83d1),
		249: uint32(0x42181322),
		250: uint32(0x190fa237),
		251: uint32(0x2ffd32c4),
		252: uint32(0xaf20c01d),
		253: uint32(0x99d250ee),
		254: uint32(0xc2c5e1fb),
		255: uint32(0xf4377108),
	},
}
var _crc_braid_big_table = [8][256]Tz_word_t{
	0: {
		1:   uint64(0xf390f23600000000),
		2:   uint64(0xe621e56d00000000),
		3:   uint64(0x15b1175b00000000),
		4:   uint64(0xcc43cadb00000000),
		5:   uint64(0x3fd338ed00000000),
		6:   uint64(0x2a622fb600000000),
		7:   uint64(0xd9f2dd8000000000),
		8:   uint64(0xd981e56c00000000),
		9:   uint64(0x2a11175a00000000),
		10:  uint64(0x3fa0000100000000),
		11:  uint64(0xcc30f23700000000),
		12:  uint64(0x15c22fb700000000),
		13:  uint64(0xe652dd8100000000),
		14:  uint64(0xf3e3cada00000000),
		15:  uint64(0x007338ec00000000),
		16:  uint64(0xb203cbd900000000),
		17:  uint64(0x419339ef00000000),
		18:  uint64(0x54222eb400000000),
		19:  uint64(0xa7b2dc8200000000),
		20:  uint64(0x7e40010200000000),
		21:  uint64(0x8dd0f33400000000),
		22:  uint64(0x9861e46f00000000),
		23:  uint64(0x6bf1165900000000),
		24:  uint64(0x6b822eb500000000),
		25:  uint64(0x9812dc8300000000),
		26:  uint64(0x8da3cbd800000000),
		27:  uint64(0x7e3339ee00000000),
		28:  uint64(0xa7c1e46e00000000),
		29:  uint64(0x5451165800000000),
		30:  uint64(0x41e0010300000000),
		31:  uint64(0xb270f33500000000),
		32:  uint64(0x2501e76800000000),
		33:  uint64(0xd691155e00000000),
		34:  uint64(0xc320020500000000),
		35:  uint64(0x30b0f03300000000),
		36:  uint64(0xe9422db300000000),
		37:  uint64(0x1ad2df8500000000),
		38:  uint64(0x0f63c8de00000000),
		39:  uint64(0xfcf33ae800000000),
		40:  uint64(0xfc80020400000000),
		41:  uint64(0x0f10f03200000000),
		42:  uint64(0x1aa1e76900000000),
		43:  uint64(0xe931155f00000000),
		44:  uint64(0x30c3c8df00000000),
		45:  uint64(0xc3533ae900000000),
		46:  uint64(0xd6e22db200000000),
		47:  uint64(0x2572df8400000000),
		48:  uint64(0x97022cb100000000),
		49:  uint64(0x6492de8700000000),
		50:  uint64(0x7123c9dc00000000),
		51:  uint64(0x82b33bea00000000),
		52:  uint64(0x5b41e66a00000000),
		53:  uint64(0xa8d1145c00000000),
		54:  uint64(0xbd60030700000000),
		55:  uint64(0x4ef0f13100000000),
		56:  uint64(0x4e83c9dd00000000),
		57:  uint64(0xbd133beb00000000),
		58:  uint64(0xa8a22cb000000000),
		59:  uint64(0x5b32de8600000000),
		60:  uint64(0x82c0030600000000),
		61:  uint64(0x7150f13000000000),
		62:  uint64(0x64e1e66b00000000),
		63:  uint64(0x9771145d00000000),
		64:  uint64(0x4a02ced100000000),
		65:  uint64(0xb9923ce700000000),
		66:  uint64(0xac232bbc00000000),
		67:  uint64(0x5fb3d98a00000000),
		68:  uint64(0x8641040a00000000),
		69:  uint64(0x75d1f63c00000000),
		70:  uint64(0x6060e16700000000),
		71:  uint64(0x93f0135100000000),
		72:  uint64(0x93832bbd00000000),
		73:  uint64(0x6013d98b00000000),
		74:  uint64(0x75a2ced000000000),
		75:  uint64(0x86323ce600000000),
		76:  uint64(0x5fc0e16600000000),
		77:  uint64(0xac50135000000000),
		78:  uint64(0xb9e1040b00000000),
		79:  uint64(0x4a71f63d00000000),
		80:  uint64(0xf801050800000000),
		81:  uint64(0x0b91f73e00000000),
		82:  uint64(0x1e20e06500000000),
		83:  uint64(0xedb0125300000000),
		84:  uint64(0x3442cfd300000000),
		85:  uint64(0xc7d23de500000000),
		86:  uint64(0xd2632abe00000000),
		87:  uint64(0x21f3d88800000000),
		88:  uint64(0x2180e06400000000),
		89:  uint64(0xd210125200000000),
		90:  uint64(0xc7a1050900000000),
		91:  uint64(0x3431f73f00000000),
		92:  uint64(0xedc32abf00000000),
		93:  uint64(0x1e53d88900000000),
		94:  uint64(0x0be2cfd200000000),
		95:  uint64(0xf8723de400000000),
		96:  uint64(0x6f0329b900000000),
		97:  uint64(0x9c93db8f00000000),
		98:  uint64(0x8922ccd400000000),
		99:  uint64(0x7ab23ee200000000),
		100: uint64(0xa340e36200000000),
		101: uint64(0x50d0115400000000),
		102: uint64(0x4561060f00000000),
		103: uint64(0xb6f1f43900000000),
		104: uint64(0xb682ccd500000000),
		105: uint64(0x45123ee300000000),
		106: uint64(0x50a329b800000000),
		107: uint64(0xa333db8e00000000),
		108: uint64(0x7ac1060e00000000),
		109: uint64(0x8951f43800000000),
		110: uint64(0x9ce0e36300000000),
		111: uint64(0x6f70115500000000),
		112: uint64(0xdd00e26000000000),
		113: uint64(0x2e90105600000000),
		114: uint64(0x3b21070d00000000),
		115: uint64(0xc8b1f53b00000000),
		116: uint64(0x114328bb00000000),
		117: uint64(0xe2d3da8d00000000),
		118: uint64(0xf762cdd600000000),
		119: uint64(0x04f23fe000000000),
		120: uint64(0x0481070c00000000),
		121: uint64(0xf711f53a00000000),
		122: uint64(0xe2a0e26100000000),
		123: uint64(0x1130105700000000),
		124: uint64(0xc8c2cdd700000000),
		125: uint64(0x3b523fe100000000),
		126: uint64(0x2ee328ba00000000),
		127: uint64(0xdd73da8c00000000),
		128: uint64(0xd502ed7800000000),
		129: uint64(0x26921f4e00000000),
		130: uint64(0x3323081500000000),
		131: uint64(0xc0b3fa2300000000),
		132: uint64(0x194127a300000000),
		133: uint64(0xead1d59500000000),
		134: uint64(0xff60c2ce00000000),
		135: uint64(0x0cf030f800000000),
		136: uint64(0x0c83081400000000),
		137: uint64(0xff13fa2200000000),
		138: uint64(0xeaa2ed7900000000),
		139: uint64(0x19321f4f00000000),
		140: uint64(0xc0c0c2cf00000000),
		141: uint64(0x335030f900000000),
		142: uint64(0x26e127a200000000),
		143: uint64(0xd571d59400000000),
		144: uint64(0x670126a100000000),
		145: uint64(0x9491d49700000000),
		146: uint64(0x8120c3cc00000000),
		147: uint64(0x72b031fa00000000),
		148: uint64(0xab42ec7a00000000),
		149: uint64(0x58d21e4c00000000),
		150: uint64(0x4d63091700000000),
		151: uint64(0xbef3fb2100000000),
		152: uint64(0xbe80c3cd00000000),
		153: uint64(0x4d1031fb00000000),
		154: uint64(0x58a126a000000000),
		155: uint64(0xab31d49600000000),
		156: uint64(0x72c3091600000000),
		157: uint64(0x8153fb2000000000),
		158: uint64(0x94e2ec7b00000000),
		159: uint64(0x67721e4d00000000),
		160: uint64(0xf0030a1000000000),
		161: uint64(0x0393f82600000000),
		162: uint64(0x1622ef7d00000000),
		163: uint64(0xe5b21d4b00000000),
		164: uint64(0x3c40c0cb00000000),
		165: uint64(0xcfd032fd00000000),
		166: uint64(0xda6125a600000000),
		167: uint64(0x29f1d79000000000),
		168: uint64(0x2982ef7c00000000),
		169: uint64(0xda121d4a00000000),
		170: uint64(0xcfa30a1100000000),
		171: uint64(0x3c33f82700000000),
		172: uint64(0xe5c125a700000000),
		173: uint64(0x1651d79100000000),
		174: uint64(0x03e0c0ca00000000),
		175: uint64(0xf07032fc00000000),
		176: uint64(0x4200c1c900000000),
		177: uint64(0xb19033ff00000000),
		178: uint64(0xa42124a400000000),
		179: uint64(0x57b1d69200000000),
		180: uint64(0x8e430b1200000000),
		181: uint64(0x7dd3f92400000000),
		182: uint64(0x6862ee7f00000000),
		183: uint64(0x9bf21c4900000000),
		184: uint64(0x9b8124a500000000),
		185: uint64(0x6811d69300000000),
		186: uint64(0x7da0c1c800000000),
		187: uint64(0x8e3033fe00000000),
		188: uint64(0x57c2ee7e00000000),
		189: uint64(0xa4521c4800000000),
		190: uint64(0xb1e30b1300000000),
		191: uint64(0x4273f92500000000),
		192: uint64(0x9f0023a900000000),
		193: uint64(0x6c90d19f00000000),
		194: uint64(0x7921c6c400000000),
		195: uint64(0x8ab134f200000000),
		196: uint64(0x5343e97200000000),
		197: uint64(0xa0d31b4400000000),
		198: uint64(0xb5620c1f00000000),
		199: uint64(0x46f2fe2900000000),
		200: uint64(0x4681c6c500000000),
		201: uint64(0xb51134f300000000),
		202: uint64(0xa0a023a800000000),
		203: uint64(0x5330d19e00000000),
		204: uint64(0x8ac20c1e00000000),
		205: uint64(0x7952fe2800000000),
		206: uint64(0x6ce3e97300000000),
		207: uint64(0x9f731b4500000000),
		208: uint64(0x2d03e87000000000),
		209: uint64(0xde931a4600000000),
		210: uint64(0xcb220d1d00000000),
		211: uint64(0x38b2ff2b00000000),
		212: uint64(0xe14022ab00000000),
		213: uint64(0x12d0d09d00000000),
		214: uint64(0x0761c7c600000000),
		215: uint64(0xf4f135f000000000),
		216: uint64(0xf4820d1c00000000),
		217: uint64(0x0712ff2a00000000),
		218: uint64(0x12a3e87100000000),
		219: uint64(0xe1331a4700000000),
		220: uint64(0x38c1c7c700000000),
		221: uint64(0xcb5135f100000000),
		222: uint64(0xdee022aa00000000),
		223: uint64(0x2d70d09c00000000),
		224: uint64(0xba01c4c100000000),
		225: uint64(0x499136f700000000),
		226: uint64(0x5c2021ac00000000),
		227: uint64(0xafb0d39a00000000),
		228: uint64(0x76420e1a00000000),
		229: uint64(0x85d2fc2c00000000),
		230: uint64(0x9063eb7700000000),
		231: uint64(0x63f3194100000000),
		232: uint64(0x638021ad00000000),
		233: uint64(0x9010d39b00000000),
		234: uint64(0x85a1c4c000000000),
		235: uint64(0x763136f600000000),
		236: uint64(0xafc3eb7600000000),
		237: uint64(0x5c53194000000000),
		238: uint64(0x49e20e1b00000000),
		239: uint64(0xba72fc2d00000000),
		240: uint64(0x08020f1800000000),
		241: uint64(0xfb92fd2e00000000),
		242: uint64(0xee23ea7500000000),
		243: uint64(0x1db3184300000000),
		244: uint64(0xc441c5c300000000),
		245: uint64(0x37d137f500000000),
		246: uint64(0x226020ae00000000),
		247: uint64(0xd1f0d29800000000),
		248: uint64(0xd183ea7400000000),
		249: uint64(0x2213184200000000),
		250: uint64(0x37a20f1900000000),
		251: uint64(0xc432fd2f00000000),
		252: uint64(0x1dc020af00000000),
		253: uint64(0xee50d29900000000),
		254: uint64(0xfbe1c5c200000000),
		255: uint64(0x087137f400000000),
	},
	1: {
		1:   uint64(0x3651822400000000),
		2:   uint64(0x6ca2044900000000),
		3:   uint64(0x5af3866d00000000),
		4:   uint64(0xd844099200000000),
		5:   uint64(0xee158bb600000000),
		6:   uint64(0xb4e60ddb00000000),
		7:   uint64(0x82b78fff00000000),
		8:   uint64(0xf18f63ff00000000),
		9:   uint64(0xc7dee1db00000000),
		10:  uint64(0x9d2d67b600000000),
		11:  uint64(0xab7ce59200000000),
		12:  uint64(0x29cb6a6d00000000),
		13:  uint64(0x1f9ae84900000000),
		14:  uint64(0x45696e2400000000),
		15:  uint64(0x7338ec0000000000),
		16:  uint64(0xa319b62500000000),
		17:  uint64(0x9548340100000000),
		18:  uint64(0xcfbbb26c00000000),
		19:  uint64(0xf9ea304800000000),
		20:  uint64(0x7b5dbfb700000000),
		21:  uint64(0x4d0c3d9300000000),
		22:  uint64(0x17ffbbfe00000000),
		23:  uint64(0x21ae39da00000000),
		24:  uint64(0x5296d5da00000000),
		25:  uint64(0x64c757fe00000000),
		26:  uint64(0x3e34d19300000000),
		27:  uint64(0x086553b700000000),
		28:  uint64(0x8ad2dc4800000000),
		29:  uint64(0xbc835e6c00000000),
		30:  uint64(0xe670d80100000000),
		31:  uint64(0xd0215a2500000000),
		32:  uint64(0x46336c4b00000000),
		33:  uint64(0x7062ee6f00000000),
		34:  uint64(0x2a91680200000000),
		35:  uint64(0x1cc0ea2600000000),
		36:  uint64(0x9e7765d900000000),
		37:  uint64(0xa826e7fd00000000),
		38:  uint64(0xf2d5619000000000),
		39:  uint64(0xc484e3b400000000),
		40:  uint64(0xb7bc0fb400000000),
		41:  uint64(0x81ed8d9000000000),
		42:  uint64(0xdb1e0bfd00000000),
		43:  uint64(0xed4f89d900000000),
		44:  uint64(0x6ff8062600000000),
		45:  uint64(0x59a9840200000000),
		46:  uint64(0x035a026f00000000),
		47:  uint64(0x350b804b00000000),
		48:  uint64(0xe52ada6e00000000),
		49:  uint64(0xd37b584a00000000),
		50:  uint64(0x8988de2700000000),
		51:  uint64(0xbfd95c0300000000),
		52:  uint64(0x3d6ed3fc00000000),
		53:  uint64(0x0b3f51d800000000),
		54:  uint64(0x51ccd7b500000000),
		55:  uint64(0x679d559100000000),
		56:  uint64(0x14a5b99100000000),
		57:  uint64(0x22f43bb500000000),
		58:  uint64(0x7807bdd800000000),
		59:  uint64(0x4e563ffc00000000),
		60:  uint64(0xcce1b00300000000),
		61:  uint64(0xfab0322700000000),
		62:  uint64(0xa043b44a00000000),
		63:  uint64(0x9612366e00000000),
		64:  uint64(0x8c66d89600000000),
		65:  uint64(0xba375ab200000000),
		66:  uint64(0xe0c4dcdf00000000),
		67:  uint64(0xd6955efb00000000),
		68:  uint64(0x5422d10400000000),
		69:  uint64(0x6273532000000000),
		70:  uint64(0x3880d54d00000000),
		71:  uint64(0x0ed1576900000000),
		72:  uint64(0x7de9bb6900000000),
		73:  uint64(0x4bb8394d00000000),
		74:  uint64(0x114bbf2000000000),
		75:  uint64(0x271a3d0400000000),
		76:  uint64(0xa5adb2fb00000000),
		77:  uint64(0x93fc30df00000000),
		78:  uint64(0xc90fb6b200000000),
		79:  uint64(0xff5e349600000000),
		80:  uint64(0x2f7f6eb300000000),
		81:  uint64(0x192eec9700000000),
		82:  uint64(0x43dd6afa00000000),
		83:  uint64(0x758ce8de00000000),
		84:  uint64(0xf73b672100000000),
		85:  uint64(0xc16ae50500000000),
		86:  uint64(0x9b99636800000000),
		87:  uint64(0xadc8e14c00000000),
		88:  uint64(0xdef00d4c00000000),
		89:  uint64(0xe8a18f6800000000),
		90:  uint64(0xb252090500000000),
		91:  uint64(0x84038b2100000000),
		92:  uint64(0x06b404de00000000),
		93:  uint64(0x30e586fa00000000),
		94:  uint64(0x6a16009700000000),
		95:  uint64(0x5c4782b300000000),
		96:  uint64(0xca55b4dd00000000),
		97:  uint64(0xfc0436f900000000),
		98:  uint64(0xa6f7b09400000000),
		99:  uint64(0x90a632b000000000),
		100: uint64(0x1211bd4f00000000),
		101: uint64(0x24403f6b00000000),
		102: uint64(0x7eb3b90600000000),
		103: uint64(0x48e23b2200000000),
		104: uint64(0x3bdad72200000000),
		105: uint64(0x0d8b550600000000),
		106: uint64(0x5778d36b00000000),
		107: uint64(0x6129514f00000000),
		108: uint64(0xe39edeb000000000),
		109: uint64(0xd5cf5c9400000000),
		110: uint64(0x8f3cdaf900000000),
		111: uint64(0xb96d58dd00000000),
		112: uint64(0x694c02f800000000),
		113: uint64(0x5f1d80dc00000000),
		114: uint64(0x05ee06b100000000),
		115: uint64(0x33bf849500000000),
		116: uint64(0xb1080b6a00000000),
		117: uint64(0x8759894e00000000),
		118: uint64(0xddaa0f2300000000),
		119: uint64(0xebfb8d0700000000),
		120: uint64(0x98c3610700000000),
		121: uint64(0xae92e32300000000),
		122: uint64(0xf461654e00000000),
		123: uint64(0xc230e76a00000000),
		124: uint64(0x4087689500000000),
		125: uint64(0x76d6eab100000000),
		126: uint64(0x2c256cdc00000000),
		127: uint64(0x1a74eef800000000),
		128: uint64(0x59cbc1f600000000),
		129: uint64(0x6f9a43d200000000),
		130: uint64(0x3569c5bf00000000),
		131: uint64(0x0338479b00000000),
		132: uint64(0x818fc86400000000),
		133: uint64(0xb7de4a4000000000),
		134: uint64(0xed2dcc2d00000000),
		135: uint64(0xdb7c4e0900000000),
		136: uint64(0xa844a20900000000),
		137: uint64(0x9e15202d00000000),
		138: uint64(0xc4e6a64000000000),
		139: uint64(0xf2b7246400000000),
		140: uint64(0x7000ab9b00000000),
		141: uint64(0x465129bf00000000),
		142: uint64(0x1ca2afd200000000),
		143: uint64(0x2af32df600000000),
		144: uint64(0xfad277d300000000),
		145: uint64(0xcc83f5f700000000),
		146: uint64(0x9670739a00000000),
		147: uint64(0xa021f1be00000000),
		148: uint64(0x22967e4100000000),
		149: uint64(0x14c7fc6500000000),
		150: uint64(0x4e347a0800000000),
		151: uint64(0x7865f82c00000000),
		152: uint64(0x0b5d142c00000000),
		153: uint64(0x3d0c960800000000),
		154: uint64(0x67ff106500000000),
		155: uint64(0x51ae924100000000),
		156: uint64(0xd3191dbe00000000),
		157: uint64(0xe5489f9a00000000),
		158: uint64(0xbfbb19f700000000),
		159: uint64(0x89ea9bd300000000),
		160: uint64(0x1ff8adbd00000000),
		161: uint64(0x29a92f9900000000),
		162: uint64(0x735aa9f400000000),
		163: uint64(0x450b2bd000000000),
		164: uint64(0xc7bca42f00000000),
		165: uint64(0xf1ed260b00000000),
		166: uint64(0xab1ea06600000000),
		167: uint64(0x9d4f224200000000),
		168: uint64(0xee77ce4200000000),
		169: uint64(0xd8264c6600000000),
		170: uint64(0x82d5ca0b00000000),
		171: uint64(0xb484482f00000000),
		172: uint64(0x3633c7d000000000),
		173: uint64(0x006245f400000000),
		174: uint64(0x5a91c39900000000),
		175: uint64(0x6cc041bd00000000),
		176: uint64(0xbce11b9800000000),
		177: uint64(0x8ab099bc00000000),
		178: uint64(0xd0431fd100000000),
		179: uint64(0xe6129df500000000),
		180: uint64(0x64a5120a00000000),
		181: uint64(0x52f4902e00000000),
		182: uint64(0x0807164300000000),
		183: uint64(0x3e56946700000000),
		184: uint64(0x4d6e786700000000),
		185: uint64(0x7b3ffa4300000000),
		186: uint64(0x21cc7c2e00000000),
		187: uint64(0x179dfe0a00000000),
		188: uint64(0x952a71f500000000),
		189: uint64(0xa37bf3d100000000),
		190: uint64(0xf98875bc00000000),
		191: uint64(0xcfd9f79800000000),
		192: uint64(0xd5ad196000000000),
		193: uint64(0xe3fc9b4400000000),
		194: uint64(0xb90f1d2900000000),
		195: uint64(0x8f5e9f0d00000000),
		196: uint64(0x0de910f200000000),
		197: uint64(0x3bb892d600000000),
		198: uint64(0x614b14bb00000000),
		199: uint64(0x571a969f00000000),
		200: uint64(0x24227a9f00000000),
		201: uint64(0x1273f8bb00000000),
		202: uint64(0x48807ed600000000),
		203: uint64(0x7ed1fcf200000000),
		204: uint64(0xfc66730d00000000),
		205: uint64(0xca37f12900000000),
		206: uint64(0x90c4774400000000),
		207: uint64(0xa695f56000000000),
		208: uint64(0x76b4af4500000000),
		209: uint64(0x40e52d6100000000),
		210: uint64(0x1a16ab0c00000000),
		211: uint64(0x2c47292800000000),
		212: uint64(0xaef0a6d700000000),
		213: uint64(0x98a124f300000000),
		214: uint64(0xc252a29e00000000),
		215: uint64(0xf40320ba00000000),
		216: uint64(0x873bccba00000000),
		217: uint64(0xb16a4e9e00000000),
		218: uint64(0xeb99c8f300000000),
		219: uint64(0xddc84ad700000000),
		220: uint64(0x5f7fc52800000000),
		221: uint64(0x692e470c00000000),
		222: uint64(0x33ddc16100000000),
		223: uint64(0x058c434500000000),
		224: uint64(0x939e752b00000000),
		225: uint64(0xa5cff70f00000000),
		226: uint64(0xff3c716200000000),
		227: uint64(0xc96df34600000000),
		228: uint64(0x4bda7cb900000000),
		229: uint64(0x7d8bfe9d00000000),
		230: uint64(0x277878f000000000),
		231: uint64(0x1129fad400000000),
		232: uint64(0x621116d400000000),
		233: uint64(0x544094f000000000),
		234: uint64(0x0eb3129d00000000),
		235: uint64(0x38e290b900000000),
		236: uint64(0xba551f4600000000),
		237: uint64(0x8c049d6200000000),
		238: uint64(0xd6f71b0f00000000),
		239: uint64(0xe0a6992b00000000),
		240: uint64(0x3087c30e00000000),
		241: uint64(0x06d6412a00000000),
		242: uint64(0x5c25c74700000000),
		243: uint64(0x6a74456300000000),
		244: uint64(0xe8c3ca9c00000000),
		245: uint64(0xde9248b800000000),
		246: uint64(0x8461ced500000000),
		247: uint64(0xb2304cf100000000),
		248: uint64(0xc108a0f100000000),
		249: uint64(0xf75922d500000000),
		250: uint64(0xadaaa4b800000000),
		251: uint64(0x9bfb269c00000000),
		252: uint64(0x194ca96300000000),
		253: uint64(0x2f1d2b4700000000),
		254: uint64(0x75eead2a00000000),
		255: uint64(0x43bf2f0e00000000),
	},
	2: {
		1:   uint64(0xc8179ecf00000000),
		2:   uint64(0xd1294d4400000000),
		3:   uint64(0x193ed38b00000000),
		4:   uint64(0xa2539a8800000000),
		5:   uint64(0x6a44044700000000),
		6:   uint64(0x737ad7cc00000000),
		7:   uint64(0xbb6d490300000000),
		8:   uint64(0x05a145ca00000000),
		9:   uint64(0xcdb6db0500000000),
		10:  uint64(0xd488088e00000000),
		11:  uint64(0x1c9f964100000000),
		12:  uint64(0xa7f2df4200000000),
		13:  uint64(0x6fe5418d00000000),
		14:  uint64(0x76db920600000000),
		15:  uint64(0xbecc0cc900000000),
		16:  uint64(0x4b44fa4f00000000),
		17:  uint64(0x8353648000000000),
		18:  uint64(0x9a6db70b00000000),
		19:  uint64(0x527a29c400000000),
		20:  uint64(0xe91760c700000000),
		21:  uint64(0x2100fe0800000000),
		22:  uint64(0x383e2d8300000000),
		23:  uint64(0xf029b34c00000000),
		24:  uint64(0x4ee5bf8500000000),
		25:  uint64(0x86f2214a00000000),
		26:  uint64(0x9fccf2c100000000),
		27:  uint64(0x57db6c0e00000000),
		28:  uint64(0xecb6250d00000000),
		29:  uint64(0x24a1bbc200000000),
		30:  uint64(0x3d9f684900000000),
		31:  uint64(0xf588f68600000000),
		32:  uint64(0x9688f49f00000000),
		33:  uint64(0x5e9f6a5000000000),
		34:  uint64(0x47a1b9db00000000),
		35:  uint64(0x8fb6271400000000),
		36:  uint64(0x34db6e1700000000),
		37:  uint64(0xfcccf0d800000000),
		38:  uint64(0xe5f2235300000000),
		39:  uint64(0x2de5bd9c00000000),
		40:  uint64(0x9329b15500000000),
		41:  uint64(0x5b3e2f9a00000000),
		42:  uint64(0x4200fc1100000000),
		43:  uint64(0x8a1762de00000000),
		44:  uint64(0x317a2bdd00000000),
		45:  uint64(0xf96db51200000000),
		46:  uint64(0xe053669900000000),
		47:  uint64(0x2844f85600000000),
		48:  uint64(0xddcc0ed000000000),
		49:  uint64(0x15db901f00000000),
		50:  uint64(0x0ce5439400000000),
		51:  uint64(0xc4f2dd5b00000000),
		52:  uint64(0x7f9f945800000000),
		53:  uint64(0xb7880a9700000000),
		54:  uint64(0xaeb6d91c00000000),
		55:  uint64(0x66a147d300000000),
		56:  uint64(0xd86d4b1a00000000),
		57:  uint64(0x107ad5d500000000),
		58:  uint64(0x0944065e00000000),
		59:  uint64(0xc153989100000000),
		60:  uint64(0x7a3ed19200000000),
		61:  uint64(0xb2294f5d00000000),
		62:  uint64(0xab179cd600000000),
		63:  uint64(0x6300021900000000),
		64:  uint64(0x6d1798e400000000),
		65:  uint64(0xa500062b00000000),
		66:  uint64(0xbc3ed5a000000000),
		67:  uint64(0x74294b6f00000000),
		68:  uint64(0xcf44026c00000000),
		69:  uint64(0x07539ca300000000),
		70:  uint64(0x1e6d4f2800000000),
		71:  uint64(0xd67ad1e700000000),
		72:  uint64(0x68b6dd2e00000000),
		73:  uint64(0xa0a143e100000000),
		74:  uint64(0xb99f906a00000000),
		75:  uint64(0x71880ea500000000),
		76:  uint64(0xcae547a600000000),
		77:  uint64(0x02f2d96900000000),
		78:  uint64(0x1bcc0ae200000000),
		79:  uint64(0xd3db942d00000000),
		80:  uint64(0x265362ab00000000),
		81:  uint64(0xee44fc6400000000),
		82:  uint64(0xf77a2fef00000000),
		83:  uint64(0x3f6db12000000000),
		84:  uint64(0x8400f82300000000),
		85:  uint64(0x4c1766ec00000000),
		86:  uint64(0x5529b56700000000),
		87:  uint64(0x9d3e2ba800000000),
		88:  uint64(0x23f2276100000000),
		89:  uint64(0xebe5b9ae00000000),
		90:  uint64(0xf2db6a2500000000),
		91:  uint64(0x3accf4ea00000000),
		92:  uint64(0x81a1bde900000000),
		93:  uint64(0x49b6232600000000),
		94:  uint64(0x5088f0ad00000000),
		95:  uint64(0x989f6e6200000000),
		96:  uint64(0xfb9f6c7b00000000),
		97:  uint64(0x3388f2b400000000),
		98:  uint64(0x2ab6213f00000000),
		99:  uint64(0xe2a1bff000000000),
		100: uint64(0x59ccf6f300000000),
		101: uint64(0x91db683c00000000),
		102: uint64(0x88e5bbb700000000),
		103: uint64(0x40f2257800000000),
		104: uint64(0xfe3e29b100000000),
		105: uint64(0x3629b77e00000000),
		106: uint64(0x2f1764f500000000),
		107: uint64(0xe700fa3a00000000),
		108: uint64(0x5c6db33900000000),
		109: uint64(0x947a2df600000000),
		110: uint64(0x8d44fe7d00000000),
		111: uint64(0x455360b200000000),
		112: uint64(0xb0db963400000000),
		113: uint64(0x78cc08fb00000000),
		114: uint64(0x61f2db7000000000),
		115: uint64(0xa9e545bf00000000),
		116: uint64(0x12880cbc00000000),
		117: uint64(0xda9f927300000000),
		118: uint64(0xc3a141f800000000),
		119: uint64(0x0bb6df3700000000),
		120: uint64(0xb57ad3fe00000000),
		121: uint64(0x7d6d4d3100000000),
		122: uint64(0x64539eba00000000),
		123: uint64(0xac44007500000000),
		124: uint64(0x1729497600000000),
		125: uint64(0xdf3ed7b900000000),
		126: uint64(0xc600043200000000),
		127: uint64(0x0e179afd00000000),
		128: uint64(0x9b28411200000000),
		129: uint64(0x533fdfdd00000000),
		130: uint64(0x4a010c5600000000),
		131: uint64(0x8216929900000000),
		132: uint64(0x397bdb9a00000000),
		133: uint64(0xf16c455500000000),
		134: uint64(0xe85296de00000000),
		135: uint64(0x2045081100000000),
		136: uint64(0x9e8904d800000000),
		137: uint64(0x569e9a1700000000),
		138: uint64(0x4fa0499c00000000),
		139: uint64(0x87b7d75300000000),
		140: uint64(0x3cda9e5000000000),
		141: uint64(0xf4cd009f00000000),
		142: uint64(0xedf3d31400000000),
		143: uint64(0x25e44ddb00000000),
		144: uint64(0xd06cbb5d00000000),
		145: uint64(0x187b259200000000),
		146: uint64(0x0145f61900000000),
		147: uint64(0xc95268d600000000),
		148: uint64(0x723f21d500000000),
		149: uint64(0xba28bf1a00000000),
		150: uint64(0xa3166c9100000000),
		151: uint64(0x6b01f25e00000000),
		152: uint64(0xd5cdfe9700000000),
		153: uint64(0x1dda605800000000),
		154: uint64(0x04e4b3d300000000),
		155: uint64(0xccf32d1c00000000),
		156: uint64(0x779e641f00000000),
		157: uint64(0xbf89fad000000000),
		158: uint64(0xa6b7295b00000000),
		159: uint64(0x6ea0b79400000000),
		160: uint64(0x0da0b58d00000000),
		161: uint64(0xc5b72b4200000000),
		162: uint64(0xdc89f8c900000000),
		163: uint64(0x149e660600000000),
		164: uint64(0xaff32f0500000000),
		165: uint64(0x67e4b1ca00000000),
		166: uint64(0x7eda624100000000),
		167: uint64(0xb6cdfc8e00000000),
		168: uint64(0x0801f04700000000),
		169: uint64(0xc0166e8800000000),
		170: uint64(0xd928bd0300000000),
		171: uint64(0x113f23cc00000000),
		172: uint64(0xaa526acf00000000),
		173: uint64(0x6245f40000000000),
		174: uint64(0x7b7b278b00000000),
		175: uint64(0xb36cb94400000000),
		176: uint64(0x46e44fc200000000),
		177: uint64(0x8ef3d10d00000000),
		178: uint64(0x97cd028600000000),
		179: uint64(0x5fda9c4900000000),
		180: uint64(0xe4b7d54a00000000),
		181: uint64(0x2ca04b8500000000),
		182: uint64(0x359e980e00000000),
		183: uint64(0xfd8906c100000000),
		184: uint64(0x43450a0800000000),
		185: uint64(0x8b5294c700000000),
		186: uint64(0x926c474c00000000),
		187: uint64(0x5a7bd98300000000),
		188: uint64(0xe116908000000000),
		189: uint64(0x29010e4f00000000),
		190: uint64(0x303fddc400000000),
		191: uint64(0xf828430b00000000),
		192: uint64(0xf63fd9f600000000),
		193: uint64(0x3e28473900000000),
		194: uint64(0x271694b200000000),
		195: uint64(0xef010a7d00000000),
		196: uint64(0x546c437e00000000),
		197: uint64(0x9c7bddb100000000),
		198: uint64(0x85450e3a00000000),
		199: uint64(0x4d5290f500000000),
		200: uint64(0xf39e9c3c00000000),
		201: uint64(0x3b8902f300000000),
		202: uint64(0x22b7d17800000000),
		203: uint64(0xeaa04fb700000000),
		204: uint64(0x51cd06b400000000),
		205: uint64(0x99da987b00000000),
		206: uint64(0x80e44bf000000000),
		207: uint64(0x48f3d53f00000000),
		208: uint64(0xbd7b23b900000000),
		209: uint64(0x756cbd7600000000),
		210: uint64(0x6c526efd00000000),
		211: uint64(0xa445f03200000000),
		212: uint64(0x1f28b93100000000),
		213: uint64(0xd73f27fe00000000),
		214: uint64(0xce01f47500000000),
		215: uint64(0x06166aba00000000),
		216: uint64(0xb8da667300000000),
		217: uint64(0x70cdf8bc00000000),
		218: uint64(0x69f32b3700000000),
		219: uint64(0xa1e4b5f800000000),
		220: uint64(0x1a89fcfb00000000),
		221: uint64(0xd29e623400000000),
		222: uint64(0xcba0b1bf00000000),
		223: uint64(0x03b72f7000000000),
		224: uint64(0x60b72d6900000000),
		225: uint64(0xa8a0b3a600000000),
		226: uint64(0xb19e602d00000000),
		227: uint64(0x7989fee200000000),
		228: uint64(0xc2e4b7e100000000),
		229: uint64(0x0af3292e00000000),
		230: uint64(0x13cdfaa500000000),
		231: uint64(0xdbda646a00000000),
		232: uint64(0x651668a300000000),
		233: uint64(0xad01f66c00000000),
		234: uint64(0xb43f25e700000000),
		235: uint64(0x7c28bb2800000000),
		236: uint64(0xc745f22b00000000),
		237: uint64(0x0f526ce400000000),
		238: uint64(0x166cbf6f00000000),
		239: uint64(0xde7b21a000000000),
		240: uint64(0x2bf3d72600000000),
		241: uint64(0xe3e449e900000000),
		242: uint64(0xfada9a6200000000),
		243: uint64(0x32cd04ad00000000),
		244: uint64(0x89a04dae00000000),
		245: uint64(0x41b7d36100000000),
		246: uint64(0x588900ea00000000),
		247: uint64(0x909e9e2500000000),
		248: uint64(0x2e5292ec00000000),
		249: uint64(0xe6450c2300000000),
		250: uint64(0xff7bdfa800000000),
		251: uint64(0x376c416700000000),
		252: uint64(0x8c01086400000000),
		253: uint64(0x441696ab00000000),
		254: uint64(0x5d28452000000000),
		255: uint64(0x953fdbef00000000),
	},
	3: {
		1:   uint64(0x95d4709500000000),
		2:   uint64(0x6baf90f100000000),
		3:   uint64(0xfe7be06400000000),
		4:   uint64(0x9758503800000000),
		5:   uint64(0x028c20ad00000000),
		6:   uint64(0xfcf7c0c900000000),
		7:   uint64(0x6923b05c00000000),
		8:   uint64(0x2eb1a07000000000),
		9:   uint64(0xbb65d0e500000000),
		10:  uint64(0x451e308100000000),
		11:  uint64(0xd0ca401400000000),
		12:  uint64(0xb9e9f04800000000),
		13:  uint64(0x2c3d80dd00000000),
		14:  uint64(0xd24660b900000000),
		15:  uint64(0x4792102c00000000),
		16:  uint64(0x5c6241e100000000),
		17:  uint64(0xc9b6317400000000),
		18:  uint64(0x37cdd11000000000),
		19:  uint64(0xa219a18500000000),
		20:  uint64(0xcb3a11d900000000),
		21:  uint64(0x5eee614c00000000),
		22:  uint64(0xa095812800000000),
		23:  uint64(0x3541f1bd00000000),
		24:  uint64(0x72d3e19100000000),
		25:  uint64(0xe707910400000000),
		26:  uint64(0x197c716000000000),
		27:  uint64(0x8ca801f500000000),
		28:  uint64(0xe58bb1a900000000),
		29:  uint64(0x705fc13c00000000),
		30:  uint64(0x8e24215800000000),
		31:  uint64(0x1bf051cd00000000),
		32:  uint64(0xf9c2f31900000000),
		33:  uint64(0x6c16838c00000000),
		34:  uint64(0x926d63e800000000),
		35:  uint64(0x07b9137d00000000),
		36:  uint64(0x6e9aa32100000000),
		37:  uint64(0xfb4ed3b400000000),
		38:  uint64(0x053533d000000000),
		39:  uint64(0x90e1434500000000),
		40:  uint64(0xd773536900000000),
		41:  uint64(0x42a723fc00000000),
		42:  uint64(0xbcdcc39800000000),
		43:  uint64(0x2908b30d00000000),
		44:  uint64(0x402b035100000000),
		45:  uint64(0xd5ff73c400000000),
		46:  uint64(0x2b8493a000000000),
		47:  uint64(0xbe50e33500000000),
		48:  uint64(0xa5a0b2f800000000),
		49:  uint64(0x3074c26d00000000),
		50:  uint64(0xce0f220900000000),
		51:  uint64(0x5bdb529c00000000),
		52:  uint64(0x32f8e2c000000000),
		53:  uint64(0xa72c925500000000),
		54:  uint64(0x5957723100000000),
		55:  uint64(0xcc8302a400000000),
		56:  uint64(0x8b11128800000000),
		57:  uint64(0x1ec5621d00000000),
		58:  uint64(0xe0be827900000000),
		59:  uint64(0x756af2ec00000000),
		60:  uint64(0x1c4942b000000000),
		61:  uint64(0x899d322500000000),
		62:  uint64(0x77e6d24100000000),
		63:  uint64(0xe232a2d400000000),
		64:  uint64(0xf285e73300000000),
		65:  uint64(0x675197a600000000),
		66:  uint64(0x992a77c200000000),
		67:  uint64(0x0cfe075700000000),
		68:  uint64(0x65ddb70b00000000),
		69:  uint64(0xf009c79e00000000),
		70:  uint64(0x0e7227fa00000000),
		71:  uint64(0x9ba6576f00000000),
		72:  uint64(0xdc34474300000000),
		73:  uint64(0x49e037d600000000),
		74:  uint64(0xb79bd7b200000000),
		75:  uint64(0x224fa72700000000),
		76:  uint64(0x4b6c177b00000000),
		77:  uint64(0xdeb867ee00000000),
		78:  uint64(0x20c3878a00000000),
		79:  uint64(0xb517f71f00000000),
		80:  uint64(0xaee7a6d200000000),
		81:  uint64(0x3b33d64700000000),
		82:  uint64(0xc548362300000000),
		83:  uint64(0x509c46b600000000),
		84:  uint64(0x39bff6ea00000000),
		85:  uint64(0xac6b867f00000000),
		86:  uint64(0x5210661b00000000),
		87:  uint64(0xc7c4168e00000000),
		88:  uint64(0x805606a200000000),
		89:  uint64(0x1582763700000000),
		90:  uint64(0xebf9965300000000),
		91:  uint64(0x7e2de6c600000000),
		92:  uint64(0x170e569a00000000),
		93:  uint64(0x82da260f00000000),
		94:  uint64(0x7ca1c66b00000000),
		95:  uint64(0xe975b6fe00000000),
		96:  uint64(0x0b47142a00000000),
		97:  uint64(0x9e9364bf00000000),
		98:  uint64(0x60e884db00000000),
		99:  uint64(0xf53cf44e00000000),
		100: uint64(0x9c1f441200000000),
		101: uint64(0x09cb348700000000),
		102: uint64(0xf7b0d4e300000000),
		103: uint64(0x6264a47600000000),
		104: uint64(0x25f6b45a00000000),
		105: uint64(0xb022c4cf00000000),
		106: uint64(0x4e5924ab00000000),
		107: uint64(0xdb8d543e00000000),
		108: uint64(0xb2aee46200000000),
		109: uint64(0x277a94f700000000),
		110: uint64(0xd901749300000000),
		111: uint64(0x4cd5040600000000),
		112: uint64(0x572555cb00000000),
		113: uint64(0xc2f1255e00000000),
		114: uint64(0x3c8ac53a00000000),
		115: uint64(0xa95eb5af00000000),
		116: uint64(0xc07d05f300000000),
		117: uint64(0x55a9756600000000),
		118: uint64(0xabd2950200000000),
		119: uint64(0x3e06e59700000000),
		120: uint64(0x7994f5bb00000000),
		121: uint64(0xec40852e00000000),
		122: uint64(0x123b654a00000000),
		123: uint64(0x87ef15df00000000),
		124: uint64(0xeecca58300000000),
		125: uint64(0x7b18d51600000000),
		126: uint64(0x8563357200000000),
		127: uint64(0x10b745e700000000),
		128: uint64(0xe40bcf6700000000),
		129: uint64(0x71dfbff200000000),
		130: uint64(0x8fa45f9600000000),
		131: uint64(0x1a702f0300000000),
		132: uint64(0x73539f5f00000000),
		133: uint64(0xe687efca00000000),
		134: uint64(0x18fc0fae00000000),
		135: uint64(0x8d287f3b00000000),
		136: uint64(0xcaba6f1700000000),
		137: uint64(0x5f6e1f8200000000),
		138: uint64(0xa115ffe600000000),
		139: uint64(0x34c18f7300000000),
		140: uint64(0x5de23f2f00000000),
		141: uint64(0xc8364fba00000000),
		142: uint64(0x364dafde00000000),
		143: uint64(0xa399df4b00000000),
		144: uint64(0xb8698e8600000000),
		145: uint64(0x2dbdfe1300000000),
		146: uint64(0xd3c61e7700000000),
		147: uint64(0x46126ee200000000),
		148: uint64(0x2f31debe00000000),
		149: uint64(0xbae5ae2b00000000),
		150: uint64(0x449e4e4f00000000),
		151: uint64(0xd14a3eda00000000),
		152: uint64(0x96d82ef600000000),
		153: uint64(0x030c5e6300000000),
		154: uint64(0xfd77be0700000000),
		155: uint64(0x68a3ce9200000000),
		156: uint64(0x01807ece00000000),
		157: uint64(0x94540e5b00000000),
		158: uint64(0x6a2fee3f00000000),
		159: uint64(0xfffb9eaa00000000),
		160: uint64(0x1dc93c7e00000000),
		161: uint64(0x881d4ceb00000000),
		162: uint64(0x7666ac8f00000000),
		163: uint64(0xe3b2dc1a00000000),
		164: uint64(0x8a916c4600000000),
		165: uint64(0x1f451cd300000000),
		166: uint64(0xe13efcb700000000),
		167: uint64(0x74ea8c2200000000),
		168: uint64(0x33789c0e00000000),
		169: uint64(0xa6acec9b00000000),
		170: uint64(0x58d70cff00000000),
		171: uint64(0xcd037c6a00000000),
		172: uint64(0xa420cc3600000000),
		173: uint64(0x31f4bca300000000),
		174: uint64(0xcf8f5cc700000000),
		175: uint64(0x5a5b2c5200000000),
		176: uint64(0x41ab7d9f00000000),
		177: uint64(0xd47f0d0a00000000),
		178: uint64(0x2a04ed6e00000000),
		179: uint64(0xbfd09dfb00000000),
		180: uint64(0xd6f32da700000000),
		181: uint64(0x43275d3200000000),
		182: uint64(0xbd5cbd5600000000),
		183: uint64(0x2888cdc300000000),
		184: uint64(0x6f1addef00000000),
		185: uint64(0xfacead7a00000000),
		186: uint64(0x04b54d1e00000000),
		187: uint64(0x91613d8b00000000),
		188: uint64(0xf8428dd700000000),
		189: uint64(0x6d96fd4200000000),
		190: uint64(0x93ed1d2600000000),
		191: uint64(0x06396db300000000),
		192: uint64(0x168e285400000000),
		193: uint64(0x835a58c100000000),
		194: uint64(0x7d21b8a500000000),
		195: uint64(0xe8f5c83000000000),
		196: uint64(0x81d6786c00000000),
		197: uint64(0x140208f900000000),
		198: uint64(0xea79e89d00000000),
		199: uint64(0x7fad980800000000),
		200: uint64(0x383f882400000000),
		201: uint64(0xadebf8b100000000),
		202: uint64(0x539018d500000000),
		203: uint64(0xc644684000000000),
		204: uint64(0xaf67d81c00000000),
		205: uint64(0x3ab3a88900000000),
		206: uint64(0xc4c848ed00000000),
		207: uint64(0x511c387800000000),
		208: uint64(0x4aec69b500000000),
		209: uint64(0xdf38192000000000),
		210: uint64(0x2143f94400000000),
		211: uint64(0xb49789d100000000),
		212: uint64(0xddb4398d00000000),
		213: uint64(0x4860491800000000),
		214: uint64(0xb61ba97c00000000),
		215: uint64(0x23cfd9e900000000),
		216: uint64(0x645dc9c500000000),
		217: uint64(0xf189b95000000000),
		218: uint64(0x0ff2593400000000),
		219: uint64(0x9a2629a100000000),
		220: uint64(0xf30599fd00000000),
		221: uint64(0x66d1e96800000000),
		222: uint64(0x98aa090c00000000),
		223: uint64(0x0d7e799900000000),
		224: uint64(0xef4cdb4d00000000),
		225: uint64(0x7a98abd800000000),
		226: uint64(0x84e34bbc00000000),
		227: uint64(0x11373b2900000000),
		228: uint64(0x78148b7500000000),
		229: uint64(0xedc0fbe000000000),
		230: uint64(0x13bb1b8400000000),
		231: uint64(0x866f6b1100000000),
		232: uint64(0xc1fd7b3d00000000),
		233: uint64(0x54290ba800000000),
		234: uint64(0xaa52ebcc00000000),
		235: uint64(0x3f869b5900000000),
		236: uint64(0x56a52b0500000000),
		237: uint64(0xc3715b9000000000),
		238: uint64(0x3d0abbf400000000),
		239: uint64(0xa8decb6100000000),
		240: uint64(0xb32e9aac00000000),
		241: uint64(0x26faea3900000000),
		242: uint64(0xd8810a5d00000000),
		243: uint64(0x4d557ac800000000),
		244: uint64(0x2476ca9400000000),
		245: uint64(0xb1a2ba0100000000),
		246: uint64(0x4fd95a6500000000),
		247: uint64(0xda0d2af000000000),
		248: uint64(0x9d9f3adc00000000),
		249: uint64(0x084b4a4900000000),
		250: uint64(0xf630aa2d00000000),
		251: uint64(0x63e4dab800000000),
		252: uint64(0x0ac76ae400000000),
		253: uint64(0x9f131a7100000000),
		254: uint64(0x6168fa1500000000),
		255: uint64(0xf4bc8a8000000000),
	},
	4: {
		1:   uint64(0x1f17f08000000000),
		2:   uint64(0x7f2891da00000000),
		3:   uint64(0x603f615a00000000),
		4:   uint64(0xbf56536e00000000),
		5:   uint64(0xa041a3ee00000000),
		6:   uint64(0xc07ec2b400000000),
		7:   uint64(0xdf69323400000000),
		8:   uint64(0x7eada6dc00000000),
		9:   uint64(0x61ba565c00000000),
		10:  uint64(0x0185370600000000),
		11:  uint64(0x1e92c78600000000),
		12:  uint64(0xc1fbf5b200000000),
		13:  uint64(0xdeec053200000000),
		14:  uint64(0xbed3646800000000),
		15:  uint64(0xa1c494e800000000),
		16:  uint64(0xbd5c3c6200000000),
		17:  uint64(0xa24bcce200000000),
		18:  uint64(0xc274adb800000000),
		19:  uint64(0xdd635d3800000000),
		20:  uint64(0x020a6f0c00000000),
		21:  uint64(0x1d1d9f8c00000000),
		22:  uint64(0x7d22fed600000000),
		23:  uint64(0x62350e5600000000),
		24:  uint64(0xc3f19abe00000000),
		25:  uint64(0xdce66a3e00000000),
		26:  uint64(0xbcd90b6400000000),
		27:  uint64(0xa3cefbe400000000),
		28:  uint64(0x7ca7c9d000000000),
		29:  uint64(0x63b0395000000000),
		30:  uint64(0x038f580a00000000),
		31:  uint64(0x1c98a88a00000000),
		32:  uint64(0x7ab978c400000000),
		33:  uint64(0x65ae884400000000),
		34:  uint64(0x0591e91e00000000),
		35:  uint64(0x1a86199e00000000),
		36:  uint64(0xc5ef2baa00000000),
		37:  uint64(0xdaf8db2a00000000),
		38:  uint64(0xbac7ba7000000000),
		39:  uint64(0xa5d04af000000000),
		40:  uint64(0x0414de1800000000),
		41:  uint64(0x1b032e9800000000),
		42:  uint64(0x7b3c4fc200000000),
		43:  uint64(0x642bbf4200000000),
		44:  uint64(0xbb428d7600000000),
		45:  uint64(0xa4557df600000000),
		46:  uint64(0xc46a1cac00000000),
		47:  uint64(0xdb7dec2c00000000),
		48:  uint64(0xc7e544a600000000),
		49:  uint64(0xd8f2b42600000000),
		50:  uint64(0xb8cdd57c00000000),
		51:  uint64(0xa7da25fc00000000),
		52:  uint64(0x78b317c800000000),
		53:  uint64(0x67a4e74800000000),
		54:  uint64(0x079b861200000000),
		55:  uint64(0x188c769200000000),
		56:  uint64(0xb948e27a00000000),
		57:  uint64(0xa65f12fa00000000),
		58:  uint64(0xc66073a000000000),
		59:  uint64(0xd977832000000000),
		60:  uint64(0x061eb11400000000),
		61:  uint64(0x1909419400000000),
		62:  uint64(0x793620ce00000000),
		63:  uint64(0x6621d04e00000000),
		64:  uint64(0xb574805300000000),
		65:  uint64(0xaa6370d300000000),
		66:  uint64(0xca5c118900000000),
		67:  uint64(0xd54be10900000000),
		68:  uint64(0x0a22d33d00000000),
		69:  uint64(0x153523bd00000000),
		70:  uint64(0x750a42e700000000),
		71:  uint64(0x6a1db26700000000),
		72:  uint64(0xcbd9268f00000000),
		73:  uint64(0xd4ced60f00000000),
		74:  uint64(0xb4f1b75500000000),
		75:  uint64(0xabe647d500000000),
		76:  uint64(0x748f75e100000000),
		77:  uint64(0x6b98856100000000),
		78:  uint64(0x0ba7e43b00000000),
		79:  uint64(0x14b014bb00000000),
		80:  uint64(0x0828bc3100000000),
		81:  uint64(0x173f4cb100000000),
		82:  uint64(0x77002deb00000000),
		83:  uint64(0x6817dd6b00000000),
		84:  uint64(0xb77eef5f00000000),
		85:  uint64(0xa8691fdf00000000),
		86:  uint64(0xc8567e8500000000),
		87:  uint64(0xd7418e0500000000),
		88:  uint64(0x76851aed00000000),
		89:  uint64(0x6992ea6d00000000),
		90:  uint64(0x09ad8b3700000000),
		91:  uint64(0x16ba7bb700000000),
		92:  uint64(0xc9d3498300000000),
		93:  uint64(0xd6c4b90300000000),
		94:  uint64(0xb6fbd85900000000),
		95:  uint64(0xa9ec28d900000000),
		96:  uint64(0xcfcdf89700000000),
		97:  uint64(0xd0da081700000000),
		98:  uint64(0xb0e5694d00000000),
		99:  uint64(0xaff299cd00000000),
		100: uint64(0x709babf900000000),
		101: uint64(0x6f8c5b7900000000),
		102: uint64(0x0fb33a2300000000),
		103: uint64(0x10a4caa300000000),
		104: uint64(0xb1605e4b00000000),
		105: uint64(0xae77aecb00000000),
		106: uint64(0xce48cf9100000000),
		107: uint64(0xd15f3f1100000000),
		108: uint64(0x0e360d2500000000),
		109: uint64(0x1121fda500000000),
		110: uint64(0x711e9cff00000000),
		111: uint64(0x6e096c7f00000000),
		112: uint64(0x7291c4f500000000),
		113: uint64(0x6d86347500000000),
		114: uint64(0x0db9552f00000000),
		115: uint64(0x12aea5af00000000),
		116: uint64(0xcdc7979b00000000),
		117: uint64(0xd2d0671b00000000),
		118: uint64(0xb2ef064100000000),
		119: uint64(0xadf8f6c100000000),
		120: uint64(0x0c3c622900000000),
		121: uint64(0x132b92a900000000),
		122: uint64(0x7314f3f300000000),
		123: uint64(0x6c03037300000000),
		124: uint64(0xb36a314700000000),
		125: uint64(0xac7dc1c700000000),
		126: uint64(0xcc42a09d00000000),
		127: uint64(0xd355501d00000000),
		128: uint64(0x6ae900a700000000),
		129: uint64(0x75fef02700000000),
		130: uint64(0x15c1917d00000000),
		131: uint64(0x0ad661fd00000000),
		132: uint64(0xd5bf53c900000000),
		133: uint64(0xcaa8a34900000000),
		134: uint64(0xaa97c21300000000),
		135: uint64(0xb580329300000000),
		136: uint64(0x1444a67b00000000),
		137: uint64(0x0b5356fb00000000),
		138: uint64(0x6b6c37a100000000),
		139: uint64(0x747bc72100000000),
		140: uint64(0xab12f51500000000),
		141: uint64(0xb405059500000000),
		142: uint64(0xd43a64cf00000000),
		143: uint64(0xcb2d944f00000000),
		144: uint64(0xd7b53cc500000000),
		145: uint64(0xc8a2cc4500000000),
		146: uint64(0xa89dad1f00000000),
		147: uint64(0xb78a5d9f00000000),
		148: uint64(0x68e36fab00000000),
		149: uint64(0x77f49f2b00000000),
		150: uint64(0x17cbfe7100000000),
		151: uint64(0x08dc0ef100000000),
		152: uint64(0xa9189a1900000000),
		153: uint64(0xb60f6a9900000000),
		154: uint64(0xd6300bc300000000),
		155: uint64(0xc927fb4300000000),
		156: uint64(0x164ec97700000000),
		157: uint64(0x095939f700000000),
		158: uint64(0x696658ad00000000),
		159: uint64(0x7671a82d00000000),
		160: uint64(0x1050786300000000),
		161: uint64(0x0f4788e300000000),
		162: uint64(0x6f78e9b900000000),
		163: uint64(0x706f193900000000),
		164: uint64(0xaf062b0d00000000),
		165: uint64(0xb011db8d00000000),
		166: uint64(0xd02ebad700000000),
		167: uint64(0xcf394a5700000000),
		168: uint64(0x6efddebf00000000),
		169: uint64(0x71ea2e3f00000000),
		170: uint64(0x11d54f6500000000),
		171: uint64(0x0ec2bfe500000000),
		172: uint64(0xd1ab8dd100000000),
		173: uint64(0xcebc7d5100000000),
		174: uint64(0xae831c0b00000000),
		175: uint64(0xb194ec8b00000000),
		176: uint64(0xad0c440100000000),
		177: uint64(0xb21bb48100000000),
		178: uint64(0xd224d5db00000000),
		179: uint64(0xcd33255b00000000),
		180: uint64(0x125a176f00000000),
		181: uint64(0x0d4de7ef00000000),
		182: uint64(0x6d7286b500000000),
		183: uint64(0x7265763500000000),
		184: uint64(0xd3a1e2dd00000000),
		185: uint64(0xccb6125d00000000),
		186: uint64(0xac89730700000000),
		187: uint64(0xb39e838700000000),
		188: uint64(0x6cf7b1b300000000),
		189: uint64(0x73e0413300000000),
		190: uint64(0x13df206900000000),
		191: uint64(0x0cc8d0e900000000),
		192: uint64(0xdf9d80f400000000),
		193: uint64(0xc08a707400000000),
		194: uint64(0xa0b5112e00000000),
		195: uint64(0xbfa2e1ae00000000),
		196: uint64(0x60cbd39a00000000),
		197: uint64(0x7fdc231a00000000),
		198: uint64(0x1fe3424000000000),
		199: uint64(0x00f4b2c000000000),
		200: uint64(0xa130262800000000),
		201: uint64(0xbe27d6a800000000),
		202: uint64(0xde18b7f200000000),
		203: uint64(0xc10f477200000000),
		204: uint64(0x1e66754600000000),
		205: uint64(0x017185c600000000),
		206: uint64(0x614ee49c00000000),
		207: uint64(0x7e59141c00000000),
		208: uint64(0x62c1bc9600000000),
		209: uint64(0x7dd64c1600000000),
		210: uint64(0x1de92d4c00000000),
		211: uint64(0x02feddcc00000000),
		212: uint64(0xdd97eff800000000),
		213: uint64(0xc2801f7800000000),
		214: uint64(0xa2bf7e2200000000),
		215: uint64(0xbda88ea200000000),
		216: uint64(0x1c6c1a4a00000000),
		217: uint64(0x037beaca00000000),
		218: uint64(0x63448b9000000000),
		219: uint64(0x7c537b1000000000),
		220: uint64(0xa33a492400000000),
		221: uint64(0xbc2db9a400000000),
		222: uint64(0xdc12d8fe00000000),
		223: uint64(0xc305287e00000000),
		224: uint64(0xa524f83000000000),
		225: uint64(0xba3308b000000000),
		226: uint64(0xda0c69ea00000000),
		227: uint64(0xc51b996a00000000),
		228: uint64(0x1a72ab5e00000000),
		229: uint64(0x05655bde00000000),
		230: uint64(0x655a3a8400000000),
		231: uint64(0x7a4dca0400000000),
		232: uint64(0xdb895eec00000000),
		233: uint64(0xc49eae6c00000000),
		234: uint64(0xa4a1cf3600000000),
		235: uint64(0xbbb63fb600000000),
		236: uint64(0x64df0d8200000000),
		237: uint64(0x7bc8fd0200000000),
		238: uint64(0x1bf79c5800000000),
		239: uint64(0x04e06cd800000000),
		240: uint64(0x1878c45200000000),
		241: uint64(0x076f34d200000000),
		242: uint64(0x6750558800000000),
		243: uint64(0x7847a50800000000),
		244: uint64(0xa72e973c00000000),
		245: uint64(0xb83967bc00000000),
		246: uint64(0xd80606e600000000),
		247: uint64(0xc711f66600000000),
		248: uint64(0x66d5628e00000000),
		249: uint64(0x79c2920e00000000),
		250: uint64(0x19fdf35400000000),
		251: uint64(0x06ea03d400000000),
		252: uint64(0xd98331e000000000),
		253: uint64(0xc694c16000000000),
		254: uint64(0xa6aba03a00000000),
		255: uint64(0xb9bc50ba00000000),
	},
	5: {
		1:   uint64(0xe2fd888d00000000),
		2:   uint64(0x85fd60c000000000),
		3:   uint64(0x6700e84d00000000),
		4:   uint64(0x4bfdb05b00000000),
		5:   uint64(0xa90038d600000000),
		6:   uint64(0xce00d09b00000000),
		7:   uint64(0x2cfd581600000000),
		8:   uint64(0x96fa61b700000000),
		9:   uint64(0x7407e93a00000000),
		10:  uint64(0x1307017700000000),
		11:  uint64(0xf1fa89fa00000000),
		12:  uint64(0xdd07d1ec00000000),
		13:  uint64(0x3ffa596100000000),
		14:  uint64(0x58fab12c00000000),
		15:  uint64(0xba0739a100000000),
		16:  uint64(0x6df3b2b500000000),
		17:  uint64(0x8f0e3a3800000000),
		18:  uint64(0xe80ed27500000000),
		19:  uint64(0x0af35af800000000),
		20:  uint64(0x260e02ee00000000),
		21:  uint64(0xc4f38a6300000000),
		22:  uint64(0xa3f3622e00000000),
		23:  uint64(0x410eeaa300000000),
		24:  uint64(0xfb09d30200000000),
		25:  uint64(0x19f45b8f00000000),
		26:  uint64(0x7ef4b3c200000000),
		27:  uint64(0x9c093b4f00000000),
		28:  uint64(0xb0f4635900000000),
		29:  uint64(0x5209ebd400000000),
		30:  uint64(0x3509039900000000),
		31:  uint64(0xd7f48b1400000000),
		32:  uint64(0x9be014b000000000),
		33:  uint64(0x791d9c3d00000000),
		34:  uint64(0x1e1d747000000000),
		35:  uint64(0xfce0fcfd00000000),
		36:  uint64(0xd01da4eb00000000),
		37:  uint64(0x32e02c6600000000),
		38:  uint64(0x55e0c42b00000000),
		39:  uint64(0xb71d4ca600000000),
		40:  uint64(0x0d1a750700000000),
		41:  uint64(0xefe7fd8a00000000),
		42:  uint64(0x88e715c700000000),
		43:  uint64(0x6a1a9d4a00000000),
		44:  uint64(0x46e7c55c00000000),
		45:  uint64(0xa41a4dd100000000),
		46:  uint64(0xc31aa59c00000000),
		47:  uint64(0x21e72d1100000000),
		48:  uint64(0xf613a60500000000),
		49:  uint64(0x14ee2e8800000000),
		50:  uint64(0x73eec6c500000000),
		51:  uint64(0x91134e4800000000),
		52:  uint64(0xbdee165e00000000),
		53:  uint64(0x5f139ed300000000),
		54:  uint64(0x3813769e00000000),
		55:  uint64(0xdaeefe1300000000),
		56:  uint64(0x60e9c7b200000000),
		57:  uint64(0x82144f3f00000000),
		58:  uint64(0xe514a77200000000),
		59:  uint64(0x07e92fff00000000),
		60:  uint64(0x2b1477e900000000),
		61:  uint64(0xc9e9ff6400000000),
		62:  uint64(0xaee9172900000000),
		63:  uint64(0x4c149fa400000000),
		64:  uint64(0x77c758bb00000000),
		65:  uint64(0x953ad03600000000),
		66:  uint64(0xf23a387b00000000),
		67:  uint64(0x10c7b0f600000000),
		68:  uint64(0x3c3ae8e000000000),
		69:  uint64(0xdec7606d00000000),
		70:  uint64(0xb9c7882000000000),
		71:  uint64(0x5b3a00ad00000000),
		72:  uint64(0xe13d390c00000000),
		73:  uint64(0x03c0b18100000000),
		74:  uint64(0x64c059cc00000000),
		75:  uint64(0x863dd14100000000),
		76:  uint64(0xaac0895700000000),
		77:  uint64(0x483d01da00000000),
		78:  uint64(0x2f3de99700000000),
		79:  uint64(0xcdc0611a00000000),
		80:  uint64(0x1a34ea0e00000000),
		81:  uint64(0xf8c9628300000000),
		82:  uint64(0x9fc98ace00000000),
		83:  uint64(0x7d34024300000000),
		84:  uint64(0x51c95a5500000000),
		85:  uint64(0xb334d2d800000000),
		86:  uint64(0xd4343a9500000000),
		87:  uint64(0x36c9b21800000000),
		88:  uint64(0x8cce8bb900000000),
		89:  uint64(0x6e33033400000000),
		90:  uint64(0x0933eb7900000000),
		91:  uint64(0xebce63f400000000),
		92:  uint64(0xc7333be200000000),
		93:  uint64(0x25ceb36f00000000),
		94:  uint64(0x42ce5b2200000000),
		95:  uint64(0xa033d3af00000000),
		96:  uint64(0xec274c0b00000000),
		97:  uint64(0x0edac48600000000),
		98:  uint64(0x69da2ccb00000000),
		99:  uint64(0x8b27a44600000000),
		100: uint64(0xa7dafc5000000000),
		101: uint64(0x452774dd00000000),
		102: uint64(0x22279c9000000000),
		103: uint64(0xc0da141d00000000),
		104: uint64(0x7add2dbc00000000),
		105: uint64(0x9820a53100000000),
		106: uint64(0xff204d7c00000000),
		107: uint64(0x1dddc5f100000000),
		108: uint64(0x31209de700000000),
		109: uint64(0xd3dd156a00000000),
		110: uint64(0xb4ddfd2700000000),
		111: uint64(0x562075aa00000000),
		112: uint64(0x81d4febe00000000),
		113: uint64(0x6329763300000000),
		114: uint64(0x04299e7e00000000),
		115: uint64(0xe6d416f300000000),
		116: uint64(0xca294ee500000000),
		117: uint64(0x28d4c66800000000),
		118: uint64(0x4fd42e2500000000),
		119: uint64(0xad29a6a800000000),
		120: uint64(0x172e9f0900000000),
		121: uint64(0xf5d3178400000000),
		122: uint64(0x92d3ffc900000000),
		123: uint64(0x702e774400000000),
		124: uint64(0x5cd32f5200000000),
		125: uint64(0xbe2ea7df00000000),
		126: uint64(0xd92e4f9200000000),
		127: uint64(0x3bd3c71f00000000),
		128: uint64(0xaf88c0ad00000000),
		129: uint64(0x4d75482000000000),
		130: uint64(0x2a75a06d00000000),
		131: uint64(0xc88828e000000000),
		132: uint64(0xe47570f600000000),
		133: uint64(0x0688f87b00000000),
		134: uint64(0x6188103600000000),
		135: uint64(0x837598bb00000000),
		136: uint64(0x3972a11a00000000),
		137: uint64(0xdb8f299700000000),
		138: uint64(0xbc8fc1da00000000),
		139: uint64(0x5e72495700000000),
		140: uint64(0x728f114100000000),
		141: uint64(0x907299cc00000000),
		142: uint64(0xf772718100000000),
		143: uint64(0x158ff90c00000000),
		144: uint64(0xc27b721800000000),
		145: uint64(0x2086fa9500000000),
		146: uint64(0x478612d800000000),
		147: uint64(0xa57b9a5500000000),
		148: uint64(0x8986c24300000000),
		149: uint64(0x6b7b4ace00000000),
		150: uint64(0x0c7ba28300000000),
		151: uint64(0xee862a0e00000000),
		152: uint64(0x548113af00000000),
		153: uint64(0xb67c9b2200000000),
		154: uint64(0xd17c736f00000000),
		155: uint64(0x3381fbe200000000),
		156: uint64(0x1f7ca3f400000000),
		157: uint64(0xfd812b7900000000),
		158: uint64(0x9a81c33400000000),
		159: uint64(0x787c4bb900000000),
		160: uint64(0x3468d41d00000000),
		161: uint64(0xd6955c9000000000),
		162: uint64(0xb195b4dd00000000),
		163: uint64(0x53683c5000000000),
		164: uint64(0x7f95644600000000),
		165: uint64(0x9d68eccb00000000),
		166: uint64(0xfa68048600000000),
		167: uint64(0x18958c0b00000000),
		168: uint64(0xa292b5aa00000000),
		169: uint64(0x406f3d2700000000),
		170: uint64(0x276fd56a00000000),
		171: uint64(0xc5925de700000000),
		172: uint64(0xe96f05f100000000),
		173: uint64(0x0b928d7c00000000),
		174: uint64(0x6c92653100000000),
		175: uint64(0x8e6fedbc00000000),
		176: uint64(0x599b66a800000000),
		177: uint64(0xbb66ee2500000000),
		178: uint64(0xdc66066800000000),
		179: uint64(0x3e9b8ee500000000),
		180: uint64(0x1266d6f300000000),
		181: uint64(0xf09b5e7e00000000),
		182: uint64(0x979bb63300000000),
		183: uint64(0x75663ebe00000000),
		184: uint64(0xcf61071f00000000),
		185: uint64(0x2d9c8f9200000000),
		186: uint64(0x4a9c67df00000000),
		187: uint64(0xa861ef5200000000),
		188: uint64(0x849cb74400000000),
		189: uint64(0x66613fc900000000),
		190: uint64(0x0161d78400000000),
		191: uint64(0xe39c5f0900000000),
		192: uint64(0xd84f981600000000),
		193: uint64(0x3ab2109b00000000),
		194: uint64(0x5db2f8d600000000),
		195: uint64(0xbf4f705b00000000),
		196: uint64(0x93b2284d00000000),
		197: uint64(0x714fa0c000000000),
		198: uint64(0x164f488d00000000),
		199: uint64(0xf4b2c00000000000),
		200: uint64(0x4eb5f9a100000000),
		201: uint64(0xac48712c00000000),
		202: uint64(0xcb48996100000000),
		203: uint64(0x29b511ec00000000),
		204: uint64(0x054849fa00000000),
		205: uint64(0xe7b5c17700000000),
		206: uint64(0x80b5293a00000000),
		207: uint64(0x6248a1b700000000),
		208: uint64(0xb5bc2aa300000000),
		209: uint64(0x5741a22e00000000),
		210: uint64(0x30414a6300000000),
		211: uint64(0xd2bcc2ee00000000),
		212: uint64(0xfe419af800000000),
		213: uint64(0x1cbc127500000000),
		214: uint64(0x7bbcfa3800000000),
		215: uint64(0x994172b500000000),
		216: uint64(0x23464b1400000000),
		217: uint64(0xc1bbc39900000000),
		218: uint64(0xa6bb2bd400000000),
		219: uint64(0x4446a35900000000),
		220: uint64(0x68bbfb4f00000000),
		221: uint64(0x8a4673c200000000),
		222: uint64(0xed469b8f00000000),
		223: uint64(0x0fbb130200000000),
		224: uint64(0x43af8ca600000000),
		225: uint64(0xa152042b00000000),
		226: uint64(0xc652ec6600000000),
		227: uint64(0x24af64eb00000000),
		228: uint64(0x08523cfd00000000),
		229: uint64(0xeaafb47000000000),
		230: uint64(0x8daf5c3d00000000),
		231: uint64(0x6f52d4b000000000),
		232: uint64(0xd555ed1100000000),
		233: uint64(0x37a8659c00000000),
		234: uint64(0x50a88dd100000000),
		235: uint64(0xb255055c00000000),
		236: uint64(0x9ea85d4a00000000),
		237: uint64(0x7c55d5c700000000),
		238: uint64(0x1b553d8a00000000),
		239: uint64(0xf9a8b50700000000),
		240: uint64(0x2e5c3e1300000000),
		241: uint64(0xcca1b69e00000000),
		242: uint64(0xaba15ed300000000),
		243: uint64(0x495cd65e00000000),
		244: uint64(0x65a18e4800000000),
		245: uint64(0x875c06c500000000),
		246: uint64(0xe05cee8800000000),
		247: uint64(0x02a1660500000000),
		248: uint64(0xb8a65fa400000000),
		249: uint64(0x5a5bd72900000000),
		250: uint64(0x3d5b3f6400000000),
		251: uint64(0xdfa6b7e900000000),
		252: uint64(0xf35befff00000000),
		253: uint64(0x11a6677200000000),
		254: uint64(0x76a68f3f00000000),
		255: uint64(0x945b07b200000000),
	},
	6: {
		1:   uint64(0xa90b894e00000000),
		2:   uint64(0x5217129d00000000),
		3:   uint64(0xfb1c9bd300000000),
		4:   uint64(0xe52855e100000000),
		5:   uint64(0x4c23dcaf00000000),
		6:   uint64(0xb73f477c00000000),
		7:   uint64(0x1e34ce3200000000),
		8:   uint64(0x8b57db1900000000),
		9:   uint64(0x225c525700000000),
		10:  uint64(0xd940c98400000000),
		11:  uint64(0x704b40ca00000000),
		12:  uint64(0x6e7f8ef800000000),
		13:  uint64(0xc77407b600000000),
		14:  uint64(0x3c689c6500000000),
		15:  uint64(0x9563152b00000000),
		16:  uint64(0x16afb63300000000),
		17:  uint64(0xbfa43f7d00000000),
		18:  uint64(0x44b8a4ae00000000),
		19:  uint64(0xedb32de000000000),
		20:  uint64(0xf387e3d200000000),
		21:  uint64(0x5a8c6a9c00000000),
		22:  uint64(0xa190f14f00000000),
		23:  uint64(0x089b780100000000),
		24:  uint64(0x9df86d2a00000000),
		25:  uint64(0x34f3e46400000000),
		26:  uint64(0xcfef7fb700000000),
		27:  uint64(0x66e4f6f900000000),
		28:  uint64(0x78d038cb00000000),
		29:  uint64(0xd1dbb18500000000),
		30:  uint64(0x2ac72a5600000000),
		31:  uint64(0x83cca31800000000),
		32:  uint64(0x2c5e6d6700000000),
		33:  uint64(0x8555e42900000000),
		34:  uint64(0x7e497ffa00000000),
		35:  uint64(0xd742f6b400000000),
		36:  uint64(0xc976388600000000),
		37:  uint64(0x607db1c800000000),
		38:  uint64(0x9b612a1b00000000),
		39:  uint64(0x326aa35500000000),
		40:  uint64(0xa709b67e00000000),
		41:  uint64(0x0e023f3000000000),
		42:  uint64(0xf51ea4e300000000),
		43:  uint64(0x5c152dad00000000),
		44:  uint64(0x4221e39f00000000),
		45:  uint64(0xeb2a6ad100000000),
		46:  uint64(0x1036f10200000000),
		47:  uint64(0xb93d784c00000000),
		48:  uint64(0x3af1db5400000000),
		49:  uint64(0x93fa521a00000000),
		50:  uint64(0x68e6c9c900000000),
		51:  uint64(0xc1ed408700000000),
		52:  uint64(0xdfd98eb500000000),
		53:  uint64(0x76d207fb00000000),
		54:  uint64(0x8dce9c2800000000),
		55:  uint64(0x24c5156600000000),
		56:  uint64(0xb1a6004d00000000),
		57:  uint64(0x18ad890300000000),
		58:  uint64(0xe3b112d000000000),
		59:  uint64(0x4aba9b9e00000000),
		60:  uint64(0x548e55ac00000000),
		61:  uint64(0xfd85dce200000000),
		62:  uint64(0x0699473100000000),
		63:  uint64(0xaf92ce7f00000000),
		64:  uint64(0x58bcdace00000000),
		65:  uint64(0xf1b7538000000000),
		66:  uint64(0x0aabc85300000000),
		67:  uint64(0xa3a0411d00000000),
		68:  uint64(0xbd948f2f00000000),
		69:  uint64(0x149f066100000000),
		70:  uint64(0xef839db200000000),
		71:  uint64(0x468814fc00000000),
		72:  uint64(0xd3eb01d700000000),
		73:  uint64(0x7ae0889900000000),
		74:  uint64(0x81fc134a00000000),
		75:  uint64(0x28f79a0400000000),
		76:  uint64(0x36c3543600000000),
		77:  uint64(0x9fc8dd7800000000),
		78:  uint64(0x64d446ab00000000),
		79:  uint64(0xcddfcfe500000000),
		80:  uint64(0x4e136cfd00000000),
		81:  uint64(0xe718e5b300000000),
		82:  uint64(0x1c047e6000000000),
		83:  uint64(0xb50ff72e00000000),
		84:  uint64(0xab3b391c00000000),
		85:  uint64(0x0230b05200000000),
		86:  uint64(0xf92c2b8100000000),
		87:  uint64(0x5027a2cf00000000),
		88:  uint64(0xc544b7e400000000),
		89:  uint64(0x6c4f3eaa00000000),
		90:  uint64(0x9753a57900000000),
		91:  uint64(0x3e582c3700000000),
		92:  uint64(0x206ce20500000000),
		93:  uint64(0x89676b4b00000000),
		94:  uint64(0x727bf09800000000),
		95:  uint64(0xdb7079d600000000),
		96:  uint64(0x74e2b7a900000000),
		97:  uint64(0xdde93ee700000000),
		98:  uint64(0x26f5a53400000000),
		99:  uint64(0x8ffe2c7a00000000),
		100: uint64(0x91cae24800000000),
		101: uint64(0x38c16b0600000000),
		102: uint64(0xc3ddf0d500000000),
		103: uint64(0x6ad6799b00000000),
		104: uint64(0xffb56cb000000000),
		105: uint64(0x56bee5fe00000000),
		106: uint64(0xada27e2d00000000),
		107: uint64(0x04a9f76300000000),
		108: uint64(0x1a9d395100000000),
		109: uint64(0xb396b01f00000000),
		110: uint64(0x488a2bcc00000000),
		111: uint64(0xe181a28200000000),
		112: uint64(0x624d019a00000000),
		113: uint64(0xcb4688d400000000),
		114: uint64(0x305a130700000000),
		115: uint64(0x99519a4900000000),
		116: uint64(0x8765547b00000000),
		117: uint64(0x2e6edd3500000000),
		118: uint64(0xd57246e600000000),
		119: uint64(0x7c79cfa800000000),
		120: uint64(0xe91ada8300000000),
		121: uint64(0x401153cd00000000),
		122: uint64(0xbb0dc81e00000000),
		123: uint64(0x1206415000000000),
		124: uint64(0x0c328f6200000000),
		125: uint64(0xa539062c00000000),
		126: uint64(0x5e259dff00000000),
		127: uint64(0xf72e14b100000000),
		128: uint64(0xf17ec44600000000),
		129: uint64(0x58754d0800000000),
		130: uint64(0xa369d6db00000000),
		131: uint64(0x0a625f9500000000),
		132: uint64(0x145691a700000000),
		133: uint64(0xbd5d18e900000000),
		134: uint64(0x4641833a00000000),
		135: uint64(0xef4a0a7400000000),
		136: uint64(0x7a291f5f00000000),
		137: uint64(0xd322961100000000),
		138: uint64(0x283e0dc200000000),
		139: uint64(0x8135848c00000000),
		140: uint64(0x9f014abe00000000),
		141: uint64(0x360ac3f000000000),
		142: uint64(0xcd16582300000000),
		143: uint64(0x641dd16d00000000),
		144: uint64(0xe7d1727500000000),
		145: uint64(0x4edafb3b00000000),
		146: uint64(0xb5c660e800000000),
		147: uint64(0x1ccde9a600000000),
		148: uint64(0x02f9279400000000),
		149: uint64(0xabf2aeda00000000),
		150: uint64(0x50ee350900000000),
		151: uint64(0xf9e5bc4700000000),
		152: uint64(0x6c86a96c00000000),
		153: uint64(0xc58d202200000000),
		154: uint64(0x3e91bbf100000000),
		155: uint64(0x979a32bf00000000),
		156: uint64(0x89aefc8d00000000),
		157: uint64(0x20a575c300000000),
		158: uint64(0xdbb9ee1000000000),
		159: uint64(0x72b2675e00000000),
		160: uint64(0xdd20a92100000000),
		161: uint64(0x742b206f00000000),
		162: uint64(0x8f37bbbc00000000),
		163: uint64(0x263c32f200000000),
		164: uint64(0x3808fcc000000000),
		165: uint64(0x9103758e00000000),
		166: uint64(0x6a1fee5d00000000),
		167: uint64(0xc314671300000000),
		168: uint64(0x5677723800000000),
		169: uint64(0xff7cfb7600000000),
		170: uint64(0x046060a500000000),
		171: uint64(0xad6be9eb00000000),
		172: uint64(0xb35f27d900000000),
		173: uint64(0x1a54ae9700000000),
		174: uint64(0xe148354400000000),
		175: uint64(0x4843bc0a00000000),
		176: uint64(0xcb8f1f1200000000),
		177: uint64(0x6284965c00000000),
		178: uint64(0x99980d8f00000000),
		179: uint64(0x309384c100000000),
		180: uint64(0x2ea74af300000000),
		181: uint64(0x87acc3bd00000000),
		182: uint64(0x7cb0586e00000000),
		183: uint64(0xd5bbd12000000000),
		184: uint64(0x40d8c40b00000000),
		185: uint64(0xe9d34d4500000000),
		186: uint64(0x12cfd69600000000),
		187: uint64(0xbbc45fd800000000),
		188: uint64(0xa5f091ea00000000),
		189: uint64(0x0cfb18a400000000),
		190: uint64(0xf7e7837700000000),
		191: uint64(0x5eec0a3900000000),
		192: uint64(0xa9c21e8800000000),
		193: uint64(0x00c997c600000000),
		194: uint64(0xfbd50c1500000000),
		195: uint64(0x52de855b00000000),
		196: uint64(0x4cea4b6900000000),
		197: uint64(0xe5e1c22700000000),
		198: uint64(0x1efd59f400000000),
		199: uint64(0xb7f6d0ba00000000),
		200: uint64(0x2295c59100000000),
		201: uint64(0x8b9e4cdf00000000),
		202: uint64(0x7082d70c00000000),
		203: uint64(0xd9895e4200000000),
		204: uint64(0xc7bd907000000000),
		205: uint64(0x6eb6193e00000000),
		206: uint64(0x95aa82ed00000000),
		207: uint64(0x3ca10ba300000000),
		208: uint64(0xbf6da8bb00000000),
		209: uint64(0x166621f500000000),
		210: uint64(0xed7aba2600000000),
		211: uint64(0x4471336800000000),
		212: uint64(0x5a45fd5a00000000),
		213: uint64(0xf34e741400000000),
		214: uint64(0x0852efc700000000),
		215: uint64(0xa159668900000000),
		216: uint64(0x343a73a200000000),
		217: uint64(0x9d31faec00000000),
		218: uint64(0x662d613f00000000),
		219: uint64(0xcf26e87100000000),
		220: uint64(0xd112264300000000),
		221: uint64(0x7819af0d00000000),
		222: uint64(0x830534de00000000),
		223: uint64(0x2a0ebd9000000000),
		224: uint64(0x859c73ef00000000),
		225: uint64(0x2c97faa100000000),
		226: uint64(0xd78b617200000000),
		227: uint64(0x7e80e83c00000000),
		228: uint64(0x60b4260e00000000),
		229: uint64(0xc9bfaf4000000000),
		230: uint64(0x32a3349300000000),
		231: uint64(0x9ba8bddd00000000),
		232: uint64(0x0ecba8f600000000),
		233: uint64(0xa7c021b800000000),
		234: uint64(0x5cdcba6b00000000),
		235: uint64(0xf5d7332500000000),
		236: uint64(0xebe3fd1700000000),
		237: uint64(0x42e8745900000000),
		238: uint64(0xb9f4ef8a00000000),
		239: uint64(0x10ff66c400000000),
		240: uint64(0x9333c5dc00000000),
		241: uint64(0x3a384c9200000000),
		242: uint64(0xc124d74100000000),
		243: uint64(0x682f5e0f00000000),
		244: uint64(0x761b903d00000000),
		245: uint64(0xdf10197300000000),
		246: uint64(0x240c82a000000000),
		247: uint64(0x8d070bee00000000),
		248: uint64(0x18641ec500000000),
		249: uint64(0xb16f978b00000000),
		250: uint64(0x4a730c5800000000),
		251: uint64(0xe378851600000000),
		252: uint64(0xfd4c4b2400000000),
		253: uint64(0x5447c26a00000000),
		254: uint64(0xaf5b59b900000000),
		255: uint64(0x0650d0f700000000),
	},
	7: {
		1:   uint64(0x479244af00000000),
		2:   uint64(0xcf22f88500000000),
		3:   uint64(0x88b0bc2a00000000),
		4:   uint64(0xdf4381d000000000),
		5:   uint64(0x98d1c57f00000000),
		6:   uint64(0x1061795500000000),
		7:   uint64(0x57f33dfa00000000),
		8:   uint64(0xff81737a00000000),
		9:   uint64(0xb81337d500000000),
		10:  uint64(0x30a38bff00000000),
		11:  uint64(0x7731cf5000000000),
		12:  uint64(0x20c2f2aa00000000),
		13:  uint64(0x6750b60500000000),
		14:  uint64(0xefe00a2f00000000),
		15:  uint64(0xa8724e8000000000),
		16:  uint64(0xfe03e7f400000000),
		17:  uint64(0xb991a35b00000000),
		18:  uint64(0x31211f7100000000),
		19:  uint64(0x76b35bde00000000),
		20:  uint64(0x2140662400000000),
		21:  uint64(0x66d2228b00000000),
		22:  uint64(0xee629ea100000000),
		23:  uint64(0xa9f0da0e00000000),
		24:  uint64(0x0182948e00000000),
		25:  uint64(0x4610d02100000000),
		26:  uint64(0xcea06c0b00000000),
		27:  uint64(0x893228a400000000),
		28:  uint64(0xdec1155e00000000),
		29:  uint64(0x995351f100000000),
		30:  uint64(0x11e3eddb00000000),
		31:  uint64(0x5671a97400000000),
		32:  uint64(0xbd01bf3200000000),
		33:  uint64(0xfa93fb9d00000000),
		34:  uint64(0x722347b700000000),
		35:  uint64(0x35b1031800000000),
		36:  uint64(0x62423ee200000000),
		37:  uint64(0x25d07a4d00000000),
		38:  uint64(0xad60c66700000000),
		39:  uint64(0xeaf282c800000000),
		40:  uint64(0x4280cc4800000000),
		41:  uint64(0x051288e700000000),
		42:  uint64(0x8da234cd00000000),
		43:  uint64(0xca30706200000000),
		44:  uint64(0x9dc34d9800000000),
		45:  uint64(0xda51093700000000),
		46:  uint64(0x52e1b51d00000000),
		47:  uint64(0x1573f1b200000000),
		48:  uint64(0x430258c600000000),
		49:  uint64(0x04901c6900000000),
		50:  uint64(0x8c20a04300000000),
		51:  uint64(0xcbb2e4ec00000000),
		52:  uint64(0x9c41d91600000000),
		53:  uint64(0xdbd39db900000000),
		54:  uint64(0x5363219300000000),
		55:  uint64(0x14f1653c00000000),
		56:  uint64(0xbc832bbc00000000),
		57:  uint64(0xfb116f1300000000),
		58:  uint64(0x73a1d33900000000),
		59:  uint64(0x3433979600000000),
		60:  uint64(0x63c0aa6c00000000),
		61:  uint64(0x2452eec300000000),
		62:  uint64(0xace252e900000000),
		63:  uint64(0xeb70164600000000),
		64:  uint64(0x7a037e6500000000),
		65:  uint64(0x3d913aca00000000),
		66:  uint64(0xb52186e000000000),
		67:  uint64(0xf2b3c24f00000000),
		68:  uint64(0xa540ffb500000000),
		69:  uint64(0xe2d2bb1a00000000),
		70:  uint64(0x6a62073000000000),
		71:  uint64(0x2df0439f00000000),
		72:  uint64(0x85820d1f00000000),
		73:  uint64(0xc21049b000000000),
		74:  uint64(0x4aa0f59a00000000),
		75:  uint64(0x0d32b13500000000),
		76:  uint64(0x5ac18ccf00000000),
		77:  uint64(0x1d53c86000000000),
		78:  uint64(0x95e3744a00000000),
		79:  uint64(0xd27130e500000000),
		80:  uint64(0x8400999100000000),
		81:  uint64(0xc392dd3e00000000),
		82:  uint64(0x4b22611400000000),
		83:  uint64(0x0cb025bb00000000),
		84:  uint64(0x5b43184100000000),
		85:  uint64(0x1cd15cee00000000),
		86:  uint64(0x9461e0c400000000),
		87:  uint64(0xd3f3a46b00000000),
		88:  uint64(0x7b81eaeb00000000),
		89:  uint64(0x3c13ae4400000000),
		90:  uint64(0xb4a3126e00000000),
		91:  uint64(0xf33156c100000000),
		92:  uint64(0xa4c26b3b00000000),
		93:  uint64(0xe3502f9400000000),
		94:  uint64(0x6be093be00000000),
		95:  uint64(0x2c72d71100000000),
		96:  uint64(0xc702c15700000000),
		97:  uint64(0x809085f800000000),
		98:  uint64(0x082039d200000000),
		99:  uint64(0x4fb27d7d00000000),
		100: uint64(0x1841408700000000),
		101: uint64(0x5fd3042800000000),
		102: uint64(0xd763b80200000000),
		103: uint64(0x90f1fcad00000000),
		104: uint64(0x3883b22d00000000),
		105: uint64(0x7f11f68200000000),
		106: uint64(0xf7a14aa800000000),
		107: uint64(0xb0330e0700000000),
		108: uint64(0xe7c033fd00000000),
		109: uint64(0xa052775200000000),
		110: uint64(0x28e2cb7800000000),
		111: uint64(0x6f708fd700000000),
		112: uint64(0x390126a300000000),
		113: uint64(0x7e93620c00000000),
		114: uint64(0xf623de2600000000),
		115: uint64(0xb1b19a8900000000),
		116: uint64(0xe642a77300000000),
		117: uint64(0xa1d0e3dc00000000),
		118: uint64(0x29605ff600000000),
		119: uint64(0x6ef21b5900000000),
		120: uint64(0xc68055d900000000),
		121: uint64(0x8112117600000000),
		122: uint64(0x09a2ad5c00000000),
		123: uint64(0x4e30e9f300000000),
		124: uint64(0x19c3d40900000000),
		125: uint64(0x5e5190a600000000),
		126: uint64(0xd6e12c8c00000000),
		127: uint64(0x9173682300000000),
		128: uint64(0xf406fcca00000000),
		129: uint64(0xb394b86500000000),
		130: uint64(0x3b24044f00000000),
		131: uint64(0x7cb640e000000000),
		132: uint64(0x2b457d1a00000000),
		133: uint64(0x6cd739b500000000),
		134: uint64(0xe467859f00000000),
		135: uint64(0xa3f5c13000000000),
		136: uint64(0x0b878fb000000000),
		137: uint64(0x4c15cb1f00000000),
		138: uint64(0xc4a5773500000000),
		139: uint64(0x8337339a00000000),
		140: uint64(0xd4c40e6000000000),
		141: uint64(0x93564acf00000000),
		142: uint64(0x1be6f6e500000000),
		143: uint64(0x5c74b24a00000000),
		144: uint64(0x0a051b3e00000000),
		145: uint64(0x4d975f9100000000),
		146: uint64(0xc527e3bb00000000),
		147: uint64(0x82b5a71400000000),
		148: uint64(0xd5469aee00000000),
		149: uint64(0x92d4de4100000000),
		150: uint64(0x1a64626b00000000),
		151: uint64(0x5df626c400000000),
		152: uint64(0xf584684400000000),
		153: uint64(0xb2162ceb00000000),
		154: uint64(0x3aa690c100000000),
		155: uint64(0x7d34d46e00000000),
		156: uint64(0x2ac7e99400000000),
		157: uint64(0x6d55ad3b00000000),
		158: uint64(0xe5e5111100000000),
		159: uint64(0xa27755be00000000),
		160: uint64(0x490743f800000000),
		161: uint64(0x0e95075700000000),
		162: uint64(0x8625bb7d00000000),
		163: uint64(0xc1b7ffd200000000),
		164: uint64(0x9644c22800000000),
		165: uint64(0xd1d6868700000000),
		166: uint64(0x59663aad00000000),
		167: uint64(0x1ef47e0200000000),
		168: uint64(0xb686308200000000),
		169: uint64(0xf114742d00000000),
		170: uint64(0x79a4c80700000000),
		171: uint64(0x3e368ca800000000),
		172: uint64(0x69c5b15200000000),
		173: uint64(0x2e57f5fd00000000),
		174: uint64(0xa6e749d700000000),
		175: uint64(0xe1750d7800000000),
		176: uint64(0xb704a40c00000000),
		177: uint64(0xf096e0a300000000),
		178: uint64(0x78265c8900000000),
		179: uint64(0x3fb4182600000000),
		180: uint64(0x684725dc00000000),
		181: uint64(0x2fd5617300000000),
		182: uint64(0xa765dd5900000000),
		183: uint64(0xe0f799f600000000),
		184: uint64(0x4885d77600000000),
		185: uint64(0x0f1793d900000000),
		186: uint64(0x87a72ff300000000),
		187: uint64(0xc0356b5c00000000),
		188: uint64(0x97c656a600000000),
		189: uint64(0xd054120900000000),
		190: uint64(0x58e4ae2300000000),
		191: uint64(0x1f76ea8c00000000),
		192: uint64(0x8e0582af00000000),
		193: uint64(0xc997c60000000000),
		194: uint64(0x41277a2a00000000),
		195: uint64(0x06b53e8500000000),
		196: uint64(0x5146037f00000000),
		197: uint64(0x16d447d000000000),
		198: uint64(0x9e64fbfa00000000),
		199: uint64(0xd9f6bf5500000000),
		200: uint64(0x7184f1d500000000),
		201: uint64(0x3616b57a00000000),
		202: uint64(0xbea6095000000000),
		203: uint64(0xf9344dff00000000),
		204: uint64(0xaec7700500000000),
		205: uint64(0xe95534aa00000000),
		206: uint64(0x61e5888000000000),
		207: uint64(0x2677cc2f00000000),
		208: uint64(0x7006655b00000000),
		209: uint64(0x379421f400000000),
		210: uint64(0xbf249dde00000000),
		211: uint64(0xf8b6d97100000000),
		212: uint64(0xaf45e48b00000000),
		213: uint64(0xe8d7a02400000000),
		214: uint64(0x60671c0e00000000),
		215: uint64(0x27f558a100000000),
		216: uint64(0x8f87162100000000),
		217: uint64(0xc815528e00000000),
		218: uint64(0x40a5eea400000000),
		219: uint64(0x0737aa0b00000000),
		220: uint64(0x50c497f100000000),
		221: uint64(0x1756d35e00000000),
		222: uint64(0x9fe66f7400000000),
		223: uint64(0xd8742bdb00000000),
		224: uint64(0x33043d9d00000000),
		225: uint64(0x7496793200000000),
		226: uint64(0xfc26c51800000000),
		227: uint64(0xbbb481b700000000),
		228: uint64(0xec47bc4d00000000),
		229: uint64(0xabd5f8e200000000),
		230: uint64(0x236544c800000000),
		231: uint64(0x64f7006700000000),
		232: uint64(0xcc854ee700000000),
		233: uint64(0x8b170a4800000000),
		234: uint64(0x03a7b66200000000),
		235: uint64(0x4435f2cd00000000),
		236: uint64(0x13c6cf3700000000),
		237: uint64(0x54548b9800000000),
		238: uint64(0xdce437b200000000),
		239: uint64(0x9b76731d00000000),
		240: uint64(0xcd07da6900000000),
		241: uint64(0x8a959ec600000000),
		242: uint64(0x022522ec00000000),
		243: uint64(0x45b7664300000000),
		244: uint64(0x12445bb900000000),
		245: uint64(0x55d61f1600000000),
		246: uint64(0xdd66a33c00000000),
		247: uint64(0x9af4e79300000000),
		248: uint64(0x3286a91300000000),
		249: uint64(0x7514edbc00000000),
		250: uint64(0xfda4519600000000),
		251: uint64(0xba36153900000000),
		252: uint64(0xedc528c300000000),
		253: uint64(0xaa576c6c00000000),
		254: uint64(0x22e7d04600000000),
		255: uint64(0x657594e900000000),
	},
}
var _x2n_table = [32]Tz_crc_t{
	0:  uint32(0x40000000),
	1:  uint32(0x20000000),
	2:  uint32(0x08000000),
	3:  uint32(0x00800000),
	4:  uint32(0x00008000),
	5:  uint32(0xedb88320),
	6:  uint32(0xb1e6b092),
	7:  uint32(0xa06a2517),
	8:  uint32(0xed627dae),
	9:  uint32(0x88d14467),
	10: uint32(0xd7bbfe6a),
	11: uint32(0xec447f11),
	12: uint32(0x8e7ea170),
	13: uint32(0x6427800e),
	14: uint32(0x4d47bae0),
	15: uint32(0x09fe548f),
	16: uint32(0x83852d0f),
	17: uint32(0x30362f1a),
	18: uint32(0x7b5a9cc3),
	19: uint32(0x31fec169),
	20: uint32(0x9fec022a),
	21: uint32(0x6c8dedc4),
	22: uint32(0x15d6874d),
	23: uint32(0x5fde7a4e),
	24: uint32(0xbad90e37),
	25: uint32(0x2e4e5eef),
	26: uint32(0x4eaba214),
	27: uint32(0xa8a472c0),
	28: uint32(0x429a969e),
	29: uint32(0x148d302a),
	30: uint32(0xc40ba6d0),
	31: uint32(0xc4e22c3c),
}

/* CRC polynomial. */

// C documentation
//
//	/*
//	  Return a(x) multiplied by b(x) modulo p(x), where p(x) is the CRC polynomial,
//	  reflected. For speed, this requires that a not be zero.
//	 */
func _multmodp(tls *libc.TLS, a Tz_crc_t, b Tz_crc_t) (r Tz_crc_t) {
	var m, p Tz_crc_t
	var v2 uint32
	_, _, _ = m, p, v2
	m = libc.Uint32FromInt32(1) << libc.Int32FromInt32(31)
	p = uint32(0)
	for {
		if a&m != 0 {
			p ^= b
			if a&(m-uint32(1)) == uint32(0) {
				break
			}
		}
		m >>= uint32(1)
		if b&uint32(1) != 0 {
			v2 = b>>libc.Int32FromInt32(1) ^ uint32(m_POLY)
		} else {
			v2 = b >> int32(1)
		}
		b = v2
		goto _1
	_1:
	}
	return p
}

// C documentation
//
//	/*
//	  Return x^(n * 2^k) modulo p(x). Requires that x2n_table[] has been
//	  initialized.
//	 */
func _x2nmodp(tls *libc.TLS, n Toff_t, k uint32) (r Tz_crc_t) {
	var p Tz_crc_t
	_ = p
	p = libc.Uint32FromInt32(1) << libc.Int32FromInt32(31) /* x^0 == 1 */
	for n != 0 {
		if n&int64(1) != 0 {
			p = _multmodp(tls, _x2n_table[k&uint32(31)], p)
		}
		n >>= int64(1)
		k++
	}
	return p
}

// C documentation
//
//	/* =========================================================================
//	 * This function can be used by asm versions of crc32(), and to force the
//	 * generation of the CRC tables in a threaded application.
//	 */
func Xget_crc_table(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&_crc_table))
}

/* =========================================================================
 * Use ARM machine instructions if available. This will compute the CRC about
 * ten times faster than the braided calculation. This code does not check for
 * the presence of the CRC instruction at run time. __ARM_FEATURE_CRC32 will
 * only be defined if the compilation specifies an ARM processor architecture
 * that has the instructions. For example, compiling with -march=armv8.1-a or
 * -march=armv8-a+crc, or -march=native if the compile machine has the crc32
 * instructions.
 */

// C documentation
//
//	/*
//	  Return the CRC of the W bytes in the word_t data, taking the
//	  least-significant byte of the word as the first byte of data, without any pre
//	  or post conditioning. This is used to combine the CRCs of each braid.
//	 */
func _crc_word(tls *libc.TLS, data Tz_word_t) (r Tz_crc_t) {
	var k int32
	_ = k
	k = 0
	for {
		if !(k < int32(m_W)) {
			break
		}
		data = data>>libc.Int32FromInt32(8) ^ uint64(_crc_table[data&uint64(0xff)])
		goto _1
	_1:
		;
		k++
	}
	return uint32(data)
}

func _crc_word_big(tls *libc.TLS, data Tz_word_t) (r Tz_word_t) {
	var k int32
	_ = k
	k = 0
	for {
		if !(k < int32(m_W)) {
			break
		}
		data = data<<libc.Int32FromInt32(8) ^ _crc_big_table[data>>((libc.Int32FromInt32(m_W)-libc.Int32FromInt32(1))<<libc.Int32FromInt32(3))&uint64(0xff)]
		goto _1
	_1:
		;
		k++
	}
	return data
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_z(tls *libc.TLS, crc uint64, buf uintptr, len1 Tz_size_t) (r uint64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var blks, v2, v4 Tz_size_t
	var comb, crc01, crc11, crc21, crc31, crc41, word0, word01, word1, word11, word2, word21, word3, word31, word4, word41 Tz_word_t
	var crc0, crc1, crc2, crc3, crc4 Tz_crc_t
	var k int32
	var words, v1, v10, v11, v12, v13, v14, v6, v7, v8, v9 uintptr
	var _ /* endian at bp+0 */ uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = blks, comb, crc0, crc01, crc1, crc11, crc2, crc21, crc3, crc31, crc4, crc41, k, word0, word01, word1, word11, word2, word21, word3, word31, word4, word41, words, v1, v10, v11, v12, v13, v14, v2, v4, v6, v7, v8, v9
	/* Return initial CRC, if requested. */
	if buf == uintptr(m_Z_NULL) {
		return uint64(0)
	}
	/* Pre-condition the CRC */
	crc = ^crc & uint64(0xffffffff)
	/* If provided enough bytes, do a braided CRC calculation. */
	if len1 >= libc.Uint64FromInt32(libc.Int32FromInt32(m_N)*libc.Int32FromInt32(m_W)+libc.Int32FromInt32(m_W)-libc.Int32FromInt32(1)) {
		/* Compute the CRC up to a z_word_t boundary. */
		for len1 != 0 && uint64(buf)&libc.Uint64FromInt32(libc.Int32FromInt32(m_W)-libc.Int32FromInt32(1)) != uint64(0) {
			len1--
			v1 = buf
			buf++
			crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v1))))&uint64(0xff)])
		}
		/* Compute the CRC on as many N z_word_t blocks as are available. */
		blks = len1 / libc.Uint64FromInt32(libc.Int32FromInt32(m_N)*libc.Int32FromInt32(m_W))
		len1 -= blks * uint64(m_N) * uint64(m_W)
		words = buf
		/* Do endian check at execution time instead of compile time, since ARM
		   processors can change the endianness at execution time. If the
		   compiler knows what the endianness will be, it can optimize out the
		   check and the unused branch. */
		*(*uint32)(unsafe.Pointer(bp)) = uint32(1)
		if *(*uint8)(unsafe.Pointer(bp)) != 0 {
			/* Initialize the CRC for each braid. */
			crc0 = uint32(crc)
			crc1 = uint32(0)
			crc2 = uint32(0)
			crc3 = uint32(0)
			crc4 = uint32(0)
			/*
			   Process the first blks-1 blocks, computing the CRCs on each braid
			   independently.
			*/
			for {
				blks--
				v2 = blks
				if !(v2 != 0) {
					break
				}
				/* Load the word for each braid into registers. */
				word0 = uint64(crc0) ^ *(*Tz_word_t)(unsafe.Pointer(words))
				word1 = uint64(crc1) ^ *(*Tz_word_t)(unsafe.Pointer(words + 1*8))
				word2 = uint64(crc2) ^ *(*Tz_word_t)(unsafe.Pointer(words + 2*8))
				word3 = uint64(crc3) ^ *(*Tz_word_t)(unsafe.Pointer(words + 3*8))
				word4 = uint64(crc4) ^ *(*Tz_word_t)(unsafe.Pointer(words + 4*8))
				words += uintptr(m_N) * 8
				/* Compute and update the CRC for each word. The loop should
				   get unrolled. */
				crc0 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word0&uint64(0xff))*4))
				crc1 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word1&uint64(0xff))*4))
				crc2 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word2&uint64(0xff))*4))
				crc3 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word3&uint64(0xff))*4))
				crc4 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word4&uint64(0xff))*4))
				k = int32(1)
				for {
					if !(k < int32(m_W)) {
						break
					}
					crc0 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word0>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*4))
					crc1 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word1>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*4))
					crc2 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word2>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*4))
					crc3 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word3>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*4))
					crc4 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word4>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*4))
					goto _3
				_3:
					;
					k++
				}
			}
			/*
			   Process the last block, combining the CRCs of the N braids at the
			   same time.
			*/
			crc = uint64(_crc_word(tls, uint64(crc0)^*(*Tz_word_t)(unsafe.Pointer(words))))
			crc = uint64(_crc_word(tls, uint64(crc1)^*(*Tz_word_t)(unsafe.Pointer(words + 1*8))^crc))
			crc = uint64(_crc_word(tls, uint64(crc2)^*(*Tz_word_t)(unsafe.Pointer(words + 2*8))^crc))
			crc = uint64(_crc_word(tls, uint64(crc3)^*(*Tz_word_t)(unsafe.Pointer(words + 3*8))^crc))
			crc = uint64(_crc_word(tls, uint64(crc4)^*(*Tz_word_t)(unsafe.Pointer(words + 4*8))^crc))
			words += uintptr(m_N) * 8
		} else {
			/* Initialize the CRC for each braid. */
			crc01 = _byte_swap(tls, crc)
			crc11 = uint64(0)
			crc21 = uint64(0)
			crc31 = uint64(0)
			crc41 = uint64(0)
			/*
			   Process the first blks-1 blocks, computing the CRCs on each braid
			   independently.
			*/
			for {
				blks--
				v4 = blks
				if !(v4 != 0) {
					break
				}
				/* Load the word for each braid into registers. */
				word01 = crc01 ^ *(*Tz_word_t)(unsafe.Pointer(words))
				word11 = crc11 ^ *(*Tz_word_t)(unsafe.Pointer(words + 1*8))
				word21 = crc21 ^ *(*Tz_word_t)(unsafe.Pointer(words + 2*8))
				word31 = crc31 ^ *(*Tz_word_t)(unsafe.Pointer(words + 3*8))
				word41 = crc41 ^ *(*Tz_word_t)(unsafe.Pointer(words + 4*8))
				words += uintptr(m_N) * 8
				/* Compute and update the CRC for each word. The loop should
				   get unrolled. */
				crc01 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word01&uint64(0xff))*8))
				crc11 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word11&uint64(0xff))*8))
				crc21 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word21&uint64(0xff))*8))
				crc31 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word31&uint64(0xff))*8))
				crc41 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word41&uint64(0xff))*8))
				k = int32(1)
				for {
					if !(k < int32(m_W)) {
						break
					}
					crc01 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*2048 + uintptr(word01>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*8))
					crc11 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*2048 + uintptr(word11>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*8))
					crc21 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*2048 + uintptr(word21>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*8))
					crc31 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*2048 + uintptr(word31>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*8))
					crc41 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*2048 + uintptr(word41>>(k<<libc.Int32FromInt32(3))&uint64(0xff))*8))
					goto _5
				_5:
					;
					k++
				}
			}
			/*
			   Process the last block, combining the CRCs of the N braids at the
			   same time.
			*/
			comb = _crc_word_big(tls, crc01^*(*Tz_word_t)(unsafe.Pointer(words)))
			comb = _crc_word_big(tls, crc11^*(*Tz_word_t)(unsafe.Pointer(words + 1*8))^comb)
			comb = _crc_word_big(tls, crc21^*(*Tz_word_t)(unsafe.Pointer(words + 2*8))^comb)
			comb = _crc_word_big(tls, crc31^*(*Tz_word_t)(unsafe.Pointer(words + 3*8))^comb)
			comb = _crc_word_big(tls, crc41^*(*Tz_word_t)(unsafe.Pointer(words + 4*8))^comb)
			words += uintptr(m_N) * 8
			crc = _byte_swap(tls, comb)
		}
		/*
		   Update the pointer to the remaining bytes to process.
		*/
		buf = words
	}
	/* Complete the computation of the CRC on any remaining bytes. */
	for len1 >= uint64(8) {
		len1 -= uint64(8)
		v6 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v6))))&uint64(0xff)])
		v7 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v7))))&uint64(0xff)])
		v8 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v8))))&uint64(0xff)])
		v9 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v9))))&uint64(0xff)])
		v10 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v10))))&uint64(0xff)])
		v11 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v11))))&uint64(0xff)])
		v12 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v12))))&uint64(0xff)])
		v13 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v13))))&uint64(0xff)])
	}
	for len1 != 0 {
		len1--
		v14 = buf
		buf++
		crc = crc>>int32(8) ^ uint64(_crc_table[(crc^uint64(*(*uint8)(unsafe.Pointer(v14))))&uint64(0xff)])
	}
	/* Return the CRC, post-conditioned. */
	return crc ^ uint64(0xffffffff)
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32(tls *libc.TLS, crc uint64, buf uintptr, len1 TuInt) (r uint64) {
	return Xcrc32_z(tls, crc, buf, uint64(len1))
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_combine64(tls *libc.TLS, crc1 TuLong, crc2 TuLong, len2 Toff_t) (r TuLong) {
	return uint64(_multmodp(tls, _x2nmodp(tls, len2, uint32(3)), uint32(crc1))) ^ crc2&uint64(0xffffffff)
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_combine(tls *libc.TLS, crc1 TuLong, crc2 TuLong, len2 Toff_t) (r TuLong) {
	return Xcrc32_combine64(tls, crc1, crc2, len2)
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_combine_gen64(tls *libc.TLS, len2 Toff_t) (r TuLong) {
	return uint64(_x2nmodp(tls, len2, uint32(3)))
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_combine_gen(tls *libc.TLS, len2 Toff_t) (r TuLong) {
	return Xcrc32_combine_gen64(tls, len2)
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_combine_op(tls *libc.TLS, crc1 TuLong, crc2 TuLong, op TuLong) (r TuLong) {
	return uint64(_multmodp(tls, uint32(op), uint32(crc1))) ^ crc2&uint64(0xffffffff)
}

const m_BL_CODES = 19
const m_BUSY_STATE = 113
const m_Buf_size = 16
const m_COMMENT_STATE = 91
const m_D_CODES = 30
const m_EXTRA_STATE = 69
const m_FINISH_STATE = 666
const m_GZIP_STATE = 57
const m_HCRC_STATE = 103
const m_INIT_STATE = 42
const m_LENGTH_CODES = 29
const m_LITERALS = 256
const m_LIT_BUFS = 4
const m_MAX_BITS = 15
const m_MAX_STORED = 65535
const m_NAME_STATE = 73
const m_NIL = 0
const m_PRESET_DICT1 = 32
const m_TOO_FAR = 4096
const m_WIN_INIT = "MAX_MATCH"
const m_max_insert_length = "max_lazy_match"

type Tinternal_state = struct {
	Fstrm             Tz_streamp
	Fstatus           int32
	Fpending_buf      uintptr
	Fpending_buf_size Tulg
	Fpending_out      uintptr
	Fpending          Tulg
	Fwrap             int32
	Fgzhead           Tgz_headerp
	Fgzindex          Tulg
	Fmethod           TByte
	Flast_flush       int32
	Fw_size           TuInt
	Fw_bits           TuInt
	Fw_mask           TuInt
	Fwindow           uintptr
	Fwindow_size      Tulg
	Fprev             uintptr
	Fhead             uintptr
	Fins_h            TuInt
	Fhash_size        TuInt
	Fhash_bits        TuInt
	Fhash_mask        TuInt
	Fhash_shift       TuInt
	Fblock_start      int64
	Fmatch_length     TuInt
	Fprev_match       TIPos
	Fmatch_available  int32
	Fstrstart         TuInt
	Fmatch_start      TuInt
	Flookahead        TuInt
	Fprev_length      TuInt
	Fmax_chain_length TuInt
	Fmax_lazy_match   TuInt
	Flevel            int32
	Fstrategy         int32
	Fgood_match       TuInt
	Fnice_match       int32
	Fdyn_ltree        [573]Tct_data_s
	Fdyn_dtree        [61]Tct_data_s
	Fbl_tree          [39]Tct_data_s
	Fl_desc           Ttree_desc_s
	Fd_desc           Ttree_desc_s
	Fbl_desc          Ttree_desc_s
	Fbl_count         [16]Tush
	Fheap             [573]int32
	Fheap_len         int32
	Fheap_max         int32
	Fdepth            [573]Tuch
	Fsym_buf          uintptr
	Flit_bufsize      TuInt
	Fsym_next         TuInt
	Fsym_end          TuInt
	Fopt_len          Tulg
	Fstatic_len       Tulg
	Fmatches          TuInt
	Finsert           TuInt
	Fbi_buf           Tush
	Fbi_valid         int32
	Fhigh_water       Tulg
}

type Tct_data = struct {
	Ffc struct {
		Fcode [0]Tush
		Ffreq Tush
	}
	Fdl struct {
		Flen1 [0]Tush
		Fdad  Tush
	}
}

type Tct_data_s = Tct_data

type Ttree_desc = struct {
	Fdyn_tree  uintptr
	Fmax_code  int32
	Fstat_desc uintptr
}

type Ttree_desc_s = Ttree_desc

type TPos = uint16

type TPosf = uint16

type TIPos = uint32

type Tdeflate_state = struct {
	Fstrm             Tz_streamp
	Fstatus           int32
	Fpending_buf      uintptr
	Fpending_buf_size Tulg
	Fpending_out      uintptr
	Fpending          Tulg
	Fwrap             int32
	Fgzhead           Tgz_headerp
	Fgzindex          Tulg
	Fmethod           TByte
	Flast_flush       int32
	Fw_size           TuInt
	Fw_bits           TuInt
	Fw_mask           TuInt
	Fwindow           uintptr
	Fwindow_size      Tulg
	Fprev             uintptr
	Fhead             uintptr
	Fins_h            TuInt
	Fhash_size        TuInt
	Fhash_bits        TuInt
	Fhash_mask        TuInt
	Fhash_shift       TuInt
	Fblock_start      int64
	Fmatch_length     TuInt
	Fprev_match       TIPos
	Fmatch_available  int32
	Fstrstart         TuInt
	Fmatch_start      TuInt
	Flookahead        TuInt
	Fprev_length      TuInt
	Fmax_chain_length TuInt
	Fmax_lazy_match   TuInt
	Flevel            int32
	Fstrategy         int32
	Fgood_match       TuInt
	Fnice_match       int32
	Fdyn_ltree        [573]Tct_data_s
	Fdyn_dtree        [61]Tct_data_s
	Fbl_tree          [39]Tct_data_s
	Fl_desc           Ttree_desc_s
	Fd_desc           Ttree_desc_s
	Fbl_desc          Ttree_desc_s
	Fbl_count         [16]Tush
	Fheap             [573]int32
	Fheap_len         int32
	Fheap_max         int32
	Fdepth            [573]Tuch
	Fsym_buf          uintptr
	Flit_bufsize      TuInt
	Fsym_next         TuInt
	Fsym_end          TuInt
	Fopt_len          Tulg
	Fstatic_len       Tulg
	Fmatches          TuInt
	Finsert           TuInt
	Fbi_buf           Tush
	Fbi_valid         int32
	Fhigh_water       Tulg
}

/*
  If you use the zlib library in a product, an acknowledgment is welcome
  in the documentation of your product. If for some reason you cannot
  include such an acknowledgment, I would appreciate that you keep this
  copyright string in the executable of your product.
*/

type Tblock_state = int32

const _need_more = 0
const /* block not completed, need more input or more output */
_block_done = 1
const /* block flush performed */
_finish_started = 2
const /* finish started, need only more output at next deflate */
_finish_done = 3

type Tcompress_func = uintptr

/* ===========================================================================
 * Local data
 */

/* Tail of hash chains */

/* Matches of length 3 are discarded if their distance exceeds TOO_FAR */

// C documentation
//
//	/* Values for max_lazy_match, good_match and max_chain_length, depending on
//	 * the desired pack level (0..9). The values given below have been tuned to
//	 * exclude worst case performance for pathological files. Better values may be
//	 * found for specific files.
//	 */
type Tconfig = struct {
	Fgood_length Tush
	Fmax_lazy    Tush
	Fnice_length Tush
	Fmax_chain   Tush
	Ffunc1       Tcompress_func
}

/* ===========================================================================
 * Local data
 */

/* Tail of hash chains */

/* Matches of length 3 are discarded if their distance exceeds TOO_FAR */

// C documentation
//
//	/* Values for max_lazy_match, good_match and max_chain_length, depending on
//	 * the desired pack level (0..9). The values given below have been tuned to
//	 * exclude worst case performance for pathological files. Better values may be
//	 * found for specific files.
//	 */
type Tconfig_s = Tconfig

var _configuration_table = [10]Tconfig{
	0: {},
	1: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(4),
		Fnice_length: uint16(8),
		Fmax_chain:   uint16(4),
	},
	2: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(5),
		Fnice_length: uint16(16),
		Fmax_chain:   uint16(8),
	},
	3: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(6),
		Fnice_length: uint16(32),
		Fmax_chain:   uint16(32),
	},
	4: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(4),
		Fnice_length: uint16(16),
		Fmax_chain:   uint16(16),
	},
	5: {
		Fgood_length: uint16(8),
		Fmax_lazy:    uint16(16),
		Fnice_length: uint16(32),
		Fmax_chain:   uint16(32),
	},
	6: {
		Fgood_length: uint16(8),
		Fmax_lazy:    uint16(16),
		Fnice_length: uint16(128),
		Fmax_chain:   uint16(128),
	},
	7: {
		Fgood_length: uint16(8),
		Fmax_lazy:    uint16(32),
		Fnice_length: uint16(128),
		Fmax_chain:   uint16(256),
	},
	8: {
		Fgood_length: uint16(32),
		Fmax_lazy:    uint16(128),
		Fnice_length: uint16(258),
		Fmax_chain:   uint16(1024),
	},
	9: {
		Fgood_length: uint16(32),
		Fmax_lazy:    uint16(258),
		Fnice_length: uint16(258),
		Fmax_chain:   uint16(4096),
	},
}

func init() {
	p := unsafe.Pointer(&_configuration_table)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(_deflate_stored)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 40)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 136)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(_deflate_slow)
}

/* max compression */

/* Note: the deflate() code requires max_lazy >= MIN_MATCH and max_chain >= 4
 * For deflate_fast() (levels <= 3) good is ignored and lazy has a different
 * meaning.
 */

/* rank Z_BLOCK between Z_NO_FLUSH and Z_PARTIAL_FLUSH */

/* ===========================================================================
 * Update a hash value with the given input byte
 * IN  assertion: all calls to UPDATE_HASH are made with consecutive input
 *    characters, so that a running hash key can be computed from the previous
 *    key instead of complete recalculation each time.
 */

/* ===========================================================================
 * Insert string str in the dictionary and set match_head to the previous head
 * of the hash chain (the most recent string with same hash key). Return
 * the previous length of the hash chain.
 * If this file is compiled with -DFASTEST, the compression level is forced
 * to 1, and no hash chains are maintained.
 * IN  assertion: all calls to INSERT_STRING are made with consecutive input
 *    characters and the first MIN_MATCH bytes of str are valid (except for
 *    the last MIN_MATCH-1 bytes of the input file).
 */

/* ===========================================================================
 * Initialize the hash table (avoiding 64K overflow for 16 bit systems).
 * prev[] will be initialized on the fly.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Slide the hash table when sliding the window down (could be avoided with 32
//	 * bit values at the expense of memory usage). We slide even when level == 0 to
//	 * keep the hash table consistent if we switch back to level > 0 later.
//	 */
func _slide_hash(tls *libc.TLS, s uintptr) {
	var m, n, v1, v4, v5, v8 uint32
	var p, v3, v7 uintptr
	var wsize TuInt
	_, _, _, _, _, _, _, _, _, _ = m, n, p, wsize, v1, v3, v4, v5, v7, v8
	wsize = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	n = (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size
	p = (*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr(n)*2
	for {
		p -= 2
		v3 = p
		m = uint32(*(*TPosf)(unsafe.Pointer(v3)))
		if m >= wsize {
			v4 = m - wsize
		} else {
			v4 = uint32(m_NIL)
		}
		*(*TPosf)(unsafe.Pointer(p)) = uint16(v4)
		goto _2
	_2:
		;
		n--
		v1 = n
		if !(v1 != 0) {
			break
		}
	}
	n = wsize
	p = (*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr(n)*2
	for {
		p -= 2
		v7 = p
		m = uint32(*(*TPosf)(unsafe.Pointer(v7)))
		if m >= wsize {
			v8 = m - wsize
		} else {
			v8 = uint32(m_NIL)
		}
		*(*TPosf)(unsafe.Pointer(p)) = uint16(v8)
		/* If n is not on any hash chain, prev[n] is garbage but
		 * its value will never be used.
		 */
		goto _6
	_6:
		;
		n--
		v5 = n
		if !(v5 != 0) {
			break
		}
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Read a new buffer from the current input stream, update the adler32
//	 * and total number of bytes read.  All deflate() input goes through
//	 * this function so some applications may wish to modify it to avoid
//	 * allocating a large strm->next_in buffer and copying from it.
//	 * (See also flush_pending()).
//	 */
func _read_buf(tls *libc.TLS, strm Tz_streamp, buf uintptr, size uint32) (r uint32) {
	var len1 uint32
	_ = len1
	len1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	if len1 > size {
		len1 = size
	}
	if len1 == uint32(0) {
		return uint32(0)
	}
	*(*TuInt)(unsafe.Pointer(strm + 8)) -= len1
	libc.Xmemcpy(tls, buf, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, uint64(len1))
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap == int32(1) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xadler32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, buf, len1)
	} else {
		if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap == int32(2) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, buf, len1)
		}
	}
	*(*uintptr)(unsafe.Pointer(strm)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 16)) += uint64(len1)
	return len1
}

// C documentation
//
//	/* ===========================================================================
//	 * Fill the window when the lookahead becomes insufficient.
//	 * Updates strstart and lookahead.
//	 *
//	 * IN assertion: lookahead < MIN_LOOKAHEAD
//	 * OUT assertions: strstart <= window_size-MIN_LOOKAHEAD
//	 *    At least one byte has been read, or avail_in == 0; reads are
//	 *    performed for at least two bytes (required for the zip translate_eol
//	 *    option -- not supported here).
//	 */
func _fill_window(tls *libc.TLS, s uintptr) {
	var curr, init1 Tulg
	var more, n uint32
	var str, wsize TuInt
	_, _, _, _, _, _ = curr, init1, more, n, str, wsize /* Amount of free space at the end of the window. */
	wsize = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	for cond := true; cond; cond = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in != uint32(0) {
		more = uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - uint64((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead) - uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart))
		/* Deal with !@#$% 64K limit: */
		if uint64(4) <= uint64(2) {
			if more == uint32(0) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart == uint32(0) && (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				more = wsize
			} else {
				if more == libc.Uint32FromInt32(-libc.Int32FromInt32(1)) {
					/* Very unlikely, but possible on 16 bit machine if
					 * strstart == 0 && lookahead == 1 (input done a byte at time)
					 */
					more--
				}
			}
		}
		/* If the window is almost full and there is insufficient lookahead,
		 * move the upper half to the lower one to make room in the upper half.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart >= wsize+((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1))) {
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr(wsize), uint64(wsize-more))
			*(*TuInt)(unsafe.Pointer(s + 176)) -= wsize
			*(*TuInt)(unsafe.Pointer(s + 172)) -= wsize /* we now have strstart >= MAX_DIST */
			*(*int64)(unsafe.Pointer(s + 152)) -= libc.Int64FromUint32(wsize)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Finsert > (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
				(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
			}
			_slide_hash(tls, s)
			more += wsize
		}
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) {
			break
		}
		/* If there was no sliding:
		 *    strstart <= WSIZE+MAX_DIST-1 && lookahead <= MIN_LOOKAHEAD - 1 &&
		 *    more == window_size - lookahead - strstart
		 * => more >= window_size - (MIN_LOOKAHEAD-1 + WSIZE + MAX_DIST-1)
		 * => more >= window_size - 2*WSIZE + 2
		 * In the BIG_MEM or MMAP case (not yet supported),
		 *   window_size == input_size + MIN_LOOKAHEAD  &&
		 *   strstart + s->lookahead <= input_size => more >= MIN_LOOKAHEAD.
		 * Otherwise, window_size == 2*WSIZE so more >= 2.
		 * If there was sliding, more >= WSIZE. So in all cases, more >= 2.
		 */
		n = _read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead), more)
		*(*TuInt)(unsafe.Pointer(s + 180)) += n
		/* Initialize the hash value now that we have some input: */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead+(*Tdeflate_state)(unsafe.Pointer(s)).Finsert >= uint32(m_MIN_MATCH) {
			str = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str))))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str+uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			for (*Tdeflate_state)(unsafe.Pointer(s)).Finsert != 0 {
				(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str+uint32(m_MIN_MATCH)-uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr(str&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16(str)
				str++
				(*Tdeflate_state)(unsafe.Pointer(s)).Finsert--
				if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead+(*Tdeflate_state)(unsafe.Pointer(s)).Finsert < uint32(m_MIN_MATCH) {
					break
				}
			}
		}
		/* If the whole input has less than MIN_MATCH bytes, ins_h is garbage,
		 * but this is not important since only literal bytes will be emitted.
		 */
	}
	/* If the WIN_INIT bytes after the end of the current data have never been
	 * written, then zero those bytes in order to avoid memory check reports of
	 * the use of uninitialized (or uninitialised as Julian writes) bytes by
	 * the longest match routines.  Update the high water mark for the next
	 * time through here.  WIN_INIT is set to MAX_MATCH since the longest match
	 * routines allow scanning to strstart + MAX_MATCH, ignoring lookahead.
	 */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size {
		curr = uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) + uint64((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < curr {
			/* Previous high water mark below current data -- zero WIN_INIT
			 * bytes or up to end of window, whichever is less.
			 */
			init1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - curr
			if init1 > uint64(m_MAX_MATCH) {
				init1 = uint64(m_MAX_MATCH)
			}
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr(curr), 0, uint64(uint32(init1)))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = curr + init1
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < curr+uint64(m_MAX_MATCH) {
				/* High water mark at or above current data, but below current data
				 * plus WIN_INIT -- zero out to current data plus WIN_INIT, or up
				 * to end of window, whichever is less.
				 */
				init1 = curr + uint64(m_MAX_MATCH) - (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water
				if init1 > (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size-(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water {
					init1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water
				}
				libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water), 0, uint64(uint32(init1)))
				*(*Tulg)(unsafe.Pointer(s + 5944)) += init1
			}
		}
	}
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateInit_(tls *libc.TLS, strm Tz_streamp, level int32, version uintptr, stream_size int32) (r int32) {
	return XdeflateInit2_(tls, strm, level, int32(m_Z_DEFLATED), int32(m_MAX_WBITS), int32(m_DEF_MEM_LEVEL), m_Z_DEFAULT_STRATEGY, version, stream_size)
	/* To do: ignore strm->next_in if we use it as window */
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateInit2_(tls *libc.TLS, strm Tz_streamp, level int32, method int32, windowBits int32, memLevel int32, strategy int32, version uintptr, stream_size int32) (r int32) {
	var s uintptr
	var wrap int32
	_, _ = s, wrap
	wrap = int32(1)
	if version == uintptr(m_Z_NULL) || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(version))) != libc.Int32FromUint8(_my_version[0]) || libc.Uint64FromInt32(stream_size) != uint64(112) {
		return -int32(6)
	}
	if strm == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = __ccgo_fp(Xzcalloc)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = libc.UintptrFromInt32(0)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = __ccgo_fp(Xzcfree)
	}
	if level == -int32(1) {
		level = int32(6)
	}
	if windowBits < 0 { /* suppress zlib wrapper */
		wrap = 0
		if windowBits < -int32(15) {
			return -int32(2)
		}
		windowBits = -windowBits
	} else {
		if windowBits > int32(15) {
			wrap = int32(2) /* write gzip wrapper instead */
			windowBits -= int32(16)
		}
	}
	if memLevel < int32(1) || memLevel > int32(m_MAX_MEM_LEVEL) || method != int32(m_Z_DEFLATED) || windowBits < int32(8) || windowBits > int32(15) || level < 0 || level > int32(9) || strategy < 0 || strategy > int32(m_Z_FIXED) || windowBits == int32(8) && wrap != int32(1) {
		return -int32(2)
	}
	if windowBits == int32(8) {
		windowBits = int32(9)
	} /* until 256-byte window bug fixed */
	s = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), uint32(libc.Uint64FromInt64(5952)))
	if s == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = s
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrm = strm
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_INIT_STATE) /* to pass state test in deflateReset() */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = wrap
	(*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead = uintptr(m_Z_NULL)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits = libc.Uint32FromInt32(windowBits)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fw_size = libc.Uint32FromInt32(int32(1) << (*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - uint32(1)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits = libc.Uint32FromInt32(memLevel) + uint32(7)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size = libc.Uint32FromInt32(int32(1) << (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask = (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size - uint32(1)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift = ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits + libc.Uint32FromInt32(m_MIN_MATCH) - libc.Uint32FromInt32(1)) / libc.Uint32FromInt32(m_MIN_MATCH)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size, uint32(libc.Uint64FromInt32(2)*libc.Uint64FromInt64(1)))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size, uint32(libc.Uint64FromInt64(2)))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhead = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size, uint32(libc.Uint64FromInt64(2)))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = uint64(0)                                                /* nothing written to s->window yet */
	(*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize = libc.Uint32FromInt32(int32(1) << (memLevel + int32(6))) /* 16K elements by default */
	/* We overlay pending_buf and sym_buf. This works since the average size
	 * for length/distance pairs over any compressed block is assured to be 31
	 * bits or less.
	 *
	 * Analysis: The longest fixed codes are a length code of 8 bits plus 5
	 * extra bits, for lengths 131 to 257. The longest fixed distance codes are
	 * 5 bits plus 13 extra bits, for distances 16385 to 32768. The longest
	 * possible fixed-codes length/distance pair is then 31 bits total.
	 *
	 * sym_buf starts one-fourth of the way into pending_buf. So there are
	 * three bytes in sym_buf for every four bytes in pending_buf. Each symbol
	 * in sym_buf is three bytes -- two for the distance and one for the
	 * literal/length. As each symbol is consumed, the pointer to the next
	 * sym_buf value to read moves forward three bytes. From that symbol, up to
	 * 31 bits are written to pending_buf. The closest the written pending_buf
	 * bits gets to the next sym_buf symbol to read is just before the last
	 * code is written. At that time, 31*(n - 2) bits have been written, just
	 * after 24*(n - 2) bits have been consumed from sym_buf. sym_buf starts at
	 * 8*n bits into pending_buf. (Note that the symbol buffer fills when n - 1
	 * symbols are written.) The closest the writing gets to what is unread is
	 * then n + 14 bits. Here n is lit_bufsize, which is 16384 by default, and
	 * can range from 128 to 32768.
	 *
	 * Therefore, at a minimum, there are 142 bits of space between what is
	 * written and what is read in the overlain buffers, so the symbols cannot
	 * be overwritten by the compressed data. That space is actually 139 bits,
	 * due to the three-bit fixed-code block header.
	 *
	 * That covers the case where either Z_FIXED is specified, forcing fixed
	 * codes, or when the use of fixed codes is chosen, because that choice
	 * results in a smaller compressed block than dynamic codes. That latter
	 * condition then assures that the above analysis also covers all dynamic
	 * blocks. A dynamic-code block will only be chosen to be emitted if it has
	 * fewer bits than a fixed-code block would for the same set of symbols.
	 * Therefore its average symbol length is assured to be less than 31. So
	 * the compressed data for a dynamic block also cannot overwrite the
	 * symbols from which it is being constructed.
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize, libc.Uint32FromInt32(libc.Int32FromInt32(m_LIT_BUFS)))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size = uint64((*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize) * uint64(4)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fprev == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fhead == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf == uintptr(m_Z_NULL) {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_FINISH_STATE)
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = Xz_errmsg[libc.Int32FromInt32(2) - -libc.Int32FromInt32(4)]
		XdeflateEnd(tls, strm)
		return -int32(4)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end = ((*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize - uint32(1)) * uint32(3)
	/* We avoid equality with lit_bufsize*3 because of wraparound at 64K
	 * on 16 bit machines and because stored blocks are restricted to
	 * 64K-1 bytes.
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Flevel = level
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy = strategy
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmethod = libc.Uint8FromInt32(method)
	return XdeflateReset(tls, strm)
}

var _my_version = [6]uint8{'1', '.', '3', '.', '1'}

// C documentation
//
//	/* =========================================================================
//	 * Check for a valid deflate stream state. Return 0 if ok, 1 if not.
//	 */
func _deflateStateCheck(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var s uintptr
	_ = s
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) || (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		return int32(1)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if s == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm != strm || (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_INIT_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_GZIP_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_EXTRA_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_NAME_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_COMMENT_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_HCRC_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_BUSY_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_FINISH_STATE) {
		return int32(1)
	}
	return 0
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateSetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength TuInt) (r int32) {
	var avail uint32
	var n, str, v1, v3 TuInt
	var next, s uintptr
	var wrap int32
	_, _, _, _, _, _, _, _ = avail, n, next, s, str, wrap, v1, v3
	if _deflateStateCheck(tls, strm) != 0 || dictionary == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	wrap = (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap
	if wrap == int32(2) || wrap == int32(1) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_INIT_STATE) || (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead != 0 {
		return -int32(2)
	}
	/* when using zlib wrappers, compute Adler-32 for provided dictionary */
	if wrap == int32(1) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xadler32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, dictionary, dictLength)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = 0 /* avoid computing Adler-32 in read_buf */
	/* if dictionary would fill window, just replace the history */
	if dictLength >= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		if wrap == 0 { /* already empty otherwise */
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint64(2))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = uint32(0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = 0
			(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
		}
		dictionary += uintptr(dictLength - (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size) /* use the tail */
		dictLength = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	}
	/* insert dictionary into window and hash */
	avail = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	next = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = dictLength
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = dictionary
	_fill_window(tls, s)
	for (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
		str = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		n = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead - libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))
		for {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str+uint32(m_MIN_MATCH)-uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr(str&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16(str)
			str++
			goto _2
		_2:
			;
			n--
			v1 = n
			if !(v1 != 0) {
				break
			}
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = str
		(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
		_fill_window(tls, s)
	}
	*(*TuInt)(unsafe.Pointer(s + 172)) += (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead = uint32(0)
	v3 = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length = v3
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = v3
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = avail
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = wrap
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateGetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength uintptr) (r int32) {
	var len1 TuInt
	var s uintptr
	_, _ = len1, s
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	len1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart + (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	if len1 > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		len1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	}
	if dictionary != uintptr(m_Z_NULL) && len1 != 0 {
		libc.Xmemcpy(tls, dictionary, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead)-uintptr(len1), uint64(len1))
	}
	if dictLength != uintptr(m_Z_NULL) {
		*(*TuInt)(unsafe.Pointer(dictLength)) = len1
	}
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateResetKeep(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var s uintptr
	var v1 TuLong
	var v2 int32
	var v3 uint64
	_, _, _, _ = s, v1, v2, v3
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	v1 = libc.Uint64FromInt32(0)
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* use zfree if we ever allocate msg dynamically */
	(*Tz_stream)(unsafe.Pointer(strm)).Fdata_type = int32(m_Z_UNKNOWN)
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending = uint64(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap < 0 {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = -(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap /* was made negative by deflate(..., Z_FINISH); */
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == int32(2) {
		v2 = int32(m_GZIP_STATE)
	} else {
		v2 = int32(m_INIT_STATE)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = v2
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == int32(2) {
		v3 = Xcrc32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
	} else {
		v3 = Xadler32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v3
	(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(2)
	x__tr_init(tls, s)
	return m_Z_OK
}

// C documentation
//
//	/* ===========================================================================
//	 * Initialize the "longest match" routines for a new zlib stream
//	 */
func _lm_init(tls *libc.TLS, s uintptr) {
	var v1 TuInt
	_ = v1
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size = libc.Uint64FromInt64(2) * uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
	*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
	libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint64(2))
	/* Set the default configuration parameters:
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match = uint32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fmax_lazy)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match = uint32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fgood_length)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match = libc.Int32FromUint16(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fnice_length)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length = uint32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fmax_chain)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = uint32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = 0
	(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead = uint32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
	v1 = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length = v1
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = v1
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
	(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = uint32(0)
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateReset(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var ret int32
	_ = ret
	ret = XdeflateResetKeep(tls, strm)
	if ret == m_Z_OK {
		_lm_init(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	}
	return ret
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateSetHeader(tls *libc.TLS, strm Tz_streamp, head Tgz_headerp) (r int32) {
	if _deflateStateCheck(tls, strm) != 0 || (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap != int32(2) {
		return -int32(2)
	}
	(*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fgzhead = head
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func XdeflatePending(tls *libc.TLS, strm Tz_streamp, pending uintptr, bits uintptr) (r int32) {
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	if pending != uintptr(m_Z_NULL) {
		*(*uint32)(unsafe.Pointer(pending)) = uint32((*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending)
	}
	if bits != uintptr(m_Z_NULL) {
		*(*int32)(unsafe.Pointer(bits)) = (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fbi_valid
	}
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func XdeflatePrime(tls *libc.TLS, strm Tz_streamp, bits int32, value int32) (r int32) {
	var put int32
	var s, p1 uintptr
	_, _, _ = put, s, p1
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if bits < 0 || bits > int32(16) || (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf < (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out+uintptr((libc.Int32FromInt32(m_Buf_size)+libc.Int32FromInt32(7))>>libc.Int32FromInt32(3)) {
		return -int32(5)
	}
	for cond := true; cond; cond = bits != 0 {
		put = int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid
		if put > bits {
			put = bits
		}
		p1 = s + 5936
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(value&(libc.Int32FromInt32(1)<<put-libc.Int32FromInt32(1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)))
		*(*int32)(unsafe.Pointer(s + 5940)) += put
		x__tr_flush_bits(tls, s)
		value >>= put
		bits -= put
	}
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateParams(tls *libc.TLS, strm Tz_streamp, level int32, strategy int32) (r int32) {
	var err int32
	var func1 Tcompress_func
	var s uintptr
	_, _, _ = err, func1, s
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if level == -int32(1) {
		level = int32(6)
	}
	if level < 0 || level > int32(9) || strategy < 0 || strategy > int32(m_Z_FIXED) {
		return -int32(2)
	}
	func1 = _configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Ffunc1
	if (strategy != (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy || func1 != _configuration_table[level].Ffunc1) && (*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush != -int32(2) {
		/* Flush the last buffer: */
		err = Xdeflate(tls, strm, int32(m_Z_BLOCK))
		if err == -int32(2) {
			return err
		}
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 || libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start+libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead) != 0 {
			return -int32(5)
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel != level {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches != uint32(0) {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches == uint32(1) {
				_slide_hash(tls, s)
			} else {
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
				libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint64(2))
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches = uint32(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Flevel = level
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match = uint32(_configuration_table[level].Fmax_lazy)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match = uint32(_configuration_table[level].Fgood_length)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match = libc.Int32FromUint16(_configuration_table[level].Fnice_length)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length = uint32(_configuration_table[level].Fmax_chain)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy = strategy
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateTune(tls *libc.TLS, strm Tz_streamp, good_length int32, max_lazy int32, nice_length int32, max_chain int32) (r int32) {
	var s uintptr
	_ = s
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match = libc.Uint32FromInt32(good_length)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match = libc.Uint32FromInt32(max_lazy)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match = nice_length
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length = libc.Uint32FromInt32(max_chain)
	return m_Z_OK
}

// C documentation
//
//	/* =========================================================================
//	 * For the default windowBits of 15 and memLevel of 8, this function returns a
//	 * close to exact, as well as small, upper bound on the compressed size. This
//	 * is an expansion of ~0.03%, plus a small constant.
//	 *
//	 * For any setting other than those defaults for windowBits and memLevel, one
//	 * of two worst case bounds is returned. This is at most an expansion of ~4% or
//	 * ~13%, plus a small constant.
//	 *
//	 * Both the 0.03% and 4% derive from the overhead of stored blocks. The first
//	 * one is for stored blocks of 16383 bytes (memLevel == 8), whereas the second
//	 * is for stored blocks of 127 bytes (the worst case memLevel == 1). The
//	 * expansion results from five bytes of header for each stored block.
//	 *
//	 * The larger expansion of 13% results from a window size less than or equal to
//	 * the symbols buffer size (windowBits <= memLevel + 7). In that case some of
//	 * the data being compressed may have slid out of the sliding window, impeding
//	 * a stored block from being emitted. Then the only choice is a fixed or
//	 * dynamic block, where a fixed block limits the maximum expansion to 9 bits
//	 * per 8-bit byte, plus 10 bits for every block. The smallest block size for
//	 * which this can occur is 255 (memLevel == 2).
//	 *
//	 * Shifts are used to approximate divisions, for speed.
//	 */
func XdeflateBound(tls *libc.TLS, strm Tz_streamp, sourceLen TuLong) (r TuLong) {
	var fixedlen, storelen, wraplen TuLong
	var s, str, v3, v5 uintptr
	var v1, v7 uint64
	var v2 int32
	_, _, _, _, _, _, _, _, _, _ = fixedlen, s, storelen, str, wraplen, v1, v2, v3, v5, v7
	/* upper bound for fixed blocks with 9-bit literals and length 255
	   (memLevel == 2, which is the lowest that may not use stored blocks) --
	   ~13% overhead plus a small constant */
	fixedlen = sourceLen + sourceLen>>libc.Int32FromInt32(3) + sourceLen>>libc.Int32FromInt32(8) + sourceLen>>libc.Int32FromInt32(9) + uint64(4)
	/* upper bound for stored blocks with length 127 (memLevel == 1) --
	   ~4% overhead plus a small constant */
	storelen = sourceLen + sourceLen>>libc.Int32FromInt32(5) + sourceLen>>libc.Int32FromInt32(7) + sourceLen>>libc.Int32FromInt32(11) + uint64(7)
	/* if can't get parameters, return larger bound plus a zlib wrapper */
	if _deflateStateCheck(tls, strm) != 0 {
		if fixedlen > storelen {
			v1 = fixedlen
		} else {
			v1 = storelen
		}
		return v1 + uint64(6)
	}
	/* compute wrapper length */
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	switch (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap {
	case 0: /* raw deflate */
		wraplen = uint64(0)
	case int32(1): /* zlib wrapper */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != 0 {
			v2 = int32(4)
		} else {
			v2 = 0
		}
		wraplen = libc.Uint64FromInt32(int32(6) + v2)
	case int32(2): /* gzip wrapper */
		wraplen = uint64(18)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead != uintptr(m_Z_NULL) {
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
				wraplen += uint64(uint32(2) + (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len)
			}
			str = (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname
			if str != uintptr(m_Z_NULL) {
				for {
					wraplen++
					goto _4
				_4:
					;
					v3 = str
					str++
					if !(*(*TBytef)(unsafe.Pointer(v3)) != 0) {
						break
					}
				}
			}
			str = (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment
			if str != uintptr(m_Z_NULL) {
				for {
					wraplen++
					goto _6
				_6:
					;
					v5 = str
					str++
					if !(*(*TBytef)(unsafe.Pointer(v5)) != 0) {
						break
					}
				}
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				wraplen += uint64(2)
			}
		}
	default: /* for compiler happiness */
		wraplen = uint64(6)
	}
	/* if not default parameters, return one of the conservative bounds */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits != uint32(15) || (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits != libc.Uint32FromInt32(libc.Int32FromInt32(8)+libc.Int32FromInt32(7)) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits <= (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits && (*Tdeflate_state)(unsafe.Pointer(s)).Flevel != 0 {
			v7 = fixedlen
		} else {
			v7 = storelen
		}
		return v7 + wraplen
	}
	/* default settings: return tight bound for that case -- ~0.03% overhead
	   plus a small constant */
	return sourceLen + sourceLen>>libc.Int32FromInt32(12) + sourceLen>>libc.Int32FromInt32(14) + sourceLen>>libc.Int32FromInt32(25) + uint64(13) - uint64(6) + wraplen
}

// C documentation
//
//	/* =========================================================================
//	 * Put a short in the pending buffer. The 16-bit value is put in MSB order.
//	 * IN assertion: the stream state is correct and there is enough room in
//	 * pending_buf.
//	 */
func _putShortMSB(tls *libc.TLS, s uintptr, b TuInt) {
	var v1, v3 Tulg
	var v2, v4 uintptr
	_, _, _, _ = v1, v2, v3, v4
	v2 = s + 40
	v1 = *(*Tulg)(unsafe.Pointer(v2))
	*(*Tulg)(unsafe.Pointer(v2))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = uint8(b >> libc.Int32FromInt32(8))
	v4 = s + 40
	v3 = *(*Tulg)(unsafe.Pointer(v4))
	*(*Tulg)(unsafe.Pointer(v4))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = uint8(b & libc.Uint32FromInt32(0xff))
}

// C documentation
//
//	/* =========================================================================
//	 * Flush as much pending output as possible. All deflate() output, except for
//	 * some deflate_stored() output, goes through this function so some
//	 * applications may wish to modify it to avoid allocating a large
//	 * strm->next_out buffer and copying into it. (See also read_buf()).
//	 */
func _flush_pending(tls *libc.TLS, strm Tz_streamp) {
	var len1 uint32
	var s uintptr
	_, _ = len1, s
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	x__tr_flush_bits(tls, s)
	len1 = uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending)
	if len1 > (*Tz_stream)(unsafe.Pointer(strm)).Favail_out {
		len1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	}
	if len1 == uint32(0) {
		return
	}
	libc.Xmemcpy(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out, uint64(len1))
	*(*uintptr)(unsafe.Pointer(strm + 24)) += uintptr(len1)
	*(*uintptr)(unsafe.Pointer(s + 32)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 40)) += uint64(len1)
	*(*TuInt)(unsafe.Pointer(strm + 32)) -= len1
	*(*Tulg)(unsafe.Pointer(s + 40)) -= uint64(len1)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == uint64(0) {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf
	}
}

/* ===========================================================================
 * Update the header CRC with the bytes s->pending_buf[beg..s->pending - 1].
 */

// C documentation
//
//	/* ========================================================================= */
func Xdeflate(tls *libc.TLS, strm Tz_streamp, flush int32) (r int32) {
	var beg, beg1, beg2, v10, v12, v14, v16, v18, v20, v24, v26, v33, v35, v37, v39, v4, v41, v45, v47, v49, v51, v53, v55, v57, v59, v6, v61, v66, v68, v70, v72, v74, v76, v78, v8, v80 Tulg
	var bstate Tblock_state
	var copy1, header, left, level_flags TuInt
	var old_flush, val, val1, v1, v2, v22, v23, v28, v29, v30, v31, v32, v43, v44, v63, v64, v65, v82 int32
	var s, v11, v13, v15, v17, v19, v21, v25, v27, v34, v36, v38, v40, v42, v46, v48, v5, v50, v52, v54, v56, v58, v60, v62, v67, v69, v7, v71, v73, v75, v77, v79, v81, v9 uintptr
	var v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = beg, beg1, beg2, bstate, copy1, header, left, level_flags, old_flush, s, val, val1, v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v5, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v7, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v8, v80, v81, v82, v9
	if _deflateStateCheck(tls, strm) != 0 || flush > int32(m_Z_BLOCK) || flush < 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) && (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_FINISH_STATE) && flush != int32(m_Z_FINISH) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = Xz_errmsg[libc.Int32FromInt32(2) - -libc.Int32FromInt32(2)]
		return -libc.Int32FromInt32(2)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = Xz_errmsg[libc.Int32FromInt32(2) - -libc.Int32FromInt32(5)]
		return -libc.Int32FromInt32(5)
	}
	old_flush = (*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush
	(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = flush
	/* Flush as much pending output as possible */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
		_flush_pending(tls, strm)
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
			/* Since avail_out is 0, deflate will be called again with
			 * more output space, but possibly with both pending and
			 * avail_in equal to zero. There won't be anything to do,
			 * but this is not an error situation so make sure we
			 * return OK instead of BUF_ERROR at next call of deflate:
			 */
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
		/* Make sure there is something to do and avoid duplicate consecutive
		 * flushes. For repeated and useless calls with Z_FINISH, we keep
		 * returning Z_STREAM_END instead of Z_BUF_ERROR.
		 */
	} else {
		if v3 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0); v3 {
			if flush > int32(4) {
				v1 = int32(9)
			} else {
				v1 = 0
			}
			if old_flush > int32(4) {
				v2 = int32(9)
			} else {
				v2 = 0
			}
		}
		if v3 && flush*int32(2)-v1 <= old_flush*int32(2)-v2 && flush != int32(m_Z_FINISH) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = Xz_errmsg[libc.Int32FromInt32(2) - -libc.Int32FromInt32(5)]
			return -libc.Int32FromInt32(5)
		}
	}
	/* User must not provide more input after the first FINISH: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_FINISH_STATE) && (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = Xz_errmsg[libc.Int32FromInt32(2) - -libc.Int32FromInt32(5)]
		return -libc.Int32FromInt32(5)
	}
	/* Write the header */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_INIT_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == 0 {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_INIT_STATE) {
		/* zlib header */
		header = (uint32(m_Z_DEFLATED) + ((*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits-uint32(8))<<int32(4)) << int32(8)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy >= int32(m_Z_HUFFMAN_ONLY) || (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(2) {
			level_flags = uint32(0)
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(6) {
				level_flags = uint32(1)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == int32(6) {
					level_flags = uint32(2)
				} else {
					level_flags = uint32(3)
				}
			}
		}
		header |= level_flags << libc.Int32FromInt32(6)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != uint32(0) {
			header |= uint32(m_PRESET_DICT1)
		}
		header += uint32(31) - header%uint32(31)
		_putShortMSB(tls, s, header)
		/* Save the adler32 of the preset dictionary: */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != uint32(0) {
			_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler>>libc.Int32FromInt32(16)))
			_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler&libc.Uint64FromInt32(0xffff)))
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xadler32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
		/* Compression must start with an empty pending buffer */
		_flush_pending(tls, strm)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_GZIP_STATE) {
		/* gzip header */
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		v5 = s + 40
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromInt32(31))
		v7 = s + 40
		v6 = *(*Tulg)(unsafe.Pointer(v7))
		*(*Tulg)(unsafe.Pointer(v7))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v6))) = libc.Uint8FromInt32(libc.Int32FromInt32(139))
		v9 = s + 40
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromInt32(8))
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead == uintptr(m_Z_NULL) {
			v11 = s + 40
			v10 = *(*Tulg)(unsafe.Pointer(v11))
			*(*Tulg)(unsafe.Pointer(v11))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v13 = s + 40
			v12 = *(*Tulg)(unsafe.Pointer(v13))
			*(*Tulg)(unsafe.Pointer(v13))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v12))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v15 = s + 40
			v14 = *(*Tulg)(unsafe.Pointer(v15))
			*(*Tulg)(unsafe.Pointer(v15))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v17 = s + 40
			v16 = *(*Tulg)(unsafe.Pointer(v17))
			*(*Tulg)(unsafe.Pointer(v17))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v16))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v19 = s + 40
			v18 = *(*Tulg)(unsafe.Pointer(v19))
			*(*Tulg)(unsafe.Pointer(v19))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v18))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v21 = s + 40
			v20 = *(*Tulg)(unsafe.Pointer(v21))
			*(*Tulg)(unsafe.Pointer(v21))++
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == int32(9) {
				v22 = int32(2)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy >= int32(m_Z_HUFFMAN_ONLY) || (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(2) {
					v23 = int32(4)
				} else {
					v23 = 0
				}
				v22 = v23
			}
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v20))) = libc.Uint8FromInt32(v22)
			v25 = s + 40
			v24 = *(*Tulg)(unsafe.Pointer(v25))
			*(*Tulg)(unsafe.Pointer(v25))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = libc.Uint8FromInt32(libc.Int32FromInt32(m_OS_CODE))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
			/* Compression must start with an empty pending buffer */
			_flush_pending(tls, strm)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
				return m_Z_OK
			}
		} else {
			v27 = s + 40
			v26 = *(*Tulg)(unsafe.Pointer(v27))
			*(*Tulg)(unsafe.Pointer(v27))++
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftext != 0 {
				v28 = int32(1)
			} else {
				v28 = 0
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				v29 = int32(2)
			} else {
				v29 = 0
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra == uintptr(m_Z_NULL) {
				v30 = 0
			} else {
				v30 = int32(4)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname == uintptr(m_Z_NULL) {
				v31 = 0
			} else {
				v31 = int32(8)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment == uintptr(m_Z_NULL) {
				v32 = 0
			} else {
				v32 = int32(16)
			}
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = libc.Uint8FromInt32(v28 + v29 + v30 + v31 + v32)
			v34 = s + 40
			v33 = *(*Tulg)(unsafe.Pointer(v34))
			*(*Tulg)(unsafe.Pointer(v34))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v33))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime & libc.Uint64FromInt32(0xff))
			v36 = s + 40
			v35 = *(*Tulg)(unsafe.Pointer(v36))
			*(*Tulg)(unsafe.Pointer(v36))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v35))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(8) & libc.Uint64FromInt32(0xff))
			v38 = s + 40
			v37 = *(*Tulg)(unsafe.Pointer(v38))
			*(*Tulg)(unsafe.Pointer(v38))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v37))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(16) & libc.Uint64FromInt32(0xff))
			v40 = s + 40
			v39 = *(*Tulg)(unsafe.Pointer(v40))
			*(*Tulg)(unsafe.Pointer(v40))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v39))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(24) & libc.Uint64FromInt32(0xff))
			v42 = s + 40
			v41 = *(*Tulg)(unsafe.Pointer(v42))
			*(*Tulg)(unsafe.Pointer(v42))++
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == int32(9) {
				v43 = int32(2)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy >= int32(m_Z_HUFFMAN_ONLY) || (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(2) {
					v44 = int32(4)
				} else {
					v44 = 0
				}
				v43 = v44
			}
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v41))) = libc.Uint8FromInt32(v43)
			v46 = s + 40
			v45 = *(*Tulg)(unsafe.Pointer(v46))
			*(*Tulg)(unsafe.Pointer(v46))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v45))) = libc.Uint8FromInt32((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fos & libc.Int32FromInt32(0xff))
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
				v48 = s + 40
				v47 = *(*Tulg)(unsafe.Pointer(v48))
				*(*Tulg)(unsafe.Pointer(v48))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v47))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len & libc.Uint32FromInt32(0xff))
				v50 = s + 40
				v49 = *(*Tulg)(unsafe.Pointer(v50))
				*(*Tulg)(unsafe.Pointer(v50))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v49))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf, uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending))
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint64(0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_EXTRA_STATE)
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_EXTRA_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
			beg = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending /* start of bytes to update crc */
			left = uint32(uint64((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len&libc.Uint32FromInt32(0xffff)) - (*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex)
			for (*Tdeflate_state)(unsafe.Pointer(s)).Fpending+uint64(left) > (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
				copy1 = uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fpending)
				libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex), uint64(copy1))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fpending = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size
				if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg {
					(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg), uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg))
				}
				*(*Tulg)(unsafe.Pointer(s + 64)) += uint64(copy1)
				_flush_pending(tls, strm)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
					return m_Z_OK
				}
				beg = uint64(0)
				left -= copy1
			}
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex), uint64(left))
			*(*Tulg)(unsafe.Pointer(s + 40)) += uint64(left)
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg), uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg))
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint64(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_NAME_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_NAME_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname != uintptr(m_Z_NULL) {
			beg1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
			for cond := true; cond; cond = val != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
					if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg1 {
						(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg1), uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg1))
					}
					_flush_pending(tls, strm)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
						(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
						return m_Z_OK
					}
					beg1 = uint64(0)
				}
				v52 = s + 64
				v51 = *(*Tulg)(unsafe.Pointer(v52))
				*(*Tulg)(unsafe.Pointer(v52))++
				val = libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname + uintptr(v51))))
				v54 = s + 40
				v53 = *(*Tulg)(unsafe.Pointer(v54))
				*(*Tulg)(unsafe.Pointer(v54))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v53))) = libc.Uint8FromInt32(val)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg1 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg1), uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg1))
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint64(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_COMMENT_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_COMMENT_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment != uintptr(m_Z_NULL) {
			beg2 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
			for cond := true; cond; cond = val1 != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
					if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg2 {
						(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg2), uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg2))
					}
					_flush_pending(tls, strm)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
						(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
						return m_Z_OK
					}
					beg2 = uint64(0)
				}
				v56 = s + 64
				v55 = *(*Tulg)(unsafe.Pointer(v56))
				*(*Tulg)(unsafe.Pointer(v56))++
				val1 = libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment + uintptr(v55))))
				v58 = s + 40
				v57 = *(*Tulg)(unsafe.Pointer(v58))
				*(*Tulg)(unsafe.Pointer(v58))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v57))) = libc.Uint8FromInt32(val1)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg2 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg2), uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg2))
			}
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_HCRC_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_HCRC_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending+uint64(2) > (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
				_flush_pending(tls, strm)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
					return m_Z_OK
				}
			}
			v60 = s + 40
			v59 = *(*Tulg)(unsafe.Pointer(v60))
			*(*Tulg)(unsafe.Pointer(v60))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v59))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler & libc.Uint64FromInt32(0xff))
			v62 = s + 40
			v61 = *(*Tulg)(unsafe.Pointer(v62))
			*(*Tulg)(unsafe.Pointer(v62))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v61))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(8) & libc.Uint64FromInt32(0xff))
			(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
		/* Compression must start with an empty pending buffer */
		_flush_pending(tls, strm)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
	}
	/* Start a new block or continue the current one.
	 */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) || (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead != uint32(0) || flush != m_Z_NO_FLUSH && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_FINISH_STATE) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == 0 {
			v63 = _deflate_stored(tls, s, flush)
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_HUFFMAN_ONLY) {
				v64 = _deflate_huff(tls, s, flush)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_RLE) {
					v65 = _deflate_rle(tls, s, flush)
				} else {
					v65 = (*(*func(*libc.TLS, uintptr, int32) Tblock_state)(unsafe.Pointer(&struct{ uintptr }{_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Ffunc1})))(tls, s, flush)
				}
				v64 = v65
			}
			v63 = v64
		}
		bstate = v63
		if bstate == int32(_finish_started) || bstate == int32(_finish_done) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_FINISH_STATE)
		}
		if bstate == int32(_need_more) || bstate == int32(_finish_started) {
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1) /* avoid BUF_ERROR next call, see above */
			}
			return m_Z_OK
			/* If flush != Z_NO_FLUSH && avail_out == 0, the next call
			 * of deflate should use the same flush parameter to make sure
			 * that the flush is complete. So we don't have to output an
			 * empty block here, this will be done at next call. This also
			 * ensures that for a very small output buffer, we emit at most
			 * one empty block.
			 */
		}
		if bstate == int32(_block_done) {
			if flush == int32(m_Z_PARTIAL_FLUSH) {
				x__tr_align(tls, s)
			} else {
				if flush != int32(m_Z_BLOCK) { /* FULL_FLUSH or SYNC_FLUSH */
					x__tr_stored_block(tls, s, libc.UintptrFromInt32(0), uint64(0), 0)
					/* For a full flush, this empty block will be recognized
					 * as a special marker by inflate_sync().
					 */
					if flush == int32(m_Z_FULL_FLUSH) {
						*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
						libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint64(2)) /* forget history */
						if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
							(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = uint32(0)
							(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = 0
							(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
						}
					}
				}
			}
			_flush_pending(tls, strm)
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1) /* avoid BUF_ERROR at next call, see above */
				return m_Z_OK
			}
		}
	}
	if flush != int32(m_Z_FINISH) {
		return m_Z_OK
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap <= 0 {
		return int32(m_Z_STREAM_END)
	}
	/* Write the trailer */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == int32(2) {
		v67 = s + 40
		v66 = *(*Tulg)(unsafe.Pointer(v67))
		*(*Tulg)(unsafe.Pointer(v67))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v66))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler & libc.Uint64FromInt32(0xff))
		v69 = s + 40
		v68 = *(*Tulg)(unsafe.Pointer(v69))
		*(*Tulg)(unsafe.Pointer(v69))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v68))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(8) & libc.Uint64FromInt32(0xff))
		v71 = s + 40
		v70 = *(*Tulg)(unsafe.Pointer(v71))
		*(*Tulg)(unsafe.Pointer(v71))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v70))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(16) & libc.Uint64FromInt32(0xff))
		v73 = s + 40
		v72 = *(*Tulg)(unsafe.Pointer(v73))
		*(*Tulg)(unsafe.Pointer(v73))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v72))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(24) & libc.Uint64FromInt32(0xff))
		v75 = s + 40
		v74 = *(*Tulg)(unsafe.Pointer(v75))
		*(*Tulg)(unsafe.Pointer(v75))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v74))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in & libc.Uint64FromInt32(0xff))
		v77 = s + 40
		v76 = *(*Tulg)(unsafe.Pointer(v77))
		*(*Tulg)(unsafe.Pointer(v77))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v76))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(8) & libc.Uint64FromInt32(0xff))
		v79 = s + 40
		v78 = *(*Tulg)(unsafe.Pointer(v79))
		*(*Tulg)(unsafe.Pointer(v79))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v78))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(16) & libc.Uint64FromInt32(0xff))
		v81 = s + 40
		v80 = *(*Tulg)(unsafe.Pointer(v81))
		*(*Tulg)(unsafe.Pointer(v81))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v80))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(24) & libc.Uint64FromInt32(0xff))
	} else {
		_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler>>libc.Int32FromInt32(16)))
		_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler&libc.Uint64FromInt32(0xffff)))
	}
	_flush_pending(tls, strm)
	/* If avail_out is zero, the application will call deflate again
	 * to flush the rest.
	 */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap > 0 {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = -(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap
	} /* write the trailer only once! */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint64(0) {
		v82 = m_Z_OK
	} else {
		v82 = int32(m_Z_STREAM_END)
	}
	return v82
}

// C documentation
//
//	/* ========================================================================= */
func XdeflateEnd(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var status, v1 int32
	_, _ = status, v1
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	status = (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fstatus
	/* Deallocate in reverse order of allocations: */
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending_buf != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending_buf)
	}
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fhead != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fhead)
	}
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fprev != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fprev)
	}
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwindow != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwindow)
	}
	(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	if status == int32(m_BUSY_STATE) {
		v1 = -int32(3)
	} else {
		v1 = m_Z_OK
	}
	return v1
}

// C documentation
//
//	/* =========================================================================
//	 * Copy the source state to the destination state.
//	 * To simplify the source, this is not supported for 16-bit MSDOS (which
//	 * doesn't have enough memory anyway to duplicate compression states).
//	 */
func XdeflateCopy(tls *libc.TLS, dest Tz_streamp, source Tz_streamp) (r int32) {
	var ds, ss uintptr
	_, _ = ds, ss
	if _deflateStateCheck(tls, source) != 0 || dest == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	ss = (*Tz_stream)(unsafe.Pointer(source)).Fstate
	libc.Xmemcpy(tls, dest, source, uint64(112))
	ds = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), uint32(libc.Uint64FromInt64(5952)))
	if ds == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(dest)).Fstate = ds
	libc.Xmemcpy(tls, ds, ss, uint64(5952))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fstrm = dest
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size, uint32(libc.Uint64FromInt32(2)*libc.Uint64FromInt64(1)))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fprev = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size, uint32(libc.Uint64FromInt64(2)))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fhead = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhash_size, uint32(libc.Uint64FromInt64(2)))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize, libc.Uint32FromInt32(libc.Int32FromInt32(m_LIT_BUFS)))
	if (*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fprev == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fhead == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf == uintptr(m_Z_NULL) {
		XdeflateEnd(tls, dest)
		return -int32(4)
	}
	/* following zmemcpy do not work for 16-bit MSDOS */
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(ss)).Fwindow, uint64((*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size*uint32(2))*uint64(1))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fprev, (*Tdeflate_state)(unsafe.Pointer(ss)).Fprev, uint64((*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size)*uint64(2))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhead, (*Tdeflate_state)(unsafe.Pointer(ss)).Fhead, uint64((*Tdeflate_state)(unsafe.Pointer(ds)).Fhash_size)*uint64(2))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_buf, uint64((*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize*uint32(m_LIT_BUFS)))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf + uintptr(int64((*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_out)-int64((*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_buf))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fsym_buf = (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize)
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fl_desc.Fdyn_tree = ds + 212
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fd_desc.Fdyn_tree = ds + 2504
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fbl_desc.Fdyn_tree = ds + 2748
	return m_Z_OK
}

// C documentation
//
//	/* ===========================================================================
//	 * Set match_start to the longest match starting at the given string and
//	 * return its length. Matches shorter or equal to prev_length are discarded,
//	 * in which case the result is equal to prev_length and match_start is
//	 * garbage.
//	 * IN assertions: cur_match is the head of the hash chain for the current
//	 *   string (strstart) and its distance is <= MAX_DIST, and prev_length >= 1
//	 * OUT assertion: the match length is not greater than s->lookahead.
//	 */
func _longest_match(tls *libc.TLS, s uintptr, cur_match TIPos) (r TuInt) {
	var best_len, len1, nice_match int32
	var chain_length, v1, v3 uint32
	var limit, v2 TIPos
	var match, prev, scan, strend, v10, v11, v13, v14, v16, v17, v19, v20, v22, v23, v25, v26, v28, v29, v6, v8, v9 uintptr
	var scan_end, scan_end1 TByte
	var wmask TuInt
	var v12, v15, v18, v21, v24, v27, v30, v4, v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = best_len, chain_length, len1, limit, match, nice_match, prev, scan, scan_end, scan_end1, strend, wmask, v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v4, v6, v7, v8, v9
	chain_length = (*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length                                         /* max hash chain length */
	scan = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) /* length of current match */
	best_len = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length)                            /* best match length so far */
	nice_match = (*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - ((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)))
	} else {
		v1 = uint32(m_NIL)
	} /* stop if match long enough */
	limit = v1
	/* Stop when cur_match becomes <= limit. To simplify the code,
	 * we prevent matches with the string of window index 0.
	 */
	prev = (*Tdeflate_state)(unsafe.Pointer(s)).Fprev
	wmask = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask
	strend = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) + uintptr(m_MAX_MATCH)
	scan_end1 = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len-int32(1))))
	scan_end = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len)))
	/* The code is optimized for HASH_BITS >= 8 and MAX_MATCH-2 multiple of 16.
	 * It is easy to get rid of this optimization if necessary.
	 */
	/* Do not waste too much time if we already have a good match: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length >= (*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match {
		chain_length >>= uint32(2)
	}
	/* Do not look for matches beyond the end of the input. This is necessary
	 * to make deflate deterministic.
	 */
	if libc.Uint32FromInt32(nice_match) > (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
		nice_match = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead)
	}
	for {
		match = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(cur_match)
		/* Skip to next match if the match length cannot increase
		 * or if the match length is less than 2.  Note that the checks below
		 * for insufficient lookahead only occur occasionally for performance
		 * reasons.  Therefore uninitialized memory will be accessed, and
		 * conditional jumps will be made that depend on those values.
		 * However the length of the match is limited to the lookahead, so
		 * the output of deflate is not affected by the uninitialized values.
		 */
		if v7 = libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(match + uintptr(best_len)))) != libc.Int32FromUint8(scan_end) || libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(match + uintptr(best_len-int32(1))))) != libc.Int32FromUint8(scan_end1) || libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(match))) != libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(scan))); !v7 {
			match++
			v6 = match
		}
		if v7 || libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v6))) != libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(scan + 1))) {
			goto _5
		}
		/* The check at best_len - 1 can be removed because it will be made
		 * again later. (This heuristic is not always a win.)
		 * It is not necessary to compare scan[2] and match[2] since they
		 * are always equal when the other bytes match, given that
		 * the hash keys are equal and that HASH_BITS >= 8.
		 */
		scan += uintptr(2)
		/* The check at best_len - 1 can be removed because it will be made
		 * again later. (This heuristic is not always a win.)
		 * It is not necessary to compare scan[2] and match[2] since they
		 * are always equal when the other bytes match, given that
		 * the hash keys are equal and that HASH_BITS >= 8.
		 */
		match++
		/* We check for insufficient lookahead only every 8th comparison;
		 * the 256th check will be made at strstart + 258.
		 */
		for {
			goto _31
		_31:
			;
			scan++
			v8 = scan
			match++
			v9 = match
			if v12 = libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v8))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v9))); v12 {
				scan++
				v10 = scan
				match++
				v11 = match
			}
			if v15 = v12 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v10))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v11))); v15 {
				scan++
				v13 = scan
				match++
				v14 = match
			}
			if v18 = v15 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v13))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v14))); v18 {
				scan++
				v16 = scan
				match++
				v17 = match
			}
			if v21 = v18 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v16))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v17))); v21 {
				scan++
				v19 = scan
				match++
				v20 = match
			}
			if v24 = v21 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v19))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v20))); v24 {
				scan++
				v22 = scan
				match++
				v23 = match
			}
			if v27 = v24 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v22))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v23))); v27 {
				scan++
				v25 = scan
				match++
				v26 = match
			}
			if v30 = v27 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v25))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v26))); v30 {
				scan++
				v28 = scan
				match++
				v29 = match
			}
			if !(v30 && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v28))) == libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer(v29))) && scan < strend) {
				break
			}
		}
		len1 = int32(m_MAX_MATCH) - int32(int64(strend)-int64(scan))
		scan = strend - uintptr(m_MAX_MATCH)
		if len1 > best_len {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start = cur_match
			best_len = len1
			if len1 >= nice_match {
				break
			}
			scan_end1 = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len-int32(1))))
			scan_end = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len)))
		}
		goto _5
	_5:
		;
		v2 = uint32(*(*TPosf)(unsafe.Pointer(prev + uintptr(cur_match&wmask)*2)))
		cur_match = v2
		if v4 = v2 > limit; v4 {
			chain_length--
			v3 = chain_length
		}
		if !(v4 && v3 != uint32(0)) {
			break
		}
	}
	if libc.Uint32FromInt32(best_len) <= (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
		return libc.Uint32FromInt32(best_len)
	}
	return (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
}

/* ===========================================================================
 * Flush the current block, with given end-of-file flag.
 * IN assertion: strstart is set to the end of the current match.
 */

/* Same but force premature exit if necessary. */

/* Maximum stored block length in deflate format (not including header). */

/* Minimum of a and b. */

// C documentation
//
//	/* ===========================================================================
//	 * Copy without compression as much as possible from the input stream, return
//	 * the current block state.
//	 *
//	 * In case deflateParams() is used to later switch to a non-zero compression
//	 * level, s->matches (otherwise unused when storing) keeps track of the number
//	 * of hash table slides to perform. If s->matches is 1, then one hash table
//	 * slide will be done when switching. If s->matches is 2, the maximum value
//	 * allowed here, then the hash table will be cleared, since two or more slides
//	 * is the same as a clear.
//	 *
//	 * deflate_stored() is written to minimize the number of times an input byte is
//	 * copied. It is most efficient with large input and output buffers, which
//	 * maximizes the opportunities to have a single copy from next_in to next_out.
//	 */
func _deflate_stored(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var have, last, left, len1, min_block, used, v3, v4, v6, v7 uint32
	var v1, v5 uint64
	var v2, v8, v9 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = have, last, left, len1, min_block, used, v1, v2, v3, v4, v5, v6, v7, v8, v9
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size-uint64(5) > uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size) {
		v1 = uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
	} else {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - uint64(5)
	}
	/* Smallest worthy block size when not flushing or finishing. By default
	 * this is 32K. This can be as small as 507 bytes for memLevel == 1. For
	 * large input and output buffers, the stored block size will be larger.
	 */
	min_block = uint32(v1)
	last = uint32(0)
	used = (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in
	for cond := true; cond; cond = last == uint32(0) {
		/* Set len to the maximum size block that we can copy directly with the
		 * available input data and output space. Set left to how much of that
		 * would be copied from what's left in the window.
		 */
		len1 = uint32(m_MAX_STORED)                                                                           /* maximum deflate stored block length */
		have = libc.Uint32FromInt32(((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid + int32(42)) >> int32(3)) /* number of header bytes */
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out < have {       /* need room for header */
			break
		}
		/* maximum stored block length that will fit in avail_out: */
		have = (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out - have
		left = libc.Uint32FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) - (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start) /* bytes left in window */
		if uint64(len1) > uint64(left)+uint64((*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in) {
			len1 = left + (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in
		} /* limit len to the input */
		if len1 > have {
			len1 = have
		} /* limit len to the output */
		/* If the stored block would be less than min_block in length, or if
		 * unable to copy all of the available input when flushing, then try
		 * copying to the window and the pending buffer instead. Also don't
		 * write an empty block when flushing -- deflate() does that.
		 */
		if len1 < min_block && (len1 == uint32(0) && flush != int32(m_Z_FINISH) || flush == m_Z_NO_FLUSH || len1 != left+(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in) {
			break
		}
		/* Make a dummy stored block in pending to get the header bytes,
		 * including any pending bits. This also updates the debugging counts.
		 */
		if flush == int32(m_Z_FINISH) && len1 == left+(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in {
			v2 = int32(1)
		} else {
			v2 = 0
		}
		last = libc.Uint32FromInt32(v2)
		x__tr_stored_block(tls, s, libc.UintptrFromInt32(0), uint64(0), libc.Int32FromUint32(last))
		/* Replace the lengths in the dummy stored block with len. */
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint64(4)))) = uint8(len1)
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint64(3)))) = uint8(len1 >> int32(8))
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint64(2)))) = uint8(^len1)
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint64(1)))) = uint8(^len1 >> int32(8))
		/* Write the stored block header bytes. */
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		/* Copy uncompressed bytes from the window to next_out. */
		if left != 0 {
			if left > len1 {
				left = len1
			}
			libc.Xmemcpy(tls, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_out, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), uint64(left))
			*(*uintptr)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 24)) += uintptr(left)
			*(*TuInt)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 32)) -= left
			*(*TuLong)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 40)) += uint64(left)
			*(*int64)(unsafe.Pointer(s + 152)) += libc.Int64FromUint32(left)
			len1 -= left
		}
		/* Copy uncompressed bytes directly from next_in to next_out, updating
		 * the check value.
		 */
		if len1 != 0 {
			_read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_out, len1)
			*(*uintptr)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 24)) += uintptr(len1)
			*(*TuInt)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 32)) -= len1
			*(*TuLong)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 40)) += uint64(len1)
		}
	}
	/* Update the sliding window with the last s->w_size bytes of the copied
	 * data, or append all of the copied data to the existing window if less
	 * than s->w_size bytes were copied. Also update the number of bytes to
	 * insert in the hash tables, in the event that deflateParams() switches to
	 * a non-zero compression level.
	 */
	used -= (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in /* number of input bytes directly copied */
	if used != 0 {
		/* If any input was used, then no unused input remains in the window,
		 * therefore s->block_start == s->strstart.
		 */
		if used >= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size { /* supplant the previous history */
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches = uint32(2) /* clear hash */
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_in-uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
			(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size-uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) <= uint64(used) {
				/* Slide the window down. */
				*(*TuInt)(unsafe.Pointer(s + 172)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
				libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches < uint32(2) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
				} /* add a pending slide_hash() */
				if (*Tdeflate_state)(unsafe.Pointer(s)).Finsert > (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
					(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
				}
			}
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart), (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_in-uintptr(used), uint64(used))
			*(*TuInt)(unsafe.Pointer(s + 172)) += used
			if used > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-(*Tdeflate_state)(unsafe.Pointer(s)).Finsert {
				v3 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
			} else {
				v3 = used
			}
			*(*TuInt)(unsafe.Pointer(s + 5932)) += v3
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
	}
	/* If the last block was written to next_out, then done. */
	if last != 0 {
		return int32(_finish_done)
	}
	/* If flushing and all input has been consumed, then done. */
	if flush != m_Z_NO_FLUSH && flush != int32(m_Z_FINISH) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) == (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start {
		return int32(_block_done)
	}
	/* Fill the window with any remaining input. */
	have = uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart))
	if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in > have && (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size) {
		/* Slide the window down. */
		*(*int64)(unsafe.Pointer(s + 152)) -= libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
		*(*TuInt)(unsafe.Pointer(s + 172)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
		libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart))
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches < uint32(2) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
		} /* add a pending slide_hash() */
		have += (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size /* more space now */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Finsert > (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
			(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		}
	}
	if have > (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in {
		have = (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in
	}
	if have != 0 {
		_read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart), have)
		*(*TuInt)(unsafe.Pointer(s + 172)) += have
		if have > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-(*Tdeflate_state)(unsafe.Pointer(s)).Finsert {
			v4 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
		} else {
			v4 = have
		}
		*(*TuInt)(unsafe.Pointer(s + 5932)) += v4
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = uint64((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
	}
	/* There was not enough avail_out to write a complete worthy or flushed
	 * stored block to next_out. Write a stored block to pending instead, if we
	 * have enough input for a worthy block, or if flushing and there is enough
	 * room for the remaining input as a stored block in the pending buffer.
	 */
	have = libc.Uint32FromInt32(((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid + int32(42)) >> int32(3)) /* number of header bytes */
	/* maximum stored block length that will fit in pending: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size-uint64(have) > libc.Uint64FromInt32(libc.Int32FromInt32(m_MAX_STORED)) {
		v5 = libc.Uint64FromInt32(libc.Int32FromInt32(m_MAX_STORED))
	} else {
		v5 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - uint64(have)
	}
	have = uint32(v5)
	if have > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		v6 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	} else {
		v6 = have
	}
	min_block = v6
	left = libc.Uint32FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) - (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start)
	if left >= min_block || (left != 0 || flush == int32(m_Z_FINISH)) && flush != m_Z_NO_FLUSH && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && left <= have {
		if left > have {
			v7 = have
		} else {
			v7 = left
		}
		len1 = v7
		if flush == int32(m_Z_FINISH) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && len1 == left {
			v8 = int32(1)
		} else {
			v8 = 0
		}
		last = libc.Uint32FromInt32(v8)
		x__tr_stored_block(tls, s, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), uint64(len1), libc.Int32FromUint32(last))
		*(*int64)(unsafe.Pointer(s + 152)) += libc.Int64FromUint32(len1)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
	}
	/* We've done all we can with the available input and output. */
	if last != 0 {
		v9 = int32(_finish_started)
	} else {
		v9 = int32(_need_more)
	}
	return v9
}

// C documentation
//
//	/* ===========================================================================
//	 * Compress as much as possible from the input stream, return the current
//	 * block state.
//	 * This function does not perform lazy evaluation of matches and inserts
//	 * new strings in the dictionary only for unmatched strings or for short
//	 * matches. It is used only for the fast compression options.
//	 */
func _deflate_fast(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush, v9 int32
	var cc, len1 Tuch
	var dist Tush
	var hash_head TIPos
	var v10, v14, v16, v18, v3, v5, v7 TuInt
	var v11, v15, v17, v19, v20, v22, v23, v4, v6, v8 uintptr
	var v13, v2 TPosf
	var v21 uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bflush, cc, dist, hash_head, len1, v10, v11, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v3, v4, v5, v6, v7, v8, v9 /* set if current block must be flushed */
	for {
		/* Make sure that we always have enough lookahead, except
		 * at the end of the input file. We need MAX_MATCH bytes
		 * for the next match, plus MIN_MATCH bytes to insert the
		 * string following the next match.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) && flush == m_Z_NO_FLUSH {
				return int32(_need_more)
			}
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				break
			} /* flush the current block */
		}
		/* Insert the string window[strstart .. strstart + 2] in the
		 * dictionary, and set hash_head to the head of the hash chain:
		 */
		hash_head = uint32(m_NIL)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			v2 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v2
			hash_head = uint32(v2)
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		}
		/* Find the longest match, discarding those <= prev_length.
		 * At this point we have always match_length < MIN_MATCH
		 */
		if hash_head != uint32(m_NIL) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-hash_head <= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			/* To simplify the code, we prevent matches with the string
			 * of window index 0 (in particular we have to avoid a match
			 * of the string with itself at the start of the input file).
			 */
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = _longest_match(tls, s, hash_head)
			/* longest_match() sets match_start */
		}
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length >= uint32(m_MIN_MATCH) {
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start)
			v4 = s + 5900
			v3 = *(*TuInt)(unsafe.Pointer(v4))
			*(*TuInt)(unsafe.Pointer(v4))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))) = uint8(dist)
			v6 = s + 5900
			v5 = *(*TuInt)(unsafe.Pointer(v6))
			*(*TuInt)(unsafe.Pointer(v6))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16(dist) >> libc.Int32FromInt32(8))
			v8 = s + 5900
			v7 = *(*TuInt)(unsafe.Pointer(v8))
			*(*TuInt)(unsafe.Pointer(v8))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v7))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 212 + uintptr(libc.Int32FromUint8(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if libc.Int32FromUint16(dist) < int32(256) {
				v9 = libc.Int32FromUint8(x__dist_code[dist])
			} else {
				v9 = libc.Int32FromUint8(x__dist_code[int32(256)+libc.Int32FromUint16(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2504 + uintptr(v9)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			*(*TuInt)(unsafe.Pointer(s + 180)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			/* Insert new strings in the hash table only if the match length
			 * is not too large. This saves time but degrades compression.
			 */
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length <= (*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match && (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length-- /* string at strstart already in table */
				for {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
					(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
					v13 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v13
					hash_head = uint32(v13)
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
					/* strstart never exceeds WSIZE-MAX_MATCH, so there are
					 * always MIN_MATCH bytes ahead.
					 */
					goto _12
				_12:
					;
					v11 = s + 160
					*(*TuInt)(unsafe.Pointer(v11))--
					v10 = *(*TuInt)(unsafe.Pointer(v11))
					if !(v10 != uint32(0)) {
						break
					}
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
			} else {
				*(*TuInt)(unsafe.Pointer(s + 172)) += (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart))))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
				/* If lookahead < MIN_MATCH, ins_h is garbage, but it does not
				 * matter since it will be recomputed at next deflate call.
				 */
			}
		} else {
			/* No match, output a literal byte */
			cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
			v15 = s + 5900
			v14 = *(*TuInt)(unsafe.Pointer(v15))
			*(*TuInt)(unsafe.Pointer(v15))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v14))) = uint8(0)
			v17 = s + 5900
			v16 = *(*TuInt)(unsafe.Pointer(v17))
			*(*TuInt)(unsafe.Pointer(v17))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v16))) = uint8(0)
			v19 = s + 5900
			v18 = *(*TuInt)(unsafe.Pointer(v19))
			*(*TuInt)(unsafe.Pointer(v19))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v18))) = cc
			*(*Tush)(unsafe.Pointer(s + 212 + uintptr(cc)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		}
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v20 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v20 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v20, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
			_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
			if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
				return int32(_need_more)
			}
		}
		goto _1
	_1:
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart < libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1)) {
		v21 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	} else {
		v21 = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = v21
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v22 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v22 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v22, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v23 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v23 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v23, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

// C documentation
//
//	/* ===========================================================================
//	 * Same as above, but achieves better compression. We use a lazy
//	 * evaluation for matches: a match is finally adopted only if there is
//	 * no better match at the next window position.
//	 */
func _deflate_slow(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush, v9 int32
	var cc, cc1, len1 Tuch
	var dist Tush
	var hash_head TIPos
	var max_insert, v10, v13, v17, v19, v21, v24, v26, v28, v3, v5, v7 TuInt
	var v11, v14, v16, v18, v20, v22, v23, v25, v27, v29, v31, v32, v4, v6, v8 uintptr
	var v15, v2 TPosf
	var v30 uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bflush, cc, cc1, dist, hash_head, len1, max_insert, v10, v11, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v4, v5, v6, v7, v8, v9 /* set if current block must be flushed */
	/* Process the input block. */
	for {
		/* Make sure that we always have enough lookahead, except
		 * at the end of the input file. We need MAX_MATCH bytes
		 * for the next match, plus MIN_MATCH bytes to insert the
		 * string following the next match.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) && flush == m_Z_NO_FLUSH {
				return int32(_need_more)
			}
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				break
			} /* flush the current block */
		}
		/* Insert the string window[strstart .. strstart + 2] in the
		 * dictionary, and set hash_head to the head of the hash chain:
		 */
		hash_head = uint32(m_NIL)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			v2 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v2
			hash_head = uint32(v2)
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		}
		/* Find the longest match, discarding those <= prev_length.
		 */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length = (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
		/* Find the longest match, discarding those <= prev_length.
		 */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_match = (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
		if hash_head != uint32(m_NIL) && (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length < (*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-hash_head <= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			/* To simplify the code, we prevent matches with the string
			 * of window index 0 (in particular we have to avoid a match
			 * of the string with itself at the start of the input file).
			 */
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = _longest_match(tls, s, hash_head)
			/* longest_match() sets match_start */
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length <= uint32(5) && ((*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_FILTERED) || (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length == uint32(m_MIN_MATCH) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start > uint32(m_TOO_FAR)) {
				/* If prev_match is also MIN_MATCH, match_start is garbage
				 * but we will ignore the current match anyway.
				 */
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
			}
		}
		/* If there was a match at the previous step and the current
		 * match is not better, output the previous match:
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length >= uint32(m_MIN_MATCH) && (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length <= (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length {
			max_insert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart + (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead - uint32(m_MIN_MATCH)
			/* Do not insert strings in hash table beyond this. */
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - libc.Uint32FromInt32(1) - (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_match)
			v4 = s + 5900
			v3 = *(*TuInt)(unsafe.Pointer(v4))
			*(*TuInt)(unsafe.Pointer(v4))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))) = uint8(dist)
			v6 = s + 5900
			v5 = *(*TuInt)(unsafe.Pointer(v6))
			*(*TuInt)(unsafe.Pointer(v6))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16(dist) >> libc.Int32FromInt32(8))
			v8 = s + 5900
			v7 = *(*TuInt)(unsafe.Pointer(v8))
			*(*TuInt)(unsafe.Pointer(v8))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v7))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 212 + uintptr(libc.Int32FromUint8(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if libc.Int32FromUint16(dist) < int32(256) {
				v9 = libc.Int32FromUint8(x__dist_code[dist])
			} else {
				v9 = libc.Int32FromUint8(x__dist_code[int32(256)+libc.Int32FromUint16(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2504 + uintptr(v9)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			/* Insert in hash table all strings up to the end of the match.
			 * strstart - 1 and strstart are already inserted. If there is not
			 * enough lookahead, the last two strings are not inserted in
			 * the hash table.
			 */
			*(*TuInt)(unsafe.Pointer(s + 180)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length - uint32(1)
			*(*TuInt)(unsafe.Pointer(s + 184)) -= uint32(2)
			for {
				v14 = s + 172
				*(*TuInt)(unsafe.Pointer(v14))++
				v13 = *(*TuInt)(unsafe.Pointer(v14))
				if v13 <= max_insert {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
					v15 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v15
					hash_head = uint32(v15)
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
				}
				goto _12
			_12:
				;
				v11 = s + 184
				*(*TuInt)(unsafe.Pointer(v11))--
				v10 = *(*TuInt)(unsafe.Pointer(v11))
				if !(v10 != uint32(0)) {
					break
				}
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
			if bflush != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
					v16 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
				} else {
					v16 = libc.UintptrFromInt32(m_Z_NULL)
				}
				x__tr_flush_block(tls, s, v16, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
				_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
				if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
					return int32(_need_more)
				}
			}
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available != 0 {
				/* If there was no match at the previous position, output a
				 * single literal. If there was a match but the current match
				 * is longer, truncate the previous match to a single literal.
				 */
				cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-uint32(1))))
				v18 = s + 5900
				v17 = *(*TuInt)(unsafe.Pointer(v18))
				*(*TuInt)(unsafe.Pointer(v18))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v17))) = uint8(0)
				v20 = s + 5900
				v19 = *(*TuInt)(unsafe.Pointer(v20))
				*(*TuInt)(unsafe.Pointer(v20))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v19))) = uint8(0)
				v22 = s + 5900
				v21 = *(*TuInt)(unsafe.Pointer(v22))
				*(*TuInt)(unsafe.Pointer(v22))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v21))) = cc
				*(*Tush)(unsafe.Pointer(s + 212 + uintptr(cc)*4))++
				bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
				if bflush != 0 {
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
						v23 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
					} else {
						v23 = libc.UintptrFromInt32(m_Z_NULL)
					}
					x__tr_flush_block(tls, s, v23, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
					(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
					_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
				(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
				if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
					return int32(_need_more)
				}
			} else {
				/* There is no previous match to compare with, wait for
				 * the next step to decide.
				 */
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = int32(1)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
				(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			}
		}
		goto _1
	_1:
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available != 0 {
		cc1 = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-uint32(1))))
		v25 = s + 5900
		v24 = *(*TuInt)(unsafe.Pointer(v25))
		*(*TuInt)(unsafe.Pointer(v25))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v24))) = uint8(0)
		v27 = s + 5900
		v26 = *(*TuInt)(unsafe.Pointer(v27))
		*(*TuInt)(unsafe.Pointer(v27))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v26))) = uint8(0)
		v29 = s + 5900
		v28 = *(*TuInt)(unsafe.Pointer(v29))
		*(*TuInt)(unsafe.Pointer(v29))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v28))) = cc1
		*(*Tush)(unsafe.Pointer(s + 212 + uintptr(cc1)*4))++
		bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart < libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1)) {
		v30 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	} else {
		v30 = libc.Uint32FromInt32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = v30
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v31 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v31 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v31, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v32 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v32 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v32, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

// C documentation
//
//	/* ===========================================================================
//	 * For Z_RLE, simply look for runs of bytes, generate matches only of distance
//	 * one.  Do not maintain a hash table.  (It will be regenerated if this run of
//	 * deflate switches away from Z_RLE.)
//	 */
func _deflate_rle(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush, v29 int32
	var cc, len1 Tuch
	var dist Tush
	var prev, v23, v25, v27, v30, v32, v34 TuInt
	var scan, strend, v10, v12, v14, v16, v18, v2, v20, v24, v26, v28, v3, v31, v33, v35, v36, v37, v38, v5, v7, v8 uintptr
	var v11, v13, v15, v17, v19, v21, v4, v6, v9 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bflush, cc, dist, len1, prev, scan, strend, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v4, v5, v6, v7, v8, v9 /* scan goes up to strend for length of run */
	for {
		/* Make sure that we always have enough lookahead, except
		 * at the end of the input file. We need MAX_MATCH bytes
		 * for the longest run, plus one for the unrolled loop.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead <= uint32(m_MAX_MATCH) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead <= uint32(m_MAX_MATCH) && flush == m_Z_NO_FLUSH {
				return int32(_need_more)
			}
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				break
			} /* flush the current block */
		}
		/* See how many times the previous byte repeats */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart > uint32(0) {
			scan = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) - uintptr(1)
			prev = uint32(*(*TBytef)(unsafe.Pointer(scan)))
			scan++
			v2 = scan
			if v4 = prev == uint32(*(*TBytef)(unsafe.Pointer(v2))); v4 {
				scan++
				v3 = scan
			}
			if v6 = v4 && prev == uint32(*(*TBytef)(unsafe.Pointer(v3))); v6 {
				scan++
				v5 = scan
			}
			if v6 && prev == uint32(*(*TBytef)(unsafe.Pointer(v5))) {
				strend = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) + uintptr(m_MAX_MATCH)
				for {
					goto _22
				_22:
					;
					scan++
					v7 = scan
					if v9 = prev == uint32(*(*TBytef)(unsafe.Pointer(v7))); v9 {
						scan++
						v8 = scan
					}
					if v11 = v9 && prev == uint32(*(*TBytef)(unsafe.Pointer(v8))); v11 {
						scan++
						v10 = scan
					}
					if v13 = v11 && prev == uint32(*(*TBytef)(unsafe.Pointer(v10))); v13 {
						scan++
						v12 = scan
					}
					if v15 = v13 && prev == uint32(*(*TBytef)(unsafe.Pointer(v12))); v15 {
						scan++
						v14 = scan
					}
					if v17 = v15 && prev == uint32(*(*TBytef)(unsafe.Pointer(v14))); v17 {
						scan++
						v16 = scan
					}
					if v19 = v17 && prev == uint32(*(*TBytef)(unsafe.Pointer(v16))); v19 {
						scan++
						v18 = scan
					}
					if v21 = v19 && prev == uint32(*(*TBytef)(unsafe.Pointer(v18))); v21 {
						scan++
						v20 = scan
					}
					if !(v21 && prev == uint32(*(*TBytef)(unsafe.Pointer(v20))) && scan < strend) {
						break
					}
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(m_MAX_MATCH) - libc.Uint32FromInt64(int64(strend)-int64(scan))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length > (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
				}
			}
		}
		/* Emit match if have run of MIN_MATCH or longer, else emit literal */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length >= uint32(m_MIN_MATCH) {
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = libc.Uint16FromInt32(libc.Int32FromInt32(1))
			v24 = s + 5900
			v23 = *(*TuInt)(unsafe.Pointer(v24))
			*(*TuInt)(unsafe.Pointer(v24))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v23))) = uint8(dist)
			v26 = s + 5900
			v25 = *(*TuInt)(unsafe.Pointer(v26))
			*(*TuInt)(unsafe.Pointer(v26))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v25))) = libc.Uint8FromInt32(libc.Int32FromUint16(dist) >> libc.Int32FromInt32(8))
			v28 = s + 5900
			v27 = *(*TuInt)(unsafe.Pointer(v28))
			*(*TuInt)(unsafe.Pointer(v28))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v27))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 212 + uintptr(libc.Int32FromUint8(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if libc.Int32FromUint16(dist) < int32(256) {
				v29 = libc.Int32FromUint8(x__dist_code[dist])
			} else {
				v29 = libc.Int32FromUint8(x__dist_code[int32(256)+libc.Int32FromUint16(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2504 + uintptr(v29)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			*(*TuInt)(unsafe.Pointer(s + 180)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			*(*TuInt)(unsafe.Pointer(s + 172)) += (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		} else {
			/* No match, output a literal byte */
			cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
			v31 = s + 5900
			v30 = *(*TuInt)(unsafe.Pointer(v31))
			*(*TuInt)(unsafe.Pointer(v31))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v30))) = uint8(0)
			v33 = s + 5900
			v32 = *(*TuInt)(unsafe.Pointer(v33))
			*(*TuInt)(unsafe.Pointer(v33))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v32))) = uint8(0)
			v35 = s + 5900
			v34 = *(*TuInt)(unsafe.Pointer(v35))
			*(*TuInt)(unsafe.Pointer(v35))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v34))) = cc
			*(*Tush)(unsafe.Pointer(s + 212 + uintptr(cc)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		}
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v36 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v36 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v36, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
			_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
			if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
				return int32(_need_more)
			}
		}
		goto _1
	_1:
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v37 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v37 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v37, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v38 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v38 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v38, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

// C documentation
//
//	/* ===========================================================================
//	 * For Z_HUFFMAN_ONLY, do not look for matches.  Do not maintain a hash table.
//	 * (It will be regenerated if this run of deflate switches away from Huffman.)
//	 */
func _deflate_huff(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush int32
	var cc Tuch
	var v10, v3, v5, v7, v8, v9 uintptr
	var v2, v4, v6 TuInt
	_, _, _, _, _, _, _, _, _, _, _ = bflush, cc, v10, v2, v3, v4, v5, v6, v7, v8, v9 /* set if current block must be flushed */
	for {
		/* Make sure that we have a literal to write. */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				if flush == m_Z_NO_FLUSH {
					return int32(_need_more)
				}
				break /* flush the current block */
			}
		}
		/* Output a literal byte */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
		v3 = s + 5900
		v2 = *(*TuInt)(unsafe.Pointer(v3))
		*(*TuInt)(unsafe.Pointer(v3))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v2))) = uint8(0)
		v5 = s + 5900
		v4 = *(*TuInt)(unsafe.Pointer(v5))
		*(*TuInt)(unsafe.Pointer(v5))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v4))) = uint8(0)
		v7 = s + 5900
		v6 = *(*TuInt)(unsafe.Pointer(v7))
		*(*TuInt)(unsafe.Pointer(v7))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v6))) = cc
		*(*Tush)(unsafe.Pointer(s + 212 + uintptr(cc)*4))++
		bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
		(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v8 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v8 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v8, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
			_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
			if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
				return int32(_need_more)
			}
		}
		goto _1
	_1:
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v9 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v9 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v9, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v10 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt64((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v10 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v10, libc.Uint64FromInt64(libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int64FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

const m_ENOUGH_DISTS = 592
const m_ENOUGH_LENS = 852
const m_PRESET_DICT2 = 0x20

type Tcode = struct {
	Fop   uint8
	Fbits uint8
	Fval  uint16
}

type Tcodetype = int32

const _CODES = 0
const _LENS = 1
const _DISTS = 2

type Tinflate_mode = int32

const _HEAD = 16180
const _FLAGS = 16181
const _TIME = 16182
const _OS = 16183
const _EXLEN = 16184
const _EXTRA = 16185
const _NAME = 16186
const _COMMENT = 16187
const _HCRC = 16188
const _DICTID = 16189
const _DICT = 16190
const _TYPE = 16191
const _TYPEDO = 16192
const _STORED = 16193
const _COPY_ = 16194
const _COPY = 16195
const _TABLE = 16196
const _LENLENS = 16197
const _CODELENS = 16198
const _LEN_ = 16199
const _LEN = 16200
const _LENEXT = 16201
const _DIST = 16202
const _DISTEXT = 16203
const _MATCH = 16204
const _LIT = 16205
const _CHECK = 16206
const _LENGTH = 16207
const _DONE = 16208
const _BAD = 16209
const _MEM = 16210
const _SYNC = 16211

type Tinflate_state = struct {
	Fstrm     Tz_streamp
	Fmode     Tinflate_mode
	Flast     int32
	Fwrap     int32
	Fhavedict int32
	Fflags    int32
	Fdmax     uint32
	Fcheck    uint64
	Ftotal    uint64
	Fhead     Tgz_headerp
	Fwbits    uint32
	Fwsize    uint32
	Fwhave    uint32
	Fwnext    uint32
	Fwindow   uintptr
	Fhold     uint64
	Fbits     uint32
	Flength   uint32
	Foffset   uint32
	Fextra    uint32
	Flencode  uintptr
	Fdistcode uintptr
	Flenbits  uint32
	Fdistbits uint32
	Fncode    uint32
	Fnlen     uint32
	Fndist    uint32
	Fhave     uint32
	Fnext     uintptr
	Flens     [320]uint16
	Fwork     [288]uint16
	Fcodes    [1444]Tcode
	Fsane     int32
	Fback     int32
	Fwas      uint32
}

// C documentation
//
//	/*
//	   strm provides memory allocation functions in zalloc and zfree, or
//	   Z_NULL to use the library memory allocation functions.
//
//	   windowBits is in the range 8..15, and window is a user-supplied
//	   window and output buffer that is 2**windowBits bytes.
//	 */
func XinflateBackInit_(tls *libc.TLS, strm Tz_streamp, windowBits int32, window uintptr, version uintptr, stream_size int32) (r int32) {
	var state uintptr
	_ = state
	if version == uintptr(m_Z_NULL) || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(version))) != libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(__ccgo_ts))) || stream_size != libc.Int32FromUint64(libc.Uint64FromInt64(112)) {
		return -int32(6)
	}
	if strm == uintptr(m_Z_NULL) || window == uintptr(m_Z_NULL) || windowBits < int32(8) || windowBits > int32(15) {
		return -int32(2)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* in case we return an error */
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = __ccgo_fp(Xzcalloc)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = libc.UintptrFromInt32(0)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = __ccgo_fp(Xzcfree)
	}
	state = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), uint32(libc.Uint64FromInt64(7160)))
	if state == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = state
	(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(32768)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = libc.Uint32FromInt32(windowBits)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwsize = uint32(1) << windowBits
	(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = window
	(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fsane = int32(1)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Return state with length and distance decoding tables and index sizes set to
//	   fixed code decoding.  Normally this returns fixed tables from inffixed.h.
//	   If BUILDFIXED is defined, then instead this routine builds the tables the
//	   first time it's called, and returns those tables the first time and
//	   thereafter.  This reduces the size of the code by about 2K bytes, in
//	   exchange for a little execution time.  However, BUILDFIXED should not be
//	   used for threaded applications, since the rewriting of the tables and virgin
//	   may not be thread-safe.
//	 */
func _fixedtables(tls *libc.TLS, state uintptr) {
	(*Tinflate_state)(unsafe.Pointer(state)).Flencode = uintptr(unsafe.Pointer(&_lenfix))
	(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = uintptr(unsafe.Pointer(&_distfix))
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(5)
}

var _lenfix = [512]Tcode{
	0: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	1: {
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	2: {
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	3: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	4: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	5: {
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	6: {
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	7: {
		Fbits: uint8(9),
		Fval:  uint16(192),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	9: {
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	10: {
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	11: {
		Fbits: uint8(9),
		Fval:  uint16(160),
	},
	12: {
		Fbits: uint8(8),
	},
	13: {
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	14: {
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	15: {
		Fbits: uint8(9),
		Fval:  uint16(224),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	17: {
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	18: {
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	19: {
		Fbits: uint8(9),
		Fval:  uint16(144),
	},
	20: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	21: {
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	22: {
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	23: {
		Fbits: uint8(9),
		Fval:  uint16(208),
	},
	24: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	25: {
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	26: {
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	27: {
		Fbits: uint8(9),
		Fval:  uint16(176),
	},
	28: {
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	29: {
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	30: {
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	31: {
		Fbits: uint8(9),
		Fval:  uint16(240),
	},
	32: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	33: {
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	34: {
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	35: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	36: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	37: {
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	38: {
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	39: {
		Fbits: uint8(9),
		Fval:  uint16(200),
	},
	40: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	41: {
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	42: {
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	43: {
		Fbits: uint8(9),
		Fval:  uint16(168),
	},
	44: {
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	45: {
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	46: {
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	47: {
		Fbits: uint8(9),
		Fval:  uint16(232),
	},
	48: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	49: {
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	50: {
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	51: {
		Fbits: uint8(9),
		Fval:  uint16(152),
	},
	52: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	53: {
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	54: {
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	55: {
		Fbits: uint8(9),
		Fval:  uint16(216),
	},
	56: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	57: {
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	58: {
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	59: {
		Fbits: uint8(9),
		Fval:  uint16(184),
	},
	60: {
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	61: {
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	62: {
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	63: {
		Fbits: uint8(9),
		Fval:  uint16(248),
	},
	64: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	65: {
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	66: {
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	67: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	68: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	69: {
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	70: {
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	71: {
		Fbits: uint8(9),
		Fval:  uint16(196),
	},
	72: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	73: {
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	74: {
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	75: {
		Fbits: uint8(9),
		Fval:  uint16(164),
	},
	76: {
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	77: {
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	78: {
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	79: {
		Fbits: uint8(9),
		Fval:  uint16(228),
	},
	80: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	81: {
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	82: {
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	83: {
		Fbits: uint8(9),
		Fval:  uint16(148),
	},
	84: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	85: {
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	86: {
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	87: {
		Fbits: uint8(9),
		Fval:  uint16(212),
	},
	88: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	89: {
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	90: {
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	91: {
		Fbits: uint8(9),
		Fval:  uint16(180),
	},
	92: {
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	93: {
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	94: {
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	95: {
		Fbits: uint8(9),
		Fval:  uint16(244),
	},
	96: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	97: {
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	98: {
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	99: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	100: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	101: {
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	102: {
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	103: {
		Fbits: uint8(9),
		Fval:  uint16(204),
	},
	104: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	105: {
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	106: {
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	107: {
		Fbits: uint8(9),
		Fval:  uint16(172),
	},
	108: {
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	109: {
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	110: {
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	111: {
		Fbits: uint8(9),
		Fval:  uint16(236),
	},
	112: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	113: {
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	114: {
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	115: {
		Fbits: uint8(9),
		Fval:  uint16(156),
	},
	116: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	117: {
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	118: {
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	119: {
		Fbits: uint8(9),
		Fval:  uint16(220),
	},
	120: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	121: {
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	122: {
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	123: {
		Fbits: uint8(9),
		Fval:  uint16(188),
	},
	124: {
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	125: {
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	126: {
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	127: {
		Fbits: uint8(9),
		Fval:  uint16(252),
	},
	128: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	129: {
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	130: {
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	131: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	132: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	133: {
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	134: {
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	135: {
		Fbits: uint8(9),
		Fval:  uint16(194),
	},
	136: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	137: {
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	138: {
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	139: {
		Fbits: uint8(9),
		Fval:  uint16(162),
	},
	140: {
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	141: {
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	142: {
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	143: {
		Fbits: uint8(9),
		Fval:  uint16(226),
	},
	144: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	145: {
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	146: {
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	147: {
		Fbits: uint8(9),
		Fval:  uint16(146),
	},
	148: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	149: {
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	150: {
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	151: {
		Fbits: uint8(9),
		Fval:  uint16(210),
	},
	152: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	153: {
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	154: {
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	155: {
		Fbits: uint8(9),
		Fval:  uint16(178),
	},
	156: {
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	157: {
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	158: {
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	159: {
		Fbits: uint8(9),
		Fval:  uint16(242),
	},
	160: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	161: {
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	162: {
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	163: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	164: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	165: {
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	166: {
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	167: {
		Fbits: uint8(9),
		Fval:  uint16(202),
	},
	168: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	169: {
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	170: {
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	171: {
		Fbits: uint8(9),
		Fval:  uint16(170),
	},
	172: {
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	173: {
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	174: {
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	175: {
		Fbits: uint8(9),
		Fval:  uint16(234),
	},
	176: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	177: {
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	178: {
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	179: {
		Fbits: uint8(9),
		Fval:  uint16(154),
	},
	180: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	181: {
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	182: {
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	183: {
		Fbits: uint8(9),
		Fval:  uint16(218),
	},
	184: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	185: {
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	186: {
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	187: {
		Fbits: uint8(9),
		Fval:  uint16(186),
	},
	188: {
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	189: {
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	190: {
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	191: {
		Fbits: uint8(9),
		Fval:  uint16(250),
	},
	192: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	193: {
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	194: {
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	195: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	196: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	197: {
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	198: {
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	199: {
		Fbits: uint8(9),
		Fval:  uint16(198),
	},
	200: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	201: {
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	202: {
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	203: {
		Fbits: uint8(9),
		Fval:  uint16(166),
	},
	204: {
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	205: {
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	206: {
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	207: {
		Fbits: uint8(9),
		Fval:  uint16(230),
	},
	208: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	209: {
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	210: {
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	211: {
		Fbits: uint8(9),
		Fval:  uint16(150),
	},
	212: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	213: {
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	214: {
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	215: {
		Fbits: uint8(9),
		Fval:  uint16(214),
	},
	216: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	217: {
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	218: {
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	219: {
		Fbits: uint8(9),
		Fval:  uint16(182),
	},
	220: {
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	221: {
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	222: {
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	223: {
		Fbits: uint8(9),
		Fval:  uint16(246),
	},
	224: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	225: {
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	226: {
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	227: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	228: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	229: {
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	230: {
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	231: {
		Fbits: uint8(9),
		Fval:  uint16(206),
	},
	232: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	233: {
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	234: {
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	235: {
		Fbits: uint8(9),
		Fval:  uint16(174),
	},
	236: {
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	237: {
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	238: {
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	239: {
		Fbits: uint8(9),
		Fval:  uint16(238),
	},
	240: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	241: {
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	242: {
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	243: {
		Fbits: uint8(9),
		Fval:  uint16(158),
	},
	244: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	245: {
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	246: {
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	247: {
		Fbits: uint8(9),
		Fval:  uint16(222),
	},
	248: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	249: {
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	250: {
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	251: {
		Fbits: uint8(9),
		Fval:  uint16(190),
	},
	252: {
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	253: {
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	254: {
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	255: {
		Fbits: uint8(9),
		Fval:  uint16(254),
	},
	256: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	257: {
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	258: {
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	259: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	260: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	261: {
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	262: {
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	263: {
		Fbits: uint8(9),
		Fval:  uint16(193),
	},
	264: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	265: {
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	266: {
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	267: {
		Fbits: uint8(9),
		Fval:  uint16(161),
	},
	268: {
		Fbits: uint8(8),
	},
	269: {
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	270: {
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	271: {
		Fbits: uint8(9),
		Fval:  uint16(225),
	},
	272: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	273: {
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	274: {
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	275: {
		Fbits: uint8(9),
		Fval:  uint16(145),
	},
	276: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	277: {
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	278: {
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	279: {
		Fbits: uint8(9),
		Fval:  uint16(209),
	},
	280: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	281: {
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	282: {
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	283: {
		Fbits: uint8(9),
		Fval:  uint16(177),
	},
	284: {
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	285: {
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	286: {
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	287: {
		Fbits: uint8(9),
		Fval:  uint16(241),
	},
	288: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	289: {
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	290: {
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	291: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	292: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	293: {
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	294: {
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	295: {
		Fbits: uint8(9),
		Fval:  uint16(201),
	},
	296: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	297: {
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	298: {
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	299: {
		Fbits: uint8(9),
		Fval:  uint16(169),
	},
	300: {
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	301: {
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	302: {
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	303: {
		Fbits: uint8(9),
		Fval:  uint16(233),
	},
	304: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	305: {
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	306: {
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	307: {
		Fbits: uint8(9),
		Fval:  uint16(153),
	},
	308: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	309: {
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	310: {
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	311: {
		Fbits: uint8(9),
		Fval:  uint16(217),
	},
	312: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	313: {
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	314: {
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	315: {
		Fbits: uint8(9),
		Fval:  uint16(185),
	},
	316: {
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	317: {
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	318: {
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	319: {
		Fbits: uint8(9),
		Fval:  uint16(249),
	},
	320: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	321: {
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	322: {
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	323: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	324: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	325: {
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	326: {
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	327: {
		Fbits: uint8(9),
		Fval:  uint16(197),
	},
	328: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	329: {
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	330: {
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	331: {
		Fbits: uint8(9),
		Fval:  uint16(165),
	},
	332: {
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	333: {
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	334: {
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	335: {
		Fbits: uint8(9),
		Fval:  uint16(229),
	},
	336: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	337: {
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	338: {
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	339: {
		Fbits: uint8(9),
		Fval:  uint16(149),
	},
	340: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	341: {
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	342: {
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	343: {
		Fbits: uint8(9),
		Fval:  uint16(213),
	},
	344: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	345: {
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	346: {
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	347: {
		Fbits: uint8(9),
		Fval:  uint16(181),
	},
	348: {
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	349: {
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	350: {
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	351: {
		Fbits: uint8(9),
		Fval:  uint16(245),
	},
	352: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	353: {
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	354: {
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	355: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	356: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	357: {
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	358: {
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	359: {
		Fbits: uint8(9),
		Fval:  uint16(205),
	},
	360: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	361: {
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	362: {
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	363: {
		Fbits: uint8(9),
		Fval:  uint16(173),
	},
	364: {
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	365: {
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	366: {
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	367: {
		Fbits: uint8(9),
		Fval:  uint16(237),
	},
	368: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	369: {
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	370: {
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	371: {
		Fbits: uint8(9),
		Fval:  uint16(157),
	},
	372: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	373: {
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	374: {
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	375: {
		Fbits: uint8(9),
		Fval:  uint16(221),
	},
	376: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	377: {
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	378: {
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	379: {
		Fbits: uint8(9),
		Fval:  uint16(189),
	},
	380: {
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	381: {
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	382: {
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	383: {
		Fbits: uint8(9),
		Fval:  uint16(253),
	},
	384: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	385: {
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	386: {
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	387: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	388: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	389: {
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	390: {
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	391: {
		Fbits: uint8(9),
		Fval:  uint16(195),
	},
	392: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	393: {
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	394: {
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	395: {
		Fbits: uint8(9),
		Fval:  uint16(163),
	},
	396: {
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	397: {
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	398: {
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	399: {
		Fbits: uint8(9),
		Fval:  uint16(227),
	},
	400: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	401: {
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	402: {
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	403: {
		Fbits: uint8(9),
		Fval:  uint16(147),
	},
	404: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	405: {
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	406: {
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	407: {
		Fbits: uint8(9),
		Fval:  uint16(211),
	},
	408: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	409: {
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	410: {
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	411: {
		Fbits: uint8(9),
		Fval:  uint16(179),
	},
	412: {
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	413: {
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	414: {
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	415: {
		Fbits: uint8(9),
		Fval:  uint16(243),
	},
	416: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	417: {
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	418: {
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	419: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	420: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	421: {
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	422: {
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	423: {
		Fbits: uint8(9),
		Fval:  uint16(203),
	},
	424: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	425: {
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	426: {
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	427: {
		Fbits: uint8(9),
		Fval:  uint16(171),
	},
	428: {
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	429: {
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	430: {
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	431: {
		Fbits: uint8(9),
		Fval:  uint16(235),
	},
	432: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	433: {
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	434: {
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	435: {
		Fbits: uint8(9),
		Fval:  uint16(155),
	},
	436: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	437: {
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	438: {
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	439: {
		Fbits: uint8(9),
		Fval:  uint16(219),
	},
	440: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	441: {
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	442: {
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	443: {
		Fbits: uint8(9),
		Fval:  uint16(187),
	},
	444: {
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	445: {
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	446: {
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	447: {
		Fbits: uint8(9),
		Fval:  uint16(251),
	},
	448: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	449: {
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	450: {
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	451: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	452: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	453: {
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	454: {
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	455: {
		Fbits: uint8(9),
		Fval:  uint16(199),
	},
	456: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	457: {
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	458: {
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	459: {
		Fbits: uint8(9),
		Fval:  uint16(167),
	},
	460: {
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	461: {
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	462: {
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	463: {
		Fbits: uint8(9),
		Fval:  uint16(231),
	},
	464: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	465: {
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	466: {
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	467: {
		Fbits: uint8(9),
		Fval:  uint16(151),
	},
	468: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	469: {
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	470: {
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	471: {
		Fbits: uint8(9),
		Fval:  uint16(215),
	},
	472: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	473: {
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	474: {
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	475: {
		Fbits: uint8(9),
		Fval:  uint16(183),
	},
	476: {
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	477: {
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	478: {
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	479: {
		Fbits: uint8(9),
		Fval:  uint16(247),
	},
	480: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	481: {
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	482: {
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	483: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	484: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	485: {
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	486: {
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	487: {
		Fbits: uint8(9),
		Fval:  uint16(207),
	},
	488: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	489: {
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	490: {
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	491: {
		Fbits: uint8(9),
		Fval:  uint16(175),
	},
	492: {
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	493: {
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	494: {
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	495: {
		Fbits: uint8(9),
		Fval:  uint16(239),
	},
	496: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	497: {
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	498: {
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	499: {
		Fbits: uint8(9),
		Fval:  uint16(159),
	},
	500: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	501: {
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	502: {
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	503: {
		Fbits: uint8(9),
		Fval:  uint16(223),
	},
	504: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	505: {
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	506: {
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	507: {
		Fbits: uint8(9),
		Fval:  uint16(191),
	},
	508: {
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	509: {
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	510: {
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	511: {
		Fbits: uint8(9),
		Fval:  uint16(255),
	},
}

var _distfix = [32]Tcode{
	0: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(1),
	},
	1: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(257),
	},
	2: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(17),
	},
	3: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(4097),
	},
	4: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(5),
	},
	5: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1025),
	},
	6: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(65),
	},
	7: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(16385),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(3),
	},
	9: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(513),
	},
	10: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(33),
	},
	11: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(8193),
	},
	12: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(9),
	},
	13: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(2049),
	},
	14: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(129),
	},
	15: {
		Fop:   uint8(64),
		Fbits: uint8(5),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(2),
	},
	17: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(385),
	},
	18: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(25),
	},
	19: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(6145),
	},
	20: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(7),
	},
	21: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1537),
	},
	22: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(97),
	},
	23: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(24577),
	},
	24: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(4),
	},
	25: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(769),
	},
	26: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(49),
	},
	27: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(12289),
	},
	28: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(13),
	},
	29: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(3073),
	},
	30: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(193),
	},
	31: {
		Fop:   uint8(64),
		Fbits: uint8(5),
	},
}

type t__ccgo_fp__XinflateBack_1 = func(*libc.TLS, uintptr, uintptr) uint32

type t__ccgo_fp__XinflateBack_3 = func(*libc.TLS, uintptr, uintptr, uint32) int32

/* Macros for inflateBack(): */

/* Load returned state from inflate_fast() */

/* Set state from registers for inflate_fast() */

/* Clear the input bit accumulator */

/* Assure that some input is available.  If input is requested, but denied,
   then return a Z_BUF_ERROR from inflateBack(). */

/* Get a byte of input into the bit accumulator, or return from inflateBack()
   with an error if there is no input available. */

/* Assure that there are at least n bits in the bit accumulator.  If there is
   not enough available input to do that, then return from inflateBack() with
   an error. */

/* Return the low n bits of the bit accumulator (n < 16) */

/* Remove n bits from the bit accumulator */

/* Remove zero to seven bits as needed to go to a byte boundary */

/* Assure that some output space is available, by writing out the window
   if it's full.  If the write fails, return from inflateBack() with a
   Z_BUF_ERROR. */

// C documentation
//
//	/*
//	   strm provides the memory allocation functions and window buffer on input,
//	   and provides information on the unused input on return.  For Z_DATA_ERROR
//	   returns, strm will also provide an error message.
//
//	   in() and out() are the call-back input and output functions.  When
//	   inflateBack() needs more input, it calls in().  When inflateBack() has
//	   filled the window with output, or when it completes with data in the
//	   window, it calls out() to write out the data.  The application must not
//	   change the provided input until in() is called again or inflateBack()
//	   returns.  The application must not change the window/output buffer until
//	   inflateBack() returns.
//
//	   in() and out() are called with a descriptor parameter provided in the
//	   inflateBack() call.  This parameter can be a structure that provides the
//	   information required to do the read or write, as well as accumulated
//	   information on the input and output such as totals and check values.
//
//	   in() should return zero on failure.  out() should return non-zero on
//	   failure.  If either in() or out() fails, than inflateBack() returns a
//	   Z_BUF_ERROR.  strm->next_in can be checked for Z_NULL to see whether it
//	   was in() or out() that caused in the error.  Otherwise,  inflateBack()
//	   returns Z_STREAM_END on success, Z_DATA_ERROR for an deflate format
//	   error, or Z_MEM_ERROR if it could not allocate memory for the state.
//	   inflateBack() can also return Z_STREAM_ERROR if the input parameters
//	   are not correct, i.e. strm is Z_NULL or the state was not initialized.
//	 */
func XinflateBack(tls *libc.TLS, strm Tz_streamp, __ccgo_fp_in Tin_func, in_desc uintptr, __ccgo_fp_out Tout_func, out_desc uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bits, copy1, have, left, len1, v1, v18, v20, v24, v29, v30, v43, v44 uint32
	var from, put, state, v11, v15, v16, v17, v19, v21, v23, v25, v26, v27, v28, v31, v33, v35, v36, v37, v39, v41, v42, v46, v47 uintptr
	var here, last Tcode
	var hold uint64
	var ret int32
	var _ /* next at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bits, copy1, from, have, here, hold, last, left, len1, put, ret, state, v1, v11, v15, v16, v17, v18, v19, v20, v21, v23, v24, v25, v26, v27, v28, v29, v30, v31, v33, v35, v36, v37, v39, v41, v42, v43, v44, v46, v47 /* return code */
	/* Check that the strm exists and that the state was initialized */
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fstate == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* Reset the state */
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
	(*Tinflate_state)(unsafe.Pointer(state)).Flast = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	*(*uintptr)(unsafe.Pointer(bp)) = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	if *(*uintptr)(unsafe.Pointer(bp)) != uintptr(m_Z_NULL) {
		v1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	} else {
		v1 = uint32(0)
	}
	have = v1
	hold = uint64(0)
	bits = uint32(0)
	put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
	left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	/* Inflate until end of block marked as last */
	for {
		switch (*Tinflate_state)(unsafe.Pointer(state)).Fmode {
		case int32(_TYPE):
			goto _3
		case int32(_STORED):
			goto _4
		case int32(_TABLE):
			goto _5
		case int32(_LEN):
			goto _6
		case int32(_DONE):
			goto _7
		case int32(_BAD):
			goto _8
		default:
			goto _9
		}
		goto _10
	_3:
		;
		/* determine and dispatch block type */
		if (*Tinflate_state)(unsafe.Pointer(state)).Flast != 0 {
			hold >>= uint64(bits & uint32(7))
			bits -= bits & uint32(7)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DONE)
			goto _10
		}
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(3)) {
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v11 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint64(*(*uint8)(unsafe.Pointer(v11))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flast = libc.Int32FromUint32(uint32(hold) & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(1) - libc.Uint32FromInt32(1)))
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(1))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(1))
		switch uint32(hold) & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2) - libc.Uint32FromInt32(1)) {
		case uint32(0): /* stored block */
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_STORED)
		case uint32(1): /* fixed block */
			_fixedtables(tls, state)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN) /* decode codes */
		case uint32(2): /* dynamic block */
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TABLE)
		case uint32(3):
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 6
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
		}
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(2))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
		goto _10
	_4:
		;
		/* get and verify stored block length */
	_14:
		;
		hold >>= uint64(bits & uint32(7))
		bits -= bits & uint32(7)
		goto _13
	_13:
		;
		if 0 != 0 {
			goto _14
		}
		goto _12
	_12:
		; /* go to byte boundary */
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v15 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint64(*(*uint8)(unsafe.Pointer(v15))) << bits
			bits += uint32(8)
		}
		if hold&uint64(0xffff) != hold>>int32(16)^uint64(0xffff) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 25
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(hold) & uint32(0xffff)
		hold = uint64(0)
		bits = uint32(0)
		/* copy stored block from input to output */
		for (*Tinflate_state)(unsafe.Pointer(state)).Flength != uint32(0) {
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			if left == uint32(0) {
				put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
				left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = left
				if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_out})))(tls, out_desc, put, left) != 0 {
					ret = -int32(5)
					goto inf_leave
				}
			}
			if copy1 > have {
				copy1 = have
			}
			if copy1 > left {
				copy1 = left
			}
			libc.Xmemcpy(tls, put, *(*uintptr)(unsafe.Pointer(bp)), uint64(copy1))
			have -= copy1
			*(*uintptr)(unsafe.Pointer(bp)) += uintptr(copy1)
			left -= copy1
			put += uintptr(copy1)
			*(*uint32)(unsafe.Pointer(state + 92)) -= copy1
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		goto _10
	_5:
		;
		/* get dynamic table entries descriptor */
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(14)) {
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v16 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint64(*(*uint8)(unsafe.Pointer(v16))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnlen = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(257)
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fndist = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(1)
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fncode = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(4)
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(4))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(4))
		if (*Tinflate_state)(unsafe.Pointer(state)).Fnlen > uint32(286) || (*Tinflate_state)(unsafe.Pointer(state)).Fndist > uint32(30) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 54
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* get code length code lengths (not a typo) */
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fncode {
			for bits < libc.Uint32FromInt32(libc.Int32FromInt32(3)) {
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v17 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint64(*(*uint8)(unsafe.Pointer(v17))) << bits
				bits += uint32(8)
			}
			v19 = state + 140
			v18 = *(*uint32)(unsafe.Pointer(v19))
			*(*uint32)(unsafe.Pointer(v19))++
			*(*uint16)(unsafe.Pointer(state + 152 + uintptr(_order[v18])*2)) = uint16(uint32(hold) & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3) - libc.Uint32FromInt32(1)))
			hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(3))
			bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
		}
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < uint32(19) {
			v21 = state + 140
			v20 = *(*uint32)(unsafe.Pointer(v21))
			*(*uint32)(unsafe.Pointer(v21))++
			*(*uint16)(unsafe.Pointer(state + 152 + uintptr(_order[v20])*2)) = uint16(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1368
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(7)
		ret = Xinflate_table(tls, int32(_CODES), state+152, uint32(19), state+144, state+120, state+792)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 90
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* get length and distance code code lengths */
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(hold)&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
				if uint32(here.Fbits) <= bits {
					break
				}
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v23 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint64(*(*uint8)(unsafe.Pointer(v23))) << bits
				bits += uint32(8)
				goto _22
			_22:
			}
			if libc.Int32FromUint16(here.Fval) < int32(16) {
				hold >>= uint64(here.Fbits)
				bits -= uint32(here.Fbits)
				v25 = state + 140
				v24 = *(*uint32)(unsafe.Pointer(v25))
				*(*uint32)(unsafe.Pointer(v25))++
				*(*uint16)(unsafe.Pointer(state + 152 + uintptr(v24)*2)) = here.Fval
			} else {
				if libc.Int32FromUint16(here.Fval) == int32(16) {
					for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(2)) {
						if have == uint32(0) {
							have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
							if have == uint32(0) {
								*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
								ret = -int32(5)
								goto inf_leave
							}
						}
						have--
						v26 = *(*uintptr)(unsafe.Pointer(bp))
						*(*uintptr)(unsafe.Pointer(bp))++
						hold += uint64(*(*uint8)(unsafe.Pointer(v26))) << bits
						bits += uint32(8)
					}
					hold >>= uint64(here.Fbits)
					bits -= uint32(here.Fbits)
					if (*Tinflate_state)(unsafe.Pointer(state)).Fhave == uint32(0) {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 115
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
					len1 = uint32(*(*uint16)(unsafe.Pointer(state + 152 + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fhave-uint32(1))*2)))
					copy1 = uint32(3) + uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2)-libc.Uint32FromInt32(1))
					hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(2))
					bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
				} else {
					if libc.Int32FromUint16(here.Fval) == int32(17) {
						for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(3)) {
							if have == uint32(0) {
								have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
								if have == uint32(0) {
									*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
									ret = -int32(5)
									goto inf_leave
								}
							}
							have--
							v27 = *(*uintptr)(unsafe.Pointer(bp))
							*(*uintptr)(unsafe.Pointer(bp))++
							hold += uint64(*(*uint8)(unsafe.Pointer(v27))) << bits
							bits += uint32(8)
						}
						hold >>= uint64(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(3) + uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(3))
						bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
					} else {
						for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(7)) {
							if have == uint32(0) {
								have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
								if have == uint32(0) {
									*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
									ret = -int32(5)
									goto inf_leave
								}
							}
							have--
							v28 = *(*uintptr)(unsafe.Pointer(bp))
							*(*uintptr)(unsafe.Pointer(bp))++
							hold += uint64(*(*uint8)(unsafe.Pointer(v28))) << bits
							bits += uint32(8)
						}
						hold >>= uint64(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(11) + uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(7)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(7))
						bits -= libc.Uint32FromInt32(libc.Int32FromInt32(7))
					}
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhave+copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
					(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 115
					(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
					break
				}
				for {
					v29 = copy1
					copy1--
					if !(v29 != 0) {
						break
					}
					v31 = state + 140
					v30 = *(*uint32)(unsafe.Pointer(v31))
					*(*uint32)(unsafe.Pointer(v31))++
					*(*uint16)(unsafe.Pointer(state + 152 + uintptr(v30)*2)) = uint16(len1)
				}
			}
		}
		/* handle error breaks in while */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_BAD) {
			goto _10
		}
		/* check for end-of-block code (better have one) */
		if libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(state + 152 + 256*2))) == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 141
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* build code tables -- note: do not change the lenbits or distbits
		   values here (9 and 6) without reading the comments in inftrees.h
		   concerning the ENOUGH constants, which depend on those values */
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1368
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
		ret = Xinflate_table(tls, int32(_LENS), state+152, (*Tinflate_state)(unsafe.Pointer(state)).Fnlen, state+144, state+120, state+792)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 178
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(6)
		ret = Xinflate_table(tls, int32(_DISTS), state+152+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fnlen)*2, (*Tinflate_state)(unsafe.Pointer(state)).Fndist, state+144, state+124, state+792)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 206
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
		/* fallthrough */
	_6:
		;
		/* use inflate_fast() if we have enough input and output */
		if have >= uint32(6) && left >= uint32(258) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = *(*uintptr)(unsafe.Pointer(bp))
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
			(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
			(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - left
			}
			Xinflate_fast(tls, strm, (*Tinflate_state)(unsafe.Pointer(state)).Fwsize)
			put = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
			left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
			*(*uintptr)(unsafe.Pointer(bp)) = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
			have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
			bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
			goto _10
		}
		/* get a literal, length, or end-of-block code */
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(hold)&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v33 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint64(*(*uint8)(unsafe.Pointer(v33))) << bits
			bits += uint32(8)
			goto _32
		_32:
		}
		if here.Fop != 0 && libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(last.Fval)+uint32(hold)&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
				if libc.Uint32FromInt32(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v35 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint64(*(*uint8)(unsafe.Pointer(v35))) << bits
				bits += uint32(8)
				goto _34
			_34:
			}
			hold >>= uint64(last.Fbits)
			bits -= uint32(last.Fbits)
		}
		hold >>= uint64(here.Fbits)
		bits -= uint32(here.Fbits)
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(here.Fval)
		/* process literal */
		if libc.Int32FromUint8(here.Fop) == 0 {
			if left == uint32(0) {
				put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
				left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = left
				if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_out})))(tls, out_desc, put, left) != 0 {
					ret = -int32(5)
					goto inf_leave
				}
			}
			v36 = put
			put++
			*(*uint8)(unsafe.Pointer(v36)) = uint8((*Tinflate_state)(unsafe.Pointer(state)).Flength)
			left--
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
			goto _10
		}
		/* process end of block */
		if libc.Int32FromUint8(here.Fop)&int32(32) != 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
			goto _10
		}
		/* invalid code */
		if libc.Int32FromUint8(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 228
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* length code -- get extra bits, if any */
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != uint32(0) {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v37 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint64(*(*uint8)(unsafe.Pointer(v37))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 92)) += uint32(hold) & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= uint64((*Tinflate_state)(unsafe.Pointer(state)).Fextra)
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
		}
		/* get distance code */
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(hold)&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v39 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint64(*(*uint8)(unsafe.Pointer(v39))) << bits
			bits += uint32(8)
			goto _38
		_38:
		}
		if libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(last.Fval)+uint32(hold)&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
				if libc.Uint32FromInt32(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v41 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint64(*(*uint8)(unsafe.Pointer(v41))) << bits
				bits += uint32(8)
				goto _40
			_40:
			}
			hold >>= uint64(last.Fbits)
			bits -= uint32(last.Fbits)
		}
		hold >>= uint64(here.Fbits)
		bits -= uint32(here.Fbits)
		if libc.Int32FromUint8(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 256
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Foffset = uint32(here.Fval)
		/* get distance extra bits, if any */
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != uint32(0) {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v42 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint64(*(*uint8)(unsafe.Pointer(v42))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 96)) += uint32(hold) & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= uint64((*Tinflate_state)(unsafe.Pointer(state)).Fextra)
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
			v43 = left
		} else {
			v43 = uint32(0)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Foffset > (*Tinflate_state)(unsafe.Pointer(state)).Fwsize-v43 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 278
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* copy match from window to output */
		for cond := true; cond; cond = (*Tinflate_state)(unsafe.Pointer(state)).Flength != uint32(0) {
			if left == uint32(0) {
				put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
				left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = left
				if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_out})))(tls, out_desc, put, left) != 0 {
					ret = -int32(5)
					goto inf_leave
				}
			}
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - (*Tinflate_state)(unsafe.Pointer(state)).Foffset
			if copy1 < left {
				from = put + uintptr(copy1)
				copy1 = left - copy1
			} else {
				from = put - uintptr((*Tinflate_state)(unsafe.Pointer(state)).Foffset)
				copy1 = left
			}
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Flength {
				copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			}
			*(*uint32)(unsafe.Pointer(state + 92)) -= copy1
			left -= copy1
			for {
				v46 = put
				put++
				v47 = from
				from++
				*(*uint8)(unsafe.Pointer(v46)) = *(*uint8)(unsafe.Pointer(v47))
				goto _45
			_45:
				;
				copy1--
				v44 = copy1
				if !(v44 != 0) {
					break
				}
			}
		}
		goto _10
	_7:
		;
		/* inflate stream terminated properly */
		ret = int32(m_Z_STREAM_END)
		goto inf_leave
	_8:
		;
		ret = -int32(3)
		goto inf_leave
	_9:
		;
		/* can't happen, but makes compilers happy */
		ret = -int32(2)
		goto inf_leave
	_10:
		;
		goto _2
	_2:
	}
	/* Write leftover output and return unused input */
	goto inf_leave
inf_leave:
	;
	if left < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
		if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_out})))(tls, out_desc, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, (*Tinflate_state)(unsafe.Pointer(state)).Fwsize-left) != 0 && ret == int32(m_Z_STREAM_END) {
			ret = -int32(5)
		}
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = *(*uintptr)(unsafe.Pointer(bp))
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
	return ret
}

var _order = [19]uint16{
	0:  uint16(16),
	1:  uint16(17),
	2:  uint16(18),
	4:  uint16(8),
	5:  uint16(7),
	6:  uint16(9),
	7:  uint16(6),
	8:  uint16(10),
	9:  uint16(5),
	10: uint16(11),
	11: uint16(4),
	12: uint16(12),
	13: uint16(3),
	14: uint16(13),
	15: uint16(2),
	16: uint16(14),
	17: uint16(1),
	18: uint16(15),
}

func XinflateBackEnd(tls *libc.TLS, strm Tz_streamp) (r int32) {
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fstate == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Decode literal, length, and distance codes and write out the resulting
//	   literal and match bytes until either not enough input or output is
//	   available, an end-of-block is encountered, or a data error is encountered.
//	   When large enough input and output buffers are supplied to inflate(), for
//	   example, a 16K input buffer and a 64K output buffer, more than 95% of the
//	   inflate execution time is spent in this routine.
//
//	   Entry assumptions:
//
//	        state->mode == LEN
//	        strm->avail_in >= 6
//	        strm->avail_out >= 258
//	        start >= strm->avail_out
//	        state->bits < 8
//
//	   On return, state->mode is one of:
//
//	        LEN -- ran out of enough output space or enough available input
//	        TYPE -- reached end of block code, inflate() to interpret next block
//	        BAD -- error in block data
//
//	   Notes:
//
//	    - The maximum input bits used by a length/distance pair is 15 bits for the
//	      length code, 5 bits for the length extra, 15 bits for the distance code,
//	      and 13 bits for the distance extra.  This totals 48 bits, or six bytes.
//	      Therefore if strm->avail_in >= 6, then there is enough input to avoid
//	      checking for available input while decoding.
//
//	    - The maximum bytes that a single length/distance pair can output is 258
//	      bytes, which is the maximum length that can be coded.  inflate_fast()
//	      requires strm->avail_out >= 258 for each loop to avoid checking for
//	      output space.
//	 */
func Xinflate_fast(tls *libc.TLS, strm Tz_streamp, start uint32) {
	var beg, dcode, end, from, here, in, last, lcode, out, state, window, v1, v11, v12, v15, v16, v19, v2, v20, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v5, v6, v7, v8 uintptr
	var bits, dist, dmask, len1, lmask, op, whave, wnext, wsize, v13, v17, v21, v9 uint32
	var hold uint64
	var v45, v46 int64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = beg, bits, dcode, dist, dmask, end, from, here, hold, in, last, lcode, len1, lmask, op, out, state, whave, window, wnext, wsize, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v5, v6, v7, v8, v9 /* where to copy match from */
	/* copy state to local variables */
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	in = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	last = in + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in-libc.Uint32FromInt32(5))
	out = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
	beg = out - uintptr(start-(*Tz_stream)(unsafe.Pointer(strm)).Favail_out)
	end = out + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_out-libc.Uint32FromInt32(257))
	wsize = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	whave = (*Tinflate_state)(unsafe.Pointer(state)).Fwhave
	wnext = (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
	window = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
	hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
	bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
	lcode = (*Tinflate_state)(unsafe.Pointer(state)).Flencode
	dcode = (*Tinflate_state)(unsafe.Pointer(state)).Fdistcode
	lmask = uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits - uint32(1)
	dmask = uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits - uint32(1)
	/* decode literals and length/distances until end-of-block or not enough
	   input data or output space */
	for cond := true; cond; cond = in < last && out < end {
		if bits < uint32(15) {
			v1 = in
			in++
			hold += uint64(*(*uint8)(unsafe.Pointer(v1))) << bits
			bits += uint32(8)
			v2 = in
			in++
			hold += uint64(*(*uint8)(unsafe.Pointer(v2))) << bits
			bits += uint32(8)
		}
		here = lcode + uintptr(hold&uint64(lmask))*4
		goto dolen
	dolen:
		;
		op = uint32((*Tcode)(unsafe.Pointer(here)).Fbits)
		hold >>= uint64(op)
		bits -= op
		op = uint32((*Tcode)(unsafe.Pointer(here)).Fop)
		if op == uint32(0) { /* literal */
			v3 = out
			out++
			*(*uint8)(unsafe.Pointer(v3)) = uint8((*Tcode)(unsafe.Pointer(here)).Fval)
		} else {
			if op&uint32(16) != 0 { /* length base */
				len1 = uint32((*Tcode)(unsafe.Pointer(here)).Fval)
				op &= uint32(15) /* number of extra bits */
				if op != 0 {
					if bits < op {
						v4 = in
						in++
						hold += uint64(*(*uint8)(unsafe.Pointer(v4))) << bits
						bits += uint32(8)
					}
					len1 += uint32(hold) & (uint32(1)<<op - uint32(1))
					hold >>= uint64(op)
					bits -= op
				}
				if bits < uint32(15) {
					v5 = in
					in++
					hold += uint64(*(*uint8)(unsafe.Pointer(v5))) << bits
					bits += uint32(8)
					v6 = in
					in++
					hold += uint64(*(*uint8)(unsafe.Pointer(v6))) << bits
					bits += uint32(8)
				}
				here = dcode + uintptr(hold&uint64(dmask))*4
				goto dodist
			dodist:
				;
				op = uint32((*Tcode)(unsafe.Pointer(here)).Fbits)
				hold >>= uint64(op)
				bits -= op
				op = uint32((*Tcode)(unsafe.Pointer(here)).Fop)
				if op&uint32(16) != 0 { /* distance base */
					dist = uint32((*Tcode)(unsafe.Pointer(here)).Fval)
					op &= uint32(15) /* number of extra bits */
					if bits < op {
						v7 = in
						in++
						hold += uint64(*(*uint8)(unsafe.Pointer(v7))) << bits
						bits += uint32(8)
						if bits < op {
							v8 = in
							in++
							hold += uint64(*(*uint8)(unsafe.Pointer(v8))) << bits
							bits += uint32(8)
						}
					}
					dist += uint32(hold) & (uint32(1)<<op - uint32(1))
					hold >>= uint64(op)
					bits -= op
					op = libc.Uint32FromInt64(int64(out) - int64(beg)) /* max distance in output */
					if dist > op {                                     /* see if copy from window */
						op = dist - op /* distance back in window */
						if op > whave {
							if (*Tinflate_state)(unsafe.Pointer(state)).Fsane != 0 {
								(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 278
								(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
								break
							}
						}
						from = window
						if wnext == uint32(0) { /* very common case */
							from += uintptr(wsize - op)
							if op < len1 { /* some from window */
								len1 -= op
								for {
									v11 = out
									out++
									v12 = from
									from++
									*(*uint8)(unsafe.Pointer(v11)) = *(*uint8)(unsafe.Pointer(v12))
									goto _10
								_10:
									;
									op--
									v9 = op
									if !(v9 != 0) {
										break
									}
								}
								from = out - uintptr(dist) /* rest from output */
							}
						} else {
							if wnext < op { /* wrap around window */
								from += uintptr(wsize + wnext - op)
								op -= wnext
								if op < len1 { /* some from end of window */
									len1 -= op
									for {
										v15 = out
										out++
										v16 = from
										from++
										*(*uint8)(unsafe.Pointer(v15)) = *(*uint8)(unsafe.Pointer(v16))
										goto _14
									_14:
										;
										op--
										v13 = op
										if !(v13 != 0) {
											break
										}
									}
									from = window
									if wnext < len1 { /* some from start of window */
										op = wnext
										len1 -= op
										for {
											v19 = out
											out++
											v20 = from
											from++
											*(*uint8)(unsafe.Pointer(v19)) = *(*uint8)(unsafe.Pointer(v20))
											goto _18
										_18:
											;
											op--
											v17 = op
											if !(v17 != 0) {
												break
											}
										}
										from = out - uintptr(dist) /* rest from output */
									}
								}
							} else { /* contiguous in window */
								from += uintptr(wnext - op)
								if op < len1 { /* some from window */
									len1 -= op
									for {
										v23 = out
										out++
										v24 = from
										from++
										*(*uint8)(unsafe.Pointer(v23)) = *(*uint8)(unsafe.Pointer(v24))
										goto _22
									_22:
										;
										op--
										v21 = op
										if !(v21 != 0) {
											break
										}
									}
									from = out - uintptr(dist) /* rest from output */
								}
							}
						}
						for len1 > uint32(2) {
							v25 = out
							out++
							v26 = from
							from++
							*(*uint8)(unsafe.Pointer(v25)) = *(*uint8)(unsafe.Pointer(v26))
							v27 = out
							out++
							v28 = from
							from++
							*(*uint8)(unsafe.Pointer(v27)) = *(*uint8)(unsafe.Pointer(v28))
							v29 = out
							out++
							v30 = from
							from++
							*(*uint8)(unsafe.Pointer(v29)) = *(*uint8)(unsafe.Pointer(v30))
							len1 -= uint32(3)
						}
						if len1 != 0 {
							v31 = out
							out++
							v32 = from
							from++
							*(*uint8)(unsafe.Pointer(v31)) = *(*uint8)(unsafe.Pointer(v32))
							if len1 > uint32(1) {
								v33 = out
								out++
								v34 = from
								from++
								*(*uint8)(unsafe.Pointer(v33)) = *(*uint8)(unsafe.Pointer(v34))
							}
						}
					} else {
						from = out - uintptr(dist)                        /* copy direct from output */
						for cond := true; cond; cond = len1 > uint32(2) { /* minimum length is three */
							v35 = out
							out++
							v36 = from
							from++
							*(*uint8)(unsafe.Pointer(v35)) = *(*uint8)(unsafe.Pointer(v36))
							v37 = out
							out++
							v38 = from
							from++
							*(*uint8)(unsafe.Pointer(v37)) = *(*uint8)(unsafe.Pointer(v38))
							v39 = out
							out++
							v40 = from
							from++
							*(*uint8)(unsafe.Pointer(v39)) = *(*uint8)(unsafe.Pointer(v40))
							len1 -= uint32(3)
						}
						if len1 != 0 {
							v41 = out
							out++
							v42 = from
							from++
							*(*uint8)(unsafe.Pointer(v41)) = *(*uint8)(unsafe.Pointer(v42))
							if len1 > uint32(1) {
								v43 = out
								out++
								v44 = from
								from++
								*(*uint8)(unsafe.Pointer(v43)) = *(*uint8)(unsafe.Pointer(v44))
							}
						}
					}
				} else {
					if op&uint32(64) == uint32(0) { /* 2nd level distance code */
						here = dcode + uintptr((*Tcode)(unsafe.Pointer(here)).Fval)*4 + uintptr(hold&uint64(libc.Uint32FromUint32(1)<<op-libc.Uint32FromInt32(1)))*4
						goto dodist
					} else {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 256
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
				}
			} else {
				if op&uint32(64) == uint32(0) { /* 2nd level length code */
					here = lcode + uintptr((*Tcode)(unsafe.Pointer(here)).Fval)*4 + uintptr(hold&uint64(libc.Uint32FromUint32(1)<<op-libc.Uint32FromInt32(1)))*4
					goto dolen
				} else {
					if op&uint32(32) != 0 { /* end-of-block */
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
						break
					} else {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 228
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
				}
			}
		}
	}
	/* return unused bytes (on entry, bits < 8, so in won't go too far back) */
	len1 = bits >> int32(3)
	in -= uintptr(len1)
	bits -= len1 << int32(3)
	hold &= uint64(uint32(1)<<bits - uint32(1))
	/* update state and return */
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = in
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = out
	if in < last {
		v45 = int64(5) + (int64(last) - int64(in))
	} else {
		v45 = int64(5) - (int64(in) - int64(last))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = libc.Uint32FromInt64(v45)
	if out < end {
		v46 = int64(257) + (int64(end) - int64(out))
	} else {
		v46 = int64(257) - (int64(out) - int64(end))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = libc.Uint32FromInt64(v46)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
	return
}

func _inflateStateCheck(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) || (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		return int32(1)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if state == uintptr(m_Z_NULL) || (*Tinflate_state)(unsafe.Pointer(state)).Fstrm != strm || (*Tinflate_state)(unsafe.Pointer(state)).Fmode < int32(_HEAD) || (*Tinflate_state)(unsafe.Pointer(state)).Fmode > int32(_SYNC) {
		return int32(1)
	}
	return 0
}

func XinflateResetKeep(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state, v3, v4 uintptr
	var v1 TuLong
	var v2 uint64
	_, _, _, _, _ = state, v1, v2, v3, v4
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	v2 = libc.Uint64FromInt32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Ftotal = v2
	v1 = v2
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 { /* to support ill-conceived Java test suite */
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = libc.Uint64FromInt32((*Tinflate_state)(unsafe.Pointer(state)).Fwrap & int32(1))
	}
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HEAD)
	(*Tinflate_state)(unsafe.Pointer(state)).Flast = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fhavedict = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fflags = -int32(1)
	(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(32768)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhead = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = uint64(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = uint32(0)
	v4 = state + 1368
	(*Tinflate_state)(unsafe.Pointer(state)).Fnext = v4
	v3 = v4
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = v3
	(*Tinflate_state)(unsafe.Pointer(state)).Flencode = v3
	(*Tinflate_state)(unsafe.Pointer(state)).Fsane = int32(1)
	(*Tinflate_state)(unsafe.Pointer(state)).Fback = -int32(1)
	return m_Z_OK
}

func XinflateReset(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tinflate_state)(unsafe.Pointer(state)).Fwsize = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
	return XinflateResetKeep(tls, strm)
}

func XinflateReset2(tls *libc.TLS, strm Tz_streamp, windowBits int32) (r int32) {
	var state uintptr
	var wrap int32
	_, _ = state, wrap
	/* get the state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* extract wrap request from windowBits parameter */
	if windowBits < 0 {
		if windowBits < -int32(15) {
			return -int32(2)
		}
		wrap = 0
		windowBits = -windowBits
	} else {
		wrap = windowBits>>int32(4) + int32(5)
		if windowBits < int32(48) {
			windowBits &= int32(15)
		}
	}
	/* set number of window bits, free window if different */
	if windowBits != 0 && (windowBits < int32(8) || windowBits > int32(15)) {
		return -int32(2)
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) && (*Tinflate_state)(unsafe.Pointer(state)).Fwbits != libc.Uint32FromInt32(windowBits) {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = uintptr(m_Z_NULL)
	}
	/* update state and reset the rest of it */
	(*Tinflate_state)(unsafe.Pointer(state)).Fwrap = wrap
	(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = libc.Uint32FromInt32(windowBits)
	return XinflateReset(tls, strm)
}

func XinflateInit2_(tls *libc.TLS, strm Tz_streamp, windowBits int32, version uintptr, stream_size int32) (r int32) {
	var ret int32
	var state uintptr
	_, _ = ret, state
	if version == uintptr(m_Z_NULL) || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(version))) != libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(__ccgo_ts))) || stream_size != libc.Int32FromUint64(libc.Uint64FromInt64(112)) {
		return -int32(6)
	}
	if strm == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* in case we return an error */
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = __ccgo_fp(Xzcalloc)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = libc.UintptrFromInt32(0)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = __ccgo_fp(Xzcfree)
	}
	state = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), uint32(libc.Uint64FromInt64(7160)))
	if state == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = state
	(*Tinflate_state)(unsafe.Pointer(state)).Fstrm = strm
	(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HEAD) /* to pass state test in inflateReset2() */
	ret = XinflateReset2(tls, strm, windowBits)
	if ret != m_Z_OK {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, state)
		(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	}
	return ret
}

func XinflateInit_(tls *libc.TLS, strm Tz_streamp, version uintptr, stream_size int32) (r int32) {
	return XinflateInit2_(tls, strm, int32(m_MAX_WBITS), version, stream_size)
}

func XinflatePrime(tls *libc.TLS, strm Tz_streamp, bits int32, value int32) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	if bits == 0 {
		return m_Z_OK
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if bits < 0 {
		(*Tinflate_state)(unsafe.Pointer(state)).Fhold = uint64(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fbits = uint32(0)
		return m_Z_OK
	}
	if bits > int32(16) || (*Tinflate_state)(unsafe.Pointer(state)).Fbits+libc.Uint32FromInt32(bits) > uint32(32) {
		return -int32(2)
	}
	value = int32(int64(value) & (libc.Int64FromInt64(1)<<bits - libc.Int64FromInt32(1)))
	*(*uint64)(unsafe.Pointer(state + 80)) += uint64(libc.Uint32FromInt32(value) << (*Tinflate_state)(unsafe.Pointer(state)).Fbits)
	*(*uint32)(unsafe.Pointer(state + 88)) += libc.Uint32FromInt32(bits)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Return state with length and distance decoding tables and index sizes set to
//	   fixed code decoding.  Normally this returns fixed tables from inffixed.h.
//	   If BUILDFIXED is defined, then instead this routine builds the tables the
//	   first time it's called, and returns those tables the first time and
//	   thereafter.  This reduces the size of the code by about 2K bytes, in
//	   exchange for a little execution time.  However, BUILDFIXED should not be
//	   used for threaded applications, since the rewriting of the tables and virgin
//	   may not be thread-safe.
//	 */
func _fixedtables1(tls *libc.TLS, state uintptr) {
	(*Tinflate_state)(unsafe.Pointer(state)).Flencode = uintptr(unsafe.Pointer(&_lenfix1))
	(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = uintptr(unsafe.Pointer(&_distfix1))
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(5)
}

var _lenfix1 = [512]Tcode{
	0: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	1: {
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	2: {
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	3: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	4: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	5: {
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	6: {
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	7: {
		Fbits: uint8(9),
		Fval:  uint16(192),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	9: {
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	10: {
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	11: {
		Fbits: uint8(9),
		Fval:  uint16(160),
	},
	12: {
		Fbits: uint8(8),
	},
	13: {
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	14: {
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	15: {
		Fbits: uint8(9),
		Fval:  uint16(224),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	17: {
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	18: {
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	19: {
		Fbits: uint8(9),
		Fval:  uint16(144),
	},
	20: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	21: {
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	22: {
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	23: {
		Fbits: uint8(9),
		Fval:  uint16(208),
	},
	24: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	25: {
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	26: {
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	27: {
		Fbits: uint8(9),
		Fval:  uint16(176),
	},
	28: {
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	29: {
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	30: {
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	31: {
		Fbits: uint8(9),
		Fval:  uint16(240),
	},
	32: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	33: {
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	34: {
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	35: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	36: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	37: {
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	38: {
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	39: {
		Fbits: uint8(9),
		Fval:  uint16(200),
	},
	40: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	41: {
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	42: {
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	43: {
		Fbits: uint8(9),
		Fval:  uint16(168),
	},
	44: {
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	45: {
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	46: {
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	47: {
		Fbits: uint8(9),
		Fval:  uint16(232),
	},
	48: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	49: {
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	50: {
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	51: {
		Fbits: uint8(9),
		Fval:  uint16(152),
	},
	52: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	53: {
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	54: {
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	55: {
		Fbits: uint8(9),
		Fval:  uint16(216),
	},
	56: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	57: {
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	58: {
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	59: {
		Fbits: uint8(9),
		Fval:  uint16(184),
	},
	60: {
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	61: {
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	62: {
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	63: {
		Fbits: uint8(9),
		Fval:  uint16(248),
	},
	64: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	65: {
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	66: {
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	67: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	68: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	69: {
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	70: {
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	71: {
		Fbits: uint8(9),
		Fval:  uint16(196),
	},
	72: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	73: {
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	74: {
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	75: {
		Fbits: uint8(9),
		Fval:  uint16(164),
	},
	76: {
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	77: {
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	78: {
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	79: {
		Fbits: uint8(9),
		Fval:  uint16(228),
	},
	80: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	81: {
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	82: {
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	83: {
		Fbits: uint8(9),
		Fval:  uint16(148),
	},
	84: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	85: {
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	86: {
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	87: {
		Fbits: uint8(9),
		Fval:  uint16(212),
	},
	88: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	89: {
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	90: {
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	91: {
		Fbits: uint8(9),
		Fval:  uint16(180),
	},
	92: {
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	93: {
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	94: {
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	95: {
		Fbits: uint8(9),
		Fval:  uint16(244),
	},
	96: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	97: {
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	98: {
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	99: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	100: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	101: {
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	102: {
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	103: {
		Fbits: uint8(9),
		Fval:  uint16(204),
	},
	104: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	105: {
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	106: {
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	107: {
		Fbits: uint8(9),
		Fval:  uint16(172),
	},
	108: {
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	109: {
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	110: {
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	111: {
		Fbits: uint8(9),
		Fval:  uint16(236),
	},
	112: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	113: {
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	114: {
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	115: {
		Fbits: uint8(9),
		Fval:  uint16(156),
	},
	116: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	117: {
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	118: {
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	119: {
		Fbits: uint8(9),
		Fval:  uint16(220),
	},
	120: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	121: {
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	122: {
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	123: {
		Fbits: uint8(9),
		Fval:  uint16(188),
	},
	124: {
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	125: {
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	126: {
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	127: {
		Fbits: uint8(9),
		Fval:  uint16(252),
	},
	128: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	129: {
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	130: {
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	131: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	132: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	133: {
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	134: {
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	135: {
		Fbits: uint8(9),
		Fval:  uint16(194),
	},
	136: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	137: {
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	138: {
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	139: {
		Fbits: uint8(9),
		Fval:  uint16(162),
	},
	140: {
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	141: {
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	142: {
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	143: {
		Fbits: uint8(9),
		Fval:  uint16(226),
	},
	144: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	145: {
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	146: {
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	147: {
		Fbits: uint8(9),
		Fval:  uint16(146),
	},
	148: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	149: {
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	150: {
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	151: {
		Fbits: uint8(9),
		Fval:  uint16(210),
	},
	152: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	153: {
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	154: {
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	155: {
		Fbits: uint8(9),
		Fval:  uint16(178),
	},
	156: {
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	157: {
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	158: {
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	159: {
		Fbits: uint8(9),
		Fval:  uint16(242),
	},
	160: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	161: {
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	162: {
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	163: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	164: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	165: {
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	166: {
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	167: {
		Fbits: uint8(9),
		Fval:  uint16(202),
	},
	168: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	169: {
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	170: {
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	171: {
		Fbits: uint8(9),
		Fval:  uint16(170),
	},
	172: {
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	173: {
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	174: {
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	175: {
		Fbits: uint8(9),
		Fval:  uint16(234),
	},
	176: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	177: {
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	178: {
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	179: {
		Fbits: uint8(9),
		Fval:  uint16(154),
	},
	180: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	181: {
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	182: {
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	183: {
		Fbits: uint8(9),
		Fval:  uint16(218),
	},
	184: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	185: {
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	186: {
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	187: {
		Fbits: uint8(9),
		Fval:  uint16(186),
	},
	188: {
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	189: {
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	190: {
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	191: {
		Fbits: uint8(9),
		Fval:  uint16(250),
	},
	192: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	193: {
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	194: {
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	195: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	196: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	197: {
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	198: {
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	199: {
		Fbits: uint8(9),
		Fval:  uint16(198),
	},
	200: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	201: {
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	202: {
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	203: {
		Fbits: uint8(9),
		Fval:  uint16(166),
	},
	204: {
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	205: {
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	206: {
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	207: {
		Fbits: uint8(9),
		Fval:  uint16(230),
	},
	208: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	209: {
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	210: {
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	211: {
		Fbits: uint8(9),
		Fval:  uint16(150),
	},
	212: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	213: {
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	214: {
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	215: {
		Fbits: uint8(9),
		Fval:  uint16(214),
	},
	216: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	217: {
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	218: {
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	219: {
		Fbits: uint8(9),
		Fval:  uint16(182),
	},
	220: {
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	221: {
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	222: {
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	223: {
		Fbits: uint8(9),
		Fval:  uint16(246),
	},
	224: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	225: {
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	226: {
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	227: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	228: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	229: {
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	230: {
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	231: {
		Fbits: uint8(9),
		Fval:  uint16(206),
	},
	232: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	233: {
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	234: {
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	235: {
		Fbits: uint8(9),
		Fval:  uint16(174),
	},
	236: {
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	237: {
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	238: {
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	239: {
		Fbits: uint8(9),
		Fval:  uint16(238),
	},
	240: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	241: {
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	242: {
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	243: {
		Fbits: uint8(9),
		Fval:  uint16(158),
	},
	244: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	245: {
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	246: {
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	247: {
		Fbits: uint8(9),
		Fval:  uint16(222),
	},
	248: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	249: {
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	250: {
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	251: {
		Fbits: uint8(9),
		Fval:  uint16(190),
	},
	252: {
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	253: {
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	254: {
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	255: {
		Fbits: uint8(9),
		Fval:  uint16(254),
	},
	256: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	257: {
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	258: {
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	259: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	260: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	261: {
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	262: {
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	263: {
		Fbits: uint8(9),
		Fval:  uint16(193),
	},
	264: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	265: {
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	266: {
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	267: {
		Fbits: uint8(9),
		Fval:  uint16(161),
	},
	268: {
		Fbits: uint8(8),
	},
	269: {
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	270: {
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	271: {
		Fbits: uint8(9),
		Fval:  uint16(225),
	},
	272: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	273: {
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	274: {
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	275: {
		Fbits: uint8(9),
		Fval:  uint16(145),
	},
	276: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	277: {
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	278: {
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	279: {
		Fbits: uint8(9),
		Fval:  uint16(209),
	},
	280: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	281: {
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	282: {
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	283: {
		Fbits: uint8(9),
		Fval:  uint16(177),
	},
	284: {
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	285: {
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	286: {
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	287: {
		Fbits: uint8(9),
		Fval:  uint16(241),
	},
	288: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	289: {
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	290: {
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	291: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	292: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	293: {
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	294: {
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	295: {
		Fbits: uint8(9),
		Fval:  uint16(201),
	},
	296: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	297: {
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	298: {
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	299: {
		Fbits: uint8(9),
		Fval:  uint16(169),
	},
	300: {
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	301: {
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	302: {
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	303: {
		Fbits: uint8(9),
		Fval:  uint16(233),
	},
	304: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	305: {
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	306: {
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	307: {
		Fbits: uint8(9),
		Fval:  uint16(153),
	},
	308: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	309: {
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	310: {
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	311: {
		Fbits: uint8(9),
		Fval:  uint16(217),
	},
	312: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	313: {
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	314: {
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	315: {
		Fbits: uint8(9),
		Fval:  uint16(185),
	},
	316: {
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	317: {
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	318: {
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	319: {
		Fbits: uint8(9),
		Fval:  uint16(249),
	},
	320: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	321: {
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	322: {
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	323: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	324: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	325: {
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	326: {
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	327: {
		Fbits: uint8(9),
		Fval:  uint16(197),
	},
	328: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	329: {
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	330: {
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	331: {
		Fbits: uint8(9),
		Fval:  uint16(165),
	},
	332: {
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	333: {
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	334: {
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	335: {
		Fbits: uint8(9),
		Fval:  uint16(229),
	},
	336: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	337: {
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	338: {
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	339: {
		Fbits: uint8(9),
		Fval:  uint16(149),
	},
	340: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	341: {
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	342: {
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	343: {
		Fbits: uint8(9),
		Fval:  uint16(213),
	},
	344: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	345: {
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	346: {
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	347: {
		Fbits: uint8(9),
		Fval:  uint16(181),
	},
	348: {
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	349: {
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	350: {
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	351: {
		Fbits: uint8(9),
		Fval:  uint16(245),
	},
	352: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	353: {
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	354: {
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	355: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	356: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	357: {
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	358: {
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	359: {
		Fbits: uint8(9),
		Fval:  uint16(205),
	},
	360: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	361: {
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	362: {
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	363: {
		Fbits: uint8(9),
		Fval:  uint16(173),
	},
	364: {
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	365: {
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	366: {
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	367: {
		Fbits: uint8(9),
		Fval:  uint16(237),
	},
	368: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	369: {
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	370: {
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	371: {
		Fbits: uint8(9),
		Fval:  uint16(157),
	},
	372: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	373: {
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	374: {
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	375: {
		Fbits: uint8(9),
		Fval:  uint16(221),
	},
	376: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	377: {
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	378: {
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	379: {
		Fbits: uint8(9),
		Fval:  uint16(189),
	},
	380: {
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	381: {
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	382: {
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	383: {
		Fbits: uint8(9),
		Fval:  uint16(253),
	},
	384: {
		Fop:   uint8(96),
		Fbits: uint8(7),
	},
	385: {
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	386: {
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	387: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	388: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	389: {
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	390: {
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	391: {
		Fbits: uint8(9),
		Fval:  uint16(195),
	},
	392: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	393: {
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	394: {
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	395: {
		Fbits: uint8(9),
		Fval:  uint16(163),
	},
	396: {
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	397: {
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	398: {
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	399: {
		Fbits: uint8(9),
		Fval:  uint16(227),
	},
	400: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	401: {
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	402: {
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	403: {
		Fbits: uint8(9),
		Fval:  uint16(147),
	},
	404: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	405: {
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	406: {
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	407: {
		Fbits: uint8(9),
		Fval:  uint16(211),
	},
	408: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	409: {
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	410: {
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	411: {
		Fbits: uint8(9),
		Fval:  uint16(179),
	},
	412: {
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	413: {
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	414: {
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	415: {
		Fbits: uint8(9),
		Fval:  uint16(243),
	},
	416: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	417: {
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	418: {
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	419: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	420: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	421: {
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	422: {
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	423: {
		Fbits: uint8(9),
		Fval:  uint16(203),
	},
	424: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	425: {
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	426: {
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	427: {
		Fbits: uint8(9),
		Fval:  uint16(171),
	},
	428: {
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	429: {
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	430: {
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	431: {
		Fbits: uint8(9),
		Fval:  uint16(235),
	},
	432: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	433: {
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	434: {
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	435: {
		Fbits: uint8(9),
		Fval:  uint16(155),
	},
	436: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	437: {
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	438: {
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	439: {
		Fbits: uint8(9),
		Fval:  uint16(219),
	},
	440: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	441: {
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	442: {
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	443: {
		Fbits: uint8(9),
		Fval:  uint16(187),
	},
	444: {
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	445: {
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	446: {
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	447: {
		Fbits: uint8(9),
		Fval:  uint16(251),
	},
	448: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	449: {
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	450: {
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	451: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	452: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	453: {
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	454: {
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	455: {
		Fbits: uint8(9),
		Fval:  uint16(199),
	},
	456: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	457: {
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	458: {
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	459: {
		Fbits: uint8(9),
		Fval:  uint16(167),
	},
	460: {
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	461: {
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	462: {
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	463: {
		Fbits: uint8(9),
		Fval:  uint16(231),
	},
	464: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	465: {
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	466: {
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	467: {
		Fbits: uint8(9),
		Fval:  uint16(151),
	},
	468: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	469: {
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	470: {
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	471: {
		Fbits: uint8(9),
		Fval:  uint16(215),
	},
	472: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	473: {
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	474: {
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	475: {
		Fbits: uint8(9),
		Fval:  uint16(183),
	},
	476: {
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	477: {
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	478: {
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	479: {
		Fbits: uint8(9),
		Fval:  uint16(247),
	},
	480: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	481: {
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	482: {
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	483: {
		Fop:   uint8(64),
		Fbits: uint8(8),
	},
	484: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	485: {
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	486: {
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	487: {
		Fbits: uint8(9),
		Fval:  uint16(207),
	},
	488: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	489: {
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	490: {
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	491: {
		Fbits: uint8(9),
		Fval:  uint16(175),
	},
	492: {
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	493: {
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	494: {
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	495: {
		Fbits: uint8(9),
		Fval:  uint16(239),
	},
	496: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	497: {
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	498: {
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	499: {
		Fbits: uint8(9),
		Fval:  uint16(159),
	},
	500: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	501: {
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	502: {
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	503: {
		Fbits: uint8(9),
		Fval:  uint16(223),
	},
	504: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	505: {
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	506: {
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	507: {
		Fbits: uint8(9),
		Fval:  uint16(191),
	},
	508: {
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	509: {
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	510: {
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	511: {
		Fbits: uint8(9),
		Fval:  uint16(255),
	},
}

var _distfix1 = [32]Tcode{
	0: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(1),
	},
	1: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(257),
	},
	2: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(17),
	},
	3: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(4097),
	},
	4: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(5),
	},
	5: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1025),
	},
	6: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(65),
	},
	7: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(16385),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(3),
	},
	9: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(513),
	},
	10: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(33),
	},
	11: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(8193),
	},
	12: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(9),
	},
	13: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(2049),
	},
	14: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(129),
	},
	15: {
		Fop:   uint8(64),
		Fbits: uint8(5),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(2),
	},
	17: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(385),
	},
	18: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(25),
	},
	19: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(6145),
	},
	20: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(7),
	},
	21: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1537),
	},
	22: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(97),
	},
	23: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(24577),
	},
	24: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(4),
	},
	25: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(769),
	},
	26: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(49),
	},
	27: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(12289),
	},
	28: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(13),
	},
	29: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(3073),
	},
	30: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(193),
	},
	31: {
		Fop:   uint8(64),
		Fbits: uint8(5),
	},
}

// C documentation
//
//	/*
//	   Update the window with the last wsize (normally 32K) bytes written before
//	   returning.  If window does not exist yet, create it.  This is only called
//	   when a window is already in use, or when output has been written during this
//	   inflate call, but the end of the deflate stream has not been reached yet.
//	   It is also called to create a window for dictionary data when a dictionary
//	   is loaded.
//
//	   Providing output buffers larger than 32K to inflate() should provide a speed
//	   advantage, since only the last 32K of output is copied to the sliding window
//	   upon return from inflate(), and since all distances after the first 32K of
//	   output will fall in the output data, making match copies simpler and faster.
//	   The advantage may be dependent on the size of the processor's data caches.
//	 */
func _updatewindow(tls *libc.TLS, strm Tz_streamp, end uintptr, copy1 uint32) (r int32) {
	var dist uint32
	var state uintptr
	_, _ = dist, state
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* if it hasn't been done already, allocate space for the window */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow == uintptr(m_Z_NULL) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fwbits, uint32(libc.Uint64FromInt64(1)))
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow == uintptr(m_Z_NULL) {
			return int32(1)
		}
	}
	/* if window not in use yet, initialize */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwsize == uint32(0) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fwsize = uint32(1) << (*Tinflate_state)(unsafe.Pointer(state)).Fwbits
		(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	}
	/* copy state->wsize or less output bytes into the circular window */
	if copy1 >= (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
		libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, end-uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwsize), uint64((*Tinflate_state)(unsafe.Pointer(state)).Fwsize))
		(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	} else {
		dist = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
		if dist > copy1 {
			dist = copy1
		}
		libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), end-uintptr(copy1), uint64(dist))
		copy1 -= dist
		if copy1 != 0 {
			libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, end-uintptr(copy1), uint64(copy1))
			(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = copy1
			(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
		} else {
			*(*uint32)(unsafe.Pointer(state + 68)) += dist
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwnext == (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				*(*uint32)(unsafe.Pointer(state + 64)) += dist
			}
		}
	}
	return 0
}

/* Macros for inflate(): */

/* check function to use adler32() for zlib or crc32() for gzip */

/* check macros for header crc */

/* Load registers with state in inflate() for speed */

/* Restore state from registers in inflate() */

/* Clear the input bit accumulator */

/* Get a byte of input into the bit accumulator, or return from inflate()
   if there is no input available. */

/* Assure that there are at least n bits in the bit accumulator.  If there is
   not enough available input to do that, then return from inflate(). */

/* Return the low n bits of the bit accumulator (n < 16) */

/* Remove n bits from the bit accumulator */

/* Remove zero to seven bits as needed to go to a byte boundary */

/*
   inflate() uses a state machine to process as much input data and generate as
   much output data as possible before returning.  The state machine is
   structured roughly as follows:

    for (;;) switch (state) {
    ...
    case STATEn:
        if (not enough input data or output space to make progress)
            return;
        ... make progress ...
        state = STATEm;
        break;
    ...
    }

   so when inflate() is called again, the same case is attempted again, and
   if the appropriate resources are provided, the machine proceeds to the
   next state.  The NEEDBITS() macro is usually the way the state evaluates
   whether it can proceed or should return.  NEEDBITS() does the return if
   the requested bits are not available.  The typical use of the BITS macros
   is:

        NEEDBITS(n);
        ... do something with BITS(n) ...
        DROPBITS(n);

   where NEEDBITS(n) either returns from inflate() if there isn't enough
   input left to load n bits into the accumulator, or it continues.  BITS(n)
   gives the low n bits in the accumulator.  When done, DROPBITS(n) drops
   the low n bits off the accumulator.  INITBITS() clears the accumulator
   and sets the number of available bits to zero.  BYTEBITS() discards just
   enough bits to put the accumulator on a byte boundary.  After BYTEBITS()
   and a NEEDBITS(8), then BITS(8) would return the next byte in the stream.

   NEEDBITS(n) uses PULLBYTE() to get an available byte of input, or to return
   if there is no input available.  The decoding of variable length codes uses
   PULLBYTE() directly in order to pull just enough bytes to decode the next
   code, and no more.

   Some states loop until they get enough input, making sure that enough
   state information is maintained to continue the loop where it left off
   if NEEDBITS() returns in the loop.  For example, want, need, and keep
   would all have to actually be part of the saved state in case NEEDBITS()
   returns:

    case STATEw:
        while (want < need) {
            NEEDBITS(n);
            keep[want++] = BITS(n);
            DROPBITS(n);
        }
        state = STATEx;
    case STATEx:

   As shown above, if the next state is also the next case, then the break
   is omitted.

   A state may also return if there is not enough output space available to
   complete that state.  Those states are copying stored data, writing a
   literal byte, and copying a matching string.

   When returning, a "goto inf_leave" is used to update the total counters,
   update the check value, and determine whether any progress has been made
   during that inflate() call in order to return the proper return code.
   Progress is defined as a change in either strm->avail_in or strm->avail_out.
   When there is a window, goto inf_leave will update the window with the last
   output written.  If a goto inf_leave occurs in the middle of decompression
   and there is no window currently, goto inf_leave will create one and copy
   output to the window for the next call of inflate().

   In this implementation, the flush parameter of inflate() only affects the
   return code (per zlib.h).  inflate() always writes as much as possible to
   strm->next_out, given the space available and the provided input--the effect
   documented in zlib.h of Z_SYNC_FLUSH.  Furthermore, inflate() always defers
   the allocation of and copying into a sliding window until necessary, which
   provides the effect documented in zlib.h for Z_FINISH when the entire input
   stream available.  So the only thing the flush parameter actually does is:
   when flush is set to Z_FINISH, inflate() cannot return Z_OK.  Instead it
   will return Z_BUF_ERROR if it has not reached the end of the stream.
*/

func Xinflate(tls *libc.TLS, strm Tz_streamp, flush int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bits, copy1, have, in, left, len1, out, v46, v48, v49, v50, v52, v53, v64, v66, v70, v75, v76, v90 uint32
	var from, next, put, state, v100, v36, v42, v43, v44, v45, v51, v54, v55, v57, v60, v61, v62, v63, v65, v67, v69, v71, v72, v73, v74, v77, v79, v81, v82, v85, v87, v88, v92, v93, v94, v95, p83, p89 uintptr
	var here, last Tcode
	var hold, v101, v102, v37, v56, v58, v59, v96, v97, v98 uint64
	var ret, v103, v104, v105, v38 int32
	var v47, v99 bool
	var _ /* hbuf at bp+0 */ [4]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bits, copy1, from, have, here, hold, in, last, left, len1, next, out, put, ret, state, v100, v101, v102, v103, v104, v105, v36, v37, v38, v42, v43, v44, v45, v46, v47, v48, v49, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v62, v63, v64, v65, v66, v67, v69, v70, v71, v72, v73, v74, v75, v76, v77, v79, v81, v82, v85, v87, v88, v90, v92, v93, v94, v95, v96, v97, v98, v99, p83, p89 /* buffer for gzip header crc calculation */
	if _inflateStateCheck(tls, strm) != 0 || (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in == uintptr(m_Z_NULL) && (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_TYPE) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPEDO)
	} /* skip check */
	put = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
	left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	next = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
	bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
	in = have
	out = left
	ret = m_Z_OK
	for {
		switch (*Tinflate_state)(unsafe.Pointer(state)).Fmode {
		case int32(_HEAD):
			goto _2
		case int32(_FLAGS):
			goto _3
		case int32(_TIME):
			goto _4
		case int32(_OS):
			goto _5
		case int32(_EXLEN):
			goto _6
		case int32(_EXTRA):
			goto _7
		case int32(_NAME):
			goto _8
		case int32(_COMMENT):
			goto _9
		case int32(_HCRC):
			goto _10
		case int32(_DICTID):
			goto _11
		case int32(_DICT):
			goto _12
		case int32(_TYPE):
			goto _13
		case int32(_TYPEDO):
			goto _14
		case int32(_STORED):
			goto _15
		case int32(_COPY_):
			goto _16
		case int32(_COPY):
			goto _17
		case int32(_TABLE):
			goto _18
		case int32(_LENLENS):
			goto _19
		case int32(_CODELENS):
			goto _20
		case int32(_LEN_):
			goto _21
		case int32(_LEN):
			goto _22
		case int32(_LENEXT):
			goto _23
		case int32(_DIST):
			goto _24
		case int32(_DISTEXT):
			goto _25
		case int32(_MATCH):
			goto _26
		case int32(_LIT):
			goto _27
		case int32(_CHECK):
			goto _28
		case int32(_LENGTH):
			goto _29
		case int32(_DONE):
			goto _30
		case int32(_BAD):
			goto _31
		case int32(_MEM):
			goto _32
		default:
			goto _33
		case int32(_SYNC):
			goto _34
		}
		goto _35
	_2:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap == 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPEDO)
			goto _35
		}
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(16)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v36 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v36))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(2) != 0 && hold == uint64(0x8b1f) { /* gzip header */
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwbits == uint32(0) {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = uint32(15)
			}
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
			hold = uint64(0)
			bits = uint32(0)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_FLAGS)
			goto _35
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fdone = -int32(1)
		}
		if !((*Tinflate_state)(unsafe.Pointer(state)).Fwrap&libc.Int32FromInt32(1) != 0) || (uint64(uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(8)-libc.Uint32FromInt32(1))<<libc.Int32FromInt32(8))+hold>>int32(8))%uint64(31) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 308
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) != uint32(m_Z_DEFLATED) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 331
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(4))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(4))
		len1 = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(8)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwbits == uint32(0) {
			(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = len1
		}
		if len1 > uint32(15) || len1 > (*Tinflate_state)(unsafe.Pointer(state)).Fwbits {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 358
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(1) << len1
		(*Tinflate_state)(unsafe.Pointer(state)).Fflags = 0 /* indicate zlib header */
		v37 = Xadler32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v37
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v37
		if hold&uint64(0x200) != 0 {
			v38 = int32(_DICTID)
		} else {
			v38 = int32(_TYPE)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = v38
		hold = uint64(0)
		bits = uint32(0)
		goto _35
	_3:
		;
	_41:
		;
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(16)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v42 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v42))) << bits
			bits += uint32(8)
		}
		goto _40
	_40:
		;
		if 0 != 0 {
			goto _41
		}
		goto _39
	_39:
		;
		(*Tinflate_state)(unsafe.Pointer(state)).Fflags = libc.Int32FromUint64(hold)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0xff) != int32(m_Z_DEFLATED) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 331
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0xe000) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 378
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Ftext = libc.Int32FromUint64(hold >> libc.Int32FromInt32(8) & libc.Uint64FromInt32(1))
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
		}
		hold = uint64(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TIME)
		/* fallthrough */
	_4:
		;
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v43 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v43))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Ftime = hold
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(2)] = uint8(hold >> libc.Int32FromInt32(16))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(3)] = uint8(hold >> libc.Int32FromInt32(24))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(4))
		}
		hold = uint64(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_OS)
		/* fallthrough */
	_5:
		;
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(16)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v44 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v44))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fxflags = libc.Int32FromUint64(hold & libc.Uint64FromInt32(0xff))
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fos = libc.Int32FromUint64(hold >> libc.Int32FromInt32(8))
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
		}
		hold = uint64(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_EXLEN)
		/* fallthrough */
	_6:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0400) != 0 {
			for bits < libc.Uint32FromInt32(libc.Int32FromInt32(16)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v45 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v45))) << bits
				bits += uint32(8)
			}
			(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(hold)
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_len = uint32(hold)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
				(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
			}
			hold = uint64(0)
			bits = uint32(0)
		} else {
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra = uintptr(m_Z_NULL)
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_EXTRA)
		/* fallthrough */
	_7:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0400) != 0 {
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			if copy1 > have {
				copy1 = have
			}
			if copy1 != 0 {
				if v47 = (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) && (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra != uintptr(m_Z_NULL); v47 {
					v46 = (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_len - (*Tinflate_state)(unsafe.Pointer(state)).Flength
					len1 = v46
				}
				if v47 && v46 < (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_max {
					if len1+copy1 > (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_max {
						v48 = (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_max - len1
					} else {
						v48 = copy1
					}
					libc.Xmemcpy(tls, (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra+uintptr(len1), next, uint64(v48))
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
					(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
				}
				have -= copy1
				next += uintptr(copy1)
				*(*uint32)(unsafe.Pointer(state + 92)) -= copy1
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Flength != 0 {
				goto inf_leave
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_NAME)
		/* fallthrough */
	_8:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0800) != 0 {
			if have == uint32(0) {
				goto inf_leave
			}
			copy1 = uint32(0)
			for cond := true; cond; cond = len1 != 0 && copy1 < have {
				v49 = copy1
				copy1++
				len1 = uint32(*(*uint8)(unsafe.Pointer(next + uintptr(v49))))
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) && (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname != uintptr(m_Z_NULL) && (*Tinflate_state)(unsafe.Pointer(state)).Flength < (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname_max {
					v51 = state + 92
					v50 = *(*uint32)(unsafe.Pointer(v51))
					*(*uint32)(unsafe.Pointer(v51))++
					*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname + uintptr(v50))) = uint8(len1)
				}
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
			}
			have -= copy1
			next += uintptr(copy1)
			if len1 != 0 {
				goto inf_leave
			}
		} else {
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname = uintptr(m_Z_NULL)
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_COMMENT)
		/* fallthrough */
	_9:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x1000) != 0 {
			if have == uint32(0) {
				goto inf_leave
			}
			copy1 = uint32(0)
			for cond := true; cond; cond = len1 != 0 && copy1 < have {
				v52 = copy1
				copy1++
				len1 = uint32(*(*uint8)(unsafe.Pointer(next + uintptr(v52))))
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) && (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomment != uintptr(m_Z_NULL) && (*Tinflate_state)(unsafe.Pointer(state)).Flength < (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomm_max {
					v54 = state + 92
					v53 = *(*uint32)(unsafe.Pointer(v54))
					*(*uint32)(unsafe.Pointer(v54))++
					*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomment + uintptr(v53))) = uint8(len1)
				}
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
			}
			have -= copy1
			next += uintptr(copy1)
			if len1 != 0 {
				goto inf_leave
			}
		} else {
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomment = uintptr(m_Z_NULL)
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HCRC)
		/* fallthrough */
	_10:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 {
			for bits < libc.Uint32FromInt32(libc.Int32FromInt32(16)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v55 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v55))) << bits
				bits += uint32(8)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && hold != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck&uint64(0xffff) {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 403
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint64(0)
			bits = uint32(0)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fhcrc = (*Tinflate_state)(unsafe.Pointer(state)).Fflags >> libc.Int32FromInt32(9) & libc.Int32FromInt32(1)
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fdone = int32(1)
		}
		v56 = Xcrc32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v56
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v56
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		goto _35
	_11:
		;
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v57 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v57))) << bits
			bits += uint32(8)
		}
		v58 = hold>>libc.Int32FromInt32(24)&libc.Uint64FromInt32(0xff) + hold>>libc.Int32FromInt32(8)&libc.Uint64FromInt32(0xff00) + hold&libc.Uint64FromInt32(0xff00)<<libc.Int32FromInt32(8) + hold&libc.Uint64FromInt32(0xff)<<libc.Int32FromInt32(24)
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v58
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v58
		hold = uint64(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DICT)
		/* fallthrough */
	_12:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhavedict == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
			(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
			(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
			return int32(m_Z_NEED_DICT)
		}
		v59 = Xadler32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v59
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v59
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		/* fallthrough */
	_13:
		;
		if flush == int32(m_Z_BLOCK) || flush == int32(m_Z_TREES) {
			goto inf_leave
		}
		/* fallthrough */
	_14:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Flast != 0 {
			hold >>= uint64(bits & uint32(7))
			bits -= bits & uint32(7)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_CHECK)
			goto _35
		}
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(3)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v60 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v60))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flast = libc.Int32FromUint32(uint32(hold) & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(1) - libc.Uint32FromInt32(1)))
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(1))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(1))
		switch uint32(hold) & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2) - libc.Uint32FromInt32(1)) {
		case uint32(0): /* stored block */
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_STORED)
		case uint32(1): /* fixed block */
			_fixedtables1(tls, state)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN_) /* decode codes */
			if flush == int32(m_Z_TREES) {
				hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(2))
				bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
				goto inf_leave
			}
		case uint32(2): /* dynamic block */
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TABLE)
		case uint32(3):
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 6
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
		}
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(2))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
		goto _35
	_15:
		;
		hold >>= uint64(bits & uint32(7))
		bits -= bits & uint32(7) /* go to byte boundary */
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v61 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v61))) << bits
			bits += uint32(8)
		}
		if hold&uint64(0xffff) != hold>>int32(16)^uint64(0xffff) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 25
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(hold) & uint32(0xffff)
		hold = uint64(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_COPY_)
		if flush == int32(m_Z_TREES) {
			goto inf_leave
		}
		/* fallthrough */
	_16:
		;
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_COPY)
		/* fallthrough */
	_17:
		;
		copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		if copy1 != 0 {
			if copy1 > have {
				copy1 = have
			}
			if copy1 > left {
				copy1 = left
			}
			if copy1 == uint32(0) {
				goto inf_leave
			}
			libc.Xmemcpy(tls, put, next, uint64(copy1))
			have -= copy1
			next += uintptr(copy1)
			left -= copy1
			put += uintptr(copy1)
			*(*uint32)(unsafe.Pointer(state + 92)) -= copy1
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		goto _35
	_18:
		;
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(14)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v62 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v62))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnlen = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(257)
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fndist = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(1)
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fncode = uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(4)
		hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(4))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(4))
		if (*Tinflate_state)(unsafe.Pointer(state)).Fnlen > uint32(286) || (*Tinflate_state)(unsafe.Pointer(state)).Fndist > uint32(30) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 54
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LENLENS)
		/* fallthrough */
	_19:
		;
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fncode {
			for bits < libc.Uint32FromInt32(libc.Int32FromInt32(3)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v63 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v63))) << bits
				bits += uint32(8)
			}
			v65 = state + 140
			v64 = *(*uint32)(unsafe.Pointer(v65))
			*(*uint32)(unsafe.Pointer(v65))++
			*(*uint16)(unsafe.Pointer(state + 152 + uintptr(_order1[v64])*2)) = uint16(uint32(hold) & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3) - libc.Uint32FromInt32(1)))
			hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(3))
			bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
		}
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < uint32(19) {
			v67 = state + 140
			v66 = *(*uint32)(unsafe.Pointer(v67))
			*(*uint32)(unsafe.Pointer(v67))++
			*(*uint16)(unsafe.Pointer(state + 152 + uintptr(_order1[v66])*2)) = uint16(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1368
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(7)
		ret = Xinflate_table(tls, int32(_CODES), state+152, uint32(19), state+144, state+120, state+792)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 90
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_CODELENS)
		/* fallthrough */
	_20:
		;
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(hold)&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
				if uint32(here.Fbits) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v69 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v69))) << bits
				bits += uint32(8)
				goto _68
			_68:
			}
			if libc.Int32FromUint16(here.Fval) < int32(16) {
				hold >>= uint64(here.Fbits)
				bits -= uint32(here.Fbits)
				v71 = state + 140
				v70 = *(*uint32)(unsafe.Pointer(v71))
				*(*uint32)(unsafe.Pointer(v71))++
				*(*uint16)(unsafe.Pointer(state + 152 + uintptr(v70)*2)) = here.Fval
			} else {
				if libc.Int32FromUint16(here.Fval) == int32(16) {
					for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(2)) {
						if have == uint32(0) {
							goto inf_leave
						}
						have--
						v72 = next
						next++
						hold += uint64(*(*uint8)(unsafe.Pointer(v72))) << bits
						bits += uint32(8)
					}
					hold >>= uint64(here.Fbits)
					bits -= uint32(here.Fbits)
					if (*Tinflate_state)(unsafe.Pointer(state)).Fhave == uint32(0) {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 115
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
					len1 = uint32(*(*uint16)(unsafe.Pointer(state + 152 + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fhave-uint32(1))*2)))
					copy1 = uint32(3) + uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2)-libc.Uint32FromInt32(1))
					hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(2))
					bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
				} else {
					if libc.Int32FromUint16(here.Fval) == int32(17) {
						for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(3)) {
							if have == uint32(0) {
								goto inf_leave
							}
							have--
							v73 = next
							next++
							hold += uint64(*(*uint8)(unsafe.Pointer(v73))) << bits
							bits += uint32(8)
						}
						hold >>= uint64(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(3) + uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(3))
						bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
					} else {
						for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(7)) {
							if have == uint32(0) {
								goto inf_leave
							}
							have--
							v74 = next
							next++
							hold += uint64(*(*uint8)(unsafe.Pointer(v74))) << bits
							bits += uint32(8)
						}
						hold >>= uint64(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(11) + uint32(hold)&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(7)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint64FromInt32(libc.Int32FromInt32(7))
						bits -= libc.Uint32FromInt32(libc.Int32FromInt32(7))
					}
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhave+copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
					(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 115
					(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
					break
				}
				for {
					v75 = copy1
					copy1--
					if !(v75 != 0) {
						break
					}
					v77 = state + 140
					v76 = *(*uint32)(unsafe.Pointer(v77))
					*(*uint32)(unsafe.Pointer(v77))++
					*(*uint16)(unsafe.Pointer(state + 152 + uintptr(v76)*2)) = uint16(len1)
				}
			}
		}
		/* handle error breaks in while */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_BAD) {
			goto _35
		}
		/* check for end-of-block code (better have one) */
		if libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(state + 152 + 256*2))) == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 141
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		/* build code tables -- note: do not change the lenbits or distbits
		   values here (9 and 6) without reading the comments in inftrees.h
		   concerning the ENOUGH constants, which depend on those values */
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1368
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
		ret = Xinflate_table(tls, int32(_LENS), state+152, (*Tinflate_state)(unsafe.Pointer(state)).Fnlen, state+144, state+120, state+792)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 178
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(6)
		ret = Xinflate_table(tls, int32(_DISTS), state+152+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fnlen)*2, (*Tinflate_state)(unsafe.Pointer(state)).Fndist, state+144, state+124, state+792)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 206
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN_)
		if flush == int32(m_Z_TREES) {
			goto inf_leave
		}
		/* fallthrough */
	_21:
		;
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
		/* fallthrough */
	_22:
		;
		if have >= uint32(6) && left >= uint32(258) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
			(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
			(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
			Xinflate_fast(tls, strm, out)
			put = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
			left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
			next = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
			have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
			bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
			if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_TYPE) {
				(*Tinflate_state)(unsafe.Pointer(state)).Fback = -int32(1)
			}
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fback = 0
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(hold)&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v79 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v79))) << bits
			bits += uint32(8)
			goto _78
		_78:
		}
		if here.Fop != 0 && libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(last.Fval)+uint32(hold)&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
				if libc.Uint32FromInt32(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v81 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v81))) << bits
				bits += uint32(8)
				goto _80
			_80:
			}
			hold >>= uint64(last.Fbits)
			bits -= uint32(last.Fbits)
			*(*int32)(unsafe.Pointer(state + 7148)) += libc.Int32FromUint8(last.Fbits)
		}
		hold >>= uint64(here.Fbits)
		bits -= uint32(here.Fbits)
		*(*int32)(unsafe.Pointer(state + 7148)) += libc.Int32FromUint8(here.Fbits)
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(here.Fval)
		if libc.Int32FromUint8(here.Fop) == 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LIT)
			goto _35
		}
		if libc.Int32FromUint8(here.Fop)&int32(32) != 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fback = -int32(1)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
			goto _35
		}
		if libc.Int32FromUint8(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 228
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LENEXT)
		/* fallthrough */
	_23:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != 0 {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v82 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v82))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 92)) += uint32(hold) & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= uint64((*Tinflate_state)(unsafe.Pointer(state)).Fextra)
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			p83 = state + 7148
			*(*int32)(unsafe.Pointer(p83)) = int32(uint32(*(*int32)(unsafe.Pointer(p83))) + (*Tinflate_state)(unsafe.Pointer(state)).Fextra)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fwas = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DIST)
		/* fallthrough */
	_24:
		;
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(hold)&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v85 = next
			next++
			hold += uint64(*(*uint8)(unsafe.Pointer(v85))) << bits
			bits += uint32(8)
			goto _84
		_84:
		}
		if libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(last.Fval)+uint32(hold)&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
				if libc.Uint32FromInt32(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v87 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v87))) << bits
				bits += uint32(8)
				goto _86
			_86:
			}
			hold >>= uint64(last.Fbits)
			bits -= uint32(last.Fbits)
			*(*int32)(unsafe.Pointer(state + 7148)) += libc.Int32FromUint8(last.Fbits)
		}
		hold >>= uint64(here.Fbits)
		bits -= uint32(here.Fbits)
		*(*int32)(unsafe.Pointer(state + 7148)) += libc.Int32FromUint8(here.Fbits)
		if libc.Int32FromUint8(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 256
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Foffset = uint32(here.Fval)
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DISTEXT)
		/* fallthrough */
	_25:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != 0 {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v88 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v88))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 96)) += uint32(hold) & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= uint64((*Tinflate_state)(unsafe.Pointer(state)).Fextra)
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			p89 = state + 7148
			*(*int32)(unsafe.Pointer(p89)) = int32(uint32(*(*int32)(unsafe.Pointer(p89))) + (*Tinflate_state)(unsafe.Pointer(state)).Fextra)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_MATCH)
		/* fallthrough */
	_26:
		;
		if left == uint32(0) {
			goto inf_leave
		}
		copy1 = out - left
		if (*Tinflate_state)(unsafe.Pointer(state)).Foffset > copy1 { /* copy from window */
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Foffset - copy1
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fwhave {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fsane != 0 {
					(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 278
					(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
					goto _35
				}
			}
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fwnext {
				copy1 -= (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
				from = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwsize-copy1)
			} else {
				from = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext-copy1)
			}
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Flength {
				copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			}
		} else { /* copy from output */
			from = put - uintptr((*Tinflate_state)(unsafe.Pointer(state)).Foffset)
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		}
		if copy1 > left {
			copy1 = left
		}
		left -= copy1
		*(*uint32)(unsafe.Pointer(state + 92)) -= copy1
		for {
			v92 = put
			put++
			v93 = from
			from++
			*(*uint8)(unsafe.Pointer(v92)) = *(*uint8)(unsafe.Pointer(v93))
			goto _91
		_91:
			;
			copy1--
			v90 = copy1
			if !(v90 != 0) {
				break
			}
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Flength == uint32(0) {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
		}
		goto _35
	_27:
		;
		if left == uint32(0) {
			goto inf_leave
		}
		v94 = put
		put++
		*(*uint8)(unsafe.Pointer(v94)) = uint8((*Tinflate_state)(unsafe.Pointer(state)).Flength)
		left--
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
		goto _35
	_28:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 {
			for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v95 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v95))) << bits
				bits += uint32(8)
			}
			out -= left
			*(*TuLong)(unsafe.Pointer(strm + 40)) += uint64(out)
			*(*uint64)(unsafe.Pointer(state + 40)) += uint64(out)
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && out != 0 {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
					v97 = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, put-uintptr(out), out)
				} else {
					v97 = Xadler32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, put-uintptr(out), out)
				}
				v96 = v97
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v96
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v96
			}
			out = left
			if v99 = (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0; v99 {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
					v98 = hold
				} else {
					v98 = hold>>int32(24)&uint64(0xff) + hold>>int32(8)&uint64(0xff00) + hold&uint64(0xff00)<<int32(8) + hold&uint64(0xff)<<int32(24)
				}
			}
			if v99 && v98 != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 423
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint64(0)
			bits = uint32(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LENGTH)
		/* fallthrough */
	_29:
		;
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
			for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v100 = next
				next++
				hold += uint64(*(*uint8)(unsafe.Pointer(v100))) << bits
				bits += uint32(8)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && hold != (*Tinflate_state)(unsafe.Pointer(state)).Ftotal&uint64(0xffffffff) {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 444
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint64(0)
			bits = uint32(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DONE)
		/* fallthrough */
	_30:
		;
		ret = int32(m_Z_STREAM_END)
		goto inf_leave
	_31:
		;
		ret = -int32(3)
		goto inf_leave
	_32:
		;
		return -int32(4)
	_34:
		;
		/* fallthrough */
	_33:
		;
		return -int32(2)
	_35:
		;
		goto _1
	_1:
	}
	/*
	   Return from inflate(), updating the total counts and the check value.
	   If there was no progress during the inflate() call, return a buffer
	   error.  Call updatewindow() to create and/or update the window state.
	   Note: a memory error from inflate() is non-recoverable.
	*/
	goto inf_leave
inf_leave:
	;
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwsize != 0 || out != (*Tz_stream)(unsafe.Pointer(strm)).Favail_out && (*Tinflate_state)(unsafe.Pointer(state)).Fmode < int32(_BAD) && ((*Tinflate_state)(unsafe.Pointer(state)).Fmode < int32(_CHECK) || flush != int32(m_Z_FINISH)) {
		if _updatewindow(tls, strm, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out, out-(*Tz_stream)(unsafe.Pointer(strm)).Favail_out) != 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_MEM)
			return -int32(4)
		}
	}
	in -= (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	out -= (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	*(*TuLong)(unsafe.Pointer(strm + 16)) += uint64(in)
	*(*TuLong)(unsafe.Pointer(strm + 40)) += uint64(out)
	*(*uint64)(unsafe.Pointer(state + 40)) += uint64(out)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && out != 0 {
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
			v102 = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out-uintptr(out), out)
		} else {
			v102 = Xadler32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out-uintptr(out), out)
		}
		v101 = v102
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v101
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v101
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Flast != 0 {
		v103 = int32(64)
	} else {
		v103 = 0
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_TYPE) {
		v104 = int32(128)
	} else {
		v104 = 0
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_LEN_) || (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_COPY_) {
		v105 = int32(256)
	} else {
		v105 = 0
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fdata_type = libc.Int32FromUint32((*Tinflate_state)(unsafe.Pointer(state)).Fbits) + v103 + v104 + v105
	if (in == uint32(0) && out == uint32(0) || flush == int32(m_Z_FINISH)) && ret == m_Z_OK {
		ret = -int32(5)
	}
	return ret
}

var _order1 = [19]uint16{
	0:  uint16(16),
	1:  uint16(17),
	2:  uint16(18),
	4:  uint16(8),
	5:  uint16(7),
	6:  uint16(9),
	7:  uint16(6),
	8:  uint16(10),
	9:  uint16(5),
	10: uint16(11),
	11: uint16(4),
	12: uint16(12),
	13: uint16(3),
	14: uint16(13),
	15: uint16(2),
	16: uint16(14),
	17: uint16(1),
	18: uint16(15),
}

func XinflateEnd(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow)
	}
	(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	return m_Z_OK
}

func XinflateGetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength uintptr) (r int32) {
	var state uintptr
	_ = state
	/* check state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* copy dictionary */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave != 0 && dictionary != uintptr(m_Z_NULL) {
		libc.Xmemcpy(tls, dictionary, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), uint64((*Tinflate_state)(unsafe.Pointer(state)).Fwhave-(*Tinflate_state)(unsafe.Pointer(state)).Fwnext))
		libc.Xmemcpy(tls, dictionary+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwhave)-uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, uint64((*Tinflate_state)(unsafe.Pointer(state)).Fwnext))
	}
	if dictLength != uintptr(m_Z_NULL) {
		*(*TuInt)(unsafe.Pointer(dictLength)) = (*Tinflate_state)(unsafe.Pointer(state)).Fwhave
	}
	return m_Z_OK
}

func XinflateSetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength TuInt) (r int32) {
	var dictid uint64
	var ret int32
	var state uintptr
	_, _, _ = dictid, ret, state
	/* check state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fmode != int32(_DICT) {
		return -int32(2)
	}
	/* check for correct dictionary identifier */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_DICT) {
		dictid = Xadler32(tls, uint64(0), uintptr(m_Z_NULL), uint32(0))
		dictid = Xadler32(tls, dictid, dictionary, dictLength)
		if dictid != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck {
			return -int32(3)
		}
	}
	/* copy dictionary to window using updatewindow(), which will amend the
	   existing dictionary if appropriate */
	ret = _updatewindow(tls, strm, dictionary+uintptr(dictLength), dictLength)
	if ret != 0 {
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_MEM)
		return -int32(4)
	}
	(*Tinflate_state)(unsafe.Pointer(state)).Fhavedict = int32(1)
	return m_Z_OK
}

func XinflateGetHeader(tls *libc.TLS, strm Tz_streamp, head Tgz_headerp) (r int32) {
	var state uintptr
	_ = state
	/* check state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(2) == 0 {
		return -int32(2)
	}
	/* save header structure */
	(*Tinflate_state)(unsafe.Pointer(state)).Fhead = head
	(*Tgz_header)(unsafe.Pointer(head)).Fdone = 0
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Search buf[0..len-1] for the pattern: 0, 0, 0xff, 0xff.  Return when found
//	   or when out of input.  When called, *have is the number of pattern bytes
//	   found in order so far, in 0..3.  On return *have is updated to the new
//	   state.  If on return *have equals four, then the pattern was found and the
//	   return value is how many bytes were read including the last byte of the
//	   pattern.  If *have is less than four, then the pattern has not been found
//	   yet and the return value is len.  In the latter case, syncsearch() can be
//	   called again with more data and the *have state.  *have is initialized to
//	   zero for the first call.
//	 */
func _syncsearch(tls *libc.TLS, have uintptr, buf uintptr, len1 uint32) (r uint32) {
	var got, next uint32
	var v1 int32
	_, _, _ = got, next, v1
	got = *(*uint32)(unsafe.Pointer(have))
	next = uint32(0)
	for next < len1 && got < uint32(4) {
		if got < uint32(2) {
			v1 = 0
		} else {
			v1 = int32(0xff)
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(buf + uintptr(next)))) == v1 {
			got++
		} else {
			if *(*uint8)(unsafe.Pointer(buf + uintptr(next))) != 0 {
				got = uint32(0)
			} else {
				got = uint32(4) - got
			}
		}
		next++
	}
	*(*uint32)(unsafe.Pointer(have)) = got
	return next
}

func XinflateSync(tls *libc.TLS, strm Tz_streamp) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var flags int32
	var in, out uint64
	var len1, v1 uint32
	var state uintptr
	var _ /* buf at bp+0 */ [4]uint8
	_, _, _, _, _, _ = flags, in, len1, out, state, v1
	/* check parameters */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) && (*Tinflate_state)(unsafe.Pointer(state)).Fbits < uint32(8) {
		return -int32(5)
	}
	/* if first time, start search in bit buffer */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode != int32(_SYNC) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_SYNC)
		*(*uint64)(unsafe.Pointer(state + 80)) >>= uint64((*Tinflate_state)(unsafe.Pointer(state)).Fbits & uint32(7))
		*(*uint32)(unsafe.Pointer(state + 88)) -= (*Tinflate_state)(unsafe.Pointer(state)).Fbits & uint32(7)
		len1 = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fbits >= uint32(8) {
			v1 = len1
			len1++
			(*(*[4]uint8)(unsafe.Pointer(bp)))[v1] = uint8((*Tinflate_state)(unsafe.Pointer(state)).Fhold)
			*(*uint64)(unsafe.Pointer(state + 80)) >>= uint64(8)
			*(*uint32)(unsafe.Pointer(state + 88)) -= uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		_syncsearch(tls, state+140, bp, len1)
	}
	/* search available input */
	len1 = _syncsearch(tls, state+140, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, (*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
	*(*TuInt)(unsafe.Pointer(strm + 8)) -= len1
	*(*uintptr)(unsafe.Pointer(strm)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 16)) += uint64(len1)
	/* return no joy or set up to restart inflate() on a new block */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fhave != uint32(4) {
		return -int32(3)
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fflags == -int32(1) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fwrap = 0
	} else {
		*(*int32)(unsafe.Pointer(state + 16)) &= ^libc.Int32FromInt32(4)
	} /* no point in computing a check value now */
	flags = (*Tinflate_state)(unsafe.Pointer(state)).Fflags
	in = (*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in
	out = (*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out
	XinflateReset(tls, strm)
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = in
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = out
	(*Tinflate_state)(unsafe.Pointer(state)).Fflags = flags
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Returns true if inflate is currently at the end of a block generated by
//	   Z_SYNC_FLUSH or Z_FULL_FLUSH. This function is used by one PPP
//	   implementation to provide an additional safety check. PPP uses
//	   Z_SYNC_FLUSH but removes the length bytes of the resulting empty stored
//	   block. When decompressing, PPP checks that at the end of input packet,
//	   inflate is waiting for these length bytes.
//	 */
func XinflateSyncPoint(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	return libc.BoolInt32((*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_STORED) && (*Tinflate_state)(unsafe.Pointer(state)).Fbits == uint32(0))
}

func XinflateCopy(tls *libc.TLS, dest Tz_streamp, source Tz_streamp) (r int32) {
	var copy1, state, window uintptr
	var wsize uint32
	_, _, _, _ = copy1, state, window, wsize
	/* check input */
	if _inflateStateCheck(tls, source) != 0 || dest == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(source)).Fstate
	/* allocate space */
	copy1 = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), uint32(libc.Uint64FromInt64(7160)))
	if copy1 == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	window = uintptr(m_Z_NULL)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) {
		window = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fwbits, uint32(libc.Uint64FromInt64(1)))
		if window == uintptr(m_Z_NULL) {
			(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, copy1)
			return -int32(4)
		}
	}
	/* copy state */
	libc.Xmemcpy(tls, dest, source, uint64(112))
	libc.Xmemcpy(tls, copy1, state, uint64(7160))
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fstrm = dest
	if (*Tinflate_state)(unsafe.Pointer(state)).Flencode >= state+1368 && (*Tinflate_state)(unsafe.Pointer(state)).Flencode <= state+1368+uintptr(libc.Int32FromInt32(m_ENOUGH_LENS)+libc.Int32FromInt32(m_ENOUGH_DISTS))*4-uintptr(1)*4 {
		(*Tinflate_state)(unsafe.Pointer(copy1)).Flencode = copy1 + 1368 + uintptr((int64((*Tinflate_state)(unsafe.Pointer(state)).Flencode)-t__predefined_ptrdiff_t(state+1368))/4)*4
		(*Tinflate_state)(unsafe.Pointer(copy1)).Fdistcode = copy1 + 1368 + uintptr((int64((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode)-t__predefined_ptrdiff_t(state+1368))/4)*4
	}
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fnext = copy1 + 1368 + uintptr((int64((*Tinflate_state)(unsafe.Pointer(state)).Fnext)-t__predefined_ptrdiff_t(state+1368))/4)*4
	if window != uintptr(m_Z_NULL) {
		wsize = uint32(1) << (*Tinflate_state)(unsafe.Pointer(state)).Fwbits
		libc.Xmemcpy(tls, window, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, uint64(wsize))
	}
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fwindow = window
	(*Tz_stream)(unsafe.Pointer(dest)).Fstate = copy1
	return m_Z_OK
}

func XinflateUndermine(tls *libc.TLS, strm Tz_streamp, subvert int32) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	_ = subvert
	(*Tinflate_state)(unsafe.Pointer(state)).Fsane = int32(1)
	return -int32(3)
}

func XinflateValidate(tls *libc.TLS, strm Tz_streamp, check int32) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if check != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 {
		*(*int32)(unsafe.Pointer(state + 16)) |= int32(4)
	} else {
		*(*int32)(unsafe.Pointer(state + 16)) &= ^libc.Int32FromInt32(4)
	}
	return m_Z_OK
}

func XinflateMark(tls *libc.TLS, strm Tz_streamp) (r int64) {
	var state uintptr
	var v1, v2 uint32
	_, _, _ = state, v1, v2
	if _inflateStateCheck(tls, strm) != 0 {
		return -(libc.Int64FromInt64(1) << libc.Int32FromInt32(16))
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_COPY) {
		v1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
	} else {
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_MATCH) {
			v2 = (*Tinflate_state)(unsafe.Pointer(state)).Fwas - (*Tinflate_state)(unsafe.Pointer(state)).Flength
		} else {
			v2 = uint32(0)
		}
		v1 = v2
	}
	return libc.Int64FromUint64(libc.Uint64FromInt64(int64((*Tinflate_state)(unsafe.Pointer(state)).Fback))<<libc.Int32FromInt32(16)) + libc.Int64FromUint32(v1)
}

func XinflateCodesUsed(tls *libc.TLS, strm Tz_streamp) (r uint64) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return libc.Uint64FromInt32(-libc.Int32FromInt32(1))
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	return libc.Uint64FromInt64((int64((*Tinflate_state)(unsafe.Pointer(state)).Fnext) - t__predefined_ptrdiff_t(state+1368)) / 4)
}

const m_MAXBITS = 15

/*
  If you use the zlib library in a product, an acknowledgment is welcome
  in the documentation of your product. If for some reason you cannot
  include such an acknowledgment, I would appreciate that you keep this
  copyright string in the executable of your product.
*/

// C documentation
//
//	/*
//	   Build a set of tables to decode the provided canonical Huffman code.
//	   The code lengths are lens[0..codes-1].  The result starts at *table,
//	   whose indices are 0..2^bits-1.  work is a writable array of at least
//	   lens shorts, which is used as a work area.  type is the type of code
//	   to be generated, CODES, LENS, or DISTS.  On return, zero is success,
//	   -1 is an invalid code, and +1 means that ENOUGH isn't enough.  table
//	   on return points to the next available entry's address.  bits is the
//	   requested root table index bits, and on return it is the actual root
//	   table index bits.  It will differ if the request is greater than the
//	   longest code or if it is less than the shortest code.
//	 */
func Xinflate_table(tls *libc.TLS, type1 Tcodetype, lens uintptr, codes uint32, table uintptr, bits uintptr, work uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var base, extra, next, v13, v14, v17, v4, v5, v6, v7 uintptr
	var curr, drop, fill, huff, incr, len1, low, mask, match, max, min, root, sym, used uint32
	var here Tcode
	var left int32
	var v12, v16 uint16
	var _ /* count at bp+0 */ [16]uint16
	var _ /* offs at bp+32 */ [16]uint16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, curr, drop, extra, fill, here, huff, incr, left, len1, low, mask, match, max, min, next, root, sym, used, v12, v13, v14, v16, v17, v4, v5, v6, v7 /* offsets in table for each length */
	/*
	   Process a set of code lengths to create a canonical Huffman code.  The
	   code lengths are lens[0..codes-1].  Each length corresponds to the
	   symbols 0..codes-1.  The Huffman code is generated by first sorting the
	   symbols by length from short to long, and retaining the symbol order
	   for codes with equal lengths.  Then the code starts with all zero bits
	   for the first code of the shortest length, and the codes are integer
	   increments for the same length, and zeros are appended as the length
	   increases.  For the deflate format, these bits are stored backwards
	   from their more natural integer increment ordering, and so when the
	   decoding tables are built in the large loop below, the integer codes
	   are incremented backwards.
	   This routine assumes, but does not check, that all of the entries in
	   lens[] are in the range 0..MAXBITS.  The caller must assure this.
	   1..MAXBITS is interpreted as that code length.  zero means that that
	   symbol does not occur in this code.
	   The codes are sorted by computing a count of codes for each length,
	   creating from that a table of starting indices for each length in the
	   sorted table, and then entering the symbols in order in the sorted
	   table.  The sorted table is work[], with that space being provided by
	   the caller.
	   The length counts are used for other purposes as well, i.e. finding
	   the minimum and maximum length codes, determining if there are any
	   codes at all, checking for a valid set of lengths, and looking ahead
	   at length counts to determine sub-table sizes when building the
	   decoding tables.
	*/
	/* accumulate lengths for codes (assumes lens[] all in 0..MAXBITS) */
	len1 = uint32(0)
	for {
		if !(len1 <= uint32(m_MAXBITS)) {
			break
		}
		(*(*[16]uint16)(unsafe.Pointer(bp)))[len1] = uint16(0)
		goto _1
	_1:
		;
		len1++
	}
	sym = uint32(0)
	for {
		if !(sym < codes) {
			break
		}
		(*(*[16]uint16)(unsafe.Pointer(bp)))[*(*uint16)(unsafe.Pointer(lens + uintptr(sym)*2))]++
		goto _2
	_2:
		;
		sym++
	}
	/* bound code lengths, force root to be within code lengths */
	root = *(*uint32)(unsafe.Pointer(bits))
	max = uint32(m_MAXBITS)
	for {
		if !(max >= uint32(1)) {
			break
		}
		if libc.Int32FromUint16((*(*[16]uint16)(unsafe.Pointer(bp)))[max]) != 0 {
			break
		}
		goto _3
	_3:
		;
		max--
	}
	if root > max {
		root = max
	}
	if max == uint32(0) { /* no symbols to code at all */
		here.Fop = libc.Uint8FromInt32(64) /* invalid code marker */
		here.Fbits = libc.Uint8FromInt32(1)
		here.Fval = libc.Uint16FromInt32(0)
		v5 = table
		v4 = *(*uintptr)(unsafe.Pointer(v5))
		*(*uintptr)(unsafe.Pointer(v5)) += 4
		*(*Tcode)(unsafe.Pointer(v4)) = here /* make a table to force an error */
		v7 = table
		v6 = *(*uintptr)(unsafe.Pointer(v7))
		*(*uintptr)(unsafe.Pointer(v7)) += 4
		*(*Tcode)(unsafe.Pointer(v6)) = here
		*(*uint32)(unsafe.Pointer(bits)) = uint32(1)
		return 0 /* no symbols, but wait for decoding to report error */
	}
	min = uint32(1)
	for {
		if !(min < max) {
			break
		}
		if libc.Int32FromUint16((*(*[16]uint16)(unsafe.Pointer(bp)))[min]) != 0 {
			break
		}
		goto _8
	_8:
		;
		min++
	}
	if root < min {
		root = min
	}
	/* check for an over-subscribed or incomplete set of lengths */
	left = int32(1)
	len1 = uint32(1)
	for {
		if !(len1 <= uint32(m_MAXBITS)) {
			break
		}
		left <<= int32(1)
		left -= libc.Int32FromUint16((*(*[16]uint16)(unsafe.Pointer(bp)))[len1])
		if left < 0 {
			return -int32(1)
		} /* over-subscribed */
		goto _9
	_9:
		;
		len1++
	}
	if left > 0 && (type1 == int32(_CODES) || max != uint32(1)) {
		return -int32(1)
	} /* incomplete set */
	/* generate offsets into symbol table for each length for sorting */
	(*(*[16]uint16)(unsafe.Pointer(bp + 32)))[int32(1)] = uint16(0)
	len1 = uint32(1)
	for {
		if !(len1 < uint32(m_MAXBITS)) {
			break
		}
		(*(*[16]uint16)(unsafe.Pointer(bp + 32)))[len1+uint32(1)] = libc.Uint16FromInt32(libc.Int32FromUint16((*(*[16]uint16)(unsafe.Pointer(bp + 32)))[len1]) + libc.Int32FromUint16((*(*[16]uint16)(unsafe.Pointer(bp)))[len1]))
		goto _10
	_10:
		;
		len1++
	}
	/* sort symbols by length, by symbol order within each length */
	sym = uint32(0)
	for {
		if !(sym < codes) {
			break
		}
		if libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(lens + uintptr(sym)*2))) != 0 {
			v13 = bp + 32 + uintptr(*(*uint16)(unsafe.Pointer(lens + uintptr(sym)*2)))*2
			v12 = *(*uint16)(unsafe.Pointer(v13))
			*(*uint16)(unsafe.Pointer(v13))++
			*(*uint16)(unsafe.Pointer(work + uintptr(v12)*2)) = uint16(sym)
		}
		goto _11
	_11:
		;
		sym++
	}
	/*
	   Create and fill in decoding tables.  In this loop, the table being
	   filled is at next and has curr index bits.  The code being used is huff
	   with length len.  That code is converted to an index by dropping drop
	   bits off of the bottom.  For codes where len is less than drop + curr,
	   those top drop + curr - len bits are incremented through all values to
	   fill the table with replicated entries.
	   root is the number of index bits for the root table.  When len exceeds
	   root, sub-tables are created pointed to by the root entry with an index
	   of the low root bits of huff.  This is saved in low to check for when a
	   new sub-table should be started.  drop is zero when the root table is
	   being filled, and drop is root when sub-tables are being filled.
	   When a new sub-table is needed, it is necessary to look ahead in the
	   code lengths to determine what size sub-table is needed.  The length
	   counts are used for this, and so count[] is decremented as codes are
	   entered in the tables.
	   used keeps track of how many table entries have been allocated from the
	   provided *table space.  It is checked for LENS and DIST tables against
	   the constants ENOUGH_LENS and ENOUGH_DISTS to guard against changes in
	   the initial root table size constants.  See the comments in inftrees.h
	   for more information.
	   sym increments through all symbols, and the loop terminates when
	   all codes of length max, i.e. all codes, have been processed.  This
	   routine permits incomplete codes, so another loop after this one fills
	   in the rest of the decoding tables with invalid code markers.
	*/
	/* set up for code type */
	switch type1 {
	case int32(_CODES):
		v14 = work
		extra = v14
		base = v14 /* dummy value--not used */
		match = uint32(20)
	case int32(_LENS):
		base = uintptr(unsafe.Pointer(&_lbase))
		extra = uintptr(unsafe.Pointer(&_lext))
		match = uint32(257)
	default: /* DISTS */
		base = uintptr(unsafe.Pointer(&_dbase))
		extra = uintptr(unsafe.Pointer(&_dext))
		match = uint32(0)
	}
	/* initialize state for loop */
	huff = uint32(0)                                    /* starting code */
	sym = uint32(0)                                     /* starting code symbol */
	len1 = min                                          /* starting code length */
	next = *(*uintptr)(unsafe.Pointer(table))           /* current table to fill in */
	curr = root                                         /* current table index bits */
	drop = uint32(0)                                    /* current bits to drop from code for index */
	low = libc.Uint32FromInt32(-libc.Int32FromInt32(1)) /* trigger new sub-table when len > root */
	used = uint32(1) << root                            /* use root table entries */
	mask = used - uint32(1)                             /* mask for comparing low */
	/* check available table space */
	if type1 == int32(_LENS) && used > uint32(m_ENOUGH_LENS) || type1 == int32(_DISTS) && used > uint32(m_ENOUGH_DISTS) {
		return int32(1)
	}
	/* process all codes and make table entries */
	for {
		/* create table entry */
		here.Fbits = uint8(len1 - drop)
		if uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))+uint32(1) < match {
			here.Fop = libc.Uint8FromInt32(0)
			here.Fval = *(*uint16)(unsafe.Pointer(work + uintptr(sym)*2))
		} else {
			if uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2))) >= match {
				here.Fop = uint8(*(*uint16)(unsafe.Pointer(extra + uintptr(uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))-match)*2)))
				here.Fval = *(*uint16)(unsafe.Pointer(base + uintptr(uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))-match)*2))
			} else {
				here.Fop = libc.Uint8FromInt32(libc.Int32FromInt32(32) + libc.Int32FromInt32(64)) /* end of block */
				here.Fval = uint16(0)
			}
		}
		/* replicate for those indices with low len bits equal to huff */
		incr = uint32(1) << (len1 - drop)
		fill = uint32(1) << curr
		min = fill /* save offset to next table */
		for cond := true; cond; cond = fill != uint32(0) {
			fill -= incr
			*(*Tcode)(unsafe.Pointer(next + uintptr(huff>>drop+fill)*4)) = here
		}
		/* backwards increment the len-bit code huff */
		incr = uint32(1) << (len1 - uint32(1))
		for huff&incr != 0 {
			incr >>= uint32(1)
		}
		if incr != uint32(0) {
			huff &= incr - uint32(1)
			huff += incr
		} else {
			huff = uint32(0)
		}
		/* go to next symbol, update count, len */
		sym++
		v17 = bp + uintptr(len1)*2
		*(*uint16)(unsafe.Pointer(v17))--
		v16 = *(*uint16)(unsafe.Pointer(v17))
		if libc.Int32FromUint16(v16) == 0 {
			if len1 == max {
				break
			}
			len1 = uint32(*(*uint16)(unsafe.Pointer(lens + uintptr(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))*2)))
		}
		/* create new sub-table if needed */
		if len1 > root && huff&mask != low {
			/* if first time, transition to sub-tables */
			if drop == uint32(0) {
				drop = root
			}
			/* increment past last table */
			next += uintptr(min) * 4 /* here min is 1 << curr */
			/* determine length of next table */
			curr = len1 - drop
			left = libc.Int32FromInt32(1) << curr
			for curr+drop < max {
				left -= libc.Int32FromUint16((*(*[16]uint16)(unsafe.Pointer(bp)))[curr+drop])
				if left <= 0 {
					break
				}
				curr++
				left <<= int32(1)
			}
			/* check for enough space */
			used += uint32(1) << curr
			if type1 == int32(_LENS) && used > uint32(m_ENOUGH_LENS) || type1 == int32(_DISTS) && used > uint32(m_ENOUGH_DISTS) {
				return int32(1)
			}
			/* point entry in root table to sub-table */
			low = huff & mask
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fop = uint8(curr)
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fbits = uint8(root)
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fval = libc.Uint16FromInt64((int64(next) - int64(*(*uintptr)(unsafe.Pointer(table)))) / 4)
		}
		goto _15
	_15:
	}
	/* fill in remaining table entry if code is incomplete (guaranteed to have
	   at most one remaining entry, since if the code is incomplete, the
	   maximum code length that was allowed to get this far is one bit) */
	if huff != uint32(0) {
		here.Fop = libc.Uint8FromInt32(64) /* invalid code marker */
		here.Fbits = uint8(len1 - drop)
		here.Fval = libc.Uint16FromInt32(0)
		*(*Tcode)(unsafe.Pointer(next + uintptr(huff)*4)) = here
	}
	/* set return parameters */
	*(*uintptr)(unsafe.Pointer(table)) += uintptr(used) * 4
	*(*uint32)(unsafe.Pointer(bits)) = root
	return 0
}

var _lbase = [31]uint16{
	0:  uint16(3),
	1:  uint16(4),
	2:  uint16(5),
	3:  uint16(6),
	4:  uint16(7),
	5:  uint16(8),
	6:  uint16(9),
	7:  uint16(10),
	8:  uint16(11),
	9:  uint16(13),
	10: uint16(15),
	11: uint16(17),
	12: uint16(19),
	13: uint16(23),
	14: uint16(27),
	15: uint16(31),
	16: uint16(35),
	17: uint16(43),
	18: uint16(51),
	19: uint16(59),
	20: uint16(67),
	21: uint16(83),
	22: uint16(99),
	23: uint16(115),
	24: uint16(131),
	25: uint16(163),
	26: uint16(195),
	27: uint16(227),
	28: uint16(258),
}

var _lext = [31]uint16{
	0:  uint16(16),
	1:  uint16(16),
	2:  uint16(16),
	3:  uint16(16),
	4:  uint16(16),
	5:  uint16(16),
	6:  uint16(16),
	7:  uint16(16),
	8:  uint16(17),
	9:  uint16(17),
	10: uint16(17),
	11: uint16(17),
	12: uint16(18),
	13: uint16(18),
	14: uint16(18),
	15: uint16(18),
	16: uint16(19),
	17: uint16(19),
	18: uint16(19),
	19: uint16(19),
	20: uint16(20),
	21: uint16(20),
	22: uint16(20),
	23: uint16(20),
	24: uint16(21),
	25: uint16(21),
	26: uint16(21),
	27: uint16(21),
	28: uint16(16),
	29: uint16(203),
	30: uint16(77),
}

var _dbase = [32]uint16{
	0:  uint16(1),
	1:  uint16(2),
	2:  uint16(3),
	3:  uint16(4),
	4:  uint16(5),
	5:  uint16(7),
	6:  uint16(9),
	7:  uint16(13),
	8:  uint16(17),
	9:  uint16(25),
	10: uint16(33),
	11: uint16(49),
	12: uint16(65),
	13: uint16(97),
	14: uint16(129),
	15: uint16(193),
	16: uint16(257),
	17: uint16(385),
	18: uint16(513),
	19: uint16(769),
	20: uint16(1025),
	21: uint16(1537),
	22: uint16(2049),
	23: uint16(3073),
	24: uint16(4097),
	25: uint16(6145),
	26: uint16(8193),
	27: uint16(12289),
	28: uint16(16385),
	29: uint16(24577),
}

var _dext = [32]uint16{
	0:  uint16(16),
	1:  uint16(16),
	2:  uint16(16),
	3:  uint16(16),
	4:  uint16(17),
	5:  uint16(17),
	6:  uint16(18),
	7:  uint16(18),
	8:  uint16(19),
	9:  uint16(19),
	10: uint16(20),
	11: uint16(20),
	12: uint16(21),
	13: uint16(21),
	14: uint16(22),
	15: uint16(22),
	16: uint16(23),
	17: uint16(23),
	18: uint16(24),
	19: uint16(24),
	20: uint16(25),
	21: uint16(25),
	22: uint16(26),
	23: uint16(26),
	24: uint16(27),
	25: uint16(27),
	26: uint16(28),
	27: uint16(28),
	28: uint16(29),
	29: uint16(29),
	30: uint16(64),
	31: uint16(64),
}

const m_DIST_CODE_LEN = 512
const m_END_BLOCK = 256
const m_MAX_BL_BITS = 7
const m_REPZ_11_138 = 18
const m_REPZ_3_10 = 17
const m_REP_3_6 = 16
const m_SMALLEST = 1
const m_TCONST = "const"

type Tstatic_tree_desc = struct {
	Fstatic_tree uintptr
	Fextra_bits  uintptr
	Fextra_base  int32
	Felems       int32
	Fmax_length  int32
}

type Tstatic_tree_desc_s = Tstatic_tree_desc

/* ===========================================================================
 * Constants
 */

/* Bit length codes must not exceed MAX_BL_BITS bits */

/* end of block literal code */

/* repeat previous bit length 3-6 times (2 bits of repeat count) */

/* repeat a zero length 3-10 times  (3 bits of repeat count) */

/* repeat a zero length 11-138 times  (7 bits of repeat count) */
var _extra_lbits = [29]int32{
	8:  int32(1),
	9:  int32(1),
	10: int32(1),
	11: int32(1),
	12: int32(2),
	13: int32(2),
	14: int32(2),
	15: int32(2),
	16: int32(3),
	17: int32(3),
	18: int32(3),
	19: int32(3),
	20: int32(4),
	21: int32(4),
	22: int32(4),
	23: int32(4),
	24: int32(5),
	25: int32(5),
	26: int32(5),
	27: int32(5),
}

var _extra_dbits = [30]int32{
	4:  int32(1),
	5:  int32(1),
	6:  int32(2),
	7:  int32(2),
	8:  int32(3),
	9:  int32(3),
	10: int32(4),
	11: int32(4),
	12: int32(5),
	13: int32(5),
	14: int32(6),
	15: int32(6),
	16: int32(7),
	17: int32(7),
	18: int32(8),
	19: int32(8),
	20: int32(9),
	21: int32(9),
	22: int32(10),
	23: int32(10),
	24: int32(11),
	25: int32(11),
	26: int32(12),
	27: int32(12),
	28: int32(13),
	29: int32(13),
}

var _extra_blbits = [19]int32{
	16: int32(2),
	17: int32(3),
	18: int32(7),
}

var _bl_order = [19]Tuch{
	0:  uint8(16),
	1:  uint8(17),
	2:  uint8(18),
	4:  uint8(8),
	5:  uint8(7),
	6:  uint8(9),
	7:  uint8(6),
	8:  uint8(10),
	9:  uint8(5),
	10: uint8(11),
	11: uint8(4),
	12: uint8(12),
	13: uint8(3),
	14: uint8(13),
	15: uint8(2),
	16: uint8(14),
	17: uint8(1),
	18: uint8(15),
}
var _static_ltree = [288]Tct_data{
	0: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(12)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	1: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(140)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	2: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(76)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	3: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(204)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	4: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(44)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	5: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(172)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	6: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(108)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	7: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(236)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	8: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(28)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	9: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(156)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	10: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(92)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	11: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(220)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	12: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(60)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	13: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(188)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	14: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(124)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	15: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(252)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	16: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(2)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	17: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(130)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	18: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(66)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	19: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(194)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	20: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(34)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	21: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(162)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	22: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(98)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	23: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(226)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	24: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(18)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	25: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(146)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	26: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(82)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	27: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(210)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	28: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(50)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	29: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(178)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	30: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(114)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	31: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(242)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	32: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(10)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	33: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(138)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	34: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(74)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	35: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(202)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	36: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(42)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	37: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(170)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	38: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(106)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	39: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(234)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	40: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(26)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	41: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(154)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	42: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(90)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	43: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(218)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	44: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(58)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	45: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(186)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	46: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(122)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	47: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(250)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	48: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(6)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	49: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(134)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	50: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(70)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	51: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(198)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	52: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(38)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	53: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(166)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	54: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(102)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	55: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(230)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	56: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(22)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	57: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(150)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	58: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(86)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	59: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(214)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	60: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(54)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	61: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(182)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	62: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(118)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	63: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(246)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	64: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(14)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	65: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(142)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	66: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(78)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	67: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(206)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	68: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(46)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	69: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(174)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	70: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(110)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	71: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(238)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	72: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(30)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	73: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(158)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	74: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(94)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	75: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(222)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	76: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(62)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	77: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(190)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	78: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(126)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	79: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(254)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	80: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(1)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	81: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(129)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	82: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(65)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	83: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(193)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	84: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(33)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	85: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(161)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	86: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(97)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	87: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(225)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	88: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(17)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	89: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(145)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	90: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(81)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	91: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(209)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	92: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(49)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	93: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(177)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	94: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(113)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	95: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(241)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	96: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	97: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(137)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	98: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(73)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	99: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(201)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	100: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(41)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	101: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(169)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	102: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(105)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	103: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(233)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	104: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(25)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	105: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(153)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	106: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(89)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	107: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(217)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	108: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(57)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	109: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(185)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	110: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(121)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	111: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(249)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	112: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	113: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(133)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	114: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(69)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	115: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(197)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	116: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(37)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	117: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(165)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	118: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(101)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	119: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(229)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	120: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(21)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	121: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(149)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	122: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(85)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	123: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(213)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	124: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(53)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	125: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(181)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	126: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(117)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	127: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(245)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	128: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(13)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	129: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(141)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	130: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(77)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	131: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(205)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	132: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(45)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	133: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(173)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	134: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(109)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	135: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(237)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	136: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(29)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	137: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(157)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	138: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(93)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	139: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(221)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	140: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(61)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	141: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(189)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	142: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(125)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	143: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(253)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	144: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(19)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	145: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(275)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	146: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(147)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	147: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(403)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	148: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(83)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	149: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(339)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	150: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(211)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	151: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(467)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	152: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(51)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	153: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(307)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	154: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(179)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	155: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(435)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	156: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(115)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	157: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(371)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	158: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(243)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	159: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(499)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	160: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(11)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	161: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(267)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	162: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(139)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	163: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(395)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	164: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(75)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	165: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(331)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	166: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(203)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	167: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(459)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	168: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(43)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	169: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(299)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	170: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(171)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	171: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(427)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	172: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(107)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	173: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(363)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	174: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(235)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	175: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(491)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	176: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(27)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	177: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(283)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	178: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(155)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	179: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(411)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	180: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(91)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	181: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(347)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	182: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(219)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	183: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(475)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	184: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(59)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	185: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(315)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	186: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(187)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	187: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(443)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	188: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(123)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	189: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(379)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	190: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(251)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	191: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(507)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	192: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	193: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(263)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	194: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(135)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	195: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(391)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	196: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(71)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	197: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(327)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	198: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(199)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	199: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(455)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	200: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(39)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	201: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(295)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	202: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(167)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	203: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(423)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	204: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(103)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	205: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(359)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	206: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(231)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	207: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(487)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	208: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(23)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	209: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(279)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	210: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(151)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	211: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(407)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	212: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(87)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	213: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(343)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	214: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(215)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	215: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(471)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	216: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(55)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	217: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(311)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	218: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(183)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	219: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(439)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	220: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(119)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	221: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(375)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	222: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(247)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	223: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(503)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	224: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(15)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	225: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(271)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	226: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(143)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	227: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(399)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	228: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(79)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	229: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(335)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	230: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(207)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	231: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(463)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	232: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(47)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	233: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(303)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	234: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(175)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	235: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(431)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	236: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(111)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	237: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(367)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	238: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(239)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	239: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(495)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	240: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(31)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	241: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(287)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	242: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(159)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	243: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(415)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	244: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(95)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	245: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(351)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	246: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(223)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	247: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(479)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	248: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(63)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	249: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(319)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	250: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(191)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	251: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(447)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	252: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(127)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	253: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(383)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	254: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(255)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	255: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(511)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	256: {
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	257: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(64)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	258: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(32)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	259: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(96)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	260: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(16)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	261: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(80)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	262: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(48)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	263: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(112)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	264: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	265: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(72)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	266: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(40)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	267: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(104)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	268: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(24)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	269: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(88)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	270: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(56)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	271: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(120)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	272: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(4)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	273: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(68)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	274: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(36)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	275: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(100)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	276: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(20)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	277: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(84)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	278: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(52)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	279: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(116)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	280: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(3)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	281: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(131)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	282: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(67)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	283: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(195)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	284: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(35)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	285: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(163)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	286: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(99)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	287: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(227)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
}
var _static_dtree = [30]Tct_data{
	0: {
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	1: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(16)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	2: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	3: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(24)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	4: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(4)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	5: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(20)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	6: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(12)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	7: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(28)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	8: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(2)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	9: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(18)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	10: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(10)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	11: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(26)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	12: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(6)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	13: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(22)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	14: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(14)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	15: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(30)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	16: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(1)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	17: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(17)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	18: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	19: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(25)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	20: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	21: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(21)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	22: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(13)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	23: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(29)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	24: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(3)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	25: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(19)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	26: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(11)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	27: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(27)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	28: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	29: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(23)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
}
var _base_length = [29]int32{
	1:  int32(1),
	2:  int32(2),
	3:  int32(3),
	4:  int32(4),
	5:  int32(5),
	6:  int32(6),
	7:  int32(7),
	8:  int32(8),
	9:  int32(10),
	10: int32(12),
	11: int32(14),
	12: int32(16),
	13: int32(20),
	14: int32(24),
	15: int32(28),
	16: int32(32),
	17: int32(40),
	18: int32(48),
	19: int32(56),
	20: int32(64),
	21: int32(80),
	22: int32(96),
	23: int32(112),
	24: int32(128),
	25: int32(160),
	26: int32(192),
	27: int32(224),
}
var _base_dist = [30]int32{
	1:  int32(1),
	2:  int32(2),
	3:  int32(3),
	4:  int32(4),
	5:  int32(6),
	6:  int32(8),
	7:  int32(12),
	8:  int32(16),
	9:  int32(24),
	10: int32(32),
	11: int32(48),
	12: int32(64),
	13: int32(96),
	14: int32(128),
	15: int32(192),
	16: int32(256),
	17: int32(384),
	18: int32(512),
	19: int32(768),
	20: int32(1024),
	21: int32(1536),
	22: int32(2048),
	23: int32(3072),
	24: int32(4096),
	25: int32(6144),
	26: int32(8192),
	27: int32(12288),
	28: int32(16384),
	29: int32(24576),
}

var _static_l_desc = Tstatic_tree_desc{
	Fstatic_tree: uintptr(unsafe.Pointer(&_static_ltree)),
	Fextra_bits:  uintptr(unsafe.Pointer(&_extra_lbits)),
	Fextra_base:  libc.Int32FromInt32(m_LITERALS) + libc.Int32FromInt32(1),
	Felems:       libc.Int32FromInt32(m_LITERALS) + libc.Int32FromInt32(1) + libc.Int32FromInt32(m_LENGTH_CODES),
	Fmax_length:  int32(m_MAX_BITS),
}

var _static_d_desc = Tstatic_tree_desc{
	Fstatic_tree: uintptr(unsafe.Pointer(&_static_dtree)),
	Fextra_bits:  uintptr(unsafe.Pointer(&_extra_dbits)),
	Felems:       int32(m_D_CODES),
	Fmax_length:  int32(m_MAX_BITS),
}

var _static_bl_desc = Tstatic_tree_desc{
	Fextra_bits: uintptr(unsafe.Pointer(&_extra_blbits)),
	Felems:      int32(m_BL_CODES),
	Fmax_length: int32(m_MAX_BL_BITS),
}

/* ===========================================================================
 * Output a short LSB first on the stream.
 * IN assertion: there is enough room in pendingBuf.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Reverse the first len bits of a code, using straightforward code (a faster
//	 * method would use a table)
//	 * IN assertion: 1 <= len <= 15
//	 */
func _bi_reverse(tls *libc.TLS, code uint32, len1 int32) (r uint32) {
	var res uint32
	var v1 int32
	_, _ = res, v1
	res = uint32(0)
	for {
		res |= code & uint32(1)
		code >>= uint32(1)
		res <<= uint32(1)
		goto _2
	_2:
		;
		len1--
		v1 = len1
		if !(v1 > 0) {
			break
		}
	}
	return res >> int32(1)
}

// C documentation
//
//	/* ===========================================================================
//	 * Flush the bit buffer, keeping at most 7 bits in it.
//	 */
func _bi_flush(tls *libc.TLS, s uintptr) {
	var v1, v3, v5 Tulg
	var v2, v4, v6, p7 uintptr
	_, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, p7
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid == int32(16) {
		v2 = s + 40
		v1 = *(*Tulg)(unsafe.Pointer(v2))
		*(*Tulg)(unsafe.Pointer(v2))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v4 = s + 40
		v3 = *(*Tulg)(unsafe.Pointer(v4))
		*(*Tulg)(unsafe.Pointer(v4))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid >= int32(8) {
			v6 = s + 40
			v5 = *(*Tulg)(unsafe.Pointer(v6))
			*(*Tulg)(unsafe.Pointer(v6))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf)
			p7 = s + 5936
			*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) >> libc.Int32FromInt32(8))
			*(*int32)(unsafe.Pointer(s + 5940)) -= int32(8)
		}
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Flush the bit buffer and align the output on a byte boundary
//	 */
func _bi_windup(tls *libc.TLS, s uintptr) {
	var v1, v3, v5 Tulg
	var v2, v4, v6 uintptr
	_, _, _, _, _, _ = v1, v2, v3, v4, v5, v6
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > int32(8) {
		v2 = s + 40
		v1 = *(*Tulg)(unsafe.Pointer(v2))
		*(*Tulg)(unsafe.Pointer(v2))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v4 = s + 40
		v3 = *(*Tulg)(unsafe.Pointer(v4))
		*(*Tulg)(unsafe.Pointer(v4))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > 0 {
			v6 = s + 40
			v5 = *(*Tulg)(unsafe.Pointer(v6))
			*(*Tulg)(unsafe.Pointer(v6))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf)
		}
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
}

// C documentation
//
//	/* ===========================================================================
//	 * Generate the codes for a given tree and bit counts (which need not be
//	 * optimal).
//	 * IN assertion: the array bl_count contains the bit length statistics for
//	 * the given tree and the field len is set for all tree elements.
//	 * OUT assertion: the field code is set for all tree elements of non
//	 *     zero code length.
//	 */
func _gen_codes(tls *libc.TLS, tree uintptr, max_code int32, bl_count uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bits, len1, n int32
	var code uint32
	var v3 Tush
	var v4 uintptr
	var _ /* next_code at bp+0 */ [16]Tush
	_, _, _, _, _, _ = bits, code, len1, n, v3, v4 /* next code value for each bit length */
	code = uint32(0)                               /* code index */
	/* The distribution counts are first used to generate the code values
	 * without bit reversal.
	 */
	bits = int32(1)
	for {
		if !(bits <= int32(m_MAX_BITS)) {
			break
		}
		code = (code + uint32(*(*Tushf)(unsafe.Pointer(bl_count + uintptr(bits-int32(1))*2)))) << int32(1)
		(*(*[16]Tush)(unsafe.Pointer(bp)))[bits] = uint16(code)
		goto _1
	_1:
		;
		bits++
	}
	/* Check that the bit counts in bl_count are consistent. The last code
	 * must be all ones.
	 */
	n = 0
	for {
		if !(n <= max_code) {
			break
		}
		len1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)))
		if len1 == 0 {
			goto _2
		}
		/* Now reverse the bits */
		v4 = bp + uintptr(len1)*2
		v3 = *(*Tush)(unsafe.Pointer(v4))
		*(*Tush)(unsafe.Pointer(v4))++
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4)) = uint16(_bi_reverse(tls, uint32(v3), len1))
		goto _2
	_2:
		;
		n++
	}
}

/* Send a code of the given tree. c and tree must not have side effects */

/* ===========================================================================
 * Send a value on a given number of bits.
 * IN assertion: length <= 16 and value fits in length bits.
 */

/* the arguments must not have side effects */

// C documentation
//
//	/* ===========================================================================
//	 * Initialize the various 'constant' tables.
//	 */
func _tr_static_init(tls *libc.TLS) {
}

/* ===========================================================================
 * Generate the file trees.h describing the static trees.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Initialize a new block.
//	 */
func _init_block(tls *libc.TLS, s uintptr) {
	var n int32
	var v4 Tulg
	var v5 TuInt
	_, _, _ = n, v4, v5 /* iterates over tree elements */
	/* Initialize the trees. */
	n = 0
	for {
		if !(n < libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 212 + uintptr(n)*4)) = uint16(0)
		goto _1
	_1:
		;
		n++
	}
	n = 0
	for {
		if !(n < int32(m_D_CODES)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 2504 + uintptr(n)*4)) = uint16(0)
		goto _2
	_2:
		;
		n++
	}
	n = 0
	for {
		if !(n < int32(m_BL_CODES)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(n)*4)) = uint16(0)
		goto _3
	_3:
		;
		n++
	}
	*(*Tush)(unsafe.Pointer(s + 212 + 256*4)) = uint16(1)
	v4 = libc.Uint64FromInt64(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstatic_len = v4
	(*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len = v4
	v5 = libc.Uint32FromInt32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches = v5
	(*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next = v5
}

// C documentation
//
//	/* ===========================================================================
//	 * Initialize the tree data structures for a new zlib stream.
//	 */
func x__tr_init(tls *libc.TLS, s uintptr) {
	_tr_static_init(tls)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fdyn_tree = s + 212
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_l_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fdyn_tree = s + 2504
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_d_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbl_desc.Fdyn_tree = s + 2748
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbl_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_bl_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
	/* Initialize the first block of the first file: */
	_init_block(tls, s)
}

/* Index within the heap array of least frequent node in the Huffman tree */

/* ===========================================================================
 * Remove the smallest element from the heap and recreate the heap with
 * one less element. Updates heap and heap_len.
 */

/* ===========================================================================
 * Compares to subtrees, using the tree depth as tie breaker when
 * the subtrees have equal frequency. This minimizes the worst case length.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Restore the heap property by moving down the tree starting at node k,
//	 * exchanging a node with the smallest of its two sons if necessary, stopping
//	 * when the heap property is re-established (each father smaller than its
//	 * two sons).
//	 */
func _pqdownheap(tls *libc.TLS, s uintptr, tree uintptr, k int32) {
	var j, v int32
	_, _ = j, v
	v = *(*int32)(unsafe.Pointer(s + 3008 + uintptr(k)*4))
	j = k << int32(1) /* left son of k */
	for j <= (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len {
		/* Set j to the smallest of the two sons: */
		if j < (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len && (libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j+int32(1))*4)))*4))) < libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4)))*4))) || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j+int32(1))*4)))*4))) == libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4)))*4))) && libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j+int32(1))*4)))))) <= libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4))))))) {
			j++
		}
		/* Exit if v is smaller than both sons */
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(v)*4))) < libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4)))*4))) || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(v)*4))) == libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4)))*4))) && libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(v)))) <= libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4)))))) {
			break
		}
		/* Exchange v with the smallest son */
		*(*int32)(unsafe.Pointer(s + 3008 + uintptr(k)*4)) = *(*int32)(unsafe.Pointer(s + 3008 + uintptr(j)*4))
		k = j
		/* And continue down the tree, setting j to the left son of k */
		j <<= int32(1)
	}
	*(*int32)(unsafe.Pointer(s + 3008 + uintptr(k)*4)) = v
}

// C documentation
//
//	/* ===========================================================================
//	 * Compute the optimal bit lengths for a tree and update the total bit length
//	 * for the current block.
//	 * IN assertion: the fields freq and dad are set, heap[heap_max] and
//	 *    above are the tree nodes sorted by increasing frequency.
//	 * OUT assertions: the field len is set to the optimal bit length, the
//	 *     array bl_count contains the frequencies for each bit length.
//	 *     The length opt_len is updated; static_len is also updated if stree is
//	 *     not null.
//	 */
func _gen_bitlen(tls *libc.TLS, s uintptr, desc uintptr) {
	var base, bits, h, m, max_code, max_length, n, overflow, xbits, v5 int32
	var extra, stree, tree, p3 uintptr
	var f Tush
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, bits, extra, f, h, m, max_code, max_length, n, overflow, stree, tree, xbits, v5, p3
	tree = (*Ttree_desc)(unsafe.Pointer(desc)).Fdyn_tree
	max_code = (*Ttree_desc)(unsafe.Pointer(desc)).Fmax_code
	stree = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fstatic_tree
	extra = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fextra_bits
	base = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fextra_base
	max_length = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fmax_length /* frequency */
	overflow = 0                                                                                                  /* number of elements with bit length too large */
	bits = 0
	for {
		if !(bits <= int32(m_MAX_BITS)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 2976 + uintptr(bits)*2)) = uint16(0)
		goto _1
	_1:
		;
		bits++
	}
	/* In a first pass, compute the optimal bit lengths (which may
	 * overflow in the case of the bit length tree).
	 */
	*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 3008 + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max)*4)))*4 + 2)) = uint16(0) /* root of the heap */
	h = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max + int32(1)
	for {
		if !(h < libc.Int32FromInt32(2)*(libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES))+libc.Int32FromInt32(1)) {
			break
		}
		n = *(*int32)(unsafe.Pointer(s + 3008 + uintptr(h)*4))
		bits = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)))*4 + 2))) + int32(1)
		if bits > max_length {
			bits = max_length
			overflow++
		}
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = libc.Uint16FromInt32(bits)
		/* We overwrite tree[n].Dad which is no longer needed */
		if n > max_code {
			goto _2
		} /* not a leaf node */
		*(*Tush)(unsafe.Pointer(s + 2976 + uintptr(bits)*2))++
		xbits = 0
		if n >= base {
			xbits = *(*Tintf)(unsafe.Pointer(extra + uintptr(n-base)*4))
		}
		f = *(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))
		*(*Tulg)(unsafe.Pointer(s + 5912)) += uint64(f) * uint64(libc.Uint32FromInt32(bits+xbits))
		if stree != 0 {
			*(*Tulg)(unsafe.Pointer(s + 5920)) += uint64(f) * uint64(libc.Uint32FromInt32(libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(stree + uintptr(n)*4 + 2)))+xbits))
		}
		goto _2
	_2:
		;
		h++
	}
	if overflow == 0 {
		return
	}
	/* This happens for example on obj2 and pic of the Calgary corpus */
	/* Find the first bit length which could increase: */
	for cond := true; cond; cond = overflow > 0 {
		bits = max_length - int32(1)
		for libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2976 + uintptr(bits)*2))) == 0 {
			bits--
		}
		*(*Tush)(unsafe.Pointer(s + 2976 + uintptr(bits)*2))-- /* move one leaf down the tree */
		p3 = s + 2976 + uintptr(bits+int32(1))*2
		*(*Tush)(unsafe.Pointer(p3)) = Tush(int32(*(*Tush)(unsafe.Pointer(p3))) + libc.Int32FromInt32(2)) /* move one overflow item as its brother */
		*(*Tush)(unsafe.Pointer(s + 2976 + uintptr(max_length)*2))--
		/* The brother of the overflow item also moves one step up,
		 * but this does not affect bl_count[max_length]
		 */
		overflow -= int32(2)
	}
	/* Now recompute all bit lengths, scanning in increasing frequency.
	 * h is still equal to HEAP_SIZE. (It is simpler to reconstruct all
	 * lengths instead of fixing only the wrong ones. This idea is taken
	 * from 'ar' written by Haruhiko Okumura.)
	 */
	bits = max_length
	for {
		if !(bits != 0) {
			break
		}
		n = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2976 + uintptr(bits)*2)))
		for n != 0 {
			h--
			v5 = h
			m = *(*int32)(unsafe.Pointer(s + 3008 + uintptr(v5)*4))
			if m > max_code {
				continue
			}
			if uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2))) != libc.Uint32FromInt32(bits) {
				*(*Tulg)(unsafe.Pointer(s + 5912)) += (libc.Uint64FromInt32(bits) - uint64(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)))) * uint64(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4)))
				*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)) = libc.Uint16FromInt32(bits)
			}
			n--
		}
		goto _4
	_4:
		;
		bits--
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Construct one Huffman tree and assigns the code bit strings and lengths.
//	 * Update the total bit length for the current block.
//	 * IN assertion: the field freq is set for all tree elements.
//	 * OUT assertions: the fields len and code are set to the optimal bit length
//	 *     and corresponding code. The length opt_len is updated; static_len is
//	 *     also updated if stree is not null. The field max_code is set.
//	 */
func _build_tree(tls *libc.TLS, s uintptr, desc uintptr) {
	var elems, m, max_code, n, node, v11, v13, v15, v17, v19, v2, v20, v4, v5, v6, v7, v8 int32
	var stree, tree, v12, v14, v16, v21, v3, v9 uintptr
	var v18 Tush
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = elems, m, max_code, n, node, stree, tree, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v3, v4, v5, v6, v7, v8, v9
	tree = (*Ttree_desc)(unsafe.Pointer(desc)).Fdyn_tree
	stree = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fstatic_tree
	elems = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Felems /* iterate over heap elements */
	max_code = -int32(1)                                                                                /* new node being created */
	/* Construct the initial heap, with least frequent element in
	 * heap[SMALLEST]. The sons of heap[n] are heap[2*n] and heap[2*n + 1].
	 * heap[0] is not used.
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len = 0 /* new node being created */
	/* Construct the initial heap, with least frequent element in
	 * heap[SMALLEST]. The sons of heap[n] are heap[2*n] and heap[2*n + 1].
	 * heap[0] is not used.
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max = libc.Int32FromInt32(2)*(libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES)) + libc.Int32FromInt32(1)
	n = 0
	for {
		if !(n < elems) {
			break
		}
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))) != 0 {
			v3 = s + 5300
			*(*int32)(unsafe.Pointer(v3))++
			v2 = *(*int32)(unsafe.Pointer(v3))
			v4 = n
			max_code = v4
			*(*int32)(unsafe.Pointer(s + 3008 + uintptr(v2)*4)) = v4
			*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(n))) = uint8(0)
		} else {
			*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = uint16(0)
		}
		goto _1
	_1:
		;
		n++
	}
	/* The pkzip format requires that at least one distance code exists,
	 * and that at least one bit should be sent even if there is only one
	 * possible code. So to avoid special checks later on we force at least
	 * two codes of non zero frequency.
	 */
	for (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len < int32(2) {
		if max_code < int32(2) {
			max_code++
			v7 = max_code
			v6 = v7
		} else {
			v6 = 0
		}
		v5 = v6
		v9 = s + 5300
		*(*int32)(unsafe.Pointer(v9))++
		v8 = *(*int32)(unsafe.Pointer(v9))
		*(*int32)(unsafe.Pointer(s + 3008 + uintptr(v8)*4)) = v5
		node = v5
		*(*Tush)(unsafe.Pointer(tree + uintptr(node)*4)) = uint16(1)
		*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(node))) = uint8(0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len--
		if stree != 0 {
			*(*Tulg)(unsafe.Pointer(s + 5920)) -= uint64(*(*Tush)(unsafe.Pointer(stree + uintptr(node)*4 + 2)))
		}
		/* node is 0 or 1 so it does not have extra bits */
	}
	(*Ttree_desc)(unsafe.Pointer(desc)).Fmax_code = max_code
	/* The elements heap[heap_len/2 + 1 .. heap_len] are leaves of the tree,
	 * establish sub-heaps of increasing lengths:
	 */
	n = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len / int32(2)
	for {
		if !(n >= int32(1)) {
			break
		}
		_pqdownheap(tls, s, tree, n)
		goto _10
	_10:
		;
		n--
	}
	/* Construct the Huffman tree by repeatedly combining the least two
	 * frequent nodes.
	 */
	node = elems /* next internal node of the tree */
	for cond := true; cond; cond = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len >= int32(2) {
		n = *(*int32)(unsafe.Pointer(s + 3008 + 1*4))
		v12 = s + 5300
		v11 = *(*int32)(unsafe.Pointer(v12))
		*(*int32)(unsafe.Pointer(v12))--
		*(*int32)(unsafe.Pointer(s + 3008 + 1*4)) = *(*int32)(unsafe.Pointer(s + 3008 + uintptr(v11)*4))
		_pqdownheap(tls, s, tree, int32(m_SMALLEST))  /* n = node of least frequency */
		m = *(*int32)(unsafe.Pointer(s + 3008 + 1*4)) /* m = node of next least frequency */
		v14 = s + 5304
		*(*int32)(unsafe.Pointer(v14))--
		v13 = *(*int32)(unsafe.Pointer(v14))
		*(*int32)(unsafe.Pointer(s + 3008 + uintptr(v13)*4)) = n /* keep the nodes sorted by frequency */
		v16 = s + 5304
		*(*int32)(unsafe.Pointer(v16))--
		v15 = *(*int32)(unsafe.Pointer(v16))
		*(*int32)(unsafe.Pointer(s + 3008 + uintptr(v15)*4)) = m
		/* Create a new node father of n and m */
		*(*Tush)(unsafe.Pointer(tree + uintptr(node)*4)) = libc.Uint16FromInt32(libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))) + libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4))))
		if libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(n)))) >= libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(m)))) {
			v17 = libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(n))))
		} else {
			v17 = libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(m))))
		}
		*(*Tuch)(unsafe.Pointer(s + 5308 + uintptr(node))) = libc.Uint8FromInt32(v17 + libc.Int32FromInt32(1))
		v18 = libc.Uint16FromInt32(node)
		*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)) = v18
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = v18
		/* and insert the new node in the heap */
		v19 = node
		node++
		*(*int32)(unsafe.Pointer(s + 3008 + 1*4)) = v19
		_pqdownheap(tls, s, tree, int32(m_SMALLEST))
	}
	v21 = s + 5304
	*(*int32)(unsafe.Pointer(v21))--
	v20 = *(*int32)(unsafe.Pointer(v21))
	*(*int32)(unsafe.Pointer(s + 3008 + uintptr(v20)*4)) = *(*int32)(unsafe.Pointer(s + 3008 + 1*4))
	/* At this point, the fields freq and dad are set. We can now
	 * generate the bit lengths.
	 */
	_gen_bitlen(tls, s, desc)
	/* The field len is now set, we can generate the bit codes */
	_gen_codes(tls, tree, max_code, s+2976)
}

// C documentation
//
//	/* ===========================================================================
//	 * Scan a literal or distance tree to determine the frequencies of the codes
//	 * in the bit length tree.
//	 */
func _scan_tree(tls *libc.TLS, s uintptr, tree uintptr, max_code int32) {
	var count, curlen, max_count, min_count, n, nextlen, prevlen, v2 int32
	var p3 uintptr
	_, _, _, _, _, _, _, _, _ = count, curlen, max_count, min_count, n, nextlen, prevlen, v2, p3 /* iterates over all tree elements */
	prevlen = -int32(1)                                                                          /* length of current code */
	nextlen = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + 2)))                           /* length of next code */
	count = 0                                                                                    /* repeat count of the current code */
	max_count = int32(7)                                                                         /* max repeat count */
	min_count = int32(4)                                                                         /* min repeat count */
	if nextlen == 0 {
		max_count = int32(138)
		min_count = libc.Int32FromInt32(3)
	}
	*(*Tush)(unsafe.Pointer(tree + uintptr(max_code+int32(1))*4 + 2)) = libc.Uint16FromInt32(0xffff) /* guard */
	n = 0
	for {
		if !(n <= max_code) {
			break
		}
		curlen = nextlen
		nextlen = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(n+int32(1))*4 + 2)))
		count++
		v2 = count
		if v2 < max_count && curlen == nextlen {
			goto _1
		} else {
			if count < min_count {
				p3 = s + 2748 + uintptr(curlen)*4
				*(*Tush)(unsafe.Pointer(p3)) = Tush(int32(*(*Tush)(unsafe.Pointer(p3))) + count)
			} else {
				if curlen != 0 {
					if curlen != prevlen {
						*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4))++
					}
					*(*Tush)(unsafe.Pointer(s + 2748 + 16*4))++
				} else {
					if count <= int32(10) {
						*(*Tush)(unsafe.Pointer(s + 2748 + 17*4))++
					} else {
						*(*Tush)(unsafe.Pointer(s + 2748 + 18*4))++
					}
				}
			}
		}
		count = 0
		prevlen = curlen
		if nextlen == 0 {
			max_count = int32(138)
			min_count = libc.Int32FromInt32(3)
		} else {
			if curlen == nextlen {
				max_count = int32(6)
				min_count = libc.Int32FromInt32(3)
			} else {
				max_count = int32(7)
				min_count = libc.Int32FromInt32(4)
			}
		}
		goto _1
	_1:
		;
		n++
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Send a literal or distance tree in compressed form, using the codes in
//	 * bl_tree.
//	 */
func _send_tree(tls *libc.TLS, s uintptr, tree uintptr, max_code int32) {
	var count, curlen, len1, len11, len2, len3, len4, len5, len6, len7, max_count, min_count, n, nextlen, prevlen, val, val1, val2, val3, val4, val5, val6, val7, v2, v3 int32
	var v12, v14, v18, v20, v24, v26, v30, v32, v36, v38, v42, v44, v48, v50, v6, v8 Tulg
	var v13, v15, v19, v21, v25, v27, v31, v33, v37, v39, v43, v45, v49, v51, v7, v9, p10, p11, p16, p17, p22, p23, p28, p29, p34, p35, p40, p41, p46, p47, p5, p52 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, curlen, len1, len11, len2, len3, len4, len5, len6, len7, max_count, min_count, n, nextlen, prevlen, val, val1, val2, val3, val4, val5, val6, val7, v12, v13, v14, v15, v18, v19, v2, v20, v21, v24, v25, v26, v27, v3, v30, v31, v32, v33, v36, v37, v38, v39, v42, v43, v44, v45, v48, v49, v50, v51, v6, v7, v8, v9, p10, p11, p16, p17, p22, p23, p28, p29, p34, p35, p40, p41, p46, p47, p5, p52 /* iterates over all tree elements */
	prevlen = -int32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     /* length of current code */
	nextlen = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + 2)))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      /* length of next code */
	count = 0                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               /* repeat count of the current code */
	max_count = int32(7)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    /* max repeat count */
	min_count = int32(4)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    /* min repeat count */
	/* tree[max_code + 1].Len = -1; */ /* guard already set */
	if nextlen == 0 {
		max_count = int32(138)
		min_count = libc.Int32FromInt32(3)
	}
	n = 0
	for {
		if !(n <= max_code) {
			break
		}
		curlen = nextlen
		nextlen = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(n+int32(1))*4 + 2)))
		count++
		v2 = count
		if v2 < max_count && curlen == nextlen {
			goto _1
		} else {
			if count < min_count {
				for {
					len1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4 + 2)))
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
						val = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4)))
						p5 = s + 5936
						*(*Tush)(unsafe.Pointer(p5)) = Tush(int32(*(*Tush)(unsafe.Pointer(p5))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v7 = s + 40
						v6 = *(*Tulg)(unsafe.Pointer(v7))
						*(*Tulg)(unsafe.Pointer(v7))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v6))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v9 = s + 40
						v8 = *(*Tulg)(unsafe.Pointer(v9))
						*(*Tulg)(unsafe.Pointer(v9))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5940)) += len1 - int32(m_Buf_size)
					} else {
						p10 = s + 5936
						*(*Tush)(unsafe.Pointer(p10)) = Tush(int32(*(*Tush)(unsafe.Pointer(p10))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5940)) += len1
					}
					goto _4
				_4:
					;
					count--
					v3 = count
					if !(v3 != 0) {
						break
					}
				}
			} else {
				if curlen != 0 {
					if curlen != prevlen {
						len11 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
							val1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4)))
							p11 = s + 5936
							*(*Tush)(unsafe.Pointer(p11)) = Tush(int32(*(*Tush)(unsafe.Pointer(p11))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v13 = s + 40
							v12 = *(*Tulg)(unsafe.Pointer(v13))
							*(*Tulg)(unsafe.Pointer(v13))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v12))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v15 = s + 40
							v14 = *(*Tulg)(unsafe.Pointer(v15))
							*(*Tulg)(unsafe.Pointer(v15))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5940)) += len11 - int32(m_Buf_size)
						} else {
							p16 = s + 5936
							*(*Tush)(unsafe.Pointer(p16)) = Tush(int32(*(*Tush)(unsafe.Pointer(p16))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(curlen)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5940)) += len11
						}
						count--
					}
					len2 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 16*4 + 2)))
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
						val2 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 16*4)))
						p17 = s + 5936
						*(*Tush)(unsafe.Pointer(p17)) = Tush(int32(*(*Tush)(unsafe.Pointer(p17))) | libc.Int32FromUint16(libc.Uint16FromInt32(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v19 = s + 40
						v18 = *(*Tulg)(unsafe.Pointer(v19))
						*(*Tulg)(unsafe.Pointer(v19))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v18))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v21 = s + 40
						v20 = *(*Tulg)(unsafe.Pointer(v21))
						*(*Tulg)(unsafe.Pointer(v21))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v20))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5940)) += len2 - int32(m_Buf_size)
					} else {
						p22 = s + 5936
						*(*Tush)(unsafe.Pointer(p22)) = Tush(int32(*(*Tush)(unsafe.Pointer(p22))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 16*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5940)) += len2
					}
					len3 = int32(2)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
						val3 = count - int32(3)
						p23 = s + 5936
						*(*Tush)(unsafe.Pointer(p23)) = Tush(int32(*(*Tush)(unsafe.Pointer(p23))) | libc.Int32FromUint16(libc.Uint16FromInt32(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v25 = s + 40
						v24 = *(*Tulg)(unsafe.Pointer(v25))
						*(*Tulg)(unsafe.Pointer(v25))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v27 = s + 40
						v26 = *(*Tulg)(unsafe.Pointer(v27))
						*(*Tulg)(unsafe.Pointer(v27))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5940)) += len3 - int32(m_Buf_size)
					} else {
						p28 = s + 5936
						*(*Tush)(unsafe.Pointer(p28)) = Tush(int32(*(*Tush)(unsafe.Pointer(p28))) | libc.Int32FromUint16(libc.Uint16FromInt32(count-libc.Int32FromInt32(3)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5940)) += len3
					}
				} else {
					if count <= int32(10) {
						len4 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 17*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len4 {
							val4 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 17*4)))
							p29 = s + 5936
							*(*Tush)(unsafe.Pointer(p29)) = Tush(int32(*(*Tush)(unsafe.Pointer(p29))) | libc.Int32FromUint16(libc.Uint16FromInt32(val4))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v31 = s + 40
							v30 = *(*Tulg)(unsafe.Pointer(v31))
							*(*Tulg)(unsafe.Pointer(v31))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v30))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v33 = s + 40
							v32 = *(*Tulg)(unsafe.Pointer(v33))
							*(*Tulg)(unsafe.Pointer(v33))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v32))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val4)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5940)) += len4 - int32(m_Buf_size)
						} else {
							p34 = s + 5936
							*(*Tush)(unsafe.Pointer(p34)) = Tush(int32(*(*Tush)(unsafe.Pointer(p34))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 17*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5940)) += len4
						}
						len5 = int32(3)
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len5 {
							val5 = count - int32(3)
							p35 = s + 5936
							*(*Tush)(unsafe.Pointer(p35)) = Tush(int32(*(*Tush)(unsafe.Pointer(p35))) | libc.Int32FromUint16(libc.Uint16FromInt32(val5))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v37 = s + 40
							v36 = *(*Tulg)(unsafe.Pointer(v37))
							*(*Tulg)(unsafe.Pointer(v37))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v36))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v39 = s + 40
							v38 = *(*Tulg)(unsafe.Pointer(v39))
							*(*Tulg)(unsafe.Pointer(v39))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v38))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val5)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5940)) += len5 - int32(m_Buf_size)
						} else {
							p40 = s + 5936
							*(*Tush)(unsafe.Pointer(p40)) = Tush(int32(*(*Tush)(unsafe.Pointer(p40))) | libc.Int32FromUint16(libc.Uint16FromInt32(count-libc.Int32FromInt32(3)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5940)) += len5
						}
					} else {
						len6 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 18*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len6 {
							val6 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 18*4)))
							p41 = s + 5936
							*(*Tush)(unsafe.Pointer(p41)) = Tush(int32(*(*Tush)(unsafe.Pointer(p41))) | libc.Int32FromUint16(libc.Uint16FromInt32(val6))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v43 = s + 40
							v42 = *(*Tulg)(unsafe.Pointer(v43))
							*(*Tulg)(unsafe.Pointer(v43))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v42))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v45 = s + 40
							v44 = *(*Tulg)(unsafe.Pointer(v45))
							*(*Tulg)(unsafe.Pointer(v45))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v44))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val6)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5940)) += len6 - int32(m_Buf_size)
						} else {
							p46 = s + 5936
							*(*Tush)(unsafe.Pointer(p46)) = Tush(int32(*(*Tush)(unsafe.Pointer(p46))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + 18*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5940)) += len6
						}
						len7 = int32(7)
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len7 {
							val7 = count - int32(11)
							p47 = s + 5936
							*(*Tush)(unsafe.Pointer(p47)) = Tush(int32(*(*Tush)(unsafe.Pointer(p47))) | libc.Int32FromUint16(libc.Uint16FromInt32(val7))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v49 = s + 40
							v48 = *(*Tulg)(unsafe.Pointer(v49))
							*(*Tulg)(unsafe.Pointer(v49))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v48))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v51 = s + 40
							v50 = *(*Tulg)(unsafe.Pointer(v51))
							*(*Tulg)(unsafe.Pointer(v51))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v50))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val7)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5940)) += len7 - int32(m_Buf_size)
						} else {
							p52 = s + 5936
							*(*Tush)(unsafe.Pointer(p52)) = Tush(int32(*(*Tush)(unsafe.Pointer(p52))) | libc.Int32FromUint16(libc.Uint16FromInt32(count-libc.Int32FromInt32(11)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5940)) += len7
						}
					}
				}
			}
		}
		count = 0
		prevlen = curlen
		if nextlen == 0 {
			max_count = int32(138)
			min_count = libc.Int32FromInt32(3)
		} else {
			if curlen == nextlen {
				max_count = int32(6)
				min_count = libc.Int32FromInt32(3)
			} else {
				max_count = int32(7)
				min_count = libc.Int32FromInt32(4)
			}
		}
		goto _1
	_1:
		;
		n++
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Construct the Huffman tree for the bit lengths and return the index in
//	 * bl_order of the last bit length code to send.
//	 */
func _build_bl_tree(tls *libc.TLS, s uintptr) (r int32) {
	var max_blindex int32
	_ = max_blindex /* index of last bit length code of non zero freq */
	/* Determine the bit length frequencies for literal and distance trees */
	_scan_tree(tls, s, s+212, (*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fmax_code)
	_scan_tree(tls, s, s+2504, (*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fmax_code)
	/* Build the bit length tree: */
	_build_tree(tls, s, s+2952)
	/* opt_len now includes the length of the tree representations, except the
	 * lengths of the bit lengths codes and the 5 + 5 + 4 bits for the counts.
	 */
	/* Determine the number of bit length codes to send. The pkzip format
	 * requires that at least 4 bit length codes be sent. (appnote.txt says
	 * 3 but the actual value used is 4.)
	 */
	max_blindex = libc.Int32FromInt32(m_BL_CODES) - libc.Int32FromInt32(1)
	for {
		if !(max_blindex >= int32(3)) {
			break
		}
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(_bl_order[max_blindex])*4 + 2))) != 0 {
			break
		}
		goto _1
	_1:
		;
		max_blindex--
	}
	/* Update opt_len to include the bit length tree and counts */
	*(*Tulg)(unsafe.Pointer(s + 5912)) += uint64(3)*(libc.Uint64FromInt32(max_blindex)+uint64(1)) + uint64(5) + uint64(5) + uint64(4)
	return max_blindex
}

// C documentation
//
//	/* ===========================================================================
//	 * Send the header for a block using dynamic Huffman trees: the counts, the
//	 * lengths of the bit length codes, the literal tree and the distance tree.
//	 * IN assertion: lcodes >= 257, dcodes >= 1, blcodes >= 4.
//	 */
func _send_all_trees(tls *libc.TLS, s uintptr, lcodes int32, dcodes int32, blcodes int32) {
	var len1, len11, len2, len3, rank, val, val1, val2, val3 int32
	var v10, v14, v16, v2, v21, v23, v4, v8 Tulg
	var v11, v15, v17, v22, v24, v3, v5, v9, p1, p12, p13, p18, p20, p25, p6, p7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, len11, len2, len3, rank, val, val1, val2, val3, v10, v11, v14, v15, v16, v17, v2, v21, v22, v23, v24, v3, v4, v5, v8, v9, p1, p12, p13, p18, p20, p25, p6, p7 /* index in bl_order */
	len1 = int32(5)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
		val = lcodes - int32(257)
		p1 = s + 5936
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 40
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 40
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5936
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | libc.Int32FromUint16(libc.Uint16FromInt32(lcodes-libc.Int32FromInt32(257)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len1
	} /* not +255 as stated in appnote.txt */
	len11 = int32(5)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
		val1 = dcodes - int32(1)
		p7 = s + 5936
		*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v9 = s + 40
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v11 = s + 40
		v10 = *(*Tulg)(unsafe.Pointer(v11))
		*(*Tulg)(unsafe.Pointer(v11))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len11 - int32(m_Buf_size)
	} else {
		p12 = s + 5936
		*(*Tush)(unsafe.Pointer(p12)) = Tush(int32(*(*Tush)(unsafe.Pointer(p12))) | libc.Int32FromUint16(libc.Uint16FromInt32(dcodes-libc.Int32FromInt32(1)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len11
	}
	len2 = int32(4)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
		val2 = blcodes - int32(4)
		p13 = s + 5936
		*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | libc.Int32FromUint16(libc.Uint16FromInt32(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v15 = s + 40
		v14 = *(*Tulg)(unsafe.Pointer(v15))
		*(*Tulg)(unsafe.Pointer(v15))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v17 = s + 40
		v16 = *(*Tulg)(unsafe.Pointer(v17))
		*(*Tulg)(unsafe.Pointer(v17))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v16))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len2 - int32(m_Buf_size)
	} else {
		p18 = s + 5936
		*(*Tush)(unsafe.Pointer(p18)) = Tush(int32(*(*Tush)(unsafe.Pointer(p18))) | libc.Int32FromUint16(libc.Uint16FromInt32(blcodes-libc.Int32FromInt32(4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len2
	} /* not -3 as stated in appnote.txt */
	rank = 0
	for {
		if !(rank < blcodes) {
			break
		}
		len3 = int32(3)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
			val3 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(_bl_order[rank])*4 + 2)))
			p20 = s + 5936
			*(*Tush)(unsafe.Pointer(p20)) = Tush(int32(*(*Tush)(unsafe.Pointer(p20))) | libc.Int32FromUint16(libc.Uint16FromInt32(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
			v22 = s + 40
			v21 = *(*Tulg)(unsafe.Pointer(v22))
			*(*Tulg)(unsafe.Pointer(v22))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v21))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
			v24 = s + 40
			v23 = *(*Tulg)(unsafe.Pointer(v24))
			*(*Tulg)(unsafe.Pointer(v24))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v23))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
			*(*int32)(unsafe.Pointer(s + 5940)) += len3 - int32(m_Buf_size)
		} else {
			p25 = s + 5936
			*(*Tush)(unsafe.Pointer(p25)) = Tush(int32(*(*Tush)(unsafe.Pointer(p25))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2748 + uintptr(_bl_order[rank])*4 + 2)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
			*(*int32)(unsafe.Pointer(s + 5940)) += len3
		}
		goto _19
	_19:
		;
		rank++
	}
	_send_tree(tls, s, s+212, lcodes-int32(1))  /* literal tree */
	_send_tree(tls, s, s+2504, dcodes-int32(1)) /* distance tree */
}

// C documentation
//
//	/* ===========================================================================
//	 * Send a stored block
//	 */
func x__tr_stored_block(tls *libc.TLS, s uintptr, buf uintptr, stored_len Tulg, last int32) {
	var len1, val int32
	var v10, v12, v14, v3, v5, v8, p1, p6 uintptr
	var v11, v13, v2, v4, v7, v9 Tulg
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, val, v10, v11, v12, v13, v14, v2, v3, v4, v5, v7, v8, v9, p1, p6
	len1 = int32(3)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
		val = libc.Int32FromInt32(m_STORED_BLOCK)<<libc.Int32FromInt32(1) + last
		p1 = s + 5936
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 40
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 40
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5936
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_STORED_BLOCK)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len1
	} /* send block type */
	_bi_windup(tls, s) /* align on byte boundary */
	v8 = s + 40
	v7 = *(*Tulg)(unsafe.Pointer(v8))
	*(*Tulg)(unsafe.Pointer(v8))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v7))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(stored_len)) & libc.Int32FromInt32(0xff))
	v10 = s + 40
	v9 = *(*Tulg)(unsafe.Pointer(v10))
	*(*Tulg)(unsafe.Pointer(v10))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(stored_len)) >> libc.Int32FromInt32(8))
	v12 = s + 40
	v11 = *(*Tulg)(unsafe.Pointer(v12))
	*(*Tulg)(unsafe.Pointer(v12))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(^stored_len)) & libc.Int32FromInt32(0xff))
	v14 = s + 40
	v13 = *(*Tulg)(unsafe.Pointer(v14))
	*(*Tulg)(unsafe.Pointer(v14))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v13))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(^stored_len)) >> libc.Int32FromInt32(8))
	if stored_len != 0 {
		libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), buf, stored_len)
	}
	*(*Tulg)(unsafe.Pointer(s + 40)) += stored_len
}

// C documentation
//
//	/* ===========================================================================
//	 * Flush the bits in the bit buffer to pending output (leaves at most 7 bits)
//	 */
func x__tr_flush_bits(tls *libc.TLS, s uintptr) {
	_bi_flush(tls, s)
}

// C documentation
//
//	/* ===========================================================================
//	 * Send one empty static block to give enough lookahead for inflate.
//	 * This takes 10 bits, of which 7 may remain in the bit buffer.
//	 */
func x__tr_align(tls *libc.TLS, s uintptr) {
	var len1, len11, val, val1 int32
	var v10, v2, v4, v8 Tulg
	var v11, v3, v5, v9, p1, p12, p6, p7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, len11, val, val1, v10, v11, v2, v3, v4, v5, v8, v9, p1, p12, p6, p7
	len1 = int32(3)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
		val = libc.Int32FromInt32(m_STATIC_TREES) << libc.Int32FromInt32(1)
		p1 = s + 5936
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 40
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 40
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5936
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len1
	}
	len11 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4 + 2)))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
		val1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4)))
		p7 = s + 5936
		*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v9 = s + 40
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v11 = s + 40
		v10 = *(*Tulg)(unsafe.Pointer(v11))
		*(*Tulg)(unsafe.Pointer(v11))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len11 - int32(m_Buf_size)
	} else {
		p12 = s + 5936
		*(*Tush)(unsafe.Pointer(p12)) = Tush(int32(*(*Tush)(unsafe.Pointer(p12))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len11
	}
	_bi_flush(tls, s)
}

// C documentation
//
//	/* ===========================================================================
//	 * Send the block data compressed using the given Huffman trees
//	 */
func _compress_block(tls *libc.TLS, s uintptr, ltree uintptr, dtree uintptr) {
	var code, dist, sx, v1, v2, v3 uint32
	var extra, lc, len1, len11, len2, len3, len4, len5, val, val1, val2, val3, val4, val5, v22 int32
	var v11, v13, v17, v19, v24, v26, v30, v32, v36, v38, v5, v7 Tulg
	var v12, v14, v18, v20, v25, v27, v31, v33, v37, v39, v6, v8, p10, p15, p16, p21, p23, p28, p29, p34, p35, p4, p40, p9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = code, dist, extra, lc, len1, len11, len2, len3, len4, len5, sx, val, val1, val2, val3, val4, val5, v1, v11, v12, v13, v14, v17, v18, v19, v2, v20, v22, v24, v25, v26, v27, v3, v30, v31, v32, v33, v36, v37, v38, v39, v5, v6, v7, v8, p10, p15, p16, p21, p23, p28, p29, p34, p35, p4, p40, p9 /* match length or unmatched char (if dist == 0) */
	sx = uint32(0)                                                                                                                                                                                                                                                                                                                                                                                                                                                               /* number of extra bits to send */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != uint32(0) {
		for cond := true; cond; cond = sx < (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next {
			v1 = sx
			sx++
			dist = libc.Uint32FromInt32(libc.Int32FromUint8(*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v1)))) & int32(0xff))
			v2 = sx
			sx++
			dist += libc.Uint32FromInt32(libc.Int32FromUint8(*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v2))))&libc.Int32FromInt32(0xff)) << int32(8)
			v3 = sx
			sx++
			lc = libc.Int32FromUint8(*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))))
			if dist == uint32(0) {
				len1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
					val = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4)))
					p4 = s + 5936
					*(*Tush)(unsafe.Pointer(p4)) = Tush(int32(*(*Tush)(unsafe.Pointer(p4))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v6 = s + 40
					v5 = *(*Tulg)(unsafe.Pointer(v6))
					*(*Tulg)(unsafe.Pointer(v6))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v8 = s + 40
					v7 = *(*Tulg)(unsafe.Pointer(v8))
					*(*Tulg)(unsafe.Pointer(v8))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v7))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5940)) += len1 - int32(m_Buf_size)
				} else {
					p9 = s + 5936
					*(*Tush)(unsafe.Pointer(p9)) = Tush(int32(*(*Tush)(unsafe.Pointer(p9))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5940)) += len1
				} /* send a literal byte */
			} else {
				/* Here, lc is the match length - MIN_MATCH */
				code = uint32(x__length_code[lc])
				len11 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
					val1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4)))
					p10 = s + 5936
					*(*Tush)(unsafe.Pointer(p10)) = Tush(int32(*(*Tush)(unsafe.Pointer(p10))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v12 = s + 40
					v11 = *(*Tulg)(unsafe.Pointer(v12))
					*(*Tulg)(unsafe.Pointer(v12))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v14 = s + 40
					v13 = *(*Tulg)(unsafe.Pointer(v14))
					*(*Tulg)(unsafe.Pointer(v14))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v13))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5940)) += len11 - int32(m_Buf_size)
				} else {
					p15 = s + 5936
					*(*Tush)(unsafe.Pointer(p15)) = Tush(int32(*(*Tush)(unsafe.Pointer(p15))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5940)) += len11
				} /* send length code */
				extra = _extra_lbits[code]
				if extra != 0 {
					lc -= _base_length[code]
					len2 = extra
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
						val2 = lc
						p16 = s + 5936
						*(*Tush)(unsafe.Pointer(p16)) = Tush(int32(*(*Tush)(unsafe.Pointer(p16))) | libc.Int32FromUint16(libc.Uint16FromInt32(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v18 = s + 40
						v17 = *(*Tulg)(unsafe.Pointer(v18))
						*(*Tulg)(unsafe.Pointer(v18))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v17))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v20 = s + 40
						v19 = *(*Tulg)(unsafe.Pointer(v20))
						*(*Tulg)(unsafe.Pointer(v20))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v19))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5940)) += len2 - int32(m_Buf_size)
					} else {
						p21 = s + 5936
						*(*Tush)(unsafe.Pointer(p21)) = Tush(int32(*(*Tush)(unsafe.Pointer(p21))) | libc.Int32FromUint16(libc.Uint16FromInt32(lc))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5940)) += len2
					} /* send the extra length bits */
				}
				dist-- /* dist is now the match distance - 1 */
				if dist < uint32(256) {
					v22 = libc.Int32FromUint8(x__dist_code[dist])
				} else {
					v22 = libc.Int32FromUint8(x__dist_code[uint32(256)+dist>>int32(7)])
				}
				code = libc.Uint32FromInt32(v22)
				len3 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
					val3 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4)))
					p23 = s + 5936
					*(*Tush)(unsafe.Pointer(p23)) = Tush(int32(*(*Tush)(unsafe.Pointer(p23))) | libc.Int32FromUint16(libc.Uint16FromInt32(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v25 = s + 40
					v24 = *(*Tulg)(unsafe.Pointer(v25))
					*(*Tulg)(unsafe.Pointer(v25))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v27 = s + 40
					v26 = *(*Tulg)(unsafe.Pointer(v27))
					*(*Tulg)(unsafe.Pointer(v27))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5940)) += len3 - int32(m_Buf_size)
				} else {
					p28 = s + 5936
					*(*Tush)(unsafe.Pointer(p28)) = Tush(int32(*(*Tush)(unsafe.Pointer(p28))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5940)) += len3
				} /* send the distance code */
				extra = _extra_dbits[code]
				if extra != 0 {
					dist -= libc.Uint32FromInt32(_base_dist[code])
					len4 = extra
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len4 {
						val4 = libc.Int32FromUint32(dist)
						p29 = s + 5936
						*(*Tush)(unsafe.Pointer(p29)) = Tush(int32(*(*Tush)(unsafe.Pointer(p29))) | libc.Int32FromUint16(libc.Uint16FromInt32(val4))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v31 = s + 40
						v30 = *(*Tulg)(unsafe.Pointer(v31))
						*(*Tulg)(unsafe.Pointer(v31))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v30))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v33 = s + 40
						v32 = *(*Tulg)(unsafe.Pointer(v33))
						*(*Tulg)(unsafe.Pointer(v33))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v32))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val4)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5940)) += len4 - int32(m_Buf_size)
					} else {
						p34 = s + 5936
						*(*Tush)(unsafe.Pointer(p34)) = Tush(int32(*(*Tush)(unsafe.Pointer(p34))) | libc.Int32FromUint16(uint16(dist))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5940)) += len4
					} /* send the extra distance bits */
				}
			} /* literal or match pair ? */
			/* Check for no overlay of pending_buf on needed symbols */
		}
	}
	len5 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + 256*4 + 2)))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len5 {
		val5 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + 256*4)))
		p35 = s + 5936
		*(*Tush)(unsafe.Pointer(p35)) = Tush(int32(*(*Tush)(unsafe.Pointer(p35))) | libc.Int32FromUint16(libc.Uint16FromInt32(val5))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v37 = s + 40
		v36 = *(*Tulg)(unsafe.Pointer(v37))
		*(*Tulg)(unsafe.Pointer(v37))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v36))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v39 = s + 40
		v38 = *(*Tulg)(unsafe.Pointer(v39))
		*(*Tulg)(unsafe.Pointer(v39))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v38))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val5)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5940)) += len5 - int32(m_Buf_size)
	} else {
		p40 = s + 5936
		*(*Tush)(unsafe.Pointer(p40)) = Tush(int32(*(*Tush)(unsafe.Pointer(p40))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + 256*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5940)) += len5
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Check if the data type is TEXT or BINARY, using the following algorithm:
//	 * - TEXT if the two conditions below are satisfied:
//	 *    a) There are no non-portable control characters belonging to the
//	 *       "block list" (0..6, 14..25, 28..31).
//	 *    b) There is at least one printable character belonging to the
//	 *       "allow list" (9 {TAB}, 10 {LF}, 13 {CR}, 32..255).
//	 * - BINARY otherwise.
//	 * - The following partially-portable control characters form a
//	 *   "gray list" that is ignored in this detection algorithm:
//	 *   (7 {BEL}, 8 {BS}, 11 {VT}, 12 {FF}, 26 {SUB}, 27 {ESC}).
//	 * IN assertion: the fields Freq of dyn_ltree are set.
//	 */
func _detect_data_type(tls *libc.TLS, s uintptr) (r int32) {
	var block_mask uint64
	var n int32
	_, _ = block_mask, n
	/* block_mask is the bit mask of block-listed bytes
	 * set bits 0..6, 14..25, and 28..31
	 * 0xf3ffc07f = binary 11110011111111111100000001111111
	 */
	block_mask = uint64(0xf3ffc07f)
	/* Check for non-textual ("block-listed") bytes. */
	n = 0
	for {
		if !(n <= int32(31)) {
			break
		}
		if block_mask&uint64(1) != 0 && libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 212 + uintptr(n)*4))) != 0 {
			return m_Z_BINARY
		}
		goto _1
	_1:
		;
		n++
		block_mask >>= uint64(1)
	}
	/* Check for textual ("allow-listed") bytes. */
	if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 212 + 9*4))) != 0 || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 212 + 10*4))) != 0 || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 212 + 13*4))) != 0 {
		return int32(m_Z_TEXT)
	}
	n = int32(32)
	for {
		if !(n < int32(m_LITERALS)) {
			break
		}
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 212 + uintptr(n)*4))) != 0 {
			return int32(m_Z_TEXT)
		}
		goto _2
	_2:
		;
		n++
	}
	/* There are no "block-listed" or "allow-listed" bytes:
	 * this stream either is empty or has tolerated ("gray-listed") bytes only.
	 */
	return m_Z_BINARY
}

// C documentation
//
//	/* ===========================================================================
//	 * Determine the best encoding for the current block: dynamic trees, static
//	 * trees or store, and write out the encoded block.
//	 */
func x__tr_flush_block(tls *libc.TLS, s uintptr, buf uintptr, stored_len Tulg, last int32) {
	var len1, len11, max_blindex, val, val1 int32
	var opt_lenb, static_lenb, v1, v11, v3, v5, v9 Tulg
	var v10, v12, v4, v6, p13, p2, p7, p8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, len11, max_blindex, opt_lenb, static_lenb, val, val1, v1, v10, v11, v12, v3, v4, v5, v6, v9, p13, p2, p7, p8 /* opt_len and static_len in bytes */
	max_blindex = 0                                                                                                                                                                 /* index of last bit length code of non zero freq */
	/* Build the Huffman trees unless a stored block is forced */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel > 0 {
		/* Check if the file is binary or text */
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fdata_type == int32(m_Z_UNKNOWN) {
			(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fdata_type = _detect_data_type(tls, s)
		}
		/* Construct the literal and distance trees */
		_build_tree(tls, s, s+2904)
		_build_tree(tls, s, s+2928)
		/* At this point, opt_len and static_len are the total bit lengths of
		 * the compressed block data, excluding the tree representations.
		 */
		/* Build the bit length tree for the above two trees, and get the index
		 * in bl_order of the last bit length code to send.
		 */
		max_blindex = _build_bl_tree(tls, s)
		/* Determine the best encoding. Compute the block lengths in bytes. */
		opt_lenb = ((*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len + uint64(3) + uint64(7)) >> int32(3)
		static_lenb = ((*Tdeflate_state)(unsafe.Pointer(s)).Fstatic_len + uint64(3) + uint64(7)) >> int32(3)
		if static_lenb <= opt_lenb || (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_FIXED) {
			opt_lenb = static_lenb
		}
	} else {
		v1 = stored_len + libc.Uint64FromInt32(5)
		static_lenb = v1
		opt_lenb = v1 /* force a stored block */
	}
	if stored_len+uint64(4) <= opt_lenb && buf != libc.UintptrFromInt32(0) {
		/* 4: two words for the lengths */
		/* The test buf != NULL is only necessary if LIT_BUFSIZE > WSIZE.
		 * Otherwise we can't have processed more than WSIZE input bytes since
		 * the last block flush, because compression would have been
		 * successful. If LIT_BUFSIZE <= WSIZE, it is never too late to
		 * transform a block into a stored block.
		 */
		x__tr_stored_block(tls, s, buf, stored_len, last)
	} else {
		if static_lenb == opt_lenb {
			len1 = int32(3)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
				val = libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1) + last
				p2 = s + 5936
				*(*Tush)(unsafe.Pointer(p2)) = Tush(int32(*(*Tush)(unsafe.Pointer(p2))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				v4 = s + 40
				v3 = *(*Tulg)(unsafe.Pointer(v4))
				*(*Tulg)(unsafe.Pointer(v4))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
				v6 = s + 40
				v5 = *(*Tulg)(unsafe.Pointer(v6))
				*(*Tulg)(unsafe.Pointer(v6))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
				*(*int32)(unsafe.Pointer(s + 5940)) += len1 - int32(m_Buf_size)
			} else {
				p7 = s + 5936
				*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				*(*int32)(unsafe.Pointer(s + 5940)) += len1
			}
			_compress_block(tls, s, uintptr(unsafe.Pointer(&_static_ltree)), uintptr(unsafe.Pointer(&_static_dtree)))
		} else {
			len11 = int32(3)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
				val1 = libc.Int32FromInt32(m_DYN_TREES)<<libc.Int32FromInt32(1) + last
				p8 = s + 5936
				*(*Tush)(unsafe.Pointer(p8)) = Tush(int32(*(*Tush)(unsafe.Pointer(p8))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				v10 = s + 40
				v9 = *(*Tulg)(unsafe.Pointer(v10))
				*(*Tulg)(unsafe.Pointer(v10))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
				v12 = s + 40
				v11 = *(*Tulg)(unsafe.Pointer(v12))
				*(*Tulg)(unsafe.Pointer(v12))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
				*(*int32)(unsafe.Pointer(s + 5940)) += len11 - int32(m_Buf_size)
			} else {
				p13 = s + 5936
				*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_DYN_TREES)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				*(*int32)(unsafe.Pointer(s + 5940)) += len11
			}
			_send_all_trees(tls, s, (*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fmax_code+int32(1), (*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fmax_code+int32(1), max_blindex+int32(1))
			_compress_block(tls, s, s+212, s+2504)
		}
	}
	/* The above check is made mod 2^32, for files larger than 512 MB
	 * and uLong implemented on 32 bits.
	 */
	_init_block(tls, s)
	if last != 0 {
		_bi_windup(tls, s)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Save the match info and tally the frequency counts. Return true if
//	 * the current block must be flushed.
//	 */
func x__tr_tally(tls *libc.TLS, s uintptr, dist uint32, lc uint32) (r int32) {
	var v1, v3, v5 TuInt
	var v2, v4, v6 uintptr
	var v7 int32
	_, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, v7
	v2 = s + 5900
	v1 = *(*TuInt)(unsafe.Pointer(v2))
	*(*TuInt)(unsafe.Pointer(v2))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v1))) = uint8(dist)
	v4 = s + 5900
	v3 = *(*TuInt)(unsafe.Pointer(v4))
	*(*TuInt)(unsafe.Pointer(v4))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))) = uint8(dist >> libc.Int32FromInt32(8))
	v6 = s + 5900
	v5 = *(*TuInt)(unsafe.Pointer(v6))
	*(*TuInt)(unsafe.Pointer(v6))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v5))) = uint8(lc)
	if dist == uint32(0) {
		/* lc is the unmatched char */
		*(*Tush)(unsafe.Pointer(s + 212 + uintptr(lc)*4))++
	} else {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
		/* Here, lc is the match length - MIN_MATCH */
		dist-- /* dist = match distance - 1 */
		*(*Tush)(unsafe.Pointer(s + 212 + uintptr(libc.Int32FromUint8(x__length_code[lc])+int32(m_LITERALS)+int32(1))*4))++
		if dist < uint32(256) {
			v7 = libc.Int32FromUint8(x__dist_code[dist])
		} else {
			v7 = libc.Int32FromUint8(x__dist_code[uint32(256)+dist>>int32(7)])
		}
		*(*Tush)(unsafe.Pointer(s + 2504 + uintptr(v7)*4))++
	}
	return libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
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
const m_BUFSIZ = 1024
const m_COPY = 1
const m_DN_ACCESS = 0x00000001
const m_DN_ATTRIB = 0x00000020
const m_DN_CREATE = 0x00000004
const m_DN_DELETE = 0x00000008
const m_DN_MODIFY = 0x00000002
const m_DN_MULTISHOT = 0x80000000
const m_DN_RENAME = 0x00000010
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
const m_FALLOC_FL_KEEP_SIZE = 1
const m_FALLOC_FL_PUNCH_HOLE = 2
const m_FAPPEND = "O_APPEND"
const m_FASYNC = "O_ASYNC"
const m_FD_CLOEXEC = 1
const m_FFSYNC = "O_SYNC"
const m_FILENAME_MAX = 4096
const m_FNDELAY = "O_NDELAY"
const m_FNONBLOCK = "O_NONBLOCK"
const m_FOPEN_MAX = 1000
const m_F_ADD_SEALS = 1033
const m_F_CANCELLK = 1029
const m_F_DUPFD = 0
const m_F_DUPFD_CLOEXEC = 1030
const m_F_GETFD = 1
const m_F_GETFL = 3
const m_F_GETLEASE = 1025
const m_F_GETLK = 5
const m_F_GETLK64 = "F_GETLK"
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
const m_F_SETLK64 = "F_SETLK"
const m_F_SETLKW = 7
const m_F_SETLKW64 = "F_SETLKW"
const m_F_SETOWN = 8
const m_F_SETOWN_EX = 15
const m_F_SETPIPE_SZ = 1031
const m_F_SETSIG = 10
const m_F_SET_FILE_RW_HINT = 1038
const m_F_SET_RW_HINT = 1036
const m_F_UNLCK = 2
const m_F_WRLCK = 1
const m_GZBUFSIZE = 8192
const m_GZIP = 2
const m_GZ_APPEND = 1
const m_GZ_NONE = 0
const m_GZ_READ = 7247
const m_GZ_WRITE = 31153
const m_LOOK = 0
const m_L_ctermid = 20
const m_L_cuserid = 20
const m_L_tmpnam = 20
const m_MAX_HANDLE_SZ = 128
const m_O_APPEND = 02000
const m_O_ASYNC = 020000
const m_O_CLOEXEC = 02000000
const m_O_CREAT = 0100
const m_O_DIRECT = 0200000
const m_O_DIRECTORY = 040000
const m_O_DSYNC = 010000
const m_O_EXCL = 0200
const m_O_EXEC = "O_PATH"
const m_O_LARGEFILE = 0400000
const m_O_NDELAY = "O_NONBLOCK"
const m_O_NOATIME = 01000000
const m_O_NOCTTY = 0400
const m_O_NOFOLLOW = 0100000
const m_O_NONBLOCK = 04000
const m_O_PATH = 010000000
const m_O_RDONLY = 00
const m_O_RDWR = 02
const m_O_RSYNC = 04010000
const m_O_SEARCH = "O_PATH"
const m_O_SYNC = 04010000
const m_O_TMPFILE = 020040000
const m_O_TRUNC = 01000
const m_O_TTY_INIT = 0
const m_O_WRONLY = 01
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
const m_P_tmpdir = "/tmp"
const m_RWF_WRITE_LIFE_NOT_SET = 0
const m_RWH_WRITE_LIFE_EXTREME = 5
const m_RWH_WRITE_LIFE_LONG = 4
const m_RWH_WRITE_LIFE_MEDIUM = 3
const m_RWH_WRITE_LIFE_NONE = 1
const m_RWH_WRITE_LIFE_SHORT = 2
const m_SPLICE_F_GIFT = 8
const m_SPLICE_F_MORE = 4
const m_SPLICE_F_MOVE = 1
const m_SPLICE_F_NONBLOCK = 2
const m_SYNC_FILE_RANGE_WAIT_AFTER = 4
const m_SYNC_FILE_RANGE_WAIT_BEFORE = 1
const m_SYNC_FILE_RANGE_WRITE = 2
const m_S_IRGRP = 0040
const m_S_IROTH = 0004
const m_S_IRUSR = 0400
const m_S_IRWXG = 0070
const m_S_IRWXO = 0007
const m_S_IRWXU = 0700
const m_S_ISGID = 02000
const m_S_ISUID = 04000
const m_S_ISVTX = 01000
const m_S_IWGRP = 0020
const m_S_IWOTH = 0002
const m_S_IWUSR = 0200
const m_S_IXGRP = 0010
const m_S_IXOTH = 0001
const m_S_IXUSR = 0100
const m_TMP_MAX = 10000
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
const m__LARGEFILE_SOURCE = 1
const m_creat64 = "creat"
const m_fallocate64 = "fallocate"
const m_fgetpos64 = "fgetpos"
const m_flock64 = "flock"
const m_fopen64 = "fopen"
const m_fpos64_t = "fpos_t"
const m_freopen64 = "freopen"
const m_fseeko64 = "fseeko"
const m_fsetpos64 = "fsetpos"
const m_ftello64 = "ftello"
const m_loff_t = "off_t"
const m_open64 = "open"
const m_openat64 = "openat"
const m_posix_fadvise64 = "posix_fadvise"
const m_posix_fallocate64 = "posix_fallocate"
const m_tmpfile64 = "tmpfile"

type t__isoc_va_list = uintptr

type Tfpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]uint8
}

type T_G_fpos64_t = Tfpos_t

type Tcookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type T_IO_cookie_io_functions_t = Tcookie_io_functions_t

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

type Tgz_state = struct {
	Fx        TgzFile_s
	Fmode     int32
	Ffd       int32
	Fpath     uintptr
	Fsize     uint32
	Fwant     uint32
	Fin       uintptr
	Fout      uintptr
	Fdirect   int32
	Fhow      int32
	Fstart    Toff_t
	Feof      int32
	Fpast     int32
	Flevel    int32
	Fstrategy int32
	Freset    int32
	Fskip     Toff_t
	Fseek     int32
	Ferr      int32
	Fmsg      uintptr
	Fstrm     Tz_stream
}

type Tgz_statep = uintptr

func XzlibVersion(tls *libc.TLS) (r uintptr) {
	return __ccgo_ts
}

func XzlibCompileFlags(tls *libc.TLS) (r TuLong) {
	var flags TuLong
	_ = flags
	flags = uint64(0)
	switch libc.Int32FromUint64(libc.Uint64FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += uint64(1)
	case int32(8):
		flags += uint64(2)
	default:
		flags += uint64(3)
	}
	switch libc.Int32FromUint64(libc.Uint64FromInt64(8)) {
	case int32(2):
	case int32(4):
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(1) << libc.Int32FromInt32(2))
	case int32(8):
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(2) << libc.Int32FromInt32(2))
	default:
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(3) << libc.Int32FromInt32(2))
	}
	switch libc.Int32FromUint64(libc.Uint64FromInt64(8)) {
	case int32(2):
	case int32(4):
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(1) << libc.Int32FromInt32(4))
	case int32(8):
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(2) << libc.Int32FromInt32(4))
	default:
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(3) << libc.Int32FromInt32(4))
	}
	switch libc.Int32FromUint64(libc.Uint64FromInt64(8)) {
	case int32(2):
	case int32(4):
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(1) << libc.Int32FromInt32(6))
	case int32(8):
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(2) << libc.Int32FromInt32(6))
	default:
		flags += libc.Uint64FromInt32(libc.Int32FromInt32(3) << libc.Int32FromInt32(6))
	}
	/*
		#if defined(ASMV) || defined(ASMINF)
		    flags += 1 << 9;
		#endif
	*/
	return flags
}

// C documentation
//
//	/* exported to allow conversion of error code to string for compress() and
//	 * uncompress()
//	 */
func XzError(tls *libc.TLS, err int32) (r uintptr) {
	var v1 int32
	_ = v1
	if err < -int32(6) || err > int32(2) {
		v1 = int32(9)
	} else {
		v1 = int32(2) - err
	}
	return Xz_errmsg[v1]
}

func Xzcalloc(tls *libc.TLS, opaque Tvoidpf, items uint32, size uint32) (r Tvoidpf) {
	_ = opaque
	return libc.Xmalloc(tls, uint64(items*size))
}

func Xzcfree(tls *libc.TLS, opaque Tvoidpf, ptr Tvoidpf) {
	_ = opaque
	libc.Xfree(tls, ptr)
}

// C documentation
//
//	/* ===========================================================================
//	     Compresses the source buffer into the destination buffer. The level
//	   parameter has the same meaning as in deflateInit.  sourceLen is the byte
//	   length of the source buffer. Upon entry, destLen is the total size of the
//	   destination buffer, which must be at least 0.1% larger than sourceLen plus
//	   12 bytes. Upon exit, destLen is the actual size of the compressed buffer.
//
//	     compress2 returns Z_OK if success, Z_MEM_ERROR if there was not enough
//	   memory, Z_BUF_ERROR if there was not enough room in the output buffer,
//	   Z_STREAM_ERROR if the level parameter is invalid.
//	*/
func Xcompress2(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen TuLong, level int32) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err, v3, v4 int32
	var left TuLong
	var max TuInt
	var v1, v2 uint32
	var _ /* stream at bp+0 */ Tz_stream
	_, _, _, _, _, _, _ = err, left, max, v1, v2, v3, v4
	max = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	left = *(*TuLongf)(unsafe.Pointer(destLen))
	*(*TuLongf)(unsafe.Pointer(destLen)) = uint64(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = XdeflateInit_(tls, bp, level, __ccgo_ts, libc.Int32FromInt64(112))
	if err != m_Z_OK {
		return err
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = dest
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = source
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	for cond := true; cond; cond = err == m_Z_OK {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out == uint32(0) {
			if left > uint64(max) {
				v1 = max
			} else {
				v1 = uint32(left)
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
			left -= uint64((*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out)
		}
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in == uint32(0) {
			if sourceLen > uint64(max) {
				v2 = max
			} else {
				v2 = uint32(sourceLen)
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v2
			sourceLen -= uint64((*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in)
		}
		if sourceLen != 0 {
			v3 = m_Z_NO_FLUSH
		} else {
			v3 = int32(m_Z_FINISH)
		}
		err = Xdeflate(tls, bp, v3)
	}
	*(*TuLongf)(unsafe.Pointer(destLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
	XdeflateEnd(tls, bp)
	if err == int32(m_Z_STREAM_END) {
		v4 = m_Z_OK
	} else {
		v4 = err
	}
	return v4
}

// C documentation
//
//	/* ===========================================================================
//	 */
func Xcompress(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen TuLong) (r int32) {
	return Xcompress2(tls, dest, destLen, source, sourceLen, -int32(1))
}

// C documentation
//
//	/* ===========================================================================
//	     If the default memLevel or windowBits for deflateInit() is changed, then
//	   this function needs to be updated.
//	 */
func XcompressBound(tls *libc.TLS, sourceLen TuLong) (r TuLong) {
	return sourceLen + sourceLen>>libc.Int32FromInt32(12) + sourceLen>>libc.Int32FromInt32(14) + sourceLen>>libc.Int32FromInt32(25) + uint64(13)
}

// C documentation
//
//	/* ===========================================================================
//	     Decompresses the source buffer into the destination buffer.  *sourceLen is
//	   the byte length of the source buffer. Upon entry, *destLen is the total size
//	   of the destination buffer, which must be large enough to hold the entire
//	   uncompressed data. (The size of the uncompressed data must have been saved
//	   previously by the compressor and transmitted to the decompressor by some
//	   mechanism outside the scope of this compression library.) Upon exit,
//	   *destLen is the size of the decompressed data and *sourceLen is the number
//	   of source bytes consumed. Upon return, source + *sourceLen points to the
//	   first unused input byte.
//
//	     uncompress returns Z_OK if success, Z_MEM_ERROR if there was not enough
//	   memory, Z_BUF_ERROR if there was not enough room in the output buffer, or
//	   Z_DATA_ERROR if the input data was corrupted, including if the input data is
//	   an incomplete zlib stream.
//	*/
func Xuncompress2(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen uintptr) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var err, v3, v4, v5 int32
	var left, len1 TuLong
	var max TuInt
	var v1, v2 uint32
	var _ /* buf at bp+112 */ [1]TByte
	var _ /* stream at bp+0 */ Tz_stream
	_, _, _, _, _, _, _, _, _ = err, left, len1, max, v1, v2, v3, v4, v5
	max = libc.Uint32FromInt32(-libc.Int32FromInt32(1)) /* for detection of incomplete stream when *destLen == 0 */
	len1 = *(*TuLong)(unsafe.Pointer(sourceLen))
	if *(*TuLongf)(unsafe.Pointer(destLen)) != 0 {
		left = *(*TuLongf)(unsafe.Pointer(destLen))
		*(*TuLongf)(unsafe.Pointer(destLen)) = uint64(0)
	} else {
		left = uint64(1)
		dest = bp + 112
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = source
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = XinflateInit_(tls, bp, __ccgo_ts, libc.Int32FromInt64(112))
	if err != m_Z_OK {
		return err
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = dest
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(0)
	for cond := true; cond; cond = err == m_Z_OK {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out == uint32(0) {
			if left > uint64(max) {
				v1 = max
			} else {
				v1 = uint32(left)
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
			left -= uint64((*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out)
		}
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in == uint32(0) {
			if len1 > uint64(max) {
				v2 = max
			} else {
				v2 = uint32(len1)
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v2
			len1 -= uint64((*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in)
		}
		err = Xinflate(tls, bp, m_Z_NO_FLUSH)
	}
	*(*TuLong)(unsafe.Pointer(sourceLen)) -= len1 + uint64((*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in)
	if dest != bp+112 {
		*(*TuLongf)(unsafe.Pointer(destLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
	} else {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out != 0 && err == -int32(5) {
			left = uint64(1)
		}
	}
	XinflateEnd(tls, bp)
	if err == int32(m_Z_STREAM_END) {
		v3 = m_Z_OK
	} else {
		if err == int32(m_Z_NEED_DICT) {
			v4 = -int32(3)
		} else {
			if err == -int32(5) && left+uint64((*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out) != 0 {
				v5 = -int32(3)
			} else {
				v5 = err
			}
			v4 = v5
		}
		v3 = v4
	}
	return v3
}

func Xuncompress(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, _sourceLen TuLong) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*TuLong)(unsafe.Pointer(bp)) = _sourceLen
	return Xuncompress2(tls, dest, destLen, source, bp)
}

// C documentation
//
//	/* gzclose() is in a separate file so that it is linked in only if it is used.
//	   That way the other gzclose functions can be used instead to avoid linking in
//	   unneeded compression or decompression routines. */
func Xgzclose(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	var v1 int32
	_, _ = state, v1
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		v1 = Xgzclose_r(tls, file)
	} else {
		v1 = Xgzclose_w(tls, file)
	}
	return v1
}

const m_INT_MAX1 = 2147483647
const m_LSEEK = "lseek"
const m_O_APPEND1 = 1024
const m_O_CLOEXEC1 = 524288
const m_O_CREAT1 = 64
const m_O_EXCL1 = 128
const m_O_LARGEFILE1 = 131072
const m_O_RDONLY1 = 0
const m_O_TRUNC1 = 512
const m_O_WRONLY1 = 1

// C documentation
//
//	/* Reset gzip file state */
func _gz_reset(tls *libc.TLS, state Tgz_statep) {
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)           /* no output data available */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) { /* for reading ... */
		(*Tgz_state)(unsafe.Pointer(state)).Feof = 0      /* not at end of file */
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0     /* have not read past end yet */
		(*Tgz_state)(unsafe.Pointer(state)).Fhow = m_LOOK /* look for gzip header */
	} else { /* for writing ... */
		(*Tgz_state)(unsafe.Pointer(state)).Freset = 0
	} /* no deflateReset pending */
	(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0                   /* no seek request pending */
	Xgz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))         /* clear error */
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos = 0                 /* no uncompressed data yet */
	(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = uint32(0) /* no input data yet */
}

// C documentation
//
//	/* Open a gzip file either by name or file descriptor. */
func _gz_open(tls *libc.TLS, path uintptr, fd int32, mode uintptr) (r TgzFile) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var cloexec, exclusive, oflag, v1, v2, v3, v4, v5 int32
	var len1 Tz_size_t
	var state Tgz_statep
	_, _, _, _, _, _, _, _, _, _ = cloexec, exclusive, len1, oflag, state, v1, v2, v3, v4, v5
	cloexec = 0
	exclusive = 0
	/* check input */
	if path == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	/* allocate gzFile structure to return */
	state = libc.Xmalloc(tls, uint64(240))
	if state == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fsize = uint32(0)               /* no buffers allocated yet */
	(*Tgz_state)(unsafe.Pointer(state)).Fwant = uint32(m_GZBUFSIZE)     /* requested buffer size */
	(*Tgz_state)(unsafe.Pointer(state)).Fmsg = libc.UintptrFromInt32(0) /* no error message yet */
	/* interpret mode */
	(*Tgz_state)(unsafe.Pointer(state)).Fmode = m_GZ_NONE
	(*Tgz_state)(unsafe.Pointer(state)).Flevel = -int32(1)
	(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = m_Z_DEFAULT_STRATEGY
	(*Tgz_state)(unsafe.Pointer(state)).Fdirect = 0
	for *(*uint8)(unsafe.Pointer(mode)) != 0 {
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(mode))) >= int32('0') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(mode))) <= int32('9') {
			(*Tgz_state)(unsafe.Pointer(state)).Flevel = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(mode))) - int32('0')
		} else {
			switch libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(mode))) {
			case int32('r'):
				(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_READ)
			case int32('w'):
				(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_WRITE)
			case int32('a'):
				(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_APPEND)
			case int32('+'): /* can't read and write at the same time */
				libc.Xfree(tls, state)
				return libc.UintptrFromInt32(0)
			case int32('b'): /* ignore -- will request binary anyway */
			case int32('e'):
				cloexec = int32(1)
			case int32('x'):
				exclusive = int32(1)
			case int32('f'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_FILTERED)
			case int32('h'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_HUFFMAN_ONLY)
			case int32('R'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_RLE)
			case int32('F'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_FIXED)
			case int32('T'):
				(*Tgz_state)(unsafe.Pointer(state)).Fdirect = int32(1)
			default: /* could consider as an error, but just ignore */
			}
		}
		mode++
	}
	/* must provide an "r", "w", or "a" */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == m_GZ_NONE {
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	/* can't force transparent read */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		if (*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0 {
			libc.Xfree(tls, state)
			return libc.UintptrFromInt32(0)
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fdirect = int32(1) /* for empty file */
	}
	/* save the path name for error messages */
	len1 = libc.Xstrlen(tls, path)
	(*Tgz_state)(unsafe.Pointer(state)).Fpath = libc.Xmalloc(tls, len1+uint64(1))
	if (*Tgz_state)(unsafe.Pointer(state)).Fpath == libc.UintptrFromInt32(0) {
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	libc.X__builtin_snprintf(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath, len1+uint64(1), __ccgo_ts+584, libc.VaList(bp+8, path))
	/* compute the flags for open() */
	if cloexec != 0 {
		v1 = int32(m_O_CLOEXEC1)
	} else {
		v1 = 0
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		v2 = m_O_RDONLY1
	} else {
		if exclusive != 0 {
			v3 = int32(m_O_EXCL1)
		} else {
			v3 = 0
		}
		if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_WRITE) {
			v4 = int32(m_O_TRUNC1)
		} else {
			v4 = int32(m_O_APPEND1)
		}
		v2 = libc.Int32FromInt32(m_O_WRONLY1) | libc.Int32FromInt32(m_O_CREAT1) | v3 | v4
	}
	oflag = int32(m_O_LARGEFILE1) | v1 | v2
	/* open the file with the appropriate flags (or just use fd) */
	if fd > -int32(1) {
		v5 = fd
	} else {
		v5 = libc.Xopen(tls, path, oflag, libc.VaList(bp+8, int32(0666)))
	}
	(*Tgz_state)(unsafe.Pointer(state)).Ffd = v5
	if (*Tgz_state)(unsafe.Pointer(state)).Ffd == -int32(1) {
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_APPEND) {
		libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, 0, int32(2)) /* so gzoffset() is correct */
		(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_WRITE)          /* simplify later checks */
	}
	/* save the current position for rewinding (only if reading) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		(*Tgz_state)(unsafe.Pointer(state)).Fstart = libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, 0, int32(1))
		if (*Tgz_state)(unsafe.Pointer(state)).Fstart == int64(-int32(1)) {
			(*Tgz_state)(unsafe.Pointer(state)).Fstart = 0
		}
	}
	/* initialize stream */
	_gz_reset(tls, state)
	/* return stream */
	return state
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzopen(tls *libc.TLS, path uintptr, mode uintptr) (r TgzFile) {
	return _gz_open(tls, path, -int32(1), mode)
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzopen64(tls *libc.TLS, path uintptr, mode uintptr) (r TgzFile) {
	return _gz_open(tls, path, -int32(1), mode)
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzdopen(tls *libc.TLS, fd int32, mode uintptr) (r TgzFile) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var gz TgzFile
	var path, v1 uintptr
	var v2 bool
	_, _, _, _ = gz, path, v1, v2
	if v2 = fd == -int32(1); !v2 {
		v1 = libc.Xmalloc(tls, libc.Uint64FromInt32(7)+libc.Uint64FromInt32(3)*libc.Uint64FromInt64(4))
		path = v1
	}
	if v2 || v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	libc.X__builtin_snprintf(tls, path, libc.Uint64FromInt32(7)+libc.Uint64FromInt32(3)*libc.Uint64FromInt64(4), __ccgo_ts+587, libc.VaList(bp+8, fd))
	gz = _gz_open(tls, path, fd, mode)
	libc.Xfree(tls, path)
	return gz
}

/* -- see zlib.h -- */

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzbuffer(tls *libc.TLS, file TgzFile, size uint32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return -int32(1)
	}
	/* make sure we haven't already allocated memory */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != uint32(0) {
		return -int32(1)
	}
	/* check and set requested size */
	if size<<int32(1) < size {
		return -int32(1)
	} /* need to be able to double it */
	if size < uint32(8) {
		size = uint32(8)
	} /* needed to behave well with flushing */
	(*Tgz_state)(unsafe.Pointer(state)).Fwant = size
	return 0
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzrewind(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* back up and start over */
	if libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tgz_state)(unsafe.Pointer(state)).Fstart, 0) == int64(-int32(1)) {
		return -int32(1)
	}
	_gz_reset(tls, state)
	return 0
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzseek64(tls *libc.TLS, file TgzFile, offset Toff_t, whence int32) (r Toff_t) {
	var n, v1 uint32
	var ret Toff_t
	var state Tgz_statep
	_, _, _, _ = n, ret, state, v1
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return int64(-int32(1))
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return int64(-int32(1))
	}
	/* check that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return int64(-int32(1))
	}
	/* can only seek from start or relative to current position */
	if whence != 0 && whence != int32(1) {
		return int64(-int32(1))
	}
	/* normalize offset to a SEEK_CUR specification */
	if whence == 0 {
		offset -= (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos
	} else {
		if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
			offset += (*Tgz_state)(unsafe.Pointer(state)).Fskip
		}
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
	/* if within raw area while reading, just go there */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fhow == int32(m_COPY) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos+offset >= 0 {
		ret = libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, offset-libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave), int32(1))
		if ret == int64(-int32(1)) {
			return int64(-int32(1))
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Feof = 0
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		Xgz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += offset
		return (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos
	}
	/* calculate skip amount, rewinding if needed for back seek when reading */
	if offset < 0 {
		if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) { /* writing -- can't go backwards */
			return int64(-int32(1))
		}
		offset += (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos
		if offset < 0 { /* before start of file! */
			return int64(-int32(1))
		}
		if Xgzrewind(tls, file) == -int32(1) { /* rewind, then skip to offset */
			return int64(-int32(1))
		}
	}
	/* if reading, skip what's in output buffer (one less gzgetc() check) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		if libc.Bool(libc.Bool(uint64(4) == uint64(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > Xgz_intmax(tls)) || libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave) > offset {
			v1 = libc.Uint32FromInt64(offset)
		} else {
			v1 = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
		}
		n = v1
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= n
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(n)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(n)
		offset -= libc.Int64FromUint32(n)
	}
	/* request skip (if not zero) */
	if offset != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = int32(1)
		(*Tgz_state)(unsafe.Pointer(state)).Fskip = offset
	}
	return (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos + offset
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzseek(tls *libc.TLS, file TgzFile, offset Toff_t, whence int32) (r Toff_t) {
	var ret Toff_t
	var v1 int64
	_, _ = ret, v1
	ret = Xgzseek64(tls, file, offset, whence)
	if ret == ret {
		v1 = ret
	} else {
		v1 = int64(-int32(1))
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgztell64(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var state Tgz_statep
	var v1 int64
	_, _ = state, v1
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return int64(-int32(1))
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return int64(-int32(1))
	}
	/* return position */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		v1 = (*Tgz_state)(unsafe.Pointer(state)).Fskip
	} else {
		v1 = 0
	}
	return (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos + v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgztell(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var ret Toff_t
	var v1 int64
	_, _ = ret, v1
	ret = Xgztell64(tls, file)
	if ret == ret {
		v1 = ret
	} else {
		v1 = int64(-int32(1))
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzoffset64(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var offset Toff_t
	var state Tgz_statep
	_, _ = offset, state
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return int64(-int32(1))
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return int64(-int32(1))
	}
	/* compute and return effective offset in file */
	offset = libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, 0, int32(1))
	if offset == int64(-int32(1)) {
		return int64(-int32(1))
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) { /* reading */
		offset -= libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in)
	} /* don't count buffered input */
	return offset
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzoffset(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var ret Toff_t
	var v1 int64
	_, _ = ret, v1
	ret = Xgzoffset64(tls, file)
	if ret == ret {
		v1 = ret
	} else {
		v1 = int64(-int32(1))
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzeof(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	var v1 int32
	_, _ = state, v1
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return 0
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return 0
	}
	/* return end-of-file state */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		v1 = (*Tgz_state)(unsafe.Pointer(state)).Fpast
	} else {
		v1 = 0
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzerror(tls *libc.TLS, file TgzFile, errnum uintptr) (r uintptr) {
	var state Tgz_statep
	var v1, v2 uintptr
	_, _, _ = state, v1, v2
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return libc.UintptrFromInt32(0)
	}
	/* return error information */
	if errnum != libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(errnum)) = (*Tgz_state)(unsafe.Pointer(state)).Ferr
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr == -int32(4) {
		v1 = __ccgo_ts + 595
	} else {
		if (*Tgz_state)(unsafe.Pointer(state)).Fmsg == libc.UintptrFromInt32(0) {
			v2 = __ccgo_ts + 494
		} else {
			v2 = (*Tgz_state)(unsafe.Pointer(state)).Fmsg
		}
		v1 = v2
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzclearerr(tls *libc.TLS, file TgzFile) {
	var state Tgz_statep
	_ = state
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return
	}
	/* clear error and end-of-file */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		(*Tgz_state)(unsafe.Pointer(state)).Feof = 0
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
	}
	Xgz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
}

// C documentation
//
//	/* Create an error message in allocated memory and set state->err and
//	   state->msg accordingly.  Free any previous error message already there.  Do
//	   not try to free or allocate space if the error is Z_MEM_ERROR (out of
//	   memory).  Simply save the error message as a static string.  If there is an
//	   allocation failure constructing the error message, then convert the error to
//	   out of memory. */
func Xgz_error(tls *libc.TLS, state Tgz_statep, err int32, msg uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1 uintptr
	_ = v1
	/* free previously allocated message and clear */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmsg != libc.UintptrFromInt32(0) {
		if (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(4) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fmsg)
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fmsg = libc.UintptrFromInt32(0)
	}
	/* if fatal, set state->x.have to 0 so that the gzgetc() macro fails */
	if err != m_Z_OK && err != -int32(5) {
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
	}
	/* set error code, and if no message, then done */
	(*Tgz_state)(unsafe.Pointer(state)).Ferr = err
	if msg == libc.UintptrFromInt32(0) {
		return
	}
	/* for an out of memory error, return literal string when requested */
	if err == -int32(4) {
		return
	}
	/* construct error message with path */
	v1 = libc.Xmalloc(tls, libc.Xstrlen(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)+libc.Xstrlen(tls, msg)+uint64(3))
	(*Tgz_state)(unsafe.Pointer(state)).Fmsg = v1
	if v1 == libc.UintptrFromInt32(0) {
		(*Tgz_state)(unsafe.Pointer(state)).Ferr = -int32(4)
		return
	}
	libc.X__builtin_snprintf(tls, (*Tgz_state)(unsafe.Pointer(state)).Fmsg, libc.Xstrlen(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)+libc.Xstrlen(tls, msg)+uint64(3), __ccgo_ts+609, libc.VaList(bp+8, (*Tgz_state)(unsafe.Pointer(state)).Fpath, __ccgo_ts+616, msg))
}

// C documentation
//
//	/* portably return maximum value for an int (when limits.h presumed not
//	   available) -- we need to do this to cover cases where 2's complement not
//	   used, since C standard permits 1's complement and sign-bit representations,
//	   otherwise we could just use ((unsigned)-1) >> 1 */
func Xgz_intmax(tls *libc.TLS) (r uint32) {
	return uint32(m_INT_MAX1)
}

const m_INT_MAX2 = 0x7fffffff
const m_O_APPEND2 = 02000
const m_O_CLOEXEC2 = 02000000
const m_O_CREAT2 = 0100
const m_O_EXCL2 = 0200
const m_O_LARGEFILE2 = 0400000
const m_O_RDONLY2 = 00
const m_O_TRUNC2 = 01000
const m_O_WRONLY2 = 01

// C documentation
//
//	/* Use read() to load a buffer -- return -1 on error, otherwise 0.  Read from
//	   state->fd, and update state->eof, state->err, and state->msg as appropriate.
//	   This function needs to loop on read(), since read() is not guaranteed to
//	   read the number of bytes requested, depending on the type of descriptor. */
func _gz_load(tls *libc.TLS, state Tgz_statep, buf uintptr, len1 uint32, have uintptr) (r int32) {
	var get, max uint32
	var ret int32
	_, _, _ = get, max, ret
	max = libc.Uint32FromInt32(-libc.Int32FromInt32(1))>>libc.Int32FromInt32(2) + libc.Uint32FromInt32(1)
	*(*uint32)(unsafe.Pointer(have)) = uint32(0)
	for cond := true; cond; cond = *(*uint32)(unsafe.Pointer(have)) < len1 {
		get = len1 - *(*uint32)(unsafe.Pointer(have))
		if get > max {
			get = max
		}
		ret = int32(libc.Xread(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, buf+uintptr(*(*uint32)(unsafe.Pointer(have))), uint64(get)))
		if ret <= 0 {
			break
		}
		*(*uint32)(unsafe.Pointer(have)) += libc.Uint32FromInt32(ret)
	}
	if ret < 0 {
		Xgz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return -int32(1)
	}
	if ret == 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Feof = int32(1)
	}
	return 0
}

// C documentation
//
//	/* Load up input buffer and set eof flag if last data loaded -- return -1 on
//	   error, 0 otherwise.  Note that the eof flag is set when the end of the input
//	   file is reached, even though there may be unused data in the buffer.  Once
//	   that data has been used, no more attempts will be made to read the file.
//	   If strm->avail_in != 0, then the current data is moved to the beginning of
//	   the input buffer, and then the remainder of the buffer is loaded with the
//	   available data from the input file. */
func _gz_avail(tls *libc.TLS, state Tgz_statep) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var n, v1 uint32
	var p, q, v3, v4 uintptr
	var strm Tz_streamp
	var _ /* got at bp+0 */ uint32
	_, _, _, _, _, _, _ = n, p, q, strm, v1, v3, v4
	strm = state + 128
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Feof == 0 {
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 { /* copy what's there to the start */
			p = (*Tgz_state)(unsafe.Pointer(state)).Fin
			q = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
			n = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			for {
				v3 = p
				p++
				v4 = q
				q++
				*(*uint8)(unsafe.Pointer(v3)) = *(*uint8)(unsafe.Pointer(v4))
				goto _2
			_2:
				;
				n--
				v1 = n
				if !(v1 != 0) {
					break
				}
			}
		}
		if _gz_load(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in), (*Tgz_state)(unsafe.Pointer(state)).Fsize-(*Tz_stream)(unsafe.Pointer(strm)).Favail_in, bp) == -int32(1) {
			return -int32(1)
		}
		*(*TuInt)(unsafe.Pointer(strm + 8)) += *(*uint32)(unsafe.Pointer(bp))
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
	}
	return 0
}

// C documentation
//
//	/* Look for gzip header, set up for inflate or copy.  state->x.have must be 0.
//	   If this is the first time in, allocate required memory.  state->how will be
//	   left unchanged if there is no more input data available, will be set to COPY
//	   if there is no gzip header and direct copying will be performed, or it will
//	   be set to GZIP for decompression.  If direct copying, then leftover input
//	   data from the input buffer will be copied to the output buffer.  In that
//	   case, all further file reads will be directly to either the output buffer or
//	   a user buffer.  If decompressing, the inflate state will be initialized.
//	   gz_look() will return 0 on success or -1 on failure. */
func _gz_look(tls *libc.TLS, state Tgz_statep) (r int32) {
	var strm Tz_streamp
	_ = strm
	strm = state + 128
	/* allocate read buffers and inflate memory */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) {
		/* allocate buffers */
		(*Tgz_state)(unsafe.Pointer(state)).Fin = libc.Xmalloc(tls, uint64((*Tgz_state)(unsafe.Pointer(state)).Fwant))
		(*Tgz_state)(unsafe.Pointer(state)).Fout = libc.Xmalloc(tls, uint64((*Tgz_state)(unsafe.Pointer(state)).Fwant<<int32(1)))
		if (*Tgz_state)(unsafe.Pointer(state)).Fin == libc.UintptrFromInt32(0) || (*Tgz_state)(unsafe.Pointer(state)).Fout == libc.UintptrFromInt32(0) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
			return -int32(1)
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fsize = (*Tgz_state)(unsafe.Pointer(state)).Fwant
		/* allocate inflate memory */
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fzalloc = uintptr(m_Z_NULL)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fzfree = uintptr(m_Z_NULL)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fopaque = uintptr(m_Z_NULL)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = uintptr(m_Z_NULL)
		if XinflateInit2_(tls, state+128, libc.Int32FromInt32(15)+libc.Int32FromInt32(16), __ccgo_ts, libc.Int32FromInt64(112)) != m_Z_OK { /* gunzip */
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			(*Tgz_state)(unsafe.Pointer(state)).Fsize = uint32(0)
			Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
			return -int32(1)
		}
	}
	/* get at least the magic bytes in the input buffer */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in < uint32(2) {
		if _gz_avail(tls, state) == -int32(1) {
			return -int32(1)
		}
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			return 0
		}
	}
	/* look for gzip magic bytes -- if there, do gzip decoding (note: there is
	   a logical dilemma here when considering the case of a partially written
	   gzip file, to wit, if a single 31 byte is written, then we cannot tell
	   whether this is a single-byte file, or just a partially written gzip
	   file -- for here we assume that if a gzip file is being written, then
	   the header will be written in a single operation, so that reading a
	   single byte is sufficient indication that it is not a gzip file) */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in > uint32(1) && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in))) == int32(31) && libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in + 1))) == int32(139) {
		XinflateReset(tls, strm)
		(*Tgz_state)(unsafe.Pointer(state)).Fhow = int32(m_GZIP)
		(*Tgz_state)(unsafe.Pointer(state)).Fdirect = 0
		return 0
	}
	/* no gzip header -- if we were decoding gzip before, then this is trailing
	   garbage.  Ignore the trailing garbage and finish. */
	if (*Tgz_state)(unsafe.Pointer(state)).Fdirect == 0 {
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Feof = int32(1)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
		return 0
	}
	/* doing raw i/o, copy any leftover input to output -- this assumes that
	   the output buffer is larger than the input buffer, which also assures
	   space for gzungetc() */
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout
	libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, uint64((*Tz_stream)(unsafe.Pointer(strm)).Favail_in))
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = uint32(0)
	(*Tgz_state)(unsafe.Pointer(state)).Fhow = int32(m_COPY)
	(*Tgz_state)(unsafe.Pointer(state)).Fdirect = int32(1)
	return 0
}

// C documentation
//
//	/* Decompress from input to the provided next_out and avail_out in the state.
//	   On return, state->x.have and state->x.next point to the just decompressed
//	   data.  If the gzip stream completes, state->how is reset to LOOK to look for
//	   the next gzip stream or raw data, once state->x.have is depleted.  Returns 0
//	   on success, -1 on failure. */
func _gz_decomp(tls *libc.TLS, state Tgz_statep) (r int32) {
	var had uint32
	var ret int32
	var strm Tz_streamp
	var v1 uintptr
	_, _, _, _ = had, ret, strm, v1
	ret = m_Z_OK
	strm = state + 128
	/* fill output buffer up to end of deflate stream */
	had = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	for cond := true; cond; cond = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out != 0 && ret != int32(m_Z_STREAM_END) {
		/* get more input for inflate() */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) && _gz_avail(tls, state) == -int32(1) {
			return -int32(1)
		}
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			Xgz_error(tls, state, -int32(5), __ccgo_ts+619)
			break
		}
		/* decompress and handle errors */
		ret = Xinflate(tls, strm, m_Z_NO_FLUSH)
		if ret == -int32(2) || ret == int32(m_Z_NEED_DICT) {
			Xgz_error(tls, state, -int32(2), __ccgo_ts+642)
			return -int32(1)
		}
		if ret == -int32(4) {
			Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
			return -int32(1)
		}
		if ret == -int32(3) { /* deflate stream invalid */
			if (*Tz_stream)(unsafe.Pointer(strm)).Fmsg == libc.UintptrFromInt32(0) {
				v1 = __ccgo_ts + 681
			} else {
				v1 = (*Tz_stream)(unsafe.Pointer(strm)).Fmsg
			}
			Xgz_error(tls, state, -int32(3), v1)
			return -int32(1)
		}
	}
	/* update available output */
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = had - (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out - uintptr((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave)
	/* if the gzip stream completed successfully, look for another */
	if ret == int32(m_Z_STREAM_END) {
		(*Tgz_state)(unsafe.Pointer(state)).Fhow = m_LOOK
	}
	/* good decompression */
	return 0
}

// C documentation
//
//	/* Fetch data and put it in the output buffer.  Assumes state->x.have is 0.
//	   Data is either copied from the input file or decompressed from the input
//	   file depending on state->how.  If state->how is LOOK, then a gzip header is
//	   looked for to determine whether to copy or decompress.  Returns -1 on error,
//	   otherwise 0.  gz_fetch() will leave state->how as COPY or GZIP unless the
//	   end of the input file has been reached and all data has been processed.  */
func _gz_fetch(tls *libc.TLS, state Tgz_statep) (r int32) {
	var strm Tz_streamp
	_ = strm
	strm = state + 128
	for cond := true; cond; cond = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) && (!((*Tgz_state)(unsafe.Pointer(state)).Feof != 0) || (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0) {
		switch (*Tgz_state)(unsafe.Pointer(state)).Fhow {
		case m_LOOK: /* -> LOOK, COPY (only if never GZIP), or GZIP */
			if _gz_look(tls, state) == -int32(1) {
				return -int32(1)
			}
			if (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK {
				return 0
			}
		case int32(m_COPY): /* -> COPY */
			if _gz_load(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fout, (*Tgz_state)(unsafe.Pointer(state)).Fsize<<int32(1), state) == -int32(1) {
				return -int32(1)
			}
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout
			return 0
		case int32(m_GZIP): /* -> GZIP or LOOK (if end of gzip stream) */
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = (*Tgz_state)(unsafe.Pointer(state)).Fsize << int32(1)
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = (*Tgz_state)(unsafe.Pointer(state)).Fout
			if _gz_decomp(tls, state) == -int32(1) {
				return -int32(1)
			}
		}
	}
	return 0
}

// C documentation
//
//	/* Skip len uncompressed bytes of output.  Return -1 on error, 0 on success. */
func _gz_skip(tls *libc.TLS, state Tgz_statep, len1 Toff_t) (r int32) {
	var n, v1 uint32
	_, _ = n, v1
	/* skip over len bytes or reach end-of-file, whichever comes first */
	for len1 != 0 {
		/* skip over whatever is in output buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
			if libc.Bool(libc.Bool(uint64(4) == uint64(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > Xgz_intmax(tls)) || libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave) > len1 {
				v1 = libc.Uint32FromInt64(len1)
			} else {
				v1 = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			n = v1
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(n)
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(n)
			len1 -= libc.Int64FromUint32(n)
		} else {
			if (*Tgz_state)(unsafe.Pointer(state)).Feof != 0 && (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				break
			} else {
				/* get more output, looking for header if required */
				if _gz_fetch(tls, state) == -int32(1) {
					return -int32(1)
				}
			}
		}
	}
	return 0
}

// C documentation
//
//	/* Read len bytes into buf from file, or less than len up to the end of the
//	   input.  Return the number of bytes read.  If zero is returned, either the
//	   end of file was reached, or there was an error.  state->err must be
//	   consulted in that case to determine which. */
func _gz_read(tls *libc.TLS, state Tgz_statep, buf Tvoidp, len1 Tz_size_t) (r Tz_size_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var got Tz_size_t
	var _ /* n at bp+0 */ uint32
	_ = got
	/* if len is zero, avoid unnecessary operations */
	if len1 == uint64(0) {
		return uint64(0)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return uint64(0)
		}
	}
	/* get len bytes to buf, or less than len if at the end */
	got = uint64(0)
	for cond := true; cond; cond = len1 != 0 {
		/* set n to the maximum amount of len that fits in an unsigned int */
		*(*uint32)(unsafe.Pointer(bp)) = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
		if uint64(*(*uint32)(unsafe.Pointer(bp))) > len1 {
			*(*uint32)(unsafe.Pointer(bp)) = uint32(len1)
		}
		/* first just try copying data from the output buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave < *(*uint32)(unsafe.Pointer(bp)) {
				*(*uint32)(unsafe.Pointer(bp)) = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			libc.Xmemcpy(tls, buf, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, uint64(*(*uint32)(unsafe.Pointer(bp))))
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(*(*uint32)(unsafe.Pointer(bp)))
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= *(*uint32)(unsafe.Pointer(bp))
		} else {
			if (*Tgz_state)(unsafe.Pointer(state)).Feof != 0 && (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				(*Tgz_state)(unsafe.Pointer(state)).Fpast = int32(1) /* tried to read past end */
				break
			} else {
				if (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK || *(*uint32)(unsafe.Pointer(bp)) < (*Tgz_state)(unsafe.Pointer(state)).Fsize<<int32(1) {
					/* get more output, looking for header if required */
					if _gz_fetch(tls, state) == -int32(1) {
						return uint64(0)
					}
					continue /* no progress yet -- go back to copy above */
					/* the copy above assures that we will leave with space in the
					   output buffer, allowing at least one gzungetc() to succeed */
				} else {
					if (*Tgz_state)(unsafe.Pointer(state)).Fhow == int32(m_COPY) { /* read directly */
						if _gz_load(tls, state, buf, *(*uint32)(unsafe.Pointer(bp)), bp) == -int32(1) {
							return uint64(0)
						}
					} else { /* state->how == GZIP */
						(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_out = *(*uint32)(unsafe.Pointer(bp))
						(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_out = buf
						if _gz_decomp(tls, state) == -int32(1) {
							return uint64(0)
						}
						*(*uint32)(unsafe.Pointer(bp)) = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
						(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
					}
				}
			}
		}
		/* update progress */
		len1 -= uint64(*(*uint32)(unsafe.Pointer(bp)))
		buf = buf + uintptr(*(*uint32)(unsafe.Pointer(bp)))
		got += uint64(*(*uint32)(unsafe.Pointer(bp)))
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(*(*uint32)(unsafe.Pointer(bp)))
	}
	/* return number of bytes read into user buffer */
	return got
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzread(tls *libc.TLS, file TgzFile, buf Tvoidp, len1 uint32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* since an int is returned, make sure len fits in one, otherwise return
	   with an error (this avoids a flaw in the interface) */
	if libc.Int32FromUint32(len1) < 0 {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+703)
		return -int32(1)
	}
	/* read len or fewer bytes to buf */
	len1 = uint32(_gz_read(tls, state, buf, uint64(len1)))
	/* check for an error */
	if len1 == uint32(0) && (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* return the number of bytes read (this is assured to fit in an int) */
	return libc.Int32FromUint32(len1)
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzfread(tls *libc.TLS, buf Tvoidp, size Tz_size_t, nitems Tz_size_t, file TgzFile) (r Tz_size_t) {
	var len1 Tz_size_t
	var state Tgz_statep
	var v1 uint64
	_, _, _ = len1, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return uint64(0)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return uint64(0)
	}
	/* compute bytes to read -- error on overflow */
	len1 = nitems * size
	if size != 0 && len1/size != nitems {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+734)
		return uint64(0)
	}
	/* read len or fewer bytes to buf, return the number of full items read */
	if len1 != 0 {
		v1 = _gz_read(tls, state, buf, len1) / size
	} else {
		v1 = uint64(0)
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzgetc(tls *libc.TLS, file TgzFile) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var state Tgz_statep
	var v1, v2 uintptr
	var v3 int32
	var _ /* buf at bp+0 */ [1]uint8
	_, _, _, _ = state, v1, v2, v3
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* try output buffer (no need to check for skip request) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave--
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos++
		v2 = state + 8
		v1 = *(*uintptr)(unsafe.Pointer(v2))
		*(*uintptr)(unsafe.Pointer(v2))++
		return libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v1)))
	}
	/* nothing there -- try gz_read() */
	if _gz_read(tls, state, bp, uint64(1)) < uint64(1) {
		v3 = -int32(1)
	} else {
		v3 = libc.Int32FromUint8((*(*[1]uint8)(unsafe.Pointer(bp)))[0])
	}
	return v3
}

func Xgzgetc_(tls *libc.TLS, file TgzFile) (r int32) {
	return Xgzgetc(tls, file)
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzungetc(tls *libc.TLS, c int32, file TgzFile) (r int32) {
	var dest, src, v1, v2 uintptr
	var state Tgz_statep
	_, _, _, _, _ = dest, src, state, v1, v2
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* in case this was just opened, set up the input buffer */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) {
		_gz_look(tls, state)
	}
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return -int32(1)
		}
	}
	/* can't push EOF */
	if c < 0 {
		return -int32(1)
	}
	/* if output buffer empty, put byte at end (allows more pushing) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) {
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(1)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize<<libc.Int32FromInt32(1)) - uintptr(1)
		*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) = libc.Uint8FromInt32(c)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos--
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
		return c
	}
	/* if no room, give up (must have already done a gzungetc()) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == (*Tgz_state)(unsafe.Pointer(state)).Fsize<<int32(1) {
		Xgz_error(tls, state, -int32(3), __ccgo_ts+767)
		return -int32(1)
	}
	/* slide output data if needed and insert byte before existing data */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext == (*Tgz_state)(unsafe.Pointer(state)).Fout {
		src = (*Tgz_state)(unsafe.Pointer(state)).Fout + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave)
		dest = (*Tgz_state)(unsafe.Pointer(state)).Fout + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize<<libc.Int32FromInt32(1))
		for src > (*Tgz_state)(unsafe.Pointer(state)).Fout {
			dest--
			v1 = dest
			src--
			v2 = src
			*(*uint8)(unsafe.Pointer(v1)) = *(*uint8)(unsafe.Pointer(v2))
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = dest
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave++
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext--
	*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) = libc.Uint8FromInt32(c)
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos--
	(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
	return c
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzgets(tls *libc.TLS, file TgzFile, buf uintptr, len1 int32) (r uintptr) {
	var eol, str uintptr
	var left, n, v1 uint32
	var state Tgz_statep
	_, _, _, _, _, _ = eol, left, n, state, str, v1
	/* check parameters and get internal structure */
	if file == libc.UintptrFromInt32(0) || buf == libc.UintptrFromInt32(0) || len1 < int32(1) {
		return libc.UintptrFromInt32(0)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return libc.UintptrFromInt32(0)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return libc.UintptrFromInt32(0)
		}
	}
	/* copy output bytes up to new line or len - 1, whichever comes first --
	   append a terminating zero to the string (we don't check for a zero in
	   the contents, let the user worry about that) */
	str = buf
	left = libc.Uint32FromInt32(len1) - uint32(1)
	if left != 0 {
		for cond := true; cond; cond = left != 0 && eol == libc.UintptrFromInt32(0) {
			/* assure that something is in the output buffer */
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) && _gz_fetch(tls, state) == -int32(1) {
				return libc.UintptrFromInt32(0)
			} /* error */
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) { /* end of file */
				(*Tgz_state)(unsafe.Pointer(state)).Fpast = int32(1) /* read past end */
				break                                                /* return what we have */
			}
			/* look for end-of-line in current output buffer */
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > left {
				v1 = left
			} else {
				v1 = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			n = v1
			eol = libc.Xmemchr(tls, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, int32('\n'), uint64(n))
			if eol != libc.UintptrFromInt32(0) {
				n = libc.Uint32FromInt64(int64(eol)-int64((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) + uint32(1)
			}
			/* copy through end-of-line, or remainder if not found */
			libc.Xmemcpy(tls, buf, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, uint64(n))
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(n)
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(n)
			left -= n
			buf += uintptr(n)
		}
	}
	/* return terminated string, or if nothing, end of file */
	if buf == str {
		return libc.UintptrFromInt32(0)
	}
	*(*uint8)(unsafe.Pointer(buf)) = uint8(0)
	return str
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzdirect(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return 0
	}
	state = file
	/* if the state is not known, but we can find out, then do so (this is
	   mainly for right after a gzopen() or gzdopen()) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) {
		_gz_look(tls, state)
	}
	/* return 1 if transparent, 0 if processing a gzip stream */
	return (*Tgz_state)(unsafe.Pointer(state)).Fdirect
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzclose_r(tls *libc.TLS, file TgzFile) (r int32) {
	var err, ret, v1, v2 int32
	var state Tgz_statep
	_, _, _, _, _ = err, ret, state, v1, v2
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	/* check that we're reading */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) {
		return -int32(2)
	}
	/* free memory and close file */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		XinflateEnd(tls, state+128)
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr == -int32(5) {
		v1 = -int32(5)
	} else {
		v1 = m_Z_OK
	}
	err = v1
	Xgz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
	libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)
	ret = libc.Xclose(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd)
	libc.Xfree(tls, state)
	if ret != 0 {
		v2 = -int32(1)
	} else {
		v2 = err
	}
	return v2
}

// C documentation
//
//	/* Initialize state for writing a gzip file.  Mark initialization by setting
//	   state->size to non-zero.  Return -1 on a memory allocation failure, or 0 on
//	   success. */
func _gz_init(tls *libc.TLS, state Tgz_statep) (r int32) {
	var ret int32
	var strm Tz_streamp
	_, _ = ret, strm
	strm = state + 128
	/* allocate input buffer (double size for gzprintf) */
	(*Tgz_state)(unsafe.Pointer(state)).Fin = libc.Xmalloc(tls, uint64((*Tgz_state)(unsafe.Pointer(state)).Fwant<<int32(1)))
	if (*Tgz_state)(unsafe.Pointer(state)).Fin == libc.UintptrFromInt32(0) {
		Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
		return -int32(1)
	}
	/* only need output buffer and deflate state if compressing */
	if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
		/* allocate output buffer */
		(*Tgz_state)(unsafe.Pointer(state)).Fout = libc.Xmalloc(tls, uint64((*Tgz_state)(unsafe.Pointer(state)).Fwant))
		if (*Tgz_state)(unsafe.Pointer(state)).Fout == libc.UintptrFromInt32(0) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
			return -int32(1)
		}
		/* allocate deflate memory, set up for gzip compression */
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = uintptr(m_Z_NULL)
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = uintptr(m_Z_NULL)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = uintptr(m_Z_NULL)
		ret = XdeflateInit2_(tls, strm, (*Tgz_state)(unsafe.Pointer(state)).Flevel, int32(m_Z_DEFLATED), libc.Int32FromInt32(m_MAX_WBITS)+libc.Int32FromInt32(16), int32(m_DEF_MEM_LEVEL), (*Tgz_state)(unsafe.Pointer(state)).Fstrategy, __ccgo_ts, libc.Int32FromInt64(112))
		if ret != m_Z_OK {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
			return -int32(1)
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = libc.UintptrFromInt32(0)
	}
	/* mark state as initialized */
	(*Tgz_state)(unsafe.Pointer(state)).Fsize = (*Tgz_state)(unsafe.Pointer(state)).Fwant
	/* initialize write buffer if compressing */
	if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = (*Tgz_state)(unsafe.Pointer(state)).Fout
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
	}
	return 0
}

// C documentation
//
//	/* Compress whatever is at avail_in and next_in and write to the output file.
//	   Return -1 if there is an error writing to the output file or if gz_init()
//	   fails to allocate memory, otherwise 0.  flush is assumed to be a valid
//	   deflate() flush value.  If flush is Z_FINISH, then the deflate() state is
//	   reset to start a new gzip stream.  If gz->direct is true, then simply write
//	   to the output file without compressing, and ignore flush. */
func _gz_comp(tls *libc.TLS, state Tgz_statep, flush int32) (r int32) {
	var have, max, put, v1, v2 uint32
	var ret, writ int32
	var strm Tz_streamp
	_, _, _, _, _, _, _, _ = have, max, put, ret, strm, writ, v1, v2
	max = libc.Uint32FromInt32(-libc.Int32FromInt32(1))>>libc.Int32FromInt32(2) + libc.Uint32FromInt32(1)
	strm = state + 128
	/* allocate memory if this is the first time through */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return -int32(1)
	}
	/* write directly if requested */
	if (*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0 {
		for (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 {
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in > max {
				v1 = max
			} else {
				v1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			}
			put = v1
			writ = int32(libc.Xwrite(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, uint64(put)))
			if writ < 0 {
				Xgz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
				return -int32(1)
			}
			*(*TuInt)(unsafe.Pointer(strm + 8)) -= libc.Uint32FromInt32(writ)
			*(*uintptr)(unsafe.Pointer(strm)) += uintptr(writ)
		}
		return 0
	}
	/* check for a pending reset */
	if (*Tgz_state)(unsafe.Pointer(state)).Freset != 0 {
		/* don't start a new gzip member unless there is data to write */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			return 0
		}
		XdeflateReset(tls, strm)
		(*Tgz_state)(unsafe.Pointer(state)).Freset = 0
	}
	/* run deflate() on provided input until it produces no more output */
	ret = m_Z_OK
	for cond := true; cond; cond = have != 0 {
		/* write out current buffer contents if full, or if flushing, but if
		   doing Z_FINISH then don't write until we get to Z_STREAM_END */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) || flush != m_Z_NO_FLUSH && (flush != int32(m_Z_FINISH) || ret == int32(m_Z_STREAM_END)) {
			for (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out > (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext {
				if int64((*Tz_stream)(unsafe.Pointer(strm)).Fnext_out)-int64((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext) > int64(libc.Int32FromUint32(max)) {
					v2 = max
				} else {
					v2 = libc.Uint32FromInt64(int64((*Tz_stream)(unsafe.Pointer(strm)).Fnext_out) - int64((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext))
				}
				put = v2
				writ = int32(libc.Xwrite(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, uint64(put)))
				if writ < 0 {
					Xgz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
					return -int32(1)
				}
				(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(writ)
			}
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
				(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = (*Tgz_state)(unsafe.Pointer(state)).Fsize
				(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = (*Tgz_state)(unsafe.Pointer(state)).Fout
				(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout
			}
		}
		/* compress */
		have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
		ret = Xdeflate(tls, strm, flush)
		if ret == -int32(2) {
			Xgz_error(tls, state, -int32(2), __ccgo_ts+798)
			return -int32(1)
		}
		have -= (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	}
	/* if that completed a deflate stream, allow another to start */
	if flush == int32(m_Z_FINISH) {
		(*Tgz_state)(unsafe.Pointer(state)).Freset = int32(1)
	}
	/* all done, no errors */
	return 0
}

// C documentation
//
//	/* Compress len zeros to output.  Return -1 on a write error or memory
//	   allocation failure by gz_comp(), or 0 on success. */
func _gz_zero(tls *libc.TLS, state Tgz_statep, len1 Toff_t) (r int32) {
	var first int32
	var n, v1 uint32
	var strm Tz_streamp
	_, _, _, _ = first, n, strm, v1
	strm = state + 128
	/* consume whatever's left in the input buffer */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
		return -int32(1)
	}
	/* compress len zeros (len guaranteed > 0) */
	first = int32(1)
	for len1 != 0 {
		if libc.Bool(libc.Bool(uint64(4) == uint64(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fsize > Xgz_intmax(tls)) || libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fsize) > len1 {
			v1 = libc.Uint32FromInt64(len1)
		} else {
			v1 = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		}
		n = v1
		if first != 0 {
			libc.Xmemset(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin, 0, uint64(n))
			first = 0
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = n
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(n)
		if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return -int32(1)
		}
		len1 -= libc.Int64FromUint32(n)
	}
	return 0
}

// C documentation
//
//	/* Write len bytes from buf to file.  Return the number of bytes written.  If
//	   the returned value is less than len, then there was an error. */
func _gz_write(tls *libc.TLS, state Tgz_statep, buf Tvoidpc, len1 Tz_size_t) (r Tz_size_t) {
	var copy1, have, n uint32
	var put Tz_size_t
	_, _, _, _ = copy1, have, n, put
	put = len1
	/* if len is zero, avoid unnecessary operations */
	if len1 == uint64(0) {
		return uint64(0)
	}
	/* allocate memory if this is the first time through */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return uint64(0)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return uint64(0)
		}
	}
	/* for small len, copy to input buffer, otherwise compress directly */
	if len1 < uint64((*Tgz_state)(unsafe.Pointer(state)).Fsize) {
		/* copy to input buffer, compress when full */
		for cond := true; cond; cond = len1 != 0 {
			if (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
			}
			have = libc.Uint32FromInt64(int64((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in+uintptr((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in)) - int64((*Tgz_state)(unsafe.Pointer(state)).Fin))
			copy1 = (*Tgz_state)(unsafe.Pointer(state)).Fsize - have
			if uint64(copy1) > len1 {
				copy1 = uint32(len1)
			}
			libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr(have), buf, uint64(copy1))
			(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in += copy1
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(copy1)
			buf = buf + uintptr(copy1)
			len1 -= uint64(copy1)
			if len1 != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
				return uint64(0)
			}
		}
	} else {
		/* consume whatever's left in the input buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return uint64(0)
		}
		/* directly compress user buffer to file */
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = buf
		for cond := true; cond; cond = len1 != 0 {
			n = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
			if uint64(n) > len1 {
				n = uint32(len1)
			}
			(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(n)
			if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
				return uint64(0)
			}
			len1 -= uint64(n)
		}
	}
	/* input was all buffered or compressed */
	return put
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzwrite(tls *libc.TLS, file TgzFile, buf Tvoidpc, len1 uint32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return 0
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return 0
	}
	/* since an int is returned, make sure len fits in one, otherwise return
	   with an error (this avoids a flaw in the interface) */
	if libc.Int32FromUint32(len1) < 0 {
		Xgz_error(tls, state, -int32(3), __ccgo_ts+837)
		return 0
	}
	/* write len bytes from buf (the return value will fit in an int) */
	return libc.Int32FromUint64(_gz_write(tls, state, buf, uint64(len1)))
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzfwrite(tls *libc.TLS, buf Tvoidpc, size Tz_size_t, nitems Tz_size_t, file TgzFile) (r Tz_size_t) {
	var len1 Tz_size_t
	var state Tgz_statep
	var v1 uint64
	_, _, _ = len1, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return uint64(0)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return uint64(0)
	}
	/* compute bytes to read -- error on overflow */
	len1 = nitems * size
	if size != 0 && len1/size != nitems {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+734)
		return uint64(0)
	}
	/* write len bytes to buf, return the number of full items written */
	if len1 != 0 {
		v1 = _gz_write(tls, state, buf, len1) / size
	} else {
		v1 = uint64(0)
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzputc(tls *libc.TLS, file TgzFile, c int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var have uint32
	var state Tgz_statep
	var strm Tz_streamp
	var _ /* buf at bp+0 */ [1]uint8
	_, _, _ = have, state, strm
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	strm = state + 128
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(1)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return -int32(1)
		}
	}
	/* try writing to input buffer for speed (state->size == 0 if buffer not
	   initialized) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
		}
		have = libc.Uint32FromInt64(int64((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in+uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in)) - int64((*Tgz_state)(unsafe.Pointer(state)).Fin))
		if have < (*Tgz_state)(unsafe.Pointer(state)).Fsize {
			*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fin + uintptr(have))) = libc.Uint8FromInt32(c)
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in++
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos++
			return c & int32(0xff)
		}
	}
	/* no room in buffer or not initialized, use gz_write() */
	(*(*[1]uint8)(unsafe.Pointer(bp)))[0] = libc.Uint8FromInt32(c)
	if _gz_write(tls, state, bp, uint64(1)) != uint64(1) {
		return -int32(1)
	}
	return c & int32(0xff)
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzputs(tls *libc.TLS, file TgzFile, s uintptr) (r int32) {
	var len1, put Tz_size_t
	var state Tgz_statep
	var v1 int32
	_, _, _, _ = len1, put, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(1)
	}
	/* write string */
	len1 = libc.Xstrlen(tls, s)
	if libc.Int32FromUint64(len1) < 0 || uint64(uint32(len1)) != len1 {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+874)
		return -int32(1)
	}
	put = _gz_write(tls, state, s, len1)
	if put < len1 {
		v1 = -int32(1)
	} else {
		v1 = libc.Int32FromUint64(len1)
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzvprintf(tls *libc.TLS, file TgzFile, format uintptr, va Tva_list) (r int32) {
	var left uint32
	var len1 int32
	var next uintptr
	var state Tgz_statep
	var strm Tz_streamp
	_, _, _, _, _ = left, len1, next, state, strm
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	strm = state + 128
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(2)
	}
	/* make sure we have some buffer space */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return (*Tgz_state)(unsafe.Pointer(state)).Ferr
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* do the printf() into the input buffer, put length in len -- the input
	   buffer is double-sized just for this function, so there is guaranteed to
	   be state->size bytes available after the current contents */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
	}
	next = (*Tgz_state)(unsafe.Pointer(state)).Fin + uintptr(int64((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in)-int64((*Tgz_state)(unsafe.Pointer(state)).Fin)) + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
	*(*uint8)(unsafe.Pointer(next + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize-uint32(1)))) = uint8(0)
	len1 = libc.X__builtin_vsnprintf(tls, next, uint64((*Tgz_state)(unsafe.Pointer(state)).Fsize), format, va)
	/* check that printf() results fit in buffer */
	if len1 == 0 || libc.Uint32FromInt32(len1) >= (*Tgz_state)(unsafe.Pointer(state)).Fsize || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(next + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize-uint32(1))))) != 0 {
		return 0
	}
	/* update buffer and position, compress first half if past that */
	*(*TuInt)(unsafe.Pointer(strm + 8)) += libc.Uint32FromInt32(len1)
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(len1)
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in >= (*Tgz_state)(unsafe.Pointer(state)).Fsize {
		left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in - (*Tgz_state)(unsafe.Pointer(state)).Fsize
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
		libc.Xmemmove(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize), uint64(left))
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = left
	}
	return len1
}

func Xgzprintf(tls *libc.TLS, file TgzFile, format uintptr, va1 uintptr) (r int32) {
	var ret int32
	var va Tva_list
	_, _ = ret, va
	va = va1
	ret = Xgzvprintf(tls, file, format, va)
	_ = va
	return ret
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzflush(tls *libc.TLS, file TgzFile, flush int32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(2)
	}
	/* check flush parameter */
	if flush < 0 || flush > int32(m_Z_FINISH) {
		return -int32(2)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* compress remaining data with requested flush */
	_gz_comp(tls, state, flush)
	return (*Tgz_state)(unsafe.Pointer(state)).Ferr
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzsetparams(tls *libc.TLS, file TgzFile, level int32, strategy int32) (r int32) {
	var state Tgz_statep
	var strm Tz_streamp
	_, _ = state, strm
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	strm = state + 128
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK || (*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0 {
		return -int32(2)
	}
	/* if no change is requested, then do nothing */
	if level == (*Tgz_state)(unsafe.Pointer(state)).Flevel && strategy == (*Tgz_state)(unsafe.Pointer(state)).Fstrategy {
		return m_Z_OK
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* change compression parameters for subsequent input */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		/* flush previous input with previous parameters before changing */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 && _gz_comp(tls, state, int32(m_Z_BLOCK)) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
		XdeflateParams(tls, strm, level, strategy)
	}
	(*Tgz_state)(unsafe.Pointer(state)).Flevel = level
	(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = strategy
	return m_Z_OK
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzclose_w(tls *libc.TLS, file TgzFile) (r int32) {
	var ret int32
	var state Tgz_statep
	_, _ = ret, state
	ret = m_Z_OK
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	/* check that we're writing */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return -int32(2)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			ret = (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* flush, free memory, and close file */
	if _gz_comp(tls, state, int32(m_Z_FINISH)) == -int32(1) {
		ret = (*Tgz_state)(unsafe.Pointer(state)).Ferr
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
			XdeflateEnd(tls, state+128)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
		}
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
	}
	Xgz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
	libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)
	if libc.Xclose(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd) == -int32(1) {
		ret = -int32(1)
	}
	libc.Xfree(tls, state)
	return ret
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var x__dist_code = [512]Tuch{
	1:   uint8(1),
	2:   uint8(2),
	3:   uint8(3),
	4:   uint8(4),
	5:   uint8(4),
	6:   uint8(5),
	7:   uint8(5),
	8:   uint8(6),
	9:   uint8(6),
	10:  uint8(6),
	11:  uint8(6),
	12:  uint8(7),
	13:  uint8(7),
	14:  uint8(7),
	15:  uint8(7),
	16:  uint8(8),
	17:  uint8(8),
	18:  uint8(8),
	19:  uint8(8),
	20:  uint8(8),
	21:  uint8(8),
	22:  uint8(8),
	23:  uint8(8),
	24:  uint8(9),
	25:  uint8(9),
	26:  uint8(9),
	27:  uint8(9),
	28:  uint8(9),
	29:  uint8(9),
	30:  uint8(9),
	31:  uint8(9),
	32:  uint8(10),
	33:  uint8(10),
	34:  uint8(10),
	35:  uint8(10),
	36:  uint8(10),
	37:  uint8(10),
	38:  uint8(10),
	39:  uint8(10),
	40:  uint8(10),
	41:  uint8(10),
	42:  uint8(10),
	43:  uint8(10),
	44:  uint8(10),
	45:  uint8(10),
	46:  uint8(10),
	47:  uint8(10),
	48:  uint8(11),
	49:  uint8(11),
	50:  uint8(11),
	51:  uint8(11),
	52:  uint8(11),
	53:  uint8(11),
	54:  uint8(11),
	55:  uint8(11),
	56:  uint8(11),
	57:  uint8(11),
	58:  uint8(11),
	59:  uint8(11),
	60:  uint8(11),
	61:  uint8(11),
	62:  uint8(11),
	63:  uint8(11),
	64:  uint8(12),
	65:  uint8(12),
	66:  uint8(12),
	67:  uint8(12),
	68:  uint8(12),
	69:  uint8(12),
	70:  uint8(12),
	71:  uint8(12),
	72:  uint8(12),
	73:  uint8(12),
	74:  uint8(12),
	75:  uint8(12),
	76:  uint8(12),
	77:  uint8(12),
	78:  uint8(12),
	79:  uint8(12),
	80:  uint8(12),
	81:  uint8(12),
	82:  uint8(12),
	83:  uint8(12),
	84:  uint8(12),
	85:  uint8(12),
	86:  uint8(12),
	87:  uint8(12),
	88:  uint8(12),
	89:  uint8(12),
	90:  uint8(12),
	91:  uint8(12),
	92:  uint8(12),
	93:  uint8(12),
	94:  uint8(12),
	95:  uint8(12),
	96:  uint8(13),
	97:  uint8(13),
	98:  uint8(13),
	99:  uint8(13),
	100: uint8(13),
	101: uint8(13),
	102: uint8(13),
	103: uint8(13),
	104: uint8(13),
	105: uint8(13),
	106: uint8(13),
	107: uint8(13),
	108: uint8(13),
	109: uint8(13),
	110: uint8(13),
	111: uint8(13),
	112: uint8(13),
	113: uint8(13),
	114: uint8(13),
	115: uint8(13),
	116: uint8(13),
	117: uint8(13),
	118: uint8(13),
	119: uint8(13),
	120: uint8(13),
	121: uint8(13),
	122: uint8(13),
	123: uint8(13),
	124: uint8(13),
	125: uint8(13),
	126: uint8(13),
	127: uint8(13),
	128: uint8(14),
	129: uint8(14),
	130: uint8(14),
	131: uint8(14),
	132: uint8(14),
	133: uint8(14),
	134: uint8(14),
	135: uint8(14),
	136: uint8(14),
	137: uint8(14),
	138: uint8(14),
	139: uint8(14),
	140: uint8(14),
	141: uint8(14),
	142: uint8(14),
	143: uint8(14),
	144: uint8(14),
	145: uint8(14),
	146: uint8(14),
	147: uint8(14),
	148: uint8(14),
	149: uint8(14),
	150: uint8(14),
	151: uint8(14),
	152: uint8(14),
	153: uint8(14),
	154: uint8(14),
	155: uint8(14),
	156: uint8(14),
	157: uint8(14),
	158: uint8(14),
	159: uint8(14),
	160: uint8(14),
	161: uint8(14),
	162: uint8(14),
	163: uint8(14),
	164: uint8(14),
	165: uint8(14),
	166: uint8(14),
	167: uint8(14),
	168: uint8(14),
	169: uint8(14),
	170: uint8(14),
	171: uint8(14),
	172: uint8(14),
	173: uint8(14),
	174: uint8(14),
	175: uint8(14),
	176: uint8(14),
	177: uint8(14),
	178: uint8(14),
	179: uint8(14),
	180: uint8(14),
	181: uint8(14),
	182: uint8(14),
	183: uint8(14),
	184: uint8(14),
	185: uint8(14),
	186: uint8(14),
	187: uint8(14),
	188: uint8(14),
	189: uint8(14),
	190: uint8(14),
	191: uint8(14),
	192: uint8(15),
	193: uint8(15),
	194: uint8(15),
	195: uint8(15),
	196: uint8(15),
	197: uint8(15),
	198: uint8(15),
	199: uint8(15),
	200: uint8(15),
	201: uint8(15),
	202: uint8(15),
	203: uint8(15),
	204: uint8(15),
	205: uint8(15),
	206: uint8(15),
	207: uint8(15),
	208: uint8(15),
	209: uint8(15),
	210: uint8(15),
	211: uint8(15),
	212: uint8(15),
	213: uint8(15),
	214: uint8(15),
	215: uint8(15),
	216: uint8(15),
	217: uint8(15),
	218: uint8(15),
	219: uint8(15),
	220: uint8(15),
	221: uint8(15),
	222: uint8(15),
	223: uint8(15),
	224: uint8(15),
	225: uint8(15),
	226: uint8(15),
	227: uint8(15),
	228: uint8(15),
	229: uint8(15),
	230: uint8(15),
	231: uint8(15),
	232: uint8(15),
	233: uint8(15),
	234: uint8(15),
	235: uint8(15),
	236: uint8(15),
	237: uint8(15),
	238: uint8(15),
	239: uint8(15),
	240: uint8(15),
	241: uint8(15),
	242: uint8(15),
	243: uint8(15),
	244: uint8(15),
	245: uint8(15),
	246: uint8(15),
	247: uint8(15),
	248: uint8(15),
	249: uint8(15),
	250: uint8(15),
	251: uint8(15),
	252: uint8(15),
	253: uint8(15),
	254: uint8(15),
	255: uint8(15),
	258: uint8(16),
	259: uint8(17),
	260: uint8(18),
	261: uint8(18),
	262: uint8(19),
	263: uint8(19),
	264: uint8(20),
	265: uint8(20),
	266: uint8(20),
	267: uint8(20),
	268: uint8(21),
	269: uint8(21),
	270: uint8(21),
	271: uint8(21),
	272: uint8(22),
	273: uint8(22),
	274: uint8(22),
	275: uint8(22),
	276: uint8(22),
	277: uint8(22),
	278: uint8(22),
	279: uint8(22),
	280: uint8(23),
	281: uint8(23),
	282: uint8(23),
	283: uint8(23),
	284: uint8(23),
	285: uint8(23),
	286: uint8(23),
	287: uint8(23),
	288: uint8(24),
	289: uint8(24),
	290: uint8(24),
	291: uint8(24),
	292: uint8(24),
	293: uint8(24),
	294: uint8(24),
	295: uint8(24),
	296: uint8(24),
	297: uint8(24),
	298: uint8(24),
	299: uint8(24),
	300: uint8(24),
	301: uint8(24),
	302: uint8(24),
	303: uint8(24),
	304: uint8(25),
	305: uint8(25),
	306: uint8(25),
	307: uint8(25),
	308: uint8(25),
	309: uint8(25),
	310: uint8(25),
	311: uint8(25),
	312: uint8(25),
	313: uint8(25),
	314: uint8(25),
	315: uint8(25),
	316: uint8(25),
	317: uint8(25),
	318: uint8(25),
	319: uint8(25),
	320: uint8(26),
	321: uint8(26),
	322: uint8(26),
	323: uint8(26),
	324: uint8(26),
	325: uint8(26),
	326: uint8(26),
	327: uint8(26),
	328: uint8(26),
	329: uint8(26),
	330: uint8(26),
	331: uint8(26),
	332: uint8(26),
	333: uint8(26),
	334: uint8(26),
	335: uint8(26),
	336: uint8(26),
	337: uint8(26),
	338: uint8(26),
	339: uint8(26),
	340: uint8(26),
	341: uint8(26),
	342: uint8(26),
	343: uint8(26),
	344: uint8(26),
	345: uint8(26),
	346: uint8(26),
	347: uint8(26),
	348: uint8(26),
	349: uint8(26),
	350: uint8(26),
	351: uint8(26),
	352: uint8(27),
	353: uint8(27),
	354: uint8(27),
	355: uint8(27),
	356: uint8(27),
	357: uint8(27),
	358: uint8(27),
	359: uint8(27),
	360: uint8(27),
	361: uint8(27),
	362: uint8(27),
	363: uint8(27),
	364: uint8(27),
	365: uint8(27),
	366: uint8(27),
	367: uint8(27),
	368: uint8(27),
	369: uint8(27),
	370: uint8(27),
	371: uint8(27),
	372: uint8(27),
	373: uint8(27),
	374: uint8(27),
	375: uint8(27),
	376: uint8(27),
	377: uint8(27),
	378: uint8(27),
	379: uint8(27),
	380: uint8(27),
	381: uint8(27),
	382: uint8(27),
	383: uint8(27),
	384: uint8(28),
	385: uint8(28),
	386: uint8(28),
	387: uint8(28),
	388: uint8(28),
	389: uint8(28),
	390: uint8(28),
	391: uint8(28),
	392: uint8(28),
	393: uint8(28),
	394: uint8(28),
	395: uint8(28),
	396: uint8(28),
	397: uint8(28),
	398: uint8(28),
	399: uint8(28),
	400: uint8(28),
	401: uint8(28),
	402: uint8(28),
	403: uint8(28),
	404: uint8(28),
	405: uint8(28),
	406: uint8(28),
	407: uint8(28),
	408: uint8(28),
	409: uint8(28),
	410: uint8(28),
	411: uint8(28),
	412: uint8(28),
	413: uint8(28),
	414: uint8(28),
	415: uint8(28),
	416: uint8(28),
	417: uint8(28),
	418: uint8(28),
	419: uint8(28),
	420: uint8(28),
	421: uint8(28),
	422: uint8(28),
	423: uint8(28),
	424: uint8(28),
	425: uint8(28),
	426: uint8(28),
	427: uint8(28),
	428: uint8(28),
	429: uint8(28),
	430: uint8(28),
	431: uint8(28),
	432: uint8(28),
	433: uint8(28),
	434: uint8(28),
	435: uint8(28),
	436: uint8(28),
	437: uint8(28),
	438: uint8(28),
	439: uint8(28),
	440: uint8(28),
	441: uint8(28),
	442: uint8(28),
	443: uint8(28),
	444: uint8(28),
	445: uint8(28),
	446: uint8(28),
	447: uint8(28),
	448: uint8(29),
	449: uint8(29),
	450: uint8(29),
	451: uint8(29),
	452: uint8(29),
	453: uint8(29),
	454: uint8(29),
	455: uint8(29),
	456: uint8(29),
	457: uint8(29),
	458: uint8(29),
	459: uint8(29),
	460: uint8(29),
	461: uint8(29),
	462: uint8(29),
	463: uint8(29),
	464: uint8(29),
	465: uint8(29),
	466: uint8(29),
	467: uint8(29),
	468: uint8(29),
	469: uint8(29),
	470: uint8(29),
	471: uint8(29),
	472: uint8(29),
	473: uint8(29),
	474: uint8(29),
	475: uint8(29),
	476: uint8(29),
	477: uint8(29),
	478: uint8(29),
	479: uint8(29),
	480: uint8(29),
	481: uint8(29),
	482: uint8(29),
	483: uint8(29),
	484: uint8(29),
	485: uint8(29),
	486: uint8(29),
	487: uint8(29),
	488: uint8(29),
	489: uint8(29),
	490: uint8(29),
	491: uint8(29),
	492: uint8(29),
	493: uint8(29),
	494: uint8(29),
	495: uint8(29),
	496: uint8(29),
	497: uint8(29),
	498: uint8(29),
	499: uint8(29),
	500: uint8(29),
	501: uint8(29),
	502: uint8(29),
	503: uint8(29),
	504: uint8(29),
	505: uint8(29),
	506: uint8(29),
	507: uint8(29),
	508: uint8(29),
	509: uint8(29),
	510: uint8(29),
	511: uint8(29),
}

var x__length_code = [256]Tuch{
	1:   uint8(1),
	2:   uint8(2),
	3:   uint8(3),
	4:   uint8(4),
	5:   uint8(5),
	6:   uint8(6),
	7:   uint8(7),
	8:   uint8(8),
	9:   uint8(8),
	10:  uint8(9),
	11:  uint8(9),
	12:  uint8(10),
	13:  uint8(10),
	14:  uint8(11),
	15:  uint8(11),
	16:  uint8(12),
	17:  uint8(12),
	18:  uint8(12),
	19:  uint8(12),
	20:  uint8(13),
	21:  uint8(13),
	22:  uint8(13),
	23:  uint8(13),
	24:  uint8(14),
	25:  uint8(14),
	26:  uint8(14),
	27:  uint8(14),
	28:  uint8(15),
	29:  uint8(15),
	30:  uint8(15),
	31:  uint8(15),
	32:  uint8(16),
	33:  uint8(16),
	34:  uint8(16),
	35:  uint8(16),
	36:  uint8(16),
	37:  uint8(16),
	38:  uint8(16),
	39:  uint8(16),
	40:  uint8(17),
	41:  uint8(17),
	42:  uint8(17),
	43:  uint8(17),
	44:  uint8(17),
	45:  uint8(17),
	46:  uint8(17),
	47:  uint8(17),
	48:  uint8(18),
	49:  uint8(18),
	50:  uint8(18),
	51:  uint8(18),
	52:  uint8(18),
	53:  uint8(18),
	54:  uint8(18),
	55:  uint8(18),
	56:  uint8(19),
	57:  uint8(19),
	58:  uint8(19),
	59:  uint8(19),
	60:  uint8(19),
	61:  uint8(19),
	62:  uint8(19),
	63:  uint8(19),
	64:  uint8(20),
	65:  uint8(20),
	66:  uint8(20),
	67:  uint8(20),
	68:  uint8(20),
	69:  uint8(20),
	70:  uint8(20),
	71:  uint8(20),
	72:  uint8(20),
	73:  uint8(20),
	74:  uint8(20),
	75:  uint8(20),
	76:  uint8(20),
	77:  uint8(20),
	78:  uint8(20),
	79:  uint8(20),
	80:  uint8(21),
	81:  uint8(21),
	82:  uint8(21),
	83:  uint8(21),
	84:  uint8(21),
	85:  uint8(21),
	86:  uint8(21),
	87:  uint8(21),
	88:  uint8(21),
	89:  uint8(21),
	90:  uint8(21),
	91:  uint8(21),
	92:  uint8(21),
	93:  uint8(21),
	94:  uint8(21),
	95:  uint8(21),
	96:  uint8(22),
	97:  uint8(22),
	98:  uint8(22),
	99:  uint8(22),
	100: uint8(22),
	101: uint8(22),
	102: uint8(22),
	103: uint8(22),
	104: uint8(22),
	105: uint8(22),
	106: uint8(22),
	107: uint8(22),
	108: uint8(22),
	109: uint8(22),
	110: uint8(22),
	111: uint8(22),
	112: uint8(23),
	113: uint8(23),
	114: uint8(23),
	115: uint8(23),
	116: uint8(23),
	117: uint8(23),
	118: uint8(23),
	119: uint8(23),
	120: uint8(23),
	121: uint8(23),
	122: uint8(23),
	123: uint8(23),
	124: uint8(23),
	125: uint8(23),
	126: uint8(23),
	127: uint8(23),
	128: uint8(24),
	129: uint8(24),
	130: uint8(24),
	131: uint8(24),
	132: uint8(24),
	133: uint8(24),
	134: uint8(24),
	135: uint8(24),
	136: uint8(24),
	137: uint8(24),
	138: uint8(24),
	139: uint8(24),
	140: uint8(24),
	141: uint8(24),
	142: uint8(24),
	143: uint8(24),
	144: uint8(24),
	145: uint8(24),
	146: uint8(24),
	147: uint8(24),
	148: uint8(24),
	149: uint8(24),
	150: uint8(24),
	151: uint8(24),
	152: uint8(24),
	153: uint8(24),
	154: uint8(24),
	155: uint8(24),
	156: uint8(24),
	157: uint8(24),
	158: uint8(24),
	159: uint8(24),
	160: uint8(25),
	161: uint8(25),
	162: uint8(25),
	163: uint8(25),
	164: uint8(25),
	165: uint8(25),
	166: uint8(25),
	167: uint8(25),
	168: uint8(25),
	169: uint8(25),
	170: uint8(25),
	171: uint8(25),
	172: uint8(25),
	173: uint8(25),
	174: uint8(25),
	175: uint8(25),
	176: uint8(25),
	177: uint8(25),
	178: uint8(25),
	179: uint8(25),
	180: uint8(25),
	181: uint8(25),
	182: uint8(25),
	183: uint8(25),
	184: uint8(25),
	185: uint8(25),
	186: uint8(25),
	187: uint8(25),
	188: uint8(25),
	189: uint8(25),
	190: uint8(25),
	191: uint8(25),
	192: uint8(26),
	193: uint8(26),
	194: uint8(26),
	195: uint8(26),
	196: uint8(26),
	197: uint8(26),
	198: uint8(26),
	199: uint8(26),
	200: uint8(26),
	201: uint8(26),
	202: uint8(26),
	203: uint8(26),
	204: uint8(26),
	205: uint8(26),
	206: uint8(26),
	207: uint8(26),
	208: uint8(26),
	209: uint8(26),
	210: uint8(26),
	211: uint8(26),
	212: uint8(26),
	213: uint8(26),
	214: uint8(26),
	215: uint8(26),
	216: uint8(26),
	217: uint8(26),
	218: uint8(26),
	219: uint8(26),
	220: uint8(26),
	221: uint8(26),
	222: uint8(26),
	223: uint8(26),
	224: uint8(27),
	225: uint8(27),
	226: uint8(27),
	227: uint8(27),
	228: uint8(27),
	229: uint8(27),
	230: uint8(27),
	231: uint8(27),
	232: uint8(27),
	233: uint8(27),
	234: uint8(27),
	235: uint8(27),
	236: uint8(27),
	237: uint8(27),
	238: uint8(27),
	239: uint8(27),
	240: uint8(27),
	241: uint8(27),
	242: uint8(27),
	243: uint8(27),
	244: uint8(27),
	245: uint8(27),
	246: uint8(27),
	247: uint8(27),
	248: uint8(27),
	249: uint8(27),
	250: uint8(27),
	251: uint8(27),
	252: uint8(27),
	253: uint8(27),
	254: uint8(27),
	255: uint8(28),
}

var Xdeflate_copyright = [68]uint8{' ', 'd', 'e', 'f', 'l', 'a', 't', 'e', ' ', '1', '.', '3', '.', '1', ' ', 'C', 'o', 'p', 'y', 'r', 'i', 'g', 'h', 't', ' ', '1', '9', '9', '5', '-', '2', '0', '2', '4', ' ', 'J', 'e', 'a', 'n', '-', 'l', 'o', 'u', 'p', ' ', 'G', 'a', 'i', 'l', 'l', 'y', ' ', 'a', 'n', 'd', ' ', 'M', 'a', 'r', 'k', ' ', 'A', 'd', 'l', 'e', 'r', ' '}

var Xinflate_copyright = [47]uint8{' ', 'i', 'n', 'f', 'l', 'a', 't', 'e', ' ', '1', '.', '3', '.', '1', ' ', 'C', 'o', 'p', 'y', 'r', 'i', 'g', 'h', 't', ' ', '1', '9', '9', '5', '-', '2', '0', '2', '4', ' ', 'M', 'a', 'r', 'k', ' ', 'A', 'd', 'l', 'e', 'r', ' '}

var Xz_errmsg = [10]uintptr{
	0: __ccgo_ts + 467,
	1: __ccgo_ts + 483,
	2: __ccgo_ts + 494,
	3: __ccgo_ts + 495,
	4: __ccgo_ts + 506,
	5: __ccgo_ts + 519,
	6: __ccgo_ts + 530,
	7: __ccgo_ts + 550,
	8: __ccgo_ts + 563,
	9: __ccgo_ts + 494,
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "1.3.1\x00invalid block type\x00invalid stored block lengths\x00too many length or distance symbols\x00invalid code lengths set\x00invalid bit length repeat\x00invalid code -- missing end-of-block\x00invalid literal/lengths set\x00invalid distances set\x00invalid literal/length code\x00invalid distance code\x00invalid distance too far back\x00incorrect header check\x00unknown compression method\x00invalid window size\x00unknown header flags set\x00header crc mismatch\x00incorrect data check\x00incorrect length check\x00need dictionary\x00stream end\x00\x00file error\x00stream error\x00data error\x00insufficient memory\x00buffer error\x00incompatible version\x00%s\x00<fd:%d>\x00out of memory\x00%s%s%s\x00: \x00unexpected end of file\x00internal error: inflate stream corrupt\x00compressed data error\x00request does not fit in an int\x00request does not fit in a size_t\x00out of room to push characters\x00internal error: deflate stream corrupt\x00requested length does not fit in int\x00string length does not fit in int\x00"
