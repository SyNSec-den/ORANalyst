// Copyright (c) 2016 Uber Technologies, Inc.
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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:21
)

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap/internal/bufferpool"
	"go.uber.org/zap/internal/exit"

	"go.uber.org/multierr"
)

var (
	_cePool = sync.Pool{New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:37
		_go_fuzz_dep_.CoverTab[130754]++

											return &CheckedEntry{
			cores: make([]Core, 4),
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:41
		// _ = "end of CoverTab[130754]"
	}}
)

func getCheckedEntry() *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:45
	_go_fuzz_dep_.CoverTab[130755]++
										ce := _cePool.Get().(*CheckedEntry)
										ce.reset()
										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:48
	// _ = "end of CoverTab[130755]"
}

func putCheckedEntry(ce *CheckedEntry) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:51
	_go_fuzz_dep_.CoverTab[130756]++
										if ce == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:52
		_go_fuzz_dep_.CoverTab[130758]++
											return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:53
		// _ = "end of CoverTab[130758]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:54
		_go_fuzz_dep_.CoverTab[130759]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:54
		// _ = "end of CoverTab[130759]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:54
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:54
	// _ = "end of CoverTab[130756]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:54
	_go_fuzz_dep_.CoverTab[130757]++
										_cePool.Put(ce)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:55
	// _ = "end of CoverTab[130757]"
}

// NewEntryCaller makes an EntryCaller from the return signature of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:58
// runtime.Caller.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:60
func NewEntryCaller(pc uintptr, file string, line int, ok bool) EntryCaller {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:60
	_go_fuzz_dep_.CoverTab[130760]++
										if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:61
		_go_fuzz_dep_.CoverTab[130762]++
											return EntryCaller{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:62
		// _ = "end of CoverTab[130762]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:63
		_go_fuzz_dep_.CoverTab[130763]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:63
		// _ = "end of CoverTab[130763]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:63
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:63
	// _ = "end of CoverTab[130760]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:63
	_go_fuzz_dep_.CoverTab[130761]++
										return EntryCaller{
		PC:		pc,
		File:		file,
		Line:		line,
		Defined:	true,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:69
	// _ = "end of CoverTab[130761]"
}

// EntryCaller represents the caller of a logging function.
type EntryCaller struct {
	Defined		bool
	PC		uintptr
	File		string
	Line		int
	Function	string
}

// String returns the full path and line number of the caller.
func (ec EntryCaller) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:82
	_go_fuzz_dep_.CoverTab[130764]++
										return ec.FullPath()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:83
	// _ = "end of CoverTab[130764]"
}

// FullPath returns a /full/path/to/package/file:line description of the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:86
// caller.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:88
func (ec EntryCaller) FullPath() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:88
	_go_fuzz_dep_.CoverTab[130765]++
										if !ec.Defined {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:89
		_go_fuzz_dep_.CoverTab[130767]++
											return "undefined"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:90
		// _ = "end of CoverTab[130767]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:91
		_go_fuzz_dep_.CoverTab[130768]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:91
		// _ = "end of CoverTab[130768]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:91
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:91
	// _ = "end of CoverTab[130765]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:91
	_go_fuzz_dep_.CoverTab[130766]++
										buf := bufferpool.Get()
										buf.AppendString(ec.File)
										buf.AppendByte(':')
										buf.AppendInt(int64(ec.Line))
										caller := buf.String()
										buf.Free()
										return caller
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:98
	// _ = "end of CoverTab[130766]"
}

// TrimmedPath returns a package/file:line description of the caller,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:101
// preserving only the leaf directory name and file name.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:103
func (ec EntryCaller) TrimmedPath() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:103
	_go_fuzz_dep_.CoverTab[130769]++
										if !ec.Defined {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:104
		_go_fuzz_dep_.CoverTab[130773]++
											return "undefined"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:105
		// _ = "end of CoverTab[130773]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:106
		_go_fuzz_dep_.CoverTab[130774]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:106
		// _ = "end of CoverTab[130774]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:106
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:106
	// _ = "end of CoverTab[130769]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:106
	_go_fuzz_dep_.CoverTab[130770]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:120
	idx := strings.LastIndexByte(ec.File, '/')
	if idx == -1 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:121
		_go_fuzz_dep_.CoverTab[130775]++
											return ec.FullPath()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:122
		// _ = "end of CoverTab[130775]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:123
		_go_fuzz_dep_.CoverTab[130776]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:123
		// _ = "end of CoverTab[130776]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:123
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:123
	// _ = "end of CoverTab[130770]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:123
	_go_fuzz_dep_.CoverTab[130771]++

										idx = strings.LastIndexByte(ec.File[:idx], '/')
										if idx == -1 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:126
		_go_fuzz_dep_.CoverTab[130777]++
											return ec.FullPath()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:127
		// _ = "end of CoverTab[130777]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:128
		_go_fuzz_dep_.CoverTab[130778]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:128
		// _ = "end of CoverTab[130778]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:128
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:128
	// _ = "end of CoverTab[130771]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:128
	_go_fuzz_dep_.CoverTab[130772]++
										buf := bufferpool.Get()

										buf.AppendString(ec.File[idx+1:])
										buf.AppendByte(':')
										buf.AppendInt(int64(ec.Line))
										caller := buf.String()
										buf.Free()
										return caller
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:136
	// _ = "end of CoverTab[130772]"
}

// An Entry represents a complete log message. The entry's structured context
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:139
// is already serialized, but the log level, time, message, and call site
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:139
// information are available for inspection and modification. Any fields left
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:139
// empty will be omitted when encoding.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:139
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:139
// Entries are pooled, so any functions that accept them MUST be careful not to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:139
// retain references to them.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:146
type Entry struct {
	Level		Level
	Time		time.Time
	LoggerName	string
	Message		string
	Caller		EntryCaller
	Stack		string
}

// CheckWriteAction indicates what action to take after a log entry is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:155
// processed. Actions are ordered in increasing severity.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:157
type CheckWriteAction uint8

const (
	// WriteThenNoop indicates that nothing special needs to be done. It's the
	// default behavior.
	WriteThenNoop	CheckWriteAction	= iota
	// WriteThenGoexit runs runtime.Goexit after Write.
	WriteThenGoexit
	// WriteThenPanic causes a panic after Write.
	WriteThenPanic
	// WriteThenFatal causes a fatal os.Exit after Write.
	WriteThenFatal
)

// CheckedEntry is an Entry together with a collection of Cores that have
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:171
// already agreed to log it.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:171
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:171
// CheckedEntry references should be created by calling AddCore or Should on a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:171
// nil *CheckedEntry. References are returned to a pool after Write, and MUST
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:171
// NOT be retained after calling their Write method.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:177
type CheckedEntry struct {
	Entry
	ErrorOutput	WriteSyncer
	dirty		bool	// best-effort detection of pool misuse
	should		CheckWriteAction
	cores		[]Core
}

func (ce *CheckedEntry) reset() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:185
	_go_fuzz_dep_.CoverTab[130779]++
										ce.Entry = Entry{}
										ce.ErrorOutput = nil
										ce.dirty = false
										ce.should = WriteThenNoop
										for i := range ce.cores {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:190
		_go_fuzz_dep_.CoverTab[130781]++

											ce.cores[i] = nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:192
		// _ = "end of CoverTab[130781]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:193
	// _ = "end of CoverTab[130779]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:193
	_go_fuzz_dep_.CoverTab[130780]++
										ce.cores = ce.cores[:0]
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:194
	// _ = "end of CoverTab[130780]"
}

// Write writes the entry to the stored Cores, returns any errors, and returns
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:197
// the CheckedEntry reference to a pool for immediate re-use. Finally, it
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:197
// executes any required CheckWriteAction.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:200
func (ce *CheckedEntry) Write(fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:200
	_go_fuzz_dep_.CoverTab[130782]++
										if ce == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:201
		_go_fuzz_dep_.CoverTab[130787]++
											return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:202
		// _ = "end of CoverTab[130787]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:203
		_go_fuzz_dep_.CoverTab[130788]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:203
		// _ = "end of CoverTab[130788]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:203
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:203
	// _ = "end of CoverTab[130782]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:203
	_go_fuzz_dep_.CoverTab[130783]++

										if ce.dirty {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:205
		_go_fuzz_dep_.CoverTab[130789]++
											if ce.ErrorOutput != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:206
			_go_fuzz_dep_.CoverTab[130791]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:211
			fmt.Fprintf(ce.ErrorOutput, "%v Unsafe CheckedEntry re-use near Entry %+v.\n", time.Now(), ce.Entry)
												ce.ErrorOutput.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:212
			// _ = "end of CoverTab[130791]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:213
			_go_fuzz_dep_.CoverTab[130792]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:213
			// _ = "end of CoverTab[130792]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:213
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:213
		// _ = "end of CoverTab[130789]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:213
		_go_fuzz_dep_.CoverTab[130790]++
											return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:214
		// _ = "end of CoverTab[130790]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:215
		_go_fuzz_dep_.CoverTab[130793]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:215
		// _ = "end of CoverTab[130793]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:215
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:215
	// _ = "end of CoverTab[130783]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:215
	_go_fuzz_dep_.CoverTab[130784]++
										ce.dirty = true

										var err error
										for i := range ce.cores {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:219
		_go_fuzz_dep_.CoverTab[130794]++
											err = multierr.Append(err, ce.cores[i].Write(ce.Entry, fields))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:220
		// _ = "end of CoverTab[130794]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:221
	// _ = "end of CoverTab[130784]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:221
	_go_fuzz_dep_.CoverTab[130785]++
										if ce.ErrorOutput != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:222
		_go_fuzz_dep_.CoverTab[130795]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:223
			_go_fuzz_dep_.CoverTab[130796]++
												fmt.Fprintf(ce.ErrorOutput, "%v write error: %v\n", time.Now(), err)
												ce.ErrorOutput.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:225
			// _ = "end of CoverTab[130796]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:226
			_go_fuzz_dep_.CoverTab[130797]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:226
			// _ = "end of CoverTab[130797]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:226
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:226
		// _ = "end of CoverTab[130795]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:227
		_go_fuzz_dep_.CoverTab[130798]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:227
		// _ = "end of CoverTab[130798]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:227
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:227
	// _ = "end of CoverTab[130785]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:227
	_go_fuzz_dep_.CoverTab[130786]++

										should, msg := ce.should, ce.Message
										putCheckedEntry(ce)

										switch should {
	case WriteThenPanic:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:233
		_go_fuzz_dep_.CoverTab[130799]++
											panic(msg)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:234
		// _ = "end of CoverTab[130799]"
	case WriteThenFatal:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:235
		_go_fuzz_dep_.CoverTab[130800]++
											exit.Exit()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:236
		// _ = "end of CoverTab[130800]"
	case WriteThenGoexit:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:237
		_go_fuzz_dep_.CoverTab[130801]++
											runtime.Goexit()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:238
		// _ = "end of CoverTab[130801]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:238
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:238
		_go_fuzz_dep_.CoverTab[130802]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:238
		// _ = "end of CoverTab[130802]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:239
	// _ = "end of CoverTab[130786]"
}

// AddCore adds a Core that has agreed to log this CheckedEntry. It's intended to be
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:242
// used by Core.Check implementations, and is safe to call on nil CheckedEntry
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:242
// references.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:245
func (ce *CheckedEntry) AddCore(ent Entry, core Core) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:245
	_go_fuzz_dep_.CoverTab[130803]++
										if ce == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:246
		_go_fuzz_dep_.CoverTab[130805]++
											ce = getCheckedEntry()
											ce.Entry = ent
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:248
		// _ = "end of CoverTab[130805]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:249
		_go_fuzz_dep_.CoverTab[130806]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:249
		// _ = "end of CoverTab[130806]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:249
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:249
	// _ = "end of CoverTab[130803]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:249
	_go_fuzz_dep_.CoverTab[130804]++
										ce.cores = append(ce.cores, core)
										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:251
	// _ = "end of CoverTab[130804]"
}

// Should sets this CheckedEntry's CheckWriteAction, which controls whether a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:254
// Core will panic or fatal after writing this log entry. Like AddCore, it's
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:254
// safe to call on nil CheckedEntry references.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:257
func (ce *CheckedEntry) Should(ent Entry, should CheckWriteAction) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:257
	_go_fuzz_dep_.CoverTab[130807]++
										if ce == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:258
		_go_fuzz_dep_.CoverTab[130809]++
											ce = getCheckedEntry()
											ce.Entry = ent
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:260
		// _ = "end of CoverTab[130809]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:261
		_go_fuzz_dep_.CoverTab[130810]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:261
		// _ = "end of CoverTab[130810]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:261
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:261
	// _ = "end of CoverTab[130807]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:261
	_go_fuzz_dep_.CoverTab[130808]++
										ce.should = should
										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:263
	// _ = "end of CoverTab[130808]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:264
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/entry.go:264
var _ = _go_fuzz_dep_.CoverTab
