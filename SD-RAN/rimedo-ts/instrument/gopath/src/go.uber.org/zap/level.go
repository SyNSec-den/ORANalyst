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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:21
)

import (
	"go.uber.org/atomic"
	"go.uber.org/zap/zapcore"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel	= zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel	= zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel	= zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel	= zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel	= zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel	= zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel	= zapcore.FatalLevel
)

// LevelEnablerFunc is a convenient way to implement zapcore.LevelEnabler with
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:49
// an anonymous function.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:49
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:49
// It's particularly useful when splitting log output between different
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:49
// outputs (e.g., standard error and standard out). For sample code, see the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:49
// package-level AdvancedConfiguration example.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:55
type LevelEnablerFunc func(zapcore.Level) bool

// Enabled calls the wrapped function.
func (f LevelEnablerFunc) Enabled(lvl zapcore.Level) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:58
	_go_fuzz_dep_.CoverTab[131662]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:58
	return f(lvl)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:58
	// _ = "end of CoverTab[131662]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:58
}

// An AtomicLevel is an atomically changeable, dynamic logging level. It lets
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
// you safely change the log level of a tree of loggers (the root logger and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
// any children created by adding context) at runtime.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
// The AtomicLevel itself is an http.Handler that serves a JSON endpoint to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
// alter its level.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
// AtomicLevels must be created with the NewAtomicLevel constructor to allocate
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:60
// their internal atomic pointer.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:69
type AtomicLevel struct {
	l *atomic.Int32
}

// NewAtomicLevel creates an AtomicLevel with InfoLevel and above logging
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:73
// enabled.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:75
func NewAtomicLevel() AtomicLevel {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:75
	_go_fuzz_dep_.CoverTab[131663]++
									return AtomicLevel{
		l: atomic.NewInt32(int32(InfoLevel)),
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:78
	// _ = "end of CoverTab[131663]"
}

// NewAtomicLevelAt is a convenience function that creates an AtomicLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:81
// and then calls SetLevel with the given level.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:83
func NewAtomicLevelAt(l zapcore.Level) AtomicLevel {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:83
	_go_fuzz_dep_.CoverTab[131664]++
									a := NewAtomicLevel()
									a.SetLevel(l)
									return a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:86
	// _ = "end of CoverTab[131664]"
}

// Enabled implements the zapcore.LevelEnabler interface, which allows the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:89
// AtomicLevel to be used in place of traditional static levels.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:91
func (lvl AtomicLevel) Enabled(l zapcore.Level) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:91
	_go_fuzz_dep_.CoverTab[131665]++
									return lvl.Level().Enabled(l)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:92
	// _ = "end of CoverTab[131665]"
}

// Level returns the minimum enabled log level.
func (lvl AtomicLevel) Level() zapcore.Level {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:96
	_go_fuzz_dep_.CoverTab[131666]++
									return zapcore.Level(int8(lvl.l.Load()))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:97
	// _ = "end of CoverTab[131666]"
}

// SetLevel alters the logging level.
func (lvl AtomicLevel) SetLevel(l zapcore.Level) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:101
	_go_fuzz_dep_.CoverTab[131667]++
									lvl.l.Store(int32(l))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:102
	// _ = "end of CoverTab[131667]"
}

// String returns the string representation of the underlying Level.
func (lvl AtomicLevel) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:106
	_go_fuzz_dep_.CoverTab[131668]++
									return lvl.Level().String()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:107
	// _ = "end of CoverTab[131668]"
}

// UnmarshalText unmarshals the text to an AtomicLevel. It uses the same text
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:110
// representations as the static zapcore.Levels ("debug", "info", "warn",
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:110
// "error", "dpanic", "panic", and "fatal").
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:113
func (lvl *AtomicLevel) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:113
	_go_fuzz_dep_.CoverTab[131669]++
									if lvl.l == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:114
		_go_fuzz_dep_.CoverTab[131672]++
										lvl.l = &atomic.Int32{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:115
		// _ = "end of CoverTab[131672]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:116
		_go_fuzz_dep_.CoverTab[131673]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:116
		// _ = "end of CoverTab[131673]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:116
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:116
	// _ = "end of CoverTab[131669]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:116
	_go_fuzz_dep_.CoverTab[131670]++

									var l zapcore.Level
									if err := l.UnmarshalText(text); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:119
		_go_fuzz_dep_.CoverTab[131674]++
										return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:120
		// _ = "end of CoverTab[131674]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:121
		_go_fuzz_dep_.CoverTab[131675]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:121
		// _ = "end of CoverTab[131675]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:121
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:121
	// _ = "end of CoverTab[131670]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:121
	_go_fuzz_dep_.CoverTab[131671]++

									lvl.SetLevel(l)
									return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:124
	// _ = "end of CoverTab[131671]"
}

// MarshalText marshals the AtomicLevel to a byte slice. It uses the same
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:127
// text representation as the static zapcore.Levels ("debug", "info", "warn",
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:127
// "error", "dpanic", "panic", and "fatal").
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:130
func (lvl AtomicLevel) MarshalText() (text []byte, err error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:130
	_go_fuzz_dep_.CoverTab[131676]++
									return lvl.Level().MarshalText()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:131
	// _ = "end of CoverTab[131676]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/level.go:132
var _ = _go_fuzz_dep_.CoverTab
