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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:21
)

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap/zapcore"
)

const (
	_loggerWriterDepth		= 2
	_programmerErrorTemplate	= "You've found a bug in zap! Please file a bug at " +
		"https://github.com/uber-go/zap/issues/new and reference this error: %v"
)

var (
	_globalMu	sync.RWMutex
	_globalL	= NewNop()
	_globalS	= _globalL.Sugar()
)

// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:45
// It's safe for concurrent use.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:47
func L() *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:47
	_go_fuzz_dep_.CoverTab[131607]++
									_globalMu.RLock()
									l := _globalL
									_globalMu.RUnlock()
									return l
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:51
	// _ = "end of CoverTab[131607]"
}

// S returns the global SugaredLogger, which can be reconfigured with
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:54
// ReplaceGlobals. It's safe for concurrent use.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:56
func S() *SugaredLogger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:56
	_go_fuzz_dep_.CoverTab[131608]++
									_globalMu.RLock()
									s := _globalS
									_globalMu.RUnlock()
									return s
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:60
	// _ = "end of CoverTab[131608]"
}

// ReplaceGlobals replaces the global Logger and SugaredLogger, and returns a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:63
// function to restore the original values. It's safe for concurrent use.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:65
func ReplaceGlobals(logger *Logger) func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:65
	_go_fuzz_dep_.CoverTab[131609]++
									_globalMu.Lock()
									prev := _globalL
									_globalL = logger
									_globalS = logger.Sugar()
									_globalMu.Unlock()
									return func() { _go_fuzz_dep_.CoverTab[131610]++; ReplaceGlobals(prev); // _ = "end of CoverTab[131610]" }
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:71
	// _ = "end of CoverTab[131609]"
}

// NewStdLog returns a *log.Logger which writes to the supplied zap Logger at
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:74
// InfoLevel. To redirect the standard library's package-global logging
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:74
// functions, use RedirectStdLog instead.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:77
func NewStdLog(l *Logger) *log.Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:77
	_go_fuzz_dep_.CoverTab[131611]++
									logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
									f := logger.Info
									return log.New(&loggerWriter{f}, "", 0)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:80
	// _ = "end of CoverTab[131611]"
}

// NewStdLogAt returns *log.Logger which writes to supplied zap logger at
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:83
// required level.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:85
func NewStdLogAt(l *Logger, level zapcore.Level) (*log.Logger, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:85
	_go_fuzz_dep_.CoverTab[131612]++
									logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
									logFunc, err := levelToFunc(logger, level)
									if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:88
		_go_fuzz_dep_.CoverTab[131614]++
										return nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:89
		// _ = "end of CoverTab[131614]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:90
		_go_fuzz_dep_.CoverTab[131615]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:90
		// _ = "end of CoverTab[131615]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:90
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:90
	// _ = "end of CoverTab[131612]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:90
	_go_fuzz_dep_.CoverTab[131613]++
									return log.New(&loggerWriter{logFunc}, "", 0), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:91
	// _ = "end of CoverTab[131613]"
}

// RedirectStdLog redirects output from the standard library's package-global
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:94
// logger to the supplied logger at InfoLevel. Since zap already handles caller
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:94
// annotations, timestamps, etc., it automatically disables the standard
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:94
// library's annotations and prefixing.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:94
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:94
// It returns a function to restore the original prefix and flags and reset the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:94
// standard library's output to os.Stderr.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:101
func RedirectStdLog(l *Logger) func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:101
	_go_fuzz_dep_.CoverTab[131616]++
										f, err := redirectStdLogAt(l, InfoLevel)
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:103
		_go_fuzz_dep_.CoverTab[131618]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:106
		panic(fmt.Sprintf(_programmerErrorTemplate, err))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:106
		// _ = "end of CoverTab[131618]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:107
		_go_fuzz_dep_.CoverTab[131619]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:107
		// _ = "end of CoverTab[131619]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:107
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:107
	// _ = "end of CoverTab[131616]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:107
	_go_fuzz_dep_.CoverTab[131617]++
										return f
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:108
	// _ = "end of CoverTab[131617]"
}

// RedirectStdLogAt redirects output from the standard library's package-global
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:111
// logger to the supplied logger at the specified level. Since zap already
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:111
// handles caller annotations, timestamps, etc., it automatically disables the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:111
// standard library's annotations and prefixing.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:111
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:111
// It returns a function to restore the original prefix and flags and reset the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:111
// standard library's output to os.Stderr.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:118
func RedirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:118
	_go_fuzz_dep_.CoverTab[131620]++
										return redirectStdLogAt(l, level)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:119
	// _ = "end of CoverTab[131620]"
}

func redirectStdLogAt(l *Logger, level zapcore.Level) (func(), error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:122
	_go_fuzz_dep_.CoverTab[131621]++
										flags := log.Flags()
										prefix := log.Prefix()
										log.SetFlags(0)
										log.SetPrefix("")
										logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
										logFunc, err := levelToFunc(logger, level)
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:129
		_go_fuzz_dep_.CoverTab[131623]++
											return nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:130
		// _ = "end of CoverTab[131623]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:131
		_go_fuzz_dep_.CoverTab[131624]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:131
		// _ = "end of CoverTab[131624]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:131
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:131
	// _ = "end of CoverTab[131621]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:131
	_go_fuzz_dep_.CoverTab[131622]++
										log.SetOutput(&loggerWriter{logFunc})
										return func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:133
		_go_fuzz_dep_.CoverTab[131625]++
											log.SetFlags(flags)
											log.SetPrefix(prefix)
											log.SetOutput(os.Stderr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:136
		// _ = "end of CoverTab[131625]"
	}, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:137
	// _ = "end of CoverTab[131622]"
}

func levelToFunc(logger *Logger, lvl zapcore.Level) (func(string, ...Field), error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:140
	_go_fuzz_dep_.CoverTab[131626]++
										switch lvl {
	case DebugLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:142
		_go_fuzz_dep_.CoverTab[131628]++
											return logger.Debug, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:143
		// _ = "end of CoverTab[131628]"
	case InfoLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:144
		_go_fuzz_dep_.CoverTab[131629]++
											return logger.Info, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:145
		// _ = "end of CoverTab[131629]"
	case WarnLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:146
		_go_fuzz_dep_.CoverTab[131630]++
											return logger.Warn, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:147
		// _ = "end of CoverTab[131630]"
	case ErrorLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:148
		_go_fuzz_dep_.CoverTab[131631]++
											return logger.Error, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:149
		// _ = "end of CoverTab[131631]"
	case DPanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:150
		_go_fuzz_dep_.CoverTab[131632]++
											return logger.DPanic, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:151
		// _ = "end of CoverTab[131632]"
	case PanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:152
		_go_fuzz_dep_.CoverTab[131633]++
											return logger.Panic, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:153
		// _ = "end of CoverTab[131633]"
	case FatalLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:154
		_go_fuzz_dep_.CoverTab[131634]++
											return logger.Fatal, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:155
		// _ = "end of CoverTab[131634]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:155
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:155
		_go_fuzz_dep_.CoverTab[131635]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:155
		// _ = "end of CoverTab[131635]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:156
	// _ = "end of CoverTab[131626]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:156
	_go_fuzz_dep_.CoverTab[131627]++
										return nil, fmt.Errorf("unrecognized level: %q", lvl)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:157
	// _ = "end of CoverTab[131627]"
}

type loggerWriter struct {
	logFunc func(msg string, fields ...Field)
}

func (l *loggerWriter) Write(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:164
	_go_fuzz_dep_.CoverTab[131636]++
										p = bytes.TrimSpace(p)
										l.logFunc(string(p))
										return len(p), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:167
	// _ = "end of CoverTab[131636]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:168
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/global.go:168
var _ = _go_fuzz_dep_.CoverTab
