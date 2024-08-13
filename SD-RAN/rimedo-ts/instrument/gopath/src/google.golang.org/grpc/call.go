//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:19
)

import (
	"context"
)

// Invoke sends the RPC request on the wire and returns after response is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:25
// received.  This is typically called by generated code.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:25
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:25
// All errors returned by Invoke are compatible with the status package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:29
func (cc *ClientConn) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...CallOption) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:29
	_go_fuzz_dep_.CoverTab[78837]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:32
	opts = combine(cc.dopts.callOptions, opts)

	if cc.dopts.unaryInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:34
		_go_fuzz_dep_.CoverTab[78839]++
											return cc.dopts.unaryInt(ctx, method, args, reply, cc, invoke, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:35
		// _ = "end of CoverTab[78839]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:36
		_go_fuzz_dep_.CoverTab[78840]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:36
		// _ = "end of CoverTab[78840]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:36
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:36
	// _ = "end of CoverTab[78837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:36
	_go_fuzz_dep_.CoverTab[78838]++
										return invoke(ctx, method, args, reply, cc, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:37
	// _ = "end of CoverTab[78838]"
}

func combine(o1 []CallOption, o2 []CallOption) []CallOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:40
	_go_fuzz_dep_.CoverTab[78841]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:44
	if len(o1) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:44
		_go_fuzz_dep_.CoverTab[78843]++
											return o2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:45
		// _ = "end of CoverTab[78843]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:46
		_go_fuzz_dep_.CoverTab[78844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:46
		if len(o2) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:46
			_go_fuzz_dep_.CoverTab[78845]++
												return o1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:47
			// _ = "end of CoverTab[78845]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
			_go_fuzz_dep_.CoverTab[78846]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
			// _ = "end of CoverTab[78846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
		// _ = "end of CoverTab[78844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
	// _ = "end of CoverTab[78841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:48
	_go_fuzz_dep_.CoverTab[78842]++
										ret := make([]CallOption, len(o1)+len(o2))
										copy(ret, o1)
										copy(ret[len(o1):], o2)
										return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:52
	// _ = "end of CoverTab[78842]"
}

// Invoke sends the RPC request on the wire and returns after response is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:55
// received.  This is typically called by generated code.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:55
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:55
// DEPRECATED: Use ClientConn.Invoke instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:59
func Invoke(ctx context.Context, method string, args, reply interface{}, cc *ClientConn, opts ...CallOption) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:59
	_go_fuzz_dep_.CoverTab[78847]++
										return cc.Invoke(ctx, method, args, reply, opts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:60
	// _ = "end of CoverTab[78847]"
}

var unaryStreamDesc = &StreamDesc{ServerStreams: false, ClientStreams: false}

func invoke(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:65
	_go_fuzz_dep_.CoverTab[78848]++
										cs, err := newClientStream(ctx, unaryStreamDesc, cc, method, opts...)
										if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:67
		_go_fuzz_dep_.CoverTab[78851]++
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:68
		// _ = "end of CoverTab[78851]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:69
		_go_fuzz_dep_.CoverTab[78852]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:69
		// _ = "end of CoverTab[78852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:69
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:69
	// _ = "end of CoverTab[78848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:69
	_go_fuzz_dep_.CoverTab[78849]++
										if err := cs.SendMsg(req); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:70
		_go_fuzz_dep_.CoverTab[78853]++
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:71
		// _ = "end of CoverTab[78853]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:72
		_go_fuzz_dep_.CoverTab[78854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:72
		// _ = "end of CoverTab[78854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:72
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:72
	// _ = "end of CoverTab[78849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:72
	_go_fuzz_dep_.CoverTab[78850]++
										return cs.RecvMsg(reply)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:73
	// _ = "end of CoverTab[78850]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/call.go:74
var _ = _go_fuzz_dep_.CoverTab
