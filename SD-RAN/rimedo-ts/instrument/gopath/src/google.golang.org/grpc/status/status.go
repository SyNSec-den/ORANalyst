//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// Package status implements errors returned by gRPC.  These errors are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// serialized and transmitted on the wire between server and client, and allow
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// for additional data to be transmitted via the Details field in the status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// proto.  gRPC service handlers should return an error created by this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// package, and gRPC clients should expect a corresponding error to be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// returned from the RPC call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// This package upholds the invariants that a non-nil error may not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:19
// contain an OK code, and an OK code must result in a nil error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
package status

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:28
)

import (
	"context"
	"errors"
	"fmt"

	spb "google.golang.org/genproto/googleapis/rpc/status"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

// Status references google.golang.org/grpc/internal/status. It represents an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:41
// RPC status code, message, and details.  It is immutable and should be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:41
// created with New, Newf, or FromProto.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:41
// https://godoc.org/google.golang.org/grpc/internal/status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:45
type Status = status.Status

// New returns a Status representing c and msg.
func New(c codes.Code, msg string) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:48
	_go_fuzz_dep_.CoverTab[68548]++
											return status.New(c, msg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:49
	// _ = "end of CoverTab[68548]"
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c codes.Code, format string, a ...interface{}) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:53
	_go_fuzz_dep_.CoverTab[68549]++
											return New(c, fmt.Sprintf(format, a...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:54
	// _ = "end of CoverTab[68549]"
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c codes.Code, msg string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:58
	_go_fuzz_dep_.CoverTab[68550]++
											return New(c, msg).Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:59
	// _ = "end of CoverTab[68550]"
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c codes.Code, format string, a ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:63
	_go_fuzz_dep_.CoverTab[68551]++
											return Error(c, fmt.Sprintf(format, a...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:64
	// _ = "end of CoverTab[68551]"
}

// ErrorProto returns an error representing s.  If s.Code is OK, returns nil.
func ErrorProto(s *spb.Status) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:68
	_go_fuzz_dep_.CoverTab[68552]++
											return FromProto(s).Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:69
	// _ = "end of CoverTab[68552]"
}

// FromProto returns a Status representing s.
func FromProto(s *spb.Status) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:73
	_go_fuzz_dep_.CoverTab[68553]++
											return status.FromProto(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:74
	// _ = "end of CoverTab[68553]"
}

// FromError returns a Status representation of err.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//   - If err was produced by this package or implements the method `GRPCStatus()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//     *Status`, the appropriate Status is returned.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//   - If err is nil, a Status is returned with codes.OK and no message.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//   - Otherwise, err is an error not compatible with this package.  In this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//     case, a Status is returned with codes.Unknown and err's Error() message,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:77
//     and ok is false.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:87
func FromError(err error) (s *Status, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:87
	_go_fuzz_dep_.CoverTab[68554]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:88
		_go_fuzz_dep_.CoverTab[68557]++
												return nil, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:89
		// _ = "end of CoverTab[68557]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:90
		_go_fuzz_dep_.CoverTab[68558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:90
		// _ = "end of CoverTab[68558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:90
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:90
	// _ = "end of CoverTab[68554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:90
	_go_fuzz_dep_.CoverTab[68555]++
											if se, ok := err.(interface {
		GRPCStatus() *Status
	}); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:93
		_go_fuzz_dep_.CoverTab[68559]++
												return se.GRPCStatus(), true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:94
		// _ = "end of CoverTab[68559]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:95
		_go_fuzz_dep_.CoverTab[68560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:95
		// _ = "end of CoverTab[68560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:95
	// _ = "end of CoverTab[68555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:95
	_go_fuzz_dep_.CoverTab[68556]++
											return New(codes.Unknown, err.Error()), false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:96
	// _ = "end of CoverTab[68556]"
}

// Convert is a convenience function which removes the need to handle the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:99
// boolean return value from FromError.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:101
func Convert(err error) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:101
	_go_fuzz_dep_.CoverTab[68561]++
											s, _ := FromError(err)
											return s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:103
	// _ = "end of CoverTab[68561]"
}

// Code returns the Code of the error if it is a Status error, codes.OK if err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:106
// is nil, or codes.Unknown otherwise.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:108
func Code(err error) codes.Code {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:108
	_go_fuzz_dep_.CoverTab[68562]++

											if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:110
		_go_fuzz_dep_.CoverTab[68565]++
												return codes.OK
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:111
		// _ = "end of CoverTab[68565]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:112
		_go_fuzz_dep_.CoverTab[68566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:112
		// _ = "end of CoverTab[68566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:112
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:112
	// _ = "end of CoverTab[68562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:112
	_go_fuzz_dep_.CoverTab[68563]++
											if se, ok := err.(interface {
		GRPCStatus() *Status
	}); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:115
		_go_fuzz_dep_.CoverTab[68567]++
												return se.GRPCStatus().Code()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:116
		// _ = "end of CoverTab[68567]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:117
		_go_fuzz_dep_.CoverTab[68568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:117
		// _ = "end of CoverTab[68568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:117
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:117
	// _ = "end of CoverTab[68563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:117
	_go_fuzz_dep_.CoverTab[68564]++
											return codes.Unknown
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:118
	// _ = "end of CoverTab[68564]"
}

// FromContextError converts a context error or wrapped context error into a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:121
// Status.  It returns a Status with codes.OK if err is nil, or a Status with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:121
// codes.Unknown if err is non-nil and not a context error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:124
func FromContextError(err error) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:124
	_go_fuzz_dep_.CoverTab[68569]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:125
		_go_fuzz_dep_.CoverTab[68573]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:126
		// _ = "end of CoverTab[68573]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:127
		_go_fuzz_dep_.CoverTab[68574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:127
		// _ = "end of CoverTab[68574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:127
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:127
	// _ = "end of CoverTab[68569]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:127
	_go_fuzz_dep_.CoverTab[68570]++
											if errors.Is(err, context.DeadlineExceeded) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:128
		_go_fuzz_dep_.CoverTab[68575]++
												return New(codes.DeadlineExceeded, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:129
		// _ = "end of CoverTab[68575]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:130
		_go_fuzz_dep_.CoverTab[68576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:130
		// _ = "end of CoverTab[68576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:130
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:130
	// _ = "end of CoverTab[68570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:130
	_go_fuzz_dep_.CoverTab[68571]++
											if errors.Is(err, context.Canceled) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:131
		_go_fuzz_dep_.CoverTab[68577]++
												return New(codes.Canceled, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:132
		// _ = "end of CoverTab[68577]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:133
		_go_fuzz_dep_.CoverTab[68578]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:133
		// _ = "end of CoverTab[68578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:133
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:133
	// _ = "end of CoverTab[68571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:133
	_go_fuzz_dep_.CoverTab[68572]++
											return New(codes.Unknown, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:134
	// _ = "end of CoverTab[68572]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:135
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/status/status.go:135
var _ = _go_fuzz_dep_.CoverTab
