//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
package grpclog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:19
)

import (
	"fmt"
)

// PrefixLogger does logging with a prefix.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:25
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:25
// Logging method on a nil logs without any prefix.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:28
type PrefixLogger struct {
	logger	DepthLoggerV2
	prefix	string
}

// Infof does info logging.
func (pl *PrefixLogger) Infof(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:34
	_go_fuzz_dep_.CoverTab[48061]++
													if pl != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:35
		_go_fuzz_dep_.CoverTab[48063]++

														format = pl.prefix + format
														pl.logger.InfoDepth(1, fmt.Sprintf(format, args...))
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:39
		// _ = "end of CoverTab[48063]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:40
		_go_fuzz_dep_.CoverTab[48064]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:40
		// _ = "end of CoverTab[48064]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:40
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:40
	// _ = "end of CoverTab[48061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:40
	_go_fuzz_dep_.CoverTab[48062]++
													InfoDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:41
	// _ = "end of CoverTab[48062]"
}

// Warningf does warning logging.
func (pl *PrefixLogger) Warningf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:45
	_go_fuzz_dep_.CoverTab[48065]++
													if pl != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:46
		_go_fuzz_dep_.CoverTab[48067]++
														format = pl.prefix + format
														pl.logger.WarningDepth(1, fmt.Sprintf(format, args...))
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:49
		// _ = "end of CoverTab[48067]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:50
		_go_fuzz_dep_.CoverTab[48068]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:50
		// _ = "end of CoverTab[48068]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:50
	// _ = "end of CoverTab[48065]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:50
	_go_fuzz_dep_.CoverTab[48066]++
													WarningDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:51
	// _ = "end of CoverTab[48066]"
}

// Errorf does error logging.
func (pl *PrefixLogger) Errorf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:55
	_go_fuzz_dep_.CoverTab[48069]++
													if pl != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:56
		_go_fuzz_dep_.CoverTab[48071]++
														format = pl.prefix + format
														pl.logger.ErrorDepth(1, fmt.Sprintf(format, args...))
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:59
		// _ = "end of CoverTab[48071]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:60
		_go_fuzz_dep_.CoverTab[48072]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:60
		// _ = "end of CoverTab[48072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:60
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:60
	// _ = "end of CoverTab[48069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:60
	_go_fuzz_dep_.CoverTab[48070]++
													ErrorDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:61
	// _ = "end of CoverTab[48070]"
}

// Debugf does info logging at verbose level 2.
func (pl *PrefixLogger) Debugf(format string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:65
	_go_fuzz_dep_.CoverTab[48073]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:69
	if !Logger.V(2) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:69
		_go_fuzz_dep_.CoverTab[48076]++
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:70
		// _ = "end of CoverTab[48076]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:71
		_go_fuzz_dep_.CoverTab[48077]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:71
		// _ = "end of CoverTab[48077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:71
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:71
	// _ = "end of CoverTab[48073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:71
	_go_fuzz_dep_.CoverTab[48074]++
													if pl != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:72
		_go_fuzz_dep_.CoverTab[48078]++

														format = pl.prefix + format
														pl.logger.InfoDepth(1, fmt.Sprintf(format, args...))
														return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:76
		// _ = "end of CoverTab[48078]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:77
		_go_fuzz_dep_.CoverTab[48079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:77
		// _ = "end of CoverTab[48079]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:77
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:77
	// _ = "end of CoverTab[48074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:77
	_go_fuzz_dep_.CoverTab[48075]++
													InfoDepth(1, fmt.Sprintf(format, args...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:78
	// _ = "end of CoverTab[48075]"

}

// V reports whether verbosity level l is at least the requested verbose level.
func (pl *PrefixLogger) V(l int) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:83
	_go_fuzz_dep_.CoverTab[48080]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:87
	return Logger.V(l)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:87
	// _ = "end of CoverTab[48080]"
}

// NewPrefixLogger creates a prefix logger with the given prefix.
func NewPrefixLogger(logger DepthLoggerV2, prefix string) *PrefixLogger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:91
	_go_fuzz_dep_.CoverTab[48081]++
													return &PrefixLogger{logger: logger, prefix: prefix}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:92
	// _ = "end of CoverTab[48081]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:93
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/grpclog/prefixLogger.go:93
var _ = _go_fuzz_dep_.CoverTab
