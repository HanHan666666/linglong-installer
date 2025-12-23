// Code generated for linux/386 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -I /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libXau/include/linux/386 -I /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libXdmcp/include/linux/386 -lXau -lXdmcp -o libxcb.go --package-name libxcb src/.libs/libxcb.a', DO NOT EDIT.

//go:build linux && 386

package libxcb

import (
	"reflect"
	"unsafe"

	"modernc.org/libXau"
	"modernc.org/libXdmcp"
	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

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
const m_ARG_MAX = 131072
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
const m_BC_BASE_MAX = 99
const m_BC_DIM_MAX = 2048
const m_BC_SCALE_MAX = 99
const m_BC_STRING_MAX = 1000
const m_BIG_ENDIAN = "__BIG_ENDIAN"
const m_BUFSIZ = 1024
const m_BYTE_ORDER = "__BYTE_ORDER"
const m_CHARCLASS_NAME_MAX = 14
const m_CHAR_BIT = 8
const m_CHAR_MAX = 255
const m_CHAR_MIN = 0
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
const m_COLL_WEIGHTS_MAX = 2
const m_CPU_SETSIZE = 1024
const m_CSIGNAL = 0x000000ff
const m_DELAYTIMER_MAX = 0x7fffffff
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
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_EXPR_NEST_MAX = 32
const m_FALLOC_FL_KEEP_SIZE = 1
const m_FALLOC_FL_PUNCH_HOLE = 2
const m_FAPPEND = "O_APPEND"
const m_FASYNC = "O_ASYNC"
const m_FD_CLOEXEC = 1
const m_FD_SETSIZE = 1024
const m_FFSYNC = "O_SYNC"
const m_FILENAME_MAX = 4096
const m_FILESIZEBITS = 64
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
const m_F_SETLK = 13
const m_F_SETLKW = 14
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
const m_HASXDMAUTH = 1
const m_HAVE_ABSTRACT_SOCKETS = 1
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_GETADDRINFO = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_SENDMSG = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDIO_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_HAVE_WCHAR_H = 1
const m_HOST_NAME_MAX = 255
const m_INET6_ADDRSTRLEN = 46
const m_INET_ADDRSTRLEN = 16
const m_INT16_MAX = 0x7fff
const m_INT32_MAX = 0x7fffffff
const m_INT64_MAX = 0x7fffffffffffffff
const m_INT8_MAX = 0x7f
const m_INTMAX_MAX = "INT64_MAX"
const m_INTMAX_MIN = "INT64_MIN"
const m_INTPTR_MAX = "INT32_MAX"
const m_INTPTR_MIN = "INT32_MIN"
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
const m_IN_CLASSA_MAX = 128
const m_IN_CLASSA_NET = 0xff000000
const m_IN_CLASSA_NSHIFT = 24
const m_IN_CLASSB_MAX = 65536
const m_IN_CLASSB_NET = 0xffff0000
const m_IN_CLASSB_NSHIFT = 16
const m_IN_CLASSC_NET = 0xffffff00
const m_IN_CLASSC_NSHIFT = 8
const m_IN_LOOPBACKNET = 127
const m_IOV_MAX = 1024
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
const m_LINE_MAX = 4096
const m_LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const m_LLONG_MAX = 0x7fffffffffffffff
const m_LOGIN_NAME_MAX = 256
const m_LONG_BIT = 32
const m_LONG_MAX = "__LONG_MAX"
const m_LT_OBJDIR = ".libs/"
const m_L_INCR = 1
const m_L_SET = 0
const m_L_XTND = 2
const m_L_ctermid = 20
const m_L_cuserid = 20
const m_L_tmpnam = 20
const m_MAX_HANDLE_SZ = 128
const m_MB_LEN_MAX = 4
const m_MCAST_BLOCK_SOURCE = 43
const m_MCAST_EXCLUDE = 0
const m_MCAST_INCLUDE = 1
const m_MCAST_JOIN_GROUP = 42
const m_MCAST_JOIN_SOURCE_GROUP = 46
const m_MCAST_LEAVE_GROUP = 45
const m_MCAST_LEAVE_SOURCE_GROUP = 47
const m_MCAST_MSFILTER = 48
const m_MCAST_UNBLOCK_SOURCE = 44
const m_MQ_PRIO_MAX = 32768
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
const m_NAME_MAX = 255
const m_NDEBUG = 1
const m_NGROUPS_MAX = 32
const m_NL_ARGMAX = 9
const m_NL_LANGMAX = 32
const m_NL_MSGMAX = 32767
const m_NL_NMAX = 16
const m_NL_SETMAX = 255
const m_NL_TEXTMAX = 2048
const m_NZERO = 20
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
const m_O_NONBLOCK = 2048
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
const m_PACKAGE = "libxcb"
const m_PACKAGE_BUGREPORT = "https://gitlab.freedesktop.org/xorg/lib/libxcb/issues"
const m_PACKAGE_NAME = "libxcb"
const m_PACKAGE_STRING = "libxcb 1.15"
const m_PACKAGE_TARNAME = "libxcb"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "1.15"
const m_PACKAGE_VERSION_MAJOR = 1
const m_PACKAGE_VERSION_MINOR = 15
const m_PACKAGE_VERSION_PATCHLEVEL = 0
const m_PAGESIZE = 4096
const m_PAGE_SIZE = "PAGESIZE"
const m_PATH_MAX = 4096
const m_PDP_ENDIAN = "__PDP_ENDIAN"
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
const m_PIPE_BUF = 4096
const m_POLLERR = 0x008
const m_POLLHUP = 0x010
const m_POLLIN = 1
const m_POLLMSG = 0x400
const m_POLLNVAL = 0x020
const m_POLLOUT = 4
const m_POLLPRI = 0x002
const m_POLLRDBAND = 0x080
const m_POLLRDHUP = 0x2000
const m_POLLRDNORM = 0x040
const m_POLLWRBAND = 0x200
const m_POLLWRNORM = 0x100
const m_POSIX_CLOSE_RESTART = 0
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
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
const m_PTHREAD_CANCEL_ASYNCHRONOUS = 1
const m_PTHREAD_CANCEL_DEFERRED = 0
const m_PTHREAD_CANCEL_DISABLE = 1
const m_PTHREAD_CANCEL_ENABLE = 0
const m_PTHREAD_CANCEL_MASKED = 2
const m_PTHREAD_CREATE_DETACHED = 1
const m_PTHREAD_CREATE_JOINABLE = 0
const m_PTHREAD_DESTRUCTOR_ITERATIONS = 4
const m_PTHREAD_EXPLICIT_SCHED = 1
const m_PTHREAD_INHERIT_SCHED = 0
const m_PTHREAD_KEYS_MAX = 128
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
const m_PTHREAD_STACK_MIN = 2048
const m_PTRDIFF_MAX = "INT32_MAX"
const m_PTRDIFF_MIN = "INT32_MIN"
const m_P_tmpdir = "/tmp"
const m_RAND_MAX = 0x7fffffff
const m_RE_DUP_MAX = 255
const m_RWF_APPEND = 0x00000010
const m_RWF_DSYNC = 0x00000002
const m_RWF_HIPRI = 0x00000001
const m_RWF_NOWAIT = 0x00000008
const m_RWF_SYNC = 0x00000004
const m_RWF_WRITE_LIFE_NOT_SET = 0
const m_RWH_WRITE_LIFE_EXTREME = 5
const m_RWH_WRITE_LIFE_LONG = 4
const m_RWH_WRITE_LIFE_MEDIUM = 3
const m_RWH_WRITE_LIFE_NONE = 1
const m_RWH_WRITE_LIFE_SHORT = 2
const m_R_OK = 4
const m_SCHAR_MAX = 127
const m_SCHED_BATCH = 3
const m_SCHED_DEADLINE = 6
const m_SCHED_FIFO = 1
const m_SCHED_IDLE = 5
const m_SCHED_OTHER = 0
const m_SCHED_RESET_ON_FORK = 0x40000000
const m_SCHED_RR = 2
const m_SCM_CREDENTIALS = 0x02
const m_SCM_RIGHTS = 1
const m_SCM_TIMESTAMP = "SO_TIMESTAMP"
const m_SCM_TIMESTAMPING = "SO_TIMESTAMPING"
const m_SCM_TIMESTAMPING_OPT_STATS = 54
const m_SCM_TIMESTAMPING_PKTINFO = 58
const m_SCM_TIMESTAMPNS = "SO_TIMESTAMPNS"
const m_SCM_TXTIME = "SO_TXTIME"
const m_SCM_WIFI_STATUS = "SO_WIFI_STATUS"
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
const m_SEEK_DATA = 3
const m_SEEK_HOLE = 4
const m_SEM_NSEMS_MAX = 256
const m_SEM_VALUE_MAX = 0x7fffffff
const m_SHRT_MAX = 0x7fff
const m_SHUT_RD = 0
const m_SHUT_RDWR = 2
const m_SHUT_WR = 1
const m_SIG_ATOMIC_MAX = "INT32_MAX"
const m_SIG_ATOMIC_MIN = "INT32_MIN"
const m_SIZE_MAX = "UINT32_MAX"
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
const m_SO_RCVTIMEO = 66
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
const m_SO_SNDTIMEO = 67
const m_SO_TIMESTAMP = 63
const m_SO_TIMESTAMPING = 65
const m_SO_TIMESTAMPNS = 64
const m_SO_TXTIME = 61
const m_SO_TYPE = 3
const m_SO_WIFI_STATUS = 41
const m_SO_ZEROCOPY = 60
const m_SPLICE_F_GIFT = 8
const m_SPLICE_F_MORE = 4
const m_SPLICE_F_MOVE = 1
const m_SPLICE_F_NONBLOCK = 2
const m_SSIZE_MAX = "LONG_MAX"
const m_STDC_HEADERS = 1
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_SYMLOOP_MAX = 40
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
const m_TIMER_ABSTIME = 1
const m_TIME_UTC = 1
const m_TMP_MAX = 10000
const m_TTY_NAME_MAX = 32
const m_TZNAME_MAX = 6
const m_UCHAR_MAX = 255
const m_UINT16_MAX = 0xffff
const m_UINT32_MAX = "0xffffffffu"
const m_UINT64_MAX = "0xffffffffffffffffu"
const m_UINT8_MAX = 0xff
const m_UINTMAX_MAX = "UINT64_MAX"
const m_UINTPTR_MAX = "UINT32_MAX"
const m_UINT_FAST16_MAX = "UINT32_MAX"
const m_UINT_FAST32_MAX = "UINT32_MAX"
const m_UINT_FAST64_MAX = "UINT64_MAX"
const m_UINT_FAST8_MAX = "UINT8_MAX"
const m_UINT_LEAST16_MAX = "UINT16_MAX"
const m_UINT_LEAST32_MAX = "UINT32_MAX"
const m_UINT_LEAST64_MAX = "UINT64_MAX"
const m_UINT_LEAST8_MAX = "UINT8_MAX"
const m_UINT_MAX = 0xffffffff
const m_UIO_MAXIOV = 1024
const m_USE_POLL = 1
const m_USHRT_MAX = 0xffff
const m_VERSION = "1.15"
const m_WINT_MAX = "UINT32_MAX"
const m_WINT_MIN = 0
const m_WNOHANG = 1
const m_WORD_BIT = 32
const m_WUNTRACED = 2
const m_W_OK = 2
const m_XCB_ACCESS = 10
const m_XCB_ALLOC = 11
const m_XCB_ALLOC_COLOR = 84
const m_XCB_ALLOC_COLOR_CELLS = 86
const m_XCB_ALLOC_COLOR_PLANES = 87
const m_XCB_ALLOC_NAMED_COLOR = 85
const m_XCB_ALLOW_EVENTS = 35
const m_XCB_ATOM = 5
const m_XCB_BELL = 104
const m_XCB_BIGREQUESTS_MAJOR_VERSION = 0
const m_XCB_BIGREQUESTS_MINOR_VERSION = 0
const m_XCB_BIG_REQUESTS_ENABLE = 0
const m_XCB_BUTTON_PRESS = 4
const m_XCB_BUTTON_RELEASE = 5
const m_XCB_CHANGE_ACTIVE_POINTER_GRAB = 30
const m_XCB_CHANGE_GC = 56
const m_XCB_CHANGE_HOSTS = 109
const m_XCB_CHANGE_KEYBOARD_CONTROL = 102
const m_XCB_CHANGE_KEYBOARD_MAPPING = 100
const m_XCB_CHANGE_POINTER_CONTROL = 105
const m_XCB_CHANGE_PROPERTY = 18
const m_XCB_CHANGE_SAVE_SET = 6
const m_XCB_CHANGE_WINDOW_ATTRIBUTES = 2
const m_XCB_CIRCULATE_NOTIFY = 26
const m_XCB_CIRCULATE_REQUEST = 27
const m_XCB_CIRCULATE_WINDOW = 13
const m_XCB_CLEAR_AREA = 61
const m_XCB_CLIENT_MESSAGE = 33
const m_XCB_CLOSE_FONT = 46
const m_XCB_COLORMAP = 12
const m_XCB_COLORMAP_NOTIFY = 32
const m_XCB_CONFIGURE_NOTIFY = 22
const m_XCB_CONFIGURE_REQUEST = 23
const m_XCB_CONFIGURE_WINDOW = 12
const m_XCB_CONN_CLOSED_EXT_NOTSUPPORTED = 2
const m_XCB_CONN_CLOSED_FDPASSING_FAILED = 7
const m_XCB_CONN_CLOSED_INVALID_SCREEN = 6
const m_XCB_CONN_CLOSED_MEM_INSUFFICIENT = 3
const m_XCB_CONN_CLOSED_PARSE_ERR = 5
const m_XCB_CONN_CLOSED_REQ_LEN_EXCEED = 4
const m_XCB_CONN_ERROR = 1
const m_XCB_CONVERT_SELECTION = 24
const m_XCB_COPY_AREA = 62
const m_XCB_COPY_COLORMAP_AND_FREE = 80
const m_XCB_COPY_FROM_PARENT = 0
const m_XCB_COPY_GC = 57
const m_XCB_COPY_PLANE = 63
const m_XCB_CREATE_COLORMAP = 78
const m_XCB_CREATE_CURSOR = 93
const m_XCB_CREATE_GC = 55
const m_XCB_CREATE_GLYPH_CURSOR = 94
const m_XCB_CREATE_NOTIFY = 16
const m_XCB_CREATE_PIXMAP = 53
const m_XCB_CREATE_WINDOW = 1
const m_XCB_CURRENT_TIME = 0
const m_XCB_CURSOR = 6
const m_XCB_DELETE_PROPERTY = 19
const m_XCB_DESTROY_NOTIFY = 17
const m_XCB_DESTROY_SUBWINDOWS = 5
const m_XCB_DESTROY_WINDOW = 4
const m_XCB_DRAWABLE = 9
const m_XCB_ENTER_NOTIFY = 7
const m_XCB_EXPOSE = 12
const m_XCB_FILL_POLY = 69
const m_XCB_FOCUS_IN = 9
const m_XCB_FOCUS_OUT = 10
const m_XCB_FONT = 7
const m_XCB_FORCE_SCREEN_SAVER = 115
const m_XCB_FREE_COLORMAP = 79
const m_XCB_FREE_COLORS = 88
const m_XCB_FREE_CURSOR = 95
const m_XCB_FREE_GC = 60
const m_XCB_FREE_PIXMAP = 54
const m_XCB_GET_ATOM_NAME = 17
const m_XCB_GET_FONT_PATH = 52
const m_XCB_GET_GEOMETRY = 14
const m_XCB_GET_IMAGE = 73
const m_XCB_GET_INPUT_FOCUS = 43
const m_XCB_GET_KEYBOARD_CONTROL = 103
const m_XCB_GET_KEYBOARD_MAPPING = 101
const m_XCB_GET_MODIFIER_MAPPING = 119
const m_XCB_GET_MOTION_EVENTS = 39
const m_XCB_GET_POINTER_CONTROL = 106
const m_XCB_GET_POINTER_MAPPING = 117
const m_XCB_GET_PROPERTY = 20
const m_XCB_GET_SCREEN_SAVER = 108
const m_XCB_GET_SELECTION_OWNER = 23
const m_XCB_GET_WINDOW_ATTRIBUTES = 3
const m_XCB_GE_GENERIC = 35
const m_XCB_GRAB_BUTTON = 28
const m_XCB_GRAB_KEY = 33
const m_XCB_GRAB_KEYBOARD = 31
const m_XCB_GRAB_POINTER = 26
const m_XCB_GRAB_SERVER = 36
const m_XCB_GRAPHICS_EXPOSURE = 13
const m_XCB_GRAVITY_NOTIFY = 24
const m_XCB_G_CONTEXT = 13
const m_XCB_ID_CHOICE = 14
const m_XCB_IMAGE_TEXT_16 = 77
const m_XCB_IMAGE_TEXT_8 = 76
const m_XCB_IMPLEMENTATION = 17
const m_XCB_INSTALL_COLORMAP = 81
const m_XCB_INTERN_ATOM = 16
const m_XCB_KEYMAP_NOTIFY = 11
const m_XCB_KEY_PRESS = 2
const m_XCB_KEY_RELEASE = 3
const m_XCB_KILL_CLIENT = 113
const m_XCB_LEAVE_NOTIFY = 8
const m_XCB_LENGTH = 16
const m_XCB_LIST_EXTENSIONS = 99
const m_XCB_LIST_FONTS = 49
const m_XCB_LIST_FONTS_WITH_INFO = 50
const m_XCB_LIST_HOSTS = 110
const m_XCB_LIST_INSTALLED_COLORMAPS = 83
const m_XCB_LIST_PROPERTIES = 21
const m_XCB_LOOKUP_COLOR = 92
const m_XCB_MAPPING_NOTIFY = 34
const m_XCB_MAP_NOTIFY = 19
const m_XCB_MAP_REQUEST = 20
const m_XCB_MAP_SUBWINDOWS = 9
const m_XCB_MAP_WINDOW = 8
const m_XCB_MATCH = 8
const m_XCB_MAX_PASS_FD = 16
const m_XCB_MOTION_NOTIFY = 6
const m_XCB_NAME = 15
const m_XCB_NONE = 0
const m_XCB_NO_EXPOSURE = 14
const m_XCB_NO_OPERATION = 127
const m_XCB_NO_SYMBOL = 0
const m_XCB_OPEN_FONT = 45
const m_XCB_PIXMAP = 4
const m_XCB_POLY_ARC = 68
const m_XCB_POLY_FILL_ARC = 71
const m_XCB_POLY_FILL_RECTANGLE = 70
const m_XCB_POLY_LINE = 65
const m_XCB_POLY_POINT = 64
const m_XCB_POLY_RECTANGLE = 67
const m_XCB_POLY_SEGMENT = 66
const m_XCB_POLY_TEXT_16 = 75
const m_XCB_POLY_TEXT_8 = 74
const m_XCB_PROPERTY_NOTIFY = 28
const m_XCB_PUT_IMAGE = 72
const m_XCB_QUERY_BEST_SIZE = 97
const m_XCB_QUERY_COLORS = 91
const m_XCB_QUERY_EXTENSION = 98
const m_XCB_QUERY_FONT = 47
const m_XCB_QUERY_KEYMAP = 44
const m_XCB_QUERY_POINTER = 38
const m_XCB_QUERY_TEXT_EXTENTS = 48
const m_XCB_QUERY_TREE = 15
const m_XCB_QUEUE_BUFFER_SIZE = 16384
const m_XCB_RECOLOR_CURSOR = 96
const m_XCB_REPARENT_NOTIFY = 21
const m_XCB_REPARENT_WINDOW = 7
const m_XCB_REQUEST = 1
const m_XCB_RESIZE_REQUEST = 25
const m_XCB_ROTATE_PROPERTIES = 114
const m_XCB_SELECTION_CLEAR = 29
const m_XCB_SELECTION_NOTIFY = 31
const m_XCB_SELECTION_REQUEST = 30
const m_XCB_SEND_EVENT = 25
const m_XCB_SET_ACCESS_CONTROL = 111
const m_XCB_SET_CLIP_RECTANGLES = 59
const m_XCB_SET_CLOSE_DOWN_MODE = 112
const m_XCB_SET_DASHES = 58
const m_XCB_SET_FONT_PATH = 51
const m_XCB_SET_INPUT_FOCUS = 42
const m_XCB_SET_MODIFIER_MAPPING = 118
const m_XCB_SET_POINTER_MAPPING = 116
const m_XCB_SET_SCREEN_SAVER = 107
const m_XCB_SET_SELECTION_OWNER = 22
const m_XCB_STORE_COLORS = 89
const m_XCB_STORE_NAMED_COLOR = 90
const m_XCB_TRANSLATE_COORDINATES = 40
const m_XCB_UNGRAB_BUTTON = 29
const m_XCB_UNGRAB_KEY = 34
const m_XCB_UNGRAB_KEYBOARD = 32
const m_XCB_UNGRAB_POINTER = 27
const m_XCB_UNGRAB_SERVER = 37
const m_XCB_UNINSTALL_COLORMAP = 82
const m_XCB_UNMAP_NOTIFY = 18
const m_XCB_UNMAP_SUBWINDOWS = 11
const m_XCB_UNMAP_WINDOW = 10
const m_XCB_VALUE = 2
const m_XCB_VISIBILITY_NOTIFY = 15
const m_XCB_WARP_POINTER = 41
const m_XCB_WINDOW = 3
const m_X_OK = 1
const m_X_PROTOCOL = 11
const m_X_PROTOCOL_REVISION = 0
const m_X_TCP_PORT = 6000
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
const m__FILE_OFFSET_BITS = 64
const m__GNU_SOURCE = 1
const m__HPUX_ALT_XOPEN_SOCKET_API = 1
const m__ILP32 = 1
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
const m__NETBSD_SOURCE = 1
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
const m__TANDEM_SOURCE = 1
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
const m___PRI64 = "ll"
const m___PRIPTR = ""
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
const m___STDC_WANT_IEC_60559_ATTRIBS_EXT__ = 1
const m___STDC_WANT_IEC_60559_BFP_EXT__ = 1
const m___STDC_WANT_IEC_60559_DFP_EXT__ = 1
const m___STDC_WANT_IEC_60559_FUNCS_EXT__ = 1
const m___STDC_WANT_IEC_60559_TYPES_EXT__ = 1
const m___STDC_WANT_LIB_EXT2__ = 1
const m___STDC_WANT_MATH_SPEC_FUNCS__ = 1
const m___STDC__ = 1
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
const m___tm_gmtoff = "tm_gmtoff"
const m___tm_zone = "tm_zone"
const m___unix = 1
const m___unix__ = 1
const m_alloca = "__builtin_alloca"
const m_i386 = 1
const m_linux = 1
const m_loff_t = "off_t"
const m_static_assert = "_Static_assert"
const m_unix = 1

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint32

type t__predefined_wchar_t = int32

type t__predefined_ptrdiff_t = int32

type Tsize_t = uint32

type Tlocale_t = uintptr

type Tssize_t = int32

type Toff_t = int64

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

type Twchar_t = int32

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

type Tmode_t = uint32

type Tpid_t = int32

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

type Tregister_t = int32

type Ttime_t = int64

type Tsuseconds_t = int64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tu_int64_t = uint64

type Tnlink_t = uint32

type Tino_t = uint64

type Tdev_t = uint64

type Tblksize_t = int32

type Tblkcnt_t = int64

type Tfsblkcnt_t = uint64

type Tfsfilcnt_t = uint64

type Ttimer_t = uintptr

type Tclockid_t = int32

type Tclock_t = int32

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

type Tuintptr_t = uint32

type Tintptr_t = int32

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

type Tsched_param = struct {
	Fsched_priority int32
	F__reserved1    int32
	F__reserved2    [4]int32
	F__reserved3    int32
}

type Tcpu_set_t = struct {
	F__bits [32]uint32
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
	Ftm_gmtoff int32
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
	F__bits [32]uint32
}

type Txcb_connection_t = struct {
	Fhas_error int32
	Fsetup     uintptr
	Ffd        int32
	Fiolock    Tpthread_mutex_t
	Fin        T_xcb_in
	Fout       T_xcb_out
	Fext       T_xcb_ext
	Fxid       T_xcb_xid
}

type Txcb_generic_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_generic_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
}

type Txcb_generic_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fpad           [7]Tuint32_t
	Ffull_sequence Tuint32_t
}

type Txcb_raw_generic_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fpad           [7]Tuint32_t
}

type Txcb_ge_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fevent_type    Tuint16_t
	Fpad1          Tuint16_t
	Fpad           [5]Tuint32_t
	Ffull_sequence Tuint32_t
}

type Txcb_generic_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fresource_id   Tuint32_t
	Fminor_code    Tuint16_t
	Fmajor_code    Tuint8_t
	Fpad0          Tuint8_t
	Fpad           [5]Tuint32_t
	Ffull_sequence Tuint32_t
}

type Txcb_void_cookie_t = struct {
	Fsequence uint32
}

type Txcb_char2b_t = struct {
	Fbyte1 Tuint8_t
	Fbyte2 Tuint8_t
}

type Txcb_char2b_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_window_t = uint32

type Txcb_window_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_pixmap_t = uint32

type Txcb_pixmap_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_cursor_t = uint32

type Txcb_cursor_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_font_t = uint32

type Txcb_font_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_gcontext_t = uint32

type Txcb_gcontext_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_colormap_t = uint32

type Txcb_colormap_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_atom_t = uint32

type Txcb_atom_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_drawable_t = uint32

type Txcb_drawable_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_fontable_t = uint32

type Txcb_fontable_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_bool32_t = uint32

type Txcb_bool32_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_visualid_t = uint32

type Txcb_visualid_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_timestamp_t = uint32

type Txcb_timestamp_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_keysym_t = uint32

type Txcb_keysym_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_keycode_t = uint8

type Txcb_keycode_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_keycode32_t = uint32

type Txcb_keycode32_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_button_t = uint8

type Txcb_button_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_point_t = struct {
	Fx Tint16_t
	Fy Tint16_t
}

type Txcb_point_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_rectangle_t = struct {
	Fx      Tint16_t
	Fy      Tint16_t
	Fwidth  Tuint16_t
	Fheight Tuint16_t
}

type Txcb_rectangle_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_arc_t = struct {
	Fx      Tint16_t
	Fy      Tint16_t
	Fwidth  Tuint16_t
	Fheight Tuint16_t
	Fangle1 Tint16_t
	Fangle2 Tint16_t
}

type Txcb_arc_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_format_t = struct {
	Fdepth          Tuint8_t
	Fbits_per_pixel Tuint8_t
	Fscanline_pad   Tuint8_t
	Fpad0           [5]Tuint8_t
}

type Txcb_format_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_visual_class_t = int32

type _xcb_visual_class_t = int32

const _XCB_VISUAL_CLASS_STATIC_GRAY = 0
const _XCB_VISUAL_CLASS_GRAY_SCALE = 1
const _XCB_VISUAL_CLASS_STATIC_COLOR = 2
const _XCB_VISUAL_CLASS_PSEUDO_COLOR = 3
const _XCB_VISUAL_CLASS_TRUE_COLOR = 4
const _XCB_VISUAL_CLASS_DIRECT_COLOR = 5

type Txcb_visualtype_t = struct {
	Fvisual_id          Txcb_visualid_t
	F_class             Tuint8_t
	Fbits_per_rgb_value Tuint8_t
	Fcolormap_entries   Tuint16_t
	Fred_mask           Tuint32_t
	Fgreen_mask         Tuint32_t
	Fblue_mask          Tuint32_t
	Fpad0               [4]Tuint8_t
}

type Txcb_visualtype_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_depth_t = struct {
	Fdepth       Tuint8_t
	Fpad0        Tuint8_t
	Fvisuals_len Tuint16_t
	Fpad1        [4]Tuint8_t
}

type Txcb_depth_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_event_mask_t = int32

type _xcb_event_mask_t = int32

const _XCB_EVENT_MASK_NO_EVENT = 0
const _XCB_EVENT_MASK_KEY_PRESS = 1
const _XCB_EVENT_MASK_KEY_RELEASE = 2
const _XCB_EVENT_MASK_BUTTON_PRESS = 4
const _XCB_EVENT_MASK_BUTTON_RELEASE = 8
const _XCB_EVENT_MASK_ENTER_WINDOW = 16
const _XCB_EVENT_MASK_LEAVE_WINDOW = 32
const _XCB_EVENT_MASK_POINTER_MOTION = 64
const _XCB_EVENT_MASK_POINTER_MOTION_HINT = 128
const _XCB_EVENT_MASK_BUTTON_1_MOTION = 256
const _XCB_EVENT_MASK_BUTTON_2_MOTION = 512
const _XCB_EVENT_MASK_BUTTON_3_MOTION = 1024
const _XCB_EVENT_MASK_BUTTON_4_MOTION = 2048
const _XCB_EVENT_MASK_BUTTON_5_MOTION = 4096
const _XCB_EVENT_MASK_BUTTON_MOTION = 8192
const _XCB_EVENT_MASK_KEYMAP_STATE = 16384
const _XCB_EVENT_MASK_EXPOSURE = 32768
const _XCB_EVENT_MASK_VISIBILITY_CHANGE = 65536
const _XCB_EVENT_MASK_STRUCTURE_NOTIFY = 131072
const _XCB_EVENT_MASK_RESIZE_REDIRECT = 262144
const _XCB_EVENT_MASK_SUBSTRUCTURE_NOTIFY = 524288
const _XCB_EVENT_MASK_SUBSTRUCTURE_REDIRECT = 1048576
const _XCB_EVENT_MASK_FOCUS_CHANGE = 2097152
const _XCB_EVENT_MASK_PROPERTY_CHANGE = 4194304
const _XCB_EVENT_MASK_COLOR_MAP_CHANGE = 8388608
const _XCB_EVENT_MASK_OWNER_GRAB_BUTTON = 16777216

type Txcb_backing_store_t = int32

type _xcb_backing_store_t = int32

const _XCB_BACKING_STORE_NOT_USEFUL = 0
const _XCB_BACKING_STORE_WHEN_MAPPED = 1
const _XCB_BACKING_STORE_ALWAYS = 2

type Txcb_screen_t = struct {
	Froot                  Txcb_window_t
	Fdefault_colormap      Txcb_colormap_t
	Fwhite_pixel           Tuint32_t
	Fblack_pixel           Tuint32_t
	Fcurrent_input_masks   Tuint32_t
	Fwidth_in_pixels       Tuint16_t
	Fheight_in_pixels      Tuint16_t
	Fwidth_in_millimeters  Tuint16_t
	Fheight_in_millimeters Tuint16_t
	Fmin_installed_maps    Tuint16_t
	Fmax_installed_maps    Tuint16_t
	Froot_visual           Txcb_visualid_t
	Fbacking_stores        Tuint8_t
	Fsave_unders           Tuint8_t
	Froot_depth            Tuint8_t
	Fallowed_depths_len    Tuint8_t
}

type Txcb_screen_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_setup_request_t = struct {
	Fbyte_order                      Tuint8_t
	Fpad0                            Tuint8_t
	Fprotocol_major_version          Tuint16_t
	Fprotocol_minor_version          Tuint16_t
	Fauthorization_protocol_name_len Tuint16_t
	Fauthorization_protocol_data_len Tuint16_t
	Fpad1                            [2]Tuint8_t
}

type Txcb_setup_request_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_setup_failed_t = struct {
	Fstatus                 Tuint8_t
	Freason_len             Tuint8_t
	Fprotocol_major_version Tuint16_t
	Fprotocol_minor_version Tuint16_t
	Flength                 Tuint16_t
}

type Txcb_setup_failed_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_setup_authenticate_t = struct {
	Fstatus Tuint8_t
	Fpad0   [5]Tuint8_t
	Flength Tuint16_t
}

type Txcb_setup_authenticate_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_image_order_t = int32

type _xcb_image_order_t = int32

const _XCB_IMAGE_ORDER_LSB_FIRST = 0
const _XCB_IMAGE_ORDER_MSB_FIRST = 1

type Txcb_setup_t = struct {
	Fstatus                      Tuint8_t
	Fpad0                        Tuint8_t
	Fprotocol_major_version      Tuint16_t
	Fprotocol_minor_version      Tuint16_t
	Flength                      Tuint16_t
	Frelease_number              Tuint32_t
	Fresource_id_base            Tuint32_t
	Fresource_id_mask            Tuint32_t
	Fmotion_buffer_size          Tuint32_t
	Fvendor_len                  Tuint16_t
	Fmaximum_request_length      Tuint16_t
	Froots_len                   Tuint8_t
	Fpixmap_formats_len          Tuint8_t
	Fimage_byte_order            Tuint8_t
	Fbitmap_format_bit_order     Tuint8_t
	Fbitmap_format_scanline_unit Tuint8_t
	Fbitmap_format_scanline_pad  Tuint8_t
	Fmin_keycode                 Txcb_keycode_t
	Fmax_keycode                 Txcb_keycode_t
	Fpad1                        [4]Tuint8_t
}

type Txcb_setup_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_mod_mask_t = int32

type _xcb_mod_mask_t = int32

const _XCB_MOD_MASK_SHIFT = 1
const _XCB_MOD_MASK_LOCK = 2
const _XCB_MOD_MASK_CONTROL = 4
const _XCB_MOD_MASK_1 = 8
const _XCB_MOD_MASK_2 = 16
const _XCB_MOD_MASK_3 = 32
const _XCB_MOD_MASK_4 = 64
const _XCB_MOD_MASK_5 = 128
const _XCB_MOD_MASK_ANY = 32768

type Txcb_key_but_mask_t = int32

type _xcb_key_but_mask_t = int32

const _XCB_KEY_BUT_MASK_SHIFT = 1
const _XCB_KEY_BUT_MASK_LOCK = 2
const _XCB_KEY_BUT_MASK_CONTROL = 4
const _XCB_KEY_BUT_MASK_MOD_1 = 8
const _XCB_KEY_BUT_MASK_MOD_2 = 16
const _XCB_KEY_BUT_MASK_MOD_3 = 32
const _XCB_KEY_BUT_MASK_MOD_4 = 64
const _XCB_KEY_BUT_MASK_MOD_5 = 128
const _XCB_KEY_BUT_MASK_BUTTON_1 = 256
const _XCB_KEY_BUT_MASK_BUTTON_2 = 512
const _XCB_KEY_BUT_MASK_BUTTON_3 = 1024
const _XCB_KEY_BUT_MASK_BUTTON_4 = 2048
const _XCB_KEY_BUT_MASK_BUTTON_5 = 4096

type Txcb_window_enum_t = int32

type _xcb_window_enum_t = int32

const _XCB_WINDOW_NONE = 0

type Txcb_key_press_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Txcb_keycode_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Froot          Txcb_window_t
	Fevent         Txcb_window_t
	Fchild         Txcb_window_t
	Froot_x        Tint16_t
	Froot_y        Tint16_t
	Fevent_x       Tint16_t
	Fevent_y       Tint16_t
	Fstate         Tuint16_t
	Fsame_screen   Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_key_release_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Txcb_keycode_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Froot          Txcb_window_t
	Fevent         Txcb_window_t
	Fchild         Txcb_window_t
	Froot_x        Tint16_t
	Froot_y        Tint16_t
	Fevent_x       Tint16_t
	Fevent_y       Tint16_t
	Fstate         Tuint16_t
	Fsame_screen   Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_key_press_event_t1 = Txcb_key_release_event_t

type Txcb_button_mask_t = int32

type _xcb_button_mask_t = int32

const _XCB_BUTTON_MASK_1 = 256
const _XCB_BUTTON_MASK_2 = 512
const _XCB_BUTTON_MASK_3 = 1024
const _XCB_BUTTON_MASK_4 = 2048
const _XCB_BUTTON_MASK_5 = 4096
const _XCB_BUTTON_MASK_ANY = 32768

type Txcb_button_press_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Txcb_button_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Froot          Txcb_window_t
	Fevent         Txcb_window_t
	Fchild         Txcb_window_t
	Froot_x        Tint16_t
	Froot_y        Tint16_t
	Fevent_x       Tint16_t
	Fevent_y       Tint16_t
	Fstate         Tuint16_t
	Fsame_screen   Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_button_release_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Txcb_button_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Froot          Txcb_window_t
	Fevent         Txcb_window_t
	Fchild         Txcb_window_t
	Froot_x        Tint16_t
	Froot_y        Tint16_t
	Fevent_x       Tint16_t
	Fevent_y       Tint16_t
	Fstate         Tuint16_t
	Fsame_screen   Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_button_press_event_t1 = Txcb_button_release_event_t

type Txcb_motion_t = int32

type _xcb_motion_t = int32

const _XCB_MOTION_NORMAL = 0
const _XCB_MOTION_HINT = 1

type Txcb_motion_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Tuint8_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Froot          Txcb_window_t
	Fevent         Txcb_window_t
	Fchild         Txcb_window_t
	Froot_x        Tint16_t
	Froot_y        Tint16_t
	Fevent_x       Tint16_t
	Fevent_y       Tint16_t
	Fstate         Tuint16_t
	Fsame_screen   Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_notify_detail_t = int32

type _xcb_notify_detail_t = int32

const _XCB_NOTIFY_DETAIL_ANCESTOR = 0
const _XCB_NOTIFY_DETAIL_VIRTUAL = 1
const _XCB_NOTIFY_DETAIL_INFERIOR = 2
const _XCB_NOTIFY_DETAIL_NONLINEAR = 3
const _XCB_NOTIFY_DETAIL_NONLINEAR_VIRTUAL = 4
const _XCB_NOTIFY_DETAIL_POINTER = 5
const _XCB_NOTIFY_DETAIL_POINTER_ROOT = 6
const _XCB_NOTIFY_DETAIL_NONE = 7

type Txcb_notify_mode_t = int32

type _xcb_notify_mode_t = int32

const _XCB_NOTIFY_MODE_NORMAL = 0
const _XCB_NOTIFY_MODE_GRAB = 1
const _XCB_NOTIFY_MODE_UNGRAB = 2
const _XCB_NOTIFY_MODE_WHILE_GRABBED = 3

type Txcb_enter_notify_event_t = struct {
	Fresponse_type     Tuint8_t
	Fdetail            Tuint8_t
	Fsequence          Tuint16_t
	Ftime              Txcb_timestamp_t
	Froot              Txcb_window_t
	Fevent             Txcb_window_t
	Fchild             Txcb_window_t
	Froot_x            Tint16_t
	Froot_y            Tint16_t
	Fevent_x           Tint16_t
	Fevent_y           Tint16_t
	Fstate             Tuint16_t
	Fmode              Tuint8_t
	Fsame_screen_focus Tuint8_t
}

type Txcb_leave_notify_event_t = struct {
	Fresponse_type     Tuint8_t
	Fdetail            Tuint8_t
	Fsequence          Tuint16_t
	Ftime              Txcb_timestamp_t
	Froot              Txcb_window_t
	Fevent             Txcb_window_t
	Fchild             Txcb_window_t
	Froot_x            Tint16_t
	Froot_y            Tint16_t
	Fevent_x           Tint16_t
	Fevent_y           Tint16_t
	Fstate             Tuint16_t
	Fmode              Tuint8_t
	Fsame_screen_focus Tuint8_t
}

type Txcb_enter_notify_event_t1 = Txcb_leave_notify_event_t

type Txcb_focus_in_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Tuint8_t
	Fsequence      Tuint16_t
	Fevent         Txcb_window_t
	Fmode          Tuint8_t
	Fpad0          [3]Tuint8_t
}

type Txcb_focus_out_event_t = struct {
	Fresponse_type Tuint8_t
	Fdetail        Tuint8_t
	Fsequence      Tuint16_t
	Fevent         Txcb_window_t
	Fmode          Tuint8_t
	Fpad0          [3]Tuint8_t
}

type Txcb_focus_in_event_t1 = Txcb_focus_out_event_t

type Txcb_keymap_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fkeys          [31]Tuint8_t
}

type Txcb_expose_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fwindow        Txcb_window_t
	Fx             Tuint16_t
	Fy             Tuint16_t
	Fwidth         Tuint16_t
	Fheight        Tuint16_t
	Fcount         Tuint16_t
	Fpad1          [2]Tuint8_t
}

type Txcb_graphics_exposure_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fdrawable      Txcb_drawable_t
	Fx             Tuint16_t
	Fy             Tuint16_t
	Fwidth         Tuint16_t
	Fheight        Tuint16_t
	Fminor_opcode  Tuint16_t
	Fcount         Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad1          [3]Tuint8_t
}

type Txcb_no_exposure_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fdrawable      Txcb_drawable_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad1          Tuint8_t
}

type Txcb_visibility_t = int32

type _xcb_visibility_t = int32

const _XCB_VISIBILITY_UNOBSCURED = 0
const _XCB_VISIBILITY_PARTIALLY_OBSCURED = 1
const _XCB_VISIBILITY_FULLY_OBSCURED = 2

type Txcb_visibility_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fwindow        Txcb_window_t
	Fstate         Tuint8_t
	Fpad1          [3]Tuint8_t
}

type Txcb_create_notify_event_t = struct {
	Fresponse_type     Tuint8_t
	Fpad0              Tuint8_t
	Fsequence          Tuint16_t
	Fparent            Txcb_window_t
	Fwindow            Txcb_window_t
	Fx                 Tint16_t
	Fy                 Tint16_t
	Fwidth             Tuint16_t
	Fheight            Tuint16_t
	Fborder_width      Tuint16_t
	Foverride_redirect Tuint8_t
	Fpad1              Tuint8_t
}

type Txcb_destroy_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fevent         Txcb_window_t
	Fwindow        Txcb_window_t
}

type Txcb_unmap_notify_event_t = struct {
	Fresponse_type  Tuint8_t
	Fpad0           Tuint8_t
	Fsequence       Tuint16_t
	Fevent          Txcb_window_t
	Fwindow         Txcb_window_t
	Ffrom_configure Tuint8_t
	Fpad1           [3]Tuint8_t
}

type Txcb_map_notify_event_t = struct {
	Fresponse_type     Tuint8_t
	Fpad0              Tuint8_t
	Fsequence          Tuint16_t
	Fevent             Txcb_window_t
	Fwindow            Txcb_window_t
	Foverride_redirect Tuint8_t
	Fpad1              [3]Tuint8_t
}

type Txcb_map_request_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fparent        Txcb_window_t
	Fwindow        Txcb_window_t
}

type Txcb_reparent_notify_event_t = struct {
	Fresponse_type     Tuint8_t
	Fpad0              Tuint8_t
	Fsequence          Tuint16_t
	Fevent             Txcb_window_t
	Fwindow            Txcb_window_t
	Fparent            Txcb_window_t
	Fx                 Tint16_t
	Fy                 Tint16_t
	Foverride_redirect Tuint8_t
	Fpad1              [3]Tuint8_t
}

type Txcb_configure_notify_event_t = struct {
	Fresponse_type     Tuint8_t
	Fpad0              Tuint8_t
	Fsequence          Tuint16_t
	Fevent             Txcb_window_t
	Fwindow            Txcb_window_t
	Fabove_sibling     Txcb_window_t
	Fx                 Tint16_t
	Fy                 Tint16_t
	Fwidth             Tuint16_t
	Fheight            Tuint16_t
	Fborder_width      Tuint16_t
	Foverride_redirect Tuint8_t
	Fpad1              Tuint8_t
}

type Txcb_configure_request_event_t = struct {
	Fresponse_type Tuint8_t
	Fstack_mode    Tuint8_t
	Fsequence      Tuint16_t
	Fparent        Txcb_window_t
	Fwindow        Txcb_window_t
	Fsibling       Txcb_window_t
	Fx             Tint16_t
	Fy             Tint16_t
	Fwidth         Tuint16_t
	Fheight        Tuint16_t
	Fborder_width  Tuint16_t
	Fvalue_mask    Tuint16_t
}

type Txcb_gravity_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fevent         Txcb_window_t
	Fwindow        Txcb_window_t
	Fx             Tint16_t
	Fy             Tint16_t
}

type Txcb_resize_request_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fwindow        Txcb_window_t
	Fwidth         Tuint16_t
	Fheight        Tuint16_t
}

type Txcb_place_t = int32

type _xcb_place_t = int32

const _XCB_PLACE_ON_TOP = 0
const _XCB_PLACE_ON_BOTTOM = 1

type Txcb_circulate_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fevent         Txcb_window_t
	Fwindow        Txcb_window_t
	Fpad1          [4]Tuint8_t
	Fplace         Tuint8_t
	Fpad2          [3]Tuint8_t
}

type Txcb_circulate_request_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fevent         Txcb_window_t
	Fwindow        Txcb_window_t
	Fpad1          [4]Tuint8_t
	Fplace         Tuint8_t
	Fpad2          [3]Tuint8_t
}

type Txcb_circulate_notify_event_t1 = Txcb_circulate_request_event_t

type Txcb_property_t = int32

type _xcb_property_t = int32

const _XCB_PROPERTY_NEW_VALUE = 0
const _XCB_PROPERTY_DELETE = 1

type Txcb_property_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fwindow        Txcb_window_t
	Fatom          Txcb_atom_t
	Ftime          Txcb_timestamp_t
	Fstate         Tuint8_t
	Fpad1          [3]Tuint8_t
}

type Txcb_selection_clear_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Fowner         Txcb_window_t
	Fselection     Txcb_atom_t
}

type Txcb_time_t = int32

type _xcb_time_t = int32

const _XCB_TIME_CURRENT_TIME = 0

type Txcb_atom_enum_t = int32

type _xcb_atom_enum_t = int32

const _XCB_ATOM_NONE = 0
const _XCB_ATOM_ANY = 0
const _XCB_ATOM_PRIMARY = 1
const _XCB_ATOM_SECONDARY = 2
const _XCB_ATOM_ARC = 3
const _XCB_ATOM_ATOM = 4
const _XCB_ATOM_BITMAP = 5
const _XCB_ATOM_CARDINAL = 6
const _XCB_ATOM_COLORMAP = 7
const _XCB_ATOM_CURSOR = 8
const _XCB_ATOM_CUT_BUFFER0 = 9
const _XCB_ATOM_CUT_BUFFER1 = 10
const _XCB_ATOM_CUT_BUFFER2 = 11
const _XCB_ATOM_CUT_BUFFER3 = 12
const _XCB_ATOM_CUT_BUFFER4 = 13
const _XCB_ATOM_CUT_BUFFER5 = 14
const _XCB_ATOM_CUT_BUFFER6 = 15
const _XCB_ATOM_CUT_BUFFER7 = 16
const _XCB_ATOM_DRAWABLE = 17
const _XCB_ATOM_FONT = 18
const _XCB_ATOM_INTEGER = 19
const _XCB_ATOM_PIXMAP = 20
const _XCB_ATOM_POINT = 21
const _XCB_ATOM_RECTANGLE = 22
const _XCB_ATOM_RESOURCE_MANAGER = 23
const _XCB_ATOM_RGB_COLOR_MAP = 24
const _XCB_ATOM_RGB_BEST_MAP = 25
const _XCB_ATOM_RGB_BLUE_MAP = 26
const _XCB_ATOM_RGB_DEFAULT_MAP = 27
const _XCB_ATOM_RGB_GRAY_MAP = 28
const _XCB_ATOM_RGB_GREEN_MAP = 29
const _XCB_ATOM_RGB_RED_MAP = 30
const _XCB_ATOM_STRING = 31
const _XCB_ATOM_VISUALID = 32
const _XCB_ATOM_WINDOW = 33
const _XCB_ATOM_WM_COMMAND = 34
const _XCB_ATOM_WM_HINTS = 35
const _XCB_ATOM_WM_CLIENT_MACHINE = 36
const _XCB_ATOM_WM_ICON_NAME = 37
const _XCB_ATOM_WM_ICON_SIZE = 38
const _XCB_ATOM_WM_NAME = 39
const _XCB_ATOM_WM_NORMAL_HINTS = 40
const _XCB_ATOM_WM_SIZE_HINTS = 41
const _XCB_ATOM_WM_ZOOM_HINTS = 42
const _XCB_ATOM_MIN_SPACE = 43
const _XCB_ATOM_NORM_SPACE = 44
const _XCB_ATOM_MAX_SPACE = 45
const _XCB_ATOM_END_SPACE = 46
const _XCB_ATOM_SUPERSCRIPT_X = 47
const _XCB_ATOM_SUPERSCRIPT_Y = 48
const _XCB_ATOM_SUBSCRIPT_X = 49
const _XCB_ATOM_SUBSCRIPT_Y = 50
const _XCB_ATOM_UNDERLINE_POSITION = 51
const _XCB_ATOM_UNDERLINE_THICKNESS = 52
const _XCB_ATOM_STRIKEOUT_ASCENT = 53
const _XCB_ATOM_STRIKEOUT_DESCENT = 54
const _XCB_ATOM_ITALIC_ANGLE = 55
const _XCB_ATOM_X_HEIGHT = 56
const _XCB_ATOM_QUAD_WIDTH = 57
const _XCB_ATOM_WEIGHT = 58
const _XCB_ATOM_POINT_SIZE = 59
const _XCB_ATOM_RESOLUTION = 60
const _XCB_ATOM_COPYRIGHT = 61
const _XCB_ATOM_NOTICE = 62
const _XCB_ATOM_FONT_NAME = 63
const _XCB_ATOM_FAMILY_NAME = 64
const _XCB_ATOM_FULL_NAME = 65
const _XCB_ATOM_CAP_HEIGHT = 66
const _XCB_ATOM_WM_CLASS = 67
const _XCB_ATOM_WM_TRANSIENT_FOR = 68

type Txcb_selection_request_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Fowner         Txcb_window_t
	Frequestor     Txcb_window_t
	Fselection     Txcb_atom_t
	Ftarget        Txcb_atom_t
	Fproperty      Txcb_atom_t
}

type Txcb_selection_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Ftime          Txcb_timestamp_t
	Frequestor     Txcb_window_t
	Fselection     Txcb_atom_t
	Ftarget        Txcb_atom_t
	Fproperty      Txcb_atom_t
}

type Txcb_colormap_state_t = int32

type _xcb_colormap_state_t = int32

const _XCB_COLORMAP_STATE_UNINSTALLED = 0
const _XCB_COLORMAP_STATE_INSTALLED = 1

type Txcb_colormap_enum_t = int32

type _xcb_colormap_enum_t = int32

const _XCB_COLORMAP_NONE = 0

type Txcb_colormap_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Fwindow        Txcb_window_t
	Fcolormap      Txcb_colormap_t
	F_new          Tuint8_t
	Fstate         Tuint8_t
	Fpad1          [2]Tuint8_t
}

type Txcb_client_message_data_t = struct {
	Fdata16 [0][10]Tuint16_t
	Fdata32 [0][5]Tuint32_t
	Fdata8  [20]Tuint8_t
}

type Txcb_client_message_data_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_client_message_event_t = struct {
	Fresponse_type Tuint8_t
	Fformat        Tuint8_t
	Fsequence      Tuint16_t
	Fwindow        Txcb_window_t
	Ftype1         Txcb_atom_t
	Fdata          Txcb_client_message_data_t
}

type Txcb_mapping_t = int32

type _xcb_mapping_t = int32

const _XCB_MAPPING_MODIFIER = 0
const _XCB_MAPPING_KEYBOARD = 1
const _XCB_MAPPING_POINTER = 2

type Txcb_mapping_notify_event_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Frequest       Tuint8_t
	Ffirst_keycode Txcb_keycode_t
	Fcount         Tuint8_t
	Fpad1          Tuint8_t
}

type Txcb_ge_generic_event_t = struct {
	Fresponse_type Tuint8_t
	Fextension     Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fevent_type    Tuint16_t
	Fpad0          [22]Tuint8_t
	Ffull_sequence Tuint32_t
}

type Txcb_request_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_value_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_window_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_value_error_t1 = Txcb_window_error_t

type Txcb_pixmap_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_atom_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_cursor_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_font_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_match_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_request_error_t1 = Txcb_match_error_t

type Txcb_drawable_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_access_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_alloc_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_colormap_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_g_context_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_id_choice_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_name_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_length_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_implementation_error_t = struct {
	Fresponse_type Tuint8_t
	Ferror_code    Tuint8_t
	Fsequence      Tuint16_t
	Fbad_value     Tuint32_t
	Fminor_opcode  Tuint16_t
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
}

type Txcb_window_class_t = int32

type _xcb_window_class_t = int32

const _XCB_WINDOW_CLASS_COPY_FROM_PARENT = 0
const _XCB_WINDOW_CLASS_INPUT_OUTPUT = 1
const _XCB_WINDOW_CLASS_INPUT_ONLY = 2

type Txcb_cw_t = int32

type _xcb_cw_t = int32

const _XCB_CW_BACK_PIXMAP = 1
const _XCB_CW_BACK_PIXEL = 2
const _XCB_CW_BORDER_PIXMAP = 4
const _XCB_CW_BORDER_PIXEL = 8
const _XCB_CW_BIT_GRAVITY = 16
const _XCB_CW_WIN_GRAVITY = 32
const _XCB_CW_BACKING_STORE = 64
const _XCB_CW_BACKING_PLANES = 128
const _XCB_CW_BACKING_PIXEL = 256
const _XCB_CW_OVERRIDE_REDIRECT = 512
const _XCB_CW_SAVE_UNDER = 1024
const _XCB_CW_EVENT_MASK = 2048
const _XCB_CW_DONT_PROPAGATE = 4096
const _XCB_CW_COLORMAP = 8192
const _XCB_CW_CURSOR = 16384

type Txcb_back_pixmap_t = int32

type _xcb_back_pixmap_t = int32

const _XCB_BACK_PIXMAP_NONE = 0
const _XCB_BACK_PIXMAP_PARENT_RELATIVE = 1

type Txcb_gravity_t = int32

type _xcb_gravity_t = int32

const _XCB_GRAVITY_BIT_FORGET = 0
const _XCB_GRAVITY_WIN_UNMAP = 0
const _XCB_GRAVITY_NORTH_WEST = 1
const _XCB_GRAVITY_NORTH = 2
const _XCB_GRAVITY_NORTH_EAST = 3
const _XCB_GRAVITY_WEST = 4
const _XCB_GRAVITY_CENTER = 5
const _XCB_GRAVITY_EAST = 6
const _XCB_GRAVITY_SOUTH_WEST = 7
const _XCB_GRAVITY_SOUTH = 8
const _XCB_GRAVITY_SOUTH_EAST = 9
const _XCB_GRAVITY_STATIC = 10

type Txcb_create_window_value_list_t = struct {
	Fbackground_pixmap     Txcb_pixmap_t
	Fbackground_pixel      Tuint32_t
	Fborder_pixmap         Txcb_pixmap_t
	Fborder_pixel          Tuint32_t
	Fbit_gravity           Tuint32_t
	Fwin_gravity           Tuint32_t
	Fbacking_store         Tuint32_t
	Fbacking_planes        Tuint32_t
	Fbacking_pixel         Tuint32_t
	Foverride_redirect     Txcb_bool32_t
	Fsave_under            Txcb_bool32_t
	Fevent_mask            Tuint32_t
	Fdo_not_propogate_mask Tuint32_t
	Fcolormap              Txcb_colormap_t
	Fcursor                Txcb_cursor_t
}

type Txcb_create_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fdepth        Tuint8_t
	Flength       Tuint16_t
	Fwid          Txcb_window_t
	Fparent       Txcb_window_t
	Fx            Tint16_t
	Fy            Tint16_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
	Fborder_width Tuint16_t
	F_class       Tuint16_t
	Fvisual       Txcb_visualid_t
	Fvalue_mask   Tuint32_t
}

type Txcb_change_window_attributes_value_list_t = struct {
	Fbackground_pixmap     Txcb_pixmap_t
	Fbackground_pixel      Tuint32_t
	Fborder_pixmap         Txcb_pixmap_t
	Fborder_pixel          Tuint32_t
	Fbit_gravity           Tuint32_t
	Fwin_gravity           Tuint32_t
	Fbacking_store         Tuint32_t
	Fbacking_planes        Tuint32_t
	Fbacking_pixel         Tuint32_t
	Foverride_redirect     Txcb_bool32_t
	Fsave_under            Txcb_bool32_t
	Fevent_mask            Tuint32_t
	Fdo_not_propogate_mask Tuint32_t
	Fcolormap              Txcb_colormap_t
	Fcursor                Txcb_cursor_t
}

type Txcb_change_window_attributes_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fvalue_mask   Tuint32_t
}

type Txcb_map_state_t = int32

type _xcb_map_state_t = int32

const _XCB_MAP_STATE_UNMAPPED = 0
const _XCB_MAP_STATE_UNVIEWABLE = 1
const _XCB_MAP_STATE_VIEWABLE = 2

type Txcb_get_window_attributes_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_window_attributes_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_get_window_attributes_reply_t = struct {
	Fresponse_type         Tuint8_t
	Fbacking_store         Tuint8_t
	Fsequence              Tuint16_t
	Flength                Tuint32_t
	Fvisual                Txcb_visualid_t
	F_class                Tuint16_t
	Fbit_gravity           Tuint8_t
	Fwin_gravity           Tuint8_t
	Fbacking_planes        Tuint32_t
	Fbacking_pixel         Tuint32_t
	Fsave_under            Tuint8_t
	Fmap_is_installed      Tuint8_t
	Fmap_state             Tuint8_t
	Foverride_redirect     Tuint8_t
	Fcolormap              Txcb_colormap_t
	Fall_event_masks       Tuint32_t
	Fyour_event_mask       Tuint32_t
	Fdo_not_propagate_mask Tuint16_t
	Fpad0                  [2]Tuint8_t
}

type Txcb_destroy_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_destroy_subwindows_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_set_mode_t = int32

type _xcb_set_mode_t = int32

const _XCB_SET_MODE_INSERT = 0
const _XCB_SET_MODE_DELETE = 1

type Txcb_change_save_set_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_reparent_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fparent       Txcb_window_t
	Fx            Tint16_t
	Fy            Tint16_t
}

type Txcb_map_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_map_subwindows_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_unmap_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_unmap_subwindows_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_config_window_t = int32

type _xcb_config_window_t = int32

const _XCB_CONFIG_WINDOW_X = 1
const _XCB_CONFIG_WINDOW_Y = 2
const _XCB_CONFIG_WINDOW_WIDTH = 4
const _XCB_CONFIG_WINDOW_HEIGHT = 8
const _XCB_CONFIG_WINDOW_BORDER_WIDTH = 16
const _XCB_CONFIG_WINDOW_SIBLING = 32
const _XCB_CONFIG_WINDOW_STACK_MODE = 64

type Txcb_stack_mode_t = int32

type _xcb_stack_mode_t = int32

const _XCB_STACK_MODE_ABOVE = 0
const _XCB_STACK_MODE_BELOW = 1
const _XCB_STACK_MODE_TOP_IF = 2
const _XCB_STACK_MODE_BOTTOM_IF = 3
const _XCB_STACK_MODE_OPPOSITE = 4

type Txcb_configure_window_value_list_t = struct {
	Fx            Tint32_t
	Fy            Tint32_t
	Fwidth        Tuint32_t
	Fheight       Tuint32_t
	Fborder_width Tuint32_t
	Fsibling      Txcb_window_t
	Fstack_mode   Tuint32_t
}

type Txcb_configure_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fvalue_mask   Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_circulate_t = int32

type _xcb_circulate_t = int32

const _XCB_CIRCULATE_RAISE_LOWEST = 0
const _XCB_CIRCULATE_LOWER_HIGHEST = 1

type Txcb_circulate_window_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fdirection    Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_get_geometry_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_geometry_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
}

type Txcb_get_geometry_reply_t = struct {
	Fresponse_type Tuint8_t
	Fdepth         Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Froot          Txcb_window_t
	Fx             Tint16_t
	Fy             Tint16_t
	Fwidth         Tuint16_t
	Fheight        Tuint16_t
	Fborder_width  Tuint16_t
	Fpad0          [2]Tuint8_t
}

type Txcb_query_tree_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_tree_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_query_tree_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Froot          Txcb_window_t
	Fparent        Txcb_window_t
	Fchildren_len  Tuint16_t
	Fpad1          [14]Tuint8_t
}

type Txcb_intern_atom_cookie_t = struct {
	Fsequence uint32
}

type Txcb_intern_atom_request_t = struct {
	Fmajor_opcode   Tuint8_t
	Fonly_if_exists Tuint8_t
	Flength         Tuint16_t
	Fname_len       Tuint16_t
	Fpad0           [2]Tuint8_t
}

type Txcb_intern_atom_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fatom          Txcb_atom_t
}

type Txcb_get_atom_name_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_atom_name_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fatom         Txcb_atom_t
}

type Txcb_get_atom_name_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fname_len      Tuint16_t
	Fpad1          [22]Tuint8_t
}

type Txcb_prop_mode_t = int32

type _xcb_prop_mode_t = int32

const _XCB_PROP_MODE_REPLACE = 0
const _XCB_PROP_MODE_PREPEND = 1
const _XCB_PROP_MODE_APPEND = 2

type Txcb_change_property_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fproperty     Txcb_atom_t
	Ftype1        Txcb_atom_t
	Fformat       Tuint8_t
	Fpad0         [3]Tuint8_t
	Fdata_len     Tuint32_t
}

type Txcb_delete_property_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fproperty     Txcb_atom_t
}

type Txcb_get_property_type_t = int32

type _xcb_get_property_type_t = int32

const _XCB_GET_PROPERTY_TYPE_ANY = 0

type Txcb_get_property_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_property_request_t = struct {
	Fmajor_opcode Tuint8_t
	F_delete      Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fproperty     Txcb_atom_t
	Ftype1        Txcb_atom_t
	Flong_offset  Tuint32_t
	Flong_length  Tuint32_t
}

type Txcb_get_property_reply_t = struct {
	Fresponse_type Tuint8_t
	Fformat        Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Ftype1         Txcb_atom_t
	Fbytes_after   Tuint32_t
	Fvalue_len     Tuint32_t
	Fpad0          [12]Tuint8_t
}

type Txcb_list_properties_cookie_t = struct {
	Fsequence uint32
}

type Txcb_list_properties_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_list_properties_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fatoms_len     Tuint16_t
	Fpad1          [22]Tuint8_t
}

type Txcb_set_selection_owner_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fowner        Txcb_window_t
	Fselection    Txcb_atom_t
	Ftime         Txcb_timestamp_t
}

type Txcb_get_selection_owner_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_selection_owner_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fselection    Txcb_atom_t
}

type Txcb_get_selection_owner_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fowner         Txcb_window_t
}

type Txcb_convert_selection_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Frequestor    Txcb_window_t
	Fselection    Txcb_atom_t
	Ftarget       Txcb_atom_t
	Fproperty     Txcb_atom_t
	Ftime         Txcb_timestamp_t
}

type Txcb_send_event_dest_t = int32

type _xcb_send_event_dest_t = int32

const _XCB_SEND_EVENT_DEST_POINTER_WINDOW = 0
const _XCB_SEND_EVENT_DEST_ITEM_FOCUS = 1

type Txcb_send_event_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpropagate    Tuint8_t
	Flength       Tuint16_t
	Fdestination  Txcb_window_t
	Fevent_mask   Tuint32_t
	Fevent        [32]int8
}

type Txcb_grab_mode_t = int32

type _xcb_grab_mode_t = int32

const _XCB_GRAB_MODE_SYNC = 0
const _XCB_GRAB_MODE_ASYNC = 1

type Txcb_grab_status_t = int32

type _xcb_grab_status_t = int32

const _XCB_GRAB_STATUS_SUCCESS = 0
const _XCB_GRAB_STATUS_ALREADY_GRABBED = 1
const _XCB_GRAB_STATUS_INVALID_TIME = 2
const _XCB_GRAB_STATUS_NOT_VIEWABLE = 3
const _XCB_GRAB_STATUS_FROZEN = 4

type Txcb_cursor_enum_t = int32

type _xcb_cursor_enum_t = int32

const _XCB_CURSOR_NONE = 0

type Txcb_grab_pointer_cookie_t = struct {
	Fsequence uint32
}

type Txcb_grab_pointer_request_t = struct {
	Fmajor_opcode  Tuint8_t
	Fowner_events  Tuint8_t
	Flength        Tuint16_t
	Fgrab_window   Txcb_window_t
	Fevent_mask    Tuint16_t
	Fpointer_mode  Tuint8_t
	Fkeyboard_mode Tuint8_t
	Fconfine_to    Txcb_window_t
	Fcursor        Txcb_cursor_t
	Ftime          Txcb_timestamp_t
}

type Txcb_grab_pointer_reply_t = struct {
	Fresponse_type Tuint8_t
	Fstatus        Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
}

type Txcb_ungrab_pointer_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Ftime         Txcb_timestamp_t
}

type Txcb_button_index_t = int32

type _xcb_button_index_t = int32

const _XCB_BUTTON_INDEX_ANY = 0
const _XCB_BUTTON_INDEX_1 = 1
const _XCB_BUTTON_INDEX_2 = 2
const _XCB_BUTTON_INDEX_3 = 3
const _XCB_BUTTON_INDEX_4 = 4
const _XCB_BUTTON_INDEX_5 = 5

type Txcb_grab_button_request_t = struct {
	Fmajor_opcode  Tuint8_t
	Fowner_events  Tuint8_t
	Flength        Tuint16_t
	Fgrab_window   Txcb_window_t
	Fevent_mask    Tuint16_t
	Fpointer_mode  Tuint8_t
	Fkeyboard_mode Tuint8_t
	Fconfine_to    Txcb_window_t
	Fcursor        Txcb_cursor_t
	Fbutton        Tuint8_t
	Fpad0          Tuint8_t
	Fmodifiers     Tuint16_t
}

type Txcb_ungrab_button_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fbutton       Tuint8_t
	Flength       Tuint16_t
	Fgrab_window  Txcb_window_t
	Fmodifiers    Tuint16_t
	Fpad0         [2]Tuint8_t
}

type Txcb_change_active_pointer_grab_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcursor       Txcb_cursor_t
	Ftime         Txcb_timestamp_t
	Fevent_mask   Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_grab_keyboard_cookie_t = struct {
	Fsequence uint32
}

type Txcb_grab_keyboard_request_t = struct {
	Fmajor_opcode  Tuint8_t
	Fowner_events  Tuint8_t
	Flength        Tuint16_t
	Fgrab_window   Txcb_window_t
	Ftime          Txcb_timestamp_t
	Fpointer_mode  Tuint8_t
	Fkeyboard_mode Tuint8_t
	Fpad0          [2]Tuint8_t
}

type Txcb_grab_keyboard_reply_t = struct {
	Fresponse_type Tuint8_t
	Fstatus        Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
}

type Txcb_ungrab_keyboard_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Ftime         Txcb_timestamp_t
}

type Txcb_grab_t = int32

type _xcb_grab_t = int32

const _XCB_GRAB_ANY = 0

type Txcb_grab_key_request_t = struct {
	Fmajor_opcode  Tuint8_t
	Fowner_events  Tuint8_t
	Flength        Tuint16_t
	Fgrab_window   Txcb_window_t
	Fmodifiers     Tuint16_t
	Fkey           Txcb_keycode_t
	Fpointer_mode  Tuint8_t
	Fkeyboard_mode Tuint8_t
	Fpad0          [3]Tuint8_t
}

type Txcb_ungrab_key_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fkey          Txcb_keycode_t
	Flength       Tuint16_t
	Fgrab_window  Txcb_window_t
	Fmodifiers    Tuint16_t
	Fpad0         [2]Tuint8_t
}

type Txcb_allow_t = int32

type _xcb_allow_t = int32

const _XCB_ALLOW_ASYNC_POINTER = 0
const _XCB_ALLOW_SYNC_POINTER = 1
const _XCB_ALLOW_REPLAY_POINTER = 2
const _XCB_ALLOW_ASYNC_KEYBOARD = 3
const _XCB_ALLOW_SYNC_KEYBOARD = 4
const _XCB_ALLOW_REPLAY_KEYBOARD = 5
const _XCB_ALLOW_ASYNC_BOTH = 6
const _XCB_ALLOW_SYNC_BOTH = 7

type Txcb_allow_events_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
	Ftime         Txcb_timestamp_t
}

type Txcb_grab_server_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_ungrab_server_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_query_pointer_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_pointer_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_query_pointer_reply_t = struct {
	Fresponse_type Tuint8_t
	Fsame_screen   Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Froot          Txcb_window_t
	Fchild         Txcb_window_t
	Froot_x        Tint16_t
	Froot_y        Tint16_t
	Fwin_x         Tint16_t
	Fwin_y         Tint16_t
	Fmask          Tuint16_t
	Fpad0          [2]Tuint8_t
}

type Txcb_timecoord_t = struct {
	Ftime Txcb_timestamp_t
	Fx    Tint16_t
	Fy    Tint16_t
}

type Txcb_timecoord_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_get_motion_events_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_motion_events_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fstart        Txcb_timestamp_t
	Fstop         Txcb_timestamp_t
}

type Txcb_get_motion_events_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fevents_len    Tuint32_t
	Fpad1          [20]Tuint8_t
}

type Txcb_translate_coordinates_cookie_t = struct {
	Fsequence uint32
}

type Txcb_translate_coordinates_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fsrc_window   Txcb_window_t
	Fdst_window   Txcb_window_t
	Fsrc_x        Tint16_t
	Fsrc_y        Tint16_t
}

type Txcb_translate_coordinates_reply_t = struct {
	Fresponse_type Tuint8_t
	Fsame_screen   Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fchild         Txcb_window_t
	Fdst_x         Tint16_t
	Fdst_y         Tint16_t
}

type Txcb_warp_pointer_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fsrc_window   Txcb_window_t
	Fdst_window   Txcb_window_t
	Fsrc_x        Tint16_t
	Fsrc_y        Tint16_t
	Fsrc_width    Tuint16_t
	Fsrc_height   Tuint16_t
	Fdst_x        Tint16_t
	Fdst_y        Tint16_t
}

type Txcb_input_focus_t = int32

type _xcb_input_focus_t = int32

const _XCB_INPUT_FOCUS_NONE = 0
const _XCB_INPUT_FOCUS_POINTER_ROOT = 1
const _XCB_INPUT_FOCUS_PARENT = 2
const _XCB_INPUT_FOCUS_FOLLOW_KEYBOARD = 3

type Txcb_set_input_focus_request_t = struct {
	Fmajor_opcode Tuint8_t
	Frevert_to    Tuint8_t
	Flength       Tuint16_t
	Ffocus        Txcb_window_t
	Ftime         Txcb_timestamp_t
}

type Txcb_get_input_focus_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_input_focus_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_input_focus_reply_t = struct {
	Fresponse_type Tuint8_t
	Frevert_to     Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Ffocus         Txcb_window_t
}

type Txcb_query_keymap_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_keymap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_query_keymap_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fkeys          [32]Tuint8_t
}

type Txcb_open_font_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Ffid          Txcb_font_t
	Fname_len     Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_close_font_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Ffont         Txcb_font_t
}

type Txcb_font_draw_t = int32

type _xcb_font_draw_t = int32

const _XCB_FONT_DRAW_LEFT_TO_RIGHT = 0
const _XCB_FONT_DRAW_RIGHT_TO_LEFT = 1

type Txcb_fontprop_t = struct {
	Fname  Txcb_atom_t
	Fvalue Tuint32_t
}

type Txcb_fontprop_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_charinfo_t = struct {
	Fleft_side_bearing  Tint16_t
	Fright_side_bearing Tint16_t
	Fcharacter_width    Tint16_t
	Fascent             Tint16_t
	Fdescent            Tint16_t
	Fattributes         Tuint16_t
}

type Txcb_charinfo_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_query_font_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_font_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Ffont         Txcb_fontable_t
}

type Txcb_query_font_reply_t = struct {
	Fresponse_type     Tuint8_t
	Fpad0              Tuint8_t
	Fsequence          Tuint16_t
	Flength            Tuint32_t
	Fmin_bounds        Txcb_charinfo_t
	Fpad1              [4]Tuint8_t
	Fmax_bounds        Txcb_charinfo_t
	Fpad2              [4]Tuint8_t
	Fmin_char_or_byte2 Tuint16_t
	Fmax_char_or_byte2 Tuint16_t
	Fdefault_char      Tuint16_t
	Fproperties_len    Tuint16_t
	Fdraw_direction    Tuint8_t
	Fmin_byte1         Tuint8_t
	Fmax_byte1         Tuint8_t
	Fall_chars_exist   Tuint8_t
	Ffont_ascent       Tint16_t
	Ffont_descent      Tint16_t
	Fchar_infos_len    Tuint32_t
}

type Txcb_query_text_extents_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_text_extents_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fodd_length   Tuint8_t
	Flength       Tuint16_t
	Ffont         Txcb_fontable_t
}

type Txcb_query_text_extents_reply_t = struct {
	Fresponse_type   Tuint8_t
	Fdraw_direction  Tuint8_t
	Fsequence        Tuint16_t
	Flength          Tuint32_t
	Ffont_ascent     Tint16_t
	Ffont_descent    Tint16_t
	Foverall_ascent  Tint16_t
	Foverall_descent Tint16_t
	Foverall_width   Tint32_t
	Foverall_left    Tint32_t
	Foverall_right   Tint32_t
}

type Txcb_str_t = struct {
	Fname_len Tuint8_t
}

type Txcb_str_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_list_fonts_cookie_t = struct {
	Fsequence uint32
}

type Txcb_list_fonts_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fmax_names    Tuint16_t
	Fpattern_len  Tuint16_t
}

type Txcb_list_fonts_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fnames_len     Tuint16_t
	Fpad1          [22]Tuint8_t
}

type Txcb_list_fonts_with_info_cookie_t = struct {
	Fsequence uint32
}

type Txcb_list_fonts_with_info_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fmax_names    Tuint16_t
	Fpattern_len  Tuint16_t
}

type Txcb_list_fonts_with_info_reply_t = struct {
	Fresponse_type     Tuint8_t
	Fname_len          Tuint8_t
	Fsequence          Tuint16_t
	Flength            Tuint32_t
	Fmin_bounds        Txcb_charinfo_t
	Fpad0              [4]Tuint8_t
	Fmax_bounds        Txcb_charinfo_t
	Fpad1              [4]Tuint8_t
	Fmin_char_or_byte2 Tuint16_t
	Fmax_char_or_byte2 Tuint16_t
	Fdefault_char      Tuint16_t
	Fproperties_len    Tuint16_t
	Fdraw_direction    Tuint8_t
	Fmin_byte1         Tuint8_t
	Fmax_byte1         Tuint8_t
	Fall_chars_exist   Tuint8_t
	Ffont_ascent       Tint16_t
	Ffont_descent      Tint16_t
	Freplies_hint      Tuint32_t
}

type Txcb_set_font_path_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Ffont_qty     Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_get_font_path_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_font_path_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_font_path_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpath_len      Tuint16_t
	Fpad1          [22]Tuint8_t
}

type Txcb_create_pixmap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fdepth        Tuint8_t
	Flength       Tuint16_t
	Fpid          Txcb_pixmap_t
	Fdrawable     Txcb_drawable_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
}

type Txcb_free_pixmap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fpixmap       Txcb_pixmap_t
}

type Txcb_gc_t = int32

type _xcb_gc_t = int32

const _XCB_GC_FUNCTION = 1
const _XCB_GC_PLANE_MASK = 2
const _XCB_GC_FOREGROUND = 4
const _XCB_GC_BACKGROUND = 8
const _XCB_GC_LINE_WIDTH = 16
const _XCB_GC_LINE_STYLE = 32
const _XCB_GC_CAP_STYLE = 64
const _XCB_GC_JOIN_STYLE = 128
const _XCB_GC_FILL_STYLE = 256
const _XCB_GC_FILL_RULE = 512
const _XCB_GC_TILE = 1024
const _XCB_GC_STIPPLE = 2048
const _XCB_GC_TILE_STIPPLE_ORIGIN_X = 4096
const _XCB_GC_TILE_STIPPLE_ORIGIN_Y = 8192
const _XCB_GC_FONT = 16384
const _XCB_GC_SUBWINDOW_MODE = 32768
const _XCB_GC_GRAPHICS_EXPOSURES = 65536
const _XCB_GC_CLIP_ORIGIN_X = 131072
const _XCB_GC_CLIP_ORIGIN_Y = 262144
const _XCB_GC_CLIP_MASK = 524288
const _XCB_GC_DASH_OFFSET = 1048576
const _XCB_GC_DASH_LIST = 2097152
const _XCB_GC_ARC_MODE = 4194304

type Txcb_gx_t = int32

type _xcb_gx_t = int32

const _XCB_GX_CLEAR = 0
const _XCB_GX_AND = 1
const _XCB_GX_AND_REVERSE = 2
const _XCB_GX_COPY = 3
const _XCB_GX_AND_INVERTED = 4
const _XCB_GX_NOOP = 5
const _XCB_GX_XOR = 6
const _XCB_GX_OR = 7
const _XCB_GX_NOR = 8
const _XCB_GX_EQUIV = 9
const _XCB_GX_INVERT = 10
const _XCB_GX_OR_REVERSE = 11
const _XCB_GX_COPY_INVERTED = 12
const _XCB_GX_OR_INVERTED = 13
const _XCB_GX_NAND = 14
const _XCB_GX_SET = 15

type Txcb_line_style_t = int32

type _xcb_line_style_t = int32

const _XCB_LINE_STYLE_SOLID = 0
const _XCB_LINE_STYLE_ON_OFF_DASH = 1
const _XCB_LINE_STYLE_DOUBLE_DASH = 2

type Txcb_cap_style_t = int32

type _xcb_cap_style_t = int32

const _XCB_CAP_STYLE_NOT_LAST = 0
const _XCB_CAP_STYLE_BUTT = 1
const _XCB_CAP_STYLE_ROUND = 2
const _XCB_CAP_STYLE_PROJECTING = 3

type Txcb_join_style_t = int32

type _xcb_join_style_t = int32

const _XCB_JOIN_STYLE_MITER = 0
const _XCB_JOIN_STYLE_ROUND = 1
const _XCB_JOIN_STYLE_BEVEL = 2

type Txcb_fill_style_t = int32

type _xcb_fill_style_t = int32

const _XCB_FILL_STYLE_SOLID = 0
const _XCB_FILL_STYLE_TILED = 1
const _XCB_FILL_STYLE_STIPPLED = 2
const _XCB_FILL_STYLE_OPAQUE_STIPPLED = 3

type Txcb_fill_rule_t = int32

type _xcb_fill_rule_t = int32

const _XCB_FILL_RULE_EVEN_ODD = 0
const _XCB_FILL_RULE_WINDING = 1

type Txcb_subwindow_mode_t = int32

type _xcb_subwindow_mode_t = int32

const _XCB_SUBWINDOW_MODE_CLIP_BY_CHILDREN = 0
const _XCB_SUBWINDOW_MODE_INCLUDE_INFERIORS = 1

type Txcb_arc_mode_t = int32

type _xcb_arc_mode_t = int32

const _XCB_ARC_MODE_CHORD = 0
const _XCB_ARC_MODE_PIE_SLICE = 1

type Txcb_create_gc_value_list_t = struct {
	Ffunction              Tuint32_t
	Fplane_mask            Tuint32_t
	Fforeground            Tuint32_t
	Fbackground            Tuint32_t
	Fline_width            Tuint32_t
	Fline_style            Tuint32_t
	Fcap_style             Tuint32_t
	Fjoin_style            Tuint32_t
	Ffill_style            Tuint32_t
	Ffill_rule             Tuint32_t
	Ftile                  Txcb_pixmap_t
	Fstipple               Txcb_pixmap_t
	Ftile_stipple_x_origin Tint32_t
	Ftile_stipple_y_origin Tint32_t
	Ffont                  Txcb_font_t
	Fsubwindow_mode        Tuint32_t
	Fgraphics_exposures    Txcb_bool32_t
	Fclip_x_origin         Tint32_t
	Fclip_y_origin         Tint32_t
	Fclip_mask             Txcb_pixmap_t
	Fdash_offset           Tuint32_t
	Fdashes                Tuint32_t
	Farc_mode              Tuint32_t
}

type Txcb_create_gc_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcid          Txcb_gcontext_t
	Fdrawable     Txcb_drawable_t
	Fvalue_mask   Tuint32_t
}

type Txcb_change_gc_value_list_t = struct {
	Ffunction              Tuint32_t
	Fplane_mask            Tuint32_t
	Fforeground            Tuint32_t
	Fbackground            Tuint32_t
	Fline_width            Tuint32_t
	Fline_style            Tuint32_t
	Fcap_style             Tuint32_t
	Fjoin_style            Tuint32_t
	Ffill_style            Tuint32_t
	Ffill_rule             Tuint32_t
	Ftile                  Txcb_pixmap_t
	Fstipple               Txcb_pixmap_t
	Ftile_stipple_x_origin Tint32_t
	Ftile_stipple_y_origin Tint32_t
	Ffont                  Txcb_font_t
	Fsubwindow_mode        Tuint32_t
	Fgraphics_exposures    Txcb_bool32_t
	Fclip_x_origin         Tint32_t
	Fclip_y_origin         Tint32_t
	Fclip_mask             Txcb_pixmap_t
	Fdash_offset           Tuint32_t
	Fdashes                Tuint32_t
	Farc_mode              Tuint32_t
}

type Txcb_change_gc_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fgc           Txcb_gcontext_t
	Fvalue_mask   Tuint32_t
}

type Txcb_copy_gc_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fsrc_gc       Txcb_gcontext_t
	Fdst_gc       Txcb_gcontext_t
	Fvalue_mask   Tuint32_t
}

type Txcb_set_dashes_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fgc           Txcb_gcontext_t
	Fdash_offset  Tuint16_t
	Fdashes_len   Tuint16_t
}

type Txcb_clip_ordering_t = int32

type _xcb_clip_ordering_t = int32

const _XCB_CLIP_ORDERING_UNSORTED = 0
const _XCB_CLIP_ORDERING_Y_SORTED = 1
const _XCB_CLIP_ORDERING_YX_SORTED = 2
const _XCB_CLIP_ORDERING_YX_BANDED = 3

type Txcb_set_clip_rectangles_request_t = struct {
	Fmajor_opcode  Tuint8_t
	Fordering      Tuint8_t
	Flength        Tuint16_t
	Fgc            Txcb_gcontext_t
	Fclip_x_origin Tint16_t
	Fclip_y_origin Tint16_t
}

type Txcb_free_gc_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fgc           Txcb_gcontext_t
}

type Txcb_clear_area_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fexposures    Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fx            Tint16_t
	Fy            Tint16_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
}

type Txcb_copy_area_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fsrc_drawable Txcb_drawable_t
	Fdst_drawable Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fsrc_x        Tint16_t
	Fsrc_y        Tint16_t
	Fdst_x        Tint16_t
	Fdst_y        Tint16_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
}

type Txcb_copy_plane_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fsrc_drawable Txcb_drawable_t
	Fdst_drawable Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fsrc_x        Tint16_t
	Fsrc_y        Tint16_t
	Fdst_x        Tint16_t
	Fdst_y        Tint16_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
	Fbit_plane    Tuint32_t
}

type Txcb_coord_mode_t = int32

type _xcb_coord_mode_t = int32

const _XCB_COORD_MODE_ORIGIN = 0
const _XCB_COORD_MODE_PREVIOUS = 1

type Txcb_poly_point_request_t = struct {
	Fmajor_opcode    Tuint8_t
	Fcoordinate_mode Tuint8_t
	Flength          Tuint16_t
	Fdrawable        Txcb_drawable_t
	Fgc              Txcb_gcontext_t
}

type Txcb_poly_line_request_t = struct {
	Fmajor_opcode    Tuint8_t
	Fcoordinate_mode Tuint8_t
	Flength          Tuint16_t
	Fdrawable        Txcb_drawable_t
	Fgc              Txcb_gcontext_t
}

type Txcb_segment_t = struct {
	Fx1 Tint16_t
	Fy1 Tint16_t
	Fx2 Tint16_t
	Fy2 Tint16_t
}

type Txcb_segment_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_poly_segment_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
}

type Txcb_poly_rectangle_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
}

type Txcb_poly_arc_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
}

type Txcb_poly_shape_t = int32

type _xcb_poly_shape_t = int32

const _XCB_POLY_SHAPE_COMPLEX = 0
const _XCB_POLY_SHAPE_NONCONVEX = 1
const _XCB_POLY_SHAPE_CONVEX = 2

type Txcb_fill_poly_request_t = struct {
	Fmajor_opcode    Tuint8_t
	Fpad0            Tuint8_t
	Flength          Tuint16_t
	Fdrawable        Txcb_drawable_t
	Fgc              Txcb_gcontext_t
	Fshape           Tuint8_t
	Fcoordinate_mode Tuint8_t
	Fpad1            [2]Tuint8_t
}

type Txcb_poly_fill_rectangle_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
}

type Txcb_poly_fill_arc_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
}

type Txcb_image_format_t = int32

type _xcb_image_format_t = int32

const _XCB_IMAGE_FORMAT_XY_BITMAP = 0
const _XCB_IMAGE_FORMAT_XY_PIXMAP = 1
const _XCB_IMAGE_FORMAT_Z_PIXMAP = 2

type Txcb_put_image_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fformat       Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
	Fdst_x        Tint16_t
	Fdst_y        Tint16_t
	Fleft_pad     Tuint8_t
	Fdepth        Tuint8_t
	Fpad0         [2]Tuint8_t
}

type Txcb_get_image_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_image_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fformat       Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fx            Tint16_t
	Fy            Tint16_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
	Fplane_mask   Tuint32_t
}

type Txcb_get_image_reply_t = struct {
	Fresponse_type Tuint8_t
	Fdepth         Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fvisual        Txcb_visualid_t
	Fpad0          [20]Tuint8_t
}

type Txcb_poly_text_8_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fx            Tint16_t
	Fy            Tint16_t
}

type Txcb_poly_text_16_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fx            Tint16_t
	Fy            Tint16_t
}

type Txcb_image_text_8_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fstring_len   Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fx            Tint16_t
	Fy            Tint16_t
}

type Txcb_image_text_16_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fstring_len   Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fgc           Txcb_gcontext_t
	Fx            Tint16_t
	Fy            Tint16_t
}

type Txcb_colormap_alloc_t = int32

type _xcb_colormap_alloc_t = int32

const _XCB_COLORMAP_ALLOC_NONE = 0
const _XCB_COLORMAP_ALLOC_ALL = 1

type Txcb_create_colormap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Falloc        Tuint8_t
	Flength       Tuint16_t
	Fmid          Txcb_colormap_t
	Fwindow       Txcb_window_t
	Fvisual       Txcb_visualid_t
}

type Txcb_free_colormap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
}

type Txcb_copy_colormap_and_free_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fmid          Txcb_colormap_t
	Fsrc_cmap     Txcb_colormap_t
}

type Txcb_install_colormap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
}

type Txcb_uninstall_colormap_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
}

type Txcb_list_installed_colormaps_cookie_t = struct {
	Fsequence uint32
}

type Txcb_list_installed_colormaps_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
}

type Txcb_list_installed_colormaps_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fcmaps_len     Tuint16_t
	Fpad1          [22]Tuint8_t
}

type Txcb_alloc_color_cookie_t = struct {
	Fsequence uint32
}

type Txcb_alloc_color_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fred          Tuint16_t
	Fgreen        Tuint16_t
	Fblue         Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_alloc_color_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fred           Tuint16_t
	Fgreen         Tuint16_t
	Fblue          Tuint16_t
	Fpad1          [2]Tuint8_t
	Fpixel         Tuint32_t
}

type Txcb_alloc_named_color_cookie_t = struct {
	Fsequence uint32
}

type Txcb_alloc_named_color_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fname_len     Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_alloc_named_color_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpixel         Tuint32_t
	Fexact_red     Tuint16_t
	Fexact_green   Tuint16_t
	Fexact_blue    Tuint16_t
	Fvisual_red    Tuint16_t
	Fvisual_green  Tuint16_t
	Fvisual_blue   Tuint16_t
}

type Txcb_alloc_color_cells_cookie_t = struct {
	Fsequence uint32
}

type Txcb_alloc_color_cells_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fcontiguous   Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fcolors       Tuint16_t
	Fplanes       Tuint16_t
}

type Txcb_alloc_color_cells_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpixels_len    Tuint16_t
	Fmasks_len     Tuint16_t
	Fpad1          [20]Tuint8_t
}

type Txcb_alloc_color_planes_cookie_t = struct {
	Fsequence uint32
}

type Txcb_alloc_color_planes_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fcontiguous   Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fcolors       Tuint16_t
	Freds         Tuint16_t
	Fgreens       Tuint16_t
	Fblues        Tuint16_t
}

type Txcb_alloc_color_planes_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpixels_len    Tuint16_t
	Fpad1          [2]Tuint8_t
	Fred_mask      Tuint32_t
	Fgreen_mask    Tuint32_t
	Fblue_mask     Tuint32_t
	Fpad2          [8]Tuint8_t
}

type Txcb_free_colors_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fplane_mask   Tuint32_t
}

type Txcb_color_flag_t = int32

type _xcb_color_flag_t = int32

const _XCB_COLOR_FLAG_RED = 1
const _XCB_COLOR_FLAG_GREEN = 2
const _XCB_COLOR_FLAG_BLUE = 4

type Txcb_coloritem_t = struct {
	Fpixel Tuint32_t
	Fred   Tuint16_t
	Fgreen Tuint16_t
	Fblue  Tuint16_t
	Fflags Tuint8_t
	Fpad0  Tuint8_t
}

type Txcb_coloritem_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_store_colors_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
}

type Txcb_store_named_color_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fflags        Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fpixel        Tuint32_t
	Fname_len     Tuint16_t
	Fpad0         [2]Tuint8_t
}

type Txcb_rgb_t = struct {
	Fred   Tuint16_t
	Fgreen Tuint16_t
	Fblue  Tuint16_t
	Fpad0  [2]Tuint8_t
}

type Txcb_rgb_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_query_colors_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_colors_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
}

type Txcb_query_colors_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fcolors_len    Tuint16_t
	Fpad1          [22]Tuint8_t
}

type Txcb_lookup_color_cookie_t = struct {
	Fsequence uint32
}

type Txcb_lookup_color_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcmap         Txcb_colormap_t
	Fname_len     Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_lookup_color_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fexact_red     Tuint16_t
	Fexact_green   Tuint16_t
	Fexact_blue    Tuint16_t
	Fvisual_red    Tuint16_t
	Fvisual_green  Tuint16_t
	Fvisual_blue   Tuint16_t
}

type Txcb_pixmap_enum_t = int32

type _xcb_pixmap_enum_t = int32

const _XCB_PIXMAP_NONE = 0

type Txcb_create_cursor_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcid          Txcb_cursor_t
	Fsource       Txcb_pixmap_t
	Fmask         Txcb_pixmap_t
	Ffore_red     Tuint16_t
	Ffore_green   Tuint16_t
	Ffore_blue    Tuint16_t
	Fback_red     Tuint16_t
	Fback_green   Tuint16_t
	Fback_blue    Tuint16_t
	Fx            Tuint16_t
	Fy            Tuint16_t
}

type Txcb_font_enum_t = int32

type _xcb_font_enum_t = int32

const _XCB_FONT_NONE = 0

type Txcb_create_glyph_cursor_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcid          Txcb_cursor_t
	Fsource_font  Txcb_font_t
	Fmask_font    Txcb_font_t
	Fsource_char  Tuint16_t
	Fmask_char    Tuint16_t
	Ffore_red     Tuint16_t
	Ffore_green   Tuint16_t
	Ffore_blue    Tuint16_t
	Fback_red     Tuint16_t
	Fback_green   Tuint16_t
	Fback_blue    Tuint16_t
}

type Txcb_free_cursor_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcursor       Txcb_cursor_t
}

type Txcb_recolor_cursor_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fcursor       Txcb_cursor_t
	Ffore_red     Tuint16_t
	Ffore_green   Tuint16_t
	Ffore_blue    Tuint16_t
	Fback_red     Tuint16_t
	Fback_green   Tuint16_t
	Fback_blue    Tuint16_t
}

type Txcb_query_shape_of_t = int32

type _xcb_query_shape_of_t = int32

const _XCB_QUERY_SHAPE_OF_LARGEST_CURSOR = 0
const _XCB_QUERY_SHAPE_OF_FASTEST_TILE = 1
const _XCB_QUERY_SHAPE_OF_FASTEST_STIPPLE = 2

type Txcb_query_best_size_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_best_size_request_t = struct {
	Fmajor_opcode Tuint8_t
	F_class       Tuint8_t
	Flength       Tuint16_t
	Fdrawable     Txcb_drawable_t
	Fwidth        Tuint16_t
	Fheight       Tuint16_t
}

type Txcb_query_best_size_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fwidth         Tuint16_t
	Fheight        Tuint16_t
}

type Txcb_query_extension_cookie_t = struct {
	Fsequence uint32
}

type Txcb_query_extension_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fname_len     Tuint16_t
	Fpad1         [2]Tuint8_t
}

type Txcb_query_extension_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpresent       Tuint8_t
	Fmajor_opcode  Tuint8_t
	Ffirst_event   Tuint8_t
	Ffirst_error   Tuint8_t
}

type Txcb_list_extensions_cookie_t = struct {
	Fsequence uint32
}

type Txcb_list_extensions_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_list_extensions_reply_t = struct {
	Fresponse_type Tuint8_t
	Fnames_len     Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpad0          [24]Tuint8_t
}

type Txcb_change_keyboard_mapping_request_t = struct {
	Fmajor_opcode        Tuint8_t
	Fkeycode_count       Tuint8_t
	Flength              Tuint16_t
	Ffirst_keycode       Txcb_keycode_t
	Fkeysyms_per_keycode Tuint8_t
	Fpad0                [2]Tuint8_t
}

type Txcb_get_keyboard_mapping_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_keyboard_mapping_request_t = struct {
	Fmajor_opcode  Tuint8_t
	Fpad0          Tuint8_t
	Flength        Tuint16_t
	Ffirst_keycode Txcb_keycode_t
	Fcount         Tuint8_t
}

type Txcb_get_keyboard_mapping_reply_t = struct {
	Fresponse_type       Tuint8_t
	Fkeysyms_per_keycode Tuint8_t
	Fsequence            Tuint16_t
	Flength              Tuint32_t
	Fpad0                [24]Tuint8_t
}

type Txcb_kb_t = int32

type _xcb_kb_t = int32

const _XCB_KB_KEY_CLICK_PERCENT = 1
const _XCB_KB_BELL_PERCENT = 2
const _XCB_KB_BELL_PITCH = 4
const _XCB_KB_BELL_DURATION = 8
const _XCB_KB_LED = 16
const _XCB_KB_LED_MODE = 32
const _XCB_KB_KEY = 64
const _XCB_KB_AUTO_REPEAT_MODE = 128

type Txcb_led_mode_t = int32

type _xcb_led_mode_t = int32

const _XCB_LED_MODE_OFF = 0
const _XCB_LED_MODE_ON = 1

type Txcb_auto_repeat_mode_t = int32

type _xcb_auto_repeat_mode_t = int32

const _XCB_AUTO_REPEAT_MODE_OFF = 0
const _XCB_AUTO_REPEAT_MODE_ON = 1
const _XCB_AUTO_REPEAT_MODE_DEFAULT = 2

type Txcb_change_keyboard_control_value_list_t = struct {
	Fkey_click_percent Tint32_t
	Fbell_percent      Tint32_t
	Fbell_pitch        Tint32_t
	Fbell_duration     Tint32_t
	Fled               Tuint32_t
	Fled_mode          Tuint32_t
	Fkey               Txcb_keycode32_t
	Fauto_repeat_mode  Tuint32_t
}

type Txcb_change_keyboard_control_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fvalue_mask   Tuint32_t
}

type Txcb_get_keyboard_control_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_keyboard_control_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_keyboard_control_reply_t = struct {
	Fresponse_type      Tuint8_t
	Fglobal_auto_repeat Tuint8_t
	Fsequence           Tuint16_t
	Flength             Tuint32_t
	Fled_mask           Tuint32_t
	Fkey_click_percent  Tuint8_t
	Fbell_percent       Tuint8_t
	Fbell_pitch         Tuint16_t
	Fbell_duration      Tuint16_t
	Fpad0               [2]Tuint8_t
	Fauto_repeats       [32]Tuint8_t
}

type Txcb_bell_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpercent      Tint8_t
	Flength       Tuint16_t
}

type Txcb_change_pointer_control_request_t = struct {
	Fmajor_opcode             Tuint8_t
	Fpad0                     Tuint8_t
	Flength                   Tuint16_t
	Facceleration_numerator   Tint16_t
	Facceleration_denominator Tint16_t
	Fthreshold                Tint16_t
	Fdo_acceleration          Tuint8_t
	Fdo_threshold             Tuint8_t
}

type Txcb_get_pointer_control_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_pointer_control_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_pointer_control_reply_t = struct {
	Fresponse_type            Tuint8_t
	Fpad0                     Tuint8_t
	Fsequence                 Tuint16_t
	Flength                   Tuint32_t
	Facceleration_numerator   Tuint16_t
	Facceleration_denominator Tuint16_t
	Fthreshold                Tuint16_t
	Fpad1                     [18]Tuint8_t
}

type Txcb_blanking_t = int32

type _xcb_blanking_t = int32

const _XCB_BLANKING_NOT_PREFERRED = 0
const _XCB_BLANKING_PREFERRED = 1
const _XCB_BLANKING_DEFAULT = 2

type Txcb_exposures_t = int32

type _xcb_exposures_t = int32

const _XCB_EXPOSURES_NOT_ALLOWED = 0
const _XCB_EXPOSURES_ALLOWED = 1
const _XCB_EXPOSURES_DEFAULT = 2

type Txcb_set_screen_saver_request_t = struct {
	Fmajor_opcode    Tuint8_t
	Fpad0            Tuint8_t
	Flength          Tuint16_t
	Ftimeout         Tint16_t
	Finterval        Tint16_t
	Fprefer_blanking Tuint8_t
	Fallow_exposures Tuint8_t
}

type Txcb_get_screen_saver_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_screen_saver_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_screen_saver_reply_t = struct {
	Fresponse_type   Tuint8_t
	Fpad0            Tuint8_t
	Fsequence        Tuint16_t
	Flength          Tuint32_t
	Ftimeout         Tuint16_t
	Finterval        Tuint16_t
	Fprefer_blanking Tuint8_t
	Fallow_exposures Tuint8_t
	Fpad1            [18]Tuint8_t
}

type Txcb_host_mode_t = int32

type _xcb_host_mode_t = int32

const _XCB_HOST_MODE_INSERT = 0
const _XCB_HOST_MODE_DELETE = 1

type Txcb_family_t = int32

type _xcb_family_t = int32

const _XCB_FAMILY_INTERNET = 0
const _XCB_FAMILY_DECNET = 1
const _XCB_FAMILY_CHAOS = 2
const _XCB_FAMILY_SERVER_INTERPRETED = 5
const _XCB_FAMILY_INTERNET_6 = 6

type Txcb_change_hosts_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
	Ffamily       Tuint8_t
	Fpad0         Tuint8_t
	Faddress_len  Tuint16_t
}

type Txcb_host_t = struct {
	Ffamily      Tuint8_t
	Fpad0        Tuint8_t
	Faddress_len Tuint16_t
}

type Txcb_host_iterator_t = struct {
	Fdata  uintptr
	Frem   int32
	Findex int32
}

type Txcb_list_hosts_cookie_t = struct {
	Fsequence uint32
}

type Txcb_list_hosts_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_list_hosts_reply_t = struct {
	Fresponse_type Tuint8_t
	Fmode          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fhosts_len     Tuint16_t
	Fpad0          [22]Tuint8_t
}

type Txcb_access_control_t = int32

type _xcb_access_control_t = int32

const _XCB_ACCESS_CONTROL_DISABLE = 0
const _XCB_ACCESS_CONTROL_ENABLE = 1

type Txcb_set_access_control_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_close_down_t = int32

type _xcb_close_down_t = int32

const _XCB_CLOSE_DOWN_DESTROY_ALL = 0
const _XCB_CLOSE_DOWN_RETAIN_PERMANENT = 1
const _XCB_CLOSE_DOWN_RETAIN_TEMPORARY = 2

type Txcb_set_close_down_mode_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_kill_t = int32

type _xcb_kill_t = int32

const _XCB_KILL_ALL_TEMPORARY = 0

type Txcb_kill_client_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fresource     Tuint32_t
}

type Txcb_rotate_properties_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
	Fwindow       Txcb_window_t
	Fatoms_len    Tuint16_t
	Fdelta        Tint16_t
}

type Txcb_screen_saver_t = int32

type _xcb_screen_saver_t = int32

const _XCB_SCREEN_SAVER_RESET = 0
const _XCB_SCREEN_SAVER_ACTIVE = 1

type Txcb_force_screen_saver_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmode         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_mapping_status_t = int32

type _xcb_mapping_status_t = int32

const _XCB_MAPPING_STATUS_SUCCESS = 0
const _XCB_MAPPING_STATUS_BUSY = 1
const _XCB_MAPPING_STATUS_FAILURE = 2

type Txcb_set_pointer_mapping_cookie_t = struct {
	Fsequence uint32
}

type Txcb_set_pointer_mapping_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fmap_len      Tuint8_t
	Flength       Tuint16_t
}

type Txcb_set_pointer_mapping_reply_t = struct {
	Fresponse_type Tuint8_t
	Fstatus        Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
}

type Txcb_get_pointer_mapping_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_pointer_mapping_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_pointer_mapping_reply_t = struct {
	Fresponse_type Tuint8_t
	Fmap_len       Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fpad0          [24]Tuint8_t
}

type Txcb_map_index_t = int32

type _xcb_map_index_t = int32

const _XCB_MAP_INDEX_SHIFT = 0
const _XCB_MAP_INDEX_LOCK = 1
const _XCB_MAP_INDEX_CONTROL = 2
const _XCB_MAP_INDEX_1 = 3
const _XCB_MAP_INDEX_2 = 4
const _XCB_MAP_INDEX_3 = 5
const _XCB_MAP_INDEX_4 = 6
const _XCB_MAP_INDEX_5 = 7

type Txcb_set_modifier_mapping_cookie_t = struct {
	Fsequence uint32
}

type Txcb_set_modifier_mapping_request_t = struct {
	Fmajor_opcode          Tuint8_t
	Fkeycodes_per_modifier Tuint8_t
	Flength                Tuint16_t
}

type Txcb_set_modifier_mapping_reply_t = struct {
	Fresponse_type Tuint8_t
	Fstatus        Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
}

type Txcb_get_modifier_mapping_cookie_t = struct {
	Fsequence uint32
}

type Txcb_get_modifier_mapping_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_get_modifier_mapping_reply_t = struct {
	Fresponse_type         Tuint8_t
	Fkeycodes_per_modifier Tuint8_t
	Fsequence              Tuint16_t
	Flength                Tuint32_t
	Fpad0                  [24]Tuint8_t
}

type Txcb_no_operation_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fpad0         Tuint8_t
	Flength       Tuint16_t
}

type Txcb_auth_info_t = struct {
	Fnamelen int32
	Fname    uintptr
	Fdatalen int32
	Fdata    uintptr
}

type Txcb_big_requests_enable_cookie_t = struct {
	Fsequence uint32
}

type Txcb_big_requests_enable_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fminor_opcode Tuint8_t
	Flength       Tuint16_t
}

type Txcb_big_requests_enable_reply_t = struct {
	Fresponse_type          Tuint8_t
	Fpad0                   Tuint8_t
	Fsequence               Tuint16_t
	Flength                 Tuint32_t
	Fmaximum_request_length Tuint32_t
}

type _workarounds = int32

const _WORKAROUND_NONE = 0
const _WORKAROUND_GLX_GET_FB_CONFIGS_BUG = 1
const _WORKAROUND_EXTERNAL_SOCKET_OWNER = 2

type _lazy_reply_tag = int32

const _LAZY_NONE = 0
const _LAZY_COOKIE = 1
const _LAZY_FORCED = 2

type Txcb_list_free_func_t = uintptr

type T_xcb_fd = struct {
	Ffd  [16]int32
	Fnfd int32
	Fifd int32
}

type T_xcb_out = struct {
	Fcond                       Tpthread_cond_t
	Fwriting                    int32
	Fsocket_cond                Tpthread_cond_t
	Freturn_socket              uintptr
	Fsocket_closure             uintptr
	Fsocket_moving              int32
	Fqueue                      [16384]int8
	Fqueue_len                  int32
	Frequest                    Tuint64_t
	Frequest_written            Tuint64_t
	Frequest_expected_written   Tuint64_t
	Ftotal_written              Tuint64_t
	Freqlenlock                 Tpthread_mutex_t
	Fmaximum_request_length_tag _lazy_reply_tag
	Fmaximum_request_length     struct {
		Fvalue  [0]Tuint32_t
		Fcookie Txcb_big_requests_enable_cookie_t
	}
	Fout_fd T_xcb_fd
}

type T_xcb_in = struct {
	Fevent_cond           Tpthread_cond_t
	Freading              int32
	Fqueue                [4096]int8
	Fqueue_len            int32
	Frequest_expected     Tuint64_t
	Frequest_read         Tuint64_t
	Frequest_completed    Tuint64_t
	Ftotal_read           Tuint64_t
	Fcurrent_reply        uintptr
	Fcurrent_reply_tail   uintptr
	Freplies              uintptr
	Fevents               uintptr
	Fevents_tail          uintptr
	Freaders              uintptr
	Fspecial_waiters      uintptr
	Fpending_replies      uintptr
	Fpending_replies_tail uintptr
	Fin_fd                T_xcb_fd
	Fspecial_events       uintptr
}

type T_xcb_xid = struct {
	Flock Tpthread_mutex_t
	Flast Tuint32_t
	Fbase Tuint32_t
	Fmax  Tuint32_t
	Finc  Tuint32_t
}

type T_xcb_ext = struct {
	Flock            Tpthread_mutex_t
	Fextensions      uintptr
	Fextensions_size int32
}

type Txcb_connection_t1 = struct {
	Fhas_error int32
	Fsetup     uintptr
	Ffd        int32
	Fiolock    Tpthread_mutex_t
	Fin        T_xcb_in
	Fout       T_xcb_out
	Fext       T_xcb_ext
	Fxid       T_xcb_xid
}

type Tnfds_t = uint32

type Tpollfd = struct {
	Ffd      int32
	Fevents  int16
	Frevents int16
}

type Tsocklen_t = uint32

type Tsa_family_t = uint16

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
	F__ss_padding [122]int8
	F__ss_align   uint32
}

type Timaxdiv_t = struct {
	Fquot Tintmax_t
	Frem  Tintmax_t
}

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

/* SHUT_RDWR is fairly recent and is not available on all platforms */

type Txcb_setup_generic_t = struct {
	Fstatus Tuint8_t
	Fpad0   [5]Tuint8_t
	Flength Tuint16_t
}

var _xcb_error_setup = Txcb_setup_t{}

// C documentation
//
//	/* Keep this list in sync with is_static_error_conn()! */
var _xcb_con_error = int32(m_XCB_CONN_ERROR)
var _xcb_con_closed_mem_er = int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT)
var _xcb_con_closed_parse_er = int32(m_XCB_CONN_CLOSED_PARSE_ERR)
var _xcb_con_closed_screen_er = int32(m_XCB_CONN_CLOSED_INVALID_SCREEN)

func _is_static_error_conn(tls *libc.TLS, c uintptr) (r int32) {
	return libc.BoolInt32(c == uintptr(unsafe.Pointer(&_xcb_con_error)) || c == uintptr(unsafe.Pointer(&_xcb_con_closed_mem_er)) || c == uintptr(unsafe.Pointer(&_xcb_con_closed_parse_er)) || c == uintptr(unsafe.Pointer(&_xcb_con_closed_screen_er)))
}

func _set_fd_flags(tls *libc.TLS, fd int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var flags int32
	_ = flags
	/* Win32 doesn't have file descriptors and the fcntl function. This block sets the socket in non-blocking mode */
	flags = libc.Xfcntl(tls, fd, int32(m_F_GETFL), libc.VaList(bp+8, 0))
	if flags == -int32(1) {
		return 0
	}
	flags |= int32(m_O_NONBLOCK)
	if libc.Xfcntl(tls, fd, int32(m_F_SETFL), libc.VaList(bp+8, flags)) == -int32(1) {
		return 0
	}
	if libc.Xfcntl(tls, fd, int32(m_F_SETFD), libc.VaList(bp+8, int32(m_FD_CLOEXEC))) == -int32(1) {
		return 0
	}
	return int32(1)
}

func _write_setup(tls *libc.TLS, c uintptr, auth_info uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var count, ret, v1, v2, v4, v5, v7, v8 int32
	var v3, v6 Tuint16_t
	var _ /* out at bp+0 */ Txcb_setup_request_t
	var _ /* parts at bp+12 */ [6]Tiovec
	_, _, _, _, _, _, _, _, _, _ = count, ret, v1, v2, v3, v4, v5, v6, v7, v8
	count = 0
	libc.Xmemset(tls, bp, 0, uint32(12))
	/* B = 0x42 = MSB first, l = 0x6c = LSB first */
	if libc.Xhtonl(tls, _endian) == _endian {
		(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fbyte_order = uint8(0x42)
	} else {
		(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fbyte_order = uint8(0x6c)
	}
	(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fprotocol_major_version = uint16(m_X_PROTOCOL)
	(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fprotocol_minor_version = uint16(m_X_PROTOCOL_REVISION)
	(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fauthorization_protocol_name_len = uint16(0)
	(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fauthorization_protocol_data_len = uint16(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[count].Fiov_len = uint32(12)
	v1 = count
	count++
	(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[v1].Fiov_base = bp
	(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[count].Fiov_len = -libc.Uint32FromInt64(12) & libc.Uint32FromInt32(3)
	v2 = count
	count++
	(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[v2].Fiov_base = uintptr(unsafe.Pointer(&_pad))
	if auth_info != 0 {
		v3 = libc.Uint16FromInt32((*Txcb_auth_info_t)(unsafe.Pointer(auth_info)).Fnamelen)
		(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fauthorization_protocol_name_len = v3
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[count].Fiov_len = uint32(v3)
		v4 = count
		count++
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[v4].Fiov_base = (*Txcb_auth_info_t)(unsafe.Pointer(auth_info)).Fname
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[count].Fiov_len = libc.Uint32FromInt32(-libc.Int32FromUint16((*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fauthorization_protocol_name_len) & libc.Int32FromInt32(3))
		v5 = count
		count++
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[v5].Fiov_base = uintptr(unsafe.Pointer(&_pad))
		v6 = libc.Uint16FromInt32((*Txcb_auth_info_t)(unsafe.Pointer(auth_info)).Fdatalen)
		(*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fauthorization_protocol_data_len = v6
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[count].Fiov_len = uint32(v6)
		v7 = count
		count++
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[v7].Fiov_base = (*Txcb_auth_info_t)(unsafe.Pointer(auth_info)).Fdata
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[count].Fiov_len = libc.Uint32FromInt32(-libc.Int32FromUint16((*(*Txcb_setup_request_t)(unsafe.Pointer(bp))).Fauthorization_protocol_data_len) & libc.Int32FromInt32(3))
		v8 = count
		count++
		(*(*[6]Tiovec)(unsafe.Pointer(bp + 12)))[v8].Fiov_base = uintptr(unsafe.Pointer(&_pad))
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	ret = X_xcb_out_send(tls, c, bp+12, count)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

var _pad [3]int8

var _endian = uint32(0x01020304)

func _read_setup(tls *libc.TLS, c uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var setup, setup1, tmp uintptr
	var _ /* newline at bp+0 */ int8
	_, _, _ = setup, setup1, tmp
	*(*int8)(unsafe.Pointer(bp)) = int8('\n')
	/* Read the server response */
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup = libc.Xmalloc(tls, uint32(8))
	if !((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup != 0) {
		return 0
	}
	if libc.Uint32FromInt32(X_xcb_in_read_block(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup, int32(8))) != uint32(8) {
		return 0
	}
	tmp = libc.Xrealloc(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup, libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Flength)*int32(4))+uint32(8))
	if !(tmp != 0) {
		return 0
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup = tmp
	if X_xcb_in_read_block(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup+uintptr(8), libc.Int32FromUint16((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Flength)*int32(4)) <= 0 {
		return 0
	}
	/* 0 = failed, 2 = authenticate, 1 = success */
	switch libc.Int32FromUint8((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fstatus) {
	case 0: /* failed */
		setup = (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup
		libc.Xwrite(tls, int32(m_STDERR_FILENO), Xxcb_setup_failed_reason(tls, setup), libc.Uint32FromInt32(Xxcb_setup_failed_reason_length(tls, setup)))
		libc.Xwrite(tls, int32(m_STDERR_FILENO), bp, uint32(1))
		return 0
	case int32(2): /* authenticate */
		setup1 = (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup
		libc.Xwrite(tls, int32(m_STDERR_FILENO), Xxcb_setup_authenticate_reason(tls, setup1), libc.Uint32FromInt32(Xxcb_setup_authenticate_reason_length(tls, setup1)))
		libc.Xwrite(tls, int32(m_STDERR_FILENO), bp, uint32(1))
		return 0
	}
	return int32(1)
}

// C documentation
//
//	/* precondition: there must be something for us to write. */
func _write_vec(tls *libc.TLS, c uintptr, vector uintptr, count uintptr) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var cur, i, n int32
	var hdr, v1 uintptr
	var _ /* cmsgbuf at bp+0 */ struct {
		Fbuf         [0][76]int8
		Fcmsghdr     Tcmsghdr
		F__ccgo_pad2 [64]byte
	}
	var _ /* msg at bp+76 */ Tmsghdr
	_, _, _, _, _ = cur, hdr, i, n, v1
	n = *(*int32)(unsafe.Pointer(count))
	if n > int32(m_IOV_MAX) {
		n = int32(m_IOV_MAX)
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd != 0 {
		*(*Tmsghdr)(unsafe.Pointer(bp + 76)) = Tmsghdr{
			Fmsg_iov:        *(*uintptr)(unsafe.Pointer(vector)),
			Fmsg_iovlen:     n,
			Fmsg_control:    bp,
			Fmsg_controllen: (libc.Uint32FromInt64(12)+libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) & ^(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) + libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd)*libc.Uint32FromInt64(4),
		}
		if (*Tmsghdr)(unsafe.Pointer(bp+76)).Fmsg_controllen >= uint32(12) {
			v1 = (*Tmsghdr)(unsafe.Pointer(bp + 76)).Fmsg_control
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		hdr = v1
		(*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_len = (*(*Tmsghdr)(unsafe.Pointer(bp + 76))).Fmsg_controllen
		(*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_level = int32(m_SOL_SOCKET)
		(*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_type = int32(m_SCM_RIGHTS)
		libc.Xmemcpy(tls, hdr+libc.UintptrFromInt32(1)*12, c+4336+16568, libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd)*uint32(4))
		n = libc.Xsendmsg(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd, bp+76, 0)
		if n < 0 && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) {
			return int32(1)
		}
		i = 0
		for {
			if !(i < (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd) {
				break
			}
			libc.Xclose(tls, *(*int32)(unsafe.Pointer(c + 4336 + 16568 + uintptr(i)*4)))
			goto _2
		_2:
			;
			i++
		}
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd = 0
	} else {
		n = libc.Xwritev(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd, *(*uintptr)(unsafe.Pointer(vector)), n)
		if n < 0 && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) {
			return int32(1)
		}
	}
	if n <= 0 {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_ERROR))
		return 0
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Ftotal_written += libc.Uint64FromInt32(n)
	for {
		if !(*(*int32)(unsafe.Pointer(count)) != 0) {
			break
		}
		cur = libc.Int32FromUint32((*Tiovec)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(vector)))).Fiov_len)
		if cur > n {
			cur = n
		}
		*(*Tsize_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(vector)) + 4)) -= libc.Uint32FromInt32(cur)
		(*Tiovec)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(vector)))).Fiov_base = (*Tiovec)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(vector)))).Fiov_base + uintptr(cur)
		n -= cur
		if (*Tiovec)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(vector)))).Fiov_len != 0 {
			break
		}
		goto _3
	_3:
		;
		*(*int32)(unsafe.Pointer(count))--
		*(*uintptr)(unsafe.Pointer(vector)) += 8
	}
	if !(*(*int32)(unsafe.Pointer(count)) != 0) {
		*(*uintptr)(unsafe.Pointer(vector)) = uintptr(0)
	}
	return int32(1)
}

/* Public interface */

func Xxcb_get_setup(tls *libc.TLS, c uintptr) (r uintptr) {
	if _is_static_error_conn(tls, c) != 0 {
		return uintptr(unsafe.Pointer(&_xcb_error_setup))
	}
	/* doesn't need locking because it's never written to. */
	return (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup
}

func Xxcb_get_file_descriptor(tls *libc.TLS, c uintptr) (r int32) {
	if _is_static_error_conn(tls, c) != 0 {
		return -int32(1)
	}
	/* doesn't need locking because it's never written to. */
	return (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd
}

func Xxcb_connection_has_error(tls *libc.TLS, c uintptr) (r int32) {
	/* doesn't need locking because it's read and written atomically. */
	return (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error
}

func Xxcb_connect_to_fd(tls *libc.TLS, fd int32, auth_info uintptr) (r uintptr) {
	var c uintptr
	_ = c
	c = libc.Xcalloc(tls, uint32(1), uint32(21056))
	if !(c != 0) {
		libc.Xclose(tls, fd)
		return X_xcb_conn_ret_error(tls, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Ffd = fd
	if !(_set_fd_flags(tls, fd) != 0 && libc.Xpthread_mutex_init(tls, c+12, uintptr(0)) == 0 && X_xcb_in_init(tls, c+40) != 0 && X_xcb_out_init(tls, c+4336) != 0 && _write_setup(tls, c, auth_info) != 0 && _read_setup(tls, c) != 0 && X_xcb_ext_init(tls, c) != 0 && X_xcb_xid_init(tls, c) != 0) {
		Xxcb_disconnect(tls, c)
		return X_xcb_conn_ret_error(tls, int32(m_XCB_CONN_ERROR))
	}
	return c
}

func Xxcb_disconnect(tls *libc.TLS, c uintptr) {
	if c == libc.UintptrFromInt32(0) || _is_static_error_conn(tls, c) != 0 {
		return
	}
	libc.Xfree(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)
	/* disallow further sends and receives */
	libc.Xshutdown(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd, int32(m_SHUT_RDWR))
	libc.Xclose(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd)
	libc.Xpthread_mutex_destroy(tls, c+12)
	X_xcb_in_destroy(tls, c+40)
	X_xcb_out_destroy(tls, c+4336)
	X_xcb_ext_destroy(tls, c)
	X_xcb_xid_destroy(tls, c)
	libc.Xfree(tls, c)
}

/* Private interface */

func X_xcb_conn_shutdown(tls *libc.TLS, c uintptr, err int32) {
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error = err
}

// C documentation
//
//	/* Return connection error state.
//	 * To make thread-safe, I need a seperate static
//	 * variable for every possible error.
//	 * has_error is the first field in xcb_connection_t, so just
//	 * return a casted int here; checking has_error (and only
//	 * has_error) will be safe.
//	 */
func X_xcb_conn_ret_error(tls *libc.TLS, err int32) (r uintptr) {
	switch err {
	case int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT):
		return uintptr(unsafe.Pointer(&_xcb_con_closed_mem_er))
	case int32(m_XCB_CONN_CLOSED_PARSE_ERR):
		return uintptr(unsafe.Pointer(&_xcb_con_closed_parse_er))
	case int32(m_XCB_CONN_CLOSED_INVALID_SCREEN):
		return uintptr(unsafe.Pointer(&_xcb_con_closed_screen_er))
	case int32(m_XCB_CONN_ERROR):
		fallthrough
	default:
		return uintptr(unsafe.Pointer(&_xcb_con_error))
	}
	return r
}

func X_xcb_conn_wait(tls *libc.TLS, c uintptr, cond uintptr, vector uintptr, count uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var may_read, ret, v1 int32
	var p2 uintptr
	var _ /* fd at bp+0 */ Tpollfd
	_, _, _, _ = may_read, ret, v1, p2
	/* If the thing I should be doing is already being done, wait for it. */
	if count != 0 {
		v1 = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fwriting
	} else {
		v1 = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading
	}
	if v1 != 0 {
		libc.Xpthread_cond_wait(tls, cond, c+12)
		return int32(1)
	}
	libc.Xmemset(tls, bp, 0, uint32(8))
	(*(*Tpollfd)(unsafe.Pointer(bp))).Ffd = (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd
	(*(*Tpollfd)(unsafe.Pointer(bp))).Fevents = int16(m_POLLIN)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading++
	if count != 0 {
		p2 = bp + 4
		*(*int16)(unsafe.Pointer(p2)) = int16(int32(*(*int16)(unsafe.Pointer(p2))) | libc.Int32FromInt32(m_POLLOUT))
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fwriting++
	}
	libc.Xpthread_mutex_unlock(tls, c+12)
	for cond1 := true; cond1; cond1 = ret == -int32(1) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EINTR) {
		ret = libc.Xpoll(tls, bp, uint32(1), -int32(1))
		/* If poll() returns an event we didn't expect, such as POLLNVAL, treat
		 * it as if it failed. */
		if ret >= 0 && int32((*(*Tpollfd)(unsafe.Pointer(bp))).Frevents) & ^int32((*(*Tpollfd)(unsafe.Pointer(bp))).Fevents) != 0 {
			ret = -int32(1)
			break
		}
	}
	if ret < 0 {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_ERROR))
		ret = 0
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	if ret != 0 {
		/* The code allows two threads to call select()/poll() at the same time.
		 * First thread just wants to read, a second thread wants to write, too.
		 * We have to make sure that we don't steal the reading thread's reply
		 * and let it get stuck in select()/poll().
		 * So a thread may read if either:
		 * - There is no other thread that wants to read (the above situation
		 *   did not occur).
		 * - It is the reading thread (above situation occurred).
		 */
		may_read = libc.BoolInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading == int32(1) || !(count != 0))
		if may_read != 0 && int32((*(*Tpollfd)(unsafe.Pointer(bp))).Frevents)&int32(m_POLLIN) != 0 {
			ret = libc.BoolInt32(ret != 0 && X_xcb_in_read(tls, c) != 0)
		}
		if int32((*(*Tpollfd)(unsafe.Pointer(bp))).Frevents)&int32(m_POLLOUT) != 0 {
			ret = libc.BoolInt32(ret != 0 && _write_vec(tls, c, vector, count) != 0)
		}
	}
	if count != 0 {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fwriting--
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading--
	return ret
}

func Xxcb_total_read(tls *libc.TLS, c uintptr) (r Tuint64_t) {
	var n Tuint64_t
	_ = n
	if Xxcb_connection_has_error(tls, c) != 0 {
		return uint64(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	n = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Ftotal_read
	libc.Xpthread_mutex_unlock(tls, c+12)
	return n
}

func Xxcb_total_written(tls *libc.TLS, c uintptr) (r Tuint64_t) {
	var n Tuint64_t
	_ = n
	if Xxcb_connection_has_error(tls, c) != 0 {
		return uint64(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	n = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Ftotal_written
	libc.Xpthread_mutex_unlock(tls, c+12)
	return n
}

const m_SEEK_CUR = 1
const m_SEEK_END = 2
const m_SEEK_SET = 0

type Txcb_extension_t = struct {
	Fname      uintptr
	Fglobal_id int32
}

type Txcb_extension_t1 = struct {
	Fname      uintptr
	Fglobal_id int32
}

type Txcb_protocol_request_t = struct {
	Fcount  Tsize_t
	Fext    uintptr
	Fopcode Tuint8_t
	Fisvoid Tuint8_t
}

type _xcb_send_request_flags_t = int32

const _XCB_REQUEST_CHECKED = 1
const _XCB_REQUEST_RAW = 2
const _XCB_REQUEST_DISCARD_REPLY = 4
const _XCB_REQUEST_REPLY_FDS = 8

/*
 * This file generated automatically from bigreq.xml by c_client.py.
 * Edit at your peril.
 */

/**
 * @defgroup XCB_BigRequests_API XCB BigRequests API
 * @brief BigRequests XCB Protocol Implementation.
 * @{
 **/

/**
 * @}
 */

func _send_request(tls *libc.TLS, c uintptr, isvoid int32, workaround _workarounds, flags int32, vector uintptr, count int32) {
	var p1 uintptr
	_ = p1
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest++
	if !(isvoid != 0) {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_expected = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest
	}
	if workaround != int32(_WORKAROUND_NONE) || flags != 0 {
		X_xcb_in_expect_reply(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest, workaround, flags)
	}
	for count != 0 && libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len)+(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len <= uint32(16384) {
		libc.Xmemcpy(tls, c+4336+112+uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len), (*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base, (*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len)
		p1 = c + 4336 + 16496
		*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + (*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len)
		(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base = (*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base + uintptr((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len)
		(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len = uint32(0)
		vector += 8
		count--
	}
	if !(count != 0) {
		return
	}
	vector -= 8
	count++
	(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base = c + 4336 + 112
	(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len = libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len = 0
	X_xcb_out_send(tls, c, vector, count)
}

func _send_sync(tls *libc.TLS, c uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* vector at bp+0 */ [2]Tiovec
	(*(*[2]Tiovec)(unsafe.Pointer(bp)))[int32(1)].Fiov_base = uintptr(unsafe.Pointer(&_sync_req))
	(*(*[2]Tiovec)(unsafe.Pointer(bp)))[int32(1)].Fiov_len = uint32(4)
	_send_request(tls, c, 0, int32(_WORKAROUND_NONE), int32(_XCB_REQUEST_DISCARD_REPLY), bp+uintptr(1)*8, int32(1))
}

var _sync_req = *(*struct {
	Fpacket [0]Tuint32_t
	Ffields struct {
		Fmajor Tuint8_t
		Fpad   Tuint8_t
		Flen1  Tuint16_t
	}
})(unsafe.Pointer(&struct {
	Fmajor uint8
	Fpad   uint8
	Flen1  uint16
}{
	Fmajor: uint8(43),
	Flen1:  uint16(1),
}))

func _get_socket_back(tls *libc.TLS, c uintptr) {
	for (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Freturn_socket != 0 && (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fsocket_moving != 0 {
		libc.Xpthread_cond_wait(tls, c+4336+52, c+12)
	}
	if !((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Freturn_socket != 0) {
		return
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fsocket_moving = int32(1)
	libc.Xpthread_mutex_unlock(tls, c+12)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Freturn_socket})))(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fsocket_closure)
	libc.Xpthread_mutex_lock(tls, c+12)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fsocket_moving = 0
	libc.Xpthread_cond_broadcast(tls, c+4336+52)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Freturn_socket = uintptr(0)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fsocket_closure = uintptr(0)
	X_xcb_in_replies_done(tls, c)
}

func _prepare_socket_request(tls *libc.TLS, c uintptr) {
	/* We're about to append data to out.queue, so we need to
	 * atomically test for an external socket owner *and* some other
	 * thread currently writing.
	 *
	 * If we have an external socket owner, we have to get the socket back
	 * before we can use it again.
	 *
	 * If some other thread is writing to the socket, we assume it's
	 * writing from out.queue, and so we can't stick data there.
	 *
	 * We satisfy this condition by first calling get_socket_back
	 * (which may drop the lock, but will return when XCB owns the
	 * socket again) and then checking for another writing thread and
	 * escaping the loop if we're ready to go.
	 */
	for {
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
			return
		}
		_get_socket_back(tls, c)
		if !((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fwriting != 0) {
			break
		}
		libc.Xpthread_cond_wait(tls, c+4336, c+12)
		goto _1
	_1:
	}
}

/* Public interface */

func Xxcb_prefetch_maximum_request_length(tls *libc.TLS, c uintptr) {
	var ext uintptr
	_ = ext
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return
	}
	libc.Xpthread_mutex_lock(tls, c+4336+16532)
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fmaximum_request_length_tag == int32(_LAZY_NONE) {
		ext = Xxcb_get_extension_data(tls, c, uintptr(unsafe.Pointer(&Xxcb_big_requests_id)))
		if ext != 0 && (*Txcb_query_extension_reply_t)(unsafe.Pointer(ext)).Fpresent != 0 {
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fmaximum_request_length_tag = int32(_LAZY_COOKIE)
			*(*Txcb_big_requests_enable_cookie_t)(unsafe.Pointer(c + 4336 + 16564)) = Xxcb_big_requests_enable(tls, c)
		} else {
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fmaximum_request_length_tag = int32(_LAZY_FORCED)
			*(*Tuint32_t)(unsafe.Pointer(c + 4336 + 16564)) = uint32((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fmaximum_request_length)
		}
	}
	libc.Xpthread_mutex_unlock(tls, c+4336+16532)
}

func Xxcb_get_maximum_request_length(tls *libc.TLS, c uintptr) (r1 Tuint32_t) {
	var r uintptr
	_ = r
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uint32(0)
	}
	Xxcb_prefetch_maximum_request_length(tls, c)
	libc.Xpthread_mutex_lock(tls, c+4336+16532)
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fmaximum_request_length_tag == int32(_LAZY_COOKIE) {
		r = Xxcb_big_requests_enable_reply(tls, c, *(*Txcb_big_requests_enable_cookie_t)(unsafe.Pointer(c + 4336 + 16564)), uintptr(0))
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fmaximum_request_length_tag = int32(_LAZY_FORCED)
		if r != 0 {
			*(*Tuint32_t)(unsafe.Pointer(c + 4336 + 16564)) = (*Txcb_big_requests_enable_reply_t)(unsafe.Pointer(r)).Fmaximum_request_length
			libc.Xfree(tls, r)
		} else {
			*(*Tuint32_t)(unsafe.Pointer(c + 4336 + 16564)) = uint32((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fmaximum_request_length)
		}
	}
	libc.Xpthread_mutex_unlock(tls, c+4336+16532)
	return *(*Tuint32_t)(unsafe.Pointer(c + 4336 + 16564))
}

func _close_fds(tls *libc.TLS, fds uintptr, num_fds uint32) {
	var index uint32
	_ = index
	index = uint32(0)
	for {
		if !(index < num_fds) {
			break
		}
		libc.Xclose(tls, *(*int32)(unsafe.Pointer(fds + uintptr(index)*4)))
		goto _1
	_1:
		;
		index++
	}
}

func _send_fds(tls *libc.TLS, c uintptr, fds uintptr, num_fds uint32) {
	var v1 int32
	var v2 uintptr
	_, _ = v1, v2
	/* Calling _xcb_out_flush_to() can drop the iolock and wait on a condition
	 * variable if another thread is currently writing (c->out.writing > 0).
	 * This call waits for writers to be done and thus _xcb_out_flush_to() will
	 * do the work itself (in which case we are a writer and
	 * prepare_socket_request() will wait for us to be done if another threads
	 * tries to send fds, too). Thanks to this, we can atomically write out FDs.
	 */
	_prepare_socket_request(tls, c)
	for num_fds > uint32(0) {
		for (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd == int32(m_XCB_MAX_PASS_FD) && !((*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0) {
			/* XXX: if c->out.writing > 0, this releases the iolock and
			 * potentially allows other threads to interfere with their own fds.
			 */
			X_xcb_out_flush_to(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest)
			if (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fout_fd.Fnfd == int32(m_XCB_MAX_PASS_FD) {
				/* We need some request to send FDs with */
				X_xcb_out_send_sync(tls, c)
			}
		}
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
			break
		}
		v2 = c + 4336 + 16568 + 64
		v1 = *(*int32)(unsafe.Pointer(v2))
		*(*int32)(unsafe.Pointer(v2))++
		*(*int32)(unsafe.Pointer(c + 4336 + 16568 + uintptr(v1)*4)) = *(*int32)(unsafe.Pointer(fds))
		fds += 4
		num_fds--
	}
	_close_fds(tls, fds, num_fds)
}

func Xxcb_send_request_with_fds64(tls *libc.TLS, c uintptr, flags int32, vector uintptr, req uintptr, num_fds uint32, fds uintptr) (r Tuint64_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var extension uintptr
	var i uint32
	var longlen, v2 Tsize_t
	var request Tuint64_t
	var shortlen Tuint16_t
	var veclen int32
	var workaround _workarounds
	var v3 uint64
	var _ /* prefix at bp+0 */ [2]Tuint32_t
	_, _, _, _, _, _, _, _, _ = extension, i, longlen, request, shortlen, veclen, workaround, v2, v3
	veclen = libc.Int32FromUint32((*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fcount)
	workaround = int32(_WORKAROUND_NONE)
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		_close_fds(tls, fds, num_fds)
		return uint64(0)
	}
	if !(flags&int32(_XCB_REQUEST_RAW) != 0) {
		shortlen = uint16(0)
		longlen = uint32(0)
		/* set the major opcode, and the minor opcode for extensions */
		if (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fext != 0 {
			extension = Xxcb_get_extension_data(tls, c, (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fext)
			if !(extension != 0 && (*Txcb_query_extension_reply_t)(unsafe.Pointer(extension)).Fpresent != 0) {
				_close_fds(tls, fds, num_fds)
				X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_EXT_NOTSUPPORTED))
				return uint64(0)
			}
			*(*Tuint8_t)(unsafe.Pointer((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base)) = (*Txcb_query_extension_reply_t)(unsafe.Pointer(extension)).Fmajor_opcode
			*(*Tuint8_t)(unsafe.Pointer((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base + 1)) = (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fopcode
		} else {
			*(*Tuint8_t)(unsafe.Pointer((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base)) = (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fopcode
		}
		/* put together the length field, possibly using BIGREQUESTS */
		i = uint32(0)
		for {
			if !(i < (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fcount) {
				break
			}
			longlen += (*(*Tiovec)(unsafe.Pointer(vector + uintptr(i)*8))).Fiov_len
			if !((*(*Tiovec)(unsafe.Pointer(vector + uintptr(i)*8))).Fiov_base != 0) {
				(*(*Tiovec)(unsafe.Pointer(vector + uintptr(i)*8))).Fiov_base = uintptr(unsafe.Pointer(&_pad1))
			}
			goto _1
		_1:
			;
			i++
		}
		longlen >>= uint32(2)
		if longlen <= uint32((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fmaximum_request_length) {
			/* we don't need BIGREQUESTS. */
			shortlen = uint16(longlen)
			longlen = uint32(0)
		} else {
			if longlen > Xxcb_get_maximum_request_length(tls, c) {
				_close_fds(tls, fds, num_fds)
				X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_REQ_LEN_EXCEED))
				return uint64(0) /* server can't take this; maybe need BIGREQUESTS? */
			}
		}
		/* set the length field. */
		*(*Tuint16_t)(unsafe.Pointer((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base + 1*2)) = shortlen
		if !(shortlen != 0) {
			(*(*[2]Tuint32_t)(unsafe.Pointer(bp)))[0] = *(*Tuint32_t)(unsafe.Pointer((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base))
			longlen++
			v2 = longlen
			(*(*[2]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] = v2
			(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base = (*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base + uintptr(1)*4
			(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len -= uint32(4)
			vector -= 8
			veclen++
			(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base = bp
			(*(*Tiovec)(unsafe.Pointer(vector))).Fiov_len = uint32(8)
		}
	}
	flags &= ^int32(_XCB_REQUEST_RAW)
	/* do we need to work around the X server bug described in glx.xml? */
	/* XXX: GetFBConfigs won't use BIG-REQUESTS in any sane
	 * configuration, but that should be handled here anyway. */
	if (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fext != 0 && !((*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fisvoid != 0) && !(libc.Xstrcmp(tls, (*Txcb_extension_t)(unsafe.Pointer((*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fext)).Fname, __ccgo_ts) != 0) && (libc.Int32FromUint8((*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fopcode) == int32(17) && *(*Tuint32_t)(unsafe.Pointer((*(*Tiovec)(unsafe.Pointer(vector))).Fiov_base + 1*4)) == uint32(0x10004) || libc.Int32FromUint8((*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fopcode) == int32(21)) {
		workaround = int32(_WORKAROUND_GLX_GET_FB_CONFIGS_BUG)
	}
	/* get a sequence number and arrange for delivery. */
	libc.Xpthread_mutex_lock(tls, c+12)
	/* send FDs before establishing a good request number, because this might
	 * call send_sync(), too
	 */
	_send_fds(tls, c, fds, num_fds)
	_prepare_socket_request(tls, c)
	/* send GetInputFocus (sync_req) when 64k-2 requests have been sent without
	 * a reply.
	 * Also send sync_req (could use NoOp) at 32-bit wrap to avoid having
	 * applications see sequence 0 as that is used to indicate
	 * an error in sending the request
	 */
	for (*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fisvoid != 0 && (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest == (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_expected+libc.Uint64FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(16))-uint64(2) || uint32((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest+libc.Uint64FromInt32(1)) == uint32(0) {
		_send_sync(tls, c)
		_prepare_socket_request(tls, c)
	}
	_send_request(tls, c, libc.Int32FromUint8((*Txcb_protocol_request_t)(unsafe.Pointer(req)).Fisvoid), workaround, flags, vector, veclen)
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		v3 = uint64(0)
	} else {
		v3 = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest
	}
	request = v3
	libc.Xpthread_mutex_unlock(tls, c+12)
	return request
}

var _pad1 [3]int8

// C documentation
//
//	/* request number are actually uint64_t internally but keep API compat with unsigned int */
func Xxcb_send_request_with_fds(tls *libc.TLS, c uintptr, flags int32, vector uintptr, req uintptr, num_fds uint32, fds uintptr) (r uint32) {
	return uint32(Xxcb_send_request_with_fds64(tls, c, flags, vector, req, num_fds, fds))
}

func Xxcb_send_request64(tls *libc.TLS, c uintptr, flags int32, vector uintptr, req uintptr) (r Tuint64_t) {
	return Xxcb_send_request_with_fds64(tls, c, flags, vector, req, uint32(0), libc.UintptrFromInt32(0))
}

// C documentation
//
//	/* request number are actually uint64_t internally but keep API compat with unsigned int */
func Xxcb_send_request(tls *libc.TLS, c uintptr, flags int32, vector uintptr, req uintptr) (r uint32) {
	return uint32(Xxcb_send_request64(tls, c, flags, vector, req))
}

func Xxcb_send_fd(tls *libc.TLS, c uintptr, fd int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* fds at bp+0 */ [1]int32
	*(*[1]int32)(unsafe.Pointer(bp)) = [1]int32{
		0: fd,
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		libc.Xclose(tls, fd)
		return
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	_send_fds(tls, c, bp, uint32(1))
	libc.Xpthread_mutex_unlock(tls, c+12)
}

func Xxcb_take_socket(tls *libc.TLS, c uintptr, return_socket uintptr, closure uintptr, flags int32, sent uintptr) (r int32) {
	var ret int32
	_ = ret
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return 0
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	_get_socket_back(tls, c)
	/* _xcb_out_flush may drop the iolock allowing other threads to
	 * write requests, so keep flushing until we're done
	 */
	for cond := true; cond; cond = ret != 0 && (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest != (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest_written {
		ret = X_xcb_out_flush_to(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest)
	}
	if ret != 0 {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Freturn_socket = return_socket
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fsocket_closure = closure
		if flags != 0 {
			/* c->out.request + 1 will be the first request sent by the external
			 * socket owner. If the socket is returned before this request is sent
			 * it will be detected in _xcb_in_replies_done and this pending_reply
			 * will be discarded.
			 */
			X_xcb_in_expect_reply(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest+uint64(1), int32(_WORKAROUND_EXTERNAL_SOCKET_OWNER), flags)
		}
		*(*Tuint64_t)(unsafe.Pointer(sent)) = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest
	}
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func Xxcb_writev(tls *libc.TLS, c uintptr, vector uintptr, count int32, requests Tuint64_t) (r int32) {
	var ret int32
	_ = ret
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return 0
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest += requests
	ret = X_xcb_out_send(tls, c, vector, count)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func Xxcb_flush(tls *libc.TLS, c uintptr) (r int32) {
	var ret int32
	_ = ret
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return 0
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	ret = X_xcb_out_flush_to(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

/* Private interface */

func X_xcb_out_init(tls *libc.TLS, out uintptr) (r int32) {
	if libc.Xpthread_cond_init(tls, out+52, uintptr(0)) != 0 {
		return 0
	}
	(*T_xcb_out)(unsafe.Pointer(out)).Freturn_socket = uintptr(0)
	(*T_xcb_out)(unsafe.Pointer(out)).Fsocket_closure = uintptr(0)
	(*T_xcb_out)(unsafe.Pointer(out)).Fsocket_moving = 0
	if libc.Xpthread_cond_init(tls, out, uintptr(0)) != 0 {
		return 0
	}
	(*T_xcb_out)(unsafe.Pointer(out)).Fwriting = 0
	(*T_xcb_out)(unsafe.Pointer(out)).Fqueue_len = 0
	(*T_xcb_out)(unsafe.Pointer(out)).Frequest = uint64(0)
	(*T_xcb_out)(unsafe.Pointer(out)).Frequest_written = uint64(0)
	(*T_xcb_out)(unsafe.Pointer(out)).Frequest_expected_written = uint64(0)
	if libc.Xpthread_mutex_init(tls, out+16532, uintptr(0)) != 0 {
		return 0
	}
	(*T_xcb_out)(unsafe.Pointer(out)).Fmaximum_request_length_tag = int32(_LAZY_NONE)
	return int32(1)
}

func X_xcb_out_destroy(tls *libc.TLS, out uintptr) {
	libc.Xpthread_mutex_destroy(tls, out+16532)
	libc.Xpthread_cond_destroy(tls, out)
	libc.Xpthread_cond_destroy(tls, out+52)
}

func X_xcb_out_send(tls *libc.TLS, c uintptr, _vector uintptr, _count int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uintptr)(unsafe.Pointer(bp)) = _vector
	*(*int32)(unsafe.Pointer(bp + 4)) = _count
	var ret int32
	_ = ret
	ret = int32(1)
	for ret != 0 && *(*int32)(unsafe.Pointer(bp + 4)) != 0 {
		ret = X_xcb_conn_wait(tls, c, c+4336, bp, bp+4)
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest_written = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest_expected_written = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_expected
	libc.Xpthread_cond_broadcast(tls, c+4336)
	X_xcb_in_wake_up_next_reader(tls, c)
	return ret
}

func X_xcb_out_send_sync(tls *libc.TLS, c uintptr) {
	_prepare_socket_request(tls, c)
	_send_sync(tls, c)
}

func X_xcb_out_flush_to(tls *libc.TLS, c uintptr, request Tuint64_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* vec at bp+0 */ Tiovec
	if libc.Int64FromUint64((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest_written-request) >= 0 {
		return int32(1)
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len != 0 {
		(*(*Tiovec)(unsafe.Pointer(bp))).Fiov_base = c + 4336 + 112
		(*(*Tiovec)(unsafe.Pointer(bp))).Fiov_len = libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len)
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fqueue_len = 0
		return X_xcb_out_send(tls, c, bp, int32(1))
	}
	for (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Fwriting != 0 {
		libc.Xpthread_cond_wait(tls, c+4336, c+12)
	}
	return int32(1)
}

const m_INT32_MAX1 = 2147483647
const m_MSG_CTRUNC1 = 8
const m_MSG_TRUNC1 = 32
const m_POLLOUT1 = 0x004
const m_UINT32_MAX1 = 4294967295
const m_XCB_ERROR = 0
const m_XCB_REPLY = 1
const m_XCB_XGE_EVENT = 35

type Txcb_special_event_t = struct {
	Fnext               uintptr
	Fextension          Tuint8_t
	Feid                Tuint32_t
	Fstamp              uintptr
	Fevents             uintptr
	Fevents_tail        uintptr
	Fspecial_event_cond Tpthread_cond_t
}

type Txcb_special_event = Txcb_special_event_t

type Tevent_list = struct {
	Fevent uintptr
	Fnext  uintptr
}

type Treply_list = struct {
	Freply uintptr
	Fnext  uintptr
}

type Tpending_reply = struct {
	Ffirst_request Tuint64_t
	Flast_request  Tuint64_t
	Fworkaround    _workarounds
	Fflags         int32
	Fnext          uintptr
}

type Treader_list = struct {
	Frequest Tuint64_t
	Fdata    uintptr
	Fnext    uintptr
}

type Tspecial_list = struct {
	Fse   uintptr
	Fnext uintptr
}

func _remove_finished_readers(tls *libc.TLS, prev_reader uintptr, completed Tuint64_t) {
	for *(*uintptr)(unsafe.Pointer(prev_reader)) != 0 && libc.Int64FromUint64((*Treader_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_reader)))).Frequest-completed) <= 0 {
		/* If you don't have what you're looking for now, you never
		 * will. Wake up and leave me alone. */
		libc.Xpthread_cond_signal(tls, (*Treader_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_reader)))).Fdata)
		*(*uintptr)(unsafe.Pointer(prev_reader)) = (*Treader_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_reader)))).Fnext
	}
}

func _read_fds(tls *libc.TLS, c uintptr, fds uintptr, nfd int32) (r int32) {
	var ifds uintptr
	var infd int32
	_, _ = ifds, infd
	ifds = c + 40 + 4220 + uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fifd)*4
	infd = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd - (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fifd
	if nfd > infd {
		return 0
	}
	libc.Xmemcpy(tls, fds, ifds, libc.Uint32FromInt32(nfd)*uint32(4))
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fifd += nfd
	return int32(1)
}

type Txcb_ge_special_event_t = struct {
	Fresponse_type Tuint8_t
	Fextension     Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fevtype        Tuint16_t
	Fpad0          [2]Tuint8_t
	Feid           Tuint32_t
	Fpad1          [16]Tuint8_t
}

func _event_special(tls *libc.TLS, c uintptr, event uintptr) (r int32) {
	var ges, special_event uintptr
	_, _ = ges, special_event
	ges = (*Tevent_list)(unsafe.Pointer(event)).Fevent
	/* Special events are always XGE events */
	if libc.Int32FromUint8((*Txcb_ge_special_event_t1)(unsafe.Pointer(ges)).Fresponse_type)&int32(0x7f) != int32(m_XCB_XGE_EVENT) {
		return 0
	}
	special_event = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fspecial_events
	for {
		if !(special_event != 0) {
			break
		}
		if libc.Int32FromUint8((*Txcb_ge_special_event_t1)(unsafe.Pointer(ges)).Fextension) == libc.Int32FromUint8((*Txcb_special_event)(unsafe.Pointer(special_event)).Fextension) && (*Txcb_ge_special_event_t1)(unsafe.Pointer(ges)).Feid == (*Txcb_special_event)(unsafe.Pointer(special_event)).Feid {
			*(*uintptr)(unsafe.Pointer((*Txcb_special_event)(unsafe.Pointer(special_event)).Fevents_tail)) = event
			(*Txcb_special_event)(unsafe.Pointer(special_event)).Fevents_tail = event + 4
			if (*Txcb_special_event)(unsafe.Pointer(special_event)).Fstamp != 0 {
				*(*Tuint32_t)(unsafe.Pointer((*Txcb_special_event)(unsafe.Pointer(special_event)).Fstamp))++
			}
			libc.Xpthread_cond_signal(tls, special_event+24)
			return int32(1)
		}
		goto _1
	_1:
		;
		special_event = (*Txcb_special_event)(unsafe.Pointer(special_event)).Fnext
	}
	return 0
}

func _read_packet(tls *libc.TLS, c uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var buf, cur, event, oldpend, p, pend uintptr
	var bufsize, eventlength, lastread, length, new_length Tuint64_t
	var nfd int32
	var v1 uint32
	var _ /* genrep at bp+0 */ Txcb_generic_reply_t
	_, _, _, _, _, _, _, _, _, _, _, _, _ = buf, bufsize, cur, event, eventlength, lastread, length, new_length, nfd, oldpend, p, pend, v1
	length = uint64(32)
	eventlength = uint64(0) /* length after first 32 bytes for GenericEvents */
	nfd = 0
	pend = uintptr(0)
	/* Wait for there to be enough data for us to read a whole packet */
	if libc.Uint64FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len) < length {
		return 0
	}
	/* Get the response type, length, and sequence number. */
	libc.Xmemcpy(tls, bp, c+40+52, uint32(8))
	/* Compute 32-bit sequence number of this packet. */
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type)&int32(0x7f) != int32(m_XCB_KEYMAP_NOTIFY) {
		lastread = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read = lastread&uint64(0xffffffffffff0000) | uint64((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fsequence)
		if libc.Int64FromUint64((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read-lastread) < 0 {
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read += uint64(0x10000)
		}
		if libc.Int64FromUint64((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_expected) > 0 {
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_expected = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read
		}
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read != lastread {
			if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply != 0 {
				X_xcb_map_put(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freplies, lastread, (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply)
				(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply = uintptr(0)
				(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply_tail = c + 40 + 4184
			}
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read - uint64(1)
		}
		for (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies != 0 && (*Tpending_reply1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies)).Fworkaround != int32(_WORKAROUND_EXTERNAL_SOCKET_OWNER) && libc.Int64FromUint64((*Tpending_reply1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies)).Flast_request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed) <= 0 {
			oldpend = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies = (*Tpending_reply)(unsafe.Pointer(oldpend)).Fnext
			if !((*Tpending_reply)(unsafe.Pointer(oldpend)).Fnext != 0) {
				(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail = c + 40 + 4212
			}
			libc.Xfree(tls, oldpend)
		}
		if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == m_XCB_ERROR {
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read
		}
		_remove_finished_readers(tls, c+40+4204, (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed)
	}
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == m_XCB_ERROR || libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == int32(m_XCB_REPLY) {
		pend = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies
		if pend != 0 && !(libc.Int64FromUint64((*Tpending_reply)(unsafe.Pointer(pend)).Ffirst_request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read) <= 0 && ((*Tpending_reply)(unsafe.Pointer(pend)).Fworkaround == int32(_WORKAROUND_EXTERNAL_SOCKET_OWNER) || libc.Int64FromUint64((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read-(*Tpending_reply)(unsafe.Pointer(pend)).Flast_request) <= 0)) {
			pend = uintptr(0)
		}
	}
	/* For reply packets, check that the entire packet is available. */
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == int32(m_XCB_REPLY) {
		if pend != 0 && (*Tpending_reply)(unsafe.Pointer(pend)).Fworkaround == int32(_WORKAROUND_GLX_GET_FB_CONFIGS_BUG) {
			p = c + 40 + 52
			new_length = uint64(*(*Tuint32_t)(unsafe.Pointer(p + 2*4))) * uint64(*(*Tuint32_t)(unsafe.Pointer(p + 3*4)))
			if new_length >= uint64(libc.Uint32FromUint32(0xffffffff)/libc.Uint32FromUint32(16)) {
				X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
				return 0
			}
			(*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Flength = uint32(new_length * libc.Uint64FromUint64(2))
		}
		length += uint64((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Flength) * uint64(4)
		/* XXX a bit of a hack -- we "know" that all FD replys place
		 * the number of fds in the pad0 byte */
		if pend != 0 && (*Tpending_reply)(unsafe.Pointer(pend)).Fflags&int32(_XCB_REQUEST_REPLY_FDS) != 0 {
			nfd = libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fpad0)
		}
	}
	/* XGE events may have sizes > 32 */
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type)&int32(0x7f) == int32(m_XCB_XGE_EVENT) {
		eventlength = uint64((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Flength) * uint64(4)
	}
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == int32(m_XCB_REPLY) {
		v1 = uint32(0)
	} else {
		v1 = uint32(4)
	}
	bufsize = length + eventlength + uint64(libc.Uint32FromInt32(nfd)*uint32(4)) + uint64(v1)
	if bufsize < libc.Uint64FromInt32(libc.Int32FromInt32(m_INT32_MAX1)) {
		buf = libc.Xmalloc(tls, uint32(bufsize))
	} else {
		buf = libc.UintptrFromInt32(0)
	}
	if !(buf != 0) {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
		return 0
	}
	if X_xcb_in_read_block(tls, c, buf, libc.Int32FromUint64(length)) <= 0 {
		libc.Xfree(tls, buf)
		return 0
	}
	/* pull in XGE event data if available, append after event struct */
	if eventlength != 0 {
		if X_xcb_in_read_block(tls, c, buf+1*36, libc.Int32FromUint64(eventlength)) <= 0 {
			libc.Xfree(tls, buf)
			return 0
		}
	}
	if nfd != 0 {
		if !(_read_fds(tls, c, buf+uintptr(length), nfd) != 0) {
			libc.Xfree(tls, buf)
			return 0
		}
	}
	if pend != 0 && (*Tpending_reply)(unsafe.Pointer(pend)).Fflags&int32(_XCB_REQUEST_DISCARD_REPLY) != 0 {
		libc.Xfree(tls, buf)
		return int32(1)
	}
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) != int32(m_XCB_REPLY) {
		(*Txcb_generic_event_t)(unsafe.Pointer(buf)).Ffull_sequence = uint32((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read)
	}
	/* reply, or checked error */
	if libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == int32(m_XCB_REPLY) || libc.Int32FromUint8((*(*Txcb_generic_reply_t)(unsafe.Pointer(bp))).Fresponse_type) == m_XCB_ERROR && pend != 0 && (*Tpending_reply)(unsafe.Pointer(pend)).Fflags&int32(_XCB_REQUEST_CHECKED) != 0 {
		cur = libc.Xmalloc(tls, uint32(8))
		if !(cur != 0) {
			X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
			libc.Xfree(tls, buf)
			return 0
		}
		(*Treply_list)(unsafe.Pointer(cur)).Freply = buf
		(*Treply_list)(unsafe.Pointer(cur)).Fnext = uintptr(0)
		*(*uintptr)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply_tail)) = cur
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply_tail = cur + 4
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freaders != 0 && (*Treader_list1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freaders)).Frequest == (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read {
			libc.Xpthread_cond_signal(tls, (*Treader_list1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freaders)).Fdata)
		}
		return int32(1)
	}
	/* event, or unchecked error */
	event = libc.Xmalloc(tls, uint32(8))
	if !(event != 0) {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
		libc.Xfree(tls, buf)
		return 0
	}
	(*Tevent_list)(unsafe.Pointer(event)).Fevent = buf
	(*Tevent_list)(unsafe.Pointer(event)).Fnext = uintptr(0)
	if !(_event_special(tls, c, event) != 0) {
		*(*uintptr)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fevents_tail)) = event
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fevents_tail = event + 4
		libc.Xpthread_cond_signal(tls, c+40)
	}
	return int32(1) /* I have something for you... */
}

func _get_event(tls *libc.TLS, c uintptr) (r uintptr) {
	var cur, ret uintptr
	_, _ = cur, ret
	cur = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fevents
	if !((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fevents != 0) {
		return uintptr(0)
	}
	ret = (*Tevent_list)(unsafe.Pointer(cur)).Fevent
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fevents = (*Tevent_list)(unsafe.Pointer(cur)).Fnext
	if !((*Tevent_list)(unsafe.Pointer(cur)).Fnext != 0) {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fevents_tail = c + 40 + 4196
	}
	libc.Xfree(tls, cur)
	return ret
}

func _free_reply_list(tls *libc.TLS, head uintptr) {
	var cur uintptr
	_ = cur
	for head != 0 {
		cur = head
		head = (*Treply_list)(unsafe.Pointer(cur)).Fnext
		libc.Xfree(tls, (*Treply_list)(unsafe.Pointer(cur)).Freply)
		libc.Xfree(tls, cur)
	}
}

func _read_block(tls *libc.TLS, fd int32, buf uintptr, len1 Tintptr_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var done, ret int32
	var _ /* pfd at bp+0 */ Tpollfd
	_, _ = done, ret
	done = 0
	for done < len1 {
		ret = libc.Xrecv(tls, fd, buf+uintptr(done), libc.Uint32FromInt32(len1-done), 0)
		if ret > 0 {
			done += ret
		}
		if ret < 0 && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) {
			(*(*Tpollfd)(unsafe.Pointer(bp))).Ffd = fd
			(*(*Tpollfd)(unsafe.Pointer(bp))).Fevents = int16(m_POLLIN)
			(*(*Tpollfd)(unsafe.Pointer(bp))).Frevents = 0
			for cond := true; cond; cond = ret == -int32(1) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EINTR) {
				ret = libc.Xpoll(tls, bp, uint32(1), -int32(1))
			}
		}
		if ret <= 0 {
			return ret
		}
	}
	return len1
}

func _poll_for_reply(tls *libc.TLS, c uintptr, request Tuint64_t, reply uintptr, error1 uintptr) (r int32) {
	var head uintptr
	_ = head
	/* If an error occurred when issuing the request, fail immediately. */
	if !(request != 0) {
		head = uintptr(0)
	} else {
		if libc.Int64FromUint64(request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read) < 0 {
			head = X_xcb_map_remove(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freplies, request)
			if head != 0 && (*Treply_list)(unsafe.Pointer(head)).Fnext != 0 {
				X_xcb_map_put(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freplies, request, (*Treply_list)(unsafe.Pointer(head)).Fnext)
			}
		} else {
			if request == (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_read && (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply != 0 {
				head = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply
				(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply = (*Treply_list)(unsafe.Pointer(head)).Fnext
				if !((*Treply_list)(unsafe.Pointer(head)).Fnext != 0) {
					(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fcurrent_reply_tail = c + 40 + 4184
				}
			} else {
				if request == (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed {
					head = uintptr(0)
				} else {
					return 0
				}
			}
		}
	}
	if error1 != 0 {
		*(*uintptr)(unsafe.Pointer(error1)) = uintptr(0)
	}
	*(*uintptr)(unsafe.Pointer(reply)) = uintptr(0)
	if head != 0 {
		if libc.Int32FromUint8((*Txcb_generic_reply_t)(unsafe.Pointer((*Treply_list)(unsafe.Pointer(head)).Freply)).Fresponse_type) == m_XCB_ERROR {
			if error1 != 0 {
				*(*uintptr)(unsafe.Pointer(error1)) = (*Treply_list)(unsafe.Pointer(head)).Freply
			} else {
				libc.Xfree(tls, (*Treply_list)(unsafe.Pointer(head)).Freply)
			}
		} else {
			*(*uintptr)(unsafe.Pointer(reply)) = (*Treply_list)(unsafe.Pointer(head)).Freply
		}
		libc.Xfree(tls, head)
	}
	return int32(1)
}

func _insert_reader(tls *libc.TLS, prev_reader uintptr, reader uintptr, request Tuint64_t, cond uintptr) {
	for *(*uintptr)(unsafe.Pointer(prev_reader)) != 0 && libc.Int64FromUint64((*Treader_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_reader)))).Frequest-request) <= 0 {
		prev_reader = *(*uintptr)(unsafe.Pointer(prev_reader)) + 12
	}
	(*Treader_list)(unsafe.Pointer(reader)).Frequest = request
	(*Treader_list)(unsafe.Pointer(reader)).Fdata = cond
	(*Treader_list)(unsafe.Pointer(reader)).Fnext = *(*uintptr)(unsafe.Pointer(prev_reader))
	*(*uintptr)(unsafe.Pointer(prev_reader)) = reader
}

func _remove_reader(tls *libc.TLS, prev_reader uintptr, reader uintptr) {
	for *(*uintptr)(unsafe.Pointer(prev_reader)) != 0 && libc.Int64FromUint64((*Treader_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_reader)))).Frequest-(*Treader_list)(unsafe.Pointer(reader)).Frequest) <= 0 {
		if *(*uintptr)(unsafe.Pointer(prev_reader)) == reader {
			*(*uintptr)(unsafe.Pointer(prev_reader)) = (*Treader_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_reader)))).Fnext
			break
		}
	}
}

func _insert_special(tls *libc.TLS, prev_special uintptr, special uintptr, se uintptr) {
	(*Tspecial_list)(unsafe.Pointer(special)).Fse = se
	(*Tspecial_list)(unsafe.Pointer(special)).Fnext = *(*uintptr)(unsafe.Pointer(prev_special))
	*(*uintptr)(unsafe.Pointer(prev_special)) = special
}

func _remove_special(tls *libc.TLS, prev_special uintptr, special uintptr) {
	for *(*uintptr)(unsafe.Pointer(prev_special)) != 0 {
		if *(*uintptr)(unsafe.Pointer(prev_special)) == special {
			*(*uintptr)(unsafe.Pointer(prev_special)) = (*Tspecial_list)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_special)))).Fnext
			break
		}
		prev_special = *(*uintptr)(unsafe.Pointer(prev_special)) + 4
	}
}

func _wait_for_reply(tls *libc.TLS, c uintptr, request Tuint64_t, e uintptr) (r uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* cond at bp+4 */ Tpthread_cond_t
	var _ /* reader at bp+52 */ Treader_list
	var _ /* ret at bp+0 */ uintptr
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	/* If this request has not been written yet, write it. */
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Freturn_socket != 0 || X_xcb_out_flush_to(tls, c, request) != 0 {
		*(*Tpthread_cond_t)(unsafe.Pointer(bp + 4)) = Tpthread_cond_t{}
		*(*int32)(unsafe.Pointer(bp + 4)) = 0
		_insert_reader(tls, c+40+4204, bp+52, request, bp+4)
		for !(_poll_for_reply(tls, c, request, bp, e) != 0) {
			if !(X_xcb_conn_wait(tls, c, bp+4, uintptr(0), uintptr(0)) != 0) {
				break
			}
		}
		_remove_reader(tls, c+40+4204, bp+52)
		libc.Xpthread_cond_destroy(tls, bp+4)
	}
	X_xcb_in_wake_up_next_reader(tls, c)
	return *(*uintptr)(unsafe.Pointer(bp))
}

func _widen(tls *libc.TLS, c uintptr, request uint32) (r Tuint64_t) {
	var widened_request Tuint64_t
	_ = widened_request
	widened_request = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest&uint64(0xffffffff00000000) | uint64(request)
	if widened_request > (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest {
		widened_request -= libc.Uint64FromUint64(1) << libc.Int32FromInt32(32)
	}
	return widened_request
}

/* Public interface */

func Xxcb_wait_for_reply(tls *libc.TLS, c uintptr, request uint32, e uintptr) (r uintptr) {
	var ret uintptr
	_ = ret
	if e != 0 {
		*(*uintptr)(unsafe.Pointer(e)) = uintptr(0)
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	ret = _wait_for_reply(tls, c, _widen(tls, c, request), e)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func Xxcb_wait_for_reply64(tls *libc.TLS, c uintptr, request Tuint64_t, e uintptr) (r uintptr) {
	var ret uintptr
	_ = ret
	if e != 0 {
		*(*uintptr)(unsafe.Pointer(e)) = uintptr(0)
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	ret = _wait_for_reply(tls, c, request, e)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func Xxcb_get_reply_fds(tls *libc.TLS, c uintptr, reply uintptr, reply_size Tsize_t) (r uintptr) {
	return reply + uintptr(reply_size)
}

func _insert_pending_discard(tls *libc.TLS, c uintptr, prev_next uintptr, seq Tuint64_t) {
	var pend uintptr
	_ = pend
	pend = libc.Xmalloc(tls, uint32(28))
	if !(pend != 0) {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
		return
	}
	(*Tpending_reply)(unsafe.Pointer(pend)).Ffirst_request = seq
	(*Tpending_reply)(unsafe.Pointer(pend)).Flast_request = seq
	(*Tpending_reply)(unsafe.Pointer(pend)).Fworkaround = 0
	(*Tpending_reply)(unsafe.Pointer(pend)).Fflags = int32(_XCB_REQUEST_DISCARD_REPLY)
	(*Tpending_reply)(unsafe.Pointer(pend)).Fnext = *(*uintptr)(unsafe.Pointer(prev_next))
	*(*uintptr)(unsafe.Pointer(prev_next)) = pend
	if !((*Tpending_reply)(unsafe.Pointer(pend)).Fnext != 0) {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail = pend + 24
	}
}

func _discard_reply(tls *libc.TLS, c uintptr, request Tuint64_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var prev_pend uintptr
	var _ /* reply at bp+0 */ uintptr
	_ = prev_pend
	/* Free any replies or errors that we've already read. Stop if
	 * xcb_wait_for_reply would block or we've run out of replies. */
	for _poll_for_reply(tls, c, request, bp, uintptr(0)) != 0 && *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp)))
	}
	/* If we've proven there are no more responses coming, we're done. */
	if libc.Int64FromUint64(request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed) <= 0 {
		return
	}
	/* Walk the list of pending requests. Mark the first match for deletion. */
	prev_pend = c + 40 + 4212
	for {
		if !(*(*uintptr)(unsafe.Pointer(prev_pend)) != 0) {
			break
		}
		if libc.Int64FromUint64((*Tpending_reply)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_pend)))).Ffirst_request-request) > 0 {
			break
		}
		if (*Tpending_reply)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_pend)))).Ffirst_request == request {
			/* Pending reply found. Mark for discard: */
			*(*int32)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(prev_pend)) + 20)) |= int32(_XCB_REQUEST_DISCARD_REPLY)
			return
		}
		goto _1
	_1:
		;
		prev_pend = *(*uintptr)(unsafe.Pointer(prev_pend)) + 24
	}
	/* Pending reply not found (likely due to _unchecked request). Create one: */
	_insert_pending_discard(tls, c, prev_pend, request)
}

func Xxcb_discard_reply(tls *libc.TLS, c uintptr, sequence uint32) {
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return
	}
	/* If an error occurred when issuing the request, fail immediately. */
	if !(sequence != 0) {
		return
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	_discard_reply(tls, c, _widen(tls, c, sequence))
	libc.Xpthread_mutex_unlock(tls, c+12)
}

func Xxcb_discard_reply64(tls *libc.TLS, c uintptr, sequence Tuint64_t) {
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return
	}
	/* If an error occurred when issuing the request, fail immediately. */
	if !(sequence != 0) {
		return
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	_discard_reply(tls, c, sequence)
	libc.Xpthread_mutex_unlock(tls, c+12)
}

func Xxcb_poll_for_reply(tls *libc.TLS, c uintptr, request uint32, reply uintptr, error1 uintptr) (r int32) {
	var ret int32
	_ = ret
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		*(*uintptr)(unsafe.Pointer(reply)) = uintptr(0)
		if error1 != 0 {
			*(*uintptr)(unsafe.Pointer(error1)) = uintptr(0)
		}
		return int32(1) /* would not block */
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	ret = _poll_for_reply(tls, c, _widen(tls, c, request), reply, error1)
	if !(ret != 0) && (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading == 0 && X_xcb_in_read(tls, c) != 0 { /* _xcb_in_read shuts down the connection on error */
		ret = _poll_for_reply(tls, c, _widen(tls, c, request), reply, error1)
	}
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func Xxcb_poll_for_reply64(tls *libc.TLS, c uintptr, request Tuint64_t, reply uintptr, error1 uintptr) (r int32) {
	var ret int32
	_ = ret
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		*(*uintptr)(unsafe.Pointer(reply)) = uintptr(0)
		if error1 != 0 {
			*(*uintptr)(unsafe.Pointer(error1)) = uintptr(0)
		}
		return int32(1) /* would not block */
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	ret = _poll_for_reply(tls, c, request, reply, error1)
	if !(ret != 0) && (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading == 0 && X_xcb_in_read(tls, c) != 0 { /* _xcb_in_read shuts down the connection on error */
		ret = _poll_for_reply(tls, c, request, reply, error1)
	}
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func Xxcb_wait_for_event(tls *libc.TLS, c uintptr) (r uintptr) {
	var ret, v1 uintptr
	_, _ = ret, v1
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	/* get_event returns 0 on empty list. */
	for {
		v1 = _get_event(tls, c)
		ret = v1
		if !!(v1 != 0) {
			break
		}
		if !(X_xcb_conn_wait(tls, c, c+40, uintptr(0), uintptr(0)) != 0) {
			break
		}
	}
	X_xcb_in_wake_up_next_reader(tls, c)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return ret
}

func _poll_for_next_event(tls *libc.TLS, c uintptr, queued int32) (r uintptr) {
	var ret uintptr
	_ = ret
	ret = uintptr(0)
	if !((*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0) {
		libc.Xpthread_mutex_lock(tls, c+12)
		/* FIXME: follow X meets Z architecture changes. */
		ret = _get_event(tls, c)
		if !(ret != 0) && !(queued != 0) && (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading == 0 && X_xcb_in_read(tls, c) != 0 { /* _xcb_in_read shuts down the connection on error */
			ret = _get_event(tls, c)
		}
		libc.Xpthread_mutex_unlock(tls, c+12)
	}
	return ret
}

func Xxcb_poll_for_event(tls *libc.TLS, c uintptr) (r uintptr) {
	return _poll_for_next_event(tls, c, 0)
}

func Xxcb_poll_for_queued_event(tls *libc.TLS, c uintptr) (r uintptr) {
	return _poll_for_next_event(tls, c, int32(1))
}

func Xxcb_request_check(tls *libc.TLS, c uintptr, cookie Txcb_void_cookie_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var reply uintptr
	var request Tuint64_t
	var _ /* ret at bp+0 */ uintptr
	_, _ = reply, request
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	request = _widen(tls, c, cookie.Fsequence)
	if libc.Int64FromUint64(request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_completed) > 0 {
		if libc.Int64FromUint64(request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Frequest_expected) >= 0 {
			X_xcb_out_send_sync(tls, c)
		}
		if libc.Int64FromUint64(request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest_expected_written) >= 0 {
			X_xcb_out_flush_to(tls, c, (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest)
		}
	}
	reply = _wait_for_reply(tls, c, request, bp)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return *(*uintptr)(unsafe.Pointer(bp))
}

func _get_special_event(tls *libc.TLS, c uintptr, se uintptr) (r uintptr) {
	var event, events, v1, v2 uintptr
	_, _, _, _ = event, events, v1, v2
	event = libc.UintptrFromInt32(0)
	v1 = (*Txcb_special_event_t)(unsafe.Pointer(se)).Fevents
	events = v1
	if v1 != libc.UintptrFromInt32(0) {
		event = (*Tevent_list)(unsafe.Pointer(events)).Fevent
		v2 = (*Tevent_list)(unsafe.Pointer(events)).Fnext
		(*Txcb_special_event_t)(unsafe.Pointer(se)).Fevents = v2
		if !(v2 != 0) {
			(*Txcb_special_event_t)(unsafe.Pointer(se)).Fevents_tail = se + 16
		}
		libc.Xfree(tls, events)
	}
	return event
}

func Xxcb_poll_for_special_event(tls *libc.TLS, c uintptr, se uintptr) (r uintptr) {
	var event uintptr
	_ = event
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	event = _get_special_event(tls, c, se)
	if !(event != 0) && (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freading == 0 && X_xcb_in_read(tls, c) != 0 { /* _xcb_in_read shuts down the connection on error */
		event = _get_special_event(tls, c, se)
	}
	libc.Xpthread_mutex_unlock(tls, c+12)
	return event
}

func Xxcb_wait_for_special_event(tls *libc.TLS, c uintptr, se uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var event, v1 uintptr
	var _ /* special at bp+0 */ Tspecial_list
	_, _ = event, v1
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	_insert_special(tls, c+40+4208, bp, se)
	/* get_special_event returns 0 on empty list. */
	for {
		v1 = _get_special_event(tls, c, se)
		event = v1
		if !!(v1 != 0) {
			break
		}
		if !(X_xcb_conn_wait(tls, c, se+24, uintptr(0), uintptr(0)) != 0) {
			break
		}
	}
	_remove_special(tls, c+40+4208, bp)
	X_xcb_in_wake_up_next_reader(tls, c)
	libc.Xpthread_mutex_unlock(tls, c+12)
	return event
}

func Xxcb_register_for_special_xge(tls *libc.TLS, c uintptr, ext uintptr, eid Tuint32_t, stamp uintptr) (r uintptr) {
	var ext_reply, se uintptr
	_, _ = ext_reply, se
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return libc.UintptrFromInt32(0)
	}
	ext_reply = Xxcb_get_extension_data(tls, c, ext)
	if !(ext_reply != 0) {
		return libc.UintptrFromInt32(0)
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	se = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fspecial_events
	for {
		if !(se != 0) {
			break
		}
		if libc.Int32FromUint8((*Txcb_special_event_t)(unsafe.Pointer(se)).Fextension) == libc.Int32FromUint8((*Txcb_query_extension_reply_t)(unsafe.Pointer(ext_reply)).Fmajor_opcode) && (*Txcb_special_event_t)(unsafe.Pointer(se)).Feid == eid {
			libc.Xpthread_mutex_unlock(tls, c+12)
			return libc.UintptrFromInt32(0)
		}
		goto _1
	_1:
		;
		se = (*Txcb_special_event_t)(unsafe.Pointer(se)).Fnext
	}
	se = libc.Xcalloc(tls, uint32(1), uint32(72))
	if !(se != 0) {
		libc.Xpthread_mutex_unlock(tls, c+12)
		return libc.UintptrFromInt32(0)
	}
	(*Txcb_special_event_t)(unsafe.Pointer(se)).Fextension = (*Txcb_query_extension_reply_t)(unsafe.Pointer(ext_reply)).Fmajor_opcode
	(*Txcb_special_event_t)(unsafe.Pointer(se)).Feid = eid
	(*Txcb_special_event_t)(unsafe.Pointer(se)).Fevents = libc.UintptrFromInt32(0)
	(*Txcb_special_event_t)(unsafe.Pointer(se)).Fevents_tail = se + 16
	(*Txcb_special_event_t)(unsafe.Pointer(se)).Fstamp = stamp
	libc.Xpthread_cond_init(tls, se+24, uintptr(0))
	(*Txcb_special_event_t)(unsafe.Pointer(se)).Fnext = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fspecial_events
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fspecial_events = se
	libc.Xpthread_mutex_unlock(tls, c+12)
	return se
}

func Xxcb_unregister_for_special_event(tls *libc.TLS, c uintptr, se uintptr) {
	var events, next, prev, s, v2 uintptr
	_, _, _, _, _ = events, next, prev, s, v2
	if !(se != 0) {
		return
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return
	}
	libc.Xpthread_mutex_lock(tls, c+12)
	prev = c + 40 + 4292
	for {
		v2 = *(*uintptr)(unsafe.Pointer(prev))
		s = v2
		if !(v2 != libc.UintptrFromInt32(0)) {
			break
		}
		if s == se {
			*(*uintptr)(unsafe.Pointer(prev)) = (*Txcb_special_event_t)(unsafe.Pointer(se)).Fnext
			events = (*Txcb_special_event_t)(unsafe.Pointer(se)).Fevents
			for {
				if !(events != 0) {
					break
				}
				next = (*Tevent_list)(unsafe.Pointer(events)).Fnext
				libc.Xfree(tls, (*Tevent_list)(unsafe.Pointer(events)).Fevent)
				libc.Xfree(tls, events)
				goto _3
			_3:
				;
				events = next
			}
			libc.Xpthread_cond_destroy(tls, se+24)
			libc.Xfree(tls, se)
			break
		}
		goto _1
	_1:
		;
		prev = s
	}
	libc.Xpthread_mutex_unlock(tls, c+12)
}

/* Private interface */

func X_xcb_in_init(tls *libc.TLS, in uintptr) (r int32) {
	if libc.Xpthread_cond_init(tls, in, uintptr(0)) != 0 {
		return 0
	}
	(*T_xcb_in)(unsafe.Pointer(in)).Freading = 0
	(*T_xcb_in)(unsafe.Pointer(in)).Fqueue_len = 0
	(*T_xcb_in)(unsafe.Pointer(in)).Frequest_read = uint64(0)
	(*T_xcb_in)(unsafe.Pointer(in)).Frequest_completed = uint64(0)
	(*T_xcb_in)(unsafe.Pointer(in)).Freplies = X_xcb_map_new(tls)
	if !((*T_xcb_in)(unsafe.Pointer(in)).Freplies != 0) {
		return 0
	}
	(*T_xcb_in)(unsafe.Pointer(in)).Fcurrent_reply_tail = in + 4184
	(*T_xcb_in)(unsafe.Pointer(in)).Fevents_tail = in + 4196
	(*T_xcb_in)(unsafe.Pointer(in)).Fpending_replies_tail = in + 4212
	return int32(1)
}

func X_xcb_in_destroy(tls *libc.TLS, in uintptr) {
	var e, pend uintptr
	_, _ = e, pend
	libc.Xpthread_cond_destroy(tls, in)
	_free_reply_list(tls, (*T_xcb_in)(unsafe.Pointer(in)).Fcurrent_reply)
	X_xcb_map_delete(tls, (*T_xcb_in)(unsafe.Pointer(in)).Freplies, __ccgo_fp(_free_reply_list))
	for (*T_xcb_in)(unsafe.Pointer(in)).Fevents != 0 {
		e = (*T_xcb_in)(unsafe.Pointer(in)).Fevents
		(*T_xcb_in)(unsafe.Pointer(in)).Fevents = (*Tevent_list)(unsafe.Pointer(e)).Fnext
		libc.Xfree(tls, (*Tevent_list)(unsafe.Pointer(e)).Fevent)
		libc.Xfree(tls, e)
	}
	for (*T_xcb_in)(unsafe.Pointer(in)).Fpending_replies != 0 {
		pend = (*T_xcb_in)(unsafe.Pointer(in)).Fpending_replies
		(*T_xcb_in)(unsafe.Pointer(in)).Fpending_replies = (*Tpending_reply)(unsafe.Pointer(pend)).Fnext
		libc.Xfree(tls, pend)
	}
}

func X_xcb_in_wake_up_next_reader(tls *libc.TLS, c uintptr) {
	var pthreadret int32
	_ = pthreadret
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freaders != 0 {
		pthreadret = libc.Xpthread_cond_signal(tls, (*Treader_list1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Freaders)).Fdata)
	} else {
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fspecial_waiters != 0 {
			pthreadret = libc.Xpthread_cond_signal(tls, (*Tspecial_list1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fspecial_waiters)).Fse+24)
		} else {
			pthreadret = libc.Xpthread_cond_signal(tls, c+40)
		}
	}
}

func X_xcb_in_expect_reply(tls *libc.TLS, c uintptr, request Tuint64_t, workaround _workarounds, flags int32) (r int32) {
	var pend uintptr
	var v1 Tuint64_t
	_, _ = pend, v1
	pend = libc.Xmalloc(tls, uint32(28))
	if !(pend != 0) {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_MEM_INSUFFICIENT))
		return 0
	}
	v1 = request
	(*Tpending_reply)(unsafe.Pointer(pend)).Flast_request = v1
	(*Tpending_reply)(unsafe.Pointer(pend)).Ffirst_request = v1
	(*Tpending_reply)(unsafe.Pointer(pend)).Fworkaround = workaround
	(*Tpending_reply)(unsafe.Pointer(pend)).Fflags = flags
	(*Tpending_reply)(unsafe.Pointer(pend)).Fnext = uintptr(0)
	*(*uintptr)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail)) = pend
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail = pend + 24
	return int32(1)
}

func X_xcb_in_replies_done(tls *libc.TLS, c uintptr) {
	var pend, prev_next uintptr
	_, _ = pend, prev_next
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail != c+40+4212 {
		pend = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail - uintptr(Tsize_t(libc.UintptrFromInt32(0)+24))
		if (*Tpending_reply1)(unsafe.Pointer(pend)).Fworkaround == int32(_WORKAROUND_EXTERNAL_SOCKET_OWNER) {
			if libc.Int64FromUint64((*Tpending_reply1)(unsafe.Pointer(pend)).Ffirst_request-(*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest) <= 0 {
				(*Tpending_reply1)(unsafe.Pointer(pend)).Flast_request = (*Txcb_connection_t)(unsafe.Pointer(c)).Fout.Frequest
				(*Tpending_reply1)(unsafe.Pointer(pend)).Fworkaround = int32(_WORKAROUND_NONE)
			} else {
				/* The socket was taken, but no requests were actually sent
				 * so just discard the pending_reply that was created.
				 */
				prev_next = c + 40 + 4212
				for *(*uintptr)(unsafe.Pointer(prev_next)) != pend {
					prev_next = *(*uintptr)(unsafe.Pointer(prev_next)) + 24
				}
				*(*uintptr)(unsafe.Pointer(prev_next)) = libc.UintptrFromInt32(0)
				(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fpending_replies_tail = prev_next
				libc.Xfree(tls, pend)
			}
		}
	}
}

func X_xcb_in_read(tls *libc.TLS, c uintptr) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var hdr, v2, v3 uintptr
	var i, n, nfd int32
	var _ /* cmsgbuf at bp+8 */ struct {
		Fbuf         [0][76]int8
		Fcmsghdr     Tcmsghdr
		F__ccgo_pad2 [64]byte
	}
	var _ /* iov at bp+0 */ Tiovec
	var _ /* msg at bp+84 */ Tmsghdr
	_, _, _, _, _, _ = hdr, i, n, nfd, v2, v3
	*(*Tiovec)(unsafe.Pointer(bp)) = Tiovec{
		Fiov_base: c + 40 + 52 + uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len),
		Fiov_len:  uint32(4096) - libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len),
	}
	*(*Tmsghdr)(unsafe.Pointer(bp + 84)) = Tmsghdr{
		Fmsg_iov:        bp,
		Fmsg_iovlen:     int32(1),
		Fmsg_control:    bp + 8,
		Fmsg_controllen: (libc.Uint32FromInt64(4)*libc.Uint32FromInt32(libc.Int32FromInt32(m_XCB_MAX_PASS_FD)-(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd)+libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) & ^(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) + (libc.Uint32FromInt64(12)+libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) & ^(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)),
	}
	n = libc.Xrecvmsg(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd, bp+84, 0)
	/* Check for truncation errors. Only MSG_CTRUNC is
	 * probably possible here, which would indicate that
	 * the sender tried to transmit more than XCB_MAX_PASS_FD
	 * file descriptors.
	 */
	if (*(*Tmsghdr)(unsafe.Pointer(bp + 84))).Fmsg_flags&(libc.Int32FromInt32(m_MSG_TRUNC1)|libc.Int32FromInt32(m_MSG_CTRUNC1)) != 0 {
		X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_FDPASSING_FAILED))
		return 0
	}
	if n > 0 {
		if (*(*Tmsghdr)(unsafe.Pointer(bp + 84))).Fmsg_controllen >= uint32(12) {
			if (*Tmsghdr)(unsafe.Pointer(bp+84)).Fmsg_controllen >= uint32(12) {
				v2 = (*Tmsghdr)(unsafe.Pointer(bp + 84)).Fmsg_control
			} else {
				v2 = libc.UintptrFromInt32(0)
			}
			hdr = v2
			for {
				if !(hdr != 0) {
					break
				}
				if (*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_level == int32(m_SOL_SOCKET) && (*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_type == int32(m_SCM_RIGHTS) {
					nfd = libc.Int32FromUint32(((*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_len - ((libc.Uint32FromInt64(12)+libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) & ^(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)) + libc.Uint32FromInt32(libc.Int32FromInt32(0)))) / uint32(4))
					libc.Xmemcpy(tls, c+40+4220+uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd)*4, hdr+libc.UintptrFromInt32(1)*12, libc.Uint32FromInt32(nfd)*uint32(4))
					(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd += nfd
				}
				goto _1
			_1:
				;
				if (*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_len < uint32(12) || uint32((*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_len+libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1))&libc.Uint32FromInt32(^libc.Int32FromUint32(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1)))+uint32(12) >= libc.Uint32FromInt32(int32((*Tmsghdr)(unsafe.Pointer(bp+84)).Fmsg_control+uintptr((*Tmsghdr)(unsafe.Pointer(bp+84)).Fmsg_controllen))-int32(hdr)) {
					v3 = uintptr(0)
				} else {
					v3 = hdr + uintptr(uint32((*Tcmsghdr)(unsafe.Pointer(hdr)).Fcmsg_len+libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1))&libc.Uint32FromInt32(^libc.Int32FromUint32(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1))))
				}
				hdr = v3
			}
		}
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Ftotal_read += libc.Uint64FromInt32(n)
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len += n
	}
	for _read_packet(tls, c) != 0 {
		/* empty */
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd != 0 {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd -= (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fifd
		libc.Xmemmove(tls, c+40+4220, c+40+4220+uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fifd)*4, libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd)*uint32(4))
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fifd = 0
		/* If we have any left-over file descriptors after emptying
		 * the input buffer, then the server sent some that we weren't
		 * expecting.  Close them and mark the connection as broken;
		 */
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len == 0 && (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd != 0 {
			i = 0
			for {
				if !(i < (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fin_fd.Fnfd) {
					break
				}
				libc.Xclose(tls, *(*int32)(unsafe.Pointer(c + 40 + 4220 + uintptr(i)*4)))
				goto _4
			_4:
				;
				i++
			}
			X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_CLOSED_FDPASSING_FAILED))
			return 0
		}
	}
	if n > 0 || n < 0 && (*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EAGAIN) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EINTR)) {
		return int32(1)
	}
	X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_ERROR))
	return 0
}

func X_xcb_in_read_block(tls *libc.TLS, c uintptr, buf uintptr, len1 int32) (r int32) {
	var done, ret int32
	_, _ = done, ret
	done = (*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len
	if len1 < done {
		done = len1
	}
	libc.Xmemcpy(tls, buf, c+40+52, libc.Uint32FromInt32(done))
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len -= done
	libc.Xmemmove(tls, c+40+52, c+40+52+uintptr(done), libc.Uint32FromInt32((*Txcb_connection_t)(unsafe.Pointer(c)).Fin.Fqueue_len))
	if len1 > done {
		ret = _read_block(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Ffd, buf+uintptr(done), len1-done)
		if ret <= 0 {
			X_xcb_conn_shutdown(tls, c, int32(m_XCB_CONN_ERROR))
			return ret
		}
	}
	return len1
}

type Tpending_reply1 = struct {
	Ffirst_request Tuint64_t
	Flast_request  Tuint64_t
	Fworkaround    _workarounds
	Fflags         int32
	Fnext          uintptr
}

type Treader_list1 = struct {
	Frequest Tuint64_t
	Fdata    uintptr
	Fnext    uintptr
}

type Tspecial_list1 = struct {
	Fse   uintptr
	Fnext uintptr
}

type Txcb_ge_special_event_t1 = struct {
	Fresponse_type Tuint8_t
	Fextension     Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fevtype        Tuint16_t
	Fpad0          [2]Tuint8_t
	Feid           Tuint32_t
	Fpad1          [16]Tuint8_t
}

const m_INT32_MAX2 = 0x7fffffff
const m_UINT32_MAX2 = "0xffffffffu"

type Tlazyreply = struct {
	Ftag   _lazy_reply_tag
	Fvalue struct {
		Freply  [0]uintptr
		Fcookie Txcb_query_extension_cookie_t
	}
}

func _get_index(tls *libc.TLS, c uintptr, idx int32) (r uintptr) {
	var new_extensions uintptr
	var new_size int32
	_, _ = new_extensions, new_size
	if idx > (*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions_size {
		new_size = idx << int32(1)
		new_extensions = libc.Xrealloc(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions, uint32(8)*libc.Uint32FromInt32(new_size))
		if !(new_extensions != 0) {
			return uintptr(0)
		}
		libc.Xmemset(tls, new_extensions+uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions_size)*8, 0, uint32(8)*libc.Uint32FromInt32(new_size-(*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions_size))
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions = new_extensions
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions_size = new_size
	}
	return (*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions + uintptr(idx)*8 - uintptr(1)*8
}

func _get_lazyreply(tls *libc.TLS, c uintptr, ext uintptr) (r uintptr) {
	var data uintptr
	var v1 int32
	_, _ = data, v1
	libc.Xpthread_mutex_lock(tls, uintptr(unsafe.Pointer(&_global_lock)))
	if !((*Txcb_extension_t)(unsafe.Pointer(ext)).Fglobal_id != 0) {
		_next_global_id++
		v1 = _next_global_id
		(*Txcb_extension_t)(unsafe.Pointer(ext)).Fglobal_id = v1
	}
	libc.Xpthread_mutex_unlock(tls, uintptr(unsafe.Pointer(&_global_lock)))
	data = _get_index(tls, c, (*Txcb_extension_t)(unsafe.Pointer(ext)).Fglobal_id)
	if data != 0 && (*Tlazyreply)(unsafe.Pointer(data)).Ftag == int32(_LAZY_NONE) {
		/* cache miss: query the server */
		(*Tlazyreply)(unsafe.Pointer(data)).Ftag = int32(_LAZY_COOKIE)
		*(*Txcb_query_extension_cookie_t)(unsafe.Pointer(data + 4)) = Xxcb_query_extension(tls, c, uint16(libc.Xstrlen(tls, (*Txcb_extension_t)(unsafe.Pointer(ext)).Fname)), (*Txcb_extension_t)(unsafe.Pointer(ext)).Fname)
	}
	return data
}

var _global_lock = Tpthread_mutex_t{}

var _next_global_id int32

/* Public interface */

// C documentation
//
//	/* Do not free the returned xcb_query_extension_reply_t - on return, it's aliased
//	 * from the cache. */
func Xxcb_get_extension_data(tls *libc.TLS, c uintptr, ext uintptr) (r uintptr) {
	var data, v1 uintptr
	_, _ = data, v1
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return uintptr(0)
	}
	libc.Xpthread_mutex_lock(tls, c+20976)
	data = _get_lazyreply(tls, c, ext)
	if data != 0 && (*Tlazyreply)(unsafe.Pointer(data)).Ftag == int32(_LAZY_COOKIE) {
		(*Tlazyreply)(unsafe.Pointer(data)).Ftag = int32(_LAZY_FORCED)
		*(*uintptr)(unsafe.Pointer(data + 4)) = Xxcb_query_extension_reply(tls, c, *(*Txcb_query_extension_cookie_t)(unsafe.Pointer(data + 4)), uintptr(0))
	}
	libc.Xpthread_mutex_unlock(tls, c+20976)
	if data != 0 {
		v1 = *(*uintptr)(unsafe.Pointer(data + 4))
	} else {
		v1 = uintptr(0)
	}
	return v1
}

func Xxcb_prefetch_extension_data(tls *libc.TLS, c uintptr, ext uintptr) {
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return
	}
	libc.Xpthread_mutex_lock(tls, c+20976)
	_get_lazyreply(tls, c, ext)
	libc.Xpthread_mutex_unlock(tls, c+20976)
}

/* Private interface */

func X_xcb_ext_init(tls *libc.TLS, c uintptr) (r int32) {
	if libc.Xpthread_mutex_init(tls, c+20976, uintptr(0)) != 0 {
		return 0
	}
	return int32(1)
}

func X_xcb_ext_destroy(tls *libc.TLS, c uintptr) {
	var v1 int32
	var v2 uintptr
	_, _ = v1, v2
	libc.Xpthread_mutex_destroy(tls, c+20976)
	for {
		v2 = c + 20976 + 32
		v1 = *(*int32)(unsafe.Pointer(v2))
		*(*int32)(unsafe.Pointer(v2))--
		if !(v1 > 0) {
			break
		}
		if (*(*Tlazyreply1)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions + uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions_size)*8))).Ftag == int32(_LAZY_FORCED) {
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions + uintptr((*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions_size)*8 + 4)))
		}
	}
	libc.Xfree(tls, (*Txcb_connection_t)(unsafe.Pointer(c)).Fext.Fextensions)
}

type Tlazyreply1 = struct {
	Ftag   _lazy_reply_tag
	Fvalue struct {
		Freply  [0]uintptr
		Fcookie Txcb_query_extension_cookie_t
	}
}

const m_XCB_XCMISC_MAJOR_VERSION = 1
const m_XCB_XCMISC_MINOR_VERSION = 1
const m_XCB_XC_MISC_GET_VERSION = 0
const m_XCB_XC_MISC_GET_XID_LIST = 2
const m_XCB_XC_MISC_GET_XID_RANGE = 1

type Txcb_xc_misc_get_version_cookie_t = struct {
	Fsequence uint32
}

type Txcb_xc_misc_get_version_request_t = struct {
	Fmajor_opcode         Tuint8_t
	Fminor_opcode         Tuint8_t
	Flength               Tuint16_t
	Fclient_major_version Tuint16_t
	Fclient_minor_version Tuint16_t
}

type Txcb_xc_misc_get_version_reply_t = struct {
	Fresponse_type        Tuint8_t
	Fpad0                 Tuint8_t
	Fsequence             Tuint16_t
	Flength               Tuint32_t
	Fserver_major_version Tuint16_t
	Fserver_minor_version Tuint16_t
}

type Txcb_xc_misc_get_xid_range_cookie_t = struct {
	Fsequence uint32
}

type Txcb_xc_misc_get_xid_range_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fminor_opcode Tuint8_t
	Flength       Tuint16_t
}

type Txcb_xc_misc_get_xid_range_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fstart_id      Tuint32_t
	Fcount         Tuint32_t
}

type Txcb_xc_misc_get_xid_list_cookie_t = struct {
	Fsequence uint32
}

type Txcb_xc_misc_get_xid_list_request_t = struct {
	Fmajor_opcode Tuint8_t
	Fminor_opcode Tuint8_t
	Flength       Tuint16_t
	Fcount        Tuint32_t
}

type Txcb_xc_misc_get_xid_list_reply_t = struct {
	Fresponse_type Tuint8_t
	Fpad0          Tuint8_t
	Fsequence      Tuint16_t
	Flength        Tuint32_t
	Fids_len       Tuint32_t
	Fpad1          [20]Tuint8_t
}

/**
 * @}
 */

/* Public interface */

func Xxcb_generate_id(tls *libc.TLS, c uintptr) (r Tuint32_t) {
	var range1, xc_misc_reply uintptr
	var ret Tuint32_t
	_, _, _ = range1, ret, xc_misc_reply
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		return libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	}
	libc.Xpthread_mutex_lock(tls, c+21012)
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Flast >= (*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Fmax-(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Finc+uint32(1) {
		if (*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Flast == uint32(0) {
			/* finish setting up initial range */
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Fmax = (*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fresource_id_mask
		} else {
			/* check for extension */
			xc_misc_reply = Xxcb_get_extension_data(tls, c, uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)))
			if !(xc_misc_reply != 0) || !((*Txcb_query_extension_reply_t)(unsafe.Pointer(xc_misc_reply)).Fpresent != 0) {
				libc.Xpthread_mutex_unlock(tls, c+21012)
				return libc.Uint32FromInt32(-libc.Int32FromInt32(1))
			}
			/* get new range */
			range1 = Xxcb_xc_misc_get_xid_range_reply(tls, c, Xxcb_xc_misc_get_xid_range(tls, c), uintptr(0))
			/* XXX The latter disjunct is what the server returns
			   when it is out of XIDs.  Sweet. */
			if !(range1 != 0) || (*Txcb_xc_misc_get_xid_range_reply_t)(unsafe.Pointer(range1)).Fstart_id == uint32(0) && (*Txcb_xc_misc_get_xid_range_reply_t)(unsafe.Pointer(range1)).Fcount == uint32(1) {
				libc.Xpthread_mutex_unlock(tls, c+21012)
				return libc.Uint32FromInt32(-libc.Int32FromInt32(1))
			}
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Flast = (*Txcb_xc_misc_get_xid_range_reply_t)(unsafe.Pointer(range1)).Fstart_id
			(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Fmax = (*Txcb_xc_misc_get_xid_range_reply_t)(unsafe.Pointer(range1)).Fstart_id + ((*Txcb_xc_misc_get_xid_range_reply_t)(unsafe.Pointer(range1)).Fcount-uint32(1))*(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Finc
			libc.Xfree(tls, range1)
		}
	} else {
		(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Flast += (*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Finc
	}
	ret = (*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Flast | (*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Fbase
	libc.Xpthread_mutex_unlock(tls, c+21012)
	return ret
}

/* Private interface */

func X_xcb_xid_init(tls *libc.TLS, c uintptr) (r int32) {
	if libc.Xpthread_mutex_init(tls, c+21012, uintptr(0)) != 0 {
		return 0
	}
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Flast = uint32(0)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Fmax = uint32(0)
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Fbase = (*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fresource_id_base
	(*Txcb_connection_t)(unsafe.Pointer(c)).Fxid.Finc = (*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fresource_id_mask & -(*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Fresource_id_mask
	return int32(1)
}

func X_xcb_xid_destroy(tls *libc.TLS, c uintptr) {
	libc.Xpthread_mutex_destroy(tls, c+21012)
}

type T_xcb_map = struct {
	Fhead uintptr
	Ftail uintptr
}

type Tnode = struct {
	Fnext uintptr
	Fkey  Tuint64_t
	Fdata uintptr
}

type T_xcb_map1 = struct {
	Fhead uintptr
	Ftail uintptr
}

/* Private interface */

func X_xcb_map_new(tls *libc.TLS) (r uintptr) {
	var list uintptr
	_ = list
	list = libc.Xmalloc(tls, uint32(8))
	if !(list != 0) {
		return uintptr(0)
	}
	(*T_xcb_map)(unsafe.Pointer(list)).Fhead = uintptr(0)
	(*T_xcb_map)(unsafe.Pointer(list)).Ftail = list
	return list
}

func X_xcb_map_delete(tls *libc.TLS, list uintptr, do_free Txcb_list_free_func_t) {
	var cur uintptr
	_ = cur
	if !(list != 0) {
		return
	}
	for (*T_xcb_map)(unsafe.Pointer(list)).Fhead != 0 {
		cur = (*T_xcb_map)(unsafe.Pointer(list)).Fhead
		if do_free != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{do_free})))(tls, (*Tnode)(unsafe.Pointer(cur)).Fdata)
		}
		(*T_xcb_map)(unsafe.Pointer(list)).Fhead = (*Tnode)(unsafe.Pointer(cur)).Fnext
		libc.Xfree(tls, cur)
	}
	libc.Xfree(tls, list)
}

func X_xcb_map_put(tls *libc.TLS, list uintptr, key Tuint64_t, data uintptr) (r int32) {
	var cur uintptr
	_ = cur
	cur = libc.Xmalloc(tls, uint32(16))
	if !(cur != 0) {
		return 0
	}
	(*Tnode)(unsafe.Pointer(cur)).Fkey = key
	(*Tnode)(unsafe.Pointer(cur)).Fdata = data
	(*Tnode)(unsafe.Pointer(cur)).Fnext = uintptr(0)
	*(*uintptr)(unsafe.Pointer((*T_xcb_map)(unsafe.Pointer(list)).Ftail)) = cur
	(*T_xcb_map)(unsafe.Pointer(list)).Ftail = cur
	return int32(1)
}

func X_xcb_map_remove(tls *libc.TLS, list uintptr, key Tuint64_t) (r uintptr) {
	var cur, ret, tmp uintptr
	_, _, _ = cur, ret, tmp
	cur = list
	for {
		if !(*(*uintptr)(unsafe.Pointer(cur)) != 0) {
			break
		}
		if (*Tnode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(cur)))).Fkey == key {
			tmp = *(*uintptr)(unsafe.Pointer(cur))
			ret = (*Tnode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(cur)))).Fdata
			*(*uintptr)(unsafe.Pointer(cur)) = (*Tnode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(cur)))).Fnext
			if !(*(*uintptr)(unsafe.Pointer(cur)) != 0) {
				(*T_xcb_map)(unsafe.Pointer(list)).Ftail = cur
			}
			libc.Xfree(tls, tmp)
			return ret
		}
		goto _1
	_1:
		;
		cur = *(*uintptr)(unsafe.Pointer(cur))
	}
	return uintptr(0)
}

const m_AI_ADDRCONFIG = 0x20
const m_AI_ALL = 0x10
const m_AI_CANONNAME = 0x02
const m_AI_NUMERICHOST = 4
const m_AI_NUMERICSERV = 1024
const m_AI_PASSIVE = 0x01
const m_AI_V4MAPPED = 0x08
const m_HOST_NOT_FOUND = 1
const m_MSG_CTRUNC2 = 0x0008
const m_MSG_TRUNC2 = 0x0020
const m_NI_DGRAM = 0x10
const m_NI_MAXHOST = 255
const m_NI_MAXSERV = 32
const m_NI_NAMEREQD = 0x08
const m_NI_NOFQDN = 0x04
const m_NI_NUMERICHOST = 0x01
const m_NI_NUMERICSCOPE = 0x100
const m_NI_NUMERICSERV = 0x02
const m_NO_ADDRESS = "NO_DATA"
const m_NO_DATA = 4
const m_NO_RECOVERY = 3
const m_O_NONBLOCK1 = 04000
const m_SCM_RIGHTS1 = 0x01
const m_SOCK_CLOEXEC1 = 524288
const m_SOL_TCP = 6
const m_TCPI_OPT_ECN = 8
const m_TCPI_OPT_SACK = 2
const m_TCPI_OPT_TIMESTAMPS = 1
const m_TCPI_OPT_WSCALE = 4
const m_TCPOLEN_MAXSEG = 4
const m_TCPOLEN_SACK_PERMITTED = 2
const m_TCPOLEN_TIMESTAMP = 10
const m_TCPOLEN_WINDOW = 3
const m_TCPOPT_EOL = 0
const m_TCPOPT_MAXSEG = 2
const m_TCPOPT_NOP = 1
const m_TCPOPT_SACK = 5
const m_TCPOPT_SACK_PERMITTED = 4
const m_TCPOPT_TIMESTAMP = 8
const m_TCPOPT_WINDOW = 3
const m_TCP_CA_CWR = 2
const m_TCP_CA_Disorder = 1
const m_TCP_CA_Loss = 4
const m_TCP_CA_Open = 0
const m_TCP_CA_Recovery = 3
const m_TCP_CC_INFO = 26
const m_TCP_CLOSE = 7
const m_TCP_CLOSE_WAIT = 8
const m_TCP_CLOSING = 11
const m_TCP_CM_INQ = "TCP_INQ"
const m_TCP_CONGESTION = 13
const m_TCP_CORK = 3
const m_TCP_DEFER_ACCEPT = 9
const m_TCP_ESTABLISHED = 1
const m_TCP_FASTOPEN = 23
const m_TCP_FASTOPEN_CONNECT = 30
const m_TCP_FASTOPEN_KEY = 33
const m_TCP_FASTOPEN_NO_COOKIE = 34
const m_TCP_FIN_WAIT1 = 4
const m_TCP_FIN_WAIT2 = 5
const m_TCP_INFO = 11
const m_TCP_INQ = 36
const m_TCP_KEEPCNT = 6
const m_TCP_KEEPIDLE = 4
const m_TCP_KEEPINTVL = 5
const m_TCP_LAST_ACK = 9
const m_TCP_LINGER2 = 8
const m_TCP_LISTEN = 10
const m_TCP_MAXSEG = 2
const m_TCP_MD5SIG = 14
const m_TCP_MD5SIG_EXT = 32
const m_TCP_MD5SIG_FLAG_IFINDEX = 0x2
const m_TCP_MD5SIG_FLAG_PREFIX = 0x1
const m_TCP_MD5SIG_MAXKEYLEN = 80
const m_TCP_NODELAY = 1
const m_TCP_NOTSENT_LOWAT = 25
const m_TCP_QUEUE_SEQ = 21
const m_TCP_QUICKACK = 12
const m_TCP_RECEIVE_ZEROCOPY_FLAG_TLB_CLEAN_HINT = 0x1
const m_TCP_REPAIR = 19
const m_TCP_REPAIR_OFF = 0
const m_TCP_REPAIR_ON = 1
const m_TCP_REPAIR_OPTIONS = 22
const m_TCP_REPAIR_QUEUE = 20
const m_TCP_REPAIR_WINDOW = 29
const m_TCP_SAVED_SYN = 28
const m_TCP_SAVE_SYN = 27
const m_TCP_SYNCNT = 7
const m_TCP_SYN_RECV = 3
const m_TCP_SYN_SENT = 2
const m_TCP_THIN_DUPACK = 17
const m_TCP_THIN_LINEAR_TIMEOUTS = 16
const m_TCP_TIMESTAMP = 24
const m_TCP_TIME_WAIT = 6
const m_TCP_TX_DELAY = 37
const m_TCP_ULP = 31
const m_TCP_USER_TIMEOUT = 18
const m_TCP_WINDOW_CLAMP = 10
const m_TCP_ZEROCOPY_RECEIVE = 35
const m_TH_ACK = 0x10
const m_TH_FIN = 0x01
const m_TH_PUSH = 0x08
const m_TH_RST = 0x04
const m_TH_SYN = 0x02
const m_TH_URG = 0x20
const m_TRY_AGAIN = 2

type Tmax_align_t = struct {
	F__ll int64
	F__ld float64
}

type Tptrdiff_t = int32

type Tsockaddr_un = struct {
	Fsun_family Tsa_family_t
	Fsun_path   [108]int8
}

const _TCP_NLA_PAD = 0
const _TCP_NLA_BUSY = 1
const _TCP_NLA_RWND_LIMITED = 2
const _TCP_NLA_SNDBUF_LIMITED = 3
const _TCP_NLA_DATA_SEGS_OUT = 4
const _TCP_NLA_TOTAL_RETRANS = 5
const _TCP_NLA_PACING_RATE = 6
const _TCP_NLA_DELIVERY_RATE = 7
const _TCP_NLA_SND_CWND = 8
const _TCP_NLA_REORDERING = 9
const _TCP_NLA_MIN_RTT = 10
const _TCP_NLA_RECUR_RETRANS = 11
const _TCP_NLA_DELIVERY_RATE_APP_LMT = 12
const _TCP_NLA_SNDQ_SIZE = 13
const _TCP_NLA_CA_STATE = 14
const _TCP_NLA_SND_SSTHRESH = 15
const _TCP_NLA_DELIVERED = 16
const _TCP_NLA_DELIVERED_CE = 17
const _TCP_NLA_BYTES_SENT = 18
const _TCP_NLA_BYTES_RETRANS = 19
const _TCP_NLA_DSACK_DUPS = 20
const _TCP_NLA_REORD_SEEN = 21
const _TCP_NLA_SRTT = 22
const _TCP_NLA_TIMEOUT_REHASH = 23
const _TCP_NLA_BYTES_NOTSENT = 24
const _TCP_NLA_EDT = 25
const _TCP_NLA_TTL = 26

type Ttcp_seq = uint32

type Ttcphdr = struct {
	F__ccgo0_0 struct {
		F__ccgo1_0 [0]struct {
			Fth_sport Tuint16_t
			Fth_dport Tuint16_t
			Fth_seq   Tuint32_t
			Fth_ack   Tuint32_t
			F__ccgo12 uint8
			Fth_flags Tuint8_t
			Fth_win   Tuint16_t
			Fth_sum   Tuint16_t
			Fth_urp   Tuint16_t
		}
		F__ccgo0_0 struct {
			Fsource   Tuint16_t
			Fdest     Tuint16_t
			Fseq      Tuint32_t
			Fack_seq  Tuint32_t
			F__ccgo12 uint16
			Fwindow   Tuint16_t
			Fcheck    Tuint16_t
			Furg_ptr  Tuint16_t
		}
	}
}

type _tcp_fastopen_client_fail = int32

const _TFO_STATUS_UNSPEC = 0
const _TFO_COOKIE_UNAVAILABLE = 1
const _TFO_DATA_NOT_ACKED = 2
const _TFO_SYN_RETRANSMITTED = 3

type Ttcp_info = struct {
	Ftcpi_state           Tuint8_t
	Ftcpi_ca_state        Tuint8_t
	Ftcpi_retransmits     Tuint8_t
	Ftcpi_probes          Tuint8_t
	Ftcpi_backoff         Tuint8_t
	Ftcpi_options         Tuint8_t
	F__ccgo6              uint8
	F__ccgo7              uint8
	Ftcpi_rto             Tuint32_t
	Ftcpi_ato             Tuint32_t
	Ftcpi_snd_mss         Tuint32_t
	Ftcpi_rcv_mss         Tuint32_t
	Ftcpi_unacked         Tuint32_t
	Ftcpi_sacked          Tuint32_t
	Ftcpi_lost            Tuint32_t
	Ftcpi_retrans         Tuint32_t
	Ftcpi_fackets         Tuint32_t
	Ftcpi_last_data_sent  Tuint32_t
	Ftcpi_last_ack_sent   Tuint32_t
	Ftcpi_last_data_recv  Tuint32_t
	Ftcpi_last_ack_recv   Tuint32_t
	Ftcpi_pmtu            Tuint32_t
	Ftcpi_rcv_ssthresh    Tuint32_t
	Ftcpi_rtt             Tuint32_t
	Ftcpi_rttvar          Tuint32_t
	Ftcpi_snd_ssthresh    Tuint32_t
	Ftcpi_snd_cwnd        Tuint32_t
	Ftcpi_advmss          Tuint32_t
	Ftcpi_reordering      Tuint32_t
	Ftcpi_rcv_rtt         Tuint32_t
	Ftcpi_rcv_space       Tuint32_t
	Ftcpi_total_retrans   Tuint32_t
	Ftcpi_pacing_rate     Tuint64_t
	Ftcpi_max_pacing_rate Tuint64_t
	Ftcpi_bytes_acked     Tuint64_t
	Ftcpi_bytes_received  Tuint64_t
	Ftcpi_segs_out        Tuint32_t
	Ftcpi_segs_in         Tuint32_t
	Ftcpi_notsent_bytes   Tuint32_t
	Ftcpi_min_rtt         Tuint32_t
	Ftcpi_data_segs_in    Tuint32_t
	Ftcpi_data_segs_out   Tuint32_t
	Ftcpi_delivery_rate   Tuint64_t
	Ftcpi_busy_time       Tuint64_t
	Ftcpi_rwnd_limited    Tuint64_t
	Ftcpi_sndbuf_limited  Tuint64_t
	Ftcpi_delivered       Tuint32_t
	Ftcpi_delivered_ce    Tuint32_t
	Ftcpi_bytes_sent      Tuint64_t
	Ftcpi_bytes_retrans   Tuint64_t
	Ftcpi_dsack_dups      Tuint32_t
	Ftcpi_reord_seen      Tuint32_t
	Ftcpi_rcv_ooopack     Tuint32_t
	Ftcpi_snd_wnd         Tuint32_t
}

type Ttcp_md5sig = struct {
	Ftcpm_addr      Tsockaddr_storage
	Ftcpm_flags     Tuint8_t
	Ftcpm_prefixlen Tuint8_t
	Ftcpm_keylen    Tuint16_t
	Ftcpm_ifindex   int32
	Ftcpm_key       [80]Tuint8_t
}

type Ttcp_diag_md5sig = struct {
	Ftcpm_family    Tuint8_t
	Ftcpm_prefixlen Tuint8_t
	Ftcpm_keylen    Tuint16_t
	Ftcpm_addr      [4]Tuint32_t
	Ftcpm_key       [80]Tuint8_t
}

type Ttcp_repair_window = struct {
	Fsnd_wl1    Tuint32_t
	Fsnd_wnd    Tuint32_t
	Fmax_window Tuint32_t
	Frcv_wnd    Tuint32_t
	Frcv_wup    Tuint32_t
}

type Ttcp_zerocopy_receive = struct {
	Faddress         Tuint64_t
	Flength          Tuint32_t
	Frecv_skip_hint  Tuint32_t
	Finq             Tuint32_t
	Ferr             Tint32_t
	Fcopybuf_address Tuint64_t
	Fcopybuf_len     Tint32_t
	Fflags           Tuint32_t
	Fmsg_control     Tuint64_t
	Fmsg_controllen  Tuint64_t
	Fmsg_flags       Tuint32_t
	Freserved        Tuint32_t
}

type Taddrinfo = struct {
	Fai_flags     int32
	Fai_family    int32
	Fai_socktype  int32
	Fai_protocol  int32
	Fai_addrlen   Tsocklen_t
	Fai_addr      uintptr
	Fai_canonname uintptr
	Fai_next      uintptr
}

type Tnetent = struct {
	Fn_name     uintptr
	Fn_aliases  uintptr
	Fn_addrtype int32
	Fn_net      Tuint32_t
}

type Thostent = struct {
	Fh_name      uintptr
	Fh_aliases   uintptr
	Fh_addrtype  int32
	Fh_length    int32
	Fh_addr_list uintptr
}

type Tservent = struct {
	Fs_name    uintptr
	Fs_aliases uintptr
	Fs_port    int32
	Fs_proto   uintptr
}

type Tprotoent = struct {
	Fp_name    uintptr
	Fp_aliases uintptr
	Fp_proto   int32
}

func Xxcb_popcount(tls *libc.TLS, mask Tuint32_t) (r int32) {
	var y Tuint32_t
	_ = y
	y = mask >> libc.Int32FromInt32(1) & uint32(033333333333)
	y = mask - y - y>>libc.Int32FromInt32(1)&uint32(033333333333)
	return libc.Int32FromUint32((y + y>>libc.Int32FromInt32(3)) & uint32(030707070707) % uint32(077))
}

func Xxcb_sumof(tls *libc.TLS, list uintptr, len1 int32) (r int32) {
	var i, s int32
	_, _ = i, s
	s = 0
	i = 0
	for {
		if !(i < len1) {
			break
		}
		s += libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(list)))
		list++
		goto _1
	_1:
		;
		i++
	}
	return s
}

func __xcb_parse_display(tls *libc.TLS, name uintptr, host uintptr, protocol uintptr, displayp uintptr, screenp uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var colon, slash uintptr
	var display, len1, screen int32
	var _ /* dot at bp+0 */ uintptr
	var _ /* end at bp+4 */ uintptr
	_, _, _, _, _ = colon, display, len1, screen, slash
	if !(name != 0) || !(*(*int8)(unsafe.Pointer(name)) != 0) {
		name = libc.Xgetenv(tls, __ccgo_ts+4)
	}
	if !(name != 0) {
		return 0
	}
	slash = libc.Xstrrchr(tls, name, int32('/'))
	if slash != 0 {
		len1 = int32(slash) - int32(name)
		if protocol != 0 {
			*(*uintptr)(unsafe.Pointer(protocol)) = libc.Xmalloc(tls, libc.Uint32FromInt32(len1+int32(1)))
			if !(*(*uintptr)(unsafe.Pointer(protocol)) != 0) {
				return 0
			}
			libc.Xmemcpy(tls, *(*uintptr)(unsafe.Pointer(protocol)), name, libc.Uint32FromInt32(len1))
			*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(protocol)) + uintptr(len1))) = int8('\000')
		}
		name = slash + uintptr(1)
	} else {
		if protocol != 0 {
			*(*uintptr)(unsafe.Pointer(protocol)) = libc.UintptrFromInt32(0)
		}
	}
	colon = libc.Xstrrchr(tls, name, int32(':'))
	if !(colon != 0) {
		goto error_out
	}
	len1 = int32(colon) - int32(name)
	colon++
	display = libc.Int32FromUint32(libc.Xstrtoul(tls, colon, bp, int32(10)))
	if *(*uintptr)(unsafe.Pointer(bp)) == colon {
		goto error_out
	}
	if int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))))) == int32('\000') {
		screen = 0
	} else {
		if int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))))) != int32('.') {
			goto error_out
		}
		*(*uintptr)(unsafe.Pointer(bp))++
		screen = libc.Int32FromUint32(libc.Xstrtoul(tls, *(*uintptr)(unsafe.Pointer(bp)), bp+4, int32(10)))
		if *(*uintptr)(unsafe.Pointer(bp + 4)) == *(*uintptr)(unsafe.Pointer(bp)) || int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 4))))) != int32('\000') {
			goto error_out
		}
	}
	/* At this point, the display string is fully parsed and valid, but
	 * the caller's memory is untouched. */
	*(*uintptr)(unsafe.Pointer(host)) = libc.Xmalloc(tls, libc.Uint32FromInt32(len1+int32(1)))
	if !(*(*uintptr)(unsafe.Pointer(host)) != 0) {
		goto error_out
	}
	libc.Xmemcpy(tls, *(*uintptr)(unsafe.Pointer(host)), name, libc.Uint32FromInt32(len1))
	*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(host)) + uintptr(len1))) = int8('\000')
	*(*int32)(unsafe.Pointer(displayp)) = display
	if screenp != 0 {
		*(*int32)(unsafe.Pointer(screenp)) = screen
	}
	return int32(1)
	goto error_out
error_out:
	;
	if protocol != 0 {
		libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(protocol)))
		*(*uintptr)(unsafe.Pointer(protocol)) = libc.UintptrFromInt32(0)
	}
	return 0
}

func Xxcb_parse_display(tls *libc.TLS, name uintptr, host uintptr, displayp uintptr, screenp uintptr) (r int32) {
	return __xcb_parse_display(tls, name, host, libc.UintptrFromInt32(0), displayp, screenp)
}

func __xcb_open(tls *libc.TLS, host uintptr, protocol uintptr, display int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var actual_filelen, fd int32
	var base, file uintptr
	var filelen Tsize_t
	var port, port1 uint16
	var v1 uint32
	_, _, _, _, _, _, _, _ = actual_filelen, base, fd, file, filelen, port, port1, v1
	base = uintptr(unsafe.Pointer(&_unix_base))
	file = libc.UintptrFromInt32(0)
	/* If protocol or host is "unix", fall through to Unix socket code below */
	if (!(protocol != 0) || libc.Xstrcmp(tls, __ccgo_ts+12, protocol) != 0) && int32(*(*int8)(unsafe.Pointer(host))) != int32('\000') && libc.Xstrcmp(tls, __ccgo_ts+12, host) != 0 {
		/* display specifies TCP */
		port = libc.Uint16FromInt32(int32(m_X_TCP_PORT) + display)
		return __xcb_open_tcp(tls, host, protocol, port)
	}
	filelen = libc.Xstrlen(tls, base) + uint32(1) + libc.Uint32FromInt64(4)*libc.Uint32FromInt32(3) + uint32(1)
	file = libc.Xmalloc(tls, filelen)
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	/* display specifies Unix socket */
	actual_filelen = libc.X__builtin_snprintf(tls, file, filelen, __ccgo_ts+17, libc.VaList(bp+8, base, display))
	if actual_filelen < 0 {
		libc.Xfree(tls, file)
		return -int32(1)
	}
	/* snprintf may truncate the file */
	if libc.Uint32FromInt32(actual_filelen) < filelen-uint32(1) {
		v1 = libc.Uint32FromInt32(actual_filelen)
	} else {
		v1 = filelen - uint32(1)
	}
	filelen = v1
	fd = __xcb_open_abstract(tls, protocol, file, filelen)
	if fd >= 0 || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(m_ENOENT) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(m_ECONNREFUSED) {
		libc.Xfree(tls, file)
		return fd
	}
	fd = __xcb_open_unix(tls, protocol, file)
	libc.Xfree(tls, file)
	if fd < 0 && !(protocol != 0) && int32(*(*int8)(unsafe.Pointer(host))) == int32('\000') {
		port1 = libc.Uint16FromInt32(int32(m_X_TCP_PORT) + display)
		fd = __xcb_open_tcp(tls, host, protocol, port1)
	}
	return fd
	return -int32(1) /* if control reaches here then something has gone wrong */
}

var _unix_base = [17]int8{'/', 't', 'm', 'p', '/', '.', 'X', '1', '1', '-', 'u', 'n', 'i', 'x', '/', 'X'}

func __xcb_socket(tls *libc.TLS, family int32, type1 int32, proto int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var fd int32
	_ = fd
	fd = libc.Xsocket(tls, family, type1|int32(m_SOCK_CLOEXEC1), proto)
	if fd == -int32(1) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_EINVAL) {
		fd = libc.Xsocket(tls, family, type1, proto)
		if fd >= 0 {
			libc.Xfcntl(tls, fd, int32(m_F_SETFD), libc.VaList(bp+8, int32(m_FD_CLOEXEC)))
		}
	}
	return fd
}

func __xcb_do_connect(tls *libc.TLS, fd int32, addr uintptr, addrlen int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* on at bp+0 */ int32
	*(*int32)(unsafe.Pointer(bp)) = int32(1)
	if fd < 0 {
		return -int32(1)
	}
	libc.Xsetsockopt(tls, fd, int32(m_IPPROTO_TCP), int32(m_TCP_NODELAY), bp, uint32(4))
	libc.Xsetsockopt(tls, fd, int32(m_SOL_SOCKET), int32(m_SO_KEEPALIVE), bp, uint32(4))
	return libc.Xconnect(tls, fd, addr, libc.Uint32FromInt32(addrlen))
}

func __xcb_open_tcp(tls *libc.TLS, host uintptr, protocol uintptr, port uint16) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var addr, bracket, v1 uintptr
	var fd int32
	var v2 bool
	var _ /* hints at bp+0 */ Taddrinfo
	var _ /* results at bp+44 */ uintptr
	var _ /* service at bp+36 */ [6]int8
	_, _, _, _, _ = addr, bracket, fd, v1, v2
	fd = -int32(1)
	if protocol != 0 && libc.Xstrcmp(tls, __ccgo_ts+22, protocol) != 0 && libc.Xstrcmp(tls, __ccgo_ts+26, protocol) != 0 && libc.Xstrcmp(tls, __ccgo_ts+31, protocol) != 0 {
		return -int32(1)
	}
	if int32(*(*int8)(unsafe.Pointer(host))) == int32('\000') {
		host = __ccgo_ts + 37
	}
	libc.Xmemset(tls, bp, 0, uint32(32))
	(*(*Taddrinfo)(unsafe.Pointer(bp))).Fai_flags |= int32(m_AI_NUMERICSERV)
	(*(*Taddrinfo)(unsafe.Pointer(bp))).Fai_family = m_PF_UNSPEC
	(*(*Taddrinfo)(unsafe.Pointer(bp))).Fai_socktype = int32(m_SOCK_STREAM)
	/* Allow IPv6 addresses enclosed in brackets. */
	if v2 = int32(*(*int8)(unsafe.Pointer(host))) == int32('['); v2 {
		v1 = libc.Xstrrchr(tls, host, int32(']'))
		bracket = v1
	}
	if v2 && v1 != 0 && int32(*(*int8)(unsafe.Pointer(bracket + 1))) == int32('\000') {
		*(*int8)(unsafe.Pointer(bracket)) = int8('\000')
		host++
		(*(*Taddrinfo)(unsafe.Pointer(bp))).Fai_flags |= int32(m_AI_NUMERICHOST)
		(*(*Taddrinfo)(unsafe.Pointer(bp))).Fai_family = int32(m_PF_INET6)
	}
	libc.X__builtin_snprintf(tls, bp+36, uint32(6), __ccgo_ts+47, libc.VaList(bp+56, libc.Int32FromUint16(port)))
	if libc.Xgetaddrinfo(tls, host, bp+36, bp, bp+44) != 0 {
		/* FIXME: use gai_strerror, and fill in error connection */
		return -int32(1)
	}
	addr = *(*uintptr)(unsafe.Pointer(bp + 44))
	for {
		if !(addr != 0) {
			break
		}
		fd = __xcb_socket(tls, (*Taddrinfo)(unsafe.Pointer(addr)).Fai_family, (*Taddrinfo)(unsafe.Pointer(addr)).Fai_socktype, (*Taddrinfo)(unsafe.Pointer(addr)).Fai_protocol)
		if __xcb_do_connect(tls, fd, (*Taddrinfo)(unsafe.Pointer(addr)).Fai_addr, libc.Int32FromUint32((*Taddrinfo)(unsafe.Pointer(addr)).Fai_addrlen)) >= 0 {
			break
		}
		libc.Xclose(tls, fd)
		fd = -int32(1)
		goto _3
	_3:
		;
		addr = (*Taddrinfo)(unsafe.Pointer(addr)).Fai_next
	}
	libc.Xfreeaddrinfo(tls, *(*uintptr)(unsafe.Pointer(bp + 44)))
	return fd
}

func __xcb_open_unix(tls *libc.TLS, protocol uintptr, file uintptr) (r int32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var fd int32
	var _ /* addr at bp+0 */ Tsockaddr_un
	var _ /* len at bp+112 */ Tsocklen_t
	var _ /* val at bp+116 */ int32
	_ = fd
	*(*Tsocklen_t)(unsafe.Pointer(bp + 112)) = uint32(4)
	if protocol != 0 && libc.Xstrcmp(tls, __ccgo_ts+12, protocol) != 0 {
		return -int32(1)
	}
	libc.Xstrcpy(tls, bp+2, file)
	(*(*Tsockaddr_un)(unsafe.Pointer(bp))).Fsun_family = uint16(m_PF_LOCAL)
	fd = __xcb_socket(tls, int32(m_PF_LOCAL), int32(m_SOCK_STREAM), 0)
	if fd == -int32(1) {
		return -int32(1)
	}
	if libc.Xgetsockopt(tls, fd, int32(m_SOL_SOCKET), int32(m_SO_SNDBUF), bp+116, bp+112) == 0 && *(*int32)(unsafe.Pointer(bp + 116)) < libc.Int32FromInt32(64)*libc.Int32FromInt32(1024) {
		*(*int32)(unsafe.Pointer(bp + 116)) = libc.Int32FromInt32(64) * libc.Int32FromInt32(1024)
		libc.Xsetsockopt(tls, fd, int32(m_SOL_SOCKET), int32(m_SO_SNDBUF), bp+116, uint32(4))
	}
	if libc.Xconnect(tls, fd, bp, uint32(110)) == -int32(1) {
		libc.Xclose(tls, fd)
		return -int32(1)
	}
	return fd
}

func __xcb_open_abstract(tls *libc.TLS, protocol uintptr, file uintptr, filelen Tsize_t) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var fd int32
	var namelen Tsocklen_t
	var _ /* addr at bp+0 */ Tsockaddr_un
	_, _ = fd, namelen
	*(*Tsockaddr_un)(unsafe.Pointer(bp)) = Tsockaddr_un{}
	if protocol != 0 && libc.Xstrcmp(tls, __ccgo_ts+12, protocol) != 0 {
		return -int32(1)
	}
	libc.Xstrcpy(tls, bp+2+uintptr(1), file)
	(*(*Tsockaddr_un)(unsafe.Pointer(bp))).Fsun_family = uint16(m_PF_LOCAL)
	namelen = uint32(libc.UintptrFromInt32(0)+2) + libc.Uint32FromInt32(1) + filelen
	fd = __xcb_socket(tls, int32(m_PF_LOCAL), int32(m_SOCK_STREAM), 0)
	if fd == -int32(1) {
		return -int32(1)
	}
	if libc.Xconnect(tls, fd, bp, namelen) == -int32(1) {
		libc.Xclose(tls, fd)
		return -int32(1)
	}
	return fd
}

func Xxcb_connect(tls *libc.TLS, displayname uintptr, screenp uintptr) (r uintptr) {
	return Xxcb_connect_to_display_with_auth_info(tls, displayname, libc.UintptrFromInt32(0), screenp)
}

func Xxcb_connect_to_display_with_auth_info(tls *libc.TLS, displayname uintptr, auth uintptr, screenp uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c uintptr
	var fd, parsed int32
	var _ /* display at bp+0 */ int32
	var _ /* host at bp+4 */ uintptr
	var _ /* ourauth at bp+12 */ Txcb_auth_info_t
	var _ /* protocol at bp+8 */ uintptr
	_, _, _ = c, fd, parsed
	*(*int32)(unsafe.Pointer(bp)) = 0
	*(*uintptr)(unsafe.Pointer(bp + 4)) = libc.UintptrFromInt32(0)
	*(*uintptr)(unsafe.Pointer(bp + 8)) = libc.UintptrFromInt32(0)
	parsed = __xcb_parse_display(tls, displayname, bp+4, bp+8, bp, screenp)
	if !(parsed != 0) {
		c = X_xcb_conn_ret_error(tls, int32(m_XCB_CONN_CLOSED_PARSE_ERR))
		goto out
	}
	fd = __xcb_open(tls, *(*uintptr)(unsafe.Pointer(bp + 4)), *(*uintptr)(unsafe.Pointer(bp + 8)), *(*int32)(unsafe.Pointer(bp)))
	if fd == -int32(1) {
		c = X_xcb_conn_ret_error(tls, int32(m_XCB_CONN_ERROR))
		goto out
	}
	if auth != 0 {
		c = Xxcb_connect_to_fd(tls, fd, auth)
		goto out
	}
	if X_xcb_get_auth_info(tls, fd, bp+12, *(*int32)(unsafe.Pointer(bp))) != 0 {
		c = Xxcb_connect_to_fd(tls, fd, bp+12)
		libc.Xfree(tls, (*(*Txcb_auth_info_t)(unsafe.Pointer(bp + 12))).Fname)
		libc.Xfree(tls, (*(*Txcb_auth_info_t)(unsafe.Pointer(bp + 12))).Fdata)
	} else {
		c = Xxcb_connect_to_fd(tls, fd, uintptr(0))
	}
	if (*Txcb_connection_t)(unsafe.Pointer(c)).Fhas_error != 0 {
		goto out
	}
	/* Make sure requested screen number is in bounds for this server */
	if screenp != libc.UintptrFromInt32(0) && *(*int32)(unsafe.Pointer(screenp)) >= libc.Int32FromUint8((*Txcb_setup_t)(unsafe.Pointer((*Txcb_connection_t)(unsafe.Pointer(c)).Fsetup)).Froots_len) {
		Xxcb_disconnect(tls, c)
		c = X_xcb_conn_ret_error(tls, int32(m_XCB_CONN_CLOSED_INVALID_SCREEN))
		goto out
	}
	goto out
out:
	;
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 4)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
	return c
}

const m_AUTH_PROTO_MIT_MAGIC_COOKIE = "MIT-MAGIC-COOKIE-1"
const m_AUTH_PROTO_XDM_AUTHORIZATION = "XDM-AUTHORIZATION-1"
const m_CANBSIZ = 255
const m_DEV_BSIZE = 512
const m_FALSE = 0
const m_FUNCPROTO = 15
const m_FamilyKrb5Principal = 253
const m_FamilyLocal = 256
const m_FamilyLocalHost = 252
const m_FamilyNetname = 254
const m_FamilyWild = 65535
const m_INITIAL_SOCKNAME_SLACK = 108
const m_ITIMER_PROF = 2
const m_ITIMER_REAL = 0
const m_ITIMER_VIRTUAL = 1
const m_LOCK_ERROR = 1
const m_LOCK_SUCCESS = 0
const m_LOCK_TIMEOUT = 2
const m_MAXHOSTNAMELEN = 64
const m_MAXNAMLEN = 255
const m_MAXPATHLEN = 4096
const m_MAXSYMLINKS = 20
const m_NBBY = 8
const m_NCARGS = 131072
const m_NGROUPS = 32
const m_NOFILE = 256
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 0
const m_PRIO_MAX = 20
const m_PRIO_PGRP = 1
const m_PRIO_PROCESS = 0
const m_PRIO_USER = 2
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
const m_SOCK_CLOEXEC2 = 02000000
const m_TRUE = 1
const m_XDM_DEFAULT_MCAST_ADDR6 = "ff02:0:0:0:0:0:0:12b"
const m_XDM_KA_RTX_LIMIT = 4
const m_XDM_MAX_MSGLEN = 8192
const m_XDM_MAX_RTX = 32
const m_XDM_MIN_RTX = 2
const m_XDM_PROTOCOL_VERSION = 1
const m_XDM_RTX_LIMIT = 7
const m_XDM_UDP_PORT = 177
const m_XMD_H = 1
const m__X_INLINE = "inline"
const m__X_RESTRICT_KYWD = "restrict"
const m__Xconst = "const"
const m_prlimit64 = "prlimit"

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

type Titimerval = struct {
	Fit_interval Ttimeval
	Fit_value    Ttimeval
}

type Ttimezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Trlim_t = uint64

type Trlimit = struct {
	Frlim_cur Trlim_t
	Frlim_max Trlim_t
}

type Trusage = struct {
	Fru_utime    Ttimeval
	Fru_stime    Ttimeval
	Fru_maxrss   int32
	Fru_ixrss    int32
	Fru_idrss    int32
	Fru_isrss    int32
	Fru_minflt   int32
	Fru_majflt   int32
	Fru_nswap    int32
	Fru_inblock  int32
	Fru_oublock  int32
	Fru_msgsnd   int32
	Fru_msgrcv   int32
	Fru_nsignals int32
	Fru_nvcsw    int32
	Fru_nivcsw   int32
	F__reserved  [16]int32
}

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

type _auth_protos = int32

const _AUTH_XA1 = 0
const _AUTH_MC1 = 1
const _N_AUTH_PROTOS = 2

var _authnames = [2]uintptr{
	0: __ccgo_ts + 51,
	1: __ccgo_ts + 71,
}

var _authnameslen = [2]int32{
	0: libc.Int32FromUint32(libc.Uint32FromInt64(20) - libc.Uint32FromInt32(1)),
	1: libc.Int32FromUint32(libc.Uint32FromInt64(19) - libc.Uint32FromInt32(1)),
}

func _memdup(tls *libc.TLS, dst uintptr, src uintptr, len1 Tsize_t) (r Tsize_t) {
	if len1 != 0 {
		*(*uintptr)(unsafe.Pointer(dst)) = libc.Xmalloc(tls, len1)
	} else {
		*(*uintptr)(unsafe.Pointer(dst)) = uintptr(0)
	}
	if !(*(*uintptr)(unsafe.Pointer(dst)) != 0) {
		return uint32(0)
	}
	libc.Xmemcpy(tls, *(*uintptr)(unsafe.Pointer(dst)), src, len1)
	return len1
}

func _authname_match(tls *libc.TLS, kind _auth_protos, name uintptr, namelen Tsize_t) (r int32) {
	if libc.Uint32FromInt32(_authnameslen[kind]) != namelen {
		return 0
	}
	if libc.Xmemcmp(tls, _authnames[kind], name, namelen) != 0 {
		return 0
	}
	return int32(1)
}

func _get_authptr(tls *libc.TLS, sockname uintptr, display int32) (r uintptr) {
	bp := tls.Alloc(320)
	defer tls.Free(320)
	var addr uintptr
	var addrlen, dispbuflen int32
	var family uint16
	var v1 uint32
	var _ /* dispbuf at bp+256 */ [40]int8
	var _ /* hostnamebuf at bp+0 */ [256]int8
	_, _, _, _, _ = addr, addrlen, dispbuflen, family, v1
	addr = uintptr(0)
	addrlen = 0
	family = libc.Uint16FromInt32(libc.Int32FromInt32(m_FamilyLocal)) /* 256 */
	switch libc.Int32FromUint16((*Tsockaddr)(unsafe.Pointer(sockname)).Fsa_family) {
	case int32(m_PF_INET6):
		addr = sockname + 8
		addrlen = int32(16)
		if !(*(*Tuint32_t)(unsafe.Pointer(sockname + 8)) == uint32(0) && *(*Tuint32_t)(unsafe.Pointer(sockname + 8 + 1*4)) == uint32(0) && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 8))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 9))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 10))) == int32(0xff) && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 11))) == int32(0xff)) {
			if !(*(*Tuint32_t)(unsafe.Pointer(sockname + 8)) == uint32(0) && *(*Tuint32_t)(unsafe.Pointer(sockname + 8 + 1*4)) == uint32(0) && *(*Tuint32_t)(unsafe.Pointer(sockname + 8 + 2*4)) == uint32(0) && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 12))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 13))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 14))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 15))) == int32(1)) {
				family = uint16(_XCB_FAMILY_INTERNET_6)
			}
			break
		}
		addr += uintptr(12)
		/* if v4-mapped, fall through. */
		fallthrough
	case int32(m_PF_INET):
		if !(addr != 0) {
			addr = sockname + 4
		}
		addrlen = int32(4)
		if *(*Tin_addr_t)(unsafe.Pointer(addr)) != libc.Xhtonl(tls, libc.Uint32FromInt32(0x7f000001)) {
			family = uint16(_XCB_FAMILY_INTERNET)
		}
	case int32(m_PF_LOCAL):
	default:
		return uintptr(0) /* cannot authenticate this family */
	}
	dispbuflen = libc.X__builtin_snprintf(tls, bp+256, uint32(40), __ccgo_ts+90, libc.VaList(bp+304, display))
	if dispbuflen < 0 {
		return uintptr(0)
	}
	/* snprintf may have truncate our text */
	if libc.Uint32FromInt32(dispbuflen) < libc.Uint32FromInt64(40)-libc.Uint32FromInt32(1) {
		v1 = libc.Uint32FromInt32(dispbuflen)
	} else {
		v1 = libc.Uint32FromInt64(40) - libc.Uint32FromInt32(1)
	}
	dispbuflen = libc.Int32FromUint32(v1)
	if libc.Int32FromUint16(family) == int32(m_FamilyLocal) {
		if libc.Xgethostname(tls, bp, uint32(256)) == -int32(1) {
			return uintptr(0)
		} /* do not know own hostname */
		addr = bp
		addrlen = libc.Int32FromUint32(libc.Xstrlen(tls, addr))
	}
	return libxau.XXauGetBestAuthByAddr(tls, family, libc.Uint16FromInt32(addrlen), addr, libc.Uint16FromInt32(dispbuflen), bp+256, int32(_N_AUTH_PROTOS), uintptr(unsafe.Pointer(&_authnames)), uintptr(unsafe.Pointer(&_authnameslen)))
}

func _next_nonce(tls *libc.TLS) (r int32) {
	var ret, v1 int32
	_, _ = ret, v1
	libc.Xpthread_mutex_lock(tls, uintptr(unsafe.Pointer(&_nonce_mutex)))
	v1 = _nonce
	_nonce++
	ret = v1
	libc.Xpthread_mutex_unlock(tls, uintptr(unsafe.Pointer(&_nonce_mutex)))
	return ret
}

var _nonce int32

var _nonce_mutex = Tpthread_mutex_t{}

func _do_append(tls *libc.TLS, buf uintptr, idxp uintptr, val uintptr, valsize Tsize_t) {
	var p1 uintptr
	_ = p1
	libc.Xmemcpy(tls, buf+uintptr(*(*int32)(unsafe.Pointer(idxp))), val, valsize)
	p1 = idxp
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + valsize)
}

func _compute_auth(tls *libc.TLS, info uintptr, authptr uintptr, sockname uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var si, si6 uintptr
	var v2 int32
	var _ /* fakeaddr at bp+12 */ Tuint32_t
	var _ /* fakeaddr at bp+4 */ Tuint32_t
	var _ /* fakeport at bp+16 */ Tuint16_t
	var _ /* fakeport at bp+8 */ Tuint16_t
	var _ /* j at bp+0 */ int32
	var _ /* now at bp+20 */ Tuint32_t
	_, _, _ = si, si6, v2
	if _authname_match(tls, int32(_AUTH_MC1), (*TXauth)(unsafe.Pointer(authptr)).Fname, uint32((*TXauth)(unsafe.Pointer(authptr)).Fname_length)) != 0 {
		(*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdatalen = libc.Int32FromUint32(_memdup(tls, info+12, (*TXauth)(unsafe.Pointer(authptr)).Fdata, uint32((*TXauth)(unsafe.Pointer(authptr)).Fdata_length)))
		if !((*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdatalen != 0) {
			return 0
		}
		return int32(1)
	}
	if _authname_match(tls, int32(_AUTH_XA1), (*TXauth)(unsafe.Pointer(authptr)).Fname, uint32((*TXauth)(unsafe.Pointer(authptr)).Fname_length)) != 0 {
		(*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(192)/libc.Int32FromInt32(8)))
		if !((*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata != 0) {
			return 0
		}
		*(*int32)(unsafe.Pointer(bp)) = 0
		for {
			if !(*(*int32)(unsafe.Pointer(bp)) < int32(8)) {
				break
			}
			*(*int8)(unsafe.Pointer((*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata + uintptr(*(*int32)(unsafe.Pointer(bp))))) = *(*int8)(unsafe.Pointer((*TXauth)(unsafe.Pointer(authptr)).Fdata + uintptr(*(*int32)(unsafe.Pointer(bp)))))
			goto _1
		_1:
			;
			*(*int32)(unsafe.Pointer(bp))++
		}
		switch libc.Int32FromUint16((*Tsockaddr)(unsafe.Pointer(sockname)).Fsa_family) {
		case int32(m_PF_INET):
			/*block*/
			si = sockname
			_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, si+4, uint32(4))
			_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, si+2, uint32(2))
		case int32(m_PF_INET6):
			/*block*/
			si6 = sockname
			if *(*Tuint32_t)(unsafe.Pointer(sockname + 8)) == uint32(0) && *(*Tuint32_t)(unsafe.Pointer(sockname + 8 + 1*4)) == uint32(0) && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 8))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 9))) == 0 && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 10))) == int32(0xff) && libc.Int32FromUint8(*(*Tuint8_t)(unsafe.Pointer(sockname + 8 + 11))) == int32(0xff) {
				_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, si6+8+12, uint32(4))
				_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, si6+2, uint32(2))
			} else {
				/* XDM-AUTHORIZATION-1 does not handle IPv6 correctly.  Do the
				   same thing Xlib does: use all zeroes for the 4-byte address
				   and 2-byte port number. */
				*(*Tuint32_t)(unsafe.Pointer(bp + 4)) = uint32(0)
				*(*Tuint16_t)(unsafe.Pointer(bp + 8)) = uint16(0)
				_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, bp+4, uint32(4))
				_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, bp+8, uint32(2))
			}
		case int32(m_PF_LOCAL):
			/*block*/
			*(*Tuint32_t)(unsafe.Pointer(bp + 12)) = libc.Xhtonl(tls, uint32(0xffffffff)-libc.Uint32FromInt32(_next_nonce(tls)))
			*(*Tuint16_t)(unsafe.Pointer(bp + 16)) = libc.Xhtons(tls, libc.Uint16FromInt32(libc.Xgetpid(tls)))
			_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, bp+12, uint32(4))
			_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, bp+16, uint32(2))
		default:
			libc.Xfree(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata)
			return 0 /* do not know how to build this */
		}
		*(*Tuint32_t)(unsafe.Pointer(bp + 20)) = libc.Xhtonl(tls, libc.Uint32FromInt64(libc.Xtime(tls, uintptr(0))))
		_do_append(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, bp, bp+20, uint32(4))
		for *(*int32)(unsafe.Pointer(bp)) < libc.Int32FromInt32(192)/libc.Int32FromInt32(8) {
			v2 = *(*int32)(unsafe.Pointer(bp))
			*(*int32)(unsafe.Pointer(bp))++
			*(*int8)(unsafe.Pointer((*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata + uintptr(v2))) = 0
		}
		(*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdatalen = *(*int32)(unsafe.Pointer(bp))
		libxdmcp.XXdmcpWrap(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, (*TXauth)(unsafe.Pointer(authptr)).Fdata+uintptr(8), (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdata, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fdatalen)
		return int32(1)
	}
	return 0 /* Unknown authorization type */
}

/* `sockaddr_un.sun_path' typical size usually ranges between 92 and 108 */

// C documentation
//
//	/* Return a dynamically allocated socket address structure according
//	   to the value returned by either getpeername() or getsockname()
//	   (according to POSIX, applications should not assume a particular
//	   length for `sockaddr_un.sun_path') */
func _get_peer_sock_name(tls *libc.TLS, socket_func uintptr, fd int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var new_sockname, sockname, v1 uintptr
	var socknamelen Tsocklen_t
	var _ /* actual_socknamelen at bp+0 */ Tsocklen_t
	_, _, _, _ = new_sockname, sockname, socknamelen, v1
	socknamelen = libc.Uint32FromInt64(16) + libc.Uint32FromInt32(m_INITIAL_SOCKNAME_SLACK)
	*(*Tsocklen_t)(unsafe.Pointer(bp)) = socknamelen
	sockname = libc.Xmalloc(tls, socknamelen)
	if sockname == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	/* Both getpeername() and getsockname() truncates sockname if
	   there is not enough space and set the required length in
	   actual_socknamelen */
	if (*(*func(*libc.TLS, int32, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{socket_func})))(tls, fd, sockname, bp) == -int32(1) {
		goto sock_or_realloc_error
	}
	if *(*Tsocklen_t)(unsafe.Pointer(bp)) > socknamelen {
		new_sockname = libc.UintptrFromInt32(0)
		socknamelen = *(*Tsocklen_t)(unsafe.Pointer(bp))
		v1 = libc.Xrealloc(tls, sockname, *(*Tsocklen_t)(unsafe.Pointer(bp)))
		new_sockname = v1
		if v1 == libc.UintptrFromInt32(0) {
			goto sock_or_realloc_error
		}
		sockname = new_sockname
		if (*(*func(*libc.TLS, int32, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{socket_func})))(tls, fd, sockname, bp) == -int32(1) || *(*Tsocklen_t)(unsafe.Pointer(bp)) > socknamelen {
			goto sock_or_realloc_error
		}
	}
	return sockname
	goto sock_or_realloc_error
sock_or_realloc_error:
	;
	libc.Xfree(tls, sockname)
	return libc.UintptrFromInt32(0)
}

func X_xcb_get_auth_info(tls *libc.TLS, fd int32, info uintptr, display int32) (r int32) {
	var authptr, sockname, v1, v2, v3 uintptr
	var gotsockname, ret int32
	_, _, _, _, _, _, _ = authptr, gotsockname, ret, sockname, v1, v2, v3
	/* code adapted from Xlib/ConnDis.c, xtrans/Xtranssocket.c,
	   xtrans/Xtransutils.c */
	sockname = libc.UintptrFromInt32(0)
	gotsockname = 0
	authptr = uintptr(0)
	ret = int32(1)
	/* Some systems like hpux or Hurd do not expose peer names
	 * for UNIX Domain Sockets, but this is irrelevant,
	 * since compute_auth() ignores the peer name in this
	 * case anyway.*/
	v1 = _get_peer_sock_name(tls, __ccgo_fp(libc.Xgetpeername), fd)
	sockname = v1
	if v1 == libc.UintptrFromInt32(0) {
		v2 = _get_peer_sock_name(tls, __ccgo_fp(libc.Xgetsockname), fd)
		sockname = v2
		if v2 == libc.UintptrFromInt32(0) {
			return 0
		} /* can only authenticate sockets */
		if libc.Int32FromUint16((*Tsockaddr)(unsafe.Pointer(sockname)).Fsa_family) != int32(m_PF_LOCAL) {
			libc.Xfree(tls, sockname)
			return 0 /* except for AF_UNIX, sockets should have peernames */
		}
		gotsockname = int32(1)
	}
	authptr = _get_authptr(tls, sockname, display)
	if authptr == uintptr(0) {
		libc.Xfree(tls, sockname)
		return 0 /* cannot find good auth data */
	}
	(*Txcb_auth_info_t)(unsafe.Pointer(info)).Fnamelen = libc.Int32FromUint32(_memdup(tls, info+4, (*TXauth)(unsafe.Pointer(authptr)).Fname, uint32((*TXauth)(unsafe.Pointer(authptr)).Fname_length)))
	if !((*Txcb_auth_info_t)(unsafe.Pointer(info)).Fnamelen != 0) {
		goto no_auth
	} /* out of memory */
	if !(gotsockname != 0) {
		libc.Xfree(tls, sockname)
		v3 = _get_peer_sock_name(tls, __ccgo_fp(libc.Xgetsockname), fd)
		sockname = v3
		if v3 == libc.UintptrFromInt32(0) {
			libc.Xfree(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fname)
			goto no_auth /* can only authenticate sockets */
		}
	}
	ret = _compute_auth(tls, info, authptr, sockname)
	if !(ret != 0) {
		libc.Xfree(tls, (*Txcb_auth_info_t)(unsafe.Pointer(info)).Fname)
		goto no_auth /* cannot build auth record */
	}
	libc.Xfree(tls, sockname)
	sockname = libc.UintptrFromInt32(0)
	libxau.XXauDisposeAuth(tls, authptr)
	return ret
	goto no_auth
no_auth:
	;
	libc.Xfree(tls, sockname)
	(*Txcb_auth_info_t)(unsafe.Pointer(info)).Fname = uintptr(0)
	(*Txcb_auth_info_t)(unsafe.Pointer(info)).Fnamelen = 0
	libxau.XXauDisposeAuth(tls, authptr)
	return 0
}

/*
 * This file generated automatically from xproto.xml by c_client.py.
 * Edit at your peril.
 */

/**
 * @defgroup XCB__API XCB  API
 * @brief  XCB Protocol Implementation.
 * @{
 **/

/**
 * @}
 */

func Xxcb_char2b_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_char2b_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_char2b_iterator_t)(unsafe.Pointer(i)).Fdata += 2
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(2))
}

func Xxcb_char2b_end(tls *libc.TLS, i Txcb_char2b_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*2
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_window_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_window_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_window_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_window_end(tls *libc.TLS, i Txcb_window_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_pixmap_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_pixmap_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_pixmap_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_pixmap_end(tls *libc.TLS, i Txcb_pixmap_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_cursor_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_cursor_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_cursor_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_cursor_end(tls *libc.TLS, i Txcb_cursor_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_font_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_font_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_font_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_font_end(tls *libc.TLS, i Txcb_font_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_gcontext_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_gcontext_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_gcontext_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_gcontext_end(tls *libc.TLS, i Txcb_gcontext_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_colormap_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_colormap_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_colormap_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_colormap_end(tls *libc.TLS, i Txcb_colormap_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_atom_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_atom_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_atom_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_atom_end(tls *libc.TLS, i Txcb_atom_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_drawable_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_drawable_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_drawable_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_drawable_end(tls *libc.TLS, i Txcb_drawable_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_fontable_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_fontable_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_fontable_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_fontable_end(tls *libc.TLS, i Txcb_fontable_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_bool32_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_bool32_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_bool32_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_bool32_end(tls *libc.TLS, i Txcb_bool32_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_visualid_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_visualid_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_visualid_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_visualid_end(tls *libc.TLS, i Txcb_visualid_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_timestamp_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_timestamp_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_timestamp_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_timestamp_end(tls *libc.TLS, i Txcb_timestamp_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_keysym_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_keysym_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_keysym_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_keysym_end(tls *libc.TLS, i Txcb_keysym_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_keycode_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_keycode_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_keycode_iterator_t)(unsafe.Pointer(i)).Fdata++
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(1))
}

func Xxcb_keycode_end(tls *libc.TLS, i Txcb_keycode_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_keycode32_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_keycode32_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_keycode32_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_keycode32_end(tls *libc.TLS, i Txcb_keycode32_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_button_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_button_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_button_iterator_t)(unsafe.Pointer(i)).Fdata++
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(1))
}

func Xxcb_button_end(tls *libc.TLS, i Txcb_button_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_point_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_point_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_point_iterator_t)(unsafe.Pointer(i)).Fdata += 4
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(4))
}

func Xxcb_point_end(tls *libc.TLS, i Txcb_point_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*4
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_rectangle_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_rectangle_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_rectangle_iterator_t)(unsafe.Pointer(i)).Fdata += 8
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(8))
}

func Xxcb_rectangle_end(tls *libc.TLS, i Txcb_rectangle_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*8
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_arc_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_arc_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_arc_iterator_t)(unsafe.Pointer(i)).Fdata += 12
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(12))
}

func Xxcb_arc_end(tls *libc.TLS, i Txcb_arc_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*12
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_format_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_format_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_format_iterator_t)(unsafe.Pointer(i)).Fdata += 8
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(8))
}

func Xxcb_format_end(tls *libc.TLS, i Txcb_format_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*8
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_visualtype_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_visualtype_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_visualtype_iterator_t)(unsafe.Pointer(i)).Fdata += 24
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(24))
}

func Xxcb_visualtype_end(tls *libc.TLS, i Txcb_visualtype_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*24
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_depth_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* visuals */
	xcb_block_len += uint32((*Txcb_depth_t)(unsafe.Pointer(_aux)).Fvisuals_len) * uint32(24)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_depth_visuals(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_depth_visuals_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_depth_t)(unsafe.Pointer(R)).Fvisuals_len)
}

func Xxcb_depth_visuals_iterator(tls *libc.TLS, R uintptr) (r Txcb_visualtype_iterator_t) {
	var i Txcb_visualtype_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8
	i.Frem = libc.Int32FromUint16((*Txcb_depth_t)(unsafe.Pointer(R)).Fvisuals_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_depth_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_depth_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_depth_sizeof(tls, R))
	(*Txcb_depth_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_depth_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_depth_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_depth_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_depth_end(tls *libc.TLS, _i Txcb_depth_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_depth_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_depth_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_depth_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_depth_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_depth_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_depth_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_screen_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp_len uint32
	_, _, _, _, _, _, _, _ = _aux, i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp, xcb_tmp_len
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(40)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* allowed_depths */
	i = uint32(0)
	for {
		if !(i < uint32((*Txcb_screen_t)(unsafe.Pointer(_aux)).Fallowed_depths_len)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_depth_sizeof(tls, xcb_tmp))
		xcb_block_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_screen_allowed_depths_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_screen_t)(unsafe.Pointer(R)).Fallowed_depths_len)
}

func Xxcb_screen_allowed_depths_iterator(tls *libc.TLS, R uintptr) (r Txcb_depth_iterator_t) {
	var i Txcb_depth_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*40
	i.Frem = libc.Int32FromUint8((*Txcb_screen_t)(unsafe.Pointer(R)).Fallowed_depths_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_screen_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_screen_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_screen_sizeof(tls, R))
	(*Txcb_screen_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_screen_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_screen_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_screen_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_screen_end(tls *libc.TLS, _i Txcb_screen_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_screen_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_screen_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_screen_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_screen_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_screen_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_screen_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_setup_request_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* authorization_protocol_name */
	xcb_block_len += uint32((*Txcb_setup_request_t)(unsafe.Pointer(_aux)).Fauthorization_protocol_name_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	xcb_align_to = uint32(4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* authorization_protocol_data */
	xcb_block_len += uint32((*Txcb_setup_request_t)(unsafe.Pointer(_aux)).Fauthorization_protocol_data_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	xcb_align_to = uint32(4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_setup_request_authorization_protocol_name(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_setup_request_authorization_protocol_name_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_setup_request_t)(unsafe.Pointer(R)).Fauthorization_protocol_name_len)
}

func Xxcb_setup_request_authorization_protocol_name_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12 + uintptr((*Txcb_setup_request_t)(unsafe.Pointer(R)).Fauthorization_protocol_name_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_request_authorization_protocol_data(tls *libc.TLS, R uintptr) (r uintptr) {
	var prev Txcb_generic_iterator_t
	_ = prev
	prev = Xxcb_setup_request_authorization_protocol_name_end(tls, R)
	return prev.Fdata + uintptr(-prev.Findex&(libc.Int32FromInt32(4)-libc.Int32FromInt32(1))) + libc.UintptrFromInt32(0)
}

func Xxcb_setup_request_authorization_protocol_data_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_setup_request_t)(unsafe.Pointer(R)).Fauthorization_protocol_data_len)
}

func Xxcb_setup_request_authorization_protocol_data_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i, prev Txcb_generic_iterator_t
	_, _ = i, prev
	prev = Xxcb_setup_request_authorization_protocol_name_end(tls, R)
	i.Fdata = prev.Fdata + uintptr(-prev.Findex&(libc.Int32FromInt32(4)-libc.Int32FromInt32(1))) + uintptr((*Txcb_setup_request_t)(unsafe.Pointer(R)).Fauthorization_protocol_data_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_request_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_setup_request_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_setup_request_sizeof(tls, R))
	(*Txcb_setup_request_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_setup_request_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_setup_request_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_setup_request_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_setup_request_end(tls *libc.TLS, _i Txcb_setup_request_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_setup_request_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_setup_request_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_setup_request_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_setup_request_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_setup_request_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_setup_request_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_setup_failed_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* reason */
	xcb_block_len += uint32((*Txcb_setup_failed_t)(unsafe.Pointer(_aux)).Freason_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_setup_failed_reason(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_setup_failed_reason_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_setup_failed_t)(unsafe.Pointer(R)).Freason_len)
}

func Xxcb_setup_failed_reason_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8 + uintptr((*Txcb_setup_failed_t)(unsafe.Pointer(R)).Freason_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_failed_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_setup_failed_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_setup_failed_sizeof(tls, R))
	(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_setup_failed_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_setup_failed_end(tls *libc.TLS, _i Txcb_setup_failed_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_setup_failed_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_setup_failed_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_setup_authenticate_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* reason */
	xcb_block_len += libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_setup_authenticate_t)(unsafe.Pointer(_aux)).Flength)*libc.Int32FromInt32(4)) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_setup_authenticate_reason(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_setup_authenticate_reason_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_setup_authenticate_t)(unsafe.Pointer(R)).Flength) * int32(4)
}

func Xxcb_setup_authenticate_reason_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8 + uintptr(libc.Int32FromUint16((*Txcb_setup_authenticate_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_authenticate_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_setup_authenticate_sizeof(tls, R))
	(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_setup_authenticate_end(tls *libc.TLS, _i Txcb_setup_authenticate_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_setup_authenticate_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_setup_authenticate_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_setup_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp_len uint32
	_, _, _, _, _, _, _, _ = _aux, i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp, xcb_tmp_len
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(40)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* vendor */
	xcb_block_len += uint32((*Txcb_setup_t)(unsafe.Pointer(_aux)).Fvendor_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	xcb_align_to = uint32(4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* pixmap_formats */
	xcb_block_len += uint32((*Txcb_setup_t)(unsafe.Pointer(_aux)).Fpixmap_formats_len) * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* roots */
	i = uint32(0)
	for {
		if !(i < uint32((*Txcb_setup_t)(unsafe.Pointer(_aux)).Froots_len)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_screen_sizeof(tls, xcb_tmp))
		xcb_block_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_setup_vendor(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*40
}

func Xxcb_setup_vendor_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_setup_t)(unsafe.Pointer(R)).Fvendor_len)
}

func Xxcb_setup_vendor_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*40 + uintptr((*Txcb_setup_t)(unsafe.Pointer(R)).Fvendor_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_pixmap_formats(tls *libc.TLS, R uintptr) (r uintptr) {
	var prev Txcb_generic_iterator_t
	_ = prev
	prev = Xxcb_setup_vendor_end(tls, R)
	return prev.Fdata + uintptr(-prev.Findex&(libc.Int32FromInt32(4)-libc.Int32FromInt32(1))) + libc.UintptrFromInt32(0)
}

func Xxcb_setup_pixmap_formats_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_setup_t)(unsafe.Pointer(R)).Fpixmap_formats_len)
}

func Xxcb_setup_pixmap_formats_iterator(tls *libc.TLS, R uintptr) (r Txcb_format_iterator_t) {
	var i Txcb_format_iterator_t
	var prev Txcb_generic_iterator_t
	_, _ = i, prev
	prev = Xxcb_setup_vendor_end(tls, R)
	i.Fdata = prev.Fdata + uintptr(-prev.Findex&(libc.Int32FromInt32(4)-libc.Int32FromInt32(1)))
	i.Frem = libc.Int32FromUint8((*Txcb_setup_t)(unsafe.Pointer(R)).Fpixmap_formats_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_roots_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_setup_t)(unsafe.Pointer(R)).Froots_len)
}

func Xxcb_setup_roots_iterator(tls *libc.TLS, R uintptr) (r Txcb_screen_iterator_t) {
	var i Txcb_screen_iterator_t
	var prev Txcb_generic_iterator_t
	_, _ = i, prev
	prev = Xxcb_format_end(tls, Xxcb_setup_pixmap_formats_iterator(tls, R))
	i.Fdata = prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&libc.Uint32FromInt32(3))
	i.Frem = libc.Int32FromUint8((*Txcb_setup_t)(unsafe.Pointer(R)).Froots_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_setup_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_setup_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_setup_sizeof(tls, R))
	(*Txcb_setup_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_setup_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_setup_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_setup_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_setup_end(tls *libc.TLS, _i Txcb_setup_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_setup_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_setup_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_setup_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_setup_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_setup_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_setup_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_client_message_data_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_client_message_data_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_client_message_data_iterator_t)(unsafe.Pointer(i)).Fdata += 20
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(20))
}

func Xxcb_client_message_data_end(tls *libc.TLS, i Txcb_client_message_data_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*20
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_create_window_value_list_serialize(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_parts_idx uint32
	var xcb_out, xcb_tmp uintptr
	var xcb_parts [16]Tiovec
	var _ /* xcb_pad0 at bp+0 */ [3]int8
	_, _, _, _, _, _, _, _, _, _ = i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_out, xcb_pad, xcb_padding_offset, xcb_parts, xcb_parts_idx, xcb_tmp
	xcb_out = *(*uintptr)(unsafe.Pointer(_buffer))
	xcb_buffer_len = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	xcb_pad = uint32(0)
	*(*[3]int8)(unsafe.Pointer(bp)) = [3]int8{}
	xcb_parts_idx = uint32(0)
	xcb_block_len = uint32(0)
	if value_mask&uint32(_XCB_CW_BACK_PIXMAP) != 0 {
		/* xcb_create_window_value_list_t.background_pixmap */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACK_PIXEL) != 0 {
		/* xcb_create_window_value_list_t.background_pixel */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 4
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXMAP) != 0 {
		/* xcb_create_window_value_list_t.border_pixmap */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 8
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXEL) != 0 {
		/* xcb_create_window_value_list_t.border_pixel */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 12
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BIT_GRAVITY) != 0 {
		/* xcb_create_window_value_list_t.bit_gravity */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 16
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_WIN_GRAVITY) != 0 {
		/* xcb_create_window_value_list_t.win_gravity */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 20
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_STORE) != 0 {
		/* xcb_create_window_value_list_t.backing_store */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 24
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PLANES) != 0 {
		/* xcb_create_window_value_list_t.backing_planes */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 28
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PIXEL) != 0 {
		/* xcb_create_window_value_list_t.backing_pixel */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 32
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_OVERRIDE_REDIRECT) != 0 {
		/* xcb_create_window_value_list_t.override_redirect */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 36
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_SAVE_UNDER) != 0 {
		/* xcb_create_window_value_list_t.save_under */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 40
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_EVENT_MASK) != 0 {
		/* xcb_create_window_value_list_t.event_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 44
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_DONT_PROPAGATE) != 0 {
		/* xcb_create_window_value_list_t.do_not_propogate_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 48
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_COLORMAP) != 0 {
		/* xcb_create_window_value_list_t.colormap */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 52
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_CURSOR) != 0 {
		/* xcb_create_window_value_list_t.cursor */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 56
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_parts[xcb_parts_idx].Fiov_base = bp
		xcb_parts[xcb_parts_idx].Fiov_len = xcb_pad
		xcb_parts_idx++
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.UintptrFromInt32(0) == xcb_out {
		/* allocate memory */
		xcb_out = libc.Xmalloc(tls, xcb_buffer_len)
		*(*uintptr)(unsafe.Pointer(_buffer)) = xcb_out
	}
	xcb_tmp = xcb_out
	i = uint32(0)
	for {
		if !(i < xcb_parts_idx) {
			break
		}
		if uintptr(0) != xcb_parts[i].Fiov_base && uint32(0) != xcb_parts[i].Fiov_len {
			libc.Xmemcpy(tls, xcb_tmp, xcb_parts[i].Fiov_base, xcb_parts[i].Fiov_len)
		}
		if uint32(0) != xcb_parts[i].Fiov_len {
			xcb_tmp += uintptr(xcb_parts[i].Fiov_len)
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_create_window_value_list_unpack(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset uint32
	var xcb_tmp uintptr
	_, _, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	if value_mask&uint32(_XCB_CW_BACK_PIXMAP) != 0 {
		/* xcb_create_window_value_list_t.background_pixmap */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fbackground_pixmap = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACK_PIXEL) != 0 {
		/* xcb_create_window_value_list_t.background_pixel */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fbackground_pixel = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXMAP) != 0 {
		/* xcb_create_window_value_list_t.border_pixmap */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fborder_pixmap = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXEL) != 0 {
		/* xcb_create_window_value_list_t.border_pixel */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fborder_pixel = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BIT_GRAVITY) != 0 {
		/* xcb_create_window_value_list_t.bit_gravity */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fbit_gravity = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_WIN_GRAVITY) != 0 {
		/* xcb_create_window_value_list_t.win_gravity */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fwin_gravity = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_STORE) != 0 {
		/* xcb_create_window_value_list_t.backing_store */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fbacking_store = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PLANES) != 0 {
		/* xcb_create_window_value_list_t.backing_planes */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fbacking_planes = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PIXEL) != 0 {
		/* xcb_create_window_value_list_t.backing_pixel */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fbacking_pixel = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_OVERRIDE_REDIRECT) != 0 {
		/* xcb_create_window_value_list_t.override_redirect */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Foverride_redirect = *(*Txcb_bool32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_SAVE_UNDER) != 0 {
		/* xcb_create_window_value_list_t.save_under */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fsave_under = *(*Txcb_bool32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_EVENT_MASK) != 0 {
		/* xcb_create_window_value_list_t.event_mask */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fevent_mask = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_DONT_PROPAGATE) != 0 {
		/* xcb_create_window_value_list_t.do_not_propogate_mask */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fdo_not_propogate_mask = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_COLORMAP) != 0 {
		/* xcb_create_window_value_list_t.colormap */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fcolormap = *(*Txcb_colormap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_CURSOR) != 0 {
		/* xcb_create_window_value_list_t.cursor */
		(*Txcb_create_window_value_list_t)(unsafe.Pointer(_aux)).Fcursor = *(*Txcb_cursor_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_create_window_value_list_sizeof(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* _aux at bp+0 */ Txcb_create_window_value_list_t
	return Xxcb_create_window_value_list_unpack(tls, _buffer, value_mask, bp)
}

func Xxcb_create_window_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value_list */
	xcb_block_len += libc.Uint32FromInt32(Xxcb_create_window_value_list_sizeof(tls, xcb_tmp, (*Txcb_create_window_request_t)(unsafe.Pointer(_aux)).Fvalue_mask))
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_create_window_checked(tls *libc.TLS, c uintptr, depth Tuint8_t, wid Txcb_window_t, parent Txcb_window_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t, border_width Tuint16_t, _class Tuint16_t, visual Txcb_visualid_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_create_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fdepth = depth
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwid = wid
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fparent = parent
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fx = x
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fy = y
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwidth = width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fheight = height
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fborder_width = border_width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).F_class = _class
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvisual = visual
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_window_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req)))
	return xcb_ret
}

var _xcb_req = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_create_window(tls *libc.TLS, c uintptr, depth Tuint8_t, wid Txcb_window_t, parent Txcb_window_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t, border_width Tuint16_t, _class Tuint16_t, visual Txcb_visualid_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_create_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fdepth = depth
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwid = wid
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fparent = parent
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fx = x
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fy = y
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwidth = width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fheight = height
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fborder_width = border_width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).F_class = _class
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvisual = visual
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_window_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req1)))
	return xcb_ret
}

var _xcb_req1 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_create_window_aux_checked(tls *libc.TLS, c uintptr, depth Tuint8_t, wid Txcb_window_t, parent Txcb_window_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t, border_width Tuint16_t, _class Tuint16_t, visual Txcb_visualid_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+72 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_create_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 72)) = uintptr(0)
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fdepth = depth
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwid = wid
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fparent = parent
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fx = x
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fy = y
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwidth = width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fheight = height
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fborder_width = border_width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).F_class = _class
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvisual = visual
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_window_value_list_serialize(tls, bp+72, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 72))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req2)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 72)))
	return xcb_ret
}

var _xcb_req2 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_create_window_aux(tls *libc.TLS, c uintptr, depth Tuint8_t, wid Txcb_window_t, parent Txcb_window_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t, border_width Tuint16_t, _class Tuint16_t, visual Txcb_visualid_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+72 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_create_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 72)) = uintptr(0)
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fdepth = depth
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwid = wid
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fparent = parent
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fx = x
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fy = y
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fwidth = width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fheight = height
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fborder_width = border_width
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).F_class = _class
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvisual = visual
	(*(*Txcb_create_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_window_value_list_serialize(tls, bp+72, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 72))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req3)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 72)))
	return xcb_ret
}

var _xcb_req3 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_create_window_value_list(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_change_window_attributes_value_list_serialize(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_parts_idx uint32
	var xcb_out, xcb_tmp uintptr
	var xcb_parts [16]Tiovec
	var _ /* xcb_pad0 at bp+0 */ [3]int8
	_, _, _, _, _, _, _, _, _, _ = i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_out, xcb_pad, xcb_padding_offset, xcb_parts, xcb_parts_idx, xcb_tmp
	xcb_out = *(*uintptr)(unsafe.Pointer(_buffer))
	xcb_buffer_len = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	xcb_pad = uint32(0)
	*(*[3]int8)(unsafe.Pointer(bp)) = [3]int8{}
	xcb_parts_idx = uint32(0)
	xcb_block_len = uint32(0)
	if value_mask&uint32(_XCB_CW_BACK_PIXMAP) != 0 {
		/* xcb_change_window_attributes_value_list_t.background_pixmap */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACK_PIXEL) != 0 {
		/* xcb_change_window_attributes_value_list_t.background_pixel */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 4
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXMAP) != 0 {
		/* xcb_change_window_attributes_value_list_t.border_pixmap */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 8
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXEL) != 0 {
		/* xcb_change_window_attributes_value_list_t.border_pixel */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 12
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BIT_GRAVITY) != 0 {
		/* xcb_change_window_attributes_value_list_t.bit_gravity */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 16
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_WIN_GRAVITY) != 0 {
		/* xcb_change_window_attributes_value_list_t.win_gravity */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 20
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_STORE) != 0 {
		/* xcb_change_window_attributes_value_list_t.backing_store */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 24
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PLANES) != 0 {
		/* xcb_change_window_attributes_value_list_t.backing_planes */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 28
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PIXEL) != 0 {
		/* xcb_change_window_attributes_value_list_t.backing_pixel */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 32
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_OVERRIDE_REDIRECT) != 0 {
		/* xcb_change_window_attributes_value_list_t.override_redirect */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 36
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_SAVE_UNDER) != 0 {
		/* xcb_change_window_attributes_value_list_t.save_under */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 40
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_EVENT_MASK) != 0 {
		/* xcb_change_window_attributes_value_list_t.event_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 44
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_DONT_PROPAGATE) != 0 {
		/* xcb_change_window_attributes_value_list_t.do_not_propogate_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 48
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_COLORMAP) != 0 {
		/* xcb_change_window_attributes_value_list_t.colormap */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 52
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_CURSOR) != 0 {
		/* xcb_change_window_attributes_value_list_t.cursor */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 56
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_parts[xcb_parts_idx].Fiov_base = bp
		xcb_parts[xcb_parts_idx].Fiov_len = xcb_pad
		xcb_parts_idx++
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.UintptrFromInt32(0) == xcb_out {
		/* allocate memory */
		xcb_out = libc.Xmalloc(tls, xcb_buffer_len)
		*(*uintptr)(unsafe.Pointer(_buffer)) = xcb_out
	}
	xcb_tmp = xcb_out
	i = uint32(0)
	for {
		if !(i < xcb_parts_idx) {
			break
		}
		if uintptr(0) != xcb_parts[i].Fiov_base && uint32(0) != xcb_parts[i].Fiov_len {
			libc.Xmemcpy(tls, xcb_tmp, xcb_parts[i].Fiov_base, xcb_parts[i].Fiov_len)
		}
		if uint32(0) != xcb_parts[i].Fiov_len {
			xcb_tmp += uintptr(xcb_parts[i].Fiov_len)
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_window_attributes_value_list_unpack(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset uint32
	var xcb_tmp uintptr
	_, _, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	if value_mask&uint32(_XCB_CW_BACK_PIXMAP) != 0 {
		/* xcb_change_window_attributes_value_list_t.background_pixmap */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fbackground_pixmap = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACK_PIXEL) != 0 {
		/* xcb_change_window_attributes_value_list_t.background_pixel */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fbackground_pixel = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXMAP) != 0 {
		/* xcb_change_window_attributes_value_list_t.border_pixmap */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fborder_pixmap = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BORDER_PIXEL) != 0 {
		/* xcb_change_window_attributes_value_list_t.border_pixel */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fborder_pixel = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BIT_GRAVITY) != 0 {
		/* xcb_change_window_attributes_value_list_t.bit_gravity */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fbit_gravity = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_WIN_GRAVITY) != 0 {
		/* xcb_change_window_attributes_value_list_t.win_gravity */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fwin_gravity = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_STORE) != 0 {
		/* xcb_change_window_attributes_value_list_t.backing_store */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fbacking_store = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PLANES) != 0 {
		/* xcb_change_window_attributes_value_list_t.backing_planes */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fbacking_planes = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_BACKING_PIXEL) != 0 {
		/* xcb_change_window_attributes_value_list_t.backing_pixel */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fbacking_pixel = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_OVERRIDE_REDIRECT) != 0 {
		/* xcb_change_window_attributes_value_list_t.override_redirect */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Foverride_redirect = *(*Txcb_bool32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_SAVE_UNDER) != 0 {
		/* xcb_change_window_attributes_value_list_t.save_under */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fsave_under = *(*Txcb_bool32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_EVENT_MASK) != 0 {
		/* xcb_change_window_attributes_value_list_t.event_mask */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fevent_mask = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_DONT_PROPAGATE) != 0 {
		/* xcb_change_window_attributes_value_list_t.do_not_propogate_mask */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fdo_not_propogate_mask = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_COLORMAP) != 0 {
		/* xcb_change_window_attributes_value_list_t.colormap */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fcolormap = *(*Txcb_colormap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_CW_CURSOR) != 0 {
		/* xcb_change_window_attributes_value_list_t.cursor */
		(*Txcb_change_window_attributes_value_list_t)(unsafe.Pointer(_aux)).Fcursor = *(*Txcb_cursor_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_window_attributes_value_list_sizeof(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* _aux at bp+0 */ Txcb_change_window_attributes_value_list_t
	return Xxcb_change_window_attributes_value_list_unpack(tls, _buffer, value_mask, bp)
}

func Xxcb_change_window_attributes_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value_list */
	xcb_block_len += libc.Uint32FromInt32(Xxcb_change_window_attributes_value_list_sizeof(tls, xcb_tmp, (*Txcb_change_window_attributes_request_t)(unsafe.Pointer(_aux)).Fvalue_mask))
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_window_attributes_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_change_window_attributes_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_window_attributes_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_window_attributes_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req4)))
	return xcb_ret
}

var _xcb_req4 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_WINDOW_ATTRIBUTES),
	Fisvoid: uint8(1),
}

func Xxcb_change_window_attributes(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_change_window_attributes_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_window_attributes_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_window_attributes_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req5)))
	return xcb_ret
}

var _xcb_req5 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_WINDOW_ATTRIBUTES),
	Fisvoid: uint8(1),
}

func Xxcb_change_window_attributes_aux_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+52 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_change_window_attributes_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 52)) = uintptr(0)
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_window_attributes_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_window_attributes_value_list_serialize(tls, bp+52, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 52))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req6)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 52)))
	return xcb_ret
}

var _xcb_req6 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_WINDOW_ATTRIBUTES),
	Fisvoid: uint8(1),
}

func Xxcb_change_window_attributes_aux(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+52 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_change_window_attributes_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 52)) = uintptr(0)
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_change_window_attributes_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_window_attributes_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_window_attributes_value_list_serialize(tls, bp+52, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 52))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req7)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 52)))
	return xcb_ret
}

var _xcb_req7 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_WINDOW_ATTRIBUTES),
	Fisvoid: uint8(1),
}

func Xxcb_change_window_attributes_value_list(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_get_window_attributes(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_get_window_attributes_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_window_attributes_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_window_attributes_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_window_attributes_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_window_attributes_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req8)))
	return xcb_ret
}

var _xcb_req8 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_WINDOW_ATTRIBUTES),
}

func Xxcb_get_window_attributes_unchecked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_get_window_attributes_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_window_attributes_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_window_attributes_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_window_attributes_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_window_attributes_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req9)))
	return xcb_ret
}

var _xcb_req9 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_WINDOW_ATTRIBUTES),
}

func Xxcb_get_window_attributes_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_window_attributes_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_destroy_window_checked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_destroy_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_destroy_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_destroy_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req10)))
	return xcb_ret
}

var _xcb_req10 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_DESTROY_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_destroy_window(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_destroy_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_destroy_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_destroy_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req11)))
	return xcb_ret
}

var _xcb_req11 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_DESTROY_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_destroy_subwindows_checked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_destroy_subwindows_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_destroy_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_destroy_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req12)))
	return xcb_ret
}

var _xcb_req12 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_DESTROY_SUBWINDOWS),
	Fisvoid: uint8(1),
}

func Xxcb_destroy_subwindows(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_destroy_subwindows_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_destroy_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_destroy_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req13)))
	return xcb_ret
}

var _xcb_req13 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_DESTROY_SUBWINDOWS),
	Fisvoid: uint8(1),
}

func Xxcb_change_save_set_checked(tls *libc.TLS, c uintptr, mode Tuint8_t, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_change_save_set_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_save_set_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*Txcb_change_save_set_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req14)))
	return xcb_ret
}

var _xcb_req14 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CHANGE_SAVE_SET),
	Fisvoid: uint8(1),
}

func Xxcb_change_save_set(tls *libc.TLS, c uintptr, mode Tuint8_t, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_change_save_set_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_save_set_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*Txcb_change_save_set_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req15)))
	return xcb_ret
}

var _xcb_req15 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CHANGE_SAVE_SET),
	Fisvoid: uint8(1),
}

func Xxcb_reparent_window_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, parent Txcb_window_t, x Tint16_t, y Tint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_reparent_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fparent = parent
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req16)))
	return xcb_ret
}

var _xcb_req16 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_REPARENT_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_reparent_window(tls *libc.TLS, c uintptr, window Txcb_window_t, parent Txcb_window_t, x Tint16_t, y Tint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_reparent_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fparent = parent
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_reparent_window_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req17)))
	return xcb_ret
}

var _xcb_req17 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_REPARENT_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_map_window_checked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_map_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_map_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_map_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req18)))
	return xcb_ret
}

var _xcb_req18 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_MAP_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_map_window(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_map_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_map_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_map_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req19)))
	return xcb_ret
}

var _xcb_req19 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_MAP_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_map_subwindows_checked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_map_subwindows_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_map_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_map_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req20)))
	return xcb_ret
}

var _xcb_req20 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_MAP_SUBWINDOWS),
	Fisvoid: uint8(1),
}

func Xxcb_map_subwindows(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_map_subwindows_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_map_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_map_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req21)))
	return xcb_ret
}

var _xcb_req21 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_MAP_SUBWINDOWS),
	Fisvoid: uint8(1),
}

func Xxcb_unmap_window_checked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_unmap_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_unmap_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_unmap_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req22)))
	return xcb_ret
}

var _xcb_req22 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNMAP_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_unmap_window(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_unmap_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_unmap_window_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_unmap_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req23)))
	return xcb_ret
}

var _xcb_req23 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNMAP_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_unmap_subwindows_checked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_unmap_subwindows_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_unmap_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_unmap_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req24)))
	return xcb_ret
}

var _xcb_req24 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNMAP_SUBWINDOWS),
	Fisvoid: uint8(1),
}

func Xxcb_unmap_subwindows(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_unmap_subwindows_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_unmap_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_unmap_subwindows_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req25)))
	return xcb_ret
}

var _xcb_req25 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNMAP_SUBWINDOWS),
	Fisvoid: uint8(1),
}

func Xxcb_configure_window_value_list_serialize(tls *libc.TLS, _buffer uintptr, value_mask Tuint16_t, _aux uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_parts_idx uint32
	var xcb_out, xcb_tmp uintptr
	var xcb_parts [8]Tiovec
	var _ /* xcb_pad0 at bp+0 */ [3]int8
	_, _, _, _, _, _, _, _, _, _ = i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_out, xcb_pad, xcb_padding_offset, xcb_parts, xcb_parts_idx, xcb_tmp
	xcb_out = *(*uintptr)(unsafe.Pointer(_buffer))
	xcb_buffer_len = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	xcb_pad = uint32(0)
	*(*[3]int8)(unsafe.Pointer(bp)) = [3]int8{}
	xcb_parts_idx = uint32(0)
	xcb_block_len = uint32(0)
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_X) != 0 {
		/* xcb_configure_window_value_list_t.x */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_Y) != 0 {
		/* xcb_configure_window_value_list_t.y */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 4
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_WIDTH) != 0 {
		/* xcb_configure_window_value_list_t.width */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 8
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_HEIGHT) != 0 {
		/* xcb_configure_window_value_list_t.height */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 12
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_BORDER_WIDTH) != 0 {
		/* xcb_configure_window_value_list_t.border_width */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 16
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_SIBLING) != 0 {
		/* xcb_configure_window_value_list_t.sibling */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 20
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_STACK_MODE) != 0 {
		/* xcb_configure_window_value_list_t.stack_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 24
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_parts[xcb_parts_idx].Fiov_base = bp
		xcb_parts[xcb_parts_idx].Fiov_len = xcb_pad
		xcb_parts_idx++
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.UintptrFromInt32(0) == xcb_out {
		/* allocate memory */
		xcb_out = libc.Xmalloc(tls, xcb_buffer_len)
		*(*uintptr)(unsafe.Pointer(_buffer)) = xcb_out
	}
	xcb_tmp = xcb_out
	i = uint32(0)
	for {
		if !(i < xcb_parts_idx) {
			break
		}
		if uintptr(0) != xcb_parts[i].Fiov_base && uint32(0) != xcb_parts[i].Fiov_len {
			libc.Xmemcpy(tls, xcb_tmp, xcb_parts[i].Fiov_base, xcb_parts[i].Fiov_len)
		}
		if uint32(0) != xcb_parts[i].Fiov_len {
			xcb_tmp += uintptr(xcb_parts[i].Fiov_len)
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_configure_window_value_list_unpack(tls *libc.TLS, _buffer uintptr, value_mask Tuint16_t, _aux uintptr) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset uint32
	var xcb_tmp uintptr
	_, _, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_X) != 0 {
		/* xcb_configure_window_value_list_t.x */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fx = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_Y) != 0 {
		/* xcb_configure_window_value_list_t.y */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fy = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_WIDTH) != 0 {
		/* xcb_configure_window_value_list_t.width */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fwidth = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_HEIGHT) != 0 {
		/* xcb_configure_window_value_list_t.height */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fheight = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_BORDER_WIDTH) != 0 {
		/* xcb_configure_window_value_list_t.border_width */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fborder_width = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_SIBLING) != 0 {
		/* xcb_configure_window_value_list_t.sibling */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fsibling = *(*Txcb_window_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if libc.Int32FromUint16(value_mask)&int32(_XCB_CONFIG_WINDOW_STACK_MODE) != 0 {
		/* xcb_configure_window_value_list_t.stack_mode */
		(*Txcb_configure_window_value_list_t)(unsafe.Pointer(_aux)).Fstack_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_configure_window_value_list_sizeof(tls *libc.TLS, _buffer uintptr, value_mask Tuint16_t) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* _aux at bp+0 */ Txcb_configure_window_value_list_t
	return Xxcb_configure_window_value_list_unpack(tls, _buffer, value_mask, bp)
}

func Xxcb_configure_window_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value_list */
	xcb_block_len += libc.Uint32FromInt32(Xxcb_configure_window_value_list_sizeof(tls, xcb_tmp, (*Txcb_configure_window_request_t)(unsafe.Pointer(_aux)).Fvalue_mask))
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_configure_window_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint16_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_configure_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	libc.Xmemset(tls, bp+40+10, 0, uint32(2))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_configure_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_configure_window_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req26)))
	return xcb_ret
}

var _xcb_req26 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CONFIGURE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_configure_window(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint16_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_configure_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	libc.Xmemset(tls, bp+40+10, 0, uint32(2))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_configure_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_configure_window_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req27)))
	return xcb_ret
}

var _xcb_req27 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CONFIGURE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_configure_window_aux_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint16_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+52 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_configure_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 52)) = uintptr(0)
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	libc.Xmemset(tls, bp+40+10, 0, uint32(2))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_configure_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_configure_window_value_list_serialize(tls, bp+52, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 52))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req28)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 52)))
	return xcb_ret
}

var _xcb_req28 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CONFIGURE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_configure_window_aux(tls *libc.TLS, c uintptr, window Txcb_window_t, value_mask Tuint16_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+52 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_configure_window_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 52)) = uintptr(0)
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fwindow = window
	(*(*Txcb_configure_window_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	libc.Xmemset(tls, bp+40+10, 0, uint32(2))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_configure_window_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_configure_window_value_list_serialize(tls, bp+52, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 52))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req29)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 52)))
	return xcb_ret
}

var _xcb_req29 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CONFIGURE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_configure_window_value_list(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_circulate_window_checked(tls *libc.TLS, c uintptr, direction Tuint8_t, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_circulate_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_circulate_window_request_t)(unsafe.Pointer(bp + 32))).Fdirection = direction
	(*(*Txcb_circulate_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req30)))
	return xcb_ret
}

var _xcb_req30 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CIRCULATE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_circulate_window(tls *libc.TLS, c uintptr, direction Tuint8_t, window Txcb_window_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_circulate_window_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_circulate_window_request_t)(unsafe.Pointer(bp + 32))).Fdirection = direction
	(*(*Txcb_circulate_window_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req31)))
	return xcb_ret
}

var _xcb_req31 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CIRCULATE_WINDOW),
	Fisvoid: uint8(1),
}

func Xxcb_get_geometry(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t) (r Txcb_get_geometry_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_geometry_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_geometry_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_geometry_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_geometry_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req32)))
	return xcb_ret
}

var _xcb_req32 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_GEOMETRY),
}

func Xxcb_get_geometry_unchecked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t) (r Txcb_get_geometry_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_geometry_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_geometry_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_geometry_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_geometry_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req33)))
	return xcb_ret
}

var _xcb_req33 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_GEOMETRY),
}

func Xxcb_get_geometry_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_geometry_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_query_tree_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* children */
	xcb_block_len += uint32((*Txcb_query_tree_reply_t)(unsafe.Pointer(_aux)).Fchildren_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_query_tree(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_query_tree_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_tree_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_tree_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_tree_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_query_tree_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req34)))
	return xcb_ret
}

var _xcb_req34 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_TREE),
}

func Xxcb_query_tree_unchecked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_query_tree_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_tree_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_tree_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_tree_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_query_tree_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req35)))
	return xcb_ret
}

var _xcb_req35 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_TREE),
}

func Xxcb_query_tree_children(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_query_tree_children_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_query_tree_reply_t)(unsafe.Pointer(R)).Fchildren_len)
}

func Xxcb_query_tree_children_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_query_tree_reply_t)(unsafe.Pointer(R)).Fchildren_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_query_tree_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_tree_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_intern_atom_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_intern_atom_request_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_intern_atom(tls *libc.TLS, c uintptr, only_if_exists Tuint8_t, name_len Tuint16_t, name uintptr) (r Txcb_intern_atom_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_intern_atom_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_intern_atom_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_intern_atom_request_t)(unsafe.Pointer(bp + 48))).Fonly_if_exists = only_if_exists
	(*(*Txcb_intern_atom_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req36)))
	return xcb_ret
}

var _xcb_req36 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_INTERN_ATOM),
}

func Xxcb_intern_atom_unchecked(tls *libc.TLS, c uintptr, only_if_exists Tuint8_t, name_len Tuint16_t, name uintptr) (r Txcb_intern_atom_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_intern_atom_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_intern_atom_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_intern_atom_request_t)(unsafe.Pointer(bp + 48))).Fonly_if_exists = only_if_exists
	(*(*Txcb_intern_atom_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req37)))
	return xcb_ret
}

var _xcb_req37 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_INTERN_ATOM),
}

func Xxcb_intern_atom_reply(tls *libc.TLS, c uintptr, cookie Txcb_intern_atom_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_get_atom_name_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_get_atom_name_reply_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_atom_name(tls *libc.TLS, c uintptr, atom Txcb_atom_t) (r Txcb_get_atom_name_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_atom_name_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_atom_name_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_atom_name_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_atom_name_request_t)(unsafe.Pointer(bp + 32))).Fatom = atom
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req38)))
	return xcb_ret
}

var _xcb_req38 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_ATOM_NAME),
}

func Xxcb_get_atom_name_unchecked(tls *libc.TLS, c uintptr, atom Txcb_atom_t) (r Txcb_get_atom_name_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_atom_name_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_atom_name_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_atom_name_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_atom_name_request_t)(unsafe.Pointer(bp + 32))).Fatom = atom
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req39)))
	return xcb_ret
}

var _xcb_req39 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_ATOM_NAME),
}

func Xxcb_get_atom_name_name(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_atom_name_name_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_get_atom_name_reply_t)(unsafe.Pointer(R)).Fname_len)
}

func Xxcb_get_atom_name_name_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_get_atom_name_reply_t)(unsafe.Pointer(R)).Fname_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_atom_name_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_atom_name_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_change_property_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(24)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* data */
	xcb_block_len += (*Txcb_change_property_request_t)(unsafe.Pointer(_aux)).Fdata_len * uint32((*Txcb_change_property_request_t)(unsafe.Pointer(_aux)).Fformat) / uint32(8) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_property_checked(tls *libc.TLS, c uintptr, mode Tuint8_t, window Txcb_window_t, property Txcb_atom_t, type1 Txcb_atom_t, format Tuint8_t, data_len Tuint32_t, data uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_change_property_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fmode = mode
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fwindow = window
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fproperty = property
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Ftype1 = type1
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fformat = format
	libc.Xmemset(tls, bp+48+17, 0, uint32(3))
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fdata_len = data_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* void data */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = data
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = data_len * uint32(format) / uint32(8) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req40)))
	return xcb_ret
}

var _xcb_req40 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_CHANGE_PROPERTY),
	Fisvoid: uint8(1),
}

func Xxcb_change_property(tls *libc.TLS, c uintptr, mode Tuint8_t, window Txcb_window_t, property Txcb_atom_t, type1 Txcb_atom_t, format Tuint8_t, data_len Tuint32_t, data uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_change_property_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fmode = mode
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fwindow = window
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fproperty = property
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Ftype1 = type1
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fformat = format
	libc.Xmemset(tls, bp+48+17, 0, uint32(3))
	(*(*Txcb_change_property_request_t)(unsafe.Pointer(bp + 48))).Fdata_len = data_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* void data */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = data
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = data_len * uint32(format) / uint32(8) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req41)))
	return xcb_ret
}

var _xcb_req41 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_CHANGE_PROPERTY),
	Fisvoid: uint8(1),
}

func Xxcb_change_property_data(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*24
}

func Xxcb_change_property_data_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_change_property_request_t)(unsafe.Pointer(R)).Fdata_len * uint32((*Txcb_change_property_request_t)(unsafe.Pointer(R)).Fformat) / libc.Uint32FromInt32(8))
}

func Xxcb_change_property_data_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*24 + uintptr((*Txcb_change_property_request_t)(unsafe.Pointer(R)).Fdata_len*uint32((*Txcb_change_property_request_t)(unsafe.Pointer(R)).Fformat)/libc.Uint32FromInt32(8))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_delete_property_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, property Txcb_atom_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_delete_property_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_delete_property_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_delete_property_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_delete_property_request_t)(unsafe.Pointer(bp + 32))).Fproperty = property
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req42)))
	return xcb_ret
}

var _xcb_req42 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_DELETE_PROPERTY),
	Fisvoid: uint8(1),
}

func Xxcb_delete_property(tls *libc.TLS, c uintptr, window Txcb_window_t, property Txcb_atom_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_delete_property_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_delete_property_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_delete_property_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_delete_property_request_t)(unsafe.Pointer(bp + 32))).Fproperty = property
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req43)))
	return xcb_ret
}

var _xcb_req43 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_DELETE_PROPERTY),
	Fisvoid: uint8(1),
}

func Xxcb_get_property_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value */
	xcb_block_len += (*Txcb_get_property_reply_t)(unsafe.Pointer(_aux)).Fvalue_len * libc.Uint32FromInt32(libc.Int32FromUint8((*Txcb_get_property_reply_t)(unsafe.Pointer(_aux)).Fformat)/libc.Int32FromInt32(8)) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_property(tls *libc.TLS, c uintptr, _delete Tuint8_t, window Txcb_window_t, property Txcb_atom_t, type1 Txcb_atom_t, long_offset Tuint32_t, long_length Tuint32_t) (r Txcb_get_property_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_get_property_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_property_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).F_delete = _delete
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Fproperty = property
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Ftype1 = type1
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Flong_offset = long_offset
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Flong_length = long_length
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req44)))
	return xcb_ret
}

var _xcb_req44 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_PROPERTY),
}

func Xxcb_get_property_unchecked(tls *libc.TLS, c uintptr, _delete Tuint8_t, window Txcb_window_t, property Txcb_atom_t, type1 Txcb_atom_t, long_offset Tuint32_t, long_length Tuint32_t) (r Txcb_get_property_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_get_property_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_property_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).F_delete = _delete
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Fproperty = property
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Ftype1 = type1
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Flong_offset = long_offset
	(*(*Txcb_get_property_request_t)(unsafe.Pointer(bp + 32))).Flong_length = long_length
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req45)))
	return xcb_ret
}

var _xcb_req45 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_PROPERTY),
}

func Xxcb_get_property_value(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_property_value_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_get_property_reply_t)(unsafe.Pointer(R)).Fvalue_len * libc.Uint32FromInt32(libc.Int32FromUint8((*Txcb_get_property_reply_t)(unsafe.Pointer(R)).Fformat)/libc.Int32FromInt32(8)))
}

func Xxcb_get_property_value_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_get_property_reply_t)(unsafe.Pointer(R)).Fvalue_len*libc.Uint32FromInt32(libc.Int32FromUint8((*Txcb_get_property_reply_t)(unsafe.Pointer(R)).Fformat)/libc.Int32FromInt32(8)))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_property_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_property_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_list_properties_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* atoms */
	xcb_block_len += uint32((*Txcb_list_properties_reply_t)(unsafe.Pointer(_aux)).Fatoms_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_list_properties(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_list_properties_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_properties_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_properties_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_properties_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_list_properties_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req46)))
	return xcb_ret
}

var _xcb_req46 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_PROPERTIES),
}

func Xxcb_list_properties_unchecked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_list_properties_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_properties_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_properties_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_properties_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_list_properties_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req47)))
	return xcb_ret
}

var _xcb_req47 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_PROPERTIES),
}

func Xxcb_list_properties_atoms(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_list_properties_atoms_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_list_properties_reply_t)(unsafe.Pointer(R)).Fatoms_len)
}

func Xxcb_list_properties_atoms_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_list_properties_reply_t)(unsafe.Pointer(R)).Fatoms_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_properties_reply(tls *libc.TLS, c uintptr, cookie Txcb_list_properties_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_set_selection_owner_checked(tls *libc.TLS, c uintptr, owner Txcb_window_t, selection Txcb_atom_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_selection_owner_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fowner = owner
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fselection = selection
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req48)))
	return xcb_ret
}

var _xcb_req48 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_SELECTION_OWNER),
	Fisvoid: uint8(1),
}

func Xxcb_set_selection_owner(tls *libc.TLS, c uintptr, owner Txcb_window_t, selection Txcb_atom_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_selection_owner_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fowner = owner
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fselection = selection
	(*(*Txcb_set_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req49)))
	return xcb_ret
}

var _xcb_req49 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_SELECTION_OWNER),
	Fisvoid: uint8(1),
}

func Xxcb_get_selection_owner(tls *libc.TLS, c uintptr, selection Txcb_atom_t) (r Txcb_get_selection_owner_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_selection_owner_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_selection_owner_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fselection = selection
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req50)))
	return xcb_ret
}

var _xcb_req50 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_SELECTION_OWNER),
}

func Xxcb_get_selection_owner_unchecked(tls *libc.TLS, c uintptr, selection Txcb_atom_t) (r Txcb_get_selection_owner_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_selection_owner_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_selection_owner_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_selection_owner_request_t)(unsafe.Pointer(bp + 32))).Fselection = selection
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req51)))
	return xcb_ret
}

var _xcb_req51 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_SELECTION_OWNER),
}

func Xxcb_get_selection_owner_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_selection_owner_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_convert_selection_checked(tls *libc.TLS, c uintptr, requestor Txcb_window_t, selection Txcb_atom_t, target Txcb_atom_t, property Txcb_atom_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_convert_selection_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Frequestor = requestor
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Fselection = selection
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Ftarget = target
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Fproperty = property
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req52)))
	return xcb_ret
}

var _xcb_req52 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CONVERT_SELECTION),
	Fisvoid: uint8(1),
}

func Xxcb_convert_selection(tls *libc.TLS, c uintptr, requestor Txcb_window_t, selection Txcb_atom_t, target Txcb_atom_t, property Txcb_atom_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_convert_selection_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Frequestor = requestor
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Fselection = selection
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Ftarget = target
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Fproperty = property
	(*(*Txcb_convert_selection_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req53)))
	return xcb_ret
}

var _xcb_req53 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CONVERT_SELECTION),
	Fisvoid: uint8(1),
}

func Xxcb_send_event_checked(tls *libc.TLS, c uintptr, propagate Tuint8_t, destination Txcb_window_t, event_mask Tuint32_t, event uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_send_event_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_send_event_request_t)(unsafe.Pointer(bp + 32))).Fpropagate = propagate
	(*(*Txcb_send_event_request_t)(unsafe.Pointer(bp + 32))).Fdestination = destination
	(*(*Txcb_send_event_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	libc.Xmemcpy(tls, bp+32+12, event, uint32(32))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(44)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req54)))
	return xcb_ret
}

var _xcb_req54 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SEND_EVENT),
	Fisvoid: uint8(1),
}

func Xxcb_send_event(tls *libc.TLS, c uintptr, propagate Tuint8_t, destination Txcb_window_t, event_mask Tuint32_t, event uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_send_event_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_send_event_request_t)(unsafe.Pointer(bp + 32))).Fpropagate = propagate
	(*(*Txcb_send_event_request_t)(unsafe.Pointer(bp + 32))).Fdestination = destination
	(*(*Txcb_send_event_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	libc.Xmemcpy(tls, bp+32+12, event, uint32(32))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(44)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req55)))
	return xcb_ret
}

var _xcb_req55 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SEND_EVENT),
	Fisvoid: uint8(1),
}

func Xxcb_grab_pointer(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, event_mask Tuint16_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t, confine_to Txcb_window_t, cursor Txcb_cursor_t, time Txcb_timestamp_t) (r Txcb_grab_pointer_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_grab_pointer_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fconfine_to = confine_to
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req56)))
	return xcb_ret
}

var _xcb_req56 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_POINTER),
}

func Xxcb_grab_pointer_unchecked(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, event_mask Tuint16_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t, confine_to Txcb_window_t, cursor Txcb_cursor_t, time Txcb_timestamp_t) (r Txcb_grab_pointer_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_grab_pointer_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fconfine_to = confine_to
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_grab_pointer_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req57)))
	return xcb_ret
}

var _xcb_req57 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_POINTER),
}

func Xxcb_grab_pointer_reply(tls *libc.TLS, c uintptr, cookie Txcb_grab_pointer_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_ungrab_pointer_checked(tls *libc.TLS, c uintptr, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_ungrab_pointer_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req58)))
	return xcb_ret
}

var _xcb_req58 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_POINTER),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_pointer(tls *libc.TLS, c uintptr, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_ungrab_pointer_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req59)))
	return xcb_ret
}

var _xcb_req59 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_POINTER),
	Fisvoid: uint8(1),
}

func Xxcb_grab_button_checked(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, event_mask Tuint16_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t, confine_to Txcb_window_t, cursor Txcb_cursor_t, button Tuint8_t, modifiers Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_button_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fconfine_to = confine_to
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fbutton = button
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req60)))
	return xcb_ret
}

var _xcb_req60 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_BUTTON),
	Fisvoid: uint8(1),
}

func Xxcb_grab_button(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, event_mask Tuint16_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t, confine_to Txcb_window_t, cursor Txcb_cursor_t, button Tuint8_t, modifiers Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_button_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fconfine_to = confine_to
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fbutton = button
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_grab_button_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req61)))
	return xcb_ret
}

var _xcb_req61 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_BUTTON),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_button_checked(tls *libc.TLS, c uintptr, button Tuint8_t, grab_window Txcb_window_t, modifiers Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_button_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_button_request_t)(unsafe.Pointer(bp + 32))).Fbutton = button
	(*(*Txcb_ungrab_button_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_ungrab_button_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	libc.Xmemset(tls, bp+32+10, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req62)))
	return xcb_ret
}

var _xcb_req62 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_BUTTON),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_button(tls *libc.TLS, c uintptr, button Tuint8_t, grab_window Txcb_window_t, modifiers Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_button_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_button_request_t)(unsafe.Pointer(bp + 32))).Fbutton = button
	(*(*Txcb_ungrab_button_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_ungrab_button_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	libc.Xmemset(tls, bp+32+10, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req63)))
	return xcb_ret
}

var _xcb_req63 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_BUTTON),
	Fisvoid: uint8(1),
}

func Xxcb_change_active_pointer_grab_checked(tls *libc.TLS, c uintptr, cursor Txcb_cursor_t, time Txcb_timestamp_t, event_mask Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_change_active_pointer_grab_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	libc.Xmemset(tls, bp+32+14, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req64)))
	return xcb_ret
}

var _xcb_req64 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CHANGE_ACTIVE_POINTER_GRAB),
	Fisvoid: uint8(1),
}

func Xxcb_change_active_pointer_grab(tls *libc.TLS, c uintptr, cursor Txcb_cursor_t, time Txcb_timestamp_t, event_mask Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_change_active_pointer_grab_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*Txcb_change_active_pointer_grab_request_t)(unsafe.Pointer(bp + 32))).Fevent_mask = event_mask
	libc.Xmemset(tls, bp+32+14, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req65)))
	return xcb_ret
}

var _xcb_req65 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CHANGE_ACTIVE_POINTER_GRAB),
	Fisvoid: uint8(1),
}

func Xxcb_grab_keyboard(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, time Txcb_timestamp_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t) (r Txcb_grab_keyboard_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_grab_keyboard_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_keyboard_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	libc.Xmemset(tls, bp+32+14, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req66)))
	return xcb_ret
}

var _xcb_req66 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_KEYBOARD),
}

func Xxcb_grab_keyboard_unchecked(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, time Txcb_timestamp_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t) (r Txcb_grab_keyboard_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_grab_keyboard_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_keyboard_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	libc.Xmemset(tls, bp+32+14, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req67)))
	return xcb_ret
}

var _xcb_req67 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_KEYBOARD),
}

func Xxcb_grab_keyboard_reply(tls *libc.TLS, c uintptr, cookie Txcb_grab_keyboard_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_ungrab_keyboard_checked(tls *libc.TLS, c uintptr, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_keyboard_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_ungrab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req68)))
	return xcb_ret
}

var _xcb_req68 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_KEYBOARD),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_keyboard(tls *libc.TLS, c uintptr, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_keyboard_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_ungrab_keyboard_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req69)))
	return xcb_ret
}

var _xcb_req69 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_KEYBOARD),
	Fisvoid: uint8(1),
}

func Xxcb_grab_key_checked(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, modifiers Tuint16_t, key Txcb_keycode_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_key_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fkey = key
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	libc.Xmemset(tls, bp+32+13, 0, uint32(3))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req70)))
	return xcb_ret
}

var _xcb_req70 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_KEY),
	Fisvoid: uint8(1),
}

func Xxcb_grab_key(tls *libc.TLS, c uintptr, owner_events Tuint8_t, grab_window Txcb_window_t, modifiers Tuint16_t, key Txcb_keycode_t, pointer_mode Tuint8_t, keyboard_mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_key_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fowner_events = owner_events
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fkey = key
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fpointer_mode = pointer_mode
	(*(*Txcb_grab_key_request_t)(unsafe.Pointer(bp + 32))).Fkeyboard_mode = keyboard_mode
	libc.Xmemset(tls, bp+32+13, 0, uint32(3))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req71)))
	return xcb_ret
}

var _xcb_req71 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_KEY),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_key_checked(tls *libc.TLS, c uintptr, key Txcb_keycode_t, grab_window Txcb_window_t, modifiers Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_key_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_key_request_t)(unsafe.Pointer(bp + 32))).Fkey = key
	(*(*Txcb_ungrab_key_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_ungrab_key_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	libc.Xmemset(tls, bp+32+10, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req72)))
	return xcb_ret
}

var _xcb_req72 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_KEY),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_key(tls *libc.TLS, c uintptr, key Txcb_keycode_t, grab_window Txcb_window_t, modifiers Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_key_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_key_request_t)(unsafe.Pointer(bp + 32))).Fkey = key
	(*(*Txcb_ungrab_key_request_t)(unsafe.Pointer(bp + 32))).Fgrab_window = grab_window
	(*(*Txcb_ungrab_key_request_t)(unsafe.Pointer(bp + 32))).Fmodifiers = modifiers
	libc.Xmemset(tls, bp+32+10, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req73)))
	return xcb_ret
}

var _xcb_req73 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_KEY),
	Fisvoid: uint8(1),
}

func Xxcb_allow_events_checked(tls *libc.TLS, c uintptr, mode Tuint8_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_allow_events_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_allow_events_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*Txcb_allow_events_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req74)))
	return xcb_ret
}

var _xcb_req74 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOW_EVENTS),
	Fisvoid: uint8(1),
}

func Xxcb_allow_events(tls *libc.TLS, c uintptr, mode Tuint8_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_allow_events_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_allow_events_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*Txcb_allow_events_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req75)))
	return xcb_ret
}

var _xcb_req75 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOW_EVENTS),
	Fisvoid: uint8(1),
}

func Xxcb_grab_server_checked(tls *libc.TLS, c uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_server_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_server_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req76)))
	return xcb_ret
}

var _xcb_req76 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_SERVER),
	Fisvoid: uint8(1),
}

func Xxcb_grab_server(tls *libc.TLS, c uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_grab_server_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_grab_server_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req77)))
	return xcb_ret
}

var _xcb_req77 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GRAB_SERVER),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_server_checked(tls *libc.TLS, c uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_server_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_server_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req78)))
	return xcb_ret
}

var _xcb_req78 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_SERVER),
	Fisvoid: uint8(1),
}

func Xxcb_ungrab_server(tls *libc.TLS, c uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_ungrab_server_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_ungrab_server_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req79)))
	return xcb_ret
}

var _xcb_req79 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNGRAB_SERVER),
	Fisvoid: uint8(1),
}

func Xxcb_query_pointer(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_query_pointer_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_pointer_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_query_pointer_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req80)))
	return xcb_ret
}

var _xcb_req80 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_POINTER),
}

func Xxcb_query_pointer_unchecked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_query_pointer_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_pointer_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_query_pointer_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req81)))
	return xcb_ret
}

var _xcb_req81 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_POINTER),
}

func Xxcb_query_pointer_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_pointer_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_timecoord_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_timecoord_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_timecoord_iterator_t)(unsafe.Pointer(i)).Fdata += 8
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(8))
}

func Xxcb_timecoord_end(tls *libc.TLS, i Txcb_timecoord_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*8
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_get_motion_events_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* events */
	xcb_block_len += (*Txcb_get_motion_events_reply_t)(unsafe.Pointer(_aux)).Fevents_len * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_motion_events(tls *libc.TLS, c uintptr, window Txcb_window_t, start Txcb_timestamp_t, stop Txcb_timestamp_t) (r Txcb_get_motion_events_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_motion_events_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_motion_events_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fstart = start
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fstop = stop
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req82)))
	return xcb_ret
}

var _xcb_req82 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_MOTION_EVENTS),
}

func Xxcb_get_motion_events_unchecked(tls *libc.TLS, c uintptr, window Txcb_window_t, start Txcb_timestamp_t, stop Txcb_timestamp_t) (r Txcb_get_motion_events_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_motion_events_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_motion_events_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fstart = start
	(*(*Txcb_get_motion_events_request_t)(unsafe.Pointer(bp + 32))).Fstop = stop
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req83)))
	return xcb_ret
}

var _xcb_req83 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_MOTION_EVENTS),
}

func Xxcb_get_motion_events_events(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_motion_events_events_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_get_motion_events_reply_t)(unsafe.Pointer(R)).Fevents_len)
}

func Xxcb_get_motion_events_events_iterator(tls *libc.TLS, R uintptr) (r Txcb_timecoord_iterator_t) {
	var i Txcb_timecoord_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32
	i.Frem = libc.Int32FromUint32((*Txcb_get_motion_events_reply_t)(unsafe.Pointer(R)).Fevents_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_motion_events_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_motion_events_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_translate_coordinates(tls *libc.TLS, c uintptr, src_window Txcb_window_t, dst_window Txcb_window_t, src_x Tint16_t, src_y Tint16_t) (r Txcb_translate_coordinates_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_translate_coordinates_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_translate_coordinates_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fsrc_window = src_window
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fdst_window = dst_window
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req84)))
	return xcb_ret
}

var _xcb_req84 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_TRANSLATE_COORDINATES),
}

func Xxcb_translate_coordinates_unchecked(tls *libc.TLS, c uintptr, src_window Txcb_window_t, dst_window Txcb_window_t, src_x Tint16_t, src_y Tint16_t) (r Txcb_translate_coordinates_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_translate_coordinates_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_translate_coordinates_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fsrc_window = src_window
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fdst_window = dst_window
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_translate_coordinates_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req85)))
	return xcb_ret
}

var _xcb_req85 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_TRANSLATE_COORDINATES),
}

func Xxcb_translate_coordinates_reply(tls *libc.TLS, c uintptr, cookie Txcb_translate_coordinates_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_warp_pointer_checked(tls *libc.TLS, c uintptr, src_window Txcb_window_t, dst_window Txcb_window_t, src_x Tint16_t, src_y Tint16_t, src_width Tuint16_t, src_height Tuint16_t, dst_x Tint16_t, dst_y Tint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_warp_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_window = src_window
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fdst_window = dst_window
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_width = src_width
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_height = src_height
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fdst_x = dst_x
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fdst_y = dst_y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req86)))
	return xcb_ret
}

var _xcb_req86 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_WARP_POINTER),
	Fisvoid: uint8(1),
}

func Xxcb_warp_pointer(tls *libc.TLS, c uintptr, src_window Txcb_window_t, dst_window Txcb_window_t, src_x Tint16_t, src_y Tint16_t, src_width Tuint16_t, src_height Tuint16_t, dst_x Tint16_t, dst_y Tint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_warp_pointer_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_window = src_window
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fdst_window = dst_window
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_width = src_width
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fsrc_height = src_height
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fdst_x = dst_x
	(*(*Txcb_warp_pointer_request_t)(unsafe.Pointer(bp + 32))).Fdst_y = dst_y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req87)))
	return xcb_ret
}

var _xcb_req87 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_WARP_POINTER),
	Fisvoid: uint8(1),
}

func Xxcb_set_input_focus_checked(tls *libc.TLS, c uintptr, revert_to Tuint8_t, focus Txcb_window_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_input_focus_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_input_focus_request_t)(unsafe.Pointer(bp + 32))).Frevert_to = revert_to
	(*(*Txcb_set_input_focus_request_t)(unsafe.Pointer(bp + 32))).Ffocus = focus
	(*(*Txcb_set_input_focus_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req88)))
	return xcb_ret
}

var _xcb_req88 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_INPUT_FOCUS),
	Fisvoid: uint8(1),
}

func Xxcb_set_input_focus(tls *libc.TLS, c uintptr, revert_to Tuint8_t, focus Txcb_window_t, time Txcb_timestamp_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_input_focus_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_input_focus_request_t)(unsafe.Pointer(bp + 32))).Frevert_to = revert_to
	(*(*Txcb_set_input_focus_request_t)(unsafe.Pointer(bp + 32))).Ffocus = focus
	(*(*Txcb_set_input_focus_request_t)(unsafe.Pointer(bp + 32))).Ftime = time
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req89)))
	return xcb_ret
}

var _xcb_req89 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_INPUT_FOCUS),
	Fisvoid: uint8(1),
}

func Xxcb_get_input_focus(tls *libc.TLS, c uintptr) (r Txcb_get_input_focus_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_input_focus_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_input_focus_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_input_focus_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req90)))
	return xcb_ret
}

var _xcb_req90 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_INPUT_FOCUS),
}

func Xxcb_get_input_focus_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_input_focus_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_input_focus_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_input_focus_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_input_focus_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req91)))
	return xcb_ret
}

var _xcb_req91 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_INPUT_FOCUS),
}

func Xxcb_get_input_focus_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_input_focus_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_query_keymap(tls *libc.TLS, c uintptr) (r Txcb_query_keymap_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_keymap_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_keymap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_keymap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req92)))
	return xcb_ret
}

var _xcb_req92 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_KEYMAP),
}

func Xxcb_query_keymap_unchecked(tls *libc.TLS, c uintptr) (r Txcb_query_keymap_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_keymap_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_keymap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_keymap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req93)))
	return xcb_ret
}

var _xcb_req93 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_KEYMAP),
}

func Xxcb_query_keymap_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_keymap_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_open_font_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_open_font_request_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_open_font_checked(tls *libc.TLS, c uintptr, fid Txcb_font_t, name_len Tuint16_t, name uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_open_font_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_open_font_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_open_font_request_t)(unsafe.Pointer(bp + 48))).Ffid = fid
	(*(*Txcb_open_font_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+10, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req94)))
	return xcb_ret
}

var _xcb_req94 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_OPEN_FONT),
	Fisvoid: uint8(1),
}

func Xxcb_open_font(tls *libc.TLS, c uintptr, fid Txcb_font_t, name_len Tuint16_t, name uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_open_font_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_open_font_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_open_font_request_t)(unsafe.Pointer(bp + 48))).Ffid = fid
	(*(*Txcb_open_font_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+10, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req95)))
	return xcb_ret
}

var _xcb_req95 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_OPEN_FONT),
	Fisvoid: uint8(1),
}

func Xxcb_open_font_name(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_open_font_name_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_open_font_request_t)(unsafe.Pointer(R)).Fname_len)
}

func Xxcb_open_font_name_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12 + uintptr((*Txcb_open_font_request_t)(unsafe.Pointer(R)).Fname_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_close_font_checked(tls *libc.TLS, c uintptr, font Txcb_font_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_close_font_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_close_font_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_close_font_request_t)(unsafe.Pointer(bp + 32))).Ffont = font
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req96)))
	return xcb_ret
}

var _xcb_req96 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CLOSE_FONT),
	Fisvoid: uint8(1),
}

func Xxcb_close_font(tls *libc.TLS, c uintptr, font Txcb_font_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_close_font_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_close_font_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_close_font_request_t)(unsafe.Pointer(bp + 32))).Ffont = font
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req97)))
	return xcb_ret
}

var _xcb_req97 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CLOSE_FONT),
	Fisvoid: uint8(1),
}

func Xxcb_fontprop_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_fontprop_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_fontprop_iterator_t)(unsafe.Pointer(i)).Fdata += 8
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(8))
}

func Xxcb_fontprop_end(tls *libc.TLS, i Txcb_fontprop_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*8
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_charinfo_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_charinfo_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_charinfo_iterator_t)(unsafe.Pointer(i)).Fdata += 12
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(12))
}

func Xxcb_charinfo_end(tls *libc.TLS, i Txcb_charinfo_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*12
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_query_font_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(60)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* properties */
	xcb_block_len += uint32((*Txcb_query_font_reply_t)(unsafe.Pointer(_aux)).Fproperties_len) * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* char_infos */
	xcb_block_len += (*Txcb_query_font_reply_t)(unsafe.Pointer(_aux)).Fchar_infos_len * uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_query_font(tls *libc.TLS, c uintptr, font Txcb_fontable_t) (r Txcb_query_font_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_font_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_font_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_font_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_query_font_request_t)(unsafe.Pointer(bp + 32))).Ffont = font
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req98)))
	return xcb_ret
}

var _xcb_req98 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_FONT),
}

func Xxcb_query_font_unchecked(tls *libc.TLS, c uintptr, font Txcb_fontable_t) (r Txcb_query_font_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_font_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_font_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_font_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_query_font_request_t)(unsafe.Pointer(bp + 32))).Ffont = font
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req99)))
	return xcb_ret
}

var _xcb_req99 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_FONT),
}

func Xxcb_query_font_properties(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*60
}

func Xxcb_query_font_properties_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_query_font_reply_t)(unsafe.Pointer(R)).Fproperties_len)
}

func Xxcb_query_font_properties_iterator(tls *libc.TLS, R uintptr) (r Txcb_fontprop_iterator_t) {
	var i Txcb_fontprop_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*60
	i.Frem = libc.Int32FromUint16((*Txcb_query_font_reply_t)(unsafe.Pointer(R)).Fproperties_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_query_font_char_infos(tls *libc.TLS, R uintptr) (r uintptr) {
	var prev Txcb_generic_iterator_t
	_ = prev
	prev = Xxcb_fontprop_end(tls, Xxcb_query_font_properties_iterator(tls, R))
	return prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&libc.Uint32FromInt32(3)) + libc.UintptrFromInt32(0)
}

func Xxcb_query_font_char_infos_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_query_font_reply_t)(unsafe.Pointer(R)).Fchar_infos_len)
}

func Xxcb_query_font_char_infos_iterator(tls *libc.TLS, R uintptr) (r Txcb_charinfo_iterator_t) {
	var i Txcb_charinfo_iterator_t
	var prev Txcb_generic_iterator_t
	_, _ = i, prev
	prev = Xxcb_fontprop_end(tls, Xxcb_query_font_properties_iterator(tls, R))
	i.Fdata = prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&libc.Uint32FromInt32(3))
	i.Frem = libc.Int32FromUint32((*Txcb_query_font_reply_t)(unsafe.Pointer(R)).Fchar_infos_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_query_font_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_font_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_query_text_extents_sizeof(tls *libc.TLS, _buffer uintptr, string_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* string */
	xcb_block_len += string_len * uint32(2)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_query_text_extents(tls *libc.TLS, c uintptr, font Txcb_fontable_t, string_len Tuint32_t, string1 uintptr) (r Txcb_query_text_extents_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_query_text_extents_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_query_text_extents_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_text_extents_request_t)(unsafe.Pointer(bp + 48))).Fodd_length = uint8(string_len & libc.Uint32FromInt32(1))
	(*(*Txcb_query_text_extents_request_t)(unsafe.Pointer(bp + 48))).Ffont = font
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_char2b_t string */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = string1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = string_len * uint32(2)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req100)))
	return xcb_ret
}

var _xcb_req100 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_QUERY_TEXT_EXTENTS),
}

func Xxcb_query_text_extents_unchecked(tls *libc.TLS, c uintptr, font Txcb_fontable_t, string_len Tuint32_t, string1 uintptr) (r Txcb_query_text_extents_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_query_text_extents_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_query_text_extents_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_text_extents_request_t)(unsafe.Pointer(bp + 48))).Fodd_length = uint8(string_len & libc.Uint32FromInt32(1))
	(*(*Txcb_query_text_extents_request_t)(unsafe.Pointer(bp + 48))).Ffont = font
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_char2b_t string */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = string1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = string_len * uint32(2)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req101)))
	return xcb_ret
}

var _xcb_req101 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_QUERY_TEXT_EXTENTS),
}

func Xxcb_query_text_extents_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_text_extents_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_str_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_str_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_str_name(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)
}

func Xxcb_str_name_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_str_t)(unsafe.Pointer(R)).Fname_len)
}

func Xxcb_str_name_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1) + uintptr((*Txcb_str_t)(unsafe.Pointer(R)).Fname_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_str_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_str_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_str_sizeof(tls, R))
	(*Txcb_str_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_str_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_str_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_str_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_str_end(tls *libc.TLS, _i Txcb_str_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_str_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_str_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_str_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_str_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_str_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_str_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_list_fonts_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* pattern */
	xcb_block_len += uint32((*Txcb_list_fonts_request_t)(unsafe.Pointer(_aux)).Fpattern_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_list_fonts(tls *libc.TLS, c uintptr, max_names Tuint16_t, pattern_len Tuint16_t, pattern uintptr) (r Txcb_list_fonts_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_list_fonts_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_list_fonts_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_fonts_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_list_fonts_request_t)(unsafe.Pointer(bp + 48))).Fmax_names = max_names
	(*(*Txcb_list_fonts_request_t)(unsafe.Pointer(bp + 48))).Fpattern_len = pattern_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char pattern */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pattern
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(pattern_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req102)))
	return xcb_ret
}

var _xcb_req102 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_LIST_FONTS),
}

func Xxcb_list_fonts_unchecked(tls *libc.TLS, c uintptr, max_names Tuint16_t, pattern_len Tuint16_t, pattern uintptr) (r Txcb_list_fonts_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_list_fonts_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_list_fonts_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_fonts_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_list_fonts_request_t)(unsafe.Pointer(bp + 48))).Fmax_names = max_names
	(*(*Txcb_list_fonts_request_t)(unsafe.Pointer(bp + 48))).Fpattern_len = pattern_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char pattern */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pattern
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(pattern_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req103)))
	return xcb_ret
}

var _xcb_req103 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_LIST_FONTS),
}

func Xxcb_list_fonts_names_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_list_fonts_reply_t)(unsafe.Pointer(R)).Fnames_len)
}

func Xxcb_list_fonts_names_iterator(tls *libc.TLS, R uintptr) (r Txcb_str_iterator_t) {
	var i Txcb_str_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32
	i.Frem = libc.Int32FromUint16((*Txcb_list_fonts_reply_t)(unsafe.Pointer(R)).Fnames_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_fonts_reply(tls *libc.TLS, c uintptr, cookie Txcb_list_fonts_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_list_fonts_with_info_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* pattern */
	xcb_block_len += uint32((*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(_aux)).Fpattern_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_list_fonts_with_info(tls *libc.TLS, c uintptr, max_names Tuint16_t, pattern_len Tuint16_t, pattern uintptr) (r Txcb_list_fonts_with_info_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_list_fonts_with_info_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_list_fonts_with_info_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(bp + 48))).Fmax_names = max_names
	(*(*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(bp + 48))).Fpattern_len = pattern_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char pattern */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pattern
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(pattern_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req104)))
	return xcb_ret
}

var _xcb_req104 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_LIST_FONTS_WITH_INFO),
}

func Xxcb_list_fonts_with_info_unchecked(tls *libc.TLS, c uintptr, max_names Tuint16_t, pattern_len Tuint16_t, pattern uintptr) (r Txcb_list_fonts_with_info_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_list_fonts_with_info_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_list_fonts_with_info_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(bp + 48))).Fmax_names = max_names
	(*(*Txcb_list_fonts_with_info_request_t)(unsafe.Pointer(bp + 48))).Fpattern_len = pattern_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char pattern */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pattern
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(pattern_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req105)))
	return xcb_ret
}

var _xcb_req105 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_LIST_FONTS_WITH_INFO),
}

func Xxcb_list_fonts_with_info_properties(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*60
}

func Xxcb_list_fonts_with_info_properties_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_list_fonts_with_info_reply_t)(unsafe.Pointer(R)).Fproperties_len)
}

func Xxcb_list_fonts_with_info_properties_iterator(tls *libc.TLS, R uintptr) (r Txcb_fontprop_iterator_t) {
	var i Txcb_fontprop_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*60
	i.Frem = libc.Int32FromUint16((*Txcb_list_fonts_with_info_reply_t)(unsafe.Pointer(R)).Fproperties_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_fonts_with_info_name(tls *libc.TLS, R uintptr) (r uintptr) {
	var prev Txcb_generic_iterator_t
	_ = prev
	prev = Xxcb_fontprop_end(tls, Xxcb_list_fonts_with_info_properties_iterator(tls, R))
	return prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&(libc.Uint32FromInt64(1)-libc.Uint32FromInt32(1))) + libc.UintptrFromInt32(0)
}

func Xxcb_list_fonts_with_info_name_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_list_fonts_with_info_reply_t)(unsafe.Pointer(R)).Fname_len)
}

func Xxcb_list_fonts_with_info_name_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i, prev Txcb_generic_iterator_t
	_, _ = i, prev
	prev = Xxcb_fontprop_end(tls, Xxcb_list_fonts_with_info_properties_iterator(tls, R))
	i.Fdata = prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&(libc.Uint32FromInt64(1)-libc.Uint32FromInt32(1))) + uintptr((*Txcb_list_fonts_with_info_reply_t)(unsafe.Pointer(R)).Fname_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_fonts_with_info_reply(tls *libc.TLS, c uintptr, cookie Txcb_list_fonts_with_info_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_set_font_path_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp_len uint32
	_, _, _, _, _, _, _, _ = _aux, i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp, xcb_tmp_len
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* font */
	i = uint32(0)
	for {
		if !(i < uint32((*Txcb_set_font_path_request_t)(unsafe.Pointer(_aux)).Ffont_qty)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_str_sizeof(tls, xcb_tmp))
		xcb_block_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_set_font_path_checked(tls *libc.TLS, c uintptr, font_qty Tuint16_t, font uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, xcb_tmp_len uint32
	var xcb_ret Txcb_void_cookie_t
	var xcb_tmp uintptr
	var _ /* xcb_out at bp+48 */ Txcb_set_font_path_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_, _, _, _ = i, xcb_ret, xcb_tmp, xcb_tmp_len
	(*(*Txcb_set_font_path_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_set_font_path_request_t)(unsafe.Pointer(bp + 48))).Ffont_qty = font_qty
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_str_t font */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = font
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(0)
	xcb_tmp = font
	i = uint32(0)
	for {
		if !(i < uint32(font_qty)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_str_sizeof(tls, xcb_tmp))
		(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req106)))
	return xcb_ret
}

var _xcb_req106 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_FONT_PATH),
	Fisvoid: uint8(1),
}

func Xxcb_set_font_path(tls *libc.TLS, c uintptr, font_qty Tuint16_t, font uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, xcb_tmp_len uint32
	var xcb_ret Txcb_void_cookie_t
	var xcb_tmp uintptr
	var _ /* xcb_out at bp+48 */ Txcb_set_font_path_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_, _, _, _ = i, xcb_ret, xcb_tmp, xcb_tmp_len
	(*(*Txcb_set_font_path_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_set_font_path_request_t)(unsafe.Pointer(bp + 48))).Ffont_qty = font_qty
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_str_t font */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = font
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(0)
	xcb_tmp = font
	i = uint32(0)
	for {
		if !(i < uint32(font_qty)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_str_sizeof(tls, xcb_tmp))
		(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req107)))
	return xcb_ret
}

var _xcb_req107 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_FONT_PATH),
	Fisvoid: uint8(1),
}

func Xxcb_set_font_path_font_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_set_font_path_request_t)(unsafe.Pointer(R)).Ffont_qty)
}

func Xxcb_set_font_path_font_iterator(tls *libc.TLS, R uintptr) (r Txcb_str_iterator_t) {
	var i Txcb_str_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8
	i.Frem = libc.Int32FromUint16((*Txcb_set_font_path_request_t)(unsafe.Pointer(R)).Ffont_qty)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_font_path_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp_len uint32
	_, _, _, _, _, _, _, _ = _aux, i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp, xcb_tmp_len
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* path */
	i = uint32(0)
	for {
		if !(i < uint32((*Txcb_get_font_path_reply_t)(unsafe.Pointer(_aux)).Fpath_len)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_str_sizeof(tls, xcb_tmp))
		xcb_block_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_font_path(tls *libc.TLS, c uintptr) (r Txcb_get_font_path_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_font_path_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_font_path_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_font_path_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req108)))
	return xcb_ret
}

var _xcb_req108 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_FONT_PATH),
}

func Xxcb_get_font_path_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_font_path_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_font_path_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_font_path_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_font_path_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req109)))
	return xcb_ret
}

var _xcb_req109 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_FONT_PATH),
}

func Xxcb_get_font_path_path_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_get_font_path_reply_t)(unsafe.Pointer(R)).Fpath_len)
}

func Xxcb_get_font_path_path_iterator(tls *libc.TLS, R uintptr) (r Txcb_str_iterator_t) {
	var i Txcb_str_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32
	i.Frem = libc.Int32FromUint16((*Txcb_get_font_path_reply_t)(unsafe.Pointer(R)).Fpath_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_font_path_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_font_path_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_create_pixmap_checked(tls *libc.TLS, c uintptr, depth Tuint8_t, pid Txcb_pixmap_t, drawable Txcb_drawable_t, width Tuint16_t, height Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_pixmap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fdepth = depth
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fpid = pid
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req110)))
	return xcb_ret
}

var _xcb_req110 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_PIXMAP),
	Fisvoid: uint8(1),
}

func Xxcb_create_pixmap(tls *libc.TLS, c uintptr, depth Tuint8_t, pid Txcb_pixmap_t, drawable Txcb_drawable_t, width Tuint16_t, height Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_pixmap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fdepth = depth
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fpid = pid
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_create_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req111)))
	return xcb_ret
}

var _xcb_req111 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_PIXMAP),
	Fisvoid: uint8(1),
}

func Xxcb_free_pixmap_checked(tls *libc.TLS, c uintptr, pixmap Txcb_pixmap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_pixmap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fpixmap = pixmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req112)))
	return xcb_ret
}

var _xcb_req112 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_PIXMAP),
	Fisvoid: uint8(1),
}

func Xxcb_free_pixmap(tls *libc.TLS, c uintptr, pixmap Txcb_pixmap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_pixmap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_pixmap_request_t)(unsafe.Pointer(bp + 32))).Fpixmap = pixmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req113)))
	return xcb_ret
}

var _xcb_req113 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_PIXMAP),
	Fisvoid: uint8(1),
}

func Xxcb_create_gc_value_list_serialize(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_parts_idx uint32
	var xcb_out, xcb_tmp uintptr
	var xcb_parts [24]Tiovec
	var _ /* xcb_pad0 at bp+0 */ [3]int8
	_, _, _, _, _, _, _, _, _, _ = i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_out, xcb_pad, xcb_padding_offset, xcb_parts, xcb_parts_idx, xcb_tmp
	xcb_out = *(*uintptr)(unsafe.Pointer(_buffer))
	xcb_buffer_len = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	xcb_pad = uint32(0)
	*(*[3]int8)(unsafe.Pointer(bp)) = [3]int8{}
	xcb_parts_idx = uint32(0)
	xcb_block_len = uint32(0)
	if value_mask&uint32(_XCB_GC_FUNCTION) != 0 {
		/* xcb_create_gc_value_list_t.function */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_PLANE_MASK) != 0 {
		/* xcb_create_gc_value_list_t.plane_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 4
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FOREGROUND) != 0 {
		/* xcb_create_gc_value_list_t.foreground */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 8
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_BACKGROUND) != 0 {
		/* xcb_create_gc_value_list_t.background */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 12
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_WIDTH) != 0 {
		/* xcb_create_gc_value_list_t.line_width */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 16
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.line_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 20
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CAP_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.cap_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 24
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_JOIN_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.join_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 28
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.fill_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 32
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_RULE) != 0 {
		/* xcb_create_gc_value_list_t.fill_rule */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 36
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE) != 0 {
		/* xcb_create_gc_value_list_t.tile */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 40
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_STIPPLE) != 0 {
		/* xcb_create_gc_value_list_t.stipple */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 44
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_X) != 0 {
		/* xcb_create_gc_value_list_t.tile_stipple_x_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 48
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_Y) != 0 {
		/* xcb_create_gc_value_list_t.tile_stipple_y_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 52
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FONT) != 0 {
		/* xcb_create_gc_value_list_t.font */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 56
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_SUBWINDOW_MODE) != 0 {
		/* xcb_create_gc_value_list_t.subwindow_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 60
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_GRAPHICS_EXPOSURES) != 0 {
		/* xcb_create_gc_value_list_t.graphics_exposures */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 64
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_X) != 0 {
		/* xcb_create_gc_value_list_t.clip_x_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 68
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_Y) != 0 {
		/* xcb_create_gc_value_list_t.clip_y_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 72
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_MASK) != 0 {
		/* xcb_create_gc_value_list_t.clip_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 76
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_OFFSET) != 0 {
		/* xcb_create_gc_value_list_t.dash_offset */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 80
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_LIST) != 0 {
		/* xcb_create_gc_value_list_t.dashes */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 84
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_ARC_MODE) != 0 {
		/* xcb_create_gc_value_list_t.arc_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 88
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_parts[xcb_parts_idx].Fiov_base = bp
		xcb_parts[xcb_parts_idx].Fiov_len = xcb_pad
		xcb_parts_idx++
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.UintptrFromInt32(0) == xcb_out {
		/* allocate memory */
		xcb_out = libc.Xmalloc(tls, xcb_buffer_len)
		*(*uintptr)(unsafe.Pointer(_buffer)) = xcb_out
	}
	xcb_tmp = xcb_out
	i = uint32(0)
	for {
		if !(i < xcb_parts_idx) {
			break
		}
		if uintptr(0) != xcb_parts[i].Fiov_base && uint32(0) != xcb_parts[i].Fiov_len {
			libc.Xmemcpy(tls, xcb_tmp, xcb_parts[i].Fiov_base, xcb_parts[i].Fiov_len)
		}
		if uint32(0) != xcb_parts[i].Fiov_len {
			xcb_tmp += uintptr(xcb_parts[i].Fiov_len)
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_create_gc_value_list_unpack(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset uint32
	var xcb_tmp uintptr
	_, _, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	if value_mask&uint32(_XCB_GC_FUNCTION) != 0 {
		/* xcb_create_gc_value_list_t.function */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ffunction = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_PLANE_MASK) != 0 {
		/* xcb_create_gc_value_list_t.plane_mask */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fplane_mask = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FOREGROUND) != 0 {
		/* xcb_create_gc_value_list_t.foreground */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fforeground = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_BACKGROUND) != 0 {
		/* xcb_create_gc_value_list_t.background */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fbackground = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_WIDTH) != 0 {
		/* xcb_create_gc_value_list_t.line_width */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fline_width = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.line_style */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fline_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CAP_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.cap_style */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fcap_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_JOIN_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.join_style */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fjoin_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_STYLE) != 0 {
		/* xcb_create_gc_value_list_t.fill_style */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ffill_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_RULE) != 0 {
		/* xcb_create_gc_value_list_t.fill_rule */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ffill_rule = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE) != 0 {
		/* xcb_create_gc_value_list_t.tile */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ftile = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_STIPPLE) != 0 {
		/* xcb_create_gc_value_list_t.stipple */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fstipple = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_X) != 0 {
		/* xcb_create_gc_value_list_t.tile_stipple_x_origin */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ftile_stipple_x_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_Y) != 0 {
		/* xcb_create_gc_value_list_t.tile_stipple_y_origin */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ftile_stipple_y_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FONT) != 0 {
		/* xcb_create_gc_value_list_t.font */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Ffont = *(*Txcb_font_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_SUBWINDOW_MODE) != 0 {
		/* xcb_create_gc_value_list_t.subwindow_mode */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fsubwindow_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_GRAPHICS_EXPOSURES) != 0 {
		/* xcb_create_gc_value_list_t.graphics_exposures */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fgraphics_exposures = *(*Txcb_bool32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_X) != 0 {
		/* xcb_create_gc_value_list_t.clip_x_origin */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fclip_x_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_Y) != 0 {
		/* xcb_create_gc_value_list_t.clip_y_origin */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fclip_y_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_MASK) != 0 {
		/* xcb_create_gc_value_list_t.clip_mask */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fclip_mask = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_OFFSET) != 0 {
		/* xcb_create_gc_value_list_t.dash_offset */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fdash_offset = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_LIST) != 0 {
		/* xcb_create_gc_value_list_t.dashes */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Fdashes = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_ARC_MODE) != 0 {
		/* xcb_create_gc_value_list_t.arc_mode */
		(*Txcb_create_gc_value_list_t)(unsafe.Pointer(_aux)).Farc_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_create_gc_value_list_sizeof(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* _aux at bp+0 */ Txcb_create_gc_value_list_t
	return Xxcb_create_gc_value_list_unpack(tls, _buffer, value_mask, bp)
}

func Xxcb_create_gc_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value_list */
	xcb_block_len += libc.Uint32FromInt32(Xxcb_create_gc_value_list_sizeof(tls, xcb_tmp, (*Txcb_create_gc_request_t)(unsafe.Pointer(_aux)).Fvalue_mask))
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_create_gc_checked(tls *libc.TLS, c uintptr, cid Txcb_gcontext_t, drawable Txcb_drawable_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_create_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fcid = cid
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fdrawable = drawable
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_gc_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req114)))
	return xcb_ret
}

var _xcb_req114 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_create_gc(tls *libc.TLS, c uintptr, cid Txcb_gcontext_t, drawable Txcb_drawable_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_create_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fcid = cid
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fdrawable = drawable
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_gc_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req115)))
	return xcb_ret
}

var _xcb_req115 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_create_gc_aux_checked(tls *libc.TLS, c uintptr, cid Txcb_gcontext_t, drawable Txcb_drawable_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+56 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_create_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 56)) = uintptr(0)
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fcid = cid
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fdrawable = drawable
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_gc_value_list_serialize(tls, bp+56, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 56))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req116)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 56)))
	return xcb_ret
}

var _xcb_req116 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_create_gc_aux(tls *libc.TLS, c uintptr, cid Txcb_gcontext_t, drawable Txcb_drawable_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+56 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_create_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 56)) = uintptr(0)
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fcid = cid
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fdrawable = drawable
	(*(*Txcb_create_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_create_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_create_gc_value_list_serialize(tls, bp+56, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 56))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req117)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 56)))
	return xcb_ret
}

var _xcb_req117 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CREATE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_create_gc_value_list(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_change_gc_value_list_serialize(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_parts_idx uint32
	var xcb_out, xcb_tmp uintptr
	var xcb_parts [24]Tiovec
	var _ /* xcb_pad0 at bp+0 */ [3]int8
	_, _, _, _, _, _, _, _, _, _ = i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_out, xcb_pad, xcb_padding_offset, xcb_parts, xcb_parts_idx, xcb_tmp
	xcb_out = *(*uintptr)(unsafe.Pointer(_buffer))
	xcb_buffer_len = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	xcb_pad = uint32(0)
	*(*[3]int8)(unsafe.Pointer(bp)) = [3]int8{}
	xcb_parts_idx = uint32(0)
	xcb_block_len = uint32(0)
	if value_mask&uint32(_XCB_GC_FUNCTION) != 0 {
		/* xcb_change_gc_value_list_t.function */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_PLANE_MASK) != 0 {
		/* xcb_change_gc_value_list_t.plane_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 4
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FOREGROUND) != 0 {
		/* xcb_change_gc_value_list_t.foreground */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 8
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_BACKGROUND) != 0 {
		/* xcb_change_gc_value_list_t.background */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 12
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_WIDTH) != 0 {
		/* xcb_change_gc_value_list_t.line_width */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 16
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.line_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 20
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CAP_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.cap_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 24
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_JOIN_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.join_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 28
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.fill_style */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 32
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_RULE) != 0 {
		/* xcb_change_gc_value_list_t.fill_rule */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 36
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE) != 0 {
		/* xcb_change_gc_value_list_t.tile */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 40
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_STIPPLE) != 0 {
		/* xcb_change_gc_value_list_t.stipple */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 44
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_X) != 0 {
		/* xcb_change_gc_value_list_t.tile_stipple_x_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 48
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_Y) != 0 {
		/* xcb_change_gc_value_list_t.tile_stipple_y_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 52
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FONT) != 0 {
		/* xcb_change_gc_value_list_t.font */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 56
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_SUBWINDOW_MODE) != 0 {
		/* xcb_change_gc_value_list_t.subwindow_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 60
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_GRAPHICS_EXPOSURES) != 0 {
		/* xcb_change_gc_value_list_t.graphics_exposures */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 64
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_X) != 0 {
		/* xcb_change_gc_value_list_t.clip_x_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 68
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_Y) != 0 {
		/* xcb_change_gc_value_list_t.clip_y_origin */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 72
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_MASK) != 0 {
		/* xcb_change_gc_value_list_t.clip_mask */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 76
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_OFFSET) != 0 {
		/* xcb_change_gc_value_list_t.dash_offset */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 80
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_LIST) != 0 {
		/* xcb_change_gc_value_list_t.dashes */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 84
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_ARC_MODE) != 0 {
		/* xcb_change_gc_value_list_t.arc_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 88
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_parts[xcb_parts_idx].Fiov_base = bp
		xcb_parts[xcb_parts_idx].Fiov_len = xcb_pad
		xcb_parts_idx++
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.UintptrFromInt32(0) == xcb_out {
		/* allocate memory */
		xcb_out = libc.Xmalloc(tls, xcb_buffer_len)
		*(*uintptr)(unsafe.Pointer(_buffer)) = xcb_out
	}
	xcb_tmp = xcb_out
	i = uint32(0)
	for {
		if !(i < xcb_parts_idx) {
			break
		}
		if uintptr(0) != xcb_parts[i].Fiov_base && uint32(0) != xcb_parts[i].Fiov_len {
			libc.Xmemcpy(tls, xcb_tmp, xcb_parts[i].Fiov_base, xcb_parts[i].Fiov_len)
		}
		if uint32(0) != xcb_parts[i].Fiov_len {
			xcb_tmp += uintptr(xcb_parts[i].Fiov_len)
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_gc_value_list_unpack(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset uint32
	var xcb_tmp uintptr
	_, _, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	if value_mask&uint32(_XCB_GC_FUNCTION) != 0 {
		/* xcb_change_gc_value_list_t.function */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ffunction = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_PLANE_MASK) != 0 {
		/* xcb_change_gc_value_list_t.plane_mask */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fplane_mask = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FOREGROUND) != 0 {
		/* xcb_change_gc_value_list_t.foreground */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fforeground = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_BACKGROUND) != 0 {
		/* xcb_change_gc_value_list_t.background */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fbackground = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_WIDTH) != 0 {
		/* xcb_change_gc_value_list_t.line_width */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fline_width = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_LINE_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.line_style */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fline_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CAP_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.cap_style */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fcap_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_JOIN_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.join_style */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fjoin_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_STYLE) != 0 {
		/* xcb_change_gc_value_list_t.fill_style */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ffill_style = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FILL_RULE) != 0 {
		/* xcb_change_gc_value_list_t.fill_rule */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ffill_rule = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE) != 0 {
		/* xcb_change_gc_value_list_t.tile */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ftile = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_STIPPLE) != 0 {
		/* xcb_change_gc_value_list_t.stipple */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fstipple = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_X) != 0 {
		/* xcb_change_gc_value_list_t.tile_stipple_x_origin */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ftile_stipple_x_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_TILE_STIPPLE_ORIGIN_Y) != 0 {
		/* xcb_change_gc_value_list_t.tile_stipple_y_origin */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ftile_stipple_y_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_FONT) != 0 {
		/* xcb_change_gc_value_list_t.font */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Ffont = *(*Txcb_font_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_SUBWINDOW_MODE) != 0 {
		/* xcb_change_gc_value_list_t.subwindow_mode */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fsubwindow_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_GRAPHICS_EXPOSURES) != 0 {
		/* xcb_change_gc_value_list_t.graphics_exposures */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fgraphics_exposures = *(*Txcb_bool32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_X) != 0 {
		/* xcb_change_gc_value_list_t.clip_x_origin */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fclip_x_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_ORIGIN_Y) != 0 {
		/* xcb_change_gc_value_list_t.clip_y_origin */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fclip_y_origin = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_CLIP_MASK) != 0 {
		/* xcb_change_gc_value_list_t.clip_mask */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fclip_mask = *(*Txcb_pixmap_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_OFFSET) != 0 {
		/* xcb_change_gc_value_list_t.dash_offset */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fdash_offset = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_DASH_LIST) != 0 {
		/* xcb_change_gc_value_list_t.dashes */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Fdashes = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_GC_ARC_MODE) != 0 {
		/* xcb_change_gc_value_list_t.arc_mode */
		(*Txcb_change_gc_value_list_t)(unsafe.Pointer(_aux)).Farc_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_gc_value_list_sizeof(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* _aux at bp+0 */ Txcb_change_gc_value_list_t
	return Xxcb_change_gc_value_list_unpack(tls, _buffer, value_mask, bp)
}

func Xxcb_change_gc_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value_list */
	xcb_block_len += libc.Uint32FromInt32(Xxcb_change_gc_value_list_sizeof(tls, xcb_tmp, (*Txcb_change_gc_request_t)(unsafe.Pointer(_aux)).Fvalue_mask))
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_gc_checked(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_change_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fgc = gc
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_gc_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req118)))
	return xcb_ret
}

var _xcb_req118 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_change_gc(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_change_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fgc = gc
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_gc_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req119)))
	return xcb_ret
}

var _xcb_req119 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_change_gc_aux_checked(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+52 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_change_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 52)) = uintptr(0)
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fgc = gc
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_gc_value_list_serialize(tls, bp+52, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 52))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req120)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 52)))
	return xcb_ret
}

var _xcb_req120 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_change_gc_aux(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+52 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_change_gc_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 52)) = uintptr(0)
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fgc = gc
	(*(*Txcb_change_gc_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_gc_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_gc_value_list_serialize(tls, bp+52, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 52))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req121)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 52)))
	return xcb_ret
}

var _xcb_req121 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_change_gc_value_list(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_copy_gc_checked(tls *libc.TLS, c uintptr, src_gc Txcb_gcontext_t, dst_gc Txcb_gcontext_t, value_mask Tuint32_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_gc_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fsrc_gc = src_gc
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fdst_gc = dst_gc
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fvalue_mask = value_mask
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req122)))
	return xcb_ret
}

var _xcb_req122 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_GC),
	Fisvoid: uint8(1),
}

func Xxcb_copy_gc(tls *libc.TLS, c uintptr, src_gc Txcb_gcontext_t, dst_gc Txcb_gcontext_t, value_mask Tuint32_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_gc_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fsrc_gc = src_gc
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fdst_gc = dst_gc
	(*(*Txcb_copy_gc_request_t)(unsafe.Pointer(bp + 32))).Fvalue_mask = value_mask
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req123)))
	return xcb_ret
}

var _xcb_req123 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_GC),
	Fisvoid: uint8(1),
}

func Xxcb_set_dashes_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* dashes */
	xcb_block_len += uint32((*Txcb_set_dashes_request_t)(unsafe.Pointer(_aux)).Fdashes_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_set_dashes_checked(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t, dash_offset Tuint16_t, dashes_len Tuint16_t, dashes uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_dashes_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fdash_offset = dash_offset
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fdashes_len = dashes_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t dashes */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = dashes
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(dashes_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req124)))
	return xcb_ret
}

var _xcb_req124 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_DASHES),
	Fisvoid: uint8(1),
}

func Xxcb_set_dashes(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t, dash_offset Tuint16_t, dashes_len Tuint16_t, dashes uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_dashes_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fdash_offset = dash_offset
	(*(*Txcb_set_dashes_request_t)(unsafe.Pointer(bp + 48))).Fdashes_len = dashes_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t dashes */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = dashes
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(dashes_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req125)))
	return xcb_ret
}

var _xcb_req125 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_DASHES),
	Fisvoid: uint8(1),
}

func Xxcb_set_dashes_dashes(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_set_dashes_dashes_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_set_dashes_request_t)(unsafe.Pointer(R)).Fdashes_len)
}

func Xxcb_set_dashes_dashes_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12 + uintptr((*Txcb_set_dashes_request_t)(unsafe.Pointer(R)).Fdashes_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_set_clip_rectangles_sizeof(tls *libc.TLS, _buffer uintptr, rectangles_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* rectangles */
	xcb_block_len += rectangles_len * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_set_clip_rectangles_checked(tls *libc.TLS, c uintptr, ordering Tuint8_t, gc Txcb_gcontext_t, clip_x_origin Tint16_t, clip_y_origin Tint16_t, rectangles_len Tuint32_t, rectangles uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_clip_rectangles_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fordering = ordering
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fclip_x_origin = clip_x_origin
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fclip_y_origin = clip_y_origin
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_rectangle_t rectangles */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = rectangles
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = rectangles_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req126)))
	return xcb_ret
}

var _xcb_req126 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_CLIP_RECTANGLES),
	Fisvoid: uint8(1),
}

func Xxcb_set_clip_rectangles(tls *libc.TLS, c uintptr, ordering Tuint8_t, gc Txcb_gcontext_t, clip_x_origin Tint16_t, clip_y_origin Tint16_t, rectangles_len Tuint32_t, rectangles uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_clip_rectangles_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fordering = ordering
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fclip_x_origin = clip_x_origin
	(*(*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(bp + 48))).Fclip_y_origin = clip_y_origin
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_rectangle_t rectangles */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = rectangles
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = rectangles_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req127)))
	return xcb_ret
}

var _xcb_req127 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_CLIP_RECTANGLES),
	Fisvoid: uint8(1),
}

func Xxcb_set_clip_rectangles_rectangles(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_set_clip_rectangles_rectangles_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
}

func Xxcb_set_clip_rectangles_rectangles_iterator(tls *libc.TLS, R uintptr) (r Txcb_rectangle_iterator_t) {
	var i Txcb_rectangle_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_set_clip_rectangles_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_free_gc_checked(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_gc_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_gc_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_gc_request_t)(unsafe.Pointer(bp + 32))).Fgc = gc
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req128)))
	return xcb_ret
}

var _xcb_req128 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_free_gc(tls *libc.TLS, c uintptr, gc Txcb_gcontext_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_gc_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_gc_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_gc_request_t)(unsafe.Pointer(bp + 32))).Fgc = gc
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req129)))
	return xcb_ret
}

var _xcb_req129 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_GC),
	Fisvoid: uint8(1),
}

func Xxcb_clear_area_checked(tls *libc.TLS, c uintptr, exposures Tuint8_t, window Txcb_window_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_clear_area_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fexposures = exposures
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req130)))
	return xcb_ret
}

var _xcb_req130 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CLEAR_AREA),
	Fisvoid: uint8(1),
}

func Xxcb_clear_area(tls *libc.TLS, c uintptr, exposures Tuint8_t, window Txcb_window_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_clear_area_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fexposures = exposures
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_clear_area_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req131)))
	return xcb_ret
}

var _xcb_req131 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CLEAR_AREA),
	Fisvoid: uint8(1),
}

func Xxcb_copy_area_checked(tls *libc.TLS, c uintptr, src_drawable Txcb_drawable_t, dst_drawable Txcb_drawable_t, gc Txcb_gcontext_t, src_x Tint16_t, src_y Tint16_t, dst_x Tint16_t, dst_y Tint16_t, width Tuint16_t, height Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_area_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fsrc_drawable = src_drawable
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fdst_drawable = dst_drawable
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fgc = gc
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fdst_x = dst_x
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fdst_y = dst_y
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(28)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req132)))
	return xcb_ret
}

var _xcb_req132 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_AREA),
	Fisvoid: uint8(1),
}

func Xxcb_copy_area(tls *libc.TLS, c uintptr, src_drawable Txcb_drawable_t, dst_drawable Txcb_drawable_t, gc Txcb_gcontext_t, src_x Tint16_t, src_y Tint16_t, dst_x Tint16_t, dst_y Tint16_t, width Tuint16_t, height Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_area_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fsrc_drawable = src_drawable
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fdst_drawable = dst_drawable
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fgc = gc
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fdst_x = dst_x
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fdst_y = dst_y
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_copy_area_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(28)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req133)))
	return xcb_ret
}

var _xcb_req133 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_AREA),
	Fisvoid: uint8(1),
}

func Xxcb_copy_plane_checked(tls *libc.TLS, c uintptr, src_drawable Txcb_drawable_t, dst_drawable Txcb_drawable_t, gc Txcb_gcontext_t, src_x Tint16_t, src_y Tint16_t, dst_x Tint16_t, dst_y Tint16_t, width Tuint16_t, height Tuint16_t, bit_plane Tuint32_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_plane_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fsrc_drawable = src_drawable
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fdst_drawable = dst_drawable
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fgc = gc
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fdst_x = dst_x
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fdst_y = dst_y
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fbit_plane = bit_plane
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req134)))
	return xcb_ret
}

var _xcb_req134 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_PLANE),
	Fisvoid: uint8(1),
}

func Xxcb_copy_plane(tls *libc.TLS, c uintptr, src_drawable Txcb_drawable_t, dst_drawable Txcb_drawable_t, gc Txcb_gcontext_t, src_x Tint16_t, src_y Tint16_t, dst_x Tint16_t, dst_y Tint16_t, width Tuint16_t, height Tuint16_t, bit_plane Tuint32_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_plane_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fsrc_drawable = src_drawable
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fdst_drawable = dst_drawable
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fgc = gc
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fsrc_x = src_x
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fsrc_y = src_y
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fdst_x = dst_x
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fdst_y = dst_y
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*Txcb_copy_plane_request_t)(unsafe.Pointer(bp + 32))).Fbit_plane = bit_plane
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req135)))
	return xcb_ret
}

var _xcb_req135 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_PLANE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_point_sizeof(tls *libc.TLS, _buffer uintptr, points_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* points */
	xcb_block_len += points_len * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_point_checked(tls *libc.TLS, c uintptr, coordinate_mode Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, points_len Tuint32_t, points uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_point_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_point_request_t)(unsafe.Pointer(bp + 48))).Fcoordinate_mode = coordinate_mode
	(*(*Txcb_poly_point_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_point_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_point_t points */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = points
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = points_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req136)))
	return xcb_ret
}

var _xcb_req136 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_POINT),
	Fisvoid: uint8(1),
}

func Xxcb_poly_point(tls *libc.TLS, c uintptr, coordinate_mode Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, points_len Tuint32_t, points uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_point_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_point_request_t)(unsafe.Pointer(bp + 48))).Fcoordinate_mode = coordinate_mode
	(*(*Txcb_poly_point_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_point_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_point_t points */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = points
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = points_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req137)))
	return xcb_ret
}

var _xcb_req137 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_POINT),
	Fisvoid: uint8(1),
}

func Xxcb_poly_point_points(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_point_points_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_point_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(4))
}

func Xxcb_poly_point_points_iterator(tls *libc.TLS, R uintptr) (r Txcb_point_iterator_t) {
	var i Txcb_point_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_point_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(4))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_poly_line_sizeof(tls *libc.TLS, _buffer uintptr, points_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* points */
	xcb_block_len += points_len * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_line_checked(tls *libc.TLS, c uintptr, coordinate_mode Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, points_len Tuint32_t, points uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_line_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_line_request_t)(unsafe.Pointer(bp + 48))).Fcoordinate_mode = coordinate_mode
	(*(*Txcb_poly_line_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_line_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_point_t points */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = points
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = points_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req138)))
	return xcb_ret
}

var _xcb_req138 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_LINE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_line(tls *libc.TLS, c uintptr, coordinate_mode Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, points_len Tuint32_t, points uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_line_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_line_request_t)(unsafe.Pointer(bp + 48))).Fcoordinate_mode = coordinate_mode
	(*(*Txcb_poly_line_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_line_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_point_t points */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = points
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = points_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req139)))
	return xcb_ret
}

var _xcb_req139 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_LINE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_line_points(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_line_points_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_line_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(4))
}

func Xxcb_poly_line_points_iterator(tls *libc.TLS, R uintptr) (r Txcb_point_iterator_t) {
	var i Txcb_point_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_line_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(4))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_segment_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_segment_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_segment_iterator_t)(unsafe.Pointer(i)).Fdata += 8
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(8))
}

func Xxcb_segment_end(tls *libc.TLS, i Txcb_segment_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*8
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_poly_segment_sizeof(tls *libc.TLS, _buffer uintptr, segments_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* segments */
	xcb_block_len += segments_len * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_segment_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, segments_len Tuint32_t, segments uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_segment_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_segment_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_segment_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_segment_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_segment_t segments */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = segments
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = segments_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req140)))
	return xcb_ret
}

var _xcb_req140 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_SEGMENT),
	Fisvoid: uint8(1),
}

func Xxcb_poly_segment(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, segments_len Tuint32_t, segments uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_segment_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_segment_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_segment_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_segment_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_segment_t segments */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = segments
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = segments_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req141)))
	return xcb_ret
}

var _xcb_req141 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_SEGMENT),
	Fisvoid: uint8(1),
}

func Xxcb_poly_segment_segments(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_segment_segments_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_segment_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
}

func Xxcb_poly_segment_segments_iterator(tls *libc.TLS, R uintptr) (r Txcb_segment_iterator_t) {
	var i Txcb_segment_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_segment_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_poly_rectangle_sizeof(tls *libc.TLS, _buffer uintptr, rectangles_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* rectangles */
	xcb_block_len += rectangles_len * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_rectangle_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, rectangles_len Tuint32_t, rectangles uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_rectangle_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_rectangle_t rectangles */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = rectangles
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = rectangles_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req142)))
	return xcb_ret
}

var _xcb_req142 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_RECTANGLE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_rectangle(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, rectangles_len Tuint32_t, rectangles uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_rectangle_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_rectangle_t rectangles */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = rectangles
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = rectangles_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req143)))
	return xcb_ret
}

var _xcb_req143 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_RECTANGLE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_rectangle_rectangles(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_rectangle_rectangles_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_rectangle_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
}

func Xxcb_poly_rectangle_rectangles_iterator(tls *libc.TLS, R uintptr) (r Txcb_rectangle_iterator_t) {
	var i Txcb_rectangle_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_rectangle_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_poly_arc_sizeof(tls *libc.TLS, _buffer uintptr, arcs_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* arcs */
	xcb_block_len += arcs_len * uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_arc_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, arcs_len Tuint32_t, arcs uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_arc_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_arc_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_arc_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_arc_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_arc_t arcs */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = arcs
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = arcs_len * uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req144)))
	return xcb_ret
}

var _xcb_req144 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_ARC),
	Fisvoid: uint8(1),
}

func Xxcb_poly_arc(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, arcs_len Tuint32_t, arcs uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_arc_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_arc_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_arc_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_arc_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_arc_t arcs */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = arcs
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = arcs_len * uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req145)))
	return xcb_ret
}

var _xcb_req145 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_ARC),
	Fisvoid: uint8(1),
}

func Xxcb_poly_arc_arcs(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_arc_arcs_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_arc_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(12))
}

func Xxcb_poly_arc_arcs_iterator(tls *libc.TLS, R uintptr) (r Txcb_arc_iterator_t) {
	var i Txcb_arc_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_arc_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(12))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_fill_poly_sizeof(tls *libc.TLS, _buffer uintptr, points_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* points */
	xcb_block_len += points_len * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_fill_poly_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, shape Tuint8_t, coordinate_mode Tuint8_t, points_len Tuint32_t, points uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_fill_poly_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fshape = shape
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fcoordinate_mode = coordinate_mode
	libc.Xmemset(tls, bp+48+14, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_point_t points */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = points
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = points_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req146)))
	return xcb_ret
}

var _xcb_req146 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_FILL_POLY),
	Fisvoid: uint8(1),
}

func Xxcb_fill_poly(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, shape Tuint8_t, coordinate_mode Tuint8_t, points_len Tuint32_t, points uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_fill_poly_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fshape = shape
	(*(*Txcb_fill_poly_request_t)(unsafe.Pointer(bp + 48))).Fcoordinate_mode = coordinate_mode
	libc.Xmemset(tls, bp+48+14, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_point_t points */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = points
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = points_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req147)))
	return xcb_ret
}

var _xcb_req147 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_FILL_POLY),
	Fisvoid: uint8(1),
}

func Xxcb_fill_poly_points(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_fill_poly_points_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_fill_poly_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(16)) / libc.Uint32FromInt64(4))
}

func Xxcb_fill_poly_points_iterator(tls *libc.TLS, R uintptr) (r Txcb_point_iterator_t) {
	var i Txcb_point_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*16
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_fill_poly_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(16)) / libc.Uint32FromInt64(4))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_poly_fill_rectangle_sizeof(tls *libc.TLS, _buffer uintptr, rectangles_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* rectangles */
	xcb_block_len += rectangles_len * uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_fill_rectangle_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, rectangles_len Tuint32_t, rectangles uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_fill_rectangle_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_rectangle_t rectangles */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = rectangles
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = rectangles_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req148)))
	return xcb_ret
}

var _xcb_req148 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_FILL_RECTANGLE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_fill_rectangle(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, rectangles_len Tuint32_t, rectangles uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_fill_rectangle_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_rectangle_t rectangles */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = rectangles
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = rectangles_len * uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req149)))
	return xcb_ret
}

var _xcb_req149 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_FILL_RECTANGLE),
	Fisvoid: uint8(1),
}

func Xxcb_poly_fill_rectangle_rectangles(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_fill_rectangle_rectangles_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
}

func Xxcb_poly_fill_rectangle_rectangles_iterator(tls *libc.TLS, R uintptr) (r Txcb_rectangle_iterator_t) {
	var i Txcb_rectangle_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_fill_rectangle_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(8))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_poly_fill_arc_sizeof(tls *libc.TLS, _buffer uintptr, arcs_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* arcs */
	xcb_block_len += arcs_len * uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_fill_arc_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, arcs_len Tuint32_t, arcs uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_fill_arc_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_arc_t arcs */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = arcs
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = arcs_len * uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req150)))
	return xcb_ret
}

var _xcb_req150 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_FILL_ARC),
	Fisvoid: uint8(1),
}

func Xxcb_poly_fill_arc(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, arcs_len Tuint32_t, arcs uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_fill_arc_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_arc_t arcs */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = arcs
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = arcs_len * uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req151)))
	return xcb_ret
}

var _xcb_req151 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_FILL_ARC),
	Fisvoid: uint8(1),
}

func Xxcb_poly_fill_arc_arcs(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_poly_fill_arc_arcs_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(12))
}

func Xxcb_poly_fill_arc_arcs_iterator(tls *libc.TLS, R uintptr) (r Txcb_arc_iterator_t) {
	var i Txcb_arc_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_fill_arc_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(12))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_put_image_sizeof(tls *libc.TLS, _buffer uintptr, data_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(24)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* data */
	xcb_block_len += data_len * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_put_image_checked(tls *libc.TLS, c uintptr, format Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, width Tuint16_t, height Tuint16_t, dst_x Tint16_t, dst_y Tint16_t, left_pad Tuint8_t, depth Tuint8_t, data_len Tuint32_t, data uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_put_image_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fformat = format
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fwidth = width
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fheight = height
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdst_x = dst_x
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdst_y = dst_y
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fleft_pad = left_pad
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdepth = depth
	libc.Xmemset(tls, bp+48+22, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t data */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = data
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = data_len * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req152)))
	return xcb_ret
}

var _xcb_req152 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_PUT_IMAGE),
	Fisvoid: uint8(1),
}

func Xxcb_put_image(tls *libc.TLS, c uintptr, format Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, width Tuint16_t, height Tuint16_t, dst_x Tint16_t, dst_y Tint16_t, left_pad Tuint8_t, depth Tuint8_t, data_len Tuint32_t, data uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_put_image_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fformat = format
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fwidth = width
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fheight = height
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdst_x = dst_x
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdst_y = dst_y
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fleft_pad = left_pad
	(*(*Txcb_put_image_request_t)(unsafe.Pointer(bp + 48))).Fdepth = depth
	libc.Xmemset(tls, bp+48+22, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(24)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t data */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = data
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = data_len * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req153)))
	return xcb_ret
}

var _xcb_req153 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_PUT_IMAGE),
	Fisvoid: uint8(1),
}

func Xxcb_put_image_data(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*24
}

func Xxcb_put_image_data_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_put_image_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(24)) / libc.Uint32FromInt64(1))
}

func Xxcb_put_image_data_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*24 + uintptr((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_put_image_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4))-libc.Uint32FromInt64(24))/libc.Uint32FromInt64(1))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_image_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* data */
	xcb_block_len += (*Txcb_get_image_reply_t)(unsafe.Pointer(_aux)).Flength * uint32(4) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_image(tls *libc.TLS, c uintptr, format Tuint8_t, drawable Txcb_drawable_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t, plane_mask Tuint32_t) (r Txcb_get_image_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_get_image_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_image_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fformat = format
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fplane_mask = plane_mask
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(20)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req154)))
	return xcb_ret
}

var _xcb_req154 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_IMAGE),
}

func Xxcb_get_image_unchecked(tls *libc.TLS, c uintptr, format Tuint8_t, drawable Txcb_drawable_t, x Tint16_t, y Tint16_t, width Tuint16_t, height Tuint16_t, plane_mask Tuint32_t) (r Txcb_get_image_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_get_image_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_image_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fformat = format
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*Txcb_get_image_request_t)(unsafe.Pointer(bp + 32))).Fplane_mask = plane_mask
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(20)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req155)))
	return xcb_ret
}

var _xcb_req155 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_IMAGE),
}

func Xxcb_get_image_data(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_image_data_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_get_image_reply_t)(unsafe.Pointer(R)).Flength * libc.Uint32FromInt32(4))
}

func Xxcb_get_image_data_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_get_image_reply_t)(unsafe.Pointer(R)).Flength*libc.Uint32FromInt32(4))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_image_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_image_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_poly_text_8_sizeof(tls *libc.TLS, _buffer uintptr, items_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* items */
	xcb_block_len += items_len * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_text_8_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, items_len Tuint32_t, items uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_text_8_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t items */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = items
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = items_len * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req156)))
	return xcb_ret
}

var _xcb_req156 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_TEXT_8),
	Fisvoid: uint8(1),
}

func Xxcb_poly_text_8(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, items_len Tuint32_t, items uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_text_8_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_poly_text_8_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t items */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = items
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = items_len * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req157)))
	return xcb_ret
}

var _xcb_req157 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_TEXT_8),
	Fisvoid: uint8(1),
}

func Xxcb_poly_text_8_items(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_poly_text_8_items_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_text_8_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(16)) / libc.Uint32FromInt64(1))
}

func Xxcb_poly_text_8_items_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*16 + uintptr((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_text_8_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4))-libc.Uint32FromInt64(16))/libc.Uint32FromInt64(1))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_poly_text_16_sizeof(tls *libc.TLS, _buffer uintptr, items_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* items */
	xcb_block_len += items_len * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_poly_text_16_checked(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, items_len Tuint32_t, items uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_text_16_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t items */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = items
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = items_len * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req158)))
	return xcb_ret
}

var _xcb_req158 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_TEXT_16),
	Fisvoid: uint8(1),
}

func Xxcb_poly_text_16(tls *libc.TLS, c uintptr, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, items_len Tuint32_t, items uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_poly_text_16_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_poly_text_16_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t items */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = items
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = items_len * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req159)))
	return xcb_ret
}

var _xcb_req159 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_POLY_TEXT_16),
	Fisvoid: uint8(1),
}

func Xxcb_poly_text_16_items(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_poly_text_16_items_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_text_16_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(16)) / libc.Uint32FromInt64(1))
}

func Xxcb_poly_text_16_items_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*16 + uintptr((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_poly_text_16_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4))-libc.Uint32FromInt64(16))/libc.Uint32FromInt64(1))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_image_text_8_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* string */
	xcb_block_len += uint32((*Txcb_image_text_8_request_t)(unsafe.Pointer(_aux)).Fstring_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_image_text_8_checked(tls *libc.TLS, c uintptr, string_len Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, string1 uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_image_text_8_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fstring_len = string_len
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char string */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = string1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(string_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req160)))
	return xcb_ret
}

var _xcb_req160 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_IMAGE_TEXT_8),
	Fisvoid: uint8(1),
}

func Xxcb_image_text_8(tls *libc.TLS, c uintptr, string_len Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, string1 uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_image_text_8_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fstring_len = string_len
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_image_text_8_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char string */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = string1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(string_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req161)))
	return xcb_ret
}

var _xcb_req161 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_IMAGE_TEXT_8),
	Fisvoid: uint8(1),
}

func Xxcb_image_text_8_string(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_image_text_8_string_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_image_text_8_request_t)(unsafe.Pointer(R)).Fstring_len)
}

func Xxcb_image_text_8_string_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*16 + uintptr((*Txcb_image_text_8_request_t)(unsafe.Pointer(R)).Fstring_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_image_text_16_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* string */
	xcb_block_len += uint32((*Txcb_image_text_16_request_t)(unsafe.Pointer(_aux)).Fstring_len) * uint32(2)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_image_text_16_checked(tls *libc.TLS, c uintptr, string_len Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, string1 uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_image_text_16_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fstring_len = string_len
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_char2b_t string */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = string1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(string_len) * uint32(2)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req162)))
	return xcb_ret
}

var _xcb_req162 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_IMAGE_TEXT_16),
	Fisvoid: uint8(1),
}

func Xxcb_image_text_16(tls *libc.TLS, c uintptr, string_len Tuint8_t, drawable Txcb_drawable_t, gc Txcb_gcontext_t, x Tint16_t, y Tint16_t, string1 uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_image_text_16_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fstring_len = string_len
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fdrawable = drawable
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fgc = gc
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fx = x
	(*(*Txcb_image_text_16_request_t)(unsafe.Pointer(bp + 48))).Fy = y
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_char2b_t string */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = string1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(string_len) * uint32(2)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req163)))
	return xcb_ret
}

var _xcb_req163 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_IMAGE_TEXT_16),
	Fisvoid: uint8(1),
}

func Xxcb_image_text_16_string(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_image_text_16_string_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_image_text_16_request_t)(unsafe.Pointer(R)).Fstring_len)
}

func Xxcb_image_text_16_string_iterator(tls *libc.TLS, R uintptr) (r Txcb_char2b_iterator_t) {
	var i Txcb_char2b_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*16
	i.Frem = libc.Int32FromUint8((*Txcb_image_text_16_request_t)(unsafe.Pointer(R)).Fstring_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_create_colormap_checked(tls *libc.TLS, c uintptr, alloc Tuint8_t, mid Txcb_colormap_t, window Txcb_window_t, visual Txcb_visualid_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Falloc = alloc
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Fmid = mid
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Fvisual = visual
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req164)))
	return xcb_ret
}

var _xcb_req164 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_create_colormap(tls *libc.TLS, c uintptr, alloc Tuint8_t, mid Txcb_colormap_t, window Txcb_window_t, visual Txcb_visualid_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Falloc = alloc
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Fmid = mid
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*Txcb_create_colormap_request_t)(unsafe.Pointer(bp + 32))).Fvisual = visual
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req165)))
	return xcb_ret
}

var _xcb_req165 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_free_colormap_checked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_colormap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_colormap_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req166)))
	return xcb_ret
}

var _xcb_req166 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_free_colormap(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_colormap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_colormap_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req167)))
	return xcb_ret
}

var _xcb_req167 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_copy_colormap_and_free_checked(tls *libc.TLS, c uintptr, mid Txcb_colormap_t, src_cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_colormap_and_free_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_colormap_and_free_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_colormap_and_free_request_t)(unsafe.Pointer(bp + 32))).Fmid = mid
	(*(*Txcb_copy_colormap_and_free_request_t)(unsafe.Pointer(bp + 32))).Fsrc_cmap = src_cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req168)))
	return xcb_ret
}

var _xcb_req168 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_COLORMAP_AND_FREE),
	Fisvoid: uint8(1),
}

func Xxcb_copy_colormap_and_free(tls *libc.TLS, c uintptr, mid Txcb_colormap_t, src_cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_copy_colormap_and_free_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_copy_colormap_and_free_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_copy_colormap_and_free_request_t)(unsafe.Pointer(bp + 32))).Fmid = mid
	(*(*Txcb_copy_colormap_and_free_request_t)(unsafe.Pointer(bp + 32))).Fsrc_cmap = src_cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req169)))
	return xcb_ret
}

var _xcb_req169 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_COPY_COLORMAP_AND_FREE),
	Fisvoid: uint8(1),
}

func Xxcb_install_colormap_checked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_install_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_install_colormap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_install_colormap_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req170)))
	return xcb_ret
}

var _xcb_req170 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_INSTALL_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_install_colormap(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_install_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_install_colormap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_install_colormap_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req171)))
	return xcb_ret
}

var _xcb_req171 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_INSTALL_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_uninstall_colormap_checked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_uninstall_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_uninstall_colormap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_uninstall_colormap_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req172)))
	return xcb_ret
}

var _xcb_req172 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNINSTALL_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_uninstall_colormap(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_uninstall_colormap_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_uninstall_colormap_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_uninstall_colormap_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req173)))
	return xcb_ret
}

var _xcb_req173 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_UNINSTALL_COLORMAP),
	Fisvoid: uint8(1),
}

func Xxcb_list_installed_colormaps_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* cmaps */
	xcb_block_len += uint32((*Txcb_list_installed_colormaps_reply_t)(unsafe.Pointer(_aux)).Fcmaps_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_list_installed_colormaps(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_list_installed_colormaps_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_installed_colormaps_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_installed_colormaps_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_installed_colormaps_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_list_installed_colormaps_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req174)))
	return xcb_ret
}

var _xcb_req174 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_INSTALLED_COLORMAPS),
}

func Xxcb_list_installed_colormaps_unchecked(tls *libc.TLS, c uintptr, window Txcb_window_t) (r Txcb_list_installed_colormaps_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_installed_colormaps_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_installed_colormaps_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_installed_colormaps_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_list_installed_colormaps_request_t)(unsafe.Pointer(bp + 32))).Fwindow = window
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req175)))
	return xcb_ret
}

var _xcb_req175 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_INSTALLED_COLORMAPS),
}

func Xxcb_list_installed_colormaps_cmaps(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_list_installed_colormaps_cmaps_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_list_installed_colormaps_reply_t)(unsafe.Pointer(R)).Fcmaps_len)
}

func Xxcb_list_installed_colormaps_cmaps_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_list_installed_colormaps_reply_t)(unsafe.Pointer(R)).Fcmaps_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_installed_colormaps_reply(tls *libc.TLS, c uintptr, cookie Txcb_list_installed_colormaps_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_alloc_color(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, red Tuint16_t, green Tuint16_t, blue Tuint16_t) (r Txcb_alloc_color_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_alloc_color_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_alloc_color_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fred = red
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fgreen = green
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fblue = blue
	libc.Xmemset(tls, bp+32+14, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req176)))
	return xcb_ret
}

var _xcb_req176 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOC_COLOR),
}

func Xxcb_alloc_color_unchecked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, red Tuint16_t, green Tuint16_t, blue Tuint16_t) (r Txcb_alloc_color_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_alloc_color_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_alloc_color_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fred = red
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fgreen = green
	(*(*Txcb_alloc_color_request_t)(unsafe.Pointer(bp + 32))).Fblue = blue
	libc.Xmemset(tls, bp+32+14, 0, uint32(2))
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req177)))
	return xcb_ret
}

var _xcb_req177 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOC_COLOR),
}

func Xxcb_alloc_color_reply(tls *libc.TLS, c uintptr, cookie Txcb_alloc_color_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_alloc_named_color_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_alloc_named_color_request_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_alloc_named_color(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, name_len Tuint16_t, name uintptr) (r Txcb_alloc_named_color_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_alloc_named_color_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_alloc_named_color_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_named_color_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_alloc_named_color_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_alloc_named_color_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+10, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req178)))
	return xcb_ret
}

var _xcb_req178 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_ALLOC_NAMED_COLOR),
}

func Xxcb_alloc_named_color_unchecked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, name_len Tuint16_t, name uintptr) (r Txcb_alloc_named_color_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_alloc_named_color_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_alloc_named_color_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_named_color_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_alloc_named_color_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_alloc_named_color_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+10, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req179)))
	return xcb_ret
}

var _xcb_req179 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_ALLOC_NAMED_COLOR),
}

func Xxcb_alloc_named_color_reply(tls *libc.TLS, c uintptr, cookie Txcb_alloc_named_color_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_alloc_color_cells_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* pixels */
	xcb_block_len += uint32((*Txcb_alloc_color_cells_reply_t)(unsafe.Pointer(_aux)).Fpixels_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* masks */
	xcb_block_len += uint32((*Txcb_alloc_color_cells_reply_t)(unsafe.Pointer(_aux)).Fmasks_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_alloc_color_cells(tls *libc.TLS, c uintptr, contiguous Tuint8_t, cmap Txcb_colormap_t, colors Tuint16_t, planes Tuint16_t) (r Txcb_alloc_color_cells_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_alloc_color_cells_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_alloc_color_cells_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fcontiguous = contiguous
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fcolors = colors
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fplanes = planes
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req180)))
	return xcb_ret
}

var _xcb_req180 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOC_COLOR_CELLS),
}

func Xxcb_alloc_color_cells_unchecked(tls *libc.TLS, c uintptr, contiguous Tuint8_t, cmap Txcb_colormap_t, colors Tuint16_t, planes Tuint16_t) (r Txcb_alloc_color_cells_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_alloc_color_cells_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_alloc_color_cells_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fcontiguous = contiguous
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fcolors = colors
	(*(*Txcb_alloc_color_cells_request_t)(unsafe.Pointer(bp + 32))).Fplanes = planes
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req181)))
	return xcb_ret
}

var _xcb_req181 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOC_COLOR_CELLS),
}

func Xxcb_alloc_color_cells_pixels(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_alloc_color_cells_pixels_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_alloc_color_cells_reply_t)(unsafe.Pointer(R)).Fpixels_len)
}

func Xxcb_alloc_color_cells_pixels_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_alloc_color_cells_reply_t)(unsafe.Pointer(R)).Fpixels_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_alloc_color_cells_masks(tls *libc.TLS, R uintptr) (r uintptr) {
	var prev Txcb_generic_iterator_t
	_ = prev
	prev = Xxcb_alloc_color_cells_pixels_end(tls, R)
	return prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1))) + libc.UintptrFromInt32(0)
}

func Xxcb_alloc_color_cells_masks_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_alloc_color_cells_reply_t)(unsafe.Pointer(R)).Fmasks_len)
}

func Xxcb_alloc_color_cells_masks_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i, prev Txcb_generic_iterator_t
	_, _ = i, prev
	prev = Xxcb_alloc_color_cells_pixels_end(tls, R)
	i.Fdata = prev.Fdata + uintptr(libc.Uint32FromInt32(-prev.Findex)&(libc.Uint32FromInt64(4)-libc.Uint32FromInt32(1))) + uintptr((*Txcb_alloc_color_cells_reply_t)(unsafe.Pointer(R)).Fmasks_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_alloc_color_cells_reply(tls *libc.TLS, c uintptr, cookie Txcb_alloc_color_cells_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_alloc_color_planes_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* pixels */
	xcb_block_len += uint32((*Txcb_alloc_color_planes_reply_t)(unsafe.Pointer(_aux)).Fpixels_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_alloc_color_planes(tls *libc.TLS, c uintptr, contiguous Tuint8_t, cmap Txcb_colormap_t, colors Tuint16_t, reds Tuint16_t, greens Tuint16_t, blues Tuint16_t) (r Txcb_alloc_color_planes_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_alloc_color_planes_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_alloc_color_planes_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fcontiguous = contiguous
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fcolors = colors
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Freds = reds
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fgreens = greens
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fblues = blues
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req182)))
	return xcb_ret
}

var _xcb_req182 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOC_COLOR_PLANES),
}

func Xxcb_alloc_color_planes_unchecked(tls *libc.TLS, c uintptr, contiguous Tuint8_t, cmap Txcb_colormap_t, colors Tuint16_t, reds Tuint16_t, greens Tuint16_t, blues Tuint16_t) (r Txcb_alloc_color_planes_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_alloc_color_planes_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_alloc_color_planes_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fcontiguous = contiguous
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fcmap = cmap
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fcolors = colors
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Freds = reds
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fgreens = greens
	(*(*Txcb_alloc_color_planes_request_t)(unsafe.Pointer(bp + 32))).Fblues = blues
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req183)))
	return xcb_ret
}

var _xcb_req183 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_ALLOC_COLOR_PLANES),
}

func Xxcb_alloc_color_planes_pixels(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_alloc_color_planes_pixels_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_alloc_color_planes_reply_t)(unsafe.Pointer(R)).Fpixels_len)
}

func Xxcb_alloc_color_planes_pixels_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_alloc_color_planes_reply_t)(unsafe.Pointer(R)).Fpixels_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_alloc_color_planes_reply(tls *libc.TLS, c uintptr, cookie Txcb_alloc_color_planes_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_free_colors_sizeof(tls *libc.TLS, _buffer uintptr, pixels_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* pixels */
	xcb_block_len += pixels_len * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_free_colors_checked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, plane_mask Tuint32_t, pixels_len Tuint32_t, pixels uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_free_colors_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_colors_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_free_colors_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_free_colors_request_t)(unsafe.Pointer(bp + 48))).Fplane_mask = plane_mask
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint32_t pixels */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pixels
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = pixels_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req184)))
	return xcb_ret
}

var _xcb_req184 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_FREE_COLORS),
	Fisvoid: uint8(1),
}

func Xxcb_free_colors(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, plane_mask Tuint32_t, pixels_len Tuint32_t, pixels uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_free_colors_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_colors_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_free_colors_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_free_colors_request_t)(unsafe.Pointer(bp + 48))).Fplane_mask = plane_mask
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint32_t pixels */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pixels
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = pixels_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req185)))
	return xcb_ret
}

var _xcb_req185 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_FREE_COLORS),
	Fisvoid: uint8(1),
}

func Xxcb_free_colors_pixels(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_free_colors_pixels_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_free_colors_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(12)) / libc.Uint32FromInt64(4))
}

func Xxcb_free_colors_pixels_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12 + uintptr((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_free_colors_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4))-libc.Uint32FromInt64(12))/libc.Uint32FromInt64(4))*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_coloritem_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_coloritem_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_coloritem_iterator_t)(unsafe.Pointer(i)).Fdata += 12
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(12))
}

func Xxcb_coloritem_end(tls *libc.TLS, i Txcb_coloritem_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*12
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_store_colors_sizeof(tls *libc.TLS, _buffer uintptr, items_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* items */
	xcb_block_len += items_len * uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_store_colors_checked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, items_len Tuint32_t, items uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_store_colors_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_store_colors_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_store_colors_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_coloritem_t items */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = items
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = items_len * uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req186)))
	return xcb_ret
}

var _xcb_req186 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_STORE_COLORS),
	Fisvoid: uint8(1),
}

func Xxcb_store_colors(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, items_len Tuint32_t, items uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_store_colors_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_store_colors_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_store_colors_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_coloritem_t items */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = items
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = items_len * uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req187)))
	return xcb_ret
}

var _xcb_req187 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_STORE_COLORS),
	Fisvoid: uint8(1),
}

func Xxcb_store_colors_items(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_store_colors_items_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_store_colors_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(8)) / libc.Uint32FromInt64(12))
}

func Xxcb_store_colors_items_iterator(tls *libc.TLS, R uintptr) (r Txcb_coloritem_iterator_t) {
	var i Txcb_coloritem_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8
	i.Frem = libc.Int32FromUint32((libc.Uint32FromInt32(libc.Int32FromUint16((*Txcb_store_colors_request_t)(unsafe.Pointer(R)).Flength)*libc.Int32FromInt32(4)) - libc.Uint32FromInt64(8)) / libc.Uint32FromInt64(12))
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_store_named_color_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(16)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_store_named_color_request_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_store_named_color_checked(tls *libc.TLS, c uintptr, flags Tuint8_t, cmap Txcb_colormap_t, pixel Tuint32_t, name_len Tuint16_t, name uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_store_named_color_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fflags = flags
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fpixel = pixel
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+14, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req188)))
	return xcb_ret
}

var _xcb_req188 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_STORE_NAMED_COLOR),
	Fisvoid: uint8(1),
}

func Xxcb_store_named_color(tls *libc.TLS, c uintptr, flags Tuint8_t, cmap Txcb_colormap_t, pixel Tuint32_t, name_len Tuint16_t, name uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_store_named_color_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fflags = flags
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fpixel = pixel
	(*(*Txcb_store_named_color_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+14, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(16)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req189)))
	return xcb_ret
}

var _xcb_req189 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_STORE_NAMED_COLOR),
	Fisvoid: uint8(1),
}

func Xxcb_store_named_color_name(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*16
}

func Xxcb_store_named_color_name_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_store_named_color_request_t)(unsafe.Pointer(R)).Fname_len)
}

func Xxcb_store_named_color_name_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*16 + uintptr((*Txcb_store_named_color_request_t)(unsafe.Pointer(R)).Fname_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_rgb_next(tls *libc.TLS, i uintptr) {
	var p1 uintptr
	_ = p1
	(*Txcb_rgb_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_rgb_iterator_t)(unsafe.Pointer(i)).Fdata += 8
	p1 = i + 8
	*(*int32)(unsafe.Pointer(p1)) = int32(uint32(*(*int32)(unsafe.Pointer(p1))) + libc.Uint32FromInt64(8))
}

func Xxcb_rgb_end(tls *libc.TLS, i Txcb_rgb_iterator_t) (r Txcb_generic_iterator_t) {
	var ret Txcb_generic_iterator_t
	_ = ret
	ret.Fdata = i.Fdata + uintptr(i.Frem)*8
	ret.Findex = i.Findex + (int32(ret.Fdata) - int32(i.Fdata))
	ret.Frem = 0
	return ret
}

func Xxcb_query_colors_sizeof(tls *libc.TLS, _buffer uintptr, pixels_len Tuint32_t) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	var xcb_tmp uintptr
	_, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* pixels */
	xcb_block_len += pixels_len * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_query_colors(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, pixels_len Tuint32_t, pixels uintptr) (r Txcb_query_colors_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_query_colors_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_query_colors_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_colors_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_query_colors_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint32_t pixels */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pixels
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = pixels_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req190)))
	return xcb_ret
}

var _xcb_req190 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_QUERY_COLORS),
}

func Xxcb_query_colors_unchecked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, pixels_len Tuint32_t, pixels uintptr) (r Txcb_query_colors_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_query_colors_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_query_colors_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_colors_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_query_colors_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint32_t pixels */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = pixels
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = pixels_len * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req191)))
	return xcb_ret
}

var _xcb_req191 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_QUERY_COLORS),
}

func Xxcb_query_colors_colors(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_query_colors_colors_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_query_colors_reply_t)(unsafe.Pointer(R)).Fcolors_len)
}

func Xxcb_query_colors_colors_iterator(tls *libc.TLS, R uintptr) (r Txcb_rgb_iterator_t) {
	var i Txcb_rgb_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32
	i.Frem = libc.Int32FromUint16((*Txcb_query_colors_reply_t)(unsafe.Pointer(R)).Fcolors_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_query_colors_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_colors_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_lookup_color_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_lookup_color_request_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_lookup_color(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, name_len Tuint16_t, name uintptr) (r Txcb_lookup_color_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_lookup_color_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_lookup_color_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_lookup_color_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_lookup_color_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_lookup_color_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+10, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req192)))
	return xcb_ret
}

var _xcb_req192 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_LOOKUP_COLOR),
}

func Xxcb_lookup_color_unchecked(tls *libc.TLS, c uintptr, cmap Txcb_colormap_t, name_len Tuint16_t, name uintptr) (r Txcb_lookup_color_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_lookup_color_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_lookup_color_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_lookup_color_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_lookup_color_request_t)(unsafe.Pointer(bp + 48))).Fcmap = cmap
	(*(*Txcb_lookup_color_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+10, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req193)))
	return xcb_ret
}

var _xcb_req193 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_LOOKUP_COLOR),
}

func Xxcb_lookup_color_reply(tls *libc.TLS, c uintptr, cookie Txcb_lookup_color_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_create_cursor_checked(tls *libc.TLS, c uintptr, cid Txcb_cursor_t, source Txcb_pixmap_t, mask Txcb_pixmap_t, fore_red Tuint16_t, fore_green Tuint16_t, fore_blue Tuint16_t, back_red Tuint16_t, back_green Tuint16_t, back_blue Tuint16_t, x Tuint16_t, y Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcid = cid
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fsource = source
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fmask = mask
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_red = fore_red
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_green = fore_green
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_blue = fore_blue
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_red = back_red
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_green = back_green
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_blue = back_blue
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req194)))
	return xcb_ret
}

var _xcb_req194 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_create_cursor(tls *libc.TLS, c uintptr, cid Txcb_cursor_t, source Txcb_pixmap_t, mask Txcb_pixmap_t, fore_red Tuint16_t, fore_green Tuint16_t, fore_blue Tuint16_t, back_red Tuint16_t, back_green Tuint16_t, back_blue Tuint16_t, x Tuint16_t, y Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcid = cid
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fsource = source
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fmask = mask
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_red = fore_red
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_green = fore_green
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_blue = fore_blue
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_red = back_red
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_green = back_green
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_blue = back_blue
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fx = x
	(*(*Txcb_create_cursor_request_t)(unsafe.Pointer(bp + 32))).Fy = y
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req195)))
	return xcb_ret
}

var _xcb_req195 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_create_glyph_cursor_checked(tls *libc.TLS, c uintptr, cid Txcb_cursor_t, source_font Txcb_font_t, mask_font Txcb_font_t, source_char Tuint16_t, mask_char Tuint16_t, fore_red Tuint16_t, fore_green Tuint16_t, fore_blue Tuint16_t, back_red Tuint16_t, back_green Tuint16_t, back_blue Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_glyph_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcid = cid
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fsource_font = source_font
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fmask_font = mask_font
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fsource_char = source_char
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fmask_char = mask_char
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_red = fore_red
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_green = fore_green
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_blue = fore_blue
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_red = back_red
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_green = back_green
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_blue = back_blue
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req196)))
	return xcb_ret
}

var _xcb_req196 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_GLYPH_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_create_glyph_cursor(tls *libc.TLS, c uintptr, cid Txcb_cursor_t, source_font Txcb_font_t, mask_font Txcb_font_t, source_char Tuint16_t, mask_char Tuint16_t, fore_red Tuint16_t, fore_green Tuint16_t, fore_blue Tuint16_t, back_red Tuint16_t, back_green Tuint16_t, back_blue Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_create_glyph_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcid = cid
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fsource_font = source_font
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fmask_font = mask_font
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fsource_char = source_char
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fmask_char = mask_char
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_red = fore_red
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_green = fore_green
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_blue = fore_blue
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_red = back_red
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_green = back_green
	(*(*Txcb_create_glyph_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_blue = back_blue
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(32)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req197)))
	return xcb_ret
}

var _xcb_req197 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CREATE_GLYPH_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_free_cursor_checked(tls *libc.TLS, c uintptr, cursor Txcb_cursor_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req198)))
	return xcb_ret
}

var _xcb_req198 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_free_cursor(tls *libc.TLS, c uintptr, cursor Txcb_cursor_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_free_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_free_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_free_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req199)))
	return xcb_ret
}

var _xcb_req199 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FREE_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_recolor_cursor_checked(tls *libc.TLS, c uintptr, cursor Txcb_cursor_t, fore_red Tuint16_t, fore_green Tuint16_t, fore_blue Tuint16_t, back_red Tuint16_t, back_green Tuint16_t, back_blue Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_recolor_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_red = fore_red
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_green = fore_green
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_blue = fore_blue
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_red = back_red
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_green = back_green
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_blue = back_blue
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(20)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req200)))
	return xcb_ret
}

var _xcb_req200 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_RECOLOR_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_recolor_cursor(tls *libc.TLS, c uintptr, cursor Txcb_cursor_t, fore_red Tuint16_t, fore_green Tuint16_t, fore_blue Tuint16_t, back_red Tuint16_t, back_green Tuint16_t, back_blue Tuint16_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_recolor_cursor_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fcursor = cursor
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_red = fore_red
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_green = fore_green
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Ffore_blue = fore_blue
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_red = back_red
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_green = back_green
	(*(*Txcb_recolor_cursor_request_t)(unsafe.Pointer(bp + 32))).Fback_blue = back_blue
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(20)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req201)))
	return xcb_ret
}

var _xcb_req201 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_RECOLOR_CURSOR),
	Fisvoid: uint8(1),
}

func Xxcb_query_best_size(tls *libc.TLS, c uintptr, _class Tuint8_t, drawable Txcb_drawable_t, width Tuint16_t, height Tuint16_t) (r Txcb_query_best_size_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_best_size_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_best_size_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).F_class = _class
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req202)))
	return xcb_ret
}

var _xcb_req202 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_BEST_SIZE),
}

func Xxcb_query_best_size_unchecked(tls *libc.TLS, c uintptr, _class Tuint8_t, drawable Txcb_drawable_t, width Tuint16_t, height Tuint16_t) (r Txcb_query_best_size_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_query_best_size_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_query_best_size_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).F_class = _class
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).Fdrawable = drawable
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).Fwidth = width
	(*(*Txcb_query_best_size_request_t)(unsafe.Pointer(bp + 32))).Fheight = height
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req203)))
	return xcb_ret
}

var _xcb_req203 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_QUERY_BEST_SIZE),
}

func Xxcb_query_best_size_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_best_size_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_query_extension_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* name */
	xcb_block_len += uint32((*Txcb_query_extension_request_t)(unsafe.Pointer(_aux)).Fname_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_query_extension(tls *libc.TLS, c uintptr, name_len Tuint16_t, name uintptr) (r Txcb_query_extension_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_query_extension_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_query_extension_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_extension_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_query_extension_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req204)))
	return xcb_ret
}

var _xcb_req204 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_QUERY_EXTENSION),
}

func Xxcb_query_extension_unchecked(tls *libc.TLS, c uintptr, name_len Tuint16_t, name uintptr) (r Txcb_query_extension_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_query_extension_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_query_extension_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_query_extension_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_query_extension_request_t)(unsafe.Pointer(bp + 48))).Fname_len = name_len
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* char name */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = name
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(name_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req205)))
	return xcb_ret
}

var _xcb_req205 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_QUERY_EXTENSION),
}

func Xxcb_query_extension_reply(tls *libc.TLS, c uintptr, cookie Txcb_query_extension_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_list_extensions_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp_len uint32
	_, _, _, _, _, _, _, _ = _aux, i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp, xcb_tmp_len
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* names */
	i = uint32(0)
	for {
		if !(i < uint32((*Txcb_list_extensions_reply_t)(unsafe.Pointer(_aux)).Fnames_len)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_str_sizeof(tls, xcb_tmp))
		xcb_block_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_list_extensions(tls *libc.TLS, c uintptr) (r Txcb_list_extensions_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_extensions_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_extensions_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_extensions_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req206)))
	return xcb_ret
}

var _xcb_req206 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_EXTENSIONS),
}

func Xxcb_list_extensions_unchecked(tls *libc.TLS, c uintptr) (r Txcb_list_extensions_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_extensions_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_extensions_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_extensions_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req207)))
	return xcb_ret
}

var _xcb_req207 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_EXTENSIONS),
}

func Xxcb_list_extensions_names_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_list_extensions_reply_t)(unsafe.Pointer(R)).Fnames_len)
}

func Xxcb_list_extensions_names_iterator(tls *libc.TLS, R uintptr) (r Txcb_str_iterator_t) {
	var i Txcb_str_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32
	i.Frem = libc.Int32FromUint8((*Txcb_list_extensions_reply_t)(unsafe.Pointer(R)).Fnames_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_extensions_reply(tls *libc.TLS, c uintptr, cookie Txcb_list_extensions_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_change_keyboard_mapping_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* keysyms */
	xcb_block_len += libc.Uint32FromInt32(libc.Int32FromUint8((*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(_aux)).Fkeycode_count)*libc.Int32FromUint8((*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(_aux)).Fkeysyms_per_keycode)) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_keyboard_mapping_checked(tls *libc.TLS, c uintptr, keycode_count Tuint8_t, first_keycode Txcb_keycode_t, keysyms_per_keycode Tuint8_t, keysyms uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_change_keyboard_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(bp + 48))).Fkeycode_count = keycode_count
	(*(*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(bp + 48))).Ffirst_keycode = first_keycode
	(*(*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(bp + 48))).Fkeysyms_per_keycode = keysyms_per_keycode
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_keysym_t keysyms */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = keysyms
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(libc.Int32FromUint8(keycode_count)*libc.Int32FromUint8(keysyms_per_keycode)) * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req208)))
	return xcb_ret
}

var _xcb_req208 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_CHANGE_KEYBOARD_MAPPING),
	Fisvoid: uint8(1),
}

func Xxcb_change_keyboard_mapping(tls *libc.TLS, c uintptr, keycode_count Tuint8_t, first_keycode Txcb_keycode_t, keysyms_per_keycode Tuint8_t, keysyms uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_change_keyboard_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(bp + 48))).Fkeycode_count = keycode_count
	(*(*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(bp + 48))).Ffirst_keycode = first_keycode
	(*(*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(bp + 48))).Fkeysyms_per_keycode = keysyms_per_keycode
	libc.Xmemset(tls, bp+48+6, 0, uint32(2))
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_keysym_t keysyms */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = keysyms
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(libc.Int32FromUint8(keycode_count)*libc.Int32FromUint8(keysyms_per_keycode)) * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req209)))
	return xcb_ret
}

var _xcb_req209 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_CHANGE_KEYBOARD_MAPPING),
	Fisvoid: uint8(1),
}

func Xxcb_change_keyboard_mapping_keysyms(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_change_keyboard_mapping_keysyms_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(R)).Fkeycode_count) * libc.Int32FromUint8((*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(R)).Fkeysyms_per_keycode)
}

func Xxcb_change_keyboard_mapping_keysyms_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8 + uintptr(libc.Int32FromUint8((*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(R)).Fkeycode_count)*libc.Int32FromUint8((*Txcb_change_keyboard_mapping_request_t)(unsafe.Pointer(R)).Fkeysyms_per_keycode))*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_keyboard_mapping_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* keysyms */
	xcb_block_len += (*Txcb_get_keyboard_mapping_reply_t)(unsafe.Pointer(_aux)).Flength * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_keyboard_mapping(tls *libc.TLS, c uintptr, first_keycode Txcb_keycode_t, count Tuint8_t) (r Txcb_get_keyboard_mapping_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_keyboard_mapping_cookie_t
	var _ /* xcb_out at bp+36 */ Txcb_get_keyboard_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_keyboard_mapping_request_t)(unsafe.Pointer(bp + 36))).Fpad0 = uint8(0)
	(*(*Txcb_get_keyboard_mapping_request_t)(unsafe.Pointer(bp + 36))).Ffirst_keycode = first_keycode
	(*(*Txcb_get_keyboard_mapping_request_t)(unsafe.Pointer(bp + 36))).Fcount = count
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 36
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(6)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req210)))
	return xcb_ret
}

var _xcb_req210 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_KEYBOARD_MAPPING),
}

func Xxcb_get_keyboard_mapping_unchecked(tls *libc.TLS, c uintptr, first_keycode Txcb_keycode_t, count Tuint8_t) (r Txcb_get_keyboard_mapping_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_keyboard_mapping_cookie_t
	var _ /* xcb_out at bp+36 */ Txcb_get_keyboard_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_keyboard_mapping_request_t)(unsafe.Pointer(bp + 36))).Fpad0 = uint8(0)
	(*(*Txcb_get_keyboard_mapping_request_t)(unsafe.Pointer(bp + 36))).Ffirst_keycode = first_keycode
	(*(*Txcb_get_keyboard_mapping_request_t)(unsafe.Pointer(bp + 36))).Fcount = count
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 36
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(6)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req211)))
	return xcb_ret
}

var _xcb_req211 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_KEYBOARD_MAPPING),
}

func Xxcb_get_keyboard_mapping_keysyms(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_keyboard_mapping_keysyms_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_get_keyboard_mapping_reply_t)(unsafe.Pointer(R)).Flength)
}

func Xxcb_get_keyboard_mapping_keysyms_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_get_keyboard_mapping_reply_t)(unsafe.Pointer(R)).Flength)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_keyboard_mapping_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_keyboard_mapping_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_change_keyboard_control_value_list_serialize(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_parts_idx uint32
	var xcb_out, xcb_tmp uintptr
	var xcb_parts [9]Tiovec
	var _ /* xcb_pad0 at bp+0 */ [3]int8
	_, _, _, _, _, _, _, _, _, _ = i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_out, xcb_pad, xcb_padding_offset, xcb_parts, xcb_parts_idx, xcb_tmp
	xcb_out = *(*uintptr)(unsafe.Pointer(_buffer))
	xcb_buffer_len = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	xcb_pad = uint32(0)
	*(*[3]int8)(unsafe.Pointer(bp)) = [3]int8{}
	xcb_parts_idx = uint32(0)
	xcb_block_len = uint32(0)
	if value_mask&uint32(_XCB_KB_KEY_CLICK_PERCENT) != 0 {
		/* xcb_change_keyboard_control_value_list_t.key_click_percent */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_BELL_PERCENT) != 0 {
		/* xcb_change_keyboard_control_value_list_t.bell_percent */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 4
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_BELL_PITCH) != 0 {
		/* xcb_change_keyboard_control_value_list_t.bell_pitch */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 8
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_BELL_DURATION) != 0 {
		/* xcb_change_keyboard_control_value_list_t.bell_duration */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 12
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_LED) != 0 {
		/* xcb_change_keyboard_control_value_list_t.led */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 16
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_LED_MODE) != 0 {
		/* xcb_change_keyboard_control_value_list_t.led_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 20
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_KEY) != 0 {
		/* xcb_change_keyboard_control_value_list_t.key */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 24
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_AUTO_REPEAT_MODE) != 0 {
		/* xcb_change_keyboard_control_value_list_t.auto_repeat_mode */
		xcb_parts[xcb_parts_idx].Fiov_base = _aux + 28
		xcb_block_len += uint32(4)
		xcb_parts[xcb_parts_idx].Fiov_len = uint32(4)
		xcb_parts_idx++
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_parts[xcb_parts_idx].Fiov_base = bp
		xcb_parts[xcb_parts_idx].Fiov_len = xcb_pad
		xcb_parts_idx++
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	if libc.UintptrFromInt32(0) == xcb_out {
		/* allocate memory */
		xcb_out = libc.Xmalloc(tls, xcb_buffer_len)
		*(*uintptr)(unsafe.Pointer(_buffer)) = xcb_out
	}
	xcb_tmp = xcb_out
	i = uint32(0)
	for {
		if !(i < xcb_parts_idx) {
			break
		}
		if uintptr(0) != xcb_parts[i].Fiov_base && uint32(0) != xcb_parts[i].Fiov_len {
			libc.Xmemcpy(tls, xcb_tmp, xcb_parts[i].Fiov_base, xcb_parts[i].Fiov_len)
		}
		if uint32(0) != xcb_parts[i].Fiov_len {
			xcb_tmp += uintptr(xcb_parts[i].Fiov_len)
		}
		goto _1
	_1:
		;
		i++
	}
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_keyboard_control_value_list_unpack(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t, _aux uintptr) (r int32) {
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset uint32
	var xcb_tmp uintptr
	_, _, _, _, _, _ = xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_padding_offset, xcb_tmp
	xcb_tmp = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_padding_offset = uint32(0)
	if value_mask&uint32(_XCB_KB_KEY_CLICK_PERCENT) != 0 {
		/* xcb_change_keyboard_control_value_list_t.key_click_percent */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fkey_click_percent = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_BELL_PERCENT) != 0 {
		/* xcb_change_keyboard_control_value_list_t.bell_percent */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fbell_percent = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_BELL_PITCH) != 0 {
		/* xcb_change_keyboard_control_value_list_t.bell_pitch */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fbell_pitch = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_BELL_DURATION) != 0 {
		/* xcb_change_keyboard_control_value_list_t.bell_duration */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fbell_duration = *(*Tint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_LED) != 0 {
		/* xcb_change_keyboard_control_value_list_t.led */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fled = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_LED_MODE) != 0 {
		/* xcb_change_keyboard_control_value_list_t.led_mode */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fled_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_KEY) != 0 {
		/* xcb_change_keyboard_control_value_list_t.key */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fkey = *(*Txcb_keycode32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	if value_mask&uint32(_XCB_KB_AUTO_REPEAT_MODE) != 0 {
		/* xcb_change_keyboard_control_value_list_t.auto_repeat_mode */
		(*Txcb_change_keyboard_control_value_list_t)(unsafe.Pointer(_aux)).Fauto_repeat_mode = *(*Tuint32_t)(unsafe.Pointer(xcb_tmp))
		xcb_block_len += uint32(4)
		xcb_tmp += uintptr(4)
		xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	}
	/* insert padding */
	xcb_pad = -(xcb_block_len + xcb_padding_offset) & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	xcb_padding_offset = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_keyboard_control_value_list_sizeof(tls *libc.TLS, _buffer uintptr, value_mask Tuint32_t) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* _aux at bp+0 */ Txcb_change_keyboard_control_value_list_t
	return Xxcb_change_keyboard_control_value_list_unpack(tls, _buffer, value_mask, bp)
}

func Xxcb_change_keyboard_control_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* value_list */
	xcb_block_len += libc.Uint32FromInt32(Xxcb_change_keyboard_control_value_list_sizeof(tls, xcb_tmp, (*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(_aux)).Fvalue_mask))
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_keyboard_control_checked(tls *libc.TLS, c uintptr, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_change_keyboard_control_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_keyboard_control_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_keyboard_control_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req212)))
	return xcb_ret
}

var _xcb_req212 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_KEYBOARD_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_change_keyboard_control(tls *libc.TLS, c uintptr, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+40 */ Txcb_change_keyboard_control_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_keyboard_control_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = value_list
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_keyboard_control_value_list_sizeof(tls, value_list, value_mask))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req213)))
	return xcb_ret
}

var _xcb_req213 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_KEYBOARD_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_change_keyboard_control_aux_checked(tls *libc.TLS, c uintptr, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+48 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_change_keyboard_control_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 48)) = uintptr(0)
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_keyboard_control_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_keyboard_control_value_list_serialize(tls, bp+48, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 48))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req214)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 48)))
	return xcb_ret
}

var _xcb_req214 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_KEYBOARD_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_change_keyboard_control_aux(tls *libc.TLS, c uintptr, value_mask Tuint32_t, value_list uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_aux0 at bp+48 */ uintptr
	var _ /* xcb_out at bp+40 */ Txcb_change_keyboard_control_request_t
	var _ /* xcb_parts at bp+0 */ [5]Tiovec
	_ = xcb_ret
	*(*uintptr)(unsafe.Pointer(bp + 48)) = uintptr(0)
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fpad0 = uint8(0)
	(*(*Txcb_change_keyboard_control_request_t)(unsafe.Pointer(bp + 40))).Fvalue_mask = value_mask
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 40
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_change_keyboard_control_value_list_t value_list */
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(Xxcb_change_keyboard_control_value_list_serialize(tls, bp+48, value_mask, value_list))
	(*(*[5]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = *(*uintptr)(unsafe.Pointer(bp + 48))
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req215)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp + 48)))
	return xcb_ret
}

var _xcb_req215 = Txcb_protocol_request_t{
	Fcount:  uint32(3),
	Fopcode: uint8(m_XCB_CHANGE_KEYBOARD_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_change_keyboard_control_value_list(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_get_keyboard_control(tls *libc.TLS, c uintptr) (r Txcb_get_keyboard_control_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_keyboard_control_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_keyboard_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_keyboard_control_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req216)))
	return xcb_ret
}

var _xcb_req216 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_KEYBOARD_CONTROL),
}

func Xxcb_get_keyboard_control_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_keyboard_control_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_keyboard_control_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_keyboard_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_keyboard_control_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req217)))
	return xcb_ret
}

var _xcb_req217 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_KEYBOARD_CONTROL),
}

func Xxcb_get_keyboard_control_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_keyboard_control_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_bell_checked(tls *libc.TLS, c uintptr, percent Tint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_bell_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_bell_request_t)(unsafe.Pointer(bp + 32))).Fpercent = percent
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req218)))
	return xcb_ret
}

var _xcb_req218 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_BELL),
	Fisvoid: uint8(1),
}

func Xxcb_bell(tls *libc.TLS, c uintptr, percent Tint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_bell_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_bell_request_t)(unsafe.Pointer(bp + 32))).Fpercent = percent
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req219)))
	return xcb_ret
}

var _xcb_req219 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_BELL),
	Fisvoid: uint8(1),
}

func Xxcb_change_pointer_control_checked(tls *libc.TLS, c uintptr, acceleration_numerator Tint16_t, acceleration_denominator Tint16_t, threshold Tint16_t, do_acceleration Tuint8_t, do_threshold Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_change_pointer_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Facceleration_numerator = acceleration_numerator
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Facceleration_denominator = acceleration_denominator
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fthreshold = threshold
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fdo_acceleration = do_acceleration
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fdo_threshold = do_threshold
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req220)))
	return xcb_ret
}

var _xcb_req220 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CHANGE_POINTER_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_change_pointer_control(tls *libc.TLS, c uintptr, acceleration_numerator Tint16_t, acceleration_denominator Tint16_t, threshold Tint16_t, do_acceleration Tuint8_t, do_threshold Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_change_pointer_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Facceleration_numerator = acceleration_numerator
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Facceleration_denominator = acceleration_denominator
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fthreshold = threshold
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fdo_acceleration = do_acceleration
	(*(*Txcb_change_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fdo_threshold = do_threshold
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req221)))
	return xcb_ret
}

var _xcb_req221 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_CHANGE_POINTER_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_get_pointer_control(tls *libc.TLS, c uintptr) (r Txcb_get_pointer_control_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_pointer_control_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_pointer_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req222)))
	return xcb_ret
}

var _xcb_req222 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_POINTER_CONTROL),
}

func Xxcb_get_pointer_control_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_pointer_control_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_pointer_control_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_pointer_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_pointer_control_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req223)))
	return xcb_ret
}

var _xcb_req223 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_POINTER_CONTROL),
}

func Xxcb_get_pointer_control_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_pointer_control_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_set_screen_saver_checked(tls *libc.TLS, c uintptr, timeout Tint16_t, interval Tint16_t, prefer_blanking Tuint8_t, allow_exposures Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_screen_saver_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Ftimeout = timeout
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Finterval = interval
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fprefer_blanking = prefer_blanking
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fallow_exposures = allow_exposures
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(10)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req224)))
	return xcb_ret
}

var _xcb_req224 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_SCREEN_SAVER),
	Fisvoid: uint8(1),
}

func Xxcb_set_screen_saver(tls *libc.TLS, c uintptr, timeout Tint16_t, interval Tint16_t, prefer_blanking Tuint8_t, allow_exposures Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_screen_saver_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Ftimeout = timeout
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Finterval = interval
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fprefer_blanking = prefer_blanking
	(*(*Txcb_set_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fallow_exposures = allow_exposures
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(10)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req225)))
	return xcb_ret
}

var _xcb_req225 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_SCREEN_SAVER),
	Fisvoid: uint8(1),
}

func Xxcb_get_screen_saver(tls *libc.TLS, c uintptr) (r Txcb_get_screen_saver_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_screen_saver_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_screen_saver_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req226)))
	return xcb_ret
}

var _xcb_req226 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_SCREEN_SAVER),
}

func Xxcb_get_screen_saver_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_screen_saver_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_screen_saver_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_screen_saver_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req227)))
	return xcb_ret
}

var _xcb_req227 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_SCREEN_SAVER),
}

func Xxcb_get_screen_saver_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_screen_saver_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_change_hosts_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(8)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* address */
	xcb_block_len += uint32((*Txcb_change_hosts_request_t)(unsafe.Pointer(_aux)).Faddress_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_change_hosts_checked(tls *libc.TLS, c uintptr, mode Tuint8_t, family Tuint8_t, address_len Tuint16_t, address uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_change_hosts_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Fmode = mode
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Ffamily = family
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Faddress_len = address_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t address */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = address
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(address_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req228)))
	return xcb_ret
}

var _xcb_req228 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_CHANGE_HOSTS),
	Fisvoid: uint8(1),
}

func Xxcb_change_hosts(tls *libc.TLS, c uintptr, mode Tuint8_t, family Tuint8_t, address_len Tuint16_t, address uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_change_hosts_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Fmode = mode
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Ffamily = family
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_change_hosts_request_t)(unsafe.Pointer(bp + 48))).Faddress_len = address_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t address */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = address
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(address_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req229)))
	return xcb_ret
}

var _xcb_req229 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_CHANGE_HOSTS),
	Fisvoid: uint8(1),
}

func Xxcb_change_hosts_address(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*8
}

func Xxcb_change_hosts_address_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_change_hosts_request_t)(unsafe.Pointer(R)).Faddress_len)
}

func Xxcb_change_hosts_address_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*8 + uintptr((*Txcb_change_hosts_request_t)(unsafe.Pointer(R)).Faddress_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_host_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* address */
	xcb_block_len += uint32((*Txcb_host_t)(unsafe.Pointer(_aux)).Faddress_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	xcb_align_to = uint32(4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_host_address(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*4
}

func Xxcb_host_address_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_host_t)(unsafe.Pointer(R)).Faddress_len)
}

func Xxcb_host_address_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*4 + uintptr((*Txcb_host_t)(unsafe.Pointer(R)).Faddress_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_host_next(tls *libc.TLS, i uintptr) {
	var R uintptr
	var child Txcb_generic_iterator_t
	_, _ = R, child
	R = (*Txcb_host_iterator_t)(unsafe.Pointer(i)).Fdata
	child.Fdata = R + uintptr(Xxcb_host_sizeof(tls, R))
	(*Txcb_host_iterator_t)(unsafe.Pointer(i)).Findex = int32(child.Fdata) - int32((*Txcb_host_iterator_t)(unsafe.Pointer(i)).Fdata)
	(*Txcb_host_iterator_t)(unsafe.Pointer(i)).Frem--
	(*Txcb_host_iterator_t)(unsafe.Pointer(i)).Fdata = child.Fdata
}

func Xxcb_host_end(tls *libc.TLS, _i Txcb_host_iterator_t) (r Txcb_generic_iterator_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Txcb_host_iterator_t)(unsafe.Pointer(bp)) = _i
	var ret Txcb_generic_iterator_t
	_ = ret
	for (*(*Txcb_host_iterator_t)(unsafe.Pointer(bp))).Frem > 0 {
		Xxcb_host_next(tls, bp)
	}
	ret.Fdata = (*(*Txcb_host_iterator_t)(unsafe.Pointer(bp))).Fdata
	ret.Frem = (*(*Txcb_host_iterator_t)(unsafe.Pointer(bp))).Frem
	ret.Findex = (*(*Txcb_host_iterator_t)(unsafe.Pointer(bp))).Findex
	return ret
}

func Xxcb_list_hosts_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp_len uint32
	_, _, _, _, _, _, _, _ = _aux, i, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp, xcb_tmp_len
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* hosts */
	i = uint32(0)
	for {
		if !(i < uint32((*Txcb_list_hosts_reply_t)(unsafe.Pointer(_aux)).Fhosts_len)) {
			break
		}
		xcb_tmp_len = libc.Uint32FromInt32(Xxcb_host_sizeof(tls, xcb_tmp))
		xcb_block_len += xcb_tmp_len
		xcb_tmp += uintptr(xcb_tmp_len)
		goto _1
	_1:
		;
		i++
	}
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 2)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_list_hosts(tls *libc.TLS, c uintptr) (r Txcb_list_hosts_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_hosts_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_hosts_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_hosts_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req230)))
	return xcb_ret
}

var _xcb_req230 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_HOSTS),
}

func Xxcb_list_hosts_unchecked(tls *libc.TLS, c uintptr) (r Txcb_list_hosts_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_list_hosts_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_list_hosts_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_list_hosts_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req231)))
	return xcb_ret
}

var _xcb_req231 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_LIST_HOSTS),
}

func Xxcb_list_hosts_hosts_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_list_hosts_reply_t)(unsafe.Pointer(R)).Fhosts_len)
}

func Xxcb_list_hosts_hosts_iterator(tls *libc.TLS, R uintptr) (r Txcb_host_iterator_t) {
	var i Txcb_host_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32
	i.Frem = libc.Int32FromUint16((*Txcb_list_hosts_reply_t)(unsafe.Pointer(R)).Fhosts_len)
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_list_hosts_reply(tls *libc.TLS, c uintptr, cookie Txcb_list_hosts_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_set_access_control_checked(tls *libc.TLS, c uintptr, mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_access_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_access_control_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req232)))
	return xcb_ret
}

var _xcb_req232 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_ACCESS_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_set_access_control(tls *libc.TLS, c uintptr, mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_access_control_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_access_control_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req233)))
	return xcb_ret
}

var _xcb_req233 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_ACCESS_CONTROL),
	Fisvoid: uint8(1),
}

func Xxcb_set_close_down_mode_checked(tls *libc.TLS, c uintptr, mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_close_down_mode_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_close_down_mode_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req234)))
	return xcb_ret
}

var _xcb_req234 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_CLOSE_DOWN_MODE),
	Fisvoid: uint8(1),
}

func Xxcb_set_close_down_mode(tls *libc.TLS, c uintptr, mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_set_close_down_mode_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_close_down_mode_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req235)))
	return xcb_ret
}

var _xcb_req235 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_SET_CLOSE_DOWN_MODE),
	Fisvoid: uint8(1),
}

func Xxcb_kill_client_checked(tls *libc.TLS, c uintptr, resource Tuint32_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_kill_client_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_kill_client_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_kill_client_request_t)(unsafe.Pointer(bp + 32))).Fresource = resource
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req236)))
	return xcb_ret
}

var _xcb_req236 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_KILL_CLIENT),
	Fisvoid: uint8(1),
}

func Xxcb_kill_client(tls *libc.TLS, c uintptr, resource Tuint32_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_kill_client_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_kill_client_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*Txcb_kill_client_request_t)(unsafe.Pointer(bp + 32))).Fresource = resource
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req237)))
	return xcb_ret
}

var _xcb_req237 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_KILL_CLIENT),
	Fisvoid: uint8(1),
}

func Xxcb_rotate_properties_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(12)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* atoms */
	xcb_block_len += uint32((*Txcb_rotate_properties_request_t)(unsafe.Pointer(_aux)).Fatoms_len) * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_rotate_properties_checked(tls *libc.TLS, c uintptr, window Txcb_window_t, atoms_len Tuint16_t, delta Tint16_t, atoms uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_rotate_properties_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fwindow = window
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fatoms_len = atoms_len
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fdelta = delta
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_atom_t atoms */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = atoms
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(atoms_len) * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req238)))
	return xcb_ret
}

var _xcb_req238 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_ROTATE_PROPERTIES),
	Fisvoid: uint8(1),
}

func Xxcb_rotate_properties(tls *libc.TLS, c uintptr, window Txcb_window_t, atoms_len Tuint16_t, delta Tint16_t, atoms uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_rotate_properties_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fpad0 = uint8(0)
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fwindow = window
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fatoms_len = atoms_len
	(*(*Txcb_rotate_properties_request_t)(unsafe.Pointer(bp + 48))).Fdelta = delta
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(12)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_atom_t atoms */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = atoms
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(atoms_len) * uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req239)))
	return xcb_ret
}

var _xcb_req239 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_ROTATE_PROPERTIES),
	Fisvoid: uint8(1),
}

func Xxcb_rotate_properties_atoms(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*12
}

func Xxcb_rotate_properties_atoms_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint16((*Txcb_rotate_properties_request_t)(unsafe.Pointer(R)).Fatoms_len)
}

func Xxcb_rotate_properties_atoms_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*12 + uintptr((*Txcb_rotate_properties_request_t)(unsafe.Pointer(R)).Fatoms_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_force_screen_saver_checked(tls *libc.TLS, c uintptr, mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_force_screen_saver_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_force_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req240)))
	return xcb_ret
}

var _xcb_req240 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FORCE_SCREEN_SAVER),
	Fisvoid: uint8(1),
}

func Xxcb_force_screen_saver(tls *libc.TLS, c uintptr, mode Tuint8_t) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_force_screen_saver_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_force_screen_saver_request_t)(unsafe.Pointer(bp + 32))).Fmode = mode
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req241)))
	return xcb_ret
}

var _xcb_req241 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_FORCE_SCREEN_SAVER),
	Fisvoid: uint8(1),
}

func Xxcb_set_pointer_mapping_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* map */
	xcb_block_len += uint32((*Txcb_set_pointer_mapping_request_t)(unsafe.Pointer(_aux)).Fmap_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_set_pointer_mapping(tls *libc.TLS, c uintptr, map_len Tuint8_t, map1 uintptr) (r Txcb_set_pointer_mapping_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_set_pointer_mapping_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_pointer_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_pointer_mapping_request_t)(unsafe.Pointer(bp + 48))).Fmap_len = map_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t map */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = map1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(map_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req242)))
	return xcb_ret
}

var _xcb_req242 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_POINTER_MAPPING),
}

func Xxcb_set_pointer_mapping_unchecked(tls *libc.TLS, c uintptr, map_len Tuint8_t, map1 uintptr) (r Txcb_set_pointer_mapping_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_set_pointer_mapping_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_pointer_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_pointer_mapping_request_t)(unsafe.Pointer(bp + 48))).Fmap_len = map_len
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* uint8_t map */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = map1
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = uint32(map_len) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req243)))
	return xcb_ret
}

var _xcb_req243 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_POINTER_MAPPING),
}

func Xxcb_set_pointer_mapping_reply(tls *libc.TLS, c uintptr, cookie Txcb_set_pointer_mapping_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_get_pointer_mapping_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* map */
	xcb_block_len += uint32((*Txcb_get_pointer_mapping_reply_t)(unsafe.Pointer(_aux)).Fmap_len) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_pointer_mapping(tls *libc.TLS, c uintptr) (r Txcb_get_pointer_mapping_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_pointer_mapping_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_pointer_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_pointer_mapping_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req244)))
	return xcb_ret
}

var _xcb_req244 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_POINTER_MAPPING),
}

func Xxcb_get_pointer_mapping_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_pointer_mapping_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_pointer_mapping_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_pointer_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_pointer_mapping_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req245)))
	return xcb_ret
}

var _xcb_req245 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_POINTER_MAPPING),
}

func Xxcb_get_pointer_mapping_map(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_pointer_mapping_map_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_get_pointer_mapping_reply_t)(unsafe.Pointer(R)).Fmap_len)
}

func Xxcb_get_pointer_mapping_map_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_get_pointer_mapping_reply_t)(unsafe.Pointer(R)).Fmap_len)
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_pointer_mapping_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_pointer_mapping_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_set_modifier_mapping_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* keycodes */
	xcb_block_len += libc.Uint32FromInt32(libc.Int32FromUint8((*Txcb_set_modifier_mapping_request_t)(unsafe.Pointer(_aux)).Fkeycodes_per_modifier)*libc.Int32FromInt32(8)) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_set_modifier_mapping(tls *libc.TLS, c uintptr, keycodes_per_modifier Tuint8_t, keycodes uintptr) (r Txcb_set_modifier_mapping_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_set_modifier_mapping_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_modifier_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_modifier_mapping_request_t)(unsafe.Pointer(bp + 48))).Fkeycodes_per_modifier = keycodes_per_modifier
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_keycode_t keycodes */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = keycodes
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(libc.Int32FromUint8(keycodes_per_modifier)*libc.Int32FromInt32(8)) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req246)))
	return xcb_ret
}

var _xcb_req246 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_MODIFIER_MAPPING),
}

func Xxcb_set_modifier_mapping_unchecked(tls *libc.TLS, c uintptr, keycodes_per_modifier Tuint8_t, keycodes uintptr) (r Txcb_set_modifier_mapping_cookie_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var xcb_ret Txcb_set_modifier_mapping_cookie_t
	var _ /* xcb_out at bp+48 */ Txcb_set_modifier_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [6]Tiovec
	_ = xcb_ret
	(*(*Txcb_set_modifier_mapping_request_t)(unsafe.Pointer(bp + 48))).Fkeycodes_per_modifier = keycodes_per_modifier
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 48
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	/* xcb_keycode_t keycodes */
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_base = keycodes
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len = libc.Uint32FromInt32(libc.Int32FromUint8(keycodes_per_modifier)*libc.Int32FromInt32(8)) * uint32(1)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_base = uintptr(0)
	(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(5)].Fiov_len = -(*(*[6]Tiovec)(unsafe.Pointer(bp)))[int32(4)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req247)))
	return xcb_ret
}

var _xcb_req247 = Txcb_protocol_request_t{
	Fcount:  uint32(4),
	Fopcode: uint8(m_XCB_SET_MODIFIER_MAPPING),
}

func Xxcb_set_modifier_mapping_reply(tls *libc.TLS, c uintptr, cookie Txcb_set_modifier_mapping_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_get_modifier_mapping_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* keycodes */
	xcb_block_len += libc.Uint32FromInt32(libc.Int32FromUint8((*Txcb_get_modifier_mapping_reply_t)(unsafe.Pointer(_aux)).Fkeycodes_per_modifier)*libc.Int32FromInt32(8)) * uint32(1)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 1)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_get_modifier_mapping(tls *libc.TLS, c uintptr) (r Txcb_get_modifier_mapping_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_modifier_mapping_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_modifier_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_modifier_mapping_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req248)))
	return xcb_ret
}

var _xcb_req248 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_MODIFIER_MAPPING),
}

func Xxcb_get_modifier_mapping_unchecked(tls *libc.TLS, c uintptr) (r Txcb_get_modifier_mapping_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_get_modifier_mapping_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_get_modifier_mapping_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_get_modifier_mapping_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req249)))
	return xcb_ret
}

var _xcb_req249 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_GET_MODIFIER_MAPPING),
}

func Xxcb_get_modifier_mapping_keycodes(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_get_modifier_mapping_keycodes_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint8((*Txcb_get_modifier_mapping_reply_t)(unsafe.Pointer(R)).Fkeycodes_per_modifier) * int32(8)
}

func Xxcb_get_modifier_mapping_keycodes_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr(libc.Int32FromUint8((*Txcb_get_modifier_mapping_reply_t)(unsafe.Pointer(R)).Fkeycodes_per_modifier)*libc.Int32FromInt32(8))
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_get_modifier_mapping_reply(tls *libc.TLS, c uintptr, cookie Txcb_get_modifier_mapping_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_no_operation_checked(tls *libc.TLS, c uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_no_operation_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_no_operation_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req250)))
	return xcb_ret
}

var _xcb_req250 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_NO_OPERATION),
	Fisvoid: uint8(1),
}

func Xxcb_no_operation(tls *libc.TLS, c uintptr) (r Txcb_void_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_void_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_no_operation_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_no_operation_request_t)(unsafe.Pointer(bp + 32))).Fpad0 = uint8(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req251)))
	return xcb_ret
}

var _xcb_req251 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fopcode: uint8(m_XCB_NO_OPERATION),
	Fisvoid: uint8(1),
}

func Xxcb_big_requests_enable(tls *libc.TLS, c uintptr) (r Txcb_big_requests_enable_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_big_requests_enable_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_big_requests_enable_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req252)))
	return xcb_ret
}

var _xcb_req252 = Txcb_protocol_request_t{
	Fcount: uint32(2),
	Fext:   uintptr(unsafe.Pointer(&Xxcb_big_requests_id)),
}

func Xxcb_big_requests_enable_unchecked(tls *libc.TLS, c uintptr) (r Txcb_big_requests_enable_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_big_requests_enable_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_big_requests_enable_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req253)))
	return xcb_ret
}

var _xcb_req253 = Txcb_protocol_request_t{
	Fcount: uint32(2),
	Fext:   uintptr(unsafe.Pointer(&Xxcb_big_requests_id)),
}

func Xxcb_big_requests_enable_reply(tls *libc.TLS, c uintptr, cookie Txcb_big_requests_enable_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_xc_misc_get_version(tls *libc.TLS, c uintptr, client_major_version Tuint16_t, client_minor_version Tuint16_t) (r Txcb_xc_misc_get_version_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_xc_misc_get_version_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_xc_misc_get_version_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_xc_misc_get_version_request_t)(unsafe.Pointer(bp + 32))).Fclient_major_version = client_major_version
	(*(*Txcb_xc_misc_get_version_request_t)(unsafe.Pointer(bp + 32))).Fclient_minor_version = client_minor_version
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req254)))
	return xcb_ret
}

var _xcb_req254 = Txcb_protocol_request_t{
	Fcount: uint32(2),
	Fext:   uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)),
}

func Xxcb_xc_misc_get_version_unchecked(tls *libc.TLS, c uintptr, client_major_version Tuint16_t, client_minor_version Tuint16_t) (r Txcb_xc_misc_get_version_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_xc_misc_get_version_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_xc_misc_get_version_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_xc_misc_get_version_request_t)(unsafe.Pointer(bp + 32))).Fclient_major_version = client_major_version
	(*(*Txcb_xc_misc_get_version_request_t)(unsafe.Pointer(bp + 32))).Fclient_minor_version = client_minor_version
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req255)))
	return xcb_ret
}

var _xcb_req255 = Txcb_protocol_request_t{
	Fcount: uint32(2),
	Fext:   uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)),
}

func Xxcb_xc_misc_get_version_reply(tls *libc.TLS, c uintptr, cookie Txcb_xc_misc_get_version_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_xc_misc_get_xid_range(tls *libc.TLS, c uintptr) (r Txcb_xc_misc_get_xid_range_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_xc_misc_get_xid_range_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_xc_misc_get_xid_range_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req256)))
	return xcb_ret
}

var _xcb_req256 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fext:    uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)),
	Fopcode: uint8(m_XCB_XC_MISC_GET_XID_RANGE),
}

func Xxcb_xc_misc_get_xid_range_unchecked(tls *libc.TLS, c uintptr) (r Txcb_xc_misc_get_xid_range_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_xc_misc_get_xid_range_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_xc_misc_get_xid_range_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(4)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req257)))
	return xcb_ret
}

var _xcb_req257 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fext:    uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)),
	Fopcode: uint8(m_XCB_XC_MISC_GET_XID_RANGE),
}

func Xxcb_xc_misc_get_xid_range_reply(tls *libc.TLS, c uintptr, cookie Txcb_xc_misc_get_xid_range_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func Xxcb_xc_misc_get_xid_list_sizeof(tls *libc.TLS, _buffer uintptr) (r int32) {
	var _aux, xcb_tmp uintptr
	var xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad uint32
	_, _, _, _, _, _ = _aux, xcb_align_to, xcb_block_len, xcb_buffer_len, xcb_pad, xcb_tmp
	xcb_tmp = _buffer
	_aux = _buffer
	xcb_buffer_len = uint32(0)
	xcb_block_len = uint32(0)
	xcb_pad = uint32(0)
	xcb_align_to = uint32(0)
	xcb_block_len += uint32(32)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_buffer_len += xcb_block_len
	xcb_block_len = uint32(0)
	/* ids */
	xcb_block_len += (*Txcb_xc_misc_get_xid_list_reply_t)(unsafe.Pointer(_aux)).Fids_len * uint32(4)
	xcb_tmp += uintptr(xcb_block_len)
	xcb_align_to = uint32(libc.UintptrFromInt32(0) + 4)
	/* insert padding */
	xcb_pad = -xcb_block_len & (xcb_align_to - uint32(1))
	xcb_buffer_len += xcb_block_len + xcb_pad
	if uint32(0) != xcb_pad {
		xcb_tmp += uintptr(xcb_pad)
		xcb_pad = uint32(0)
	}
	xcb_block_len = uint32(0)
	return libc.Int32FromUint32(xcb_buffer_len)
}

func Xxcb_xc_misc_get_xid_list(tls *libc.TLS, c uintptr, count Tuint32_t) (r Txcb_xc_misc_get_xid_list_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_xc_misc_get_xid_list_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_xc_misc_get_xid_list_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_xc_misc_get_xid_list_request_t)(unsafe.Pointer(bp + 32))).Fcount = count
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, int32(_XCB_REQUEST_CHECKED), bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req258)))
	return xcb_ret
}

var _xcb_req258 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fext:    uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)),
	Fopcode: uint8(m_XCB_XC_MISC_GET_XID_LIST),
}

func Xxcb_xc_misc_get_xid_list_unchecked(tls *libc.TLS, c uintptr, count Tuint32_t) (r Txcb_xc_misc_get_xid_list_cookie_t) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var xcb_ret Txcb_xc_misc_get_xid_list_cookie_t
	var _ /* xcb_out at bp+32 */ Txcb_xc_misc_get_xid_list_request_t
	var _ /* xcb_parts at bp+0 */ [4]Tiovec
	_ = xcb_ret
	(*(*Txcb_xc_misc_get_xid_list_request_t)(unsafe.Pointer(bp + 32))).Fcount = count
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_base = bp + 32
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len = uint32(8)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_base = uintptr(0)
	(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(3)].Fiov_len = -(*(*[4]Tiovec)(unsafe.Pointer(bp)))[int32(2)].Fiov_len & uint32(3)
	xcb_ret.Fsequence = Xxcb_send_request(tls, c, 0, bp+uintptr(2)*8, uintptr(unsafe.Pointer(&_xcb_req259)))
	return xcb_ret
}

var _xcb_req259 = Txcb_protocol_request_t{
	Fcount:  uint32(2),
	Fext:    uintptr(unsafe.Pointer(&Xxcb_xc_misc_id)),
	Fopcode: uint8(m_XCB_XC_MISC_GET_XID_LIST),
}

func Xxcb_xc_misc_get_xid_list_ids(tls *libc.TLS, R uintptr) (r uintptr) {
	return R + libc.UintptrFromInt32(1)*32
}

func Xxcb_xc_misc_get_xid_list_ids_length(tls *libc.TLS, R uintptr) (r int32) {
	return libc.Int32FromUint32((*Txcb_xc_misc_get_xid_list_reply_t)(unsafe.Pointer(R)).Fids_len)
}

func Xxcb_xc_misc_get_xid_list_ids_end(tls *libc.TLS, R uintptr) (r Txcb_generic_iterator_t) {
	var i Txcb_generic_iterator_t
	_ = i
	i.Fdata = R + libc.UintptrFromInt32(1)*32 + uintptr((*Txcb_xc_misc_get_xid_list_reply_t)(unsafe.Pointer(R)).Fids_len)*4
	i.Frem = 0
	i.Findex = int32(i.Fdata) - int32(R)
	return i
}

func Xxcb_xc_misc_get_xid_list_reply(tls *libc.TLS, c uintptr, cookie Txcb_xc_misc_get_xid_list_cookie_t, e uintptr) (r uintptr) {
	return Xxcb_wait_for_reply(tls, c, cookie.Fsequence, e)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

/**
 * @}
 */

var Xxcb_big_requests_id = Txcb_extension_t{
	Fname: __ccgo_ts + 93,
}

/**
 * @}
 */

var Xxcb_xc_misc_id = Txcb_extension_t{
	Fname: __ccgo_ts + 106,
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "GLX\x00DISPLAY\x00unix\x00%s%d\x00tcp\x00inet\x00inet6\x00localhost\x00%hu\x00XDM-AUTHORIZATION-1\x00MIT-MAGIC-COOKIE-1\x00%d\x00BIG-REQUESTS\x00XC-MISC\x00"
