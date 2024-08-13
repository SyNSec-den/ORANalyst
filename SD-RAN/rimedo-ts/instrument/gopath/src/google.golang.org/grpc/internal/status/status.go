//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// Package status implements errors returned by gRPC.  These errors are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// serialized and transmitted on the wire between server and client, and allow
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// for additional data to be transmitted via the Details field in the status
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// proto.  gRPC service handlers should return an error created by this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// package, and gRPC clients should expect a corresponding error to be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// returned from the RPC call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// This package upholds the invariants that a non-nil error may not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:19
// contain an OK code, and an OK code must result in a nil error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
package status

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:28
)

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

// Status represents an RPC status code, message, and details.  It is immutable
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:40
// and should be created with New, Newf, or FromProto.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:42
type Status struct {
	s *spb.Status
}

// New returns a Status representing c and msg.
func New(c codes.Code, msg string) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:47
	_go_fuzz_dep_.CoverTab[68495]++
												return &Status{s: &spb.Status{Code: int32(c), Message: msg}}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:48
	// _ = "end of CoverTab[68495]"
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c codes.Code, format string, a ...interface{}) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:52
	_go_fuzz_dep_.CoverTab[68496]++
												return New(c, fmt.Sprintf(format, a...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:53
	// _ = "end of CoverTab[68496]"
}

// FromProto returns a Status representing s.
func FromProto(s *spb.Status) *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:57
	_go_fuzz_dep_.CoverTab[68497]++
												return &Status{s: proto.Clone(s).(*spb.Status)}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:58
	// _ = "end of CoverTab[68497]"
}

// Err returns an error representing c and msg.  If c is OK, returns nil.
func Err(c codes.Code, msg string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:62
	_go_fuzz_dep_.CoverTab[68498]++
												return New(c, msg).Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:63
	// _ = "end of CoverTab[68498]"
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c codes.Code, format string, a ...interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:67
	_go_fuzz_dep_.CoverTab[68499]++
												return Err(c, fmt.Sprintf(format, a...))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:68
	// _ = "end of CoverTab[68499]"
}

// Code returns the status code contained in s.
func (s *Status) Code() codes.Code {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:72
	_go_fuzz_dep_.CoverTab[68500]++
												if s == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:73
		_go_fuzz_dep_.CoverTab[68502]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:73
		return s.s == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:73
		// _ = "end of CoverTab[68502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:73
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:73
		_go_fuzz_dep_.CoverTab[68503]++
													return codes.OK
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:74
		// _ = "end of CoverTab[68503]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:75
		_go_fuzz_dep_.CoverTab[68504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:75
		// _ = "end of CoverTab[68504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:75
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:75
	// _ = "end of CoverTab[68500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:75
	_go_fuzz_dep_.CoverTab[68501]++
												return codes.Code(s.s.Code)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:76
	// _ = "end of CoverTab[68501]"
}

// Message returns the message contained in s.
func (s *Status) Message() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:80
	_go_fuzz_dep_.CoverTab[68505]++
												if s == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:81
		_go_fuzz_dep_.CoverTab[68507]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:81
		return s.s == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:81
		// _ = "end of CoverTab[68507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:81
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:81
		_go_fuzz_dep_.CoverTab[68508]++
													return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:82
		// _ = "end of CoverTab[68508]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:83
		_go_fuzz_dep_.CoverTab[68509]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:83
		// _ = "end of CoverTab[68509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:83
	// _ = "end of CoverTab[68505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:83
	_go_fuzz_dep_.CoverTab[68506]++
												return s.s.Message
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:84
	// _ = "end of CoverTab[68506]"
}

// Proto returns s's status as an spb.Status proto message.
func (s *Status) Proto() *spb.Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:88
	_go_fuzz_dep_.CoverTab[68510]++
												if s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:89
		_go_fuzz_dep_.CoverTab[68512]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:90
		// _ = "end of CoverTab[68512]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:91
		_go_fuzz_dep_.CoverTab[68513]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:91
		// _ = "end of CoverTab[68513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:91
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:91
	// _ = "end of CoverTab[68510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:91
	_go_fuzz_dep_.CoverTab[68511]++
												return proto.Clone(s.s).(*spb.Status)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:92
	// _ = "end of CoverTab[68511]"
}

// Err returns an immutable error representing s; returns nil if s.Code() is OK.
func (s *Status) Err() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:96
	_go_fuzz_dep_.CoverTab[68514]++
												if s.Code() == codes.OK {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:97
		_go_fuzz_dep_.CoverTab[68516]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:98
		// _ = "end of CoverTab[68516]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:99
		_go_fuzz_dep_.CoverTab[68517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:99
		// _ = "end of CoverTab[68517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:99
	// _ = "end of CoverTab[68514]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:99
	_go_fuzz_dep_.CoverTab[68515]++
												return &Error{s: s}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:100
	// _ = "end of CoverTab[68515]"
}

// WithDetails returns a new status with the provided details messages appended to the status.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:103
// If any errors are encountered, it returns nil and the first error encountered.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:105
func (s *Status) WithDetails(details ...proto.Message) (*Status, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:105
	_go_fuzz_dep_.CoverTab[68518]++
												if s.Code() == codes.OK {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:106
		_go_fuzz_dep_.CoverTab[68521]++
													return nil, errors.New("no error details for status with code OK")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:107
		// _ = "end of CoverTab[68521]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:108
		_go_fuzz_dep_.CoverTab[68522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:108
		// _ = "end of CoverTab[68522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:108
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:108
	// _ = "end of CoverTab[68518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:108
	_go_fuzz_dep_.CoverTab[68519]++

												p := s.Proto()
												for _, detail := range details {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:111
		_go_fuzz_dep_.CoverTab[68523]++
													any, err := ptypes.MarshalAny(detail)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:113
			_go_fuzz_dep_.CoverTab[68525]++
														return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:114
			// _ = "end of CoverTab[68525]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:115
			_go_fuzz_dep_.CoverTab[68526]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:115
			// _ = "end of CoverTab[68526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:115
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:115
		// _ = "end of CoverTab[68523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:115
		_go_fuzz_dep_.CoverTab[68524]++
													p.Details = append(p.Details, any)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:116
		// _ = "end of CoverTab[68524]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:117
	// _ = "end of CoverTab[68519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:117
	_go_fuzz_dep_.CoverTab[68520]++
												return &Status{s: p}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:118
	// _ = "end of CoverTab[68520]"
}

// Details returns a slice of details messages attached to the status.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:121
// If a detail cannot be decoded, the error is returned in place of the detail.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:123
func (s *Status) Details() []interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:123
	_go_fuzz_dep_.CoverTab[68527]++
												if s == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:124
		_go_fuzz_dep_.CoverTab[68530]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:124
		return s.s == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:124
		// _ = "end of CoverTab[68530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:124
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:124
		_go_fuzz_dep_.CoverTab[68531]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:125
		// _ = "end of CoverTab[68531]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:126
		_go_fuzz_dep_.CoverTab[68532]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:126
		// _ = "end of CoverTab[68532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:126
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:126
	// _ = "end of CoverTab[68527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:126
	_go_fuzz_dep_.CoverTab[68528]++
												details := make([]interface{}, 0, len(s.s.Details))
												for _, any := range s.s.Details {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:128
		_go_fuzz_dep_.CoverTab[68533]++
													detail := &ptypes.DynamicAny{}
													if err := ptypes.UnmarshalAny(any, detail); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:130
			_go_fuzz_dep_.CoverTab[68535]++
														details = append(details, err)
														continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:132
			// _ = "end of CoverTab[68535]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:133
			_go_fuzz_dep_.CoverTab[68536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:133
			// _ = "end of CoverTab[68536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:133
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:133
		// _ = "end of CoverTab[68533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:133
		_go_fuzz_dep_.CoverTab[68534]++
													details = append(details, detail.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:134
		// _ = "end of CoverTab[68534]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:135
	// _ = "end of CoverTab[68528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:135
	_go_fuzz_dep_.CoverTab[68529]++
												return details
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:136
	// _ = "end of CoverTab[68529]"
}

func (s *Status) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:139
	_go_fuzz_dep_.CoverTab[68537]++
												return fmt.Sprintf("rpc error: code = %s desc = %s", s.Code(), s.Message())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:140
	// _ = "end of CoverTab[68537]"
}

// Error wraps a pointer of a status proto. It implements error and Status,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:143
// and a nil *Error should never be returned by this package.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:145
type Error struct {
	s *Status
}

func (e *Error) Error() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:149
	_go_fuzz_dep_.CoverTab[68538]++
												return e.s.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:150
	// _ = "end of CoverTab[68538]"
}

// GRPCStatus returns the Status represented by se.
func (e *Error) GRPCStatus() *Status {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:154
	_go_fuzz_dep_.CoverTab[68539]++
												return e.s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:155
	// _ = "end of CoverTab[68539]"
}

// Is implements future error.Is functionality.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:158
// A Error is equivalent if the code and message are identical.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:160
func (e *Error) Is(target error) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:160
	_go_fuzz_dep_.CoverTab[68540]++
												tse, ok := target.(*Error)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:162
		_go_fuzz_dep_.CoverTab[68542]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:163
		// _ = "end of CoverTab[68542]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:164
		_go_fuzz_dep_.CoverTab[68543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:164
		// _ = "end of CoverTab[68543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:164
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:164
	// _ = "end of CoverTab[68540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:164
	_go_fuzz_dep_.CoverTab[68541]++
												return proto.Equal(e.s.s, tse.s.s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:165
	// _ = "end of CoverTab[68541]"
}

// IsRestrictedControlPlaneCode returns whether the status includes a code
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:168
// restricted for control plane usage as defined by gRFC A54.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:170
func IsRestrictedControlPlaneCode(s *Status) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:170
	_go_fuzz_dep_.CoverTab[68544]++
												switch s.Code() {
	case codes.InvalidArgument, codes.NotFound, codes.AlreadyExists, codes.FailedPrecondition, codes.Aborted, codes.OutOfRange, codes.DataLoss:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:172
		_go_fuzz_dep_.CoverTab[68546]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:173
		// _ = "end of CoverTab[68546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:173
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:173
		_go_fuzz_dep_.CoverTab[68547]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:173
		// _ = "end of CoverTab[68547]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:174
	// _ = "end of CoverTab[68544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:174
	_go_fuzz_dep_.CoverTab[68545]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:175
	// _ = "end of CoverTab[68545]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:176
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/status/status.go:176
var _ = _go_fuzz_dep_.CoverTab
