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

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:21
)

import (
	"fmt"
	"io"
	"io/ioutil"

	"go.uber.org/zap/zapcore"

	"go.uber.org/multierr"
)

// Open is a high-level wrapper that takes a variadic number of URLs, opens or
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// creates each of the specified resources, and combines them into a locked
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// WriteSyncer. It also returns any error encountered and a function to close
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// any opened files.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// Passing no URLs returns a no-op WriteSyncer. Zap handles URLs without a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// scheme and URLs with the "file" scheme. Third-party code may register
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// factories for other schemes using RegisterSink.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// URLs with the "file" scheme must use absolute paths on the local
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// filesystem. No user, password, port, fragments, or query parameters are
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// allowed, and the hostname must be empty or "localhost".
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// Since it's common to write logs to the local filesystem, URLs without a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// scheme (e.g., "/var/log/foo.log") are treated as local file paths. Without
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// a scheme, the special paths "stdout" and "stderr" are interpreted as
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// os.Stdout and os.Stderr. When specified without a scheme, relative file
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:33
// paths also work.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:51
func Open(paths ...string) (zapcore.WriteSyncer, func(), error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:51
	_go_fuzz_dep_.CoverTab[131923]++
									writers, close, err := open(paths)
									if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:53
		_go_fuzz_dep_.CoverTab[131925]++
										return nil, nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:54
		// _ = "end of CoverTab[131925]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:55
		_go_fuzz_dep_.CoverTab[131926]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:55
		// _ = "end of CoverTab[131926]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:55
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:55
	// _ = "end of CoverTab[131923]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:55
	_go_fuzz_dep_.CoverTab[131924]++

									writer := CombineWriteSyncers(writers...)
									return writer, close, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:58
	// _ = "end of CoverTab[131924]"
}

func open(paths []string) ([]zapcore.WriteSyncer, func(), error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:61
	_go_fuzz_dep_.CoverTab[131927]++
									writers := make([]zapcore.WriteSyncer, 0, len(paths))
									closers := make([]io.Closer, 0, len(paths))
									close := func() {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:64
		_go_fuzz_dep_.CoverTab[131931]++
										for _, c := range closers {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:65
			_go_fuzz_dep_.CoverTab[131932]++
											c.Close()
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:66
			// _ = "end of CoverTab[131932]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:67
		// _ = "end of CoverTab[131931]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:68
	// _ = "end of CoverTab[131927]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:68
	_go_fuzz_dep_.CoverTab[131928]++

									var openErr error
									for _, path := range paths {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:71
		_go_fuzz_dep_.CoverTab[131933]++
										sink, err := newSink(path)
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:73
			_go_fuzz_dep_.CoverTab[131935]++
											openErr = multierr.Append(openErr, fmt.Errorf("couldn't open sink %q: %v", path, err))
											continue
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:75
			// _ = "end of CoverTab[131935]"
		} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:76
			_go_fuzz_dep_.CoverTab[131936]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:76
			// _ = "end of CoverTab[131936]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:76
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:76
		// _ = "end of CoverTab[131933]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:76
		_go_fuzz_dep_.CoverTab[131934]++
										writers = append(writers, sink)
										closers = append(closers, sink)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:78
		// _ = "end of CoverTab[131934]"
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:79
	// _ = "end of CoverTab[131928]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:79
	_go_fuzz_dep_.CoverTab[131929]++
									if openErr != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:80
		_go_fuzz_dep_.CoverTab[131937]++
										close()
										return writers, nil, openErr
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:82
		// _ = "end of CoverTab[131937]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:83
		_go_fuzz_dep_.CoverTab[131938]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:83
		// _ = "end of CoverTab[131938]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:83
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:83
	// _ = "end of CoverTab[131929]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:83
	_go_fuzz_dep_.CoverTab[131930]++

									return writers, close, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:85
	// _ = "end of CoverTab[131930]"
}

// CombineWriteSyncers is a utility that combines multiple WriteSyncers into a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:88
// single, locked WriteSyncer. If no inputs are supplied, it returns a no-op
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:88
// WriteSyncer.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:88
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:88
// It's provided purely as a convenience; the result is no different from
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:88
// using zapcore.NewMultiWriteSyncer and zapcore.Lock individually.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:94
func CombineWriteSyncers(writers ...zapcore.WriteSyncer) zapcore.WriteSyncer {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:94
	_go_fuzz_dep_.CoverTab[131939]++
									if len(writers) == 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:95
		_go_fuzz_dep_.CoverTab[131941]++
										return zapcore.AddSync(ioutil.Discard)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:96
		// _ = "end of CoverTab[131941]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:97
		_go_fuzz_dep_.CoverTab[131942]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:97
		// _ = "end of CoverTab[131942]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:97
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:97
	// _ = "end of CoverTab[131939]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:97
	_go_fuzz_dep_.CoverTab[131940]++
									return zapcore.Lock(zapcore.NewMultiWriteSyncer(writers...))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:98
	// _ = "end of CoverTab[131940]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:99
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/writer.go:99
var _ = _go_fuzz_dep_.CoverTab
