// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/bufio/scan.go:5
package bufio

//line /usr/local/go/src/bufio/scan.go:5
import (
//line /usr/local/go/src/bufio/scan.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/bufio/scan.go:5
)
//line /usr/local/go/src/bufio/scan.go:5
import (
//line /usr/local/go/src/bufio/scan.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/bufio/scan.go:5
)

import (
	"bytes"
	"errors"
	"io"
	"unicode/utf8"
)

// Scanner provides a convenient interface for reading data such as
//line /usr/local/go/src/bufio/scan.go:14
// a file of newline-delimited lines of text. Successive calls to
//line /usr/local/go/src/bufio/scan.go:14
// the Scan method will step through the 'tokens' of a file, skipping
//line /usr/local/go/src/bufio/scan.go:14
// the bytes between the tokens. The specification of a token is
//line /usr/local/go/src/bufio/scan.go:14
// defined by a split function of type SplitFunc; the default split
//line /usr/local/go/src/bufio/scan.go:14
// function breaks the input into lines with line termination stripped. Split
//line /usr/local/go/src/bufio/scan.go:14
// functions are defined in this package for scanning a file into
//line /usr/local/go/src/bufio/scan.go:14
// lines, bytes, UTF-8-encoded runes, and space-delimited words. The
//line /usr/local/go/src/bufio/scan.go:14
// client may instead provide a custom split function.
//line /usr/local/go/src/bufio/scan.go:14
//
//line /usr/local/go/src/bufio/scan.go:14
// Scanning stops unrecoverably at EOF, the first I/O error, or a token too
//line /usr/local/go/src/bufio/scan.go:14
// large to fit in the buffer. When a scan stops, the reader may have
//line /usr/local/go/src/bufio/scan.go:14
// advanced arbitrarily far past the last token. Programs that need more
//line /usr/local/go/src/bufio/scan.go:14
// control over error handling or large tokens, or must run sequential scans
//line /usr/local/go/src/bufio/scan.go:14
// on a reader, should use bufio.Reader instead.
//line /usr/local/go/src/bufio/scan.go:29
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
//line /usr/local/go/src/bufio/scan.go:43
// input. The arguments are an initial substring of the remaining unprocessed
//line /usr/local/go/src/bufio/scan.go:43
// data and a flag, atEOF, that reports whether the Reader has no more data
//line /usr/local/go/src/bufio/scan.go:43
// to give. The return values are the number of bytes to advance the input
//line /usr/local/go/src/bufio/scan.go:43
// and the next token to return to the user, if any, plus an error, if any.
//line /usr/local/go/src/bufio/scan.go:43
//
//line /usr/local/go/src/bufio/scan.go:43
// Scanning stops if the function returns an error, in which case some of
//line /usr/local/go/src/bufio/scan.go:43
// the input may be discarded. If that error is ErrFinalToken, scanning
//line /usr/local/go/src/bufio/scan.go:43
// stops with no error.
//line /usr/local/go/src/bufio/scan.go:43
//
//line /usr/local/go/src/bufio/scan.go:43
// Otherwise, the Scanner advances the input. If the token is not nil,
//line /usr/local/go/src/bufio/scan.go:43
// the Scanner returns it to the user. If the token is nil, the
//line /usr/local/go/src/bufio/scan.go:43
// Scanner reads more data and continues scanning; if there is no more
//line /usr/local/go/src/bufio/scan.go:43
// data--if atEOF was true--the Scanner returns. If the data does not
//line /usr/local/go/src/bufio/scan.go:43
// yet hold a complete token, for instance if it has no newline while
//line /usr/local/go/src/bufio/scan.go:43
// scanning lines, a SplitFunc can return (0, nil, nil) to signal the
//line /usr/local/go/src/bufio/scan.go:43
// Scanner to read more data into the slice and try again with a
//line /usr/local/go/src/bufio/scan.go:43
// longer slice starting at the same point in the input.
//line /usr/local/go/src/bufio/scan.go:43
//
//line /usr/local/go/src/bufio/scan.go:43
// The function is never called with an empty data slice unless atEOF
//line /usr/local/go/src/bufio/scan.go:43
// is true. If atEOF is true, however, data may be non-empty and,
//line /usr/local/go/src/bufio/scan.go:43
// as always, holds unprocessed text.
//line /usr/local/go/src/bufio/scan.go:65
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
//line /usr/local/go/src/bufio/scan.go:85
// The split function defaults to ScanLines.
//line /usr/local/go/src/bufio/scan.go:87
func NewScanner(r io.Reader) *Scanner {
//line /usr/local/go/src/bufio/scan.go:87
	_go_fuzz_dep_.CoverTab[25521]++
						return &Scanner{
		r:		r,
		split:		ScanLines,
		maxTokenSize:	MaxScanTokenSize,
	}
//line /usr/local/go/src/bufio/scan.go:92
	// _ = "end of CoverTab[25521]"
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (s *Scanner) Err() error {
//line /usr/local/go/src/bufio/scan.go:96
	_go_fuzz_dep_.CoverTab[25522]++
						if s.err == io.EOF {
//line /usr/local/go/src/bufio/scan.go:97
		_go_fuzz_dep_.CoverTab[25524]++
							return nil
//line /usr/local/go/src/bufio/scan.go:98
		// _ = "end of CoverTab[25524]"
	} else {
//line /usr/local/go/src/bufio/scan.go:99
		_go_fuzz_dep_.CoverTab[25525]++
//line /usr/local/go/src/bufio/scan.go:99
		// _ = "end of CoverTab[25525]"
//line /usr/local/go/src/bufio/scan.go:99
	}
//line /usr/local/go/src/bufio/scan.go:99
	// _ = "end of CoverTab[25522]"
//line /usr/local/go/src/bufio/scan.go:99
	_go_fuzz_dep_.CoverTab[25523]++
						return s.err
//line /usr/local/go/src/bufio/scan.go:100
	// _ = "end of CoverTab[25523]"
}

// Bytes returns the most recent token generated by a call to Scan.
//line /usr/local/go/src/bufio/scan.go:103
// The underlying array may point to data that will be overwritten
//line /usr/local/go/src/bufio/scan.go:103
// by a subsequent call to Scan. It does no allocation.
//line /usr/local/go/src/bufio/scan.go:106
func (s *Scanner) Bytes() []byte {
//line /usr/local/go/src/bufio/scan.go:106
	_go_fuzz_dep_.CoverTab[25526]++
						return s.token
//line /usr/local/go/src/bufio/scan.go:107
	// _ = "end of CoverTab[25526]"
}

// Text returns the most recent token generated by a call to Scan
//line /usr/local/go/src/bufio/scan.go:110
// as a newly allocated string holding its bytes.
//line /usr/local/go/src/bufio/scan.go:112
func (s *Scanner) Text() string {
//line /usr/local/go/src/bufio/scan.go:112
	_go_fuzz_dep_.CoverTab[25527]++
						return string(s.token)
//line /usr/local/go/src/bufio/scan.go:113
	// _ = "end of CoverTab[25527]"
}

// ErrFinalToken is a special sentinel error value. It is intended to be
//line /usr/local/go/src/bufio/scan.go:116
// returned by a Split function to indicate that the token being delivered
//line /usr/local/go/src/bufio/scan.go:116
// with the error is the last token and scanning should stop after this one.
//line /usr/local/go/src/bufio/scan.go:116
// After ErrFinalToken is received by Scan, scanning stops with no error.
//line /usr/local/go/src/bufio/scan.go:116
// The value is useful to stop processing early or when it is necessary to
//line /usr/local/go/src/bufio/scan.go:116
// deliver a final empty token. One could achieve the same behavior
//line /usr/local/go/src/bufio/scan.go:116
// with a custom error value but providing one here is tidier.
//line /usr/local/go/src/bufio/scan.go:116
// See the emptyFinalToken example for a use of this value.
//line /usr/local/go/src/bufio/scan.go:124
var ErrFinalToken = errors.New("final token")

// Scan advances the Scanner to the next token, which will then be
//line /usr/local/go/src/bufio/scan.go:126
// available through the Bytes or Text method. It returns false when the
//line /usr/local/go/src/bufio/scan.go:126
// scan stops, either by reaching the end of the input or an error.
//line /usr/local/go/src/bufio/scan.go:126
// After Scan returns false, the Err method will return any error that
//line /usr/local/go/src/bufio/scan.go:126
// occurred during scanning, except that if it was io.EOF, Err
//line /usr/local/go/src/bufio/scan.go:126
// will return nil.
//line /usr/local/go/src/bufio/scan.go:126
// Scan panics if the split function returns too many empty
//line /usr/local/go/src/bufio/scan.go:126
// tokens without advancing the input. This is a common error mode for
//line /usr/local/go/src/bufio/scan.go:126
// scanners.
//line /usr/local/go/src/bufio/scan.go:135
func (s *Scanner) Scan() bool {
//line /usr/local/go/src/bufio/scan.go:135
	_go_fuzz_dep_.CoverTab[25528]++
						if s.done {
//line /usr/local/go/src/bufio/scan.go:136
		_go_fuzz_dep_.CoverTab[25530]++
							return false
//line /usr/local/go/src/bufio/scan.go:137
		// _ = "end of CoverTab[25530]"
	} else {
//line /usr/local/go/src/bufio/scan.go:138
		_go_fuzz_dep_.CoverTab[25531]++
//line /usr/local/go/src/bufio/scan.go:138
		// _ = "end of CoverTab[25531]"
//line /usr/local/go/src/bufio/scan.go:138
	}
//line /usr/local/go/src/bufio/scan.go:138
	// _ = "end of CoverTab[25528]"
//line /usr/local/go/src/bufio/scan.go:138
	_go_fuzz_dep_.CoverTab[25529]++
						s.scanCalled = true

						for {
//line /usr/local/go/src/bufio/scan.go:141
		_go_fuzz_dep_.CoverTab[25532]++

//line /usr/local/go/src/bufio/scan.go:145
		if s.end > s.start || func() bool {
//line /usr/local/go/src/bufio/scan.go:145
			_go_fuzz_dep_.CoverTab[25537]++
//line /usr/local/go/src/bufio/scan.go:145
			return s.err != nil
//line /usr/local/go/src/bufio/scan.go:145
			// _ = "end of CoverTab[25537]"
//line /usr/local/go/src/bufio/scan.go:145
		}() {
//line /usr/local/go/src/bufio/scan.go:145
			_go_fuzz_dep_.CoverTab[25538]++
								advance, token, err := s.split(s.buf[s.start:s.end], s.err != nil)
								if err != nil {
//line /usr/local/go/src/bufio/scan.go:147
				_go_fuzz_dep_.CoverTab[25541]++
									if err == ErrFinalToken {
//line /usr/local/go/src/bufio/scan.go:148
					_go_fuzz_dep_.CoverTab[25543]++
										s.token = token
										s.done = true
										return true
//line /usr/local/go/src/bufio/scan.go:151
					// _ = "end of CoverTab[25543]"
				} else {
//line /usr/local/go/src/bufio/scan.go:152
					_go_fuzz_dep_.CoverTab[25544]++
//line /usr/local/go/src/bufio/scan.go:152
					// _ = "end of CoverTab[25544]"
//line /usr/local/go/src/bufio/scan.go:152
				}
//line /usr/local/go/src/bufio/scan.go:152
				// _ = "end of CoverTab[25541]"
//line /usr/local/go/src/bufio/scan.go:152
				_go_fuzz_dep_.CoverTab[25542]++
									s.setErr(err)
									return false
//line /usr/local/go/src/bufio/scan.go:154
				// _ = "end of CoverTab[25542]"
			} else {
//line /usr/local/go/src/bufio/scan.go:155
				_go_fuzz_dep_.CoverTab[25545]++
//line /usr/local/go/src/bufio/scan.go:155
				// _ = "end of CoverTab[25545]"
//line /usr/local/go/src/bufio/scan.go:155
			}
//line /usr/local/go/src/bufio/scan.go:155
			// _ = "end of CoverTab[25538]"
//line /usr/local/go/src/bufio/scan.go:155
			_go_fuzz_dep_.CoverTab[25539]++
								if !s.advance(advance) {
//line /usr/local/go/src/bufio/scan.go:156
				_go_fuzz_dep_.CoverTab[25546]++
									return false
//line /usr/local/go/src/bufio/scan.go:157
				// _ = "end of CoverTab[25546]"
			} else {
//line /usr/local/go/src/bufio/scan.go:158
				_go_fuzz_dep_.CoverTab[25547]++
//line /usr/local/go/src/bufio/scan.go:158
				// _ = "end of CoverTab[25547]"
//line /usr/local/go/src/bufio/scan.go:158
			}
//line /usr/local/go/src/bufio/scan.go:158
			// _ = "end of CoverTab[25539]"
//line /usr/local/go/src/bufio/scan.go:158
			_go_fuzz_dep_.CoverTab[25540]++
								s.token = token
								if token != nil {
//line /usr/local/go/src/bufio/scan.go:160
				_go_fuzz_dep_.CoverTab[25548]++
									if s.err == nil || func() bool {
//line /usr/local/go/src/bufio/scan.go:161
					_go_fuzz_dep_.CoverTab[25550]++
//line /usr/local/go/src/bufio/scan.go:161
					return advance > 0
//line /usr/local/go/src/bufio/scan.go:161
					// _ = "end of CoverTab[25550]"
//line /usr/local/go/src/bufio/scan.go:161
				}() {
//line /usr/local/go/src/bufio/scan.go:161
					_go_fuzz_dep_.CoverTab[25551]++
										s.empties = 0
//line /usr/local/go/src/bufio/scan.go:162
					// _ = "end of CoverTab[25551]"
				} else {
//line /usr/local/go/src/bufio/scan.go:163
					_go_fuzz_dep_.CoverTab[25552]++

										s.empties++
										if s.empties > maxConsecutiveEmptyReads {
//line /usr/local/go/src/bufio/scan.go:166
						_go_fuzz_dep_.CoverTab[25553]++
											panic("bufio.Scan: too many empty tokens without progressing")
//line /usr/local/go/src/bufio/scan.go:167
						// _ = "end of CoverTab[25553]"
					} else {
//line /usr/local/go/src/bufio/scan.go:168
						_go_fuzz_dep_.CoverTab[25554]++
//line /usr/local/go/src/bufio/scan.go:168
						// _ = "end of CoverTab[25554]"
//line /usr/local/go/src/bufio/scan.go:168
					}
//line /usr/local/go/src/bufio/scan.go:168
					// _ = "end of CoverTab[25552]"
				}
//line /usr/local/go/src/bufio/scan.go:169
				// _ = "end of CoverTab[25548]"
//line /usr/local/go/src/bufio/scan.go:169
				_go_fuzz_dep_.CoverTab[25549]++
									return true
//line /usr/local/go/src/bufio/scan.go:170
				// _ = "end of CoverTab[25549]"
			} else {
//line /usr/local/go/src/bufio/scan.go:171
				_go_fuzz_dep_.CoverTab[25555]++
//line /usr/local/go/src/bufio/scan.go:171
				// _ = "end of CoverTab[25555]"
//line /usr/local/go/src/bufio/scan.go:171
			}
//line /usr/local/go/src/bufio/scan.go:171
			// _ = "end of CoverTab[25540]"
		} else {
//line /usr/local/go/src/bufio/scan.go:172
			_go_fuzz_dep_.CoverTab[25556]++
//line /usr/local/go/src/bufio/scan.go:172
			// _ = "end of CoverTab[25556]"
//line /usr/local/go/src/bufio/scan.go:172
		}
//line /usr/local/go/src/bufio/scan.go:172
		// _ = "end of CoverTab[25532]"
//line /usr/local/go/src/bufio/scan.go:172
		_go_fuzz_dep_.CoverTab[25533]++

//line /usr/local/go/src/bufio/scan.go:175
		if s.err != nil {
//line /usr/local/go/src/bufio/scan.go:175
			_go_fuzz_dep_.CoverTab[25557]++

								s.start = 0
								s.end = 0
								return false
//line /usr/local/go/src/bufio/scan.go:179
			// _ = "end of CoverTab[25557]"
		} else {
//line /usr/local/go/src/bufio/scan.go:180
			_go_fuzz_dep_.CoverTab[25558]++
//line /usr/local/go/src/bufio/scan.go:180
			// _ = "end of CoverTab[25558]"
//line /usr/local/go/src/bufio/scan.go:180
		}
//line /usr/local/go/src/bufio/scan.go:180
		// _ = "end of CoverTab[25533]"
//line /usr/local/go/src/bufio/scan.go:180
		_go_fuzz_dep_.CoverTab[25534]++

//line /usr/local/go/src/bufio/scan.go:184
		if s.start > 0 && func() bool {
//line /usr/local/go/src/bufio/scan.go:184
			_go_fuzz_dep_.CoverTab[25559]++
//line /usr/local/go/src/bufio/scan.go:184
			return (s.end == len(s.buf) || func() bool {
//line /usr/local/go/src/bufio/scan.go:184
				_go_fuzz_dep_.CoverTab[25560]++
//line /usr/local/go/src/bufio/scan.go:184
				return s.start > len(s.buf)/2
//line /usr/local/go/src/bufio/scan.go:184
				// _ = "end of CoverTab[25560]"
//line /usr/local/go/src/bufio/scan.go:184
			}())
//line /usr/local/go/src/bufio/scan.go:184
			// _ = "end of CoverTab[25559]"
//line /usr/local/go/src/bufio/scan.go:184
		}() {
//line /usr/local/go/src/bufio/scan.go:184
			_go_fuzz_dep_.CoverTab[25561]++
								copy(s.buf, s.buf[s.start:s.end])
								s.end -= s.start
								s.start = 0
//line /usr/local/go/src/bufio/scan.go:187
			// _ = "end of CoverTab[25561]"
		} else {
//line /usr/local/go/src/bufio/scan.go:188
			_go_fuzz_dep_.CoverTab[25562]++
//line /usr/local/go/src/bufio/scan.go:188
			// _ = "end of CoverTab[25562]"
//line /usr/local/go/src/bufio/scan.go:188
		}
//line /usr/local/go/src/bufio/scan.go:188
		// _ = "end of CoverTab[25534]"
//line /usr/local/go/src/bufio/scan.go:188
		_go_fuzz_dep_.CoverTab[25535]++

							if s.end == len(s.buf) {
//line /usr/local/go/src/bufio/scan.go:190
			_go_fuzz_dep_.CoverTab[25563]++
			// Guarantee no overflow in the multiplication below.
			const maxInt = int(^uint(0) >> 1)
			if len(s.buf) >= s.maxTokenSize || func() bool {
//line /usr/local/go/src/bufio/scan.go:193
				_go_fuzz_dep_.CoverTab[25567]++
//line /usr/local/go/src/bufio/scan.go:193
				return len(s.buf) > maxInt/2
//line /usr/local/go/src/bufio/scan.go:193
				// _ = "end of CoverTab[25567]"
//line /usr/local/go/src/bufio/scan.go:193
			}() {
//line /usr/local/go/src/bufio/scan.go:193
				_go_fuzz_dep_.CoverTab[25568]++
									s.setErr(ErrTooLong)
									return false
//line /usr/local/go/src/bufio/scan.go:195
				// _ = "end of CoverTab[25568]"
			} else {
//line /usr/local/go/src/bufio/scan.go:196
				_go_fuzz_dep_.CoverTab[25569]++
//line /usr/local/go/src/bufio/scan.go:196
				// _ = "end of CoverTab[25569]"
//line /usr/local/go/src/bufio/scan.go:196
			}
//line /usr/local/go/src/bufio/scan.go:196
			// _ = "end of CoverTab[25563]"
//line /usr/local/go/src/bufio/scan.go:196
			_go_fuzz_dep_.CoverTab[25564]++
								newSize := len(s.buf) * 2
								if newSize == 0 {
//line /usr/local/go/src/bufio/scan.go:198
				_go_fuzz_dep_.CoverTab[25570]++
									newSize = startBufSize
//line /usr/local/go/src/bufio/scan.go:199
				// _ = "end of CoverTab[25570]"
			} else {
//line /usr/local/go/src/bufio/scan.go:200
				_go_fuzz_dep_.CoverTab[25571]++
//line /usr/local/go/src/bufio/scan.go:200
				// _ = "end of CoverTab[25571]"
//line /usr/local/go/src/bufio/scan.go:200
			}
//line /usr/local/go/src/bufio/scan.go:200
			// _ = "end of CoverTab[25564]"
//line /usr/local/go/src/bufio/scan.go:200
			_go_fuzz_dep_.CoverTab[25565]++
								if newSize > s.maxTokenSize {
//line /usr/local/go/src/bufio/scan.go:201
				_go_fuzz_dep_.CoverTab[25572]++
									newSize = s.maxTokenSize
//line /usr/local/go/src/bufio/scan.go:202
				// _ = "end of CoverTab[25572]"
			} else {
//line /usr/local/go/src/bufio/scan.go:203
				_go_fuzz_dep_.CoverTab[25573]++
//line /usr/local/go/src/bufio/scan.go:203
				// _ = "end of CoverTab[25573]"
//line /usr/local/go/src/bufio/scan.go:203
			}
//line /usr/local/go/src/bufio/scan.go:203
			// _ = "end of CoverTab[25565]"
//line /usr/local/go/src/bufio/scan.go:203
			_go_fuzz_dep_.CoverTab[25566]++
								newBuf := make([]byte, newSize)
								copy(newBuf, s.buf[s.start:s.end])
								s.buf = newBuf
								s.end -= s.start
								s.start = 0
//line /usr/local/go/src/bufio/scan.go:208
			// _ = "end of CoverTab[25566]"
		} else {
//line /usr/local/go/src/bufio/scan.go:209
			_go_fuzz_dep_.CoverTab[25574]++
//line /usr/local/go/src/bufio/scan.go:209
			// _ = "end of CoverTab[25574]"
//line /usr/local/go/src/bufio/scan.go:209
		}
//line /usr/local/go/src/bufio/scan.go:209
		// _ = "end of CoverTab[25535]"
//line /usr/local/go/src/bufio/scan.go:209
		_go_fuzz_dep_.CoverTab[25536]++

//line /usr/local/go/src/bufio/scan.go:213
		for loop := 0; ; {
//line /usr/local/go/src/bufio/scan.go:213
			_go_fuzz_dep_.CoverTab[25575]++
								n, err := s.r.Read(s.buf[s.end:len(s.buf)])
								if n < 0 || func() bool {
//line /usr/local/go/src/bufio/scan.go:215
				_go_fuzz_dep_.CoverTab[25579]++
//line /usr/local/go/src/bufio/scan.go:215
				return len(s.buf)-s.end < n
//line /usr/local/go/src/bufio/scan.go:215
				// _ = "end of CoverTab[25579]"
//line /usr/local/go/src/bufio/scan.go:215
			}() {
//line /usr/local/go/src/bufio/scan.go:215
				_go_fuzz_dep_.CoverTab[25580]++
									s.setErr(ErrBadReadCount)
									break
//line /usr/local/go/src/bufio/scan.go:217
				// _ = "end of CoverTab[25580]"
			} else {
//line /usr/local/go/src/bufio/scan.go:218
				_go_fuzz_dep_.CoverTab[25581]++
//line /usr/local/go/src/bufio/scan.go:218
				// _ = "end of CoverTab[25581]"
//line /usr/local/go/src/bufio/scan.go:218
			}
//line /usr/local/go/src/bufio/scan.go:218
			// _ = "end of CoverTab[25575]"
//line /usr/local/go/src/bufio/scan.go:218
			_go_fuzz_dep_.CoverTab[25576]++
								s.end += n
								if err != nil {
//line /usr/local/go/src/bufio/scan.go:220
				_go_fuzz_dep_.CoverTab[25582]++
									s.setErr(err)
									break
//line /usr/local/go/src/bufio/scan.go:222
				// _ = "end of CoverTab[25582]"
			} else {
//line /usr/local/go/src/bufio/scan.go:223
				_go_fuzz_dep_.CoverTab[25583]++
//line /usr/local/go/src/bufio/scan.go:223
				// _ = "end of CoverTab[25583]"
//line /usr/local/go/src/bufio/scan.go:223
			}
//line /usr/local/go/src/bufio/scan.go:223
			// _ = "end of CoverTab[25576]"
//line /usr/local/go/src/bufio/scan.go:223
			_go_fuzz_dep_.CoverTab[25577]++
								if n > 0 {
//line /usr/local/go/src/bufio/scan.go:224
				_go_fuzz_dep_.CoverTab[25584]++
									s.empties = 0
									break
//line /usr/local/go/src/bufio/scan.go:226
				// _ = "end of CoverTab[25584]"
			} else {
//line /usr/local/go/src/bufio/scan.go:227
				_go_fuzz_dep_.CoverTab[25585]++
//line /usr/local/go/src/bufio/scan.go:227
				// _ = "end of CoverTab[25585]"
//line /usr/local/go/src/bufio/scan.go:227
			}
//line /usr/local/go/src/bufio/scan.go:227
			// _ = "end of CoverTab[25577]"
//line /usr/local/go/src/bufio/scan.go:227
			_go_fuzz_dep_.CoverTab[25578]++
								loop++
								if loop > maxConsecutiveEmptyReads {
//line /usr/local/go/src/bufio/scan.go:229
				_go_fuzz_dep_.CoverTab[25586]++
									s.setErr(io.ErrNoProgress)
									break
//line /usr/local/go/src/bufio/scan.go:231
				// _ = "end of CoverTab[25586]"
			} else {
//line /usr/local/go/src/bufio/scan.go:232
				_go_fuzz_dep_.CoverTab[25587]++
//line /usr/local/go/src/bufio/scan.go:232
				// _ = "end of CoverTab[25587]"
//line /usr/local/go/src/bufio/scan.go:232
			}
//line /usr/local/go/src/bufio/scan.go:232
			// _ = "end of CoverTab[25578]"
		}
//line /usr/local/go/src/bufio/scan.go:233
		// _ = "end of CoverTab[25536]"
	}
//line /usr/local/go/src/bufio/scan.go:234
	// _ = "end of CoverTab[25529]"
}

// advance consumes n bytes of the buffer. It reports whether the advance was legal.
func (s *Scanner) advance(n int) bool {
//line /usr/local/go/src/bufio/scan.go:238
	_go_fuzz_dep_.CoverTab[25588]++
						if n < 0 {
//line /usr/local/go/src/bufio/scan.go:239
		_go_fuzz_dep_.CoverTab[25591]++
							s.setErr(ErrNegativeAdvance)
							return false
//line /usr/local/go/src/bufio/scan.go:241
		// _ = "end of CoverTab[25591]"
	} else {
//line /usr/local/go/src/bufio/scan.go:242
		_go_fuzz_dep_.CoverTab[25592]++
//line /usr/local/go/src/bufio/scan.go:242
		// _ = "end of CoverTab[25592]"
//line /usr/local/go/src/bufio/scan.go:242
	}
//line /usr/local/go/src/bufio/scan.go:242
	// _ = "end of CoverTab[25588]"
//line /usr/local/go/src/bufio/scan.go:242
	_go_fuzz_dep_.CoverTab[25589]++
						if n > s.end-s.start {
//line /usr/local/go/src/bufio/scan.go:243
		_go_fuzz_dep_.CoverTab[25593]++
							s.setErr(ErrAdvanceTooFar)
							return false
//line /usr/local/go/src/bufio/scan.go:245
		// _ = "end of CoverTab[25593]"
	} else {
//line /usr/local/go/src/bufio/scan.go:246
		_go_fuzz_dep_.CoverTab[25594]++
//line /usr/local/go/src/bufio/scan.go:246
		// _ = "end of CoverTab[25594]"
//line /usr/local/go/src/bufio/scan.go:246
	}
//line /usr/local/go/src/bufio/scan.go:246
	// _ = "end of CoverTab[25589]"
//line /usr/local/go/src/bufio/scan.go:246
	_go_fuzz_dep_.CoverTab[25590]++
						s.start += n
						return true
//line /usr/local/go/src/bufio/scan.go:248
	// _ = "end of CoverTab[25590]"
}

// setErr records the first error encountered.
func (s *Scanner) setErr(err error) {
//line /usr/local/go/src/bufio/scan.go:252
	_go_fuzz_dep_.CoverTab[25595]++
						if s.err == nil || func() bool {
//line /usr/local/go/src/bufio/scan.go:253
		_go_fuzz_dep_.CoverTab[25596]++
//line /usr/local/go/src/bufio/scan.go:253
		return s.err == io.EOF
//line /usr/local/go/src/bufio/scan.go:253
		// _ = "end of CoverTab[25596]"
//line /usr/local/go/src/bufio/scan.go:253
	}() {
//line /usr/local/go/src/bufio/scan.go:253
		_go_fuzz_dep_.CoverTab[25597]++
							s.err = err
//line /usr/local/go/src/bufio/scan.go:254
		// _ = "end of CoverTab[25597]"
	} else {
//line /usr/local/go/src/bufio/scan.go:255
		_go_fuzz_dep_.CoverTab[25598]++
//line /usr/local/go/src/bufio/scan.go:255
		// _ = "end of CoverTab[25598]"
//line /usr/local/go/src/bufio/scan.go:255
	}
//line /usr/local/go/src/bufio/scan.go:255
	// _ = "end of CoverTab[25595]"
}

// Buffer sets the initial buffer to use when scanning and the maximum
//line /usr/local/go/src/bufio/scan.go:258
// size of buffer that may be allocated during scanning. The maximum
//line /usr/local/go/src/bufio/scan.go:258
// token size is the larger of max and cap(buf). If max <= cap(buf),
//line /usr/local/go/src/bufio/scan.go:258
// Scan will use this buffer only and do no allocation.
//line /usr/local/go/src/bufio/scan.go:258
//
//line /usr/local/go/src/bufio/scan.go:258
// By default, Scan uses an internal buffer and sets the
//line /usr/local/go/src/bufio/scan.go:258
// maximum token size to MaxScanTokenSize.
//line /usr/local/go/src/bufio/scan.go:258
//
//line /usr/local/go/src/bufio/scan.go:258
// Buffer panics if it is called after scanning has started.
//line /usr/local/go/src/bufio/scan.go:267
func (s *Scanner) Buffer(buf []byte, max int) {
//line /usr/local/go/src/bufio/scan.go:267
	_go_fuzz_dep_.CoverTab[25599]++
						if s.scanCalled {
//line /usr/local/go/src/bufio/scan.go:268
		_go_fuzz_dep_.CoverTab[25601]++
							panic("Buffer called after Scan")
//line /usr/local/go/src/bufio/scan.go:269
		// _ = "end of CoverTab[25601]"
	} else {
//line /usr/local/go/src/bufio/scan.go:270
		_go_fuzz_dep_.CoverTab[25602]++
//line /usr/local/go/src/bufio/scan.go:270
		// _ = "end of CoverTab[25602]"
//line /usr/local/go/src/bufio/scan.go:270
	}
//line /usr/local/go/src/bufio/scan.go:270
	// _ = "end of CoverTab[25599]"
//line /usr/local/go/src/bufio/scan.go:270
	_go_fuzz_dep_.CoverTab[25600]++
						s.buf = buf[0:cap(buf)]
						s.maxTokenSize = max
//line /usr/local/go/src/bufio/scan.go:272
	// _ = "end of CoverTab[25600]"
}

// Split sets the split function for the Scanner.
//line /usr/local/go/src/bufio/scan.go:275
// The default split function is ScanLines.
//line /usr/local/go/src/bufio/scan.go:275
//
//line /usr/local/go/src/bufio/scan.go:275
// Split panics if it is called after scanning has started.
//line /usr/local/go/src/bufio/scan.go:279
func (s *Scanner) Split(split SplitFunc) {
//line /usr/local/go/src/bufio/scan.go:279
	_go_fuzz_dep_.CoverTab[25603]++
						if s.scanCalled {
//line /usr/local/go/src/bufio/scan.go:280
		_go_fuzz_dep_.CoverTab[25605]++
							panic("Split called after Scan")
//line /usr/local/go/src/bufio/scan.go:281
		// _ = "end of CoverTab[25605]"
	} else {
//line /usr/local/go/src/bufio/scan.go:282
		_go_fuzz_dep_.CoverTab[25606]++
//line /usr/local/go/src/bufio/scan.go:282
		// _ = "end of CoverTab[25606]"
//line /usr/local/go/src/bufio/scan.go:282
	}
//line /usr/local/go/src/bufio/scan.go:282
	// _ = "end of CoverTab[25603]"
//line /usr/local/go/src/bufio/scan.go:282
	_go_fuzz_dep_.CoverTab[25604]++
						s.split = split
//line /usr/local/go/src/bufio/scan.go:283
	// _ = "end of CoverTab[25604]"
}

//line /usr/local/go/src/bufio/scan.go:288
// ScanBytes is a split function for a Scanner that returns each byte as a token.
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /usr/local/go/src/bufio/scan.go:289
	_go_fuzz_dep_.CoverTab[25607]++
						if atEOF && func() bool {
//line /usr/local/go/src/bufio/scan.go:290
		_go_fuzz_dep_.CoverTab[25609]++
//line /usr/local/go/src/bufio/scan.go:290
		return len(data) == 0
//line /usr/local/go/src/bufio/scan.go:290
		// _ = "end of CoverTab[25609]"
//line /usr/local/go/src/bufio/scan.go:290
	}() {
//line /usr/local/go/src/bufio/scan.go:290
		_go_fuzz_dep_.CoverTab[25610]++
							return 0, nil, nil
//line /usr/local/go/src/bufio/scan.go:291
		// _ = "end of CoverTab[25610]"
	} else {
//line /usr/local/go/src/bufio/scan.go:292
		_go_fuzz_dep_.CoverTab[25611]++
//line /usr/local/go/src/bufio/scan.go:292
		// _ = "end of CoverTab[25611]"
//line /usr/local/go/src/bufio/scan.go:292
	}
//line /usr/local/go/src/bufio/scan.go:292
	// _ = "end of CoverTab[25607]"
//line /usr/local/go/src/bufio/scan.go:292
	_go_fuzz_dep_.CoverTab[25608]++
						return 1, data[0:1], nil
//line /usr/local/go/src/bufio/scan.go:293
	// _ = "end of CoverTab[25608]"
}

var errorRune = []byte(string(utf8.RuneError))

// ScanRunes is a split function for a Scanner that returns each
//line /usr/local/go/src/bufio/scan.go:298
// UTF-8-encoded rune as a token. The sequence of runes returned is
//line /usr/local/go/src/bufio/scan.go:298
// equivalent to that from a range loop over the input as a string, which
//line /usr/local/go/src/bufio/scan.go:298
// means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
//line /usr/local/go/src/bufio/scan.go:298
// Because of the Scan interface, this makes it impossible for the client to
//line /usr/local/go/src/bufio/scan.go:298
// distinguish correctly encoded replacement runes from encoding errors.
//line /usr/local/go/src/bufio/scan.go:304
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /usr/local/go/src/bufio/scan.go:304
	_go_fuzz_dep_.CoverTab[25612]++
						if atEOF && func() bool {
//line /usr/local/go/src/bufio/scan.go:305
		_go_fuzz_dep_.CoverTab[25617]++
//line /usr/local/go/src/bufio/scan.go:305
		return len(data) == 0
//line /usr/local/go/src/bufio/scan.go:305
		// _ = "end of CoverTab[25617]"
//line /usr/local/go/src/bufio/scan.go:305
	}() {
//line /usr/local/go/src/bufio/scan.go:305
		_go_fuzz_dep_.CoverTab[25618]++
							return 0, nil, nil
//line /usr/local/go/src/bufio/scan.go:306
		// _ = "end of CoverTab[25618]"
	} else {
//line /usr/local/go/src/bufio/scan.go:307
		_go_fuzz_dep_.CoverTab[25619]++
//line /usr/local/go/src/bufio/scan.go:307
		// _ = "end of CoverTab[25619]"
//line /usr/local/go/src/bufio/scan.go:307
	}
//line /usr/local/go/src/bufio/scan.go:307
	// _ = "end of CoverTab[25612]"
//line /usr/local/go/src/bufio/scan.go:307
	_go_fuzz_dep_.CoverTab[25613]++

//line /usr/local/go/src/bufio/scan.go:310
	if data[0] < utf8.RuneSelf {
//line /usr/local/go/src/bufio/scan.go:310
		_go_fuzz_dep_.CoverTab[25620]++
							return 1, data[0:1], nil
//line /usr/local/go/src/bufio/scan.go:311
		// _ = "end of CoverTab[25620]"
	} else {
//line /usr/local/go/src/bufio/scan.go:312
		_go_fuzz_dep_.CoverTab[25621]++
//line /usr/local/go/src/bufio/scan.go:312
		// _ = "end of CoverTab[25621]"
//line /usr/local/go/src/bufio/scan.go:312
	}
//line /usr/local/go/src/bufio/scan.go:312
	// _ = "end of CoverTab[25613]"
//line /usr/local/go/src/bufio/scan.go:312
	_go_fuzz_dep_.CoverTab[25614]++

//line /usr/local/go/src/bufio/scan.go:315
	_, width := utf8.DecodeRune(data)
	if width > 1 {
//line /usr/local/go/src/bufio/scan.go:316
		_go_fuzz_dep_.CoverTab[25622]++

//line /usr/local/go/src/bufio/scan.go:319
		return width, data[0:width], nil
//line /usr/local/go/src/bufio/scan.go:319
		// _ = "end of CoverTab[25622]"
	} else {
//line /usr/local/go/src/bufio/scan.go:320
		_go_fuzz_dep_.CoverTab[25623]++
//line /usr/local/go/src/bufio/scan.go:320
		// _ = "end of CoverTab[25623]"
//line /usr/local/go/src/bufio/scan.go:320
	}
//line /usr/local/go/src/bufio/scan.go:320
	// _ = "end of CoverTab[25614]"
//line /usr/local/go/src/bufio/scan.go:320
	_go_fuzz_dep_.CoverTab[25615]++

//line /usr/local/go/src/bufio/scan.go:325
	if !atEOF && func() bool {
//line /usr/local/go/src/bufio/scan.go:325
		_go_fuzz_dep_.CoverTab[25624]++
//line /usr/local/go/src/bufio/scan.go:325
		return !utf8.FullRune(data)
//line /usr/local/go/src/bufio/scan.go:325
		// _ = "end of CoverTab[25624]"
//line /usr/local/go/src/bufio/scan.go:325
	}() {
//line /usr/local/go/src/bufio/scan.go:325
		_go_fuzz_dep_.CoverTab[25625]++

							return 0, nil, nil
//line /usr/local/go/src/bufio/scan.go:327
		// _ = "end of CoverTab[25625]"
	} else {
//line /usr/local/go/src/bufio/scan.go:328
		_go_fuzz_dep_.CoverTab[25626]++
//line /usr/local/go/src/bufio/scan.go:328
		// _ = "end of CoverTab[25626]"
//line /usr/local/go/src/bufio/scan.go:328
	}
//line /usr/local/go/src/bufio/scan.go:328
	// _ = "end of CoverTab[25615]"
//line /usr/local/go/src/bufio/scan.go:328
	_go_fuzz_dep_.CoverTab[25616]++

//line /usr/local/go/src/bufio/scan.go:333
	return 1, errorRune, nil
//line /usr/local/go/src/bufio/scan.go:333
	// _ = "end of CoverTab[25616]"
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
//line /usr/local/go/src/bufio/scan.go:337
	_go_fuzz_dep_.CoverTab[25627]++
						if len(data) > 0 && func() bool {
//line /usr/local/go/src/bufio/scan.go:338
		_go_fuzz_dep_.CoverTab[25629]++
//line /usr/local/go/src/bufio/scan.go:338
		return data[len(data)-1] == '\r'
//line /usr/local/go/src/bufio/scan.go:338
		// _ = "end of CoverTab[25629]"
//line /usr/local/go/src/bufio/scan.go:338
	}() {
//line /usr/local/go/src/bufio/scan.go:338
		_go_fuzz_dep_.CoverTab[25630]++
							return data[0 : len(data)-1]
//line /usr/local/go/src/bufio/scan.go:339
		// _ = "end of CoverTab[25630]"
	} else {
//line /usr/local/go/src/bufio/scan.go:340
		_go_fuzz_dep_.CoverTab[25631]++
//line /usr/local/go/src/bufio/scan.go:340
		// _ = "end of CoverTab[25631]"
//line /usr/local/go/src/bufio/scan.go:340
	}
//line /usr/local/go/src/bufio/scan.go:340
	// _ = "end of CoverTab[25627]"
//line /usr/local/go/src/bufio/scan.go:340
	_go_fuzz_dep_.CoverTab[25628]++
						return data
//line /usr/local/go/src/bufio/scan.go:341
	// _ = "end of CoverTab[25628]"
}

// ScanLines is a split function for a Scanner that returns each line of
//line /usr/local/go/src/bufio/scan.go:344
// text, stripped of any trailing end-of-line marker. The returned line may
//line /usr/local/go/src/bufio/scan.go:344
// be empty. The end-of-line marker is one optional carriage return followed
//line /usr/local/go/src/bufio/scan.go:344
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
//line /usr/local/go/src/bufio/scan.go:344
// The last non-empty line of input will be returned even if it has no
//line /usr/local/go/src/bufio/scan.go:344
// newline.
//line /usr/local/go/src/bufio/scan.go:350
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /usr/local/go/src/bufio/scan.go:350
	_go_fuzz_dep_.CoverTab[25632]++
						if atEOF && func() bool {
//line /usr/local/go/src/bufio/scan.go:351
		_go_fuzz_dep_.CoverTab[25636]++
//line /usr/local/go/src/bufio/scan.go:351
		return len(data) == 0
//line /usr/local/go/src/bufio/scan.go:351
		// _ = "end of CoverTab[25636]"
//line /usr/local/go/src/bufio/scan.go:351
	}() {
//line /usr/local/go/src/bufio/scan.go:351
		_go_fuzz_dep_.CoverTab[25637]++
							return 0, nil, nil
//line /usr/local/go/src/bufio/scan.go:352
		// _ = "end of CoverTab[25637]"
	} else {
//line /usr/local/go/src/bufio/scan.go:353
		_go_fuzz_dep_.CoverTab[25638]++
//line /usr/local/go/src/bufio/scan.go:353
		// _ = "end of CoverTab[25638]"
//line /usr/local/go/src/bufio/scan.go:353
	}
//line /usr/local/go/src/bufio/scan.go:353
	// _ = "end of CoverTab[25632]"
//line /usr/local/go/src/bufio/scan.go:353
	_go_fuzz_dep_.CoverTab[25633]++
						if i := bytes.IndexByte(data, '\n'); i >= 0 {
//line /usr/local/go/src/bufio/scan.go:354
		_go_fuzz_dep_.CoverTab[25639]++

							return i + 1, dropCR(data[0:i]), nil
//line /usr/local/go/src/bufio/scan.go:356
		// _ = "end of CoverTab[25639]"
	} else {
//line /usr/local/go/src/bufio/scan.go:357
		_go_fuzz_dep_.CoverTab[25640]++
//line /usr/local/go/src/bufio/scan.go:357
		// _ = "end of CoverTab[25640]"
//line /usr/local/go/src/bufio/scan.go:357
	}
//line /usr/local/go/src/bufio/scan.go:357
	// _ = "end of CoverTab[25633]"
//line /usr/local/go/src/bufio/scan.go:357
	_go_fuzz_dep_.CoverTab[25634]++

						if atEOF {
//line /usr/local/go/src/bufio/scan.go:359
		_go_fuzz_dep_.CoverTab[25641]++
							return len(data), dropCR(data), nil
//line /usr/local/go/src/bufio/scan.go:360
		// _ = "end of CoverTab[25641]"
	} else {
//line /usr/local/go/src/bufio/scan.go:361
		_go_fuzz_dep_.CoverTab[25642]++
//line /usr/local/go/src/bufio/scan.go:361
		// _ = "end of CoverTab[25642]"
//line /usr/local/go/src/bufio/scan.go:361
	}
//line /usr/local/go/src/bufio/scan.go:361
	// _ = "end of CoverTab[25634]"
//line /usr/local/go/src/bufio/scan.go:361
	_go_fuzz_dep_.CoverTab[25635]++

						return 0, nil, nil
//line /usr/local/go/src/bufio/scan.go:363
	// _ = "end of CoverTab[25635]"
}

// isSpace reports whether the character is a Unicode white space character.
//line /usr/local/go/src/bufio/scan.go:366
// We avoid dependency on the unicode package, but check validity of the implementation
//line /usr/local/go/src/bufio/scan.go:366
// in the tests.
//line /usr/local/go/src/bufio/scan.go:369
func isSpace(r rune) bool {
//line /usr/local/go/src/bufio/scan.go:369
	_go_fuzz_dep_.CoverTab[25643]++
						if r <= '\u00FF' {
//line /usr/local/go/src/bufio/scan.go:370
		_go_fuzz_dep_.CoverTab[25647]++

							switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
//line /usr/local/go/src/bufio/scan.go:373
			_go_fuzz_dep_.CoverTab[25649]++
								return true
//line /usr/local/go/src/bufio/scan.go:374
			// _ = "end of CoverTab[25649]"
		case '\u0085', '\u00A0':
//line /usr/local/go/src/bufio/scan.go:375
			_go_fuzz_dep_.CoverTab[25650]++
								return true
//line /usr/local/go/src/bufio/scan.go:376
			// _ = "end of CoverTab[25650]"
//line /usr/local/go/src/bufio/scan.go:376
		default:
//line /usr/local/go/src/bufio/scan.go:376
			_go_fuzz_dep_.CoverTab[25651]++
//line /usr/local/go/src/bufio/scan.go:376
			// _ = "end of CoverTab[25651]"
		}
//line /usr/local/go/src/bufio/scan.go:377
		// _ = "end of CoverTab[25647]"
//line /usr/local/go/src/bufio/scan.go:377
		_go_fuzz_dep_.CoverTab[25648]++
							return false
//line /usr/local/go/src/bufio/scan.go:378
		// _ = "end of CoverTab[25648]"
	} else {
//line /usr/local/go/src/bufio/scan.go:379
		_go_fuzz_dep_.CoverTab[25652]++
//line /usr/local/go/src/bufio/scan.go:379
		// _ = "end of CoverTab[25652]"
//line /usr/local/go/src/bufio/scan.go:379
	}
//line /usr/local/go/src/bufio/scan.go:379
	// _ = "end of CoverTab[25643]"
//line /usr/local/go/src/bufio/scan.go:379
	_go_fuzz_dep_.CoverTab[25644]++

						if '\u2000' <= r && func() bool {
//line /usr/local/go/src/bufio/scan.go:381
		_go_fuzz_dep_.CoverTab[25653]++
//line /usr/local/go/src/bufio/scan.go:381
		return r <= '\u200a'
//line /usr/local/go/src/bufio/scan.go:381
		// _ = "end of CoverTab[25653]"
//line /usr/local/go/src/bufio/scan.go:381
	}() {
//line /usr/local/go/src/bufio/scan.go:381
		_go_fuzz_dep_.CoverTab[25654]++
							return true
//line /usr/local/go/src/bufio/scan.go:382
		// _ = "end of CoverTab[25654]"
	} else {
//line /usr/local/go/src/bufio/scan.go:383
		_go_fuzz_dep_.CoverTab[25655]++
//line /usr/local/go/src/bufio/scan.go:383
		// _ = "end of CoverTab[25655]"
//line /usr/local/go/src/bufio/scan.go:383
	}
//line /usr/local/go/src/bufio/scan.go:383
	// _ = "end of CoverTab[25644]"
//line /usr/local/go/src/bufio/scan.go:383
	_go_fuzz_dep_.CoverTab[25645]++
						switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
//line /usr/local/go/src/bufio/scan.go:385
		_go_fuzz_dep_.CoverTab[25656]++
							return true
//line /usr/local/go/src/bufio/scan.go:386
		// _ = "end of CoverTab[25656]"
//line /usr/local/go/src/bufio/scan.go:386
	default:
//line /usr/local/go/src/bufio/scan.go:386
		_go_fuzz_dep_.CoverTab[25657]++
//line /usr/local/go/src/bufio/scan.go:386
		// _ = "end of CoverTab[25657]"
	}
//line /usr/local/go/src/bufio/scan.go:387
	// _ = "end of CoverTab[25645]"
//line /usr/local/go/src/bufio/scan.go:387
	_go_fuzz_dep_.CoverTab[25646]++
						return false
//line /usr/local/go/src/bufio/scan.go:388
	// _ = "end of CoverTab[25646]"
}

// ScanWords is a split function for a Scanner that returns each
//line /usr/local/go/src/bufio/scan.go:391
// space-separated word of text, with surrounding spaces deleted. It will
//line /usr/local/go/src/bufio/scan.go:391
// never return an empty string. The definition of space is set by
//line /usr/local/go/src/bufio/scan.go:391
// unicode.IsSpace.
//line /usr/local/go/src/bufio/scan.go:395
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
//line /usr/local/go/src/bufio/scan.go:395
	_go_fuzz_dep_.CoverTab[25658]++

						start := 0
						for width := 0; start < len(data); start += width {
//line /usr/local/go/src/bufio/scan.go:398
		_go_fuzz_dep_.CoverTab[25662]++
							var r rune
							r, width = utf8.DecodeRune(data[start:])
							if !isSpace(r) {
//line /usr/local/go/src/bufio/scan.go:401
			_go_fuzz_dep_.CoverTab[25663]++
								break
//line /usr/local/go/src/bufio/scan.go:402
			// _ = "end of CoverTab[25663]"
		} else {
//line /usr/local/go/src/bufio/scan.go:403
			_go_fuzz_dep_.CoverTab[25664]++
//line /usr/local/go/src/bufio/scan.go:403
			// _ = "end of CoverTab[25664]"
//line /usr/local/go/src/bufio/scan.go:403
		}
//line /usr/local/go/src/bufio/scan.go:403
		// _ = "end of CoverTab[25662]"
	}
//line /usr/local/go/src/bufio/scan.go:404
	// _ = "end of CoverTab[25658]"
//line /usr/local/go/src/bufio/scan.go:404
	_go_fuzz_dep_.CoverTab[25659]++

						for width, i := 0, start; i < len(data); i += width {
//line /usr/local/go/src/bufio/scan.go:406
		_go_fuzz_dep_.CoverTab[25665]++
							var r rune
							r, width = utf8.DecodeRune(data[i:])
							if isSpace(r) {
//line /usr/local/go/src/bufio/scan.go:409
			_go_fuzz_dep_.CoverTab[25666]++
								return i + width, data[start:i], nil
//line /usr/local/go/src/bufio/scan.go:410
			// _ = "end of CoverTab[25666]"
		} else {
//line /usr/local/go/src/bufio/scan.go:411
			_go_fuzz_dep_.CoverTab[25667]++
//line /usr/local/go/src/bufio/scan.go:411
			// _ = "end of CoverTab[25667]"
//line /usr/local/go/src/bufio/scan.go:411
		}
//line /usr/local/go/src/bufio/scan.go:411
		// _ = "end of CoverTab[25665]"
	}
//line /usr/local/go/src/bufio/scan.go:412
	// _ = "end of CoverTab[25659]"
//line /usr/local/go/src/bufio/scan.go:412
	_go_fuzz_dep_.CoverTab[25660]++

						if atEOF && func() bool {
//line /usr/local/go/src/bufio/scan.go:414
		_go_fuzz_dep_.CoverTab[25668]++
//line /usr/local/go/src/bufio/scan.go:414
		return len(data) > start
//line /usr/local/go/src/bufio/scan.go:414
		// _ = "end of CoverTab[25668]"
//line /usr/local/go/src/bufio/scan.go:414
	}() {
//line /usr/local/go/src/bufio/scan.go:414
		_go_fuzz_dep_.CoverTab[25669]++
							return len(data), data[start:], nil
//line /usr/local/go/src/bufio/scan.go:415
		// _ = "end of CoverTab[25669]"
	} else {
//line /usr/local/go/src/bufio/scan.go:416
		_go_fuzz_dep_.CoverTab[25670]++
//line /usr/local/go/src/bufio/scan.go:416
		// _ = "end of CoverTab[25670]"
//line /usr/local/go/src/bufio/scan.go:416
	}
//line /usr/local/go/src/bufio/scan.go:416
	// _ = "end of CoverTab[25660]"
//line /usr/local/go/src/bufio/scan.go:416
	_go_fuzz_dep_.CoverTab[25661]++

						return start, nil, nil
//line /usr/local/go/src/bufio/scan.go:418
	// _ = "end of CoverTab[25661]"
}

//line /usr/local/go/src/bufio/scan.go:419
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/bufio/scan.go:419
var _ = _go_fuzz_dep_.CoverTab
