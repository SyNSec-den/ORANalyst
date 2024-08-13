//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:19
)

import (
	"syscall"
)

// GetSocketOption gets the socket option info of the conn.
func GetSocketOption(socket interface{}) *SocketOptionData {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:26
	_go_fuzz_dep_.CoverTab[62951]++
													c, ok := socket.(syscall.Conn)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:28
		_go_fuzz_dep_.CoverTab[62954]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:29
		// _ = "end of CoverTab[62954]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:30
		_go_fuzz_dep_.CoverTab[62955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:30
		// _ = "end of CoverTab[62955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:30
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:30
	// _ = "end of CoverTab[62951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:30
	_go_fuzz_dep_.CoverTab[62952]++
													data := &SocketOptionData{}
													if rawConn, err := c.SyscallConn(); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:32
		_go_fuzz_dep_.CoverTab[62956]++
														rawConn.Control(data.Getsockopt)
														return data
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:34
		// _ = "end of CoverTab[62956]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:35
		_go_fuzz_dep_.CoverTab[62957]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:35
		// _ = "end of CoverTab[62957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:35
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:35
	// _ = "end of CoverTab[62952]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:35
	_go_fuzz_dep_.CoverTab[62953]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:36
	// _ = "end of CoverTab[62953]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:37
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/util_linux.go:37
var _ = _go_fuzz_dep_.CoverTab
