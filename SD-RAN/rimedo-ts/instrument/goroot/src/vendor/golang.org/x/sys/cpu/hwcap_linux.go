// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
package cpu

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:5
)

import (
	"io/ioutil"
)

const (
	_AT_HWCAP	= 16
	_AT_HWCAP2	= 26

	procAuxv	= "/proc/self/auxv"

	uintSize	= int(32 << (^uint(0) >> 63))
)

// For those platforms don't have a 'cpuid' equivalent we use HWCAP/HWCAP2
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:20
// These are initialized in cpu_$GOARCH.go
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:20
// and should not be changed after they are initialized.
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:23
var hwCap uint
var hwCap2 uint

func readHWCAP() error {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:26
	_go_fuzz_dep_.CoverTab[20893]++
									buf, err := ioutil.ReadFile(procAuxv)
									if err != nil {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:28
		_go_fuzz_dep_.CoverTab[20896]++

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:33
		return err
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:33
		// _ = "end of CoverTab[20896]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:34
		_go_fuzz_dep_.CoverTab[20897]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:34
		// _ = "end of CoverTab[20897]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:34
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:34
	// _ = "end of CoverTab[20893]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:34
	_go_fuzz_dep_.CoverTab[20894]++
									bo := hostByteOrder()
									for len(buf) >= 2*(uintSize/8) {
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:36
		_go_fuzz_dep_.CoverTab[20898]++
										var tag, val uint
										switch uintSize {
		case 32:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:39
			_go_fuzz_dep_.CoverTab[20900]++
											tag = uint(bo.Uint32(buf[0:]))
											val = uint(bo.Uint32(buf[4:]))
											buf = buf[8:]
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:42
			// _ = "end of CoverTab[20900]"
		case 64:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:43
			_go_fuzz_dep_.CoverTab[20901]++
											tag = uint(bo.Uint64(buf[0:]))
											val = uint(bo.Uint64(buf[8:]))
											buf = buf[16:]
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:46
			// _ = "end of CoverTab[20901]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:46
		default:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:46
			_go_fuzz_dep_.CoverTab[20902]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:46
			// _ = "end of CoverTab[20902]"
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:47
		// _ = "end of CoverTab[20898]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:47
		_go_fuzz_dep_.CoverTab[20899]++
										switch tag {
		case _AT_HWCAP:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:49
			_go_fuzz_dep_.CoverTab[20903]++
											hwCap = val
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:50
			// _ = "end of CoverTab[20903]"
		case _AT_HWCAP2:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:51
			_go_fuzz_dep_.CoverTab[20904]++
											hwCap2 = val
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:52
			// _ = "end of CoverTab[20904]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:52
		default:
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:52
			_go_fuzz_dep_.CoverTab[20905]++
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:52
			// _ = "end of CoverTab[20905]"
		}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:53
		// _ = "end of CoverTab[20899]"
	}
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:54
	// _ = "end of CoverTab[20894]"
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:54
	_go_fuzz_dep_.CoverTab[20895]++
									return nil
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:55
	// _ = "end of CoverTab[20895]"
}

//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:56
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/sys/cpu/hwcap_linux.go:56
var _ = _go_fuzz_dep_.CoverTab
