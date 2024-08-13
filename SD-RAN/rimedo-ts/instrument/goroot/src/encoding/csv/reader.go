// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/csv/reader.go:5
// Package csv reads and writes comma-separated values (CSV) files.
//line /usr/local/go/src/encoding/csv/reader.go:5
// There are many kinds of CSV files; this package supports the format
//line /usr/local/go/src/encoding/csv/reader.go:5
// described in RFC 4180.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// A csv file contains zero or more records of one or more fields per record.
//line /usr/local/go/src/encoding/csv/reader.go:5
// Each record is separated by the newline character. The final record may
//line /usr/local/go/src/encoding/csv/reader.go:5
// optionally be followed by a newline character.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	field1,field2,field3
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// White space is considered part of a field.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// Carriage returns before newline characters are silently removed.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// Blank lines are ignored. A line with only whitespace characters (excluding
//line /usr/local/go/src/encoding/csv/reader.go:5
// the ending newline character) is not considered a blank line.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// Fields which start and stop with the quote character " are called
//line /usr/local/go/src/encoding/csv/reader.go:5
// quoted-fields. The beginning and ending quote are not part of the
//line /usr/local/go/src/encoding/csv/reader.go:5
// field.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// The source:
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	normal string,"quoted-field"
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// results in the fields
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	{`normal string`, `quoted-field`}
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// Within a quoted-field a quote character followed by a second quote
//line /usr/local/go/src/encoding/csv/reader.go:5
// character is considered a single quote.
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	"the ""word"" is true","a ""quoted-field"""
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// results in
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	{`the "word" is true`, `a "quoted-field"`}
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// Newlines and commas may be included in a quoted-field
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	"Multi-line
//line /usr/local/go/src/encoding/csv/reader.go:5
//	field","comma is ,"
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
// results in
//line /usr/local/go/src/encoding/csv/reader.go:5
//
//line /usr/local/go/src/encoding/csv/reader.go:5
//	{`Multi-line
//line /usr/local/go/src/encoding/csv/reader.go:5
//	field`, `comma is ,`}
//line /usr/local/go/src/encoding/csv/reader.go:52
package csv

//line /usr/local/go/src/encoding/csv/reader.go:52
import (
//line /usr/local/go/src/encoding/csv/reader.go:52
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/csv/reader.go:52
)
//line /usr/local/go/src/encoding/csv/reader.go:52
import (
//line /usr/local/go/src/encoding/csv/reader.go:52
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/csv/reader.go:52
)

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// A ParseError is returned for parsing errors.
//line /usr/local/go/src/encoding/csv/reader.go:64
// Line numbers are 1-indexed and columns are 0-indexed.
//line /usr/local/go/src/encoding/csv/reader.go:66
type ParseError struct {
	StartLine	int	// Line where the record starts
	Line		int	// Line where the error occurred
	Column		int	// Column (1-based byte index) where the error occurred
	Err		error	// The actual error
}

func (e *ParseError) Error() string {
//line /usr/local/go/src/encoding/csv/reader.go:73
	_go_fuzz_dep_.CoverTab[114774]++
							if e.Err == ErrFieldCount {
//line /usr/local/go/src/encoding/csv/reader.go:74
		_go_fuzz_dep_.CoverTab[114777]++
								return fmt.Sprintf("record on line %d: %v", e.Line, e.Err)
//line /usr/local/go/src/encoding/csv/reader.go:75
		// _ = "end of CoverTab[114777]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:76
		_go_fuzz_dep_.CoverTab[114778]++
//line /usr/local/go/src/encoding/csv/reader.go:76
		// _ = "end of CoverTab[114778]"
//line /usr/local/go/src/encoding/csv/reader.go:76
	}
//line /usr/local/go/src/encoding/csv/reader.go:76
	// _ = "end of CoverTab[114774]"
//line /usr/local/go/src/encoding/csv/reader.go:76
	_go_fuzz_dep_.CoverTab[114775]++
							if e.StartLine != e.Line {
//line /usr/local/go/src/encoding/csv/reader.go:77
		_go_fuzz_dep_.CoverTab[114779]++
								return fmt.Sprintf("record on line %d; parse error on line %d, column %d: %v", e.StartLine, e.Line, e.Column, e.Err)
//line /usr/local/go/src/encoding/csv/reader.go:78
		// _ = "end of CoverTab[114779]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:79
		_go_fuzz_dep_.CoverTab[114780]++
//line /usr/local/go/src/encoding/csv/reader.go:79
		// _ = "end of CoverTab[114780]"
//line /usr/local/go/src/encoding/csv/reader.go:79
	}
//line /usr/local/go/src/encoding/csv/reader.go:79
	// _ = "end of CoverTab[114775]"
//line /usr/local/go/src/encoding/csv/reader.go:79
	_go_fuzz_dep_.CoverTab[114776]++
							return fmt.Sprintf("parse error on line %d, column %d: %v", e.Line, e.Column, e.Err)
//line /usr/local/go/src/encoding/csv/reader.go:80
	// _ = "end of CoverTab[114776]"
}

func (e *ParseError) Unwrap() error {
//line /usr/local/go/src/encoding/csv/reader.go:83
	_go_fuzz_dep_.CoverTab[114781]++
//line /usr/local/go/src/encoding/csv/reader.go:83
	return e.Err
//line /usr/local/go/src/encoding/csv/reader.go:83
	// _ = "end of CoverTab[114781]"
//line /usr/local/go/src/encoding/csv/reader.go:83
}

// These are the errors that can be returned in ParseError.Err.
var (
	ErrBareQuote	= errors.New("bare \" in non-quoted-field")
	ErrQuote	= errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount	= errors.New("wrong number of fields")

	// Deprecated: ErrTrailingComma is no longer used.
	ErrTrailingComma	= errors.New("extra delimiter at end of line")
)

var errInvalidDelim = errors.New("csv: invalid field or comment delimiter")

func validDelim(r rune) bool {
//line /usr/local/go/src/encoding/csv/reader.go:97
	_go_fuzz_dep_.CoverTab[114782]++
							return r != 0 && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:98
		_go_fuzz_dep_.CoverTab[114783]++
//line /usr/local/go/src/encoding/csv/reader.go:98
		return r != '"'
//line /usr/local/go/src/encoding/csv/reader.go:98
		// _ = "end of CoverTab[114783]"
//line /usr/local/go/src/encoding/csv/reader.go:98
	}() && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:98
		_go_fuzz_dep_.CoverTab[114784]++
//line /usr/local/go/src/encoding/csv/reader.go:98
		return r != '\r'
//line /usr/local/go/src/encoding/csv/reader.go:98
		// _ = "end of CoverTab[114784]"
//line /usr/local/go/src/encoding/csv/reader.go:98
	}() && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:98
		_go_fuzz_dep_.CoverTab[114785]++
//line /usr/local/go/src/encoding/csv/reader.go:98
		return r != '\n'
//line /usr/local/go/src/encoding/csv/reader.go:98
		// _ = "end of CoverTab[114785]"
//line /usr/local/go/src/encoding/csv/reader.go:98
	}() && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:98
		_go_fuzz_dep_.CoverTab[114786]++
//line /usr/local/go/src/encoding/csv/reader.go:98
		return utf8.ValidRune(r)
//line /usr/local/go/src/encoding/csv/reader.go:98
		// _ = "end of CoverTab[114786]"
//line /usr/local/go/src/encoding/csv/reader.go:98
	}() && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:98
		_go_fuzz_dep_.CoverTab[114787]++
//line /usr/local/go/src/encoding/csv/reader.go:98
		return r != utf8.RuneError
//line /usr/local/go/src/encoding/csv/reader.go:98
		// _ = "end of CoverTab[114787]"
//line /usr/local/go/src/encoding/csv/reader.go:98
	}()
//line /usr/local/go/src/encoding/csv/reader.go:98
	// _ = "end of CoverTab[114782]"
}

// A Reader reads records from a CSV-encoded file.
//line /usr/local/go/src/encoding/csv/reader.go:101
//
//line /usr/local/go/src/encoding/csv/reader.go:101
// As returned by NewReader, a Reader expects input conforming to RFC 4180.
//line /usr/local/go/src/encoding/csv/reader.go:101
// The exported fields can be changed to customize the details before the
//line /usr/local/go/src/encoding/csv/reader.go:101
// first call to Read or ReadAll.
//line /usr/local/go/src/encoding/csv/reader.go:101
//
//line /usr/local/go/src/encoding/csv/reader.go:101
// The Reader converts all \r\n sequences in its input to plain \n,
//line /usr/local/go/src/encoding/csv/reader.go:101
// including in multiline field values, so that the returned data does
//line /usr/local/go/src/encoding/csv/reader.go:101
// not depend on which line-ending convention an input file uses.
//line /usr/local/go/src/encoding/csv/reader.go:110
type Reader struct {
	// Comma is the field delimiter.
	// It is set to comma (',') by NewReader.
	// Comma must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	Comma	rune

	// Comment, if not 0, is the comment character. Lines beginning with the
	// Comment character without preceding whitespace are ignored.
	// With leading whitespace the Comment character becomes part of the
	// field, even if TrimLeadingSpace is true.
	// Comment must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	// It must also not be equal to Comma.
	Comment	rune

	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	FieldsPerRecord	int

	// If LazyQuotes is true, a quote may appear in an unquoted field and a
	// non-doubled quote may appear in a quoted field.
	LazyQuotes	bool

	// If TrimLeadingSpace is true, leading white space in a field is ignored.
	// This is done even if the field delimiter, Comma, is white space.
	TrimLeadingSpace	bool

	// ReuseRecord controls whether calls to Read may return a slice sharing
	// the backing array of the previous call's returned slice for performance.
	// By default, each call to Read returns newly allocated memory owned by the caller.
	ReuseRecord	bool

	// Deprecated: TrailingComma is no longer used.
	TrailingComma	bool

	r	*bufio.Reader

	// numLine is the current line being read in the CSV file.
	numLine	int

	// offset is the input stream byte offset of the current reader position.
	offset	int64

	// rawBuffer is a line buffer only used by the readLine method.
	rawBuffer	[]byte

	// recordBuffer holds the unescaped fields, one after another.
	// The fields can be accessed by using the indexes in fieldIndexes.
	// E.g., For the row `a,"b","c""d",e`, recordBuffer will contain `abc"de`
	// and fieldIndexes will contain the indexes [1, 2, 5, 6].
	recordBuffer	[]byte

	// fieldIndexes is an index of fields inside recordBuffer.
	// The i'th field ends at offset fieldIndexes[i] in recordBuffer.
	fieldIndexes	[]int

	// fieldPositions is an index of field positions for the
	// last record returned by Read.
	fieldPositions	[]position

	// lastRecord is a record cache and only used when ReuseRecord == true.
	lastRecord	[]string
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
//line /usr/local/go/src/encoding/csv/reader.go:180
	_go_fuzz_dep_.CoverTab[114788]++
							return &Reader{
		Comma:	',',
		r:	bufio.NewReader(r),
	}
//line /usr/local/go/src/encoding/csv/reader.go:184
	// _ = "end of CoverTab[114788]"
}

// Read reads one record (a slice of fields) from r.
//line /usr/local/go/src/encoding/csv/reader.go:187
// If the record has an unexpected number of fields,
//line /usr/local/go/src/encoding/csv/reader.go:187
// Read returns the record along with the error ErrFieldCount.
//line /usr/local/go/src/encoding/csv/reader.go:187
// Except for that case, Read always returns either a non-nil
//line /usr/local/go/src/encoding/csv/reader.go:187
// record or a non-nil error, but not both.
//line /usr/local/go/src/encoding/csv/reader.go:187
// If there is no data left to be read, Read returns nil, io.EOF.
//line /usr/local/go/src/encoding/csv/reader.go:187
// If ReuseRecord is true, the returned slice may be shared
//line /usr/local/go/src/encoding/csv/reader.go:187
// between multiple calls to Read.
//line /usr/local/go/src/encoding/csv/reader.go:195
func (r *Reader) Read() (record []string, err error) {
//line /usr/local/go/src/encoding/csv/reader.go:195
	_go_fuzz_dep_.CoverTab[114789]++
							if r.ReuseRecord {
//line /usr/local/go/src/encoding/csv/reader.go:196
		_go_fuzz_dep_.CoverTab[114791]++
								record, err = r.readRecord(r.lastRecord)
								r.lastRecord = record
//line /usr/local/go/src/encoding/csv/reader.go:198
		// _ = "end of CoverTab[114791]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:199
		_go_fuzz_dep_.CoverTab[114792]++
								record, err = r.readRecord(nil)
//line /usr/local/go/src/encoding/csv/reader.go:200
		// _ = "end of CoverTab[114792]"
	}
//line /usr/local/go/src/encoding/csv/reader.go:201
	// _ = "end of CoverTab[114789]"
//line /usr/local/go/src/encoding/csv/reader.go:201
	_go_fuzz_dep_.CoverTab[114790]++
							return record, err
//line /usr/local/go/src/encoding/csv/reader.go:202
	// _ = "end of CoverTab[114790]"
}

// FieldPos returns the line and column corresponding to
//line /usr/local/go/src/encoding/csv/reader.go:205
// the start of the field with the given index in the slice most recently
//line /usr/local/go/src/encoding/csv/reader.go:205
// returned by Read. Numbering of lines and columns starts at 1;
//line /usr/local/go/src/encoding/csv/reader.go:205
// columns are counted in bytes, not runes.
//line /usr/local/go/src/encoding/csv/reader.go:205
//
//line /usr/local/go/src/encoding/csv/reader.go:205
// If this is called with an out-of-bounds index, it panics.
//line /usr/local/go/src/encoding/csv/reader.go:211
func (r *Reader) FieldPos(field int) (line, column int) {
//line /usr/local/go/src/encoding/csv/reader.go:211
	_go_fuzz_dep_.CoverTab[114793]++
							if field < 0 || func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:212
		_go_fuzz_dep_.CoverTab[114795]++
//line /usr/local/go/src/encoding/csv/reader.go:212
		return field >= len(r.fieldPositions)
//line /usr/local/go/src/encoding/csv/reader.go:212
		// _ = "end of CoverTab[114795]"
//line /usr/local/go/src/encoding/csv/reader.go:212
	}() {
//line /usr/local/go/src/encoding/csv/reader.go:212
		_go_fuzz_dep_.CoverTab[114796]++
								panic("out of range index passed to FieldPos")
//line /usr/local/go/src/encoding/csv/reader.go:213
		// _ = "end of CoverTab[114796]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:214
		_go_fuzz_dep_.CoverTab[114797]++
//line /usr/local/go/src/encoding/csv/reader.go:214
		// _ = "end of CoverTab[114797]"
//line /usr/local/go/src/encoding/csv/reader.go:214
	}
//line /usr/local/go/src/encoding/csv/reader.go:214
	// _ = "end of CoverTab[114793]"
//line /usr/local/go/src/encoding/csv/reader.go:214
	_go_fuzz_dep_.CoverTab[114794]++
							p := &r.fieldPositions[field]
							return p.line, p.col
//line /usr/local/go/src/encoding/csv/reader.go:216
	// _ = "end of CoverTab[114794]"
}

// InputOffset returns the input stream byte offset of the current reader
//line /usr/local/go/src/encoding/csv/reader.go:219
// position. The offset gives the location of the end of the most recently
//line /usr/local/go/src/encoding/csv/reader.go:219
// read row and the beginning of the next row.
//line /usr/local/go/src/encoding/csv/reader.go:222
func (r *Reader) InputOffset() int64 {
//line /usr/local/go/src/encoding/csv/reader.go:222
	_go_fuzz_dep_.CoverTab[114798]++
							return r.offset
//line /usr/local/go/src/encoding/csv/reader.go:223
	// _ = "end of CoverTab[114798]"
}

// pos holds the position of a field in the current line.
type position struct {
	line, col int
}

// ReadAll reads all the remaining records from r.
//line /usr/local/go/src/encoding/csv/reader.go:231
// Each record is a slice of fields.
//line /usr/local/go/src/encoding/csv/reader.go:231
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
//line /usr/local/go/src/encoding/csv/reader.go:231
// defined to read until EOF, it does not treat end of file as an error to be
//line /usr/local/go/src/encoding/csv/reader.go:231
// reported.
//line /usr/local/go/src/encoding/csv/reader.go:236
func (r *Reader) ReadAll() (records [][]string, err error) {
//line /usr/local/go/src/encoding/csv/reader.go:236
	_go_fuzz_dep_.CoverTab[114799]++
							for {
//line /usr/local/go/src/encoding/csv/reader.go:237
		_go_fuzz_dep_.CoverTab[114800]++
								record, err := r.readRecord(nil)
								if err == io.EOF {
//line /usr/local/go/src/encoding/csv/reader.go:239
			_go_fuzz_dep_.CoverTab[114803]++
									return records, nil
//line /usr/local/go/src/encoding/csv/reader.go:240
			// _ = "end of CoverTab[114803]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:241
			_go_fuzz_dep_.CoverTab[114804]++
//line /usr/local/go/src/encoding/csv/reader.go:241
			// _ = "end of CoverTab[114804]"
//line /usr/local/go/src/encoding/csv/reader.go:241
		}
//line /usr/local/go/src/encoding/csv/reader.go:241
		// _ = "end of CoverTab[114800]"
//line /usr/local/go/src/encoding/csv/reader.go:241
		_go_fuzz_dep_.CoverTab[114801]++
								if err != nil {
//line /usr/local/go/src/encoding/csv/reader.go:242
			_go_fuzz_dep_.CoverTab[114805]++
									return nil, err
//line /usr/local/go/src/encoding/csv/reader.go:243
			// _ = "end of CoverTab[114805]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:244
			_go_fuzz_dep_.CoverTab[114806]++
//line /usr/local/go/src/encoding/csv/reader.go:244
			// _ = "end of CoverTab[114806]"
//line /usr/local/go/src/encoding/csv/reader.go:244
		}
//line /usr/local/go/src/encoding/csv/reader.go:244
		// _ = "end of CoverTab[114801]"
//line /usr/local/go/src/encoding/csv/reader.go:244
		_go_fuzz_dep_.CoverTab[114802]++
								records = append(records, record)
//line /usr/local/go/src/encoding/csv/reader.go:245
		// _ = "end of CoverTab[114802]"
	}
//line /usr/local/go/src/encoding/csv/reader.go:246
	// _ = "end of CoverTab[114799]"
}

// readLine reads the next line (with the trailing endline).
//line /usr/local/go/src/encoding/csv/reader.go:249
// If EOF is hit without a trailing endline, it will be omitted.
//line /usr/local/go/src/encoding/csv/reader.go:249
// If some bytes were read, then the error is never io.EOF.
//line /usr/local/go/src/encoding/csv/reader.go:249
// The result is only valid until the next call to readLine.
//line /usr/local/go/src/encoding/csv/reader.go:253
func (r *Reader) readLine() ([]byte, error) {
//line /usr/local/go/src/encoding/csv/reader.go:253
	_go_fuzz_dep_.CoverTab[114807]++
							line, err := r.r.ReadSlice('\n')
							if err == bufio.ErrBufferFull {
//line /usr/local/go/src/encoding/csv/reader.go:255
		_go_fuzz_dep_.CoverTab[114811]++
								r.rawBuffer = append(r.rawBuffer[:0], line...)
								for err == bufio.ErrBufferFull {
//line /usr/local/go/src/encoding/csv/reader.go:257
			_go_fuzz_dep_.CoverTab[114813]++
									line, err = r.r.ReadSlice('\n')
									r.rawBuffer = append(r.rawBuffer, line...)
//line /usr/local/go/src/encoding/csv/reader.go:259
			// _ = "end of CoverTab[114813]"
		}
//line /usr/local/go/src/encoding/csv/reader.go:260
		// _ = "end of CoverTab[114811]"
//line /usr/local/go/src/encoding/csv/reader.go:260
		_go_fuzz_dep_.CoverTab[114812]++
								line = r.rawBuffer
//line /usr/local/go/src/encoding/csv/reader.go:261
		// _ = "end of CoverTab[114812]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:262
		_go_fuzz_dep_.CoverTab[114814]++
//line /usr/local/go/src/encoding/csv/reader.go:262
		// _ = "end of CoverTab[114814]"
//line /usr/local/go/src/encoding/csv/reader.go:262
	}
//line /usr/local/go/src/encoding/csv/reader.go:262
	// _ = "end of CoverTab[114807]"
//line /usr/local/go/src/encoding/csv/reader.go:262
	_go_fuzz_dep_.CoverTab[114808]++
							readSize := len(line)
							if readSize > 0 && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:264
		_go_fuzz_dep_.CoverTab[114815]++
//line /usr/local/go/src/encoding/csv/reader.go:264
		return err == io.EOF
//line /usr/local/go/src/encoding/csv/reader.go:264
		// _ = "end of CoverTab[114815]"
//line /usr/local/go/src/encoding/csv/reader.go:264
	}() {
//line /usr/local/go/src/encoding/csv/reader.go:264
		_go_fuzz_dep_.CoverTab[114816]++
								err = nil

								if line[readSize-1] == '\r' {
//line /usr/local/go/src/encoding/csv/reader.go:267
			_go_fuzz_dep_.CoverTab[114817]++
									line = line[:readSize-1]
//line /usr/local/go/src/encoding/csv/reader.go:268
			// _ = "end of CoverTab[114817]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:269
			_go_fuzz_dep_.CoverTab[114818]++
//line /usr/local/go/src/encoding/csv/reader.go:269
			// _ = "end of CoverTab[114818]"
//line /usr/local/go/src/encoding/csv/reader.go:269
		}
//line /usr/local/go/src/encoding/csv/reader.go:269
		// _ = "end of CoverTab[114816]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:270
		_go_fuzz_dep_.CoverTab[114819]++
//line /usr/local/go/src/encoding/csv/reader.go:270
		// _ = "end of CoverTab[114819]"
//line /usr/local/go/src/encoding/csv/reader.go:270
	}
//line /usr/local/go/src/encoding/csv/reader.go:270
	// _ = "end of CoverTab[114808]"
//line /usr/local/go/src/encoding/csv/reader.go:270
	_go_fuzz_dep_.CoverTab[114809]++
							r.numLine++
							r.offset += int64(readSize)

							if n := len(line); n >= 2 && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:274
		_go_fuzz_dep_.CoverTab[114820]++
//line /usr/local/go/src/encoding/csv/reader.go:274
		return line[n-2] == '\r'
//line /usr/local/go/src/encoding/csv/reader.go:274
		// _ = "end of CoverTab[114820]"
//line /usr/local/go/src/encoding/csv/reader.go:274
	}() && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:274
		_go_fuzz_dep_.CoverTab[114821]++
//line /usr/local/go/src/encoding/csv/reader.go:274
		return line[n-1] == '\n'
//line /usr/local/go/src/encoding/csv/reader.go:274
		// _ = "end of CoverTab[114821]"
//line /usr/local/go/src/encoding/csv/reader.go:274
	}() {
//line /usr/local/go/src/encoding/csv/reader.go:274
		_go_fuzz_dep_.CoverTab[114822]++
								line[n-2] = '\n'
								line = line[:n-1]
//line /usr/local/go/src/encoding/csv/reader.go:276
		// _ = "end of CoverTab[114822]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:277
		_go_fuzz_dep_.CoverTab[114823]++
//line /usr/local/go/src/encoding/csv/reader.go:277
		// _ = "end of CoverTab[114823]"
//line /usr/local/go/src/encoding/csv/reader.go:277
	}
//line /usr/local/go/src/encoding/csv/reader.go:277
	// _ = "end of CoverTab[114809]"
//line /usr/local/go/src/encoding/csv/reader.go:277
	_go_fuzz_dep_.CoverTab[114810]++
							return line, err
//line /usr/local/go/src/encoding/csv/reader.go:278
	// _ = "end of CoverTab[114810]"
}

// lengthNL reports the number of bytes for the trailing \n.
func lengthNL(b []byte) int {
//line /usr/local/go/src/encoding/csv/reader.go:282
	_go_fuzz_dep_.CoverTab[114824]++
							if len(b) > 0 && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:283
		_go_fuzz_dep_.CoverTab[114826]++
//line /usr/local/go/src/encoding/csv/reader.go:283
		return b[len(b)-1] == '\n'
//line /usr/local/go/src/encoding/csv/reader.go:283
		// _ = "end of CoverTab[114826]"
//line /usr/local/go/src/encoding/csv/reader.go:283
	}() {
//line /usr/local/go/src/encoding/csv/reader.go:283
		_go_fuzz_dep_.CoverTab[114827]++
								return 1
//line /usr/local/go/src/encoding/csv/reader.go:284
		// _ = "end of CoverTab[114827]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:285
		_go_fuzz_dep_.CoverTab[114828]++
//line /usr/local/go/src/encoding/csv/reader.go:285
		// _ = "end of CoverTab[114828]"
//line /usr/local/go/src/encoding/csv/reader.go:285
	}
//line /usr/local/go/src/encoding/csv/reader.go:285
	// _ = "end of CoverTab[114824]"
//line /usr/local/go/src/encoding/csv/reader.go:285
	_go_fuzz_dep_.CoverTab[114825]++
							return 0
//line /usr/local/go/src/encoding/csv/reader.go:286
	// _ = "end of CoverTab[114825]"
}

// nextRune returns the next rune in b or utf8.RuneError.
func nextRune(b []byte) rune {
//line /usr/local/go/src/encoding/csv/reader.go:290
	_go_fuzz_dep_.CoverTab[114829]++
							r, _ := utf8.DecodeRune(b)
							return r
//line /usr/local/go/src/encoding/csv/reader.go:292
	// _ = "end of CoverTab[114829]"
}

func (r *Reader) readRecord(dst []string) ([]string, error) {
//line /usr/local/go/src/encoding/csv/reader.go:295
	_go_fuzz_dep_.CoverTab[114830]++
							if r.Comma == r.Comment || func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:296
		_go_fuzz_dep_.CoverTab[114839]++
//line /usr/local/go/src/encoding/csv/reader.go:296
		return !validDelim(r.Comma)
//line /usr/local/go/src/encoding/csv/reader.go:296
		// _ = "end of CoverTab[114839]"
//line /usr/local/go/src/encoding/csv/reader.go:296
	}() || func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:296
		_go_fuzz_dep_.CoverTab[114840]++
//line /usr/local/go/src/encoding/csv/reader.go:296
		return (r.Comment != 0 && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:296
			_go_fuzz_dep_.CoverTab[114841]++
//line /usr/local/go/src/encoding/csv/reader.go:296
			return !validDelim(r.Comment)
//line /usr/local/go/src/encoding/csv/reader.go:296
			// _ = "end of CoverTab[114841]"
//line /usr/local/go/src/encoding/csv/reader.go:296
		}())
//line /usr/local/go/src/encoding/csv/reader.go:296
		// _ = "end of CoverTab[114840]"
//line /usr/local/go/src/encoding/csv/reader.go:296
	}() {
//line /usr/local/go/src/encoding/csv/reader.go:296
		_go_fuzz_dep_.CoverTab[114842]++
								return nil, errInvalidDelim
//line /usr/local/go/src/encoding/csv/reader.go:297
		// _ = "end of CoverTab[114842]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:298
		_go_fuzz_dep_.CoverTab[114843]++
//line /usr/local/go/src/encoding/csv/reader.go:298
		// _ = "end of CoverTab[114843]"
//line /usr/local/go/src/encoding/csv/reader.go:298
	}
//line /usr/local/go/src/encoding/csv/reader.go:298
	// _ = "end of CoverTab[114830]"
//line /usr/local/go/src/encoding/csv/reader.go:298
	_go_fuzz_dep_.CoverTab[114831]++

	// Read line (automatically skipping past empty lines and any comments).
	var line []byte
	var errRead error
	for errRead == nil {
//line /usr/local/go/src/encoding/csv/reader.go:303
		_go_fuzz_dep_.CoverTab[114844]++
								line, errRead = r.readLine()
								if r.Comment != 0 && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:305
			_go_fuzz_dep_.CoverTab[114847]++
//line /usr/local/go/src/encoding/csv/reader.go:305
			return nextRune(line) == r.Comment
//line /usr/local/go/src/encoding/csv/reader.go:305
			// _ = "end of CoverTab[114847]"
//line /usr/local/go/src/encoding/csv/reader.go:305
		}() {
//line /usr/local/go/src/encoding/csv/reader.go:305
			_go_fuzz_dep_.CoverTab[114848]++
									line = nil
									continue
//line /usr/local/go/src/encoding/csv/reader.go:307
			// _ = "end of CoverTab[114848]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:308
			_go_fuzz_dep_.CoverTab[114849]++
//line /usr/local/go/src/encoding/csv/reader.go:308
			// _ = "end of CoverTab[114849]"
//line /usr/local/go/src/encoding/csv/reader.go:308
		}
//line /usr/local/go/src/encoding/csv/reader.go:308
		// _ = "end of CoverTab[114844]"
//line /usr/local/go/src/encoding/csv/reader.go:308
		_go_fuzz_dep_.CoverTab[114845]++
								if errRead == nil && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:309
			_go_fuzz_dep_.CoverTab[114850]++
//line /usr/local/go/src/encoding/csv/reader.go:309
			return len(line) == lengthNL(line)
//line /usr/local/go/src/encoding/csv/reader.go:309
			// _ = "end of CoverTab[114850]"
//line /usr/local/go/src/encoding/csv/reader.go:309
		}() {
//line /usr/local/go/src/encoding/csv/reader.go:309
			_go_fuzz_dep_.CoverTab[114851]++
									line = nil
									continue
//line /usr/local/go/src/encoding/csv/reader.go:311
			// _ = "end of CoverTab[114851]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:312
			_go_fuzz_dep_.CoverTab[114852]++
//line /usr/local/go/src/encoding/csv/reader.go:312
			// _ = "end of CoverTab[114852]"
//line /usr/local/go/src/encoding/csv/reader.go:312
		}
//line /usr/local/go/src/encoding/csv/reader.go:312
		// _ = "end of CoverTab[114845]"
//line /usr/local/go/src/encoding/csv/reader.go:312
		_go_fuzz_dep_.CoverTab[114846]++
								break
//line /usr/local/go/src/encoding/csv/reader.go:313
		// _ = "end of CoverTab[114846]"
	}
//line /usr/local/go/src/encoding/csv/reader.go:314
	// _ = "end of CoverTab[114831]"
//line /usr/local/go/src/encoding/csv/reader.go:314
	_go_fuzz_dep_.CoverTab[114832]++
							if errRead == io.EOF {
//line /usr/local/go/src/encoding/csv/reader.go:315
		_go_fuzz_dep_.CoverTab[114853]++
								return nil, errRead
//line /usr/local/go/src/encoding/csv/reader.go:316
		// _ = "end of CoverTab[114853]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:317
		_go_fuzz_dep_.CoverTab[114854]++
//line /usr/local/go/src/encoding/csv/reader.go:317
		// _ = "end of CoverTab[114854]"
//line /usr/local/go/src/encoding/csv/reader.go:317
	}
//line /usr/local/go/src/encoding/csv/reader.go:317
	// _ = "end of CoverTab[114832]"
//line /usr/local/go/src/encoding/csv/reader.go:317
	_go_fuzz_dep_.CoverTab[114833]++

	// Parse each field in the record.
	var err error
	const quoteLen = len(`"`)
	commaLen := utf8.RuneLen(r.Comma)
	recLine := r.numLine
	r.recordBuffer = r.recordBuffer[:0]
	r.fieldIndexes = r.fieldIndexes[:0]
	r.fieldPositions = r.fieldPositions[:0]
	pos := position{line: r.numLine, col: 1}
parseField:
	for {
//line /usr/local/go/src/encoding/csv/reader.go:329
		_go_fuzz_dep_.CoverTab[114855]++
								if r.TrimLeadingSpace {
//line /usr/local/go/src/encoding/csv/reader.go:330
			_go_fuzz_dep_.CoverTab[114857]++
									i := bytes.IndexFunc(line, func(r rune) bool {
//line /usr/local/go/src/encoding/csv/reader.go:331
				_go_fuzz_dep_.CoverTab[114860]++
										return !unicode.IsSpace(r)
//line /usr/local/go/src/encoding/csv/reader.go:332
				// _ = "end of CoverTab[114860]"
			})
//line /usr/local/go/src/encoding/csv/reader.go:333
			// _ = "end of CoverTab[114857]"
//line /usr/local/go/src/encoding/csv/reader.go:333
			_go_fuzz_dep_.CoverTab[114858]++
									if i < 0 {
//line /usr/local/go/src/encoding/csv/reader.go:334
				_go_fuzz_dep_.CoverTab[114861]++
										i = len(line)
										pos.col -= lengthNL(line)
//line /usr/local/go/src/encoding/csv/reader.go:336
				// _ = "end of CoverTab[114861]"
			} else {
//line /usr/local/go/src/encoding/csv/reader.go:337
				_go_fuzz_dep_.CoverTab[114862]++
//line /usr/local/go/src/encoding/csv/reader.go:337
				// _ = "end of CoverTab[114862]"
//line /usr/local/go/src/encoding/csv/reader.go:337
			}
//line /usr/local/go/src/encoding/csv/reader.go:337
			// _ = "end of CoverTab[114858]"
//line /usr/local/go/src/encoding/csv/reader.go:337
			_go_fuzz_dep_.CoverTab[114859]++
									line = line[i:]
									pos.col += i
//line /usr/local/go/src/encoding/csv/reader.go:339
			// _ = "end of CoverTab[114859]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:340
			_go_fuzz_dep_.CoverTab[114863]++
//line /usr/local/go/src/encoding/csv/reader.go:340
			// _ = "end of CoverTab[114863]"
//line /usr/local/go/src/encoding/csv/reader.go:340
		}
//line /usr/local/go/src/encoding/csv/reader.go:340
		// _ = "end of CoverTab[114855]"
//line /usr/local/go/src/encoding/csv/reader.go:340
		_go_fuzz_dep_.CoverTab[114856]++
								if len(line) == 0 || func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:341
			_go_fuzz_dep_.CoverTab[114864]++
//line /usr/local/go/src/encoding/csv/reader.go:341
			return line[0] != '"'
//line /usr/local/go/src/encoding/csv/reader.go:341
			// _ = "end of CoverTab[114864]"
//line /usr/local/go/src/encoding/csv/reader.go:341
		}() {
//line /usr/local/go/src/encoding/csv/reader.go:341
			_go_fuzz_dep_.CoverTab[114865]++

									i := bytes.IndexRune(line, r.Comma)
									field := line
									if i >= 0 {
//line /usr/local/go/src/encoding/csv/reader.go:345
				_go_fuzz_dep_.CoverTab[114869]++
										field = field[:i]
//line /usr/local/go/src/encoding/csv/reader.go:346
				// _ = "end of CoverTab[114869]"
			} else {
//line /usr/local/go/src/encoding/csv/reader.go:347
				_go_fuzz_dep_.CoverTab[114870]++
										field = field[:len(field)-lengthNL(field)]
//line /usr/local/go/src/encoding/csv/reader.go:348
				// _ = "end of CoverTab[114870]"
			}
//line /usr/local/go/src/encoding/csv/reader.go:349
			// _ = "end of CoverTab[114865]"
//line /usr/local/go/src/encoding/csv/reader.go:349
			_go_fuzz_dep_.CoverTab[114866]++

									if !r.LazyQuotes {
//line /usr/local/go/src/encoding/csv/reader.go:351
				_go_fuzz_dep_.CoverTab[114871]++
										if j := bytes.IndexByte(field, '"'); j >= 0 {
//line /usr/local/go/src/encoding/csv/reader.go:352
					_go_fuzz_dep_.CoverTab[114872]++
											col := pos.col + j
											err = &ParseError{StartLine: recLine, Line: r.numLine, Column: col, Err: ErrBareQuote}
											break parseField
//line /usr/local/go/src/encoding/csv/reader.go:355
					// _ = "end of CoverTab[114872]"
				} else {
//line /usr/local/go/src/encoding/csv/reader.go:356
					_go_fuzz_dep_.CoverTab[114873]++
//line /usr/local/go/src/encoding/csv/reader.go:356
					// _ = "end of CoverTab[114873]"
//line /usr/local/go/src/encoding/csv/reader.go:356
				}
//line /usr/local/go/src/encoding/csv/reader.go:356
				// _ = "end of CoverTab[114871]"
			} else {
//line /usr/local/go/src/encoding/csv/reader.go:357
				_go_fuzz_dep_.CoverTab[114874]++
//line /usr/local/go/src/encoding/csv/reader.go:357
				// _ = "end of CoverTab[114874]"
//line /usr/local/go/src/encoding/csv/reader.go:357
			}
//line /usr/local/go/src/encoding/csv/reader.go:357
			// _ = "end of CoverTab[114866]"
//line /usr/local/go/src/encoding/csv/reader.go:357
			_go_fuzz_dep_.CoverTab[114867]++
									r.recordBuffer = append(r.recordBuffer, field...)
									r.fieldIndexes = append(r.fieldIndexes, len(r.recordBuffer))
									r.fieldPositions = append(r.fieldPositions, pos)
									if i >= 0 {
//line /usr/local/go/src/encoding/csv/reader.go:361
				_go_fuzz_dep_.CoverTab[114875]++
										line = line[i+commaLen:]
										pos.col += i + commaLen
										continue parseField
//line /usr/local/go/src/encoding/csv/reader.go:364
				// _ = "end of CoverTab[114875]"
			} else {
//line /usr/local/go/src/encoding/csv/reader.go:365
				_go_fuzz_dep_.CoverTab[114876]++
//line /usr/local/go/src/encoding/csv/reader.go:365
				// _ = "end of CoverTab[114876]"
//line /usr/local/go/src/encoding/csv/reader.go:365
			}
//line /usr/local/go/src/encoding/csv/reader.go:365
			// _ = "end of CoverTab[114867]"
//line /usr/local/go/src/encoding/csv/reader.go:365
			_go_fuzz_dep_.CoverTab[114868]++
									break parseField
//line /usr/local/go/src/encoding/csv/reader.go:366
			// _ = "end of CoverTab[114868]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:367
			_go_fuzz_dep_.CoverTab[114877]++

									fieldPos := pos
									line = line[quoteLen:]
									pos.col += quoteLen
									for {
//line /usr/local/go/src/encoding/csv/reader.go:372
				_go_fuzz_dep_.CoverTab[114878]++
										i := bytes.IndexByte(line, '"')
										if i >= 0 {
//line /usr/local/go/src/encoding/csv/reader.go:374
					_go_fuzz_dep_.CoverTab[114879]++

											r.recordBuffer = append(r.recordBuffer, line[:i]...)
											line = line[i+quoteLen:]
											pos.col += i + quoteLen
											switch rn := nextRune(line); {
					case rn == '"':
//line /usr/local/go/src/encoding/csv/reader.go:380
						_go_fuzz_dep_.CoverTab[114880]++

												r.recordBuffer = append(r.recordBuffer, '"')
												line = line[quoteLen:]
												pos.col += quoteLen
//line /usr/local/go/src/encoding/csv/reader.go:384
						// _ = "end of CoverTab[114880]"
					case rn == r.Comma:
//line /usr/local/go/src/encoding/csv/reader.go:385
						_go_fuzz_dep_.CoverTab[114881]++

												line = line[commaLen:]
												pos.col += commaLen
												r.fieldIndexes = append(r.fieldIndexes, len(r.recordBuffer))
												r.fieldPositions = append(r.fieldPositions, fieldPos)
												continue parseField
//line /usr/local/go/src/encoding/csv/reader.go:391
						// _ = "end of CoverTab[114881]"
					case lengthNL(line) == len(line):
//line /usr/local/go/src/encoding/csv/reader.go:392
						_go_fuzz_dep_.CoverTab[114882]++

												r.fieldIndexes = append(r.fieldIndexes, len(r.recordBuffer))
												r.fieldPositions = append(r.fieldPositions, fieldPos)
												break parseField
//line /usr/local/go/src/encoding/csv/reader.go:396
						// _ = "end of CoverTab[114882]"
					case r.LazyQuotes:
//line /usr/local/go/src/encoding/csv/reader.go:397
						_go_fuzz_dep_.CoverTab[114883]++

												r.recordBuffer = append(r.recordBuffer, '"')
//line /usr/local/go/src/encoding/csv/reader.go:399
						// _ = "end of CoverTab[114883]"
					default:
//line /usr/local/go/src/encoding/csv/reader.go:400
						_go_fuzz_dep_.CoverTab[114884]++

												err = &ParseError{StartLine: recLine, Line: r.numLine, Column: pos.col - quoteLen, Err: ErrQuote}
												break parseField
//line /usr/local/go/src/encoding/csv/reader.go:403
						// _ = "end of CoverTab[114884]"
					}
//line /usr/local/go/src/encoding/csv/reader.go:404
					// _ = "end of CoverTab[114879]"
				} else {
//line /usr/local/go/src/encoding/csv/reader.go:405
					_go_fuzz_dep_.CoverTab[114885]++
//line /usr/local/go/src/encoding/csv/reader.go:405
					if len(line) > 0 {
//line /usr/local/go/src/encoding/csv/reader.go:405
						_go_fuzz_dep_.CoverTab[114886]++

												r.recordBuffer = append(r.recordBuffer, line...)
												if errRead != nil {
//line /usr/local/go/src/encoding/csv/reader.go:408
							_go_fuzz_dep_.CoverTab[114889]++
													break parseField
//line /usr/local/go/src/encoding/csv/reader.go:409
							// _ = "end of CoverTab[114889]"
						} else {
//line /usr/local/go/src/encoding/csv/reader.go:410
							_go_fuzz_dep_.CoverTab[114890]++
//line /usr/local/go/src/encoding/csv/reader.go:410
							// _ = "end of CoverTab[114890]"
//line /usr/local/go/src/encoding/csv/reader.go:410
						}
//line /usr/local/go/src/encoding/csv/reader.go:410
						// _ = "end of CoverTab[114886]"
//line /usr/local/go/src/encoding/csv/reader.go:410
						_go_fuzz_dep_.CoverTab[114887]++
												pos.col += len(line)
												line, errRead = r.readLine()
												if len(line) > 0 {
//line /usr/local/go/src/encoding/csv/reader.go:413
							_go_fuzz_dep_.CoverTab[114891]++
													pos.line++
													pos.col = 1
//line /usr/local/go/src/encoding/csv/reader.go:415
							// _ = "end of CoverTab[114891]"
						} else {
//line /usr/local/go/src/encoding/csv/reader.go:416
							_go_fuzz_dep_.CoverTab[114892]++
//line /usr/local/go/src/encoding/csv/reader.go:416
							// _ = "end of CoverTab[114892]"
//line /usr/local/go/src/encoding/csv/reader.go:416
						}
//line /usr/local/go/src/encoding/csv/reader.go:416
						// _ = "end of CoverTab[114887]"
//line /usr/local/go/src/encoding/csv/reader.go:416
						_go_fuzz_dep_.CoverTab[114888]++
												if errRead == io.EOF {
//line /usr/local/go/src/encoding/csv/reader.go:417
							_go_fuzz_dep_.CoverTab[114893]++
													errRead = nil
//line /usr/local/go/src/encoding/csv/reader.go:418
							// _ = "end of CoverTab[114893]"
						} else {
//line /usr/local/go/src/encoding/csv/reader.go:419
							_go_fuzz_dep_.CoverTab[114894]++
//line /usr/local/go/src/encoding/csv/reader.go:419
							// _ = "end of CoverTab[114894]"
//line /usr/local/go/src/encoding/csv/reader.go:419
						}
//line /usr/local/go/src/encoding/csv/reader.go:419
						// _ = "end of CoverTab[114888]"
					} else {
//line /usr/local/go/src/encoding/csv/reader.go:420
						_go_fuzz_dep_.CoverTab[114895]++

												if !r.LazyQuotes && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:422
							_go_fuzz_dep_.CoverTab[114897]++
//line /usr/local/go/src/encoding/csv/reader.go:422
							return errRead == nil
//line /usr/local/go/src/encoding/csv/reader.go:422
							// _ = "end of CoverTab[114897]"
//line /usr/local/go/src/encoding/csv/reader.go:422
						}() {
//line /usr/local/go/src/encoding/csv/reader.go:422
							_go_fuzz_dep_.CoverTab[114898]++
													err = &ParseError{StartLine: recLine, Line: pos.line, Column: pos.col, Err: ErrQuote}
													break parseField
//line /usr/local/go/src/encoding/csv/reader.go:424
							// _ = "end of CoverTab[114898]"
						} else {
//line /usr/local/go/src/encoding/csv/reader.go:425
							_go_fuzz_dep_.CoverTab[114899]++
//line /usr/local/go/src/encoding/csv/reader.go:425
							// _ = "end of CoverTab[114899]"
//line /usr/local/go/src/encoding/csv/reader.go:425
						}
//line /usr/local/go/src/encoding/csv/reader.go:425
						// _ = "end of CoverTab[114895]"
//line /usr/local/go/src/encoding/csv/reader.go:425
						_go_fuzz_dep_.CoverTab[114896]++
												r.fieldIndexes = append(r.fieldIndexes, len(r.recordBuffer))
												r.fieldPositions = append(r.fieldPositions, fieldPos)
												break parseField
//line /usr/local/go/src/encoding/csv/reader.go:428
						// _ = "end of CoverTab[114896]"
					}
//line /usr/local/go/src/encoding/csv/reader.go:429
					// _ = "end of CoverTab[114885]"
//line /usr/local/go/src/encoding/csv/reader.go:429
				}
//line /usr/local/go/src/encoding/csv/reader.go:429
				// _ = "end of CoverTab[114878]"
			}
//line /usr/local/go/src/encoding/csv/reader.go:430
			// _ = "end of CoverTab[114877]"
		}
//line /usr/local/go/src/encoding/csv/reader.go:431
		// _ = "end of CoverTab[114856]"
	}
//line /usr/local/go/src/encoding/csv/reader.go:432
	// _ = "end of CoverTab[114833]"
//line /usr/local/go/src/encoding/csv/reader.go:432
	_go_fuzz_dep_.CoverTab[114834]++
							if err == nil {
//line /usr/local/go/src/encoding/csv/reader.go:433
		_go_fuzz_dep_.CoverTab[114900]++
								err = errRead
//line /usr/local/go/src/encoding/csv/reader.go:434
		// _ = "end of CoverTab[114900]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:435
		_go_fuzz_dep_.CoverTab[114901]++
//line /usr/local/go/src/encoding/csv/reader.go:435
		// _ = "end of CoverTab[114901]"
//line /usr/local/go/src/encoding/csv/reader.go:435
	}
//line /usr/local/go/src/encoding/csv/reader.go:435
	// _ = "end of CoverTab[114834]"
//line /usr/local/go/src/encoding/csv/reader.go:435
	_go_fuzz_dep_.CoverTab[114835]++

//line /usr/local/go/src/encoding/csv/reader.go:439
	str := string(r.recordBuffer)
	dst = dst[:0]
	if cap(dst) < len(r.fieldIndexes) {
//line /usr/local/go/src/encoding/csv/reader.go:441
		_go_fuzz_dep_.CoverTab[114902]++
								dst = make([]string, len(r.fieldIndexes))
//line /usr/local/go/src/encoding/csv/reader.go:442
		// _ = "end of CoverTab[114902]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:443
		_go_fuzz_dep_.CoverTab[114903]++
//line /usr/local/go/src/encoding/csv/reader.go:443
		// _ = "end of CoverTab[114903]"
//line /usr/local/go/src/encoding/csv/reader.go:443
	}
//line /usr/local/go/src/encoding/csv/reader.go:443
	// _ = "end of CoverTab[114835]"
//line /usr/local/go/src/encoding/csv/reader.go:443
	_go_fuzz_dep_.CoverTab[114836]++
							dst = dst[:len(r.fieldIndexes)]
							var preIdx int
							for i, idx := range r.fieldIndexes {
//line /usr/local/go/src/encoding/csv/reader.go:446
		_go_fuzz_dep_.CoverTab[114904]++
								dst[i] = str[preIdx:idx]
								preIdx = idx
//line /usr/local/go/src/encoding/csv/reader.go:448
		// _ = "end of CoverTab[114904]"
	}
//line /usr/local/go/src/encoding/csv/reader.go:449
	// _ = "end of CoverTab[114836]"
//line /usr/local/go/src/encoding/csv/reader.go:449
	_go_fuzz_dep_.CoverTab[114837]++

//line /usr/local/go/src/encoding/csv/reader.go:452
	if r.FieldsPerRecord > 0 {
//line /usr/local/go/src/encoding/csv/reader.go:452
		_go_fuzz_dep_.CoverTab[114905]++
								if len(dst) != r.FieldsPerRecord && func() bool {
//line /usr/local/go/src/encoding/csv/reader.go:453
			_go_fuzz_dep_.CoverTab[114906]++
//line /usr/local/go/src/encoding/csv/reader.go:453
			return err == nil
//line /usr/local/go/src/encoding/csv/reader.go:453
			// _ = "end of CoverTab[114906]"
//line /usr/local/go/src/encoding/csv/reader.go:453
		}() {
//line /usr/local/go/src/encoding/csv/reader.go:453
			_go_fuzz_dep_.CoverTab[114907]++
									err = &ParseError{
				StartLine:	recLine,
				Line:		recLine,
				Column:		1,
				Err:		ErrFieldCount,
			}
//line /usr/local/go/src/encoding/csv/reader.go:459
			// _ = "end of CoverTab[114907]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:460
			_go_fuzz_dep_.CoverTab[114908]++
//line /usr/local/go/src/encoding/csv/reader.go:460
			// _ = "end of CoverTab[114908]"
//line /usr/local/go/src/encoding/csv/reader.go:460
		}
//line /usr/local/go/src/encoding/csv/reader.go:460
		// _ = "end of CoverTab[114905]"
	} else {
//line /usr/local/go/src/encoding/csv/reader.go:461
		_go_fuzz_dep_.CoverTab[114909]++
//line /usr/local/go/src/encoding/csv/reader.go:461
		if r.FieldsPerRecord == 0 {
//line /usr/local/go/src/encoding/csv/reader.go:461
			_go_fuzz_dep_.CoverTab[114910]++
									r.FieldsPerRecord = len(dst)
//line /usr/local/go/src/encoding/csv/reader.go:462
			// _ = "end of CoverTab[114910]"
		} else {
//line /usr/local/go/src/encoding/csv/reader.go:463
			_go_fuzz_dep_.CoverTab[114911]++
//line /usr/local/go/src/encoding/csv/reader.go:463
			// _ = "end of CoverTab[114911]"
//line /usr/local/go/src/encoding/csv/reader.go:463
		}
//line /usr/local/go/src/encoding/csv/reader.go:463
		// _ = "end of CoverTab[114909]"
//line /usr/local/go/src/encoding/csv/reader.go:463
	}
//line /usr/local/go/src/encoding/csv/reader.go:463
	// _ = "end of CoverTab[114837]"
//line /usr/local/go/src/encoding/csv/reader.go:463
	_go_fuzz_dep_.CoverTab[114838]++
							return dst, err
//line /usr/local/go/src/encoding/csv/reader.go:464
	// _ = "end of CoverTab[114838]"
}

//line /usr/local/go/src/encoding/csv/reader.go:465
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/csv/reader.go:465
var _ = _go_fuzz_dep_.CoverTab
