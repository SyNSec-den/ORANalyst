//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
package backoff

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:1
)

import (
	"sync"
	"time"
)

// Ticker holds a channel that delivers `ticks' of a clock at times reported by a BackOff.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:8
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:8
// Ticks will continue to arrive when the previous operation is still running,
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:8
// so operations that take a while to fail could run in quick succession.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:12
type Ticker struct {
	C		<-chan time.Time
	c		chan time.Time
	b		BackOffContext
	stop		chan struct{}
	stopOnce	sync.Once
}

// NewTicker returns a new Ticker containing a channel that will send
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:20
// the time at times specified by the BackOff argument. Ticker is
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:20
// guaranteed to tick at least once.  The channel is closed when Stop
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:20
// method is called or BackOff stops. It is not safe to manipulate the
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:20
// provided backoff policy (notably calling NextBackOff or Reset)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:20
// while the ticker is running.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:26
func NewTicker(b BackOff) *Ticker {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:26
	_go_fuzz_dep_.CoverTab[182725]++
												c := make(chan time.Time)
												t := &Ticker{
		C:	c,
		c:	c,
		b:	ensureContext(b),
		stop:	make(chan struct{}),
	}
												t.b.Reset()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:34
	_curRoutineNum162_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:34
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum162_)
												go t.run()
												return t
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:36
	// _ = "end of CoverTab[182725]"
}

// Stop turns off a ticker. After Stop, no more ticks will be sent.
func (t *Ticker) Stop() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:40
	_go_fuzz_dep_.CoverTab[182726]++
												t.stopOnce.Do(func() { _go_fuzz_dep_.CoverTab[182727]++; close(t.stop); // _ = "end of CoverTab[182727]" })
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:41
	// _ = "end of CoverTab[182726]"
}

func (t *Ticker) run() {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:44
	_go_fuzz_dep_.CoverTab[182728]++
												c := t.c
												defer close(c)

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:49
	afterC := t.send(time.Now())

	for {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:51
		_go_fuzz_dep_.CoverTab[182729]++
													if afterC == nil {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:52
			_go_fuzz_dep_.CoverTab[182731]++
														return
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:53
			// _ = "end of CoverTab[182731]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:54
			_go_fuzz_dep_.CoverTab[182732]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:54
			// _ = "end of CoverTab[182732]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:54
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:54
		// _ = "end of CoverTab[182729]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:54
		_go_fuzz_dep_.CoverTab[182730]++

													select {
		case tick := <-afterC:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:57
			_go_fuzz_dep_.CoverTab[182733]++
														afterC = t.send(tick)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:58
			// _ = "end of CoverTab[182733]"
		case <-t.stop:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:59
			_go_fuzz_dep_.CoverTab[182734]++
														t.c = nil
														return
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:61
			// _ = "end of CoverTab[182734]"
		case <-t.b.Context().Done():
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:62
			_go_fuzz_dep_.CoverTab[182735]++
														return
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:63
			// _ = "end of CoverTab[182735]"
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:64
		// _ = "end of CoverTab[182730]"
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:65
	// _ = "end of CoverTab[182728]"
}

func (t *Ticker) send(tick time.Time) <-chan time.Time {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:68
	_go_fuzz_dep_.CoverTab[182736]++
												select {
	case t.c <- tick:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:70
		_go_fuzz_dep_.CoverTab[182739]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:70
		// _ = "end of CoverTab[182739]"
	case <-t.stop:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:71
		_go_fuzz_dep_.CoverTab[182740]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:72
		// _ = "end of CoverTab[182740]"
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:73
	// _ = "end of CoverTab[182736]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:73
	_go_fuzz_dep_.CoverTab[182737]++

												next := t.b.NextBackOff()
												if next == Stop {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:76
		_go_fuzz_dep_.CoverTab[182741]++
													t.Stop()
													return nil
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:78
		// _ = "end of CoverTab[182741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:79
		_go_fuzz_dep_.CoverTab[182742]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:79
		// _ = "end of CoverTab[182742]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:79
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:79
	// _ = "end of CoverTab[182737]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:79
	_go_fuzz_dep_.CoverTab[182738]++

												return time.After(next)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:81
	// _ = "end of CoverTab[182738]"
}

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:82
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/ticker.go:82
var _ = _go_fuzz_dep_.CoverTab
