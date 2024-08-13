// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
package norm

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:5
)

import "unicode/utf8"

type input struct {
	str	string
	bytes	[]byte
}

func inputBytes(str []byte) input {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:14
	_go_fuzz_dep_.CoverTab[70428]++
											return input{bytes: str}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:15
	// _ = "end of CoverTab[70428]"
}

func inputString(str string) input {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:18
	_go_fuzz_dep_.CoverTab[70429]++
											return input{str: str}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:19
	// _ = "end of CoverTab[70429]"
}

func (in *input) setBytes(str []byte) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:22
	_go_fuzz_dep_.CoverTab[70430]++
											in.str = ""
											in.bytes = str
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:24
	// _ = "end of CoverTab[70430]"
}

func (in *input) setString(str string) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:27
	_go_fuzz_dep_.CoverTab[70431]++
											in.str = str
											in.bytes = nil
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:29
	// _ = "end of CoverTab[70431]"
}

func (in *input) _byte(p int) byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:32
	_go_fuzz_dep_.CoverTab[70432]++
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:33
		_go_fuzz_dep_.CoverTab[70434]++
												return in.str[p]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:34
		// _ = "end of CoverTab[70434]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:35
		_go_fuzz_dep_.CoverTab[70435]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:35
		// _ = "end of CoverTab[70435]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:35
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:35
	// _ = "end of CoverTab[70432]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:35
	_go_fuzz_dep_.CoverTab[70433]++
											return in.bytes[p]
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:36
	// _ = "end of CoverTab[70433]"
}

func (in *input) skipASCII(p, max int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:39
	_go_fuzz_dep_.CoverTab[70436]++
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:40
		_go_fuzz_dep_.CoverTab[70438]++
												for ; p < max && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:41
			_go_fuzz_dep_.CoverTab[70439]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:41
			return in.str[p] < utf8.RuneSelf
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:41
			// _ = "end of CoverTab[70439]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:41
		}(); p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:41
			_go_fuzz_dep_.CoverTab[70440]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:41
			// _ = "end of CoverTab[70440]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:42
		// _ = "end of CoverTab[70438]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:43
		_go_fuzz_dep_.CoverTab[70441]++
												for ; p < max && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:44
			_go_fuzz_dep_.CoverTab[70442]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:44
			return in.bytes[p] < utf8.RuneSelf
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:44
			// _ = "end of CoverTab[70442]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:44
		}(); p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:44
			_go_fuzz_dep_.CoverTab[70443]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:44
			// _ = "end of CoverTab[70443]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:45
		// _ = "end of CoverTab[70441]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:46
	// _ = "end of CoverTab[70436]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:46
	_go_fuzz_dep_.CoverTab[70437]++
											return p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:47
	// _ = "end of CoverTab[70437]"
}

func (in *input) skipContinuationBytes(p int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:50
	_go_fuzz_dep_.CoverTab[70444]++
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:51
		_go_fuzz_dep_.CoverTab[70446]++
												for ; p < len(in.str) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:52
			_go_fuzz_dep_.CoverTab[70447]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:52
			return !utf8.RuneStart(in.str[p])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:52
			// _ = "end of CoverTab[70447]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:52
		}(); p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:52
			_go_fuzz_dep_.CoverTab[70448]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:52
			// _ = "end of CoverTab[70448]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:53
		// _ = "end of CoverTab[70446]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:54
		_go_fuzz_dep_.CoverTab[70449]++
												for ; p < len(in.bytes) && func() bool {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:55
			_go_fuzz_dep_.CoverTab[70450]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:55
			return !utf8.RuneStart(in.bytes[p])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:55
			// _ = "end of CoverTab[70450]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:55
		}(); p++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:55
			_go_fuzz_dep_.CoverTab[70451]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:55
			// _ = "end of CoverTab[70451]"
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:56
		// _ = "end of CoverTab[70449]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:57
	// _ = "end of CoverTab[70444]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:57
	_go_fuzz_dep_.CoverTab[70445]++
											return p
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:58
	// _ = "end of CoverTab[70445]"
}

func (in *input) appendSlice(buf []byte, b, e int) []byte {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:61
	_go_fuzz_dep_.CoverTab[70452]++
											if in.bytes != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:62
		_go_fuzz_dep_.CoverTab[70455]++
												return append(buf, in.bytes[b:e]...)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:63
		// _ = "end of CoverTab[70455]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:64
		_go_fuzz_dep_.CoverTab[70456]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:64
		// _ = "end of CoverTab[70456]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:64
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:64
	// _ = "end of CoverTab[70452]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:64
	_go_fuzz_dep_.CoverTab[70453]++
											for i := b; i < e; i++ {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:65
		_go_fuzz_dep_.CoverTab[70457]++
												buf = append(buf, in.str[i])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:66
		// _ = "end of CoverTab[70457]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:67
	// _ = "end of CoverTab[70453]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:67
	_go_fuzz_dep_.CoverTab[70454]++
											return buf
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:68
	// _ = "end of CoverTab[70454]"
}

func (in *input) copySlice(buf []byte, b, e int) int {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:71
	_go_fuzz_dep_.CoverTab[70458]++
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:72
		_go_fuzz_dep_.CoverTab[70460]++
												return copy(buf, in.str[b:e])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:73
		// _ = "end of CoverTab[70460]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:74
		_go_fuzz_dep_.CoverTab[70461]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:74
		// _ = "end of CoverTab[70461]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:74
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:74
	// _ = "end of CoverTab[70458]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:74
	_go_fuzz_dep_.CoverTab[70459]++
											return copy(buf, in.bytes[b:e])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:75
	// _ = "end of CoverTab[70459]"
}

func (in *input) charinfoNFC(p int) (uint16, int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:78
	_go_fuzz_dep_.CoverTab[70462]++
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:79
		_go_fuzz_dep_.CoverTab[70464]++
												return nfcData.lookupString(in.str[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:80
		// _ = "end of CoverTab[70464]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:81
		_go_fuzz_dep_.CoverTab[70465]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:81
		// _ = "end of CoverTab[70465]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:81
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:81
	// _ = "end of CoverTab[70462]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:81
	_go_fuzz_dep_.CoverTab[70463]++
											return nfcData.lookup(in.bytes[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:82
	// _ = "end of CoverTab[70463]"
}

func (in *input) charinfoNFKC(p int) (uint16, int) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:85
	_go_fuzz_dep_.CoverTab[70466]++
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:86
		_go_fuzz_dep_.CoverTab[70468]++
												return nfkcData.lookupString(in.str[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:87
		// _ = "end of CoverTab[70468]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:88
		_go_fuzz_dep_.CoverTab[70469]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:88
		// _ = "end of CoverTab[70469]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:88
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:88
	// _ = "end of CoverTab[70466]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:88
	_go_fuzz_dep_.CoverTab[70467]++
											return nfkcData.lookup(in.bytes[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:89
	// _ = "end of CoverTab[70467]"
}

func (in *input) hangul(p int) (r rune) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:92
	_go_fuzz_dep_.CoverTab[70470]++
											var size int
											if in.bytes == nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:94
		_go_fuzz_dep_.CoverTab[70473]++
												if !isHangulString(in.str[p:]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:95
			_go_fuzz_dep_.CoverTab[70475]++
													return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:96
			// _ = "end of CoverTab[70475]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:97
			_go_fuzz_dep_.CoverTab[70476]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:97
			// _ = "end of CoverTab[70476]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:97
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:97
		// _ = "end of CoverTab[70473]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:97
		_go_fuzz_dep_.CoverTab[70474]++
												r, size = utf8.DecodeRuneInString(in.str[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:98
		// _ = "end of CoverTab[70474]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:99
		_go_fuzz_dep_.CoverTab[70477]++
												if !isHangul(in.bytes[p:]) {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:100
			_go_fuzz_dep_.CoverTab[70479]++
													return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:101
			// _ = "end of CoverTab[70479]"
		} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:102
			_go_fuzz_dep_.CoverTab[70480]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:102
			// _ = "end of CoverTab[70480]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:102
		}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:102
		// _ = "end of CoverTab[70477]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:102
		_go_fuzz_dep_.CoverTab[70478]++
												r, size = utf8.DecodeRune(in.bytes[p:])
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:103
		// _ = "end of CoverTab[70478]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:104
	// _ = "end of CoverTab[70470]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:104
	_go_fuzz_dep_.CoverTab[70471]++
											if size != hangulUTF8Size {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:105
		_go_fuzz_dep_.CoverTab[70481]++
												return 0
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:106
		// _ = "end of CoverTab[70481]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:107
		_go_fuzz_dep_.CoverTab[70482]++
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:107
		// _ = "end of CoverTab[70482]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:107
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:107
	// _ = "end of CoverTab[70471]"
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:107
	_go_fuzz_dep_.CoverTab[70472]++
											return r
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:108
	// _ = "end of CoverTab[70472]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/text@v0.9.0/unicode/norm/input.go:109
var _ = _go_fuzz_dep_.CoverTab
