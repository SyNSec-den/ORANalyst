// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:5
)

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)

// UnmarshalOptions configures the unmarshaler.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:19
// Example usage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:19
//	err := UnmarshalOptions{DiscardUnknown: true}.Unmarshal(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:24
type UnmarshalOptions struct {
	pragma.NoUnkeyedLiterals

	// Merge merges the input into the destination message.
	// The default behavior is to always reset the message before unmarshaling,
	// unless Merge is specified.
	Merge	bool

	// AllowPartial accepts input for messages that will result in missing
	// required fields. If AllowPartial is false (the default), Unmarshal will
	// return an error if there are any missing required fields.
	AllowPartial	bool

	// If DiscardUnknown is set, unknown fields are ignored.
	DiscardUnknown	bool

	// Resolver is used for looking up types when unmarshaling extension fields.
	// If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver	interface {
		FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)
		FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
	}

	// RecursionLimit limits how deeply messages may be nested.
	// If zero, a default limit is applied.
	RecursionLimit	int
}

// Unmarshal parses the wire-format message in b and places the result in m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:52
// The provided message must be mutable (e.g., a non-nil pointer to a message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:54
func Unmarshal(b []byte, m Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:54
	_go_fuzz_dep_.CoverTab[50628]++
												_, err := UnmarshalOptions{RecursionLimit: protowire.DefaultRecursionLimit}.unmarshal(b, m.ProtoReflect())
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:56
	// _ = "end of CoverTab[50628]"
}

// Unmarshal parses the wire-format message in b and places the result in m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:59
// The provided message must be mutable (e.g., a non-nil pointer to a message).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:61
func (o UnmarshalOptions) Unmarshal(b []byte, m Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:61
	_go_fuzz_dep_.CoverTab[50629]++
												if o.RecursionLimit == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:62
		_go_fuzz_dep_.CoverTab[50631]++
													o.RecursionLimit = protowire.DefaultRecursionLimit
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:63
		// _ = "end of CoverTab[50631]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:64
		_go_fuzz_dep_.CoverTab[50632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:64
		// _ = "end of CoverTab[50632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:64
	// _ = "end of CoverTab[50629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:64
	_go_fuzz_dep_.CoverTab[50630]++
												_, err := o.unmarshal(b, m.ProtoReflect())
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:66
	// _ = "end of CoverTab[50630]"
}

// UnmarshalState parses a wire-format message and places the result in m.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:69
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:69
// This method permits fine-grained control over the unmarshaler.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:69
// Most users should use Unmarshal instead.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:73
func (o UnmarshalOptions) UnmarshalState(in protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:73
	_go_fuzz_dep_.CoverTab[50633]++
												if o.RecursionLimit == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:74
		_go_fuzz_dep_.CoverTab[50635]++
													o.RecursionLimit = protowire.DefaultRecursionLimit
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:75
		// _ = "end of CoverTab[50635]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:76
		_go_fuzz_dep_.CoverTab[50636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:76
		// _ = "end of CoverTab[50636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:76
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:76
	// _ = "end of CoverTab[50633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:76
	_go_fuzz_dep_.CoverTab[50634]++
												return o.unmarshal(in.Buf, in.Message)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:77
	// _ = "end of CoverTab[50634]"
}

// unmarshal is a centralized function that all unmarshal operations go through.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:80
// For profiling purposes, avoid changing the name of this function or
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:80
// introducing other code paths for unmarshal that do not go through this.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:83
func (o UnmarshalOptions) unmarshal(b []byte, m protoreflect.Message) (out protoiface.UnmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:83
	_go_fuzz_dep_.CoverTab[50637]++
												if o.Resolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:84
		_go_fuzz_dep_.CoverTab[50643]++
													o.Resolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:85
		// _ = "end of CoverTab[50643]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:86
		_go_fuzz_dep_.CoverTab[50644]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:86
		// _ = "end of CoverTab[50644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:86
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:86
	// _ = "end of CoverTab[50637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:86
	_go_fuzz_dep_.CoverTab[50638]++
												if !o.Merge {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:87
		_go_fuzz_dep_.CoverTab[50645]++
													Reset(m.Interface())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:88
		// _ = "end of CoverTab[50645]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:89
		_go_fuzz_dep_.CoverTab[50646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:89
		// _ = "end of CoverTab[50646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:89
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:89
	// _ = "end of CoverTab[50638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:89
	_go_fuzz_dep_.CoverTab[50639]++
												allowPartial := o.AllowPartial
												o.Merge = true
												o.AllowPartial = true
												methods := protoMethods(m)
												if methods != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:94
		_go_fuzz_dep_.CoverTab[50647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:94
		return methods.Unmarshal != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:94
		// _ = "end of CoverTab[50647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:94
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:94
		_go_fuzz_dep_.CoverTab[50648]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:94
		return !(o.DiscardUnknown && func() bool {
														_go_fuzz_dep_.CoverTab[50649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:95
			return methods.Flags&protoiface.SupportUnmarshalDiscardUnknown == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:95
			// _ = "end of CoverTab[50649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:95
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:95
		// _ = "end of CoverTab[50648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:95
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:95
		_go_fuzz_dep_.CoverTab[50650]++
													in := protoiface.UnmarshalInput{
			Message:	m,
			Buf:		b,
			Resolver:	o.Resolver,
			Depth:		o.RecursionLimit,
		}
		if o.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:102
			_go_fuzz_dep_.CoverTab[50652]++
														in.Flags |= protoiface.UnmarshalDiscardUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:103
			// _ = "end of CoverTab[50652]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:104
			_go_fuzz_dep_.CoverTab[50653]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:104
			// _ = "end of CoverTab[50653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:104
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:104
		// _ = "end of CoverTab[50650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:104
		_go_fuzz_dep_.CoverTab[50651]++
													out, err = methods.Unmarshal(in)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:105
		// _ = "end of CoverTab[50651]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:106
		_go_fuzz_dep_.CoverTab[50654]++
													o.RecursionLimit--
													if o.RecursionLimit < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:108
			_go_fuzz_dep_.CoverTab[50656]++
														return out, errors.New("exceeded max recursion depth")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:109
			// _ = "end of CoverTab[50656]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:110
			_go_fuzz_dep_.CoverTab[50657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:110
			// _ = "end of CoverTab[50657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:110
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:110
		// _ = "end of CoverTab[50654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:110
		_go_fuzz_dep_.CoverTab[50655]++
													err = o.unmarshalMessageSlow(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:111
		// _ = "end of CoverTab[50655]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:112
	// _ = "end of CoverTab[50639]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:112
	_go_fuzz_dep_.CoverTab[50640]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:113
		_go_fuzz_dep_.CoverTab[50658]++
													return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:114
		// _ = "end of CoverTab[50658]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:115
		_go_fuzz_dep_.CoverTab[50659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:115
		// _ = "end of CoverTab[50659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:115
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:115
	// _ = "end of CoverTab[50640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:115
	_go_fuzz_dep_.CoverTab[50641]++
												if allowPartial || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:116
		_go_fuzz_dep_.CoverTab[50660]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:116
		return (out.Flags&protoiface.UnmarshalInitialized != 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:116
		// _ = "end of CoverTab[50660]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:116
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:116
		_go_fuzz_dep_.CoverTab[50661]++
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:117
		// _ = "end of CoverTab[50661]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:118
		_go_fuzz_dep_.CoverTab[50662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:118
		// _ = "end of CoverTab[50662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:118
	// _ = "end of CoverTab[50641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:118
	_go_fuzz_dep_.CoverTab[50642]++
												return out, checkInitialized(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:119
	// _ = "end of CoverTab[50642]"
}

func (o UnmarshalOptions) unmarshalMessage(b []byte, m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:122
	_go_fuzz_dep_.CoverTab[50663]++
												_, err := o.unmarshal(b, m)
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:124
	// _ = "end of CoverTab[50663]"
}

func (o UnmarshalOptions) unmarshalMessageSlow(b []byte, m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:127
	_go_fuzz_dep_.CoverTab[50664]++
												md := m.Descriptor()
												if messageset.IsMessageSet(md) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:129
		_go_fuzz_dep_.CoverTab[50667]++
													return o.unmarshalMessageSet(b, m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:130
		// _ = "end of CoverTab[50667]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:131
		_go_fuzz_dep_.CoverTab[50668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:131
		// _ = "end of CoverTab[50668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:131
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:131
	// _ = "end of CoverTab[50664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:131
	_go_fuzz_dep_.CoverTab[50665]++
												fields := md.Fields()
												for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:133
		_go_fuzz_dep_.CoverTab[50669]++

													num, wtyp, tagLen := protowire.ConsumeTag(b)
													if tagLen < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:136
			_go_fuzz_dep_.CoverTab[50676]++
														return errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:137
			// _ = "end of CoverTab[50676]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:138
			_go_fuzz_dep_.CoverTab[50677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:138
			// _ = "end of CoverTab[50677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:138
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:138
		// _ = "end of CoverTab[50669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:138
		_go_fuzz_dep_.CoverTab[50670]++
													if num > protowire.MaxValidNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:139
			_go_fuzz_dep_.CoverTab[50678]++
														return errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:140
			// _ = "end of CoverTab[50678]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:141
			_go_fuzz_dep_.CoverTab[50679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:141
			// _ = "end of CoverTab[50679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:141
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:141
		// _ = "end of CoverTab[50670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:141
		_go_fuzz_dep_.CoverTab[50671]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:144
		fd := fields.ByNumber(num)
		if fd == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:145
			_go_fuzz_dep_.CoverTab[50680]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:145
			return md.ExtensionRanges().Has(num)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:145
			// _ = "end of CoverTab[50680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:145
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:145
			_go_fuzz_dep_.CoverTab[50681]++
														extType, err := o.Resolver.FindExtensionByNumber(md.FullName(), num)
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:147
				_go_fuzz_dep_.CoverTab[50683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:147
				return err != protoregistry.NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:147
				// _ = "end of CoverTab[50683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:147
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:147
				_go_fuzz_dep_.CoverTab[50684]++
															return errors.New("%v: unable to resolve extension %v: %v", md.FullName(), num, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:148
				// _ = "end of CoverTab[50684]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:149
				_go_fuzz_dep_.CoverTab[50685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:149
				// _ = "end of CoverTab[50685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:149
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:149
			// _ = "end of CoverTab[50681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:149
			_go_fuzz_dep_.CoverTab[50682]++
														if extType != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:150
				_go_fuzz_dep_.CoverTab[50686]++
															fd = extType.TypeDescriptor()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:151
				// _ = "end of CoverTab[50686]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:152
				_go_fuzz_dep_.CoverTab[50687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:152
				// _ = "end of CoverTab[50687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:152
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:152
			// _ = "end of CoverTab[50682]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:153
			_go_fuzz_dep_.CoverTab[50688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:153
			// _ = "end of CoverTab[50688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:153
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:153
		// _ = "end of CoverTab[50671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:153
		_go_fuzz_dep_.CoverTab[50672]++
													var err error
													if fd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:155
			_go_fuzz_dep_.CoverTab[50689]++
														err = errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:156
			// _ = "end of CoverTab[50689]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:157
			_go_fuzz_dep_.CoverTab[50690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:157
			if flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:157
				_go_fuzz_dep_.CoverTab[50691]++
															if fd.IsWeak() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:158
					_go_fuzz_dep_.CoverTab[50692]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:158
					return fd.Message().IsPlaceholder()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:158
					// _ = "end of CoverTab[50692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:158
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:158
					_go_fuzz_dep_.CoverTab[50693]++
																err = errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:159
					// _ = "end of CoverTab[50693]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:160
					_go_fuzz_dep_.CoverTab[50694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:160
					// _ = "end of CoverTab[50694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:160
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:160
				// _ = "end of CoverTab[50691]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
				_go_fuzz_dep_.CoverTab[50695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
				// _ = "end of CoverTab[50695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
			// _ = "end of CoverTab[50690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
		// _ = "end of CoverTab[50672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:161
		_go_fuzz_dep_.CoverTab[50673]++

		// Parse the field value.
		var valLen int
		switch {
		case err != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:166
			_go_fuzz_dep_.CoverTab[50696]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:166
			// _ = "end of CoverTab[50696]"
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:167
			_go_fuzz_dep_.CoverTab[50697]++
														valLen, err = o.unmarshalList(b[tagLen:], wtyp, m.Mutable(fd).List(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:168
			// _ = "end of CoverTab[50697]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:169
			_go_fuzz_dep_.CoverTab[50698]++
														valLen, err = o.unmarshalMap(b[tagLen:], wtyp, m.Mutable(fd).Map(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:170
			// _ = "end of CoverTab[50698]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:171
			_go_fuzz_dep_.CoverTab[50699]++
														valLen, err = o.unmarshalSingular(b[tagLen:], wtyp, m, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:172
			// _ = "end of CoverTab[50699]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:173
		// _ = "end of CoverTab[50673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:173
		_go_fuzz_dep_.CoverTab[50674]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:174
			_go_fuzz_dep_.CoverTab[50700]++
														if err != errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:175
				_go_fuzz_dep_.CoverTab[50703]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:176
				// _ = "end of CoverTab[50703]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:177
				_go_fuzz_dep_.CoverTab[50704]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:177
				// _ = "end of CoverTab[50704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:177
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:177
			// _ = "end of CoverTab[50700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:177
			_go_fuzz_dep_.CoverTab[50701]++
														valLen = protowire.ConsumeFieldValue(num, wtyp, b[tagLen:])
														if valLen < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:179
				_go_fuzz_dep_.CoverTab[50705]++
															return errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:180
				// _ = "end of CoverTab[50705]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:181
				_go_fuzz_dep_.CoverTab[50706]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:181
				// _ = "end of CoverTab[50706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:181
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:181
			// _ = "end of CoverTab[50701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:181
			_go_fuzz_dep_.CoverTab[50702]++
														if !o.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:182
				_go_fuzz_dep_.CoverTab[50707]++
															m.SetUnknown(append(m.GetUnknown(), b[:tagLen+valLen]...))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:183
				// _ = "end of CoverTab[50707]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:184
				_go_fuzz_dep_.CoverTab[50708]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:184
				// _ = "end of CoverTab[50708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:184
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:184
			// _ = "end of CoverTab[50702]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:185
			_go_fuzz_dep_.CoverTab[50709]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:185
			// _ = "end of CoverTab[50709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:185
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:185
		// _ = "end of CoverTab[50674]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:185
		_go_fuzz_dep_.CoverTab[50675]++
													b = b[tagLen+valLen:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:186
		// _ = "end of CoverTab[50675]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:187
	// _ = "end of CoverTab[50665]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:187
	_go_fuzz_dep_.CoverTab[50666]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:188
	// _ = "end of CoverTab[50666]"
}

func (o UnmarshalOptions) unmarshalSingular(b []byte, wtyp protowire.Type, m protoreflect.Message, fd protoreflect.FieldDescriptor) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:191
	_go_fuzz_dep_.CoverTab[50710]++
												v, n, err := o.unmarshalScalar(b, wtyp, fd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:193
		_go_fuzz_dep_.CoverTab[50713]++
													return 0, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:194
		// _ = "end of CoverTab[50713]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:195
		_go_fuzz_dep_.CoverTab[50714]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:195
		// _ = "end of CoverTab[50714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:195
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:195
	// _ = "end of CoverTab[50710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:195
	_go_fuzz_dep_.CoverTab[50711]++
												switch fd.Kind() {
	case protoreflect.GroupKind, protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:197
		_go_fuzz_dep_.CoverTab[50715]++
													m2 := m.Mutable(fd).Message()
													if err := o.unmarshalMessage(v.Bytes(), m2); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:199
			_go_fuzz_dep_.CoverTab[50717]++
														return n, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:200
			// _ = "end of CoverTab[50717]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:201
			_go_fuzz_dep_.CoverTab[50718]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:201
			// _ = "end of CoverTab[50718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:201
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:201
		// _ = "end of CoverTab[50715]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:202
		_go_fuzz_dep_.CoverTab[50716]++

													m.Set(fd, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:204
		// _ = "end of CoverTab[50716]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:205
	// _ = "end of CoverTab[50711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:205
	_go_fuzz_dep_.CoverTab[50712]++
												return n, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:206
	// _ = "end of CoverTab[50712]"
}

func (o UnmarshalOptions) unmarshalMap(b []byte, wtyp protowire.Type, mapv protoreflect.Map, fd protoreflect.FieldDescriptor) (n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:209
	_go_fuzz_dep_.CoverTab[50719]++
												if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:210
		_go_fuzz_dep_.CoverTab[50726]++
													return 0, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:211
		// _ = "end of CoverTab[50726]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:212
		_go_fuzz_dep_.CoverTab[50727]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:212
		// _ = "end of CoverTab[50727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:212
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:212
	// _ = "end of CoverTab[50719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:212
	_go_fuzz_dep_.CoverTab[50720]++
												b, n = protowire.ConsumeBytes(b)
												if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:214
		_go_fuzz_dep_.CoverTab[50728]++
													return 0, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:215
		// _ = "end of CoverTab[50728]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:216
		_go_fuzz_dep_.CoverTab[50729]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:216
		// _ = "end of CoverTab[50729]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:216
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:216
	// _ = "end of CoverTab[50720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:216
	_go_fuzz_dep_.CoverTab[50721]++
												var (
		keyField	= fd.MapKey()
		valField	= fd.MapValue()
		key		protoreflect.Value
		val		protoreflect.Value
		haveKey		bool
		haveVal		bool
	)
	switch valField.Kind() {
	case protoreflect.GroupKind, protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:226
		_go_fuzz_dep_.CoverTab[50730]++
													val = mapv.NewValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:227
		// _ = "end of CoverTab[50730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:227
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:227
		_go_fuzz_dep_.CoverTab[50731]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:227
		// _ = "end of CoverTab[50731]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:228
	// _ = "end of CoverTab[50721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:228
	_go_fuzz_dep_.CoverTab[50722]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:231
	for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:231
		_go_fuzz_dep_.CoverTab[50732]++
													num, wtyp, n := protowire.ConsumeTag(b)
													if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:233
			_go_fuzz_dep_.CoverTab[50737]++
														return 0, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:234
			// _ = "end of CoverTab[50737]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:235
			_go_fuzz_dep_.CoverTab[50738]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:235
			// _ = "end of CoverTab[50738]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:235
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:235
		// _ = "end of CoverTab[50732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:235
		_go_fuzz_dep_.CoverTab[50733]++
													if num > protowire.MaxValidNumber {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:236
			_go_fuzz_dep_.CoverTab[50739]++
														return 0, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:237
			// _ = "end of CoverTab[50739]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:238
			_go_fuzz_dep_.CoverTab[50740]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:238
			// _ = "end of CoverTab[50740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:238
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:238
		// _ = "end of CoverTab[50733]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:238
		_go_fuzz_dep_.CoverTab[50734]++
													b = b[n:]
													err = errUnknown
													switch num {
		case genid.MapEntry_Key_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:242
			_go_fuzz_dep_.CoverTab[50741]++
														key, n, err = o.unmarshalScalar(b, wtyp, keyField)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:244
				_go_fuzz_dep_.CoverTab[50747]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:245
				// _ = "end of CoverTab[50747]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:246
				_go_fuzz_dep_.CoverTab[50748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:246
				// _ = "end of CoverTab[50748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:246
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:246
			// _ = "end of CoverTab[50741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:246
			_go_fuzz_dep_.CoverTab[50742]++
														haveKey = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:247
			// _ = "end of CoverTab[50742]"
		case genid.MapEntry_Value_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:248
			_go_fuzz_dep_.CoverTab[50743]++
														var v protoreflect.Value
														v, n, err = o.unmarshalScalar(b, wtyp, valField)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:251
				_go_fuzz_dep_.CoverTab[50749]++
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:252
				// _ = "end of CoverTab[50749]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:253
				_go_fuzz_dep_.CoverTab[50750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:253
				// _ = "end of CoverTab[50750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:253
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:253
			// _ = "end of CoverTab[50743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:253
			_go_fuzz_dep_.CoverTab[50744]++
														switch valField.Kind() {
			case protoreflect.GroupKind, protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:255
				_go_fuzz_dep_.CoverTab[50751]++
															if err := o.unmarshalMessage(v.Bytes(), val.Message()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:256
					_go_fuzz_dep_.CoverTab[50753]++
																return 0, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:257
					// _ = "end of CoverTab[50753]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:258
					_go_fuzz_dep_.CoverTab[50754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:258
					// _ = "end of CoverTab[50754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:258
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:258
				// _ = "end of CoverTab[50751]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:259
				_go_fuzz_dep_.CoverTab[50752]++
															val = v
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:260
				// _ = "end of CoverTab[50752]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:261
			// _ = "end of CoverTab[50744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:261
			_go_fuzz_dep_.CoverTab[50745]++
														haveVal = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:262
			// _ = "end of CoverTab[50745]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:262
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:262
			_go_fuzz_dep_.CoverTab[50746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:262
			// _ = "end of CoverTab[50746]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:263
		// _ = "end of CoverTab[50734]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:263
		_go_fuzz_dep_.CoverTab[50735]++
													if err == errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:264
			_go_fuzz_dep_.CoverTab[50755]++
														n = protowire.ConsumeFieldValue(num, wtyp, b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:266
				_go_fuzz_dep_.CoverTab[50756]++
															return 0, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:267
				// _ = "end of CoverTab[50756]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:268
				_go_fuzz_dep_.CoverTab[50757]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:268
				// _ = "end of CoverTab[50757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:268
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:268
			// _ = "end of CoverTab[50755]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:269
			_go_fuzz_dep_.CoverTab[50758]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:269
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:269
				_go_fuzz_dep_.CoverTab[50759]++
															return 0, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:270
				// _ = "end of CoverTab[50759]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
				_go_fuzz_dep_.CoverTab[50760]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
				// _ = "end of CoverTab[50760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
			// _ = "end of CoverTab[50758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
		// _ = "end of CoverTab[50735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:271
		_go_fuzz_dep_.CoverTab[50736]++
													b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:272
		// _ = "end of CoverTab[50736]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:273
	// _ = "end of CoverTab[50722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:273
	_go_fuzz_dep_.CoverTab[50723]++

												if !haveKey {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:275
		_go_fuzz_dep_.CoverTab[50761]++
													key = keyField.Default()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:276
		// _ = "end of CoverTab[50761]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:277
		_go_fuzz_dep_.CoverTab[50762]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:277
		// _ = "end of CoverTab[50762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:277
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:277
	// _ = "end of CoverTab[50723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:277
	_go_fuzz_dep_.CoverTab[50724]++
												if !haveVal {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:278
		_go_fuzz_dep_.CoverTab[50763]++
													switch valField.Kind() {
		case protoreflect.GroupKind, protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:280
			_go_fuzz_dep_.CoverTab[50764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:280
			// _ = "end of CoverTab[50764]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:281
			_go_fuzz_dep_.CoverTab[50765]++
														val = valField.Default()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:282
			// _ = "end of CoverTab[50765]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:283
		// _ = "end of CoverTab[50763]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:284
		_go_fuzz_dep_.CoverTab[50766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:284
		// _ = "end of CoverTab[50766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:284
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:284
	// _ = "end of CoverTab[50724]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:284
	_go_fuzz_dep_.CoverTab[50725]++
												mapv.Set(key.MapKey(), val)
												return n, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:286
	// _ = "end of CoverTab[50725]"
}

// errUnknown is used internally to indicate fields which should be added
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:289
// to the unknown field set of a message. It is never returned from an exported
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:289
// function.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:292
var errUnknown = errors.New("BUG: internal error (unknown)")

var errDecode = errors.New("cannot parse invalid wire-format data")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:294
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/decode.go:294
var _ = _go_fuzz_dep_.CoverTab
