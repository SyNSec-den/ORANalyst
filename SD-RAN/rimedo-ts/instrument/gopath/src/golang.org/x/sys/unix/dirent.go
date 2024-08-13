// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:8
)

import "unsafe"

// readInt returns the size-bytes unsigned integer in native byte order at offset off.
func readInt(b []byte, off, size uintptr) (u uint64, ok bool) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:13
	_go_fuzz_dep_.CoverTab[45722]++
										if len(b) < int(off+size) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:14
		_go_fuzz_dep_.CoverTab[45725]++
											return 0, false
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:15
		// _ = "end of CoverTab[45725]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:16
		_go_fuzz_dep_.CoverTab[45726]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:16
		// _ = "end of CoverTab[45726]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:16
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:16
	// _ = "end of CoverTab[45722]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:16
	_go_fuzz_dep_.CoverTab[45723]++
										if isBigEndian {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:17
		_go_fuzz_dep_.CoverTab[45727]++
											return readIntBE(b[off:], size), true
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:18
		// _ = "end of CoverTab[45727]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:19
		_go_fuzz_dep_.CoverTab[45728]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:19
		// _ = "end of CoverTab[45728]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:19
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:19
	// _ = "end of CoverTab[45723]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:19
	_go_fuzz_dep_.CoverTab[45724]++
										return readIntLE(b[off:], size), true
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:20
	// _ = "end of CoverTab[45724]"
}

func readIntBE(b []byte, size uintptr) uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:23
	_go_fuzz_dep_.CoverTab[45729]++
										switch size {
	case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:25
		_go_fuzz_dep_.CoverTab[45730]++
											return uint64(b[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:26
		// _ = "end of CoverTab[45730]"
	case 2:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:27
		_go_fuzz_dep_.CoverTab[45731]++
											_ = b[1]
											return uint64(b[1]) | uint64(b[0])<<8
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:29
		// _ = "end of CoverTab[45731]"
	case 4:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:30
		_go_fuzz_dep_.CoverTab[45732]++
											_ = b[3]
											return uint64(b[3]) | uint64(b[2])<<8 | uint64(b[1])<<16 | uint64(b[0])<<24
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:32
		// _ = "end of CoverTab[45732]"
	case 8:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:33
		_go_fuzz_dep_.CoverTab[45733]++
											_ = b[7]
											return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
			uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:36
		// _ = "end of CoverTab[45733]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:37
		_go_fuzz_dep_.CoverTab[45734]++
											panic("syscall: readInt with unsupported size")
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:38
		// _ = "end of CoverTab[45734]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:39
	// _ = "end of CoverTab[45729]"
}

func readIntLE(b []byte, size uintptr) uint64 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:42
	_go_fuzz_dep_.CoverTab[45735]++
										switch size {
	case 1:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:44
		_go_fuzz_dep_.CoverTab[45736]++
											return uint64(b[0])
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:45
		// _ = "end of CoverTab[45736]"
	case 2:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:46
		_go_fuzz_dep_.CoverTab[45737]++
											_ = b[1]
											return uint64(b[0]) | uint64(b[1])<<8
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:48
		// _ = "end of CoverTab[45737]"
	case 4:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:49
		_go_fuzz_dep_.CoverTab[45738]++
											_ = b[3]
											return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:51
		// _ = "end of CoverTab[45738]"
	case 8:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:52
		_go_fuzz_dep_.CoverTab[45739]++
											_ = b[7]
											return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
			uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:55
		// _ = "end of CoverTab[45739]"
	default:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:56
		_go_fuzz_dep_.CoverTab[45740]++
											panic("syscall: readInt with unsupported size")
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:57
		// _ = "end of CoverTab[45740]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:58
	// _ = "end of CoverTab[45735]"
}

// ParseDirent parses up to max directory entries in buf,
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:61
// appending the names to names. It returns the number of
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:61
// bytes consumed from buf, the number of entries added
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:61
// to names, and the new names slice.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:65
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:65
	_go_fuzz_dep_.CoverTab[45741]++
										origlen := len(buf)
										count = 0
										for max != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:68
		_go_fuzz_dep_.CoverTab[45743]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:68
		return len(buf) > 0
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:68
		// _ = "end of CoverTab[45743]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:68
	}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:68
		_go_fuzz_dep_.CoverTab[45744]++
											reclen, ok := direntReclen(buf)
											if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:70
			_go_fuzz_dep_.CoverTab[45751]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:70
			return reclen > uint64(len(buf))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:70
			// _ = "end of CoverTab[45751]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:70
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:70
			_go_fuzz_dep_.CoverTab[45752]++
												return origlen, count, names
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:71
			// _ = "end of CoverTab[45752]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:72
			_go_fuzz_dep_.CoverTab[45753]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:72
			// _ = "end of CoverTab[45753]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:72
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:72
		// _ = "end of CoverTab[45744]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:72
		_go_fuzz_dep_.CoverTab[45745]++
											rec := buf[:reclen]
											buf = buf[reclen:]
											ino, ok := direntIno(rec)
											if !ok {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:76
			_go_fuzz_dep_.CoverTab[45754]++
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:77
			// _ = "end of CoverTab[45754]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:78
			_go_fuzz_dep_.CoverTab[45755]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:78
			// _ = "end of CoverTab[45755]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:78
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:78
		// _ = "end of CoverTab[45745]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:78
		_go_fuzz_dep_.CoverTab[45746]++
											if ino == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:79
			_go_fuzz_dep_.CoverTab[45756]++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:80
			// _ = "end of CoverTab[45756]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:81
			_go_fuzz_dep_.CoverTab[45757]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:81
			// _ = "end of CoverTab[45757]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:81
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:81
		// _ = "end of CoverTab[45746]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:81
		_go_fuzz_dep_.CoverTab[45747]++
											const namoff = uint64(unsafe.Offsetof(Dirent{}.Name))
											namlen, ok := direntNamlen(rec)
											if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:84
			_go_fuzz_dep_.CoverTab[45758]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:84
			return namoff+namlen > uint64(len(rec))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:84
			// _ = "end of CoverTab[45758]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:84
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:84
			_go_fuzz_dep_.CoverTab[45759]++
												break
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:85
			// _ = "end of CoverTab[45759]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:86
			_go_fuzz_dep_.CoverTab[45760]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:86
			// _ = "end of CoverTab[45760]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:86
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:86
		// _ = "end of CoverTab[45747]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:86
		_go_fuzz_dep_.CoverTab[45748]++
											name := rec[namoff : namoff+namlen]
											for i, c := range name {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:88
			_go_fuzz_dep_.CoverTab[45761]++
												if c == 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:89
				_go_fuzz_dep_.CoverTab[45762]++
													name = name[:i]
													break
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:91
				// _ = "end of CoverTab[45762]"
			} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:92
				_go_fuzz_dep_.CoverTab[45763]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:92
				// _ = "end of CoverTab[45763]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:92
			}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:92
			// _ = "end of CoverTab[45761]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:93
		// _ = "end of CoverTab[45748]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:93
		_go_fuzz_dep_.CoverTab[45749]++

											if string(name) == "." || func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:95
			_go_fuzz_dep_.CoverTab[45764]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:95
			return string(name) == ".."
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:95
			// _ = "end of CoverTab[45764]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:95
		}() {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:95
			_go_fuzz_dep_.CoverTab[45765]++
												continue
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:96
			// _ = "end of CoverTab[45765]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:97
			_go_fuzz_dep_.CoverTab[45766]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:97
			// _ = "end of CoverTab[45766]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:97
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:97
		// _ = "end of CoverTab[45749]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:97
		_go_fuzz_dep_.CoverTab[45750]++
											max--
											count++
											names = append(names, string(name))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:100
		// _ = "end of CoverTab[45750]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:101
	// _ = "end of CoverTab[45741]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:101
	_go_fuzz_dep_.CoverTab[45742]++
										return origlen - len(buf), count, names
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:102
	// _ = "end of CoverTab[45742]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:103
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/dirent.go:103
var _ = _go_fuzz_dep_.CoverTab
