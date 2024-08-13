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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:21
)

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

// A Logger provides fast, leveled, structured logging. All methods are safe
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:34
// for concurrent use.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:34
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:34
// The Logger is designed for contexts in which every microsecond and every
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:34
// allocation matters, so its API intentionally favors performance and type
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:34
// safety over brevity. For most applications, the SugaredLogger strikes a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:34
// better balance between performance and ergonomics.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:41
type Logger struct {
	core	zapcore.Core

	development	bool
	addCaller	bool
	onFatal		zapcore.CheckWriteAction	// default is WriteThenFatal

	name		string
	errorOutput	zapcore.WriteSyncer

	addStack	zapcore.LevelEnabler

	callerSkip	int
}

// New constructs a new Logger from the provided zapcore.Core and Options. If
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// the passed zapcore.Core is nil, it falls back to using a no-op
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// implementation.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// This is the most flexible way to construct a Logger, but also the most
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// verbose. For typical use cases, the highly-opinionated presets
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// (NewProduction, NewDevelopment, and NewExample) or the Config struct are
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// more convenient.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:56
// For sample code, see the package-level AdvancedConfiguration example.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:66
func New(core zapcore.Core, options ...Option) *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:66
	_go_fuzz_dep_.CoverTab[131677]++
									if core == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:67
		_go_fuzz_dep_.CoverTab[131679]++
										return NewNop()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:68
		// _ = "end of CoverTab[131679]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:69
		_go_fuzz_dep_.CoverTab[131680]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:69
		// _ = "end of CoverTab[131680]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:69
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:69
	// _ = "end of CoverTab[131677]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:69
	_go_fuzz_dep_.CoverTab[131678]++
									log := &Logger{
		core:		core,
		errorOutput:	zapcore.Lock(os.Stderr),
		addStack:	zapcore.FatalLevel + 1,
	}
									return log.WithOptions(options...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:75
	// _ = "end of CoverTab[131678]"
}

// NewNop returns a no-op Logger. It never writes out logs or internal errors,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:78
// and it never runs user-defined hooks.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:78
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:78
// Using WithOptions to replace the Core or error output of a no-op Logger can
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:78
// re-enable logging.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:83
func NewNop() *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:83
	_go_fuzz_dep_.CoverTab[131681]++
									return &Logger{
		core:		zapcore.NewNopCore(),
		errorOutput:	zapcore.AddSync(ioutil.Discard),
		addStack:	zapcore.FatalLevel + 1,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:88
	// _ = "end of CoverTab[131681]"
}

// NewProduction builds a sensible production Logger that writes InfoLevel and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:91
// above logs to standard error as JSON.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:91
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:91
// It's a shortcut for NewProductionConfig().Build(...Option).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:95
func NewProduction(options ...Option) (*Logger, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:95
	_go_fuzz_dep_.CoverTab[131682]++
									return NewProductionConfig().Build(options...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:96
	// _ = "end of CoverTab[131682]"
}

// NewDevelopment builds a development Logger that writes DebugLevel and above
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:99
// logs to standard error in a human-friendly format.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:99
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:99
// It's a shortcut for NewDevelopmentConfig().Build(...Option).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:103
func NewDevelopment(options ...Option) (*Logger, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:103
	_go_fuzz_dep_.CoverTab[131683]++
										return NewDevelopmentConfig().Build(options...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:104
	// _ = "end of CoverTab[131683]"
}

// NewExample builds a Logger that's designed for use in zap's testable
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:107
// examples. It writes DebugLevel and above logs to standard out as JSON, but
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:107
// omits the timestamp and calling function to keep example output
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:107
// short and deterministic.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:111
func NewExample(options ...Option) *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:111
	_go_fuzz_dep_.CoverTab[131684]++
										encoderCfg := zapcore.EncoderConfig{
		MessageKey:	"msg",
		LevelKey:	"level",
		NameKey:	"logger",
		EncodeLevel:	zapcore.LowercaseLevelEncoder,
		EncodeTime:	zapcore.ISO8601TimeEncoder,
		EncodeDuration:	zapcore.StringDurationEncoder,
	}
										core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, DebugLevel)
										return New(core).WithOptions(options...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:121
	// _ = "end of CoverTab[131684]"
}

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:124
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:124
// single application to use both Loggers and SugaredLoggers, converting
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:124
// between them on the boundaries of performance-sensitive code.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:128
func (log *Logger) Sugar() *SugaredLogger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:128
	_go_fuzz_dep_.CoverTab[131685]++
										core := log.clone()
										core.callerSkip += 2
										return &SugaredLogger{core}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:131
	// _ = "end of CoverTab[131685]"
}

// Named adds a new path segment to the logger's name. Segments are joined by
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:134
// periods. By default, Loggers are unnamed.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:136
func (log *Logger) Named(s string) *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:136
	_go_fuzz_dep_.CoverTab[131686]++
										if s == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:137
		_go_fuzz_dep_.CoverTab[131689]++
											return log
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:138
		// _ = "end of CoverTab[131689]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:139
		_go_fuzz_dep_.CoverTab[131690]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:139
		// _ = "end of CoverTab[131690]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:139
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:139
	// _ = "end of CoverTab[131686]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:139
	_go_fuzz_dep_.CoverTab[131687]++
										l := log.clone()
										if log.name == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:141
		_go_fuzz_dep_.CoverTab[131691]++
											l.name = s
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:142
		// _ = "end of CoverTab[131691]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:143
		_go_fuzz_dep_.CoverTab[131692]++
											l.name = strings.Join([]string{l.name, s}, ".")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:144
		// _ = "end of CoverTab[131692]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:145
	// _ = "end of CoverTab[131687]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:145
	_go_fuzz_dep_.CoverTab[131688]++
										return l
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:146
	// _ = "end of CoverTab[131688]"
}

// WithOptions clones the current Logger, applies the supplied Options, and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:149
// returns the resulting Logger. It's safe to use concurrently.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:151
func (log *Logger) WithOptions(opts ...Option) *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:151
	_go_fuzz_dep_.CoverTab[131693]++
										c := log.clone()
										for _, opt := range opts {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:153
		_go_fuzz_dep_.CoverTab[131695]++
											opt.apply(c)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:154
		// _ = "end of CoverTab[131695]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:155
	// _ = "end of CoverTab[131693]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:155
	_go_fuzz_dep_.CoverTab[131694]++
										return c
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:156
	// _ = "end of CoverTab[131694]"
}

// With creates a child logger and adds structured context to it. Fields added
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:159
// to the child don't affect the parent, and vice versa.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:161
func (log *Logger) With(fields ...Field) *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:161
	_go_fuzz_dep_.CoverTab[131696]++
										if len(fields) == 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:162
		_go_fuzz_dep_.CoverTab[131698]++
											return log
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:163
		// _ = "end of CoverTab[131698]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:164
		_go_fuzz_dep_.CoverTab[131699]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:164
		// _ = "end of CoverTab[131699]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:164
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:164
	// _ = "end of CoverTab[131696]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:164
	_go_fuzz_dep_.CoverTab[131697]++
										l := log.clone()
										l.core = l.core.With(fields)
										return l
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:167
	// _ = "end of CoverTab[131697]"
}

// Check returns a CheckedEntry if logging a message at the specified level
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:170
// is enabled. It's a completely optional optimization; in high-performance
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:170
// applications, Check can help avoid allocating a slice to hold fields.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:173
func (log *Logger) Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:173
	_go_fuzz_dep_.CoverTab[131700]++
										return log.check(lvl, msg)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:174
	// _ = "end of CoverTab[131700]"
}

// Debug logs a message at DebugLevel. The message includes any fields passed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:177
// at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:179
func (log *Logger) Debug(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:179
	_go_fuzz_dep_.CoverTab[131701]++
										if ce := log.check(DebugLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:180
		_go_fuzz_dep_.CoverTab[131702]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:181
		// _ = "end of CoverTab[131702]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:182
		_go_fuzz_dep_.CoverTab[131703]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:182
		// _ = "end of CoverTab[131703]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:182
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:182
	// _ = "end of CoverTab[131701]"
}

// Info logs a message at InfoLevel. The message includes any fields passed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:185
// at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:187
func (log *Logger) Info(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:187
	_go_fuzz_dep_.CoverTab[131704]++
										if ce := log.check(InfoLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:188
		_go_fuzz_dep_.CoverTab[131705]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:189
		// _ = "end of CoverTab[131705]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:190
		_go_fuzz_dep_.CoverTab[131706]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:190
		// _ = "end of CoverTab[131706]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:190
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:190
	// _ = "end of CoverTab[131704]"
}

// Warn logs a message at WarnLevel. The message includes any fields passed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:193
// at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:195
func (log *Logger) Warn(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:195
	_go_fuzz_dep_.CoverTab[131707]++
										if ce := log.check(WarnLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:196
		_go_fuzz_dep_.CoverTab[131708]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:197
		// _ = "end of CoverTab[131708]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:198
		_go_fuzz_dep_.CoverTab[131709]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:198
		// _ = "end of CoverTab[131709]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:198
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:198
	// _ = "end of CoverTab[131707]"
}

// Error logs a message at ErrorLevel. The message includes any fields passed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:201
// at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:203
func (log *Logger) Error(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:203
	_go_fuzz_dep_.CoverTab[131710]++
										if ce := log.check(ErrorLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:204
		_go_fuzz_dep_.CoverTab[131711]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:205
		// _ = "end of CoverTab[131711]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:206
		_go_fuzz_dep_.CoverTab[131712]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:206
		// _ = "end of CoverTab[131712]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:206
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:206
	// _ = "end of CoverTab[131710]"
}

// DPanic logs a message at DPanicLevel. The message includes any fields
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:209
// passed at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:209
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:209
// If the logger is in development mode, it then panics (DPanic means
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:209
// "development panic"). This is useful for catching errors that are
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:209
// recoverable, but shouldn't ever happen.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:215
func (log *Logger) DPanic(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:215
	_go_fuzz_dep_.CoverTab[131713]++
										if ce := log.check(DPanicLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:216
		_go_fuzz_dep_.CoverTab[131714]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:217
		// _ = "end of CoverTab[131714]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:218
		_go_fuzz_dep_.CoverTab[131715]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:218
		// _ = "end of CoverTab[131715]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:218
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:218
	// _ = "end of CoverTab[131713]"
}

// Panic logs a message at PanicLevel. The message includes any fields passed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:221
// at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:221
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:221
// The logger then panics, even if logging at PanicLevel is disabled.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:225
func (log *Logger) Panic(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:225
	_go_fuzz_dep_.CoverTab[131716]++
										if ce := log.check(PanicLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:226
		_go_fuzz_dep_.CoverTab[131717]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:227
		// _ = "end of CoverTab[131717]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:228
		_go_fuzz_dep_.CoverTab[131718]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:228
		// _ = "end of CoverTab[131718]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:228
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:228
	// _ = "end of CoverTab[131716]"
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:231
// at the log site, as well as any fields accumulated on the logger.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:231
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:231
// The logger then calls os.Exit(1), even if logging at FatalLevel is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:231
// disabled.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:236
func (log *Logger) Fatal(msg string, fields ...Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:236
	_go_fuzz_dep_.CoverTab[131719]++
										if ce := log.check(FatalLevel, msg); ce != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:237
		_go_fuzz_dep_.CoverTab[131720]++
											ce.Write(fields...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:238
		// _ = "end of CoverTab[131720]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:239
		_go_fuzz_dep_.CoverTab[131721]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:239
		// _ = "end of CoverTab[131721]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:239
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:239
	// _ = "end of CoverTab[131719]"
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:242
// entries. Applications should take care to call Sync before exiting.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:244
func (log *Logger) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:244
	_go_fuzz_dep_.CoverTab[131722]++
										return log.core.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:245
	// _ = "end of CoverTab[131722]"
}

// Core returns the Logger's underlying zapcore.Core.
func (log *Logger) Core() zapcore.Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:249
	_go_fuzz_dep_.CoverTab[131723]++
										return log.core
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:250
	// _ = "end of CoverTab[131723]"
}

func (log *Logger) clone() *Logger {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:253
	_go_fuzz_dep_.CoverTab[131724]++
										copy := *log
										return &copy
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:255
	// _ = "end of CoverTab[131724]"
}

func (log *Logger) check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:258
	_go_fuzz_dep_.CoverTab[131725]++
	// check must always be called directly by a method in the Logger interface
										// (e.g., Check, Info, Fatal).
										const callerSkipOffset = 2

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:265
	if lvl < zapcore.DPanicLevel && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:265
		_go_fuzz_dep_.CoverTab[131731]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:265
		return !log.core.Enabled(lvl)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:265
		// _ = "end of CoverTab[131731]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:265
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:265
		_go_fuzz_dep_.CoverTab[131732]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:266
		// _ = "end of CoverTab[131732]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:267
		_go_fuzz_dep_.CoverTab[131733]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:267
		// _ = "end of CoverTab[131733]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:267
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:267
	// _ = "end of CoverTab[131725]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:267
	_go_fuzz_dep_.CoverTab[131726]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:271
	ent := zapcore.Entry{
		LoggerName:	log.name,
		Time:		time.Now(),
		Level:		lvl,
		Message:	msg,
	}
										ce := log.core.Check(ent, nil)
										willWrite := ce != nil

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:281
	switch ent.Level {
	case zapcore.PanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:282
		_go_fuzz_dep_.CoverTab[131734]++
											ce = ce.Should(ent, zapcore.WriteThenPanic)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:283
		// _ = "end of CoverTab[131734]"
	case zapcore.FatalLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:284
		_go_fuzz_dep_.CoverTab[131735]++
											onFatal := log.onFatal

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:288
		if onFatal == zapcore.WriteThenNoop {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:288
			_go_fuzz_dep_.CoverTab[131739]++
												onFatal = zapcore.WriteThenFatal
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:289
			// _ = "end of CoverTab[131739]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:290
			_go_fuzz_dep_.CoverTab[131740]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:290
			// _ = "end of CoverTab[131740]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:290
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:290
		// _ = "end of CoverTab[131735]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:290
		_go_fuzz_dep_.CoverTab[131736]++
											ce = ce.Should(ent, onFatal)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:291
		// _ = "end of CoverTab[131736]"
	case zapcore.DPanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:292
		_go_fuzz_dep_.CoverTab[131737]++
											if log.development {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:293
			_go_fuzz_dep_.CoverTab[131741]++
												ce = ce.Should(ent, zapcore.WriteThenPanic)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:294
			// _ = "end of CoverTab[131741]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
			_go_fuzz_dep_.CoverTab[131742]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
			// _ = "end of CoverTab[131742]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
		// _ = "end of CoverTab[131737]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
		_go_fuzz_dep_.CoverTab[131738]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:295
		// _ = "end of CoverTab[131738]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:296
	// _ = "end of CoverTab[131726]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:296
	_go_fuzz_dep_.CoverTab[131727]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:301
	if !willWrite {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:301
		_go_fuzz_dep_.CoverTab[131743]++
											return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:302
		// _ = "end of CoverTab[131743]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:303
		_go_fuzz_dep_.CoverTab[131744]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:303
		// _ = "end of CoverTab[131744]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:303
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:303
	// _ = "end of CoverTab[131727]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:303
	_go_fuzz_dep_.CoverTab[131728]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:306
	ce.ErrorOutput = log.errorOutput
	if log.addCaller {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:307
		_go_fuzz_dep_.CoverTab[131745]++
											frame, defined := getCallerFrame(log.callerSkip + callerSkipOffset)
											if !defined {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:309
			_go_fuzz_dep_.CoverTab[131747]++
												fmt.Fprintf(log.errorOutput, "%v Logger.check error: failed to get caller\n", time.Now().UTC())
												log.errorOutput.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:311
			// _ = "end of CoverTab[131747]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:312
			_go_fuzz_dep_.CoverTab[131748]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:312
			// _ = "end of CoverTab[131748]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:312
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:312
		// _ = "end of CoverTab[131745]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:312
		_go_fuzz_dep_.CoverTab[131746]++

											ce.Entry.Caller = zapcore.EntryCaller{
			Defined:	defined,
			PC:		frame.PC,
			File:		frame.File,
			Line:		frame.Line,
			Function:	frame.Function,
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:320
		// _ = "end of CoverTab[131746]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:321
		_go_fuzz_dep_.CoverTab[131749]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:321
		// _ = "end of CoverTab[131749]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:321
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:321
	// _ = "end of CoverTab[131728]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:321
	_go_fuzz_dep_.CoverTab[131729]++
										if log.addStack.Enabled(ce.Entry.Level) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:322
		_go_fuzz_dep_.CoverTab[131750]++
											ce.Entry.Stack = StackSkip("", log.callerSkip+callerSkipOffset).String
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:323
		// _ = "end of CoverTab[131750]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:324
		_go_fuzz_dep_.CoverTab[131751]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:324
		// _ = "end of CoverTab[131751]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:324
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:324
	// _ = "end of CoverTab[131729]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:324
	_go_fuzz_dep_.CoverTab[131730]++

										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:326
	// _ = "end of CoverTab[131730]"
}

// getCallerFrame gets caller frame. The argument skip is the number of stack
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:329
// frames to ascend, with 0 identifying the caller of getCallerFrame. The
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:329
// boolean ok is false if it was not possible to recover the information.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:329
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:329
// Note: This implementation is similar to runtime.Caller, but it returns the whole frame.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:334
func getCallerFrame(skip int) (frame runtime.Frame, ok bool) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:334
	_go_fuzz_dep_.CoverTab[131752]++
										const skipOffset = 2	// skip getCallerFrame and Callers

										pc := make([]uintptr, 1)
										numFrames := runtime.Callers(skip+skipOffset, pc)
										if numFrames < 1 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:339
		_go_fuzz_dep_.CoverTab[131754]++
											return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:340
		// _ = "end of CoverTab[131754]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:341
		_go_fuzz_dep_.CoverTab[131755]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:341
		// _ = "end of CoverTab[131755]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:341
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:341
	// _ = "end of CoverTab[131752]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:341
	_go_fuzz_dep_.CoverTab[131753]++

										frame, _ = runtime.CallersFrames(pc).Next()
										return frame, frame.PC != 0
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:344
	// _ = "end of CoverTab[131753]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:345
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/logger.go:345
var _ = _go_fuzz_dep_.CoverTab
