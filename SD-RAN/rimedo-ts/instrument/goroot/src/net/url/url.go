// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/url/url.go:5
// Package url parses URLs and implements query escaping.
package url

//line /usr/local/go/src/net/url/url.go:6
import (
//line /usr/local/go/src/net/url/url.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/url/url.go:6
)
//line /usr/local/go/src/net/url/url.go:6
import (
//line /usr/local/go/src/net/url/url.go:6
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/url/url.go:6
)

//line /usr/local/go/src/net/url/url.go:13
import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"
)

// Error reports an error and the operation and URL that caused it.
type Error struct {
	Op	string
	URL	string
	Err	error
}

func (e *Error) Unwrap() error {
//line /usr/local/go/src/net/url/url.go:29
	_go_fuzz_dep_.CoverTab[17216]++
//line /usr/local/go/src/net/url/url.go:29
	return e.Err
//line /usr/local/go/src/net/url/url.go:29
	// _ = "end of CoverTab[17216]"
//line /usr/local/go/src/net/url/url.go:29
}
func (e *Error) Error() string {
//line /usr/local/go/src/net/url/url.go:30
	_go_fuzz_dep_.CoverTab[17217]++
//line /usr/local/go/src/net/url/url.go:30
	return fmt.Sprintf("%s %q: %s", e.Op, e.URL, e.Err)
//line /usr/local/go/src/net/url/url.go:30
	// _ = "end of CoverTab[17217]"
//line /usr/local/go/src/net/url/url.go:30
}

func (e *Error) Timeout() bool {
//line /usr/local/go/src/net/url/url.go:32
	_go_fuzz_dep_.CoverTab[17218]++
						t, ok := e.Err.(interface {
		Timeout() bool
	})
	return ok && func() bool {
//line /usr/local/go/src/net/url/url.go:36
		_go_fuzz_dep_.CoverTab[17219]++
//line /usr/local/go/src/net/url/url.go:36
		return t.Timeout()
//line /usr/local/go/src/net/url/url.go:36
		// _ = "end of CoverTab[17219]"
//line /usr/local/go/src/net/url/url.go:36
	}()
//line /usr/local/go/src/net/url/url.go:36
	// _ = "end of CoverTab[17218]"
}

func (e *Error) Temporary() bool {
//line /usr/local/go/src/net/url/url.go:39
	_go_fuzz_dep_.CoverTab[17220]++
						t, ok := e.Err.(interface {
		Temporary() bool
	})
	return ok && func() bool {
//line /usr/local/go/src/net/url/url.go:43
		_go_fuzz_dep_.CoverTab[17221]++
//line /usr/local/go/src/net/url/url.go:43
		return t.Temporary()
//line /usr/local/go/src/net/url/url.go:43
		// _ = "end of CoverTab[17221]"
//line /usr/local/go/src/net/url/url.go:43
	}()
//line /usr/local/go/src/net/url/url.go:43
	// _ = "end of CoverTab[17220]"
}

const upperhex = "0123456789ABCDEF"

func ishex(c byte) bool {
//line /usr/local/go/src/net/url/url.go:48
	_go_fuzz_dep_.CoverTab[17222]++
						switch {
	case '0' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:50
		_go_fuzz_dep_.CoverTab[17228]++
//line /usr/local/go/src/net/url/url.go:50
		return c <= '9'
//line /usr/local/go/src/net/url/url.go:50
		// _ = "end of CoverTab[17228]"
//line /usr/local/go/src/net/url/url.go:50
	}():
//line /usr/local/go/src/net/url/url.go:50
		_go_fuzz_dep_.CoverTab[17224]++
							return true
//line /usr/local/go/src/net/url/url.go:51
		// _ = "end of CoverTab[17224]"
	case 'a' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:52
		_go_fuzz_dep_.CoverTab[17229]++
//line /usr/local/go/src/net/url/url.go:52
		return c <= 'f'
//line /usr/local/go/src/net/url/url.go:52
		// _ = "end of CoverTab[17229]"
//line /usr/local/go/src/net/url/url.go:52
	}():
//line /usr/local/go/src/net/url/url.go:52
		_go_fuzz_dep_.CoverTab[17225]++
							return true
//line /usr/local/go/src/net/url/url.go:53
		// _ = "end of CoverTab[17225]"
	case 'A' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:54
		_go_fuzz_dep_.CoverTab[17230]++
//line /usr/local/go/src/net/url/url.go:54
		return c <= 'F'
//line /usr/local/go/src/net/url/url.go:54
		// _ = "end of CoverTab[17230]"
//line /usr/local/go/src/net/url/url.go:54
	}():
//line /usr/local/go/src/net/url/url.go:54
		_go_fuzz_dep_.CoverTab[17226]++
							return true
//line /usr/local/go/src/net/url/url.go:55
		// _ = "end of CoverTab[17226]"
//line /usr/local/go/src/net/url/url.go:55
	default:
//line /usr/local/go/src/net/url/url.go:55
		_go_fuzz_dep_.CoverTab[17227]++
//line /usr/local/go/src/net/url/url.go:55
		// _ = "end of CoverTab[17227]"
	}
//line /usr/local/go/src/net/url/url.go:56
	// _ = "end of CoverTab[17222]"
//line /usr/local/go/src/net/url/url.go:56
	_go_fuzz_dep_.CoverTab[17223]++
						return false
//line /usr/local/go/src/net/url/url.go:57
	// _ = "end of CoverTab[17223]"
}

func unhex(c byte) byte {
//line /usr/local/go/src/net/url/url.go:60
	_go_fuzz_dep_.CoverTab[17231]++
						switch {
	case '0' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:62
		_go_fuzz_dep_.CoverTab[17237]++
//line /usr/local/go/src/net/url/url.go:62
		return c <= '9'
//line /usr/local/go/src/net/url/url.go:62
		// _ = "end of CoverTab[17237]"
//line /usr/local/go/src/net/url/url.go:62
	}():
//line /usr/local/go/src/net/url/url.go:62
		_go_fuzz_dep_.CoverTab[17233]++
							return c - '0'
//line /usr/local/go/src/net/url/url.go:63
		// _ = "end of CoverTab[17233]"
	case 'a' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:64
		_go_fuzz_dep_.CoverTab[17238]++
//line /usr/local/go/src/net/url/url.go:64
		return c <= 'f'
//line /usr/local/go/src/net/url/url.go:64
		// _ = "end of CoverTab[17238]"
//line /usr/local/go/src/net/url/url.go:64
	}():
//line /usr/local/go/src/net/url/url.go:64
		_go_fuzz_dep_.CoverTab[17234]++
							return c - 'a' + 10
//line /usr/local/go/src/net/url/url.go:65
		// _ = "end of CoverTab[17234]"
	case 'A' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:66
		_go_fuzz_dep_.CoverTab[17239]++
//line /usr/local/go/src/net/url/url.go:66
		return c <= 'F'
//line /usr/local/go/src/net/url/url.go:66
		// _ = "end of CoverTab[17239]"
//line /usr/local/go/src/net/url/url.go:66
	}():
//line /usr/local/go/src/net/url/url.go:66
		_go_fuzz_dep_.CoverTab[17235]++
							return c - 'A' + 10
//line /usr/local/go/src/net/url/url.go:67
		// _ = "end of CoverTab[17235]"
//line /usr/local/go/src/net/url/url.go:67
	default:
//line /usr/local/go/src/net/url/url.go:67
		_go_fuzz_dep_.CoverTab[17236]++
//line /usr/local/go/src/net/url/url.go:67
		// _ = "end of CoverTab[17236]"
	}
//line /usr/local/go/src/net/url/url.go:68
	// _ = "end of CoverTab[17231]"
//line /usr/local/go/src/net/url/url.go:68
	_go_fuzz_dep_.CoverTab[17232]++
						return 0
//line /usr/local/go/src/net/url/url.go:69
	// _ = "end of CoverTab[17232]"
}

type encoding int

const (
	encodePath	encoding	= 1 + iota
	encodePathSegment
	encodeHost
	encodeZone
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

type EscapeError string

func (e EscapeError) Error() string {
//line /usr/local/go/src/net/url/url.go:86
	_go_fuzz_dep_.CoverTab[17240]++
						return "invalid URL escape " + strconv.Quote(string(e))
//line /usr/local/go/src/net/url/url.go:87
	// _ = "end of CoverTab[17240]"
}

type InvalidHostError string

func (e InvalidHostError) Error() string {
//line /usr/local/go/src/net/url/url.go:92
	_go_fuzz_dep_.CoverTab[17241]++
						return "invalid character " + strconv.Quote(string(e)) + " in host name"
//line /usr/local/go/src/net/url/url.go:93
	// _ = "end of CoverTab[17241]"
}

// Return true if the specified character should be escaped when
//line /usr/local/go/src/net/url/url.go:96
// appearing in a URL string, according to RFC 3986.
//line /usr/local/go/src/net/url/url.go:96
//
//line /usr/local/go/src/net/url/url.go:96
// Please be informed that for now shouldEscape does not check all
//line /usr/local/go/src/net/url/url.go:96
// reserved characters correctly. See golang.org/issue/5684.
//line /usr/local/go/src/net/url/url.go:101
func shouldEscape(c byte, mode encoding) bool {
//line /usr/local/go/src/net/url/url.go:101
	_go_fuzz_dep_.CoverTab[17242]++

						if 'a' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:103
		_go_fuzz_dep_.CoverTab[17247]++
//line /usr/local/go/src/net/url/url.go:103
		return c <= 'z'
//line /usr/local/go/src/net/url/url.go:103
		// _ = "end of CoverTab[17247]"
//line /usr/local/go/src/net/url/url.go:103
	}() || func() bool {
//line /usr/local/go/src/net/url/url.go:103
		_go_fuzz_dep_.CoverTab[17248]++
//line /usr/local/go/src/net/url/url.go:103
		return 'A' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:103
			_go_fuzz_dep_.CoverTab[17249]++
//line /usr/local/go/src/net/url/url.go:103
			return c <= 'Z'
//line /usr/local/go/src/net/url/url.go:103
			// _ = "end of CoverTab[17249]"
//line /usr/local/go/src/net/url/url.go:103
		}()
//line /usr/local/go/src/net/url/url.go:103
		// _ = "end of CoverTab[17248]"
//line /usr/local/go/src/net/url/url.go:103
	}() || func() bool {
//line /usr/local/go/src/net/url/url.go:103
		_go_fuzz_dep_.CoverTab[17250]++
//line /usr/local/go/src/net/url/url.go:103
		return '0' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:103
			_go_fuzz_dep_.CoverTab[17251]++
//line /usr/local/go/src/net/url/url.go:103
			return c <= '9'
//line /usr/local/go/src/net/url/url.go:103
			// _ = "end of CoverTab[17251]"
//line /usr/local/go/src/net/url/url.go:103
		}()
//line /usr/local/go/src/net/url/url.go:103
		// _ = "end of CoverTab[17250]"
//line /usr/local/go/src/net/url/url.go:103
	}() {
//line /usr/local/go/src/net/url/url.go:103
		_go_fuzz_dep_.CoverTab[17252]++
							return false
//line /usr/local/go/src/net/url/url.go:104
		// _ = "end of CoverTab[17252]"
	} else {
//line /usr/local/go/src/net/url/url.go:105
		_go_fuzz_dep_.CoverTab[17253]++
//line /usr/local/go/src/net/url/url.go:105
		// _ = "end of CoverTab[17253]"
//line /usr/local/go/src/net/url/url.go:105
	}
//line /usr/local/go/src/net/url/url.go:105
	// _ = "end of CoverTab[17242]"
//line /usr/local/go/src/net/url/url.go:105
	_go_fuzz_dep_.CoverTab[17243]++

						if mode == encodeHost || func() bool {
//line /usr/local/go/src/net/url/url.go:107
		_go_fuzz_dep_.CoverTab[17254]++
//line /usr/local/go/src/net/url/url.go:107
		return mode == encodeZone
//line /usr/local/go/src/net/url/url.go:107
		// _ = "end of CoverTab[17254]"
//line /usr/local/go/src/net/url/url.go:107
	}() {
//line /usr/local/go/src/net/url/url.go:107
		_go_fuzz_dep_.CoverTab[17255]++

//line /usr/local/go/src/net/url/url.go:117
		switch c {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '[', ']', '<', '>', '"':
//line /usr/local/go/src/net/url/url.go:118
			_go_fuzz_dep_.CoverTab[17256]++
								return false
//line /usr/local/go/src/net/url/url.go:119
			// _ = "end of CoverTab[17256]"
//line /usr/local/go/src/net/url/url.go:119
		default:
//line /usr/local/go/src/net/url/url.go:119
			_go_fuzz_dep_.CoverTab[17257]++
//line /usr/local/go/src/net/url/url.go:119
			// _ = "end of CoverTab[17257]"
		}
//line /usr/local/go/src/net/url/url.go:120
		// _ = "end of CoverTab[17255]"
	} else {
//line /usr/local/go/src/net/url/url.go:121
		_go_fuzz_dep_.CoverTab[17258]++
//line /usr/local/go/src/net/url/url.go:121
		// _ = "end of CoverTab[17258]"
//line /usr/local/go/src/net/url/url.go:121
	}
//line /usr/local/go/src/net/url/url.go:121
	// _ = "end of CoverTab[17243]"
//line /usr/local/go/src/net/url/url.go:121
	_go_fuzz_dep_.CoverTab[17244]++

						switch c {
	case '-', '_', '.', '~':
//line /usr/local/go/src/net/url/url.go:124
		_go_fuzz_dep_.CoverTab[17259]++
							return false
//line /usr/local/go/src/net/url/url.go:125
		// _ = "end of CoverTab[17259]"

	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@':
//line /usr/local/go/src/net/url/url.go:127
		_go_fuzz_dep_.CoverTab[17260]++

//line /usr/local/go/src/net/url/url.go:130
		switch mode {
		case encodePath:
//line /usr/local/go/src/net/url/url.go:131
			_go_fuzz_dep_.CoverTab[17262]++

//line /usr/local/go/src/net/url/url.go:136
			return c == '?'
//line /usr/local/go/src/net/url/url.go:136
			// _ = "end of CoverTab[17262]"

		case encodePathSegment:
//line /usr/local/go/src/net/url/url.go:138
			_go_fuzz_dep_.CoverTab[17263]++

//line /usr/local/go/src/net/url/url.go:141
			return c == '/' || func() bool {
//line /usr/local/go/src/net/url/url.go:141
				_go_fuzz_dep_.CoverTab[17268]++
//line /usr/local/go/src/net/url/url.go:141
				return c == ';'
//line /usr/local/go/src/net/url/url.go:141
				// _ = "end of CoverTab[17268]"
//line /usr/local/go/src/net/url/url.go:141
			}() || func() bool {
//line /usr/local/go/src/net/url/url.go:141
				_go_fuzz_dep_.CoverTab[17269]++
//line /usr/local/go/src/net/url/url.go:141
				return c == ','
//line /usr/local/go/src/net/url/url.go:141
				// _ = "end of CoverTab[17269]"
//line /usr/local/go/src/net/url/url.go:141
			}() || func() bool {
//line /usr/local/go/src/net/url/url.go:141
				_go_fuzz_dep_.CoverTab[17270]++
//line /usr/local/go/src/net/url/url.go:141
				return c == '?'
//line /usr/local/go/src/net/url/url.go:141
				// _ = "end of CoverTab[17270]"
//line /usr/local/go/src/net/url/url.go:141
			}()
//line /usr/local/go/src/net/url/url.go:141
			// _ = "end of CoverTab[17263]"

		case encodeUserPassword:
//line /usr/local/go/src/net/url/url.go:143
			_go_fuzz_dep_.CoverTab[17264]++

//line /usr/local/go/src/net/url/url.go:148
			return c == '@' || func() bool {
//line /usr/local/go/src/net/url/url.go:148
				_go_fuzz_dep_.CoverTab[17271]++
//line /usr/local/go/src/net/url/url.go:148
				return c == '/'
//line /usr/local/go/src/net/url/url.go:148
				// _ = "end of CoverTab[17271]"
//line /usr/local/go/src/net/url/url.go:148
			}() || func() bool {
//line /usr/local/go/src/net/url/url.go:148
				_go_fuzz_dep_.CoverTab[17272]++
//line /usr/local/go/src/net/url/url.go:148
				return c == '?'
//line /usr/local/go/src/net/url/url.go:148
				// _ = "end of CoverTab[17272]"
//line /usr/local/go/src/net/url/url.go:148
			}() || func() bool {
//line /usr/local/go/src/net/url/url.go:148
				_go_fuzz_dep_.CoverTab[17273]++
//line /usr/local/go/src/net/url/url.go:148
				return c == ':'
//line /usr/local/go/src/net/url/url.go:148
				// _ = "end of CoverTab[17273]"
//line /usr/local/go/src/net/url/url.go:148
			}()
//line /usr/local/go/src/net/url/url.go:148
			// _ = "end of CoverTab[17264]"

		case encodeQueryComponent:
//line /usr/local/go/src/net/url/url.go:150
			_go_fuzz_dep_.CoverTab[17265]++

								return true
//line /usr/local/go/src/net/url/url.go:152
			// _ = "end of CoverTab[17265]"

		case encodeFragment:
//line /usr/local/go/src/net/url/url.go:154
			_go_fuzz_dep_.CoverTab[17266]++

//line /usr/local/go/src/net/url/url.go:157
			return false
//line /usr/local/go/src/net/url/url.go:157
			// _ = "end of CoverTab[17266]"
//line /usr/local/go/src/net/url/url.go:157
		default:
//line /usr/local/go/src/net/url/url.go:157
			_go_fuzz_dep_.CoverTab[17267]++
//line /usr/local/go/src/net/url/url.go:157
			// _ = "end of CoverTab[17267]"
		}
//line /usr/local/go/src/net/url/url.go:158
		// _ = "end of CoverTab[17260]"
//line /usr/local/go/src/net/url/url.go:158
	default:
//line /usr/local/go/src/net/url/url.go:158
		_go_fuzz_dep_.CoverTab[17261]++
//line /usr/local/go/src/net/url/url.go:158
		// _ = "end of CoverTab[17261]"
	}
//line /usr/local/go/src/net/url/url.go:159
	// _ = "end of CoverTab[17244]"
//line /usr/local/go/src/net/url/url.go:159
	_go_fuzz_dep_.CoverTab[17245]++

						if mode == encodeFragment {
//line /usr/local/go/src/net/url/url.go:161
		_go_fuzz_dep_.CoverTab[17274]++

//line /usr/local/go/src/net/url/url.go:168
		switch c {
		case '!', '(', ')', '*':
//line /usr/local/go/src/net/url/url.go:169
			_go_fuzz_dep_.CoverTab[17275]++
								return false
//line /usr/local/go/src/net/url/url.go:170
			// _ = "end of CoverTab[17275]"
//line /usr/local/go/src/net/url/url.go:170
		default:
//line /usr/local/go/src/net/url/url.go:170
			_go_fuzz_dep_.CoverTab[17276]++
//line /usr/local/go/src/net/url/url.go:170
			// _ = "end of CoverTab[17276]"
		}
//line /usr/local/go/src/net/url/url.go:171
		// _ = "end of CoverTab[17274]"
	} else {
//line /usr/local/go/src/net/url/url.go:172
		_go_fuzz_dep_.CoverTab[17277]++
//line /usr/local/go/src/net/url/url.go:172
		// _ = "end of CoverTab[17277]"
//line /usr/local/go/src/net/url/url.go:172
	}
//line /usr/local/go/src/net/url/url.go:172
	// _ = "end of CoverTab[17245]"
//line /usr/local/go/src/net/url/url.go:172
	_go_fuzz_dep_.CoverTab[17246]++

//line /usr/local/go/src/net/url/url.go:175
	return true
//line /usr/local/go/src/net/url/url.go:175
	// _ = "end of CoverTab[17246]"
}

// QueryUnescape does the inverse transformation of QueryEscape,
//line /usr/local/go/src/net/url/url.go:178
// converting each 3-byte encoded substring of the form "%AB" into the
//line /usr/local/go/src/net/url/url.go:178
// hex-decoded byte 0xAB.
//line /usr/local/go/src/net/url/url.go:178
// It returns an error if any % is not followed by two hexadecimal
//line /usr/local/go/src/net/url/url.go:178
// digits.
//line /usr/local/go/src/net/url/url.go:183
func QueryUnescape(s string) (string, error) {
//line /usr/local/go/src/net/url/url.go:183
	_go_fuzz_dep_.CoverTab[17278]++
						return unescape(s, encodeQueryComponent)
//line /usr/local/go/src/net/url/url.go:184
	// _ = "end of CoverTab[17278]"
}

// PathUnescape does the inverse transformation of PathEscape,
//line /usr/local/go/src/net/url/url.go:187
// converting each 3-byte encoded substring of the form "%AB" into the
//line /usr/local/go/src/net/url/url.go:187
// hex-decoded byte 0xAB. It returns an error if any % is not followed
//line /usr/local/go/src/net/url/url.go:187
// by two hexadecimal digits.
//line /usr/local/go/src/net/url/url.go:187
//
//line /usr/local/go/src/net/url/url.go:187
// PathUnescape is identical to QueryUnescape except that it does not
//line /usr/local/go/src/net/url/url.go:187
// unescape '+' to ' ' (space).
//line /usr/local/go/src/net/url/url.go:194
func PathUnescape(s string) (string, error) {
//line /usr/local/go/src/net/url/url.go:194
	_go_fuzz_dep_.CoverTab[17279]++
						return unescape(s, encodePathSegment)
//line /usr/local/go/src/net/url/url.go:195
	// _ = "end of CoverTab[17279]"
}

// unescape unescapes a string; the mode specifies
//line /usr/local/go/src/net/url/url.go:198
// which section of the URL string is being unescaped.
//line /usr/local/go/src/net/url/url.go:200
func unescape(s string, mode encoding) (string, error) {
//line /usr/local/go/src/net/url/url.go:200
	_go_fuzz_dep_.CoverTab[17280]++

						n := 0
						hasPlus := false
						for i := 0; i < len(s); {
//line /usr/local/go/src/net/url/url.go:204
		_go_fuzz_dep_.CoverTab[17284]++
							switch s[i] {
		case '%':
//line /usr/local/go/src/net/url/url.go:206
			_go_fuzz_dep_.CoverTab[17285]++
								n++
								if i+2 >= len(s) || func() bool {
//line /usr/local/go/src/net/url/url.go:208
				_go_fuzz_dep_.CoverTab[17292]++
//line /usr/local/go/src/net/url/url.go:208
				return !ishex(s[i+1])
//line /usr/local/go/src/net/url/url.go:208
				// _ = "end of CoverTab[17292]"
//line /usr/local/go/src/net/url/url.go:208
			}() || func() bool {
//line /usr/local/go/src/net/url/url.go:208
				_go_fuzz_dep_.CoverTab[17293]++
//line /usr/local/go/src/net/url/url.go:208
				return !ishex(s[i+2])
//line /usr/local/go/src/net/url/url.go:208
				// _ = "end of CoverTab[17293]"
//line /usr/local/go/src/net/url/url.go:208
			}() {
//line /usr/local/go/src/net/url/url.go:208
				_go_fuzz_dep_.CoverTab[17294]++
									s = s[i:]
									if len(s) > 3 {
//line /usr/local/go/src/net/url/url.go:210
					_go_fuzz_dep_.CoverTab[17296]++
										s = s[:3]
//line /usr/local/go/src/net/url/url.go:211
					// _ = "end of CoverTab[17296]"
				} else {
//line /usr/local/go/src/net/url/url.go:212
					_go_fuzz_dep_.CoverTab[17297]++
//line /usr/local/go/src/net/url/url.go:212
					// _ = "end of CoverTab[17297]"
//line /usr/local/go/src/net/url/url.go:212
				}
//line /usr/local/go/src/net/url/url.go:212
				// _ = "end of CoverTab[17294]"
//line /usr/local/go/src/net/url/url.go:212
				_go_fuzz_dep_.CoverTab[17295]++
									return "", EscapeError(s)
//line /usr/local/go/src/net/url/url.go:213
				// _ = "end of CoverTab[17295]"
			} else {
//line /usr/local/go/src/net/url/url.go:214
				_go_fuzz_dep_.CoverTab[17298]++
//line /usr/local/go/src/net/url/url.go:214
				// _ = "end of CoverTab[17298]"
//line /usr/local/go/src/net/url/url.go:214
			}
//line /usr/local/go/src/net/url/url.go:214
			// _ = "end of CoverTab[17285]"
//line /usr/local/go/src/net/url/url.go:214
			_go_fuzz_dep_.CoverTab[17286]++

//line /usr/local/go/src/net/url/url.go:221
			if mode == encodeHost && func() bool {
//line /usr/local/go/src/net/url/url.go:221
				_go_fuzz_dep_.CoverTab[17299]++
//line /usr/local/go/src/net/url/url.go:221
				return unhex(s[i+1]) < 8
//line /usr/local/go/src/net/url/url.go:221
				// _ = "end of CoverTab[17299]"
//line /usr/local/go/src/net/url/url.go:221
			}() && func() bool {
//line /usr/local/go/src/net/url/url.go:221
				_go_fuzz_dep_.CoverTab[17300]++
//line /usr/local/go/src/net/url/url.go:221
				return s[i:i+3] != "%25"
//line /usr/local/go/src/net/url/url.go:221
				// _ = "end of CoverTab[17300]"
//line /usr/local/go/src/net/url/url.go:221
			}() {
//line /usr/local/go/src/net/url/url.go:221
				_go_fuzz_dep_.CoverTab[17301]++
									return "", EscapeError(s[i : i+3])
//line /usr/local/go/src/net/url/url.go:222
				// _ = "end of CoverTab[17301]"
			} else {
//line /usr/local/go/src/net/url/url.go:223
				_go_fuzz_dep_.CoverTab[17302]++
//line /usr/local/go/src/net/url/url.go:223
				// _ = "end of CoverTab[17302]"
//line /usr/local/go/src/net/url/url.go:223
			}
//line /usr/local/go/src/net/url/url.go:223
			// _ = "end of CoverTab[17286]"
//line /usr/local/go/src/net/url/url.go:223
			_go_fuzz_dep_.CoverTab[17287]++
								if mode == encodeZone {
//line /usr/local/go/src/net/url/url.go:224
				_go_fuzz_dep_.CoverTab[17303]++

//line /usr/local/go/src/net/url/url.go:232
				v := unhex(s[i+1])<<4 | unhex(s[i+2])
				if s[i:i+3] != "%25" && func() bool {
//line /usr/local/go/src/net/url/url.go:233
					_go_fuzz_dep_.CoverTab[17304]++
//line /usr/local/go/src/net/url/url.go:233
					return v != ' '
//line /usr/local/go/src/net/url/url.go:233
					// _ = "end of CoverTab[17304]"
//line /usr/local/go/src/net/url/url.go:233
				}() && func() bool {
//line /usr/local/go/src/net/url/url.go:233
					_go_fuzz_dep_.CoverTab[17305]++
//line /usr/local/go/src/net/url/url.go:233
					return shouldEscape(v, encodeHost)
//line /usr/local/go/src/net/url/url.go:233
					// _ = "end of CoverTab[17305]"
//line /usr/local/go/src/net/url/url.go:233
				}() {
//line /usr/local/go/src/net/url/url.go:233
					_go_fuzz_dep_.CoverTab[17306]++
										return "", EscapeError(s[i : i+3])
//line /usr/local/go/src/net/url/url.go:234
					// _ = "end of CoverTab[17306]"
				} else {
//line /usr/local/go/src/net/url/url.go:235
					_go_fuzz_dep_.CoverTab[17307]++
//line /usr/local/go/src/net/url/url.go:235
					// _ = "end of CoverTab[17307]"
//line /usr/local/go/src/net/url/url.go:235
				}
//line /usr/local/go/src/net/url/url.go:235
				// _ = "end of CoverTab[17303]"
			} else {
//line /usr/local/go/src/net/url/url.go:236
				_go_fuzz_dep_.CoverTab[17308]++
//line /usr/local/go/src/net/url/url.go:236
				// _ = "end of CoverTab[17308]"
//line /usr/local/go/src/net/url/url.go:236
			}
//line /usr/local/go/src/net/url/url.go:236
			// _ = "end of CoverTab[17287]"
//line /usr/local/go/src/net/url/url.go:236
			_go_fuzz_dep_.CoverTab[17288]++
								i += 3
//line /usr/local/go/src/net/url/url.go:237
			// _ = "end of CoverTab[17288]"
		case '+':
//line /usr/local/go/src/net/url/url.go:238
			_go_fuzz_dep_.CoverTab[17289]++
								hasPlus = mode == encodeQueryComponent
								i++
//line /usr/local/go/src/net/url/url.go:240
			// _ = "end of CoverTab[17289]"
		default:
//line /usr/local/go/src/net/url/url.go:241
			_go_fuzz_dep_.CoverTab[17290]++
								if (mode == encodeHost || func() bool {
//line /usr/local/go/src/net/url/url.go:242
				_go_fuzz_dep_.CoverTab[17309]++
//line /usr/local/go/src/net/url/url.go:242
				return mode == encodeZone
//line /usr/local/go/src/net/url/url.go:242
				// _ = "end of CoverTab[17309]"
//line /usr/local/go/src/net/url/url.go:242
			}()) && func() bool {
//line /usr/local/go/src/net/url/url.go:242
				_go_fuzz_dep_.CoverTab[17310]++
//line /usr/local/go/src/net/url/url.go:242
				return s[i] < 0x80
//line /usr/local/go/src/net/url/url.go:242
				// _ = "end of CoverTab[17310]"
//line /usr/local/go/src/net/url/url.go:242
			}() && func() bool {
//line /usr/local/go/src/net/url/url.go:242
				_go_fuzz_dep_.CoverTab[17311]++
//line /usr/local/go/src/net/url/url.go:242
				return shouldEscape(s[i], mode)
//line /usr/local/go/src/net/url/url.go:242
				// _ = "end of CoverTab[17311]"
//line /usr/local/go/src/net/url/url.go:242
			}() {
//line /usr/local/go/src/net/url/url.go:242
				_go_fuzz_dep_.CoverTab[17312]++
									return "", InvalidHostError(s[i : i+1])
//line /usr/local/go/src/net/url/url.go:243
				// _ = "end of CoverTab[17312]"
			} else {
//line /usr/local/go/src/net/url/url.go:244
				_go_fuzz_dep_.CoverTab[17313]++
//line /usr/local/go/src/net/url/url.go:244
				// _ = "end of CoverTab[17313]"
//line /usr/local/go/src/net/url/url.go:244
			}
//line /usr/local/go/src/net/url/url.go:244
			// _ = "end of CoverTab[17290]"
//line /usr/local/go/src/net/url/url.go:244
			_go_fuzz_dep_.CoverTab[17291]++
								i++
//line /usr/local/go/src/net/url/url.go:245
			// _ = "end of CoverTab[17291]"
		}
//line /usr/local/go/src/net/url/url.go:246
		// _ = "end of CoverTab[17284]"
	}
//line /usr/local/go/src/net/url/url.go:247
	// _ = "end of CoverTab[17280]"
//line /usr/local/go/src/net/url/url.go:247
	_go_fuzz_dep_.CoverTab[17281]++

						if n == 0 && func() bool {
//line /usr/local/go/src/net/url/url.go:249
		_go_fuzz_dep_.CoverTab[17314]++
//line /usr/local/go/src/net/url/url.go:249
		return !hasPlus
//line /usr/local/go/src/net/url/url.go:249
		// _ = "end of CoverTab[17314]"
//line /usr/local/go/src/net/url/url.go:249
	}() {
//line /usr/local/go/src/net/url/url.go:249
		_go_fuzz_dep_.CoverTab[17315]++
							return s, nil
//line /usr/local/go/src/net/url/url.go:250
		// _ = "end of CoverTab[17315]"
	} else {
//line /usr/local/go/src/net/url/url.go:251
		_go_fuzz_dep_.CoverTab[17316]++
//line /usr/local/go/src/net/url/url.go:251
		// _ = "end of CoverTab[17316]"
//line /usr/local/go/src/net/url/url.go:251
	}
//line /usr/local/go/src/net/url/url.go:251
	// _ = "end of CoverTab[17281]"
//line /usr/local/go/src/net/url/url.go:251
	_go_fuzz_dep_.CoverTab[17282]++

						var t strings.Builder
						t.Grow(len(s) - 2*n)
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/url/url.go:255
		_go_fuzz_dep_.CoverTab[17317]++
							switch s[i] {
		case '%':
//line /usr/local/go/src/net/url/url.go:257
			_go_fuzz_dep_.CoverTab[17318]++
								t.WriteByte(unhex(s[i+1])<<4 | unhex(s[i+2]))
								i += 2
//line /usr/local/go/src/net/url/url.go:259
			// _ = "end of CoverTab[17318]"
		case '+':
//line /usr/local/go/src/net/url/url.go:260
			_go_fuzz_dep_.CoverTab[17319]++
								if mode == encodeQueryComponent {
//line /usr/local/go/src/net/url/url.go:261
				_go_fuzz_dep_.CoverTab[17321]++
									t.WriteByte(' ')
//line /usr/local/go/src/net/url/url.go:262
				// _ = "end of CoverTab[17321]"
			} else {
//line /usr/local/go/src/net/url/url.go:263
				_go_fuzz_dep_.CoverTab[17322]++
									t.WriteByte('+')
//line /usr/local/go/src/net/url/url.go:264
				// _ = "end of CoverTab[17322]"
			}
//line /usr/local/go/src/net/url/url.go:265
			// _ = "end of CoverTab[17319]"
		default:
//line /usr/local/go/src/net/url/url.go:266
			_go_fuzz_dep_.CoverTab[17320]++
								t.WriteByte(s[i])
//line /usr/local/go/src/net/url/url.go:267
			// _ = "end of CoverTab[17320]"
		}
//line /usr/local/go/src/net/url/url.go:268
		// _ = "end of CoverTab[17317]"
	}
//line /usr/local/go/src/net/url/url.go:269
	// _ = "end of CoverTab[17282]"
//line /usr/local/go/src/net/url/url.go:269
	_go_fuzz_dep_.CoverTab[17283]++
						return t.String(), nil
//line /usr/local/go/src/net/url/url.go:270
	// _ = "end of CoverTab[17283]"
}

// QueryEscape escapes the string so it can be safely placed
//line /usr/local/go/src/net/url/url.go:273
// inside a URL query.
//line /usr/local/go/src/net/url/url.go:275
func QueryEscape(s string) string {
//line /usr/local/go/src/net/url/url.go:275
	_go_fuzz_dep_.CoverTab[17323]++
						return escape(s, encodeQueryComponent)
//line /usr/local/go/src/net/url/url.go:276
	// _ = "end of CoverTab[17323]"
}

// PathEscape escapes the string so it can be safely placed inside a URL path segment,
//line /usr/local/go/src/net/url/url.go:279
// replacing special characters (including /) with %XX sequences as needed.
//line /usr/local/go/src/net/url/url.go:281
func PathEscape(s string) string {
//line /usr/local/go/src/net/url/url.go:281
	_go_fuzz_dep_.CoverTab[17324]++
						return escape(s, encodePathSegment)
//line /usr/local/go/src/net/url/url.go:282
	// _ = "end of CoverTab[17324]"
}

func escape(s string, mode encoding) string {
//line /usr/local/go/src/net/url/url.go:285
	_go_fuzz_dep_.CoverTab[17325]++
						spaceCount, hexCount := 0, 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/url/url.go:287
		_go_fuzz_dep_.CoverTab[17331]++
							c := s[i]
							if shouldEscape(c, mode) {
//line /usr/local/go/src/net/url/url.go:289
			_go_fuzz_dep_.CoverTab[17332]++
								if c == ' ' && func() bool {
//line /usr/local/go/src/net/url/url.go:290
				_go_fuzz_dep_.CoverTab[17333]++
//line /usr/local/go/src/net/url/url.go:290
				return mode == encodeQueryComponent
//line /usr/local/go/src/net/url/url.go:290
				// _ = "end of CoverTab[17333]"
//line /usr/local/go/src/net/url/url.go:290
			}() {
//line /usr/local/go/src/net/url/url.go:290
				_go_fuzz_dep_.CoverTab[17334]++
									spaceCount++
//line /usr/local/go/src/net/url/url.go:291
				// _ = "end of CoverTab[17334]"
			} else {
//line /usr/local/go/src/net/url/url.go:292
				_go_fuzz_dep_.CoverTab[17335]++
									hexCount++
//line /usr/local/go/src/net/url/url.go:293
				// _ = "end of CoverTab[17335]"
			}
//line /usr/local/go/src/net/url/url.go:294
			// _ = "end of CoverTab[17332]"
		} else {
//line /usr/local/go/src/net/url/url.go:295
			_go_fuzz_dep_.CoverTab[17336]++
//line /usr/local/go/src/net/url/url.go:295
			// _ = "end of CoverTab[17336]"
//line /usr/local/go/src/net/url/url.go:295
		}
//line /usr/local/go/src/net/url/url.go:295
		// _ = "end of CoverTab[17331]"
	}
//line /usr/local/go/src/net/url/url.go:296
	// _ = "end of CoverTab[17325]"
//line /usr/local/go/src/net/url/url.go:296
	_go_fuzz_dep_.CoverTab[17326]++

						if spaceCount == 0 && func() bool {
//line /usr/local/go/src/net/url/url.go:298
		_go_fuzz_dep_.CoverTab[17337]++
//line /usr/local/go/src/net/url/url.go:298
		return hexCount == 0
//line /usr/local/go/src/net/url/url.go:298
		// _ = "end of CoverTab[17337]"
//line /usr/local/go/src/net/url/url.go:298
	}() {
//line /usr/local/go/src/net/url/url.go:298
		_go_fuzz_dep_.CoverTab[17338]++
							return s
//line /usr/local/go/src/net/url/url.go:299
		// _ = "end of CoverTab[17338]"
	} else {
//line /usr/local/go/src/net/url/url.go:300
		_go_fuzz_dep_.CoverTab[17339]++
//line /usr/local/go/src/net/url/url.go:300
		// _ = "end of CoverTab[17339]"
//line /usr/local/go/src/net/url/url.go:300
	}
//line /usr/local/go/src/net/url/url.go:300
	// _ = "end of CoverTab[17326]"
//line /usr/local/go/src/net/url/url.go:300
	_go_fuzz_dep_.CoverTab[17327]++

						var buf [64]byte
						var t []byte

						required := len(s) + 2*hexCount
						if required <= len(buf) {
//line /usr/local/go/src/net/url/url.go:306
		_go_fuzz_dep_.CoverTab[17340]++
							t = buf[:required]
//line /usr/local/go/src/net/url/url.go:307
		// _ = "end of CoverTab[17340]"
	} else {
//line /usr/local/go/src/net/url/url.go:308
		_go_fuzz_dep_.CoverTab[17341]++
							t = make([]byte, required)
//line /usr/local/go/src/net/url/url.go:309
		// _ = "end of CoverTab[17341]"
	}
//line /usr/local/go/src/net/url/url.go:310
	// _ = "end of CoverTab[17327]"
//line /usr/local/go/src/net/url/url.go:310
	_go_fuzz_dep_.CoverTab[17328]++

						if hexCount == 0 {
//line /usr/local/go/src/net/url/url.go:312
		_go_fuzz_dep_.CoverTab[17342]++
							copy(t, s)
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/url/url.go:314
			_go_fuzz_dep_.CoverTab[17344]++
								if s[i] == ' ' {
//line /usr/local/go/src/net/url/url.go:315
				_go_fuzz_dep_.CoverTab[17345]++
									t[i] = '+'
//line /usr/local/go/src/net/url/url.go:316
				// _ = "end of CoverTab[17345]"
			} else {
//line /usr/local/go/src/net/url/url.go:317
				_go_fuzz_dep_.CoverTab[17346]++
//line /usr/local/go/src/net/url/url.go:317
				// _ = "end of CoverTab[17346]"
//line /usr/local/go/src/net/url/url.go:317
			}
//line /usr/local/go/src/net/url/url.go:317
			// _ = "end of CoverTab[17344]"
		}
//line /usr/local/go/src/net/url/url.go:318
		// _ = "end of CoverTab[17342]"
//line /usr/local/go/src/net/url/url.go:318
		_go_fuzz_dep_.CoverTab[17343]++
							return string(t)
//line /usr/local/go/src/net/url/url.go:319
		// _ = "end of CoverTab[17343]"
	} else {
//line /usr/local/go/src/net/url/url.go:320
		_go_fuzz_dep_.CoverTab[17347]++
//line /usr/local/go/src/net/url/url.go:320
		// _ = "end of CoverTab[17347]"
//line /usr/local/go/src/net/url/url.go:320
	}
//line /usr/local/go/src/net/url/url.go:320
	// _ = "end of CoverTab[17328]"
//line /usr/local/go/src/net/url/url.go:320
	_go_fuzz_dep_.CoverTab[17329]++

						j := 0
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/url/url.go:323
		_go_fuzz_dep_.CoverTab[17348]++
							switch c := s[i]; {
		case c == ' ' && func() bool {
//line /usr/local/go/src/net/url/url.go:325
			_go_fuzz_dep_.CoverTab[17352]++
//line /usr/local/go/src/net/url/url.go:325
			return mode == encodeQueryComponent
//line /usr/local/go/src/net/url/url.go:325
			// _ = "end of CoverTab[17352]"
//line /usr/local/go/src/net/url/url.go:325
		}():
//line /usr/local/go/src/net/url/url.go:325
			_go_fuzz_dep_.CoverTab[17349]++
								t[j] = '+'
								j++
//line /usr/local/go/src/net/url/url.go:327
			// _ = "end of CoverTab[17349]"
		case shouldEscape(c, mode):
//line /usr/local/go/src/net/url/url.go:328
			_go_fuzz_dep_.CoverTab[17350]++
								t[j] = '%'
								t[j+1] = upperhex[c>>4]
								t[j+2] = upperhex[c&15]
								j += 3
//line /usr/local/go/src/net/url/url.go:332
			// _ = "end of CoverTab[17350]"
		default:
//line /usr/local/go/src/net/url/url.go:333
			_go_fuzz_dep_.CoverTab[17351]++
								t[j] = s[i]
								j++
//line /usr/local/go/src/net/url/url.go:335
			// _ = "end of CoverTab[17351]"
		}
//line /usr/local/go/src/net/url/url.go:336
		// _ = "end of CoverTab[17348]"
	}
//line /usr/local/go/src/net/url/url.go:337
	// _ = "end of CoverTab[17329]"
//line /usr/local/go/src/net/url/url.go:337
	_go_fuzz_dep_.CoverTab[17330]++
						return string(t)
//line /usr/local/go/src/net/url/url.go:338
	// _ = "end of CoverTab[17330]"
}

// A URL represents a parsed URL (technically, a URI reference).
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
// The general form represented is:
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
// URLs that do not start with a slash after the scheme are interpreted as:
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
//	scheme:opaque[?query][#fragment]
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
//line /usr/local/go/src/net/url/url.go:341
// A consequence is that it is impossible to tell which slashes in the Path were
//line /usr/local/go/src/net/url/url.go:341
// slashes in the raw URL and which were %2f. This distinction is rarely important,
//line /usr/local/go/src/net/url/url.go:341
// but when it is, the code should use the EscapedPath method, which preserves
//line /usr/local/go/src/net/url/url.go:341
// the original encoding of Path.
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
// The RawPath field is an optional field which is only set when the default
//line /usr/local/go/src/net/url/url.go:341
// encoding of Path is different from the escaped path. See the EscapedPath method
//line /usr/local/go/src/net/url/url.go:341
// for more details.
//line /usr/local/go/src/net/url/url.go:341
//
//line /usr/local/go/src/net/url/url.go:341
// URL's String method uses the EscapedPath method to obtain the path.
//line /usr/local/go/src/net/url/url.go:362
type URL struct {
	Scheme		string
	Opaque		string		// encoded opaque data
	User		*Userinfo	// username and password information
	Host		string		// host or host:port
	Path		string		// path (relative paths may omit leading slash)
	RawPath		string		// encoded path hint (see EscapedPath method)
	OmitHost	bool		// do not emit empty host (authority)
	ForceQuery	bool		// append a query ('?') even if RawQuery is empty
	RawQuery	string		// encoded query values, without '?'
	Fragment	string		// fragment for references, without '#'
	RawFragment	string		// encoded fragment hint (see EscapedFragment method)
}

// User returns a Userinfo containing the provided username
//line /usr/local/go/src/net/url/url.go:376
// and no password set.
//line /usr/local/go/src/net/url/url.go:378
func User(username string) *Userinfo {
//line /usr/local/go/src/net/url/url.go:378
	_go_fuzz_dep_.CoverTab[17353]++
						return &Userinfo{username, "", false}
//line /usr/local/go/src/net/url/url.go:379
	// _ = "end of CoverTab[17353]"
}

// UserPassword returns a Userinfo containing the provided username
//line /usr/local/go/src/net/url/url.go:382
// and password.
//line /usr/local/go/src/net/url/url.go:382
//
//line /usr/local/go/src/net/url/url.go:382
// This functionality should only be used with legacy web sites.
//line /usr/local/go/src/net/url/url.go:382
// RFC 2396 warns that interpreting Userinfo this way
//line /usr/local/go/src/net/url/url.go:382
// “is NOT RECOMMENDED, because the passing of authentication
//line /usr/local/go/src/net/url/url.go:382
// information in clear text (such as URI) has proven to be a
//line /usr/local/go/src/net/url/url.go:382
// security risk in almost every case where it has been used.”
//line /usr/local/go/src/net/url/url.go:390
func UserPassword(username, password string) *Userinfo {
//line /usr/local/go/src/net/url/url.go:390
	_go_fuzz_dep_.CoverTab[17354]++
						return &Userinfo{username, password, true}
//line /usr/local/go/src/net/url/url.go:391
	// _ = "end of CoverTab[17354]"
}

// The Userinfo type is an immutable encapsulation of username and
//line /usr/local/go/src/net/url/url.go:394
// password details for a URL. An existing Userinfo value is guaranteed
//line /usr/local/go/src/net/url/url.go:394
// to have a username set (potentially empty, as allowed by RFC 2396),
//line /usr/local/go/src/net/url/url.go:394
// and optionally a password.
//line /usr/local/go/src/net/url/url.go:398
type Userinfo struct {
	username	string
	password	string
	passwordSet	bool
}

// Username returns the username.
func (u *Userinfo) Username() string {
//line /usr/local/go/src/net/url/url.go:405
	_go_fuzz_dep_.CoverTab[17355]++
						if u == nil {
//line /usr/local/go/src/net/url/url.go:406
		_go_fuzz_dep_.CoverTab[17357]++
							return ""
//line /usr/local/go/src/net/url/url.go:407
		// _ = "end of CoverTab[17357]"
	} else {
//line /usr/local/go/src/net/url/url.go:408
		_go_fuzz_dep_.CoverTab[17358]++
//line /usr/local/go/src/net/url/url.go:408
		// _ = "end of CoverTab[17358]"
//line /usr/local/go/src/net/url/url.go:408
	}
//line /usr/local/go/src/net/url/url.go:408
	// _ = "end of CoverTab[17355]"
//line /usr/local/go/src/net/url/url.go:408
	_go_fuzz_dep_.CoverTab[17356]++
						return u.username
//line /usr/local/go/src/net/url/url.go:409
	// _ = "end of CoverTab[17356]"
}

// Password returns the password in case it is set, and whether it is set.
func (u *Userinfo) Password() (string, bool) {
//line /usr/local/go/src/net/url/url.go:413
	_go_fuzz_dep_.CoverTab[17359]++
						if u == nil {
//line /usr/local/go/src/net/url/url.go:414
		_go_fuzz_dep_.CoverTab[17361]++
							return "", false
//line /usr/local/go/src/net/url/url.go:415
		// _ = "end of CoverTab[17361]"
	} else {
//line /usr/local/go/src/net/url/url.go:416
		_go_fuzz_dep_.CoverTab[17362]++
//line /usr/local/go/src/net/url/url.go:416
		// _ = "end of CoverTab[17362]"
//line /usr/local/go/src/net/url/url.go:416
	}
//line /usr/local/go/src/net/url/url.go:416
	// _ = "end of CoverTab[17359]"
//line /usr/local/go/src/net/url/url.go:416
	_go_fuzz_dep_.CoverTab[17360]++
						return u.password, u.passwordSet
//line /usr/local/go/src/net/url/url.go:417
	// _ = "end of CoverTab[17360]"
}

// String returns the encoded userinfo information in the standard form
//line /usr/local/go/src/net/url/url.go:420
// of "username[:password]".
//line /usr/local/go/src/net/url/url.go:422
func (u *Userinfo) String() string {
//line /usr/local/go/src/net/url/url.go:422
	_go_fuzz_dep_.CoverTab[17363]++
						if u == nil {
//line /usr/local/go/src/net/url/url.go:423
		_go_fuzz_dep_.CoverTab[17366]++
							return ""
//line /usr/local/go/src/net/url/url.go:424
		// _ = "end of CoverTab[17366]"
	} else {
//line /usr/local/go/src/net/url/url.go:425
		_go_fuzz_dep_.CoverTab[17367]++
//line /usr/local/go/src/net/url/url.go:425
		// _ = "end of CoverTab[17367]"
//line /usr/local/go/src/net/url/url.go:425
	}
//line /usr/local/go/src/net/url/url.go:425
	// _ = "end of CoverTab[17363]"
//line /usr/local/go/src/net/url/url.go:425
	_go_fuzz_dep_.CoverTab[17364]++
						s := escape(u.username, encodeUserPassword)
						if u.passwordSet {
//line /usr/local/go/src/net/url/url.go:427
		_go_fuzz_dep_.CoverTab[17368]++
							s += ":" + escape(u.password, encodeUserPassword)
//line /usr/local/go/src/net/url/url.go:428
		// _ = "end of CoverTab[17368]"
	} else {
//line /usr/local/go/src/net/url/url.go:429
		_go_fuzz_dep_.CoverTab[17369]++
//line /usr/local/go/src/net/url/url.go:429
		// _ = "end of CoverTab[17369]"
//line /usr/local/go/src/net/url/url.go:429
	}
//line /usr/local/go/src/net/url/url.go:429
	// _ = "end of CoverTab[17364]"
//line /usr/local/go/src/net/url/url.go:429
	_go_fuzz_dep_.CoverTab[17365]++
						return s
//line /usr/local/go/src/net/url/url.go:430
	// _ = "end of CoverTab[17365]"
}

// Maybe rawURL is of the form scheme:path.
//line /usr/local/go/src/net/url/url.go:433
// (Scheme must be [a-zA-Z][a-zA-Z0-9+.-]*)
//line /usr/local/go/src/net/url/url.go:433
// If so, return scheme, path; else return "", rawURL.
//line /usr/local/go/src/net/url/url.go:436
func getScheme(rawURL string) (scheme, path string, err error) {
//line /usr/local/go/src/net/url/url.go:436
	_go_fuzz_dep_.CoverTab[17370]++
						for i := 0; i < len(rawURL); i++ {
//line /usr/local/go/src/net/url/url.go:437
		_go_fuzz_dep_.CoverTab[17372]++
							c := rawURL[i]
							switch {
		case 'a' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:440
			_go_fuzz_dep_.CoverTab[17378]++
//line /usr/local/go/src/net/url/url.go:440
			return c <= 'z'
//line /usr/local/go/src/net/url/url.go:440
			// _ = "end of CoverTab[17378]"
//line /usr/local/go/src/net/url/url.go:440
		}() || func() bool {
//line /usr/local/go/src/net/url/url.go:440
			_go_fuzz_dep_.CoverTab[17379]++
//line /usr/local/go/src/net/url/url.go:440
			return 'A' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:440
				_go_fuzz_dep_.CoverTab[17380]++
//line /usr/local/go/src/net/url/url.go:440
				return c <= 'Z'
//line /usr/local/go/src/net/url/url.go:440
				// _ = "end of CoverTab[17380]"
//line /usr/local/go/src/net/url/url.go:440
			}()
//line /usr/local/go/src/net/url/url.go:440
			// _ = "end of CoverTab[17379]"
//line /usr/local/go/src/net/url/url.go:440
		}():
//line /usr/local/go/src/net/url/url.go:440
			_go_fuzz_dep_.CoverTab[17373]++
//line /usr/local/go/src/net/url/url.go:440
			// _ = "end of CoverTab[17373]"

		case '0' <= c && func() bool {
//line /usr/local/go/src/net/url/url.go:442
			_go_fuzz_dep_.CoverTab[17381]++
//line /usr/local/go/src/net/url/url.go:442
			return c <= '9'
//line /usr/local/go/src/net/url/url.go:442
			// _ = "end of CoverTab[17381]"
//line /usr/local/go/src/net/url/url.go:442
		}() || func() bool {
//line /usr/local/go/src/net/url/url.go:442
			_go_fuzz_dep_.CoverTab[17382]++
//line /usr/local/go/src/net/url/url.go:442
			return c == '+'
//line /usr/local/go/src/net/url/url.go:442
			// _ = "end of CoverTab[17382]"
//line /usr/local/go/src/net/url/url.go:442
		}() || func() bool {
//line /usr/local/go/src/net/url/url.go:442
			_go_fuzz_dep_.CoverTab[17383]++
//line /usr/local/go/src/net/url/url.go:442
			return c == '-'
//line /usr/local/go/src/net/url/url.go:442
			// _ = "end of CoverTab[17383]"
//line /usr/local/go/src/net/url/url.go:442
		}() || func() bool {
//line /usr/local/go/src/net/url/url.go:442
			_go_fuzz_dep_.CoverTab[17384]++
//line /usr/local/go/src/net/url/url.go:442
			return c == '.'
//line /usr/local/go/src/net/url/url.go:442
			// _ = "end of CoverTab[17384]"
//line /usr/local/go/src/net/url/url.go:442
		}():
//line /usr/local/go/src/net/url/url.go:442
			_go_fuzz_dep_.CoverTab[17374]++
								if i == 0 {
//line /usr/local/go/src/net/url/url.go:443
				_go_fuzz_dep_.CoverTab[17385]++
									return "", rawURL, nil
//line /usr/local/go/src/net/url/url.go:444
				// _ = "end of CoverTab[17385]"
			} else {
//line /usr/local/go/src/net/url/url.go:445
				_go_fuzz_dep_.CoverTab[17386]++
//line /usr/local/go/src/net/url/url.go:445
				// _ = "end of CoverTab[17386]"
//line /usr/local/go/src/net/url/url.go:445
			}
//line /usr/local/go/src/net/url/url.go:445
			// _ = "end of CoverTab[17374]"
		case c == ':':
//line /usr/local/go/src/net/url/url.go:446
			_go_fuzz_dep_.CoverTab[17375]++
								if i == 0 {
//line /usr/local/go/src/net/url/url.go:447
				_go_fuzz_dep_.CoverTab[17387]++
									return "", "", errors.New("missing protocol scheme")
//line /usr/local/go/src/net/url/url.go:448
				// _ = "end of CoverTab[17387]"
			} else {
//line /usr/local/go/src/net/url/url.go:449
				_go_fuzz_dep_.CoverTab[17388]++
//line /usr/local/go/src/net/url/url.go:449
				// _ = "end of CoverTab[17388]"
//line /usr/local/go/src/net/url/url.go:449
			}
//line /usr/local/go/src/net/url/url.go:449
			// _ = "end of CoverTab[17375]"
//line /usr/local/go/src/net/url/url.go:449
			_go_fuzz_dep_.CoverTab[17376]++
								return rawURL[:i], rawURL[i+1:], nil
//line /usr/local/go/src/net/url/url.go:450
			// _ = "end of CoverTab[17376]"
		default:
//line /usr/local/go/src/net/url/url.go:451
			_go_fuzz_dep_.CoverTab[17377]++

//line /usr/local/go/src/net/url/url.go:454
			return "", rawURL, nil
//line /usr/local/go/src/net/url/url.go:454
			// _ = "end of CoverTab[17377]"
		}
//line /usr/local/go/src/net/url/url.go:455
		// _ = "end of CoverTab[17372]"
	}
//line /usr/local/go/src/net/url/url.go:456
	// _ = "end of CoverTab[17370]"
//line /usr/local/go/src/net/url/url.go:456
	_go_fuzz_dep_.CoverTab[17371]++
						return "", rawURL, nil
//line /usr/local/go/src/net/url/url.go:457
	// _ = "end of CoverTab[17371]"
}

// Parse parses a raw url into a URL structure.
//line /usr/local/go/src/net/url/url.go:460
//
//line /usr/local/go/src/net/url/url.go:460
// The url may be relative (a path, without a host) or absolute
//line /usr/local/go/src/net/url/url.go:460
// (starting with a scheme). Trying to parse a hostname and path
//line /usr/local/go/src/net/url/url.go:460
// without a scheme is invalid but may not necessarily return an
//line /usr/local/go/src/net/url/url.go:460
// error, due to parsing ambiguities.
//line /usr/local/go/src/net/url/url.go:466
func Parse(rawURL string) (*URL, error) {
//line /usr/local/go/src/net/url/url.go:466
	_go_fuzz_dep_.CoverTab[17389]++

						u, frag, _ := strings.Cut(rawURL, "#")
						url, err := parse(u, false)
						if err != nil {
//line /usr/local/go/src/net/url/url.go:470
		_go_fuzz_dep_.CoverTab[17393]++
							return nil, &Error{"parse", u, err}
//line /usr/local/go/src/net/url/url.go:471
		// _ = "end of CoverTab[17393]"
	} else {
//line /usr/local/go/src/net/url/url.go:472
		_go_fuzz_dep_.CoverTab[17394]++
//line /usr/local/go/src/net/url/url.go:472
		// _ = "end of CoverTab[17394]"
//line /usr/local/go/src/net/url/url.go:472
	}
//line /usr/local/go/src/net/url/url.go:472
	// _ = "end of CoverTab[17389]"
//line /usr/local/go/src/net/url/url.go:472
	_go_fuzz_dep_.CoverTab[17390]++
						if frag == "" {
//line /usr/local/go/src/net/url/url.go:473
		_go_fuzz_dep_.CoverTab[17395]++
							return url, nil
//line /usr/local/go/src/net/url/url.go:474
		// _ = "end of CoverTab[17395]"
	} else {
//line /usr/local/go/src/net/url/url.go:475
		_go_fuzz_dep_.CoverTab[17396]++
//line /usr/local/go/src/net/url/url.go:475
		// _ = "end of CoverTab[17396]"
//line /usr/local/go/src/net/url/url.go:475
	}
//line /usr/local/go/src/net/url/url.go:475
	// _ = "end of CoverTab[17390]"
//line /usr/local/go/src/net/url/url.go:475
	_go_fuzz_dep_.CoverTab[17391]++
						if err = url.setFragment(frag); err != nil {
//line /usr/local/go/src/net/url/url.go:476
		_go_fuzz_dep_.CoverTab[17397]++
							return nil, &Error{"parse", rawURL, err}
//line /usr/local/go/src/net/url/url.go:477
		// _ = "end of CoverTab[17397]"
	} else {
//line /usr/local/go/src/net/url/url.go:478
		_go_fuzz_dep_.CoverTab[17398]++
//line /usr/local/go/src/net/url/url.go:478
		// _ = "end of CoverTab[17398]"
//line /usr/local/go/src/net/url/url.go:478
	}
//line /usr/local/go/src/net/url/url.go:478
	// _ = "end of CoverTab[17391]"
//line /usr/local/go/src/net/url/url.go:478
	_go_fuzz_dep_.CoverTab[17392]++
						return url, nil
//line /usr/local/go/src/net/url/url.go:479
	// _ = "end of CoverTab[17392]"
}

// ParseRequestURI parses a raw url into a URL structure. It assumes that
//line /usr/local/go/src/net/url/url.go:482
// url was received in an HTTP request, so the url is interpreted
//line /usr/local/go/src/net/url/url.go:482
// only as an absolute URI or an absolute path.
//line /usr/local/go/src/net/url/url.go:482
// The string url is assumed not to have a #fragment suffix.
//line /usr/local/go/src/net/url/url.go:482
// (Web browsers strip #fragment before sending the URL to a web server.)
//line /usr/local/go/src/net/url/url.go:487
func ParseRequestURI(rawURL string) (*URL, error) {
//line /usr/local/go/src/net/url/url.go:487
	_go_fuzz_dep_.CoverTab[17399]++
						url, err := parse(rawURL, true)
						if err != nil {
//line /usr/local/go/src/net/url/url.go:489
		_go_fuzz_dep_.CoverTab[17401]++
							return nil, &Error{"parse", rawURL, err}
//line /usr/local/go/src/net/url/url.go:490
		// _ = "end of CoverTab[17401]"
	} else {
//line /usr/local/go/src/net/url/url.go:491
		_go_fuzz_dep_.CoverTab[17402]++
//line /usr/local/go/src/net/url/url.go:491
		// _ = "end of CoverTab[17402]"
//line /usr/local/go/src/net/url/url.go:491
	}
//line /usr/local/go/src/net/url/url.go:491
	// _ = "end of CoverTab[17399]"
//line /usr/local/go/src/net/url/url.go:491
	_go_fuzz_dep_.CoverTab[17400]++
						return url, nil
//line /usr/local/go/src/net/url/url.go:492
	// _ = "end of CoverTab[17400]"
}

// parse parses a URL from a string in one of two contexts. If
//line /usr/local/go/src/net/url/url.go:495
// viaRequest is true, the URL is assumed to have arrived via an HTTP request,
//line /usr/local/go/src/net/url/url.go:495
// in which case only absolute URLs or path-absolute relative URLs are allowed.
//line /usr/local/go/src/net/url/url.go:495
// If viaRequest is false, all forms of relative URLs are allowed.
//line /usr/local/go/src/net/url/url.go:499
func parse(rawURL string, viaRequest bool) (*URL, error) {
//line /usr/local/go/src/net/url/url.go:499
	_go_fuzz_dep_.CoverTab[17403]++
						var rest string
						var err error

						if stringContainsCTLByte(rawURL) {
//line /usr/local/go/src/net/url/url.go:503
		_go_fuzz_dep_.CoverTab[17412]++
							return nil, errors.New("net/url: invalid control character in URL")
//line /usr/local/go/src/net/url/url.go:504
		// _ = "end of CoverTab[17412]"
	} else {
//line /usr/local/go/src/net/url/url.go:505
		_go_fuzz_dep_.CoverTab[17413]++
//line /usr/local/go/src/net/url/url.go:505
		// _ = "end of CoverTab[17413]"
//line /usr/local/go/src/net/url/url.go:505
	}
//line /usr/local/go/src/net/url/url.go:505
	// _ = "end of CoverTab[17403]"
//line /usr/local/go/src/net/url/url.go:505
	_go_fuzz_dep_.CoverTab[17404]++

						if rawURL == "" && func() bool {
//line /usr/local/go/src/net/url/url.go:507
		_go_fuzz_dep_.CoverTab[17414]++
//line /usr/local/go/src/net/url/url.go:507
		return viaRequest
//line /usr/local/go/src/net/url/url.go:507
		// _ = "end of CoverTab[17414]"
//line /usr/local/go/src/net/url/url.go:507
	}() {
//line /usr/local/go/src/net/url/url.go:507
		_go_fuzz_dep_.CoverTab[17415]++
							return nil, errors.New("empty url")
//line /usr/local/go/src/net/url/url.go:508
		// _ = "end of CoverTab[17415]"
	} else {
//line /usr/local/go/src/net/url/url.go:509
		_go_fuzz_dep_.CoverTab[17416]++
//line /usr/local/go/src/net/url/url.go:509
		// _ = "end of CoverTab[17416]"
//line /usr/local/go/src/net/url/url.go:509
	}
//line /usr/local/go/src/net/url/url.go:509
	// _ = "end of CoverTab[17404]"
//line /usr/local/go/src/net/url/url.go:509
	_go_fuzz_dep_.CoverTab[17405]++
						url := new(URL)

						if rawURL == "*" {
//line /usr/local/go/src/net/url/url.go:512
		_go_fuzz_dep_.CoverTab[17417]++
							url.Path = "*"
							return url, nil
//line /usr/local/go/src/net/url/url.go:514
		// _ = "end of CoverTab[17417]"
	} else {
//line /usr/local/go/src/net/url/url.go:515
		_go_fuzz_dep_.CoverTab[17418]++
//line /usr/local/go/src/net/url/url.go:515
		// _ = "end of CoverTab[17418]"
//line /usr/local/go/src/net/url/url.go:515
	}
//line /usr/local/go/src/net/url/url.go:515
	// _ = "end of CoverTab[17405]"
//line /usr/local/go/src/net/url/url.go:515
	_go_fuzz_dep_.CoverTab[17406]++

//line /usr/local/go/src/net/url/url.go:519
	if url.Scheme, rest, err = getScheme(rawURL); err != nil {
//line /usr/local/go/src/net/url/url.go:519
		_go_fuzz_dep_.CoverTab[17419]++
							return nil, err
//line /usr/local/go/src/net/url/url.go:520
		// _ = "end of CoverTab[17419]"
	} else {
//line /usr/local/go/src/net/url/url.go:521
		_go_fuzz_dep_.CoverTab[17420]++
//line /usr/local/go/src/net/url/url.go:521
		// _ = "end of CoverTab[17420]"
//line /usr/local/go/src/net/url/url.go:521
	}
//line /usr/local/go/src/net/url/url.go:521
	// _ = "end of CoverTab[17406]"
//line /usr/local/go/src/net/url/url.go:521
	_go_fuzz_dep_.CoverTab[17407]++
						url.Scheme = strings.ToLower(url.Scheme)

						if strings.HasSuffix(rest, "?") && func() bool {
//line /usr/local/go/src/net/url/url.go:524
		_go_fuzz_dep_.CoverTab[17421]++
//line /usr/local/go/src/net/url/url.go:524
		return strings.Count(rest, "?") == 1
//line /usr/local/go/src/net/url/url.go:524
		// _ = "end of CoverTab[17421]"
//line /usr/local/go/src/net/url/url.go:524
	}() {
//line /usr/local/go/src/net/url/url.go:524
		_go_fuzz_dep_.CoverTab[17422]++
							url.ForceQuery = true
							rest = rest[:len(rest)-1]
//line /usr/local/go/src/net/url/url.go:526
		// _ = "end of CoverTab[17422]"
	} else {
//line /usr/local/go/src/net/url/url.go:527
		_go_fuzz_dep_.CoverTab[17423]++
							rest, url.RawQuery, _ = strings.Cut(rest, "?")
//line /usr/local/go/src/net/url/url.go:528
		// _ = "end of CoverTab[17423]"
	}
//line /usr/local/go/src/net/url/url.go:529
	// _ = "end of CoverTab[17407]"
//line /usr/local/go/src/net/url/url.go:529
	_go_fuzz_dep_.CoverTab[17408]++

						if !strings.HasPrefix(rest, "/") {
//line /usr/local/go/src/net/url/url.go:531
		_go_fuzz_dep_.CoverTab[17424]++
							if url.Scheme != "" {
//line /usr/local/go/src/net/url/url.go:532
			_go_fuzz_dep_.CoverTab[17427]++

								url.Opaque = rest
								return url, nil
//line /usr/local/go/src/net/url/url.go:535
			// _ = "end of CoverTab[17427]"
		} else {
//line /usr/local/go/src/net/url/url.go:536
			_go_fuzz_dep_.CoverTab[17428]++
//line /usr/local/go/src/net/url/url.go:536
			// _ = "end of CoverTab[17428]"
//line /usr/local/go/src/net/url/url.go:536
		}
//line /usr/local/go/src/net/url/url.go:536
		// _ = "end of CoverTab[17424]"
//line /usr/local/go/src/net/url/url.go:536
		_go_fuzz_dep_.CoverTab[17425]++
							if viaRequest {
//line /usr/local/go/src/net/url/url.go:537
			_go_fuzz_dep_.CoverTab[17429]++
								return nil, errors.New("invalid URI for request")
//line /usr/local/go/src/net/url/url.go:538
			// _ = "end of CoverTab[17429]"
		} else {
//line /usr/local/go/src/net/url/url.go:539
			_go_fuzz_dep_.CoverTab[17430]++
//line /usr/local/go/src/net/url/url.go:539
			// _ = "end of CoverTab[17430]"
//line /usr/local/go/src/net/url/url.go:539
		}
//line /usr/local/go/src/net/url/url.go:539
		// _ = "end of CoverTab[17425]"
//line /usr/local/go/src/net/url/url.go:539
		_go_fuzz_dep_.CoverTab[17426]++

//line /usr/local/go/src/net/url/url.go:547
		if segment, _, _ := strings.Cut(rest, "/"); strings.Contains(segment, ":") {
//line /usr/local/go/src/net/url/url.go:547
			_go_fuzz_dep_.CoverTab[17431]++

								return nil, errors.New("first path segment in URL cannot contain colon")
//line /usr/local/go/src/net/url/url.go:549
			// _ = "end of CoverTab[17431]"
		} else {
//line /usr/local/go/src/net/url/url.go:550
			_go_fuzz_dep_.CoverTab[17432]++
//line /usr/local/go/src/net/url/url.go:550
			// _ = "end of CoverTab[17432]"
//line /usr/local/go/src/net/url/url.go:550
		}
//line /usr/local/go/src/net/url/url.go:550
		// _ = "end of CoverTab[17426]"
	} else {
//line /usr/local/go/src/net/url/url.go:551
		_go_fuzz_dep_.CoverTab[17433]++
//line /usr/local/go/src/net/url/url.go:551
		// _ = "end of CoverTab[17433]"
//line /usr/local/go/src/net/url/url.go:551
	}
//line /usr/local/go/src/net/url/url.go:551
	// _ = "end of CoverTab[17408]"
//line /usr/local/go/src/net/url/url.go:551
	_go_fuzz_dep_.CoverTab[17409]++

						if (url.Scheme != "" || func() bool {
//line /usr/local/go/src/net/url/url.go:553
		_go_fuzz_dep_.CoverTab[17434]++
//line /usr/local/go/src/net/url/url.go:553
		return !viaRequest && func() bool {
//line /usr/local/go/src/net/url/url.go:553
			_go_fuzz_dep_.CoverTab[17435]++
//line /usr/local/go/src/net/url/url.go:553
			return !strings.HasPrefix(rest, "///")
//line /usr/local/go/src/net/url/url.go:553
			// _ = "end of CoverTab[17435]"
//line /usr/local/go/src/net/url/url.go:553
		}()
//line /usr/local/go/src/net/url/url.go:553
		// _ = "end of CoverTab[17434]"
//line /usr/local/go/src/net/url/url.go:553
	}()) && func() bool {
//line /usr/local/go/src/net/url/url.go:553
		_go_fuzz_dep_.CoverTab[17436]++
//line /usr/local/go/src/net/url/url.go:553
		return strings.HasPrefix(rest, "//")
//line /usr/local/go/src/net/url/url.go:553
		// _ = "end of CoverTab[17436]"
//line /usr/local/go/src/net/url/url.go:553
	}() {
//line /usr/local/go/src/net/url/url.go:553
		_go_fuzz_dep_.CoverTab[17437]++
							var authority string
							authority, rest = rest[2:], ""
							if i := strings.Index(authority, "/"); i >= 0 {
//line /usr/local/go/src/net/url/url.go:556
			_go_fuzz_dep_.CoverTab[17439]++
								authority, rest = authority[:i], authority[i:]
//line /usr/local/go/src/net/url/url.go:557
			// _ = "end of CoverTab[17439]"
		} else {
//line /usr/local/go/src/net/url/url.go:558
			_go_fuzz_dep_.CoverTab[17440]++
//line /usr/local/go/src/net/url/url.go:558
			// _ = "end of CoverTab[17440]"
//line /usr/local/go/src/net/url/url.go:558
		}
//line /usr/local/go/src/net/url/url.go:558
		// _ = "end of CoverTab[17437]"
//line /usr/local/go/src/net/url/url.go:558
		_go_fuzz_dep_.CoverTab[17438]++
							url.User, url.Host, err = parseAuthority(authority)
							if err != nil {
//line /usr/local/go/src/net/url/url.go:560
			_go_fuzz_dep_.CoverTab[17441]++
								return nil, err
//line /usr/local/go/src/net/url/url.go:561
			// _ = "end of CoverTab[17441]"
		} else {
//line /usr/local/go/src/net/url/url.go:562
			_go_fuzz_dep_.CoverTab[17442]++
//line /usr/local/go/src/net/url/url.go:562
			// _ = "end of CoverTab[17442]"
//line /usr/local/go/src/net/url/url.go:562
		}
//line /usr/local/go/src/net/url/url.go:562
		// _ = "end of CoverTab[17438]"
	} else {
//line /usr/local/go/src/net/url/url.go:563
		_go_fuzz_dep_.CoverTab[17443]++
//line /usr/local/go/src/net/url/url.go:563
		if url.Scheme != "" && func() bool {
//line /usr/local/go/src/net/url/url.go:563
			_go_fuzz_dep_.CoverTab[17444]++
//line /usr/local/go/src/net/url/url.go:563
			return strings.HasPrefix(rest, "/")
//line /usr/local/go/src/net/url/url.go:563
			// _ = "end of CoverTab[17444]"
//line /usr/local/go/src/net/url/url.go:563
		}() {
//line /usr/local/go/src/net/url/url.go:563
			_go_fuzz_dep_.CoverTab[17445]++

//line /usr/local/go/src/net/url/url.go:566
			url.OmitHost = true
//line /usr/local/go/src/net/url/url.go:566
			// _ = "end of CoverTab[17445]"
		} else {
//line /usr/local/go/src/net/url/url.go:567
			_go_fuzz_dep_.CoverTab[17446]++
//line /usr/local/go/src/net/url/url.go:567
			// _ = "end of CoverTab[17446]"
//line /usr/local/go/src/net/url/url.go:567
		}
//line /usr/local/go/src/net/url/url.go:567
		// _ = "end of CoverTab[17443]"
//line /usr/local/go/src/net/url/url.go:567
	}
//line /usr/local/go/src/net/url/url.go:567
	// _ = "end of CoverTab[17409]"
//line /usr/local/go/src/net/url/url.go:567
	_go_fuzz_dep_.CoverTab[17410]++

//line /usr/local/go/src/net/url/url.go:573
	if err := url.setPath(rest); err != nil {
//line /usr/local/go/src/net/url/url.go:573
		_go_fuzz_dep_.CoverTab[17447]++
							return nil, err
//line /usr/local/go/src/net/url/url.go:574
		// _ = "end of CoverTab[17447]"
	} else {
//line /usr/local/go/src/net/url/url.go:575
		_go_fuzz_dep_.CoverTab[17448]++
//line /usr/local/go/src/net/url/url.go:575
		// _ = "end of CoverTab[17448]"
//line /usr/local/go/src/net/url/url.go:575
	}
//line /usr/local/go/src/net/url/url.go:575
	// _ = "end of CoverTab[17410]"
//line /usr/local/go/src/net/url/url.go:575
	_go_fuzz_dep_.CoverTab[17411]++
						return url, nil
//line /usr/local/go/src/net/url/url.go:576
	// _ = "end of CoverTab[17411]"
}

func parseAuthority(authority string) (user *Userinfo, host string, err error) {
//line /usr/local/go/src/net/url/url.go:579
	_go_fuzz_dep_.CoverTab[17449]++
						i := strings.LastIndex(authority, "@")
						if i < 0 {
//line /usr/local/go/src/net/url/url.go:581
		_go_fuzz_dep_.CoverTab[17455]++
							host, err = parseHost(authority)
//line /usr/local/go/src/net/url/url.go:582
		// _ = "end of CoverTab[17455]"
	} else {
//line /usr/local/go/src/net/url/url.go:583
		_go_fuzz_dep_.CoverTab[17456]++
							host, err = parseHost(authority[i+1:])
//line /usr/local/go/src/net/url/url.go:584
		// _ = "end of CoverTab[17456]"
	}
//line /usr/local/go/src/net/url/url.go:585
	// _ = "end of CoverTab[17449]"
//line /usr/local/go/src/net/url/url.go:585
	_go_fuzz_dep_.CoverTab[17450]++
						if err != nil {
//line /usr/local/go/src/net/url/url.go:586
		_go_fuzz_dep_.CoverTab[17457]++
							return nil, "", err
//line /usr/local/go/src/net/url/url.go:587
		// _ = "end of CoverTab[17457]"
	} else {
//line /usr/local/go/src/net/url/url.go:588
		_go_fuzz_dep_.CoverTab[17458]++
//line /usr/local/go/src/net/url/url.go:588
		// _ = "end of CoverTab[17458]"
//line /usr/local/go/src/net/url/url.go:588
	}
//line /usr/local/go/src/net/url/url.go:588
	// _ = "end of CoverTab[17450]"
//line /usr/local/go/src/net/url/url.go:588
	_go_fuzz_dep_.CoverTab[17451]++
						if i < 0 {
//line /usr/local/go/src/net/url/url.go:589
		_go_fuzz_dep_.CoverTab[17459]++
							return nil, host, nil
//line /usr/local/go/src/net/url/url.go:590
		// _ = "end of CoverTab[17459]"
	} else {
//line /usr/local/go/src/net/url/url.go:591
		_go_fuzz_dep_.CoverTab[17460]++
//line /usr/local/go/src/net/url/url.go:591
		// _ = "end of CoverTab[17460]"
//line /usr/local/go/src/net/url/url.go:591
	}
//line /usr/local/go/src/net/url/url.go:591
	// _ = "end of CoverTab[17451]"
//line /usr/local/go/src/net/url/url.go:591
	_go_fuzz_dep_.CoverTab[17452]++
						userinfo := authority[:i]
						if !validUserinfo(userinfo) {
//line /usr/local/go/src/net/url/url.go:593
		_go_fuzz_dep_.CoverTab[17461]++
							return nil, "", errors.New("net/url: invalid userinfo")
//line /usr/local/go/src/net/url/url.go:594
		// _ = "end of CoverTab[17461]"
	} else {
//line /usr/local/go/src/net/url/url.go:595
		_go_fuzz_dep_.CoverTab[17462]++
//line /usr/local/go/src/net/url/url.go:595
		// _ = "end of CoverTab[17462]"
//line /usr/local/go/src/net/url/url.go:595
	}
//line /usr/local/go/src/net/url/url.go:595
	// _ = "end of CoverTab[17452]"
//line /usr/local/go/src/net/url/url.go:595
	_go_fuzz_dep_.CoverTab[17453]++
						if !strings.Contains(userinfo, ":") {
//line /usr/local/go/src/net/url/url.go:596
		_go_fuzz_dep_.CoverTab[17463]++
							if userinfo, err = unescape(userinfo, encodeUserPassword); err != nil {
//line /usr/local/go/src/net/url/url.go:597
			_go_fuzz_dep_.CoverTab[17465]++
								return nil, "", err
//line /usr/local/go/src/net/url/url.go:598
			// _ = "end of CoverTab[17465]"
		} else {
//line /usr/local/go/src/net/url/url.go:599
			_go_fuzz_dep_.CoverTab[17466]++
//line /usr/local/go/src/net/url/url.go:599
			// _ = "end of CoverTab[17466]"
//line /usr/local/go/src/net/url/url.go:599
		}
//line /usr/local/go/src/net/url/url.go:599
		// _ = "end of CoverTab[17463]"
//line /usr/local/go/src/net/url/url.go:599
		_go_fuzz_dep_.CoverTab[17464]++
							user = User(userinfo)
//line /usr/local/go/src/net/url/url.go:600
		// _ = "end of CoverTab[17464]"
	} else {
//line /usr/local/go/src/net/url/url.go:601
		_go_fuzz_dep_.CoverTab[17467]++
							username, password, _ := strings.Cut(userinfo, ":")
							if username, err = unescape(username, encodeUserPassword); err != nil {
//line /usr/local/go/src/net/url/url.go:603
			_go_fuzz_dep_.CoverTab[17470]++
								return nil, "", err
//line /usr/local/go/src/net/url/url.go:604
			// _ = "end of CoverTab[17470]"
		} else {
//line /usr/local/go/src/net/url/url.go:605
			_go_fuzz_dep_.CoverTab[17471]++
//line /usr/local/go/src/net/url/url.go:605
			// _ = "end of CoverTab[17471]"
//line /usr/local/go/src/net/url/url.go:605
		}
//line /usr/local/go/src/net/url/url.go:605
		// _ = "end of CoverTab[17467]"
//line /usr/local/go/src/net/url/url.go:605
		_go_fuzz_dep_.CoverTab[17468]++
							if password, err = unescape(password, encodeUserPassword); err != nil {
//line /usr/local/go/src/net/url/url.go:606
			_go_fuzz_dep_.CoverTab[17472]++
								return nil, "", err
//line /usr/local/go/src/net/url/url.go:607
			// _ = "end of CoverTab[17472]"
		} else {
//line /usr/local/go/src/net/url/url.go:608
			_go_fuzz_dep_.CoverTab[17473]++
//line /usr/local/go/src/net/url/url.go:608
			// _ = "end of CoverTab[17473]"
//line /usr/local/go/src/net/url/url.go:608
		}
//line /usr/local/go/src/net/url/url.go:608
		// _ = "end of CoverTab[17468]"
//line /usr/local/go/src/net/url/url.go:608
		_go_fuzz_dep_.CoverTab[17469]++
							user = UserPassword(username, password)
//line /usr/local/go/src/net/url/url.go:609
		// _ = "end of CoverTab[17469]"
	}
//line /usr/local/go/src/net/url/url.go:610
	// _ = "end of CoverTab[17453]"
//line /usr/local/go/src/net/url/url.go:610
	_go_fuzz_dep_.CoverTab[17454]++
						return user, host, nil
//line /usr/local/go/src/net/url/url.go:611
	// _ = "end of CoverTab[17454]"
}

// parseHost parses host as an authority without user
//line /usr/local/go/src/net/url/url.go:614
// information. That is, as host[:port].
//line /usr/local/go/src/net/url/url.go:616
func parseHost(host string) (string, error) {
//line /usr/local/go/src/net/url/url.go:616
	_go_fuzz_dep_.CoverTab[17474]++
						if strings.HasPrefix(host, "[") {
//line /usr/local/go/src/net/url/url.go:617
		_go_fuzz_dep_.CoverTab[17477]++

//line /usr/local/go/src/net/url/url.go:620
		i := strings.LastIndex(host, "]")
		if i < 0 {
//line /usr/local/go/src/net/url/url.go:621
			_go_fuzz_dep_.CoverTab[17480]++
								return "", errors.New("missing ']' in host")
//line /usr/local/go/src/net/url/url.go:622
			// _ = "end of CoverTab[17480]"
		} else {
//line /usr/local/go/src/net/url/url.go:623
			_go_fuzz_dep_.CoverTab[17481]++
//line /usr/local/go/src/net/url/url.go:623
			// _ = "end of CoverTab[17481]"
//line /usr/local/go/src/net/url/url.go:623
		}
//line /usr/local/go/src/net/url/url.go:623
		// _ = "end of CoverTab[17477]"
//line /usr/local/go/src/net/url/url.go:623
		_go_fuzz_dep_.CoverTab[17478]++
							colonPort := host[i+1:]
							if !validOptionalPort(colonPort) {
//line /usr/local/go/src/net/url/url.go:625
			_go_fuzz_dep_.CoverTab[17482]++
								return "", fmt.Errorf("invalid port %q after host", colonPort)
//line /usr/local/go/src/net/url/url.go:626
			// _ = "end of CoverTab[17482]"
		} else {
//line /usr/local/go/src/net/url/url.go:627
			_go_fuzz_dep_.CoverTab[17483]++
//line /usr/local/go/src/net/url/url.go:627
			// _ = "end of CoverTab[17483]"
//line /usr/local/go/src/net/url/url.go:627
		}
//line /usr/local/go/src/net/url/url.go:627
		// _ = "end of CoverTab[17478]"
//line /usr/local/go/src/net/url/url.go:627
		_go_fuzz_dep_.CoverTab[17479]++

//line /usr/local/go/src/net/url/url.go:635
		zone := strings.Index(host[:i], "%25")
		if zone >= 0 {
//line /usr/local/go/src/net/url/url.go:636
			_go_fuzz_dep_.CoverTab[17484]++
								host1, err := unescape(host[:zone], encodeHost)
								if err != nil {
//line /usr/local/go/src/net/url/url.go:638
				_go_fuzz_dep_.CoverTab[17488]++
									return "", err
//line /usr/local/go/src/net/url/url.go:639
				// _ = "end of CoverTab[17488]"
			} else {
//line /usr/local/go/src/net/url/url.go:640
				_go_fuzz_dep_.CoverTab[17489]++
//line /usr/local/go/src/net/url/url.go:640
				// _ = "end of CoverTab[17489]"
//line /usr/local/go/src/net/url/url.go:640
			}
//line /usr/local/go/src/net/url/url.go:640
			// _ = "end of CoverTab[17484]"
//line /usr/local/go/src/net/url/url.go:640
			_go_fuzz_dep_.CoverTab[17485]++
								host2, err := unescape(host[zone:i], encodeZone)
								if err != nil {
//line /usr/local/go/src/net/url/url.go:642
				_go_fuzz_dep_.CoverTab[17490]++
									return "", err
//line /usr/local/go/src/net/url/url.go:643
				// _ = "end of CoverTab[17490]"
			} else {
//line /usr/local/go/src/net/url/url.go:644
				_go_fuzz_dep_.CoverTab[17491]++
//line /usr/local/go/src/net/url/url.go:644
				// _ = "end of CoverTab[17491]"
//line /usr/local/go/src/net/url/url.go:644
			}
//line /usr/local/go/src/net/url/url.go:644
			// _ = "end of CoverTab[17485]"
//line /usr/local/go/src/net/url/url.go:644
			_go_fuzz_dep_.CoverTab[17486]++
								host3, err := unescape(host[i:], encodeHost)
								if err != nil {
//line /usr/local/go/src/net/url/url.go:646
				_go_fuzz_dep_.CoverTab[17492]++
									return "", err
//line /usr/local/go/src/net/url/url.go:647
				// _ = "end of CoverTab[17492]"
			} else {
//line /usr/local/go/src/net/url/url.go:648
				_go_fuzz_dep_.CoverTab[17493]++
//line /usr/local/go/src/net/url/url.go:648
				// _ = "end of CoverTab[17493]"
//line /usr/local/go/src/net/url/url.go:648
			}
//line /usr/local/go/src/net/url/url.go:648
			// _ = "end of CoverTab[17486]"
//line /usr/local/go/src/net/url/url.go:648
			_go_fuzz_dep_.CoverTab[17487]++
								return host1 + host2 + host3, nil
//line /usr/local/go/src/net/url/url.go:649
			// _ = "end of CoverTab[17487]"
		} else {
//line /usr/local/go/src/net/url/url.go:650
			_go_fuzz_dep_.CoverTab[17494]++
//line /usr/local/go/src/net/url/url.go:650
			// _ = "end of CoverTab[17494]"
//line /usr/local/go/src/net/url/url.go:650
		}
//line /usr/local/go/src/net/url/url.go:650
		// _ = "end of CoverTab[17479]"
	} else {
//line /usr/local/go/src/net/url/url.go:651
		_go_fuzz_dep_.CoverTab[17495]++
//line /usr/local/go/src/net/url/url.go:651
		if i := strings.LastIndex(host, ":"); i != -1 {
//line /usr/local/go/src/net/url/url.go:651
			_go_fuzz_dep_.CoverTab[17496]++
								colonPort := host[i:]
								if !validOptionalPort(colonPort) {
//line /usr/local/go/src/net/url/url.go:653
				_go_fuzz_dep_.CoverTab[17497]++
									return "", fmt.Errorf("invalid port %q after host", colonPort)
//line /usr/local/go/src/net/url/url.go:654
				// _ = "end of CoverTab[17497]"
			} else {
//line /usr/local/go/src/net/url/url.go:655
				_go_fuzz_dep_.CoverTab[17498]++
//line /usr/local/go/src/net/url/url.go:655
				// _ = "end of CoverTab[17498]"
//line /usr/local/go/src/net/url/url.go:655
			}
//line /usr/local/go/src/net/url/url.go:655
			// _ = "end of CoverTab[17496]"
		} else {
//line /usr/local/go/src/net/url/url.go:656
			_go_fuzz_dep_.CoverTab[17499]++
//line /usr/local/go/src/net/url/url.go:656
			// _ = "end of CoverTab[17499]"
//line /usr/local/go/src/net/url/url.go:656
		}
//line /usr/local/go/src/net/url/url.go:656
		// _ = "end of CoverTab[17495]"
//line /usr/local/go/src/net/url/url.go:656
	}
//line /usr/local/go/src/net/url/url.go:656
	// _ = "end of CoverTab[17474]"
//line /usr/local/go/src/net/url/url.go:656
	_go_fuzz_dep_.CoverTab[17475]++

						var err error
						if host, err = unescape(host, encodeHost); err != nil {
//line /usr/local/go/src/net/url/url.go:659
		_go_fuzz_dep_.CoverTab[17500]++
							return "", err
//line /usr/local/go/src/net/url/url.go:660
		// _ = "end of CoverTab[17500]"
	} else {
//line /usr/local/go/src/net/url/url.go:661
		_go_fuzz_dep_.CoverTab[17501]++
//line /usr/local/go/src/net/url/url.go:661
		// _ = "end of CoverTab[17501]"
//line /usr/local/go/src/net/url/url.go:661
	}
//line /usr/local/go/src/net/url/url.go:661
	// _ = "end of CoverTab[17475]"
//line /usr/local/go/src/net/url/url.go:661
	_go_fuzz_dep_.CoverTab[17476]++
						return host, nil
//line /usr/local/go/src/net/url/url.go:662
	// _ = "end of CoverTab[17476]"
}

// setPath sets the Path and RawPath fields of the URL based on the provided
//line /usr/local/go/src/net/url/url.go:665
// escaped path p. It maintains the invariant that RawPath is only specified
//line /usr/local/go/src/net/url/url.go:665
// when it differs from the default encoding of the path.
//line /usr/local/go/src/net/url/url.go:665
// For example:
//line /usr/local/go/src/net/url/url.go:665
// - setPath("/foo/bar")   will set Path="/foo/bar" and RawPath=""
//line /usr/local/go/src/net/url/url.go:665
// - setPath("/foo%2fbar") will set Path="/foo/bar" and RawPath="/foo%2fbar"
//line /usr/local/go/src/net/url/url.go:665
// setPath will return an error only if the provided path contains an invalid
//line /usr/local/go/src/net/url/url.go:665
// escaping.
//line /usr/local/go/src/net/url/url.go:673
func (u *URL) setPath(p string) error {
//line /usr/local/go/src/net/url/url.go:673
	_go_fuzz_dep_.CoverTab[17502]++
						path, err := unescape(p, encodePath)
						if err != nil {
//line /usr/local/go/src/net/url/url.go:675
		_go_fuzz_dep_.CoverTab[17505]++
							return err
//line /usr/local/go/src/net/url/url.go:676
		// _ = "end of CoverTab[17505]"
	} else {
//line /usr/local/go/src/net/url/url.go:677
		_go_fuzz_dep_.CoverTab[17506]++
//line /usr/local/go/src/net/url/url.go:677
		// _ = "end of CoverTab[17506]"
//line /usr/local/go/src/net/url/url.go:677
	}
//line /usr/local/go/src/net/url/url.go:677
	// _ = "end of CoverTab[17502]"
//line /usr/local/go/src/net/url/url.go:677
	_go_fuzz_dep_.CoverTab[17503]++
						u.Path = path
						if escp := escape(path, encodePath); p == escp {
//line /usr/local/go/src/net/url/url.go:679
		_go_fuzz_dep_.CoverTab[17507]++

							u.RawPath = ""
//line /usr/local/go/src/net/url/url.go:681
		// _ = "end of CoverTab[17507]"
	} else {
//line /usr/local/go/src/net/url/url.go:682
		_go_fuzz_dep_.CoverTab[17508]++
							u.RawPath = p
//line /usr/local/go/src/net/url/url.go:683
		// _ = "end of CoverTab[17508]"
	}
//line /usr/local/go/src/net/url/url.go:684
	// _ = "end of CoverTab[17503]"
//line /usr/local/go/src/net/url/url.go:684
	_go_fuzz_dep_.CoverTab[17504]++
						return nil
//line /usr/local/go/src/net/url/url.go:685
	// _ = "end of CoverTab[17504]"
}

// EscapedPath returns the escaped form of u.Path.
//line /usr/local/go/src/net/url/url.go:688
// In general there are multiple possible escaped forms of any path.
//line /usr/local/go/src/net/url/url.go:688
// EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
//line /usr/local/go/src/net/url/url.go:688
// Otherwise EscapedPath ignores u.RawPath and computes an escaped
//line /usr/local/go/src/net/url/url.go:688
// form on its own.
//line /usr/local/go/src/net/url/url.go:688
// The String and RequestURI methods use EscapedPath to construct
//line /usr/local/go/src/net/url/url.go:688
// their results.
//line /usr/local/go/src/net/url/url.go:688
// In general, code should call EscapedPath instead of
//line /usr/local/go/src/net/url/url.go:688
// reading u.RawPath directly.
//line /usr/local/go/src/net/url/url.go:697
func (u *URL) EscapedPath() string {
//line /usr/local/go/src/net/url/url.go:697
	_go_fuzz_dep_.CoverTab[17509]++
						if u.RawPath != "" && func() bool {
//line /usr/local/go/src/net/url/url.go:698
		_go_fuzz_dep_.CoverTab[17512]++
//line /usr/local/go/src/net/url/url.go:698
		return validEncoded(u.RawPath, encodePath)
//line /usr/local/go/src/net/url/url.go:698
		// _ = "end of CoverTab[17512]"
//line /usr/local/go/src/net/url/url.go:698
	}() {
//line /usr/local/go/src/net/url/url.go:698
		_go_fuzz_dep_.CoverTab[17513]++
							p, err := unescape(u.RawPath, encodePath)
							if err == nil && func() bool {
//line /usr/local/go/src/net/url/url.go:700
			_go_fuzz_dep_.CoverTab[17514]++
//line /usr/local/go/src/net/url/url.go:700
			return p == u.Path
//line /usr/local/go/src/net/url/url.go:700
			// _ = "end of CoverTab[17514]"
//line /usr/local/go/src/net/url/url.go:700
		}() {
//line /usr/local/go/src/net/url/url.go:700
			_go_fuzz_dep_.CoverTab[17515]++
								return u.RawPath
//line /usr/local/go/src/net/url/url.go:701
			// _ = "end of CoverTab[17515]"
		} else {
//line /usr/local/go/src/net/url/url.go:702
			_go_fuzz_dep_.CoverTab[17516]++
//line /usr/local/go/src/net/url/url.go:702
			// _ = "end of CoverTab[17516]"
//line /usr/local/go/src/net/url/url.go:702
		}
//line /usr/local/go/src/net/url/url.go:702
		// _ = "end of CoverTab[17513]"
	} else {
//line /usr/local/go/src/net/url/url.go:703
		_go_fuzz_dep_.CoverTab[17517]++
//line /usr/local/go/src/net/url/url.go:703
		// _ = "end of CoverTab[17517]"
//line /usr/local/go/src/net/url/url.go:703
	}
//line /usr/local/go/src/net/url/url.go:703
	// _ = "end of CoverTab[17509]"
//line /usr/local/go/src/net/url/url.go:703
	_go_fuzz_dep_.CoverTab[17510]++
						if u.Path == "*" {
//line /usr/local/go/src/net/url/url.go:704
		_go_fuzz_dep_.CoverTab[17518]++
							return "*"
//line /usr/local/go/src/net/url/url.go:705
		// _ = "end of CoverTab[17518]"
	} else {
//line /usr/local/go/src/net/url/url.go:706
		_go_fuzz_dep_.CoverTab[17519]++
//line /usr/local/go/src/net/url/url.go:706
		// _ = "end of CoverTab[17519]"
//line /usr/local/go/src/net/url/url.go:706
	}
//line /usr/local/go/src/net/url/url.go:706
	// _ = "end of CoverTab[17510]"
//line /usr/local/go/src/net/url/url.go:706
	_go_fuzz_dep_.CoverTab[17511]++
						return escape(u.Path, encodePath)
//line /usr/local/go/src/net/url/url.go:707
	// _ = "end of CoverTab[17511]"
}

// validEncoded reports whether s is a valid encoded path or fragment,
//line /usr/local/go/src/net/url/url.go:710
// according to mode.
//line /usr/local/go/src/net/url/url.go:710
// It must not contain any bytes that require escaping during encoding.
//line /usr/local/go/src/net/url/url.go:713
func validEncoded(s string, mode encoding) bool {
//line /usr/local/go/src/net/url/url.go:713
	_go_fuzz_dep_.CoverTab[17520]++
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/url/url.go:714
		_go_fuzz_dep_.CoverTab[17522]++

//line /usr/local/go/src/net/url/url.go:720
		switch s[i] {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '@':
//line /usr/local/go/src/net/url/url.go:721
			_go_fuzz_dep_.CoverTab[17523]++
//line /usr/local/go/src/net/url/url.go:721
			// _ = "end of CoverTab[17523]"

		case '[', ']':
//line /usr/local/go/src/net/url/url.go:723
			_go_fuzz_dep_.CoverTab[17524]++
//line /usr/local/go/src/net/url/url.go:723
			// _ = "end of CoverTab[17524]"

		case '%':
//line /usr/local/go/src/net/url/url.go:725
			_go_fuzz_dep_.CoverTab[17525]++
//line /usr/local/go/src/net/url/url.go:725
			// _ = "end of CoverTab[17525]"

		default:
//line /usr/local/go/src/net/url/url.go:727
			_go_fuzz_dep_.CoverTab[17526]++
								if shouldEscape(s[i], mode) {
//line /usr/local/go/src/net/url/url.go:728
				_go_fuzz_dep_.CoverTab[17527]++
									return false
//line /usr/local/go/src/net/url/url.go:729
				// _ = "end of CoverTab[17527]"
			} else {
//line /usr/local/go/src/net/url/url.go:730
				_go_fuzz_dep_.CoverTab[17528]++
//line /usr/local/go/src/net/url/url.go:730
				// _ = "end of CoverTab[17528]"
//line /usr/local/go/src/net/url/url.go:730
			}
//line /usr/local/go/src/net/url/url.go:730
			// _ = "end of CoverTab[17526]"
		}
//line /usr/local/go/src/net/url/url.go:731
		// _ = "end of CoverTab[17522]"
	}
//line /usr/local/go/src/net/url/url.go:732
	// _ = "end of CoverTab[17520]"
//line /usr/local/go/src/net/url/url.go:732
	_go_fuzz_dep_.CoverTab[17521]++
						return true
//line /usr/local/go/src/net/url/url.go:733
	// _ = "end of CoverTab[17521]"
}

// setFragment is like setPath but for Fragment/RawFragment.
func (u *URL) setFragment(f string) error {
//line /usr/local/go/src/net/url/url.go:737
	_go_fuzz_dep_.CoverTab[17529]++
						frag, err := unescape(f, encodeFragment)
						if err != nil {
//line /usr/local/go/src/net/url/url.go:739
		_go_fuzz_dep_.CoverTab[17532]++
							return err
//line /usr/local/go/src/net/url/url.go:740
		// _ = "end of CoverTab[17532]"
	} else {
//line /usr/local/go/src/net/url/url.go:741
		_go_fuzz_dep_.CoverTab[17533]++
//line /usr/local/go/src/net/url/url.go:741
		// _ = "end of CoverTab[17533]"
//line /usr/local/go/src/net/url/url.go:741
	}
//line /usr/local/go/src/net/url/url.go:741
	// _ = "end of CoverTab[17529]"
//line /usr/local/go/src/net/url/url.go:741
	_go_fuzz_dep_.CoverTab[17530]++
						u.Fragment = frag
						if escf := escape(frag, encodeFragment); f == escf {
//line /usr/local/go/src/net/url/url.go:743
		_go_fuzz_dep_.CoverTab[17534]++

							u.RawFragment = ""
//line /usr/local/go/src/net/url/url.go:745
		// _ = "end of CoverTab[17534]"
	} else {
//line /usr/local/go/src/net/url/url.go:746
		_go_fuzz_dep_.CoverTab[17535]++
							u.RawFragment = f
//line /usr/local/go/src/net/url/url.go:747
		// _ = "end of CoverTab[17535]"
	}
//line /usr/local/go/src/net/url/url.go:748
	// _ = "end of CoverTab[17530]"
//line /usr/local/go/src/net/url/url.go:748
	_go_fuzz_dep_.CoverTab[17531]++
						return nil
//line /usr/local/go/src/net/url/url.go:749
	// _ = "end of CoverTab[17531]"
}

// EscapedFragment returns the escaped form of u.Fragment.
//line /usr/local/go/src/net/url/url.go:752
// In general there are multiple possible escaped forms of any fragment.
//line /usr/local/go/src/net/url/url.go:752
// EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment.
//line /usr/local/go/src/net/url/url.go:752
// Otherwise EscapedFragment ignores u.RawFragment and computes an escaped
//line /usr/local/go/src/net/url/url.go:752
// form on its own.
//line /usr/local/go/src/net/url/url.go:752
// The String method uses EscapedFragment to construct its result.
//line /usr/local/go/src/net/url/url.go:752
// In general, code should call EscapedFragment instead of
//line /usr/local/go/src/net/url/url.go:752
// reading u.RawFragment directly.
//line /usr/local/go/src/net/url/url.go:760
func (u *URL) EscapedFragment() string {
//line /usr/local/go/src/net/url/url.go:760
	_go_fuzz_dep_.CoverTab[17536]++
						if u.RawFragment != "" && func() bool {
//line /usr/local/go/src/net/url/url.go:761
		_go_fuzz_dep_.CoverTab[17538]++
//line /usr/local/go/src/net/url/url.go:761
		return validEncoded(u.RawFragment, encodeFragment)
//line /usr/local/go/src/net/url/url.go:761
		// _ = "end of CoverTab[17538]"
//line /usr/local/go/src/net/url/url.go:761
	}() {
//line /usr/local/go/src/net/url/url.go:761
		_go_fuzz_dep_.CoverTab[17539]++
							f, err := unescape(u.RawFragment, encodeFragment)
							if err == nil && func() bool {
//line /usr/local/go/src/net/url/url.go:763
			_go_fuzz_dep_.CoverTab[17540]++
//line /usr/local/go/src/net/url/url.go:763
			return f == u.Fragment
//line /usr/local/go/src/net/url/url.go:763
			// _ = "end of CoverTab[17540]"
//line /usr/local/go/src/net/url/url.go:763
		}() {
//line /usr/local/go/src/net/url/url.go:763
			_go_fuzz_dep_.CoverTab[17541]++
								return u.RawFragment
//line /usr/local/go/src/net/url/url.go:764
			// _ = "end of CoverTab[17541]"
		} else {
//line /usr/local/go/src/net/url/url.go:765
			_go_fuzz_dep_.CoverTab[17542]++
//line /usr/local/go/src/net/url/url.go:765
			// _ = "end of CoverTab[17542]"
//line /usr/local/go/src/net/url/url.go:765
		}
//line /usr/local/go/src/net/url/url.go:765
		// _ = "end of CoverTab[17539]"
	} else {
//line /usr/local/go/src/net/url/url.go:766
		_go_fuzz_dep_.CoverTab[17543]++
//line /usr/local/go/src/net/url/url.go:766
		// _ = "end of CoverTab[17543]"
//line /usr/local/go/src/net/url/url.go:766
	}
//line /usr/local/go/src/net/url/url.go:766
	// _ = "end of CoverTab[17536]"
//line /usr/local/go/src/net/url/url.go:766
	_go_fuzz_dep_.CoverTab[17537]++
						return escape(u.Fragment, encodeFragment)
//line /usr/local/go/src/net/url/url.go:767
	// _ = "end of CoverTab[17537]"
}

// validOptionalPort reports whether port is either an empty string
//line /usr/local/go/src/net/url/url.go:770
// or matches /^:\d*$/
//line /usr/local/go/src/net/url/url.go:772
func validOptionalPort(port string) bool {
//line /usr/local/go/src/net/url/url.go:772
	_go_fuzz_dep_.CoverTab[17544]++
						if port == "" {
//line /usr/local/go/src/net/url/url.go:773
		_go_fuzz_dep_.CoverTab[17548]++
							return true
//line /usr/local/go/src/net/url/url.go:774
		// _ = "end of CoverTab[17548]"
	} else {
//line /usr/local/go/src/net/url/url.go:775
		_go_fuzz_dep_.CoverTab[17549]++
//line /usr/local/go/src/net/url/url.go:775
		// _ = "end of CoverTab[17549]"
//line /usr/local/go/src/net/url/url.go:775
	}
//line /usr/local/go/src/net/url/url.go:775
	// _ = "end of CoverTab[17544]"
//line /usr/local/go/src/net/url/url.go:775
	_go_fuzz_dep_.CoverTab[17545]++
						if port[0] != ':' {
//line /usr/local/go/src/net/url/url.go:776
		_go_fuzz_dep_.CoverTab[17550]++
							return false
//line /usr/local/go/src/net/url/url.go:777
		// _ = "end of CoverTab[17550]"
	} else {
//line /usr/local/go/src/net/url/url.go:778
		_go_fuzz_dep_.CoverTab[17551]++
//line /usr/local/go/src/net/url/url.go:778
		// _ = "end of CoverTab[17551]"
//line /usr/local/go/src/net/url/url.go:778
	}
//line /usr/local/go/src/net/url/url.go:778
	// _ = "end of CoverTab[17545]"
//line /usr/local/go/src/net/url/url.go:778
	_go_fuzz_dep_.CoverTab[17546]++
						for _, b := range port[1:] {
//line /usr/local/go/src/net/url/url.go:779
		_go_fuzz_dep_.CoverTab[17552]++
							if b < '0' || func() bool {
//line /usr/local/go/src/net/url/url.go:780
			_go_fuzz_dep_.CoverTab[17553]++
//line /usr/local/go/src/net/url/url.go:780
			return b > '9'
//line /usr/local/go/src/net/url/url.go:780
			// _ = "end of CoverTab[17553]"
//line /usr/local/go/src/net/url/url.go:780
		}() {
//line /usr/local/go/src/net/url/url.go:780
			_go_fuzz_dep_.CoverTab[17554]++
								return false
//line /usr/local/go/src/net/url/url.go:781
			// _ = "end of CoverTab[17554]"
		} else {
//line /usr/local/go/src/net/url/url.go:782
			_go_fuzz_dep_.CoverTab[17555]++
//line /usr/local/go/src/net/url/url.go:782
			// _ = "end of CoverTab[17555]"
//line /usr/local/go/src/net/url/url.go:782
		}
//line /usr/local/go/src/net/url/url.go:782
		// _ = "end of CoverTab[17552]"
	}
//line /usr/local/go/src/net/url/url.go:783
	// _ = "end of CoverTab[17546]"
//line /usr/local/go/src/net/url/url.go:783
	_go_fuzz_dep_.CoverTab[17547]++
						return true
//line /usr/local/go/src/net/url/url.go:784
	// _ = "end of CoverTab[17547]"
}

// String reassembles the URL into a valid URL string.
//line /usr/local/go/src/net/url/url.go:787
// The general form of the result is one of:
//line /usr/local/go/src/net/url/url.go:787
//
//line /usr/local/go/src/net/url/url.go:787
//	scheme:opaque?query#fragment
//line /usr/local/go/src/net/url/url.go:787
//	scheme://userinfo@host/path?query#fragment
//line /usr/local/go/src/net/url/url.go:787
//
//line /usr/local/go/src/net/url/url.go:787
// If u.Opaque is non-empty, String uses the first form;
//line /usr/local/go/src/net/url/url.go:787
// otherwise it uses the second form.
//line /usr/local/go/src/net/url/url.go:787
// Any non-ASCII characters in host are escaped.
//line /usr/local/go/src/net/url/url.go:787
// To obtain the path, String uses u.EscapedPath().
//line /usr/local/go/src/net/url/url.go:787
//
//line /usr/local/go/src/net/url/url.go:787
// In the second form, the following rules apply:
//line /usr/local/go/src/net/url/url.go:787
//   - if u.Scheme is empty, scheme: is omitted.
//line /usr/local/go/src/net/url/url.go:787
//   - if u.User is nil, userinfo@ is omitted.
//line /usr/local/go/src/net/url/url.go:787
//   - if u.Host is empty, host/ is omitted.
//line /usr/local/go/src/net/url/url.go:787
//   - if u.Scheme and u.Host are empty and u.User is nil,
//line /usr/local/go/src/net/url/url.go:787
//     the entire scheme://userinfo@host/ is omitted.
//line /usr/local/go/src/net/url/url.go:787
//   - if u.Host is non-empty and u.Path begins with a /,
//line /usr/local/go/src/net/url/url.go:787
//     the form host/path does not add its own /.
//line /usr/local/go/src/net/url/url.go:787
//   - if u.RawQuery is empty, ?query is omitted.
//line /usr/local/go/src/net/url/url.go:787
//   - if u.Fragment is empty, #fragment is omitted.
//line /usr/local/go/src/net/url/url.go:808
func (u *URL) String() string {
//line /usr/local/go/src/net/url/url.go:808
	_go_fuzz_dep_.CoverTab[17556]++
						var buf strings.Builder
						if u.Scheme != "" {
//line /usr/local/go/src/net/url/url.go:810
		_go_fuzz_dep_.CoverTab[17561]++
							buf.WriteString(u.Scheme)
							buf.WriteByte(':')
//line /usr/local/go/src/net/url/url.go:812
		// _ = "end of CoverTab[17561]"
	} else {
//line /usr/local/go/src/net/url/url.go:813
		_go_fuzz_dep_.CoverTab[17562]++
//line /usr/local/go/src/net/url/url.go:813
		// _ = "end of CoverTab[17562]"
//line /usr/local/go/src/net/url/url.go:813
	}
//line /usr/local/go/src/net/url/url.go:813
	// _ = "end of CoverTab[17556]"
//line /usr/local/go/src/net/url/url.go:813
	_go_fuzz_dep_.CoverTab[17557]++
						if u.Opaque != "" {
//line /usr/local/go/src/net/url/url.go:814
		_go_fuzz_dep_.CoverTab[17563]++
							buf.WriteString(u.Opaque)
//line /usr/local/go/src/net/url/url.go:815
		// _ = "end of CoverTab[17563]"
	} else {
//line /usr/local/go/src/net/url/url.go:816
		_go_fuzz_dep_.CoverTab[17564]++
							if u.Scheme != "" || func() bool {
//line /usr/local/go/src/net/url/url.go:817
			_go_fuzz_dep_.CoverTab[17568]++
//line /usr/local/go/src/net/url/url.go:817
			return u.Host != ""
//line /usr/local/go/src/net/url/url.go:817
			// _ = "end of CoverTab[17568]"
//line /usr/local/go/src/net/url/url.go:817
		}() || func() bool {
//line /usr/local/go/src/net/url/url.go:817
			_go_fuzz_dep_.CoverTab[17569]++
//line /usr/local/go/src/net/url/url.go:817
			return u.User != nil
//line /usr/local/go/src/net/url/url.go:817
			// _ = "end of CoverTab[17569]"
//line /usr/local/go/src/net/url/url.go:817
		}() {
//line /usr/local/go/src/net/url/url.go:817
			_go_fuzz_dep_.CoverTab[17570]++
								if u.OmitHost && func() bool {
//line /usr/local/go/src/net/url/url.go:818
				_go_fuzz_dep_.CoverTab[17571]++
//line /usr/local/go/src/net/url/url.go:818
				return u.Host == ""
//line /usr/local/go/src/net/url/url.go:818
				// _ = "end of CoverTab[17571]"
//line /usr/local/go/src/net/url/url.go:818
			}() && func() bool {
//line /usr/local/go/src/net/url/url.go:818
				_go_fuzz_dep_.CoverTab[17572]++
//line /usr/local/go/src/net/url/url.go:818
				return u.User == nil
//line /usr/local/go/src/net/url/url.go:818
				// _ = "end of CoverTab[17572]"
//line /usr/local/go/src/net/url/url.go:818
			}() {
//line /usr/local/go/src/net/url/url.go:818
				_go_fuzz_dep_.CoverTab[17573]++
//line /usr/local/go/src/net/url/url.go:818
				// _ = "end of CoverTab[17573]"

			} else {
//line /usr/local/go/src/net/url/url.go:820
				_go_fuzz_dep_.CoverTab[17574]++
									if u.Host != "" || func() bool {
//line /usr/local/go/src/net/url/url.go:821
					_go_fuzz_dep_.CoverTab[17577]++
//line /usr/local/go/src/net/url/url.go:821
					return u.Path != ""
//line /usr/local/go/src/net/url/url.go:821
					// _ = "end of CoverTab[17577]"
//line /usr/local/go/src/net/url/url.go:821
				}() || func() bool {
//line /usr/local/go/src/net/url/url.go:821
					_go_fuzz_dep_.CoverTab[17578]++
//line /usr/local/go/src/net/url/url.go:821
					return u.User != nil
//line /usr/local/go/src/net/url/url.go:821
					// _ = "end of CoverTab[17578]"
//line /usr/local/go/src/net/url/url.go:821
				}() {
//line /usr/local/go/src/net/url/url.go:821
					_go_fuzz_dep_.CoverTab[17579]++
										buf.WriteString("//")
//line /usr/local/go/src/net/url/url.go:822
					// _ = "end of CoverTab[17579]"
				} else {
//line /usr/local/go/src/net/url/url.go:823
					_go_fuzz_dep_.CoverTab[17580]++
//line /usr/local/go/src/net/url/url.go:823
					// _ = "end of CoverTab[17580]"
//line /usr/local/go/src/net/url/url.go:823
				}
//line /usr/local/go/src/net/url/url.go:823
				// _ = "end of CoverTab[17574]"
//line /usr/local/go/src/net/url/url.go:823
				_go_fuzz_dep_.CoverTab[17575]++
									if ui := u.User; ui != nil {
//line /usr/local/go/src/net/url/url.go:824
					_go_fuzz_dep_.CoverTab[17581]++
										buf.WriteString(ui.String())
										buf.WriteByte('@')
//line /usr/local/go/src/net/url/url.go:826
					// _ = "end of CoverTab[17581]"
				} else {
//line /usr/local/go/src/net/url/url.go:827
					_go_fuzz_dep_.CoverTab[17582]++
//line /usr/local/go/src/net/url/url.go:827
					// _ = "end of CoverTab[17582]"
//line /usr/local/go/src/net/url/url.go:827
				}
//line /usr/local/go/src/net/url/url.go:827
				// _ = "end of CoverTab[17575]"
//line /usr/local/go/src/net/url/url.go:827
				_go_fuzz_dep_.CoverTab[17576]++
									if h := u.Host; h != "" {
//line /usr/local/go/src/net/url/url.go:828
					_go_fuzz_dep_.CoverTab[17583]++
										buf.WriteString(escape(h, encodeHost))
//line /usr/local/go/src/net/url/url.go:829
					// _ = "end of CoverTab[17583]"
				} else {
//line /usr/local/go/src/net/url/url.go:830
					_go_fuzz_dep_.CoverTab[17584]++
//line /usr/local/go/src/net/url/url.go:830
					// _ = "end of CoverTab[17584]"
//line /usr/local/go/src/net/url/url.go:830
				}
//line /usr/local/go/src/net/url/url.go:830
				// _ = "end of CoverTab[17576]"
			}
//line /usr/local/go/src/net/url/url.go:831
			// _ = "end of CoverTab[17570]"
		} else {
//line /usr/local/go/src/net/url/url.go:832
			_go_fuzz_dep_.CoverTab[17585]++
//line /usr/local/go/src/net/url/url.go:832
			// _ = "end of CoverTab[17585]"
//line /usr/local/go/src/net/url/url.go:832
		}
//line /usr/local/go/src/net/url/url.go:832
		// _ = "end of CoverTab[17564]"
//line /usr/local/go/src/net/url/url.go:832
		_go_fuzz_dep_.CoverTab[17565]++
							path := u.EscapedPath()
							if path != "" && func() bool {
//line /usr/local/go/src/net/url/url.go:834
			_go_fuzz_dep_.CoverTab[17586]++
//line /usr/local/go/src/net/url/url.go:834
			return path[0] != '/'
//line /usr/local/go/src/net/url/url.go:834
			// _ = "end of CoverTab[17586]"
//line /usr/local/go/src/net/url/url.go:834
		}() && func() bool {
//line /usr/local/go/src/net/url/url.go:834
			_go_fuzz_dep_.CoverTab[17587]++
//line /usr/local/go/src/net/url/url.go:834
			return u.Host != ""
//line /usr/local/go/src/net/url/url.go:834
			// _ = "end of CoverTab[17587]"
//line /usr/local/go/src/net/url/url.go:834
		}() {
//line /usr/local/go/src/net/url/url.go:834
			_go_fuzz_dep_.CoverTab[17588]++
								buf.WriteByte('/')
//line /usr/local/go/src/net/url/url.go:835
			// _ = "end of CoverTab[17588]"
		} else {
//line /usr/local/go/src/net/url/url.go:836
			_go_fuzz_dep_.CoverTab[17589]++
//line /usr/local/go/src/net/url/url.go:836
			// _ = "end of CoverTab[17589]"
//line /usr/local/go/src/net/url/url.go:836
		}
//line /usr/local/go/src/net/url/url.go:836
		// _ = "end of CoverTab[17565]"
//line /usr/local/go/src/net/url/url.go:836
		_go_fuzz_dep_.CoverTab[17566]++
							if buf.Len() == 0 {
//line /usr/local/go/src/net/url/url.go:837
			_go_fuzz_dep_.CoverTab[17590]++

//line /usr/local/go/src/net/url/url.go:844
			if segment, _, _ := strings.Cut(path, "/"); strings.Contains(segment, ":") {
//line /usr/local/go/src/net/url/url.go:844
				_go_fuzz_dep_.CoverTab[17591]++
									buf.WriteString("./")
//line /usr/local/go/src/net/url/url.go:845
				// _ = "end of CoverTab[17591]"
			} else {
//line /usr/local/go/src/net/url/url.go:846
				_go_fuzz_dep_.CoverTab[17592]++
//line /usr/local/go/src/net/url/url.go:846
				// _ = "end of CoverTab[17592]"
//line /usr/local/go/src/net/url/url.go:846
			}
//line /usr/local/go/src/net/url/url.go:846
			// _ = "end of CoverTab[17590]"
		} else {
//line /usr/local/go/src/net/url/url.go:847
			_go_fuzz_dep_.CoverTab[17593]++
//line /usr/local/go/src/net/url/url.go:847
			// _ = "end of CoverTab[17593]"
//line /usr/local/go/src/net/url/url.go:847
		}
//line /usr/local/go/src/net/url/url.go:847
		// _ = "end of CoverTab[17566]"
//line /usr/local/go/src/net/url/url.go:847
		_go_fuzz_dep_.CoverTab[17567]++
							buf.WriteString(path)
//line /usr/local/go/src/net/url/url.go:848
		// _ = "end of CoverTab[17567]"
	}
//line /usr/local/go/src/net/url/url.go:849
	// _ = "end of CoverTab[17557]"
//line /usr/local/go/src/net/url/url.go:849
	_go_fuzz_dep_.CoverTab[17558]++
						if u.ForceQuery || func() bool {
//line /usr/local/go/src/net/url/url.go:850
		_go_fuzz_dep_.CoverTab[17594]++
//line /usr/local/go/src/net/url/url.go:850
		return u.RawQuery != ""
//line /usr/local/go/src/net/url/url.go:850
		// _ = "end of CoverTab[17594]"
//line /usr/local/go/src/net/url/url.go:850
	}() {
//line /usr/local/go/src/net/url/url.go:850
		_go_fuzz_dep_.CoverTab[17595]++
							buf.WriteByte('?')
							buf.WriteString(u.RawQuery)
//line /usr/local/go/src/net/url/url.go:852
		// _ = "end of CoverTab[17595]"
	} else {
//line /usr/local/go/src/net/url/url.go:853
		_go_fuzz_dep_.CoverTab[17596]++
//line /usr/local/go/src/net/url/url.go:853
		// _ = "end of CoverTab[17596]"
//line /usr/local/go/src/net/url/url.go:853
	}
//line /usr/local/go/src/net/url/url.go:853
	// _ = "end of CoverTab[17558]"
//line /usr/local/go/src/net/url/url.go:853
	_go_fuzz_dep_.CoverTab[17559]++
						if u.Fragment != "" {
//line /usr/local/go/src/net/url/url.go:854
		_go_fuzz_dep_.CoverTab[17597]++
							buf.WriteByte('#')
							buf.WriteString(u.EscapedFragment())
//line /usr/local/go/src/net/url/url.go:856
		// _ = "end of CoverTab[17597]"
	} else {
//line /usr/local/go/src/net/url/url.go:857
		_go_fuzz_dep_.CoverTab[17598]++
//line /usr/local/go/src/net/url/url.go:857
		// _ = "end of CoverTab[17598]"
//line /usr/local/go/src/net/url/url.go:857
	}
//line /usr/local/go/src/net/url/url.go:857
	// _ = "end of CoverTab[17559]"
//line /usr/local/go/src/net/url/url.go:857
	_go_fuzz_dep_.CoverTab[17560]++
						return buf.String()
//line /usr/local/go/src/net/url/url.go:858
	// _ = "end of CoverTab[17560]"
}

// Redacted is like String but replaces any password with "xxxxx".
//line /usr/local/go/src/net/url/url.go:861
// Only the password in u.URL is redacted.
//line /usr/local/go/src/net/url/url.go:863
func (u *URL) Redacted() string {
//line /usr/local/go/src/net/url/url.go:863
	_go_fuzz_dep_.CoverTab[17599]++
						if u == nil {
//line /usr/local/go/src/net/url/url.go:864
		_go_fuzz_dep_.CoverTab[17602]++
							return ""
//line /usr/local/go/src/net/url/url.go:865
		// _ = "end of CoverTab[17602]"
	} else {
//line /usr/local/go/src/net/url/url.go:866
		_go_fuzz_dep_.CoverTab[17603]++
//line /usr/local/go/src/net/url/url.go:866
		// _ = "end of CoverTab[17603]"
//line /usr/local/go/src/net/url/url.go:866
	}
//line /usr/local/go/src/net/url/url.go:866
	// _ = "end of CoverTab[17599]"
//line /usr/local/go/src/net/url/url.go:866
	_go_fuzz_dep_.CoverTab[17600]++

						ru := *u
						if _, has := ru.User.Password(); has {
//line /usr/local/go/src/net/url/url.go:869
		_go_fuzz_dep_.CoverTab[17604]++
							ru.User = UserPassword(ru.User.Username(), "xxxxx")
//line /usr/local/go/src/net/url/url.go:870
		// _ = "end of CoverTab[17604]"
	} else {
//line /usr/local/go/src/net/url/url.go:871
		_go_fuzz_dep_.CoverTab[17605]++
//line /usr/local/go/src/net/url/url.go:871
		// _ = "end of CoverTab[17605]"
//line /usr/local/go/src/net/url/url.go:871
	}
//line /usr/local/go/src/net/url/url.go:871
	// _ = "end of CoverTab[17600]"
//line /usr/local/go/src/net/url/url.go:871
	_go_fuzz_dep_.CoverTab[17601]++
						return ru.String()
//line /usr/local/go/src/net/url/url.go:872
	// _ = "end of CoverTab[17601]"
}

// Values maps a string key to a list of values.
//line /usr/local/go/src/net/url/url.go:875
// It is typically used for query parameters and form values.
//line /usr/local/go/src/net/url/url.go:875
// Unlike in the http.Header map, the keys in a Values map
//line /usr/local/go/src/net/url/url.go:875
// are case-sensitive.
//line /usr/local/go/src/net/url/url.go:879
type Values map[string][]string

// Get gets the first value associated with the given key.
//line /usr/local/go/src/net/url/url.go:881
// If there are no values associated with the key, Get returns
//line /usr/local/go/src/net/url/url.go:881
// the empty string. To access multiple values, use the map
//line /usr/local/go/src/net/url/url.go:881
// directly.
//line /usr/local/go/src/net/url/url.go:885
func (v Values) Get(key string) string {
//line /usr/local/go/src/net/url/url.go:885
	_go_fuzz_dep_.CoverTab[17606]++
						if v == nil {
//line /usr/local/go/src/net/url/url.go:886
		_go_fuzz_dep_.CoverTab[17609]++
							return ""
//line /usr/local/go/src/net/url/url.go:887
		// _ = "end of CoverTab[17609]"
	} else {
//line /usr/local/go/src/net/url/url.go:888
		_go_fuzz_dep_.CoverTab[17610]++
//line /usr/local/go/src/net/url/url.go:888
		// _ = "end of CoverTab[17610]"
//line /usr/local/go/src/net/url/url.go:888
	}
//line /usr/local/go/src/net/url/url.go:888
	// _ = "end of CoverTab[17606]"
//line /usr/local/go/src/net/url/url.go:888
	_go_fuzz_dep_.CoverTab[17607]++
						vs := v[key]
						if len(vs) == 0 {
//line /usr/local/go/src/net/url/url.go:890
		_go_fuzz_dep_.CoverTab[17611]++
							return ""
//line /usr/local/go/src/net/url/url.go:891
		// _ = "end of CoverTab[17611]"
	} else {
//line /usr/local/go/src/net/url/url.go:892
		_go_fuzz_dep_.CoverTab[17612]++
//line /usr/local/go/src/net/url/url.go:892
		// _ = "end of CoverTab[17612]"
//line /usr/local/go/src/net/url/url.go:892
	}
//line /usr/local/go/src/net/url/url.go:892
	// _ = "end of CoverTab[17607]"
//line /usr/local/go/src/net/url/url.go:892
	_go_fuzz_dep_.CoverTab[17608]++
						return vs[0]
//line /usr/local/go/src/net/url/url.go:893
	// _ = "end of CoverTab[17608]"
}

// Set sets the key to value. It replaces any existing
//line /usr/local/go/src/net/url/url.go:896
// values.
//line /usr/local/go/src/net/url/url.go:898
func (v Values) Set(key, value string) {
//line /usr/local/go/src/net/url/url.go:898
	_go_fuzz_dep_.CoverTab[17613]++
						v[key] = []string{value}
//line /usr/local/go/src/net/url/url.go:899
	// _ = "end of CoverTab[17613]"
}

// Add adds the value to key. It appends to any existing
//line /usr/local/go/src/net/url/url.go:902
// values associated with key.
//line /usr/local/go/src/net/url/url.go:904
func (v Values) Add(key, value string) {
//line /usr/local/go/src/net/url/url.go:904
	_go_fuzz_dep_.CoverTab[17614]++
						v[key] = append(v[key], value)
//line /usr/local/go/src/net/url/url.go:905
	// _ = "end of CoverTab[17614]"
}

// Del deletes the values associated with key.
func (v Values) Del(key string) {
//line /usr/local/go/src/net/url/url.go:909
	_go_fuzz_dep_.CoverTab[17615]++
						delete(v, key)
//line /usr/local/go/src/net/url/url.go:910
	// _ = "end of CoverTab[17615]"
}

// Has checks whether a given key is set.
func (v Values) Has(key string) bool {
//line /usr/local/go/src/net/url/url.go:914
	_go_fuzz_dep_.CoverTab[17616]++
						_, ok := v[key]
						return ok
//line /usr/local/go/src/net/url/url.go:916
	// _ = "end of CoverTab[17616]"
}

// ParseQuery parses the URL-encoded query string and returns
//line /usr/local/go/src/net/url/url.go:919
// a map listing the values specified for each key.
//line /usr/local/go/src/net/url/url.go:919
// ParseQuery always returns a non-nil map containing all the
//line /usr/local/go/src/net/url/url.go:919
// valid query parameters found; err describes the first decoding error
//line /usr/local/go/src/net/url/url.go:919
// encountered, if any.
//line /usr/local/go/src/net/url/url.go:919
//
//line /usr/local/go/src/net/url/url.go:919
// Query is expected to be a list of key=value settings separated by ampersands.
//line /usr/local/go/src/net/url/url.go:919
// A setting without an equals sign is interpreted as a key set to an empty
//line /usr/local/go/src/net/url/url.go:919
// value.
//line /usr/local/go/src/net/url/url.go:919
// Settings containing a non-URL-encoded semicolon are considered invalid.
//line /usr/local/go/src/net/url/url.go:929
func ParseQuery(query string) (Values, error) {
//line /usr/local/go/src/net/url/url.go:929
	_go_fuzz_dep_.CoverTab[17617]++
						m := make(Values)
						err := parseQuery(m, query)
						return m, err
//line /usr/local/go/src/net/url/url.go:932
	// _ = "end of CoverTab[17617]"
}

func parseQuery(m Values, query string) (err error) {
//line /usr/local/go/src/net/url/url.go:935
	_go_fuzz_dep_.CoverTab[17618]++
						for query != "" {
//line /usr/local/go/src/net/url/url.go:936
		_go_fuzz_dep_.CoverTab[17620]++
							var key string
							key, query, _ = strings.Cut(query, "&")
							if strings.Contains(key, ";") {
//line /usr/local/go/src/net/url/url.go:939
			_go_fuzz_dep_.CoverTab[17625]++
								err = fmt.Errorf("invalid semicolon separator in query")
								continue
//line /usr/local/go/src/net/url/url.go:941
			// _ = "end of CoverTab[17625]"
		} else {
//line /usr/local/go/src/net/url/url.go:942
			_go_fuzz_dep_.CoverTab[17626]++
//line /usr/local/go/src/net/url/url.go:942
			// _ = "end of CoverTab[17626]"
//line /usr/local/go/src/net/url/url.go:942
		}
//line /usr/local/go/src/net/url/url.go:942
		// _ = "end of CoverTab[17620]"
//line /usr/local/go/src/net/url/url.go:942
		_go_fuzz_dep_.CoverTab[17621]++
							if key == "" {
//line /usr/local/go/src/net/url/url.go:943
			_go_fuzz_dep_.CoverTab[17627]++
								continue
//line /usr/local/go/src/net/url/url.go:944
			// _ = "end of CoverTab[17627]"
		} else {
//line /usr/local/go/src/net/url/url.go:945
			_go_fuzz_dep_.CoverTab[17628]++
//line /usr/local/go/src/net/url/url.go:945
			// _ = "end of CoverTab[17628]"
//line /usr/local/go/src/net/url/url.go:945
		}
//line /usr/local/go/src/net/url/url.go:945
		// _ = "end of CoverTab[17621]"
//line /usr/local/go/src/net/url/url.go:945
		_go_fuzz_dep_.CoverTab[17622]++
							key, value, _ := strings.Cut(key, "=")
							key, err1 := QueryUnescape(key)
							if err1 != nil {
//line /usr/local/go/src/net/url/url.go:948
			_go_fuzz_dep_.CoverTab[17629]++
								if err == nil {
//line /usr/local/go/src/net/url/url.go:949
				_go_fuzz_dep_.CoverTab[17631]++
									err = err1
//line /usr/local/go/src/net/url/url.go:950
				// _ = "end of CoverTab[17631]"
			} else {
//line /usr/local/go/src/net/url/url.go:951
				_go_fuzz_dep_.CoverTab[17632]++
//line /usr/local/go/src/net/url/url.go:951
				// _ = "end of CoverTab[17632]"
//line /usr/local/go/src/net/url/url.go:951
			}
//line /usr/local/go/src/net/url/url.go:951
			// _ = "end of CoverTab[17629]"
//line /usr/local/go/src/net/url/url.go:951
			_go_fuzz_dep_.CoverTab[17630]++
								continue
//line /usr/local/go/src/net/url/url.go:952
			// _ = "end of CoverTab[17630]"
		} else {
//line /usr/local/go/src/net/url/url.go:953
			_go_fuzz_dep_.CoverTab[17633]++
//line /usr/local/go/src/net/url/url.go:953
			// _ = "end of CoverTab[17633]"
//line /usr/local/go/src/net/url/url.go:953
		}
//line /usr/local/go/src/net/url/url.go:953
		// _ = "end of CoverTab[17622]"
//line /usr/local/go/src/net/url/url.go:953
		_go_fuzz_dep_.CoverTab[17623]++
							value, err1 = QueryUnescape(value)
							if err1 != nil {
//line /usr/local/go/src/net/url/url.go:955
			_go_fuzz_dep_.CoverTab[17634]++
								if err == nil {
//line /usr/local/go/src/net/url/url.go:956
				_go_fuzz_dep_.CoverTab[17636]++
									err = err1
//line /usr/local/go/src/net/url/url.go:957
				// _ = "end of CoverTab[17636]"
			} else {
//line /usr/local/go/src/net/url/url.go:958
				_go_fuzz_dep_.CoverTab[17637]++
//line /usr/local/go/src/net/url/url.go:958
				// _ = "end of CoverTab[17637]"
//line /usr/local/go/src/net/url/url.go:958
			}
//line /usr/local/go/src/net/url/url.go:958
			// _ = "end of CoverTab[17634]"
//line /usr/local/go/src/net/url/url.go:958
			_go_fuzz_dep_.CoverTab[17635]++
								continue
//line /usr/local/go/src/net/url/url.go:959
			// _ = "end of CoverTab[17635]"
		} else {
//line /usr/local/go/src/net/url/url.go:960
			_go_fuzz_dep_.CoverTab[17638]++
//line /usr/local/go/src/net/url/url.go:960
			// _ = "end of CoverTab[17638]"
//line /usr/local/go/src/net/url/url.go:960
		}
//line /usr/local/go/src/net/url/url.go:960
		// _ = "end of CoverTab[17623]"
//line /usr/local/go/src/net/url/url.go:960
		_go_fuzz_dep_.CoverTab[17624]++
							m[key] = append(m[key], value)
//line /usr/local/go/src/net/url/url.go:961
		// _ = "end of CoverTab[17624]"
	}
//line /usr/local/go/src/net/url/url.go:962
	// _ = "end of CoverTab[17618]"
//line /usr/local/go/src/net/url/url.go:962
	_go_fuzz_dep_.CoverTab[17619]++
						return err
//line /usr/local/go/src/net/url/url.go:963
	// _ = "end of CoverTab[17619]"
}

// Encode encodes the values into “URL encoded” form
//line /usr/local/go/src/net/url/url.go:966
// ("bar=baz&foo=quux") sorted by key.
//line /usr/local/go/src/net/url/url.go:968
func (v Values) Encode() string {
//line /usr/local/go/src/net/url/url.go:968
	_go_fuzz_dep_.CoverTab[17639]++
						if v == nil {
//line /usr/local/go/src/net/url/url.go:969
		_go_fuzz_dep_.CoverTab[17643]++
							return ""
//line /usr/local/go/src/net/url/url.go:970
		// _ = "end of CoverTab[17643]"
	} else {
//line /usr/local/go/src/net/url/url.go:971
		_go_fuzz_dep_.CoverTab[17644]++
//line /usr/local/go/src/net/url/url.go:971
		// _ = "end of CoverTab[17644]"
//line /usr/local/go/src/net/url/url.go:971
	}
//line /usr/local/go/src/net/url/url.go:971
	// _ = "end of CoverTab[17639]"
//line /usr/local/go/src/net/url/url.go:971
	_go_fuzz_dep_.CoverTab[17640]++
						var buf strings.Builder
						keys := make([]string, 0, len(v))
						for k := range v {
//line /usr/local/go/src/net/url/url.go:974
		_go_fuzz_dep_.CoverTab[17645]++
							keys = append(keys, k)
//line /usr/local/go/src/net/url/url.go:975
		// _ = "end of CoverTab[17645]"
	}
//line /usr/local/go/src/net/url/url.go:976
	// _ = "end of CoverTab[17640]"
//line /usr/local/go/src/net/url/url.go:976
	_go_fuzz_dep_.CoverTab[17641]++
						sort.Strings(keys)
						for _, k := range keys {
//line /usr/local/go/src/net/url/url.go:978
		_go_fuzz_dep_.CoverTab[17646]++
							vs := v[k]
							keyEscaped := QueryEscape(k)
							for _, v := range vs {
//line /usr/local/go/src/net/url/url.go:981
			_go_fuzz_dep_.CoverTab[17647]++
								if buf.Len() > 0 {
//line /usr/local/go/src/net/url/url.go:982
				_go_fuzz_dep_.CoverTab[17649]++
									buf.WriteByte('&')
//line /usr/local/go/src/net/url/url.go:983
				// _ = "end of CoverTab[17649]"
			} else {
//line /usr/local/go/src/net/url/url.go:984
				_go_fuzz_dep_.CoverTab[17650]++
//line /usr/local/go/src/net/url/url.go:984
				// _ = "end of CoverTab[17650]"
//line /usr/local/go/src/net/url/url.go:984
			}
//line /usr/local/go/src/net/url/url.go:984
			// _ = "end of CoverTab[17647]"
//line /usr/local/go/src/net/url/url.go:984
			_go_fuzz_dep_.CoverTab[17648]++
								buf.WriteString(keyEscaped)
								buf.WriteByte('=')
								buf.WriteString(QueryEscape(v))
//line /usr/local/go/src/net/url/url.go:987
			// _ = "end of CoverTab[17648]"
		}
//line /usr/local/go/src/net/url/url.go:988
		// _ = "end of CoverTab[17646]"
	}
//line /usr/local/go/src/net/url/url.go:989
	// _ = "end of CoverTab[17641]"
//line /usr/local/go/src/net/url/url.go:989
	_go_fuzz_dep_.CoverTab[17642]++
						return buf.String()
//line /usr/local/go/src/net/url/url.go:990
	// _ = "end of CoverTab[17642]"
}

// resolvePath applies special path segments from refs and applies
//line /usr/local/go/src/net/url/url.go:993
// them to base, per RFC 3986.
//line /usr/local/go/src/net/url/url.go:995
func resolvePath(base, ref string) string {
//line /usr/local/go/src/net/url/url.go:995
	_go_fuzz_dep_.CoverTab[17651]++
						var full string
						if ref == "" {
//line /usr/local/go/src/net/url/url.go:997
		_go_fuzz_dep_.CoverTab[17657]++
							full = base
//line /usr/local/go/src/net/url/url.go:998
		// _ = "end of CoverTab[17657]"
	} else {
//line /usr/local/go/src/net/url/url.go:999
		_go_fuzz_dep_.CoverTab[17658]++
//line /usr/local/go/src/net/url/url.go:999
		if ref[0] != '/' {
//line /usr/local/go/src/net/url/url.go:999
			_go_fuzz_dep_.CoverTab[17659]++
								i := strings.LastIndex(base, "/")
								full = base[:i+1] + ref
//line /usr/local/go/src/net/url/url.go:1001
			// _ = "end of CoverTab[17659]"
		} else {
//line /usr/local/go/src/net/url/url.go:1002
			_go_fuzz_dep_.CoverTab[17660]++
								full = ref
//line /usr/local/go/src/net/url/url.go:1003
			// _ = "end of CoverTab[17660]"
		}
//line /usr/local/go/src/net/url/url.go:1004
		// _ = "end of CoverTab[17658]"
//line /usr/local/go/src/net/url/url.go:1004
	}
//line /usr/local/go/src/net/url/url.go:1004
	// _ = "end of CoverTab[17651]"
//line /usr/local/go/src/net/url/url.go:1004
	_go_fuzz_dep_.CoverTab[17652]++
						if full == "" {
//line /usr/local/go/src/net/url/url.go:1005
		_go_fuzz_dep_.CoverTab[17661]++
							return ""
//line /usr/local/go/src/net/url/url.go:1006
		// _ = "end of CoverTab[17661]"
	} else {
//line /usr/local/go/src/net/url/url.go:1007
		_go_fuzz_dep_.CoverTab[17662]++
//line /usr/local/go/src/net/url/url.go:1007
		// _ = "end of CoverTab[17662]"
//line /usr/local/go/src/net/url/url.go:1007
	}
//line /usr/local/go/src/net/url/url.go:1007
	// _ = "end of CoverTab[17652]"
//line /usr/local/go/src/net/url/url.go:1007
	_go_fuzz_dep_.CoverTab[17653]++

						var (
		elem	string
		dst	strings.Builder
	)
	first := true
	remaining := full

	dst.WriteByte('/')
	found := true
	for found {
//line /usr/local/go/src/net/url/url.go:1018
		_go_fuzz_dep_.CoverTab[17663]++
							elem, remaining, found = strings.Cut(remaining, "/")
							if elem == "." {
//line /usr/local/go/src/net/url/url.go:1020
			_go_fuzz_dep_.CoverTab[17665]++
								first = false

								continue
//line /usr/local/go/src/net/url/url.go:1023
			// _ = "end of CoverTab[17665]"
		} else {
//line /usr/local/go/src/net/url/url.go:1024
			_go_fuzz_dep_.CoverTab[17666]++
//line /usr/local/go/src/net/url/url.go:1024
			// _ = "end of CoverTab[17666]"
//line /usr/local/go/src/net/url/url.go:1024
		}
//line /usr/local/go/src/net/url/url.go:1024
		// _ = "end of CoverTab[17663]"
//line /usr/local/go/src/net/url/url.go:1024
		_go_fuzz_dep_.CoverTab[17664]++

							if elem == ".." {
//line /usr/local/go/src/net/url/url.go:1026
			_go_fuzz_dep_.CoverTab[17667]++

								str := dst.String()[1:]
								index := strings.LastIndexByte(str, '/')

								dst.Reset()
								dst.WriteByte('/')
								if index == -1 {
//line /usr/local/go/src/net/url/url.go:1033
				_go_fuzz_dep_.CoverTab[17668]++
									first = true
//line /usr/local/go/src/net/url/url.go:1034
				// _ = "end of CoverTab[17668]"
			} else {
//line /usr/local/go/src/net/url/url.go:1035
				_go_fuzz_dep_.CoverTab[17669]++
									dst.WriteString(str[:index])
//line /usr/local/go/src/net/url/url.go:1036
				// _ = "end of CoverTab[17669]"
			}
//line /usr/local/go/src/net/url/url.go:1037
			// _ = "end of CoverTab[17667]"
		} else {
//line /usr/local/go/src/net/url/url.go:1038
			_go_fuzz_dep_.CoverTab[17670]++
								if !first {
//line /usr/local/go/src/net/url/url.go:1039
				_go_fuzz_dep_.CoverTab[17672]++
									dst.WriteByte('/')
//line /usr/local/go/src/net/url/url.go:1040
				// _ = "end of CoverTab[17672]"
			} else {
//line /usr/local/go/src/net/url/url.go:1041
				_go_fuzz_dep_.CoverTab[17673]++
//line /usr/local/go/src/net/url/url.go:1041
				// _ = "end of CoverTab[17673]"
//line /usr/local/go/src/net/url/url.go:1041
			}
//line /usr/local/go/src/net/url/url.go:1041
			// _ = "end of CoverTab[17670]"
//line /usr/local/go/src/net/url/url.go:1041
			_go_fuzz_dep_.CoverTab[17671]++
								dst.WriteString(elem)
								first = false
//line /usr/local/go/src/net/url/url.go:1043
			// _ = "end of CoverTab[17671]"
		}
//line /usr/local/go/src/net/url/url.go:1044
		// _ = "end of CoverTab[17664]"
	}
//line /usr/local/go/src/net/url/url.go:1045
	// _ = "end of CoverTab[17653]"
//line /usr/local/go/src/net/url/url.go:1045
	_go_fuzz_dep_.CoverTab[17654]++

						if elem == "." || func() bool {
//line /usr/local/go/src/net/url/url.go:1047
		_go_fuzz_dep_.CoverTab[17674]++
//line /usr/local/go/src/net/url/url.go:1047
		return elem == ".."
//line /usr/local/go/src/net/url/url.go:1047
		// _ = "end of CoverTab[17674]"
//line /usr/local/go/src/net/url/url.go:1047
	}() {
//line /usr/local/go/src/net/url/url.go:1047
		_go_fuzz_dep_.CoverTab[17675]++
							dst.WriteByte('/')
//line /usr/local/go/src/net/url/url.go:1048
		// _ = "end of CoverTab[17675]"
	} else {
//line /usr/local/go/src/net/url/url.go:1049
		_go_fuzz_dep_.CoverTab[17676]++
//line /usr/local/go/src/net/url/url.go:1049
		// _ = "end of CoverTab[17676]"
//line /usr/local/go/src/net/url/url.go:1049
	}
//line /usr/local/go/src/net/url/url.go:1049
	// _ = "end of CoverTab[17654]"
//line /usr/local/go/src/net/url/url.go:1049
	_go_fuzz_dep_.CoverTab[17655]++

//line /usr/local/go/src/net/url/url.go:1052
	r := dst.String()
	if len(r) > 1 && func() bool {
//line /usr/local/go/src/net/url/url.go:1053
		_go_fuzz_dep_.CoverTab[17677]++
//line /usr/local/go/src/net/url/url.go:1053
		return r[1] == '/'
//line /usr/local/go/src/net/url/url.go:1053
		// _ = "end of CoverTab[17677]"
//line /usr/local/go/src/net/url/url.go:1053
	}() {
//line /usr/local/go/src/net/url/url.go:1053
		_go_fuzz_dep_.CoverTab[17678]++
							r = r[1:]
//line /usr/local/go/src/net/url/url.go:1054
		// _ = "end of CoverTab[17678]"
	} else {
//line /usr/local/go/src/net/url/url.go:1055
		_go_fuzz_dep_.CoverTab[17679]++
//line /usr/local/go/src/net/url/url.go:1055
		// _ = "end of CoverTab[17679]"
//line /usr/local/go/src/net/url/url.go:1055
	}
//line /usr/local/go/src/net/url/url.go:1055
	// _ = "end of CoverTab[17655]"
//line /usr/local/go/src/net/url/url.go:1055
	_go_fuzz_dep_.CoverTab[17656]++
						return r
//line /usr/local/go/src/net/url/url.go:1056
	// _ = "end of CoverTab[17656]"
}

// IsAbs reports whether the URL is absolute.
//line /usr/local/go/src/net/url/url.go:1059
// Absolute means that it has a non-empty scheme.
//line /usr/local/go/src/net/url/url.go:1061
func (u *URL) IsAbs() bool {
//line /usr/local/go/src/net/url/url.go:1061
	_go_fuzz_dep_.CoverTab[17680]++
						return u.Scheme != ""
//line /usr/local/go/src/net/url/url.go:1062
	// _ = "end of CoverTab[17680]"
}

// Parse parses a URL in the context of the receiver. The provided URL
//line /usr/local/go/src/net/url/url.go:1065
// may be relative or absolute. Parse returns nil, err on parse
//line /usr/local/go/src/net/url/url.go:1065
// failure, otherwise its return value is the same as ResolveReference.
//line /usr/local/go/src/net/url/url.go:1068
func (u *URL) Parse(ref string) (*URL, error) {
//line /usr/local/go/src/net/url/url.go:1068
	_go_fuzz_dep_.CoverTab[17681]++
						refURL, err := Parse(ref)
						if err != nil {
//line /usr/local/go/src/net/url/url.go:1070
		_go_fuzz_dep_.CoverTab[17683]++
							return nil, err
//line /usr/local/go/src/net/url/url.go:1071
		// _ = "end of CoverTab[17683]"
	} else {
//line /usr/local/go/src/net/url/url.go:1072
		_go_fuzz_dep_.CoverTab[17684]++
//line /usr/local/go/src/net/url/url.go:1072
		// _ = "end of CoverTab[17684]"
//line /usr/local/go/src/net/url/url.go:1072
	}
//line /usr/local/go/src/net/url/url.go:1072
	// _ = "end of CoverTab[17681]"
//line /usr/local/go/src/net/url/url.go:1072
	_go_fuzz_dep_.CoverTab[17682]++
						return u.ResolveReference(refURL), nil
//line /usr/local/go/src/net/url/url.go:1073
	// _ = "end of CoverTab[17682]"
}

// ResolveReference resolves a URI reference to an absolute URI from
//line /usr/local/go/src/net/url/url.go:1076
// an absolute base URI u, per RFC 3986 Section 5.2. The URI reference
//line /usr/local/go/src/net/url/url.go:1076
// may be relative or absolute. ResolveReference always returns a new
//line /usr/local/go/src/net/url/url.go:1076
// URL instance, even if the returned URL is identical to either the
//line /usr/local/go/src/net/url/url.go:1076
// base or reference. If ref is an absolute URL, then ResolveReference
//line /usr/local/go/src/net/url/url.go:1076
// ignores base and returns a copy of ref.
//line /usr/local/go/src/net/url/url.go:1082
func (u *URL) ResolveReference(ref *URL) *URL {
//line /usr/local/go/src/net/url/url.go:1082
	_go_fuzz_dep_.CoverTab[17685]++
						url := *ref
						if ref.Scheme == "" {
//line /usr/local/go/src/net/url/url.go:1084
		_go_fuzz_dep_.CoverTab[17690]++
							url.Scheme = u.Scheme
//line /usr/local/go/src/net/url/url.go:1085
		// _ = "end of CoverTab[17690]"
	} else {
//line /usr/local/go/src/net/url/url.go:1086
		_go_fuzz_dep_.CoverTab[17691]++
//line /usr/local/go/src/net/url/url.go:1086
		// _ = "end of CoverTab[17691]"
//line /usr/local/go/src/net/url/url.go:1086
	}
//line /usr/local/go/src/net/url/url.go:1086
	// _ = "end of CoverTab[17685]"
//line /usr/local/go/src/net/url/url.go:1086
	_go_fuzz_dep_.CoverTab[17686]++
						if ref.Scheme != "" || func() bool {
//line /usr/local/go/src/net/url/url.go:1087
		_go_fuzz_dep_.CoverTab[17692]++
//line /usr/local/go/src/net/url/url.go:1087
		return ref.Host != ""
//line /usr/local/go/src/net/url/url.go:1087
		// _ = "end of CoverTab[17692]"
//line /usr/local/go/src/net/url/url.go:1087
	}() || func() bool {
//line /usr/local/go/src/net/url/url.go:1087
		_go_fuzz_dep_.CoverTab[17693]++
//line /usr/local/go/src/net/url/url.go:1087
		return ref.User != nil
//line /usr/local/go/src/net/url/url.go:1087
		// _ = "end of CoverTab[17693]"
//line /usr/local/go/src/net/url/url.go:1087
	}() {
//line /usr/local/go/src/net/url/url.go:1087
		_go_fuzz_dep_.CoverTab[17694]++

//line /usr/local/go/src/net/url/url.go:1091
		url.setPath(resolvePath(ref.EscapedPath(), ""))
							return &url
//line /usr/local/go/src/net/url/url.go:1092
		// _ = "end of CoverTab[17694]"
	} else {
//line /usr/local/go/src/net/url/url.go:1093
		_go_fuzz_dep_.CoverTab[17695]++
//line /usr/local/go/src/net/url/url.go:1093
		// _ = "end of CoverTab[17695]"
//line /usr/local/go/src/net/url/url.go:1093
	}
//line /usr/local/go/src/net/url/url.go:1093
	// _ = "end of CoverTab[17686]"
//line /usr/local/go/src/net/url/url.go:1093
	_go_fuzz_dep_.CoverTab[17687]++
						if ref.Opaque != "" {
//line /usr/local/go/src/net/url/url.go:1094
		_go_fuzz_dep_.CoverTab[17696]++
							url.User = nil
							url.Host = ""
							url.Path = ""
							return &url
//line /usr/local/go/src/net/url/url.go:1098
		// _ = "end of CoverTab[17696]"
	} else {
//line /usr/local/go/src/net/url/url.go:1099
		_go_fuzz_dep_.CoverTab[17697]++
//line /usr/local/go/src/net/url/url.go:1099
		// _ = "end of CoverTab[17697]"
//line /usr/local/go/src/net/url/url.go:1099
	}
//line /usr/local/go/src/net/url/url.go:1099
	// _ = "end of CoverTab[17687]"
//line /usr/local/go/src/net/url/url.go:1099
	_go_fuzz_dep_.CoverTab[17688]++
						if ref.Path == "" && func() bool {
//line /usr/local/go/src/net/url/url.go:1100
		_go_fuzz_dep_.CoverTab[17698]++
//line /usr/local/go/src/net/url/url.go:1100
		return !ref.ForceQuery
//line /usr/local/go/src/net/url/url.go:1100
		// _ = "end of CoverTab[17698]"
//line /usr/local/go/src/net/url/url.go:1100
	}() && func() bool {
//line /usr/local/go/src/net/url/url.go:1100
		_go_fuzz_dep_.CoverTab[17699]++
//line /usr/local/go/src/net/url/url.go:1100
		return ref.RawQuery == ""
//line /usr/local/go/src/net/url/url.go:1100
		// _ = "end of CoverTab[17699]"
//line /usr/local/go/src/net/url/url.go:1100
	}() {
//line /usr/local/go/src/net/url/url.go:1100
		_go_fuzz_dep_.CoverTab[17700]++
							url.RawQuery = u.RawQuery
							if ref.Fragment == "" {
//line /usr/local/go/src/net/url/url.go:1102
			_go_fuzz_dep_.CoverTab[17701]++
								url.Fragment = u.Fragment
								url.RawFragment = u.RawFragment
//line /usr/local/go/src/net/url/url.go:1104
			// _ = "end of CoverTab[17701]"
		} else {
//line /usr/local/go/src/net/url/url.go:1105
			_go_fuzz_dep_.CoverTab[17702]++
//line /usr/local/go/src/net/url/url.go:1105
			// _ = "end of CoverTab[17702]"
//line /usr/local/go/src/net/url/url.go:1105
		}
//line /usr/local/go/src/net/url/url.go:1105
		// _ = "end of CoverTab[17700]"
	} else {
//line /usr/local/go/src/net/url/url.go:1106
		_go_fuzz_dep_.CoverTab[17703]++
//line /usr/local/go/src/net/url/url.go:1106
		// _ = "end of CoverTab[17703]"
//line /usr/local/go/src/net/url/url.go:1106
	}
//line /usr/local/go/src/net/url/url.go:1106
	// _ = "end of CoverTab[17688]"
//line /usr/local/go/src/net/url/url.go:1106
	_go_fuzz_dep_.CoverTab[17689]++

						url.Host = u.Host
						url.User = u.User
						url.setPath(resolvePath(u.EscapedPath(), ref.EscapedPath()))
						return &url
//line /usr/local/go/src/net/url/url.go:1111
	// _ = "end of CoverTab[17689]"
}

// Query parses RawQuery and returns the corresponding values.
//line /usr/local/go/src/net/url/url.go:1114
// It silently discards malformed value pairs.
//line /usr/local/go/src/net/url/url.go:1114
// To check errors use ParseQuery.
//line /usr/local/go/src/net/url/url.go:1117
func (u *URL) Query() Values {
//line /usr/local/go/src/net/url/url.go:1117
	_go_fuzz_dep_.CoverTab[17704]++
						v, _ := ParseQuery(u.RawQuery)
						return v
//line /usr/local/go/src/net/url/url.go:1119
	// _ = "end of CoverTab[17704]"
}

// RequestURI returns the encoded path?query or opaque?query
//line /usr/local/go/src/net/url/url.go:1122
// string that would be used in an HTTP request for u.
//line /usr/local/go/src/net/url/url.go:1124
func (u *URL) RequestURI() string {
//line /usr/local/go/src/net/url/url.go:1124
	_go_fuzz_dep_.CoverTab[17705]++
						result := u.Opaque
						if result == "" {
//line /usr/local/go/src/net/url/url.go:1126
		_go_fuzz_dep_.CoverTab[17708]++
							result = u.EscapedPath()
							if result == "" {
//line /usr/local/go/src/net/url/url.go:1128
			_go_fuzz_dep_.CoverTab[17709]++
								result = "/"
//line /usr/local/go/src/net/url/url.go:1129
			// _ = "end of CoverTab[17709]"
		} else {
//line /usr/local/go/src/net/url/url.go:1130
			_go_fuzz_dep_.CoverTab[17710]++
//line /usr/local/go/src/net/url/url.go:1130
			// _ = "end of CoverTab[17710]"
//line /usr/local/go/src/net/url/url.go:1130
		}
//line /usr/local/go/src/net/url/url.go:1130
		// _ = "end of CoverTab[17708]"
	} else {
//line /usr/local/go/src/net/url/url.go:1131
		_go_fuzz_dep_.CoverTab[17711]++
							if strings.HasPrefix(result, "//") {
//line /usr/local/go/src/net/url/url.go:1132
			_go_fuzz_dep_.CoverTab[17712]++
								result = u.Scheme + ":" + result
//line /usr/local/go/src/net/url/url.go:1133
			// _ = "end of CoverTab[17712]"
		} else {
//line /usr/local/go/src/net/url/url.go:1134
			_go_fuzz_dep_.CoverTab[17713]++
//line /usr/local/go/src/net/url/url.go:1134
			// _ = "end of CoverTab[17713]"
//line /usr/local/go/src/net/url/url.go:1134
		}
//line /usr/local/go/src/net/url/url.go:1134
		// _ = "end of CoverTab[17711]"
	}
//line /usr/local/go/src/net/url/url.go:1135
	// _ = "end of CoverTab[17705]"
//line /usr/local/go/src/net/url/url.go:1135
	_go_fuzz_dep_.CoverTab[17706]++
						if u.ForceQuery || func() bool {
//line /usr/local/go/src/net/url/url.go:1136
		_go_fuzz_dep_.CoverTab[17714]++
//line /usr/local/go/src/net/url/url.go:1136
		return u.RawQuery != ""
//line /usr/local/go/src/net/url/url.go:1136
		// _ = "end of CoverTab[17714]"
//line /usr/local/go/src/net/url/url.go:1136
	}() {
//line /usr/local/go/src/net/url/url.go:1136
		_go_fuzz_dep_.CoverTab[17715]++
							result += "?" + u.RawQuery
//line /usr/local/go/src/net/url/url.go:1137
		// _ = "end of CoverTab[17715]"
	} else {
//line /usr/local/go/src/net/url/url.go:1138
		_go_fuzz_dep_.CoverTab[17716]++
//line /usr/local/go/src/net/url/url.go:1138
		// _ = "end of CoverTab[17716]"
//line /usr/local/go/src/net/url/url.go:1138
	}
//line /usr/local/go/src/net/url/url.go:1138
	// _ = "end of CoverTab[17706]"
//line /usr/local/go/src/net/url/url.go:1138
	_go_fuzz_dep_.CoverTab[17707]++
						return result
//line /usr/local/go/src/net/url/url.go:1139
	// _ = "end of CoverTab[17707]"
}

// Hostname returns u.Host, stripping any valid port number if present.
//line /usr/local/go/src/net/url/url.go:1142
//
//line /usr/local/go/src/net/url/url.go:1142
// If the result is enclosed in square brackets, as literal IPv6 addresses are,
//line /usr/local/go/src/net/url/url.go:1142
// the square brackets are removed from the result.
//line /usr/local/go/src/net/url/url.go:1146
func (u *URL) Hostname() string {
//line /usr/local/go/src/net/url/url.go:1146
	_go_fuzz_dep_.CoverTab[17717]++
						host, _ := splitHostPort(u.Host)
						return host
//line /usr/local/go/src/net/url/url.go:1148
	// _ = "end of CoverTab[17717]"
}

// Port returns the port part of u.Host, without the leading colon.
//line /usr/local/go/src/net/url/url.go:1151
//
//line /usr/local/go/src/net/url/url.go:1151
// If u.Host doesn't contain a valid numeric port, Port returns an empty string.
//line /usr/local/go/src/net/url/url.go:1154
func (u *URL) Port() string {
//line /usr/local/go/src/net/url/url.go:1154
	_go_fuzz_dep_.CoverTab[17718]++
						_, port := splitHostPort(u.Host)
						return port
//line /usr/local/go/src/net/url/url.go:1156
	// _ = "end of CoverTab[17718]"
}

// splitHostPort separates host and port. If the port is not valid, it returns
//line /usr/local/go/src/net/url/url.go:1159
// the entire input as host, and it doesn't check the validity of the host.
//line /usr/local/go/src/net/url/url.go:1159
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
//line /usr/local/go/src/net/url/url.go:1162
func splitHostPort(hostPort string) (host, port string) {
//line /usr/local/go/src/net/url/url.go:1162
	_go_fuzz_dep_.CoverTab[17719]++
						host = hostPort

						colon := strings.LastIndexByte(host, ':')
						if colon != -1 && func() bool {
//line /usr/local/go/src/net/url/url.go:1166
		_go_fuzz_dep_.CoverTab[17722]++
//line /usr/local/go/src/net/url/url.go:1166
		return validOptionalPort(host[colon:])
//line /usr/local/go/src/net/url/url.go:1166
		// _ = "end of CoverTab[17722]"
//line /usr/local/go/src/net/url/url.go:1166
	}() {
//line /usr/local/go/src/net/url/url.go:1166
		_go_fuzz_dep_.CoverTab[17723]++
							host, port = host[:colon], host[colon+1:]
//line /usr/local/go/src/net/url/url.go:1167
		// _ = "end of CoverTab[17723]"
	} else {
//line /usr/local/go/src/net/url/url.go:1168
		_go_fuzz_dep_.CoverTab[17724]++
//line /usr/local/go/src/net/url/url.go:1168
		// _ = "end of CoverTab[17724]"
//line /usr/local/go/src/net/url/url.go:1168
	}
//line /usr/local/go/src/net/url/url.go:1168
	// _ = "end of CoverTab[17719]"
//line /usr/local/go/src/net/url/url.go:1168
	_go_fuzz_dep_.CoverTab[17720]++

						if strings.HasPrefix(host, "[") && func() bool {
//line /usr/local/go/src/net/url/url.go:1170
		_go_fuzz_dep_.CoverTab[17725]++
//line /usr/local/go/src/net/url/url.go:1170
		return strings.HasSuffix(host, "]")
//line /usr/local/go/src/net/url/url.go:1170
		// _ = "end of CoverTab[17725]"
//line /usr/local/go/src/net/url/url.go:1170
	}() {
//line /usr/local/go/src/net/url/url.go:1170
		_go_fuzz_dep_.CoverTab[17726]++
							host = host[1 : len(host)-1]
//line /usr/local/go/src/net/url/url.go:1171
		// _ = "end of CoverTab[17726]"
	} else {
//line /usr/local/go/src/net/url/url.go:1172
		_go_fuzz_dep_.CoverTab[17727]++
//line /usr/local/go/src/net/url/url.go:1172
		// _ = "end of CoverTab[17727]"
//line /usr/local/go/src/net/url/url.go:1172
	}
//line /usr/local/go/src/net/url/url.go:1172
	// _ = "end of CoverTab[17720]"
//line /usr/local/go/src/net/url/url.go:1172
	_go_fuzz_dep_.CoverTab[17721]++

						return
//line /usr/local/go/src/net/url/url.go:1174
	// _ = "end of CoverTab[17721]"
}

//line /usr/local/go/src/net/url/url.go:1180
func (u *URL) MarshalBinary() (text []byte, err error) {
//line /usr/local/go/src/net/url/url.go:1180
	_go_fuzz_dep_.CoverTab[17728]++
						return []byte(u.String()), nil
//line /usr/local/go/src/net/url/url.go:1181
	// _ = "end of CoverTab[17728]"
}

func (u *URL) UnmarshalBinary(text []byte) error {
//line /usr/local/go/src/net/url/url.go:1184
	_go_fuzz_dep_.CoverTab[17729]++
						u1, err := Parse(string(text))
						if err != nil {
//line /usr/local/go/src/net/url/url.go:1186
		_go_fuzz_dep_.CoverTab[17731]++
							return err
//line /usr/local/go/src/net/url/url.go:1187
		// _ = "end of CoverTab[17731]"
	} else {
//line /usr/local/go/src/net/url/url.go:1188
		_go_fuzz_dep_.CoverTab[17732]++
//line /usr/local/go/src/net/url/url.go:1188
		// _ = "end of CoverTab[17732]"
//line /usr/local/go/src/net/url/url.go:1188
	}
//line /usr/local/go/src/net/url/url.go:1188
	// _ = "end of CoverTab[17729]"
//line /usr/local/go/src/net/url/url.go:1188
	_go_fuzz_dep_.CoverTab[17730]++
						*u = *u1
						return nil
//line /usr/local/go/src/net/url/url.go:1190
	// _ = "end of CoverTab[17730]"
}

// JoinPath returns a new URL with the provided path elements joined to
//line /usr/local/go/src/net/url/url.go:1193
// any existing path and the resulting path cleaned of any ./ or ../ elements.
//line /usr/local/go/src/net/url/url.go:1193
// Any sequences of multiple / characters will be reduced to a single /.
//line /usr/local/go/src/net/url/url.go:1196
func (u *URL) JoinPath(elem ...string) *URL {
//line /usr/local/go/src/net/url/url.go:1196
	_go_fuzz_dep_.CoverTab[17733]++
						elem = append([]string{u.EscapedPath()}, elem...)
						var p string
						if !strings.HasPrefix(elem[0], "/") {
//line /usr/local/go/src/net/url/url.go:1199
		_go_fuzz_dep_.CoverTab[17736]++

//line /usr/local/go/src/net/url/url.go:1202
		elem[0] = "/" + elem[0]
							p = path.Join(elem...)[1:]
//line /usr/local/go/src/net/url/url.go:1203
		// _ = "end of CoverTab[17736]"
	} else {
//line /usr/local/go/src/net/url/url.go:1204
		_go_fuzz_dep_.CoverTab[17737]++
							p = path.Join(elem...)
//line /usr/local/go/src/net/url/url.go:1205
		// _ = "end of CoverTab[17737]"
	}
//line /usr/local/go/src/net/url/url.go:1206
	// _ = "end of CoverTab[17733]"
//line /usr/local/go/src/net/url/url.go:1206
	_go_fuzz_dep_.CoverTab[17734]++

//line /usr/local/go/src/net/url/url.go:1209
	if strings.HasSuffix(elem[len(elem)-1], "/") && func() bool {
//line /usr/local/go/src/net/url/url.go:1209
		_go_fuzz_dep_.CoverTab[17738]++
//line /usr/local/go/src/net/url/url.go:1209
		return !strings.HasSuffix(p, "/")
//line /usr/local/go/src/net/url/url.go:1209
		// _ = "end of CoverTab[17738]"
//line /usr/local/go/src/net/url/url.go:1209
	}() {
//line /usr/local/go/src/net/url/url.go:1209
		_go_fuzz_dep_.CoverTab[17739]++
							p += "/"
//line /usr/local/go/src/net/url/url.go:1210
		// _ = "end of CoverTab[17739]"
	} else {
//line /usr/local/go/src/net/url/url.go:1211
		_go_fuzz_dep_.CoverTab[17740]++
//line /usr/local/go/src/net/url/url.go:1211
		// _ = "end of CoverTab[17740]"
//line /usr/local/go/src/net/url/url.go:1211
	}
//line /usr/local/go/src/net/url/url.go:1211
	// _ = "end of CoverTab[17734]"
//line /usr/local/go/src/net/url/url.go:1211
	_go_fuzz_dep_.CoverTab[17735]++
						url := *u
						url.setPath(p)
						return &url
//line /usr/local/go/src/net/url/url.go:1214
	// _ = "end of CoverTab[17735]"
}

// validUserinfo reports whether s is a valid userinfo string per RFC 3986
//line /usr/local/go/src/net/url/url.go:1217
// Section 3.2.1:
//line /usr/local/go/src/net/url/url.go:1217
//
//line /usr/local/go/src/net/url/url.go:1217
//	userinfo    = *( unreserved / pct-encoded / sub-delims / ":" )
//line /usr/local/go/src/net/url/url.go:1217
//	unreserved  = ALPHA / DIGIT / "-" / "." / "_" / "~"
//line /usr/local/go/src/net/url/url.go:1217
//	sub-delims  = "!" / "$" / "&" / "'" / "(" / ")"
//line /usr/local/go/src/net/url/url.go:1217
//	              / "*" / "+" / "," / ";" / "="
//line /usr/local/go/src/net/url/url.go:1217
//
//line /usr/local/go/src/net/url/url.go:1217
// It doesn't validate pct-encoded. The caller does that via func unescape.
//line /usr/local/go/src/net/url/url.go:1226
func validUserinfo(s string) bool {
//line /usr/local/go/src/net/url/url.go:1226
	_go_fuzz_dep_.CoverTab[17741]++
						for _, r := range s {
//line /usr/local/go/src/net/url/url.go:1227
		_go_fuzz_dep_.CoverTab[17743]++
							if 'A' <= r && func() bool {
//line /usr/local/go/src/net/url/url.go:1228
			_go_fuzz_dep_.CoverTab[17747]++
//line /usr/local/go/src/net/url/url.go:1228
			return r <= 'Z'
//line /usr/local/go/src/net/url/url.go:1228
			// _ = "end of CoverTab[17747]"
//line /usr/local/go/src/net/url/url.go:1228
		}() {
//line /usr/local/go/src/net/url/url.go:1228
			_go_fuzz_dep_.CoverTab[17748]++
								continue
//line /usr/local/go/src/net/url/url.go:1229
			// _ = "end of CoverTab[17748]"
		} else {
//line /usr/local/go/src/net/url/url.go:1230
			_go_fuzz_dep_.CoverTab[17749]++
//line /usr/local/go/src/net/url/url.go:1230
			// _ = "end of CoverTab[17749]"
//line /usr/local/go/src/net/url/url.go:1230
		}
//line /usr/local/go/src/net/url/url.go:1230
		// _ = "end of CoverTab[17743]"
//line /usr/local/go/src/net/url/url.go:1230
		_go_fuzz_dep_.CoverTab[17744]++
							if 'a' <= r && func() bool {
//line /usr/local/go/src/net/url/url.go:1231
			_go_fuzz_dep_.CoverTab[17750]++
//line /usr/local/go/src/net/url/url.go:1231
			return r <= 'z'
//line /usr/local/go/src/net/url/url.go:1231
			// _ = "end of CoverTab[17750]"
//line /usr/local/go/src/net/url/url.go:1231
		}() {
//line /usr/local/go/src/net/url/url.go:1231
			_go_fuzz_dep_.CoverTab[17751]++
								continue
//line /usr/local/go/src/net/url/url.go:1232
			// _ = "end of CoverTab[17751]"
		} else {
//line /usr/local/go/src/net/url/url.go:1233
			_go_fuzz_dep_.CoverTab[17752]++
//line /usr/local/go/src/net/url/url.go:1233
			// _ = "end of CoverTab[17752]"
//line /usr/local/go/src/net/url/url.go:1233
		}
//line /usr/local/go/src/net/url/url.go:1233
		// _ = "end of CoverTab[17744]"
//line /usr/local/go/src/net/url/url.go:1233
		_go_fuzz_dep_.CoverTab[17745]++
							if '0' <= r && func() bool {
//line /usr/local/go/src/net/url/url.go:1234
			_go_fuzz_dep_.CoverTab[17753]++
//line /usr/local/go/src/net/url/url.go:1234
			return r <= '9'
//line /usr/local/go/src/net/url/url.go:1234
			// _ = "end of CoverTab[17753]"
//line /usr/local/go/src/net/url/url.go:1234
		}() {
//line /usr/local/go/src/net/url/url.go:1234
			_go_fuzz_dep_.CoverTab[17754]++
								continue
//line /usr/local/go/src/net/url/url.go:1235
			// _ = "end of CoverTab[17754]"
		} else {
//line /usr/local/go/src/net/url/url.go:1236
			_go_fuzz_dep_.CoverTab[17755]++
//line /usr/local/go/src/net/url/url.go:1236
			// _ = "end of CoverTab[17755]"
//line /usr/local/go/src/net/url/url.go:1236
		}
//line /usr/local/go/src/net/url/url.go:1236
		// _ = "end of CoverTab[17745]"
//line /usr/local/go/src/net/url/url.go:1236
		_go_fuzz_dep_.CoverTab[17746]++
							switch r {
		case '-', '.', '_', ':', '~', '!', '$', '&', '\'',
			'(', ')', '*', '+', ',', ';', '=', '%', '@':
//line /usr/local/go/src/net/url/url.go:1239
			_go_fuzz_dep_.CoverTab[17756]++
								continue
//line /usr/local/go/src/net/url/url.go:1240
			// _ = "end of CoverTab[17756]"
		default:
//line /usr/local/go/src/net/url/url.go:1241
			_go_fuzz_dep_.CoverTab[17757]++
								return false
//line /usr/local/go/src/net/url/url.go:1242
			// _ = "end of CoverTab[17757]"
		}
//line /usr/local/go/src/net/url/url.go:1243
		// _ = "end of CoverTab[17746]"
	}
//line /usr/local/go/src/net/url/url.go:1244
	// _ = "end of CoverTab[17741]"
//line /usr/local/go/src/net/url/url.go:1244
	_go_fuzz_dep_.CoverTab[17742]++
						return true
//line /usr/local/go/src/net/url/url.go:1245
	// _ = "end of CoverTab[17742]"
}

// stringContainsCTLByte reports whether s contains any ASCII control character.
func stringContainsCTLByte(s string) bool {
//line /usr/local/go/src/net/url/url.go:1249
	_go_fuzz_dep_.CoverTab[17758]++
						for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/url/url.go:1250
		_go_fuzz_dep_.CoverTab[17760]++
							b := s[i]
							if b < ' ' || func() bool {
//line /usr/local/go/src/net/url/url.go:1252
			_go_fuzz_dep_.CoverTab[17761]++
//line /usr/local/go/src/net/url/url.go:1252
			return b == 0x7f
//line /usr/local/go/src/net/url/url.go:1252
			// _ = "end of CoverTab[17761]"
//line /usr/local/go/src/net/url/url.go:1252
		}() {
//line /usr/local/go/src/net/url/url.go:1252
			_go_fuzz_dep_.CoverTab[17762]++
								return true
//line /usr/local/go/src/net/url/url.go:1253
			// _ = "end of CoverTab[17762]"
		} else {
//line /usr/local/go/src/net/url/url.go:1254
			_go_fuzz_dep_.CoverTab[17763]++
//line /usr/local/go/src/net/url/url.go:1254
			// _ = "end of CoverTab[17763]"
//line /usr/local/go/src/net/url/url.go:1254
		}
//line /usr/local/go/src/net/url/url.go:1254
		// _ = "end of CoverTab[17760]"
	}
//line /usr/local/go/src/net/url/url.go:1255
	// _ = "end of CoverTab[17758]"
//line /usr/local/go/src/net/url/url.go:1255
	_go_fuzz_dep_.CoverTab[17759]++
						return false
//line /usr/local/go/src/net/url/url.go:1256
	// _ = "end of CoverTab[17759]"
}

// JoinPath returns a URL string with the provided path elements joined to
//line /usr/local/go/src/net/url/url.go:1259
// the existing path of base and the resulting path cleaned of any ./ or ../ elements.
//line /usr/local/go/src/net/url/url.go:1261
func JoinPath(base string, elem ...string) (result string, err error) {
//line /usr/local/go/src/net/url/url.go:1261
	_go_fuzz_dep_.CoverTab[17764]++
						url, err := Parse(base)
						if err != nil {
//line /usr/local/go/src/net/url/url.go:1263
		_go_fuzz_dep_.CoverTab[17766]++
							return
//line /usr/local/go/src/net/url/url.go:1264
		// _ = "end of CoverTab[17766]"
	} else {
//line /usr/local/go/src/net/url/url.go:1265
		_go_fuzz_dep_.CoverTab[17767]++
//line /usr/local/go/src/net/url/url.go:1265
		// _ = "end of CoverTab[17767]"
//line /usr/local/go/src/net/url/url.go:1265
	}
//line /usr/local/go/src/net/url/url.go:1265
	// _ = "end of CoverTab[17764]"
//line /usr/local/go/src/net/url/url.go:1265
	_go_fuzz_dep_.CoverTab[17765]++
						result = url.JoinPath(elem...).String()
						return
//line /usr/local/go/src/net/url/url.go:1267
	// _ = "end of CoverTab[17765]"
}

//line /usr/local/go/src/net/url/url.go:1268
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/url/url.go:1268
var _ = _go_fuzz_dep_.CoverTab
