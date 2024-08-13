// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
package jsonpb

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:5
)

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	protoV2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const wrapJSONUnmarshalV2 = false

// UnmarshalNext unmarshals the next JSON object from d into m.
func UnmarshalNext(d *json.Decoder, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:28
	_go_fuzz_dep_.CoverTab[66720]++
												return new(Unmarshaler).UnmarshalNext(d, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:29
	// _ = "end of CoverTab[66720]"
}

// Unmarshal unmarshals a JSON object from r into m.
func Unmarshal(r io.Reader, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:33
	_go_fuzz_dep_.CoverTab[66721]++
												return new(Unmarshaler).Unmarshal(r, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:34
	// _ = "end of CoverTab[66721]"
}

// UnmarshalString unmarshals a JSON object from s into m.
func UnmarshalString(s string, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:38
	_go_fuzz_dep_.CoverTab[66722]++
												return new(Unmarshaler).Unmarshal(strings.NewReader(s), m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:39
	// _ = "end of CoverTab[66722]"
}

// Unmarshaler is a configurable object for converting from a JSON
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:42
// representation to a protocol buffer object.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:44
type Unmarshaler struct {
	// AllowUnknownFields specifies whether to allow messages to contain
	// unknown JSON fields, as opposed to failing to unmarshal.
	AllowUnknownFields	bool

	// AnyResolver is used to resolve the google.protobuf.Any well-known type.
	// If unset, the global registry is used by default.
	AnyResolver	AnyResolver
}

// JSONPBUnmarshaler is implemented by protobuf messages that customize the way
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
// they are unmarshaled from JSON. Messages that implement this should also
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
// implement JSONPBMarshaler so that the custom format can be produced.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
// The JSON unmarshaling must follow the JSON to proto specification:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
//	https://developers.google.com/protocol-buffers/docs/proto3#json
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:54
// Deprecated: Custom types should implement protobuf reflection instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:62
type JSONPBUnmarshaler interface {
	UnmarshalJSONPB(*Unmarshaler, []byte) error
}

// Unmarshal unmarshals a JSON object from r into m.
func (u *Unmarshaler) Unmarshal(r io.Reader, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:67
	_go_fuzz_dep_.CoverTab[66723]++
												return u.UnmarshalNext(json.NewDecoder(r), m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:68
	// _ = "end of CoverTab[66723]"
}

// UnmarshalNext unmarshals the next JSON object from d into m.
func (u *Unmarshaler) UnmarshalNext(d *json.Decoder, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:72
	_go_fuzz_dep_.CoverTab[66724]++
												if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:73
		_go_fuzz_dep_.CoverTab[66729]++
													return errors.New("invalid nil message")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:74
		// _ = "end of CoverTab[66729]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:75
		_go_fuzz_dep_.CoverTab[66730]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:75
		// _ = "end of CoverTab[66730]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:75
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:75
	// _ = "end of CoverTab[66724]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:75
	_go_fuzz_dep_.CoverTab[66725]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:78
	raw := json.RawMessage{}
	if err := d.Decode(&raw); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:79
		_go_fuzz_dep_.CoverTab[66731]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:80
		// _ = "end of CoverTab[66731]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:81
		_go_fuzz_dep_.CoverTab[66732]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:81
		// _ = "end of CoverTab[66732]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:81
	// _ = "end of CoverTab[66725]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:81
	_go_fuzz_dep_.CoverTab[66726]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:85
	if jsu, ok := m.(JSONPBUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:85
		_go_fuzz_dep_.CoverTab[66733]++
													return jsu.UnmarshalJSONPB(u, raw)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:86
		// _ = "end of CoverTab[66733]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:87
		_go_fuzz_dep_.CoverTab[66734]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:87
		// _ = "end of CoverTab[66734]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:87
	// _ = "end of CoverTab[66726]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:87
	_go_fuzz_dep_.CoverTab[66727]++

												mr := proto.MessageReflect(m)

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:93
	if string(raw) == "null" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:93
		_go_fuzz_dep_.CoverTab[66735]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:93
		return mr.Descriptor().FullName() != "google.protobuf.Value"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:93
		// _ = "end of CoverTab[66735]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:93
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:93
		_go_fuzz_dep_.CoverTab[66736]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:94
		// _ = "end of CoverTab[66736]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:95
		_go_fuzz_dep_.CoverTab[66737]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:95
		// _ = "end of CoverTab[66737]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:95
	// _ = "end of CoverTab[66727]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:95
	_go_fuzz_dep_.CoverTab[66728]++

												if wrapJSONUnmarshalV2 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:97
		_go_fuzz_dep_.CoverTab[66738]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:101
		isEmpty := true
		mr.Range(func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:102
			_go_fuzz_dep_.CoverTab[66742]++
														isEmpty = false
														return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:104
			// _ = "end of CoverTab[66742]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:105
		// _ = "end of CoverTab[66738]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:105
		_go_fuzz_dep_.CoverTab[66739]++
													if !isEmpty {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:106
			_go_fuzz_dep_.CoverTab[66743]++

														mr = mr.New()

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:111
			dst := proto.MessageReflect(m)
			defer mr.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:112
				_go_fuzz_dep_.CoverTab[66744]++
															dst.Set(fd, v)
															return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:114
				// _ = "end of CoverTab[66744]"
			})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:115
			// _ = "end of CoverTab[66743]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:116
			_go_fuzz_dep_.CoverTab[66745]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:116
			// _ = "end of CoverTab[66745]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:116
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:116
		// _ = "end of CoverTab[66739]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:116
		_go_fuzz_dep_.CoverTab[66740]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:119
		opts := protojson.UnmarshalOptions{
			DiscardUnknown: u.AllowUnknownFields,
		}
		if u.AnyResolver != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:122
			_go_fuzz_dep_.CoverTab[66746]++
														opts.Resolver = anyResolver{u.AnyResolver}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:123
			// _ = "end of CoverTab[66746]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:124
			_go_fuzz_dep_.CoverTab[66747]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:124
			// _ = "end of CoverTab[66747]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:124
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:124
		// _ = "end of CoverTab[66740]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:124
		_go_fuzz_dep_.CoverTab[66741]++
													return opts.Unmarshal(raw, mr.Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:125
		// _ = "end of CoverTab[66741]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:126
		_go_fuzz_dep_.CoverTab[66748]++
													if err := u.unmarshalMessage(mr, raw); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:127
			_go_fuzz_dep_.CoverTab[66750]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:128
			// _ = "end of CoverTab[66750]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:129
			_go_fuzz_dep_.CoverTab[66751]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:129
			// _ = "end of CoverTab[66751]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:129
		// _ = "end of CoverTab[66748]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:129
		_go_fuzz_dep_.CoverTab[66749]++
													return protoV2.CheckInitialized(mr.Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:130
		// _ = "end of CoverTab[66749]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:131
	// _ = "end of CoverTab[66728]"
}

func (u *Unmarshaler) unmarshalMessage(m protoreflect.Message, in []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:134
	_go_fuzz_dep_.CoverTab[66752]++
												md := m.Descriptor()
												fds := md.Fields()

												if jsu, ok := proto.MessageV1(m.Interface()).(JSONPBUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:138
		_go_fuzz_dep_.CoverTab[66760]++
													return jsu.UnmarshalJSONPB(u, in)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:139
		// _ = "end of CoverTab[66760]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:140
		_go_fuzz_dep_.CoverTab[66761]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:140
		// _ = "end of CoverTab[66761]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:140
	// _ = "end of CoverTab[66752]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:140
	_go_fuzz_dep_.CoverTab[66753]++

												if string(in) == "null" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:142
		_go_fuzz_dep_.CoverTab[66762]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:142
		return md.FullName() != "google.protobuf.Value"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:142
		// _ = "end of CoverTab[66762]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:142
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:142
		_go_fuzz_dep_.CoverTab[66763]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:143
		// _ = "end of CoverTab[66763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:144
		_go_fuzz_dep_.CoverTab[66764]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:144
		// _ = "end of CoverTab[66764]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:144
	// _ = "end of CoverTab[66753]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:144
	_go_fuzz_dep_.CoverTab[66754]++

												switch wellKnownType(md.FullName()) {
	case "Any":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:147
		_go_fuzz_dep_.CoverTab[66765]++
													var jsonObject map[string]json.RawMessage
													if err := json.Unmarshal(in, &jsonObject); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:149
			_go_fuzz_dep_.CoverTab[66789]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:150
			// _ = "end of CoverTab[66789]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:151
			_go_fuzz_dep_.CoverTab[66790]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:151
			// _ = "end of CoverTab[66790]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:151
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:151
		// _ = "end of CoverTab[66765]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:151
		_go_fuzz_dep_.CoverTab[66766]++

													rawTypeURL, ok := jsonObject["@type"]
													if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:154
			_go_fuzz_dep_.CoverTab[66791]++
														return errors.New("Any JSON doesn't have '@type'")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:155
			// _ = "end of CoverTab[66791]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:156
			_go_fuzz_dep_.CoverTab[66792]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:156
			// _ = "end of CoverTab[66792]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:156
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:156
		// _ = "end of CoverTab[66766]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:156
		_go_fuzz_dep_.CoverTab[66767]++
													typeURL, err := unquoteString(string(rawTypeURL))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:158
			_go_fuzz_dep_.CoverTab[66793]++
														return fmt.Errorf("can't unmarshal Any's '@type': %q", rawTypeURL)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:159
			// _ = "end of CoverTab[66793]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:160
			_go_fuzz_dep_.CoverTab[66794]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:160
			// _ = "end of CoverTab[66794]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:160
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:160
		// _ = "end of CoverTab[66767]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:160
		_go_fuzz_dep_.CoverTab[66768]++
													m.Set(fds.ByNumber(1), protoreflect.ValueOfString(typeURL))

													var m2 protoreflect.Message
													if u.AnyResolver != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:164
			_go_fuzz_dep_.CoverTab[66795]++
														mi, err := u.AnyResolver.Resolve(typeURL)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:166
				_go_fuzz_dep_.CoverTab[66797]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:167
				// _ = "end of CoverTab[66797]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:168
				_go_fuzz_dep_.CoverTab[66798]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:168
				// _ = "end of CoverTab[66798]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:168
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:168
			// _ = "end of CoverTab[66795]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:168
			_go_fuzz_dep_.CoverTab[66796]++
														m2 = proto.MessageReflect(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:169
			// _ = "end of CoverTab[66796]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:170
			_go_fuzz_dep_.CoverTab[66799]++
														mt, err := protoregistry.GlobalTypes.FindMessageByURL(typeURL)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:172
				_go_fuzz_dep_.CoverTab[66801]++
															if err == protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:173
					_go_fuzz_dep_.CoverTab[66803]++
																return fmt.Errorf("could not resolve Any message type: %v", typeURL)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:174
					// _ = "end of CoverTab[66803]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:175
					_go_fuzz_dep_.CoverTab[66804]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:175
					// _ = "end of CoverTab[66804]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:175
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:175
				// _ = "end of CoverTab[66801]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:175
				_go_fuzz_dep_.CoverTab[66802]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:176
				// _ = "end of CoverTab[66802]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:177
				_go_fuzz_dep_.CoverTab[66805]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:177
				// _ = "end of CoverTab[66805]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:177
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:177
			// _ = "end of CoverTab[66799]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:177
			_go_fuzz_dep_.CoverTab[66800]++
														m2 = mt.New()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:178
			// _ = "end of CoverTab[66800]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:179
		// _ = "end of CoverTab[66768]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:179
		_go_fuzz_dep_.CoverTab[66769]++

													if wellKnownType(m2.Descriptor().FullName()) != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:181
			_go_fuzz_dep_.CoverTab[66806]++
														rawValue, ok := jsonObject["value"]
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:183
				_go_fuzz_dep_.CoverTab[66808]++
															return errors.New("Any JSON doesn't have 'value'")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:184
				// _ = "end of CoverTab[66808]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:185
				_go_fuzz_dep_.CoverTab[66809]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:185
				// _ = "end of CoverTab[66809]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:185
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:185
			// _ = "end of CoverTab[66806]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:185
			_go_fuzz_dep_.CoverTab[66807]++
														if err := u.unmarshalMessage(m2, rawValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:186
				_go_fuzz_dep_.CoverTab[66810]++
															return fmt.Errorf("can't unmarshal Any nested proto %v: %v", typeURL, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:187
				// _ = "end of CoverTab[66810]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:188
				_go_fuzz_dep_.CoverTab[66811]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:188
				// _ = "end of CoverTab[66811]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:188
			// _ = "end of CoverTab[66807]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:189
			_go_fuzz_dep_.CoverTab[66812]++
														delete(jsonObject, "@type")
														rawJSON, err := json.Marshal(jsonObject)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:192
				_go_fuzz_dep_.CoverTab[66814]++
															return fmt.Errorf("can't generate JSON for Any's nested proto to be unmarshaled: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:193
				// _ = "end of CoverTab[66814]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:194
				_go_fuzz_dep_.CoverTab[66815]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:194
				// _ = "end of CoverTab[66815]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:194
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:194
			// _ = "end of CoverTab[66812]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:194
			_go_fuzz_dep_.CoverTab[66813]++
														if err = u.unmarshalMessage(m2, rawJSON); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:195
				_go_fuzz_dep_.CoverTab[66816]++
															return fmt.Errorf("can't unmarshal Any nested proto %v: %v", typeURL, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:196
				// _ = "end of CoverTab[66816]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:197
				_go_fuzz_dep_.CoverTab[66817]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:197
				// _ = "end of CoverTab[66817]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:197
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:197
			// _ = "end of CoverTab[66813]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:198
		// _ = "end of CoverTab[66769]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:198
		_go_fuzz_dep_.CoverTab[66770]++

													rawWire, err := protoV2.Marshal(m2.Interface())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:201
			_go_fuzz_dep_.CoverTab[66818]++
														return fmt.Errorf("can't marshal proto %v into Any.Value: %v", typeURL, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:202
			// _ = "end of CoverTab[66818]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:203
			_go_fuzz_dep_.CoverTab[66819]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:203
			// _ = "end of CoverTab[66819]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:203
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:203
		// _ = "end of CoverTab[66770]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:203
		_go_fuzz_dep_.CoverTab[66771]++
													m.Set(fds.ByNumber(2), protoreflect.ValueOfBytes(rawWire))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:205
		// _ = "end of CoverTab[66771]"
	case "BoolValue", "BytesValue", "StringValue",
		"Int32Value", "UInt32Value", "FloatValue",
		"Int64Value", "UInt64Value", "DoubleValue":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:208
		_go_fuzz_dep_.CoverTab[66772]++
													fd := fds.ByNumber(1)
													v, err := u.unmarshalValue(m.NewField(fd), in, fd)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:211
			_go_fuzz_dep_.CoverTab[66820]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:212
			// _ = "end of CoverTab[66820]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:213
			_go_fuzz_dep_.CoverTab[66821]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:213
			// _ = "end of CoverTab[66821]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:213
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:213
		// _ = "end of CoverTab[66772]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:213
		_go_fuzz_dep_.CoverTab[66773]++
													m.Set(fd, v)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:215
		// _ = "end of CoverTab[66773]"
	case "Duration":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:216
		_go_fuzz_dep_.CoverTab[66774]++
													v, err := unquoteString(string(in))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:218
			_go_fuzz_dep_.CoverTab[66822]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:219
			// _ = "end of CoverTab[66822]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:220
			_go_fuzz_dep_.CoverTab[66823]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:220
			// _ = "end of CoverTab[66823]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:220
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:220
		// _ = "end of CoverTab[66774]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:220
		_go_fuzz_dep_.CoverTab[66775]++
													d, err := time.ParseDuration(v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:222
			_go_fuzz_dep_.CoverTab[66824]++
														return fmt.Errorf("bad Duration: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:223
			// _ = "end of CoverTab[66824]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:224
			_go_fuzz_dep_.CoverTab[66825]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:224
			// _ = "end of CoverTab[66825]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:224
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:224
		// _ = "end of CoverTab[66775]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:224
		_go_fuzz_dep_.CoverTab[66776]++

													sec := d.Nanoseconds() / 1e9
													nsec := d.Nanoseconds() % 1e9
													m.Set(fds.ByNumber(1), protoreflect.ValueOfInt64(int64(sec)))
													m.Set(fds.ByNumber(2), protoreflect.ValueOfInt32(int32(nsec)))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:230
		// _ = "end of CoverTab[66776]"
	case "Timestamp":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:231
		_go_fuzz_dep_.CoverTab[66777]++
													v, err := unquoteString(string(in))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:233
			_go_fuzz_dep_.CoverTab[66826]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:234
			// _ = "end of CoverTab[66826]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:235
			_go_fuzz_dep_.CoverTab[66827]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:235
			// _ = "end of CoverTab[66827]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:235
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:235
		// _ = "end of CoverTab[66777]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:235
		_go_fuzz_dep_.CoverTab[66778]++
													t, err := time.Parse(time.RFC3339Nano, v)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:237
			_go_fuzz_dep_.CoverTab[66828]++
														return fmt.Errorf("bad Timestamp: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:238
			// _ = "end of CoverTab[66828]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:239
			_go_fuzz_dep_.CoverTab[66829]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:239
			// _ = "end of CoverTab[66829]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:239
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:239
		// _ = "end of CoverTab[66778]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:239
		_go_fuzz_dep_.CoverTab[66779]++

													sec := t.Unix()
													nsec := t.Nanosecond()
													m.Set(fds.ByNumber(1), protoreflect.ValueOfInt64(int64(sec)))
													m.Set(fds.ByNumber(2), protoreflect.ValueOfInt32(int32(nsec)))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:245
		// _ = "end of CoverTab[66779]"
	case "Value":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:246
		_go_fuzz_dep_.CoverTab[66780]++
													switch {
		case string(in) == "null":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:248
			_go_fuzz_dep_.CoverTab[66830]++
														m.Set(fds.ByNumber(1), protoreflect.ValueOfEnum(0))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:249
			// _ = "end of CoverTab[66830]"
		case string(in) == "true":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:250
			_go_fuzz_dep_.CoverTab[66831]++
														m.Set(fds.ByNumber(4), protoreflect.ValueOfBool(true))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:251
			// _ = "end of CoverTab[66831]"
		case string(in) == "false":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:252
			_go_fuzz_dep_.CoverTab[66832]++
														m.Set(fds.ByNumber(4), protoreflect.ValueOfBool(false))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:253
			// _ = "end of CoverTab[66832]"
		case hasPrefixAndSuffix('"', in, '"'):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:254
			_go_fuzz_dep_.CoverTab[66833]++
														s, err := unquoteString(string(in))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:256
				_go_fuzz_dep_.CoverTab[66839]++
															return fmt.Errorf("unrecognized type for Value %q", in)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:257
				// _ = "end of CoverTab[66839]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:258
				_go_fuzz_dep_.CoverTab[66840]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:258
				// _ = "end of CoverTab[66840]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:258
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:258
			// _ = "end of CoverTab[66833]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:258
			_go_fuzz_dep_.CoverTab[66834]++
														m.Set(fds.ByNumber(3), protoreflect.ValueOfString(s))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:259
			// _ = "end of CoverTab[66834]"
		case hasPrefixAndSuffix('[', in, ']'):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:260
			_go_fuzz_dep_.CoverTab[66835]++
														v := m.Mutable(fds.ByNumber(6))
														return u.unmarshalMessage(v.Message(), in)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:262
			// _ = "end of CoverTab[66835]"
		case hasPrefixAndSuffix('{', in, '}'):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:263
			_go_fuzz_dep_.CoverTab[66836]++
														v := m.Mutable(fds.ByNumber(5))
														return u.unmarshalMessage(v.Message(), in)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:265
			// _ = "end of CoverTab[66836]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:266
			_go_fuzz_dep_.CoverTab[66837]++
														f, err := strconv.ParseFloat(string(in), 0)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:268
				_go_fuzz_dep_.CoverTab[66841]++
															return fmt.Errorf("unrecognized type for Value %q", in)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:269
				// _ = "end of CoverTab[66841]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:270
				_go_fuzz_dep_.CoverTab[66842]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:270
				// _ = "end of CoverTab[66842]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:270
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:270
			// _ = "end of CoverTab[66837]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:270
			_go_fuzz_dep_.CoverTab[66838]++
														m.Set(fds.ByNumber(2), protoreflect.ValueOfFloat64(f))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:271
			// _ = "end of CoverTab[66838]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:272
		// _ = "end of CoverTab[66780]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:272
		_go_fuzz_dep_.CoverTab[66781]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:273
		// _ = "end of CoverTab[66781]"
	case "ListValue":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:274
		_go_fuzz_dep_.CoverTab[66782]++
													var jsonArray []json.RawMessage
													if err := json.Unmarshal(in, &jsonArray); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:276
			_go_fuzz_dep_.CoverTab[66843]++
														return fmt.Errorf("bad ListValue: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:277
			// _ = "end of CoverTab[66843]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:278
			_go_fuzz_dep_.CoverTab[66844]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:278
			// _ = "end of CoverTab[66844]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:278
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:278
		// _ = "end of CoverTab[66782]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:278
		_go_fuzz_dep_.CoverTab[66783]++

													lv := m.Mutable(fds.ByNumber(1)).List()
													for _, raw := range jsonArray {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:281
			_go_fuzz_dep_.CoverTab[66845]++
														ve := lv.NewElement()
														if err := u.unmarshalMessage(ve.Message(), raw); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:283
				_go_fuzz_dep_.CoverTab[66847]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:284
				// _ = "end of CoverTab[66847]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:285
				_go_fuzz_dep_.CoverTab[66848]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:285
				// _ = "end of CoverTab[66848]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:285
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:285
			// _ = "end of CoverTab[66845]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:285
			_go_fuzz_dep_.CoverTab[66846]++
														lv.Append(ve)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:286
			// _ = "end of CoverTab[66846]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:287
		// _ = "end of CoverTab[66783]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:287
		_go_fuzz_dep_.CoverTab[66784]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:288
		// _ = "end of CoverTab[66784]"
	case "Struct":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:289
		_go_fuzz_dep_.CoverTab[66785]++
													var jsonObject map[string]json.RawMessage
													if err := json.Unmarshal(in, &jsonObject); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:291
			_go_fuzz_dep_.CoverTab[66849]++
														return fmt.Errorf("bad StructValue: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:292
			// _ = "end of CoverTab[66849]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:293
			_go_fuzz_dep_.CoverTab[66850]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:293
			// _ = "end of CoverTab[66850]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:293
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:293
		// _ = "end of CoverTab[66785]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:293
		_go_fuzz_dep_.CoverTab[66786]++

													mv := m.Mutable(fds.ByNumber(1)).Map()
													for key, raw := range jsonObject {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:296
			_go_fuzz_dep_.CoverTab[66851]++
														kv := protoreflect.ValueOf(key).MapKey()
														vv := mv.NewValue()
														if err := u.unmarshalMessage(vv.Message(), raw); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:299
				_go_fuzz_dep_.CoverTab[66853]++
															return fmt.Errorf("bad value in StructValue for key %q: %v", key, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:300
				// _ = "end of CoverTab[66853]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:301
				_go_fuzz_dep_.CoverTab[66854]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:301
				// _ = "end of CoverTab[66854]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:301
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:301
			// _ = "end of CoverTab[66851]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:301
			_go_fuzz_dep_.CoverTab[66852]++
														mv.Set(kv, vv)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:302
			// _ = "end of CoverTab[66852]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:303
		// _ = "end of CoverTab[66786]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:303
		_go_fuzz_dep_.CoverTab[66787]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:304
		// _ = "end of CoverTab[66787]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:304
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:304
		_go_fuzz_dep_.CoverTab[66788]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:304
		// _ = "end of CoverTab[66788]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:305
	// _ = "end of CoverTab[66754]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:305
	_go_fuzz_dep_.CoverTab[66755]++

												var jsonObject map[string]json.RawMessage
												if err := json.Unmarshal(in, &jsonObject); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:308
		_go_fuzz_dep_.CoverTab[66855]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:309
		// _ = "end of CoverTab[66855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:310
		_go_fuzz_dep_.CoverTab[66856]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:310
		// _ = "end of CoverTab[66856]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:310
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:310
	// _ = "end of CoverTab[66755]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:310
	_go_fuzz_dep_.CoverTab[66756]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:313
	for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:313
		_go_fuzz_dep_.CoverTab[66857]++
													fd := fds.Get(i)
													if fd.IsWeak() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:315
			_go_fuzz_dep_.CoverTab[66864]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:315
			return fd.Message().IsPlaceholder()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:315
			// _ = "end of CoverTab[66864]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:315
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:315
			_go_fuzz_dep_.CoverTab[66865]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:316
			// _ = "end of CoverTab[66865]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:317
			_go_fuzz_dep_.CoverTab[66866]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:317
			// _ = "end of CoverTab[66866]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:317
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:317
		// _ = "end of CoverTab[66857]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:317
		_go_fuzz_dep_.CoverTab[66858]++

		// Search for any raw JSON value associated with this field.
		var raw json.RawMessage
		name := string(fd.Name())
		if fd.Kind() == protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:322
			_go_fuzz_dep_.CoverTab[66867]++
														name = string(fd.Message().Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:323
			// _ = "end of CoverTab[66867]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:324
			_go_fuzz_dep_.CoverTab[66868]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:324
			// _ = "end of CoverTab[66868]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:324
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:324
		// _ = "end of CoverTab[66858]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:324
		_go_fuzz_dep_.CoverTab[66859]++
													if v, ok := jsonObject[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:325
			_go_fuzz_dep_.CoverTab[66869]++
														delete(jsonObject, name)
														raw = v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:327
			// _ = "end of CoverTab[66869]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:328
			_go_fuzz_dep_.CoverTab[66870]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:328
			// _ = "end of CoverTab[66870]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:328
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:328
		// _ = "end of CoverTab[66859]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:328
		_go_fuzz_dep_.CoverTab[66860]++
													name = string(fd.JSONName())
													if v, ok := jsonObject[name]; ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:330
			_go_fuzz_dep_.CoverTab[66871]++
														delete(jsonObject, name)
														raw = v
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:332
			// _ = "end of CoverTab[66871]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:333
			_go_fuzz_dep_.CoverTab[66872]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:333
			// _ = "end of CoverTab[66872]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:333
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:333
		// _ = "end of CoverTab[66860]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:333
		_go_fuzz_dep_.CoverTab[66861]++

													field := m.NewField(fd)

													if raw == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
			_go_fuzz_dep_.CoverTab[66873]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
			return (string(raw) == "null" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
				_go_fuzz_dep_.CoverTab[66874]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
				return !isSingularWellKnownValue(fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
				// _ = "end of CoverTab[66874]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
				_go_fuzz_dep_.CoverTab[66875]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
				return !isSingularJSONPBUnmarshaler(field, fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
				// _ = "end of CoverTab[66875]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
			}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
			// _ = "end of CoverTab[66873]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:337
			_go_fuzz_dep_.CoverTab[66876]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:338
			// _ = "end of CoverTab[66876]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:339
			_go_fuzz_dep_.CoverTab[66877]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:339
			// _ = "end of CoverTab[66877]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:339
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:339
		// _ = "end of CoverTab[66861]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:339
		_go_fuzz_dep_.CoverTab[66862]++
													v, err := u.unmarshalValue(field, raw, fd)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:341
			_go_fuzz_dep_.CoverTab[66878]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:342
			// _ = "end of CoverTab[66878]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:343
			_go_fuzz_dep_.CoverTab[66879]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:343
			// _ = "end of CoverTab[66879]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:343
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:343
		// _ = "end of CoverTab[66862]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:343
		_go_fuzz_dep_.CoverTab[66863]++
													m.Set(fd, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:344
		// _ = "end of CoverTab[66863]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:345
	// _ = "end of CoverTab[66756]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:345
	_go_fuzz_dep_.CoverTab[66757]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:348
	for name, raw := range jsonObject {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:348
		_go_fuzz_dep_.CoverTab[66880]++
													if !strings.HasPrefix(name, "[") || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:349
			_go_fuzz_dep_.CoverTab[66887]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:349
			return !strings.HasSuffix(name, "]")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:349
			// _ = "end of CoverTab[66887]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:349
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:349
			_go_fuzz_dep_.CoverTab[66888]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:350
			// _ = "end of CoverTab[66888]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:351
			_go_fuzz_dep_.CoverTab[66889]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:351
			// _ = "end of CoverTab[66889]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:351
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:351
		// _ = "end of CoverTab[66880]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:351
		_go_fuzz_dep_.CoverTab[66881]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:354
		xname := protoreflect.FullName(name[len("[") : len(name)-len("]")])
		xt, _ := protoregistry.GlobalTypes.FindExtensionByName(xname)
		if xt == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:356
			_go_fuzz_dep_.CoverTab[66890]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:356
			return isMessageSet(md)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:356
			// _ = "end of CoverTab[66890]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:356
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:356
			_go_fuzz_dep_.CoverTab[66891]++
														xt, _ = protoregistry.GlobalTypes.FindExtensionByName(xname.Append("message_set_extension"))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:357
			// _ = "end of CoverTab[66891]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:358
			_go_fuzz_dep_.CoverTab[66892]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:358
			// _ = "end of CoverTab[66892]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:358
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:358
		// _ = "end of CoverTab[66881]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:358
		_go_fuzz_dep_.CoverTab[66882]++
													if xt == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:359
			_go_fuzz_dep_.CoverTab[66893]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:360
			// _ = "end of CoverTab[66893]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:361
			_go_fuzz_dep_.CoverTab[66894]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:361
			// _ = "end of CoverTab[66894]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:361
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:361
		// _ = "end of CoverTab[66882]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:361
		_go_fuzz_dep_.CoverTab[66883]++
													delete(jsonObject, name)
													fd := xt.TypeDescriptor()
													if fd.ContainingMessage().FullName() != m.Descriptor().FullName() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:364
			_go_fuzz_dep_.CoverTab[66895]++
														return fmt.Errorf("extension field %q does not extend message %q", xname, m.Descriptor().FullName())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:365
			// _ = "end of CoverTab[66895]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:366
			_go_fuzz_dep_.CoverTab[66896]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:366
			// _ = "end of CoverTab[66896]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:366
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:366
		// _ = "end of CoverTab[66883]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:366
		_go_fuzz_dep_.CoverTab[66884]++

													field := m.NewField(fd)

													if raw == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
			_go_fuzz_dep_.CoverTab[66897]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
			return (string(raw) == "null" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
				_go_fuzz_dep_.CoverTab[66898]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
				return !isSingularWellKnownValue(fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
				// _ = "end of CoverTab[66898]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
				_go_fuzz_dep_.CoverTab[66899]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
				return !isSingularJSONPBUnmarshaler(field, fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
				// _ = "end of CoverTab[66899]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
			}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
			// _ = "end of CoverTab[66897]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:370
			_go_fuzz_dep_.CoverTab[66900]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:371
			// _ = "end of CoverTab[66900]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:372
			_go_fuzz_dep_.CoverTab[66901]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:372
			// _ = "end of CoverTab[66901]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:372
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:372
		// _ = "end of CoverTab[66884]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:372
		_go_fuzz_dep_.CoverTab[66885]++
													v, err := u.unmarshalValue(field, raw, fd)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:374
			_go_fuzz_dep_.CoverTab[66902]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:375
			// _ = "end of CoverTab[66902]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:376
			_go_fuzz_dep_.CoverTab[66903]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:376
			// _ = "end of CoverTab[66903]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:376
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:376
		// _ = "end of CoverTab[66885]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:376
		_go_fuzz_dep_.CoverTab[66886]++
													m.Set(fd, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:377
		// _ = "end of CoverTab[66886]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:378
	// _ = "end of CoverTab[66757]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:378
	_go_fuzz_dep_.CoverTab[66758]++

												if !u.AllowUnknownFields && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:380
		_go_fuzz_dep_.CoverTab[66904]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:380
		return len(jsonObject) > 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:380
		// _ = "end of CoverTab[66904]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:380
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:380
		_go_fuzz_dep_.CoverTab[66905]++
													for name := range jsonObject {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:381
			_go_fuzz_dep_.CoverTab[66906]++
														return fmt.Errorf("unknown field %q in %v", name, md.FullName())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:382
			// _ = "end of CoverTab[66906]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:383
		// _ = "end of CoverTab[66905]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:384
		_go_fuzz_dep_.CoverTab[66907]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:384
		// _ = "end of CoverTab[66907]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:384
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:384
	// _ = "end of CoverTab[66758]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:384
	_go_fuzz_dep_.CoverTab[66759]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:385
	// _ = "end of CoverTab[66759]"
}

func isSingularWellKnownValue(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:388
	_go_fuzz_dep_.CoverTab[66908]++
												if fd.Cardinality() == protoreflect.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:389
		_go_fuzz_dep_.CoverTab[66912]++
													return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:390
		// _ = "end of CoverTab[66912]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:391
		_go_fuzz_dep_.CoverTab[66913]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:391
		// _ = "end of CoverTab[66913]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:391
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:391
	// _ = "end of CoverTab[66908]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:391
	_go_fuzz_dep_.CoverTab[66909]++
												if md := fd.Message(); md != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:392
		_go_fuzz_dep_.CoverTab[66914]++
													return md.FullName() == "google.protobuf.Value"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:393
		// _ = "end of CoverTab[66914]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:394
		_go_fuzz_dep_.CoverTab[66915]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:394
		// _ = "end of CoverTab[66915]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:394
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:394
	// _ = "end of CoverTab[66909]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:394
	_go_fuzz_dep_.CoverTab[66910]++
												if ed := fd.Enum(); ed != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:395
		_go_fuzz_dep_.CoverTab[66916]++
													return ed.FullName() == "google.protobuf.NullValue"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:396
		// _ = "end of CoverTab[66916]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:397
		_go_fuzz_dep_.CoverTab[66917]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:397
		// _ = "end of CoverTab[66917]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:397
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:397
	// _ = "end of CoverTab[66910]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:397
	_go_fuzz_dep_.CoverTab[66911]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:398
	// _ = "end of CoverTab[66911]"
}

func isSingularJSONPBUnmarshaler(v protoreflect.Value, fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:401
	_go_fuzz_dep_.CoverTab[66918]++
												if fd.Message() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:402
		_go_fuzz_dep_.CoverTab[66920]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:402
		return fd.Cardinality() != protoreflect.Repeated
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:402
		// _ = "end of CoverTab[66920]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:402
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:402
		_go_fuzz_dep_.CoverTab[66921]++
													_, ok := proto.MessageV1(v.Interface()).(JSONPBUnmarshaler)
													return ok
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:404
		// _ = "end of CoverTab[66921]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:405
		_go_fuzz_dep_.CoverTab[66922]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:405
		// _ = "end of CoverTab[66922]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:405
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:405
	// _ = "end of CoverTab[66918]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:405
	_go_fuzz_dep_.CoverTab[66919]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:406
	// _ = "end of CoverTab[66919]"
}

func (u *Unmarshaler) unmarshalValue(v protoreflect.Value, in []byte, fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:409
	_go_fuzz_dep_.CoverTab[66923]++
												switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:411
		_go_fuzz_dep_.CoverTab[66924]++
													var jsonArray []json.RawMessage
													if err := json.Unmarshal(in, &jsonArray); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:413
			_go_fuzz_dep_.CoverTab[66931]++
														return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:414
			// _ = "end of CoverTab[66931]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:415
			_go_fuzz_dep_.CoverTab[66932]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:415
			// _ = "end of CoverTab[66932]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:415
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:415
		// _ = "end of CoverTab[66924]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:415
		_go_fuzz_dep_.CoverTab[66925]++
													lv := v.List()
													for _, raw := range jsonArray {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:417
			_go_fuzz_dep_.CoverTab[66933]++
														ve, err := u.unmarshalSingularValue(lv.NewElement(), raw, fd)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:419
				_go_fuzz_dep_.CoverTab[66935]++
															return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:420
				// _ = "end of CoverTab[66935]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:421
				_go_fuzz_dep_.CoverTab[66936]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:421
				// _ = "end of CoverTab[66936]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:421
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:421
			// _ = "end of CoverTab[66933]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:421
			_go_fuzz_dep_.CoverTab[66934]++
														lv.Append(ve)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:422
			// _ = "end of CoverTab[66934]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:423
		// _ = "end of CoverTab[66925]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:423
		_go_fuzz_dep_.CoverTab[66926]++
													return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:424
		// _ = "end of CoverTab[66926]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:425
		_go_fuzz_dep_.CoverTab[66927]++
													var jsonObject map[string]json.RawMessage
													if err := json.Unmarshal(in, &jsonObject); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:427
			_go_fuzz_dep_.CoverTab[66937]++
														return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:428
			// _ = "end of CoverTab[66937]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:429
			_go_fuzz_dep_.CoverTab[66938]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:429
			// _ = "end of CoverTab[66938]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:429
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:429
		// _ = "end of CoverTab[66927]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:429
		_go_fuzz_dep_.CoverTab[66928]++
													kfd := fd.MapKey()
													vfd := fd.MapValue()
													mv := v.Map()
													for key, raw := range jsonObject {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:433
			_go_fuzz_dep_.CoverTab[66939]++
														var kv protoreflect.MapKey
														if kfd.Kind() == protoreflect.StringKind {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:435
				_go_fuzz_dep_.CoverTab[66942]++
															kv = protoreflect.ValueOf(key).MapKey()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:436
				// _ = "end of CoverTab[66942]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:437
				_go_fuzz_dep_.CoverTab[66943]++
															v, err := u.unmarshalSingularValue(kfd.Default(), []byte(key), kfd)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:439
					_go_fuzz_dep_.CoverTab[66945]++
																return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:440
					// _ = "end of CoverTab[66945]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:441
					_go_fuzz_dep_.CoverTab[66946]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:441
					// _ = "end of CoverTab[66946]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:441
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:441
				// _ = "end of CoverTab[66943]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:441
				_go_fuzz_dep_.CoverTab[66944]++
															kv = v.MapKey()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:442
				// _ = "end of CoverTab[66944]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:443
			// _ = "end of CoverTab[66939]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:443
			_go_fuzz_dep_.CoverTab[66940]++

														vv, err := u.unmarshalSingularValue(mv.NewValue(), raw, vfd)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:446
				_go_fuzz_dep_.CoverTab[66947]++
															return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:447
				// _ = "end of CoverTab[66947]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:448
				_go_fuzz_dep_.CoverTab[66948]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:448
				// _ = "end of CoverTab[66948]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:448
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:448
			// _ = "end of CoverTab[66940]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:448
			_go_fuzz_dep_.CoverTab[66941]++
														mv.Set(kv, vv)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:449
			// _ = "end of CoverTab[66941]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:450
		// _ = "end of CoverTab[66928]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:450
		_go_fuzz_dep_.CoverTab[66929]++
													return v, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:451
		// _ = "end of CoverTab[66929]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:452
		_go_fuzz_dep_.CoverTab[66930]++
													return u.unmarshalSingularValue(v, in, fd)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:453
		// _ = "end of CoverTab[66930]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:454
	// _ = "end of CoverTab[66923]"
}

var nonFinite = map[string]float64{
	`"NaN"`:	math.NaN(),
	`"Infinity"`:	math.Inf(+1),
	`"-Infinity"`:	math.Inf(-1),
}

func (u *Unmarshaler) unmarshalSingularValue(v protoreflect.Value, in []byte, fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:463
	_go_fuzz_dep_.CoverTab[66949]++
												switch fd.Kind() {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:465
		_go_fuzz_dep_.CoverTab[66950]++
													return unmarshalValue(in, new(bool))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:466
		// _ = "end of CoverTab[66950]"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:467
		_go_fuzz_dep_.CoverTab[66951]++
													return unmarshalValue(trimQuote(in), new(int32))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:468
		// _ = "end of CoverTab[66951]"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:469
		_go_fuzz_dep_.CoverTab[66952]++
													return unmarshalValue(trimQuote(in), new(int64))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:470
		// _ = "end of CoverTab[66952]"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:471
		_go_fuzz_dep_.CoverTab[66953]++
													return unmarshalValue(trimQuote(in), new(uint32))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:472
		// _ = "end of CoverTab[66953]"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:473
		_go_fuzz_dep_.CoverTab[66954]++
													return unmarshalValue(trimQuote(in), new(uint64))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:474
		// _ = "end of CoverTab[66954]"
	case protoreflect.FloatKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:475
		_go_fuzz_dep_.CoverTab[66955]++
													if f, ok := nonFinite[string(in)]; ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:476
			_go_fuzz_dep_.CoverTab[66965]++
														return protoreflect.ValueOfFloat32(float32(f)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:477
			// _ = "end of CoverTab[66965]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:478
			_go_fuzz_dep_.CoverTab[66966]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:478
			// _ = "end of CoverTab[66966]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:478
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:478
		// _ = "end of CoverTab[66955]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:478
		_go_fuzz_dep_.CoverTab[66956]++
													return unmarshalValue(trimQuote(in), new(float32))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:479
		// _ = "end of CoverTab[66956]"
	case protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:480
		_go_fuzz_dep_.CoverTab[66957]++
													if f, ok := nonFinite[string(in)]; ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:481
			_go_fuzz_dep_.CoverTab[66967]++
														return protoreflect.ValueOfFloat64(float64(f)), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:482
			// _ = "end of CoverTab[66967]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:483
			_go_fuzz_dep_.CoverTab[66968]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:483
			// _ = "end of CoverTab[66968]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:483
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:483
		// _ = "end of CoverTab[66957]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:483
		_go_fuzz_dep_.CoverTab[66958]++
													return unmarshalValue(trimQuote(in), new(float64))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:484
		// _ = "end of CoverTab[66958]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:485
		_go_fuzz_dep_.CoverTab[66959]++
													return unmarshalValue(in, new(string))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:486
		// _ = "end of CoverTab[66959]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:487
		_go_fuzz_dep_.CoverTab[66960]++
													return unmarshalValue(in, new([]byte))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:488
		// _ = "end of CoverTab[66960]"
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:489
		_go_fuzz_dep_.CoverTab[66961]++
													if hasPrefixAndSuffix('"', in, '"') {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:490
			_go_fuzz_dep_.CoverTab[66969]++
														vd := fd.Enum().Values().ByName(protoreflect.Name(trimQuote(in)))
														if vd == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:492
				_go_fuzz_dep_.CoverTab[66971]++
															return v, fmt.Errorf("unknown value %q for enum %s", in, fd.Enum().FullName())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:493
				// _ = "end of CoverTab[66971]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:494
				_go_fuzz_dep_.CoverTab[66972]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:494
				// _ = "end of CoverTab[66972]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:494
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:494
			// _ = "end of CoverTab[66969]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:494
			_go_fuzz_dep_.CoverTab[66970]++
														return protoreflect.ValueOfEnum(vd.Number()), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:495
			// _ = "end of CoverTab[66970]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:496
			_go_fuzz_dep_.CoverTab[66973]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:496
			// _ = "end of CoverTab[66973]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:496
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:496
		// _ = "end of CoverTab[66961]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:496
		_go_fuzz_dep_.CoverTab[66962]++
													return unmarshalValue(in, new(protoreflect.EnumNumber))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:497
		// _ = "end of CoverTab[66962]"
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:498
		_go_fuzz_dep_.CoverTab[66963]++
													err := u.unmarshalMessage(v.Message(), in)
													return v, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:500
		// _ = "end of CoverTab[66963]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:501
		_go_fuzz_dep_.CoverTab[66964]++
													panic(fmt.Sprintf("invalid kind %v", fd.Kind()))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:502
		// _ = "end of CoverTab[66964]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:503
	// _ = "end of CoverTab[66949]"
}

func unmarshalValue(in []byte, v interface{}) (protoreflect.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:506
	_go_fuzz_dep_.CoverTab[66974]++
												err := json.Unmarshal(in, v)
												return protoreflect.ValueOf(reflect.ValueOf(v).Elem().Interface()), err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:508
	// _ = "end of CoverTab[66974]"
}

func unquoteString(in string) (out string, err error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:511
	_go_fuzz_dep_.CoverTab[66975]++
												err = json.Unmarshal([]byte(in), &out)
												return out, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:513
	// _ = "end of CoverTab[66975]"
}

func hasPrefixAndSuffix(prefix byte, in []byte, suffix byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:516
	_go_fuzz_dep_.CoverTab[66976]++
												if len(in) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		_go_fuzz_dep_.CoverTab[66978]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		return in[0] == prefix
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		// _ = "end of CoverTab[66978]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		_go_fuzz_dep_.CoverTab[66979]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		return in[len(in)-1] == suffix
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		// _ = "end of CoverTab[66979]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:517
		_go_fuzz_dep_.CoverTab[66980]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:518
		// _ = "end of CoverTab[66980]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:519
		_go_fuzz_dep_.CoverTab[66981]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:519
		// _ = "end of CoverTab[66981]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:519
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:519
	// _ = "end of CoverTab[66976]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:519
	_go_fuzz_dep_.CoverTab[66977]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:520
	// _ = "end of CoverTab[66977]"
}

// trimQuote is like unquoteString but simply strips surrounding quotes.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:523
// This is incorrect, but is behavior done by the legacy implementation.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:525
func trimQuote(in []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:525
	_go_fuzz_dep_.CoverTab[66982]++
												if len(in) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		_go_fuzz_dep_.CoverTab[66984]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		return in[0] == '"'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		// _ = "end of CoverTab[66984]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		_go_fuzz_dep_.CoverTab[66985]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		return in[len(in)-1] == '"'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		// _ = "end of CoverTab[66985]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:526
		_go_fuzz_dep_.CoverTab[66986]++
													in = in[1 : len(in)-1]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:527
		// _ = "end of CoverTab[66986]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:528
		_go_fuzz_dep_.CoverTab[66987]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:528
		// _ = "end of CoverTab[66987]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:528
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:528
	// _ = "end of CoverTab[66982]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:528
	_go_fuzz_dep_.CoverTab[66983]++
												return in
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:529
	// _ = "end of CoverTab[66983]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:530
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/decode.go:530
var _ = _go_fuzz_dep_.CoverTab
