// Code generated for freebsd/amd64 by 'generator --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -ignore-link-errors -extended-errors -o libxau.go --package-name libxau .libs/libXau.a', DO NOT EDIT.

//go:build freebsd && amd64

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
const m_FILENAME_MAX = 1024
const m_FOPEN_MAX = 20
const m_FUNCPROTO = 15
const m_FamilyKrb5Principal = 253
const m_FamilyLocal = 256
const m_FamilyLocalHost = 252
const m_FamilyNetname = 254
const m_FamilyWild = 65535
const m_HAVE_CONFIG_H = 1
const m_HAVE_DLFCN_H = 1
const m_HAVE_EXPLICIT_BZERO = 1
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
const m_LOCK_ERROR = 1
const m_LOCK_SUCCESS = 0
const m_LOCK_TIMEOUT = 2
const m_LT_OBJDIR = ".libs/"
const m_L_ctermid = 1024
const m_L_cuserid = 17
const m_L_tmpnam = 1024
const m_NDEBUG = 1
const m_NeedFunctionPrototypes = 1
const m_NeedNestedPrototypes = 1
const m_NeedVarargsPrototypes = 1
const m_NeedWidePrototypes = 0
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
const m_P_tmpdir = "/tmp/"
const m_RAND_MAX = 0x7fffffff
const m_SEEK_CUR = 1
const m_SEEK_END = 2
const m_SEEK_SET = 0
const m_STDC_HEADERS = 1
const m_TMP_MAX = 308915776
const m_VERSION = "1.0.11"
const m_XTHREADS = 1
const m_XUSE_MTSAFE_API = 1
const m__ALL_SOURCE = 1
const m__DARWIN_C_SOURCE = 1
const m__GNU_SOURCE = 1
const m__HPUX_ALT_XOPEN_SOCKET_API = 1
const m__IOFBF = 0
const m__IOLBF = 1
const m__IONBF = 2
const m__LP64 = 1
const m__NETBSD_SOURCE = 1
const m__OPENBSD_SOURCE = 1
const m__POSIX_PTHREAD_SEMANTICS = 1
const m__TANDEM_SOURCE = 1
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
const m___EXTENSIONS__ = 1
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
const m___S2OAP = 0x0001
const m___SALC = 0x4000
const m___SAPP = 0x0100
const m___SCHAR_MAX = 0x7f
const m___SCHAR_MAX__ = 127
const m___SEG_FS = 1
const m___SEG_GS = 1
const m___SEOF = 0x0020
const m___SERR = 0x0040
const m___SHRT_MAX = 0x7fff
const m___SHRT_MAX__ = 32767
const m___SHRT_WIDTH__ = 16
const m___SIGN = 0x8000
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
const m___SLBF = 0x0001
const m___SMBF = 0x0080
const m___SMOD = 0x2000
const m___SNBF = 0x0002
const m___SNPT = 0x0800
const m___SOFF = 0x1000
const m___SOPT = 0x0400
const m___SRD = 0x0004
const m___SRW = 0x0010
const m___SSE2_MATH__ = 1
const m___SSE2__ = 1
const m___SSE_MATH__ = 1
const m___SSE__ = 1
const m___SSIZE_MAX = "__LONG_MAX"
const m___SSTR = 0x0200
const m___STDC_HOSTED__ = 1
const m___STDC_MB_MIGHT_NEQ_WC__ = 1
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
const m___SWR = 0x0008
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
const m_stderr = "__stderrp"
const m_stdin = "__stdinp"
const m_stdout = "__stdoutp"
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

type Tsize_t = uint64

type Tlocale_t = uintptr

type Tmode_t = uint16

type Tssize_t = int64

type Trsize_t = uint64

type Terrno_t = int32

type Tfpos_t = int64

type Toff_t = int64

type Toff64_t = int64

type Tva_list = uintptr

type t__sbuf = struct {
	F_base uintptr
	F_size int32
}

type t__sFILE = struct {
	F_p           uintptr
	F_r           int32
	F_w           int32
	F_flags       int16
	F_file        int16
	F_bf          t__sbuf
	F_lbfsize     int32
	F_cookie      uintptr
	F_close       uintptr
	F_read        uintptr
	F_seek        uintptr
	F_write       uintptr
	F_ub          t__sbuf
	F_up          uintptr
	F_ur          int32
	F_ubuf        [3]uint8
	F_nbuf        [1]uint8
	F_lb          t__sbuf
	F_blksize     int32
	F_offset      Tfpos_t
	F_fl_mutex    uintptr
	F_fl_owner    uintptr
	F_fl_count    int32
	F_orientation int32
	F_mbstate     t__mbstate_t
	F_flags2      int32
}

type TFILE = struct {
	F_p           uintptr
	F_r           int32
	F_w           int32
	F_flags       int16
	F_file        int16
	F_bf          t__sbuf
	F_lbfsize     int32
	F_cookie      uintptr
	F_close       uintptr
	F_read        uintptr
	F_seek        uintptr
	F_write       uintptr
	F_ub          t__sbuf
	F_up          uintptr
	F_ur          int32
	F_ubuf        [3]uint8
	F_nbuf        [1]uint8
	F_lb          t__sbuf
	F_blksize     int32
	F_offset      Tfpos_t
	F_fl_mutex    uintptr
	F_fl_owner    uintptr
	F_fl_count    int32
	F_orientation int32
	F_mbstate     t__mbstate_t
	F_flags2      int32
}

type Tcookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

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

type Tconstraint_handler_t = uintptr

func XXauDisposeAuth(tls *libc.TLS, auth uintptr) {
	if auth != 0 {
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Faddress)
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fnumber)
		libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fname)
		if (*TXauth)(unsafe.Pointer(auth)).Fdata != 0 {
			libc.Xexplicit_bzero(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata, uint64((*TXauth)(unsafe.Pointer(auth)).Fdata_length))
			libc.Xfree(tls, (*TXauth)(unsafe.Pointer(auth)).Fdata)
		}
		libc.Xfree(tls, auth)
	}
	return
}

const m_AT_EACCESS = 0x0100
const m_AT_EMPTY_PATH = 0x4000
const m_AT_REMOVEDIR = 0x0800
const m_AT_RESOLVE_BENEATH = 0x2000
const m_AT_SYMLINK_FOLLOW = 0x0400
const m_AT_SYMLINK_NOFOLLOW = 0x0200
const m_BIG_ENDIAN = "_BIG_ENDIAN"
const m_BYTE_ORDER = "_BYTE_ORDER"
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
const m_CPUCLOCK_WHICH_PID = 0
const m_CPUCLOCK_WHICH_TID = 1
const m_DST_AUST = 2
const m_DST_CAN = 6
const m_DST_EET = 5
const m_DST_MET = 4
const m_DST_NONE = 0
const m_DST_USA = 1
const m_DST_WET = 3
const m_FAPPEND = "O_APPEND"
const m_FASYNC = "O_ASYNC"
const m_FDSYNC = "O_DSYNC"
const m_FD_CLOEXEC = 1
const m_FD_SETSIZE = 1024
const m_FFSYNC = "O_FSYNC"
const m_FNDELAY = "O_NONBLOCK"
const m_FNONBLOCK = "O_NONBLOCK"
const m_FRDAHEAD = "O_CREAT"
const m_FREAD = 0x0001
const m_FWRITE = 0x0002
const m_F_ADD_SEALS = 19
const m_F_CANCEL = 5
const m_F_DUP2FD = 10
const m_F_DUP2FD_CLOEXEC = 18
const m_F_DUPFD = 0
const m_F_DUPFD_CLOEXEC = 17
const m_F_GETFD = 1
const m_F_GETFL = 3
const m_F_GETLK = 11
const m_F_GETOWN = 5
const m_F_GET_SEALS = 20
const m_F_ISUNIONSTACK = 21
const m_F_KINFO = 22
const m_F_LOCK = 1
const m_F_OGETLK = 7
const m_F_OK = 0
const m_F_OSETLK = 8
const m_F_OSETLKW = 9
const m_F_RDAHEAD = 16
const m_F_RDLCK = 1
const m_F_READAHEAD = 15
const m_F_SEAL_GROW = 0x0004
const m_F_SEAL_SEAL = 0x0001
const m_F_SEAL_SHRINK = 0x0002
const m_F_SEAL_WRITE = 0x0008
const m_F_SETFD = 2
const m_F_SETFL = 4
const m_F_SETLK = 12
const m_F_SETLKW = 13
const m_F_SETLK_REMOTE = 14
const m_F_SETOWN = 6
const m_F_TEST = 3
const m_F_TLOCK = 2
const m_F_ULOCK = 0
const m_F_UNLCK = 2
const m_F_UNLCKSYS = 4
const m_F_WRLCK = 3
const m_ITIMER_PROF = 2
const m_ITIMER_REAL = 0
const m_ITIMER_VIRTUAL = 1
const m_KCMP_FILE = 100
const m_KCMP_FILEOBJ = 101
const m_KCMP_FILES = 102
const m_KCMP_SIGHAND = 103
const m_KCMP_VM = 104
const m_LITTLE_ENDIAN = "_LITTLE_ENDIAN"
const m_LOCK_EX = 0x02
const m_LOCK_NB = 0x04
const m_LOCK_SH = 0x01
const m_LOCK_UN = 0x08
const m_L_INCR = "SEEK_CUR"
const m_L_SET = "SEEK_SET"
const m_L_XTND = "SEEK_END"
const m_NFDBITS = "_NFDBITS"
const m_O_ACCMODE = 0x0003
const m_O_APPEND = 0x0008
const m_O_ASYNC = 0x0040
const m_O_CLOEXEC = 0x00100000
const m_O_CREAT = 0x0200
const m_O_DIRECT = 0x00010000
const m_O_DIRECTORY = 0x00020000
const m_O_DSYNC = 0x01000000
const m_O_EMPTY_PATH = 0x02000000
const m_O_EXCL = 0x0800
const m_O_EXEC = 0x00040000
const m_O_EXLOCK = 0x0020
const m_O_FSYNC = 0x0080
const m_O_NDELAY = "O_NONBLOCK"
const m_O_NOCTTY = 0x8000
const m_O_NOFOLLOW = 0x0100
const m_O_NONBLOCK = 0x0004
const m_O_PATH = 0x00400000
const m_O_RDONLY = 0x0000
const m_O_RDWR = 0x0002
const m_O_RESOLVE_BENEATH = 0x00800000
const m_O_SEARCH = "O_EXEC"
const m_O_SHLOCK = 0x0010
const m_O_SYNC = 0x0080
const m_O_TRUNC = 0x0400
const m_O_TTY_INIT = 0x00080000
const m_O_VERIFY = 0x00200000
const m_O_WRONLY = 0x0001
const m_PDP_ENDIAN = "_PDP_ENDIAN"
const m_POSIX_FADV_DONTNEED = 4
const m_POSIX_FADV_NOREUSE = 5
const m_POSIX_FADV_NORMAL = 0
const m_POSIX_FADV_RANDOM = 1
const m_POSIX_FADV_SEQUENTIAL = 2
const m_POSIX_FADV_WILLNEED = 3
const m_RFTSIGMASK = 0xFF
const m_RFTSIGSHIFT = 20
const m_R_OK = 0x04
const m_SBT_MAX = 0x7fffffffffffffff
const m_SEEK_DATA = 3
const m_SEEK_HOLE = 4
const m_SPACECTL_DEALLOC = 1
const m_SPACECTL_F_SUPPORTED = 0
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_SWAPOFF_FORCE = 0x00000001
const m_TIMER_ABSTIME = 0x1
const m_TIMER_RELTIME = 0x0
const m_TIME_MONOTONIC = 2
const m_TIME_UTC = 1
const m_W_OK = 0x02
const m_X_BIG_ENDIAN = "BIG_ENDIAN"
const m_X_BYTE_ORDER = "BYTE_ORDER"
const m_X_LITTLE_ENDIAN = "LITTLE_ENDIAN"
const m_X_OK = 0x01
const m__BYTE_ORDER = "__BYTE_ORDER__"
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
const m__PC_ACL_EXTENDED = 59
const m__PC_ACL_NFS4 = 64
const m__PC_ACL_PATH_MAX = 60
const m__PC_ALLOC_SIZE_MIN = 10
const m__PC_ASYNC_IO = 53
const m__PC_CAP_PRESENT = 61
const m__PC_CHOWN_RESTRICTED = 7
const m__PC_DEALLOC_PRESENT = 65
const m__PC_FILESIZEBITS = 12
const m__PC_INF_PRESENT = 62
const m__PC_LINK_MAX = 1
const m__PC_MAC_PRESENT = 63
const m__PC_MAX_CANON = 2
const m__PC_MAX_INPUT = 3
const m__PC_MIN_HOLE_SIZE = 21
const m__PC_NAME_MAX = 4
const m__PC_NO_TRUNC = 8
const m__PC_PATH_MAX = 5
const m__PC_PIPE_BUF = 6
const m__PC_PRIO_IO = 54
const m__PC_REC_INCR_XFER_SIZE = 14
const m__PC_REC_MAX_XFER_SIZE = 15
const m__PC_REC_MIN_XFER_SIZE = 16
const m__PC_REC_XFER_ALIGN = 17
const m__PC_SYMLINK_MAX = 18
const m__PC_SYNC_IO = 55
const m__PC_VDISABLE = 9
const m__PDP_ENDIAN = "__ORDER_PDP_ENDIAN__"
const m__POSIX2_CHAR_TERM = 1
const m__POSIX2_C_BIND = 200112
const m__POSIX2_FORT_RUN = 200112
const m__POSIX2_UPE = 200112
const m__POSIX2_VERSION = 199212
const m__POSIX_ADVISORY_INFO = 200112
const m__POSIX_ASYNCHRONOUS_IO = 200112
const m__POSIX_BARRIERS = 200112
const m__POSIX_CHOWN_RESTRICTED = 1
const m__POSIX_CPUTIME = 200112
const m__POSIX_FSYNC = 200112
const m__POSIX_IPV6 = 0
const m__POSIX_JOB_CONTROL = 1
const m__POSIX_MAPPED_FILES = 200112
const m__POSIX_MEMLOCK_RANGE = 200112
const m__POSIX_MEMORY_PROTECTION = 200112
const m__POSIX_MESSAGE_PASSING = 200112
const m__POSIX_MONOTONIC_CLOCK = 200112
const m__POSIX_NO_TRUNC = 1
const m__POSIX_PRIORITY_SCHEDULING = 0
const m__POSIX_RAW_SOCKETS = 200112
const m__POSIX_READER_WRITER_LOCKS = 200112
const m__POSIX_REALTIME_SIGNALS = 200112
const m__POSIX_REGEXP = 1
const m__POSIX_SEMAPHORES = 200112
const m__POSIX_SHARED_MEMORY_OBJECTS = 200112
const m__POSIX_SHELL = 1
const m__POSIX_SPAWN = 200112
const m__POSIX_SPIN_LOCKS = 200112
const m__POSIX_THREADS = 200112
const m__POSIX_THREAD_ATTR_STACKADDR = 200112
const m__POSIX_THREAD_ATTR_STACKSIZE = 200112
const m__POSIX_THREAD_CPUTIME = 200112
const m__POSIX_THREAD_PRIORITY_SCHEDULING = 200112
const m__POSIX_THREAD_PRIO_INHERIT = 200112
const m__POSIX_THREAD_PRIO_PROTECT = 200112
const m__POSIX_THREAD_PROCESS_SHARED = 200112
const m__POSIX_TIMEOUTS = 200112
const m__POSIX_TIMERS = 200112
const m__POSIX_VDISABLE = 0xff
const m__POSIX_VERSION = 200112
const m__QUAD_HIGHWORD = 1
const m__QUAD_LOWWORD = 0
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
const m__SC_CPUSET_SIZE = 122
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
const m__SC_MAPPED_FILES = 29
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
const m__SC_PAGESIZE = 47
const m__SC_PAGE_SIZE = "_SC_PAGESIZE"
const m__SC_PHYS_PAGES = 121
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
const m__SC_TRACE_INHERIT = 99
const m__SC_TRACE_LOG = 100
const m__SC_TTY_NAME_MAX = 101
const m__SC_TYPED_MEMORY_OBJECTS = 102
const m__SC_TZNAME_MAX = 27
const m__SC_V6_ILP32_OFF32 = 103
const m__SC_V6_ILP32_OFFBIG = 104
const m__SC_V6_LP64_OFF64 = 105
const m__SC_V6_LPBIG_OFFBIG = 106
const m__SC_VERSION = 8
const m__SC_XOPEN_CRYPT = 108
const m__SC_XOPEN_ENH_I18N = 109
const m__SC_XOPEN_LEGACY = 110
const m__SC_XOPEN_REALTIME = 111
const m__SC_XOPEN_REALTIME_THREADS = 112
const m__SC_XOPEN_SHM = 113
const m__SC_XOPEN_STREAMS = 114
const m__SC_XOPEN_UNIX = 115
const m__SC_XOPEN_VERSION = 116
const m__SC_XOPEN_XCU_VERSION = 117
const m__SIG_MAXSIG = 128
const m__SIG_WORDS = 4
const m__V6_ILP32_OFFBIG = 0
const m__V6_LP64_OFF64 = 0
const m__XOPEN_SHM = 1
const m_fds_bits = "__fds_bits"
const m_static_assert = "_Static_assert"

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

type Taccmode_t = int32

type Tnlink_t = uint64

type Tpid_t = int32

type Tregister_t = int64

type Trlim_t = int64

type Tsbintime_t = int64

type Tsegsz_t = int64

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

type Tflock = struct {
	Fl_start  Toff_t
	Fl_len    Toff_t
	Fl_pid    Tpid_t
	Fl_type   int16
	Fl_whence int16
	Fl_sysid  int32
}

type t__oflock = struct {
	Fl_start  Toff_t
	Fl_len    Toff_t
	Fl_pid    Tpid_t
	Fl_type   int16
	Fl_whence int16
}

type Tspacectl_range = struct {
	Fr_offset Toff_t
	Fr_len    Toff_t
}

type Tcrypt_data = struct {
	Finitialized int32
	F__buf       [256]int8
}

type Ttimezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Tbintime = struct {
	Fsec  Ttime_t
	Ffrac Tuint64_t
}

type Titimerval = struct {
	Fit_interval Ttimeval
	Fit_value    Ttimeval
}

type Tclockinfo = struct {
	Fhz     int32
	Ftick   int32
	Fspare  int32
	Fstathz int32
	Fprofhz int32
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
	libc.X__builtin_snprintf(tls, _buf, _bsize, __ccgo_ts+29, libc.VaList(bp+8, name, slashDotXauthority+uintptr(v2)))
	return _buf
}

var _bsize Tsize_t

var _atexit_registered int32

const m_R_OK1 = 4

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
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
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
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

/*-
 * This file is in the public domain.
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1987, 1991 Regents of the University of California.
 * All rights reserved.
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
 *	@(#)endian.h	7.8 (Berkeley) 4/3/91
 */

func XXauGetAuthByAddr(tls *libc.TLS, family uint16, address_length uint16, address uintptr, number_length uint16, number uintptr, name_length uint16, name uintptr) (r uintptr) {
	var auth_file, auth_name, entry uintptr
	_, _, _ = auth_file, auth_name, entry
	auth_name = XXauFileName(tls)
	if !(auth_name != 0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xaccess(tls, auth_name, int32(m_R_OK1)) != 0 { /* checks REAL id */
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
const m_SCHED_FIFO = 1
const m_SCHED_OTHER = 2
const m_SCHED_RR = 3
const m_XMUTEX_INITIALIZER = "PTHREAD_MUTEX_INITIALIZER"
const m__NCPUBITS = "_BITSET_BITS"
const m_pthread_attr_default = "NULL"
const m_pthread_condattr_default = "NULL"
const m_pthread_mutexattr_default = "NULL"
const m_xfree = "free"
const m_xmalloc = "malloc"
const m_xthread_self = "pthread_self"

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

/* aids understood by some debuggers */

func XXauGetBestAuthByAddr(tls *libc.TLS, family uint16, address_length uint16, address uintptr, number_length uint16, number uintptr, types_length int32, types uintptr, type_lengths uintptr) (r uintptr) {
	var auth_file, auth_name, best, entry uintptr
	var best_type, type1 int32
	_, _, _, _, _, _ = auth_file, auth_name, best, best_type, entry, type1
	auth_name = XXauFileName(tls)
	if !(auth_name != 0) {
		return libc.UintptrFromInt32(0)
	}
	if libc.Xaccess(tls, auth_name, int32(m_R_OK1)) != 0 { /* checks REAL id */
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
const m_O_CREAT1 = 512
const m_O_EXCL1 = 2048
const m_O_WRONLY1 = 1
const m_R_OK2 = 0x04
const m_SF_APPEND = 0x00040000
const m_SF_ARCHIVED = 0x00010000
const m_SF_IMMUTABLE = 0x00020000
const m_SF_NOUNLINK = 0x00100000
const m_SF_SETTABLE = 0xffff0000
const m_SF_SNAPSHOT = 0x00200000
const m_S_BLKSIZE = 512
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
const m_S_ISTXT = 0001000
const m_S_ISUID = 0004000
const m_S_ISVTX = 0001000
const m_S_IWGRP = 0000020
const m_S_IWOTH = 0000002
const m_S_IWRITE = "S_IWUSR"
const m_S_IWUSR = 0000200
const m_S_IXGRP = 0000010
const m_S_IXOTH = 0000001
const m_S_IXUSR = 0000100
const m_Time_t = "time_t"
const m_UF_APPEND = 0x00000004
const m_UF_ARCHIVE = 0x00000800
const m_UF_HIDDEN = 0x00008000
const m_UF_IMMUTABLE = 0x00000002
const m_UF_NODUMP = 0x00000001
const m_UF_NOUNLINK = 0x00000010
const m_UF_OFFLINE = 0x00000200
const m_UF_OPAQUE = 0x00000008
const m_UF_READONLY = 0x00001000
const m_UF_REPARSE = 0x00000400
const m_UF_SETTABLE = 0x0000ffff
const m_UF_SPARSE = 0x00000100
const m_UF_SYSTEM = 0x00000080
const m_st_atimespec = "st_atim"
const m_st_birthtimespec = "st_birthtim"
const m_st_ctimespec = "st_ctim"
const m_st_mtimespec = "st_mtim"

type Tstat = struct {
	Fst_dev      Tdev_t
	Fst_ino      Tino_t
	Fst_nlink    Tnlink_t
	Fst_mode     Tmode_t
	Fst_padding0 t__int16_t
	Fst_uid      Tuid_t
	Fst_gid      Tgid_t
	Fst_padding1 t__int32_t
	Fst_rdev     Tdev_t
	Fst_atim     Ttimespec
	Fst_mtim     Ttimespec
	Fst_ctim     Ttimespec
	Fst_birthtim Ttimespec
	Fst_size     Toff_t
	Fst_blocks   Tblkcnt_t
	Fst_blksize  Tblksize_t
	Fst_flags    Tfflags_t
	Fst_gen      t__uint64_t
	Fst_spare    [10]t__uint64_t
}

/* 11 was EAGAIN */

/* math software */

/* non-blocking and interrupt i/o */

/* ipc/network software -- argument errors */

/* ipc/network software -- operational errors */

/* should be rearranged */

/* quotas & mush */

/* Network File System */

/* ISO/IEC 9899:2011 K.3.2.2 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
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
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
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
 *	@(#)unistd.h	8.12 (Berkeley) 4/27/95
 */

func XXauLockAuth(tls *libc.TLS, file_name uintptr, retries int32, timeout int32, dead int64) (r int32) {
	bp := tls.Alloc(2304)
	defer tls.Free(2304)
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
	libc.X__builtin_snprintf(tls, bp, uint64(1025), __ccgo_ts+37, libc.VaList(bp+2288, file_name))
	libc.X__builtin_snprintf(tls, bp+1025, uint64(1025), __ccgo_ts+42, libc.VaList(bp+2288, file_name))
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
			creat_fd = libc.Xopen(tls, bp, libc.Int32FromInt32(m_O_WRONLY1)|libc.Int32FromInt32(m_O_CREAT1)|libc.Int32FromInt32(m_O_EXCL1), libc.VaList(bp+2288, int32(0600)))
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
			libc.Xexplicit_bzero(tls, data, uint64(*(*uint16)(unsafe.Pointer(bp))))
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
		libc.Xexplicit_bzero(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fdata, uint64((*(*TXauth)(unsafe.Pointer(bp))).Fdata_length))
		libc.Xfree(tls, (*(*TXauth)(unsafe.Pointer(bp))).Fdata)
	}
	return libc.UintptrFromInt32(0)
}

const m_O_CREAT2 = 0x0200
const m_O_EXCL2 = 0x0800
const m_O_WRONLY2 = 0x0001

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
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
 *	@(#)time.h	8.3 (Berkeley) 1/21/94
 */

/*
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

/*-
 * This file is in the public domain.
 */

/*-
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * Copyright (c) 1987, 1991 Regents of the University of California.
 * All rights reserved.
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
 *	@(#)endian.h	7.8 (Berkeley) 4/3/91
 */

func XXauUnlockAuth(tls *libc.TLS, file_name uintptr) (r int32) {
	bp := tls.Alloc(2080)
	defer tls.Free(2080)
	var _ /* creat_name at bp+0 */ [1025]int8
	var _ /* link_name at bp+1025 */ [1025]int8
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
