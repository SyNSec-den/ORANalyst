//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:19
// Package grpcsync implements additional synchronization primitives built upon
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:19
// the sync package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
package grpcsync

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:21
)

import (
	"sync"
	"sync/atomic"
)

// Event represents a one-time event that may occur in the future.
type Event struct {
	fired	int32
	c	chan struct{}
	o	sync.Once
}

// Fire causes e to complete.  It is safe to call multiple times, and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:35
// concurrently.  It returns true iff this call to Fire caused the signaling
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:35
// channel returned by Done to close.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:38
func (e *Event) Fire() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:38
	_go_fuzz_dep_.CoverTab[68894]++
												ret := false
												e.o.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:40
		_go_fuzz_dep_.CoverTab[68896]++
													atomic.StoreInt32(&e.fired, 1)
													close(e.c)
													ret = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:43
		// _ = "end of CoverTab[68896]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:44
	// _ = "end of CoverTab[68894]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:44
	_go_fuzz_dep_.CoverTab[68895]++
												return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:45
	// _ = "end of CoverTab[68895]"
}

// Done returns a channel that will be closed when Fire is called.
func (e *Event) Done() <-chan struct{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:49
	_go_fuzz_dep_.CoverTab[68897]++
												return e.c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:50
	// _ = "end of CoverTab[68897]"
}

// HasFired returns true if Fire has been called.
func (e *Event) HasFired() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:54
	_go_fuzz_dep_.CoverTab[68898]++
												return atomic.LoadInt32(&e.fired) == 1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:55
	// _ = "end of CoverTab[68898]"
}

// NewEvent returns a new, ready-to-use Event.
func NewEvent() *Event {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:59
	_go_fuzz_dep_.CoverTab[68899]++
												return &Event{c: make(chan struct{})}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:60
	// _ = "end of CoverTab[68899]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:61
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcsync/event.go:61
var _ = _go_fuzz_dep_.CoverTab
