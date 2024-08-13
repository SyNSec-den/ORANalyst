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
	_go_fuzz_dep_.CoverTab[15765]++
//line /usr/local/go/src/net/parse.go:23
	f.file.Close()
//line /usr/local/go/src/net/parse.go:23
	// _ = "end of CoverTab[15765]"
//line /usr/local/go/src/net/parse.go:23
}

func (f *file) getLineFromData() (s string, ok bool) {
//line /usr/local/go/src/net/parse.go:25
	_go_fuzz_dep_.CoverTab[15766]++
						data := f.data
						i := 0
						for i = 0; i < len(data); i++ {
//line /usr/local/go/src/net/parse.go:28
		_go_fuzz_dep_.CoverTab[15769]++
							if data[i] == '\n' {
//line /usr/local/go/src/net/parse.go:29
			_go_fuzz_dep_.CoverTab[15770]++
								s = string(data[0:i])
								ok = true

								i++
								n := len(data) - i
								copy(data[0:], data[i:])
								f.data = data[0:n]
								return
//line /usr/local/go/src/net/parse.go:37
			// _ = "end of CoverTab[15770]"
		} else {
//line /usr/local/go/src/net/parse.go:38
			_go_fuzz_dep_.CoverTab[15771]++
//line /usr/local/go/src/net/parse.go:38
			// _ = "end of CoverTab[15771]"
//line /usr/local/go/src/net/parse.go:38
		}
//line /usr/local/go/src/net/parse.go:38
		// _ = "end of CoverTab[15769]"
	}
//line /usr/local/go/src/net/parse.go:39
	// _ = "end of CoverTab[15766]"
//line /usr/local/go/src/net/parse.go:39
	_go_fuzz_dep_.CoverTab[15767]++
						if f.atEOF && func() bool {
//line /usr/local/go/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[15772]++
//line /usr/local/go/src/net/parse.go:40
		return len(f.data) > 0
//line /usr/local/go/src/net/parse.go:40
		// _ = "end of CoverTab[15772]"
//line /usr/local/go/src/net/parse.go:40
	}() {
//line /usr/local/go/src/net/parse.go:40
		_go_fuzz_dep_.CoverTab[15773]++

							s = string(data)
							f.data = f.data[0:0]
							ok = true
//line /usr/local/go/src/net/parse.go:44
		// _ = "end of CoverTab[15773]"
	} else {
//line /usr/local/go/src/net/parse.go:45
		_go_fuzz_dep_.CoverTab[15774]++
//line /usr/local/go/src/net/parse.go:45
		// _ = "end of CoverTab[15774]"
//line /usr/local/go/src/net/parse.go:45
	}
//line /usr/local/go/src/net/parse.go:45
	// _ = "end of CoverTab[15767]"
//line /usr/local/go/src/net/parse.go:45
	_go_fuzz_dep_.CoverTab[15768]++
						return
//line /usr/local/go/src/net/parse.go:46
	// _ = "end of CoverTab[15768]"
}

func (f *file) readLine() (s string, ok bool) {
//line /usr/local/go/src/net/parse.go:49
	_go_fuzz_dep_.CoverTab[15775]++
						if s, ok = f.getLineFromData(); ok {
//line /usr/local/go/src/net/parse.go:50
		_go_fuzz_dep_.CoverTab[15778]++
							return
//line /usr/local/go/src/net/parse.go:51
		// _ = "end of CoverTab[15778]"
	} else {
//line /usr/local/go/src/net/parse.go:52
		_go_fuzz_dep_.CoverTab[15779]++
//line /usr/local/go/src/net/parse.go:52
		// _ = "end of CoverTab[15779]"
//line /usr/local/go/src/net/parse.go:52
	}
//line /usr/local/go/src/net/parse.go:52
	// _ = "end of CoverTab[15775]"
//line /usr/local/go/src/net/parse.go:52
	_go_fuzz_dep_.CoverTab[15776]++
						if len(f.data) < cap(f.data) {
//line /usr/local/go/src/net/parse.go:53
		_go_fuzz_dep_.CoverTab[15780]++
							ln := len(f.data)
							n, err := io.ReadFull(f.file, f.data[ln:cap(f.data)])
							if n >= 0 {
//line /usr/local/go/src/net/parse.go:56
			_go_fuzz_dep_.CoverTab[15782]++
								f.data = f.data[0 : ln+n]
//line /usr/local/go/src/net/parse.go:57
			// _ = "end of CoverTab[15782]"
		} else {
//line /usr/local/go/src/net/parse.go:58
			_go_fuzz_dep_.CoverTab[15783]++
//line /usr/local/go/src/net/parse.go:58
			// _ = "end of CoverTab[15783]"
//line /usr/local/go/src/net/parse.go:58
		}
//line /usr/local/go/src/net/parse.go:58
		// _ = "end of CoverTab[15780]"
//line /usr/local/go/src/net/parse.go:58
		_go_fuzz_dep_.CoverTab[15781]++
							if err == io.EOF || func() bool {
//line /usr/local/go/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[15784]++
//line /usr/local/go/src/net/parse.go:59
			return err == io.ErrUnexpectedEOF
//line /usr/local/go/src/net/parse.go:59
			// _ = "end of CoverTab[15784]"
//line /usr/local/go/src/net/parse.go:59
		}() {
//line /usr/local/go/src/net/parse.go:59
			_go_fuzz_dep_.CoverTab[15785]++
								f.atEOF = true
//line /usr/local/go/src/net/parse.go:60
			// _ = "end of CoverTab[15785]"
		} else {
//line /usr/local/go/src/net/parse.go:61
			_go_fuzz_dep_.CoverTab[15786]++
//line /usr/local/go/src/net/parse.go:61
			// _ = "end of CoverTab[15786]"
//line /usr/local/go/src/net/parse.go:61
		}
//line /usr/local/go/src/net/parse.go:61
		// _ = "end of CoverTab[15781]"
	} else {
//line /usr/local/go/src/net/parse.go:62
		_go_fuzz_dep_.CoverTab[15787]++
//line /usr/local/go/src/net/parse.go:62
		// _ = "end of CoverTab[15787]"
//line /usr/local/go/src/net/parse.go:62
	}
//line /usr/local/go/src/net/parse.go:62
	// _ = "end of CoverTab[15776]"
//line /usr/local/go/src/net/parse.go:62
	_go_fuzz_dep_.CoverTab[15777]++
						s, ok = f.getLineFromData()
						return
//line /usr/local/go/src/net/parse.go:64
	// _ = "end of CoverTab[15777]"
}

func (f *file) stat() (mtime time.Time, size int64, err error) {
//line /usr/local/go/src/net/parse.go:67
	_go_fuzz_dep_.CoverTab[15788]++
						st, err := f.file.Stat()
						if err != nil {
//line /usr/local/go/src/net/parse.go:69
		_go_fuzz_dep_.CoverTab[15790]++
							return time.Time{}, 0, err
//line /usr/local/go/src/net/parse.go:70
		// _ = "end of CoverTab[15790]"
	} else {
//line /usr/local/go/src/net/parse.go:71
		_go_fuzz_dep_.CoverTab[15791]++
//line /usr/local/go/src/net/parse.go:71
		// _ = "end of CoverTab[15791]"
//line /usr/local/go/src/net/parse.go:71
	}
//line /usr/local/go/src/net/parse.go:71
	// _ = "end of CoverTab[15788]"
//line /usr/local/go/src/net/parse.go:71
	_go_fuzz_dep_.CoverTab[15789]++
						return st.ModTime(), st.Size(), nil
//line /usr/local/go/src/net/parse.go:72
	// _ = "end of CoverTab[15789]"
}

func open(name string) (*file, error) {
//line /usr/local/go/src/net/parse.go:75
	_go_fuzz_dep_.CoverTab[15792]++
						fd, err := os.Open(name)
						if err != nil {
//line /usr/local/go/src/net/parse.go:77
		_go_fuzz_dep_.CoverTab[15794]++
							return nil, err
//line /usr/local/go/src/net/parse.go:78
		// _ = "end of CoverTab[15794]"
	} else {
//line /usr/local/go/src/net/parse.go:79
		_go_fuzz_dep_.CoverTab[15795]++
//line /usr/local/go/src/net/parse.go:79
		// _ = "end of CoverTab[15795]"
//line /usr/local/go/src/net/parse.go:79
	}
//line /usr/local/go/src/net/parse.go:79
	// _ = "end of CoverTab[15792]"
//line /usr/local/go/src/net/parse.go:79
	_go_fuzz_dep_.CoverTab[15793]++
						return &file{fd, make([]byte, 0, 64*1024), false}, nil
//line /usr/local/go/src/net/parse.go:80
	// _ = "end of CoverTab[15793]"
}

func stat(name string) (mtime time.Time, size int64, err error) {
//line /usr/local/go/src/net/parse.go:83
	_go_fuzz_dep_.CoverTab[15796]++
						st, err := os.Stat(name)
						if err != nil {
//line /usr/local/go/src/net/parse.go:85
		_go_fuzz_dep_.CoverTab[15798]++
							return time.Time{}, 0, err
//line /usr/local/go/src/net/parse.go:86
		// _ = "end of CoverTab[15798]"
	} else {
//line /usr/local/go/src/net/parse.go:87
		_go_fuzz_dep_.CoverTab[15799]++
//line /usr/local/go/src/net/parse.go:87
		// _ = "end of CoverTab[15799]"
//line /usr/local/go/src/net/parse.go:87
	}
//line /usr/local/go/src/net/parse.go:87
	// _ = "end of CoverTab[15796]"
//line /usr/local/go/src/net/parse.go:87
	_go_fuzz_dep_.CoverTab[15797]++
						return st.ModTime(), st.Size(), nil
//line /usr/local/go/src/net/parse.go:88
	// _ = "end of CoverTab[15797]"
}

// Count occurrences in s of any bytes in t.
func countAnyByte(s string, t string) int {
//line /usr/local/go/src/net/parse.go:92
	_go_fuzz_dep_.CoverTab[15800]++
						n := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:94
		_go_fuzz_dep_.CoverTab[15802]++
							if bytealg.IndexByteString(t, s[i]) >= 0 {
//line /usr/local/go/src/net/parse.go:95
			_go_fuzz_dep_.CoverTab[15803]++
								n++
//line /usr/local/go/src/net/parse.go:96
			// _ = "end of CoverTab[15803]"
		} else {
//line /usr/local/go/src/net/parse.go:97
			_go_fuzz_dep_.CoverTab[15804]++
//line /usr/local/go/src/net/parse.go:97
			// _ = "end of CoverTab[15804]"
//line /usr/local/go/src/net/parse.go:97
		}
//line /usr/local/go/src/net/parse.go:97
		// _ = "end of CoverTab[15802]"
	}
//line /usr/local/go/src/net/parse.go:98
	// _ = "end of CoverTab[15800]"
//line /usr/local/go/src/net/parse.go:98
	_go_fuzz_dep_.CoverTab[15801]++
						return n
//line /usr/local/go/src/net/parse.go:99
	// _ = "end of CoverTab[15801]"
}

// Split s at any bytes in t.
func splitAtBytes(s string, t string) []string {
//line /usr/local/go/src/net/parse.go:103
	_go_fuzz_dep_.CoverTab[15805]++
						a := make([]string, 1+countAnyByte(s, t))
						n := 0
						last := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:107
		_go_fuzz_dep_.CoverTab[15808]++
							if bytealg.IndexByteString(t, s[i]) >= 0 {
//line /usr/local/go/src/net/parse.go:108
			_go_fuzz_dep_.CoverTab[15809]++
								if last < i {
//line /usr/local/go/src/net/parse.go:109
				_go_fuzz_dep_.CoverTab[15811]++
									a[n] = s[last:i]
									n++
//line /usr/local/go/src/net/parse.go:111
				// _ = "end of CoverTab[15811]"
			} else {
//line /usr/local/go/src/net/parse.go:112
				_go_fuzz_dep_.CoverTab[15812]++
//line /usr/local/go/src/net/parse.go:112
				// _ = "end of CoverTab[15812]"
//line /usr/local/go/src/net/parse.go:112
			}
//line /usr/local/go/src/net/parse.go:112
			// _ = "end of CoverTab[15809]"
//line /usr/local/go/src/net/parse.go:112
			_go_fuzz_dep_.CoverTab[15810]++
								last = i + 1
//line /usr/local/go/src/net/parse.go:113
			// _ = "end of CoverTab[15810]"
		} else {
//line /usr/local/go/src/net/parse.go:114
			_go_fuzz_dep_.CoverTab[15813]++
//line /usr/local/go/src/net/parse.go:114
			// _ = "end of CoverTab[15813]"
//line /usr/local/go/src/net/parse.go:114
		}
//line /usr/local/go/src/net/parse.go:114
		// _ = "end of CoverTab[15808]"
	}
//line /usr/local/go/src/net/parse.go:115
	// _ = "end of CoverTab[15805]"
//line /usr/local/go/src/net/parse.go:115
	_go_fuzz_dep_.CoverTab[15806]++
						if last < len(s) {
//line /usr/local/go/src/net/parse.go:116
		_go_fuzz_dep_.CoverTab[15814]++
							a[n] = s[last:]
							n++
//line /usr/local/go/src/net/parse.go:118
		// _ = "end of CoverTab[15814]"
	} else {
//line /usr/local/go/src/net/parse.go:119
		_go_fuzz_dep_.CoverTab[15815]++
//line /usr/local/go/src/net/parse.go:119
		// _ = "end of CoverTab[15815]"
//line /usr/local/go/src/net/parse.go:119
	}
//line /usr/local/go/src/net/parse.go:119
	// _ = "end of CoverTab[15806]"
//line /usr/local/go/src/net/parse.go:119
	_go_fuzz_dep_.CoverTab[15807]++
						return a[0:n]
//line /usr/local/go/src/net/parse.go:120
	// _ = "end of CoverTab[15807]"
}

func getFields(s string) []string {
//line /usr/local/go/src/net/parse.go:123
	_go_fuzz_dep_.CoverTab[15816]++
//line /usr/local/go/src/net/parse.go:123
	return splitAtBytes(s, " \r\t\n")
//line /usr/local/go/src/net/parse.go:123
	// _ = "end of CoverTab[15816]"
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
	_go_fuzz_dep_.CoverTab[15817]++
						n = 0
						for i = 0; i < len(s) && func() bool {
//line /usr/local/go/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[15820]++
//line /usr/local/go/src/net/parse.go:132
		return '0' <= s[i]
//line /usr/local/go/src/net/parse.go:132
		// _ = "end of CoverTab[15820]"
//line /usr/local/go/src/net/parse.go:132
	}() && func() bool {
//line /usr/local/go/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[15821]++
//line /usr/local/go/src/net/parse.go:132
		return s[i] <= '9'
//line /usr/local/go/src/net/parse.go:132
		// _ = "end of CoverTab[15821]"
//line /usr/local/go/src/net/parse.go:132
	}(); i++ {
//line /usr/local/go/src/net/parse.go:132
		_go_fuzz_dep_.CoverTab[15822]++
							n = n*10 + int(s[i]-'0')
							if n >= big {
//line /usr/local/go/src/net/parse.go:134
			_go_fuzz_dep_.CoverTab[15823]++
								return big, i, false
//line /usr/local/go/src/net/parse.go:135
			// _ = "end of CoverTab[15823]"
		} else {
//line /usr/local/go/src/net/parse.go:136
			_go_fuzz_dep_.CoverTab[15824]++
//line /usr/local/go/src/net/parse.go:136
			// _ = "end of CoverTab[15824]"
//line /usr/local/go/src/net/parse.go:136
		}
//line /usr/local/go/src/net/parse.go:136
		// _ = "end of CoverTab[15822]"
	}
//line /usr/local/go/src/net/parse.go:137
	// _ = "end of CoverTab[15817]"
//line /usr/local/go/src/net/parse.go:137
	_go_fuzz_dep_.CoverTab[15818]++
						if i == 0 {
//line /usr/local/go/src/net/parse.go:138
		_go_fuzz_dep_.CoverTab[15825]++
							return 0, 0, false
//line /usr/local/go/src/net/parse.go:139
		// _ = "end of CoverTab[15825]"
	} else {
//line /usr/local/go/src/net/parse.go:140
		_go_fuzz_dep_.CoverTab[15826]++
//line /usr/local/go/src/net/parse.go:140
		// _ = "end of CoverTab[15826]"
//line /usr/local/go/src/net/parse.go:140
	}
//line /usr/local/go/src/net/parse.go:140
	// _ = "end of CoverTab[15818]"
//line /usr/local/go/src/net/parse.go:140
	_go_fuzz_dep_.CoverTab[15819]++
						return n, i, true
//line /usr/local/go/src/net/parse.go:141
	// _ = "end of CoverTab[15819]"
}

// Hexadecimal to integer.
//line /usr/local/go/src/net/parse.go:144
// Returns number, characters consumed, success.
//line /usr/local/go/src/net/parse.go:146
func xtoi(s string) (n int, i int, ok bool) {
//line /usr/local/go/src/net/parse.go:146
	_go_fuzz_dep_.CoverTab[15827]++
						n = 0
						for i = 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:148
		_go_fuzz_dep_.CoverTab[15830]++
							if '0' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[15832]++
//line /usr/local/go/src/net/parse.go:149
			return s[i] <= '9'
//line /usr/local/go/src/net/parse.go:149
			// _ = "end of CoverTab[15832]"
//line /usr/local/go/src/net/parse.go:149
		}() {
//line /usr/local/go/src/net/parse.go:149
			_go_fuzz_dep_.CoverTab[15833]++
								n *= 16
								n += int(s[i] - '0')
//line /usr/local/go/src/net/parse.go:151
			// _ = "end of CoverTab[15833]"
		} else {
//line /usr/local/go/src/net/parse.go:152
			_go_fuzz_dep_.CoverTab[15834]++
//line /usr/local/go/src/net/parse.go:152
			if 'a' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[15835]++
//line /usr/local/go/src/net/parse.go:152
				return s[i] <= 'f'
//line /usr/local/go/src/net/parse.go:152
				// _ = "end of CoverTab[15835]"
//line /usr/local/go/src/net/parse.go:152
			}() {
//line /usr/local/go/src/net/parse.go:152
				_go_fuzz_dep_.CoverTab[15836]++
									n *= 16
									n += int(s[i]-'a') + 10
//line /usr/local/go/src/net/parse.go:154
				// _ = "end of CoverTab[15836]"
			} else {
//line /usr/local/go/src/net/parse.go:155
				_go_fuzz_dep_.CoverTab[15837]++
//line /usr/local/go/src/net/parse.go:155
				if 'A' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[15838]++
//line /usr/local/go/src/net/parse.go:155
					return s[i] <= 'F'
//line /usr/local/go/src/net/parse.go:155
					// _ = "end of CoverTab[15838]"
//line /usr/local/go/src/net/parse.go:155
				}() {
//line /usr/local/go/src/net/parse.go:155
					_go_fuzz_dep_.CoverTab[15839]++
										n *= 16
										n += int(s[i]-'A') + 10
//line /usr/local/go/src/net/parse.go:157
					// _ = "end of CoverTab[15839]"
				} else {
//line /usr/local/go/src/net/parse.go:158
					_go_fuzz_dep_.CoverTab[15840]++
										break
//line /usr/local/go/src/net/parse.go:159
					// _ = "end of CoverTab[15840]"
				}
//line /usr/local/go/src/net/parse.go:160
				// _ = "end of CoverTab[15837]"
//line /usr/local/go/src/net/parse.go:160
			}
//line /usr/local/go/src/net/parse.go:160
			// _ = "end of CoverTab[15834]"
//line /usr/local/go/src/net/parse.go:160
		}
//line /usr/local/go/src/net/parse.go:160
		// _ = "end of CoverTab[15830]"
//line /usr/local/go/src/net/parse.go:160
		_go_fuzz_dep_.CoverTab[15831]++
							if n >= big {
//line /usr/local/go/src/net/parse.go:161
			_go_fuzz_dep_.CoverTab[15841]++
								return 0, i, false
//line /usr/local/go/src/net/parse.go:162
			// _ = "end of CoverTab[15841]"
		} else {
//line /usr/local/go/src/net/parse.go:163
			_go_fuzz_dep_.CoverTab[15842]++
//line /usr/local/go/src/net/parse.go:163
			// _ = "end of CoverTab[15842]"
//line /usr/local/go/src/net/parse.go:163
		}
//line /usr/local/go/src/net/parse.go:163
		// _ = "end of CoverTab[15831]"
	}
//line /usr/local/go/src/net/parse.go:164
	// _ = "end of CoverTab[15827]"
//line /usr/local/go/src/net/parse.go:164
	_go_fuzz_dep_.CoverTab[15828]++
						if i == 0 {
//line /usr/local/go/src/net/parse.go:165
		_go_fuzz_dep_.CoverTab[15843]++
							return 0, i, false
//line /usr/local/go/src/net/parse.go:166
		// _ = "end of CoverTab[15843]"
	} else {
//line /usr/local/go/src/net/parse.go:167
		_go_fuzz_dep_.CoverTab[15844]++
//line /usr/local/go/src/net/parse.go:167
		// _ = "end of CoverTab[15844]"
//line /usr/local/go/src/net/parse.go:167
	}
//line /usr/local/go/src/net/parse.go:167
	// _ = "end of CoverTab[15828]"
//line /usr/local/go/src/net/parse.go:167
	_go_fuzz_dep_.CoverTab[15829]++
						return n, i, true
//line /usr/local/go/src/net/parse.go:168
	// _ = "end of CoverTab[15829]"
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
	_go_fuzz_dep_.CoverTab[15845]++
						if len(s) > 2 && func() bool {
//line /usr/local/go/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[15847]++
//line /usr/local/go/src/net/parse.go:176
		return s[2] != e
//line /usr/local/go/src/net/parse.go:176
		// _ = "end of CoverTab[15847]"
//line /usr/local/go/src/net/parse.go:176
	}() {
//line /usr/local/go/src/net/parse.go:176
		_go_fuzz_dep_.CoverTab[15848]++
							return 0, false
//line /usr/local/go/src/net/parse.go:177
		// _ = "end of CoverTab[15848]"
	} else {
//line /usr/local/go/src/net/parse.go:178
		_go_fuzz_dep_.CoverTab[15849]++
//line /usr/local/go/src/net/parse.go:178
		// _ = "end of CoverTab[15849]"
//line /usr/local/go/src/net/parse.go:178
	}
//line /usr/local/go/src/net/parse.go:178
	// _ = "end of CoverTab[15845]"
//line /usr/local/go/src/net/parse.go:178
	_go_fuzz_dep_.CoverTab[15846]++
						n, ei, ok := xtoi(s[:2])
						return byte(n), ok && func() bool {
//line /usr/local/go/src/net/parse.go:180
		_go_fuzz_dep_.CoverTab[15850]++
//line /usr/local/go/src/net/parse.go:180
		return ei == 2
//line /usr/local/go/src/net/parse.go:180
		// _ = "end of CoverTab[15850]"
//line /usr/local/go/src/net/parse.go:180
	}()
//line /usr/local/go/src/net/parse.go:180
	// _ = "end of CoverTab[15846]"
}

// Convert i to a hexadecimal string. Leading zeros are not printed.
func appendHex(dst []byte, i uint32) []byte {
//line /usr/local/go/src/net/parse.go:184
	_go_fuzz_dep_.CoverTab[15851]++
						if i == 0 {
//line /usr/local/go/src/net/parse.go:185
		_go_fuzz_dep_.CoverTab[15854]++
							return append(dst, '0')
//line /usr/local/go/src/net/parse.go:186
		// _ = "end of CoverTab[15854]"
	} else {
//line /usr/local/go/src/net/parse.go:187
		_go_fuzz_dep_.CoverTab[15855]++
//line /usr/local/go/src/net/parse.go:187
		// _ = "end of CoverTab[15855]"
//line /usr/local/go/src/net/parse.go:187
	}
//line /usr/local/go/src/net/parse.go:187
	// _ = "end of CoverTab[15851]"
//line /usr/local/go/src/net/parse.go:187
	_go_fuzz_dep_.CoverTab[15852]++
						for j := 7; j >= 0; j-- {
//line /usr/local/go/src/net/parse.go:188
		_go_fuzz_dep_.CoverTab[15856]++
							v := i >> uint(j*4)
							if v > 0 {
//line /usr/local/go/src/net/parse.go:190
			_go_fuzz_dep_.CoverTab[15857]++
								dst = append(dst, hexDigit[v&0xf])
//line /usr/local/go/src/net/parse.go:191
			// _ = "end of CoverTab[15857]"
		} else {
//line /usr/local/go/src/net/parse.go:192
			_go_fuzz_dep_.CoverTab[15858]++
//line /usr/local/go/src/net/parse.go:192
			// _ = "end of CoverTab[15858]"
//line /usr/local/go/src/net/parse.go:192
		}
//line /usr/local/go/src/net/parse.go:192
		// _ = "end of CoverTab[15856]"
	}
//line /usr/local/go/src/net/parse.go:193
	// _ = "end of CoverTab[15852]"
//line /usr/local/go/src/net/parse.go:193
	_go_fuzz_dep_.CoverTab[15853]++
						return dst
//line /usr/local/go/src/net/parse.go:194
	// _ = "end of CoverTab[15853]"
}

// Number of occurrences of b in s.
func count(s string, b byte) int {
//line /usr/local/go/src/net/parse.go:198
	_go_fuzz_dep_.CoverTab[15859]++
						n := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:200
		_go_fuzz_dep_.CoverTab[15861]++
							if s[i] == b {
//line /usr/local/go/src/net/parse.go:201
			_go_fuzz_dep_.CoverTab[15862]++
								n++
//line /usr/local/go/src/net/parse.go:202
			// _ = "end of CoverTab[15862]"
		} else {
//line /usr/local/go/src/net/parse.go:203
			_go_fuzz_dep_.CoverTab[15863]++
//line /usr/local/go/src/net/parse.go:203
			// _ = "end of CoverTab[15863]"
//line /usr/local/go/src/net/parse.go:203
		}
//line /usr/local/go/src/net/parse.go:203
		// _ = "end of CoverTab[15861]"
	}
//line /usr/local/go/src/net/parse.go:204
	// _ = "end of CoverTab[15859]"
//line /usr/local/go/src/net/parse.go:204
	_go_fuzz_dep_.CoverTab[15860]++
						return n
//line /usr/local/go/src/net/parse.go:205
	// _ = "end of CoverTab[15860]"
}

// Index of rightmost occurrence of b in s.
func last(s string, b byte) int {
//line /usr/local/go/src/net/parse.go:209
	_go_fuzz_dep_.CoverTab[15864]++
						i := len(s)
						for i--; i >= 0; i-- {
//line /usr/local/go/src/net/parse.go:211
		_go_fuzz_dep_.CoverTab[15866]++
							if s[i] == b {
//line /usr/local/go/src/net/parse.go:212
			_go_fuzz_dep_.CoverTab[15867]++
								break
//line /usr/local/go/src/net/parse.go:213
			// _ = "end of CoverTab[15867]"
		} else {
//line /usr/local/go/src/net/parse.go:214
			_go_fuzz_dep_.CoverTab[15868]++
//line /usr/local/go/src/net/parse.go:214
			// _ = "end of CoverTab[15868]"
//line /usr/local/go/src/net/parse.go:214
		}
//line /usr/local/go/src/net/parse.go:214
		// _ = "end of CoverTab[15866]"
	}
//line /usr/local/go/src/net/parse.go:215
	// _ = "end of CoverTab[15864]"
//line /usr/local/go/src/net/parse.go:215
	_go_fuzz_dep_.CoverTab[15865]++
						return i
//line /usr/local/go/src/net/parse.go:216
	// _ = "end of CoverTab[15865]"
}

// hasUpperCase tells whether the given string contains at least one upper-case.
func hasUpperCase(s string) bool {
//line /usr/local/go/src/net/parse.go:220
	_go_fuzz_dep_.CoverTab[15869]++
						for i := range s {
//line /usr/local/go/src/net/parse.go:221
		_go_fuzz_dep_.CoverTab[15871]++
							if 'A' <= s[i] && func() bool {
//line /usr/local/go/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[15872]++
//line /usr/local/go/src/net/parse.go:222
			return s[i] <= 'Z'
//line /usr/local/go/src/net/parse.go:222
			// _ = "end of CoverTab[15872]"
//line /usr/local/go/src/net/parse.go:222
		}() {
//line /usr/local/go/src/net/parse.go:222
			_go_fuzz_dep_.CoverTab[15873]++
								return true
//line /usr/local/go/src/net/parse.go:223
			// _ = "end of CoverTab[15873]"
		} else {
//line /usr/local/go/src/net/parse.go:224
			_go_fuzz_dep_.CoverTab[15874]++
//line /usr/local/go/src/net/parse.go:224
			// _ = "end of CoverTab[15874]"
//line /usr/local/go/src/net/parse.go:224
		}
//line /usr/local/go/src/net/parse.go:224
		// _ = "end of CoverTab[15871]"
	}
//line /usr/local/go/src/net/parse.go:225
	// _ = "end of CoverTab[15869]"
//line /usr/local/go/src/net/parse.go:225
	_go_fuzz_dep_.CoverTab[15870]++
						return false
//line /usr/local/go/src/net/parse.go:226
	// _ = "end of CoverTab[15870]"
}

// lowerASCIIBytes makes x ASCII lowercase in-place.
func lowerASCIIBytes(x []byte) {
//line /usr/local/go/src/net/parse.go:230
	_go_fuzz_dep_.CoverTab[15875]++
						for i, b := range x {
//line /usr/local/go/src/net/parse.go:231
		_go_fuzz_dep_.CoverTab[15876]++
							if 'A' <= b && func() bool {
//line /usr/local/go/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[15877]++
//line /usr/local/go/src/net/parse.go:232
			return b <= 'Z'
//line /usr/local/go/src/net/parse.go:232
			// _ = "end of CoverTab[15877]"
//line /usr/local/go/src/net/parse.go:232
		}() {
//line /usr/local/go/src/net/parse.go:232
			_go_fuzz_dep_.CoverTab[15878]++
								x[i] += 'a' - 'A'
//line /usr/local/go/src/net/parse.go:233
			// _ = "end of CoverTab[15878]"
		} else {
//line /usr/local/go/src/net/parse.go:234
			_go_fuzz_dep_.CoverTab[15879]++
//line /usr/local/go/src/net/parse.go:234
			// _ = "end of CoverTab[15879]"
//line /usr/local/go/src/net/parse.go:234
		}
//line /usr/local/go/src/net/parse.go:234
		// _ = "end of CoverTab[15876]"
	}
//line /usr/local/go/src/net/parse.go:235
	// _ = "end of CoverTab[15875]"
}

// lowerASCII returns the ASCII lowercase version of b.
func lowerASCII(b byte) byte {
//line /usr/local/go/src/net/parse.go:239
	_go_fuzz_dep_.CoverTab[15880]++
						if 'A' <= b && func() bool {
//line /usr/local/go/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[15882]++
//line /usr/local/go/src/net/parse.go:240
		return b <= 'Z'
//line /usr/local/go/src/net/parse.go:240
		// _ = "end of CoverTab[15882]"
//line /usr/local/go/src/net/parse.go:240
	}() {
//line /usr/local/go/src/net/parse.go:240
		_go_fuzz_dep_.CoverTab[15883]++
							return b + ('a' - 'A')
//line /usr/local/go/src/net/parse.go:241
		// _ = "end of CoverTab[15883]"
	} else {
//line /usr/local/go/src/net/parse.go:242
		_go_fuzz_dep_.CoverTab[15884]++
//line /usr/local/go/src/net/parse.go:242
		// _ = "end of CoverTab[15884]"
//line /usr/local/go/src/net/parse.go:242
	}
//line /usr/local/go/src/net/parse.go:242
	// _ = "end of CoverTab[15880]"
//line /usr/local/go/src/net/parse.go:242
	_go_fuzz_dep_.CoverTab[15881]++
						return b
//line /usr/local/go/src/net/parse.go:243
	// _ = "end of CoverTab[15881]"
}

// trimSpace returns x without any leading or trailing ASCII whitespace.
func trimSpace(x string) string {
//line /usr/local/go/src/net/parse.go:247
	_go_fuzz_dep_.CoverTab[15885]++
						for len(x) > 0 && func() bool {
//line /usr/local/go/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[15888]++
//line /usr/local/go/src/net/parse.go:248
		return isSpace(x[0])
//line /usr/local/go/src/net/parse.go:248
		// _ = "end of CoverTab[15888]"
//line /usr/local/go/src/net/parse.go:248
	}() {
//line /usr/local/go/src/net/parse.go:248
		_go_fuzz_dep_.CoverTab[15889]++
							x = x[1:]
//line /usr/local/go/src/net/parse.go:249
		// _ = "end of CoverTab[15889]"
	}
//line /usr/local/go/src/net/parse.go:250
	// _ = "end of CoverTab[15885]"
//line /usr/local/go/src/net/parse.go:250
	_go_fuzz_dep_.CoverTab[15886]++
						for len(x) > 0 && func() bool {
//line /usr/local/go/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[15890]++
//line /usr/local/go/src/net/parse.go:251
		return isSpace(x[len(x)-1])
//line /usr/local/go/src/net/parse.go:251
		// _ = "end of CoverTab[15890]"
//line /usr/local/go/src/net/parse.go:251
	}() {
//line /usr/local/go/src/net/parse.go:251
		_go_fuzz_dep_.CoverTab[15891]++
							x = x[:len(x)-1]
//line /usr/local/go/src/net/parse.go:252
		// _ = "end of CoverTab[15891]"
	}
//line /usr/local/go/src/net/parse.go:253
	// _ = "end of CoverTab[15886]"
//line /usr/local/go/src/net/parse.go:253
	_go_fuzz_dep_.CoverTab[15887]++
						return x
//line /usr/local/go/src/net/parse.go:254
	// _ = "end of CoverTab[15887]"
}

// isSpace reports whether b is an ASCII space character.
func isSpace(b byte) bool {
//line /usr/local/go/src/net/parse.go:258
	_go_fuzz_dep_.CoverTab[15892]++
						return b == ' ' || func() bool {
//line /usr/local/go/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[15893]++
//line /usr/local/go/src/net/parse.go:259
		return b == '\t'
//line /usr/local/go/src/net/parse.go:259
		// _ = "end of CoverTab[15893]"
//line /usr/local/go/src/net/parse.go:259
	}() || func() bool {
//line /usr/local/go/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[15894]++
//line /usr/local/go/src/net/parse.go:259
		return b == '\n'
//line /usr/local/go/src/net/parse.go:259
		// _ = "end of CoverTab[15894]"
//line /usr/local/go/src/net/parse.go:259
	}() || func() bool {
//line /usr/local/go/src/net/parse.go:259
		_go_fuzz_dep_.CoverTab[15895]++
//line /usr/local/go/src/net/parse.go:259
		return b == '\r'
//line /usr/local/go/src/net/parse.go:259
		// _ = "end of CoverTab[15895]"
//line /usr/local/go/src/net/parse.go:259
	}()
//line /usr/local/go/src/net/parse.go:259
	// _ = "end of CoverTab[15892]"
}

// removeComment returns line, removing any '#' byte and any following
//line /usr/local/go/src/net/parse.go:262
// bytes.
//line /usr/local/go/src/net/parse.go:264
func removeComment(line string) string {
//line /usr/local/go/src/net/parse.go:264
	_go_fuzz_dep_.CoverTab[15896]++
						if i := bytealg.IndexByteString(line, '#'); i != -1 {
//line /usr/local/go/src/net/parse.go:265
		_go_fuzz_dep_.CoverTab[15898]++
							return line[:i]
//line /usr/local/go/src/net/parse.go:266
		// _ = "end of CoverTab[15898]"
	} else {
//line /usr/local/go/src/net/parse.go:267
		_go_fuzz_dep_.CoverTab[15899]++
//line /usr/local/go/src/net/parse.go:267
		// _ = "end of CoverTab[15899]"
//line /usr/local/go/src/net/parse.go:267
	}
//line /usr/local/go/src/net/parse.go:267
	// _ = "end of CoverTab[15896]"
//line /usr/local/go/src/net/parse.go:267
	_go_fuzz_dep_.CoverTab[15897]++
						return line
//line /usr/local/go/src/net/parse.go:268
	// _ = "end of CoverTab[15897]"
}

// foreachField runs fn on each non-empty run of non-space bytes in x.
//line /usr/local/go/src/net/parse.go:271
// It returns the first non-nil error returned by fn.
//line /usr/local/go/src/net/parse.go:273
func foreachField(x string, fn func(field string) error) error {
//line /usr/local/go/src/net/parse.go:273
	_go_fuzz_dep_.CoverTab[15900]++
						x = trimSpace(x)
						for len(x) > 0 {
//line /usr/local/go/src/net/parse.go:275
		_go_fuzz_dep_.CoverTab[15902]++
							sp := bytealg.IndexByteString(x, ' ')
							if sp == -1 {
//line /usr/local/go/src/net/parse.go:277
			_go_fuzz_dep_.CoverTab[15905]++
								return fn(x)
//line /usr/local/go/src/net/parse.go:278
			// _ = "end of CoverTab[15905]"
		} else {
//line /usr/local/go/src/net/parse.go:279
			_go_fuzz_dep_.CoverTab[15906]++
//line /usr/local/go/src/net/parse.go:279
			// _ = "end of CoverTab[15906]"
//line /usr/local/go/src/net/parse.go:279
		}
//line /usr/local/go/src/net/parse.go:279
		// _ = "end of CoverTab[15902]"
//line /usr/local/go/src/net/parse.go:279
		_go_fuzz_dep_.CoverTab[15903]++
							if field := trimSpace(x[:sp]); len(field) > 0 {
//line /usr/local/go/src/net/parse.go:280
			_go_fuzz_dep_.CoverTab[15907]++
								if err := fn(field); err != nil {
//line /usr/local/go/src/net/parse.go:281
				_go_fuzz_dep_.CoverTab[15908]++
									return err
//line /usr/local/go/src/net/parse.go:282
				// _ = "end of CoverTab[15908]"
			} else {
//line /usr/local/go/src/net/parse.go:283
				_go_fuzz_dep_.CoverTab[15909]++
//line /usr/local/go/src/net/parse.go:283
				// _ = "end of CoverTab[15909]"
//line /usr/local/go/src/net/parse.go:283
			}
//line /usr/local/go/src/net/parse.go:283
			// _ = "end of CoverTab[15907]"
		} else {
//line /usr/local/go/src/net/parse.go:284
			_go_fuzz_dep_.CoverTab[15910]++
//line /usr/local/go/src/net/parse.go:284
			// _ = "end of CoverTab[15910]"
//line /usr/local/go/src/net/parse.go:284
		}
//line /usr/local/go/src/net/parse.go:284
		// _ = "end of CoverTab[15903]"
//line /usr/local/go/src/net/parse.go:284
		_go_fuzz_dep_.CoverTab[15904]++
							x = trimSpace(x[sp+1:])
//line /usr/local/go/src/net/parse.go:285
		// _ = "end of CoverTab[15904]"
	}
//line /usr/local/go/src/net/parse.go:286
	// _ = "end of CoverTab[15900]"
//line /usr/local/go/src/net/parse.go:286
	_go_fuzz_dep_.CoverTab[15901]++
						return nil
//line /usr/local/go/src/net/parse.go:287
	// _ = "end of CoverTab[15901]"
}

// stringsHasSuffix is strings.HasSuffix. It reports whether s ends in
//line /usr/local/go/src/net/parse.go:290
// suffix.
//line /usr/local/go/src/net/parse.go:292
func stringsHasSuffix(s, suffix string) bool {
//line /usr/local/go/src/net/parse.go:292
	_go_fuzz_dep_.CoverTab[15911]++
						return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/net/parse.go:293
		_go_fuzz_dep_.CoverTab[15912]++
//line /usr/local/go/src/net/parse.go:293
		return s[len(s)-len(suffix):] == suffix
//line /usr/local/go/src/net/parse.go:293
		// _ = "end of CoverTab[15912]"
//line /usr/local/go/src/net/parse.go:293
	}()
//line /usr/local/go/src/net/parse.go:293
	// _ = "end of CoverTab[15911]"
}

// stringsHasSuffixFold reports whether s ends in suffix,
//line /usr/local/go/src/net/parse.go:296
// ASCII-case-insensitively.
//line /usr/local/go/src/net/parse.go:298
func stringsHasSuffixFold(s, suffix string) bool {
//line /usr/local/go/src/net/parse.go:298
	_go_fuzz_dep_.CoverTab[15913]++
						return len(s) >= len(suffix) && func() bool {
//line /usr/local/go/src/net/parse.go:299
		_go_fuzz_dep_.CoverTab[15914]++
//line /usr/local/go/src/net/parse.go:299
		return stringsEqualFold(s[len(s)-len(suffix):], suffix)
//line /usr/local/go/src/net/parse.go:299
		// _ = "end of CoverTab[15914]"
//line /usr/local/go/src/net/parse.go:299
	}()
//line /usr/local/go/src/net/parse.go:299
	// _ = "end of CoverTab[15913]"
}

// stringsHasPrefix is strings.HasPrefix. It reports whether s begins with prefix.
func stringsHasPrefix(s, prefix string) bool {
//line /usr/local/go/src/net/parse.go:303
	_go_fuzz_dep_.CoverTab[15915]++
						return len(s) >= len(prefix) && func() bool {
//line /usr/local/go/src/net/parse.go:304
		_go_fuzz_dep_.CoverTab[15916]++
//line /usr/local/go/src/net/parse.go:304
		return s[:len(prefix)] == prefix
//line /usr/local/go/src/net/parse.go:304
		// _ = "end of CoverTab[15916]"
//line /usr/local/go/src/net/parse.go:304
	}()
//line /usr/local/go/src/net/parse.go:304
	// _ = "end of CoverTab[15915]"
}

// stringsEqualFold is strings.EqualFold, ASCII only. It reports whether s and t
//line /usr/local/go/src/net/parse.go:307
// are equal, ASCII-case-insensitively.
//line /usr/local/go/src/net/parse.go:309
func stringsEqualFold(s, t string) bool {
//line /usr/local/go/src/net/parse.go:309
	_go_fuzz_dep_.CoverTab[15917]++
						if len(s) != len(t) {
//line /usr/local/go/src/net/parse.go:310
		_go_fuzz_dep_.CoverTab[15920]++
							return false
//line /usr/local/go/src/net/parse.go:311
		// _ = "end of CoverTab[15920]"
	} else {
//line /usr/local/go/src/net/parse.go:312
		_go_fuzz_dep_.CoverTab[15921]++
//line /usr/local/go/src/net/parse.go:312
		// _ = "end of CoverTab[15921]"
//line /usr/local/go/src/net/parse.go:312
	}
//line /usr/local/go/src/net/parse.go:312
	// _ = "end of CoverTab[15917]"
//line /usr/local/go/src/net/parse.go:312
	_go_fuzz_dep_.CoverTab[15918]++
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/parse.go:313
		_go_fuzz_dep_.CoverTab[15922]++
							if lowerASCII(s[i]) != lowerASCII(t[i]) {
//line /usr/local/go/src/net/parse.go:314
			_go_fuzz_dep_.CoverTab[15923]++
								return false
//line /usr/local/go/src/net/parse.go:315
			// _ = "end of CoverTab[15923]"
		} else {
//line /usr/local/go/src/net/parse.go:316
			_go_fuzz_dep_.CoverTab[15924]++
//line /usr/local/go/src/net/parse.go:316
			// _ = "end of CoverTab[15924]"
//line /usr/local/go/src/net/parse.go:316
		}
//line /usr/local/go/src/net/parse.go:316
		// _ = "end of CoverTab[15922]"
	}
//line /usr/local/go/src/net/parse.go:317
	// _ = "end of CoverTab[15918]"
//line /usr/local/go/src/net/parse.go:317
	_go_fuzz_dep_.CoverTab[15919]++
						return true
//line /usr/local/go/src/net/parse.go:318
	// _ = "end of CoverTab[15919]"
}

//line /usr/local/go/src/net/parse.go:319
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/parse.go:319
var _ = _go_fuzz_dep_.CoverTab
