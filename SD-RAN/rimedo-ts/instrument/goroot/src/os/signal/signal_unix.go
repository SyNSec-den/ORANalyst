// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || windows

//line /usr/local/go/src/os/signal/signal_unix.go:7
package signal

//line /usr/local/go/src/os/signal/signal_unix.go:7
import (
//line /usr/local/go/src/os/signal/signal_unix.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/os/signal/signal_unix.go:7
)
//line /usr/local/go/src/os/signal/signal_unix.go:7
import (
//line /usr/local/go/src/os/signal/signal_unix.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/os/signal/signal_unix.go:7
)

import (
	"os"
	"syscall"
)

// Defined by the runtime package.
func signal_disable(uint32)
func signal_enable(uint32)
func signal_ignore(uint32)
func signal_ignored(uint32) bool
func signal_recv() uint32

func loop() {
//line /usr/local/go/src/os/signal/signal_unix.go:21
	_go_fuzz_dep_.CoverTab[197204]++
							for {
//line /usr/local/go/src/os/signal/signal_unix.go:22
		_go_fuzz_dep_.CoverTab[197205]++
								process(syscall.Signal(signal_recv()))
//line /usr/local/go/src/os/signal/signal_unix.go:23
		// _ = "end of CoverTab[197205]"
	}
//line /usr/local/go/src/os/signal/signal_unix.go:24
	// _ = "end of CoverTab[197204]"
}

func init() {
	watchSignalLoop = loop
}

const (
	numSig = 65	// max across all systems
)

func signum(sig os.Signal) int {
//line /usr/local/go/src/os/signal/signal_unix.go:35
	_go_fuzz_dep_.CoverTab[197206]++
							switch sig := sig.(type) {
	case syscall.Signal:
//line /usr/local/go/src/os/signal/signal_unix.go:37
		_go_fuzz_dep_.CoverTab[197207]++
								i := int(sig)
								if i < 0 || func() bool {
//line /usr/local/go/src/os/signal/signal_unix.go:39
			_go_fuzz_dep_.CoverTab[197210]++
//line /usr/local/go/src/os/signal/signal_unix.go:39
			return i >= numSig
//line /usr/local/go/src/os/signal/signal_unix.go:39
			// _ = "end of CoverTab[197210]"
//line /usr/local/go/src/os/signal/signal_unix.go:39
		}() {
//line /usr/local/go/src/os/signal/signal_unix.go:39
			_go_fuzz_dep_.CoverTab[197211]++
									return -1
//line /usr/local/go/src/os/signal/signal_unix.go:40
			// _ = "end of CoverTab[197211]"
		} else {
//line /usr/local/go/src/os/signal/signal_unix.go:41
			_go_fuzz_dep_.CoverTab[197212]++
//line /usr/local/go/src/os/signal/signal_unix.go:41
			// _ = "end of CoverTab[197212]"
//line /usr/local/go/src/os/signal/signal_unix.go:41
		}
//line /usr/local/go/src/os/signal/signal_unix.go:41
		// _ = "end of CoverTab[197207]"
//line /usr/local/go/src/os/signal/signal_unix.go:41
		_go_fuzz_dep_.CoverTab[197208]++
								return i
//line /usr/local/go/src/os/signal/signal_unix.go:42
		// _ = "end of CoverTab[197208]"
	default:
//line /usr/local/go/src/os/signal/signal_unix.go:43
		_go_fuzz_dep_.CoverTab[197209]++
								return -1
//line /usr/local/go/src/os/signal/signal_unix.go:44
		// _ = "end of CoverTab[197209]"
	}
//line /usr/local/go/src/os/signal/signal_unix.go:45
	// _ = "end of CoverTab[197206]"
}

func enableSignal(sig int) {
//line /usr/local/go/src/os/signal/signal_unix.go:48
	_go_fuzz_dep_.CoverTab[197213]++
							signal_enable(uint32(sig))
//line /usr/local/go/src/os/signal/signal_unix.go:49
	// _ = "end of CoverTab[197213]"
}

func disableSignal(sig int) {
//line /usr/local/go/src/os/signal/signal_unix.go:52
	_go_fuzz_dep_.CoverTab[197214]++
							signal_disable(uint32(sig))
//line /usr/local/go/src/os/signal/signal_unix.go:53
	// _ = "end of CoverTab[197214]"
}

func ignoreSignal(sig int) {
//line /usr/local/go/src/os/signal/signal_unix.go:56
	_go_fuzz_dep_.CoverTab[197215]++
							signal_ignore(uint32(sig))
//line /usr/local/go/src/os/signal/signal_unix.go:57
	// _ = "end of CoverTab[197215]"
}

func signalIgnored(sig int) bool {
//line /usr/local/go/src/os/signal/signal_unix.go:60
	_go_fuzz_dep_.CoverTab[197216]++
							return signal_ignored(uint32(sig))
//line /usr/local/go/src/os/signal/signal_unix.go:61
	// _ = "end of CoverTab[197216]"
}

//line /usr/local/go/src/os/signal/signal_unix.go:62
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/os/signal/signal_unix.go:62
var _ = _go_fuzz_dep_.CoverTab
