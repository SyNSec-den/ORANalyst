// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:5
)

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/order"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// MarshalOptions configures the marshaler.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:16
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:16
// Example usage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:16
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:16
//	b, err := MarshalOptions{Deterministic: true}.Marshal(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:21
type MarshalOptions struct {
	pragma.NoUnkeyedLiterals

	// AllowPartial allows messages that have missing required fields to marshal
	// without returning an error. If AllowPartial is false (the default),
	// Marshal will return an error if there are any missing required fields.
	AllowPartial	bool

	// Deterministic controls whether the same message will always be
	// serialized to the same bytes within the same binary.
	//
	// Setting this option guarantees that repeated serialization of
	// the same message will return the same bytes, and that different
	// processes of the same binary (which may be executing on different
	// machines) will serialize equal messages to the same bytes.
	// It has no effect on the resulting size of the encoded message compared
	// to a non-deterministic marshal.
	//
	// Note that the deterministic serialization is NOT canonical across
	// languages. It is not guaranteed to remain stable over time. It is
	// unstable across different builds with schema changes due to unknown
	// fields. Users who need canonical serialization (e.g., persistent
	// storage in a canonical form, fingerprinting, etc.) must define
	// their own canonicalization specification and implement their own
	// serializer rather than relying on this API.
	//
	// If deterministic serialization is requested, map entries will be
	// sorted by keys in lexographical order. This is an implementation
	// detail and subject to change.
	Deterministic	bool

	// UseCachedSize indicates that the result of a previous Size call
	// may be reused.
	//
	// Setting this option asserts that:
	//
	// 1. Size has previously been called on this message with identical
	// options (except for UseCachedSize itself).
	//
	// 2. The message and all its submessages have not changed in any
	// way since the Size call.
	//
	// If either of these invariants is violated,
	// the results are undefined and may include panics or corrupted output.
	//
	// Implementations MAY take this option into account to provide
	// better performance, but there is no guarantee that they will do so.
	// There is absolutely no guarantee that Size followed by Marshal with
	// UseCachedSize set will perform equivalently to Marshal alone.
	UseCachedSize	bool
}

// Marshal returns the wire-format encoding of m.
func Marshal(m Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:74
	_go_fuzz_dep_.CoverTab[51191]++

												if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:76
		_go_fuzz_dep_.CoverTab[51194]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:77
		// _ = "end of CoverTab[51194]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:78
		_go_fuzz_dep_.CoverTab[51195]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:78
		// _ = "end of CoverTab[51195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:78
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:78
	// _ = "end of CoverTab[51191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:78
	_go_fuzz_dep_.CoverTab[51192]++

												out, err := MarshalOptions{}.marshal(nil, m.ProtoReflect())
												if len(out.Buf) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:81
		_go_fuzz_dep_.CoverTab[51196]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:81
		return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:81
		// _ = "end of CoverTab[51196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:81
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:81
		_go_fuzz_dep_.CoverTab[51197]++
													out.Buf = emptyBytesForMessage(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:82
		// _ = "end of CoverTab[51197]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:83
		_go_fuzz_dep_.CoverTab[51198]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:83
		// _ = "end of CoverTab[51198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:83
	// _ = "end of CoverTab[51192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:83
	_go_fuzz_dep_.CoverTab[51193]++
												return out.Buf, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:84
	// _ = "end of CoverTab[51193]"
}

// Marshal returns the wire-format encoding of m.
func (o MarshalOptions) Marshal(m Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:88
	_go_fuzz_dep_.CoverTab[51199]++

												if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:90
		_go_fuzz_dep_.CoverTab[51202]++
													return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:91
		// _ = "end of CoverTab[51202]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:92
		_go_fuzz_dep_.CoverTab[51203]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:92
		// _ = "end of CoverTab[51203]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:92
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:92
	// _ = "end of CoverTab[51199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:92
	_go_fuzz_dep_.CoverTab[51200]++

												out, err := o.marshal(nil, m.ProtoReflect())
												if len(out.Buf) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:95
		_go_fuzz_dep_.CoverTab[51204]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:95
		return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:95
		// _ = "end of CoverTab[51204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:95
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:95
		_go_fuzz_dep_.CoverTab[51205]++
													out.Buf = emptyBytesForMessage(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:96
		// _ = "end of CoverTab[51205]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:97
		_go_fuzz_dep_.CoverTab[51206]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:97
		// _ = "end of CoverTab[51206]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:97
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:97
	// _ = "end of CoverTab[51200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:97
	_go_fuzz_dep_.CoverTab[51201]++
												return out.Buf, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:98
	// _ = "end of CoverTab[51201]"
}

// emptyBytesForMessage returns a nil buffer if and only if m is invalid,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
// otherwise it returns a non-nil empty buffer.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
// This is to assist the edge-case where user-code does the following:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
//	m1.OptionalBytes, _ = proto.Marshal(m2)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
// where they expect the proto2 "optional_bytes" field to be populated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:101
// if any only if m2 is a valid message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:110
func emptyBytesForMessage(m Message) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:110
	_go_fuzz_dep_.CoverTab[51207]++
												if m == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:111
		_go_fuzz_dep_.CoverTab[51209]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:111
		return !m.ProtoReflect().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:111
		// _ = "end of CoverTab[51209]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:111
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:111
		_go_fuzz_dep_.CoverTab[51210]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:112
		// _ = "end of CoverTab[51210]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:113
		_go_fuzz_dep_.CoverTab[51211]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:113
		// _ = "end of CoverTab[51211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:113
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:113
	// _ = "end of CoverTab[51207]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:113
	_go_fuzz_dep_.CoverTab[51208]++
												return emptyBuf[:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:114
	// _ = "end of CoverTab[51208]"
}

// MarshalAppend appends the wire-format encoding of m to b,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:117
// returning the result.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:119
func (o MarshalOptions) MarshalAppend(b []byte, m Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:119
	_go_fuzz_dep_.CoverTab[51212]++

												if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:121
		_go_fuzz_dep_.CoverTab[51214]++
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:122
		// _ = "end of CoverTab[51214]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:123
		_go_fuzz_dep_.CoverTab[51215]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:123
		// _ = "end of CoverTab[51215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:123
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:123
	// _ = "end of CoverTab[51212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:123
	_go_fuzz_dep_.CoverTab[51213]++

												out, err := o.marshal(b, m.ProtoReflect())
												return out.Buf, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:126
	// _ = "end of CoverTab[51213]"
}

// MarshalState returns the wire-format encoding of a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:129
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:129
// This method permits fine-grained control over the marshaler.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:129
// Most users should use Marshal instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:133
func (o MarshalOptions) MarshalState(in protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:133
	_go_fuzz_dep_.CoverTab[51216]++
												return o.marshal(in.Buf, in.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:134
	// _ = "end of CoverTab[51216]"
}

// marshal is a centralized function that all marshal operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:137
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:137
// introducing other code paths for marshal that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:140
func (o MarshalOptions) marshal(b []byte, m protoreflect.Message) (out protoiface.MarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:140
	_go_fuzz_dep_.CoverTab[51217]++
												allowPartial := o.AllowPartial
												o.AllowPartial = true
												if methods := protoMethods(m); methods != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:143
		_go_fuzz_dep_.CoverTab[51221]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:143
		return methods.Marshal != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:143
		// _ = "end of CoverTab[51221]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:143
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:143
		_go_fuzz_dep_.CoverTab[51222]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:143
		return !(o.Deterministic && func() bool {
														_go_fuzz_dep_.CoverTab[51223]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:144
			return methods.Flags&protoiface.SupportMarshalDeterministic == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:144
			// _ = "end of CoverTab[51223]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:144
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:144
		// _ = "end of CoverTab[51222]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:144
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:144
		_go_fuzz_dep_.CoverTab[51224]++
													in := protoiface.MarshalInput{
			Message:	m,
			Buf:		b,
		}
		if o.Deterministic {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:149
			_go_fuzz_dep_.CoverTab[51228]++
														in.Flags |= protoiface.MarshalDeterministic
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:150
			// _ = "end of CoverTab[51228]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:151
			_go_fuzz_dep_.CoverTab[51229]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:151
			// _ = "end of CoverTab[51229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:151
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:151
		// _ = "end of CoverTab[51224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:151
		_go_fuzz_dep_.CoverTab[51225]++
													if o.UseCachedSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:152
			_go_fuzz_dep_.CoverTab[51230]++
														in.Flags |= protoiface.MarshalUseCachedSize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:153
			// _ = "end of CoverTab[51230]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:154
			_go_fuzz_dep_.CoverTab[51231]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:154
			// _ = "end of CoverTab[51231]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:154
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:154
		// _ = "end of CoverTab[51225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:154
		_go_fuzz_dep_.CoverTab[51226]++
													if methods.Size != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:155
			_go_fuzz_dep_.CoverTab[51232]++
														sout := methods.Size(protoiface.SizeInput{
				Message:	m,
				Flags:		in.Flags,
			})
			if cap(b) < len(b)+sout.Size {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:160
				_go_fuzz_dep_.CoverTab[51234]++
															in.Buf = make([]byte, len(b), growcap(cap(b), len(b)+sout.Size))
															copy(in.Buf, b)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:162
				// _ = "end of CoverTab[51234]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:163
				_go_fuzz_dep_.CoverTab[51235]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:163
				// _ = "end of CoverTab[51235]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:163
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:163
			// _ = "end of CoverTab[51232]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:163
			_go_fuzz_dep_.CoverTab[51233]++
														in.Flags |= protoiface.MarshalUseCachedSize
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:164
			// _ = "end of CoverTab[51233]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:165
			_go_fuzz_dep_.CoverTab[51236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:165
			// _ = "end of CoverTab[51236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:165
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:165
		// _ = "end of CoverTab[51226]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:165
		_go_fuzz_dep_.CoverTab[51227]++
													out, err = methods.Marshal(in)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:166
		// _ = "end of CoverTab[51227]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:167
		_go_fuzz_dep_.CoverTab[51237]++
													out.Buf, err = o.marshalMessageSlow(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:168
		// _ = "end of CoverTab[51237]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:169
	// _ = "end of CoverTab[51217]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:169
	_go_fuzz_dep_.CoverTab[51218]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:170
		_go_fuzz_dep_.CoverTab[51238]++
													return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:171
		// _ = "end of CoverTab[51238]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:172
		_go_fuzz_dep_.CoverTab[51239]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:172
		// _ = "end of CoverTab[51239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:172
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:172
	// _ = "end of CoverTab[51218]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:172
	_go_fuzz_dep_.CoverTab[51219]++
												if allowPartial {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:173
		_go_fuzz_dep_.CoverTab[51240]++
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:174
		// _ = "end of CoverTab[51240]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:175
		_go_fuzz_dep_.CoverTab[51241]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:175
		// _ = "end of CoverTab[51241]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:175
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:175
	// _ = "end of CoverTab[51219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:175
	_go_fuzz_dep_.CoverTab[51220]++
												return out, checkInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:176
	// _ = "end of CoverTab[51220]"
}

func (o MarshalOptions) marshalMessage(b []byte, m protoreflect.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:179
	_go_fuzz_dep_.CoverTab[51242]++
												out, err := o.marshal(b, m)
												return out.Buf, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:181
	// _ = "end of CoverTab[51242]"
}

// growcap scales up the capacity of a slice.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:184
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:184
// Given a slice with a current capacity of oldcap and a desired
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:184
// capacity of wantcap, growcap returns a new capacity >= wantcap.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:184
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:184
// The algorithm is mostly identical to the one used by append as of Go 1.14.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:190
func growcap(oldcap, wantcap int) (newcap int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:190
	_go_fuzz_dep_.CoverTab[51243]++
												if wantcap > oldcap*2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:191
		_go_fuzz_dep_.CoverTab[51245]++
													newcap = wantcap
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:192
		// _ = "end of CoverTab[51245]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:193
		_go_fuzz_dep_.CoverTab[51246]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:193
		if oldcap < 1024 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:193
			_go_fuzz_dep_.CoverTab[51247]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:197
			newcap = oldcap * 2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:197
			// _ = "end of CoverTab[51247]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:198
			_go_fuzz_dep_.CoverTab[51248]++
														newcap = oldcap
														for 0 < newcap && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:200
				_go_fuzz_dep_.CoverTab[51250]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:200
				return newcap < wantcap
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:200
				// _ = "end of CoverTab[51250]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:200
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:200
				_go_fuzz_dep_.CoverTab[51251]++
															newcap += newcap / 4
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:201
				// _ = "end of CoverTab[51251]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:202
			// _ = "end of CoverTab[51248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:202
			_go_fuzz_dep_.CoverTab[51249]++
														if newcap <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:203
				_go_fuzz_dep_.CoverTab[51252]++
															newcap = wantcap
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:204
				// _ = "end of CoverTab[51252]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:205
				_go_fuzz_dep_.CoverTab[51253]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:205
				// _ = "end of CoverTab[51253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:205
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:205
			// _ = "end of CoverTab[51249]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:206
		// _ = "end of CoverTab[51246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:206
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:206
	// _ = "end of CoverTab[51243]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:206
	_go_fuzz_dep_.CoverTab[51244]++
												return newcap
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:207
	// _ = "end of CoverTab[51244]"
}

func (o MarshalOptions) marshalMessageSlow(b []byte, m protoreflect.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:210
	_go_fuzz_dep_.CoverTab[51254]++
												if messageset.IsMessageSet(m.Descriptor()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:211
		_go_fuzz_dep_.CoverTab[51259]++
													return o.marshalMessageSet(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:212
		// _ = "end of CoverTab[51259]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:213
		_go_fuzz_dep_.CoverTab[51260]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:213
		// _ = "end of CoverTab[51260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:213
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:213
	// _ = "end of CoverTab[51254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:213
	_go_fuzz_dep_.CoverTab[51255]++
												fieldOrder := order.AnyFieldOrder
												if o.Deterministic {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:215
		_go_fuzz_dep_.CoverTab[51261]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:219
		fieldOrder = order.LegacyFieldOrder
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:219
		// _ = "end of CoverTab[51261]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:220
		_go_fuzz_dep_.CoverTab[51262]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:220
		// _ = "end of CoverTab[51262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:220
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:220
	// _ = "end of CoverTab[51255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:220
	_go_fuzz_dep_.CoverTab[51256]++
												var err error
												order.RangeFields(m, fieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:222
		_go_fuzz_dep_.CoverTab[51263]++
													b, err = o.marshalField(b, fd, v)
													return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:224
		// _ = "end of CoverTab[51263]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:225
	// _ = "end of CoverTab[51256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:225
	_go_fuzz_dep_.CoverTab[51257]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:226
		_go_fuzz_dep_.CoverTab[51264]++
													return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:227
		// _ = "end of CoverTab[51264]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:228
		_go_fuzz_dep_.CoverTab[51265]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:228
		// _ = "end of CoverTab[51265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:228
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:228
	// _ = "end of CoverTab[51257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:228
	_go_fuzz_dep_.CoverTab[51258]++
												b = append(b, m.GetUnknown()...)
												return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:230
	// _ = "end of CoverTab[51258]"
}

func (o MarshalOptions) marshalField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:233
	_go_fuzz_dep_.CoverTab[51266]++
												switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:235
		_go_fuzz_dep_.CoverTab[51267]++
													return o.marshalList(b, fd, value.List())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:236
		// _ = "end of CoverTab[51267]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:237
		_go_fuzz_dep_.CoverTab[51268]++
													return o.marshalMap(b, fd, value.Map())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:238
		// _ = "end of CoverTab[51268]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:239
		_go_fuzz_dep_.CoverTab[51269]++
													b = protowire.AppendTag(b, fd.Number(), wireTypes[fd.Kind()])
													return o.marshalSingular(b, fd, value)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:241
		// _ = "end of CoverTab[51269]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:242
	// _ = "end of CoverTab[51266]"
}

func (o MarshalOptions) marshalList(b []byte, fd protoreflect.FieldDescriptor, list protoreflect.List) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:245
	_go_fuzz_dep_.CoverTab[51270]++
												if fd.IsPacked() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:246
		_go_fuzz_dep_.CoverTab[51273]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:246
		return list.Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:246
		// _ = "end of CoverTab[51273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:246
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:246
		_go_fuzz_dep_.CoverTab[51274]++
													b = protowire.AppendTag(b, fd.Number(), protowire.BytesType)
													b, pos := appendSpeculativeLength(b)
													for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:249
			_go_fuzz_dep_.CoverTab[51276]++
														var err error
														b, err = o.marshalSingular(b, fd, list.Get(i))
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:252
				_go_fuzz_dep_.CoverTab[51277]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:253
				// _ = "end of CoverTab[51277]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:254
				_go_fuzz_dep_.CoverTab[51278]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:254
				// _ = "end of CoverTab[51278]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:254
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:254
			// _ = "end of CoverTab[51276]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:255
		// _ = "end of CoverTab[51274]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:255
		_go_fuzz_dep_.CoverTab[51275]++
													b = finishSpeculativeLength(b, pos)
													return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:257
		// _ = "end of CoverTab[51275]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:258
		_go_fuzz_dep_.CoverTab[51279]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:258
		// _ = "end of CoverTab[51279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:258
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:258
	// _ = "end of CoverTab[51270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:258
	_go_fuzz_dep_.CoverTab[51271]++

												kind := fd.Kind()
												for i, llen := 0, list.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:261
		_go_fuzz_dep_.CoverTab[51280]++
													var err error
													b = protowire.AppendTag(b, fd.Number(), wireTypes[kind])
													b, err = o.marshalSingular(b, fd, list.Get(i))
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:265
			_go_fuzz_dep_.CoverTab[51281]++
														return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:266
			// _ = "end of CoverTab[51281]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:267
			_go_fuzz_dep_.CoverTab[51282]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:267
			// _ = "end of CoverTab[51282]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:267
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:267
		// _ = "end of CoverTab[51280]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:268
	// _ = "end of CoverTab[51271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:268
	_go_fuzz_dep_.CoverTab[51272]++
												return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:269
	// _ = "end of CoverTab[51272]"
}

func (o MarshalOptions) marshalMap(b []byte, fd protoreflect.FieldDescriptor, mapv protoreflect.Map) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:272
	_go_fuzz_dep_.CoverTab[51283]++
												keyf := fd.MapKey()
												valf := fd.MapValue()
												keyOrder := order.AnyKeyOrder
												if o.Deterministic {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:276
		_go_fuzz_dep_.CoverTab[51286]++
													keyOrder = order.GenericKeyOrder
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:277
		// _ = "end of CoverTab[51286]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:278
		_go_fuzz_dep_.CoverTab[51287]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:278
		// _ = "end of CoverTab[51287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:278
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:278
	// _ = "end of CoverTab[51283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:278
	_go_fuzz_dep_.CoverTab[51284]++
												var err error
												order.RangeEntries(mapv, keyOrder, func(key protoreflect.MapKey, value protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:280
		_go_fuzz_dep_.CoverTab[51288]++
													b = protowire.AppendTag(b, fd.Number(), protowire.BytesType)
													var pos int
													b, pos = appendSpeculativeLength(b)

													b, err = o.marshalField(b, keyf, key.Value())
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:286
			_go_fuzz_dep_.CoverTab[51291]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:287
			// _ = "end of CoverTab[51291]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:288
			_go_fuzz_dep_.CoverTab[51292]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:288
			// _ = "end of CoverTab[51292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:288
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:288
		// _ = "end of CoverTab[51288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:288
		_go_fuzz_dep_.CoverTab[51289]++
													b, err = o.marshalField(b, valf, value)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:290
			_go_fuzz_dep_.CoverTab[51293]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:291
			// _ = "end of CoverTab[51293]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:292
			_go_fuzz_dep_.CoverTab[51294]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:292
			// _ = "end of CoverTab[51294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:292
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:292
		// _ = "end of CoverTab[51289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:292
		_go_fuzz_dep_.CoverTab[51290]++
													b = finishSpeculativeLength(b, pos)
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:294
		// _ = "end of CoverTab[51290]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:295
	// _ = "end of CoverTab[51284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:295
	_go_fuzz_dep_.CoverTab[51285]++
												return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:296
	// _ = "end of CoverTab[51285]"
}

// When encoding length-prefixed fields, we speculatively set aside some number of bytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:299
// for the length, encode the data, and then encode the length (shifting the data if necessary
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:299
// to make room).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:302
const speculativeLength = 1

func appendSpeculativeLength(b []byte) ([]byte, int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:304
	_go_fuzz_dep_.CoverTab[51295]++
												pos := len(b)
												b = append(b, "\x00\x00\x00\x00"[:speculativeLength]...)
												return b, pos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:307
	// _ = "end of CoverTab[51295]"
}

func finishSpeculativeLength(b []byte, pos int) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:310
	_go_fuzz_dep_.CoverTab[51296]++
												mlen := len(b) - pos - speculativeLength
												msiz := protowire.SizeVarint(uint64(mlen))
												if msiz != speculativeLength {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:313
		_go_fuzz_dep_.CoverTab[51298]++
													for i := 0; i < msiz-speculativeLength; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:314
			_go_fuzz_dep_.CoverTab[51300]++
														b = append(b, 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:315
			// _ = "end of CoverTab[51300]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:316
		// _ = "end of CoverTab[51298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:316
		_go_fuzz_dep_.CoverTab[51299]++
													copy(b[pos+msiz:], b[pos+speculativeLength:])
													b = b[:pos+msiz+mlen]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:318
		// _ = "end of CoverTab[51299]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:319
		_go_fuzz_dep_.CoverTab[51301]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:319
		// _ = "end of CoverTab[51301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:319
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:319
	// _ = "end of CoverTab[51296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:319
	_go_fuzz_dep_.CoverTab[51297]++
												protowire.AppendVarint(b[:pos], uint64(mlen))
												return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:321
	// _ = "end of CoverTab[51297]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:322
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/encode.go:322
var _ = _go_fuzz_dep_.CoverTab
