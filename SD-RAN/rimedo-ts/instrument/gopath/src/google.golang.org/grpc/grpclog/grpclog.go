//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
// Package grpclog defines logging for grpc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
// All logs in transport and grpclb packages only go to verbose level 2.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
// All logs in other packages in grpc are logged in spite of the verbosity level.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
// In the default logger,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
// severity level can be set by environment variable GRPC_GO_LOG_SEVERITY_LEVEL,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:19
// verbosity level can be set by GRPC_GO_LOG_VERBOSITY_LEVEL.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
package grpclog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:27
)

import (
	"os"

	"google.golang.org/grpc/internal/grpclog"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:40
	_go_fuzz_dep_.CoverTab[48103]++
											return grpclog.Logger.V(l)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:41
	// _ = "end of CoverTab[48103]"
}

// Info logs to the INFO log.
func Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:45
	_go_fuzz_dep_.CoverTab[48104]++
											grpclog.Logger.Info(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:46
	// _ = "end of CoverTab[48104]"
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:50
	_go_fuzz_dep_.CoverTab[48105]++
											grpclog.Logger.Infof(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:51
	// _ = "end of CoverTab[48105]"
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:55
	_go_fuzz_dep_.CoverTab[48106]++
											grpclog.Logger.Infoln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:56
	// _ = "end of CoverTab[48106]"
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:60
	_go_fuzz_dep_.CoverTab[48107]++
											grpclog.Logger.Warning(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:61
	// _ = "end of CoverTab[48107]"
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:65
	_go_fuzz_dep_.CoverTab[48108]++
											grpclog.Logger.Warningf(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:66
	// _ = "end of CoverTab[48108]"
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:70
	_go_fuzz_dep_.CoverTab[48109]++
											grpclog.Logger.Warningln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:71
	// _ = "end of CoverTab[48109]"
}

// Error logs to the ERROR log.
func Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:75
	_go_fuzz_dep_.CoverTab[48110]++
											grpclog.Logger.Error(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:76
	// _ = "end of CoverTab[48110]"
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:80
	_go_fuzz_dep_.CoverTab[48111]++
											grpclog.Logger.Errorf(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:81
	// _ = "end of CoverTab[48111]"
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:85
	_go_fuzz_dep_.CoverTab[48112]++
											grpclog.Logger.Errorln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:86
	// _ = "end of CoverTab[48112]"
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:89
// It calls os.Exit() with exit code 1.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:91
func Fatal(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:91
	_go_fuzz_dep_.CoverTab[48113]++
											grpclog.Logger.Fatal(args...)

											os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:94
	// _ = "end of CoverTab[48113]"
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:97
// It calls os.Exit() with exit code 1.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:99
func Fatalf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:99
		_go_fuzz_dep_.CoverTab[48114]++
												grpclog.Logger.Fatalf(format, args...)

												os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:102
	// _ = "end of CoverTab[48114]"
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:105
// It calle os.Exit()) with exit code 1.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:107
func Fatalln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:107
	_go_fuzz_dep_.CoverTab[48115]++
												grpclog.Logger.Fatalln(args...)

												os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:110
	// _ = "end of CoverTab[48115]"
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:113
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:113
// Deprecated: use Info.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:116
func Print(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:116
	_go_fuzz_dep_.CoverTab[48116]++
												grpclog.Logger.Info(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:117
	// _ = "end of CoverTab[48116]"
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:120
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:120
// Deprecated: use Infof.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:123
func Printf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:123
	_go_fuzz_dep_.CoverTab[48117]++
												grpclog.Logger.Infof(format, args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:124
	// _ = "end of CoverTab[48117]"
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:127
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:127
// Deprecated: use Infoln.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:130
func Println(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:130
	_go_fuzz_dep_.CoverTab[48118]++
												grpclog.Logger.Infoln(args...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:131
	// _ = "end of CoverTab[48118]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/grpclog.go:132
var _ = _go_fuzz_dep_.CoverTab
