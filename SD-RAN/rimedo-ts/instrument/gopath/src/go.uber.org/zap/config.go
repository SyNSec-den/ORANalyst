// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
package zap

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
import (
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:21
)

import (
	"fmt"
	"sort"
	"time"

	"go.uber.org/zap/zapcore"
)

// SamplingConfig sets a sampling strategy for the logger. Sampling caps the
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
// global CPU and I/O load that logging puts on your process while attempting
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
// to preserve a representative subset of your logs.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
// If specified, the Sampler will invoke the Hook after each decision.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
// Values configured here are per-second. See zapcore.NewSamplerWithOptions for
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:31
// details.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:39
type SamplingConfig struct {
	Initial		int						`json:"initial" yaml:"initial"`
	Thereafter	int						`json:"thereafter" yaml:"thereafter"`
	Hook		func(zapcore.Entry, zapcore.SamplingDecision)	`json:"-" yaml:"-"`
}

// Config offers a declarative way to construct a logger. It doesn't do
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// anything that can't be done with New, Options, and the various
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// zapcore.WriteSyncer and zapcore.Core wrappers, but it's a simpler way to
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// toggle common options.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// Note that Config intentionally supports only the most common options. More
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// unusual logging setups (logging to network connections or message queues,
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// splitting output between multiple files, etc.) are possible, but require
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// direct use of the zapcore package. For sample code, see the package-level
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// BasicConfiguration and AdvancedConfiguration examples.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// For an example showing runtime log level changes, see the documentation for
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:45
// AtomicLevel.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:58
type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level	AtomicLevel	`json:"level" yaml:"level"`
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development	bool	`json:"development" yaml:"development"`
	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller	bool	`json:"disableCaller" yaml:"disableCaller"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace	bool	`json:"disableStacktrace" yaml:"disableStacktrace"`
	// Sampling sets a sampling policy. A nil SamplingConfig disables sampling.
	Sampling	*SamplingConfig	`json:"sampling" yaml:"sampling"`
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding	string	`json:"encoding" yaml:"encoding"`
	// EncoderConfig sets options for the chosen encoder. See
	// zapcore.EncoderConfig for details.
	EncoderConfig	zapcore.EncoderConfig	`json:"encoderConfig" yaml:"encoderConfig"`
	// OutputPaths is a list of URLs or file paths to write logging output to.
	// See Open for details.
	OutputPaths	[]string	`json:"outputPaths" yaml:"outputPaths"`
	// ErrorOutputPaths is a list of URLs to write internal logger errors to.
	// The default is standard error.
	//
	// Note that this setting only affects internal errors; for sample code that
	// sends error-level logs to a different location from info- and debug-level
	// logs, see the package-level AdvancedConfiguration example.
	ErrorOutputPaths	[]string	`json:"errorOutputPaths" yaml:"errorOutputPaths"`
	// InitialFields is a collection of fields to add to the root logger.
	InitialFields	map[string]interface{}	`json:"initialFields" yaml:"initialFields"`
}

// NewProductionEncoderConfig returns an opinionated EncoderConfig for
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:96
// production environments.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:98
func NewProductionEncoderConfig() zapcore.EncoderConfig {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:98
	_go_fuzz_dep_.CoverTab[131344]++
									return zapcore.EncoderConfig{
		TimeKey:	"ts",
		LevelKey:	"level",
		NameKey:	"logger",
		CallerKey:	"caller",
		FunctionKey:	zapcore.OmitKey,
		MessageKey:	"msg",
		StacktraceKey:	"stacktrace",
		LineEnding:	zapcore.DefaultLineEnding,
		EncodeLevel:	zapcore.LowercaseLevelEncoder,
		EncodeTime:	zapcore.EpochTimeEncoder,
		EncodeDuration:	zapcore.SecondsDurationEncoder,
		EncodeCaller:	zapcore.ShortCallerEncoder,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:112
	// _ = "end of CoverTab[131344]"
}

// NewProductionConfig is a reasonable production logging configuration.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:115
// Logging is enabled at InfoLevel and above.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:115
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:115
// It uses a JSON encoder, writes to standard error, and enables sampling.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:115
// Stacktraces are automatically included on logs of ErrorLevel and above.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:120
func NewProductionConfig() Config {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:120
	_go_fuzz_dep_.CoverTab[131345]++
										return Config{
		Level:		NewAtomicLevelAt(InfoLevel),
		Development:	false,
		Sampling: &SamplingConfig{
			Initial:	100,
			Thereafter:	100,
		},
		Encoding:		"json",
		EncoderConfig:		NewProductionEncoderConfig(),
		OutputPaths:		[]string{"stderr"},
		ErrorOutputPaths:	[]string{"stderr"},
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:132
	// _ = "end of CoverTab[131345]"
}

// NewDevelopmentEncoderConfig returns an opinionated EncoderConfig for
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:135
// development environments.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:137
func NewDevelopmentEncoderConfig() zapcore.EncoderConfig {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:137
	_go_fuzz_dep_.CoverTab[131346]++
										return zapcore.EncoderConfig{

		TimeKey:	"T",
		LevelKey:	"L",
		NameKey:	"N",
		CallerKey:	"C",
		FunctionKey:	zapcore.OmitKey,
		MessageKey:	"M",
		StacktraceKey:	"S",
		LineEnding:	zapcore.DefaultLineEnding,
		EncodeLevel:	zapcore.CapitalLevelEncoder,
		EncodeTime:	zapcore.ISO8601TimeEncoder,
		EncodeDuration:	zapcore.StringDurationEncoder,
		EncodeCaller:	zapcore.ShortCallerEncoder,
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:152
	// _ = "end of CoverTab[131346]"
}

// NewDevelopmentConfig is a reasonable development logging configuration.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:155
// Logging is enabled at DebugLevel and above.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:155
//
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:155
// It enables development mode (which makes DPanicLevel logs panic), uses a
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:155
// console encoder, writes to standard error, and disables sampling.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:155
// Stacktraces are automatically included on logs of WarnLevel and above.
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:161
func NewDevelopmentConfig() Config {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:161
	_go_fuzz_dep_.CoverTab[131347]++
										return Config{
		Level:			NewAtomicLevelAt(DebugLevel),
		Development:		true,
		Encoding:		"console",
		EncoderConfig:		NewDevelopmentEncoderConfig(),
		OutputPaths:		[]string{"stderr"},
		ErrorOutputPaths:	[]string{"stderr"},
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:169
	// _ = "end of CoverTab[131347]"
}

// Build constructs a logger from the Config and Options.
func (cfg Config) Build(opts ...Option) (*Logger, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:173
	_go_fuzz_dep_.CoverTab[131348]++
										enc, err := cfg.buildEncoder()
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:175
		_go_fuzz_dep_.CoverTab[131353]++
											return nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:176
		// _ = "end of CoverTab[131353]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:177
		_go_fuzz_dep_.CoverTab[131354]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:177
		// _ = "end of CoverTab[131354]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:177
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:177
	// _ = "end of CoverTab[131348]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:177
	_go_fuzz_dep_.CoverTab[131349]++

										sink, errSink, err := cfg.openSinks()
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:180
		_go_fuzz_dep_.CoverTab[131355]++
											return nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:181
		// _ = "end of CoverTab[131355]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:182
		_go_fuzz_dep_.CoverTab[131356]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:182
		// _ = "end of CoverTab[131356]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:182
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:182
	// _ = "end of CoverTab[131349]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:182
	_go_fuzz_dep_.CoverTab[131350]++

										if cfg.Level == (AtomicLevel{}) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:184
		_go_fuzz_dep_.CoverTab[131357]++
											return nil, fmt.Errorf("missing Level")
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:185
		// _ = "end of CoverTab[131357]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:186
		_go_fuzz_dep_.CoverTab[131358]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:186
		// _ = "end of CoverTab[131358]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:186
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:186
	// _ = "end of CoverTab[131350]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:186
	_go_fuzz_dep_.CoverTab[131351]++

										log := New(
		zapcore.NewCore(enc, sink, cfg.Level),
		cfg.buildOptions(errSink)...,
	)
	if len(opts) > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:192
		_go_fuzz_dep_.CoverTab[131359]++
											log = log.WithOptions(opts...)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:193
		// _ = "end of CoverTab[131359]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:194
		_go_fuzz_dep_.CoverTab[131360]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:194
		// _ = "end of CoverTab[131360]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:194
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:194
	// _ = "end of CoverTab[131351]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:194
	_go_fuzz_dep_.CoverTab[131352]++
										return log, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:195
	// _ = "end of CoverTab[131352]"
}

func (cfg Config) buildOptions(errSink zapcore.WriteSyncer) []Option {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:198
	_go_fuzz_dep_.CoverTab[131361]++
										opts := []Option{ErrorOutput(errSink)}

										if cfg.Development {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:201
		_go_fuzz_dep_.CoverTab[131368]++
											opts = append(opts, Development())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:202
		// _ = "end of CoverTab[131368]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:203
		_go_fuzz_dep_.CoverTab[131369]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:203
		// _ = "end of CoverTab[131369]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:203
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:203
	// _ = "end of CoverTab[131361]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:203
	_go_fuzz_dep_.CoverTab[131362]++

										if !cfg.DisableCaller {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:205
		_go_fuzz_dep_.CoverTab[131370]++
											opts = append(opts, AddCaller())
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:206
		// _ = "end of CoverTab[131370]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:207
		_go_fuzz_dep_.CoverTab[131371]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:207
		// _ = "end of CoverTab[131371]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:207
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:207
	// _ = "end of CoverTab[131362]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:207
	_go_fuzz_dep_.CoverTab[131363]++

										stackLevel := ErrorLevel
										if cfg.Development {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:210
		_go_fuzz_dep_.CoverTab[131372]++
											stackLevel = WarnLevel
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:211
		// _ = "end of CoverTab[131372]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:212
		_go_fuzz_dep_.CoverTab[131373]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:212
		// _ = "end of CoverTab[131373]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:212
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:212
	// _ = "end of CoverTab[131363]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:212
	_go_fuzz_dep_.CoverTab[131364]++
										if !cfg.DisableStacktrace {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:213
		_go_fuzz_dep_.CoverTab[131374]++
											opts = append(opts, AddStacktrace(stackLevel))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:214
		// _ = "end of CoverTab[131374]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:215
		_go_fuzz_dep_.CoverTab[131375]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:215
		// _ = "end of CoverTab[131375]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:215
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:215
	// _ = "end of CoverTab[131364]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:215
	_go_fuzz_dep_.CoverTab[131365]++

										if scfg := cfg.Sampling; scfg != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:217
		_go_fuzz_dep_.CoverTab[131376]++
											opts = append(opts, WrapCore(func(core zapcore.Core) zapcore.Core {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:218
			_go_fuzz_dep_.CoverTab[131377]++
												var samplerOpts []zapcore.SamplerOption
												if scfg.Hook != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:220
				_go_fuzz_dep_.CoverTab[131379]++
													samplerOpts = append(samplerOpts, zapcore.SamplerHook(scfg.Hook))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:221
				// _ = "end of CoverTab[131379]"
			} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:222
				_go_fuzz_dep_.CoverTab[131380]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:222
				// _ = "end of CoverTab[131380]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:222
			}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:222
			// _ = "end of CoverTab[131377]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:222
			_go_fuzz_dep_.CoverTab[131378]++
												return zapcore.NewSamplerWithOptions(
				core,
				time.Second,
				cfg.Sampling.Initial,
				cfg.Sampling.Thereafter,
				samplerOpts...,
			)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:229
			// _ = "end of CoverTab[131378]"
		}))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:230
		// _ = "end of CoverTab[131376]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:231
		_go_fuzz_dep_.CoverTab[131381]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:231
		// _ = "end of CoverTab[131381]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:231
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:231
	// _ = "end of CoverTab[131365]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:231
	_go_fuzz_dep_.CoverTab[131366]++

										if len(cfg.InitialFields) > 0 {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:233
		_go_fuzz_dep_.CoverTab[131382]++
											fs := make([]Field, 0, len(cfg.InitialFields))
											keys := make([]string, 0, len(cfg.InitialFields))
											for k := range cfg.InitialFields {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:236
			_go_fuzz_dep_.CoverTab[131385]++
												keys = append(keys, k)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:237
			// _ = "end of CoverTab[131385]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:238
		// _ = "end of CoverTab[131382]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:238
		_go_fuzz_dep_.CoverTab[131383]++
											sort.Strings(keys)
											for _, k := range keys {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:240
			_go_fuzz_dep_.CoverTab[131386]++
												fs = append(fs, Any(k, cfg.InitialFields[k]))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:241
			// _ = "end of CoverTab[131386]"
		}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:242
		// _ = "end of CoverTab[131383]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:242
		_go_fuzz_dep_.CoverTab[131384]++
											opts = append(opts, Fields(fs...))
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:243
		// _ = "end of CoverTab[131384]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:244
		_go_fuzz_dep_.CoverTab[131387]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:244
		// _ = "end of CoverTab[131387]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:244
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:244
	// _ = "end of CoverTab[131366]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:244
	_go_fuzz_dep_.CoverTab[131367]++

										return opts
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:246
	// _ = "end of CoverTab[131367]"
}

func (cfg Config) openSinks() (zapcore.WriteSyncer, zapcore.WriteSyncer, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:249
	_go_fuzz_dep_.CoverTab[131388]++
										sink, closeOut, err := Open(cfg.OutputPaths...)
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:251
		_go_fuzz_dep_.CoverTab[131391]++
											return nil, nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:252
		// _ = "end of CoverTab[131391]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:253
		_go_fuzz_dep_.CoverTab[131392]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:253
		// _ = "end of CoverTab[131392]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:253
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:253
	// _ = "end of CoverTab[131388]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:253
	_go_fuzz_dep_.CoverTab[131389]++
										errSink, _, err := Open(cfg.ErrorOutputPaths...)
										if err != nil {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:255
		_go_fuzz_dep_.CoverTab[131393]++
											closeOut()
											return nil, nil, err
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:257
		// _ = "end of CoverTab[131393]"
	} else {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:258
		_go_fuzz_dep_.CoverTab[131394]++
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:258
		// _ = "end of CoverTab[131394]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:258
	}
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:258
	// _ = "end of CoverTab[131389]"
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:258
	_go_fuzz_dep_.CoverTab[131390]++
										return sink, errSink, nil
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:259
	// _ = "end of CoverTab[131390]"
}

func (cfg Config) buildEncoder() (zapcore.Encoder, error) {
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:262
	_go_fuzz_dep_.CoverTab[131395]++
										return newEncoder(cfg.Encoding, cfg.EncoderConfig)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:263
	// _ = "end of CoverTab[131395]"
}

//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:264
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/go.uber.org/zap@v1.17.0/config.go:264
var _ = _go_fuzz_dep_.CoverTab
