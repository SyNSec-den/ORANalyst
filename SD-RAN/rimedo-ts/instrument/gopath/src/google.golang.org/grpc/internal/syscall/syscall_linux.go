//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:19
// Package syscall provides functionalities that grpc uses to get low-level operating system
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:19
// stats/info.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
package syscall

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:21
)

import (
	"fmt"
	"net"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.Component("core")

// GetCPUTime returns the how much CPU time has passed since the start of this process.
func GetCPUTime() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:36
	_go_fuzz_dep_.CoverTab[75988]++
													var ts unix.Timespec
													if err := unix.ClockGettime(unix.CLOCK_PROCESS_CPUTIME_ID, &ts); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:38
		_go_fuzz_dep_.CoverTab[75990]++
														logger.Fatal(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:39
		// _ = "end of CoverTab[75990]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:40
		_go_fuzz_dep_.CoverTab[75991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:40
		// _ = "end of CoverTab[75991]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:40
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:40
	// _ = "end of CoverTab[75988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:40
	_go_fuzz_dep_.CoverTab[75989]++
													return ts.Nano()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:41
	// _ = "end of CoverTab[75989]"
}

// Rusage is an alias for syscall.Rusage under linux environment.
type Rusage = syscall.Rusage

// GetRusage returns the resource usage of current process.
func GetRusage() *Rusage {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:48
	_go_fuzz_dep_.CoverTab[75992]++
													rusage := new(Rusage)
													syscall.Getrusage(syscall.RUSAGE_SELF, rusage)
													return rusage
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:51
	// _ = "end of CoverTab[75992]"
}

// CPUTimeDiff returns the differences of user CPU time and system CPU time used
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:54
// between two Rusage structs.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:56
func CPUTimeDiff(first *Rusage, latest *Rusage) (float64, float64) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:56
	_go_fuzz_dep_.CoverTab[75993]++
													var (
		utimeDiffs	= latest.Utime.Sec - first.Utime.Sec
		utimeDiffus	= latest.Utime.Usec - first.Utime.Usec
		stimeDiffs	= latest.Stime.Sec - first.Stime.Sec
		stimeDiffus	= latest.Stime.Usec - first.Stime.Usec
	)

													uTimeElapsed := float64(utimeDiffs) + float64(utimeDiffus)*1.0e-6
													sTimeElapsed := float64(stimeDiffs) + float64(stimeDiffus)*1.0e-6

													return uTimeElapsed, sTimeElapsed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:67
	// _ = "end of CoverTab[75993]"
}

// SetTCPUserTimeout sets the TCP user timeout on a connection's socket
func SetTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:71
	_go_fuzz_dep_.CoverTab[75994]++
													tcpconn, ok := conn.(*net.TCPConn)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:73
		_go_fuzz_dep_.CoverTab[75999]++

														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:75
		// _ = "end of CoverTab[75999]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:76
		_go_fuzz_dep_.CoverTab[76000]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:76
		// _ = "end of CoverTab[76000]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:76
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:76
	// _ = "end of CoverTab[75994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:76
	_go_fuzz_dep_.CoverTab[75995]++
													rawConn, err := tcpconn.SyscallConn()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:78
		_go_fuzz_dep_.CoverTab[76001]++
														return fmt.Errorf("error getting raw connection: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:79
		// _ = "end of CoverTab[76001]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:80
		_go_fuzz_dep_.CoverTab[76002]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:80
		// _ = "end of CoverTab[76002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:80
	// _ = "end of CoverTab[75995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:80
	_go_fuzz_dep_.CoverTab[75996]++
													err = rawConn.Control(func(fd uintptr) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:81
		_go_fuzz_dep_.CoverTab[76003]++
														err = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, unix.TCP_USER_TIMEOUT, int(timeout/time.Millisecond))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:82
		// _ = "end of CoverTab[76003]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:83
	// _ = "end of CoverTab[75996]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:83
	_go_fuzz_dep_.CoverTab[75997]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:84
		_go_fuzz_dep_.CoverTab[76004]++
														return fmt.Errorf("error setting option on socket: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:85
		// _ = "end of CoverTab[76004]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:86
		_go_fuzz_dep_.CoverTab[76005]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:86
		// _ = "end of CoverTab[76005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:86
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:86
	// _ = "end of CoverTab[75997]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:86
	_go_fuzz_dep_.CoverTab[75998]++

													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:88
	// _ = "end of CoverTab[75998]"
}

// GetTCPUserTimeout gets the TCP user timeout on a connection's socket
func GetTCPUserTimeout(conn net.Conn) (opt int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:92
	_go_fuzz_dep_.CoverTab[76006]++
													tcpconn, ok := conn.(*net.TCPConn)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:94
		_go_fuzz_dep_.CoverTab[76011]++
														err = fmt.Errorf("conn is not *net.TCPConn. got %T", conn)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:96
		// _ = "end of CoverTab[76011]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:97
		_go_fuzz_dep_.CoverTab[76012]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:97
		// _ = "end of CoverTab[76012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:97
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:97
	// _ = "end of CoverTab[76006]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:97
	_go_fuzz_dep_.CoverTab[76007]++
													rawConn, err := tcpconn.SyscallConn()
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:99
		_go_fuzz_dep_.CoverTab[76013]++
														err = fmt.Errorf("error getting raw connection: %v", err)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:101
		// _ = "end of CoverTab[76013]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:102
		_go_fuzz_dep_.CoverTab[76014]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:102
		// _ = "end of CoverTab[76014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:102
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:102
	// _ = "end of CoverTab[76007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:102
	_go_fuzz_dep_.CoverTab[76008]++
													err = rawConn.Control(func(fd uintptr) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:103
		_go_fuzz_dep_.CoverTab[76015]++
														opt, err = syscall.GetsockoptInt(int(fd), syscall.IPPROTO_TCP, unix.TCP_USER_TIMEOUT)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:104
		// _ = "end of CoverTab[76015]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:105
	// _ = "end of CoverTab[76008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:105
	_go_fuzz_dep_.CoverTab[76009]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:106
		_go_fuzz_dep_.CoverTab[76016]++
														err = fmt.Errorf("error getting option on socket: %v", err)
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:108
		// _ = "end of CoverTab[76016]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:109
		_go_fuzz_dep_.CoverTab[76017]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:109
		// _ = "end of CoverTab[76017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:109
	// _ = "end of CoverTab[76009]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:109
	_go_fuzz_dep_.CoverTab[76010]++

													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:111
	// _ = "end of CoverTab[76010]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:112
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/syscall/syscall_linux.go:112
var _ = _go_fuzz_dep_.CoverTab
