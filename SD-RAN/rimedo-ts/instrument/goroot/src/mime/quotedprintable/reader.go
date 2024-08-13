// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/mime/quotedprintable/reader.go:5
// Package quotedprintable implements quoted-printable encoding as specified by
//line /usr/local/go/src/mime/quotedprintable/reader.go:5
// RFC 2045.
//line /usr/local/go/src/mime/quotedprintable/reader.go:7
package quotedprintable

//line /usr/local/go/src/mime/quotedprintable/reader.go:7
import (
//line /usr/local/go/src/mime/quotedprintable/reader.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/mime/quotedprintable/reader.go:7
)
//line /usr/local/go/src/mime/quotedprintable/reader.go:7
import (
//line /usr/local/go/src/mime/quotedprintable/reader.go:7
	_atomic_ "sync/atomic"
//line /usr/local/go/src/mime/quotedprintable/reader.go:7
)

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Reader is a quoted-printable decoder.
type Reader struct {
	br	*bufio.Reader
	rerr	error	// last read error
	line	[]byte	// to be consumed before more of br
}

// NewReader returns a quoted-printable reader, decoding from r.
func NewReader(r io.Reader) *Reader {
//line /usr/local/go/src/mime/quotedprintable/reader.go:24
	_go_fuzz_dep_.CoverTab[35912]++
								return &Reader{
		br: bufio.NewReader(r),
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:27
	// _ = "end of CoverTab[35912]"
}

func fromHex(b byte) (byte, error) {
//line /usr/local/go/src/mime/quotedprintable/reader.go:30
	_go_fuzz_dep_.CoverTab[35913]++
								switch {
	case b >= '0' && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:32
		_go_fuzz_dep_.CoverTab[35919]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:32
		return b <= '9'
//line /usr/local/go/src/mime/quotedprintable/reader.go:32
		// _ = "end of CoverTab[35919]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:32
	}():
//line /usr/local/go/src/mime/quotedprintable/reader.go:32
		_go_fuzz_dep_.CoverTab[35915]++
									return b - '0', nil
//line /usr/local/go/src/mime/quotedprintable/reader.go:33
		// _ = "end of CoverTab[35915]"
	case b >= 'A' && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:34
		_go_fuzz_dep_.CoverTab[35920]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:34
		return b <= 'F'
//line /usr/local/go/src/mime/quotedprintable/reader.go:34
		// _ = "end of CoverTab[35920]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:34
	}():
//line /usr/local/go/src/mime/quotedprintable/reader.go:34
		_go_fuzz_dep_.CoverTab[35916]++
									return b - 'A' + 10, nil
//line /usr/local/go/src/mime/quotedprintable/reader.go:35
		// _ = "end of CoverTab[35916]"

	case b >= 'a' && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:37
		_go_fuzz_dep_.CoverTab[35921]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:37
		return b <= 'f'
//line /usr/local/go/src/mime/quotedprintable/reader.go:37
		// _ = "end of CoverTab[35921]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:37
	}():
//line /usr/local/go/src/mime/quotedprintable/reader.go:37
		_go_fuzz_dep_.CoverTab[35917]++
									return b - 'a' + 10, nil
//line /usr/local/go/src/mime/quotedprintable/reader.go:38
		// _ = "end of CoverTab[35917]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:38
	default:
//line /usr/local/go/src/mime/quotedprintable/reader.go:38
		_go_fuzz_dep_.CoverTab[35918]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:38
		// _ = "end of CoverTab[35918]"
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:39
	// _ = "end of CoverTab[35913]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:39
	_go_fuzz_dep_.CoverTab[35914]++
								return 0, fmt.Errorf("quotedprintable: invalid hex byte 0x%02x", b)
//line /usr/local/go/src/mime/quotedprintable/reader.go:40
	// _ = "end of CoverTab[35914]"
}

func readHexByte(v []byte) (b byte, err error) {
//line /usr/local/go/src/mime/quotedprintable/reader.go:43
	_go_fuzz_dep_.CoverTab[35922]++
								if len(v) < 2 {
//line /usr/local/go/src/mime/quotedprintable/reader.go:44
		_go_fuzz_dep_.CoverTab[35926]++
									return 0, io.ErrUnexpectedEOF
//line /usr/local/go/src/mime/quotedprintable/reader.go:45
		// _ = "end of CoverTab[35926]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:46
		_go_fuzz_dep_.CoverTab[35927]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:46
		// _ = "end of CoverTab[35927]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:46
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:46
	// _ = "end of CoverTab[35922]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:46
	_go_fuzz_dep_.CoverTab[35923]++
								var hb, lb byte
								if hb, err = fromHex(v[0]); err != nil {
//line /usr/local/go/src/mime/quotedprintable/reader.go:48
		_go_fuzz_dep_.CoverTab[35928]++
									return 0, err
//line /usr/local/go/src/mime/quotedprintable/reader.go:49
		// _ = "end of CoverTab[35928]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:50
		_go_fuzz_dep_.CoverTab[35929]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:50
		// _ = "end of CoverTab[35929]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:50
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:50
	// _ = "end of CoverTab[35923]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:50
	_go_fuzz_dep_.CoverTab[35924]++
								if lb, err = fromHex(v[1]); err != nil {
//line /usr/local/go/src/mime/quotedprintable/reader.go:51
		_go_fuzz_dep_.CoverTab[35930]++
									return 0, err
//line /usr/local/go/src/mime/quotedprintable/reader.go:52
		// _ = "end of CoverTab[35930]"
	} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:53
		_go_fuzz_dep_.CoverTab[35931]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:53
		// _ = "end of CoverTab[35931]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:53
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:53
	// _ = "end of CoverTab[35924]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:53
	_go_fuzz_dep_.CoverTab[35925]++
								return hb<<4 | lb, nil
//line /usr/local/go/src/mime/quotedprintable/reader.go:54
	// _ = "end of CoverTab[35925]"
}

func isQPDiscardWhitespace(r rune) bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:57
	_go_fuzz_dep_.CoverTab[35932]++
								switch r {
	case '\n', '\r', ' ', '\t':
//line /usr/local/go/src/mime/quotedprintable/reader.go:59
		_go_fuzz_dep_.CoverTab[35934]++
									return true
//line /usr/local/go/src/mime/quotedprintable/reader.go:60
		// _ = "end of CoverTab[35934]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:60
	default:
//line /usr/local/go/src/mime/quotedprintable/reader.go:60
		_go_fuzz_dep_.CoverTab[35935]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:60
		// _ = "end of CoverTab[35935]"
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:61
	// _ = "end of CoverTab[35932]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:61
	_go_fuzz_dep_.CoverTab[35933]++
								return false
//line /usr/local/go/src/mime/quotedprintable/reader.go:62
	// _ = "end of CoverTab[35933]"
}

var (
	crlf		= []byte("\r\n")
	lf		= []byte("\n")
	softSuffix	= []byte("=")
)

// Read reads and decodes quoted-printable data from the underlying reader.
func (r *Reader) Read(p []byte) (n int, err error) {
//line /usr/local/go/src/mime/quotedprintable/reader.go:72
	_go_fuzz_dep_.CoverTab[35936]++

//line /usr/local/go/src/mime/quotedprintable/reader.go:82
	for len(p) > 0 {
//line /usr/local/go/src/mime/quotedprintable/reader.go:82
		_go_fuzz_dep_.CoverTab[35938]++
									if len(r.line) == 0 {
//line /usr/local/go/src/mime/quotedprintable/reader.go:83
			_go_fuzz_dep_.CoverTab[35941]++
										if r.rerr != nil {
//line /usr/local/go/src/mime/quotedprintable/reader.go:84
				_go_fuzz_dep_.CoverTab[35944]++
											return n, r.rerr
//line /usr/local/go/src/mime/quotedprintable/reader.go:85
				// _ = "end of CoverTab[35944]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:86
				_go_fuzz_dep_.CoverTab[35945]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:86
				// _ = "end of CoverTab[35945]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:86
			}
//line /usr/local/go/src/mime/quotedprintable/reader.go:86
			// _ = "end of CoverTab[35941]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:86
			_go_fuzz_dep_.CoverTab[35942]++
										r.line, r.rerr = r.br.ReadSlice('\n')

//line /usr/local/go/src/mime/quotedprintable/reader.go:90
			hasLF := bytes.HasSuffix(r.line, lf)
			hasCR := bytes.HasSuffix(r.line, crlf)
			wholeLine := r.line
			r.line = bytes.TrimRightFunc(wholeLine, isQPDiscardWhitespace)
			if bytes.HasSuffix(r.line, softSuffix) {
//line /usr/local/go/src/mime/quotedprintable/reader.go:94
				_go_fuzz_dep_.CoverTab[35946]++
											rightStripped := wholeLine[len(r.line):]
											r.line = r.line[:len(r.line)-1]
											if !bytes.HasPrefix(rightStripped, lf) && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:97
					_go_fuzz_dep_.CoverTab[35947]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:97
					return !bytes.HasPrefix(rightStripped, crlf)
//line /usr/local/go/src/mime/quotedprintable/reader.go:97
					// _ = "end of CoverTab[35947]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:97
				}() && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:97
					_go_fuzz_dep_.CoverTab[35948]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:97
					return !(len(rightStripped) == 0 && func() bool {
													_go_fuzz_dep_.CoverTab[35949]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
						return len(r.line) > 0
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
						// _ = "end of CoverTab[35949]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
					}() && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
						_go_fuzz_dep_.CoverTab[35950]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
						return r.rerr == io.EOF
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
						// _ = "end of CoverTab[35950]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
					}())
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
					// _ = "end of CoverTab[35948]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
				}() {
//line /usr/local/go/src/mime/quotedprintable/reader.go:98
					_go_fuzz_dep_.CoverTab[35951]++
												r.rerr = fmt.Errorf("quotedprintable: invalid bytes after =: %q", rightStripped)
//line /usr/local/go/src/mime/quotedprintable/reader.go:99
					// _ = "end of CoverTab[35951]"
				} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:100
					_go_fuzz_dep_.CoverTab[35952]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:100
					// _ = "end of CoverTab[35952]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:100
				}
//line /usr/local/go/src/mime/quotedprintable/reader.go:100
				// _ = "end of CoverTab[35946]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:101
				_go_fuzz_dep_.CoverTab[35953]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:101
				if hasLF {
//line /usr/local/go/src/mime/quotedprintable/reader.go:101
					_go_fuzz_dep_.CoverTab[35954]++
												if hasCR {
//line /usr/local/go/src/mime/quotedprintable/reader.go:102
						_go_fuzz_dep_.CoverTab[35955]++
													r.line = append(r.line, '\r', '\n')
//line /usr/local/go/src/mime/quotedprintable/reader.go:103
						// _ = "end of CoverTab[35955]"
					} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:104
						_go_fuzz_dep_.CoverTab[35956]++
													r.line = append(r.line, '\n')
//line /usr/local/go/src/mime/quotedprintable/reader.go:105
						// _ = "end of CoverTab[35956]"
					}
//line /usr/local/go/src/mime/quotedprintable/reader.go:106
					// _ = "end of CoverTab[35954]"
				} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
					_go_fuzz_dep_.CoverTab[35957]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
					// _ = "end of CoverTab[35957]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
				}
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
				// _ = "end of CoverTab[35953]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
			}
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
			// _ = "end of CoverTab[35942]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:107
			_go_fuzz_dep_.CoverTab[35943]++
										continue
//line /usr/local/go/src/mime/quotedprintable/reader.go:108
			// _ = "end of CoverTab[35943]"
		} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:109
			_go_fuzz_dep_.CoverTab[35958]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:109
			// _ = "end of CoverTab[35958]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:109
		}
//line /usr/local/go/src/mime/quotedprintable/reader.go:109
		// _ = "end of CoverTab[35938]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:109
		_go_fuzz_dep_.CoverTab[35939]++
									b := r.line[0]

									switch {
		case b == '=':
//line /usr/local/go/src/mime/quotedprintable/reader.go:113
			_go_fuzz_dep_.CoverTab[35959]++
										b, err = readHexByte(r.line[1:])
										if err != nil {
//line /usr/local/go/src/mime/quotedprintable/reader.go:115
				_go_fuzz_dep_.CoverTab[35965]++
											if len(r.line) >= 2 && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					_go_fuzz_dep_.CoverTab[35967]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					return r.line[1] != '\r'
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					// _ = "end of CoverTab[35967]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
				}() && func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					_go_fuzz_dep_.CoverTab[35968]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					return r.line[1] != '\n'
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					// _ = "end of CoverTab[35968]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
				}() {
//line /usr/local/go/src/mime/quotedprintable/reader.go:116
					_go_fuzz_dep_.CoverTab[35969]++

												b = '='
												break
//line /usr/local/go/src/mime/quotedprintable/reader.go:119
					// _ = "end of CoverTab[35969]"
				} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:120
					_go_fuzz_dep_.CoverTab[35970]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:120
					// _ = "end of CoverTab[35970]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:120
				}
//line /usr/local/go/src/mime/quotedprintable/reader.go:120
				// _ = "end of CoverTab[35965]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:120
				_go_fuzz_dep_.CoverTab[35966]++
											return n, err
//line /usr/local/go/src/mime/quotedprintable/reader.go:121
				// _ = "end of CoverTab[35966]"
			} else {
//line /usr/local/go/src/mime/quotedprintable/reader.go:122
				_go_fuzz_dep_.CoverTab[35971]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:122
				// _ = "end of CoverTab[35971]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:122
			}
//line /usr/local/go/src/mime/quotedprintable/reader.go:122
			// _ = "end of CoverTab[35959]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:122
			_go_fuzz_dep_.CoverTab[35960]++
										r.line = r.line[2:]
//line /usr/local/go/src/mime/quotedprintable/reader.go:123
			// _ = "end of CoverTab[35960]"
		case b == '\t' || func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			_go_fuzz_dep_.CoverTab[35972]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			return b == '\r'
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			// _ = "end of CoverTab[35972]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
		}() || func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			_go_fuzz_dep_.CoverTab[35973]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			return b == '\n'
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			// _ = "end of CoverTab[35973]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
		}():
//line /usr/local/go/src/mime/quotedprintable/reader.go:124
			_go_fuzz_dep_.CoverTab[35961]++
										break
//line /usr/local/go/src/mime/quotedprintable/reader.go:125
			// _ = "end of CoverTab[35961]"
		case b >= 0x80:
//line /usr/local/go/src/mime/quotedprintable/reader.go:126
			_go_fuzz_dep_.CoverTab[35962]++

//line /usr/local/go/src/mime/quotedprintable/reader.go:129
			break
//line /usr/local/go/src/mime/quotedprintable/reader.go:129
			// _ = "end of CoverTab[35962]"
		case b < ' ' || func() bool {
//line /usr/local/go/src/mime/quotedprintable/reader.go:130
			_go_fuzz_dep_.CoverTab[35974]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:130
			return b > '~'
//line /usr/local/go/src/mime/quotedprintable/reader.go:130
			// _ = "end of CoverTab[35974]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:130
		}():
//line /usr/local/go/src/mime/quotedprintable/reader.go:130
			_go_fuzz_dep_.CoverTab[35963]++
										return n, fmt.Errorf("quotedprintable: invalid unescaped byte 0x%02x in body", b)
//line /usr/local/go/src/mime/quotedprintable/reader.go:131
			// _ = "end of CoverTab[35963]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:131
		default:
//line /usr/local/go/src/mime/quotedprintable/reader.go:131
			_go_fuzz_dep_.CoverTab[35964]++
//line /usr/local/go/src/mime/quotedprintable/reader.go:131
			// _ = "end of CoverTab[35964]"
		}
//line /usr/local/go/src/mime/quotedprintable/reader.go:132
		// _ = "end of CoverTab[35939]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:132
		_go_fuzz_dep_.CoverTab[35940]++
									p[0] = b
									p = p[1:]
									r.line = r.line[1:]
									n++
//line /usr/local/go/src/mime/quotedprintable/reader.go:136
		// _ = "end of CoverTab[35940]"
	}
//line /usr/local/go/src/mime/quotedprintable/reader.go:137
	// _ = "end of CoverTab[35936]"
//line /usr/local/go/src/mime/quotedprintable/reader.go:137
	_go_fuzz_dep_.CoverTab[35937]++
								return n, nil
//line /usr/local/go/src/mime/quotedprintable/reader.go:138
	// _ = "end of CoverTab[35937]"
}

//line /usr/local/go/src/mime/quotedprintable/reader.go:139
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/mime/quotedprintable/reader.go:139
var _ = _go_fuzz_dep_.CoverTab
