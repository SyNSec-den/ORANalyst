//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
package backoff

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:1
)

import (
	"context"
	"time"
)

// BackOffContext is a backoff policy that stops retrying after the context
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:8
// is canceled.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:10
type BackOffContext interface {
	BackOff
	Context() context.Context
}

type backOffContext struct {
	BackOff
	ctx	context.Context
}

// WithContext returns a BackOffContext with context ctx
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:20
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:20
// ctx must not be nil
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:23
func WithContext(b BackOff, ctx context.Context) BackOffContext {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:23
	_go_fuzz_dep_.CoverTab[182670]++
													if ctx == nil {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:24
		_go_fuzz_dep_.CoverTab[182673]++
														panic("nil context")
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:25
		// _ = "end of CoverTab[182673]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:26
		_go_fuzz_dep_.CoverTab[182674]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:26
		// _ = "end of CoverTab[182674]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:26
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:26
	// _ = "end of CoverTab[182670]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:26
	_go_fuzz_dep_.CoverTab[182671]++

													if b, ok := b.(*backOffContext); ok {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:28
		_go_fuzz_dep_.CoverTab[182675]++
														return &backOffContext{
			BackOff:	b.BackOff,
			ctx:		ctx,
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:32
		// _ = "end of CoverTab[182675]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:33
		_go_fuzz_dep_.CoverTab[182676]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:33
		// _ = "end of CoverTab[182676]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:33
	// _ = "end of CoverTab[182671]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:33
	_go_fuzz_dep_.CoverTab[182672]++

													return &backOffContext{
		BackOff:	b,
		ctx:		ctx,
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:38
	// _ = "end of CoverTab[182672]"
}

func ensureContext(b BackOff) BackOffContext {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:41
	_go_fuzz_dep_.CoverTab[182677]++
													if cb, ok := b.(BackOffContext); ok {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:42
		_go_fuzz_dep_.CoverTab[182679]++
														return cb
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:43
		// _ = "end of CoverTab[182679]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:44
		_go_fuzz_dep_.CoverTab[182680]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:44
		// _ = "end of CoverTab[182680]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:44
	// _ = "end of CoverTab[182677]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:44
	_go_fuzz_dep_.CoverTab[182678]++
													return WithContext(b, context.Background())
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:45
	// _ = "end of CoverTab[182678]"
}

func (b *backOffContext) Context() context.Context {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:48
	_go_fuzz_dep_.CoverTab[182681]++
													return b.ctx
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:49
	// _ = "end of CoverTab[182681]"
}

func (b *backOffContext) NextBackOff() time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:52
	_go_fuzz_dep_.CoverTab[182682]++
													select {
	case <-b.ctx.Done():
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:54
		_go_fuzz_dep_.CoverTab[182685]++
														return Stop
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:55
		// _ = "end of CoverTab[182685]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:56
		_go_fuzz_dep_.CoverTab[182686]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:56
		// _ = "end of CoverTab[182686]"
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:57
	// _ = "end of CoverTab[182682]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:57
	_go_fuzz_dep_.CoverTab[182683]++
													next := b.BackOff.NextBackOff()
													if deadline, ok := b.ctx.Deadline(); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:59
		_go_fuzz_dep_.CoverTab[182687]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:59
		return deadline.Sub(time.Now()) < next
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:59
		// _ = "end of CoverTab[182687]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:59
	}() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:59
		_go_fuzz_dep_.CoverTab[182688]++
														return Stop
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:60
		// _ = "end of CoverTab[182688]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:61
		_go_fuzz_dep_.CoverTab[182689]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:61
		// _ = "end of CoverTab[182689]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:61
	// _ = "end of CoverTab[182683]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:61
	_go_fuzz_dep_.CoverTab[182684]++
													return next
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:62
	// _ = "end of CoverTab[182684]"
}

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/context.go:63
var _ = _go_fuzz_dep_.CoverTab
