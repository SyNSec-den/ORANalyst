// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// Package multierr allows combining one or more errors together.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// # Overview
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// Errors can be combined with the use of the Combine function.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	multierr.Combine(
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		reader.Close(),
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		writer.Close(),
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		conn.Close(),
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// If only two errors are being combined, the Append function may be used
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// instead.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	err = multierr.Append(reader.Close(), writer.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// This makes it possible to record resource cleanup failures from deferred
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// blocks with the help of named return values.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	func sendRequest(req Request) (err error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		conn, err := openConnection()
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//			return err
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		defer func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//			err = multierr.Append(err, conn.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		}()
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		// ...
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// The underlying list of errors for a returned error object may be retrieved
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// with the Errors function.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	errors := multierr.Errors(err)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	if len(errors) > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		fmt.Println("The following errors occurred:", errors)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// # Advanced Usage
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// Errors returned by Combine and Append MAY implement the following
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// interface.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	type errorGroup interface {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		// Returns a slice containing the underlying list of errors.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		// This slice MUST NOT be modified by the caller.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		Errors() []error
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// Note that if you need access to list of errors behind a multierr error, you
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// should prefer using the Errors function. That said, if you need cheap
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// read-only access to the underlying errors slice, you can attempt to cast
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// the error to this interface. You MUST handle the failure case gracefully
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// because errors returned by Combine and Append are not guaranteed to
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
// implement this interface.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	var errors []error
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	group, ok := err.(errorGroup)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	if ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		errors = group.Errors()
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//		errors = []error{err}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:21
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
package multierr

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:86
)

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"

	"go.uber.org/atomic"
)

var (
	// Separator for single-line error messages.
	_singlelineSeparator	= []byte("; ")

	// Prefix for multi-line messages
	_multilinePrefix	= []byte("the following errors occurred:")

	// Prefix for the first and following lines of an item in a list of
	// multi-line error messages.
	//
	// For example, if a single item is:
	//
	// 	foo
	// 	bar
	//
	// It will become,
	//
	// 	 -  foo
	// 	    bar
	_multilineSeparator	= []byte("\n -  ")
	_multilineIndent	= []byte("    ")
)

// _bufferPool is a pool of bytes.Buffers.
var _bufferPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:123
		_go_fuzz_dep_.CoverTab[130485]++
											return &bytes.Buffer{}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:124
		// _ = "end of CoverTab[130485]"
	},
}

type errorGroup interface {
	Errors() []error
}

// Errors returns a slice containing zero or more errors that the supplied
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
// error is composed of. If the error is nil, a nil slice is returned.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
//	err := multierr.Append(r.Close(), w.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
//	errors := multierr.Errors(err)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
// If the error is not composed of other errors, the returned slice contains
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
// just the error that was passed in.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:132
// Callers of this function are free to modify the returned slice.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:142
func Errors(err error) []error {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:142
	_go_fuzz_dep_.CoverTab[130486]++
										if err == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:143
		_go_fuzz_dep_.CoverTab[130489]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:144
		// _ = "end of CoverTab[130489]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:145
		_go_fuzz_dep_.CoverTab[130490]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:145
		// _ = "end of CoverTab[130490]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:145
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:145
	// _ = "end of CoverTab[130486]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:145
	_go_fuzz_dep_.CoverTab[130487]++

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:154
	eg, ok := err.(*multiError)
	if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:155
		_go_fuzz_dep_.CoverTab[130491]++
											return []error{err}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:156
		// _ = "end of CoverTab[130491]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:157
		_go_fuzz_dep_.CoverTab[130492]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:157
		// _ = "end of CoverTab[130492]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:157
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:157
	// _ = "end of CoverTab[130487]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:157
	_go_fuzz_dep_.CoverTab[130488]++

										errors := eg.Errors()
										result := make([]error, len(errors))
										copy(result, errors)
										return result
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:162
	// _ = "end of CoverTab[130488]"
}

// multiError is an error that holds one or more errors.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:165
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:165
// An instance of this is guaranteed to be non-empty and flattened. That is,
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:165
// none of the errors inside multiError are other multiErrors.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:165
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:165
// multiError formats to a semi-colon delimited list of error messages with
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:165
// %v and with a more readable multi-line format with %+v.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:172
type multiError struct {
	copyNeeded	atomic.Bool
	errors		[]error
}

var _ errorGroup = (*multiError)(nil)

// Errors returns the list of underlying errors.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:179
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:179
// This slice MUST NOT be modified.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:182
func (merr *multiError) Errors() []error {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:182
	_go_fuzz_dep_.CoverTab[130493]++
										if merr == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:183
		_go_fuzz_dep_.CoverTab[130495]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:184
		// _ = "end of CoverTab[130495]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:185
		_go_fuzz_dep_.CoverTab[130496]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:185
		// _ = "end of CoverTab[130496]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:185
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:185
	// _ = "end of CoverTab[130493]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:185
	_go_fuzz_dep_.CoverTab[130494]++
										return merr.errors
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:186
	// _ = "end of CoverTab[130494]"
}

func (merr *multiError) Error() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:189
	_go_fuzz_dep_.CoverTab[130497]++
										if merr == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:190
		_go_fuzz_dep_.CoverTab[130499]++
											return ""
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:191
		// _ = "end of CoverTab[130499]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:192
		_go_fuzz_dep_.CoverTab[130500]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:192
		// _ = "end of CoverTab[130500]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:192
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:192
	// _ = "end of CoverTab[130497]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:192
	_go_fuzz_dep_.CoverTab[130498]++

										buff := _bufferPool.Get().(*bytes.Buffer)
										buff.Reset()

										merr.writeSingleline(buff)

										result := buff.String()
										_bufferPool.Put(buff)
										return result
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:201
	// _ = "end of CoverTab[130498]"
}

func (merr *multiError) Format(f fmt.State, c rune) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:204
	_go_fuzz_dep_.CoverTab[130501]++
										if c == 'v' && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:205
		_go_fuzz_dep_.CoverTab[130502]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:205
		return f.Flag('+')
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:205
		// _ = "end of CoverTab[130502]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:205
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:205
		_go_fuzz_dep_.CoverTab[130503]++
											merr.writeMultiline(f)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:206
		// _ = "end of CoverTab[130503]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:207
		_go_fuzz_dep_.CoverTab[130504]++
											merr.writeSingleline(f)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:208
		// _ = "end of CoverTab[130504]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:209
	// _ = "end of CoverTab[130501]"
}

func (merr *multiError) writeSingleline(w io.Writer) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:212
	_go_fuzz_dep_.CoverTab[130505]++
										first := true
										for _, item := range merr.errors {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:214
		_go_fuzz_dep_.CoverTab[130506]++
											if first {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:215
			_go_fuzz_dep_.CoverTab[130508]++
												first = false
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:216
			// _ = "end of CoverTab[130508]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:217
			_go_fuzz_dep_.CoverTab[130509]++
												w.Write(_singlelineSeparator)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:218
			// _ = "end of CoverTab[130509]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:219
		// _ = "end of CoverTab[130506]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:219
		_go_fuzz_dep_.CoverTab[130507]++
											io.WriteString(w, item.Error())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:220
		// _ = "end of CoverTab[130507]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:221
	// _ = "end of CoverTab[130505]"
}

func (merr *multiError) writeMultiline(w io.Writer) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:224
	_go_fuzz_dep_.CoverTab[130510]++
										w.Write(_multilinePrefix)
										for _, item := range merr.errors {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:226
		_go_fuzz_dep_.CoverTab[130511]++
											w.Write(_multilineSeparator)
											writePrefixLine(w, _multilineIndent, fmt.Sprintf("%+v", item))
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:228
		// _ = "end of CoverTab[130511]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:229
	// _ = "end of CoverTab[130510]"
}

// Writes s to the writer with the given prefix added before each line after
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:232
// the first.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:234
func writePrefixLine(w io.Writer, prefix []byte, s string) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:234
	_go_fuzz_dep_.CoverTab[130512]++
										first := true
										for len(s) > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:236
		_go_fuzz_dep_.CoverTab[130513]++
											if first {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:237
			_go_fuzz_dep_.CoverTab[130516]++
												first = false
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:238
			// _ = "end of CoverTab[130516]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:239
			_go_fuzz_dep_.CoverTab[130517]++
												w.Write(prefix)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:240
			// _ = "end of CoverTab[130517]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:241
		// _ = "end of CoverTab[130513]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:241
		_go_fuzz_dep_.CoverTab[130514]++

											idx := strings.IndexByte(s, '\n')
											if idx < 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:244
			_go_fuzz_dep_.CoverTab[130518]++
												idx = len(s) - 1
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:245
			// _ = "end of CoverTab[130518]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:246
			_go_fuzz_dep_.CoverTab[130519]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:246
			// _ = "end of CoverTab[130519]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:246
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:246
		// _ = "end of CoverTab[130514]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:246
		_go_fuzz_dep_.CoverTab[130515]++

											io.WriteString(w, s[:idx+1])
											s = s[idx+1:]
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:249
		// _ = "end of CoverTab[130515]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:250
	// _ = "end of CoverTab[130512]"
}

type inspectResult struct {
	// Number of top-level non-nil errors
	Count	int

	// Total number of errors including multiErrors
	Capacity	int

	// Index of the first non-nil error in the list. Value is meaningless if
	// Count is zero.
	FirstErrorIdx	int

	// Whether the list contains at least one multiError
	ContainsMultiError	bool
}

// Inspects the given slice of errors so that we can efficiently allocate
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:268
// space for it.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:270
func inspect(errors []error) (res inspectResult) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:270
	_go_fuzz_dep_.CoverTab[130520]++
										first := true
										for i, err := range errors {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:272
		_go_fuzz_dep_.CoverTab[130522]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:273
			_go_fuzz_dep_.CoverTab[130525]++
												continue
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:274
			// _ = "end of CoverTab[130525]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:275
			_go_fuzz_dep_.CoverTab[130526]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:275
			// _ = "end of CoverTab[130526]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:275
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:275
		// _ = "end of CoverTab[130522]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:275
		_go_fuzz_dep_.CoverTab[130523]++

											res.Count++
											if first {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:278
			_go_fuzz_dep_.CoverTab[130527]++
												first = false
												res.FirstErrorIdx = i
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:280
			// _ = "end of CoverTab[130527]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:281
			_go_fuzz_dep_.CoverTab[130528]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:281
			// _ = "end of CoverTab[130528]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:281
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:281
		// _ = "end of CoverTab[130523]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:281
		_go_fuzz_dep_.CoverTab[130524]++

											if merr, ok := err.(*multiError); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:283
			_go_fuzz_dep_.CoverTab[130529]++
												res.Capacity += len(merr.errors)
												res.ContainsMultiError = true
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:285
			// _ = "end of CoverTab[130529]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:286
			_go_fuzz_dep_.CoverTab[130530]++
												res.Capacity++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:287
			// _ = "end of CoverTab[130530]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:288
		// _ = "end of CoverTab[130524]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:289
	// _ = "end of CoverTab[130520]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:289
	_go_fuzz_dep_.CoverTab[130521]++
										return
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:290
	// _ = "end of CoverTab[130521]"
}

// fromSlice converts the given list of errors into a single error.
func fromSlice(errors []error) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:294
	_go_fuzz_dep_.CoverTab[130531]++
										res := inspect(errors)
										switch res.Count {
	case 0:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:297
		_go_fuzz_dep_.CoverTab[130534]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:298
		// _ = "end of CoverTab[130534]"
	case 1:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:299
		_go_fuzz_dep_.CoverTab[130535]++

											return errors[res.FirstErrorIdx]
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:301
		// _ = "end of CoverTab[130535]"
	case len(errors):
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:302
		_go_fuzz_dep_.CoverTab[130536]++
											if !res.ContainsMultiError {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:303
			_go_fuzz_dep_.CoverTab[130538]++

												return &multiError{errors: errors}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:305
			// _ = "end of CoverTab[130538]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
			_go_fuzz_dep_.CoverTab[130539]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
			// _ = "end of CoverTab[130539]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
		// _ = "end of CoverTab[130536]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
		_go_fuzz_dep_.CoverTab[130537]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:306
		// _ = "end of CoverTab[130537]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:307
	// _ = "end of CoverTab[130531]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:307
	_go_fuzz_dep_.CoverTab[130532]++

										nonNilErrs := make([]error, 0, res.Capacity)
										for _, err := range errors[res.FirstErrorIdx:] {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:310
		_go_fuzz_dep_.CoverTab[130540]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:311
			_go_fuzz_dep_.CoverTab[130542]++
												continue
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:312
			// _ = "end of CoverTab[130542]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:313
			_go_fuzz_dep_.CoverTab[130543]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:313
			// _ = "end of CoverTab[130543]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:313
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:313
		// _ = "end of CoverTab[130540]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:313
		_go_fuzz_dep_.CoverTab[130541]++

											if nested, ok := err.(*multiError); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:315
			_go_fuzz_dep_.CoverTab[130544]++
												nonNilErrs = append(nonNilErrs, nested.errors...)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:316
			// _ = "end of CoverTab[130544]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:317
			_go_fuzz_dep_.CoverTab[130545]++
												nonNilErrs = append(nonNilErrs, err)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:318
			// _ = "end of CoverTab[130545]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:319
		// _ = "end of CoverTab[130541]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:320
	// _ = "end of CoverTab[130532]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:320
	_go_fuzz_dep_.CoverTab[130533]++

										return &multiError{errors: nonNilErrs}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:322
	// _ = "end of CoverTab[130533]"
}

// Combine combines the passed errors into a single error.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// If zero arguments were passed or if all items are nil, a nil error is
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// returned.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	Combine(nil, nil)  // == nil
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// If only a single error was passed, it is returned as-is.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	Combine(err)  // == err
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// Combine skips over nil arguments so this function may be used to combine
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// together errors from operations that fail independently of each other.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	multierr.Combine(
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//		reader.Close(),
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//		writer.Close(),
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//		pipe.Close(),
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// If any of the passed errors is a multierr error, it will be flattened along
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// with the other errors.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	multierr.Combine(multierr.Combine(err1, err2), err3)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	// is the same as
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	multierr.Combine(err1, err2, err3)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// The returned error formats into a readable multi-line error message if
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
// formatted with %+v.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:325
//	fmt.Sprintf("%+v", multierr.Combine(err1, err2))
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:356
func Combine(errors ...error) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:356
	_go_fuzz_dep_.CoverTab[130546]++
										return fromSlice(errors)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:357
	// _ = "end of CoverTab[130546]"
}

// Append appends the given errors together. Either value may be nil.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
// This function is a specialization of Combine for the common case where
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
// there are only two errors.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//	err = multierr.Append(reader.Close(), writer.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
// The following pattern may also be used to record failure of deferred
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
// operations without losing information about the original error.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//	func doSomething(..) (err error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//		f := acquireResource()
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//		defer func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//			err = multierr.Append(err, f.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:360
//		}()
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:375
func Append(left error, right error) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:375
	_go_fuzz_dep_.CoverTab[130547]++
										switch {
	case left == nil:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:377
		_go_fuzz_dep_.CoverTab[130550]++
											return right
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:378
		// _ = "end of CoverTab[130550]"
	case right == nil:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:379
		_go_fuzz_dep_.CoverTab[130551]++
											return left
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:380
		// _ = "end of CoverTab[130551]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:380
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:380
		_go_fuzz_dep_.CoverTab[130552]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:380
		// _ = "end of CoverTab[130552]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:381
	// _ = "end of CoverTab[130547]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:381
	_go_fuzz_dep_.CoverTab[130548]++

										if _, ok := right.(*multiError); !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:383
		_go_fuzz_dep_.CoverTab[130553]++
											if l, ok := left.(*multiError); ok && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:384
			_go_fuzz_dep_.CoverTab[130554]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:384
			return !l.copyNeeded.Swap(true)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:384
			// _ = "end of CoverTab[130554]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:384
		}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:384
			_go_fuzz_dep_.CoverTab[130555]++

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:387
			errs := append(l.errors, right)
												return &multiError{errors: errs}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:388
			// _ = "end of CoverTab[130555]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:389
			_go_fuzz_dep_.CoverTab[130556]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:389
			if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:389
				_go_fuzz_dep_.CoverTab[130557]++

													return &multiError{errors: []error{left, right}}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:391
				// _ = "end of CoverTab[130557]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:392
				_go_fuzz_dep_.CoverTab[130558]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:392
				// _ = "end of CoverTab[130558]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:392
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:392
			// _ = "end of CoverTab[130556]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:392
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:392
		// _ = "end of CoverTab[130553]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:393
		_go_fuzz_dep_.CoverTab[130559]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:393
		// _ = "end of CoverTab[130559]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:393
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:393
	// _ = "end of CoverTab[130548]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:393
	_go_fuzz_dep_.CoverTab[130549]++

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:397
	errors := [2]error{left, right}
										return fromSlice(errors[0:])
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:398
	// _ = "end of CoverTab[130549]"
}

// AppendInto appends an error into the destination of an error pointer and
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
// returns whether the error being appended was non-nil.
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	var err error
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	multierr.AppendInto(&err, r.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	multierr.AppendInto(&err, w.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
// The above is equivalent to,
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	err := multierr.Append(r.Close(), w.Close())
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
// As AppendInto reports whether the provided error was non-nil, it may be
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
// used to build a multierr error in a loop more ergonomically. For example:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	var err error
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	for line := range lines {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		var item Item
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		if multierr.AppendInto(&err, parse(line, &item)) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//			continue
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		items = append(items, item)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
// Compare this with a verison that relies solely on Append:
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	var err error
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	for line := range lines {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		var item Item
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		if parseErr := parse(line, &item); parseErr != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//			err = multierr.Append(err, parseErr)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//			continue
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//		items = append(items, item)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:401
//	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:435
func AppendInto(into *error, err error) (errored bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:435
	_go_fuzz_dep_.CoverTab[130560]++
										if into == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:436
		_go_fuzz_dep_.CoverTab[130563]++

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:441
		panic("misuse of multierr.AppendInto: into pointer must not be nil")
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:441
		// _ = "end of CoverTab[130563]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:442
		_go_fuzz_dep_.CoverTab[130564]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:442
		// _ = "end of CoverTab[130564]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:442
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:442
	// _ = "end of CoverTab[130560]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:442
	_go_fuzz_dep_.CoverTab[130561]++

										if err == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:444
		_go_fuzz_dep_.CoverTab[130565]++
											return false
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:445
		// _ = "end of CoverTab[130565]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:446
		_go_fuzz_dep_.CoverTab[130566]++
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:446
		// _ = "end of CoverTab[130566]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:446
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:446
	// _ = "end of CoverTab[130561]"
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:446
	_go_fuzz_dep_.CoverTab[130562]++
										*into = Append(*into, err)
										return true
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:448
	// _ = "end of CoverTab[130562]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:449
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/multierr@v1.6.0/error.go:449
var _ = _go_fuzz_dep_.CoverTab
