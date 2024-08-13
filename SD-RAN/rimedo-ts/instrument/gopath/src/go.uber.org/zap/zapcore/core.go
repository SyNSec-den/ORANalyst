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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:21
)

// Core is a minimal, fast logger interface. It's designed for library authors
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:23
// to wrap in a more user-friendly API.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:25
type Core interface {
	LevelEnabler

	// With adds structured context to the Core.
	With([]Field) Core
	// Check determines whether the supplied Entry should be logged (using the
	// embedded LevelEnabler and possibly some extra logic). If the entry
	// should be logged, the Core adds itself to the CheckedEntry and returns
	// the result.
	//
	// Callers must use Check before calling Write.
	Check(Entry, *CheckedEntry) *CheckedEntry
	// Write serializes the Entry and any Fields supplied at the log site and
	// writes them to their destination.
	//
	// If called, Write should always log the Entry and Fields; it should not
	// replicate the logic of Check.
	Write(Entry, []Field) error
	// Sync flushes buffered logs (if any).
	Sync() error
}

type nopCore struct{}

// NewNopCore returns a no-op Core.
func NewNopCore() Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:50
	_go_fuzz_dep_.CoverTab[130664]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:50
	return nopCore{}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:50
	// _ = "end of CoverTab[130664]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:50
}
func (nopCore) Enabled(Level) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:51
	_go_fuzz_dep_.CoverTab[130665]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:51
	return false
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:51
	// _ = "end of CoverTab[130665]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:51
}
func (n nopCore) With([]Field) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:52
	_go_fuzz_dep_.CoverTab[130666]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:52
	return n
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:52
	// _ = "end of CoverTab[130666]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:52
}
func (nopCore) Check(_ Entry, ce *CheckedEntry) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:53
	_go_fuzz_dep_.CoverTab[130667]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:53
	return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:53
	// _ = "end of CoverTab[130667]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:53
}
func (nopCore) Write(Entry, []Field) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:54
	_go_fuzz_dep_.CoverTab[130668]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:54
	return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:54
	// _ = "end of CoverTab[130668]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:54
}
func (nopCore) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:55
	_go_fuzz_dep_.CoverTab[130669]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:55
	return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:55
	// _ = "end of CoverTab[130669]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:55
}

// NewCore creates a Core that writes logs to a WriteSyncer.
func NewCore(enc Encoder, ws WriteSyncer, enab LevelEnabler) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:58
	_go_fuzz_dep_.CoverTab[130670]++
										return &ioCore{
		LevelEnabler:	enab,
		enc:		enc,
		out:		ws,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:63
	// _ = "end of CoverTab[130670]"
}

type ioCore struct {
	LevelEnabler
	enc	Encoder
	out	WriteSyncer
}

func (c *ioCore) With(fields []Field) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:72
	_go_fuzz_dep_.CoverTab[130671]++
										clone := c.clone()
										addFields(clone.enc, fields)
										return clone
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:75
	// _ = "end of CoverTab[130671]"
}

func (c *ioCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:78
	_go_fuzz_dep_.CoverTab[130672]++
										if c.Enabled(ent.Level) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:79
		_go_fuzz_dep_.CoverTab[130674]++
											return ce.AddCore(ent, c)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:80
		// _ = "end of CoverTab[130674]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:81
		_go_fuzz_dep_.CoverTab[130675]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:81
		// _ = "end of CoverTab[130675]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:81
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:81
	// _ = "end of CoverTab[130672]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:81
	_go_fuzz_dep_.CoverTab[130673]++
										return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:82
	// _ = "end of CoverTab[130673]"
}

func (c *ioCore) Write(ent Entry, fields []Field) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:85
	_go_fuzz_dep_.CoverTab[130676]++
										buf, err := c.enc.EncodeEntry(ent, fields)
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:87
		_go_fuzz_dep_.CoverTab[130680]++
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:88
		// _ = "end of CoverTab[130680]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:89
		_go_fuzz_dep_.CoverTab[130681]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:89
		// _ = "end of CoverTab[130681]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:89
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:89
	// _ = "end of CoverTab[130676]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:89
	_go_fuzz_dep_.CoverTab[130677]++
										_, err = c.out.Write(buf.Bytes())
										buf.Free()
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:92
		_go_fuzz_dep_.CoverTab[130682]++
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:93
		// _ = "end of CoverTab[130682]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:94
		_go_fuzz_dep_.CoverTab[130683]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:94
		// _ = "end of CoverTab[130683]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:94
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:94
	// _ = "end of CoverTab[130677]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:94
	_go_fuzz_dep_.CoverTab[130678]++
										if ent.Level > ErrorLevel {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:95
		_go_fuzz_dep_.CoverTab[130684]++

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:98
		c.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:98
		// _ = "end of CoverTab[130684]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:99
		_go_fuzz_dep_.CoverTab[130685]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:99
		// _ = "end of CoverTab[130685]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:99
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:99
	// _ = "end of CoverTab[130678]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:99
	_go_fuzz_dep_.CoverTab[130679]++
										return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:100
	// _ = "end of CoverTab[130679]"
}

func (c *ioCore) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:103
	_go_fuzz_dep_.CoverTab[130686]++
										return c.out.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:104
	// _ = "end of CoverTab[130686]"
}

func (c *ioCore) clone() *ioCore {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:107
	_go_fuzz_dep_.CoverTab[130687]++
										return &ioCore{
		LevelEnabler:	c.LevelEnabler,
		enc:		c.enc.Clone(),
		out:		c.out,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:112
	// _ = "end of CoverTab[130687]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/core.go:113
var _ = _go_fuzz_dep_.CoverTab
