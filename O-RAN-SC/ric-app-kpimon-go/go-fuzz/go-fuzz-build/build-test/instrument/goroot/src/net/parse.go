// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple file i/o and string manipulation, to avoid
// depending on strconv and bufio and strings.

//line /snap/go/10455/src/net/parse.go:8
package net

//line /snap/go/10455/src/net/parse.go:8
import (
//line /snap/go/10455/src/net/parse.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/net/parse.go:8
)
//line /snap/go/10455/src/net/parse.go:8
import (
//line /snap/go/10455/src/net/parse.go:8
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/net/parse.go:8
)

import (
	"internal/bytealg"
	"io"
	"os"
	"time"
)

type file struct {
	file	*os.File
	data	[]byte
	atEOF	bool
}

func (f *file) close() {
//line /snap/go/10455/src/net/parse.go:23
	_go_fuzz_dep_.CoverTab[7655]++
//line /snap/go/10455/src/net/parse.go:23
	f.file.Close()
//line /snap/go/10455/src/net/parse.go:23
	// _ = "end of CoverTab[7655]"
//line /snap/go/10455/src/net/parse.go:23
}

func (f *file) getLineFromData() (s string, ok bool) {
//line /snap/go/10455/src/net/parse.go:25
	_go_fuzz_dep_.CoverTab[7656]++
						data := f.data
						i := 0
//line /snap/go/10455/src/net/parse.go:27
	_go_fuzz_dep_.CoverTab[786726] = 0
						for i = 0; i < len(data); i++ {
//line /snap/go/10455/src/net/parse.go:28
		if _go_fuzz_dep_.CoverTab[786726] == 0 {
//line /snap/go/10455/src/net/parse.go:28
			_go_fuzz_dep_.CoverTab[529450]++
//line /snap/go/10455/src/net/parse.go:28
		} else {
//line /snap/go/10455/src/net/parse.go:28
			_go_fuzz_dep_.CoverTab[529451]++
//line /snap/go/10455/src/net/parse.go:28
		}
//line /snap/go/10455/src/net/parse.go:28
		_go_fuzz_dep_.CoverTab[786726] = 1
//line /snap/go/10455/src/net/parse.go:28
		_go_fuzz_dep_.CoverTab[7659]++
							if data[i] == '\n' {
//line /snap/go/10455/src/net/parse.go:29
			_go_fuzz_dep_.CoverTab[529382]++
//line /snap/go/10455/src/net/parse.go:29
			_go_fuzz_dep_.CoverTab[7660]++
								s = string(data[0:i])
								ok = true

								i++
								n := len(data) - i
								copy(data[0:], data[i:])
								f.data = data[0:n]
								return
//line /snap/go/10455/src/net/parse.go:37
			// _ = "end of CoverTab[7660]"
		} else {
//line /snap/go/10455/src/net/parse.go:38
			_go_fuzz_dep_.CoverTab[529383]++
//line /snap/go/10455/src/net/parse.go:38
			_go_fuzz_dep_.CoverTab[7661]++
//line /snap/go/10455/src/net/parse.go:38
			// _ = "end of CoverTab[7661]"
//line /snap/go/10455/src/net/parse.go:38
		}
//line /snap/go/10455/src/net/parse.go:38
		// _ = "end of CoverTab[7659]"
	}
//line /snap/go/10455/src/net/parse.go:39
	if _go_fuzz_dep_.CoverTab[786726] == 0 {
//line /snap/go/10455/src/net/parse.go:39
		_go_fuzz_dep_.CoverTab[529452]++
//line /snap/go/10455/src/net/parse.go:39
	} else {
//line /snap/go/10455/src/net/parse.go:39
		_go_fuzz_dep_.CoverTab[529453]++
//line /snap/go/10455/src/net/parse.go:39
	}
//line /snap/go/10455/src/net/parse.go:39
	// _ = "end of CoverTab[7656]"
//line /snap/go/10455/src/net/parse.go:39
	_go_fuzz_dep_.CoverTab[7657]++
						if f.atEOF && func() bool {
//line /snap/go/10455/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[7662]++
//line /snap/go/10455/src/net/parse.go:40
		return len(f.data) > 0
//line /snap/go/10455/src/net/parse.go:40
		// _ = "end of CoverTab[7662]"
//line /snap/go/10455/src/net/parse.go:40
	}() {
//line /snap/go/10455/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[529384]++
//line /snap/go/10455/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[7663]++

							s = string(data)
							f.data = f.data[0:0]
							ok = true
//line /snap/go/10455/src/net/parse.go:44
		// _ = "end of CoverTab[7663]"
	} else {
//line /snap/go/10455/src/net/parse.go:45
		_go_fuzz_dep_.CoverTab[529385]++
//line /snap/go/10455/src/net/parse.go:45
		_go_fuzz_dep_.CoverTab[7664]++
//line /snap/go/10455/src/net/parse.go:45
		// _ = "end of CoverTab[7664]"
//line /snap/go/10455/src/net/parse.go:45
	}
//line /snap/go/10455/src/net/parse.go:45
	// _ = "end of CoverTab[7657]"
//line /snap/go/10455/src/net/parse.go:45
	_go_fuzz_dep_.CoverTab[7658]++
						return
//line /snap/go/10455/src/net/parse.go:46
	// _ = "end of CoverTab[7658]"
}

func (f *file) readLine() (s string, ok bool) {
//line /snap/go/10455/src/net/parse.go:49
	_go_fuzz_dep_.CoverTab[7665]++
						if s, ok = f.getLineFromData(); ok {
//line /snap/go/10455/src/net/parse.go:50
		_go_fuzz_dep_.CoverTab[529386]++
//line /snap/go/10455/src/net/parse.go:50
		_go_fuzz_dep_.CoverTab[7668]++
							return
//line /snap/go/10455/src/net/parse.go:51
		// _ = "end of CoverTab[7668]"
	} else {
//line /snap/go/10455/src/net/parse.go:52
		_go_fuzz_dep_.CoverTab[529387]++
//line /snap/go/10455/src/net/parse.go:52
		_go_fuzz_dep_.CoverTab[7669]++
//line /snap/go/10455/src/net/parse.go:52
		// _ = "end of CoverTab[7669]"
//line /snap/go/10455/src/net/parse.go:52
	}
//line /snap/go/10455/src/net/parse.go:52
	// _ = "end of CoverTab[7665]"
//line /snap/go/10455/src/net/parse.go:52
	_go_fuzz_dep_.CoverTab[7666]++
						if len(f.data) < cap(f.data) {
//line /snap/go/10455/src/net/parse.go:53
		_go_fuzz_dep_.CoverTab[529388]++
//line /snap/go/10455/src/net/parse.go:53
		_go_fuzz_dep_.CoverTab[7670]++
							ln := len(f.data)
							n, err := io.ReadFull(f.file, f.data[ln:cap(f.data)])
							if n >= 0 {
//line /snap/go/10455/src/net/parse.go:56
			_go_fuzz_dep_.CoverTab[529390]++
//line /snap/go/10455/src/net/parse.go:56
			_go_fuzz_dep_.CoverTab[7672]++
								f.data = f.data[0 : ln+n]
//line /snap/go/10455/src/net/parse.go:57
			// _ = "end of CoverTab[7672]"
		} else {
//line /snap/go/10455/src/net/parse.go:58
			_go_fuzz_dep_.CoverTab[529391]++
//line /snap/go/10455/src/net/parse.go:58
			_go_fuzz_dep_.CoverTab[7673]++
//line /snap/go/10455/src/net/parse.go:58
			// _ = "end of CoverTab[7673]"
//line /snap/go/10455/src/net/parse.go:58
		}
//line /snap/go/10455/src/net/parse.go:58
		// _ = "end of CoverTab[7670]"
//line /snap/go/10455/src/net/parse.go:58
		_go_fuzz_dep_.CoverTab[7671]++
							if err == io.EOF || func() bool {
//line /snap/go/10455/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[7674]++
//line /snap/go/10455/src/net/parse.go:59
			return err == io.ErrUnexpectedEOF
//line /snap/go/10455/src/net/parse.go:59
			// _ = "end of CoverTab[7674]"
//line /snap/go/10455/src/net/parse.go:59
		}() {
//line /snap/go/10455/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[529392]++
//line /snap/go/10455/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[7675]++
								f.atEOF = true
//line /snap/go/10455/src/net/parse.go:60
			// _ = "end of CoverTab[7675]"
		} else {
//line /snap/go/10455/src/net/parse.go:61
			_go_fuzz_dep_.CoverTab[529393]++
//line /snap/go/10455/src/net/parse.go:61
			_go_fuzz_dep_.CoverTab[7676]++
//line /snap/go/10455/src/net/parse.go:61
			// _ = "end of CoverTab[7676]"
//line /snap/go/10455/src/net/parse.go:61
		}
//line /snap/go/10455/src/net/parse.go:61
		// _ = "end of CoverTab[7671]"
	} else {
//line /snap/go/10455/src/net/parse.go:62
		_go_fuzz_dep_.CoverTab[529389]++
//line /snap/go/10455/src/net/parse.go:62
		_go_fuzz_dep_.CoverTab[7677]++
//line /snap/go/10455/src/net/parse.go:62
		// _ = "end of CoverTab[7677]"
//line /snap/go/10455/src/net/parse.go:62
	}
//line /snap/go/10455/src/net/parse.go:62
	// _ = "end of CoverTab[7666]"
//line /snap/go/10455/src/net/parse.go:62
	_go_fuzz_dep_.CoverTab[7667]++
						s, ok = f.getLineFromData()
						return
//line /snap/go/10455/src/net/parse.go:64
	// _ = "end of CoverTab[7667]"
}

func (f *file) stat() (mtime time.Time, size int64, err error) {
//line /snap/go/10455/src/net/parse.go:67
	_go_fuzz_dep_.CoverTab[7678]++
						st, err := f.file.Stat()
						if err != nil {
//line /snap/go/10455/src/net/parse.go:69
		_go_fuzz_dep_.CoverTab[529394]++
//line /snap/go/10455/src/net/parse.go:69
		_go_fuzz_dep_.CoverTab[7680]++
							return time.Time{}, 0, err
//line /snap/go/10455/src/net/parse.go:70
		// _ = "end of CoverTab[7680]"
	} else {
//line /snap/go/10455/src/net/parse.go:71
		_go_fuzz_dep_.CoverTab[529395]++
//line /snap/go/10455/src/net/parse.go:71
		_go_fuzz_dep_.CoverTab[7681]++
//line /snap/go/10455/src/net/parse.go:71
		// _ = "end of CoverTab[7681]"
//line /snap/go/10455/src/net/parse.go:71
	}
//line /snap/go/10455/src/net/parse.go:71
	// _ = "end of CoverTab[7678]"
//line /snap/go/10455/src/net/parse.go:71
	_go_fuzz_dep_.CoverTab[7679]++
						return st.ModTime(), st.Size(), nil
//line /snap/go/10455/src/net/parse.go:72
	// _ = "end of CoverTab[7679]"
}

func open(name string) (*file, error) {
//line /snap/go/10455/src/net/parse.go:75
	_go_fuzz_dep_.CoverTab[7682]++
						fd, err := os.Open(name)
						if err != nil {
//line /snap/go/10455/src/net/parse.go:77
		_go_fuzz_dep_.CoverTab[529396]++
//line /snap/go/10455/src/net/parse.go:77
		_go_fuzz_dep_.CoverTab[7684]++
							return nil, err
//line /snap/go/10455/src/net/parse.go:78
		// _ = "end of CoverTab[7684]"
	} else {
//line /snap/go/10455/src/net/parse.go:79
		_go_fuzz_dep_.CoverTab[529397]++
//line /snap/go/10455/src/net/parse.go:79
		_go_fuzz_dep_.CoverTab[7685]++
//line /snap/go/10455/src/net/parse.go:79
		// _ = "end of CoverTab[7685]"
//line /snap/go/10455/src/net/parse.go:79
	}
//line /snap/go/10455/src/net/parse.go:79
	// _ = "end of CoverTab[7682]"
//line /snap/go/10455/src/net/parse.go:79
	_go_fuzz_dep_.CoverTab[7683]++
						return &file{fd, make([]byte, 0, 64*1024), false}, nil
//line /snap/go/10455/src/net/parse.go:80
	// _ = "end of CoverTab[7683]"
}

func stat(name string) (mtime time.Time, size int64, err error) {
//line /snap/go/10455/src/net/parse.go:83
	_go_fuzz_dep_.CoverTab[7686]++
						st, err := os.Stat(name)
						if err != nil {
//line /snap/go/10455/src/net/parse.go:85
		_go_fuzz_dep_.CoverTab[529398]++
//line /snap/go/10455/src/net/parse.go:85
		_go_fuzz_dep_.CoverTab[7688]++
							return time.Time{}, 0, err
//line /snap/go/10455/src/net/parse.go:86
		// _ = "end of CoverTab[7688]"
	} else {
//line /snap/go/10455/src/net/parse.go:87
		_go_fuzz_dep_.CoverTab[529399]++
//line /snap/go/10455/src/net/parse.go:87
		_go_fuzz_dep_.CoverTab[7689]++
//line /snap/go/10455/src/net/parse.go:87
		// _ = "end of CoverTab[7689]"
//line /snap/go/10455/src/net/parse.go:87
	}
//line /snap/go/10455/src/net/parse.go:87
	// _ = "end of CoverTab[7686]"
//line /snap/go/10455/src/net/parse.go:87
	_go_fuzz_dep_.CoverTab[7687]++
						return st.ModTime(), st.Size(), nil
//line /snap/go/10455/src/net/parse.go:88
	// _ = "end of CoverTab[7687]"
}

// Count occurrences in s of any bytes in t.
func countAnyByte(s string, t string) int {
//line /snap/go/10455/src/net/parse.go:92
	_go_fuzz_dep_.CoverTab[7690]++
						n := 0
//line /snap/go/10455/src/net/parse.go:93
	_go_fuzz_dep_.CoverTab[786727] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/parse.go:94
		if _go_fuzz_dep_.CoverTab[786727] == 0 {
//line /snap/go/10455/src/net/parse.go:94
			_go_fuzz_dep_.CoverTab[529454]++
//line /snap/go/10455/src/net/parse.go:94
		} else {
//line /snap/go/10455/src/net/parse.go:94
			_go_fuzz_dep_.CoverTab[529455]++
//line /snap/go/10455/src/net/parse.go:94
		}
//line /snap/go/10455/src/net/parse.go:94
		_go_fuzz_dep_.CoverTab[786727] = 1
//line /snap/go/10455/src/net/parse.go:94
		_go_fuzz_dep_.CoverTab[7692]++
							if bytealg.IndexByteString(t, s[i]) >= 0 {
//line /snap/go/10455/src/net/parse.go:95
			_go_fuzz_dep_.CoverTab[529400]++
//line /snap/go/10455/src/net/parse.go:95
			_go_fuzz_dep_.CoverTab[7693]++
								n++
//line /snap/go/10455/src/net/parse.go:96
			// _ = "end of CoverTab[7693]"
		} else {
//line /snap/go/10455/src/net/parse.go:97
			_go_fuzz_dep_.CoverTab[529401]++
//line /snap/go/10455/src/net/parse.go:97
			_go_fuzz_dep_.CoverTab[7694]++
//line /snap/go/10455/src/net/parse.go:97
			// _ = "end of CoverTab[7694]"
//line /snap/go/10455/src/net/parse.go:97
		}
//line /snap/go/10455/src/net/parse.go:97
		// _ = "end of CoverTab[7692]"
	}
//line /snap/go/10455/src/net/parse.go:98
	if _go_fuzz_dep_.CoverTab[786727] == 0 {
//line /snap/go/10455/src/net/parse.go:98
		_go_fuzz_dep_.CoverTab[529456]++
//line /snap/go/10455/src/net/parse.go:98
	} else {
//line /snap/go/10455/src/net/parse.go:98
		_go_fuzz_dep_.CoverTab[529457]++
//line /snap/go/10455/src/net/parse.go:98
	}
//line /snap/go/10455/src/net/parse.go:98
	// _ = "end of CoverTab[7690]"
//line /snap/go/10455/src/net/parse.go:98
	_go_fuzz_dep_.CoverTab[7691]++
						return n
//line /snap/go/10455/src/net/parse.go:99
	// _ = "end of CoverTab[7691]"
}

// Split s at any bytes in t.
func splitAtBytes(s string, t string) []string {
//line /snap/go/10455/src/net/parse.go:103
	_go_fuzz_dep_.CoverTab[7695]++
						a := make([]string, 1+countAnyByte(s, t))
						n := 0
						last := 0
//line /snap/go/10455/src/net/parse.go:106
	_go_fuzz_dep_.CoverTab[786728] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/parse.go:107
		if _go_fuzz_dep_.CoverTab[786728] == 0 {
//line /snap/go/10455/src/net/parse.go:107
			_go_fuzz_dep_.CoverTab[529458]++
//line /snap/go/10455/src/net/parse.go:107
		} else {
//line /snap/go/10455/src/net/parse.go:107
			_go_fuzz_dep_.CoverTab[529459]++
//line /snap/go/10455/src/net/parse.go:107
		}
//line /snap/go/10455/src/net/parse.go:107
		_go_fuzz_dep_.CoverTab[786728] = 1
//line /snap/go/10455/src/net/parse.go:107
		_go_fuzz_dep_.CoverTab[7698]++
							if bytealg.IndexByteString(t, s[i]) >= 0 {
//line /snap/go/10455/src/net/parse.go:108
			_go_fuzz_dep_.CoverTab[529402]++
//line /snap/go/10455/src/net/parse.go:108
			_go_fuzz_dep_.CoverTab[7699]++
								if last < i {
//line /snap/go/10455/src/net/parse.go:109
				_go_fuzz_dep_.CoverTab[529404]++
//line /snap/go/10455/src/net/parse.go:109
				_go_fuzz_dep_.CoverTab[7701]++
									a[n] = s[last:i]
									n++
//line /snap/go/10455/src/net/parse.go:111
				// _ = "end of CoverTab[7701]"
			} else {
//line /snap/go/10455/src/net/parse.go:112
				_go_fuzz_dep_.CoverTab[529405]++
//line /snap/go/10455/src/net/parse.go:112
				_go_fuzz_dep_.CoverTab[7702]++
//line /snap/go/10455/src/net/parse.go:112
				// _ = "end of CoverTab[7702]"
//line /snap/go/10455/src/net/parse.go:112
			}
//line /snap/go/10455/src/net/parse.go:112
			// _ = "end of CoverTab[7699]"
//line /snap/go/10455/src/net/parse.go:112
			_go_fuzz_dep_.CoverTab[7700]++
								last = i + 1
//line /snap/go/10455/src/net/parse.go:113
			// _ = "end of CoverTab[7700]"
		} else {
//line /snap/go/10455/src/net/parse.go:114
			_go_fuzz_dep_.CoverTab[529403]++
//line /snap/go/10455/src/net/parse.go:114
			_go_fuzz_dep_.CoverTab[7703]++
//line /snap/go/10455/src/net/parse.go:114
			// _ = "end of CoverTab[7703]"
//line /snap/go/10455/src/net/parse.go:114
		}
//line /snap/go/10455/src/net/parse.go:114
		// _ = "end of CoverTab[7698]"
	}
//line /snap/go/10455/src/net/parse.go:115
	if _go_fuzz_dep_.CoverTab[786728] == 0 {
//line /snap/go/10455/src/net/parse.go:115
		_go_fuzz_dep_.CoverTab[529460]++
//line /snap/go/10455/src/net/parse.go:115
	} else {
//line /snap/go/10455/src/net/parse.go:115
		_go_fuzz_dep_.CoverTab[529461]++
//line /snap/go/10455/src/net/parse.go:115
	}
//line /snap/go/10455/src/net/parse.go:115
	// _ = "end of CoverTab[7695]"
//line /snap/go/10455/src/net/parse.go:115
	_go_fuzz_dep_.CoverTab[7696]++
						if last < len(s) {
//line /snap/go/10455/src/net/parse.go:116
		_go_fuzz_dep_.CoverTab[529406]++
//line /snap/go/10455/src/net/parse.go:116
		_go_fuzz_dep_.CoverTab[7704]++
							a[n] = s[last:]
							n++
//line /snap/go/10455/src/net/parse.go:118
		// _ = "end of CoverTab[7704]"
	} else {
//line /snap/go/10455/src/net/parse.go:119
		_go_fuzz_dep_.CoverTab[529407]++
//line /snap/go/10455/src/net/parse.go:119
		_go_fuzz_dep_.CoverTab[7705]++
//line /snap/go/10455/src/net/parse.go:119
		// _ = "end of CoverTab[7705]"
//line /snap/go/10455/src/net/parse.go:119
	}
//line /snap/go/10455/src/net/parse.go:119
	// _ = "end of CoverTab[7696]"
//line /snap/go/10455/src/net/parse.go:119
	_go_fuzz_dep_.CoverTab[7697]++
						return a[0:n]
//line /snap/go/10455/src/net/parse.go:120
	// _ = "end of CoverTab[7697]"
}

func getFields(s string) []string {
//line /snap/go/10455/src/net/parse.go:123
	_go_fuzz_dep_.CoverTab[7706]++
//line /snap/go/10455/src/net/parse.go:123
	return splitAtBytes(s, " \r\t\n")
//line /snap/go/10455/src/net/parse.go:123
	// _ = "end of CoverTab[7706]"
//line /snap/go/10455/src/net/parse.go:123
}

// Bigger than we need, not too big to worry about overflow
const big = 0xFFFFFF

// Decimal to integer.
//line /snap/go/10455/src/net/parse.go:128
// Returns number, characters consumed, success.
//line /snap/go/10455/src/net/parse.go:130
func dtoi(s string) (n int, i int, ok bool) {
//line /snap/go/10455/src/net/parse.go:130
	_go_fuzz_dep_.CoverTab[7707]++
						n = 0
//line /snap/go/10455/src/net/parse.go:131
	_go_fuzz_dep_.CoverTab[786729] = 0
						for i = 0; i < len(s) && func() bool {
//line /snap/go/10455/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[7710]++
//line /snap/go/10455/src/net/parse.go:132
		return '0' <= s[i]
//line /snap/go/10455/src/net/parse.go:132
		// _ = "end of CoverTab[7710]"
//line /snap/go/10455/src/net/parse.go:132
	}() && func() bool {
//line /snap/go/10455/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[7711]++
//line /snap/go/10455/src/net/parse.go:132
		return s[i] <= '9'
//line /snap/go/10455/src/net/parse.go:132
		// _ = "end of CoverTab[7711]"
//line /snap/go/10455/src/net/parse.go:132
	}(); i++ {
//line /snap/go/10455/src/net/parse.go:132
		if _go_fuzz_dep_.CoverTab[786729] == 0 {
//line /snap/go/10455/src/net/parse.go:132
			_go_fuzz_dep_.CoverTab[529462]++
//line /snap/go/10455/src/net/parse.go:132
		} else {
//line /snap/go/10455/src/net/parse.go:132
			_go_fuzz_dep_.CoverTab[529463]++
//line /snap/go/10455/src/net/parse.go:132
		}
//line /snap/go/10455/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[786729] = 1
//line /snap/go/10455/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[7712]++
							n = n*10 + int(s[i]-'0')
							if n >= big {
//line /snap/go/10455/src/net/parse.go:134
			_go_fuzz_dep_.CoverTab[529408]++
//line /snap/go/10455/src/net/parse.go:134
			_go_fuzz_dep_.CoverTab[7713]++
								return big, i, false
//line /snap/go/10455/src/net/parse.go:135
			// _ = "end of CoverTab[7713]"
		} else {
//line /snap/go/10455/src/net/parse.go:136
			_go_fuzz_dep_.CoverTab[529409]++
//line /snap/go/10455/src/net/parse.go:136
			_go_fuzz_dep_.CoverTab[7714]++
//line /snap/go/10455/src/net/parse.go:136
			// _ = "end of CoverTab[7714]"
//line /snap/go/10455/src/net/parse.go:136
		}
//line /snap/go/10455/src/net/parse.go:136
		// _ = "end of CoverTab[7712]"
	}
//line /snap/go/10455/src/net/parse.go:137
	if _go_fuzz_dep_.CoverTab[786729] == 0 {
//line /snap/go/10455/src/net/parse.go:137
		_go_fuzz_dep_.CoverTab[529464]++
//line /snap/go/10455/src/net/parse.go:137
	} else {
//line /snap/go/10455/src/net/parse.go:137
		_go_fuzz_dep_.CoverTab[529465]++
//line /snap/go/10455/src/net/parse.go:137
	}
//line /snap/go/10455/src/net/parse.go:137
	// _ = "end of CoverTab[7707]"
//line /snap/go/10455/src/net/parse.go:137
	_go_fuzz_dep_.CoverTab[7708]++
						if i == 0 {
//line /snap/go/10455/src/net/parse.go:138
		_go_fuzz_dep_.CoverTab[529410]++
//line /snap/go/10455/src/net/parse.go:138
		_go_fuzz_dep_.CoverTab[7715]++
							return 0, 0, false
//line /snap/go/10455/src/net/parse.go:139
		// _ = "end of CoverTab[7715]"
	} else {
//line /snap/go/10455/src/net/parse.go:140
		_go_fuzz_dep_.CoverTab[529411]++
//line /snap/go/10455/src/net/parse.go:140
		_go_fuzz_dep_.CoverTab[7716]++
//line /snap/go/10455/src/net/parse.go:140
		// _ = "end of CoverTab[7716]"
//line /snap/go/10455/src/net/parse.go:140
	}
//line /snap/go/10455/src/net/parse.go:140
	// _ = "end of CoverTab[7708]"
//line /snap/go/10455/src/net/parse.go:140
	_go_fuzz_dep_.CoverTab[7709]++
						return n, i, true
//line /snap/go/10455/src/net/parse.go:141
	// _ = "end of CoverTab[7709]"
}

// Hexadecimal to integer.
//line /snap/go/10455/src/net/parse.go:144
// Returns number, characters consumed, success.
//line /snap/go/10455/src/net/parse.go:146
func xtoi(s string) (n int, i int, ok bool) {
//line /snap/go/10455/src/net/parse.go:146
	_go_fuzz_dep_.CoverTab[7717]++
						n = 0
//line /snap/go/10455/src/net/parse.go:147
	_go_fuzz_dep_.CoverTab[786730] = 0
						for i = 0; i < len(s); i++ {
//line /snap/go/10455/src/net/parse.go:148
		if _go_fuzz_dep_.CoverTab[786730] == 0 {
//line /snap/go/10455/src/net/parse.go:148
			_go_fuzz_dep_.CoverTab[529466]++
//line /snap/go/10455/src/net/parse.go:148
		} else {
//line /snap/go/10455/src/net/parse.go:148
			_go_fuzz_dep_.CoverTab[529467]++
//line /snap/go/10455/src/net/parse.go:148
		}
//line /snap/go/10455/src/net/parse.go:148
		_go_fuzz_dep_.CoverTab[786730] = 1
//line /snap/go/10455/src/net/parse.go:148
		_go_fuzz_dep_.CoverTab[7720]++
							if '0' <= s[i] && func() bool {
//line /snap/go/10455/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[7722]++
//line /snap/go/10455/src/net/parse.go:149
			return s[i] <= '9'
//line /snap/go/10455/src/net/parse.go:149
			// _ = "end of CoverTab[7722]"
//line /snap/go/10455/src/net/parse.go:149
		}() {
//line /snap/go/10455/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[529412]++
//line /snap/go/10455/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[7723]++
								n *= 16
								n += int(s[i] - '0')
//line /snap/go/10455/src/net/parse.go:151
			// _ = "end of CoverTab[7723]"
		} else {
//line /snap/go/10455/src/net/parse.go:152
			_go_fuzz_dep_.CoverTab[529413]++
//line /snap/go/10455/src/net/parse.go:152
			_go_fuzz_dep_.CoverTab[7724]++
//line /snap/go/10455/src/net/parse.go:152
			if 'a' <= s[i] && func() bool {
//line /snap/go/10455/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[7725]++
//line /snap/go/10455/src/net/parse.go:152
				return s[i] <= 'f'
//line /snap/go/10455/src/net/parse.go:152
				// _ = "end of CoverTab[7725]"
//line /snap/go/10455/src/net/parse.go:152
			}() {
//line /snap/go/10455/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[529414]++
//line /snap/go/10455/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[7726]++
									n *= 16
									n += int(s[i]-'a') + 10
//line /snap/go/10455/src/net/parse.go:154
				// _ = "end of CoverTab[7726]"
			} else {
//line /snap/go/10455/src/net/parse.go:155
				_go_fuzz_dep_.CoverTab[529415]++
//line /snap/go/10455/src/net/parse.go:155
				_go_fuzz_dep_.CoverTab[7727]++
//line /snap/go/10455/src/net/parse.go:155
				if 'A' <= s[i] && func() bool {
//line /snap/go/10455/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[7728]++
//line /snap/go/10455/src/net/parse.go:155
					return s[i] <= 'F'
//line /snap/go/10455/src/net/parse.go:155
					// _ = "end of CoverTab[7728]"
//line /snap/go/10455/src/net/parse.go:155
				}() {
//line /snap/go/10455/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[529416]++
//line /snap/go/10455/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[7729]++
										n *= 16
										n += int(s[i]-'A') + 10
//line /snap/go/10455/src/net/parse.go:157
					// _ = "end of CoverTab[7729]"
				} else {
//line /snap/go/10455/src/net/parse.go:158
					_go_fuzz_dep_.CoverTab[529417]++
//line /snap/go/10455/src/net/parse.go:158
					_go_fuzz_dep_.CoverTab[7730]++
										break
//line /snap/go/10455/src/net/parse.go:159
					// _ = "end of CoverTab[7730]"
				}
//line /snap/go/10455/src/net/parse.go:160
				// _ = "end of CoverTab[7727]"
//line /snap/go/10455/src/net/parse.go:160
			}
//line /snap/go/10455/src/net/parse.go:160
			// _ = "end of CoverTab[7724]"
//line /snap/go/10455/src/net/parse.go:160
		}
//line /snap/go/10455/src/net/parse.go:160
		// _ = "end of CoverTab[7720]"
//line /snap/go/10455/src/net/parse.go:160
		_go_fuzz_dep_.CoverTab[7721]++
							if n >= big {
//line /snap/go/10455/src/net/parse.go:161
			_go_fuzz_dep_.CoverTab[529418]++
//line /snap/go/10455/src/net/parse.go:161
			_go_fuzz_dep_.CoverTab[7731]++
								return 0, i, false
//line /snap/go/10455/src/net/parse.go:162
			// _ = "end of CoverTab[7731]"
		} else {
//line /snap/go/10455/src/net/parse.go:163
			_go_fuzz_dep_.CoverTab[529419]++
//line /snap/go/10455/src/net/parse.go:163
			_go_fuzz_dep_.CoverTab[7732]++
//line /snap/go/10455/src/net/parse.go:163
			// _ = "end of CoverTab[7732]"
//line /snap/go/10455/src/net/parse.go:163
		}
//line /snap/go/10455/src/net/parse.go:163
		// _ = "end of CoverTab[7721]"
	}
//line /snap/go/10455/src/net/parse.go:164
	if _go_fuzz_dep_.CoverTab[786730] == 0 {
//line /snap/go/10455/src/net/parse.go:164
		_go_fuzz_dep_.CoverTab[529468]++
//line /snap/go/10455/src/net/parse.go:164
	} else {
//line /snap/go/10455/src/net/parse.go:164
		_go_fuzz_dep_.CoverTab[529469]++
//line /snap/go/10455/src/net/parse.go:164
	}
//line /snap/go/10455/src/net/parse.go:164
	// _ = "end of CoverTab[7717]"
//line /snap/go/10455/src/net/parse.go:164
	_go_fuzz_dep_.CoverTab[7718]++
						if i == 0 {
//line /snap/go/10455/src/net/parse.go:165
		_go_fuzz_dep_.CoverTab[529420]++
//line /snap/go/10455/src/net/parse.go:165
		_go_fuzz_dep_.CoverTab[7733]++
							return 0, i, false
//line /snap/go/10455/src/net/parse.go:166
		// _ = "end of CoverTab[7733]"
	} else {
//line /snap/go/10455/src/net/parse.go:167
		_go_fuzz_dep_.CoverTab[529421]++
//line /snap/go/10455/src/net/parse.go:167
		_go_fuzz_dep_.CoverTab[7734]++
//line /snap/go/10455/src/net/parse.go:167
		// _ = "end of CoverTab[7734]"
//line /snap/go/10455/src/net/parse.go:167
	}
//line /snap/go/10455/src/net/parse.go:167
	// _ = "end of CoverTab[7718]"
//line /snap/go/10455/src/net/parse.go:167
	_go_fuzz_dep_.CoverTab[7719]++
						return n, i, true
//line /snap/go/10455/src/net/parse.go:168
	// _ = "end of CoverTab[7719]"
}

// xtoi2 converts the next two hex digits of s into a byte.
//line /snap/go/10455/src/net/parse.go:171
// If s is longer than 2 bytes then the third byte must be e.
//line /snap/go/10455/src/net/parse.go:171
// If the first two bytes of s are not hex digits or the third byte
//line /snap/go/10455/src/net/parse.go:171
// does not match e, false is returned.
//line /snap/go/10455/src/net/parse.go:175
func xtoi2(s string, e byte) (byte, bool) {
//line /snap/go/10455/src/net/parse.go:175
	_go_fuzz_dep_.CoverTab[7735]++
						if len(s) > 2 && func() bool {
//line /snap/go/10455/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[7737]++
//line /snap/go/10455/src/net/parse.go:176
		return s[2] != e
//line /snap/go/10455/src/net/parse.go:176
		// _ = "end of CoverTab[7737]"
//line /snap/go/10455/src/net/parse.go:176
	}() {
//line /snap/go/10455/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[529422]++
//line /snap/go/10455/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[7738]++
							return 0, false
//line /snap/go/10455/src/net/parse.go:177
		// _ = "end of CoverTab[7738]"
	} else {
//line /snap/go/10455/src/net/parse.go:178
		_go_fuzz_dep_.CoverTab[529423]++
//line /snap/go/10455/src/net/parse.go:178
		_go_fuzz_dep_.CoverTab[7739]++
//line /snap/go/10455/src/net/parse.go:178
		// _ = "end of CoverTab[7739]"
//line /snap/go/10455/src/net/parse.go:178
	}
//line /snap/go/10455/src/net/parse.go:178
	// _ = "end of CoverTab[7735]"
//line /snap/go/10455/src/net/parse.go:178
	_go_fuzz_dep_.CoverTab[7736]++
						n, ei, ok := xtoi(s[:2])
						return byte(n), ok && func() bool {
//line /snap/go/10455/src/net/parse.go:180
		_go_fuzz_dep_.CoverTab[7740]++
//line /snap/go/10455/src/net/parse.go:180
		return ei == 2
//line /snap/go/10455/src/net/parse.go:180
		// _ = "end of CoverTab[7740]"
//line /snap/go/10455/src/net/parse.go:180
	}()
//line /snap/go/10455/src/net/parse.go:180
	// _ = "end of CoverTab[7736]"
}

// Convert i to a hexadecimal string. Leading zeros are not printed.
func appendHex(dst []byte, i uint32) []byte {
//line /snap/go/10455/src/net/parse.go:184
	_go_fuzz_dep_.CoverTab[7741]++
						if i == 0 {
//line /snap/go/10455/src/net/parse.go:185
		_go_fuzz_dep_.CoverTab[529424]++
//line /snap/go/10455/src/net/parse.go:185
		_go_fuzz_dep_.CoverTab[7744]++
							return append(dst, '0')
//line /snap/go/10455/src/net/parse.go:186
		// _ = "end of CoverTab[7744]"
	} else {
//line /snap/go/10455/src/net/parse.go:187
		_go_fuzz_dep_.CoverTab[529425]++
//line /snap/go/10455/src/net/parse.go:187
		_go_fuzz_dep_.CoverTab[7745]++
//line /snap/go/10455/src/net/parse.go:187
		// _ = "end of CoverTab[7745]"
//line /snap/go/10455/src/net/parse.go:187
	}
//line /snap/go/10455/src/net/parse.go:187
	// _ = "end of CoverTab[7741]"
//line /snap/go/10455/src/net/parse.go:187
	_go_fuzz_dep_.CoverTab[7742]++
//line /snap/go/10455/src/net/parse.go:187
	_go_fuzz_dep_.CoverTab[786731] = 0
						for j := 7; j >= 0; j-- {
//line /snap/go/10455/src/net/parse.go:188
		if _go_fuzz_dep_.CoverTab[786731] == 0 {
//line /snap/go/10455/src/net/parse.go:188
			_go_fuzz_dep_.CoverTab[529470]++
//line /snap/go/10455/src/net/parse.go:188
		} else {
//line /snap/go/10455/src/net/parse.go:188
			_go_fuzz_dep_.CoverTab[529471]++
//line /snap/go/10455/src/net/parse.go:188
		}
//line /snap/go/10455/src/net/parse.go:188
		_go_fuzz_dep_.CoverTab[786731] = 1
//line /snap/go/10455/src/net/parse.go:188
		_go_fuzz_dep_.CoverTab[7746]++
							v := i >> uint(j*4)
							if v > 0 {
//line /snap/go/10455/src/net/parse.go:190
			_go_fuzz_dep_.CoverTab[529426]++
//line /snap/go/10455/src/net/parse.go:190
			_go_fuzz_dep_.CoverTab[7747]++
								dst = append(dst, hexDigit[v&0xf])
//line /snap/go/10455/src/net/parse.go:191
			// _ = "end of CoverTab[7747]"
		} else {
//line /snap/go/10455/src/net/parse.go:192
			_go_fuzz_dep_.CoverTab[529427]++
//line /snap/go/10455/src/net/parse.go:192
			_go_fuzz_dep_.CoverTab[7748]++
//line /snap/go/10455/src/net/parse.go:192
			// _ = "end of CoverTab[7748]"
//line /snap/go/10455/src/net/parse.go:192
		}
//line /snap/go/10455/src/net/parse.go:192
		// _ = "end of CoverTab[7746]"
	}
//line /snap/go/10455/src/net/parse.go:193
	if _go_fuzz_dep_.CoverTab[786731] == 0 {
//line /snap/go/10455/src/net/parse.go:193
		_go_fuzz_dep_.CoverTab[529472]++
//line /snap/go/10455/src/net/parse.go:193
	} else {
//line /snap/go/10455/src/net/parse.go:193
		_go_fuzz_dep_.CoverTab[529473]++
//line /snap/go/10455/src/net/parse.go:193
	}
//line /snap/go/10455/src/net/parse.go:193
	// _ = "end of CoverTab[7742]"
//line /snap/go/10455/src/net/parse.go:193
	_go_fuzz_dep_.CoverTab[7743]++
						return dst
//line /snap/go/10455/src/net/parse.go:194
	// _ = "end of CoverTab[7743]"
}

// Number of occurrences of b in s.
func count(s string, b byte) int {
//line /snap/go/10455/src/net/parse.go:198
	_go_fuzz_dep_.CoverTab[7749]++
						n := 0
//line /snap/go/10455/src/net/parse.go:199
	_go_fuzz_dep_.CoverTab[786732] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/parse.go:200
		if _go_fuzz_dep_.CoverTab[786732] == 0 {
//line /snap/go/10455/src/net/parse.go:200
			_go_fuzz_dep_.CoverTab[529474]++
//line /snap/go/10455/src/net/parse.go:200
		} else {
//line /snap/go/10455/src/net/parse.go:200
			_go_fuzz_dep_.CoverTab[529475]++
//line /snap/go/10455/src/net/parse.go:200
		}
//line /snap/go/10455/src/net/parse.go:200
		_go_fuzz_dep_.CoverTab[786732] = 1
//line /snap/go/10455/src/net/parse.go:200
		_go_fuzz_dep_.CoverTab[7751]++
							if s[i] == b {
//line /snap/go/10455/src/net/parse.go:201
			_go_fuzz_dep_.CoverTab[529428]++
//line /snap/go/10455/src/net/parse.go:201
			_go_fuzz_dep_.CoverTab[7752]++
								n++
//line /snap/go/10455/src/net/parse.go:202
			// _ = "end of CoverTab[7752]"
		} else {
//line /snap/go/10455/src/net/parse.go:203
			_go_fuzz_dep_.CoverTab[529429]++
//line /snap/go/10455/src/net/parse.go:203
			_go_fuzz_dep_.CoverTab[7753]++
//line /snap/go/10455/src/net/parse.go:203
			// _ = "end of CoverTab[7753]"
//line /snap/go/10455/src/net/parse.go:203
		}
//line /snap/go/10455/src/net/parse.go:203
		// _ = "end of CoverTab[7751]"
	}
//line /snap/go/10455/src/net/parse.go:204
	if _go_fuzz_dep_.CoverTab[786732] == 0 {
//line /snap/go/10455/src/net/parse.go:204
		_go_fuzz_dep_.CoverTab[529476]++
//line /snap/go/10455/src/net/parse.go:204
	} else {
//line /snap/go/10455/src/net/parse.go:204
		_go_fuzz_dep_.CoverTab[529477]++
//line /snap/go/10455/src/net/parse.go:204
	}
//line /snap/go/10455/src/net/parse.go:204
	// _ = "end of CoverTab[7749]"
//line /snap/go/10455/src/net/parse.go:204
	_go_fuzz_dep_.CoverTab[7750]++
						return n
//line /snap/go/10455/src/net/parse.go:205
	// _ = "end of CoverTab[7750]"
}

// Index of rightmost occurrence of b in s.
func last(s string, b byte) int {
//line /snap/go/10455/src/net/parse.go:209
	_go_fuzz_dep_.CoverTab[7754]++
						i := len(s)
//line /snap/go/10455/src/net/parse.go:210
	_go_fuzz_dep_.CoverTab[786733] = 0
						for i--; i >= 0; i-- {
//line /snap/go/10455/src/net/parse.go:211
		if _go_fuzz_dep_.CoverTab[786733] == 0 {
//line /snap/go/10455/src/net/parse.go:211
			_go_fuzz_dep_.CoverTab[529478]++
//line /snap/go/10455/src/net/parse.go:211
		} else {
//line /snap/go/10455/src/net/parse.go:211
			_go_fuzz_dep_.CoverTab[529479]++
//line /snap/go/10455/src/net/parse.go:211
		}
//line /snap/go/10455/src/net/parse.go:211
		_go_fuzz_dep_.CoverTab[786733] = 1
//line /snap/go/10455/src/net/parse.go:211
		_go_fuzz_dep_.CoverTab[7756]++
							if s[i] == b {
//line /snap/go/10455/src/net/parse.go:212
			_go_fuzz_dep_.CoverTab[529430]++
//line /snap/go/10455/src/net/parse.go:212
			_go_fuzz_dep_.CoverTab[7757]++
								break
//line /snap/go/10455/src/net/parse.go:213
			// _ = "end of CoverTab[7757]"
		} else {
//line /snap/go/10455/src/net/parse.go:214
			_go_fuzz_dep_.CoverTab[529431]++
//line /snap/go/10455/src/net/parse.go:214
			_go_fuzz_dep_.CoverTab[7758]++
//line /snap/go/10455/src/net/parse.go:214
			// _ = "end of CoverTab[7758]"
//line /snap/go/10455/src/net/parse.go:214
		}
//line /snap/go/10455/src/net/parse.go:214
		// _ = "end of CoverTab[7756]"
	}
//line /snap/go/10455/src/net/parse.go:215
	if _go_fuzz_dep_.CoverTab[786733] == 0 {
//line /snap/go/10455/src/net/parse.go:215
		_go_fuzz_dep_.CoverTab[529480]++
//line /snap/go/10455/src/net/parse.go:215
	} else {
//line /snap/go/10455/src/net/parse.go:215
		_go_fuzz_dep_.CoverTab[529481]++
//line /snap/go/10455/src/net/parse.go:215
	}
//line /snap/go/10455/src/net/parse.go:215
	// _ = "end of CoverTab[7754]"
//line /snap/go/10455/src/net/parse.go:215
	_go_fuzz_dep_.CoverTab[7755]++
						return i
//line /snap/go/10455/src/net/parse.go:216
	// _ = "end of CoverTab[7755]"
}

// hasUpperCase tells whether the given string contains at least one upper-case.
func hasUpperCase(s string) bool {
//line /snap/go/10455/src/net/parse.go:220
	_go_fuzz_dep_.CoverTab[7759]++
//line /snap/go/10455/src/net/parse.go:220
	_go_fuzz_dep_.CoverTab[786734] = 0
						for i := range s {
//line /snap/go/10455/src/net/parse.go:221
		if _go_fuzz_dep_.CoverTab[786734] == 0 {
//line /snap/go/10455/src/net/parse.go:221
			_go_fuzz_dep_.CoverTab[529482]++
//line /snap/go/10455/src/net/parse.go:221
		} else {
//line /snap/go/10455/src/net/parse.go:221
			_go_fuzz_dep_.CoverTab[529483]++
//line /snap/go/10455/src/net/parse.go:221
		}
//line /snap/go/10455/src/net/parse.go:221
		_go_fuzz_dep_.CoverTab[786734] = 1
//line /snap/go/10455/src/net/parse.go:221
		_go_fuzz_dep_.CoverTab[7761]++
							if 'A' <= s[i] && func() bool {
//line /snap/go/10455/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[7762]++
//line /snap/go/10455/src/net/parse.go:222
			return s[i] <= 'Z'
//line /snap/go/10455/src/net/parse.go:222
			// _ = "end of CoverTab[7762]"
//line /snap/go/10455/src/net/parse.go:222
		}() {
//line /snap/go/10455/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[529432]++
//line /snap/go/10455/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[7763]++
								return true
//line /snap/go/10455/src/net/parse.go:223
			// _ = "end of CoverTab[7763]"
		} else {
//line /snap/go/10455/src/net/parse.go:224
			_go_fuzz_dep_.CoverTab[529433]++
//line /snap/go/10455/src/net/parse.go:224
			_go_fuzz_dep_.CoverTab[7764]++
//line /snap/go/10455/src/net/parse.go:224
			// _ = "end of CoverTab[7764]"
//line /snap/go/10455/src/net/parse.go:224
		}
//line /snap/go/10455/src/net/parse.go:224
		// _ = "end of CoverTab[7761]"
	}
//line /snap/go/10455/src/net/parse.go:225
	if _go_fuzz_dep_.CoverTab[786734] == 0 {
//line /snap/go/10455/src/net/parse.go:225
		_go_fuzz_dep_.CoverTab[529484]++
//line /snap/go/10455/src/net/parse.go:225
	} else {
//line /snap/go/10455/src/net/parse.go:225
		_go_fuzz_dep_.CoverTab[529485]++
//line /snap/go/10455/src/net/parse.go:225
	}
//line /snap/go/10455/src/net/parse.go:225
	// _ = "end of CoverTab[7759]"
//line /snap/go/10455/src/net/parse.go:225
	_go_fuzz_dep_.CoverTab[7760]++
						return false
//line /snap/go/10455/src/net/parse.go:226
	// _ = "end of CoverTab[7760]"
}

// lowerASCIIBytes makes x ASCII lowercase in-place.
func lowerASCIIBytes(x []byte) {
//line /snap/go/10455/src/net/parse.go:230
	_go_fuzz_dep_.CoverTab[7765]++
//line /snap/go/10455/src/net/parse.go:230
	_go_fuzz_dep_.CoverTab[786735] = 0
						for i, b := range x {
//line /snap/go/10455/src/net/parse.go:231
		if _go_fuzz_dep_.CoverTab[786735] == 0 {
//line /snap/go/10455/src/net/parse.go:231
			_go_fuzz_dep_.CoverTab[529486]++
//line /snap/go/10455/src/net/parse.go:231
		} else {
//line /snap/go/10455/src/net/parse.go:231
			_go_fuzz_dep_.CoverTab[529487]++
//line /snap/go/10455/src/net/parse.go:231
		}
//line /snap/go/10455/src/net/parse.go:231
		_go_fuzz_dep_.CoverTab[786735] = 1
//line /snap/go/10455/src/net/parse.go:231
		_go_fuzz_dep_.CoverTab[7766]++
							if 'A' <= b && func() bool {
//line /snap/go/10455/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[7767]++
//line /snap/go/10455/src/net/parse.go:232
			return b <= 'Z'
//line /snap/go/10455/src/net/parse.go:232
			// _ = "end of CoverTab[7767]"
//line /snap/go/10455/src/net/parse.go:232
		}() {
//line /snap/go/10455/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[529434]++
//line /snap/go/10455/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[7768]++
								x[i] += 'a' - 'A'
//line /snap/go/10455/src/net/parse.go:233
			// _ = "end of CoverTab[7768]"
		} else {
//line /snap/go/10455/src/net/parse.go:234
			_go_fuzz_dep_.CoverTab[529435]++
//line /snap/go/10455/src/net/parse.go:234
			_go_fuzz_dep_.CoverTab[7769]++
//line /snap/go/10455/src/net/parse.go:234
			// _ = "end of CoverTab[7769]"
//line /snap/go/10455/src/net/parse.go:234
		}
//line /snap/go/10455/src/net/parse.go:234
		// _ = "end of CoverTab[7766]"
	}
//line /snap/go/10455/src/net/parse.go:235
	if _go_fuzz_dep_.CoverTab[786735] == 0 {
//line /snap/go/10455/src/net/parse.go:235
		_go_fuzz_dep_.CoverTab[529488]++
//line /snap/go/10455/src/net/parse.go:235
	} else {
//line /snap/go/10455/src/net/parse.go:235
		_go_fuzz_dep_.CoverTab[529489]++
//line /snap/go/10455/src/net/parse.go:235
	}
//line /snap/go/10455/src/net/parse.go:235
	// _ = "end of CoverTab[7765]"
}

// lowerASCII returns the ASCII lowercase version of b.
func lowerASCII(b byte) byte {
//line /snap/go/10455/src/net/parse.go:239
	_go_fuzz_dep_.CoverTab[7770]++
						if 'A' <= b && func() bool {
//line /snap/go/10455/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[7772]++
//line /snap/go/10455/src/net/parse.go:240
		return b <= 'Z'
//line /snap/go/10455/src/net/parse.go:240
		// _ = "end of CoverTab[7772]"
//line /snap/go/10455/src/net/parse.go:240
	}() {
//line /snap/go/10455/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[529436]++
//line /snap/go/10455/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[7773]++
							return b + ('a' - 'A')
//line /snap/go/10455/src/net/parse.go:241
		// _ = "end of CoverTab[7773]"
	} else {
//line /snap/go/10455/src/net/parse.go:242
		_go_fuzz_dep_.CoverTab[529437]++
//line /snap/go/10455/src/net/parse.go:242
		_go_fuzz_dep_.CoverTab[7774]++
//line /snap/go/10455/src/net/parse.go:242
		// _ = "end of CoverTab[7774]"
//line /snap/go/10455/src/net/parse.go:242
	}
//line /snap/go/10455/src/net/parse.go:242
	// _ = "end of CoverTab[7770]"
//line /snap/go/10455/src/net/parse.go:242
	_go_fuzz_dep_.CoverTab[7771]++
						return b
//line /snap/go/10455/src/net/parse.go:243
	// _ = "end of CoverTab[7771]"
}

// trimSpace returns x without any leading or trailing ASCII whitespace.
func trimSpace(x string) string {
//line /snap/go/10455/src/net/parse.go:247
	_go_fuzz_dep_.CoverTab[7775]++
//line /snap/go/10455/src/net/parse.go:247
	_go_fuzz_dep_.CoverTab[786736] = 0
						for len(x) > 0 && func() bool {
//line /snap/go/10455/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[7778]++
//line /snap/go/10455/src/net/parse.go:248
		return isSpace(x[0])
//line /snap/go/10455/src/net/parse.go:248
		// _ = "end of CoverTab[7778]"
//line /snap/go/10455/src/net/parse.go:248
	}() {
//line /snap/go/10455/src/net/parse.go:248
		if _go_fuzz_dep_.CoverTab[786736] == 0 {
//line /snap/go/10455/src/net/parse.go:248
			_go_fuzz_dep_.CoverTab[529490]++
//line /snap/go/10455/src/net/parse.go:248
		} else {
//line /snap/go/10455/src/net/parse.go:248
			_go_fuzz_dep_.CoverTab[529491]++
//line /snap/go/10455/src/net/parse.go:248
		}
//line /snap/go/10455/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[786736] = 1
//line /snap/go/10455/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[7779]++
							x = x[1:]
//line /snap/go/10455/src/net/parse.go:249
		// _ = "end of CoverTab[7779]"
	}
//line /snap/go/10455/src/net/parse.go:250
	if _go_fuzz_dep_.CoverTab[786736] == 0 {
//line /snap/go/10455/src/net/parse.go:250
		_go_fuzz_dep_.CoverTab[529492]++
//line /snap/go/10455/src/net/parse.go:250
	} else {
//line /snap/go/10455/src/net/parse.go:250
		_go_fuzz_dep_.CoverTab[529493]++
//line /snap/go/10455/src/net/parse.go:250
	}
//line /snap/go/10455/src/net/parse.go:250
	// _ = "end of CoverTab[7775]"
//line /snap/go/10455/src/net/parse.go:250
	_go_fuzz_dep_.CoverTab[7776]++
//line /snap/go/10455/src/net/parse.go:250
	_go_fuzz_dep_.CoverTab[786737] = 0
						for len(x) > 0 && func() bool {
//line /snap/go/10455/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[7780]++
//line /snap/go/10455/src/net/parse.go:251
		return isSpace(x[len(x)-1])
//line /snap/go/10455/src/net/parse.go:251
		// _ = "end of CoverTab[7780]"
//line /snap/go/10455/src/net/parse.go:251
	}() {
//line /snap/go/10455/src/net/parse.go:251
		if _go_fuzz_dep_.CoverTab[786737] == 0 {
//line /snap/go/10455/src/net/parse.go:251
			_go_fuzz_dep_.CoverTab[529494]++
//line /snap/go/10455/src/net/parse.go:251
		} else {
//line /snap/go/10455/src/net/parse.go:251
			_go_fuzz_dep_.CoverTab[529495]++
//line /snap/go/10455/src/net/parse.go:251
		}
//line /snap/go/10455/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[786737] = 1
//line /snap/go/10455/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[7781]++
							x = x[:len(x)-1]
//line /snap/go/10455/src/net/parse.go:252
		// _ = "end of CoverTab[7781]"
	}
//line /snap/go/10455/src/net/parse.go:253
	if _go_fuzz_dep_.CoverTab[786737] == 0 {
//line /snap/go/10455/src/net/parse.go:253
		_go_fuzz_dep_.CoverTab[529496]++
//line /snap/go/10455/src/net/parse.go:253
	} else {
//line /snap/go/10455/src/net/parse.go:253
		_go_fuzz_dep_.CoverTab[529497]++
//line /snap/go/10455/src/net/parse.go:253
	}
//line /snap/go/10455/src/net/parse.go:253
	// _ = "end of CoverTab[7776]"
//line /snap/go/10455/src/net/parse.go:253
	_go_fuzz_dep_.CoverTab[7777]++
						return x
//line /snap/go/10455/src/net/parse.go:254
	// _ = "end of CoverTab[7777]"
}

// isSpace reports whether b is an ASCII space character.
func isSpace(b byte) bool {
//line /snap/go/10455/src/net/parse.go:258
	_go_fuzz_dep_.CoverTab[7782]++
						return b == ' ' || func() bool {
//line /snap/go/10455/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[7783]++
//line /snap/go/10455/src/net/parse.go:259
		return b == '\t'
//line /snap/go/10455/src/net/parse.go:259
		// _ = "end of CoverTab[7783]"
//line /snap/go/10455/src/net/parse.go:259
	}() || func() bool {
//line /snap/go/10455/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[7784]++
//line /snap/go/10455/src/net/parse.go:259
		return b == '\n'
//line /snap/go/10455/src/net/parse.go:259
		// _ = "end of CoverTab[7784]"
//line /snap/go/10455/src/net/parse.go:259
	}() || func() bool {
//line /snap/go/10455/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[7785]++
//line /snap/go/10455/src/net/parse.go:259
		return b == '\r'
//line /snap/go/10455/src/net/parse.go:259
		// _ = "end of CoverTab[7785]"
//line /snap/go/10455/src/net/parse.go:259
	}()
//line /snap/go/10455/src/net/parse.go:259
	// _ = "end of CoverTab[7782]"
}

// removeComment returns line, removing any '#' byte and any following
//line /snap/go/10455/src/net/parse.go:262
// bytes.
//line /snap/go/10455/src/net/parse.go:264
func removeComment(line string) string {
//line /snap/go/10455/src/net/parse.go:264
	_go_fuzz_dep_.CoverTab[7786]++
						if i := bytealg.IndexByteString(line, '#'); i != -1 {
//line /snap/go/10455/src/net/parse.go:265
		_go_fuzz_dep_.CoverTab[529438]++
//line /snap/go/10455/src/net/parse.go:265
		_go_fuzz_dep_.CoverTab[7788]++
							return line[:i]
//line /snap/go/10455/src/net/parse.go:266
		// _ = "end of CoverTab[7788]"
	} else {
//line /snap/go/10455/src/net/parse.go:267
		_go_fuzz_dep_.CoverTab[529439]++
//line /snap/go/10455/src/net/parse.go:267
		_go_fuzz_dep_.CoverTab[7789]++
//line /snap/go/10455/src/net/parse.go:267
		// _ = "end of CoverTab[7789]"
//line /snap/go/10455/src/net/parse.go:267
	}
//line /snap/go/10455/src/net/parse.go:267
	// _ = "end of CoverTab[7786]"
//line /snap/go/10455/src/net/parse.go:267
	_go_fuzz_dep_.CoverTab[7787]++
						return line
//line /snap/go/10455/src/net/parse.go:268
	// _ = "end of CoverTab[7787]"
}

// foreachField runs fn on each non-empty run of non-space bytes in x.
//line /snap/go/10455/src/net/parse.go:271
// It returns the first non-nil error returned by fn.
//line /snap/go/10455/src/net/parse.go:273
func foreachField(x string, fn func(field string) error) error {
//line /snap/go/10455/src/net/parse.go:273
	_go_fuzz_dep_.CoverTab[7790]++
						x = trimSpace(x)
//line /snap/go/10455/src/net/parse.go:274
	_go_fuzz_dep_.CoverTab[786738] = 0
						for len(x) > 0 {
//line /snap/go/10455/src/net/parse.go:275
		if _go_fuzz_dep_.CoverTab[786738] == 0 {
//line /snap/go/10455/src/net/parse.go:275
			_go_fuzz_dep_.CoverTab[529498]++
//line /snap/go/10455/src/net/parse.go:275
		} else {
//line /snap/go/10455/src/net/parse.go:275
			_go_fuzz_dep_.CoverTab[529499]++
//line /snap/go/10455/src/net/parse.go:275
		}
//line /snap/go/10455/src/net/parse.go:275
		_go_fuzz_dep_.CoverTab[786738] = 1
//line /snap/go/10455/src/net/parse.go:275
		_go_fuzz_dep_.CoverTab[7792]++
							sp := bytealg.IndexByteString(x, ' ')
							if sp == -1 {
//line /snap/go/10455/src/net/parse.go:277
			_go_fuzz_dep_.CoverTab[529440]++
//line /snap/go/10455/src/net/parse.go:277
			_go_fuzz_dep_.CoverTab[7795]++
								return fn(x)
//line /snap/go/10455/src/net/parse.go:278
			// _ = "end of CoverTab[7795]"
		} else {
//line /snap/go/10455/src/net/parse.go:279
			_go_fuzz_dep_.CoverTab[529441]++
//line /snap/go/10455/src/net/parse.go:279
			_go_fuzz_dep_.CoverTab[7796]++
//line /snap/go/10455/src/net/parse.go:279
			// _ = "end of CoverTab[7796]"
//line /snap/go/10455/src/net/parse.go:279
		}
//line /snap/go/10455/src/net/parse.go:279
		// _ = "end of CoverTab[7792]"
//line /snap/go/10455/src/net/parse.go:279
		_go_fuzz_dep_.CoverTab[7793]++
							if field := trimSpace(x[:sp]); len(field) > 0 {
//line /snap/go/10455/src/net/parse.go:280
			_go_fuzz_dep_.CoverTab[529442]++
//line /snap/go/10455/src/net/parse.go:280
			_go_fuzz_dep_.CoverTab[7797]++
								if err := fn(field); err != nil {
//line /snap/go/10455/src/net/parse.go:281
				_go_fuzz_dep_.CoverTab[529444]++
//line /snap/go/10455/src/net/parse.go:281
				_go_fuzz_dep_.CoverTab[7798]++
									return err
//line /snap/go/10455/src/net/parse.go:282
				// _ = "end of CoverTab[7798]"
			} else {
//line /snap/go/10455/src/net/parse.go:283
				_go_fuzz_dep_.CoverTab[529445]++
//line /snap/go/10455/src/net/parse.go:283
				_go_fuzz_dep_.CoverTab[7799]++
//line /snap/go/10455/src/net/parse.go:283
				// _ = "end of CoverTab[7799]"
//line /snap/go/10455/src/net/parse.go:283
			}
//line /snap/go/10455/src/net/parse.go:283
			// _ = "end of CoverTab[7797]"
		} else {
//line /snap/go/10455/src/net/parse.go:284
			_go_fuzz_dep_.CoverTab[529443]++
//line /snap/go/10455/src/net/parse.go:284
			_go_fuzz_dep_.CoverTab[7800]++
//line /snap/go/10455/src/net/parse.go:284
			// _ = "end of CoverTab[7800]"
//line /snap/go/10455/src/net/parse.go:284
		}
//line /snap/go/10455/src/net/parse.go:284
		// _ = "end of CoverTab[7793]"
//line /snap/go/10455/src/net/parse.go:284
		_go_fuzz_dep_.CoverTab[7794]++
							x = trimSpace(x[sp+1:])
//line /snap/go/10455/src/net/parse.go:285
		// _ = "end of CoverTab[7794]"
	}
//line /snap/go/10455/src/net/parse.go:286
	if _go_fuzz_dep_.CoverTab[786738] == 0 {
//line /snap/go/10455/src/net/parse.go:286
		_go_fuzz_dep_.CoverTab[529500]++
//line /snap/go/10455/src/net/parse.go:286
	} else {
//line /snap/go/10455/src/net/parse.go:286
		_go_fuzz_dep_.CoverTab[529501]++
//line /snap/go/10455/src/net/parse.go:286
	}
//line /snap/go/10455/src/net/parse.go:286
	// _ = "end of CoverTab[7790]"
//line /snap/go/10455/src/net/parse.go:286
	_go_fuzz_dep_.CoverTab[7791]++
						return nil
//line /snap/go/10455/src/net/parse.go:287
	// _ = "end of CoverTab[7791]"
}

// stringsHasSuffix is strings.HasSuffix. It reports whether s ends in
//line /snap/go/10455/src/net/parse.go:290
// suffix.
//line /snap/go/10455/src/net/parse.go:292
func stringsHasSuffix(s, suffix string) bool {
//line /snap/go/10455/src/net/parse.go:292
	_go_fuzz_dep_.CoverTab[7801]++
						return len(s) >= len(suffix) && func() bool {
//line /snap/go/10455/src/net/parse.go:293
		_go_fuzz_dep_.CoverTab[7802]++
//line /snap/go/10455/src/net/parse.go:293
		return s[len(s)-len(suffix):] == suffix
//line /snap/go/10455/src/net/parse.go:293
		// _ = "end of CoverTab[7802]"
//line /snap/go/10455/src/net/parse.go:293
	}()
//line /snap/go/10455/src/net/parse.go:293
	// _ = "end of CoverTab[7801]"
}

// stringsHasSuffixFold reports whether s ends in suffix,
//line /snap/go/10455/src/net/parse.go:296
// ASCII-case-insensitively.
//line /snap/go/10455/src/net/parse.go:298
func stringsHasSuffixFold(s, suffix string) bool {
//line /snap/go/10455/src/net/parse.go:298
	_go_fuzz_dep_.CoverTab[7803]++
						return len(s) >= len(suffix) && func() bool {
//line /snap/go/10455/src/net/parse.go:299
		_go_fuzz_dep_.CoverTab[7804]++
//line /snap/go/10455/src/net/parse.go:299
		return stringsEqualFold(s[len(s)-len(suffix):], suffix)
//line /snap/go/10455/src/net/parse.go:299
		// _ = "end of CoverTab[7804]"
//line /snap/go/10455/src/net/parse.go:299
	}()
//line /snap/go/10455/src/net/parse.go:299
	// _ = "end of CoverTab[7803]"
}

// stringsHasPrefix is strings.HasPrefix. It reports whether s begins with prefix.
func stringsHasPrefix(s, prefix string) bool {
//line /snap/go/10455/src/net/parse.go:303
	_go_fuzz_dep_.CoverTab[7805]++
						return len(s) >= len(prefix) && func() bool {
//line /snap/go/10455/src/net/parse.go:304
		_go_fuzz_dep_.CoverTab[7806]++
//line /snap/go/10455/src/net/parse.go:304
		return s[:len(prefix)] == prefix
//line /snap/go/10455/src/net/parse.go:304
		// _ = "end of CoverTab[7806]"
//line /snap/go/10455/src/net/parse.go:304
	}()
//line /snap/go/10455/src/net/parse.go:304
	// _ = "end of CoverTab[7805]"
}

// stringsEqualFold is strings.EqualFold, ASCII only. It reports whether s and t
//line /snap/go/10455/src/net/parse.go:307
// are equal, ASCII-case-insensitively.
//line /snap/go/10455/src/net/parse.go:309
func stringsEqualFold(s, t string) bool {
//line /snap/go/10455/src/net/parse.go:309
	_go_fuzz_dep_.CoverTab[7807]++
						if len(s) != len(t) {
//line /snap/go/10455/src/net/parse.go:310
		_go_fuzz_dep_.CoverTab[529446]++
//line /snap/go/10455/src/net/parse.go:310
		_go_fuzz_dep_.CoverTab[7810]++
							return false
//line /snap/go/10455/src/net/parse.go:311
		// _ = "end of CoverTab[7810]"
	} else {
//line /snap/go/10455/src/net/parse.go:312
		_go_fuzz_dep_.CoverTab[529447]++
//line /snap/go/10455/src/net/parse.go:312
		_go_fuzz_dep_.CoverTab[7811]++
//line /snap/go/10455/src/net/parse.go:312
		// _ = "end of CoverTab[7811]"
//line /snap/go/10455/src/net/parse.go:312
	}
//line /snap/go/10455/src/net/parse.go:312
	// _ = "end of CoverTab[7807]"
//line /snap/go/10455/src/net/parse.go:312
	_go_fuzz_dep_.CoverTab[7808]++
//line /snap/go/10455/src/net/parse.go:312
	_go_fuzz_dep_.CoverTab[786739] = 0
						for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/net/parse.go:313
		if _go_fuzz_dep_.CoverTab[786739] == 0 {
//line /snap/go/10455/src/net/parse.go:313
			_go_fuzz_dep_.CoverTab[529502]++
//line /snap/go/10455/src/net/parse.go:313
		} else {
//line /snap/go/10455/src/net/parse.go:313
			_go_fuzz_dep_.CoverTab[529503]++
//line /snap/go/10455/src/net/parse.go:313
		}
//line /snap/go/10455/src/net/parse.go:313
		_go_fuzz_dep_.CoverTab[786739] = 1
//line /snap/go/10455/src/net/parse.go:313
		_go_fuzz_dep_.CoverTab[7812]++
							if lowerASCII(s[i]) != lowerASCII(t[i]) {
//line /snap/go/10455/src/net/parse.go:314
			_go_fuzz_dep_.CoverTab[529448]++
//line /snap/go/10455/src/net/parse.go:314
			_go_fuzz_dep_.CoverTab[7813]++
								return false
//line /snap/go/10455/src/net/parse.go:315
			// _ = "end of CoverTab[7813]"
		} else {
//line /snap/go/10455/src/net/parse.go:316
			_go_fuzz_dep_.CoverTab[529449]++
//line /snap/go/10455/src/net/parse.go:316
			_go_fuzz_dep_.CoverTab[7814]++
//line /snap/go/10455/src/net/parse.go:316
			// _ = "end of CoverTab[7814]"
//line /snap/go/10455/src/net/parse.go:316
		}
//line /snap/go/10455/src/net/parse.go:316
		// _ = "end of CoverTab[7812]"
	}
//line /snap/go/10455/src/net/parse.go:317
	if _go_fuzz_dep_.CoverTab[786739] == 0 {
//line /snap/go/10455/src/net/parse.go:317
		_go_fuzz_dep_.CoverTab[529504]++
//line /snap/go/10455/src/net/parse.go:317
	} else {
//line /snap/go/10455/src/net/parse.go:317
		_go_fuzz_dep_.CoverTab[529505]++
//line /snap/go/10455/src/net/parse.go:317
	}
//line /snap/go/10455/src/net/parse.go:317
	// _ = "end of CoverTab[7808]"
//line /snap/go/10455/src/net/parse.go:317
	_go_fuzz_dep_.CoverTab[7809]++
						return true
//line /snap/go/10455/src/net/parse.go:318
	// _ = "end of CoverTab[7809]"
}

//line /snap/go/10455/src/net/parse.go:319
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/net/parse.go:319
var _ = _go_fuzz_dep_.CoverTab
