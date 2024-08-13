// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
package logging

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:15
)

import (
	zp "go.uber.org/zap"
	zc "go.uber.org/zap/zapcore"
	"strings"
)

// Level :
type Level int

const (
	// DebugLevel logs a message at debug level
	DebugLevel	Level	= iota
	// InfoLevel logs a message at info level
	InfoLevel
	// WarnLevel logs a message at warning level
	WarnLevel
	// ErrorLevel logs a message at error level
	ErrorLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// DPanicLevel logs at PanicLevel; otherwise, it logs at ErrorLevel
	DPanicLevel

	// EmptyLevel :
	EmptyLevel	= InfoLevel
)

// String :
func (l Level) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:47
	_go_fuzz_dep_.CoverTab[132057]++
													return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC", "DPANIC", ""}[l]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:48
	// _ = "end of CoverTab[132057]"
}

func levelToAtomicLevel(l Level) zp.AtomicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:51
	_go_fuzz_dep_.CoverTab[132058]++
													switch l {
	case DebugLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:53
		_go_fuzz_dep_.CoverTab[132060]++
														return zp.NewAtomicLevelAt(zc.DebugLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:54
		// _ = "end of CoverTab[132060]"
	case InfoLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:55
		_go_fuzz_dep_.CoverTab[132061]++
														return zp.NewAtomicLevelAt(zc.InfoLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:56
		// _ = "end of CoverTab[132061]"
	case WarnLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:57
		_go_fuzz_dep_.CoverTab[132062]++
														return zp.NewAtomicLevelAt(zc.WarnLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:58
		// _ = "end of CoverTab[132062]"
	case ErrorLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:59
		_go_fuzz_dep_.CoverTab[132063]++
														return zp.NewAtomicLevelAt(zc.ErrorLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:60
		// _ = "end of CoverTab[132063]"
	case FatalLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:61
		_go_fuzz_dep_.CoverTab[132064]++
														return zp.NewAtomicLevelAt(zc.FatalLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:62
		// _ = "end of CoverTab[132064]"
	case PanicLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:63
		_go_fuzz_dep_.CoverTab[132065]++
														return zp.NewAtomicLevelAt(zc.PanicLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:64
		// _ = "end of CoverTab[132065]"
	case DPanicLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:65
		_go_fuzz_dep_.CoverTab[132066]++
														return zp.NewAtomicLevelAt(zc.DPanicLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:66
		// _ = "end of CoverTab[132066]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:66
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:66
		_go_fuzz_dep_.CoverTab[132067]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:66
		// _ = "end of CoverTab[132067]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:67
	// _ = "end of CoverTab[132058]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:67
	_go_fuzz_dep_.CoverTab[132059]++
													return zp.NewAtomicLevelAt(zc.ErrorLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:68
	// _ = "end of CoverTab[132059]"
}

func levelStringToLevel(l string) Level {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:71
	_go_fuzz_dep_.CoverTab[132068]++
													switch strings.ToUpper(l) {
	case DebugLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:73
		_go_fuzz_dep_.CoverTab[132070]++
														return DebugLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:74
		// _ = "end of CoverTab[132070]"
	case InfoLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:75
		_go_fuzz_dep_.CoverTab[132071]++
														return InfoLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:76
		// _ = "end of CoverTab[132071]"
	case WarnLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:77
		_go_fuzz_dep_.CoverTab[132072]++
														return WarnLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:78
		// _ = "end of CoverTab[132072]"
	case ErrorLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:79
		_go_fuzz_dep_.CoverTab[132073]++
														return ErrorLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:80
		// _ = "end of CoverTab[132073]"
	case FatalLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:81
		_go_fuzz_dep_.CoverTab[132074]++
														return FatalLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:82
		// _ = "end of CoverTab[132074]"
	case PanicLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:83
		_go_fuzz_dep_.CoverTab[132075]++
														return PanicLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:84
		// _ = "end of CoverTab[132075]"
	case DPanicLevel.String():
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:85
		_go_fuzz_dep_.CoverTab[132076]++
														return DPanicLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:86
		// _ = "end of CoverTab[132076]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:86
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:86
		_go_fuzz_dep_.CoverTab[132077]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:86
		// _ = "end of CoverTab[132077]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:87
	// _ = "end of CoverTab[132068]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:87
	_go_fuzz_dep_.CoverTab[132069]++
													return ErrorLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:88
	// _ = "end of CoverTab[132069]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/level.go:89
var _ = _go_fuzz_dep_.CoverTab
