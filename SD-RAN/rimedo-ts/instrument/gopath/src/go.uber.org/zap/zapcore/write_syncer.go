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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
package zapcore

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:21
)

import (
	"io"
	"sync"

	"go.uber.org/multierr"
)

// A WriteSyncer is an io.Writer that can also flush any buffered data. Note
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:30
// that *os.File (and thus, os.Stderr and os.Stdout) implement WriteSyncer.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:32
type WriteSyncer interface {
	io.Writer
	Sync() error
}

// AddSync converts an io.Writer to a WriteSyncer. It attempts to be
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:37
// intelligent: if the concrete type of the io.Writer implements WriteSyncer,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:37
// we'll use the existing Sync method. If it doesn't, we'll add a no-op Sync.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:40
func AddSync(w io.Writer) WriteSyncer {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:40
	_go_fuzz_dep_.CoverTab[131237]++
											switch w := w.(type) {
	case WriteSyncer:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:42
		_go_fuzz_dep_.CoverTab[131238]++
												return w
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:43
		// _ = "end of CoverTab[131238]"
	default:
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:44
		_go_fuzz_dep_.CoverTab[131239]++
												return writerWrapper{w}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:45
		// _ = "end of CoverTab[131239]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:46
	// _ = "end of CoverTab[131237]"
}

type lockedWriteSyncer struct {
	sync.Mutex
	ws	WriteSyncer
}

// Lock wraps a WriteSyncer in a mutex to make it safe for concurrent use. In
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:54
// particular, *os.Files must be locked before use.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:56
func Lock(ws WriteSyncer) WriteSyncer {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:56
	_go_fuzz_dep_.CoverTab[131240]++
											if _, ok := ws.(*lockedWriteSyncer); ok {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:57
		_go_fuzz_dep_.CoverTab[131242]++

												return ws
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:59
		// _ = "end of CoverTab[131242]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:60
		_go_fuzz_dep_.CoverTab[131243]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:60
		// _ = "end of CoverTab[131243]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:60
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:60
	// _ = "end of CoverTab[131240]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:60
	_go_fuzz_dep_.CoverTab[131241]++
											return &lockedWriteSyncer{ws: ws}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:61
	// _ = "end of CoverTab[131241]"
}

func (s *lockedWriteSyncer) Write(bs []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:64
	_go_fuzz_dep_.CoverTab[131244]++
											s.Lock()
											n, err := s.ws.Write(bs)
											s.Unlock()
											return n, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:68
	// _ = "end of CoverTab[131244]"
}

func (s *lockedWriteSyncer) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:71
	_go_fuzz_dep_.CoverTab[131245]++
											s.Lock()
											err := s.ws.Sync()
											s.Unlock()
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:75
	// _ = "end of CoverTab[131245]"
}

type writerWrapper struct {
	io.Writer
}

func (w writerWrapper) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:82
	_go_fuzz_dep_.CoverTab[131246]++
											return nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:83
	// _ = "end of CoverTab[131246]"
}

type multiWriteSyncer []WriteSyncer

// NewMultiWriteSyncer creates a WriteSyncer that duplicates its writes
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:88
// and sync calls, much like io.MultiWriter.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:90
func NewMultiWriteSyncer(ws ...WriteSyncer) WriteSyncer {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:90
	_go_fuzz_dep_.CoverTab[131247]++
											if len(ws) == 1 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:91
		_go_fuzz_dep_.CoverTab[131249]++
												return ws[0]
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:92
		// _ = "end of CoverTab[131249]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:93
		_go_fuzz_dep_.CoverTab[131250]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:93
		// _ = "end of CoverTab[131250]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:93
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:93
	// _ = "end of CoverTab[131247]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:93
	_go_fuzz_dep_.CoverTab[131248]++
											return multiWriteSyncer(ws)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:94
	// _ = "end of CoverTab[131248]"
}

// See https://golang.org/src/io/multi.go
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:97
// When not all underlying syncers write the same number of bytes,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:97
// the smallest number is returned even though Write() is called on
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:97
// all of them.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:101
func (ws multiWriteSyncer) Write(p []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:101
	_go_fuzz_dep_.CoverTab[131251]++
											var writeErr error
											nWritten := 0
											for _, w := range ws {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:104
		_go_fuzz_dep_.CoverTab[131253]++
												n, err := w.Write(p)
												writeErr = multierr.Append(writeErr, err)
												if nWritten == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:107
			_go_fuzz_dep_.CoverTab[131254]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:107
			return n != 0
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:107
			// _ = "end of CoverTab[131254]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:107
		}() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:107
			_go_fuzz_dep_.CoverTab[131255]++
													nWritten = n
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:108
			// _ = "end of CoverTab[131255]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:109
			_go_fuzz_dep_.CoverTab[131256]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:109
			if n < nWritten {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:109
				_go_fuzz_dep_.CoverTab[131257]++
														nWritten = n
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:110
				// _ = "end of CoverTab[131257]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:111
				_go_fuzz_dep_.CoverTab[131258]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:111
				// _ = "end of CoverTab[131258]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:111
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:111
			// _ = "end of CoverTab[131256]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:111
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:111
		// _ = "end of CoverTab[131253]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:112
	// _ = "end of CoverTab[131251]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:112
	_go_fuzz_dep_.CoverTab[131252]++
											return nWritten, writeErr
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:113
	// _ = "end of CoverTab[131252]"
}

func (ws multiWriteSyncer) Sync() error {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:116
	_go_fuzz_dep_.CoverTab[131259]++
											var err error
											for _, w := range ws {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:118
		_go_fuzz_dep_.CoverTab[131261]++
												err = multierr.Append(err, w.Sync())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:119
		// _ = "end of CoverTab[131261]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:120
	// _ = "end of CoverTab[131259]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:120
	_go_fuzz_dep_.CoverTab[131260]++
											return err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:121
	// _ = "end of CoverTab[131260]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:122
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/zapcore/write_syncer.go:122
var _ = _go_fuzz_dep_.CoverTab
