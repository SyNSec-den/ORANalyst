//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:19
// Package grpclog (internal) defines depth logging for grpc.
package grpclog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:20
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:20
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:20
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:20
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:20
)

import (
	"os"
)

// Logger is the logger used for the non-depth log functions.
var Logger LoggerV2

// DepthLogger is the logger used for the depth log functions.
var DepthLogger DepthLoggerV2

// InfoDepth logs to the INFO log at the specified depth.
func InfoDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:33
	_go_fuzz_dep_.CoverTab[48048]++
													if DepthLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:34
		_go_fuzz_dep_.CoverTab[48049]++
														DepthLogger.InfoDepth(depth, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:35
		// _ = "end of CoverTab[48049]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:36
		_go_fuzz_dep_.CoverTab[48050]++
														Logger.Infoln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:37
		// _ = "end of CoverTab[48050]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:38
	// _ = "end of CoverTab[48048]"
}

// WarningDepth logs to the WARNING log at the specified depth.
func WarningDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:42
	_go_fuzz_dep_.CoverTab[48051]++
													if DepthLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:43
		_go_fuzz_dep_.CoverTab[48052]++
														DepthLogger.WarningDepth(depth, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:44
		// _ = "end of CoverTab[48052]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:45
		_go_fuzz_dep_.CoverTab[48053]++
														Logger.Warningln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:46
		// _ = "end of CoverTab[48053]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:47
	// _ = "end of CoverTab[48051]"
}

// ErrorDepth logs to the ERROR log at the specified depth.
func ErrorDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:51
	_go_fuzz_dep_.CoverTab[48054]++
													if DepthLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:52
		_go_fuzz_dep_.CoverTab[48055]++
														DepthLogger.ErrorDepth(depth, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:53
		// _ = "end of CoverTab[48055]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:54
		_go_fuzz_dep_.CoverTab[48056]++
														Logger.Errorln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:55
		// _ = "end of CoverTab[48056]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:56
	// _ = "end of CoverTab[48054]"
}

// FatalDepth logs to the FATAL log at the specified depth.
func FatalDepth(depth int, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:60
	_go_fuzz_dep_.CoverTab[48057]++
													if DepthLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:61
		_go_fuzz_dep_.CoverTab[48059]++
														DepthLogger.FatalDepth(depth, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:62
		// _ = "end of CoverTab[48059]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:63
		_go_fuzz_dep_.CoverTab[48060]++
														Logger.Fatalln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:64
		// _ = "end of CoverTab[48060]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:65
	// _ = "end of CoverTab[48057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:65
	_go_fuzz_dep_.CoverTab[48058]++
													os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:66
	// _ = "end of CoverTab[48058]"
}

// LoggerV2 does underlying logging work for grpclog.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:69
// This is a copy of the LoggerV2 defined in the external grpclog package. It
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:69
// is defined here to avoid a circular dependency.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:72
type LoggerV2 interface {
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})
	// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
	Infoln(args ...interface{})
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})
	// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
	Warning(args ...interface{})
	// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
	Warningln(args ...interface{})
	// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...interface{})
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})
	// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	Errorln(args ...interface{})
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
	// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...interface{})
	// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalln(args ...interface{})
	// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalf(format string, args ...interface{})
	// V reports whether verbosity level l is at least the requested verbose level.
	V(l int) bool
}

// DepthLoggerV2 logs at a specified call frame. If a LoggerV2 also implements
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// DepthLoggerV2, the below functions will be called with the appropriate stack
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// depth set for trivial functions the logger may ignore.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// This is a copy of the DepthLoggerV2 defined in the external grpclog package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// It is defined here to avoid a circular dependency.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:107
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:117
type DepthLoggerV2 interface {
	// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Println.
	InfoDepth(depth int, args ...interface{})
	// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Println.
	WarningDepth(depth int, args ...interface{})
	// ErrorDepth logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Println.
	ErrorDepth(depth int, args ...interface{})
	// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Println.
	FatalDepth(depth int, args ...interface{})
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:126
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/grpclog.go:126
var _ = _go_fuzz_dep_.CoverTab
