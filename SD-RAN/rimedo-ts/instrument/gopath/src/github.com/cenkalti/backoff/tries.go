//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
package backoff

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:1
)

import "time"

/*
WithMaxRetries creates a wrapper around another BackOff, which will
return Stop if NextBackOff() has been called too many times since
the last time Reset() was called

Note: Implementation is not thread-safe.
*/
func WithMaxRetries(b BackOff, max uint64) BackOff {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:12
	_go_fuzz_dep_.CoverTab[182743]++
												return &backOffTries{delegate: b, maxTries: max}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:13
	// _ = "end of CoverTab[182743]"
}

type backOffTries struct {
	delegate	BackOff
	maxTries	uint64
	numTries	uint64
}

func (b *backOffTries) NextBackOff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:22
	_go_fuzz_dep_.CoverTab[182744]++
												if b.maxTries > 0 {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:23
		_go_fuzz_dep_.CoverTab[182746]++
													if b.maxTries <= b.numTries {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:24
			_go_fuzz_dep_.CoverTab[182748]++
														return Stop
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:25
			// _ = "end of CoverTab[182748]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:26
			_go_fuzz_dep_.CoverTab[182749]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:26
			// _ = "end of CoverTab[182749]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:26
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:26
		// _ = "end of CoverTab[182746]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:26
		_go_fuzz_dep_.CoverTab[182747]++
													b.numTries++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:27
		// _ = "end of CoverTab[182747]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:28
		_go_fuzz_dep_.CoverTab[182750]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:28
		// _ = "end of CoverTab[182750]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:28
	// _ = "end of CoverTab[182744]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:28
	_go_fuzz_dep_.CoverTab[182745]++
												return b.delegate.NextBackOff()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:29
	// _ = "end of CoverTab[182745]"
}

func (b *backOffTries) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:32
	_go_fuzz_dep_.CoverTab[182751]++
												b.numTries = 0
												b.delegate.Reset()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:34
	// _ = "end of CoverTab[182751]"
}

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:35
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/tries.go:35
var _ = _go_fuzz_dep_.CoverTab
