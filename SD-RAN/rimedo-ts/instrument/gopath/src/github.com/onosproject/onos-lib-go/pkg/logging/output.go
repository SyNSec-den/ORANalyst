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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
package logging

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:15
)

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"strings"
	"sync"
)

func newZapOutput(logger LoggerConfig, output OutputConfig, sink SinkConfig) (*zapOutput, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:26
	_go_fuzz_dep_.CoverTab[132242]++
													zapConfig := zap.Config{}
													zapConfig.Level = levelToAtomicLevel(output.GetLevel())
													zapConfig.Encoding = string(sink.GetEncoding())
													zapConfig.EncoderConfig.EncodeName = zapcore.FullNameEncoder
													zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
													zapConfig.EncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
													zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
													zapConfig.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
													zapConfig.EncoderConfig.NameKey = "logger"
													zapConfig.EncoderConfig.MessageKey = "message"
													zapConfig.EncoderConfig.LevelKey = "level"
													zapConfig.EncoderConfig.TimeKey = "timestamp"
													zapConfig.EncoderConfig.CallerKey = "caller"
													zapConfig.EncoderConfig.StacktraceKey = "trace"

													var encoder zapcore.Encoder
													switch sink.GetEncoding() {
	case ConsoleEncoding:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:44
		_go_fuzz_dep_.CoverTab[132249]++
														encoder = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:45
		// _ = "end of CoverTab[132249]"
	case JSONEncoding:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:46
		_go_fuzz_dep_.CoverTab[132250]++
														encoder = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:47
		// _ = "end of CoverTab[132250]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:47
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:47
		_go_fuzz_dep_.CoverTab[132251]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:47
		// _ = "end of CoverTab[132251]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:48
	// _ = "end of CoverTab[132242]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:48
	_go_fuzz_dep_.CoverTab[132243]++

													var path string
													switch sink.GetType() {
	case StdoutSinkType:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:52
		_go_fuzz_dep_.CoverTab[132252]++
														path = StdoutSinkType.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:53
		// _ = "end of CoverTab[132252]"
	case StderrSinkType:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:54
		_go_fuzz_dep_.CoverTab[132253]++
														path = StderrSinkType.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:55
		// _ = "end of CoverTab[132253]"
	case FileSinkType:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:56
		_go_fuzz_dep_.CoverTab[132254]++
														path = sink.GetFileSinkConfig().Path
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:57
		// _ = "end of CoverTab[132254]"
	case KafkaSinkType:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:58
		_go_fuzz_dep_.CoverTab[132255]++
														kafkaConfig := sink.GetKafkaSinkConfig()
														var rawQuery bytes.Buffer
														if kafkaConfig.Topic != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:61
			_go_fuzz_dep_.CoverTab[132259]++
															rawQuery.WriteString("topic=")
															rawQuery.WriteString(kafkaConfig.Topic)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:63
			// _ = "end of CoverTab[132259]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:64
			_go_fuzz_dep_.CoverTab[132260]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:64
			// _ = "end of CoverTab[132260]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:64
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:64
		// _ = "end of CoverTab[132255]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:64
		_go_fuzz_dep_.CoverTab[132256]++

														if kafkaConfig.Key != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:66
			_go_fuzz_dep_.CoverTab[132261]++
															rawQuery.WriteString("&")
															rawQuery.WriteString("key=")
															rawQuery.WriteString(kafkaConfig.Key)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:69
			// _ = "end of CoverTab[132261]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:70
			_go_fuzz_dep_.CoverTab[132262]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:70
			// _ = "end of CoverTab[132262]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:70
		// _ = "end of CoverTab[132256]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:70
		_go_fuzz_dep_.CoverTab[132257]++
														kafkaURL := url.URL{Scheme: KafkaSinkType.String(), Host: strings.Join(kafkaConfig.Brokers, ","), RawQuery: rawQuery.String()}
														path = kafkaURL.String()
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:72
		// _ = "end of CoverTab[132257]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:72
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:72
		_go_fuzz_dep_.CoverTab[132258]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:72
		// _ = "end of CoverTab[132258]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:73
	// _ = "end of CoverTab[132243]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:73
	_go_fuzz_dep_.CoverTab[132244]++

													writer, err := getWriter(path)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:76
		_go_fuzz_dep_.CoverTab[132263]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:77
		// _ = "end of CoverTab[132263]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:78
		_go_fuzz_dep_.CoverTab[132264]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:78
		// _ = "end of CoverTab[132264]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:78
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:78
	// _ = "end of CoverTab[132244]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:78
	_go_fuzz_dep_.CoverTab[132245]++

													atomLevel := zap.AtomicLevel{}
													switch output.GetLevel() {
	case DebugLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:82
		_go_fuzz_dep_.CoverTab[132265]++
														atomLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:83
		// _ = "end of CoverTab[132265]"
	case InfoLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:84
		_go_fuzz_dep_.CoverTab[132266]++
														atomLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:85
		// _ = "end of CoverTab[132266]"
	case WarnLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:86
		_go_fuzz_dep_.CoverTab[132267]++
														atomLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:87
		// _ = "end of CoverTab[132267]"
	case ErrorLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:88
		_go_fuzz_dep_.CoverTab[132268]++
														atomLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:89
		// _ = "end of CoverTab[132268]"
	case PanicLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:90
		_go_fuzz_dep_.CoverTab[132269]++
														atomLevel = zap.NewAtomicLevelAt(zapcore.PanicLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:91
		// _ = "end of CoverTab[132269]"
	case FatalLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:92
		_go_fuzz_dep_.CoverTab[132270]++
														atomLevel = zap.NewAtomicLevelAt(zapcore.FatalLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:93
		// _ = "end of CoverTab[132270]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:93
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:93
		_go_fuzz_dep_.CoverTab[132271]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:93
		// _ = "end of CoverTab[132271]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:94
	// _ = "end of CoverTab[132245]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:94
	_go_fuzz_dep_.CoverTab[132246]++

													zapLogger, err := zapConfig.Build(zap.AddCallerSkip(2))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:97
		_go_fuzz_dep_.CoverTab[132272]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:98
		// _ = "end of CoverTab[132272]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:99
		_go_fuzz_dep_.CoverTab[132273]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:99
		// _ = "end of CoverTab[132273]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:99
	// _ = "end of CoverTab[132246]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:99
	_go_fuzz_dep_.CoverTab[132247]++

													zapLogger = zapLogger.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:103
				_go_fuzz_dep_.CoverTab[132274]++
																return zapcore.NewCore(encoder, writer, &atomLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:104
				// _ = "end of CoverTab[132274]"
			}))
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:105
	// _ = "end of CoverTab[132247]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:105
	_go_fuzz_dep_.CoverTab[132248]++
													return &zapOutput{
		config:	output,
		logger:	zapLogger.Named(logger.Name),
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:109
	// _ = "end of CoverTab[132248]"
}

var writers = make(map[string]zapcore.WriteSyncer)
var writersMu = &sync.Mutex{}

func getWriter(url string) (zapcore.WriteSyncer, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:115
	_go_fuzz_dep_.CoverTab[132275]++
													writersMu.Lock()
													defer writersMu.Unlock()
													writer, ok := writers[url]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:119
		_go_fuzz_dep_.CoverTab[132277]++
														ws, _, err := zap.Open(url)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:121
			_go_fuzz_dep_.CoverTab[132279]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:122
			// _ = "end of CoverTab[132279]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:123
			_go_fuzz_dep_.CoverTab[132280]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:123
			// _ = "end of CoverTab[132280]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:123
		// _ = "end of CoverTab[132277]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:123
		_go_fuzz_dep_.CoverTab[132278]++
														writer = ws
														writers[url] = writer
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:125
		// _ = "end of CoverTab[132278]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:126
		_go_fuzz_dep_.CoverTab[132281]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:126
		// _ = "end of CoverTab[132281]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:126
	// _ = "end of CoverTab[132275]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:126
	_go_fuzz_dep_.CoverTab[132276]++
													return writer, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:127
	// _ = "end of CoverTab[132276]"
}

// Output is a logging output
type Output interface {
	Debug(...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Info(...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Error(...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Fatal(...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Panic(...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	DPanic(...interface{})
	DPanicf(template string, args ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})

	Warn(...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
}

// zapOutput is a logging output implementation
type zapOutput struct {
	config	OutputConfig
	logger	*zap.Logger
}

func (o *zapOutput) Debug(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:167
	_go_fuzz_dep_.CoverTab[132282]++
													o.logger.Sugar().Debug(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:168
	// _ = "end of CoverTab[132282]"
}

func (o *zapOutput) Debugf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:171
	_go_fuzz_dep_.CoverTab[132283]++
													o.logger.Sugar().Debugf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:172
	// _ = "end of CoverTab[132283]"
}

func (o *zapOutput) Debugw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:175
	_go_fuzz_dep_.CoverTab[132284]++
													o.logger.Sugar().Debugw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:176
	// _ = "end of CoverTab[132284]"
}

func (o *zapOutput) Info(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:179
	_go_fuzz_dep_.CoverTab[132285]++
													o.logger.Sugar().Info(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:180
	// _ = "end of CoverTab[132285]"
}

func (o *zapOutput) Infof(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:183
	_go_fuzz_dep_.CoverTab[132286]++
													o.logger.Sugar().Infof(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:184
	// _ = "end of CoverTab[132286]"
}

func (o *zapOutput) Infow(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:187
	_go_fuzz_dep_.CoverTab[132287]++
													o.logger.Sugar().Infow(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:188
	// _ = "end of CoverTab[132287]"
}

func (o *zapOutput) Error(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:191
	_go_fuzz_dep_.CoverTab[132288]++
													o.logger.Sugar().Error(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:192
	// _ = "end of CoverTab[132288]"
}

func (o *zapOutput) Errorf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:195
	_go_fuzz_dep_.CoverTab[132289]++
													o.logger.Sugar().Errorf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:196
	// _ = "end of CoverTab[132289]"
}

func (o *zapOutput) Errorw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:199
	_go_fuzz_dep_.CoverTab[132290]++
													o.logger.Sugar().Errorw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:200
	// _ = "end of CoverTab[132290]"
}

func (o *zapOutput) Fatal(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:203
	_go_fuzz_dep_.CoverTab[132291]++
													o.logger.Sugar().Fatal(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:204
	// _ = "end of CoverTab[132291]"
}

func (o *zapOutput) Fatalf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:207
	_go_fuzz_dep_.CoverTab[132292]++
													o.logger.Sugar().Fatalf(template, args)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:208
	// _ = "end of CoverTab[132292]"
}

func (o *zapOutput) Fatalw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:211
	_go_fuzz_dep_.CoverTab[132293]++
													o.logger.Sugar().Fatalw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:212
	// _ = "end of CoverTab[132293]"
}

func (o *zapOutput) Panic(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:215
	_go_fuzz_dep_.CoverTab[132294]++
													o.logger.Sugar().Panic(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:216
	// _ = "end of CoverTab[132294]"
}

func (o *zapOutput) Panicf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:219
	_go_fuzz_dep_.CoverTab[132295]++
													o.logger.Sugar().Panicf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:220
	// _ = "end of CoverTab[132295]"
}

func (o *zapOutput) Panicw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:223
	_go_fuzz_dep_.CoverTab[132296]++
													o.logger.Sugar().Panicw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:224
	// _ = "end of CoverTab[132296]"
}

func (o *zapOutput) DPanic(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:227
	_go_fuzz_dep_.CoverTab[132297]++
													o.logger.Sugar().DPanic(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:228
	// _ = "end of CoverTab[132297]"
}

func (o *zapOutput) DPanicf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:231
	_go_fuzz_dep_.CoverTab[132298]++
													o.logger.Sugar().DPanicf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:232
	// _ = "end of CoverTab[132298]"
}

func (o *zapOutput) DPanicw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:235
	_go_fuzz_dep_.CoverTab[132299]++
													o.logger.Sugar().DPanicw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:236
	// _ = "end of CoverTab[132299]"
}

func (o *zapOutput) Warn(args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:239
	_go_fuzz_dep_.CoverTab[132300]++
													o.logger.Sugar().Warn(args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:240
	// _ = "end of CoverTab[132300]"
}

func (o *zapOutput) Warnf(template string, args ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:243
	_go_fuzz_dep_.CoverTab[132301]++
													o.logger.Sugar().Warnf(template, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:244
	// _ = "end of CoverTab[132301]"
}

func (o *zapOutput) Warnw(msg string, keysAndValues ...interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:247
	_go_fuzz_dep_.CoverTab[132302]++
													o.logger.Sugar().Warnw(msg, keysAndValues...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:248
	// _ = "end of CoverTab[132302]"
}

var _ Output = &zapOutput{}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:251
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/output.go:251
var _ = _go_fuzz_dep_.CoverTab
