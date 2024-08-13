//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
package grpclog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:19
)

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc/internal/grpclog"
)

// LoggerV2 does underlying logging work for grpclog.
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

// SetLoggerV2 sets logger that is used in grpc to a V2 logger.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:69
// Not mutex-protected, should be called before any gRPC functions.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:71
func SetLoggerV2(l LoggerV2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:71
	_go_fuzz_dep_.CoverTab[48130]++
												if _, ok := l.(*componentData); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:72
		_go_fuzz_dep_.CoverTab[48132]++
													panic("cannot use component logger as grpclog logger")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:73
		// _ = "end of CoverTab[48132]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:74
		_go_fuzz_dep_.CoverTab[48133]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:74
		// _ = "end of CoverTab[48133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:74
	// _ = "end of CoverTab[48130]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:74
	_go_fuzz_dep_.CoverTab[48131]++
												grpclog.Logger = l
												grpclog.DepthLogger, _ = l.(grpclog.DepthLoggerV2)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:76
	// _ = "end of CoverTab[48131]"
}

const (
	// infoLog indicates Info severity.
	infoLog	int	= iota
	// warningLog indicates Warning severity.
	warningLog
	// errorLog indicates Error severity.
	errorLog
	// fatalLog indicates Fatal severity.
	fatalLog
)

// severityName contains the string representation of each severity.
var severityName = []string{
	infoLog:	"INFO",
	warningLog:	"WARNING",
	errorLog:	"ERROR",
	fatalLog:	"FATAL",
}

// loggerT is the default logger used by grpclog.
type loggerT struct {
	m		[]*log.Logger
	v		int
	jsonFormat	bool
}

// NewLoggerV2 creates a loggerV2 with the provided writers.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:105
// Fatal logs will be written to errorW, warningW, infoW, followed by exit(1).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:105
// Error logs will be written to errorW, warningW and infoW.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:105
// Warning logs will be written to warningW and infoW.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:105
// Info logs will be written to infoW.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:110
func NewLoggerV2(infoW, warningW, errorW io.Writer) LoggerV2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:110
	_go_fuzz_dep_.CoverTab[48134]++
												return newLoggerV2WithConfig(infoW, warningW, errorW, loggerV2Config{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:111
	// _ = "end of CoverTab[48134]"
}

// NewLoggerV2WithVerbosity creates a loggerV2 with the provided writers and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:114
// verbosity level.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:116
func NewLoggerV2WithVerbosity(infoW, warningW, errorW io.Writer, v int) LoggerV2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:116
	_go_fuzz_dep_.CoverTab[48135]++
												return newLoggerV2WithConfig(infoW, warningW, errorW, loggerV2Config{verbose: v})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:117
	// _ = "end of CoverTab[48135]"
}

type loggerV2Config struct {
	verbose		int
	jsonFormat	bool
}

func newLoggerV2WithConfig(infoW, warningW, errorW io.Writer, c loggerV2Config) LoggerV2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:125
	_go_fuzz_dep_.CoverTab[48136]++
												var m []*log.Logger
												flag := log.LstdFlags
												if c.jsonFormat {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:128
		_go_fuzz_dep_.CoverTab[48138]++
													flag = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:129
		// _ = "end of CoverTab[48138]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:130
		_go_fuzz_dep_.CoverTab[48139]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:130
		// _ = "end of CoverTab[48139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:130
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:130
	// _ = "end of CoverTab[48136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:130
	_go_fuzz_dep_.CoverTab[48137]++
												m = append(m, log.New(infoW, "", flag))
												m = append(m, log.New(io.MultiWriter(infoW, warningW), "", flag))
												ew := io.MultiWriter(infoW, warningW, errorW)
												m = append(m, log.New(ew, "", flag))
												m = append(m, log.New(ew, "", flag))
												return &loggerT{m: m, v: c.verbose, jsonFormat: c.jsonFormat}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:136
	// _ = "end of CoverTab[48137]"
}

// newLoggerV2 creates a loggerV2 to be used as default logger.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:139
// All logs are written to stderr.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:141
func newLoggerV2() LoggerV2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:141
	_go_fuzz_dep_.CoverTab[48140]++
												errorW := io.Discard
												warningW := io.Discard
												infoW := io.Discard

												logLevel := os.Getenv("GRPC_GO_LOG_SEVERITY_LEVEL")
												switch logLevel {
	case "", "ERROR", "error":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:148
		_go_fuzz_dep_.CoverTab[48143]++
													errorW = os.Stderr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:149
		// _ = "end of CoverTab[48143]"
	case "WARNING", "warning":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:150
		_go_fuzz_dep_.CoverTab[48144]++
													warningW = os.Stderr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:151
		// _ = "end of CoverTab[48144]"
	case "INFO", "info":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:152
		_go_fuzz_dep_.CoverTab[48145]++
													infoW = os.Stderr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:153
		// _ = "end of CoverTab[48145]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:153
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:153
		_go_fuzz_dep_.CoverTab[48146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:153
		// _ = "end of CoverTab[48146]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:154
	// _ = "end of CoverTab[48140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:154
	_go_fuzz_dep_.CoverTab[48141]++

												var v int
												vLevel := os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL")
												if vl, err := strconv.Atoi(vLevel); err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:158
		_go_fuzz_dep_.CoverTab[48147]++
													v = vl
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:159
		// _ = "end of CoverTab[48147]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:160
		_go_fuzz_dep_.CoverTab[48148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:160
		// _ = "end of CoverTab[48148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:160
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:160
	// _ = "end of CoverTab[48141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:160
	_go_fuzz_dep_.CoverTab[48142]++

												jsonFormat := strings.EqualFold(os.Getenv("GRPC_GO_LOG_FORMATTER"), "json")

												return newLoggerV2WithConfig(infoW, warningW, errorW, loggerV2Config{
		verbose:	v,
		jsonFormat:	jsonFormat,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:167
	// _ = "end of CoverTab[48142]"
}

func (g *loggerT) output(severity int, s string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:170
	_go_fuzz_dep_.CoverTab[48149]++
												sevStr := severityName[severity]
												if !g.jsonFormat {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:172
		_go_fuzz_dep_.CoverTab[48151]++
													g.m[severity].Output(2, fmt.Sprintf("%v: %v", sevStr, s))
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:174
		// _ = "end of CoverTab[48151]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:175
		_go_fuzz_dep_.CoverTab[48152]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:175
		// _ = "end of CoverTab[48152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:175
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:175
	// _ = "end of CoverTab[48149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:175
	_go_fuzz_dep_.CoverTab[48150]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:178
	b, _ := json.Marshal(map[string]string{
		"severity":	sevStr,
		"message":	s,
	})
												g.m[severity].Output(2, string(b))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:182
	// _ = "end of CoverTab[48150]"
}

func (g *loggerT) Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:185
	_go_fuzz_dep_.CoverTab[48153]++
												g.output(infoLog, fmt.Sprint(args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:186
	// _ = "end of CoverTab[48153]"
}

func (g *loggerT) Infoln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:189
	_go_fuzz_dep_.CoverTab[48154]++
												g.output(infoLog, fmt.Sprintln(args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:190
	// _ = "end of CoverTab[48154]"
}

func (g *loggerT) Infof(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:193
	_go_fuzz_dep_.CoverTab[48155]++
												g.output(infoLog, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:194
	// _ = "end of CoverTab[48155]"
}

func (g *loggerT) Warning(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:197
	_go_fuzz_dep_.CoverTab[48156]++
												g.output(warningLog, fmt.Sprint(args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:198
	// _ = "end of CoverTab[48156]"
}

func (g *loggerT) Warningln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:201
	_go_fuzz_dep_.CoverTab[48157]++
												g.output(warningLog, fmt.Sprintln(args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:202
	// _ = "end of CoverTab[48157]"
}

func (g *loggerT) Warningf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:205
	_go_fuzz_dep_.CoverTab[48158]++
												g.output(warningLog, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:206
	// _ = "end of CoverTab[48158]"
}

func (g *loggerT) Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:209
	_go_fuzz_dep_.CoverTab[48159]++
												g.output(errorLog, fmt.Sprint(args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:210
	// _ = "end of CoverTab[48159]"
}

func (g *loggerT) Errorln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:213
	_go_fuzz_dep_.CoverTab[48160]++
												g.output(errorLog, fmt.Sprintln(args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:214
	// _ = "end of CoverTab[48160]"
}

func (g *loggerT) Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:217
	_go_fuzz_dep_.CoverTab[48161]++
												g.output(errorLog, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:218
	// _ = "end of CoverTab[48161]"
}

func (g *loggerT) Fatal(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:221
	_go_fuzz_dep_.CoverTab[48162]++
												g.output(fatalLog, fmt.Sprint(args...))
												os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:223
	// _ = "end of CoverTab[48162]"
}

func (g *loggerT) Fatalln(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:226
	_go_fuzz_dep_.CoverTab[48163]++
												g.output(fatalLog, fmt.Sprintln(args...))
												os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:228
	// _ = "end of CoverTab[48163]"
}

func (g *loggerT) Fatalf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:231
	_go_fuzz_dep_.CoverTab[48164]++
												g.output(fatalLog, fmt.Sprintf(format, args...))
												os.Exit(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:233
	// _ = "end of CoverTab[48164]"
}

func (g *loggerT) V(l int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:236
	_go_fuzz_dep_.CoverTab[48165]++
												return l <= g.v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:237
	// _ = "end of CoverTab[48165]"
}

// DepthLoggerV2 logs at a specified call frame. If a LoggerV2 also implements
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
// DepthLoggerV2, the below functions will be called with the appropriate stack
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
// depth set for trivial functions the logger may ignore.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:240
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:248
type DepthLoggerV2 interface {
	LoggerV2
	// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Println.
	InfoDepth(depth int, args ...interface{})
	// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Println.
	WarningDepth(depth int, args ...interface{})
	// ErrorDepth logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Println.
	ErrorDepth(depth int, args ...interface{})
	// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Println.
	FatalDepth(depth int, args ...interface{})
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:258
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/grpclog/loggerv2.go:258
var _ = _go_fuzz_dep_.CoverTab
