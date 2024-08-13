//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
package grpclog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:19
)

import "google.golang.org/grpc/internal/grpclog"

// Logger mimics golang's standard Logger as an interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:23
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:23
// Deprecated: use LoggerV2.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:26
type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// SetLogger sets the logger that is used in grpc. Call only from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:35
// init() functions.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:35
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:35
// Deprecated: use SetLoggerV2.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:39
func SetLogger(l Logger) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:39
	_go_fuzz_dep_.CoverTab[48119]++
											grpclog.Logger = &loggerWrapper{Logger: l}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:40
	// _ = "end of CoverTab[48119]"
}

// loggerWrapper wraps Logger into a LoggerV2.
type loggerWrapper struct {
	Logger
}

func (g *loggerWrapper) Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:48
	_go_fuzz_dep_.CoverTab[48120]++
											g.Logger.Print(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:49
	// _ = "end of CoverTab[48120]"
}

func (g *loggerWrapper) Infoln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:52
	_go_fuzz_dep_.CoverTab[48121]++
											g.Logger.Println(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:53
	// _ = "end of CoverTab[48121]"
}

func (g *loggerWrapper) Infof(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:56
	_go_fuzz_dep_.CoverTab[48122]++
											g.Logger.Printf(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:57
	// _ = "end of CoverTab[48122]"
}

func (g *loggerWrapper) Warning(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:60
	_go_fuzz_dep_.CoverTab[48123]++
											g.Logger.Print(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:61
	// _ = "end of CoverTab[48123]"
}

func (g *loggerWrapper) Warningln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:64
	_go_fuzz_dep_.CoverTab[48124]++
											g.Logger.Println(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:65
	// _ = "end of CoverTab[48124]"
}

func (g *loggerWrapper) Warningf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:68
	_go_fuzz_dep_.CoverTab[48125]++
											g.Logger.Printf(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:69
	// _ = "end of CoverTab[48125]"
}

func (g *loggerWrapper) Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:72
	_go_fuzz_dep_.CoverTab[48126]++
											g.Logger.Print(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:73
	// _ = "end of CoverTab[48126]"
}

func (g *loggerWrapper) Errorln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:76
	_go_fuzz_dep_.CoverTab[48127]++
											g.Logger.Println(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:77
	// _ = "end of CoverTab[48127]"
}

func (g *loggerWrapper) Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:80
	_go_fuzz_dep_.CoverTab[48128]++
											g.Logger.Printf(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:81
	// _ = "end of CoverTab[48128]"
}

func (g *loggerWrapper) V(l int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:84
	_go_fuzz_dep_.CoverTab[48129]++

											return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:86
	// _ = "end of CoverTab[48129]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:87
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/logger.go:87
var _ = _go_fuzz_dep_.CoverTab
