//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:1
// Package breaker implements the circuit-breaker resiliency pattern for Go.
package breaker

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:2
)

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// ErrBreakerOpen is the error returned from Run() when the function is not executed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:11
// because the breaker is currently open.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:13
var ErrBreakerOpen = errors.New("circuit breaker is open")

const (
	closed	uint32	= iota
	open
	halfOpen
)

// Breaker implements the circuit-breaker resiliency pattern
type Breaker struct {
	errorThreshold, successThreshold	int
	timeout					time.Duration

	lock			sync.Mutex
	state			uint32
	errors, successes	int
	lastError		time.Time
}

// New constructs a new circuit-breaker that starts closed.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:32
// From closed, the breaker opens if "errorThreshold" errors are seen
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:32
// without an error-free period of at least "timeout". From open, the
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:32
// breaker half-closes after "timeout". From half-open, the breaker closes
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:32
// after "successThreshold" consecutive successes, or opens on a single error.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:37
func New(errorThreshold, successThreshold int, timeout time.Duration) *Breaker {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:37
	_go_fuzz_dep_.CoverTab[81989]++
													return &Breaker{
		errorThreshold:		errorThreshold,
		successThreshold:	successThreshold,
		timeout:		timeout,
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:42
	// _ = "end of CoverTab[81989]"
}

// Run will either return ErrBreakerOpen immediately if the circuit-breaker is
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:45
// already open, or it will run the given function and pass along its return
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:45
// value. It is safe to call Run concurrently on the same Breaker.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:48
func (b *Breaker) Run(work func() error) error {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:48
	_go_fuzz_dep_.CoverTab[81990]++
													state := atomic.LoadUint32(&b.state)

													if state == open {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:51
		_go_fuzz_dep_.CoverTab[81992]++
														return ErrBreakerOpen
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:52
		// _ = "end of CoverTab[81992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:53
		_go_fuzz_dep_.CoverTab[81993]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:53
		// _ = "end of CoverTab[81993]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:53
	// _ = "end of CoverTab[81990]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:53
	_go_fuzz_dep_.CoverTab[81991]++

													return b.doWork(state, work)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:55
	// _ = "end of CoverTab[81991]"
}

// Go will either return ErrBreakerOpen immediately if the circuit-breaker is
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:58
// already open, or it will run the given function in a separate goroutine.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:58
// If the function is run, Go will return nil immediately, and will *not* return
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:58
// the return value of the function. It is safe to call Go concurrently on the
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:58
// same Breaker.
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:63
func (b *Breaker) Go(work func() error) error {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:63
	_go_fuzz_dep_.CoverTab[81994]++
													state := atomic.LoadUint32(&b.state)

													if state == open {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:66
		_go_fuzz_dep_.CoverTab[81996]++
														return ErrBreakerOpen
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:67
		// _ = "end of CoverTab[81996]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
		_go_fuzz_dep_.CoverTab[81997]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
		// _ = "end of CoverTab[81997]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
	// _ = "end of CoverTab[81994]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
	_go_fuzz_dep_.CoverTab[81995]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
	_curRoutineNum99_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:68
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum99_)

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:73
	go b.doWork(state, work)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:75
	// _ = "end of CoverTab[81995]"
}

func (b *Breaker) doWork(state uint32, work func() error) error {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:78
	_go_fuzz_dep_.CoverTab[81998]++
													var panicValue interface{}

													result := func() error {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:81
		_go_fuzz_dep_.CoverTab[82002]++
														defer func() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:82
			_go_fuzz_dep_.CoverTab[82004]++
															panicValue = recover()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:83
			// _ = "end of CoverTab[82004]"
		}()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:84
		// _ = "end of CoverTab[82002]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:84
		_go_fuzz_dep_.CoverTab[82003]++
														return work()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:85
		// _ = "end of CoverTab[82003]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:86
	// _ = "end of CoverTab[81998]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:86
	_go_fuzz_dep_.CoverTab[81999]++

													if result == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		_go_fuzz_dep_.CoverTab[82005]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		return panicValue == nil
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		// _ = "end of CoverTab[82005]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		_go_fuzz_dep_.CoverTab[82006]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		return state == closed
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		// _ = "end of CoverTab[82006]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
	}() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:88
		_go_fuzz_dep_.CoverTab[82007]++

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:91
		return nil
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:91
		// _ = "end of CoverTab[82007]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:92
		_go_fuzz_dep_.CoverTab[82008]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:92
		// _ = "end of CoverTab[82008]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:92
	// _ = "end of CoverTab[81999]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:92
	_go_fuzz_dep_.CoverTab[82000]++

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:95
	b.processResult(result, panicValue)

	if panicValue != nil {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:97
		_go_fuzz_dep_.CoverTab[82009]++

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:100
		panic(panicValue)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:100
		// _ = "end of CoverTab[82009]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:101
		_go_fuzz_dep_.CoverTab[82010]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:101
		// _ = "end of CoverTab[82010]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:101
	// _ = "end of CoverTab[82000]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:101
	_go_fuzz_dep_.CoverTab[82001]++

													return result
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:103
	// _ = "end of CoverTab[82001]"
}

func (b *Breaker) processResult(result error, panicValue interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:106
	_go_fuzz_dep_.CoverTab[82011]++
													b.lock.Lock()
													defer b.lock.Unlock()

													if result == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:110
		_go_fuzz_dep_.CoverTab[82012]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:110
		return panicValue == nil
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:110
		// _ = "end of CoverTab[82012]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:110
	}() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:110
		_go_fuzz_dep_.CoverTab[82013]++
														if b.state == halfOpen {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:111
			_go_fuzz_dep_.CoverTab[82014]++
															b.successes++
															if b.successes == b.successThreshold {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:113
				_go_fuzz_dep_.CoverTab[82015]++
																b.closeBreaker()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:114
				// _ = "end of CoverTab[82015]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:115
				_go_fuzz_dep_.CoverTab[82016]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:115
				// _ = "end of CoverTab[82016]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:115
			}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:115
			// _ = "end of CoverTab[82014]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:116
			_go_fuzz_dep_.CoverTab[82017]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:116
			// _ = "end of CoverTab[82017]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:116
		// _ = "end of CoverTab[82013]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:117
		_go_fuzz_dep_.CoverTab[82018]++
														if b.errors > 0 {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:118
			_go_fuzz_dep_.CoverTab[82020]++
															expiry := b.lastError.Add(b.timeout)
															if time.Now().After(expiry) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:120
				_go_fuzz_dep_.CoverTab[82021]++
																b.errors = 0
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:121
				// _ = "end of CoverTab[82021]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:122
				_go_fuzz_dep_.CoverTab[82022]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:122
				// _ = "end of CoverTab[82022]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:122
			}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:122
			// _ = "end of CoverTab[82020]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:123
			_go_fuzz_dep_.CoverTab[82023]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:123
			// _ = "end of CoverTab[82023]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:123
		// _ = "end of CoverTab[82018]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:123
		_go_fuzz_dep_.CoverTab[82019]++

														switch b.state {
		case closed:
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:126
			_go_fuzz_dep_.CoverTab[82024]++
															b.errors++
															if b.errors == b.errorThreshold {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:128
				_go_fuzz_dep_.CoverTab[82027]++
																b.openBreaker()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:129
				// _ = "end of CoverTab[82027]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:130
				_go_fuzz_dep_.CoverTab[82028]++
																b.lastError = time.Now()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:131
				// _ = "end of CoverTab[82028]"
			}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:132
			// _ = "end of CoverTab[82024]"
		case halfOpen:
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:133
			_go_fuzz_dep_.CoverTab[82025]++
															b.openBreaker()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:134
			// _ = "end of CoverTab[82025]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:134
		default:
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:134
			_go_fuzz_dep_.CoverTab[82026]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:134
			// _ = "end of CoverTab[82026]"
		}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:135
		// _ = "end of CoverTab[82019]"
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:136
	// _ = "end of CoverTab[82011]"
}

func (b *Breaker) openBreaker() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:139
	_go_fuzz_dep_.CoverTab[82029]++
													b.changeState(open)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:140
	_curRoutineNum100_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:140
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum100_)
													go b.timer()
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:141
	// _ = "end of CoverTab[82029]"
}

func (b *Breaker) closeBreaker() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:144
	_go_fuzz_dep_.CoverTab[82030]++
													b.changeState(closed)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:145
	// _ = "end of CoverTab[82030]"
}

func (b *Breaker) timer() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:148
	_go_fuzz_dep_.CoverTab[82031]++
													time.Sleep(b.timeout)

													b.lock.Lock()
													defer b.lock.Unlock()

													b.changeState(halfOpen)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:154
	// _ = "end of CoverTab[82031]"
}

func (b *Breaker) changeState(newState uint32) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:157
	_go_fuzz_dep_.CoverTab[82032]++
													b.errors = 0
													b.successes = 0
													atomic.StoreUint32(&b.state, newState)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:160
	// _ = "end of CoverTab[82032]"
}

//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:161
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/eapache/go-resiliency@v1.2.0/breaker/breaker.go:161
var _ = _go_fuzz_dep_.CoverTab
