// Code generated for freebsd/amd64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -I /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libX11/include/freebsd/amd64 -lX11 -o libxrender.go --package-name libxrender src/.libs/libXrender.a', DO NOT EDIT.

//go:build freebsd && amd64

package libxrender

import (
	"reflect"
	"unsafe"

	"modernc.org/libX11"
	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_Above = 0
const m_AllTemporary = 0
const m_AllValues = 0x000F
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
const m_BIG_ENDIAN = "_BIG_ENDIAN"
const m_BUFSIZE = 2048
const m_BYTE_ORDER = "_BYTE_ORDER"
const m_BadAccess = 10
const m_BadAlloc = 11
const m_BadAtom = 5
const m_BadColor = 12
const m_BadCursor = 6
const m_BadDrawable = 9
const m_BadFont = 7
const m_BadGC = 13
const m_BadGlyph = 4
const m_BadGlyphSet = 3
const m_BadIDChoice = 14
const m_BadImplementation = 17
const m_BadLength = 16
const m_BadMatch = 8
const m_BadName = 15
const m_BadPictFormat = 0
const m_BadPictOp = 2
const m_BadPicture = 1
const m_BadPixmap = 4
const m_BadRequest = 1
const m_BadValue = 2
const m_BadWindow = 3
const m_Below = 1
const m_BitmapFileInvalid = 2
const m_BitmapNoMemory = 3
const m_BitmapOpenFailed = 1
const m_BitmapSuccess = 0
const m_Bool = "int"
const m_BottomIf = 3
const m_Button1 = 1
const m_Button2 = 2
const m_Button3 = 3
const m_Button4 = 4
const m_Button5 = 5
const m_ButtonPress = 4
const m_ButtonRelease = 5
const m_CLK_TCK = 128
const m_CLOCKS_PER_SEC = 128
const m_CLOCK_BOOTTIME = "CLOCK_UPTIME"
const m_CLOCK_MONOTONIC = 4
const m_CLOCK_MONOTONIC_COARSE = "CLOCK_MONOTONIC_FAST"
const m_CLOCK_MONOTONIC_FAST = 12
const m_CLOCK_MONOTONIC_PRECISE = 11
const m_CLOCK_PROCESS_CPUTIME_ID = 15
const m_CLOCK_PROF = 2
const m_CLOCK_REALTIME = 0
const m_CLOCK_REALTIME_COARSE = "CLOCK_REALTIME_FAST"
const m_CLOCK_REALTIME_FAST = 10
const m_CLOCK_REALTIME_PRECISE = 9
const m_CLOCK_SECOND = 13
const m_CLOCK_THREAD_CPUTIME_ID = 14
const m_CLOCK_UPTIME = 5
const m_CLOCK_UPTIME_FAST = 8
const m_CLOCK_UPTIME_PRECISE = 7
const m_CLOCK_VIRTUAL = 1
const m_CPLastBit = 12
const m_CPUSET_DEFAULT = 0
const m_CPU_LEVEL_CPUSET = 2
const m_CPU_LEVEL_ROOT = 1
const m_CPU_LEVEL_WHICH = 3
const m_CPU_MAXSIZE = 1024
const m_CPU_SETSIZE = "CPU_MAXSIZE"
const m_CPU_WHICH_CPUSET = 3
const m_CPU_WHICH_DOMAIN = 6
const m_CPU_WHICH_INTRHANDLER = 7
const m_CPU_WHICH_IRQ = 4
const m_CPU_WHICH_ITHREAD = 8
const m_CPU_WHICH_JAIL = 5
const m_CPU_WHICH_PID = 2
const m_CPU_WHICH_TID = 1
const m_CPU_WHICH_TIDPID = 9
const m_CURSORFONT = "cursor"
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
const m_DefaultBlanking = 2
const m_DefaultExposures = 2
const m_DestroyAll = 0
const m_DestroyNotify = 17
const m_DirectColor = 5
const m_DisableAccess = 0
const m_DisableScreenInterval = 0
const m_DisableScreenSaver = 0
const m_DontAllowExposures = 0
const m_DontCareState = 0
const m_DontPreferBlanking = 0
const m_E2BIG = 7
const m_EACCES = 13
const m_EADDRINUSE = 48
const m_EADDRNOTAVAIL = 49
const m_EAFNOSUPPORT = 47
const m_EAGAIN = 35
const m_EALREADY = 37
const m_EAUTH = 80
const m_EBADF = 9
const m_EBADMSG = 89
const m_EBADRPC = 72
const m_EBUSY = 16
const m_ECANCELED = 85
const m_ECAPMODE = 94
const m_ECHILD = 10
const m_ECONNABORTED = 53
const m_ECONNREFUSED = 61
const m_ECONNRESET = 54
const m_EDEADLK = 11
const m_EDESTADDRREQ = 39
const m_EDOM = 33
const m_EDOOFUS = 88
const m_EDQUOT = 69
const m_EEXIST = 17
const m_EFAULT = 14
const m_EFBIG = 27
const m_EFTYPE = 79
const m_EHOSTDOWN = 64
const m_EHOSTUNREACH = 65
const m_EIDRM = 82
const m_EILSEQ = 86
const m_EINPROGRESS = 36
const m_EINTEGRITY = 97
const m_EINTR = 4
const m_EINVAL = 22
const m_EIO = 5
const m_EISCONN = 56
const m_EISDIR = 21
const m_ELAST = 97
const m_ELOOP = 62
const m_EMFILE = 24
const m_EMLINK = 31
const m_EMSGSIZE = 40
const m_EMULTIHOP = 90
const m_ENAMETOOLONG = 63
const m_ENEEDAUTH = 81
const m_ENETDOWN = 50
const m_ENETRESET = 52
const m_ENETUNREACH = 51
const m_ENFILE = 23
const m_ENOATTR = 87
const m_ENOBUFS = 55
const m_ENODEV = 19
const m_ENOENT = 2
const m_ENOEXEC = 8
const m_ENOLCK = 77
const m_ENOLINK = 91
const m_ENOMEM = 12
const m_ENOMSG = 83
const m_ENOPROTOOPT = 42
const m_ENOSPC = 28
const m_ENOSYS = 78
const m_ENOTBLK = 15
const m_ENOTCAPABLE = 93
const m_ENOTCONN = 57
const m_ENOTDIR = 20
const m_ENOTEMPTY = 66
const m_ENOTRECOVERABLE = 95
const m_ENOTSOCK = 38
const m_ENOTSUP = "EOPNOTSUPP"
const m_ENOTTY = 25
const m_ENXIO = 6
const m_EOPNOTSUPP = 45
const m_EOVERFLOW = 84
const m_EOWNERDEAD = 96
const m_EPERM = 1
const m_EPFNOSUPPORT = 46
const m_EPIPE = 32
const m_EPROCLIM = 67
const m_EPROCUNAVAIL = 76
const m_EPROGMISMATCH = 75
const m_EPROGUNAVAIL = 74
const m_EPROTO = 92
const m_EPROTONOSUPPORT = 43
const m_EPROTOTYPE = 41
const m_ERANGE = 34
const m_EREMOTE = 71
const m_EROFS = 30
const m_ERPCMISMATCH = 73
const m_ESHUTDOWN = 58
const m_ESOCKTNOSUPPORT = 44
const m_ESPIPE = 29
const m_ESRCH = 3
const m_ESTALE = 70
const m_ETIMEDOUT = 60
const m_ETOOMANYREFS = 59
const m_ETXTBSY = 26
const m_EUSERS = 68
const m_EWOULDBLOCK = "EAGAIN"
const m_EXDEV = 18
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_EastGravity = 6
const m_EnableAccess = 1
const m_EnterNotify = 7
const m_EvenOddRule = 0
const m_Expose = 12
const m_FARCSPERBATCH = 256
const m_FD_SETSIZE = 1024
const m_FRCTSPERBATCH = 256
const m_FUNCPROTO = 15
const m_False = 0
const m_FamilyChaos = 2
const m_FamilyDECnet = 1
const m_FamilyInternet = 0
const m_FamilyInternet6 = 6
const m_FamilyServerInterpreted = 5
const m_FillOpaqueStippled = 3
const m_FillSolid = 0
const m_FillStippled = 2
const m_FillTiled = 1
const m_FilterBest = "best"
const m_FilterBilinear = "bilinear"
const m_FilterConvolution = "convolution"
const m_FilterFast = "fast"
const m_FilterGood = "good"
const m_FilterNearest = "nearest"
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
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_MEMORY_H = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_HeightValue = 0x0008
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
const m_INT_FAST8_MAX = "INT32_MAX"
const m_INT_FAST8_MIN = "INT32_MIN"
const m_INT_LEAST16_MAX = "INT16_MAX"
const m_INT_LEAST16_MIN = "INT16_MIN"
const m_INT_LEAST32_MAX = "INT32_MAX"
const m_INT_LEAST32_MIN = "INT32_MIN"
const m_INT_LEAST64_MAX = "INT64_MAX"
const m_INT_LEAST64_MIN = "INT64_MIN"
const m_INT_LEAST8_MAX = "INT8_MAX"
const m_INT_LEAST8_MIN = "INT8_MIN"
const m_IconicState = 3
const m_InactiveState = 4
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
const m_LITTLE_ENDIAN = "_LITTLE_ENDIAN"
const m_LOCKED = 1
const m_LSBFirst = 0
const m_LT_OBJDIR = ".libs/"
const m_LastExtensionError = 255
const m_LeaveNotify = 8
const m_LedModeOff = 0
const m_LedModeOn = 1
const m_LineDoubleDash = 2
const m_LineOnOffDash = 1
const m_LineSolid = 0
const m_LockMapIndex = 1
const m_LowerHighest = 1
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
const m_NFDBITS = "_NFDBITS"
const m_NLOCAL = 256
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 0
const m_NoEventMask = 0
const m_NoExpose = 14
const m_NoSymbol = 0
const m_NoValue = 0x0000
const m_Nonconvex = 1
const m_None = 0
const m_NormalState = 1
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
const m_Opposite = 4
const m_PACKAGE = "libXrender"
const m_PACKAGE_BUGREPORT = "https://bugs.freedesktop.org/enter_bug.cgi?product=xorg"
const m_PACKAGE_NAME = "libXrender"
const m_PACKAGE_STRING = "libXrender 0.9.10"
const m_PACKAGE_TARNAME = "libXrender"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "0.9.10"
const m_PACKAGE_VERSION_MAJOR = 0
const m_PACKAGE_VERSION_MINOR = 9
const m_PACKAGE_VERSION_PATCHLEVEL = 10
const m_PDP_ENDIAN = "_PDP_ENDIAN"
const m_PTHREAD_CANCEL_ASYNCHRONOUS = 2
const m_PTHREAD_CANCEL_DEFERRED = 0
const m_PTHREAD_CANCEL_DISABLE = 1
const m_PTHREAD_CANCEL_ENABLE = 0
const m_PTHREAD_COND_INITIALIZER = "NULL"
const m_PTHREAD_CREATE_DETACHED = "PTHREAD_DETACHED"
const m_PTHREAD_CREATE_JOINABLE = 0
const m_PTHREAD_DESTRUCTOR_ITERATIONS = 4
const m_PTHREAD_DETACHED = 0x1
const m_PTHREAD_DONE_INIT = 1
const m_PTHREAD_EXPLICIT_SCHED = 0
const m_PTHREAD_INHERIT_SCHED = 0x4
const m_PTHREAD_KEYS_MAX = 256
const m_PTHREAD_MUTEX_DEFAULT = "PTHREAD_MUTEX_ERRORCHECK"
const m_PTHREAD_MUTEX_INITIALIZER = "NULL"
const m_PTHREAD_MUTEX_ROBUST = 1
const m_PTHREAD_MUTEX_STALLED = 0
const m_PTHREAD_NEEDS_INIT = 0
const m_PTHREAD_NOFLOAT = 0x8
const m_PTHREAD_PRIO_INHERIT = 1
const m_PTHREAD_PRIO_NONE = 0
const m_PTHREAD_PRIO_PROTECT = 2
const m_PTHREAD_PROCESS_PRIVATE = 0
const m_PTHREAD_PROCESS_SHARED = 1
const m_PTHREAD_RWLOCK_INITIALIZER = "NULL"
const m_PTHREAD_SCOPE_PROCESS = 0
const m_PTHREAD_SCOPE_SYSTEM = 0x2
const m_PTHREAD_STACK_MIN = "__MINSIGSTKSZ"
const m_PTHREAD_THREADS_MAX = "__ULONG_MAX"
const m_PTRDIFF_MAX = "INT64_MAX"
const m_PTRDIFF_MIN = "INT64_MIN"
const m_PTSPERBATCH = 1024
const m_ParentRelative = 1
const m_PictOpAdd = 12
const m_PictOpAtop = 9
const m_PictOpAtopReverse = 10
const m_PictOpBlendMaximum = 0x3e
const m_PictOpBlendMinimum = 0x30
const m_PictOpClear = 0
const m_PictOpColorBurn = 0x36
const m_PictOpColorDodge = 0x35
const m_PictOpConjointAtop = 0x29
const m_PictOpConjointAtopReverse = 0x2a
const m_PictOpConjointClear = 0x20
const m_PictOpConjointDst = 0x22
const m_PictOpConjointIn = 0x25
const m_PictOpConjointInReverse = 0x26
const m_PictOpConjointMaximum = 0x2b
const m_PictOpConjointMinimum = 0x20
const m_PictOpConjointOut = 0x27
const m_PictOpConjointOutReverse = 0x28
const m_PictOpConjointOver = 0x23
const m_PictOpConjointOverReverse = 0x24
const m_PictOpConjointSrc = 0x21
const m_PictOpConjointXor = 0x2b
const m_PictOpDarken = 0x33
const m_PictOpDifference = 0x39
const m_PictOpDisjointAtop = 0x19
const m_PictOpDisjointAtopReverse = 0x1a
const m_PictOpDisjointClear = 0x10
const m_PictOpDisjointDst = 0x12
const m_PictOpDisjointIn = 0x15
const m_PictOpDisjointInReverse = 0x16
const m_PictOpDisjointMaximum = 0x1b
const m_PictOpDisjointMinimum = 0x10
const m_PictOpDisjointOut = 0x17
const m_PictOpDisjointOutReverse = 0x18
const m_PictOpDisjointOver = 0x13
const m_PictOpDisjointOverReverse = 0x14
const m_PictOpDisjointSrc = 0x11
const m_PictOpDisjointXor = 0x1b
const m_PictOpDst = 2
const m_PictOpExclusion = 0x3a
const m_PictOpHSLColor = 0x3d
const m_PictOpHSLHue = 0x3b
const m_PictOpHSLLuminosity = 0x3e
const m_PictOpHSLSaturation = 0x3c
const m_PictOpHardLight = 0x37
const m_PictOpIn = 5
const m_PictOpInReverse = 6
const m_PictOpLighten = 0x34
const m_PictOpMaximum = 13
const m_PictOpMinimum = 0
const m_PictOpMultiply = 0x30
const m_PictOpOut = 7
const m_PictOpOutReverse = 8
const m_PictOpOver = 3
const m_PictOpOverReverse = 4
const m_PictOpOverlay = 0x32
const m_PictOpSaturate = 13
const m_PictOpScreen = 0x31
const m_PictOpSoftLight = 0x38
const m_PictOpSrc = 1
const m_PictOpXor = 11
const m_PictStandardA1 = 4
const m_PictStandardA4 = 3
const m_PictStandardA8 = 2
const m_PictStandardARGB32 = 0
const m_PictStandardNUM = 5
const m_PictStandardRGB24 = 1
const m_PictTypeDirect = 1
const m_PictTypeIndexed = 0
const m_PlaceOnBottom = 1
const m_PlaceOnTop = 0
const m_PointerRoot = 1
const m_PointerWindow = 0
const m_PolyEdgeSharp = 0
const m_PolyEdgeSmooth = 1
const m_PolyModeImprecise = 1
const m_PolyModePrecise = 0
const m_PreferBlanking = 1
const m_PropModeAppend = 2
const m_PropModePrepend = 1
const m_PropModeReplace = 0
const m_PropertyDelete = 1
const m_PropertyNewValue = 0
const m_PropertyNotify = 28
const m_PseudoColor = 3
const m_QueuedAfterFlush = 2
const m_QueuedAfterReading = 1
const m_QueuedAlready = 0
const m_RAND_MAX = 0x7fffffff
const m_RENDER_MAJOR = 0
const m_RENDER_MINOR = 11
const m_RENDER_NAME = "RENDER"
const m_RaiseLowest = 0
const m_RectangleIn = 1
const m_RectangleOut = 0
const m_RectanglePart = 2
const m_ReparentNotify = 21
const m_RepeatNone = 0
const m_RepeatNormal = 1
const m_RepeatPad = 2
const m_RepeatReflect = 3
const m_ReplayKeyboard = 5
const m_ReplayPointer = 2
const m_ResizeRequest = 25
const m_RetainPermanent = 1
const m_RetainTemporary = 2
const m_RevertToParent = 2
const m_SCHED_FIFO = 1
const m_SCHED_OTHER = 2
const m_SCHED_RR = 3
const m_SIG_ATOMIC_MAX = "INT64_MAX"
const m_SIG_ATOMIC_MIN = "INT64_MIN"
const m_SIZE_MAX = "UINT64_MAX"
const m_STDC_HEADERS = 1
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
const m_Status = "int"
const m_StippleShape = 2
const m_SubPixelHorizontalBGR = 2
const m_SubPixelHorizontalRGB = 1
const m_SubPixelNone = 5
const m_SubPixelUnknown = 0
const m_SubPixelVerticalBGR = 4
const m_SubPixelVerticalRGB = 3
const m_Success = 0
const m_SyncBoth = 7
const m_SyncKeyboard = 4
const m_SyncPointer = 1
const m_TIMER_ABSTIME = 0x1
const m_TIMER_RELTIME = 0x0
const m_TIME_MONOTONIC = 2
const m_TIME_UTC = 1
const m_TileShape = 1
const m_TopIf = 2
const m_True = 1
const m_TrueColor = 4
const m_UINT16_MAX = 0xffff
const m_UINT32_MAX = 0xffffffff
const m_UINT64_MAX = 0xffffffffffffffff
const m_UINT8_MAX = 0xff
const m_UINTMAX_MAX = "UINT64_MAX"
const m_UINTPTR_MAX = "UINT64_MAX"
const m_UINT_FAST16_MAX = "UINT32_MAX"
const m_UINT_FAST32_MAX = "UINT32_MAX"
const m_UINT_FAST64_MAX = "UINT64_MAX"
const m_UINT_FAST8_MAX = "UINT32_MAX"
const m_UINT_LEAST16_MAX = "UINT16_MAX"
const m_UINT_LEAST32_MAX = "UINT32_MAX"
const m_UINT_LEAST64_MAX = "UINT64_MAX"
const m_UINT_LEAST8_MAX = "UINT8_MAX"
const m_UNLOCKED = 0
const m_UnmapGravity = 0
const m_UnmapNotify = 18
const m_Unsorted = 0
const m_VERSION = "0.9.10"
const m_VisibilityFullyObscured = 2
const m_VisibilityNotify = 15
const m_VisibilityPartiallyObscured = 1
const m_VisibilityUnobscured = 0
const m_VisualAllMask = 0x1FF
const m_VisualBitsPerRGBMask = 0x100
const m_VisualBlueMaskMask = 0x40
const m_VisualClassMask = 0x8
const m_VisualColormapSizeMask = 0x80
const m_VisualDepthMask = 0x4
const m_VisualGreenMaskMask = 0x20
const m_VisualIDMask = 0x1
const m_VisualNoMask = 0x0
const m_VisualRedMaskMask = 0x10
const m_VisualScreenMask = 0x2
const m_WCHAR_MAX = "__WCHAR_MAX"
const m_WCHAR_MIN = "__WCHAR_MIN"
const m_WINT_MAX = "INT32_MAX"
const m_WINT_MIN = "INT32_MIN"
const m_WLNSPERBATCH = 50
const m_WRCTSPERBATCH = 10
const m_WestGravity = 4
const m_WhenMapped = 1
const m_WidthValue = 0x0004
const m_WindingRule = 1
const m_WithdrawnState = 0
const m_XCNOENT = 2
const m_XCNOMEM = 1
const m_XCSUCCESS = 0
const m_XIMHotKeyStateOFF = 0x0002
const m_XIMHotKeyStateON = 0x0001
const m_XIMInitialState = 1
const m_XIMPreeditArea = 0x0001
const m_XIMPreeditCallbacks = 0x0002
const m_XIMPreeditEnable = 1
const m_XIMPreeditNone = 0x0010
const m_XIMPreeditNothing = 0x0008
const m_XIMPreeditPosition = 0x0004
const m_XIMPreeditUnKnown = 0
const m_XIMReverse = 1
const m_XIMStatusArea = 0x0100
const m_XIMStatusCallbacks = 0x0200
const m_XIMStatusNone = 0x0800
const m_XIMStatusNothing = 0x0400
const m_XIMStringConversionBottomEdge = 0x00000008
const m_XIMStringConversionBuffer = 0x0001
const m_XIMStringConversionChar = 0x0004
const m_XIMStringConversionConcealed = 0x00000010
const m_XIMStringConversionLeftEdge = 0x00000001
const m_XIMStringConversionLine = 0x0002
const m_XIMStringConversionRetrieval = 0x0002
const m_XIMStringConversionRightEdge = 0x00000002
const m_XIMStringConversionSubstitution = 0x0001
const m_XIMStringConversionTopEdge = 0x00000004
const m_XIMStringConversionWord = 0x0003
const m_XIMStringConversionWrapped = 0x00000020
const m_XK_0 = 0x0030
const m_XK_1 = 0x0031
const m_XK_2 = 0x0032
const m_XK_3 = 0x0033
const m_XK_4 = 0x0034
const m_XK_5 = 0x0035
const m_XK_6 = 0x0036
const m_XK_7 = 0x0037
const m_XK_8 = 0x0038
const m_XK_9 = 0x0039
const m_XK_A = 0x0041
const m_XK_AE = 0x00c6
const m_XK_Aacute = 0x00c1
const m_XK_Abelowdot = 0x1001ea0
const m_XK_Abreve = 0x01c3
const m_XK_Abreveacute = 0x1001eae
const m_XK_Abrevebelowdot = 0x1001eb6
const m_XK_Abrevegrave = 0x1001eb0
const m_XK_Abrevehook = 0x1001eb2
const m_XK_Abrevetilde = 0x1001eb4
const m_XK_AccessX_Enable = 0xfe70
const m_XK_AccessX_Feedback_Enable = 0xfe71
const m_XK_Acircumflex = 0x00c2
const m_XK_Acircumflexacute = 0x1001ea4
const m_XK_Acircumflexbelowdot = 0x1001eac
const m_XK_Acircumflexgrave = 0x1001ea6
const m_XK_Acircumflexhook = 0x1001ea8
const m_XK_Acircumflextilde = 0x1001eaa
const m_XK_Adiaeresis = 0x00c4
const m_XK_Agrave = 0x00c0
const m_XK_Ahook = 0x1001ea2
const m_XK_Alt_L = 0xffe9
const m_XK_Alt_R = 0xffea
const m_XK_Amacron = 0x03c0
const m_XK_Aogonek = 0x01a1
const m_XK_Arabic_0 = 0x1000660
const m_XK_Arabic_1 = 0x1000661
const m_XK_Arabic_2 = 0x1000662
const m_XK_Arabic_3 = 0x1000663
const m_XK_Arabic_4 = 0x1000664
const m_XK_Arabic_5 = 0x1000665
const m_XK_Arabic_6 = 0x1000666
const m_XK_Arabic_7 = 0x1000667
const m_XK_Arabic_8 = 0x1000668
const m_XK_Arabic_9 = 0x1000669
const m_XK_Arabic_ain = 0x05d9
const m_XK_Arabic_alef = 0x05c7
const m_XK_Arabic_alefmaksura = 0x05e9
const m_XK_Arabic_beh = 0x05c8
const m_XK_Arabic_comma = 0x05ac
const m_XK_Arabic_dad = 0x05d6
const m_XK_Arabic_dal = 0x05cf
const m_XK_Arabic_damma = 0x05ef
const m_XK_Arabic_dammatan = 0x05ec
const m_XK_Arabic_ddal = 0x1000688
const m_XK_Arabic_farsi_yeh = 0x10006cc
const m_XK_Arabic_fatha = 0x05ee
const m_XK_Arabic_fathatan = 0x05eb
const m_XK_Arabic_feh = 0x05e1
const m_XK_Arabic_fullstop = 0x10006d4
const m_XK_Arabic_gaf = 0x10006af
const m_XK_Arabic_ghain = 0x05da
const m_XK_Arabic_ha = 0x05e7
const m_XK_Arabic_hah = 0x05cd
const m_XK_Arabic_hamza = 0x05c1
const m_XK_Arabic_hamza_above = 0x1000654
const m_XK_Arabic_hamza_below = 0x1000655
const m_XK_Arabic_hamzaonalef = 0x05c3
const m_XK_Arabic_hamzaonwaw = 0x05c4
const m_XK_Arabic_hamzaonyeh = 0x05c6
const m_XK_Arabic_hamzaunderalef = 0x05c5
const m_XK_Arabic_heh = 0x05e7
const m_XK_Arabic_heh_doachashmee = 0x10006be
const m_XK_Arabic_heh_goal = 0x10006c1
const m_XK_Arabic_jeem = 0x05cc
const m_XK_Arabic_jeh = 0x1000698
const m_XK_Arabic_kaf = 0x05e3
const m_XK_Arabic_kasra = 0x05f0
const m_XK_Arabic_kasratan = 0x05ed
const m_XK_Arabic_keheh = 0x10006a9
const m_XK_Arabic_khah = 0x05ce
const m_XK_Arabic_lam = 0x05e4
const m_XK_Arabic_madda_above = 0x1000653
const m_XK_Arabic_maddaonalef = 0x05c2
const m_XK_Arabic_meem = 0x05e5
const m_XK_Arabic_noon = 0x05e6
const m_XK_Arabic_noon_ghunna = 0x10006ba
const m_XK_Arabic_peh = 0x100067e
const m_XK_Arabic_percent = 0x100066a
const m_XK_Arabic_qaf = 0x05e2
const m_XK_Arabic_question_mark = 0x05bf
const m_XK_Arabic_ra = 0x05d1
const m_XK_Arabic_rreh = 0x1000691
const m_XK_Arabic_sad = 0x05d5
const m_XK_Arabic_seen = 0x05d3
const m_XK_Arabic_semicolon = 0x05bb
const m_XK_Arabic_shadda = 0x05f1
const m_XK_Arabic_sheen = 0x05d4
const m_XK_Arabic_sukun = 0x05f2
const m_XK_Arabic_superscript_alef = 0x1000670
const m_XK_Arabic_switch = 0xff7e
const m_XK_Arabic_tah = 0x05d7
const m_XK_Arabic_tatweel = 0x05e0
const m_XK_Arabic_tcheh = 0x1000686
const m_XK_Arabic_teh = 0x05ca
const m_XK_Arabic_tehmarbuta = 0x05c9
const m_XK_Arabic_thal = 0x05d0
const m_XK_Arabic_theh = 0x05cb
const m_XK_Arabic_tteh = 0x1000679
const m_XK_Arabic_veh = 0x10006a4
const m_XK_Arabic_waw = 0x05e8
const m_XK_Arabic_yeh = 0x05ea
const m_XK_Arabic_yeh_baree = 0x10006d2
const m_XK_Arabic_zah = 0x05d8
const m_XK_Arabic_zain = 0x05d2
const m_XK_Aring = 0x00c5
const m_XK_Armenian_AT = 0x1000538
const m_XK_Armenian_AYB = 0x1000531
const m_XK_Armenian_BEN = 0x1000532
const m_XK_Armenian_CHA = 0x1000549
const m_XK_Armenian_DA = 0x1000534
const m_XK_Armenian_DZA = 0x1000541
const m_XK_Armenian_E = 0x1000537
const m_XK_Armenian_FE = 0x1000556
const m_XK_Armenian_GHAT = 0x1000542
const m_XK_Armenian_GIM = 0x1000533
const m_XK_Armenian_HI = 0x1000545
const m_XK_Armenian_HO = 0x1000540
const m_XK_Armenian_INI = 0x100053b
const m_XK_Armenian_JE = 0x100054b
const m_XK_Armenian_KE = 0x1000554
const m_XK_Armenian_KEN = 0x100053f
const m_XK_Armenian_KHE = 0x100053d
const m_XK_Armenian_LYUN = 0x100053c
const m_XK_Armenian_MEN = 0x1000544
const m_XK_Armenian_NU = 0x1000546
const m_XK_Armenian_O = 0x1000555
const m_XK_Armenian_PE = 0x100054a
const m_XK_Armenian_PYUR = 0x1000553
const m_XK_Armenian_RA = 0x100054c
const m_XK_Armenian_RE = 0x1000550
const m_XK_Armenian_SE = 0x100054d
const m_XK_Armenian_SHA = 0x1000547
const m_XK_Armenian_TCHE = 0x1000543
const m_XK_Armenian_TO = 0x1000539
const m_XK_Armenian_TSA = 0x100053e
const m_XK_Armenian_TSO = 0x1000551
const m_XK_Armenian_TYUN = 0x100054f
const m_XK_Armenian_VEV = 0x100054e
const m_XK_Armenian_VO = 0x1000548
const m_XK_Armenian_VYUN = 0x1000552
const m_XK_Armenian_YECH = 0x1000535
const m_XK_Armenian_ZA = 0x1000536
const m_XK_Armenian_ZHE = 0x100053a
const m_XK_Armenian_accent = 0x100055b
const m_XK_Armenian_amanak = 0x100055c
const m_XK_Armenian_apostrophe = 0x100055a
const m_XK_Armenian_at = 0x1000568
const m_XK_Armenian_ayb = 0x1000561
const m_XK_Armenian_ben = 0x1000562
const m_XK_Armenian_but = 0x100055d
const m_XK_Armenian_cha = 0x1000579
const m_XK_Armenian_da = 0x1000564
const m_XK_Armenian_dza = 0x1000571
const m_XK_Armenian_e = 0x1000567
const m_XK_Armenian_exclam = 0x100055c
const m_XK_Armenian_fe = 0x1000586
const m_XK_Armenian_full_stop = 0x1000589
const m_XK_Armenian_ghat = 0x1000572
const m_XK_Armenian_gim = 0x1000563
const m_XK_Armenian_hi = 0x1000575
const m_XK_Armenian_ho = 0x1000570
const m_XK_Armenian_hyphen = 0x100058a
const m_XK_Armenian_ini = 0x100056b
const m_XK_Armenian_je = 0x100057b
const m_XK_Armenian_ke = 0x1000584
const m_XK_Armenian_ken = 0x100056f
const m_XK_Armenian_khe = 0x100056d
const m_XK_Armenian_ligature_ew = 0x1000587
const m_XK_Armenian_lyun = 0x100056c
const m_XK_Armenian_men = 0x1000574
const m_XK_Armenian_nu = 0x1000576
const m_XK_Armenian_o = 0x1000585
const m_XK_Armenian_paruyk = 0x100055e
const m_XK_Armenian_pe = 0x100057a
const m_XK_Armenian_pyur = 0x1000583
const m_XK_Armenian_question = 0x100055e
const m_XK_Armenian_ra = 0x100057c
const m_XK_Armenian_re = 0x1000580
const m_XK_Armenian_se = 0x100057d
const m_XK_Armenian_separation_mark = 0x100055d
const m_XK_Armenian_sha = 0x1000577
const m_XK_Armenian_shesht = 0x100055b
const m_XK_Armenian_tche = 0x1000573
const m_XK_Armenian_to = 0x1000569
const m_XK_Armenian_tsa = 0x100056e
const m_XK_Armenian_tso = 0x1000581
const m_XK_Armenian_tyun = 0x100057f
const m_XK_Armenian_verjaket = 0x1000589
const m_XK_Armenian_vev = 0x100057e
const m_XK_Armenian_vo = 0x1000578
const m_XK_Armenian_vyun = 0x1000582
const m_XK_Armenian_yech = 0x1000565
const m_XK_Armenian_yentamna = 0x100058a
const m_XK_Armenian_za = 0x1000566
const m_XK_Armenian_zhe = 0x100056a
const m_XK_Atilde = 0x00c3
const m_XK_AudibleBell_Enable = 0xfe7a
const m_XK_B = 0x0042
const m_XK_Babovedot = 0x1001e02
const m_XK_BackSpace = 0xff08
const m_XK_Begin = 0xff58
const m_XK_BounceKeys_Enable = 0xfe74
const m_XK_Break = 0xff6b
const m_XK_Byelorussian_SHORTU = 0x06be
const m_XK_Byelorussian_shortu = 0x06ae
const m_XK_C = 0x0043
const m_XK_CH = 0xfea2
const m_XK_C_H = 0xfea5
const m_XK_C_h = 0xfea4
const m_XK_Cabovedot = 0x02c5
const m_XK_Cacute = 0x01c6
const m_XK_Cancel = 0xff69
const m_XK_Caps_Lock = 0xffe5
const m_XK_Ccaron = 0x01c8
const m_XK_Ccedilla = 0x00c7
const m_XK_Ccircumflex = 0x02c6
const m_XK_Ch = 0xfea1
const m_XK_Clear = 0xff0b
const m_XK_Codeinput = 0xff37
const m_XK_ColonSign = 0x10020a1
const m_XK_Control_L = 0xffe3
const m_XK_Control_R = 0xffe4
const m_XK_CruzeiroSign = 0x10020a2
const m_XK_Cyrillic_A = 0x06e1
const m_XK_Cyrillic_BE = 0x06e2
const m_XK_Cyrillic_CHE = 0x06fe
const m_XK_Cyrillic_CHE_descender = 0x10004b6
const m_XK_Cyrillic_CHE_vertstroke = 0x10004b8
const m_XK_Cyrillic_DE = 0x06e4
const m_XK_Cyrillic_DZHE = 0x06bf
const m_XK_Cyrillic_E = 0x06fc
const m_XK_Cyrillic_EF = 0x06e6
const m_XK_Cyrillic_EL = 0x06ec
const m_XK_Cyrillic_EM = 0x06ed
const m_XK_Cyrillic_EN = 0x06ee
const m_XK_Cyrillic_EN_descender = 0x10004a2
const m_XK_Cyrillic_ER = 0x06f2
const m_XK_Cyrillic_ES = 0x06f3
const m_XK_Cyrillic_GHE = 0x06e7
const m_XK_Cyrillic_GHE_bar = 0x1000492
const m_XK_Cyrillic_HA = 0x06e8
const m_XK_Cyrillic_HARDSIGN = 0x06ff
const m_XK_Cyrillic_HA_descender = 0x10004b2
const m_XK_Cyrillic_I = 0x06e9
const m_XK_Cyrillic_IE = 0x06e5
const m_XK_Cyrillic_IO = 0x06b3
const m_XK_Cyrillic_I_macron = 0x10004e2
const m_XK_Cyrillic_JE = 0x06b8
const m_XK_Cyrillic_KA = 0x06eb
const m_XK_Cyrillic_KA_descender = 0x100049a
const m_XK_Cyrillic_KA_vertstroke = 0x100049c
const m_XK_Cyrillic_LJE = 0x06b9
const m_XK_Cyrillic_NJE = 0x06ba
const m_XK_Cyrillic_O = 0x06ef
const m_XK_Cyrillic_O_bar = 0x10004e8
const m_XK_Cyrillic_PE = 0x06f0
const m_XK_Cyrillic_SCHWA = 0x10004d8
const m_XK_Cyrillic_SHA = 0x06fb
const m_XK_Cyrillic_SHCHA = 0x06fd
const m_XK_Cyrillic_SHHA = 0x10004ba
const m_XK_Cyrillic_SHORTI = 0x06ea
const m_XK_Cyrillic_SOFTSIGN = 0x06f8
const m_XK_Cyrillic_TE = 0x06f4
const m_XK_Cyrillic_TSE = 0x06e3
const m_XK_Cyrillic_U = 0x06f5
const m_XK_Cyrillic_U_macron = 0x10004ee
const m_XK_Cyrillic_U_straight = 0x10004ae
const m_XK_Cyrillic_U_straight_bar = 0x10004b0
const m_XK_Cyrillic_VE = 0x06f7
const m_XK_Cyrillic_YA = 0x06f1
const m_XK_Cyrillic_YERU = 0x06f9
const m_XK_Cyrillic_YU = 0x06e0
const m_XK_Cyrillic_ZE = 0x06fa
const m_XK_Cyrillic_ZHE = 0x06f6
const m_XK_Cyrillic_ZHE_descender = 0x1000496
const m_XK_Cyrillic_a = 0x06c1
const m_XK_Cyrillic_be = 0x06c2
const m_XK_Cyrillic_che = 0x06de
const m_XK_Cyrillic_che_descender = 0x10004b7
const m_XK_Cyrillic_che_vertstroke = 0x10004b9
const m_XK_Cyrillic_de = 0x06c4
const m_XK_Cyrillic_dzhe = 0x06af
const m_XK_Cyrillic_e = 0x06dc
const m_XK_Cyrillic_ef = 0x06c6
const m_XK_Cyrillic_el = 0x06cc
const m_XK_Cyrillic_em = 0x06cd
const m_XK_Cyrillic_en = 0x06ce
const m_XK_Cyrillic_en_descender = 0x10004a3
const m_XK_Cyrillic_er = 0x06d2
const m_XK_Cyrillic_es = 0x06d3
const m_XK_Cyrillic_ghe = 0x06c7
const m_XK_Cyrillic_ghe_bar = 0x1000493
const m_XK_Cyrillic_ha = 0x06c8
const m_XK_Cyrillic_ha_descender = 0x10004b3
const m_XK_Cyrillic_hardsign = 0x06df
const m_XK_Cyrillic_i = 0x06c9
const m_XK_Cyrillic_i_macron = 0x10004e3
const m_XK_Cyrillic_ie = 0x06c5
const m_XK_Cyrillic_io = 0x06a3
const m_XK_Cyrillic_je = 0x06a8
const m_XK_Cyrillic_ka = 0x06cb
const m_XK_Cyrillic_ka_descender = 0x100049b
const m_XK_Cyrillic_ka_vertstroke = 0x100049d
const m_XK_Cyrillic_lje = 0x06a9
const m_XK_Cyrillic_nje = 0x06aa
const m_XK_Cyrillic_o = 0x06cf
const m_XK_Cyrillic_o_bar = 0x10004e9
const m_XK_Cyrillic_pe = 0x06d0
const m_XK_Cyrillic_schwa = 0x10004d9
const m_XK_Cyrillic_sha = 0x06db
const m_XK_Cyrillic_shcha = 0x06dd
const m_XK_Cyrillic_shha = 0x10004bb
const m_XK_Cyrillic_shorti = 0x06ca
const m_XK_Cyrillic_softsign = 0x06d8
const m_XK_Cyrillic_te = 0x06d4
const m_XK_Cyrillic_tse = 0x06c3
const m_XK_Cyrillic_u = 0x06d5
const m_XK_Cyrillic_u_macron = 0x10004ef
const m_XK_Cyrillic_u_straight = 0x10004af
const m_XK_Cyrillic_u_straight_bar = 0x10004b1
const m_XK_Cyrillic_ve = 0x06d7
const m_XK_Cyrillic_ya = 0x06d1
const m_XK_Cyrillic_yeru = 0x06d9
const m_XK_Cyrillic_yu = 0x06c0
const m_XK_Cyrillic_ze = 0x06da
const m_XK_Cyrillic_zhe = 0x06d6
const m_XK_Cyrillic_zhe_descender = 0x1000497
const m_XK_D = 0x0044
const m_XK_Dabovedot = 0x1001e0a
const m_XK_Dcaron = 0x01cf
const m_XK_Delete = 0xffff
const m_XK_DongSign = 0x10020ab
const m_XK_Down = 0xff54
const m_XK_Dstroke = 0x01d0
const m_XK_E = 0x0045
const m_XK_ENG = 0x03bd
const m_XK_ETH = 0x00d0
const m_XK_EZH = 0x10001b7
const m_XK_Eabovedot = 0x03cc
const m_XK_Eacute = 0x00c9
const m_XK_Ebelowdot = 0x1001eb8
const m_XK_Ecaron = 0x01cc
const m_XK_Ecircumflex = 0x00ca
const m_XK_Ecircumflexacute = 0x1001ebe
const m_XK_Ecircumflexbelowdot = 0x1001ec6
const m_XK_Ecircumflexgrave = 0x1001ec0
const m_XK_Ecircumflexhook = 0x1001ec2
const m_XK_Ecircumflextilde = 0x1001ec4
const m_XK_EcuSign = 0x10020a0
const m_XK_Ediaeresis = 0x00cb
const m_XK_Egrave = 0x00c8
const m_XK_Ehook = 0x1001eba
const m_XK_Eisu_Shift = 0xff2f
const m_XK_Eisu_toggle = 0xff30
const m_XK_Emacron = 0x03aa
const m_XK_End = 0xff57
const m_XK_Eogonek = 0x01ca
const m_XK_Escape = 0xff1b
const m_XK_Eth = 0x00d0
const m_XK_Etilde = 0x1001ebc
const m_XK_EuroSign = 0x20ac
const m_XK_Execute = 0xff62
const m_XK_F = 0x0046
const m_XK_F1 = 0xffbe
const m_XK_F10 = 0xffc7
const m_XK_F11 = 0xffc8
const m_XK_F12 = 0xffc9
const m_XK_F13 = 0xffca
const m_XK_F14 = 0xffcb
const m_XK_F15 = 0xffcc
const m_XK_F16 = 0xffcd
const m_XK_F17 = 0xffce
const m_XK_F18 = 0xffcf
const m_XK_F19 = 0xffd0
const m_XK_F2 = 0xffbf
const m_XK_F20 = 0xffd1
const m_XK_F21 = 0xffd2
const m_XK_F22 = 0xffd3
const m_XK_F23 = 0xffd4
const m_XK_F24 = 0xffd5
const m_XK_F25 = 0xffd6
const m_XK_F26 = 0xffd7
const m_XK_F27 = 0xffd8
const m_XK_F28 = 0xffd9
const m_XK_F29 = 0xffda
const m_XK_F3 = 0xffc0
const m_XK_F30 = 0xffdb
const m_XK_F31 = 0xffdc
const m_XK_F32 = 0xffdd
const m_XK_F33 = 0xffde
const m_XK_F34 = 0xffdf
const m_XK_F35 = 0xffe0
const m_XK_F4 = 0xffc1
const m_XK_F5 = 0xffc2
const m_XK_F6 = 0xffc3
const m_XK_F7 = 0xffc4
const m_XK_F8 = 0xffc5
const m_XK_F9 = 0xffc6
const m_XK_FFrancSign = 0x10020a3
const m_XK_Fabovedot = 0x1001e1e
const m_XK_Farsi_0 = 0x10006f0
const m_XK_Farsi_1 = 0x10006f1
const m_XK_Farsi_2 = 0x10006f2
const m_XK_Farsi_3 = 0x10006f3
const m_XK_Farsi_4 = 0x10006f4
const m_XK_Farsi_5 = 0x10006f5
const m_XK_Farsi_6 = 0x10006f6
const m_XK_Farsi_7 = 0x10006f7
const m_XK_Farsi_8 = 0x10006f8
const m_XK_Farsi_9 = 0x10006f9
const m_XK_Farsi_yeh = 0x10006cc
const m_XK_Find = 0xff68
const m_XK_First_Virtual_Screen = 0xfed0
const m_XK_G = 0x0047
const m_XK_Gabovedot = 0x02d5
const m_XK_Gbreve = 0x02ab
const m_XK_Gcaron = 0x10001e6
const m_XK_Gcedilla = 0x03ab
const m_XK_Gcircumflex = 0x02d8
const m_XK_Georgian_an = 0x10010d0
const m_XK_Georgian_ban = 0x10010d1
const m_XK_Georgian_can = 0x10010ea
const m_XK_Georgian_char = 0x10010ed
const m_XK_Georgian_chin = 0x10010e9
const m_XK_Georgian_cil = 0x10010ec
const m_XK_Georgian_don = 0x10010d3
const m_XK_Georgian_en = 0x10010d4
const m_XK_Georgian_fi = 0x10010f6
const m_XK_Georgian_gan = 0x10010d2
const m_XK_Georgian_ghan = 0x10010e6
const m_XK_Georgian_hae = 0x10010f0
const m_XK_Georgian_har = 0x10010f4
const m_XK_Georgian_he = 0x10010f1
const m_XK_Georgian_hie = 0x10010f2
const m_XK_Georgian_hoe = 0x10010f5
const m_XK_Georgian_in = 0x10010d8
const m_XK_Georgian_jhan = 0x10010ef
const m_XK_Georgian_jil = 0x10010eb
const m_XK_Georgian_kan = 0x10010d9
const m_XK_Georgian_khar = 0x10010e5
const m_XK_Georgian_las = 0x10010da
const m_XK_Georgian_man = 0x10010db
const m_XK_Georgian_nar = 0x10010dc
const m_XK_Georgian_on = 0x10010dd
const m_XK_Georgian_par = 0x10010de
const m_XK_Georgian_phar = 0x10010e4
const m_XK_Georgian_qar = 0x10010e7
const m_XK_Georgian_rae = 0x10010e0
const m_XK_Georgian_san = 0x10010e1
const m_XK_Georgian_shin = 0x10010e8
const m_XK_Georgian_tan = 0x10010d7
const m_XK_Georgian_tar = 0x10010e2
const m_XK_Georgian_un = 0x10010e3
const m_XK_Georgian_vin = 0x10010d5
const m_XK_Georgian_we = 0x10010f3
const m_XK_Georgian_xan = 0x10010ee
const m_XK_Georgian_zen = 0x10010d6
const m_XK_Georgian_zhar = 0x10010df
const m_XK_Greek_ALPHA = 0x07c1
const m_XK_Greek_ALPHAaccent = 0x07a1
const m_XK_Greek_BETA = 0x07c2
const m_XK_Greek_CHI = 0x07d7
const m_XK_Greek_DELTA = 0x07c4
const m_XK_Greek_EPSILON = 0x07c5
const m_XK_Greek_EPSILONaccent = 0x07a2
const m_XK_Greek_ETA = 0x07c7
const m_XK_Greek_ETAaccent = 0x07a3
const m_XK_Greek_GAMMA = 0x07c3
const m_XK_Greek_IOTA = 0x07c9
const m_XK_Greek_IOTAaccent = 0x07a4
const m_XK_Greek_IOTAdiaeresis = 0x07a5
const m_XK_Greek_IOTAdieresis = 0x07a5
const m_XK_Greek_KAPPA = 0x07ca
const m_XK_Greek_LAMBDA = 0x07cb
const m_XK_Greek_LAMDA = 0x07cb
const m_XK_Greek_MU = 0x07cc
const m_XK_Greek_NU = 0x07cd
const m_XK_Greek_OMEGA = 0x07d9
const m_XK_Greek_OMEGAaccent = 0x07ab
const m_XK_Greek_OMICRON = 0x07cf
const m_XK_Greek_OMICRONaccent = 0x07a7
const m_XK_Greek_PHI = 0x07d6
const m_XK_Greek_PI = 0x07d0
const m_XK_Greek_PSI = 0x07d8
const m_XK_Greek_RHO = 0x07d1
const m_XK_Greek_SIGMA = 0x07d2
const m_XK_Greek_TAU = 0x07d4
const m_XK_Greek_THETA = 0x07c8
const m_XK_Greek_UPSILON = 0x07d5
const m_XK_Greek_UPSILONaccent = 0x07a8
const m_XK_Greek_UPSILONdieresis = 0x07a9
const m_XK_Greek_XI = 0x07ce
const m_XK_Greek_ZETA = 0x07c6
const m_XK_Greek_accentdieresis = 0x07ae
const m_XK_Greek_alpha = 0x07e1
const m_XK_Greek_alphaaccent = 0x07b1
const m_XK_Greek_beta = 0x07e2
const m_XK_Greek_chi = 0x07f7
const m_XK_Greek_delta = 0x07e4
const m_XK_Greek_epsilon = 0x07e5
const m_XK_Greek_epsilonaccent = 0x07b2
const m_XK_Greek_eta = 0x07e7
const m_XK_Greek_etaaccent = 0x07b3
const m_XK_Greek_finalsmallsigma = 0x07f3
const m_XK_Greek_gamma = 0x07e3
const m_XK_Greek_horizbar = 0x07af
const m_XK_Greek_iota = 0x07e9
const m_XK_Greek_iotaaccent = 0x07b4
const m_XK_Greek_iotaaccentdieresis = 0x07b6
const m_XK_Greek_iotadieresis = 0x07b5
const m_XK_Greek_kappa = 0x07ea
const m_XK_Greek_lambda = 0x07eb
const m_XK_Greek_lamda = 0x07eb
const m_XK_Greek_mu = 0x07ec
const m_XK_Greek_nu = 0x07ed
const m_XK_Greek_omega = 0x07f9
const m_XK_Greek_omegaaccent = 0x07bb
const m_XK_Greek_omicron = 0x07ef
const m_XK_Greek_omicronaccent = 0x07b7
const m_XK_Greek_phi = 0x07f6
const m_XK_Greek_pi = 0x07f0
const m_XK_Greek_psi = 0x07f8
const m_XK_Greek_rho = 0x07f1
const m_XK_Greek_sigma = 0x07f2
const m_XK_Greek_switch = 0xff7e
const m_XK_Greek_tau = 0x07f4
const m_XK_Greek_theta = 0x07e8
const m_XK_Greek_upsilon = 0x07f5
const m_XK_Greek_upsilonaccent = 0x07b8
const m_XK_Greek_upsilonaccentdieresis = 0x07ba
const m_XK_Greek_upsilondieresis = 0x07b9
const m_XK_Greek_xi = 0x07ee
const m_XK_Greek_zeta = 0x07e6
const m_XK_H = 0x0048
const m_XK_Hangul = 0xff31
const m_XK_Hangul_A = 0x0ebf
const m_XK_Hangul_AE = 0x0ec0
const m_XK_Hangul_AraeA = 0x0ef6
const m_XK_Hangul_AraeAE = 0x0ef7
const m_XK_Hangul_Banja = 0xff39
const m_XK_Hangul_Cieuc = 0x0eba
const m_XK_Hangul_Codeinput = 0xff37
const m_XK_Hangul_Dikeud = 0x0ea7
const m_XK_Hangul_E = 0x0ec4
const m_XK_Hangul_EO = 0x0ec3
const m_XK_Hangul_EU = 0x0ed1
const m_XK_Hangul_End = 0xff33
const m_XK_Hangul_Hanja = 0xff34
const m_XK_Hangul_Hieuh = 0x0ebe
const m_XK_Hangul_I = 0x0ed3
const m_XK_Hangul_Ieung = 0x0eb7
const m_XK_Hangul_J_Cieuc = 0x0eea
const m_XK_Hangul_J_Dikeud = 0x0eda
const m_XK_Hangul_J_Hieuh = 0x0eee
const m_XK_Hangul_J_Ieung = 0x0ee8
const m_XK_Hangul_J_Jieuj = 0x0ee9
const m_XK_Hangul_J_Khieuq = 0x0eeb
const m_XK_Hangul_J_Kiyeog = 0x0ed4
const m_XK_Hangul_J_KiyeogSios = 0x0ed6
const m_XK_Hangul_J_KkogjiDalrinIeung = 0x0ef9
const m_XK_Hangul_J_Mieum = 0x0ee3
const m_XK_Hangul_J_Nieun = 0x0ed7
const m_XK_Hangul_J_NieunHieuh = 0x0ed9
const m_XK_Hangul_J_NieunJieuj = 0x0ed8
const m_XK_Hangul_J_PanSios = 0x0ef8
const m_XK_Hangul_J_Phieuf = 0x0eed
const m_XK_Hangul_J_Pieub = 0x0ee4
const m_XK_Hangul_J_PieubSios = 0x0ee5
const m_XK_Hangul_J_Rieul = 0x0edb
const m_XK_Hangul_J_RieulHieuh = 0x0ee2
const m_XK_Hangul_J_RieulKiyeog = 0x0edc
const m_XK_Hangul_J_RieulMieum = 0x0edd
const m_XK_Hangul_J_RieulPhieuf = 0x0ee1
const m_XK_Hangul_J_RieulPieub = 0x0ede
const m_XK_Hangul_J_RieulSios = 0x0edf
const m_XK_Hangul_J_RieulTieut = 0x0ee0
const m_XK_Hangul_J_Sios = 0x0ee6
const m_XK_Hangul_J_SsangKiyeog = 0x0ed5
const m_XK_Hangul_J_SsangSios = 0x0ee7
const m_XK_Hangul_J_Tieut = 0x0eec
const m_XK_Hangul_J_YeorinHieuh = 0x0efa
const m_XK_Hangul_Jamo = 0xff35
const m_XK_Hangul_Jeonja = 0xff38
const m_XK_Hangul_Jieuj = 0x0eb8
const m_XK_Hangul_Khieuq = 0x0ebb
const m_XK_Hangul_Kiyeog = 0x0ea1
const m_XK_Hangul_KiyeogSios = 0x0ea3
const m_XK_Hangul_KkogjiDalrinIeung = 0x0ef3
const m_XK_Hangul_Mieum = 0x0eb1
const m_XK_Hangul_MultipleCandidate = 0xff3d
const m_XK_Hangul_Nieun = 0x0ea4
const m_XK_Hangul_NieunHieuh = 0x0ea6
const m_XK_Hangul_NieunJieuj = 0x0ea5
const m_XK_Hangul_O = 0x0ec7
const m_XK_Hangul_OE = 0x0eca
const m_XK_Hangul_PanSios = 0x0ef2
const m_XK_Hangul_Phieuf = 0x0ebd
const m_XK_Hangul_Pieub = 0x0eb2
const m_XK_Hangul_PieubSios = 0x0eb4
const m_XK_Hangul_PostHanja = 0xff3b
const m_XK_Hangul_PreHanja = 0xff3a
const m_XK_Hangul_PreviousCandidate = 0xff3e
const m_XK_Hangul_Rieul = 0x0ea9
const m_XK_Hangul_RieulHieuh = 0x0eb0
const m_XK_Hangul_RieulKiyeog = 0x0eaa
const m_XK_Hangul_RieulMieum = 0x0eab
const m_XK_Hangul_RieulPhieuf = 0x0eaf
const m_XK_Hangul_RieulPieub = 0x0eac
const m_XK_Hangul_RieulSios = 0x0ead
const m_XK_Hangul_RieulTieut = 0x0eae
const m_XK_Hangul_RieulYeorinHieuh = 0x0eef
const m_XK_Hangul_Romaja = 0xff36
const m_XK_Hangul_SingleCandidate = 0xff3c
const m_XK_Hangul_Sios = 0x0eb5
const m_XK_Hangul_Special = 0xff3f
const m_XK_Hangul_SsangDikeud = 0x0ea8
const m_XK_Hangul_SsangJieuj = 0x0eb9
const m_XK_Hangul_SsangKiyeog = 0x0ea2
const m_XK_Hangul_SsangPieub = 0x0eb3
const m_XK_Hangul_SsangSios = 0x0eb6
const m_XK_Hangul_Start = 0xff32
const m_XK_Hangul_SunkyeongeumMieum = 0x0ef0
const m_XK_Hangul_SunkyeongeumPhieuf = 0x0ef4
const m_XK_Hangul_SunkyeongeumPieub = 0x0ef1
const m_XK_Hangul_Tieut = 0x0ebc
const m_XK_Hangul_U = 0x0ecc
const m_XK_Hangul_WA = 0x0ec8
const m_XK_Hangul_WAE = 0x0ec9
const m_XK_Hangul_WE = 0x0ece
const m_XK_Hangul_WEO = 0x0ecd
const m_XK_Hangul_WI = 0x0ecf
const m_XK_Hangul_YA = 0x0ec1
const m_XK_Hangul_YAE = 0x0ec2
const m_XK_Hangul_YE = 0x0ec6
const m_XK_Hangul_YEO = 0x0ec5
const m_XK_Hangul_YI = 0x0ed2
const m_XK_Hangul_YO = 0x0ecb
const m_XK_Hangul_YU = 0x0ed0
const m_XK_Hangul_YeorinHieuh = 0x0ef5
const m_XK_Hangul_switch = 0xff7e
const m_XK_Hankaku = 0xff29
const m_XK_Hcircumflex = 0x02a6
const m_XK_Hebrew_switch = 0xff7e
const m_XK_Help = 0xff6a
const m_XK_Henkan = 0xff23
const m_XK_Henkan_Mode = 0xff23
const m_XK_Hiragana = 0xff25
const m_XK_Hiragana_Katakana = 0xff27
const m_XK_Home = 0xff50
const m_XK_Hstroke = 0x02a1
const m_XK_Hyper_L = 0xffed
const m_XK_Hyper_R = 0xffee
const m_XK_I = 0x0049
const m_XK_ISO_Center_Object = 0xfe33
const m_XK_ISO_Continuous_Underline = 0xfe30
const m_XK_ISO_Discontinuous_Underline = 0xfe31
const m_XK_ISO_Emphasize = 0xfe32
const m_XK_ISO_Enter = 0xfe34
const m_XK_ISO_Fast_Cursor_Down = 0xfe2f
const m_XK_ISO_Fast_Cursor_Left = 0xfe2c
const m_XK_ISO_Fast_Cursor_Right = 0xfe2d
const m_XK_ISO_Fast_Cursor_Up = 0xfe2e
const m_XK_ISO_First_Group = 0xfe0c
const m_XK_ISO_First_Group_Lock = 0xfe0d
const m_XK_ISO_Group_Latch = 0xfe06
const m_XK_ISO_Group_Lock = 0xfe07
const m_XK_ISO_Group_Shift = 0xff7e
const m_XK_ISO_Last_Group = 0xfe0e
const m_XK_ISO_Last_Group_Lock = 0xfe0f
const m_XK_ISO_Left_Tab = 0xfe20
const m_XK_ISO_Level2_Latch = 0xfe02
const m_XK_ISO_Level3_Latch = 0xfe04
const m_XK_ISO_Level3_Lock = 0xfe05
const m_XK_ISO_Level3_Shift = 0xfe03
const m_XK_ISO_Level5_Latch = 0xfe12
const m_XK_ISO_Level5_Lock = 0xfe13
const m_XK_ISO_Level5_Shift = 0xfe11
const m_XK_ISO_Lock = 0xfe01
const m_XK_ISO_Move_Line_Down = 0xfe22
const m_XK_ISO_Move_Line_Up = 0xfe21
const m_XK_ISO_Next_Group = 0xfe08
const m_XK_ISO_Next_Group_Lock = 0xfe09
const m_XK_ISO_Partial_Line_Down = 0xfe24
const m_XK_ISO_Partial_Line_Up = 0xfe23
const m_XK_ISO_Partial_Space_Left = 0xfe25
const m_XK_ISO_Partial_Space_Right = 0xfe26
const m_XK_ISO_Prev_Group = 0xfe0a
const m_XK_ISO_Prev_Group_Lock = 0xfe0b
const m_XK_ISO_Release_Both_Margins = 0xfe2b
const m_XK_ISO_Release_Margin_Left = 0xfe29
const m_XK_ISO_Release_Margin_Right = 0xfe2a
const m_XK_ISO_Set_Margin_Left = 0xfe27
const m_XK_ISO_Set_Margin_Right = 0xfe28
const m_XK_Iabovedot = 0x02a9
const m_XK_Iacute = 0x00cd
const m_XK_Ibelowdot = 0x1001eca
const m_XK_Ibreve = 0x100012c
const m_XK_Icircumflex = 0x00ce
const m_XK_Idiaeresis = 0x00cf
const m_XK_Igrave = 0x00cc
const m_XK_Ihook = 0x1001ec8
const m_XK_Imacron = 0x03cf
const m_XK_Insert = 0xff63
const m_XK_Iogonek = 0x03c7
const m_XK_Itilde = 0x03a5
const m_XK_J = 0x004a
const m_XK_Jcircumflex = 0x02ac
const m_XK_K = 0x004b
const m_XK_KP_0 = 0xffb0
const m_XK_KP_1 = 0xffb1
const m_XK_KP_2 = 0xffb2
const m_XK_KP_3 = 0xffb3
const m_XK_KP_4 = 0xffb4
const m_XK_KP_5 = 0xffb5
const m_XK_KP_6 = 0xffb6
const m_XK_KP_7 = 0xffb7
const m_XK_KP_8 = 0xffb8
const m_XK_KP_9 = 0xffb9
const m_XK_KP_Add = 0xffab
const m_XK_KP_Begin = 0xff9d
const m_XK_KP_Decimal = 0xffae
const m_XK_KP_Delete = 0xff9f
const m_XK_KP_Divide = 0xffaf
const m_XK_KP_Down = 0xff99
const m_XK_KP_End = 0xff9c
const m_XK_KP_Enter = 0xff8d
const m_XK_KP_Equal = 0xffbd
const m_XK_KP_F1 = 0xff91
const m_XK_KP_F2 = 0xff92
const m_XK_KP_F3 = 0xff93
const m_XK_KP_F4 = 0xff94
const m_XK_KP_Home = 0xff95
const m_XK_KP_Insert = 0xff9e
const m_XK_KP_Left = 0xff96
const m_XK_KP_Multiply = 0xffaa
const m_XK_KP_Next = 0xff9b
const m_XK_KP_Page_Down = 0xff9b
const m_XK_KP_Page_Up = 0xff9a
const m_XK_KP_Prior = 0xff9a
const m_XK_KP_Right = 0xff98
const m_XK_KP_Separator = 0xffac
const m_XK_KP_Space = 0xff80
const m_XK_KP_Subtract = 0xffad
const m_XK_KP_Tab = 0xff89
const m_XK_KP_Up = 0xff97
const m_XK_Kana_Lock = 0xff2d
const m_XK_Kana_Shift = 0xff2e
const m_XK_Kanji = 0xff21
const m_XK_Kanji_Bangou = 0xff37
const m_XK_Katakana = 0xff26
const m_XK_Kcedilla = 0x03d3
const m_XK_Korean_Won = 0x0eff
const m_XK_L = 0x004c
const m_XK_L1 = 0xffc8
const m_XK_L10 = 0xffd1
const m_XK_L2 = 0xffc9
const m_XK_L3 = 0xffca
const m_XK_L4 = 0xffcb
const m_XK_L5 = 0xffcc
const m_XK_L6 = 0xffcd
const m_XK_L7 = 0xffce
const m_XK_L8 = 0xffcf
const m_XK_L9 = 0xffd0
const m_XK_Lacute = 0x01c5
const m_XK_Last_Virtual_Screen = 0xfed4
const m_XK_Lbelowdot = 0x1001e36
const m_XK_Lcaron = 0x01a5
const m_XK_Lcedilla = 0x03a6
const m_XK_Left = 0xff51
const m_XK_Linefeed = 0xff0a
const m_XK_LiraSign = 0x10020a4
const m_XK_Lstroke = 0x01a3
const m_XK_M = 0x004d
const m_XK_Mabovedot = 0x1001e40
const m_XK_Macedonia_DSE = 0x06b5
const m_XK_Macedonia_GJE = 0x06b2
const m_XK_Macedonia_KJE = 0x06bc
const m_XK_Macedonia_dse = 0x06a5
const m_XK_Macedonia_gje = 0x06a2
const m_XK_Macedonia_kje = 0x06ac
const m_XK_Mae_Koho = 0xff3e
const m_XK_Massyo = 0xff2c
const m_XK_Menu = 0xff67
const m_XK_Meta_L = 0xffe7
const m_XK_Meta_R = 0xffe8
const m_XK_MillSign = 0x10020a5
const m_XK_Mode_switch = 0xff7e
const m_XK_MouseKeys_Accel_Enable = 0xfe77
const m_XK_MouseKeys_Enable = 0xfe76
const m_XK_Muhenkan = 0xff22
const m_XK_Multi_key = 0xff20
const m_XK_MultipleCandidate = 0xff3d
const m_XK_N = 0x004e
const m_XK_Nacute = 0x01d1
const m_XK_NairaSign = 0x10020a6
const m_XK_Ncaron = 0x01d2
const m_XK_Ncedilla = 0x03d1
const m_XK_NewSheqelSign = 0x10020aa
const m_XK_Next = 0xff56
const m_XK_Next_Virtual_Screen = 0xfed2
const m_XK_Ntilde = 0x00d1
const m_XK_Num_Lock = 0xff7f
const m_XK_O = 0x004f
const m_XK_OE = 0x13bc
const m_XK_Oacute = 0x00d3
const m_XK_Obarred = 0x100019f
const m_XK_Obelowdot = 0x1001ecc
const m_XK_Ocaron = 0x10001d1
const m_XK_Ocircumflex = 0x00d4
const m_XK_Ocircumflexacute = 0x1001ed0
const m_XK_Ocircumflexbelowdot = 0x1001ed8
const m_XK_Ocircumflexgrave = 0x1001ed2
const m_XK_Ocircumflexhook = 0x1001ed4
const m_XK_Ocircumflextilde = 0x1001ed6
const m_XK_Odiaeresis = 0x00d6
const m_XK_Odoubleacute = 0x01d5
const m_XK_Ograve = 0x00d2
const m_XK_Ohook = 0x1001ece
const m_XK_Ohorn = 0x10001a0
const m_XK_Ohornacute = 0x1001eda
const m_XK_Ohornbelowdot = 0x1001ee2
const m_XK_Ohorngrave = 0x1001edc
const m_XK_Ohornhook = 0x1001ede
const m_XK_Ohorntilde = 0x1001ee0
const m_XK_Omacron = 0x03d2
const m_XK_Ooblique = 0x00d8
const m_XK_Oslash = 0x00d8
const m_XK_Otilde = 0x00d5
const m_XK_Overlay1_Enable = 0xfe78
const m_XK_Overlay2_Enable = 0xfe79
const m_XK_P = 0x0050
const m_XK_Pabovedot = 0x1001e56
const m_XK_Page_Down = 0xff56
const m_XK_Page_Up = 0xff55
const m_XK_Pause = 0xff13
const m_XK_PesetaSign = 0x10020a7
const m_XK_Pointer_Accelerate = 0xfefa
const m_XK_Pointer_Button1 = 0xfee9
const m_XK_Pointer_Button2 = 0xfeea
const m_XK_Pointer_Button3 = 0xfeeb
const m_XK_Pointer_Button4 = 0xfeec
const m_XK_Pointer_Button5 = 0xfeed
const m_XK_Pointer_Button_Dflt = 0xfee8
const m_XK_Pointer_DblClick1 = 0xfeef
const m_XK_Pointer_DblClick2 = 0xfef0
const m_XK_Pointer_DblClick3 = 0xfef1
const m_XK_Pointer_DblClick4 = 0xfef2
const m_XK_Pointer_DblClick5 = 0xfef3
const m_XK_Pointer_DblClick_Dflt = 0xfeee
const m_XK_Pointer_DfltBtnNext = 0xfefb
const m_XK_Pointer_DfltBtnPrev = 0xfefc
const m_XK_Pointer_Down = 0xfee3
const m_XK_Pointer_DownLeft = 0xfee6
const m_XK_Pointer_DownRight = 0xfee7
const m_XK_Pointer_Drag1 = 0xfef5
const m_XK_Pointer_Drag2 = 0xfef6
const m_XK_Pointer_Drag3 = 0xfef7
const m_XK_Pointer_Drag4 = 0xfef8
const m_XK_Pointer_Drag5 = 0xfefd
const m_XK_Pointer_Drag_Dflt = 0xfef4
const m_XK_Pointer_EnableKeys = 0xfef9
const m_XK_Pointer_Left = 0xfee0
const m_XK_Pointer_Right = 0xfee1
const m_XK_Pointer_Up = 0xfee2
const m_XK_Pointer_UpLeft = 0xfee4
const m_XK_Pointer_UpRight = 0xfee5
const m_XK_Prev_Virtual_Screen = 0xfed1
const m_XK_PreviousCandidate = 0xff3e
const m_XK_Print = 0xff61
const m_XK_Prior = 0xff55
const m_XK_Q = 0x0051
const m_XK_R = 0x0052
const m_XK_R1 = 0xffd2
const m_XK_R10 = 0xffdb
const m_XK_R11 = 0xffdc
const m_XK_R12 = 0xffdd
const m_XK_R13 = 0xffde
const m_XK_R14 = 0xffdf
const m_XK_R15 = 0xffe0
const m_XK_R2 = 0xffd3
const m_XK_R3 = 0xffd4
const m_XK_R4 = 0xffd5
const m_XK_R5 = 0xffd6
const m_XK_R6 = 0xffd7
const m_XK_R7 = 0xffd8
const m_XK_R8 = 0xffd9
const m_XK_R9 = 0xffda
const m_XK_Racute = 0x01c0
const m_XK_Rcaron = 0x01d8
const m_XK_Rcedilla = 0x03a3
const m_XK_Redo = 0xff66
const m_XK_RepeatKeys_Enable = 0xfe72
const m_XK_Return = 0xff0d
const m_XK_Right = 0xff53
const m_XK_Romaji = 0xff24
const m_XK_RupeeSign = 0x10020a8
const m_XK_S = 0x0053
const m_XK_SCHWA = 0x100018f
const m_XK_Sabovedot = 0x1001e60
const m_XK_Sacute = 0x01a6
const m_XK_Scaron = 0x01a9
const m_XK_Scedilla = 0x01aa
const m_XK_Scircumflex = 0x02de
const m_XK_Scroll_Lock = 0xff14
const m_XK_Select = 0xff60
const m_XK_Serbian_DJE = 0x06b1
const m_XK_Serbian_DZE = 0x06bf
const m_XK_Serbian_JE = 0x06b8
const m_XK_Serbian_LJE = 0x06b9
const m_XK_Serbian_NJE = 0x06ba
const m_XK_Serbian_TSHE = 0x06bb
const m_XK_Serbian_dje = 0x06a1
const m_XK_Serbian_dze = 0x06af
const m_XK_Serbian_je = 0x06a8
const m_XK_Serbian_lje = 0x06a9
const m_XK_Serbian_nje = 0x06aa
const m_XK_Serbian_tshe = 0x06ab
const m_XK_Shift_L = 0xffe1
const m_XK_Shift_Lock = 0xffe6
const m_XK_Shift_R = 0xffe2
const m_XK_SingleCandidate = 0xff3c
const m_XK_Sinh_a = 0x1000d85
const m_XK_Sinh_aa = 0x1000d86
const m_XK_Sinh_aa2 = 0x1000dcf
const m_XK_Sinh_ae = 0x1000d87
const m_XK_Sinh_ae2 = 0x1000dd0
const m_XK_Sinh_aee = 0x1000d88
const m_XK_Sinh_aee2 = 0x1000dd1
const m_XK_Sinh_ai = 0x1000d93
const m_XK_Sinh_ai2 = 0x1000ddb
const m_XK_Sinh_al = 0x1000dca
const m_XK_Sinh_au = 0x1000d96
const m_XK_Sinh_au2 = 0x1000dde
const m_XK_Sinh_ba = 0x1000db6
const m_XK_Sinh_bha = 0x1000db7
const m_XK_Sinh_ca = 0x1000da0
const m_XK_Sinh_cha = 0x1000da1
const m_XK_Sinh_dda = 0x1000da9
const m_XK_Sinh_ddha = 0x1000daa
const m_XK_Sinh_dha = 0x1000daf
const m_XK_Sinh_dhha = 0x1000db0
const m_XK_Sinh_e = 0x1000d91
const m_XK_Sinh_e2 = 0x1000dd9
const m_XK_Sinh_ee = 0x1000d92
const m_XK_Sinh_ee2 = 0x1000dda
const m_XK_Sinh_fa = 0x1000dc6
const m_XK_Sinh_ga = 0x1000d9c
const m_XK_Sinh_gha = 0x1000d9d
const m_XK_Sinh_h2 = 0x1000d83
const m_XK_Sinh_ha = 0x1000dc4
const m_XK_Sinh_i = 0x1000d89
const m_XK_Sinh_i2 = 0x1000dd2
const m_XK_Sinh_ii = 0x1000d8a
const m_XK_Sinh_ii2 = 0x1000dd3
const m_XK_Sinh_ja = 0x1000da2
const m_XK_Sinh_jha = 0x1000da3
const m_XK_Sinh_jnya = 0x1000da5
const m_XK_Sinh_ka = 0x1000d9a
const m_XK_Sinh_kha = 0x1000d9b
const m_XK_Sinh_kunddaliya = 0x1000df4
const m_XK_Sinh_la = 0x1000dbd
const m_XK_Sinh_lla = 0x1000dc5
const m_XK_Sinh_lu = 0x1000d8f
const m_XK_Sinh_lu2 = 0x1000ddf
const m_XK_Sinh_luu = 0x1000d90
const m_XK_Sinh_luu2 = 0x1000df3
const m_XK_Sinh_ma = 0x1000db8
const m_XK_Sinh_mba = 0x1000db9
const m_XK_Sinh_na = 0x1000db1
const m_XK_Sinh_ndda = 0x1000dac
const m_XK_Sinh_ndha = 0x1000db3
const m_XK_Sinh_ng = 0x1000d82
const m_XK_Sinh_ng2 = 0x1000d9e
const m_XK_Sinh_nga = 0x1000d9f
const m_XK_Sinh_nja = 0x1000da6
const m_XK_Sinh_nna = 0x1000dab
const m_XK_Sinh_nya = 0x1000da4
const m_XK_Sinh_o = 0x1000d94
const m_XK_Sinh_o2 = 0x1000ddc
const m_XK_Sinh_oo = 0x1000d95
const m_XK_Sinh_oo2 = 0x1000ddd
const m_XK_Sinh_pa = 0x1000db4
const m_XK_Sinh_pha = 0x1000db5
const m_XK_Sinh_ra = 0x1000dbb
const m_XK_Sinh_ri = 0x1000d8d
const m_XK_Sinh_rii = 0x1000d8e
const m_XK_Sinh_ru2 = 0x1000dd8
const m_XK_Sinh_ruu2 = 0x1000df2
const m_XK_Sinh_sa = 0x1000dc3
const m_XK_Sinh_sha = 0x1000dc1
const m_XK_Sinh_ssha = 0x1000dc2
const m_XK_Sinh_tha = 0x1000dad
const m_XK_Sinh_thha = 0x1000dae
const m_XK_Sinh_tta = 0x1000da7
const m_XK_Sinh_ttha = 0x1000da8
const m_XK_Sinh_u = 0x1000d8b
const m_XK_Sinh_u2 = 0x1000dd4
const m_XK_Sinh_uu = 0x1000d8c
const m_XK_Sinh_uu2 = 0x1000dd6
const m_XK_Sinh_va = 0x1000dc0
const m_XK_Sinh_ya = 0x1000dba
const m_XK_SlowKeys_Enable = 0xfe73
const m_XK_StickyKeys_Enable = 0xfe75
const m_XK_Super_L = 0xffeb
const m_XK_Super_R = 0xffec
const m_XK_Sys_Req = 0xff15
const m_XK_T = 0x0054
const m_XK_THORN = 0x00de
const m_XK_Tab = 0xff09
const m_XK_Tabovedot = 0x1001e6a
const m_XK_Tcaron = 0x01ab
const m_XK_Tcedilla = 0x01de
const m_XK_Terminate_Server = 0xfed5
const m_XK_Thai_baht = 0x0ddf
const m_XK_Thai_bobaimai = 0x0dba
const m_XK_Thai_chochan = 0x0da8
const m_XK_Thai_chochang = 0x0daa
const m_XK_Thai_choching = 0x0da9
const m_XK_Thai_chochoe = 0x0dac
const m_XK_Thai_dochada = 0x0dae
const m_XK_Thai_dodek = 0x0db4
const m_XK_Thai_fofa = 0x0dbd
const m_XK_Thai_fofan = 0x0dbf
const m_XK_Thai_hohip = 0x0dcb
const m_XK_Thai_honokhuk = 0x0dce
const m_XK_Thai_khokhai = 0x0da2
const m_XK_Thai_khokhon = 0x0da5
const m_XK_Thai_khokhuat = 0x0da3
const m_XK_Thai_khokhwai = 0x0da4
const m_XK_Thai_khorakhang = 0x0da6
const m_XK_Thai_kokai = 0x0da1
const m_XK_Thai_lakkhangyao = 0x0de5
const m_XK_Thai_lekchet = 0x0df7
const m_XK_Thai_lekha = 0x0df5
const m_XK_Thai_lekhok = 0x0df6
const m_XK_Thai_lekkao = 0x0df9
const m_XK_Thai_leknung = 0x0df1
const m_XK_Thai_lekpaet = 0x0df8
const m_XK_Thai_leksam = 0x0df3
const m_XK_Thai_leksi = 0x0df4
const m_XK_Thai_leksong = 0x0df2
const m_XK_Thai_leksun = 0x0df0
const m_XK_Thai_lochula = 0x0dcc
const m_XK_Thai_loling = 0x0dc5
const m_XK_Thai_lu = 0x0dc6
const m_XK_Thai_maichattawa = 0x0deb
const m_XK_Thai_maiek = 0x0de8
const m_XK_Thai_maihanakat = 0x0dd1
const m_XK_Thai_maihanakat_maitho = 0x0dde
const m_XK_Thai_maitaikhu = 0x0de7
const m_XK_Thai_maitho = 0x0de9
const m_XK_Thai_maitri = 0x0dea
const m_XK_Thai_maiyamok = 0x0de6
const m_XK_Thai_moma = 0x0dc1
const m_XK_Thai_ngongu = 0x0da7
const m_XK_Thai_nikhahit = 0x0ded
const m_XK_Thai_nonen = 0x0db3
const m_XK_Thai_nonu = 0x0db9
const m_XK_Thai_oang = 0x0dcd
const m_XK_Thai_paiyannoi = 0x0dcf
const m_XK_Thai_phinthu = 0x0dda
const m_XK_Thai_phophan = 0x0dbe
const m_XK_Thai_phophung = 0x0dbc
const m_XK_Thai_phosamphao = 0x0dc0
const m_XK_Thai_popla = 0x0dbb
const m_XK_Thai_rorua = 0x0dc3
const m_XK_Thai_ru = 0x0dc4
const m_XK_Thai_saraa = 0x0dd0
const m_XK_Thai_saraaa = 0x0dd2
const m_XK_Thai_saraae = 0x0de1
const m_XK_Thai_saraaimaimalai = 0x0de4
const m_XK_Thai_saraaimaimuan = 0x0de3
const m_XK_Thai_saraam = 0x0dd3
const m_XK_Thai_sarae = 0x0de0
const m_XK_Thai_sarai = 0x0dd4
const m_XK_Thai_saraii = 0x0dd5
const m_XK_Thai_sarao = 0x0de2
const m_XK_Thai_sarau = 0x0dd8
const m_XK_Thai_saraue = 0x0dd6
const m_XK_Thai_sarauee = 0x0dd7
const m_XK_Thai_sarauu = 0x0dd9
const m_XK_Thai_sorusi = 0x0dc9
const m_XK_Thai_sosala = 0x0dc8
const m_XK_Thai_soso = 0x0dab
const m_XK_Thai_sosua = 0x0dca
const m_XK_Thai_thanthakhat = 0x0dec
const m_XK_Thai_thonangmontho = 0x0db1
const m_XK_Thai_thophuthao = 0x0db2
const m_XK_Thai_thothahan = 0x0db7
const m_XK_Thai_thothan = 0x0db0
const m_XK_Thai_thothong = 0x0db8
const m_XK_Thai_thothung = 0x0db6
const m_XK_Thai_topatak = 0x0daf
const m_XK_Thai_totao = 0x0db5
const m_XK_Thai_wowaen = 0x0dc7
const m_XK_Thai_yoyak = 0x0dc2
const m_XK_Thai_yoying = 0x0dad
const m_XK_Thorn = 0x00de
const m_XK_Touroku = 0xff2b
const m_XK_Tslash = 0x03ac
const m_XK_U = 0x0055
const m_XK_Uacute = 0x00da
const m_XK_Ubelowdot = 0x1001ee4
const m_XK_Ubreve = 0x02dd
const m_XK_Ucircumflex = 0x00db
const m_XK_Udiaeresis = 0x00dc
const m_XK_Udoubleacute = 0x01db
const m_XK_Ugrave = 0x00d9
const m_XK_Uhook = 0x1001ee6
const m_XK_Uhorn = 0x10001af
const m_XK_Uhornacute = 0x1001ee8
const m_XK_Uhornbelowdot = 0x1001ef0
const m_XK_Uhorngrave = 0x1001eea
const m_XK_Uhornhook = 0x1001eec
const m_XK_Uhorntilde = 0x1001eee
const m_XK_Ukrainian_GHE_WITH_UPTURN = 0x06bd
const m_XK_Ukrainian_I = 0x06b6
const m_XK_Ukrainian_IE = 0x06b4
const m_XK_Ukrainian_YI = 0x06b7
const m_XK_Ukrainian_ghe_with_upturn = 0x06ad
const m_XK_Ukrainian_i = 0x06a6
const m_XK_Ukrainian_ie = 0x06a4
const m_XK_Ukrainian_yi = 0x06a7
const m_XK_Ukranian_I = 0x06b6
const m_XK_Ukranian_JE = 0x06b4
const m_XK_Ukranian_YI = 0x06b7
const m_XK_Ukranian_i = 0x06a6
const m_XK_Ukranian_je = 0x06a4
const m_XK_Ukranian_yi = 0x06a7
const m_XK_Umacron = 0x03de
const m_XK_Undo = 0xff65
const m_XK_Uogonek = 0x03d9
const m_XK_Up = 0xff52
const m_XK_Uring = 0x01d9
const m_XK_Utilde = 0x03dd
const m_XK_V = 0x0056
const m_XK_VoidSymbol = 0xffffff
const m_XK_W = 0x0057
const m_XK_Wacute = 0x1001e82
const m_XK_Wcircumflex = 0x1000174
const m_XK_Wdiaeresis = 0x1001e84
const m_XK_Wgrave = 0x1001e80
const m_XK_WonSign = 0x10020a9
const m_XK_X = 0x0058
const m_XK_Xabovedot = 0x1001e8a
const m_XK_Y = 0x0059
const m_XK_Yacute = 0x00dd
const m_XK_Ybelowdot = 0x1001ef4
const m_XK_Ycircumflex = 0x1000176
const m_XK_Ydiaeresis = 0x13be
const m_XK_Ygrave = 0x1001ef2
const m_XK_Yhook = 0x1001ef6
const m_XK_Ytilde = 0x1001ef8
const m_XK_Z = 0x005a
const m_XK_Zabovedot = 0x01af
const m_XK_Zacute = 0x01ac
const m_XK_Zcaron = 0x01ae
const m_XK_Zen_Koho = 0xff3d
const m_XK_Zenkaku = 0xff28
const m_XK_Zenkaku_Hankaku = 0xff2a
const m_XK_Zstroke = 0x10001b5
const m_XK_a = 0x0061
const m_XK_aacute = 0x00e1
const m_XK_abelowdot = 0x1001ea1
const m_XK_abovedot = 0x01ff
const m_XK_abreve = 0x01e3
const m_XK_abreveacute = 0x1001eaf
const m_XK_abrevebelowdot = 0x1001eb7
const m_XK_abrevegrave = 0x1001eb1
const m_XK_abrevehook = 0x1001eb3
const m_XK_abrevetilde = 0x1001eb5
const m_XK_acircumflex = 0x00e2
const m_XK_acircumflexacute = 0x1001ea5
const m_XK_acircumflexbelowdot = 0x1001ead
const m_XK_acircumflexgrave = 0x1001ea7
const m_XK_acircumflexhook = 0x1001ea9
const m_XK_acircumflextilde = 0x1001eab
const m_XK_acute = 0x00b4
const m_XK_adiaeresis = 0x00e4
const m_XK_ae = 0x00e6
const m_XK_agrave = 0x00e0
const m_XK_ahook = 0x1001ea3
const m_XK_amacron = 0x03e0
const m_XK_ampersand = 0x0026
const m_XK_aogonek = 0x01b1
const m_XK_apostrophe = 0x0027
const m_XK_approxeq = 0x1002248
const m_XK_aring = 0x00e5
const m_XK_asciicircum = 0x005e
const m_XK_asciitilde = 0x007e
const m_XK_asterisk = 0x002a
const m_XK_at = 0x0040
const m_XK_atilde = 0x00e3
const m_XK_b = 0x0062
const m_XK_babovedot = 0x1001e03
const m_XK_backslash = 0x005c
const m_XK_bar = 0x007c
const m_XK_because = 0x1002235
const m_XK_braceleft = 0x007b
const m_XK_braceright = 0x007d
const m_XK_bracketleft = 0x005b
const m_XK_bracketright = 0x005d
const m_XK_braille_blank = 0x1002800
const m_XK_braille_dot_1 = 0xfff1
const m_XK_braille_dot_10 = 0xfffa
const m_XK_braille_dot_2 = 0xfff2
const m_XK_braille_dot_3 = 0xfff3
const m_XK_braille_dot_4 = 0xfff4
const m_XK_braille_dot_5 = 0xfff5
const m_XK_braille_dot_6 = 0xfff6
const m_XK_braille_dot_7 = 0xfff7
const m_XK_braille_dot_8 = 0xfff8
const m_XK_braille_dot_9 = 0xfff9
const m_XK_braille_dots_1 = 0x1002801
const m_XK_braille_dots_12 = 0x1002803
const m_XK_braille_dots_123 = 0x1002807
const m_XK_braille_dots_1234 = 0x100280f
const m_XK_braille_dots_12345 = 0x100281f
const m_XK_braille_dots_123456 = 0x100283f
const m_XK_braille_dots_1234567 = 0x100287f
const m_XK_braille_dots_12345678 = 0x10028ff
const m_XK_braille_dots_1234568 = 0x10028bf
const m_XK_braille_dots_123457 = 0x100285f
const m_XK_braille_dots_1234578 = 0x10028df
const m_XK_braille_dots_123458 = 0x100289f
const m_XK_braille_dots_12346 = 0x100282f
const m_XK_braille_dots_123467 = 0x100286f
const m_XK_braille_dots_1234678 = 0x10028ef
const m_XK_braille_dots_123468 = 0x10028af
const m_XK_braille_dots_12347 = 0x100284f
const m_XK_braille_dots_123478 = 0x10028cf
const m_XK_braille_dots_12348 = 0x100288f
const m_XK_braille_dots_1235 = 0x1002817
const m_XK_braille_dots_12356 = 0x1002837
const m_XK_braille_dots_123567 = 0x1002877
const m_XK_braille_dots_1235678 = 0x10028f7
const m_XK_braille_dots_123568 = 0x10028b7
const m_XK_braille_dots_12357 = 0x1002857
const m_XK_braille_dots_123578 = 0x10028d7
const m_XK_braille_dots_12358 = 0x1002897
const m_XK_braille_dots_1236 = 0x1002827
const m_XK_braille_dots_12367 = 0x1002867
const m_XK_braille_dots_123678 = 0x10028e7
const m_XK_braille_dots_12368 = 0x10028a7
const m_XK_braille_dots_1237 = 0x1002847
const m_XK_braille_dots_12378 = 0x10028c7
const m_XK_braille_dots_1238 = 0x1002887
const m_XK_braille_dots_124 = 0x100280b
const m_XK_braille_dots_1245 = 0x100281b
const m_XK_braille_dots_12456 = 0x100283b
const m_XK_braille_dots_124567 = 0x100287b
const m_XK_braille_dots_1245678 = 0x10028fb
const m_XK_braille_dots_124568 = 0x10028bb
const m_XK_braille_dots_12457 = 0x100285b
const m_XK_braille_dots_124578 = 0x10028db
const m_XK_braille_dots_12458 = 0x100289b
const m_XK_braille_dots_1246 = 0x100282b
const m_XK_braille_dots_12467 = 0x100286b
const m_XK_braille_dots_124678 = 0x10028eb
const m_XK_braille_dots_12468 = 0x10028ab
const m_XK_braille_dots_1247 = 0x100284b
const m_XK_braille_dots_12478 = 0x10028cb
const m_XK_braille_dots_1248 = 0x100288b
const m_XK_braille_dots_125 = 0x1002813
const m_XK_braille_dots_1256 = 0x1002833
const m_XK_braille_dots_12567 = 0x1002873
const m_XK_braille_dots_125678 = 0x10028f3
const m_XK_braille_dots_12568 = 0x10028b3
const m_XK_braille_dots_1257 = 0x1002853
const m_XK_braille_dots_12578 = 0x10028d3
const m_XK_braille_dots_1258 = 0x1002893
const m_XK_braille_dots_126 = 0x1002823
const m_XK_braille_dots_1267 = 0x1002863
const m_XK_braille_dots_12678 = 0x10028e3
const m_XK_braille_dots_1268 = 0x10028a3
const m_XK_braille_dots_127 = 0x1002843
const m_XK_braille_dots_1278 = 0x10028c3
const m_XK_braille_dots_128 = 0x1002883
const m_XK_braille_dots_13 = 0x1002805
const m_XK_braille_dots_134 = 0x100280d
const m_XK_braille_dots_1345 = 0x100281d
const m_XK_braille_dots_13456 = 0x100283d
const m_XK_braille_dots_134567 = 0x100287d
const m_XK_braille_dots_1345678 = 0x10028fd
const m_XK_braille_dots_134568 = 0x10028bd
const m_XK_braille_dots_13457 = 0x100285d
const m_XK_braille_dots_134578 = 0x10028dd
const m_XK_braille_dots_13458 = 0x100289d
const m_XK_braille_dots_1346 = 0x100282d
const m_XK_braille_dots_13467 = 0x100286d
const m_XK_braille_dots_134678 = 0x10028ed
const m_XK_braille_dots_13468 = 0x10028ad
const m_XK_braille_dots_1347 = 0x100284d
const m_XK_braille_dots_13478 = 0x10028cd
const m_XK_braille_dots_1348 = 0x100288d
const m_XK_braille_dots_135 = 0x1002815
const m_XK_braille_dots_1356 = 0x1002835
const m_XK_braille_dots_13567 = 0x1002875
const m_XK_braille_dots_135678 = 0x10028f5
const m_XK_braille_dots_13568 = 0x10028b5
const m_XK_braille_dots_1357 = 0x1002855
const m_XK_braille_dots_13578 = 0x10028d5
const m_XK_braille_dots_1358 = 0x1002895
const m_XK_braille_dots_136 = 0x1002825
const m_XK_braille_dots_1367 = 0x1002865
const m_XK_braille_dots_13678 = 0x10028e5
const m_XK_braille_dots_1368 = 0x10028a5
const m_XK_braille_dots_137 = 0x1002845
const m_XK_braille_dots_1378 = 0x10028c5
const m_XK_braille_dots_138 = 0x1002885
const m_XK_braille_dots_14 = 0x1002809
const m_XK_braille_dots_145 = 0x1002819
const m_XK_braille_dots_1456 = 0x1002839
const m_XK_braille_dots_14567 = 0x1002879
const m_XK_braille_dots_145678 = 0x10028f9
const m_XK_braille_dots_14568 = 0x10028b9
const m_XK_braille_dots_1457 = 0x1002859
const m_XK_braille_dots_14578 = 0x10028d9
const m_XK_braille_dots_1458 = 0x1002899
const m_XK_braille_dots_146 = 0x1002829
const m_XK_braille_dots_1467 = 0x1002869
const m_XK_braille_dots_14678 = 0x10028e9
const m_XK_braille_dots_1468 = 0x10028a9
const m_XK_braille_dots_147 = 0x1002849
const m_XK_braille_dots_1478 = 0x10028c9
const m_XK_braille_dots_148 = 0x1002889
const m_XK_braille_dots_15 = 0x1002811
const m_XK_braille_dots_156 = 0x1002831
const m_XK_braille_dots_1567 = 0x1002871
const m_XK_braille_dots_15678 = 0x10028f1
const m_XK_braille_dots_1568 = 0x10028b1
const m_XK_braille_dots_157 = 0x1002851
const m_XK_braille_dots_1578 = 0x10028d1
const m_XK_braille_dots_158 = 0x1002891
const m_XK_braille_dots_16 = 0x1002821
const m_XK_braille_dots_167 = 0x1002861
const m_XK_braille_dots_1678 = 0x10028e1
const m_XK_braille_dots_168 = 0x10028a1
const m_XK_braille_dots_17 = 0x1002841
const m_XK_braille_dots_178 = 0x10028c1
const m_XK_braille_dots_18 = 0x1002881
const m_XK_braille_dots_2 = 0x1002802
const m_XK_braille_dots_23 = 0x1002806
const m_XK_braille_dots_234 = 0x100280e
const m_XK_braille_dots_2345 = 0x100281e
const m_XK_braille_dots_23456 = 0x100283e
const m_XK_braille_dots_234567 = 0x100287e
const m_XK_braille_dots_2345678 = 0x10028fe
const m_XK_braille_dots_234568 = 0x10028be
const m_XK_braille_dots_23457 = 0x100285e
const m_XK_braille_dots_234578 = 0x10028de
const m_XK_braille_dots_23458 = 0x100289e
const m_XK_braille_dots_2346 = 0x100282e
const m_XK_braille_dots_23467 = 0x100286e
const m_XK_braille_dots_234678 = 0x10028ee
const m_XK_braille_dots_23468 = 0x10028ae
const m_XK_braille_dots_2347 = 0x100284e
const m_XK_braille_dots_23478 = 0x10028ce
const m_XK_braille_dots_2348 = 0x100288e
const m_XK_braille_dots_235 = 0x1002816
const m_XK_braille_dots_2356 = 0x1002836
const m_XK_braille_dots_23567 = 0x1002876
const m_XK_braille_dots_235678 = 0x10028f6
const m_XK_braille_dots_23568 = 0x10028b6
const m_XK_braille_dots_2357 = 0x1002856
const m_XK_braille_dots_23578 = 0x10028d6
const m_XK_braille_dots_2358 = 0x1002896
const m_XK_braille_dots_236 = 0x1002826
const m_XK_braille_dots_2367 = 0x1002866
const m_XK_braille_dots_23678 = 0x10028e6
const m_XK_braille_dots_2368 = 0x10028a6
const m_XK_braille_dots_237 = 0x1002846
const m_XK_braille_dots_2378 = 0x10028c6
const m_XK_braille_dots_238 = 0x1002886
const m_XK_braille_dots_24 = 0x100280a
const m_XK_braille_dots_245 = 0x100281a
const m_XK_braille_dots_2456 = 0x100283a
const m_XK_braille_dots_24567 = 0x100287a
const m_XK_braille_dots_245678 = 0x10028fa
const m_XK_braille_dots_24568 = 0x10028ba
const m_XK_braille_dots_2457 = 0x100285a
const m_XK_braille_dots_24578 = 0x10028da
const m_XK_braille_dots_2458 = 0x100289a
const m_XK_braille_dots_246 = 0x100282a
const m_XK_braille_dots_2467 = 0x100286a
const m_XK_braille_dots_24678 = 0x10028ea
const m_XK_braille_dots_2468 = 0x10028aa
const m_XK_braille_dots_247 = 0x100284a
const m_XK_braille_dots_2478 = 0x10028ca
const m_XK_braille_dots_248 = 0x100288a
const m_XK_braille_dots_25 = 0x1002812
const m_XK_braille_dots_256 = 0x1002832
const m_XK_braille_dots_2567 = 0x1002872
const m_XK_braille_dots_25678 = 0x10028f2
const m_XK_braille_dots_2568 = 0x10028b2
const m_XK_braille_dots_257 = 0x1002852
const m_XK_braille_dots_2578 = 0x10028d2
const m_XK_braille_dots_258 = 0x1002892
const m_XK_braille_dots_26 = 0x1002822
const m_XK_braille_dots_267 = 0x1002862
const m_XK_braille_dots_2678 = 0x10028e2
const m_XK_braille_dots_268 = 0x10028a2
const m_XK_braille_dots_27 = 0x1002842
const m_XK_braille_dots_278 = 0x10028c2
const m_XK_braille_dots_28 = 0x1002882
const m_XK_braille_dots_3 = 0x1002804
const m_XK_braille_dots_34 = 0x100280c
const m_XK_braille_dots_345 = 0x100281c
const m_XK_braille_dots_3456 = 0x100283c
const m_XK_braille_dots_34567 = 0x100287c
const m_XK_braille_dots_345678 = 0x10028fc
const m_XK_braille_dots_34568 = 0x10028bc
const m_XK_braille_dots_3457 = 0x100285c
const m_XK_braille_dots_34578 = 0x10028dc
const m_XK_braille_dots_3458 = 0x100289c
const m_XK_braille_dots_346 = 0x100282c
const m_XK_braille_dots_3467 = 0x100286c
const m_XK_braille_dots_34678 = 0x10028ec
const m_XK_braille_dots_3468 = 0x10028ac
const m_XK_braille_dots_347 = 0x100284c
const m_XK_braille_dots_3478 = 0x10028cc
const m_XK_braille_dots_348 = 0x100288c
const m_XK_braille_dots_35 = 0x1002814
const m_XK_braille_dots_356 = 0x1002834
const m_XK_braille_dots_3567 = 0x1002874
const m_XK_braille_dots_35678 = 0x10028f4
const m_XK_braille_dots_3568 = 0x10028b4
const m_XK_braille_dots_357 = 0x1002854
const m_XK_braille_dots_3578 = 0x10028d4
const m_XK_braille_dots_358 = 0x1002894
const m_XK_braille_dots_36 = 0x1002824
const m_XK_braille_dots_367 = 0x1002864
const m_XK_braille_dots_3678 = 0x10028e4
const m_XK_braille_dots_368 = 0x10028a4
const m_XK_braille_dots_37 = 0x1002844
const m_XK_braille_dots_378 = 0x10028c4
const m_XK_braille_dots_38 = 0x1002884
const m_XK_braille_dots_4 = 0x1002808
const m_XK_braille_dots_45 = 0x1002818
const m_XK_braille_dots_456 = 0x1002838
const m_XK_braille_dots_4567 = 0x1002878
const m_XK_braille_dots_45678 = 0x10028f8
const m_XK_braille_dots_4568 = 0x10028b8
const m_XK_braille_dots_457 = 0x1002858
const m_XK_braille_dots_4578 = 0x10028d8
const m_XK_braille_dots_458 = 0x1002898
const m_XK_braille_dots_46 = 0x1002828
const m_XK_braille_dots_467 = 0x1002868
const m_XK_braille_dots_4678 = 0x10028e8
const m_XK_braille_dots_468 = 0x10028a8
const m_XK_braille_dots_47 = 0x1002848
const m_XK_braille_dots_478 = 0x10028c8
const m_XK_braille_dots_48 = 0x1002888
const m_XK_braille_dots_5 = 0x1002810
const m_XK_braille_dots_56 = 0x1002830
const m_XK_braille_dots_567 = 0x1002870
const m_XK_braille_dots_5678 = 0x10028f0
const m_XK_braille_dots_568 = 0x10028b0
const m_XK_braille_dots_57 = 0x1002850
const m_XK_braille_dots_578 = 0x10028d0
const m_XK_braille_dots_58 = 0x1002890
const m_XK_braille_dots_6 = 0x1002820
const m_XK_braille_dots_67 = 0x1002860
const m_XK_braille_dots_678 = 0x10028e0
const m_XK_braille_dots_68 = 0x10028a0
const m_XK_braille_dots_7 = 0x1002840
const m_XK_braille_dots_78 = 0x10028c0
const m_XK_braille_dots_8 = 0x1002880
const m_XK_breve = 0x01a2
const m_XK_brokenbar = 0x00a6
const m_XK_c = 0x0063
const m_XK_c_h = 0xfea3
const m_XK_cabovedot = 0x02e5
const m_XK_cacute = 0x01e6
const m_XK_caron = 0x01b7
const m_XK_ccaron = 0x01e8
const m_XK_ccedilla = 0x00e7
const m_XK_ccircumflex = 0x02e6
const m_XK_cedilla = 0x00b8
const m_XK_cent = 0x00a2
const m_XK_ch = 0xfea0
const m_XK_colon = 0x003a
const m_XK_combining_acute = 0x1000301
const m_XK_combining_belowdot = 0x1000323
const m_XK_combining_grave = 0x1000300
const m_XK_combining_hook = 0x1000309
const m_XK_combining_tilde = 0x1000303
const m_XK_comma = 0x002c
const m_XK_containsas = 0x100220b
const m_XK_copyright = 0x00a9
const m_XK_cuberoot = 0x100221b
const m_XK_currency = 0x00a4
const m_XK_d = 0x0064
const m_XK_dabovedot = 0x1001e0b
const m_XK_dcaron = 0x01ef
const m_XK_dead_A = 0xfe81
const m_XK_dead_E = 0xfe83
const m_XK_dead_I = 0xfe85
const m_XK_dead_O = 0xfe87
const m_XK_dead_SCHWA = 0xfe8b
const m_XK_dead_U = 0xfe89
const m_XK_dead_a = 0xfe80
const m_XK_dead_abovecomma = 0xfe64
const m_XK_dead_abovedot = 0xfe56
const m_XK_dead_abovereversedcomma = 0xfe65
const m_XK_dead_abovering = 0xfe58
const m_XK_dead_aboveverticalline = 0xfe91
const m_XK_dead_acute = 0xfe51
const m_XK_dead_belowbreve = 0xfe6b
const m_XK_dead_belowcircumflex = 0xfe69
const m_XK_dead_belowcomma = 0xfe6e
const m_XK_dead_belowdiaeresis = 0xfe6c
const m_XK_dead_belowdot = 0xfe60
const m_XK_dead_belowmacron = 0xfe68
const m_XK_dead_belowring = 0xfe67
const m_XK_dead_belowtilde = 0xfe6a
const m_XK_dead_belowverticalline = 0xfe92
const m_XK_dead_breve = 0xfe55
const m_XK_dead_capital_schwa = 0xfe8b
const m_XK_dead_caron = 0xfe5a
const m_XK_dead_cedilla = 0xfe5b
const m_XK_dead_circumflex = 0xfe52
const m_XK_dead_currency = 0xfe6f
const m_XK_dead_dasia = 0xfe65
const m_XK_dead_diaeresis = 0xfe57
const m_XK_dead_doubleacute = 0xfe59
const m_XK_dead_doublegrave = 0xfe66
const m_XK_dead_e = 0xfe82
const m_XK_dead_grave = 0xfe50
const m_XK_dead_greek = 0xfe8c
const m_XK_dead_hamza = 0xfe8d
const m_XK_dead_hook = 0xfe61
const m_XK_dead_horn = 0xfe62
const m_XK_dead_i = 0xfe84
const m_XK_dead_invertedbreve = 0xfe6d
const m_XK_dead_iota = 0xfe5d
const m_XK_dead_longsolidusoverlay = 0xfe93
const m_XK_dead_lowline = 0xfe90
const m_XK_dead_macron = 0xfe54
const m_XK_dead_o = 0xfe86
const m_XK_dead_ogonek = 0xfe5c
const m_XK_dead_perispomeni = 0xfe53
const m_XK_dead_psili = 0xfe64
const m_XK_dead_schwa = 0xfe8a
const m_XK_dead_semivoiced_sound = 0xfe5f
const m_XK_dead_small_schwa = 0xfe8a
const m_XK_dead_stroke = 0xfe63
const m_XK_dead_tilde = 0xfe53
const m_XK_dead_u = 0xfe88
const m_XK_dead_voiced_sound = 0xfe5e
const m_XK_degree = 0x00b0
const m_XK_diaeresis = 0x00a8
const m_XK_dintegral = 0x100222c
const m_XK_division = 0x00f7
const m_XK_dollar = 0x0024
const m_XK_doubleacute = 0x01bd
const m_XK_dstroke = 0x01f0
const m_XK_e = 0x0065
const m_XK_eabovedot = 0x03ec
const m_XK_eacute = 0x00e9
const m_XK_ebelowdot = 0x1001eb9
const m_XK_ecaron = 0x01ec
const m_XK_ecircumflex = 0x00ea
const m_XK_ecircumflexacute = 0x1001ebf
const m_XK_ecircumflexbelowdot = 0x1001ec7
const m_XK_ecircumflexgrave = 0x1001ec1
const m_XK_ecircumflexhook = 0x1001ec3
const m_XK_ecircumflextilde = 0x1001ec5
const m_XK_ediaeresis = 0x00eb
const m_XK_egrave = 0x00e8
const m_XK_ehook = 0x1001ebb
const m_XK_eightsubscript = 0x1002088
const m_XK_eightsuperior = 0x1002078
const m_XK_elementof = 0x1002208
const m_XK_emacron = 0x03ba
const m_XK_emptyset = 0x1002205
const m_XK_eng = 0x03bf
const m_XK_eogonek = 0x01ea
const m_XK_equal = 0x003d
const m_XK_eth = 0x00f0
const m_XK_etilde = 0x1001ebd
const m_XK_exclam = 0x0021
const m_XK_exclamdown = 0x00a1
const m_XK_ezh = 0x1000292
const m_XK_f = 0x0066
const m_XK_fabovedot = 0x1001e1f
const m_XK_fivesubscript = 0x1002085
const m_XK_fivesuperior = 0x1002075
const m_XK_foursubscript = 0x1002084
const m_XK_foursuperior = 0x1002074
const m_XK_fourthroot = 0x100221c
const m_XK_g = 0x0067
const m_XK_gabovedot = 0x02f5
const m_XK_gbreve = 0x02bb
const m_XK_gcaron = 0x10001e7
const m_XK_gcedilla = 0x03bb
const m_XK_gcircumflex = 0x02f8
const m_XK_grave = 0x0060
const m_XK_greater = 0x003e
const m_XK_guillemetleft = 0x00ab
const m_XK_guillemetright = 0x00bb
const m_XK_guillemotleft = 0x00ab
const m_XK_guillemotright = 0x00bb
const m_XK_h = 0x0068
const m_XK_hcircumflex = 0x02b6
const m_XK_hebrew_aleph = 0x0ce0
const m_XK_hebrew_ayin = 0x0cf2
const m_XK_hebrew_bet = 0x0ce1
const m_XK_hebrew_beth = 0x0ce1
const m_XK_hebrew_chet = 0x0ce7
const m_XK_hebrew_dalet = 0x0ce3
const m_XK_hebrew_daleth = 0x0ce3
const m_XK_hebrew_doublelowline = 0x0cdf
const m_XK_hebrew_finalkaph = 0x0cea
const m_XK_hebrew_finalmem = 0x0ced
const m_XK_hebrew_finalnun = 0x0cef
const m_XK_hebrew_finalpe = 0x0cf3
const m_XK_hebrew_finalzade = 0x0cf5
const m_XK_hebrew_finalzadi = 0x0cf5
const m_XK_hebrew_gimel = 0x0ce2
const m_XK_hebrew_gimmel = 0x0ce2
const m_XK_hebrew_he = 0x0ce4
const m_XK_hebrew_het = 0x0ce7
const m_XK_hebrew_kaph = 0x0ceb
const m_XK_hebrew_kuf = 0x0cf7
const m_XK_hebrew_lamed = 0x0cec
const m_XK_hebrew_mem = 0x0cee
const m_XK_hebrew_nun = 0x0cf0
const m_XK_hebrew_pe = 0x0cf4
const m_XK_hebrew_qoph = 0x0cf7
const m_XK_hebrew_resh = 0x0cf8
const m_XK_hebrew_samech = 0x0cf1
const m_XK_hebrew_samekh = 0x0cf1
const m_XK_hebrew_shin = 0x0cf9
const m_XK_hebrew_taf = 0x0cfa
const m_XK_hebrew_taw = 0x0cfa
const m_XK_hebrew_tet = 0x0ce8
const m_XK_hebrew_teth = 0x0ce8
const m_XK_hebrew_waw = 0x0ce5
const m_XK_hebrew_yod = 0x0ce9
const m_XK_hebrew_zade = 0x0cf6
const m_XK_hebrew_zadi = 0x0cf6
const m_XK_hebrew_zain = 0x0ce6
const m_XK_hebrew_zayin = 0x0ce6
const m_XK_hstroke = 0x02b1
const m_XK_hyphen = 0x00ad
const m_XK_i = 0x0069
const m_XK_iacute = 0x00ed
const m_XK_ibelowdot = 0x1001ecb
const m_XK_ibreve = 0x100012d
const m_XK_icircumflex = 0x00ee
const m_XK_idiaeresis = 0x00ef
const m_XK_idotless = 0x02b9
const m_XK_igrave = 0x00ec
const m_XK_ihook = 0x1001ec9
const m_XK_imacron = 0x03ef
const m_XK_iogonek = 0x03e7
const m_XK_itilde = 0x03b5
const m_XK_j = 0x006a
const m_XK_jcircumflex = 0x02bc
const m_XK_k = 0x006b
const m_XK_kana_A = 0x04b1
const m_XK_kana_CHI = 0x04c1
const m_XK_kana_E = 0x04b4
const m_XK_kana_FU = 0x04cc
const m_XK_kana_HA = 0x04ca
const m_XK_kana_HE = 0x04cd
const m_XK_kana_HI = 0x04cb
const m_XK_kana_HO = 0x04ce
const m_XK_kana_HU = 0x04cc
const m_XK_kana_I = 0x04b2
const m_XK_kana_KA = 0x04b6
const m_XK_kana_KE = 0x04b9
const m_XK_kana_KI = 0x04b7
const m_XK_kana_KO = 0x04ba
const m_XK_kana_KU = 0x04b8
const m_XK_kana_MA = 0x04cf
const m_XK_kana_ME = 0x04d2
const m_XK_kana_MI = 0x04d0
const m_XK_kana_MO = 0x04d3
const m_XK_kana_MU = 0x04d1
const m_XK_kana_N = 0x04dd
const m_XK_kana_NA = 0x04c5
const m_XK_kana_NE = 0x04c8
const m_XK_kana_NI = 0x04c6
const m_XK_kana_NO = 0x04c9
const m_XK_kana_NU = 0x04c7
const m_XK_kana_O = 0x04b5
const m_XK_kana_RA = 0x04d7
const m_XK_kana_RE = 0x04da
const m_XK_kana_RI = 0x04d8
const m_XK_kana_RO = 0x04db
const m_XK_kana_RU = 0x04d9
const m_XK_kana_SA = 0x04bb
const m_XK_kana_SE = 0x04be
const m_XK_kana_SHI = 0x04bc
const m_XK_kana_SO = 0x04bf
const m_XK_kana_SU = 0x04bd
const m_XK_kana_TA = 0x04c0
const m_XK_kana_TE = 0x04c3
const m_XK_kana_TI = 0x04c1
const m_XK_kana_TO = 0x04c4
const m_XK_kana_TSU = 0x04c2
const m_XK_kana_TU = 0x04c2
const m_XK_kana_U = 0x04b3
const m_XK_kana_WA = 0x04dc
const m_XK_kana_WO = 0x04a6
const m_XK_kana_YA = 0x04d4
const m_XK_kana_YO = 0x04d6
const m_XK_kana_YU = 0x04d5
const m_XK_kana_a = 0x04a7
const m_XK_kana_closingbracket = 0x04a3
const m_XK_kana_comma = 0x04a4
const m_XK_kana_conjunctive = 0x04a5
const m_XK_kana_e = 0x04aa
const m_XK_kana_fullstop = 0x04a1
const m_XK_kana_i = 0x04a8
const m_XK_kana_middledot = 0x04a5
const m_XK_kana_o = 0x04ab
const m_XK_kana_openingbracket = 0x04a2
const m_XK_kana_switch = 0xff7e
const m_XK_kana_tsu = 0x04af
const m_XK_kana_tu = 0x04af
const m_XK_kana_u = 0x04a9
const m_XK_kana_ya = 0x04ac
const m_XK_kana_yo = 0x04ae
const m_XK_kana_yu = 0x04ad
const m_XK_kappa = 0x03a2
const m_XK_kcedilla = 0x03f3
const m_XK_kra = 0x03a2
const m_XK_l = 0x006c
const m_XK_lacute = 0x01e5
const m_XK_lbelowdot = 0x1001e37
const m_XK_lcaron = 0x01b5
const m_XK_lcedilla = 0x03b6
const m_XK_less = 0x003c
const m_XK_lstroke = 0x01b3
const m_XK_m = 0x006d
const m_XK_mabovedot = 0x1001e41
const m_XK_macron = 0x00af
const m_XK_masculine = 0x00ba
const m_XK_minus = 0x002d
const m_XK_mu = 0x00b5
const m_XK_multiply = 0x00d7
const m_XK_n = 0x006e
const m_XK_nacute = 0x01f1
const m_XK_ncaron = 0x01f2
const m_XK_ncedilla = 0x03f1
const m_XK_ninesubscript = 0x1002089
const m_XK_ninesuperior = 0x1002079
const m_XK_nobreakspace = 0x00a0
const m_XK_notapproxeq = 0x1002247
const m_XK_notelementof = 0x1002209
const m_XK_notidentical = 0x1002262
const m_XK_notsign = 0x00ac
const m_XK_ntilde = 0x00f1
const m_XK_numbersign = 0x0023
const m_XK_numerosign = 0x06b0
const m_XK_o = 0x006f
const m_XK_oacute = 0x00f3
const m_XK_obarred = 0x1000275
const m_XK_obelowdot = 0x1001ecd
const m_XK_ocaron = 0x10001d2
const m_XK_ocircumflex = 0x00f4
const m_XK_ocircumflexacute = 0x1001ed1
const m_XK_ocircumflexbelowdot = 0x1001ed9
const m_XK_ocircumflexgrave = 0x1001ed3
const m_XK_ocircumflexhook = 0x1001ed5
const m_XK_ocircumflextilde = 0x1001ed7
const m_XK_odiaeresis = 0x00f6
const m_XK_odoubleacute = 0x01f5
const m_XK_oe = 0x13bd
const m_XK_ogonek = 0x01b2
const m_XK_ograve = 0x00f2
const m_XK_ohook = 0x1001ecf
const m_XK_ohorn = 0x10001a1
const m_XK_ohornacute = 0x1001edb
const m_XK_ohornbelowdot = 0x1001ee3
const m_XK_ohorngrave = 0x1001edd
const m_XK_ohornhook = 0x1001edf
const m_XK_ohorntilde = 0x1001ee1
const m_XK_omacron = 0x03f2
const m_XK_onehalf = 0x00bd
const m_XK_onequarter = 0x00bc
const m_XK_onesubscript = 0x1002081
const m_XK_onesuperior = 0x00b9
const m_XK_ooblique = 0x00f8
const m_XK_ordfeminine = 0x00aa
const m_XK_ordmasculine = 0x00ba
const m_XK_oslash = 0x00f8
const m_XK_otilde = 0x00f5
const m_XK_overline = 0x047e
const m_XK_p = 0x0070
const m_XK_pabovedot = 0x1001e57
const m_XK_paragraph = 0x00b6
const m_XK_parenleft = 0x0028
const m_XK_parenright = 0x0029
const m_XK_partdifferential = 0x1002202
const m_XK_percent = 0x0025
const m_XK_period = 0x002e
const m_XK_periodcentered = 0x00b7
const m_XK_plus = 0x002b
const m_XK_plusminus = 0x00b1
const m_XK_prolongedsound = 0x04b0
const m_XK_q = 0x0071
const m_XK_question = 0x003f
const m_XK_questiondown = 0x00bf
const m_XK_quotedbl = 0x0022
const m_XK_quoteleft = 0x0060
const m_XK_quoteright = 0x0027
const m_XK_r = 0x0072
const m_XK_racute = 0x01e0
const m_XK_rcaron = 0x01f8
const m_XK_rcedilla = 0x03b3
const m_XK_registered = 0x00ae
const m_XK_s = 0x0073
const m_XK_sabovedot = 0x1001e61
const m_XK_sacute = 0x01b6
const m_XK_scaron = 0x01b9
const m_XK_scedilla = 0x01ba
const m_XK_schwa = 0x1000259
const m_XK_scircumflex = 0x02fe
const m_XK_script_switch = 0xff7e
const m_XK_section = 0x00a7
const m_XK_semicolon = 0x003b
const m_XK_semivoicedsound = 0x04df
const m_XK_sevensubscript = 0x1002087
const m_XK_sevensuperior = 0x1002077
const m_XK_sixsubscript = 0x1002086
const m_XK_sixsuperior = 0x1002076
const m_XK_slash = 0x002f
const m_XK_space = 0x0020
const m_XK_squareroot = 0x100221a
const m_XK_ssharp = 0x00df
const m_XK_sterling = 0x00a3
const m_XK_stricteq = 0x1002263
const m_XK_t = 0x0074
const m_XK_tabovedot = 0x1001e6b
const m_XK_tcaron = 0x01bb
const m_XK_tcedilla = 0x01fe
const m_XK_thorn = 0x00fe
const m_XK_threequarters = 0x00be
const m_XK_threesubscript = 0x1002083
const m_XK_threesuperior = 0x00b3
const m_XK_tintegral = 0x100222d
const m_XK_tslash = 0x03bc
const m_XK_twosubscript = 0x1002082
const m_XK_twosuperior = 0x00b2
const m_XK_u = 0x0075
const m_XK_uacute = 0x00fa
const m_XK_ubelowdot = 0x1001ee5
const m_XK_ubreve = 0x02fd
const m_XK_ucircumflex = 0x00fb
const m_XK_udiaeresis = 0x00fc
const m_XK_udoubleacute = 0x01fb
const m_XK_ugrave = 0x00f9
const m_XK_uhook = 0x1001ee7
const m_XK_uhorn = 0x10001b0
const m_XK_uhornacute = 0x1001ee9
const m_XK_uhornbelowdot = 0x1001ef1
const m_XK_uhorngrave = 0x1001eeb
const m_XK_uhornhook = 0x1001eed
const m_XK_uhorntilde = 0x1001eef
const m_XK_umacron = 0x03fe
const m_XK_underscore = 0x005f
const m_XK_uogonek = 0x03f9
const m_XK_uring = 0x01f9
const m_XK_utilde = 0x03fd
const m_XK_v = 0x0076
const m_XK_voicedsound = 0x04de
const m_XK_w = 0x0077
const m_XK_wacute = 0x1001e83
const m_XK_wcircumflex = 0x1000175
const m_XK_wdiaeresis = 0x1001e85
const m_XK_wgrave = 0x1001e81
const m_XK_x = 0x0078
const m_XK_xabovedot = 0x1001e8b
const m_XK_y = 0x0079
const m_XK_yacute = 0x00fd
const m_XK_ybelowdot = 0x1001ef5
const m_XK_ycircumflex = 0x1000177
const m_XK_ydiaeresis = 0x00ff
const m_XK_yen = 0x00a5
const m_XK_ygrave = 0x1001ef3
const m_XK_yhook = 0x1001ef7
const m_XK_ytilde = 0x1001ef9
const m_XK_z = 0x007a
const m_XK_zabovedot = 0x01bf
const m_XK_zacute = 0x01bc
const m_XK_zcaron = 0x01be
const m_XK_zerosubscript = 0x1002080
const m_XK_zerosuperior = 0x1002070
const m_XK_zstroke = 0x10001b6
const m_XLookupBoth = 4
const m_XLookupChars = 2
const m_XLookupKeySym = 3
const m_XLookupNone = 1
const m_XMD_H = 1
const m_XMUTEX_INITIALIZER = "PTHREAD_MUTEX_INITIALIZER"
const m_XNArea = "area"
const m_XNAreaNeeded = "areaNeeded"
const m_XNBackground = "background"
const m_XNBackgroundPixmap = "backgroundPixmap"
const m_XNBaseFontName = "baseFontName"
const m_XNClientWindow = "clientWindow"
const m_XNColormap = "colorMap"
const m_XNContextualDrawing = "contextualDrawing"
const m_XNCursor = "cursor"
const m_XNDefaultString = "defaultString"
const m_XNDestroyCallback = "destroyCallback"
const m_XNDirectionalDependentDrawing = "directionalDependentDrawing"
const m_XNFilterEvents = "filterEvents"
const m_XNFocusWindow = "focusWindow"
const m_XNFontInfo = "fontInfo"
const m_XNFontSet = "fontSet"
const m_XNForeground = "foreground"
const m_XNGeometryCallback = "geometryCallback"
const m_XNHotKey = "hotKey"
const m_XNHotKeyState = "hotKeyState"
const m_XNInputStyle = "inputStyle"
const m_XNLineSpace = "lineSpace"
const m_XNMissingCharSet = "missingCharSet"
const m_XNOMAutomatic = "omAutomatic"
const m_XNOrientation = "orientation"
const m_XNPreeditAttributes = "preeditAttributes"
const m_XNPreeditCaretCallback = "preeditCaretCallback"
const m_XNPreeditDoneCallback = "preeditDoneCallback"
const m_XNPreeditDrawCallback = "preeditDrawCallback"
const m_XNPreeditStartCallback = "preeditStartCallback"
const m_XNPreeditState = "preeditState"
const m_XNPreeditStateNotifyCallback = "preeditStateNotifyCallback"
const m_XNQueryICValuesList = "queryICValuesList"
const m_XNQueryIMValuesList = "queryIMValuesList"
const m_XNQueryInputStyle = "queryInputStyle"
const m_XNQueryOrientation = "queryOrientation"
const m_XNR6PreeditCallback = "r6PreeditCallback"
const m_XNRequiredCharSet = "requiredCharSet"
const m_XNResetState = "resetState"
const m_XNResourceClass = "resourceClass"
const m_XNResourceName = "resourceName"
const m_XNSeparatorofNestedList = "separatorofNestedList"
const m_XNSpotLocation = "spotLocation"
const m_XNStatusAttributes = "statusAttributes"
const m_XNStatusDoneCallback = "statusDoneCallback"
const m_XNStatusDrawCallback = "statusDrawCallback"
const m_XNStatusStartCallback = "statusStartCallback"
const m_XNStdColormap = "stdColorMap"
const m_XNStringConversion = "stringConversion"
const m_XNStringConversionCallback = "stringConversionCallback"
const m_XNVaNestedList = "XNVaNestedList"
const m_XNVisiblePosition = "visiblePosition"
const m_XNegative = 0x0010
const m_XTHREADS = 1
const m_XUSE_MTSAFE_API = 1
const m_XValue = 0x0001
const m_XYBitmap = 0
const m_XYPixmap = 1
const m_X_AllocColor = 84
const m_X_AllocColorCells = 86
const m_X_AllocColorPlanes = 87
const m_X_AllocNamedColor = 85
const m_X_AllowEvents = 35
const m_X_Bell = 104
const m_X_ChangeActivePointerGrab = 30
const m_X_ChangeGC = 56
const m_X_ChangeHosts = 109
const m_X_ChangeKeyboardControl = 102
const m_X_ChangeKeyboardMapping = 100
const m_X_ChangePointerControl = 105
const m_X_ChangeProperty = 18
const m_X_ChangeSaveSet = 6
const m_X_ChangeWindowAttributes = 2
const m_X_CirculateWindow = 13
const m_X_ClearArea = 61
const m_X_CloseFont = 46
const m_X_ConfigureWindow = 12
const m_X_ConvertSelection = 24
const m_X_CopyArea = 62
const m_X_CopyColormapAndFree = 80
const m_X_CopyGC = 57
const m_X_CopyPlane = 63
const m_X_CreateColormap = 78
const m_X_CreateCursor = 93
const m_X_CreateGC = 55
const m_X_CreateGlyphCursor = 94
const m_X_CreatePixmap = 53
const m_X_CreateWindow = 1
const m_X_DeleteProperty = 19
const m_X_DestroySubwindows = 5
const m_X_DestroyWindow = 4
const m_X_Error = 0
const m_X_FillPoly = 69
const m_X_ForceScreenSaver = 115
const m_X_FreeColormap = 79
const m_X_FreeColors = 88
const m_X_FreeCursor = 95
const m_X_FreeGC = 60
const m_X_FreePixmap = 54
const m_X_GetAtomName = 17
const m_X_GetFontPath = 52
const m_X_GetGeometry = 14
const m_X_GetImage = 73
const m_X_GetInputFocus = 43
const m_X_GetKeyboardControl = 103
const m_X_GetKeyboardMapping = 101
const m_X_GetModifierMapping = 119
const m_X_GetMotionEvents = 39
const m_X_GetPointerControl = 106
const m_X_GetPointerMapping = 117
const m_X_GetProperty = 20
const m_X_GetScreenSaver = 108
const m_X_GetSelectionOwner = 23
const m_X_GetWindowAttributes = 3
const m_X_GrabButton = 28
const m_X_GrabKey = 33
const m_X_GrabKeyboard = 31
const m_X_GrabPointer = 26
const m_X_GrabServer = 36
const m_X_HAVE_UTF8_STRING = 1
const m_X_ImageText16 = 77
const m_X_ImageText8 = 76
const m_X_InstallColormap = 81
const m_X_InternAtom = 16
const m_X_KillClient = 113
const m_X_ListExtensions = 99
const m_X_ListFonts = 49
const m_X_ListFontsWithInfo = 50
const m_X_ListHosts = 110
const m_X_ListInstalledColormaps = 83
const m_X_ListProperties = 21
const m_X_LookupColor = 92
const m_X_MapSubwindows = 9
const m_X_MapWindow = 8
const m_X_NoOperation = 127
const m_X_OpenFont = 45
const m_X_PROTOCOL = 11
const m_X_PROTOCOL_REVISION = 0
const m_X_PolyArc = 68
const m_X_PolyFillArc = 71
const m_X_PolyFillRectangle = 70
const m_X_PolyLine = 65
const m_X_PolyPoint = 64
const m_X_PolyRectangle = 67
const m_X_PolySegment = 66
const m_X_PolyText16 = 75
const m_X_PolyText8 = 74
const m_X_PutImage = 72
const m_X_QueryBestSize = 97
const m_X_QueryColors = 91
const m_X_QueryExtension = 98
const m_X_QueryFont = 47
const m_X_QueryKeymap = 44
const m_X_QueryPointer = 38
const m_X_QueryTextExtents = 48
const m_X_QueryTree = 15
const m_X_RecolorCursor = 96
const m_X_RenderAddGlyphs = 20
const m_X_RenderAddGlyphsFromPicture = 21
const m_X_RenderAddTraps = 32
const m_X_RenderChangePicture = 5
const m_X_RenderColorTrapezoids = 14
const m_X_RenderColorTriangles = 15
const m_X_RenderComposite = 8
const m_X_RenderCompositeGlyphs16 = 24
const m_X_RenderCompositeGlyphs32 = 25
const m_X_RenderCompositeGlyphs8 = 23
const m_X_RenderCreateAnimCursor = 31
const m_X_RenderCreateConicalGradient = 36
const m_X_RenderCreateCursor = 27
const m_X_RenderCreateGlyphSet = 17
const m_X_RenderCreateLinearGradient = 34
const m_X_RenderCreatePicture = 4
const m_X_RenderCreateRadialGradient = 35
const m_X_RenderCreateSolidFill = 33
const m_X_RenderFillRectangles = 26
const m_X_RenderFreeGlyphSet = 19
const m_X_RenderFreeGlyphs = 22
const m_X_RenderFreePicture = 7
const m_X_RenderQueryDithers = 3
const m_X_RenderQueryFilters = 29
const m_X_RenderQueryPictFormats = 1
const m_X_RenderQueryPictIndexValues = 2
const m_X_RenderQueryVersion = 0
const m_X_RenderReferenceGlyphSet = 18
const m_X_RenderScale = 9
const m_X_RenderSetPictureClipRectangles = 6
const m_X_RenderSetPictureFilter = 30
const m_X_RenderSetPictureTransform = 28
const m_X_RenderTrapezoids = 10
const m_X_RenderTriFan = 13
const m_X_RenderTriStrip = 12
const m_X_RenderTriangles = 11
const m_X_ReparentWindow = 7
const m_X_Reply = 1
const m_X_RotateProperties = 114
const m_X_SendEvent = 25
const m_X_SetAccessControl = 111
const m_X_SetClipRectangles = 59
const m_X_SetCloseDownMode = 112
const m_X_SetDashes = 58
const m_X_SetFontPath = 51
const m_X_SetInputFocus = 42
const m_X_SetModifierMapping = 118
const m_X_SetPointerMapping = 116
const m_X_SetScreenSaver = 107
const m_X_SetSelectionOwner = 22
const m_X_StoreColors = 89
const m_X_StoreNamedColor = 90
const m_X_TCP_PORT = 6000
const m_X_TranslateCoords = 40
const m_X_UngrabButton = 29
const m_X_UngrabKey = 34
const m_X_UngrabKeyboard = 32
const m_X_UngrabPointer = 27
const m_X_UngrabServer = 37
const m_X_UninstallColormap = 82
const m_X_UnmapSubwindows = 11
const m_X_UnmapWindow = 10
const m_X_WarpPointer = 41
const m_XlibSpecificationRelease = 6
const m_YNegative = 0x0020
const m_YSorted = 1
const m_YValue = 0x0002
const m_YXBanded = 3
const m_YXSorted = 2
const m_ZLNSPERBATCH = 1024
const m_ZPixmap = 2
const m_ZRCTSPERBATCH = 256
const m_ZoomState = 2
const m__BYTE_ORDER = "__BYTE_ORDER__"
const m__LP64 = 1
const m__NCPUBITS = "_BITSET_BITS"
const m__PDP_ENDIAN = "__ORDER_PDP_ENDIAN__"
const m__QUAD_HIGHWORD = 1
const m__QUAD_LOWWORD = 0
const m__SIG_MAXSIG = 128
const m__SIG_WORDS = 4
const m__THREAD_SAFE = 1
const m__X11_XLIBINT_H_ = 1
const m__XBCOPYFUNC = "_Xbcopy"
const m__X_INLINE = "inline"
const m__X_RESTRICT_KYWD = "restrict"
const m__Xconst = "const"
const m___ATOMIC_ACQUIRE = 2
const m___ATOMIC_ACQ_REL = 4
const m___ATOMIC_CONSUME = 1
const m___ATOMIC_RELAXED = 0
const m___ATOMIC_RELEASE = 3
const m___ATOMIC_SEQ_CST = 5
const m___BIGGEST_ALIGNMENT__ = 16
const m___BITINT_MAXWIDTH__ = 8388608
const m___BOOL_WIDTH__ = 8
const m___BSD_VISIBLE = 1
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___CCGO__ = 1
const m___CC_SUPPORTS_DYNAMIC_ARRAY_INIT = 1
const m___CC_SUPPORTS_INLINE = 1
const m___CC_SUPPORTS_VARADIC_XXX = 1
const m___CC_SUPPORTS_WARNING = 1
const m___CC_SUPPORTS___FUNC__ = 1
const m___CC_SUPPORTS___INLINE = 1
const m___CC_SUPPORTS___INLINE__ = 1
const m___CHAR_BIT = 8
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
const m___DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const m___ELF__ = 1
const m___EXT1_VISIBLE = 1
const m___FINITE_MATH_ONLY__ = 0
const m___FLOAT128__ = 1
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
const m___FUNCTION__ = "__func__"
const m___FXSR__ = 1
const m___FreeBSD__ = 14
const m___FreeBSD_cc_version = 1400006
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
const m___GCC_HAVE_DWARF2_CFI_ASM = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const m___GNUCLIKE_ASM = 3
const m___GNUCLIKE_BUILTIN_CONSTANT_P = 1
const m___GNUCLIKE_BUILTIN_MEMCPY = 1
const m___GNUCLIKE_BUILTIN_NEXT_ARG = 1
const m___GNUCLIKE_BUILTIN_STDARG = 1
const m___GNUCLIKE_BUILTIN_VAALIST = 1
const m___GNUCLIKE_BUILTIN_VARARGS = 1
const m___GNUCLIKE_CTOR_SECTION_HANDLING = 1
const m___GNUCLIKE___SECTION = 1
const m___GNUCLIKE___TYPEOF = 1
const m___GNUC_MINOR__ = 2
const m___GNUC_PATCHLEVEL__ = 1
const m___GNUC_STDC_INLINE__ = 1
const m___GNUC_VA_LIST_COMPATIBILITY = 1
const m___GNUC__ = 4
const m___GXX_ABI_VERSION = 1002
const m___INT16_FMTd__ = "hd"
const m___INT16_FMTi__ = "hi"
const m___INT16_MAX__ = 32767
const m___INT16_TYPE__ = "short"
const m___INT32_FMTd__ = "d"
const m___INT32_FMTi__ = "i"
const m___INT32_MAX__ = 2147483647
const m___INT32_TYPE__ = "int"
const m___INT64_C_SUFFIX__ = "L"
const m___INT64_FMTd__ = "ld"
const m___INT64_FMTi__ = "li"
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
const m___INT_FAST64_FMTd__ = "ld"
const m___INT_FAST64_FMTi__ = "li"
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
const m___INT_LEAST64_FMTd__ = "ld"
const m___INT_LEAST64_FMTi__ = "li"
const m___INT_LEAST64_MAX__ = 9223372036854775807
const m___INT_LEAST64_WIDTH__ = 64
const m___INT_LEAST8_FMTd__ = "hhd"
const m___INT_LEAST8_FMTi__ = "hhi"
const m___INT_LEAST8_MAX__ = 127
const m___INT_LEAST8_WIDTH__ = 8
const m___INT_MAX = 0x7fffffff
const m___INT_MAX__ = 2147483647
const m___INT_WIDTH__ = 32
const m___ISO_C_VISIBLE = 2011
const m___KPRINTF_ATTRIBUTE__ = 1
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
const m___LITTLE_ENDIAN__ = 1
const m___LLONG_MAX = 0x7fffffffffffffff
const m___LLONG_WIDTH__ = 64
const m___LONG_BIT = 64
const m___LONG_LONG_MAX__ = 9223372036854775807
const m___LONG_MAX = 0x7fffffffffffffff
const m___LONG_MAX__ = 9223372036854775807
const m___LONG_WIDTH__ = 64
const m___LP64__ = 1
const m___MEMORY_SCOPE_DEVICE = 1
const m___MEMORY_SCOPE_SINGLE = 4
const m___MEMORY_SCOPE_SYSTEM = 0
const m___MEMORY_SCOPE_WRKGRP = 2
const m___MEMORY_SCOPE_WVFRNT = 3
const m___MMX__ = 1
const m___NO_INLINE__ = 1
const m___NO_MATH_ERRNO__ = 1
const m___NO_MATH_INLINES = 1
const m___OBJC_BOOL_IS_BOOL = 0
const m___OFF_MAX = "__LONG_MAX"
const m___OFF_MIN = "__LONG_MIN"
const m___OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const m___OPENCL_MEMORY_SCOPE_DEVICE = 2
const m___OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const m___OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const m___OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const m___ORDER_BIG_ENDIAN__ = 4321
const m___ORDER_LITTLE_ENDIAN__ = 1234
const m___ORDER_PDP_ENDIAN__ = 3412
const m___POINTER_WIDTH__ = 64
const m___POSIX_VISIBLE = 200809
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTRDIFF_FMTd__ = "ld"
const m___PTRDIFF_FMTi__ = "li"
const m___PTRDIFF_MAX__ = 9223372036854775807
const m___PTRDIFF_WIDTH__ = 64
const m___QUAD_MAX = "__LONG_MAX"
const m___QUAD_MIN = "__LONG_MIN"
const m___SCHAR_MAX = 0x7f
const m___SCHAR_MAX__ = 127
const m___SEG_FS = 1
const m___SEG_GS = 1
const m___SHRT_MAX = 0x7fff
const m___SHRT_MAX__ = 32767
const m___SHRT_WIDTH__ = 16
const m___SIG_ATOMIC_MAX__ = 2147483647
const m___SIG_ATOMIC_WIDTH__ = 32
const m___SIZEOF_DOUBLE__ = 8
const m___SIZEOF_FLOAT128__ = 16
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
const m___SIZE_T_MAX = "__ULONG_MAX"
const m___SIZE_WIDTH__ = 64
const m___SSE2_MATH__ = 1
const m___SSE2__ = 1
const m___SSE_MATH__ = 1
const m___SSE__ = 1
const m___SSIZE_MAX = "__LONG_MAX"
const m___STDC_HOSTED__ = 1
const m___STDC_MB_MIGHT_NEQ_WC__ = 1
const m___STDC_UTF_16__ = 1
const m___STDC_UTF_32__ = 1
const m___STDC_VERSION__ = 201710
const m___STDC__ = 1
const m___UCHAR_MAX = 0xff
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
const m___UINT64_C_SUFFIX__ = "UL"
const m___UINT64_FMTX__ = "lX"
const m___UINT64_FMTo__ = "lo"
const m___UINT64_FMTu__ = "lu"
const m___UINT64_FMTx__ = "lx"
const m___UINT64_MAX__ = 18446744073709551615
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
const m___UINT_FAST64_FMTX__ = "lX"
const m___UINT_FAST64_FMTo__ = "lo"
const m___UINT_FAST64_FMTu__ = "lu"
const m___UINT_FAST64_FMTx__ = "lx"
const m___UINT_FAST64_MAX__ = 18446744073709551615
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
const m___UINT_LEAST64_FMTX__ = "lX"
const m___UINT_LEAST64_FMTo__ = "lo"
const m___UINT_LEAST64_FMTu__ = "lu"
const m___UINT_LEAST64_FMTx__ = "lx"
const m___UINT_LEAST64_MAX__ = 18446744073709551615
const m___UINT_LEAST8_FMTX__ = "hhX"
const m___UINT_LEAST8_FMTo__ = "hho"
const m___UINT_LEAST8_FMTu__ = "hhu"
const m___UINT_LEAST8_FMTx__ = "hhx"
const m___UINT_LEAST8_MAX__ = 255
const m___UINT_MAX = 0xffffffff
const m___ULLONG_MAX = "0xffffffffffffffffU"
const m___ULONG_MAX = 0xffffffffffffffff
const m___UQUAD_MAX = "__ULONG_MAX"
const m___USHRT_MAX = 0xffff
const m___VERSION__ = "FreeBSD Clang 18.1.6 (https://github.com/llvm/llvm-project.git llvmorg-18.1.6-0-g1118c2e05e67)"
const m___WCHAR_MAX = "__INT_MAX"
const m___WCHAR_MAX__ = 2147483647
const m___WCHAR_MIN = "__INT_MIN"
const m___WCHAR_TYPE__ = "int"
const m___WCHAR_WIDTH__ = 32
const m___WINT_MAX__ = 2147483647
const m___WINT_TYPE__ = "int"
const m___WINT_WIDTH__ = 32
const m___WORDSIZE = 64
const m___WORD_BIT = 32
const m___XSI_VISIBLE = 700
const m___amd64 = 1
const m___amd64__ = 1
const m___clang__ = 1
const m___clang_literal_encoding__ = "UTF-8"
const m___clang_major__ = 18
const m___clang_minor__ = 1
const m___clang_patchlevel__ = 6
const m___clang_version__ = "18.1.6 (https://github.com/llvm/llvm-project.git llvmorg-18.1.6-0-g1118c2e05e67)"
const m___clang_wide_literal_encoding__ = "UTF-32"
const m___code_model_small__ = 1
const m___const = "const"
const m___has_extension = "__has_feature"
const m___k8 = 1
const m___k8__ = 1
const m___llvm__ = 1
const m___restrict = "restrict"
const m___restrict_arr = "restrict"
const m___signed = "signed"
const m___tune_k8__ = 1
const m___unix = 1
const m___unix__ = 1
const m___volatile = "volatile"
const m___writeonly = "__unused"
const m___x86_64 = 1
const m___x86_64__ = 1
const m_fds_bits = "__fds_bits"
const m_pthread_attr_default = "NULL"
const m_pthread_condattr_default = "NULL"
const m_pthread_mutexattr_default = "NULL"
const m_sz_xAllocColorCellsReply = 32
const m_sz_xAllocColorCellsReq = 12
const m_sz_xAllocColorPlanesReply = 32
const m_sz_xAllocColorPlanesReq = 16
const m_sz_xAllocColorReply = 32
const m_sz_xAllocColorReq = 16
const m_sz_xAllocNamedColorReply = 32
const m_sz_xAllocNamedColorReq = 12
const m_sz_xAllowEventsReq = 8
const m_sz_xAnimCursorElt = 8
const m_sz_xArc = 12
const m_sz_xBellReq = 4
const m_sz_xChangeActivePointerGrabReq = 16
const m_sz_xChangeGCReq = 12
const m_sz_xChangeHostsReq = 8
const m_sz_xChangeKeyboardControlReq = 8
const m_sz_xChangeKeyboardMappingReq = 8
const m_sz_xChangeModeReq = 4
const m_sz_xChangePointerControlReq = 12
const m_sz_xChangePropertyReq = 24
const m_sz_xChangeSaveSetReq = 8
const m_sz_xChangeWindowAttributesReq = 12
const m_sz_xCharInfo = 12
const m_sz_xCirculateWindowReq = 8
const m_sz_xClearAreaReq = 16
const m_sz_xColorItem = 12
const m_sz_xConfigureWindowReq = 12
const m_sz_xConnClientPrefix = 12
const m_sz_xConnSetup = 32
const m_sz_xConnSetupPrefix = 8
const m_sz_xConvertSelectionReq = 24
const m_sz_xCopyAreaReq = 28
const m_sz_xCopyColormapAndFreeReq = 12
const m_sz_xCopyGCReq = 16
const m_sz_xCopyPlaneReq = 32
const m_sz_xCreateColormapReq = 16
const m_sz_xCreateCursorReq = 32
const m_sz_xCreateGCReq = 16
const m_sz_xCreateGlyphCursorReq = 32
const m_sz_xCreatePixmapReq = 16
const m_sz_xCreateWindowReq = 32
const m_sz_xDeletePropertyReq = 12
const m_sz_xDepth = 8
const m_sz_xDirectFormat = 16
const m_sz_xError = 32
const m_sz_xEvent = 32
const m_sz_xFillPolyReq = 16
const m_sz_xFontProp = 8
const m_sz_xForceScreenSaverReq = 4
const m_sz_xFreeColorsReq = 12
const m_sz_xGenericReply = 32
const m_sz_xGetAtomNameReply = 32
const m_sz_xGetFontPathReply = 32
const m_sz_xGetGeometryReply = 32
const m_sz_xGetImageReply = 32
const m_sz_xGetImageReq = 20
const m_sz_xGetInputFocusReply = 32
const m_sz_xGetKeyboardControlReply = 52
const m_sz_xGetKeyboardMappingReply = 32
const m_sz_xGetKeyboardMappingReq = 8
const m_sz_xGetModifierMappingReply = 32
const m_sz_xGetMotionEventsReply = 32
const m_sz_xGetMotionEventsReq = 16
const m_sz_xGetPointerControlReply = 32
const m_sz_xGetPointerMappingReply = 32
const m_sz_xGetPropertyReply = 32
const m_sz_xGetPropertyReq = 24
const m_sz_xGetScreenSaverReply = 32
const m_sz_xGetSelectionOwnerReply = 32
const m_sz_xGetWindowAttributesReply = 44
const m_sz_xGlyphElt = 8
const m_sz_xGlyphInfo = 12
const m_sz_xGrabButtonReq = 24
const m_sz_xGrabKeyReq = 16
const m_sz_xGrabKeyboardReply = 32
const m_sz_xGrabKeyboardReq = 16
const m_sz_xGrabPointerReply = 32
const m_sz_xGrabPointerReq = 24
const m_sz_xHostEntry = 4
const m_sz_xImageText16Req = 16
const m_sz_xImageText8Req = 16
const m_sz_xImageTextReq = 16
const m_sz_xIndexValue = 12
const m_sz_xInternAtomReply = 32
const m_sz_xInternAtomReq = 8
const m_sz_xKeymapEvent = 32
const m_sz_xLineFixed = 16
const m_sz_xListExtensionsReply = 32
const m_sz_xListFontsReply = 32
const m_sz_xListFontsReq = 8
const m_sz_xListFontsWithInfoReply = 60
const m_sz_xListFontsWithInfoReq = 8
const m_sz_xListHostsReply = 32
const m_sz_xListHostsReq = 4
const m_sz_xListInstalledColormapsReply = 32
const m_sz_xListPropertiesReply = 32
const m_sz_xLookupColorReply = 32
const m_sz_xLookupColorReq = 12
const m_sz_xOpenFontReq = 12
const m_sz_xPictDepth = 8
const m_sz_xPictFormInfo = 28
const m_sz_xPictScreen = 8
const m_sz_xPictVisual = 8
const m_sz_xPixmapFormat = 8
const m_sz_xPoint = 4
const m_sz_xPointFixed = 8
const m_sz_xPolyArcReq = 12
const m_sz_xPolyFillArcReq = 12
const m_sz_xPolyFillRectangleReq = 12
const m_sz_xPolyLineReq = 12
const m_sz_xPolyPointReq = 12
const m_sz_xPolyRectangleReq = 12
const m_sz_xPolySegmentReq = 12
const m_sz_xPolyText16Req = 16
const m_sz_xPolyText8Req = 16
const m_sz_xPolyTextReq = 16
const m_sz_xPropIconSize = 24
const m_sz_xPutImageReq = 24
const m_sz_xQueryBestSizeReply = 32
const m_sz_xQueryBestSizeReq = 12
const m_sz_xQueryColorsReply = 32
const m_sz_xQueryColorsReq = 8
const m_sz_xQueryExtensionReply = 32
const m_sz_xQueryExtensionReq = 8
const m_sz_xQueryFontReply = 60
const m_sz_xQueryKeymapReply = 40
const m_sz_xQueryPointerReply = 32
const m_sz_xQueryTextExtentsReply = 32
const m_sz_xQueryTextExtentsReq = 8
const m_sz_xQueryTreeReply = 32
const m_sz_xRecolorCursorReq = 20
const m_sz_xRectangle = 8
const m_sz_xRenderAddGlyphsReq = 12
const m_sz_xRenderAddTrapsReq = 12
const m_sz_xRenderChangePictureReq = 12
const m_sz_xRenderColor = 8
const m_sz_xRenderCompositeGlyphs16Req = 28
const m_sz_xRenderCompositeGlyphs32Req = 28
const m_sz_xRenderCompositeGlyphs8Req = 28
const m_sz_xRenderCompositeReq = 36
const m_sz_xRenderCreateAnimCursorReq = 8
const m_sz_xRenderCreateConicalGradientReq = 24
const m_sz_xRenderCreateCursorReq = 16
const m_sz_xRenderCreateGlyphSetReq = 12
const m_sz_xRenderCreateLinearGradientReq = 28
const m_sz_xRenderCreatePictureReq = 20
const m_sz_xRenderCreateRadialGradientReq = 36
const m_sz_xRenderCreateSolidFillReq = 16
const m_sz_xRenderFillRectanglesReq = 20
const m_sz_xRenderFreeGlyphSetReq = 8
const m_sz_xRenderFreeGlyphsReq = 8
const m_sz_xRenderFreePictureReq = 8
const m_sz_xRenderQueryFiltersReply = 32
const m_sz_xRenderQueryFiltersReq = 8
const m_sz_xRenderQueryPictFormatsReply = 32
const m_sz_xRenderQueryPictFormatsReq = 4
const m_sz_xRenderQueryPictIndexValuesReply = 32
const m_sz_xRenderQueryPictIndexValuesReq = 8
const m_sz_xRenderQueryVersionReply = 32
const m_sz_xRenderQueryVersionReq = 12
const m_sz_xRenderReferenceGlyphSetReq = 24
const m_sz_xRenderScaleReq = 32
const m_sz_xRenderSetPictureClipRectanglesReq = 12
const m_sz_xRenderSetPictureFilterReq = 12
const m_sz_xRenderSetPictureTransformReq = 44
const m_sz_xRenderTransform = 36
const m_sz_xRenderTrapezoidsReq = 24
const m_sz_xRenderTriFanReq = 24
const m_sz_xRenderTriStripReq = 24
const m_sz_xRenderTrianglesReq = 24
const m_sz_xReparentWindowReq = 16
const m_sz_xReply = 32
const m_sz_xReq = 4
const m_sz_xResourceReq = 8
const m_sz_xRotatePropertiesReq = 12
const m_sz_xSegment = 8
const m_sz_xSendEventReq = 44
const m_sz_xSetAccessControlReq = 4
const m_sz_xSetClipRectanglesReq = 12
const m_sz_xSetCloseDownModeReq = 4
const m_sz_xSetDashesReq = 12
const m_sz_xSetFontPathReq = 8
const m_sz_xSetInputFocusReq = 12
const m_sz_xSetMappingReply = 32
const m_sz_xSetModifierMappingReply = 32
const m_sz_xSetModifierMappingReq = 4
const m_sz_xSetPointerMappingReply = 32
const m_sz_xSetPointerMappingReq = 4
const m_sz_xSetScreenSaverReq = 12
const m_sz_xSetSelectionOwnerReq = 16
const m_sz_xSpanFix = 12
const m_sz_xStoreColorsReq = 8
const m_sz_xStoreNamedColorReq = 16
const m_sz_xTextElt = 2
const m_sz_xTimecoord = 8
const m_sz_xTranslateCoordsReply = 32
const m_sz_xTranslateCoordsReq = 16
const m_sz_xTrap = 24
const m_sz_xTrapezoid = 40
const m_sz_xTriangle = 24
const m_sz_xUngrabButtonReq = 12
const m_sz_xUngrabKeyReq = 12
const m_sz_xVisualType = 24
const m_sz_xWarpPointerReq = 24
const m_sz_xWindowRoot = 40
const m_sz_xrgb = 8
const m_unix = 1
const m_xFalse = 0
const m_xTrue = 1
const m_xfree = "free"
const m_xmalloc = "malloc"
const m_xthread_self = "pthread_self"

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

type t__int_least8_t = int8

type t__int_least16_t = int16

type t__int_least32_t = int32

type t__int_least64_t = int64

type t__intmax_t = int64

type t__uint_least8_t = uint8

type t__uint_least16_t = uint16

type t__uint_least32_t = uint32

type t__uint_least64_t = uint64

type t__uintmax_t = uint64

type t__intptr_t = int64

type t__intfptr_t = int64

type t__uintptr_t = uint64

type t__uintfptr_t = uint64

type t__vm_offset_t = uint64

type t__vm_size_t = uint64

type t__size_t = uint64

type t__ssize_t = int64

type t__ptrdiff_t = int64

type t__clock_t = int32

type t__critical_t = int64

type t__double_t = float64

type t__float_t = float32

type t__int_fast8_t = int32

type t__int_fast16_t = int32

type t__int_fast32_t = int32

type t__int_fast64_t = int64

type t__register_t = int64

type t__segsz_t = int64

type t__time_t = int64

type t__uint_fast8_t = uint32

type t__uint_fast16_t = uint32

type t__uint_fast32_t = uint32

type t__uint_fast64_t = uint64

type t__u_register_t = uint64

type t__vm_paddr_t = uint64

type T___wchar_t = int32

type t__blksize_t = int32

type t__blkcnt_t = int64

type t__clockid_t = int32

type t__fflags_t = uint32

type t__fsblkcnt_t = uint64

type t__fsfilcnt_t = uint64

type t__gid_t = uint32

type t__id_t = int64

type t__ino_t = uint64

type t__key_t = int64

type t__lwpid_t = int32

type t__mode_t = uint16

type t__accmode_t = int32

type t__nl_item = int32

type t__nlink_t = uint64

type t__off_t = int64

type t__off64_t = int64

type t__pid_t = int32

type t__sbintime_t = int64

type t__rlim_t = int64

type t__sa_family_t = uint8

type t__socklen_t = uint32

type t__suseconds_t = int64

type t__timer_t = uintptr

type t__mqd_t = uintptr

type t__uid_t = uint32

type t__useconds_t = uint32

type t__cpuwhich_t = int32

type t__cpulevel_t = int32

type t__cpusetid_t = int32

type t__daddr_t = int64

type t__ct_rune_t = int32

type t__rune_t = int32

type t__wint_t = int32

type t__char16_t = uint16

type t__char32_t = uint32

type t__max_align_t = struct {
	F__max_align1 int64
	F__max_align2 float64
}

type t__dev_t = uint64

type t__fixpt_t = uint32

type t__mbstate_t = struct {
	F_mbstateL  [0]t__int64_t
	F__mbstate8 [128]int8
}

type t__rman_res_t = uint64

type t__va_list = uintptr

type t__gnuc_va_list = uintptr

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tuint8_t = uint8

type Tuint16_t = uint16

type Tuint32_t = uint32

type Tuint64_t = uint64

type Tintptr_t = int64

type Tuintptr_t = uint64

type Tintmax_t = int64

type Tuintmax_t = uint64

type Tint_least8_t = int8

type Tint_least16_t = int16

type Tint_least32_t = int32

type Tint_least64_t = int64

type Tuint_least8_t = uint8

type Tuint_least16_t = uint16

type Tuint_least32_t = uint32

type Tuint_least64_t = uint64

type Tint_fast8_t = int32

type Tint_fast16_t = int32

type Tint_fast32_t = int32

type Tint_fast64_t = int64

type Tuint_fast8_t = uint32

type Tuint_fast16_t = uint32

type Tuint_fast32_t = uint32

type Tuint_fast64_t = uint64

type Tpthread_once = struct {
	Fstate int32
	Fmutex Tpthread_mutex_t
}

type Tpthread_t = uintptr

type Tpthread_attr_t = uintptr

type Tpthread_mutex_t = uintptr

type Tpthread_mutexattr_t = uintptr

type Tpthread_cond_t = uintptr

type Tpthread_condattr_t = uintptr

type Tpthread_key_t = int32

type Tpthread_once_t = struct {
	Fstate int32
	Fmutex Tpthread_mutex_t
}

type Tpthread_rwlock_t = uintptr

type Tpthread_rwlockattr_t = uintptr

type Tpthread_barrier_t = uintptr

type Tpthread_barrierattr_t = uintptr

type Tpthread_spinlock_t = uintptr

type Tpthread_addr_t = uintptr

type Tpthread_startroutine_t = uintptr

type Tu_char = uint8

type Tu_short = uint16

type Tu_int = uint32

type Tu_long = uint64

type Tushort = uint16

type Tuint = uint32

type Tu_int8_t = uint8

type Tu_int16_t = uint16

type Tu_int32_t = uint32

type Tu_int64_t = uint64

type Tu_quad_t = uint64

type Tquad_t = int64

type Tqaddr_t = uintptr

type Tcaddr_t = uintptr

type Tc_caddr_t = uintptr

type Tblksize_t = int32

type Tcpuwhich_t = int32

type Tcpulevel_t = int32

type Tcpusetid_t = int32

type Tblkcnt_t = int64

type Tclock_t = int32

type Tclockid_t = int32

type Tcritical_t = int64

type Tdaddr_t = int64

type Tdev_t = uint64

type Tfflags_t = uint32

type Tfixpt_t = uint32

type Tfsblkcnt_t = uint64

type Tfsfilcnt_t = uint64

type Tgid_t = uint32

type Tin_addr_t = uint32

type Tin_port_t = uint16

type Tid_t = int64

type Tino_t = uint64

type Tkey_t = int64

type Tlwpid_t = int32

type Tmode_t = uint16

type Taccmode_t = int32

type Tnlink_t = uint64

type Toff_t = int64

type Toff64_t = int64

type Tpid_t = int32

type Tregister_t = int64

type Trlim_t = int64

type Tsbintime_t = int64

type Tsegsz_t = int64

type Tsize_t = uint64

type Tssize_t = int64

type Tsuseconds_t = int64

type Ttime_t = int64

type Ttimer_t = uintptr

type Tmqd_t = uintptr

type Tu_register_t = uint64

type Tuid_t = uint32

type Tuseconds_t = uint32

type Tcap_ioctl_t = uint64

type Tkpaddr_t = uint64

type Tkvaddr_t = uint64

type Tksize_t = uint64

type Tkssize_t = int64

type Tvm_offset_t = uint64

type Tvm_ooffset_t = uint64

type Tvm_paddr_t = uint64

type Tvm_pindex_t = uint64

type Tvm_size_t = uint64

type Trman_res_t = uint64

type Tsyscallarg_t = int64

type t__sigset_t = struct {
	F__bits [4]t__uint32_t
}

type t__sigset = t__sigset_t

type Ttimeval = struct {
	Ftv_sec  Ttime_t
	Ftv_usec Tsuseconds_t
}

type Ttimespec = struct {
	Ftv_sec  Ttime_t
	Ftv_nsec int64
}

type Titimerspec = struct {
	Fit_interval Ttimespec
	Fit_value    Ttimespec
}

type t__fd_mask = uint64

type Tfd_mask = uint64

type Tsigset_t = struct {
	F__bits [4]t__uint32_t
}

type Tfd_set = struct {
	F__fds_bits [16]t__fd_mask
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

type Tptrdiff_t = int64

type Trune_t = int32

type Twchar_t = int32

type Tmax_align_t = struct {
	F__max_align1 int64
	F__max_align2 float64
}

type Trsize_t = uint64

type TXPointer = uintptr

type TXExtData = struct {
	Fnumber       int32
	Fnext         uintptr
	Ffree_private uintptr
	Fprivate_data TXPointer
}

type T_XExtData = TXExtData

type TXExtCodes = struct {
	Fextension    int32
	Fmajor_opcode int32
	Ffirst_event  int32
	Ffirst_error  int32
}

type TXPixmapFormatValues = struct {
	Fdepth          int32
	Fbits_per_pixel int32
	Fscanline_pad   int32
}

type TXGCValues = struct {
	Ffunction           int32
	Fplane_mask         uint64
	Fforeground         uint64
	Fbackground         uint64
	Fline_width         int32
	Fline_style         int32
	Fcap_style          int32
	Fjoin_style         int32
	Ffill_style         int32
	Ffill_rule          int32
	Farc_mode           int32
	Ftile               TPixmap
	Fstipple            TPixmap
	Fts_x_origin        int32
	Fts_y_origin        int32
	Ffont               TFont
	Fsubwindow_mode     int32
	Fgraphics_exposures int32
	Fclip_x_origin      int32
	Fclip_y_origin      int32
	Fclip_mask          TPixmap
	Fdash_offset        int32
	Fdashes             int8
}

type TGC = uintptr

type T_XGC = struct {
	Fext_data uintptr
	Fgid      TGContext
	Frects    int32
	Fdashes   int32
	Fdirty    uint64
	Fvalues   TXGCValues
}

type TVisual = struct {
	Fext_data     uintptr
	Fvisualid     TVisualID
	Fclass        int32
	Fred_mask     uint64
	Fgreen_mask   uint64
	Fblue_mask    uint64
	Fbits_per_rgb int32
	Fmap_entries  int32
}

type TDepth = struct {
	Fdepth    int32
	Fnvisuals int32
	Fvisuals  uintptr
}

type T_XDisplay = struct {
	Fext_data            uintptr
	Ffree_funcs          uintptr
	Ffd                  int32
	Fconn_checker        int32
	Fproto_major_version int32
	Fproto_minor_version int32
	Fvendor              uintptr
	Fresource_base       TXID
	Fresource_mask       TXID
	Fresource_id         TXID
	Fresource_shift      int32
	Fresource_alloc      uintptr
	Fbyte_order          int32
	Fbitmap_unit         int32
	Fbitmap_pad          int32
	Fbitmap_bit_order    int32
	Fnformats            int32
	Fpixmap_format       uintptr
	Fvnumber             int32
	Frelease             int32
	Fhead                uintptr
	Ftail                uintptr
	Fqlen                int32
	Flast_request_read   uint64
	Frequest             uint64
	Flast_req            uintptr
	Fbuffer              uintptr
	Fbufptr              uintptr
	Fbufmax              uintptr
	Fmax_request_size    uint32
	Fdb                  uintptr
	Fsynchandler         uintptr
	Fdisplay_name        uintptr
	Fdefault_screen      int32
	Fnscreens            int32
	Fscreens             uintptr
	Fmotion_buffer       uint64
	Fflags               uint64
	Fmin_keycode         int32
	Fmax_keycode         int32
	Fkeysyms             uintptr
	Fmodifiermap         uintptr
	Fkeysyms_per_keycode int32
	Fxdefaults           uintptr
	Fscratch_buffer      uintptr
	Fscratch_length      uint64
	Fext_number          int32
	Fext_procs           uintptr
	Fevent_vec           [128]uintptr
	Fwire_vec            [128]uintptr
	Flock_meaning        TKeySym
	Flock                uintptr
	Fasync_handlers      uintptr
	Fbigreq_size         uint64
	Flock_fns            uintptr
	Fidlist_alloc        uintptr
	Fkey_bindings        uintptr
	Fcursor_font         TFont
	Fatoms               uintptr
	Fmode_switch         uint32
	Fnum_lock            uint32
	Fcontext_db          uintptr
	Ferror_vec           uintptr
	Fcms                 struct {
		FdefaultCCCs            TXPointer
		FclientCmaps            TXPointer
		FperVisualIntensityMaps TXPointer
	}
	Fim_filters             uintptr
	Fqfree                  uintptr
	Fnext_event_serial_num  uint64
	Fflushes                uintptr
	Fim_fd_info             uintptr
	Fim_fd_length           int32
	Fconn_watchers          uintptr
	Fwatcher_count          int32
	Ffiledes                TXPointer
	Fsavedsynchandler       uintptr
	Fresource_max           TXID
	Fxcmisc_opcode          int32
	Fxkb_info               uintptr
	Ftrans_conn             uintptr
	Fxcb                    uintptr
	Fnext_cookie            uint32
	Fgeneric_event_vec      [128]uintptr
	Fgeneric_event_copy_vec [128]uintptr
	Fcookiejar              uintptr
	Ferror_threads          uintptr
	Fexit_handler           TXIOErrorExitHandler
	Fexit_handler_data      uintptr
	Fin_ifevent             TCARD32
	Fifevent_thread         Txthread_t
}

type TScreen = struct {
	Fext_data        uintptr
	Fdisplay         uintptr
	Froot            TWindow
	Fwidth           int32
	Fheight          int32
	Fmwidth          int32
	Fmheight         int32
	Fndepths         int32
	Fdepths          uintptr
	Froot_depth      int32
	Froot_visual     uintptr
	Fdefault_gc      TGC
	Fcmap            TColormap
	Fwhite_pixel     uint64
	Fblack_pixel     uint64
	Fmax_maps        int32
	Fmin_maps        int32
	Fbacking_store   int32
	Fsave_unders     int32
	Froot_input_mask int64
}

type TScreenFormat = struct {
	Fext_data       uintptr
	Fdepth          int32
	Fbits_per_pixel int32
	Fscanline_pad   int32
}

type TXSetWindowAttributes = struct {
	Fbackground_pixmap     TPixmap
	Fbackground_pixel      uint64
	Fborder_pixmap         TPixmap
	Fborder_pixel          uint64
	Fbit_gravity           int32
	Fwin_gravity           int32
	Fbacking_store         int32
	Fbacking_planes        uint64
	Fbacking_pixel         uint64
	Fsave_under            int32
	Fevent_mask            int64
	Fdo_not_propagate_mask int64
	Foverride_redirect     int32
	Fcolormap              TColormap
	Fcursor                TCursor
}

type TXWindowAttributes = struct {
	Fx                     int32
	Fy                     int32
	Fwidth                 int32
	Fheight                int32
	Fborder_width          int32
	Fdepth                 int32
	Fvisual                uintptr
	Froot                  TWindow
	Fclass                 int32
	Fbit_gravity           int32
	Fwin_gravity           int32
	Fbacking_store         int32
	Fbacking_planes        uint64
	Fbacking_pixel         uint64
	Fsave_under            int32
	Fcolormap              TColormap
	Fmap_installed         int32
	Fmap_state             int32
	Fall_event_masks       int64
	Fyour_event_mask       int64
	Fdo_not_propagate_mask int64
	Foverride_redirect     int32
	Fscreen                uintptr
}

type TXHostAddress = struct {
	Ffamily  int32
	Flength  int32
	Faddress uintptr
}

type TXServerInterpretedAddress = struct {
	Ftypelength  int32
	Fvaluelength int32
	Ftype1       uintptr
	Fvalue       uintptr
}

type TXImage = struct {
	Fwidth            int32
	Fheight           int32
	Fxoffset          int32
	Fformat           int32
	Fdata             uintptr
	Fbyte_order       int32
	Fbitmap_unit      int32
	Fbitmap_bit_order int32
	Fbitmap_pad       int32
	Fdepth            int32
	Fbytes_per_line   int32
	Fbits_per_pixel   int32
	Fred_mask         uint64
	Fgreen_mask       uint64
	Fblue_mask        uint64
	Fobdata           TXPointer
	Ff                Tfuncs
}

type T_XImage = TXImage

type TXWindowChanges = struct {
	Fx            int32
	Fy            int32
	Fwidth        int32
	Fheight       int32
	Fborder_width int32
	Fsibling      TWindow
	Fstack_mode   int32
}

type TXColor = struct {
	Fpixel uint64
	Fred   uint16
	Fgreen uint16
	Fblue  uint16
	Fflags int8
	Fpad   int8
}

type TXSegment = struct {
	Fx1 int16
	Fy1 int16
	Fx2 int16
	Fy2 int16
}

type TXPoint = struct {
	Fx int16
	Fy int16
}

type TXRectangle = struct {
	Fx      int16
	Fy      int16
	Fwidth  uint16
	Fheight uint16
}

type TXArc = struct {
	Fx      int16
	Fy      int16
	Fwidth  uint16
	Fheight uint16
	Fangle1 int16
	Fangle2 int16
}

type TXKeyboardControl = struct {
	Fkey_click_percent int32
	Fbell_percent      int32
	Fbell_pitch        int32
	Fbell_duration     int32
	Fled               int32
	Fled_mode          int32
	Fkey               int32
	Fauto_repeat_mode  int32
}

type TXKeyboardState = struct {
	Fkey_click_percent  int32
	Fbell_percent       int32
	Fbell_pitch         uint32
	Fbell_duration      uint32
	Fled_mask           uint64
	Fglobal_auto_repeat int32
	Fauto_repeats       [32]int8
}

type TXTimeCoord = struct {
	Ftime TTime
	Fx    int16
	Fy    int16
}

type TXModifierKeymap = struct {
	Fmax_keypermod int32
	Fmodifiermap   uintptr
}

type TDisplay = struct {
	Fext_data            uintptr
	Ffree_funcs          uintptr
	Ffd                  int32
	Fconn_checker        int32
	Fproto_major_version int32
	Fproto_minor_version int32
	Fvendor              uintptr
	Fresource_base       TXID
	Fresource_mask       TXID
	Fresource_id         TXID
	Fresource_shift      int32
	Fresource_alloc      uintptr
	Fbyte_order          int32
	Fbitmap_unit         int32
	Fbitmap_pad          int32
	Fbitmap_bit_order    int32
	Fnformats            int32
	Fpixmap_format       uintptr
	Fvnumber             int32
	Frelease             int32
	Fhead                uintptr
	Ftail                uintptr
	Fqlen                int32
	Flast_request_read   uint64
	Frequest             uint64
	Flast_req            uintptr
	Fbuffer              uintptr
	Fbufptr              uintptr
	Fbufmax              uintptr
	Fmax_request_size    uint32
	Fdb                  uintptr
	Fsynchandler         uintptr
	Fdisplay_name        uintptr
	Fdefault_screen      int32
	Fnscreens            int32
	Fscreens             uintptr
	Fmotion_buffer       uint64
	Fflags               uint64
	Fmin_keycode         int32
	Fmax_keycode         int32
	Fkeysyms             uintptr
	Fmodifiermap         uintptr
	Fkeysyms_per_keycode int32
	Fxdefaults           uintptr
	Fscratch_buffer      uintptr
	Fscratch_length      uint64
	Fext_number          int32
	Fext_procs           uintptr
	Fevent_vec           [128]uintptr
	Fwire_vec            [128]uintptr
	Flock_meaning        TKeySym
	Flock                uintptr
	Fasync_handlers      uintptr
	Fbigreq_size         uint64
	Flock_fns            uintptr
	Fidlist_alloc        uintptr
	Fkey_bindings        uintptr
	Fcursor_font         TFont
	Fatoms               uintptr
	Fmode_switch         uint32
	Fnum_lock            uint32
	Fcontext_db          uintptr
	Ferror_vec           uintptr
	Fcms                 struct {
		FdefaultCCCs            TXPointer
		FclientCmaps            TXPointer
		FperVisualIntensityMaps TXPointer
	}
	Fim_filters             uintptr
	Fqfree                  uintptr
	Fnext_event_serial_num  uint64
	Fflushes                uintptr
	Fim_fd_info             uintptr
	Fim_fd_length           int32
	Fconn_watchers          uintptr
	Fwatcher_count          int32
	Ffiledes                TXPointer
	Fsavedsynchandler       uintptr
	Fresource_max           TXID
	Fxcmisc_opcode          int32
	Fxkb_info               uintptr
	Ftrans_conn             uintptr
	Fxcb                    uintptr
	Fnext_cookie            uint32
	Fgeneric_event_vec      [128]uintptr
	Fgeneric_event_copy_vec [128]uintptr
	Fcookiejar              uintptr
	Ferror_threads          uintptr
	Fexit_handler           TXIOErrorExitHandler
	Fexit_handler_data      uintptr
	Fin_ifevent             TCARD32
	Fifevent_thread         Txthread_t
}

type T_XPrivDisplay = uintptr

type TXKeyEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fkeycode     uint32
	Fsame_screen int32
}

type TXKeyPressedEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fkeycode     uint32
	Fsame_screen int32
}

type TXKeyReleasedEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fkeycode     uint32
	Fsame_screen int32
}

type TXButtonEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fbutton      uint32
	Fsame_screen int32
}

type TXButtonPressedEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fbutton      uint32
	Fsame_screen int32
}

type TXButtonReleasedEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fbutton      uint32
	Fsame_screen int32
}

type TXMotionEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fis_hint     int8
	Fsame_screen int32
}

type TXPointerMovedEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fstate       uint32
	Fis_hint     int8
	Fsame_screen int32
}

type TXCrossingEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fmode        int32
	Fdetail      int32
	Fsame_screen int32
	Ffocus       int32
	Fstate       uint32
}

type TXEnterWindowEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fmode        int32
	Fdetail      int32
	Fsame_screen int32
	Ffocus       int32
	Fstate       uint32
}

type TXLeaveWindowEvent = struct {
	Ftype1       int32
	Fserial      uint64
	Fsend_event  int32
	Fdisplay     uintptr
	Fwindow      TWindow
	Froot        TWindow
	Fsubwindow   TWindow
	Ftime        TTime
	Fx           int32
	Fy           int32
	Fx_root      int32
	Fy_root      int32
	Fmode        int32
	Fdetail      int32
	Fsame_screen int32
	Ffocus       int32
	Fstate       uint32
}

type TXFocusChangeEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fmode       int32
	Fdetail     int32
}

type TXFocusInEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fmode       int32
	Fdetail     int32
}

type TXFocusOutEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fmode       int32
	Fdetail     int32
}

type TXKeymapEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fkey_vector [32]int8
}

type TXExposeEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fx          int32
	Fy          int32
	Fwidth      int32
	Fheight     int32
	Fcount      int32
}

type TXGraphicsExposeEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fdrawable   TDrawable
	Fx          int32
	Fy          int32
	Fwidth      int32
	Fheight     int32
	Fcount      int32
	Fmajor_code int32
	Fminor_code int32
}

type TXNoExposeEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fdrawable   TDrawable
	Fmajor_code int32
	Fminor_code int32
}

type TXVisibilityEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fstate      int32
}

type TXCreateWindowEvent = struct {
	Ftype1             int32
	Fserial            uint64
	Fsend_event        int32
	Fdisplay           uintptr
	Fparent            TWindow
	Fwindow            TWindow
	Fx                 int32
	Fy                 int32
	Fwidth             int32
	Fheight            int32
	Fborder_width      int32
	Foverride_redirect int32
}

type TXDestroyWindowEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fevent      TWindow
	Fwindow     TWindow
}

type TXUnmapEvent = struct {
	Ftype1          int32
	Fserial         uint64
	Fsend_event     int32
	Fdisplay        uintptr
	Fevent          TWindow
	Fwindow         TWindow
	Ffrom_configure int32
}

type TXMapEvent = struct {
	Ftype1             int32
	Fserial            uint64
	Fsend_event        int32
	Fdisplay           uintptr
	Fevent             TWindow
	Fwindow            TWindow
	Foverride_redirect int32
}

type TXMapRequestEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fparent     TWindow
	Fwindow     TWindow
}

type TXReparentEvent = struct {
	Ftype1             int32
	Fserial            uint64
	Fsend_event        int32
	Fdisplay           uintptr
	Fevent             TWindow
	Fwindow            TWindow
	Fparent            TWindow
	Fx                 int32
	Fy                 int32
	Foverride_redirect int32
}

type TXConfigureEvent = struct {
	Ftype1             int32
	Fserial            uint64
	Fsend_event        int32
	Fdisplay           uintptr
	Fevent             TWindow
	Fwindow            TWindow
	Fx                 int32
	Fy                 int32
	Fwidth             int32
	Fheight            int32
	Fborder_width      int32
	Fabove             TWindow
	Foverride_redirect int32
}

type TXGravityEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fevent      TWindow
	Fwindow     TWindow
	Fx          int32
	Fy          int32
}

type TXResizeRequestEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fwidth      int32
	Fheight     int32
}

type TXConfigureRequestEvent = struct {
	Ftype1        int32
	Fserial       uint64
	Fsend_event   int32
	Fdisplay      uintptr
	Fparent       TWindow
	Fwindow       TWindow
	Fx            int32
	Fy            int32
	Fwidth        int32
	Fheight       int32
	Fborder_width int32
	Fabove        TWindow
	Fdetail       int32
	Fvalue_mask   uint64
}

type TXCirculateEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fevent      TWindow
	Fwindow     TWindow
	Fplace      int32
}

type TXCirculateRequestEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fparent     TWindow
	Fwindow     TWindow
	Fplace      int32
}

type TXPropertyEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fatom       TAtom
	Ftime       TTime
	Fstate      int32
}

type TXSelectionClearEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fselection  TAtom
	Ftime       TTime
}

type TXSelectionRequestEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fowner      TWindow
	Frequestor  TWindow
	Fselection  TAtom
	Ftarget     TAtom
	Fproperty   TAtom
	Ftime       TTime
}

type TXSelectionEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Frequestor  TWindow
	Fselection  TAtom
	Ftarget     TAtom
	Fproperty   TAtom
	Ftime       TTime
}

type TXColormapEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
	Fcolormap   TColormap
	Fnew1       int32
	Fstate      int32
}

type TXClientMessageEvent = struct {
	Ftype1        int32
	Fserial       uint64
	Fsend_event   int32
	Fdisplay      uintptr
	Fwindow       TWindow
	Fmessage_type TAtom
	Fformat       int32
	Fdata         struct {
		Fs           [0][10]int16
		Fl           [0][5]int64
		Fb           [20]int8
		F__ccgo_pad3 [20]byte
	}
}

type TXMappingEvent = struct {
	Ftype1         int32
	Fserial        uint64
	Fsend_event    int32
	Fdisplay       uintptr
	Fwindow        TWindow
	Frequest       int32
	Ffirst_keycode int32
	Fcount         int32
}

type TXErrorEvent = struct {
	Ftype1        int32
	Fdisplay      uintptr
	Fresourceid   TXID
	Fserial       uint64
	Ferror_code   uint8
	Frequest_code uint8
	Fminor_code   uint8
}

type TXAnyEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fwindow     TWindow
}

type TXGenericEvent = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fextension  int32
	Fevtype     int32
}

type TXGenericEventCookie = struct {
	Ftype1      int32
	Fserial     uint64
	Fsend_event int32
	Fdisplay    uintptr
	Fextension  int32
	Fevtype     int32
	Fcookie     uint32
	Fdata       uintptr
}

type TXEvent = struct {
	Fxany              [0]TXAnyEvent
	Fxkey              [0]TXKeyEvent
	Fxbutton           [0]TXButtonEvent
	Fxmotion           [0]TXMotionEvent
	Fxcrossing         [0]TXCrossingEvent
	Fxfocus            [0]TXFocusChangeEvent
	Fxexpose           [0]TXExposeEvent
	Fxgraphicsexpose   [0]TXGraphicsExposeEvent
	Fxnoexpose         [0]TXNoExposeEvent
	Fxvisibility       [0]TXVisibilityEvent
	Fxcreatewindow     [0]TXCreateWindowEvent
	Fxdestroywindow    [0]TXDestroyWindowEvent
	Fxunmap            [0]TXUnmapEvent
	Fxmap              [0]TXMapEvent
	Fxmaprequest       [0]TXMapRequestEvent
	Fxreparent         [0]TXReparentEvent
	Fxconfigure        [0]TXConfigureEvent
	Fxgravity          [0]TXGravityEvent
	Fxresizerequest    [0]TXResizeRequestEvent
	Fxconfigurerequest [0]TXConfigureRequestEvent
	Fxcirculate        [0]TXCirculateEvent
	Fxcirculaterequest [0]TXCirculateRequestEvent
	Fxproperty         [0]TXPropertyEvent
	Fxselectionclear   [0]TXSelectionClearEvent
	Fxselectionrequest [0]TXSelectionRequestEvent
	Fxselection        [0]TXSelectionEvent
	Fxcolormap         [0]TXColormapEvent
	Fxclient           [0]TXClientMessageEvent
	Fxmapping          [0]TXMappingEvent
	Fxerror            [0]TXErrorEvent
	Fxkeymap           [0]TXKeymapEvent
	Fxgeneric          [0]TXGenericEvent
	Fxcookie           [0]TXGenericEventCookie
	Fpad               [0][24]int64
	Ftype1             int32
	F__ccgo_pad35      [188]byte
}

type T_XEvent = TXEvent

type TXCharStruct = struct {
	Flbearing   int16
	Frbearing   int16
	Fwidth      int16
	Fascent     int16
	Fdescent    int16
	Fattributes uint16
}

type TXFontProp = struct {
	Fname   TAtom
	Fcard32 uint64
}

type TXFontStruct = struct {
	Fext_data          uintptr
	Ffid               TFont
	Fdirection         uint32
	Fmin_char_or_byte2 uint32
	Fmax_char_or_byte2 uint32
	Fmin_byte1         uint32
	Fmax_byte1         uint32
	Fall_chars_exist   int32
	Fdefault_char      uint32
	Fn_properties      int32
	Fproperties        uintptr
	Fmin_bounds        TXCharStruct
	Fmax_bounds        TXCharStruct
	Fper_char          uintptr
	Fascent            int32
	Fdescent           int32
}

type TXTextItem = struct {
	Fchars  uintptr
	Fnchars int32
	Fdelta  int32
	Ffont   TFont
}

type TXChar2b = struct {
	Fbyte1 uint8
	Fbyte2 uint8
}

type TXTextItem16 = struct {
	Fchars  uintptr
	Fnchars int32
	Fdelta  int32
	Ffont   TFont
}

type TXEDataObject = struct {
	Fgc            [0]TGC
	Fvisual        [0]uintptr
	Fscreen        [0]uintptr
	Fpixmap_format [0]uintptr
	Ffont          [0]uintptr
	Fdisplay       uintptr
}

type TXFontSetExtents = struct {
	Fmax_ink_extent     TXRectangle
	Fmax_logical_extent TXRectangle
}

type TXOM = uintptr

type TXOC = uintptr

type TXFontSet = uintptr

type TXmbTextItem = struct {
	Fchars    uintptr
	Fnchars   int32
	Fdelta    int32
	Ffont_set TXFontSet
}

type TXwcTextItem = struct {
	Fchars    uintptr
	Fnchars   int32
	Fdelta    int32
	Ffont_set TXFontSet
}

type TXOMCharSetList = struct {
	Fcharset_count int32
	Fcharset_list  uintptr
}

type TXOrientation = int32

const _XOMOrientation_LTR_TTB = 0
const _XOMOrientation_RTL_TTB = 1
const _XOMOrientation_TTB_LTR = 2
const _XOMOrientation_TTB_RTL = 3
const _XOMOrientation_Context = 4

type TXOMOrientation = struct {
	Fnum_orientation int32
	Forientation     uintptr
}

type TXOMFontInfo = struct {
	Fnum_font         int32
	Ffont_struct_list uintptr
	Ffont_name_list   uintptr
}

type TXIM = uintptr

type TXIC = uintptr

type TXIMProc = uintptr

type TXICProc = uintptr

type TXIDProc = uintptr

type TXIMStyle = uint64

type TXIMStyles = struct {
	Fcount_styles     uint16
	Fsupported_styles uintptr
}

type TXVaNestedList = uintptr

type TXIMCallback = struct {
	Fclient_data TXPointer
	Fcallback    TXIMProc
}

type TXICCallback = struct {
	Fclient_data TXPointer
	Fcallback    TXICProc
}

type TXIMFeedback = uint64

type TXIMText = struct {
	Flength            uint16
	Ffeedback          uintptr
	Fencoding_is_wchar int32
	Fstring1           struct {
		Fwide_char  [0]uintptr
		Fmulti_byte uintptr
	}
}

type T_XIMText = TXIMText

type TXIMPreeditState = uint64

type TXIMPreeditStateNotifyCallbackStruct = struct {
	Fstate TXIMPreeditState
}

type T_XIMPreeditStateNotifyCallbackStruct = TXIMPreeditStateNotifyCallbackStruct

type TXIMResetState = uint64

type TXIMStringConversionFeedback = uint64

type TXIMStringConversionText = struct {
	Flength            uint16
	Ffeedback          uintptr
	Fencoding_is_wchar int32
	Fstring1           struct {
		Fwcs [0]uintptr
		Fmbs uintptr
	}
}

type T_XIMStringConversionText = TXIMStringConversionText

type TXIMStringConversionPosition = uint16

type TXIMStringConversionType = uint16

type TXIMStringConversionOperation = uint16

type TXIMCaretDirection = int32

const _XIMForwardChar = 0
const _XIMBackwardChar = 1
const _XIMForwardWord = 2
const _XIMBackwardWord = 3
const _XIMCaretUp = 4
const _XIMCaretDown = 5
const _XIMNextLine = 6
const _XIMPreviousLine = 7
const _XIMLineStart = 8
const _XIMLineEnd = 9
const _XIMAbsolutePosition = 10
const _XIMDontChange = 11

type TXIMStringConversionCallbackStruct = struct {
	Fposition  TXIMStringConversionPosition
	Fdirection TXIMCaretDirection
	Foperation TXIMStringConversionOperation
	Ffactor    uint16
	Ftext      uintptr
}

type T_XIMStringConversionCallbackStruct = TXIMStringConversionCallbackStruct

type TXIMPreeditDrawCallbackStruct = struct {
	Fcaret      int32
	Fchg_first  int32
	Fchg_length int32
	Ftext       uintptr
}

type T_XIMPreeditDrawCallbackStruct = TXIMPreeditDrawCallbackStruct

type TXIMCaretStyle = int32

const _XIMIsInvisible = 0
const _XIMIsPrimary = 1
const _XIMIsSecondary = 2

type TXIMPreeditCaretCallbackStruct = struct {
	Fposition  int32
	Fdirection TXIMCaretDirection
	Fstyle     TXIMCaretStyle
}

type T_XIMPreeditCaretCallbackStruct = TXIMPreeditCaretCallbackStruct

type TXIMStatusDataType = int32

const _XIMTextType = 0
const _XIMBitmapType = 1

type TXIMStatusDrawCallbackStruct = struct {
	Ftype1 TXIMStatusDataType
	Fdata  struct {
		Fbitmap [0]TPixmap
		Ftext   uintptr
	}
}

type T_XIMStatusDrawCallbackStruct = TXIMStatusDrawCallbackStruct

type TXIMHotKeyTrigger = struct {
	Fkeysym        TKeySym
	Fmodifier      int32
	Fmodifier_mask int32
}

type T_XIMHotKeyTrigger = TXIMHotKeyTrigger

type TXIMHotKeyTriggers = struct {
	Fnum_hot_key int32
	Fkey         uintptr
}

type T_XIMHotKeyTriggers = TXIMHotKeyTriggers

type TXIMHotKeyState = uint64

type TXIMValuesList = struct {
	Fcount_values     uint16
	Fsupported_values uintptr
}

type TXErrorHandler = uintptr

type TXIOErrorHandler = uintptr

type TXIOErrorExitHandler = uintptr

type TXConnectionWatchProc = uintptr

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

type TxSegment = struct {
	Fx1 TINT16
	Fy1 TINT16
	Fx2 TINT16
	Fy2 TINT16
}

type T_xSegment = TxSegment

type TxPoint = struct {
	Fx TINT16
	Fy TINT16
}

type T_xPoint = TxPoint

type TxRectangle = struct {
	Fx      TINT16
	Fy      TINT16
	Fwidth  TCARD16
	Fheight TCARD16
}

type T_xRectangle = TxRectangle

type TxArc = struct {
	Fx      TINT16
	Fy      TINT16
	Fwidth  TCARD16
	Fheight TCARD16
	Fangle1 TINT16
	Fangle2 TINT16
}

type T_xArc = TxArc

type TKeyButMask = uint16

type TxConnClientPrefix = struct {
	FbyteOrder        TCARD8
	Fpad              TBYTE
	FmajorVersion     TCARD16
	FminorVersion     TCARD16
	FnbytesAuthProto  TCARD16
	FnbytesAuthString TCARD16
	Fpad2             TCARD16
}

type TxConnSetupPrefix = struct {
	Fsuccess      TCARD8
	FlengthReason TBYTE
	FmajorVersion TCARD16
	FminorVersion TCARD16
	Flength       TCARD16
}

type TxConnSetup = struct {
	Frelease            TCARD32
	FridBase            TCARD32
	FridMask            TCARD32
	FmotionBufferSize   TCARD32
	FnbytesVendor       TCARD16
	FmaxRequestSize     TCARD16
	FnumRoots           TCARD8
	FnumFormats         TCARD8
	FimageByteOrder     TCARD8
	FbitmapBitOrder     TCARD8
	FbitmapScanlineUnit TCARD8
	FbitmapScanlinePad  TCARD8
	FminKeyCode         TCARD8
	FmaxKeyCode         TCARD8
	Fpad2               TCARD32
}

type TxPixmapFormat = struct {
	Fdepth        TCARD8
	FbitsPerPixel TCARD8
	FscanLinePad  TCARD8
	Fpad1         TCARD8
	Fpad2         TCARD32
}

type TxDepth = struct {
	Fdepth    TCARD8
	Fpad1     TCARD8
	FnVisuals TCARD16
	Fpad2     TCARD32
}

type TxVisualType = struct {
	FvisualID        TCARD32
	Fclass           TCARD8
	FbitsPerRGB      TCARD8
	FcolormapEntries TCARD16
	FredMask         TCARD32
	FgreenMask       TCARD32
	FblueMask        TCARD32
	Fpad             TCARD32
}

type TxWindowRoot = struct {
	FwindowId         TCARD32
	FdefaultColormap  TCARD32
	FwhitePixel       TCARD32
	FblackPixel       TCARD32
	FcurrentInputMask TCARD32
	FpixWidth         TCARD16
	FpixHeight        TCARD16
	FmmWidth          TCARD16
	FmmHeight         TCARD16
	FminInstalledMaps TCARD16
	FmaxInstalledMaps TCARD16
	FrootVisualID     TCARD32
	FbackingStore     TCARD8
	FsaveUnders       TBOOL
	FrootDepth        TCARD8
	FnDepths          TCARD8
}

type TxTimecoord = struct {
	Ftime TCARD32
	Fx    TINT16
	Fy    TINT16
}

type TxHostEntry = struct {
	Ffamily TCARD8
	Fpad    TBYTE
	Flength TCARD16
}

type TxCharInfo = struct {
	FleftSideBearing  TINT16
	FrightSideBearing TINT16
	FcharacterWidth   TINT16
	Fascent           TINT16
	Fdescent          TINT16
	Fattributes       TCARD16
}

type TxFontProp = struct {
	Fname  TCARD32
	Fvalue TCARD32
}

type TxTextElt = struct {
	Flen1  TCARD8
	Fdelta TINT8
}

type TxColorItem = struct {
	Fpixel TCARD32
	Fred   TCARD16
	Fgreen TCARD16
	Fblue  TCARD16
	Fflags TCARD8
	Fpad   TCARD8
}

type Txrgb = struct {
	Fred   TCARD16
	Fgreen TCARD16
	Fblue  TCARD16
	Fpad   TCARD16
}

type TKEYCODE = uint8

type TxGenericReply = struct {
	Ftype1          TBYTE
	Fdata1          TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fdata00         TCARD32
	Fdata01         TCARD32
	Fdata02         TCARD32
	Fdata03         TCARD32
	Fdata04         TCARD32
	Fdata05         TCARD32
}

type TxGetWindowAttributesReply = struct {
	Ftype1              TBYTE
	FbackingStore       TCARD8
	FsequenceNumber     TCARD16
	Flength             TCARD32
	FvisualID           TCARD32
	Fclass              TCARD16
	FbitGravity         TCARD8
	FwinGravity         TCARD8
	FbackingBitPlanes   TCARD32
	FbackingPixel       TCARD32
	FsaveUnder          TBOOL
	FmapInstalled       TBOOL
	FmapState           TCARD8
	Foverride           TBOOL
	Fcolormap           TCARD32
	FallEventMasks      TCARD32
	FyourEventMask      TCARD32
	FdoNotPropagateMask TCARD16
	Fpad                TCARD16
}

type TxGetGeometryReply = struct {
	Ftype1          TBYTE
	Fdepth          TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Froot           TCARD32
	Fx              TINT16
	Fy              TINT16
	Fwidth          TCARD16
	Fheight         TCARD16
	FborderWidth    TCARD16
	Fpad1           TCARD16
	Fpad2           TCARD32
	Fpad3           TCARD32
}

type TxQueryTreeReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Froot           TCARD32
	Fparent         TCARD32
	FnChildren      TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxInternAtomReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fatom           TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxGetAtomNameReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnameLength     TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxGetPropertyReply = struct {
	Ftype1          TBYTE
	Fformat         TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	FpropertyType   TCARD32
	FbytesAfter     TCARD32
	FnItems         TCARD32
	Fpad1           TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
}

type TxListPropertiesReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnProperties    TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxGetSelectionOwnerReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fowner          TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxGrabPointerReply = struct {
	Ftype1          TBYTE
	Fstatus         TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad1           TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxGrabKeyboardReply = struct {
	Ftype1          TBYTE
	Fstatus         TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad1           TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxQueryPointerReply = struct {
	Ftype1          TBYTE
	FsameScreen     TBOOL
	FsequenceNumber TCARD16
	Flength         TCARD32
	Froot           TCARD32
	Fchild          TCARD32
	FrootX          TINT16
	FrootY          TINT16
	FwinX           TINT16
	FwinY           TINT16
	Fmask           TCARD16
	Fpad1           TCARD16
	Fpad            TCARD32
}

type TxGetMotionEventsReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnEvents        TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxTranslateCoordsReply = struct {
	Ftype1          TBYTE
	FsameScreen     TBOOL
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fchild          TCARD32
	FdstX           TINT16
	FdstY           TINT16
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxGetInputFocusReply = struct {
	Ftype1          TBYTE
	FrevertTo       TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Ffocus          TCARD32
	Fpad1           TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxQueryKeymapReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fmap1           [32]TBYTE
}

type TxQueryFontReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FminBounds      TxCharInfo
	Fwalign1        TCARD32
	FmaxBounds      TxCharInfo
	Fwalign2        TCARD32
	FminCharOrByte2 TCARD16
	FmaxCharOrByte2 TCARD16
	FdefaultChar    TCARD16
	FnFontProps     TCARD16
	FdrawDirection  TCARD8
	FminByte1       TCARD8
	FmaxByte1       TCARD8
	FallCharsExist  TBOOL
	FfontAscent     TINT16
	FfontDescent    TINT16
	FnCharInfos     TCARD32
}

type T_xQueryFontReply = TxQueryFontReply

type TxQueryTextExtentsReply = struct {
	Ftype1          TBYTE
	FdrawDirection  TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	FfontAscent     TINT16
	FfontDescent    TINT16
	FoverallAscent  TINT16
	FoverallDescent TINT16
	FoverallWidth   TINT32
	FoverallLeft    TINT32
	FoverallRight   TINT32
	Fpad            TCARD32
}

type TxListFontsReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnFonts         TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxListFontsWithInfoReply = struct {
	Ftype1          TBYTE
	FnameLength     TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	FminBounds      TxCharInfo
	Fwalign1        TCARD32
	FmaxBounds      TxCharInfo
	Fwalign2        TCARD32
	FminCharOrByte2 TCARD16
	FmaxCharOrByte2 TCARD16
	FdefaultChar    TCARD16
	FnFontProps     TCARD16
	FdrawDirection  TCARD8
	FminByte1       TCARD8
	FmaxByte1       TCARD8
	FallCharsExist  TBOOL
	FfontAscent     TINT16
	FfontDescent    TINT16
	FnReplies       TCARD32
}

type TxGetFontPathReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnPaths         TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxGetImageReply = struct {
	Ftype1          TBYTE
	Fdepth          TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fvisual         TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxListInstalledColormapsReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnColormaps     TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxAllocColorReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fred            TCARD16
	Fgreen          TCARD16
	Fblue           TCARD16
	Fpad2           TCARD16
	Fpixel          TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxAllocNamedColorReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpixel          TCARD32
	FexactRed       TCARD16
	FexactGreen     TCARD16
	FexactBlue      TCARD16
	FscreenRed      TCARD16
	FscreenGreen    TCARD16
	FscreenBlue     TCARD16
	Fpad2           TCARD32
	Fpad3           TCARD32
}

type TxAllocColorCellsReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnPixels        TCARD16
	FnMasks         TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxAllocColorPlanesReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnPixels        TCARD16
	Fpad2           TCARD16
	FredMask        TCARD32
	FgreenMask      TCARD32
	FblueMask       TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
}

type TxQueryColorsReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnColors        TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxLookupColorReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FexactRed       TCARD16
	FexactGreen     TCARD16
	FexactBlue      TCARD16
	FscreenRed      TCARD16
	FscreenGreen    TCARD16
	FscreenBlue     TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxQueryBestSizeReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fwidth          TCARD16
	Fheight         TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxQueryExtensionReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpresent        TBOOL
	Fmajor_opcode   TCARD8
	Ffirst_event    TCARD8
	Ffirst_error    TCARD8
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxListExtensionsReply = struct {
	Ftype1          TBYTE
	FnExtensions    TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxSetMappingReply = struct {
	Ftype1          TBYTE
	Fsuccess        TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxSetPointerMappingReply = struct {
	Ftype1          TBYTE
	Fsuccess        TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxSetModifierMappingReply = struct {
	Ftype1          TBYTE
	Fsuccess        TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxGetPointerMappingReply = struct {
	Ftype1          TBYTE
	FnElts          TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxGetKeyboardMappingReply = struct {
	Ftype1             TBYTE
	FkeySymsPerKeyCode TCARD8
	FsequenceNumber    TCARD16
	Flength            TCARD32
	Fpad2              TCARD32
	Fpad3              TCARD32
	Fpad4              TCARD32
	Fpad5              TCARD32
	Fpad6              TCARD32
	Fpad7              TCARD32
}

type TxGetModifierMappingReply = struct {
	Ftype1             TBYTE
	FnumKeyPerModifier TCARD8
	FsequenceNumber    TCARD16
	Flength            TCARD32
	Fpad1              TCARD32
	Fpad2              TCARD32
	Fpad3              TCARD32
	Fpad4              TCARD32
	Fpad5              TCARD32
	Fpad6              TCARD32
}

type TxGetKeyboardControlReply = struct {
	Ftype1            TBYTE
	FglobalAutoRepeat TBOOL
	FsequenceNumber   TCARD16
	Flength           TCARD32
	FledMask          TCARD32
	FkeyClickPercent  TCARD8
	FbellPercent      TCARD8
	FbellPitch        TCARD16
	FbellDuration     TCARD16
	Fpad              TCARD16
	Fmap1             [32]TBYTE
}

type TxGetPointerControlReply = struct {
	Ftype1            TBYTE
	Fpad1             TBYTE
	FsequenceNumber   TCARD16
	Flength           TCARD32
	FaccelNumerator   TCARD16
	FaccelDenominator TCARD16
	Fthreshold        TCARD16
	Fpad2             TCARD16
	Fpad3             TCARD32
	Fpad4             TCARD32
	Fpad5             TCARD32
	Fpad6             TCARD32
}

type TxGetScreenSaverReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	Ftimeout        TCARD16
	Finterval       TCARD16
	FpreferBlanking TBOOL
	FallowExposures TBOOL
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxListHostsReply = struct {
	Ftype1          TBYTE
	Fenabled        TBOOL
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnHosts         TCARD16
	Fpad1           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxError = struct {
	Ftype1          TBYTE
	FerrorCode      TBYTE
	FsequenceNumber TCARD16
	FresourceID     TCARD32
	FminorCode      TCARD16
	FmajorCode      TCARD8
	Fpad1           TBYTE
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxEvent = struct {
	Fu struct {
		FkeyButtonPointer [0]struct {
			Fpad00      TCARD32
			Ftime       TCARD32
			Froot       TCARD32
			Fevent      TCARD32
			Fchild      TCARD32
			FrootX      TINT16
			FrootY      TINT16
			FeventX     TINT16
			FeventY     TINT16
			Fstate      TKeyButMask
			FsameScreen TBOOL
			Fpad1       TBYTE
		}
		FenterLeave [0]struct {
			Fpad00  TCARD32
			Ftime   TCARD32
			Froot   TCARD32
			Fevent  TCARD32
			Fchild  TCARD32
			FrootX  TINT16
			FrootY  TINT16
			FeventX TINT16
			FeventY TINT16
			Fstate  TKeyButMask
			Fmode   TBYTE
			Fflags  TBYTE
		}
		Ffocus [0]struct {
			Fpad00  TCARD32
			Fwindow TCARD32
			Fmode   TBYTE
			Fpad1   TBYTE
			Fpad2   TBYTE
			Fpad3   TBYTE
		}
		Fexpose [0]struct {
			Fpad00  TCARD32
			Fwindow TCARD32
			Fx      TCARD16
			Fy      TCARD16
			Fwidth  TCARD16
			Fheight TCARD16
			Fcount  TCARD16
			Fpad2   TCARD16
		}
		FgraphicsExposure [0]struct {
			Fpad00      TCARD32
			Fdrawable   TCARD32
			Fx          TCARD16
			Fy          TCARD16
			Fwidth      TCARD16
			Fheight     TCARD16
			FminorEvent TCARD16
			Fcount      TCARD16
			FmajorEvent TBYTE
			Fpad1       TBYTE
			Fpad2       TBYTE
			Fpad3       TBYTE
		}
		FnoExposure [0]struct {
			Fpad00      TCARD32
			Fdrawable   TCARD32
			FminorEvent TCARD16
			FmajorEvent TBYTE
			Fbpad       TBYTE
		}
		Fvisibility [0]struct {
			Fpad00  TCARD32
			Fwindow TCARD32
			Fstate  TCARD8
			Fpad1   TBYTE
			Fpad2   TBYTE
			Fpad3   TBYTE
		}
		FcreateNotify [0]struct {
			Fpad00       TCARD32
			Fparent      TCARD32
			Fwindow      TCARD32
			Fx           TINT16
			Fy           TINT16
			Fwidth       TCARD16
			Fheight      TCARD16
			FborderWidth TCARD16
			Foverride    TBOOL
			Fbpad        TBYTE
		}
		FdestroyNotify [0]struct {
			Fpad00  TCARD32
			Fevent  TCARD32
			Fwindow TCARD32
		}
		FunmapNotify [0]struct {
			Fpad00         TCARD32
			Fevent         TCARD32
			Fwindow        TCARD32
			FfromConfigure TBOOL
			Fpad1          TBYTE
			Fpad2          TBYTE
			Fpad3          TBYTE
		}
		FmapNotify [0]struct {
			Fpad00    TCARD32
			Fevent    TCARD32
			Fwindow   TCARD32
			Foverride TBOOL
			Fpad1     TBYTE
			Fpad2     TBYTE
			Fpad3     TBYTE
		}
		FmapRequest [0]struct {
			Fpad00  TCARD32
			Fparent TCARD32
			Fwindow TCARD32
		}
		Freparent [0]struct {
			Fpad00    TCARD32
			Fevent    TCARD32
			Fwindow   TCARD32
			Fparent   TCARD32
			Fx        TINT16
			Fy        TINT16
			Foverride TBOOL
			Fpad1     TBYTE
			Fpad2     TBYTE
			Fpad3     TBYTE
		}
		FconfigureNotify [0]struct {
			Fpad00        TCARD32
			Fevent        TCARD32
			Fwindow       TCARD32
			FaboveSibling TCARD32
			Fx            TINT16
			Fy            TINT16
			Fwidth        TCARD16
			Fheight       TCARD16
			FborderWidth  TCARD16
			Foverride     TBOOL
			Fbpad         TBYTE
		}
		FconfigureRequest [0]struct {
			Fpad00       TCARD32
			Fparent      TCARD32
			Fwindow      TCARD32
			Fsibling     TCARD32
			Fx           TINT16
			Fy           TINT16
			Fwidth       TCARD16
			Fheight      TCARD16
			FborderWidth TCARD16
			FvalueMask   TCARD16
			Fpad1        TCARD32
		}
		Fgravity [0]struct {
			Fpad00  TCARD32
			Fevent  TCARD32
			Fwindow TCARD32
			Fx      TINT16
			Fy      TINT16
			Fpad1   TCARD32
			Fpad2   TCARD32
			Fpad3   TCARD32
			Fpad4   TCARD32
		}
		FresizeRequest [0]struct {
			Fpad00  TCARD32
			Fwindow TCARD32
			Fwidth  TCARD16
			Fheight TCARD16
		}
		Fcirculate [0]struct {
			Fpad00  TCARD32
			Fevent  TCARD32
			Fwindow TCARD32
			Fparent TCARD32
			Fplace  TBYTE
			Fpad1   TBYTE
			Fpad2   TBYTE
			Fpad3   TBYTE
		}
		Fproperty [0]struct {
			Fpad00  TCARD32
			Fwindow TCARD32
			Fatom   TCARD32
			Ftime   TCARD32
			Fstate  TBYTE
			Fpad1   TBYTE
			Fpad2   TCARD16
		}
		FselectionClear [0]struct {
			Fpad00  TCARD32
			Ftime   TCARD32
			Fwindow TCARD32
			Fatom   TCARD32
		}
		FselectionRequest [0]struct {
			Fpad00     TCARD32
			Ftime      TCARD32
			Fowner     TCARD32
			Frequestor TCARD32
			Fselection TCARD32
			Ftarget    TCARD32
			Fproperty  TCARD32
		}
		FselectionNotify [0]struct {
			Fpad00     TCARD32
			Ftime      TCARD32
			Frequestor TCARD32
			Fselection TCARD32
			Ftarget    TCARD32
			Fproperty  TCARD32
		}
		Fcolormap [0]struct {
			Fpad00    TCARD32
			Fwindow   TCARD32
			Fcolormap TCARD32
			Fnew1     TBOOL
			Fstate    TBYTE
			Fpad1     TBYTE
			Fpad2     TBYTE
		}
		FmappingNotify [0]struct {
			Fpad00        TCARD32
			Frequest      TCARD8
			FfirstKeyCode TCARD8
			Fcount        TCARD8
			Fpad1         TBYTE
		}
		FclientMessage [0]struct {
			Fpad00  TCARD32
			Fwindow TCARD32
			Fu      struct {
				Fs [0]struct {
					Ftype1   TCARD32
					Fshorts0 TINT16
					Fshorts1 TINT16
					Fshorts2 TINT16
					Fshorts3 TINT16
					Fshorts4 TINT16
					Fshorts5 TINT16
					Fshorts6 TINT16
					Fshorts7 TINT16
					Fshorts8 TINT16
					Fshorts9 TINT16
				}
				Fb [0]struct {
					Ftype1 TCARD32
					Fbytes [20]TINT8
				}
				Fl struct {
					Ftype1  TCARD32
					Flongs0 TINT32
					Flongs1 TINT32
					Flongs2 TINT32
					Flongs3 TINT32
					Flongs4 TINT32
				}
			}
		}
		Fu struct {
			Ftype1          TBYTE
			Fdetail         TBYTE
			FsequenceNumber TCARD16
		}
		F__ccgo_pad26 [28]byte
	}
}

type T_xEvent = TxEvent

type TxGenericEvent = struct {
	Ftype1          TBYTE
	Fextension      TCARD8
	FsequenceNumber TCARD16
	Flength         TCARD32
	Fevtype         TCARD16
	Fpad2           TCARD16
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
	Fpad7           TCARD32
}

type TxKeymapEvent = struct {
	Ftype1 TBYTE
	Fmap1  [31]TBYTE
}

type TxReply = struct {
	Fgeom               [0]TxGetGeometryReply
	Ftree               [0]TxQueryTreeReply
	Fatom               [0]TxInternAtomReply
	FatomName           [0]TxGetAtomNameReply
	Fproperty           [0]TxGetPropertyReply
	FlistProperties     [0]TxListPropertiesReply
	Fselection          [0]TxGetSelectionOwnerReply
	FgrabPointer        [0]TxGrabPointerReply
	FgrabKeyboard       [0]TxGrabKeyboardReply
	Fpointer            [0]TxQueryPointerReply
	FmotionEvents       [0]TxGetMotionEventsReply
	Fcoords             [0]TxTranslateCoordsReply
	FinputFocus         [0]TxGetInputFocusReply
	FtextExtents        [0]TxQueryTextExtentsReply
	Ffonts              [0]TxListFontsReply
	FfontPath           [0]TxGetFontPathReply
	Fimage              [0]TxGetImageReply
	Fcolormaps          [0]TxListInstalledColormapsReply
	FallocColor         [0]TxAllocColorReply
	FallocNamedColor    [0]TxAllocNamedColorReply
	FcolorCells         [0]TxAllocColorCellsReply
	FcolorPlanes        [0]TxAllocColorPlanesReply
	Fcolors             [0]TxQueryColorsReply
	FlookupColor        [0]TxLookupColorReply
	FbestSize           [0]TxQueryBestSizeReply
	Fextension          [0]TxQueryExtensionReply
	Fextensions         [0]TxListExtensionsReply
	FsetModifierMapping [0]TxSetModifierMappingReply
	FgetModifierMapping [0]TxGetModifierMappingReply
	FsetPointerMapping  [0]TxSetPointerMappingReply
	FgetKeyboardMapping [0]TxGetKeyboardMappingReply
	FgetPointerMapping  [0]TxGetPointerMappingReply
	FpointerControl     [0]TxGetPointerControlReply
	FscreenSaver        [0]TxGetScreenSaverReply
	Fhosts              [0]TxListHostsReply
	Ferror1             [0]TxError
	Fevent              [0]TxEvent
	Fgeneric            TxGenericReply
}

type TxReq = struct {
	FreqType TCARD8
	Fdata    TCARD8
	Flength  TCARD16
}

type T_xReq = TxReq

type TxResourceReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fid      TCARD32
}

type TxCreateWindowReq = struct {
	FreqType     TCARD8
	Fdepth       TCARD8
	Flength      TCARD16
	Fwid         TCARD32
	Fparent      TCARD32
	Fx           TINT16
	Fy           TINT16
	Fwidth       TCARD16
	Fheight      TCARD16
	FborderWidth TCARD16
	Fclass       TCARD16
	Fvisual      TCARD32
	Fmask        TCARD32
}

type TxChangeWindowAttributesReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fwindow    TCARD32
	FvalueMask TCARD32
}

type TxChangeSaveSetReq = struct {
	FreqType TCARD8
	Fmode    TBYTE
	Flength  TCARD16
	Fwindow  TCARD32
}

type TxReparentWindowReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fwindow  TCARD32
	Fparent  TCARD32
	Fx       TINT16
	Fy       TINT16
}

type TxConfigureWindowReq = struct {
	FreqType TCARD8
	Fpad     TCARD8
	Flength  TCARD16
	Fwindow  TCARD32
	Fmask    TCARD16
	Fpad2    TCARD16
}

type TxCirculateWindowReq = struct {
	FreqType   TCARD8
	Fdirection TCARD8
	Flength    TCARD16
	Fwindow    TCARD32
}

type TxInternAtomReq = struct {
	FreqType      TCARD8
	FonlyIfExists TBOOL
	Flength       TCARD16
	Fnbytes       TCARD16
	Fpad          TCARD16
}

type TxChangePropertyReq = struct {
	FreqType  TCARD8
	Fmode     TCARD8
	Flength   TCARD16
	Fwindow   TCARD32
	Fproperty TCARD32
	Ftype1    TCARD32
	Fformat   TCARD8
	Fpad      [3]TBYTE
	FnUnits   TCARD32
}

type TxDeletePropertyReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fwindow   TCARD32
	Fproperty TCARD32
}

type TxGetPropertyReq = struct {
	FreqType    TCARD8
	Fdelete1    TBOOL
	Flength     TCARD16
	Fwindow     TCARD32
	Fproperty   TCARD32
	Ftype1      TCARD32
	FlongOffset TCARD32
	FlongLength TCARD32
}

type TxSetSelectionOwnerReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fwindow    TCARD32
	Fselection TCARD32
	Ftime      TCARD32
}

type TxConvertSelectionReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Frequestor TCARD32
	Fselection TCARD32
	Ftarget    TCARD32
	Fproperty  TCARD32
	Ftime      TCARD32
}

type TxSendEventReq = struct {
	FreqType     TCARD8
	Fpropagate   TBOOL
	Flength      TCARD16
	Fdestination TCARD32
	FeventMask   TCARD32
	Fevent       TxEvent
}

type TxGrabPointerReq = struct {
	FreqType      TCARD8
	FownerEvents  TBOOL
	Flength       TCARD16
	FgrabWindow   TCARD32
	FeventMask    TCARD16
	FpointerMode  TBYTE
	FkeyboardMode TBYTE
	FconfineTo    TCARD32
	Fcursor       TCARD32
	Ftime         TCARD32
}

type TxGrabButtonReq = struct {
	FreqType      TCARD8
	FownerEvents  TBOOL
	Flength       TCARD16
	FgrabWindow   TCARD32
	FeventMask    TCARD16
	FpointerMode  TBYTE
	FkeyboardMode TBYTE
	FconfineTo    TCARD32
	Fcursor       TCARD32
	Fbutton       TCARD8
	Fpad          TBYTE
	Fmodifiers    TCARD16
}

type TxUngrabButtonReq = struct {
	FreqType    TCARD8
	Fbutton     TCARD8
	Flength     TCARD16
	FgrabWindow TCARD32
	Fmodifiers  TCARD16
	Fpad        TCARD16
}

type TxChangeActivePointerGrabReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fcursor    TCARD32
	Ftime      TCARD32
	FeventMask TCARD16
	Fpad2      TCARD16
}

type TxGrabKeyboardReq = struct {
	FreqType      TCARD8
	FownerEvents  TBOOL
	Flength       TCARD16
	FgrabWindow   TCARD32
	Ftime         TCARD32
	FpointerMode  TBYTE
	FkeyboardMode TBYTE
	Fpad          TCARD16
}

type TxGrabKeyReq = struct {
	FreqType      TCARD8
	FownerEvents  TBOOL
	Flength       TCARD16
	FgrabWindow   TCARD32
	Fmodifiers    TCARD16
	Fkey          TCARD8
	FpointerMode  TBYTE
	FkeyboardMode TBYTE
	Fpad1         TBYTE
	Fpad2         TBYTE
	Fpad3         TBYTE
}

type TxUngrabKeyReq = struct {
	FreqType    TCARD8
	Fkey        TCARD8
	Flength     TCARD16
	FgrabWindow TCARD32
	Fmodifiers  TCARD16
	Fpad        TCARD16
}

type TxAllowEventsReq = struct {
	FreqType TCARD8
	Fmode    TCARD8
	Flength  TCARD16
	Ftime    TCARD32
}

type TxGetMotionEventsReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fwindow  TCARD32
	Fstart   TCARD32
	Fstop    TCARD32
}

type TxTranslateCoordsReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	FsrcWid  TCARD32
	FdstWid  TCARD32
	FsrcX    TINT16
	FsrcY    TINT16
}

type TxWarpPointerReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	FsrcWid    TCARD32
	FdstWid    TCARD32
	FsrcX      TINT16
	FsrcY      TINT16
	FsrcWidth  TCARD16
	FsrcHeight TCARD16
	FdstX      TINT16
	FdstY      TINT16
}

type TxSetInputFocusReq = struct {
	FreqType  TCARD8
	FrevertTo TCARD8
	Flength   TCARD16
	Ffocus    TCARD32
	Ftime     TCARD32
}

type TxOpenFontReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Ffid     TCARD32
	Fnbytes  TCARD16
	Fpad1    TBYTE
	Fpad2    TBYTE
}

type TxQueryTextExtentsReq = struct {
	FreqType   TCARD8
	FoddLength TBOOL
	Flength    TCARD16
	Ffid       TCARD32
}

type TxListFontsReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	FmaxNames TCARD16
	Fnbytes   TCARD16
}

type TxListFontsWithInfoReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	FmaxNames TCARD16
	Fnbytes   TCARD16
}

type TxSetFontPathReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	FnFonts  TCARD16
	Fpad1    TBYTE
	Fpad2    TBYTE
}

type TxCreatePixmapReq = struct {
	FreqType  TCARD8
	Fdepth    TCARD8
	Flength   TCARD16
	Fpid      TCARD32
	Fdrawable TCARD32
	Fwidth    TCARD16
	Fheight   TCARD16
}

type TxCreateGCReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fgc       TCARD32
	Fdrawable TCARD32
	Fmask     TCARD32
}

type TxChangeGCReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fgc      TCARD32
	Fmask    TCARD32
}

type TxCopyGCReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	FsrcGC   TCARD32
	FdstGC   TCARD32
	Fmask    TCARD32
}

type TxSetDashesReq = struct {
	FreqType    TCARD8
	Fpad        TBYTE
	Flength     TCARD16
	Fgc         TCARD32
	FdashOffset TCARD16
	FnDashes    TCARD16
}

type TxSetClipRectanglesReq = struct {
	FreqType  TCARD8
	Fordering TBYTE
	Flength   TCARD16
	Fgc       TCARD32
	FxOrigin  TINT16
	FyOrigin  TINT16
}

type TxClearAreaReq = struct {
	FreqType   TCARD8
	Fexposures TBOOL
	Flength    TCARD16
	Fwindow    TCARD32
	Fx         TINT16
	Fy         TINT16
	Fwidth     TCARD16
	Fheight    TCARD16
}

type TxCopyAreaReq = struct {
	FreqType     TCARD8
	Fpad         TBYTE
	Flength      TCARD16
	FsrcDrawable TCARD32
	FdstDrawable TCARD32
	Fgc          TCARD32
	FsrcX        TINT16
	FsrcY        TINT16
	FdstX        TINT16
	FdstY        TINT16
	Fwidth       TCARD16
	Fheight      TCARD16
}

type TxCopyPlaneReq = struct {
	FreqType     TCARD8
	Fpad         TBYTE
	Flength      TCARD16
	FsrcDrawable TCARD32
	FdstDrawable TCARD32
	Fgc          TCARD32
	FsrcX        TINT16
	FsrcY        TINT16
	FdstX        TINT16
	FdstY        TINT16
	Fwidth       TCARD16
	Fheight      TCARD16
	FbitPlane    TCARD32
}

type TxPolyPointReq = struct {
	FreqType   TCARD8
	FcoordMode TBYTE
	Flength    TCARD16
	Fdrawable  TCARD32
	Fgc        TCARD32
}

type TxPolyLineReq = struct {
	FreqType   TCARD8
	FcoordMode TBYTE
	Flength    TCARD16
	Fdrawable  TCARD32
	Fgc        TCARD32
}

type TxPolySegmentReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
}

type TxPolyArcReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
}

type TxPolyRectangleReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
}

type TxPolyFillRectangleReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
}

type TxPolyFillArcReq = struct {
	FreqType  TCARD8
	Fpad      TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
}

type TxFillPolyReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fdrawable  TCARD32
	Fgc        TCARD32
	Fshape     TBYTE
	FcoordMode TBYTE
	Fpad1      TCARD16
}

type T_FillPolyReq = TxFillPolyReq

type TxPutImageReq = struct {
	FreqType  TCARD8
	Fformat   TCARD8
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fwidth    TCARD16
	Fheight   TCARD16
	FdstX     TINT16
	FdstY     TINT16
	FleftPad  TCARD8
	Fdepth    TCARD8
	Fpad      TCARD16
}

type T_PutImageReq = TxPutImageReq

type TxGetImageReq = struct {
	FreqType   TCARD8
	Fformat    TCARD8
	Flength    TCARD16
	Fdrawable  TCARD32
	Fx         TINT16
	Fy         TINT16
	Fwidth     TCARD16
	Fheight    TCARD16
	FplaneMask TCARD32
}

type TxPolyTextReq = struct {
	FreqType  TCARD8
	Fpad      TCARD8
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fx        TINT16
	Fy        TINT16
}

type TxPolyText8Req = struct {
	FreqType  TCARD8
	Fpad      TCARD8
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fx        TINT16
	Fy        TINT16
}

type TxPolyText16Req = struct {
	FreqType  TCARD8
	Fpad      TCARD8
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fx        TINT16
	Fy        TINT16
}

type TxImageTextReq = struct {
	FreqType  TCARD8
	FnChars   TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fx        TINT16
	Fy        TINT16
}

type TxImageText8Req = struct {
	FreqType  TCARD8
	FnChars   TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fx        TINT16
	Fy        TINT16
}

type TxImageText16Req = struct {
	FreqType  TCARD8
	FnChars   TBYTE
	Flength   TCARD16
	Fdrawable TCARD32
	Fgc       TCARD32
	Fx        TINT16
	Fy        TINT16
}

type TxCreateColormapReq = struct {
	FreqType TCARD8
	Falloc   TBYTE
	Flength  TCARD16
	Fmid     TCARD32
	Fwindow  TCARD32
	Fvisual  TCARD32
}

type TxCopyColormapAndFreeReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fmid     TCARD32
	FsrcCmap TCARD32
}

type TxAllocColorReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fcmap    TCARD32
	Fred     TCARD16
	Fgreen   TCARD16
	Fblue    TCARD16
	Fpad2    TCARD16
}

type TxAllocNamedColorReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fcmap    TCARD32
	Fnbytes  TCARD16
	Fpad1    TBYTE
	Fpad2    TBYTE
}

type TxAllocColorCellsReq = struct {
	FreqType    TCARD8
	Fcontiguous TBOOL
	Flength     TCARD16
	Fcmap       TCARD32
	Fcolors     TCARD16
	Fplanes     TCARD16
}

type TxAllocColorPlanesReq = struct {
	FreqType    TCARD8
	Fcontiguous TBOOL
	Flength     TCARD16
	Fcmap       TCARD32
	Fcolors     TCARD16
	Fred        TCARD16
	Fgreen      TCARD16
	Fblue       TCARD16
}

type TxFreeColorsReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fcmap      TCARD32
	FplaneMask TCARD32
}

type TxStoreColorsReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fcmap    TCARD32
}

type TxStoreNamedColorReq = struct {
	FreqType TCARD8
	Fflags   TCARD8
	Flength  TCARD16
	Fcmap    TCARD32
	Fpixel   TCARD32
	Fnbytes  TCARD16
	Fpad1    TBYTE
	Fpad2    TBYTE
}

type TxQueryColorsReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fcmap    TCARD32
}

type TxLookupColorReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fcmap    TCARD32
	Fnbytes  TCARD16
	Fpad1    TBYTE
	Fpad2    TBYTE
}

type TxCreateCursorReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fcid       TCARD32
	Fsource    TCARD32
	Fmask      TCARD32
	FforeRed   TCARD16
	FforeGreen TCARD16
	FforeBlue  TCARD16
	FbackRed   TCARD16
	FbackGreen TCARD16
	FbackBlue  TCARD16
	Fx         TCARD16
	Fy         TCARD16
}

type TxCreateGlyphCursorReq = struct {
	FreqType    TCARD8
	Fpad        TBYTE
	Flength     TCARD16
	Fcid        TCARD32
	Fsource     TCARD32
	Fmask       TCARD32
	FsourceChar TCARD16
	FmaskChar   TCARD16
	FforeRed    TCARD16
	FforeGreen  TCARD16
	FforeBlue   TCARD16
	FbackRed    TCARD16
	FbackGreen  TCARD16
	FbackBlue   TCARD16
}

type TxRecolorCursorReq = struct {
	FreqType   TCARD8
	Fpad       TBYTE
	Flength    TCARD16
	Fcursor    TCARD32
	FforeRed   TCARD16
	FforeGreen TCARD16
	FforeBlue  TCARD16
	FbackRed   TCARD16
	FbackGreen TCARD16
	FbackBlue  TCARD16
}

type TxQueryBestSizeReq = struct {
	FreqType  TCARD8
	Fclass    TCARD8
	Flength   TCARD16
	Fdrawable TCARD32
	Fwidth    TCARD16
	Fheight   TCARD16
}

type TxQueryExtensionReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fnbytes  TCARD16
	Fpad1    TBYTE
	Fpad2    TBYTE
}

type TxSetModifierMappingReq = struct {
	FreqType           TCARD8
	FnumKeyPerModifier TCARD8
	Flength            TCARD16
}

type TxSetPointerMappingReq = struct {
	FreqType TCARD8
	FnElts   TCARD8
	Flength  TCARD16
}

type TxGetKeyboardMappingReq = struct {
	FreqType      TCARD8
	Fpad          TBYTE
	Flength       TCARD16
	FfirstKeyCode TCARD8
	Fcount        TCARD8
	Fpad1         TCARD16
}

type TxChangeKeyboardMappingReq = struct {
	FreqType           TCARD8
	FkeyCodes          TCARD8
	Flength            TCARD16
	FfirstKeyCode      TCARD8
	FkeySymsPerKeyCode TCARD8
	Fpad1              TCARD16
}

type TxChangeKeyboardControlReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
	Fmask    TCARD32
}

type TxBellReq = struct {
	FreqType TCARD8
	Fpercent TINT8
	Flength  TCARD16
}

type TxChangePointerControlReq = struct {
	FreqType    TCARD8
	Fpad        TBYTE
	Flength     TCARD16
	FaccelNum   TINT16
	FaccelDenum TINT16
	Fthreshold  TINT16
	FdoAccel    TBOOL
	FdoThresh   TBOOL
}

type TxSetScreenSaverReq = struct {
	FreqType     TCARD8
	Fpad         TBYTE
	Flength      TCARD16
	Ftimeout     TINT16
	Finterval    TINT16
	FpreferBlank TBYTE
	FallowExpose TBYTE
	Fpad2        TCARD16
}

type TxChangeHostsReq = struct {
	FreqType    TCARD8
	Fmode       TBYTE
	Flength     TCARD16
	FhostFamily TCARD8
	Fpad        TBYTE
	FhostLength TCARD16
}

type TxListHostsReq = struct {
	FreqType TCARD8
	Fpad     TBYTE
	Flength  TCARD16
}

type TxChangeModeReq = struct {
	FreqType TCARD8
	Fmode    TBYTE
	Flength  TCARD16
}

type TxSetAccessControlReq = struct {
	FreqType TCARD8
	Fmode    TBYTE
	Flength  TCARD16
}

type TxSetCloseDownModeReq = struct {
	FreqType TCARD8
	Fmode    TBYTE
	Flength  TCARD16
}

type TxForceScreenSaverReq = struct {
	FreqType TCARD8
	Fmode    TBYTE
	Flength  TCARD16
}

type TxRotatePropertiesReq = struct {
	FreqType    TCARD8
	Fpad        TBYTE
	Flength     TCARD16
	Fwindow     TCARD32
	FnAtoms     TCARD16
	FnPositions TINT16
}

type Tsigval = struct {
	Fsival_ptr   [0]uintptr
	Fsigval_int  [0]int32
	Fsigval_ptr  [0]uintptr
	Fsival_int   int32
	F__ccgo_pad4 [4]byte
}

type Tsched_param = struct {
	Fsched_priority int32
}

type T_cpuset = struct {
	F__bits [16]int64
}

type Tcpuset_t = struct {
	F__bits [16]int64
}

type Tcpu_set_t = struct {
	F__bits [16]int64
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

type Tlocale_t = uintptr

type _pthread_mutextype = int32

const _PTHREAD_MUTEX_ERRORCHECK = 1
const _PTHREAD_MUTEX_RECURSIVE = 2
const _PTHREAD_MUTEX_NORMAL = 3
const _PTHREAD_MUTEX_ADAPTIVE_NP = 4
const _PTHREAD_MUTEX_TYPE_MAX = 5

type T_pthread_cleanup_info = struct {
	Fpthread_cleanup_pad [8]t__uintptr_t
}

type Txthread_t = uintptr

type Txthread_key_t = int32

type Txcondition_rec = uintptr

type Txmutex_rec = uintptr

type Txcondition_t = uintptr

type Txmutex_t = uintptr

type T_XQEvent = struct {
	Fnext        uintptr
	Fevent       TXEvent
	Fqserial_num uint64
}

type T_XSQEvent = T_XQEvent

type Terrno_t = int32

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

type Tconstraint_handler_t = uintptr

type TLockInfoPtr = uintptr

type T_XLockPtrs = struct {
	Flock_display   uintptr
	Funlock_display uintptr
}

type T_XAsyncHandler = struct {
	Fnext    uintptr
	Fhandler uintptr
	Fdata    TXPointer
}

type T_XInternalAsync = T_XAsyncHandler

type T_XAsyncErrorState = struct {
	Fmin_sequence_number uint64
	Fmax_sequence_number uint64
	Ferror_code          uint8
	Fmajor_opcode        uint8
	Fminor_opcode        uint16
	Flast_error_received uint8
	Ferror_count         int32
}

type T_XAsyncEState = T_XAsyncErrorState

type TFreeFuncType = uintptr

type TFreeModmapType = uintptr

type T_XFreeFuncRec = struct {
	Fatoms         TFreeFuncType
	Fmodifiermap   TFreeModmapType
	Fkey_bindings  TFreeFuncType
	Fcontext_db    TFreeFuncType
	FdefaultCCCs   TFreeFuncType
	FclientCmaps   TFreeFuncType
	FintensityMaps TFreeFuncType
	Fim_filters    TFreeFuncType
	Fxkb           TFreeFuncType
}

type T_XFreeFuncs = T_XFreeFuncRec

type TCreateGCType = uintptr

type TCopyGCType = uintptr

type TFlushGCType = uintptr

type TFreeGCType = uintptr

type TCreateFontType = uintptr

type TFreeFontType = uintptr

type TCloseDisplayType = uintptr

type TErrorType = uintptr

type TErrorStringType = uintptr

type TPrintErrorType = uintptr

type TBeforeFlushType = uintptr

type T_XExtension = struct {
	Fnext          uintptr
	Fcodes         TXExtCodes
	Fcreate_GC     TCreateGCType
	Fcopy_GC       TCopyGCType
	Fflush_GC      TFlushGCType
	Ffree_GC       TFreeGCType
	Fcreate_Font   TCreateFontType
	Ffree_Font     TFreeFontType
	Fclose_display TCloseDisplayType
	Ferror1        TErrorType
	Ferror_string  TErrorStringType
	Fname          uintptr
	Ferror_values  TPrintErrorType
	Fbefore_flush  TBeforeFlushType
	Fnext_flush    uintptr
}

type T_XExten = T_XExtension

type T_XInternalConnectionProc = uintptr

type T_XConnectionInfo = struct {
	Ffd            int32
	Fread_callback T_XInternalConnectionProc
	Fcall_data     TXPointer
	Fwatch_data    uintptr
	Fnext          uintptr
}

type T_XConnWatchInfo = struct {
	Ffn          TXConnectionWatchProc
	Fclient_data TXPointer
	Fnext        uintptr
}

type TXSizeHints = struct {
	Fflags      int64
	Fx          int32
	Fy          int32
	Fwidth      int32
	Fheight     int32
	Fmin_width  int32
	Fmin_height int32
	Fmax_width  int32
	Fmax_height int32
	Fwidth_inc  int32
	Fheight_inc int32
	Fmin_aspect struct {
		Fx int32
		Fy int32
	}
	Fmax_aspect struct {
		Fx int32
		Fy int32
	}
	Fbase_width  int32
	Fbase_height int32
	Fwin_gravity int32
}

type TXWMHints = struct {
	Fflags         int64
	Finput         int32
	Finitial_state int32
	Ficon_pixmap   TPixmap
	Ficon_window   TWindow
	Ficon_x        int32
	Ficon_y        int32
	Ficon_mask     TPixmap
	Fwindow_group  TXID
}

type TXTextProperty = struct {
	Fvalue    uintptr
	Fencoding TAtom
	Fformat   int32
	Fnitems   uint64
}

type TXICCEncodingStyle = int32

const _XStringStyle = 0
const _XCompoundTextStyle = 1
const _XTextStyle = 2
const _XStdICCTextStyle = 3
const _XUTF8StringStyle = 4

type TXIconSize = struct {
	Fmin_width  int32
	Fmin_height int32
	Fmax_width  int32
	Fmax_height int32
	Fwidth_inc  int32
	Fheight_inc int32
}

type TXClassHint = struct {
	Fres_name  uintptr
	Fres_class uintptr
}

type TXComposeStatus = struct {
	Fcompose_ptr   TXPointer
	Fchars_matched int32
}

type T_XComposeStatus = TXComposeStatus

type TRegion = uintptr

type TXVisualInfo = struct {
	Fvisual        uintptr
	Fvisualid      TVisualID
	Fscreen        int32
	Fdepth         int32
	Fclass         int32
	Fred_mask      uint64
	Fgreen_mask    uint64
	Fblue_mask     uint64
	Fcolormap_size int32
	Fbits_per_rgb  int32
}

type TXStandardColormap = struct {
	Fcolormap   TColormap
	Fred_max    uint64
	Fred_mult   uint64
	Fgreen_max  uint64
	Fgreen_mult uint64
	Fblue_max   uint64
	Fblue_mult  uint64
	Fbase_pixel uint64
	Fvisualid   TVisualID
	Fkillid     TXID
}

type TXContext = int32

type Tpointer = uintptr

type TClientPtr = uintptr

type TFontPtr = uintptr

type TFSID = uint64

type TAccContext = uint64

type TOSTimePtr = uintptr

type TBlockHandlerProcPtr = uintptr

type TGlyph = uint64

type TGlyphSet = uint64

type TPicture = uint64

type TPictFormat = uint64

type TxDirectFormat = struct {
	Fred       TCARD16
	FredMask   TCARD16
	Fgreen     TCARD16
	FgreenMask TCARD16
	Fblue      TCARD16
	FblueMask  TCARD16
	Falpha     TCARD16
	FalphaMask TCARD16
}

type TxPictFormInfo = struct {
	Fid       TCARD32
	Ftype1    TCARD8
	Fdepth    TCARD8
	Fpad1     TCARD16
	Fdirect   TxDirectFormat
	Fcolormap TCARD32
}

type TxPictVisual = struct {
	Fvisual TCARD32
	Fformat TCARD32
}

type TxPictDepth = struct {
	Fdepth        TCARD8
	Fpad1         TCARD8
	FnPictVisuals TCARD16
	Fpad2         TCARD32
}

type TxPictScreen = struct {
	FnDepth   TCARD32
	Ffallback TCARD32
}

type TxIndexValue = struct {
	Fpixel TCARD32
	Fred   TCARD16
	Fgreen TCARD16
	Fblue  TCARD16
	Falpha TCARD16
}

type TxRenderColor = struct {
	Fred   TCARD16
	Fgreen TCARD16
	Fblue  TCARD16
	Falpha TCARD16
}

type TxPointFixed = struct {
	Fx TINT32
	Fy TINT32
}

type TxLineFixed = struct {
	Fp1 TxPointFixed
	Fp2 TxPointFixed
}

type TxTriangle = struct {
	Fp1 TxPointFixed
	Fp2 TxPointFixed
	Fp3 TxPointFixed
}

type TxTrapezoid = struct {
	Ftop    TINT32
	Fbottom TINT32
	Fleft   TxLineFixed
	Fright  TxLineFixed
}

type TxGlyphInfo = struct {
	Fwidth  TCARD16
	Fheight TCARD16
	Fx      TINT16
	Fy      TINT16
	FxOff   TINT16
	FyOff   TINT16
}

type TxGlyphElt = struct {
	Flen1   TCARD8
	Fpad1   TCARD8
	Fpad2   TCARD16
	Fdeltax TINT16
	Fdeltay TINT16
}

type TxSpanFix = struct {
	Fl TINT32
	Fr TINT32
	Fy TINT32
}

type TxTrap = struct {
	Ftop TxSpanFix
	Fbot TxSpanFix
}

type TxRenderQueryVersionReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	FmajorVersion  TCARD32
	FminorVersion  TCARD32
}

type TxRenderQueryVersionReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FmajorVersion   TCARD32
	FminorVersion   TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxRenderQueryPictFormatsReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
}

type TxRenderQueryPictFormatsReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnumFormats     TCARD32
	FnumScreens     TCARD32
	FnumDepths      TCARD32
	FnumVisuals     TCARD32
	FnumSubpixel    TCARD32
	Fpad5           TCARD32
}

type TxRenderQueryPictIndexValuesReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fformat        TCARD32
}

type TxRenderQueryPictIndexValuesReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnumIndexValues TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
	Fpad6           TCARD32
}

type TxRenderCreatePictureReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpid           TCARD32
	Fdrawable      TCARD32
	Fformat        TCARD32
	Fmask          TCARD32
}

type TxRenderChangePictureReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpicture       TCARD32
	Fmask          TCARD32
}

type TxRenderSetPictureClipRectanglesReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpicture       TCARD32
	FxOrigin       TINT16
	FyOrigin       TINT16
}

type TxRenderFreePictureReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpicture       TCARD32
}

type TxRenderCompositeReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fmask          TCARD32
	Fdst           TCARD32
	FxSrc          TINT16
	FySrc          TINT16
	FxMask         TINT16
	FyMask         TINT16
	FxDst          TINT16
	FyDst          TINT16
	Fwidth         TCARD16
	Fheight        TCARD16
}

type TxRenderScaleReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FcolorScale    TCARD32
	FalphaScale    TCARD32
	FxSrc          TINT16
	FySrc          TINT16
	FxDst          TINT16
	FyDst          TINT16
	Fwidth         TCARD16
	Fheight        TCARD16
}

type TxRenderTrapezoidsReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderTrianglesReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderTriStripReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderTriFanReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderCreateGlyphSetReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fgsid          TCARD32
	Fformat        TCARD32
}

type TxRenderReferenceGlyphSetReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fgsid          TCARD32
	Fexisting      TCARD32
}

type TxRenderFreeGlyphSetReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fglyphset      TCARD32
}

type TxRenderAddGlyphsReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fglyphset      TCARD32
	Fnglyphs       TCARD32
}

type TxRenderFreeGlyphsReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fglyphset      TCARD32
}

type TxRenderCompositeGlyphsReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	Fglyphset      TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderCompositeGlyphs8Req = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	Fglyphset      TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderCompositeGlyphs16Req = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	Fglyphset      TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderCompositeGlyphs32Req = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fsrc           TCARD32
	Fdst           TCARD32
	FmaskFormat    TCARD32
	Fglyphset      TCARD32
	FxSrc          TINT16
	FySrc          TINT16
}

type TxRenderFillRectanglesReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fop            TCARD8
	Fpad1          TCARD8
	Fpad2          TCARD16
	Fdst           TCARD32
	Fcolor         TxRenderColor
}

type TxRenderCreateCursorReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fcid           TCARD32
	Fsrc           TCARD32
	Fx             TCARD16
	Fy             TCARD16
}

type TxRenderTransform = struct {
	Fmatrix11 TINT32
	Fmatrix12 TINT32
	Fmatrix13 TINT32
	Fmatrix21 TINT32
	Fmatrix22 TINT32
	Fmatrix23 TINT32
	Fmatrix31 TINT32
	Fmatrix32 TINT32
	Fmatrix33 TINT32
}

type TxRenderSetPictureTransformReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpicture       TCARD32
	Ftransform     TxRenderTransform
}

type TxRenderQueryFiltersReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fdrawable      TCARD32
}

type TxRenderQueryFiltersReply = struct {
	Ftype1          TBYTE
	Fpad1           TBYTE
	FsequenceNumber TCARD16
	Flength         TCARD32
	FnumAliases     TCARD32
	FnumFilters     TCARD32
	Fpad2           TCARD32
	Fpad3           TCARD32
	Fpad4           TCARD32
	Fpad5           TCARD32
}

type TxRenderSetPictureFilterReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpicture       TCARD32
	Fnbytes        TCARD16
	Fpad           TCARD16
}

type TxAnimCursorElt = struct {
	Fcursor TCARD32
	Fdelay  TCARD32
}

type TxRenderCreateAnimCursorReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fcid           TCARD32
}

type TxRenderAddTrapsReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpicture       TCARD32
	FxOff          TINT16
	FyOff          TINT16
}

type TxRenderCreateSolidFillReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpid           TCARD32
	Fcolor         TxRenderColor
}

type TxRenderCreateLinearGradientReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpid           TCARD32
	Fp1            TxPointFixed
	Fp2            TxPointFixed
	FnStops        TCARD32
}

type TxRenderCreateRadialGradientReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpid           TCARD32
	Finner         TxPointFixed
	Fouter         TxPointFixed
	Finner_radius  TINT32
	Fouter_radius  TINT32
	FnStops        TCARD32
}

type TxRenderCreateConicalGradientReq = struct {
	FreqType       TCARD8
	FrenderReqType TCARD8
	Flength        TCARD16
	Fpid           TCARD32
	Fcenter        TxPointFixed
	Fangle         TINT32
	FnStops        TCARD32
}

type TXRenderDirectFormat = struct {
	Fred       int16
	FredMask   int16
	Fgreen     int16
	FgreenMask int16
	Fblue      int16
	FblueMask  int16
	Falpha     int16
	FalphaMask int16
}

type TXRenderPictFormat = struct {
	Fid       TPictFormat
	Ftype1    int32
	Fdepth    int32
	Fdirect   TXRenderDirectFormat
	Fcolormap TColormap
}

type TXRenderPictureAttributes = struct {
	Frepeat             int32
	Falpha_map          TPicture
	Falpha_x_origin     int32
	Falpha_y_origin     int32
	Fclip_x_origin      int32
	Fclip_y_origin      int32
	Fclip_mask          TPixmap
	Fgraphics_exposures int32
	Fsubwindow_mode     int32
	Fpoly_edge          int32
	Fpoly_mode          int32
	Fdither             TAtom
	Fcomponent_alpha    int32
}

type T_XRenderPictureAttributes = TXRenderPictureAttributes

type TXRenderColor = struct {
	Fred   uint16
	Fgreen uint16
	Fblue  uint16
	Falpha uint16
}

type TXGlyphInfo = struct {
	Fwidth  uint16
	Fheight uint16
	Fx      int16
	Fy      int16
	FxOff   int16
	FyOff   int16
}

type T_XGlyphInfo = TXGlyphInfo

type TXGlyphElt8 = struct {
	Fglyphset TGlyphSet
	Fchars    uintptr
	Fnchars   int32
	FxOff     int32
	FyOff     int32
}

type T_XGlyphElt8 = TXGlyphElt8

type TXGlyphElt16 = struct {
	Fglyphset TGlyphSet
	Fchars    uintptr
	Fnchars   int32
	FxOff     int32
	FyOff     int32
}

type T_XGlyphElt16 = TXGlyphElt16

type TXGlyphElt32 = struct {
	Fglyphset TGlyphSet
	Fchars    uintptr
	Fnchars   int32
	FxOff     int32
	FyOff     int32
}

type T_XGlyphElt32 = TXGlyphElt32

type TXDouble = float64

type TXPointDouble = struct {
	Fx TXDouble
	Fy TXDouble
}

type T_XPointDouble = TXPointDouble

type TXFixed = int32

type TXPointFixed = struct {
	Fx TXFixed
	Fy TXFixed
}

type T_XPointFixed = TXPointFixed

type TXLineFixed = struct {
	Fp1 TXPointFixed
	Fp2 TXPointFixed
}

type T_XLineFixed = TXLineFixed

type TXTriangle = struct {
	Fp1 TXPointFixed
	Fp2 TXPointFixed
	Fp3 TXPointFixed
}

type T_XTriangle = TXTriangle

type TXCircle = struct {
	Fx      TXFixed
	Fy      TXFixed
	Fradius TXFixed
}

type T_XCircle = TXCircle

type TXTrapezoid = struct {
	Ftop    TXFixed
	Fbottom TXFixed
	Fleft   TXLineFixed
	Fright  TXLineFixed
}

type T_XTrapezoid = TXTrapezoid

type TXTransform = struct {
	Fmatrix [3][3]TXFixed
}

type T_XTransform = TXTransform

type TXFilters = struct {
	Fnfilter int32
	Ffilter  uintptr
	Fnalias  int32
	Falias   uintptr
}

type T_XFilters = TXFilters

type TXIndexValue = struct {
	Fpixel uint64
	Fred   uint16
	Fgreen uint16
	Fblue  uint16
	Falpha uint16
}

type T_XIndexValue = TXIndexValue

type TXAnimCursor = struct {
	Fcursor TCursor
	Fdelay  uint64
}

type T_XAnimCursor = TXAnimCursor

type TXSpanFix = struct {
	Fleft  TXFixed
	Fright TXFixed
	Fy     TXFixed
}

type T_XSpanFix = TXSpanFix

type TXTrap = struct {
	Ftop    TXSpanFix
	Fbottom TXSpanFix
}

type T_XTrap = TXTrap

type TXLinearGradient = struct {
	Fp1 TXPointFixed
	Fp2 TXPointFixed
}

type T_XLinearGradient = TXLinearGradient

type TXRadialGradient = struct {
	Finner TXCircle
	Fouter TXCircle
}

type T_XRadialGradient = TXRadialGradient

type TXConicalGradient = struct {
	Fcenter TXPointFixed
	Fangle  TXFixed
}

type T_XConicalGradient = TXConicalGradient

type TXRenderVisual = struct {
	Fvisual uintptr
	Fformat uintptr
}

type TXRenderDepth = struct {
	Fdepth    int32
	Fnvisuals int32
	Fvisuals  uintptr
}

type TXRenderScreen = struct {
	Fdepths   uintptr
	Fndepths  int32
	Ffallback uintptr
	Fsubpixel int32
}

type TXRenderInfo = struct {
	Fmajor_version int32
	Fminor_version int32
	Fformat        uintptr
	Fnformat       int32
	Fscreen        uintptr
	Fnscreen       int32
	Fdepth         uintptr
	Fndepth        int32
	Fvisual        uintptr
	Fnvisual       int32
	Fsubpixel      uintptr
	Fnsubpixel     int32
	Ffilter        uintptr
	Fnfilter       int32
	Ffilter_alias  uintptr
	Fnfilter_alias int32
}

type T_XRenderInfo = TXRenderInfo

type TXRenderExtDisplayInfo = struct {
	Fnext    uintptr
	Fdisplay uintptr
	Fcodes   uintptr
	Finfo    uintptr
}

type T_XRenderExtDisplayInfo = TXRenderExtDisplayInfo

type TXRenderExtInfo = struct {
	Fhead      uintptr
	Fcur       uintptr
	Fndisplays int32
}

type T_XRenderExtInfo = TXRenderExtInfo

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

func XXRenderAddTraps(tls *libc.TLS, dpy uintptr, picture TPicture, xOff int32, yOff int32, traps uintptr, ntrap int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var max_req, v1 uint64
	var n int32
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _, _ = _BRlen, info, len1, max_req, n, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
		v1 = (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size
	} else {
		v1 = uint64((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size)
	}
	max_req = v1
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	for ntrap != 0 {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderAddTraps), uint64(m_sz_xRenderAddTrapsReq))
		(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderAddTraps)
		(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Fpicture = uint32(picture)
		(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).FxOff = int16(xOff)
		(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).FyOff = int16(yOff)
		n = ntrap
		len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xTrap)>>libc.Int32FromInt32(2))
		if libc.Uint64FromInt64(len1) > max_req-uint64((*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength) {
			n = libc.Int32FromUint64((max_req - uint64((*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength)) / libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xTrap)>>libc.Int32FromInt32(2)))
			len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xTrap)>>libc.Int32FromInt32(2))
		}
		if libc.Int64FromUint16((*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
				_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
				(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength = uint16(0)
				*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
				libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
				*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
				libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
			} else {
				len1 = len1
				(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength) + len1)
			}
		} else {
			(*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderAddTrapsReq)(unsafe.Pointer(req)).Flength) + len1)
		}
		len1 <<= int64(2)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, traps, libc.Uint64FromInt64(len1))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
		} else {
			libx11.X_XSend(tls, dpy, traps, len1)
		}
		ntrap -= n
		traps += uintptr(n) * 24
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

type Tfuncs = struct {
	Fcreate_image  uintptr
	Fdestroy_image uintptr
	Fget_pixel     uintptr
	Fput_pixel     uintptr
	Fsub_image     uintptr
	Fadd_pixel     uintptr
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */
func XXRenderParseColor(tls *libc.TLS, dpy uintptr, spec uintptr, def uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c int8
	var colormap TColormap
	var i, n, v2 int32
	var pShort, v3, p4, p5, p6 uintptr
	var _ /* coreColor at bp+8 */ TXColor
	var _ /* elements at bp+0 */ [4]uint16
	_, _, _, _, _, _, _, _, _, _ = c, colormap, i, n, pShort, v2, v3, p4, p5, p6
	if !(libc.Xstrncmp(tls, spec, __ccgo_ts, uint64(5)) != 0) {
		spec += uintptr(5)
		/*
		 * Attempt to parse the value portion.
		 */
		pShort = bp
		i = 0
		for {
			if !(i < int32(4)) {
				break
			}
			n = 0
			*(*uint16)(unsafe.Pointer(pShort)) = uint16(0)
			for int32(*(*int8)(unsafe.Pointer(spec))) != int32('/') && int32(*(*int8)(unsafe.Pointer(spec))) != int32('\000') {
				n++
				v2 = n
				if v2 > int32(4) {
					return 0
				}
				v3 = spec
				spec++
				c = *(*int8)(unsafe.Pointer(v3))
				p4 = pShort
				*(*uint16)(unsafe.Pointer(p4)) = uint16(int32(*(*uint16)(unsafe.Pointer(p4))) << libc.Int32FromInt32(4))
				if int32(c) >= int32('0') && int32(c) <= int32('9') {
					p5 = pShort
					*(*uint16)(unsafe.Pointer(p5)) = uint16(int32(*(*uint16)(unsafe.Pointer(p5))) | (int32(c) - libc.Int32FromUint8('0')))
				} else {
					if int32(c) >= int32('a') && int32(c) <= int32('f') {
						p6 = pShort
						*(*uint16)(unsafe.Pointer(p6)) = uint16(int32(*(*uint16)(unsafe.Pointer(p6))) | (int32(c) - (libc.Int32FromUint8('a') - libc.Int32FromInt32(10))))
					} else {
						return 0
					}
				}
			}
			if n == 0 {
				return 0
			}
			if n < int32(4) {
				*(*uint16)(unsafe.Pointer(pShort)) = uint16(uint64(*(*uint16)(unsafe.Pointer(pShort))) * uint64(0xFFFF) / libc.Uint64FromInt32(libc.Int32FromInt32(1)<<(n*libc.Int32FromInt32(4))-libc.Int32FromInt32(1)))
			}
			goto _1
		_1:
			;
			i++
			pShort += 2
			spec++
		}
		(*TXRenderColor)(unsafe.Pointer(def)).Fred = (*(*[4]uint16)(unsafe.Pointer(bp)))[0]
		(*TXRenderColor)(unsafe.Pointer(def)).Fgreen = (*(*[4]uint16)(unsafe.Pointer(bp)))[int32(1)]
		(*TXRenderColor)(unsafe.Pointer(def)).Fblue = (*(*[4]uint16)(unsafe.Pointer(bp)))[int32(2)]
		(*TXRenderColor)(unsafe.Pointer(def)).Falpha = (*(*[4]uint16)(unsafe.Pointer(bp)))[int32(3)]
	} else {
		colormap = (*TScreen)(unsafe.Pointer((*struct {
			Fext_data            uintptr
			Fprivate1            uintptr
			Ffd                  int32
			Fprivate2            int32
			Fproto_major_version int32
			Fproto_minor_version int32
			Fvendor              uintptr
			Fprivate3            TXID
			Fprivate4            TXID
			Fprivate5            TXID
			Fprivate6            int32
			Fresource_alloc      uintptr
			Fbyte_order          int32
			Fbitmap_unit         int32
			Fbitmap_pad          int32
			Fbitmap_bit_order    int32
			Fnformats            int32
			Fpixmap_format       uintptr
			Fprivate8            int32
			Frelease             int32
			Fprivate9            uintptr
			Fprivate10           uintptr
			Fqlen                int32
			Flast_request_read   uint64
			Frequest             uint64
			Fprivate11           TXPointer
			Fprivate12           TXPointer
			Fprivate13           TXPointer
			Fprivate14           TXPointer
			Fmax_request_size    uint32
			Fdb                  uintptr
			Fprivate15           uintptr
			Fdisplay_name        uintptr
			Fdefault_screen      int32
			Fnscreens            int32
			Fscreens             uintptr
			Fmotion_buffer       uint64
			Fprivate16           uint64
			Fmin_keycode         int32
			Fmax_keycode         int32
			Fprivate17           TXPointer
			Fprivate18           TXPointer
			Fprivate19           int32
			Fxdefaults           uintptr
		})(unsafe.Pointer(dpy)).Fscreens + uintptr((*struct {
			Fext_data            uintptr
			Fprivate1            uintptr
			Ffd                  int32
			Fprivate2            int32
			Fproto_major_version int32
			Fproto_minor_version int32
			Fvendor              uintptr
			Fprivate3            TXID
			Fprivate4            TXID
			Fprivate5            TXID
			Fprivate6            int32
			Fresource_alloc      uintptr
			Fbyte_order          int32
			Fbitmap_unit         int32
			Fbitmap_pad          int32
			Fbitmap_bit_order    int32
			Fnformats            int32
			Fpixmap_format       uintptr
			Fprivate8            int32
			Frelease             int32
			Fprivate9            uintptr
			Fprivate10           uintptr
			Fqlen                int32
			Flast_request_read   uint64
			Frequest             uint64
			Fprivate11           TXPointer
			Fprivate12           TXPointer
			Fprivate13           TXPointer
			Fprivate14           TXPointer
			Fmax_request_size    uint32
			Fdb                  uintptr
			Fprivate15           uintptr
			Fdisplay_name        uintptr
			Fdefault_screen      int32
			Fnscreens            int32
			Fscreens             uintptr
			Fmotion_buffer       uint64
			Fprivate16           uint64
			Fmin_keycode         int32
			Fmax_keycode         int32
			Fprivate17           TXPointer
			Fprivate18           TXPointer
			Fprivate19           int32
			Fxdefaults           uintptr
		})(unsafe.Pointer(dpy)).Fdefault_screen)*128)).Fcmap
		if !(libx11.XXParseColor(tls, dpy, colormap, spec, bp+8) != 0) {
			return 0
		}
		(*TXRenderColor)(unsafe.Pointer(def)).Fred = (*(*TXColor)(unsafe.Pointer(bp + 8))).Fred
		(*TXRenderColor)(unsafe.Pointer(def)).Fgreen = (*(*TXColor)(unsafe.Pointer(bp + 8))).Fgreen
		(*TXRenderColor)(unsafe.Pointer(def)).Fblue = (*(*TXColor)(unsafe.Pointer(bp + 8))).Fblue
		(*TXRenderColor)(unsafe.Pointer(def)).Falpha = uint16(0xffff)
	}
	(*TXRenderColor)(unsafe.Pointer(def)).Fred = uint16(libc.Uint32FromInt32(libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(def)).Fred)*libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(def)).Falpha)) / uint32(0xffff))
	(*TXRenderColor)(unsafe.Pointer(def)).Fgreen = uint16(libc.Uint32FromInt32(libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(def)).Fgreen)*libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(def)).Falpha)) / uint32(0xffff))
	(*TXRenderColor)(unsafe.Pointer(def)).Fblue = uint16(libc.Uint32FromInt32(libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(def)).Fblue)*libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(def)).Falpha)) / uint32(0xffff))
	return int32(1)
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

func XXRenderComposite(tls *libc.TLS, dpy uintptr, op int32, src TPicture, mask TPicture, dst TPicture, src_x int32, src_y int32, mask_x int32, mask_y int32, dst_x int32, dst_y int32, width uint32, height uint32) {
	var info, req uintptr
	_, _ = info, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderComposite), uint64(m_sz_xRenderCompositeReq))
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderComposite)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).Fmask = uint32(mask)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FxSrc = int16(src_x)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FySrc = int16(src_y)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FxMask = int16(mask_x)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FyMask = int16(mask_y)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FxDst = int16(dst_x)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).FyDst = int16(dst_y)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).Fwidth = uint16(width)
	(*TxRenderCompositeReq)(unsafe.Pointer(req)).Fheight = uint16(height)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

func XXRenderCreateCursor(tls *libc.TLS, dpy uintptr, source TPicture, x uint32, y uint32) (r TCursor) {
	var cid, v1 TCursor
	var info, req uintptr
	_, _, _, _ = cid, info, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateCursor), uint64(m_sz_xRenderCreateCursorReq))
	(*TxRenderCreateCursorReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateCursorReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateCursor)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	cid = v1
	(*TxRenderCreateCursorReq)(unsafe.Pointer(req)).Fcid = uint32(v1)
	(*TxRenderCreateCursorReq)(unsafe.Pointer(req)).Fsrc = uint32(source)
	(*TxRenderCreateCursorReq)(unsafe.Pointer(req)).Fx = uint16(x)
	(*TxRenderCreateCursorReq)(unsafe.Pointer(req)).Fy = uint16(y)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return cid
}

func XXRenderCreateAnimCursor(tls *libc.TLS, dpy uintptr, ncursor int32, cursors uintptr) (r TCursor) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var cid, v1 TCursor
	var info, req uintptr
	var len1 int64
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _ = _BRlen, cid, info, len1, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateAnimCursor), uint64(m_sz_xRenderCreateAnimCursorReq))
	(*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateAnimCursor)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	cid = v1
	(*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Fcid = uint32(v1)
	len1 = int64(ncursor) * int64(m_sz_xAnimCursorElt) >> int32(2)
	if libc.Int64FromUint16((*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = len1
			(*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateAnimCursorReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	len1 <<= int64(2)
	libx11.X_XData32(tls, dpy, cursors, libc.Uint32FromInt64(len1))
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return cid
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

/* precompute the maximum size of batching request allowed */

func XXRenderFillRectangle(tls *libc.TLS, dpy uintptr, op int32, dst TPicture, color uintptr, x int32, y int32, width uint32, height uint32) {
	var info, rect, req, p1 uintptr
	_, _, _, _ = info, rect, req, p1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = (*TDisplay)(unsafe.Pointer(dpy)).Flast_req
	/* if same as previous request, with same drawable, batch requests */
	if libc.Int32FromUint8((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).FreqType) == (*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode && libc.Int32FromUint8((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).FrenderReqType) == int32(m_X_RenderFillRectangles) && libc.Int32FromUint8((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fop) == op && uint64((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fdst) == dst && libc.Int32FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fred) == libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(color)).Fred) && libc.Int32FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fgreen) == libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(color)).Fgreen) && libc.Int32FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fblue) == libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(color)).Fblue) && libc.Int32FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Falpha) == libc.Int32FromUint16((*TXRenderColor)(unsafe.Pointer(color)).Falpha) && (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(m_sz_xRectangle) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax && int64((*TDisplay)(unsafe.Pointer(dpy)).Fbufptr)-int64(req) < int64(libc.Int32FromInt32(m_sz_xRenderFillRectanglesReq)+libc.Int32FromInt32(m_FRCTSPERBATCH)*libc.Int32FromInt32(m_sz_xRectangle)) {
		p1 = req + 2
		*(*TCARD16)(unsafe.Pointer(p1)) = TCARD16(int32(*(*TCARD16)(unsafe.Pointer(p1))) + libc.Int32FromInt32(m_sz_xRectangle)>>libc.Int32FromInt32(2))
		rect = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(m_sz_xRectangle)
	} else {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderFillRectangles), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xRenderFillRectanglesReq)+libc.Int32FromInt32(m_sz_xRectangle)))
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderFillRectangles)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fred = (*TXRenderColor)(unsafe.Pointer(color)).Fred
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fgreen = (*TXRenderColor)(unsafe.Pointer(color)).Fgreen
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fblue = (*TXRenderColor)(unsafe.Pointer(color)).Fblue
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Falpha = (*TXRenderColor)(unsafe.Pointer(color)).Falpha
		rect = req + libc.UintptrFromInt32(1)*20
	}
	(*TxRectangle)(unsafe.Pointer(rect)).Fx = int16(x)
	(*TxRectangle)(unsafe.Pointer(rect)).Fy = int16(y)
	(*TxRectangle)(unsafe.Pointer(rect)).Fwidth = uint16(width)
	(*TxRectangle)(unsafe.Pointer(rect)).Fheight = uint16(height)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

/* precompute the maximum size of batching request allowed */

func XXRenderFillRectangles(tls *libc.TLS, dpy uintptr, op int32, dst TPicture, color uintptr, rectangles uintptr, n_rects int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var n int32
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _ = _BRlen, info, len1, n, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	for n_rects != 0 {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderFillRectangles), uint64(m_sz_xRenderFillRectanglesReq))
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderFillRectangles)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fred = (*TXRenderColor)(unsafe.Pointer(color)).Fred
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fgreen = (*TXRenderColor)(unsafe.Pointer(color)).Fgreen
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Fblue = (*TXRenderColor)(unsafe.Pointer(color)).Fblue
		(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Fcolor.Falpha = (*TXRenderColor)(unsafe.Pointer(color)).Falpha
		n = n_rects
		len1 = int64(n) << int32(1)
		if !((*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0) && len1 > libc.Int64FromUint32((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size-uint32((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength)) {
			n = libc.Int32FromUint32(((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size - uint32((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength)) >> int32(1))
			len1 = int64(n) << int32(1)
		}
		if libc.Int64FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
				_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
				(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength = uint16(0)
				*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
				libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
				*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
				libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
			} else {
				len1 = len1
				(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength) + len1)
			}
		} else {
			(*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderFillRectanglesReq)(unsafe.Pointer(req)).Flength) + len1)
		}
		len1 <<= int64(2) /* watch out for macros... */
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, rectangles, libc.Uint64FromInt64(len1))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
		} else {
			libx11.X_XSend(tls, dpy, rectangles, len1)
		}
		n_rects -= n
		rectangles += uintptr(n) * 8
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

const m_BC_BASE_MAX = 99
const m_BC_DIM_MAX = 2048
const m_BC_SCALE_MAX = 99
const m_BC_STRING_MAX = 1000
const m_CHARCLASS_NAME_MAX = 14
const m_CHAR_BIT = "__CHAR_BIT"
const m_CHAR_MAX = "SCHAR_MAX"
const m_CHAR_MIN = "SCHAR_MIN"
const m_CHILD_MAX = 40
const m_COLL_WEIGHTS_MAX = 10
const m_EXPR_NEST_MAX = 32
const m_GID_MAX = "UINT_MAX"
const m_INT_MAX = "__INT_MAX"
const m_INT_MIN = "__INT_MIN"
const m_IOV_MAX = 1024
const m_LINE_MAX = 2048
const m_LLONG_MAX = "__LLONG_MAX"
const m_LLONG_MIN = "__LLONG_MIN"
const m_LONG_BIT = "__LONG_BIT"
const m_LONG_MAX = "__LONG_MAX"
const m_LONG_MIN = "__LONG_MIN"
const m_MAX_CANON = 255
const m_MAX_INPUT = 255
const m_MB_LEN_MAX = 6
const m_MQ_PRIO_MAX = 64
const m_NAME_MAX = 255
const m_NGROUPS_MAX = 1023
const m_NL_ARGMAX = 4096
const m_NL_LANGMAX = 31
const m_NL_MSGMAX = 32767
const m_NL_NMAX = 1
const m_NL_SETMAX = 255
const m_NL_TEXTMAX = 2048
const m_OFF_MAX = "__OFF_MAX"
const m_OFF_MIN = "__OFF_MIN"
const m_OPEN_MAX = 64
const m_PASS_MAX = 128
const m_PATH_MAX = 1024
const m_PIPE_BUF = 512
const m_QUAD_MAX = "__QUAD_MAX"
const m_QUAD_MIN = "__QUAD_MIN"
const m_RE_DUP_MAX = 255
const m_SCHAR_MAX = "__SCHAR_MAX"
const m_SCHAR_MIN = "__SCHAR_MIN"
const m_SHRT_MAX = "__SHRT_MAX"
const m_SHRT_MIN = "__SHRT_MIN"
const m_SIZE_T_MAX = "__SIZE_T_MAX"
const m_SSIZE_MAX = "__SSIZE_MAX"
const m_UCHAR_MAX = "__UCHAR_MAX"
const m_UID_MAX = "UINT_MAX"
const m_UINT_MAX = "__UINT_MAX"
const m_ULLONG_MAX = "__ULLONG_MAX"
const m_ULONG_MAX = "__ULONG_MAX"
const m_UQUAD_MAX = "__UQUAD_MAX"
const m_USHRT_MAX = "__USHRT_MAX"
const m_WORD_BIT = "__WORD_BIT"
const m__POSIX2_BC_BASE_MAX = 99
const m__POSIX2_BC_DIM_MAX = 2048
const m__POSIX2_BC_SCALE_MAX = 99
const m__POSIX2_BC_STRING_MAX = 1000
const m__POSIX2_CHARCLASS_NAME_MAX = 14
const m__POSIX2_COLL_WEIGHTS_MAX = 2
const m__POSIX2_EQUIV_CLASS_MAX = 2
const m__POSIX2_EXPR_NEST_MAX = 32
const m__POSIX2_LINE_MAX = 2048
const m__POSIX2_RE_DUP_MAX = 255
const m__POSIX_AIO_LISTIO_MAX = 2
const m__POSIX_AIO_MAX = 1
const m__POSIX_ARG_MAX = 4096
const m__POSIX_CHILD_MAX = 25
const m__POSIX_CLOCKRES_MIN = 20000000
const m__POSIX_DELAYTIMER_MAX = 32
const m__POSIX_HOST_NAME_MAX = 255
const m__POSIX_LINK_MAX = 8
const m__POSIX_LOGIN_NAME_MAX = 9
const m__POSIX_MAX_CANON = 255
const m__POSIX_MAX_INPUT = 255
const m__POSIX_MQ_OPEN_MAX = 8
const m__POSIX_MQ_PRIO_MAX = 32
const m__POSIX_NAME_MAX = 14
const m__POSIX_NGROUPS_MAX = 8
const m__POSIX_OPEN_MAX = 20
const m__POSIX_PATH_MAX = 256
const m__POSIX_PIPE_BUF = 512
const m__POSIX_RE_DUP_MAX = "_POSIX2_RE_DUP_MAX"
const m__POSIX_RTSIG_MAX = 8
const m__POSIX_SEM_NSEMS_MAX = 256
const m__POSIX_SEM_VALUE_MAX = 32767
const m__POSIX_SIGQUEUE_MAX = 32
const m__POSIX_SSIZE_MAX = 32767
const m__POSIX_SS_REPL_MAX = 4
const m__POSIX_STREAM_MAX = 8
const m__POSIX_SYMLINK_MAX = 255
const m__POSIX_SYMLOOP_MAX = 8
const m__POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const m__POSIX_THREAD_KEYS_MAX = 128
const m__POSIX_THREAD_THREADS_MAX = 64
const m__POSIX_TIMER_MAX = 32
const m__POSIX_TRACE_EVENT_NAME_MAX = 30
const m__POSIX_TRACE_NAME_MAX = 8
const m__POSIX_TRACE_SYS_MAX = 8
const m__POSIX_TRACE_USER_EVENT_MAX = 32
const m__POSIX_TTY_NAME_MAX = 9
const m__POSIX_TZNAME_MAX = 6
const m__XOPEN_IOV_MAX = 16
const m__XOPEN_NAME_MAX = 255
const m__XOPEN_PATH_MAX = 1024
const m___INT_MAX1 = 2147483647

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)limits.h	8.2 (Berkeley) 1/4/94
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1991, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * This code is derived from software contributed to Berkeley by
 * Berkeley Software Design, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)cdefs.h	8.8 (Berkeley) 1/9/95
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1991, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * This code is derived from software contributed to Berkeley by
 * Berkeley Software Design, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)cdefs.h	8.8 (Berkeley) 1/9/95
 */

/*-
 * This file is in the public domain.
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)limits.h	8.3 (Berkeley) 1/4/94
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)syslimits.h	8.1 (Berkeley) 6/2/93
 */

/*
 * Do not add any new variables here.  (See the comment at the end of
 * the file for why.)
 */

/*
 * We leave the following values undefined to force applications to either
 * assume conservative values or call sysconf() to get the current value.
 *
 * HOST_NAME_MAX
 *
 * (We should do this for most of the values currently defined here,
 * but many programs are not prepared to deal with this yet.)
 */

func XXRenderQueryFilters(tls *libc.TLS, dpy uintptr, drawable TDrawable) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var filters, info, name, req, xri uintptr
	var i, l int32
	var nbytes, nbytesAlias, nbytesName, reply_left uint64
	var _ /* len at bp+32 */ int8
	var _ /* rep at bp+0 */ TxRenderQueryFiltersReply
	_, _, _, _, _, _, _, _, _, _, _ = filters, i, info, l, name, nbytes, nbytesAlias, nbytesName, reply_left, req, xri
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return libc.UintptrFromInt32(0)
	}
	if !(XXRenderQueryFormats(tls, dpy) != 0) {
		return libc.UintptrFromInt32(0)
	}
	xri = (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo
	if (*TXRenderInfo)(unsafe.Pointer(xri)).Fminor_version < int32(6) {
		return libc.UintptrFromInt32(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderQueryFilters), uint64(m_sz_xRenderQueryFiltersReq))
	(*TxRenderQueryFiltersReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderQueryFiltersReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderQueryFilters)
	(*TxRenderQueryFiltersReq)(unsafe.Pointer(req)).Fdrawable = uint32(drawable)
	if !(libx11.X_XReply(tls, dpy, bp, 0, m_xFalse) != 0) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return libc.UintptrFromInt32(0)
	}
	/*
	 * Limit each component of combined size to 1/4 the max, which is far
	 * more than they should ever possibly need.
	 */
	if (*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).Flength < libc.Uint32FromInt32(libc.Int32FromInt32(m___INT_MAX1)>>libc.Int32FromInt32(2)) && uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumFilters) < libc.Uint64FromInt32(libc.Int32FromInt32(m___INT_MAX1)/libc.Int32FromInt32(4))/libc.Uint64FromInt64(8) && uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases) < libc.Uint64FromInt32(libc.Int32FromInt32(m___INT_MAX1)/libc.Int32FromInt32(4))/libc.Uint64FromInt64(2) {
		/*
		 * Compute total number of bytes for filter names
		 */
		nbytes = uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).Flength) << int32(2)
		nbytesAlias = uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases * uint32(2))
		if (*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases&uint32(1) != 0 {
			nbytesAlias += uint64(2)
		}
		nbytesName = nbytes - nbytesAlias
		/*
		 * Allocate one giant block for the whole data structure
		 */
		filters = libc.Xmalloc(tls, libc.Uint64FromInt64(32)+uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumFilters)*libc.Uint64FromInt64(8)+uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases)*libc.Uint64FromInt64(2)+nbytesName)
	} else {
		filters = libc.UintptrFromInt32(0)
	}
	if !(filters != 0) {
		libx11.X_XEatDataWords(tls, dpy, uint64((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).Flength))
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return libc.UintptrFromInt32(0)
	}
	/*
	 * Layout:
	 *	XFilters
	 *	numFilters  char * pointers to filter names
	 *	numAliases  short alias values
	 *	nbytesName  char strings
	 */
	(*TXFilters)(unsafe.Pointer(filters)).Fnfilter = libc.Int32FromUint32((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumFilters)
	(*TXFilters)(unsafe.Pointer(filters)).Fnalias = libc.Int32FromUint32((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases)
	(*TXFilters)(unsafe.Pointer(filters)).Ffilter = filters + libc.UintptrFromInt32(1)*32
	(*TXFilters)(unsafe.Pointer(filters)).Falias = (*TXFilters)(unsafe.Pointer(filters)).Ffilter + uintptr((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumFilters)*8
	name = (*TXFilters)(unsafe.Pointer(filters)).Falias + uintptr((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases)*2
	/*
	 * Read the filter aliases
	 */
	libx11.X_XReadPad(tls, dpy, (*TXFilters)(unsafe.Pointer(filters)).Falias, libc.Int64FromUint32(libc.Uint32FromInt32(2)*(*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases))
	reply_left = uint64(uint32(8) + (*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).Flength - uint32(2)*(*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases)
	/*
	 * Read the filter names
	 */
	i = 0
	for {
		if !(libc.Uint32FromInt32(i) < (*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumFilters) {
			break
		}
		libx11.X_XRead(tls, dpy, bp+32, int64(1))
		reply_left--
		l = int32(*(*int8)(unsafe.Pointer(bp + 32))) & int32(0xff)
		if libc.Uint64FromInt32(l)+uint64(1) > nbytesName {
			libx11.X_XEatDataWords(tls, dpy, reply_left)
			libc.Xfree(tls, filters)
			if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
			}
			if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
				(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
			}
			return libc.UintptrFromInt32(0)
		}
		nbytesName -= libc.Uint64FromInt32(l + int32(1))
		*(*uintptr)(unsafe.Pointer((*TXFilters)(unsafe.Pointer(filters)).Ffilter + uintptr(i)*8)) = name
		libx11.X_XRead(tls, dpy, name, int64(l))
		reply_left -= libc.Uint64FromInt32(l)
		*(*int8)(unsafe.Pointer(name + uintptr(l))) = int8('\000')
		name += uintptr(l + int32(1))
		goto _1
	_1:
		;
		i++
	}
	i = int32(int64(name) - int64((*TXFilters)(unsafe.Pointer(filters)).Falias+uintptr((*(*TxRenderQueryFiltersReply)(unsafe.Pointer(bp))).FnumAliases)*2))
	if i&int32(3) != 0 {
		libx11.X_XEatData(tls, dpy, libc.Uint64FromInt32(int32(4)-i&int32(3)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return filters
}

func XXRenderSetPictureFilter(tls *libc.TLS, dpy uintptr, picture TPicture, filter uintptr, params uintptr, nparams int32) {
	var info, req, p1 uintptr
	var nbytes int32
	_, _, _, _ = info, nbytes, req, p1
	info = XXRenderFindDisplay(tls, dpy)
	nbytes = libc.Int32FromUint64(libc.Xstrlen(tls, filter))
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderSetPictureFilter), uint64(m_sz_xRenderSetPictureFilterReq))
	(*TxRenderSetPictureFilterReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderSetPictureFilterReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderSetPictureFilter)
	(*TxRenderSetPictureFilterReq)(unsafe.Pointer(req)).Fpicture = uint32(picture)
	(*TxRenderSetPictureFilterReq)(unsafe.Pointer(req)).Fnbytes = libc.Uint16FromInt32(nbytes)
	p1 = req + 2
	*(*TCARD16)(unsafe.Pointer(p1)) = TCARD16(int32(*(*TCARD16)(unsafe.Pointer(p1))) + ((nbytes+libc.Int32FromInt32(3))>>libc.Int32FromInt32(2) + nparams))
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, filter, libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nbytes+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, filter, int64(nbytes))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nparams<<libc.Int32FromInt32(2)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, params, libc.Uint64FromInt32(nparams<<libc.Int32FromInt32(2)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nparams<<libc.Int32FromInt32(2)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, params, int64(nparams<<libc.Int32FromInt32(2)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

const m___INT_MAX2 = 0x7fffffff

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

func XXRenderCreateGlyphSet(tls *libc.TLS, dpy uintptr, format uintptr) (r TGlyphSet) {
	var gsid, v1 TGlyphSet
	var info, req uintptr
	_, _, _, _ = gsid, info, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateGlyphSet), uint64(m_sz_xRenderCreateGlyphSetReq))
	(*TxRenderCreateGlyphSetReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateGlyphSetReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateGlyphSet)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	gsid = v1
	(*TxRenderCreateGlyphSetReq)(unsafe.Pointer(req)).Fgsid = uint32(v1)
	(*TxRenderCreateGlyphSetReq)(unsafe.Pointer(req)).Fformat = uint32((*TXRenderPictFormat)(unsafe.Pointer(format)).Fid)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return gsid
}

func XXRenderReferenceGlyphSet(tls *libc.TLS, dpy uintptr, existing TGlyphSet) (r TGlyphSet) {
	var gsid, v1 TGlyphSet
	var info, req uintptr
	_, _, _, _ = gsid, info, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderReferenceGlyphSet), uint64(m_sz_xRenderReferenceGlyphSetReq))
	(*TxRenderReferenceGlyphSetReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderReferenceGlyphSetReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderReferenceGlyphSet)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	gsid = v1
	(*TxRenderReferenceGlyphSetReq)(unsafe.Pointer(req)).Fgsid = uint32(v1)
	(*TxRenderReferenceGlyphSetReq)(unsafe.Pointer(req)).Fexisting = uint32(existing)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return gsid
}

func XXRenderFreeGlyphSet(tls *libc.TLS, dpy uintptr, glyphset TGlyphSet) {
	var info, req uintptr
	_, _ = info, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderFreeGlyphSet), uint64(m_sz_xRenderFreeGlyphSetReq))
	(*TxRenderFreeGlyphSetReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderFreeGlyphSetReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderFreeGlyphSet)
	(*TxRenderFreeGlyphSetReq)(unsafe.Pointer(req)).Fglyphset = uint32(glyphset)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderAddGlyphs(tls *libc.TLS, dpy uintptr, glyphset TGlyphSet, gids uintptr, glyphs uintptr, nglyphs int32, images uintptr, nbyte_images int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _ = _BRlen, info, len1, req
	info = XXRenderFindDisplay(tls, dpy)
	if nbyte_images&int32(3) != 0 {
		nbyte_images += int32(4) - nbyte_images&int32(3)
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderAddGlyphs), uint64(m_sz_xRenderAddGlyphsReq))
	(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderAddGlyphs)
	(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Fglyphset = uint32(glyphset)
	(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Fnglyphs = libc.Uint32FromInt32(nglyphs)
	len1 = int64((nglyphs*(libc.Int32FromInt32(m_sz_xGlyphInfo)+libc.Int32FromInt32(4)) + nbyte_images) >> int32(2))
	if libc.Int64FromUint16((*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = len1
			(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderAddGlyphsReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	libx11.X_XData32(tls, dpy, gids, libc.Uint32FromInt32(nglyphs*libc.Int32FromInt32(4)))
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nglyphs*libc.Int32FromInt32(m_sz_xGlyphInfo)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, glyphs, libc.Uint64FromInt32(nglyphs*libc.Int32FromInt32(m_sz_xGlyphInfo)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nglyphs*libc.Int32FromInt32(m_sz_xGlyphInfo)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, glyphs, int64(nglyphs*libc.Int32FromInt32(m_sz_xGlyphInfo)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbyte_images) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, images, libc.Uint64FromInt32(nbyte_images))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nbyte_images+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, images, int64(nbyte_images))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderFreeGlyphs(tls *libc.TLS, dpy uintptr, glyphset TGlyphSet, gids uintptr, nglyphs int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _ = _BRlen, info, len1, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderFreeGlyphs), uint64(m_sz_xRenderFreeGlyphsReq))
	(*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderFreeGlyphs)
	(*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Fglyphset = uint32(glyphset)
	len1 = int64(nglyphs)
	if libc.Int64FromUint16((*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = len1
			(*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderFreeGlyphsReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	len1 <<= int64(2)
	libx11.X_XData32(tls, dpy, gids, libc.Uint32FromInt64(len1))
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeString8(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, glyphset TGlyphSet, xSrc int32, ySrc int32, xDst int32, yDst int32, string1 uintptr, nchar int32) {
	var elt, info, req, p2 uintptr
	var len1 int64
	var nbytes int32
	var v1 uint64
	_, _, _, _, _, _, _ = elt, info, len1, nbytes, req, v1, p2
	info = XXRenderFindDisplay(tls, dpy)
	if !(nchar != 0) {
		return
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCompositeGlyphs8), uint64(m_sz_xRenderCompositeGlyphs8Req))
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCompositeGlyphs8)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fdst = uint32(dst)
	if maskFormat != 0 {
		v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
	} else {
		v1 = uint64(0)
	}
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fglyphset = uint32(glyphset)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FySrc = int16(ySrc)
	/*
	 * xGlyphElt must be aligned on a 32-bit boundary; this is
	 * easily done by filling no more than 252 glyphs in each
	 * bucket
	 */
	len1 = int64(int32(m_sz_xGlyphElt)*((nchar+int32(252)-int32(1))/int32(252)) + nchar)
	p2 = req + 2
	*(*TCARD16)(unsafe.Pointer(p2)) = TCARD16(int64(*(*TCARD16)(unsafe.Pointer(p2))) + (len1+libc.Int64FromInt32(3))>>libc.Int32FromInt32(2)) /* convert to number of 32-bit words */
	/*
	 * If the entire request does not fit into the remaining space in the
	 * buffer, flush the buffer first.
	 */
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libx11.X_XFlush(tls, dpy)
	}
	for nchar > int32(252) {
		nbytes = libc.Int32FromInt32(252) + libc.Int32FromInt32(m_sz_xGlyphElt)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libx11.X_XFlush(tls, dpy)
		}
		elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(nbytes)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = uint8(252)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
		xDst = 0
		yDst = 0
		libc.Xmemcpy(tls, elt+libc.UintptrFromInt32(1)*8, string1, uint64(252))
		nchar = nchar - int32(252)
		string1 += uintptr(252)
	}
	if nchar != 0 {
		nbytes = (nchar + int32(m_sz_xGlyphElt) + int32(3)) & ^libc.Int32FromInt32(3)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libx11.X_XFlush(tls, dpy)
		}
		elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(nbytes)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = libc.Uint8FromInt32(nchar)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
		libc.Xmemcpy(tls, elt+libc.UintptrFromInt32(1)*8, string1, libc.Uint64FromInt32(nchar))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeString16(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, glyphset TGlyphSet, xSrc int32, ySrc int32, xDst int32, yDst int32, string1 uintptr, nchar int32) {
	var elt, info, req, p2 uintptr
	var len1 int64
	var nbytes int32
	var v1 uint64
	_, _, _, _, _, _, _ = elt, info, len1, nbytes, req, v1, p2
	info = XXRenderFindDisplay(tls, dpy)
	if !(nchar != 0) {
		return
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCompositeGlyphs16), uint64(m_sz_xRenderCompositeGlyphs16Req))
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCompositeGlyphs16)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fdst = uint32(dst)
	if maskFormat != 0 {
		v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
	} else {
		v1 = uint64(0)
	}
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fglyphset = uint32(glyphset)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FySrc = int16(ySrc)
	len1 = int64(int32(m_sz_xGlyphElt)*((nchar+int32(254)-int32(1))/int32(254)) + nchar*int32(2))
	p2 = req + 2
	*(*TCARD16)(unsafe.Pointer(p2)) = TCARD16(int64(*(*TCARD16)(unsafe.Pointer(p2))) + (len1+libc.Int64FromInt32(3))>>libc.Int32FromInt32(2)) /* convert to number of 32-bit words */
	/*
	 * If the entire request does not fit into the remaining space in the
	 * buffer, flush the buffer first.
	 */
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libx11.X_XFlush(tls, dpy)
	}
	for nchar > int32(254) {
		nbytes = libc.Int32FromInt32(254)*libc.Int32FromInt32(2) + libc.Int32FromInt32(m_sz_xGlyphElt)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libx11.X_XFlush(tls, dpy)
		}
		elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(nbytes)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = uint8(254)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
		xDst = 0
		yDst = 0
		libc.Xmemcpy(tls, elt+libc.UintptrFromInt32(1)*8, string1, libc.Uint64FromInt32(libc.Int32FromInt32(254)*libc.Int32FromInt32(2)))
		nchar = nchar - int32(254)
		string1 += uintptr(254) * 2
	}
	if nchar != 0 {
		nbytes = (nchar*int32(2) + int32(m_sz_xGlyphElt) + int32(3)) & ^libc.Int32FromInt32(3)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libx11.X_XFlush(tls, dpy)
		}
		elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(nbytes)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = libc.Uint8FromInt32(nchar)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
		libc.Xmemcpy(tls, elt+libc.UintptrFromInt32(1)*8, string1, libc.Uint64FromInt32(nchar*int32(2)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeString32(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, glyphset TGlyphSet, xSrc int32, ySrc int32, xDst int32, yDst int32, string1 uintptr, nchar int32) {
	var elt, info, req, p2 uintptr
	var len1 int64
	var nbytes int32
	var v1 uint64
	_, _, _, _, _, _, _ = elt, info, len1, nbytes, req, v1, p2
	info = XXRenderFindDisplay(tls, dpy)
	if !(nchar != 0) {
		return
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCompositeGlyphs32), uint64(m_sz_xRenderCompositeGlyphs32Req))
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCompositeGlyphs32)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fdst = uint32(dst)
	if maskFormat != 0 {
		v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
	} else {
		v1 = uint64(0)
	}
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fglyphset = uint32(glyphset)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FySrc = int16(ySrc)
	len1 = int64(int32(m_sz_xGlyphElt)*((nchar+int32(254)-int32(1))/int32(254)) + nchar*int32(4))
	p2 = req + 2
	*(*TCARD16)(unsafe.Pointer(p2)) = TCARD16(int64(*(*TCARD16)(unsafe.Pointer(p2))) + (len1+libc.Int64FromInt32(3))>>libc.Int32FromInt32(2)) /* convert to number of 32-bit words */
	/*
	 * If the entire request does not fit into the remaining space in the
	 * buffer, flush the buffer first.
	 */
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libx11.X_XFlush(tls, dpy)
	}
	for nchar > int32(254) {
		nbytes = libc.Int32FromInt32(254)*libc.Int32FromInt32(4) + libc.Int32FromInt32(m_sz_xGlyphElt)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libx11.X_XFlush(tls, dpy)
		}
		elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(nbytes)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = uint8(254)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
		xDst = 0
		yDst = 0
		libc.Xmemcpy(tls, elt+libc.UintptrFromInt32(1)*8, string1, libc.Uint64FromInt32(libc.Int32FromInt32(254)*libc.Int32FromInt32(4)))
		nchar = nchar - int32(254)
		string1 += uintptr(254) * 4
	}
	if nchar != 0 {
		nbytes = nchar*int32(4) + int32(m_sz_xGlyphElt)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nbytes) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libx11.X_XFlush(tls, dpy)
		}
		elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
		libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(nbytes))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(nbytes)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = libc.Uint8FromInt32(nchar)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
		(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
		libc.Xmemcpy(tls, elt+libc.UintptrFromInt32(1)*8, string1, libc.Uint64FromInt32(nchar*int32(4)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeText8(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, xDst int32, yDst int32, elts uintptr, nelt int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var chars, elt, info, req, p3 uintptr
	var elen, len1 int64
	var i, nchars, this_chars, v5 int32
	var v1 uint64
	var _ /* glyphset at bp+0 */ TGlyphSet
	_, _, _, _, _, _, _, _, _, _, _, _ = chars, elen, elt, i, info, len1, nchars, req, this_chars, v1, v5, p3
	info = XXRenderFindDisplay(tls, dpy)
	if !(nelt != 0) {
		return
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCompositeGlyphs8), uint64(m_sz_xRenderCompositeGlyphs8Req))
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCompositeGlyphs8)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fdst = uint32(dst)
	if maskFormat != 0 {
		v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
	} else {
		v1 = uint64(0)
	}
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).Fglyphset = uint32((*(*TXGlyphElt8)(unsafe.Pointer(elts))).Fglyphset)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
	(*TxRenderCompositeGlyphs8Req)(unsafe.Pointer(req)).FySrc = int16(ySrc)
	/*
	 * Compute the space necessary
	 */
	len1 = 0
	*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt8)(unsafe.Pointer(elts))).Fglyphset
	i = 0
	for {
		if !(i < nelt) {
			break
		}
		/*
		 * Check for glyphset change
		 */
		if (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset != *(*TGlyphSet)(unsafe.Pointer(bp)) {
			*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset
			len1 += int64((libc.Int32FromInt32(m_sz_xGlyphElt) + libc.Int32FromInt32(4)) >> libc.Int32FromInt32(2))
		}
		nchars = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fnchars
		/*
		 * xGlyphElt must be aligned on a 32-bit boundary; this is
		 * easily done by filling no more than 252 glyphs in each
		 * bucket
		 */
		elen = int64(int32(m_sz_xGlyphElt)*((nchars+int32(252)-int32(1))/int32(252)) + nchars)
		len1 += (elen + int64(3)) >> int32(2)
		goto _2
	_2:
		;
		i++
	}
	p3 = req + 2
	*(*TCARD16)(unsafe.Pointer(p3)) = TCARD16(int64(*(*TCARD16)(unsafe.Pointer(p3))) + len1)
	/*
	 * Send the glyphs
	 */
	*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt8)(unsafe.Pointer(elts))).Fglyphset
	i = 0
	for {
		if !(i < nelt) {
			break
		}
		/*
		 * Switch glyphsets
		 */
		if (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset != *(*TGlyphSet)(unsafe.Pointer(bp)) {
			*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(libc.Int32FromInt32(m_sz_xGlyphElt)) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libx11.X_XFlush(tls, dpy)
			}
			elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
			libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xGlyphElt)))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Int32FromInt32(m_sz_xGlyphElt))
			(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = uint8(0xff)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = 0
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = 0
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		}
		nchars = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fnchars
		xDst = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).FxOff
		yDst = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).FyOff
		chars = (*(*TXGlyphElt8)(unsafe.Pointer(elts + uintptr(i)*32))).Fchars
		for nchars != 0 {
			if nchars > int32(252) {
				v5 = int32(252)
			} else {
				v5 = nchars
			}
			this_chars = v5
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(libc.Int32FromInt32(m_sz_xGlyphElt)) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libx11.X_XFlush(tls, dpy)
			}
			elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
			libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xGlyphElt)))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Int32FromInt32(m_sz_xGlyphElt))
			(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = libc.Uint8FromInt32(this_chars)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
			xDst = 0
			yDst = 0
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(this_chars) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, chars, libc.Uint64FromInt32(this_chars))
				*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(this_chars+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
			} else {
				libx11.X_XSend(tls, dpy, chars, int64(this_chars))
			}
			nchars -= this_chars
			chars += uintptr(this_chars)
		}
		goto _4
	_4:
		;
		i++
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeText16(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, xDst int32, yDst int32, elts uintptr, nelt int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var chars, elt, info, req, p3 uintptr
	var elen, len1 int64
	var i, nchars, this_bytes, this_chars, v5 int32
	var v1 uint64
	var _ /* glyphset at bp+0 */ TGlyphSet
	_, _, _, _, _, _, _, _, _, _, _, _, _ = chars, elen, elt, i, info, len1, nchars, req, this_bytes, this_chars, v1, v5, p3
	info = XXRenderFindDisplay(tls, dpy)
	if !(nelt != 0) {
		return
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCompositeGlyphs16), uint64(m_sz_xRenderCompositeGlyphs16Req))
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCompositeGlyphs16)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).Fdst = uint32(dst)
	if maskFormat != 0 {
		v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
	} else {
		v1 = uint64(0)
	}
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).Fglyphset = uint32((*(*TXGlyphElt16)(unsafe.Pointer(elts))).Fglyphset)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
	(*TxRenderCompositeGlyphs16Req)(unsafe.Pointer(req)).FySrc = int16(ySrc)
	/*
	 * Compute the space necessary
	 */
	len1 = 0
	*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt16)(unsafe.Pointer(elts))).Fglyphset
	i = 0
	for {
		if !(i < nelt) {
			break
		}
		/*
		 * Check for glyphset change
		 */
		if (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset != *(*TGlyphSet)(unsafe.Pointer(bp)) {
			*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset
			len1 += int64((libc.Int32FromInt32(m_sz_xGlyphElt) + libc.Int32FromInt32(4)) >> libc.Int32FromInt32(2))
		}
		nchars = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fnchars
		/*
		 * xGlyphElt must be aligned on a 32-bit boundary; this is
		 * easily done by filling no more than 254 glyphs in each
		 * bucket
		 */
		elen = int64(int32(m_sz_xGlyphElt)*((nchars+int32(254)-int32(1))/int32(254)) + nchars*int32(2))
		len1 += (elen + int64(3)) >> int32(2)
		goto _2
	_2:
		;
		i++
	}
	p3 = req + 2
	*(*TCARD16)(unsafe.Pointer(p3)) = TCARD16(int64(*(*TCARD16)(unsafe.Pointer(p3))) + len1)
	*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt16)(unsafe.Pointer(elts))).Fglyphset
	i = 0
	for {
		if !(i < nelt) {
			break
		}
		/*
		 * Switch glyphsets
		 */
		if (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset != *(*TGlyphSet)(unsafe.Pointer(bp)) {
			*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(libc.Int32FromInt32(m_sz_xGlyphElt)) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libx11.X_XFlush(tls, dpy)
			}
			elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
			libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xGlyphElt)))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Int32FromInt32(m_sz_xGlyphElt))
			(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = uint8(0xff)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = 0
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = 0
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		}
		nchars = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fnchars
		xDst = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).FxOff
		yDst = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).FyOff
		chars = (*(*TXGlyphElt16)(unsafe.Pointer(elts + uintptr(i)*32))).Fchars
		for nchars != 0 {
			if nchars > int32(254) {
				v5 = int32(254)
			} else {
				v5 = nchars
			}
			this_chars = v5
			this_bytes = this_chars * int32(2)
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(libc.Int32FromInt32(m_sz_xGlyphElt)) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libx11.X_XFlush(tls, dpy)
			}
			elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
			libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xGlyphElt)))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Int32FromInt32(m_sz_xGlyphElt))
			(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = libc.Uint8FromInt32(this_chars)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
			xDst = 0
			yDst = 0
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(this_bytes) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, chars, libc.Uint64FromInt32(this_bytes))
				*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(this_bytes+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
			} else {
				libx11.X_XSend(tls, dpy, chars, int64(this_bytes))
			}
			nchars -= this_chars
			chars += uintptr(this_chars) * 2
		}
		goto _4
	_4:
		;
		i++
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeText32(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, xDst int32, yDst int32, elts uintptr, nelt int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var chars, elt, info, req, p3 uintptr
	var elen, len1 int64
	var i, nchars, this_bytes, this_chars, v5 int32
	var v1 uint64
	var _ /* glyphset at bp+0 */ TGlyphSet
	_, _, _, _, _, _, _, _, _, _, _, _, _ = chars, elen, elt, i, info, len1, nchars, req, this_bytes, this_chars, v1, v5, p3
	info = XXRenderFindDisplay(tls, dpy)
	if !(nelt != 0) {
		return
	}
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCompositeGlyphs32), uint64(m_sz_xRenderCompositeGlyphs32Req))
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCompositeGlyphs32)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).Fsrc = uint32(src)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).Fdst = uint32(dst)
	if maskFormat != 0 {
		v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
	} else {
		v1 = uint64(0)
	}
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).Fglyphset = uint32((*(*TXGlyphElt32)(unsafe.Pointer(elts))).Fglyphset)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
	(*TxRenderCompositeGlyphs32Req)(unsafe.Pointer(req)).FySrc = int16(ySrc)
	/*
	 * Compute the space necessary
	 */
	len1 = 0
	*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt32)(unsafe.Pointer(elts))).Fglyphset
	i = 0
	for {
		if !(i < nelt) {
			break
		}
		/*
		 * Check for glyphset change
		 */
		if (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset != *(*TGlyphSet)(unsafe.Pointer(bp)) {
			*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset
			len1 += int64((libc.Int32FromInt32(m_sz_xGlyphElt) + libc.Int32FromInt32(4)) >> libc.Int32FromInt32(2))
		}
		nchars = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fnchars
		elen = int64(int32(m_sz_xGlyphElt)*((nchars+int32(254)-int32(1))/int32(254)) + nchars*int32(4))
		len1 += (elen + int64(3)) >> int32(2)
		goto _2
	_2:
		;
		i++
	}
	p3 = req + 2
	*(*TCARD16)(unsafe.Pointer(p3)) = TCARD16(int64(*(*TCARD16)(unsafe.Pointer(p3))) + len1)
	*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt32)(unsafe.Pointer(elts))).Fglyphset
	i = 0
	for {
		if !(i < nelt) {
			break
		}
		/*
		 * Switch glyphsets
		 */
		if (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset != *(*TGlyphSet)(unsafe.Pointer(bp)) {
			*(*TGlyphSet)(unsafe.Pointer(bp)) = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fglyphset
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(libc.Int32FromInt32(m_sz_xGlyphElt)) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libx11.X_XFlush(tls, dpy)
			}
			elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
			libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xGlyphElt)))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Int32FromInt32(m_sz_xGlyphElt))
			(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = uint8(0xff)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = 0
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = 0
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		}
		nchars = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fnchars
		xDst = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).FxOff
		yDst = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).FyOff
		chars = (*(*TXGlyphElt32)(unsafe.Pointer(elts + uintptr(i)*32))).Fchars
		for nchars != 0 {
			if nchars > int32(254) {
				v5 = int32(254)
			} else {
				v5 = nchars
			}
			this_chars = v5
			this_bytes = this_chars * int32(4)
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(libc.Int32FromInt32(m_sz_xGlyphElt)) > (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libx11.X_XFlush(tls, dpy)
			}
			elt = (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr
			libc.Xmemset(tls, elt, int32('\000'), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xGlyphElt)))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Int32FromInt32(m_sz_xGlyphElt))
			(*TxGlyphElt)(unsafe.Pointer(elt)).Flen1 = libc.Uint8FromInt32(this_chars)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltax = int16(xDst)
			(*TxGlyphElt)(unsafe.Pointer(elt)).Fdeltay = int16(yDst)
			xDst = 0
			yDst = 0
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(this_bytes) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
				libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, chars, libc.Uint64FromInt32(this_bytes))
				*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(this_bytes+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
			} else {
				libx11.X_XSend(tls, dpy, chars, int64(this_bytes))
			}
			nchars -= this_chars
			chars += uintptr(this_chars) * 4
		}
		goto _4
	_4:
		;
		i++
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

const m_FALSE = 0
const m_MAXSHORT = 32767
const m_NUMPTSTOBUFFER = 200
const m_TRUE = 1

type T_XRegion = struct {
	Fsize     int64
	FnumRects int64
	Frects    uintptr
	Fextents  TBOX
}

type TBox = struct {
	Fx1 int16
	Fx2 int16
	Fy1 int16
	Fy2 int16
}

type TBOX = struct {
	Fx1 int16
	Fx2 int16
	Fy1 int16
	Fy2 int16
}

type TBoxRec = struct {
	Fx1 int16
	Fx2 int16
	Fy1 int16
	Fy2 int16
}

type TBoxPtr = uintptr

type TRECTANGLE = struct {
	Fx      int16
	Fy      int16
	Fwidth  int16
	Fheight int16
}

type TRectangleRec = struct {
	Fx      int16
	Fy      int16
	Fwidth  int16
	Fheight int16
}

type TRectanglePtr = uintptr

type TREGION = struct {
	Fsize     int64
	FnumRects int64
	Frects    uintptr
	Fextents  TBOX
}

type TPOINTBLOCK = struct {
	Fpts  [200]TXPoint
	Fnext uintptr
}

type T_POINTBLOCK = TPOINTBLOCK

func __XRenderProcessPictureAttributes(tls *libc.TLS, dpy uintptr, req uintptr, valuemask uint64, attributes uintptr) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var nvalues, v15 uint32
	var value, v1, v10, v11, v12, v13, v2, v3, v4, v5, v6, v7, v8, v9, p14 uintptr
	var _ /* values at bp+0 */ [32]uint64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = nvalues, value, v1, v10, v11, v12, v13, v15, v2, v3, v4, v5, v6, v7, v8, v9, p14
	value = bp
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(0)) != 0 {
		v1 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v1)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Frepeat)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1)) != 0 {
		v2 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v2)) = (*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Falpha_map
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(2)) != 0 {
		v3 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v3)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Falpha_x_origin)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(3)) != 0 {
		v4 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v4)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Falpha_y_origin)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(4)) != 0 {
		v5 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v5)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fclip_x_origin)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(5)) != 0 {
		v6 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v6)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fclip_y_origin)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(6)) != 0 {
		v7 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v7)) = (*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fclip_mask
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(7)) != 0 {
		v8 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v8)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fgraphics_exposures)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(8)) != 0 {
		v9 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v9)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fsubwindow_mode)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(9)) != 0 {
		v10 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v10)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fpoly_edge)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)) != 0 {
		v11 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v11)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fpoly_mode)
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(11)) != 0 {
		v12 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v12)) = (*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fdither
	}
	if valuemask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(12)) != 0 {
		v13 = value
		value += 8
		*(*uint64)(unsafe.Pointer(v13)) = libc.Uint64FromInt32((*TXRenderPictureAttributes)(unsafe.Pointer(attributes)).Fcomponent_alpha)
	}
	p14 = req + 2
	v15 = libc.Uint32FromInt64((int64(value) - t__predefined_ptrdiff_t(bp)) / 8)
	nvalues = v15
	*(*TCARD16)(unsafe.Pointer(p14)) = TCARD16(uint32(*(*TCARD16)(unsafe.Pointer(p14))) + v15)
	nvalues <<= uint32(2) /* watch out for macros... */
	libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt64(libc.Int64FromUint32(nvalues)))
}

func XXRenderCreatePicture(tls *libc.TLS, dpy uintptr, drawable TDrawable, format uintptr, valuemask uint64, attributes uintptr) (r TPicture) {
	var info, req uintptr
	var pid, v1 TPicture
	var v2 TCARD32
	_, _, _, _, _ = info, pid, req, v1, v2
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreatePicture), uint64(m_sz_xRenderCreatePictureReq))
	(*TxRenderCreatePictureReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreatePictureReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreatePicture)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	pid = v1
	(*TxRenderCreatePictureReq)(unsafe.Pointer(req)).Fpid = uint32(v1)
	(*TxRenderCreatePictureReq)(unsafe.Pointer(req)).Fdrawable = uint32(drawable)
	(*TxRenderCreatePictureReq)(unsafe.Pointer(req)).Fformat = uint32((*TXRenderPictFormat)(unsafe.Pointer(format)).Fid)
	v2 = uint32(valuemask)
	(*TxRenderCreatePictureReq)(unsafe.Pointer(req)).Fmask = v2
	if v2 != 0 {
		__XRenderProcessPictureAttributes(tls, dpy, req, valuemask, attributes)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return pid
}

func XXRenderChangePicture(tls *libc.TLS, dpy uintptr, picture TPicture, valuemask uint64, attributes uintptr) {
	var info, req uintptr
	_, _ = info, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderChangePicture), uint64(m_sz_xRenderChangePictureReq))
	(*TxRenderChangePictureReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderChangePictureReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderChangePicture)
	(*TxRenderChangePictureReq)(unsafe.Pointer(req)).Fpicture = uint32(picture)
	(*TxRenderChangePictureReq)(unsafe.Pointer(req)).Fmask = uint32(valuemask)
	__XRenderProcessPictureAttributes(tls, dpy, req, valuemask, attributes)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func __XRenderSetPictureClipRectangles(tls *libc.TLS, dpy uintptr, info uintptr, picture TPicture, xOrigin int32, yOrigin int32, rects uintptr, n int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var len1 int64
	var req uintptr
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _ = _BRlen, len1, req
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderSetPictureClipRectangles), uint64(m_sz_xRenderSetPictureClipRectanglesReq))
	(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderSetPictureClipRectangles)
	(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Fpicture = uint32(picture)
	(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).FxOrigin = int16(xOrigin)
	(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).FyOrigin = int16(yOrigin)
	len1 = int64(n) << int32(1)
	if libc.Int64FromUint16((*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = int64(1)
			(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderSetPictureClipRectanglesReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	len1 <<= int64(2)
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, rects, libc.Uint64FromInt64(len1))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, rects, len1)
	}
}

func XXRenderSetPictureClipRectangles(tls *libc.TLS, dpy uintptr, picture TPicture, xOrigin int32, yOrigin int32, rects uintptr, n int32) {
	var info uintptr
	_ = info
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	__XRenderSetPictureClipRectangles(tls, dpy, info, picture, xOrigin, yOrigin, rects, n)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderSetPictureClipRegion(tls *libc.TLS, dpy uintptr, picture TPicture, r TRegion) {
	var i, v3 int32
	var info, pb, pr, xr, v1 uintptr
	var total uint64
	_, _, _, _, _, _, _, _ = i, info, pb, pr, total, xr, v1, v3
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	total = libc.Uint64FromInt64((*T_XRegion)(unsafe.Pointer(r)).FnumRects) * uint64(8)
	v1 = libx11.X_XAllocTemp(tls, dpy, total)
	xr = v1
	if v1 != 0 {
		pr = xr
		pb = (*T_XRegion)(unsafe.Pointer(r)).Frects
		i = int32((*T_XRegion)(unsafe.Pointer(r)).FnumRects)
		for {
			i--
			v3 = i
			if !(v3 >= 0) {
				break
			}
			(*TXRectangle)(unsafe.Pointer(pr)).Fx = (*TBOX)(unsafe.Pointer(pb)).Fx1
			(*TXRectangle)(unsafe.Pointer(pr)).Fy = (*TBOX)(unsafe.Pointer(pb)).Fy1
			(*TXRectangle)(unsafe.Pointer(pr)).Fwidth = libc.Uint16FromInt32(int32((*TBOX)(unsafe.Pointer(pb)).Fx2) - int32((*TBOX)(unsafe.Pointer(pb)).Fx1))
			(*TXRectangle)(unsafe.Pointer(pr)).Fheight = libc.Uint16FromInt32(int32((*TBOX)(unsafe.Pointer(pb)).Fy2) - int32((*TBOX)(unsafe.Pointer(pb)).Fy1))
			goto _2
		_2:
			;
			pr += 8
			pb += 8
		}
	}
	if xr != 0 || !((*T_XRegion)(unsafe.Pointer(r)).FnumRects != 0) {
		__XRenderSetPictureClipRectangles(tls, dpy, info, picture, 0, 0, xr, int32((*T_XRegion)(unsafe.Pointer(r)).FnumRects))
	}
	if xr != 0 {
		libx11.X_XFreeTemp(tls, dpy, xr, total)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderSetPictureTransform(tls *libc.TLS, dpy uintptr, picture TPicture, transform uintptr) {
	var info, req uintptr
	_, _ = info, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderSetPictureTransform), uint64(m_sz_xRenderSetPictureTransformReq))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderSetPictureTransform)
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Fpicture = uint32(picture)
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix11 = *(*TXFixed)(unsafe.Pointer(transform))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix12 = *(*TXFixed)(unsafe.Pointer(transform + 1*4))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix13 = *(*TXFixed)(unsafe.Pointer(transform + 2*4))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix21 = *(*TXFixed)(unsafe.Pointer(transform + 1*12))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix22 = *(*TXFixed)(unsafe.Pointer(transform + 1*12 + 1*4))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix23 = *(*TXFixed)(unsafe.Pointer(transform + 1*12 + 2*4))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix31 = *(*TXFixed)(unsafe.Pointer(transform + 2*12))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix32 = *(*TXFixed)(unsafe.Pointer(transform + 2*12 + 1*4))
	(*TxRenderSetPictureTransformReq)(unsafe.Pointer(req)).Ftransform.Fmatrix33 = *(*TXFixed)(unsafe.Pointer(transform + 2*12 + 2*4))
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderFreePicture(tls *libc.TLS, dpy uintptr, picture TPicture) {
	var info, req uintptr
	_, _ = info, req
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderFreePicture), uint64(m_sz_xRenderFreePictureReq))
	(*TxRenderFreePictureReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderFreePictureReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderFreePicture)
	(*TxRenderFreePictureReq)(unsafe.Pointer(req)).Fpicture = uint32(picture)
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCreateSolidFill(tls *libc.TLS, dpy uintptr, color uintptr) (r TPicture) {
	var info, req uintptr
	var pid, v1 TPicture
	_, _, _, _ = info, pid, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateSolidFill), uint64(m_sz_xRenderCreateSolidFillReq))
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateSolidFill)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	pid = v1
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).Fpid = uint32(v1)
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).Fcolor.Fred = (*TXRenderColor)(unsafe.Pointer(color)).Fred
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).Fcolor.Fgreen = (*TXRenderColor)(unsafe.Pointer(color)).Fgreen
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).Fcolor.Fblue = (*TXRenderColor)(unsafe.Pointer(color)).Fblue
	(*TxRenderCreateSolidFillReq)(unsafe.Pointer(req)).Fcolor.Falpha = (*TXRenderColor)(unsafe.Pointer(color)).Falpha
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return pid
}

func XXRenderCreateLinearGradient(tls *libc.TLS, dpy uintptr, gradient uintptr, stops uintptr, colors uintptr, nStops int32) (r TPicture) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var pid, v1 TPicture
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _ = _BRlen, info, len1, pid, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateLinearGradient), uint64(m_sz_xRenderCreateLinearGradientReq))
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateLinearGradient)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	pid = v1
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Fpid = uint32(v1)
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Fp1.Fx = (*TXLinearGradient)(unsafe.Pointer(gradient)).Fp1.Fx
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Fp1.Fy = (*TXLinearGradient)(unsafe.Pointer(gradient)).Fp1.Fy
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Fp2.Fx = (*TXLinearGradient)(unsafe.Pointer(gradient)).Fp2.Fx
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Fp2.Fy = (*TXLinearGradient)(unsafe.Pointer(gradient)).Fp2.Fy
	(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).FnStops = libc.Uint32FromInt32(nStops)
	len1 = int64(nStops) * int64(3)
	if libc.Int64FromUint16((*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = int64(6)
			(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateLinearGradientReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nStops*libc.Int32FromInt32(4)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, stops, libc.Uint64FromInt32(nStops*libc.Int32FromInt32(4)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nStops*libc.Int32FromInt32(4)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, stops, int64(nStops*libc.Int32FromInt32(4)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nStops*libc.Int32FromInt32(8)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, colors, libc.Uint64FromInt32(nStops*libc.Int32FromInt32(8)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nStops*libc.Int32FromInt32(8)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, colors, int64(nStops*libc.Int32FromInt32(8)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return pid
}

func XXRenderCreateRadialGradient(tls *libc.TLS, dpy uintptr, gradient uintptr, stops uintptr, colors uintptr, nStops int32) (r TPicture) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var pid, v1 TPicture
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _ = _BRlen, info, len1, pid, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateRadialGradient), uint64(m_sz_xRenderCreateRadialGradientReq))
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateRadialGradient)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	pid = v1
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Fpid = uint32(v1)
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Finner.Fx = (*TXRadialGradient)(unsafe.Pointer(gradient)).Finner.Fx
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Finner.Fy = (*TXRadialGradient)(unsafe.Pointer(gradient)).Finner.Fy
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Fouter.Fx = (*TXRadialGradient)(unsafe.Pointer(gradient)).Fouter.Fx
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Fouter.Fy = (*TXRadialGradient)(unsafe.Pointer(gradient)).Fouter.Fy
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Finner_radius = (*TXRadialGradient)(unsafe.Pointer(gradient)).Finner.Fradius
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Fouter_radius = (*TXRadialGradient)(unsafe.Pointer(gradient)).Fouter.Fradius
	(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).FnStops = libc.Uint32FromInt32(nStops)
	len1 = int64(nStops) * int64(3)
	if libc.Int64FromUint16((*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = int64(6)
			(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateRadialGradientReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nStops*libc.Int32FromInt32(4)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, stops, libc.Uint64FromInt32(nStops*libc.Int32FromInt32(4)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nStops*libc.Int32FromInt32(4)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, stops, int64(nStops*libc.Int32FromInt32(4)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nStops*libc.Int32FromInt32(8)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, colors, libc.Uint64FromInt32(nStops*libc.Int32FromInt32(8)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nStops*libc.Int32FromInt32(8)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, colors, int64(nStops*libc.Int32FromInt32(8)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return pid
}

func XXRenderCreateConicalGradient(tls *libc.TLS, dpy uintptr, gradient uintptr, stops uintptr, colors uintptr, nStops int32) (r TPicture) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var pid, v1 TPicture
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _ = _BRlen, info, len1, pid, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return uint64(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderCreateConicalGradient), uint64(m_sz_xRenderCreateConicalGradientReq))
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderCreateConicalGradient)
	v1 = (*(*func(*libc.TLS, uintptr) TXID)(unsafe.Pointer(&struct{ uintptr }{(*struct {
		Fext_data            uintptr
		Fprivate1            uintptr
		Ffd                  int32
		Fprivate2            int32
		Fproto_major_version int32
		Fproto_minor_version int32
		Fvendor              uintptr
		Fprivate3            TXID
		Fprivate4            TXID
		Fprivate5            TXID
		Fprivate6            int32
		Fresource_alloc      uintptr
		Fbyte_order          int32
		Fbitmap_unit         int32
		Fbitmap_pad          int32
		Fbitmap_bit_order    int32
		Fnformats            int32
		Fpixmap_format       uintptr
		Fprivate8            int32
		Frelease             int32
		Fprivate9            uintptr
		Fprivate10           uintptr
		Fqlen                int32
		Flast_request_read   uint64
		Frequest             uint64
		Fprivate11           TXPointer
		Fprivate12           TXPointer
		Fprivate13           TXPointer
		Fprivate14           TXPointer
		Fmax_request_size    uint32
		Fdb                  uintptr
		Fprivate15           uintptr
		Fdisplay_name        uintptr
		Fdefault_screen      int32
		Fnscreens            int32
		Fscreens             uintptr
		Fmotion_buffer       uint64
		Fprivate16           uint64
		Fmin_keycode         int32
		Fmax_keycode         int32
		Fprivate17           TXPointer
		Fprivate18           TXPointer
		Fprivate19           int32
		Fxdefaults           uintptr
	})(unsafe.Pointer(dpy)).Fresource_alloc})))(tls, dpy)
	pid = v1
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Fpid = uint32(v1)
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Fcenter.Fx = (*TXConicalGradient)(unsafe.Pointer(gradient)).Fcenter.Fx
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Fcenter.Fy = (*TXConicalGradient)(unsafe.Pointer(gradient)).Fcenter.Fy
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Fangle = (*TXConicalGradient)(unsafe.Pointer(gradient)).Fangle
	(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).FnStops = libc.Uint32FromInt32(nStops)
	len1 = int64(nStops) * int64(3)
	if libc.Int64FromUint16((*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
			_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
			(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength = uint16(0)
			*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
			libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
			*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
			libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
		} else {
			len1 = int64(6)
			(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength) + len1)
		}
	} else {
		(*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderCreateConicalGradientReq)(unsafe.Pointer(req)).Flength) + len1)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nStops*libc.Int32FromInt32(4)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, stops, libc.Uint64FromInt32(nStops*libc.Int32FromInt32(4)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nStops*libc.Int32FromInt32(4)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, stops, int64(nStops*libc.Int32FromInt32(4)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(nStops*libc.Int32FromInt32(8)) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
		libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, colors, libc.Uint64FromInt32(nStops*libc.Int32FromInt32(8)))
		*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt32(nStops*libc.Int32FromInt32(8)+libc.Int32FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
	} else {
		libx11.X_XSend(tls, dpy, colors, int64(nStops*libc.Int32FromInt32(8)))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return pid
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

type TEdge = struct {
	Fedge      TXLineFixed
	Fcurrent_x TXFixed
	FclockWise int32
	Fnext      uintptr
	Fprev      uintptr
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

type T_Edge = TEdge

func _CompareEdge(tls *libc.TLS, o1 uintptr, o2 uintptr) (r int32) {
	var e1, e2 uintptr
	_, _ = e1, e2
	e1 = o1
	e2 = o2
	return (*TEdge)(unsafe.Pointer(e1)).Fedge.Fp1.Fy - (*TEdge)(unsafe.Pointer(e2)).Fedge.Fp1.Fy
}

func _XRenderComputeX(tls *libc.TLS, line uintptr, y TXFixed) (r TXFixed) {
	var dx, dy TXFixed
	var ex float64
	_, _, _ = dx, dy, ex
	dx = (*TXLineFixed)(unsafe.Pointer(line)).Fp2.Fx - (*TXLineFixed)(unsafe.Pointer(line)).Fp1.Fx
	ex = float64(float64(y-(*TXLineFixed)(unsafe.Pointer(line)).Fp1.Fy) * float64(dx))
	dy = (*TXLineFixed)(unsafe.Pointer(line)).Fp2.Fy - (*TXLineFixed)(unsafe.Pointer(line)).Fp1.Fy
	return (*TXLineFixed)(unsafe.Pointer(line)).Fp1.Fx + int32(ex/float64(dy))
}

func _XRenderComputeInverseSlope(tls *libc.TLS, l uintptr) (r float64) {
	return float64((*TXLineFixed)(unsafe.Pointer(l)).Fp2.Fx-(*TXLineFixed)(unsafe.Pointer(l)).Fp1.Fx) / libc.Float64FromInt32(65536) / (float64((*TXLineFixed)(unsafe.Pointer(l)).Fp2.Fy-(*TXLineFixed)(unsafe.Pointer(l)).Fp1.Fy) / libc.Float64FromInt32(65536))
}

func _XRenderComputeXIntercept(tls *libc.TLS, l uintptr, inverse_slope float64) (r float64) {
	return float64((*TXLineFixed)(unsafe.Pointer(l)).Fp1.Fx)/libc.Float64FromInt32(65536) - TXDouble(inverse_slope*(float64((*TXLineFixed)(unsafe.Pointer(l)).Fp1.Fy)/libc.Float64FromInt32(65536)))
}

func _XRenderComputeIntersect(tls *libc.TLS, l1 uintptr, l2 uintptr) (r TXFixed) {
	var b1, b2, m1, m2 float64
	_, _, _, _ = b1, b2, m1, m2
	/*
	 * x = m1y + b1
	 * x = m2y + b2
	 * m1y + b1 = m2y + b2
	 * y * (m1 - m2) = b2 - b1
	 * y = (b2 - b1) / (m1 - m2)
	 */
	m1 = _XRenderComputeInverseSlope(tls, l1)
	b1 = _XRenderComputeXIntercept(tls, l1, m1)
	m2 = _XRenderComputeInverseSlope(tls, l2)
	b2 = _XRenderComputeXIntercept(tls, l2, m2)
	return int32(float64((b2 - b1) / (m1 - m2) * libc.Float64FromInt32(65536)))
}

func _XRenderComputeTrapezoids(tls *libc.TLS, edges uintptr, nedges int32, winding int32, traps uintptr) (r int32) {
	var active, e, en, next, v6 uintptr
	var inactive, ntraps int32
	var intersect, next_y, y TXFixed
	var v7 bool
	_, _, _, _, _, _, _, _, _, _, _ = active, e, en, inactive, intersect, next, next_y, ntraps, y, v6, v7
	ntraps = 0
	libc.Xqsort(tls, edges, libc.Uint64FromInt32(nedges), uint64(40), __ccgo_fp(_CompareEdge))
	y = (*(*TEdge)(unsafe.Pointer(edges))).Fedge.Fp1.Fy
	active = libc.UintptrFromInt32(0)
	inactive = 0
	for active != 0 || inactive < nedges {
		/* insert new active edges into list */
		for inactive < nedges {
			e = edges + uintptr(inactive)*40
			if (*TEdge)(unsafe.Pointer(e)).Fedge.Fp1.Fy > y {
				break
			}
			/* move this edge into the active list */
			inactive++
			(*TEdge)(unsafe.Pointer(e)).Fnext = active
			(*TEdge)(unsafe.Pointer(e)).Fprev = libc.UintptrFromInt32(0)
			if active != 0 {
				(*TEdge)(unsafe.Pointer(active)).Fprev = e
			}
			active = e
		}
		/* compute x coordinates along this group */
		e = active
		for {
			if !(e != 0) {
				break
			}
			(*TEdge)(unsafe.Pointer(e)).Fcurrent_x = _XRenderComputeX(tls, e, y)
			goto _1
		_1:
			;
			e = (*TEdge)(unsafe.Pointer(e)).Fnext
		}
		/* sort active list */
		e = active
		for {
			if !(e != 0) {
				break
			}
			next = (*TEdge)(unsafe.Pointer(e)).Fnext
			/*
			 * Find one later in the list that belongs before the
			 * current one
			 */
			en = next
			for {
				if !(en != 0) {
					break
				}
				if (*TEdge)(unsafe.Pointer(en)).Fcurrent_x < (*TEdge)(unsafe.Pointer(e)).Fcurrent_x || (*TEdge)(unsafe.Pointer(en)).Fcurrent_x == (*TEdge)(unsafe.Pointer(e)).Fcurrent_x && (*TEdge)(unsafe.Pointer(en)).Fedge.Fp2.Fx < (*TEdge)(unsafe.Pointer(e)).Fedge.Fp2.Fx {
					/*
					 * insert en before e
					 *
					 * extract en
					 */
					(*TEdge)(unsafe.Pointer((*TEdge)(unsafe.Pointer(en)).Fprev)).Fnext = (*TEdge)(unsafe.Pointer(en)).Fnext
					if (*TEdge)(unsafe.Pointer(en)).Fnext != 0 {
						(*TEdge)(unsafe.Pointer((*TEdge)(unsafe.Pointer(en)).Fnext)).Fprev = (*TEdge)(unsafe.Pointer(en)).Fprev
					}
					/*
					 * insert en
					 */
					if (*TEdge)(unsafe.Pointer(e)).Fprev != 0 {
						(*TEdge)(unsafe.Pointer((*TEdge)(unsafe.Pointer(e)).Fprev)).Fnext = en
					} else {
						active = en
					}
					(*TEdge)(unsafe.Pointer(en)).Fprev = (*TEdge)(unsafe.Pointer(e)).Fprev
					(*TEdge)(unsafe.Pointer(e)).Fprev = en
					(*TEdge)(unsafe.Pointer(en)).Fnext = e
					/*
					 * start over at en
					 */
					next = en
					break
				}
				goto _3
			_3:
				;
				en = (*TEdge)(unsafe.Pointer(en)).Fnext
			}
			goto _2
		_2:
			;
			e = next
		}
		/* find next inflection point */
		next_y = (*TEdge)(unsafe.Pointer(active)).Fedge.Fp2.Fy
		e = active
		for {
			if !(e != 0) {
				break
			}
			if (*TEdge)(unsafe.Pointer(e)).Fedge.Fp2.Fy < next_y {
				next_y = (*TEdge)(unsafe.Pointer(e)).Fedge.Fp2.Fy
			}
			en = (*TEdge)(unsafe.Pointer(e)).Fnext
			/* check intersect */
			if en != 0 && (*TEdge)(unsafe.Pointer(e)).Fedge.Fp2.Fx > (*TEdge)(unsafe.Pointer(en)).Fedge.Fp2.Fx {
				intersect = _XRenderComputeIntersect(tls, e, (*TEdge)(unsafe.Pointer(e)).Fnext)
				/* make sure this point is below the actual intersection */
				intersect = intersect + int32(1)
				if intersect < next_y {
					next_y = intersect
				}
			}
			goto _4
		_4:
			;
			e = en
		}
		/* check next inactive point */
		if inactive < nedges && (*(*TEdge)(unsafe.Pointer(edges + uintptr(inactive)*40))).Fedge.Fp1.Fy < next_y {
			next_y = (*(*TEdge)(unsafe.Pointer(edges + uintptr(inactive)*40))).Fedge.Fp1.Fy
		}
		/* walk the list generating trapezoids */
		e = active
		for {
			if v7 = e != 0; v7 {
				v6 = (*TEdge)(unsafe.Pointer(e)).Fnext
				en = v6
			}
			if !(v7 && v6 != 0) {
				break
			}
			(*TXTrapezoid)(unsafe.Pointer(traps)).Ftop = y
			(*TXTrapezoid)(unsafe.Pointer(traps)).Fbottom = next_y
			(*TXTrapezoid)(unsafe.Pointer(traps)).Fleft = (*TEdge)(unsafe.Pointer(e)).Fedge
			(*TXTrapezoid)(unsafe.Pointer(traps)).Fright = (*TEdge)(unsafe.Pointer(en)).Fedge
			traps += 40
			ntraps++
			goto _5
		_5:
			;
			e = (*TEdge)(unsafe.Pointer(en)).Fnext
		}
		y = next_y
		/* delete inactive edges from list */
		e = active
		for {
			if !(e != 0) {
				break
			}
			next = (*TEdge)(unsafe.Pointer(e)).Fnext
			if (*TEdge)(unsafe.Pointer(e)).Fedge.Fp2.Fy <= y {
				if (*TEdge)(unsafe.Pointer(e)).Fprev != 0 {
					(*TEdge)(unsafe.Pointer((*TEdge)(unsafe.Pointer(e)).Fprev)).Fnext = (*TEdge)(unsafe.Pointer(e)).Fnext
				} else {
					active = (*TEdge)(unsafe.Pointer(e)).Fnext
				}
				if (*TEdge)(unsafe.Pointer(e)).Fnext != 0 {
					(*TEdge)(unsafe.Pointer((*TEdge)(unsafe.Pointer(e)).Fnext)).Fprev = (*TEdge)(unsafe.Pointer(e)).Fprev
				}
			}
			goto _8
		_8:
			;
			e = next
		}
	}
	return ntraps
}

func XXRenderCompositeDoublePoly(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, xDst int32, yDst int32, fpoints uintptr, npoints int32, winding int32) {
	var bottom, firstx, firsty, prevx, prevy, top, x, y TXFixed
	var edges, traps uintptr
	var i, nedges, ntraps int32
	_, _, _, _, _, _, _, _, _, _, _, _, _ = bottom, edges, firstx, firsty, i, nedges, ntraps, prevx, prevy, top, traps, x, y
	prevx = 0
	prevy = 0
	firstx = 0
	firsty = 0
	top = 0
	bottom = 0 /* GCCism */
	edges = libc.Xmalloc(tls, libc.Uint64FromInt32(npoints)*libc.Uint64FromInt64(40)+libc.Uint64FromInt32(npoints*npoints)*libc.Uint64FromInt64(40))
	if !(edges != 0) {
		return
	}
	traps = edges + uintptr(npoints)*40
	nedges = 0
	i = 0
	for {
		if !(i <= npoints) {
			break
		}
		if i == npoints {
			x = firstx
			y = firsty
		} else {
			x = int32(TXDouble((*(*TXPointDouble)(unsafe.Pointer(fpoints + uintptr(i)*16))).Fx * libc.Float64FromInt32(65536)))
			y = int32(TXDouble((*(*TXPointDouble)(unsafe.Pointer(fpoints + uintptr(i)*16))).Fy * libc.Float64FromInt32(65536)))
		}
		if i != 0 {
			if y < top {
				top = y
			} else {
				if y > bottom {
					bottom = y
				}
			}
			if prevy < y {
				(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp1.Fx = prevx
				(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp1.Fy = prevy
				(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp2.Fx = x
				(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp2.Fy = y
				(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).FclockWise = int32(m_True)
				nedges++
			} else {
				if prevy > y {
					(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp1.Fx = x
					(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp1.Fy = y
					(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp2.Fx = prevx
					(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).Fedge.Fp2.Fy = prevy
					(*(*TEdge)(unsafe.Pointer(edges + uintptr(nedges)*40))).FclockWise = m_False
					nedges++
				}
			}
			/* drop horizontal edges */
		} else {
			top = y
			bottom = y
			firstx = x
			firsty = y
		}
		prevx = x
		prevy = y
		goto _1
	_1:
		;
		i++
	}
	ntraps = _XRenderComputeTrapezoids(tls, edges, nedges, winding, traps)
	/* XXX adjust xSrc/xDst */
	XXRenderCompositeTrapezoids(tls, dpy, op, src, dst, maskFormat, xSrc, ySrc, traps, ntraps)
	libc.Xfree(tls, edges)
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

func XXRenderCompositeTrapezoids(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, traps uintptr, ntrap int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var max_req, v1, v2 uint64
	var n int32
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _, _, _ = _BRlen, info, len1, max_req, n, req, v1, v2
	info = XXRenderFindDisplay(tls, dpy)
	if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
		v1 = (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size
	} else {
		v1 = uint64((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size)
	}
	max_req = v1
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	for ntrap != 0 {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderTrapezoids), uint64(m_sz_xRenderTrapezoidsReq))
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderTrapezoids)
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Fsrc = uint32(src)
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
		if maskFormat != 0 {
			v2 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
		} else {
			v2 = uint64(0)
		}
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).FmaskFormat = uint32(v2)
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
		(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).FySrc = int16(ySrc)
		n = ntrap
		len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xTrapezoid)>>libc.Int32FromInt32(2))
		if libc.Uint64FromInt64(len1) > max_req-uint64((*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength) {
			n = libc.Int32FromUint64((max_req - uint64((*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength)) / libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xTrapezoid)>>libc.Int32FromInt32(2)))
			len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xTrapezoid)>>libc.Int32FromInt32(2))
		}
		if libc.Int64FromUint16((*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
				_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
				(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength = uint16(0)
				*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
				libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
				*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
				libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
			} else {
				len1 = len1
				(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength) + len1)
			}
		} else {
			(*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTrapezoidsReq)(unsafe.Pointer(req)).Flength) + len1)
		}
		len1 <<= int64(2)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, traps, libc.Uint64FromInt64(len1))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
		} else {
			libx11.X_XSend(tls, dpy, traps, len1)
		}
		ntrap -= n
		traps += uintptr(n) * 40
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

func XXRenderCompositeTriangles(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, triangles uintptr, ntriangle int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var n int32
	var v1 uint64
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _ = _BRlen, info, len1, n, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	for ntriangle != 0 {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderTriangles), uint64(m_sz_xRenderTrianglesReq))
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderTriangles)
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).Fsrc = uint32(src)
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
		if maskFormat != 0 {
			v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
		} else {
			v1 = uint64(0)
		}
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
		(*TxRenderTrianglesReq)(unsafe.Pointer(req)).FySrc = int16(ySrc)
		n = ntriangle
		len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xTriangle)>>libc.Int32FromInt32(2))
		if !((*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0) && len1 > libc.Int64FromUint32((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size-uint32((*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength)) {
			n = libc.Int32FromUint32(((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size - uint32((*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength)) / libc.Uint32FromInt32(libc.Int32FromInt32(m_sz_xTriangle)>>libc.Int32FromInt32(2)))
			len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xTriangle)>>libc.Int32FromInt32(2))
		}
		if libc.Int64FromUint16((*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
				_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
				(*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength = uint16(0)
				*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
				libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
				*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
				libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
			} else {
				len1 = len1
				(*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength) + len1)
			}
		} else {
			(*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTrianglesReq)(unsafe.Pointer(req)).Flength) + len1)
		}
		len1 <<= int64(2)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, triangles, libc.Uint64FromInt64(len1))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
		} else {
			libx11.X_XSend(tls, dpy, triangles, len1)
		}
		ntriangle -= n
		triangles += uintptr(n) * 24
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeTriStrip(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, points uintptr, npoint int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var info, req uintptr
	var len1 int64
	var n int32
	var v1 uint64
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _ = _BRlen, info, len1, n, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	for npoint > int32(2) {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderTriStrip), uint64(m_sz_xRenderTriStripReq))
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderTriStrip)
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).Fsrc = uint32(src)
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
		if maskFormat != 0 {
			v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
		} else {
			v1 = uint64(0)
		}
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
		(*TxRenderTriStripReq)(unsafe.Pointer(req)).FySrc = int16(ySrc)
		n = npoint
		len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xPointFixed)>>libc.Int32FromInt32(2))
		if !((*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0) && len1 > libc.Int64FromUint32((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size-uint32((*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength)) {
			n = libc.Int32FromUint32(((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size - uint32((*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength)) / libc.Uint32FromInt32(libc.Int32FromInt32(m_sz_xPointFixed)>>libc.Int32FromInt32(2)))
			len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xPointFixed)>>libc.Int32FromInt32(2))
		}
		if libc.Int64FromUint16((*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
				_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
				(*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength = uint16(0)
				*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
				libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
				*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
				libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
			} else {
				len1 = len1
				(*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength) + len1)
			}
		} else {
			(*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTriStripReq)(unsafe.Pointer(req)).Flength) + len1)
		}
		len1 <<= int64(2)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, points, libc.Uint64FromInt64(len1))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
		} else {
			libx11.X_XSend(tls, dpy, points, len1)
		}
		npoint -= n - int32(2)
		points += uintptr(n-libc.Int32FromInt32(2)) * 8
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

func XXRenderCompositeTriFan(tls *libc.TLS, dpy uintptr, op int32, src TPicture, dst TPicture, maskFormat uintptr, xSrc int32, ySrc int32, points uintptr, npoint int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _BRlen TCARD32
	var first, info, p, req uintptr
	var len1 int64
	var n int32
	var v1 uint64
	var _ /* _BRdat at bp+0 */ TCARD64
	_, _, _, _, _, _, _, _ = _BRlen, first, info, len1, n, p, req, v1
	info = XXRenderFindDisplay(tls, dpy)
	first = points
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	points += 8
	npoint--
	for npoint > int32(1) {
		req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderTriFan), libc.Uint64FromInt32(libc.Int32FromInt32(m_sz_xRenderTriFanReq)+libc.Int32FromInt32(m_sz_xPointFixed)))
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderTriFan)
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).Fop = libc.Uint8FromInt32(op)
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).Fsrc = uint32(src)
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).Fdst = uint32(dst)
		if maskFormat != 0 {
			v1 = (*TXRenderPictFormat)(unsafe.Pointer(maskFormat)).Fid
		} else {
			v1 = uint64(0)
		}
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).FmaskFormat = uint32(v1)
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).FxSrc = int16(xSrc)
		(*TxRenderTriFanReq)(unsafe.Pointer(req)).FySrc = int16(ySrc)
		p = req + libc.UintptrFromInt32(1)*24
		(*TxPointFixed)(unsafe.Pointer(p)).Fx = (*TXPointFixed)(unsafe.Pointer(first)).Fx
		(*TxPointFixed)(unsafe.Pointer(p)).Fy = (*TXPointFixed)(unsafe.Pointer(first)).Fy
		n = npoint
		len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xPointFixed)>>libc.Int32FromInt32(2))
		if !((*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0) && len1 > libc.Int64FromUint32((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size-uint32((*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength)) {
			n = libc.Int32FromUint32(((*TDisplay)(unsafe.Pointer(dpy)).Fmax_request_size - uint32((*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength)) / libc.Uint32FromInt32(libc.Int32FromInt32(m_sz_xPointFixed)>>libc.Int32FromInt32(2)))
			len1 = int64(n) * int64(libc.Int32FromInt32(m_sz_xPointFixed)>>libc.Int32FromInt32(2))
		}
		if libc.Int64FromUint16((*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength)+len1 > libc.Int64FromUint32(libc.Uint32FromInt32(65535)) {
			if (*TDisplay)(unsafe.Pointer(dpy)).Fbigreq_size != 0 {
				_BRlen = libc.Uint32FromInt32(libc.Int32FromUint16((*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength) - libc.Int32FromInt32(1))
				(*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength = uint16(0)
				*(*TCARD64)(unsafe.Pointer(bp)) = uint64(*(*TCARD32)(unsafe.Pointer(req + uintptr(_BRlen)*4)))
				libc.Xmemmove(tls, req+uintptr(8), req+uintptr(4), uint64((_BRlen-uint32(1))<<int32(2)))
				*(*TCARD32)(unsafe.Pointer(req + 1*4)) = _BRlen + libc.Uint32FromInt64(len1) + uint32(2)
				libx11.X_XData32(tls, dpy, bp, libc.Uint32FromInt32(libc.Int32FromInt32(4)))
			} else {
				len1 = len1
				(*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength) + len1)
			}
		} else {
			(*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength = libc.Uint16FromInt64(libc.Int64FromUint16((*TxRenderTriFanReq)(unsafe.Pointer(req)).Flength) + len1)
		}
		len1 <<= int64(2)
		if (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr+uintptr(len1) <= (*TDisplay)(unsafe.Pointer(dpy)).Fbufmax {
			libc.Xmemcpy(tls, (*TDisplay)(unsafe.Pointer(dpy)).Fbufptr, points, libc.Uint64FromInt64(len1))
			*(*uintptr)(unsafe.Pointer(dpy + 176)) += uintptr(libc.Uint64FromInt64(len1+libc.Int64FromInt32(3)) & libc.Uint64FromInt32(^libc.Int32FromInt32(3)))
		} else {
			libx11.X_XSend(tls, dpy, points, len1)
		}
		npoint -= n - int32(1)
		points += uintptr(n-libc.Int32FromInt32(1)) * 8
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
}

const m___INT_MAX3 = 2147483647

// C documentation
//
//	/*
//	 * XRenderExtFindDisplay - look for a display in this extension; keeps a
//	 * cache of the most-recently used for efficiency. (Replaces
//	 * XextFindDisplay.)
//	 */
func _XRenderExtFindDisplay(tls *libc.TLS, extinfo uintptr, dpy uintptr) (r uintptr) {
	var dpyinfo, v1 uintptr
	_, _ = dpyinfo, v1
	/*
	 * see if this was the most recently accessed display
	 */
	v1 = (*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fcur
	dpyinfo = v1
	if v1 != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fdisplay == dpy {
		return dpyinfo
	}
	/*
	 * look for display in list
	 */
	if libx11.X_XLockMutex_fn != 0 {
		(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XLockMutex_fn})))(tls, libx11.X_Xglobal_lock)
	}
	dpyinfo = (*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fhead
	for {
		if !(dpyinfo != 0) {
			break
		}
		if (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fdisplay == dpy {
			(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fcur = dpyinfo /* cache most recently used */
			if libx11.X_XUnlockMutex_fn != 0 {
				(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
			}
			return dpyinfo
		}
		goto _2
	_2:
		;
		dpyinfo = (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fnext
	}
	if libx11.X_XUnlockMutex_fn != 0 {
		(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
	}
	return libc.UintptrFromInt32(0)
}

/*
 * If the server is missing support for any of the required depths on
 * any screen, tell the application that Render is not present.
 */

/*
 * Render requires support for depth 1, 4, 8, 24 and 32 pixmaps
 */

type TDepthCheckRec = struct {
	Fnext    uintptr
	Fdpy     uintptr
	Fmissing TCARD32
	Fserial  uint64
}

/*
 * If the server is missing support for any of the required depths on
 * any screen, tell the application that Render is not present.
 */

/*
 * Render requires support for depth 1, 4, 8, 24 and 32 pixmaps
 */

type T_DepthCheckRec = TDepthCheckRec

type TDepthCheckPtr = uintptr

var _depthChecks TDepthCheckPtr

func _XRenderDepthCheckErrorHandler(tls *libc.TLS, dpy uintptr, evt uintptr) (r int32) {
	var d TDepthCheckPtr
	_ = d
	if libc.Int32FromUint8((*TXErrorEvent)(unsafe.Pointer(evt)).Frequest_code) == int32(m_X_CreatePixmap) && libc.Int32FromUint8((*TXErrorEvent)(unsafe.Pointer(evt)).Ferror_code) == int32(m_BadValue) {
		if libx11.X_XLockMutex_fn != 0 {
			(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XLockMutex_fn})))(tls, libx11.X_Xglobal_lock)
		}
		d = _depthChecks
		for {
			if !(d != 0) {
				break
			}
			if (*T_DepthCheckRec)(unsafe.Pointer(d)).Fdpy == dpy {
				if libc.Int64FromUint64((*TXErrorEvent)(unsafe.Pointer(evt)).Fserial-(*T_DepthCheckRec)(unsafe.Pointer(d)).Fserial) >= 0 {
					*(*TCARD32)(unsafe.Pointer(d + 16)) |= uint32(1) << ((*TXErrorEvent)(unsafe.Pointer(evt)).Fresourceid - uint64(1))
				}
				break
			}
			goto _1
		_1:
			;
			d = (*T_DepthCheckRec)(unsafe.Pointer(d)).Fnext
		}
		if libx11.X_XUnlockMutex_fn != 0 {
			(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
		}
	}
	return 0
}

func _XRenderHasDepths(tls *libc.TLS, dpy uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var d, s int32
	var depths, missing TCARD32
	var dp, scr uintptr
	var p TPixmap
	var previousHandler TXErrorHandler
	var _ /* dc at bp+0 */ TDepthCheckRec
	_, _, _, _, _, _, _, _ = d, depths, dp, missing, p, previousHandler, s, scr
	s = 0
	for {
		if !(s < (*struct {
			Fext_data            uintptr
			Fprivate1            uintptr
			Ffd                  int32
			Fprivate2            int32
			Fproto_major_version int32
			Fproto_minor_version int32
			Fvendor              uintptr
			Fprivate3            TXID
			Fprivate4            TXID
			Fprivate5            TXID
			Fprivate6            int32
			Fresource_alloc      uintptr
			Fbyte_order          int32
			Fbitmap_unit         int32
			Fbitmap_pad          int32
			Fbitmap_bit_order    int32
			Fnformats            int32
			Fpixmap_format       uintptr
			Fprivate8            int32
			Frelease             int32
			Fprivate9            uintptr
			Fprivate10           uintptr
			Fqlen                int32
			Flast_request_read   uint64
			Frequest             uint64
			Fprivate11           TXPointer
			Fprivate12           TXPointer
			Fprivate13           TXPointer
			Fprivate14           TXPointer
			Fmax_request_size    uint32
			Fdb                  uintptr
			Fprivate15           uintptr
			Fdisplay_name        uintptr
			Fdefault_screen      int32
			Fnscreens            int32
			Fscreens             uintptr
			Fmotion_buffer       uint64
			Fprivate16           uint64
			Fmin_keycode         int32
			Fmax_keycode         int32
			Fprivate17           TXPointer
			Fprivate18           TXPointer
			Fprivate19           int32
			Fxdefaults           uintptr
		})(unsafe.Pointer(dpy)).Fnscreens) {
			break
		}
		depths = uint32(0)
		scr = (*struct {
			Fext_data            uintptr
			Fprivate1            uintptr
			Ffd                  int32
			Fprivate2            int32
			Fproto_major_version int32
			Fproto_minor_version int32
			Fvendor              uintptr
			Fprivate3            TXID
			Fprivate4            TXID
			Fprivate5            TXID
			Fprivate6            int32
			Fresource_alloc      uintptr
			Fbyte_order          int32
			Fbitmap_unit         int32
			Fbitmap_pad          int32
			Fbitmap_bit_order    int32
			Fnformats            int32
			Fpixmap_format       uintptr
			Fprivate8            int32
			Frelease             int32
			Fprivate9            uintptr
			Fprivate10           uintptr
			Fqlen                int32
			Flast_request_read   uint64
			Frequest             uint64
			Fprivate11           TXPointer
			Fprivate12           TXPointer
			Fprivate13           TXPointer
			Fprivate14           TXPointer
			Fmax_request_size    uint32
			Fdb                  uintptr
			Fprivate15           uintptr
			Fdisplay_name        uintptr
			Fdefault_screen      int32
			Fnscreens            int32
			Fscreens             uintptr
			Fmotion_buffer       uint64
			Fprivate16           uint64
			Fmin_keycode         int32
			Fmax_keycode         int32
			Fprivate17           TXPointer
			Fprivate18           TXPointer
			Fprivate19           int32
			Fxdefaults           uintptr
		})(unsafe.Pointer(dpy)).Fscreens + uintptr(s)*128
		d = 0
		for {
			if !(d < (*TScreen)(unsafe.Pointer(scr)).Fndepths) {
				break
			}
			depths |= uint32(1) << ((*(*TDepth)(unsafe.Pointer((*TScreen)(unsafe.Pointer(scr)).Fdepths + uintptr(d)*16))).Fdepth - int32(1))
			goto _2
		_2:
			;
			d++
		}
		missing = ^depths & (libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(1)-libc.Int32FromInt32(1)) | libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(4)-libc.Int32FromInt32(1)) | libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(8)-libc.Int32FromInt32(1)) | libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(24)-libc.Int32FromInt32(1)) | libc.Uint32FromUint32(1)<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(1)))
		if missing != 0 {
			/*
			 * Ok, this is ugly.  It should be sufficient at this
			 * point to just return False, but Xinerama is broken at
			 * this point and only advertises depths which have an
			 * associated visual.  Of course, the other depths still
			 * work, but the only way to find out is to try them.
			 */
			(*(*TDepthCheckRec)(unsafe.Pointer(bp))).Fdpy = dpy
			(*(*TDepthCheckRec)(unsafe.Pointer(bp))).Fmissing = uint32(0)
			(*(*TDepthCheckRec)(unsafe.Pointer(bp))).Fserial = libx11.XXNextRequest(tls, dpy)
			if libx11.X_XLockMutex_fn != 0 {
				(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XLockMutex_fn})))(tls, libx11.X_Xglobal_lock)
			}
			(*(*TDepthCheckRec)(unsafe.Pointer(bp))).Fnext = _depthChecks
			_depthChecks = bp
			if libx11.X_XUnlockMutex_fn != 0 {
				(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
			}
			/*
			 * I suspect this is not really thread safe, but Xlib doesn't
			 * provide a lot of options here
			 */
			previousHandler = libx11.XXSetErrorHandler(tls, __ccgo_fp(_XRenderDepthCheckErrorHandler))
			/*
			 * Try each missing depth and see if pixmap creation succeeds
			 */
			d = int32(1)
			for {
				if !(d <= int32(32)) {
					break
				}
				/* don't check depth 1 == Xcursor recurses... */
				if missing&(uint32(1)<<(d-int32(1))) != 0 && d != int32(1) {
					p = libx11.XXCreatePixmap(tls, dpy, (*TScreen)(unsafe.Pointer((*struct {
						Fext_data            uintptr
						Fprivate1            uintptr
						Ffd                  int32
						Fprivate2            int32
						Fproto_major_version int32
						Fproto_minor_version int32
						Fvendor              uintptr
						Fprivate3            TXID
						Fprivate4            TXID
						Fprivate5            TXID
						Fprivate6            int32
						Fresource_alloc      uintptr
						Fbyte_order          int32
						Fbitmap_unit         int32
						Fbitmap_pad          int32
						Fbitmap_bit_order    int32
						Fnformats            int32
						Fpixmap_format       uintptr
						Fprivate8            int32
						Frelease             int32
						Fprivate9            uintptr
						Fprivate10           uintptr
						Fqlen                int32
						Flast_request_read   uint64
						Frequest             uint64
						Fprivate11           TXPointer
						Fprivate12           TXPointer
						Fprivate13           TXPointer
						Fprivate14           TXPointer
						Fmax_request_size    uint32
						Fdb                  uintptr
						Fprivate15           uintptr
						Fdisplay_name        uintptr
						Fdefault_screen      int32
						Fnscreens            int32
						Fscreens             uintptr
						Fmotion_buffer       uint64
						Fprivate16           uint64
						Fmin_keycode         int32
						Fmax_keycode         int32
						Fprivate17           TXPointer
						Fprivate18           TXPointer
						Fprivate19           int32
						Fxdefaults           uintptr
					})(unsafe.Pointer(dpy)).Fscreens+uintptr(s)*128)).Froot, uint32(1), uint32(1), libc.Uint32FromInt32(d))
					libx11.XXFreePixmap(tls, dpy, p)
				}
				goto _3
			_3:
				;
				d++
			}
			libx11.XXSync(tls, dpy, m_False)
			libx11.XXSetErrorHandler(tls, previousHandler)
			/*
			 * Unhook from the list of depth check records
			 */
			if libx11.X_XLockMutex_fn != 0 {
				(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XLockMutex_fn})))(tls, libx11.X_Xglobal_lock)
			}
			dp = uintptr(unsafe.Pointer(&_depthChecks))
			for {
				if !(*(*uintptr)(unsafe.Pointer(dp)) != 0) {
					break
				}
				if *(*uintptr)(unsafe.Pointer(dp)) == bp {
					*(*uintptr)(unsafe.Pointer(dp)) = (*(*TDepthCheckRec)(unsafe.Pointer(bp))).Fnext
					break
				}
				goto _4
			_4:
				;
				dp = *(*uintptr)(unsafe.Pointer(dp))
			}
			if libx11.X_XUnlockMutex_fn != 0 {
				(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
			}
			if (*(*TDepthCheckRec)(unsafe.Pointer(bp))).Fmissing != 0 {
				return m_False
			}
		}
		goto _1
	_1:
		;
		s++
	}
	return int32(m_True)
}

// C documentation
//
//	/*
//	 * XRenderExtAddDisplay - add a display to this extension. (Replaces
//	 * XextAddDisplay)
//	 */
func _XRenderExtAddDisplay(tls *libc.TLS, extinfo uintptr, dpy uintptr, ext_name uintptr) (r uintptr) {
	var codes, dpyinfo uintptr
	_, _ = codes, dpyinfo
	dpyinfo = libc.Xmalloc(tls, libc.Uint64FromInt64(32))
	if !(dpyinfo != 0) {
		return libc.UintptrFromInt32(0)
	}
	(*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fdisplay = dpy
	(*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Finfo = libc.UintptrFromInt32(0)
	if _XRenderHasDepths(tls, dpy) != 0 {
		(*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fcodes = libx11.XXInitExtension(tls, dpy, ext_name)
	} else {
		(*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fcodes = libc.UintptrFromInt32(0)
	}
	/*
	 * if the server has the extension, then we can initialize the
	 * appropriate function vectors
	 */
	if (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fcodes != 0 {
		libx11.XXESetCloseDisplay(tls, dpy, (*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fcodes)).Fextension, __ccgo_fp(_XRenderCloseDisplay))
	} else {
		/* The server doesn't have this extension.
		 * Use a private Xlib-internal extension to hang the close_display
		 * hook on so that the "cache" (extinfo->cur) is properly cleaned.
		 * (XBUG 7955)
		 */
		codes = libx11.XXAddExtension(tls, dpy)
		if !(codes != 0) {
			libx11.XXFree(tls, dpyinfo)
			return libc.UintptrFromInt32(0)
		}
		libx11.XXESetCloseDisplay(tls, dpy, (*TXExtCodes)(unsafe.Pointer(codes)).Fextension, __ccgo_fp(_XRenderCloseDisplay))
	}
	/*
	 * now, chain it onto the list
	 */
	if libx11.X_XLockMutex_fn != 0 {
		(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XLockMutex_fn})))(tls, libx11.X_Xglobal_lock)
	}
	(*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fnext = (*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fhead
	(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fhead = dpyinfo
	(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fcur = dpyinfo
	(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fndisplays++
	if libx11.X_XUnlockMutex_fn != 0 {
		(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
	}
	return dpyinfo
}

// C documentation
//
//	/*
//	 * XRenderExtRemoveDisplay - remove the indicated display from the
//	 * extension object. (Replaces XextRemoveDisplay.)
//	 */
func _XRenderExtRemoveDisplay(tls *libc.TLS, extinfo uintptr, dpy uintptr) (r int32) {
	var dpyinfo, prev uintptr
	_, _ = dpyinfo, prev
	/*
	 * locate this display and its back link so that it can be removed
	 */
	if libx11.X_XLockMutex_fn != 0 {
		(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XLockMutex_fn})))(tls, libx11.X_Xglobal_lock)
	}
	prev = libc.UintptrFromInt32(0)
	dpyinfo = (*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fhead
	for {
		if !(dpyinfo != 0) {
			break
		}
		if (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fdisplay == dpy {
			break
		}
		prev = dpyinfo
		goto _1
	_1:
		;
		dpyinfo = (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fnext
	}
	if !(dpyinfo != 0) {
		if libx11.X_XUnlockMutex_fn != 0 {
			(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
		}
		return 0 /* hmm, actually an error */
	}
	/*
	 * remove the display from the list; handles going to zero
	 */
	if prev != 0 {
		(*TXRenderExtDisplayInfo)(unsafe.Pointer(prev)).Fnext = (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fnext
	} else {
		(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fhead = (*TXRenderExtDisplayInfo)(unsafe.Pointer(dpyinfo)).Fnext
	}
	(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fndisplays--
	if dpyinfo == (*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fcur {
		(*TXRenderExtInfo)(unsafe.Pointer(extinfo)).Fcur = libc.UintptrFromInt32(0)
	} /* flush cache */
	if libx11.X_XUnlockMutex_fn != 0 {
		(*(*func(*libc.TLS, TLockInfoPtr))(unsafe.Pointer(&struct{ uintptr }{libx11.X_XUnlockMutex_fn})))(tls, libx11.X_Xglobal_lock)
	}
	libc.Xfree(tls, dpyinfo)
	return int32(1)
}

func XXRenderFindDisplay(tls *libc.TLS, dpy uintptr) (r uintptr) {
	var dpyinfo uintptr
	_ = dpyinfo
	dpyinfo = _XRenderExtFindDisplay(tls, uintptr(unsafe.Pointer(&XXRenderExtensionInfo)), dpy)
	if !(dpyinfo != 0) {
		dpyinfo = _XRenderExtAddDisplay(tls, uintptr(unsafe.Pointer(&XXRenderExtensionInfo)), dpy, uintptr(unsafe.Pointer(&XXRenderExtensionName)))
	}
	return dpyinfo
}

func _XRenderCloseDisplay(tls *libc.TLS, dpy uintptr, codes uintptr) (r int32) {
	var info uintptr
	_ = info
	info = XXRenderFindDisplay(tls, dpy)
	if info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo != 0 {
		libx11.XXFree(tls, (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo)
	}
	return _XRenderExtRemoveDisplay(tls, uintptr(unsafe.Pointer(&XXRenderExtensionInfo)), dpy)
}

/****************************************************************************
 *                                                                          *
 *			    Render public interfaces                        *
 *                                                                          *
 ****************************************************************************/
func XXRenderQueryExtension(tls *libc.TLS, dpy uintptr, event_basep uintptr, error_basep uintptr) (r int32) {
	var info uintptr
	_ = info
	info = XXRenderFindDisplay(tls, dpy)
	if info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0 {
		*(*int32)(unsafe.Pointer(event_basep)) = (*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Ffirst_event
		*(*int32)(unsafe.Pointer(error_basep)) = (*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Ffirst_error
		return int32(m_True)
	} else {
		return m_False
	}
	return r
}

func XXRenderQueryVersion(tls *libc.TLS, dpy uintptr, major_versionp uintptr, minor_versionp uintptr) (r int32) {
	var info, xri uintptr
	_, _ = info, xri
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return 0
	}
	if !(XXRenderQueryFormats(tls, dpy) != 0) {
		return 0
	}
	xri = (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo
	*(*int32)(unsafe.Pointer(major_versionp)) = (*TXRenderInfo)(unsafe.Pointer(xri)).Fmajor_version
	*(*int32)(unsafe.Pointer(minor_versionp)) = (*TXRenderInfo)(unsafe.Pointer(xri)).Fminor_version
	return int32(1)
}

func __XRenderFindFormat(tls *libc.TLS, xri uintptr, format TPictFormat) (r uintptr) {
	var nf int32
	_ = nf
	nf = 0
	for {
		if !(nf < (*TXRenderInfo)(unsafe.Pointer(xri)).Fnformat) {
			break
		}
		if (*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fid == format {
			return (*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40
		}
		goto _1
	_1:
		;
		nf++
	}
	return libc.UintptrFromInt32(0)
}

func __XRenderFindVisual(tls *libc.TLS, dpy uintptr, vid TVisualID) (r uintptr) {
	return libx11.X_XVIDtoVisual(tls, dpy, vid)
}

type T_XrenderVersionState = struct {
	Fversion_seq   uint64
	Ferror1        int32
	Fmajor_version int32
	Fminor_version int32
}

type T_renderVersionState = T_XrenderVersionState

func __XRenderVersionHandler(tls *libc.TLS, dpy uintptr, rep uintptr, buf uintptr, len1 int32, data TXPointer) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var repl, state uintptr
	var _ /* replbuf at bp+0 */ TxRenderQueryVersionReply
	_, _ = repl, state
	state = data
	if (*TDisplay)(unsafe.Pointer(dpy)).Flast_request_read != (*T_XrenderVersionState)(unsafe.Pointer(state)).Fversion_seq {
		return m_False
	}
	if libc.Int32FromUint8((*TxReply)(unsafe.Pointer(rep)).Fgeneric.Ftype1) == m_X_Error {
		(*T_XrenderVersionState)(unsafe.Pointer(state)).Ferror1 = int32(m_True)
		return m_False
	}
	repl = libx11.X_XGetAsyncReply(tls, dpy, bp, rep, buf, len1, (libc.Int32FromInt32(m_sz_xRenderQueryVersionReply)-libc.Int32FromInt32(m_sz_xReply))>>libc.Int32FromInt32(2), int32(m_True))
	(*T_XrenderVersionState)(unsafe.Pointer(state)).Fmajor_version = libc.Int32FromUint32((*TxRenderQueryVersionReply)(unsafe.Pointer(repl)).FmajorVersion)
	(*T_XrenderVersionState)(unsafe.Pointer(state)).Fminor_version = libc.Int32FromUint32((*TxRenderQueryVersionReply)(unsafe.Pointer(repl)).FminorVersion)
	return int32(m_True)
}

func XXRenderQueryFormats(tls *libc.TLS, dpy uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var depth, format, info, req, screen, visual, vreq, xData, xDepth, xFormat, xScreen, xSubpixel, xVisual, xri uintptr
	var nbytes, rlength, v1 uint64
	var nd, nf, ns, nv int32
	var _ /* async at bp+0 */ T_XAsyncHandler
	var _ /* async_state at bp+24 */ T_XrenderVersionState
	var _ /* rep at bp+48 */ TxRenderQueryPictFormatsReply
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = depth, format, info, nbytes, nd, nf, ns, nv, req, rlength, screen, visual, vreq, xData, xDepth, xFormat, xScreen, xSubpixel, xVisual, xri, v1
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return 0
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	if (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo != 0 {
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		return int32(1)
	}
	vreq = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderQueryVersion), uint64(m_sz_xRenderQueryVersionReq))
	(*TxRenderQueryVersionReq)(unsafe.Pointer(vreq)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderQueryVersionReq)(unsafe.Pointer(vreq)).FrenderReqType = uint8(m_X_RenderQueryVersion)
	(*TxRenderQueryVersionReq)(unsafe.Pointer(vreq)).FmajorVersion = uint32(m_RENDER_MAJOR)
	(*TxRenderQueryVersionReq)(unsafe.Pointer(vreq)).FminorVersion = uint32(m_RENDER_MINOR)
	(*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Fversion_seq = (*TDisplay)(unsafe.Pointer(dpy)).Frequest
	(*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Ferror1 = m_False
	(*(*T_XAsyncHandler)(unsafe.Pointer(bp))).Fnext = (*TDisplay)(unsafe.Pointer(dpy)).Fasync_handlers
	(*(*T_XAsyncHandler)(unsafe.Pointer(bp))).Fhandler = __ccgo_fp(__XRenderVersionHandler)
	(*(*T_XAsyncHandler)(unsafe.Pointer(bp))).Fdata = bp + 24
	(*TDisplay)(unsafe.Pointer(dpy)).Fasync_handlers = bp
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderQueryPictFormats), uint64(m_sz_xRenderQueryPictFormatsReq))
	(*TxRenderQueryPictFormatsReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderQueryPictFormatsReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderQueryPictFormats)
	if !(libx11.X_XReply(tls, dpy, bp+48, 0, m_xFalse) != 0) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Fasync_handlers == bp {
			(*TDisplay)(unsafe.Pointer(dpy)).Fasync_handlers = (*T_XAsyncHandler)(unsafe.Pointer(bp)).Fnext
		} else {
			libx11.X_XDeqAsyncHandler(tls, dpy, bp)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return 0
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fasync_handlers == bp {
		(*TDisplay)(unsafe.Pointer(dpy)).Fasync_handlers = (*T_XAsyncHandler)(unsafe.Pointer(bp)).Fnext
	} else {
		libx11.X_XDeqAsyncHandler(tls, dpy, bp)
	}
	if (*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Ferror1 != 0 {
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return 0
	}
	/*
	 * Check for the lack of sub-pixel data
	 */
	if (*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Fmajor_version == 0 && (*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Fminor_version < int32(6) {
		(*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumSubpixel = uint32(0)
	}
	if uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumFormats) < libc.Uint64FromInt32(libc.Int32FromInt32(m___INT_MAX3)/libc.Int32FromInt32(4))/libc.Uint64FromInt64(40) && uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumScreens) < libc.Uint64FromInt32(libc.Int32FromInt32(m___INT_MAX3)/libc.Int32FromInt32(4))/libc.Uint64FromInt64(32) && uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths) < libc.Uint64FromInt32(libc.Int32FromInt32(m___INT_MAX3)/libc.Int32FromInt32(4))/libc.Uint64FromInt64(16) && uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumVisuals) < libc.Uint64FromInt32(libc.Int32FromInt32(m___INT_MAX3)/libc.Int32FromInt32(4))/libc.Uint64FromInt64(16) && (*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumSubpixel < libc.Uint32FromInt32(libc.Int32FromInt32(m___INT_MAX3)/libc.Int32FromInt32(4)/libc.Int32FromInt32(4)) && (*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).Flength < libc.Uint32FromInt32(libc.Int32FromInt32(m___INT_MAX3)>>libc.Int32FromInt32(2)) {
		xri = libc.Xmalloc(tls, libc.Uint64FromInt64(120)+uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumFormats)*libc.Uint64FromInt64(40)+uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumScreens)*libc.Uint64FromInt64(32)+uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths)*libc.Uint64FromInt64(16)+uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumVisuals)*libc.Uint64FromInt64(16))
		rlength = uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumFormats)*uint64(28) + uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumScreens)*uint64(8) + uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths)*uint64(8) + uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumVisuals)*uint64(8) + uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumSubpixel*libc.Uint32FromInt32(4))
		xData = libc.Xmalloc(tls, rlength)
		nbytes = uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).Flength) << int32(2)
	} else {
		xri = libc.UintptrFromInt32(0)
		xData = libc.UintptrFromInt32(0)
		v1 = libc.Uint64FromInt32(0)
		nbytes = v1
		rlength = v1
	}
	if !(xri != 0) || !(xData != 0) || nbytes < rlength {
		if xri != 0 {
			libc.Xfree(tls, xri)
		}
		if xData != 0 {
			libc.Xfree(tls, xData)
		}
		libx11.X_XEatDataWords(tls, dpy, uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).Flength))
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return 0
	}
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fmajor_version = (*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Fmajor_version
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fminor_version = (*(*T_XrenderVersionState)(unsafe.Pointer(bp + 24))).Fminor_version
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fformat = xri + libc.UintptrFromInt32(1)*120
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fnformat = libc.Int32FromUint32((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumFormats)
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fscreen = (*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumFormats)*40
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fnscreen = libc.Int32FromUint32((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumScreens)
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fdepth = (*TXRenderInfo)(unsafe.Pointer(xri)).Fscreen + uintptr((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumScreens)*32
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fndepth = libc.Int32FromUint32((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths)
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fvisual = (*TXRenderInfo)(unsafe.Pointer(xri)).Fdepth + uintptr((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths)*16
	(*TXRenderInfo)(unsafe.Pointer(xri)).Fnvisual = libc.Int32FromUint32((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumVisuals)
	libx11.X_XRead(tls, dpy, xData, libc.Int64FromUint64(rlength))
	format = (*TXRenderInfo)(unsafe.Pointer(xri)).Fformat
	xFormat = xData
	nf = 0
	for {
		if !(libc.Uint32FromInt32(nf) < (*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumFormats) {
			break
		}
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fid = uint64((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fid)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Ftype1 = libc.Int32FromUint8((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Ftype1)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdepth = libc.Int32FromUint8((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdepth)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.Fred = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.Fred)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.FredMask = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.FredMask)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.Fgreen = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.Fgreen)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.FgreenMask = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.FgreenMask)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.Fblue = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.Fblue)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.FblueMask = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.FblueMask)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.Falpha = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.Falpha)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fdirect.FalphaMask = libc.Int16FromUint16((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fdirect.FalphaMask)
		(*TXRenderPictFormat)(unsafe.Pointer(format)).Fcolormap = uint64((*TxPictFormInfo)(unsafe.Pointer(xFormat)).Fcolormap)
		format += 40
		xFormat += 28
		goto _2
	_2:
		;
		nf++
	}
	xScreen = xFormat
	screen = (*TXRenderInfo)(unsafe.Pointer(xri)).Fscreen
	depth = (*TXRenderInfo)(unsafe.Pointer(xri)).Fdepth
	visual = (*TXRenderInfo)(unsafe.Pointer(xri)).Fvisual
	ns = 0
	for {
		if !(ns < (*TXRenderInfo)(unsafe.Pointer(xri)).Fnscreen) {
			break
		}
		(*TXRenderScreen)(unsafe.Pointer(screen)).Fdepths = depth
		(*TXRenderScreen)(unsafe.Pointer(screen)).Fndepths = libc.Int32FromUint32((*TxPictScreen)(unsafe.Pointer(xScreen)).FnDepth)
		(*TXRenderScreen)(unsafe.Pointer(screen)).Ffallback = __XRenderFindFormat(tls, xri, uint64((*TxPictScreen)(unsafe.Pointer(xScreen)).Ffallback))
		(*TXRenderScreen)(unsafe.Pointer(screen)).Fsubpixel = m_SubPixelUnknown
		xDepth = xScreen + libc.UintptrFromInt32(1)*8
		if libc.Uint32FromInt32((*TXRenderScreen)(unsafe.Pointer(screen)).Fndepths) > (*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths {
			libc.Xfree(tls, xri)
			libc.Xfree(tls, xData)
			libx11.X_XEatDataWords(tls, dpy, uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).Flength))
			if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
			}
			if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
				(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
			}
			return 0
		}
		(*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumDepths -= libc.Uint32FromInt32((*TXRenderScreen)(unsafe.Pointer(screen)).Fndepths)
		nd = 0
		for {
			if !(nd < (*TXRenderScreen)(unsafe.Pointer(screen)).Fndepths) {
				break
			}
			(*TXRenderDepth)(unsafe.Pointer(depth)).Fdepth = libc.Int32FromUint8((*TxPictDepth)(unsafe.Pointer(xDepth)).Fdepth)
			(*TXRenderDepth)(unsafe.Pointer(depth)).Fnvisuals = libc.Int32FromUint16((*TxPictDepth)(unsafe.Pointer(xDepth)).FnPictVisuals)
			(*TXRenderDepth)(unsafe.Pointer(depth)).Fvisuals = visual
			xVisual = xDepth + libc.UintptrFromInt32(1)*8
			if libc.Uint32FromInt32((*TXRenderDepth)(unsafe.Pointer(depth)).Fnvisuals) > (*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumVisuals {
				libc.Xfree(tls, xri)
				libc.Xfree(tls, xData)
				libx11.X_XEatDataWords(tls, dpy, uint64((*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).Flength))
				if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
				}
				if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
					(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
				}
				return 0
			}
			(*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumVisuals -= libc.Uint32FromInt32((*TXRenderDepth)(unsafe.Pointer(depth)).Fnvisuals)
			nv = 0
			for {
				if !(nv < (*TXRenderDepth)(unsafe.Pointer(depth)).Fnvisuals) {
					break
				}
				(*TXRenderVisual)(unsafe.Pointer(visual)).Fvisual = __XRenderFindVisual(tls, dpy, uint64((*TxPictVisual)(unsafe.Pointer(xVisual)).Fvisual))
				(*TXRenderVisual)(unsafe.Pointer(visual)).Fformat = __XRenderFindFormat(tls, xri, uint64((*TxPictVisual)(unsafe.Pointer(xVisual)).Fformat))
				visual += 16
				xVisual += 8
				goto _5
			_5:
				;
				nv++
			}
			depth += 16
			xDepth = xVisual
			goto _4
		_4:
			;
			nd++
		}
		screen += 32
		xScreen = xDepth
		goto _3
	_3:
		;
		ns++
	}
	xSubpixel = xScreen
	screen = (*TXRenderInfo)(unsafe.Pointer(xri)).Fscreen
	ns = 0
	for {
		if !(libc.Uint32FromInt32(ns) < (*(*TxRenderQueryPictFormatsReply)(unsafe.Pointer(bp + 48))).FnumSubpixel) {
			break
		}
		(*TXRenderScreen)(unsafe.Pointer(screen)).Fsubpixel = libc.Int32FromUint32(*(*TCARD32)(unsafe.Pointer(xSubpixel)))
		xSubpixel += 4
		screen += 32
		goto _6
	_6:
		;
		ns++
	}
	(*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo = xri
	/*
	 * Skip any extra data
	 */
	if nbytes > rlength {
		libx11.X_XEatData(tls, dpy, nbytes-rlength)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	libc.Xfree(tls, xData)
	return int32(1)
}

func XXRenderQuerySubpixelOrder(tls *libc.TLS, dpy uintptr, screen int32) (r int32) {
	var info, xri uintptr
	_, _ = info, xri
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return m_SubPixelUnknown
	}
	if !(XXRenderQueryFormats(tls, dpy) != 0) {
		return m_SubPixelUnknown
	}
	xri = (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo
	return (*(*TXRenderScreen)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fscreen + uintptr(screen)*32))).Fsubpixel
}

func XXRenderSetSubpixelOrder(tls *libc.TLS, dpy uintptr, screen int32, subpixel int32) (r int32) {
	var info, xri uintptr
	_, _ = info, xri
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return m_False
	}
	if !(XXRenderQueryFormats(tls, dpy) != 0) {
		return m_False
	}
	xri = (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo
	(*(*TXRenderScreen)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fscreen + uintptr(screen)*32))).Fsubpixel = subpixel
	return int32(m_True)
}

func XXRenderFindVisualFormat(tls *libc.TLS, dpy uintptr, visual uintptr) (r uintptr) {
	var info, xri, xrv uintptr
	var nv int32
	_, _, _, _ = info, nv, xri, xrv
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return libc.UintptrFromInt32(0)
	}
	if !(XXRenderQueryFormats(tls, dpy) != 0) {
		return libc.UintptrFromInt32(0)
	}
	xri = (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo
	nv = 0
	xrv = (*TXRenderInfo)(unsafe.Pointer(xri)).Fvisual
	for {
		if !(nv < (*TXRenderInfo)(unsafe.Pointer(xri)).Fnvisual) {
			break
		}
		if (*TXRenderVisual)(unsafe.Pointer(xrv)).Fvisual == visual {
			return (*TXRenderVisual)(unsafe.Pointer(xrv)).Fformat
		}
		goto _1
	_1:
		;
		nv++
		xrv += 16
	}
	return libc.UintptrFromInt32(0)
}

func XXRenderFindFormat(tls *libc.TLS, dpy uintptr, mask uint64, template uintptr, count int32) (r uintptr) {
	var info, xri uintptr
	var nf, v2 int32
	_, _, _, _ = info, nf, xri, v2
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return libc.UintptrFromInt32(0)
	}
	if !(XXRenderQueryFormats(tls, dpy) != 0) {
		return libc.UintptrFromInt32(0)
	}
	xri = (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Finfo
	nf = 0
	for {
		if !(nf < (*TXRenderInfo)(unsafe.Pointer(xri)).Fnformat) {
			break
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(0)) != 0 {
			if (*TXRenderPictFormat)(unsafe.Pointer(template)).Fid != (*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fid {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1)) != 0 {
			if (*TXRenderPictFormat)(unsafe.Pointer(template)).Ftype1 != (*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Ftype1 {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(2)) != 0 {
			if (*TXRenderPictFormat)(unsafe.Pointer(template)).Fdepth != (*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdepth {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(3)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.Fred) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.Fred) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(4)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.FredMask) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.FredMask) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(5)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.Fgreen) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.Fgreen) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(6)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.FgreenMask) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.FgreenMask) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(7)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.Fblue) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.Fblue) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(8)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.FblueMask) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.FblueMask) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(9)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.Falpha) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.Falpha) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)) != 0 {
			if int32((*TXRenderPictFormat)(unsafe.Pointer(template)).Fdirect.FalphaMask) != int32((*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fdirect.FalphaMask) {
				goto _1
			}
		}
		if mask&libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(11)) != 0 {
			if (*TXRenderPictFormat)(unsafe.Pointer(template)).Fcolormap != (*(*TXRenderPictFormat)(unsafe.Pointer((*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40))).Fcolormap {
				goto _1
			}
		}
		v2 = count
		count--
		if v2 == 0 {
			return (*TXRenderInfo)(unsafe.Pointer(xri)).Fformat + uintptr(nf)*40
		}
		goto _1
	_1:
		;
		nf++
	}
	return libc.UintptrFromInt32(0)
}

func XXRenderFindStandardFormat(tls *libc.TLS, dpy uintptr, format int32) (r uintptr) {
	if 0 <= format && format < int32(m_PictStandardNUM) {
		return XXRenderFindFormat(tls, dpy, _standardFormats[format].Fmask, uintptr(unsafe.Pointer(&_standardFormats))+uintptr(format)*48, 0)
	}
	return libc.UintptrFromInt32(0)
}

var _standardFormats = [5]struct {
	Ftempl TXRenderPictFormat
	Fmask  uint64
}{
	0: {
		Ftempl: TXRenderPictFormat{
			Ftype1: int32(m_PictTypeDirect),
			Fdepth: int32(32),
			Fdirect: TXRenderDirectFormat{
				Fred:       int16(16),
				FredMask:   int16(0xff),
				Fgreen:     int16(8),
				FgreenMask: int16(0xff),
				FblueMask:  int16(0xff),
				Falpha:     int16(24),
				FalphaMask: int16(0xff),
			},
		},
		Fmask: libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(2) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(3) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(4) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(5) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(6) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(7) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(9) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)),
	},
	1: {
		Ftempl: TXRenderPictFormat{
			Ftype1: int32(m_PictTypeDirect),
			Fdepth: int32(24),
			Fdirect: TXRenderDirectFormat{
				Fred:       int16(16),
				FredMask:   int16(0xff),
				Fgreen:     int16(8),
				FgreenMask: int16(0xff),
				FblueMask:  int16(0xff),
			},
		},
		Fmask: libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(2) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(3) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(4) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(5) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(6) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(7) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)),
	},
	2: {
		Ftempl: TXRenderPictFormat{
			Ftype1: int32(m_PictTypeDirect),
			Fdepth: int32(8),
			Fdirect: TXRenderDirectFormat{
				FalphaMask: int16(0xff),
			},
		},
		Fmask: libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(2) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(4) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(6) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(9) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)),
	},
	3: {
		Ftempl: TXRenderPictFormat{
			Ftype1: int32(m_PictTypeDirect),
			Fdepth: int32(4),
			Fdirect: TXRenderDirectFormat{
				FalphaMask: int16(0x0f),
			},
		},
		Fmask: libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(2) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(4) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(6) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(9) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)),
	},
	4: {
		Ftempl: TXRenderPictFormat{
			Ftype1: int32(m_PictTypeDirect),
			Fdepth: int32(1),
			Fdirect: TXRenderDirectFormat{
				FalphaMask: int16(0x01),
			},
		},
		Fmask: libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(1) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(2) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(4) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(6) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(9) | libc.Int32FromInt32(1)<<libc.Int32FromInt32(10)),
	},
}

func XXRenderQueryPictIndexValues(tls *libc.TLS, dpy uintptr, format uintptr, num uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i, nbytes, nread, rlength, v1, v2 uint32
	var info, req, values uintptr
	var _ /* rep at bp+0 */ TxRenderQueryPictIndexValuesReply
	var _ /* value at bp+32 */ TxIndexValue
	_, _, _, _, _, _, _, _, _ = i, info, nbytes, nread, req, rlength, values, v1, v2
	info = XXRenderFindDisplay(tls, dpy)
	if !(info != 0 && (*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes != 0) {
		return libc.UintptrFromInt32(0)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Flock_display})))(tls, dpy)
	}
	req = libx11.X_XGetRequest(tls, dpy, uint8(m_X_RenderQueryPictIndexValues), uint64(m_sz_xRenderQueryPictIndexValuesReq))
	(*TxRenderQueryPictIndexValuesReq)(unsafe.Pointer(req)).FreqType = libc.Uint8FromInt32((*TXExtCodes)(unsafe.Pointer((*TXRenderExtDisplayInfo)(unsafe.Pointer(info)).Fcodes)).Fmajor_opcode)
	(*TxRenderQueryPictIndexValuesReq)(unsafe.Pointer(req)).FrenderReqType = uint8(m_X_RenderQueryPictIndexValues)
	(*TxRenderQueryPictIndexValuesReq)(unsafe.Pointer(req)).Fformat = uint32((*TXRenderPictFormat)(unsafe.Pointer(format)).Fid)
	if !(libx11.X_XReply(tls, dpy, bp, 0, m_xFalse) != 0) {
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return libc.UintptrFromInt32(0)
	}
	if (*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).Flength < libc.Uint32FromInt32(libc.Int32FromInt32(m___INT_MAX3)>>libc.Int32FromInt32(2)) && uint64((*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).FnumIndexValues) < libc.Uint64FromInt32(m___INT_MAX3)/libc.Uint64FromInt64(16) {
		/* request data length */
		nbytes = (*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).Flength << int32(2)
		/* bytes of actual data in the request */
		nread = (*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).FnumIndexValues * uint32(m_sz_xIndexValue)
		/* size of array returned to application */
		rlength = uint32(uint64((*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).FnumIndexValues) * uint64(16))
		/* allocate returned data */
		values = libc.Xmalloc(tls, uint64(rlength))
	} else {
		v2 = libc.Uint32FromInt32(0)
		rlength = v2
		v1 = v2
		nread = v1
		nbytes = v1
		values = libc.UintptrFromInt32(0)
	}
	if !(values != 0) {
		libx11.X_XEatDataWords(tls, dpy, uint64((*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).Flength))
		if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
		}
		if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
			(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
		}
		return libc.UintptrFromInt32(0)
	}
	/* read the values one at a time and convert */
	*(*int32)(unsafe.Pointer(num)) = libc.Int32FromUint32((*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).FnumIndexValues)
	i = uint32(0)
	for {
		if !(i < (*(*TxRenderQueryPictIndexValuesReply)(unsafe.Pointer(bp))).FnumIndexValues) {
			break
		}
		libx11.X_XRead(tls, dpy, bp+32, int64(m_sz_xIndexValue))
		(*(*TXIndexValue)(unsafe.Pointer(values + uintptr(i)*16))).Fpixel = uint64((*(*TxIndexValue)(unsafe.Pointer(bp + 32))).Fpixel)
		(*(*TXIndexValue)(unsafe.Pointer(values + uintptr(i)*16))).Fred = (*(*TxIndexValue)(unsafe.Pointer(bp + 32))).Fred
		(*(*TXIndexValue)(unsafe.Pointer(values + uintptr(i)*16))).Fgreen = (*(*TxIndexValue)(unsafe.Pointer(bp + 32))).Fgreen
		(*(*TXIndexValue)(unsafe.Pointer(values + uintptr(i)*16))).Fblue = (*(*TxIndexValue)(unsafe.Pointer(bp + 32))).Fblue
		(*(*TXIndexValue)(unsafe.Pointer(values + uintptr(i)*16))).Falpha = (*(*TxIndexValue)(unsafe.Pointer(bp + 32))).Falpha
		goto _3
	_3:
		;
		i++
	}
	/* skip any padding */
	if nbytes > nread {
		libx11.X_XEatData(tls, dpy, uint64(nbytes-nread))
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Flock_fns != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*T_XLockPtrs)(unsafe.Pointer((*TDisplay)(unsafe.Pointer(dpy)).Flock_fns)).Funlock_display})))(tls, dpy)
	}
	if (*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler != 0 {
		(*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*TDisplay)(unsafe.Pointer(dpy)).Fsynchandler})))(tls, dpy)
	}
	return values
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

/*
 * Xlib uses long for 32-bit values.  Xrender uses int.  This
 * matters on alpha.  Note that this macro assumes that int is 32 bits
 * except on WORD64 machines where it is 64 bits.
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)limits.h	8.2 (Berkeley) 1/4/94
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1991, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * This code is derived from software contributed to Berkeley by
 * Berkeley Software Design, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)cdefs.h	8.8 (Berkeley) 1/9/95
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1991, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * This code is derived from software contributed to Berkeley by
 * Berkeley Software Design, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)cdefs.h	8.8 (Berkeley) 1/9/95
 */

/*-
 * This file is in the public domain.
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)limits.h	8.3 (Berkeley) 1/4/94
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1988, 1993
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
 * 3. Neither the name of the University nor the names of its contributors
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
 *	@(#)syslimits.h	8.1 (Berkeley) 6/2/93
 */

/*
 * Do not add any new variables here.  (See the comment at the end of
 * the file for why.)
 */

/*
 * We leave the following values undefined to force applications to either
 * assume conservative values or call sysconf() to get the current value.
 *
 * HOST_NAME_MAX
 *
 * (We should do this for most of the values currently defined here,
 * but many programs are not prepared to deal with this yet.)
 */

var XXRenderExtensionInfo TXRenderExtInfo

var XXRenderExtensionName = [7]int8{'R', 'E', 'N', 'D', 'E', 'R'}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "rgba:\x00"
