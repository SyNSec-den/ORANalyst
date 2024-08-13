//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
package backoff

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:1
)

import "time"

// An Operation is executing by Retry() or RetryNotify().
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:5
// The operation will be retried using a backoff policy if it returns an error.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:7
type Operation func() error

// Notify is a notify-on-error function. It receives an operation error and
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:9
// backoff delay if the operation failed (with an error).
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:9
// NOTE that if the backoff policy stated to stop retrying,
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:9
// the notify function isn't called.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:14
type Notify func(error, time.Duration)

// Retry the operation o until it does not return error or BackOff stops.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
// o is guaranteed to be run at least once.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
// If o returns a *PermanentError, the operation is not retried, and the
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
// wrapped error is returned.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
//
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
// Retry sleeps the goroutine for the duration returned by BackOff after a
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:16
// failed operation returns.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:24
func Retry(o Operation, b BackOff) error {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:24
	_go_fuzz_dep_.CoverTab[182703]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:24
	return RetryNotify(o, b, nil)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:24
	// _ = "end of CoverTab[182703]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:24
}

// RetryNotify calls notify function with the error and wait duration
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:26
// for each failed attempt before sleep.
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:28
func RetryNotify(operation Operation, b BackOff, notify Notify) error {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:28
	_go_fuzz_dep_.CoverTab[182704]++
												var err error
												var next time.Duration
												var t *time.Timer

												cb := ensureContext(b)

												b.Reset()
												for {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:36
		_go_fuzz_dep_.CoverTab[182705]++
													if err = operation(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:37
			_go_fuzz_dep_.CoverTab[182711]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:38
			// _ = "end of CoverTab[182711]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:39
			_go_fuzz_dep_.CoverTab[182712]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:39
			// _ = "end of CoverTab[182712]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:39
		// _ = "end of CoverTab[182705]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:39
		_go_fuzz_dep_.CoverTab[182706]++

													if permanent, ok := err.(*PermanentError); ok {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:41
			_go_fuzz_dep_.CoverTab[182713]++
														return permanent.Err
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:42
			// _ = "end of CoverTab[182713]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:43
			_go_fuzz_dep_.CoverTab[182714]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:43
			// _ = "end of CoverTab[182714]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:43
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:43
		// _ = "end of CoverTab[182706]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:43
		_go_fuzz_dep_.CoverTab[182707]++

													if next = cb.NextBackOff(); next == Stop {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:45
			_go_fuzz_dep_.CoverTab[182715]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:46
			// _ = "end of CoverTab[182715]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:47
			_go_fuzz_dep_.CoverTab[182716]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:47
			// _ = "end of CoverTab[182716]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:47
		// _ = "end of CoverTab[182707]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:47
		_go_fuzz_dep_.CoverTab[182708]++

													if notify != nil {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:49
			_go_fuzz_dep_.CoverTab[182717]++
														notify(err, next)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:50
			// _ = "end of CoverTab[182717]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:51
			_go_fuzz_dep_.CoverTab[182718]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:51
			// _ = "end of CoverTab[182718]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:51
		// _ = "end of CoverTab[182708]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:51
		_go_fuzz_dep_.CoverTab[182709]++

													if t == nil {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:53
			_go_fuzz_dep_.CoverTab[182719]++
														t = time.NewTimer(next)
														defer t.Stop()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:55
			// _ = "end of CoverTab[182719]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:56
			_go_fuzz_dep_.CoverTab[182720]++
														t.Reset(next)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:57
			// _ = "end of CoverTab[182720]"
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:58
		// _ = "end of CoverTab[182709]"
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:58
		_go_fuzz_dep_.CoverTab[182710]++

													select {
		case <-cb.Context().Done():
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:61
			_go_fuzz_dep_.CoverTab[182721]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:62
			// _ = "end of CoverTab[182721]"
		case <-t.C:
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:63
			_go_fuzz_dep_.CoverTab[182722]++
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:63
			// _ = "end of CoverTab[182722]"
		}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:64
		// _ = "end of CoverTab[182710]"
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:65
	// _ = "end of CoverTab[182704]"
}

// PermanentError signals that the operation should not be retried.
type PermanentError struct {
	Err error
}

func (e *PermanentError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:73
	_go_fuzz_dep_.CoverTab[182723]++
												return e.Err.Error()
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:74
	// _ = "end of CoverTab[182723]"
}

// Permanent wraps the given err in a *PermanentError.
func Permanent(err error) *PermanentError {
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:78
	_go_fuzz_dep_.CoverTab[182724]++
												return &PermanentError{
		Err: err,
	}
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:81
	// _ = "end of CoverTab[182724]"
}

//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:82
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/cenkalti/backoff@v2.2.1+incompatible/retry.go:82
var _ = _go_fuzz_dep_.CoverTab
