//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:19
// Package binarylog implementation binary logging as defined in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:19
// https://github.com/grpc/proposal/blob/master/A16-binary-logging.md.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
package binarylog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:21
)

import (
	"fmt"
	"os"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/grpcutil"
)

var grpclogLogger = grpclog.Component("binarylog")

// Logger specifies MethodLoggers for method names with a Log call that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:33
// takes a context.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:35
type Logger interface {
	GetMethodLogger(methodName string) MethodLogger
}

// binLogger is the global binary logger for the binary. One of this should be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:39
// built at init time from the configuration (environment variable or flags).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:39
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:39
// It is used to get a MethodLogger for each individual method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:43
var binLogger Logger

// SetLogger sets the binary logger.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:45
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:45
// Only call this at init time.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:48
func SetLogger(l Logger) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:48
	_go_fuzz_dep_.CoverTab[68579]++
													binLogger = l
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:49
	// _ = "end of CoverTab[68579]"
}

// GetLogger gets the binary logger.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:52
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:52
// Only call this at init time.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:55
func GetLogger() Logger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:55
	_go_fuzz_dep_.CoverTab[68580]++
													return binLogger
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:56
	// _ = "end of CoverTab[68580]"
}

// GetMethodLogger returns the MethodLogger for the given methodName.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:59
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:59
// methodName should be in the format of "/service/method".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:59
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:59
// Each MethodLogger returned by this method is a new instance. This is to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:59
// generate sequence id within the call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:65
func GetMethodLogger(methodName string) MethodLogger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:65
	_go_fuzz_dep_.CoverTab[68581]++
													if binLogger == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:66
		_go_fuzz_dep_.CoverTab[68583]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:67
		// _ = "end of CoverTab[68583]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:68
		_go_fuzz_dep_.CoverTab[68584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:68
		// _ = "end of CoverTab[68584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:68
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:68
	// _ = "end of CoverTab[68581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:68
	_go_fuzz_dep_.CoverTab[68582]++
													return binLogger.GetMethodLogger(methodName)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:69
	// _ = "end of CoverTab[68582]"
}

func init() {
	const envStr = "GRPC_BINARY_LOG_FILTER"
	configStr := os.Getenv(envStr)
	binLogger = NewLoggerFromConfigString(configStr)
}

// MethodLoggerConfig contains the setting for logging behavior of a method
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:78
// logger. Currently, it contains the max length of header and message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:80
type MethodLoggerConfig struct {
	// Max length of header and message.
	Header, Message uint64
}

// LoggerConfig contains the config for loggers to create method loggers.
type LoggerConfig struct {
	All		*MethodLoggerConfig
	Services	map[string]*MethodLoggerConfig
	Methods		map[string]*MethodLoggerConfig

	Blacklist	map[string]struct{}
}

type logger struct {
	config LoggerConfig
}

// NewLoggerFromConfig builds a logger with the given LoggerConfig.
func NewLoggerFromConfig(config LoggerConfig) Logger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:99
	_go_fuzz_dep_.CoverTab[68585]++
													return &logger{config: config}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:100
	// _ = "end of CoverTab[68585]"
}

// newEmptyLogger creates an empty logger. The map fields need to be filled in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:103
// using the set* functions.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:105
func newEmptyLogger() *logger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:105
	_go_fuzz_dep_.CoverTab[68586]++
													return &logger{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:106
	// _ = "end of CoverTab[68586]"
}

// Set method logger for "*".
func (l *logger) setDefaultMethodLogger(ml *MethodLoggerConfig) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:110
	_go_fuzz_dep_.CoverTab[68587]++
													if l.config.All != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:111
		_go_fuzz_dep_.CoverTab[68589]++
														return fmt.Errorf("conflicting global rules found")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:112
		// _ = "end of CoverTab[68589]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:113
		_go_fuzz_dep_.CoverTab[68590]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:113
		// _ = "end of CoverTab[68590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:113
	// _ = "end of CoverTab[68587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:113
	_go_fuzz_dep_.CoverTab[68588]++
													l.config.All = ml
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:115
	// _ = "end of CoverTab[68588]"
}

// Set method logger for "service/*".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:118
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:118
// New MethodLogger with same service overrides the old one.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:121
func (l *logger) setServiceMethodLogger(service string, ml *MethodLoggerConfig) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:121
	_go_fuzz_dep_.CoverTab[68591]++
													if _, ok := l.config.Services[service]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:122
		_go_fuzz_dep_.CoverTab[68594]++
														return fmt.Errorf("conflicting service rules for service %v found", service)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:123
		// _ = "end of CoverTab[68594]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:124
		_go_fuzz_dep_.CoverTab[68595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:124
		// _ = "end of CoverTab[68595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:124
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:124
	// _ = "end of CoverTab[68591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:124
	_go_fuzz_dep_.CoverTab[68592]++
													if l.config.Services == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:125
		_go_fuzz_dep_.CoverTab[68596]++
														l.config.Services = make(map[string]*MethodLoggerConfig)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:126
		// _ = "end of CoverTab[68596]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:127
		_go_fuzz_dep_.CoverTab[68597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:127
		// _ = "end of CoverTab[68597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:127
	// _ = "end of CoverTab[68592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:127
	_go_fuzz_dep_.CoverTab[68593]++
													l.config.Services[service] = ml
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:129
	// _ = "end of CoverTab[68593]"
}

// Set method logger for "service/method".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:132
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:132
// New MethodLogger with same method overrides the old one.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:135
func (l *logger) setMethodMethodLogger(method string, ml *MethodLoggerConfig) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:135
	_go_fuzz_dep_.CoverTab[68598]++
													if _, ok := l.config.Blacklist[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:136
		_go_fuzz_dep_.CoverTab[68602]++
														return fmt.Errorf("conflicting blacklist rules for method %v found", method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:137
		// _ = "end of CoverTab[68602]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:138
		_go_fuzz_dep_.CoverTab[68603]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:138
		// _ = "end of CoverTab[68603]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:138
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:138
	// _ = "end of CoverTab[68598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:138
	_go_fuzz_dep_.CoverTab[68599]++
													if _, ok := l.config.Methods[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:139
		_go_fuzz_dep_.CoverTab[68604]++
														return fmt.Errorf("conflicting method rules for method %v found", method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:140
		// _ = "end of CoverTab[68604]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:141
		_go_fuzz_dep_.CoverTab[68605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:141
		// _ = "end of CoverTab[68605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:141
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:141
	// _ = "end of CoverTab[68599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:141
	_go_fuzz_dep_.CoverTab[68600]++
													if l.config.Methods == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:142
		_go_fuzz_dep_.CoverTab[68606]++
														l.config.Methods = make(map[string]*MethodLoggerConfig)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:143
		// _ = "end of CoverTab[68606]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:144
		_go_fuzz_dep_.CoverTab[68607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:144
		// _ = "end of CoverTab[68607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:144
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:144
	// _ = "end of CoverTab[68600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:144
	_go_fuzz_dep_.CoverTab[68601]++
													l.config.Methods[method] = ml
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:146
	// _ = "end of CoverTab[68601]"
}

// Set blacklist method for "-service/method".
func (l *logger) setBlacklist(method string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:150
	_go_fuzz_dep_.CoverTab[68608]++
													if _, ok := l.config.Blacklist[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:151
		_go_fuzz_dep_.CoverTab[68612]++
														return fmt.Errorf("conflicting blacklist rules for method %v found", method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:152
		// _ = "end of CoverTab[68612]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:153
		_go_fuzz_dep_.CoverTab[68613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:153
		// _ = "end of CoverTab[68613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:153
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:153
	// _ = "end of CoverTab[68608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:153
	_go_fuzz_dep_.CoverTab[68609]++
													if _, ok := l.config.Methods[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:154
		_go_fuzz_dep_.CoverTab[68614]++
														return fmt.Errorf("conflicting method rules for method %v found", method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:155
		// _ = "end of CoverTab[68614]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:156
		_go_fuzz_dep_.CoverTab[68615]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:156
		// _ = "end of CoverTab[68615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:156
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:156
	// _ = "end of CoverTab[68609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:156
	_go_fuzz_dep_.CoverTab[68610]++
													if l.config.Blacklist == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:157
		_go_fuzz_dep_.CoverTab[68616]++
														l.config.Blacklist = make(map[string]struct{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:158
		// _ = "end of CoverTab[68616]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:159
		_go_fuzz_dep_.CoverTab[68617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:159
		// _ = "end of CoverTab[68617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:159
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:159
	// _ = "end of CoverTab[68610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:159
	_go_fuzz_dep_.CoverTab[68611]++
													l.config.Blacklist[method] = struct{}{}
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:161
	// _ = "end of CoverTab[68611]"
}

// getMethodLogger returns the MethodLogger for the given methodName.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:164
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:164
// methodName should be in the format of "/service/method".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:164
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:164
// Each MethodLogger returned by this method is a new instance. This is to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:164
// generate sequence id within the call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:170
func (l *logger) GetMethodLogger(methodName string) MethodLogger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:170
	_go_fuzz_dep_.CoverTab[68618]++
													s, m, err := grpcutil.ParseMethod(methodName)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:172
		_go_fuzz_dep_.CoverTab[68624]++
														grpclogLogger.Infof("binarylogging: failed to parse %q: %v", methodName, err)
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:174
		// _ = "end of CoverTab[68624]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:175
		_go_fuzz_dep_.CoverTab[68625]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:175
		// _ = "end of CoverTab[68625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:175
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:175
	// _ = "end of CoverTab[68618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:175
	_go_fuzz_dep_.CoverTab[68619]++
													if ml, ok := l.config.Methods[s+"/"+m]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:176
		_go_fuzz_dep_.CoverTab[68626]++
														return NewTruncatingMethodLogger(ml.Header, ml.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:177
		// _ = "end of CoverTab[68626]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:178
		_go_fuzz_dep_.CoverTab[68627]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:178
		// _ = "end of CoverTab[68627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:178
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:178
	// _ = "end of CoverTab[68619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:178
	_go_fuzz_dep_.CoverTab[68620]++
													if _, ok := l.config.Blacklist[s+"/"+m]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:179
		_go_fuzz_dep_.CoverTab[68628]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:180
		// _ = "end of CoverTab[68628]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:181
		_go_fuzz_dep_.CoverTab[68629]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:181
		// _ = "end of CoverTab[68629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:181
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:181
	// _ = "end of CoverTab[68620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:181
	_go_fuzz_dep_.CoverTab[68621]++
													if ml, ok := l.config.Services[s]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:182
		_go_fuzz_dep_.CoverTab[68630]++
														return NewTruncatingMethodLogger(ml.Header, ml.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:183
		// _ = "end of CoverTab[68630]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:184
		_go_fuzz_dep_.CoverTab[68631]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:184
		// _ = "end of CoverTab[68631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:184
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:184
	// _ = "end of CoverTab[68621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:184
	_go_fuzz_dep_.CoverTab[68622]++
													if l.config.All == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:185
		_go_fuzz_dep_.CoverTab[68632]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:186
		// _ = "end of CoverTab[68632]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:187
		_go_fuzz_dep_.CoverTab[68633]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:187
		// _ = "end of CoverTab[68633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:187
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:187
	// _ = "end of CoverTab[68622]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:187
	_go_fuzz_dep_.CoverTab[68623]++
													return NewTruncatingMethodLogger(l.config.All.Header, l.config.All.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:188
	// _ = "end of CoverTab[68623]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:189
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/binarylog.go:189
var _ = _go_fuzz_dep_.CoverTab
