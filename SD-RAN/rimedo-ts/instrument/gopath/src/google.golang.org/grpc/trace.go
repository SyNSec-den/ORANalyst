//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:19
)

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/trace"
)

// EnableTracing controls whether to trace RPCs using the golang.org/x/net/trace package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:33
// This should only be set before any RPCs are sent or received by this program.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:35
var EnableTracing bool

// methodFamily returns the trace family for the given method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:37
// It turns "/pkg.Service/GetFoo" into "pkg.Service".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:39
func methodFamily(m string) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:39
	_go_fuzz_dep_.CoverTab[81447]++
										m = strings.TrimPrefix(m, "/")
										if i := strings.Index(m, "/"); i >= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:41
		_go_fuzz_dep_.CoverTab[81449]++
											m = m[:i]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:42
		// _ = "end of CoverTab[81449]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:43
		_go_fuzz_dep_.CoverTab[81450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:43
		// _ = "end of CoverTab[81450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:43
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:43
	// _ = "end of CoverTab[81447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:43
	_go_fuzz_dep_.CoverTab[81448]++
										return m
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:44
	// _ = "end of CoverTab[81448]"
}

// traceInfo contains tracing information for an RPC.
type traceInfo struct {
	tr		trace.Trace
	firstLine	firstLine
}

// firstLine is the first line of an RPC trace.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:53
// It may be mutated after construction; remoteAddr specifically may change
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:53
// during client-side use.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:56
type firstLine struct {
	mu		sync.Mutex
	client		bool	// whether this is a client (outgoing) RPC
	remoteAddr	net.Addr
	deadline	time.Duration	// may be zero
}

func (f *firstLine) SetRemoteAddr(addr net.Addr) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:63
	_go_fuzz_dep_.CoverTab[81451]++
										f.mu.Lock()
										f.remoteAddr = addr
										f.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:66
	// _ = "end of CoverTab[81451]"
}

func (f *firstLine) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:69
	_go_fuzz_dep_.CoverTab[81452]++
										f.mu.Lock()
										defer f.mu.Unlock()

										var line bytes.Buffer
										io.WriteString(&line, "RPC: ")
										if f.client {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:75
		_go_fuzz_dep_.CoverTab[81455]++
											io.WriteString(&line, "to")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:76
		// _ = "end of CoverTab[81455]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:77
		_go_fuzz_dep_.CoverTab[81456]++
											io.WriteString(&line, "from")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:78
		// _ = "end of CoverTab[81456]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:79
	// _ = "end of CoverTab[81452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:79
	_go_fuzz_dep_.CoverTab[81453]++
										fmt.Fprintf(&line, " %v deadline:", f.remoteAddr)
										if f.deadline != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:81
		_go_fuzz_dep_.CoverTab[81457]++
											fmt.Fprint(&line, f.deadline)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:82
		// _ = "end of CoverTab[81457]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:83
		_go_fuzz_dep_.CoverTab[81458]++
											io.WriteString(&line, "none")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:84
		// _ = "end of CoverTab[81458]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:85
	// _ = "end of CoverTab[81453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:85
	_go_fuzz_dep_.CoverTab[81454]++
										return line.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:86
	// _ = "end of CoverTab[81454]"
}

const truncateSize = 100

func truncate(x string, l int) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:91
	_go_fuzz_dep_.CoverTab[81459]++
										if l > len(x) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:92
		_go_fuzz_dep_.CoverTab[81461]++
											return x
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:93
		// _ = "end of CoverTab[81461]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:94
		_go_fuzz_dep_.CoverTab[81462]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:94
		// _ = "end of CoverTab[81462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:94
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:94
	// _ = "end of CoverTab[81459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:94
	_go_fuzz_dep_.CoverTab[81460]++
										return x[:l]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:95
	// _ = "end of CoverTab[81460]"
}

// payload represents an RPC request or response payload.
type payload struct {
	sent	bool		// whether this is an outgoing payload
	msg	interface{}	// e.g. a proto.Message

}

func (p payload) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:105
	_go_fuzz_dep_.CoverTab[81463]++
										if p.sent {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:106
		_go_fuzz_dep_.CoverTab[81465]++
											return truncate(fmt.Sprintf("sent: %v", p.msg), truncateSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:107
		// _ = "end of CoverTab[81465]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:108
		_go_fuzz_dep_.CoverTab[81466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:108
		// _ = "end of CoverTab[81466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:108
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:108
	// _ = "end of CoverTab[81463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:108
	_go_fuzz_dep_.CoverTab[81464]++
										return truncate(fmt.Sprintf("recv: %v", p.msg), truncateSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:109
	// _ = "end of CoverTab[81464]"
}

type fmtStringer struct {
	format	string
	a	[]interface{}
}

func (f *fmtStringer) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:117
	_go_fuzz_dep_.CoverTab[81467]++
										return fmt.Sprintf(f.format, f.a...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:118
	// _ = "end of CoverTab[81467]"
}

type stringer string

func (s stringer) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:123
	_go_fuzz_dep_.CoverTab[81468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:123
	return string(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:123
	// _ = "end of CoverTab[81468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:123
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/trace.go:123
var _ = _go_fuzz_dep_.CoverTab
