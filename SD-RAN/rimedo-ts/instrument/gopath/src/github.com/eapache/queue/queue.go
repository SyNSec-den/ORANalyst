//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:1
/*
Package queue provides a fast, ring-buffer queue based on the version suggested by Dariusz Górecki.
Using this instead of other, simpler, queue implementations (slice+append or linked list) provides
substantial memory and time benefits, and fewer GC pauses.

The queue implemented here is as fast as it is for an additional reason: it is *not* thread-safe.
*/
package queue

//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:8
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:8
)
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:8
import (
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:8
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:8
)

// minQueueLen is smallest capacity that queue may have.
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:10
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:12
const minQueueLen = 16

// Queue represents a single instance of the queue data structure.
type Queue struct {
	buf			[]interface{}
	head, tail, count	int
}

// New constructs and returns a new Queue.
func New() *Queue {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:21
	_go_fuzz_dep_.CoverTab[82263]++
										return &Queue{
		buf: make([]interface{}, minQueueLen),
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:24
	// _ = "end of CoverTab[82263]"
}

// Length returns the number of elements currently stored in the queue.
func (q *Queue) Length() int {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:28
	_go_fuzz_dep_.CoverTab[82264]++
										return q.count
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:29
	// _ = "end of CoverTab[82264]"
}

// resizes the queue to fit exactly twice its current contents
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:32
// this can result in shrinking if the queue is less than half-full
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:34
func (q *Queue) resize() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:34
	_go_fuzz_dep_.CoverTab[82265]++
										newBuf := make([]interface{}, q.count<<1)

										if q.tail > q.head {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:37
		_go_fuzz_dep_.CoverTab[82267]++
											copy(newBuf, q.buf[q.head:q.tail])
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:38
		// _ = "end of CoverTab[82267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:39
		_go_fuzz_dep_.CoverTab[82268]++
											n := copy(newBuf, q.buf[q.head:])
											copy(newBuf[n:], q.buf[:q.tail])
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:41
		// _ = "end of CoverTab[82268]"
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:42
	// _ = "end of CoverTab[82265]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:42
	_go_fuzz_dep_.CoverTab[82266]++

										q.head = 0
										q.tail = q.count
										q.buf = newBuf
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:46
	// _ = "end of CoverTab[82266]"
}

// Add puts an element on the end of the queue.
func (q *Queue) Add(elem interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:50
	_go_fuzz_dep_.CoverTab[82269]++
										if q.count == len(q.buf) {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:51
		_go_fuzz_dep_.CoverTab[82271]++
											q.resize()
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:52
		// _ = "end of CoverTab[82271]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:53
		_go_fuzz_dep_.CoverTab[82272]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:53
		// _ = "end of CoverTab[82272]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:53
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:53
	// _ = "end of CoverTab[82269]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:53
	_go_fuzz_dep_.CoverTab[82270]++

										q.buf[q.tail] = elem

										q.tail = (q.tail + 1) & (len(q.buf) - 1)
										q.count++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:58
	// _ = "end of CoverTab[82270]"
}

// Peek returns the element at the head of the queue. This call panics
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:61
// if the queue is empty.
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:63
func (q *Queue) Peek() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:63
	_go_fuzz_dep_.CoverTab[82273]++
										if q.count <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:64
		_go_fuzz_dep_.CoverTab[82275]++
											panic("queue: Peek() called on empty queue")
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:65
		// _ = "end of CoverTab[82275]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:66
		_go_fuzz_dep_.CoverTab[82276]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:66
		// _ = "end of CoverTab[82276]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:66
	// _ = "end of CoverTab[82273]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:66
	_go_fuzz_dep_.CoverTab[82274]++
										return q.buf[q.head]
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:67
	// _ = "end of CoverTab[82274]"
}

// Get returns the element at index i in the queue. If the index is
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:70
// invalid, the call will panic. This method accepts both positive and
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:70
// negative index values. Index 0 refers to the first element, and
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:70
// index -1 refers to the last.
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:74
func (q *Queue) Get(i int) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:74
	_go_fuzz_dep_.CoverTab[82277]++

										if i < 0 {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:76
		_go_fuzz_dep_.CoverTab[82280]++
											i += q.count
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:77
		// _ = "end of CoverTab[82280]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:78
		_go_fuzz_dep_.CoverTab[82281]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:78
		// _ = "end of CoverTab[82281]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:78
	// _ = "end of CoverTab[82277]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:78
	_go_fuzz_dep_.CoverTab[82278]++
										if i < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:79
		_go_fuzz_dep_.CoverTab[82282]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:79
		return i >= q.count
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:79
		// _ = "end of CoverTab[82282]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:79
	}() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:79
		_go_fuzz_dep_.CoverTab[82283]++
											panic("queue: Get() called with index out of range")
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:80
		// _ = "end of CoverTab[82283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:81
		_go_fuzz_dep_.CoverTab[82284]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:81
		// _ = "end of CoverTab[82284]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:81
	// _ = "end of CoverTab[82278]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:81
	_go_fuzz_dep_.CoverTab[82279]++

										return q.buf[(q.head+i)&(len(q.buf)-1)]
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:83
	// _ = "end of CoverTab[82279]"
}

// Remove removes and returns the element from the front of the queue. If the
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:86
// queue is empty, the call will panic.
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:88
func (q *Queue) Remove() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:88
	_go_fuzz_dep_.CoverTab[82285]++
										if q.count <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:89
		_go_fuzz_dep_.CoverTab[82288]++
											panic("queue: Remove() called on empty queue")
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:90
		// _ = "end of CoverTab[82288]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:91
		_go_fuzz_dep_.CoverTab[82289]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:91
		// _ = "end of CoverTab[82289]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:91
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:91
	// _ = "end of CoverTab[82285]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:91
	_go_fuzz_dep_.CoverTab[82286]++
										ret := q.buf[q.head]
										q.buf[q.head] = nil

										q.head = (q.head + 1) & (len(q.buf) - 1)
										q.count--

										if len(q.buf) > minQueueLen && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:98
		_go_fuzz_dep_.CoverTab[82290]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:98
		return (q.count << 2) == len(q.buf)
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:98
		// _ = "end of CoverTab[82290]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:98
	}() {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:98
		_go_fuzz_dep_.CoverTab[82291]++
											q.resize()
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:99
		// _ = "end of CoverTab[82291]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:100
		_go_fuzz_dep_.CoverTab[82292]++
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:100
		// _ = "end of CoverTab[82292]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:100
	}
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:100
	// _ = "end of CoverTab[82286]"
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:100
	_go_fuzz_dep_.CoverTab[82287]++
										return ret
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:101
	// _ = "end of CoverTab[82287]"
}

//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:102
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/eapache/queue@v1.1.0/queue.go:102
var _ = _go_fuzz_dep_.CoverTab
