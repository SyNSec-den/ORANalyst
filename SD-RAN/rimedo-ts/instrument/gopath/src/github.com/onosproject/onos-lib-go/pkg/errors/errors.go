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

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
package errors

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:15
)

import (
	"fmt"
	atomixerrors "github.com/atomix/atomix-go-framework/pkg/atomix/errors"
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
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:62
	_go_fuzz_dep_.CoverTab[114607]++
													return e.Message
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:63
	// _ = "end of CoverTab[114607]"
}

var _ error = &TypedError{}

// Status gets the gRPC status for the given error
func Status(err error) *status.Status {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:69
	_go_fuzz_dep_.CoverTab[114608]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:70
		_go_fuzz_dep_.CoverTab[114611]++
														return status.New(codes.OK, "")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:71
		// _ = "end of CoverTab[114611]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:72
		_go_fuzz_dep_.CoverTab[114612]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:72
		// _ = "end of CoverTab[114612]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:72
	// _ = "end of CoverTab[114608]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:72
	_go_fuzz_dep_.CoverTab[114609]++

													typed, ok := err.(*TypedError)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:75
		_go_fuzz_dep_.CoverTab[114613]++
														return status.New(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:76
		// _ = "end of CoverTab[114613]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:77
		_go_fuzz_dep_.CoverTab[114614]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:77
		// _ = "end of CoverTab[114614]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:77
	// _ = "end of CoverTab[114609]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:77
	_go_fuzz_dep_.CoverTab[114610]++

													switch typed.Type {
	case Unknown:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:80
		_go_fuzz_dep_.CoverTab[114615]++
														return status.New(codes.Unknown, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:81
		// _ = "end of CoverTab[114615]"
	case Canceled:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:82
		_go_fuzz_dep_.CoverTab[114616]++
														return status.New(codes.Canceled, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:83
		// _ = "end of CoverTab[114616]"
	case NotFound:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:84
		_go_fuzz_dep_.CoverTab[114617]++
														return status.New(codes.NotFound, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:85
		// _ = "end of CoverTab[114617]"
	case AlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:86
		_go_fuzz_dep_.CoverTab[114618]++
														return status.New(codes.AlreadyExists, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:87
		// _ = "end of CoverTab[114618]"
	case Unauthorized:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:88
		_go_fuzz_dep_.CoverTab[114619]++
														return status.New(codes.Unauthenticated, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:89
		// _ = "end of CoverTab[114619]"
	case Forbidden:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:90
		_go_fuzz_dep_.CoverTab[114620]++
														return status.New(codes.PermissionDenied, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:91
		// _ = "end of CoverTab[114620]"
	case Conflict:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:92
		_go_fuzz_dep_.CoverTab[114621]++
														return status.New(codes.FailedPrecondition, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:93
		// _ = "end of CoverTab[114621]"
	case Invalid:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:94
		_go_fuzz_dep_.CoverTab[114622]++
														return status.New(codes.InvalidArgument, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:95
		// _ = "end of CoverTab[114622]"
	case Unavailable:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:96
		_go_fuzz_dep_.CoverTab[114623]++
														return status.New(codes.Unavailable, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:97
		// _ = "end of CoverTab[114623]"
	case NotSupported:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:98
		_go_fuzz_dep_.CoverTab[114624]++
														return status.New(codes.Unimplemented, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:99
		// _ = "end of CoverTab[114624]"
	case Timeout:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:100
		_go_fuzz_dep_.CoverTab[114625]++
														return status.New(codes.DeadlineExceeded, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:101
		// _ = "end of CoverTab[114625]"
	case Internal:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:102
		_go_fuzz_dep_.CoverTab[114626]++
														return status.New(codes.Internal, typed.Message)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:103
		// _ = "end of CoverTab[114626]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:104
		_go_fuzz_dep_.CoverTab[114627]++
														return status.New(codes.Internal, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:105
		// _ = "end of CoverTab[114627]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:106
	// _ = "end of CoverTab[114610]"
}

// FromStatus creates a typed error from a gRPC status
func FromStatus(status *status.Status) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:110
	_go_fuzz_dep_.CoverTab[114628]++
													switch status.Code() {
	case codes.OK:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:112
		_go_fuzz_dep_.CoverTab[114629]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:113
		// _ = "end of CoverTab[114629]"
	case codes.Unknown:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:114
		_go_fuzz_dep_.CoverTab[114630]++
														return NewUnknown(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:115
		// _ = "end of CoverTab[114630]"
	case codes.Canceled:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:116
		_go_fuzz_dep_.CoverTab[114631]++
														return NewCanceled(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:117
		// _ = "end of CoverTab[114631]"
	case codes.NotFound:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:118
		_go_fuzz_dep_.CoverTab[114632]++
														return NewNotFound(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:119
		// _ = "end of CoverTab[114632]"
	case codes.AlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:120
		_go_fuzz_dep_.CoverTab[114633]++
														return NewAlreadyExists(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:121
		// _ = "end of CoverTab[114633]"
	case codes.Unauthenticated:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:122
		_go_fuzz_dep_.CoverTab[114634]++
														return NewUnauthorized(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:123
		// _ = "end of CoverTab[114634]"
	case codes.PermissionDenied:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:124
		_go_fuzz_dep_.CoverTab[114635]++
														return NewForbidden(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:125
		// _ = "end of CoverTab[114635]"
	case codes.FailedPrecondition:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:126
		_go_fuzz_dep_.CoverTab[114636]++
														return NewConflict(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:127
		// _ = "end of CoverTab[114636]"
	case codes.InvalidArgument:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:128
		_go_fuzz_dep_.CoverTab[114637]++
														return NewInvalid(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:129
		// _ = "end of CoverTab[114637]"
	case codes.Unavailable:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:130
		_go_fuzz_dep_.CoverTab[114638]++
														return NewUnavailable(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:131
		// _ = "end of CoverTab[114638]"
	case codes.Unimplemented:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:132
		_go_fuzz_dep_.CoverTab[114639]++
														return NewNotSupported(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:133
		// _ = "end of CoverTab[114639]"
	case codes.DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:134
		_go_fuzz_dep_.CoverTab[114640]++
														return NewTimeout(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:135
		// _ = "end of CoverTab[114640]"
	case codes.Internal:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:136
		_go_fuzz_dep_.CoverTab[114641]++
														return NewInternal(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:137
		// _ = "end of CoverTab[114641]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:138
		_go_fuzz_dep_.CoverTab[114642]++
														return NewUnknown(status.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:139
		// _ = "end of CoverTab[114642]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:140
	// _ = "end of CoverTab[114628]"
}

// FromGRPC creates a typed error from a gRPC error
func FromGRPC(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:144
	_go_fuzz_dep_.CoverTab[114643]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:145
		_go_fuzz_dep_.CoverTab[114646]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:146
		// _ = "end of CoverTab[114646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:147
		_go_fuzz_dep_.CoverTab[114647]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:147
		// _ = "end of CoverTab[114647]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:147
	// _ = "end of CoverTab[114643]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:147
	_go_fuzz_dep_.CoverTab[114644]++

													stat, ok := status.FromError(err)
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:150
		_go_fuzz_dep_.CoverTab[114648]++
														return New(Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:151
		// _ = "end of CoverTab[114648]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:152
		_go_fuzz_dep_.CoverTab[114649]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:152
		// _ = "end of CoverTab[114649]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:152
	// _ = "end of CoverTab[114644]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:152
	_go_fuzz_dep_.CoverTab[114645]++

													switch stat.Code() {
	case codes.OK:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:155
		_go_fuzz_dep_.CoverTab[114650]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:156
		// _ = "end of CoverTab[114650]"
	case codes.Unknown:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:157
		_go_fuzz_dep_.CoverTab[114651]++
														return New(Unknown, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:158
		// _ = "end of CoverTab[114651]"
	case codes.Canceled:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:159
		_go_fuzz_dep_.CoverTab[114652]++
														return New(Canceled, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:160
		// _ = "end of CoverTab[114652]"
	case codes.NotFound:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:161
		_go_fuzz_dep_.CoverTab[114653]++
														return New(NotFound, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:162
		// _ = "end of CoverTab[114653]"
	case codes.AlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:163
		_go_fuzz_dep_.CoverTab[114654]++
														return New(AlreadyExists, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:164
		// _ = "end of CoverTab[114654]"
	case codes.Unauthenticated:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:165
		_go_fuzz_dep_.CoverTab[114655]++
														return New(Unauthorized, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:166
		// _ = "end of CoverTab[114655]"
	case codes.PermissionDenied:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:167
		_go_fuzz_dep_.CoverTab[114656]++
														return New(Forbidden, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:168
		// _ = "end of CoverTab[114656]"
	case codes.FailedPrecondition:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:169
		_go_fuzz_dep_.CoverTab[114657]++
														return New(Conflict, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:170
		// _ = "end of CoverTab[114657]"
	case codes.InvalidArgument:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:171
		_go_fuzz_dep_.CoverTab[114658]++
														return New(Invalid, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:172
		// _ = "end of CoverTab[114658]"
	case codes.Unavailable:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:173
		_go_fuzz_dep_.CoverTab[114659]++
														return New(Unavailable, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:174
		// _ = "end of CoverTab[114659]"
	case codes.Unimplemented:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:175
		_go_fuzz_dep_.CoverTab[114660]++
														return New(NotSupported, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:176
		// _ = "end of CoverTab[114660]"
	case codes.DeadlineExceeded:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:177
		_go_fuzz_dep_.CoverTab[114661]++
														return New(Timeout, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:178
		// _ = "end of CoverTab[114661]"
	case codes.Internal:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:179
		_go_fuzz_dep_.CoverTab[114662]++
														return New(Internal, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:180
		// _ = "end of CoverTab[114662]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:181
		_go_fuzz_dep_.CoverTab[114663]++
														return New(Unknown, stat.Message())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:182
		// _ = "end of CoverTab[114663]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:183
	// _ = "end of CoverTab[114645]"
}

// FromAtomix creates a typed error from an Atomix error
func FromAtomix(err error) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:187
	_go_fuzz_dep_.CoverTab[114664]++
													if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:188
		_go_fuzz_dep_.CoverTab[114667]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:189
		// _ = "end of CoverTab[114667]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:190
		_go_fuzz_dep_.CoverTab[114668]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:190
		// _ = "end of CoverTab[114668]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:190
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:190
	// _ = "end of CoverTab[114664]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:190
	_go_fuzz_dep_.CoverTab[114665]++

													if typed, ok := err.(*atomixerrors.TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:192
		_go_fuzz_dep_.CoverTab[114669]++
														switch typed.Type {
		case atomixerrors.Unknown:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:194
			_go_fuzz_dep_.CoverTab[114670]++
															return New(Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:195
			// _ = "end of CoverTab[114670]"
		case atomixerrors.Canceled:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:196
			_go_fuzz_dep_.CoverTab[114671]++
															return New(Canceled, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:197
			// _ = "end of CoverTab[114671]"
		case atomixerrors.NotFound:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:198
			_go_fuzz_dep_.CoverTab[114672]++
															return New(NotFound, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:199
			// _ = "end of CoverTab[114672]"
		case atomixerrors.AlreadyExists:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:200
			_go_fuzz_dep_.CoverTab[114673]++
															return New(AlreadyExists, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:201
			// _ = "end of CoverTab[114673]"
		case atomixerrors.Unauthorized:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:202
			_go_fuzz_dep_.CoverTab[114674]++
															return New(Unauthorized, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:203
			// _ = "end of CoverTab[114674]"
		case atomixerrors.Forbidden:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:204
			_go_fuzz_dep_.CoverTab[114675]++
															return New(Forbidden, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:205
			// _ = "end of CoverTab[114675]"
		case atomixerrors.Conflict:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:206
			_go_fuzz_dep_.CoverTab[114676]++
															return New(Conflict, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:207
			// _ = "end of CoverTab[114676]"
		case atomixerrors.Invalid:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:208
			_go_fuzz_dep_.CoverTab[114677]++
															return New(Invalid, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:209
			// _ = "end of CoverTab[114677]"
		case atomixerrors.Unavailable:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:210
			_go_fuzz_dep_.CoverTab[114678]++
															return New(Unavailable, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:211
			// _ = "end of CoverTab[114678]"
		case atomixerrors.NotSupported:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:212
			_go_fuzz_dep_.CoverTab[114679]++
															return New(NotSupported, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:213
			// _ = "end of CoverTab[114679]"
		case atomixerrors.Timeout:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:214
			_go_fuzz_dep_.CoverTab[114680]++
															return New(Timeout, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:215
			// _ = "end of CoverTab[114680]"
		case atomixerrors.Internal:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:216
			_go_fuzz_dep_.CoverTab[114681]++
															return New(Internal, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:217
			// _ = "end of CoverTab[114681]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:218
			_go_fuzz_dep_.CoverTab[114682]++
															return New(Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:219
			// _ = "end of CoverTab[114682]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:220
		// _ = "end of CoverTab[114669]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:221
		_go_fuzz_dep_.CoverTab[114683]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:221
		// _ = "end of CoverTab[114683]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:221
	// _ = "end of CoverTab[114665]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:221
	_go_fuzz_dep_.CoverTab[114666]++
													return New(Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:222
	// _ = "end of CoverTab[114666]"
}

// New creates a new typed error
func New(t Type, msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:226
	_go_fuzz_dep_.CoverTab[114684]++
													if len(args) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:227
		_go_fuzz_dep_.CoverTab[114686]++
														msg = fmt.Sprintf(msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:228
		// _ = "end of CoverTab[114686]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:229
		_go_fuzz_dep_.CoverTab[114687]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:229
		// _ = "end of CoverTab[114687]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:229
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:229
	// _ = "end of CoverTab[114684]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:229
	_go_fuzz_dep_.CoverTab[114685]++
													return &TypedError{
		Type:		t,
		Message:	msg,
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:233
	// _ = "end of CoverTab[114685]"
}

// NewUnknown returns a new Unknown error
func NewUnknown(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:237
	_go_fuzz_dep_.CoverTab[114688]++
													return New(Unknown, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:238
	// _ = "end of CoverTab[114688]"
}

// NewCanceled returns a new Canceled error
func NewCanceled(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:242
	_go_fuzz_dep_.CoverTab[114689]++
													return New(Canceled, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:243
	// _ = "end of CoverTab[114689]"
}

// NewNotFound returns a new NotFound error
func NewNotFound(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:247
	_go_fuzz_dep_.CoverTab[114690]++
													return New(NotFound, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:248
	// _ = "end of CoverTab[114690]"
}

// NewAlreadyExists returns a new AlreadyExists error
func NewAlreadyExists(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:252
	_go_fuzz_dep_.CoverTab[114691]++
													return New(AlreadyExists, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:253
	// _ = "end of CoverTab[114691]"
}

// NewUnauthorized returns a new Unauthorized error
func NewUnauthorized(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:257
	_go_fuzz_dep_.CoverTab[114692]++
													return New(Unauthorized, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:258
	// _ = "end of CoverTab[114692]"
}

// NewForbidden returns a new Forbidden error
func NewForbidden(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:262
	_go_fuzz_dep_.CoverTab[114693]++
													return New(Forbidden, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:263
	// _ = "end of CoverTab[114693]"
}

// NewConflict returns a new Conflict error
func NewConflict(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:267
	_go_fuzz_dep_.CoverTab[114694]++
													return New(Conflict, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:268
	// _ = "end of CoverTab[114694]"
}

// NewInvalid returns a new Invalid error
func NewInvalid(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:272
	_go_fuzz_dep_.CoverTab[114695]++
													return New(Invalid, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:273
	// _ = "end of CoverTab[114695]"
}

// NewUnavailable returns a new Unavailable error
func NewUnavailable(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:277
	_go_fuzz_dep_.CoverTab[114696]++
													return New(Unavailable, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:278
	// _ = "end of CoverTab[114696]"
}

// NewNotSupported returns a new NotSupported error
func NewNotSupported(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:282
	_go_fuzz_dep_.CoverTab[114697]++
													return New(NotSupported, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:283
	// _ = "end of CoverTab[114697]"
}

// NewTimeout returns a new Timeout error
func NewTimeout(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:287
	_go_fuzz_dep_.CoverTab[114698]++
													return New(Timeout, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:288
	// _ = "end of CoverTab[114698]"
}

// NewInternal returns a new Internal error
func NewInternal(msg string, args ...interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:292
	_go_fuzz_dep_.CoverTab[114699]++
													return New(Internal, msg, args...)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:293
	// _ = "end of CoverTab[114699]"
}

// TypeOf returns the type of the given error
func TypeOf(err error) Type {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:297
	_go_fuzz_dep_.CoverTab[114700]++
													if typed, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:298
		_go_fuzz_dep_.CoverTab[114702]++
														return typed.Type
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:299
		// _ = "end of CoverTab[114702]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:300
		_go_fuzz_dep_.CoverTab[114703]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:300
		// _ = "end of CoverTab[114703]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:300
	// _ = "end of CoverTab[114700]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:300
	_go_fuzz_dep_.CoverTab[114701]++
													return Unknown
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:301
	// _ = "end of CoverTab[114701]"
}

// IsType checks whether the given error is of the given type
func IsType(err error, t Type) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:305
	_go_fuzz_dep_.CoverTab[114704]++
													if typed, ok := err.(*TypedError); ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:306
		_go_fuzz_dep_.CoverTab[114706]++
														return typed.Type == t
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:307
		// _ = "end of CoverTab[114706]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:308
		_go_fuzz_dep_.CoverTab[114707]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:308
		// _ = "end of CoverTab[114707]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:308
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:308
	// _ = "end of CoverTab[114704]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:308
	_go_fuzz_dep_.CoverTab[114705]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:309
	// _ = "end of CoverTab[114705]"
}

// IsUnknown checks whether the given error is an Unknown error
func IsUnknown(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:313
	_go_fuzz_dep_.CoverTab[114708]++
													return IsType(err, Unknown)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:314
	// _ = "end of CoverTab[114708]"
}

// IsCanceled checks whether the given error is an Canceled error
func IsCanceled(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:318
	_go_fuzz_dep_.CoverTab[114709]++
													return IsType(err, Canceled)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:319
	// _ = "end of CoverTab[114709]"
}

// IsNotFound checks whether the given error is a NotFound error
func IsNotFound(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:323
	_go_fuzz_dep_.CoverTab[114710]++
													return IsType(err, NotFound)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:324
	// _ = "end of CoverTab[114710]"
}

// IsAlreadyExists checks whether the given error is a AlreadyExists error
func IsAlreadyExists(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:328
	_go_fuzz_dep_.CoverTab[114711]++
													return IsType(err, AlreadyExists)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:329
	// _ = "end of CoverTab[114711]"
}

// IsUnauthorized checks whether the given error is a Unauthorized error
func IsUnauthorized(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:333
	_go_fuzz_dep_.CoverTab[114712]++
													return IsType(err, Unauthorized)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:334
	// _ = "end of CoverTab[114712]"
}

// IsForbidden checks whether the given error is a Forbidden error
func IsForbidden(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:338
	_go_fuzz_dep_.CoverTab[114713]++
													return IsType(err, Forbidden)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:339
	// _ = "end of CoverTab[114713]"
}

// IsConflict checks whether the given error is a Conflict error
func IsConflict(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:343
	_go_fuzz_dep_.CoverTab[114714]++
													return IsType(err, Conflict)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:344
	// _ = "end of CoverTab[114714]"
}

// IsInvalid checks whether the given error is an Invalid error
func IsInvalid(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:348
	_go_fuzz_dep_.CoverTab[114715]++
													return IsType(err, Invalid)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:349
	// _ = "end of CoverTab[114715]"
}

// IsUnavailable checks whether the given error is an Unavailable error
func IsUnavailable(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:353
	_go_fuzz_dep_.CoverTab[114716]++
													return IsType(err, Unavailable)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:354
	// _ = "end of CoverTab[114716]"
}

// IsNotSupported checks whether the given error is a NotSupported error
func IsNotSupported(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:358
	_go_fuzz_dep_.CoverTab[114717]++
													return IsType(err, NotSupported)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:359
	// _ = "end of CoverTab[114717]"
}

// IsTimeout checks whether the given error is a Timeout error
func IsTimeout(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:363
	_go_fuzz_dep_.CoverTab[114718]++
													return IsType(err, Timeout)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:364
	// _ = "end of CoverTab[114718]"
}

// IsInternal checks whether the given error is an Internal error
func IsInternal(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:368
	_go_fuzz_dep_.CoverTab[114719]++
													return IsType(err, Internal)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:369
	// _ = "end of CoverTab[114719]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:370
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/errors/errors.go:370
var _ = _go_fuzz_dep_.CoverTab
