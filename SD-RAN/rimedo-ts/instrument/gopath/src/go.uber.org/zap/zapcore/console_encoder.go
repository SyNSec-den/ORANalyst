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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:21
)

import (
	"fmt"
	"sync"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/internal/bufferpool"
)

var _sliceEncoderPool = sync.Pool{
	New: func() interface{} {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:32
		_go_fuzz_dep_.CoverTab[130607]++
													return &sliceArrayEncoder{elems: make([]interface{}, 0, 2)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:33
		// _ = "end of CoverTab[130607]"
	},
}

func getSliceEncoder() *sliceArrayEncoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:37
	_go_fuzz_dep_.CoverTab[130608]++
												return _sliceEncoderPool.Get().(*sliceArrayEncoder)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:38
	// _ = "end of CoverTab[130608]"
}

func putSliceEncoder(e *sliceArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:41
	_go_fuzz_dep_.CoverTab[130609]++
												e.elems = e.elems[:0]
												_sliceEncoderPool.Put(e)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:43
	// _ = "end of CoverTab[130609]"
}

type consoleEncoder struct {
	*jsonEncoder
}

// NewConsoleEncoder creates an encoder whose output is designed for human -
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
// rather than machine - consumption. It serializes the core log entry data
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
// (message, level, timestamp, etc.) in a plain-text format and leaves the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
// structured context as JSON.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
// Note that although the console encoder doesn't use the keys specified in the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
// encoder configuration, it will omit any element whose key is set to the empty
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:50
// string.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:58
func NewConsoleEncoder(cfg EncoderConfig) Encoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:58
	_go_fuzz_dep_.CoverTab[130610]++
												if cfg.ConsoleSeparator == "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:59
		_go_fuzz_dep_.CoverTab[130612]++

													cfg.ConsoleSeparator = "\t"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:61
		// _ = "end of CoverTab[130612]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:62
		_go_fuzz_dep_.CoverTab[130613]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:62
		// _ = "end of CoverTab[130613]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:62
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:62
	// _ = "end of CoverTab[130610]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:62
	_go_fuzz_dep_.CoverTab[130611]++
												return consoleEncoder{newJSONEncoder(cfg, true)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:63
	// _ = "end of CoverTab[130611]"
}

func (c consoleEncoder) Clone() Encoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:66
	_go_fuzz_dep_.CoverTab[130614]++
												return consoleEncoder{c.jsonEncoder.Clone().(*jsonEncoder)}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:67
	// _ = "end of CoverTab[130614]"
}

func (c consoleEncoder) EncodeEntry(ent Entry, fields []Field) (*buffer.Buffer, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:70
	_go_fuzz_dep_.CoverTab[130615]++
												line := bufferpool.Get()

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:79
	arr := getSliceEncoder()
	if c.TimeKey != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:80
		_go_fuzz_dep_.CoverTab[130624]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:80
		return c.EncodeTime != nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:80
		// _ = "end of CoverTab[130624]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:80
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:80
		_go_fuzz_dep_.CoverTab[130625]++
													c.EncodeTime(ent.Time, arr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:81
		// _ = "end of CoverTab[130625]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:82
		_go_fuzz_dep_.CoverTab[130626]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:82
		// _ = "end of CoverTab[130626]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:82
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:82
	// _ = "end of CoverTab[130615]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:82
	_go_fuzz_dep_.CoverTab[130616]++
												if c.LevelKey != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:83
		_go_fuzz_dep_.CoverTab[130627]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:83
		return c.EncodeLevel != nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:83
		// _ = "end of CoverTab[130627]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:83
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:83
		_go_fuzz_dep_.CoverTab[130628]++
													c.EncodeLevel(ent.Level, arr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:84
		// _ = "end of CoverTab[130628]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:85
		_go_fuzz_dep_.CoverTab[130629]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:85
		// _ = "end of CoverTab[130629]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:85
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:85
	// _ = "end of CoverTab[130616]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:85
	_go_fuzz_dep_.CoverTab[130617]++
												if ent.LoggerName != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:86
		_go_fuzz_dep_.CoverTab[130630]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:86
		return c.NameKey != ""
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:86
		// _ = "end of CoverTab[130630]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:86
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:86
		_go_fuzz_dep_.CoverTab[130631]++
													nameEncoder := c.EncodeName

													if nameEncoder == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:89
			_go_fuzz_dep_.CoverTab[130633]++

														nameEncoder = FullNameEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:91
			// _ = "end of CoverTab[130633]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:92
			_go_fuzz_dep_.CoverTab[130634]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:92
			// _ = "end of CoverTab[130634]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:92
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:92
		// _ = "end of CoverTab[130631]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:92
		_go_fuzz_dep_.CoverTab[130632]++

													nameEncoder(ent.LoggerName, arr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:94
		// _ = "end of CoverTab[130632]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:95
		_go_fuzz_dep_.CoverTab[130635]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:95
		// _ = "end of CoverTab[130635]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:95
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:95
	// _ = "end of CoverTab[130617]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:95
	_go_fuzz_dep_.CoverTab[130618]++
												if ent.Caller.Defined {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:96
		_go_fuzz_dep_.CoverTab[130636]++
													if c.CallerKey != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:97
			_go_fuzz_dep_.CoverTab[130638]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:97
			return c.EncodeCaller != nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:97
			// _ = "end of CoverTab[130638]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:97
		}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:97
			_go_fuzz_dep_.CoverTab[130639]++
														c.EncodeCaller(ent.Caller, arr)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:98
			// _ = "end of CoverTab[130639]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:99
			_go_fuzz_dep_.CoverTab[130640]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:99
			// _ = "end of CoverTab[130640]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:99
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:99
		// _ = "end of CoverTab[130636]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:99
		_go_fuzz_dep_.CoverTab[130637]++
													if c.FunctionKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:100
			_go_fuzz_dep_.CoverTab[130641]++
														arr.AppendString(ent.Caller.Function)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:101
			// _ = "end of CoverTab[130641]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:102
			_go_fuzz_dep_.CoverTab[130642]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:102
			// _ = "end of CoverTab[130642]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:102
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:102
		// _ = "end of CoverTab[130637]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:103
		_go_fuzz_dep_.CoverTab[130643]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:103
		// _ = "end of CoverTab[130643]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:103
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:103
	// _ = "end of CoverTab[130618]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:103
	_go_fuzz_dep_.CoverTab[130619]++
												for i := range arr.elems {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:104
		_go_fuzz_dep_.CoverTab[130644]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:105
			_go_fuzz_dep_.CoverTab[130646]++
														line.AppendString(c.ConsoleSeparator)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:106
			// _ = "end of CoverTab[130646]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:107
			_go_fuzz_dep_.CoverTab[130647]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:107
			// _ = "end of CoverTab[130647]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:107
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:107
		// _ = "end of CoverTab[130644]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:107
		_go_fuzz_dep_.CoverTab[130645]++
													fmt.Fprint(line, arr.elems[i])
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:108
		// _ = "end of CoverTab[130645]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:109
	// _ = "end of CoverTab[130619]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:109
	_go_fuzz_dep_.CoverTab[130620]++
												putSliceEncoder(arr)

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:113
	if c.MessageKey != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:113
		_go_fuzz_dep_.CoverTab[130648]++
													c.addSeparatorIfNecessary(line)
													line.AppendString(ent.Message)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:115
		// _ = "end of CoverTab[130648]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:116
		_go_fuzz_dep_.CoverTab[130649]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:116
		// _ = "end of CoverTab[130649]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:116
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:116
	// _ = "end of CoverTab[130620]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:116
	_go_fuzz_dep_.CoverTab[130621]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:119
	c.writeContext(line, fields)

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:123
	if ent.Stack != "" && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:123
		_go_fuzz_dep_.CoverTab[130650]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:123
		return c.StacktraceKey != ""
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:123
		// _ = "end of CoverTab[130650]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:123
	}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:123
		_go_fuzz_dep_.CoverTab[130651]++
													line.AppendByte('\n')
													line.AppendString(ent.Stack)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:125
		// _ = "end of CoverTab[130651]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:126
		_go_fuzz_dep_.CoverTab[130652]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:126
		// _ = "end of CoverTab[130652]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:126
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:126
	// _ = "end of CoverTab[130621]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:126
	_go_fuzz_dep_.CoverTab[130622]++

												if c.LineEnding != "" {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:128
		_go_fuzz_dep_.CoverTab[130653]++
													line.AppendString(c.LineEnding)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:129
		// _ = "end of CoverTab[130653]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:130
		_go_fuzz_dep_.CoverTab[130654]++
													line.AppendString(DefaultLineEnding)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:131
		// _ = "end of CoverTab[130654]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:132
	// _ = "end of CoverTab[130622]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:132
	_go_fuzz_dep_.CoverTab[130623]++
												return line, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:133
	// _ = "end of CoverTab[130623]"
}

func (c consoleEncoder) writeContext(line *buffer.Buffer, extra []Field) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:136
	_go_fuzz_dep_.CoverTab[130655]++
												context := c.jsonEncoder.Clone().(*jsonEncoder)
												defer func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:138
		_go_fuzz_dep_.CoverTab[130658]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:141
		context.buf.Free()
													putJSONEncoder(context)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:142
		// _ = "end of CoverTab[130658]"
	}()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:143
	// _ = "end of CoverTab[130655]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:143
	_go_fuzz_dep_.CoverTab[130656]++

												addFields(context, extra)
												context.closeOpenNamespaces()
												if context.buf.Len() == 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:147
		_go_fuzz_dep_.CoverTab[130659]++
													return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:148
		// _ = "end of CoverTab[130659]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:149
		_go_fuzz_dep_.CoverTab[130660]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:149
		// _ = "end of CoverTab[130660]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:149
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:149
	// _ = "end of CoverTab[130656]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:149
	_go_fuzz_dep_.CoverTab[130657]++

												c.addSeparatorIfNecessary(line)
												line.AppendByte('{')
												line.Write(context.buf.Bytes())
												line.AppendByte('}')
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:154
	// _ = "end of CoverTab[130657]"
}

func (c consoleEncoder) addSeparatorIfNecessary(line *buffer.Buffer) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:157
	_go_fuzz_dep_.CoverTab[130661]++
												if line.Len() > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:158
		_go_fuzz_dep_.CoverTab[130662]++
													line.AppendString(c.ConsoleSeparator)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:159
		// _ = "end of CoverTab[130662]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:160
		_go_fuzz_dep_.CoverTab[130663]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:160
		// _ = "end of CoverTab[130663]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:160
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:160
	// _ = "end of CoverTab[130661]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:161
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/console_encoder.go:161
var _ = _go_fuzz_dep_.CoverTab
