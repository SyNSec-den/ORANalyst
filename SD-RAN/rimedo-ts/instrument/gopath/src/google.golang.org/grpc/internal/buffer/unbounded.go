//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:18
// Package buffer provides an implementation of an unbounded buffer.
package buffer

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:19
)

import "sync"

// Unbounded is an implementation of an unbounded buffer which does not use
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// extra goroutines. This is typically used for passing updates from one entity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// to another within gRPC.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// All methods on this type are thread-safe and don't block on anything except
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// the underlying mutex used for synchronization.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// Unbounded supports values of any type to be stored in it by using a channel
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// of `interface{}`. This means that a call to Put() incurs an extra memory
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// allocation, and also that users need a type assertion while reading. For
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// performance critical code paths, using Unbounded is strongly discouraged and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// defining a new type specific implementation of this buffer is preferred. See
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:23
// internal/transport/transport.go for an example of this.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:36
type Unbounded struct {
	c	chan interface{}
	mu	sync.Mutex
	backlog	[]interface{}
}

// NewUnbounded returns a new instance of Unbounded.
func NewUnbounded() *Unbounded {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:43
	_go_fuzz_dep_.CoverTab[68880]++
													return &Unbounded{c: make(chan interface{}, 1)}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:44
	// _ = "end of CoverTab[68880]"
}

// Put adds t to the unbounded buffer.
func (b *Unbounded) Put(t interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:48
	_go_fuzz_dep_.CoverTab[68881]++
													b.mu.Lock()
													if len(b.backlog) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:50
		_go_fuzz_dep_.CoverTab[68883]++
														select {
		case b.c <- t:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:52
			_go_fuzz_dep_.CoverTab[68884]++
															b.mu.Unlock()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:54
			// _ = "end of CoverTab[68884]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:55
			_go_fuzz_dep_.CoverTab[68885]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:55
			// _ = "end of CoverTab[68885]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:56
		// _ = "end of CoverTab[68883]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:57
		_go_fuzz_dep_.CoverTab[68886]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:57
		// _ = "end of CoverTab[68886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:57
	// _ = "end of CoverTab[68881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:57
	_go_fuzz_dep_.CoverTab[68882]++
													b.backlog = append(b.backlog, t)
													b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:59
	// _ = "end of CoverTab[68882]"
}

// Load sends the earliest buffered data, if any, onto the read channel
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:62
// returned by Get(). Users are expected to call this every time they read a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:62
// value from the read channel.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:65
func (b *Unbounded) Load() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:65
	_go_fuzz_dep_.CoverTab[68887]++
													b.mu.Lock()
													if len(b.backlog) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:67
		_go_fuzz_dep_.CoverTab[68889]++
														select {
		case b.c <- b.backlog[0]:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:69
			_go_fuzz_dep_.CoverTab[68890]++
															b.backlog[0] = nil
															b.backlog = b.backlog[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:71
			// _ = "end of CoverTab[68890]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:72
			_go_fuzz_dep_.CoverTab[68891]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:72
			// _ = "end of CoverTab[68891]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:73
		// _ = "end of CoverTab[68889]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:74
		_go_fuzz_dep_.CoverTab[68892]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:74
		// _ = "end of CoverTab[68892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:74
	// _ = "end of CoverTab[68887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:74
	_go_fuzz_dep_.CoverTab[68888]++
													b.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:75
	// _ = "end of CoverTab[68888]"
}

// Get returns a read channel on which values added to the buffer, via Put(),
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:78
// are sent on.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:78
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:78
// Upon reading a value from this channel, users are expected to call Load() to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:78
// send the next buffered value onto the channel if there is any.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:83
func (b *Unbounded) Get() <-chan interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:83
	_go_fuzz_dep_.CoverTab[68893]++
													return b.c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:84
	// _ = "end of CoverTab[68893]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:85
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/buffer/unbounded.go:85
var _ = _go_fuzz_dep_.CoverTab
