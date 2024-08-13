// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:8
)

import (
	"bytes"
	"sort"
	"sync"
	"syscall"
	"unsafe"
)

var (
	Stdin	= 0
	Stdout	= 1
	Stderr	= 2
)

// Do the interface allocations only once for common
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:24
// Errno values.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:26
var (
	errEAGAIN	error	= syscall.EAGAIN
	errEINVAL	error	= syscall.EINVAL
	errENOENT	error	= syscall.ENOENT
)

var (
	signalNameMapOnce	sync.Once
	signalNameMap		map[string]syscall.Signal
)

// errnoErr returns common boxed Errno values, to prevent
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:37
// allocations at runtime.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:39
func errnoErr(e syscall.Errno) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:39
	_go_fuzz_dep_.CoverTab[46640]++
											switch e {
	case 0:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:41
		_go_fuzz_dep_.CoverTab[46642]++
												return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:42
		// _ = "end of CoverTab[46642]"
	case EAGAIN:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:43
		_go_fuzz_dep_.CoverTab[46643]++
												return errEAGAIN
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:44
		// _ = "end of CoverTab[46643]"
	case EINVAL:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:45
		_go_fuzz_dep_.CoverTab[46644]++
												return errEINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:46
		// _ = "end of CoverTab[46644]"
	case ENOENT:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:47
		_go_fuzz_dep_.CoverTab[46645]++
												return errENOENT
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:48
		// _ = "end of CoverTab[46645]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:48
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:48
		_go_fuzz_dep_.CoverTab[46646]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:48
		// _ = "end of CoverTab[46646]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:49
	// _ = "end of CoverTab[46640]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:49
	_go_fuzz_dep_.CoverTab[46641]++
											return e
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:50
	// _ = "end of CoverTab[46641]"
}

// ErrnoName returns the error name for error number e.
func ErrnoName(e syscall.Errno) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:54
	_go_fuzz_dep_.CoverTab[46647]++
											i := sort.Search(len(errorList), func(i int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:55
		_go_fuzz_dep_.CoverTab[46650]++
												return errorList[i].num >= e
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:56
		// _ = "end of CoverTab[46650]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:57
	// _ = "end of CoverTab[46647]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:57
	_go_fuzz_dep_.CoverTab[46648]++
											if i < len(errorList) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:58
		_go_fuzz_dep_.CoverTab[46651]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:58
		return errorList[i].num == e
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:58
		// _ = "end of CoverTab[46651]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:58
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:58
		_go_fuzz_dep_.CoverTab[46652]++
												return errorList[i].name
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:59
		// _ = "end of CoverTab[46652]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:60
		_go_fuzz_dep_.CoverTab[46653]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:60
		// _ = "end of CoverTab[46653]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:60
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:60
	// _ = "end of CoverTab[46648]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:60
	_go_fuzz_dep_.CoverTab[46649]++
											return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:61
	// _ = "end of CoverTab[46649]"
}

// SignalName returns the signal name for signal number s.
func SignalName(s syscall.Signal) string {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:65
	_go_fuzz_dep_.CoverTab[46654]++
											i := sort.Search(len(signalList), func(i int) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:66
		_go_fuzz_dep_.CoverTab[46657]++
												return signalList[i].num >= s
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:67
		// _ = "end of CoverTab[46657]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:68
	// _ = "end of CoverTab[46654]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:68
	_go_fuzz_dep_.CoverTab[46655]++
											if i < len(signalList) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:69
		_go_fuzz_dep_.CoverTab[46658]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:69
		return signalList[i].num == s
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:69
		// _ = "end of CoverTab[46658]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:69
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:69
		_go_fuzz_dep_.CoverTab[46659]++
												return signalList[i].name
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:70
		// _ = "end of CoverTab[46659]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:71
		_go_fuzz_dep_.CoverTab[46660]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:71
		// _ = "end of CoverTab[46660]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:71
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:71
	// _ = "end of CoverTab[46655]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:71
	_go_fuzz_dep_.CoverTab[46656]++
											return ""
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:72
	// _ = "end of CoverTab[46656]"
}

// SignalNum returns the syscall.Signal for signal named s,
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:75
// or 0 if a signal with such name is not found.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:75
// The signal name should start with "SIG".
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:78
func SignalNum(s string) syscall.Signal {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:78
	_go_fuzz_dep_.CoverTab[46661]++
											signalNameMapOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:79
		_go_fuzz_dep_.CoverTab[46663]++
												signalNameMap = make(map[string]syscall.Signal, len(signalList))
												for _, signal := range signalList {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:81
			_go_fuzz_dep_.CoverTab[46664]++
													signalNameMap[signal.name] = signal.num
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:82
			// _ = "end of CoverTab[46664]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:83
		// _ = "end of CoverTab[46663]"
	})
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:84
	// _ = "end of CoverTab[46661]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:84
	_go_fuzz_dep_.CoverTab[46662]++
											return signalNameMap[s]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:85
	// _ = "end of CoverTab[46662]"
}

// clen returns the index of the first NULL byte in n or len(n) if n contains no NULL byte.
func clen(n []byte) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:89
	_go_fuzz_dep_.CoverTab[46665]++
											i := bytes.IndexByte(n, 0)
											if i == -1 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:91
		_go_fuzz_dep_.CoverTab[46667]++
												i = len(n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:92
		// _ = "end of CoverTab[46667]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:93
		_go_fuzz_dep_.CoverTab[46668]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:93
		// _ = "end of CoverTab[46668]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:93
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:93
	// _ = "end of CoverTab[46665]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:93
	_go_fuzz_dep_.CoverTab[46666]++
											return i
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:94
	// _ = "end of CoverTab[46666]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:99
type mmapper struct {
	sync.Mutex
	active	map[*byte][]byte	// active mappings; key is last byte in mapping
	mmap	func(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error)
	munmap	func(addr uintptr, length uintptr) error
}

func (m *mmapper) Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:106
	_go_fuzz_dep_.CoverTab[46669]++
											if length <= 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:107
		_go_fuzz_dep_.CoverTab[46672]++
												return nil, EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:108
		// _ = "end of CoverTab[46672]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:109
		_go_fuzz_dep_.CoverTab[46673]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:109
		// _ = "end of CoverTab[46673]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:109
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:109
	// _ = "end of CoverTab[46669]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:109
	_go_fuzz_dep_.CoverTab[46670]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:112
	addr, errno := m.mmap(0, uintptr(length), prot, flags, fd, offset)
	if errno != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:113
		_go_fuzz_dep_.CoverTab[46674]++
												return nil, errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:114
		// _ = "end of CoverTab[46674]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:115
		_go_fuzz_dep_.CoverTab[46675]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:115
		// _ = "end of CoverTab[46675]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:115
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:115
	// _ = "end of CoverTab[46670]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:115
	_go_fuzz_dep_.CoverTab[46671]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:118
	b := unsafe.Slice((*byte)(unsafe.Pointer(addr)), length)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:121
	p := &b[cap(b)-1]
											m.Lock()
											defer m.Unlock()
											m.active[p] = b
											return b, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:125
	// _ = "end of CoverTab[46671]"
}

func (m *mmapper) Munmap(data []byte) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:128
	_go_fuzz_dep_.CoverTab[46676]++
											if len(data) == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:129
		_go_fuzz_dep_.CoverTab[46680]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:129
		return len(data) != cap(data)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:129
		// _ = "end of CoverTab[46680]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:129
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:129
		_go_fuzz_dep_.CoverTab[46681]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:130
		// _ = "end of CoverTab[46681]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:131
		_go_fuzz_dep_.CoverTab[46682]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:131
		// _ = "end of CoverTab[46682]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:131
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:131
	// _ = "end of CoverTab[46676]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:131
	_go_fuzz_dep_.CoverTab[46677]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:134
	p := &data[cap(data)-1]
	m.Lock()
	defer m.Unlock()
	b := m.active[p]
	if b == nil || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:138
		_go_fuzz_dep_.CoverTab[46683]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:138
		return &b[0] != &data[0]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:138
		// _ = "end of CoverTab[46683]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:138
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:138
		_go_fuzz_dep_.CoverTab[46684]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:139
		// _ = "end of CoverTab[46684]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:140
		_go_fuzz_dep_.CoverTab[46685]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:140
		// _ = "end of CoverTab[46685]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:140
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:140
	// _ = "end of CoverTab[46677]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:140
	_go_fuzz_dep_.CoverTab[46678]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:143
	if errno := m.munmap(uintptr(unsafe.Pointer(&b[0])), uintptr(len(b))); errno != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:143
		_go_fuzz_dep_.CoverTab[46686]++
												return errno
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:144
		// _ = "end of CoverTab[46686]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:145
		_go_fuzz_dep_.CoverTab[46687]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:145
		// _ = "end of CoverTab[46687]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:145
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:145
	// _ = "end of CoverTab[46678]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:145
	_go_fuzz_dep_.CoverTab[46679]++
											delete(m.active, p)
											return nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:147
	// _ = "end of CoverTab[46679]"
}

func Read(fd int, p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:150
	_go_fuzz_dep_.CoverTab[46688]++
											n, err = read(fd, p)
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:152
		_go_fuzz_dep_.CoverTab[46690]++
												if n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:153
			_go_fuzz_dep_.CoverTab[46692]++
													raceWriteRange(unsafe.Pointer(&p[0]), n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:154
			// _ = "end of CoverTab[46692]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:155
			_go_fuzz_dep_.CoverTab[46693]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:155
			// _ = "end of CoverTab[46693]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:155
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:155
		// _ = "end of CoverTab[46690]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:155
		_go_fuzz_dep_.CoverTab[46691]++
												if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:156
			_go_fuzz_dep_.CoverTab[46694]++
													raceAcquire(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:157
			// _ = "end of CoverTab[46694]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:158
			_go_fuzz_dep_.CoverTab[46695]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:158
			// _ = "end of CoverTab[46695]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:158
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:158
		// _ = "end of CoverTab[46691]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:159
		_go_fuzz_dep_.CoverTab[46696]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:159
		// _ = "end of CoverTab[46696]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:159
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:159
	// _ = "end of CoverTab[46688]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:159
	_go_fuzz_dep_.CoverTab[46689]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:160
	// _ = "end of CoverTab[46689]"
}

func Write(fd int, p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:163
	_go_fuzz_dep_.CoverTab[46697]++
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:164
		_go_fuzz_dep_.CoverTab[46700]++
												raceReleaseMerge(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:165
		// _ = "end of CoverTab[46700]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:166
		_go_fuzz_dep_.CoverTab[46701]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:166
		// _ = "end of CoverTab[46701]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:166
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:166
	// _ = "end of CoverTab[46697]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:166
	_go_fuzz_dep_.CoverTab[46698]++
											n, err = write(fd, p)
											if raceenabled && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:168
		_go_fuzz_dep_.CoverTab[46702]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:168
		return n > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:168
		// _ = "end of CoverTab[46702]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:168
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:168
		_go_fuzz_dep_.CoverTab[46703]++
												raceReadRange(unsafe.Pointer(&p[0]), n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:169
		// _ = "end of CoverTab[46703]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:170
		_go_fuzz_dep_.CoverTab[46704]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:170
		// _ = "end of CoverTab[46704]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:170
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:170
	// _ = "end of CoverTab[46698]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:170
	_go_fuzz_dep_.CoverTab[46699]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:171
	// _ = "end of CoverTab[46699]"
}

func Pread(fd int, p []byte, offset int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:174
	_go_fuzz_dep_.CoverTab[46705]++
											n, err = pread(fd, p, offset)
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:176
		_go_fuzz_dep_.CoverTab[46707]++
												if n > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:177
			_go_fuzz_dep_.CoverTab[46709]++
													raceWriteRange(unsafe.Pointer(&p[0]), n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:178
			// _ = "end of CoverTab[46709]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:179
			_go_fuzz_dep_.CoverTab[46710]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:179
			// _ = "end of CoverTab[46710]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:179
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:179
		// _ = "end of CoverTab[46707]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:179
		_go_fuzz_dep_.CoverTab[46708]++
												if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:180
			_go_fuzz_dep_.CoverTab[46711]++
													raceAcquire(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:181
			// _ = "end of CoverTab[46711]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:182
			_go_fuzz_dep_.CoverTab[46712]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:182
			// _ = "end of CoverTab[46712]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:182
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:182
		// _ = "end of CoverTab[46708]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:183
		_go_fuzz_dep_.CoverTab[46713]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:183
		// _ = "end of CoverTab[46713]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:183
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:183
	// _ = "end of CoverTab[46705]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:183
	_go_fuzz_dep_.CoverTab[46706]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:184
	// _ = "end of CoverTab[46706]"
}

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:187
	_go_fuzz_dep_.CoverTab[46714]++
											if raceenabled {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:188
		_go_fuzz_dep_.CoverTab[46717]++
												raceReleaseMerge(unsafe.Pointer(&ioSync))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:189
		// _ = "end of CoverTab[46717]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:190
		_go_fuzz_dep_.CoverTab[46718]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:190
		// _ = "end of CoverTab[46718]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:190
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:190
	// _ = "end of CoverTab[46714]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:190
	_go_fuzz_dep_.CoverTab[46715]++
											n, err = pwrite(fd, p, offset)
											if raceenabled && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:192
		_go_fuzz_dep_.CoverTab[46719]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:192
		return n > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:192
		// _ = "end of CoverTab[46719]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:192
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:192
		_go_fuzz_dep_.CoverTab[46720]++
												raceReadRange(unsafe.Pointer(&p[0]), n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:193
		// _ = "end of CoverTab[46720]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:194
		_go_fuzz_dep_.CoverTab[46721]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:194
		// _ = "end of CoverTab[46721]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:194
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:194
	// _ = "end of CoverTab[46715]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:194
	_go_fuzz_dep_.CoverTab[46716]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:195
	// _ = "end of CoverTab[46716]"
}

// For testing: clients can set this flag to force
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:198
// creation of IPv6 sockets to return EAFNOSUPPORT.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:200
var SocketDisableIPv6 bool

// Sockaddr represents a socket address.
type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len _Socklen, err error)	// lowercase; only we can define Sockaddrs
}

// SockaddrInet4 implements the Sockaddr interface for AF_INET type sockets.
type SockaddrInet4 struct {
	Port	int
	Addr	[4]byte
	raw	RawSockaddrInet4
}

// SockaddrInet6 implements the Sockaddr interface for AF_INET6 type sockets.
type SockaddrInet6 struct {
	Port	int
	ZoneId	uint32
	Addr	[16]byte
	raw	RawSockaddrInet6
}

// SockaddrUnix implements the Sockaddr interface for AF_UNIX type sockets.
type SockaddrUnix struct {
	Name	string
	raw	RawSockaddrUnix
}

func Bind(fd int, sa Sockaddr) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:228
	_go_fuzz_dep_.CoverTab[46722]++
											ptr, n, err := sa.sockaddr()
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:230
		_go_fuzz_dep_.CoverTab[46724]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:231
		// _ = "end of CoverTab[46724]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:232
		_go_fuzz_dep_.CoverTab[46725]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:232
		// _ = "end of CoverTab[46725]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:232
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:232
	// _ = "end of CoverTab[46722]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:232
	_go_fuzz_dep_.CoverTab[46723]++
											return bind(fd, ptr, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:233
	// _ = "end of CoverTab[46723]"
}

func Connect(fd int, sa Sockaddr) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:236
	_go_fuzz_dep_.CoverTab[46726]++
											ptr, n, err := sa.sockaddr()
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:238
		_go_fuzz_dep_.CoverTab[46728]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:239
		// _ = "end of CoverTab[46728]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:240
		_go_fuzz_dep_.CoverTab[46729]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:240
		// _ = "end of CoverTab[46729]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:240
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:240
	// _ = "end of CoverTab[46726]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:240
	_go_fuzz_dep_.CoverTab[46727]++
											return connect(fd, ptr, n)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:241
	// _ = "end of CoverTab[46727]"
}

func Getpeername(fd int) (sa Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:244
	_go_fuzz_dep_.CoverTab[46730]++
											var rsa RawSockaddrAny
											var len _Socklen = SizeofSockaddrAny
											if err = getpeername(fd, &rsa, &len); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:247
		_go_fuzz_dep_.CoverTab[46732]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:248
		// _ = "end of CoverTab[46732]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:249
		_go_fuzz_dep_.CoverTab[46733]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:249
		// _ = "end of CoverTab[46733]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:249
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:249
	// _ = "end of CoverTab[46730]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:249
	_go_fuzz_dep_.CoverTab[46731]++
											return anyToSockaddr(fd, &rsa)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:250
	// _ = "end of CoverTab[46731]"
}

func GetsockoptByte(fd, level, opt int) (value byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:253
	_go_fuzz_dep_.CoverTab[46734]++
											var n byte
											vallen := _Socklen(1)
											err = getsockopt(fd, level, opt, unsafe.Pointer(&n), &vallen)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:257
	// _ = "end of CoverTab[46734]"
}

func GetsockoptInt(fd, level, opt int) (value int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:260
	_go_fuzz_dep_.CoverTab[46735]++
											var n int32
											vallen := _Socklen(4)
											err = getsockopt(fd, level, opt, unsafe.Pointer(&n), &vallen)
											return int(n), err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:264
	// _ = "end of CoverTab[46735]"
}

func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:267
	_go_fuzz_dep_.CoverTab[46736]++
											vallen := _Socklen(4)
											err = getsockopt(fd, level, opt, unsafe.Pointer(&value[0]), &vallen)
											return value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:270
	// _ = "end of CoverTab[46736]"
}

func GetsockoptIPMreq(fd, level, opt int) (*IPMreq, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:273
	_go_fuzz_dep_.CoverTab[46737]++
											var value IPMreq
											vallen := _Socklen(SizeofIPMreq)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:277
	// _ = "end of CoverTab[46737]"
}

func GetsockoptIPv6Mreq(fd, level, opt int) (*IPv6Mreq, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:280
	_go_fuzz_dep_.CoverTab[46738]++
											var value IPv6Mreq
											vallen := _Socklen(SizeofIPv6Mreq)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:284
	// _ = "end of CoverTab[46738]"
}

func GetsockoptIPv6MTUInfo(fd, level, opt int) (*IPv6MTUInfo, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:287
	_go_fuzz_dep_.CoverTab[46739]++
											var value IPv6MTUInfo
											vallen := _Socklen(SizeofIPv6MTUInfo)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:291
	// _ = "end of CoverTab[46739]"
}

func GetsockoptICMPv6Filter(fd, level, opt int) (*ICMPv6Filter, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:294
	_go_fuzz_dep_.CoverTab[46740]++
											var value ICMPv6Filter
											vallen := _Socklen(SizeofICMPv6Filter)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:298
	// _ = "end of CoverTab[46740]"
}

func GetsockoptLinger(fd, level, opt int) (*Linger, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:301
	_go_fuzz_dep_.CoverTab[46741]++
											var linger Linger
											vallen := _Socklen(SizeofLinger)
											err := getsockopt(fd, level, opt, unsafe.Pointer(&linger), &vallen)
											return &linger, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:305
	// _ = "end of CoverTab[46741]"
}

func GetsockoptTimeval(fd, level, opt int) (*Timeval, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:308
	_go_fuzz_dep_.CoverTab[46742]++
											var tv Timeval
											vallen := _Socklen(unsafe.Sizeof(tv))
											err := getsockopt(fd, level, opt, unsafe.Pointer(&tv), &vallen)
											return &tv, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:312
	// _ = "end of CoverTab[46742]"
}

func GetsockoptUint64(fd, level, opt int) (value uint64, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:315
	_go_fuzz_dep_.CoverTab[46743]++
											var n uint64
											vallen := _Socklen(8)
											err = getsockopt(fd, level, opt, unsafe.Pointer(&n), &vallen)
											return n, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:319
	// _ = "end of CoverTab[46743]"
}

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:322
	_go_fuzz_dep_.CoverTab[46744]++
											var rsa RawSockaddrAny
											var len _Socklen = SizeofSockaddrAny
											if n, err = recvfrom(fd, p, flags, &rsa, &len); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:325
		_go_fuzz_dep_.CoverTab[46747]++
												return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:326
		// _ = "end of CoverTab[46747]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:327
		_go_fuzz_dep_.CoverTab[46748]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:327
		// _ = "end of CoverTab[46748]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:327
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:327
	// _ = "end of CoverTab[46744]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:327
	_go_fuzz_dep_.CoverTab[46745]++
											if rsa.Addr.Family != AF_UNSPEC {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:328
		_go_fuzz_dep_.CoverTab[46749]++
												from, err = anyToSockaddr(fd, &rsa)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:329
		// _ = "end of CoverTab[46749]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:330
		_go_fuzz_dep_.CoverTab[46750]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:330
		// _ = "end of CoverTab[46750]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:330
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:330
	// _ = "end of CoverTab[46745]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:330
	_go_fuzz_dep_.CoverTab[46746]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:331
	// _ = "end of CoverTab[46746]"
}

// Recvmsg receives a message from a socket using the recvmsg system call. The
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
// received non-control data will be written to p, and any "out of band"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
// control data will be written to oob. The flags are passed to recvmsg.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
// The results are:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
//   - n is the number of non-control data bytes read into p
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
//   - oobn is the number of control data bytes read into oob; this may be interpreted using [ParseSocketControlMessage]
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
//   - recvflags is flags returned by recvmsg
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
//   - from is the address of the sender
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
// If the underlying socket type is not SOCK_DGRAM, a received message
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
// containing oob data and a single '\0' of non-control data is treated as if
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:334
// the message contained only control data, i.e. n will be zero on return.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:347
func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:347
	_go_fuzz_dep_.CoverTab[46751]++
											var iov [1]Iovec
											if len(p) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:349
		_go_fuzz_dep_.CoverTab[46754]++
												iov[0].Base = &p[0]
												iov[0].SetLen(len(p))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:351
		// _ = "end of CoverTab[46754]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:352
		_go_fuzz_dep_.CoverTab[46755]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:352
		// _ = "end of CoverTab[46755]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:352
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:352
	// _ = "end of CoverTab[46751]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:352
	_go_fuzz_dep_.CoverTab[46752]++
											var rsa RawSockaddrAny
											n, oobn, recvflags, err = recvmsgRaw(fd, iov[:], oob, flags, &rsa)

											if rsa.Addr.Family != AF_UNSPEC {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:356
		_go_fuzz_dep_.CoverTab[46756]++
												from, err = anyToSockaddr(fd, &rsa)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:357
		// _ = "end of CoverTab[46756]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:358
		_go_fuzz_dep_.CoverTab[46757]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:358
		// _ = "end of CoverTab[46757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:358
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:358
	// _ = "end of CoverTab[46752]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:358
	_go_fuzz_dep_.CoverTab[46753]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:359
	// _ = "end of CoverTab[46753]"
}

// RecvmsgBuffers receives a message from a socket using the recvmsg system
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:362
// call. This function is equivalent to Recvmsg, but non-control data read is
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:362
// scattered into the buffers slices.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:365
func RecvmsgBuffers(fd int, buffers [][]byte, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:365
	_go_fuzz_dep_.CoverTab[46758]++
											iov := make([]Iovec, len(buffers))
											for i := range buffers {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:367
		_go_fuzz_dep_.CoverTab[46761]++
												if len(buffers[i]) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:368
			_go_fuzz_dep_.CoverTab[46762]++
													iov[i].Base = &buffers[i][0]
													iov[i].SetLen(len(buffers[i]))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:370
			// _ = "end of CoverTab[46762]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:371
			_go_fuzz_dep_.CoverTab[46763]++
													iov[i].Base = (*byte)(unsafe.Pointer(&_zero))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:372
			// _ = "end of CoverTab[46763]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:373
		// _ = "end of CoverTab[46761]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:374
	// _ = "end of CoverTab[46758]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:374
	_go_fuzz_dep_.CoverTab[46759]++
											var rsa RawSockaddrAny
											n, oobn, recvflags, err = recvmsgRaw(fd, iov, oob, flags, &rsa)
											if err == nil && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:377
		_go_fuzz_dep_.CoverTab[46764]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:377
		return rsa.Addr.Family != AF_UNSPEC
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:377
		// _ = "end of CoverTab[46764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:377
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:377
		_go_fuzz_dep_.CoverTab[46765]++
												from, err = anyToSockaddr(fd, &rsa)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:378
		// _ = "end of CoverTab[46765]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:379
		_go_fuzz_dep_.CoverTab[46766]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:379
		// _ = "end of CoverTab[46766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:379
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:379
	// _ = "end of CoverTab[46759]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:379
	_go_fuzz_dep_.CoverTab[46760]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:380
	// _ = "end of CoverTab[46760]"
}

// Sendmsg sends a message on a socket to an address using the sendmsg system
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:383
// call. This function is equivalent to SendmsgN, but does not return the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:383
// number of bytes actually sent.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:386
func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:386
	_go_fuzz_dep_.CoverTab[46767]++
											_, err = SendmsgN(fd, p, oob, to, flags)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:388
	// _ = "end of CoverTab[46767]"
}

// SendmsgN sends a message on a socket to an address using the sendmsg system
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// call. p contains the non-control data to send, and oob contains the "out of
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// band" control data. The flags are passed to sendmsg. The number of
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// non-control bytes actually written to the socket is returned.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// Some socket types do not support sending control data without accompanying
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// non-control data. If p is empty, and oob contains control data, and the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// underlying socket type is not SOCK_DGRAM, p will be treated as containing a
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// single '\0' and the return value will indicate zero bytes sent.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// The Go function Recvmsg, if called with an empty p and a non-empty oob,
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// will read and ignore this additional '\0'.  If the message is received by
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// code that does not use Recvmsg, or that does not use Go at all, that code
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// will need to be written to expect and ignore the additional '\0'.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// If you need to send non-empty oob with p actually empty, and if the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// underlying socket type supports it, you can do so via a raw system call as
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
// follows:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//	msg := &unix.Msghdr{
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//	    Control: &oob[0],
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//	msg.SetControllen(len(oob))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:391
//	n, _, errno := unix.Syscall(unix.SYS_SENDMSG, uintptr(fd), uintptr(unsafe.Pointer(msg)), flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:415
func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:415
	_go_fuzz_dep_.CoverTab[46768]++
											var iov [1]Iovec
											if len(p) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:417
		_go_fuzz_dep_.CoverTab[46771]++
												iov[0].Base = &p[0]
												iov[0].SetLen(len(p))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:419
		// _ = "end of CoverTab[46771]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:420
		_go_fuzz_dep_.CoverTab[46772]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:420
		// _ = "end of CoverTab[46772]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:420
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:420
	// _ = "end of CoverTab[46768]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:420
	_go_fuzz_dep_.CoverTab[46769]++
											var ptr unsafe.Pointer
											var salen _Socklen
											if to != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:423
		_go_fuzz_dep_.CoverTab[46773]++
												ptr, salen, err = to.sockaddr()
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:425
			_go_fuzz_dep_.CoverTab[46774]++
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:426
			// _ = "end of CoverTab[46774]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:427
			_go_fuzz_dep_.CoverTab[46775]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:427
			// _ = "end of CoverTab[46775]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:427
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:427
		// _ = "end of CoverTab[46773]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:428
		_go_fuzz_dep_.CoverTab[46776]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:428
		// _ = "end of CoverTab[46776]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:428
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:428
	// _ = "end of CoverTab[46769]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:428
	_go_fuzz_dep_.CoverTab[46770]++
											return sendmsgN(fd, iov[:], oob, ptr, salen, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:429
	// _ = "end of CoverTab[46770]"
}

// SendmsgBuffers sends a message on a socket to an address using the sendmsg
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:432
// system call. This function is equivalent to SendmsgN, but the non-control
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:432
// data is gathered from buffers.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:435
func SendmsgBuffers(fd int, buffers [][]byte, oob []byte, to Sockaddr, flags int) (n int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:435
	_go_fuzz_dep_.CoverTab[46777]++
											iov := make([]Iovec, len(buffers))
											for i := range buffers {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:437
		_go_fuzz_dep_.CoverTab[46780]++
												if len(buffers[i]) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:438
			_go_fuzz_dep_.CoverTab[46781]++
													iov[i].Base = &buffers[i][0]
													iov[i].SetLen(len(buffers[i]))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:440
			// _ = "end of CoverTab[46781]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:441
			_go_fuzz_dep_.CoverTab[46782]++
													iov[i].Base = (*byte)(unsafe.Pointer(&_zero))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:442
			// _ = "end of CoverTab[46782]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:443
		// _ = "end of CoverTab[46780]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:444
	// _ = "end of CoverTab[46777]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:444
	_go_fuzz_dep_.CoverTab[46778]++
											var ptr unsafe.Pointer
											var salen _Socklen
											if to != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:447
		_go_fuzz_dep_.CoverTab[46783]++
												ptr, salen, err = to.sockaddr()
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:449
			_go_fuzz_dep_.CoverTab[46784]++
													return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:450
			// _ = "end of CoverTab[46784]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:451
			_go_fuzz_dep_.CoverTab[46785]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:451
			// _ = "end of CoverTab[46785]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:451
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:451
		// _ = "end of CoverTab[46783]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:452
		_go_fuzz_dep_.CoverTab[46786]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:452
		// _ = "end of CoverTab[46786]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:452
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:452
	// _ = "end of CoverTab[46778]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:452
	_go_fuzz_dep_.CoverTab[46779]++
											return sendmsgN(fd, iov, oob, ptr, salen, flags)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:453
	// _ = "end of CoverTab[46779]"
}

func Send(s int, buf []byte, flags int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:456
	_go_fuzz_dep_.CoverTab[46787]++
											return sendto(s, buf, flags, nil, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:457
	// _ = "end of CoverTab[46787]"
}

func Sendto(fd int, p []byte, flags int, to Sockaddr) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:460
	_go_fuzz_dep_.CoverTab[46788]++
											var ptr unsafe.Pointer
											var salen _Socklen
											if to != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:463
		_go_fuzz_dep_.CoverTab[46790]++
												ptr, salen, err = to.sockaddr()
												if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:465
			_go_fuzz_dep_.CoverTab[46791]++
													return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:466
			// _ = "end of CoverTab[46791]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:467
			_go_fuzz_dep_.CoverTab[46792]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:467
			// _ = "end of CoverTab[46792]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:467
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:467
		// _ = "end of CoverTab[46790]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:468
		_go_fuzz_dep_.CoverTab[46793]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:468
		// _ = "end of CoverTab[46793]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:468
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:468
	// _ = "end of CoverTab[46788]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:468
	_go_fuzz_dep_.CoverTab[46789]++
											return sendto(fd, p, flags, ptr, salen)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:469
	// _ = "end of CoverTab[46789]"
}

func SetsockoptByte(fd, level, opt int, value byte) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:472
	_go_fuzz_dep_.CoverTab[46794]++
											return setsockopt(fd, level, opt, unsafe.Pointer(&value), 1)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:473
	// _ = "end of CoverTab[46794]"
}

func SetsockoptInt(fd, level, opt int, value int) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:476
	_go_fuzz_dep_.CoverTab[46795]++
											var n = int32(value)
											return setsockopt(fd, level, opt, unsafe.Pointer(&n), 4)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:478
	// _ = "end of CoverTab[46795]"
}

func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:481
	_go_fuzz_dep_.CoverTab[46796]++
											return setsockopt(fd, level, opt, unsafe.Pointer(&value[0]), 4)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:482
	// _ = "end of CoverTab[46796]"
}

func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:485
	_go_fuzz_dep_.CoverTab[46797]++
											return setsockopt(fd, level, opt, unsafe.Pointer(mreq), SizeofIPMreq)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:486
	// _ = "end of CoverTab[46797]"
}

func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:489
	_go_fuzz_dep_.CoverTab[46798]++
											return setsockopt(fd, level, opt, unsafe.Pointer(mreq), SizeofIPv6Mreq)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:490
	// _ = "end of CoverTab[46798]"
}

func SetsockoptICMPv6Filter(fd, level, opt int, filter *ICMPv6Filter) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:493
	_go_fuzz_dep_.CoverTab[46799]++
											return setsockopt(fd, level, opt, unsafe.Pointer(filter), SizeofICMPv6Filter)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:494
	// _ = "end of CoverTab[46799]"
}

func SetsockoptLinger(fd, level, opt int, l *Linger) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:497
	_go_fuzz_dep_.CoverTab[46800]++
											return setsockopt(fd, level, opt, unsafe.Pointer(l), SizeofLinger)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:498
	// _ = "end of CoverTab[46800]"
}

func SetsockoptString(fd, level, opt int, s string) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:501
	_go_fuzz_dep_.CoverTab[46801]++
											var p unsafe.Pointer
											if len(s) > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:503
		_go_fuzz_dep_.CoverTab[46803]++
												p = unsafe.Pointer(&[]byte(s)[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:504
		// _ = "end of CoverTab[46803]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:505
		_go_fuzz_dep_.CoverTab[46804]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:505
		// _ = "end of CoverTab[46804]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:505
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:505
	// _ = "end of CoverTab[46801]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:505
	_go_fuzz_dep_.CoverTab[46802]++
											return setsockopt(fd, level, opt, p, uintptr(len(s)))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:506
	// _ = "end of CoverTab[46802]"
}

func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:509
	_go_fuzz_dep_.CoverTab[46805]++
											return setsockopt(fd, level, opt, unsafe.Pointer(tv), unsafe.Sizeof(*tv))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:510
	// _ = "end of CoverTab[46805]"
}

func SetsockoptUint64(fd, level, opt int, value uint64) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:513
	_go_fuzz_dep_.CoverTab[46806]++
											return setsockopt(fd, level, opt, unsafe.Pointer(&value), 8)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:514
	// _ = "end of CoverTab[46806]"
}

func Socket(domain, typ, proto int) (fd int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:517
	_go_fuzz_dep_.CoverTab[46807]++
											if domain == AF_INET6 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:518
		_go_fuzz_dep_.CoverTab[46809]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:518
		return SocketDisableIPv6
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:518
		// _ = "end of CoverTab[46809]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:518
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:518
		_go_fuzz_dep_.CoverTab[46810]++
												return -1, EAFNOSUPPORT
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:519
		// _ = "end of CoverTab[46810]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:520
		_go_fuzz_dep_.CoverTab[46811]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:520
		// _ = "end of CoverTab[46811]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:520
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:520
	// _ = "end of CoverTab[46807]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:520
	_go_fuzz_dep_.CoverTab[46808]++
											fd, err = socket(domain, typ, proto)
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:522
	// _ = "end of CoverTab[46808]"
}

func Socketpair(domain, typ, proto int) (fd [2]int, err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:525
	_go_fuzz_dep_.CoverTab[46812]++
											var fdx [2]int32
											err = socketpair(domain, typ, proto, &fdx)
											if err == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:528
		_go_fuzz_dep_.CoverTab[46814]++
												fd[0] = int(fdx[0])
												fd[1] = int(fdx[1])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:530
		// _ = "end of CoverTab[46814]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:531
		_go_fuzz_dep_.CoverTab[46815]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:531
		// _ = "end of CoverTab[46815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:531
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:531
	// _ = "end of CoverTab[46812]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:531
	_go_fuzz_dep_.CoverTab[46813]++
											return
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:532
	// _ = "end of CoverTab[46813]"
}

var ioSync int64

func CloseOnExec(fd int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:537
	_go_fuzz_dep_.CoverTab[46816]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:537
	fcntl(fd, F_SETFD, FD_CLOEXEC)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:537
	// _ = "end of CoverTab[46816]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:537
}

func SetNonblock(fd int, nonblocking bool) (err error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:539
	_go_fuzz_dep_.CoverTab[46817]++
											flag, err := fcntl(fd, F_GETFL, 0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:541
		_go_fuzz_dep_.CoverTab[46820]++
												return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:542
		// _ = "end of CoverTab[46820]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:543
		_go_fuzz_dep_.CoverTab[46821]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:543
		// _ = "end of CoverTab[46821]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:543
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:543
	// _ = "end of CoverTab[46817]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:543
	_go_fuzz_dep_.CoverTab[46818]++
											if nonblocking {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:544
		_go_fuzz_dep_.CoverTab[46822]++
												flag |= O_NONBLOCK
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:545
		// _ = "end of CoverTab[46822]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:546
		_go_fuzz_dep_.CoverTab[46823]++
												flag &= ^O_NONBLOCK
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:547
		// _ = "end of CoverTab[46823]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:548
	// _ = "end of CoverTab[46818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:548
	_go_fuzz_dep_.CoverTab[46819]++
											_, err = fcntl(fd, F_SETFL, flag)
											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:550
	// _ = "end of CoverTab[46819]"
}

// Exec calls execve(2), which replaces the calling executable in the process
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:553
// tree. argv0 should be the full path to an executable ("/bin/ls") and the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:553
// executable name should also be the first argument in argv (["ls", "-l"]).
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:553
// envv are the environment variables that should be passed to the new
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:553
// process (["USER=go", "PWD=/tmp"]).
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:558
func Exec(argv0 string, argv []string, envv []string) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:558
	_go_fuzz_dep_.CoverTab[46824]++
											return syscall.Exec(argv0, argv, envv)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:559
	// _ = "end of CoverTab[46824]"
}

// Lutimes sets the access and modification times tv on path. If path refers to
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:562
// a symlink, it is not dereferenced and the timestamps are set on the symlink.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:562
// If tv is nil, the access and modification times are set to the current time.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:562
// Otherwise tv must contain exactly 2 elements, with access time as the first
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:562
// element and modification time as the second element.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:567
func Lutimes(path string, tv []Timeval) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:567
	_go_fuzz_dep_.CoverTab[46825]++
											if tv == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:568
		_go_fuzz_dep_.CoverTab[46828]++
												return UtimesNanoAt(AT_FDCWD, path, nil, AT_SYMLINK_NOFOLLOW)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:569
		// _ = "end of CoverTab[46828]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:570
		_go_fuzz_dep_.CoverTab[46829]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:570
		// _ = "end of CoverTab[46829]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:570
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:570
	// _ = "end of CoverTab[46825]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:570
	_go_fuzz_dep_.CoverTab[46826]++
											if len(tv) != 2 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:571
		_go_fuzz_dep_.CoverTab[46830]++
												return EINVAL
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:572
		// _ = "end of CoverTab[46830]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:573
		_go_fuzz_dep_.CoverTab[46831]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:573
		// _ = "end of CoverTab[46831]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:573
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:573
	// _ = "end of CoverTab[46826]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:573
	_go_fuzz_dep_.CoverTab[46827]++
											ts := []Timespec{
		NsecToTimespec(TimevalToNsec(tv[0])),
		NsecToTimespec(TimevalToNsec(tv[1])),
	}
											return UtimesNanoAt(AT_FDCWD, path, ts, AT_SYMLINK_NOFOLLOW)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:578
	// _ = "end of CoverTab[46827]"
}

// emptyIovecs reports whether there are no bytes in the slice of Iovec.
func emptyIovecs(iov []Iovec) bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:582
	_go_fuzz_dep_.CoverTab[46832]++
											for i := range iov {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:583
		_go_fuzz_dep_.CoverTab[46834]++
												if iov[i].Len > 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:584
			_go_fuzz_dep_.CoverTab[46835]++
													return false
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:585
			// _ = "end of CoverTab[46835]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:586
			_go_fuzz_dep_.CoverTab[46836]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:586
			// _ = "end of CoverTab[46836]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:586
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:586
		// _ = "end of CoverTab[46834]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:587
	// _ = "end of CoverTab[46832]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:587
	_go_fuzz_dep_.CoverTab[46833]++
											return true
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:588
	// _ = "end of CoverTab[46833]"
}

// Setrlimit sets a resource limit.
func Setrlimit(resource int, rlim *Rlimit) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:592
	_go_fuzz_dep_.CoverTab[46837]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:595
	return syscall.Setrlimit(resource, (*syscall.Rlimit)(rlim))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:595
	// _ = "end of CoverTab[46837]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:596
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/syscall_unix.go:596
var _ = _go_fuzz_dep_.CoverTab
