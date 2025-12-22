// Code generated for linux/386 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors -ignore-unsupported-alignment -ignore-link-errors -o libmd.go --package-name libmd src/.libs/libmd.a', DO NOT EDIT.

//go:build linux && 386

package libmd

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_BIG_ENDIAN = "__BIG_ENDIAN"
const m_BYTE_ORDER = "__BYTE_ORDER"
const m_FD_SETSIZE = 1024
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDIO_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_HAVE_WCHAR_H = 1
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
const m_LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const m_LT_OBJDIR = ".libs/"
const m_MD2_DIGEST_LENGTH = 16
const m_MD2_DIGEST_STRING_LENGTH = 33
const m_NDEBUG = 1
const m_PACKAGE = "libmd"
const m_PACKAGE_BUGREPORT = "libbsd@lists.freedesktop.org"
const m_PACKAGE_NAME = "libmd"
const m_PACKAGE_STRING = "libmd 1.0.4"
const m_PACKAGE_TARNAME = "libmd"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "1.0.4"
const m_PDP_ENDIAN = "__PDP_ENDIAN"
const m_PTRDIFF_MAX = "INT32_MAX"
const m_PTRDIFF_MIN = "INT32_MIN"
const m_SIG_ATOMIC_MAX = "INT32_MAX"
const m_SIG_ATOMIC_MIN = "INT32_MIN"
const m_SIZE_MAX = "UINT32_MAX"
const m_STDC_HEADERS = 1
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
const m_VERSION = "1.0.4"
const m_WINT_MAX = "UINT32_MAX"
const m_WINT_MIN = 0
const m__ALL_SOURCE = 1
const m__DARWIN_C_SOURCE = 1
const m__FILE_OFFSET_BITS = 64
const m__GNU_SOURCE = 1
const m__HPUX_ALT_XOPEN_SOCKET_API = 1
const m__ILP32 = 1
const m__NETBSD_SOURCE = 1
const m__OPENBSD_SOURCE = 1
const m__POSIX_PTHREAD_SEMANTICS = 1
const m__REDIR_TIME64 = 1
const m__STDC_PREDEF_H = 1
const m__TANDEM_SOURCE = 1
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
const m_i386 = 1
const m_linux = 1
const m_static_assert = "_Static_assert"
const m_unix = 1

type t__builtin_va_list = uintptr

type t__predefined_size_t = uint32

type t__predefined_wchar_t = int32

type t__predefined_ptrdiff_t = int32

type Tsize_t = uint32

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

type TMD2_CTX = struct {
	Fi Tuint32_t
	FC [16]uint8
	FX [48]uint8
}

type TMD2Context = TMD2_CTX

type Tlocale_t = uintptr

// C documentation
//
//	/* cut-n-pasted from rfc1319 */
var _S = [256]uint8{
	0:   uint8(41),
	1:   uint8(46),
	2:   uint8(67),
	3:   uint8(201),
	4:   uint8(162),
	5:   uint8(216),
	6:   uint8(124),
	7:   uint8(1),
	8:   uint8(61),
	9:   uint8(54),
	10:  uint8(84),
	11:  uint8(161),
	12:  uint8(236),
	13:  uint8(240),
	14:  uint8(6),
	15:  uint8(19),
	16:  uint8(98),
	17:  uint8(167),
	18:  uint8(5),
	19:  uint8(243),
	20:  uint8(192),
	21:  uint8(199),
	22:  uint8(115),
	23:  uint8(140),
	24:  uint8(152),
	25:  uint8(147),
	26:  uint8(43),
	27:  uint8(217),
	28:  uint8(188),
	29:  uint8(76),
	30:  uint8(130),
	31:  uint8(202),
	32:  uint8(30),
	33:  uint8(155),
	34:  uint8(87),
	35:  uint8(60),
	36:  uint8(253),
	37:  uint8(212),
	38:  uint8(224),
	39:  uint8(22),
	40:  uint8(103),
	41:  uint8(66),
	42:  uint8(111),
	43:  uint8(24),
	44:  uint8(138),
	45:  uint8(23),
	46:  uint8(229),
	47:  uint8(18),
	48:  uint8(190),
	49:  uint8(78),
	50:  uint8(196),
	51:  uint8(214),
	52:  uint8(218),
	53:  uint8(158),
	54:  uint8(222),
	55:  uint8(73),
	56:  uint8(160),
	57:  uint8(251),
	58:  uint8(245),
	59:  uint8(142),
	60:  uint8(187),
	61:  uint8(47),
	62:  uint8(238),
	63:  uint8(122),
	64:  uint8(169),
	65:  uint8(104),
	66:  uint8(121),
	67:  uint8(145),
	68:  uint8(21),
	69:  uint8(178),
	70:  uint8(7),
	71:  uint8(63),
	72:  uint8(148),
	73:  uint8(194),
	74:  uint8(16),
	75:  uint8(137),
	76:  uint8(11),
	77:  uint8(34),
	78:  uint8(95),
	79:  uint8(33),
	80:  uint8(128),
	81:  uint8(127),
	82:  uint8(93),
	83:  uint8(154),
	84:  uint8(90),
	85:  uint8(144),
	86:  uint8(50),
	87:  uint8(39),
	88:  uint8(53),
	89:  uint8(62),
	90:  uint8(204),
	91:  uint8(231),
	92:  uint8(191),
	93:  uint8(247),
	94:  uint8(151),
	95:  uint8(3),
	96:  uint8(255),
	97:  uint8(25),
	98:  uint8(48),
	99:  uint8(179),
	100: uint8(72),
	101: uint8(165),
	102: uint8(181),
	103: uint8(209),
	104: uint8(215),
	105: uint8(94),
	106: uint8(146),
	107: uint8(42),
	108: uint8(172),
	109: uint8(86),
	110: uint8(170),
	111: uint8(198),
	112: uint8(79),
	113: uint8(184),
	114: uint8(56),
	115: uint8(210),
	116: uint8(150),
	117: uint8(164),
	118: uint8(125),
	119: uint8(182),
	120: uint8(118),
	121: uint8(252),
	122: uint8(107),
	123: uint8(226),
	124: uint8(156),
	125: uint8(116),
	126: uint8(4),
	127: uint8(241),
	128: uint8(69),
	129: uint8(157),
	130: uint8(112),
	131: uint8(89),
	132: uint8(100),
	133: uint8(113),
	134: uint8(135),
	135: uint8(32),
	136: uint8(134),
	137: uint8(91),
	138: uint8(207),
	139: uint8(101),
	140: uint8(230),
	141: uint8(45),
	142: uint8(168),
	143: uint8(2),
	144: uint8(27),
	145: uint8(96),
	146: uint8(37),
	147: uint8(173),
	148: uint8(174),
	149: uint8(176),
	150: uint8(185),
	151: uint8(246),
	152: uint8(28),
	153: uint8(70),
	154: uint8(97),
	155: uint8(105),
	156: uint8(52),
	157: uint8(64),
	158: uint8(126),
	159: uint8(15),
	160: uint8(85),
	161: uint8(71),
	162: uint8(163),
	163: uint8(35),
	164: uint8(221),
	165: uint8(81),
	166: uint8(175),
	167: uint8(58),
	168: uint8(195),
	169: uint8(92),
	170: uint8(249),
	171: uint8(206),
	172: uint8(186),
	173: uint8(197),
	174: uint8(234),
	175: uint8(38),
	176: uint8(44),
	177: uint8(83),
	178: uint8(13),
	179: uint8(110),
	180: uint8(133),
	181: uint8(40),
	182: uint8(132),
	183: uint8(9),
	184: uint8(211),
	185: uint8(223),
	186: uint8(205),
	187: uint8(244),
	188: uint8(65),
	189: uint8(129),
	190: uint8(77),
	191: uint8(82),
	192: uint8(106),
	193: uint8(220),
	194: uint8(55),
	195: uint8(200),
	196: uint8(108),
	197: uint8(193),
	198: uint8(171),
	199: uint8(250),
	200: uint8(36),
	201: uint8(225),
	202: uint8(123),
	203: uint8(8),
	204: uint8(12),
	205: uint8(189),
	206: uint8(177),
	207: uint8(74),
	208: uint8(120),
	209: uint8(136),
	210: uint8(149),
	211: uint8(139),
	212: uint8(227),
	213: uint8(99),
	214: uint8(232),
	215: uint8(109),
	216: uint8(233),
	217: uint8(203),
	218: uint8(213),
	219: uint8(254),
	220: uint8(59),
	222: uint8(29),
	223: uint8(57),
	224: uint8(242),
	225: uint8(239),
	226: uint8(183),
	227: uint8(14),
	228: uint8(102),
	229: uint8(88),
	230: uint8(208),
	231: uint8(228),
	232: uint8(166),
	233: uint8(119),
	234: uint8(114),
	235: uint8(248),
	236: uint8(235),
	237: uint8(117),
	238: uint8(75),
	239: uint8(10),
	240: uint8(49),
	241: uint8(68),
	242: uint8(80),
	243: uint8(180),
	244: uint8(143),
	245: uint8(237),
	246: uint8(31),
	247: uint8(26),
	248: uint8(219),
	249: uint8(153),
	250: uint8(141),
	251: uint8(51),
	252: uint8(159),
	253: uint8(17),
	254: uint8(131),
	255: uint8(20),
}

// C documentation
//
//	/* cut-n-pasted from rfc1319 */
var _pad = [17]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 1,
	2:  __ccgo_ts + 3,
	3:  __ccgo_ts + 6,
	4:  __ccgo_ts + 10,
	5:  __ccgo_ts + 15,
	6:  __ccgo_ts + 21,
	7:  __ccgo_ts + 28,
	8:  __ccgo_ts + 36,
	9:  __ccgo_ts + 45,
	10: __ccgo_ts + 55,
	11: __ccgo_ts + 66,
	12: __ccgo_ts + 78,
	13: __ccgo_ts + 91,
	14: __ccgo_ts + 105,
	15: __ccgo_ts + 120,
	16: __ccgo_ts + 136,
}

func XMD2Init(tls *libc.TLS, context uintptr) {
	(*TMD2_CTX)(unsafe.Pointer(context)).Fi = uint32(16)
	libc.Xmemset(tls, context+4, 0, uint32(16))
	libc.Xmemset(tls, context+20, 0, uint32(48))
}

func XMD2Update(tls *libc.TLS, context uintptr, input uintptr, inputLen uint32) {
	var idx, piece uint32
	var p2 uintptr
	_, _, _ = idx, piece, p2
	idx = uint32(0)
	for {
		if !(idx < inputLen) {
			break
		}
		piece = uint32(32) - (*TMD2_CTX)(unsafe.Pointer(context)).Fi
		if inputLen-idx < piece {
			piece = inputLen - idx
		}
		libc.Xmemcpy(tls, context+20+uintptr((*TMD2_CTX)(unsafe.Pointer(context)).Fi), input+uintptr(idx), piece)
		p2 = context
		*(*Tuint32_t)(unsafe.Pointer(p2)) += piece
		if *(*Tuint32_t)(unsafe.Pointer(p2)) == uint32(32) {
			XMD2Transform(tls, context)
		} /* resets i */
		goto _1
	_1:
		;
		idx += piece
	}
}

func XMD2Final(tls *libc.TLS, digest uintptr, context uintptr) {
	var padlen uint32
	_ = padlen
	/* padlen should be 1..16 */
	padlen = uint32(32) - (*TMD2_CTX)(unsafe.Pointer(context)).Fi
	/* add padding */
	XMD2Update(tls, context, _pad[padlen], padlen)
	/* add checksum */
	XMD2Update(tls, context, context+4, libc.Uint32FromInt64(16))
	/* copy out final digest */
	libc.Xmemcpy(tls, digest, context+20, libc.Uint32FromInt32(16))
	/* reset the context */
	XMD2Init(tls, context)
}

// C documentation
//
//	/*static*/
func XMD2Transform(tls *libc.TLS, context uintptr) {
	var j, k, l, t, v4 Tuint32_t
	var v6 uint8
	var p2 uintptr
	_, _, _, _, _, _, _ = j, k, l, t, v4, v6, p2
	/* set block "3" and update "checksum" */
	l = uint32(*(*uint8)(unsafe.Pointer(context + 4 + 15)))
	j = libc.Uint32FromInt32(0)
	for {
		if !(j < uint32(16)) {
			break
		}
		*(*uint8)(unsafe.Pointer(context + 20 + uintptr(uint32(32)+j))) = libc.Uint8FromInt32(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(context + 20 + uintptr(j)))) ^ libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(context + 20 + uintptr(uint32(16)+j)))))
		p2 = context + 4 + uintptr(j)
		*(*uint8)(unsafe.Pointer(p2)) = uint8(int32(*(*uint8)(unsafe.Pointer(p2))) ^ libc.Int32FromUint8(_S[uint32(*(*uint8)(unsafe.Pointer(context + 20 + uintptr(uint32(16)+j))))^l]))
		l = uint32(*(*uint8)(unsafe.Pointer(p2)))
		goto _1
	_1:
		;
		j++
	}
	/* mangle input block */
	v4 = libc.Uint32FromInt32(0)
	j = v4
	t = v4
	for {
		if !(j < uint32(18)) {
			break
		}
		k = uint32(0)
		for {
			if !(k < uint32(48)) {
				break
			}
			v6 = libc.Uint8FromInt32(libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(context + 20 + uintptr(k)))) ^ libc.Int32FromUint8(_S[t]))
			*(*uint8)(unsafe.Pointer(context + 20 + uintptr(k))) = v6
			t = uint32(v6)
			goto _5
		_5:
			;
			k++
		}
		goto _3
	_3:
		;
		t = (t + j) % uint32(256)
		j++
	}
	/* reset input pointer */
	(*TMD2_CTX)(unsafe.Pointer(context)).Fi = uint32(16)
}

const m_MD4_BLOCK_LENGTH = 64
const m_MD4_DIGEST_LENGTH = 16

type TMD4_CTX = struct {
	Fstate  [4]Tuint32_t
	Fcount  Tuint64_t
	Fbuffer [64]Tuint8_t
}

type TMD4Context = TMD4_CTX

var _PADDING = [64]Tuint8_t{
	0: uint8(0x80),
}

// C documentation
//
//	/*
//	 * Start MD4 accumulation.
//	 * Set bit count to 0 and buffer to mysterious initialization constants.
//	 */
func XMD4Init(tls *libc.TLS, ctx uintptr) {
	(*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount = uint64(0)
	*(*Tuint32_t)(unsafe.Pointer(ctx)) = uint32(0x67452301)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 1*4)) = uint32(0xefcdab89)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 2*4)) = uint32(0x98badcfe)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 3*4)) = uint32(0x10325476)
}

// C documentation
//
//	/*
//	 * Update context to reflect the concatenation of another buffer full
//	 * of bytes.
//	 */
func XMD4Update(tls *libc.TLS, ctx uintptr, input uintptr, len1 Tsize_t) {
	var have, need Tsize_t
	_, _ = have, need
	/* Check how many bytes we already have and how many more we need. */
	have = uint32((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> libc.Int32FromInt32(3) & libc.Uint64FromInt32(libc.Int32FromInt32(m_MD4_BLOCK_LENGTH)-libc.Int32FromInt32(1)))
	need = uint32(m_MD4_BLOCK_LENGTH) - have
	/* Update bitcount */
	*(*Tuint64_t)(unsafe.Pointer(ctx + 16)) += uint64(len1) << int32(3)
	if len1 >= need {
		if have != uint32(0) {
			libc.Xmemcpy(tls, ctx+24+uintptr(have), input, need)
			XMD4Transform(tls, ctx, ctx+24)
			input += uintptr(need)
			len1 -= need
			have = uint32(0)
		}
		/* Process data in MD4_BLOCK_LENGTH-byte chunks. */
		for len1 >= uint32(m_MD4_BLOCK_LENGTH) {
			XMD4Transform(tls, ctx, input)
			input += uintptr(m_MD4_BLOCK_LENGTH)
			len1 -= uint32(m_MD4_BLOCK_LENGTH)
		}
	}
	/* Handle any remaining bytes of data. */
	if len1 != uint32(0) {
		libc.Xmemcpy(tls, ctx+24+uintptr(have), input, len1)
	}
}

// C documentation
//
//	/*
//	 * Pad pad to 64-byte boundary with the bit pattern
//	 * 1 0* (64-bit count of bits processed, MSB-first)
//	 */
func XMD4Pad(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var padlen Tsize_t
	var _ /* count at bp+0 */ [8]Tuint8_t
	_ = padlen
	/* Convert count to 8 bytes in little endian order. */
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(7)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(56))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(6)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(48))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(5)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(40))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(4)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(32))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(3)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(24))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(2)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(16))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(1)] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(8))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[0] = uint8((*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount)
	/* Pad out to 56 mod 64. */
	padlen = uint32(uint64(m_MD4_BLOCK_LENGTH) - (*TMD4_CTX)(unsafe.Pointer(ctx)).Fcount>>libc.Int32FromInt32(3)&libc.Uint64FromInt32(libc.Int32FromInt32(m_MD4_BLOCK_LENGTH)-libc.Int32FromInt32(1)))
	if padlen < libc.Uint32FromInt32(libc.Int32FromInt32(1)+libc.Int32FromInt32(8)) {
		padlen += uint32(m_MD4_BLOCK_LENGTH)
	}
	XMD4Update(tls, ctx, uintptr(unsafe.Pointer(&_PADDING)), padlen-uint32(8)) /* padlen - 8 <= 64 */
	XMD4Update(tls, ctx, bp, uint32(8))
}

// C documentation
//
//	/*
//	 * Final wrapup--call MD4Pad, fill in digest and zero out ctx.
//	 */
func XMD4Final(tls *libc.TLS, digest uintptr, ctx uintptr) {
	var i int32
	_ = i
	XMD4Pad(tls, ctx)
	if digest != libc.UintptrFromInt32(0) {
		i = 0
		for {
			if !(i < int32(4)) {
				break
			}
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 3)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(24))
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 2)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(16))
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 1)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(8))
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)))) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)))
			goto _1
		_1:
			;
			i++
		}
		libc.Xmemset(tls, ctx, 0, uint32(88))
	}
}

/* The three core functions - F1 is optimized somewhat */

/* #define F1(x, y, z) (x & y | ~x & z) */

/* This is the central step in the MD4 algorithm. */

// C documentation
//
//	/*
//	 * The core of the MD4 algorithm, this alters an existing MD4 hash to
//	 * reflect the addition of 16 longwords of new data.  MD4Update blocks
//	 * the data and converts bytes into longwords for this routine.
//	 */
func XMD4Transform(tls *libc.TLS, state uintptr, block uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var a, b, c, d Tuint32_t
	var _ /* in at bp+0 */ [16]Tuint32_t
	_, _, _, _ = a, b, c, d
	libc.Xmemcpy(tls, bp, block, uint32(64))
	a = *(*Tuint32_t)(unsafe.Pointer(state))
	b = *(*Tuint32_t)(unsafe.Pointer(state + 1*4))
	c = *(*Tuint32_t)(unsafe.Pointer(state + 2*4))
	d = *(*Tuint32_t)(unsafe.Pointer(state + 3*4))
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]
	d = d<<libc.Int32FromInt32(7) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]
	b = b<<libc.Int32FromInt32(19) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(19))
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]
	d = d<<libc.Int32FromInt32(7) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]
	b = b<<libc.Int32FromInt32(19) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(19))
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]
	d = d<<libc.Int32FromInt32(7) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]
	b = b<<libc.Int32FromInt32(19) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(19))
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]
	d = d<<libc.Int32FromInt32(7) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]
	b = b<<libc.Int32FromInt32(19) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(19))
	a += b&c | b&d | c&d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0] + uint32(0x5a827999)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a&b | a&c | b&c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)] + uint32(0x5a827999)
	d = d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	c += d&a | d&b | a&b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)] + uint32(0x5a827999)
	c = c<<libc.Int32FromInt32(9) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	b += c&d | c&a | d&a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)] + uint32(0x5a827999)
	b = b<<libc.Int32FromInt32(13) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13))
	a += b&c | b&d | c&d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] + uint32(0x5a827999)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a&b | a&c | b&c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)] + uint32(0x5a827999)
	d = d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	c += d&a | d&b | a&b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)] + uint32(0x5a827999)
	c = c<<libc.Int32FromInt32(9) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	b += c&d | c&a | d&a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)] + uint32(0x5a827999)
	b = b<<libc.Int32FromInt32(13) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13))
	a += b&c | b&d | c&d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)] + uint32(0x5a827999)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a&b | a&c | b&c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)] + uint32(0x5a827999)
	d = d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	c += d&a | d&b | a&b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)] + uint32(0x5a827999)
	c = c<<libc.Int32FromInt32(9) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	b += c&d | c&a | d&a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)] + uint32(0x5a827999)
	b = b<<libc.Int32FromInt32(13) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13))
	a += b&c | b&d | c&d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)] + uint32(0x5a827999)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a&b | a&c | b&c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)] + uint32(0x5a827999)
	d = d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	c += d&a | d&b | a&b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)] + uint32(0x5a827999)
	c = c<<libc.Int32FromInt32(9) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	b += c&d | c&a | d&a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)] + uint32(0x5a827999)
	b = b<<libc.Int32FromInt32(13) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13))
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0] + uint32(0x6ed9eba1)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)] + uint32(0x6ed9eba1)
	d = d<<libc.Int32FromInt32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)] + uint32(0x6ed9eba1)
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)] + uint32(0x6ed9eba1)
	b = b<<libc.Int32FromInt32(15) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)] + uint32(0x6ed9eba1)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)] + uint32(0x6ed9eba1)
	d = d<<libc.Int32FromInt32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)] + uint32(0x6ed9eba1)
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)] + uint32(0x6ed9eba1)
	b = b<<libc.Int32FromInt32(15) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] + uint32(0x6ed9eba1)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)] + uint32(0x6ed9eba1)
	d = d<<libc.Int32FromInt32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)] + uint32(0x6ed9eba1)
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)] + uint32(0x6ed9eba1)
	b = b<<libc.Int32FromInt32(15) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)] + uint32(0x6ed9eba1)
	a = a<<libc.Int32FromInt32(3) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(3))
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)] + uint32(0x6ed9eba1)
	d = d<<libc.Int32FromInt32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)] + uint32(0x6ed9eba1)
	c = c<<libc.Int32FromInt32(11) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)] + uint32(0x6ed9eba1)
	b = b<<libc.Int32FromInt32(15) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	*(*Tuint32_t)(unsafe.Pointer(state)) += a
	*(*Tuint32_t)(unsafe.Pointer(state + 1*4)) += b
	*(*Tuint32_t)(unsafe.Pointer(state + 2*4)) += c
	*(*Tuint32_t)(unsafe.Pointer(state + 3*4)) += d
}

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

var _PADDING1 = [64]Tuint8_t{
	0: uint8(0x80),
}

// C documentation
//
//	/*
//	 * Start MD5 accumulation.  Set bit count to 0 and buffer to mysterious
//	 * initialization constants.
//	 */
func XMD5Init(tls *libc.TLS, ctx uintptr) {
	(*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount = uint64(0)
	*(*Tuint32_t)(unsafe.Pointer(ctx)) = uint32(0x67452301)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 1*4)) = uint32(0xefcdab89)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 2*4)) = uint32(0x98badcfe)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 3*4)) = uint32(0x10325476)
}

// C documentation
//
//	/*
//	 * Update context to reflect the concatenation of another buffer full
//	 * of bytes.
//	 */
func XMD5Update(tls *libc.TLS, ctx uintptr, input uintptr, len1 Tsize_t) {
	var have, need Tsize_t
	_, _ = have, need
	/* Check how many bytes we already have and how many more we need. */
	have = uint32((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> libc.Int32FromInt32(3) & libc.Uint64FromInt32(libc.Int32FromInt32(m_MD5_BLOCK_LENGTH)-libc.Int32FromInt32(1)))
	need = uint32(m_MD5_BLOCK_LENGTH) - have
	/* Update bitcount */
	*(*Tuint64_t)(unsafe.Pointer(ctx + 16)) += uint64(len1) << int32(3)
	if len1 >= need {
		if have != uint32(0) {
			libc.Xmemcpy(tls, ctx+24+uintptr(have), input, need)
			XMD5Transform(tls, ctx, ctx+24)
			input += uintptr(need)
			len1 -= need
			have = uint32(0)
		}
		/* Process data in MD5_BLOCK_LENGTH-byte chunks. */
		for len1 >= uint32(m_MD5_BLOCK_LENGTH) {
			XMD5Transform(tls, ctx, input)
			input += uintptr(m_MD5_BLOCK_LENGTH)
			len1 -= uint32(m_MD5_BLOCK_LENGTH)
		}
	}
	/* Handle any remaining bytes of data. */
	if len1 != uint32(0) {
		libc.Xmemcpy(tls, ctx+24+uintptr(have), input, len1)
	}
}

// C documentation
//
//	/*
//	 * Pad pad to 64-byte boundary with the bit pattern
//	 * 1 0* (64-bit count of bits processed, MSB-first)
//	 */
func XMD5Pad(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var padlen Tsize_t
	var _ /* count at bp+0 */ [8]Tuint8_t
	_ = padlen
	/* Convert count to 8 bytes in little endian order. */
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(7)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(56))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(6)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(48))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(5)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(40))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(4)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(32))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(3)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(24))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(2)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(16))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(1)] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(8))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[0] = uint8((*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount)
	/* Pad out to 56 mod 64. */
	padlen = uint32(uint64(m_MD5_BLOCK_LENGTH) - (*TMD5_CTX)(unsafe.Pointer(ctx)).Fcount>>libc.Int32FromInt32(3)&libc.Uint64FromInt32(libc.Int32FromInt32(m_MD5_BLOCK_LENGTH)-libc.Int32FromInt32(1)))
	if padlen < libc.Uint32FromInt32(libc.Int32FromInt32(1)+libc.Int32FromInt32(8)) {
		padlen += uint32(m_MD5_BLOCK_LENGTH)
	}
	XMD5Update(tls, ctx, uintptr(unsafe.Pointer(&_PADDING1)), padlen-uint32(8)) /* padlen - 8 <= 64 */
	XMD5Update(tls, ctx, bp, uint32(8))
}

// C documentation
//
//	/*
//	 * Final wrapup--call MD5Pad, fill in digest and zero out ctx.
//	 */
func XMD5Final(tls *libc.TLS, digest uintptr, ctx uintptr) {
	var i int32
	_ = i
	XMD5Pad(tls, ctx)
	if digest != libc.UintptrFromInt32(0) {
		i = 0
		for {
			if !(i < int32(4)) {
				break
			}
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 3)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(24))
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 2)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(16))
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 1)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(8))
			*(*uint8)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)))) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)))
			goto _1
		_1:
			;
			i++
		}
		libc.Xmemset(tls, ctx, 0, uint32(88))
	}
}

/* The four core functions - F1 is optimized somewhat */

/* #define F1(x, y, z) (x & y | ~x & z) */

/* This is the central step in the MD5 algorithm. */

// C documentation
//
//	/*
//	 * The core of the MD5 algorithm, this alters an existing MD5 hash to
//	 * reflect the addition of 16 longwords of new data.  MD5Update blocks
//	 * the data and converts bytes into longwords for this routine.
//	 */
func XMD5Transform(tls *libc.TLS, state uintptr, block uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var a, b, c, d Tuint32_t
	var _ /* in at bp+0 */ [16]Tuint32_t
	_, _, _, _ = a, b, c, d
	libc.Xmemcpy(tls, bp, block, uint32(64))
	a = *(*Tuint32_t)(unsafe.Pointer(state))
	b = *(*Tuint32_t)(unsafe.Pointer(state + 1*4))
	c = *(*Tuint32_t)(unsafe.Pointer(state + 2*4))
	d = *(*Tuint32_t)(unsafe.Pointer(state + 3*4))
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0] + uint32(0xd76aa478)
	a = a<<int32(7) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	a += b
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] + uint32(0xe8c7b756)
	d = d<<int32(12) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
	d += a
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)] + uint32(0x242070db)
	c = c<<int32(17) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(17))
	c += d
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)] + uint32(0xc1bdceee)
	b = b<<int32(22) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(22))
	b += c
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)] + uint32(0xf57c0faf)
	a = a<<int32(7) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	a += b
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)] + uint32(0x4787c62a)
	d = d<<int32(12) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
	d += a
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)] + uint32(0xa8304613)
	c = c<<int32(17) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(17))
	c += d
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)] + uint32(0xfd469501)
	b = b<<int32(22) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(22))
	b += c
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)] + uint32(0x698098d8)
	a = a<<int32(7) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	a += b
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)] + uint32(0x8b44f7af)
	d = d<<int32(12) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
	d += a
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)] + uint32(0xffff5bb1)
	c = c<<int32(17) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(17))
	c += d
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)] + uint32(0x895cd7be)
	b = b<<int32(22) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(22))
	b += c
	a += d ^ b&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)] + uint32(0x6b901122)
	a = a<<int32(7) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7))
	a += b
	d += c ^ a&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)] + uint32(0xfd987193)
	d = d<<int32(12) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12))
	d += a
	c += b ^ d&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)] + uint32(0xa679438e)
	c = c<<int32(17) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(17))
	c += d
	b += a ^ c&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)] + uint32(0x49b40821)
	b = b<<int32(22) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(22))
	b += c
	a += c ^ d&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] + uint32(0xf61e2562)
	a = a<<int32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	a += b
	d += b ^ c&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)] + uint32(0xc040b340)
	d = d<<int32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	d += a
	c += a ^ b&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)] + uint32(0x265e5a51)
	c = c<<int32(14) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14))
	c += d
	b += d ^ a&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0] + uint32(0xe9b6c7aa)
	b = b<<int32(20) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(20))
	b += c
	a += c ^ d&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)] + uint32(0xd62f105d)
	a = a<<int32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	a += b
	d += b ^ c&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)] + uint32(0x02441453)
	d = d<<int32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	d += a
	c += a ^ b&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)] + uint32(0xd8a1e681)
	c = c<<int32(14) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14))
	c += d
	b += d ^ a&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)] + uint32(0xe7d3fbc8)
	b = b<<int32(20) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(20))
	b += c
	a += c ^ d&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)] + uint32(0x21e1cde6)
	a = a<<int32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	a += b
	d += b ^ c&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)] + uint32(0xc33707d6)
	d = d<<int32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	d += a
	c += a ^ b&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)] + uint32(0xf4d50d87)
	c = c<<int32(14) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14))
	c += d
	b += d ^ a&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)] + uint32(0x455a14ed)
	b = b<<int32(20) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(20))
	b += c
	a += c ^ d&(b^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)] + uint32(0xa9e3e905)
	a = a<<int32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5))
	a += b
	d += b ^ c&(a^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)] + uint32(0xfcefa3f8)
	d = d<<int32(9) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9))
	d += a
	c += a ^ b&(d^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)] + uint32(0x676f02d9)
	c = c<<int32(14) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14))
	c += d
	b += d ^ a&(c^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)] + uint32(0x8d2a4c8a)
	b = b<<int32(20) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(20))
	b += c
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)] + uint32(0xfffa3942)
	a = a<<int32(4) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(4))
	a += b
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)] + uint32(0x8771f681)
	d = d<<int32(11) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	d += a
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)] + uint32(0x6d9d6122)
	c = c<<int32(16) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
	c += d
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)] + uint32(0xfde5380c)
	b = b<<int32(23) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(23))
	b += c
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] + uint32(0xa4beea44)
	a = a<<int32(4) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(4))
	a += b
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)] + uint32(0x4bdecfa9)
	d = d<<int32(11) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	d += a
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)] + uint32(0xf6bb4b60)
	c = c<<int32(16) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
	c += d
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)] + uint32(0xbebfbc70)
	b = b<<int32(23) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(23))
	b += c
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)] + uint32(0x289b7ec6)
	a = a<<int32(4) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(4))
	a += b
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0] + uint32(0xeaa127fa)
	d = d<<int32(11) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	d += a
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)] + uint32(0xd4ef3085)
	c = c<<int32(16) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
	c += d
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)] + uint32(0x04881d05)
	b = b<<int32(23) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(23))
	b += c
	a += b ^ c ^ d + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)] + uint32(0xd9d4d039)
	a = a<<int32(4) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(4))
	a += b
	d += a ^ b ^ c + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)] + uint32(0xe6db99e5)
	d = d<<int32(11) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))
	d += a
	c += d ^ a ^ b + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)] + uint32(0x1fa27cf8)
	c = c<<int32(16) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(16))
	c += d
	b += c ^ d ^ a + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)] + uint32(0xc4ac5665)
	b = b<<int32(23) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(23))
	b += c
	a += c ^ (b | ^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0] + uint32(0xf4292244)
	a = a<<int32(6) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6))
	a += b
	d += b ^ (a | ^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)] + uint32(0x432aff97)
	d = d<<int32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d += a
	c += a ^ (d | ^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)] + uint32(0xab9423a7)
	c = c<<int32(15) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	c += d
	b += d ^ (c | ^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)] + uint32(0xfc93a039)
	b = b<<int32(21) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(21))
	b += c
	a += c ^ (b | ^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)] + uint32(0x655b59c3)
	a = a<<int32(6) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6))
	a += b
	d += b ^ (a | ^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)] + uint32(0x8f0ccc92)
	d = d<<int32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d += a
	c += a ^ (d | ^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)] + uint32(0xffeff47d)
	c = c<<int32(15) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	c += d
	b += d ^ (c | ^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)] + uint32(0x85845dd1)
	b = b<<int32(21) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(21))
	b += c
	a += c ^ (b | ^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)] + uint32(0x6fa87e4f)
	a = a<<int32(6) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6))
	a += b
	d += b ^ (a | ^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)] + uint32(0xfe2ce6e0)
	d = d<<int32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d += a
	c += a ^ (d | ^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)] + uint32(0xa3014314)
	c = c<<int32(15) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	c += d
	b += d ^ (c | ^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)] + uint32(0x4e0811a1)
	b = b<<int32(21) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(21))
	b += c
	a += c ^ (b | ^d) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)] + uint32(0xf7537e82)
	a = a<<int32(6) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6))
	a += b
	d += b ^ (a | ^c) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)] + uint32(0xbd3af235)
	d = d<<int32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d += a
	c += a ^ (d | ^b) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)] + uint32(0x2ad7d2bb)
	c = c<<int32(15) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15))
	c += d
	b += d ^ (c | ^a) + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)] + uint32(0xeb86d391)
	b = b<<int32(21) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(21))
	b += c
	*(*Tuint32_t)(unsafe.Pointer(state)) += a
	*(*Tuint32_t)(unsafe.Pointer(state + 1*4)) += b
	*(*Tuint32_t)(unsafe.Pointer(state + 2*4)) += c
	*(*Tuint32_t)(unsafe.Pointer(state + 3*4)) += d
}

const m_H0 = 1732584193
const m_H1 = 4023233417
const m_H2 = 2562383102
const m_H3 = 271733878
const m_H4 = 3285377520
const m_K0 = 0
const m_K1 = 1518500249
const m_K2 = 1859775393
const m_K3 = 2400959708
const m_K4 = 2840853838
const m_KK0 = 1352829926
const m_KK1 = 1548603684
const m_KK2 = 1836072691
const m_KK3 = 2053994217
const m_KK4 = 0
const m_RMD160_BLOCK_LENGTH = 64
const m_RMD160_DIGEST_LENGTH = 20

type TRMD160_CTX = struct {
	Fstate  [5]Tuint32_t
	Fcount  Tuint64_t
	Fbuffer [64]Tuint8_t
}

type TRMD160Context = TRMD160_CTX

/* rotate x left n bits.  */

var _PADDING2 = [64]Tuint8_t{
	0: uint8(0x80),
}

func XRMD160Init(tls *libc.TLS, ctx uintptr) {
	(*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount = uint64(0)
	*(*Tuint32_t)(unsafe.Pointer(ctx)) = uint32(0x67452301)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 1*4)) = uint32(0xEFCDAB89)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 2*4)) = uint32(0x98BADCFE)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 3*4)) = uint32(0x10325476)
	*(*Tuint32_t)(unsafe.Pointer(ctx + 4*4)) = uint32(0xC3D2E1F0)
}

func XRMD160Update(tls *libc.TLS, ctx uintptr, input uintptr, len1 Tsize_t) {
	var have, need, off Tsize_t
	_, _, _ = have, need, off
	have = uint32((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount / uint64(8) % uint64(m_RMD160_BLOCK_LENGTH))
	need = uint32(m_RMD160_BLOCK_LENGTH) - have
	*(*Tuint64_t)(unsafe.Pointer(ctx + 20)) += uint64(uint32(8) * len1)
	off = uint32(0)
	if len1 >= need {
		if have != 0 {
			libc.Xmemcpy(tls, ctx+28+uintptr(have), input, need)
			XRMD160Transform(tls, ctx, ctx+28)
			off = need
			have = uint32(0)
		}
		/* now the buffer is empty */
		for off+uint32(m_RMD160_BLOCK_LENGTH) <= len1 {
			XRMD160Transform(tls, ctx, input+uintptr(off))
			off += uint32(m_RMD160_BLOCK_LENGTH)
		}
	}
	if off < len1 {
		libc.Xmemcpy(tls, ctx+28+uintptr(have), input+uintptr(off), len1-off)
	}
}

func XRMD160Pad(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var padlen Tsize_t
	var _ /* size at bp+0 */ [8]Tuint8_t
	_ = padlen
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(7)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(56))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(6)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(48))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(5)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(40))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(4)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(32))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(3)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(24))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(2)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(16))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[int32(1)] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount >> int32(8))
	(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[0] = uint8((*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount)
	/*
	 * pad to RMD160_BLOCK_LENGTH byte blocks, at least one byte from
	 * PADDING plus 8 bytes for the size
	 */
	padlen = uint32(uint64(m_RMD160_BLOCK_LENGTH) - (*TRMD160_CTX)(unsafe.Pointer(ctx)).Fcount/uint64(8)%uint64(m_RMD160_BLOCK_LENGTH))
	if padlen < libc.Uint32FromInt32(libc.Int32FromInt32(1)+libc.Int32FromInt32(8)) {
		padlen += uint32(m_RMD160_BLOCK_LENGTH)
	}
	XRMD160Update(tls, ctx, uintptr(unsafe.Pointer(&_PADDING2)), padlen-uint32(8)) /* padlen - 8 <= 64 */
	XRMD160Update(tls, ctx, bp, uint32(8))
}

func XRMD160Final(tls *libc.TLS, digest uintptr, ctx uintptr) {
	var i int32
	_ = i
	XRMD160Pad(tls, ctx)
	if digest != libc.UintptrFromInt32(0) {
		i = 0
		for {
			if !(i < int32(5)) {
				break
			}
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 3)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(24))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 2)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(16))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 1)) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)) >> int32(8))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)))) = uint8(*(*Tuint32_t)(unsafe.Pointer(ctx + uintptr(i)*4)))
			goto _1
		_1:
			;
			i++
		}
		libc.Xmemset(tls, ctx, 0, uint32(92))
	}
}

func XRMD160Transform(tls *libc.TLS, state uintptr, block uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var a, aa, b, bb, c, cc, d, dd, e, ee, t Tuint32_t
	var _ /* x at bp+0 */ [16]Tuint32_t
	_, _, _, _, _, _, _, _, _, _, _ = a, aa, b, bb, c, cc, d, dd, e, ee, t
	libc.Xmemcpy(tls, bp, block, uint32(m_RMD160_BLOCK_LENGTH))
	a = *(*Tuint32_t)(unsafe.Pointer(state))
	b = *(*Tuint32_t)(unsafe.Pointer(state + 1*4))
	c = *(*Tuint32_t)(unsafe.Pointer(state + 2*4))
	d = *(*Tuint32_t)(unsafe.Pointer(state + 3*4))
	e = *(*Tuint32_t)(unsafe.Pointer(state + 4*4))
	/* Round 1 */
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+0x00000000)<<int32(11) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+0x00000000)<<int32(14) | (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+0x00000000)<<int32(15) | (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+0x00000000)<<int32(12) | (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+0x00000000)<<int32(5) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+0x00000000)<<int32(8) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+0x00000000)<<int32(7) | (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+0x00000000)<<int32(9) | (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+0x00000000)<<int32(11) | (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+0x00000000)<<int32(13) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+0x00000000)<<int32(14) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+0x00000000)<<int32(15) | (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+0x00000000)<<int32(6) | (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+0x00000000)<<int32(7) | (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+0x00000000)<<int32(9) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+0x00000000)<<int32(8) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #15 */
	/* Round 2 */
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x5A827999))<<int32(7) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x5A827999))<<int32(6) | (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x5A827999))<<int32(8) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x5A827999))<<int32(13) | (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x5A827999))<<int32(11) | (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x5A827999))<<int32(9) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x5A827999))<<int32(7) | (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x5A827999))<<int32(15) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x5A827999))<<int32(7) | (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x5A827999))<<int32(12) | (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x5A827999))<<int32(15) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x5A827999))<<int32(9) | (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x5A827999))<<int32(11) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x5A827999))<<int32(7) | (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x5A827999))<<int32(13) | (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x5A827999))<<int32(12) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x5A827999))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #31 */
	/* Round 3 */
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x6ED9EBA1))<<int32(11) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x6ED9EBA1))<<int32(13) | (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x6ED9EBA1))<<int32(6) | (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x6ED9EBA1))<<int32(7) | (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x6ED9EBA1))<<int32(14) | (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x6ED9EBA1))<<int32(9) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x6ED9EBA1))<<int32(13) | (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x6ED9EBA1))<<int32(15) | (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x6ED9EBA1))<<int32(14) | (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x6ED9EBA1))<<int32(8) | (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x6ED9EBA1))<<int32(13) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x6ED9EBA1))<<int32(6) | (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x6ED9EBA1))<<int32(5) | (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x6ED9EBA1))<<int32(12) | (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x6ED9EBA1))<<int32(7) | (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x6ED9EBA1))<<int32(5) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x6ED9EBA1))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #47 */
	/* Round 4 */
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x8F1BBCDC))<<int32(11) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x8F1BBCDC))<<int32(12) | (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x8F1BBCDC))<<int32(14) | (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x8F1BBCDC))<<int32(15) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x8F1BBCDC))<<int32(14) | (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x8F1BBCDC))<<int32(15) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x8F1BBCDC))<<int32(9) | (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x8F1BBCDC))<<int32(8) | (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x8F1BBCDC))<<int32(9) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x8F1BBCDC))<<int32(14) | (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x8F1BBCDC))<<int32(5) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x8F1BBCDC))<<int32(6) | (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x8F1BBCDC))<<int32(8) | (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x8F1BBCDC))<<int32(6) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x8F1BBCDC))<<int32(5) | (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x8F1BBCDC))<<int32(12) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x8F1BBCDC))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #63 */
	/* Round 5 */
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0xA953FD4E))<<int32(9) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0xA953FD4E))<<int32(15) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0xA953FD4E))<<int32(5) | (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0xA953FD4E))<<int32(11) | (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0xA953FD4E))<<int32(6) | (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0xA953FD4E))<<int32(8) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0xA953FD4E))<<int32(13) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0xA953FD4E))<<int32(12) | (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0xA953FD4E))<<int32(5) | (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0xA953FD4E))<<int32(12) | (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0xA953FD4E))<<int32(13) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0xA953FD4E))<<int32(14) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0xA953FD4E))<<int32(11) | (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0xA953FD4E))<<int32(8) | (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0xA953FD4E))<<int32(5) | (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0xA953FD4E))<<int32(6) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0xA953FD4E))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #79 */
	aa = a
	bb = b
	cc = c
	dd = d
	ee = e
	a = *(*Tuint32_t)(unsafe.Pointer(state))
	b = *(*Tuint32_t)(unsafe.Pointer(state + 1*4))
	c = *(*Tuint32_t)(unsafe.Pointer(state + 2*4))
	d = *(*Tuint32_t)(unsafe.Pointer(state + 3*4))
	e = *(*Tuint32_t)(unsafe.Pointer(state + 4*4))
	/* Parallel round 1 */
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x50A28BE6))<<int32(8) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x50A28BE6))<<int32(9) | (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x50A28BE6))<<int32(9) | (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x50A28BE6))<<int32(11) | (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x50A28BE6))<<int32(13) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x50A28BE6))<<int32(15) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x50A28BE6))<<int32(15) | (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x50A28BE6))<<int32(5) | (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x50A28BE6))<<int32(7) | (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x50A28BE6))<<int32(7) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x50A28BE6))<<int32(8) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x50A28BE6))<<int32(11) | (e+(a^(b|^c))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x50A28BE6))<<int32(14) | (d+(e^(a|^b))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x50A28BE6))<<int32(14) | (c+(d^(e|^a))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x50A28BE6))<<int32(12) | (b+(c^(d|^e))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x50A28BE6))<<int32(6) | (a+(b^(c|^d))+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x50A28BE6))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #15 */
	/* Parallel round 2 */
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x5C4DD124))<<int32(9) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x5C4DD124))<<int32(13) | (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x5C4DD124))<<int32(15) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x5C4DD124))<<int32(7) | (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x5C4DD124))<<int32(12) | (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x5C4DD124))<<int32(8) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x5C4DD124))<<int32(9) | (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x5C4DD124))<<int32(11) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x5C4DD124))<<int32(7) | (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x5C4DD124))<<int32(7) | (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x5C4DD124))<<int32(12) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x5C4DD124))<<int32(7) | (d+(e&b | a & ^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x5C4DD124))<<int32(6) | (c+(d&a | e & ^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x5C4DD124))<<int32(15) | (b+(c&e | d & ^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x5C4DD124))<<int32(13) | (a+(b&d | c & ^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x5C4DD124))<<int32(11) | (e+(a&c | b & ^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x5C4DD124))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #31 */
	/* Parallel round 3 */
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x6D703EF3))<<int32(9) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x6D703EF3))<<int32(7) | (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x6D703EF3))<<int32(15) | (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x6D703EF3))<<int32(11) | (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x6D703EF3))<<int32(8) | (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x6D703EF3))<<int32(6) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x6D703EF3))<<int32(6) | (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x6D703EF3))<<int32(14) | (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x6D703EF3))<<int32(12) | (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x6D703EF3))<<int32(13) | (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x6D703EF3))<<int32(5) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x6D703EF3))<<int32(14) | (c+(d|^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x6D703EF3))<<int32(13) | (b+(c|^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x6D703EF3))<<int32(13) | (a+(b|^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x6D703EF3))<<int32(7) | (e+(a|^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x6D703EF3))<<int32(5) | (d+(e|^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x6D703EF3))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #47 */
	/* Parallel round 4 */
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x7A6D76E9))<<int32(15) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x7A6D76E9))<<int32(5) | (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x7A6D76E9))<<int32(8) | (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x7A6D76E9))<<int32(11) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x7A6D76E9))<<int32(14) | (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x7A6D76E9))<<int32(14) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x7A6D76E9))<<int32(6) | (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x7A6D76E9))<<int32(14) | (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x7A6D76E9))<<int32(6) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x7A6D76E9))<<int32(9) | (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x7A6D76E9))<<int32(12) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x7A6D76E9))<<int32(9) | (b+(c&d|^c&e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x7A6D76E9))<<int32(12) | (a+(b&c|^b&d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x7A6D76E9))<<int32(5) | (e+(a&b|^a&c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x7A6D76E9))<<int32(15) | (d+(e&a|^e&b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x7A6D76E9))<<int32(8) | (c+(d&e|^d&a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+uint32(0x7A6D76E9))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #63 */
	/* Parallel round 5 */
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+0x00000000)<<int32(8) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(12)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+0x00000000)<<int32(5) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(15)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+0x00000000)<<int32(12) | (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(10)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+0x00000000)<<int32(9) | (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(4)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(9)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+0x00000000)<<int32(12) | (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(1)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(12)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+0x00000000)<<int32(5) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(5)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+0x00000000)<<int32(14) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(8)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(14)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+0x00000000)<<int32(6) | (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(7)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+0x00000000)<<int32(8) | (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(6)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+0x00000000)<<int32(13) | (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(2)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+0x00000000)<<int32(6) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(13)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	a = (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+0x00000000)<<int32(5) | (a+(b^c^d)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(14)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)) + e
	c = c<<libc.Int32FromInt32(10) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	e = (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+0x00000000)<<int32(15) | (e+(a^b^c)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[0]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(15)) + d
	b = b<<libc.Int32FromInt32(10) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	d = (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+0x00000000)<<int32(13) | (d+(e^a^b)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(3)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(13)) + c
	a = a<<libc.Int32FromInt32(10) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	c = (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+0x00000000)<<int32(11) | (c+(d^e^a)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(9)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + b
	e = e<<libc.Int32FromInt32(10) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10))
	b = (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+0x00000000)<<int32(11) | (b+(c^d^e)+(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[int32(11)]+0x00000000)>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(11)) + a
	d = d<<libc.Int32FromInt32(10) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(10)) /* #79 */
	t = *(*Tuint32_t)(unsafe.Pointer(state + 1*4)) + cc + d
	*(*Tuint32_t)(unsafe.Pointer(state + 1*4)) = *(*Tuint32_t)(unsafe.Pointer(state + 2*4)) + dd + e
	*(*Tuint32_t)(unsafe.Pointer(state + 2*4)) = *(*Tuint32_t)(unsafe.Pointer(state + 3*4)) + ee + a
	*(*Tuint32_t)(unsafe.Pointer(state + 3*4)) = *(*Tuint32_t)(unsafe.Pointer(state + 4*4)) + aa + b
	*(*Tuint32_t)(unsafe.Pointer(state + 4*4)) = *(*Tuint32_t)(unsafe.Pointer(state)) + bb + c
	*(*Tuint32_t)(unsafe.Pointer(state)) = t
}

const m_SHA1_BLOCK_LENGTH = 64
const m_SHA1_DIGEST_LENGTH = 20

type TSHA1_CTX = struct {
	Fstate  [5]Tuint32_t
	Fcount  Tuint64_t
	Fbuffer [64]Tuint8_t
}

/*
 * blk0() and blk() perform the initial expand.
 * I got the idea of expanding during the round function from SSLeay
 */

/*
 * (R0+R1), R2, R3, R4 are the different operations (rounds) used in SHA1
 */

type TCHAR64LONG16 = struct {
	Fl [0][16]Tuint32_t
	Fc [64]Tuint8_t
}

// C documentation
//
//	/*
//	 * Hash a single 512-bit block. This is the core of the algorithm.
//	 */
func XSHA1Transform(tls *libc.TLS, state uintptr, buffer uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var a, b, c, d, e, v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v5, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v7, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v8, v80, v81, v82, v83, v84, v9 Tuint32_t
	var block uintptr
	var _ /* workspace at bp+0 */ [64]Tuint8_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a, b, block, c, d, e, v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v5, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v7, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v8, v80, v81, v82, v83, v84, v9
	block = bp
	libc.Xmemcpy(tls, block, buffer, uint32(m_SHA1_BLOCK_LENGTH))
	/* Copy context->state[] to working vars */
	a = *(*Tuint32_t)(unsafe.Pointer(state))
	b = *(*Tuint32_t)(unsafe.Pointer(state + 1*4))
	c = *(*Tuint32_t)(unsafe.Pointer(state + 2*4))
	d = *(*Tuint32_t)(unsafe.Pointer(state + 3*4))
	e = *(*Tuint32_t)(unsafe.Pointer(state + 4*4))
	/* 4 rounds of 20 operations each. Loop unrolled. */
	v1 = (*(*Tuint32_t)(unsafe.Pointer(block))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block)) = v1
	e += b&(c^d) ^ d + v1 + uint32(0x5A827999) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v2 = (*(*Tuint32_t)(unsafe.Pointer(block + 1*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 1*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 1*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 1*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 1*4)) = v2
	d += a&(b^c) ^ c + v2 + uint32(0x5A827999) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v3 = (*(*Tuint32_t)(unsafe.Pointer(block + 2*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 2*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 2*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 2*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 2*4)) = v3
	c += e&(a^b) ^ b + v3 + uint32(0x5A827999) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v4 = (*(*Tuint32_t)(unsafe.Pointer(block + 3*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 3*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 3*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 3*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 3*4)) = v4
	b += d&(e^a) ^ a + v4 + uint32(0x5A827999) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v5 = (*(*Tuint32_t)(unsafe.Pointer(block + 4*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 4*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 4*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 4*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 4*4)) = v5
	a += c&(d^e) ^ e + v5 + uint32(0x5A827999) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v6 = (*(*Tuint32_t)(unsafe.Pointer(block + 5*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 5*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 5*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 5*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 5*4)) = v6
	e += b&(c^d) ^ d + v6 + uint32(0x5A827999) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v7 = (*(*Tuint32_t)(unsafe.Pointer(block + 6*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 6*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 6*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 6*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 6*4)) = v7
	d += a&(b^c) ^ c + v7 + uint32(0x5A827999) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v8 = (*(*Tuint32_t)(unsafe.Pointer(block + 7*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 7*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 7*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 7*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 7*4)) = v8
	c += e&(a^b) ^ b + v8 + uint32(0x5A827999) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v9 = (*(*Tuint32_t)(unsafe.Pointer(block + 8*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 8*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 8*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 8*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 8*4)) = v9
	b += d&(e^a) ^ a + v9 + uint32(0x5A827999) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v10 = (*(*Tuint32_t)(unsafe.Pointer(block + 9*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 9*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 9*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 9*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 9*4)) = v10
	a += c&(d^e) ^ e + v10 + uint32(0x5A827999) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v11 = (*(*Tuint32_t)(unsafe.Pointer(block + 10*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 10*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 10*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 10*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 10*4)) = v11
	e += b&(c^d) ^ d + v11 + uint32(0x5A827999) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v12 = (*(*Tuint32_t)(unsafe.Pointer(block + 11*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 11*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 11*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 11*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 11*4)) = v12
	d += a&(b^c) ^ c + v12 + uint32(0x5A827999) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v13 = (*(*Tuint32_t)(unsafe.Pointer(block + 12*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 12*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 12*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 12*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 12*4)) = v13
	c += e&(a^b) ^ b + v13 + uint32(0x5A827999) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v14 = (*(*Tuint32_t)(unsafe.Pointer(block + 13*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 13*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 13*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 13*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 13*4)) = v14
	b += d&(e^a) ^ a + v14 + uint32(0x5A827999) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v15 = (*(*Tuint32_t)(unsafe.Pointer(block + 14*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 14*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 14*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 14*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 14*4)) = v15
	a += c&(d^e) ^ e + v15 + uint32(0x5A827999) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v16 = (*(*Tuint32_t)(unsafe.Pointer(block + 15*4))<<libc.Int32FromInt32(24)|*(*Tuint32_t)(unsafe.Pointer(block + 15*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(24)))&libc.Uint32FromUint32(0xFF00FF00) | (*(*Tuint32_t)(unsafe.Pointer(block + 15*4))<<libc.Int32FromInt32(8)|*(*Tuint32_t)(unsafe.Pointer(block + 15*4))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(8)))&libc.Uint32FromInt32(0x00FF00FF)
	*(*Tuint32_t)(unsafe.Pointer(block + 15*4)) = v16
	e += b&(c^d) ^ d + v16 + uint32(0x5A827999) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v17 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(16)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(16)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(16)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(16)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(16)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(16)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(16)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(16)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(16)&libc.Int32FromInt32(15))*4)) = v17
	d += a&(b^c) ^ c + v17 + uint32(0x5A827999) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v18 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(17)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(17)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(17)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(17)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(17)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(17)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(17)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(17)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(17)&libc.Int32FromInt32(15))*4)) = v18
	c += e&(a^b) ^ b + v18 + uint32(0x5A827999) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v19 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(18)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(18)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(18)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(18)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(18)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(18)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(18)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(18)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(18)&libc.Int32FromInt32(15))*4)) = v19
	b += d&(e^a) ^ a + v19 + uint32(0x5A827999) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v20 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(19)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(19)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(19)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(19)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(19)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(19)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(19)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(19)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(19)&libc.Int32FromInt32(15))*4)) = v20
	a += c&(d^e) ^ e + v20 + uint32(0x5A827999) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v21 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(20)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(20)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(20)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(20)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(20)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(20)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(20)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(20)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(20)&libc.Int32FromInt32(15))*4)) = v21
	e += b ^ c ^ d + v21 + uint32(0x6ED9EBA1) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v22 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(21)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(21)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(21)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(21)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(21)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(21)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(21)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(21)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(21)&libc.Int32FromInt32(15))*4)) = v22
	d += a ^ b ^ c + v22 + uint32(0x6ED9EBA1) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v23 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(22)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(22)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(22)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(22)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(22)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(22)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(22)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(22)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(22)&libc.Int32FromInt32(15))*4)) = v23
	c += e ^ a ^ b + v23 + uint32(0x6ED9EBA1) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v24 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(23)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(23)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(23)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(23)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(23)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(23)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(23)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(23)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(23)&libc.Int32FromInt32(15))*4)) = v24
	b += d ^ e ^ a + v24 + uint32(0x6ED9EBA1) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v25 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(24)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(24)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(24)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(24)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(24)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(24)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(24)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(24)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(24)&libc.Int32FromInt32(15))*4)) = v25
	a += c ^ d ^ e + v25 + uint32(0x6ED9EBA1) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v26 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(25)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(25)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(25)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(25)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(25)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(25)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(25)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(25)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(25)&libc.Int32FromInt32(15))*4)) = v26
	e += b ^ c ^ d + v26 + uint32(0x6ED9EBA1) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v27 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(26)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(26)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(26)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(26)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(26)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(26)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(26)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(26)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(26)&libc.Int32FromInt32(15))*4)) = v27
	d += a ^ b ^ c + v27 + uint32(0x6ED9EBA1) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v28 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(27)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(27)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(27)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(27)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(27)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(27)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(27)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(27)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(27)&libc.Int32FromInt32(15))*4)) = v28
	c += e ^ a ^ b + v28 + uint32(0x6ED9EBA1) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v29 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(28)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(28)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(28)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(28)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(28)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(28)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(28)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(28)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(28)&libc.Int32FromInt32(15))*4)) = v29
	b += d ^ e ^ a + v29 + uint32(0x6ED9EBA1) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v30 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(29)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(29)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(29)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(29)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(29)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(29)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(29)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(29)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(29)&libc.Int32FromInt32(15))*4)) = v30
	a += c ^ d ^ e + v30 + uint32(0x6ED9EBA1) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v31 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(30)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(30)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(30)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(30)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(30)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(30)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(30)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(30)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(30)&libc.Int32FromInt32(15))*4)) = v31
	e += b ^ c ^ d + v31 + uint32(0x6ED9EBA1) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v32 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(31)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(31)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(31)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(31)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(31)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(31)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(31)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(31)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(31)&libc.Int32FromInt32(15))*4)) = v32
	d += a ^ b ^ c + v32 + uint32(0x6ED9EBA1) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v33 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(32)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(32)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(32)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(32)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(32)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(32)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(32)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(32)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(32)&libc.Int32FromInt32(15))*4)) = v33
	c += e ^ a ^ b + v33 + uint32(0x6ED9EBA1) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v34 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(33)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(33)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(33)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(33)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(33)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(33)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(33)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(33)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(33)&libc.Int32FromInt32(15))*4)) = v34
	b += d ^ e ^ a + v34 + uint32(0x6ED9EBA1) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v35 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(34)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(34)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(34)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(34)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(34)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(34)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(34)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(34)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(34)&libc.Int32FromInt32(15))*4)) = v35
	a += c ^ d ^ e + v35 + uint32(0x6ED9EBA1) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v36 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(35)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(35)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(35)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(35)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(35)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(35)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(35)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(35)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(35)&libc.Int32FromInt32(15))*4)) = v36
	e += b ^ c ^ d + v36 + uint32(0x6ED9EBA1) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v37 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(36)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(36)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(36)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(36)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(36)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(36)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(36)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(36)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(36)&libc.Int32FromInt32(15))*4)) = v37
	d += a ^ b ^ c + v37 + uint32(0x6ED9EBA1) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v38 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(37)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(37)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(37)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(37)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(37)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(37)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(37)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(37)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(37)&libc.Int32FromInt32(15))*4)) = v38
	c += e ^ a ^ b + v38 + uint32(0x6ED9EBA1) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v39 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(38)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(38)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(38)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(38)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(38)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(38)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(38)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(38)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(38)&libc.Int32FromInt32(15))*4)) = v39
	b += d ^ e ^ a + v39 + uint32(0x6ED9EBA1) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v40 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(39)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(39)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(39)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(39)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(39)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(39)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(39)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(39)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(39)&libc.Int32FromInt32(15))*4)) = v40
	a += c ^ d ^ e + v40 + uint32(0x6ED9EBA1) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v41 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(40)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(40)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(40)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(40)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(40)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(40)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(40)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(40)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(40)&libc.Int32FromInt32(15))*4)) = v41
	e += (b|c)&d | b&c + v41 + uint32(0x8F1BBCDC) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v42 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(41)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(41)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(41)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(41)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(41)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(41)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(41)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(41)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(41)&libc.Int32FromInt32(15))*4)) = v42
	d += (a|b)&c | a&b + v42 + uint32(0x8F1BBCDC) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v43 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(42)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(42)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(42)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(42)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(42)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(42)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(42)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(42)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(42)&libc.Int32FromInt32(15))*4)) = v43
	c += (e|a)&b | e&a + v43 + uint32(0x8F1BBCDC) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v44 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(43)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(43)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(43)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(43)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(43)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(43)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(43)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(43)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(43)&libc.Int32FromInt32(15))*4)) = v44
	b += (d|e)&a | d&e + v44 + uint32(0x8F1BBCDC) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v45 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(44)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(44)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(44)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(44)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(44)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(44)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(44)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(44)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(44)&libc.Int32FromInt32(15))*4)) = v45
	a += (c|d)&e | c&d + v45 + uint32(0x8F1BBCDC) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v46 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(45)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(45)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(45)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(45)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(45)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(45)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(45)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(45)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(45)&libc.Int32FromInt32(15))*4)) = v46
	e += (b|c)&d | b&c + v46 + uint32(0x8F1BBCDC) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v47 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(46)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(46)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(46)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(46)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(46)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(46)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(46)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(46)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(46)&libc.Int32FromInt32(15))*4)) = v47
	d += (a|b)&c | a&b + v47 + uint32(0x8F1BBCDC) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v48 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(47)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(47)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(47)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(47)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(47)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(47)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(47)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(47)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(47)&libc.Int32FromInt32(15))*4)) = v48
	c += (e|a)&b | e&a + v48 + uint32(0x8F1BBCDC) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v49 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(48)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(48)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(48)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(48)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(48)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(48)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(48)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(48)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(48)&libc.Int32FromInt32(15))*4)) = v49
	b += (d|e)&a | d&e + v49 + uint32(0x8F1BBCDC) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v50 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(49)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(49)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(49)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(49)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(49)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(49)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(49)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(49)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(49)&libc.Int32FromInt32(15))*4)) = v50
	a += (c|d)&e | c&d + v50 + uint32(0x8F1BBCDC) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v51 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(50)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(50)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(50)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(50)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(50)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(50)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(50)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(50)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(50)&libc.Int32FromInt32(15))*4)) = v51
	e += (b|c)&d | b&c + v51 + uint32(0x8F1BBCDC) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v52 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(51)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(51)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(51)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(51)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(51)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(51)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(51)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(51)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(51)&libc.Int32FromInt32(15))*4)) = v52
	d += (a|b)&c | a&b + v52 + uint32(0x8F1BBCDC) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v53 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(52)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(52)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(52)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(52)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(52)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(52)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(52)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(52)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(52)&libc.Int32FromInt32(15))*4)) = v53
	c += (e|a)&b | e&a + v53 + uint32(0x8F1BBCDC) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v54 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(53)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(53)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(53)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(53)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(53)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(53)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(53)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(53)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(53)&libc.Int32FromInt32(15))*4)) = v54
	b += (d|e)&a | d&e + v54 + uint32(0x8F1BBCDC) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v55 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(54)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(54)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(54)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(54)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(54)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(54)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(54)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(54)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(54)&libc.Int32FromInt32(15))*4)) = v55
	a += (c|d)&e | c&d + v55 + uint32(0x8F1BBCDC) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v56 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(55)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(55)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(55)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(55)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(55)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(55)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(55)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(55)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(55)&libc.Int32FromInt32(15))*4)) = v56
	e += (b|c)&d | b&c + v56 + uint32(0x8F1BBCDC) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v57 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(56)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(56)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(56)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(56)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(56)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(56)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(56)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(56)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(56)&libc.Int32FromInt32(15))*4)) = v57
	d += (a|b)&c | a&b + v57 + uint32(0x8F1BBCDC) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v58 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(57)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(57)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(57)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(57)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(57)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(57)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(57)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(57)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(57)&libc.Int32FromInt32(15))*4)) = v58
	c += (e|a)&b | e&a + v58 + uint32(0x8F1BBCDC) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v59 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(58)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(58)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(58)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(58)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(58)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(58)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(58)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(58)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(58)&libc.Int32FromInt32(15))*4)) = v59
	b += (d|e)&a | d&e + v59 + uint32(0x8F1BBCDC) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v60 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(59)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(59)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(59)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(59)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(59)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(59)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(59)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(59)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(59)&libc.Int32FromInt32(15))*4)) = v60
	a += (c|d)&e | c&d + v60 + uint32(0x8F1BBCDC) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v61 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(60)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(60)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(60)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(60)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(60)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(60)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(60)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(60)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(60)&libc.Int32FromInt32(15))*4)) = v61
	e += b ^ c ^ d + v61 + uint32(0xCA62C1D6) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v62 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(61)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(61)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(61)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(61)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(61)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(61)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(61)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(61)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(61)&libc.Int32FromInt32(15))*4)) = v62
	d += a ^ b ^ c + v62 + uint32(0xCA62C1D6) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v63 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(62)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(62)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(62)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(62)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(62)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(62)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(62)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(62)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(62)&libc.Int32FromInt32(15))*4)) = v63
	c += e ^ a ^ b + v63 + uint32(0xCA62C1D6) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v64 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(63)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(63)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(63)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(63)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(63)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(63)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(63)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(63)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(63)&libc.Int32FromInt32(15))*4)) = v64
	b += d ^ e ^ a + v64 + uint32(0xCA62C1D6) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v65 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(64)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(64)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(64)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(64)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(64)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(64)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(64)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(64)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(64)&libc.Int32FromInt32(15))*4)) = v65
	a += c ^ d ^ e + v65 + uint32(0xCA62C1D6) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v66 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(65)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(65)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(65)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(65)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(65)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(65)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(65)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(65)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(65)&libc.Int32FromInt32(15))*4)) = v66
	e += b ^ c ^ d + v66 + uint32(0xCA62C1D6) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v67 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(66)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(66)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(66)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(66)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(66)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(66)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(66)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(66)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(66)&libc.Int32FromInt32(15))*4)) = v67
	d += a ^ b ^ c + v67 + uint32(0xCA62C1D6) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v68 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(67)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(67)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(67)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(67)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(67)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(67)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(67)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(67)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(67)&libc.Int32FromInt32(15))*4)) = v68
	c += e ^ a ^ b + v68 + uint32(0xCA62C1D6) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v69 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(68)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(68)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(68)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(68)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(68)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(68)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(68)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(68)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(68)&libc.Int32FromInt32(15))*4)) = v69
	b += d ^ e ^ a + v69 + uint32(0xCA62C1D6) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v70 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(69)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(69)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(69)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(69)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(69)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(69)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(69)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(69)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(69)&libc.Int32FromInt32(15))*4)) = v70
	a += c ^ d ^ e + v70 + uint32(0xCA62C1D6) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v71 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(70)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(70)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(70)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(70)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(70)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(70)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(70)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(70)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(70)&libc.Int32FromInt32(15))*4)) = v71
	e += b ^ c ^ d + v71 + uint32(0xCA62C1D6) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v72 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(71)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(71)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(71)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(71)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(71)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(71)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(71)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(71)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(71)&libc.Int32FromInt32(15))*4)) = v72
	d += a ^ b ^ c + v72 + uint32(0xCA62C1D6) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v73 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(72)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(72)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(72)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(72)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(72)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(72)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(72)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(72)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(72)&libc.Int32FromInt32(15))*4)) = v73
	c += e ^ a ^ b + v73 + uint32(0xCA62C1D6) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v74 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(73)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(73)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(73)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(73)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(73)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(73)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(73)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(73)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(73)&libc.Int32FromInt32(15))*4)) = v74
	b += d ^ e ^ a + v74 + uint32(0xCA62C1D6) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v75 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(74)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(74)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(74)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(74)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(74)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(74)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(74)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(74)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(74)&libc.Int32FromInt32(15))*4)) = v75
	a += c ^ d ^ e + v75 + uint32(0xCA62C1D6) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v76 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(75)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(75)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(75)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(75)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(75)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(75)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(75)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(75)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(75)&libc.Int32FromInt32(15))*4)) = v76
	e += b ^ c ^ d + v76 + uint32(0xCA62C1D6) + (a<<libc.Int32FromInt32(5) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	b = b<<libc.Int32FromInt32(30) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v77 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(76)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(76)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(76)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(76)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(76)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(76)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(76)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(76)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(76)&libc.Int32FromInt32(15))*4)) = v77
	d += a ^ b ^ c + v77 + uint32(0xCA62C1D6) + (e<<libc.Int32FromInt32(5) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	a = a<<libc.Int32FromInt32(30) | a>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v78 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(77)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(77)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(77)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(77)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(77)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(77)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(77)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(77)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(77)&libc.Int32FromInt32(15))*4)) = v78
	c += e ^ a ^ b + v78 + uint32(0xCA62C1D6) + (d<<libc.Int32FromInt32(5) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	e = e<<libc.Int32FromInt32(30) | e>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v79 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(78)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(78)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(78)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(78)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(78)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(78)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(78)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(78)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(78)&libc.Int32FromInt32(15))*4)) = v79
	b += d ^ e ^ a + v79 + uint32(0xCA62C1D6) + (c<<libc.Int32FromInt32(5) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	d = d<<libc.Int32FromInt32(30) | d>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	v80 = (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(79)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(79)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(79)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(79)&libc.Int32FromInt32(15))*4)))<<libc.Int32FromInt32(1) | (*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(79)+libc.Int32FromInt32(13))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(79)+libc.Int32FromInt32(8))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr((libc.Int32FromInt32(79)+libc.Int32FromInt32(2))&libc.Int32FromInt32(15))*4))^*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(79)&libc.Int32FromInt32(15))*4)))>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(1))
	*(*Tuint32_t)(unsafe.Pointer(block + uintptr(libc.Int32FromInt32(79)&libc.Int32FromInt32(15))*4)) = v80
	a += c ^ d ^ e + v80 + uint32(0xCA62C1D6) + (b<<libc.Int32FromInt32(5) | b>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(5)))
	c = c<<libc.Int32FromInt32(30) | c>>(libc.Int32FromInt32(32)-libc.Int32FromInt32(30))
	/* Add the working vars back into context.state[] */
	*(*Tuint32_t)(unsafe.Pointer(state)) += a
	*(*Tuint32_t)(unsafe.Pointer(state + 1*4)) += b
	*(*Tuint32_t)(unsafe.Pointer(state + 2*4)) += c
	*(*Tuint32_t)(unsafe.Pointer(state + 3*4)) += d
	*(*Tuint32_t)(unsafe.Pointer(state + 4*4)) += e
	/* Wipe variables */
	v84 = libc.Uint32FromInt32(0)
	e = v84
	v83 = v84
	d = v83
	v82 = v83
	c = v82
	v81 = v82
	b = v81
	a = v81
}

// C documentation
//
//	/*
//	 * SHA1Init - Initialize new context
//	 */
func XSHA1Init(tls *libc.TLS, context uintptr) {
	/* SHA1 initialization constants */
	(*TSHA1_CTX)(unsafe.Pointer(context)).Fcount = uint64(0)
	*(*Tuint32_t)(unsafe.Pointer(context)) = uint32(0x67452301)
	*(*Tuint32_t)(unsafe.Pointer(context + 1*4)) = uint32(0xEFCDAB89)
	*(*Tuint32_t)(unsafe.Pointer(context + 2*4)) = uint32(0x98BADCFE)
	*(*Tuint32_t)(unsafe.Pointer(context + 3*4)) = uint32(0x10325476)
	*(*Tuint32_t)(unsafe.Pointer(context + 4*4)) = uint32(0xC3D2E1F0)
}

// C documentation
//
//	/*
//	 * Run your data through this.
//	 */
func XSHA1Update(tls *libc.TLS, context uintptr, data uintptr, len1 Tsize_t) {
	var i, j, v1 Tsize_t
	_, _, _ = i, j, v1
	j = uint32((*TSHA1_CTX)(unsafe.Pointer(context)).Fcount >> libc.Int32FromInt32(3) & libc.Uint64FromInt32(63))
	*(*Tuint64_t)(unsafe.Pointer(context + 20)) += uint64(len1 << libc.Int32FromInt32(3))
	if j+len1 > uint32(63) {
		v1 = libc.Uint32FromInt32(64) - j
		i = v1
		libc.Xmemcpy(tls, context+28+uintptr(j), data, v1)
		XSHA1Transform(tls, context, context+28)
		for {
			if !(i+uint32(63) < len1) {
				break
			}
			XSHA1Transform(tls, context, data+uintptr(i))
			goto _2
		_2:
			;
			i += uint32(64)
		}
		j = uint32(0)
	} else {
		i = uint32(0)
	}
	libc.Xmemcpy(tls, context+28+uintptr(j), data+uintptr(i), len1-i)
}

// C documentation
//
//	/*
//	 * Add padding and return the message digest.
//	 */
func XSHA1Pad(tls *libc.TLS, context uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i uint32
	var _ /* finalcount at bp+0 */ [8]Tuint8_t
	_ = i
	i = uint32(0)
	for {
		if !(i < uint32(8)) {
			break
		}
		(*(*[8]Tuint8_t)(unsafe.Pointer(bp)))[i] = uint8((*TSHA1_CTX)(unsafe.Pointer(context)).Fcount >> ((libc.Uint32FromInt32(7) - i&libc.Uint32FromInt32(7)) * libc.Uint32FromInt32(8)) & libc.Uint64FromInt32(255)) /* Endian independent */
		goto _1
	_1:
		;
		i++
	}
	XSHA1Update(tls, context, __ccgo_ts+153, uint32(1))
	for (*TSHA1_CTX)(unsafe.Pointer(context)).Fcount&uint64(504) != uint64(448) {
		XSHA1Update(tls, context, __ccgo_ts+155, uint32(1))
	}
	XSHA1Update(tls, context, bp, uint32(8)) /* Should cause a SHA1Transform() */
}

func XSHA1Final(tls *libc.TLS, digest uintptr, context uintptr) {
	var i uint32
	_ = i
	XSHA1Pad(tls, context)
	if digest != 0 {
		i = uint32(0)
		for {
			if !(i < uint32(m_SHA1_DIGEST_LENGTH)) {
				break
			}
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i))) = uint8(*(*Tuint32_t)(unsafe.Pointer(context + uintptr(i>>int32(2))*4)) >> ((libc.Uint32FromInt32(3) - i&libc.Uint32FromInt32(3)) * libc.Uint32FromInt32(8)) & libc.Uint32FromInt32(255))
			goto _1
		_1:
			;
			i++
		}
		libc.Xmemset(tls, context, 0, uint32(92))
	}
}

const m_SHA256_BLOCK_LENGTH = 64
const m_SHA256_DIGEST_LENGTH = 32
const m_SHA384_BLOCK_LENGTH = 128
const m_SHA384_DIGEST_LENGTH = 48
const m_SHA512_BLOCK_LENGTH = 128
const m_SHA512_DIGEST_LENGTH = 64

type TSHA2_CTX = struct {
	Fstate struct {
		Fst64        [0][8]Tuint64_t
		Fst32        [8]Tuint32_t
		F__ccgo_pad2 [32]byte
	}
	Fbitcount [2]Tuint64_t
	Fbuffer   [128]Tuint8_t
}

type T_SHA2_CTX = TSHA2_CTX

/*
 * UNROLLED TRANSFORM LOOP NOTE:
 * You can define SHA2_UNROLL_TRANSFORM to use the unrolled transform
 * loop version for the hash transform rounds (defined using macros
 * later in this file).  Either define on the command line, for example:
 *
 *   cc -DSHA2_UNROLL_TRANSFORM -o sha2 sha2.c sha2prog.c
 *
 * or define below:
 *
 *   #define SHA2_UNROLL_TRANSFORM
 *
 */

/*** SHA-256/384/512 Various Length Definitions ***********************/
/* NOTE: Most of these are in sha2.h */

/*** ENDIAN SPECIFIC COPY MACROS **************************************/

/*
 * Macro for incrementally adding the unsigned 64-bit integer n to the
 * unsigned 128-bit integer (represented using a two-element array of
 * 64-bit words):
 */

/*** THE SIX LOGICAL FUNCTIONS ****************************************/
/*
 * Bit shifting and rotation (used by the six SHA-XYZ logical functions:
 *
 *   NOTE:  The naming of R and S appears backwards here (R is a SHIFT and
 *   S is a ROTATION) because the SHA-256/384/512 description document
 *   (see http://csrc.nist.gov/cryptval/shs/sha256-384-512.pdf) uses this
 *   same "backwards" definition.
 */
/* Shift-right (used in SHA-256, SHA-384, and SHA-512): */
/* 32-bit Rotate-right (used in SHA-256): */
/* 64-bit Rotate-right (used in SHA-384 and SHA-512): */

/* Two of six logical functions used in SHA-256, SHA-384, and SHA-512: */

/* Four of six logical functions used in SHA-256: */

/* Four of six logical functions used in SHA-384 and SHA-512: */

// C documentation
//
//	/*** SHA-XYZ INITIAL HASH VALUES AND CONSTANTS ************************/
//	/* Hash constant words K for SHA-256: */
var _K256 = [64]Tuint32_t{
	0:  uint32(0x428a2f98),
	1:  uint32(0x71374491),
	2:  uint32(0xb5c0fbcf),
	3:  uint32(0xe9b5dba5),
	4:  uint32(0x3956c25b),
	5:  uint32(0x59f111f1),
	6:  uint32(0x923f82a4),
	7:  uint32(0xab1c5ed5),
	8:  uint32(0xd807aa98),
	9:  uint32(0x12835b01),
	10: uint32(0x243185be),
	11: uint32(0x550c7dc3),
	12: uint32(0x72be5d74),
	13: uint32(0x80deb1fe),
	14: uint32(0x9bdc06a7),
	15: uint32(0xc19bf174),
	16: uint32(0xe49b69c1),
	17: uint32(0xefbe4786),
	18: uint32(0x0fc19dc6),
	19: uint32(0x240ca1cc),
	20: uint32(0x2de92c6f),
	21: uint32(0x4a7484aa),
	22: uint32(0x5cb0a9dc),
	23: uint32(0x76f988da),
	24: uint32(0x983e5152),
	25: uint32(0xa831c66d),
	26: uint32(0xb00327c8),
	27: uint32(0xbf597fc7),
	28: uint32(0xc6e00bf3),
	29: uint32(0xd5a79147),
	30: uint32(0x06ca6351),
	31: uint32(0x14292967),
	32: uint32(0x27b70a85),
	33: uint32(0x2e1b2138),
	34: uint32(0x4d2c6dfc),
	35: uint32(0x53380d13),
	36: uint32(0x650a7354),
	37: uint32(0x766a0abb),
	38: uint32(0x81c2c92e),
	39: uint32(0x92722c85),
	40: uint32(0xa2bfe8a1),
	41: uint32(0xa81a664b),
	42: uint32(0xc24b8b70),
	43: uint32(0xc76c51a3),
	44: uint32(0xd192e819),
	45: uint32(0xd6990624),
	46: uint32(0xf40e3585),
	47: uint32(0x106aa070),
	48: uint32(0x19a4c116),
	49: uint32(0x1e376c08),
	50: uint32(0x2748774c),
	51: uint32(0x34b0bcb5),
	52: uint32(0x391c0cb3),
	53: uint32(0x4ed8aa4a),
	54: uint32(0x5b9cca4f),
	55: uint32(0x682e6ff3),
	56: uint32(0x748f82ee),
	57: uint32(0x78a5636f),
	58: uint32(0x84c87814),
	59: uint32(0x8cc70208),
	60: uint32(0x90befffa),
	61: uint32(0xa4506ceb),
	62: uint32(0xbef9a3f7),
	63: uint32(0xc67178f2),
}

// C documentation
//
//	/* Initial hash value H for SHA-256: */
var _sha256_initial_hash_value = [8]Tuint32_t{
	0: uint32(0x6a09e667),
	1: uint32(0xbb67ae85),
	2: uint32(0x3c6ef372),
	3: uint32(0xa54ff53a),
	4: uint32(0x510e527f),
	5: uint32(0x9b05688c),
	6: uint32(0x1f83d9ab),
	7: uint32(0x5be0cd19),
}

// C documentation
//
//	/* Hash constant words K for SHA-384 and SHA-512: */
var _K512 = [80]Tuint64_t{
	0:  uint64(0x428a2f98d728ae22),
	1:  uint64(0x7137449123ef65cd),
	2:  uint64(0xb5c0fbcfec4d3b2f),
	3:  uint64(0xe9b5dba58189dbbc),
	4:  uint64(0x3956c25bf348b538),
	5:  uint64(0x59f111f1b605d019),
	6:  uint64(0x923f82a4af194f9b),
	7:  uint64(0xab1c5ed5da6d8118),
	8:  uint64(0xd807aa98a3030242),
	9:  uint64(0x12835b0145706fbe),
	10: uint64(0x243185be4ee4b28c),
	11: uint64(0x550c7dc3d5ffb4e2),
	12: uint64(0x72be5d74f27b896f),
	13: uint64(0x80deb1fe3b1696b1),
	14: uint64(0x9bdc06a725c71235),
	15: uint64(0xc19bf174cf692694),
	16: uint64(0xe49b69c19ef14ad2),
	17: uint64(0xefbe4786384f25e3),
	18: uint64(0x0fc19dc68b8cd5b5),
	19: uint64(0x240ca1cc77ac9c65),
	20: uint64(0x2de92c6f592b0275),
	21: uint64(0x4a7484aa6ea6e483),
	22: uint64(0x5cb0a9dcbd41fbd4),
	23: uint64(0x76f988da831153b5),
	24: uint64(0x983e5152ee66dfab),
	25: uint64(0xa831c66d2db43210),
	26: uint64(0xb00327c898fb213f),
	27: uint64(0xbf597fc7beef0ee4),
	28: uint64(0xc6e00bf33da88fc2),
	29: uint64(0xd5a79147930aa725),
	30: uint64(0x06ca6351e003826f),
	31: uint64(0x142929670a0e6e70),
	32: uint64(0x27b70a8546d22ffc),
	33: uint64(0x2e1b21385c26c926),
	34: uint64(0x4d2c6dfc5ac42aed),
	35: uint64(0x53380d139d95b3df),
	36: uint64(0x650a73548baf63de),
	37: uint64(0x766a0abb3c77b2a8),
	38: uint64(0x81c2c92e47edaee6),
	39: uint64(0x92722c851482353b),
	40: uint64(0xa2bfe8a14cf10364),
	41: uint64(0xa81a664bbc423001),
	42: uint64(0xc24b8b70d0f89791),
	43: uint64(0xc76c51a30654be30),
	44: uint64(0xd192e819d6ef5218),
	45: uint64(0xd69906245565a910),
	46: uint64(0xf40e35855771202a),
	47: uint64(0x106aa07032bbd1b8),
	48: uint64(0x19a4c116b8d2d0c8),
	49: uint64(0x1e376c085141ab53),
	50: uint64(0x2748774cdf8eeb99),
	51: uint64(0x34b0bcb5e19b48a8),
	52: uint64(0x391c0cb3c5c95a63),
	53: uint64(0x4ed8aa4ae3418acb),
	54: uint64(0x5b9cca4f7763e373),
	55: uint64(0x682e6ff3d6b2b8a3),
	56: uint64(0x748f82ee5defb2fc),
	57: uint64(0x78a5636f43172f60),
	58: uint64(0x84c87814a1f0ab72),
	59: uint64(0x8cc702081a6439ec),
	60: uint64(0x90befffa23631e28),
	61: uint64(0xa4506cebde82bde9),
	62: uint64(0xbef9a3f7b2c67915),
	63: uint64(0xc67178f2e372532b),
	64: uint64(0xca273eceea26619c),
	65: uint64(0xd186b8c721c0c207),
	66: uint64(0xeada7dd6cde0eb1e),
	67: uint64(0xf57d4f7fee6ed178),
	68: uint64(0x06f067aa72176fba),
	69: uint64(0x0a637dc5a2c898a6),
	70: uint64(0x113f9804bef90dae),
	71: uint64(0x1b710b35131c471b),
	72: uint64(0x28db77f523047d84),
	73: uint64(0x32caab7b40c72493),
	74: uint64(0x3c9ebe0a15c9bebc),
	75: uint64(0x431d67c49c100d4c),
	76: uint64(0x4cc5d4becb3e42b6),
	77: uint64(0x597f299cfc657e2a),
	78: uint64(0x5fcb6fab3ad6faec),
	79: uint64(0x6c44198c4a475817),
}

// C documentation
//
//	/* Initial hash value H for SHA-384 */
var _sha384_initial_hash_value = [8]Tuint64_t{
	0: uint64(0xcbbb9d5dc1059ed8),
	1: uint64(0x629a292a367cd507),
	2: uint64(0x9159015a3070dd17),
	3: uint64(0x152fecd8f70e5939),
	4: uint64(0x67332667ffc00b31),
	5: uint64(0x8eb44a8768581511),
	6: uint64(0xdb0c2e0d64f98fa7),
	7: uint64(0x47b5481dbefa4fa4),
}

// C documentation
//
//	/* Initial hash value H for SHA-512 */
var _sha512_initial_hash_value = [8]Tuint64_t{
	0: uint64(0x6a09e667f3bcc908),
	1: uint64(0xbb67ae8584caa73b),
	2: uint64(0x3c6ef372fe94f82b),
	3: uint64(0xa54ff53a5f1d36f1),
	4: uint64(0x510e527fade682d1),
	5: uint64(0x9b05688c2b3e6c1f),
	6: uint64(0x1f83d9abfb41bd6b),
	7: uint64(0x5be0cd19137e2179),
}

// C documentation
//
//	/*** SHA-256: *********************************************************/
func XSHA256Init(tls *libc.TLS, context uintptr) {
	if context == libc.UintptrFromInt32(0) {
		return
	}
	libc.Xmemcpy(tls, context, uintptr(unsafe.Pointer(&_sha256_initial_hash_value)), uint32(32))
	libc.Xmemset(tls, context+80, 0, uint32(128))
	*(*Tuint64_t)(unsafe.Pointer(context + 64)) = uint64(0)
}

func XSHA256Transform(tls *libc.TLS, state uintptr, data uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var T1, T2, a, b, c, d, e, f, g, h, s0, s1, v10, v2, v3, v4, v5, v6, v7, v8, v9 Tuint32_t
	var j int32
	var p1 uintptr
	var _ /* W256 at bp+0 */ [16]Tuint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = T1, T2, a, b, c, d, e, f, g, h, j, s0, s1, v10, v2, v3, v4, v5, v6, v7, v8, v9, p1
	/* Initialize registers with the prev. intermediate value */
	a = *(*Tuint32_t)(unsafe.Pointer(state))
	b = *(*Tuint32_t)(unsafe.Pointer(state + 1*4))
	c = *(*Tuint32_t)(unsafe.Pointer(state + 2*4))
	d = *(*Tuint32_t)(unsafe.Pointer(state + 3*4))
	e = *(*Tuint32_t)(unsafe.Pointer(state + 4*4))
	f = *(*Tuint32_t)(unsafe.Pointer(state + 5*4))
	g = *(*Tuint32_t)(unsafe.Pointer(state + 6*4))
	h = *(*Tuint32_t)(unsafe.Pointer(state + 7*4))
	j = 0
	for cond := true; cond; cond = j < int32(16) {
		(*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[j] = uint32(*(*Tuint8_t)(unsafe.Pointer(data + 3))) | uint32(*(*Tuint8_t)(unsafe.Pointer(data + 2)))<<libc.Int32FromInt32(8) | uint32(*(*Tuint8_t)(unsafe.Pointer(data + 1)))<<libc.Int32FromInt32(16) | uint32(*(*Tuint8_t)(unsafe.Pointer(data)))<<libc.Int32FromInt32(24)
		data += uintptr(4)
		/* Apply the SHA-256 compression function to update a..h */
		T1 = h + (e>>libc.Int32FromInt32(6) | e<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) ^ (e>>libc.Int32FromInt32(11) | e<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))) ^ (e>>libc.Int32FromInt32(25) | e<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(25)))) + (e&f ^ ^e&g) + _K256[j] + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[j]
		T2 = a>>libc.Int32FromInt32(2) | a<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(2)) ^ (a>>libc.Int32FromInt32(13) | a<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(13))) ^ (a>>libc.Int32FromInt32(22) | a<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(22))) + (a&b ^ a&c ^ b&c)
		h = g
		g = f
		f = e
		e = d + T1
		d = c
		c = b
		b = a
		a = T1 + T2
		j++
	}
	for cond := true; cond; cond = j < int32(64) {
		/* Part of the message block expansion: */
		s0 = (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[(j+int32(1))&int32(0x0f)]
		s0 = s0>>libc.Int32FromInt32(7) | s0<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(7)) ^ (s0>>libc.Int32FromInt32(18) | s0<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(18))) ^ s0>>libc.Int32FromInt32(3)
		s1 = (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[(j+int32(14))&int32(0x0f)]
		s1 = s1>>libc.Int32FromInt32(17) | s1<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(17)) ^ (s1>>libc.Int32FromInt32(19) | s1<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(19))) ^ s1>>libc.Int32FromInt32(10)
		/* Apply the SHA-256 compression function to update a..h */
		p1 = bp + uintptr(j&int32(0x0f))*4
		*(*Tuint32_t)(unsafe.Pointer(p1)) += s1 + (*(*[16]Tuint32_t)(unsafe.Pointer(bp)))[(j+int32(9))&int32(0x0f)] + s0
		T1 = h + (e>>libc.Int32FromInt32(6) | e<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(6)) ^ (e>>libc.Int32FromInt32(11) | e<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(11))) ^ (e>>libc.Int32FromInt32(25) | e<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(25)))) + (e&f ^ ^e&g) + _K256[j] + *(*Tuint32_t)(unsafe.Pointer(p1))
		T2 = a>>libc.Int32FromInt32(2) | a<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(2)) ^ (a>>libc.Int32FromInt32(13) | a<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(13))) ^ (a>>libc.Int32FromInt32(22) | a<<(libc.Int32FromInt32(32)-libc.Int32FromInt32(22))) + (a&b ^ a&c ^ b&c)
		h = g
		g = f
		f = e
		e = d + T1
		d = c
		c = b
		b = a
		a = T1 + T2
		j++
	}
	/* Compute the current intermediate hash value */
	*(*Tuint32_t)(unsafe.Pointer(state)) += a
	*(*Tuint32_t)(unsafe.Pointer(state + 1*4)) += b
	*(*Tuint32_t)(unsafe.Pointer(state + 2*4)) += c
	*(*Tuint32_t)(unsafe.Pointer(state + 3*4)) += d
	*(*Tuint32_t)(unsafe.Pointer(state + 4*4)) += e
	*(*Tuint32_t)(unsafe.Pointer(state + 5*4)) += f
	*(*Tuint32_t)(unsafe.Pointer(state + 6*4)) += g
	*(*Tuint32_t)(unsafe.Pointer(state + 7*4)) += h
	/* Clean up */
	v10 = libc.Uint32FromInt32(0)
	T2 = v10
	v9 = v10
	T1 = v9
	v8 = v9
	h = v8
	v7 = v8
	g = v7
	v6 = v7
	f = v6
	v5 = v6
	e = v5
	v4 = v5
	d = v4
	v3 = v4
	c = v3
	v2 = v3
	b = v2
	a = v2
}

func XSHA256Update(tls *libc.TLS, context uintptr, data uintptr, len1 Tsize_t) {
	var freespace, usedspace, v1, v2 Tsize_t
	_, _, _, _ = freespace, usedspace, v1, v2
	/* Calling with no data is valid (we do nothing) */
	if len1 == uint32(0) {
		return
	}
	usedspace = uint32(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> libc.Int32FromInt32(3) % uint64(m_SHA256_BLOCK_LENGTH))
	if usedspace > uint32(0) {
		/* Calculate how much free space is available in the buffer */
		freespace = uint32(m_SHA256_BLOCK_LENGTH) - usedspace
		if len1 >= freespace {
			/* Fill the buffer completely and process it */
			libc.Xmemcpy(tls, context+80+uintptr(usedspace), data, freespace)
			*(*Tuint64_t)(unsafe.Pointer(context + 64)) += uint64(freespace << int32(3))
			len1 -= freespace
			data += uintptr(freespace)
			XSHA256Transform(tls, context, context+80)
		} else {
			/* The buffer is not yet full */
			libc.Xmemcpy(tls, context+80+uintptr(usedspace), data, len1)
			*(*Tuint64_t)(unsafe.Pointer(context + 64)) += uint64(len1 << int32(3))
			/* Clean up: */
			v1 = libc.Uint32FromInt32(0)
			freespace = v1
			usedspace = v1
			return
		}
	}
	for len1 >= uint32(m_SHA256_BLOCK_LENGTH) {
		/* Process as many complete blocks as we can */
		XSHA256Transform(tls, context, data)
		*(*Tuint64_t)(unsafe.Pointer(context + 64)) += libc.Uint64FromInt32(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH) << libc.Int32FromInt32(3))
		len1 -= uint32(m_SHA256_BLOCK_LENGTH)
		data += uintptr(m_SHA256_BLOCK_LENGTH)
	}
	if len1 > uint32(0) {
		/* There's left-overs, so save 'em */
		libc.Xmemcpy(tls, context+80, data, len1)
		*(*Tuint64_t)(unsafe.Pointer(context + 64)) += uint64(len1 << int32(3))
	}
	/* Clean up: */
	v2 = libc.Uint32FromInt32(0)
	freespace = v2
	usedspace = v2
}

func XSHA256Pad(tls *libc.TLS, context uintptr) {
	var usedspace, v1 uint32
	_, _ = usedspace, v1
	usedspace = uint32(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> libc.Int32FromInt32(3) % uint64(m_SHA256_BLOCK_LENGTH))
	if usedspace > uint32(0) {
		/* Begin padding with a 1 bit: */
		v1 = usedspace
		usedspace++
		*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(v1))) = uint8(0x80)
		if usedspace <= libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) {
			/* Set-up for the last transform: */
			libc.Xmemset(tls, context+80+uintptr(usedspace), 0, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8))-usedspace)
		} else {
			if usedspace < uint32(m_SHA256_BLOCK_LENGTH) {
				libc.Xmemset(tls, context+80+uintptr(usedspace), 0, uint32(m_SHA256_BLOCK_LENGTH)-usedspace)
			}
			/* Do second-to-last transform: */
			XSHA256Transform(tls, context, context+80)
			/* Prepare for last transform: */
			libc.Xmemset(tls, context+80, 0, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)))
		}
	} else {
		/* Set-up for the last transform: */
		libc.Xmemset(tls, context+80, 0, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)))
		/* Begin padding with a 1 bit: */
		*(*Tuint8_t)(unsafe.Pointer(context + 80)) = uint8(0x80)
	}
	/* Store the length of input data (in bits) in big endian format: */
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)))) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(56))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 1)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(48))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 2)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(40))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 3)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(32))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 4)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(24))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 5)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(16))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 6)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(8))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA256_BLOCK_LENGTH)-libc.Int32FromInt32(8)) + 7)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)))
	/* Final transform: */
	XSHA256Transform(tls, context, context+80)
	/* Clean up: */
	usedspace = uint32(0)
}

func XSHA256Final(tls *libc.TLS, digest uintptr, context uintptr) {
	var i int32
	_ = i
	XSHA256Pad(tls, context)
	/* If no digest buffer is passed, we don't bother doing this: */
	if digest != libc.UintptrFromInt32(0) {
		/* Convert TO host byte order */
		i = 0
		for {
			if !(i < int32(8)) {
				break
			}
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)))) = uint8(*(*Tuint32_t)(unsafe.Pointer(context + uintptr(i)*4)) >> int32(24))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 1)) = uint8(*(*Tuint32_t)(unsafe.Pointer(context + uintptr(i)*4)) >> int32(16))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 2)) = uint8(*(*Tuint32_t)(unsafe.Pointer(context + uintptr(i)*4)) >> int32(8))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(4)) + 3)) = uint8(*(*Tuint32_t)(unsafe.Pointer(context + uintptr(i)*4)))
			goto _1
		_1:
			;
			i++
		}
		libc.Xmemset(tls, context, 0, uint32(208))
	}
}

// C documentation
//
//	/*** SHA-512: *********************************************************/
func XSHA512Init(tls *libc.TLS, context uintptr) {
	var v1 Tuint64_t
	_ = v1
	if context == libc.UintptrFromInt32(0) {
		return
	}
	libc.Xmemcpy(tls, context, uintptr(unsafe.Pointer(&_sha512_initial_hash_value)), uint32(64))
	libc.Xmemset(tls, context+80, 0, uint32(128))
	v1 = libc.Uint64FromInt32(0)
	*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) = v1
	*(*Tuint64_t)(unsafe.Pointer(context + 64)) = v1
}

func XSHA512Transform(tls *libc.TLS, state uintptr, data uintptr) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var T1, T2, a, b, c, d, e, f, g, h, s0, s1, v10, v2, v3, v4, v5, v6, v7, v8, v9 Tuint64_t
	var j int32
	var p1 uintptr
	var _ /* W512 at bp+0 */ [16]Tuint64_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = T1, T2, a, b, c, d, e, f, g, h, j, s0, s1, v10, v2, v3, v4, v5, v6, v7, v8, v9, p1
	/* Initialize registers with the prev. intermediate value */
	a = *(*Tuint64_t)(unsafe.Pointer(state))
	b = *(*Tuint64_t)(unsafe.Pointer(state + 1*8))
	c = *(*Tuint64_t)(unsafe.Pointer(state + 2*8))
	d = *(*Tuint64_t)(unsafe.Pointer(state + 3*8))
	e = *(*Tuint64_t)(unsafe.Pointer(state + 4*8))
	f = *(*Tuint64_t)(unsafe.Pointer(state + 5*8))
	g = *(*Tuint64_t)(unsafe.Pointer(state + 6*8))
	h = *(*Tuint64_t)(unsafe.Pointer(state + 7*8))
	j = 0
	for cond := true; cond; cond = j < int32(16) {
		(*(*[16]Tuint64_t)(unsafe.Pointer(bp)))[j] = uint64(*(*Tuint8_t)(unsafe.Pointer(data + 7))) | uint64(*(*Tuint8_t)(unsafe.Pointer(data + 6)))<<libc.Int32FromInt32(8) | uint64(*(*Tuint8_t)(unsafe.Pointer(data + 5)))<<libc.Int32FromInt32(16) | uint64(*(*Tuint8_t)(unsafe.Pointer(data + 4)))<<libc.Int32FromInt32(24) | uint64(*(*Tuint8_t)(unsafe.Pointer(data + 3)))<<libc.Int32FromInt32(32) | uint64(*(*Tuint8_t)(unsafe.Pointer(data + 2)))<<libc.Int32FromInt32(40) | uint64(*(*Tuint8_t)(unsafe.Pointer(data + 1)))<<libc.Int32FromInt32(48) | uint64(*(*Tuint8_t)(unsafe.Pointer(data)))<<libc.Int32FromInt32(56)
		data += uintptr(8)
		/* Apply the SHA-512 compression function to update a..h */
		T1 = h + (e>>libc.Int32FromInt32(14) | e<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(14)) ^ (e>>libc.Int32FromInt32(18) | e<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(18))) ^ (e>>libc.Int32FromInt32(41) | e<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(41)))) + (e&f ^ ^e&g) + _K512[j] + (*(*[16]Tuint64_t)(unsafe.Pointer(bp)))[j]
		T2 = a>>libc.Int32FromInt32(28) | a<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(28)) ^ (a>>libc.Int32FromInt32(34) | a<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(34))) ^ (a>>libc.Int32FromInt32(39) | a<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(39))) + (a&b ^ a&c ^ b&c)
		h = g
		g = f
		f = e
		e = d + T1
		d = c
		c = b
		b = a
		a = T1 + T2
		j++
	}
	for cond := true; cond; cond = j < int32(80) {
		/* Part of the message block expansion: */
		s0 = (*(*[16]Tuint64_t)(unsafe.Pointer(bp)))[(j+int32(1))&int32(0x0f)]
		s0 = s0>>libc.Int32FromInt32(1) | s0<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(1)) ^ (s0>>libc.Int32FromInt32(8) | s0<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(8))) ^ s0>>libc.Int32FromInt32(7)
		s1 = (*(*[16]Tuint64_t)(unsafe.Pointer(bp)))[(j+int32(14))&int32(0x0f)]
		s1 = s1>>libc.Int32FromInt32(19) | s1<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(19)) ^ (s1>>libc.Int32FromInt32(61) | s1<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(61))) ^ s1>>libc.Int32FromInt32(6)
		/* Apply the SHA-512 compression function to update a..h */
		p1 = bp + uintptr(j&int32(0x0f))*8
		*(*Tuint64_t)(unsafe.Pointer(p1)) += s1 + (*(*[16]Tuint64_t)(unsafe.Pointer(bp)))[(j+int32(9))&int32(0x0f)] + s0
		T1 = h + (e>>libc.Int32FromInt32(14) | e<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(14)) ^ (e>>libc.Int32FromInt32(18) | e<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(18))) ^ (e>>libc.Int32FromInt32(41) | e<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(41)))) + (e&f ^ ^e&g) + _K512[j] + *(*Tuint64_t)(unsafe.Pointer(p1))
		T2 = a>>libc.Int32FromInt32(28) | a<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(28)) ^ (a>>libc.Int32FromInt32(34) | a<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(34))) ^ (a>>libc.Int32FromInt32(39) | a<<(libc.Int32FromInt32(64)-libc.Int32FromInt32(39))) + (a&b ^ a&c ^ b&c)
		h = g
		g = f
		f = e
		e = d + T1
		d = c
		c = b
		b = a
		a = T1 + T2
		j++
	}
	/* Compute the current intermediate hash value */
	*(*Tuint64_t)(unsafe.Pointer(state)) += a
	*(*Tuint64_t)(unsafe.Pointer(state + 1*8)) += b
	*(*Tuint64_t)(unsafe.Pointer(state + 2*8)) += c
	*(*Tuint64_t)(unsafe.Pointer(state + 3*8)) += d
	*(*Tuint64_t)(unsafe.Pointer(state + 4*8)) += e
	*(*Tuint64_t)(unsafe.Pointer(state + 5*8)) += f
	*(*Tuint64_t)(unsafe.Pointer(state + 6*8)) += g
	*(*Tuint64_t)(unsafe.Pointer(state + 7*8)) += h
	/* Clean up */
	v10 = libc.Uint64FromInt32(0)
	T2 = v10
	v9 = v10
	T1 = v9
	v8 = v9
	h = v8
	v7 = v8
	g = v7
	v6 = v7
	f = v6
	v5 = v6
	e = v5
	v4 = v5
	d = v4
	v3 = v4
	c = v3
	v2 = v3
	b = v2
	a = v2
}

func XSHA512Update(tls *libc.TLS, context uintptr, data uintptr, len1 Tsize_t) {
	var freespace, usedspace, v1, v2 Tsize_t
	_, _, _, _ = freespace, usedspace, v1, v2
	/* Calling with no data is valid (we do nothing) */
	if len1 == uint32(0) {
		return
	}
	usedspace = uint32(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> libc.Int32FromInt32(3) % uint64(m_SHA512_BLOCK_LENGTH))
	if usedspace > uint32(0) {
		/* Calculate how much free space is available in the buffer */
		freespace = uint32(m_SHA512_BLOCK_LENGTH) - usedspace
		if len1 >= freespace {
			/* Fill the buffer completely and process it */
			libc.Xmemcpy(tls, context+80+uintptr(usedspace), data, freespace)
			*(*Tuint64_t)(unsafe.Pointer(context + 64)) += uint64(freespace << libc.Int32FromInt32(3))
			if *(*Tuint64_t)(unsafe.Pointer(context + 64)) < uint64(freespace<<libc.Int32FromInt32(3)) {
				*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8))++
			}
			len1 -= freespace
			data += uintptr(freespace)
			XSHA512Transform(tls, context, context+80)
		} else {
			/* The buffer is not yet full */
			libc.Xmemcpy(tls, context+80+uintptr(usedspace), data, len1)
			*(*Tuint64_t)(unsafe.Pointer(context + 64)) += uint64(len1 << libc.Int32FromInt32(3))
			if *(*Tuint64_t)(unsafe.Pointer(context + 64)) < uint64(len1<<libc.Int32FromInt32(3)) {
				*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8))++
			}
			/* Clean up: */
			v1 = libc.Uint32FromInt32(0)
			freespace = v1
			usedspace = v1
			return
		}
	}
	for len1 >= uint32(m_SHA512_BLOCK_LENGTH) {
		/* Process as many complete blocks as we can */
		XSHA512Transform(tls, context, data)
		*(*Tuint64_t)(unsafe.Pointer(context + 64)) += libc.Uint64FromInt32(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH) << libc.Int32FromInt32(3))
		if *(*Tuint64_t)(unsafe.Pointer(context + 64)) < libc.Uint64FromInt32(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)<<libc.Int32FromInt32(3)) {
			*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8))++
		}
		len1 -= uint32(m_SHA512_BLOCK_LENGTH)
		data += uintptr(m_SHA512_BLOCK_LENGTH)
	}
	if len1 > uint32(0) {
		/* There's left-overs, so save 'em */
		libc.Xmemcpy(tls, context+80, data, len1)
		*(*Tuint64_t)(unsafe.Pointer(context + 64)) += uint64(len1 << libc.Int32FromInt32(3))
		if *(*Tuint64_t)(unsafe.Pointer(context + 64)) < uint64(len1<<libc.Int32FromInt32(3)) {
			*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8))++
		}
	}
	/* Clean up: */
	v2 = libc.Uint32FromInt32(0)
	freespace = v2
	usedspace = v2
}

func XSHA512Pad(tls *libc.TLS, context uintptr) {
	var usedspace, v1 uint32
	_, _ = usedspace, v1
	usedspace = uint32(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> libc.Int32FromInt32(3) % uint64(m_SHA512_BLOCK_LENGTH))
	if usedspace > uint32(0) {
		/* Begin padding with a 1 bit: */
		v1 = usedspace
		usedspace++
		*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(v1))) = uint8(0x80)
		if usedspace <= libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) {
			/* Set-up for the last transform: */
			libc.Xmemset(tls, context+80+uintptr(usedspace), 0, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16))-usedspace)
		} else {
			if usedspace < uint32(m_SHA512_BLOCK_LENGTH) {
				libc.Xmemset(tls, context+80+uintptr(usedspace), 0, uint32(m_SHA512_BLOCK_LENGTH)-usedspace)
			}
			/* Do second-to-last transform: */
			XSHA512Transform(tls, context, context+80)
			/* And set-up for the last transform: */
			libc.Xmemset(tls, context+80, 0, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(2)))
		}
	} else {
		/* Prepare for final transform: */
		libc.Xmemset(tls, context+80, 0, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)))
		/* Begin padding with a 1 bit: */
		*(*Tuint8_t)(unsafe.Pointer(context + 80)) = uint8(0x80)
	}
	/* Store the length of input data (in bits) in big endian format: */
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)))) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(56))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 1)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(48))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 2)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(40))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 3)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(32))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 4)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(24))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 5)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(16))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 6)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) >> int32(8))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)) + 7)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)))) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(56))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 1)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(48))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 2)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(40))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 3)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(32))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 4)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(24))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 5)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(16))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 6)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)) >> int32(8))
	*(*Tuint8_t)(unsafe.Pointer(context + 80 + uintptr(libc.Int32FromInt32(m_SHA512_BLOCK_LENGTH)-libc.Int32FromInt32(16)+libc.Int32FromInt32(8)) + 7)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + 64)))
	/* Final transform: */
	XSHA512Transform(tls, context, context+80)
	/* Clean up: */
	usedspace = uint32(0)
}

func XSHA512Final(tls *libc.TLS, digest uintptr, context uintptr) {
	var i int32
	_ = i
	XSHA512Pad(tls, context)
	/* If no digest buffer is passed, we don't bother doing this: */
	if digest != libc.UintptrFromInt32(0) {
		/* Convert TO host byte order */
		i = 0
		for {
			if !(i < int32(8)) {
				break
			}
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)))) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(56))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 1)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(48))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 2)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(40))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 3)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(32))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 4)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(24))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 5)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(16))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 6)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(8))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 7)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)))
			goto _1
		_1:
			;
			i++
		}
		libc.Xmemset(tls, context, 0, uint32(208))
	}
}

// C documentation
//
//	/*** SHA-384: *********************************************************/
func XSHA384Init(tls *libc.TLS, context uintptr) {
	var v1 Tuint64_t
	_ = v1
	if context == libc.UintptrFromInt32(0) {
		return
	}
	libc.Xmemcpy(tls, context, uintptr(unsafe.Pointer(&_sha384_initial_hash_value)), uint32(64))
	libc.Xmemset(tls, context+80, 0, uint32(128))
	v1 = libc.Uint64FromInt32(0)
	*(*Tuint64_t)(unsafe.Pointer(context + 64 + 1*8)) = v1
	*(*Tuint64_t)(unsafe.Pointer(context + 64)) = v1
}

func XSHA384Final(tls *libc.TLS, digest uintptr, context uintptr) {
	var i int32
	_ = i
	XSHA512Pad(tls, context)
	/* If no digest buffer is passed, we don't bother doing this: */
	if digest != libc.UintptrFromInt32(0) {
		/* Convert TO host byte order */
		i = 0
		for {
			if !(i < int32(6)) {
				break
			}
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)))) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(56))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 1)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(48))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 2)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(40))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 3)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(32))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 4)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(24))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 5)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(16))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 6)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)) >> int32(8))
			*(*Tuint8_t)(unsafe.Pointer(digest + uintptr(i*libc.Int32FromInt32(8)) + 7)) = uint8(*(*Tuint64_t)(unsafe.Pointer(context + uintptr(i)*8)))
			goto _1
		_1:
			;
			i++
		}
	}
	/* Zero out state data */
	libc.Xmemset(tls, context, 0, uint32(208))
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
const m_L_INCR = 1
const m_L_SET = 0
const m_L_XTND = 2
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
const m_O_RDONLY = 0
const m_O_RDWR = 02
const m_O_RSYNC = 04010000
const m_O_SEARCH = "O_PATH"
const m_O_SYNC = 04010000
const m_O_TMPFILE = 020200000
const m_O_TRUNC = 01000
const m_O_TTY_INIT = 0
const m_O_WRONLY = 01
const m_POSIX_CLOSE_RESTART = 0
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
const m_P_tmpdir = "/tmp"
const m_RAND_MAX = 0x7fffffff
const m_RWF_WRITE_LIFE_NOT_SET = 0
const m_RWH_WRITE_LIFE_EXTREME = 5
const m_RWH_WRITE_LIFE_LONG = 4
const m_RWH_WRITE_LIFE_MEDIUM = 3
const m_RWH_WRITE_LIFE_NONE = 1
const m_RWH_WRITE_LIFE_SHORT = 2
const m_R_OK = 4
const m_SEEK_DATA = 3
const m_SEEK_HOLE = 4
const m_SPLICE_F_GIFT = 8
const m_SPLICE_F_MORE = 4
const m_SPLICE_F_MOVE = 1
const m_SPLICE_F_NONBLOCK = 2
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
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_SYNC_FILE_RANGE_WAIT_AFTER = 4
const m_SYNC_FILE_RANGE_WAIT_BEFORE = 1
const m_SYNC_FILE_RANGE_WRITE = 2
const m_S_IEXEC = "S_IXUSR"
const m_S_IFBLK = 0060000
const m_S_IFCHR = 0020000
const m_S_IFDIR = 0040000
const m_S_IFIFO = 0010000
const m_S_IFLNK = 0120000
const m_S_IFMT = 0170000
const m_S_IFREG = 0100000
const m_S_IFSOCK = 0140000
const m_S_IREAD = "S_IRUSR"
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
const m_S_IWRITE = "S_IWUSR"
const m_S_IWUSR = 0200
const m_S_IXGRP = 0010
const m_S_IXOTH = 0001
const m_S_IXUSR = 0100
const m_TMP_MAX = 10000
const m_UTIME_NOW = 0x3fffffff
const m_UTIME_OMIT = 0x3ffffffe
const m_WNOHANG = 1
const m_WUNTRACED = 2
const m_W_OK = 2
const m_X_OK = 1
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
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
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
const m__POSIX_V6_ILP32_OFFBIG = 1
const m__POSIX_V7_ILP32_OFFBIG = 1
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
const m__XOPEN_ENH_I18N = 1
const m__XOPEN_UNIX = 1
const m__XOPEN_VERSION = 700
const m_alloca = "__builtin_alloca"
const m_loff_t = "off_t"

type Tstat = struct {
	Fst_dev             Tdev_t
	F__st_dev_padding   int32
	F__st_ino_truncated int32
	Fst_mode            Tmode_t
	Fst_nlink           Tnlink_t
	Fst_uid             Tuid_t
	Fst_gid             Tgid_t
	Fst_rdev            Tdev_t
	F__st_rdev_padding  int32
	Fst_size            Toff_t
	Fst_blksize         Tblksize_t
	Fst_blocks          Tblkcnt_t
	F__st_atim32        struct {
		Ftv_sec  int32
		Ftv_nsec int32
	}
	F__st_mtim32 struct {
		Ftv_sec  int32
		Ftv_nsec int32
	}
	F__st_ctim32 struct {
		Ftv_sec  int32
		Ftv_nsec int32
	}
	Fst_ino  Tino_t
	Fst_atim Ttimespec
	Fst_mtim Ttimespec
	Fst_ctim Ttimespec
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

// C documentation
//
//	/* ARGSUSED */
func XMD2End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [16]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, uint32(m_MD2_DIGEST_STRING_LENGTH))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XMD2Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_MD2_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex[libc.Int32FromUint8((*(*[16]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex[libc.Int32FromUint8((*(*[16]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(16))
	return buf
}

var _hex = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XMD2FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1248)
	defer tls.Free(1248)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TMD2_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XMD2Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XMD2Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XMD2End(tls, bp+1168, buf)
	}
	return v5
}

func XMD2File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XMD2FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XMD2Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var _ /* ctx at bp+0 */ TMD2_CTX
	XMD2Init(tls, bp)
	XMD2Update(tls, bp, data, len1)
	return XMD2End(tls, bp, buf)
}

// C documentation
//
//	/* ARGSUSED */
func XMD4End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [16]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_MD4_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XMD4Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_MD4_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex1[libc.Int32FromUint8((*(*[16]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex1[libc.Int32FromUint8((*(*[16]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(16))
	return buf
}

var _hex1 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XMD4FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1264)
	defer tls.Free(1264)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TMD4_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XMD4Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XMD4Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XMD4End(tls, bp+1168, buf)
	}
	return v5
}

func XMD4File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XMD4FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XMD4Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* ctx at bp+0 */ TMD4_CTX
	XMD4Init(tls, bp)
	XMD4Update(tls, bp, data, len1)
	return XMD4End(tls, bp, buf)
}

/* Avoid polluting the namespace. Even though this makes this usage
 * implementation-specific, defining it unconditionally should not be
 * a problem, and better than possibly breaking unexpecting code. */

// C documentation
//
//	/* ARGSUSED */
func XMD5End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [16]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_MD5_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XMD5Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_MD5_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex2[libc.Int32FromUint8((*(*[16]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex2[libc.Int32FromUint8((*(*[16]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(16))
	return buf
}

var _hex2 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XMD5FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1264)
	defer tls.Free(1264)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TMD5_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XMD5Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XMD5Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XMD5End(tls, bp+1168, buf)
	}
	return v5
}

func XMD5File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XMD5FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XMD5Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* ctx at bp+0 */ TMD5_CTX
	XMD5Init(tls, bp)
	XMD5Update(tls, bp, data, len1)
	return XMD5End(tls, bp, buf)
}

// C documentation
//
//	/* ARGSUSED */
func XRMD160End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [20]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_RMD160_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XRMD160Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_RMD160_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex3[libc.Int32FromUint8((*(*[20]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex3[libc.Int32FromUint8((*(*[20]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(20))
	return buf
}

var _hex3 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XRMD160FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1264)
	defer tls.Free(1264)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TRMD160_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XRMD160Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XRMD160Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XRMD160End(tls, bp+1168, buf)
	}
	return v5
}

func XRMD160File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XRMD160FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XRMD160Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* ctx at bp+0 */ TRMD160_CTX
	XRMD160Init(tls, bp)
	XRMD160Update(tls, bp, data, len1)
	return XRMD160End(tls, bp, buf)
}

// C documentation
//
//	/* ARGSUSED */
func XSHA1End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [20]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA1_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XSHA1Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_SHA1_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex4[libc.Int32FromUint8((*(*[20]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex4[libc.Int32FromUint8((*(*[20]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(20))
	return buf
}

var _hex4 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XSHA1FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1264)
	defer tls.Free(1264)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TSHA1_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XSHA1Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XSHA1Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XSHA1End(tls, bp+1168, buf)
	}
	return v5
}

func XSHA1File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XSHA1FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XSHA1Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* ctx at bp+0 */ TSHA1_CTX
	XSHA1Init(tls, bp)
	XSHA1Update(tls, bp, data, len1)
	return XSHA1End(tls, bp, buf)
}

// C documentation
//
//	/* ARGSUSED */
func XSHA256End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [32]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA256_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XSHA256Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_SHA256_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex5[libc.Int32FromUint8((*(*[32]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex5[libc.Int32FromUint8((*(*[32]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(32))
	return buf
}

var _hex5 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XSHA256FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1376)
	defer tls.Free(1376)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TSHA2_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XSHA256Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XSHA256Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XSHA256End(tls, bp+1168, buf)
	}
	return v5
}

func XSHA256File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XSHA256FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XSHA256Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var _ /* ctx at bp+0 */ TSHA2_CTX
	XSHA256Init(tls, bp)
	XSHA256Update(tls, bp, data, len1)
	return XSHA256End(tls, bp, buf)
}

// C documentation
//
//	/* ARGSUSED */
func XSHA384End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [48]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA384_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XSHA384Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_SHA384_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex6[libc.Int32FromUint8((*(*[48]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex6[libc.Int32FromUint8((*(*[48]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(48))
	return buf
}

var _hex6 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XSHA384FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1376)
	defer tls.Free(1376)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TSHA2_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XSHA384Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XSHA512Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XSHA384End(tls, bp+1168, buf)
	}
	return v5
}

func XSHA384File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XSHA384FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XSHA384Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var _ /* ctx at bp+0 */ TSHA2_CTX
	XSHA384Init(tls, bp)
	XSHA512Update(tls, bp, data, len1)
	return XSHA384End(tls, bp, buf)
}

// C documentation
//
//	/* ARGSUSED */
func XSHA512End(tls *libc.TLS, ctx uintptr, buf uintptr) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i int32
	var v1 uintptr
	var v2 bool
	var _ /* digest at bp+0 */ [64]Tuint8_t
	_, _, _ = i, v1, v2
	if v2 = buf == libc.UintptrFromInt32(0); v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(libc.Int32FromInt32(m_SHA512_DIGEST_LENGTH)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))
		buf = v1
	}
	if v2 && v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	XSHA512Final(tls, bp, ctx)
	i = 0
	for {
		if !(i < int32(m_SHA512_DIGEST_LENGTH)) {
			break
		}
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = _hex7[libc.Int32FromUint8((*(*[64]Tuint8_t)(unsafe.Pointer(bp)))[i])>>int32(4)]
		*(*int8)(unsafe.Pointer(buf + uintptr(i+i+int32(1)))) = _hex7[libc.Int32FromUint8((*(*[64]Tuint8_t)(unsafe.Pointer(bp)))[i])&int32(0x0f)]
		goto _3
	_3:
		;
		i++
	}
	*(*int8)(unsafe.Pointer(buf + uintptr(i+i))) = int8('\000')
	libc.Xmemset(tls, bp, 0, uint32(64))
	return buf
}

var _hex7 = [17]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func XSHA512FileChunk(tls *libc.TLS, filename uintptr, buf uintptr, off Toff_t, len1 Toff_t) (r uintptr) {
	bp := tls.Alloc(1376)
	defer tls.Free(1376)
	var fd, save_errno, v1 int32
	var nr, v2 Tssize_t
	var v3 int64
	var v4 bool
	var v5 uintptr
	var _ /* buffer at bp+144 */ [1024]Tuint8_t
	var _ /* ctx at bp+1168 */ TSHA2_CTX
	var _ /* sb at bp+0 */ Tstat
	_, _, _, _, _, _, _, _ = fd, nr, save_errno, v1, v2, v3, v4, v5
	XSHA512Init(tls, bp+1168)
	v1 = libc.Xopen(tls, filename, m_O_RDONLY, 0)
	fd = v1
	if v1 < 0 {
		return libc.UintptrFromInt32(0)
	}
	if len1 == 0 {
		if libc.Xfstat(tls, fd, bp) == -int32(1) {
			libc.Xclose(tls, fd)
			return libc.UintptrFromInt32(0)
		}
		len1 = (*(*Tstat)(unsafe.Pointer(bp))).Fst_size
	}
	if off > 0 && libc.Xlseek(tls, fd, off, 0) < 0 {
		libc.Xclose(tls, fd)
		return libc.UintptrFromInt32(0)
	}
	for {
		if libc.Int64FromInt64(1024) < len1 {
			v3 = libc.Int64FromInt64(1024)
		} else {
			v3 = len1
		}
		v2 = libc.Xread(tls, fd, bp+144, libc.Uint32FromInt64(v3))
		nr = v2
		if !(v2 > 0) {
			break
		}
		XSHA512Update(tls, bp+1168, bp+144, libc.Uint32FromInt32(nr))
		if v4 = len1 > 0; v4 {
			len1 -= int64(nr)
		}
		if v4 && len1 == 0 {
			break
		}
	}
	save_errno = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	libc.Xclose(tls, fd)
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = save_errno
	if nr < 0 {
		v5 = libc.UintptrFromInt32(0)
	} else {
		v5 = XSHA512End(tls, bp+1168, buf)
	}
	return v5
}

func XSHA512File(tls *libc.TLS, filename uintptr, buf uintptr) (r uintptr) {
	return XSHA512FileChunk(tls, filename, buf, libc.Int64FromInt32(0), libc.Int64FromInt32(0))
}

func XSHA512Data(tls *libc.TLS, data uintptr, len1 Tsize_t, buf uintptr) (r uintptr) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var _ /* ctx at bp+0 */ TSHA2_CTX
	XSHA512Init(tls, bp)
	XSHA512Update(tls, bp, data, len1)
	return XSHA512End(tls, bp, buf)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "\x00\x01\x00\x02\x02\x00\x03\x03\x03\x00\x04\x04\x04\x04\x00\x05\x05\x05\x05\x05\x00\x06\x06\x06\x06\x06\x06\x00\a\a\a\a\a\a\a\x00\b\b\b\b\b\b\b\b\x00\t\t\t\t\t\t\t\t\t\x00\n\n\n\n\n\n\n\n\n\n\x00\v\v\v\v\v\v\v\v\v\v\v\x00\f\f\f\f\f\f\f\f\f\f\f\f\x00\r\r\r\r\r\r\r\r\r\r\r\r\r\x00\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x00\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x0f\x00\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x00\x80\x00\x00\x00"
