//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:19
)

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
)

var logger = grpclog.Component("channelz")

func withParens(id *Identifier) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:29
	_go_fuzz_dep_.CoverTab[62831]++
													return "[" + id.String() + "] "
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:30
	// _ = "end of CoverTab[62831]"
}

// Info logs and adds a trace event if channelz is on.
func Info(l grpclog.DepthLoggerV2, id *Identifier, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:34
	_go_fuzz_dep_.CoverTab[62832]++
													AddTraceEvent(l, id, 1, &TraceEventDesc{
		Desc:		fmt.Sprint(args...),
		Severity:	CtInfo,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:38
	// _ = "end of CoverTab[62832]"
}

// Infof logs and adds a trace event if channelz is on.
func Infof(l grpclog.DepthLoggerV2, id *Identifier, format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:42
	_go_fuzz_dep_.CoverTab[62833]++
													AddTraceEvent(l, id, 1, &TraceEventDesc{
		Desc:		fmt.Sprintf(format, args...),
		Severity:	CtInfo,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:46
	// _ = "end of CoverTab[62833]"
}

// Warning logs and adds a trace event if channelz is on.
func Warning(l grpclog.DepthLoggerV2, id *Identifier, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:50
	_go_fuzz_dep_.CoverTab[62834]++
													AddTraceEvent(l, id, 1, &TraceEventDesc{
		Desc:		fmt.Sprint(args...),
		Severity:	CtWarning,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:54
	// _ = "end of CoverTab[62834]"
}

// Warningf logs and adds a trace event if channelz is on.
func Warningf(l grpclog.DepthLoggerV2, id *Identifier, format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:58
	_go_fuzz_dep_.CoverTab[62835]++
													AddTraceEvent(l, id, 1, &TraceEventDesc{
		Desc:		fmt.Sprintf(format, args...),
		Severity:	CtWarning,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:62
	// _ = "end of CoverTab[62835]"
}

// Error logs and adds a trace event if channelz is on.
func Error(l grpclog.DepthLoggerV2, id *Identifier, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:66
	_go_fuzz_dep_.CoverTab[62836]++
													AddTraceEvent(l, id, 1, &TraceEventDesc{
		Desc:		fmt.Sprint(args...),
		Severity:	CtError,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:70
	// _ = "end of CoverTab[62836]"
}

// Errorf logs and adds a trace event if channelz is on.
func Errorf(l grpclog.DepthLoggerV2, id *Identifier, format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:74
	_go_fuzz_dep_.CoverTab[62837]++
													AddTraceEvent(l, id, 1, &TraceEventDesc{
		Desc:		fmt.Sprintf(format, args...),
		Severity:	CtError,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:78
	// _ = "end of CoverTab[62837]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:79
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/logging.go:79
var _ = _go_fuzz_dep_.CoverTab
