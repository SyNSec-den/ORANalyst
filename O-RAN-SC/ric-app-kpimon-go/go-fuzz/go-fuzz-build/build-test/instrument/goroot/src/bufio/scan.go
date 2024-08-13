// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/bufio/scan.go:5
package bufio

//line /snap/go/10455/src/bufio/scan.go:5
import (
//line /snap/go/10455/src/bufio/scan.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/bufio/scan.go:5
)
//line /snap/go/10455/src/bufio/scan.go:5
import (
//line /snap/go/10455/src/bufio/scan.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/bufio/scan.go:5
)

import (
	"bytes"
	"errors"
	"io"
	"unicode/utf8"
)

// Scanner provides a convenient interface for reading data such as
//line /snap/go/10455/src/bufio/scan.go:14
// a file of newline-delimited lines of text. Successive calls to
//line /snap/go/10455/src/bufio/scan.go:14
// the Scan method will step through the 'tokens' of a file, skipping
//line /snap/go/10455/src/bufio/scan.go:14
// the bytes between the tokens. The specification of a token is
//line /snap/go/10455/src/bufio/scan.go:14
// defined by a split function of type SplitFunc; the default split
//line /snap/go/10455/src/bufio/scan.go:14
// function breaks the input into lines with line termination stripped. Split
//line /snap/go/10455/src/bufio/scan.go:14
// functions are defined in this package for scanning a file into
//line /snap/go/10455/src/bufio/scan.go:14
// lines, bytes, UTF-8-encoded runes, and space-delimited words. The
//line /snap/go/10455/src/bufio/scan.go:14
// client may instead provide a custom split function.
//line /snap/go/10455/src/bufio/scan.go:14
//
//line /snap/go/10455/src/bufio/scan.go:14
// Scanning stops unrecoverably at EOF, the first I/O error, or a token too
//line /snap/go/10455/src/bufio/scan.go:14
// large to fit in the buffer. When a scan stops, the reader may have
//line /snap/go/10455/src/bufio/scan.go:14
// advanced arbitrarily far past the last token. Programs that need more
//line /snap/go/10455/src/bufio/scan.go:14
// control over error handling or large tokens, or must run sequential scans
//line /snap/go/10455/src/bufio/scan.go:14
// on a reader, should use bufio.Reader instead.
//line /snap/go/10455/src/bufio/scan.go:29
type Scanner struct {
	r		io.Reader	// The reader provided by the client.
	split		SplitFunc	// The function to split the tokens.
	maxTokenSize	int		// Maximum size of a token; modified by tests.
	token		[]byte		// Last token returned by split.
	buf		[]byte		// Buffer used as argument to split.
	start		int		// First non-processed byte in buf.
	end		int		// End of data in buf.
	err		error		// Sticky error.
	empties		int		// Count of successive empty tokens.
	scanCalled	bool		// Scan has been called; buffer is in use.
	done		bool		// Scan has finished.
}

// SplitFunc is the signature of the split function used to tokenize the
//line /snap/go/10455/src/bufio/scan.go:43
// input. The arguments are an initial substring of the remaining unprocessed
//line /snap/go/10455/src/bufio/scan.go:43
// data and a flag, atEOF, that reports whether the Reader has no more data
//line /snap/go/10455/src/bufio/scan.go:43
// to give. The return values are the number of bytes to advance the input
//line /snap/go/10455/src/bufio/scan.go:43
// and the next token to return to the user, if any, plus an error, if any.
//line /snap/go/10455/src/bufio/scan.go:43
//
//line /snap/go/10455/src/bufio/scan.go:43
// Scanning stops if the function returns an error, in which case some of
//line /snap/go/10455/src/bufio/scan.go:43
// the input may be discarded. If that error is ErrFinalToken, scanning
//line /snap/go/10455/src/bufio/scan.go:43
// stops with no error.
//line /snap/go/10455/src/bufio/scan.go:43
//
//line /snap/go/10455/src/bufio/scan.go:43
// Otherwise, the Scanner advances the input. If the token is not nil,
//line /snap/go/10455/src/bufio/scan.go:43
// the Scanner returns it to the user. If the token is nil, the
//line /snap/go/10455/src/bufio/scan.go:43
// Scanner reads more data and continues scanning; if there is no more
//line /snap/go/10455/src/bufio/scan.go:43
// data--if atEOF was true--the Scanner returns. If the data does not
//line /snap/go/10455/src/bufio/scan.go:43
// yet hold a complete token, for instance if it has no newline while
//line /snap/go/10455/src/bufio/scan.go:43
// scanning lines, a SplitFunc can return (0, nil, nil) to signal the
//line /snap/go/10455/src/bufio/scan.go:43
// Scanner to read more data into the slice and try again with a
//line /snap/go/10455/src/bufio/scan.go:43
// longer slice starting at the same point in the input.
//line /snap/go/10455/src/bufio/scan.go:43
//
//line /snap/go/10455/src/bufio/scan.go:43
// The function is never called with an empty data slice unless atEOF
//line /snap/go/10455/src/bufio/scan.go:43
// is true. If atEOF is true, however, data may be non-empty and,
//line /snap/go/10455/src/bufio/scan.go:43
// as always, holds unprocessed text.
//line /snap/go/10455/src/bufio/scan.go:65
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// Errors returned by Scanner.
var (
	ErrTooLong		= errors.New("bufio.Scanner: token too long")
	ErrNegativeAdvance	= errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	ErrAdvanceTooFar	= errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
	ErrBadReadCount		= errors.New("bufio.Scanner: Read returned impossible count")
)

const (
	// MaxScanTokenSize is the maximum size used to buffer a token
	// unless the user provides an explicit buffer with Scanner.Buffer.
	// The actual maximum token size may be smaller as the buffer
	// may need to include, for instance, a newline.
	MaxScanTokenSize	= 64 * 1024

	startBufSize	= 4096	// Size of initial allocation for buffer.
)

// NewScanner returns a new Scanner to read from r.
//line /snap/go/10455/src/bufio/scan.go:85
// The split function defaults to ScanLines.
//line /snap/go/10455/src/bufio/scan.go:87
func NewScanner(r io.Reader) *Scanner {
//line /snap/go/10455/src/bufio/scan.go:87
	_go_fuzz_dep_.CoverTab[2142]++
						return &Scanner{
		r:		r,
		split:		ScanLines,
		maxTokenSize:	MaxScanTokenSize,
	}
//line /snap/go/10455/src/bufio/scan.go:92
	// _ = "end of CoverTab[2142]"
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (s *Scanner) Err() error {
//line /snap/go/10455/src/bufio/scan.go:96
	_go_fuzz_dep_.CoverTab[2143]++
						if s.err == io.EOF {
//line /snap/go/10455/src/bufio/scan.go:97
		_go_fuzz_dep_.CoverTab[525897]++
//line /snap/go/10455/src/bufio/scan.go:97
		_go_fuzz_dep_.CoverTab[2145]++
							return nil
//line /snap/go/10455/src/bufio/scan.go:98
		// _ = "end of CoverTab[2145]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:99
		_go_fuzz_dep_.CoverTab[525898]++
//line /snap/go/10455/src/bufio/scan.go:99
		_go_fuzz_dep_.CoverTab[2146]++
//line /snap/go/10455/src/bufio/scan.go:99
		// _ = "end of CoverTab[2146]"
//line /snap/go/10455/src/bufio/scan.go:99
	}
//line /snap/go/10455/src/bufio/scan.go:99
	// _ = "end of CoverTab[2143]"
//line /snap/go/10455/src/bufio/scan.go:99
	_go_fuzz_dep_.CoverTab[2144]++
						return s.err
//line /snap/go/10455/src/bufio/scan.go:100
	// _ = "end of CoverTab[2144]"
}

// Bytes returns the most recent token generated by a call to Scan.
//line /snap/go/10455/src/bufio/scan.go:103
// The underlying array may point to data that will be overwritten
//line /snap/go/10455/src/bufio/scan.go:103
// by a subsequent call to Scan. It does no allocation.
//line /snap/go/10455/src/bufio/scan.go:106
func (s *Scanner) Bytes() []byte {
//line /snap/go/10455/src/bufio/scan.go:106
	_go_fuzz_dep_.CoverTab[2147]++
						return s.token
//line /snap/go/10455/src/bufio/scan.go:107
	// _ = "end of CoverTab[2147]"
}

// Text returns the most recent token generated by a call to Scan
//line /snap/go/10455/src/bufio/scan.go:110
// as a newly allocated string holding its bytes.
//line /snap/go/10455/src/bufio/scan.go:112
func (s *Scanner) Text() string {
//line /snap/go/10455/src/bufio/scan.go:112
	_go_fuzz_dep_.CoverTab[2148]++
						return string(s.token)
//line /snap/go/10455/src/bufio/scan.go:113
	// _ = "end of CoverTab[2148]"
}

// ErrFinalToken is a special sentinel error value. It is intended to be
//line /snap/go/10455/src/bufio/scan.go:116
// returned by a Split function to indicate that the token being delivered
//line /snap/go/10455/src/bufio/scan.go:116
// with the error is the last token and scanning should stop after this one.
//line /snap/go/10455/src/bufio/scan.go:116
// After ErrFinalToken is received by Scan, scanning stops with no error.
//line /snap/go/10455/src/bufio/scan.go:116
// The value is useful to stop processing early or when it is necessary to
//line /snap/go/10455/src/bufio/scan.go:116
// deliver a final empty token. One could achieve the same behavior
//line /snap/go/10455/src/bufio/scan.go:116
// with a custom error value but providing one here is tidier.
//line /snap/go/10455/src/bufio/scan.go:116
// See the emptyFinalToken example for a use of this value.
//line /snap/go/10455/src/bufio/scan.go:124
var ErrFinalToken = errors.New("final token")

// Scan advances the Scanner to the next token, which will then be
//line /snap/go/10455/src/bufio/scan.go:126
// available through the Bytes or Text method. It returns false when the
//line /snap/go/10455/src/bufio/scan.go:126
// scan stops, either by reaching the end of the input or an error.
//line /snap/go/10455/src/bufio/scan.go:126
// After Scan returns false, the Err method will return any error that
//line /snap/go/10455/src/bufio/scan.go:126
// occurred during scanning, except that if it was io.EOF, Err
//line /snap/go/10455/src/bufio/scan.go:126
// will return nil.
//line /snap/go/10455/src/bufio/scan.go:126
// Scan panics if the split function returns too many empty
//line /snap/go/10455/src/bufio/scan.go:126
// tokens without advancing the input. This is a common error mode for
//line /snap/go/10455/src/bufio/scan.go:126
// scanners.
//line /snap/go/10455/src/bufio/scan.go:135
func (s *Scanner) Scan() bool {
//line /snap/go/10455/src/bufio/scan.go:135
	_go_fuzz_dep_.CoverTab[2149]++
						if s.done {
//line /snap/go/10455/src/bufio/scan.go:136
		_go_fuzz_dep_.CoverTab[525899]++
//line /snap/go/10455/src/bufio/scan.go:136
		_go_fuzz_dep_.CoverTab[2151]++
							return false
//line /snap/go/10455/src/bufio/scan.go:137
		// _ = "end of CoverTab[2151]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:138
		_go_fuzz_dep_.CoverTab[525900]++
//line /snap/go/10455/src/bufio/scan.go:138
		_go_fuzz_dep_.CoverTab[2152]++
//line /snap/go/10455/src/bufio/scan.go:138
		// _ = "end of CoverTab[2152]"
//line /snap/go/10455/src/bufio/scan.go:138
	}
//line /snap/go/10455/src/bufio/scan.go:138
	// _ = "end of CoverTab[2149]"
//line /snap/go/10455/src/bufio/scan.go:138
	_go_fuzz_dep_.CoverTab[2150]++
						s.scanCalled = true
//line /snap/go/10455/src/bufio/scan.go:139
	_go_fuzz_dep_.CoverTab[786568] = 0

						for {
//line /snap/go/10455/src/bufio/scan.go:141
		if _go_fuzz_dep_.CoverTab[786568] == 0 {
//line /snap/go/10455/src/bufio/scan.go:141
			_go_fuzz_dep_.CoverTab[525978]++
//line /snap/go/10455/src/bufio/scan.go:141
		} else {
//line /snap/go/10455/src/bufio/scan.go:141
			_go_fuzz_dep_.CoverTab[525979]++
//line /snap/go/10455/src/bufio/scan.go:141
		}
//line /snap/go/10455/src/bufio/scan.go:141
		_go_fuzz_dep_.CoverTab[786568] = 1
//line /snap/go/10455/src/bufio/scan.go:141
		_go_fuzz_dep_.CoverTab[2153]++

//line /snap/go/10455/src/bufio/scan.go:145
		if s.end > s.start || func() bool {
//line /snap/go/10455/src/bufio/scan.go:145
			_go_fuzz_dep_.CoverTab[2158]++
//line /snap/go/10455/src/bufio/scan.go:145
			return s.err != nil
//line /snap/go/10455/src/bufio/scan.go:145
			// _ = "end of CoverTab[2158]"
//line /snap/go/10455/src/bufio/scan.go:145
		}() {
//line /snap/go/10455/src/bufio/scan.go:145
			_go_fuzz_dep_.CoverTab[525901]++
//line /snap/go/10455/src/bufio/scan.go:145
			_go_fuzz_dep_.CoverTab[2159]++
								advance, token, err := s.split(s.buf[s.start:s.end], s.err != nil)
								if err != nil {
//line /snap/go/10455/src/bufio/scan.go:147
				_go_fuzz_dep_.CoverTab[525903]++
//line /snap/go/10455/src/bufio/scan.go:147
				_go_fuzz_dep_.CoverTab[2162]++
									if err == ErrFinalToken {
//line /snap/go/10455/src/bufio/scan.go:148
					_go_fuzz_dep_.CoverTab[525905]++
//line /snap/go/10455/src/bufio/scan.go:148
					_go_fuzz_dep_.CoverTab[2164]++
										s.token = token
										s.done = true
										return true
//line /snap/go/10455/src/bufio/scan.go:151
					// _ = "end of CoverTab[2164]"
				} else {
//line /snap/go/10455/src/bufio/scan.go:152
					_go_fuzz_dep_.CoverTab[525906]++
//line /snap/go/10455/src/bufio/scan.go:152
					_go_fuzz_dep_.CoverTab[2165]++
//line /snap/go/10455/src/bufio/scan.go:152
					// _ = "end of CoverTab[2165]"
//line /snap/go/10455/src/bufio/scan.go:152
				}
//line /snap/go/10455/src/bufio/scan.go:152
				// _ = "end of CoverTab[2162]"
//line /snap/go/10455/src/bufio/scan.go:152
				_go_fuzz_dep_.CoverTab[2163]++
									s.setErr(err)
									return false
//line /snap/go/10455/src/bufio/scan.go:154
				// _ = "end of CoverTab[2163]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:155
				_go_fuzz_dep_.CoverTab[525904]++
//line /snap/go/10455/src/bufio/scan.go:155
				_go_fuzz_dep_.CoverTab[2166]++
//line /snap/go/10455/src/bufio/scan.go:155
				// _ = "end of CoverTab[2166]"
//line /snap/go/10455/src/bufio/scan.go:155
			}
//line /snap/go/10455/src/bufio/scan.go:155
			// _ = "end of CoverTab[2159]"
//line /snap/go/10455/src/bufio/scan.go:155
			_go_fuzz_dep_.CoverTab[2160]++
								if !s.advance(advance) {
//line /snap/go/10455/src/bufio/scan.go:156
				_go_fuzz_dep_.CoverTab[525907]++
//line /snap/go/10455/src/bufio/scan.go:156
				_go_fuzz_dep_.CoverTab[2167]++
									return false
//line /snap/go/10455/src/bufio/scan.go:157
				// _ = "end of CoverTab[2167]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:158
				_go_fuzz_dep_.CoverTab[525908]++
//line /snap/go/10455/src/bufio/scan.go:158
				_go_fuzz_dep_.CoverTab[2168]++
//line /snap/go/10455/src/bufio/scan.go:158
				// _ = "end of CoverTab[2168]"
//line /snap/go/10455/src/bufio/scan.go:158
			}
//line /snap/go/10455/src/bufio/scan.go:158
			// _ = "end of CoverTab[2160]"
//line /snap/go/10455/src/bufio/scan.go:158
			_go_fuzz_dep_.CoverTab[2161]++
								s.token = token
								if token != nil {
//line /snap/go/10455/src/bufio/scan.go:160
				_go_fuzz_dep_.CoverTab[525909]++
//line /snap/go/10455/src/bufio/scan.go:160
				_go_fuzz_dep_.CoverTab[2169]++
									if s.err == nil || func() bool {
//line /snap/go/10455/src/bufio/scan.go:161
					_go_fuzz_dep_.CoverTab[2171]++
//line /snap/go/10455/src/bufio/scan.go:161
					return advance > 0
//line /snap/go/10455/src/bufio/scan.go:161
					// _ = "end of CoverTab[2171]"
//line /snap/go/10455/src/bufio/scan.go:161
				}() {
//line /snap/go/10455/src/bufio/scan.go:161
					_go_fuzz_dep_.CoverTab[525911]++
//line /snap/go/10455/src/bufio/scan.go:161
					_go_fuzz_dep_.CoverTab[2172]++
										s.empties = 0
//line /snap/go/10455/src/bufio/scan.go:162
					// _ = "end of CoverTab[2172]"
				} else {
//line /snap/go/10455/src/bufio/scan.go:163
					_go_fuzz_dep_.CoverTab[525912]++
//line /snap/go/10455/src/bufio/scan.go:163
					_go_fuzz_dep_.CoverTab[2173]++

										s.empties++
										if s.empties > maxConsecutiveEmptyReads {
//line /snap/go/10455/src/bufio/scan.go:166
						_go_fuzz_dep_.CoverTab[525913]++
//line /snap/go/10455/src/bufio/scan.go:166
						_go_fuzz_dep_.CoverTab[2174]++
											panic("bufio.Scan: too many empty tokens without progressing")
//line /snap/go/10455/src/bufio/scan.go:167
						// _ = "end of CoverTab[2174]"
					} else {
//line /snap/go/10455/src/bufio/scan.go:168
						_go_fuzz_dep_.CoverTab[525914]++
//line /snap/go/10455/src/bufio/scan.go:168
						_go_fuzz_dep_.CoverTab[2175]++
//line /snap/go/10455/src/bufio/scan.go:168
						// _ = "end of CoverTab[2175]"
//line /snap/go/10455/src/bufio/scan.go:168
					}
//line /snap/go/10455/src/bufio/scan.go:168
					// _ = "end of CoverTab[2173]"
				}
//line /snap/go/10455/src/bufio/scan.go:169
				// _ = "end of CoverTab[2169]"
//line /snap/go/10455/src/bufio/scan.go:169
				_go_fuzz_dep_.CoverTab[2170]++
									return true
//line /snap/go/10455/src/bufio/scan.go:170
				// _ = "end of CoverTab[2170]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:171
				_go_fuzz_dep_.CoverTab[525910]++
//line /snap/go/10455/src/bufio/scan.go:171
				_go_fuzz_dep_.CoverTab[2176]++
//line /snap/go/10455/src/bufio/scan.go:171
				// _ = "end of CoverTab[2176]"
//line /snap/go/10455/src/bufio/scan.go:171
			}
//line /snap/go/10455/src/bufio/scan.go:171
			// _ = "end of CoverTab[2161]"
		} else {
//line /snap/go/10455/src/bufio/scan.go:172
			_go_fuzz_dep_.CoverTab[525902]++
//line /snap/go/10455/src/bufio/scan.go:172
			_go_fuzz_dep_.CoverTab[2177]++
//line /snap/go/10455/src/bufio/scan.go:172
			// _ = "end of CoverTab[2177]"
//line /snap/go/10455/src/bufio/scan.go:172
		}
//line /snap/go/10455/src/bufio/scan.go:172
		// _ = "end of CoverTab[2153]"
//line /snap/go/10455/src/bufio/scan.go:172
		_go_fuzz_dep_.CoverTab[2154]++

//line /snap/go/10455/src/bufio/scan.go:175
		if s.err != nil {
//line /snap/go/10455/src/bufio/scan.go:175
			_go_fuzz_dep_.CoverTab[525915]++
//line /snap/go/10455/src/bufio/scan.go:175
			_go_fuzz_dep_.CoverTab[2178]++

								s.start = 0
								s.end = 0
								return false
//line /snap/go/10455/src/bufio/scan.go:179
			// _ = "end of CoverTab[2178]"
		} else {
//line /snap/go/10455/src/bufio/scan.go:180
			_go_fuzz_dep_.CoverTab[525916]++
//line /snap/go/10455/src/bufio/scan.go:180
			_go_fuzz_dep_.CoverTab[2179]++
//line /snap/go/10455/src/bufio/scan.go:180
			// _ = "end of CoverTab[2179]"
//line /snap/go/10455/src/bufio/scan.go:180
		}
//line /snap/go/10455/src/bufio/scan.go:180
		// _ = "end of CoverTab[2154]"
//line /snap/go/10455/src/bufio/scan.go:180
		_go_fuzz_dep_.CoverTab[2155]++

//line /snap/go/10455/src/bufio/scan.go:184
		if s.start > 0 && func() bool {
//line /snap/go/10455/src/bufio/scan.go:184
			_go_fuzz_dep_.CoverTab[2180]++
//line /snap/go/10455/src/bufio/scan.go:184
			return (s.end == len(s.buf) || func() bool {
//line /snap/go/10455/src/bufio/scan.go:184
				_go_fuzz_dep_.CoverTab[2181]++
//line /snap/go/10455/src/bufio/scan.go:184
				return s.start > len(s.buf)/2
//line /snap/go/10455/src/bufio/scan.go:184
				// _ = "end of CoverTab[2181]"
//line /snap/go/10455/src/bufio/scan.go:184
			}())
//line /snap/go/10455/src/bufio/scan.go:184
			// _ = "end of CoverTab[2180]"
//line /snap/go/10455/src/bufio/scan.go:184
		}() {
//line /snap/go/10455/src/bufio/scan.go:184
			_go_fuzz_dep_.CoverTab[525917]++
//line /snap/go/10455/src/bufio/scan.go:184
			_go_fuzz_dep_.CoverTab[2182]++
								copy(s.buf, s.buf[s.start:s.end])
								s.end -= s.start
								s.start = 0
//line /snap/go/10455/src/bufio/scan.go:187
			// _ = "end of CoverTab[2182]"
		} else {
//line /snap/go/10455/src/bufio/scan.go:188
			_go_fuzz_dep_.CoverTab[525918]++
//line /snap/go/10455/src/bufio/scan.go:188
			_go_fuzz_dep_.CoverTab[2183]++
//line /snap/go/10455/src/bufio/scan.go:188
			// _ = "end of CoverTab[2183]"
//line /snap/go/10455/src/bufio/scan.go:188
		}
//line /snap/go/10455/src/bufio/scan.go:188
		// _ = "end of CoverTab[2155]"
//line /snap/go/10455/src/bufio/scan.go:188
		_go_fuzz_dep_.CoverTab[2156]++

							if s.end == len(s.buf) {
//line /snap/go/10455/src/bufio/scan.go:190
			_go_fuzz_dep_.CoverTab[525919]++
//line /snap/go/10455/src/bufio/scan.go:190
			_go_fuzz_dep_.CoverTab[2184]++
			// Guarantee no overflow in the multiplication below.
			const maxInt = int(^uint(0) >> 1)
			if len(s.buf) >= s.maxTokenSize || func() bool {
//line /snap/go/10455/src/bufio/scan.go:193
				_go_fuzz_dep_.CoverTab[2188]++
//line /snap/go/10455/src/bufio/scan.go:193
				return len(s.buf) > maxInt/2
//line /snap/go/10455/src/bufio/scan.go:193
				// _ = "end of CoverTab[2188]"
//line /snap/go/10455/src/bufio/scan.go:193
			}() {
//line /snap/go/10455/src/bufio/scan.go:193
				_go_fuzz_dep_.CoverTab[525921]++
//line /snap/go/10455/src/bufio/scan.go:193
				_go_fuzz_dep_.CoverTab[2189]++
									s.setErr(ErrTooLong)
									return false
//line /snap/go/10455/src/bufio/scan.go:195
				// _ = "end of CoverTab[2189]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:196
				_go_fuzz_dep_.CoverTab[525922]++
//line /snap/go/10455/src/bufio/scan.go:196
				_go_fuzz_dep_.CoverTab[2190]++
//line /snap/go/10455/src/bufio/scan.go:196
				// _ = "end of CoverTab[2190]"
//line /snap/go/10455/src/bufio/scan.go:196
			}
//line /snap/go/10455/src/bufio/scan.go:196
			// _ = "end of CoverTab[2184]"
//line /snap/go/10455/src/bufio/scan.go:196
			_go_fuzz_dep_.CoverTab[2185]++
								newSize := len(s.buf) * 2
								if newSize == 0 {
//line /snap/go/10455/src/bufio/scan.go:198
				_go_fuzz_dep_.CoverTab[525923]++
//line /snap/go/10455/src/bufio/scan.go:198
				_go_fuzz_dep_.CoverTab[2191]++
									newSize = startBufSize
//line /snap/go/10455/src/bufio/scan.go:199
				// _ = "end of CoverTab[2191]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:200
				_go_fuzz_dep_.CoverTab[525924]++
//line /snap/go/10455/src/bufio/scan.go:200
				_go_fuzz_dep_.CoverTab[2192]++
//line /snap/go/10455/src/bufio/scan.go:200
				// _ = "end of CoverTab[2192]"
//line /snap/go/10455/src/bufio/scan.go:200
			}
//line /snap/go/10455/src/bufio/scan.go:200
			// _ = "end of CoverTab[2185]"
//line /snap/go/10455/src/bufio/scan.go:200
			_go_fuzz_dep_.CoverTab[2186]++
								if newSize > s.maxTokenSize {
//line /snap/go/10455/src/bufio/scan.go:201
				_go_fuzz_dep_.CoverTab[525925]++
//line /snap/go/10455/src/bufio/scan.go:201
				_go_fuzz_dep_.CoverTab[2193]++
									newSize = s.maxTokenSize
//line /snap/go/10455/src/bufio/scan.go:202
				// _ = "end of CoverTab[2193]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:203
				_go_fuzz_dep_.CoverTab[525926]++
//line /snap/go/10455/src/bufio/scan.go:203
				_go_fuzz_dep_.CoverTab[2194]++
//line /snap/go/10455/src/bufio/scan.go:203
				// _ = "end of CoverTab[2194]"
//line /snap/go/10455/src/bufio/scan.go:203
			}
//line /snap/go/10455/src/bufio/scan.go:203
			// _ = "end of CoverTab[2186]"
//line /snap/go/10455/src/bufio/scan.go:203
			_go_fuzz_dep_.CoverTab[2187]++
								newBuf := make([]byte, newSize)
								copy(newBuf, s.buf[s.start:s.end])
								s.buf = newBuf
								s.end -= s.start
								s.start = 0
//line /snap/go/10455/src/bufio/scan.go:208
			// _ = "end of CoverTab[2187]"
		} else {
//line /snap/go/10455/src/bufio/scan.go:209
			_go_fuzz_dep_.CoverTab[525920]++
//line /snap/go/10455/src/bufio/scan.go:209
			_go_fuzz_dep_.CoverTab[2195]++
//line /snap/go/10455/src/bufio/scan.go:209
			// _ = "end of CoverTab[2195]"
//line /snap/go/10455/src/bufio/scan.go:209
		}
//line /snap/go/10455/src/bufio/scan.go:209
		// _ = "end of CoverTab[2156]"
//line /snap/go/10455/src/bufio/scan.go:209
		_go_fuzz_dep_.CoverTab[2157]++
//line /snap/go/10455/src/bufio/scan.go:209
		_go_fuzz_dep_.CoverTab[786569] = 0

//line /snap/go/10455/src/bufio/scan.go:213
		for loop := 0; ; {
//line /snap/go/10455/src/bufio/scan.go:213
			if _go_fuzz_dep_.CoverTab[786569] == 0 {
//line /snap/go/10455/src/bufio/scan.go:213
				_go_fuzz_dep_.CoverTab[525982]++
//line /snap/go/10455/src/bufio/scan.go:213
			} else {
//line /snap/go/10455/src/bufio/scan.go:213
				_go_fuzz_dep_.CoverTab[525983]++
//line /snap/go/10455/src/bufio/scan.go:213
			}
//line /snap/go/10455/src/bufio/scan.go:213
			_go_fuzz_dep_.CoverTab[786569] = 1
//line /snap/go/10455/src/bufio/scan.go:213
			_go_fuzz_dep_.CoverTab[2196]++
								n, err := s.r.Read(s.buf[s.end:len(s.buf)])
								if n < 0 || func() bool {
//line /snap/go/10455/src/bufio/scan.go:215
				_go_fuzz_dep_.CoverTab[2200]++
//line /snap/go/10455/src/bufio/scan.go:215
				return len(s.buf)-s.end < n
//line /snap/go/10455/src/bufio/scan.go:215
				// _ = "end of CoverTab[2200]"
//line /snap/go/10455/src/bufio/scan.go:215
			}() {
//line /snap/go/10455/src/bufio/scan.go:215
				_go_fuzz_dep_.CoverTab[525927]++
//line /snap/go/10455/src/bufio/scan.go:215
				_go_fuzz_dep_.CoverTab[2201]++
									s.setErr(ErrBadReadCount)
									break
//line /snap/go/10455/src/bufio/scan.go:217
				// _ = "end of CoverTab[2201]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:218
				_go_fuzz_dep_.CoverTab[525928]++
//line /snap/go/10455/src/bufio/scan.go:218
				_go_fuzz_dep_.CoverTab[2202]++
//line /snap/go/10455/src/bufio/scan.go:218
				// _ = "end of CoverTab[2202]"
//line /snap/go/10455/src/bufio/scan.go:218
			}
//line /snap/go/10455/src/bufio/scan.go:218
			// _ = "end of CoverTab[2196]"
//line /snap/go/10455/src/bufio/scan.go:218
			_go_fuzz_dep_.CoverTab[2197]++
								s.end += n
								if err != nil {
//line /snap/go/10455/src/bufio/scan.go:220
				_go_fuzz_dep_.CoverTab[525929]++
//line /snap/go/10455/src/bufio/scan.go:220
				_go_fuzz_dep_.CoverTab[2203]++
									s.setErr(err)
									break
//line /snap/go/10455/src/bufio/scan.go:222
				// _ = "end of CoverTab[2203]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:223
				_go_fuzz_dep_.CoverTab[525930]++
//line /snap/go/10455/src/bufio/scan.go:223
				_go_fuzz_dep_.CoverTab[2204]++
//line /snap/go/10455/src/bufio/scan.go:223
				// _ = "end of CoverTab[2204]"
//line /snap/go/10455/src/bufio/scan.go:223
			}
//line /snap/go/10455/src/bufio/scan.go:223
			// _ = "end of CoverTab[2197]"
//line /snap/go/10455/src/bufio/scan.go:223
			_go_fuzz_dep_.CoverTab[2198]++
								if n > 0 {
//line /snap/go/10455/src/bufio/scan.go:224
				_go_fuzz_dep_.CoverTab[525931]++
//line /snap/go/10455/src/bufio/scan.go:224
				_go_fuzz_dep_.CoverTab[2205]++
									s.empties = 0
									break
//line /snap/go/10455/src/bufio/scan.go:226
				// _ = "end of CoverTab[2205]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:227
				_go_fuzz_dep_.CoverTab[525932]++
//line /snap/go/10455/src/bufio/scan.go:227
				_go_fuzz_dep_.CoverTab[2206]++
//line /snap/go/10455/src/bufio/scan.go:227
				// _ = "end of CoverTab[2206]"
//line /snap/go/10455/src/bufio/scan.go:227
			}
//line /snap/go/10455/src/bufio/scan.go:227
			// _ = "end of CoverTab[2198]"
//line /snap/go/10455/src/bufio/scan.go:227
			_go_fuzz_dep_.CoverTab[2199]++
								loop++
								if loop > maxConsecutiveEmptyReads {
//line /snap/go/10455/src/bufio/scan.go:229
				_go_fuzz_dep_.CoverTab[525933]++
//line /snap/go/10455/src/bufio/scan.go:229
				_go_fuzz_dep_.CoverTab[2207]++
									s.setErr(io.ErrNoProgress)
									break
//line /snap/go/10455/src/bufio/scan.go:231
				// _ = "end of CoverTab[2207]"
			} else {
//line /snap/go/10455/src/bufio/scan.go:232
				_go_fuzz_dep_.CoverTab[525934]++
//line /snap/go/10455/src/bufio/scan.go:232
				_go_fuzz_dep_.CoverTab[2208]++
//line /snap/go/10455/src/bufio/scan.go:232
				// _ = "end of CoverTab[2208]"
//line /snap/go/10455/src/bufio/scan.go:232
			}
//line /snap/go/10455/src/bufio/scan.go:232
			// _ = "end of CoverTab[2199]"
		}
//line /snap/go/10455/src/bufio/scan.go:233
		// _ = "end of CoverTab[2157]"
	}
//line /snap/go/10455/src/bufio/scan.go:234
	// _ = "end of CoverTab[2150]"
}

// advance consumes n bytes of the buffer. It reports whether the advance was legal.
func (s *Scanner) advance(n int) bool {
//line /snap/go/10455/src/bufio/scan.go:238
	_go_fuzz_dep_.CoverTab[2209]++
						if n < 0 {
//line /snap/go/10455/src/bufio/scan.go:239
		_go_fuzz_dep_.CoverTab[525935]++
//line /snap/go/10455/src/bufio/scan.go:239
		_go_fuzz_dep_.CoverTab[2212]++
							s.setErr(ErrNegativeAdvance)
							return false
//line /snap/go/10455/src/bufio/scan.go:241
		// _ = "end of CoverTab[2212]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:242
		_go_fuzz_dep_.CoverTab[525936]++
//line /snap/go/10455/src/bufio/scan.go:242
		_go_fuzz_dep_.CoverTab[2213]++
//line /snap/go/10455/src/bufio/scan.go:242
		// _ = "end of CoverTab[2213]"
//line /snap/go/10455/src/bufio/scan.go:242
	}
//line /snap/go/10455/src/bufio/scan.go:242
	// _ = "end of CoverTab[2209]"
//line /snap/go/10455/src/bufio/scan.go:242
	_go_fuzz_dep_.CoverTab[2210]++
						if n > s.end-s.start {
//line /snap/go/10455/src/bufio/scan.go:243
		_go_fuzz_dep_.CoverTab[525937]++
//line /snap/go/10455/src/bufio/scan.go:243
		_go_fuzz_dep_.CoverTab[2214]++
							s.setErr(ErrAdvanceTooFar)
							return false
//line /snap/go/10455/src/bufio/scan.go:245
		// _ = "end of CoverTab[2214]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:246
		_go_fuzz_dep_.CoverTab[525938]++
//line /snap/go/10455/src/bufio/scan.go:246
		_go_fuzz_dep_.CoverTab[2215]++
//line /snap/go/10455/src/bufio/scan.go:246
		// _ = "end of CoverTab[2215]"
//line /snap/go/10455/src/bufio/scan.go:246
	}
//line /snap/go/10455/src/bufio/scan.go:246
	// _ = "end of CoverTab[2210]"
//line /snap/go/10455/src/bufio/scan.go:246
	_go_fuzz_dep_.CoverTab[2211]++
						s.start += n
						return true
//line /snap/go/10455/src/bufio/scan.go:248
	// _ = "end of CoverTab[2211]"
}

// setErr records the first error encountered.
func (s *Scanner) setErr(err error) {
//line /snap/go/10455/src/bufio/scan.go:252
	_go_fuzz_dep_.CoverTab[2216]++
						if s.err == nil || func() bool {
//line /snap/go/10455/src/bufio/scan.go:253
		_go_fuzz_dep_.CoverTab[2217]++
//line /snap/go/10455/src/bufio/scan.go:253
		return s.err == io.EOF
//line /snap/go/10455/src/bufio/scan.go:253
		// _ = "end of CoverTab[2217]"
//line /snap/go/10455/src/bufio/scan.go:253
	}() {
//line /snap/go/10455/src/bufio/scan.go:253
		_go_fuzz_dep_.CoverTab[525939]++
//line /snap/go/10455/src/bufio/scan.go:253
		_go_fuzz_dep_.CoverTab[2218]++
							s.err = err
//line /snap/go/10455/src/bufio/scan.go:254
		// _ = "end of CoverTab[2218]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:255
		_go_fuzz_dep_.CoverTab[525940]++
//line /snap/go/10455/src/bufio/scan.go:255
		_go_fuzz_dep_.CoverTab[2219]++
//line /snap/go/10455/src/bufio/scan.go:255
		// _ = "end of CoverTab[2219]"
//line /snap/go/10455/src/bufio/scan.go:255
	}
//line /snap/go/10455/src/bufio/scan.go:255
	// _ = "end of CoverTab[2216]"
}

// Buffer sets the initial buffer to use when scanning and the maximum
//line /snap/go/10455/src/bufio/scan.go:258
// size of buffer that may be allocated during scanning. The maximum
//line /snap/go/10455/src/bufio/scan.go:258
// token size is the larger of max and cap(buf). If max <= cap(buf),
//line /snap/go/10455/src/bufio/scan.go:258
// Scan will use this buffer only and do no allocation.
//line /snap/go/10455/src/bufio/scan.go:258
//
//line /snap/go/10455/src/bufio/scan.go:258
// By default, Scan uses an internal buffer and sets the
//line /snap/go/10455/src/bufio/scan.go:258
// maximum token size to MaxScanTokenSize.
//line /snap/go/10455/src/bufio/scan.go:258
//
//line /snap/go/10455/src/bufio/scan.go:258
// Buffer panics if it is called after scanning has started.
//line /snap/go/10455/src/bufio/scan.go:267
func (s *Scanner) Buffer(buf []byte, max int) {
//line /snap/go/10455/src/bufio/scan.go:267
	_go_fuzz_dep_.CoverTab[2220]++
						if s.scanCalled {
//line /snap/go/10455/src/bufio/scan.go:268
		_go_fuzz_dep_.CoverTab[525941]++
//line /snap/go/10455/src/bufio/scan.go:268
		_go_fuzz_dep_.CoverTab[2222]++
							panic("Buffer called after Scan")
//line /snap/go/10455/src/bufio/scan.go:269
		// _ = "end of CoverTab[2222]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:270
		_go_fuzz_dep_.CoverTab[525942]++
//line /snap/go/10455/src/bufio/scan.go:270
		_go_fuzz_dep_.CoverTab[2223]++
//line /snap/go/10455/src/bufio/scan.go:270
		// _ = "end of CoverTab[2223]"
//line /snap/go/10455/src/bufio/scan.go:270
	}
//line /snap/go/10455/src/bufio/scan.go:270
	// _ = "end of CoverTab[2220]"
//line /snap/go/10455/src/bufio/scan.go:270
	_go_fuzz_dep_.CoverTab[2221]++
						s.buf = buf[0:cap(buf)]
						s.maxTokenSize = max
//line /snap/go/10455/src/bufio/scan.go:272
	// _ = "end of CoverTab[2221]"
}

// Split sets the split function for the Scanner.
//line /snap/go/10455/src/bufio/scan.go:275
// The default split function is ScanLines.
//line /snap/go/10455/src/bufio/scan.go:275
//
//line /snap/go/10455/src/bufio/scan.go:275
// Split panics if it is called after scanning has started.
//line /snap/go/10455/src/bufio/scan.go:279
func (s *Scanner) Split(split SplitFunc) {
//line /snap/go/10455/src/bufio/scan.go:279
	_go_fuzz_dep_.CoverTab[2224]++
						if s.scanCalled {
//line /snap/go/10455/src/bufio/scan.go:280
		_go_fuzz_dep_.CoverTab[525943]++
//line /snap/go/10455/src/bufio/scan.go:280
		_go_fuzz_dep_.CoverTab[2226]++
							panic("Split called after Scan")
//line /snap/go/10455/src/bufio/scan.go:281
		// _ = "end of CoverTab[2226]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:282
		_go_fuzz_dep_.CoverTab[525944]++
//line /snap/go/10455/src/bufio/scan.go:282
		_go_fuzz_dep_.CoverTab[2227]++
//line /snap/go/10455/src/bufio/scan.go:282
		// _ = "end of CoverTab[2227]"
//line /snap/go/10455/src/bufio/scan.go:282
	}
//line /snap/go/10455/src/bufio/scan.go:282
	// _ = "end of CoverTab[2224]"
//line /snap/go/10455/src/bufio/scan.go:282
	_go_fuzz_dep_.CoverTab[2225]++
						s.split = split
//line /snap/go/10455/src/bufio/scan.go:283
	// _ = "end of CoverTab[2225]"
}

//line /snap/go/10455/src/bufio/scan.go:288
// ScanBytes is a split function for a Scanner that returns each byte as a token.
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /snap/go/10455/src/bufio/scan.go:289
	_go_fuzz_dep_.CoverTab[2228]++
						if atEOF && func() bool {
//line /snap/go/10455/src/bufio/scan.go:290
		_go_fuzz_dep_.CoverTab[2230]++
//line /snap/go/10455/src/bufio/scan.go:290
		return len(data) == 0
//line /snap/go/10455/src/bufio/scan.go:290
		// _ = "end of CoverTab[2230]"
//line /snap/go/10455/src/bufio/scan.go:290
	}() {
//line /snap/go/10455/src/bufio/scan.go:290
		_go_fuzz_dep_.CoverTab[525945]++
//line /snap/go/10455/src/bufio/scan.go:290
		_go_fuzz_dep_.CoverTab[2231]++
							return 0, nil, nil
//line /snap/go/10455/src/bufio/scan.go:291
		// _ = "end of CoverTab[2231]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:292
		_go_fuzz_dep_.CoverTab[525946]++
//line /snap/go/10455/src/bufio/scan.go:292
		_go_fuzz_dep_.CoverTab[2232]++
//line /snap/go/10455/src/bufio/scan.go:292
		// _ = "end of CoverTab[2232]"
//line /snap/go/10455/src/bufio/scan.go:292
	}
//line /snap/go/10455/src/bufio/scan.go:292
	// _ = "end of CoverTab[2228]"
//line /snap/go/10455/src/bufio/scan.go:292
	_go_fuzz_dep_.CoverTab[2229]++
						return 1, data[0:1], nil
//line /snap/go/10455/src/bufio/scan.go:293
	// _ = "end of CoverTab[2229]"
}

var errorRune = []byte(string(utf8.RuneError))

// ScanRunes is a split function for a Scanner that returns each
//line /snap/go/10455/src/bufio/scan.go:298
// UTF-8-encoded rune as a token. The sequence of runes returned is
//line /snap/go/10455/src/bufio/scan.go:298
// equivalent to that from a range loop over the input as a string, which
//line /snap/go/10455/src/bufio/scan.go:298
// means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
//line /snap/go/10455/src/bufio/scan.go:298
// Because of the Scan interface, this makes it impossible for the client to
//line /snap/go/10455/src/bufio/scan.go:298
// distinguish correctly encoded replacement runes from encoding errors.
//line /snap/go/10455/src/bufio/scan.go:304
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /snap/go/10455/src/bufio/scan.go:304
	_go_fuzz_dep_.CoverTab[2233]++
						if atEOF && func() bool {
//line /snap/go/10455/src/bufio/scan.go:305
		_go_fuzz_dep_.CoverTab[2238]++
//line /snap/go/10455/src/bufio/scan.go:305
		return len(data) == 0
//line /snap/go/10455/src/bufio/scan.go:305
		// _ = "end of CoverTab[2238]"
//line /snap/go/10455/src/bufio/scan.go:305
	}() {
//line /snap/go/10455/src/bufio/scan.go:305
		_go_fuzz_dep_.CoverTab[525947]++
//line /snap/go/10455/src/bufio/scan.go:305
		_go_fuzz_dep_.CoverTab[2239]++
							return 0, nil, nil
//line /snap/go/10455/src/bufio/scan.go:306
		// _ = "end of CoverTab[2239]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:307
		_go_fuzz_dep_.CoverTab[525948]++
//line /snap/go/10455/src/bufio/scan.go:307
		_go_fuzz_dep_.CoverTab[2240]++
//line /snap/go/10455/src/bufio/scan.go:307
		// _ = "end of CoverTab[2240]"
//line /snap/go/10455/src/bufio/scan.go:307
	}
//line /snap/go/10455/src/bufio/scan.go:307
	// _ = "end of CoverTab[2233]"
//line /snap/go/10455/src/bufio/scan.go:307
	_go_fuzz_dep_.CoverTab[2234]++

//line /snap/go/10455/src/bufio/scan.go:310
	if data[0] < utf8.RuneSelf {
//line /snap/go/10455/src/bufio/scan.go:310
		_go_fuzz_dep_.CoverTab[525949]++
//line /snap/go/10455/src/bufio/scan.go:310
		_go_fuzz_dep_.CoverTab[2241]++
							return 1, data[0:1], nil
//line /snap/go/10455/src/bufio/scan.go:311
		// _ = "end of CoverTab[2241]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:312
		_go_fuzz_dep_.CoverTab[525950]++
//line /snap/go/10455/src/bufio/scan.go:312
		_go_fuzz_dep_.CoverTab[2242]++
//line /snap/go/10455/src/bufio/scan.go:312
		// _ = "end of CoverTab[2242]"
//line /snap/go/10455/src/bufio/scan.go:312
	}
//line /snap/go/10455/src/bufio/scan.go:312
	// _ = "end of CoverTab[2234]"
//line /snap/go/10455/src/bufio/scan.go:312
	_go_fuzz_dep_.CoverTab[2235]++

//line /snap/go/10455/src/bufio/scan.go:315
	_, width := utf8.DecodeRune(data)
	if width > 1 {
//line /snap/go/10455/src/bufio/scan.go:316
		_go_fuzz_dep_.CoverTab[525951]++
//line /snap/go/10455/src/bufio/scan.go:316
		_go_fuzz_dep_.CoverTab[2243]++

//line /snap/go/10455/src/bufio/scan.go:319
		return width, data[0:width], nil
//line /snap/go/10455/src/bufio/scan.go:319
		// _ = "end of CoverTab[2243]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:320
		_go_fuzz_dep_.CoverTab[525952]++
//line /snap/go/10455/src/bufio/scan.go:320
		_go_fuzz_dep_.CoverTab[2244]++
//line /snap/go/10455/src/bufio/scan.go:320
		// _ = "end of CoverTab[2244]"
//line /snap/go/10455/src/bufio/scan.go:320
	}
//line /snap/go/10455/src/bufio/scan.go:320
	// _ = "end of CoverTab[2235]"
//line /snap/go/10455/src/bufio/scan.go:320
	_go_fuzz_dep_.CoverTab[2236]++

//line /snap/go/10455/src/bufio/scan.go:325
	if !atEOF && func() bool {
//line /snap/go/10455/src/bufio/scan.go:325
		_go_fuzz_dep_.CoverTab[2245]++
//line /snap/go/10455/src/bufio/scan.go:325
		return !utf8.FullRune(data)
//line /snap/go/10455/src/bufio/scan.go:325
		// _ = "end of CoverTab[2245]"
//line /snap/go/10455/src/bufio/scan.go:325
	}() {
//line /snap/go/10455/src/bufio/scan.go:325
		_go_fuzz_dep_.CoverTab[525953]++
//line /snap/go/10455/src/bufio/scan.go:325
		_go_fuzz_dep_.CoverTab[2246]++

							return 0, nil, nil
//line /snap/go/10455/src/bufio/scan.go:327
		// _ = "end of CoverTab[2246]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:328
		_go_fuzz_dep_.CoverTab[525954]++
//line /snap/go/10455/src/bufio/scan.go:328
		_go_fuzz_dep_.CoverTab[2247]++
//line /snap/go/10455/src/bufio/scan.go:328
		// _ = "end of CoverTab[2247]"
//line /snap/go/10455/src/bufio/scan.go:328
	}
//line /snap/go/10455/src/bufio/scan.go:328
	// _ = "end of CoverTab[2236]"
//line /snap/go/10455/src/bufio/scan.go:328
	_go_fuzz_dep_.CoverTab[2237]++

//line /snap/go/10455/src/bufio/scan.go:333
	return 1, errorRune, nil
//line /snap/go/10455/src/bufio/scan.go:333
	// _ = "end of CoverTab[2237]"
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
//line /snap/go/10455/src/bufio/scan.go:337
	_go_fuzz_dep_.CoverTab[2248]++
						if len(data) > 0 && func() bool {
//line /snap/go/10455/src/bufio/scan.go:338
		_go_fuzz_dep_.CoverTab[2250]++
//line /snap/go/10455/src/bufio/scan.go:338
		return data[len(data)-1] == '\r'
//line /snap/go/10455/src/bufio/scan.go:338
		// _ = "end of CoverTab[2250]"
//line /snap/go/10455/src/bufio/scan.go:338
	}() {
//line /snap/go/10455/src/bufio/scan.go:338
		_go_fuzz_dep_.CoverTab[525955]++
//line /snap/go/10455/src/bufio/scan.go:338
		_go_fuzz_dep_.CoverTab[2251]++
							return data[0 : len(data)-1]
//line /snap/go/10455/src/bufio/scan.go:339
		// _ = "end of CoverTab[2251]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:340
		_go_fuzz_dep_.CoverTab[525956]++
//line /snap/go/10455/src/bufio/scan.go:340
		_go_fuzz_dep_.CoverTab[2252]++
//line /snap/go/10455/src/bufio/scan.go:340
		// _ = "end of CoverTab[2252]"
//line /snap/go/10455/src/bufio/scan.go:340
	}
//line /snap/go/10455/src/bufio/scan.go:340
	// _ = "end of CoverTab[2248]"
//line /snap/go/10455/src/bufio/scan.go:340
	_go_fuzz_dep_.CoverTab[2249]++
						return data
//line /snap/go/10455/src/bufio/scan.go:341
	// _ = "end of CoverTab[2249]"
}

// ScanLines is a split function for a Scanner that returns each line of
//line /snap/go/10455/src/bufio/scan.go:344
// text, stripped of any trailing end-of-line marker. The returned line may
//line /snap/go/10455/src/bufio/scan.go:344
// be empty. The end-of-line marker is one optional carriage return followed
//line /snap/go/10455/src/bufio/scan.go:344
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
//line /snap/go/10455/src/bufio/scan.go:344
// The last non-empty line of input will be returned even if it has no
//line /snap/go/10455/src/bufio/scan.go:344
// newline.
//line /snap/go/10455/src/bufio/scan.go:350
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /snap/go/10455/src/bufio/scan.go:350
	_go_fuzz_dep_.CoverTab[2253]++
						if atEOF && func() bool {
//line /snap/go/10455/src/bufio/scan.go:351
		_go_fuzz_dep_.CoverTab[2257]++
//line /snap/go/10455/src/bufio/scan.go:351
		return len(data) == 0
//line /snap/go/10455/src/bufio/scan.go:351
		// _ = "end of CoverTab[2257]"
//line /snap/go/10455/src/bufio/scan.go:351
	}() {
//line /snap/go/10455/src/bufio/scan.go:351
		_go_fuzz_dep_.CoverTab[525957]++
//line /snap/go/10455/src/bufio/scan.go:351
		_go_fuzz_dep_.CoverTab[2258]++
							return 0, nil, nil
//line /snap/go/10455/src/bufio/scan.go:352
		// _ = "end of CoverTab[2258]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:353
		_go_fuzz_dep_.CoverTab[525958]++
//line /snap/go/10455/src/bufio/scan.go:353
		_go_fuzz_dep_.CoverTab[2259]++
//line /snap/go/10455/src/bufio/scan.go:353
		// _ = "end of CoverTab[2259]"
//line /snap/go/10455/src/bufio/scan.go:353
	}
//line /snap/go/10455/src/bufio/scan.go:353
	// _ = "end of CoverTab[2253]"
//line /snap/go/10455/src/bufio/scan.go:353
	_go_fuzz_dep_.CoverTab[2254]++
						if i := bytes.IndexByte(data, '\n'); i >= 0 {
//line /snap/go/10455/src/bufio/scan.go:354
		_go_fuzz_dep_.CoverTab[525959]++
//line /snap/go/10455/src/bufio/scan.go:354
		_go_fuzz_dep_.CoverTab[2260]++

							return i + 1, dropCR(data[0:i]), nil
//line /snap/go/10455/src/bufio/scan.go:356
		// _ = "end of CoverTab[2260]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:357
		_go_fuzz_dep_.CoverTab[525960]++
//line /snap/go/10455/src/bufio/scan.go:357
		_go_fuzz_dep_.CoverTab[2261]++
//line /snap/go/10455/src/bufio/scan.go:357
		// _ = "end of CoverTab[2261]"
//line /snap/go/10455/src/bufio/scan.go:357
	}
//line /snap/go/10455/src/bufio/scan.go:357
	// _ = "end of CoverTab[2254]"
//line /snap/go/10455/src/bufio/scan.go:357
	_go_fuzz_dep_.CoverTab[2255]++

						if atEOF {
//line /snap/go/10455/src/bufio/scan.go:359
		_go_fuzz_dep_.CoverTab[525961]++
//line /snap/go/10455/src/bufio/scan.go:359
		_go_fuzz_dep_.CoverTab[2262]++
							return len(data), dropCR(data), nil
//line /snap/go/10455/src/bufio/scan.go:360
		// _ = "end of CoverTab[2262]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:361
		_go_fuzz_dep_.CoverTab[525962]++
//line /snap/go/10455/src/bufio/scan.go:361
		_go_fuzz_dep_.CoverTab[2263]++
//line /snap/go/10455/src/bufio/scan.go:361
		// _ = "end of CoverTab[2263]"
//line /snap/go/10455/src/bufio/scan.go:361
	}
//line /snap/go/10455/src/bufio/scan.go:361
	// _ = "end of CoverTab[2255]"
//line /snap/go/10455/src/bufio/scan.go:361
	_go_fuzz_dep_.CoverTab[2256]++

						return 0, nil, nil
//line /snap/go/10455/src/bufio/scan.go:363
	// _ = "end of CoverTab[2256]"
}

// isSpace reports whether the character is a Unicode white space character.
//line /snap/go/10455/src/bufio/scan.go:366
// We avoid dependency on the unicode package, but check validity of the implementation
//line /snap/go/10455/src/bufio/scan.go:366
// in the tests.
//line /snap/go/10455/src/bufio/scan.go:369
func isSpace(r rune) bool {
//line /snap/go/10455/src/bufio/scan.go:369
	_go_fuzz_dep_.CoverTab[2264]++
						if r <= '\u00FF' {
//line /snap/go/10455/src/bufio/scan.go:370
		_go_fuzz_dep_.CoverTab[525963]++
//line /snap/go/10455/src/bufio/scan.go:370
		_go_fuzz_dep_.CoverTab[2268]++

							switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
//line /snap/go/10455/src/bufio/scan.go:373
			_go_fuzz_dep_.CoverTab[525965]++
//line /snap/go/10455/src/bufio/scan.go:373
			_go_fuzz_dep_.CoverTab[2270]++
								return true
//line /snap/go/10455/src/bufio/scan.go:374
			// _ = "end of CoverTab[2270]"
		case '\u0085', '\u00A0':
//line /snap/go/10455/src/bufio/scan.go:375
			_go_fuzz_dep_.CoverTab[525966]++
//line /snap/go/10455/src/bufio/scan.go:375
			_go_fuzz_dep_.CoverTab[2271]++
								return true
//line /snap/go/10455/src/bufio/scan.go:376
			// _ = "end of CoverTab[2271]"
//line /snap/go/10455/src/bufio/scan.go:376
		default:
//line /snap/go/10455/src/bufio/scan.go:376
			_go_fuzz_dep_.CoverTab[525967]++
//line /snap/go/10455/src/bufio/scan.go:376
			_go_fuzz_dep_.CoverTab[2272]++
//line /snap/go/10455/src/bufio/scan.go:376
			// _ = "end of CoverTab[2272]"
		}
//line /snap/go/10455/src/bufio/scan.go:377
		// _ = "end of CoverTab[2268]"
//line /snap/go/10455/src/bufio/scan.go:377
		_go_fuzz_dep_.CoverTab[2269]++
							return false
//line /snap/go/10455/src/bufio/scan.go:378
		// _ = "end of CoverTab[2269]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:379
		_go_fuzz_dep_.CoverTab[525964]++
//line /snap/go/10455/src/bufio/scan.go:379
		_go_fuzz_dep_.CoverTab[2273]++
//line /snap/go/10455/src/bufio/scan.go:379
		// _ = "end of CoverTab[2273]"
//line /snap/go/10455/src/bufio/scan.go:379
	}
//line /snap/go/10455/src/bufio/scan.go:379
	// _ = "end of CoverTab[2264]"
//line /snap/go/10455/src/bufio/scan.go:379
	_go_fuzz_dep_.CoverTab[2265]++

						if '\u2000' <= r && func() bool {
//line /snap/go/10455/src/bufio/scan.go:381
		_go_fuzz_dep_.CoverTab[2274]++
//line /snap/go/10455/src/bufio/scan.go:381
		return r <= '\u200a'
//line /snap/go/10455/src/bufio/scan.go:381
		// _ = "end of CoverTab[2274]"
//line /snap/go/10455/src/bufio/scan.go:381
	}() {
//line /snap/go/10455/src/bufio/scan.go:381
		_go_fuzz_dep_.CoverTab[525968]++
//line /snap/go/10455/src/bufio/scan.go:381
		_go_fuzz_dep_.CoverTab[2275]++
							return true
//line /snap/go/10455/src/bufio/scan.go:382
		// _ = "end of CoverTab[2275]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:383
		_go_fuzz_dep_.CoverTab[525969]++
//line /snap/go/10455/src/bufio/scan.go:383
		_go_fuzz_dep_.CoverTab[2276]++
//line /snap/go/10455/src/bufio/scan.go:383
		// _ = "end of CoverTab[2276]"
//line /snap/go/10455/src/bufio/scan.go:383
	}
//line /snap/go/10455/src/bufio/scan.go:383
	// _ = "end of CoverTab[2265]"
//line /snap/go/10455/src/bufio/scan.go:383
	_go_fuzz_dep_.CoverTab[2266]++
						switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
//line /snap/go/10455/src/bufio/scan.go:385
		_go_fuzz_dep_.CoverTab[525970]++
//line /snap/go/10455/src/bufio/scan.go:385
		_go_fuzz_dep_.CoverTab[2277]++
							return true
//line /snap/go/10455/src/bufio/scan.go:386
		// _ = "end of CoverTab[2277]"
//line /snap/go/10455/src/bufio/scan.go:386
	default:
//line /snap/go/10455/src/bufio/scan.go:386
		_go_fuzz_dep_.CoverTab[525971]++
//line /snap/go/10455/src/bufio/scan.go:386
		_go_fuzz_dep_.CoverTab[2278]++
//line /snap/go/10455/src/bufio/scan.go:386
		// _ = "end of CoverTab[2278]"
	}
//line /snap/go/10455/src/bufio/scan.go:387
	// _ = "end of CoverTab[2266]"
//line /snap/go/10455/src/bufio/scan.go:387
	_go_fuzz_dep_.CoverTab[2267]++
						return false
//line /snap/go/10455/src/bufio/scan.go:388
	// _ = "end of CoverTab[2267]"
}

// ScanWords is a split function for a Scanner that returns each
//line /snap/go/10455/src/bufio/scan.go:391
// space-separated word of text, with surrounding spaces deleted. It will
//line /snap/go/10455/src/bufio/scan.go:391
// never return an empty string. The definition of space is set by
//line /snap/go/10455/src/bufio/scan.go:391
// unicode.IsSpace.
//line /snap/go/10455/src/bufio/scan.go:395
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /snap/go/10455/src/bufio/scan.go:395
	_go_fuzz_dep_.CoverTab[2279]++

						start := 0
//line /snap/go/10455/src/bufio/scan.go:397
	_go_fuzz_dep_.CoverTab[786570] = 0
						for width := 0; start < len(data); start += width {
//line /snap/go/10455/src/bufio/scan.go:398
		if _go_fuzz_dep_.CoverTab[786570] == 0 {
//line /snap/go/10455/src/bufio/scan.go:398
			_go_fuzz_dep_.CoverTab[525986]++
//line /snap/go/10455/src/bufio/scan.go:398
		} else {
//line /snap/go/10455/src/bufio/scan.go:398
			_go_fuzz_dep_.CoverTab[525987]++
//line /snap/go/10455/src/bufio/scan.go:398
		}
//line /snap/go/10455/src/bufio/scan.go:398
		_go_fuzz_dep_.CoverTab[786570] = 1
//line /snap/go/10455/src/bufio/scan.go:398
		_go_fuzz_dep_.CoverTab[2283]++
							var r rune
							r, width = utf8.DecodeRune(data[start:])
							if !isSpace(r) {
//line /snap/go/10455/src/bufio/scan.go:401
			_go_fuzz_dep_.CoverTab[525972]++
//line /snap/go/10455/src/bufio/scan.go:401
			_go_fuzz_dep_.CoverTab[2284]++
								break
//line /snap/go/10455/src/bufio/scan.go:402
			// _ = "end of CoverTab[2284]"
		} else {
//line /snap/go/10455/src/bufio/scan.go:403
			_go_fuzz_dep_.CoverTab[525973]++
//line /snap/go/10455/src/bufio/scan.go:403
			_go_fuzz_dep_.CoverTab[2285]++
//line /snap/go/10455/src/bufio/scan.go:403
			// _ = "end of CoverTab[2285]"
//line /snap/go/10455/src/bufio/scan.go:403
		}
//line /snap/go/10455/src/bufio/scan.go:403
		// _ = "end of CoverTab[2283]"
	}
//line /snap/go/10455/src/bufio/scan.go:404
	if _go_fuzz_dep_.CoverTab[786570] == 0 {
//line /snap/go/10455/src/bufio/scan.go:404
		_go_fuzz_dep_.CoverTab[525988]++
//line /snap/go/10455/src/bufio/scan.go:404
	} else {
//line /snap/go/10455/src/bufio/scan.go:404
		_go_fuzz_dep_.CoverTab[525989]++
//line /snap/go/10455/src/bufio/scan.go:404
	}
//line /snap/go/10455/src/bufio/scan.go:404
	// _ = "end of CoverTab[2279]"
//line /snap/go/10455/src/bufio/scan.go:404
	_go_fuzz_dep_.CoverTab[2280]++
//line /snap/go/10455/src/bufio/scan.go:404
	_go_fuzz_dep_.CoverTab[786571] = 0

						for width, i := 0, start; i < len(data); i += width {
//line /snap/go/10455/src/bufio/scan.go:406
		if _go_fuzz_dep_.CoverTab[786571] == 0 {
//line /snap/go/10455/src/bufio/scan.go:406
			_go_fuzz_dep_.CoverTab[525990]++
//line /snap/go/10455/src/bufio/scan.go:406
		} else {
//line /snap/go/10455/src/bufio/scan.go:406
			_go_fuzz_dep_.CoverTab[525991]++
//line /snap/go/10455/src/bufio/scan.go:406
		}
//line /snap/go/10455/src/bufio/scan.go:406
		_go_fuzz_dep_.CoverTab[786571] = 1
//line /snap/go/10455/src/bufio/scan.go:406
		_go_fuzz_dep_.CoverTab[2286]++
							var r rune
							r, width = utf8.DecodeRune(data[i:])
							if isSpace(r) {
//line /snap/go/10455/src/bufio/scan.go:409
			_go_fuzz_dep_.CoverTab[525974]++
//line /snap/go/10455/src/bufio/scan.go:409
			_go_fuzz_dep_.CoverTab[2287]++
								return i + width, data[start:i], nil
//line /snap/go/10455/src/bufio/scan.go:410
			// _ = "end of CoverTab[2287]"
		} else {
//line /snap/go/10455/src/bufio/scan.go:411
			_go_fuzz_dep_.CoverTab[525975]++
//line /snap/go/10455/src/bufio/scan.go:411
			_go_fuzz_dep_.CoverTab[2288]++
//line /snap/go/10455/src/bufio/scan.go:411
			// _ = "end of CoverTab[2288]"
//line /snap/go/10455/src/bufio/scan.go:411
		}
//line /snap/go/10455/src/bufio/scan.go:411
		// _ = "end of CoverTab[2286]"
	}
//line /snap/go/10455/src/bufio/scan.go:412
	if _go_fuzz_dep_.CoverTab[786571] == 0 {
//line /snap/go/10455/src/bufio/scan.go:412
		_go_fuzz_dep_.CoverTab[525992]++
//line /snap/go/10455/src/bufio/scan.go:412
	} else {
//line /snap/go/10455/src/bufio/scan.go:412
		_go_fuzz_dep_.CoverTab[525993]++
//line /snap/go/10455/src/bufio/scan.go:412
	}
//line /snap/go/10455/src/bufio/scan.go:412
	// _ = "end of CoverTab[2280]"
//line /snap/go/10455/src/bufio/scan.go:412
	_go_fuzz_dep_.CoverTab[2281]++

						if atEOF && func() bool {
//line /snap/go/10455/src/bufio/scan.go:414
		_go_fuzz_dep_.CoverTab[2289]++
//line /snap/go/10455/src/bufio/scan.go:414
		return len(data) > start
//line /snap/go/10455/src/bufio/scan.go:414
		// _ = "end of CoverTab[2289]"
//line /snap/go/10455/src/bufio/scan.go:414
	}() {
//line /snap/go/10455/src/bufio/scan.go:414
		_go_fuzz_dep_.CoverTab[525976]++
//line /snap/go/10455/src/bufio/scan.go:414
		_go_fuzz_dep_.CoverTab[2290]++
							return len(data), data[start:], nil
//line /snap/go/10455/src/bufio/scan.go:415
		// _ = "end of CoverTab[2290]"
	} else {
//line /snap/go/10455/src/bufio/scan.go:416
		_go_fuzz_dep_.CoverTab[525977]++
//line /snap/go/10455/src/bufio/scan.go:416
		_go_fuzz_dep_.CoverTab[2291]++
//line /snap/go/10455/src/bufio/scan.go:416
		// _ = "end of CoverTab[2291]"
//line /snap/go/10455/src/bufio/scan.go:416
	}
//line /snap/go/10455/src/bufio/scan.go:416
	// _ = "end of CoverTab[2281]"
//line /snap/go/10455/src/bufio/scan.go:416
	_go_fuzz_dep_.CoverTab[2282]++

						return start, nil, nil
//line /snap/go/10455/src/bufio/scan.go:418
	// _ = "end of CoverTab[2282]"
}

//line /snap/go/10455/src/bufio/scan.go:419
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/bufio/scan.go:419
var _ = _go_fuzz_dep_.CoverTab
