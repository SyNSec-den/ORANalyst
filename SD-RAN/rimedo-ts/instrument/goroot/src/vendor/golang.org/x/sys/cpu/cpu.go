// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:5
// Package cpu implements processor feature detection for
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:5
// various CPU architectures.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
package cpu

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:7
)

import (
	"os"
	"strings"
)

// Initialized reports whether the CPU features were initialized.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:14
//
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:14
// For some GOOS/GOARCH combinations initialization of the CPU features depends
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:14
// on reading an operating specific file, e.g. /proc/self/auxv on linux/arm
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:14
// Initialized will report false if reading the file fails.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:19
var Initialized bool

// CacheLinePad is used to pad structs to avoid false sharing.
type CacheLinePad struct{ _ [cacheLineSize]byte }

// X86 contains the supported CPU features of the
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:24
// current X86/AMD64 platform. If the current platform
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:24
// is not X86/AMD64 then all feature flags are false.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:24
//
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:24
// X86 is padded to avoid false sharing. Further the HasAVX
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:24
// and HasAVX2 are only set if the OS supports XMM and YMM
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:24
// registers in addition to the CPUID feature bit being set.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:31
var X86 struct {
	_			CacheLinePad
	HasAES			bool	// AES hardware implementation (AES NI)
	HasADX			bool	// Multi-precision add-carry instruction extensions
	HasAVX			bool	// Advanced vector extension
	HasAVX2			bool	// Advanced vector extension 2
	HasAVX512		bool	// Advanced vector extension 512
	HasAVX512F		bool	// Advanced vector extension 512 Foundation Instructions
	HasAVX512CD		bool	// Advanced vector extension 512 Conflict Detection Instructions
	HasAVX512ER		bool	// Advanced vector extension 512 Exponential and Reciprocal Instructions
	HasAVX512PF		bool	// Advanced vector extension 512 Prefetch Instructions Instructions
	HasAVX512VL		bool	// Advanced vector extension 512 Vector Length Extensions
	HasAVX512BW		bool	// Advanced vector extension 512 Byte and Word Instructions
	HasAVX512DQ		bool	// Advanced vector extension 512 Doubleword and Quadword Instructions
	HasAVX512IFMA		bool	// Advanced vector extension 512 Integer Fused Multiply Add
	HasAVX512VBMI		bool	// Advanced vector extension 512 Vector Byte Manipulation Instructions
	HasAVX5124VNNIW		bool	// Advanced vector extension 512 Vector Neural Network Instructions Word variable precision
	HasAVX5124FMAPS		bool	// Advanced vector extension 512 Fused Multiply Accumulation Packed Single precision
	HasAVX512VPOPCNTDQ	bool	// Advanced vector extension 512 Double and quad word population count instructions
	HasAVX512VPCLMULQDQ	bool	// Advanced vector extension 512 Vector carry-less multiply operations
	HasAVX512VNNI		bool	// Advanced vector extension 512 Vector Neural Network Instructions
	HasAVX512GFNI		bool	// Advanced vector extension 512 Galois field New Instructions
	HasAVX512VAES		bool	// Advanced vector extension 512 Vector AES instructions
	HasAVX512VBMI2		bool	// Advanced vector extension 512 Vector Byte Manipulation Instructions 2
	HasAVX512BITALG		bool	// Advanced vector extension 512 Bit Algorithms
	HasAVX512BF16		bool	// Advanced vector extension 512 BFloat16 Instructions
	HasBMI1			bool	// Bit manipulation instruction set 1
	HasBMI2			bool	// Bit manipulation instruction set 2
	HasCX16			bool	// Compare and exchange 16 Bytes
	HasERMS			bool	// Enhanced REP for MOVSB and STOSB
	HasFMA			bool	// Fused-multiply-add instructions
	HasOSXSAVE		bool	// OS supports XSAVE/XRESTOR for saving/restoring XMM registers.
	HasPCLMULQDQ		bool	// PCLMULQDQ instruction - most often used for AES-GCM
	HasPOPCNT		bool	// Hamming weight instruction POPCNT.
	HasRDRAND		bool	// RDRAND instruction (on-chip random number generator)
	HasRDSEED		bool	// RDSEED instruction (on-chip random number generator)
	HasSSE2			bool	// Streaming SIMD extension 2 (always available on amd64)
	HasSSE3			bool	// Streaming SIMD extension 3
	HasSSSE3		bool	// Supplemental streaming SIMD extension 3
	HasSSE41		bool	// Streaming SIMD extension 4 and 4.1
	HasSSE42		bool	// Streaming SIMD extension 4 and 4.2
	_			CacheLinePad
}

// ARM64 contains the supported CPU features of the
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:75
// current ARMv8(aarch64) platform. If the current platform
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:75
// is not arm64 then all feature flags are false.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:78
var ARM64 struct {
	_		CacheLinePad
	HasFP		bool	// Floating-point instruction set (always available)
	HasASIMD	bool	// Advanced SIMD (always available)
	HasEVTSTRM	bool	// Event stream support
	HasAES		bool	// AES hardware implementation
	HasPMULL	bool	// Polynomial multiplication instruction set
	HasSHA1		bool	// SHA1 hardware implementation
	HasSHA2		bool	// SHA2 hardware implementation
	HasCRC32	bool	// CRC32 hardware implementation
	HasATOMICS	bool	// Atomic memory operation instruction set
	HasFPHP		bool	// Half precision floating-point instruction set
	HasASIMDHP	bool	// Advanced SIMD half precision instruction set
	HasCPUID	bool	// CPUID identification scheme registers
	HasASIMDRDM	bool	// Rounding double multiply add/subtract instruction set
	HasJSCVT	bool	// Javascript conversion from floating-point to integer
	HasFCMA		bool	// Floating-point multiplication and addition of complex numbers
	HasLRCPC	bool	// Release Consistent processor consistent support
	HasDCPOP	bool	// Persistent memory support
	HasSHA3		bool	// SHA3 hardware implementation
	HasSM3		bool	// SM3 hardware implementation
	HasSM4		bool	// SM4 hardware implementation
	HasASIMDDP	bool	// Advanced SIMD double precision instruction set
	HasSHA512	bool	// SHA512 hardware implementation
	HasSVE		bool	// Scalable Vector Extensions
	HasASIMDFHM	bool	// Advanced SIMD multiplication FP16 to FP32
	_		CacheLinePad
}

// ARM contains the supported CPU features of the current ARM (32-bit) platform.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:107
// All feature flags are false if:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:107
//  1. the current platform is not arm, or
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:107
//  2. the current operating system is not Linux.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:111
var ARM struct {
	_		CacheLinePad
	HasSWP		bool	// SWP instruction support
	HasHALF		bool	// Half-word load and store support
	HasTHUMB	bool	// ARM Thumb instruction set
	Has26BIT	bool	// Address space limited to 26-bits
	HasFASTMUL	bool	// 32-bit operand, 64-bit result multiplication support
	HasFPA		bool	// Floating point arithmetic support
	HasVFP		bool	// Vector floating point support
	HasEDSP		bool	// DSP Extensions support
	HasJAVA		bool	// Java instruction set
	HasIWMMXT	bool	// Intel Wireless MMX technology support
	HasCRUNCH	bool	// MaverickCrunch context switching and handling
	HasTHUMBEE	bool	// Thumb EE instruction set
	HasNEON		bool	// NEON instruction set
	HasVFPv3	bool	// Vector floating point version 3 support
	HasVFPv3D16	bool	// Vector floating point version 3 D8-D15
	HasTLS		bool	// Thread local storage support
	HasVFPv4	bool	// Vector floating point version 4 support
	HasIDIVA	bool	// Integer divide instruction support in ARM mode
	HasIDIVT	bool	// Integer divide instruction support in Thumb mode
	HasVFPD32	bool	// Vector floating point version 3 D15-D31
	HasLPAE		bool	// Large Physical Address Extensions
	HasEVTSTRM	bool	// Event stream support
	HasAES		bool	// AES hardware implementation
	HasPMULL	bool	// Polynomial multiplication instruction set
	HasSHA1		bool	// SHA1 hardware implementation
	HasSHA2		bool	// SHA2 hardware implementation
	HasCRC32	bool	// CRC32 hardware implementation
	_		CacheLinePad
}

// MIPS64X contains the supported CPU features of the current mips64/mips64le
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:143
// platforms. If the current platform is not mips64/mips64le or the current
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:143
// operating system is not Linux then all feature flags are false.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:146
var MIPS64X struct {
	_	CacheLinePad
	HasMSA	bool	// MIPS SIMD architecture
	_	CacheLinePad
}

// PPC64 contains the supported CPU features of the current ppc64/ppc64le platforms.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:152
// If the current platform is not ppc64/ppc64le then all feature flags are false.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:152
//
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:152
// For ppc64/ppc64le, it is safe to check only for ISA level starting on ISA v3.00,
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:152
// since there are no optional categories. There are some exceptions that also
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:152
// require kernel support to work (DARN, SCV), so there are feature bits for
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:152
// those as well. The struct is padded to avoid false sharing.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:159
var PPC64 struct {
	_		CacheLinePad
	HasDARN		bool	// Hardware random number generator (requires kernel enablement)
	HasSCV		bool	// Syscall vectored (requires kernel enablement)
	IsPOWER8	bool	// ISA v2.07 (POWER8)
	IsPOWER9	bool	// ISA v3.00 (POWER9), implies IsPOWER8
	_		CacheLinePad
}

// S390X contains the supported CPU features of the current IBM Z
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:168
// (s390x) platform. If the current platform is not IBM Z then all
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:168
// feature flags are false.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:168
//
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:168
// S390X is padded to avoid false sharing. Further HasVX is only set
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:168
// if the OS supports vector registers in addition to the STFLE
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:168
// feature bit being set.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:175
var S390X struct {
	_		CacheLinePad
	HasZARCH	bool	// z/Architecture mode is active [mandatory]
	HasSTFLE	bool	// store facility list extended
	HasLDISP	bool	// long (20-bit) displacements
	HasEIMM		bool	// 32-bit immediates
	HasDFP		bool	// decimal floating point
	HasETF3EH	bool	// ETF-3 enhanced
	HasMSA		bool	// message security assist (CPACF)
	HasAES		bool	// KM-AES{128,192,256} functions
	HasAESCBC	bool	// KMC-AES{128,192,256} functions
	HasAESCTR	bool	// KMCTR-AES{128,192,256} functions
	HasAESGCM	bool	// KMA-GCM-AES{128,192,256} functions
	HasGHASH	bool	// KIMD-GHASH function
	HasSHA1		bool	// K{I,L}MD-SHA-1 functions
	HasSHA256	bool	// K{I,L}MD-SHA-256 functions
	HasSHA512	bool	// K{I,L}MD-SHA-512 functions
	HasSHA3		bool	// K{I,L}MD-SHA3-{224,256,384,512} and K{I,L}MD-SHAKE-{128,256} functions
	HasVX		bool	// vector facility
	HasVXE		bool	// vector-enhancements facility 1
	_		CacheLinePad
}

func init() {
	archInit()
	initOptions()
	processOptions()
}

// options contains the cpu debug options that can be used in GODEBUG.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:204
// Options are arch dependent and are added by the arch specific initOptions functions.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:204
// Features that are mandatory for the specific GOARCH should have the Required field set
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:204
// (e.g. SSE2 on amd64).
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:208
var options []option

// Option names should be lower case. e.g. avx instead of AVX.
type option struct {
	Name		string
	Feature		*bool
	Specified	bool	// whether feature value was specified in GODEBUG
	Enable		bool	// whether feature should be enabled
	Required	bool	// whether feature is mandatory and can not be disabled
}

func processOptions() {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:219
	_go_fuzz_dep_.CoverTab[20830]++
									env := os.Getenv("GODEBUG")
field:
	for env != "" {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:222
		_go_fuzz_dep_.CoverTab[20832]++
										field := ""
										i := strings.IndexByte(env, ',')
										if i < 0 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:225
			_go_fuzz_dep_.CoverTab[20839]++
											field, env = env, ""
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:226
			// _ = "end of CoverTab[20839]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:227
			_go_fuzz_dep_.CoverTab[20840]++
											field, env = env[:i], env[i+1:]
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:228
			// _ = "end of CoverTab[20840]"
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:229
		// _ = "end of CoverTab[20832]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:229
		_go_fuzz_dep_.CoverTab[20833]++
										if len(field) < 4 || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:230
			_go_fuzz_dep_.CoverTab[20841]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:230
			return field[:4] != "cpu."
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:230
			// _ = "end of CoverTab[20841]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:230
		}() {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:230
			_go_fuzz_dep_.CoverTab[20842]++
											continue
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:231
			// _ = "end of CoverTab[20842]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:232
			_go_fuzz_dep_.CoverTab[20843]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:232
			// _ = "end of CoverTab[20843]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:232
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:232
		// _ = "end of CoverTab[20833]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:232
		_go_fuzz_dep_.CoverTab[20834]++
										i = strings.IndexByte(field, '=')
										if i < 0 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:234
			_go_fuzz_dep_.CoverTab[20844]++
											print("GODEBUG sys/cpu: no value specified for \"", field, "\"\n")
											continue
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:236
			// _ = "end of CoverTab[20844]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:237
			_go_fuzz_dep_.CoverTab[20845]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:237
			// _ = "end of CoverTab[20845]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:237
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:237
		// _ = "end of CoverTab[20834]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:237
		_go_fuzz_dep_.CoverTab[20835]++
										key, value := field[4:i], field[i+1:]

										var enable bool
										switch value {
		case "on":
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:242
			_go_fuzz_dep_.CoverTab[20846]++
											enable = true
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:243
			// _ = "end of CoverTab[20846]"
		case "off":
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:244
			_go_fuzz_dep_.CoverTab[20847]++
											enable = false
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:245
			// _ = "end of CoverTab[20847]"
		default:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:246
			_go_fuzz_dep_.CoverTab[20848]++
											print("GODEBUG sys/cpu: value \"", value, "\" not supported for cpu option \"", key, "\"\n")
											continue field
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:248
			// _ = "end of CoverTab[20848]"
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:249
		// _ = "end of CoverTab[20835]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:249
		_go_fuzz_dep_.CoverTab[20836]++

										if key == "all" {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:251
			_go_fuzz_dep_.CoverTab[20849]++
											for i := range options {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:252
				_go_fuzz_dep_.CoverTab[20851]++
												options[i].Specified = true
												options[i].Enable = enable || func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:254
					_go_fuzz_dep_.CoverTab[20852]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:254
					return options[i].Required
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:254
					// _ = "end of CoverTab[20852]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:254
				}()
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:254
				// _ = "end of CoverTab[20851]"
			}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:255
			// _ = "end of CoverTab[20849]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:255
			_go_fuzz_dep_.CoverTab[20850]++
											continue field
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:256
			// _ = "end of CoverTab[20850]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:257
			_go_fuzz_dep_.CoverTab[20853]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:257
			// _ = "end of CoverTab[20853]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:257
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:257
		// _ = "end of CoverTab[20836]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:257
		_go_fuzz_dep_.CoverTab[20837]++

										for i := range options {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:259
			_go_fuzz_dep_.CoverTab[20854]++
											if options[i].Name == key {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:260
				_go_fuzz_dep_.CoverTab[20855]++
												options[i].Specified = true
												options[i].Enable = enable
												continue field
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:263
				// _ = "end of CoverTab[20855]"
			} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:264
				_go_fuzz_dep_.CoverTab[20856]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:264
				// _ = "end of CoverTab[20856]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:264
			}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:264
			// _ = "end of CoverTab[20854]"
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:265
		// _ = "end of CoverTab[20837]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:265
		_go_fuzz_dep_.CoverTab[20838]++

										print("GODEBUG sys/cpu: unknown cpu feature \"", key, "\"\n")
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:267
		// _ = "end of CoverTab[20838]"
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:268
	// _ = "end of CoverTab[20830]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:268
	_go_fuzz_dep_.CoverTab[20831]++

									for _, o := range options {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:270
		_go_fuzz_dep_.CoverTab[20857]++
										if !o.Specified {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:271
			_go_fuzz_dep_.CoverTab[20861]++
											continue
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:272
			// _ = "end of CoverTab[20861]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:273
			_go_fuzz_dep_.CoverTab[20862]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:273
			// _ = "end of CoverTab[20862]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:273
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:273
		// _ = "end of CoverTab[20857]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:273
		_go_fuzz_dep_.CoverTab[20858]++

										if o.Enable && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:275
			_go_fuzz_dep_.CoverTab[20863]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:275
			return !*o.Feature
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:275
			// _ = "end of CoverTab[20863]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:275
		}() {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:275
			_go_fuzz_dep_.CoverTab[20864]++
											print("GODEBUG sys/cpu: can not enable \"", o.Name, "\", missing CPU support\n")
											continue
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:277
			// _ = "end of CoverTab[20864]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:278
			_go_fuzz_dep_.CoverTab[20865]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:278
			// _ = "end of CoverTab[20865]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:278
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:278
		// _ = "end of CoverTab[20858]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:278
		_go_fuzz_dep_.CoverTab[20859]++

										if !o.Enable && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:280
			_go_fuzz_dep_.CoverTab[20866]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:280
			return o.Required
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:280
			// _ = "end of CoverTab[20866]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:280
		}() {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:280
			_go_fuzz_dep_.CoverTab[20867]++
											print("GODEBUG sys/cpu: can not disable \"", o.Name, "\", required CPU feature\n")
											continue
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:282
			// _ = "end of CoverTab[20867]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:283
			_go_fuzz_dep_.CoverTab[20868]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:283
			// _ = "end of CoverTab[20868]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:283
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:283
		// _ = "end of CoverTab[20859]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:283
		_go_fuzz_dep_.CoverTab[20860]++

										*o.Feature = o.Enable
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:285
		// _ = "end of CoverTab[20860]"
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:286
	// _ = "end of CoverTab[20831]"
}

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:287
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/cpu.go:287
var _ = _go_fuzz_dep_.CoverTab
