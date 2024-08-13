//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
package credentials

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:19
)

import (
	"net"
	"syscall"
)

type sysConn = syscall.Conn

// syscallConn keeps reference of rawConn to support syscall.Conn for channelz.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// SyscallConn() (the method in interface syscall.Conn) is explicitly
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// implemented on this type,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// Interface syscall.Conn is implemented by most net.Conn implementations (e.g.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// TCPConn, UnixConn), but is not part of net.Conn interface. So wrapper conns
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// that embed net.Conn don't implement syscall.Conn. (Side note: tls.Conn
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// doesn't embed net.Conn, so even if syscall.Conn is part of net.Conn, it won't
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:28
// help here).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:37
type syscallConn struct {
	net.Conn
	// sysConn is a type alias of syscall.Conn. It's necessary because the name
	// `Conn` collides with `net.Conn`.
	sysConn
}

// WrapSyscallConn tries to wrap rawConn and newConn into a net.Conn that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:44
// implements syscall.Conn. rawConn will be used to support syscall, and newConn
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:44
// will be used for read/write.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:44
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:44
// This function returns newConn if rawConn doesn't implement syscall.Conn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:49
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:49
	_go_fuzz_dep_.CoverTab[62493]++
														sysConn, ok := rawConn.(syscall.Conn)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:51
		_go_fuzz_dep_.CoverTab[62495]++
															return newConn
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:52
		// _ = "end of CoverTab[62495]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:53
		_go_fuzz_dep_.CoverTab[62496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:53
		// _ = "end of CoverTab[62496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:53
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:53
	// _ = "end of CoverTab[62493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:53
	_go_fuzz_dep_.CoverTab[62494]++
														return &syscallConn{
		Conn:		newConn,
		sysConn:	sysConn,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:57
	// _ = "end of CoverTab[62494]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:58
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/credentials/syscallconn.go:58
var _ = _go_fuzz_dep_.CoverTab
