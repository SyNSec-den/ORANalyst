// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/internal/saferio/io.go:5
// Package saferio provides I/O functions that avoid allocating large
//line /usr/local/go/src/internal/saferio/io.go:5
// amounts of memory unnecessarily. This is intended for packages that
//line /usr/local/go/src/internal/saferio/io.go:5
// read data from an [io.Reader] where the size is part of the input
//line /usr/local/go/src/internal/saferio/io.go:5
// data but the input may be corrupt, or may be provided by an
//line /usr/local/go/src/internal/saferio/io.go:5
// untrustworthy attacker.
//line /usr/local/go/src/internal/saferio/io.go:10
package saferio

//line /usr/local/go/src/internal/saferio/io.go:10
import (
//line /usr/local/go/src/internal/saferio/io.go:10
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/internal/saferio/io.go:10
)
//line /usr/local/go/src/internal/saferio/io.go:10
import (
//line /usr/local/go/src/internal/saferio/io.go:10
	_atomic_ "sync/atomic"
//line /usr/local/go/src/internal/saferio/io.go:10
)

import (
	"io"
	"reflect"
)

// chunk is an arbitrary limit on how much memory we are willing
//line /usr/local/go/src/internal/saferio/io.go:17
// to allocate without concern.
//line /usr/local/go/src/internal/saferio/io.go:19
const chunk = 10 << 20	// 10M

// ReadData reads n bytes from the input stream, but avoids allocating
//line /usr/local/go/src/internal/saferio/io.go:21
// all n bytes if n is large. This avoids crashing the program by
//line /usr/local/go/src/internal/saferio/io.go:21
// allocating all n bytes in cases where n is incorrect.
//line /usr/local/go/src/internal/saferio/io.go:21
//
//line /usr/local/go/src/internal/saferio/io.go:21
// The error is io.EOF only if no bytes were read.
//line /usr/local/go/src/internal/saferio/io.go:21
// If an io.EOF happens after reading some but not all the bytes,
//line /usr/local/go/src/internal/saferio/io.go:21
// ReadData returns io.ErrUnexpectedEOF.
//line /usr/local/go/src/internal/saferio/io.go:28
func ReadData(r io.Reader, n uint64) ([]byte, error) {
//line /usr/local/go/src/internal/saferio/io.go:28
	_go_fuzz_dep_.CoverTab[83676]++
							if int64(n) < 0 || func() bool {
//line /usr/local/go/src/internal/saferio/io.go:29
		_go_fuzz_dep_.CoverTab[83680]++
//line /usr/local/go/src/internal/saferio/io.go:29
		return n != uint64(int(n))
//line /usr/local/go/src/internal/saferio/io.go:29
		// _ = "end of CoverTab[83680]"
//line /usr/local/go/src/internal/saferio/io.go:29
	}() {
//line /usr/local/go/src/internal/saferio/io.go:29
		_go_fuzz_dep_.CoverTab[83681]++

//line /usr/local/go/src/internal/saferio/io.go:32
		return nil, io.ErrUnexpectedEOF
//line /usr/local/go/src/internal/saferio/io.go:32
		// _ = "end of CoverTab[83681]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:33
		_go_fuzz_dep_.CoverTab[83682]++
//line /usr/local/go/src/internal/saferio/io.go:33
		// _ = "end of CoverTab[83682]"
//line /usr/local/go/src/internal/saferio/io.go:33
	}
//line /usr/local/go/src/internal/saferio/io.go:33
	// _ = "end of CoverTab[83676]"
//line /usr/local/go/src/internal/saferio/io.go:33
	_go_fuzz_dep_.CoverTab[83677]++

							if n < chunk {
//line /usr/local/go/src/internal/saferio/io.go:35
		_go_fuzz_dep_.CoverTab[83683]++
								buf := make([]byte, n)
								_, err := io.ReadFull(r, buf)
								if err != nil {
//line /usr/local/go/src/internal/saferio/io.go:38
			_go_fuzz_dep_.CoverTab[83685]++
									return nil, err
//line /usr/local/go/src/internal/saferio/io.go:39
			// _ = "end of CoverTab[83685]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:40
			_go_fuzz_dep_.CoverTab[83686]++
//line /usr/local/go/src/internal/saferio/io.go:40
			// _ = "end of CoverTab[83686]"
//line /usr/local/go/src/internal/saferio/io.go:40
		}
//line /usr/local/go/src/internal/saferio/io.go:40
		// _ = "end of CoverTab[83683]"
//line /usr/local/go/src/internal/saferio/io.go:40
		_go_fuzz_dep_.CoverTab[83684]++
								return buf, nil
//line /usr/local/go/src/internal/saferio/io.go:41
		// _ = "end of CoverTab[83684]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:42
		_go_fuzz_dep_.CoverTab[83687]++
//line /usr/local/go/src/internal/saferio/io.go:42
		// _ = "end of CoverTab[83687]"
//line /usr/local/go/src/internal/saferio/io.go:42
	}
//line /usr/local/go/src/internal/saferio/io.go:42
	// _ = "end of CoverTab[83677]"
//line /usr/local/go/src/internal/saferio/io.go:42
	_go_fuzz_dep_.CoverTab[83678]++

							var buf []byte
							buf1 := make([]byte, chunk)
							for n > 0 {
//line /usr/local/go/src/internal/saferio/io.go:46
		_go_fuzz_dep_.CoverTab[83688]++
								next := n
								if next > chunk {
//line /usr/local/go/src/internal/saferio/io.go:48
			_go_fuzz_dep_.CoverTab[83691]++
									next = chunk
//line /usr/local/go/src/internal/saferio/io.go:49
			// _ = "end of CoverTab[83691]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:50
			_go_fuzz_dep_.CoverTab[83692]++
//line /usr/local/go/src/internal/saferio/io.go:50
			// _ = "end of CoverTab[83692]"
//line /usr/local/go/src/internal/saferio/io.go:50
		}
//line /usr/local/go/src/internal/saferio/io.go:50
		// _ = "end of CoverTab[83688]"
//line /usr/local/go/src/internal/saferio/io.go:50
		_go_fuzz_dep_.CoverTab[83689]++
								_, err := io.ReadFull(r, buf1[:next])
								if err != nil {
//line /usr/local/go/src/internal/saferio/io.go:52
			_go_fuzz_dep_.CoverTab[83693]++
									if len(buf) > 0 && func() bool {
//line /usr/local/go/src/internal/saferio/io.go:53
				_go_fuzz_dep_.CoverTab[83695]++
//line /usr/local/go/src/internal/saferio/io.go:53
				return err == io.EOF
//line /usr/local/go/src/internal/saferio/io.go:53
				// _ = "end of CoverTab[83695]"
//line /usr/local/go/src/internal/saferio/io.go:53
			}() {
//line /usr/local/go/src/internal/saferio/io.go:53
				_go_fuzz_dep_.CoverTab[83696]++
										err = io.ErrUnexpectedEOF
//line /usr/local/go/src/internal/saferio/io.go:54
				// _ = "end of CoverTab[83696]"
			} else {
//line /usr/local/go/src/internal/saferio/io.go:55
				_go_fuzz_dep_.CoverTab[83697]++
//line /usr/local/go/src/internal/saferio/io.go:55
				// _ = "end of CoverTab[83697]"
//line /usr/local/go/src/internal/saferio/io.go:55
			}
//line /usr/local/go/src/internal/saferio/io.go:55
			// _ = "end of CoverTab[83693]"
//line /usr/local/go/src/internal/saferio/io.go:55
			_go_fuzz_dep_.CoverTab[83694]++
									return nil, err
//line /usr/local/go/src/internal/saferio/io.go:56
			// _ = "end of CoverTab[83694]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:57
			_go_fuzz_dep_.CoverTab[83698]++
//line /usr/local/go/src/internal/saferio/io.go:57
			// _ = "end of CoverTab[83698]"
//line /usr/local/go/src/internal/saferio/io.go:57
		}
//line /usr/local/go/src/internal/saferio/io.go:57
		// _ = "end of CoverTab[83689]"
//line /usr/local/go/src/internal/saferio/io.go:57
		_go_fuzz_dep_.CoverTab[83690]++
								buf = append(buf, buf1[:next]...)
								n -= next
//line /usr/local/go/src/internal/saferio/io.go:59
		// _ = "end of CoverTab[83690]"
	}
//line /usr/local/go/src/internal/saferio/io.go:60
	// _ = "end of CoverTab[83678]"
//line /usr/local/go/src/internal/saferio/io.go:60
	_go_fuzz_dep_.CoverTab[83679]++
							return buf, nil
//line /usr/local/go/src/internal/saferio/io.go:61
	// _ = "end of CoverTab[83679]"
}

// ReadDataAt reads n bytes from the input stream at off, but avoids
//line /usr/local/go/src/internal/saferio/io.go:64
// allocating all n bytes if n is large. This avoids crashing the program
//line /usr/local/go/src/internal/saferio/io.go:64
// by allocating all n bytes in cases where n is incorrect.
//line /usr/local/go/src/internal/saferio/io.go:67
func ReadDataAt(r io.ReaderAt, n uint64, off int64) ([]byte, error) {
//line /usr/local/go/src/internal/saferio/io.go:67
	_go_fuzz_dep_.CoverTab[83699]++
							if int64(n) < 0 || func() bool {
//line /usr/local/go/src/internal/saferio/io.go:68
		_go_fuzz_dep_.CoverTab[83703]++
//line /usr/local/go/src/internal/saferio/io.go:68
		return n != uint64(int(n))
//line /usr/local/go/src/internal/saferio/io.go:68
		// _ = "end of CoverTab[83703]"
//line /usr/local/go/src/internal/saferio/io.go:68
	}() {
//line /usr/local/go/src/internal/saferio/io.go:68
		_go_fuzz_dep_.CoverTab[83704]++

//line /usr/local/go/src/internal/saferio/io.go:71
		return nil, io.ErrUnexpectedEOF
//line /usr/local/go/src/internal/saferio/io.go:71
		// _ = "end of CoverTab[83704]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:72
		_go_fuzz_dep_.CoverTab[83705]++
//line /usr/local/go/src/internal/saferio/io.go:72
		// _ = "end of CoverTab[83705]"
//line /usr/local/go/src/internal/saferio/io.go:72
	}
//line /usr/local/go/src/internal/saferio/io.go:72
	// _ = "end of CoverTab[83699]"
//line /usr/local/go/src/internal/saferio/io.go:72
	_go_fuzz_dep_.CoverTab[83700]++

							if n < chunk {
//line /usr/local/go/src/internal/saferio/io.go:74
		_go_fuzz_dep_.CoverTab[83706]++
								buf := make([]byte, n)
								_, err := r.ReadAt(buf, off)
								if err != nil {
//line /usr/local/go/src/internal/saferio/io.go:77
			_go_fuzz_dep_.CoverTab[83708]++

//line /usr/local/go/src/internal/saferio/io.go:80
			if err != io.EOF || func() bool {
//line /usr/local/go/src/internal/saferio/io.go:80
				_go_fuzz_dep_.CoverTab[83709]++
//line /usr/local/go/src/internal/saferio/io.go:80
				return n > 0
//line /usr/local/go/src/internal/saferio/io.go:80
				// _ = "end of CoverTab[83709]"
//line /usr/local/go/src/internal/saferio/io.go:80
			}() {
//line /usr/local/go/src/internal/saferio/io.go:80
				_go_fuzz_dep_.CoverTab[83710]++
										return nil, err
//line /usr/local/go/src/internal/saferio/io.go:81
				// _ = "end of CoverTab[83710]"
			} else {
//line /usr/local/go/src/internal/saferio/io.go:82
				_go_fuzz_dep_.CoverTab[83711]++
//line /usr/local/go/src/internal/saferio/io.go:82
				// _ = "end of CoverTab[83711]"
//line /usr/local/go/src/internal/saferio/io.go:82
			}
//line /usr/local/go/src/internal/saferio/io.go:82
			// _ = "end of CoverTab[83708]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:83
			_go_fuzz_dep_.CoverTab[83712]++
//line /usr/local/go/src/internal/saferio/io.go:83
			// _ = "end of CoverTab[83712]"
//line /usr/local/go/src/internal/saferio/io.go:83
		}
//line /usr/local/go/src/internal/saferio/io.go:83
		// _ = "end of CoverTab[83706]"
//line /usr/local/go/src/internal/saferio/io.go:83
		_go_fuzz_dep_.CoverTab[83707]++
								return buf, nil
//line /usr/local/go/src/internal/saferio/io.go:84
		// _ = "end of CoverTab[83707]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:85
		_go_fuzz_dep_.CoverTab[83713]++
//line /usr/local/go/src/internal/saferio/io.go:85
		// _ = "end of CoverTab[83713]"
//line /usr/local/go/src/internal/saferio/io.go:85
	}
//line /usr/local/go/src/internal/saferio/io.go:85
	// _ = "end of CoverTab[83700]"
//line /usr/local/go/src/internal/saferio/io.go:85
	_go_fuzz_dep_.CoverTab[83701]++

							var buf []byte
							buf1 := make([]byte, chunk)
							for n > 0 {
//line /usr/local/go/src/internal/saferio/io.go:89
		_go_fuzz_dep_.CoverTab[83714]++
								next := n
								if next > chunk {
//line /usr/local/go/src/internal/saferio/io.go:91
			_go_fuzz_dep_.CoverTab[83717]++
									next = chunk
//line /usr/local/go/src/internal/saferio/io.go:92
			// _ = "end of CoverTab[83717]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:93
			_go_fuzz_dep_.CoverTab[83718]++
//line /usr/local/go/src/internal/saferio/io.go:93
			// _ = "end of CoverTab[83718]"
//line /usr/local/go/src/internal/saferio/io.go:93
		}
//line /usr/local/go/src/internal/saferio/io.go:93
		// _ = "end of CoverTab[83714]"
//line /usr/local/go/src/internal/saferio/io.go:93
		_go_fuzz_dep_.CoverTab[83715]++
								_, err := r.ReadAt(buf1[:next], off)
								if err != nil {
//line /usr/local/go/src/internal/saferio/io.go:95
			_go_fuzz_dep_.CoverTab[83719]++
									return nil, err
//line /usr/local/go/src/internal/saferio/io.go:96
			// _ = "end of CoverTab[83719]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:97
			_go_fuzz_dep_.CoverTab[83720]++
//line /usr/local/go/src/internal/saferio/io.go:97
			// _ = "end of CoverTab[83720]"
//line /usr/local/go/src/internal/saferio/io.go:97
		}
//line /usr/local/go/src/internal/saferio/io.go:97
		// _ = "end of CoverTab[83715]"
//line /usr/local/go/src/internal/saferio/io.go:97
		_go_fuzz_dep_.CoverTab[83716]++
								buf = append(buf, buf1[:next]...)
								n -= next
								off += int64(next)
//line /usr/local/go/src/internal/saferio/io.go:100
		// _ = "end of CoverTab[83716]"
	}
//line /usr/local/go/src/internal/saferio/io.go:101
	// _ = "end of CoverTab[83701]"
//line /usr/local/go/src/internal/saferio/io.go:101
	_go_fuzz_dep_.CoverTab[83702]++
							return buf, nil
//line /usr/local/go/src/internal/saferio/io.go:102
	// _ = "end of CoverTab[83702]"
}

// SliceCap returns the capacity to use when allocating a slice.
//line /usr/local/go/src/internal/saferio/io.go:105
// After the slice is allocated with the capacity, it should be
//line /usr/local/go/src/internal/saferio/io.go:105
// built using append. This will avoid allocating too much memory
//line /usr/local/go/src/internal/saferio/io.go:105
// if the capacity is large and incorrect.
//line /usr/local/go/src/internal/saferio/io.go:105
//
//line /usr/local/go/src/internal/saferio/io.go:105
// A negative result means that the value is always too big.
//line /usr/local/go/src/internal/saferio/io.go:105
//
//line /usr/local/go/src/internal/saferio/io.go:105
// The element type is described by passing a pointer to a value of that type.
//line /usr/local/go/src/internal/saferio/io.go:105
// This would ideally use generics, but this code is built with
//line /usr/local/go/src/internal/saferio/io.go:105
// the bootstrap compiler which need not support generics.
//line /usr/local/go/src/internal/saferio/io.go:105
// We use a pointer so that we can handle slices of interface type.
//line /usr/local/go/src/internal/saferio/io.go:116
func SliceCap(v any, c uint64) int {
//line /usr/local/go/src/internal/saferio/io.go:116
	_go_fuzz_dep_.CoverTab[83721]++
							if int64(c) < 0 || func() bool {
//line /usr/local/go/src/internal/saferio/io.go:117
		_go_fuzz_dep_.CoverTab[83726]++
//line /usr/local/go/src/internal/saferio/io.go:117
		return c != uint64(int(c))
//line /usr/local/go/src/internal/saferio/io.go:117
		// _ = "end of CoverTab[83726]"
//line /usr/local/go/src/internal/saferio/io.go:117
	}() {
//line /usr/local/go/src/internal/saferio/io.go:117
		_go_fuzz_dep_.CoverTab[83727]++
								return -1
//line /usr/local/go/src/internal/saferio/io.go:118
		// _ = "end of CoverTab[83727]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:119
		_go_fuzz_dep_.CoverTab[83728]++
//line /usr/local/go/src/internal/saferio/io.go:119
		// _ = "end of CoverTab[83728]"
//line /usr/local/go/src/internal/saferio/io.go:119
	}
//line /usr/local/go/src/internal/saferio/io.go:119
	// _ = "end of CoverTab[83721]"
//line /usr/local/go/src/internal/saferio/io.go:119
	_go_fuzz_dep_.CoverTab[83722]++
							typ := reflect.TypeOf(v)
							if typ.Kind() != reflect.Ptr {
//line /usr/local/go/src/internal/saferio/io.go:121
		_go_fuzz_dep_.CoverTab[83729]++
								panic("SliceCap called with non-pointer type")
//line /usr/local/go/src/internal/saferio/io.go:122
		// _ = "end of CoverTab[83729]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:123
		_go_fuzz_dep_.CoverTab[83730]++
//line /usr/local/go/src/internal/saferio/io.go:123
		// _ = "end of CoverTab[83730]"
//line /usr/local/go/src/internal/saferio/io.go:123
	}
//line /usr/local/go/src/internal/saferio/io.go:123
	// _ = "end of CoverTab[83722]"
//line /usr/local/go/src/internal/saferio/io.go:123
	_go_fuzz_dep_.CoverTab[83723]++
							size := uint64(typ.Elem().Size())
							if size > 0 && func() bool {
//line /usr/local/go/src/internal/saferio/io.go:125
		_go_fuzz_dep_.CoverTab[83731]++
//line /usr/local/go/src/internal/saferio/io.go:125
		return c > (1<<64-1)/size
//line /usr/local/go/src/internal/saferio/io.go:125
		// _ = "end of CoverTab[83731]"
//line /usr/local/go/src/internal/saferio/io.go:125
	}() {
//line /usr/local/go/src/internal/saferio/io.go:125
		_go_fuzz_dep_.CoverTab[83732]++
								return -1
//line /usr/local/go/src/internal/saferio/io.go:126
		// _ = "end of CoverTab[83732]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:127
		_go_fuzz_dep_.CoverTab[83733]++
//line /usr/local/go/src/internal/saferio/io.go:127
		// _ = "end of CoverTab[83733]"
//line /usr/local/go/src/internal/saferio/io.go:127
	}
//line /usr/local/go/src/internal/saferio/io.go:127
	// _ = "end of CoverTab[83723]"
//line /usr/local/go/src/internal/saferio/io.go:127
	_go_fuzz_dep_.CoverTab[83724]++
							if c*size > chunk {
//line /usr/local/go/src/internal/saferio/io.go:128
		_go_fuzz_dep_.CoverTab[83734]++
								c = uint64(chunk / size)
								if c == 0 {
//line /usr/local/go/src/internal/saferio/io.go:130
			_go_fuzz_dep_.CoverTab[83735]++
									c = 1
//line /usr/local/go/src/internal/saferio/io.go:131
			// _ = "end of CoverTab[83735]"
		} else {
//line /usr/local/go/src/internal/saferio/io.go:132
			_go_fuzz_dep_.CoverTab[83736]++
//line /usr/local/go/src/internal/saferio/io.go:132
			// _ = "end of CoverTab[83736]"
//line /usr/local/go/src/internal/saferio/io.go:132
		}
//line /usr/local/go/src/internal/saferio/io.go:132
		// _ = "end of CoverTab[83734]"
	} else {
//line /usr/local/go/src/internal/saferio/io.go:133
		_go_fuzz_dep_.CoverTab[83737]++
//line /usr/local/go/src/internal/saferio/io.go:133
		// _ = "end of CoverTab[83737]"
//line /usr/local/go/src/internal/saferio/io.go:133
	}
//line /usr/local/go/src/internal/saferio/io.go:133
	// _ = "end of CoverTab[83724]"
//line /usr/local/go/src/internal/saferio/io.go:133
	_go_fuzz_dep_.CoverTab[83725]++
							return int(c)
//line /usr/local/go/src/internal/saferio/io.go:134
	// _ = "end of CoverTab[83725]"
}

//line /usr/local/go/src/internal/saferio/io.go:135
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/saferio/io.go:135
var _ = _go_fuzz_dep_.CoverTab
