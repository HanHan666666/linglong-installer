// Code generated for darwin/arm64 by 'generator -no-main-minimize -DHAVE_LIBBSD --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -I /Users/jnml/src/modernc.org/builder/.exclude/modernc.org/libbsd/include/darwin/arm64 -lbsd -o libxdmcp.go --package-name libxdmcp .libs/libXdmcp.a', DO NOT EDIT.

//go:build darwin && arm64

package libxdmcp

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_ACCESSX_MAX_DESCRIPTORS = 100
const m_ATTRIBUTION_NAME_MAX = 255
const m_AT_EACCESS = 0x0010
const m_AT_FDONLY = 0x0400
const m_AT_REALDEV = 0x0200
const m_AT_REMOVEDIR = 0x0080
const m_AT_SYMLINK_FOLLOW = 0x0040
const m_AT_SYMLINK_NOFOLLOW = 0x0020
const m_AT_SYMLINK_NOFOLLOW_ANY = 0x0800
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
const m_BADSIG = "SIG_ERR"
const m_BC_BASE_MAX = 99
const m_BC_DIM_MAX = 2048
const m_BC_SCALE_MAX = 99
const m_BC_STRING_MAX = 1000
const m_BIG_ENDIAN = "__DARWIN_BIG_ENDIAN"
const m_BUS_ADRALN = 1
const m_BUS_ADRERR = 2
const m_BUS_NOOP = 0
const m_BUS_OBJERR = 3
const m_BYTE_ORDER = "__DARWIN_BYTE_ORDER"
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
const m_CHARCLASS_NAME_MAX = 14
const m_CHILD_MAX = 266
const m_CLD_CONTINUED = 6
const m_CLD_DUMPED = 3
const m_CLD_EXITED = 1
const m_CLD_KILLED = 2
const m_CLD_NOOP = 0
const m_CLD_STOPPED = 5
const m_CLD_TRAPPED = 4
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
const m_CPUMON_MAKE_FATAL = 0x1000
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
const m_DST_AUST = 2
const m_DST_CAN = 6
const m_DST_EET = 5
const m_DST_MET = 4
const m_DST_NONE = 0
const m_DST_USA = 1
const m_DST_WET = 3
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
const m_EQUIV_CLASS_MAX = 2
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_EXPR_NEST_MAX = 32
const m_EastGravity = 6
const m_EnableAccess = 1
const m_EnterNotify = 7
const m_EvenOddRule = 0
const m_Expose = 12
const m_FALSE = 0
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
const m_FOOTPRINT_INTERVAL_RESET = 0x1
const m_FPE_FLTDIV = 1
const m_FPE_FLTINV = 5
const m_FPE_FLTOVF = 2
const m_FPE_FLTRES = 4
const m_FPE_FLTSUB = 6
const m_FPE_FLTUND = 3
const m_FPE_INTDIV = 7
const m_FPE_INTOVF = 8
const m_FPE_NOOP = 0
const m_FREAD = 0x00000001
const m_FUNCPROTO = 15
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
const m_GETSIGSINFO_PLATFORM_BINARY = 1
const m_GID_MAX = 2147483647
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
const m_HAVE_GETENTROPY = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_LIBBSD = 1
const m_HAVE_LRAND48 = 1
const m_HAVE_SRAND48 = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDIO_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_RANDOM_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_HAVE_WCHAR_H = 1
const m_HostDelete = 1
const m_HostInsert = 0
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
const m_LINE_MAX = 2048
const m_LINK_MAX = 32767
const m_LITTLE_ENDIAN = "__DARWIN_LITTLE_ENDIAN"
const m_LOCK_EX = 0x02
const m_LOCK_NB = 0x04
const m_LOCK_SH = 0x01
const m_LOCK_UN = 0x08
const m_LSBFirst = 0
const m_LT_OBJDIR = ".libs/"
const m_L_INCR = "SEEK_CUR"
const m_L_SET = "SEEK_SET"
const m_L_XTND = "SEEK_END"
const m_L_ctermid = 1024
const m_LastExtensionError = 255
const m_LeaveNotify = 8
const m_LedModeOff = 0
const m_LedModeOn = 1
const m_LineDoubleDash = 2
const m_LineOnOffDash = 1
const m_LineSolid = 0
const m_LockMapIndex = 1
const m_LowerHighest = 1
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
const m_MAX_CANON = 1024
const m_MAX_INPUT = 1024
const m_MB_CUR_MAX = "__mb_cur_max"
const m_MINSIGSTKSZ = 32768
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
const m_NAME_MAX = 255
const m_NBBY = "__DARWIN_NBBY"
const m_NDEBUG = 1
const m_NFDBITS = "__DARWIN_NFDBITS"
const m_NGROUPS_MAX = 16
const m_NSIG = "__DARWIN_NSIG"
const m_NULL = "__DARWIN_NULL"
const m_NZERO = 20
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 1
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
const m_Opposite = 4
const m_PACKAGE = "libXdmcp"
const m_PACKAGE_BUGREPORT = "https://gitlab.freedesktop.org/xorg/lib/libxdmcp/-/issues"
const m_PACKAGE_NAME = "libXdmcp"
const m_PACKAGE_STRING = "libXdmcp 1.1.5"
const m_PACKAGE_TARNAME = "libXdmcp"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "1.1.5"
const m_PACKAGE_VERSION_MAJOR = 1
const m_PACKAGE_VERSION_MINOR = 1
const m_PACKAGE_VERSION_PATCHLEVEL = 5
const m_PATH_MAX = 1024
const m_PDP_ENDIAN = "__DARWIN_PDP_ENDIAN"
const m_PIPE_BUF = 512
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
const m_RE_DUP_MAX = 255
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
const m_RaiseLowest = 0
const m_ReparentNotify = 21
const m_ReplayKeyboard = 5
const m_ReplayPointer = 2
const m_ResizeRequest = 25
const m_RetainPermanent = 1
const m_RetainTemporary = 2
const m_RevertToParent = 2
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
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_SV_INTERRUPT = "SA_RESTART"
const m_SV_NOCLDSTOP = "SA_NOCLDSTOP"
const m_SV_NODEFER = "SA_NODEFER"
const m_SV_ONSTACK = "SA_ONSTACK"
const m_SV_RESETHAND = "SA_RESETHAND"
const m_SV_SIGINFO = "SA_SIGINFO"
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
const m_TIME_UTC = 1
const m_TRAP_BRKPT = 1
const m_TRAP_TRACE = 2
const m_TRUE = 1
const m_TileShape = 1
const m_TopIf = 2
const m_TrueColor = 4
const m_UID_MAX = 2147483647
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
const m_USER_FSIGNATURES_CDHASH_LEN = 20
const m_UnmapGravity = 0
const m_UnmapNotify = 18
const m_Unsorted = 0
const m_VERSION = "1.1.5"
const m_VisibilityFullyObscured = 2
const m_VisibilityNotify = 15
const m_VisibilityPartiallyObscured = 1
const m_VisibilityUnobscured = 0
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
const m_X_PROTOCOL = 11
const m_X_PROTOCOL_REVISION = 0
const m_YSorted = 1
const m_YXBanded = 3
const m_YXSorted = 2
const m_ZPixmap = 2
const m__ALL_SOURCE = 1
const m__ARM_SIGNAL_ = 1
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
const m__DARWIN_C_SOURCE = 1
const m__DARWIN_FEATURE_64_BIT_INODE = 1
const m__DARWIN_FEATURE_ONLY_64_BIT_INODE = 1
const m__DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
const m__DARWIN_FEATURE_ONLY_VERS_1050 = 1
const m__DARWIN_FEATURE_UNIX_CONFORMANCE = 3
const m__FORTIFY_SOURCE = 2
const m__GNU_SOURCE = 1
const m__HPUX_ALT_XOPEN_SOCKET_API = 1
const m__LIBC_COUNT__MB_LEN_MAX = "_LIBC_UNSAFE_INDEXABLE"
const m__LIBC_COUNT__PATH_MAX = "_LIBC_UNSAFE_INDEXABLE"
const m__LP64 = 1
const m__NETBSD_SOURCE = 1
const m__OPENBSD_SOURCE = 1
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
const m__POSIX_PTHREAD_SEMANTICS = 1
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
const m__QUAD_HIGHWORD = 1
const m__QUAD_LOWWORD = 0
const m__RLIMIT_POSIX_FLAG = 0x1000
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
const m__STRUCT_MCONTEXT = "_STRUCT_MCONTEXT64"
const m__TANDEM_SOURCE = 1
const m__V6_ILP32_OFF32 = "__ILP32_OFF32"
const m__V6_ILP32_OFFBIG = "__ILP32_OFFBIG"
const m__V6_LP64_OFF64 = "__LP64_OFF64"
const m__V6_LPBIG_OFFBIG = "__LPBIG_OFFBIG"
const m__WSTOPPED = 0177
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
const m___DARWIN_FD_SETSIZE = 1024
const m___DARWIN_LITTLE_ENDIAN = 1234
const m___DARWIN_NBBY = 8
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
const m___LP64_OFF64 = 1
const m___LP64__ = 1
const m___LPBIG_OFFBIG = 1
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
const m___SCHAR_MAX__ = 127
const m___SHRT_MAX__ = 32767
const m___SHRT_WIDTH__ = 16
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
const m___SSP__ = 1
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
const m___STDC_WANT_IEC_60559_EXT__ = 1
const m___STDC_WANT_IEC_60559_FUNCS_EXT__ = 1
const m___STDC_WANT_IEC_60559_TYPES_EXT__ = 1
const m___STDC_WANT_LIB_EXT1__ = 1
const m___STDC_WANT_LIB_EXT2__ = 1
const m___STDC_WANT_MATH_SPEC_FUNCS__ = 1
const m___STDC__ = 1
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
const m_sv_onstack = "sv_flags"

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint64

type t__predefined_wchar_t = int32

type t__predefined_ptrdiff_t = int64

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

type Tdev_t = int32

type Tfixpt_t = uint32

type Tblkcnt_t = int64

type Tblksize_t = int32

type Tgid_t = uint32

type Tin_addr_t = uint32

type Tin_port_t = uint16

type Tino_t = uint64

type Tino64_t = uint64

type Tkey_t = int32

type Tmode_t = uint16

type Tnlink_t = uint16

type Tid_t = uint32

type Tpid_t = int32

type Toff_t = int64

type Tsegsz_t = int32

type Tswblk_t = int32

type Tuid_t = uint32

type Tclock_t = uint64

type Tsize_t = uint64

type Tssize_t = int64

type Ttime_t = int64

type Tuseconds_t = uint32

type Tsuseconds_t = int32

type Trsize_t = uint64

type Terrno_t = int32

type Tfd_set = struct {
	Ffds_bits [32]t__int32_t
}

type Tfd_mask = int32

type Tpthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

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

type t__darwin_nl_item = int32

type t__darwin_wctrans_t = int32

type t__darwin_wctype_t = uint32

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

type Tuint64_t = uint64

type Tuint32_t = uint32

type Ttimeval = struct {
	Ftv_sec  t__darwin_time_t
	Ftv_usec t__darwin_suseconds_t
}

type Tsigset_t = uint32

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
const __CLOCK_THREAD_CPUTIME_ID = 16

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

type Tint_least64_t = int64

type Tuint_least64_t = uint64

type Tint_fast64_t = int64

type Tuint_fast64_t = uint64

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

type Tidtype_t = int32

const _P_ALL = 0
const _P_PID = 1
const _P_PGID = 2

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

type t__darwin_arm_neon_state64 = struct {
	F__ccgo_align [0]uint64
	F__v          [32][2]uint64
	F__fpsr       t__uint32_t
	F__fpcr       t__uint32_t
	F__ccgo_pad3  [8]byte
}

type t__darwin_arm_neon_state = struct {
	F__ccgo_align [0]uint64
	F__v          [16][2]uint64
	F__fpsr       t__uint32_t
	F__fpcr       t__uint32_t
	F__ccgo_pad3  [8]byte
}

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

type t__darwin_mcontext64 = struct {
	F__ccgo_align [0]uint64
	F__es         t__darwin_arm_exception_state64
	F__ss         t__darwin_arm_thread_state64
	F__ns         t__darwin_arm_neon_state64
}

type Tmcontext_t = uintptr

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

type Tmalloc_type_id_t = uint64 /* getsubopt(3) external variable */
/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

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
	if libc.Uint32FromInt32(length) > uint32(65535) || length < 0 {
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
	if libc.Uint32FromInt32(length) > uint32(255) || length < 0 {
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
	if libc.Uint32FromInt32(length) > uint32(255) || length < 0 {
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
	if libc.Uint32FromInt32(length) > uint32(255) || length < 0 {
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
	libc.X__builtin___memcpy_chk(tls, (*T_ARRAY8)(unsafe.Pointer(dst)).Fdata, (*T_ARRAY8)(unsafe.Pointer(src)).Fdata, uint64((*T_ARRAY8)(unsafe.Pointer(src)).Flength)*uint64(1), ^t__predefined_size_t(0))
	return int32(m_TRUE)
}

func XXdmcpReallocARRAY8(tls *libc.TLS, array TARRAY8Ptr, length int32) (r int32) {
	var newData TCARD8Ptr
	_ = newData
	/* length defined in ARRAY8 struct is a CARD16 (not CARD8 like the rest) */
	if libc.Uint32FromInt32(length) > uint32(65535) || length < 0 {
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
	if libc.Uint32FromInt32(length) > uint32(255) || length < 0 {
		return m_FALSE
	}
	newData = _xrealloc(tls, (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata, libc.Uint64FromInt32(length)*uint64(16))
	if !(newData != 0) {
		return m_FALSE
	}
	if length > libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength) {
		libc.X__builtin___memset_chk(tls, newData+uintptr((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)*16, 0, libc.Uint64FromInt32(length-libc.Int32FromUint8((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength))*uint64(16), ^t__predefined_size_t(0))
	}
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength = libc.Uint8FromInt32(length)
	(*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata = newData
	return int32(m_TRUE)
}

func XXdmcpReallocARRAY16(tls *libc.TLS, array TARRAY16Ptr, length int32) (r int32) {
	var newData TCARD16Ptr
	_ = newData
	/* length defined in ARRAY16 struct is a CARD8 */
	if libc.Uint32FromInt32(length) > uint32(255) || length < 0 {
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
	if libc.Uint32FromInt32(length) > uint32(255) || length < 0 {
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
	var i uint32
	_ = i
	if (*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Fdata != libc.UintptrFromInt32(0) {
		i = uint32(0)
		for {
			if !(i < uint32((*T_ARRAYofARRAY8)(unsafe.Pointer(array)).Flength)) {
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

const m_AF_APPLETALK = 16
const m_AF_CCITT = 10
const m_AF_CHAOS = 5
const m_AF_CNT = 21
const m_AF_COIP = 20
const m_AF_DATAKIT = 9
const m_AF_DECnet = 12
const m_AF_DLI = 13
const m_AF_E164 = "AF_ISDN"
const m_AF_ECMA = 8
const m_AF_HYLINK = 15
const m_AF_IEEE80211 = 37
const m_AF_IMPLINK = 3
const m_AF_INET = 2
const m_AF_INET6 = 30
const m_AF_IPX = 23
const m_AF_ISDN = 28
const m_AF_ISO = 7
const m_AF_LAT = 14
const m_AF_LINK = 18
const m_AF_LOCAL = "AF_UNIX"
const m_AF_MAX = 41
const m_AF_NATM = 31
const m_AF_NDRV = 27
const m_AF_NETBIOS = 33
const m_AF_NS = 6
const m_AF_OSI = "AF_ISO"
const m_AF_PPP = 34
const m_AF_PUP = 4
const m_AF_RESERVED_36 = 36
const m_AF_ROUTE = 17
const m_AF_SIP = 24
const m_AF_SNA = 11
const m_AF_SYSTEM = 32
const m_AF_UNIX = 1
const m_AF_UNSPEC = 0
const m_AF_UTUN = 38
const m_AF_VSOCK = 40
const m_CONNECT_DATA_AUTHENTICATED = 0x4
const m_CONNECT_DATA_IDEMPOTENT = 0x2
const m_CONNECT_RESUME_ON_READ_WRITE = 0x1
const m_KEV_DL_ADDMULTI = 7
const m_KEV_DL_AWDL_RESTRICTED = 26
const m_KEV_DL_AWDL_UNRESTRICTED = 27
const m_KEV_DL_DELMULTI = 8
const m_KEV_DL_IFCAP_CHANGED = 19
const m_KEV_DL_IFDELEGATE_CHANGED = 25
const m_KEV_DL_IF_ATTACHED = 9
const m_KEV_DL_IF_DETACHED = 11
const m_KEV_DL_IF_DETACHING = 10
const m_KEV_DL_IF_IDLE_ROUTE_REFCNT = 18
const m_KEV_DL_ISSUES = 24
const m_KEV_DL_LINK_ADDRESS_CHANGED = 16
const m_KEV_DL_LINK_OFF = 12
const m_KEV_DL_LINK_ON = 13
const m_KEV_DL_LINK_QUALITY_METRIC_CHANGED = 20
const m_KEV_DL_LOW_POWER_MODE_CHANGED = 30
const m_KEV_DL_NODE_ABSENCE = 22
const m_KEV_DL_NODE_PRESENCE = 21
const m_KEV_DL_PRIMARY_ELECTED = 23
const m_KEV_DL_PROTO_ATTACHED = 14
const m_KEV_DL_PROTO_DETACHED = 15
const m_KEV_DL_QOS_MODE_CHANGED = 29
const m_KEV_DL_RRC_STATE_CHANGED = 28
const m_KEV_DL_SIFFLAGS = 1
const m_KEV_DL_SIFGENERIC = 6
const m_KEV_DL_SIFMEDIA = 5
const m_KEV_DL_SIFMETRICS = 2
const m_KEV_DL_SIFMTU = 3
const m_KEV_DL_SIFPHYS = 4
const m_KEV_DL_SUBCLASS = 2
const m_KEV_DL_WAKEFLAGS_CHANGED = 17
const m_KEV_INET6_ADDR_DELETED = 3
const m_KEV_INET6_CHANGED_ADDR = 2
const m_KEV_INET6_DEFROUTER = 6
const m_KEV_INET6_NEW_LL_ADDR = 4
const m_KEV_INET6_NEW_RTADV_ADDR = 5
const m_KEV_INET6_NEW_USER_ADDR = 1
const m_KEV_INET6_REQUEST_NAT64_PREFIX = 7
const m_KEV_INET6_SUBCLASS = 6
const m_KEV_INET_ADDR_DELETED = 3
const m_KEV_INET_ARPCOLLISION = 7
const m_KEV_INET_ARPRTRALIVE = 10
const m_KEV_INET_ARPRTRFAILURE = 9
const m_KEV_INET_CHANGED_ADDR = 2
const m_KEV_INET_NEW_ADDR = 1
const m_KEV_INET_PORTINUSE = 8
const m_KEV_INET_SIFBRDADDR = 5
const m_KEV_INET_SIFDSTADDR = 4
const m_KEV_INET_SIFNETMASK = 6
const m_KEV_INET_SUBCLASS = 1
const m_MSG_CTRUNC = 0x20
const m_MSG_DONTROUTE = 0x4
const m_MSG_DONTWAIT = 0x80
const m_MSG_EOF = 0x100
const m_MSG_EOR = 0x8
const m_MSG_FLUSH = 0x400
const m_MSG_HAVEMORE = 0x2000
const m_MSG_HOLD = 0x800
const m_MSG_NEEDSA = 0x10000
const m_MSG_NOSIGNAL = 0x80000
const m_MSG_OOB = 0x1
const m_MSG_PEEK = 0x2
const m_MSG_RCVMORE = 0x4000
const m_MSG_SEND = 0x1000
const m_MSG_TRUNC = 0x10
const m_MSG_WAITALL = 0x40
const m_MSG_WAITSTREAM = 0x200
const m_NETSVC_MRKNG_LVL_L2 = 1
const m_NETSVC_MRKNG_LVL_L3L2_ALL = 2
const m_NETSVC_MRKNG_LVL_L3L2_BK = 3
const m_NETSVC_MRKNG_UNKNOWN = 0
const m_NET_MAXID = "AF_MAX"
const m_NET_RT_DUMP = 1
const m_NET_RT_DUMP2 = 7
const m_NET_RT_FLAGS = 2
const m_NET_RT_FLAGS_PRIV = 10
const m_NET_RT_IFLIST = 3
const m_NET_RT_IFLIST2 = 6
const m_NET_RT_MAXID = 11
const m_NET_RT_STAT = 4
const m_NET_RT_TRASH = 5
const m_NET_SERVICE_TYPE_AV = 6
const m_NET_SERVICE_TYPE_BE = 0
const m_NET_SERVICE_TYPE_BK = 1
const m_NET_SERVICE_TYPE_OAM = 7
const m_NET_SERVICE_TYPE_RD = 8
const m_NET_SERVICE_TYPE_RV = 5
const m_NET_SERVICE_TYPE_SIG = 2
const m_NET_SERVICE_TYPE_VI = 3
const m_NET_SERVICE_TYPE_VO = 4
const m_PF_APPLETALK = "AF_APPLETALK"
const m_PF_CCITT = "AF_CCITT"
const m_PF_CHAOS = "AF_CHAOS"
const m_PF_CNT = "AF_CNT"
const m_PF_COIP = "AF_COIP"
const m_PF_DATAKIT = "AF_DATAKIT"
const m_PF_DECnet = "AF_DECnet"
const m_PF_DLI = "AF_DLI"
const m_PF_ECMA = "AF_ECMA"
const m_PF_HYLINK = "AF_HYLINK"
const m_PF_IMPLINK = "AF_IMPLINK"
const m_PF_INET = "AF_INET"
const m_PF_INET6 = "AF_INET6"
const m_PF_IPX = "AF_IPX"
const m_PF_ISDN = "AF_ISDN"
const m_PF_ISO = "AF_ISO"
const m_PF_KEY = "pseudo_AF_KEY"
const m_PF_LAT = "AF_LAT"
const m_PF_LINK = "AF_LINK"
const m_PF_LOCAL = "AF_LOCAL"
const m_PF_MAX = "AF_MAX"
const m_PF_NATM = "AF_NATM"
const m_PF_NDRV = "AF_NDRV"
const m_PF_NETBIOS = "AF_NETBIOS"
const m_PF_NS = "AF_NS"
const m_PF_OSI = "AF_ISO"
const m_PF_PIP = "pseudo_AF_PIP"
const m_PF_PPP = "AF_PPP"
const m_PF_PUP = "AF_PUP"
const m_PF_RESERVED_36 = "AF_RESERVED_36"
const m_PF_ROUTE = "AF_ROUTE"
const m_PF_RTIP = "pseudo_AF_RTIP"
const m_PF_SIP = "AF_SIP"
const m_PF_SNA = "AF_SNA"
const m_PF_SYSTEM = "AF_SYSTEM"
const m_PF_UNIX = "PF_LOCAL"
const m_PF_UNSPEC = "AF_UNSPEC"
const m_PF_UTUN = "AF_UTUN"
const m_PF_VSOCK = "AF_VSOCK"
const m_PF_XTP = "pseudo_AF_XTP"
const m_SAE_ASSOCID_ANY = 0
const m_SAE_CONNID_ANY = 0
const m_SCM_CREDS = 0x03
const m_SCM_RIGHTS = 0x01
const m_SCM_TIMESTAMP = 0x02
const m_SCM_TIMESTAMP_MONOTONIC = 0x04
const m_SHUT_RD = 0
const m_SHUT_RDWR = 2
const m_SHUT_WR = 1
const m_SOCK_DGRAM = 2
const m_SOCK_MAXADDRLEN = 255
const m_SOCK_RAW = 3
const m_SOCK_RDM = 4
const m_SOCK_SEQPACKET = 5
const m_SOCK_STREAM = 1
const m_SOL_SOCKET = 0xffff
const m_SOMAXCONN = 128
const m_SONPX_SETOPTSHUT = 0x000000001
const m_SO_ACCEPTCONN = 0x0002
const m_SO_BINDTODEVICE = 0x1134
const m_SO_BROADCAST = 0x0020
const m_SO_DEBUG = 0x0001
const m_SO_DONTROUTE = 0x0010
const m_SO_DONTTRUNC = 0x2000
const m_SO_ERROR = 0x1007
const m_SO_KEEPALIVE = 0x0008
const m_SO_LABEL = 0x1010
const m_SO_LINGER = 0x0080
const m_SO_LINGER_SEC = 0x1080
const m_SO_NETSVC_MARKING_LEVEL = 0x1119
const m_SO_NET_SERVICE_TYPE = 0x1116
const m_SO_NKE = 0x1021
const m_SO_NOADDRERR = 0x1023
const m_SO_NOSIGPIPE = 0x1022
const m_SO_NOTIFYCONFLICT = 0x1026
const m_SO_NP_EXTENSIONS = 0x1083
const m_SO_NREAD = 0x1020
const m_SO_NUMRCVPKT = 0x1112
const m_SO_NWRITE = 0x1024
const m_SO_OOBINLINE = 0x0100
const m_SO_PEERLABEL = 0x1011
const m_SO_RANDOMPORT = 0x1082
const m_SO_RCVBUF = 0x1002
const m_SO_RCVLOWAT = 0x1004
const m_SO_RCVTIMEO = 0x1006
const m_SO_RESOLVER_SIGNATURE = 0x1131
const m_SO_REUSEADDR = 0x0004
const m_SO_REUSEPORT = 0x0200
const m_SO_REUSESHAREUID = 0x1025
const m_SO_SNDBUF = 0x1001
const m_SO_SNDLOWAT = 0x1003
const m_SO_SNDTIMEO = 0x1005
const m_SO_TIMESTAMP = 0x0400
const m_SO_TIMESTAMP_MONOTONIC = 0x0800
const m_SO_TYPE = 0x1008
const m_SO_UPCALLCLOSEWAIT = 0x1027
const m_SO_USELOOPBACK = 0x0040
const m_SO_WANTMORE = 0x4000
const m_SO_WANTOOBFLAG = 0x8000
const m__SS_MAXSIZE = 128
const m_pseudo_AF_HDRCMPLT = 35
const m_pseudo_AF_KEY = 29
const m_pseudo_AF_PIP = 25
const m_pseudo_AF_RTIP = 22
const m_pseudo_AF_XTP = 19

type Tsa_family_t = uint8

type Tsocklen_t = uint32

type Tiovec = struct {
	Fiov_base uintptr
	Fiov_len  Tsize_t
}

type Tsae_associd_t = uint32

type Tsae_connid_t = uint32

type Tsa_endpoints_t = struct {
	Fsae_srcif      uint32
	Fsae_srcaddr    uintptr
	Fsae_srcaddrlen Tsocklen_t
	Fsae_dstaddr    uintptr
	Fsae_dstaddrlen Tsocklen_t
}

type Tsa_endpoints = Tsa_endpoints_t

type Tlinger = struct {
	Fl_onoff  int32
	Fl_linger int32
}

type Tso_np_extensions = struct {
	Fnpx_flags Tu_int32_t
	Fnpx_mask  Tu_int32_t
}

type Tsockaddr = struct {
	Fsa_len    t__uint8_t
	Fsa_family Tsa_family_t
	Fsa_data   [14]int8
}

type t__sockaddr_header = struct {
	Fsa_len    t__uint8_t
	Fsa_family Tsa_family_t
}

type Tsockproto = struct {
	Fsp_family   t__uint16_t
	Fsp_protocol t__uint16_t
}

type Tsockaddr_storage = struct {
	Fss_len     t__uint8_t
	Fss_family  Tsa_family_t
	F__ss_pad1  [6]int8
	F__ss_align t__int64_t
	F__ss_pad2  [112]int8
}

type Tmsghdr = struct {
	Fmsg_name       uintptr
	Fmsg_namelen    Tsocklen_t
	Fmsg_iov        uintptr
	Fmsg_iovlen     int32
	Fmsg_control    uintptr
	Fmsg_controllen Tsocklen_t
	Fmsg_flags      int32
}

type Tcmsghdr = struct {
	Fcmsg_len   Tsocklen_t
	Fcmsg_level int32
	Fcmsg_type  int32
}

type Tsf_hdtr = struct {
	Fheaders  uintptr
	Fhdr_cnt  int32
	Ftrailers uintptr
	Ftrl_cnt  int32
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

/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

func XXdmcpGenerateKey(tls *libc.TLS, key TXdmAuthKeyPtr) {
	libc.Xarc4random_buf(tls, key, libc.Uint64FromInt32(8))
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

/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

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

/* valloc is now declared in _malloc.h */

/* Poison the following routines if -fshort-wchar is set */

func XXdmcpWriteHeader(tls *libc.TLS, buffer TXdmcpBufferPtr, header TXdmcpHeaderPtr) (r int32) {
	var newData uintptr
	_ = newData
	if (*T_XdmcpBuffer)(unsafe.Pointer(buffer)).Fsize < int32(6)+libc.Int32FromUint16((*T_XdmcpHeader)(unsafe.Pointer(header)).Flength) {
		newData = libc.Xcalloc(tls, uint64(m_XDM_MAX_MSGLEN), uint64(1))
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
 * This particular version derived from OpenBSD Revision 1.3.
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

var _shifts2 = [16]int8{
	2:  int8(1),
	3:  int8(1),
	4:  int8(1),
	5:  int8(1),
	6:  int8(1),
	7:  int8(1),
	9:  int8(1),
	10: int8(1),
	11: int8(1),
	12: int8(1),
	13: int8(1),
	14: int8(1),
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
