// Code generated for linux/ppc64le by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -ignore-link-errors -extended-errors -o libxau.go --package-name libxau .libs/libXau.a', DO NOT EDIT.

//go:build linux && ppc64le

package libxau

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const m_BUFSIZ = 1024
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_FILENAME_MAX = 4096
const m_FOPEN_MAX = 1000
const m_FUNCPROTO = 15
const m_FamilyKrb5Principal = 253
const m_FamilyLocal = 256
const m_FamilyLocalHost = 252
const m_FamilyNetname = 254
const m_FamilyWild = 65535
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_INTTYPES_H = 1
const m_HAVE_MEMORY_H = 1
const m_HAVE_PATHCONF = 1
const m_HAVE_STDINT_H = 1
const m_HAVE_STDLIB_H = 1
const m_HAVE_STRINGS_H = 1
const m_HAVE_STRING_H = 1
const m_HAVE_SYS_STAT_H = 1
const m_HAVE_SYS_TYPES_H = 1
const m_HAVE_UNISTD_H = 1
const m_LOCK_ERROR = 1
const m_LOCK_SUCCESS = 0
const m_LOCK_TIMEOUT = 2
const m_LT_OBJDIR = ".libs/"
const m_L_ctermid = 20
const m_L_cuserid = 20
const m_L_tmpnam = 20
const m_NDEBUG = 1
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 0
const m_PACKAGE = "libXau"
const m_PACKAGE_BUGREPORT = "https://gitlab.freedesktop.org/xorg/lib/libXau/issues"
const m_PACKAGE_NAME = "libXau"
const m_PACKAGE_STRING = "libXau 1.0.9"
const m_PACKAGE_TARNAME = "libXau"
const m_PACKAGE_URL = ""
const m_PACKAGE_VERSION = "1.0.9"
const m_PACKAGE_VERSION_MAJOR = 1
const m_PACKAGE_VERSION_MINOR = 0
const m_PACKAGE_VERSION_PATCHLEVEL = 9
const m_P_tmpdir = "/tmp"
const m_RAND_MAX = 0x7fffffff
const m_STDC_HEADERS = 1
const m_TMP_MAX = 10000
const m_VERSION = "1.0.9"
const m_WNOHANG = 1
const m_WUNTRACED = 2
const m_XTHREADS = 1
const m_XUSE_MTSAFE_API = 1
const m__ALL_SOURCE = 1
const m__ARCH_PPC = 1
const m__ARCH_PPC64 = 1
const m__ARCH_PPCGR = 1
const m__ARCH_PPCSQ = 1
const m__ARCH_PWR4 = 1
const m__ARCH_PWR5 = 1
const m__ARCH_PWR5X = 1
const m__ARCH_PWR6 = 1
const m__ARCH_PWR7 = 1
const m__ARCH_PWR8 = 1
const m__CALL_ELF = 2
const m__CALL_LINUX = 1
const m__GNU_SOURCE = 1
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
const m__LITTLE_ENDIAN = 1
const m__LP64 = 1
const m__POSIX_PTHREAD_SEMANTICS = 1
const m__STDC_PREDEF_H = 1
const m__TANDEM_SOURCE = 1
const m__X_INLINE = "inline"
const m__X_RESTRICT_KYWD = "restrict"
const m__Xconst = "const"
const m___ALTIVEC__ = 1
const m___APPLE_ALTIVEC__ = 1
const m___ATOMIC_ACQUIRE = 2
const m___ATOMIC_ACQ_REL = 4
const m___ATOMIC_CONSUME = 1
const m___ATOMIC_RELAXED = 0
const m___ATOMIC_RELEASE = 3
const m___ATOMIC_SEQ_CST = 5
const m___BIGGEST_ALIGNMENT__ = 16
const m___BIG_ENDIAN = 4321
const m___BUILTIN_CPU_SUPPORTS__ = 1
const m___BYTE_ORDER = 1234
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___CCGO__ = 1
const m___CHAR_BIT__ = 8
const m___CHAR_UNSIGNED__ = 1
const m___CMODEL_MEDIUM__ = 1
const m___CRYPTO__ = 1
const m___DBL_DECIMAL_DIG__ = 17
const m___DBL_DIG__ = 15
const m___DBL_HAS_DENORM__ = 1
const m___DBL_HAS_INFINITY__ = 1
const m___DBL_HAS_QUIET_NAN__ = 1
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
const m___DECIMAL_DIG__ = 17
const m___DEC_EVAL_METHOD__ = 2
const m___ELF__ = 1
const m___EXTENSIONS__ = 1
const m___FINITE_MATH_ONLY__ = 0
const m___FLOAT128_TYPE__ = 1
const m___FLOAT128__ = 1
const m___FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___FLT128_DECIMAL_DIG__ = 36
const m___FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___FLT128_DIG__ = 33
const m___FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___FLT128_HAS_DENORM__ = 1
const m___FLT128_HAS_INFINITY__ = 1
const m___FLT128_HAS_QUIET_NAN__ = 1
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
const m___FP_FAST_FMAL = 1
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
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const m___GCC_IEC_559 = 2
const m___GCC_IEC_559_COMPLEX = 2
const m___GNUC_MINOR__ = 2
const m___GNUC_PATCHLEVEL__ = 1
const m___GNUC_STDC_INLINE__ = 1
const m___GNUC__ = 10
const m___GXX_ABI_VERSION = 1014
const m___HAVE_BSWAP__ = 1
const m___HAVE_SPECULATION_SAFE_VALUE = 1
const m___HTM__ = 1
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
const m___LDBL_DECIMAL_DIG__ = 17
const m___LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___LDBL_DIG__ = 15
const m___LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___LDBL_HAS_DENORM__ = 1
const m___LDBL_HAS_INFINITY__ = 1
const m___LDBL_HAS_QUIET_NAN__ = 1
const m___LDBL_MANT_DIG__ = 53
const m___LDBL_MAX_10_EXP__ = 308
const m___LDBL_MAX_EXP__ = 1024
const m___LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const m___LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const m___LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___LITTLE_ENDIAN = 1234
const m___LITTLE_ENDIAN__ = 1
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
const m___PIC__ = 2
const m___PIE__ = 2
const m___POWER8_VECTOR__ = 1
const m___PPC64__ = 1
const m___PPC__ = 1
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTRDIFF_MAX__ = 0x7fffffffffffffff
const m___PTRDIFF_WIDTH__ = 64
const m___QUAD_MEMORY_ATOMIC__ = 1
const m___RECIPF__ = 1
const m___RECIP_PRECISION__ = 1
const m___RECIP__ = 1
const m___RSQRTEF__ = 1
const m___RSQRTE__ = 1
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
const m___STDC_ISO_10646__ = 201706
const m___STDC_UTF_16__ = 1
const m___STDC_UTF_32__ = 1
const m___STDC_VERSION__ = 201710
const m___STDC__ = 1
const m___STRUCT_PARM_ALIGN__ = 16
const m___TM_FENCE__ = 1
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
const m___VEC_ELEMENT_REG_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___VEC__ = 10206
const m___VERSION__ = "10.2.1 20210110"
const m___VSX__ = 1
const m___WCHAR_MAX__ = 0x7fffffff
const m___WCHAR_TYPE__ = "int"
const m___WCHAR_WIDTH__ = 32
const m___WINT_MAX__ = 0xffffffff
const m___WINT_MIN__ = 0
const m___WINT_WIDTH__ = 32
const m___builtin_copysignq = "__builtin_copysignf128"
const m___builtin_fabsq = "__builtin_fabsf128"
const m___builtin_huge_valq = "__builtin_huge_valf128"
const m___builtin_infq = "__builtin_inff128"
const m___builtin_nanq = "__builtin_nanf128"
const m___builtin_nansq = "__builtin_nansf128"
const m___builtin_vsx_vperm = "__builtin_vec_perm"
const m___builtin_vsx_xvmaddadp = "__builtin_vsx_xvmadddp"
const m___builtin_vsx_xvmaddasp = "__builtin_vsx_xvmaddsp"
const m___builtin_vsx_xvmaddmdp = "__builtin_vsx_xvmadddp"
const m___builtin_vsx_xvmaddmsp = "__builtin_vsx_xvmaddsp"
const m___builtin_vsx_xvmsubadp = "__builtin_vsx_xvmsubdp"
const m___builtin_vsx_xvmsubasp = "__builtin_vsx_xvmsubsp"
const m___builtin_vsx_xvmsubmdp = "__builtin_vsx_xvmsubdp"
const m___builtin_vsx_xvmsubmsp = "__builtin_vsx_xvmsubsp"
const m___builtin_vsx_xvnmaddadp = "__builtin_vsx_xvnmadddp"
const m___builtin_vsx_xvnmaddasp = "__builtin_vsx_xvnmaddsp"
const m___builtin_vsx_xvnmaddmdp = "__builtin_vsx_xvnmadddp"
const m___builtin_vsx_xvnmaddmsp = "__builtin_vsx_xvnmaddsp"
const m___builtin_vsx_xvnmsubadp = "__builtin_vsx_xvnmsubdp"
const m___builtin_vsx_xvnmsubasp = "__builtin_vsx_xvnmsubsp"
const m___builtin_vsx_xvnmsubmdp = "__builtin_vsx_xvnmsubdp"
const m___builtin_vsx_xvnmsubmsp = "__builtin_vsx_xvnmsubsp"
const m___builtin_vsx_xxland = "__builtin_vec_and"
const m___builtin_vsx_xxlandc = "__builtin_vec_andc"
const m___builtin_vsx_xxlnor = "__builtin_vec_nor"
const m___builtin_vsx_xxlor = "__builtin_vec_or"
const m___builtin_vsx_xxlxor = "__builtin_vec_xor"
const m___builtin_vsx_xxsel = "__builtin_vec_sel"
const m___float128 = "__ieee128"
const m___gnu_linux__ = 1
const m___inline = "inline"
const m___linux = 1
const m___linux__ = 1
const m___pic__ = 2
const m___pie__ = 2
const m___powerpc64__ = 1
const m___powerpc__ = 1
const m___restrict = "restrict"
const m___restrict_arr = "restrict"
const m___unix = 1
const m___unix__ = 1
const m_alloca = "__builtin_alloca"
const m_linux = 1
const m_unix = 1

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

type Tsize_t = uint64

type Tlocale_t = uintptr

type Tssize_t = int64

type Toff_t = int64

type Tva_list = uintptr

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

func XXauDisposeAuth(tls *libc.TLS, auth uintptr) {
	if auth != 0 {
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Faddress)
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fnumber)
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fname)
		if (*TXauth)(unsafe.Pointer(auth)).Fdata != 0 {
			libc.Xmemset(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata, 0, uint64((*TXauth)(unsafe.Pointer(auth)).Fdata_length))
			libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata)
		}
		libc.Xfree(tls, auth)
	}
	return
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
const m_BIG_ENDIAN = "__BIG_ENDIAN"
const m_BYTE_ORDER = "__BYTE_ORDER"
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
const m_FD_SETSIZE = 1024
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
const m_ITIMER_PROF = 2
const m_ITIMER_REAL = 0
const m_ITIMER_VIRTUAL = 1
const m_LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const m_L_INCR = 1
const m_L_SET = 0
const m_L_XTND = 2
const m_MAX_HANDLE_SZ = 128
const m_O_APPEND = 02000
const m_O_ASYNC = 020000
const m_O_CLOEXEC = 02000000
const m_O_CREAT = 0100
const m_O_DIRECT = 0400000
const m_O_DIRECTORY = 040000
const m_O_DSYNC = 010000
const m_O_EXCL = 0200
const m_O_EXEC = "O_PATH"
const m_O_LARGEFILE = 0200000
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
const m_PDP_ENDIAN = "__PDP_ENDIAN"
const m_POSIX_CLOSE_RESTART = 0
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
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
const m_TIMER_ABSTIME = 1
const m_TIME_UTC = 1
const m_W_OK = 2
const m_X_BIG_ENDIAN = "BIG_ENDIAN"
const m_X_BYTE_ORDER = "BYTE_ORDER"
const m_X_LITTLE_ENDIAN = "LITTLE_ENDIAN"
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
const m__XOPEN_ENH_I18N = 1
const m__XOPEN_UNIX = 1
const m__XOPEN_VERSION = 700
const m___PDP_ENDIAN = 3412
const m___tm_gmtoff = "tm_gmtoff"
const m___tm_zone = "tm_zone"
const m_loff_t = "off_t"
const m_static_assert = "_Static_assert"

type Tregister_t = int64

type Ttime_t = int64

type Tsuseconds_t = int64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tu_int64_t = uint64

type Tmode_t = uint32

type Tnlink_t = uint64

type Tino_t = uint64

type Tdev_t = uint64

type Tblksize_t = int64

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
	if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(name))) == int32('/') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(name + 1))) == int32('\000') {
		v2 = int32(1)
	} else {
		v2 = 0
	}
	libc.X__builtin_snprintf(tls, _buf, _bsize, __ccgo_ts+29, libc.VaList(bp+8, name, slashDotXauthority+uintptr(v2)))
	return _buf
}

var _bsize Tsize_t

var _atexit_registered int32

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

/* 'endian.h' might have been included before 'Xarch.h' */

func XXauGetAuthByAddr(tls *libc.TLS, family uint16, address_length uint16, address uintptr, number_length uint16, number uintptr, name_length uint16, name uintptr) (r uintptr) {
	var auth_file, auth_name, entry uintptr
	_, _, _ = auth_file, auth_name, entry
	auth_name = XXauFileName(tls)
	if !(auth_name != 0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xaccess(tls, auth_name, int32(m_R_OK)) != 0 { /* checks REAL id */
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
		if (libc.Int32FromUint16(family) == int32(m_FamilyWild) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Ffamily) == int32(m_FamilyWild) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Ffamily) == libc.Int32FromUint16(family) && libc.Int32FromUint16(address_length) == libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Faddress_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Faddress, address, uint64(address_length)) == 0) && (libc.Int32FromUint16(number_length) == 0 || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) == 0 || libc.Int32FromUint16(number_length) == libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Fnumber, number, uint64(number_length)) == 0) && (libc.Int32FromUint16(name_length) == 0 || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fname_length) == 0 || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fname_length) == libc.Int32FromUint16(name_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Fname, name, uint64(name_length)) == 0) {
			break
		}
		XXauDisposeAuth(tls, entry)
		goto _1
	_1:
	}
	libc.Xfclose(tls, auth_file)
	return entry
}

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
const m_SCHED_BATCH = 3
const m_SCHED_DEADLINE = 6
const m_SCHED_FIFO = 1
const m_SCHED_IDLE = 5
const m_SCHED_OTHER = 0
const m_SCHED_RESET_ON_FORK = 0x40000000
const m_SCHED_RR = 2
const m_XMUTEX_INITIALIZER = "PTHREAD_MUTEX_INITIALIZER"
const m_xfree = "free"
const m_xmalloc = "malloc"
const m_xthread_self = "pthread_self"

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

type t__ptcb = struct {
	F__f    uintptr
	F__x    uintptr
	F__next uintptr
}

type Tcpu_set_t1 = struct {
	F__bits [16]uint64
}

type Txthread_t = uintptr

type Txthread_key_t = uint32

type Txcondition_rec = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][6]uintptr
		F__i  [12]int32
	}
}

type Txmutex_rec = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

type Txcondition_t = uintptr

type Txmutex_t = uintptr

/* aids understood by some debuggers */

func XXauGetBestAuthByAddr(tls *libc.TLS, family uint16, address_length uint16, address uintptr, number_length uint16, number uintptr, types_length int32, types uintptr, type_lengths uintptr) (r uintptr) {
	var auth_file, auth_name, best, entry uintptr
	var best_type, type1 int32
	_, _, _, _, _, _ = auth_file, auth_name, best, best_type, entry, type1
	auth_name = XXauFileName(tls)
	if !(auth_name != 0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xaccess(tls, auth_name, int32(m_R_OK)) != 0 { /* checks REAL id */
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
		if (libc.Int32FromUint16(family) == int32(m_FamilyWild) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Ffamily) == int32(m_FamilyWild) || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Ffamily) == libc.Int32FromUint16(family) && (libc.Int32FromUint16(address_length) == libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Faddress_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Faddress, address, uint64(address_length)) == 0)) && (libc.Int32FromUint16(number_length) == 0 || libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) == 0 || libc.Int32FromUint16(number_length) == libc.Int32FromUint16((*TXauth)(unsafe.Pointer(entry)).Fnumber_length) && libc.Xmemcmp(tls, (*TXauth)(unsafe.Pointer(entry)).Fnumber, number, uint64(number_length)) == 0) {
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
const m_EDEADLOCK = 58
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
const m_O_CREAT1 = 64
const m_O_EXCL1 = 128
const m_O_WRONLY1 = 1
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
const m_S_IWRITE = "S_IWUSR"
const m_Time_t = "time_t"
const m_UTIME_NOW = 0x3fffffff
const m_UTIME_OMIT = 0x3ffffffe

type Tstat = struct {
	Fst_dev     Tdev_t
	Fst_ino     Tino_t
	Fst_nlink   Tnlink_t
	Fst_mode    Tmode_t
	Fst_uid     Tuid_t
	Fst_gid     Tgid_t
	Fst_rdev    Tdev_t
	Fst_size    Toff_t
	Fst_blksize Tblksize_t
	Fst_blocks  Tblkcnt_t
	Fst_atim    Ttimespec
	Fst_mtim    Ttimespec
	Fst_ctim    Ttimespec
	F__unused   [3]uint64
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

func XXauLockAuth(tls *libc.TLS, file_name uintptr, retries int32, timeout int32, dead int64) (r int32) {
	bp := tls.Alloc(2224)
	defer tls.Free(2224)
	var creat_fd int32
	var now Ttime_t
	var _ /* creat_name at bp+0 */ [1025]uint8
	var _ /* link_name at bp+1025 */ [1025]uint8
	var _ /* statb at bp+2056 */ Tstat
	_, _ = creat_fd, now
	creat_fd = -int32(1)
	if libc.Xstrlen(tls, file_name) > uint64(1022) {
		return int32(m_LOCK_ERROR)
	}
	libc.X__builtin_snprintf(tls, bp, uint64(1025), __ccgo_ts+37, libc.VaList(bp+2208, file_name))
	libc.X__builtin_snprintf(tls, bp+1025, uint64(1025), __ccgo_ts+42, libc.VaList(bp+2208, file_name))
	if libc.Xstat(tls, bp, bp+2056) != -int32(1) {
		now = libc.Xtime(tls, libc.UintptrFromInt32(0))
		/*
		 * NFS may cause ctime to be before now, special
		 * case a 0 deadtime to force lock removal
		 */
		if dead == 0 || now-(*(*Tstat)(unsafe.Pointer(bp + 2056))).Fst_ctim.Ftv_sec > dead {
			libc.Xremove(tls, bp)
			libc.Xremove(tls, bp+1025)
		}
	}
	for retries > 0 {
		if creat_fd == -int32(1) {
			creat_fd = libc.Xopen(tls, bp, libc.Int32FromInt32(m_O_WRONLY1)|libc.Int32FromInt32(m_O_CREAT1)|libc.Int32FromInt32(m_O_EXCL1), libc.VaList(bp+2208, int32(0600)))
			if creat_fd == -int32(1) {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(m_EACCES) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(m_EEXIST) {
					return int32(m_LOCK_ERROR)
				}
			} else {
				libc.Xclose(tls, creat_fd)
			}
		}
		if creat_fd != -int32(1) {
			/* The file system may not support hard links, and pathconf should tell us that. */
			if int64(1) == libc.Xpathconf(tls, bp, m__PC_LINK_MAX) {
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
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(m_ENOENT) {
					creat_fd = -int32(1) /* force re-creat next time around */
					continue
				}
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(m_EEXIST) {
					return int32(m_LOCK_ERROR)
				}
			}
		}
		libc.Xsleep(tls, libc.Uint32FromInt32(timeout))
		retries--
	}
	return int32(m_LOCK_TIMEOUT)
}

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
			libc.Xmemset(tls, data, 0, uint64(*(*uint16)(unsafe.Pointer(bp))))
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
	if _read_short(tls, bp, auth_file) == 0 {
		return libc.UintptrFromInt32(0)
	}
	if _read_counted_string(tls, bp+2, bp+8, auth_file) == 0 {
		return libc.UintptrFromInt32(0)
	}
	if _read_counted_string(tls, bp+16, bp+24, auth_file) == 0 {
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Faddress)
		return libc.UintptrFromInt32(0)
	}
	if _read_counted_string(tls, bp+32, bp+40, auth_file) == 0 {
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Faddress)
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fnumber)
		return libc.UintptrFromInt32(0)
	}
	if _read_counted_string(tls, bp+48, bp+56, auth_file) == 0 {
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Faddress)
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fnumber)
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fname)
		return libc.UintptrFromInt32(0)
	}
	ret = libc.Xmalloc(tls, uint64(64))
	if !(ret != 0) {
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Faddress)
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fnumber)
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fname)
		if (*(*TXauth)(unsafe.Pointer(bp))).Fdata != 0 {
			libc.Xmemset(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fdata, 0, uint64((*(*TXauth)(unsafe.Pointer(bp))).Fdata_length))
			libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fdata)
		}
		return libc.UintptrFromInt32(0)
	}
	*(*TXauth)(unsafe.Pointer(ret)) = *(*TXauth)(unsafe.Pointer(bp))
	return ret
}

const m_O_CREAT2 = 0100
const m_O_EXCL2 = 0200
const m_O_WRONLY2 = 01

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

/* 'endian.h' might have been included before 'Xarch.h' */

func XXauUnlockAuth(tls *libc.TLS, file_name uintptr) (r int32) {
	bp := tls.Alloc(2080)
	defer tls.Free(2080)
	var _ /* creat_name at bp+0 */ [1025]uint8
	var _ /* link_name at bp+1025 */ [1025]uint8
	if libc.Xstrlen(tls, file_name) > uint64(1022) {
		return 0
	}
	libc.X__builtin_snprintf(tls, bp, uint64(1025), __ccgo_ts+37, libc.VaList(bp+2064, file_name))
	libc.X__builtin_snprintf(tls, bp+1025, uint64(1025), __ccgo_ts+42, libc.VaList(bp+2064, file_name))
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
