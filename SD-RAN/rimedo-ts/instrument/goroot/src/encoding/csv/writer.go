// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/csv/writer.go:5
package csv

//line /usr/local/go/src/encoding/csv/writer.go:5
import (
//line /usr/local/go/src/encoding/csv/writer.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/csv/writer.go:5
)
//line /usr/local/go/src/encoding/csv/writer.go:5
import (
//line /usr/local/go/src/encoding/csv/writer.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/csv/writer.go:5
)

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

// A Writer writes records using CSV encoding.
//line /usr/local/go/src/encoding/csv/writer.go:15
//
//line /usr/local/go/src/encoding/csv/writer.go:15
// As returned by NewWriter, a Writer writes records terminated by a
//line /usr/local/go/src/encoding/csv/writer.go:15
// newline and uses ',' as the field delimiter. The exported fields can be
//line /usr/local/go/src/encoding/csv/writer.go:15
// changed to customize the details before the first call to Write or WriteAll.
//line /usr/local/go/src/encoding/csv/writer.go:15
//
//line /usr/local/go/src/encoding/csv/writer.go:15
// Comma is the field delimiter.
//line /usr/local/go/src/encoding/csv/writer.go:15
//
//line /usr/local/go/src/encoding/csv/writer.go:15
// If UseCRLF is true, the Writer ends each output line with \r\n instead of \n.
//line /usr/local/go/src/encoding/csv/writer.go:15
//
//line /usr/local/go/src/encoding/csv/writer.go:15
// The writes of individual records are buffered.
//line /usr/local/go/src/encoding/csv/writer.go:15
// After all data has been written, the client should call the
//line /usr/local/go/src/encoding/csv/writer.go:15
// Flush method to guarantee all data has been forwarded to
//line /usr/local/go/src/encoding/csv/writer.go:15
// the underlying io.Writer.  Any errors that occurred should
//line /usr/local/go/src/encoding/csv/writer.go:15
// be checked by calling the Error method.
//line /usr/local/go/src/encoding/csv/writer.go:30
type Writer struct {
	Comma	rune	// Field delimiter (set to ',' by NewWriter)
	UseCRLF	bool	// True to use \r\n as the line terminator
	w	*bufio.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
//line /usr/local/go/src/encoding/csv/writer.go:37
	_go_fuzz_dep_.CoverTab[114912]++
							return &Writer{
		Comma:	',',
		w:	bufio.NewWriter(w),
	}
//line /usr/local/go/src/encoding/csv/writer.go:41
	// _ = "end of CoverTab[114912]"
}

// Write writes a single CSV record to w along with any necessary quoting.
//line /usr/local/go/src/encoding/csv/writer.go:44
// A record is a slice of strings with each string being one field.
//line /usr/local/go/src/encoding/csv/writer.go:44
// Writes are buffered, so Flush must eventually be called to ensure
//line /usr/local/go/src/encoding/csv/writer.go:44
// that the record is written to the underlying io.Writer.
//line /usr/local/go/src/encoding/csv/writer.go:48
func (w *Writer) Write(record []string) error {
//line /usr/local/go/src/encoding/csv/writer.go:48
	_go_fuzz_dep_.CoverTab[114913]++
							if !validDelim(w.Comma) {
//line /usr/local/go/src/encoding/csv/writer.go:49
		_go_fuzz_dep_.CoverTab[114917]++
								return errInvalidDelim
//line /usr/local/go/src/encoding/csv/writer.go:50
		// _ = "end of CoverTab[114917]"
	} else {
//line /usr/local/go/src/encoding/csv/writer.go:51
		_go_fuzz_dep_.CoverTab[114918]++
//line /usr/local/go/src/encoding/csv/writer.go:51
		// _ = "end of CoverTab[114918]"
//line /usr/local/go/src/encoding/csv/writer.go:51
	}
//line /usr/local/go/src/encoding/csv/writer.go:51
	// _ = "end of CoverTab[114913]"
//line /usr/local/go/src/encoding/csv/writer.go:51
	_go_fuzz_dep_.CoverTab[114914]++

							for n, field := range record {
//line /usr/local/go/src/encoding/csv/writer.go:53
		_go_fuzz_dep_.CoverTab[114919]++
								if n > 0 {
//line /usr/local/go/src/encoding/csv/writer.go:54
			_go_fuzz_dep_.CoverTab[114924]++
									if _, err := w.w.WriteRune(w.Comma); err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:55
				_go_fuzz_dep_.CoverTab[114925]++
										return err
//line /usr/local/go/src/encoding/csv/writer.go:56
				// _ = "end of CoverTab[114925]"
			} else {
//line /usr/local/go/src/encoding/csv/writer.go:57
				_go_fuzz_dep_.CoverTab[114926]++
//line /usr/local/go/src/encoding/csv/writer.go:57
				// _ = "end of CoverTab[114926]"
//line /usr/local/go/src/encoding/csv/writer.go:57
			}
//line /usr/local/go/src/encoding/csv/writer.go:57
			// _ = "end of CoverTab[114924]"
		} else {
//line /usr/local/go/src/encoding/csv/writer.go:58
			_go_fuzz_dep_.CoverTab[114927]++
//line /usr/local/go/src/encoding/csv/writer.go:58
			// _ = "end of CoverTab[114927]"
//line /usr/local/go/src/encoding/csv/writer.go:58
		}
//line /usr/local/go/src/encoding/csv/writer.go:58
		// _ = "end of CoverTab[114919]"
//line /usr/local/go/src/encoding/csv/writer.go:58
		_go_fuzz_dep_.CoverTab[114920]++

//line /usr/local/go/src/encoding/csv/writer.go:62
		if !w.fieldNeedsQuotes(field) {
//line /usr/local/go/src/encoding/csv/writer.go:62
			_go_fuzz_dep_.CoverTab[114928]++
									if _, err := w.w.WriteString(field); err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:63
				_go_fuzz_dep_.CoverTab[114930]++
										return err
//line /usr/local/go/src/encoding/csv/writer.go:64
				// _ = "end of CoverTab[114930]"
			} else {
//line /usr/local/go/src/encoding/csv/writer.go:65
				_go_fuzz_dep_.CoverTab[114931]++
//line /usr/local/go/src/encoding/csv/writer.go:65
				// _ = "end of CoverTab[114931]"
//line /usr/local/go/src/encoding/csv/writer.go:65
			}
//line /usr/local/go/src/encoding/csv/writer.go:65
			// _ = "end of CoverTab[114928]"
//line /usr/local/go/src/encoding/csv/writer.go:65
			_go_fuzz_dep_.CoverTab[114929]++
									continue
//line /usr/local/go/src/encoding/csv/writer.go:66
			// _ = "end of CoverTab[114929]"
		} else {
//line /usr/local/go/src/encoding/csv/writer.go:67
			_go_fuzz_dep_.CoverTab[114932]++
//line /usr/local/go/src/encoding/csv/writer.go:67
			// _ = "end of CoverTab[114932]"
//line /usr/local/go/src/encoding/csv/writer.go:67
		}
//line /usr/local/go/src/encoding/csv/writer.go:67
		// _ = "end of CoverTab[114920]"
//line /usr/local/go/src/encoding/csv/writer.go:67
		_go_fuzz_dep_.CoverTab[114921]++

								if err := w.w.WriteByte('"'); err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:69
			_go_fuzz_dep_.CoverTab[114933]++
									return err
//line /usr/local/go/src/encoding/csv/writer.go:70
			// _ = "end of CoverTab[114933]"
		} else {
//line /usr/local/go/src/encoding/csv/writer.go:71
			_go_fuzz_dep_.CoverTab[114934]++
//line /usr/local/go/src/encoding/csv/writer.go:71
			// _ = "end of CoverTab[114934]"
//line /usr/local/go/src/encoding/csv/writer.go:71
		}
//line /usr/local/go/src/encoding/csv/writer.go:71
		// _ = "end of CoverTab[114921]"
//line /usr/local/go/src/encoding/csv/writer.go:71
		_go_fuzz_dep_.CoverTab[114922]++
								for len(field) > 0 {
//line /usr/local/go/src/encoding/csv/writer.go:72
			_go_fuzz_dep_.CoverTab[114935]++

									i := strings.IndexAny(field, "\"\r\n")
									if i < 0 {
//line /usr/local/go/src/encoding/csv/writer.go:75
				_go_fuzz_dep_.CoverTab[114938]++
										i = len(field)
//line /usr/local/go/src/encoding/csv/writer.go:76
				// _ = "end of CoverTab[114938]"
			} else {
//line /usr/local/go/src/encoding/csv/writer.go:77
				_go_fuzz_dep_.CoverTab[114939]++
//line /usr/local/go/src/encoding/csv/writer.go:77
				// _ = "end of CoverTab[114939]"
//line /usr/local/go/src/encoding/csv/writer.go:77
			}
//line /usr/local/go/src/encoding/csv/writer.go:77
			// _ = "end of CoverTab[114935]"
//line /usr/local/go/src/encoding/csv/writer.go:77
			_go_fuzz_dep_.CoverTab[114936]++

//line /usr/local/go/src/encoding/csv/writer.go:80
			if _, err := w.w.WriteString(field[:i]); err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:80
				_go_fuzz_dep_.CoverTab[114940]++
										return err
//line /usr/local/go/src/encoding/csv/writer.go:81
				// _ = "end of CoverTab[114940]"
			} else {
//line /usr/local/go/src/encoding/csv/writer.go:82
				_go_fuzz_dep_.CoverTab[114941]++
//line /usr/local/go/src/encoding/csv/writer.go:82
				// _ = "end of CoverTab[114941]"
//line /usr/local/go/src/encoding/csv/writer.go:82
			}
//line /usr/local/go/src/encoding/csv/writer.go:82
			// _ = "end of CoverTab[114936]"
//line /usr/local/go/src/encoding/csv/writer.go:82
			_go_fuzz_dep_.CoverTab[114937]++
									field = field[i:]

//line /usr/local/go/src/encoding/csv/writer.go:86
			if len(field) > 0 {
//line /usr/local/go/src/encoding/csv/writer.go:86
				_go_fuzz_dep_.CoverTab[114942]++
										var err error
										switch field[0] {
				case '"':
//line /usr/local/go/src/encoding/csv/writer.go:89
					_go_fuzz_dep_.CoverTab[114944]++
											_, err = w.w.WriteString(`""`)
//line /usr/local/go/src/encoding/csv/writer.go:90
					// _ = "end of CoverTab[114944]"
				case '\r':
//line /usr/local/go/src/encoding/csv/writer.go:91
					_go_fuzz_dep_.CoverTab[114945]++
											if !w.UseCRLF {
//line /usr/local/go/src/encoding/csv/writer.go:92
						_go_fuzz_dep_.CoverTab[114948]++
												err = w.w.WriteByte('\r')
//line /usr/local/go/src/encoding/csv/writer.go:93
						// _ = "end of CoverTab[114948]"
					} else {
//line /usr/local/go/src/encoding/csv/writer.go:94
						_go_fuzz_dep_.CoverTab[114949]++
//line /usr/local/go/src/encoding/csv/writer.go:94
						// _ = "end of CoverTab[114949]"
//line /usr/local/go/src/encoding/csv/writer.go:94
					}
//line /usr/local/go/src/encoding/csv/writer.go:94
					// _ = "end of CoverTab[114945]"
				case '\n':
//line /usr/local/go/src/encoding/csv/writer.go:95
					_go_fuzz_dep_.CoverTab[114946]++
											if w.UseCRLF {
//line /usr/local/go/src/encoding/csv/writer.go:96
						_go_fuzz_dep_.CoverTab[114950]++
												_, err = w.w.WriteString("\r\n")
//line /usr/local/go/src/encoding/csv/writer.go:97
						// _ = "end of CoverTab[114950]"
					} else {
//line /usr/local/go/src/encoding/csv/writer.go:98
						_go_fuzz_dep_.CoverTab[114951]++
												err = w.w.WriteByte('\n')
//line /usr/local/go/src/encoding/csv/writer.go:99
						// _ = "end of CoverTab[114951]"
					}
//line /usr/local/go/src/encoding/csv/writer.go:100
					// _ = "end of CoverTab[114946]"
//line /usr/local/go/src/encoding/csv/writer.go:100
				default:
//line /usr/local/go/src/encoding/csv/writer.go:100
					_go_fuzz_dep_.CoverTab[114947]++
//line /usr/local/go/src/encoding/csv/writer.go:100
					// _ = "end of CoverTab[114947]"
				}
//line /usr/local/go/src/encoding/csv/writer.go:101
				// _ = "end of CoverTab[114942]"
//line /usr/local/go/src/encoding/csv/writer.go:101
				_go_fuzz_dep_.CoverTab[114943]++
										field = field[1:]
										if err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:103
					_go_fuzz_dep_.CoverTab[114952]++
											return err
//line /usr/local/go/src/encoding/csv/writer.go:104
					// _ = "end of CoverTab[114952]"
				} else {
//line /usr/local/go/src/encoding/csv/writer.go:105
					_go_fuzz_dep_.CoverTab[114953]++
//line /usr/local/go/src/encoding/csv/writer.go:105
					// _ = "end of CoverTab[114953]"
//line /usr/local/go/src/encoding/csv/writer.go:105
				}
//line /usr/local/go/src/encoding/csv/writer.go:105
				// _ = "end of CoverTab[114943]"
			} else {
//line /usr/local/go/src/encoding/csv/writer.go:106
				_go_fuzz_dep_.CoverTab[114954]++
//line /usr/local/go/src/encoding/csv/writer.go:106
				// _ = "end of CoverTab[114954]"
//line /usr/local/go/src/encoding/csv/writer.go:106
			}
//line /usr/local/go/src/encoding/csv/writer.go:106
			// _ = "end of CoverTab[114937]"
		}
//line /usr/local/go/src/encoding/csv/writer.go:107
		// _ = "end of CoverTab[114922]"
//line /usr/local/go/src/encoding/csv/writer.go:107
		_go_fuzz_dep_.CoverTab[114923]++
								if err := w.w.WriteByte('"'); err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:108
			_go_fuzz_dep_.CoverTab[114955]++
									return err
//line /usr/local/go/src/encoding/csv/writer.go:109
			// _ = "end of CoverTab[114955]"
		} else {
//line /usr/local/go/src/encoding/csv/writer.go:110
			_go_fuzz_dep_.CoverTab[114956]++
//line /usr/local/go/src/encoding/csv/writer.go:110
			// _ = "end of CoverTab[114956]"
//line /usr/local/go/src/encoding/csv/writer.go:110
		}
//line /usr/local/go/src/encoding/csv/writer.go:110
		// _ = "end of CoverTab[114923]"
	}
//line /usr/local/go/src/encoding/csv/writer.go:111
	// _ = "end of CoverTab[114914]"
//line /usr/local/go/src/encoding/csv/writer.go:111
	_go_fuzz_dep_.CoverTab[114915]++
							var err error
							if w.UseCRLF {
//line /usr/local/go/src/encoding/csv/writer.go:113
		_go_fuzz_dep_.CoverTab[114957]++
								_, err = w.w.WriteString("\r\n")
//line /usr/local/go/src/encoding/csv/writer.go:114
		// _ = "end of CoverTab[114957]"
	} else {
//line /usr/local/go/src/encoding/csv/writer.go:115
		_go_fuzz_dep_.CoverTab[114958]++
								err = w.w.WriteByte('\n')
//line /usr/local/go/src/encoding/csv/writer.go:116
		// _ = "end of CoverTab[114958]"
	}
//line /usr/local/go/src/encoding/csv/writer.go:117
	// _ = "end of CoverTab[114915]"
//line /usr/local/go/src/encoding/csv/writer.go:117
	_go_fuzz_dep_.CoverTab[114916]++
							return err
//line /usr/local/go/src/encoding/csv/writer.go:118
	// _ = "end of CoverTab[114916]"
}

// Flush writes any buffered data to the underlying io.Writer.
//line /usr/local/go/src/encoding/csv/writer.go:121
// To check if an error occurred during the Flush, call Error.
//line /usr/local/go/src/encoding/csv/writer.go:123
func (w *Writer) Flush() {
//line /usr/local/go/src/encoding/csv/writer.go:123
	_go_fuzz_dep_.CoverTab[114959]++
							w.w.Flush()
//line /usr/local/go/src/encoding/csv/writer.go:124
	// _ = "end of CoverTab[114959]"
}

// Error reports any error that has occurred during a previous Write or Flush.
func (w *Writer) Error() error {
//line /usr/local/go/src/encoding/csv/writer.go:128
	_go_fuzz_dep_.CoverTab[114960]++
							_, err := w.w.Write(nil)
							return err
//line /usr/local/go/src/encoding/csv/writer.go:130
	// _ = "end of CoverTab[114960]"
}

// WriteAll writes multiple CSV records to w using Write and then calls Flush,
//line /usr/local/go/src/encoding/csv/writer.go:133
// returning any error from the Flush.
//line /usr/local/go/src/encoding/csv/writer.go:135
func (w *Writer) WriteAll(records [][]string) error {
//line /usr/local/go/src/encoding/csv/writer.go:135
	_go_fuzz_dep_.CoverTab[114961]++
							for _, record := range records {
//line /usr/local/go/src/encoding/csv/writer.go:136
		_go_fuzz_dep_.CoverTab[114963]++
								err := w.Write(record)
								if err != nil {
//line /usr/local/go/src/encoding/csv/writer.go:138
			_go_fuzz_dep_.CoverTab[114964]++
									return err
//line /usr/local/go/src/encoding/csv/writer.go:139
			// _ = "end of CoverTab[114964]"
		} else {
//line /usr/local/go/src/encoding/csv/writer.go:140
			_go_fuzz_dep_.CoverTab[114965]++
//line /usr/local/go/src/encoding/csv/writer.go:140
			// _ = "end of CoverTab[114965]"
//line /usr/local/go/src/encoding/csv/writer.go:140
		}
//line /usr/local/go/src/encoding/csv/writer.go:140
		// _ = "end of CoverTab[114963]"
	}
//line /usr/local/go/src/encoding/csv/writer.go:141
	// _ = "end of CoverTab[114961]"
//line /usr/local/go/src/encoding/csv/writer.go:141
	_go_fuzz_dep_.CoverTab[114962]++
							return w.w.Flush()
//line /usr/local/go/src/encoding/csv/writer.go:142
	// _ = "end of CoverTab[114962]"
}

// fieldNeedsQuotes reports whether our field must be enclosed in quotes.
//line /usr/local/go/src/encoding/csv/writer.go:145
// Fields with a Comma, fields with a quote or newline, and
//line /usr/local/go/src/encoding/csv/writer.go:145
// fields which start with a space must be enclosed in quotes.
//line /usr/local/go/src/encoding/csv/writer.go:145
// We used to quote empty strings, but we do not anymore (as of Go 1.4).
//line /usr/local/go/src/encoding/csv/writer.go:145
// The two representations should be equivalent, but Postgres distinguishes
//line /usr/local/go/src/encoding/csv/writer.go:145
// quoted vs non-quoted empty string during database imports, and it has
//line /usr/local/go/src/encoding/csv/writer.go:145
// an option to force the quoted behavior for non-quoted CSV but it has
//line /usr/local/go/src/encoding/csv/writer.go:145
// no option to force the non-quoted behavior for quoted CSV, making
//line /usr/local/go/src/encoding/csv/writer.go:145
// CSV with quoted empty strings strictly less useful.
//line /usr/local/go/src/encoding/csv/writer.go:145
// Not quoting the empty string also makes this package match the behavior
//line /usr/local/go/src/encoding/csv/writer.go:145
// of Microsoft Excel and Google Drive.
//line /usr/local/go/src/encoding/csv/writer.go:145
// For Postgres, quote the data terminating string `\.`.
//line /usr/local/go/src/encoding/csv/writer.go:157
func (w *Writer) fieldNeedsQuotes(field string) bool {
//line /usr/local/go/src/encoding/csv/writer.go:157
	_go_fuzz_dep_.CoverTab[114966]++
							if field == "" {
//line /usr/local/go/src/encoding/csv/writer.go:158
		_go_fuzz_dep_.CoverTab[114970]++
								return false
//line /usr/local/go/src/encoding/csv/writer.go:159
		// _ = "end of CoverTab[114970]"
	} else {
//line /usr/local/go/src/encoding/csv/writer.go:160
		_go_fuzz_dep_.CoverTab[114971]++
//line /usr/local/go/src/encoding/csv/writer.go:160
		// _ = "end of CoverTab[114971]"
//line /usr/local/go/src/encoding/csv/writer.go:160
	}
//line /usr/local/go/src/encoding/csv/writer.go:160
	// _ = "end of CoverTab[114966]"
//line /usr/local/go/src/encoding/csv/writer.go:160
	_go_fuzz_dep_.CoverTab[114967]++

							if field == `\.` {
//line /usr/local/go/src/encoding/csv/writer.go:162
		_go_fuzz_dep_.CoverTab[114972]++
								return true
//line /usr/local/go/src/encoding/csv/writer.go:163
		// _ = "end of CoverTab[114972]"
	} else {
//line /usr/local/go/src/encoding/csv/writer.go:164
		_go_fuzz_dep_.CoverTab[114973]++
//line /usr/local/go/src/encoding/csv/writer.go:164
		// _ = "end of CoverTab[114973]"
//line /usr/local/go/src/encoding/csv/writer.go:164
	}
//line /usr/local/go/src/encoding/csv/writer.go:164
	// _ = "end of CoverTab[114967]"
//line /usr/local/go/src/encoding/csv/writer.go:164
	_go_fuzz_dep_.CoverTab[114968]++

							if w.Comma < utf8.RuneSelf {
//line /usr/local/go/src/encoding/csv/writer.go:166
		_go_fuzz_dep_.CoverTab[114974]++
								for i := 0; i < len(field); i++ {
//line /usr/local/go/src/encoding/csv/writer.go:167
			_go_fuzz_dep_.CoverTab[114975]++
									c := field[i]
									if c == '\n' || func() bool {
//line /usr/local/go/src/encoding/csv/writer.go:169
				_go_fuzz_dep_.CoverTab[114976]++
//line /usr/local/go/src/encoding/csv/writer.go:169
				return c == '\r'
//line /usr/local/go/src/encoding/csv/writer.go:169
				// _ = "end of CoverTab[114976]"
//line /usr/local/go/src/encoding/csv/writer.go:169
			}() || func() bool {
//line /usr/local/go/src/encoding/csv/writer.go:169
				_go_fuzz_dep_.CoverTab[114977]++
//line /usr/local/go/src/encoding/csv/writer.go:169
				return c == '"'
//line /usr/local/go/src/encoding/csv/writer.go:169
				// _ = "end of CoverTab[114977]"
//line /usr/local/go/src/encoding/csv/writer.go:169
			}() || func() bool {
//line /usr/local/go/src/encoding/csv/writer.go:169
				_go_fuzz_dep_.CoverTab[114978]++
//line /usr/local/go/src/encoding/csv/writer.go:169
				return c == byte(w.Comma)
//line /usr/local/go/src/encoding/csv/writer.go:169
				// _ = "end of CoverTab[114978]"
//line /usr/local/go/src/encoding/csv/writer.go:169
			}() {
//line /usr/local/go/src/encoding/csv/writer.go:169
				_go_fuzz_dep_.CoverTab[114979]++
										return true
//line /usr/local/go/src/encoding/csv/writer.go:170
				// _ = "end of CoverTab[114979]"
			} else {
//line /usr/local/go/src/encoding/csv/writer.go:171
				_go_fuzz_dep_.CoverTab[114980]++
//line /usr/local/go/src/encoding/csv/writer.go:171
				// _ = "end of CoverTab[114980]"
//line /usr/local/go/src/encoding/csv/writer.go:171
			}
//line /usr/local/go/src/encoding/csv/writer.go:171
			// _ = "end of CoverTab[114975]"
		}
//line /usr/local/go/src/encoding/csv/writer.go:172
		// _ = "end of CoverTab[114974]"
	} else {
//line /usr/local/go/src/encoding/csv/writer.go:173
		_go_fuzz_dep_.CoverTab[114981]++
								if strings.ContainsRune(field, w.Comma) || func() bool {
//line /usr/local/go/src/encoding/csv/writer.go:174
			_go_fuzz_dep_.CoverTab[114982]++
//line /usr/local/go/src/encoding/csv/writer.go:174
			return strings.ContainsAny(field, "\"\r\n")
//line /usr/local/go/src/encoding/csv/writer.go:174
			// _ = "end of CoverTab[114982]"
//line /usr/local/go/src/encoding/csv/writer.go:174
		}() {
//line /usr/local/go/src/encoding/csv/writer.go:174
			_go_fuzz_dep_.CoverTab[114983]++
									return true
//line /usr/local/go/src/encoding/csv/writer.go:175
			// _ = "end of CoverTab[114983]"
		} else {
//line /usr/local/go/src/encoding/csv/writer.go:176
			_go_fuzz_dep_.CoverTab[114984]++
//line /usr/local/go/src/encoding/csv/writer.go:176
			// _ = "end of CoverTab[114984]"
//line /usr/local/go/src/encoding/csv/writer.go:176
		}
//line /usr/local/go/src/encoding/csv/writer.go:176
		// _ = "end of CoverTab[114981]"
	}
//line /usr/local/go/src/encoding/csv/writer.go:177
	// _ = "end of CoverTab[114968]"
//line /usr/local/go/src/encoding/csv/writer.go:177
	_go_fuzz_dep_.CoverTab[114969]++

							r1, _ := utf8.DecodeRuneInString(field)
							return unicode.IsSpace(r1)
//line /usr/local/go/src/encoding/csv/writer.go:180
	// _ = "end of CoverTab[114969]"
}

//line /usr/local/go/src/encoding/csv/writer.go:181
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/csv/writer.go:181
var _ = _go_fuzz_dep_.CoverTab
