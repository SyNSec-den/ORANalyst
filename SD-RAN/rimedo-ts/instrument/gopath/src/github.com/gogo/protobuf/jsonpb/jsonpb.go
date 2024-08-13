// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2015 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:32
/*
Package jsonpb provides marshaling and unmarshaling between protocol buffers and JSON.
It follows the specification at https://developers.google.com/protocol-buffers/docs/proto3#json.

This package produces a different output than the standard "encoding/json" package,
which does not operate correctly on protocol buffers.
*/
package jsonpb

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:39
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:39
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:39
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:39
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:39
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:39
)

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

const secondInNanos = int64(time.Second / time.Nanosecond)
const maxSecondsInDuration = 315576000000

// Marshaler is a configurable object for converting between
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:61
// protocol buffer objects and a JSON representation for them.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:63
type Marshaler struct {
	// Whether to render enum values as integers, as opposed to string values.
	EnumsAsInts	bool

	// Whether to render fields with zero values.
	EmitDefaults	bool

	// A string to indent each level by. The presence of this field will
	// also cause a space to appear between the field separator and
	// value, and for newlines to be appear between fields and array
	// elements.
	Indent	string

	// Whether to use the original (.proto) name for fields.
	OrigName	bool

	// A custom URL resolver to use when marshaling Any messages to JSON.
	// If unset, the default resolution strategy is to extract the
	// fully-qualified type name from the type URL and pass that to
	// proto.MessageType(string).
	AnyResolver	AnyResolver
}

// AnyResolver takes a type URL, present in an Any message, and resolves it into
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:86
// an instance of the associated message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:88
type AnyResolver interface {
	Resolve(typeUrl string) (proto.Message, error)
}

func defaultResolveAny(typeUrl string) (proto.Message, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:92
	_go_fuzz_dep_.CoverTab[141861]++

											mname := typeUrl
											if slash := strings.LastIndex(mname, "/"); slash >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:95
		_go_fuzz_dep_.CoverTab[141864]++
												mname = mname[slash+1:]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:96
		// _ = "end of CoverTab[141864]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:97
		_go_fuzz_dep_.CoverTab[141865]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:97
		// _ = "end of CoverTab[141865]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:97
	// _ = "end of CoverTab[141861]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:97
	_go_fuzz_dep_.CoverTab[141862]++
											mt := proto.MessageType(mname)
											if mt == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:99
		_go_fuzz_dep_.CoverTab[141866]++
												return nil, fmt.Errorf("unknown message type %q", mname)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:100
		// _ = "end of CoverTab[141866]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:101
		_go_fuzz_dep_.CoverTab[141867]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:101
		// _ = "end of CoverTab[141867]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:101
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:101
	// _ = "end of CoverTab[141862]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:101
	_go_fuzz_dep_.CoverTab[141863]++
											return reflect.New(mt.Elem()).Interface().(proto.Message), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:102
	// _ = "end of CoverTab[141863]"
}

// JSONPBMarshaler is implemented by protobuf messages that customize the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
// way they are marshaled to JSON. Messages that implement this should
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
// also implement JSONPBUnmarshaler so that the custom format can be
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
// parsed.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
// The JSON marshaling must follow the proto to JSON specification:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:105
//	https://developers.google.com/protocol-buffers/docs/proto3#json
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:112
type JSONPBMarshaler interface {
	MarshalJSONPB(*Marshaler) ([]byte, error)
}

// JSONPBUnmarshaler is implemented by protobuf messages that customize
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
// the way they are unmarshaled from JSON. Messages that implement this
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
// should also implement JSONPBMarshaler so that the custom format can be
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
// produced.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
// The JSON unmarshaling must follow the JSON to proto specification:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:116
//	https://developers.google.com/protocol-buffers/docs/proto3#json
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:123
type JSONPBUnmarshaler interface {
	UnmarshalJSONPB(*Unmarshaler, []byte) error
}

// Marshal marshals a protocol buffer into JSON.
func (m *Marshaler) Marshal(out io.Writer, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:128
	_go_fuzz_dep_.CoverTab[141868]++
											v := reflect.ValueOf(pb)
											if pb == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
		_go_fuzz_dep_.CoverTab[141871]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
		return (v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
			_go_fuzz_dep_.CoverTab[141872]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
			return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
			// _ = "end of CoverTab[141872]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
		}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
		// _ = "end of CoverTab[141871]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:130
		_go_fuzz_dep_.CoverTab[141873]++
												return errors.New("Marshal called with nil")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:131
		// _ = "end of CoverTab[141873]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:132
		_go_fuzz_dep_.CoverTab[141874]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:132
		// _ = "end of CoverTab[141874]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:132
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:132
	// _ = "end of CoverTab[141868]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:132
	_go_fuzz_dep_.CoverTab[141869]++

											if err := checkRequiredFields(pb); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:134
		_go_fuzz_dep_.CoverTab[141875]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:135
		// _ = "end of CoverTab[141875]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:136
		_go_fuzz_dep_.CoverTab[141876]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:136
		// _ = "end of CoverTab[141876]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:136
	// _ = "end of CoverTab[141869]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:136
	_go_fuzz_dep_.CoverTab[141870]++
											writer := &errWriter{writer: out}
											return m.marshalObject(writer, pb, "", "")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:138
	// _ = "end of CoverTab[141870]"
}

// MarshalToString converts a protocol buffer object to JSON string.
func (m *Marshaler) MarshalToString(pb proto.Message) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:142
	_go_fuzz_dep_.CoverTab[141877]++
											var buf bytes.Buffer
											if err := m.Marshal(&buf, pb); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:144
		_go_fuzz_dep_.CoverTab[141879]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:145
		// _ = "end of CoverTab[141879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:146
		_go_fuzz_dep_.CoverTab[141880]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:146
		// _ = "end of CoverTab[141880]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:146
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:146
	// _ = "end of CoverTab[141877]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:146
	_go_fuzz_dep_.CoverTab[141878]++
											return buf.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:147
	// _ = "end of CoverTab[141878]"
}

type int32Slice []int32

var nonFinite = map[string]float64{
	`"NaN"`:	math.NaN(),
	`"Infinity"`:	math.Inf(1),
	`"-Infinity"`:	math.Inf(-1),
}

// For sorting extensions ids to ensure stable output.
func (s int32Slice) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:159
	_go_fuzz_dep_.CoverTab[141881]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:159
	return len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:159
	// _ = "end of CoverTab[141881]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:159
}
func (s int32Slice) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:160
	_go_fuzz_dep_.CoverTab[141882]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:160
	return s[i] < s[j]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:160
	// _ = "end of CoverTab[141882]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:160
}
func (s int32Slice) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:161
	_go_fuzz_dep_.CoverTab[141883]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:161
	s[i], s[j] = s[j], s[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:161
	// _ = "end of CoverTab[141883]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:161
}

type isWkt interface {
	XXX_WellKnownType() string
}

var (
	wktType		= reflect.TypeOf((*isWkt)(nil)).Elem()
	messageType	= reflect.TypeOf((*proto.Message)(nil)).Elem()
)

// marshalObject writes a struct to the Writer.
func (m *Marshaler) marshalObject(out *errWriter, v proto.Message, indent, typeURL string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:173
	_go_fuzz_dep_.CoverTab[141884]++
											if jsm, ok := v.(JSONPBMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:174
		_go_fuzz_dep_.CoverTab[141892]++
												b, err := jsm.MarshalJSONPB(m)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:176
			_go_fuzz_dep_.CoverTab[141895]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:177
			// _ = "end of CoverTab[141895]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:178
			_go_fuzz_dep_.CoverTab[141896]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:178
			// _ = "end of CoverTab[141896]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:178
		// _ = "end of CoverTab[141892]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:178
		_go_fuzz_dep_.CoverTab[141893]++
												if typeURL != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:179
			_go_fuzz_dep_.CoverTab[141897]++
			// we are marshaling this object to an Any type
			var js map[string]*json.RawMessage
			if err = json.Unmarshal(b, &js); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:182
				_go_fuzz_dep_.CoverTab[141901]++
														return fmt.Errorf("type %T produced invalid JSON: %v", v, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:183
				// _ = "end of CoverTab[141901]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:184
				_go_fuzz_dep_.CoverTab[141902]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:184
				// _ = "end of CoverTab[141902]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:184
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:184
			// _ = "end of CoverTab[141897]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:184
			_go_fuzz_dep_.CoverTab[141898]++
													turl, err := json.Marshal(typeURL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:186
				_go_fuzz_dep_.CoverTab[141903]++
														return fmt.Errorf("failed to marshal type URL %q to JSON: %v", typeURL, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:187
				// _ = "end of CoverTab[141903]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:188
				_go_fuzz_dep_.CoverTab[141904]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:188
				// _ = "end of CoverTab[141904]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:188
			// _ = "end of CoverTab[141898]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:188
			_go_fuzz_dep_.CoverTab[141899]++
													js["@type"] = (*json.RawMessage)(&turl)
													if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:190
				_go_fuzz_dep_.CoverTab[141905]++
														b, err = json.MarshalIndent(js, indent, m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:191
				// _ = "end of CoverTab[141905]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:192
				_go_fuzz_dep_.CoverTab[141906]++
														b, err = json.Marshal(js)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:193
				// _ = "end of CoverTab[141906]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:194
			// _ = "end of CoverTab[141899]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:194
			_go_fuzz_dep_.CoverTab[141900]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:195
				_go_fuzz_dep_.CoverTab[141907]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:196
				// _ = "end of CoverTab[141907]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:197
				_go_fuzz_dep_.CoverTab[141908]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:197
				// _ = "end of CoverTab[141908]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:197
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:197
			// _ = "end of CoverTab[141900]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:198
			_go_fuzz_dep_.CoverTab[141909]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:198
			// _ = "end of CoverTab[141909]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:198
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:198
		// _ = "end of CoverTab[141893]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:198
		_go_fuzz_dep_.CoverTab[141894]++

												out.write(string(b))
												return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:201
		// _ = "end of CoverTab[141894]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:202
		_go_fuzz_dep_.CoverTab[141910]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:202
		// _ = "end of CoverTab[141910]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:202
	// _ = "end of CoverTab[141884]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:202
	_go_fuzz_dep_.CoverTab[141885]++

											s := reflect.ValueOf(v).Elem()

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:207
	if wkt, ok := v.(isWkt); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:207
		_go_fuzz_dep_.CoverTab[141911]++
												switch wkt.XXX_WellKnownType() {
		case "DoubleValue", "FloatValue", "Int64Value", "UInt64Value",
			"Int32Value", "UInt32Value", "BoolValue", "StringValue", "BytesValue":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:210
			_go_fuzz_dep_.CoverTab[141912]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:213
			sprop := proto.GetProperties(s.Type())
													return m.marshalValue(out, sprop.Prop[0], s.Field(0), indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:214
			// _ = "end of CoverTab[141912]"
		case "Any":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:215
			_go_fuzz_dep_.CoverTab[141913]++

													return m.marshalAny(out, v, indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:217
			// _ = "end of CoverTab[141913]"
		case "Duration":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:218
			_go_fuzz_dep_.CoverTab[141914]++
													s, ns := s.Field(0).Int(), s.Field(1).Int()
													if s < -maxSecondsInDuration || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:220
				_go_fuzz_dep_.CoverTab[141925]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:220
				return s > maxSecondsInDuration
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:220
				// _ = "end of CoverTab[141925]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:220
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:220
				_go_fuzz_dep_.CoverTab[141926]++
														return fmt.Errorf("seconds out of range %v", s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:221
				// _ = "end of CoverTab[141926]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:222
				_go_fuzz_dep_.CoverTab[141927]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:222
				// _ = "end of CoverTab[141927]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:222
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:222
			// _ = "end of CoverTab[141914]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:222
			_go_fuzz_dep_.CoverTab[141915]++
													if ns <= -secondInNanos || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:223
				_go_fuzz_dep_.CoverTab[141928]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:223
				return ns >= secondInNanos
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:223
				// _ = "end of CoverTab[141928]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:223
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:223
				_go_fuzz_dep_.CoverTab[141929]++
														return fmt.Errorf("ns out of range (%v, %v)", -secondInNanos, secondInNanos)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:224
				// _ = "end of CoverTab[141929]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:225
				_go_fuzz_dep_.CoverTab[141930]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:225
				// _ = "end of CoverTab[141930]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:225
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:225
			// _ = "end of CoverTab[141915]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:225
			_go_fuzz_dep_.CoverTab[141916]++
													if (s > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				_go_fuzz_dep_.CoverTab[141931]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				return ns < 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				// _ = "end of CoverTab[141931]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
			}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				_go_fuzz_dep_.CoverTab[141932]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				return (s < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
					_go_fuzz_dep_.CoverTab[141933]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
					return ns > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
					// _ = "end of CoverTab[141933]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				// _ = "end of CoverTab[141932]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:226
				_go_fuzz_dep_.CoverTab[141934]++
														return errors.New("signs of seconds and nanos do not match")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:227
				// _ = "end of CoverTab[141934]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:228
				_go_fuzz_dep_.CoverTab[141935]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:228
				// _ = "end of CoverTab[141935]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:228
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:228
			// _ = "end of CoverTab[141916]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:228
			_go_fuzz_dep_.CoverTab[141917]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:231
			f := "%d.%09d"
			if ns < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:232
				_go_fuzz_dep_.CoverTab[141936]++
														ns = -ns
														if s == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:234
					_go_fuzz_dep_.CoverTab[141937]++
															f = "-%d.%09d"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:235
					// _ = "end of CoverTab[141937]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:236
					_go_fuzz_dep_.CoverTab[141938]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:236
					// _ = "end of CoverTab[141938]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:236
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:236
				// _ = "end of CoverTab[141936]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:237
				_go_fuzz_dep_.CoverTab[141939]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:237
				// _ = "end of CoverTab[141939]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:237
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:237
			// _ = "end of CoverTab[141917]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:237
			_go_fuzz_dep_.CoverTab[141918]++
													x := fmt.Sprintf(f, s, ns)
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, ".000")
													out.write(`"`)
													out.write(x)
													out.write(`s"`)
													return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:245
			// _ = "end of CoverTab[141918]"
		case "Struct", "ListValue":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:246
			_go_fuzz_dep_.CoverTab[141919]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:249
			return m.marshalValue(out, &proto.Properties{}, s.Field(0), indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:249
			// _ = "end of CoverTab[141919]"
		case "Timestamp":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:250
			_go_fuzz_dep_.CoverTab[141920]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:253
			s, ns := s.Field(0).Int(), s.Field(1).Int()
			if ns < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:254
				_go_fuzz_dep_.CoverTab[141940]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:254
				return ns >= secondInNanos
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:254
				// _ = "end of CoverTab[141940]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:254
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:254
				_go_fuzz_dep_.CoverTab[141941]++
														return fmt.Errorf("ns out of range [0, %v)", secondInNanos)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:255
				// _ = "end of CoverTab[141941]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:256
				_go_fuzz_dep_.CoverTab[141942]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:256
				// _ = "end of CoverTab[141942]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:256
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:256
			// _ = "end of CoverTab[141920]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:256
			_go_fuzz_dep_.CoverTab[141921]++
													t := time.Unix(s, ns).UTC()

													x := t.Format("2006-01-02T15:04:05.000000000")
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, ".000")
													out.write(`"`)
													out.write(x)
													out.write(`Z"`)
													return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:266
			// _ = "end of CoverTab[141921]"
		case "Value":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:267
			_go_fuzz_dep_.CoverTab[141922]++

													kind := s.Field(0)
													if kind.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:270
				_go_fuzz_dep_.CoverTab[141943]++

														return errors.New("nil Value")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:272
				// _ = "end of CoverTab[141943]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:273
				_go_fuzz_dep_.CoverTab[141944]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:273
				// _ = "end of CoverTab[141944]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:273
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:273
			// _ = "end of CoverTab[141922]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:273
			_go_fuzz_dep_.CoverTab[141923]++

													x := kind.Elem().Elem().Field(0)

													return m.marshalValue(out, &proto.Properties{}, x, indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:277
			// _ = "end of CoverTab[141923]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:277
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:277
			_go_fuzz_dep_.CoverTab[141924]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:277
			// _ = "end of CoverTab[141924]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:278
		// _ = "end of CoverTab[141911]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:279
		_go_fuzz_dep_.CoverTab[141945]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:279
		// _ = "end of CoverTab[141945]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:279
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:279
	// _ = "end of CoverTab[141885]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:279
	_go_fuzz_dep_.CoverTab[141886]++

											out.write("{")
											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:282
		_go_fuzz_dep_.CoverTab[141946]++
												out.write("\n")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:283
		// _ = "end of CoverTab[141946]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:284
		_go_fuzz_dep_.CoverTab[141947]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:284
		// _ = "end of CoverTab[141947]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:284
	// _ = "end of CoverTab[141886]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:284
	_go_fuzz_dep_.CoverTab[141887]++

											firstField := true

											if typeURL != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:288
		_go_fuzz_dep_.CoverTab[141948]++
												if err := m.marshalTypeURL(out, indent, typeURL); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:289
			_go_fuzz_dep_.CoverTab[141950]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:290
			// _ = "end of CoverTab[141950]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:291
			_go_fuzz_dep_.CoverTab[141951]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:291
			// _ = "end of CoverTab[141951]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:291
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:291
		// _ = "end of CoverTab[141948]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:291
		_go_fuzz_dep_.CoverTab[141949]++
												firstField = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:292
		// _ = "end of CoverTab[141949]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:293
		_go_fuzz_dep_.CoverTab[141952]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:293
		// _ = "end of CoverTab[141952]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:293
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:293
	// _ = "end of CoverTab[141887]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:293
	_go_fuzz_dep_.CoverTab[141888]++

											for i := 0; i < s.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:295
		_go_fuzz_dep_.CoverTab[141953]++
												value := s.Field(i)
												valueField := s.Type().Field(i)
												if strings.HasPrefix(valueField.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:298
			_go_fuzz_dep_.CoverTab[141962]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:299
			// _ = "end of CoverTab[141962]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:300
			_go_fuzz_dep_.CoverTab[141963]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:300
			// _ = "end of CoverTab[141963]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:300
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:300
		// _ = "end of CoverTab[141953]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:300
		_go_fuzz_dep_.CoverTab[141954]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:303
		if valueField.Tag.Get("protobuf") == "" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:303
			_go_fuzz_dep_.CoverTab[141964]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:303
			return valueField.Tag.Get("protobuf_oneof") == ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:303
			// _ = "end of CoverTab[141964]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:303
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:303
			_go_fuzz_dep_.CoverTab[141965]++
													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:304
			// _ = "end of CoverTab[141965]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:305
			_go_fuzz_dep_.CoverTab[141966]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:305
			// _ = "end of CoverTab[141966]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:305
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:305
		// _ = "end of CoverTab[141954]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:305
		_go_fuzz_dep_.CoverTab[141955]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:308
		switch value.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:309
			_go_fuzz_dep_.CoverTab[141967]++
													if value.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:310
				_go_fuzz_dep_.CoverTab[141969]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:311
				// _ = "end of CoverTab[141969]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
				_go_fuzz_dep_.CoverTab[141970]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
				// _ = "end of CoverTab[141970]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
			// _ = "end of CoverTab[141967]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
			_go_fuzz_dep_.CoverTab[141968]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:312
			// _ = "end of CoverTab[141968]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:313
		// _ = "end of CoverTab[141955]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:313
		_go_fuzz_dep_.CoverTab[141956]++

												if !m.EmitDefaults {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:315
			_go_fuzz_dep_.CoverTab[141971]++
													switch value.Kind() {
			case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:317
				_go_fuzz_dep_.CoverTab[141972]++
														if !value.Bool() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:318
					_go_fuzz_dep_.CoverTab[141979]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:319
					// _ = "end of CoverTab[141979]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:320
					_go_fuzz_dep_.CoverTab[141980]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:320
					// _ = "end of CoverTab[141980]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:320
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:320
				// _ = "end of CoverTab[141972]"
			case reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:321
				_go_fuzz_dep_.CoverTab[141973]++
														if value.Int() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:322
					_go_fuzz_dep_.CoverTab[141981]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:323
					// _ = "end of CoverTab[141981]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:324
					_go_fuzz_dep_.CoverTab[141982]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:324
					// _ = "end of CoverTab[141982]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:324
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:324
				// _ = "end of CoverTab[141973]"
			case reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:325
				_go_fuzz_dep_.CoverTab[141974]++
														if value.Uint() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:326
					_go_fuzz_dep_.CoverTab[141983]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:327
					// _ = "end of CoverTab[141983]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:328
					_go_fuzz_dep_.CoverTab[141984]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:328
					// _ = "end of CoverTab[141984]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:328
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:328
				// _ = "end of CoverTab[141974]"
			case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:329
				_go_fuzz_dep_.CoverTab[141975]++
														if value.Float() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:330
					_go_fuzz_dep_.CoverTab[141985]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:331
					// _ = "end of CoverTab[141985]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:332
					_go_fuzz_dep_.CoverTab[141986]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:332
					// _ = "end of CoverTab[141986]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:332
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:332
				// _ = "end of CoverTab[141975]"
			case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:333
				_go_fuzz_dep_.CoverTab[141976]++
														if value.Len() == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:334
					_go_fuzz_dep_.CoverTab[141987]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:335
					// _ = "end of CoverTab[141987]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:336
					_go_fuzz_dep_.CoverTab[141988]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:336
					// _ = "end of CoverTab[141988]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:336
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:336
				// _ = "end of CoverTab[141976]"
			case reflect.Map, reflect.Ptr, reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:337
				_go_fuzz_dep_.CoverTab[141977]++
														if value.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:338
					_go_fuzz_dep_.CoverTab[141989]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:339
					// _ = "end of CoverTab[141989]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
					_go_fuzz_dep_.CoverTab[141990]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
					// _ = "end of CoverTab[141990]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
				// _ = "end of CoverTab[141977]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
			default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
				_go_fuzz_dep_.CoverTab[141978]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:340
				// _ = "end of CoverTab[141978]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:341
			// _ = "end of CoverTab[141971]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:342
			_go_fuzz_dep_.CoverTab[141991]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:342
			// _ = "end of CoverTab[141991]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:342
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:342
		// _ = "end of CoverTab[141956]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:342
		_go_fuzz_dep_.CoverTab[141957]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:345
		if valueField.Tag.Get("protobuf_oneof") != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:345
			_go_fuzz_dep_.CoverTab[141992]++

													sv := value.Elem().Elem()
													value = sv.Field(0)
													valueField = sv.Type().Field(0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:349
			// _ = "end of CoverTab[141992]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:350
			_go_fuzz_dep_.CoverTab[141993]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:350
			// _ = "end of CoverTab[141993]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:350
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:350
		// _ = "end of CoverTab[141957]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:350
		_go_fuzz_dep_.CoverTab[141958]++
												prop := jsonProperties(valueField, m.OrigName)
												if !firstField {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:352
			_go_fuzz_dep_.CoverTab[141994]++
													m.writeSep(out)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:353
			// _ = "end of CoverTab[141994]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:354
			_go_fuzz_dep_.CoverTab[141995]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:354
			// _ = "end of CoverTab[141995]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:354
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:354
		// _ = "end of CoverTab[141958]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:354
		_go_fuzz_dep_.CoverTab[141959]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:359
		if value.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:359
			_go_fuzz_dep_.CoverTab[141996]++
													if tag := valueField.Tag.Get("protobuf"); tag != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:360
				_go_fuzz_dep_.CoverTab[141997]++
														for _, v := range strings.Split(tag, ",") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:361
					_go_fuzz_dep_.CoverTab[141998]++
															if !strings.HasPrefix(v, "castvaluetype=") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:362
						_go_fuzz_dep_.CoverTab[142000]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:363
						// _ = "end of CoverTab[142000]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:364
						_go_fuzz_dep_.CoverTab[142001]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:364
						// _ = "end of CoverTab[142001]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:364
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:364
					// _ = "end of CoverTab[141998]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:364
					_go_fuzz_dep_.CoverTab[141999]++
															v = strings.TrimPrefix(v, "castvaluetype=")
															prop.MapValProp.CustomType = v
															break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:367
					// _ = "end of CoverTab[141999]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:368
				// _ = "end of CoverTab[141997]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:369
				_go_fuzz_dep_.CoverTab[142002]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:369
				// _ = "end of CoverTab[142002]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:369
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:369
			// _ = "end of CoverTab[141996]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:370
			_go_fuzz_dep_.CoverTab[142003]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:370
			// _ = "end of CoverTab[142003]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:370
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:370
		// _ = "end of CoverTab[141959]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:370
		_go_fuzz_dep_.CoverTab[141960]++
												if err := m.marshalField(out, prop, value, indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:371
			_go_fuzz_dep_.CoverTab[142004]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:372
			// _ = "end of CoverTab[142004]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:373
			_go_fuzz_dep_.CoverTab[142005]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:373
			// _ = "end of CoverTab[142005]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:373
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:373
		// _ = "end of CoverTab[141960]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:373
		_go_fuzz_dep_.CoverTab[141961]++
												firstField = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:374
		// _ = "end of CoverTab[141961]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:375
	// _ = "end of CoverTab[141888]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:375
	_go_fuzz_dep_.CoverTab[141889]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:378
	if ep, ok := v.(proto.Message); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:378
		_go_fuzz_dep_.CoverTab[142006]++
												extensions := proto.RegisteredExtensions(v)

												ids := make([]int32, 0, len(extensions))
												for id, desc := range extensions {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:382
			_go_fuzz_dep_.CoverTab[142008]++
													if !proto.HasExtension(ep, desc) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:383
				_go_fuzz_dep_.CoverTab[142010]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:384
				// _ = "end of CoverTab[142010]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:385
				_go_fuzz_dep_.CoverTab[142011]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:385
				// _ = "end of CoverTab[142011]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:385
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:385
			// _ = "end of CoverTab[142008]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:385
			_go_fuzz_dep_.CoverTab[142009]++
													ids = append(ids, id)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:386
			// _ = "end of CoverTab[142009]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:387
		// _ = "end of CoverTab[142006]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:387
		_go_fuzz_dep_.CoverTab[142007]++
												sort.Sort(int32Slice(ids))
												for _, id := range ids {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:389
			_go_fuzz_dep_.CoverTab[142012]++
													desc := extensions[id]
													if desc == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:391
				_go_fuzz_dep_.CoverTab[142017]++

														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:393
				// _ = "end of CoverTab[142017]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:394
				_go_fuzz_dep_.CoverTab[142018]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:394
				// _ = "end of CoverTab[142018]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:394
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:394
			// _ = "end of CoverTab[142012]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:394
			_go_fuzz_dep_.CoverTab[142013]++
													ext, extErr := proto.GetExtension(ep, desc)
													if extErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:396
				_go_fuzz_dep_.CoverTab[142019]++
														return extErr
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:397
				// _ = "end of CoverTab[142019]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:398
				_go_fuzz_dep_.CoverTab[142020]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:398
				// _ = "end of CoverTab[142020]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:398
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:398
			// _ = "end of CoverTab[142013]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:398
			_go_fuzz_dep_.CoverTab[142014]++
													value := reflect.ValueOf(ext)
													var prop proto.Properties
													prop.Parse(desc.Tag)
													prop.JSONName = fmt.Sprintf("[%s]", desc.Name)
													if !firstField {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:403
				_go_fuzz_dep_.CoverTab[142021]++
														m.writeSep(out)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:404
				// _ = "end of CoverTab[142021]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:405
				_go_fuzz_dep_.CoverTab[142022]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:405
				// _ = "end of CoverTab[142022]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:405
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:405
			// _ = "end of CoverTab[142014]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:405
			_go_fuzz_dep_.CoverTab[142015]++
													if err := m.marshalField(out, &prop, value, indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:406
				_go_fuzz_dep_.CoverTab[142023]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:407
				// _ = "end of CoverTab[142023]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:408
				_go_fuzz_dep_.CoverTab[142024]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:408
				// _ = "end of CoverTab[142024]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:408
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:408
			// _ = "end of CoverTab[142015]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:408
			_go_fuzz_dep_.CoverTab[142016]++
													firstField = false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:409
			// _ = "end of CoverTab[142016]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:410
		// _ = "end of CoverTab[142007]"

	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:412
		_go_fuzz_dep_.CoverTab[142025]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:412
		// _ = "end of CoverTab[142025]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:412
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:412
	// _ = "end of CoverTab[141889]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:412
	_go_fuzz_dep_.CoverTab[141890]++

											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:414
		_go_fuzz_dep_.CoverTab[142026]++
												out.write("\n")
												out.write(indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:416
		// _ = "end of CoverTab[142026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:417
		_go_fuzz_dep_.CoverTab[142027]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:417
		// _ = "end of CoverTab[142027]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:417
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:417
	// _ = "end of CoverTab[141890]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:417
	_go_fuzz_dep_.CoverTab[141891]++
											out.write("}")
											return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:419
	// _ = "end of CoverTab[141891]"
}

func (m *Marshaler) writeSep(out *errWriter) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:422
	_go_fuzz_dep_.CoverTab[142028]++
											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:423
		_go_fuzz_dep_.CoverTab[142029]++
												out.write(",\n")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:424
		// _ = "end of CoverTab[142029]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:425
		_go_fuzz_dep_.CoverTab[142030]++
												out.write(",")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:426
		// _ = "end of CoverTab[142030]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:427
	// _ = "end of CoverTab[142028]"
}

func (m *Marshaler) marshalAny(out *errWriter, any proto.Message, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:430
	_go_fuzz_dep_.CoverTab[142031]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:435
	v := reflect.ValueOf(any).Elem()
	turl := v.Field(0).String()
	val := v.Field(1).Bytes()

	var msg proto.Message
	var err error
	if m.AnyResolver != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:441
		_go_fuzz_dep_.CoverTab[142036]++
												msg, err = m.AnyResolver.Resolve(turl)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:442
		// _ = "end of CoverTab[142036]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:443
		_go_fuzz_dep_.CoverTab[142037]++
												msg, err = defaultResolveAny(turl)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:444
		// _ = "end of CoverTab[142037]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:445
	// _ = "end of CoverTab[142031]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:445
	_go_fuzz_dep_.CoverTab[142032]++
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:446
		_go_fuzz_dep_.CoverTab[142038]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:447
		// _ = "end of CoverTab[142038]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:448
		_go_fuzz_dep_.CoverTab[142039]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:448
		// _ = "end of CoverTab[142039]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:448
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:448
	// _ = "end of CoverTab[142032]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:448
	_go_fuzz_dep_.CoverTab[142033]++

											if err := proto.Unmarshal(val, msg); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:450
		_go_fuzz_dep_.CoverTab[142040]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:451
		// _ = "end of CoverTab[142040]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:452
		_go_fuzz_dep_.CoverTab[142041]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:452
		// _ = "end of CoverTab[142041]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:452
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:452
	// _ = "end of CoverTab[142033]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:452
	_go_fuzz_dep_.CoverTab[142034]++

											if _, ok := msg.(isWkt); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:454
		_go_fuzz_dep_.CoverTab[142042]++
												out.write("{")
												if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:456
			_go_fuzz_dep_.CoverTab[142048]++
													out.write("\n")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:457
			// _ = "end of CoverTab[142048]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:458
			_go_fuzz_dep_.CoverTab[142049]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:458
			// _ = "end of CoverTab[142049]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:458
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:458
		// _ = "end of CoverTab[142042]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:458
		_go_fuzz_dep_.CoverTab[142043]++
												if err := m.marshalTypeURL(out, indent, turl); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:459
			_go_fuzz_dep_.CoverTab[142050]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:460
			// _ = "end of CoverTab[142050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:461
			_go_fuzz_dep_.CoverTab[142051]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:461
			// _ = "end of CoverTab[142051]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:461
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:461
		// _ = "end of CoverTab[142043]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:461
		_go_fuzz_dep_.CoverTab[142044]++
												m.writeSep(out)
												if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:463
			_go_fuzz_dep_.CoverTab[142052]++
													out.write(indent)
													out.write(m.Indent)
													out.write(`"value": `)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:466
			// _ = "end of CoverTab[142052]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:467
			_go_fuzz_dep_.CoverTab[142053]++
													out.write(`"value":`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:468
			// _ = "end of CoverTab[142053]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:469
		// _ = "end of CoverTab[142044]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:469
		_go_fuzz_dep_.CoverTab[142045]++
												if err := m.marshalObject(out, msg, indent+m.Indent, ""); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:470
			_go_fuzz_dep_.CoverTab[142054]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:471
			// _ = "end of CoverTab[142054]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:472
			_go_fuzz_dep_.CoverTab[142055]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:472
			// _ = "end of CoverTab[142055]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:472
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:472
		// _ = "end of CoverTab[142045]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:472
		_go_fuzz_dep_.CoverTab[142046]++
												if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:473
			_go_fuzz_dep_.CoverTab[142056]++
													out.write("\n")
													out.write(indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:475
			// _ = "end of CoverTab[142056]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:476
			_go_fuzz_dep_.CoverTab[142057]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:476
			// _ = "end of CoverTab[142057]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:476
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:476
		// _ = "end of CoverTab[142046]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:476
		_go_fuzz_dep_.CoverTab[142047]++
												out.write("}")
												return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:478
		// _ = "end of CoverTab[142047]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:479
		_go_fuzz_dep_.CoverTab[142058]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:479
		// _ = "end of CoverTab[142058]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:479
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:479
	// _ = "end of CoverTab[142034]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:479
	_go_fuzz_dep_.CoverTab[142035]++

											return m.marshalObject(out, msg, indent, turl)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:481
	// _ = "end of CoverTab[142035]"
}

func (m *Marshaler) marshalTypeURL(out *errWriter, indent, typeURL string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:484
	_go_fuzz_dep_.CoverTab[142059]++
											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:485
		_go_fuzz_dep_.CoverTab[142063]++
												out.write(indent)
												out.write(m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:487
		// _ = "end of CoverTab[142063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:488
		_go_fuzz_dep_.CoverTab[142064]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:488
		// _ = "end of CoverTab[142064]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:488
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:488
	// _ = "end of CoverTab[142059]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:488
	_go_fuzz_dep_.CoverTab[142060]++
											out.write(`"@type":`)
											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:490
		_go_fuzz_dep_.CoverTab[142065]++
												out.write(" ")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:491
		// _ = "end of CoverTab[142065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:492
		_go_fuzz_dep_.CoverTab[142066]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:492
		// _ = "end of CoverTab[142066]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:492
	// _ = "end of CoverTab[142060]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:492
	_go_fuzz_dep_.CoverTab[142061]++
											b, err := json.Marshal(typeURL)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:494
		_go_fuzz_dep_.CoverTab[142067]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:495
		// _ = "end of CoverTab[142067]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:496
		_go_fuzz_dep_.CoverTab[142068]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:496
		// _ = "end of CoverTab[142068]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:496
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:496
	// _ = "end of CoverTab[142061]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:496
	_go_fuzz_dep_.CoverTab[142062]++
											out.write(string(b))
											return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:498
	// _ = "end of CoverTab[142062]"
}

// marshalField writes field description and value to the Writer.
func (m *Marshaler) marshalField(out *errWriter, prop *proto.Properties, v reflect.Value, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:502
	_go_fuzz_dep_.CoverTab[142069]++
											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:503
		_go_fuzz_dep_.CoverTab[142073]++
												out.write(indent)
												out.write(m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:505
		// _ = "end of CoverTab[142073]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:506
		_go_fuzz_dep_.CoverTab[142074]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:506
		// _ = "end of CoverTab[142074]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:506
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:506
	// _ = "end of CoverTab[142069]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:506
	_go_fuzz_dep_.CoverTab[142070]++
											out.write(`"`)
											out.write(prop.JSONName)
											out.write(`":`)
											if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:510
		_go_fuzz_dep_.CoverTab[142075]++
												out.write(" ")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:511
		// _ = "end of CoverTab[142075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:512
		_go_fuzz_dep_.CoverTab[142076]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:512
		// _ = "end of CoverTab[142076]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:512
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:512
	// _ = "end of CoverTab[142070]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:512
	_go_fuzz_dep_.CoverTab[142071]++
											if err := m.marshalValue(out, prop, v, indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:513
		_go_fuzz_dep_.CoverTab[142077]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:514
		// _ = "end of CoverTab[142077]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:515
		_go_fuzz_dep_.CoverTab[142078]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:515
		// _ = "end of CoverTab[142078]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:515
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:515
	// _ = "end of CoverTab[142071]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:515
	_go_fuzz_dep_.CoverTab[142072]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:516
	// _ = "end of CoverTab[142072]"
}

// marshalValue writes the value to the Writer.
func (m *Marshaler) marshalValue(out *errWriter, prop *proto.Properties, v reflect.Value, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:520
	_go_fuzz_dep_.CoverTab[142079]++

											v = reflect.Indirect(v)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:525
	if v.Kind() == reflect.Invalid {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:525
		_go_fuzz_dep_.CoverTab[142092]++
												out.write("null")
												return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:527
		// _ = "end of CoverTab[142092]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:528
		_go_fuzz_dep_.CoverTab[142093]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:528
		// _ = "end of CoverTab[142093]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:528
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:528
	// _ = "end of CoverTab[142079]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:528
	_go_fuzz_dep_.CoverTab[142080]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:531
	if v.Kind() == reflect.Slice && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:531
		_go_fuzz_dep_.CoverTab[142094]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:531
		return v.Type().Elem().Kind() != reflect.Uint8
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:531
		// _ = "end of CoverTab[142094]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:531
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:531
		_go_fuzz_dep_.CoverTab[142095]++
												out.write("[")
												comma := ""
												for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:534
			_go_fuzz_dep_.CoverTab[142098]++
													sliceVal := v.Index(i)
													out.write(comma)
													if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:537
				_go_fuzz_dep_.CoverTab[142101]++
														out.write("\n")
														out.write(indent)
														out.write(m.Indent)
														out.write(m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:541
				// _ = "end of CoverTab[142101]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:542
				_go_fuzz_dep_.CoverTab[142102]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:542
				// _ = "end of CoverTab[142102]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:542
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:542
			// _ = "end of CoverTab[142098]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:542
			_go_fuzz_dep_.CoverTab[142099]++
													if err := m.marshalValue(out, prop, sliceVal, indent+m.Indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:543
				_go_fuzz_dep_.CoverTab[142103]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:544
				// _ = "end of CoverTab[142103]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:545
				_go_fuzz_dep_.CoverTab[142104]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:545
				// _ = "end of CoverTab[142104]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:545
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:545
			// _ = "end of CoverTab[142099]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:545
			_go_fuzz_dep_.CoverTab[142100]++
													comma = ","
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:546
			// _ = "end of CoverTab[142100]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:547
		// _ = "end of CoverTab[142095]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:547
		_go_fuzz_dep_.CoverTab[142096]++
												if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:548
			_go_fuzz_dep_.CoverTab[142105]++
													out.write("\n")
													out.write(indent)
													out.write(m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:551
			// _ = "end of CoverTab[142105]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:552
			_go_fuzz_dep_.CoverTab[142106]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:552
			// _ = "end of CoverTab[142106]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:552
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:552
		// _ = "end of CoverTab[142096]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:552
		_go_fuzz_dep_.CoverTab[142097]++
												out.write("]")
												return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:554
		// _ = "end of CoverTab[142097]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:555
		_go_fuzz_dep_.CoverTab[142107]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:555
		// _ = "end of CoverTab[142107]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:555
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:555
	// _ = "end of CoverTab[142080]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:555
	_go_fuzz_dep_.CoverTab[142081]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:559
	if v.Type().Implements(wktType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:559
		_go_fuzz_dep_.CoverTab[142108]++
												wkt := v.Interface().(isWkt)
												switch wkt.XXX_WellKnownType() {
		case "NullValue":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:562
			_go_fuzz_dep_.CoverTab[142109]++
													out.write("null")
													return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:564
			// _ = "end of CoverTab[142109]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:564
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:564
			_go_fuzz_dep_.CoverTab[142110]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:564
			// _ = "end of CoverTab[142110]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:565
		// _ = "end of CoverTab[142108]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:566
		_go_fuzz_dep_.CoverTab[142111]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:566
		// _ = "end of CoverTab[142111]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:566
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:566
	// _ = "end of CoverTab[142081]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:566
	_go_fuzz_dep_.CoverTab[142082]++

											if t, ok := v.Interface().(time.Time); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:568
		_go_fuzz_dep_.CoverTab[142112]++
												ts, err := types.TimestampProto(t)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:570
			_go_fuzz_dep_.CoverTab[142114]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:571
			// _ = "end of CoverTab[142114]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:572
			_go_fuzz_dep_.CoverTab[142115]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:572
			// _ = "end of CoverTab[142115]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:572
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:572
		// _ = "end of CoverTab[142112]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:572
		_go_fuzz_dep_.CoverTab[142113]++
												return m.marshalValue(out, prop, reflect.ValueOf(ts), indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:573
		// _ = "end of CoverTab[142113]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:574
		_go_fuzz_dep_.CoverTab[142116]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:574
		// _ = "end of CoverTab[142116]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:574
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:574
	// _ = "end of CoverTab[142082]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:574
	_go_fuzz_dep_.CoverTab[142083]++

											if d, ok := v.Interface().(time.Duration); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:576
		_go_fuzz_dep_.CoverTab[142117]++
												dur := types.DurationProto(d)
												return m.marshalValue(out, prop, reflect.ValueOf(dur), indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:578
		// _ = "end of CoverTab[142117]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:579
		_go_fuzz_dep_.CoverTab[142118]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:579
		// _ = "end of CoverTab[142118]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:579
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:579
	// _ = "end of CoverTab[142083]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:579
	_go_fuzz_dep_.CoverTab[142084]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:582
	if !m.EnumsAsInts && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:582
		_go_fuzz_dep_.CoverTab[142119]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:582
		return prop.Enum != ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:582
		// _ = "end of CoverTab[142119]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:582
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:582
		_go_fuzz_dep_.CoverTab[142120]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:586
		enumStr := v.Interface().(fmt.Stringer).String()
		var valStr string
		if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:588
			_go_fuzz_dep_.CoverTab[142125]++
													valStr = strconv.Itoa(int(v.Elem().Int()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:589
			// _ = "end of CoverTab[142125]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:590
			_go_fuzz_dep_.CoverTab[142126]++
													valStr = strconv.Itoa(int(v.Int()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:591
			// _ = "end of CoverTab[142126]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:592
		// _ = "end of CoverTab[142120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:592
		_go_fuzz_dep_.CoverTab[142121]++

												if m, ok := v.Interface().(interface {
			MarshalJSON() ([]byte, error)
		}); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:596
			_go_fuzz_dep_.CoverTab[142127]++
													data, err := m.MarshalJSON()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:598
				_go_fuzz_dep_.CoverTab[142129]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:599
				// _ = "end of CoverTab[142129]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:600
				_go_fuzz_dep_.CoverTab[142130]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:600
				// _ = "end of CoverTab[142130]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:600
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:600
			// _ = "end of CoverTab[142127]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:600
			_go_fuzz_dep_.CoverTab[142128]++
													enumStr = string(data)
													enumStr, err = strconv.Unquote(enumStr)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:603
				_go_fuzz_dep_.CoverTab[142131]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:604
				// _ = "end of CoverTab[142131]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:605
				_go_fuzz_dep_.CoverTab[142132]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:605
				// _ = "end of CoverTab[142132]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:605
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:605
			// _ = "end of CoverTab[142128]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:606
			_go_fuzz_dep_.CoverTab[142133]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:606
			// _ = "end of CoverTab[142133]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:606
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:606
		// _ = "end of CoverTab[142121]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:606
		_go_fuzz_dep_.CoverTab[142122]++

												isKnownEnum := enumStr != valStr

												if isKnownEnum {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:610
			_go_fuzz_dep_.CoverTab[142134]++
													out.write(`"`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:611
			// _ = "end of CoverTab[142134]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:612
			_go_fuzz_dep_.CoverTab[142135]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:612
			// _ = "end of CoverTab[142135]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:612
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:612
		// _ = "end of CoverTab[142122]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:612
		_go_fuzz_dep_.CoverTab[142123]++
												out.write(enumStr)
												if isKnownEnum {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:614
			_go_fuzz_dep_.CoverTab[142136]++
													out.write(`"`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:615
			// _ = "end of CoverTab[142136]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:616
			_go_fuzz_dep_.CoverTab[142137]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:616
			// _ = "end of CoverTab[142137]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:616
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:616
		// _ = "end of CoverTab[142123]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:616
		_go_fuzz_dep_.CoverTab[142124]++
												return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:617
		// _ = "end of CoverTab[142124]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:618
		_go_fuzz_dep_.CoverTab[142138]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:618
		// _ = "end of CoverTab[142138]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:618
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:618
	// _ = "end of CoverTab[142084]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:618
	_go_fuzz_dep_.CoverTab[142085]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:621
	if v.Kind() == reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:621
		_go_fuzz_dep_.CoverTab[142139]++
												i := v
												if v.CanAddr() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:623
			_go_fuzz_dep_.CoverTab[142144]++
													i = v.Addr()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:624
			// _ = "end of CoverTab[142144]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:625
			_go_fuzz_dep_.CoverTab[142145]++
													i = reflect.New(v.Type())
													i.Elem().Set(v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:627
			// _ = "end of CoverTab[142145]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:628
		// _ = "end of CoverTab[142139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:628
		_go_fuzz_dep_.CoverTab[142140]++
												iface := i.Interface()
												if iface == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:630
			_go_fuzz_dep_.CoverTab[142146]++
													out.write(`null`)
													return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:632
			// _ = "end of CoverTab[142146]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:633
			_go_fuzz_dep_.CoverTab[142147]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:633
			// _ = "end of CoverTab[142147]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:633
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:633
		// _ = "end of CoverTab[142140]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:633
		_go_fuzz_dep_.CoverTab[142141]++

												if m, ok := v.Interface().(interface {
			MarshalJSON() ([]byte, error)
		}); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:637
			_go_fuzz_dep_.CoverTab[142148]++
													data, err := m.MarshalJSON()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:639
				_go_fuzz_dep_.CoverTab[142150]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:640
				// _ = "end of CoverTab[142150]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:641
				_go_fuzz_dep_.CoverTab[142151]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:641
				// _ = "end of CoverTab[142151]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:641
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:641
			// _ = "end of CoverTab[142148]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:641
			_go_fuzz_dep_.CoverTab[142149]++
													out.write(string(data))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:643
			// _ = "end of CoverTab[142149]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:644
			_go_fuzz_dep_.CoverTab[142152]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:644
			// _ = "end of CoverTab[142152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:644
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:644
		// _ = "end of CoverTab[142141]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:644
		_go_fuzz_dep_.CoverTab[142142]++

												pm, ok := iface.(proto.Message)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:647
			_go_fuzz_dep_.CoverTab[142153]++
													if prop.CustomType == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:648
				_go_fuzz_dep_.CoverTab[142156]++
														return fmt.Errorf("%v does not implement proto.Message", v.Type())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:649
				// _ = "end of CoverTab[142156]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:650
				_go_fuzz_dep_.CoverTab[142157]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:650
				// _ = "end of CoverTab[142157]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:650
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:650
			// _ = "end of CoverTab[142153]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:650
			_go_fuzz_dep_.CoverTab[142154]++
													t := proto.MessageType(prop.CustomType)
													if t == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:652
				_go_fuzz_dep_.CoverTab[142158]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:652
				return !i.Type().ConvertibleTo(t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:652
				// _ = "end of CoverTab[142158]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:652
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:652
				_go_fuzz_dep_.CoverTab[142159]++
														return fmt.Errorf("%v declared custom type %s but it is not convertible to %v", v.Type(), prop.CustomType, t)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:653
				// _ = "end of CoverTab[142159]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:654
				_go_fuzz_dep_.CoverTab[142160]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:654
				// _ = "end of CoverTab[142160]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:654
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:654
			// _ = "end of CoverTab[142154]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:654
			_go_fuzz_dep_.CoverTab[142155]++
													pm = i.Convert(t).Interface().(proto.Message)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:655
			// _ = "end of CoverTab[142155]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:656
			_go_fuzz_dep_.CoverTab[142161]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:656
			// _ = "end of CoverTab[142161]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:656
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:656
		// _ = "end of CoverTab[142142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:656
		_go_fuzz_dep_.CoverTab[142143]++
												return m.marshalObject(out, pm, indent+m.Indent, "")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:657
		// _ = "end of CoverTab[142143]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:658
		_go_fuzz_dep_.CoverTab[142162]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:658
		// _ = "end of CoverTab[142162]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:658
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:658
	// _ = "end of CoverTab[142085]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:658
	_go_fuzz_dep_.CoverTab[142086]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:662
	if v.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:662
		_go_fuzz_dep_.CoverTab[142163]++
												out.write(`{`)
												keys := v.MapKeys()
												sort.Sort(mapKeys(keys))
												for i, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:666
			_go_fuzz_dep_.CoverTab[142166]++
													if i > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:667
				_go_fuzz_dep_.CoverTab[142173]++
														out.write(`,`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:668
				// _ = "end of CoverTab[142173]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:669
				_go_fuzz_dep_.CoverTab[142174]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:669
				// _ = "end of CoverTab[142174]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:669
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:669
			// _ = "end of CoverTab[142166]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:669
			_go_fuzz_dep_.CoverTab[142167]++
													if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:670
				_go_fuzz_dep_.CoverTab[142175]++
														out.write("\n")
														out.write(indent)
														out.write(m.Indent)
														out.write(m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:674
				// _ = "end of CoverTab[142175]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:675
				_go_fuzz_dep_.CoverTab[142176]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:675
				// _ = "end of CoverTab[142176]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:675
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:675
			// _ = "end of CoverTab[142167]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:675
			_go_fuzz_dep_.CoverTab[142168]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:678
			b, err := json.Marshal(k.Interface())
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:679
				_go_fuzz_dep_.CoverTab[142177]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:680
				// _ = "end of CoverTab[142177]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:681
				_go_fuzz_dep_.CoverTab[142178]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:681
				// _ = "end of CoverTab[142178]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:681
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:681
			// _ = "end of CoverTab[142168]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:681
			_go_fuzz_dep_.CoverTab[142169]++
													s := string(b)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:685
			if !strings.HasPrefix(s, `"`) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:685
				_go_fuzz_dep_.CoverTab[142179]++
														b, err := json.Marshal(s)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:687
					_go_fuzz_dep_.CoverTab[142181]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:688
					// _ = "end of CoverTab[142181]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:689
					_go_fuzz_dep_.CoverTab[142182]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:689
					// _ = "end of CoverTab[142182]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:689
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:689
				// _ = "end of CoverTab[142179]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:689
				_go_fuzz_dep_.CoverTab[142180]++
														s = string(b)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:690
				// _ = "end of CoverTab[142180]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:691
				_go_fuzz_dep_.CoverTab[142183]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:691
				// _ = "end of CoverTab[142183]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:691
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:691
			// _ = "end of CoverTab[142169]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:691
			_go_fuzz_dep_.CoverTab[142170]++

													out.write(s)
													out.write(`:`)
													if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:695
				_go_fuzz_dep_.CoverTab[142184]++
														out.write(` `)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:696
				// _ = "end of CoverTab[142184]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:697
				_go_fuzz_dep_.CoverTab[142185]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:697
				// _ = "end of CoverTab[142185]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:697
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:697
			// _ = "end of CoverTab[142170]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:697
			_go_fuzz_dep_.CoverTab[142171]++

													vprop := prop
													if prop != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:700
				_go_fuzz_dep_.CoverTab[142186]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:700
				return prop.MapValProp != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:700
				// _ = "end of CoverTab[142186]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:700
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:700
				_go_fuzz_dep_.CoverTab[142187]++
														vprop = prop.MapValProp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:701
				// _ = "end of CoverTab[142187]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:702
				_go_fuzz_dep_.CoverTab[142188]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:702
				// _ = "end of CoverTab[142188]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:702
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:702
			// _ = "end of CoverTab[142171]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:702
			_go_fuzz_dep_.CoverTab[142172]++
													if err := m.marshalValue(out, vprop, v.MapIndex(k), indent+m.Indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:703
				_go_fuzz_dep_.CoverTab[142189]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:704
				// _ = "end of CoverTab[142189]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:705
				_go_fuzz_dep_.CoverTab[142190]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:705
				// _ = "end of CoverTab[142190]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:705
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:705
			// _ = "end of CoverTab[142172]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:706
		// _ = "end of CoverTab[142163]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:706
		_go_fuzz_dep_.CoverTab[142164]++
												if m.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:707
			_go_fuzz_dep_.CoverTab[142191]++
													out.write("\n")
													out.write(indent)
													out.write(m.Indent)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:710
			// _ = "end of CoverTab[142191]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:711
			_go_fuzz_dep_.CoverTab[142192]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:711
			// _ = "end of CoverTab[142192]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:711
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:711
		// _ = "end of CoverTab[142164]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:711
		_go_fuzz_dep_.CoverTab[142165]++
												out.write(`}`)
												return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:713
		// _ = "end of CoverTab[142165]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:714
		_go_fuzz_dep_.CoverTab[142193]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:714
		// _ = "end of CoverTab[142193]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:714
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:714
	// _ = "end of CoverTab[142086]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:714
	_go_fuzz_dep_.CoverTab[142087]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:717
	if v.Kind() == reflect.Float32 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:717
		_go_fuzz_dep_.CoverTab[142194]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:717
		return v.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:717
		// _ = "end of CoverTab[142194]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:717
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:717
		_go_fuzz_dep_.CoverTab[142195]++
												f := v.Float()
												var sval string
												switch {
		case math.IsInf(f, 1):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:721
			_go_fuzz_dep_.CoverTab[142197]++
													sval = `"Infinity"`
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:722
			// _ = "end of CoverTab[142197]"
		case math.IsInf(f, -1):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:723
			_go_fuzz_dep_.CoverTab[142198]++
													sval = `"-Infinity"`
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:724
			// _ = "end of CoverTab[142198]"
		case math.IsNaN(f):
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:725
			_go_fuzz_dep_.CoverTab[142199]++
													sval = `"NaN"`
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:726
			// _ = "end of CoverTab[142199]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:726
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:726
			_go_fuzz_dep_.CoverTab[142200]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:726
			// _ = "end of CoverTab[142200]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:727
		// _ = "end of CoverTab[142195]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:727
		_go_fuzz_dep_.CoverTab[142196]++
												if sval != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:728
			_go_fuzz_dep_.CoverTab[142201]++
													out.write(sval)
													return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:730
			// _ = "end of CoverTab[142201]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:731
			_go_fuzz_dep_.CoverTab[142202]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:731
			// _ = "end of CoverTab[142202]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:731
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:731
		// _ = "end of CoverTab[142196]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:732
		_go_fuzz_dep_.CoverTab[142203]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:732
		// _ = "end of CoverTab[142203]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:732
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:732
	// _ = "end of CoverTab[142087]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:732
	_go_fuzz_dep_.CoverTab[142088]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:735
	b, err := json.Marshal(v.Interface())
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:736
		_go_fuzz_dep_.CoverTab[142204]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:737
		// _ = "end of CoverTab[142204]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:738
		_go_fuzz_dep_.CoverTab[142205]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:738
		// _ = "end of CoverTab[142205]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:738
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:738
	// _ = "end of CoverTab[142088]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:738
	_go_fuzz_dep_.CoverTab[142089]++
											needToQuote := string(b[0]) != `"` && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
		_go_fuzz_dep_.CoverTab[142206]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
		return (v.Kind() == reflect.Int64 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
			_go_fuzz_dep_.CoverTab[142207]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
			return v.Kind() == reflect.Uint64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
			// _ = "end of CoverTab[142207]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
		}())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
		// _ = "end of CoverTab[142206]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:739
	}()
											if needToQuote {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:740
		_go_fuzz_dep_.CoverTab[142208]++
												out.write(`"`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:741
		// _ = "end of CoverTab[142208]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:742
		_go_fuzz_dep_.CoverTab[142209]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:742
		// _ = "end of CoverTab[142209]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:742
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:742
	// _ = "end of CoverTab[142089]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:742
	_go_fuzz_dep_.CoverTab[142090]++
											out.write(string(b))
											if needToQuote {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:744
		_go_fuzz_dep_.CoverTab[142210]++
												out.write(`"`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:745
		// _ = "end of CoverTab[142210]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:746
		_go_fuzz_dep_.CoverTab[142211]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:746
		// _ = "end of CoverTab[142211]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:746
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:746
	// _ = "end of CoverTab[142090]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:746
	_go_fuzz_dep_.CoverTab[142091]++
											return out.err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:747
	// _ = "end of CoverTab[142091]"
}

// Unmarshaler is a configurable object for converting from a JSON
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:750
// representation to a protocol buffer object.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:752
type Unmarshaler struct {
	// Whether to allow messages to contain unknown fields, as opposed to
	// failing to unmarshal.
	AllowUnknownFields	bool

	// A custom URL resolver to use when unmarshaling Any messages from JSON.
	// If unset, the default resolution strategy is to extract the
	// fully-qualified type name from the type URL and pass that to
	// proto.MessageType(string).
	AnyResolver	AnyResolver
}

// UnmarshalNext unmarshals the next protocol buffer from a JSON object stream.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:764
// This function is lenient and will decode any options permutations of the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:764
// related Marshaler.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:767
func (u *Unmarshaler) UnmarshalNext(dec *json.Decoder, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:767
	_go_fuzz_dep_.CoverTab[142212]++
											inputValue := json.RawMessage{}
											if err := dec.Decode(&inputValue); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:769
		_go_fuzz_dep_.CoverTab[142215]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:770
		// _ = "end of CoverTab[142215]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:771
		_go_fuzz_dep_.CoverTab[142216]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:771
		// _ = "end of CoverTab[142216]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:771
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:771
	// _ = "end of CoverTab[142212]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:771
	_go_fuzz_dep_.CoverTab[142213]++
											if err := u.unmarshalValue(reflect.ValueOf(pb).Elem(), inputValue, nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:772
		_go_fuzz_dep_.CoverTab[142217]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:773
		// _ = "end of CoverTab[142217]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:774
		_go_fuzz_dep_.CoverTab[142218]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:774
		// _ = "end of CoverTab[142218]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:774
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:774
	// _ = "end of CoverTab[142213]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:774
	_go_fuzz_dep_.CoverTab[142214]++
											return checkRequiredFields(pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:775
	// _ = "end of CoverTab[142214]"
}

// Unmarshal unmarshals a JSON object stream into a protocol
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:778
// buffer. This function is lenient and will decode any options
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:778
// permutations of the related Marshaler.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:781
func (u *Unmarshaler) Unmarshal(r io.Reader, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:781
	_go_fuzz_dep_.CoverTab[142219]++
											dec := json.NewDecoder(r)
											return u.UnmarshalNext(dec, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:783
	// _ = "end of CoverTab[142219]"
}

// UnmarshalNext unmarshals the next protocol buffer from a JSON object stream.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:786
// This function is lenient and will decode any options permutations of the
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:786
// related Marshaler.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:789
func UnmarshalNext(dec *json.Decoder, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:789
	_go_fuzz_dep_.CoverTab[142220]++
											return new(Unmarshaler).UnmarshalNext(dec, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:790
	// _ = "end of CoverTab[142220]"
}

// Unmarshal unmarshals a JSON object stream into a protocol
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:793
// buffer. This function is lenient and will decode any options
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:793
// permutations of the related Marshaler.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:796
func Unmarshal(r io.Reader, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:796
	_go_fuzz_dep_.CoverTab[142221]++
											return new(Unmarshaler).Unmarshal(r, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:797
	// _ = "end of CoverTab[142221]"
}

// UnmarshalString will populate the fields of a protocol buffer based
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:800
// on a JSON string. This function is lenient and will decode any options
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:800
// permutations of the related Marshaler.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:803
func UnmarshalString(str string, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:803
	_go_fuzz_dep_.CoverTab[142222]++
											return new(Unmarshaler).Unmarshal(strings.NewReader(str), pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:804
	// _ = "end of CoverTab[142222]"
}

// unmarshalValue converts/copies a value into the target.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:807
// prop may be nil.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:809
func (u *Unmarshaler) unmarshalValue(target reflect.Value, inputValue json.RawMessage, prop *proto.Properties) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:809
	_go_fuzz_dep_.CoverTab[142223]++
											targetType := target.Type()

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:813
	if targetType.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:813
		_go_fuzz_dep_.CoverTab[142236]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:816
		_, isJSONPBUnmarshaler := target.Interface().(JSONPBUnmarshaler)
		if string(inputValue) == "null" && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			_go_fuzz_dep_.CoverTab[142238]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			return targetType != reflect.TypeOf(&types.Value{})
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			// _ = "end of CoverTab[142238]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			_go_fuzz_dep_.CoverTab[142239]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			return !isJSONPBUnmarshaler
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			// _ = "end of CoverTab[142239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:817
			_go_fuzz_dep_.CoverTab[142240]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:818
			// _ = "end of CoverTab[142240]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:819
			_go_fuzz_dep_.CoverTab[142241]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:819
			// _ = "end of CoverTab[142241]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:819
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:819
		// _ = "end of CoverTab[142236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:819
		_go_fuzz_dep_.CoverTab[142237]++
												target.Set(reflect.New(targetType.Elem()))

												return u.unmarshalValue(target.Elem(), inputValue, prop)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:822
		// _ = "end of CoverTab[142237]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:823
		_go_fuzz_dep_.CoverTab[142242]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:823
		// _ = "end of CoverTab[142242]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:823
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:823
	// _ = "end of CoverTab[142223]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:823
	_go_fuzz_dep_.CoverTab[142224]++

											if jsu, ok := target.Addr().Interface().(JSONPBUnmarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:825
		_go_fuzz_dep_.CoverTab[142243]++
												return jsu.UnmarshalJSONPB(u, []byte(inputValue))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:826
		// _ = "end of CoverTab[142243]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:827
		_go_fuzz_dep_.CoverTab[142244]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:827
		// _ = "end of CoverTab[142244]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:827
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:827
	// _ = "end of CoverTab[142224]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:827
	_go_fuzz_dep_.CoverTab[142225]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:830
	if w, ok := target.Addr().Interface().(isWkt); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:830
		_go_fuzz_dep_.CoverTab[142245]++
												switch w.XXX_WellKnownType() {
		case "DoubleValue", "FloatValue", "Int64Value", "UInt64Value",
			"Int32Value", "UInt32Value", "BoolValue", "StringValue", "BytesValue":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:833
			_go_fuzz_dep_.CoverTab[142246]++
													return u.unmarshalValue(target.Field(0), inputValue, prop)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:834
			// _ = "end of CoverTab[142246]"
		case "Any":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:835
			_go_fuzz_dep_.CoverTab[142247]++
			// Use json.RawMessage pointer type instead of value to support pre-1.8 version.
			// 1.8 changed RawMessage.MarshalJSON from pointer type to value type, see
			// https://github.com/golang/go/issues/14493
			var jsonFields map[string]*json.RawMessage
			if err := json.Unmarshal(inputValue, &jsonFields); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:840
				_go_fuzz_dep_.CoverTab[142270]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:841
				// _ = "end of CoverTab[142270]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:842
				_go_fuzz_dep_.CoverTab[142271]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:842
				// _ = "end of CoverTab[142271]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:842
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:842
			// _ = "end of CoverTab[142247]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:842
			_go_fuzz_dep_.CoverTab[142248]++

													val, ok := jsonFields["@type"]
													if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:845
				_go_fuzz_dep_.CoverTab[142272]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:845
				return val == nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:845
				// _ = "end of CoverTab[142272]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:845
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:845
				_go_fuzz_dep_.CoverTab[142273]++
														return errors.New("Any JSON doesn't have '@type'")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:846
				// _ = "end of CoverTab[142273]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:847
				_go_fuzz_dep_.CoverTab[142274]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:847
				// _ = "end of CoverTab[142274]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:847
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:847
			// _ = "end of CoverTab[142248]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:847
			_go_fuzz_dep_.CoverTab[142249]++

													var turl string
													if err := json.Unmarshal([]byte(*val), &turl); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:850
				_go_fuzz_dep_.CoverTab[142275]++
														return fmt.Errorf("can't unmarshal Any's '@type': %q", *val)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:851
				// _ = "end of CoverTab[142275]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:852
				_go_fuzz_dep_.CoverTab[142276]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:852
				// _ = "end of CoverTab[142276]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:852
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:852
			// _ = "end of CoverTab[142249]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:852
			_go_fuzz_dep_.CoverTab[142250]++
													target.Field(0).SetString(turl)

													var m proto.Message
													var err error
													if u.AnyResolver != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:857
				_go_fuzz_dep_.CoverTab[142277]++
														m, err = u.AnyResolver.Resolve(turl)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:858
				// _ = "end of CoverTab[142277]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:859
				_go_fuzz_dep_.CoverTab[142278]++
														m, err = defaultResolveAny(turl)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:860
				// _ = "end of CoverTab[142278]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:861
			// _ = "end of CoverTab[142250]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:861
			_go_fuzz_dep_.CoverTab[142251]++
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:862
				_go_fuzz_dep_.CoverTab[142279]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:863
				// _ = "end of CoverTab[142279]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:864
				_go_fuzz_dep_.CoverTab[142280]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:864
				// _ = "end of CoverTab[142280]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:864
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:864
			// _ = "end of CoverTab[142251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:864
			_go_fuzz_dep_.CoverTab[142252]++

													if _, ok := m.(isWkt); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:866
				_go_fuzz_dep_.CoverTab[142281]++
														val, ok := jsonFields["value"]
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:868
					_go_fuzz_dep_.CoverTab[142283]++
															return errors.New("Any JSON doesn't have 'value'")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:869
					// _ = "end of CoverTab[142283]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:870
					_go_fuzz_dep_.CoverTab[142284]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:870
					// _ = "end of CoverTab[142284]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:870
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:870
				// _ = "end of CoverTab[142281]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:870
				_go_fuzz_dep_.CoverTab[142282]++

														if err = u.unmarshalValue(reflect.ValueOf(m).Elem(), *val, nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:872
					_go_fuzz_dep_.CoverTab[142285]++
															return fmt.Errorf("can't unmarshal Any nested proto %T: %v", m, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:873
					// _ = "end of CoverTab[142285]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:874
					_go_fuzz_dep_.CoverTab[142286]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:874
					// _ = "end of CoverTab[142286]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:874
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:874
				// _ = "end of CoverTab[142282]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:875
				_go_fuzz_dep_.CoverTab[142287]++
														delete(jsonFields, "@type")
														nestedProto, uerr := json.Marshal(jsonFields)
														if uerr != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:878
					_go_fuzz_dep_.CoverTab[142289]++
															return fmt.Errorf("can't generate JSON for Any's nested proto to be unmarshaled: %v", uerr)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:879
					// _ = "end of CoverTab[142289]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:880
					_go_fuzz_dep_.CoverTab[142290]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:880
					// _ = "end of CoverTab[142290]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:880
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:880
				// _ = "end of CoverTab[142287]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:880
				_go_fuzz_dep_.CoverTab[142288]++

														if err = u.unmarshalValue(reflect.ValueOf(m).Elem(), nestedProto, nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:882
					_go_fuzz_dep_.CoverTab[142291]++
															return fmt.Errorf("can't unmarshal Any nested proto %T: %v", m, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:883
					// _ = "end of CoverTab[142291]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:884
					_go_fuzz_dep_.CoverTab[142292]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:884
					// _ = "end of CoverTab[142292]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:884
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:884
				// _ = "end of CoverTab[142288]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:885
			// _ = "end of CoverTab[142252]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:885
			_go_fuzz_dep_.CoverTab[142253]++

													b, err := proto.Marshal(m)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:888
				_go_fuzz_dep_.CoverTab[142293]++
														return fmt.Errorf("can't marshal proto %T into Any.Value: %v", m, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:889
				// _ = "end of CoverTab[142293]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:890
				_go_fuzz_dep_.CoverTab[142294]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:890
				// _ = "end of CoverTab[142294]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:890
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:890
			// _ = "end of CoverTab[142253]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:890
			_go_fuzz_dep_.CoverTab[142254]++
													target.Field(1).SetBytes(b)

													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:893
			// _ = "end of CoverTab[142254]"
		case "Duration":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:894
			_go_fuzz_dep_.CoverTab[142255]++
													unq, err := unquote(string(inputValue))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:896
				_go_fuzz_dep_.CoverTab[142295]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:897
				// _ = "end of CoverTab[142295]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:898
				_go_fuzz_dep_.CoverTab[142296]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:898
				// _ = "end of CoverTab[142296]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:898
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:898
			// _ = "end of CoverTab[142255]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:898
			_go_fuzz_dep_.CoverTab[142256]++

													d, err := time.ParseDuration(unq)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:901
				_go_fuzz_dep_.CoverTab[142297]++
														return fmt.Errorf("bad Duration: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:902
				// _ = "end of CoverTab[142297]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:903
				_go_fuzz_dep_.CoverTab[142298]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:903
				// _ = "end of CoverTab[142298]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:903
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:903
			// _ = "end of CoverTab[142256]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:903
			_go_fuzz_dep_.CoverTab[142257]++

													ns := d.Nanoseconds()
													s := ns / 1e9
													ns %= 1e9
													target.Field(0).SetInt(s)
													target.Field(1).SetInt(ns)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:910
			// _ = "end of CoverTab[142257]"
		case "Timestamp":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:911
			_go_fuzz_dep_.CoverTab[142258]++
													unq, err := unquote(string(inputValue))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:913
				_go_fuzz_dep_.CoverTab[142299]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:914
				// _ = "end of CoverTab[142299]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:915
				_go_fuzz_dep_.CoverTab[142300]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:915
				// _ = "end of CoverTab[142300]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:915
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:915
			// _ = "end of CoverTab[142258]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:915
			_go_fuzz_dep_.CoverTab[142259]++

													t, err := time.Parse(time.RFC3339Nano, unq)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:918
				_go_fuzz_dep_.CoverTab[142301]++
														return fmt.Errorf("bad Timestamp: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:919
				// _ = "end of CoverTab[142301]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:920
				_go_fuzz_dep_.CoverTab[142302]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:920
				// _ = "end of CoverTab[142302]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:920
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:920
			// _ = "end of CoverTab[142259]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:920
			_go_fuzz_dep_.CoverTab[142260]++

													target.Field(0).SetInt(t.Unix())
													target.Field(1).SetInt(int64(t.Nanosecond()))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:924
			// _ = "end of CoverTab[142260]"
		case "Struct":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:925
			_go_fuzz_dep_.CoverTab[142261]++
													var m map[string]json.RawMessage
													if err := json.Unmarshal(inputValue, &m); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:927
				_go_fuzz_dep_.CoverTab[142303]++
														return fmt.Errorf("bad StructValue: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:928
				// _ = "end of CoverTab[142303]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:929
				_go_fuzz_dep_.CoverTab[142304]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:929
				// _ = "end of CoverTab[142304]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:929
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:929
			// _ = "end of CoverTab[142261]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:929
			_go_fuzz_dep_.CoverTab[142262]++
													target.Field(0).Set(reflect.ValueOf(map[string]*types.Value{}))
													for k, jv := range m {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:931
				_go_fuzz_dep_.CoverTab[142305]++
														pv := &types.Value{}
														if err := u.unmarshalValue(reflect.ValueOf(pv).Elem(), jv, prop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:933
					_go_fuzz_dep_.CoverTab[142307]++
															return fmt.Errorf("bad value in StructValue for key %q: %v", k, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:934
					// _ = "end of CoverTab[142307]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:935
					_go_fuzz_dep_.CoverTab[142308]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:935
					// _ = "end of CoverTab[142308]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:935
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:935
				// _ = "end of CoverTab[142305]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:935
				_go_fuzz_dep_.CoverTab[142306]++
														target.Field(0).SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(pv))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:936
				// _ = "end of CoverTab[142306]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:937
			// _ = "end of CoverTab[142262]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:937
			_go_fuzz_dep_.CoverTab[142263]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:938
			// _ = "end of CoverTab[142263]"
		case "ListValue":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:939
			_go_fuzz_dep_.CoverTab[142264]++
													var s []json.RawMessage
													if err := json.Unmarshal(inputValue, &s); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:941
				_go_fuzz_dep_.CoverTab[142309]++
														return fmt.Errorf("bad ListValue: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:942
				// _ = "end of CoverTab[142309]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:943
				_go_fuzz_dep_.CoverTab[142310]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:943
				// _ = "end of CoverTab[142310]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:943
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:943
			// _ = "end of CoverTab[142264]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:943
			_go_fuzz_dep_.CoverTab[142265]++

													target.Field(0).Set(reflect.ValueOf(make([]*types.Value, len(s))))
													for i, sv := range s {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:946
				_go_fuzz_dep_.CoverTab[142311]++
														if err := u.unmarshalValue(target.Field(0).Index(i), sv, prop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:947
					_go_fuzz_dep_.CoverTab[142312]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:948
					// _ = "end of CoverTab[142312]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:949
					_go_fuzz_dep_.CoverTab[142313]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:949
					// _ = "end of CoverTab[142313]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:949
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:949
				// _ = "end of CoverTab[142311]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:950
			// _ = "end of CoverTab[142265]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:950
			_go_fuzz_dep_.CoverTab[142266]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:951
			// _ = "end of CoverTab[142266]"
		case "Value":
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:952
			_go_fuzz_dep_.CoverTab[142267]++
													ivStr := string(inputValue)
													if ivStr == "null" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:954
				_go_fuzz_dep_.CoverTab[142314]++
														target.Field(0).Set(reflect.ValueOf(&types.Value_NullValue{}))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:955
				// _ = "end of CoverTab[142314]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:956
				_go_fuzz_dep_.CoverTab[142315]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:956
				if v, err := strconv.ParseFloat(ivStr, 0); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:956
					_go_fuzz_dep_.CoverTab[142316]++
															target.Field(0).Set(reflect.ValueOf(&types.Value_NumberValue{NumberValue: v}))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:957
					// _ = "end of CoverTab[142316]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:958
					_go_fuzz_dep_.CoverTab[142317]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:958
					if v, err := unquote(ivStr); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:958
						_go_fuzz_dep_.CoverTab[142318]++
																target.Field(0).Set(reflect.ValueOf(&types.Value_StringValue{StringValue: v}))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:959
						// _ = "end of CoverTab[142318]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:960
						_go_fuzz_dep_.CoverTab[142319]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:960
						if v, err := strconv.ParseBool(ivStr); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:960
							_go_fuzz_dep_.CoverTab[142320]++
																	target.Field(0).Set(reflect.ValueOf(&types.Value_BoolValue{BoolValue: v}))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:961
							// _ = "end of CoverTab[142320]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:962
							_go_fuzz_dep_.CoverTab[142321]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:962
							if err := json.Unmarshal(inputValue, &[]json.RawMessage{}); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:962
								_go_fuzz_dep_.CoverTab[142322]++
																		lv := &types.ListValue{}
																		target.Field(0).Set(reflect.ValueOf(&types.Value_ListValue{ListValue: lv}))
																		return u.unmarshalValue(reflect.ValueOf(lv).Elem(), inputValue, prop)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:965
								// _ = "end of CoverTab[142322]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:966
								_go_fuzz_dep_.CoverTab[142323]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:966
								if err := json.Unmarshal(inputValue, &map[string]json.RawMessage{}); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:966
									_go_fuzz_dep_.CoverTab[142324]++
																			sv := &types.Struct{}
																			target.Field(0).Set(reflect.ValueOf(&types.Value_StructValue{StructValue: sv}))
																			return u.unmarshalValue(reflect.ValueOf(sv).Elem(), inputValue, prop)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:969
									// _ = "end of CoverTab[142324]"
								} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:970
									_go_fuzz_dep_.CoverTab[142325]++
																			return fmt.Errorf("unrecognized type for Value %q", ivStr)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:971
									// _ = "end of CoverTab[142325]"
								}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
								// _ = "end of CoverTab[142323]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
							}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
							// _ = "end of CoverTab[142321]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
						}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
						// _ = "end of CoverTab[142319]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
					// _ = "end of CoverTab[142317]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
				// _ = "end of CoverTab[142315]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
			// _ = "end of CoverTab[142267]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:972
			_go_fuzz_dep_.CoverTab[142268]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:973
			// _ = "end of CoverTab[142268]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:973
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:973
			_go_fuzz_dep_.CoverTab[142269]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:973
			// _ = "end of CoverTab[142269]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:974
		// _ = "end of CoverTab[142245]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:975
		_go_fuzz_dep_.CoverTab[142326]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:975
		// _ = "end of CoverTab[142326]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:975
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:975
	// _ = "end of CoverTab[142225]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:975
	_go_fuzz_dep_.CoverTab[142226]++

											if t, ok := target.Addr().Interface().(*time.Time); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:977
		_go_fuzz_dep_.CoverTab[142327]++
												ts := &types.Timestamp{}
												if err := u.unmarshalValue(reflect.ValueOf(ts).Elem(), inputValue, prop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:979
			_go_fuzz_dep_.CoverTab[142330]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:980
			// _ = "end of CoverTab[142330]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:981
			_go_fuzz_dep_.CoverTab[142331]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:981
			// _ = "end of CoverTab[142331]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:981
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:981
		// _ = "end of CoverTab[142327]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:981
		_go_fuzz_dep_.CoverTab[142328]++
												tt, err := types.TimestampFromProto(ts)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:983
			_go_fuzz_dep_.CoverTab[142332]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:984
			// _ = "end of CoverTab[142332]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:985
			_go_fuzz_dep_.CoverTab[142333]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:985
			// _ = "end of CoverTab[142333]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:985
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:985
		// _ = "end of CoverTab[142328]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:985
		_go_fuzz_dep_.CoverTab[142329]++
												*t = tt
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:987
		// _ = "end of CoverTab[142329]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:988
		_go_fuzz_dep_.CoverTab[142334]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:988
		// _ = "end of CoverTab[142334]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:988
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:988
	// _ = "end of CoverTab[142226]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:988
	_go_fuzz_dep_.CoverTab[142227]++

											if d, ok := target.Addr().Interface().(*time.Duration); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:990
		_go_fuzz_dep_.CoverTab[142335]++
												dur := &types.Duration{}
												if err := u.unmarshalValue(reflect.ValueOf(dur).Elem(), inputValue, prop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:992
			_go_fuzz_dep_.CoverTab[142338]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:993
			// _ = "end of CoverTab[142338]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:994
			_go_fuzz_dep_.CoverTab[142339]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:994
			// _ = "end of CoverTab[142339]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:994
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:994
		// _ = "end of CoverTab[142335]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:994
		_go_fuzz_dep_.CoverTab[142336]++
												dd, err := types.DurationFromProto(dur)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:996
			_go_fuzz_dep_.CoverTab[142340]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:997
			// _ = "end of CoverTab[142340]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:998
			_go_fuzz_dep_.CoverTab[142341]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:998
			// _ = "end of CoverTab[142341]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:998
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:998
			// _ = "end of CoverTab[142336]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:998
			_go_fuzz_dep_.CoverTab[142337]++
													*d = dd
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1000
		// _ = "end of CoverTab[142337]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1001
		_go_fuzz_dep_.CoverTab[142342]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1001
		// _ = "end of CoverTab[142342]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1001
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1001
	// _ = "end of CoverTab[142227]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1001
	_go_fuzz_dep_.CoverTab[142228]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
	if inputValue[0] == '"' && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		_go_fuzz_dep_.CoverTab[142343]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		return prop != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		// _ = "end of CoverTab[142343]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		_go_fuzz_dep_.CoverTab[142344]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		return prop.Enum != ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		// _ = "end of CoverTab[142344]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1007
		_go_fuzz_dep_.CoverTab[142345]++
													vmap := proto.EnumValueMap(prop.Enum)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1011
		s := inputValue[1 : len(inputValue)-1]
		n, ok := vmap[string(s)]
		if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1013
			_go_fuzz_dep_.CoverTab[142349]++
														return fmt.Errorf("unknown value %q for enum %s", s, prop.Enum)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1014
			// _ = "end of CoverTab[142349]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1015
			_go_fuzz_dep_.CoverTab[142350]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1015
			// _ = "end of CoverTab[142350]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1015
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1015
		// _ = "end of CoverTab[142345]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1015
		_go_fuzz_dep_.CoverTab[142346]++
													if target.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1016
			_go_fuzz_dep_.CoverTab[142351]++
														target.Set(reflect.New(targetType.Elem()))
														target = target.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1018
			// _ = "end of CoverTab[142351]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1019
			_go_fuzz_dep_.CoverTab[142352]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1019
			// _ = "end of CoverTab[142352]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1019
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1019
		// _ = "end of CoverTab[142346]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1019
		_go_fuzz_dep_.CoverTab[142347]++
													if targetType.Kind() != reflect.Int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1020
			_go_fuzz_dep_.CoverTab[142353]++
														return fmt.Errorf("invalid target %q for enum %s", targetType.Kind(), prop.Enum)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1021
			// _ = "end of CoverTab[142353]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1022
			_go_fuzz_dep_.CoverTab[142354]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1022
			// _ = "end of CoverTab[142354]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1022
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1022
		// _ = "end of CoverTab[142347]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1022
		_go_fuzz_dep_.CoverTab[142348]++
													target.SetInt(int64(n))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1024
		// _ = "end of CoverTab[142348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1025
		_go_fuzz_dep_.CoverTab[142355]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1025
		// _ = "end of CoverTab[142355]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1025
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1025
	// _ = "end of CoverTab[142228]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1025
	_go_fuzz_dep_.CoverTab[142229]++

												if prop != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		_go_fuzz_dep_.CoverTab[142356]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		return len(prop.CustomType) > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		// _ = "end of CoverTab[142356]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		_go_fuzz_dep_.CoverTab[142357]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		return target.CanAddr()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		// _ = "end of CoverTab[142357]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1027
		_go_fuzz_dep_.CoverTab[142358]++
													if m, ok := target.Addr().Interface().(interface {
			UnmarshalJSON([]byte) error
		}); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1030
			_go_fuzz_dep_.CoverTab[142359]++
														return json.Unmarshal(inputValue, m)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1031
			// _ = "end of CoverTab[142359]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1032
			_go_fuzz_dep_.CoverTab[142360]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1032
			// _ = "end of CoverTab[142360]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1032
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1032
		// _ = "end of CoverTab[142358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1033
		_go_fuzz_dep_.CoverTab[142361]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1033
		// _ = "end of CoverTab[142361]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1033
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1033
	// _ = "end of CoverTab[142229]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1033
	_go_fuzz_dep_.CoverTab[142230]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1036
	if targetType.Kind() == reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1036
		_go_fuzz_dep_.CoverTab[142362]++
													var jsonFields map[string]json.RawMessage
													if err := json.Unmarshal(inputValue, &jsonFields); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1038
			_go_fuzz_dep_.CoverTab[142369]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1039
			// _ = "end of CoverTab[142369]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1040
			_go_fuzz_dep_.CoverTab[142370]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1040
			// _ = "end of CoverTab[142370]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1040
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1040
		// _ = "end of CoverTab[142362]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1040
		_go_fuzz_dep_.CoverTab[142363]++

													consumeField := func(prop *proto.Properties) (json.RawMessage, bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1042
			_go_fuzz_dep_.CoverTab[142371]++

														fieldNames := acceptedJSONFieldNames(prop)

														vOrig, okOrig := jsonFields[fieldNames.orig]
														vCamel, okCamel := jsonFields[fieldNames.camel]
														if !okOrig && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1048
				_go_fuzz_dep_.CoverTab[142375]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1048
				return !okCamel
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1048
				// _ = "end of CoverTab[142375]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1048
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1048
				_go_fuzz_dep_.CoverTab[142376]++
															return nil, false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1049
				// _ = "end of CoverTab[142376]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1050
				_go_fuzz_dep_.CoverTab[142377]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1050
				// _ = "end of CoverTab[142377]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1050
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1050
			// _ = "end of CoverTab[142371]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1050
			_go_fuzz_dep_.CoverTab[142372]++
			// If, for some reason, both are present in the data, favour the camelName.
			var raw json.RawMessage
			if okOrig {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1053
				_go_fuzz_dep_.CoverTab[142378]++
															raw = vOrig
															delete(jsonFields, fieldNames.orig)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1055
				// _ = "end of CoverTab[142378]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1056
				_go_fuzz_dep_.CoverTab[142379]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1056
				// _ = "end of CoverTab[142379]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1056
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1056
			// _ = "end of CoverTab[142372]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1056
			_go_fuzz_dep_.CoverTab[142373]++
														if okCamel {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1057
				_go_fuzz_dep_.CoverTab[142380]++
															raw = vCamel
															delete(jsonFields, fieldNames.camel)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1059
				// _ = "end of CoverTab[142380]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1060
				_go_fuzz_dep_.CoverTab[142381]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1060
				// _ = "end of CoverTab[142381]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1060
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1060
			// _ = "end of CoverTab[142373]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1060
			_go_fuzz_dep_.CoverTab[142374]++
														return raw, true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1061
			// _ = "end of CoverTab[142374]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1062
		// _ = "end of CoverTab[142363]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1062
		_go_fuzz_dep_.CoverTab[142364]++

													sprops := proto.GetProperties(targetType)
													for i := 0; i < target.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1065
			_go_fuzz_dep_.CoverTab[142382]++
														ft := target.Type().Field(i)
														if strings.HasPrefix(ft.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1067
				_go_fuzz_dep_.CoverTab[142385]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1068
				// _ = "end of CoverTab[142385]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1069
				_go_fuzz_dep_.CoverTab[142386]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1069
				// _ = "end of CoverTab[142386]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1069
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1069
			// _ = "end of CoverTab[142382]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1069
			_go_fuzz_dep_.CoverTab[142383]++
														valueForField, ok := consumeField(sprops.Prop[i])
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1071
				_go_fuzz_dep_.CoverTab[142387]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1072
				// _ = "end of CoverTab[142387]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1073
				_go_fuzz_dep_.CoverTab[142388]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1073
				// _ = "end of CoverTab[142388]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1073
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1073
			// _ = "end of CoverTab[142383]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1073
			_go_fuzz_dep_.CoverTab[142384]++

														if err := u.unmarshalValue(target.Field(i), valueForField, sprops.Prop[i]); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1075
				_go_fuzz_dep_.CoverTab[142389]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1076
				// _ = "end of CoverTab[142389]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1077
				_go_fuzz_dep_.CoverTab[142390]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1077
				// _ = "end of CoverTab[142390]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1077
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1077
			// _ = "end of CoverTab[142384]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1078
		// _ = "end of CoverTab[142364]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1078
		_go_fuzz_dep_.CoverTab[142365]++

													if len(jsonFields) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1080
			_go_fuzz_dep_.CoverTab[142391]++
														for _, oop := range sprops.OneofTypes {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1081
				_go_fuzz_dep_.CoverTab[142392]++
															raw, ok := consumeField(oop.Prop)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1083
					_go_fuzz_dep_.CoverTab[142394]++
																continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1084
					// _ = "end of CoverTab[142394]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1085
					_go_fuzz_dep_.CoverTab[142395]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1085
					// _ = "end of CoverTab[142395]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1085
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1085
				// _ = "end of CoverTab[142392]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1085
				_go_fuzz_dep_.CoverTab[142393]++
															nv := reflect.New(oop.Type.Elem())
															target.Field(oop.Field).Set(nv)
															if err := u.unmarshalValue(nv.Elem().Field(0), raw, oop.Prop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1088
					_go_fuzz_dep_.CoverTab[142396]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1089
					// _ = "end of CoverTab[142396]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1090
					_go_fuzz_dep_.CoverTab[142397]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1090
					// _ = "end of CoverTab[142397]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1090
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1090
				// _ = "end of CoverTab[142393]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1091
			// _ = "end of CoverTab[142391]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1092
			_go_fuzz_dep_.CoverTab[142398]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1092
			// _ = "end of CoverTab[142398]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1092
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1092
		// _ = "end of CoverTab[142365]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1092
		_go_fuzz_dep_.CoverTab[142366]++

													if len(jsonFields) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1094
			_go_fuzz_dep_.CoverTab[142399]++
														if ep, ok := target.Addr().Interface().(proto.Message); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1095
				_go_fuzz_dep_.CoverTab[142400]++
															for _, ext := range proto.RegisteredExtensions(ep) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1096
					_go_fuzz_dep_.CoverTab[142401]++
																name := fmt.Sprintf("[%s]", ext.Name)
																raw, ok := jsonFields[name]
																if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1099
						_go_fuzz_dep_.CoverTab[142404]++
																	continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1100
						// _ = "end of CoverTab[142404]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1101
						_go_fuzz_dep_.CoverTab[142405]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1101
						// _ = "end of CoverTab[142405]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1101
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1101
					// _ = "end of CoverTab[142401]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1101
					_go_fuzz_dep_.CoverTab[142402]++
																delete(jsonFields, name)
																nv := reflect.New(reflect.TypeOf(ext.ExtensionType).Elem())
																if err := u.unmarshalValue(nv.Elem(), raw, nil); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1104
						_go_fuzz_dep_.CoverTab[142406]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1105
						// _ = "end of CoverTab[142406]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1106
						_go_fuzz_dep_.CoverTab[142407]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1106
						// _ = "end of CoverTab[142407]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1106
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1106
					// _ = "end of CoverTab[142402]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1106
					_go_fuzz_dep_.CoverTab[142403]++
																if err := proto.SetExtension(ep, ext, nv.Interface()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1107
						_go_fuzz_dep_.CoverTab[142408]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1108
						// _ = "end of CoverTab[142408]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1109
						_go_fuzz_dep_.CoverTab[142409]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1109
						// _ = "end of CoverTab[142409]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1109
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1109
					// _ = "end of CoverTab[142403]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1110
				// _ = "end of CoverTab[142400]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1111
				_go_fuzz_dep_.CoverTab[142410]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1111
				// _ = "end of CoverTab[142410]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1111
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1111
			// _ = "end of CoverTab[142399]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1112
			_go_fuzz_dep_.CoverTab[142411]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1112
			// _ = "end of CoverTab[142411]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1112
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1112
		// _ = "end of CoverTab[142366]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1112
		_go_fuzz_dep_.CoverTab[142367]++
													if !u.AllowUnknownFields && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1113
			_go_fuzz_dep_.CoverTab[142412]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1113
			return len(jsonFields) > 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1113
			// _ = "end of CoverTab[142412]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1113
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1113
			_go_fuzz_dep_.CoverTab[142413]++
			// Pick any field to be the scapegoat.
			var f string
			for fname := range jsonFields {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1116
				_go_fuzz_dep_.CoverTab[142415]++
															f = fname
															break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1118
				// _ = "end of CoverTab[142415]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1119
			// _ = "end of CoverTab[142413]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1119
			_go_fuzz_dep_.CoverTab[142414]++
														return fmt.Errorf("unknown field %q in %v", f, targetType)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1120
			// _ = "end of CoverTab[142414]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1121
			_go_fuzz_dep_.CoverTab[142416]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1121
			// _ = "end of CoverTab[142416]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1121
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1121
		// _ = "end of CoverTab[142367]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1121
		_go_fuzz_dep_.CoverTab[142368]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1122
		// _ = "end of CoverTab[142368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1123
		_go_fuzz_dep_.CoverTab[142417]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1123
		// _ = "end of CoverTab[142417]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1123
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1123
	// _ = "end of CoverTab[142230]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1123
	_go_fuzz_dep_.CoverTab[142231]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1126
	if targetType.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1126
		_go_fuzz_dep_.CoverTab[142418]++
													if targetType.Elem().Kind() == reflect.Uint8 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1127
			_go_fuzz_dep_.CoverTab[142422]++
														outRef := reflect.New(targetType)
														outVal := outRef.Interface()

														if _, ok := outVal.(interface {
				UnmarshalJSON([]byte) error
			}); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1133
				_go_fuzz_dep_.CoverTab[142425]++
															if err := json.Unmarshal(inputValue, outVal); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1134
					_go_fuzz_dep_.CoverTab[142427]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1135
					// _ = "end of CoverTab[142427]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1136
					_go_fuzz_dep_.CoverTab[142428]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1136
					// _ = "end of CoverTab[142428]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1136
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1136
				// _ = "end of CoverTab[142425]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1136
				_go_fuzz_dep_.CoverTab[142426]++
															target.Set(outRef.Elem())
															return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1138
				// _ = "end of CoverTab[142426]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1139
				_go_fuzz_dep_.CoverTab[142429]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1139
				// _ = "end of CoverTab[142429]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1139
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1139
			// _ = "end of CoverTab[142422]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1139
			_go_fuzz_dep_.CoverTab[142423]++
			// Special case for encoded bytes. Pre-go1.5 doesn't support unmarshalling
			// strings into aliased []byte types.
			// https://github.com/golang/go/commit/4302fd0409da5e4f1d71471a6770dacdc3301197
			// https://github.com/golang/go/commit/c60707b14d6be26bf4213114d13070bff00d0b0a
			var out []byte
			if err := json.Unmarshal(inputValue, &out); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1145
				_go_fuzz_dep_.CoverTab[142430]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1146
				// _ = "end of CoverTab[142430]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1147
				_go_fuzz_dep_.CoverTab[142431]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1147
				// _ = "end of CoverTab[142431]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1147
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1147
			// _ = "end of CoverTab[142423]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1147
			_go_fuzz_dep_.CoverTab[142424]++
														target.SetBytes(out)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1149
			// _ = "end of CoverTab[142424]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1150
			_go_fuzz_dep_.CoverTab[142432]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1150
			// _ = "end of CoverTab[142432]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1150
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1150
		// _ = "end of CoverTab[142418]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1150
		_go_fuzz_dep_.CoverTab[142419]++

													var slc []json.RawMessage
													if err := json.Unmarshal(inputValue, &slc); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1153
			_go_fuzz_dep_.CoverTab[142433]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1154
			// _ = "end of CoverTab[142433]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1155
			_go_fuzz_dep_.CoverTab[142434]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1155
			// _ = "end of CoverTab[142434]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1155
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1155
		// _ = "end of CoverTab[142419]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1155
		_go_fuzz_dep_.CoverTab[142420]++
													if slc != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1156
			_go_fuzz_dep_.CoverTab[142435]++
														l := len(slc)
														target.Set(reflect.MakeSlice(targetType, l, l))
														for i := 0; i < l; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1159
				_go_fuzz_dep_.CoverTab[142436]++
															if err := u.unmarshalValue(target.Index(i), slc[i], prop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1160
					_go_fuzz_dep_.CoverTab[142437]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1161
					// _ = "end of CoverTab[142437]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1162
					_go_fuzz_dep_.CoverTab[142438]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1162
					// _ = "end of CoverTab[142438]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1162
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1162
				// _ = "end of CoverTab[142436]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1163
			// _ = "end of CoverTab[142435]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1164
			_go_fuzz_dep_.CoverTab[142439]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1164
			// _ = "end of CoverTab[142439]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1164
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1164
		// _ = "end of CoverTab[142420]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1164
		_go_fuzz_dep_.CoverTab[142421]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1165
		// _ = "end of CoverTab[142421]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1166
		_go_fuzz_dep_.CoverTab[142440]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1166
		// _ = "end of CoverTab[142440]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1166
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1166
	// _ = "end of CoverTab[142231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1166
	_go_fuzz_dep_.CoverTab[142232]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1169
	if targetType.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1169
		_go_fuzz_dep_.CoverTab[142441]++
													var mp map[string]json.RawMessage
													if err := json.Unmarshal(inputValue, &mp); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1171
			_go_fuzz_dep_.CoverTab[142444]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1172
			// _ = "end of CoverTab[142444]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1173
			_go_fuzz_dep_.CoverTab[142445]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1173
			// _ = "end of CoverTab[142445]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1173
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1173
		// _ = "end of CoverTab[142441]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1173
		_go_fuzz_dep_.CoverTab[142442]++
													if mp != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1174
			_go_fuzz_dep_.CoverTab[142446]++
														target.Set(reflect.MakeMap(targetType))
														for ks, raw := range mp {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1176
				_go_fuzz_dep_.CoverTab[142447]++
				// Unmarshal map key. The core json library already decoded the key into a
				// string, so we handle that specially. Other types were quoted post-serialization.
				var k reflect.Value
				if targetType.Key().Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1180
					_go_fuzz_dep_.CoverTab[142452]++
																k = reflect.ValueOf(ks)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1181
					// _ = "end of CoverTab[142452]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1182
					_go_fuzz_dep_.CoverTab[142453]++
																k = reflect.New(targetType.Key()).Elem()
																var kprop *proto.Properties
																if prop != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1185
						_go_fuzz_dep_.CoverTab[142455]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1185
						return prop.MapKeyProp != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1185
						// _ = "end of CoverTab[142455]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1185
					}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1185
						_go_fuzz_dep_.CoverTab[142456]++
																	kprop = prop.MapKeyProp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1186
						// _ = "end of CoverTab[142456]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1187
						_go_fuzz_dep_.CoverTab[142457]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1187
						// _ = "end of CoverTab[142457]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1187
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1187
					// _ = "end of CoverTab[142453]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1187
					_go_fuzz_dep_.CoverTab[142454]++
																if err := u.unmarshalValue(k, json.RawMessage(ks), kprop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1188
						_go_fuzz_dep_.CoverTab[142458]++
																	return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1189
						// _ = "end of CoverTab[142458]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1190
						_go_fuzz_dep_.CoverTab[142459]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1190
						// _ = "end of CoverTab[142459]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1190
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1190
					// _ = "end of CoverTab[142454]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1191
				// _ = "end of CoverTab[142447]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1191
				_go_fuzz_dep_.CoverTab[142448]++

															if !k.Type().AssignableTo(targetType.Key()) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1193
					_go_fuzz_dep_.CoverTab[142460]++
																k = k.Convert(targetType.Key())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1194
					// _ = "end of CoverTab[142460]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1195
					_go_fuzz_dep_.CoverTab[142461]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1195
					// _ = "end of CoverTab[142461]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1195
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1195
				// _ = "end of CoverTab[142448]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1195
				_go_fuzz_dep_.CoverTab[142449]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1198
				v := reflect.New(targetType.Elem()).Elem()
				var vprop *proto.Properties
				if prop != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1200
					_go_fuzz_dep_.CoverTab[142462]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1200
					return prop.MapValProp != nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1200
					// _ = "end of CoverTab[142462]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1200
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1200
					_go_fuzz_dep_.CoverTab[142463]++
																vprop = prop.MapValProp
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1201
					// _ = "end of CoverTab[142463]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1202
					_go_fuzz_dep_.CoverTab[142464]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1202
					// _ = "end of CoverTab[142464]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1202
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1202
				// _ = "end of CoverTab[142449]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1202
				_go_fuzz_dep_.CoverTab[142450]++
															if err := u.unmarshalValue(v, raw, vprop); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1203
					_go_fuzz_dep_.CoverTab[142465]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1204
					// _ = "end of CoverTab[142465]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1205
					_go_fuzz_dep_.CoverTab[142466]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1205
					// _ = "end of CoverTab[142466]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1205
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1205
				// _ = "end of CoverTab[142450]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1205
				_go_fuzz_dep_.CoverTab[142451]++
															target.SetMapIndex(k, v)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1206
				// _ = "end of CoverTab[142451]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1207
			// _ = "end of CoverTab[142446]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1208
			_go_fuzz_dep_.CoverTab[142467]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1208
			// _ = "end of CoverTab[142467]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1208
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1208
		// _ = "end of CoverTab[142442]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1208
		_go_fuzz_dep_.CoverTab[142443]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1209
		// _ = "end of CoverTab[142443]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1210
		_go_fuzz_dep_.CoverTab[142468]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1210
		// _ = "end of CoverTab[142468]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1210
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1210
	// _ = "end of CoverTab[142232]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1210
	_go_fuzz_dep_.CoverTab[142233]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1213
	isFloat := targetType.Kind() == reflect.Float32 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1213
		_go_fuzz_dep_.CoverTab[142469]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1213
		return targetType.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1213
		// _ = "end of CoverTab[142469]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1213
	}()
	if isFloat {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1214
		_go_fuzz_dep_.CoverTab[142470]++
													if num, ok := nonFinite[string(inputValue)]; ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1215
			_go_fuzz_dep_.CoverTab[142471]++
														target.SetFloat(num)
														return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1217
			// _ = "end of CoverTab[142471]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1218
			_go_fuzz_dep_.CoverTab[142472]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1218
			// _ = "end of CoverTab[142472]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1218
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1218
		// _ = "end of CoverTab[142470]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1219
		_go_fuzz_dep_.CoverTab[142473]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1219
		// _ = "end of CoverTab[142473]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1219
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1219
	// _ = "end of CoverTab[142233]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1219
	_go_fuzz_dep_.CoverTab[142234]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
	isNum := targetType.Kind() == reflect.Int64 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
		_go_fuzz_dep_.CoverTab[142474]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
		return targetType.Kind() == reflect.Uint64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
		// _ = "end of CoverTab[142474]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
		_go_fuzz_dep_.CoverTab[142475]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1223
		return targetType.Kind() == reflect.Int32
													// _ = "end of CoverTab[142475]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
		_go_fuzz_dep_.CoverTab[142476]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
		return targetType.Kind() == reflect.Uint32
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
		// _ = "end of CoverTab[142476]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
		_go_fuzz_dep_.CoverTab[142477]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1224
		return targetType.Kind() == reflect.Float32
													// _ = "end of CoverTab[142477]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1225
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1225
		_go_fuzz_dep_.CoverTab[142478]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1225
		return targetType.Kind() == reflect.Float64
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1225
		// _ = "end of CoverTab[142478]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1225
	}()
	if isNum && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1226
		_go_fuzz_dep_.CoverTab[142479]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1226
		return strings.HasPrefix(string(inputValue), `"`)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1226
		// _ = "end of CoverTab[142479]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1226
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1226
		_go_fuzz_dep_.CoverTab[142480]++
													inputValue = inputValue[1 : len(inputValue)-1]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1227
		// _ = "end of CoverTab[142480]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1228
		_go_fuzz_dep_.CoverTab[142481]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1228
		// _ = "end of CoverTab[142481]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1228
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1228
	// _ = "end of CoverTab[142234]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1228
	_go_fuzz_dep_.CoverTab[142235]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1231
	return json.Unmarshal(inputValue, target.Addr().Interface())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1231
	// _ = "end of CoverTab[142235]"
}

func unquote(s string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1234
	_go_fuzz_dep_.CoverTab[142482]++
												var ret string
												err := json.Unmarshal([]byte(s), &ret)
												return ret, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1237
	// _ = "end of CoverTab[142482]"
}

// jsonProperties returns parsed proto.Properties for the field and corrects JSONName attribute.
func jsonProperties(f reflect.StructField, origName bool) *proto.Properties {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1241
	_go_fuzz_dep_.CoverTab[142483]++
												var prop proto.Properties
												prop.Init(f.Type, f.Name, f.Tag.Get("protobuf"), &f)
												if origName || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1244
		_go_fuzz_dep_.CoverTab[142485]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1244
		return prop.JSONName == ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1244
		// _ = "end of CoverTab[142485]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1244
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1244
		_go_fuzz_dep_.CoverTab[142486]++
													prop.JSONName = prop.OrigName
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1245
		// _ = "end of CoverTab[142486]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1246
		_go_fuzz_dep_.CoverTab[142487]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1246
		// _ = "end of CoverTab[142487]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1246
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1246
	// _ = "end of CoverTab[142483]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1246
	_go_fuzz_dep_.CoverTab[142484]++
												return &prop
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1247
	// _ = "end of CoverTab[142484]"
}

type fieldNames struct {
	orig, camel string
}

func acceptedJSONFieldNames(prop *proto.Properties) fieldNames {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1254
	_go_fuzz_dep_.CoverTab[142488]++
												opts := fieldNames{orig: prop.OrigName, camel: prop.OrigName}
												if prop.JSONName != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1256
		_go_fuzz_dep_.CoverTab[142490]++
													opts.camel = prop.JSONName
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1257
		// _ = "end of CoverTab[142490]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1258
		_go_fuzz_dep_.CoverTab[142491]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1258
		// _ = "end of CoverTab[142491]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1258
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1258
	// _ = "end of CoverTab[142488]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1258
	_go_fuzz_dep_.CoverTab[142489]++
												return opts
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1259
	// _ = "end of CoverTab[142489]"
}

// Writer wrapper inspired by https://blog.golang.org/errors-are-values
type errWriter struct {
	writer	io.Writer
	err	error
}

func (w *errWriter) write(str string) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1268
	_go_fuzz_dep_.CoverTab[142492]++
												if w.err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1269
		_go_fuzz_dep_.CoverTab[142494]++
													return
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1270
		// _ = "end of CoverTab[142494]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1271
		_go_fuzz_dep_.CoverTab[142495]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1271
		// _ = "end of CoverTab[142495]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1271
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1271
	// _ = "end of CoverTab[142492]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1271
	_go_fuzz_dep_.CoverTab[142493]++
												_, w.err = w.writer.Write([]byte(str))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1272
	// _ = "end of CoverTab[142493]"
}

// Map fields may have key types of non-float scalars, strings and enums.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1275
// The easiest way to sort them in some deterministic order is to use fmt.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1275
// If this turns out to be inefficient we can always consider other options,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1275
// such as doing a Schwartzian transform.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1275
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1275
// Numeric keys are sorted in numeric order per
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1275
// https://developers.google.com/protocol-buffers/docs/proto#maps.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1282
type mapKeys []reflect.Value

func (s mapKeys) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1284
	_go_fuzz_dep_.CoverTab[142496]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1284
	return len(s)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1284
	// _ = "end of CoverTab[142496]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1284
}
func (s mapKeys) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1285
	_go_fuzz_dep_.CoverTab[142497]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1285
	s[i], s[j] = s[j], s[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1285
	// _ = "end of CoverTab[142497]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1285
}
func (s mapKeys) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1286
	_go_fuzz_dep_.CoverTab[142498]++
												if k := s[i].Kind(); k == s[j].Kind() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1287
		_go_fuzz_dep_.CoverTab[142500]++
													switch k {
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1289
			_go_fuzz_dep_.CoverTab[142501]++
														return s[i].String() < s[j].String()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1290
			// _ = "end of CoverTab[142501]"
		case reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1291
			_go_fuzz_dep_.CoverTab[142502]++
														return s[i].Int() < s[j].Int()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1292
			// _ = "end of CoverTab[142502]"
		case reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1293
			_go_fuzz_dep_.CoverTab[142503]++
														return s[i].Uint() < s[j].Uint()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1294
			// _ = "end of CoverTab[142503]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1294
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1294
			_go_fuzz_dep_.CoverTab[142504]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1294
			// _ = "end of CoverTab[142504]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1295
		// _ = "end of CoverTab[142500]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1296
		_go_fuzz_dep_.CoverTab[142505]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1296
		// _ = "end of CoverTab[142505]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1296
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1296
	// _ = "end of CoverTab[142498]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1296
	_go_fuzz_dep_.CoverTab[142499]++
												return fmt.Sprint(s[i].Interface()) < fmt.Sprint(s[j].Interface())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1297
	// _ = "end of CoverTab[142499]"
}

// checkRequiredFields returns an error if any required field in the given proto message is not set.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1300
// This function is used by both Marshal and Unmarshal.  While required fields only exist in a
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1300
// proto2 message, a proto3 message can contain proto2 message(s).
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1303
func checkRequiredFields(pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1303
	_go_fuzz_dep_.CoverTab[142506]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1314
	if _, ok := pb.(isWkt); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1314
		_go_fuzz_dep_.CoverTab[142512]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1315
		// _ = "end of CoverTab[142512]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1316
		_go_fuzz_dep_.CoverTab[142513]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1316
		// _ = "end of CoverTab[142513]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1316
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1316
	// _ = "end of CoverTab[142506]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1316
	_go_fuzz_dep_.CoverTab[142507]++

												v := reflect.ValueOf(pb)

												if v.Kind() != reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1320
		_go_fuzz_dep_.CoverTab[142514]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1321
		// _ = "end of CoverTab[142514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1322
		_go_fuzz_dep_.CoverTab[142515]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1322
		// _ = "end of CoverTab[142515]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1322
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1322
	// _ = "end of CoverTab[142507]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1322
	_go_fuzz_dep_.CoverTab[142508]++
												v = v.Elem()
												if v.Kind() != reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1324
		_go_fuzz_dep_.CoverTab[142516]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1325
		// _ = "end of CoverTab[142516]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1326
		_go_fuzz_dep_.CoverTab[142517]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1326
		// _ = "end of CoverTab[142517]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1326
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1326
	// _ = "end of CoverTab[142508]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1326
	_go_fuzz_dep_.CoverTab[142509]++

												for i := 0; i < v.NumField(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1328
		_go_fuzz_dep_.CoverTab[142518]++
													field := v.Field(i)
													sfield := v.Type().Field(i)

													if sfield.PkgPath != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1332
			_go_fuzz_dep_.CoverTab[142523]++

														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1334
			// _ = "end of CoverTab[142523]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1335
			_go_fuzz_dep_.CoverTab[142524]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1335
			// _ = "end of CoverTab[142524]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1335
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1335
		// _ = "end of CoverTab[142518]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1335
		_go_fuzz_dep_.CoverTab[142519]++

													if strings.HasPrefix(sfield.Name, "XXX_") {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1337
			_go_fuzz_dep_.CoverTab[142525]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1338
			// _ = "end of CoverTab[142525]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1339
			_go_fuzz_dep_.CoverTab[142526]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1339
			// _ = "end of CoverTab[142526]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1339
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1339
		// _ = "end of CoverTab[142519]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1339
		_go_fuzz_dep_.CoverTab[142520]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1343
		if sfield.Tag.Get("protobuf_oneof") != "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1343
			_go_fuzz_dep_.CoverTab[142527]++
														if field.Kind() != reflect.Interface {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1344
				_go_fuzz_dep_.CoverTab[142531]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1345
				// _ = "end of CoverTab[142531]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1346
				_go_fuzz_dep_.CoverTab[142532]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1346
				// _ = "end of CoverTab[142532]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1346
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1346
			// _ = "end of CoverTab[142527]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1346
			_go_fuzz_dep_.CoverTab[142528]++
														v := field.Elem()
														if v.Kind() != reflect.Ptr || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1348
				_go_fuzz_dep_.CoverTab[142533]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1348
				return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1348
				// _ = "end of CoverTab[142533]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1348
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1348
				_go_fuzz_dep_.CoverTab[142534]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1349
				// _ = "end of CoverTab[142534]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1350
				_go_fuzz_dep_.CoverTab[142535]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1350
				// _ = "end of CoverTab[142535]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1350
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1350
			// _ = "end of CoverTab[142528]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1350
			_go_fuzz_dep_.CoverTab[142529]++
														v = v.Elem()
														if v.Kind() != reflect.Struct || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1352
				_go_fuzz_dep_.CoverTab[142536]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1352
				return v.NumField() < 1
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1352
				// _ = "end of CoverTab[142536]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1352
			}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1352
				_go_fuzz_dep_.CoverTab[142537]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1353
				// _ = "end of CoverTab[142537]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1354
				_go_fuzz_dep_.CoverTab[142538]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1354
				// _ = "end of CoverTab[142538]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1354
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1354
			// _ = "end of CoverTab[142529]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1354
			_go_fuzz_dep_.CoverTab[142530]++
														field = v.Field(0)
														sfield = v.Type().Field(0)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1356
			// _ = "end of CoverTab[142530]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1357
			_go_fuzz_dep_.CoverTab[142539]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1357
			// _ = "end of CoverTab[142539]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1357
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1357
		// _ = "end of CoverTab[142520]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1357
		_go_fuzz_dep_.CoverTab[142521]++

													protoTag := sfield.Tag.Get("protobuf")
													if protoTag == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1360
			_go_fuzz_dep_.CoverTab[142540]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1361
			// _ = "end of CoverTab[142540]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1362
			_go_fuzz_dep_.CoverTab[142541]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1362
			// _ = "end of CoverTab[142541]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1362
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1362
		// _ = "end of CoverTab[142521]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1362
		_go_fuzz_dep_.CoverTab[142522]++
													var prop proto.Properties
													prop.Init(sfield.Type, sfield.Name, protoTag, &sfield)

													switch field.Kind() {
		case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1367
			_go_fuzz_dep_.CoverTab[142542]++
														if field.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1368
				_go_fuzz_dep_.CoverTab[142550]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1369
				// _ = "end of CoverTab[142550]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1370
				_go_fuzz_dep_.CoverTab[142551]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1370
				// _ = "end of CoverTab[142551]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1370
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1370
			// _ = "end of CoverTab[142542]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1370
			_go_fuzz_dep_.CoverTab[142543]++

														keys := field.MapKeys()
														for _, k := range keys {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1373
				_go_fuzz_dep_.CoverTab[142552]++
															v := field.MapIndex(k)
															if err := checkRequiredFieldsInValue(v); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1375
					_go_fuzz_dep_.CoverTab[142553]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1376
					// _ = "end of CoverTab[142553]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1377
					_go_fuzz_dep_.CoverTab[142554]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1377
					// _ = "end of CoverTab[142554]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1377
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1377
				// _ = "end of CoverTab[142552]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1378
			// _ = "end of CoverTab[142543]"
		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1379
			_go_fuzz_dep_.CoverTab[142544]++

														if !prop.Repeated {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1381
				_go_fuzz_dep_.CoverTab[142555]++
															if prop.Required && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1382
					_go_fuzz_dep_.CoverTab[142557]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1382
					return field.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1382
					// _ = "end of CoverTab[142557]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1382
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1382
					_go_fuzz_dep_.CoverTab[142558]++
																return fmt.Errorf("required field %q is not set", prop.Name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1383
					// _ = "end of CoverTab[142558]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1384
					_go_fuzz_dep_.CoverTab[142559]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1384
					// _ = "end of CoverTab[142559]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1384
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1384
				// _ = "end of CoverTab[142555]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1384
				_go_fuzz_dep_.CoverTab[142556]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1385
				// _ = "end of CoverTab[142556]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1386
				_go_fuzz_dep_.CoverTab[142560]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1386
				// _ = "end of CoverTab[142560]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1386
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1386
			// _ = "end of CoverTab[142544]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1386
			_go_fuzz_dep_.CoverTab[142545]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1389
			if field.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1389
				_go_fuzz_dep_.CoverTab[142561]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1390
				// _ = "end of CoverTab[142561]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1391
				_go_fuzz_dep_.CoverTab[142562]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1391
				// _ = "end of CoverTab[142562]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1391
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1391
			// _ = "end of CoverTab[142545]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1391
			_go_fuzz_dep_.CoverTab[142546]++

														for i := 0; i < field.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1393
				_go_fuzz_dep_.CoverTab[142563]++
															v := field.Index(i)
															if err := checkRequiredFieldsInValue(v); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1395
					_go_fuzz_dep_.CoverTab[142564]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1396
					// _ = "end of CoverTab[142564]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1397
					_go_fuzz_dep_.CoverTab[142565]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1397
					// _ = "end of CoverTab[142565]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1397
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1397
				// _ = "end of CoverTab[142563]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1398
			// _ = "end of CoverTab[142546]"
		case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1399
			_go_fuzz_dep_.CoverTab[142547]++
														if field.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1400
				_go_fuzz_dep_.CoverTab[142566]++
															if prop.Required {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1401
					_go_fuzz_dep_.CoverTab[142568]++
																return fmt.Errorf("required field %q is not set", prop.Name)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1402
					// _ = "end of CoverTab[142568]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1403
					_go_fuzz_dep_.CoverTab[142569]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1403
					// _ = "end of CoverTab[142569]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1403
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1403
				// _ = "end of CoverTab[142566]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1403
				_go_fuzz_dep_.CoverTab[142567]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1404
				// _ = "end of CoverTab[142567]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1405
				_go_fuzz_dep_.CoverTab[142570]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1405
				// _ = "end of CoverTab[142570]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1405
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1405
			// _ = "end of CoverTab[142547]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1405
			_go_fuzz_dep_.CoverTab[142548]++
														if err := checkRequiredFieldsInValue(field); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1406
				_go_fuzz_dep_.CoverTab[142571]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1407
				// _ = "end of CoverTab[142571]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
				_go_fuzz_dep_.CoverTab[142572]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
				// _ = "end of CoverTab[142572]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
			// _ = "end of CoverTab[142548]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
			_go_fuzz_dep_.CoverTab[142549]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1408
			// _ = "end of CoverTab[142549]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1409
		// _ = "end of CoverTab[142522]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1410
	// _ = "end of CoverTab[142509]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1410
	_go_fuzz_dep_.CoverTab[142510]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1413
	for _, ext := range proto.RegisteredExtensions(pb) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1413
		_go_fuzz_dep_.CoverTab[142573]++
													if !proto.HasExtension(pb, ext) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1414
			_go_fuzz_dep_.CoverTab[142576]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1415
			// _ = "end of CoverTab[142576]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1416
			_go_fuzz_dep_.CoverTab[142577]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1416
			// _ = "end of CoverTab[142577]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1416
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1416
		// _ = "end of CoverTab[142573]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1416
		_go_fuzz_dep_.CoverTab[142574]++
													ep, err := proto.GetExtension(pb, ext)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1418
			_go_fuzz_dep_.CoverTab[142578]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1419
			// _ = "end of CoverTab[142578]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1420
			_go_fuzz_dep_.CoverTab[142579]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1420
			// _ = "end of CoverTab[142579]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1420
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1420
		// _ = "end of CoverTab[142574]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1420
		_go_fuzz_dep_.CoverTab[142575]++
													err = checkRequiredFieldsInValue(reflect.ValueOf(ep))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1422
			_go_fuzz_dep_.CoverTab[142580]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1423
			// _ = "end of CoverTab[142580]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1424
			_go_fuzz_dep_.CoverTab[142581]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1424
			// _ = "end of CoverTab[142581]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1424
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1424
		// _ = "end of CoverTab[142575]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1425
	// _ = "end of CoverTab[142510]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1425
	_go_fuzz_dep_.CoverTab[142511]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1427
	// _ = "end of CoverTab[142511]"
}

func checkRequiredFieldsInValue(v reflect.Value) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1430
	_go_fuzz_dep_.CoverTab[142582]++
												if v.Type().Implements(messageType) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1431
		_go_fuzz_dep_.CoverTab[142584]++
													return checkRequiredFields(v.Interface().(proto.Message))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1432
		// _ = "end of CoverTab[142584]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1433
		_go_fuzz_dep_.CoverTab[142585]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1433
		// _ = "end of CoverTab[142585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1433
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1433
	// _ = "end of CoverTab[142582]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1433
	_go_fuzz_dep_.CoverTab[142583]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1434
	// _ = "end of CoverTab[142583]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1435
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/jsonpb/jsonpb.go:1435
var _ = _go_fuzz_dep_.CoverTab
