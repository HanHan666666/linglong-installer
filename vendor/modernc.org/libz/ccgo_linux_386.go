// Code generated for linux/386 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors --libc modernc.org/libc --package-name=libz -o libz.a.go libz.a', DO NOT EDIT.

//go:build linux && 386

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
const m_LONG_BIT = 32
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
const m_PAGESIZE = 4096
const m_PAGE_SIZE = "PAGESIZE"
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
const m__FILE_OFFSET_BITS = 64
const m__GNU_SOURCE = 1
const m__ILP32 = 1
const m__LARGEFILE64_SOURCE = 1
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
const m__POSIX_V6_ILP32_OFFBIG = 1
const m__POSIX_V7_ILP32_OFFBIG = 1
const m__POSIX_VDISABLE = 0
const m__POSIX_VERSION = 200809
const m__REDIR_TIME64 = 1
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
const m___ATOMIC_ACQUIRE = 2
const m___ATOMIC_ACQ_REL = 4
const m___ATOMIC_CONSUME = 1
const m___ATOMIC_HLE_ACQUIRE = 65536
const m___ATOMIC_HLE_RELEASE = 131072
const m___ATOMIC_RELAXED = 0
const m___ATOMIC_RELEASE = 3
const m___ATOMIC_SEQ_CST = 5
const m___BIGGEST_ALIGNMENT__ = 16
const m___BIG_ENDIAN = 4321
const m___BYTE_ORDER = 1234
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___CCGO__ = 1
const m___CHAR_BIT__ = 8
const m___DBL_DECIMAL_DIG__ = 17
const m___DBL_DIG__ = 15
const m___DBL_HAS_DENORM__ = 1
const m___DBL_HAS_INFINITY__ = 1
const m___DBL_HAS_QUIET_NAN__ = 1
const m___DBL_IS_IEC_60559__ = 2
const m___DBL_MANT_DIG__ = 53
const m___DBL_MAX_10_EXP__ = 308
const m___DBL_MAX_EXP__ = 1024
const m___DEC128_EPSILON__ = 1e-33
const m___DEC128_MANT_DIG__ = 34
const m___DEC128_MAX_EXP__ = 6145
const m___DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const m___DEC128_MIN__ = 1e-6143
const m___DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const m___DEC32_EPSILON__ = 1e-6
const m___DEC32_MANT_DIG__ = 7
const m___DEC32_MAX_EXP__ = 97
const m___DEC32_MAX__ = 9.999999e96
const m___DEC32_MIN__ = 1e-95
const m___DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const m___DEC64_EPSILON__ = 1e-15
const m___DEC64_MANT_DIG__ = 16
const m___DEC64_MAX_EXP__ = 385
const m___DEC64_MAX__ = "9.999999999999999E384"
const m___DEC64_MIN__ = 1e-383
const m___DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const m___DECIMAL_BID_FORMAT__ = 1
const m___DECIMAL_DIG__ = 17
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
const m___FLT_EVAL_METHOD_TS_18661_3__ = 2
const m___FLT_EVAL_METHOD__ = 2
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
const m___ILP32__ = 1
const m___INT16_MAX__ = 0x7fff
const m___INT32_MAX__ = 0x7fffffff
const m___INT32_TYPE__ = "int"
const m___INT64_MAX__ = 0x7fffffffffffffff
const m___INT8_MAX__ = 0x7f
const m___INTMAX_MAX__ = 0x7fffffffffffffff
const m___INTMAX_WIDTH__ = 64
const m___INTPTR_MAX__ = 0x7fffffff
const m___INTPTR_TYPE__ = "int"
const m___INTPTR_WIDTH__ = 32
const m___INT_FAST16_MAX__ = 0x7fffffff
const m___INT_FAST16_TYPE__ = "int"
const m___INT_FAST16_WIDTH__ = 32
const m___INT_FAST32_MAX__ = 0x7fffffff
const m___INT_FAST32_TYPE__ = "int"
const m___INT_FAST32_WIDTH__ = 32
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
const m___LAHF_SAHF__ = 1
const m___LDBL_DECIMAL_DIG__ = 17
const m___LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___LDBL_DIG__ = 15
const m___LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___LDBL_HAS_DENORM__ = 1
const m___LDBL_HAS_INFINITY__ = 1
const m___LDBL_HAS_QUIET_NAN__ = 1
const m___LDBL_IS_IEC_60559__ = 2
const m___LDBL_MANT_DIG__ = 53
const m___LDBL_MAX_10_EXP__ = 308
const m___LDBL_MAX_EXP__ = 1024
const m___LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const m___LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const m___LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___LITTLE_ENDIAN = 1234
const m___LONG_DOUBLE_64__ = 1
const m___LONG_LONG_MAX__ = 0x7fffffffffffffff
const m___LONG_LONG_WIDTH__ = 64
const m___LONG_MAX = 0x7fffffff
const m___LONG_MAX__ = 0x7fffffff
const m___LONG_WIDTH__ = 32
const m___NO_INLINE__ = 1
const m___ORDER_BIG_ENDIAN__ = 4321
const m___ORDER_LITTLE_ENDIAN__ = 1234
const m___ORDER_PDP_ENDIAN__ = 3412
const m___PDP_ENDIAN = 3412
const m___PIC__ = 2
const m___PIE__ = 2
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTRDIFF_MAX__ = 0x7fffffff
const m___PTRDIFF_TYPE__ = "int"
const m___PTRDIFF_WIDTH__ = 32
const m___SCHAR_MAX__ = 0x7f
const m___SCHAR_WIDTH__ = 8
const m___SEG_FS = 1
const m___SEG_GS = 1
const m___SHRT_MAX__ = 0x7fff
const m___SHRT_WIDTH__ = 16
const m___SIG_ATOMIC_MAX__ = 0x7fffffff
const m___SIG_ATOMIC_TYPE__ = "int"
const m___SIG_ATOMIC_WIDTH__ = 32
const m___SIZEOF_DOUBLE__ = 8
const m___SIZEOF_FLOAT128__ = 16
const m___SIZEOF_FLOAT80__ = 12
const m___SIZEOF_FLOAT__ = 4
const m___SIZEOF_INT__ = 4
const m___SIZEOF_LONG_DOUBLE__ = 8
const m___SIZEOF_LONG_LONG__ = 8
const m___SIZEOF_LONG__ = 4
const m___SIZEOF_POINTER__ = 4
const m___SIZEOF_PTRDIFF_T__ = 4
const m___SIZEOF_SHORT__ = 2
const m___SIZEOF_SIZE_T__ = 4
const m___SIZEOF_WCHAR_T__ = 4
const m___SIZEOF_WINT_T__ = 4
const m___SIZE_MAX__ = 0xffffffff
const m___SIZE_WIDTH__ = 32
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
const m___UINT64_MAX__ = "0xffffffffffffffffU"
const m___UINT8_MAX__ = 0xff
const m___UINTMAX_MAX__ = "0xffffffffffffffffU"
const m___UINTPTR_MAX__ = 0xffffffff
const m___UINT_FAST16_MAX__ = 0xffffffff
const m___UINT_FAST32_MAX__ = 0xffffffff
const m___UINT_FAST64_MAX__ = "0xffffffffffffffffU"
const m___UINT_FAST8_MAX__ = 0xff
const m___UINT_LEAST16_MAX__ = 0xffff
const m___UINT_LEAST32_MAX__ = 0xffffffff
const m___UINT_LEAST64_MAX__ = "0xffffffffffffffffU"
const m___UINT_LEAST8_MAX__ = 0xff
const m___USE_TIME_BITS64 = 1
const m___VERSION__ = "12.2.0"
const m___WCHAR_MAX__ = 0x7fffffff
const m___WCHAR_WIDTH__ = 32
const m___WINT_MAX__ = 0xffffffff
const m___WINT_MIN__ = 0
const m___WINT_WIDTH__ = 32
const m___code_model_32__ = 1
const m___gnu_linux__ = 1
const m___i386 = 1
const m___i386__ = 1
const m___i686 = 1
const m___i686__ = 1
const m___inline = "inline"
const m___linux = 1
const m___linux__ = 1
const m___pentiumpro = 1
const m___pentiumpro__ = 1
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
const m_i386 = 1
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

type t__predefined_size_t = uint32

type t__predefined_wchar_t = int32

type t__predefined_ptrdiff_t = int32

type Twchar_t = int32

type Tmax_align_t = struct {
	F__ll int64
	F__ld float64
}

type Tsize_t = uint32

type Tptrdiff_t = int32

type Tz_size_t = uint32

type TByte = uint8

type TuInt = uint32

type TuLong = uint32

type TBytef = uint8

type Tcharf = int8

type Tintf = int32

type TuIntf = uint32

type TuLongf = uint32

type Tvoidpc = uintptr

type Tvoidpf = uintptr

type Tvoidp = uintptr

type Tz_crc_t = uint32

type Tssize_t = int32

type Tregister_t = int32

type Ttime_t = int64

type Tsuseconds_t = int64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tu_int64_t = uint64

type Tmode_t = uint32

type Tnlink_t = uint32

type Toff_t = int64

type Tino_t = uint64

type Tdev_t = uint64

type Tblksize_t = int32

type Tblkcnt_t = int64

type Tfsblkcnt_t = uint64

type Tfsfilcnt_t = uint64

type Ttimer_t = uintptr

type Tclockid_t = int32

type Tclock_t = int32

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
		F__vi [0][9]int32
		F__s  [0][9]uint32
		F__i  [9]int32
	}
}

type Tpthread_mutex_t = struct {
	F__u struct {
		F__vi [0][6]int32
		F__p  [0][6]uintptr
		F__i  [6]int32
	}
	F__ccgo_room int32
}

type Tpthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][12]uintptr
		F__i  [12]int32
	}
}

type Tpthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][8]uintptr
		F__i  [8]int32
	}
}

type Tpthread_barrier_t = struct {
	F__u struct {
		F__vi [0][5]int32
		F__p  [0][5]uintptr
		F__i  [5]int32
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

type Tu_long = uint32

type Tulong = uint32

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
	Ftv_sec   Ttime_t
	Ftv_nsec  int32
	F__ccgo12 uint32
}

type Tsigset_t = struct {
	F__bits [32]uint32
}

type t__sigset_t = Tsigset_t

type Tfd_mask = uint32

type Tfd_set = struct {
	Ffds_bits [32]uint32
}

type Tva_list = uintptr

type Tintptr_t = int32

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
	Fquot int32
	Frem  int32
}

type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

type Tuch = uint8

type Tuchf = uint8

type Tush = uint16

type Tushf = uint16

type Tulg = uint32

/* Reverse the bytes in a 32-bit value */

/* NMAX is the largest n such that 255n(n+1)/2 + (n+1)(BASE-1) <= 2^32-1 */

/* use NO_DIVIDE if your processor does not do division in hardware --
   try it both ways to see which is faster */

// C documentation
//
//	/* ========================================================================= */
func Xadler32_z(tls *libc.TLS, adler TuLong, buf uintptr, len1 Tz_size_t) (r TuLong) {
	var n, sum2, v3 uint32
	var v1, v5 Tz_size_t
	var v2, v6 uintptr
	_, _, _, _, _, _, _ = n, sum2, v1, v2, v3, v5, v6
	/* split Adler-32 into component sums */
	sum2 = adler >> libc.Int32FromInt32(16) & uint32(0xffff)
	adler &= uint32(0xffff)
	/* in case user likes doing a byte at a time, keep it fast */
	if len1 == uint32(1) {
		adler += uint32(*(*TBytef)(unsafe.Pointer(buf)))
		if adler >= uint32(65521) {
			adler -= uint32(65521)
		}
		sum2 += adler
		if sum2 >= uint32(65521) {
			sum2 -= uint32(65521)
		}
		return adler | sum2<<int32(16)
	}
	/* initial Adler-32 value (deferred check for len == 1 speed) */
	if buf == uintptr(m_Z_NULL) {
		return uint32(1)
	}
	/* in case short lengths are provided, keep it somewhat fast */
	if len1 < uint32(16) {
		for {
			v1 = len1
			len1--
			if !(v1 != 0) {
				break
			}
			v2 = buf
			buf++
			adler += uint32(*(*TBytef)(unsafe.Pointer(v2)))
			sum2 += adler
		}
		if adler >= uint32(65521) {
			adler -= uint32(65521)
		}
		sum2 %= uint32(65521) /* only added so many BASE's */
		return adler | sum2<<int32(16)
	}
	/* do length NMAX blocks -- requires just one modulo operation */
	for len1 >= uint32(m_NMAX) {
		len1 -= uint32(m_NMAX)
		n = libc.Uint32FromInt32(libc.Int32FromInt32(m_NMAX) / libc.Int32FromInt32(16)) /* NMAX is divisible by 16 */
		for {
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + 8)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
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
		adler %= uint32(65521)
		sum2 %= uint32(65521)
	}
	/* do remaining bytes (less than NMAX, still just one modulo) */
	if len1 != 0 { /* avoid modulos if none remaining */
		for len1 >= uint32(16) {
			len1 -= uint32(16)
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + 8)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
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
			adler += uint32(*(*TBytef)(unsafe.Pointer(v6)))
			sum2 += adler
		}
		adler %= uint32(65521)
		sum2 %= uint32(65521)
	}
	/* return recombined sums */
	return adler | sum2<<int32(16)
}

// C documentation
//
//	/* ========================================================================= */
func Xadler32(tls *libc.TLS, adler TuLong, buf uintptr, len1 TuInt) (r TuLong) {
	return Xadler32_z(tls, adler, buf, len1)
}

// C documentation
//
//	/* ========================================================================= */
func _adler32_combine_(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	var rem, sum1, sum2 uint32
	_, _, _ = rem, sum1, sum2
	/* for negative len, return invalid adler32 as a clue for debugging */
	if len2 < 0 {
		return uint32(0xffffffff)
	}
	/* the derivation of this formula is left as an exercise for the reader */
	len2 %= libc.Int64FromUint32(65521) /* assumes len2 >= 0 */
	rem = libc.Uint32FromInt64(len2)
	sum1 = adler1 & uint32(0xffff)
	sum2 = rem * sum1
	sum2 %= uint32(65521)
	sum1 += adler2&uint32(0xffff) + uint32(65521) - uint32(1)
	sum2 += adler1>>libc.Int32FromInt32(16)&uint32(0xffff) + adler2>>libc.Int32FromInt32(16)&uint32(0xffff) + uint32(65521) - rem
	if sum1 >= uint32(65521) {
		sum1 -= uint32(65521)
	}
	if sum1 >= uint32(65521) {
		sum1 -= uint32(65521)
	}
	if sum2 >= libc.Uint32FromUint32(65521)<<libc.Int32FromInt32(1) {
		sum2 -= libc.Uint32FromUint32(65521) << libc.Int32FromInt32(1)
	}
	if sum2 >= uint32(65521) {
		sum2 -= uint32(65521)
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
type Tz_word_t = uint32

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
	return word&uint32(0xff000000)>>int32(24) | word&uint32(0xff0000)>>int32(8) | word&uint32(0xff00)<<int32(8) | word&uint32(0xff)<<int32(24)
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
	1:   uint32(0x96300777),
	2:   uint32(0x2c610eee),
	3:   uint32(0xba510999),
	4:   uint32(0x19c46d07),
	5:   uint32(0x8ff46a70),
	6:   uint32(0x35a563e9),
	7:   uint32(0xa395649e),
	8:   uint32(0x3288db0e),
	9:   uint32(0xa4b8dc79),
	10:  uint32(0x1ee9d5e0),
	11:  uint32(0x88d9d297),
	12:  uint32(0x2b4cb609),
	13:  uint32(0xbd7cb17e),
	14:  uint32(0x072db8e7),
	15:  uint32(0x911dbf90),
	16:  uint32(0x6410b71d),
	17:  uint32(0xf220b06a),
	18:  uint32(0x4871b9f3),
	19:  uint32(0xde41be84),
	20:  uint32(0x7dd4da1a),
	21:  uint32(0xebe4dd6d),
	22:  uint32(0x51b5d4f4),
	23:  uint32(0xc785d383),
	24:  uint32(0x56986c13),
	25:  uint32(0xc0a86b64),
	26:  uint32(0x7af962fd),
	27:  uint32(0xecc9658a),
	28:  uint32(0x4f5c0114),
	29:  uint32(0xd96c0663),
	30:  uint32(0x633d0ffa),
	31:  uint32(0xf50d088d),
	32:  uint32(0xc8206e3b),
	33:  uint32(0x5e10694c),
	34:  uint32(0xe44160d5),
	35:  uint32(0x727167a2),
	36:  uint32(0xd1e4033c),
	37:  uint32(0x47d4044b),
	38:  uint32(0xfd850dd2),
	39:  uint32(0x6bb50aa5),
	40:  uint32(0xfaa8b535),
	41:  uint32(0x6c98b242),
	42:  uint32(0xd6c9bbdb),
	43:  uint32(0x40f9bcac),
	44:  uint32(0xe36cd832),
	45:  uint32(0x755cdf45),
	46:  uint32(0xcf0dd6dc),
	47:  uint32(0x593dd1ab),
	48:  uint32(0xac30d926),
	49:  uint32(0x3a00de51),
	50:  uint32(0x8051d7c8),
	51:  uint32(0x1661d0bf),
	52:  uint32(0xb5f4b421),
	53:  uint32(0x23c4b356),
	54:  uint32(0x9995bacf),
	55:  uint32(0x0fa5bdb8),
	56:  uint32(0x9eb80228),
	57:  uint32(0x0888055f),
	58:  uint32(0xb2d90cc6),
	59:  uint32(0x24e90bb1),
	60:  uint32(0x877c6f2f),
	61:  uint32(0x114c6858),
	62:  uint32(0xab1d61c1),
	63:  uint32(0x3d2d66b6),
	64:  uint32(0x9041dc76),
	65:  uint32(0x0671db01),
	66:  uint32(0xbc20d298),
	67:  uint32(0x2a10d5ef),
	68:  uint32(0x8985b171),
	69:  uint32(0x1fb5b606),
	70:  uint32(0xa5e4bf9f),
	71:  uint32(0x33d4b8e8),
	72:  uint32(0xa2c90778),
	73:  uint32(0x34f9000f),
	74:  uint32(0x8ea80996),
	75:  uint32(0x18980ee1),
	76:  uint32(0xbb0d6a7f),
	77:  uint32(0x2d3d6d08),
	78:  uint32(0x976c6491),
	79:  uint32(0x015c63e6),
	80:  uint32(0xf4516b6b),
	81:  uint32(0x62616c1c),
	82:  uint32(0xd8306585),
	83:  uint32(0x4e0062f2),
	84:  uint32(0xed95066c),
	85:  uint32(0x7ba5011b),
	86:  uint32(0xc1f40882),
	87:  uint32(0x57c40ff5),
	88:  uint32(0xc6d9b065),
	89:  uint32(0x50e9b712),
	90:  uint32(0xeab8be8b),
	91:  uint32(0x7c88b9fc),
	92:  uint32(0xdf1ddd62),
	93:  uint32(0x492dda15),
	94:  uint32(0xf37cd38c),
	95:  uint32(0x654cd4fb),
	96:  uint32(0x5861b24d),
	97:  uint32(0xce51b53a),
	98:  uint32(0x7400bca3),
	99:  uint32(0xe230bbd4),
	100: uint32(0x41a5df4a),
	101: uint32(0xd795d83d),
	102: uint32(0x6dc4d1a4),
	103: uint32(0xfbf4d6d3),
	104: uint32(0x6ae96943),
	105: uint32(0xfcd96e34),
	106: uint32(0x468867ad),
	107: uint32(0xd0b860da),
	108: uint32(0x732d0444),
	109: uint32(0xe51d0333),
	110: uint32(0x5f4c0aaa),
	111: uint32(0xc97c0ddd),
	112: uint32(0x3c710550),
	113: uint32(0xaa410227),
	114: uint32(0x10100bbe),
	115: uint32(0x86200cc9),
	116: uint32(0x25b56857),
	117: uint32(0xb3856f20),
	118: uint32(0x09d466b9),
	119: uint32(0x9fe461ce),
	120: uint32(0x0ef9de5e),
	121: uint32(0x98c9d929),
	122: uint32(0x2298d0b0),
	123: uint32(0xb4a8d7c7),
	124: uint32(0x173db359),
	125: uint32(0x810db42e),
	126: uint32(0x3b5cbdb7),
	127: uint32(0xad6cbac0),
	128: uint32(0x2083b8ed),
	129: uint32(0xb6b3bf9a),
	130: uint32(0x0ce2b603),
	131: uint32(0x9ad2b174),
	132: uint32(0x3947d5ea),
	133: uint32(0xaf77d29d),
	134: uint32(0x1526db04),
	135: uint32(0x8316dc73),
	136: uint32(0x120b63e3),
	137: uint32(0x843b6494),
	138: uint32(0x3e6a6d0d),
	139: uint32(0xa85a6a7a),
	140: uint32(0x0bcf0ee4),
	141: uint32(0x9dff0993),
	142: uint32(0x27ae000a),
	143: uint32(0xb19e077d),
	144: uint32(0x44930ff0),
	145: uint32(0xd2a30887),
	146: uint32(0x68f2011e),
	147: uint32(0xfec20669),
	148: uint32(0x5d5762f7),
	149: uint32(0xcb676580),
	150: uint32(0x71366c19),
	151: uint32(0xe7066b6e),
	152: uint32(0x761bd4fe),
	153: uint32(0xe02bd389),
	154: uint32(0x5a7ada10),
	155: uint32(0xcc4add67),
	156: uint32(0x6fdfb9f9),
	157: uint32(0xf9efbe8e),
	158: uint32(0x43beb717),
	159: uint32(0xd58eb060),
	160: uint32(0xe8a3d6d6),
	161: uint32(0x7e93d1a1),
	162: uint32(0xc4c2d838),
	163: uint32(0x52f2df4f),
	164: uint32(0xf167bbd1),
	165: uint32(0x6757bca6),
	166: uint32(0xdd06b53f),
	167: uint32(0x4b36b248),
	168: uint32(0xda2b0dd8),
	169: uint32(0x4c1b0aaf),
	170: uint32(0xf64a0336),
	171: uint32(0x607a0441),
	172: uint32(0xc3ef60df),
	173: uint32(0x55df67a8),
	174: uint32(0xef8e6e31),
	175: uint32(0x79be6946),
	176: uint32(0x8cb361cb),
	177: uint32(0x1a8366bc),
	178: uint32(0xa0d26f25),
	179: uint32(0x36e26852),
	180: uint32(0x95770ccc),
	181: uint32(0x03470bbb),
	182: uint32(0xb9160222),
	183: uint32(0x2f260555),
	184: uint32(0xbe3bbac5),
	185: uint32(0x280bbdb2),
	186: uint32(0x925ab42b),
	187: uint32(0x046ab35c),
	188: uint32(0xa7ffd7c2),
	189: uint32(0x31cfd0b5),
	190: uint32(0x8b9ed92c),
	191: uint32(0x1daede5b),
	192: uint32(0xb0c2649b),
	193: uint32(0x26f263ec),
	194: uint32(0x9ca36a75),
	195: uint32(0x0a936d02),
	196: uint32(0xa906099c),
	197: uint32(0x3f360eeb),
	198: uint32(0x85670772),
	199: uint32(0x13570005),
	200: uint32(0x824abf95),
	201: uint32(0x147ab8e2),
	202: uint32(0xae2bb17b),
	203: uint32(0x381bb60c),
	204: uint32(0x9b8ed292),
	205: uint32(0x0dbed5e5),
	206: uint32(0xb7efdc7c),
	207: uint32(0x21dfdb0b),
	208: uint32(0xd4d2d386),
	209: uint32(0x42e2d4f1),
	210: uint32(0xf8b3dd68),
	211: uint32(0x6e83da1f),
	212: uint32(0xcd16be81),
	213: uint32(0x5b26b9f6),
	214: uint32(0xe177b06f),
	215: uint32(0x7747b718),
	216: uint32(0xe65a0888),
	217: uint32(0x706a0fff),
	218: uint32(0xca3b0666),
	219: uint32(0x5c0b0111),
	220: uint32(0xff9e658f),
	221: uint32(0x69ae62f8),
	222: uint32(0xd3ff6b61),
	223: uint32(0x45cf6c16),
	224: uint32(0x78e20aa0),
	225: uint32(0xeed20dd7),
	226: uint32(0x5483044e),
	227: uint32(0xc2b30339),
	228: uint32(0x612667a7),
	229: uint32(0xf71660d0),
	230: uint32(0x4d476949),
	231: uint32(0xdb776e3e),
	232: uint32(0x4a6ad1ae),
	233: uint32(0xdc5ad6d9),
	234: uint32(0x660bdf40),
	235: uint32(0xf03bd837),
	236: uint32(0x53aebca9),
	237: uint32(0xc59ebbde),
	238: uint32(0x7fcfb247),
	239: uint32(0xe9ffb530),
	240: uint32(0x1cf2bdbd),
	241: uint32(0x8ac2baca),
	242: uint32(0x3093b353),
	243: uint32(0xa6a3b424),
	244: uint32(0x0536d0ba),
	245: uint32(0x9306d7cd),
	246: uint32(0x2957de54),
	247: uint32(0xbf67d923),
	248: uint32(0x2e7a66b3),
	249: uint32(0xb84a61c4),
	250: uint32(0x021b685d),
	251: uint32(0x942b6f2a),
	252: uint32(0x37be0bb4),
	253: uint32(0xa18e0cc3),
	254: uint32(0x1bdf055a),
	255: uint32(0x8def022d),
}
var _crc_braid_table = [4][256]Tz_crc_t{
	0: {
		1:   uint32(0x65673b46),
		2:   uint32(0xcace768c),
		3:   uint32(0xafa94dca),
		4:   uint32(0x4eedeb59),
		5:   uint32(0x2b8ad01f),
		6:   uint32(0x84239dd5),
		7:   uint32(0xe144a693),
		8:   uint32(0x9ddbd6b2),
		9:   uint32(0xf8bcedf4),
		10:  uint32(0x5715a03e),
		11:  uint32(0x32729b78),
		12:  uint32(0xd3363deb),
		13:  uint32(0xb65106ad),
		14:  uint32(0x19f84b67),
		15:  uint32(0x7c9f7021),
		16:  uint32(0xe0c6ab25),
		17:  uint32(0x85a19063),
		18:  uint32(0x2a08dda9),
		19:  uint32(0x4f6fe6ef),
		20:  uint32(0xae2b407c),
		21:  uint32(0xcb4c7b3a),
		22:  uint32(0x64e536f0),
		23:  uint32(0x01820db6),
		24:  uint32(0x7d1d7d97),
		25:  uint32(0x187a46d1),
		26:  uint32(0xb7d30b1b),
		27:  uint32(0xd2b4305d),
		28:  uint32(0x33f096ce),
		29:  uint32(0x5697ad88),
		30:  uint32(0xf93ee042),
		31:  uint32(0x9c59db04),
		32:  uint32(0x1afc500b),
		33:  uint32(0x7f9b6b4d),
		34:  uint32(0xd0322687),
		35:  uint32(0xb5551dc1),
		36:  uint32(0x5411bb52),
		37:  uint32(0x31768014),
		38:  uint32(0x9edfcdde),
		39:  uint32(0xfbb8f698),
		40:  uint32(0x872786b9),
		41:  uint32(0xe240bdff),
		42:  uint32(0x4de9f035),
		43:  uint32(0x288ecb73),
		44:  uint32(0xc9ca6de0),
		45:  uint32(0xacad56a6),
		46:  uint32(0x03041b6c),
		47:  uint32(0x6663202a),
		48:  uint32(0xfa3afb2e),
		49:  uint32(0x9f5dc068),
		50:  uint32(0x30f48da2),
		51:  uint32(0x5593b6e4),
		52:  uint32(0xb4d71077),
		53:  uint32(0xd1b02b31),
		54:  uint32(0x7e1966fb),
		55:  uint32(0x1b7e5dbd),
		56:  uint32(0x67e12d9c),
		57:  uint32(0x028616da),
		58:  uint32(0xad2f5b10),
		59:  uint32(0xc8486056),
		60:  uint32(0x290cc6c5),
		61:  uint32(0x4c6bfd83),
		62:  uint32(0xe3c2b049),
		63:  uint32(0x86a58b0f),
		64:  uint32(0x35f8a016),
		65:  uint32(0x509f9b50),
		66:  uint32(0xff36d69a),
		67:  uint32(0x9a51eddc),
		68:  uint32(0x7b154b4f),
		69:  uint32(0x1e727009),
		70:  uint32(0xb1db3dc3),
		71:  uint32(0xd4bc0685),
		72:  uint32(0xa82376a4),
		73:  uint32(0xcd444de2),
		74:  uint32(0x62ed0028),
		75:  uint32(0x078a3b6e),
		76:  uint32(0xe6ce9dfd),
		77:  uint32(0x83a9a6bb),
		78:  uint32(0x2c00eb71),
		79:  uint32(0x4967d037),
		80:  uint32(0xd53e0b33),
		81:  uint32(0xb0593075),
		82:  uint32(0x1ff07dbf),
		83:  uint32(0x7a9746f9),
		84:  uint32(0x9bd3e06a),
		85:  uint32(0xfeb4db2c),
		86:  uint32(0x511d96e6),
		87:  uint32(0x347aada0),
		88:  uint32(0x48e5dd81),
		89:  uint32(0x2d82e6c7),
		90:  uint32(0x822bab0d),
		91:  uint32(0xe74c904b),
		92:  uint32(0x060836d8),
		93:  uint32(0x636f0d9e),
		94:  uint32(0xccc64054),
		95:  uint32(0xa9a17b12),
		96:  uint32(0x2f04f01d),
		97:  uint32(0x4a63cb5b),
		98:  uint32(0xe5ca8691),
		99:  uint32(0x80adbdd7),
		100: uint32(0x61e91b44),
		101: uint32(0x048e2002),
		102: uint32(0xab276dc8),
		103: uint32(0xce40568e),
		104: uint32(0xb2df26af),
		105: uint32(0xd7b81de9),
		106: uint32(0x78115023),
		107: uint32(0x1d766b65),
		108: uint32(0xfc32cdf6),
		109: uint32(0x9955f6b0),
		110: uint32(0x36fcbb7a),
		111: uint32(0x539b803c),
		112: uint32(0xcfc25b38),
		113: uint32(0xaaa5607e),
		114: uint32(0x050c2db4),
		115: uint32(0x606b16f2),
		116: uint32(0x812fb061),
		117: uint32(0xe4488b27),
		118: uint32(0x4be1c6ed),
		119: uint32(0x2e86fdab),
		120: uint32(0x52198d8a),
		121: uint32(0x377eb6cc),
		122: uint32(0x98d7fb06),
		123: uint32(0xfdb0c040),
		124: uint32(0x1cf466d3),
		125: uint32(0x79935d95),
		126: uint32(0xd63a105f),
		127: uint32(0xb35d2b19),
		128: uint32(0x6bf1402c),
		129: uint32(0x0e967b6a),
		130: uint32(0xa13f36a0),
		131: uint32(0xc4580de6),
		132: uint32(0x251cab75),
		133: uint32(0x407b9033),
		134: uint32(0xefd2ddf9),
		135: uint32(0x8ab5e6bf),
		136: uint32(0xf62a969e),
		137: uint32(0x934dadd8),
		138: uint32(0x3ce4e012),
		139: uint32(0x5983db54),
		140: uint32(0xb8c77dc7),
		141: uint32(0xdda04681),
		142: uint32(0x72090b4b),
		143: uint32(0x176e300d),
		144: uint32(0x8b37eb09),
		145: uint32(0xee50d04f),
		146: uint32(0x41f99d85),
		147: uint32(0x249ea6c3),
		148: uint32(0xc5da0050),
		149: uint32(0xa0bd3b16),
		150: uint32(0x0f1476dc),
		151: uint32(0x6a734d9a),
		152: uint32(0x16ec3dbb),
		153: uint32(0x738b06fd),
		154: uint32(0xdc224b37),
		155: uint32(0xb9457071),
		156: uint32(0x5801d6e2),
		157: uint32(0x3d66eda4),
		158: uint32(0x92cfa06e),
		159: uint32(0xf7a89b28),
		160: uint32(0x710d1027),
		161: uint32(0x146a2b61),
		162: uint32(0xbbc366ab),
		163: uint32(0xdea45ded),
		164: uint32(0x3fe0fb7e),
		165: uint32(0x5a87c038),
		166: uint32(0xf52e8df2),
		167: uint32(0x9049b6b4),
		168: uint32(0xecd6c695),
		169: uint32(0x89b1fdd3),
		170: uint32(0x2618b019),
		171: uint32(0x437f8b5f),
		172: uint32(0xa23b2dcc),
		173: uint32(0xc75c168a),
		174: uint32(0x68f55b40),
		175: uint32(0x0d926006),
		176: uint32(0x91cbbb02),
		177: uint32(0xf4ac8044),
		178: uint32(0x5b05cd8e),
		179: uint32(0x3e62f6c8),
		180: uint32(0xdf26505b),
		181: uint32(0xba416b1d),
		182: uint32(0x15e826d7),
		183: uint32(0x708f1d91),
		184: uint32(0x0c106db0),
		185: uint32(0x697756f6),
		186: uint32(0xc6de1b3c),
		187: uint32(0xa3b9207a),
		188: uint32(0x42fd86e9),
		189: uint32(0x279abdaf),
		190: uint32(0x8833f065),
		191: uint32(0xed54cb23),
		192: uint32(0x5e09e03a),
		193: uint32(0x3b6edb7c),
		194: uint32(0x94c796b6),
		195: uint32(0xf1a0adf0),
		196: uint32(0x10e40b63),
		197: uint32(0x75833025),
		198: uint32(0xda2a7def),
		199: uint32(0xbf4d46a9),
		200: uint32(0xc3d23688),
		201: uint32(0xa6b50dce),
		202: uint32(0x091c4004),
		203: uint32(0x6c7b7b42),
		204: uint32(0x8d3fddd1),
		205: uint32(0xe858e697),
		206: uint32(0x47f1ab5d),
		207: uint32(0x2296901b),
		208: uint32(0xbecf4b1f),
		209: uint32(0xdba87059),
		210: uint32(0x74013d93),
		211: uint32(0x116606d5),
		212: uint32(0xf022a046),
		213: uint32(0x95459b00),
		214: uint32(0x3aecd6ca),
		215: uint32(0x5f8bed8c),
		216: uint32(0x23149dad),
		217: uint32(0x4673a6eb),
		218: uint32(0xe9daeb21),
		219: uint32(0x8cbdd067),
		220: uint32(0x6df976f4),
		221: uint32(0x089e4db2),
		222: uint32(0xa7370078),
		223: uint32(0xc2503b3e),
		224: uint32(0x44f5b031),
		225: uint32(0x21928b77),
		226: uint32(0x8e3bc6bd),
		227: uint32(0xeb5cfdfb),
		228: uint32(0x0a185b68),
		229: uint32(0x6f7f602e),
		230: uint32(0xc0d62de4),
		231: uint32(0xa5b116a2),
		232: uint32(0xd92e6683),
		233: uint32(0xbc495dc5),
		234: uint32(0x13e0100f),
		235: uint32(0x76872b49),
		236: uint32(0x97c38dda),
		237: uint32(0xf2a4b69c),
		238: uint32(0x5d0dfb56),
		239: uint32(0x386ac010),
		240: uint32(0xa4331b14),
		241: uint32(0xc1542052),
		242: uint32(0x6efd6d98),
		243: uint32(0x0b9a56de),
		244: uint32(0xeadef04d),
		245: uint32(0x8fb9cb0b),
		246: uint32(0x201086c1),
		247: uint32(0x4577bd87),
		248: uint32(0x39e8cda6),
		249: uint32(0x5c8ff6e0),
		250: uint32(0xf326bb2a),
		251: uint32(0x9641806c),
		252: uint32(0x770526ff),
		253: uint32(0x12621db9),
		254: uint32(0xbdcb5073),
		255: uint32(0xd8ac6b35),
	},
	1: {
		1:   uint32(0xd7e28058),
		2:   uint32(0x74b406f1),
		3:   uint32(0xa35686a9),
		4:   uint32(0xe9680de2),
		5:   uint32(0x3e8a8dba),
		6:   uint32(0x9ddc0b13),
		7:   uint32(0x4a3e8b4b),
		8:   uint32(0x09a11d85),
		9:   uint32(0xde439ddd),
		10:  uint32(0x7d151b74),
		11:  uint32(0xaaf79b2c),
		12:  uint32(0xe0c91067),
		13:  uint32(0x372b903f),
		14:  uint32(0x947d1696),
		15:  uint32(0x439f96ce),
		16:  uint32(0x13423b0a),
		17:  uint32(0xc4a0bb52),
		18:  uint32(0x67f63dfb),
		19:  uint32(0xb014bda3),
		20:  uint32(0xfa2a36e8),
		21:  uint32(0x2dc8b6b0),
		22:  uint32(0x8e9e3019),
		23:  uint32(0x597cb041),
		24:  uint32(0x1ae3268f),
		25:  uint32(0xcd01a6d7),
		26:  uint32(0x6e57207e),
		27:  uint32(0xb9b5a026),
		28:  uint32(0xf38b2b6d),
		29:  uint32(0x2469ab35),
		30:  uint32(0x873f2d9c),
		31:  uint32(0x50ddadc4),
		32:  uint32(0x26847614),
		33:  uint32(0xf166f64c),
		34:  uint32(0x523070e5),
		35:  uint32(0x85d2f0bd),
		36:  uint32(0xcfec7bf6),
		37:  uint32(0x180efbae),
		38:  uint32(0xbb587d07),
		39:  uint32(0x6cbafd5f),
		40:  uint32(0x2f256b91),
		41:  uint32(0xf8c7ebc9),
		42:  uint32(0x5b916d60),
		43:  uint32(0x8c73ed38),
		44:  uint32(0xc64d6673),
		45:  uint32(0x11afe62b),
		46:  uint32(0xb2f96082),
		47:  uint32(0x651be0da),
		48:  uint32(0x35c64d1e),
		49:  uint32(0xe224cd46),
		50:  uint32(0x41724bef),
		51:  uint32(0x9690cbb7),
		52:  uint32(0xdcae40fc),
		53:  uint32(0x0b4cc0a4),
		54:  uint32(0xa81a460d),
		55:  uint32(0x7ff8c655),
		56:  uint32(0x3c67509b),
		57:  uint32(0xeb85d0c3),
		58:  uint32(0x48d3566a),
		59:  uint32(0x9f31d632),
		60:  uint32(0xd50f5d79),
		61:  uint32(0x02eddd21),
		62:  uint32(0xa1bb5b88),
		63:  uint32(0x7659dbd0),
		64:  uint32(0x4d08ec28),
		65:  uint32(0x9aea6c70),
		66:  uint32(0x39bcead9),
		67:  uint32(0xee5e6a81),
		68:  uint32(0xa460e1ca),
		69:  uint32(0x73826192),
		70:  uint32(0xd0d4e73b),
		71:  uint32(0x07366763),
		72:  uint32(0x44a9f1ad),
		73:  uint32(0x934b71f5),
		74:  uint32(0x301df75c),
		75:  uint32(0xe7ff7704),
		76:  uint32(0xadc1fc4f),
		77:  uint32(0x7a237c17),
		78:  uint32(0xd975fabe),
		79:  uint32(0x0e977ae6),
		80:  uint32(0x5e4ad722),
		81:  uint32(0x89a8577a),
		82:  uint32(0x2afed1d3),
		83:  uint32(0xfd1c518b),
		84:  uint32(0xb722dac0),
		85:  uint32(0x60c05a98),
		86:  uint32(0xc396dc31),
		87:  uint32(0x14745c69),
		88:  uint32(0x57ebcaa7),
		89:  uint32(0x80094aff),
		90:  uint32(0x235fcc56),
		91:  uint32(0xf4bd4c0e),
		92:  uint32(0xbe83c745),
		93:  uint32(0x6961471d),
		94:  uint32(0xca37c1b4),
		95:  uint32(0x1dd541ec),
		96:  uint32(0x6b8c9a3c),
		97:  uint32(0xbc6e1a64),
		98:  uint32(0x1f389ccd),
		99:  uint32(0xc8da1c95),
		100: uint32(0x82e497de),
		101: uint32(0x55061786),
		102: uint32(0xf650912f),
		103: uint32(0x21b21177),
		104: uint32(0x622d87b9),
		105: uint32(0xb5cf07e1),
		106: uint32(0x16998148),
		107: uint32(0xc17b0110),
		108: uint32(0x8b458a5b),
		109: uint32(0x5ca70a03),
		110: uint32(0xfff18caa),
		111: uint32(0x28130cf2),
		112: uint32(0x78cea136),
		113: uint32(0xaf2c216e),
		114: uint32(0x0c7aa7c7),
		115: uint32(0xdb98279f),
		116: uint32(0x91a6acd4),
		117: uint32(0x46442c8c),
		118: uint32(0xe512aa25),
		119: uint32(0x32f02a7d),
		120: uint32(0x716fbcb3),
		121: uint32(0xa68d3ceb),
		122: uint32(0x05dbba42),
		123: uint32(0xd2393a1a),
		124: uint32(0x9807b151),
		125: uint32(0x4fe53109),
		126: uint32(0xecb3b7a0),
		127: uint32(0x3b5137f8),
		128: uint32(0x9a11d850),
		129: uint32(0x4df35808),
		130: uint32(0xeea5dea1),
		131: uint32(0x39475ef9),
		132: uint32(0x7379d5b2),
		133: uint32(0xa49b55ea),
		134: uint32(0x07cdd343),
		135: uint32(0xd02f531b),
		136: uint32(0x93b0c5d5),
		137: uint32(0x4452458d),
		138: uint32(0xe704c324),
		139: uint32(0x30e6437c),
		140: uint32(0x7ad8c837),
		141: uint32(0xad3a486f),
		142: uint32(0x0e6ccec6),
		143: uint32(0xd98e4e9e),
		144: uint32(0x8953e35a),
		145: uint32(0x5eb16302),
		146: uint32(0xfde7e5ab),
		147: uint32(0x2a0565f3),
		148: uint32(0x603beeb8),
		149: uint32(0xb7d96ee0),
		150: uint32(0x148fe849),
		151: uint32(0xc36d6811),
		152: uint32(0x80f2fedf),
		153: uint32(0x57107e87),
		154: uint32(0xf446f82e),
		155: uint32(0x23a47876),
		156: uint32(0x699af33d),
		157: uint32(0xbe787365),
		158: uint32(0x1d2ef5cc),
		159: uint32(0xcacc7594),
		160: uint32(0xbc95ae44),
		161: uint32(0x6b772e1c),
		162: uint32(0xc821a8b5),
		163: uint32(0x1fc328ed),
		164: uint32(0x55fda3a6),
		165: uint32(0x821f23fe),
		166: uint32(0x2149a557),
		167: uint32(0xf6ab250f),
		168: uint32(0xb534b3c1),
		169: uint32(0x62d63399),
		170: uint32(0xc180b530),
		171: uint32(0x16623568),
		172: uint32(0x5c5cbe23),
		173: uint32(0x8bbe3e7b),
		174: uint32(0x28e8b8d2),
		175: uint32(0xff0a388a),
		176: uint32(0xafd7954e),
		177: uint32(0x78351516),
		178: uint32(0xdb6393bf),
		179: uint32(0x0c8113e7),
		180: uint32(0x46bf98ac),
		181: uint32(0x915d18f4),
		182: uint32(0x320b9e5d),
		183: uint32(0xe5e91e05),
		184: uint32(0xa67688cb),
		185: uint32(0x71940893),
		186: uint32(0xd2c28e3a),
		187: uint32(0x05200e62),
		188: uint32(0x4f1e8529),
		189: uint32(0x98fc0571),
		190: uint32(0x3baa83d8),
		191: uint32(0xec480380),
		192: uint32(0xd7193478),
		193: uint32(0x00fbb420),
		194: uint32(0xa3ad3289),
		195: uint32(0x744fb2d1),
		196: uint32(0x3e71399a),
		197: uint32(0xe993b9c2),
		198: uint32(0x4ac53f6b),
		199: uint32(0x9d27bf33),
		200: uint32(0xdeb829fd),
		201: uint32(0x095aa9a5),
		202: uint32(0xaa0c2f0c),
		203: uint32(0x7deeaf54),
		204: uint32(0x37d0241f),
		205: uint32(0xe032a447),
		206: uint32(0x436422ee),
		207: uint32(0x9486a2b6),
		208: uint32(0xc45b0f72),
		209: uint32(0x13b98f2a),
		210: uint32(0xb0ef0983),
		211: uint32(0x670d89db),
		212: uint32(0x2d330290),
		213: uint32(0xfad182c8),
		214: uint32(0x59870461),
		215: uint32(0x8e658439),
		216: uint32(0xcdfa12f7),
		217: uint32(0x1a1892af),
		218: uint32(0xb94e1406),
		219: uint32(0x6eac945e),
		220: uint32(0x24921f15),
		221: uint32(0xf3709f4d),
		222: uint32(0x502619e4),
		223: uint32(0x87c499bc),
		224: uint32(0xf19d426c),
		225: uint32(0x267fc234),
		226: uint32(0x8529449d),
		227: uint32(0x52cbc4c5),
		228: uint32(0x18f54f8e),
		229: uint32(0xcf17cfd6),
		230: uint32(0x6c41497f),
		231: uint32(0xbba3c927),
		232: uint32(0xf83c5fe9),
		233: uint32(0x2fdedfb1),
		234: uint32(0x8c885918),
		235: uint32(0x5b6ad940),
		236: uint32(0x1154520b),
		237: uint32(0xc6b6d253),
		238: uint32(0x65e054fa),
		239: uint32(0xb202d4a2),
		240: uint32(0xe2df7966),
		241: uint32(0x353df93e),
		242: uint32(0x966b7f97),
		243: uint32(0x4189ffcf),
		244: uint32(0x0bb77484),
		245: uint32(0xdc55f4dc),
		246: uint32(0x7f037275),
		247: uint32(0xa8e1f22d),
		248: uint32(0xeb7e64e3),
		249: uint32(0x3c9ce4bb),
		250: uint32(0x9fca6212),
		251: uint32(0x4828e24a),
		252: uint32(0x02166901),
		253: uint32(0xd5f4e959),
		254: uint32(0x76a26ff0),
		255: uint32(0xa140efa8),
	},
	2: {
		1:   uint32(0xef52b6e1),
		2:   uint32(0x05d46b83),
		3:   uint32(0xea86dd62),
		4:   uint32(0x0ba8d706),
		5:   uint32(0xe4fa61e7),
		6:   uint32(0x0e7cbc85),
		7:   uint32(0xe12e0a64),
		8:   uint32(0x1751ae0c),
		9:   uint32(0xf80318ed),
		10:  uint32(0x1285c58f),
		11:  uint32(0xfdd7736e),
		12:  uint32(0x1cf9790a),
		13:  uint32(0xf3abcfeb),
		14:  uint32(0x192d1289),
		15:  uint32(0xf67fa468),
		16:  uint32(0x2ea35c18),
		17:  uint32(0xc1f1eaf9),
		18:  uint32(0x2b77379b),
		19:  uint32(0xc425817a),
		20:  uint32(0x250b8b1e),
		21:  uint32(0xca593dff),
		22:  uint32(0x20dfe09d),
		23:  uint32(0xcf8d567c),
		24:  uint32(0x39f2f214),
		25:  uint32(0xd6a044f5),
		26:  uint32(0x3c269997),
		27:  uint32(0xd3742f76),
		28:  uint32(0x325a2512),
		29:  uint32(0xdd0893f3),
		30:  uint32(0x378e4e91),
		31:  uint32(0xd8dcf870),
		32:  uint32(0x5d46b830),
		33:  uint32(0xb2140ed1),
		34:  uint32(0x5892d3b3),
		35:  uint32(0xb7c06552),
		36:  uint32(0x56ee6f36),
		37:  uint32(0xb9bcd9d7),
		38:  uint32(0x533a04b5),
		39:  uint32(0xbc68b254),
		40:  uint32(0x4a17163c),
		41:  uint32(0xa545a0dd),
		42:  uint32(0x4fc37dbf),
		43:  uint32(0xa091cb5e),
		44:  uint32(0x41bfc13a),
		45:  uint32(0xaeed77db),
		46:  uint32(0x446baab9),
		47:  uint32(0xab391c58),
		48:  uint32(0x73e5e428),
		49:  uint32(0x9cb752c9),
		50:  uint32(0x76318fab),
		51:  uint32(0x9963394a),
		52:  uint32(0x784d332e),
		53:  uint32(0x971f85cf),
		54:  uint32(0x7d9958ad),
		55:  uint32(0x92cbee4c),
		56:  uint32(0x64b44a24),
		57:  uint32(0x8be6fcc5),
		58:  uint32(0x616021a7),
		59:  uint32(0x8e329746),
		60:  uint32(0x6f1c9d22),
		61:  uint32(0x804e2bc3),
		62:  uint32(0x6ac8f6a1),
		63:  uint32(0x859a4040),
		64:  uint32(0xba8d7060),
		65:  uint32(0x55dfc681),
		66:  uint32(0xbf591be3),
		67:  uint32(0x500bad02),
		68:  uint32(0xb125a766),
		69:  uint32(0x5e771187),
		70:  uint32(0xb4f1cce5),
		71:  uint32(0x5ba37a04),
		72:  uint32(0xaddcde6c),
		73:  uint32(0x428e688d),
		74:  uint32(0xa808b5ef),
		75:  uint32(0x475a030e),
		76:  uint32(0xa674096a),
		77:  uint32(0x4926bf8b),
		78:  uint32(0xa3a062e9),
		79:  uint32(0x4cf2d408),
		80:  uint32(0x942e2c78),
		81:  uint32(0x7b7c9a99),
		82:  uint32(0x91fa47fb),
		83:  uint32(0x7ea8f11a),
		84:  uint32(0x9f86fb7e),
		85:  uint32(0x70d44d9f),
		86:  uint32(0x9a5290fd),
		87:  uint32(0x7500261c),
		88:  uint32(0x837f8274),
		89:  uint32(0x6c2d3495),
		90:  uint32(0x86abe9f7),
		91:  uint32(0x69f95f16),
		92:  uint32(0x88d75572),
		93:  uint32(0x6785e393),
		94:  uint32(0x8d033ef1),
		95:  uint32(0x62518810),
		96:  uint32(0xe7cbc850),
		97:  uint32(0x08997eb1),
		98:  uint32(0xe21fa3d3),
		99:  uint32(0x0d4d1532),
		100: uint32(0xec631f56),
		101: uint32(0x0331a9b7),
		102: uint32(0xe9b774d5),
		103: uint32(0x06e5c234),
		104: uint32(0xf09a665c),
		105: uint32(0x1fc8d0bd),
		106: uint32(0xf54e0ddf),
		107: uint32(0x1a1cbb3e),
		108: uint32(0xfb32b15a),
		109: uint32(0x146007bb),
		110: uint32(0xfee6dad9),
		111: uint32(0x11b46c38),
		112: uint32(0xc9689448),
		113: uint32(0x263a22a9),
		114: uint32(0xccbcffcb),
		115: uint32(0x23ee492a),
		116: uint32(0xc2c0434e),
		117: uint32(0x2d92f5af),
		118: uint32(0xc71428cd),
		119: uint32(0x28469e2c),
		120: uint32(0xde393a44),
		121: uint32(0x316b8ca5),
		122: uint32(0xdbed51c7),
		123: uint32(0x34bfe726),
		124: uint32(0xd591ed42),
		125: uint32(0x3ac35ba3),
		126: uint32(0xd04586c1),
		127: uint32(0x3f173020),
		128: uint32(0xae6be681),
		129: uint32(0x41395060),
		130: uint32(0xabbf8d02),
		131: uint32(0x44ed3be3),
		132: uint32(0xa5c33187),
		133: uint32(0x4a918766),
		134: uint32(0xa0175a04),
		135: uint32(0x4f45ece5),
		136: uint32(0xb93a488d),
		137: uint32(0x5668fe6c),
		138: uint32(0xbcee230e),
		139: uint32(0x53bc95ef),
		140: uint32(0xb2929f8b),
		141: uint32(0x5dc0296a),
		142: uint32(0xb746f408),
		143: uint32(0x581442e9),
		144: uint32(0x80c8ba99),
		145: uint32(0x6f9a0c78),
		146: uint32(0x851cd11a),
		147: uint32(0x6a4e67fb),
		148: uint32(0x8b606d9f),
		149: uint32(0x6432db7e),
		150: uint32(0x8eb4061c),
		151: uint32(0x61e6b0fd),
		152: uint32(0x97991495),
		153: uint32(0x78cba274),
		154: uint32(0x924d7f16),
		155: uint32(0x7d1fc9f7),
		156: uint32(0x9c31c393),
		157: uint32(0x73637572),
		158: uint32(0x99e5a810),
		159: uint32(0x76b71ef1),
		160: uint32(0xf32d5eb1),
		161: uint32(0x1c7fe850),
		162: uint32(0xf6f93532),
		163: uint32(0x19ab83d3),
		164: uint32(0xf88589b7),
		165: uint32(0x17d73f56),
		166: uint32(0xfd51e234),
		167: uint32(0x120354d5),
		168: uint32(0xe47cf0bd),
		169: uint32(0x0b2e465c),
		170: uint32(0xe1a89b3e),
		171: uint32(0x0efa2ddf),
		172: uint32(0xefd427bb),
		173: uint32(0x0086915a),
		174: uint32(0xea004c38),
		175: uint32(0x0552fad9),
		176: uint32(0xdd8e02a9),
		177: uint32(0x32dcb448),
		178: uint32(0xd85a692a),
		179: uint32(0x3708dfcb),
		180: uint32(0xd626d5af),
		181: uint32(0x3974634e),
		182: uint32(0xd3f2be2c),
		183: uint32(0x3ca008cd),
		184: uint32(0xcadfaca5),
		185: uint32(0x258d1a44),
		186: uint32(0xcf0bc726),
		187: uint32(0x205971c7),
		188: uint32(0xc1777ba3),
		189: uint32(0x2e25cd42),
		190: uint32(0xc4a31020),
		191: uint32(0x2bf1a6c1),
		192: uint32(0x14e696e1),
		193: uint32(0xfbb42000),
		194: uint32(0x1132fd62),
		195: uint32(0xfe604b83),
		196: uint32(0x1f4e41e7),
		197: uint32(0xf01cf706),
		198: uint32(0x1a9a2a64),
		199: uint32(0xf5c89c85),
		200: uint32(0x03b738ed),
		201: uint32(0xece58e0c),
		202: uint32(0x0663536e),
		203: uint32(0xe931e58f),
		204: uint32(0x081fefeb),
		205: uint32(0xe74d590a),
		206: uint32(0x0dcb8468),
		207: uint32(0xe2993289),
		208: uint32(0x3a45caf9),
		209: uint32(0xd5177c18),
		210: uint32(0x3f91a17a),
		211: uint32(0xd0c3179b),
		212: uint32(0x31ed1dff),
		213: uint32(0xdebfab1e),
		214: uint32(0x3439767c),
		215: uint32(0xdb6bc09d),
		216: uint32(0x2d1464f5),
		217: uint32(0xc246d214),
		218: uint32(0x28c00f76),
		219: uint32(0xc792b997),
		220: uint32(0x26bcb3f3),
		221: uint32(0xc9ee0512),
		222: uint32(0x2368d870),
		223: uint32(0xcc3a6e91),
		224: uint32(0x49a02ed1),
		225: uint32(0xa6f29830),
		226: uint32(0x4c744552),
		227: uint32(0xa326f3b3),
		228: uint32(0x4208f9d7),
		229: uint32(0xad5a4f36),
		230: uint32(0x47dc9254),
		231: uint32(0xa88e24b5),
		232: uint32(0x5ef180dd),
		233: uint32(0xb1a3363c),
		234: uint32(0x5b25eb5e),
		235: uint32(0xb4775dbf),
		236: uint32(0x555957db),
		237: uint32(0xba0be13a),
		238: uint32(0x508d3c58),
		239: uint32(0xbfdf8ab9),
		240: uint32(0x670372c9),
		241: uint32(0x8851c428),
		242: uint32(0x62d7194a),
		243: uint32(0x8d85afab),
		244: uint32(0x6caba5cf),
		245: uint32(0x83f9132e),
		246: uint32(0x697fce4c),
		247: uint32(0x862d78ad),
		248: uint32(0x7052dcc5),
		249: uint32(0x9f006a24),
		250: uint32(0x7586b746),
		251: uint32(0x9ad401a7),
		252: uint32(0x7bfa0bc3),
		253: uint32(0x94a8bd22),
		254: uint32(0x7e2e6040),
		255: uint32(0x917cd6a1),
	},
	3: {
		1:   uint32(0x87a6cb43),
		2:   uint32(0xd43c90c7),
		3:   uint32(0x539a5b84),
		4:   uint32(0x730827cf),
		5:   uint32(0xf4aeec8c),
		6:   uint32(0xa734b708),
		7:   uint32(0x20927c4b),
		8:   uint32(0xe6104f9e),
		9:   uint32(0x61b684dd),
		10:  uint32(0x322cdf59),
		11:  uint32(0xb58a141a),
		12:  uint32(0x95186851),
		13:  uint32(0x12bea312),
		14:  uint32(0x4124f896),
		15:  uint32(0xc68233d5),
		16:  uint32(0x1751997d),
		17:  uint32(0x90f7523e),
		18:  uint32(0xc36d09ba),
		19:  uint32(0x44cbc2f9),
		20:  uint32(0x6459beb2),
		21:  uint32(0xe3ff75f1),
		22:  uint32(0xb0652e75),
		23:  uint32(0x37c3e536),
		24:  uint32(0xf141d6e3),
		25:  uint32(0x76e71da0),
		26:  uint32(0x257d4624),
		27:  uint32(0xa2db8d67),
		28:  uint32(0x8249f12c),
		29:  uint32(0x05ef3a6f),
		30:  uint32(0x567561eb),
		31:  uint32(0xd1d3aaa8),
		32:  uint32(0x2ea332fa),
		33:  uint32(0xa905f9b9),
		34:  uint32(0xfa9fa23d),
		35:  uint32(0x7d39697e),
		36:  uint32(0x5dab1535),
		37:  uint32(0xda0dde76),
		38:  uint32(0x899785f2),
		39:  uint32(0x0e314eb1),
		40:  uint32(0xc8b37d64),
		41:  uint32(0x4f15b627),
		42:  uint32(0x1c8feda3),
		43:  uint32(0x9b2926e0),
		44:  uint32(0xbbbb5aab),
		45:  uint32(0x3c1d91e8),
		46:  uint32(0x6f87ca6c),
		47:  uint32(0xe821012f),
		48:  uint32(0x39f2ab87),
		49:  uint32(0xbe5460c4),
		50:  uint32(0xedce3b40),
		51:  uint32(0x6a68f003),
		52:  uint32(0x4afa8c48),
		53:  uint32(0xcd5c470b),
		54:  uint32(0x9ec61c8f),
		55:  uint32(0x1960d7cc),
		56:  uint32(0xdfe2e419),
		57:  uint32(0x58442f5a),
		58:  uint32(0x0bde74de),
		59:  uint32(0x8c78bf9d),
		60:  uint32(0xaceac3d6),
		61:  uint32(0x2b4c0895),
		62:  uint32(0x78d65311),
		63:  uint32(0xff709852),
		64:  uint32(0x5d4665f4),
		65:  uint32(0xdae0aeb7),
		66:  uint32(0x897af533),
		67:  uint32(0x0edc3e70),
		68:  uint32(0x2e4e423b),
		69:  uint32(0xa9e88978),
		70:  uint32(0xfa72d2fc),
		71:  uint32(0x7dd419bf),
		72:  uint32(0xbb562a6a),
		73:  uint32(0x3cf0e129),
		74:  uint32(0x6f6abaad),
		75:  uint32(0xe8cc71ee),
		76:  uint32(0xc85e0da5),
		77:  uint32(0x4ff8c6e6),
		78:  uint32(0x1c629d62),
		79:  uint32(0x9bc45621),
		80:  uint32(0x4a17fc89),
		81:  uint32(0xcdb137ca),
		82:  uint32(0x9e2b6c4e),
		83:  uint32(0x198da70d),
		84:  uint32(0x391fdb46),
		85:  uint32(0xbeb91005),
		86:  uint32(0xed234b81),
		87:  uint32(0x6a8580c2),
		88:  uint32(0xac07b317),
		89:  uint32(0x2ba17854),
		90:  uint32(0x783b23d0),
		91:  uint32(0xff9de893),
		92:  uint32(0xdf0f94d8),
		93:  uint32(0x58a95f9b),
		94:  uint32(0x0b33041f),
		95:  uint32(0x8c95cf5c),
		96:  uint32(0x73e5570e),
		97:  uint32(0xf4439c4d),
		98:  uint32(0xa7d9c7c9),
		99:  uint32(0x207f0c8a),
		100: uint32(0x00ed70c1),
		101: uint32(0x874bbb82),
		102: uint32(0xd4d1e006),
		103: uint32(0x53772b45),
		104: uint32(0x95f51890),
		105: uint32(0x1253d3d3),
		106: uint32(0x41c98857),
		107: uint32(0xc66f4314),
		108: uint32(0xe6fd3f5f),
		109: uint32(0x615bf41c),
		110: uint32(0x32c1af98),
		111: uint32(0xb56764db),
		112: uint32(0x64b4ce73),
		113: uint32(0xe3120530),
		114: uint32(0xb0885eb4),
		115: uint32(0x372e95f7),
		116: uint32(0x17bce9bc),
		117: uint32(0x901a22ff),
		118: uint32(0xc380797b),
		119: uint32(0x4426b238),
		120: uint32(0x82a481ed),
		121: uint32(0x05024aae),
		122: uint32(0x5698112a),
		123: uint32(0xd13eda69),
		124: uint32(0xf1aca622),
		125: uint32(0x760a6d61),
		126: uint32(0x259036e5),
		127: uint32(0xa236fda6),
		128: uint32(0xba8ccbe8),
		129: uint32(0x3d2a00ab),
		130: uint32(0x6eb05b2f),
		131: uint32(0xe916906c),
		132: uint32(0xc984ec27),
		133: uint32(0x4e222764),
		134: uint32(0x1db87ce0),
		135: uint32(0x9a1eb7a3),
		136: uint32(0x5c9c8476),
		137: uint32(0xdb3a4f35),
		138: uint32(0x88a014b1),
		139: uint32(0x0f06dff2),
		140: uint32(0x2f94a3b9),
		141: uint32(0xa83268fa),
		142: uint32(0xfba8337e),
		143: uint32(0x7c0ef83d),
		144: uint32(0xaddd5295),
		145: uint32(0x2a7b99d6),
		146: uint32(0x79e1c252),
		147: uint32(0xfe470911),
		148: uint32(0xded5755a),
		149: uint32(0x5973be19),
		150: uint32(0x0ae9e59d),
		151: uint32(0x8d4f2ede),
		152: uint32(0x4bcd1d0b),
		153: uint32(0xcc6bd648),
		154: uint32(0x9ff18dcc),
		155: uint32(0x1857468f),
		156: uint32(0x38c53ac4),
		157: uint32(0xbf63f187),
		158: uint32(0xecf9aa03),
		159: uint32(0x6b5f6140),
		160: uint32(0x942ff912),
		161: uint32(0x13893251),
		162: uint32(0x401369d5),
		163: uint32(0xc7b5a296),
		164: uint32(0xe727dedd),
		165: uint32(0x6081159e),
		166: uint32(0x331b4e1a),
		167: uint32(0xb4bd8559),
		168: uint32(0x723fb68c),
		169: uint32(0xf5997dcf),
		170: uint32(0xa603264b),
		171: uint32(0x21a5ed08),
		172: uint32(0x01379143),
		173: uint32(0x86915a00),
		174: uint32(0xd50b0184),
		175: uint32(0x52adcac7),
		176: uint32(0x837e606f),
		177: uint32(0x04d8ab2c),
		178: uint32(0x5742f0a8),
		179: uint32(0xd0e43beb),
		180: uint32(0xf07647a0),
		181: uint32(0x77d08ce3),
		182: uint32(0x244ad767),
		183: uint32(0xa3ec1c24),
		184: uint32(0x656e2ff1),
		185: uint32(0xe2c8e4b2),
		186: uint32(0xb152bf36),
		187: uint32(0x36f47475),
		188: uint32(0x1666083e),
		189: uint32(0x91c0c37d),
		190: uint32(0xc25a98f9),
		191: uint32(0x45fc53ba),
		192: uint32(0xe7caae1c),
		193: uint32(0x606c655f),
		194: uint32(0x33f63edb),
		195: uint32(0xb450f598),
		196: uint32(0x94c289d3),
		197: uint32(0x13644290),
		198: uint32(0x40fe1914),
		199: uint32(0xc758d257),
		200: uint32(0x01dae182),
		201: uint32(0x867c2ac1),
		202: uint32(0xd5e67145),
		203: uint32(0x5240ba06),
		204: uint32(0x72d2c64d),
		205: uint32(0xf5740d0e),
		206: uint32(0xa6ee568a),
		207: uint32(0x21489dc9),
		208: uint32(0xf09b3761),
		209: uint32(0x773dfc22),
		210: uint32(0x24a7a7a6),
		211: uint32(0xa3016ce5),
		212: uint32(0x839310ae),
		213: uint32(0x0435dbed),
		214: uint32(0x57af8069),
		215: uint32(0xd0094b2a),
		216: uint32(0x168b78ff),
		217: uint32(0x912db3bc),
		218: uint32(0xc2b7e838),
		219: uint32(0x4511237b),
		220: uint32(0x65835f30),
		221: uint32(0xe2259473),
		222: uint32(0xb1bfcff7),
		223: uint32(0x361904b4),
		224: uint32(0xc9699ce6),
		225: uint32(0x4ecf57a5),
		226: uint32(0x1d550c21),
		227: uint32(0x9af3c762),
		228: uint32(0xba61bb29),
		229: uint32(0x3dc7706a),
		230: uint32(0x6e5d2bee),
		231: uint32(0xe9fbe0ad),
		232: uint32(0x2f79d378),
		233: uint32(0xa8df183b),
		234: uint32(0xfb4543bf),
		235: uint32(0x7ce388fc),
		236: uint32(0x5c71f4b7),
		237: uint32(0xdbd73ff4),
		238: uint32(0x884d6470),
		239: uint32(0x0febaf33),
		240: uint32(0xde38059b),
		241: uint32(0x599eced8),
		242: uint32(0x0a04955c),
		243: uint32(0x8da25e1f),
		244: uint32(0xad302254),
		245: uint32(0x2a96e917),
		246: uint32(0x790cb293),
		247: uint32(0xfeaa79d0),
		248: uint32(0x38284a05),
		249: uint32(0xbf8e8146),
		250: uint32(0xec14dac2),
		251: uint32(0x6bb21181),
		252: uint32(0x4b206dca),
		253: uint32(0xcc86a689),
		254: uint32(0x9f1cfd0d),
		255: uint32(0x18ba364e),
	},
}
var _crc_braid_big_table = [4][256]Tz_word_t{
	0: {
		1:   uint32(0x43cba687),
		2:   uint32(0xc7903cd4),
		3:   uint32(0x845b9a53),
		4:   uint32(0xcf270873),
		5:   uint32(0x8cecaef4),
		6:   uint32(0x08b734a7),
		7:   uint32(0x4b7c9220),
		8:   uint32(0x9e4f10e6),
		9:   uint32(0xdd84b661),
		10:  uint32(0x59df2c32),
		11:  uint32(0x1a148ab5),
		12:  uint32(0x51681895),
		13:  uint32(0x12a3be12),
		14:  uint32(0x96f82441),
		15:  uint32(0xd53382c6),
		16:  uint32(0x7d995117),
		17:  uint32(0x3e52f790),
		18:  uint32(0xba096dc3),
		19:  uint32(0xf9c2cb44),
		20:  uint32(0xb2be5964),
		21:  uint32(0xf175ffe3),
		22:  uint32(0x752e65b0),
		23:  uint32(0x36e5c337),
		24:  uint32(0xe3d641f1),
		25:  uint32(0xa01de776),
		26:  uint32(0x24467d25),
		27:  uint32(0x678ddba2),
		28:  uint32(0x2cf14982),
		29:  uint32(0x6f3aef05),
		30:  uint32(0xeb617556),
		31:  uint32(0xa8aad3d1),
		32:  uint32(0xfa32a32e),
		33:  uint32(0xb9f905a9),
		34:  uint32(0x3da29ffa),
		35:  uint32(0x7e69397d),
		36:  uint32(0x3515ab5d),
		37:  uint32(0x76de0dda),
		38:  uint32(0xf2859789),
		39:  uint32(0xb14e310e),
		40:  uint32(0x647db3c8),
		41:  uint32(0x27b6154f),
		42:  uint32(0xa3ed8f1c),
		43:  uint32(0xe026299b),
		44:  uint32(0xab5abbbb),
		45:  uint32(0xe8911d3c),
		46:  uint32(0x6cca876f),
		47:  uint32(0x2f0121e8),
		48:  uint32(0x87abf239),
		49:  uint32(0xc46054be),
		50:  uint32(0x403bceed),
		51:  uint32(0x03f0686a),
		52:  uint32(0x488cfa4a),
		53:  uint32(0x0b475ccd),
		54:  uint32(0x8f1cc69e),
		55:  uint32(0xccd76019),
		56:  uint32(0x19e4e2df),
		57:  uint32(0x5a2f4458),
		58:  uint32(0xde74de0b),
		59:  uint32(0x9dbf788c),
		60:  uint32(0xd6c3eaac),
		61:  uint32(0x95084c2b),
		62:  uint32(0x1153d678),
		63:  uint32(0x529870ff),
		64:  uint32(0xf465465d),
		65:  uint32(0xb7aee0da),
		66:  uint32(0x33f57a89),
		67:  uint32(0x703edc0e),
		68:  uint32(0x3b424e2e),
		69:  uint32(0x7889e8a9),
		70:  uint32(0xfcd272fa),
		71:  uint32(0xbf19d47d),
		72:  uint32(0x6a2a56bb),
		73:  uint32(0x29e1f03c),
		74:  uint32(0xadba6a6f),
		75:  uint32(0xee71cce8),
		76:  uint32(0xa50d5ec8),
		77:  uint32(0xe6c6f84f),
		78:  uint32(0x629d621c),
		79:  uint32(0x2156c49b),
		80:  uint32(0x89fc174a),
		81:  uint32(0xca37b1cd),
		82:  uint32(0x4e6c2b9e),
		83:  uint32(0x0da78d19),
		84:  uint32(0x46db1f39),
		85:  uint32(0x0510b9be),
		86:  uint32(0x814b23ed),
		87:  uint32(0xc280856a),
		88:  uint32(0x17b307ac),
		89:  uint32(0x5478a12b),
		90:  uint32(0xd0233b78),
		91:  uint32(0x93e89dff),
		92:  uint32(0xd8940fdf),
		93:  uint32(0x9b5fa958),
		94:  uint32(0x1f04330b),
		95:  uint32(0x5ccf958c),
		96:  uint32(0x0e57e573),
		97:  uint32(0x4d9c43f4),
		98:  uint32(0xc9c7d9a7),
		99:  uint32(0x8a0c7f20),
		100: uint32(0xc170ed00),
		101: uint32(0x82bb4b87),
		102: uint32(0x06e0d1d4),
		103: uint32(0x452b7753),
		104: uint32(0x9018f595),
		105: uint32(0xd3d35312),
		106: uint32(0x5788c941),
		107: uint32(0x14436fc6),
		108: uint32(0x5f3ffde6),
		109: uint32(0x1cf45b61),
		110: uint32(0x98afc132),
		111: uint32(0xdb6467b5),
		112: uint32(0x73ceb464),
		113: uint32(0x300512e3),
		114: uint32(0xb45e88b0),
		115: uint32(0xf7952e37),
		116: uint32(0xbce9bc17),
		117: uint32(0xff221a90),
		118: uint32(0x7b7980c3),
		119: uint32(0x38b22644),
		120: uint32(0xed81a482),
		121: uint32(0xae4a0205),
		122: uint32(0x2a119856),
		123: uint32(0x69da3ed1),
		124: uint32(0x22a6acf1),
		125: uint32(0x616d0a76),
		126: uint32(0xe5369025),
		127: uint32(0xa6fd36a2),
		128: uint32(0xe8cb8cba),
		129: uint32(0xab002a3d),
		130: uint32(0x2f5bb06e),
		131: uint32(0x6c9016e9),
		132: uint32(0x27ec84c9),
		133: uint32(0x6427224e),
		134: uint32(0xe07cb81d),
		135: uint32(0xa3b71e9a),
		136: uint32(0x76849c5c),
		137: uint32(0x354f3adb),
		138: uint32(0xb114a088),
		139: uint32(0xf2df060f),
		140: uint32(0xb9a3942f),
		141: uint32(0xfa6832a8),
		142: uint32(0x7e33a8fb),
		143: uint32(0x3df80e7c),
		144: uint32(0x9552ddad),
		145: uint32(0xd6997b2a),
		146: uint32(0x52c2e179),
		147: uint32(0x110947fe),
		148: uint32(0x5a75d5de),
		149: uint32(0x19be7359),
		150: uint32(0x9de5e90a),
		151: uint32(0xde2e4f8d),
		152: uint32(0x0b1dcd4b),
		153: uint32(0x48d66bcc),
		154: uint32(0xcc8df19f),
		155: uint32(0x8f465718),
		156: uint32(0xc43ac538),
		157: uint32(0x87f163bf),
		158: uint32(0x03aaf9ec),
		159: uint32(0x40615f6b),
		160: uint32(0x12f92f94),
		161: uint32(0x51328913),
		162: uint32(0xd5691340),
		163: uint32(0x96a2b5c7),
		164: uint32(0xddde27e7),
		165: uint32(0x9e158160),
		166: uint32(0x1a4e1b33),
		167: uint32(0x5985bdb4),
		168: uint32(0x8cb63f72),
		169: uint32(0xcf7d99f5),
		170: uint32(0x4b2603a6),
		171: uint32(0x08eda521),
		172: uint32(0x43913701),
		173: uint32(0x005a9186),
		174: uint32(0x84010bd5),
		175: uint32(0xc7caad52),
		176: uint32(0x6f607e83),
		177: uint32(0x2cabd804),
		178: uint32(0xa8f04257),
		179: uint32(0xeb3be4d0),
		180: uint32(0xa04776f0),
		181: uint32(0xe38cd077),
		182: uint32(0x67d74a24),
		183: uint32(0x241ceca3),
		184: uint32(0xf12f6e65),
		185: uint32(0xb2e4c8e2),
		186: uint32(0x36bf52b1),
		187: uint32(0x7574f436),
		188: uint32(0x3e086616),
		189: uint32(0x7dc3c091),
		190: uint32(0xf9985ac2),
		191: uint32(0xba53fc45),
		192: uint32(0x1caecae7),
		193: uint32(0x5f656c60),
		194: uint32(0xdb3ef633),
		195: uint32(0x98f550b4),
		196: uint32(0xd389c294),
		197: uint32(0x90426413),
		198: uint32(0x1419fe40),
		199: uint32(0x57d258c7),
		200: uint32(0x82e1da01),
		201: uint32(0xc12a7c86),
		202: uint32(0x4571e6d5),
		203: uint32(0x06ba4052),
		204: uint32(0x4dc6d272),
		205: uint32(0x0e0d74f5),
		206: uint32(0x8a56eea6),
		207: uint32(0xc99d4821),
		208: uint32(0x61379bf0),
		209: uint32(0x22fc3d77),
		210: uint32(0xa6a7a724),
		211: uint32(0xe56c01a3),
		212: uint32(0xae109383),
		213: uint32(0xeddb3504),
		214: uint32(0x6980af57),
		215: uint32(0x2a4b09d0),
		216: uint32(0xff788b16),
		217: uint32(0xbcb32d91),
		218: uint32(0x38e8b7c2),
		219: uint32(0x7b231145),
		220: uint32(0x305f8365),
		221: uint32(0x739425e2),
		222: uint32(0xf7cfbfb1),
		223: uint32(0xb4041936),
		224: uint32(0xe69c69c9),
		225: uint32(0xa557cf4e),
		226: uint32(0x210c551d),
		227: uint32(0x62c7f39a),
		228: uint32(0x29bb61ba),
		229: uint32(0x6a70c73d),
		230: uint32(0xee2b5d6e),
		231: uint32(0xade0fbe9),
		232: uint32(0x78d3792f),
		233: uint32(0x3b18dfa8),
		234: uint32(0xbf4345fb),
		235: uint32(0xfc88e37c),
		236: uint32(0xb7f4715c),
		237: uint32(0xf43fd7db),
		238: uint32(0x70644d88),
		239: uint32(0x33afeb0f),
		240: uint32(0x9b0538de),
		241: uint32(0xd8ce9e59),
		242: uint32(0x5c95040a),
		243: uint32(0x1f5ea28d),
		244: uint32(0x542230ad),
		245: uint32(0x17e9962a),
		246: uint32(0x93b20c79),
		247: uint32(0xd079aafe),
		248: uint32(0x054a2838),
		249: uint32(0x46818ebf),
		250: uint32(0xc2da14ec),
		251: uint32(0x8111b26b),
		252: uint32(0xca6d204b),
		253: uint32(0x89a686cc),
		254: uint32(0x0dfd1c9f),
		255: uint32(0x4e36ba18),
	},
	1: {
		1:   uint32(0xe1b652ef),
		2:   uint32(0x836bd405),
		3:   uint32(0x62dd86ea),
		4:   uint32(0x06d7a80b),
		5:   uint32(0xe761fae4),
		6:   uint32(0x85bc7c0e),
		7:   uint32(0x640a2ee1),
		8:   uint32(0x0cae5117),
		9:   uint32(0xed1803f8),
		10:  uint32(0x8fc58512),
		11:  uint32(0x6e73d7fd),
		12:  uint32(0x0a79f91c),
		13:  uint32(0xebcfabf3),
		14:  uint32(0x89122d19),
		15:  uint32(0x68a47ff6),
		16:  uint32(0x185ca32e),
		17:  uint32(0xf9eaf1c1),
		18:  uint32(0x9b37772b),
		19:  uint32(0x7a8125c4),
		20:  uint32(0x1e8b0b25),
		21:  uint32(0xff3d59ca),
		22:  uint32(0x9de0df20),
		23:  uint32(0x7c568dcf),
		24:  uint32(0x14f2f239),
		25:  uint32(0xf544a0d6),
		26:  uint32(0x9799263c),
		27:  uint32(0x762f74d3),
		28:  uint32(0x12255a32),
		29:  uint32(0xf39308dd),
		30:  uint32(0x914e8e37),
		31:  uint32(0x70f8dcd8),
		32:  uint32(0x30b8465d),
		33:  uint32(0xd10e14b2),
		34:  uint32(0xb3d39258),
		35:  uint32(0x5265c0b7),
		36:  uint32(0x366fee56),
		37:  uint32(0xd7d9bcb9),
		38:  uint32(0xb5043a53),
		39:  uint32(0x54b268bc),
		40:  uint32(0x3c16174a),
		41:  uint32(0xdda045a5),
		42:  uint32(0xbf7dc34f),
		43:  uint32(0x5ecb91a0),
		44:  uint32(0x3ac1bf41),
		45:  uint32(0xdb77edae),
		46:  uint32(0xb9aa6b44),
		47:  uint32(0x581c39ab),
		48:  uint32(0x28e4e573),
		49:  uint32(0xc952b79c),
		50:  uint32(0xab8f3176),
		51:  uint32(0x4a396399),
		52:  uint32(0x2e334d78),
		53:  uint32(0xcf851f97),
		54:  uint32(0xad58997d),
		55:  uint32(0x4ceecb92),
		56:  uint32(0x244ab464),
		57:  uint32(0xc5fce68b),
		58:  uint32(0xa7216061),
		59:  uint32(0x4697328e),
		60:  uint32(0x229d1c6f),
		61:  uint32(0xc32b4e80),
		62:  uint32(0xa1f6c86a),
		63:  uint32(0x40409a85),
		64:  uint32(0x60708dba),
		65:  uint32(0x81c6df55),
		66:  uint32(0xe31b59bf),
		67:  uint32(0x02ad0b50),
		68:  uint32(0x66a725b1),
		69:  uint32(0x8711775e),
		70:  uint32(0xe5ccf1b4),
		71:  uint32(0x047aa35b),
		72:  uint32(0x6cdedcad),
		73:  uint32(0x8d688e42),
		74:  uint32(0xefb508a8),
		75:  uint32(0x0e035a47),
		76:  uint32(0x6a0974a6),
		77:  uint32(0x8bbf2649),
		78:  uint32(0xe962a0a3),
		79:  uint32(0x08d4f24c),
		80:  uint32(0x782c2e94),
		81:  uint32(0x999a7c7b),
		82:  uint32(0xfb47fa91),
		83:  uint32(0x1af1a87e),
		84:  uint32(0x7efb869f),
		85:  uint32(0x9f4dd470),
		86:  uint32(0xfd90529a),
		87:  uint32(0x1c260075),
		88:  uint32(0x74827f83),
		89:  uint32(0x95342d6c),
		90:  uint32(0xf7e9ab86),
		91:  uint32(0x165ff969),
		92:  uint32(0x7255d788),
		93:  uint32(0x93e38567),
		94:  uint32(0xf13e038d),
		95:  uint32(0x10885162),
		96:  uint32(0x50c8cbe7),
		97:  uint32(0xb17e9908),
		98:  uint32(0xd3a31fe2),
		99:  uint32(0x32154d0d),
		100: uint32(0x561f63ec),
		101: uint32(0xb7a93103),
		102: uint32(0xd574b7e9),
		103: uint32(0x34c2e506),
		104: uint32(0x5c669af0),
		105: uint32(0xbdd0c81f),
		106: uint32(0xdf0d4ef5),
		107: uint32(0x3ebb1c1a),
		108: uint32(0x5ab132fb),
		109: uint32(0xbb076014),
		110: uint32(0xd9dae6fe),
		111: uint32(0x386cb411),
		112: uint32(0x489468c9),
		113: uint32(0xa9223a26),
		114: uint32(0xcbffbccc),
		115: uint32(0x2a49ee23),
		116: uint32(0x4e43c0c2),
		117: uint32(0xaff5922d),
		118: uint32(0xcd2814c7),
		119: uint32(0x2c9e4628),
		120: uint32(0x443a39de),
		121: uint32(0xa58c6b31),
		122: uint32(0xc751eddb),
		123: uint32(0x26e7bf34),
		124: uint32(0x42ed91d5),
		125: uint32(0xa35bc33a),
		126: uint32(0xc18645d0),
		127: uint32(0x2030173f),
		128: uint32(0x81e66bae),
		129: uint32(0x60503941),
		130: uint32(0x028dbfab),
		131: uint32(0xe33bed44),
		132: uint32(0x8731c3a5),
		133: uint32(0x6687914a),
		134: uint32(0x045a17a0),
		135: uint32(0xe5ec454f),
		136: uint32(0x8d483ab9),
		137: uint32(0x6cfe6856),
		138: uint32(0x0e23eebc),
		139: uint32(0xef95bc53),
		140: uint32(0x8b9f92b2),
		141: uint32(0x6a29c05d),
		142: uint32(0x08f446b7),
		143: uint32(0xe9421458),
		144: uint32(0x99bac880),
		145: uint32(0x780c9a6f),
		146: uint32(0x1ad11c85),
		147: uint32(0xfb674e6a),
		148: uint32(0x9f6d608b),
		149: uint32(0x7edb3264),
		150: uint32(0x1c06b48e),
		151: uint32(0xfdb0e661),
		152: uint32(0x95149997),
		153: uint32(0x74a2cb78),
		154: uint32(0x167f4d92),
		155: uint32(0xf7c91f7d),
		156: uint32(0x93c3319c),
		157: uint32(0x72756373),
		158: uint32(0x10a8e599),
		159: uint32(0xf11eb776),
		160: uint32(0xb15e2df3),
		161: uint32(0x50e87f1c),
		162: uint32(0x3235f9f6),
		163: uint32(0xd383ab19),
		164: uint32(0xb78985f8),
		165: uint32(0x563fd717),
		166: uint32(0x34e251fd),
		167: uint32(0xd5540312),
		168: uint32(0xbdf07ce4),
		169: uint32(0x5c462e0b),
		170: uint32(0x3e9ba8e1),
		171: uint32(0xdf2dfa0e),
		172: uint32(0xbb27d4ef),
		173: uint32(0x5a918600),
		174: uint32(0x384c00ea),
		175: uint32(0xd9fa5205),
		176: uint32(0xa9028edd),
		177: uint32(0x48b4dc32),
		178: uint32(0x2a695ad8),
		179: uint32(0xcbdf0837),
		180: uint32(0xafd526d6),
		181: uint32(0x4e637439),
		182: uint32(0x2cbef2d3),
		183: uint32(0xcd08a03c),
		184: uint32(0xa5acdfca),
		185: uint32(0x441a8d25),
		186: uint32(0x26c70bcf),
		187: uint32(0xc7715920),
		188: uint32(0xa37b77c1),
		189: uint32(0x42cd252e),
		190: uint32(0x2010a3c4),
		191: uint32(0xc1a6f12b),
		192: uint32(0xe196e614),
		193: uint32(0x0020b4fb),
		194: uint32(0x62fd3211),
		195: uint32(0x834b60fe),
		196: uint32(0xe7414e1f),
		197: uint32(0x06f71cf0),
		198: uint32(0x642a9a1a),
		199: uint32(0x859cc8f5),
		200: uint32(0xed38b703),
		201: uint32(0x0c8ee5ec),
		202: uint32(0x6e536306),
		203: uint32(0x8fe531e9),
		204: uint32(0xebef1f08),
		205: uint32(0x0a594de7),
		206: uint32(0x6884cb0d),
		207: uint32(0x893299e2),
		208: uint32(0xf9ca453a),
		209: uint32(0x187c17d5),
		210: uint32(0x7aa1913f),
		211: uint32(0x9b17c3d0),
		212: uint32(0xff1ded31),
		213: uint32(0x1eabbfde),
		214: uint32(0x7c763934),
		215: uint32(0x9dc06bdb),
		216: uint32(0xf564142d),
		217: uint32(0x14d246c2),
		218: uint32(0x760fc028),
		219: uint32(0x97b992c7),
		220: uint32(0xf3b3bc26),
		221: uint32(0x1205eec9),
		222: uint32(0x70d86823),
		223: uint32(0x916e3acc),
		224: uint32(0xd12ea049),
		225: uint32(0x3098f2a6),
		226: uint32(0x5245744c),
		227: uint32(0xb3f326a3),
		228: uint32(0xd7f90842),
		229: uint32(0x364f5aad),
		230: uint32(0x5492dc47),
		231: uint32(0xb5248ea8),
		232: uint32(0xdd80f15e),
		233: uint32(0x3c36a3b1),
		234: uint32(0x5eeb255b),
		235: uint32(0xbf5d77b4),
		236: uint32(0xdb575955),
		237: uint32(0x3ae10bba),
		238: uint32(0x583c8d50),
		239: uint32(0xb98adfbf),
		240: uint32(0xc9720367),
		241: uint32(0x28c45188),
		242: uint32(0x4a19d762),
		243: uint32(0xabaf858d),
		244: uint32(0xcfa5ab6c),
		245: uint32(0x2e13f983),
		246: uint32(0x4cce7f69),
		247: uint32(0xad782d86),
		248: uint32(0xc5dc5270),
		249: uint32(0x246a009f),
		250: uint32(0x46b78675),
		251: uint32(0xa701d49a),
		252: uint32(0xc30bfa7b),
		253: uint32(0x22bda894),
		254: uint32(0x40602e7e),
		255: uint32(0xa1d67c91),
	},
	2: {
		1:   uint32(0x5880e2d7),
		2:   uint32(0xf106b474),
		3:   uint32(0xa98656a3),
		4:   uint32(0xe20d68e9),
		5:   uint32(0xba8d8a3e),
		6:   uint32(0x130bdc9d),
		7:   uint32(0x4b8b3e4a),
		8:   uint32(0x851da109),
		9:   uint32(0xdd9d43de),
		10:  uint32(0x741b157d),
		11:  uint32(0x2c9bf7aa),
		12:  uint32(0x6710c9e0),
		13:  uint32(0x3f902b37),
		14:  uint32(0x96167d94),
		15:  uint32(0xce969f43),
		16:  uint32(0x0a3b4213),
		17:  uint32(0x52bba0c4),
		18:  uint32(0xfb3df667),
		19:  uint32(0xa3bd14b0),
		20:  uint32(0xe8362afa),
		21:  uint32(0xb0b6c82d),
		22:  uint32(0x19309e8e),
		23:  uint32(0x41b07c59),
		24:  uint32(0x8f26e31a),
		25:  uint32(0xd7a601cd),
		26:  uint32(0x7e20576e),
		27:  uint32(0x26a0b5b9),
		28:  uint32(0x6d2b8bf3),
		29:  uint32(0x35ab6924),
		30:  uint32(0x9c2d3f87),
		31:  uint32(0xc4addd50),
		32:  uint32(0x14768426),
		33:  uint32(0x4cf666f1),
		34:  uint32(0xe5703052),
		35:  uint32(0xbdf0d285),
		36:  uint32(0xf67beccf),
		37:  uint32(0xaefb0e18),
		38:  uint32(0x077d58bb),
		39:  uint32(0x5ffdba6c),
		40:  uint32(0x916b252f),
		41:  uint32(0xc9ebc7f8),
		42:  uint32(0x606d915b),
		43:  uint32(0x38ed738c),
		44:  uint32(0x73664dc6),
		45:  uint32(0x2be6af11),
		46:  uint32(0x8260f9b2),
		47:  uint32(0xdae01b65),
		48:  uint32(0x1e4dc635),
		49:  uint32(0x46cd24e2),
		50:  uint32(0xef4b7241),
		51:  uint32(0xb7cb9096),
		52:  uint32(0xfc40aedc),
		53:  uint32(0xa4c04c0b),
		54:  uint32(0x0d461aa8),
		55:  uint32(0x55c6f87f),
		56:  uint32(0x9b50673c),
		57:  uint32(0xc3d085eb),
		58:  uint32(0x6a56d348),
		59:  uint32(0x32d6319f),
		60:  uint32(0x795d0fd5),
		61:  uint32(0x21dded02),
		62:  uint32(0x885bbba1),
		63:  uint32(0xd0db5976),
		64:  uint32(0x28ec084d),
		65:  uint32(0x706cea9a),
		66:  uint32(0xd9eabc39),
		67:  uint32(0x816a5eee),
		68:  uint32(0xcae160a4),
		69:  uint32(0x92618273),
		70:  uint32(0x3be7d4d0),
		71:  uint32(0x63673607),
		72:  uint32(0xadf1a944),
		73:  uint32(0xf5714b93),
		74:  uint32(0x5cf71d30),
		75:  uint32(0x0477ffe7),
		76:  uint32(0x4ffcc1ad),
		77:  uint32(0x177c237a),
		78:  uint32(0xbefa75d9),
		79:  uint32(0xe67a970e),
		80:  uint32(0x22d74a5e),
		81:  uint32(0x7a57a889),
		82:  uint32(0xd3d1fe2a),
		83:  uint32(0x8b511cfd),
		84:  uint32(0xc0da22b7),
		85:  uint32(0x985ac060),
		86:  uint32(0x31dc96c3),
		87:  uint32(0x695c7414),
		88:  uint32(0xa7caeb57),
		89:  uint32(0xff4a0980),
		90:  uint32(0x56cc5f23),
		91:  uint32(0x0e4cbdf4),
		92:  uint32(0x45c783be),
		93:  uint32(0x1d476169),
		94:  uint32(0xb4c137ca),
		95:  uint32(0xec41d51d),
		96:  uint32(0x3c9a8c6b),
		97:  uint32(0x641a6ebc),
		98:  uint32(0xcd9c381f),
		99:  uint32(0x951cdac8),
		100: uint32(0xde97e482),
		101: uint32(0x86170655),
		102: uint32(0x2f9150f6),
		103: uint32(0x7711b221),
		104: uint32(0xb9872d62),
		105: uint32(0xe107cfb5),
		106: uint32(0x48819916),
		107: uint32(0x10017bc1),
		108: uint32(0x5b8a458b),
		109: uint32(0x030aa75c),
		110: uint32(0xaa8cf1ff),
		111: uint32(0xf20c1328),
		112: uint32(0x36a1ce78),
		113: uint32(0x6e212caf),
		114: uint32(0xc7a77a0c),
		115: uint32(0x9f2798db),
		116: uint32(0xd4aca691),
		117: uint32(0x8c2c4446),
		118: uint32(0x25aa12e5),
		119: uint32(0x7d2af032),
		120: uint32(0xb3bc6f71),
		121: uint32(0xeb3c8da6),
		122: uint32(0x42badb05),
		123: uint32(0x1a3a39d2),
		124: uint32(0x51b10798),
		125: uint32(0x0931e54f),
		126: uint32(0xa0b7b3ec),
		127: uint32(0xf837513b),
		128: uint32(0x50d8119a),
		129: uint32(0x0858f34d),
		130: uint32(0xa1dea5ee),
		131: uint32(0xf95e4739),
		132: uint32(0xb2d57973),
		133: uint32(0xea559ba4),
		134: uint32(0x43d3cd07),
		135: uint32(0x1b532fd0),
		136: uint32(0xd5c5b093),
		137: uint32(0x8d455244),
		138: uint32(0x24c304e7),
		139: uint32(0x7c43e630),
		140: uint32(0x37c8d87a),
		141: uint32(0x6f483aad),
		142: uint32(0xc6ce6c0e),
		143: uint32(0x9e4e8ed9),
		144: uint32(0x5ae35389),
		145: uint32(0x0263b15e),
		146: uint32(0xabe5e7fd),
		147: uint32(0xf365052a),
		148: uint32(0xb8ee3b60),
		149: uint32(0xe06ed9b7),
		150: uint32(0x49e88f14),
		151: uint32(0x11686dc3),
		152: uint32(0xdffef280),
		153: uint32(0x877e1057),
		154: uint32(0x2ef846f4),
		155: uint32(0x7678a423),
		156: uint32(0x3df39a69),
		157: uint32(0x657378be),
		158: uint32(0xccf52e1d),
		159: uint32(0x9475ccca),
		160: uint32(0x44ae95bc),
		161: uint32(0x1c2e776b),
		162: uint32(0xb5a821c8),
		163: uint32(0xed28c31f),
		164: uint32(0xa6a3fd55),
		165: uint32(0xfe231f82),
		166: uint32(0x57a54921),
		167: uint32(0x0f25abf6),
		168: uint32(0xc1b334b5),
		169: uint32(0x9933d662),
		170: uint32(0x30b580c1),
		171: uint32(0x68356216),
		172: uint32(0x23be5c5c),
		173: uint32(0x7b3ebe8b),
		174: uint32(0xd2b8e828),
		175: uint32(0x8a380aff),
		176: uint32(0x4e95d7af),
		177: uint32(0x16153578),
		178: uint32(0xbf9363db),
		179: uint32(0xe713810c),
		180: uint32(0xac98bf46),
		181: uint32(0xf4185d91),
		182: uint32(0x5d9e0b32),
		183: uint32(0x051ee9e5),
		184: uint32(0xcb8876a6),
		185: uint32(0x93089471),
		186: uint32(0x3a8ec2d2),
		187: uint32(0x620e2005),
		188: uint32(0x29851e4f),
		189: uint32(0x7105fc98),
		190: uint32(0xd883aa3b),
		191: uint32(0x800348ec),
		192: uint32(0x783419d7),
		193: uint32(0x20b4fb00),
		194: uint32(0x8932ada3),
		195: uint32(0xd1b24f74),
		196: uint32(0x9a39713e),
		197: uint32(0xc2b993e9),
		198: uint32(0x6b3fc54a),
		199: uint32(0x33bf279d),
		200: uint32(0xfd29b8de),
		201: uint32(0xa5a95a09),
		202: uint32(0x0c2f0caa),
		203: uint32(0x54afee7d),
		204: uint32(0x1f24d037),
		205: uint32(0x47a432e0),
		206: uint32(0xee226443),
		207: uint32(0xb6a28694),
		208: uint32(0x720f5bc4),
		209: uint32(0x2a8fb913),
		210: uint32(0x8309efb0),
		211: uint32(0xdb890d67),
		212: uint32(0x9002332d),
		213: uint32(0xc882d1fa),
		214: uint32(0x61048759),
		215: uint32(0x3984658e),
		216: uint32(0xf712facd),
		217: uint32(0xaf92181a),
		218: uint32(0x06144eb9),
		219: uint32(0x5e94ac6e),
		220: uint32(0x151f9224),
		221: uint32(0x4d9f70f3),
		222: uint32(0xe4192650),
		223: uint32(0xbc99c487),
		224: uint32(0x6c429df1),
		225: uint32(0x34c27f26),
		226: uint32(0x9d442985),
		227: uint32(0xc5c4cb52),
		228: uint32(0x8e4ff518),
		229: uint32(0xd6cf17cf),
		230: uint32(0x7f49416c),
		231: uint32(0x27c9a3bb),
		232: uint32(0xe95f3cf8),
		233: uint32(0xb1dfde2f),
		234: uint32(0x1859888c),
		235: uint32(0x40d96a5b),
		236: uint32(0x0b525411),
		237: uint32(0x53d2b6c6),
		238: uint32(0xfa54e065),
		239: uint32(0xa2d402b2),
		240: uint32(0x6679dfe2),
		241: uint32(0x3ef93d35),
		242: uint32(0x977f6b96),
		243: uint32(0xcfff8941),
		244: uint32(0x8474b70b),
		245: uint32(0xdcf455dc),
		246: uint32(0x7572037f),
		247: uint32(0x2df2e1a8),
		248: uint32(0xe3647eeb),
		249: uint32(0xbbe49c3c),
		250: uint32(0x1262ca9f),
		251: uint32(0x4ae22848),
		252: uint32(0x01691602),
		253: uint32(0x59e9f4d5),
		254: uint32(0xf06fa276),
		255: uint32(0xa8ef40a1),
	},
	3: {
		1:   uint32(0x463b6765),
		2:   uint32(0x8c76ceca),
		3:   uint32(0xca4da9af),
		4:   uint32(0x59ebed4e),
		5:   uint32(0x1fd08a2b),
		6:   uint32(0xd59d2384),
		7:   uint32(0x93a644e1),
		8:   uint32(0xb2d6db9d),
		9:   uint32(0xf4edbcf8),
		10:  uint32(0x3ea01557),
		11:  uint32(0x789b7232),
		12:  uint32(0xeb3d36d3),
		13:  uint32(0xad0651b6),
		14:  uint32(0x674bf819),
		15:  uint32(0x21709f7c),
		16:  uint32(0x25abc6e0),
		17:  uint32(0x6390a185),
		18:  uint32(0xa9dd082a),
		19:  uint32(0xefe66f4f),
		20:  uint32(0x7c402bae),
		21:  uint32(0x3a7b4ccb),
		22:  uint32(0xf036e564),
		23:  uint32(0xb60d8201),
		24:  uint32(0x977d1d7d),
		25:  uint32(0xd1467a18),
		26:  uint32(0x1b0bd3b7),
		27:  uint32(0x5d30b4d2),
		28:  uint32(0xce96f033),
		29:  uint32(0x88ad9756),
		30:  uint32(0x42e03ef9),
		31:  uint32(0x04db599c),
		32:  uint32(0x0b50fc1a),
		33:  uint32(0x4d6b9b7f),
		34:  uint32(0x872632d0),
		35:  uint32(0xc11d55b5),
		36:  uint32(0x52bb1154),
		37:  uint32(0x14807631),
		38:  uint32(0xdecddf9e),
		39:  uint32(0x98f6b8fb),
		40:  uint32(0xb9862787),
		41:  uint32(0xffbd40e2),
		42:  uint32(0x35f0e94d),
		43:  uint32(0x73cb8e28),
		44:  uint32(0xe06dcac9),
		45:  uint32(0xa656adac),
		46:  uint32(0x6c1b0403),
		47:  uint32(0x2a206366),
		48:  uint32(0x2efb3afa),
		49:  uint32(0x68c05d9f),
		50:  uint32(0xa28df430),
		51:  uint32(0xe4b69355),
		52:  uint32(0x7710d7b4),
		53:  uint32(0x312bb0d1),
		54:  uint32(0xfb66197e),
		55:  uint32(0xbd5d7e1b),
		56:  uint32(0x9c2de167),
		57:  uint32(0xda168602),
		58:  uint32(0x105b2fad),
		59:  uint32(0x566048c8),
		60:  uint32(0xc5c60c29),
		61:  uint32(0x83fd6b4c),
		62:  uint32(0x49b0c2e3),
		63:  uint32(0x0f8ba586),
		64:  uint32(0x16a0f835),
		65:  uint32(0x509b9f50),
		66:  uint32(0x9ad636ff),
		67:  uint32(0xdced519a),
		68:  uint32(0x4f4b157b),
		69:  uint32(0x0970721e),
		70:  uint32(0xc33ddbb1),
		71:  uint32(0x8506bcd4),
		72:  uint32(0xa47623a8),
		73:  uint32(0xe24d44cd),
		74:  uint32(0x2800ed62),
		75:  uint32(0x6e3b8a07),
		76:  uint32(0xfd9dcee6),
		77:  uint32(0xbba6a983),
		78:  uint32(0x71eb002c),
		79:  uint32(0x37d06749),
		80:  uint32(0x330b3ed5),
		81:  uint32(0x753059b0),
		82:  uint32(0xbf7df01f),
		83:  uint32(0xf946977a),
		84:  uint32(0x6ae0d39b),
		85:  uint32(0x2cdbb4fe),
		86:  uint32(0xe6961d51),
		87:  uint32(0xa0ad7a34),
		88:  uint32(0x81dde548),
		89:  uint32(0xc7e6822d),
		90:  uint32(0x0dab2b82),
		91:  uint32(0x4b904ce7),
		92:  uint32(0xd8360806),
		93:  uint32(0x9e0d6f63),
		94:  uint32(0x5440c6cc),
		95:  uint32(0x127ba1a9),
		96:  uint32(0x1df0042f),
		97:  uint32(0x5bcb634a),
		98:  uint32(0x9186cae5),
		99:  uint32(0xd7bdad80),
		100: uint32(0x441be961),
		101: uint32(0x02208e04),
		102: uint32(0xc86d27ab),
		103: uint32(0x8e5640ce),
		104: uint32(0xaf26dfb2),
		105: uint32(0xe91db8d7),
		106: uint32(0x23501178),
		107: uint32(0x656b761d),
		108: uint32(0xf6cd32fc),
		109: uint32(0xb0f65599),
		110: uint32(0x7abbfc36),
		111: uint32(0x3c809b53),
		112: uint32(0x385bc2cf),
		113: uint32(0x7e60a5aa),
		114: uint32(0xb42d0c05),
		115: uint32(0xf2166b60),
		116: uint32(0x61b02f81),
		117: uint32(0x278b48e4),
		118: uint32(0xedc6e14b),
		119: uint32(0xabfd862e),
		120: uint32(0x8a8d1952),
		121: uint32(0xccb67e37),
		122: uint32(0x06fbd798),
		123: uint32(0x40c0b0fd),
		124: uint32(0xd366f41c),
		125: uint32(0x955d9379),
		126: uint32(0x5f103ad6),
		127: uint32(0x192b5db3),
		128: uint32(0x2c40f16b),
		129: uint32(0x6a7b960e),
		130: uint32(0xa0363fa1),
		131: uint32(0xe60d58c4),
		132: uint32(0x75ab1c25),
		133: uint32(0x33907b40),
		134: uint32(0xf9ddd2ef),
		135: uint32(0xbfe6b58a),
		136: uint32(0x9e962af6),
		137: uint32(0xd8ad4d93),
		138: uint32(0x12e0e43c),
		139: uint32(0x54db8359),
		140: uint32(0xc77dc7b8),
		141: uint32(0x8146a0dd),
		142: uint32(0x4b0b0972),
		143: uint32(0x0d306e17),
		144: uint32(0x09eb378b),
		145: uint32(0x4fd050ee),
		146: uint32(0x859df941),
		147: uint32(0xc3a69e24),
		148: uint32(0x5000dac5),
		149: uint32(0x163bbda0),
		150: uint32(0xdc76140f),
		151: uint32(0x9a4d736a),
		152: uint32(0xbb3dec16),
		153: uint32(0xfd068b73),
		154: uint32(0x374b22dc),
		155: uint32(0x717045b9),
		156: uint32(0xe2d60158),
		157: uint32(0xa4ed663d),
		158: uint32(0x6ea0cf92),
		159: uint32(0x289ba8f7),
		160: uint32(0x27100d71),
		161: uint32(0x612b6a14),
		162: uint32(0xab66c3bb),
		163: uint32(0xed5da4de),
		164: uint32(0x7efbe03f),
		165: uint32(0x38c0875a),
		166: uint32(0xf28d2ef5),
		167: uint32(0xb4b64990),
		168: uint32(0x95c6d6ec),
		169: uint32(0xd3fdb189),
		170: uint32(0x19b01826),
		171: uint32(0x5f8b7f43),
		172: uint32(0xcc2d3ba2),
		173: uint32(0x8a165cc7),
		174: uint32(0x405bf568),
		175: uint32(0x0660920d),
		176: uint32(0x02bbcb91),
		177: uint32(0x4480acf4),
		178: uint32(0x8ecd055b),
		179: uint32(0xc8f6623e),
		180: uint32(0x5b5026df),
		181: uint32(0x1d6b41ba),
		182: uint32(0xd726e815),
		183: uint32(0x911d8f70),
		184: uint32(0xb06d100c),
		185: uint32(0xf6567769),
		186: uint32(0x3c1bdec6),
		187: uint32(0x7a20b9a3),
		188: uint32(0xe986fd42),
		189: uint32(0xafbd9a27),
		190: uint32(0x65f03388),
		191: uint32(0x23cb54ed),
		192: uint32(0x3ae0095e),
		193: uint32(0x7cdb6e3b),
		194: uint32(0xb696c794),
		195: uint32(0xf0ada0f1),
		196: uint32(0x630be410),
		197: uint32(0x25308375),
		198: uint32(0xef7d2ada),
		199: uint32(0xa9464dbf),
		200: uint32(0x8836d2c3),
		201: uint32(0xce0db5a6),
		202: uint32(0x04401c09),
		203: uint32(0x427b7b6c),
		204: uint32(0xd1dd3f8d),
		205: uint32(0x97e658e8),
		206: uint32(0x5dabf147),
		207: uint32(0x1b909622),
		208: uint32(0x1f4bcfbe),
		209: uint32(0x5970a8db),
		210: uint32(0x933d0174),
		211: uint32(0xd5066611),
		212: uint32(0x46a022f0),
		213: uint32(0x009b4595),
		214: uint32(0xcad6ec3a),
		215: uint32(0x8ced8b5f),
		216: uint32(0xad9d1423),
		217: uint32(0xeba67346),
		218: uint32(0x21ebdae9),
		219: uint32(0x67d0bd8c),
		220: uint32(0xf476f96d),
		221: uint32(0xb24d9e08),
		222: uint32(0x780037a7),
		223: uint32(0x3e3b50c2),
		224: uint32(0x31b0f544),
		225: uint32(0x778b9221),
		226: uint32(0xbdc63b8e),
		227: uint32(0xfbfd5ceb),
		228: uint32(0x685b180a),
		229: uint32(0x2e607f6f),
		230: uint32(0xe42dd6c0),
		231: uint32(0xa216b1a5),
		232: uint32(0x83662ed9),
		233: uint32(0xc55d49bc),
		234: uint32(0x0f10e013),
		235: uint32(0x492b8776),
		236: uint32(0xda8dc397),
		237: uint32(0x9cb6a4f2),
		238: uint32(0x56fb0d5d),
		239: uint32(0x10c06a38),
		240: uint32(0x141b33a4),
		241: uint32(0x522054c1),
		242: uint32(0x986dfd6e),
		243: uint32(0xde569a0b),
		244: uint32(0x4df0deea),
		245: uint32(0x0bcbb98f),
		246: uint32(0xc1861020),
		247: uint32(0x87bd7745),
		248: uint32(0xa6cde839),
		249: uint32(0xe0f68f5c),
		250: uint32(0x2abb26f3),
		251: uint32(0x6c804196),
		252: uint32(0xff260577),
		253: uint32(0xb91d6212),
		254: uint32(0x7350cbbd),
		255: uint32(0x356bacd8),
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
		if !(k < int32(4)) {
			break
		}
		data = data>>libc.Int32FromInt32(8) ^ _crc_table[data&uint32(0xff)]
		goto _1
	_1:
		;
		k++
	}
	return data
}

func _crc_word_big(tls *libc.TLS, data Tz_word_t) (r Tz_word_t) {
	var k int32
	_ = k
	k = 0
	for {
		if !(k < int32(4)) {
			break
		}
		data = data<<libc.Int32FromInt32(8) ^ _crc_big_table[data>>((libc.Int32FromInt32(4)-libc.Int32FromInt32(1))<<libc.Int32FromInt32(3))&uint32(0xff)]
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
func Xcrc32_z(tls *libc.TLS, crc uint32, buf uintptr, len1 Tz_size_t) (r uint32) {
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
		return uint32(0)
	}
	/* Pre-condition the CRC */
	crc = ^crc & uint32(0xffffffff)
	/* If provided enough bytes, do a braided CRC calculation. */
	if len1 >= libc.Uint32FromInt32(libc.Int32FromInt32(m_N)*libc.Int32FromInt32(4)+libc.Int32FromInt32(4)-libc.Int32FromInt32(1)) {
		/* Compute the CRC up to a z_word_t boundary. */
		for len1 != 0 && uint32(buf)&libc.Uint32FromInt32(libc.Int32FromInt32(4)-libc.Int32FromInt32(1)) != uint32(0) {
			len1--
			v1 = buf
			buf++
			crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v1))))&uint32(0xff)]
		}
		/* Compute the CRC on as many N z_word_t blocks as are available. */
		blks = len1 / libc.Uint32FromInt32(libc.Int32FromInt32(m_N)*libc.Int32FromInt32(4))
		len1 -= blks * uint32(m_N) * uint32(4)
		words = buf
		/* Do endian check at execution time instead of compile time, since ARM
		   processors can change the endianness at execution time. If the
		   compiler knows what the endianness will be, it can optimize out the
		   check and the unused branch. */
		*(*uint32)(unsafe.Pointer(bp)) = uint32(1)
		if *(*uint8)(unsafe.Pointer(bp)) != 0 {
			/* Initialize the CRC for each braid. */
			crc0 = crc
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
				word0 = crc0 ^ *(*Tz_word_t)(unsafe.Pointer(words))
				word1 = crc1 ^ *(*Tz_word_t)(unsafe.Pointer(words + 1*4))
				word2 = crc2 ^ *(*Tz_word_t)(unsafe.Pointer(words + 2*4))
				word3 = crc3 ^ *(*Tz_word_t)(unsafe.Pointer(words + 3*4))
				word4 = crc4 ^ *(*Tz_word_t)(unsafe.Pointer(words + 4*4))
				words += uintptr(m_N) * 4
				/* Compute and update the CRC for each word. The loop should
				   get unrolled. */
				crc0 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word0&uint32(0xff))*4))
				crc1 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word1&uint32(0xff))*4))
				crc2 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word2&uint32(0xff))*4))
				crc3 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word3&uint32(0xff))*4))
				crc4 = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(word4&uint32(0xff))*4))
				k = int32(1)
				for {
					if !(k < int32(4)) {
						break
					}
					crc0 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word0>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc1 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word1>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc2 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word2>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc3 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word3>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc4 ^= *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_table)) + uintptr(k)*1024 + uintptr(word4>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
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
			crc = _crc_word(tls, crc0^*(*Tz_word_t)(unsafe.Pointer(words)))
			crc = _crc_word(tls, crc1^*(*Tz_word_t)(unsafe.Pointer(words + 1*4))^crc)
			crc = _crc_word(tls, crc2^*(*Tz_word_t)(unsafe.Pointer(words + 2*4))^crc)
			crc = _crc_word(tls, crc3^*(*Tz_word_t)(unsafe.Pointer(words + 3*4))^crc)
			crc = _crc_word(tls, crc4^*(*Tz_word_t)(unsafe.Pointer(words + 4*4))^crc)
			words += uintptr(m_N) * 4
		} else {
			/* Initialize the CRC for each braid. */
			crc01 = _byte_swap(tls, crc)
			crc11 = uint32(0)
			crc21 = uint32(0)
			crc31 = uint32(0)
			crc41 = uint32(0)
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
				word11 = crc11 ^ *(*Tz_word_t)(unsafe.Pointer(words + 1*4))
				word21 = crc21 ^ *(*Tz_word_t)(unsafe.Pointer(words + 2*4))
				word31 = crc31 ^ *(*Tz_word_t)(unsafe.Pointer(words + 3*4))
				word41 = crc41 ^ *(*Tz_word_t)(unsafe.Pointer(words + 4*4))
				words += uintptr(m_N) * 4
				/* Compute and update the CRC for each word. The loop should
				   get unrolled. */
				crc01 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word01&uint32(0xff))*4))
				crc11 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word11&uint32(0xff))*4))
				crc21 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word21&uint32(0xff))*4))
				crc31 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word31&uint32(0xff))*4))
				crc41 = *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(word41&uint32(0xff))*4))
				k = int32(1)
				for {
					if !(k < int32(4)) {
						break
					}
					crc01 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*1024 + uintptr(word01>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc11 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*1024 + uintptr(word11>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc21 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*1024 + uintptr(word21>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc31 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*1024 + uintptr(word31>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
					crc41 ^= *(*Tz_word_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_braid_big_table)) + uintptr(k)*1024 + uintptr(word41>>(k<<libc.Int32FromInt32(3))&uint32(0xff))*4))
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
			comb = _crc_word_big(tls, crc11^*(*Tz_word_t)(unsafe.Pointer(words + 1*4))^comb)
			comb = _crc_word_big(tls, crc21^*(*Tz_word_t)(unsafe.Pointer(words + 2*4))^comb)
			comb = _crc_word_big(tls, crc31^*(*Tz_word_t)(unsafe.Pointer(words + 3*4))^comb)
			comb = _crc_word_big(tls, crc41^*(*Tz_word_t)(unsafe.Pointer(words + 4*4))^comb)
			words += uintptr(m_N) * 4
			crc = _byte_swap(tls, comb)
		}
		/*
		   Update the pointer to the remaining bytes to process.
		*/
		buf = words
	}
	/* Complete the computation of the CRC on any remaining bytes. */
	for len1 >= uint32(8) {
		len1 -= uint32(8)
		v6 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v6))))&uint32(0xff)]
		v7 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v7))))&uint32(0xff)]
		v8 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v8))))&uint32(0xff)]
		v9 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v9))))&uint32(0xff)]
		v10 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v10))))&uint32(0xff)]
		v11 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v11))))&uint32(0xff)]
		v12 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v12))))&uint32(0xff)]
		v13 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v13))))&uint32(0xff)]
	}
	for len1 != 0 {
		len1--
		v14 = buf
		buf++
		crc = crc>>int32(8) ^ _crc_table[(crc^uint32(*(*uint8)(unsafe.Pointer(v14))))&uint32(0xff)]
	}
	/* Return the CRC, post-conditioned. */
	return crc ^ uint32(0xffffffff)
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32(tls *libc.TLS, crc uint32, buf uintptr, len1 TuInt) (r uint32) {
	return Xcrc32_z(tls, crc, buf, len1)
}

// C documentation
//
//	/* ========================================================================= */
func Xcrc32_combine64(tls *libc.TLS, crc1 TuLong, crc2 TuLong, len2 Toff_t) (r TuLong) {
	return _multmodp(tls, _x2nmodp(tls, len2, uint32(3)), crc1) ^ crc2&uint32(0xffffffff)
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
	return _x2nmodp(tls, len2, uint32(3))
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
	return _multmodp(tls, op, crc1) ^ crc2&uint32(0xffffffff)
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
	Fblock_start      int32
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
	Fblock_start      int32
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
	*(*uintptr)(unsafe.Add(p, 20)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 44)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 68)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 116)) = __ccgo_fp(_deflate_slow)
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
	*(*TuInt)(unsafe.Pointer(strm + 4)) -= len1
	libc.Xmemcpy(tls, buf, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, len1)
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap == int32(1) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xadler32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, buf, len1)
	} else {
		if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap == int32(2) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, buf, len1)
		}
	}
	*(*uintptr)(unsafe.Pointer(strm)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 8)) += len1
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
		more = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead - (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		/* Deal with !@#$% 64K limit: */
		if uint32(4) <= uint32(2) {
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
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr(wsize), wsize-more)
			*(*TuInt)(unsafe.Pointer(s + 112)) -= wsize
			*(*TuInt)(unsafe.Pointer(s + 108)) -= wsize /* we now have strstart >= MAX_DIST */
			*(*int32)(unsafe.Pointer(s + 92)) -= libc.Int32FromUint32(wsize)
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
		*(*TuInt)(unsafe.Pointer(s + 116)) += n
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
		curr = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart + (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < curr {
			/* Previous high water mark below current data -- zero WIN_INIT
			 * bytes or up to end of window, whichever is less.
			 */
			init1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - curr
			if init1 > uint32(m_MAX_MATCH) {
				init1 = uint32(m_MAX_MATCH)
			}
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr(curr), 0, init1)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = curr + init1
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < curr+uint32(m_MAX_MATCH) {
				/* High water mark at or above current data, but below current data
				 * plus WIN_INIT -- zero out to current data plus WIN_INIT, or up
				 * to end of window, whichever is less.
				 */
				init1 = curr + uint32(m_MAX_MATCH) - (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water
				if init1 > (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size-(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water {
					init1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water
				}
				libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water), 0, init1)
				*(*Tulg)(unsafe.Pointer(s + 5824)) += init1
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
	if version == uintptr(m_Z_NULL) || int32(*(*int8)(unsafe.Pointer(version))) != int32(_my_version[0]) || libc.Uint32FromInt32(stream_size) != uint32(56) {
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
	s = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(5828))
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
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size, libc.Uint32FromInt32(2)*libc.Uint32FromInt64(1))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhead = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = uint32(0)                                                /* nothing written to s->window yet */
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
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size = (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize * uint32(4)
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

var _my_version = [6]int8{'1', '.', '3', '.', '1'}

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
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2))
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
	*(*TuInt)(unsafe.Pointer(s + 108)) += (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
		libc.Xmemcpy(tls, dictionary, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead)-uintptr(len1), len1)
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
	var v3 uint32
	_, _, _, _ = s, v1, v2, v3
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	v1 = libc.Uint32FromInt32(0)
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* use zfree if we ever allocate msg dynamically */
	(*Tz_stream)(unsafe.Pointer(strm)).Fdata_type = int32(m_Z_UNKNOWN)
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending = uint32(0)
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
		v3 = Xcrc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
	} else {
		v3 = Xadler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
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
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size = libc.Uint32FromInt32(2) * (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
	libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2))
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
		*(*uint32)(unsafe.Pointer(pending)) = (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending
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
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(value&(libc.Int32FromInt32(1)<<put-libc.Int32FromInt32(1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)))
		*(*int32)(unsafe.Pointer(s + 5820)) += put
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
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 || (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start)+(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead != 0 {
			return -int32(5)
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel != level {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches != uint32(0) {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches == uint32(1) {
				_slide_hash(tls, s)
			} else {
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
				libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2))
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
	var v1, v7 uint32
	var v2 int32
	_, _, _, _, _, _, _, _, _, _ = fixedlen, s, storelen, str, wraplen, v1, v2, v3, v5, v7
	/* upper bound for fixed blocks with 9-bit literals and length 255
	   (memLevel == 2, which is the lowest that may not use stored blocks) --
	   ~13% overhead plus a small constant */
	fixedlen = sourceLen + sourceLen>>libc.Int32FromInt32(3) + sourceLen>>libc.Int32FromInt32(8) + sourceLen>>libc.Int32FromInt32(9) + uint32(4)
	/* upper bound for stored blocks with length 127 (memLevel == 1) --
	   ~4% overhead plus a small constant */
	storelen = sourceLen + sourceLen>>libc.Int32FromInt32(5) + sourceLen>>libc.Int32FromInt32(7) + sourceLen>>libc.Int32FromInt32(11) + uint32(7)
	/* if can't get parameters, return larger bound plus a zlib wrapper */
	if _deflateStateCheck(tls, strm) != 0 {
		if fixedlen > storelen {
			v1 = fixedlen
		} else {
			v1 = storelen
		}
		return v1 + uint32(6)
	}
	/* compute wrapper length */
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	switch (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap {
	case 0: /* raw deflate */
		wraplen = uint32(0)
	case int32(1): /* zlib wrapper */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != 0 {
			v2 = int32(4)
		} else {
			v2 = 0
		}
		wraplen = libc.Uint32FromInt32(int32(6) + v2)
	case int32(2): /* gzip wrapper */
		wraplen = uint32(18)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead != uintptr(m_Z_NULL) {
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
				wraplen += uint32(2) + (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len
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
				wraplen += uint32(2)
			}
		}
	default: /* for compiler happiness */
		wraplen = uint32(6)
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
	return sourceLen + sourceLen>>libc.Int32FromInt32(12) + sourceLen>>libc.Int32FromInt32(14) + sourceLen>>libc.Int32FromInt32(25) + uint32(13) - uint32(6) + wraplen
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
	v2 = s + 20
	v1 = *(*Tulg)(unsafe.Pointer(v2))
	*(*Tulg)(unsafe.Pointer(v2))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = uint8(b >> libc.Int32FromInt32(8))
	v4 = s + 20
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
	len1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
	if len1 > (*Tz_stream)(unsafe.Pointer(strm)).Favail_out {
		len1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	}
	if len1 == uint32(0) {
		return
	}
	libc.Xmemcpy(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out, len1)
	*(*uintptr)(unsafe.Pointer(strm + 12)) += uintptr(len1)
	*(*uintptr)(unsafe.Pointer(s + 16)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 20)) += len1
	*(*TuInt)(unsafe.Pointer(strm + 16)) -= len1
	*(*Tulg)(unsafe.Pointer(s + 20)) -= len1
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == uint32(0) {
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
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
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
			_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler&libc.Uint32FromInt32(0xffff)))
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xadler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
		/* Compression must start with an empty pending buffer */
		_flush_pending(tls, strm)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_GZIP_STATE) {
		/* gzip header */
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromInt32(31))
		v7 = s + 20
		v6 = *(*Tulg)(unsafe.Pointer(v7))
		*(*Tulg)(unsafe.Pointer(v7))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v6))) = libc.Uint8FromInt32(libc.Int32FromInt32(139))
		v9 = s + 20
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromInt32(8))
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead == uintptr(m_Z_NULL) {
			v11 = s + 20
			v10 = *(*Tulg)(unsafe.Pointer(v11))
			*(*Tulg)(unsafe.Pointer(v11))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v13 = s + 20
			v12 = *(*Tulg)(unsafe.Pointer(v13))
			*(*Tulg)(unsafe.Pointer(v13))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v12))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v15 = s + 20
			v14 = *(*Tulg)(unsafe.Pointer(v15))
			*(*Tulg)(unsafe.Pointer(v15))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v17 = s + 20
			v16 = *(*Tulg)(unsafe.Pointer(v17))
			*(*Tulg)(unsafe.Pointer(v17))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v16))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v19 = s + 20
			v18 = *(*Tulg)(unsafe.Pointer(v19))
			*(*Tulg)(unsafe.Pointer(v19))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v18))) = libc.Uint8FromInt32(libc.Int32FromInt32(0))
			v21 = s + 20
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
			v25 = s + 20
			v24 = *(*Tulg)(unsafe.Pointer(v25))
			*(*Tulg)(unsafe.Pointer(v25))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = libc.Uint8FromInt32(libc.Int32FromInt32(m_OS_CODE))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
			/* Compression must start with an empty pending buffer */
			_flush_pending(tls, strm)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
				return m_Z_OK
			}
		} else {
			v27 = s + 20
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
			v34 = s + 20
			v33 = *(*Tulg)(unsafe.Pointer(v34))
			*(*Tulg)(unsafe.Pointer(v34))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v33))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime & libc.Uint32FromInt32(0xff))
			v36 = s + 20
			v35 = *(*Tulg)(unsafe.Pointer(v36))
			*(*Tulg)(unsafe.Pointer(v36))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v35))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			v38 = s + 20
			v37 = *(*Tulg)(unsafe.Pointer(v38))
			*(*Tulg)(unsafe.Pointer(v38))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v37))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
			v40 = s + 20
			v39 = *(*Tulg)(unsafe.Pointer(v40))
			*(*Tulg)(unsafe.Pointer(v40))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v39))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
			v42 = s + 20
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
			v46 = s + 20
			v45 = *(*Tulg)(unsafe.Pointer(v46))
			*(*Tulg)(unsafe.Pointer(v46))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v45))) = libc.Uint8FromInt32((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fos & libc.Int32FromInt32(0xff))
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
				v48 = s + 20
				v47 = *(*Tulg)(unsafe.Pointer(v48))
				*(*Tulg)(unsafe.Pointer(v48))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v47))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len & libc.Uint32FromInt32(0xff))
				v50 = s + 20
				v49 = *(*Tulg)(unsafe.Pointer(v50))
				*(*Tulg)(unsafe.Pointer(v50))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v49))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending)
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint32(0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_EXTRA_STATE)
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_EXTRA_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
			beg = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending /* start of bytes to update crc */
			left = uint32((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len&libc.Uint32FromInt32(0xffff)) - (*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex
			for (*Tdeflate_state)(unsafe.Pointer(s)).Fpending+left > (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
				copy1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
				libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex), copy1)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fpending = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size
				if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg {
					(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg)
				}
				*(*Tulg)(unsafe.Pointer(s + 32)) += copy1
				_flush_pending(tls, strm)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
					return m_Z_OK
				}
				beg = uint32(0)
				left -= copy1
			}
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex), left)
			*(*Tulg)(unsafe.Pointer(s + 20)) += left
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg)
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint32(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_NAME_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_NAME_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname != uintptr(m_Z_NULL) {
			beg1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
			for cond := true; cond; cond = val != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
					if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg1 {
						(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg1), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg1)
					}
					_flush_pending(tls, strm)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
						(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
						return m_Z_OK
					}
					beg1 = uint32(0)
				}
				v52 = s + 32
				v51 = *(*Tulg)(unsafe.Pointer(v52))
				*(*Tulg)(unsafe.Pointer(v52))++
				val = libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname + uintptr(v51))))
				v54 = s + 20
				v53 = *(*Tulg)(unsafe.Pointer(v54))
				*(*Tulg)(unsafe.Pointer(v54))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v53))) = libc.Uint8FromInt32(val)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg1 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg1), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg1)
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint32(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_COMMENT_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_COMMENT_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment != uintptr(m_Z_NULL) {
			beg2 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
			for cond := true; cond; cond = val1 != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
					if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg2 {
						(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg2), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg2)
					}
					_flush_pending(tls, strm)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
						(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
						return m_Z_OK
					}
					beg2 = uint32(0)
				}
				v56 = s + 32
				v55 = *(*Tulg)(unsafe.Pointer(v56))
				*(*Tulg)(unsafe.Pointer(v56))++
				val1 = libc.Int32FromUint8(*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment + uintptr(v55))))
				v58 = s + 20
				v57 = *(*Tulg)(unsafe.Pointer(v58))
				*(*Tulg)(unsafe.Pointer(v58))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v57))) = libc.Uint8FromInt32(val1)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg2 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg2), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg2)
			}
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_HCRC_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_HCRC_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending+uint32(2) > (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
				_flush_pending(tls, strm)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
					return m_Z_OK
				}
			}
			v60 = s + 20
			v59 = *(*Tulg)(unsafe.Pointer(v60))
			*(*Tulg)(unsafe.Pointer(v60))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v59))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler & libc.Uint32FromInt32(0xff))
			v62 = s + 20
			v61 = *(*Tulg)(unsafe.Pointer(v62))
			*(*Tulg)(unsafe.Pointer(v62))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v61))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			(*Tz_stream)(unsafe.Pointer(strm)).Fadler = Xcrc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
		/* Compression must start with an empty pending buffer */
		_flush_pending(tls, strm)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
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
					x__tr_stored_block(tls, s, libc.UintptrFromInt32(0), uint32(0), 0)
					/* For a full flush, this empty block will be recognized
					 * as a special marker by inflate_sync().
					 */
					if flush == int32(m_Z_FULL_FLUSH) {
						*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
						libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2)) /* forget history */
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
		v67 = s + 20
		v66 = *(*Tulg)(unsafe.Pointer(v67))
		*(*Tulg)(unsafe.Pointer(v67))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v66))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler & libc.Uint32FromInt32(0xff))
		v69 = s + 20
		v68 = *(*Tulg)(unsafe.Pointer(v69))
		*(*Tulg)(unsafe.Pointer(v69))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v68))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
		v71 = s + 20
		v70 = *(*Tulg)(unsafe.Pointer(v71))
		*(*Tulg)(unsafe.Pointer(v71))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v70))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
		v73 = s + 20
		v72 = *(*Tulg)(unsafe.Pointer(v73))
		*(*Tulg)(unsafe.Pointer(v73))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v72))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
		v75 = s + 20
		v74 = *(*Tulg)(unsafe.Pointer(v75))
		*(*Tulg)(unsafe.Pointer(v75))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v74))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in & libc.Uint32FromInt32(0xff))
		v77 = s + 20
		v76 = *(*Tulg)(unsafe.Pointer(v77))
		*(*Tulg)(unsafe.Pointer(v77))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v76))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
		v79 = s + 20
		v78 = *(*Tulg)(unsafe.Pointer(v79))
		*(*Tulg)(unsafe.Pointer(v79))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v78))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
		v81 = s + 20
		v80 = *(*Tulg)(unsafe.Pointer(v81))
		*(*Tulg)(unsafe.Pointer(v81))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v80))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
	} else {
		_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler>>libc.Int32FromInt32(16)))
		_putShortMSB(tls, s, uint32((*Tz_stream)(unsafe.Pointer(strm)).Fadler&libc.Uint32FromInt32(0xffff)))
	}
	_flush_pending(tls, strm)
	/* If avail_out is zero, the application will call deflate again
	 * to flush the rest.
	 */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap > 0 {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = -(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap
	} /* write the trailer only once! */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
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
	libc.Xmemcpy(tls, dest, source, uint32(56))
	ds = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(5828))
	if ds == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(dest)).Fstate = ds
	libc.Xmemcpy(tls, ds, ss, uint32(5828))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fstrm = dest
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size, libc.Uint32FromInt32(2)*libc.Uint32FromInt64(1))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fprev = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fhead = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhash_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize, libc.Uint32FromInt32(libc.Int32FromInt32(m_LIT_BUFS)))
	if (*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fprev == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fhead == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf == uintptr(m_Z_NULL) {
		XdeflateEnd(tls, dest)
		return -int32(4)
	}
	/* following zmemcpy do not work for 16-bit MSDOS */
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(ss)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size*uint32(2)*uint32(1))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fprev, (*Tdeflate_state)(unsafe.Pointer(ss)).Fprev, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size*uint32(2))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhead, (*Tdeflate_state)(unsafe.Pointer(ss)).Fhead, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhash_size*uint32(2))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize*uint32(m_LIT_BUFS))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf + uintptr(int32((*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_out)-int32((*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_buf))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fsym_buf = (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize)
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fl_desc.Fdyn_tree = ds + 148
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fd_desc.Fdyn_tree = ds + 2440
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fbl_desc.Fdyn_tree = ds + 2684
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
		len1 = int32(m_MAX_MATCH) - (int32(strend) - int32(scan))
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
	var have, last, left, len1, min_block, used, v1, v4, v6, v7, v8, v9 uint32
	var v10, v12, v2 int32
	var p11, p3, p5 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = have, last, left, len1, min_block, used, v1, v10, v12, v2, v4, v6, v7, v8, v9, p11, p3, p5
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size-uint32(5) > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	} else {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - uint32(5)
	}
	/* Smallest worthy block size when not flushing or finishing. By default
	 * this is 32K. This can be as small as 507 bytes for memLevel == 1. For
	 * large input and output buffers, the stored block size will be larger.
	 */
	min_block = v1
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
		left = uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start)) /* bytes left in window */
		if len1 > left+(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in {
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
		x__tr_stored_block(tls, s, libc.UintptrFromInt32(0), uint32(0), libc.Int32FromUint32(last))
		/* Replace the lengths in the dummy stored block with len. */
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(4)))) = uint8(len1)
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(3)))) = uint8(len1 >> int32(8))
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(2)))) = uint8(^len1)
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(1)))) = uint8(^len1 >> int32(8))
		/* Write the stored block header bytes. */
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		/* Copy uncompressed bytes from the window to next_out. */
		if left != 0 {
			if left > len1 {
				left = len1
			}
			libc.Xmemcpy(tls, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_out, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), left)
			*(*uintptr)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 12)) += uintptr(left)
			*(*TuInt)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 16)) -= left
			*(*TuLong)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 20)) += left
			p3 = s + 92
			*(*int32)(unsafe.Pointer(p3)) = int32(uint32(*(*int32)(unsafe.Pointer(p3))) + left)
			len1 -= left
		}
		/* Copy uncompressed bytes directly from next_in to next_out, updating
		 * the check value.
		 */
		if len1 != 0 {
			_read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_out, len1)
			*(*uintptr)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 12)) += uintptr(len1)
			*(*TuInt)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 16)) -= len1
			*(*TuLong)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 20)) += len1
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
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_in-uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
			(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size-(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart <= used {
				/* Slide the window down. */
				*(*TuInt)(unsafe.Pointer(s + 108)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
				libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches < uint32(2) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
				} /* add a pending slide_hash() */
				if (*Tdeflate_state)(unsafe.Pointer(s)).Finsert > (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
					(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
				}
			}
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart), (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_in-uintptr(used), used)
			*(*TuInt)(unsafe.Pointer(s + 108)) += used
			if used > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-(*Tdeflate_state)(unsafe.Pointer(s)).Finsert {
				v4 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
			} else {
				v4 = used
			}
			*(*TuInt)(unsafe.Pointer(s + 5812)) += v4
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	}
	/* If the last block was written to next_out, then done. */
	if last != 0 {
		return int32(_finish_done)
	}
	/* If flushing and all input has been consumed, then done. */
	if flush != m_Z_NO_FLUSH && flush != int32(m_Z_FINISH) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) == (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start {
		return int32(_block_done)
	}
	/* Fill the window with any remaining input. */
	have = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in > have && (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size) {
		/* Slide the window down. */
		p5 = s + 92
		*(*int32)(unsafe.Pointer(p5)) = int32(uint32(*(*int32)(unsafe.Pointer(p5))) - (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
		*(*TuInt)(unsafe.Pointer(s + 108)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
		libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
		*(*TuInt)(unsafe.Pointer(s + 108)) += have
		if have > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-(*Tdeflate_state)(unsafe.Pointer(s)).Finsert {
			v6 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
		} else {
			v6 = have
		}
		*(*TuInt)(unsafe.Pointer(s + 5812)) += v6
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	}
	/* There was not enough avail_out to write a complete worthy or flushed
	 * stored block to next_out. Write a stored block to pending instead, if we
	 * have enough input for a worthy block, or if flushing and there is enough
	 * room for the remaining input as a stored block in the pending buffer.
	 */
	have = libc.Uint32FromInt32(((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid + int32(42)) >> int32(3)) /* number of header bytes */
	/* maximum stored block length that will fit in pending: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size-have > libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_STORED)) {
		v7 = libc.Uint32FromInt32(libc.Int32FromInt32(m_MAX_STORED))
	} else {
		v7 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - have
	}
	have = v7
	if have > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		v8 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	} else {
		v8 = have
	}
	min_block = v8
	left = uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
	if left >= min_block || (left != 0 || flush == int32(m_Z_FINISH)) && flush != m_Z_NO_FLUSH && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && left <= have {
		if left > have {
			v9 = have
		} else {
			v9 = left
		}
		len1 = v9
		if flush == int32(m_Z_FINISH) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && len1 == left {
			v10 = int32(1)
		} else {
			v10 = 0
		}
		last = libc.Uint32FromInt32(v10)
		x__tr_stored_block(tls, s, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), len1, libc.Int32FromUint32(last))
		p11 = s + 92
		*(*int32)(unsafe.Pointer(p11)) = int32(uint32(*(*int32)(unsafe.Pointer(p11))) + len1)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
	}
	/* We've done all we can with the available input and output. */
	if last != 0 {
		v12 = int32(_finish_started)
	} else {
		v12 = int32(_need_more)
	}
	return v12
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
			v4 = s + 5792
			v3 = *(*TuInt)(unsafe.Pointer(v4))
			*(*TuInt)(unsafe.Pointer(v4))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))) = uint8(dist)
			v6 = s + 5792
			v5 = *(*TuInt)(unsafe.Pointer(v6))
			*(*TuInt)(unsafe.Pointer(v6))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16(dist) >> libc.Int32FromInt32(8))
			v8 = s + 5792
			v7 = *(*TuInt)(unsafe.Pointer(v8))
			*(*TuInt)(unsafe.Pointer(v8))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v7))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(libc.Int32FromUint8(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if libc.Int32FromUint16(dist) < int32(256) {
				v9 = libc.Int32FromUint8(x__dist_code[dist])
			} else {
				v9 = libc.Int32FromUint8(x__dist_code[int32(256)+libc.Int32FromUint16(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v9)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			*(*TuInt)(unsafe.Pointer(s + 116)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
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
					v11 = s + 96
					*(*TuInt)(unsafe.Pointer(v11))--
					v10 = *(*TuInt)(unsafe.Pointer(v11))
					if !(v10 != uint32(0)) {
						break
					}
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
			} else {
				*(*TuInt)(unsafe.Pointer(s + 108)) += (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
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
			v15 = s + 5792
			v14 = *(*TuInt)(unsafe.Pointer(v15))
			*(*TuInt)(unsafe.Pointer(v15))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v14))) = uint8(0)
			v17 = s + 5792
			v16 = *(*TuInt)(unsafe.Pointer(v17))
			*(*TuInt)(unsafe.Pointer(v17))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v16))) = uint8(0)
			v19 = s + 5792
			v18 = *(*TuInt)(unsafe.Pointer(v19))
			*(*TuInt)(unsafe.Pointer(v19))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v18))) = cc
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		}
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v20 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v20 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v20, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
			v22 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v22 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v22, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v23 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v23 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v23, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
			v4 = s + 5792
			v3 = *(*TuInt)(unsafe.Pointer(v4))
			*(*TuInt)(unsafe.Pointer(v4))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))) = uint8(dist)
			v6 = s + 5792
			v5 = *(*TuInt)(unsafe.Pointer(v6))
			*(*TuInt)(unsafe.Pointer(v6))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16(dist) >> libc.Int32FromInt32(8))
			v8 = s + 5792
			v7 = *(*TuInt)(unsafe.Pointer(v8))
			*(*TuInt)(unsafe.Pointer(v8))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v7))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(libc.Int32FromUint8(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if libc.Int32FromUint16(dist) < int32(256) {
				v9 = libc.Int32FromUint8(x__dist_code[dist])
			} else {
				v9 = libc.Int32FromUint8(x__dist_code[int32(256)+libc.Int32FromUint16(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v9)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			/* Insert in hash table all strings up to the end of the match.
			 * strstart - 1 and strstart are already inserted. If there is not
			 * enough lookahead, the last two strings are not inserted in
			 * the hash table.
			 */
			*(*TuInt)(unsafe.Pointer(s + 116)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length - uint32(1)
			*(*TuInt)(unsafe.Pointer(s + 120)) -= uint32(2)
			for {
				v14 = s + 108
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
				v11 = s + 120
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
					v16 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
				} else {
					v16 = libc.UintptrFromInt32(m_Z_NULL)
				}
				x__tr_flush_block(tls, s, v16, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
				v18 = s + 5792
				v17 = *(*TuInt)(unsafe.Pointer(v18))
				*(*TuInt)(unsafe.Pointer(v18))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v17))) = uint8(0)
				v20 = s + 5792
				v19 = *(*TuInt)(unsafe.Pointer(v20))
				*(*TuInt)(unsafe.Pointer(v20))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v19))) = uint8(0)
				v22 = s + 5792
				v21 = *(*TuInt)(unsafe.Pointer(v22))
				*(*TuInt)(unsafe.Pointer(v22))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v21))) = cc
				*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
				bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
				if bflush != 0 {
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
						v23 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
					} else {
						v23 = libc.UintptrFromInt32(m_Z_NULL)
					}
					x__tr_flush_block(tls, s, v23, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
					(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
		v25 = s + 5792
		v24 = *(*TuInt)(unsafe.Pointer(v25))
		*(*TuInt)(unsafe.Pointer(v25))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v24))) = uint8(0)
		v27 = s + 5792
		v26 = *(*TuInt)(unsafe.Pointer(v27))
		*(*TuInt)(unsafe.Pointer(v27))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v26))) = uint8(0)
		v29 = s + 5792
		v28 = *(*TuInt)(unsafe.Pointer(v29))
		*(*TuInt)(unsafe.Pointer(v29))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v28))) = cc1
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc1)*4))++
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
			v31 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v31 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v31, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v32 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v32 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v32, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(m_MAX_MATCH) - libc.Uint32FromInt32(int32(strend)-int32(scan))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length > (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
				}
			}
		}
		/* Emit match if have run of MIN_MATCH or longer, else emit literal */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length >= uint32(m_MIN_MATCH) {
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = libc.Uint16FromInt32(libc.Int32FromInt32(1))
			v24 = s + 5792
			v23 = *(*TuInt)(unsafe.Pointer(v24))
			*(*TuInt)(unsafe.Pointer(v24))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v23))) = uint8(dist)
			v26 = s + 5792
			v25 = *(*TuInt)(unsafe.Pointer(v26))
			*(*TuInt)(unsafe.Pointer(v26))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v25))) = libc.Uint8FromInt32(libc.Int32FromUint16(dist) >> libc.Int32FromInt32(8))
			v28 = s + 5792
			v27 = *(*TuInt)(unsafe.Pointer(v28))
			*(*TuInt)(unsafe.Pointer(v28))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v27))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(libc.Int32FromUint8(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if libc.Int32FromUint16(dist) < int32(256) {
				v29 = libc.Int32FromUint8(x__dist_code[dist])
			} else {
				v29 = libc.Int32FromUint8(x__dist_code[int32(256)+libc.Int32FromUint16(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v29)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			*(*TuInt)(unsafe.Pointer(s + 116)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			*(*TuInt)(unsafe.Pointer(s + 108)) += (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		} else {
			/* No match, output a literal byte */
			cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
			v31 = s + 5792
			v30 = *(*TuInt)(unsafe.Pointer(v31))
			*(*TuInt)(unsafe.Pointer(v31))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v30))) = uint8(0)
			v33 = s + 5792
			v32 = *(*TuInt)(unsafe.Pointer(v33))
			*(*TuInt)(unsafe.Pointer(v33))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v32))) = uint8(0)
			v35 = s + 5792
			v34 = *(*TuInt)(unsafe.Pointer(v35))
			*(*TuInt)(unsafe.Pointer(v35))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v34))) = cc
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
			(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		}
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v36 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v36 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v36, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
			v37 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v37 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v37, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v38 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v38 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v38, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
		v3 = s + 5792
		v2 = *(*TuInt)(unsafe.Pointer(v3))
		*(*TuInt)(unsafe.Pointer(v3))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v2))) = uint8(0)
		v5 = s + 5792
		v4 = *(*TuInt)(unsafe.Pointer(v5))
		*(*TuInt)(unsafe.Pointer(v5))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v4))) = uint8(0)
		v7 = s + 5792
		v6 = *(*TuInt)(unsafe.Pointer(v7))
		*(*TuInt)(unsafe.Pointer(v7))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v6))) = cc
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
		bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next == (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_end)
		(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v8 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v8 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v8, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
			v9 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v9 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v9, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fsym_next != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v10 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(libc.Uint32FromInt32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v10 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v10, libc.Uint32FromInt32(libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = libc.Int32FromUint32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
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
	Fcheck    uint32
	Ftotal    uint32
	Fhead     Tgz_headerp
	Fwbits    uint32
	Fwsize    uint32
	Fwhave    uint32
	Fwnext    uint32
	Fwindow   uintptr
	Fhold     uint32
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
	if version == uintptr(m_Z_NULL) || int32(*(*int8)(unsafe.Pointer(version))) != int32(*(*int8)(unsafe.Pointer(__ccgo_ts))) || stream_size != libc.Int32FromUint32(libc.Uint32FromInt64(56)) {
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
	state = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(7120))
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
	var bits, copy1, have, hold, left, len1, v1, v18, v20, v24, v29, v30, v43, v44 uint32
	var from, put, state, v11, v15, v16, v17, v19, v21, v23, v25, v26, v27, v28, v31, v33, v35, v36, v37, v39, v41, v42, v46, v47 uintptr
	var here, last Tcode
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
	hold = uint32(0)
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
			hold >>= bits & uint32(7)
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v11))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flast = libc.Int32FromUint32(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(1) - libc.Uint32FromInt32(1)))
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(1))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(1))
		switch hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2) - libc.Uint32FromInt32(1)) {
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
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(2))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
		goto _10
	_4:
		;
		/* get and verify stored block length */
	_14:
		;
		hold >>= bits & uint32(7)
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v15))) << bits
			bits += uint32(8)
		}
		if hold&uint32(0xffff) != hold>>int32(16)^uint32(0xffff) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 25
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = hold & uint32(0xffff)
		hold = uint32(0)
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
			libc.Xmemcpy(tls, put, *(*uintptr)(unsafe.Pointer(bp)), copy1)
			have -= copy1
			*(*uintptr)(unsafe.Pointer(bp)) += uintptr(copy1)
			left -= copy1
			put += uintptr(copy1)
			*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v16))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnlen = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(257)
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fndist = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(1)
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fncode = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(4)
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(4))
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v17))) << bits
				bits += uint32(8)
			}
			v19 = state + 108
			v18 = *(*uint32)(unsafe.Pointer(v19))
			*(*uint32)(unsafe.Pointer(v19))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order[v18])*2)) = uint16(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3) - libc.Uint32FromInt32(1)))
			hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(3))
			bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
		}
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < uint32(19) {
			v21 = state + 108
			v20 = *(*uint32)(unsafe.Pointer(v21))
			*(*uint32)(unsafe.Pointer(v21))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order[v20])*2)) = uint16(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(7)
		ret = Xinflate_table(tls, int32(_CODES), state+116, uint32(19), state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 90
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* get length and distance code code lengths */
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v23))) << bits
				bits += uint32(8)
				goto _22
			_22:
			}
			if libc.Int32FromUint16(here.Fval) < int32(16) {
				hold >>= uint32(here.Fbits)
				bits -= uint32(here.Fbits)
				v25 = state + 108
				v24 = *(*uint32)(unsafe.Pointer(v25))
				*(*uint32)(unsafe.Pointer(v25))++
				*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v24)*2)) = here.Fval
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
						hold += uint32(*(*uint8)(unsafe.Pointer(v26))) << bits
						bits += uint32(8)
					}
					hold >>= uint32(here.Fbits)
					bits -= uint32(here.Fbits)
					if (*Tinflate_state)(unsafe.Pointer(state)).Fhave == uint32(0) {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 115
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
					len1 = uint32(*(*uint16)(unsafe.Pointer(state + 116 + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fhave-uint32(1))*2)))
					copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2)-libc.Uint32FromInt32(1))
					hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(2))
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
							hold += uint32(*(*uint8)(unsafe.Pointer(v27))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(3))
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
							hold += uint32(*(*uint8)(unsafe.Pointer(v28))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(11) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(7)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(7))
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
					v31 = state + 108
					v30 = *(*uint32)(unsafe.Pointer(v31))
					*(*uint32)(unsafe.Pointer(v31))++
					*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v30)*2)) = uint16(len1)
				}
			}
		}
		/* handle error breaks in while */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_BAD) {
			goto _10
		}
		/* check for end-of-block code (better have one) */
		if libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(state + 116 + 256*2))) == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 141
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* build code tables -- note: do not change the lenbits or distbits
		   values here (9 and 6) without reading the comments in inftrees.h
		   concerning the ENOUGH constants, which depend on those values */
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
		ret = Xinflate_table(tls, int32(_LENS), state+116, (*Tinflate_state)(unsafe.Pointer(state)).Fnlen, state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 178
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(6)
		ret = Xinflate_table(tls, int32(_DISTS), state+116+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fnlen)*2, (*Tinflate_state)(unsafe.Pointer(state)).Fndist, state+112, state+92, state+756)
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
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v33))) << bits
			bits += uint32(8)
			goto _32
		_32:
		}
		if here.Fop != 0 && libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v35))) << bits
				bits += uint32(8)
				goto _34
			_34:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v37))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 68)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
		}
		/* get distance code */
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits-libc.Uint32FromInt32(1)))*4))
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v39))) << bits
			bits += uint32(8)
			goto _38
		_38:
		}
		if libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v41))) << bits
				bits += uint32(8)
				goto _40
			_40:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v42))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 72)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
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
			*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
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
	var bits, dist, dmask, hold, len1, lmask, op, whave, wnext, wsize, v13, v17, v21, v9 uint32
	var v45, v46 int32
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v1))) << bits
			bits += uint32(8)
			v2 = in
			in++
			hold += uint32(*(*uint8)(unsafe.Pointer(v2))) << bits
			bits += uint32(8)
		}
		here = lcode + uintptr(hold&lmask)*4
		goto dolen
	dolen:
		;
		op = uint32((*Tcode)(unsafe.Pointer(here)).Fbits)
		hold >>= op
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
						hold += uint32(*(*uint8)(unsafe.Pointer(v4))) << bits
						bits += uint32(8)
					}
					len1 += hold & (uint32(1)<<op - uint32(1))
					hold >>= op
					bits -= op
				}
				if bits < uint32(15) {
					v5 = in
					in++
					hold += uint32(*(*uint8)(unsafe.Pointer(v5))) << bits
					bits += uint32(8)
					v6 = in
					in++
					hold += uint32(*(*uint8)(unsafe.Pointer(v6))) << bits
					bits += uint32(8)
				}
				here = dcode + uintptr(hold&dmask)*4
				goto dodist
			dodist:
				;
				op = uint32((*Tcode)(unsafe.Pointer(here)).Fbits)
				hold >>= op
				bits -= op
				op = uint32((*Tcode)(unsafe.Pointer(here)).Fop)
				if op&uint32(16) != 0 { /* distance base */
					dist = uint32((*Tcode)(unsafe.Pointer(here)).Fval)
					op &= uint32(15) /* number of extra bits */
					if bits < op {
						v7 = in
						in++
						hold += uint32(*(*uint8)(unsafe.Pointer(v7))) << bits
						bits += uint32(8)
						if bits < op {
							v8 = in
							in++
							hold += uint32(*(*uint8)(unsafe.Pointer(v8))) << bits
							bits += uint32(8)
						}
					}
					dist += hold & (uint32(1)<<op - uint32(1))
					hold >>= op
					bits -= op
					op = libc.Uint32FromInt32(int32(out) - int32(beg)) /* max distance in output */
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
						here = dcode + uintptr((*Tcode)(unsafe.Pointer(here)).Fval)*4 + uintptr(hold&uint32(libc.Uint32FromUint32(1)<<op-libc.Uint32FromInt32(1)))*4
						goto dodist
					} else {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 256
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
				}
			} else {
				if op&uint32(64) == uint32(0) { /* 2nd level length code */
					here = lcode + uintptr((*Tcode)(unsafe.Pointer(here)).Fval)*4 + uintptr(hold&uint32(libc.Uint32FromUint32(1)<<op-libc.Uint32FromInt32(1)))*4
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
	hold &= uint32(uint32(1)<<bits - uint32(1))
	/* update state and return */
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = in
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = out
	if in < last {
		v45 = int32(5) + (int32(last) - int32(in))
	} else {
		v45 = int32(5) - (int32(in) - int32(last))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = libc.Uint32FromInt32(v45)
	if out < end {
		v46 = int32(257) + (int32(end) - int32(out))
	} else {
		v46 = int32(257) - (int32(out) - int32(end))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = libc.Uint32FromInt32(v46)
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
	var v2 uint32
	_, _, _, _, _ = state, v1, v2, v3, v4
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	v2 = libc.Uint32FromInt32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Ftotal = v2
	v1 = v2
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 { /* to support ill-conceived Java test suite */
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = libc.Uint32FromInt32((*Tinflate_state)(unsafe.Pointer(state)).Fwrap & int32(1))
	}
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HEAD)
	(*Tinflate_state)(unsafe.Pointer(state)).Flast = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fhavedict = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fflags = -int32(1)
	(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(32768)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhead = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = uint32(0)
	v4 = state + 1332
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
	if version == uintptr(m_Z_NULL) || int32(*(*int8)(unsafe.Pointer(version))) != int32(*(*int8)(unsafe.Pointer(__ccgo_ts))) || stream_size != libc.Int32FromUint32(libc.Uint32FromInt64(56)) {
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
	state = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(7120))
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
		(*Tinflate_state)(unsafe.Pointer(state)).Fhold = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fbits = uint32(0)
		return m_Z_OK
	}
	if bits > int32(16) || (*Tinflate_state)(unsafe.Pointer(state)).Fbits+libc.Uint32FromInt32(bits) > uint32(32) {
		return -int32(2)
	}
	value = int32(value & (libc.Int32FromInt32(1)<<bits - libc.Int32FromInt32(1)))
	*(*uint32)(unsafe.Pointer(state + 60)) += uint32(libc.Uint32FromInt32(value) << (*Tinflate_state)(unsafe.Pointer(state)).Fbits)
	*(*uint32)(unsafe.Pointer(state + 64)) += libc.Uint32FromInt32(bits)
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
		(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fwbits, libc.Uint32FromInt64(1))
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
		libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, end-uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwsize), (*Tinflate_state)(unsafe.Pointer(state)).Fwsize)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	} else {
		dist = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
		if dist > copy1 {
			dist = copy1
		}
		libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), end-uintptr(copy1), dist)
		copy1 -= dist
		if copy1 != 0 {
			libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, end-uintptr(copy1), copy1)
			(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = copy1
			(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
		} else {
			*(*uint32)(unsafe.Pointer(state + 52)) += dist
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwnext == (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				*(*uint32)(unsafe.Pointer(state + 48)) += dist
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
	var bits, copy1, have, hold, in, left, len1, out, v101, v102, v37, v46, v48, v49, v50, v52, v53, v56, v58, v59, v64, v66, v70, v75, v76, v90, v96, v97, v98 uint32
	var from, next, put, state, v100, v36, v42, v43, v44, v45, v51, v54, v55, v57, v60, v61, v62, v63, v65, v67, v69, v71, v72, v73, v74, v77, v79, v81, v82, v85, v87, v88, v92, v93, v94, v95, p83, p89 uintptr
	var here, last Tcode
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v36))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(2) != 0 && hold == uint32(0x8b1f) { /* gzip header */
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwbits == uint32(0) {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = uint32(15)
			}
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
			hold = uint32(0)
			bits = uint32(0)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_FLAGS)
			goto _35
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fdone = -int32(1)
		}
		if !((*Tinflate_state)(unsafe.Pointer(state)).Fwrap&libc.Int32FromInt32(1) != 0) || (uint32(hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(8)-libc.Uint32FromInt32(1))<<libc.Int32FromInt32(8))+hold>>int32(8))%uint32(31) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 308
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) != uint32(m_Z_DEFLATED) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 331
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(4))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(4))
		len1 = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(8)
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
		v37 = Xadler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v37
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v37
		if hold&uint32(0x200) != 0 {
			v38 = int32(_DICTID)
		} else {
			v38 = int32(_TYPE)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = v38
		hold = uint32(0)
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v42))) << bits
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
		(*Tinflate_state)(unsafe.Pointer(state)).Fflags = libc.Int32FromUint32(hold)
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
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Ftext = libc.Int32FromUint32(hold >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(1))
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
		}
		hold = uint32(0)
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v43))) << bits
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
		hold = uint32(0)
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v44))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fxflags = libc.Int32FromUint32(hold & libc.Uint32FromInt32(0xff))
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fos = libc.Int32FromUint32(hold >> libc.Int32FromInt32(8))
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
		}
		hold = uint32(0)
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v45))) << bits
				bits += uint32(8)
			}
			(*Tinflate_state)(unsafe.Pointer(state)).Flength = hold
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_len = hold
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
				(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
			}
			hold = uint32(0)
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
					libc.Xmemcpy(tls, (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra+uintptr(len1), next, v48)
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
					(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = Xcrc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
				}
				have -= copy1
				next += uintptr(copy1)
				*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
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
					v51 = state + 68
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
					v54 = state + 68
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v55))) << bits
				bits += uint32(8)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && hold != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck&uint32(0xffff) {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 403
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint32(0)
			bits = uint32(0)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fhcrc = (*Tinflate_state)(unsafe.Pointer(state)).Fflags >> libc.Int32FromInt32(9) & libc.Int32FromInt32(1)
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fdone = int32(1)
		}
		v56 = Xcrc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v57))) << bits
			bits += uint32(8)
		}
		v58 = hold>>libc.Int32FromInt32(24)&libc.Uint32FromInt32(0xff) + hold>>libc.Int32FromInt32(8)&libc.Uint32FromInt32(0xff00) + hold&libc.Uint32FromInt32(0xff00)<<libc.Int32FromInt32(8) + hold&libc.Uint32FromInt32(0xff)<<libc.Int32FromInt32(24)
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v58
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v58
		hold = uint32(0)
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
		v59 = Xadler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
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
			hold >>= bits & uint32(7)
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v60))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flast = libc.Int32FromUint32(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(1) - libc.Uint32FromInt32(1)))
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(1))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(1))
		switch hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2) - libc.Uint32FromInt32(1)) {
		case uint32(0): /* stored block */
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_STORED)
		case uint32(1): /* fixed block */
			_fixedtables1(tls, state)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN_) /* decode codes */
			if flush == int32(m_Z_TREES) {
				hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(2))
				bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
				goto inf_leave
			}
		case uint32(2): /* dynamic block */
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TABLE)
		case uint32(3):
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 6
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
		}
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(2))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(2))
		goto _35
	_15:
		;
		hold >>= bits & uint32(7)
		bits -= bits & uint32(7) /* go to byte boundary */
		for bits < libc.Uint32FromInt32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v61 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v61))) << bits
			bits += uint32(8)
		}
		if hold&uint32(0xffff) != hold>>int32(16)^uint32(0xffff) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 25
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = hold & uint32(0xffff)
		hold = uint32(0)
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
			libc.Xmemcpy(tls, put, next, copy1)
			have -= copy1
			next += uintptr(copy1)
			left -= copy1
			put += uintptr(copy1)
			*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
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
			hold += uint32(*(*uint8)(unsafe.Pointer(v62))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnlen = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(257)
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fndist = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(1)
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		bits -= libc.Uint32FromInt32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fncode = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(4)
		hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(4))
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v63))) << bits
				bits += uint32(8)
			}
			v65 = state + 108
			v64 = *(*uint32)(unsafe.Pointer(v65))
			*(*uint32)(unsafe.Pointer(v65))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order1[v64])*2)) = uint16(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3) - libc.Uint32FromInt32(1)))
			hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(3))
			bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
		}
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < uint32(19) {
			v67 = state + 108
			v66 = *(*uint32)(unsafe.Pointer(v67))
			*(*uint32)(unsafe.Pointer(v67))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order1[v66])*2)) = uint16(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(7)
		ret = Xinflate_table(tls, int32(_CODES), state+116, uint32(19), state+112, state+88, state+756)
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
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
				if uint32(here.Fbits) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v69 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v69))) << bits
				bits += uint32(8)
				goto _68
			_68:
			}
			if libc.Int32FromUint16(here.Fval) < int32(16) {
				hold >>= uint32(here.Fbits)
				bits -= uint32(here.Fbits)
				v71 = state + 108
				v70 = *(*uint32)(unsafe.Pointer(v71))
				*(*uint32)(unsafe.Pointer(v71))++
				*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v70)*2)) = here.Fval
			} else {
				if libc.Int32FromUint16(here.Fval) == int32(16) {
					for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(2)) {
						if have == uint32(0) {
							goto inf_leave
						}
						have--
						v72 = next
						next++
						hold += uint32(*(*uint8)(unsafe.Pointer(v72))) << bits
						bits += uint32(8)
					}
					hold >>= uint32(here.Fbits)
					bits -= uint32(here.Fbits)
					if (*Tinflate_state)(unsafe.Pointer(state)).Fhave == uint32(0) {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 115
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
					len1 = uint32(*(*uint16)(unsafe.Pointer(state + 116 + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fhave-uint32(1))*2)))
					copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2)-libc.Uint32FromInt32(1))
					hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(2))
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
							hold += uint32(*(*uint8)(unsafe.Pointer(v73))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(3))
						bits -= libc.Uint32FromInt32(libc.Int32FromInt32(3))
					} else {
						for bits < libc.Uint32FromInt32(libc.Int32FromUint8(here.Fbits)+libc.Int32FromInt32(7)) {
							if have == uint32(0) {
								goto inf_leave
							}
							have--
							v74 = next
							next++
							hold += uint32(*(*uint8)(unsafe.Pointer(v74))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(11) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(7)-libc.Uint32FromInt32(1))
						hold >>= libc.Uint32FromInt32(libc.Int32FromInt32(7))
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
					v77 = state + 108
					v76 = *(*uint32)(unsafe.Pointer(v77))
					*(*uint32)(unsafe.Pointer(v77))++
					*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v76)*2)) = uint16(len1)
				}
			}
		}
		/* handle error breaks in while */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_BAD) {
			goto _35
		}
		/* check for end-of-block code (better have one) */
		if libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(state + 116 + 256*2))) == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 141
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		/* build code tables -- note: do not change the lenbits or distbits
		   values here (9 and 6) without reading the comments in inftrees.h
		   concerning the ENOUGH constants, which depend on those values */
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
		ret = Xinflate_table(tls, int32(_LENS), state+116, (*Tinflate_state)(unsafe.Pointer(state)).Fnlen, state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 178
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(6)
		ret = Xinflate_table(tls, int32(_DISTS), state+116+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fnlen)*2, (*Tinflate_state)(unsafe.Pointer(state)).Fndist, state+112, state+92, state+756)
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
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v79 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v79))) << bits
			bits += uint32(8)
			goto _78
		_78:
		}
		if here.Fop != 0 && libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
				if libc.Uint32FromInt32(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v81 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v81))) << bits
				bits += uint32(8)
				goto _80
			_80:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
			*(*int32)(unsafe.Pointer(state + 7112)) += libc.Int32FromUint8(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
		bits -= uint32(here.Fbits)
		*(*int32)(unsafe.Pointer(state + 7112)) += libc.Int32FromUint8(here.Fbits)
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v82))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 68)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			p83 = state + 7112
			*(*int32)(unsafe.Pointer(p83)) = int32(uint32(*(*int32)(unsafe.Pointer(p83))) + (*Tinflate_state)(unsafe.Pointer(state)).Fextra)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fwas = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DIST)
		/* fallthrough */
	_24:
		;
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v85 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v85))) << bits
			bits += uint32(8)
			goto _84
		_84:
		}
		if libc.Int32FromUint8(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(last.Fop))-uint32(1))>>last.Fbits)*4))
				if libc.Uint32FromInt32(libc.Int32FromUint8(last.Fbits)+libc.Int32FromUint8(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v87 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v87))) << bits
				bits += uint32(8)
				goto _86
			_86:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
			*(*int32)(unsafe.Pointer(state + 7112)) += libc.Int32FromUint8(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
		bits -= uint32(here.Fbits)
		*(*int32)(unsafe.Pointer(state + 7112)) += libc.Int32FromUint8(here.Fbits)
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v88))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 72)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			p89 = state + 7112
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
		*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v95))) << bits
				bits += uint32(8)
			}
			out -= left
			*(*TuLong)(unsafe.Pointer(strm + 20)) += out
			*(*uint32)(unsafe.Pointer(state + 32)) += out
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
					v98 = hold>>int32(24)&uint32(0xff) + hold>>int32(8)&uint32(0xff00) + hold&uint32(0xff00)<<int32(8) + hold&uint32(0xff)<<int32(24)
				}
			}
			if v99 && v98 != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 423
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint32(0)
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
				hold += uint32(*(*uint8)(unsafe.Pointer(v100))) << bits
				bits += uint32(8)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && hold != (*Tinflate_state)(unsafe.Pointer(state)).Ftotal&uint32(0xffffffff) {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 444
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint32(0)
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
	*(*TuLong)(unsafe.Pointer(strm + 8)) += in
	*(*TuLong)(unsafe.Pointer(strm + 20)) += out
	*(*uint32)(unsafe.Pointer(state + 32)) += out
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
		libc.Xmemcpy(tls, dictionary, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), (*Tinflate_state)(unsafe.Pointer(state)).Fwhave-(*Tinflate_state)(unsafe.Pointer(state)).Fwnext)
		libc.Xmemcpy(tls, dictionary+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwhave)-uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, (*Tinflate_state)(unsafe.Pointer(state)).Fwnext)
	}
	if dictLength != uintptr(m_Z_NULL) {
		*(*TuInt)(unsafe.Pointer(dictLength)) = (*Tinflate_state)(unsafe.Pointer(state)).Fwhave
	}
	return m_Z_OK
}

func XinflateSetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength TuInt) (r int32) {
	var dictid uint32
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
		dictid = Xadler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
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
	var in, len1, out, v1 uint32
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
		*(*uint32)(unsafe.Pointer(state + 60)) >>= (*Tinflate_state)(unsafe.Pointer(state)).Fbits & uint32(7)
		*(*uint32)(unsafe.Pointer(state + 64)) -= (*Tinflate_state)(unsafe.Pointer(state)).Fbits & uint32(7)
		len1 = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fbits >= uint32(8) {
			v1 = len1
			len1++
			(*(*[4]uint8)(unsafe.Pointer(bp)))[v1] = uint8((*Tinflate_state)(unsafe.Pointer(state)).Fhold)
			*(*uint32)(unsafe.Pointer(state + 60)) >>= uint32(8)
			*(*uint32)(unsafe.Pointer(state + 64)) -= uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		_syncsearch(tls, state+108, bp, len1)
	}
	/* search available input */
	len1 = _syncsearch(tls, state+108, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, (*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
	*(*TuInt)(unsafe.Pointer(strm + 4)) -= len1
	*(*uintptr)(unsafe.Pointer(strm)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 8)) += len1
	/* return no joy or set up to restart inflate() on a new block */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fhave != uint32(4) {
		return -int32(3)
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fflags == -int32(1) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fwrap = 0
	} else {
		*(*int32)(unsafe.Pointer(state + 12)) &= ^libc.Int32FromInt32(4)
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
	copy1 = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, libc.Uint32FromInt32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(7120))
	if copy1 == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	window = uintptr(m_Z_NULL)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) {
		window = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fwbits, libc.Uint32FromInt64(1))
		if window == uintptr(m_Z_NULL) {
			(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, copy1)
			return -int32(4)
		}
	}
	/* copy state */
	libc.Xmemcpy(tls, dest, source, uint32(56))
	libc.Xmemcpy(tls, copy1, state, uint32(7120))
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fstrm = dest
	if (*Tinflate_state)(unsafe.Pointer(state)).Flencode >= state+1332 && (*Tinflate_state)(unsafe.Pointer(state)).Flencode <= state+1332+uintptr(libc.Int32FromInt32(m_ENOUGH_LENS)+libc.Int32FromInt32(m_ENOUGH_DISTS))*4-uintptr(1)*4 {
		(*Tinflate_state)(unsafe.Pointer(copy1)).Flencode = copy1 + 1332 + uintptr((int32((*Tinflate_state)(unsafe.Pointer(state)).Flencode)-t__predefined_ptrdiff_t(state+1332))/4)*4
		(*Tinflate_state)(unsafe.Pointer(copy1)).Fdistcode = copy1 + 1332 + uintptr((int32((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode)-t__predefined_ptrdiff_t(state+1332))/4)*4
	}
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fnext = copy1 + 1332 + uintptr((int32((*Tinflate_state)(unsafe.Pointer(state)).Fnext)-t__predefined_ptrdiff_t(state+1332))/4)*4
	if window != uintptr(m_Z_NULL) {
		wsize = uint32(1) << (*Tinflate_state)(unsafe.Pointer(state)).Fwbits
		libc.Xmemcpy(tls, window, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, wsize)
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
		*(*int32)(unsafe.Pointer(state + 12)) |= int32(4)
	} else {
		*(*int32)(unsafe.Pointer(state + 12)) &= ^libc.Int32FromInt32(4)
	}
	return m_Z_OK
}

func XinflateMark(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	var v1, v2 uint32
	_, _, _ = state, v1, v2
	if _inflateStateCheck(tls, strm) != 0 {
		return -(libc.Int32FromInt32(1) << libc.Int32FromInt32(16))
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
	return libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromUint32(libc.Uint32FromInt32((*Tinflate_state)(unsafe.Pointer(state)).Fback)<<libc.Int32FromInt32(16))) + v1)
}

func XinflateCodesUsed(tls *libc.TLS, strm Tz_streamp) (r uint32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	return libc.Uint32FromInt32((int32((*Tinflate_state)(unsafe.Pointer(state)).Fnext) - t__predefined_ptrdiff_t(state+1332)) / 4)
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
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fval = libc.Uint16FromInt32((int32(next) - int32(*(*uintptr)(unsafe.Pointer(table)))) / 4)
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
		v2 = s + 20
		v1 = *(*Tulg)(unsafe.Pointer(v2))
		*(*Tulg)(unsafe.Pointer(v2))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v4 = s + 20
		v3 = *(*Tulg)(unsafe.Pointer(v4))
		*(*Tulg)(unsafe.Pointer(v4))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid >= int32(8) {
			v6 = s + 20
			v5 = *(*Tulg)(unsafe.Pointer(v6))
			*(*Tulg)(unsafe.Pointer(v6))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf)
			p7 = s + 5816
			*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) >> libc.Int32FromInt32(8))
			*(*int32)(unsafe.Pointer(s + 5820)) -= int32(8)
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
		v2 = s + 20
		v1 = *(*Tulg)(unsafe.Pointer(v2))
		*(*Tulg)(unsafe.Pointer(v2))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v4 = s + 20
		v3 = *(*Tulg)(unsafe.Pointer(v4))
		*(*Tulg)(unsafe.Pointer(v4))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > 0 {
			v6 = s + 20
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
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(n)*4)) = uint16(0)
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
		*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(n)*4)) = uint16(0)
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
		*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(n)*4)) = uint16(0)
		goto _3
	_3:
		;
		n++
	}
	*(*Tush)(unsafe.Pointer(s + 148 + 256*4)) = uint16(1)
	v4 = libc.Uint32FromInt32(0)
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
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fdyn_tree = s + 148
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_l_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fdyn_tree = s + 2440
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_d_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbl_desc.Fdyn_tree = s + 2684
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
	v = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(k)*4))
	j = k << int32(1) /* left son of k */
	for j <= (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len {
		/* Set j to the smallest of the two sons: */
		if j < (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len && (libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j+int32(1))*4)))*4))) < libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j+int32(1))*4)))*4))) == libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) && libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j+int32(1))*4)))))) <= libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4))))))) {
			j++
		}
		/* Exit if v is smaller than both sons */
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(v)*4))) < libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(v)*4))) == libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) && libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(v)))) <= libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))))) {
			break
		}
		/* Exchange v with the smallest son */
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(k)*4)) = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4))
		k = j
		/* And continue down the tree, setting j to the left son of k */
		j <<= int32(1)
	}
	*(*int32)(unsafe.Pointer(s + 2908 + uintptr(k)*4)) = v
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
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2)) = uint16(0)
		goto _1
	_1:
		;
		bits++
	}
	/* In a first pass, compute the optimal bit lengths (which may
	 * overflow in the case of the bit length tree).
	 */
	*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max)*4)))*4 + 2)) = uint16(0) /* root of the heap */
	h = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max + int32(1)
	for {
		if !(h < libc.Int32FromInt32(2)*(libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES))+libc.Int32FromInt32(1)) {
			break
		}
		n = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(h)*4))
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
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2))++
		xbits = 0
		if n >= base {
			xbits = *(*Tintf)(unsafe.Pointer(extra + uintptr(n-base)*4))
		}
		f = *(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))
		*(*Tulg)(unsafe.Pointer(s + 5800)) += uint32(f) * uint32(libc.Uint32FromInt32(bits+xbits))
		if stree != 0 {
			*(*Tulg)(unsafe.Pointer(s + 5804)) += uint32(f) * uint32(libc.Uint32FromInt32(libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(stree + uintptr(n)*4 + 2)))+xbits))
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
		for libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2))) == 0 {
			bits--
		}
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2))-- /* move one leaf down the tree */
		p3 = s + 2876 + uintptr(bits+int32(1))*2
		*(*Tush)(unsafe.Pointer(p3)) = Tush(int32(*(*Tush)(unsafe.Pointer(p3))) + libc.Int32FromInt32(2)) /* move one overflow item as its brother */
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(max_length)*2))--
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
		n = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2)))
		for n != 0 {
			h--
			v5 = h
			m = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(v5)*4))
			if m > max_code {
				continue
			}
			if uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2))) != libc.Uint32FromInt32(bits) {
				*(*Tulg)(unsafe.Pointer(s + 5800)) += (libc.Uint32FromInt32(bits) - uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)))) * uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4)))
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
			v3 = s + 5200
			*(*int32)(unsafe.Pointer(v3))++
			v2 = *(*int32)(unsafe.Pointer(v3))
			v4 = n
			max_code = v4
			*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v2)*4)) = v4
			*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(n))) = uint8(0)
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
		v9 = s + 5200
		*(*int32)(unsafe.Pointer(v9))++
		v8 = *(*int32)(unsafe.Pointer(v9))
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v8)*4)) = v5
		node = v5
		*(*Tush)(unsafe.Pointer(tree + uintptr(node)*4)) = uint16(1)
		*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(node))) = uint8(0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len--
		if stree != 0 {
			*(*Tulg)(unsafe.Pointer(s + 5804)) -= uint32(*(*Tush)(unsafe.Pointer(stree + uintptr(node)*4 + 2)))
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
		n = *(*int32)(unsafe.Pointer(s + 2908 + 1*4))
		v12 = s + 5200
		v11 = *(*int32)(unsafe.Pointer(v12))
		*(*int32)(unsafe.Pointer(v12))--
		*(*int32)(unsafe.Pointer(s + 2908 + 1*4)) = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(v11)*4))
		_pqdownheap(tls, s, tree, int32(m_SMALLEST))  /* n = node of least frequency */
		m = *(*int32)(unsafe.Pointer(s + 2908 + 1*4)) /* m = node of next least frequency */
		v14 = s + 5204
		*(*int32)(unsafe.Pointer(v14))--
		v13 = *(*int32)(unsafe.Pointer(v14))
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v13)*4)) = n /* keep the nodes sorted by frequency */
		v16 = s + 5204
		*(*int32)(unsafe.Pointer(v16))--
		v15 = *(*int32)(unsafe.Pointer(v16))
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v15)*4)) = m
		/* Create a new node father of n and m */
		*(*Tush)(unsafe.Pointer(tree + uintptr(node)*4)) = libc.Uint16FromInt32(libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))) + libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4))))
		if libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(n)))) >= libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(m)))) {
			v17 = libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(n))))
		} else {
			v17 = libc.Int32FromUint8(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(m))))
		}
		*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(node))) = libc.Uint8FromInt32(v17 + libc.Int32FromInt32(1))
		v18 = libc.Uint16FromInt32(node)
		*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)) = v18
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = v18
		/* and insert the new node in the heap */
		v19 = node
		node++
		*(*int32)(unsafe.Pointer(s + 2908 + 1*4)) = v19
		_pqdownheap(tls, s, tree, int32(m_SMALLEST))
	}
	v21 = s + 5204
	*(*int32)(unsafe.Pointer(v21))--
	v20 = *(*int32)(unsafe.Pointer(v21))
	*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v20)*4)) = *(*int32)(unsafe.Pointer(s + 2908 + 1*4))
	/* At this point, the fields freq and dad are set. We can now
	 * generate the bit lengths.
	 */
	_gen_bitlen(tls, s, desc)
	/* The field len is now set, we can generate the bit codes */
	_gen_codes(tls, tree, max_code, s+2876)
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
				p3 = s + 2684 + uintptr(curlen)*4
				*(*Tush)(unsafe.Pointer(p3)) = Tush(int32(*(*Tush)(unsafe.Pointer(p3))) + count)
			} else {
				if curlen != 0 {
					if curlen != prevlen {
						*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4))++
					}
					*(*Tush)(unsafe.Pointer(s + 2684 + 16*4))++
				} else {
					if count <= int32(10) {
						*(*Tush)(unsafe.Pointer(s + 2684 + 17*4))++
					} else {
						*(*Tush)(unsafe.Pointer(s + 2684 + 18*4))++
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
					len1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4 + 2)))
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
						val = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))
						p5 = s + 5816
						*(*Tush)(unsafe.Pointer(p5)) = Tush(int32(*(*Tush)(unsafe.Pointer(p5))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v7 = s + 20
						v6 = *(*Tulg)(unsafe.Pointer(v7))
						*(*Tulg)(unsafe.Pointer(v7))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v6))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v9 = s + 20
						v8 = *(*Tulg)(unsafe.Pointer(v9))
						*(*Tulg)(unsafe.Pointer(v9))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
					} else {
						p10 = s + 5816
						*(*Tush)(unsafe.Pointer(p10)) = Tush(int32(*(*Tush)(unsafe.Pointer(p10))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len1
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
						len11 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
							val1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))
							p11 = s + 5816
							*(*Tush)(unsafe.Pointer(p11)) = Tush(int32(*(*Tush)(unsafe.Pointer(p11))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v13 = s + 20
							v12 = *(*Tulg)(unsafe.Pointer(v13))
							*(*Tulg)(unsafe.Pointer(v13))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v12))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v15 = s + 20
							v14 = *(*Tulg)(unsafe.Pointer(v15))
							*(*Tulg)(unsafe.Pointer(v15))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
						} else {
							p16 = s + 5816
							*(*Tush)(unsafe.Pointer(p16)) = Tush(int32(*(*Tush)(unsafe.Pointer(p16))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len11
						}
						count--
					}
					len2 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 16*4 + 2)))
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
						val2 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 16*4)))
						p17 = s + 5816
						*(*Tush)(unsafe.Pointer(p17)) = Tush(int32(*(*Tush)(unsafe.Pointer(p17))) | libc.Int32FromUint16(libc.Uint16FromInt32(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v19 = s + 20
						v18 = *(*Tulg)(unsafe.Pointer(v19))
						*(*Tulg)(unsafe.Pointer(v19))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v18))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v21 = s + 20
						v20 = *(*Tulg)(unsafe.Pointer(v21))
						*(*Tulg)(unsafe.Pointer(v21))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v20))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len2 - int32(m_Buf_size)
					} else {
						p22 = s + 5816
						*(*Tush)(unsafe.Pointer(p22)) = Tush(int32(*(*Tush)(unsafe.Pointer(p22))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 16*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len2
					}
					len3 = int32(2)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
						val3 = count - int32(3)
						p23 = s + 5816
						*(*Tush)(unsafe.Pointer(p23)) = Tush(int32(*(*Tush)(unsafe.Pointer(p23))) | libc.Int32FromUint16(libc.Uint16FromInt32(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v25 = s + 20
						v24 = *(*Tulg)(unsafe.Pointer(v25))
						*(*Tulg)(unsafe.Pointer(v25))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v27 = s + 20
						v26 = *(*Tulg)(unsafe.Pointer(v27))
						*(*Tulg)(unsafe.Pointer(v27))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len3 - int32(m_Buf_size)
					} else {
						p28 = s + 5816
						*(*Tush)(unsafe.Pointer(p28)) = Tush(int32(*(*Tush)(unsafe.Pointer(p28))) | libc.Int32FromUint16(libc.Uint16FromInt32(count-libc.Int32FromInt32(3)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len3
					}
				} else {
					if count <= int32(10) {
						len4 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 17*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len4 {
							val4 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 17*4)))
							p29 = s + 5816
							*(*Tush)(unsafe.Pointer(p29)) = Tush(int32(*(*Tush)(unsafe.Pointer(p29))) | libc.Int32FromUint16(libc.Uint16FromInt32(val4))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v31 = s + 20
							v30 = *(*Tulg)(unsafe.Pointer(v31))
							*(*Tulg)(unsafe.Pointer(v31))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v30))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v33 = s + 20
							v32 = *(*Tulg)(unsafe.Pointer(v33))
							*(*Tulg)(unsafe.Pointer(v33))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v32))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val4)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len4 - int32(m_Buf_size)
						} else {
							p34 = s + 5816
							*(*Tush)(unsafe.Pointer(p34)) = Tush(int32(*(*Tush)(unsafe.Pointer(p34))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 17*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len4
						}
						len5 = int32(3)
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len5 {
							val5 = count - int32(3)
							p35 = s + 5816
							*(*Tush)(unsafe.Pointer(p35)) = Tush(int32(*(*Tush)(unsafe.Pointer(p35))) | libc.Int32FromUint16(libc.Uint16FromInt32(val5))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v37 = s + 20
							v36 = *(*Tulg)(unsafe.Pointer(v37))
							*(*Tulg)(unsafe.Pointer(v37))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v36))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v39 = s + 20
							v38 = *(*Tulg)(unsafe.Pointer(v39))
							*(*Tulg)(unsafe.Pointer(v39))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v38))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val5)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len5 - int32(m_Buf_size)
						} else {
							p40 = s + 5816
							*(*Tush)(unsafe.Pointer(p40)) = Tush(int32(*(*Tush)(unsafe.Pointer(p40))) | libc.Int32FromUint16(libc.Uint16FromInt32(count-libc.Int32FromInt32(3)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len5
						}
					} else {
						len6 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 18*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len6 {
							val6 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 18*4)))
							p41 = s + 5816
							*(*Tush)(unsafe.Pointer(p41)) = Tush(int32(*(*Tush)(unsafe.Pointer(p41))) | libc.Int32FromUint16(libc.Uint16FromInt32(val6))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v43 = s + 20
							v42 = *(*Tulg)(unsafe.Pointer(v43))
							*(*Tulg)(unsafe.Pointer(v43))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v42))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v45 = s + 20
							v44 = *(*Tulg)(unsafe.Pointer(v45))
							*(*Tulg)(unsafe.Pointer(v45))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v44))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val6)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len6 - int32(m_Buf_size)
						} else {
							p46 = s + 5816
							*(*Tush)(unsafe.Pointer(p46)) = Tush(int32(*(*Tush)(unsafe.Pointer(p46))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + 18*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len6
						}
						len7 = int32(7)
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len7 {
							val7 = count - int32(11)
							p47 = s + 5816
							*(*Tush)(unsafe.Pointer(p47)) = Tush(int32(*(*Tush)(unsafe.Pointer(p47))) | libc.Int32FromUint16(libc.Uint16FromInt32(val7))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v49 = s + 20
							v48 = *(*Tulg)(unsafe.Pointer(v49))
							*(*Tulg)(unsafe.Pointer(v49))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v48))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v51 = s + 20
							v50 = *(*Tulg)(unsafe.Pointer(v51))
							*(*Tulg)(unsafe.Pointer(v51))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v50))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val7)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len7 - int32(m_Buf_size)
						} else {
							p52 = s + 5816
							*(*Tush)(unsafe.Pointer(p52)) = Tush(int32(*(*Tush)(unsafe.Pointer(p52))) | libc.Int32FromUint16(libc.Uint16FromInt32(count-libc.Int32FromInt32(11)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len7
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
	_scan_tree(tls, s, s+148, (*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fmax_code)
	_scan_tree(tls, s, s+2440, (*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fmax_code)
	/* Build the bit length tree: */
	_build_tree(tls, s, s+2864)
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
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(_bl_order[max_blindex])*4 + 2))) != 0 {
			break
		}
		goto _1
	_1:
		;
		max_blindex--
	}
	/* Update opt_len to include the bit length tree and counts */
	*(*Tulg)(unsafe.Pointer(s + 5800)) += uint32(3)*(libc.Uint32FromInt32(max_blindex)+uint32(1)) + uint32(5) + uint32(5) + uint32(4)
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
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 20
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5816
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | libc.Int32FromUint16(libc.Uint16FromInt32(lcodes-libc.Int32FromInt32(257)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len1
	} /* not +255 as stated in appnote.txt */
	len11 = int32(5)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
		val1 = dcodes - int32(1)
		p7 = s + 5816
		*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v9 = s + 20
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v11 = s + 20
		v10 = *(*Tulg)(unsafe.Pointer(v11))
		*(*Tulg)(unsafe.Pointer(v11))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
	} else {
		p12 = s + 5816
		*(*Tush)(unsafe.Pointer(p12)) = Tush(int32(*(*Tush)(unsafe.Pointer(p12))) | libc.Int32FromUint16(libc.Uint16FromInt32(dcodes-libc.Int32FromInt32(1)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len11
	}
	len2 = int32(4)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
		val2 = blcodes - int32(4)
		p13 = s + 5816
		*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | libc.Int32FromUint16(libc.Uint16FromInt32(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v15 = s + 20
		v14 = *(*Tulg)(unsafe.Pointer(v15))
		*(*Tulg)(unsafe.Pointer(v15))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v17 = s + 20
		v16 = *(*Tulg)(unsafe.Pointer(v17))
		*(*Tulg)(unsafe.Pointer(v17))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v16))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len2 - int32(m_Buf_size)
	} else {
		p18 = s + 5816
		*(*Tush)(unsafe.Pointer(p18)) = Tush(int32(*(*Tush)(unsafe.Pointer(p18))) | libc.Int32FromUint16(libc.Uint16FromInt32(blcodes-libc.Int32FromInt32(4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len2
	} /* not -3 as stated in appnote.txt */
	rank = 0
	for {
		if !(rank < blcodes) {
			break
		}
		len3 = int32(3)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
			val3 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(_bl_order[rank])*4 + 2)))
			p20 = s + 5816
			*(*Tush)(unsafe.Pointer(p20)) = Tush(int32(*(*Tush)(unsafe.Pointer(p20))) | libc.Int32FromUint16(libc.Uint16FromInt32(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
			v22 = s + 20
			v21 = *(*Tulg)(unsafe.Pointer(v22))
			*(*Tulg)(unsafe.Pointer(v22))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v21))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
			v24 = s + 20
			v23 = *(*Tulg)(unsafe.Pointer(v24))
			*(*Tulg)(unsafe.Pointer(v24))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v23))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
			*(*int32)(unsafe.Pointer(s + 5820)) += len3 - int32(m_Buf_size)
		} else {
			p25 = s + 5816
			*(*Tush)(unsafe.Pointer(p25)) = Tush(int32(*(*Tush)(unsafe.Pointer(p25))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(_bl_order[rank])*4 + 2)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
			*(*int32)(unsafe.Pointer(s + 5820)) += len3
		}
		goto _19
	_19:
		;
		rank++
	}
	_send_tree(tls, s, s+148, lcodes-int32(1))  /* literal tree */
	_send_tree(tls, s, s+2440, dcodes-int32(1)) /* distance tree */
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
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 20
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5816
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_STORED_BLOCK)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len1
	} /* send block type */
	_bi_windup(tls, s) /* align on byte boundary */
	v8 = s + 20
	v7 = *(*Tulg)(unsafe.Pointer(v8))
	*(*Tulg)(unsafe.Pointer(v8))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v7))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(stored_len)) & libc.Int32FromInt32(0xff))
	v10 = s + 20
	v9 = *(*Tulg)(unsafe.Pointer(v10))
	*(*Tulg)(unsafe.Pointer(v10))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(stored_len)) >> libc.Int32FromInt32(8))
	v12 = s + 20
	v11 = *(*Tulg)(unsafe.Pointer(v12))
	*(*Tulg)(unsafe.Pointer(v12))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(^stored_len)) & libc.Int32FromInt32(0xff))
	v14 = s + 20
	v13 = *(*Tulg)(unsafe.Pointer(v14))
	*(*Tulg)(unsafe.Pointer(v14))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v13))) = libc.Uint8FromInt32(libc.Int32FromUint16(uint16(^stored_len)) >> libc.Int32FromInt32(8))
	if stored_len != 0 {
		libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), buf, stored_len)
	}
	*(*Tulg)(unsafe.Pointer(s + 20)) += stored_len
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
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 20
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5816
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len1
	}
	len11 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4 + 2)))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
		val1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4)))
		p7 = s + 5816
		*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v9 = s + 20
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v11 = s + 20
		v10 = *(*Tulg)(unsafe.Pointer(v11))
		*(*Tulg)(unsafe.Pointer(v11))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
	} else {
		p12 = s + 5816
		*(*Tush)(unsafe.Pointer(p12)) = Tush(int32(*(*Tush)(unsafe.Pointer(p12))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len11
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
					p4 = s + 5816
					*(*Tush)(unsafe.Pointer(p4)) = Tush(int32(*(*Tush)(unsafe.Pointer(p4))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v6 = s + 20
					v5 = *(*Tulg)(unsafe.Pointer(v6))
					*(*Tulg)(unsafe.Pointer(v6))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v8 = s + 20
					v7 = *(*Tulg)(unsafe.Pointer(v8))
					*(*Tulg)(unsafe.Pointer(v8))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v7))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
				} else {
					p9 = s + 5816
					*(*Tush)(unsafe.Pointer(p9)) = Tush(int32(*(*Tush)(unsafe.Pointer(p9))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5820)) += len1
				} /* send a literal byte */
			} else {
				/* Here, lc is the match length - MIN_MATCH */
				code = uint32(x__length_code[lc])
				len11 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
					val1 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4)))
					p10 = s + 5816
					*(*Tush)(unsafe.Pointer(p10)) = Tush(int32(*(*Tush)(unsafe.Pointer(p10))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v12 = s + 20
					v11 = *(*Tulg)(unsafe.Pointer(v12))
					*(*Tulg)(unsafe.Pointer(v12))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v14 = s + 20
					v13 = *(*Tulg)(unsafe.Pointer(v14))
					*(*Tulg)(unsafe.Pointer(v14))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v13))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
				} else {
					p15 = s + 5816
					*(*Tush)(unsafe.Pointer(p15)) = Tush(int32(*(*Tush)(unsafe.Pointer(p15))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5820)) += len11
				} /* send length code */
				extra = _extra_lbits[code]
				if extra != 0 {
					lc -= _base_length[code]
					len2 = extra
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
						val2 = lc
						p16 = s + 5816
						*(*Tush)(unsafe.Pointer(p16)) = Tush(int32(*(*Tush)(unsafe.Pointer(p16))) | libc.Int32FromUint16(libc.Uint16FromInt32(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v18 = s + 20
						v17 = *(*Tulg)(unsafe.Pointer(v18))
						*(*Tulg)(unsafe.Pointer(v18))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v17))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v20 = s + 20
						v19 = *(*Tulg)(unsafe.Pointer(v20))
						*(*Tulg)(unsafe.Pointer(v20))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v19))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len2 - int32(m_Buf_size)
					} else {
						p21 = s + 5816
						*(*Tush)(unsafe.Pointer(p21)) = Tush(int32(*(*Tush)(unsafe.Pointer(p21))) | libc.Int32FromUint16(libc.Uint16FromInt32(lc))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len2
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
					p23 = s + 5816
					*(*Tush)(unsafe.Pointer(p23)) = Tush(int32(*(*Tush)(unsafe.Pointer(p23))) | libc.Int32FromUint16(libc.Uint16FromInt32(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v25 = s + 20
					v24 = *(*Tulg)(unsafe.Pointer(v25))
					*(*Tulg)(unsafe.Pointer(v25))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v27 = s + 20
					v26 = *(*Tulg)(unsafe.Pointer(v27))
					*(*Tulg)(unsafe.Pointer(v27))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5820)) += len3 - int32(m_Buf_size)
				} else {
					p28 = s + 5816
					*(*Tush)(unsafe.Pointer(p28)) = Tush(int32(*(*Tush)(unsafe.Pointer(p28))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5820)) += len3
				} /* send the distance code */
				extra = _extra_dbits[code]
				if extra != 0 {
					dist -= libc.Uint32FromInt32(_base_dist[code])
					len4 = extra
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len4 {
						val4 = libc.Int32FromUint32(dist)
						p29 = s + 5816
						*(*Tush)(unsafe.Pointer(p29)) = Tush(int32(*(*Tush)(unsafe.Pointer(p29))) | libc.Int32FromUint16(libc.Uint16FromInt32(val4))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v31 = s + 20
						v30 = *(*Tulg)(unsafe.Pointer(v31))
						*(*Tulg)(unsafe.Pointer(v31))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v30))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v33 = s + 20
						v32 = *(*Tulg)(unsafe.Pointer(v33))
						*(*Tulg)(unsafe.Pointer(v33))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v32))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val4)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len4 - int32(m_Buf_size)
					} else {
						p34 = s + 5816
						*(*Tush)(unsafe.Pointer(p34)) = Tush(int32(*(*Tush)(unsafe.Pointer(p34))) | libc.Int32FromUint16(uint16(dist))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len4
					} /* send the extra distance bits */
				}
			} /* literal or match pair ? */
			/* Check for no overlay of pending_buf on needed symbols */
		}
	}
	len5 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + 256*4 + 2)))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len5 {
		val5 = libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + 256*4)))
		p35 = s + 5816
		*(*Tush)(unsafe.Pointer(p35)) = Tush(int32(*(*Tush)(unsafe.Pointer(p35))) | libc.Int32FromUint16(libc.Uint16FromInt32(val5))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v37 = s + 20
		v36 = *(*Tulg)(unsafe.Pointer(v37))
		*(*Tulg)(unsafe.Pointer(v37))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v36))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v39 = s + 20
		v38 = *(*Tulg)(unsafe.Pointer(v39))
		*(*Tulg)(unsafe.Pointer(v39))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v38))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val5)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len5 - int32(m_Buf_size)
	} else {
		p40 = s + 5816
		*(*Tush)(unsafe.Pointer(p40)) = Tush(int32(*(*Tush)(unsafe.Pointer(p40))) | libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(ltree + 256*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len5
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
	var block_mask uint32
	var n int32
	_, _ = block_mask, n
	/* block_mask is the bit mask of block-listed bytes
	 * set bits 0..6, 14..25, and 28..31
	 * 0xf3ffc07f = binary 11110011111111111100000001111111
	 */
	block_mask = uint32(0xf3ffc07f)
	/* Check for non-textual ("block-listed") bytes. */
	n = 0
	for {
		if !(n <= int32(31)) {
			break
		}
		if block_mask&uint32(1) != 0 && libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 148 + uintptr(n)*4))) != 0 {
			return m_Z_BINARY
		}
		goto _1
	_1:
		;
		n++
		block_mask >>= uint32(1)
	}
	/* Check for textual ("allow-listed") bytes. */
	if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 148 + 9*4))) != 0 || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 148 + 10*4))) != 0 || libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 148 + 13*4))) != 0 {
		return int32(m_Z_TEXT)
	}
	n = int32(32)
	for {
		if !(n < int32(m_LITERALS)) {
			break
		}
		if libc.Int32FromUint16(*(*Tush)(unsafe.Pointer(s + 148 + uintptr(n)*4))) != 0 {
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
		_build_tree(tls, s, s+2840)
		_build_tree(tls, s, s+2852)
		/* At this point, opt_len and static_len are the total bit lengths of
		 * the compressed block data, excluding the tree representations.
		 */
		/* Build the bit length tree for the above two trees, and get the index
		 * in bl_order of the last bit length code to send.
		 */
		max_blindex = _build_bl_tree(tls, s)
		/* Determine the best encoding. Compute the block lengths in bytes. */
		opt_lenb = ((*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len + uint32(3) + uint32(7)) >> int32(3)
		static_lenb = ((*Tdeflate_state)(unsafe.Pointer(s)).Fstatic_len + uint32(3) + uint32(7)) >> int32(3)
		if static_lenb <= opt_lenb || (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_FIXED) {
			opt_lenb = static_lenb
		}
	} else {
		v1 = stored_len + libc.Uint32FromInt32(5)
		static_lenb = v1
		opt_lenb = v1 /* force a stored block */
	}
	if stored_len+uint32(4) <= opt_lenb && buf != libc.UintptrFromInt32(0) {
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
				p2 = s + 5816
				*(*Tush)(unsafe.Pointer(p2)) = Tush(int32(*(*Tush)(unsafe.Pointer(p2))) | libc.Int32FromUint16(libc.Uint16FromInt32(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				v4 = s + 20
				v3 = *(*Tulg)(unsafe.Pointer(v4))
				*(*Tulg)(unsafe.Pointer(v4))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
				v6 = s + 20
				v5 = *(*Tulg)(unsafe.Pointer(v6))
				*(*Tulg)(unsafe.Pointer(v6))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
				*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
			} else {
				p7 = s + 5816
				*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				*(*int32)(unsafe.Pointer(s + 5820)) += len1
			}
			_compress_block(tls, s, uintptr(unsafe.Pointer(&_static_ltree)), uintptr(unsafe.Pointer(&_static_dtree)))
		} else {
			len11 = int32(3)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
				val1 = libc.Int32FromInt32(m_DYN_TREES)<<libc.Int32FromInt32(1) + last
				p8 = s + 5816
				*(*Tush)(unsafe.Pointer(p8)) = Tush(int32(*(*Tush)(unsafe.Pointer(p8))) | libc.Int32FromUint16(libc.Uint16FromInt32(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				v10 = s + 20
				v9 = *(*Tulg)(unsafe.Pointer(v10))
				*(*Tulg)(unsafe.Pointer(v10))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
				v12 = s + 20
				v11 = *(*Tulg)(unsafe.Pointer(v12))
				*(*Tulg)(unsafe.Pointer(v12))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = libc.Uint8FromInt32(libc.Int32FromUint16((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
				*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
			} else {
				p13 = s + 5816
				*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | libc.Int32FromUint16(libc.Uint16FromInt32(libc.Int32FromInt32(m_DYN_TREES)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				*(*int32)(unsafe.Pointer(s + 5820)) += len11
			}
			_send_all_trees(tls, s, (*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fmax_code+int32(1), (*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fmax_code+int32(1), max_blindex+int32(1))
			_compress_block(tls, s, s+148, s+2440)
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
	v2 = s + 5792
	v1 = *(*TuInt)(unsafe.Pointer(v2))
	*(*TuInt)(unsafe.Pointer(v2))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v1))) = uint8(dist)
	v4 = s + 5792
	v3 = *(*TuInt)(unsafe.Pointer(v4))
	*(*TuInt)(unsafe.Pointer(v4))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v3))) = uint8(dist >> libc.Int32FromInt32(8))
	v6 = s + 5792
	v5 = *(*TuInt)(unsafe.Pointer(v6))
	*(*TuInt)(unsafe.Pointer(v6))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fsym_buf + uintptr(v5))) = uint8(lc)
	if dist == uint32(0) {
		/* lc is the unmatched char */
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(lc)*4))++
	} else {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
		/* Here, lc is the match length - MIN_MATCH */
		dist-- /* dist = match distance - 1 */
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(libc.Int32FromUint8(x__length_code[lc])+int32(m_LITERALS)+int32(1))*4))++
		if dist < uint32(256) {
			v7 = libc.Int32FromUint8(x__dist_code[dist])
		} else {
			v7 = libc.Int32FromUint8(x__dist_code[uint32(256)+dist>>int32(7)])
		}
		*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v7)*4))++
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
const m_F_GETLK = 12
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
const m_F_SETLK = 13
const m_F_SETLK64 = "F_SETLK"
const m_F_SETLKW = 14
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
	flags = uint32(0)
	switch libc.Int32FromUint32(libc.Uint32FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += uint32(1)
	case int32(8):
		flags += uint32(2)
	default:
		flags += uint32(3)
	}
	switch libc.Int32FromUint32(libc.Uint32FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(1) << libc.Int32FromInt32(2))
	case int32(8):
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(2) << libc.Int32FromInt32(2))
	default:
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(3) << libc.Int32FromInt32(2))
	}
	switch libc.Int32FromUint32(libc.Uint32FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(1) << libc.Int32FromInt32(4))
	case int32(8):
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(2) << libc.Int32FromInt32(4))
	default:
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(3) << libc.Int32FromInt32(4))
	}
	switch libc.Int32FromUint32(libc.Uint32FromInt64(8)) {
	case int32(2):
	case int32(4):
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(1) << libc.Int32FromInt32(6))
	case int32(8):
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(2) << libc.Int32FromInt32(6))
	default:
		flags += libc.Uint32FromInt32(libc.Int32FromInt32(3) << libc.Int32FromInt32(6))
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
	return libc.Xmalloc(tls, items*size)
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
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var err, v3, v4 int32
	var left TuLong
	var max TuInt
	var v1, v2 uint32
	var _ /* stream at bp+0 */ Tz_stream
	_, _, _, _, _, _, _ = err, left, max, v1, v2, v3, v4
	max = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	left = *(*TuLongf)(unsafe.Pointer(destLen))
	*(*TuLongf)(unsafe.Pointer(destLen)) = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = XdeflateInit_(tls, bp, level, __ccgo_ts, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		return err
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = dest
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = source
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	for cond := true; cond; cond = err == m_Z_OK {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out == uint32(0) {
			if left > max {
				v1 = max
			} else {
				v1 = left
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
			left -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out
		}
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in == uint32(0) {
			if sourceLen > max {
				v2 = max
			} else {
				v2 = sourceLen
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v2
			sourceLen -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in
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
	return sourceLen + sourceLen>>libc.Int32FromInt32(12) + sourceLen>>libc.Int32FromInt32(14) + sourceLen>>libc.Int32FromInt32(25) + uint32(13)
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
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var err, v3, v4, v5 int32
	var left, len1 TuLong
	var max TuInt
	var v1, v2 uint32
	var _ /* buf at bp+56 */ [1]TByte
	var _ /* stream at bp+0 */ Tz_stream
	_, _, _, _, _, _, _, _, _ = err, left, len1, max, v1, v2, v3, v4, v5
	max = libc.Uint32FromInt32(-libc.Int32FromInt32(1)) /* for detection of incomplete stream when *destLen == 0 */
	len1 = *(*TuLong)(unsafe.Pointer(sourceLen))
	if *(*TuLongf)(unsafe.Pointer(destLen)) != 0 {
		left = *(*TuLongf)(unsafe.Pointer(destLen))
		*(*TuLongf)(unsafe.Pointer(destLen)) = uint32(0)
	} else {
		left = uint32(1)
		dest = bp + 56
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = source
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = XinflateInit_(tls, bp, __ccgo_ts, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		return err
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = dest
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(0)
	for cond := true; cond; cond = err == m_Z_OK {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out == uint32(0) {
			if left > max {
				v1 = max
			} else {
				v1 = left
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
			left -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out
		}
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in == uint32(0) {
			if len1 > max {
				v2 = max
			} else {
				v2 = len1
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v2
			len1 -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in
		}
		err = Xinflate(tls, bp, m_Z_NO_FLUSH)
	}
	*(*TuLong)(unsafe.Pointer(sourceLen)) -= len1 + (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in
	if dest != bp+56 {
		*(*TuLongf)(unsafe.Pointer(destLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
	} else {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out != 0 && err == -int32(5) {
			left = uint32(1)
		}
	}
	XinflateEnd(tls, bp)
	if err == int32(m_Z_STREAM_END) {
		v3 = m_Z_OK
	} else {
		if err == int32(m_Z_NEED_DICT) {
			v4 = -int32(3)
		} else {
			if err == -int32(5) && left+(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out != 0 {
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
const m_O_LARGEFILE1 = 32768
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
	state = libc.Xmalloc(tls, uint32(156))
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
	for *(*int8)(unsafe.Pointer(mode)) != 0 {
		if int32(*(*int8)(unsafe.Pointer(mode))) >= int32('0') && int32(*(*int8)(unsafe.Pointer(mode))) <= int32('9') {
			(*Tgz_state)(unsafe.Pointer(state)).Flevel = int32(*(*int8)(unsafe.Pointer(mode))) - int32('0')
		} else {
			switch int32(*(*int8)(unsafe.Pointer(mode))) {
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
	(*Tgz_state)(unsafe.Pointer(state)).Fpath = libc.Xmalloc(tls, len1+uint32(1))
	if (*Tgz_state)(unsafe.Pointer(state)).Fpath == libc.UintptrFromInt32(0) {
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	libc.X__builtin_snprintf(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath, len1+uint32(1), __ccgo_ts+584, libc.VaList(bp+8, path))
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
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(7)+libc.Uint32FromInt32(3)*libc.Uint32FromInt64(4))
		path = v1
	}
	if v2 || v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	libc.X__builtin_snprintf(tls, path, libc.Uint32FromInt32(7)+libc.Uint32FromInt32(3)*libc.Uint32FromInt64(4), __ccgo_ts+587, libc.VaList(bp+8, fd))
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
		if libc.Bool(libc.Bool(uint32(4) == uint32(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > Xgz_intmax(tls)) || libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave) > offset {
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
	v1 = libc.Xmalloc(tls, libc.Xstrlen(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)+libc.Xstrlen(tls, msg)+uint32(3))
	(*Tgz_state)(unsafe.Pointer(state)).Fmsg = v1
	if v1 == libc.UintptrFromInt32(0) {
		(*Tgz_state)(unsafe.Pointer(state)).Ferr = -int32(4)
		return
	}
	libc.X__builtin_snprintf(tls, (*Tgz_state)(unsafe.Pointer(state)).Fmsg, libc.Xstrlen(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)+libc.Xstrlen(tls, msg)+uint32(3), __ccgo_ts+609, libc.VaList(bp+8, (*Tgz_state)(unsafe.Pointer(state)).Fpath, __ccgo_ts+616, msg))
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
const m_O_LARGEFILE2 = 0100000
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
		ret = libc.Xread(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, buf+uintptr(*(*uint32)(unsafe.Pointer(have))), get)
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
	strm = state + 100
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
		*(*TuInt)(unsafe.Pointer(strm + 4)) += *(*uint32)(unsafe.Pointer(bp))
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
	strm = state + 100
	/* allocate read buffers and inflate memory */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) {
		/* allocate buffers */
		(*Tgz_state)(unsafe.Pointer(state)).Fin = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant)
		(*Tgz_state)(unsafe.Pointer(state)).Fout = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant<<int32(1))
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
		if XinflateInit2_(tls, state+100, libc.Int32FromInt32(15)+libc.Int32FromInt32(16), __ccgo_ts, libc.Int32FromInt64(56)) != m_Z_OK { /* gunzip */
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
	libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, (*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
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
	strm = state + 100
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
	strm = state + 100
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
			if libc.Bool(libc.Bool(uint32(4) == uint32(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > Xgz_intmax(tls)) || libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave) > len1 {
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
	if len1 == uint32(0) {
		return uint32(0)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return uint32(0)
		}
	}
	/* get len bytes to buf, or less than len if at the end */
	got = uint32(0)
	for cond := true; cond; cond = len1 != 0 {
		/* set n to the maximum amount of len that fits in an unsigned int */
		*(*uint32)(unsafe.Pointer(bp)) = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
		if *(*uint32)(unsafe.Pointer(bp)) > len1 {
			*(*uint32)(unsafe.Pointer(bp)) = len1
		}
		/* first just try copying data from the output buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave < *(*uint32)(unsafe.Pointer(bp)) {
				*(*uint32)(unsafe.Pointer(bp)) = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			libc.Xmemcpy(tls, buf, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, *(*uint32)(unsafe.Pointer(bp)))
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
						return uint32(0)
					}
					continue /* no progress yet -- go back to copy above */
					/* the copy above assures that we will leave with space in the
					   output buffer, allowing at least one gzungetc() to succeed */
				} else {
					if (*Tgz_state)(unsafe.Pointer(state)).Fhow == int32(m_COPY) { /* read directly */
						if _gz_load(tls, state, buf, *(*uint32)(unsafe.Pointer(bp)), bp) == -int32(1) {
							return uint32(0)
						}
					} else { /* state->how == GZIP */
						(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_out = *(*uint32)(unsafe.Pointer(bp))
						(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_out = buf
						if _gz_decomp(tls, state) == -int32(1) {
							return uint32(0)
						}
						*(*uint32)(unsafe.Pointer(bp)) = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
						(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
					}
				}
			}
		}
		/* update progress */
		len1 -= *(*uint32)(unsafe.Pointer(bp))
		buf = buf + uintptr(*(*uint32)(unsafe.Pointer(bp)))
		got += *(*uint32)(unsafe.Pointer(bp))
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
	len1 = _gz_read(tls, state, buf, len1)
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
	var v1 uint32
	_, _, _ = len1, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return uint32(0)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return uint32(0)
	}
	/* compute bytes to read -- error on overflow */
	len1 = nitems * size
	if size != 0 && len1/size != nitems {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+734)
		return uint32(0)
	}
	/* read len or fewer bytes to buf, return the number of full items read */
	if len1 != 0 {
		v1 = _gz_read(tls, state, buf, len1) / size
	} else {
		v1 = uint32(0)
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
		v2 = state + 4
		v1 = *(*uintptr)(unsafe.Pointer(v2))
		*(*uintptr)(unsafe.Pointer(v2))++
		return libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(v1)))
	}
	/* nothing there -- try gz_read() */
	if _gz_read(tls, state, bp, uint32(1)) < uint32(1) {
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
			eol = libc.Xmemchr(tls, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, int32('\n'), n)
			if eol != libc.UintptrFromInt32(0) {
				n = libc.Uint32FromInt32(int32(eol)-int32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) + uint32(1)
			}
			/* copy through end-of-line, or remainder if not found */
			libc.Xmemcpy(tls, buf, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, n)
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
	*(*int8)(unsafe.Pointer(buf)) = 0
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
		XinflateEnd(tls, state+100)
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
	strm = state + 100
	/* allocate input buffer (double size for gzprintf) */
	(*Tgz_state)(unsafe.Pointer(state)).Fin = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant<<int32(1))
	if (*Tgz_state)(unsafe.Pointer(state)).Fin == libc.UintptrFromInt32(0) {
		Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
		return -int32(1)
	}
	/* only need output buffer and deflate state if compressing */
	if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
		/* allocate output buffer */
		(*Tgz_state)(unsafe.Pointer(state)).Fout = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant)
		if (*Tgz_state)(unsafe.Pointer(state)).Fout == libc.UintptrFromInt32(0) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			Xgz_error(tls, state, -int32(4), __ccgo_ts+595)
			return -int32(1)
		}
		/* allocate deflate memory, set up for gzip compression */
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = uintptr(m_Z_NULL)
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = uintptr(m_Z_NULL)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = uintptr(m_Z_NULL)
		ret = XdeflateInit2_(tls, strm, (*Tgz_state)(unsafe.Pointer(state)).Flevel, int32(m_Z_DEFLATED), libc.Int32FromInt32(m_MAX_WBITS)+libc.Int32FromInt32(16), int32(m_DEF_MEM_LEVEL), (*Tgz_state)(unsafe.Pointer(state)).Fstrategy, __ccgo_ts, libc.Int32FromInt64(56))
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
	strm = state + 100
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
			writ = libc.Xwrite(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, put)
			if writ < 0 {
				Xgz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
				return -int32(1)
			}
			*(*TuInt)(unsafe.Pointer(strm + 4)) -= libc.Uint32FromInt32(writ)
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
				if int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_out)-int32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext) > libc.Int32FromUint32(max) {
					v2 = max
				} else {
					v2 = libc.Uint32FromInt32(int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_out) - int32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext))
				}
				put = v2
				writ = libc.Xwrite(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, put)
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
	strm = state + 100
	/* consume whatever's left in the input buffer */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
		return -int32(1)
	}
	/* compress len zeros (len guaranteed > 0) */
	first = int32(1)
	for len1 != 0 {
		if libc.Bool(libc.Bool(uint32(4) == uint32(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fsize > Xgz_intmax(tls)) || libc.Int64FromUint32((*Tgz_state)(unsafe.Pointer(state)).Fsize) > len1 {
			v1 = libc.Uint32FromInt64(len1)
		} else {
			v1 = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		}
		n = v1
		if first != 0 {
			libc.Xmemset(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin, 0, n)
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
	if len1 == uint32(0) {
		return uint32(0)
	}
	/* allocate memory if this is the first time through */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return uint32(0)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return uint32(0)
		}
	}
	/* for small len, copy to input buffer, otherwise compress directly */
	if len1 < (*Tgz_state)(unsafe.Pointer(state)).Fsize {
		/* copy to input buffer, compress when full */
		for cond := true; cond; cond = len1 != 0 {
			if (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
			}
			have = libc.Uint32FromInt32(int32((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in+uintptr((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in)) - int32((*Tgz_state)(unsafe.Pointer(state)).Fin))
			copy1 = (*Tgz_state)(unsafe.Pointer(state)).Fsize - have
			if copy1 > len1 {
				copy1 = len1
			}
			libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr(have), buf, copy1)
			(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in += copy1
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(copy1)
			buf = buf + uintptr(copy1)
			len1 -= copy1
			if len1 != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
				return uint32(0)
			}
		}
	} else {
		/* consume whatever's left in the input buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return uint32(0)
		}
		/* directly compress user buffer to file */
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = buf
		for cond := true; cond; cond = len1 != 0 {
			n = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
			if n > len1 {
				n = len1
			}
			(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += libc.Int64FromUint32(n)
			if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
				return uint32(0)
			}
			len1 -= n
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
	return libc.Int32FromUint32(_gz_write(tls, state, buf, len1))
}

// C documentation
//
//	/* -- see zlib.h -- */
func Xgzfwrite(tls *libc.TLS, buf Tvoidpc, size Tz_size_t, nitems Tz_size_t, file TgzFile) (r Tz_size_t) {
	var len1 Tz_size_t
	var state Tgz_statep
	var v1 uint32
	_, _, _ = len1, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return uint32(0)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return uint32(0)
	}
	/* compute bytes to read -- error on overflow */
	len1 = nitems * size
	if size != 0 && len1/size != nitems {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+734)
		return uint32(0)
	}
	/* write len bytes to buf, return the number of full items written */
	if len1 != 0 {
		v1 = _gz_write(tls, state, buf, len1) / size
	} else {
		v1 = uint32(0)
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
	strm = state + 100
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
		have = libc.Uint32FromInt32(int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in+uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in)) - int32((*Tgz_state)(unsafe.Pointer(state)).Fin))
		if have < (*Tgz_state)(unsafe.Pointer(state)).Fsize {
			*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fin + uintptr(have))) = libc.Uint8FromInt32(c)
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in++
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos++
			return c & int32(0xff)
		}
	}
	/* no room in buffer or not initialized, use gz_write() */
	(*(*[1]uint8)(unsafe.Pointer(bp)))[0] = libc.Uint8FromInt32(c)
	if _gz_write(tls, state, bp, uint32(1)) != uint32(1) {
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
	if libc.Int32FromUint32(len1) < 0 || len1 != len1 {
		Xgz_error(tls, state, -int32(2), __ccgo_ts+874)
		return -int32(1)
	}
	put = _gz_write(tls, state, s, len1)
	if put < len1 {
		v1 = -int32(1)
	} else {
		v1 = libc.Int32FromUint32(len1)
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
	strm = state + 100
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
	next = (*Tgz_state)(unsafe.Pointer(state)).Fin + uintptr(int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in)-int32((*Tgz_state)(unsafe.Pointer(state)).Fin)) + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
	*(*int8)(unsafe.Pointer(next + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize-uint32(1)))) = 0
	len1 = libc.X__builtin_vsnprintf(tls, next, (*Tgz_state)(unsafe.Pointer(state)).Fsize, format, va)
	/* check that printf() results fit in buffer */
	if len1 == 0 || libc.Uint32FromInt32(len1) >= (*Tgz_state)(unsafe.Pointer(state)).Fsize || int32(*(*int8)(unsafe.Pointer(next + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize-uint32(1))))) != 0 {
		return 0
	}
	/* update buffer and position, compress first half if past that */
	*(*TuInt)(unsafe.Pointer(strm + 4)) += libc.Uint32FromInt32(len1)
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(len1)
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in >= (*Tgz_state)(unsafe.Pointer(state)).Fsize {
		left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in - (*Tgz_state)(unsafe.Pointer(state)).Fsize
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
		libc.Xmemmove(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize), left)
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
	strm = state + 100
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
			XdeflateEnd(tls, state+100)
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

var Xdeflate_copyright = [68]int8{' ', 'd', 'e', 'f', 'l', 'a', 't', 'e', ' ', '1', '.', '3', '.', '1', ' ', 'C', 'o', 'p', 'y', 'r', 'i', 'g', 'h', 't', ' ', '1', '9', '9', '5', '-', '2', '0', '2', '4', ' ', 'J', 'e', 'a', 'n', '-', 'l', 'o', 'u', 'p', ' ', 'G', 'a', 'i', 'l', 'l', 'y', ' ', 'a', 'n', 'd', ' ', 'M', 'a', 'r', 'k', ' ', 'A', 'd', 'l', 'e', 'r', ' '}

var Xinflate_copyright = [47]int8{' ', 'i', 'n', 'f', 'l', 'a', 't', 'e', ' ', '1', '.', '3', '.', '1', ' ', 'C', 'o', 'p', 'y', 'r', 'i', 'g', 'h', 't', ' ', '1', '9', '9', '5', '-', '2', '0', '2', '4', ' ', 'M', 'a', 'r', 'k', ' ', 'A', 'd', 'l', 'e', 'r', ' '}

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
