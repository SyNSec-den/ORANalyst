//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
package grpcutil

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:19
)

import (
	"strconv"
	"time"
)

const maxTimeoutValue int64 = 100000000 - 1

// div does integer division and round-up the result. Note that this is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:28
// equivalent to (d+r-1)/r but has less chance to overflow.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:30
func div(d, r time.Duration) int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:30
	_go_fuzz_dep_.CoverTab[67605]++
														if d%r > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:31
		_go_fuzz_dep_.CoverTab[67607]++
															return int64(d/r + 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:32
		// _ = "end of CoverTab[67607]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:33
		_go_fuzz_dep_.CoverTab[67608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:33
		// _ = "end of CoverTab[67608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:33
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:33
	// _ = "end of CoverTab[67605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:33
	_go_fuzz_dep_.CoverTab[67606]++
														return int64(d / r)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:34
	// _ = "end of CoverTab[67606]"
}

// EncodeDuration encodes the duration to the format grpc-timeout header
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:37
// accepts.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:37
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:37
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:41
func EncodeDuration(t time.Duration) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:41
	_go_fuzz_dep_.CoverTab[67609]++

														if t <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:43
		_go_fuzz_dep_.CoverTab[67616]++
															return "0n"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:44
		// _ = "end of CoverTab[67616]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:45
		_go_fuzz_dep_.CoverTab[67617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:45
		// _ = "end of CoverTab[67617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:45
	// _ = "end of CoverTab[67609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:45
	_go_fuzz_dep_.CoverTab[67610]++
														if d := div(t, time.Nanosecond); d <= maxTimeoutValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:46
		_go_fuzz_dep_.CoverTab[67618]++
															return strconv.FormatInt(d, 10) + "n"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:47
		// _ = "end of CoverTab[67618]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:48
		_go_fuzz_dep_.CoverTab[67619]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:48
		// _ = "end of CoverTab[67619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:48
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:48
	// _ = "end of CoverTab[67610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:48
	_go_fuzz_dep_.CoverTab[67611]++
														if d := div(t, time.Microsecond); d <= maxTimeoutValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:49
		_go_fuzz_dep_.CoverTab[67620]++
															return strconv.FormatInt(d, 10) + "u"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:50
		// _ = "end of CoverTab[67620]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:51
		_go_fuzz_dep_.CoverTab[67621]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:51
		// _ = "end of CoverTab[67621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:51
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:51
	// _ = "end of CoverTab[67611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:51
	_go_fuzz_dep_.CoverTab[67612]++
														if d := div(t, time.Millisecond); d <= maxTimeoutValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:52
		_go_fuzz_dep_.CoverTab[67622]++
															return strconv.FormatInt(d, 10) + "m"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:53
		// _ = "end of CoverTab[67622]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:54
		_go_fuzz_dep_.CoverTab[67623]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:54
		// _ = "end of CoverTab[67623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:54
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:54
	// _ = "end of CoverTab[67612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:54
	_go_fuzz_dep_.CoverTab[67613]++
														if d := div(t, time.Second); d <= maxTimeoutValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:55
		_go_fuzz_dep_.CoverTab[67624]++
															return strconv.FormatInt(d, 10) + "S"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:56
		// _ = "end of CoverTab[67624]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:57
		_go_fuzz_dep_.CoverTab[67625]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:57
		// _ = "end of CoverTab[67625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:57
	// _ = "end of CoverTab[67613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:57
	_go_fuzz_dep_.CoverTab[67614]++
														if d := div(t, time.Minute); d <= maxTimeoutValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:58
		_go_fuzz_dep_.CoverTab[67626]++
															return strconv.FormatInt(d, 10) + "M"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:59
		// _ = "end of CoverTab[67626]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:60
		_go_fuzz_dep_.CoverTab[67627]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:60
		// _ = "end of CoverTab[67627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:60
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:60
	// _ = "end of CoverTab[67614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:60
	_go_fuzz_dep_.CoverTab[67615]++

														return strconv.FormatInt(div(t, time.Hour), 10) + "H"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:62
	// _ = "end of CoverTab[67615]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpcutil/encode_duration.go:63
var _ = _go_fuzz_dep_.CoverTab
