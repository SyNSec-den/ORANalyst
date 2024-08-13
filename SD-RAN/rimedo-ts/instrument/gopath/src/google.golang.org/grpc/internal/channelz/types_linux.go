//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:19
)

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// SocketOptionData defines the struct to hold socket option data, and related
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:27
// getter function to obtain info from fd.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:29
type SocketOptionData struct {
	Linger		*unix.Linger
	RecvTimeout	*unix.Timeval
	SendTimeout	*unix.Timeval
	TCPInfo		*unix.TCPInfo
}

// Getsockopt defines the function to get socket options requested by channelz.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:36
// It is to be passed to syscall.RawConn.Control().
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:38
func (s *SocketOptionData) Getsockopt(fd uintptr) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:38
	_go_fuzz_dep_.CoverTab[62939]++
													if v, err := unix.GetsockoptLinger(int(fd), syscall.SOL_SOCKET, syscall.SO_LINGER); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:39
		_go_fuzz_dep_.CoverTab[62943]++
														s.Linger = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:40
		// _ = "end of CoverTab[62943]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:41
		_go_fuzz_dep_.CoverTab[62944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:41
		// _ = "end of CoverTab[62944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:41
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:41
	// _ = "end of CoverTab[62939]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:41
	_go_fuzz_dep_.CoverTab[62940]++
													if v, err := unix.GetsockoptTimeval(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVTIMEO); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:42
		_go_fuzz_dep_.CoverTab[62945]++
														s.RecvTimeout = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:43
		// _ = "end of CoverTab[62945]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:44
		_go_fuzz_dep_.CoverTab[62946]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:44
		// _ = "end of CoverTab[62946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:44
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:44
	// _ = "end of CoverTab[62940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:44
	_go_fuzz_dep_.CoverTab[62941]++
													if v, err := unix.GetsockoptTimeval(int(fd), syscall.SOL_SOCKET, syscall.SO_SNDTIMEO); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:45
		_go_fuzz_dep_.CoverTab[62947]++
														s.SendTimeout = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:46
		// _ = "end of CoverTab[62947]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:47
		_go_fuzz_dep_.CoverTab[62948]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:47
		// _ = "end of CoverTab[62948]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:47
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:47
	// _ = "end of CoverTab[62941]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:47
	_go_fuzz_dep_.CoverTab[62942]++
													if v, err := unix.GetsockoptTCPInfo(int(fd), syscall.SOL_TCP, syscall.TCP_INFO); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:48
		_go_fuzz_dep_.CoverTab[62949]++
														s.TCPInfo = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:49
		// _ = "end of CoverTab[62949]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:50
		_go_fuzz_dep_.CoverTab[62950]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:50
		// _ = "end of CoverTab[62950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:50
	// _ = "end of CoverTab[62942]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/types_linux.go:51
var _ = _go_fuzz_dep_.CoverTab
