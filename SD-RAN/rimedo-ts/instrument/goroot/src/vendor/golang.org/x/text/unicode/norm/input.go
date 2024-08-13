// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
package norm

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
import (
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:5
)

import "unicode/utf8"

type input struct {
	str	string
	bytes	[]byte
}

func inputBytes(str []byte) input {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:14
	_go_fuzz_dep_.CoverTab[33147]++
										return input{bytes: str}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:15
	// _ = "end of CoverTab[33147]"
}

func inputString(str string) input {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:18
	_go_fuzz_dep_.CoverTab[33148]++
										return input{str: str}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:19
	// _ = "end of CoverTab[33148]"
}

func (in *input) setBytes(str []byte) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:22
	_go_fuzz_dep_.CoverTab[33149]++
										in.str = ""
										in.bytes = str
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:24
	// _ = "end of CoverTab[33149]"
}

func (in *input) setString(str string) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:27
	_go_fuzz_dep_.CoverTab[33150]++
										in.str = str
										in.bytes = nil
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:29
	// _ = "end of CoverTab[33150]"
}

func (in *input) _byte(p int) byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:32
	_go_fuzz_dep_.CoverTab[33151]++
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:33
		_go_fuzz_dep_.CoverTab[33153]++
											return in.str[p]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:34
		// _ = "end of CoverTab[33153]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:35
		_go_fuzz_dep_.CoverTab[33154]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:35
		// _ = "end of CoverTab[33154]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:35
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:35
	// _ = "end of CoverTab[33151]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:35
	_go_fuzz_dep_.CoverTab[33152]++
										return in.bytes[p]
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:36
	// _ = "end of CoverTab[33152]"
}

func (in *input) skipASCII(p, max int) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:39
	_go_fuzz_dep_.CoverTab[33155]++
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:40
		_go_fuzz_dep_.CoverTab[33157]++
											for ; p < max && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:41
			_go_fuzz_dep_.CoverTab[33158]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:41
			return in.str[p] < utf8.RuneSelf
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:41
			// _ = "end of CoverTab[33158]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:41
		}(); p++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:41
			_go_fuzz_dep_.CoverTab[33159]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:41
			// _ = "end of CoverTab[33159]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:42
		// _ = "end of CoverTab[33157]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:43
		_go_fuzz_dep_.CoverTab[33160]++
											for ; p < max && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:44
			_go_fuzz_dep_.CoverTab[33161]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:44
			return in.bytes[p] < utf8.RuneSelf
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:44
			// _ = "end of CoverTab[33161]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:44
		}(); p++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:44
			_go_fuzz_dep_.CoverTab[33162]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:44
			// _ = "end of CoverTab[33162]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:45
		// _ = "end of CoverTab[33160]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:46
	// _ = "end of CoverTab[33155]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:46
	_go_fuzz_dep_.CoverTab[33156]++
										return p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:47
	// _ = "end of CoverTab[33156]"
}

func (in *input) skipContinuationBytes(p int) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:50
	_go_fuzz_dep_.CoverTab[33163]++
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:51
		_go_fuzz_dep_.CoverTab[33165]++
											for ; p < len(in.str) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:52
			_go_fuzz_dep_.CoverTab[33166]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:52
			return !utf8.RuneStart(in.str[p])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:52
			// _ = "end of CoverTab[33166]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:52
		}(); p++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:52
			_go_fuzz_dep_.CoverTab[33167]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:52
			// _ = "end of CoverTab[33167]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:53
		// _ = "end of CoverTab[33165]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:54
		_go_fuzz_dep_.CoverTab[33168]++
											for ; p < len(in.bytes) && func() bool {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:55
			_go_fuzz_dep_.CoverTab[33169]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:55
			return !utf8.RuneStart(in.bytes[p])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:55
			// _ = "end of CoverTab[33169]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:55
		}(); p++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:55
			_go_fuzz_dep_.CoverTab[33170]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:55
			// _ = "end of CoverTab[33170]"
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:56
		// _ = "end of CoverTab[33168]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:57
	// _ = "end of CoverTab[33163]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:57
	_go_fuzz_dep_.CoverTab[33164]++
										return p
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:58
	// _ = "end of CoverTab[33164]"
}

func (in *input) appendSlice(buf []byte, b, e int) []byte {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:61
	_go_fuzz_dep_.CoverTab[33171]++
										if in.bytes != nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:62
		_go_fuzz_dep_.CoverTab[33174]++
											return append(buf, in.bytes[b:e]...)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:63
		// _ = "end of CoverTab[33174]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:64
		_go_fuzz_dep_.CoverTab[33175]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:64
		// _ = "end of CoverTab[33175]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:64
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:64
	// _ = "end of CoverTab[33171]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:64
	_go_fuzz_dep_.CoverTab[33172]++
										for i := b; i < e; i++ {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:65
		_go_fuzz_dep_.CoverTab[33176]++
											buf = append(buf, in.str[i])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:66
		// _ = "end of CoverTab[33176]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:67
	// _ = "end of CoverTab[33172]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:67
	_go_fuzz_dep_.CoverTab[33173]++
										return buf
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:68
	// _ = "end of CoverTab[33173]"
}

func (in *input) copySlice(buf []byte, b, e int) int {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:71
	_go_fuzz_dep_.CoverTab[33177]++
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:72
		_go_fuzz_dep_.CoverTab[33179]++
											return copy(buf, in.str[b:e])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:73
		// _ = "end of CoverTab[33179]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:74
		_go_fuzz_dep_.CoverTab[33180]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:74
		// _ = "end of CoverTab[33180]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:74
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:74
	// _ = "end of CoverTab[33177]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:74
	_go_fuzz_dep_.CoverTab[33178]++
										return copy(buf, in.bytes[b:e])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:75
	// _ = "end of CoverTab[33178]"
}

func (in *input) charinfoNFC(p int) (uint16, int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:78
	_go_fuzz_dep_.CoverTab[33181]++
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:79
		_go_fuzz_dep_.CoverTab[33183]++
											return nfcData.lookupString(in.str[p:])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:80
		// _ = "end of CoverTab[33183]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:81
		_go_fuzz_dep_.CoverTab[33184]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:81
		// _ = "end of CoverTab[33184]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:81
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:81
	// _ = "end of CoverTab[33181]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:81
	_go_fuzz_dep_.CoverTab[33182]++
										return nfcData.lookup(in.bytes[p:])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:82
	// _ = "end of CoverTab[33182]"
}

func (in *input) charinfoNFKC(p int) (uint16, int) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:85
	_go_fuzz_dep_.CoverTab[33185]++
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:86
		_go_fuzz_dep_.CoverTab[33187]++
											return nfkcData.lookupString(in.str[p:])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:87
		// _ = "end of CoverTab[33187]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:88
		_go_fuzz_dep_.CoverTab[33188]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:88
		// _ = "end of CoverTab[33188]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:88
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:88
	// _ = "end of CoverTab[33185]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:88
	_go_fuzz_dep_.CoverTab[33186]++
										return nfkcData.lookup(in.bytes[p:])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:89
	// _ = "end of CoverTab[33186]"
}

func (in *input) hangul(p int) (r rune) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:92
	_go_fuzz_dep_.CoverTab[33189]++
										var size int
										if in.bytes == nil {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:94
		_go_fuzz_dep_.CoverTab[33192]++
											if !isHangulString(in.str[p:]) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:95
			_go_fuzz_dep_.CoverTab[33194]++
												return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:96
			// _ = "end of CoverTab[33194]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:97
			_go_fuzz_dep_.CoverTab[33195]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:97
			// _ = "end of CoverTab[33195]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:97
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:97
		// _ = "end of CoverTab[33192]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:97
		_go_fuzz_dep_.CoverTab[33193]++
											r, size = utf8.DecodeRuneInString(in.str[p:])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:98
		// _ = "end of CoverTab[33193]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:99
		_go_fuzz_dep_.CoverTab[33196]++
											if !isHangul(in.bytes[p:]) {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:100
			_go_fuzz_dep_.CoverTab[33198]++
												return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:101
			// _ = "end of CoverTab[33198]"
		} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:102
			_go_fuzz_dep_.CoverTab[33199]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:102
			// _ = "end of CoverTab[33199]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:102
		}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:102
		// _ = "end of CoverTab[33196]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:102
		_go_fuzz_dep_.CoverTab[33197]++
											r, size = utf8.DecodeRune(in.bytes[p:])
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:103
		// _ = "end of CoverTab[33197]"
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:104
	// _ = "end of CoverTab[33189]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:104
	_go_fuzz_dep_.CoverTab[33190]++
										if size != hangulUTF8Size {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:105
		_go_fuzz_dep_.CoverTab[33200]++
											return 0
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:106
		// _ = "end of CoverTab[33200]"
	} else {
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:107
		_go_fuzz_dep_.CoverTab[33201]++
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:107
		// _ = "end of CoverTab[33201]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:107
	}
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:107
	// _ = "end of CoverTab[33190]"
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:107
	_go_fuzz_dep_.CoverTab[33191]++
										return r
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:108
	// _ = "end of CoverTab[33191]"
}

//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:109
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/vendor/golang.org/x/text/unicode/norm/input.go:109
var _ = _go_fuzz_dep_.CoverTab
