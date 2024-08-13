// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || freebsd || dragonfly || solaris

//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
package rand

//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
import (
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
)
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
import (
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:7
)

import (
	"internal/syscall/unix"
	"runtime"
	"syscall"
)

func init() {
	var maxGetRandomRead int
	switch runtime.GOOS {
	case "linux", "android":

//line /usr/local/go/src/crypto/rand/rand_getrandom.go:23
		maxGetRandomRead = (1 << 25) - 1
	case "freebsd", "dragonfly", "solaris", "illumos":
		maxGetRandomRead = 1 << 8
	default:
		panic("no maximum specified for GetRandom")
	}
	altGetRandom = batched(getRandom, maxGetRandomRead)
}

// If the kernel is too old to support the getrandom syscall(),
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:32
// unix.GetRandom will immediately return ENOSYS and we will then fall back to
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:32
// reading from /dev/urandom in rand_unix.go. unix.GetRandom caches the ENOSYS
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:32
// result so we only suffer the syscall overhead once in this case.
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:32
// If the kernel supports the getrandom() syscall, unix.GetRandom will block
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:32
// until the kernel has sufficient randomness (as we don't use GRND_NONBLOCK).
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:32
// In this case, unix.GetRandom will not return an error.
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:39
func getRandom(p []byte) error {
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:39
	_go_fuzz_dep_.CoverTab[9332]++
								n, err := unix.GetRandom(p, 0)
								if err != nil {
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:41
		_go_fuzz_dep_.CoverTab[9335]++
									return err
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:42
		// _ = "end of CoverTab[9335]"
	} else {
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:43
		_go_fuzz_dep_.CoverTab[9336]++
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:43
		// _ = "end of CoverTab[9336]"
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:43
	}
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:43
	// _ = "end of CoverTab[9332]"
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:43
	_go_fuzz_dep_.CoverTab[9333]++
								if n != len(p) {
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:44
		_go_fuzz_dep_.CoverTab[9337]++
									return syscall.EIO
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:45
		// _ = "end of CoverTab[9337]"
	} else {
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:46
		_go_fuzz_dep_.CoverTab[9338]++
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:46
		// _ = "end of CoverTab[9338]"
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:46
	}
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:46
	// _ = "end of CoverTab[9333]"
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:46
	_go_fuzz_dep_.CoverTab[9334]++
								return nil
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:47
	// _ = "end of CoverTab[9334]"
}

//line /usr/local/go/src/crypto/rand/rand_getrandom.go:48
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/crypto/rand/rand_getrandom.go:48
var _ = _go_fuzz_dep_.CoverTab
