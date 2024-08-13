// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
package logging

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:15
)

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

var root *zapLogger

const nameSep = "/"

func init() {
	config := Config{}
	if err := load(&config); err != nil {
		panic(err)
	} else if err := configure(config); err != nil {
		panic(err)
	}
}

// configure configures the loggers
func configure(config Config) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:39
	_go_fuzz_dep_.CoverTab[132078]++
													rootLogger, err := newZapLogger(config, config.GetRootLogger())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:41
		_go_fuzz_dep_.CoverTab[132080]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:42
		// _ = "end of CoverTab[132080]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:43
		_go_fuzz_dep_.CoverTab[132081]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:43
		// _ = "end of CoverTab[132081]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:43
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:43
	// _ = "end of CoverTab[132078]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:43
	_go_fuzz_dep_.CoverTab[132079]++
													root = rootLogger
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:45
	// _ = "end of CoverTab[132079]"
}

// GetLogger gets a logger by name
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:48
// If no name is provided, the caller's package name will be used if available.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:48
// If a single name is provided, the ancestry will be determined by splitting the
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:48
// string on backslashes.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:48
// If multiple names are provided, the set of names defines the logger ancestry.
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:53
func GetLogger(names ...string) Logger {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:53
	_go_fuzz_dep_.CoverTab[132082]++
													if len(names) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:54
		_go_fuzz_dep_.CoverTab[132084]++
														pkg, ok := getCallerPackage()
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:56
			_go_fuzz_dep_.CoverTab[132086]++
															panic("could not retrieve logger package")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:57
			// _ = "end of CoverTab[132086]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:58
			_go_fuzz_dep_.CoverTab[132087]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:58
			// _ = "end of CoverTab[132087]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:58
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:58
		// _ = "end of CoverTab[132084]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:58
		_go_fuzz_dep_.CoverTab[132085]++
														names = []string{pkg}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:59
		// _ = "end of CoverTab[132085]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:60
		_go_fuzz_dep_.CoverTab[132088]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:60
		// _ = "end of CoverTab[132088]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:60
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:60
	// _ = "end of CoverTab[132082]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:60
	_go_fuzz_dep_.CoverTab[132083]++
													return root.GetLogger(names...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:61
	// _ = "end of CoverTab[132083]"
}

// getCallerPackage gets the package name of the calling function'ss caller
func getCallerPackage() (string, bool) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:65
	_go_fuzz_dep_.CoverTab[132089]++
													var pkg string
													pc, _, _, ok := runtime.Caller(2)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:68
		_go_fuzz_dep_.CoverTab[132092]++
														return pkg, false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:69
		// _ = "end of CoverTab[132092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:70
		_go_fuzz_dep_.CoverTab[132093]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:70
		// _ = "end of CoverTab[132093]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:70
	// _ = "end of CoverTab[132089]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:70
	_go_fuzz_dep_.CoverTab[132090]++
													parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
													if parts[len(parts)-2][0] == '(' {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:72
		_go_fuzz_dep_.CoverTab[132094]++
														pkg = strings.Join(parts[0:len(parts)-2], ".")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:73
		// _ = "end of CoverTab[132094]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:74
		_go_fuzz_dep_.CoverTab[132095]++
														pkg = strings.Join(parts[0:len(parts)-1], ".")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:75
		// _ = "end of CoverTab[132095]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:76
	// _ = "end of CoverTab[132090]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:76
	_go_fuzz_dep_.CoverTab[132091]++
													return pkg, true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:77
	// _ = "end of CoverTab[132091]"
}

// SetLevel sets the root logger level
func SetLevel(level Level) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:81
	_go_fuzz_dep_.CoverTab[132096]++
													root.SetLevel(level)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:82
	// _ = "end of CoverTab[132096]"
}

// Logger represents an abstract logging interface.
type Logger interface {
	Output

	// Name returns the logger name
	Name() string

	// GetLogger gets a descendant of this Logger
	GetLogger(names ...string) Logger

	// GetLevel returns the logger's level
	GetLevel() Level

	// SetLevel sets the logger's level
	SetLevel(level Level)
}

func newZapLogger(config Config, loggerConfig LoggerConfig) (*zapLogger, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:102
	_go_fuzz_dep_.CoverTab[132097]++
													var outputs []*zapOutput
													outputConfigs := loggerConfig.GetOutputs()
													outputs = make([]*zapOutput, len(outputConfigs))
													for i, outputConfig := range outputConfigs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:106
		_go_fuzz_dep_.CoverTab[132100]++
														var sinkConfig SinkConfig
														if outputConfig.Sink == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:108
			_go_fuzz_dep_.CoverTab[132104]++
															return nil, fmt.Errorf("output sink not configured for output %s", outputConfig.Name)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:109
			// _ = "end of CoverTab[132104]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:110
			_go_fuzz_dep_.CoverTab[132105]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:110
			// _ = "end of CoverTab[132105]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:110
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:110
		// _ = "end of CoverTab[132100]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:110
		_go_fuzz_dep_.CoverTab[132101]++
														sink, ok := config.GetSink(*outputConfig.Sink)
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:112
			_go_fuzz_dep_.CoverTab[132106]++
															panic(fmt.Sprintf("unknown sink %s", *outputConfig.Sink))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:113
			// _ = "end of CoverTab[132106]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:114
			_go_fuzz_dep_.CoverTab[132107]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:114
			// _ = "end of CoverTab[132107]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:114
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:114
		// _ = "end of CoverTab[132101]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:114
		_go_fuzz_dep_.CoverTab[132102]++
														sinkConfig = sink
														output, err := newZapOutput(loggerConfig, outputConfig, sinkConfig)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:117
			_go_fuzz_dep_.CoverTab[132108]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:118
			// _ = "end of CoverTab[132108]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:119
			_go_fuzz_dep_.CoverTab[132109]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:119
			// _ = "end of CoverTab[132109]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:119
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:119
		// _ = "end of CoverTab[132102]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:119
		_go_fuzz_dep_.CoverTab[132103]++
														outputs[i] = output
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:120
		// _ = "end of CoverTab[132103]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:121
	// _ = "end of CoverTab[132097]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:121
	_go_fuzz_dep_.CoverTab[132098]++

													var level *Level
													var defaultLevel *Level
													if loggerConfig.Level != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:125
		_go_fuzz_dep_.CoverTab[132110]++
														loggerLevel := loggerConfig.GetLevel()
														level = &loggerLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:127
		// _ = "end of CoverTab[132110]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:128
		_go_fuzz_dep_.CoverTab[132111]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:128
		// _ = "end of CoverTab[132111]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:128
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:128
	// _ = "end of CoverTab[132098]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:128
	_go_fuzz_dep_.CoverTab[132099]++

													logger := &zapLogger{
		config:		config,
		loggerConfig:	loggerConfig,
		children:	make(map[string]*zapLogger),
		outputs:	outputs,
	}
													logger.level.Store(level)
													logger.defaultLevel.Store(defaultLevel)
													return logger, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:138
	// _ = "end of CoverTab[132099]"
}

// zapLogger is the default Logger implementation
type zapLogger struct {
	config		Config
	loggerConfig	LoggerConfig
	children	map[string]*zapLogger
	outputs		[]*zapOutput
	mu		sync.RWMutex
	level		atomic.Value
	defaultLevel	atomic.Value
}

func (l *zapLogger) Name() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:152
	_go_fuzz_dep_.CoverTab[132112]++
													return l.loggerConfig.Name
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:153
	// _ = "end of CoverTab[132112]"
}

func (l *zapLogger) GetLogger(names ...string) Logger {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:156
	_go_fuzz_dep_.CoverTab[132113]++
													if len(names) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:157
		_go_fuzz_dep_.CoverTab[132116]++
														names = strings.Split(names[0], nameSep)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:158
		// _ = "end of CoverTab[132116]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:159
		_go_fuzz_dep_.CoverTab[132117]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:159
		// _ = "end of CoverTab[132117]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:159
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:159
	// _ = "end of CoverTab[132113]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:159
	_go_fuzz_dep_.CoverTab[132114]++

													logger := l
													for _, name := range names {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:162
		_go_fuzz_dep_.CoverTab[132118]++
														child, err := logger.getChild(name)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:164
			_go_fuzz_dep_.CoverTab[132120]++
															panic(err)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:165
			// _ = "end of CoverTab[132120]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:166
			_go_fuzz_dep_.CoverTab[132121]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:166
			// _ = "end of CoverTab[132121]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:166
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:166
		// _ = "end of CoverTab[132118]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:166
		_go_fuzz_dep_.CoverTab[132119]++
														logger = child
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:167
		// _ = "end of CoverTab[132119]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:168
	// _ = "end of CoverTab[132114]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:168
	_go_fuzz_dep_.CoverTab[132115]++
													return logger
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:169
	// _ = "end of CoverTab[132115]"
}

func (l *zapLogger) getChild(name string) (*zapLogger, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:172
	_go_fuzz_dep_.CoverTab[132122]++
													l.mu.RLock()
													child, ok := l.children[name]
													l.mu.RUnlock()
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:176
		_go_fuzz_dep_.CoverTab[132128]++
														return child, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:177
		// _ = "end of CoverTab[132128]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:178
		_go_fuzz_dep_.CoverTab[132129]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:178
		// _ = "end of CoverTab[132129]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:178
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:178
	// _ = "end of CoverTab[132122]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:178
	_go_fuzz_dep_.CoverTab[132123]++

													l.mu.Lock()
													defer l.mu.Unlock()

													child, ok = l.children[name]
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:184
		_go_fuzz_dep_.CoverTab[132130]++
														return child, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:185
		// _ = "end of CoverTab[132130]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:186
		_go_fuzz_dep_.CoverTab[132131]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:186
		// _ = "end of CoverTab[132131]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:186
	// _ = "end of CoverTab[132123]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:186
	_go_fuzz_dep_.CoverTab[132124]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:189
	qualifiedName := strings.Trim(fmt.Sprintf("%s%s%s", l.loggerConfig.Name, nameSep, name), nameSep)

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:192
	loggerConfig, ok := l.config.GetLogger(qualifiedName)
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:193
		_go_fuzz_dep_.CoverTab[132132]++
														loggerConfig = l.loggerConfig
														loggerConfig.Name = qualifiedName
														loggerConfig.Level = nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:196
		// _ = "end of CoverTab[132132]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:197
		_go_fuzz_dep_.CoverTab[132133]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:197
		// _ = "end of CoverTab[132133]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:197
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:197
	// _ = "end of CoverTab[132124]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:197
	_go_fuzz_dep_.CoverTab[132125]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:200
	for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:200
		_go_fuzz_dep_.CoverTab[132134]++
														outputConfig, ok := loggerConfig.GetOutput(output.config.Name)
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:202
			_go_fuzz_dep_.CoverTab[132135]++
															loggerConfig.Output[output.config.Name] = output.config
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:203
			// _ = "end of CoverTab[132135]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:204
			_go_fuzz_dep_.CoverTab[132136]++
															if outputConfig.Sink == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:205
				_go_fuzz_dep_.CoverTab[132139]++
																outputConfig.Sink = output.config.Sink
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:206
				// _ = "end of CoverTab[132139]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:207
				_go_fuzz_dep_.CoverTab[132140]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:207
				// _ = "end of CoverTab[132140]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:207
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:207
			// _ = "end of CoverTab[132136]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:207
			_go_fuzz_dep_.CoverTab[132137]++
															if outputConfig.Level == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:208
				_go_fuzz_dep_.CoverTab[132141]++
																outputConfig.Level = output.config.Level
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:209
				// _ = "end of CoverTab[132141]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:210
				_go_fuzz_dep_.CoverTab[132142]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:210
				// _ = "end of CoverTab[132142]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:210
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:210
			// _ = "end of CoverTab[132137]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:210
			_go_fuzz_dep_.CoverTab[132138]++
															loggerConfig.Output[outputConfig.Name] = outputConfig
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:211
			// _ = "end of CoverTab[132138]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:212
		// _ = "end of CoverTab[132134]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:213
	// _ = "end of CoverTab[132125]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:213
	_go_fuzz_dep_.CoverTab[132126]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:216
	logger, err := newZapLogger(l.config, loggerConfig)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:217
		_go_fuzz_dep_.CoverTab[132143]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:218
		// _ = "end of CoverTab[132143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:219
		_go_fuzz_dep_.CoverTab[132144]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:219
		// _ = "end of CoverTab[132144]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:219
	// _ = "end of CoverTab[132126]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:219
	_go_fuzz_dep_.CoverTab[132127]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:222
	logger.setDefaultLevel(l.GetLevel())
													l.children[name] = logger
													return logger, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:224
	// _ = "end of CoverTab[132127]"
}

func (l *zapLogger) GetLevel() Level {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:227
	_go_fuzz_dep_.CoverTab[132145]++
													level := l.level.Load().(*Level)
													if level != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:229
		_go_fuzz_dep_.CoverTab[132148]++
														return *level
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:230
		// _ = "end of CoverTab[132148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:231
		_go_fuzz_dep_.CoverTab[132149]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:231
		// _ = "end of CoverTab[132149]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:231
	// _ = "end of CoverTab[132145]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:231
	_go_fuzz_dep_.CoverTab[132146]++

													defaultLevel := l.defaultLevel.Load().(*Level)
													if defaultLevel != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:234
		_go_fuzz_dep_.CoverTab[132150]++
														return *defaultLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:235
		// _ = "end of CoverTab[132150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:236
		_go_fuzz_dep_.CoverTab[132151]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:236
		// _ = "end of CoverTab[132151]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:236
	// _ = "end of CoverTab[132146]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:236
	_go_fuzz_dep_.CoverTab[132147]++
													return EmptyLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:237
	// _ = "end of CoverTab[132147]"
}

func (l *zapLogger) SetLevel(level Level) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:240
	_go_fuzz_dep_.CoverTab[132152]++
													l.level.Store(&level)
													for _, child := range l.children {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:242
		_go_fuzz_dep_.CoverTab[132153]++
														child.setDefaultLevel(level)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:243
		// _ = "end of CoverTab[132153]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:244
	// _ = "end of CoverTab[132152]"
}

func (l *zapLogger) setDefaultLevel(level Level) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:247
	_go_fuzz_dep_.CoverTab[132154]++
													l.defaultLevel.Store(&level)
													if l.level.Load().(*Level) == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:249
		_go_fuzz_dep_.CoverTab[132155]++
														for _, child := range l.children {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:250
			_go_fuzz_dep_.CoverTab[132156]++
															child.setDefaultLevel(level)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:251
			// _ = "end of CoverTab[132156]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:252
		// _ = "end of CoverTab[132155]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:253
		_go_fuzz_dep_.CoverTab[132157]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:253
		// _ = "end of CoverTab[132157]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:253
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:253
	// _ = "end of CoverTab[132154]"
}

func (l *zapLogger) Debug(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:256
	_go_fuzz_dep_.CoverTab[132158]++
													if l.GetLevel() <= DebugLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:257
		_go_fuzz_dep_.CoverTab[132159]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:258
			_go_fuzz_dep_.CoverTab[132160]++
															output.Debug(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:259
			// _ = "end of CoverTab[132160]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:260
		// _ = "end of CoverTab[132159]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:261
		_go_fuzz_dep_.CoverTab[132161]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:261
		// _ = "end of CoverTab[132161]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:261
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:261
	// _ = "end of CoverTab[132158]"
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:264
	_go_fuzz_dep_.CoverTab[132162]++
													if l.GetLevel() <= DebugLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:265
		_go_fuzz_dep_.CoverTab[132163]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:266
			_go_fuzz_dep_.CoverTab[132164]++
															output.Debugf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:267
			// _ = "end of CoverTab[132164]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:268
		// _ = "end of CoverTab[132163]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:269
		_go_fuzz_dep_.CoverTab[132165]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:269
		// _ = "end of CoverTab[132165]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:269
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:269
	// _ = "end of CoverTab[132162]"
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:272
	_go_fuzz_dep_.CoverTab[132166]++
													if l.GetLevel() <= DebugLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:273
		_go_fuzz_dep_.CoverTab[132167]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:274
			_go_fuzz_dep_.CoverTab[132168]++
															output.Debugw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:275
			// _ = "end of CoverTab[132168]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:276
		// _ = "end of CoverTab[132167]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:277
		_go_fuzz_dep_.CoverTab[132169]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:277
		// _ = "end of CoverTab[132169]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:277
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:277
	// _ = "end of CoverTab[132166]"
}

func (l *zapLogger) Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:280
	_go_fuzz_dep_.CoverTab[132170]++
													if l.GetLevel() <= InfoLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:281
		_go_fuzz_dep_.CoverTab[132171]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:282
			_go_fuzz_dep_.CoverTab[132172]++
															output.Info(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:283
			// _ = "end of CoverTab[132172]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:284
		// _ = "end of CoverTab[132171]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:285
		_go_fuzz_dep_.CoverTab[132173]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:285
		// _ = "end of CoverTab[132173]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:285
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:285
	// _ = "end of CoverTab[132170]"
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:288
	_go_fuzz_dep_.CoverTab[132174]++
													if l.GetLevel() <= InfoLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:289
		_go_fuzz_dep_.CoverTab[132175]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:290
			_go_fuzz_dep_.CoverTab[132176]++
															output.Infof(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:291
			// _ = "end of CoverTab[132176]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:292
		// _ = "end of CoverTab[132175]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:293
		_go_fuzz_dep_.CoverTab[132177]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:293
		// _ = "end of CoverTab[132177]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:293
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:293
	// _ = "end of CoverTab[132174]"
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:296
	_go_fuzz_dep_.CoverTab[132178]++
													if l.GetLevel() <= InfoLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:297
		_go_fuzz_dep_.CoverTab[132179]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:298
			_go_fuzz_dep_.CoverTab[132180]++
															output.Infow(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:299
			// _ = "end of CoverTab[132180]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:300
		// _ = "end of CoverTab[132179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:301
		_go_fuzz_dep_.CoverTab[132181]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:301
		// _ = "end of CoverTab[132181]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:301
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:301
	// _ = "end of CoverTab[132178]"
}

func (l *zapLogger) Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:304
	_go_fuzz_dep_.CoverTab[132182]++
													if l.GetLevel() <= ErrorLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:305
		_go_fuzz_dep_.CoverTab[132183]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:306
			_go_fuzz_dep_.CoverTab[132184]++
															output.Error(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:307
			// _ = "end of CoverTab[132184]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:308
		// _ = "end of CoverTab[132183]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:309
		_go_fuzz_dep_.CoverTab[132185]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:309
		// _ = "end of CoverTab[132185]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:309
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:309
	// _ = "end of CoverTab[132182]"
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:312
	_go_fuzz_dep_.CoverTab[132186]++
													if l.GetLevel() <= ErrorLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:313
		_go_fuzz_dep_.CoverTab[132187]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:314
			_go_fuzz_dep_.CoverTab[132188]++
															output.Errorf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:315
			// _ = "end of CoverTab[132188]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:316
		// _ = "end of CoverTab[132187]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:317
		_go_fuzz_dep_.CoverTab[132189]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:317
		// _ = "end of CoverTab[132189]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:317
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:317
	// _ = "end of CoverTab[132186]"
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:320
	_go_fuzz_dep_.CoverTab[132190]++
													if l.GetLevel() <= ErrorLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:321
		_go_fuzz_dep_.CoverTab[132191]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:322
			_go_fuzz_dep_.CoverTab[132192]++
															output.Errorw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:323
			// _ = "end of CoverTab[132192]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:324
		// _ = "end of CoverTab[132191]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:325
		_go_fuzz_dep_.CoverTab[132193]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:325
		// _ = "end of CoverTab[132193]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:325
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:325
	// _ = "end of CoverTab[132190]"
}

func (l *zapLogger) Fatal(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:328
	_go_fuzz_dep_.CoverTab[132194]++
													if l.GetLevel() <= FatalLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:329
		_go_fuzz_dep_.CoverTab[132195]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:330
			_go_fuzz_dep_.CoverTab[132196]++
															output.Fatal(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:331
			// _ = "end of CoverTab[132196]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:332
		// _ = "end of CoverTab[132195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:333
		_go_fuzz_dep_.CoverTab[132197]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:333
		// _ = "end of CoverTab[132197]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:333
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:333
	// _ = "end of CoverTab[132194]"
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:336
	_go_fuzz_dep_.CoverTab[132198]++
													if l.GetLevel() <= FatalLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:337
		_go_fuzz_dep_.CoverTab[132199]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:338
			_go_fuzz_dep_.CoverTab[132200]++
															output.Fatalf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:339
			// _ = "end of CoverTab[132200]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:340
		// _ = "end of CoverTab[132199]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:341
		_go_fuzz_dep_.CoverTab[132201]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:341
		// _ = "end of CoverTab[132201]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:341
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:341
	// _ = "end of CoverTab[132198]"
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:344
	_go_fuzz_dep_.CoverTab[132202]++
													if l.GetLevel() <= FatalLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:345
		_go_fuzz_dep_.CoverTab[132203]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:346
			_go_fuzz_dep_.CoverTab[132204]++
															output.Fatalw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:347
			// _ = "end of CoverTab[132204]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:348
		// _ = "end of CoverTab[132203]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:349
		_go_fuzz_dep_.CoverTab[132205]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:349
		// _ = "end of CoverTab[132205]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:349
	// _ = "end of CoverTab[132202]"
}

func (l *zapLogger) Panic(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:352
	_go_fuzz_dep_.CoverTab[132206]++
													if l.GetLevel() <= PanicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:353
		_go_fuzz_dep_.CoverTab[132207]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:354
			_go_fuzz_dep_.CoverTab[132208]++
															output.Panic(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:355
			// _ = "end of CoverTab[132208]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:356
		// _ = "end of CoverTab[132207]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:357
		_go_fuzz_dep_.CoverTab[132209]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:357
		// _ = "end of CoverTab[132209]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:357
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:357
	// _ = "end of CoverTab[132206]"
}

func (l *zapLogger) Panicf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:360
	_go_fuzz_dep_.CoverTab[132210]++
													if l.GetLevel() <= PanicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:361
		_go_fuzz_dep_.CoverTab[132211]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:362
			_go_fuzz_dep_.CoverTab[132212]++
															output.Panicf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:363
			// _ = "end of CoverTab[132212]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:364
		// _ = "end of CoverTab[132211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:365
		_go_fuzz_dep_.CoverTab[132213]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:365
		// _ = "end of CoverTab[132213]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:365
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:365
	// _ = "end of CoverTab[132210]"
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:368
	_go_fuzz_dep_.CoverTab[132214]++
													if l.GetLevel() <= PanicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:369
		_go_fuzz_dep_.CoverTab[132215]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:370
			_go_fuzz_dep_.CoverTab[132216]++
															output.Panicw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:371
			// _ = "end of CoverTab[132216]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:372
		// _ = "end of CoverTab[132215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:373
		_go_fuzz_dep_.CoverTab[132217]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:373
		// _ = "end of CoverTab[132217]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:373
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:373
	// _ = "end of CoverTab[132214]"
}

func (l *zapLogger) DPanic(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:376
	_go_fuzz_dep_.CoverTab[132218]++
													if l.GetLevel() <= DPanicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:377
		_go_fuzz_dep_.CoverTab[132219]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:378
			_go_fuzz_dep_.CoverTab[132220]++
															output.DPanic(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:379
			// _ = "end of CoverTab[132220]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:380
		// _ = "end of CoverTab[132219]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:381
		_go_fuzz_dep_.CoverTab[132221]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:381
		// _ = "end of CoverTab[132221]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:381
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:381
	// _ = "end of CoverTab[132218]"
}

func (l *zapLogger) DPanicf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:384
	_go_fuzz_dep_.CoverTab[132222]++
													if l.GetLevel() <= DPanicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:385
		_go_fuzz_dep_.CoverTab[132223]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:386
			_go_fuzz_dep_.CoverTab[132224]++
															output.DPanicf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:387
			// _ = "end of CoverTab[132224]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:388
		// _ = "end of CoverTab[132223]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:389
		_go_fuzz_dep_.CoverTab[132225]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:389
		// _ = "end of CoverTab[132225]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:389
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:389
	// _ = "end of CoverTab[132222]"
}

func (l *zapLogger) DPanicw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:392
	_go_fuzz_dep_.CoverTab[132226]++
													if l.GetLevel() <= DPanicLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:393
		_go_fuzz_dep_.CoverTab[132227]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:394
			_go_fuzz_dep_.CoverTab[132228]++
															output.DPanicw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:395
			// _ = "end of CoverTab[132228]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:396
		// _ = "end of CoverTab[132227]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:397
		_go_fuzz_dep_.CoverTab[132229]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:397
		// _ = "end of CoverTab[132229]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:397
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:397
	// _ = "end of CoverTab[132226]"
}

func (l *zapLogger) Warn(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:400
	_go_fuzz_dep_.CoverTab[132230]++
													if l.GetLevel() <= WarnLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:401
		_go_fuzz_dep_.CoverTab[132231]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:402
			_go_fuzz_dep_.CoverTab[132232]++
															output.Warn(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:403
			// _ = "end of CoverTab[132232]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:404
		// _ = "end of CoverTab[132231]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:405
		_go_fuzz_dep_.CoverTab[132233]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:405
		// _ = "end of CoverTab[132233]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:405
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:405
	// _ = "end of CoverTab[132230]"
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:408
	_go_fuzz_dep_.CoverTab[132234]++
													if l.GetLevel() <= WarnLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:409
		_go_fuzz_dep_.CoverTab[132235]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:410
			_go_fuzz_dep_.CoverTab[132236]++
															output.Warnf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:411
			// _ = "end of CoverTab[132236]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:412
		// _ = "end of CoverTab[132235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:413
		_go_fuzz_dep_.CoverTab[132237]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:413
		// _ = "end of CoverTab[132237]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:413
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:413
	// _ = "end of CoverTab[132234]"
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:416
	_go_fuzz_dep_.CoverTab[132238]++
													if l.GetLevel() <= WarnLevel {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:417
		_go_fuzz_dep_.CoverTab[132239]++
														for _, output := range l.outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:418
			_go_fuzz_dep_.CoverTab[132240]++
															output.Warnw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:419
			// _ = "end of CoverTab[132240]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:420
		// _ = "end of CoverTab[132239]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:421
		_go_fuzz_dep_.CoverTab[132241]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:421
		// _ = "end of CoverTab[132241]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:421
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:421
	// _ = "end of CoverTab[132238]"
}

var _ Logger = &zapLogger{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:424
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/logger.go:424
var _ = _go_fuzz_dep_.CoverTab
