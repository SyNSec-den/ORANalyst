//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:19
)

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

// writeQuota is a soft limit on the amount of data a stream can
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:28
// schedule before some of it is written out.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:30
type writeQuota struct {
	quota	int32
	// get waits on read from when quota goes less than or equal to zero.
	// replenish writes on it when quota goes positive again.
	ch	chan struct{}
	// done is triggered in error case.
	done	<-chan struct{}
	// replenish is called by loopyWriter to give quota back to.
	// It is implemented as a field so that it can be updated
	// by tests.
	replenish	func(n int)
}

func newWriteQuota(sz int32, done <-chan struct{}) *writeQuota {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:43
	_go_fuzz_dep_.CoverTab[76950]++
													w := &writeQuota{
		quota:	sz,
		ch:	make(chan struct{}, 1),
		done:	done,
	}
													w.replenish = w.realReplenish
													return w
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:50
	// _ = "end of CoverTab[76950]"
}

func (w *writeQuota) get(sz int32) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:53
	_go_fuzz_dep_.CoverTab[76951]++
													for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:54
		_go_fuzz_dep_.CoverTab[76952]++
														if atomic.LoadInt32(&w.quota) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:55
			_go_fuzz_dep_.CoverTab[76954]++
															atomic.AddInt32(&w.quota, -sz)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:57
			// _ = "end of CoverTab[76954]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:58
			_go_fuzz_dep_.CoverTab[76955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:58
			// _ = "end of CoverTab[76955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:58
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:58
		// _ = "end of CoverTab[76952]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:58
		_go_fuzz_dep_.CoverTab[76953]++
														select {
		case <-w.ch:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:60
			_go_fuzz_dep_.CoverTab[76956]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:61
			// _ = "end of CoverTab[76956]"
		case <-w.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:62
			_go_fuzz_dep_.CoverTab[76957]++
															return errStreamDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:63
			// _ = "end of CoverTab[76957]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:64
		// _ = "end of CoverTab[76953]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:65
	// _ = "end of CoverTab[76951]"
}

func (w *writeQuota) realReplenish(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:68
	_go_fuzz_dep_.CoverTab[76958]++
													sz := int32(n)
													a := atomic.AddInt32(&w.quota, sz)
													b := a - sz
													if b <= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:72
		_go_fuzz_dep_.CoverTab[76959]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:72
		return a > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:72
		// _ = "end of CoverTab[76959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:72
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:72
		_go_fuzz_dep_.CoverTab[76960]++
														select {
		case w.ch <- struct{}{}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:74
			_go_fuzz_dep_.CoverTab[76961]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:74
			// _ = "end of CoverTab[76961]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:75
			_go_fuzz_dep_.CoverTab[76962]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:75
			// _ = "end of CoverTab[76962]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:76
		// _ = "end of CoverTab[76960]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:77
		_go_fuzz_dep_.CoverTab[76963]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:77
		// _ = "end of CoverTab[76963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:77
	// _ = "end of CoverTab[76958]"
}

type trInFlow struct {
	limit			uint32
	unacked			uint32
	effectiveWindowSize	uint32
}

func (f *trInFlow) newLimit(n uint32) uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:86
	_go_fuzz_dep_.CoverTab[76964]++
													d := n - f.limit
													f.limit = n
													f.updateEffectiveWindowSize()
													return d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:90
	// _ = "end of CoverTab[76964]"
}

func (f *trInFlow) onData(n uint32) uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:93
	_go_fuzz_dep_.CoverTab[76965]++
													f.unacked += n
													if f.unacked >= f.limit/4 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:95
		_go_fuzz_dep_.CoverTab[76967]++
														w := f.unacked
														f.unacked = 0
														f.updateEffectiveWindowSize()
														return w
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:99
		// _ = "end of CoverTab[76967]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:100
		_go_fuzz_dep_.CoverTab[76968]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:100
		// _ = "end of CoverTab[76968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:100
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:100
	// _ = "end of CoverTab[76965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:100
	_go_fuzz_dep_.CoverTab[76966]++
													f.updateEffectiveWindowSize()
													return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:102
	// _ = "end of CoverTab[76966]"
}

func (f *trInFlow) reset() uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:105
	_go_fuzz_dep_.CoverTab[76969]++
													w := f.unacked
													f.unacked = 0
													f.updateEffectiveWindowSize()
													return w
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:109
	// _ = "end of CoverTab[76969]"
}

func (f *trInFlow) updateEffectiveWindowSize() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:112
	_go_fuzz_dep_.CoverTab[76970]++
													atomic.StoreUint32(&f.effectiveWindowSize, f.limit-f.unacked)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:113
	// _ = "end of CoverTab[76970]"
}

func (f *trInFlow) getSize() uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:116
	_go_fuzz_dep_.CoverTab[76971]++
													return atomic.LoadUint32(&f.effectiveWindowSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:117
	// _ = "end of CoverTab[76971]"
}

// TODO(mmukhi): Simplify this code.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:120
// inFlow deals with inbound flow control
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:122
type inFlow struct {
	mu	sync.Mutex
	// The inbound flow control limit for pending data.
	limit	uint32
	// pendingData is the overall data which have been received but not been
	// consumed by applications.
	pendingData	uint32
	// The amount of data the application has consumed but grpc has not sent
	// window update for them. Used to reduce window update frequency.
	pendingUpdate	uint32
	// delta is the extra window update given by receiver when an application
	// is reading data bigger in size than the inFlow limit.
	delta	uint32
}

// newLimit updates the inflow window to a new value n.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:137
// It assumes that n is always greater than the old limit.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:139
func (f *inFlow) newLimit(n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:139
	_go_fuzz_dep_.CoverTab[76972]++
													f.mu.Lock()
													f.limit = n
													f.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:142
	// _ = "end of CoverTab[76972]"
}

func (f *inFlow) maybeAdjust(n uint32) uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:145
	_go_fuzz_dep_.CoverTab[76973]++
													if n > uint32(math.MaxInt32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:146
		_go_fuzz_dep_.CoverTab[76976]++
														n = uint32(math.MaxInt32)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:147
		// _ = "end of CoverTab[76976]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:148
		_go_fuzz_dep_.CoverTab[76977]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:148
		// _ = "end of CoverTab[76977]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:148
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:148
	// _ = "end of CoverTab[76973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:148
	_go_fuzz_dep_.CoverTab[76974]++
													f.mu.Lock()
													defer f.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:153
	estSenderQuota := int32(f.limit - (f.pendingData + f.pendingUpdate))

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:157
	estUntransmittedData := int32(n - f.pendingData)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:161
	if estUntransmittedData > estSenderQuota {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:161
		_go_fuzz_dep_.CoverTab[76978]++

														if f.limit+n > maxWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:163
			_go_fuzz_dep_.CoverTab[76980]++
															f.delta = maxWindowSize - f.limit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:164
			// _ = "end of CoverTab[76980]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:165
			_go_fuzz_dep_.CoverTab[76981]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:169
			f.delta = n
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:169
			// _ = "end of CoverTab[76981]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:170
		// _ = "end of CoverTab[76978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:170
		_go_fuzz_dep_.CoverTab[76979]++
														return f.delta
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:171
		// _ = "end of CoverTab[76979]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:172
		_go_fuzz_dep_.CoverTab[76982]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:172
		// _ = "end of CoverTab[76982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:172
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:172
	// _ = "end of CoverTab[76974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:172
	_go_fuzz_dep_.CoverTab[76975]++
													return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:173
	// _ = "end of CoverTab[76975]"
}

// onData is invoked when some data frame is received. It updates pendingData.
func (f *inFlow) onData(n uint32) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:177
	_go_fuzz_dep_.CoverTab[76983]++
													f.mu.Lock()
													f.pendingData += n
													if f.pendingData+f.pendingUpdate > f.limit+f.delta {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:180
		_go_fuzz_dep_.CoverTab[76985]++
														limit := f.limit
														rcvd := f.pendingData + f.pendingUpdate
														f.mu.Unlock()
														return fmt.Errorf("received %d-bytes data exceeding the limit %d bytes", rcvd, limit)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:184
		// _ = "end of CoverTab[76985]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:185
		_go_fuzz_dep_.CoverTab[76986]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:185
		// _ = "end of CoverTab[76986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:185
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:185
	// _ = "end of CoverTab[76983]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:185
	_go_fuzz_dep_.CoverTab[76984]++
													f.mu.Unlock()
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:187
	// _ = "end of CoverTab[76984]"
}

// onRead is invoked when the application reads the data. It returns the window size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:190
// to be sent to the peer.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:192
func (f *inFlow) onRead(n uint32) uint32 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:192
	_go_fuzz_dep_.CoverTab[76987]++
													f.mu.Lock()
													if f.pendingData == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:194
		_go_fuzz_dep_.CoverTab[76991]++
														f.mu.Unlock()
														return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:196
		// _ = "end of CoverTab[76991]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:197
		_go_fuzz_dep_.CoverTab[76992]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:197
		// _ = "end of CoverTab[76992]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:197
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:197
	// _ = "end of CoverTab[76987]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:197
	_go_fuzz_dep_.CoverTab[76988]++
													f.pendingData -= n
													if n > f.delta {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:199
		_go_fuzz_dep_.CoverTab[76993]++
														n -= f.delta
														f.delta = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:201
		// _ = "end of CoverTab[76993]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:202
		_go_fuzz_dep_.CoverTab[76994]++
														f.delta -= n
														n = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:204
		// _ = "end of CoverTab[76994]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:205
	// _ = "end of CoverTab[76988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:205
	_go_fuzz_dep_.CoverTab[76989]++
													f.pendingUpdate += n
													if f.pendingUpdate >= f.limit/4 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:207
		_go_fuzz_dep_.CoverTab[76995]++
														wu := f.pendingUpdate
														f.pendingUpdate = 0
														f.mu.Unlock()
														return wu
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:211
		// _ = "end of CoverTab[76995]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:212
		_go_fuzz_dep_.CoverTab[76996]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:212
		// _ = "end of CoverTab[76996]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:212
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:212
	// _ = "end of CoverTab[76989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:212
	_go_fuzz_dep_.CoverTab[76990]++
													f.mu.Unlock()
													return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:214
	// _ = "end of CoverTab[76990]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:215
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/flowcontrol.go:215
var _ = _go_fuzz_dep_.CoverTab
