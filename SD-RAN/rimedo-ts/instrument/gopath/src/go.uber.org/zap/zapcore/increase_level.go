// Copyright (c) 2020 Uber Technologies, Inc.
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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:21
)

import "fmt"

type levelFilterCore struct {
	core	Core
	level	LevelEnabler
}

// NewIncreaseLevelCore creates a core that can be used to increase the level of
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:30
// an existing Core. It cannot be used to decrease the logging level, as it acts
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:30
// as a filter before calling the underlying core. If level decreases the log level,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:30
// an error is returned.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:34
func NewIncreaseLevelCore(core Core, level LevelEnabler) (Core, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:34
	_go_fuzz_dep_.CoverTab[130901]++
											for l := _maxLevel; l >= _minLevel; l-- {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:35
		_go_fuzz_dep_.CoverTab[130903]++
												if !core.Enabled(l) && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:36
			_go_fuzz_dep_.CoverTab[130904]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:36
			return level.Enabled(l)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:36
			// _ = "end of CoverTab[130904]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:36
		}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:36
			_go_fuzz_dep_.CoverTab[130905]++
													return nil, fmt.Errorf("invalid increase level, as level %q is allowed by increased level, but not by existing core", l)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:37
			// _ = "end of CoverTab[130905]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:38
			_go_fuzz_dep_.CoverTab[130906]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:38
			// _ = "end of CoverTab[130906]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:38
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:38
		// _ = "end of CoverTab[130903]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:39
	// _ = "end of CoverTab[130901]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:39
	_go_fuzz_dep_.CoverTab[130902]++

											return &levelFilterCore{core, level}, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:41
	// _ = "end of CoverTab[130902]"
}

func (c *levelFilterCore) Enabled(lvl Level) bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:44
	_go_fuzz_dep_.CoverTab[130907]++
											return c.level.Enabled(lvl)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:45
	// _ = "end of CoverTab[130907]"
}

func (c *levelFilterCore) With(fields []Field) Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:48
	_go_fuzz_dep_.CoverTab[130908]++
											return &levelFilterCore{c.core.With(fields), c.level}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:49
	// _ = "end of CoverTab[130908]"
}

func (c *levelFilterCore) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:52
	_go_fuzz_dep_.CoverTab[130909]++
											if !c.Enabled(ent.Level) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:53
		_go_fuzz_dep_.CoverTab[130911]++
												return ce
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:54
		// _ = "end of CoverTab[130911]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:55
		_go_fuzz_dep_.CoverTab[130912]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:55
		// _ = "end of CoverTab[130912]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:55
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:55
	// _ = "end of CoverTab[130909]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:55
	_go_fuzz_dep_.CoverTab[130910]++

											return c.core.Check(ent, ce)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:57
	// _ = "end of CoverTab[130910]"
}

func (c *levelFilterCore) Write(ent Entry, fields []Field) error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:60
	_go_fuzz_dep_.CoverTab[130913]++
											return c.core.Write(ent, fields)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:61
	// _ = "end of CoverTab[130913]"
}

func (c *levelFilterCore) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:64
	_go_fuzz_dep_.CoverTab[130914]++
											return c.core.Sync()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:65
	// _ = "end of CoverTab[130914]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:66
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/increase_level.go:66
var _ = _go_fuzz_dep_.CoverTab
