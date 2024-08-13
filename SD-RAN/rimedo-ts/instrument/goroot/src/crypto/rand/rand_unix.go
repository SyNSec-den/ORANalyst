// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

// Unix cryptographically secure pseudorandom number
// generator.

//line /usr/local/go/src/crypto/rand/rand_unix.go:10
package rand

//line /usr/local/go/src/crypto/rand/rand_unix.go:10
import (
//line /usr/local/go/src/crypto/rand/rand_unix.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rand/rand_unix.go:10
)
//line /usr/local/go/src/crypto/rand/rand_unix.go:10
import (
//line /usr/local/go/src/crypto/rand/rand_unix.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rand/rand_unix.go:10
)

import (
	"crypto/internal/boring"
	"errors"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

const urandomDevice = "/dev/urandom"

func init() {
	if boring.Enabled {
		Reader = boring.RandReader
		return
	}
	Reader = &reader{}
}

// A reader satisfies reads by reading from urandomDevice
type reader struct {
	f	io.Reader
	mu	sync.Mutex
	used	atomic.Uint32	// Atomic: 0 - never used, 1 - used, but f == nil, 2 - used, and f != nil
}

// altGetRandom if non-nil specifies an OS-specific function to get
//line /usr/local/go/src/crypto/rand/rand_unix.go:40
// urandom-style randomness.
//line /usr/local/go/src/crypto/rand/rand_unix.go:42
var altGetRandom func([]byte) (err error)

func warnBlocked() {
//line /usr/local/go/src/crypto/rand/rand_unix.go:44
	_go_fuzz_dep_.CoverTab[9339]++
							println("crypto/rand: blocked for 60 seconds waiting to read random data from the kernel")
//line /usr/local/go/src/crypto/rand/rand_unix.go:45
	// _ = "end of CoverTab[9339]"
}

func (r *reader) Read(b []byte) (n int, err error) {
//line /usr/local/go/src/crypto/rand/rand_unix.go:48
	_go_fuzz_dep_.CoverTab[9340]++
							boring.Unreachable()
							if r.used.CompareAndSwap(0, 1) {
//line /usr/local/go/src/crypto/rand/rand_unix.go:50
		_go_fuzz_dep_.CoverTab[9344]++

//line /usr/local/go/src/crypto/rand/rand_unix.go:53
		t := time.AfterFunc(time.Minute, warnBlocked)
								defer t.Stop()
//line /usr/local/go/src/crypto/rand/rand_unix.go:54
		// _ = "end of CoverTab[9344]"
	} else {
//line /usr/local/go/src/crypto/rand/rand_unix.go:55
		_go_fuzz_dep_.CoverTab[9345]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:55
		// _ = "end of CoverTab[9345]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:55
	}
//line /usr/local/go/src/crypto/rand/rand_unix.go:55
	// _ = "end of CoverTab[9340]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:55
	_go_fuzz_dep_.CoverTab[9341]++
							if altGetRandom != nil && func() bool {
//line /usr/local/go/src/crypto/rand/rand_unix.go:56
		_go_fuzz_dep_.CoverTab[9346]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:56
		return altGetRandom(b) == nil
//line /usr/local/go/src/crypto/rand/rand_unix.go:56
		// _ = "end of CoverTab[9346]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:56
	}() {
//line /usr/local/go/src/crypto/rand/rand_unix.go:56
		_go_fuzz_dep_.CoverTab[9347]++
								return len(b), nil
//line /usr/local/go/src/crypto/rand/rand_unix.go:57
		// _ = "end of CoverTab[9347]"
	} else {
//line /usr/local/go/src/crypto/rand/rand_unix.go:58
		_go_fuzz_dep_.CoverTab[9348]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:58
		// _ = "end of CoverTab[9348]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:58
	}
//line /usr/local/go/src/crypto/rand/rand_unix.go:58
	// _ = "end of CoverTab[9341]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:58
	_go_fuzz_dep_.CoverTab[9342]++
							if r.used.Load() != 2 {
//line /usr/local/go/src/crypto/rand/rand_unix.go:59
		_go_fuzz_dep_.CoverTab[9349]++
								r.mu.Lock()
								if r.used.Load() != 2 {
//line /usr/local/go/src/crypto/rand/rand_unix.go:61
			_go_fuzz_dep_.CoverTab[9351]++
									f, err := os.Open(urandomDevice)
									if err != nil {
//line /usr/local/go/src/crypto/rand/rand_unix.go:63
				_go_fuzz_dep_.CoverTab[9353]++
										r.mu.Unlock()
										return 0, err
//line /usr/local/go/src/crypto/rand/rand_unix.go:65
				// _ = "end of CoverTab[9353]"
			} else {
//line /usr/local/go/src/crypto/rand/rand_unix.go:66
				_go_fuzz_dep_.CoverTab[9354]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:66
				// _ = "end of CoverTab[9354]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:66
			}
//line /usr/local/go/src/crypto/rand/rand_unix.go:66
			// _ = "end of CoverTab[9351]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:66
			_go_fuzz_dep_.CoverTab[9352]++
									r.f = hideAgainReader{f}
									r.used.Store(2)
//line /usr/local/go/src/crypto/rand/rand_unix.go:68
			// _ = "end of CoverTab[9352]"
		} else {
//line /usr/local/go/src/crypto/rand/rand_unix.go:69
			_go_fuzz_dep_.CoverTab[9355]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:69
			// _ = "end of CoverTab[9355]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:69
		}
//line /usr/local/go/src/crypto/rand/rand_unix.go:69
		// _ = "end of CoverTab[9349]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:69
		_go_fuzz_dep_.CoverTab[9350]++
								r.mu.Unlock()
//line /usr/local/go/src/crypto/rand/rand_unix.go:70
		// _ = "end of CoverTab[9350]"
	} else {
//line /usr/local/go/src/crypto/rand/rand_unix.go:71
		_go_fuzz_dep_.CoverTab[9356]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:71
		// _ = "end of CoverTab[9356]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:71
	}
//line /usr/local/go/src/crypto/rand/rand_unix.go:71
	// _ = "end of CoverTab[9342]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:71
	_go_fuzz_dep_.CoverTab[9343]++
							return io.ReadFull(r.f, b)
//line /usr/local/go/src/crypto/rand/rand_unix.go:72
	// _ = "end of CoverTab[9343]"
}

// hideAgainReader masks EAGAIN reads from /dev/urandom.
//line /usr/local/go/src/crypto/rand/rand_unix.go:75
// See golang.org/issue/9205
//line /usr/local/go/src/crypto/rand/rand_unix.go:77
type hideAgainReader struct {
	r io.Reader
}

func (hr hideAgainReader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/crypto/rand/rand_unix.go:81
	_go_fuzz_dep_.CoverTab[9357]++
							n, err = hr.r.Read(p)
							if errors.Is(err, syscall.EAGAIN) {
//line /usr/local/go/src/crypto/rand/rand_unix.go:83
		_go_fuzz_dep_.CoverTab[9359]++
								err = nil
//line /usr/local/go/src/crypto/rand/rand_unix.go:84
		// _ = "end of CoverTab[9359]"
	} else {
//line /usr/local/go/src/crypto/rand/rand_unix.go:85
		_go_fuzz_dep_.CoverTab[9360]++
//line /usr/local/go/src/crypto/rand/rand_unix.go:85
		// _ = "end of CoverTab[9360]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:85
	}
//line /usr/local/go/src/crypto/rand/rand_unix.go:85
	// _ = "end of CoverTab[9357]"
//line /usr/local/go/src/crypto/rand/rand_unix.go:85
	_go_fuzz_dep_.CoverTab[9358]++
							return
//line /usr/local/go/src/crypto/rand/rand_unix.go:86
	// _ = "end of CoverTab[9358]"
}

//line /usr/local/go/src/crypto/rand/rand_unix.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rand/rand_unix.go:87
var _ = _go_fuzz_dep_.CoverTab
