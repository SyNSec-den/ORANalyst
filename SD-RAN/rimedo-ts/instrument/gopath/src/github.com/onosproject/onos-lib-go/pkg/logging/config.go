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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
package logging

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:15
)

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const configDir = ".onos"

// SinkType is the type of a sink
type SinkType string

func (t SinkType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:27
	_go_fuzz_dep_.CoverTab[131949]++
													return string(t)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:28
	// _ = "end of CoverTab[131949]"
}

const (
	// StdoutSinkType is the sink type for stdout
	StdoutSinkType	SinkType	= "stdout"
	// StderrSinkType is the sink type for stderr
	StderrSinkType	SinkType	= "stderr"
	// FileSinkType is the type for a file sink
	FileSinkType	SinkType	= "file"
	// KafkaSinkType is the sink type for the Kafka sink
	KafkaSinkType	SinkType	= "kafka"
)

// SinkEncoding is the encoding for a sink
type SinkEncoding string

func (e SinkEncoding) String() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:45
	_go_fuzz_dep_.CoverTab[131950]++
													return string(e)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:46
	// _ = "end of CoverTab[131950]"
}

const (
	// ConsoleEncoding is an encoding for outputs to the console
	ConsoleEncoding	SinkEncoding	= "console"
	// JSONEncoding is an encoding for JSON outputs
	JSONEncoding	SinkEncoding	= "json"
)

const (
	rootLoggerName = "root"
)

// Config logging configuration
type Config struct {
	Loggers	map[string]LoggerConfig	`yaml:"loggers"`
	Sinks	map[string]SinkConfig	`yaml:"sinks"`
}

// GetRootLogger returns the root logger configuration
func (c Config) GetRootLogger() LoggerConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:67
	_go_fuzz_dep_.CoverTab[131951]++
													root := c.Loggers[rootLoggerName]
													if root.Output == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:69
		_go_fuzz_dep_.CoverTab[131953]++
														defaultSink := ""
														root.Output = map[string]OutputConfig{
			"": {
				Sink: &defaultSink,
			},
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:75
		// _ = "end of CoverTab[131953]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:76
		_go_fuzz_dep_.CoverTab[131954]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:76
		// _ = "end of CoverTab[131954]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:76
	// _ = "end of CoverTab[131951]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:76
	_go_fuzz_dep_.CoverTab[131952]++
													return root
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:77
	// _ = "end of CoverTab[131952]"
}

// GetLoggers returns the configured loggers
func (c Config) GetLoggers() []LoggerConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:81
	_go_fuzz_dep_.CoverTab[131955]++
													loggers := make([]LoggerConfig, 0, len(c.Loggers))
													for name, logger := range c.Loggers {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:83
		_go_fuzz_dep_.CoverTab[131957]++
														if name != rootLoggerName {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:84
			_go_fuzz_dep_.CoverTab[131958]++
															logger.Name = name
															if logger.Output == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:86
				_go_fuzz_dep_.CoverTab[131960]++
																logger.Output = map[string]OutputConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:87
				// _ = "end of CoverTab[131960]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:88
				_go_fuzz_dep_.CoverTab[131961]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:88
				// _ = "end of CoverTab[131961]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:88
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:88
			// _ = "end of CoverTab[131958]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:88
			_go_fuzz_dep_.CoverTab[131959]++
															loggers = append(loggers, logger)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:89
			// _ = "end of CoverTab[131959]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:90
			_go_fuzz_dep_.CoverTab[131962]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:90
			// _ = "end of CoverTab[131962]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:90
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:90
		// _ = "end of CoverTab[131957]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:91
	// _ = "end of CoverTab[131955]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:91
	_go_fuzz_dep_.CoverTab[131956]++
													return loggers
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:92
	// _ = "end of CoverTab[131956]"
}

// GetLogger returns a logger by name
func (c Config) GetLogger(name string) (LoggerConfig, bool) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:96
	_go_fuzz_dep_.CoverTab[131963]++
													if name == rootLoggerName {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:97
		_go_fuzz_dep_.CoverTab[131966]++
														return LoggerConfig{}, false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:98
		// _ = "end of CoverTab[131966]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:99
		_go_fuzz_dep_.CoverTab[131967]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:99
		// _ = "end of CoverTab[131967]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:99
	// _ = "end of CoverTab[131963]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:99
	_go_fuzz_dep_.CoverTab[131964]++

													logger, ok := c.Loggers[name]
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:102
		_go_fuzz_dep_.CoverTab[131968]++
														logger.Name = name
														if logger.Output == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:104
			_go_fuzz_dep_.CoverTab[131970]++
															logger.Output = map[string]OutputConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:105
			// _ = "end of CoverTab[131970]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:106
			_go_fuzz_dep_.CoverTab[131971]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:106
			// _ = "end of CoverTab[131971]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:106
		// _ = "end of CoverTab[131968]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:106
		_go_fuzz_dep_.CoverTab[131969]++
														return logger, true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:107
		// _ = "end of CoverTab[131969]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:108
		_go_fuzz_dep_.CoverTab[131972]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:108
		// _ = "end of CoverTab[131972]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:108
	// _ = "end of CoverTab[131964]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:108
	_go_fuzz_dep_.CoverTab[131965]++
													return LoggerConfig{Output: map[string]OutputConfig{}}, false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:109
	// _ = "end of CoverTab[131965]"
}

// GetSinks returns the configured sinks
func (c Config) GetSinks() []SinkConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:113
	_go_fuzz_dep_.CoverTab[131973]++
													sinks := []SinkConfig{
		{
			Name: "",
		},
	}
	for name, sink := range c.Sinks {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:119
		_go_fuzz_dep_.CoverTab[131975]++
														sink.Name = name
														sinks = append(sinks, sink)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:121
		// _ = "end of CoverTab[131975]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:122
	// _ = "end of CoverTab[131973]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:122
	_go_fuzz_dep_.CoverTab[131974]++
													return sinks
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:123
	// _ = "end of CoverTab[131974]"
}

// GetSink returns a sink by name
func (c Config) GetSink(name string) (SinkConfig, bool) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:127
	_go_fuzz_dep_.CoverTab[131976]++
													if name == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:128
		_go_fuzz_dep_.CoverTab[131979]++
														return SinkConfig{
			Name: "",
		}, true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:131
		// _ = "end of CoverTab[131979]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:132
		_go_fuzz_dep_.CoverTab[131980]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:132
		// _ = "end of CoverTab[131980]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:132
	// _ = "end of CoverTab[131976]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:132
	_go_fuzz_dep_.CoverTab[131977]++
													sink, ok := c.Sinks[name]
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:134
		_go_fuzz_dep_.CoverTab[131981]++
														sink.Name = name
														return sink, true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:136
		// _ = "end of CoverTab[131981]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:137
		_go_fuzz_dep_.CoverTab[131982]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:137
		// _ = "end of CoverTab[131982]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:137
	// _ = "end of CoverTab[131977]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:137
	_go_fuzz_dep_.CoverTab[131978]++
													return SinkConfig{}, false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:138
	// _ = "end of CoverTab[131978]"
}

// LoggerConfig is the configuration for a logger
type LoggerConfig struct {
	Name	string			`yaml:"name"`
	Level	*string			`yaml:"level,omitempty"`
	Output	map[string]OutputConfig	`yaml:"output"`
}

// GetLevel returns the logger level
func (c LoggerConfig) GetLevel() Level {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:149
	_go_fuzz_dep_.CoverTab[131983]++
													level := c.Level
													if level != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:151
		_go_fuzz_dep_.CoverTab[131985]++
														return levelStringToLevel(*level)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:152
		// _ = "end of CoverTab[131985]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:153
		_go_fuzz_dep_.CoverTab[131986]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:153
		// _ = "end of CoverTab[131986]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:153
	// _ = "end of CoverTab[131983]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:153
	_go_fuzz_dep_.CoverTab[131984]++
													return ErrorLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:154
	// _ = "end of CoverTab[131984]"
}

// GetOutputs returns the logger outputs
func (c LoggerConfig) GetOutputs() []OutputConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:158
	_go_fuzz_dep_.CoverTab[131987]++
													outputs := c.Output
													outputsList := make([]OutputConfig, 0, len(outputs))
													for name, output := range outputs {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:161
		_go_fuzz_dep_.CoverTab[131989]++
														output.Name = name
														outputsList = append(outputsList, output)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:163
		// _ = "end of CoverTab[131989]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:164
	// _ = "end of CoverTab[131987]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:164
	_go_fuzz_dep_.CoverTab[131988]++
													return outputsList
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:165
	// _ = "end of CoverTab[131988]"
}

// GetOutput returns an output by name
func (c LoggerConfig) GetOutput(name string) (OutputConfig, bool) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:169
	_go_fuzz_dep_.CoverTab[131990]++
													output, ok := c.Output[name]
													if ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:171
		_go_fuzz_dep_.CoverTab[131992]++
														output.Name = name
														return output, true
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:173
		// _ = "end of CoverTab[131992]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:174
		_go_fuzz_dep_.CoverTab[131993]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:174
		// _ = "end of CoverTab[131993]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:174
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:174
	// _ = "end of CoverTab[131990]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:174
	_go_fuzz_dep_.CoverTab[131991]++
													return OutputConfig{}, false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:175
	// _ = "end of CoverTab[131991]"
}

// OutputConfig is the configuration for a sink output
type OutputConfig struct {
	Name	string	`yaml:"name"`
	Sink	*string	`yaml:"sink"`
	Level	*string	`yaml:"level,omitempty"`
}

// GetSink returns the output sink
func (c OutputConfig) GetSink() string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:186
	_go_fuzz_dep_.CoverTab[131994]++
													sink := c.Sink
													if sink != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:188
		_go_fuzz_dep_.CoverTab[131996]++
														return *sink
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:189
		// _ = "end of CoverTab[131996]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:190
		_go_fuzz_dep_.CoverTab[131997]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:190
		// _ = "end of CoverTab[131997]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:190
	// _ = "end of CoverTab[131994]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:190
	_go_fuzz_dep_.CoverTab[131995]++
													return ""
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:191
	// _ = "end of CoverTab[131995]"
}

// GetLevel returns the output level
func (c OutputConfig) GetLevel() Level {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:195
	_go_fuzz_dep_.CoverTab[131998]++
													level := c.Level
													if level != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:197
		_go_fuzz_dep_.CoverTab[132000]++
														return levelStringToLevel(*level)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:198
		// _ = "end of CoverTab[132000]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:199
		_go_fuzz_dep_.CoverTab[132001]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:199
		// _ = "end of CoverTab[132001]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:199
	// _ = "end of CoverTab[131998]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:199
	_go_fuzz_dep_.CoverTab[131999]++
													return DebugLevel
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:200
	// _ = "end of CoverTab[131999]"
}

// SinkConfig is the configuration for a sink
type SinkConfig struct {
	Name		string			`yaml:"name"`
	Type		*SinkType		`yaml:"type,omitempty"`
	Encoding	*SinkEncoding		`yaml:"encoding,omitempty"`
	Stdout		*StdoutSinkConfig	`yaml:"stdout,omitempty"`
	Stderr		*StderrSinkConfig	`yaml:"stderr,omitempty"`
	File		*FileSinkConfig		`yaml:"file,omitempty"`
	Kafka		*KafkaSinkConfig	`yaml:"kafka,omitempty"`
}

// GetType returns the sink type
func (c SinkConfig) GetType() SinkType {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:215
	_go_fuzz_dep_.CoverTab[132002]++
													sinkType := c.Type
													if sinkType != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:217
		_go_fuzz_dep_.CoverTab[132004]++
														return *sinkType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:218
		// _ = "end of CoverTab[132004]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:219
		_go_fuzz_dep_.CoverTab[132005]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:219
		// _ = "end of CoverTab[132005]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:219
	// _ = "end of CoverTab[132002]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:219
	_go_fuzz_dep_.CoverTab[132003]++
													return StdoutSinkType
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:220
	// _ = "end of CoverTab[132003]"
}

// GetEncoding returns the sink encoding
func (c SinkConfig) GetEncoding() SinkEncoding {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:224
	_go_fuzz_dep_.CoverTab[132006]++
													encoding := c.Encoding
													if encoding != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:226
		_go_fuzz_dep_.CoverTab[132008]++
														return *encoding
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:227
		// _ = "end of CoverTab[132008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:228
		_go_fuzz_dep_.CoverTab[132009]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:228
		// _ = "end of CoverTab[132009]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:228
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:228
	// _ = "end of CoverTab[132006]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:228
	_go_fuzz_dep_.CoverTab[132007]++
													return ConsoleEncoding
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:229
	// _ = "end of CoverTab[132007]"
}

// GetStdoutSinkConfig returns the stdout sink configuration
func (c SinkConfig) GetStdoutSinkConfig() StdoutSinkConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:233
	_go_fuzz_dep_.CoverTab[132010]++
													config := c.Stdout
													if config != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:235
		_go_fuzz_dep_.CoverTab[132012]++
														return *config
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:236
		// _ = "end of CoverTab[132012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:237
		_go_fuzz_dep_.CoverTab[132013]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:237
		// _ = "end of CoverTab[132013]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:237
	// _ = "end of CoverTab[132010]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:237
	_go_fuzz_dep_.CoverTab[132011]++
													return StdoutSinkConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:238
	// _ = "end of CoverTab[132011]"
}

// GetStderrSinkConfig returns the stderr sink configuration
func (c SinkConfig) GetStderrSinkConfig() StderrSinkConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:242
	_go_fuzz_dep_.CoverTab[132014]++
													config := c.Stderr
													if config != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:244
		_go_fuzz_dep_.CoverTab[132016]++
														return *config
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:245
		// _ = "end of CoverTab[132016]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:246
		_go_fuzz_dep_.CoverTab[132017]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:246
		// _ = "end of CoverTab[132017]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:246
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:246
	// _ = "end of CoverTab[132014]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:246
	_go_fuzz_dep_.CoverTab[132015]++
													return StderrSinkConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:247
	// _ = "end of CoverTab[132015]"
}

// GetFileSinkConfig returns the file sink configuration
func (c SinkConfig) GetFileSinkConfig() FileSinkConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:251
	_go_fuzz_dep_.CoverTab[132018]++
													config := c.File
													if config != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:253
		_go_fuzz_dep_.CoverTab[132020]++
														return *config
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:254
		// _ = "end of CoverTab[132020]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:255
		_go_fuzz_dep_.CoverTab[132021]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:255
		// _ = "end of CoverTab[132021]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:255
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:255
	// _ = "end of CoverTab[132018]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:255
	_go_fuzz_dep_.CoverTab[132019]++
													return FileSinkConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:256
	// _ = "end of CoverTab[132019]"
}

// GetKafkaSinkConfig returns the Kafka sink configuration
func (c SinkConfig) GetKafkaSinkConfig() KafkaSinkConfig {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:260
	_go_fuzz_dep_.CoverTab[132022]++
													config := c.Kafka
													if config != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:262
		_go_fuzz_dep_.CoverTab[132024]++
														return *config
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:263
		// _ = "end of CoverTab[132024]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:264
		_go_fuzz_dep_.CoverTab[132025]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:264
		// _ = "end of CoverTab[132025]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:264
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:264
	// _ = "end of CoverTab[132022]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:264
	_go_fuzz_dep_.CoverTab[132023]++
													return KafkaSinkConfig{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:265
	// _ = "end of CoverTab[132023]"
}

// StdoutSinkConfig is the configuration for an stdout sink
type StdoutSinkConfig struct {
}

// StderrSinkConfig is the configuration for an stderr sink
type StderrSinkConfig struct {
}

// FileSinkConfig is the configuration for a file sink
type FileSinkConfig struct {
	Path string `yaml:"path"`
}

// KafkaSinkConfig is the configuration for a Kafka sink
type KafkaSinkConfig struct {
	Topic	string		`yaml:"topic"`
	Key	string		`yaml:"key"`
	Brokers	[]string	`yaml:"brokers"`
}

// load loads the configuration
func load(config *Config) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:289
	_go_fuzz_dep_.CoverTab[132026]++
													home, err := homedir.Dir()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:291
		_go_fuzz_dep_.CoverTab[132030]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:292
		// _ = "end of CoverTab[132030]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:293
		_go_fuzz_dep_.CoverTab[132031]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:293
		// _ = "end of CoverTab[132031]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:293
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:293
	// _ = "end of CoverTab[132026]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:293
	_go_fuzz_dep_.CoverTab[132027]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:296
	viper.SetConfigName("logging")

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:299
	viper.AddConfigPath("./" + configDir + "/config")
	viper.AddConfigPath(home + "/" + configDir + "/config")
	viper.AddConfigPath("/etc/onos/config")
	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:306
		_go_fuzz_dep_.CoverTab[132032]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:307
		// _ = "end of CoverTab[132032]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:308
		_go_fuzz_dep_.CoverTab[132033]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:308
		// _ = "end of CoverTab[132033]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:308
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:308
	// _ = "end of CoverTab[132027]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:308
	_go_fuzz_dep_.CoverTab[132028]++

													err = viper.Unmarshal(config)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:311
		_go_fuzz_dep_.CoverTab[132034]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:312
		// _ = "end of CoverTab[132034]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:313
		_go_fuzz_dep_.CoverTab[132035]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:313
		// _ = "end of CoverTab[132035]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:313
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:313
	// _ = "end of CoverTab[132028]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:313
	_go_fuzz_dep_.CoverTab[132029]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:314
	// _ = "end of CoverTab[132029]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:315
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/config.go:315
var _ = _go_fuzz_dep_.CoverTab
