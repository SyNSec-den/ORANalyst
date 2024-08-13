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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
package logging

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:15
)

import (
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"strings"

	"github.com/onosproject/onos-lib-go/api/logging"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewService returns a new device Service
func NewService() (service.Service, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:29
	_go_fuzz_dep_.CoverTab[132303]++
													return &Service{}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:30
	// _ = "end of CoverTab[132303]"
}

// Service is an implementation of C1 service.
type Service struct {
	service.Service
}

// Register registers the logging Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:39
	_go_fuzz_dep_.CoverTab[132304]++
													server := &Server{}
													logging.RegisterLoggerServer(r, server)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:41
	// _ = "end of CoverTab[132304]"
}

// Server implements the logging gRPC service
type Server struct {
}

func splitLoggerName(name string) []string {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:48
	_go_fuzz_dep_.CoverTab[132305]++
													names := strings.Split(name, nameSep)
													return names
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:50
	// _ = "end of CoverTab[132305]"
}

// GetLevel implements GetLevel rpc function to get a logger level
func (s *Server) GetLevel(ctx context.Context, req *logging.GetLevelRequest) (*logging.GetLevelResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:54
	_go_fuzz_dep_.CoverTab[132306]++

													name := req.GetLoggerName()
													if name == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:57
		_go_fuzz_dep_.CoverTab[132309]++
														return &logging.GetLevelResponse{}, errors.NewInvalid("precondition for get level request is failed")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:58
		// _ = "end of CoverTab[132309]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:59
		_go_fuzz_dep_.CoverTab[132310]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:59
		// _ = "end of CoverTab[132310]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:59
	// _ = "end of CoverTab[132306]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:59
	_go_fuzz_dep_.CoverTab[132307]++

													names := splitLoggerName(name)
													logger := GetLogger(names...)
													level := logger.GetLevel()

													var loggerLevel logging.Level
													switch level {
	case InfoLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:67
		_go_fuzz_dep_.CoverTab[132311]++
														loggerLevel = logging.Level_INFO
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:68
		// _ = "end of CoverTab[132311]"
	case DebugLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:69
		_go_fuzz_dep_.CoverTab[132312]++
														loggerLevel = logging.Level_DEBUG
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:70
		// _ = "end of CoverTab[132312]"
	case WarnLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:71
		_go_fuzz_dep_.CoverTab[132313]++
														loggerLevel = logging.Level_WARN
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:72
		// _ = "end of CoverTab[132313]"
	case ErrorLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:73
		_go_fuzz_dep_.CoverTab[132314]++
														loggerLevel = logging.Level_ERROR
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:74
		// _ = "end of CoverTab[132314]"
	case PanicLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:75
		_go_fuzz_dep_.CoverTab[132315]++
														loggerLevel = logging.Level_PANIC
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:76
		// _ = "end of CoverTab[132315]"
	case DPanicLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:77
		_go_fuzz_dep_.CoverTab[132316]++
														loggerLevel = logging.Level_DPANIC
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:78
		// _ = "end of CoverTab[132316]"
	case FatalLevel:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:79
		_go_fuzz_dep_.CoverTab[132317]++
														loggerLevel = logging.Level_FATAL
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:80
		// _ = "end of CoverTab[132317]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:80
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:80
		_go_fuzz_dep_.CoverTab[132318]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:80
		// _ = "end of CoverTab[132318]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:82
	// _ = "end of CoverTab[132307]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:82
	_go_fuzz_dep_.CoverTab[132308]++

													return &logging.GetLevelResponse{
		Level: loggerLevel,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:86
	// _ = "end of CoverTab[132308]"

}

// SetLevel implements SetLevel rpc function to set a logger level
func (s *Server) SetLevel(ctx context.Context, req *logging.SetLevelRequest) (*logging.SetLevelResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:91
	_go_fuzz_dep_.CoverTab[132319]++
													name := req.GetLoggerName()
													level := req.GetLevel()
													if name == "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:94
		_go_fuzz_dep_.CoverTab[132322]++
														return &logging.SetLevelResponse{
			ResponseStatus: logging.ResponseStatus_PRECONDITION_FAILED,
		}, errors.NewInvalid("precondition for set level request is failed")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:97
		// _ = "end of CoverTab[132322]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:98
		_go_fuzz_dep_.CoverTab[132323]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:98
		// _ = "end of CoverTab[132323]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:98
	// _ = "end of CoverTab[132319]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:98
	_go_fuzz_dep_.CoverTab[132320]++

													names := splitLoggerName(name)
													logger := GetLogger(names...)

													switch level {
	case logging.Level_INFO:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:104
		_go_fuzz_dep_.CoverTab[132324]++
														logger.SetLevel(InfoLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:105
		// _ = "end of CoverTab[132324]"
	case logging.Level_DEBUG:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:106
		_go_fuzz_dep_.CoverTab[132325]++
														logger.SetLevel(DebugLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:107
		// _ = "end of CoverTab[132325]"
	case logging.Level_WARN:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:108
		_go_fuzz_dep_.CoverTab[132326]++
														logger.SetLevel(WarnLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:109
		// _ = "end of CoverTab[132326]"
	case logging.Level_ERROR:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:110
		_go_fuzz_dep_.CoverTab[132327]++
														logger.SetLevel(ErrorLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:111
		// _ = "end of CoverTab[132327]"
	case logging.Level_PANIC:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:112
		_go_fuzz_dep_.CoverTab[132328]++
														logger.SetLevel(PanicLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:113
		// _ = "end of CoverTab[132328]"
	case logging.Level_DPANIC:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:114
		_go_fuzz_dep_.CoverTab[132329]++
														logger.SetLevel(DPanicLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:115
		// _ = "end of CoverTab[132329]"
	case logging.Level_FATAL:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:116
		_go_fuzz_dep_.CoverTab[132330]++
														logger.SetLevel(FatalLevel)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:117
		// _ = "end of CoverTab[132330]"

	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:119
		_go_fuzz_dep_.CoverTab[132331]++
														return &logging.SetLevelResponse{
			ResponseStatus: logging.ResponseStatus_PRECONDITION_FAILED,
		}, errors.NewNotSupported("the requested level is not supported")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:122
		// _ = "end of CoverTab[132331]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:124
	// _ = "end of CoverTab[132320]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:124
	_go_fuzz_dep_.CoverTab[132321]++

													return &logging.SetLevelResponse{
		ResponseStatus: logging.ResponseStatus_OK,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:128
	// _ = "end of CoverTab[132321]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:129
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/server.go:129
var _ = _go_fuzz_dep_.CoverTab
