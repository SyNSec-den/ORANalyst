//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
package binarylog

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:19
)

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// NewLoggerFromConfigString reads the string and build a logger. It can be used
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
// to build a new logger and assign it to binarylog.Logger.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
// Example filter config strings:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "" Nothing will be logged
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "*" All headers and messages will be fully logged.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "*{h}" Only headers will be logged.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "*{m:256}" Only the first 256 bytes of each message will be logged.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "Foo/*" Logs every method in service Foo
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "Foo/*,-Foo/Bar" Logs every method in service Foo except method /Foo/Bar
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//   - "Foo/*,Foo/Bar{m:256}" Logs the first 256 bytes of each message in method
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//     /Foo/Bar, logs all headers and messages in every other method in service
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//     Foo.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
// If two configs exist for one certain method or service, the one specified
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:29
// later overrides the previous config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:45
func NewLoggerFromConfigString(s string) Logger {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:45
	_go_fuzz_dep_.CoverTab[68634]++
													if s == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:46
		_go_fuzz_dep_.CoverTab[68637]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:47
		// _ = "end of CoverTab[68637]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:48
		_go_fuzz_dep_.CoverTab[68638]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:48
		// _ = "end of CoverTab[68638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:48
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:48
	// _ = "end of CoverTab[68634]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:48
	_go_fuzz_dep_.CoverTab[68635]++
													l := newEmptyLogger()
													methods := strings.Split(s, ",")
													for _, method := range methods {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:51
		_go_fuzz_dep_.CoverTab[68639]++
														if err := l.fillMethodLoggerWithConfigString(method); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:52
			_go_fuzz_dep_.CoverTab[68640]++
															grpclogLogger.Warningf("failed to parse binary log config: %v", err)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:54
			// _ = "end of CoverTab[68640]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:55
			_go_fuzz_dep_.CoverTab[68641]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:55
			// _ = "end of CoverTab[68641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:55
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:55
		// _ = "end of CoverTab[68639]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:56
	// _ = "end of CoverTab[68635]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:56
	_go_fuzz_dep_.CoverTab[68636]++
													return l
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:57
	// _ = "end of CoverTab[68636]"
}

// fillMethodLoggerWithConfigString parses config, creates TruncatingMethodLogger and adds
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:60
// it to the right map in the logger.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:62
func (l *logger) fillMethodLoggerWithConfigString(config string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:62
	_go_fuzz_dep_.CoverTab[68642]++

													if config == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:64
		_go_fuzz_dep_.CoverTab[68649]++
														return errors.New("empty string is not a valid method binary logging config")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:65
		// _ = "end of CoverTab[68649]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:66
		_go_fuzz_dep_.CoverTab[68650]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:66
		// _ = "end of CoverTab[68650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:66
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:66
	// _ = "end of CoverTab[68642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:66
	_go_fuzz_dep_.CoverTab[68643]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:69
	if config[0] == '-' {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:69
		_go_fuzz_dep_.CoverTab[68651]++
														s, m, suffix, err := parseMethodConfigAndSuffix(config[1:])
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:71
			_go_fuzz_dep_.CoverTab[68656]++
															return fmt.Errorf("invalid config: %q, %v", config, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:72
			// _ = "end of CoverTab[68656]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:73
			_go_fuzz_dep_.CoverTab[68657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:73
			// _ = "end of CoverTab[68657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:73
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:73
		// _ = "end of CoverTab[68651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:73
		_go_fuzz_dep_.CoverTab[68652]++
														if m == "*" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:74
			_go_fuzz_dep_.CoverTab[68658]++
															return fmt.Errorf("invalid config: %q, %v", config, "* not allowed in blacklist config")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:75
			// _ = "end of CoverTab[68658]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:76
			_go_fuzz_dep_.CoverTab[68659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:76
			// _ = "end of CoverTab[68659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:76
		// _ = "end of CoverTab[68652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:76
		_go_fuzz_dep_.CoverTab[68653]++
														if suffix != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:77
			_go_fuzz_dep_.CoverTab[68660]++
															return fmt.Errorf("invalid config: %q, %v", config, "header/message limit not allowed in blacklist config")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:78
			// _ = "end of CoverTab[68660]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:79
			_go_fuzz_dep_.CoverTab[68661]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:79
			// _ = "end of CoverTab[68661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:79
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:79
		// _ = "end of CoverTab[68653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:79
		_go_fuzz_dep_.CoverTab[68654]++
														if err := l.setBlacklist(s + "/" + m); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:80
			_go_fuzz_dep_.CoverTab[68662]++
															return fmt.Errorf("invalid config: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:81
			// _ = "end of CoverTab[68662]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:82
			_go_fuzz_dep_.CoverTab[68663]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:82
			// _ = "end of CoverTab[68663]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:82
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:82
		// _ = "end of CoverTab[68654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:82
		_go_fuzz_dep_.CoverTab[68655]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:83
		// _ = "end of CoverTab[68655]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:84
		_go_fuzz_dep_.CoverTab[68664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:84
		// _ = "end of CoverTab[68664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:84
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:84
	// _ = "end of CoverTab[68643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:84
	_go_fuzz_dep_.CoverTab[68644]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:87
	if config[0] == '*' {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:87
		_go_fuzz_dep_.CoverTab[68665]++
														hdr, msg, err := parseHeaderMessageLengthConfig(config[1:])
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:89
			_go_fuzz_dep_.CoverTab[68668]++
															return fmt.Errorf("invalid config: %q, %v", config, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:90
			// _ = "end of CoverTab[68668]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:91
			_go_fuzz_dep_.CoverTab[68669]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:91
			// _ = "end of CoverTab[68669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:91
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:91
		// _ = "end of CoverTab[68665]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:91
		_go_fuzz_dep_.CoverTab[68666]++
														if err := l.setDefaultMethodLogger(&MethodLoggerConfig{Header: hdr, Message: msg}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:92
			_go_fuzz_dep_.CoverTab[68670]++
															return fmt.Errorf("invalid config: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:93
			// _ = "end of CoverTab[68670]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:94
			_go_fuzz_dep_.CoverTab[68671]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:94
			// _ = "end of CoverTab[68671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:94
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:94
		// _ = "end of CoverTab[68666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:94
		_go_fuzz_dep_.CoverTab[68667]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:95
		// _ = "end of CoverTab[68667]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:96
		_go_fuzz_dep_.CoverTab[68672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:96
		// _ = "end of CoverTab[68672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:96
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:96
	// _ = "end of CoverTab[68644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:96
	_go_fuzz_dep_.CoverTab[68645]++

													s, m, suffix, err := parseMethodConfigAndSuffix(config)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:99
		_go_fuzz_dep_.CoverTab[68673]++
														return fmt.Errorf("invalid config: %q, %v", config, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:100
		// _ = "end of CoverTab[68673]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:101
		_go_fuzz_dep_.CoverTab[68674]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:101
		// _ = "end of CoverTab[68674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:101
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:101
	// _ = "end of CoverTab[68645]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:101
	_go_fuzz_dep_.CoverTab[68646]++
													hdr, msg, err := parseHeaderMessageLengthConfig(suffix)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:103
		_go_fuzz_dep_.CoverTab[68675]++
														return fmt.Errorf("invalid header/message length config: %q, %v", suffix, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:104
		// _ = "end of CoverTab[68675]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:105
		_go_fuzz_dep_.CoverTab[68676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:105
		// _ = "end of CoverTab[68676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:105
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:105
	// _ = "end of CoverTab[68646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:105
	_go_fuzz_dep_.CoverTab[68647]++
													if m == "*" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:106
		_go_fuzz_dep_.CoverTab[68677]++
														if err := l.setServiceMethodLogger(s, &MethodLoggerConfig{Header: hdr, Message: msg}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:107
			_go_fuzz_dep_.CoverTab[68678]++
															return fmt.Errorf("invalid config: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:108
			// _ = "end of CoverTab[68678]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:109
			_go_fuzz_dep_.CoverTab[68679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:109
			// _ = "end of CoverTab[68679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:109
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:109
		// _ = "end of CoverTab[68677]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:110
		_go_fuzz_dep_.CoverTab[68680]++
														if err := l.setMethodMethodLogger(s+"/"+m, &MethodLoggerConfig{Header: hdr, Message: msg}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:111
			_go_fuzz_dep_.CoverTab[68681]++
															return fmt.Errorf("invalid config: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:112
			// _ = "end of CoverTab[68681]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:113
			_go_fuzz_dep_.CoverTab[68682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:113
			// _ = "end of CoverTab[68682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:113
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:113
		// _ = "end of CoverTab[68680]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:114
	// _ = "end of CoverTab[68647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:114
	_go_fuzz_dep_.CoverTab[68648]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:115
	// _ = "end of CoverTab[68648]"
}

const (
	// TODO: this const is only used by env_config now. But could be useful for
	// other config. Move to binarylog.go if necessary.
	maxUInt	= ^uint64(0)

	// For "p.s/m" plus any suffix. Suffix will be parsed again. See test for
	// expected output.
	longMethodConfigRegexpStr	= `^([\w./]+)/((?:\w+)|[*])(.+)?$`

	// For suffix from above, "{h:123,m:123}". See test for expected output.
	optionalLengthRegexpStr		= `(?::(\d+))?`	// Optional ":123".
	headerConfigRegexpStr		= `^{h` + optionalLengthRegexpStr + `}$`
	messageConfigRegexpStr		= `^{m` + optionalLengthRegexpStr + `}$`
	headerMessageConfigRegexpStr	= `^{h` + optionalLengthRegexpStr + `;m` + optionalLengthRegexpStr + `}$`
)

var (
	longMethodConfigRegexp		= regexp.MustCompile(longMethodConfigRegexpStr)
	headerConfigRegexp		= regexp.MustCompile(headerConfigRegexpStr)
	messageConfigRegexp		= regexp.MustCompile(messageConfigRegexpStr)
	headerMessageConfigRegexp	= regexp.MustCompile(headerMessageConfigRegexpStr)
)

// Turn "service/method{h;m}" into "service", "method", "{h;m}".
func parseMethodConfigAndSuffix(c string) (service, method, suffix string, _ error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:142
	_go_fuzz_dep_.CoverTab[68683]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:147
	match := longMethodConfigRegexp.FindStringSubmatch(c)
	if match == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:148
		_go_fuzz_dep_.CoverTab[68685]++
														return "", "", "", fmt.Errorf("%q contains invalid substring", c)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:149
		// _ = "end of CoverTab[68685]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:150
		_go_fuzz_dep_.CoverTab[68686]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:150
		// _ = "end of CoverTab[68686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:150
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:150
	// _ = "end of CoverTab[68683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:150
	_go_fuzz_dep_.CoverTab[68684]++
													service = match[1]
													method = match[2]
													suffix = match[3]
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:154
	// _ = "end of CoverTab[68684]"
}

// Turn "{h:123;m:345}" into 123, 345.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:157
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:157
// Return maxUInt if length is unspecified.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:160
func parseHeaderMessageLengthConfig(c string) (hdrLenStr, msgLenStr uint64, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:160
	_go_fuzz_dep_.CoverTab[68687]++
													if c == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:161
		_go_fuzz_dep_.CoverTab[68692]++
														return maxUInt, maxUInt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:162
		// _ = "end of CoverTab[68692]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:163
		_go_fuzz_dep_.CoverTab[68693]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:163
		// _ = "end of CoverTab[68693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:163
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:163
	// _ = "end of CoverTab[68687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:163
	_go_fuzz_dep_.CoverTab[68688]++

													if match := headerConfigRegexp.FindStringSubmatch(c); match != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:165
		_go_fuzz_dep_.CoverTab[68694]++
														if s := match[1]; s != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:166
			_go_fuzz_dep_.CoverTab[68696]++
															hdrLenStr, err = strconv.ParseUint(s, 10, 64)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:168
				_go_fuzz_dep_.CoverTab[68698]++
																return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:169
				// _ = "end of CoverTab[68698]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:170
				_go_fuzz_dep_.CoverTab[68699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:170
				// _ = "end of CoverTab[68699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:170
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:170
			// _ = "end of CoverTab[68696]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:170
			_go_fuzz_dep_.CoverTab[68697]++
															return hdrLenStr, 0, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:171
			// _ = "end of CoverTab[68697]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:172
			_go_fuzz_dep_.CoverTab[68700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:172
			// _ = "end of CoverTab[68700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:172
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:172
		// _ = "end of CoverTab[68694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:172
		_go_fuzz_dep_.CoverTab[68695]++
														return maxUInt, 0, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:173
		// _ = "end of CoverTab[68695]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:174
		_go_fuzz_dep_.CoverTab[68701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:174
		// _ = "end of CoverTab[68701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:174
	// _ = "end of CoverTab[68688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:174
	_go_fuzz_dep_.CoverTab[68689]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:177
	if match := messageConfigRegexp.FindStringSubmatch(c); match != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:177
		_go_fuzz_dep_.CoverTab[68702]++
														if s := match[1]; s != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:178
			_go_fuzz_dep_.CoverTab[68704]++
															msgLenStr, err = strconv.ParseUint(s, 10, 64)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:180
				_go_fuzz_dep_.CoverTab[68706]++
																return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:181
				// _ = "end of CoverTab[68706]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:182
				_go_fuzz_dep_.CoverTab[68707]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:182
				// _ = "end of CoverTab[68707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:182
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:182
			// _ = "end of CoverTab[68704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:182
			_go_fuzz_dep_.CoverTab[68705]++
															return 0, msgLenStr, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:183
			// _ = "end of CoverTab[68705]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:184
			_go_fuzz_dep_.CoverTab[68708]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:184
			// _ = "end of CoverTab[68708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:184
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:184
		// _ = "end of CoverTab[68702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:184
		_go_fuzz_dep_.CoverTab[68703]++
														return 0, maxUInt, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:185
		// _ = "end of CoverTab[68703]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:186
		_go_fuzz_dep_.CoverTab[68709]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:186
		// _ = "end of CoverTab[68709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:186
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:186
	// _ = "end of CoverTab[68689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:186
	_go_fuzz_dep_.CoverTab[68690]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:189
	if match := headerMessageConfigRegexp.FindStringSubmatch(c); match != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:189
		_go_fuzz_dep_.CoverTab[68710]++

														hdrLenStr = maxUInt
														msgLenStr = maxUInt
														if s := match[1]; s != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:193
			_go_fuzz_dep_.CoverTab[68713]++
															hdrLenStr, err = strconv.ParseUint(s, 10, 64)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:195
				_go_fuzz_dep_.CoverTab[68714]++
																return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:196
				// _ = "end of CoverTab[68714]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:197
				_go_fuzz_dep_.CoverTab[68715]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:197
				// _ = "end of CoverTab[68715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:197
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:197
			// _ = "end of CoverTab[68713]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:198
			_go_fuzz_dep_.CoverTab[68716]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:198
			// _ = "end of CoverTab[68716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:198
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:198
		// _ = "end of CoverTab[68710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:198
		_go_fuzz_dep_.CoverTab[68711]++
														if s := match[2]; s != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:199
			_go_fuzz_dep_.CoverTab[68717]++
															msgLenStr, err = strconv.ParseUint(s, 10, 64)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:201
				_go_fuzz_dep_.CoverTab[68718]++
																return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:202
				// _ = "end of CoverTab[68718]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:203
				_go_fuzz_dep_.CoverTab[68719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:203
				// _ = "end of CoverTab[68719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:203
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:203
			// _ = "end of CoverTab[68717]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:204
			_go_fuzz_dep_.CoverTab[68720]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:204
			// _ = "end of CoverTab[68720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:204
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:204
		// _ = "end of CoverTab[68711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:204
		_go_fuzz_dep_.CoverTab[68712]++
														return hdrLenStr, msgLenStr, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:205
		// _ = "end of CoverTab[68712]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:206
		_go_fuzz_dep_.CoverTab[68721]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:206
		// _ = "end of CoverTab[68721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:206
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:206
	// _ = "end of CoverTab[68690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:206
	_go_fuzz_dep_.CoverTab[68691]++
													return 0, 0, fmt.Errorf("%q contains invalid substring", c)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:207
	// _ = "end of CoverTab[68691]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:208
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/binarylog/env_config.go:208
var _ = _go_fuzz_dep_.CoverTab
