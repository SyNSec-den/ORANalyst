// Copyright 2019-present Open Networking Foundation.
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

//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
package errors

//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:15
)

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Type is an error type
type Type int

const (
	// Unknown is an unknown error type
	Unknown	Type	= iota
	// Canceled indicates a request context was canceled
	Canceled
	// NotFound indicates a resource was not found
	NotFound
	// AlreadyExists indicates a resource already exists
	AlreadyExists
	// Unauthorized indicates access to a resource is not authorized
	Unauthorized
	// Forbidden indicates the operation requested to be performed on a resource is forbidden
	Forbidden
	// Conflict indicates a conflict occurred during concurrent modifications to a resource
	Conflict
	// Invalid indicates a message or request is invalid
	Invalid
	// Unavailable indicates a service is not available
	Unavailable
	// NotSupported indicates a method is not supported
	NotSupported
	// Timeout indicates a request timed out
	Timeout
	// Internal indicates an unexpected internal error occurred
	Internal
)

// TypedError is an typed error
type TypedError struct {
	// Type is the error type
	Type	Type
	// Message is the error message
	Message	string
}

func (e *TypedError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:62
	_go_fuzz_dep_.CoverTab[114519]++
														return e.Message
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:63
	// _ = "end of CoverTab[114519]"
}

var _ error = &TypedError{}

// From returns the given gRPC error as an Atomix error
func From(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:69
	_go_fuzz_dep_.CoverTab[114520]++
														if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:70
		_go_fuzz_dep_.CoverTab[114526]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:71
		// _ = "end of CoverTab[114526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:72
		_go_fuzz_dep_.CoverTab[114527]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:72
		// _ = "end of CoverTab[114527]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:72
	// _ = "end of CoverTab[114520]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:72
	_go_fuzz_dep_.CoverTab[114521]++

														if _, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:74
		_go_fuzz_dep_.CoverTab[114528]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:75
		// _ = "end of CoverTab[114528]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:76
		_go_fuzz_dep_.CoverTab[114529]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:76
		// _ = "end of CoverTab[114529]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:76
	// _ = "end of CoverTab[114521]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:76
	_go_fuzz_dep_.CoverTab[114522]++

														if err == context.Canceled {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:78
		_go_fuzz_dep_.CoverTab[114530]++
															return NewCanceled(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:79
		// _ = "end of CoverTab[114530]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:80
		_go_fuzz_dep_.CoverTab[114531]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:80
		// _ = "end of CoverTab[114531]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:80
	// _ = "end of CoverTab[114522]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:80
	_go_fuzz_dep_.CoverTab[114523]++
														if err == context.DeadlineExceeded {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:81
		_go_fuzz_dep_.CoverTab[114532]++
															return NewTimeout(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:82
		// _ = "end of CoverTab[114532]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:83
		_go_fuzz_dep_.CoverTab[114533]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:83
		// _ = "end of CoverTab[114533]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:83
	// _ = "end of CoverTab[114523]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:83
	_go_fuzz_dep_.CoverTab[114524]++

														status, ok := status.FromError(err)
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:86
		_go_fuzz_dep_.CoverTab[114534]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:87
		// _ = "end of CoverTab[114534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:88
		_go_fuzz_dep_.CoverTab[114535]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:88
		// _ = "end of CoverTab[114535]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:88
	// _ = "end of CoverTab[114524]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:88
	_go_fuzz_dep_.CoverTab[114525]++

														switch status.Code() {
	case codes.Unknown:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:91
		_go_fuzz_dep_.CoverTab[114536]++
															return NewUnknown(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:92
		// _ = "end of CoverTab[114536]"
	case codes.Canceled:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:93
		_go_fuzz_dep_.CoverTab[114537]++
															return NewCanceled(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:94
		// _ = "end of CoverTab[114537]"
	case codes.NotFound:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:95
		_go_fuzz_dep_.CoverTab[114538]++
															return NewNotFound(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:96
		// _ = "end of CoverTab[114538]"
	case codes.AlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:97
		_go_fuzz_dep_.CoverTab[114539]++
															return NewAlreadyExists(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:98
		// _ = "end of CoverTab[114539]"
	case codes.Unauthenticated:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:99
			_go_fuzz_dep_.CoverTab[114540]++
																return NewUnauthorized(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:100
		// _ = "end of CoverTab[114540]"
	case codes.PermissionDenied:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:101
		_go_fuzz_dep_.CoverTab[114541]++
																return NewForbidden(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:102
		// _ = "end of CoverTab[114541]"
	case codes.FailedPrecondition:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:103
		_go_fuzz_dep_.CoverTab[114542]++
																return NewConflict(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:104
		// _ = "end of CoverTab[114542]"
	case codes.InvalidArgument:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:105
		_go_fuzz_dep_.CoverTab[114543]++
																return NewInvalid(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:106
		// _ = "end of CoverTab[114543]"
	case codes.Unavailable:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:107
		_go_fuzz_dep_.CoverTab[114544]++
																return NewUnavailable(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:108
		// _ = "end of CoverTab[114544]"
	case codes.Unimplemented:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:109
		_go_fuzz_dep_.CoverTab[114545]++
																return NewNotSupported(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:110
		// _ = "end of CoverTab[114545]"
	case codes.DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:111
		_go_fuzz_dep_.CoverTab[114546]++
																return NewTimeout(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:112
		// _ = "end of CoverTab[114546]"
	case codes.Internal:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:113
		_go_fuzz_dep_.CoverTab[114547]++
																return NewInternal(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:114
		// _ = "end of CoverTab[114547]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:115
		_go_fuzz_dep_.CoverTab[114548]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:116
		// _ = "end of CoverTab[114548]"
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:117
	// _ = "end of CoverTab[114525]"
}

// Proto returns the given error as a gRPC error
func Proto(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:121
	_go_fuzz_dep_.CoverTab[114549]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:122
		_go_fuzz_dep_.CoverTab[114552]++
																return nil
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:123
		// _ = "end of CoverTab[114552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:124
		_go_fuzz_dep_.CoverTab[114553]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:124
		// _ = "end of CoverTab[114553]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:124
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:124
	// _ = "end of CoverTab[114549]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:124
	_go_fuzz_dep_.CoverTab[114550]++

															typed, ok := err.(*TypedError)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:127
		_go_fuzz_dep_.CoverTab[114554]++
																return status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:128
		// _ = "end of CoverTab[114554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:129
		_go_fuzz_dep_.CoverTab[114555]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:129
		// _ = "end of CoverTab[114555]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:129
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:129
	// _ = "end of CoverTab[114550]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:129
	_go_fuzz_dep_.CoverTab[114551]++

															switch typed.Type {
	case Unknown:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:132
		_go_fuzz_dep_.CoverTab[114556]++
																return status.Error(codes.Unknown, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:133
		// _ = "end of CoverTab[114556]"
	case Canceled:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:134
		_go_fuzz_dep_.CoverTab[114557]++
																return status.Error(codes.Canceled, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:135
		// _ = "end of CoverTab[114557]"
	case NotFound:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:136
		_go_fuzz_dep_.CoverTab[114558]++
																return status.Error(codes.NotFound, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:137
		// _ = "end of CoverTab[114558]"
	case AlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:138
		_go_fuzz_dep_.CoverTab[114559]++
																return status.Error(codes.AlreadyExists, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:139
		// _ = "end of CoverTab[114559]"
	case Unauthorized:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:140
		_go_fuzz_dep_.CoverTab[114560]++
																return status.Error(codes.Unauthenticated, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:141
		// _ = "end of CoverTab[114560]"
	case Forbidden:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:142
		_go_fuzz_dep_.CoverTab[114561]++
																return status.Error(codes.PermissionDenied, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:143
		// _ = "end of CoverTab[114561]"
	case Conflict:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:144
		_go_fuzz_dep_.CoverTab[114562]++
																return status.Error(codes.FailedPrecondition, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:145
		// _ = "end of CoverTab[114562]"
	case Invalid:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:146
		_go_fuzz_dep_.CoverTab[114563]++
																return status.Error(codes.InvalidArgument, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:147
		// _ = "end of CoverTab[114563]"
	case Unavailable:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:148
		_go_fuzz_dep_.CoverTab[114564]++
																return status.Error(codes.Unavailable, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:149
		// _ = "end of CoverTab[114564]"
	case NotSupported:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:150
		_go_fuzz_dep_.CoverTab[114565]++
																return status.Error(codes.Unimplemented, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:151
		// _ = "end of CoverTab[114565]"
	case Timeout:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:152
		_go_fuzz_dep_.CoverTab[114566]++
																return status.Error(codes.DeadlineExceeded, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:153
		// _ = "end of CoverTab[114566]"
	case Internal:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:154
		_go_fuzz_dep_.CoverTab[114567]++
																return status.Error(codes.Internal, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:155
		// _ = "end of CoverTab[114567]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:156
		_go_fuzz_dep_.CoverTab[114568]++
																return status.Error(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:157
		// _ = "end of CoverTab[114568]"
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:158
	// _ = "end of CoverTab[114551]"
}

// New creates a new typed error
func New(t Type, msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:162
	_go_fuzz_dep_.CoverTab[114569]++
															if len(args) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:163
		_go_fuzz_dep_.CoverTab[114571]++
																msg = fmt.Sprintf(msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:164
		// _ = "end of CoverTab[114571]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:165
		_go_fuzz_dep_.CoverTab[114572]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:165
		// _ = "end of CoverTab[114572]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:165
	// _ = "end of CoverTab[114569]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:165
	_go_fuzz_dep_.CoverTab[114570]++
															return &TypedError{
		Type:		t,
		Message:	msg,
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:169
	// _ = "end of CoverTab[114570]"
}

// NewUnknown returns a new Unknown error
func NewUnknown(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:173
	_go_fuzz_dep_.CoverTab[114573]++
															return New(Unknown, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:174
	// _ = "end of CoverTab[114573]"
}

// NewCanceled returns a new Canceled error
func NewCanceled(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:178
	_go_fuzz_dep_.CoverTab[114574]++
															return New(Canceled, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:179
	// _ = "end of CoverTab[114574]"
}

// NewNotFound returns a new NotFound error
func NewNotFound(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:183
	_go_fuzz_dep_.CoverTab[114575]++
															return New(NotFound, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:184
	// _ = "end of CoverTab[114575]"
}

// NewAlreadyExists returns a new AlreadyExists error
func NewAlreadyExists(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:188
	_go_fuzz_dep_.CoverTab[114576]++
															return New(AlreadyExists, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:189
	// _ = "end of CoverTab[114576]"
}

// NewUnauthorized returns a new Unauthorized error
func NewUnauthorized(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:193
	_go_fuzz_dep_.CoverTab[114577]++
															return New(Unauthorized, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:194
	// _ = "end of CoverTab[114577]"
}

// NewForbidden returns a new Forbidden error
func NewForbidden(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:198
	_go_fuzz_dep_.CoverTab[114578]++
															return New(Forbidden, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:199
	// _ = "end of CoverTab[114578]"
}

// NewConflict returns a new Conflict error
func NewConflict(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:203
	_go_fuzz_dep_.CoverTab[114579]++
															return New(Conflict, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:204
	// _ = "end of CoverTab[114579]"
}

// NewInvalid returns a new Invalid error
func NewInvalid(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:208
	_go_fuzz_dep_.CoverTab[114580]++
															return New(Invalid, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:209
	// _ = "end of CoverTab[114580]"
}

// NewUnavailable returns a new Unavailable error
func NewUnavailable(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:213
	_go_fuzz_dep_.CoverTab[114581]++
															return New(Unavailable, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:214
	// _ = "end of CoverTab[114581]"
}

// NewNotSupported returns a new NotSupported error
func NewNotSupported(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:218
	_go_fuzz_dep_.CoverTab[114582]++
															return New(NotSupported, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:219
	// _ = "end of CoverTab[114582]"
}

// NewTimeout returns a new Timeout error
func NewTimeout(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:223
	_go_fuzz_dep_.CoverTab[114583]++
															return New(Timeout, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:224
	// _ = "end of CoverTab[114583]"
}

// NewInternal returns a new Internal error
func NewInternal(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:228
	_go_fuzz_dep_.CoverTab[114584]++
															return New(Internal, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:229
	// _ = "end of CoverTab[114584]"
}

// TypeOf returns the type of the given error
func TypeOf(err error) Type {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:233
	_go_fuzz_dep_.CoverTab[114585]++
															if typed, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:234
		_go_fuzz_dep_.CoverTab[114587]++
																return typed.Type
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:235
		// _ = "end of CoverTab[114587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:236
		_go_fuzz_dep_.CoverTab[114588]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:236
		// _ = "end of CoverTab[114588]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:236
	// _ = "end of CoverTab[114585]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:236
	_go_fuzz_dep_.CoverTab[114586]++
															return Unknown
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:237
	// _ = "end of CoverTab[114586]"
}

// IsType checks whether the given error is of the given type
func IsType(err error, t Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:241
	_go_fuzz_dep_.CoverTab[114589]++
															if typed, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:242
		_go_fuzz_dep_.CoverTab[114591]++
																return typed.Type == t
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:243
		// _ = "end of CoverTab[114591]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:244
		_go_fuzz_dep_.CoverTab[114592]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:244
		// _ = "end of CoverTab[114592]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:244
	}
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:244
	// _ = "end of CoverTab[114589]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:244
	_go_fuzz_dep_.CoverTab[114590]++
															return false
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:245
	// _ = "end of CoverTab[114590]"
}

// IsUnknown checks whether the given error is an Unknown error
func IsUnknown(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:249
	_go_fuzz_dep_.CoverTab[114593]++
															return IsType(err, Unknown)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:250
	// _ = "end of CoverTab[114593]"
}

// IsCanceled checks whether the given error is an Canceled error
func IsCanceled(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:254
	_go_fuzz_dep_.CoverTab[114594]++
															return IsType(err, Canceled) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:255
		_go_fuzz_dep_.CoverTab[114595]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:255
		return err == context.Canceled
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:255
		// _ = "end of CoverTab[114595]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:255
	}()
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:255
	// _ = "end of CoverTab[114594]"
}

// IsNotFound checks whether the given error is a NotFound error
func IsNotFound(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:259
	_go_fuzz_dep_.CoverTab[114596]++
															return IsType(err, NotFound)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:260
	// _ = "end of CoverTab[114596]"
}

// IsAlreadyExists checks whether the given error is a AlreadyExists error
func IsAlreadyExists(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:264
	_go_fuzz_dep_.CoverTab[114597]++
															return IsType(err, AlreadyExists)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:265
	// _ = "end of CoverTab[114597]"
}

// IsUnauthorized checks whether the given error is a Unauthorized error
func IsUnauthorized(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:269
	_go_fuzz_dep_.CoverTab[114598]++
															return IsType(err, Unauthorized)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:270
	// _ = "end of CoverTab[114598]"
}

// IsForbidden checks whether the given error is a Forbidden error
func IsForbidden(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:274
	_go_fuzz_dep_.CoverTab[114599]++
															return IsType(err, Forbidden)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:275
	// _ = "end of CoverTab[114599]"
}

// IsConflict checks whether the given error is a Conflict error
func IsConflict(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:279
	_go_fuzz_dep_.CoverTab[114600]++
															return IsType(err, Conflict)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:280
	// _ = "end of CoverTab[114600]"
}

// IsInvalid checks whether the given error is an Invalid error
func IsInvalid(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:284
	_go_fuzz_dep_.CoverTab[114601]++
															return IsType(err, Invalid)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:285
	// _ = "end of CoverTab[114601]"
}

// IsUnavailable checks whether the given error is an Unavailable error
func IsUnavailable(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:289
	_go_fuzz_dep_.CoverTab[114602]++
															return IsType(err, Unavailable)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:290
	// _ = "end of CoverTab[114602]"
}

// IsNotSupported checks whether the given error is a NotSupported error
func IsNotSupported(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:294
	_go_fuzz_dep_.CoverTab[114603]++
															return IsType(err, NotSupported)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:295
	// _ = "end of CoverTab[114603]"
}

// IsTimeout checks whether the given error is a Timeout error
func IsTimeout(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:299
	_go_fuzz_dep_.CoverTab[114604]++
															return IsType(err, Timeout) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:300
		_go_fuzz_dep_.CoverTab[114605]++
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:300
		return err == context.DeadlineExceeded
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:300
		// _ = "end of CoverTab[114605]"
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:300
	}()
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:300
	// _ = "end of CoverTab[114604]"
}

// IsInternal checks whether the given error is an Internal error
func IsInternal(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:304
	_go_fuzz_dep_.CoverTab[114606]++
															return IsType(err, Internal)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:305
	// _ = "end of CoverTab[114606]"
}

//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:306
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/atomix/atomix-go-framework@v0.10.1/pkg/atomix/errors/errors.go:306
var _ = _go_fuzz_dep_.CoverTab
