// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/net/textproto/reader.go:5
package textproto

//line /usr/local/go/src/net/textproto/reader.go:5
import (
//line /usr/local/go/src/net/textproto/reader.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/net/textproto/reader.go:5
)
//line /usr/local/go/src/net/textproto/reader.go:5
import (
//line /usr/local/go/src/net/textproto/reader.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/net/textproto/reader.go:5
)

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
)

// A Reader implements convenience methods for reading requests
//line /usr/local/go/src/net/textproto/reader.go:19
// or responses from a text protocol network connection.
//line /usr/local/go/src/net/textproto/reader.go:21
type Reader struct {
	R	*bufio.Reader
	dot	*dotReader
	buf	[]byte	// a re-usable buffer for readContinuedLineSlice
}

// NewReader returns a new Reader reading from r.
//line /usr/local/go/src/net/textproto/reader.go:27
//
//line /usr/local/go/src/net/textproto/reader.go:27
// To avoid denial of service attacks, the provided bufio.Reader
//line /usr/local/go/src/net/textproto/reader.go:27
// should be reading from an io.LimitReader or similar Reader to bound
//line /usr/local/go/src/net/textproto/reader.go:27
// the size of responses.
//line /usr/local/go/src/net/textproto/reader.go:32
func NewReader(r *bufio.Reader) *Reader {
//line /usr/local/go/src/net/textproto/reader.go:32
	_go_fuzz_dep_.CoverTab[34541]++
							return &Reader{R: r}
//line /usr/local/go/src/net/textproto/reader.go:33
	// _ = "end of CoverTab[34541]"
}

// ReadLine reads a single line from r,
//line /usr/local/go/src/net/textproto/reader.go:36
// eliding the final \n or \r\n from the returned string.
//line /usr/local/go/src/net/textproto/reader.go:38
func (r *Reader) ReadLine() (string, error) {
//line /usr/local/go/src/net/textproto/reader.go:38
	_go_fuzz_dep_.CoverTab[34542]++
							line, err := r.readLineSlice()
							return string(line), err
//line /usr/local/go/src/net/textproto/reader.go:40
	// _ = "end of CoverTab[34542]"
}

// ReadLineBytes is like ReadLine but returns a []byte instead of a string.
func (r *Reader) ReadLineBytes() ([]byte, error) {
//line /usr/local/go/src/net/textproto/reader.go:44
	_go_fuzz_dep_.CoverTab[34543]++
							line, err := r.readLineSlice()
							if line != nil {
//line /usr/local/go/src/net/textproto/reader.go:46
		_go_fuzz_dep_.CoverTab[34545]++
								line = bytes.Clone(line)
//line /usr/local/go/src/net/textproto/reader.go:47
		// _ = "end of CoverTab[34545]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:48
		_go_fuzz_dep_.CoverTab[34546]++
//line /usr/local/go/src/net/textproto/reader.go:48
		// _ = "end of CoverTab[34546]"
//line /usr/local/go/src/net/textproto/reader.go:48
	}
//line /usr/local/go/src/net/textproto/reader.go:48
	// _ = "end of CoverTab[34543]"
//line /usr/local/go/src/net/textproto/reader.go:48
	_go_fuzz_dep_.CoverTab[34544]++
							return line, err
//line /usr/local/go/src/net/textproto/reader.go:49
	// _ = "end of CoverTab[34544]"
}

func (r *Reader) readLineSlice() ([]byte, error) {
//line /usr/local/go/src/net/textproto/reader.go:52
	_go_fuzz_dep_.CoverTab[34547]++
							r.closeDot()
							var line []byte
							for {
//line /usr/local/go/src/net/textproto/reader.go:55
		_go_fuzz_dep_.CoverTab[34549]++
								l, more, err := r.R.ReadLine()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:57
			_go_fuzz_dep_.CoverTab[34552]++
									return nil, err
//line /usr/local/go/src/net/textproto/reader.go:58
			// _ = "end of CoverTab[34552]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:59
			_go_fuzz_dep_.CoverTab[34553]++
//line /usr/local/go/src/net/textproto/reader.go:59
			// _ = "end of CoverTab[34553]"
//line /usr/local/go/src/net/textproto/reader.go:59
		}
//line /usr/local/go/src/net/textproto/reader.go:59
		// _ = "end of CoverTab[34549]"
//line /usr/local/go/src/net/textproto/reader.go:59
		_go_fuzz_dep_.CoverTab[34550]++

								if line == nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:61
			_go_fuzz_dep_.CoverTab[34554]++
//line /usr/local/go/src/net/textproto/reader.go:61
			return !more
//line /usr/local/go/src/net/textproto/reader.go:61
			// _ = "end of CoverTab[34554]"
//line /usr/local/go/src/net/textproto/reader.go:61
		}() {
//line /usr/local/go/src/net/textproto/reader.go:61
			_go_fuzz_dep_.CoverTab[34555]++
									return l, nil
//line /usr/local/go/src/net/textproto/reader.go:62
			// _ = "end of CoverTab[34555]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:63
			_go_fuzz_dep_.CoverTab[34556]++
//line /usr/local/go/src/net/textproto/reader.go:63
			// _ = "end of CoverTab[34556]"
//line /usr/local/go/src/net/textproto/reader.go:63
		}
//line /usr/local/go/src/net/textproto/reader.go:63
		// _ = "end of CoverTab[34550]"
//line /usr/local/go/src/net/textproto/reader.go:63
		_go_fuzz_dep_.CoverTab[34551]++
								line = append(line, l...)
								if !more {
//line /usr/local/go/src/net/textproto/reader.go:65
			_go_fuzz_dep_.CoverTab[34557]++
									break
//line /usr/local/go/src/net/textproto/reader.go:66
			// _ = "end of CoverTab[34557]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:67
			_go_fuzz_dep_.CoverTab[34558]++
//line /usr/local/go/src/net/textproto/reader.go:67
			// _ = "end of CoverTab[34558]"
//line /usr/local/go/src/net/textproto/reader.go:67
		}
//line /usr/local/go/src/net/textproto/reader.go:67
		// _ = "end of CoverTab[34551]"
	}
//line /usr/local/go/src/net/textproto/reader.go:68
	// _ = "end of CoverTab[34547]"
//line /usr/local/go/src/net/textproto/reader.go:68
	_go_fuzz_dep_.CoverTab[34548]++
							return line, nil
//line /usr/local/go/src/net/textproto/reader.go:69
	// _ = "end of CoverTab[34548]"
}

// ReadContinuedLine reads a possibly continued line from r,
//line /usr/local/go/src/net/textproto/reader.go:72
// eliding the final trailing ASCII white space.
//line /usr/local/go/src/net/textproto/reader.go:72
// Lines after the first are considered continuations if they
//line /usr/local/go/src/net/textproto/reader.go:72
// begin with a space or tab character. In the returned data,
//line /usr/local/go/src/net/textproto/reader.go:72
// continuation lines are separated from the previous line
//line /usr/local/go/src/net/textproto/reader.go:72
// only by a single space: the newline and leading white space
//line /usr/local/go/src/net/textproto/reader.go:72
// are removed.
//line /usr/local/go/src/net/textproto/reader.go:72
//
//line /usr/local/go/src/net/textproto/reader.go:72
// For example, consider this input:
//line /usr/local/go/src/net/textproto/reader.go:72
//
//line /usr/local/go/src/net/textproto/reader.go:72
//	Line 1
//line /usr/local/go/src/net/textproto/reader.go:72
//	  continued...
//line /usr/local/go/src/net/textproto/reader.go:72
//	Line 2
//line /usr/local/go/src/net/textproto/reader.go:72
//
//line /usr/local/go/src/net/textproto/reader.go:72
// The first call to ReadContinuedLine will return "Line 1 continued..."
//line /usr/local/go/src/net/textproto/reader.go:72
// and the second will return "Line 2".
//line /usr/local/go/src/net/textproto/reader.go:72
//
//line /usr/local/go/src/net/textproto/reader.go:72
// Empty lines are never continued.
//line /usr/local/go/src/net/textproto/reader.go:90
func (r *Reader) ReadContinuedLine() (string, error) {
//line /usr/local/go/src/net/textproto/reader.go:90
	_go_fuzz_dep_.CoverTab[34559]++
							line, err := r.readContinuedLineSlice(noValidation)
							return string(line), err
//line /usr/local/go/src/net/textproto/reader.go:92
	// _ = "end of CoverTab[34559]"
}

// trim returns s with leading and trailing spaces and tabs removed.
//line /usr/local/go/src/net/textproto/reader.go:95
// It does not assume Unicode or UTF-8.
//line /usr/local/go/src/net/textproto/reader.go:97
func trim(s []byte) []byte {
//line /usr/local/go/src/net/textproto/reader.go:97
	_go_fuzz_dep_.CoverTab[34560]++
							i := 0
							for i < len(s) && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:99
		_go_fuzz_dep_.CoverTab[34563]++
//line /usr/local/go/src/net/textproto/reader.go:99
		return (s[i] == ' ' || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:99
			_go_fuzz_dep_.CoverTab[34564]++
//line /usr/local/go/src/net/textproto/reader.go:99
			return s[i] == '\t'
//line /usr/local/go/src/net/textproto/reader.go:99
			// _ = "end of CoverTab[34564]"
//line /usr/local/go/src/net/textproto/reader.go:99
		}())
//line /usr/local/go/src/net/textproto/reader.go:99
		// _ = "end of CoverTab[34563]"
//line /usr/local/go/src/net/textproto/reader.go:99
	}() {
//line /usr/local/go/src/net/textproto/reader.go:99
		_go_fuzz_dep_.CoverTab[34565]++
								i++
//line /usr/local/go/src/net/textproto/reader.go:100
		// _ = "end of CoverTab[34565]"
	}
//line /usr/local/go/src/net/textproto/reader.go:101
	// _ = "end of CoverTab[34560]"
//line /usr/local/go/src/net/textproto/reader.go:101
	_go_fuzz_dep_.CoverTab[34561]++
							n := len(s)
							for n > i && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:103
		_go_fuzz_dep_.CoverTab[34566]++
//line /usr/local/go/src/net/textproto/reader.go:103
		return (s[n-1] == ' ' || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:103
			_go_fuzz_dep_.CoverTab[34567]++
//line /usr/local/go/src/net/textproto/reader.go:103
			return s[n-1] == '\t'
//line /usr/local/go/src/net/textproto/reader.go:103
			// _ = "end of CoverTab[34567]"
//line /usr/local/go/src/net/textproto/reader.go:103
		}())
//line /usr/local/go/src/net/textproto/reader.go:103
		// _ = "end of CoverTab[34566]"
//line /usr/local/go/src/net/textproto/reader.go:103
	}() {
//line /usr/local/go/src/net/textproto/reader.go:103
		_go_fuzz_dep_.CoverTab[34568]++
								n--
//line /usr/local/go/src/net/textproto/reader.go:104
		// _ = "end of CoverTab[34568]"
	}
//line /usr/local/go/src/net/textproto/reader.go:105
	// _ = "end of CoverTab[34561]"
//line /usr/local/go/src/net/textproto/reader.go:105
	_go_fuzz_dep_.CoverTab[34562]++
							return s[i:n]
//line /usr/local/go/src/net/textproto/reader.go:106
	// _ = "end of CoverTab[34562]"
}

// ReadContinuedLineBytes is like ReadContinuedLine but
//line /usr/local/go/src/net/textproto/reader.go:109
// returns a []byte instead of a string.
//line /usr/local/go/src/net/textproto/reader.go:111
func (r *Reader) ReadContinuedLineBytes() ([]byte, error) {
//line /usr/local/go/src/net/textproto/reader.go:111
	_go_fuzz_dep_.CoverTab[34569]++
							line, err := r.readContinuedLineSlice(noValidation)
							if line != nil {
//line /usr/local/go/src/net/textproto/reader.go:113
		_go_fuzz_dep_.CoverTab[34571]++
								line = bytes.Clone(line)
//line /usr/local/go/src/net/textproto/reader.go:114
		// _ = "end of CoverTab[34571]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:115
		_go_fuzz_dep_.CoverTab[34572]++
//line /usr/local/go/src/net/textproto/reader.go:115
		// _ = "end of CoverTab[34572]"
//line /usr/local/go/src/net/textproto/reader.go:115
	}
//line /usr/local/go/src/net/textproto/reader.go:115
	// _ = "end of CoverTab[34569]"
//line /usr/local/go/src/net/textproto/reader.go:115
	_go_fuzz_dep_.CoverTab[34570]++
							return line, err
//line /usr/local/go/src/net/textproto/reader.go:116
	// _ = "end of CoverTab[34570]"
}

// readContinuedLineSlice reads continued lines from the reader buffer,
//line /usr/local/go/src/net/textproto/reader.go:119
// returning a byte slice with all lines. The validateFirstLine function
//line /usr/local/go/src/net/textproto/reader.go:119
// is run on the first read line, and if it returns an error then this
//line /usr/local/go/src/net/textproto/reader.go:119
// error is returned from readContinuedLineSlice.
//line /usr/local/go/src/net/textproto/reader.go:123
func (r *Reader) readContinuedLineSlice(validateFirstLine func([]byte) error) ([]byte, error) {
//line /usr/local/go/src/net/textproto/reader.go:123
	_go_fuzz_dep_.CoverTab[34573]++
							if validateFirstLine == nil {
//line /usr/local/go/src/net/textproto/reader.go:124
		_go_fuzz_dep_.CoverTab[34580]++
								return nil, fmt.Errorf("missing validateFirstLine func")
//line /usr/local/go/src/net/textproto/reader.go:125
		// _ = "end of CoverTab[34580]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:126
		_go_fuzz_dep_.CoverTab[34581]++
//line /usr/local/go/src/net/textproto/reader.go:126
		// _ = "end of CoverTab[34581]"
//line /usr/local/go/src/net/textproto/reader.go:126
	}
//line /usr/local/go/src/net/textproto/reader.go:126
	// _ = "end of CoverTab[34573]"
//line /usr/local/go/src/net/textproto/reader.go:126
	_go_fuzz_dep_.CoverTab[34574]++

//line /usr/local/go/src/net/textproto/reader.go:129
	line, err := r.readLineSlice()
	if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:130
		_go_fuzz_dep_.CoverTab[34582]++
								return nil, err
//line /usr/local/go/src/net/textproto/reader.go:131
		// _ = "end of CoverTab[34582]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:132
		_go_fuzz_dep_.CoverTab[34583]++
//line /usr/local/go/src/net/textproto/reader.go:132
		// _ = "end of CoverTab[34583]"
//line /usr/local/go/src/net/textproto/reader.go:132
	}
//line /usr/local/go/src/net/textproto/reader.go:132
	// _ = "end of CoverTab[34574]"
//line /usr/local/go/src/net/textproto/reader.go:132
	_go_fuzz_dep_.CoverTab[34575]++
							if len(line) == 0 {
//line /usr/local/go/src/net/textproto/reader.go:133
		_go_fuzz_dep_.CoverTab[34584]++
								return line, nil
//line /usr/local/go/src/net/textproto/reader.go:134
		// _ = "end of CoverTab[34584]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:135
		_go_fuzz_dep_.CoverTab[34585]++
//line /usr/local/go/src/net/textproto/reader.go:135
		// _ = "end of CoverTab[34585]"
//line /usr/local/go/src/net/textproto/reader.go:135
	}
//line /usr/local/go/src/net/textproto/reader.go:135
	// _ = "end of CoverTab[34575]"
//line /usr/local/go/src/net/textproto/reader.go:135
	_go_fuzz_dep_.CoverTab[34576]++

							if err := validateFirstLine(line); err != nil {
//line /usr/local/go/src/net/textproto/reader.go:137
		_go_fuzz_dep_.CoverTab[34586]++
								return nil, err
//line /usr/local/go/src/net/textproto/reader.go:138
		// _ = "end of CoverTab[34586]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:139
		_go_fuzz_dep_.CoverTab[34587]++
//line /usr/local/go/src/net/textproto/reader.go:139
		// _ = "end of CoverTab[34587]"
//line /usr/local/go/src/net/textproto/reader.go:139
	}
//line /usr/local/go/src/net/textproto/reader.go:139
	// _ = "end of CoverTab[34576]"
//line /usr/local/go/src/net/textproto/reader.go:139
	_go_fuzz_dep_.CoverTab[34577]++

//line /usr/local/go/src/net/textproto/reader.go:145
	if r.R.Buffered() > 1 {
//line /usr/local/go/src/net/textproto/reader.go:145
		_go_fuzz_dep_.CoverTab[34588]++
								peek, _ := r.R.Peek(2)
								if len(peek) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:147
			_go_fuzz_dep_.CoverTab[34589]++
//line /usr/local/go/src/net/textproto/reader.go:147
			return (isASCIILetter(peek[0]) || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:147
				_go_fuzz_dep_.CoverTab[34590]++
//line /usr/local/go/src/net/textproto/reader.go:147
				return peek[0] == '\n'
//line /usr/local/go/src/net/textproto/reader.go:147
				// _ = "end of CoverTab[34590]"
//line /usr/local/go/src/net/textproto/reader.go:147
			}())
//line /usr/local/go/src/net/textproto/reader.go:147
			// _ = "end of CoverTab[34589]"
//line /usr/local/go/src/net/textproto/reader.go:147
		}() || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:147
			_go_fuzz_dep_.CoverTab[34591]++
//line /usr/local/go/src/net/textproto/reader.go:147
			return len(peek) == 2 && func() bool {
										_go_fuzz_dep_.CoverTab[34592]++
//line /usr/local/go/src/net/textproto/reader.go:148
				return peek[0] == '\r'
//line /usr/local/go/src/net/textproto/reader.go:148
				// _ = "end of CoverTab[34592]"
//line /usr/local/go/src/net/textproto/reader.go:148
			}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:148
				_go_fuzz_dep_.CoverTab[34593]++
//line /usr/local/go/src/net/textproto/reader.go:148
				return peek[1] == '\n'
//line /usr/local/go/src/net/textproto/reader.go:148
				// _ = "end of CoverTab[34593]"
//line /usr/local/go/src/net/textproto/reader.go:148
			}()
//line /usr/local/go/src/net/textproto/reader.go:148
			// _ = "end of CoverTab[34591]"
//line /usr/local/go/src/net/textproto/reader.go:148
		}() {
//line /usr/local/go/src/net/textproto/reader.go:148
			_go_fuzz_dep_.CoverTab[34594]++
									return trim(line), nil
//line /usr/local/go/src/net/textproto/reader.go:149
			// _ = "end of CoverTab[34594]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:150
			_go_fuzz_dep_.CoverTab[34595]++
//line /usr/local/go/src/net/textproto/reader.go:150
			// _ = "end of CoverTab[34595]"
//line /usr/local/go/src/net/textproto/reader.go:150
		}
//line /usr/local/go/src/net/textproto/reader.go:150
		// _ = "end of CoverTab[34588]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:151
		_go_fuzz_dep_.CoverTab[34596]++
//line /usr/local/go/src/net/textproto/reader.go:151
		// _ = "end of CoverTab[34596]"
//line /usr/local/go/src/net/textproto/reader.go:151
	}
//line /usr/local/go/src/net/textproto/reader.go:151
	// _ = "end of CoverTab[34577]"
//line /usr/local/go/src/net/textproto/reader.go:151
	_go_fuzz_dep_.CoverTab[34578]++

//line /usr/local/go/src/net/textproto/reader.go:155
	r.buf = append(r.buf[:0], trim(line)...)

//line /usr/local/go/src/net/textproto/reader.go:158
	for r.skipSpace() > 0 {
//line /usr/local/go/src/net/textproto/reader.go:158
		_go_fuzz_dep_.CoverTab[34597]++
								line, err := r.readLineSlice()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:160
			_go_fuzz_dep_.CoverTab[34599]++
									break
//line /usr/local/go/src/net/textproto/reader.go:161
			// _ = "end of CoverTab[34599]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:162
			_go_fuzz_dep_.CoverTab[34600]++
//line /usr/local/go/src/net/textproto/reader.go:162
			// _ = "end of CoverTab[34600]"
//line /usr/local/go/src/net/textproto/reader.go:162
		}
//line /usr/local/go/src/net/textproto/reader.go:162
		// _ = "end of CoverTab[34597]"
//line /usr/local/go/src/net/textproto/reader.go:162
		_go_fuzz_dep_.CoverTab[34598]++
								r.buf = append(r.buf, ' ')
								r.buf = append(r.buf, trim(line)...)
//line /usr/local/go/src/net/textproto/reader.go:164
		// _ = "end of CoverTab[34598]"
	}
//line /usr/local/go/src/net/textproto/reader.go:165
	// _ = "end of CoverTab[34578]"
//line /usr/local/go/src/net/textproto/reader.go:165
	_go_fuzz_dep_.CoverTab[34579]++
							return r.buf, nil
//line /usr/local/go/src/net/textproto/reader.go:166
	// _ = "end of CoverTab[34579]"
}

// skipSpace skips R over all spaces and returns the number of bytes skipped.
func (r *Reader) skipSpace() int {
//line /usr/local/go/src/net/textproto/reader.go:170
	_go_fuzz_dep_.CoverTab[34601]++
							n := 0
							for {
//line /usr/local/go/src/net/textproto/reader.go:172
		_go_fuzz_dep_.CoverTab[34603]++
								c, err := r.R.ReadByte()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:174
			_go_fuzz_dep_.CoverTab[34606]++

									break
//line /usr/local/go/src/net/textproto/reader.go:176
			// _ = "end of CoverTab[34606]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:177
			_go_fuzz_dep_.CoverTab[34607]++
//line /usr/local/go/src/net/textproto/reader.go:177
			// _ = "end of CoverTab[34607]"
//line /usr/local/go/src/net/textproto/reader.go:177
		}
//line /usr/local/go/src/net/textproto/reader.go:177
		// _ = "end of CoverTab[34603]"
//line /usr/local/go/src/net/textproto/reader.go:177
		_go_fuzz_dep_.CoverTab[34604]++
								if c != ' ' && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:178
			_go_fuzz_dep_.CoverTab[34608]++
//line /usr/local/go/src/net/textproto/reader.go:178
			return c != '\t'
//line /usr/local/go/src/net/textproto/reader.go:178
			// _ = "end of CoverTab[34608]"
//line /usr/local/go/src/net/textproto/reader.go:178
		}() {
//line /usr/local/go/src/net/textproto/reader.go:178
			_go_fuzz_dep_.CoverTab[34609]++
									r.R.UnreadByte()
									break
//line /usr/local/go/src/net/textproto/reader.go:180
			// _ = "end of CoverTab[34609]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:181
			_go_fuzz_dep_.CoverTab[34610]++
//line /usr/local/go/src/net/textproto/reader.go:181
			// _ = "end of CoverTab[34610]"
//line /usr/local/go/src/net/textproto/reader.go:181
		}
//line /usr/local/go/src/net/textproto/reader.go:181
		// _ = "end of CoverTab[34604]"
//line /usr/local/go/src/net/textproto/reader.go:181
		_go_fuzz_dep_.CoverTab[34605]++
								n++
//line /usr/local/go/src/net/textproto/reader.go:182
		// _ = "end of CoverTab[34605]"
	}
//line /usr/local/go/src/net/textproto/reader.go:183
	// _ = "end of CoverTab[34601]"
//line /usr/local/go/src/net/textproto/reader.go:183
	_go_fuzz_dep_.CoverTab[34602]++
							return n
//line /usr/local/go/src/net/textproto/reader.go:184
	// _ = "end of CoverTab[34602]"
}

func (r *Reader) readCodeLine(expectCode int) (code int, continued bool, message string, err error) {
//line /usr/local/go/src/net/textproto/reader.go:187
	_go_fuzz_dep_.CoverTab[34611]++
							line, err := r.ReadLine()
							if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:189
		_go_fuzz_dep_.CoverTab[34613]++
								return
//line /usr/local/go/src/net/textproto/reader.go:190
		// _ = "end of CoverTab[34613]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:191
		_go_fuzz_dep_.CoverTab[34614]++
//line /usr/local/go/src/net/textproto/reader.go:191
		// _ = "end of CoverTab[34614]"
//line /usr/local/go/src/net/textproto/reader.go:191
	}
//line /usr/local/go/src/net/textproto/reader.go:191
	// _ = "end of CoverTab[34611]"
//line /usr/local/go/src/net/textproto/reader.go:191
	_go_fuzz_dep_.CoverTab[34612]++
							return parseCodeLine(line, expectCode)
//line /usr/local/go/src/net/textproto/reader.go:192
	// _ = "end of CoverTab[34612]"
}

func parseCodeLine(line string, expectCode int) (code int, continued bool, message string, err error) {
//line /usr/local/go/src/net/textproto/reader.go:195
	_go_fuzz_dep_.CoverTab[34615]++
							if len(line) < 4 || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:196
		_go_fuzz_dep_.CoverTab[34619]++
//line /usr/local/go/src/net/textproto/reader.go:196
		return line[3] != ' ' && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:196
			_go_fuzz_dep_.CoverTab[34620]++
//line /usr/local/go/src/net/textproto/reader.go:196
			return line[3] != '-'
//line /usr/local/go/src/net/textproto/reader.go:196
			// _ = "end of CoverTab[34620]"
//line /usr/local/go/src/net/textproto/reader.go:196
		}()
//line /usr/local/go/src/net/textproto/reader.go:196
		// _ = "end of CoverTab[34619]"
//line /usr/local/go/src/net/textproto/reader.go:196
	}() {
//line /usr/local/go/src/net/textproto/reader.go:196
		_go_fuzz_dep_.CoverTab[34621]++
								err = ProtocolError("short response: " + line)
								return
//line /usr/local/go/src/net/textproto/reader.go:198
		// _ = "end of CoverTab[34621]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:199
		_go_fuzz_dep_.CoverTab[34622]++
//line /usr/local/go/src/net/textproto/reader.go:199
		// _ = "end of CoverTab[34622]"
//line /usr/local/go/src/net/textproto/reader.go:199
	}
//line /usr/local/go/src/net/textproto/reader.go:199
	// _ = "end of CoverTab[34615]"
//line /usr/local/go/src/net/textproto/reader.go:199
	_go_fuzz_dep_.CoverTab[34616]++
							continued = line[3] == '-'
							code, err = strconv.Atoi(line[0:3])
							if err != nil || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:202
		_go_fuzz_dep_.CoverTab[34623]++
//line /usr/local/go/src/net/textproto/reader.go:202
		return code < 100
//line /usr/local/go/src/net/textproto/reader.go:202
		// _ = "end of CoverTab[34623]"
//line /usr/local/go/src/net/textproto/reader.go:202
	}() {
//line /usr/local/go/src/net/textproto/reader.go:202
		_go_fuzz_dep_.CoverTab[34624]++
								err = ProtocolError("invalid response code: " + line)
								return
//line /usr/local/go/src/net/textproto/reader.go:204
		// _ = "end of CoverTab[34624]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:205
		_go_fuzz_dep_.CoverTab[34625]++
//line /usr/local/go/src/net/textproto/reader.go:205
		// _ = "end of CoverTab[34625]"
//line /usr/local/go/src/net/textproto/reader.go:205
	}
//line /usr/local/go/src/net/textproto/reader.go:205
	// _ = "end of CoverTab[34616]"
//line /usr/local/go/src/net/textproto/reader.go:205
	_go_fuzz_dep_.CoverTab[34617]++
							message = line[4:]
							if 1 <= expectCode && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:207
		_go_fuzz_dep_.CoverTab[34626]++
//line /usr/local/go/src/net/textproto/reader.go:207
		return expectCode < 10
//line /usr/local/go/src/net/textproto/reader.go:207
		// _ = "end of CoverTab[34626]"
//line /usr/local/go/src/net/textproto/reader.go:207
	}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:207
		_go_fuzz_dep_.CoverTab[34627]++
//line /usr/local/go/src/net/textproto/reader.go:207
		return code/100 != expectCode
//line /usr/local/go/src/net/textproto/reader.go:207
		// _ = "end of CoverTab[34627]"
//line /usr/local/go/src/net/textproto/reader.go:207
	}() || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:207
		_go_fuzz_dep_.CoverTab[34628]++
//line /usr/local/go/src/net/textproto/reader.go:207
		return 10 <= expectCode && func() bool {
									_go_fuzz_dep_.CoverTab[34629]++
//line /usr/local/go/src/net/textproto/reader.go:208
			return expectCode < 100
//line /usr/local/go/src/net/textproto/reader.go:208
			// _ = "end of CoverTab[34629]"
//line /usr/local/go/src/net/textproto/reader.go:208
		}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:208
			_go_fuzz_dep_.CoverTab[34630]++
//line /usr/local/go/src/net/textproto/reader.go:208
			return code/10 != expectCode
//line /usr/local/go/src/net/textproto/reader.go:208
			// _ = "end of CoverTab[34630]"
//line /usr/local/go/src/net/textproto/reader.go:208
		}()
//line /usr/local/go/src/net/textproto/reader.go:208
		// _ = "end of CoverTab[34628]"
//line /usr/local/go/src/net/textproto/reader.go:208
	}() || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:208
		_go_fuzz_dep_.CoverTab[34631]++
//line /usr/local/go/src/net/textproto/reader.go:208
		return 100 <= expectCode && func() bool {
									_go_fuzz_dep_.CoverTab[34632]++
//line /usr/local/go/src/net/textproto/reader.go:209
			return expectCode < 1000
//line /usr/local/go/src/net/textproto/reader.go:209
			// _ = "end of CoverTab[34632]"
//line /usr/local/go/src/net/textproto/reader.go:209
		}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:209
			_go_fuzz_dep_.CoverTab[34633]++
//line /usr/local/go/src/net/textproto/reader.go:209
			return code != expectCode
//line /usr/local/go/src/net/textproto/reader.go:209
			// _ = "end of CoverTab[34633]"
//line /usr/local/go/src/net/textproto/reader.go:209
		}()
//line /usr/local/go/src/net/textproto/reader.go:209
		// _ = "end of CoverTab[34631]"
//line /usr/local/go/src/net/textproto/reader.go:209
	}() {
//line /usr/local/go/src/net/textproto/reader.go:209
		_go_fuzz_dep_.CoverTab[34634]++
								err = &Error{code, message}
//line /usr/local/go/src/net/textproto/reader.go:210
		// _ = "end of CoverTab[34634]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:211
		_go_fuzz_dep_.CoverTab[34635]++
//line /usr/local/go/src/net/textproto/reader.go:211
		// _ = "end of CoverTab[34635]"
//line /usr/local/go/src/net/textproto/reader.go:211
	}
//line /usr/local/go/src/net/textproto/reader.go:211
	// _ = "end of CoverTab[34617]"
//line /usr/local/go/src/net/textproto/reader.go:211
	_go_fuzz_dep_.CoverTab[34618]++
							return
//line /usr/local/go/src/net/textproto/reader.go:212
	// _ = "end of CoverTab[34618]"
}

// ReadCodeLine reads a response code line of the form
//line /usr/local/go/src/net/textproto/reader.go:215
//
//line /usr/local/go/src/net/textproto/reader.go:215
//	code message
//line /usr/local/go/src/net/textproto/reader.go:215
//
//line /usr/local/go/src/net/textproto/reader.go:215
// where code is a three-digit status code and the message
//line /usr/local/go/src/net/textproto/reader.go:215
// extends to the rest of the line. An example of such a line is:
//line /usr/local/go/src/net/textproto/reader.go:215
//
//line /usr/local/go/src/net/textproto/reader.go:215
//	220 plan9.bell-labs.com ESMTP
//line /usr/local/go/src/net/textproto/reader.go:215
//
//line /usr/local/go/src/net/textproto/reader.go:215
// If the prefix of the status does not match the digits in expectCode,
//line /usr/local/go/src/net/textproto/reader.go:215
// ReadCodeLine returns with err set to &Error{code, message}.
//line /usr/local/go/src/net/textproto/reader.go:215
// For example, if expectCode is 31, an error will be returned if
//line /usr/local/go/src/net/textproto/reader.go:215
// the status is not in the range [310,319].
//line /usr/local/go/src/net/textproto/reader.go:215
//
//line /usr/local/go/src/net/textproto/reader.go:215
// If the response is multi-line, ReadCodeLine returns an error.
//line /usr/local/go/src/net/textproto/reader.go:215
//
//line /usr/local/go/src/net/textproto/reader.go:215
// An expectCode <= 0 disables the check of the status code.
//line /usr/local/go/src/net/textproto/reader.go:232
func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error) {
//line /usr/local/go/src/net/textproto/reader.go:232
	_go_fuzz_dep_.CoverTab[34636]++
							code, continued, message, err := r.readCodeLine(expectCode)
							if err == nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:234
		_go_fuzz_dep_.CoverTab[34638]++
//line /usr/local/go/src/net/textproto/reader.go:234
		return continued
//line /usr/local/go/src/net/textproto/reader.go:234
		// _ = "end of CoverTab[34638]"
//line /usr/local/go/src/net/textproto/reader.go:234
	}() {
//line /usr/local/go/src/net/textproto/reader.go:234
		_go_fuzz_dep_.CoverTab[34639]++
								err = ProtocolError("unexpected multi-line response: " + message)
//line /usr/local/go/src/net/textproto/reader.go:235
		// _ = "end of CoverTab[34639]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:236
		_go_fuzz_dep_.CoverTab[34640]++
//line /usr/local/go/src/net/textproto/reader.go:236
		// _ = "end of CoverTab[34640]"
//line /usr/local/go/src/net/textproto/reader.go:236
	}
//line /usr/local/go/src/net/textproto/reader.go:236
	// _ = "end of CoverTab[34636]"
//line /usr/local/go/src/net/textproto/reader.go:236
	_go_fuzz_dep_.CoverTab[34637]++
							return
//line /usr/local/go/src/net/textproto/reader.go:237
	// _ = "end of CoverTab[34637]"
}

// ReadResponse reads a multi-line response of the form:
//line /usr/local/go/src/net/textproto/reader.go:240
//
//line /usr/local/go/src/net/textproto/reader.go:240
//	code-message line 1
//line /usr/local/go/src/net/textproto/reader.go:240
//	code-message line 2
//line /usr/local/go/src/net/textproto/reader.go:240
//	...
//line /usr/local/go/src/net/textproto/reader.go:240
//	code message line n
//line /usr/local/go/src/net/textproto/reader.go:240
//
//line /usr/local/go/src/net/textproto/reader.go:240
// where code is a three-digit status code. The first line starts with the
//line /usr/local/go/src/net/textproto/reader.go:240
// code and a hyphen. The response is terminated by a line that starts
//line /usr/local/go/src/net/textproto/reader.go:240
// with the same code followed by a space. Each line in message is
//line /usr/local/go/src/net/textproto/reader.go:240
// separated by a newline (\n).
//line /usr/local/go/src/net/textproto/reader.go:240
//
//line /usr/local/go/src/net/textproto/reader.go:240
// See page 36 of RFC 959 (https://www.ietf.org/rfc/rfc959.txt) for
//line /usr/local/go/src/net/textproto/reader.go:240
// details of another form of response accepted:
//line /usr/local/go/src/net/textproto/reader.go:240
//
//line /usr/local/go/src/net/textproto/reader.go:240
//	code-message line 1
//line /usr/local/go/src/net/textproto/reader.go:240
//	message line 2
//line /usr/local/go/src/net/textproto/reader.go:240
//	...
//line /usr/local/go/src/net/textproto/reader.go:240
//	code message line n
//line /usr/local/go/src/net/textproto/reader.go:240
//
//line /usr/local/go/src/net/textproto/reader.go:240
// If the prefix of the status does not match the digits in expectCode,
//line /usr/local/go/src/net/textproto/reader.go:240
// ReadResponse returns with err set to &Error{code, message}.
//line /usr/local/go/src/net/textproto/reader.go:240
// For example, if expectCode is 31, an error will be returned if
//line /usr/local/go/src/net/textproto/reader.go:240
// the status is not in the range [310,319].
//line /usr/local/go/src/net/textproto/reader.go:240
//
//line /usr/local/go/src/net/textproto/reader.go:240
// An expectCode <= 0 disables the check of the status code.
//line /usr/local/go/src/net/textproto/reader.go:266
func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error) {
//line /usr/local/go/src/net/textproto/reader.go:266
	_go_fuzz_dep_.CoverTab[34641]++
							code, continued, message, err := r.readCodeLine(expectCode)
							multi := continued
							for continued {
//line /usr/local/go/src/net/textproto/reader.go:269
		_go_fuzz_dep_.CoverTab[34644]++
								line, err := r.ReadLine()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:271
			_go_fuzz_dep_.CoverTab[34647]++
									return 0, "", err
//line /usr/local/go/src/net/textproto/reader.go:272
			// _ = "end of CoverTab[34647]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:273
			_go_fuzz_dep_.CoverTab[34648]++
//line /usr/local/go/src/net/textproto/reader.go:273
			// _ = "end of CoverTab[34648]"
//line /usr/local/go/src/net/textproto/reader.go:273
		}
//line /usr/local/go/src/net/textproto/reader.go:273
		// _ = "end of CoverTab[34644]"
//line /usr/local/go/src/net/textproto/reader.go:273
		_go_fuzz_dep_.CoverTab[34645]++

								var code2 int
								var moreMessage string
								code2, continued, moreMessage, err = parseCodeLine(line, 0)
								if err != nil || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:278
			_go_fuzz_dep_.CoverTab[34649]++
//line /usr/local/go/src/net/textproto/reader.go:278
			return code2 != code
//line /usr/local/go/src/net/textproto/reader.go:278
			// _ = "end of CoverTab[34649]"
//line /usr/local/go/src/net/textproto/reader.go:278
		}() {
//line /usr/local/go/src/net/textproto/reader.go:278
			_go_fuzz_dep_.CoverTab[34650]++
									message += "\n" + strings.TrimRight(line, "\r\n")
									continued = true
									continue
//line /usr/local/go/src/net/textproto/reader.go:281
			// _ = "end of CoverTab[34650]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:282
			_go_fuzz_dep_.CoverTab[34651]++
//line /usr/local/go/src/net/textproto/reader.go:282
			// _ = "end of CoverTab[34651]"
//line /usr/local/go/src/net/textproto/reader.go:282
		}
//line /usr/local/go/src/net/textproto/reader.go:282
		// _ = "end of CoverTab[34645]"
//line /usr/local/go/src/net/textproto/reader.go:282
		_go_fuzz_dep_.CoverTab[34646]++
								message += "\n" + moreMessage
//line /usr/local/go/src/net/textproto/reader.go:283
		// _ = "end of CoverTab[34646]"
	}
//line /usr/local/go/src/net/textproto/reader.go:284
	// _ = "end of CoverTab[34641]"
//line /usr/local/go/src/net/textproto/reader.go:284
	_go_fuzz_dep_.CoverTab[34642]++
							if err != nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:285
		_go_fuzz_dep_.CoverTab[34652]++
//line /usr/local/go/src/net/textproto/reader.go:285
		return multi
//line /usr/local/go/src/net/textproto/reader.go:285
		// _ = "end of CoverTab[34652]"
//line /usr/local/go/src/net/textproto/reader.go:285
	}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:285
		_go_fuzz_dep_.CoverTab[34653]++
//line /usr/local/go/src/net/textproto/reader.go:285
		return message != ""
//line /usr/local/go/src/net/textproto/reader.go:285
		// _ = "end of CoverTab[34653]"
//line /usr/local/go/src/net/textproto/reader.go:285
	}() {
//line /usr/local/go/src/net/textproto/reader.go:285
		_go_fuzz_dep_.CoverTab[34654]++

								err = &Error{code, message}
//line /usr/local/go/src/net/textproto/reader.go:287
		// _ = "end of CoverTab[34654]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:288
		_go_fuzz_dep_.CoverTab[34655]++
//line /usr/local/go/src/net/textproto/reader.go:288
		// _ = "end of CoverTab[34655]"
//line /usr/local/go/src/net/textproto/reader.go:288
	}
//line /usr/local/go/src/net/textproto/reader.go:288
	// _ = "end of CoverTab[34642]"
//line /usr/local/go/src/net/textproto/reader.go:288
	_go_fuzz_dep_.CoverTab[34643]++
							return
//line /usr/local/go/src/net/textproto/reader.go:289
	// _ = "end of CoverTab[34643]"
}

// DotReader returns a new Reader that satisfies Reads using the
//line /usr/local/go/src/net/textproto/reader.go:292
// decoded text of a dot-encoded block read from r.
//line /usr/local/go/src/net/textproto/reader.go:292
// The returned Reader is only valid until the next call
//line /usr/local/go/src/net/textproto/reader.go:292
// to a method on r.
//line /usr/local/go/src/net/textproto/reader.go:292
//
//line /usr/local/go/src/net/textproto/reader.go:292
// Dot encoding is a common framing used for data blocks
//line /usr/local/go/src/net/textproto/reader.go:292
// in text protocols such as SMTP.  The data consists of a sequence
//line /usr/local/go/src/net/textproto/reader.go:292
// of lines, each of which ends in "\r\n".  The sequence itself
//line /usr/local/go/src/net/textproto/reader.go:292
// ends at a line containing just a dot: ".\r\n".  Lines beginning
//line /usr/local/go/src/net/textproto/reader.go:292
// with a dot are escaped with an additional dot to avoid
//line /usr/local/go/src/net/textproto/reader.go:292
// looking like the end of the sequence.
//line /usr/local/go/src/net/textproto/reader.go:292
//
//line /usr/local/go/src/net/textproto/reader.go:292
// The decoded form returned by the Reader's Read method
//line /usr/local/go/src/net/textproto/reader.go:292
// rewrites the "\r\n" line endings into the simpler "\n",
//line /usr/local/go/src/net/textproto/reader.go:292
// removes leading dot escapes if present, and stops with error io.EOF
//line /usr/local/go/src/net/textproto/reader.go:292
// after consuming (and discarding) the end-of-sequence line.
//line /usr/local/go/src/net/textproto/reader.go:308
func (r *Reader) DotReader() io.Reader {
//line /usr/local/go/src/net/textproto/reader.go:308
	_go_fuzz_dep_.CoverTab[34656]++
							r.closeDot()
							r.dot = &dotReader{r: r}
							return r.dot
//line /usr/local/go/src/net/textproto/reader.go:311
	// _ = "end of CoverTab[34656]"
}

type dotReader struct {
	r	*Reader
	state	int
}

// Read satisfies reads by decoding dot-encoded data read from d.r.
func (d *dotReader) Read(b []byte) (n int, err error) {
//line /usr/local/go/src/net/textproto/reader.go:320
	_go_fuzz_dep_.CoverTab[34657]++
	// Run data through a simple state machine to
	// elide leading dots, rewrite trailing \r\n into \n,
	// and detect ending .\r\n line.
	const (
		stateBeginLine	= iota	// beginning of line; initial state; must be zero
		stateDot		// read . at beginning of line
		stateDotCR		// read .\r at beginning of line
		stateCR			// read \r (possibly at end of line)
		stateData		// reading data in middle of line
		stateEOF		// reached .\r\n end marker line
	)
	br := d.r.R
	for n < len(b) && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:333
		_go_fuzz_dep_.CoverTab[34661]++
//line /usr/local/go/src/net/textproto/reader.go:333
		return d.state != stateEOF
//line /usr/local/go/src/net/textproto/reader.go:333
		// _ = "end of CoverTab[34661]"
//line /usr/local/go/src/net/textproto/reader.go:333
	}() {
//line /usr/local/go/src/net/textproto/reader.go:333
		_go_fuzz_dep_.CoverTab[34662]++
								var c byte
								c, err = br.ReadByte()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:336
			_go_fuzz_dep_.CoverTab[34665]++
									if err == io.EOF {
//line /usr/local/go/src/net/textproto/reader.go:337
				_go_fuzz_dep_.CoverTab[34667]++
										err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/textproto/reader.go:338
				// _ = "end of CoverTab[34667]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:339
				_go_fuzz_dep_.CoverTab[34668]++
//line /usr/local/go/src/net/textproto/reader.go:339
				// _ = "end of CoverTab[34668]"
//line /usr/local/go/src/net/textproto/reader.go:339
			}
//line /usr/local/go/src/net/textproto/reader.go:339
			// _ = "end of CoverTab[34665]"
//line /usr/local/go/src/net/textproto/reader.go:339
			_go_fuzz_dep_.CoverTab[34666]++
									break
//line /usr/local/go/src/net/textproto/reader.go:340
			// _ = "end of CoverTab[34666]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:341
			_go_fuzz_dep_.CoverTab[34669]++
//line /usr/local/go/src/net/textproto/reader.go:341
			// _ = "end of CoverTab[34669]"
//line /usr/local/go/src/net/textproto/reader.go:341
		}
//line /usr/local/go/src/net/textproto/reader.go:341
		// _ = "end of CoverTab[34662]"
//line /usr/local/go/src/net/textproto/reader.go:341
		_go_fuzz_dep_.CoverTab[34663]++
								switch d.state {
		case stateBeginLine:
//line /usr/local/go/src/net/textproto/reader.go:343
			_go_fuzz_dep_.CoverTab[34670]++
									if c == '.' {
//line /usr/local/go/src/net/textproto/reader.go:344
				_go_fuzz_dep_.CoverTab[34683]++
										d.state = stateDot
										continue
//line /usr/local/go/src/net/textproto/reader.go:346
				// _ = "end of CoverTab[34683]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:347
				_go_fuzz_dep_.CoverTab[34684]++
//line /usr/local/go/src/net/textproto/reader.go:347
				// _ = "end of CoverTab[34684]"
//line /usr/local/go/src/net/textproto/reader.go:347
			}
//line /usr/local/go/src/net/textproto/reader.go:347
			// _ = "end of CoverTab[34670]"
//line /usr/local/go/src/net/textproto/reader.go:347
			_go_fuzz_dep_.CoverTab[34671]++
									if c == '\r' {
//line /usr/local/go/src/net/textproto/reader.go:348
				_go_fuzz_dep_.CoverTab[34685]++
										d.state = stateCR
										continue
//line /usr/local/go/src/net/textproto/reader.go:350
				// _ = "end of CoverTab[34685]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:351
				_go_fuzz_dep_.CoverTab[34686]++
//line /usr/local/go/src/net/textproto/reader.go:351
				// _ = "end of CoverTab[34686]"
//line /usr/local/go/src/net/textproto/reader.go:351
			}
//line /usr/local/go/src/net/textproto/reader.go:351
			// _ = "end of CoverTab[34671]"
//line /usr/local/go/src/net/textproto/reader.go:351
			_go_fuzz_dep_.CoverTab[34672]++
									d.state = stateData
//line /usr/local/go/src/net/textproto/reader.go:352
			// _ = "end of CoverTab[34672]"

		case stateDot:
//line /usr/local/go/src/net/textproto/reader.go:354
			_go_fuzz_dep_.CoverTab[34673]++
									if c == '\r' {
//line /usr/local/go/src/net/textproto/reader.go:355
				_go_fuzz_dep_.CoverTab[34687]++
										d.state = stateDotCR
										continue
//line /usr/local/go/src/net/textproto/reader.go:357
				// _ = "end of CoverTab[34687]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:358
				_go_fuzz_dep_.CoverTab[34688]++
//line /usr/local/go/src/net/textproto/reader.go:358
				// _ = "end of CoverTab[34688]"
//line /usr/local/go/src/net/textproto/reader.go:358
			}
//line /usr/local/go/src/net/textproto/reader.go:358
			// _ = "end of CoverTab[34673]"
//line /usr/local/go/src/net/textproto/reader.go:358
			_go_fuzz_dep_.CoverTab[34674]++
									if c == '\n' {
//line /usr/local/go/src/net/textproto/reader.go:359
				_go_fuzz_dep_.CoverTab[34689]++
										d.state = stateEOF
										continue
//line /usr/local/go/src/net/textproto/reader.go:361
				// _ = "end of CoverTab[34689]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:362
				_go_fuzz_dep_.CoverTab[34690]++
//line /usr/local/go/src/net/textproto/reader.go:362
				// _ = "end of CoverTab[34690]"
//line /usr/local/go/src/net/textproto/reader.go:362
			}
//line /usr/local/go/src/net/textproto/reader.go:362
			// _ = "end of CoverTab[34674]"
//line /usr/local/go/src/net/textproto/reader.go:362
			_go_fuzz_dep_.CoverTab[34675]++
									d.state = stateData
//line /usr/local/go/src/net/textproto/reader.go:363
			// _ = "end of CoverTab[34675]"

		case stateDotCR:
//line /usr/local/go/src/net/textproto/reader.go:365
			_go_fuzz_dep_.CoverTab[34676]++
									if c == '\n' {
//line /usr/local/go/src/net/textproto/reader.go:366
				_go_fuzz_dep_.CoverTab[34691]++
										d.state = stateEOF
										continue
//line /usr/local/go/src/net/textproto/reader.go:368
				// _ = "end of CoverTab[34691]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:369
				_go_fuzz_dep_.CoverTab[34692]++
//line /usr/local/go/src/net/textproto/reader.go:369
				// _ = "end of CoverTab[34692]"
//line /usr/local/go/src/net/textproto/reader.go:369
			}
//line /usr/local/go/src/net/textproto/reader.go:369
			// _ = "end of CoverTab[34676]"
//line /usr/local/go/src/net/textproto/reader.go:369
			_go_fuzz_dep_.CoverTab[34677]++

//line /usr/local/go/src/net/textproto/reader.go:372
			br.UnreadByte()
									c = '\r'
									d.state = stateData
//line /usr/local/go/src/net/textproto/reader.go:374
			// _ = "end of CoverTab[34677]"

		case stateCR:
//line /usr/local/go/src/net/textproto/reader.go:376
			_go_fuzz_dep_.CoverTab[34678]++
									if c == '\n' {
//line /usr/local/go/src/net/textproto/reader.go:377
				_go_fuzz_dep_.CoverTab[34693]++
										d.state = stateBeginLine
										break
//line /usr/local/go/src/net/textproto/reader.go:379
				// _ = "end of CoverTab[34693]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:380
				_go_fuzz_dep_.CoverTab[34694]++
//line /usr/local/go/src/net/textproto/reader.go:380
				// _ = "end of CoverTab[34694]"
//line /usr/local/go/src/net/textproto/reader.go:380
			}
//line /usr/local/go/src/net/textproto/reader.go:380
			// _ = "end of CoverTab[34678]"
//line /usr/local/go/src/net/textproto/reader.go:380
			_go_fuzz_dep_.CoverTab[34679]++

									br.UnreadByte()
									c = '\r'
									d.state = stateData
//line /usr/local/go/src/net/textproto/reader.go:384
			// _ = "end of CoverTab[34679]"

		case stateData:
//line /usr/local/go/src/net/textproto/reader.go:386
			_go_fuzz_dep_.CoverTab[34680]++
									if c == '\r' {
//line /usr/local/go/src/net/textproto/reader.go:387
				_go_fuzz_dep_.CoverTab[34695]++
										d.state = stateCR
										continue
//line /usr/local/go/src/net/textproto/reader.go:389
				// _ = "end of CoverTab[34695]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:390
				_go_fuzz_dep_.CoverTab[34696]++
//line /usr/local/go/src/net/textproto/reader.go:390
				// _ = "end of CoverTab[34696]"
//line /usr/local/go/src/net/textproto/reader.go:390
			}
//line /usr/local/go/src/net/textproto/reader.go:390
			// _ = "end of CoverTab[34680]"
//line /usr/local/go/src/net/textproto/reader.go:390
			_go_fuzz_dep_.CoverTab[34681]++
									if c == '\n' {
//line /usr/local/go/src/net/textproto/reader.go:391
				_go_fuzz_dep_.CoverTab[34697]++
										d.state = stateBeginLine
//line /usr/local/go/src/net/textproto/reader.go:392
				// _ = "end of CoverTab[34697]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:393
				_go_fuzz_dep_.CoverTab[34698]++
//line /usr/local/go/src/net/textproto/reader.go:393
				// _ = "end of CoverTab[34698]"
//line /usr/local/go/src/net/textproto/reader.go:393
			}
//line /usr/local/go/src/net/textproto/reader.go:393
			// _ = "end of CoverTab[34681]"
//line /usr/local/go/src/net/textproto/reader.go:393
		default:
//line /usr/local/go/src/net/textproto/reader.go:393
			_go_fuzz_dep_.CoverTab[34682]++
//line /usr/local/go/src/net/textproto/reader.go:393
			// _ = "end of CoverTab[34682]"
		}
//line /usr/local/go/src/net/textproto/reader.go:394
		// _ = "end of CoverTab[34663]"
//line /usr/local/go/src/net/textproto/reader.go:394
		_go_fuzz_dep_.CoverTab[34664]++
								b[n] = c
								n++
//line /usr/local/go/src/net/textproto/reader.go:396
		// _ = "end of CoverTab[34664]"
	}
//line /usr/local/go/src/net/textproto/reader.go:397
	// _ = "end of CoverTab[34657]"
//line /usr/local/go/src/net/textproto/reader.go:397
	_go_fuzz_dep_.CoverTab[34658]++
							if err == nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:398
		_go_fuzz_dep_.CoverTab[34699]++
//line /usr/local/go/src/net/textproto/reader.go:398
		return d.state == stateEOF
//line /usr/local/go/src/net/textproto/reader.go:398
		// _ = "end of CoverTab[34699]"
//line /usr/local/go/src/net/textproto/reader.go:398
	}() {
//line /usr/local/go/src/net/textproto/reader.go:398
		_go_fuzz_dep_.CoverTab[34700]++
								err = io.EOF
//line /usr/local/go/src/net/textproto/reader.go:399
		// _ = "end of CoverTab[34700]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:400
		_go_fuzz_dep_.CoverTab[34701]++
//line /usr/local/go/src/net/textproto/reader.go:400
		// _ = "end of CoverTab[34701]"
//line /usr/local/go/src/net/textproto/reader.go:400
	}
//line /usr/local/go/src/net/textproto/reader.go:400
	// _ = "end of CoverTab[34658]"
//line /usr/local/go/src/net/textproto/reader.go:400
	_go_fuzz_dep_.CoverTab[34659]++
							if err != nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:401
		_go_fuzz_dep_.CoverTab[34702]++
//line /usr/local/go/src/net/textproto/reader.go:401
		return d.r.dot == d
//line /usr/local/go/src/net/textproto/reader.go:401
		// _ = "end of CoverTab[34702]"
//line /usr/local/go/src/net/textproto/reader.go:401
	}() {
//line /usr/local/go/src/net/textproto/reader.go:401
		_go_fuzz_dep_.CoverTab[34703]++
								d.r.dot = nil
//line /usr/local/go/src/net/textproto/reader.go:402
		// _ = "end of CoverTab[34703]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:403
		_go_fuzz_dep_.CoverTab[34704]++
//line /usr/local/go/src/net/textproto/reader.go:403
		// _ = "end of CoverTab[34704]"
//line /usr/local/go/src/net/textproto/reader.go:403
	}
//line /usr/local/go/src/net/textproto/reader.go:403
	// _ = "end of CoverTab[34659]"
//line /usr/local/go/src/net/textproto/reader.go:403
	_go_fuzz_dep_.CoverTab[34660]++
							return
//line /usr/local/go/src/net/textproto/reader.go:404
	// _ = "end of CoverTab[34660]"
}

// closeDot drains the current DotReader if any,
//line /usr/local/go/src/net/textproto/reader.go:407
// making sure that it reads until the ending dot line.
//line /usr/local/go/src/net/textproto/reader.go:409
func (r *Reader) closeDot() {
//line /usr/local/go/src/net/textproto/reader.go:409
	_go_fuzz_dep_.CoverTab[34705]++
							if r.dot == nil {
//line /usr/local/go/src/net/textproto/reader.go:410
		_go_fuzz_dep_.CoverTab[34707]++
								return
//line /usr/local/go/src/net/textproto/reader.go:411
		// _ = "end of CoverTab[34707]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:412
		_go_fuzz_dep_.CoverTab[34708]++
//line /usr/local/go/src/net/textproto/reader.go:412
		// _ = "end of CoverTab[34708]"
//line /usr/local/go/src/net/textproto/reader.go:412
	}
//line /usr/local/go/src/net/textproto/reader.go:412
	// _ = "end of CoverTab[34705]"
//line /usr/local/go/src/net/textproto/reader.go:412
	_go_fuzz_dep_.CoverTab[34706]++
							buf := make([]byte, 128)
							for r.dot != nil {
//line /usr/local/go/src/net/textproto/reader.go:414
		_go_fuzz_dep_.CoverTab[34709]++

//line /usr/local/go/src/net/textproto/reader.go:417
		r.dot.Read(buf)
//line /usr/local/go/src/net/textproto/reader.go:417
		// _ = "end of CoverTab[34709]"
	}
//line /usr/local/go/src/net/textproto/reader.go:418
	// _ = "end of CoverTab[34706]"
}

// ReadDotBytes reads a dot-encoding and returns the decoded data.
//line /usr/local/go/src/net/textproto/reader.go:421
//
//line /usr/local/go/src/net/textproto/reader.go:421
// See the documentation for the DotReader method for details about dot-encoding.
//line /usr/local/go/src/net/textproto/reader.go:424
func (r *Reader) ReadDotBytes() ([]byte, error) {
//line /usr/local/go/src/net/textproto/reader.go:424
	_go_fuzz_dep_.CoverTab[34710]++
							return io.ReadAll(r.DotReader())
//line /usr/local/go/src/net/textproto/reader.go:425
	// _ = "end of CoverTab[34710]"
}

// ReadDotLines reads a dot-encoding and returns a slice
//line /usr/local/go/src/net/textproto/reader.go:428
// containing the decoded lines, with the final \r\n or \n elided from each.
//line /usr/local/go/src/net/textproto/reader.go:428
//
//line /usr/local/go/src/net/textproto/reader.go:428
// See the documentation for the DotReader method for details about dot-encoding.
//line /usr/local/go/src/net/textproto/reader.go:432
func (r *Reader) ReadDotLines() ([]string, error) {
//line /usr/local/go/src/net/textproto/reader.go:432
	_go_fuzz_dep_.CoverTab[34711]++
	// We could use ReadDotBytes and then Split it,
	// but reading a line at a time avoids needing a
	// large contiguous block of memory and is simpler.
	var v []string
	var err error
	for {
//line /usr/local/go/src/net/textproto/reader.go:438
		_go_fuzz_dep_.CoverTab[34713]++
								var line string
								line, err = r.ReadLine()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:441
			_go_fuzz_dep_.CoverTab[34716]++
									if err == io.EOF {
//line /usr/local/go/src/net/textproto/reader.go:442
				_go_fuzz_dep_.CoverTab[34718]++
										err = io.ErrUnexpectedEOF
//line /usr/local/go/src/net/textproto/reader.go:443
				// _ = "end of CoverTab[34718]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:444
				_go_fuzz_dep_.CoverTab[34719]++
//line /usr/local/go/src/net/textproto/reader.go:444
				// _ = "end of CoverTab[34719]"
//line /usr/local/go/src/net/textproto/reader.go:444
			}
//line /usr/local/go/src/net/textproto/reader.go:444
			// _ = "end of CoverTab[34716]"
//line /usr/local/go/src/net/textproto/reader.go:444
			_go_fuzz_dep_.CoverTab[34717]++
									break
//line /usr/local/go/src/net/textproto/reader.go:445
			// _ = "end of CoverTab[34717]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:446
			_go_fuzz_dep_.CoverTab[34720]++
//line /usr/local/go/src/net/textproto/reader.go:446
			// _ = "end of CoverTab[34720]"
//line /usr/local/go/src/net/textproto/reader.go:446
		}
//line /usr/local/go/src/net/textproto/reader.go:446
		// _ = "end of CoverTab[34713]"
//line /usr/local/go/src/net/textproto/reader.go:446
		_go_fuzz_dep_.CoverTab[34714]++

//line /usr/local/go/src/net/textproto/reader.go:449
		if len(line) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:449
			_go_fuzz_dep_.CoverTab[34721]++
//line /usr/local/go/src/net/textproto/reader.go:449
			return line[0] == '.'
//line /usr/local/go/src/net/textproto/reader.go:449
			// _ = "end of CoverTab[34721]"
//line /usr/local/go/src/net/textproto/reader.go:449
		}() {
//line /usr/local/go/src/net/textproto/reader.go:449
			_go_fuzz_dep_.CoverTab[34722]++
									if len(line) == 1 {
//line /usr/local/go/src/net/textproto/reader.go:450
				_go_fuzz_dep_.CoverTab[34724]++
										break
//line /usr/local/go/src/net/textproto/reader.go:451
				// _ = "end of CoverTab[34724]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:452
				_go_fuzz_dep_.CoverTab[34725]++
//line /usr/local/go/src/net/textproto/reader.go:452
				// _ = "end of CoverTab[34725]"
//line /usr/local/go/src/net/textproto/reader.go:452
			}
//line /usr/local/go/src/net/textproto/reader.go:452
			// _ = "end of CoverTab[34722]"
//line /usr/local/go/src/net/textproto/reader.go:452
			_go_fuzz_dep_.CoverTab[34723]++
									line = line[1:]
//line /usr/local/go/src/net/textproto/reader.go:453
			// _ = "end of CoverTab[34723]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:454
			_go_fuzz_dep_.CoverTab[34726]++
//line /usr/local/go/src/net/textproto/reader.go:454
			// _ = "end of CoverTab[34726]"
//line /usr/local/go/src/net/textproto/reader.go:454
		}
//line /usr/local/go/src/net/textproto/reader.go:454
		// _ = "end of CoverTab[34714]"
//line /usr/local/go/src/net/textproto/reader.go:454
		_go_fuzz_dep_.CoverTab[34715]++
								v = append(v, line)
//line /usr/local/go/src/net/textproto/reader.go:455
		// _ = "end of CoverTab[34715]"
	}
//line /usr/local/go/src/net/textproto/reader.go:456
	// _ = "end of CoverTab[34711]"
//line /usr/local/go/src/net/textproto/reader.go:456
	_go_fuzz_dep_.CoverTab[34712]++
							return v, err
//line /usr/local/go/src/net/textproto/reader.go:457
	// _ = "end of CoverTab[34712]"
}

var colon = []byte(":")

// ReadMIMEHeader reads a MIME-style header from r.
//line /usr/local/go/src/net/textproto/reader.go:462
// The header is a sequence of possibly continued Key: Value lines
//line /usr/local/go/src/net/textproto/reader.go:462
// ending in a blank line.
//line /usr/local/go/src/net/textproto/reader.go:462
// The returned map m maps CanonicalMIMEHeaderKey(key) to a
//line /usr/local/go/src/net/textproto/reader.go:462
// sequence of values in the same order encountered in the input.
//line /usr/local/go/src/net/textproto/reader.go:462
//
//line /usr/local/go/src/net/textproto/reader.go:462
// For example, consider this input:
//line /usr/local/go/src/net/textproto/reader.go:462
//
//line /usr/local/go/src/net/textproto/reader.go:462
//	My-Key: Value 1
//line /usr/local/go/src/net/textproto/reader.go:462
//	Long-Key: Even
//line /usr/local/go/src/net/textproto/reader.go:462
//	       Longer Value
//line /usr/local/go/src/net/textproto/reader.go:462
//	My-Key: Value 2
//line /usr/local/go/src/net/textproto/reader.go:462
//
//line /usr/local/go/src/net/textproto/reader.go:462
// Given that input, ReadMIMEHeader returns the map:
//line /usr/local/go/src/net/textproto/reader.go:462
//
//line /usr/local/go/src/net/textproto/reader.go:462
//	map[string][]string{
//line /usr/local/go/src/net/textproto/reader.go:462
//		"My-Key": {"Value 1", "Value 2"},
//line /usr/local/go/src/net/textproto/reader.go:462
//		"Long-Key": {"Even Longer Value"},
//line /usr/local/go/src/net/textproto/reader.go:462
//	}
//line /usr/local/go/src/net/textproto/reader.go:481
func (r *Reader) ReadMIMEHeader() (MIMEHeader, error) {
//line /usr/local/go/src/net/textproto/reader.go:481
	_go_fuzz_dep_.CoverTab[34727]++
							return readMIMEHeader(r, math.MaxInt64, math.MaxInt64)
//line /usr/local/go/src/net/textproto/reader.go:482
	// _ = "end of CoverTab[34727]"
}

// readMIMEHeader is a version of ReadMIMEHeader which takes a limit on the header size.
//line /usr/local/go/src/net/textproto/reader.go:485
// It is called by the mime/multipart package.
//line /usr/local/go/src/net/textproto/reader.go:487
func readMIMEHeader(r *Reader, maxMemory, maxHeaders int64) (MIMEHeader, error) {
//line /usr/local/go/src/net/textproto/reader.go:487
	_go_fuzz_dep_.CoverTab[34728]++
	// Avoid lots of small slice allocations later by allocating one
	// large one ahead of time which we'll cut up into smaller
	// slices. If this isn't big enough later, we allocate small ones.
	var strs []string
	hint := r.upcomingHeaderKeys()
	if hint > 0 {
//line /usr/local/go/src/net/textproto/reader.go:493
		_go_fuzz_dep_.CoverTab[34731]++
								if hint > 1000 {
//line /usr/local/go/src/net/textproto/reader.go:494
			_go_fuzz_dep_.CoverTab[34733]++
									hint = 1000
//line /usr/local/go/src/net/textproto/reader.go:495
			// _ = "end of CoverTab[34733]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:496
			_go_fuzz_dep_.CoverTab[34734]++
//line /usr/local/go/src/net/textproto/reader.go:496
			// _ = "end of CoverTab[34734]"
//line /usr/local/go/src/net/textproto/reader.go:496
		}
//line /usr/local/go/src/net/textproto/reader.go:496
		// _ = "end of CoverTab[34731]"
//line /usr/local/go/src/net/textproto/reader.go:496
		_go_fuzz_dep_.CoverTab[34732]++
								strs = make([]string, hint)
//line /usr/local/go/src/net/textproto/reader.go:497
		// _ = "end of CoverTab[34732]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:498
		_go_fuzz_dep_.CoverTab[34735]++
//line /usr/local/go/src/net/textproto/reader.go:498
		// _ = "end of CoverTab[34735]"
//line /usr/local/go/src/net/textproto/reader.go:498
	}
//line /usr/local/go/src/net/textproto/reader.go:498
	// _ = "end of CoverTab[34728]"
//line /usr/local/go/src/net/textproto/reader.go:498
	_go_fuzz_dep_.CoverTab[34729]++

							m := make(MIMEHeader, hint)

//line /usr/local/go/src/net/textproto/reader.go:505
	maxMemory -= 400
							const mapEntryOverhead = 200

//line /usr/local/go/src/net/textproto/reader.go:509
	if buf, err := r.R.Peek(1); err == nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:509
		_go_fuzz_dep_.CoverTab[34736]++
//line /usr/local/go/src/net/textproto/reader.go:509
		return (buf[0] == ' ' || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:509
			_go_fuzz_dep_.CoverTab[34737]++
//line /usr/local/go/src/net/textproto/reader.go:509
			return buf[0] == '\t'
//line /usr/local/go/src/net/textproto/reader.go:509
			// _ = "end of CoverTab[34737]"
//line /usr/local/go/src/net/textproto/reader.go:509
		}())
//line /usr/local/go/src/net/textproto/reader.go:509
		// _ = "end of CoverTab[34736]"
//line /usr/local/go/src/net/textproto/reader.go:509
	}() {
//line /usr/local/go/src/net/textproto/reader.go:509
		_go_fuzz_dep_.CoverTab[34738]++
								line, err := r.readLineSlice()
								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:511
			_go_fuzz_dep_.CoverTab[34740]++
									return m, err
//line /usr/local/go/src/net/textproto/reader.go:512
			// _ = "end of CoverTab[34740]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:513
			_go_fuzz_dep_.CoverTab[34741]++
//line /usr/local/go/src/net/textproto/reader.go:513
			// _ = "end of CoverTab[34741]"
//line /usr/local/go/src/net/textproto/reader.go:513
		}
//line /usr/local/go/src/net/textproto/reader.go:513
		// _ = "end of CoverTab[34738]"
//line /usr/local/go/src/net/textproto/reader.go:513
		_go_fuzz_dep_.CoverTab[34739]++
								return m, ProtocolError("malformed MIME header initial line: " + string(line))
//line /usr/local/go/src/net/textproto/reader.go:514
		// _ = "end of CoverTab[34739]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:515
		_go_fuzz_dep_.CoverTab[34742]++
//line /usr/local/go/src/net/textproto/reader.go:515
		// _ = "end of CoverTab[34742]"
//line /usr/local/go/src/net/textproto/reader.go:515
	}
//line /usr/local/go/src/net/textproto/reader.go:515
	// _ = "end of CoverTab[34729]"
//line /usr/local/go/src/net/textproto/reader.go:515
	_go_fuzz_dep_.CoverTab[34730]++

							for {
//line /usr/local/go/src/net/textproto/reader.go:517
		_go_fuzz_dep_.CoverTab[34743]++
								kv, err := r.readContinuedLineSlice(mustHaveFieldNameColon)
								if len(kv) == 0 {
//line /usr/local/go/src/net/textproto/reader.go:519
			_go_fuzz_dep_.CoverTab[34753]++
									return m, err
//line /usr/local/go/src/net/textproto/reader.go:520
			// _ = "end of CoverTab[34753]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:521
			_go_fuzz_dep_.CoverTab[34754]++
//line /usr/local/go/src/net/textproto/reader.go:521
			// _ = "end of CoverTab[34754]"
//line /usr/local/go/src/net/textproto/reader.go:521
		}
//line /usr/local/go/src/net/textproto/reader.go:521
		// _ = "end of CoverTab[34743]"
//line /usr/local/go/src/net/textproto/reader.go:521
		_go_fuzz_dep_.CoverTab[34744]++

//line /usr/local/go/src/net/textproto/reader.go:524
		k, v, ok := bytes.Cut(kv, colon)
		if !ok {
//line /usr/local/go/src/net/textproto/reader.go:525
			_go_fuzz_dep_.CoverTab[34755]++
									return m, ProtocolError("malformed MIME header line: " + string(kv))
//line /usr/local/go/src/net/textproto/reader.go:526
			// _ = "end of CoverTab[34755]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:527
			_go_fuzz_dep_.CoverTab[34756]++
//line /usr/local/go/src/net/textproto/reader.go:527
			// _ = "end of CoverTab[34756]"
//line /usr/local/go/src/net/textproto/reader.go:527
		}
//line /usr/local/go/src/net/textproto/reader.go:527
		// _ = "end of CoverTab[34744]"
//line /usr/local/go/src/net/textproto/reader.go:527
		_go_fuzz_dep_.CoverTab[34745]++
								key, ok := canonicalMIMEHeaderKey(k)
								if !ok {
//line /usr/local/go/src/net/textproto/reader.go:529
			_go_fuzz_dep_.CoverTab[34757]++
									return m, ProtocolError("malformed MIME header line: " + string(kv))
//line /usr/local/go/src/net/textproto/reader.go:530
			// _ = "end of CoverTab[34757]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:531
			_go_fuzz_dep_.CoverTab[34758]++
//line /usr/local/go/src/net/textproto/reader.go:531
			// _ = "end of CoverTab[34758]"
//line /usr/local/go/src/net/textproto/reader.go:531
		}
//line /usr/local/go/src/net/textproto/reader.go:531
		// _ = "end of CoverTab[34745]"
//line /usr/local/go/src/net/textproto/reader.go:531
		_go_fuzz_dep_.CoverTab[34746]++
								for _, c := range v {
//line /usr/local/go/src/net/textproto/reader.go:532
			_go_fuzz_dep_.CoverTab[34759]++
									if !validHeaderValueByte(c) {
//line /usr/local/go/src/net/textproto/reader.go:533
				_go_fuzz_dep_.CoverTab[34760]++
										return m, ProtocolError("malformed MIME header line: " + string(kv))
//line /usr/local/go/src/net/textproto/reader.go:534
				// _ = "end of CoverTab[34760]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:535
				_go_fuzz_dep_.CoverTab[34761]++
//line /usr/local/go/src/net/textproto/reader.go:535
				// _ = "end of CoverTab[34761]"
//line /usr/local/go/src/net/textproto/reader.go:535
			}
//line /usr/local/go/src/net/textproto/reader.go:535
			// _ = "end of CoverTab[34759]"
		}
//line /usr/local/go/src/net/textproto/reader.go:536
		// _ = "end of CoverTab[34746]"
//line /usr/local/go/src/net/textproto/reader.go:536
		_go_fuzz_dep_.CoverTab[34747]++

//line /usr/local/go/src/net/textproto/reader.go:541
		if key == "" {
//line /usr/local/go/src/net/textproto/reader.go:541
			_go_fuzz_dep_.CoverTab[34762]++
									continue
//line /usr/local/go/src/net/textproto/reader.go:542
			// _ = "end of CoverTab[34762]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:543
			_go_fuzz_dep_.CoverTab[34763]++
//line /usr/local/go/src/net/textproto/reader.go:543
			// _ = "end of CoverTab[34763]"
//line /usr/local/go/src/net/textproto/reader.go:543
		}
//line /usr/local/go/src/net/textproto/reader.go:543
		// _ = "end of CoverTab[34747]"
//line /usr/local/go/src/net/textproto/reader.go:543
		_go_fuzz_dep_.CoverTab[34748]++

								maxHeaders--
								if maxHeaders < 0 {
//line /usr/local/go/src/net/textproto/reader.go:546
			_go_fuzz_dep_.CoverTab[34764]++
									return nil, errors.New("message too large")
//line /usr/local/go/src/net/textproto/reader.go:547
			// _ = "end of CoverTab[34764]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:548
			_go_fuzz_dep_.CoverTab[34765]++
//line /usr/local/go/src/net/textproto/reader.go:548
			// _ = "end of CoverTab[34765]"
//line /usr/local/go/src/net/textproto/reader.go:548
		}
//line /usr/local/go/src/net/textproto/reader.go:548
		// _ = "end of CoverTab[34748]"
//line /usr/local/go/src/net/textproto/reader.go:548
		_go_fuzz_dep_.CoverTab[34749]++

//line /usr/local/go/src/net/textproto/reader.go:551
		value := string(bytes.TrimLeft(v, " \t"))

		vv := m[key]
		if vv == nil {
//line /usr/local/go/src/net/textproto/reader.go:554
			_go_fuzz_dep_.CoverTab[34766]++
									maxMemory -= int64(len(key))
									maxMemory -= mapEntryOverhead
//line /usr/local/go/src/net/textproto/reader.go:556
			// _ = "end of CoverTab[34766]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:557
			_go_fuzz_dep_.CoverTab[34767]++
//line /usr/local/go/src/net/textproto/reader.go:557
			// _ = "end of CoverTab[34767]"
//line /usr/local/go/src/net/textproto/reader.go:557
		}
//line /usr/local/go/src/net/textproto/reader.go:557
		// _ = "end of CoverTab[34749]"
//line /usr/local/go/src/net/textproto/reader.go:557
		_go_fuzz_dep_.CoverTab[34750]++
								maxMemory -= int64(len(value))
								if maxMemory < 0 {
//line /usr/local/go/src/net/textproto/reader.go:559
			_go_fuzz_dep_.CoverTab[34768]++

//line /usr/local/go/src/net/textproto/reader.go:562
			return m, errors.New("message too large")
//line /usr/local/go/src/net/textproto/reader.go:562
			// _ = "end of CoverTab[34768]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:563
			_go_fuzz_dep_.CoverTab[34769]++
//line /usr/local/go/src/net/textproto/reader.go:563
			// _ = "end of CoverTab[34769]"
//line /usr/local/go/src/net/textproto/reader.go:563
		}
//line /usr/local/go/src/net/textproto/reader.go:563
		// _ = "end of CoverTab[34750]"
//line /usr/local/go/src/net/textproto/reader.go:563
		_go_fuzz_dep_.CoverTab[34751]++
								if vv == nil && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:564
			_go_fuzz_dep_.CoverTab[34770]++
//line /usr/local/go/src/net/textproto/reader.go:564
			return len(strs) > 0
//line /usr/local/go/src/net/textproto/reader.go:564
			// _ = "end of CoverTab[34770]"
//line /usr/local/go/src/net/textproto/reader.go:564
		}() {
//line /usr/local/go/src/net/textproto/reader.go:564
			_go_fuzz_dep_.CoverTab[34771]++

//line /usr/local/go/src/net/textproto/reader.go:569
			vv, strs = strs[:1:1], strs[1:]
									vv[0] = value
									m[key] = vv
//line /usr/local/go/src/net/textproto/reader.go:571
			// _ = "end of CoverTab[34771]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:572
			_go_fuzz_dep_.CoverTab[34772]++
									m[key] = append(vv, value)
//line /usr/local/go/src/net/textproto/reader.go:573
			// _ = "end of CoverTab[34772]"
		}
//line /usr/local/go/src/net/textproto/reader.go:574
		// _ = "end of CoverTab[34751]"
//line /usr/local/go/src/net/textproto/reader.go:574
		_go_fuzz_dep_.CoverTab[34752]++

								if err != nil {
//line /usr/local/go/src/net/textproto/reader.go:576
			_go_fuzz_dep_.CoverTab[34773]++
									return m, err
//line /usr/local/go/src/net/textproto/reader.go:577
			// _ = "end of CoverTab[34773]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:578
			_go_fuzz_dep_.CoverTab[34774]++
//line /usr/local/go/src/net/textproto/reader.go:578
			// _ = "end of CoverTab[34774]"
//line /usr/local/go/src/net/textproto/reader.go:578
		}
//line /usr/local/go/src/net/textproto/reader.go:578
		// _ = "end of CoverTab[34752]"
	}
//line /usr/local/go/src/net/textproto/reader.go:579
	// _ = "end of CoverTab[34730]"
}

// noValidation is a no-op validation func for readContinuedLineSlice
//line /usr/local/go/src/net/textproto/reader.go:582
// that permits any lines.
//line /usr/local/go/src/net/textproto/reader.go:584
func noValidation(_ []byte) error {
//line /usr/local/go/src/net/textproto/reader.go:584
	_go_fuzz_dep_.CoverTab[34775]++
//line /usr/local/go/src/net/textproto/reader.go:584
	return nil
//line /usr/local/go/src/net/textproto/reader.go:584
	// _ = "end of CoverTab[34775]"
//line /usr/local/go/src/net/textproto/reader.go:584
}

// mustHaveFieldNameColon ensures that, per RFC 7230, the
//line /usr/local/go/src/net/textproto/reader.go:586
// field-name is on a single line, so the first line must
//line /usr/local/go/src/net/textproto/reader.go:586
// contain a colon.
//line /usr/local/go/src/net/textproto/reader.go:589
func mustHaveFieldNameColon(line []byte) error {
//line /usr/local/go/src/net/textproto/reader.go:589
	_go_fuzz_dep_.CoverTab[34776]++
							if bytes.IndexByte(line, ':') < 0 {
//line /usr/local/go/src/net/textproto/reader.go:590
		_go_fuzz_dep_.CoverTab[34778]++
								return ProtocolError(fmt.Sprintf("malformed MIME header: missing colon: %q", line))
//line /usr/local/go/src/net/textproto/reader.go:591
		// _ = "end of CoverTab[34778]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:592
		_go_fuzz_dep_.CoverTab[34779]++
//line /usr/local/go/src/net/textproto/reader.go:592
		// _ = "end of CoverTab[34779]"
//line /usr/local/go/src/net/textproto/reader.go:592
	}
//line /usr/local/go/src/net/textproto/reader.go:592
	// _ = "end of CoverTab[34776]"
//line /usr/local/go/src/net/textproto/reader.go:592
	_go_fuzz_dep_.CoverTab[34777]++
							return nil
//line /usr/local/go/src/net/textproto/reader.go:593
	// _ = "end of CoverTab[34777]"
}

var nl = []byte("\n")

// upcomingHeaderKeys returns an approximation of the number of keys
//line /usr/local/go/src/net/textproto/reader.go:598
// that will be in this header. If it gets confused, it returns 0.
//line /usr/local/go/src/net/textproto/reader.go:600
func (r *Reader) upcomingHeaderKeys() (n int) {
//line /usr/local/go/src/net/textproto/reader.go:600
	_go_fuzz_dep_.CoverTab[34780]++

							r.R.Peek(1)
							s := r.R.Buffered()
							if s == 0 {
//line /usr/local/go/src/net/textproto/reader.go:604
		_go_fuzz_dep_.CoverTab[34783]++
								return
//line /usr/local/go/src/net/textproto/reader.go:605
		// _ = "end of CoverTab[34783]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:606
		_go_fuzz_dep_.CoverTab[34784]++
//line /usr/local/go/src/net/textproto/reader.go:606
		// _ = "end of CoverTab[34784]"
//line /usr/local/go/src/net/textproto/reader.go:606
	}
//line /usr/local/go/src/net/textproto/reader.go:606
	// _ = "end of CoverTab[34780]"
//line /usr/local/go/src/net/textproto/reader.go:606
	_go_fuzz_dep_.CoverTab[34781]++
							peek, _ := r.R.Peek(s)
							for len(peek) > 0 && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:608
		_go_fuzz_dep_.CoverTab[34785]++
//line /usr/local/go/src/net/textproto/reader.go:608
		return n < 1000
//line /usr/local/go/src/net/textproto/reader.go:608
		// _ = "end of CoverTab[34785]"
//line /usr/local/go/src/net/textproto/reader.go:608
	}() {
//line /usr/local/go/src/net/textproto/reader.go:608
		_go_fuzz_dep_.CoverTab[34786]++
								var line []byte
								line, peek, _ = bytes.Cut(peek, nl)
								if len(line) == 0 || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:611
			_go_fuzz_dep_.CoverTab[34789]++
//line /usr/local/go/src/net/textproto/reader.go:611
			return (len(line) == 1 && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:611
				_go_fuzz_dep_.CoverTab[34790]++
//line /usr/local/go/src/net/textproto/reader.go:611
				return line[0] == '\r'
//line /usr/local/go/src/net/textproto/reader.go:611
				// _ = "end of CoverTab[34790]"
//line /usr/local/go/src/net/textproto/reader.go:611
			}())
//line /usr/local/go/src/net/textproto/reader.go:611
			// _ = "end of CoverTab[34789]"
//line /usr/local/go/src/net/textproto/reader.go:611
		}() {
//line /usr/local/go/src/net/textproto/reader.go:611
			_go_fuzz_dep_.CoverTab[34791]++

									break
//line /usr/local/go/src/net/textproto/reader.go:613
			// _ = "end of CoverTab[34791]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:614
			_go_fuzz_dep_.CoverTab[34792]++
//line /usr/local/go/src/net/textproto/reader.go:614
			// _ = "end of CoverTab[34792]"
//line /usr/local/go/src/net/textproto/reader.go:614
		}
//line /usr/local/go/src/net/textproto/reader.go:614
		// _ = "end of CoverTab[34786]"
//line /usr/local/go/src/net/textproto/reader.go:614
		_go_fuzz_dep_.CoverTab[34787]++
								if line[0] == ' ' || func() bool {
//line /usr/local/go/src/net/textproto/reader.go:615
			_go_fuzz_dep_.CoverTab[34793]++
//line /usr/local/go/src/net/textproto/reader.go:615
			return line[0] == '\t'
//line /usr/local/go/src/net/textproto/reader.go:615
			// _ = "end of CoverTab[34793]"
//line /usr/local/go/src/net/textproto/reader.go:615
		}() {
//line /usr/local/go/src/net/textproto/reader.go:615
			_go_fuzz_dep_.CoverTab[34794]++

									continue
//line /usr/local/go/src/net/textproto/reader.go:617
			// _ = "end of CoverTab[34794]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:618
			_go_fuzz_dep_.CoverTab[34795]++
//line /usr/local/go/src/net/textproto/reader.go:618
			// _ = "end of CoverTab[34795]"
//line /usr/local/go/src/net/textproto/reader.go:618
		}
//line /usr/local/go/src/net/textproto/reader.go:618
		// _ = "end of CoverTab[34787]"
//line /usr/local/go/src/net/textproto/reader.go:618
		_go_fuzz_dep_.CoverTab[34788]++
								n++
//line /usr/local/go/src/net/textproto/reader.go:619
		// _ = "end of CoverTab[34788]"
	}
//line /usr/local/go/src/net/textproto/reader.go:620
	// _ = "end of CoverTab[34781]"
//line /usr/local/go/src/net/textproto/reader.go:620
	_go_fuzz_dep_.CoverTab[34782]++
							return n
//line /usr/local/go/src/net/textproto/reader.go:621
	// _ = "end of CoverTab[34782]"
}

// CanonicalMIMEHeaderKey returns the canonical format of the
//line /usr/local/go/src/net/textproto/reader.go:624
// MIME header key s. The canonicalization converts the first
//line /usr/local/go/src/net/textproto/reader.go:624
// letter and any letter following a hyphen to upper case;
//line /usr/local/go/src/net/textproto/reader.go:624
// the rest are converted to lowercase. For example, the
//line /usr/local/go/src/net/textproto/reader.go:624
// canonical key for "accept-encoding" is "Accept-Encoding".
//line /usr/local/go/src/net/textproto/reader.go:624
// MIME header keys are assumed to be ASCII only.
//line /usr/local/go/src/net/textproto/reader.go:624
// If s contains a space or invalid header field bytes, it is
//line /usr/local/go/src/net/textproto/reader.go:624
// returned without modifications.
//line /usr/local/go/src/net/textproto/reader.go:632
func CanonicalMIMEHeaderKey(s string) string {
//line /usr/local/go/src/net/textproto/reader.go:632
	_go_fuzz_dep_.CoverTab[34796]++

							upper := true
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/net/textproto/reader.go:635
		_go_fuzz_dep_.CoverTab[34798]++
								c := s[i]
								if !validHeaderFieldByte(c) {
//line /usr/local/go/src/net/textproto/reader.go:637
			_go_fuzz_dep_.CoverTab[34802]++
									return s
//line /usr/local/go/src/net/textproto/reader.go:638
			// _ = "end of CoverTab[34802]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:639
			_go_fuzz_dep_.CoverTab[34803]++
//line /usr/local/go/src/net/textproto/reader.go:639
			// _ = "end of CoverTab[34803]"
//line /usr/local/go/src/net/textproto/reader.go:639
		}
//line /usr/local/go/src/net/textproto/reader.go:639
		// _ = "end of CoverTab[34798]"
//line /usr/local/go/src/net/textproto/reader.go:639
		_go_fuzz_dep_.CoverTab[34799]++
								if upper && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:640
			_go_fuzz_dep_.CoverTab[34804]++
//line /usr/local/go/src/net/textproto/reader.go:640
			return 'a' <= c
//line /usr/local/go/src/net/textproto/reader.go:640
			// _ = "end of CoverTab[34804]"
//line /usr/local/go/src/net/textproto/reader.go:640
		}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:640
			_go_fuzz_dep_.CoverTab[34805]++
//line /usr/local/go/src/net/textproto/reader.go:640
			return c <= 'z'
//line /usr/local/go/src/net/textproto/reader.go:640
			// _ = "end of CoverTab[34805]"
//line /usr/local/go/src/net/textproto/reader.go:640
		}() {
//line /usr/local/go/src/net/textproto/reader.go:640
			_go_fuzz_dep_.CoverTab[34806]++
									s, _ = canonicalMIMEHeaderKey([]byte(s))
									return s
//line /usr/local/go/src/net/textproto/reader.go:642
			// _ = "end of CoverTab[34806]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:643
			_go_fuzz_dep_.CoverTab[34807]++
//line /usr/local/go/src/net/textproto/reader.go:643
			// _ = "end of CoverTab[34807]"
//line /usr/local/go/src/net/textproto/reader.go:643
		}
//line /usr/local/go/src/net/textproto/reader.go:643
		// _ = "end of CoverTab[34799]"
//line /usr/local/go/src/net/textproto/reader.go:643
		_go_fuzz_dep_.CoverTab[34800]++
								if !upper && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:644
			_go_fuzz_dep_.CoverTab[34808]++
//line /usr/local/go/src/net/textproto/reader.go:644
			return 'A' <= c
//line /usr/local/go/src/net/textproto/reader.go:644
			// _ = "end of CoverTab[34808]"
//line /usr/local/go/src/net/textproto/reader.go:644
		}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:644
			_go_fuzz_dep_.CoverTab[34809]++
//line /usr/local/go/src/net/textproto/reader.go:644
			return c <= 'Z'
//line /usr/local/go/src/net/textproto/reader.go:644
			// _ = "end of CoverTab[34809]"
//line /usr/local/go/src/net/textproto/reader.go:644
		}() {
//line /usr/local/go/src/net/textproto/reader.go:644
			_go_fuzz_dep_.CoverTab[34810]++
									s, _ = canonicalMIMEHeaderKey([]byte(s))
									return s
//line /usr/local/go/src/net/textproto/reader.go:646
			// _ = "end of CoverTab[34810]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:647
			_go_fuzz_dep_.CoverTab[34811]++
//line /usr/local/go/src/net/textproto/reader.go:647
			// _ = "end of CoverTab[34811]"
//line /usr/local/go/src/net/textproto/reader.go:647
		}
//line /usr/local/go/src/net/textproto/reader.go:647
		// _ = "end of CoverTab[34800]"
//line /usr/local/go/src/net/textproto/reader.go:647
		_go_fuzz_dep_.CoverTab[34801]++
								upper = c == '-'
//line /usr/local/go/src/net/textproto/reader.go:648
		// _ = "end of CoverTab[34801]"
	}
//line /usr/local/go/src/net/textproto/reader.go:649
	// _ = "end of CoverTab[34796]"
//line /usr/local/go/src/net/textproto/reader.go:649
	_go_fuzz_dep_.CoverTab[34797]++
							return s
//line /usr/local/go/src/net/textproto/reader.go:650
	// _ = "end of CoverTab[34797]"
}

const toLower = 'a' - 'A'

// validHeaderFieldByte reports whether c is a valid byte in a header
//line /usr/local/go/src/net/textproto/reader.go:655
// field name. RFC 7230 says:
//line /usr/local/go/src/net/textproto/reader.go:655
//
//line /usr/local/go/src/net/textproto/reader.go:655
//	header-field   = field-name ":" OWS field-value OWS
//line /usr/local/go/src/net/textproto/reader.go:655
//	field-name     = token
//line /usr/local/go/src/net/textproto/reader.go:655
//	tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//line /usr/local/go/src/net/textproto/reader.go:655
//	        "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
//line /usr/local/go/src/net/textproto/reader.go:655
//	token = 1*tchar
//line /usr/local/go/src/net/textproto/reader.go:663
func validHeaderFieldByte(c byte) bool {
//line /usr/local/go/src/net/textproto/reader.go:663
	_go_fuzz_dep_.CoverTab[34812]++
	// mask is a 128-bit bitmap with 1s for allowed bytes,
	// so that the byte c can be tested with a shift and an and.
	// If c >= 128, then 1<<c and 1<<(c-64) will both be zero,
	// and this function will return false.
	const mask = 0 |
		(1<<(10)-1)<<'0' |
		(1<<(26)-1)<<'a' |
		(1<<(26)-1)<<'A' |
		1<<'!' |
		1<<'#' |
		1<<'$' |
		1<<'%' |
		1<<'&' |
		1<<'\'' |
		1<<'*' |
		1<<'+' |
		1<<'-' |
		1<<'.' |
		1<<'^' |
		1<<'_' |
		1<<'`' |
		1<<'|' |
		1<<'~'
	return ((uint64(1)<<c)&(mask&(1<<64-1)) |
		(uint64(1)<<(c-64))&(mask>>64)) != 0
//line /usr/local/go/src/net/textproto/reader.go:688
	// _ = "end of CoverTab[34812]"
}

// validHeaderValueByte reports whether c is a valid byte in a header
//line /usr/local/go/src/net/textproto/reader.go:691
// field value. RFC 7230 says:
//line /usr/local/go/src/net/textproto/reader.go:691
//
//line /usr/local/go/src/net/textproto/reader.go:691
//	field-content  = field-vchar [ 1*( SP / HTAB ) field-vchar ]
//line /usr/local/go/src/net/textproto/reader.go:691
//	field-vchar    = VCHAR / obs-text
//line /usr/local/go/src/net/textproto/reader.go:691
//	obs-text       = %x80-FF
//line /usr/local/go/src/net/textproto/reader.go:691
//
//line /usr/local/go/src/net/textproto/reader.go:691
// RFC 5234 says:
//line /usr/local/go/src/net/textproto/reader.go:691
//
//line /usr/local/go/src/net/textproto/reader.go:691
//	HTAB           =  %x09
//line /usr/local/go/src/net/textproto/reader.go:691
//	SP             =  %x20
//line /usr/local/go/src/net/textproto/reader.go:691
//	VCHAR          =  %x21-7E
//line /usr/local/go/src/net/textproto/reader.go:703
func validHeaderValueByte(c byte) bool {
//line /usr/local/go/src/net/textproto/reader.go:703
	_go_fuzz_dep_.CoverTab[34813]++
	// mask is a 128-bit bitmap with 1s for allowed bytes,
	// so that the byte c can be tested with a shift and an and.
	// If c >= 128, then 1<<c and 1<<(c-64) will both be zero.
	// Since this is the obs-text range, we invert the mask to
	// create a bitmap with 1s for disallowed bytes.
	const mask = 0 |
		(1<<(0x7f-0x21)-1)<<0x21 |
		1<<0x20 |
		1<<0x09	// HTAB: %x09
	return ((uint64(1)<<c)&^(mask&(1<<64-1)) |
		(uint64(1)<<(c-64))&^(mask>>64)) == 0
//line /usr/local/go/src/net/textproto/reader.go:714
	// _ = "end of CoverTab[34813]"
}

// canonicalMIMEHeaderKey is like CanonicalMIMEHeaderKey but is
//line /usr/local/go/src/net/textproto/reader.go:717
// allowed to mutate the provided byte slice before returning the
//line /usr/local/go/src/net/textproto/reader.go:717
// string.
//line /usr/local/go/src/net/textproto/reader.go:717
//
//line /usr/local/go/src/net/textproto/reader.go:717
// For invalid inputs (if a contains spaces or non-token bytes), a
//line /usr/local/go/src/net/textproto/reader.go:717
// is unchanged and a string copy is returned.
//line /usr/local/go/src/net/textproto/reader.go:717
//
//line /usr/local/go/src/net/textproto/reader.go:717
// ok is true if the header key contains only valid characters and spaces.
//line /usr/local/go/src/net/textproto/reader.go:717
// ReadMIMEHeader accepts header keys containing spaces, but does not
//line /usr/local/go/src/net/textproto/reader.go:717
// canonicalize them.
//line /usr/local/go/src/net/textproto/reader.go:727
func canonicalMIMEHeaderKey(a []byte) (_ string, ok bool) {
//line /usr/local/go/src/net/textproto/reader.go:727
	_go_fuzz_dep_.CoverTab[34814]++

							noCanon := false
							for _, c := range a {
//line /usr/local/go/src/net/textproto/reader.go:730
		_go_fuzz_dep_.CoverTab[34819]++
								if validHeaderFieldByte(c) {
//line /usr/local/go/src/net/textproto/reader.go:731
			_go_fuzz_dep_.CoverTab[34822]++
									continue
//line /usr/local/go/src/net/textproto/reader.go:732
			// _ = "end of CoverTab[34822]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:733
			_go_fuzz_dep_.CoverTab[34823]++
//line /usr/local/go/src/net/textproto/reader.go:733
			// _ = "end of CoverTab[34823]"
//line /usr/local/go/src/net/textproto/reader.go:733
		}
//line /usr/local/go/src/net/textproto/reader.go:733
		// _ = "end of CoverTab[34819]"
//line /usr/local/go/src/net/textproto/reader.go:733
		_go_fuzz_dep_.CoverTab[34820]++

								if c == ' ' {
//line /usr/local/go/src/net/textproto/reader.go:735
			_go_fuzz_dep_.CoverTab[34824]++

//line /usr/local/go/src/net/textproto/reader.go:739
			noCanon = true
									continue
//line /usr/local/go/src/net/textproto/reader.go:740
			// _ = "end of CoverTab[34824]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:741
			_go_fuzz_dep_.CoverTab[34825]++
//line /usr/local/go/src/net/textproto/reader.go:741
			// _ = "end of CoverTab[34825]"
//line /usr/local/go/src/net/textproto/reader.go:741
		}
//line /usr/local/go/src/net/textproto/reader.go:741
		// _ = "end of CoverTab[34820]"
//line /usr/local/go/src/net/textproto/reader.go:741
		_go_fuzz_dep_.CoverTab[34821]++
								return string(a), false
//line /usr/local/go/src/net/textproto/reader.go:742
		// _ = "end of CoverTab[34821]"
	}
//line /usr/local/go/src/net/textproto/reader.go:743
	// _ = "end of CoverTab[34814]"
//line /usr/local/go/src/net/textproto/reader.go:743
	_go_fuzz_dep_.CoverTab[34815]++
							if noCanon {
//line /usr/local/go/src/net/textproto/reader.go:744
		_go_fuzz_dep_.CoverTab[34826]++
								return string(a), true
//line /usr/local/go/src/net/textproto/reader.go:745
		// _ = "end of CoverTab[34826]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:746
		_go_fuzz_dep_.CoverTab[34827]++
//line /usr/local/go/src/net/textproto/reader.go:746
		// _ = "end of CoverTab[34827]"
//line /usr/local/go/src/net/textproto/reader.go:746
	}
//line /usr/local/go/src/net/textproto/reader.go:746
	// _ = "end of CoverTab[34815]"
//line /usr/local/go/src/net/textproto/reader.go:746
	_go_fuzz_dep_.CoverTab[34816]++

							upper := true
							for i, c := range a {
//line /usr/local/go/src/net/textproto/reader.go:749
		_go_fuzz_dep_.CoverTab[34828]++

//line /usr/local/go/src/net/textproto/reader.go:754
		if upper && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:754
			_go_fuzz_dep_.CoverTab[34830]++
//line /usr/local/go/src/net/textproto/reader.go:754
			return 'a' <= c
//line /usr/local/go/src/net/textproto/reader.go:754
			// _ = "end of CoverTab[34830]"
//line /usr/local/go/src/net/textproto/reader.go:754
		}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:754
			_go_fuzz_dep_.CoverTab[34831]++
//line /usr/local/go/src/net/textproto/reader.go:754
			return c <= 'z'
//line /usr/local/go/src/net/textproto/reader.go:754
			// _ = "end of CoverTab[34831]"
//line /usr/local/go/src/net/textproto/reader.go:754
		}() {
//line /usr/local/go/src/net/textproto/reader.go:754
			_go_fuzz_dep_.CoverTab[34832]++
									c -= toLower
//line /usr/local/go/src/net/textproto/reader.go:755
			// _ = "end of CoverTab[34832]"
		} else {
//line /usr/local/go/src/net/textproto/reader.go:756
			_go_fuzz_dep_.CoverTab[34833]++
//line /usr/local/go/src/net/textproto/reader.go:756
			if !upper && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:756
				_go_fuzz_dep_.CoverTab[34834]++
//line /usr/local/go/src/net/textproto/reader.go:756
				return 'A' <= c
//line /usr/local/go/src/net/textproto/reader.go:756
				// _ = "end of CoverTab[34834]"
//line /usr/local/go/src/net/textproto/reader.go:756
			}() && func() bool {
//line /usr/local/go/src/net/textproto/reader.go:756
				_go_fuzz_dep_.CoverTab[34835]++
//line /usr/local/go/src/net/textproto/reader.go:756
				return c <= 'Z'
//line /usr/local/go/src/net/textproto/reader.go:756
				// _ = "end of CoverTab[34835]"
//line /usr/local/go/src/net/textproto/reader.go:756
			}() {
//line /usr/local/go/src/net/textproto/reader.go:756
				_go_fuzz_dep_.CoverTab[34836]++
										c += toLower
//line /usr/local/go/src/net/textproto/reader.go:757
				// _ = "end of CoverTab[34836]"
			} else {
//line /usr/local/go/src/net/textproto/reader.go:758
				_go_fuzz_dep_.CoverTab[34837]++
//line /usr/local/go/src/net/textproto/reader.go:758
				// _ = "end of CoverTab[34837]"
//line /usr/local/go/src/net/textproto/reader.go:758
			}
//line /usr/local/go/src/net/textproto/reader.go:758
			// _ = "end of CoverTab[34833]"
//line /usr/local/go/src/net/textproto/reader.go:758
		}
//line /usr/local/go/src/net/textproto/reader.go:758
		// _ = "end of CoverTab[34828]"
//line /usr/local/go/src/net/textproto/reader.go:758
		_go_fuzz_dep_.CoverTab[34829]++
								a[i] = c
								upper = c == '-'
//line /usr/local/go/src/net/textproto/reader.go:760
		// _ = "end of CoverTab[34829]"
	}
//line /usr/local/go/src/net/textproto/reader.go:761
	// _ = "end of CoverTab[34816]"
//line /usr/local/go/src/net/textproto/reader.go:761
	_go_fuzz_dep_.CoverTab[34817]++
							commonHeaderOnce.Do(initCommonHeader)

//line /usr/local/go/src/net/textproto/reader.go:766
	if v := commonHeader[string(a)]; v != "" {
//line /usr/local/go/src/net/textproto/reader.go:766
		_go_fuzz_dep_.CoverTab[34838]++
								return v, true
//line /usr/local/go/src/net/textproto/reader.go:767
		// _ = "end of CoverTab[34838]"
	} else {
//line /usr/local/go/src/net/textproto/reader.go:768
		_go_fuzz_dep_.CoverTab[34839]++
//line /usr/local/go/src/net/textproto/reader.go:768
		// _ = "end of CoverTab[34839]"
//line /usr/local/go/src/net/textproto/reader.go:768
	}
//line /usr/local/go/src/net/textproto/reader.go:768
	// _ = "end of CoverTab[34817]"
//line /usr/local/go/src/net/textproto/reader.go:768
	_go_fuzz_dep_.CoverTab[34818]++
							return string(a), true
//line /usr/local/go/src/net/textproto/reader.go:769
	// _ = "end of CoverTab[34818]"
}

// commonHeader interns common header strings.
var commonHeader map[string]string

var commonHeaderOnce sync.Once

func initCommonHeader() {
//line /usr/local/go/src/net/textproto/reader.go:777
	_go_fuzz_dep_.CoverTab[34840]++
							commonHeader = make(map[string]string)
							for _, v := range []string{
		"Accept",
		"Accept-Charset",
		"Accept-Encoding",
		"Accept-Language",
		"Accept-Ranges",
		"Cache-Control",
		"Cc",
		"Connection",
		"Content-Id",
		"Content-Language",
		"Content-Length",
		"Content-Transfer-Encoding",
		"Content-Type",
		"Cookie",
		"Date",
		"Dkim-Signature",
		"Etag",
		"Expires",
		"From",
		"Host",
		"If-Modified-Since",
		"If-None-Match",
		"In-Reply-To",
		"Last-Modified",
		"Location",
		"Message-Id",
		"Mime-Version",
		"Pragma",
		"Received",
		"Return-Path",
		"Server",
		"Set-Cookie",
		"Subject",
		"To",
		"User-Agent",
		"Via",
		"X-Forwarded-For",
		"X-Imforwards",
		"X-Powered-By",
	} {
//line /usr/local/go/src/net/textproto/reader.go:819
		_go_fuzz_dep_.CoverTab[34841]++
								commonHeader[v] = v
//line /usr/local/go/src/net/textproto/reader.go:820
		// _ = "end of CoverTab[34841]"
	}
//line /usr/local/go/src/net/textproto/reader.go:821
	// _ = "end of CoverTab[34840]"
}

//line /usr/local/go/src/net/textproto/reader.go:822
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/net/textproto/reader.go:822
var _ = _go_fuzz_dep_.CoverTab
