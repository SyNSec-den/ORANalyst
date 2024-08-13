// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
package cpu

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:5
)

import (
	"runtime"
)

// byteOrder is a subset of encoding/binary.ByteOrder.
type byteOrder interface {
	Uint32([]byte) uint32
	Uint64([]byte) uint64
}

type littleEndian struct{}
type bigEndian struct{}

func (littleEndian) Uint32(b []byte) uint32 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:20
	_go_fuzz_dep_.CoverTab[20821]++
									_ = b[3]
									return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:22
	// _ = "end of CoverTab[20821]"
}

func (littleEndian) Uint64(b []byte) uint64 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:25
	_go_fuzz_dep_.CoverTab[20822]++
									_ = b[7]
									return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:28
	// _ = "end of CoverTab[20822]"
}

func (bigEndian) Uint32(b []byte) uint32 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:31
	_go_fuzz_dep_.CoverTab[20823]++
									_ = b[3]
									return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:33
	// _ = "end of CoverTab[20823]"
}

func (bigEndian) Uint64(b []byte) uint64 {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:36
	_go_fuzz_dep_.CoverTab[20824]++
									_ = b[7]
									return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:39
	// _ = "end of CoverTab[20824]"
}

// hostByteOrder returns littleEndian on little-endian machines and
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:42
// bigEndian on big-endian machines.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:44
func hostByteOrder() byteOrder {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:44
	_go_fuzz_dep_.CoverTab[20825]++
									switch runtime.GOARCH {
	case "386", "amd64", "amd64p32",
		"alpha",
		"arm", "arm64",
		"loong64",
		"mipsle", "mips64le", "mips64p32le",
		"nios2",
		"ppc64le",
		"riscv", "riscv64",
		"sh":
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:54
		_go_fuzz_dep_.CoverTab[20827]++
										return littleEndian{}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:55
		// _ = "end of CoverTab[20827]"
	case "armbe", "arm64be",
		"m68k",
		"mips", "mips64", "mips64p32",
		"ppc", "ppc64",
		"s390", "s390x",
		"shbe",
		"sparc", "sparc64":
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:62
		_go_fuzz_dep_.CoverTab[20828]++
										return bigEndian{}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:63
		// _ = "end of CoverTab[20828]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:63
	default:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:63
		_go_fuzz_dep_.CoverTab[20829]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:63
		// _ = "end of CoverTab[20829]"
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:64
	// _ = "end of CoverTab[20825]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:64
	_go_fuzz_dep_.CoverTab[20826]++
									panic("unknown architecture")
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:65
	// _ = "end of CoverTab[20826]"
}

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:66
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/byteorder.go:66
var _ = _go_fuzz_dep_.CoverTab
