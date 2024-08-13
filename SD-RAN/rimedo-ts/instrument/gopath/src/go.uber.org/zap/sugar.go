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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:21
)

import (
	"fmt"

	"go.uber.org/zap/zapcore"

	"go.uber.org/multierr"
)

const (
	_oddNumberErrMsg	= "Ignored key without a value."
	_nonStringKeyErrMsg	= "Ignored key-value pairs with non-string keys."
)

// A SugaredLogger wraps the base Logger functionality in a slower, but less
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// verbose, API. Any Logger can be converted to a SugaredLogger with its Sugar
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// method.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// Unlike the Logger, the SugaredLogger doesn't insist on structured logging.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// For each log level, it exposes three methods: one for loosely-typed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// structured logging, one for println-style formatting, and one for
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// printf-style formatting. For example, SugaredLoggers can produce InfoLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:36
// output with Infow ("info with" structured context), Info, or Infof.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:45
type SugaredLogger struct {
	base *Logger
}

// Desugar unwraps a SugaredLogger, exposing the original Logger. Desugaring
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:49
// is quite inexpensive, so it's reasonable for a single application to use
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:49
// both Loggers and SugaredLoggers, converting between them on the boundaries
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:49
// of performance-sensitive code.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:53
func (s *SugaredLogger) Desugar() *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:53
	_go_fuzz_dep_.CoverTab[131853]++
									base := s.base.clone()
									base.callerSkip -= 2
									return base
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:56
	// _ = "end of CoverTab[131853]"
}

// Named adds a sub-scope to the logger's name. See Logger.Named for details.
func (s *SugaredLogger) Named(name string) *SugaredLogger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:60
	_go_fuzz_dep_.CoverTab[131854]++
									return &SugaredLogger{base: s.base.Named(name)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:61
	// _ = "end of CoverTab[131854]"
}

// With adds a variadic number of fields to the logging context. It accepts a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// processing pairs, the first element of the pair is used as the field key
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// and the second as the field value.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// For example,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	 sugaredLogger.With(
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	   "hello", "world",
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	   "failure", errors.New("oh no"),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	   Stack(),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	   "count", 42,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	   "user", User{Name: "alice"},
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// is the equivalent of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	unsugared.With(
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	  String("hello", "world"),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	  String("failure", "oh no"),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	  Stack(),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	  Int("count", 42),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	  Object("user", User{Name: "alice"}),
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//	)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// Note that the keys in key-value pairs should be strings. In development,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// passing a non-string key panics. In production, the logger is more
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// forgiving: a separate error is logged, but the key-value pair is skipped
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// and execution continues. Passing an orphaned key triggers similar behavior:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:64
// panics in development and errors in production.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:91
func (s *SugaredLogger) With(args ...interface{}) *SugaredLogger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:91
	_go_fuzz_dep_.CoverTab[131855]++
									return &SugaredLogger{base: s.base.With(s.sweetenFields(args)...)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:92
	// _ = "end of CoverTab[131855]"
}

// Debug uses fmt.Sprint to construct and log a message.
func (s *SugaredLogger) Debug(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:96
	_go_fuzz_dep_.CoverTab[131856]++
									s.log(DebugLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:97
	// _ = "end of CoverTab[131856]"
}

// Info uses fmt.Sprint to construct and log a message.
func (s *SugaredLogger) Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:101
	_go_fuzz_dep_.CoverTab[131857]++
									s.log(InfoLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:102
	// _ = "end of CoverTab[131857]"
}

// Warn uses fmt.Sprint to construct and log a message.
func (s *SugaredLogger) Warn(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:106
	_go_fuzz_dep_.CoverTab[131858]++
									s.log(WarnLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:107
	// _ = "end of CoverTab[131858]"
}

// Error uses fmt.Sprint to construct and log a message.
func (s *SugaredLogger) Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:111
	_go_fuzz_dep_.CoverTab[131859]++
									s.log(ErrorLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:112
	// _ = "end of CoverTab[131859]"
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:115
// logger then panics. (See DPanicLevel for details.)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:117
func (s *SugaredLogger) DPanic(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:117
	_go_fuzz_dep_.CoverTab[131860]++
									s.log(DPanicLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:118
	// _ = "end of CoverTab[131860]"
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (s *SugaredLogger) Panic(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:122
	_go_fuzz_dep_.CoverTab[131861]++
									s.log(PanicLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:123
	// _ = "end of CoverTab[131861]"
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (s *SugaredLogger) Fatal(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:127
	_go_fuzz_dep_.CoverTab[131862]++
									s.log(FatalLevel, "", args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:128
	// _ = "end of CoverTab[131862]"
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) Debugf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:132
	_go_fuzz_dep_.CoverTab[131863]++
									s.log(DebugLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:133
	// _ = "end of CoverTab[131863]"
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) Infof(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:137
	_go_fuzz_dep_.CoverTab[131864]++
									s.log(InfoLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:138
	// _ = "end of CoverTab[131864]"
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) Warnf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:142
	_go_fuzz_dep_.CoverTab[131865]++
									s.log(WarnLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:143
	// _ = "end of CoverTab[131865]"
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) Errorf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:147
	_go_fuzz_dep_.CoverTab[131866]++
									s.log(ErrorLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:148
	// _ = "end of CoverTab[131866]"
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:151
// logger then panics. (See DPanicLevel for details.)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:153
func (s *SugaredLogger) DPanicf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:153
	_go_fuzz_dep_.CoverTab[131867]++
									s.log(DPanicLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:154
	// _ = "end of CoverTab[131867]"
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *SugaredLogger) Panicf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:158
	_go_fuzz_dep_.CoverTab[131868]++
									s.log(PanicLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:159
	// _ = "end of CoverTab[131868]"
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *SugaredLogger) Fatalf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:163
	_go_fuzz_dep_.CoverTab[131869]++
									s.log(FatalLevel, template, args, nil)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:164
	// _ = "end of CoverTab[131869]"
}

// Debugw logs a message with some additional context. The variadic key-value
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:167
// pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:167
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:167
// When debug-level logging is disabled, this is much faster than
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:167
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:167
//	s.With(keysAndValues).Debug(msg)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:172
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:172
	_go_fuzz_dep_.CoverTab[131870]++
									s.log(DebugLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:173
	// _ = "end of CoverTab[131870]"
}

// Infow logs a message with some additional context. The variadic key-value
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:176
// pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:178
func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:178
	_go_fuzz_dep_.CoverTab[131871]++
									s.log(InfoLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:179
	// _ = "end of CoverTab[131871]"
}

// Warnw logs a message with some additional context. The variadic key-value
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:182
// pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:184
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:184
	_go_fuzz_dep_.CoverTab[131872]++
									s.log(WarnLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:185
	// _ = "end of CoverTab[131872]"
}

// Errorw logs a message with some additional context. The variadic key-value
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:188
// pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:190
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:190
	_go_fuzz_dep_.CoverTab[131873]++
									s.log(ErrorLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:191
	// _ = "end of CoverTab[131873]"
}

// DPanicw logs a message with some additional context. In development, the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:194
// logger then panics. (See DPanicLevel for details.) The variadic key-value
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:194
// pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:197
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:197
	_go_fuzz_dep_.CoverTab[131874]++
									s.log(DPanicLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:198
	// _ = "end of CoverTab[131874]"
}

// Panicw logs a message with some additional context, then panics. The
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:201
// variadic key-value pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:203
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:203
	_go_fuzz_dep_.CoverTab[131875]++
									s.log(PanicLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:204
	// _ = "end of CoverTab[131875]"
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:207
// variadic key-value pairs are treated as they are in With.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:209
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:209
	_go_fuzz_dep_.CoverTab[131876]++
									s.log(FatalLevel, msg, nil, keysAndValues)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:210
	// _ = "end of CoverTab[131876]"
}

// Sync flushes any buffered log entries.
func (s *SugaredLogger) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:214
	_go_fuzz_dep_.CoverTab[131877]++
									return s.base.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:215
	// _ = "end of CoverTab[131877]"
}

func (s *SugaredLogger) log(lvl zapcore.Level, template string, fmtArgs []interface{}, context []interface{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:218
	_go_fuzz_dep_.CoverTab[131878]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:221
	if lvl < DPanicLevel && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:221
		_go_fuzz_dep_.CoverTab[131880]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:221
		return !s.base.Core().Enabled(lvl)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:221
		// _ = "end of CoverTab[131880]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:221
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:221
		_go_fuzz_dep_.CoverTab[131881]++
										return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:222
		// _ = "end of CoverTab[131881]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:223
		_go_fuzz_dep_.CoverTab[131882]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:223
		// _ = "end of CoverTab[131882]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:223
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:223
	// _ = "end of CoverTab[131878]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:223
	_go_fuzz_dep_.CoverTab[131879]++

									msg := getMessage(template, fmtArgs)
									if ce := s.base.Check(lvl, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:226
		_go_fuzz_dep_.CoverTab[131883]++
										ce.Write(s.sweetenFields(context)...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:227
		// _ = "end of CoverTab[131883]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:228
		_go_fuzz_dep_.CoverTab[131884]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:228
		// _ = "end of CoverTab[131884]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:228
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:228
	// _ = "end of CoverTab[131879]"
}

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:232
	_go_fuzz_dep_.CoverTab[131885]++
									if len(fmtArgs) == 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:233
		_go_fuzz_dep_.CoverTab[131889]++
										return template
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:234
		// _ = "end of CoverTab[131889]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:235
		_go_fuzz_dep_.CoverTab[131890]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:235
		// _ = "end of CoverTab[131890]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:235
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:235
	// _ = "end of CoverTab[131885]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:235
	_go_fuzz_dep_.CoverTab[131886]++

									if template != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:237
		_go_fuzz_dep_.CoverTab[131891]++
										return fmt.Sprintf(template, fmtArgs...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:238
		// _ = "end of CoverTab[131891]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:239
		_go_fuzz_dep_.CoverTab[131892]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:239
		// _ = "end of CoverTab[131892]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:239
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:239
	// _ = "end of CoverTab[131886]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:239
	_go_fuzz_dep_.CoverTab[131887]++

									if len(fmtArgs) == 1 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:241
		_go_fuzz_dep_.CoverTab[131893]++
										if str, ok := fmtArgs[0].(string); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:242
			_go_fuzz_dep_.CoverTab[131894]++
											return str
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:243
			// _ = "end of CoverTab[131894]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:244
			_go_fuzz_dep_.CoverTab[131895]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:244
			// _ = "end of CoverTab[131895]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:244
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:244
		// _ = "end of CoverTab[131893]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:245
		_go_fuzz_dep_.CoverTab[131896]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:245
		// _ = "end of CoverTab[131896]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:245
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:245
	// _ = "end of CoverTab[131887]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:245
	_go_fuzz_dep_.CoverTab[131888]++
									return fmt.Sprint(fmtArgs...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:246
	// _ = "end of CoverTab[131888]"
}

func (s *SugaredLogger) sweetenFields(args []interface{}) []Field {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:249
	_go_fuzz_dep_.CoverTab[131897]++
									if len(args) == 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:250
		_go_fuzz_dep_.CoverTab[131901]++
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:251
		// _ = "end of CoverTab[131901]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:252
		_go_fuzz_dep_.CoverTab[131902]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:252
		// _ = "end of CoverTab[131902]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:252
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:252
	// _ = "end of CoverTab[131897]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:252
	_go_fuzz_dep_.CoverTab[131898]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:256
	fields := make([]Field, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:259
		_go_fuzz_dep_.CoverTab[131903]++

										if f, ok := args[i].(Field); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:261
			_go_fuzz_dep_.CoverTab[131907]++
											fields = append(fields, f)
											i++
											continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:264
			// _ = "end of CoverTab[131907]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:265
			_go_fuzz_dep_.CoverTab[131908]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:265
			// _ = "end of CoverTab[131908]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:265
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:265
		// _ = "end of CoverTab[131903]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:265
		_go_fuzz_dep_.CoverTab[131904]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:268
		if i == len(args)-1 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:268
			_go_fuzz_dep_.CoverTab[131909]++
											s.base.DPanic(_oddNumberErrMsg, Any("ignored", args[i]))
											break
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:270
			// _ = "end of CoverTab[131909]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:271
			_go_fuzz_dep_.CoverTab[131910]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:271
			// _ = "end of CoverTab[131910]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:271
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:271
		// _ = "end of CoverTab[131904]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:271
		_go_fuzz_dep_.CoverTab[131905]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:275
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:276
			_go_fuzz_dep_.CoverTab[131911]++

											if cap(invalid) == 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:278
				_go_fuzz_dep_.CoverTab[131913]++
												invalid = make(invalidPairs, 0, len(args)/2)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:279
				// _ = "end of CoverTab[131913]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:280
				_go_fuzz_dep_.CoverTab[131914]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:280
				// _ = "end of CoverTab[131914]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:280
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:280
			// _ = "end of CoverTab[131911]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:280
			_go_fuzz_dep_.CoverTab[131912]++
											invalid = append(invalid, invalidPair{i, key, val})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:281
			// _ = "end of CoverTab[131912]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:282
			_go_fuzz_dep_.CoverTab[131915]++
											fields = append(fields, Any(keyStr, val))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:283
			// _ = "end of CoverTab[131915]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:284
		// _ = "end of CoverTab[131905]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:284
		_go_fuzz_dep_.CoverTab[131906]++
										i += 2
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:285
		// _ = "end of CoverTab[131906]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:286
	// _ = "end of CoverTab[131898]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:286
	_go_fuzz_dep_.CoverTab[131899]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:289
	if len(invalid) > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:289
		_go_fuzz_dep_.CoverTab[131916]++
										s.base.DPanic(_nonStringKeyErrMsg, Array("invalid", invalid))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:290
		// _ = "end of CoverTab[131916]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:291
		_go_fuzz_dep_.CoverTab[131917]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:291
		// _ = "end of CoverTab[131917]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:291
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:291
	// _ = "end of CoverTab[131899]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:291
	_go_fuzz_dep_.CoverTab[131900]++
									return fields
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:292
	// _ = "end of CoverTab[131900]"
}

type invalidPair struct {
	position	int
	key, value	interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:300
	_go_fuzz_dep_.CoverTab[131918]++
									enc.AddInt64("position", int64(p.position))
									Any("key", p.key).AddTo(enc)
									Any("value", p.value).AddTo(enc)
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:304
	// _ = "end of CoverTab[131918]"
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:309
	_go_fuzz_dep_.CoverTab[131919]++
									var err error
									for i := range ps {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:311
		_go_fuzz_dep_.CoverTab[131921]++
										err = multierr.Append(err, enc.AppendObject(ps[i]))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:312
		// _ = "end of CoverTab[131921]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:313
	// _ = "end of CoverTab[131919]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:313
	_go_fuzz_dep_.CoverTab[131920]++
									return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:314
	// _ = "end of CoverTab[131920]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:315
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/sugar.go:315
var _ = _go_fuzz_dep_.CoverTab
