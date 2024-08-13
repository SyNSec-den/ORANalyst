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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:21
)

import (
	"bytes"
	"errors"
	"fmt"
)

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

// A Level is a logging priority. Higher levels are more important.
type Level int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel	Level	= iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	_minLevel	= DebugLevel
	_maxLevel	= FatalLevel
)

// String returns a lower-case ASCII representation of the log level.
func (l Level) String() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:59
	_go_fuzz_dep_.CoverTab[131096]++
										switch l {
	case DebugLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:61
		_go_fuzz_dep_.CoverTab[131097]++
											return "debug"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:62
		// _ = "end of CoverTab[131097]"
	case InfoLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:63
		_go_fuzz_dep_.CoverTab[131098]++
											return "info"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:64
		// _ = "end of CoverTab[131098]"
	case WarnLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:65
		_go_fuzz_dep_.CoverTab[131099]++
											return "warn"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:66
		// _ = "end of CoverTab[131099]"
	case ErrorLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:67
		_go_fuzz_dep_.CoverTab[131100]++
											return "error"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:68
		// _ = "end of CoverTab[131100]"
	case DPanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:69
		_go_fuzz_dep_.CoverTab[131101]++
											return "dpanic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:70
		// _ = "end of CoverTab[131101]"
	case PanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:71
		_go_fuzz_dep_.CoverTab[131102]++
											return "panic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:72
		// _ = "end of CoverTab[131102]"
	case FatalLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:73
		_go_fuzz_dep_.CoverTab[131103]++
											return "fatal"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:74
		// _ = "end of CoverTab[131103]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:75
		_go_fuzz_dep_.CoverTab[131104]++
											return fmt.Sprintf("Level(%d)", l)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:76
		// _ = "end of CoverTab[131104]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:77
	// _ = "end of CoverTab[131096]"
}

// CapitalString returns an all-caps ASCII representation of the log level.
func (l Level) CapitalString() string {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:81
	_go_fuzz_dep_.CoverTab[131105]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:84
	switch l {
	case DebugLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:85
		_go_fuzz_dep_.CoverTab[131106]++
											return "DEBUG"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:86
		// _ = "end of CoverTab[131106]"
	case InfoLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:87
		_go_fuzz_dep_.CoverTab[131107]++
											return "INFO"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:88
		// _ = "end of CoverTab[131107]"
	case WarnLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:89
		_go_fuzz_dep_.CoverTab[131108]++
											return "WARN"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:90
		// _ = "end of CoverTab[131108]"
	case ErrorLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:91
		_go_fuzz_dep_.CoverTab[131109]++
											return "ERROR"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:92
		// _ = "end of CoverTab[131109]"
	case DPanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:93
		_go_fuzz_dep_.CoverTab[131110]++
											return "DPANIC"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:94
		// _ = "end of CoverTab[131110]"
	case PanicLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:95
		_go_fuzz_dep_.CoverTab[131111]++
											return "PANIC"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:96
		// _ = "end of CoverTab[131111]"
	case FatalLevel:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:97
		_go_fuzz_dep_.CoverTab[131112]++
											return "FATAL"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:98
		// _ = "end of CoverTab[131112]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:99
		_go_fuzz_dep_.CoverTab[131113]++
											return fmt.Sprintf("LEVEL(%d)", l)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:100
		// _ = "end of CoverTab[131113]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:101
	// _ = "end of CoverTab[131105]"
}

// MarshalText marshals the Level to text. Note that the text representation
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:104
// drops the -Level suffix (see example).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:106
func (l Level) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:106
	_go_fuzz_dep_.CoverTab[131114]++
										return []byte(l.String()), nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:107
	// _ = "end of CoverTab[131114]"
}

// UnmarshalText unmarshals text to a level. Like MarshalText, UnmarshalText
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:110
// expects the text representation of a Level to drop the -Level suffix (see
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:110
// example).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:110
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:110
// In particular, this makes it easy to configure logging levels using YAML,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:110
// TOML, or JSON files.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:116
func (l *Level) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:116
	_go_fuzz_dep_.CoverTab[131115]++
										if l == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:117
		_go_fuzz_dep_.CoverTab[131118]++
											return errUnmarshalNilLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:118
		// _ = "end of CoverTab[131118]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:119
		_go_fuzz_dep_.CoverTab[131119]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:119
		// _ = "end of CoverTab[131119]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:119
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:119
	// _ = "end of CoverTab[131115]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:119
	_go_fuzz_dep_.CoverTab[131116]++
										if !l.unmarshalText(text) && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:120
		_go_fuzz_dep_.CoverTab[131120]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:120
		return !l.unmarshalText(bytes.ToLower(text))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:120
		// _ = "end of CoverTab[131120]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:120
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:120
		_go_fuzz_dep_.CoverTab[131121]++
											return fmt.Errorf("unrecognized level: %q", text)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:121
		// _ = "end of CoverTab[131121]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:122
		_go_fuzz_dep_.CoverTab[131122]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:122
		// _ = "end of CoverTab[131122]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:122
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:122
	// _ = "end of CoverTab[131116]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:122
	_go_fuzz_dep_.CoverTab[131117]++
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:123
	// _ = "end of CoverTab[131117]"
}

func (l *Level) unmarshalText(text []byte) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:126
	_go_fuzz_dep_.CoverTab[131123]++
										switch string(text) {
	case "debug", "DEBUG":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:128
		_go_fuzz_dep_.CoverTab[131125]++
											*l = DebugLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:129
		// _ = "end of CoverTab[131125]"
	case "info", "INFO", "":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:130
		_go_fuzz_dep_.CoverTab[131126]++
											*l = InfoLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:131
		// _ = "end of CoverTab[131126]"
	case "warn", "WARN":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:132
		_go_fuzz_dep_.CoverTab[131127]++
											*l = WarnLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:133
		// _ = "end of CoverTab[131127]"
	case "error", "ERROR":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:134
		_go_fuzz_dep_.CoverTab[131128]++
											*l = ErrorLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:135
		// _ = "end of CoverTab[131128]"
	case "dpanic", "DPANIC":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:136
		_go_fuzz_dep_.CoverTab[131129]++
											*l = DPanicLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:137
		// _ = "end of CoverTab[131129]"
	case "panic", "PANIC":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:138
		_go_fuzz_dep_.CoverTab[131130]++
											*l = PanicLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:139
		// _ = "end of CoverTab[131130]"
	case "fatal", "FATAL":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:140
		_go_fuzz_dep_.CoverTab[131131]++
											*l = FatalLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:141
		// _ = "end of CoverTab[131131]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:142
		_go_fuzz_dep_.CoverTab[131132]++
											return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:143
		// _ = "end of CoverTab[131132]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:144
	// _ = "end of CoverTab[131123]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:144
	_go_fuzz_dep_.CoverTab[131124]++
										return true
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:145
	// _ = "end of CoverTab[131124]"
}

// Set sets the level for the flag.Value interface.
func (l *Level) Set(s string) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:149
	_go_fuzz_dep_.CoverTab[131133]++
										return l.UnmarshalText([]byte(s))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:150
	// _ = "end of CoverTab[131133]"
}

// Get gets the level for the flag.Getter interface.
func (l *Level) Get() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:154
	_go_fuzz_dep_.CoverTab[131134]++
										return *l
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:155
	// _ = "end of CoverTab[131134]"
}

// Enabled returns true if the given level is at or above this level.
func (l Level) Enabled(lvl Level) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:159
	_go_fuzz_dep_.CoverTab[131135]++
										return lvl >= l
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:160
	// _ = "end of CoverTab[131135]"
}

// LevelEnabler decides whether a given logging level is enabled when logging a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// message.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// Enablers are intended to be used to implement deterministic filters;
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// concerns like sampling are better implemented as a Core.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// Each concrete Level value implements a static LevelEnabler which returns
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// true for itself and all higher logging levels. For example WarnLevel.Enabled()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// will return true for WarnLevel, ErrorLevel, DPanicLevel, PanicLevel, and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:163
// FatalLevel, but return false for InfoLevel and DebugLevel.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:173
type LevelEnabler interface {
	Enabled(Level) bool
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:175
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/level.go:175
var _ = _go_fuzz_dep_.CoverTab
