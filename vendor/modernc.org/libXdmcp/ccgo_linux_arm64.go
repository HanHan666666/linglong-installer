// Code generated for linux/arm64 by 'generator -no-main-minimize -DHAVE_LIBBSD --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -I /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libbsd/include/linux/arm64 -lbsd -o libxdmcp.go --package-name libxdmcp .libs/libXdmcp.a', DO NOT EDIT.

//go:build linux && arm64

package libxdmcp

import (
	"reflect"
	"unsafe"

	"modernc.org/libbsd"
	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

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
const m_Above = 0
const m_AllTemporary = 0
const m_AllocAll = 1
const m_AllocNone = 0
const m_AllowExposures = 1
const m_AlreadyGrabbed = 1
const m_Always = 2
const m_AnyButton = 0
const m_AnyKey = 0
const m_AnyPropertyType = 0
const m_ArcChord = 0
const m_ArcPieSlice = 1
const m_AsyncBoth = 6
const m_AsyncKeyboard = 3
const m_AsyncPointer = 0
const m_AutoRepeatModeDefault = 2
const m_AutoRepeatModeOff = 0
const m_AutoRepeatModeOn = 1
const m_BIG_ENDIAN = "__BIG_ENDIAN"
const m_BYTE_ORDER = "__BYTE_ORDER"
const m_BadAccess = 10
const m_BadAlloc = 11
const m_BadAtom = 5
const m_BadColor = 12
const m_BadCursor = 6
const m_BadDrawable = 9
const m_BadFont = 7
const m_BadGC = 13
const m_BadIDChoice = 14
const m_BadImplementation = 17
const m_BadLength = 16
const m_BadMatch = 8
const m_BadName = 15
const m_BadPixmap = 4
const m_BadRequest = 1
const m_BadValue = 2
const m_BadWindow = 3
const m_Below = 1
const m_BottomIf = 3
const m_Button1 = 1
const m_Button2 = 2
const m_Button3 = 3
const m_Button4 = 4
const m_Button5 = 5
const m_ButtonPress = 4
const m_ButtonRelease = 5
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
const m_CapButt = 1
const m_CapNotLast = 0
const m_CapProjecting = 3
const m_CapRound = 2
const m_CenterGravity = 5
const m_CirculateNotify = 26
const m_CirculateRequest = 27
const m_ClientMessage = 33
const m_ClipByChildren = 0
const m_ColormapInstalled = 1
const m_ColormapNotify = 32
const m_ColormapUninstalled = 0
const m_Complex = 0
const m_ConfigureNotify = 22
const m_ConfigureRequest = 23
const m_ControlMapIndex = 2
const m_Convex = 2
const m_CoordModeOrigin = 0
const m_CoordModePrevious = 1
const m_CopyFromParent = 0
const m_CreateNotify = 16
const m_CurrentTime = 0
const m_CursorShape = 0
const m_DN_ACCESS = 0x00000001
const m_DN_ATTRIB = 0x00000020
const m_DN_CREATE = 0x00000004
const m_DN_DELETE = 0x00000008
const m_DN_MODIFY = 0x00000002
const m_DN_MULTISHOT = 0x80000000
const m_DN_RENAME = 0x00000010
const m_DefaultBlanking = 2
const m_DefaultExposures = 2
const m_DestroyAll = 0
const m_DestroyNotify = 17
const m_DirectColor = 5
const m_DisableAccess = 0
const m_DisableScreenInterval = 0
const m_DisableScreenSaver = 0
const m_DontAllowExposures = 0
const m_DontPreferBlanking = 0
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_EastGravity = 6
const m_EnableAccess = 1
const m_EnterNotify = 7
const m_EvenOddRule = 0
const m_Expose = 12
const m_FALLOC_FL_KEEP_SIZE = 1
const m_FALLOC_FL_PUNCH_HOLE = 2
const m_FALSE = 0
const m_FAPPEND = "O_APPEND"
const m_FASYNC = "O_ASYNC"
const m_FD_CLOEXEC = 1
const m_FD_SETSIZE = 1024
const m_FFSYNC = "O_SYNC"
const m_FNDELAY = "O_NDELAY"
const m_FNONBLOCK = "O_NONBLOCK"
const m_FUNCPROTO = 15
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
const m_F_LOCK = 1
const m_F_NOTIFY = 1026
const m_F_OFD_GETLK = 36
const m_F_OFD_SETLK = 37
const m_F_OFD_SETLKW = 38
const m_F_OK = 0
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
const m_F_TEST = 3
const m_F_TLOCK = 2
const m_F_ULOCK = 0
const m_F_UNLCK = 2
const m_F_WRLCK = 1
const m_FamilyChaos = 2
const m_FamilyDECnet = 1
const m_FamilyInternet = 0
const m_FamilyInternet6 = 6
const m_FamilyServerInterpreted = 5
const m_FillOpaqueStippled = 3
const m_FillSolid = 0
const m_FillStippled = 2
const m_FillTiled = 1
const m_FirstExtensionError = 128
const m_FocusIn = 9
const m_FocusOut = 10
const m_FontChange = 255
const m_FontLeftToRight = 0
const m_FontRightToLeft = 1
const m_ForgetGravity = 0
const m_GCLastBit = 22
const m_GXand = 0x1
const m_GXandInverted = 0x4
const m_GXandReverse = 0x2
const m_GXclear = 0x0
const m_GXcopy = 0x3
const m_GXcopyInverted = 0xc
const m_GXequiv = 0x9
const m_GXinvert = 0xa
const m_GXnand = 0xe
const m_GXnoop = 0x5
const m_GXnor = 0x8
const m_GXor = 0x7
const m_GXorInverted = 0xd
const m_GXorReverse = 0xb
const m_GXset = 0xf
const m_GXxor = 0x6
const m_GenericEvent = 35
const m_GrabFrozen = 4
const m_GrabInvalidTime = 2
const m_GrabModeAsync = 1
const m_GrabModeSync = 0
const m_GrabNotViewable = 3
const m_GrabSuccess = 0
const m_GraphicsExpose = 13
const m_GravityNotify = 24
const m_GrayScale = 1
const m_HASXDMAUTH = 1
const m_HAVE_ARC4RANDOM_BUF = 1
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_LIBBSD = 1
const m_HAVE_LRAND48 = 1
const m_HAVE_MEMORY_H = 1
const m_HAVE_SRAND48 = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_HostDelete = 1
const m_HostInsert = 0
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
const m_ITIMER_PROF = 2
const m_ITIMER_REAL = 0
const m_ITIMER_VIRTUAL = 1
const m_IncludeInferiors = 1
const m_InputFocus = 1
const m_InputOnly = 2
const m_InputOutput = 1
const m_IsUnmapped = 0
const m_IsUnviewable = 1
const m_IsViewable = 2
const m_JoinBevel = 2
const m_JoinMiter = 0
const m_JoinRound = 1
const m_KeyPress = 2
const m_KeyRelease = 3
const m_KeymapNotify = 11
const m_LASTEvent = 36
const m_LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const m_LSBFirst = 0
const m_LT_OBJDIR = ".libs/"
const m_L_INCR = 1
const m_L_SET = 0
const m_L_XTND = 2
const m_LastExtensionError = 255
const m_LeaveNotify = 8
const m_LedModeOff = 0
const m_LedModeOn = 1
const m_LineDoubleDash = 2
const m_LineOnOffDash = 1
const m_LineSolid = 0
const m_LockMapIndex = 1
const m_LowerHighest = 1
const m_MAX_HANDLE_SZ = 128
const m_MSBFirst = 1
const m_MapNotify = 19
const m_MapRequest = 20
const m_MappingBusy = 1
const m_MappingFailed = 2
const m_MappingKeyboard = 1
const m_MappingModifier = 0
const m_MappingNotify = 34
const m_MappingPointer = 2
const m_MappingSuccess = 0
const m_Mod1MapIndex = 3
const m_Mod2MapIndex = 4
const m_Mod3MapIndex = 5
const m_Mod4MapIndex = 6
const m_Mod5MapIndex = 7
const m_MotionNotify = 6
const m_NDEBUG = 1
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 0
const m_NoEventMask = 0
const m_NoExpose = 14
const m_NoSymbol = 0
const m_Nonconvex = 1
const m_None = 0
const m_NorthEastGravity = 3
const m_NorthGravity = 2
const m_NorthWestGravity = 1
const m_NotUseful = 0
const m_NotifyAncestor = 0
const m_NotifyDetailNone = 7
const m_NotifyGrab = 1
const m_NotifyHint = 1
const m_NotifyInferior = 2
const m_NotifyNonlinear = 3
const m_NotifyNonlinearVirtual = 4
const m_NotifyNormal = 0
const m_NotifyPointer = 5
const m_NotifyPointerRoot = 6
const m_NotifyUngrab = 2
const m_NotifyVirtual = 1
const m_NotifyWhileGrabbed = 3
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
const m_Opposite = 4
const m_PACKAGE = "libXdmcp"
const m_PACKAGE_BUGREPORT = "https://bugs.freedesktop.org/enter_bug.cgi?product=xorg"
const m_PACKAGE_NAME = "libXdmcp"
const m_PACKAGE_STRING = "libXdmcp 1.1.2"
const m_PACKAGE_TARNAME = "libXdmcp"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "1.1.2"
const m_PACKAGE_VERSION_MAJOR = 1
const m_PACKAGE_VERSION_MINOR = 1
const m_PACKAGE_VERSION_PATCHLEVEL = 2
const m_PDP_ENDIAN = "__PDP_ENDIAN"
const m_POSIX_CLOSE_RESTART = 0
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
const m_PTRDIFF_MAX = "INT64_MAX"
const m_PTRDIFF_MIN = "INT64_MIN"
const m_ParentRelative = 1
const m_PlaceOnBottom = 1
const m_PlaceOnTop = 0
const m_PointerRoot = 1
const m_PointerWindow = 0
const m_PreferBlanking = 1
const m_PropModeAppend = 2
const m_PropModePrepend = 1
const m_PropModeReplace = 0
const m_PropertyDelete = 1
const m_PropertyNewValue = 0
const m_PropertyNotify = 28
const m_PseudoColor = 3
const m_RAND_MAX = 0x7fffffff
const m_RWF_WRITE_LIFE_NOT_SET = 0
const m_RWH_WRITE_LIFE_EXTREME = 5
const m_RWH_WRITE_LIFE_LONG = 4
const m_RWH_WRITE_LIFE_MEDIUM = 3
const m_RWH_WRITE_LIFE_NONE = 1
const m_RWH_WRITE_LIFE_SHORT = 2
const m_R_OK = 4
const m_RaiseLowest = 0
const m_ReparentNotify = 21
const m_ReplayKeyboard = 5
const m_ReplayPointer = 2
const m_ResizeRequest = 25
const m_RetainPermanent = 1
const m_RetainTemporary = 2
const m_RevertToParent = 2
const m_SEEK_DATA = 3
const m_SEEK_HOLE = 4
const m_SIG_ATOMIC_MAX = "INT32_MAX"
const m_SIG_ATOMIC_MIN = "INT32_MIN"
const m_SIZE_MAX = "UINT64_MAX"
const m_SPLICE_F_GIFT = 8
const m_SPLICE_F_MORE = 4
const m_SPLICE_F_MOVE = 1
const m_SPLICE_F_NONBLOCK = 2
const m_STDC_HEADERS = 1
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
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
const m_ScreenSaverActive = 1
const m_ScreenSaverReset = 0
const m_SelectionClear = 29
const m_SelectionNotify = 31
const m_SelectionRequest = 30
const m_SetModeDelete = 1
const m_SetModeInsert = 0
const m_ShiftMapIndex = 0
const m_SouthEastGravity = 9
const m_SouthGravity = 8
const m_SouthWestGravity = 7
const m_StaticColor = 2
const m_StaticGravity = 10
const m_StaticGray = 0
const m_StippleShape = 2
const m_Success = 0
const m_SyncBoth = 7
const m_SyncKeyboard = 4
const m_SyncPointer = 1
const m_TIMER_ABSTIME = 1
const m_TIME_UTC = 1
const m_TRUE = 1
const m_TileShape = 1
const m_TopIf = 2
const m_TrueColor = 4
const m_UINT16_MAX = 65535
const m_UINT32_MAX = "0xffffffffu"
const m_UINT64_MAX = "0xffffffffffffffffu"
const m_UINT8_MAX = 255
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
const m_UnmapGravity = 0
const m_UnmapNotify = 18
const m_Unsorted = 0
const m_VERSION = "1.1.2"
const m_VisibilityFullyObscured = 2
const m_VisibilityNotify = 15
const m_VisibilityPartiallyObscured = 1
const m_VisibilityUnobscured = 0
const m_WINT_MAX = "UINT32_MAX"
const m_WINT_MIN = 0
const m_WNOHANG = 1
const m_WUNTRACED = 2
const m_W_OK = 2
const m_WestGravity = 4
const m_WhenMapped = 1
const m_WindingRule = 1
const m_XDM_DEFAULT_MCAST_ADDR6 = "ff02:0:0:0:0:0:0:12b"
const m_XDM_KA_RTX_LIMIT = 4
const m_XDM_MAX_MSGLEN = 8192
const m_XDM_MAX_RTX = 32
const m_XDM_MIN_RTX = 2
const m_XDM_PROTOCOL_VERSION = 1
const m_XDM_RTX_LIMIT = 7
const m_XDM_UDP_PORT = 177
const m_XMD_H = 1
const m_XYBitmap = 0
const m_XYPixmap = 1
const m_X_BIG_ENDIAN = "BIG_ENDIAN"
const m_X_BYTE_ORDER = "BYTE_ORDER"
const m_X_LITTLE_ENDIAN = "LITTLE_ENDIAN"
const m_X_OK = 1
const m_X_PROTOCOL = 11
const m_X_PROTOCOL_REVISION = 0
const m_YSorted = 1
const m_YXBanded = 3
const m_YXSorted = 2
const m_ZPixmap = 2
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
const m__GNU_SOURCE = 1
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
const m__POSIX2_C_BIND = "_POSIX_VERSION"
const m__POSIX2_VERSION = "_POSIX_VERSION"
const m__POSIX_ADVISORY_INFO = "_POSIX_VERSION"
const m__POSIX_ASYNCHRONOUS_IO = "_POSIX_VERSION"
const m__POSIX_BARRIERS = "_POSIX_VERSION"
const m__POSIX_CHOWN_RESTRICTED = 1
const m__POSIX_CLOCK_SELECTION = "_POSIX_VERSION"
const m__POSIX_CPUTIME = "_POSIX_VERSION"
const m__POSIX_FSYNC = "_POSIX_VERSION"
const m__POSIX_IPV6 = "_POSIX_VERSION"
const m__POSIX_JOB_CONTROL = 1
const m__POSIX_MAPPED_FILES = "_POSIX_VERSION"
const m__POSIX_MEMLOCK = "_POSIX_VERSION"
const m__POSIX_MEMLOCK_RANGE = "_POSIX_VERSION"
const m__POSIX_MEMORY_PROTECTION = "_POSIX_VERSION"
const m__POSIX_MESSAGE_PASSING = "_POSIX_VERSION"
const m__POSIX_MONOTONIC_CLOCK = "_POSIX_VERSION"
const m__POSIX_NO_TRUNC = 1
const m__POSIX_PTHREAD_SEMANTICS = 1
const m__POSIX_RAW_SOCKETS = "_POSIX_VERSION"
const m__POSIX_READER_WRITER_LOCKS = "_POSIX_VERSION"
const m__POSIX_REALTIME_SIGNALS = "_POSIX_VERSION"
const m__POSIX_REGEXP = 1
const m__POSIX_SAVED_IDS = 1
const m__POSIX_SEMAPHORES = "_POSIX_VERSION"
const m__POSIX_SHARED_MEMORY_OBJECTS = "_POSIX_VERSION"
const m__POSIX_SHELL = 1
const m__POSIX_SPAWN = "_POSIX_VERSION"
const m__POSIX_SPIN_LOCKS = "_POSIX_VERSION"
const m__POSIX_THREADS = "_POSIX_VERSION"
const m__POSIX_THREAD_ATTR_STACKADDR = "_POSIX_VERSION"
const m__POSIX_THREAD_ATTR_STACKSIZE = "_POSIX_VERSION"
const m__POSIX_THREAD_CPUTIME = "_POSIX_VERSION"
const m__POSIX_THREAD_PRIORITY_SCHEDULING = "_POSIX_VERSION"
const m__POSIX_THREAD_PROCESS_SHARED = "_POSIX_VERSION"
const m__POSIX_THREAD_SAFE_FUNCTIONS = "_POSIX_VERSION"
const m__POSIX_TIMEOUTS = "_POSIX_VERSION"
const m__POSIX_TIMERS = "_POSIX_VERSION"
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
const m__TANDEM_SOURCE = 1
const m__XOPEN_ENH_I18N = 1
const m__XOPEN_UNIX = 1
const m__XOPEN_VERSION = 700
const m__X_INLINE = "inline"
const m__X_RESTRICT_KYWD = "restrict"
const m__Xconst = "const"
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
const m___EXTENSIONS__ = 1
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
const m___tm_gmtoff = "tm_gmtoff"
const m___tm_zone = "tm_zone"
const m___unix = 1
const m___unix__ = 1
const m_alloca = "__builtin_alloca"
const m_linux = 1
const m_loff_t = "off_t"
const m_unix = 1

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint64

type t__predefined_wchar_t = uint32

type t__predefined_ptrdiff_t = int64

type Tblksize_t = int32

type Tnlink_t = uint32

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

type Tlocale_t = uintptr

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

type Tintptr_t = int64

type Titimerval = struct {
	Fit_interval Ttimeval
	Fit_value    Ttimeval
}

type Ttimezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
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

type TXID = uint64

type TMask = uint64

type TAtom = uint64

type TVisualID = uint64

type TTime = uint64

type TWindow = uint64

type TDrawable = uint64

type TFont = uint64

type TPixmap = uint64

type TCursor = uint64

type TColormap = uint64

type TGContext = uint64

type TKeySym = uint64

type TKeyCode = uint8

type TINT64 = int64

type TINT32 = int32

type TINT16 = int16

type TINT8 = int8

type TCARD64 = uint64

type TCARD32 = uint32

type TCARD16 = uint16

type TCARD8 = uint8

type TBITS32 = uint32

type TBITS16 = uint16

type TBYTE = uint8

type TBOOL = uint8

type TxdmOpCode = int32

const _BROADCAST_QUERY = 1
const _QUERY = 2
const _INDIRECT_QUERY = 3
const _FORWARD_QUERY = 4
const _WILLING = 5
const _UNWILLING = 6
const _REQUEST = 7
const _ACCEPT = 8
const _DECLINE = 9
const _MANAGE = 10
const _REFUSE = 11
const _FAILED = 12
const _KEEPALIVE = 13
const _ALIVE = 14

type Txdmcp_states = int32

const _XDM_QUERY = 0
const _XDM_BROADCAST = 1
const _XDM_INDIRECT = 2
const _XDM_COLLECT_QUERY = 3
const _XDM_COLLECT_BROADCAST_QUERY = 4
const _XDM_COLLECT_INDIRECT_QUERY = 5
const _XDM_START_CONNECTION = 6
const _XDM_AWAIT_REQUEST_RESPONSE = 7
const _XDM_AWAIT_MANAGE_RESPONSE = 8
const _XDM_MANAGE = 9
const _XDM_RUN_SESSION = 10
const _XDM_OFF = 11
const _XDM_AWAIT_USER_INPUT = 12
const _XDM_KEEPALIVE = 13
const _XDM_AWAIT_ALIVE_RESPONSE = 14
const _XDM_KEEP_ME_LAST = 15

type TCARD8Ptr = uintptr

type TCARD16Ptr = uintptr

type TCARD32Ptr = uintptr

type TARRAY8 = struct {
	Flength TCARD16
	Fdata   TCARD8Ptr
}

type T_ARRAY8 = TARRAY8

type TARRAY8Ptr = uintptr

type TARRAY16 = struct {
	Flength TCARD8
	Fdata   TCARD16Ptr
}

type T_ARRAY16 = TARRAY16

type TARRAY16Ptr = uintptr

type TARRAY32 = struct {
	Flength TCARD8
	Fdata   TCARD32Ptr
}

type T_ARRAY32 = TARRAY32

type TARRAY32Ptr = uintptr

type TARRAYofARRAY8 = struct {
	Flength TCARD8
	Fdata   TARRAY8Ptr
}

type T_ARRAYofARRAY8 = TARRAYofARRAY8

type TARRAYofARRAY8Ptr = uintptr

type TXdmcpHeader = struct {
	Fversion TCARD16
	Fopcode  TCARD16
	Flength  TCARD16
}

type T_XdmcpHeader = TXdmcpHeader

type TXdmcpHeaderPtr = uintptr

type TXdmcpBuffer = struct {
	Fdata    uintptr
	Fsize    int32
	Fpointer int32
	Fcount   int32
}

type T_XdmcpBuffer = TXdmcpBuffer

type TXdmcpBufferPtr = uintptr

type TXdmAuthKeyRec = struct {
	Fdata [8]TBYTE
}

type T_XdmAuthKey = TXdmAuthKeyRec

type TXdmAuthKeyPtr = uintptr

type TXdmcpNetaddr = uintptr

type Tuintptr_t = uint64

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

type Twchar_t = uint32

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

// C documentation
//
//	/*
//	 * This variant of malloc does not return NULL if zero size is passed into.
//	 */
func _xmalloc(tls *libc.TLS, size Tsize_t) (r uintptr) {
	var v1 uint64
	_ = v1
	if size != 0 {
		v1 = size
	} else {
		v1 = uint64(1)
	}
	return libc.Xmalloc(tls, v1)
}

// C documentation
//
//	/*
//	 * This variant of calloc does not return NULL if zero count is passed into.
//	 */
func _xcalloc(tls *libc.TLS, n Tsize_t, size Tsize_t) (r uintptr) {
	var v1 uint64
	_ = v1
	if n != 0 {
		v1 = n
	} else {
		v1 = uint64(1)
	}
	return libc.Xcalloc(tls, v1, size)
}

// C documentation
//
//	/*
//	 * This variant of realloc does not return NULL if zero size is passed into
//	 */
func _xrealloc(tls *libc.TLS, ptr uintptr, size Tsize_t) (r uintptr) {
	var v1 uint64
	_ = v1
	if size != 0 {
		v1 = size
	} else {
		v1 = uint64(1)
	}
	return libc.Xrealloc(tls, ptr, v1)
}

func XXdmcpAllocARRAY8(tls *libc.TLS, array TARRAY8Ptr, length int32) (r int32) {
	/* length defined in ARRAY8 struct is a CARD16 (not CARD8 like the rest) */
	if length > int32(m_UINT16_MAX) || length < 0 {
		(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
	} else {
		(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = _xmalloc(tls, libc.Uint64FromInt32(length)*uint64(1))
	}
	if (*T_ARRAY8)(unsafe.Pointer(array)).Fdata == libc.UintptrFromInt32(0) {
		(*T_ARRAY8)(unsafe.Pointer(array)).Flength = uint16(0)
		return m_FALSE
	}
	(*T_ARRAY8)(unsafe.Pointer(array)).Flength = libc.Uint16FromInt32(length)
	return int32(m_TRUE)
}

func XXdmcpAllocARRAY16(tls *libc.TLS, array TARRAY16Ptr, length int32) (r int32) {
	/* length defined in ARRAY16 struct is a CARD8 */
	if length > int32(m_UINT8_MAX) || length < 0 {
		(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
	} else {
		(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = _xmalloc(tls, libc.Uint64FromInt32(length)*uint64(2))
	}
	if (*T_ARRAY16)(unsafe.Pointer(array)).Fdata == libc.UintptrFromInt32(0) {
		(*T_ARRAY16)(unsafe.Pointer(array)).Flength = uint8(0)
		return m_FALSE
	}
	(*T_ARRAY16)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	return int32(m_TRUE)
}

func XXdmcpAllocARRAY32(tls *libc.TLS, array TARRAY32Ptr, length int32) (r int32) {
	/* length defined in ARRAY32 struct is a CARD8 */
	if length > int32(m_UINT8_MAX) || length < 0 {
		(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
	} else {
		(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = _xmalloc(tls, libc.Uint64FromInt32(length)*uint64(4))
	}
	if (*T_ARRAY32)(unsafe.Pointer(array)).Fdata == libc.UintptrFromInt32(0) {
		(*T_ARRAY32)(unsafe.Pointer(array)).Flength = uint8(0)
		return m_FALSE
	}
	(*T_ARRAY32)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	return int32(m_TRUE)
}

func XXdmcpAllocARRAYofARRAY8(tls *libc.TLS, array TARRAYofARRAY8Ptr, length int32) (r int32) {
	/* length defined in ARRAYofARRAY8 struct is a CARD8 */
	if length > int32(m_UINT8_MAX) || length < 0 {
		(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
	} else {
		/*
		 * Use calloc to ensure the pointers are cleared out so we
		 * don't try to free garbage if XdmcpDisposeARRAYofARRAY8()
		 * is called before the caller sets them to valid pointers.
		 */
		(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = _xcalloc(tls, libc.Uint64FromInt32(length), uint64(16))
	}
	if (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata == libc.UintptrFromInt32(0) {
		(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength = uint8(0)
		return m_FALSE
	}
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	return int32(m_TRUE)
}

func XXdmcpARRAY8Equal(tls *libc.TLS, array1 TARRAY8Ptr, array2 TARRAY8Ptr) (r int32) {
	if libc.Int32FromUint16((*T_ARRAY8)(unsafe.Pointer(array1)).Flength) != libc.Int32FromUint16((*T_ARRAY8)(unsafe.Pointer(array2)).Flength) {
		return m_FALSE
	}
	if libc.Xmemcmp(tls, (*T_ARRAY8)(unsafe.Pointer(array1)).Fdata, (*T_ARRAY8)(unsafe.Pointer(array2)).Fdata, uint64((*T_ARRAY8)(unsafe.Pointer(array1)).Flength)) != 0 {
		return m_FALSE
	}
	return int32(m_TRUE)
}

func XXdmcpCopyARRAY8(tls *libc.TLS, src TARRAY8Ptr, dst TARRAY8Ptr) (r int32) {
	if !(XXdmcpAllocARRAY8(tls, dst, libc.Int32FromUint16((*T_ARRAY8)(unsafe.Pointer(src)).Flength)) != 0) {
		return m_FALSE
	}
	libc.Xmemmove(tls, (*T_ARRAY8)(unsafe.Pointer(dst)).Fdata, (*T_ARRAY8)(unsafe.Pointer(src)).Fdata, uint64((*T_ARRAY8)(unsafe.Pointer(src)).Flength)*uint64(1))
	return int32(m_TRUE)
}

func XXdmcpReallocARRAY8(tls *libc.TLS, array TARRAY8Ptr, length int32) (r int32) {
	var newData TCARD8Ptr
	_ = newData
	/* length defined in ARRAY8 struct is a CARD16 (not CARD8 like the rest) */
	if length > int32(m_UINT16_MAX) || length < 0 {
		return m_FALSE
	}
	newData = _xrealloc(tls, (*T_ARRAY8)(unsafe.Pointer(array)).Fdata, libc.Uint64FromInt32(length)*uint64(1))
	if !(newData != 0) {
		return m_FALSE
	}
	(*T_ARRAY8)(unsafe.Pointer(array)).Flength = libc.Uint16FromInt32(length)
	(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = newData
	return int32(m_TRUE)
}

func XXdmcpReallocARRAYofARRAY8(tls *libc.TLS, array TARRAYofARRAY8Ptr, length int32) (r int32) {
	var newData TARRAY8Ptr
	_ = newData
	/* length defined in ARRAYofARRAY8 struct is a CARD8 */
	if length > int32(m_UINT8_MAX) || length < 0 {
		return m_FALSE
	}
	newData = _xrealloc(tls, (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata, libc.Uint64FromInt32(length)*uint64(16))
	if !(newData != 0) {
		return m_FALSE
	}
	if length > libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength) {
		libc.Xmemset(tls, newData+uintptr((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)*16, 0, libc.Uint64FromInt32(length-libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength))*uint64(16))
	}
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = newData
	return int32(m_TRUE)
}

func XXdmcpReallocARRAY16(tls *libc.TLS, array TARRAY16Ptr, length int32) (r int32) {
	var newData TCARD16Ptr
	_ = newData
	/* length defined in ARRAY16 struct is a CARD8 */
	if length > int32(m_UINT8_MAX) || length < 0 {
		return m_FALSE
	}
	newData = _xrealloc(tls, (*T_ARRAY16)(unsafe.Pointer(array)).Fdata, libc.Uint64FromInt32(length)*uint64(2))
	if !(newData != 0) {
		return m_FALSE
	}
	(*T_ARRAY16)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = newData
	return int32(m_TRUE)
}

func XXdmcpReallocARRAY32(tls *libc.TLS, array TARRAY32Ptr, length int32) (r int32) {
	var newData TCARD32Ptr
	_ = newData
	/* length defined in ARRAY32 struct is a CARD8 */
	if length > int32(m_UINT8_MAX) || length < 0 {
		return m_FALSE
	}
	newData = _xrealloc(tls, (*T_ARRAY32)(unsafe.Pointer(array)).Fdata, libc.Uint64FromInt32(length)*uint64(4))
	if !(newData != 0) {
		return m_FALSE
	}
	(*T_ARRAY32)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = newData
	return int32(m_TRUE)
}

func XXdmcpDisposeARRAY8(tls *libc.TLS, array TARRAY8Ptr) {
	libc.Xfree(tls, (*T_ARRAY8)(unsafe.Pointer(array)).Fdata)
	(*T_ARRAY8)(unsafe.Pointer(array)).Flength = uint16(0)
	(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
}

func XXdmcpDisposeARRAY16(tls *libc.TLS, array TARRAY16Ptr) {
	libc.Xfree(tls, (*T_ARRAY16)(unsafe.Pointer(array)).Fdata)
	(*T_ARRAY16)(unsafe.Pointer(array)).Flength = uint8(0)
	(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
}

func XXdmcpDisposeARRAY32(tls *libc.TLS, array TARRAY32Ptr) {
	libc.Xfree(tls, (*T_ARRAY32)(unsafe.Pointer(array)).Fdata)
	(*T_ARRAY32)(unsafe.Pointer(array)).Flength = uint8(0)
	(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
}

func XXdmcpDisposeARRAYofARRAY8(tls *libc.TLS, array TARRAYofARRAY8Ptr) {
	var i int32
	_ = i
	if (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata != libc.UintptrFromInt32(0) {
		i = 0
		for {
			if !(i < libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)) {
				break
			}
			XXdmcpDisposeARRAY8(tls, (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata+uintptr(i)*16)
			goto _1
		_1:
			;
			i++
		}
		libc.Xfree(tls, (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata)
	}
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength = uint8(0)
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
}

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
	Fsa_data   [14]uint8
}

type Tsockaddr_storage = struct {
	Fss_family    Tsa_family_t
	F__ss_padding [118]uint8
	F__ss_align   uint64
}

func XXdmcpFill(tls *libc.TLS, fd int32, buffer TXdmcpBufferPtr, from TXdmcpNetaddr, fromlen uintptr) (r int32) {
	var newBuf uintptr
	_ = newBuf
	if (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize < int32(m_XDM_MAX_MSGLEN) {
		newBuf = libc.Xmalloc(tls, uint64(m_XDM_MAX_MSGLEN))
		if newBuf != 0 {
			libc.Xfree(tls, (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata)
			(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata = newBuf
			(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize = int32(m_XDM_MAX_MSGLEN)
		}
	}
	(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer = 0
	(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fcount = int32(libc.Xrecvfrom(tls, fd, (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata, libc.Uint64FromInt32((*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize), 0, from, fromlen))
	if (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fcount < int32(6) {
		(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fcount = 0
		return m_FALSE
	}
	return int32(m_TRUE)
}

func XXdmcpFlush(tls *libc.TLS, fd int32, buffer TXdmcpBufferPtr, to TXdmcpNetaddr, tolen int32) (r int32) {
	var result int32
	_ = result
	result = int32(libc.Xsendto(tls, fd, (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata, libc.Uint64FromInt32((*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer), 0, to, libc.Uint32FromInt32(tolen)))
	if result != (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer {
		return m_FALSE
	}
	return int32(m_TRUE)
}

const m_Time_t = "time_t"
const m_random = "lrand48"
const m_srandom = "srand48"

func XXdmcpGenerateKey(tls *libc.TLS, key TXdmAuthKeyPtr) {
	libbsd.Xarc4random_buf(tls, key, libc.Uint64FromInt32(8))
}

func XXdmcpCompareKeys(tls *libc.TLS, a TXdmAuthKeyPtr, b TXdmAuthKeyPtr) (r int32) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < int32(8)) {
			break
		}
		if libc.Int32FromUint8(*(*TBYTE)(unsafe.Pointer(a + uintptr(i)))) != libc.Int32FromUint8(*(*TBYTE)(unsafe.Pointer(b + uintptr(i)))) {
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpIncrementKey(tls *libc.TLS, key TXdmAuthKeyPtr) {
	var i, v3 int32
	var v1 TBYTE
	var v2 uintptr
	_, _, _, _ = i, v1, v2, v3
	i = int32(7)
	for {
		v2 = key + uintptr(i)
		*(*TBYTE)(unsafe.Pointer(v2))++
		v1 = *(*TBYTE)(unsafe.Pointer(v2))
		if !(libc.Int32FromUint8(v1) == 0) {
			break
		}
		i--
		v3 = i
		if v3 < 0 {
			break
		}
	}
}

func XXdmcpDecrementKey(tls *libc.TLS, key TXdmAuthKeyPtr) {
	var i, v3 int32
	var v1 TBYTE
	var v2 uintptr
	_, _, _, _ = i, v1, v2, v3
	i = int32(7)
	for {
		v2 = key + uintptr(i)
		v1 = *(*TBYTE)(unsafe.Pointer(v2))
		*(*TBYTE)(unsafe.Pointer(v2))--
		if !(libc.Int32FromUint8(v1) == 0) {
			break
		}
		i--
		v3 = i
		if v3 < 0 {
			break
		}
	}
}

func XXdmcpReadHeader(tls *libc.TLS, buffer TXdmcpBufferPtr, header TXdmcpHeaderPtr) (r int32) {
	if XXdmcpReadCARD16(tls, buffer, header) != 0 && XXdmcpReadCARD16(tls, buffer, header+2) != 0 && XXdmcpReadCARD16(tls, buffer, header+4) != 0 {
		return int32(m_TRUE)
	}
	return m_FALSE
}

func XXdmcpReadRemaining(tls *libc.TLS, buffer TXdmcpBufferPtr) (r int32) {
	return (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fcount - (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer
}

func XXdmcpReadARRAY8(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAY8Ptr) (r int32) {
	var i int32
	_ = i
	/*
	 * When returning FALSE, guarantee that array->data = 0.
	 * This allows the user to safely call XdmcpDisposeARRAY8(array)
	 * regardless of the return value below.
	 * Note that XdmcpDisposeARRAY*(array) will call free(array->data),
	 * so we must guarantee that array->data is NULL or a malloced pointer.
	 */
	if !(XXdmcpReadCARD16(tls, buffer, array) != 0) {
		(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return m_FALSE
	}
	if !((*T_ARRAY8)(unsafe.Pointer(array)).Flength != 0) {
		(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return int32(m_TRUE)
	}
	(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = libc.Xmalloc(tls, uint64((*T_ARRAY8)(unsafe.Pointer(array)).Flength)*uint64(1))
	if !((*T_ARRAY8)(unsafe.Pointer(array)).Fdata != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint16((*T_ARRAY8)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpReadCARD8(tls, buffer, (*T_ARRAY8)(unsafe.Pointer(array)).Fdata+uintptr(i)) != 0) {
			libc.Xfree(tls, (*T_ARRAY8)(unsafe.Pointer(array)).Fdata)
			(*T_ARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
			(*T_ARRAY8)(unsafe.Pointer(array)).Flength = uint16(0)
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpReadARRAY16(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAY16Ptr) (r int32) {
	var i int32
	_ = i
	/*
	 * When returning FALSE, guarantee that array->data = 0.
	 * This allows the user to safely call XdmcpDisposeARRAY16(array)
	 * regardless of the return value below.
	 * Note that XdmcpDisposeARRAY*(array) will call free(array->data),
	 * so we must guarantee that array->data is NULL or a malloced pointer.
	 */
	if !(XXdmcpReadCARD8(tls, buffer, array) != 0) {
		(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return m_FALSE
	}
	if !((*T_ARRAY16)(unsafe.Pointer(array)).Flength != 0) {
		(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return int32(m_TRUE)
	}
	(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = libc.Xmalloc(tls, uint64((*T_ARRAY16)(unsafe.Pointer(array)).Flength)*uint64(2))
	if !((*T_ARRAY16)(unsafe.Pointer(array)).Fdata != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint8((*T_ARRAY16)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpReadCARD16(tls, buffer, (*T_ARRAY16)(unsafe.Pointer(array)).Fdata+uintptr(i)*2) != 0) {
			libc.Xfree(tls, (*T_ARRAY16)(unsafe.Pointer(array)).Fdata)
			(*T_ARRAY16)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
			(*T_ARRAY16)(unsafe.Pointer(array)).Flength = uint8(0)
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpReadARRAY32(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAY32Ptr) (r int32) {
	var i int32
	_ = i
	/*
	 * When returning FALSE, guarantee that array->data = 0.
	 * This allows the user to safely call XdmcpDisposeARRAY32(array)
	 * regardless of the return value below.
	 * Note that XdmcpDisposeARRAY*(array) will call free(array->data),
	 * so we must guarantee that array->data is NULL or a malloced pointer.
	 */
	if !(XXdmcpReadCARD8(tls, buffer, array) != 0) {
		(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return m_FALSE
	}
	if !((*T_ARRAY32)(unsafe.Pointer(array)).Flength != 0) {
		(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return int32(m_TRUE)
	}
	(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = libc.Xmalloc(tls, uint64((*T_ARRAY32)(unsafe.Pointer(array)).Flength)*uint64(4))
	if !((*T_ARRAY32)(unsafe.Pointer(array)).Fdata != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint8((*T_ARRAY32)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpReadCARD32(tls, buffer, (*T_ARRAY32)(unsafe.Pointer(array)).Fdata+uintptr(i)*4) != 0) {
			libc.Xfree(tls, (*T_ARRAY32)(unsafe.Pointer(array)).Fdata)
			(*T_ARRAY32)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
			(*T_ARRAY32)(unsafe.Pointer(array)).Flength = uint8(0)
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpReadARRAYofARRAY8(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAYofARRAY8Ptr) (r int32) {
	var i TCARD8
	_ = i
	/*
	 * When returning FALSE, guarantee that array->data = 0.
	 * This allows the user to safely call XdmcpDisposeARRAYofARRAY8(array)
	 * regardless of the return value below.
	 * Note that XdmcpDisposeARRAY*(array) will call free(array->data),
	 * so we must guarantee that array->data is NULL or a malloced pointer.
	 */
	if !(XXdmcpReadCARD8(tls, buffer, array) != 0) {
		(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return m_FALSE
	}
	if !((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength != 0) {
		(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = libc.UintptrFromInt32(0)
		return int32(m_TRUE)
	}
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = libc.Xmalloc(tls, uint64((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)*uint64(16))
	if !((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata != 0) {
		return m_FALSE
	}
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpReadARRAY8(tls, buffer, (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata+uintptr(i)*16) != 0) {
			/*
			 * We must free all of the arrays allocated thus far in the loop
			 * and free array->data and finally set array->data = 0;
			 * The easiest way to do this is to reset the length and call
			 * XdmcpDisposeARRAYofARRAY8(array).
			 */
			(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength = i
			XXdmcpDisposeARRAYofARRAY8(tls, array)
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpReadCARD8(tls *libc.TLS, buffer TXdmcpBufferPtr, valuep TCARD8Ptr) (r int32) {
	var v1 int32
	var v2 uintptr
	_, _ = v1, v2
	if (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer >= (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fcount {
		return m_FALSE
	}
	v2 = buffer + 12
	v1 = *(*int32)(unsafe.Pointer(v2))
	*(*int32)(unsafe.Pointer(v2))++
	*(*TCARD8)(unsafe.Pointer(valuep)) = *(*TBYTE)(unsafe.Pointer((*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata + uintptr(v1)))
	return int32(m_TRUE)
}

func XXdmcpReadCARD16(tls *libc.TLS, buffer TXdmcpBufferPtr, valuep TCARD16Ptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* high at bp+0 */ TCARD8
	var _ /* low at bp+1 */ TCARD8
	if XXdmcpReadCARD8(tls, buffer, bp) != 0 && XXdmcpReadCARD8(tls, buffer, bp+1) != 0 {
		*(*TCARD16)(unsafe.Pointer(valuep)) = libc.Uint16FromInt32(libc.Int32FromUint16(uint16(*(*TCARD8)(unsafe.Pointer(bp))))<<int32(8) | libc.Int32FromUint16(uint16(*(*TCARD8)(unsafe.Pointer(bp + 1)))))
		return int32(m_TRUE)
	}
	return m_FALSE
}

func XXdmcpReadCARD32(tls *libc.TLS, buffer TXdmcpBufferPtr, valuep TCARD32Ptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* byte0 at bp+0 */ TCARD8
	var _ /* byte1 at bp+1 */ TCARD8
	var _ /* byte2 at bp+2 */ TCARD8
	var _ /* byte3 at bp+3 */ TCARD8
	if XXdmcpReadCARD8(tls, buffer, bp) != 0 && XXdmcpReadCARD8(tls, buffer, bp+1) != 0 && XXdmcpReadCARD8(tls, buffer, bp+2) != 0 && XXdmcpReadCARD8(tls, buffer, bp+3) != 0 {
		*(*TCARD32)(unsafe.Pointer(valuep)) = uint32(*(*TCARD8)(unsafe.Pointer(bp)))<<libc.Int32FromInt32(24) | uint32(*(*TCARD8)(unsafe.Pointer(bp + 1)))<<libc.Int32FromInt32(16) | uint32(*(*TCARD8)(unsafe.Pointer(bp + 2)))<<libc.Int32FromInt32(8) | uint32(*(*TCARD8)(unsafe.Pointer(bp + 3)))
		return int32(m_TRUE)
	}
	return m_FALSE
}

type Tauth_cblock = [8]uint8

type Tauth_wrapper_schedule = [16]Tauth_ks_struct

type Tauth_ks_struct = struct {
	F_1 Tauth_cblock
}

func XXdmcpUnwrap(tls *libc.TLS, input uintptr, wrapper uintptr, output uintptr, bytes int32) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var i, j, k, v3 int32
	var _ /* blocks at bp+8 */ [2][8]uint8
	var _ /* expand_wrapper at bp+24 */ [8]uint8
	var _ /* schedule at bp+32 */ Tauth_wrapper_schedule
	var _ /* tmp at bp+0 */ [8]uint8
	_, _, _, _ = i, j, k, v3
	X_XdmcpWrapperToOddParity(tls, wrapper, bp+24)
	X_XdmcpAuthSetup(tls, bp+24, bp+32)
	k = 0
	j = 0
	for {
		if !(j < bytes) {
			break
		}
		if bytes-j < int32(8) {
			return
		} /* bad input length */
		i = 0
		for {
			if !(i < int32(8)) {
				break
			}
			*(*uint8)(unsafe.Pointer(bp + 8 + uintptr(k)*8 + uintptr(i))) = *(*uint8)(unsafe.Pointer(input + uintptr(j+i)))
			goto _2
		_2:
			;
			i++
		}
		X_XdmcpAuthDoIt(tls, input+uintptr(j), bp, bp+32, 0)
		/* block chaining */
		if k == 0 {
			v3 = int32(1)
		} else {
			v3 = 0
		}
		k = v3
		i = 0
		for {
			if !(i < int32(8)) {
				break
			}
			if j == 0 {
				*(*uint8)(unsafe.Pointer(output + uintptr(j+i))) = (*(*[8]uint8)(unsafe.Pointer(bp)))[i]
			} else {
				*(*uint8)(unsafe.Pointer(output + uintptr(j+i))) = libc.Uint8FromInt32(libc.Int32FromUint8((*(*[8]uint8)(unsafe.Pointer(bp)))[i]) ^ libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(bp + 8 + uintptr(k)*8 + uintptr(i)))))
			}
			goto _4
		_4:
			;
			i++
		}
		goto _1
	_1:
		;
		j += int32(8)
	}
}

func XXdmcpWrap(tls *libc.TLS, input uintptr, wrapper uintptr, output uintptr, bytes int32) {
	bp := tls.Alloc(144)
	defer tls.Free(144)
	var i, j, len1 int32
	var _ /* expand_wrapper at bp+8 */ [8]uint8
	var _ /* schedule at bp+16 */ Tauth_wrapper_schedule
	var _ /* tmp at bp+0 */ [8]uint8
	_, _, _ = i, j, len1
	X_XdmcpWrapperToOddParity(tls, wrapper, bp+8)
	X_XdmcpAuthSetup(tls, bp+8, bp+16)
	j = 0
	for {
		if !(j < bytes) {
			break
		}
		len1 = int32(8)
		if bytes-j < len1 {
			len1 = bytes - j
		}
		/* block chaining */
		i = 0
		for {
			if !(i < len1) {
				break
			}
			if j == 0 {
				(*(*[8]uint8)(unsafe.Pointer(bp)))[i] = *(*uint8)(unsafe.Pointer(input + uintptr(i)))
			} else {
				(*(*[8]uint8)(unsafe.Pointer(bp)))[i] = libc.Uint8FromInt32(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(input + uintptr(j+i)))) ^ libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(output + uintptr(j-int32(8)+i)))))
			}
			goto _2
		_2:
			;
			i++
		}
		for {
			if !(i < int32(8)) {
				break
			}
			if j == 0 {
				(*(*[8]uint8)(unsafe.Pointer(bp)))[i] = uint8(0)
			} else {
				(*(*[8]uint8)(unsafe.Pointer(bp)))[i] = libc.Uint8FromInt32(0 ^ libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(output + uintptr(j-int32(8)+i)))))
			}
			goto _3
		_3:
			;
			i++
		}
		X_XdmcpAuthDoIt(tls, bp, output+uintptr(j), bp+16, int32(1))
		goto _1
	_1:
		;
		j += int32(8)
	}
}

/*
 * Given a 56 bit wrapper in XDMCP format, create a 56
 * bit wrapper in 7-bits + odd parity format
 */

func _OddParity(tls *libc.TLS, c uint8) (r int32) {
	c = libc.Uint8FromInt32(libc.Int32FromUint8(c) ^ libc.Int32FromUint8(c)>>int32(4))
	c = libc.Uint8FromInt32(libc.Int32FromUint8(c) ^ libc.Int32FromUint8(c)>>int32(2))
	c = libc.Uint8FromInt32(libc.Int32FromUint8(c) ^ libc.Int32FromUint8(c)>>int32(1))
	return ^libc.Int32FromUint8(c) & int32(0x1)
}

/*
 * Spread the 56 bit wrapper among 8 bytes, using the upper 7 bits
 * of each byte, and storing an odd parity bit in the low bit
 */

func X_XdmcpWrapperToOddParity(tls *libc.TLS, in uintptr, out uintptr) {
	var ashift, bshift, i int32
	var c uint8
	_, _, _, _ = ashift, bshift, c, i
	ashift = int32(7)
	bshift = int32(1)
	i = 0
	for {
		if !(i < int32(7)) {
			break
		}
		c = libc.Uint8FromInt32((libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(in + uintptr(i))))<<ashift | libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(in + uintptr(i+int32(1)))))>>bshift) & int32(0x7f))
		*(*uint8)(unsafe.Pointer(out + uintptr(i))) = libc.Uint8FromInt32(libc.Int32FromUint8(c)<<int32(1) | _OddParity(tls, c))
		ashift--
		bshift++
		goto _1
	_1:
		;
		i++
	}
	c = *(*uint8)(unsafe.Pointer(in + uintptr(i)))
	*(*uint8)(unsafe.Pointer(out + uintptr(i))) = libc.Uint8FromInt32(libc.Int32FromUint8(c)<<int32(1) | _OddParity(tls, c))
}

func XXdmcpWriteHeader(tls *libc.TLS, buffer TXdmcpBufferPtr, header TXdmcpHeaderPtr) (r int32) {
	var newData uintptr
	_ = newData
	if (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize < int32(6)+libc.Int32FromUint16((*T_XdmcpHeader)(unsafe.Pointer(header)).Flength) {
		newData = libc.Xmalloc(tls, libc.Uint64FromInt32(m_XDM_MAX_MSGLEN)*libc.Uint64FromInt64(1))
		if !(newData != 0) {
			return m_FALSE
		}
		libc.Xfree(tls, (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata)
		(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata = newData
		(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize = int32(m_XDM_MAX_MSGLEN)
	}
	(*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer = 0
	if !(XXdmcpWriteCARD16(tls, buffer, uint32((*T_XdmcpHeader)(unsafe.Pointer(header)).Fversion)) != 0) {
		return m_FALSE
	}
	if !(XXdmcpWriteCARD16(tls, buffer, uint32((*T_XdmcpHeader)(unsafe.Pointer(header)).Fopcode)) != 0) {
		return m_FALSE
	}
	if !(XXdmcpWriteCARD16(tls, buffer, uint32((*T_XdmcpHeader)(unsafe.Pointer(header)).Flength)) != 0) {
		return m_FALSE
	}
	return int32(m_TRUE)
}

func XXdmcpWriteARRAY8(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAY8Ptr) (r int32) {
	var i int32
	_ = i
	if !(XXdmcpWriteCARD16(tls, buffer, uint32((*T_ARRAY8)(unsafe.Pointer(array)).Flength)) != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint16((*T_ARRAY8)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpWriteCARD8(tls, buffer, uint32(*(*TCARD8)(unsafe.Pointer((*T_ARRAY8)(unsafe.Pointer(array)).Fdata + uintptr(i))))) != 0) {
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpWriteARRAY16(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAY16Ptr) (r int32) {
	var i int32
	_ = i
	if !(XXdmcpWriteCARD8(tls, buffer, uint32((*T_ARRAY16)(unsafe.Pointer(array)).Flength)) != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint8((*T_ARRAY16)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpWriteCARD16(tls, buffer, uint32(*(*TCARD16)(unsafe.Pointer((*T_ARRAY16)(unsafe.Pointer(array)).Fdata + uintptr(i)*2)))) != 0) {
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpWriteARRAY32(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAY32Ptr) (r int32) {
	var i int32
	_ = i
	if !(XXdmcpWriteCARD8(tls, buffer, uint32((*T_ARRAY32)(unsafe.Pointer(array)).Flength)) != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint8((*T_ARRAY32)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpWriteCARD32(tls, buffer, *(*TCARD32)(unsafe.Pointer((*T_ARRAY32)(unsafe.Pointer(array)).Fdata + uintptr(i)*4))) != 0) {
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpWriteARRAYofARRAY8(tls *libc.TLS, buffer TXdmcpBufferPtr, array TARRAYofARRAY8Ptr) (r int32) {
	var i int32
	_ = i
	if !(XXdmcpWriteCARD8(tls, buffer, uint32((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)) != 0) {
		return m_FALSE
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)) {
			break
		}
		if !(XXdmcpWriteARRAY8(tls, buffer, (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata+uintptr(i)*16) != 0) {
			return m_FALSE
		}
		goto _1
	_1:
		;
		i++
	}
	return int32(m_TRUE)
}

func XXdmcpWriteCARD8(tls *libc.TLS, buffer TXdmcpBufferPtr, value uint32) (r int32) {
	var v1 int32
	var v2 uintptr
	_, _ = v1, v2
	if (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fpointer >= (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize {
		return m_FALSE
	}
	v2 = buffer + 12
	v1 = *(*int32)(unsafe.Pointer(v2))
	*(*int32)(unsafe.Pointer(v2))++
	*(*TBYTE)(unsafe.Pointer((*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fdata + uintptr(v1))) = uint8(value)
	return int32(m_TRUE)
}

func XXdmcpWriteCARD16(tls *libc.TLS, buffer TXdmcpBufferPtr, value uint32) (r int32) {
	if !(XXdmcpWriteCARD8(tls, buffer, value>>int32(8)) != 0) {
		return m_FALSE
	}
	if !(XXdmcpWriteCARD8(tls, buffer, value&uint32(0xff)) != 0) {
		return m_FALSE
	}
	return int32(m_TRUE)
}

func XXdmcpWriteCARD32(tls *libc.TLS, buffer TXdmcpBufferPtr, value uint32) (r int32) {
	if !(XXdmcpWriteCARD8(tls, buffer, value>>int32(24)) != 0) {
		return m_FALSE
	}
	if !(XXdmcpWriteCARD8(tls, buffer, value>>int32(16)&uint32(0xff)) != 0) {
		return m_FALSE
	}
	if !(XXdmcpWriteCARD8(tls, buffer, value>>int32(8)&uint32(0xff)) != 0) {
		return m_FALSE
	}
	if !(XXdmcpWriteCARD8(tls, buffer, value&uint32(0xff)) != 0) {
		return m_FALSE
	}
	return int32(m_TRUE)
}

const m_HALF_ITERATIONS = 8
const m_ITERATIONS = 16

/* des routines for non-usa - eay 10/9/1991 eay@psych.psy.uq.oz.au
 * These routines were written for speed not size so they are bigger than
 * needed.  I have removed some of the loop unrolling, this will reduce
 * code size at the expense of some speed.
 * 25/9/1991 eay - much faster _XdmcpAuthSetup (4 times faster).
 * 19/9/1991 eay - cleaned up the IP and FP code.
 * 10/9/1991 eay - first release.
 * The des routines this file has been made from can be found in
 * ftp.psy.uq.oz.au /pub/DES
 * This particular version derived from OpenBSD Revsion 1.3.
 */

/*
 *
 * Export Requirements.
 * You may not export or re-export this software or any copy or
 * adaptation in violation of any applicable laws or regulations.
 *
 * Without limiting the generality of the foregoing, hardware, software,
 * technology or services provided under this license agreement may not
 * be exported, reexported, transferred or downloaded to or within (or to
 * a national resident of) countries under U.S. economic embargo
 * including the following countries:
 *
 * Cuba, Iran, Libya, North Korea, Sudan and Syria. This list is subject
 * to change.
 *
 * Hardware, software, technology or services may not be exported,
 * reexported, transferred or downloaded to persons or entities listed on
 * the U.S. Department of Commerce Denied Persons List, Entity List of
 * proliferation concern or on any U.S. Treasury Department Designated
 * Nationals exclusion list, or to parties directly or indirectly
 * involved in the development or production of nuclear, chemical,
 * biological weapons or in missile technology programs as specified in
 * the U.S. Export Administration Regulations (15 CFR 744).
 *
 * By accepting this license agreement you confirm that you are not
 * located in (or a national resident of) any country under U.S. economic
 * embargo, not identified on any U.S. Department of Commerce Denied
 * Persons List, Entity List or Treasury Department Designated Nationals
 * exclusion list, and not directly or indirectly involved in the
 * development or production of nuclear, chemical, biological weapons or
 * in missile technology programs as specified in the U.S. Export
 * Administration Regulations.
 *
 *
 * Local Country Import Requirements. The software you are about to
 * download contains cryptography technology. Some countries regulate the
 * import, use and/or export of certain products with cryptography. The
 * X.org Foundation makes no claims as to the applicability of local
 * country import, use and/or export regulations in relation to the
 * download of this product. If you are located outside the U.S. and
 * Canada you are advised to consult your local country regulations to
 * insure compliance.
 */

var _skb = [8][64]TCARD32{
	0: {
		1:  uint32(0x00000010),
		2:  uint32(0x20000000),
		3:  uint32(0x20000010),
		4:  uint32(0x00010000),
		5:  uint32(0x00010010),
		6:  uint32(0x20010000),
		7:  uint32(0x20010010),
		8:  uint32(0x00000800),
		9:  uint32(0x00000810),
		10: uint32(0x20000800),
		11: uint32(0x20000810),
		12: uint32(0x00010800),
		13: uint32(0x00010810),
		14: uint32(0x20010800),
		15: uint32(0x20010810),
		16: uint32(0x00000020),
		17: uint32(0x00000030),
		18: uint32(0x20000020),
		19: uint32(0x20000030),
		20: uint32(0x00010020),
		21: uint32(0x00010030),
		22: uint32(0x20010020),
		23: uint32(0x20010030),
		24: uint32(0x00000820),
		25: uint32(0x00000830),
		26: uint32(0x20000820),
		27: uint32(0x20000830),
		28: uint32(0x00010820),
		29: uint32(0x00010830),
		30: uint32(0x20010820),
		31: uint32(0x20010830),
		32: uint32(0x00080000),
		33: uint32(0x00080010),
		34: uint32(0x20080000),
		35: uint32(0x20080010),
		36: uint32(0x00090000),
		37: uint32(0x00090010),
		38: uint32(0x20090000),
		39: uint32(0x20090010),
		40: uint32(0x00080800),
		41: uint32(0x00080810),
		42: uint32(0x20080800),
		43: uint32(0x20080810),
		44: uint32(0x00090800),
		45: uint32(0x00090810),
		46: uint32(0x20090800),
		47: uint32(0x20090810),
		48: uint32(0x00080020),
		49: uint32(0x00080030),
		50: uint32(0x20080020),
		51: uint32(0x20080030),
		52: uint32(0x00090020),
		53: uint32(0x00090030),
		54: uint32(0x20090020),
		55: uint32(0x20090030),
		56: uint32(0x00080820),
		57: uint32(0x00080830),
		58: uint32(0x20080820),
		59: uint32(0x20080830),
		60: uint32(0x00090820),
		61: uint32(0x00090830),
		62: uint32(0x20090820),
		63: uint32(0x20090830),
	},
	1: {
		1:  uint32(0x02000000),
		2:  uint32(0x00002000),
		3:  uint32(0x02002000),
		4:  uint32(0x00200000),
		5:  uint32(0x02200000),
		6:  uint32(0x00202000),
		7:  uint32(0x02202000),
		8:  uint32(0x00000004),
		9:  uint32(0x02000004),
		10: uint32(0x00002004),
		11: uint32(0x02002004),
		12: uint32(0x00200004),
		13: uint32(0x02200004),
		14: uint32(0x00202004),
		15: uint32(0x02202004),
		16: uint32(0x00000400),
		17: uint32(0x02000400),
		18: uint32(0x00002400),
		19: uint32(0x02002400),
		20: uint32(0x00200400),
		21: uint32(0x02200400),
		22: uint32(0x00202400),
		23: uint32(0x02202400),
		24: uint32(0x00000404),
		25: uint32(0x02000404),
		26: uint32(0x00002404),
		27: uint32(0x02002404),
		28: uint32(0x00200404),
		29: uint32(0x02200404),
		30: uint32(0x00202404),
		31: uint32(0x02202404),
		32: uint32(0x10000000),
		33: uint32(0x12000000),
		34: uint32(0x10002000),
		35: uint32(0x12002000),
		36: uint32(0x10200000),
		37: uint32(0x12200000),
		38: uint32(0x10202000),
		39: uint32(0x12202000),
		40: uint32(0x10000004),
		41: uint32(0x12000004),
		42: uint32(0x10002004),
		43: uint32(0x12002004),
		44: uint32(0x10200004),
		45: uint32(0x12200004),
		46: uint32(0x10202004),
		47: uint32(0x12202004),
		48: uint32(0x10000400),
		49: uint32(0x12000400),
		50: uint32(0x10002400),
		51: uint32(0x12002400),
		52: uint32(0x10200400),
		53: uint32(0x12200400),
		54: uint32(0x10202400),
		55: uint32(0x12202400),
		56: uint32(0x10000404),
		57: uint32(0x12000404),
		58: uint32(0x10002404),
		59: uint32(0x12002404),
		60: uint32(0x10200404),
		61: uint32(0x12200404),
		62: uint32(0x10202404),
		63: uint32(0x12202404),
	},
	2: {
		1:  uint32(0x00000001),
		2:  uint32(0x00040000),
		3:  uint32(0x00040001),
		4:  uint32(0x01000000),
		5:  uint32(0x01000001),
		6:  uint32(0x01040000),
		7:  uint32(0x01040001),
		8:  uint32(0x00000002),
		9:  uint32(0x00000003),
		10: uint32(0x00040002),
		11: uint32(0x00040003),
		12: uint32(0x01000002),
		13: uint32(0x01000003),
		14: uint32(0x01040002),
		15: uint32(0x01040003),
		16: uint32(0x00000200),
		17: uint32(0x00000201),
		18: uint32(0x00040200),
		19: uint32(0x00040201),
		20: uint32(0x01000200),
		21: uint32(0x01000201),
		22: uint32(0x01040200),
		23: uint32(0x01040201),
		24: uint32(0x00000202),
		25: uint32(0x00000203),
		26: uint32(0x00040202),
		27: uint32(0x00040203),
		28: uint32(0x01000202),
		29: uint32(0x01000203),
		30: uint32(0x01040202),
		31: uint32(0x01040203),
		32: uint32(0x08000000),
		33: uint32(0x08000001),
		34: uint32(0x08040000),
		35: uint32(0x08040001),
		36: uint32(0x09000000),
		37: uint32(0x09000001),
		38: uint32(0x09040000),
		39: uint32(0x09040001),
		40: uint32(0x08000002),
		41: uint32(0x08000003),
		42: uint32(0x08040002),
		43: uint32(0x08040003),
		44: uint32(0x09000002),
		45: uint32(0x09000003),
		46: uint32(0x09040002),
		47: uint32(0x09040003),
		48: uint32(0x08000200),
		49: uint32(0x08000201),
		50: uint32(0x08040200),
		51: uint32(0x08040201),
		52: uint32(0x09000200),
		53: uint32(0x09000201),
		54: uint32(0x09040200),
		55: uint32(0x09040201),
		56: uint32(0x08000202),
		57: uint32(0x08000203),
		58: uint32(0x08040202),
		59: uint32(0x08040203),
		60: uint32(0x09000202),
		61: uint32(0x09000203),
		62: uint32(0x09040202),
		63: uint32(0x09040203),
	},
	3: {
		1:  uint32(0x00100000),
		2:  uint32(0x00000100),
		3:  uint32(0x00100100),
		4:  uint32(0x00000008),
		5:  uint32(0x00100008),
		6:  uint32(0x00000108),
		7:  uint32(0x00100108),
		8:  uint32(0x00001000),
		9:  uint32(0x00101000),
		10: uint32(0x00001100),
		11: uint32(0x00101100),
		12: uint32(0x00001008),
		13: uint32(0x00101008),
		14: uint32(0x00001108),
		15: uint32(0x00101108),
		16: uint32(0x04000000),
		17: uint32(0x04100000),
		18: uint32(0x04000100),
		19: uint32(0x04100100),
		20: uint32(0x04000008),
		21: uint32(0x04100008),
		22: uint32(0x04000108),
		23: uint32(0x04100108),
		24: uint32(0x04001000),
		25: uint32(0x04101000),
		26: uint32(0x04001100),
		27: uint32(0x04101100),
		28: uint32(0x04001008),
		29: uint32(0x04101008),
		30: uint32(0x04001108),
		31: uint32(0x04101108),
		32: uint32(0x00020000),
		33: uint32(0x00120000),
		34: uint32(0x00020100),
		35: uint32(0x00120100),
		36: uint32(0x00020008),
		37: uint32(0x00120008),
		38: uint32(0x00020108),
		39: uint32(0x00120108),
		40: uint32(0x00021000),
		41: uint32(0x00121000),
		42: uint32(0x00021100),
		43: uint32(0x00121100),
		44: uint32(0x00021008),
		45: uint32(0x00121008),
		46: uint32(0x00021108),
		47: uint32(0x00121108),
		48: uint32(0x04020000),
		49: uint32(0x04120000),
		50: uint32(0x04020100),
		51: uint32(0x04120100),
		52: uint32(0x04020008),
		53: uint32(0x04120008),
		54: uint32(0x04020108),
		55: uint32(0x04120108),
		56: uint32(0x04021000),
		57: uint32(0x04121000),
		58: uint32(0x04021100),
		59: uint32(0x04121100),
		60: uint32(0x04021008),
		61: uint32(0x04121008),
		62: uint32(0x04021108),
		63: uint32(0x04121108),
	},
	4: {
		1:  uint32(0x10000000),
		2:  uint32(0x00010000),
		3:  uint32(0x10010000),
		4:  uint32(0x00000004),
		5:  uint32(0x10000004),
		6:  uint32(0x00010004),
		7:  uint32(0x10010004),
		8:  uint32(0x20000000),
		9:  uint32(0x30000000),
		10: uint32(0x20010000),
		11: uint32(0x30010000),
		12: uint32(0x20000004),
		13: uint32(0x30000004),
		14: uint32(0x20010004),
		15: uint32(0x30010004),
		16: uint32(0x00100000),
		17: uint32(0x10100000),
		18: uint32(0x00110000),
		19: uint32(0x10110000),
		20: uint32(0x00100004),
		21: uint32(0x10100004),
		22: uint32(0x00110004),
		23: uint32(0x10110004),
		24: uint32(0x20100000),
		25: uint32(0x30100000),
		26: uint32(0x20110000),
		27: uint32(0x30110000),
		28: uint32(0x20100004),
		29: uint32(0x30100004),
		30: uint32(0x20110004),
		31: uint32(0x30110004),
		32: uint32(0x00001000),
		33: uint32(0x10001000),
		34: uint32(0x00011000),
		35: uint32(0x10011000),
		36: uint32(0x00001004),
		37: uint32(0x10001004),
		38: uint32(0x00011004),
		39: uint32(0x10011004),
		40: uint32(0x20001000),
		41: uint32(0x30001000),
		42: uint32(0x20011000),
		43: uint32(0x30011000),
		44: uint32(0x20001004),
		45: uint32(0x30001004),
		46: uint32(0x20011004),
		47: uint32(0x30011004),
		48: uint32(0x00101000),
		49: uint32(0x10101000),
		50: uint32(0x00111000),
		51: uint32(0x10111000),
		52: uint32(0x00101004),
		53: uint32(0x10101004),
		54: uint32(0x00111004),
		55: uint32(0x10111004),
		56: uint32(0x20101000),
		57: uint32(0x30101000),
		58: uint32(0x20111000),
		59: uint32(0x30111000),
		60: uint32(0x20101004),
		61: uint32(0x30101004),
		62: uint32(0x20111004),
		63: uint32(0x30111004),
	},
	5: {
		1:  uint32(0x08000000),
		2:  uint32(0x00000008),
		3:  uint32(0x08000008),
		4:  uint32(0x00000400),
		5:  uint32(0x08000400),
		6:  uint32(0x00000408),
		7:  uint32(0x08000408),
		8:  uint32(0x00020000),
		9:  uint32(0x08020000),
		10: uint32(0x00020008),
		11: uint32(0x08020008),
		12: uint32(0x00020400),
		13: uint32(0x08020400),
		14: uint32(0x00020408),
		15: uint32(0x08020408),
		16: uint32(0x00000001),
		17: uint32(0x08000001),
		18: uint32(0x00000009),
		19: uint32(0x08000009),
		20: uint32(0x00000401),
		21: uint32(0x08000401),
		22: uint32(0x00000409),
		23: uint32(0x08000409),
		24: uint32(0x00020001),
		25: uint32(0x08020001),
		26: uint32(0x00020009),
		27: uint32(0x08020009),
		28: uint32(0x00020401),
		29: uint32(0x08020401),
		30: uint32(0x00020409),
		31: uint32(0x08020409),
		32: uint32(0x02000000),
		33: uint32(0x0A000000),
		34: uint32(0x02000008),
		35: uint32(0x0A000008),
		36: uint32(0x02000400),
		37: uint32(0x0A000400),
		38: uint32(0x02000408),
		39: uint32(0x0A000408),
		40: uint32(0x02020000),
		41: uint32(0x0A020000),
		42: uint32(0x02020008),
		43: uint32(0x0A020008),
		44: uint32(0x02020400),
		45: uint32(0x0A020400),
		46: uint32(0x02020408),
		47: uint32(0x0A020408),
		48: uint32(0x02000001),
		49: uint32(0x0A000001),
		50: uint32(0x02000009),
		51: uint32(0x0A000009),
		52: uint32(0x02000401),
		53: uint32(0x0A000401),
		54: uint32(0x02000409),
		55: uint32(0x0A000409),
		56: uint32(0x02020001),
		57: uint32(0x0A020001),
		58: uint32(0x02020009),
		59: uint32(0x0A020009),
		60: uint32(0x02020401),
		61: uint32(0x0A020401),
		62: uint32(0x02020409),
		63: uint32(0x0A020409),
	},
	6: {
		1:  uint32(0x00000100),
		2:  uint32(0x00080000),
		3:  uint32(0x00080100),
		4:  uint32(0x01000000),
		5:  uint32(0x01000100),
		6:  uint32(0x01080000),
		7:  uint32(0x01080100),
		8:  uint32(0x00000010),
		9:  uint32(0x00000110),
		10: uint32(0x00080010),
		11: uint32(0x00080110),
		12: uint32(0x01000010),
		13: uint32(0x01000110),
		14: uint32(0x01080010),
		15: uint32(0x01080110),
		16: uint32(0x00200000),
		17: uint32(0x00200100),
		18: uint32(0x00280000),
		19: uint32(0x00280100),
		20: uint32(0x01200000),
		21: uint32(0x01200100),
		22: uint32(0x01280000),
		23: uint32(0x01280100),
		24: uint32(0x00200010),
		25: uint32(0x00200110),
		26: uint32(0x00280010),
		27: uint32(0x00280110),
		28: uint32(0x01200010),
		29: uint32(0x01200110),
		30: uint32(0x01280010),
		31: uint32(0x01280110),
		32: uint32(0x00000200),
		33: uint32(0x00000300),
		34: uint32(0x00080200),
		35: uint32(0x00080300),
		36: uint32(0x01000200),
		37: uint32(0x01000300),
		38: uint32(0x01080200),
		39: uint32(0x01080300),
		40: uint32(0x00000210),
		41: uint32(0x00000310),
		42: uint32(0x00080210),
		43: uint32(0x00080310),
		44: uint32(0x01000210),
		45: uint32(0x01000310),
		46: uint32(0x01080210),
		47: uint32(0x01080310),
		48: uint32(0x00200200),
		49: uint32(0x00200300),
		50: uint32(0x00280200),
		51: uint32(0x00280300),
		52: uint32(0x01200200),
		53: uint32(0x01200300),
		54: uint32(0x01280200),
		55: uint32(0x01280300),
		56: uint32(0x00200210),
		57: uint32(0x00200310),
		58: uint32(0x00280210),
		59: uint32(0x00280310),
		60: uint32(0x01200210),
		61: uint32(0x01200310),
		62: uint32(0x01280210),
		63: uint32(0x01280310),
	},
	7: {
		1:  uint32(0x04000000),
		2:  uint32(0x00040000),
		3:  uint32(0x04040000),
		4:  uint32(0x00000002),
		5:  uint32(0x04000002),
		6:  uint32(0x00040002),
		7:  uint32(0x04040002),
		8:  uint32(0x00002000),
		9:  uint32(0x04002000),
		10: uint32(0x00042000),
		11: uint32(0x04042000),
		12: uint32(0x00002002),
		13: uint32(0x04002002),
		14: uint32(0x00042002),
		15: uint32(0x04042002),
		16: uint32(0x00000020),
		17: uint32(0x04000020),
		18: uint32(0x00040020),
		19: uint32(0x04040020),
		20: uint32(0x00000022),
		21: uint32(0x04000022),
		22: uint32(0x00040022),
		23: uint32(0x04040022),
		24: uint32(0x00002020),
		25: uint32(0x04002020),
		26: uint32(0x00042020),
		27: uint32(0x04042020),
		28: uint32(0x00002022),
		29: uint32(0x04002022),
		30: uint32(0x00042022),
		31: uint32(0x04042022),
		32: uint32(0x00000800),
		33: uint32(0x04000800),
		34: uint32(0x00040800),
		35: uint32(0x04040800),
		36: uint32(0x00000802),
		37: uint32(0x04000802),
		38: uint32(0x00040802),
		39: uint32(0x04040802),
		40: uint32(0x00002800),
		41: uint32(0x04002800),
		42: uint32(0x00042800),
		43: uint32(0x04042800),
		44: uint32(0x00002802),
		45: uint32(0x04002802),
		46: uint32(0x00042802),
		47: uint32(0x04042802),
		48: uint32(0x00000820),
		49: uint32(0x04000820),
		50: uint32(0x00040820),
		51: uint32(0x04040820),
		52: uint32(0x00000822),
		53: uint32(0x04000822),
		54: uint32(0x00040822),
		55: uint32(0x04040822),
		56: uint32(0x00002820),
		57: uint32(0x04002820),
		58: uint32(0x00042820),
		59: uint32(0x04042820),
		60: uint32(0x00002822),
		61: uint32(0x04002822),
		62: uint32(0x00042822),
		63: uint32(0x04042822),
	},
}

var _SPtrans = [8][64]TCARD32{
	0: {
		0:  uint32(0x00410100),
		1:  uint32(0x00010000),
		2:  uint32(0x40400000),
		3:  uint32(0x40410100),
		4:  uint32(0x00400000),
		5:  uint32(0x40010100),
		6:  uint32(0x40010000),
		7:  uint32(0x40400000),
		8:  uint32(0x40010100),
		9:  uint32(0x00410100),
		10: uint32(0x00410000),
		11: uint32(0x40000100),
		12: uint32(0x40400100),
		13: uint32(0x00400000),
		15: uint32(0x40010000),
		16: uint32(0x00010000),
		17: uint32(0x40000000),
		18: uint32(0x00400100),
		19: uint32(0x00010100),
		20: uint32(0x40410100),
		21: uint32(0x00410000),
		22: uint32(0x40000100),
		23: uint32(0x00400100),
		24: uint32(0x40000000),
		25: uint32(0x00000100),
		26: uint32(0x00010100),
		27: uint32(0x40410000),
		28: uint32(0x00000100),
		29: uint32(0x40400100),
		30: uint32(0x40410000),
		33: uint32(0x40410100),
		34: uint32(0x00400100),
		35: uint32(0x40010000),
		36: uint32(0x00410100),
		37: uint32(0x00010000),
		38: uint32(0x40000100),
		39: uint32(0x00400100),
		40: uint32(0x40410000),
		41: uint32(0x00000100),
		42: uint32(0x00010100),
		43: uint32(0x40400000),
		44: uint32(0x40010100),
		45: uint32(0x40000000),
		46: uint32(0x40400000),
		47: uint32(0x00410000),
		48: uint32(0x40410100),
		49: uint32(0x00010100),
		50: uint32(0x00410000),
		51: uint32(0x40400100),
		52: uint32(0x00400000),
		53: uint32(0x40000100),
		54: uint32(0x40010000),
		56: uint32(0x00010000),
		57: uint32(0x00400000),
		58: uint32(0x40400100),
		59: uint32(0x00410100),
		60: uint32(0x40000000),
		61: uint32(0x40410000),
		62: uint32(0x00000100),
		63: uint32(0x40010100),
	},
	1: {
		0:  uint32(0x08021002),
		2:  uint32(0x00021000),
		3:  uint32(0x08020000),
		4:  uint32(0x08000002),
		5:  uint32(0x00001002),
		6:  uint32(0x08001000),
		7:  uint32(0x00021000),
		8:  uint32(0x00001000),
		9:  uint32(0x08020002),
		10: uint32(0x00000002),
		11: uint32(0x08001000),
		12: uint32(0x00020002),
		13: uint32(0x08021000),
		14: uint32(0x08020000),
		15: uint32(0x00000002),
		16: uint32(0x00020000),
		17: uint32(0x08001002),
		18: uint32(0x08020002),
		19: uint32(0x00001000),
		20: uint32(0x00021002),
		21: uint32(0x08000000),
		23: uint32(0x00020002),
		24: uint32(0x08001002),
		25: uint32(0x00021002),
		26: uint32(0x08021000),
		27: uint32(0x08000002),
		28: uint32(0x08000000),
		29: uint32(0x00020000),
		30: uint32(0x00001002),
		31: uint32(0x08021002),
		32: uint32(0x00020002),
		33: uint32(0x08021000),
		34: uint32(0x08001000),
		35: uint32(0x00021002),
		36: uint32(0x08021002),
		37: uint32(0x00020002),
		38: uint32(0x08000002),
		40: uint32(0x08000000),
		41: uint32(0x00001002),
		42: uint32(0x00020000),
		43: uint32(0x08020002),
		44: uint32(0x00001000),
		45: uint32(0x08000000),
		46: uint32(0x00021002),
		47: uint32(0x08001002),
		48: uint32(0x08021000),
		49: uint32(0x00001000),
		51: uint32(0x08000002),
		52: uint32(0x00000002),
		53: uint32(0x08021002),
		54: uint32(0x00021000),
		55: uint32(0x08020000),
		56: uint32(0x08020002),
		57: uint32(0x00020000),
		58: uint32(0x00001002),
		59: uint32(0x08001000),
		60: uint32(0x08001002),
		61: uint32(0x00000002),
		62: uint32(0x08020000),
		63: uint32(0x00021000),
	},
	2: {
		0:  uint32(0x20800000),
		1:  uint32(0x00808020),
		2:  uint32(0x00000020),
		3:  uint32(0x20800020),
		4:  uint32(0x20008000),
		5:  uint32(0x00800000),
		6:  uint32(0x20800020),
		7:  uint32(0x00008020),
		8:  uint32(0x00800020),
		9:  uint32(0x00008000),
		10: uint32(0x00808000),
		11: uint32(0x20000000),
		12: uint32(0x20808020),
		13: uint32(0x20000020),
		14: uint32(0x20000000),
		15: uint32(0x20808000),
		17: uint32(0x20008000),
		18: uint32(0x00808020),
		19: uint32(0x00000020),
		20: uint32(0x20000020),
		21: uint32(0x20808020),
		22: uint32(0x00008000),
		23: uint32(0x20800000),
		24: uint32(0x20808000),
		25: uint32(0x00800020),
		26: uint32(0x20008020),
		27: uint32(0x00808000),
		28: uint32(0x00008020),
		30: uint32(0x00800000),
		31: uint32(0x20008020),
		32: uint32(0x00808020),
		33: uint32(0x00000020),
		34: uint32(0x20000000),
		35: uint32(0x00008000),
		36: uint32(0x20000020),
		37: uint32(0x20008000),
		38: uint32(0x00808000),
		39: uint32(0x20800020),
		41: uint32(0x00808020),
		42: uint32(0x00008020),
		43: uint32(0x20808000),
		44: uint32(0x20008000),
		45: uint32(0x00800000),
		46: uint32(0x20808020),
		47: uint32(0x20000000),
		48: uint32(0x20008020),
		49: uint32(0x20800000),
		50: uint32(0x00800000),
		51: uint32(0x20808020),
		52: uint32(0x00008000),
		53: uint32(0x00800020),
		54: uint32(0x20800020),
		55: uint32(0x00008020),
		56: uint32(0x00800020),
		58: uint32(0x20808000),
		59: uint32(0x20000020),
		60: uint32(0x20800000),
		61: uint32(0x20008020),
		62: uint32(0x00000020),
		63: uint32(0x00808000),
	},
	3: {
		0:  uint32(0x00080201),
		1:  uint32(0x02000200),
		2:  uint32(0x00000001),
		3:  uint32(0x02080201),
		5:  uint32(0x02080000),
		6:  uint32(0x02000201),
		7:  uint32(0x00080001),
		8:  uint32(0x02080200),
		9:  uint32(0x02000001),
		10: uint32(0x02000000),
		11: uint32(0x00000201),
		12: uint32(0x02000001),
		13: uint32(0x00080201),
		14: uint32(0x00080000),
		15: uint32(0x02000000),
		16: uint32(0x02080001),
		17: uint32(0x00080200),
		18: uint32(0x00000200),
		19: uint32(0x00000001),
		20: uint32(0x00080200),
		21: uint32(0x02000201),
		22: uint32(0x02080000),
		23: uint32(0x00000200),
		24: uint32(0x00000201),
		26: uint32(0x00080001),
		27: uint32(0x02080200),
		28: uint32(0x02000200),
		29: uint32(0x02080001),
		30: uint32(0x02080201),
		31: uint32(0x00080000),
		32: uint32(0x02080001),
		33: uint32(0x00000201),
		34: uint32(0x00080000),
		35: uint32(0x02000001),
		36: uint32(0x00080200),
		37: uint32(0x02000200),
		38: uint32(0x00000001),
		39: uint32(0x02080000),
		40: uint32(0x02000201),
		42: uint32(0x00000200),
		43: uint32(0x00080001),
		45: uint32(0x02080001),
		46: uint32(0x02080200),
		47: uint32(0x00000200),
		48: uint32(0x02000000),
		49: uint32(0x02080201),
		50: uint32(0x00080201),
		51: uint32(0x00080000),
		52: uint32(0x02080201),
		53: uint32(0x00000001),
		54: uint32(0x02000200),
		55: uint32(0x00080201),
		56: uint32(0x00080001),
		57: uint32(0x00080200),
		58: uint32(0x02080000),
		59: uint32(0x02000201),
		60: uint32(0x00000201),
		61: uint32(0x02000000),
		62: uint32(0x02000001),
		63: uint32(0x02080200),
	},
	4: {
		0:  uint32(0x01000000),
		1:  uint32(0x00002000),
		2:  uint32(0x00000080),
		3:  uint32(0x01002084),
		4:  uint32(0x01002004),
		5:  uint32(0x01000080),
		6:  uint32(0x00002084),
		7:  uint32(0x01002000),
		8:  uint32(0x00002000),
		9:  uint32(0x00000004),
		10: uint32(0x01000004),
		11: uint32(0x00002080),
		12: uint32(0x01000084),
		13: uint32(0x01002004),
		14: uint32(0x01002080),
		16: uint32(0x00002080),
		17: uint32(0x01000000),
		18: uint32(0x00002004),
		19: uint32(0x00000084),
		20: uint32(0x01000080),
		21: uint32(0x00002084),
		23: uint32(0x01000004),
		24: uint32(0x00000004),
		25: uint32(0x01000084),
		26: uint32(0x01002084),
		27: uint32(0x00002004),
		28: uint32(0x01002000),
		29: uint32(0x00000080),
		30: uint32(0x00000084),
		31: uint32(0x01002080),
		32: uint32(0x01002080),
		33: uint32(0x01000084),
		34: uint32(0x00002004),
		35: uint32(0x01002000),
		36: uint32(0x00002000),
		37: uint32(0x00000004),
		38: uint32(0x01000004),
		39: uint32(0x01000080),
		40: uint32(0x01000000),
		41: uint32(0x00002080),
		42: uint32(0x01002084),
		44: uint32(0x00002084),
		45: uint32(0x01000000),
		46: uint32(0x00000080),
		47: uint32(0x00002004),
		48: uint32(0x01000084),
		49: uint32(0x00000080),
		51: uint32(0x01002084),
		52: uint32(0x01002004),
		53: uint32(0x01002080),
		54: uint32(0x00000084),
		55: uint32(0x00002000),
		56: uint32(0x00002080),
		57: uint32(0x01002004),
		58: uint32(0x01000080),
		59: uint32(0x00000084),
		60: uint32(0x00000004),
		61: uint32(0x00002084),
		62: uint32(0x01002000),
		63: uint32(0x01000004),
	},
	5: {
		0:  uint32(0x10000008),
		1:  uint32(0x00040008),
		3:  uint32(0x10040400),
		4:  uint32(0x00040008),
		5:  uint32(0x00000400),
		6:  uint32(0x10000408),
		7:  uint32(0x00040000),
		8:  uint32(0x00000408),
		9:  uint32(0x10040408),
		10: uint32(0x00040400),
		11: uint32(0x10000000),
		12: uint32(0x10000400),
		13: uint32(0x10000008),
		14: uint32(0x10040000),
		15: uint32(0x00040408),
		16: uint32(0x00040000),
		17: uint32(0x10000408),
		18: uint32(0x10040008),
		20: uint32(0x00000400),
		21: uint32(0x00000008),
		22: uint32(0x10040400),
		23: uint32(0x10040008),
		24: uint32(0x10040408),
		25: uint32(0x10040000),
		26: uint32(0x10000000),
		27: uint32(0x00000408),
		28: uint32(0x00000008),
		29: uint32(0x00040400),
		30: uint32(0x00040408),
		31: uint32(0x10000400),
		32: uint32(0x00000408),
		33: uint32(0x10000000),
		34: uint32(0x10000400),
		35: uint32(0x00040408),
		36: uint32(0x10040400),
		37: uint32(0x00040008),
		39: uint32(0x10000400),
		40: uint32(0x10000000),
		41: uint32(0x00000400),
		42: uint32(0x10040008),
		43: uint32(0x00040000),
		44: uint32(0x00040008),
		45: uint32(0x10040408),
		46: uint32(0x00040400),
		47: uint32(0x00000008),
		48: uint32(0x10040408),
		49: uint32(0x00040400),
		50: uint32(0x00040000),
		51: uint32(0x10000408),
		52: uint32(0x10000008),
		53: uint32(0x10040000),
		54: uint32(0x00040408),
		56: uint32(0x00000400),
		57: uint32(0x10000008),
		58: uint32(0x10000408),
		59: uint32(0x10040400),
		60: uint32(0x10040000),
		61: uint32(0x00000408),
		62: uint32(0x00000008),
		63: uint32(0x10040008),
	},
	6: {
		0:  uint32(0x00000800),
		1:  uint32(0x00000040),
		2:  uint32(0x00200040),
		3:  uint32(0x80200000),
		4:  uint32(0x80200840),
		5:  uint32(0x80000800),
		6:  uint32(0x00000840),
		8:  uint32(0x00200000),
		9:  uint32(0x80200040),
		10: uint32(0x80000040),
		11: uint32(0x00200800),
		12: uint32(0x80000000),
		13: uint32(0x00200840),
		14: uint32(0x00200800),
		15: uint32(0x80000040),
		16: uint32(0x80200040),
		17: uint32(0x00000800),
		18: uint32(0x80000800),
		19: uint32(0x80200840),
		21: uint32(0x00200040),
		22: uint32(0x80200000),
		23: uint32(0x00000840),
		24: uint32(0x80200800),
		25: uint32(0x80000840),
		26: uint32(0x00200840),
		27: uint32(0x80000000),
		28: uint32(0x80000840),
		29: uint32(0x80200800),
		30: uint32(0x00000040),
		31: uint32(0x00200000),
		32: uint32(0x80000840),
		33: uint32(0x00200800),
		34: uint32(0x80200800),
		35: uint32(0x80000040),
		36: uint32(0x00000800),
		37: uint32(0x00000040),
		38: uint32(0x00200000),
		39: uint32(0x80200800),
		40: uint32(0x80200040),
		41: uint32(0x80000840),
		42: uint32(0x00000840),
		44: uint32(0x00000040),
		45: uint32(0x80200000),
		46: uint32(0x80000000),
		47: uint32(0x00200040),
		49: uint32(0x80200040),
		50: uint32(0x00200040),
		51: uint32(0x00000840),
		52: uint32(0x80000040),
		53: uint32(0x00000800),
		54: uint32(0x80200840),
		55: uint32(0x00200000),
		56: uint32(0x00200840),
		57: uint32(0x80000000),
		58: uint32(0x80000800),
		59: uint32(0x80200840),
		60: uint32(0x80200000),
		61: uint32(0x00200840),
		62: uint32(0x00200800),
		63: uint32(0x80000800),
	},
	7: {
		0:  uint32(0x04100010),
		1:  uint32(0x04104000),
		2:  uint32(0x00004010),
		4:  uint32(0x04004000),
		5:  uint32(0x00100010),
		6:  uint32(0x04100000),
		7:  uint32(0x04104010),
		8:  uint32(0x00000010),
		9:  uint32(0x04000000),
		10: uint32(0x00104000),
		11: uint32(0x00004010),
		12: uint32(0x00104010),
		13: uint32(0x04004010),
		14: uint32(0x04000010),
		15: uint32(0x04100000),
		16: uint32(0x00004000),
		17: uint32(0x00104010),
		18: uint32(0x00100010),
		19: uint32(0x04004000),
		20: uint32(0x04104010),
		21: uint32(0x04000010),
		23: uint32(0x00104000),
		24: uint32(0x04000000),
		25: uint32(0x00100000),
		26: uint32(0x04004010),
		27: uint32(0x04100010),
		28: uint32(0x00100000),
		29: uint32(0x00004000),
		30: uint32(0x04104000),
		31: uint32(0x00000010),
		32: uint32(0x00100000),
		33: uint32(0x00004000),
		34: uint32(0x04000010),
		35: uint32(0x04104010),
		36: uint32(0x00004010),
		37: uint32(0x04000000),
		39: uint32(0x00104000),
		40: uint32(0x04100010),
		41: uint32(0x04004010),
		42: uint32(0x04004000),
		43: uint32(0x00100010),
		44: uint32(0x04104000),
		45: uint32(0x00000010),
		46: uint32(0x00100010),
		47: uint32(0x04004000),
		48: uint32(0x04104010),
		49: uint32(0x00100000),
		50: uint32(0x04100000),
		51: uint32(0x04000010),
		52: uint32(0x00104000),
		53: uint32(0x00004010),
		54: uint32(0x04004010),
		55: uint32(0x04100000),
		56: uint32(0x00000010),
		57: uint32(0x04104000),
		58: uint32(0x00104010),
		60: uint32(0x04000000),
		61: uint32(0x04100010),
		62: uint32(0x00004000),
		63: uint32(0x00104010),
	},
}

var _shifts2 = [16]uint8{
	2:  uint8(1),
	3:  uint8(1),
	4:  uint8(1),
	5:  uint8(1),
	6:  uint8(1),
	7:  uint8(1),
	9:  uint8(1),
	10: uint8(1),
	11: uint8(1),
	12: uint8(1),
	13: uint8(1),
	14: uint8(1),
}

func X_XdmcpAuthSetup(tls *libc.TLS, key uintptr, schedule uintptr) {
	var c, d, s, t TCARD32
	var i int32
	var in, k, v1, v10, v11, v2, v3, v4, v5, v6, v7, v8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, d, i, in, k, s, t, v1, v10, v11, v2, v3, v4, v5, v6, v7, v8
	k = schedule
	in = key
	v1 = in
	in++
	c = uint32(*(*TCARD8)(unsafe.Pointer(v1)))
	v2 = in
	in++
	c |= uint32(*(*TCARD8)(unsafe.Pointer(v2))) << int32(8)
	v3 = in
	in++
	c |= uint32(*(*TCARD8)(unsafe.Pointer(v3))) << int32(16)
	v4 = in
	in++
	c |= uint32(*(*TCARD8)(unsafe.Pointer(v4))) << int32(24)
	v5 = in
	in++
	d = uint32(*(*TCARD8)(unsafe.Pointer(v5)))
	v6 = in
	in++
	d |= uint32(*(*TCARD8)(unsafe.Pointer(v6))) << int32(8)
	v7 = in
	in++
	d |= uint32(*(*TCARD8)(unsafe.Pointer(v7))) << int32(16)
	v8 = in
	in++
	d |= uint32(*(*TCARD8)(unsafe.Pointer(v8))) << int32(24)
	/* do PC1 in 60 simple operations */
	t = (d>>libc.Int32FromInt32(4) ^ c) & libc.Uint32FromInt32(libc.Int32FromInt32(0x0f0f0f0f))
	c ^= t
	d ^= t << libc.Int32FromInt32(4)
	t = (c<<(libc.Int32FromInt32(16) - -libc.Int32FromInt32(2)) ^ c) & libc.Uint32FromUint32(0xcccc0000)
	c = c ^ t ^ t>>(libc.Int32FromInt32(16) - -libc.Int32FromInt32(2))
	t = (c<<(libc.Int32FromInt32(16) - -libc.Int32FromInt32(1)) ^ c) & libc.Uint32FromUint32(0xaaaa0000)
	c = c ^ t ^ t>>(libc.Int32FromInt32(16) - -libc.Int32FromInt32(1))
	t = (c<<(libc.Int32FromInt32(16)-libc.Int32FromInt32(8)) ^ c) & libc.Uint32FromInt32(libc.Int32FromInt32(0x00ff0000))
	c = c ^ t ^ t>>(libc.Int32FromInt32(16)-libc.Int32FromInt32(8))
	t = (c<<(libc.Int32FromInt32(16) - -libc.Int32FromInt32(1)) ^ c) & libc.Uint32FromUint32(0xaaaa0000)
	c = c ^ t ^ t>>(libc.Int32FromInt32(16) - -libc.Int32FromInt32(1))
	t = (d<<(libc.Int32FromInt32(16) - -libc.Int32FromInt32(8)) ^ d) & libc.Uint32FromUint32(0xff000000)
	d = d ^ t ^ t>>(libc.Int32FromInt32(16) - -libc.Int32FromInt32(8))
	t = (d<<(libc.Int32FromInt32(16)-libc.Int32FromInt32(8)) ^ d) & libc.Uint32FromInt32(libc.Int32FromInt32(0x00ff0000))
	d = d ^ t ^ t>>(libc.Int32FromInt32(16)-libc.Int32FromInt32(8))
	t = (d<<(libc.Int32FromInt32(16)-libc.Int32FromInt32(2)) ^ d) & libc.Uint32FromInt32(libc.Int32FromInt32(0x33330000))
	d = d ^ t ^ t>>(libc.Int32FromInt32(16)-libc.Int32FromInt32(2))
	d = d&uint32(0x00aa00aa)<<int32(7) | d&uint32(0x55005500)>>int32(7) | d&uint32(0xaa55aa55)
	d = d>>libc.Int32FromInt32(8) | c&uint32(0xf0000000)>>int32(4)
	c &= uint32(0x0fffffff)
	i = 0
	for {
		if !(i < int32(m_ITERATIONS)) {
			break
		}
		if _shifts2[i] != 0 {
			c = c>>libc.Int32FromInt32(2) | c<<libc.Int32FromInt32(26)
			d = d>>libc.Int32FromInt32(2) | d<<libc.Int32FromInt32(26)
		} else {
			c = c>>libc.Int32FromInt32(1) | c<<libc.Int32FromInt32(27)
			d = d>>libc.Int32FromInt32(1) | d<<libc.Int32FromInt32(27)
		}
		c &= uint32(0x0fffffff)
		d &= uint32(0x0fffffff)
		/* could be a few less shifts but I am to lazy at this
		 * point in time to investigate */
		s = *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + uintptr(c&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 1*256 + uintptr(c>>libc.Int32FromInt32(6)&uint32(0x03)|c>>libc.Int32FromInt32(7)&uint32(0x3c))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 2*256 + uintptr(c>>libc.Int32FromInt32(13)&uint32(0x0f)|c>>libc.Int32FromInt32(14)&uint32(0x30))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 3*256 + uintptr(c>>libc.Int32FromInt32(20)&uint32(0x01)|c>>libc.Int32FromInt32(21)&uint32(0x06)|c>>libc.Int32FromInt32(22)&uint32(0x38))*4))
		t = *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 4*256 + uintptr(d&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 5*256 + uintptr(d>>libc.Int32FromInt32(7)&uint32(0x03)|d>>libc.Int32FromInt32(8)&uint32(0x3c))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 6*256 + uintptr(d>>libc.Int32FromInt32(15)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_skb)) + 7*256 + uintptr(d>>libc.Int32FromInt32(21)&uint32(0x0f)|d>>libc.Int32FromInt32(22)&uint32(0x30))*4))
		/* table contained 0213 4657 */
		v10 = k
		k += 4
		*(*TCARD32)(unsafe.Pointer(v10)) = t<<libc.Int32FromInt32(16) | s&libc.Uint32FromInt32(0x0000ffff)
		s = s>>libc.Int32FromInt32(16) | t&libc.Uint32FromUint32(0xffff0000)
		s = s<<libc.Int32FromInt32(4) | s>>libc.Int32FromInt32(28)
		v11 = k
		k += 4
		*(*TCARD32)(unsafe.Pointer(v11)) = s
		goto _9
	_9:
		;
		i++
	}
	return
}

func X_XdmcpAuthDoIt(tls *libc.TLS, input uintptr, output uintptr, ks uintptr, encrypt int32) {
	var i int32
	var in, out, s, v1, v11, v12, v13, v14, v15, v16, v17, v18, v2, v3, v4, v5, v6, v7, v8 uintptr
	var l, r, t, u TCARD32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, in, l, out, r, s, t, u, v1, v11, v12, v13, v14, v15, v16, v17, v18, v2, v3, v4, v5, v6, v7, v8
	in = input
	out = output
	v1 = in
	in++
	l = uint32(*(*TCARD8)(unsafe.Pointer(v1)))
	v2 = in
	in++
	l |= uint32(*(*TCARD8)(unsafe.Pointer(v2))) << int32(8)
	v3 = in
	in++
	l |= uint32(*(*TCARD8)(unsafe.Pointer(v3))) << int32(16)
	v4 = in
	in++
	l |= uint32(*(*TCARD8)(unsafe.Pointer(v4))) << int32(24)
	v5 = in
	in++
	r = uint32(*(*TCARD8)(unsafe.Pointer(v5)))
	v6 = in
	in++
	r |= uint32(*(*TCARD8)(unsafe.Pointer(v6))) << int32(8)
	v7 = in
	in++
	r |= uint32(*(*TCARD8)(unsafe.Pointer(v7))) << int32(16)
	v8 = in
	in++
	r |= uint32(*(*TCARD8)(unsafe.Pointer(v8))) << int32(24)
	/* do IP */
	t = (r>>libc.Int32FromInt32(4) ^ l) & libc.Uint32FromInt32(libc.Int32FromInt32(0x0f0f0f0f))
	l ^= t
	r ^= t << libc.Int32FromInt32(4)
	t = (l>>libc.Int32FromInt32(16) ^ r) & libc.Uint32FromInt32(libc.Int32FromInt32(0x0000ffff))
	r ^= t
	l ^= t << libc.Int32FromInt32(16)
	t = (r>>libc.Int32FromInt32(2) ^ l) & libc.Uint32FromInt32(libc.Int32FromInt32(0x33333333))
	l ^= t
	r ^= t << libc.Int32FromInt32(2)
	t = (l>>libc.Int32FromInt32(8) ^ r) & libc.Uint32FromInt32(libc.Int32FromInt32(0x00ff00ff))
	r ^= t
	l ^= t << libc.Int32FromInt32(8)
	t = (r>>libc.Int32FromInt32(1) ^ l) & libc.Uint32FromInt32(libc.Int32FromInt32(0x55555555))
	l ^= t
	r ^= t << libc.Int32FromInt32(1)
	/* r and l are reversed - remember that :-) */
	t = l
	l = r
	r = t
	s = ks
	if encrypt != 0 {
		i = 0
		for {
			if !(i < libc.Int32FromInt32(m_ITERATIONS)*libc.Int32FromInt32(2)) {
				break
			}
			t = r<<libc.Int32FromInt32(1) | r>>libc.Int32FromInt32(31)
			u = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i)*4))
			t = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i+int32(1))*4))
			t = t>>libc.Int32FromInt32(4) | t<<libc.Int32FromInt32(28)
			l ^= *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 1*256 + uintptr(t&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 3*256 + uintptr(t>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 5*256 + uintptr(t>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 7*256 + uintptr(t>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + uintptr(u&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 2*256 + uintptr(u>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 4*256 + uintptr(u>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 6*256 + uintptr(u>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) /*	1 */
			t = l<<libc.Int32FromInt32(1) | l>>libc.Int32FromInt32(31)
			u = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i+int32(2))*4))
			t = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i+int32(2)+int32(1))*4))
			t = t>>libc.Int32FromInt32(4) | t<<libc.Int32FromInt32(28)
			r ^= *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 1*256 + uintptr(t&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 3*256 + uintptr(t>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 5*256 + uintptr(t>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 7*256 + uintptr(t>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + uintptr(u&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 2*256 + uintptr(u>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 4*256 + uintptr(u>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 6*256 + uintptr(u>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) /*  2 */
			goto _9
		_9:
			;
			i += int32(4)
		}
	} else {
		i = libc.Int32FromInt32(m_ITERATIONS)*libc.Int32FromInt32(2) - libc.Int32FromInt32(2)
		for {
			if !(i >= 0) {
				break
			}
			t = r<<libc.Int32FromInt32(1) | r>>libc.Int32FromInt32(31)
			u = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i)*4))
			t = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i+int32(1))*4))
			t = t>>libc.Int32FromInt32(4) | t<<libc.Int32FromInt32(28)
			l ^= *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 1*256 + uintptr(t&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 3*256 + uintptr(t>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 5*256 + uintptr(t>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 7*256 + uintptr(t>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + uintptr(u&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 2*256 + uintptr(u>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 4*256 + uintptr(u>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 6*256 + uintptr(u>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) /*	1 */
			t = l<<libc.Int32FromInt32(1) | l>>libc.Int32FromInt32(31)
			u = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i-int32(2))*4))
			t = t ^ *(*TCARD32)(unsafe.Pointer(s + uintptr(i-int32(2)+int32(1))*4))
			t = t>>libc.Int32FromInt32(4) | t<<libc.Int32FromInt32(28)
			r ^= *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 1*256 + uintptr(t&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 3*256 + uintptr(t>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 5*256 + uintptr(t>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 7*256 + uintptr(t>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + uintptr(u&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 2*256 + uintptr(u>>libc.Int32FromInt32(8)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 4*256 + uintptr(u>>libc.Int32FromInt32(16)&uint32(0x3f))*4)) | *(*TCARD32)(unsafe.Pointer(uintptr(unsafe.Pointer(&_SPtrans)) + 6*256 + uintptr(u>>libc.Int32FromInt32(24)&uint32(0x3f))*4)) /*  2 */
			goto _10
		_10:
			;
			i -= int32(4)
		}
	}
	/* swap l and r
	 * we will not do the swap so just remember they are
	 * reversed for the rest of the subroutine
	 * luckily by FP fixes this problem :-) */
	t = (r>>libc.Int32FromInt32(1) ^ l) & libc.Uint32FromInt32(libc.Int32FromInt32(0x55555555))
	l ^= t
	r ^= t << libc.Int32FromInt32(1)
	t = (l>>libc.Int32FromInt32(8) ^ r) & libc.Uint32FromInt32(libc.Int32FromInt32(0x00ff00ff))
	r ^= t
	l ^= t << libc.Int32FromInt32(8)
	t = (r>>libc.Int32FromInt32(2) ^ l) & libc.Uint32FromInt32(libc.Int32FromInt32(0x33333333))
	l ^= t
	r ^= t << libc.Int32FromInt32(2)
	t = (l>>libc.Int32FromInt32(16) ^ r) & libc.Uint32FromInt32(libc.Int32FromInt32(0x0000ffff))
	r ^= t
	l ^= t << libc.Int32FromInt32(16)
	t = (r>>libc.Int32FromInt32(4) ^ l) & libc.Uint32FromInt32(libc.Int32FromInt32(0x0f0f0f0f))
	l ^= t
	r ^= t << libc.Int32FromInt32(4)
	v11 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v11)) = uint8(l & libc.Uint32FromInt32(0xff))
	v12 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v12)) = uint8(l >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
	v13 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v13)) = uint8(l >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
	v14 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v14)) = uint8(l >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
	v15 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v15)) = uint8(r & libc.Uint32FromInt32(0xff))
	v16 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v16)) = uint8(r >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
	v17 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v17)) = uint8(r >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
	v18 = out
	out++
	*(*TCARD8)(unsafe.Pointer(v18)) = uint8(r >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
	return
}
