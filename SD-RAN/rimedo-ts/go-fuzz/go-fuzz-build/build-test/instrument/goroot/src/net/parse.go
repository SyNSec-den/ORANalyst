// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple file i/o and string manipulation, to avoid
// depending on strconv and bufio and strings.

//line /usr/local/go/src/net/parse.go:8
package net

//line /usr/local/go/src/net/parse.go:8
import (
//line /usr/local/go/src/net/parse.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/parse.go:8
)
//line /usr/local/go/src/net/parse.go:8
import (
//line /usr/local/go/src/net/parse.go:8
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/parse.go:8
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
//line /usr/local/go/src/net/parse.go:23
	_go_fuzz_dep_.CoverTab[7375]++
//line /usr/local/go/src/net/parse.go:23
	f.file.Close()
//line /usr/local/go/src/net/parse.go:23
	// _ = "end of CoverTab[7375]"
//line /usr/local/go/src/net/parse.go:23
}

func (f *file) getLineFromData() (s string, ok bool) {
//line /usr/local/go/src/net/parse.go:25
	_go_fuzz_dep_.CoverTab[7376]++
						data := f.data
						i := 0
						for i = 0; i < len(data); i++ {
//line /usr/local/go/src/net/parse.go:28
		_go_fuzz_dep_.CoverTab[7379]++
							if data[i] == '\n' {
//line /usr/local/go/src/net/parse.go:29
			_go_fuzz_dep_.CoverTab[7380]++
								s = string(data[0:i])
								ok = true

								i++
								n := len(data) - i
								copy(data[0:], data[i:])
								f.data = data[0:n]
								return
//line /usr/local/go/src/net/parse.go:37
			// _ = "end of CoverTab[7380]"
		} else {
//line /usr/local/go/src/net/parse.go:38
			_go_fuzz_dep_.CoverTab[7381]++
//line /usr/local/go/src/net/parse.go:38
			// _ = "end of CoverTab[7381]"
//line /usr/local/go/src/net/parse.go:38
		}
//line /usr/local/go/src/net/parse.go:38
		// _ = "end of CoverTab[7379]"
	}
//line /usr/local/go/src/net/parse.go:39
	// _ = "end of CoverTab[7376]"
//line /usr/local/go/src/net/parse.go:39
	_go_fuzz_dep_.CoverTab[7377]++
						if f.atEOF && func() bool {
//line /usr/local/go/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[7382]++
//line /usr/local/go/src/net/parse.go:40
		return len(f.data) > 0
//line /usr/local/go/src/net/parse.go:40
		// _ = "end of CoverTab[7382]"
//line /usr/local/go/src/net/parse.go:40
	}() {
//line /usr/local/go/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[7383]++

							s = string(data)
							f.data = f.data[0:0]
							ok = true
//line /usr/local/go/src/net/parse.go:44
		// _ = "end of CoverTab[7383]"
	} else {
//line /usr/local/go/src/net/parse.go:45
		_go_fuzz_dep_.CoverTab[7384]++
//line /usr/local/go/src/net/parse.go:45
		// _ = "end of CoverTab[7384]"
//line /usr/local/go/src/net/parse.go:45
	}
//line /usr/local/go/src/net/parse.go:45
	// _ = "end of CoverTab[7377]"
//line /usr/local/go/src/net/parse.go:45
	_go_fuzz_dep_.CoverTab[7378]++
						return
//line /usr/local/go/src/net/parse.go:46
	// _ = "end of CoverTab[7378]"
}

func (f *file) readLine() (s string, ok bool) {
//line /usr/local/go/src/net/parse.go:49
	_go_fuzz_dep_.CoverTab[7385]++
						if s, ok = f.getLineFromData(); ok {
//line /usr/local/go/src/net/parse.go:50
		_go_fuzz_dep_.CoverTab[7388]++
							return
//line /usr/local/go/src/net/parse.go:51
		// _ = "end of CoverTab[7388]"
	} else {
//line /usr/local/go/src/net/parse.go:52
		_go_fuzz_dep_.CoverTab[7389]++
//line /usr/local/go/src/net/parse.go:52
		// _ = "end of CoverTab[7389]"
//line /usr/local/go/src/net/parse.go:52
	}
//line /usr/local/go/src/net/parse.go:52
	// _ = "end of CoverTab[7385]"
//line /usr/local/go/src/net/parse.go:52
	_go_fuzz_dep_.CoverTab[7386]++
						if len(f.data) < cap(f.data) {
//line /usr/local/go/src/net/parse.go:53
		_go_fuzz_dep_.CoverTab[7390]++
							ln := len(f.data)
							n, err := io.ReadFull(f.file, f.data[ln:cap(f.data)])
							if n >= 0 {
//line /usr/local/go/src/net/parse.go:56
			_go_fuzz_dep_.CoverTab[7392]++
								f.data = f.data[0 : ln+n]
//line /usr/local/go/src/net/parse.go:57
			// _ = "end of CoverTab[7392]"
		} else {
//line /usr/local/go/src/net/parse.go:58
			_go_fuzz_dep_.CoverTab[7393]++
//line /usr/local/go/src/net/parse.go:58
			// _ = "end of CoverTab[7393]"
//line /usr/local/go/src/net/parse.go:58
		}
//line /usr/local/go/src/net/parse.go:58
		// _ = "end of CoverTab[7390]"
//line /usr/local/go/src/net/parse.go:58
		_go_fuzz_dep_.CoverTab[7391]++
							if err == io.EOF || func() bool {
//line /usr/local/go/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[7394]++
//line /usr/local/go/src/net/parse.go:59
			return err == io.ErrUnexpectedEOF
//line /usr/local/go/src/net/parse.go:59
			// _ = "end of CoverTab[7394]"
//line /usr/local/go/src/net/parse.go:59
		}() {
//line /usr/local/go/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[7395]++
								f.atEOF = true
//line /usr/local/go/src/net/parse.go:60
			// _ = "end of CoverTab[7395]"
		} else {
//line /usr/local/go/src/net/parse.go:61
			_go_fuzz_dep_.CoverTab[7396]++
//line /usr/local/go/src/net/parse.go:61
			// _ = "end of CoverTab[7396]"
//line /usr/local/go/src/net/parse.go:61
		}
//line /usr/local/go/src/net/parse.go:61
		// _ = "end of CoverTab[7391]"
	} else {
//line /usr/local/go/src/net/parse.go:62
		_go_fuzz_dep_.CoverTab[7397]++
//line /usr/local/go/src/net/parse.go:62
		// _ = "end of CoverTab[7397]"
//line /usr/local/go/src/net/parse.go:62
	}
//line /usr/local/go/src/net/parse.go:62
	// _ = "end of CoverTab[7386]"
//line /usr/local/go/src/net/parse.go:62
	_go_fuzz_dep_.CoverTab[7387]++
						s, ok = f.getLineFromData()
						return
//line /usr/local/go/src/net/parse.go:64
	// _ = "end of CoverTab[7387]"
}

func (f *file) stat() (mtime time.Time, size int64, err error) {
//line /usr/local/go/src/net/parse.go:67
	_go_fuzz_dep_.CoverTab[7398]++
						st, err := f.file.Stat()
						if err != nil {
//line /usr/local/go/src/net/parse.go:69
		_go_fuzz_dep_.CoverTab[7400]++
							return time.Time{}, 0, err
//line /usr/local/go/src/net/parse.go:70
		// _ = "end of CoverTab[7400]"
	} else {
//line /usr/local/go/src/net/parse.go:71
		_go_fuzz_dep_.CoverTab[7401]++
//line /usr/local/go/src/net/parse.go:71
		// _ = "end of CoverTab[7401]"
//line /usr/local/go/src/net/parse.go:71
	}
//line /usr/local/go/src/net/parse.go:71
	// _ = "end of CoverTab[7398]"
//line /usr/local/go/src/net/parse.go:71
	_go_fuzz_dep_.CoverTab[7399]++
						return st.ModTime(), st.Size(), nil
//line /usr/local/go/src/net/parse.go:72
	// _ = "end of CoverTab[7399]"
}

func open(name string) (*file, error) {
//line /usr/local/go/src/net/parse.go:75
	_go_fuzz_dep_.CoverTab[7402]++
						fd, err := os.Open(name)
						if err != nil {
//line /usr/local/go/src/net/parse.go:77
		_go_fuzz_dep_.CoverTab[7404]++
							return nil, err
//line /usr/local/go/src/net/parse.go:78
		// _ = "end of CoverTab[7404]"
	} else {
//line /usr/local/go/src/net/parse.go:79
		_go_fuzz_dep_.CoverTab[7405]++
//line /usr/local/go/src/net/parse.go:79
		// _ = "end of CoverTab[7405]"
//line /usr/local/go/src/net/parse.go:79
	}
//line /usr/local/go/src/net/parse.go:79
	// _ = "end of CoverTab[7402]"
//line /usr/local/go/src/net/parse.go:79
	_go_fuzz_dep_.CoverTab[7403]++
						return &file{fd, make([]byte, 0, 64*1024), false}, nil
//line /usr/local/go/src/net/parse.go:80
	// _ = "end of CoverTab[7403]"
}

func stat(name string) (mtime time.Time, size int64, err error) {
//line /usr/local/go/src/net/parse.go:83
	_go_fuzz_dep_.CoverTab[7406]++
						st, err := os.Stat(name)
						if err != nil {
//line /usr/local/go/src/net/parse.go:85
		_go_fuzz_dep_.CoverTab[7408]++
							return time.Time{}, 0, err
//line /usr/local/go/src/net/parse.go:86
		// _ = "end of CoverTab[7408]"
	} else {
//line /usr/local/go/src/net/parse.go:87
		_go_fuzz_dep_.CoverTab[7409]++
//line /usr/local/go/src/net/parse.go:87
		// _ = "end of CoverTab[7409]"
//line /usr/local/go/src/net/parse.go:87
	}
//line /usr/local/go/src/net/parse.go:87
	// _ = "end of CoverTab[7406]"
//line /usr/local/go/src/net/parse.go:87
	_go_fuzz_dep_.CoverTab[7407]++
						return st.ModTime(), st.Size(), nil
//line /usr/local/go/src/net/parse.go:88
	// _ = "end of CoverTab[7407]"
}

// Count occurrences in s of any bytes in t.
func countAnyByte(s string, t string) int {
//line /usr/local/go/src/net/parse.go:92
	_go_fuzz_dep_.CoverTab[7410]++
						n := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:94
		_go_fuzz_dep_.CoverTab[7412]++
							if bytealg.IndexByteString(t, s[i]) >= 0 {
//line /usr/local/go/src/net/parse.go:95
			_go_fuzz_dep_.CoverTab[7413]++
								n++
//line /usr/local/go/src/net/parse.go:96
			// _ = "end of CoverTab[7413]"
		} else {
//line /usr/local/go/src/net/parse.go:97
			_go_fuzz_dep_.CoverTab[7414]++
//line /usr/local/go/src/net/parse.go:97
			// _ = "end of CoverTab[7414]"
//line /usr/local/go/src/net/parse.go:97
		}
//line /usr/local/go/src/net/parse.go:97
		// _ = "end of CoverTab[7412]"
	}
//line /usr/local/go/src/net/parse.go:98
	// _ = "end of CoverTab[7410]"
//line /usr/local/go/src/net/parse.go:98
	_go_fuzz_dep_.CoverTab[7411]++
						return n
//line /usr/local/go/src/net/parse.go:99
	// _ = "end of CoverTab[7411]"
}

// Split s at any bytes in t.
func splitAtBytes(s string, t string) []string {
//line /usr/local/go/src/net/parse.go:103
	_go_fuzz_dep_.CoverTab[7415]++
						a := make([]string, 1+countAnyByte(s, t))
						n := 0
						last := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:107
		_go_fuzz_dep_.CoverTab[7418]++
							if bytealg.IndexByteString(t, s[i]) >= 0 {
//line /usr/local/go/src/net/parse.go:108
			_go_fuzz_dep_.CoverTab[7419]++
								if last < i {
//line /usr/local/go/src/net/parse.go:109
				_go_fuzz_dep_.CoverTab[7421]++
									a[n] = s[last:i]
									n++
//line /usr/local/go/src/net/parse.go:111
				// _ = "end of CoverTab[7421]"
			} else {
//line /usr/local/go/src/net/parse.go:112
				_go_fuzz_dep_.CoverTab[7422]++
//line /usr/local/go/src/net/parse.go:112
				// _ = "end of CoverTab[7422]"
//line /usr/local/go/src/net/parse.go:112
			}
//line /usr/local/go/src/net/parse.go:112
			// _ = "end of CoverTab[7419]"
//line /usr/local/go/src/net/parse.go:112
			_go_fuzz_dep_.CoverTab[7420]++
								last = i + 1
//line /usr/local/go/src/net/parse.go:113
			// _ = "end of CoverTab[7420]"
		} else {
//line /usr/local/go/src/net/parse.go:114
			_go_fuzz_dep_.CoverTab[7423]++
//line /usr/local/go/src/net/parse.go:114
			// _ = "end of CoverTab[7423]"
//line /usr/local/go/src/net/parse.go:114
		}
//line /usr/local/go/src/net/parse.go:114
		// _ = "end of CoverTab[7418]"
	}
//line /usr/local/go/src/net/parse.go:115
	// _ = "end of CoverTab[7415]"
//line /usr/local/go/src/net/parse.go:115
	_go_fuzz_dep_.CoverTab[7416]++
						if last < len(s) {
//line /usr/local/go/src/net/parse.go:116
		_go_fuzz_dep_.CoverTab[7424]++
							a[n] = s[last:]
							n++
//line /usr/local/go/src/net/parse.go:118
		// _ = "end of CoverTab[7424]"
	} else {
//line /usr/local/go/src/net/parse.go:119
		_go_fuzz_dep_.CoverTab[7425]++
//line /usr/local/go/src/net/parse.go:119
		// _ = "end of CoverTab[7425]"
//line /usr/local/go/src/net/parse.go:119
	}
//line /usr/local/go/src/net/parse.go:119
	// _ = "end of CoverTab[7416]"
//line /usr/local/go/src/net/parse.go:119
	_go_fuzz_dep_.CoverTab[7417]++
						return a[0:n]
//line /usr/local/go/src/net/parse.go:120
	// _ = "end of CoverTab[7417]"
}

func getFields(s string) []string {
//line /usr/local/go/src/net/parse.go:123
	_go_fuzz_dep_.CoverTab[7426]++
//line /usr/local/go/src/net/parse.go:123
	return splitAtBytes(s, " \r\t\n")
//line /usr/local/go/src/net/parse.go:123
	// _ = "end of CoverTab[7426]"
//line /usr/local/go/src/net/parse.go:123
}

// Bigger than we need, not too big to worry about overflow
const big = 0xFFFFFF

// Decimal to integer.
//line /usr/local/go/src/net/parse.go:128
// Returns number, characters consumed, success.
//line /usr/local/go/src/net/parse.go:130
func dtoi(s string) (n int, i int, ok bool) {
//line /usr/local/go/src/net/parse.go:130
	_go_fuzz_dep_.CoverTab[7427]++
						n = 0
						for i = 0; i < len(s) && func() bool {
//line /usr/local/go/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[7430]++
//line /usr/local/go/src/net/parse.go:132
		return '0' <= s[i]
//line /usr/local/go/src/net/parse.go:132
		// _ = "end of CoverTab[7430]"
//line /usr/local/go/src/net/parse.go:132
	}() && func() bool {
//line /usr/local/go/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[7431]++
//line /usr/local/go/src/net/parse.go:132
		return s[i] <= '9'
//line /usr/local/go/src/net/parse.go:132
		// _ = "end of CoverTab[7431]"
//line /usr/local/go/src/net/parse.go:132
	}(); i++ {
//line /usr/local/go/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[7432]++
							n = n*10 + int(s[i]-'0')
							if n >= big {
//line /usr/local/go/src/net/parse.go:134
			_go_fuzz_dep_.CoverTab[7433]++
								return big, i, false
//line /usr/local/go/src/net/parse.go:135
			// _ = "end of CoverTab[7433]"
		} else {
//line /usr/local/go/src/net/parse.go:136
			_go_fuzz_dep_.CoverTab[7434]++
//line /usr/local/go/src/net/parse.go:136
			// _ = "end of CoverTab[7434]"
//line /usr/local/go/src/net/parse.go:136
		}
//line /usr/local/go/src/net/parse.go:136
		// _ = "end of CoverTab[7432]"
	}
//line /usr/local/go/src/net/parse.go:137
	// _ = "end of CoverTab[7427]"
//line /usr/local/go/src/net/parse.go:137
	_go_fuzz_dep_.CoverTab[7428]++
						if i == 0 {
//line /usr/local/go/src/net/parse.go:138
		_go_fuzz_dep_.CoverTab[7435]++
							return 0, 0, false
//line /usr/local/go/src/net/parse.go:139
		// _ = "end of CoverTab[7435]"
	} else {
//line /usr/local/go/src/net/parse.go:140
		_go_fuzz_dep_.CoverTab[7436]++
//line /usr/local/go/src/net/parse.go:140
		// _ = "end of CoverTab[7436]"
//line /usr/local/go/src/net/parse.go:140
	}
//line /usr/local/go/src/net/parse.go:140
	// _ = "end of CoverTab[7428]"
//line /usr/local/go/src/net/parse.go:140
	_go_fuzz_dep_.CoverTab[7429]++
						return n, i, true
//line /usr/local/go/src/net/parse.go:141
	// _ = "end of CoverTab[7429]"
}

// Hexadecimal to integer.
//line /usr/local/go/src/net/parse.go:144
// Returns number, characters consumed, success.
//line /usr/local/go/src/net/parse.go:146
func xtoi(s string) (n int, i int, ok bool) {
//line /usr/local/go/src/net/parse.go:146
	_go_fuzz_dep_.CoverTab[7437]++
						n = 0
						for i = 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:148
		_go_fuzz_dep_.CoverTab[7440]++
							if '0' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[7442]++
//line /usr/local/go/src/net/parse.go:149
			return s[i] <= '9'
//line /usr/local/go/src/net/parse.go:149
			// _ = "end of CoverTab[7442]"
//line /usr/local/go/src/net/parse.go:149
		}() {
//line /usr/local/go/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[7443]++
								n *= 16
								n += int(s[i] - '0')
//line /usr/local/go/src/net/parse.go:151
			// _ = "end of CoverTab[7443]"
		} else {
//line /usr/local/go/src/net/parse.go:152
			_go_fuzz_dep_.CoverTab[7444]++
//line /usr/local/go/src/net/parse.go:152
			if 'a' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[7445]++
//line /usr/local/go/src/net/parse.go:152
				return s[i] <= 'f'
//line /usr/local/go/src/net/parse.go:152
				// _ = "end of CoverTab[7445]"
//line /usr/local/go/src/net/parse.go:152
			}() {
//line /usr/local/go/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[7446]++
									n *= 16
									n += int(s[i]-'a') + 10
//line /usr/local/go/src/net/parse.go:154
				// _ = "end of CoverTab[7446]"
			} else {
//line /usr/local/go/src/net/parse.go:155
				_go_fuzz_dep_.CoverTab[7447]++
//line /usr/local/go/src/net/parse.go:155
				if 'A' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[7448]++
//line /usr/local/go/src/net/parse.go:155
					return s[i] <= 'F'
//line /usr/local/go/src/net/parse.go:155
					// _ = "end of CoverTab[7448]"
//line /usr/local/go/src/net/parse.go:155
				}() {
//line /usr/local/go/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[7449]++
										n *= 16
										n += int(s[i]-'A') + 10
//line /usr/local/go/src/net/parse.go:157
					// _ = "end of CoverTab[7449]"
				} else {
//line /usr/local/go/src/net/parse.go:158
					_go_fuzz_dep_.CoverTab[7450]++
										break
//line /usr/local/go/src/net/parse.go:159
					// _ = "end of CoverTab[7450]"
				}
//line /usr/local/go/src/net/parse.go:160
				// _ = "end of CoverTab[7447]"
//line /usr/local/go/src/net/parse.go:160
			}
//line /usr/local/go/src/net/parse.go:160
			// _ = "end of CoverTab[7444]"
//line /usr/local/go/src/net/parse.go:160
		}
//line /usr/local/go/src/net/parse.go:160
		// _ = "end of CoverTab[7440]"
//line /usr/local/go/src/net/parse.go:160
		_go_fuzz_dep_.CoverTab[7441]++
							if n >= big {
//line /usr/local/go/src/net/parse.go:161
			_go_fuzz_dep_.CoverTab[7451]++
								return 0, i, false
//line /usr/local/go/src/net/parse.go:162
			// _ = "end of CoverTab[7451]"
		} else {
//line /usr/local/go/src/net/parse.go:163
			_go_fuzz_dep_.CoverTab[7452]++
//line /usr/local/go/src/net/parse.go:163
			// _ = "end of CoverTab[7452]"
//line /usr/local/go/src/net/parse.go:163
		}
//line /usr/local/go/src/net/parse.go:163
		// _ = "end of CoverTab[7441]"
	}
//line /usr/local/go/src/net/parse.go:164
	// _ = "end of CoverTab[7437]"
//line /usr/local/go/src/net/parse.go:164
	_go_fuzz_dep_.CoverTab[7438]++
						if i == 0 {
//line /usr/local/go/src/net/parse.go:165
		_go_fuzz_dep_.CoverTab[7453]++
							return 0, i, false
//line /usr/local/go/src/net/parse.go:166
		// _ = "end of CoverTab[7453]"
	} else {
//line /usr/local/go/src/net/parse.go:167
		_go_fuzz_dep_.CoverTab[7454]++
//line /usr/local/go/src/net/parse.go:167
		// _ = "end of CoverTab[7454]"
//line /usr/local/go/src/net/parse.go:167
	}
//line /usr/local/go/src/net/parse.go:167
	// _ = "end of CoverTab[7438]"
//line /usr/local/go/src/net/parse.go:167
	_go_fuzz_dep_.CoverTab[7439]++
						return n, i, true
//line /usr/local/go/src/net/parse.go:168
	// _ = "end of CoverTab[7439]"
}

// xtoi2 converts the next two hex digits of s into a byte.
//line /usr/local/go/src/net/parse.go:171
// If s is longer than 2 bytes then the third byte must be e.
//line /usr/local/go/src/net/parse.go:171
// If the first two bytes of s are not hex digits or the third byte
//line /usr/local/go/src/net/parse.go:171
// does not match e, false is returned.
//line /usr/local/go/src/net/parse.go:175
func xtoi2(s string, e byte) (byte, bool) {
//line /usr/local/go/src/net/parse.go:175
	_go_fuzz_dep_.CoverTab[7455]++
						if len(s) > 2 && func() bool {
//line /usr/local/go/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[7457]++
//line /usr/local/go/src/net/parse.go:176
		return s[2] != e
//line /usr/local/go/src/net/parse.go:176
		// _ = "end of CoverTab[7457]"
//line /usr/local/go/src/net/parse.go:176
	}() {
//line /usr/local/go/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[7458]++
							return 0, false
//line /usr/local/go/src/net/parse.go:177
		// _ = "end of CoverTab[7458]"
	} else {
//line /usr/local/go/src/net/parse.go:178
		_go_fuzz_dep_.CoverTab[7459]++
//line /usr/local/go/src/net/parse.go:178
		// _ = "end of CoverTab[7459]"
//line /usr/local/go/src/net/parse.go:178
	}
//line /usr/local/go/src/net/parse.go:178
	// _ = "end of CoverTab[7455]"
//line /usr/local/go/src/net/parse.go:178
	_go_fuzz_dep_.CoverTab[7456]++
						n, ei, ok := xtoi(s[:2])
						return byte(n), ok && func() bool {
//line /usr/local/go/src/net/parse.go:180
		_go_fuzz_dep_.CoverTab[7460]++
//line /usr/local/go/src/net/parse.go:180
		return ei == 2
//line /usr/local/go/src/net/parse.go:180
		// _ = "end of CoverTab[7460]"
//line /usr/local/go/src/net/parse.go:180
	}()
//line /usr/local/go/src/net/parse.go:180
	// _ = "end of CoverTab[7456]"
}

// Convert i to a hexadecimal string. Leading zeros are not printed.
func appendHex(dst []byte, i uint32) []byte {
//line /usr/local/go/src/net/parse.go:184
	_go_fuzz_dep_.CoverTab[7461]++
						if i == 0 {
//line /usr/local/go/src/net/parse.go:185
		_go_fuzz_dep_.CoverTab[7464]++
							return append(dst, '0')
//line /usr/local/go/src/net/parse.go:186
		// _ = "end of CoverTab[7464]"
	} else {
//line /usr/local/go/src/net/parse.go:187
		_go_fuzz_dep_.CoverTab[7465]++
//line /usr/local/go/src/net/parse.go:187
		// _ = "end of CoverTab[7465]"
//line /usr/local/go/src/net/parse.go:187
	}
//line /usr/local/go/src/net/parse.go:187
	// _ = "end of CoverTab[7461]"
//line /usr/local/go/src/net/parse.go:187
	_go_fuzz_dep_.CoverTab[7462]++
						for j := 7; j >= 0; j-- {
//line /usr/local/go/src/net/parse.go:188
		_go_fuzz_dep_.CoverTab[7466]++
							v := i >> uint(j*4)
							if v > 0 {
//line /usr/local/go/src/net/parse.go:190
			_go_fuzz_dep_.CoverTab[7467]++
								dst = append(dst, hexDigit[v&0xf])
//line /usr/local/go/src/net/parse.go:191
			// _ = "end of CoverTab[7467]"
		} else {
//line /usr/local/go/src/net/parse.go:192
			_go_fuzz_dep_.CoverTab[7468]++
//line /usr/local/go/src/net/parse.go:192
			// _ = "end of CoverTab[7468]"
//line /usr/local/go/src/net/parse.go:192
		}
//line /usr/local/go/src/net/parse.go:192
		// _ = "end of CoverTab[7466]"
	}
//line /usr/local/go/src/net/parse.go:193
	// _ = "end of CoverTab[7462]"
//line /usr/local/go/src/net/parse.go:193
	_go_fuzz_dep_.CoverTab[7463]++
						return dst
//line /usr/local/go/src/net/parse.go:194
	// _ = "end of CoverTab[7463]"
}

// Number of occurrences of b in s.
func count(s string, b byte) int {
//line /usr/local/go/src/net/parse.go:198
	_go_fuzz_dep_.CoverTab[7469]++
						n := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:200
		_go_fuzz_dep_.CoverTab[7471]++
							if s[i] == b {
//line /usr/local/go/src/net/parse.go:201
			_go_fuzz_dep_.CoverTab[7472]++
								n++
//line /usr/local/go/src/net/parse.go:202
			// _ = "end of CoverTab[7472]"
		} else {
//line /usr/local/go/src/net/parse.go:203
			_go_fuzz_dep_.CoverTab[7473]++
//line /usr/local/go/src/net/parse.go:203
			// _ = "end of CoverTab[7473]"
//line /usr/local/go/src/net/parse.go:203
		}
//line /usr/local/go/src/net/parse.go:203
		// _ = "end of CoverTab[7471]"
	}
//line /usr/local/go/src/net/parse.go:204
	// _ = "end of CoverTab[7469]"
//line /usr/local/go/src/net/parse.go:204
	_go_fuzz_dep_.CoverTab[7470]++
						return n
//line /usr/local/go/src/net/parse.go:205
	// _ = "end of CoverTab[7470]"
}

// Index of rightmost occurrence of b in s.
func last(s string, b byte) int {
//line /usr/local/go/src/net/parse.go:209
	_go_fuzz_dep_.CoverTab[7474]++
						i := len(s)
						for i--; i >= 0; i-- {
//line /usr/local/go/src/net/parse.go:211
		_go_fuzz_dep_.CoverTab[7476]++
							if s[i] == b {
//line /usr/local/go/src/net/parse.go:212
			_go_fuzz_dep_.CoverTab[7477]++
								break
//line /usr/local/go/src/net/parse.go:213
			// _ = "end of CoverTab[7477]"
		} else {
//line /usr/local/go/src/net/parse.go:214
			_go_fuzz_dep_.CoverTab[7478]++
//line /usr/local/go/src/net/parse.go:214
			// _ = "end of CoverTab[7478]"
//line /usr/local/go/src/net/parse.go:214
		}
//line /usr/local/go/src/net/parse.go:214
		// _ = "end of CoverTab[7476]"
	}
//line /usr/local/go/src/net/parse.go:215
	// _ = "end of CoverTab[7474]"
//line /usr/local/go/src/net/parse.go:215
	_go_fuzz_dep_.CoverTab[7475]++
						return i
//line /usr/local/go/src/net/parse.go:216
	// _ = "end of CoverTab[7475]"
}

// hasUpperCase tells whether the given string contains at least one upper-case.
func hasUpperCase(s string) bool {
//line /usr/local/go/src/net/parse.go:220
	_go_fuzz_dep_.CoverTab[7479]++
						for i := range s {
//line /usr/local/go/src/net/parse.go:221
		_go_fuzz_dep_.CoverTab[7481]++
							if 'A' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[7482]++
//line /usr/local/go/src/net/parse.go:222
			return s[i] <= 'Z'
//line /usr/local/go/src/net/parse.go:222
			// _ = "end of CoverTab[7482]"
//line /usr/local/go/src/net/parse.go:222
		}() {
//line /usr/local/go/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[7483]++
								return true
//line /usr/local/go/src/net/parse.go:223
			// _ = "end of CoverTab[7483]"
		} else {
//line /usr/local/go/src/net/parse.go:224
			_go_fuzz_dep_.CoverTab[7484]++
//line /usr/local/go/src/net/parse.go:224
			// _ = "end of CoverTab[7484]"
//line /usr/local/go/src/net/parse.go:224
		}
//line /usr/local/go/src/net/parse.go:224
		// _ = "end of CoverTab[7481]"
	}
//line /usr/local/go/src/net/parse.go:225
	// _ = "end of CoverTab[7479]"
//line /usr/local/go/src/net/parse.go:225
	_go_fuzz_dep_.CoverTab[7480]++
						return false
//line /usr/local/go/src/net/parse.go:226
	// _ = "end of CoverTab[7480]"
}

// lowerASCIIBytes makes x ASCII lowercase in-place.
func lowerASCIIBytes(x []byte) {
//line /usr/local/go/src/net/parse.go:230
	_go_fuzz_dep_.CoverTab[7485]++
						for i, b := range x {
//line /usr/local/go/src/net/parse.go:231
		_go_fuzz_dep_.CoverTab[7486]++
							if 'A' <= b && func() bool {
//line /usr/local/go/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[7487]++
//line /usr/local/go/src/net/parse.go:232
			return b <= 'Z'
//line /usr/local/go/src/net/parse.go:232
			// _ = "end of CoverTab[7487]"
//line /usr/local/go/src/net/parse.go:232
		}() {
//line /usr/local/go/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[7488]++
								x[i] += 'a' - 'A'
//line /usr/local/go/src/net/parse.go:233
			// _ = "end of CoverTab[7488]"
		} else {
//line /usr/local/go/src/net/parse.go:234
			_go_fuzz_dep_.CoverTab[7489]++
//line /usr/local/go/src/net/parse.go:234
			// _ = "end of CoverTab[7489]"
//line /usr/local/go/src/net/parse.go:234
		}
//line /usr/local/go/src/net/parse.go:234
		// _ = "end of CoverTab[7486]"
	}
//line /usr/local/go/src/net/parse.go:235
	// _ = "end of CoverTab[7485]"
}

// lowerASCII returns the ASCII lowercase version of b.
func lowerASCII(b byte) byte {
//line /usr/local/go/src/net/parse.go:239
	_go_fuzz_dep_.CoverTab[7490]++
						if 'A' <= b && func() bool {
//line /usr/local/go/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[7492]++
//line /usr/local/go/src/net/parse.go:240
		return b <= 'Z'
//line /usr/local/go/src/net/parse.go:240
		// _ = "end of CoverTab[7492]"
//line /usr/local/go/src/net/parse.go:240
	}() {
//line /usr/local/go/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[7493]++
							return b + ('a' - 'A')
//line /usr/local/go/src/net/parse.go:241
		// _ = "end of CoverTab[7493]"
	} else {
//line /usr/local/go/src/net/parse.go:242
		_go_fuzz_dep_.CoverTab[7494]++
//line /usr/local/go/src/net/parse.go:242
		// _ = "end of CoverTab[7494]"
//line /usr/local/go/src/net/parse.go:242
	}
//line /usr/local/go/src/net/parse.go:242
	// _ = "end of CoverTab[7490]"
//line /usr/local/go/src/net/parse.go:242
	_go_fuzz_dep_.CoverTab[7491]++
						return b
//line /usr/local/go/src/net/parse.go:243
	// _ = "end of CoverTab[7491]"
}

// trimSpace returns x without any leading or trailing ASCII whitespace.
func trimSpace(x string) string {
//line /usr/local/go/src/net/parse.go:247
	_go_fuzz_dep_.CoverTab[7495]++
						for len(x) > 0 && func() bool {
//line /usr/local/go/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[7498]++
//line /usr/local/go/src/net/parse.go:248
		return isSpace(x[0])
//line /usr/local/go/src/net/parse.go:248
		// _ = "end of CoverTab[7498]"
//line /usr/local/go/src/net/parse.go:248
	}() {
//line /usr/local/go/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[7499]++
							x = x[1:]
//line /usr/local/go/src/net/parse.go:249
		// _ = "end of CoverTab[7499]"
	}
//line /usr/local/go/src/net/parse.go:250
	// _ = "end of CoverTab[7495]"
//line /usr/local/go/src/net/parse.go:250
	_go_fuzz_dep_.CoverTab[7496]++
						for len(x) > 0 && func() bool {
//line /usr/local/go/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[7500]++
//line /usr/local/go/src/net/parse.go:251
		return isSpace(x[len(x)-1])
//line /usr/local/go/src/net/parse.go:251
		// _ = "end of CoverTab[7500]"
//line /usr/local/go/src/net/parse.go:251
	}() {
//line /usr/local/go/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[7501]++
							x = x[:len(x)-1]
//line /usr/local/go/src/net/parse.go:252
		// _ = "end of CoverTab[7501]"
	}
//line /usr/local/go/src/net/parse.go:253
	// _ = "end of CoverTab[7496]"
//line /usr/local/go/src/net/parse.go:253
	_go_fuzz_dep_.CoverTab[7497]++
						return x
//line /usr/local/go/src/net/parse.go:254
	// _ = "end of CoverTab[7497]"
}

// isSpace reports whether b is an ASCII space character.
func isSpace(b byte) bool {
//line /usr/local/go/src/net/parse.go:258
	_go_fuzz_dep_.CoverTab[7502]++
						return b == ' ' || func() bool {
//line /usr/local/go/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[7503]++
//line /usr/local/go/src/net/parse.go:259
		return b == '\t'
//line /usr/local/go/src/net/parse.go:259
		// _ = "end of CoverTab[7503]"
//line /usr/local/go/src/net/parse.go:259
	}() || func() bool {
//line /usr/local/go/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[7504]++
//line /usr/local/go/src/net/parse.go:259
		return b == '\n'
//line /usr/local/go/src/net/parse.go:259
		// _ = "end of CoverTab[7504]"
//line /usr/local/go/src/net/parse.go:259
	}() || func() bool {
//line /usr/local/go/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[7505]++
//line /usr/local/go/src/net/parse.go:259
		return b == '\r'
//line /usr/local/go/src/net/parse.go:259
		// _ = "end of CoverTab[7505]"
//line /usr/local/go/src/net/parse.go:259
	}()
//line /usr/local/go/src/net/parse.go:259
	// _ = "end of CoverTab[7502]"
}

// removeComment returns line, removing any '#' byte and any following
//line /usr/local/go/src/net/parse.go:262
// bytes.
//line /usr/local/go/src/net/parse.go:264
func removeComment(line string) string {
//line /usr/local/go/src/net/parse.go:264
	_go_fuzz_dep_.CoverTab[7506]++
						if i := bytealg.IndexByteString(line, '#'); i != -1 {
//line /usr/local/go/src/net/parse.go:265
		_go_fuzz_dep_.CoverTab[7508]++
							return line[:i]
//line /usr/local/go/src/net/parse.go:266
		// _ = "end of CoverTab[7508]"
	} else {
//line /usr/local/go/src/net/parse.go:267
		_go_fuzz_dep_.CoverTab[7509]++
//line /usr/local/go/src/net/parse.go:267
		// _ = "end of CoverTab[7509]"
//line /usr/local/go/src/net/parse.go:267
	}
//line /usr/local/go/src/net/parse.go:267
	// _ = "end of CoverTab[7506]"
//line /usr/local/go/src/net/parse.go:267
	_go_fuzz_dep_.CoverTab[7507]++
						return line
//line /usr/local/go/src/net/parse.go:268
	// _ = "end of CoverTab[7507]"
}

// foreachField runs fn on each non-empty run of non-space bytes in x.
//line /usr/local/go/src/net/parse.go:271
// It returns the first non-nil error returned by fn.
//line /usr/local/go/src/net/parse.go:273
func foreachField(x string, fn func(field string) error) error {
//line /usr/local/go/src/net/parse.go:273
	_go_fuzz_dep_.CoverTab[7510]++
						x = trimSpace(x)
						for len(x) > 0 {
//line /usr/local/go/src/net/parse.go:275
		_go_fuzz_dep_.CoverTab[7512]++
							sp := bytealg.IndexByteString(x, ' ')
							if sp == -1 {
//line /usr/local/go/src/net/parse.go:277
			_go_fuzz_dep_.CoverTab[7515]++
								return fn(x)
//line /usr/local/go/src/net/parse.go:278
			// _ = "end of CoverTab[7515]"
		} else {
//line /usr/local/go/src/net/parse.go:279
			_go_fuzz_dep_.CoverTab[7516]++
//line /usr/local/go/src/net/parse.go:279
			// _ = "end of CoverTab[7516]"
//line /usr/local/go/src/net/parse.go:279
		}
//line /usr/local/go/src/net/parse.go:279
		// _ = "end of CoverTab[7512]"
//line /usr/local/go/src/net/parse.go:279
		_go_fuzz_dep_.CoverTab[7513]++
							if field := trimSpace(x[:sp]); len(field) > 0 {
//line /usr/local/go/src/net/parse.go:280
			_go_fuzz_dep_.CoverTab[7517]++
								if err := fn(field); err != nil {
//line /usr/local/go/src/net/parse.go:281
				_go_fuzz_dep_.CoverTab[7518]++
									return err
//line /usr/local/go/src/net/parse.go:282
				// _ = "end of CoverTab[7518]"
			} else {
//line /usr/local/go/src/net/parse.go:283
				_go_fuzz_dep_.CoverTab[7519]++
//line /usr/local/go/src/net/parse.go:283
				// _ = "end of CoverTab[7519]"
//line /usr/local/go/src/net/parse.go:283
			}
//line /usr/local/go/src/net/parse.go:283
			// _ = "end of CoverTab[7517]"
		} else {
//line /usr/local/go/src/net/parse.go:284
			_go_fuzz_dep_.CoverTab[7520]++
//line /usr/local/go/src/net/parse.go:284
			// _ = "end of CoverTab[7520]"
//line /usr/local/go/src/net/parse.go:284
		}
//line /usr/local/go/src/net/parse.go:284
		// _ = "end of CoverTab[7513]"
//line /usr/local/go/src/net/parse.go:284
		_go_fuzz_dep_.CoverTab[7514]++
							x = trimSpace(x[sp+1:])
//line /usr/local/go/src/net/parse.go:285
		// _ = "end of CoverTab[7514]"
	}
//line /usr/local/go/src/net/parse.go:286
	// _ = "end of CoverTab[7510]"
//line /usr/local/go/src/net/parse.go:286
	_go_fuzz_dep_.CoverTab[7511]++
						return nil
//line /usr/local/go/src/net/parse.go:287
	// _ = "end of CoverTab[7511]"
}

// stringsHasSuffix is strings.HasSuffix. It reports whether s ends in
//line /usr/local/go/src/net/parse.go:290
// suffix.
//line /usr/local/go/src/net/parse.go:292
func stringsHasSuffix(s, suffix string) bool {
//line /usr/local/go/src/net/parse.go:292
	_go_fuzz_dep_.CoverTab[7521]++
						return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/net/parse.go:293
		_go_fuzz_dep_.CoverTab[7522]++
//line /usr/local/go/src/net/parse.go:293
		return s[len(s)-len(suffix):] == suffix
//line /usr/local/go/src/net/parse.go:293
		// _ = "end of CoverTab[7522]"
//line /usr/local/go/src/net/parse.go:293
	}()
//line /usr/local/go/src/net/parse.go:293
	// _ = "end of CoverTab[7521]"
}

// stringsHasSuffixFold reports whether s ends in suffix,
//line /usr/local/go/src/net/parse.go:296
// ASCII-case-insensitively.
//line /usr/local/go/src/net/parse.go:298
func stringsHasSuffixFold(s, suffix string) bool {
//line /usr/local/go/src/net/parse.go:298
	_go_fuzz_dep_.CoverTab[7523]++
						return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/net/parse.go:299
		_go_fuzz_dep_.CoverTab[7524]++
//line /usr/local/go/src/net/parse.go:299
		return stringsEqualFold(s[len(s)-len(suffix):], suffix)
//line /usr/local/go/src/net/parse.go:299
		// _ = "end of CoverTab[7524]"
//line /usr/local/go/src/net/parse.go:299
	}()
//line /usr/local/go/src/net/parse.go:299
	// _ = "end of CoverTab[7523]"
}

// stringsHasPrefix is strings.HasPrefix. It reports whether s begins with prefix.
func stringsHasPrefix(s, prefix string) bool {
//line /usr/local/go/src/net/parse.go:303
	_go_fuzz_dep_.CoverTab[7525]++
						return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/net/parse.go:304
		_go_fuzz_dep_.CoverTab[7526]++
//line /usr/local/go/src/net/parse.go:304
		return s[:len(prefix)] == prefix
//line /usr/local/go/src/net/parse.go:304
		// _ = "end of CoverTab[7526]"
//line /usr/local/go/src/net/parse.go:304
	}()
//line /usr/local/go/src/net/parse.go:304
	// _ = "end of CoverTab[7525]"
}

// stringsEqualFold is strings.EqualFold, ASCII only. It reports whether s and t
//line /usr/local/go/src/net/parse.go:307
// are equal, ASCII-case-insensitively.
//line /usr/local/go/src/net/parse.go:309
func stringsEqualFold(s, t string) bool {
//line /usr/local/go/src/net/parse.go:309
	_go_fuzz_dep_.CoverTab[7527]++
						if len(s) != len(t) {
//line /usr/local/go/src/net/parse.go:310
		_go_fuzz_dep_.CoverTab[7530]++
							return false
//line /usr/local/go/src/net/parse.go:311
		// _ = "end of CoverTab[7530]"
	} else {
//line /usr/local/go/src/net/parse.go:312
		_go_fuzz_dep_.CoverTab[7531]++
//line /usr/local/go/src/net/parse.go:312
		// _ = "end of CoverTab[7531]"
//line /usr/local/go/src/net/parse.go:312
	}
//line /usr/local/go/src/net/parse.go:312
	// _ = "end of CoverTab[7527]"
//line /usr/local/go/src/net/parse.go:312
	_go_fuzz_dep_.CoverTab[7528]++
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:313
		_go_fuzz_dep_.CoverTab[7532]++
							if lowerASCII(s[i]) != lowerASCII(t[i]) {
//line /usr/local/go/src/net/parse.go:314
			_go_fuzz_dep_.CoverTab[7533]++
								return false
//line /usr/local/go/src/net/parse.go:315
			// _ = "end of CoverTab[7533]"
		} else {
//line /usr/local/go/src/net/parse.go:316
			_go_fuzz_dep_.CoverTab[7534]++
//line /usr/local/go/src/net/parse.go:316
			// _ = "end of CoverTab[7534]"
//line /usr/local/go/src/net/parse.go:316
		}
//line /usr/local/go/src/net/parse.go:316
		// _ = "end of CoverTab[7532]"
	}
//line /usr/local/go/src/net/parse.go:317
	// _ = "end of CoverTab[7528]"
//line /usr/local/go/src/net/parse.go:317
	_go_fuzz_dep_.CoverTab[7529]++
						return true
//line /usr/local/go/src/net/parse.go:318
	// _ = "end of CoverTab[7529]"
}

//line /usr/local/go/src/net/parse.go:319
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/parse.go:319
var _ = _go_fuzz_dep_.CoverTab
