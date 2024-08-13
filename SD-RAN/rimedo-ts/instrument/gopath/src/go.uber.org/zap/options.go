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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:21
)

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// An Option configures a Logger.
type Option interface {
	apply(*Logger)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:37
	_go_fuzz_dep_.CoverTab[131756]++
										f(log)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:38
	// _ = "end of CoverTab[131756]"
}

// WrapCore wraps or replaces the Logger's underlying zapcore.Core.
func WrapCore(f func(zapcore.Core) zapcore.Core) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:42
	_go_fuzz_dep_.CoverTab[131757]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:43
		_go_fuzz_dep_.CoverTab[131758]++
											log.core = f(log.core)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:44
		// _ = "end of CoverTab[131758]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:45
	// _ = "end of CoverTab[131757]"
}

// Hooks registers functions which will be called each time the Logger writes
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:48
// out an Entry. Repeated use of Hooks is additive.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:48
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:48
// Hooks are useful for simple side effects, like capturing metrics for the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:48
// number of emitted logs. More complex side effects, including anything that
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:48
// requires access to the Entry's structured fields, should be implemented as
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:48
// a zapcore.Core instead. See zapcore.RegisterHooks for details.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:55
func Hooks(hooks ...func(zapcore.Entry) error) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:55
	_go_fuzz_dep_.CoverTab[131759]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:56
		_go_fuzz_dep_.CoverTab[131760]++
											log.core = zapcore.RegisterHooks(log.core, hooks...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:57
		// _ = "end of CoverTab[131760]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:58
	// _ = "end of CoverTab[131759]"
}

// Fields adds fields to the Logger.
func Fields(fs ...Field) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:62
	_go_fuzz_dep_.CoverTab[131761]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:63
		_go_fuzz_dep_.CoverTab[131762]++
											log.core = log.core.With(fs)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:64
		// _ = "end of CoverTab[131762]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:65
	// _ = "end of CoverTab[131761]"
}

// ErrorOutput sets the destination for errors generated by the Logger. Note
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:68
// that this option only affects internal errors; for sample code that sends
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:68
// error-level logs to a different location from info- and debug-level logs,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:68
// see the package-level AdvancedConfiguration example.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:68
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:68
// The supplied WriteSyncer must be safe for concurrent use. The Open and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:68
// zapcore.Lock functions are the simplest ways to protect files with a mutex.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:75
func ErrorOutput(w zapcore.WriteSyncer) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:75
	_go_fuzz_dep_.CoverTab[131763]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:76
		_go_fuzz_dep_.CoverTab[131764]++
											log.errorOutput = w
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:77
		// _ = "end of CoverTab[131764]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:78
	// _ = "end of CoverTab[131763]"
}

// Development puts the logger in development mode, which makes DPanic-level
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:81
// logs panic instead of simply logging an error.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:83
func Development() Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:83
	_go_fuzz_dep_.CoverTab[131765]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:84
		_go_fuzz_dep_.CoverTab[131766]++
											log.development = true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:85
		// _ = "end of CoverTab[131766]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:86
	// _ = "end of CoverTab[131765]"
}

// AddCaller configures the Logger to annotate each message with the filename,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:89
// line number, and function name of zap's caller. See also WithCaller.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:91
func AddCaller() Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:91
	_go_fuzz_dep_.CoverTab[131767]++
										return WithCaller(true)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:92
	// _ = "end of CoverTab[131767]"
}

// WithCaller configures the Logger to annotate each message with the filename,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:95
// line number, and function name of zap's caller, or not, depending on the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:95
// value of enabled. This is a generalized form of AddCaller.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:98
func WithCaller(enabled bool) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:98
	_go_fuzz_dep_.CoverTab[131768]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:99
		_go_fuzz_dep_.CoverTab[131769]++
											log.addCaller = enabled
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:100
		// _ = "end of CoverTab[131769]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:101
	// _ = "end of CoverTab[131768]"
}

// AddCallerSkip increases the number of callers skipped by caller annotation
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:104
// (as enabled by the AddCaller option). When building wrappers around the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:104
// Logger and SugaredLogger, supplying this Option prevents zap from always
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:104
// reporting the wrapper code as the caller.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:108
func AddCallerSkip(skip int) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:108
	_go_fuzz_dep_.CoverTab[131770]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:109
		_go_fuzz_dep_.CoverTab[131771]++
											log.callerSkip += skip
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:110
		// _ = "end of CoverTab[131771]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:111
	// _ = "end of CoverTab[131770]"
}

// AddStacktrace configures the Logger to record a stack trace for all messages at
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:114
// or above a given level.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:116
func AddStacktrace(lvl zapcore.LevelEnabler) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:116
	_go_fuzz_dep_.CoverTab[131772]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:117
		_go_fuzz_dep_.CoverTab[131773]++
											log.addStack = lvl
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:118
		// _ = "end of CoverTab[131773]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:119
	// _ = "end of CoverTab[131772]"
}

// IncreaseLevel increase the level of the logger. It has no effect if
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:122
// the passed in level tries to decrease the level of the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:124
func IncreaseLevel(lvl zapcore.LevelEnabler) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:124
	_go_fuzz_dep_.CoverTab[131774]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:125
		_go_fuzz_dep_.CoverTab[131775]++
											core, err := zapcore.NewIncreaseLevelCore(log.core, lvl)
											if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:127
			_go_fuzz_dep_.CoverTab[131776]++
												fmt.Fprintf(log.errorOutput, "failed to IncreaseLevel: %v\n", err)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:128
			// _ = "end of CoverTab[131776]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:129
			_go_fuzz_dep_.CoverTab[131777]++
												log.core = core
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:130
			// _ = "end of CoverTab[131777]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:131
		// _ = "end of CoverTab[131775]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:132
	// _ = "end of CoverTab[131774]"
}

// OnFatal sets the action to take on fatal logs.
func OnFatal(action zapcore.CheckWriteAction) Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:136
	_go_fuzz_dep_.CoverTab[131778]++
										return optionFunc(func(log *Logger) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:137
		_go_fuzz_dep_.CoverTab[131779]++
											log.onFatal = action
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:138
		// _ = "end of CoverTab[131779]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:139
	// _ = "end of CoverTab[131778]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:140
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/options.go:140
var _ = _go_fuzz_dep_.CoverTab
