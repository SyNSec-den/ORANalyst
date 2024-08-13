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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:21
)

import (
	"encoding/json"
	"time"

	"go.uber.org/zap/buffer"
)

// DefaultLineEnding defines the default line ending when writing logs.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:30
// Alternate line endings specified in EncoderConfig can override this
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:30
// behavior.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:33
const DefaultLineEnding = "\n"

// OmitKey defines the key to use when callers want to remove a key from log output.
const OmitKey = ""

// A LevelEncoder serializes a Level to a primitive type.
type LevelEncoder func(Level, PrimitiveArrayEncoder)

// LowercaseLevelEncoder serializes a Level to a lowercase string. For example,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:41
// InfoLevel is serialized to "info".
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:43
func LowercaseLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:43
	_go_fuzz_dep_.CoverTab[130688]++
											enc.AppendString(l.String())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:44
	// _ = "end of CoverTab[130688]"
}

// LowercaseColorLevelEncoder serializes a Level to a lowercase string and adds coloring.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:47
// For example, InfoLevel is serialized to "info" and colored blue.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:49
func LowercaseColorLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:49
	_go_fuzz_dep_.CoverTab[130689]++
											s, ok := _levelToLowercaseColorString[l]
											if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:51
		_go_fuzz_dep_.CoverTab[130691]++
												s = _unknownLevelColor.Add(l.String())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:52
		// _ = "end of CoverTab[130691]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:53
		_go_fuzz_dep_.CoverTab[130692]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:53
		// _ = "end of CoverTab[130692]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:53
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:53
	// _ = "end of CoverTab[130689]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:53
	_go_fuzz_dep_.CoverTab[130690]++
											enc.AppendString(s)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:54
	// _ = "end of CoverTab[130690]"
}

// CapitalLevelEncoder serializes a Level to an all-caps string. For example,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:57
// InfoLevel is serialized to "INFO".
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:59
func CapitalLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:59
	_go_fuzz_dep_.CoverTab[130693]++
											enc.AppendString(l.CapitalString())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:60
	// _ = "end of CoverTab[130693]"
}

// CapitalColorLevelEncoder serializes a Level to an all-caps string and adds color.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:63
// For example, InfoLevel is serialized to "INFO" and colored blue.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:65
func CapitalColorLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:65
	_go_fuzz_dep_.CoverTab[130694]++
											s, ok := _levelToCapitalColorString[l]
											if !ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:67
		_go_fuzz_dep_.CoverTab[130696]++
												s = _unknownLevelColor.Add(l.CapitalString())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:68
		// _ = "end of CoverTab[130696]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:69
		_go_fuzz_dep_.CoverTab[130697]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:69
		// _ = "end of CoverTab[130697]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:69
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:69
	// _ = "end of CoverTab[130694]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:69
	_go_fuzz_dep_.CoverTab[130695]++
											enc.AppendString(s)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:70
	// _ = "end of CoverTab[130695]"
}

// UnmarshalText unmarshals text to a LevelEncoder. "capital" is unmarshaled to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:73
// CapitalLevelEncoder, "coloredCapital" is unmarshaled to CapitalColorLevelEncoder,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:73
// "colored" is unmarshaled to LowercaseColorLevelEncoder, and anything else
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:73
// is unmarshaled to LowercaseLevelEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:77
func (e *LevelEncoder) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:77
	_go_fuzz_dep_.CoverTab[130698]++
											switch string(text) {
	case "capital":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:79
		_go_fuzz_dep_.CoverTab[130700]++
												*e = CapitalLevelEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:80
		// _ = "end of CoverTab[130700]"
	case "capitalColor":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:81
		_go_fuzz_dep_.CoverTab[130701]++
												*e = CapitalColorLevelEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:82
		// _ = "end of CoverTab[130701]"
	case "color":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:83
		_go_fuzz_dep_.CoverTab[130702]++
												*e = LowercaseColorLevelEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:84
		// _ = "end of CoverTab[130702]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:85
		_go_fuzz_dep_.CoverTab[130703]++
												*e = LowercaseLevelEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:86
		// _ = "end of CoverTab[130703]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:87
	// _ = "end of CoverTab[130698]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:87
	_go_fuzz_dep_.CoverTab[130699]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:88
	// _ = "end of CoverTab[130699]"
}

// A TimeEncoder serializes a time.Time to a primitive type.
type TimeEncoder func(time.Time, PrimitiveArrayEncoder)

// EpochTimeEncoder serializes a time.Time to a floating-point number of seconds
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:94
// since the Unix epoch.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:96
func EpochTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:96
	_go_fuzz_dep_.CoverTab[130704]++
											nanos := t.UnixNano()
											sec := float64(nanos) / float64(time.Second)
											enc.AppendFloat64(sec)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:99
	// _ = "end of CoverTab[130704]"
}

// EpochMillisTimeEncoder serializes a time.Time to a floating-point number of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:102
// milliseconds since the Unix epoch.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:104
func EpochMillisTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:104
	_go_fuzz_dep_.CoverTab[130705]++
											nanos := t.UnixNano()
											millis := float64(nanos) / float64(time.Millisecond)
											enc.AppendFloat64(millis)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:107
	// _ = "end of CoverTab[130705]"
}

// EpochNanosTimeEncoder serializes a time.Time to an integer number of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:110
// nanoseconds since the Unix epoch.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:112
func EpochNanosTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:112
	_go_fuzz_dep_.CoverTab[130706]++
											enc.AppendInt64(t.UnixNano())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:113
	// _ = "end of CoverTab[130706]"
}

func encodeTimeLayout(t time.Time, layout string, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:116
	_go_fuzz_dep_.CoverTab[130707]++
											type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:121
		_go_fuzz_dep_.CoverTab[130709]++
												enc.AppendTimeLayout(t, layout)
												return
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:123
		// _ = "end of CoverTab[130709]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:124
		_go_fuzz_dep_.CoverTab[130710]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:124
		// _ = "end of CoverTab[130710]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:124
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:124
	// _ = "end of CoverTab[130707]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:124
	_go_fuzz_dep_.CoverTab[130708]++

											enc.AppendString(t.Format(layout))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:126
	// _ = "end of CoverTab[130708]"
}

// ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:129
// with millisecond precision.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:129
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:129
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:129
// instead of appending a pre-formatted string value.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:134
func ISO8601TimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:134
	_go_fuzz_dep_.CoverTab[130711]++
											encodeTimeLayout(t, "2006-01-02T15:04:05.000Z0700", enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:135
	// _ = "end of CoverTab[130711]"
}

// RFC3339TimeEncoder serializes a time.Time to an RFC3339-formatted string.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:138
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:138
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:138
// instead of appending a pre-formatted string value.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:142
func RFC3339TimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:142
	_go_fuzz_dep_.CoverTab[130712]++
											encodeTimeLayout(t, time.RFC3339, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:143
	// _ = "end of CoverTab[130712]"
}

// RFC3339NanoTimeEncoder serializes a time.Time to an RFC3339-formatted string
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:146
// with nanosecond precision.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:146
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:146
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:146
// instead of appending a pre-formatted string value.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:151
func RFC3339NanoTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:151
	_go_fuzz_dep_.CoverTab[130713]++
											encodeTimeLayout(t, time.RFC3339Nano, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:152
	// _ = "end of CoverTab[130713]"
}

// TimeEncoderOfLayout returns TimeEncoder which serializes a time.Time using
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:155
// given layout.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:157
func TimeEncoderOfLayout(layout string) TimeEncoder {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:157
	_go_fuzz_dep_.CoverTab[130714]++
											return func(t time.Time, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:158
		_go_fuzz_dep_.CoverTab[130715]++
												encodeTimeLayout(t, layout, enc)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:159
		// _ = "end of CoverTab[130715]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:160
	// _ = "end of CoverTab[130714]"
}

// UnmarshalText unmarshals text to a TimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:163
// "rfc3339nano" and "RFC3339Nano" are unmarshaled to RFC3339NanoTimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:163
// "rfc3339" and "RFC3339" are unmarshaled to RFC3339TimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:163
// "iso8601" and "ISO8601" are unmarshaled to ISO8601TimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:163
// "millis" is unmarshaled to EpochMillisTimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:163
// "nanos" is unmarshaled to EpochNanosEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:163
// Anything else is unmarshaled to EpochTimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:170
func (e *TimeEncoder) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:170
	_go_fuzz_dep_.CoverTab[130716]++
											switch string(text) {
	case "rfc3339nano", "RFC3339Nano":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:172
		_go_fuzz_dep_.CoverTab[130718]++
												*e = RFC3339NanoTimeEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:173
		// _ = "end of CoverTab[130718]"
	case "rfc3339", "RFC3339":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:174
		_go_fuzz_dep_.CoverTab[130719]++
												*e = RFC3339TimeEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:175
		// _ = "end of CoverTab[130719]"
	case "iso8601", "ISO8601":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:176
		_go_fuzz_dep_.CoverTab[130720]++
												*e = ISO8601TimeEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:177
		// _ = "end of CoverTab[130720]"
	case "millis":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:178
		_go_fuzz_dep_.CoverTab[130721]++
												*e = EpochMillisTimeEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:179
		// _ = "end of CoverTab[130721]"
	case "nanos":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:180
		_go_fuzz_dep_.CoverTab[130722]++
												*e = EpochNanosTimeEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:181
		// _ = "end of CoverTab[130722]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:182
		_go_fuzz_dep_.CoverTab[130723]++
												*e = EpochTimeEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:183
		// _ = "end of CoverTab[130723]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:184
	// _ = "end of CoverTab[130716]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:184
	_go_fuzz_dep_.CoverTab[130717]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:185
	// _ = "end of CoverTab[130717]"
}

// UnmarshalYAML unmarshals YAML to a TimeEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
// If value is an object with a "layout" field, it will be unmarshaled to  TimeEncoder with given layout.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
//	timeEncoder:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
//	  layout: 06/01/02 03:04pm
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
// If value is string, it uses UnmarshalText.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:188
//	timeEncoder: iso8601
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:194
func (e *TimeEncoder) UnmarshalYAML(unmarshal func(interface{}) error) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:194
	_go_fuzz_dep_.CoverTab[130724]++
											var o struct {
		Layout string `json:"layout" yaml:"layout"`
	}
	if err := unmarshal(&o); err == nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:198
		_go_fuzz_dep_.CoverTab[130727]++
												*e = TimeEncoderOfLayout(o.Layout)
												return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:200
		// _ = "end of CoverTab[130727]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:201
		_go_fuzz_dep_.CoverTab[130728]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:201
		// _ = "end of CoverTab[130728]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:201
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:201
	// _ = "end of CoverTab[130724]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:201
	_go_fuzz_dep_.CoverTab[130725]++

											var s string
											if err := unmarshal(&s); err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:204
		_go_fuzz_dep_.CoverTab[130729]++
												return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:205
		// _ = "end of CoverTab[130729]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:206
		_go_fuzz_dep_.CoverTab[130730]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:206
		// _ = "end of CoverTab[130730]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:206
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:206
	// _ = "end of CoverTab[130725]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:206
	_go_fuzz_dep_.CoverTab[130726]++
											return e.UnmarshalText([]byte(s))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:207
	// _ = "end of CoverTab[130726]"
}

// UnmarshalJSON unmarshals JSON to a TimeEncoder as same way UnmarshalYAML does.
func (e *TimeEncoder) UnmarshalJSON(data []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:211
	_go_fuzz_dep_.CoverTab[130731]++
											return e.UnmarshalYAML(func(v interface{}) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:212
		_go_fuzz_dep_.CoverTab[130732]++
												return json.Unmarshal(data, v)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:213
		// _ = "end of CoverTab[130732]"
	})
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:214
	// _ = "end of CoverTab[130731]"
}

// A DurationEncoder serializes a time.Duration to a primitive type.
type DurationEncoder func(time.Duration, PrimitiveArrayEncoder)

// SecondsDurationEncoder serializes a time.Duration to a floating-point number of seconds elapsed.
func SecondsDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:221
	_go_fuzz_dep_.CoverTab[130733]++
											enc.AppendFloat64(float64(d) / float64(time.Second))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:222
	// _ = "end of CoverTab[130733]"
}

// NanosDurationEncoder serializes a time.Duration to an integer number of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:225
// nanoseconds elapsed.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:227
func NanosDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:227
	_go_fuzz_dep_.CoverTab[130734]++
											enc.AppendInt64(int64(d))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:228
	// _ = "end of CoverTab[130734]"
}

// MillisDurationEncoder serializes a time.Duration to an integer number of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:231
// milliseconds elapsed.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:233
func MillisDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:233
	_go_fuzz_dep_.CoverTab[130735]++
											enc.AppendInt64(d.Nanoseconds() / 1e6)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:234
	// _ = "end of CoverTab[130735]"
}

// StringDurationEncoder serializes a time.Duration using its built-in String
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:237
// method.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:239
func StringDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:239
	_go_fuzz_dep_.CoverTab[130736]++
											enc.AppendString(d.String())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:240
	// _ = "end of CoverTab[130736]"
}

// UnmarshalText unmarshals text to a DurationEncoder. "string" is unmarshaled
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:243
// to StringDurationEncoder, and anything else is unmarshaled to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:243
// NanosDurationEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:246
func (e *DurationEncoder) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:246
	_go_fuzz_dep_.CoverTab[130737]++
											switch string(text) {
	case "string":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:248
		_go_fuzz_dep_.CoverTab[130739]++
												*e = StringDurationEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:249
		// _ = "end of CoverTab[130739]"
	case "nanos":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:250
		_go_fuzz_dep_.CoverTab[130740]++
												*e = NanosDurationEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:251
		// _ = "end of CoverTab[130740]"
	case "ms":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:252
		_go_fuzz_dep_.CoverTab[130741]++
												*e = MillisDurationEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:253
		// _ = "end of CoverTab[130741]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:254
		_go_fuzz_dep_.CoverTab[130742]++
												*e = SecondsDurationEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:255
		// _ = "end of CoverTab[130742]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:256
	// _ = "end of CoverTab[130737]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:256
	_go_fuzz_dep_.CoverTab[130738]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:257
	// _ = "end of CoverTab[130738]"
}

// A CallerEncoder serializes an EntryCaller to a primitive type.
type CallerEncoder func(EntryCaller, PrimitiveArrayEncoder)

// FullCallerEncoder serializes a caller in /full/path/to/package/file:line
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:263
// format.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:265
func FullCallerEncoder(caller EntryCaller, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:265
	_go_fuzz_dep_.CoverTab[130743]++

											enc.AppendString(caller.String())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:267
	// _ = "end of CoverTab[130743]"
}

// ShortCallerEncoder serializes a caller in package/file:line format, trimming
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:270
// all but the final directory from the full path.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:272
func ShortCallerEncoder(caller EntryCaller, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:272
	_go_fuzz_dep_.CoverTab[130744]++

											enc.AppendString(caller.TrimmedPath())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:274
	// _ = "end of CoverTab[130744]"
}

// UnmarshalText unmarshals text to a CallerEncoder. "full" is unmarshaled to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:277
// FullCallerEncoder and anything else is unmarshaled to ShortCallerEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:279
func (e *CallerEncoder) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:279
	_go_fuzz_dep_.CoverTab[130745]++
											switch string(text) {
	case "full":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:281
		_go_fuzz_dep_.CoverTab[130747]++
												*e = FullCallerEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:282
		// _ = "end of CoverTab[130747]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:283
		_go_fuzz_dep_.CoverTab[130748]++
												*e = ShortCallerEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:284
		// _ = "end of CoverTab[130748]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:285
	// _ = "end of CoverTab[130745]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:285
	_go_fuzz_dep_.CoverTab[130746]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:286
	// _ = "end of CoverTab[130746]"
}

// A NameEncoder serializes a period-separated logger name to a primitive
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:289
// type.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:291
type NameEncoder func(string, PrimitiveArrayEncoder)

// FullNameEncoder serializes the logger name as-is.
func FullNameEncoder(loggerName string, enc PrimitiveArrayEncoder) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:294
	_go_fuzz_dep_.CoverTab[130749]++
											enc.AppendString(loggerName)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:295
	// _ = "end of CoverTab[130749]"
}

// UnmarshalText unmarshals text to a NameEncoder. Currently, everything is
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:298
// unmarshaled to FullNameEncoder.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:300
func (e *NameEncoder) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:300
	_go_fuzz_dep_.CoverTab[130750]++
											switch string(text) {
	case "full":
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:302
		_go_fuzz_dep_.CoverTab[130752]++
												*e = FullNameEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:303
		// _ = "end of CoverTab[130752]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:304
		_go_fuzz_dep_.CoverTab[130753]++
												*e = FullNameEncoder
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:305
		// _ = "end of CoverTab[130753]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:306
	// _ = "end of CoverTab[130750]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:306
	_go_fuzz_dep_.CoverTab[130751]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:307
	// _ = "end of CoverTab[130751]"
}

// An EncoderConfig allows users to configure the concrete encoders supplied by
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:310
// zapcore.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:312
type EncoderConfig struct {
	// Set the keys used for each log entry. If any key is empty, that portion
	// of the entry is omitted.
	MessageKey	string	`json:"messageKey" yaml:"messageKey"`
	LevelKey	string	`json:"levelKey" yaml:"levelKey"`
	TimeKey		string	`json:"timeKey" yaml:"timeKey"`
	NameKey		string	`json:"nameKey" yaml:"nameKey"`
	CallerKey	string	`json:"callerKey" yaml:"callerKey"`
	FunctionKey	string	`json:"functionKey" yaml:"functionKey"`
	StacktraceKey	string	`json:"stacktraceKey" yaml:"stacktraceKey"`
	LineEnding	string	`json:"lineEnding" yaml:"lineEnding"`
	// Configure the primitive representations of common complex types. For
	// example, some users may want all time.Times serialized as floating-point
	// seconds since epoch, while others may prefer ISO8601 strings.
	EncodeLevel	LevelEncoder	`json:"levelEncoder" yaml:"levelEncoder"`
	EncodeTime	TimeEncoder	`json:"timeEncoder" yaml:"timeEncoder"`
	EncodeDuration	DurationEncoder	`json:"durationEncoder" yaml:"durationEncoder"`
	EncodeCaller	CallerEncoder	`json:"callerEncoder" yaml:"callerEncoder"`
	// Unlike the other primitive type encoders, EncodeName is optional. The
	// zero value falls back to FullNameEncoder.
	EncodeName	NameEncoder	`json:"nameEncoder" yaml:"nameEncoder"`
	// Configures the field separator used by the console encoder. Defaults
	// to tab.
	ConsoleSeparator	string	`json:"consoleSeparator" yaml:"consoleSeparator"`
}

// ObjectEncoder is a strongly-typed, encoding-agnostic interface for adding a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:338
// map- or struct-like object to the logging context. Like maps, ObjectEncoders
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:338
// aren't safe for concurrent use (though typical use shouldn't require locks).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:341
type ObjectEncoder interface {
	// Logging-specific marshalers.
	AddArray(key string, marshaler ArrayMarshaler) error
	AddObject(key string, marshaler ObjectMarshaler) error

	// Built-in types.
	AddBinary(key string, value []byte)	// for arbitrary bytes
	AddByteString(key string, value []byte)	// for UTF-8 encoded bytes
	AddBool(key string, value bool)
	AddComplex128(key string, value complex128)
	AddComplex64(key string, value complex64)
	AddDuration(key string, value time.Duration)
	AddFloat64(key string, value float64)
	AddFloat32(key string, value float32)
	AddInt(key string, value int)
	AddInt64(key string, value int64)
	AddInt32(key string, value int32)
	AddInt16(key string, value int16)
	AddInt8(key string, value int8)
	AddString(key, value string)
	AddTime(key string, value time.Time)
	AddUint(key string, value uint)
	AddUint64(key string, value uint64)
	AddUint32(key string, value uint32)
	AddUint16(key string, value uint16)
	AddUint8(key string, value uint8)
	AddUintptr(key string, value uintptr)

	// AddReflected uses reflection to serialize arbitrary objects, so it can be
	// slow and allocation-heavy.
	AddReflected(key string, value interface{}) error
	// OpenNamespace opens an isolated namespace where all subsequent fields will
	// be added. Applications can use namespaces to prevent key collisions when
	// injecting loggers into sub-components or third-party libraries.
	OpenNamespace(key string)
}

// ArrayEncoder is a strongly-typed, encoding-agnostic interface for adding
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:378
// array-like objects to the logging context. Of note, it supports mixed-type
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:378
// arrays even though they aren't typical in Go. Like slices, ArrayEncoders
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:378
// aren't safe for concurrent use (though typical use shouldn't require locks).
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:382
type ArrayEncoder interface {
	// Built-in types.
	PrimitiveArrayEncoder

	// Time-related types.
	AppendDuration(time.Duration)
	AppendTime(time.Time)

	// Logging-specific marshalers.
	AppendArray(ArrayMarshaler) error
	AppendObject(ObjectMarshaler) error

	// AppendReflected uses reflection to serialize arbitrary objects, so it's
	// slow and allocation-heavy.
	AppendReflected(value interface{}) error
}

// PrimitiveArrayEncoder is the subset of the ArrayEncoder interface that deals
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:399
// only in Go's built-in types. It's included only so that Duration- and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:399
// TimeEncoders cannot trigger infinite recursion.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:402
type PrimitiveArrayEncoder interface {
	// Built-in types.
	AppendBool(bool)
	AppendByteString([]byte)	// for UTF-8 encoded bytes
	AppendComplex128(complex128)
	AppendComplex64(complex64)
	AppendFloat64(float64)
	AppendFloat32(float32)
	AppendInt(int)
	AppendInt64(int64)
	AppendInt32(int32)
	AppendInt16(int16)
	AppendInt8(int8)
	AppendString(string)
	AppendUint(uint)
	AppendUint64(uint64)
	AppendUint32(uint32)
	AppendUint16(uint16)
	AppendUint8(uint8)
	AppendUintptr(uintptr)
}

// Encoder is a format-agnostic interface for all log entry marshalers. Since
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
// log encoders don't need to support the same wide range of use cases as
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
// general-purpose marshalers, it's possible to make them faster and
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
// lower-allocation.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
// Implementations of the ObjectEncoder interface's methods can, of course,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
// freely modify the receiver. However, the Clone and EncodeEntry methods will
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:424
// be called concurrently and shouldn't modify the receiver.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:432
type Encoder interface {
	ObjectEncoder

	// Clone copies the encoder, ensuring that adding fields to the copy doesn't
	// affect the original.
	Clone() Encoder

	// EncodeEntry encodes an entry and fields, along with any accumulated
	// context, into a byte buffer and returns it. Any fields that are empty,
	// including fields on the `Entry` type, should be omitted.
	EncodeEntry(Entry, []Field) (*buffer.Buffer, error)
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:443
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/encoder.go:443
var _ = _go_fuzz_dep_.CoverTab
