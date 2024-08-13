// Copyright © 2016 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
package jwalterweatherman

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:6
)

import (
	"io"
	"sync/atomic"
)

// Counter is an io.Writer that increments a counter on Write.
type Counter struct {
	count uint64
}

func (c *Counter) incr() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:18
	_go_fuzz_dep_.CoverTab[119222]++
												atomic.AddUint64(&c.count, 1)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:19
	// _ = "end of CoverTab[119222]"
}

// Reset resets the counter.
func (c *Counter) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:23
	_go_fuzz_dep_.CoverTab[119223]++
												atomic.StoreUint64(&c.count, 0)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:24
	// _ = "end of CoverTab[119223]"
}

// Count returns the current count.
func (c *Counter) Count() uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:28
	_go_fuzz_dep_.CoverTab[119224]++
												return atomic.LoadUint64(&c.count)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:29
	// _ = "end of CoverTab[119224]"
}

func (c *Counter) Write(p []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:32
	_go_fuzz_dep_.CoverTab[119225]++
												c.incr()
												return len(p), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:34
	// _ = "end of CoverTab[119225]"
}

// LogCounter creates a LogListener that counts log statements >= the given threshold.
func LogCounter(counter *Counter, t1 Threshold) LogListener {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:38
	_go_fuzz_dep_.CoverTab[119226]++
												return func(t2 Threshold) io.Writer {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:39
		_go_fuzz_dep_.CoverTab[119227]++
													if t2 < t1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:40
			_go_fuzz_dep_.CoverTab[119229]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:42
			// _ = "end of CoverTab[119229]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:43
			_go_fuzz_dep_.CoverTab[119230]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:43
			// _ = "end of CoverTab[119230]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:43
		// _ = "end of CoverTab[119227]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:43
		_go_fuzz_dep_.CoverTab[119228]++
													return counter
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:44
		// _ = "end of CoverTab[119228]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:45
	// _ = "end of CoverTab[119226]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/jwalterweatherman@v1.1.0/log_counter.go:46
var _ = _go_fuzz_dep_.CoverTab
