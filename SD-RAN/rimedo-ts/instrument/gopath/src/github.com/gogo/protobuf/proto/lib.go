// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:32
/*
Package proto converts data structures to and from the wire format of
protocol buffers.  It works in concert with the Go source code generated
for .proto files by the protocol compiler.

A summary of the properties of the protocol buffer interface
for a protocol buffer variable v:

  - Names are turned from camel_case to CamelCase for export.
  - There are no methods on v to set fields; just treat
    them as structure fields.
  - There are getters that return a field's value if set,
    and return the field's default value if unset.
    The getters work even if the receiver is a nil message.
  - The zero value for a struct is its correct initialization state.
    All desired fields must be set before marshaling.
  - A Reset() method will restore a protobuf struct to its zero state.
  - Non-repeated fields are pointers to the values; nil means unset.
    That is, optional or required field int32 f becomes F *int32.
  - Repeated fields are slices.
  - Helper functions are available to aid the setting of fields.
    msg.Foo = proto.String("hello") // set field
  - Constants are defined to hold the default values of all fields that
    have them.  They have the form Default_StructName_FieldName.
    Because the getter methods handle defaulted values,
    direct use of these constants should be rare.
  - Enums are given type names and maps from names to values.
    Enum values are prefixed by the enclosing message's name, or by the
    enum's type name if it is a top-level enum. Enum types have a String
    method, and a Enum method to assist in message construction.
  - Nested messages, groups and enums have type names prefixed with the name of
    the surrounding message type.
  - Extensions are given descriptor names that start with E_,
    followed by an underscore-delimited list of the nested messages
    that contain it (if any) followed by the CamelCased name of the
    extension field itself.  HasExtension, ClearExtension, GetExtension
    and SetExtension are functions for manipulating extensions.
  - Oneof field sets are given a single field in their message,
    with distinguished wrapper types for each possible field value.
  - Marshal and Unmarshal are functions to encode and decode the wire format.

When the .proto file specifies `syntax="proto3"`, there are some differences:

  - Non-repeated fields of non-message type are values instead of pointers.
  - Enum types do not get an Enum method.

The simplest way to describe this is to see an example.
Given file test.proto, containing

	package example;

	enum FOO { X = 17; }

	message Test {
	  required string label = 1;
	  optional int32 type = 2 [default=77];
	  repeated int64 reps = 3;
	  optional group OptionalGroup = 4 {
	    required string RequiredField = 5;
	  }
	  oneof union {
	    int32 number = 6;
	    string name = 7;
	  }
	}

The resulting file, test.pb.go, is:

	package example

	import proto "github.com/gogo/protobuf/proto"
	import math "math"

	type FOO int32
	const (
		FOO_X FOO = 17
	)
	var FOO_name = map[int32]string{
		17: "X",
	}
	var FOO_value = map[string]int32{
		"X": 17,
	}

	func (x FOO) Enum() *FOO {
		p := new(FOO)
		*p = x
		return p
	}
	func (x FOO) String() string {
		return proto.EnumName(FOO_name, int32(x))
	}
	func (x *FOO) UnmarshalJSON(data []byte) error {
		value, err := proto.UnmarshalJSONEnum(FOO_value, data)
		if err != nil {
			return err
		}
		*x = FOO(value)
		return nil
	}

	type Test struct {
		Label         *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
		Type          *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
		Reps          []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
		Optionalgroup *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup" json:"optionalgroup,omitempty"`
		// Types that are valid to be assigned to Union:
		//	*Test_Number
		//	*Test_Name
		Union            isTest_Union `protobuf_oneof:"union"`
		XXX_unrecognized []byte       `json:"-"`
	}
	func (m *Test) Reset()         { *m = Test{} }
	func (m *Test) String() string { return proto.CompactTextString(m) }
	func (*Test) ProtoMessage() {}

	type isTest_Union interface {
		isTest_Union()
	}

	type Test_Number struct {
		Number int32 `protobuf:"varint,6,opt,name=number"`
	}
	type Test_Name struct {
		Name string `protobuf:"bytes,7,opt,name=name"`
	}

	func (*Test_Number) isTest_Union() {}
	func (*Test_Name) isTest_Union()   {}

	func (m *Test) GetUnion() isTest_Union {
		if m != nil {
			return m.Union
		}
		return nil
	}
	const Default_Test_Type int32 = 77

	func (m *Test) GetLabel() string {
		if m != nil && m.Label != nil {
			return *m.Label
		}
		return ""
	}

	func (m *Test) GetType() int32 {
		if m != nil && m.Type != nil {
			return *m.Type
		}
		return Default_Test_Type
	}

	func (m *Test) GetOptionalgroup() *Test_OptionalGroup {
		if m != nil {
			return m.Optionalgroup
		}
		return nil
	}

	type Test_OptionalGroup struct {
		RequiredField *string `protobuf:"bytes,5,req" json:"RequiredField,omitempty"`
	}
	func (m *Test_OptionalGroup) Reset()         { *m = Test_OptionalGroup{} }
	func (m *Test_OptionalGroup) String() string { return proto.CompactTextString(m) }

	func (m *Test_OptionalGroup) GetRequiredField() string {
		if m != nil && m.RequiredField != nil {
			return *m.RequiredField
		}
		return ""
	}

	func (m *Test) GetNumber() int32 {
		if x, ok := m.GetUnion().(*Test_Number); ok {
			return x.Number
		}
		return 0
	}

	func (m *Test) GetName() string {
		if x, ok := m.GetUnion().(*Test_Name); ok {
			return x.Name
		}
		return ""
	}

	func init() {
		proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
	}

To create and play with a Test object:

	package main

	import (
		"log"

		"github.com/gogo/protobuf/proto"
		pb "./example.pb"
	)

	func main() {
		test := &pb.Test{
			Label: proto.String("hello"),
			Type:  proto.Int32(17),
			Reps:  []int64{1, 2, 3},
			Optionalgroup: &pb.Test_OptionalGroup{
				RequiredField: proto.String("good bye"),
			},
			Union: &pb.Test_Name{"fred"},
		}
		data, err := proto.Marshal(test)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		newTest := &pb.Test{}
		err = proto.Unmarshal(data, newTest)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		// Now test and newTest contain the same data.
		if test.GetLabel() != newTest.GetLabel() {
			log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
		}
		// Use a type switch to determine which oneof was set.
		switch u := test.Union.(type) {
		case *pb.Test_Number: // u.Number contains the number.
		case *pb.Test_Name: // u.Name contains the string.
		}
		// etc.
	}
*/
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:264
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:264
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:264
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:264
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:264
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:264
)

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"sync"
)

// RequiredNotSetError is an error type returned by either Marshal or Unmarshal.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:276
// Marshal reports this when a required field is not initialized.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:276
// Unmarshal reports this when a required field is missing from the wire data.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:279
type RequiredNotSetError struct{ field string }

func (e *RequiredNotSetError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:281
	_go_fuzz_dep_.CoverTab[108535]++
											if e.field == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:282
		_go_fuzz_dep_.CoverTab[108537]++
												return fmt.Sprintf("proto: required field not set")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:283
		// _ = "end of CoverTab[108537]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:284
		_go_fuzz_dep_.CoverTab[108538]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:284
		// _ = "end of CoverTab[108538]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:284
	// _ = "end of CoverTab[108535]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:284
	_go_fuzz_dep_.CoverTab[108536]++
											return fmt.Sprintf("proto: required field %q not set", e.field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:285
	// _ = "end of CoverTab[108536]"
}
func (e *RequiredNotSetError) RequiredNotSet() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:287
	_go_fuzz_dep_.CoverTab[108539]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:288
	// _ = "end of CoverTab[108539]"
}

type invalidUTF8Error struct{ field string }

func (e *invalidUTF8Error) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:293
	_go_fuzz_dep_.CoverTab[108540]++
											if e.field == "" {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:294
		_go_fuzz_dep_.CoverTab[108542]++
												return "proto: invalid UTF-8 detected"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:295
		// _ = "end of CoverTab[108542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:296
		_go_fuzz_dep_.CoverTab[108543]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:296
		// _ = "end of CoverTab[108543]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:296
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:296
	// _ = "end of CoverTab[108540]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:296
	_go_fuzz_dep_.CoverTab[108541]++
											return fmt.Sprintf("proto: field %q contains invalid UTF-8", e.field)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:297
	// _ = "end of CoverTab[108541]"
}
func (e *invalidUTF8Error) InvalidUTF8() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:299
	_go_fuzz_dep_.CoverTab[108544]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:300
	// _ = "end of CoverTab[108544]"
}

// errInvalidUTF8 is a sentinel error to identify fields with invalid UTF-8.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:303
// This error should not be exposed to the external API as such errors should
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:303
// be recreated with the field information.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:306
var errInvalidUTF8 = &invalidUTF8Error{}

// isNonFatal reports whether the error is either a RequiredNotSet error
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:308
// or a InvalidUTF8 error.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:310
func isNonFatal(err error) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:310
	_go_fuzz_dep_.CoverTab[108545]++
											if re, ok := err.(interface{ RequiredNotSet() bool }); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:311
		_go_fuzz_dep_.CoverTab[108548]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:311
		return re.RequiredNotSet()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:311
		// _ = "end of CoverTab[108548]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:311
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:311
		_go_fuzz_dep_.CoverTab[108549]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:312
		// _ = "end of CoverTab[108549]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:313
		_go_fuzz_dep_.CoverTab[108550]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:313
		// _ = "end of CoverTab[108550]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:313
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:313
	// _ = "end of CoverTab[108545]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:313
	_go_fuzz_dep_.CoverTab[108546]++
											if re, ok := err.(interface{ InvalidUTF8() bool }); ok && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:314
		_go_fuzz_dep_.CoverTab[108551]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:314
		return re.InvalidUTF8()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:314
		// _ = "end of CoverTab[108551]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:314
	}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:314
		_go_fuzz_dep_.CoverTab[108552]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:315
		// _ = "end of CoverTab[108552]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:316
		_go_fuzz_dep_.CoverTab[108553]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:316
		// _ = "end of CoverTab[108553]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:316
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:316
	// _ = "end of CoverTab[108546]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:316
	_go_fuzz_dep_.CoverTab[108547]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:317
	// _ = "end of CoverTab[108547]"
}

type nonFatal struct{ E error }

// Merge merges err into nf and reports whether it was successful.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:322
// Otherwise it returns false for any fatal non-nil errors.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:324
func (nf *nonFatal) Merge(err error) (ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:324
	_go_fuzz_dep_.CoverTab[108554]++
											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:325
		_go_fuzz_dep_.CoverTab[108558]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:326
		// _ = "end of CoverTab[108558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:327
		_go_fuzz_dep_.CoverTab[108559]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:327
		// _ = "end of CoverTab[108559]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:327
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:327
	// _ = "end of CoverTab[108554]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:327
	_go_fuzz_dep_.CoverTab[108555]++
											if !isNonFatal(err) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:328
		_go_fuzz_dep_.CoverTab[108560]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:329
		// _ = "end of CoverTab[108560]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:330
		_go_fuzz_dep_.CoverTab[108561]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:330
		// _ = "end of CoverTab[108561]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:330
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:330
	// _ = "end of CoverTab[108555]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:330
	_go_fuzz_dep_.CoverTab[108556]++
											if nf.E == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:331
		_go_fuzz_dep_.CoverTab[108562]++
												nf.E = err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:332
		// _ = "end of CoverTab[108562]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:333
		_go_fuzz_dep_.CoverTab[108563]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:333
		// _ = "end of CoverTab[108563]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:333
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:333
	// _ = "end of CoverTab[108556]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:333
	_go_fuzz_dep_.CoverTab[108557]++
											return true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:334
	// _ = "end of CoverTab[108557]"
}

// Message is implemented by generated protocol buffer messages.
type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

// A Buffer is a buffer manager for marshaling and unmarshaling
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:344
// protocol buffers.  It may be reused between invocations to
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:344
// reduce memory usage.  It is not necessary to use a Buffer;
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:344
// the global functions Marshal and Unmarshal create a
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:344
// temporary Buffer and are fine for most applications.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:349
type Buffer struct {
	buf	[]byte	// encode/decode byte stream
	index	int	// read point

	deterministic	bool
}

// NewBuffer allocates a new Buffer and initializes its internal data to
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:356
// the contents of the argument slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:358
func NewBuffer(e []byte) *Buffer {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:358
	_go_fuzz_dep_.CoverTab[108564]++
											return &Buffer{buf: e}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:359
	// _ = "end of CoverTab[108564]"
}

// Reset resets the Buffer, ready for marshaling a new protocol buffer.
func (p *Buffer) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:363
	_go_fuzz_dep_.CoverTab[108565]++
											p.buf = p.buf[0:0]
											p.index = 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:365
	// _ = "end of CoverTab[108565]"
}

// SetBuf replaces the internal buffer with the slice,
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:368
// ready for unmarshaling the contents of the slice.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:370
func (p *Buffer) SetBuf(s []byte) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:370
	_go_fuzz_dep_.CoverTab[108566]++
											p.buf = s
											p.index = 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:372
	// _ = "end of CoverTab[108566]"
}

// Bytes returns the contents of the Buffer.
func (p *Buffer) Bytes() []byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:376
	_go_fuzz_dep_.CoverTab[108567]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:376
	return p.buf
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:376
	// _ = "end of CoverTab[108567]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:376
}

// SetDeterministic sets whether to use deterministic serialization.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// Deterministic serialization guarantees that for a given binary, equal
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// messages will always be serialized to the same bytes. This implies:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//   - Repeated serialization of a message will return the same bytes.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//   - Different processes of the same binary (which may be executing on
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//     different machines) will serialize equal messages to the same bytes.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// Note that the deterministic serialization is NOT canonical across
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// languages. It is not guaranteed to remain stable over time. It is unstable
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// across different builds with schema changes due to unknown fields.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// Users who need canonical serialization (e.g., persistent storage in a
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// canonical form, fingerprinting, etc.) should define their own
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// canonicalization specification and implement their own serializer rather
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// than relying on this API.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// If deterministic serialization is requested, map entries will be sorted
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// by keys in lexographical order. This is an implementation detail and
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:378
// subject to change.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:398
func (p *Buffer) SetDeterministic(deterministic bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:398
	_go_fuzz_dep_.CoverTab[108568]++
											p.deterministic = deterministic
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:399
	// _ = "end of CoverTab[108568]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:406
// Bool is a helper routine that allocates a new bool value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:406
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:408
func Bool(v bool) *bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:408
	_go_fuzz_dep_.CoverTab[108569]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:409
	// _ = "end of CoverTab[108569]"
}

// Int32 is a helper routine that allocates a new int32 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:412
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:414
func Int32(v int32) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:414
	_go_fuzz_dep_.CoverTab[108570]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:415
	// _ = "end of CoverTab[108570]"
}

// Int is a helper routine that allocates a new int32 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:418
// to store v and returns a pointer to it, but unlike Int32
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:418
// its argument value is an int.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:421
func Int(v int) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:421
	_go_fuzz_dep_.CoverTab[108571]++
											p := new(int32)
											*p = int32(v)
											return p
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:424
	// _ = "end of CoverTab[108571]"
}

// Int64 is a helper routine that allocates a new int64 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:427
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:429
func Int64(v int64) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:429
	_go_fuzz_dep_.CoverTab[108572]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:430
	// _ = "end of CoverTab[108572]"
}

// Float32 is a helper routine that allocates a new float32 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:433
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:435
func Float32(v float32) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:435
	_go_fuzz_dep_.CoverTab[108573]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:436
	// _ = "end of CoverTab[108573]"
}

// Float64 is a helper routine that allocates a new float64 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:439
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:441
func Float64(v float64) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:441
	_go_fuzz_dep_.CoverTab[108574]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:442
	// _ = "end of CoverTab[108574]"
}

// Uint32 is a helper routine that allocates a new uint32 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:445
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:447
func Uint32(v uint32) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:447
	_go_fuzz_dep_.CoverTab[108575]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:448
	// _ = "end of CoverTab[108575]"
}

// Uint64 is a helper routine that allocates a new uint64 value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:451
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:453
func Uint64(v uint64) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:453
	_go_fuzz_dep_.CoverTab[108576]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:454
	// _ = "end of CoverTab[108576]"
}

// String is a helper routine that allocates a new string value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:457
// to store v and returns a pointer to it.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:459
func String(v string) *string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:459
	_go_fuzz_dep_.CoverTab[108577]++
											return &v
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:460
	// _ = "end of CoverTab[108577]"
}

// EnumName is a helper function to simplify printing protocol buffer enums
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:463
// by name.  Given an enum map and a value, it returns a useful string.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:465
func EnumName(m map[int32]string, v int32) string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:465
	_go_fuzz_dep_.CoverTab[108578]++
											s, ok := m[v]
											if ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:467
		_go_fuzz_dep_.CoverTab[108580]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:468
		// _ = "end of CoverTab[108580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:469
		_go_fuzz_dep_.CoverTab[108581]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:469
		// _ = "end of CoverTab[108581]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:469
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:469
	// _ = "end of CoverTab[108578]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:469
	_go_fuzz_dep_.CoverTab[108579]++
											return strconv.Itoa(int(v))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:470
	// _ = "end of CoverTab[108579]"
}

// UnmarshalJSONEnum is a helper function to simplify recovering enum int values
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:473
// from their JSON-encoded representation. Given a map from the enum's symbolic
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:473
// names to its int values, and a byte buffer containing the JSON-encoded
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:473
// value, it returns an int32 that can be cast to the enum type by the caller.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:473
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:473
// The function can deal with both JSON representations, numeric and symbolic.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:479
func UnmarshalJSONEnum(m map[string]int32, data []byte, enumName string) (int32, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:479
	_go_fuzz_dep_.CoverTab[108582]++
											if data[0] == '"' {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:480
		_go_fuzz_dep_.CoverTab[108585]++
		// New style: enums are strings.
		var repr string
		if err := json.Unmarshal(data, &repr); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:483
			_go_fuzz_dep_.CoverTab[108588]++
													return -1, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:484
			// _ = "end of CoverTab[108588]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:485
			_go_fuzz_dep_.CoverTab[108589]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:485
			// _ = "end of CoverTab[108589]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:485
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:485
		// _ = "end of CoverTab[108585]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:485
		_go_fuzz_dep_.CoverTab[108586]++
												val, ok := m[repr]
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:487
			_go_fuzz_dep_.CoverTab[108590]++
													return 0, fmt.Errorf("unrecognized enum %s value %q", enumName, repr)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:488
			// _ = "end of CoverTab[108590]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:489
			_go_fuzz_dep_.CoverTab[108591]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:489
			// _ = "end of CoverTab[108591]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:489
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:489
		// _ = "end of CoverTab[108586]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:489
		_go_fuzz_dep_.CoverTab[108587]++
												return val, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:490
		// _ = "end of CoverTab[108587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:491
		_go_fuzz_dep_.CoverTab[108592]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:491
		// _ = "end of CoverTab[108592]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:491
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:491
	// _ = "end of CoverTab[108582]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:491
	_go_fuzz_dep_.CoverTab[108583]++
	// Old style: enums are ints.
	var val int32
	if err := json.Unmarshal(data, &val); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:494
		_go_fuzz_dep_.CoverTab[108593]++
												return 0, fmt.Errorf("cannot unmarshal %#q into enum %s", data, enumName)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:495
		// _ = "end of CoverTab[108593]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:496
		_go_fuzz_dep_.CoverTab[108594]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:496
		// _ = "end of CoverTab[108594]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:496
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:496
	// _ = "end of CoverTab[108583]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:496
	_go_fuzz_dep_.CoverTab[108584]++
											return val, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:497
	// _ = "end of CoverTab[108584]"
}

// DebugPrint dumps the encoded data in b in a debugging format with a header
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:500
// including the string s. Used in testing but made available for general debugging.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:502
func (p *Buffer) DebugPrint(s string, b []byte) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:502
	_go_fuzz_dep_.CoverTab[108595]++
											var u uint64

											obuf := p.buf
											sindex := p.index
											p.buf = b
											p.index = 0
											depth := 0

											fmt.Printf("\n--- %s ---\n", s)

out:
	for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:514
		_go_fuzz_dep_.CoverTab[108598]++
												for i := 0; i < depth; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:515
			_go_fuzz_dep_.CoverTab[108602]++
													fmt.Print("  ")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:516
			// _ = "end of CoverTab[108602]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:517
		// _ = "end of CoverTab[108598]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:517
		_go_fuzz_dep_.CoverTab[108599]++

												index := p.index
												if index == len(p.buf) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:520
			_go_fuzz_dep_.CoverTab[108603]++
													break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:521
			// _ = "end of CoverTab[108603]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:522
			_go_fuzz_dep_.CoverTab[108604]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:522
			// _ = "end of CoverTab[108604]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:522
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:522
		// _ = "end of CoverTab[108599]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:522
		_go_fuzz_dep_.CoverTab[108600]++

												op, err := p.DecodeVarint()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:525
			_go_fuzz_dep_.CoverTab[108605]++
													fmt.Printf("%3d: fetching op err %v\n", index, err)
													break out
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:527
			// _ = "end of CoverTab[108605]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:528
			_go_fuzz_dep_.CoverTab[108606]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:528
			// _ = "end of CoverTab[108606]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:528
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:528
		// _ = "end of CoverTab[108600]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:528
		_go_fuzz_dep_.CoverTab[108601]++
												tag := op >> 3
												wire := op & 7

												switch wire {
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:533
			_go_fuzz_dep_.CoverTab[108607]++
													fmt.Printf("%3d: t=%3d unknown wire=%d\n",
				index, tag, wire)
													break out
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:536
			// _ = "end of CoverTab[108607]"

		case WireBytes:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:538
			_go_fuzz_dep_.CoverTab[108608]++
													var r []byte

													r, err = p.DecodeRawBytes(false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:542
				_go_fuzz_dep_.CoverTab[108619]++
														break out
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:543
				// _ = "end of CoverTab[108619]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:544
				_go_fuzz_dep_.CoverTab[108620]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:544
				// _ = "end of CoverTab[108620]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:544
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:544
			// _ = "end of CoverTab[108608]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:544
			_go_fuzz_dep_.CoverTab[108609]++
													fmt.Printf("%3d: t=%3d bytes [%d]", index, tag, len(r))
													if len(r) <= 6 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:546
				_go_fuzz_dep_.CoverTab[108621]++
														for i := 0; i < len(r); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:547
					_go_fuzz_dep_.CoverTab[108622]++
															fmt.Printf(" %.2x", r[i])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:548
					// _ = "end of CoverTab[108622]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:549
				// _ = "end of CoverTab[108621]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:550
				_go_fuzz_dep_.CoverTab[108623]++
														for i := 0; i < 3; i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:551
					_go_fuzz_dep_.CoverTab[108625]++
															fmt.Printf(" %.2x", r[i])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:552
					// _ = "end of CoverTab[108625]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:553
				// _ = "end of CoverTab[108623]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:553
				_go_fuzz_dep_.CoverTab[108624]++
														fmt.Printf(" ..")
														for i := len(r) - 3; i < len(r); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:555
					_go_fuzz_dep_.CoverTab[108626]++
															fmt.Printf(" %.2x", r[i])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:556
					// _ = "end of CoverTab[108626]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:557
				// _ = "end of CoverTab[108624]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:558
			// _ = "end of CoverTab[108609]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:558
			_go_fuzz_dep_.CoverTab[108610]++
													fmt.Printf("\n")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:559
			// _ = "end of CoverTab[108610]"

		case WireFixed32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:561
			_go_fuzz_dep_.CoverTab[108611]++
													u, err = p.DecodeFixed32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:563
				_go_fuzz_dep_.CoverTab[108627]++
														fmt.Printf("%3d: t=%3d fix32 err %v\n", index, tag, err)
														break out
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:565
				// _ = "end of CoverTab[108627]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:566
				_go_fuzz_dep_.CoverTab[108628]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:566
				// _ = "end of CoverTab[108628]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:566
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:566
			// _ = "end of CoverTab[108611]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:566
			_go_fuzz_dep_.CoverTab[108612]++
													fmt.Printf("%3d: t=%3d fix32 %d\n", index, tag, u)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:567
			// _ = "end of CoverTab[108612]"

		case WireFixed64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:569
			_go_fuzz_dep_.CoverTab[108613]++
													u, err = p.DecodeFixed64()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:571
				_go_fuzz_dep_.CoverTab[108629]++
														fmt.Printf("%3d: t=%3d fix64 err %v\n", index, tag, err)
														break out
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:573
				// _ = "end of CoverTab[108629]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:574
				_go_fuzz_dep_.CoverTab[108630]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:574
				// _ = "end of CoverTab[108630]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:574
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:574
			// _ = "end of CoverTab[108613]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:574
			_go_fuzz_dep_.CoverTab[108614]++
													fmt.Printf("%3d: t=%3d fix64 %d\n", index, tag, u)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:575
			// _ = "end of CoverTab[108614]"

		case WireVarint:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:577
			_go_fuzz_dep_.CoverTab[108615]++
													u, err = p.DecodeVarint()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:579
				_go_fuzz_dep_.CoverTab[108631]++
														fmt.Printf("%3d: t=%3d varint err %v\n", index, tag, err)
														break out
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:581
				// _ = "end of CoverTab[108631]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:582
				_go_fuzz_dep_.CoverTab[108632]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:582
				// _ = "end of CoverTab[108632]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:582
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:582
			// _ = "end of CoverTab[108615]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:582
			_go_fuzz_dep_.CoverTab[108616]++
													fmt.Printf("%3d: t=%3d varint %d\n", index, tag, u)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:583
			// _ = "end of CoverTab[108616]"

		case WireStartGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:585
			_go_fuzz_dep_.CoverTab[108617]++
													fmt.Printf("%3d: t=%3d start\n", index, tag)
													depth++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:587
			// _ = "end of CoverTab[108617]"

		case WireEndGroup:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:589
			_go_fuzz_dep_.CoverTab[108618]++
													depth--
													fmt.Printf("%3d: t=%3d end\n", index, tag)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:591
			// _ = "end of CoverTab[108618]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:592
		// _ = "end of CoverTab[108601]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:593
	// _ = "end of CoverTab[108595]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:593
	_go_fuzz_dep_.CoverTab[108596]++

											if depth != 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:595
		_go_fuzz_dep_.CoverTab[108633]++
												fmt.Printf("%3d: start-end not balanced %d\n", p.index, depth)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:596
		// _ = "end of CoverTab[108633]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:597
		_go_fuzz_dep_.CoverTab[108634]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:597
		// _ = "end of CoverTab[108634]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:597
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:597
	// _ = "end of CoverTab[108596]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:597
	_go_fuzz_dep_.CoverTab[108597]++
											fmt.Printf("\n")

											p.buf = obuf
											p.index = sindex
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:601
	// _ = "end of CoverTab[108597]"
}

// SetDefaults sets unset protocol buffer fields to their default values.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:604
// It only modifies fields that are both unset and have defined defaults.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:604
// It recursively sets default values in any non-nil sub-messages.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:607
func SetDefaults(pb Message) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:607
	_go_fuzz_dep_.CoverTab[108635]++
											setDefaults(reflect.ValueOf(pb), true, false)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:608
	// _ = "end of CoverTab[108635]"
}

// v is a struct.
func setDefaults(v reflect.Value, recur, zeros bool) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:612
	_go_fuzz_dep_.CoverTab[108636]++
											if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:613
		_go_fuzz_dep_.CoverTab[108640]++
												v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:614
		// _ = "end of CoverTab[108640]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:615
		_go_fuzz_dep_.CoverTab[108641]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:615
		// _ = "end of CoverTab[108641]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:615
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:615
	// _ = "end of CoverTab[108636]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:615
	_go_fuzz_dep_.CoverTab[108637]++

											defaultMu.RLock()
											dm, ok := defaults[v.Type()]
											defaultMu.RUnlock()
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:620
		_go_fuzz_dep_.CoverTab[108642]++
												dm = buildDefaultMessage(v.Type())
												defaultMu.Lock()
												defaults[v.Type()] = dm
												defaultMu.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:624
		// _ = "end of CoverTab[108642]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:625
		_go_fuzz_dep_.CoverTab[108643]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:625
		// _ = "end of CoverTab[108643]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:625
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:625
	// _ = "end of CoverTab[108637]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:625
	_go_fuzz_dep_.CoverTab[108638]++

											for _, sf := range dm.scalars {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:627
		_go_fuzz_dep_.CoverTab[108644]++
												f := v.Field(sf.index)
												if !f.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:629
			_go_fuzz_dep_.CoverTab[108647]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:631
			// _ = "end of CoverTab[108647]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:632
			_go_fuzz_dep_.CoverTab[108648]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:632
			// _ = "end of CoverTab[108648]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:632
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:632
		// _ = "end of CoverTab[108644]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:632
		_go_fuzz_dep_.CoverTab[108645]++
												dv := sf.value
												if dv == nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:634
			_go_fuzz_dep_.CoverTab[108649]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:634
			return !zeros
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:634
			// _ = "end of CoverTab[108649]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:634
		}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:634
			_go_fuzz_dep_.CoverTab[108650]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:636
			// _ = "end of CoverTab[108650]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:637
			_go_fuzz_dep_.CoverTab[108651]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:637
			// _ = "end of CoverTab[108651]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:637
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:637
		// _ = "end of CoverTab[108645]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:637
		_go_fuzz_dep_.CoverTab[108646]++
												fptr := f.Addr().Interface()

												switch sf.kind {
		case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:641
			_go_fuzz_dep_.CoverTab[108652]++
													b := new(bool)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:643
				_go_fuzz_dep_.CoverTab[108670]++
														*b = dv.(bool)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:644
				// _ = "end of CoverTab[108670]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:645
				_go_fuzz_dep_.CoverTab[108671]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:645
				// _ = "end of CoverTab[108671]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:645
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:645
			// _ = "end of CoverTab[108652]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:645
			_go_fuzz_dep_.CoverTab[108653]++
													*(fptr.(**bool)) = b
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:646
			// _ = "end of CoverTab[108653]"
		case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:647
			_go_fuzz_dep_.CoverTab[108654]++
													f := new(float32)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:649
				_go_fuzz_dep_.CoverTab[108672]++
														*f = dv.(float32)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:650
				// _ = "end of CoverTab[108672]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:651
				_go_fuzz_dep_.CoverTab[108673]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:651
				// _ = "end of CoverTab[108673]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:651
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:651
			// _ = "end of CoverTab[108654]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:651
			_go_fuzz_dep_.CoverTab[108655]++
													*(fptr.(**float32)) = f
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:652
			// _ = "end of CoverTab[108655]"
		case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:653
			_go_fuzz_dep_.CoverTab[108656]++
													f := new(float64)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:655
				_go_fuzz_dep_.CoverTab[108674]++
														*f = dv.(float64)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:656
				// _ = "end of CoverTab[108674]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:657
				_go_fuzz_dep_.CoverTab[108675]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:657
				// _ = "end of CoverTab[108675]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:657
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:657
			// _ = "end of CoverTab[108656]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:657
			_go_fuzz_dep_.CoverTab[108657]++
													*(fptr.(**float64)) = f
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:658
			// _ = "end of CoverTab[108657]"
		case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:659
			_go_fuzz_dep_.CoverTab[108658]++

													if ft := f.Type(); ft != int32PtrType {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:661
				_go_fuzz_dep_.CoverTab[108676]++

														f.Set(reflect.New(ft.Elem()))
														if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:664
					_go_fuzz_dep_.CoverTab[108677]++
															f.Elem().SetInt(int64(dv.(int32)))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:665
					// _ = "end of CoverTab[108677]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:666
					_go_fuzz_dep_.CoverTab[108678]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:666
					// _ = "end of CoverTab[108678]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:666
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:666
				// _ = "end of CoverTab[108676]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:667
				_go_fuzz_dep_.CoverTab[108679]++

														i := new(int32)
														if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:670
					_go_fuzz_dep_.CoverTab[108681]++
															*i = dv.(int32)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:671
					// _ = "end of CoverTab[108681]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:672
					_go_fuzz_dep_.CoverTab[108682]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:672
					// _ = "end of CoverTab[108682]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:672
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:672
				// _ = "end of CoverTab[108679]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:672
				_go_fuzz_dep_.CoverTab[108680]++
														*(fptr.(**int32)) = i
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:673
				// _ = "end of CoverTab[108680]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:674
			// _ = "end of CoverTab[108658]"
		case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:675
			_go_fuzz_dep_.CoverTab[108659]++
													i := new(int64)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:677
				_go_fuzz_dep_.CoverTab[108683]++
														*i = dv.(int64)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:678
				// _ = "end of CoverTab[108683]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:679
				_go_fuzz_dep_.CoverTab[108684]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:679
				// _ = "end of CoverTab[108684]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:679
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:679
			// _ = "end of CoverTab[108659]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:679
			_go_fuzz_dep_.CoverTab[108660]++
													*(fptr.(**int64)) = i
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:680
			// _ = "end of CoverTab[108660]"
		case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:681
			_go_fuzz_dep_.CoverTab[108661]++
													s := new(string)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:683
				_go_fuzz_dep_.CoverTab[108685]++
														*s = dv.(string)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:684
				// _ = "end of CoverTab[108685]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:685
				_go_fuzz_dep_.CoverTab[108686]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:685
				// _ = "end of CoverTab[108686]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:685
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:685
			// _ = "end of CoverTab[108661]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:685
			_go_fuzz_dep_.CoverTab[108662]++
													*(fptr.(**string)) = s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:686
			// _ = "end of CoverTab[108662]"
		case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:687
			_go_fuzz_dep_.CoverTab[108663]++
			// exceptional case: []byte
			var b []byte
			if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:690
				_go_fuzz_dep_.CoverTab[108687]++
														db := dv.([]byte)
														b = make([]byte, len(db))
														copy(b, db)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:693
				// _ = "end of CoverTab[108687]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:694
				_go_fuzz_dep_.CoverTab[108688]++
														b = []byte{}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:695
				// _ = "end of CoverTab[108688]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:696
			// _ = "end of CoverTab[108663]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:696
			_go_fuzz_dep_.CoverTab[108664]++
													*(fptr.(*[]byte)) = b
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:697
			// _ = "end of CoverTab[108664]"
		case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:698
			_go_fuzz_dep_.CoverTab[108665]++
													u := new(uint32)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:700
				_go_fuzz_dep_.CoverTab[108689]++
														*u = dv.(uint32)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:701
				// _ = "end of CoverTab[108689]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:702
				_go_fuzz_dep_.CoverTab[108690]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:702
				// _ = "end of CoverTab[108690]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:702
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:702
			// _ = "end of CoverTab[108665]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:702
			_go_fuzz_dep_.CoverTab[108666]++
													*(fptr.(**uint32)) = u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:703
			// _ = "end of CoverTab[108666]"
		case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:704
			_go_fuzz_dep_.CoverTab[108667]++
													u := new(uint64)
													if dv != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:706
				_go_fuzz_dep_.CoverTab[108691]++
														*u = dv.(uint64)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:707
				// _ = "end of CoverTab[108691]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:708
				_go_fuzz_dep_.CoverTab[108692]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:708
				// _ = "end of CoverTab[108692]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:708
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:708
			// _ = "end of CoverTab[108667]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:708
			_go_fuzz_dep_.CoverTab[108668]++
													*(fptr.(**uint64)) = u
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:709
			// _ = "end of CoverTab[108668]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:710
			_go_fuzz_dep_.CoverTab[108669]++
													log.Printf("proto: can't set default for field %v (sf.kind=%v)", f, sf.kind)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:711
			// _ = "end of CoverTab[108669]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:712
		// _ = "end of CoverTab[108646]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:713
	// _ = "end of CoverTab[108638]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:713
	_go_fuzz_dep_.CoverTab[108639]++

											for _, ni := range dm.nested {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:715
		_go_fuzz_dep_.CoverTab[108693]++
												f := v.Field(ni)

												switch f.Kind() {
		case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:719
			_go_fuzz_dep_.CoverTab[108694]++
													setDefaults(f, recur, zeros)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:720
			// _ = "end of CoverTab[108694]"

		case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:722
			_go_fuzz_dep_.CoverTab[108695]++
													if f.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:723
				_go_fuzz_dep_.CoverTab[108700]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:724
				// _ = "end of CoverTab[108700]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:725
				_go_fuzz_dep_.CoverTab[108701]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:725
				// _ = "end of CoverTab[108701]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:725
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:725
			// _ = "end of CoverTab[108695]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:725
			_go_fuzz_dep_.CoverTab[108696]++
													setDefaults(f, recur, zeros)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:726
			// _ = "end of CoverTab[108696]"

		case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:728
			_go_fuzz_dep_.CoverTab[108697]++
													for i := 0; i < f.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:729
				_go_fuzz_dep_.CoverTab[108702]++
														e := f.Index(i)
														if e.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:731
					_go_fuzz_dep_.CoverTab[108704]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:731
					return e.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:731
					// _ = "end of CoverTab[108704]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:731
				}() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:731
					_go_fuzz_dep_.CoverTab[108705]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:732
					// _ = "end of CoverTab[108705]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:733
					_go_fuzz_dep_.CoverTab[108706]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:733
					// _ = "end of CoverTab[108706]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:733
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:733
				// _ = "end of CoverTab[108702]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:733
				_go_fuzz_dep_.CoverTab[108703]++
														setDefaults(e, recur, zeros)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:734
				// _ = "end of CoverTab[108703]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:735
			// _ = "end of CoverTab[108697]"

		case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:737
			_go_fuzz_dep_.CoverTab[108698]++
													for _, k := range f.MapKeys() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:738
				_go_fuzz_dep_.CoverTab[108707]++
														e := f.MapIndex(k)
														if e.IsNil() {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:740
					_go_fuzz_dep_.CoverTab[108709]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:741
					// _ = "end of CoverTab[108709]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:742
					_go_fuzz_dep_.CoverTab[108710]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:742
					// _ = "end of CoverTab[108710]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:742
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:742
				// _ = "end of CoverTab[108707]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:742
				_go_fuzz_dep_.CoverTab[108708]++
														setDefaults(e, recur, zeros)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:743
				// _ = "end of CoverTab[108708]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:744
			// _ = "end of CoverTab[108698]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:744
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:744
			_go_fuzz_dep_.CoverTab[108699]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:744
			// _ = "end of CoverTab[108699]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:745
		// _ = "end of CoverTab[108693]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:746
	// _ = "end of CoverTab[108639]"
}

var (
	// defaults maps a protocol buffer struct type to a slice of the fields,
	// with its scalar fields set to their proto-declared non-zero default values.
	defaultMu	sync.RWMutex
	defaults	= make(map[reflect.Type]defaultMessage)

	int32PtrType	= reflect.TypeOf((*int32)(nil))
)

// defaultMessage represents information about the default values of a message.
type defaultMessage struct {
	scalars	[]scalarField
	nested	[]int	// struct field index of nested messages
}

type scalarField struct {
	index	int		// struct field index
	kind	reflect.Kind	// element type (the T in *T or []T)
	value	interface{}	// the proto-declared default value, or nil
}

// t is a struct type.
func buildDefaultMessage(t reflect.Type) (dm defaultMessage) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:771
	_go_fuzz_dep_.CoverTab[108711]++
											sprop := GetProperties(t)
											for _, prop := range sprop.Prop {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:773
		_go_fuzz_dep_.CoverTab[108713]++
												fi, ok := sprop.decoderTags.get(prop.Tag)
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:775
			_go_fuzz_dep_.CoverTab[108715]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:777
			// _ = "end of CoverTab[108715]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:778
			_go_fuzz_dep_.CoverTab[108716]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:778
			// _ = "end of CoverTab[108716]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:778
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:778
		// _ = "end of CoverTab[108713]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:778
		_go_fuzz_dep_.CoverTab[108714]++
												ft := t.Field(fi).Type

												sf, nested, err := fieldDefault(ft, prop)
												switch {
		case err != nil:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:783
			_go_fuzz_dep_.CoverTab[108717]++
													log.Print(err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:784
			// _ = "end of CoverTab[108717]"
		case nested:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:785
			_go_fuzz_dep_.CoverTab[108718]++
													dm.nested = append(dm.nested, fi)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:786
			// _ = "end of CoverTab[108718]"
		case sf != nil:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:787
			_go_fuzz_dep_.CoverTab[108719]++
													sf.index = fi
													dm.scalars = append(dm.scalars, *sf)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:789
			// _ = "end of CoverTab[108719]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:789
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:789
			_go_fuzz_dep_.CoverTab[108720]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:789
			// _ = "end of CoverTab[108720]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:790
		// _ = "end of CoverTab[108714]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:791
	// _ = "end of CoverTab[108711]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:791
	_go_fuzz_dep_.CoverTab[108712]++

											return dm
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:793
	// _ = "end of CoverTab[108712]"
}

// fieldDefault returns the scalarField for field type ft.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:796
// sf will be nil if the field can not have a default.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:796
// nestedMessage will be true if this is a nested message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:796
// Note that sf.index is not set on return.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:800
func fieldDefault(ft reflect.Type, prop *Properties) (sf *scalarField, nestedMessage bool, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:800
	_go_fuzz_dep_.CoverTab[108721]++
											var canHaveDefault bool
											switch ft.Kind() {
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:803
		_go_fuzz_dep_.CoverTab[108726]++
												nestedMessage = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:804
		// _ = "end of CoverTab[108726]"

	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:806
		_go_fuzz_dep_.CoverTab[108727]++
												if ft.Elem().Kind() == reflect.Struct {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:807
			_go_fuzz_dep_.CoverTab[108731]++
													nestedMessage = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:808
			// _ = "end of CoverTab[108731]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:809
			_go_fuzz_dep_.CoverTab[108732]++
													canHaveDefault = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:810
			// _ = "end of CoverTab[108732]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:811
		// _ = "end of CoverTab[108727]"

	case reflect.Slice:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:813
		_go_fuzz_dep_.CoverTab[108728]++
												switch ft.Elem().Kind() {
		case reflect.Ptr, reflect.Struct:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:815
			_go_fuzz_dep_.CoverTab[108733]++
													nestedMessage = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:816
			// _ = "end of CoverTab[108733]"
		case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:817
			_go_fuzz_dep_.CoverTab[108734]++
													canHaveDefault = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:818
			// _ = "end of CoverTab[108734]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:818
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:818
			_go_fuzz_dep_.CoverTab[108735]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:818
			// _ = "end of CoverTab[108735]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:819
		// _ = "end of CoverTab[108728]"

	case reflect.Map:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:821
		_go_fuzz_dep_.CoverTab[108729]++
												if ft.Elem().Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:822
			_go_fuzz_dep_.CoverTab[108736]++
													nestedMessage = true
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:823
			// _ = "end of CoverTab[108736]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
			_go_fuzz_dep_.CoverTab[108737]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
			// _ = "end of CoverTab[108737]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
		// _ = "end of CoverTab[108729]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
		_go_fuzz_dep_.CoverTab[108730]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:824
		// _ = "end of CoverTab[108730]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:825
	// _ = "end of CoverTab[108721]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:825
	_go_fuzz_dep_.CoverTab[108722]++

											if !canHaveDefault {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:827
		_go_fuzz_dep_.CoverTab[108738]++
												if nestedMessage {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:828
			_go_fuzz_dep_.CoverTab[108740]++
													return nil, true, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:829
			// _ = "end of CoverTab[108740]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:830
			_go_fuzz_dep_.CoverTab[108741]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:830
			// _ = "end of CoverTab[108741]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:830
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:830
		// _ = "end of CoverTab[108738]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:830
		_go_fuzz_dep_.CoverTab[108739]++
												return nil, false, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:831
		// _ = "end of CoverTab[108739]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:832
		_go_fuzz_dep_.CoverTab[108742]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:832
		// _ = "end of CoverTab[108742]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:832
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:832
	// _ = "end of CoverTab[108722]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:832
	_go_fuzz_dep_.CoverTab[108723]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:835
	sf = &scalarField{kind: ft.Elem().Kind()}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:838
	if !prop.HasDefault {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:838
		_go_fuzz_dep_.CoverTab[108743]++
												return sf, false, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:839
		// _ = "end of CoverTab[108743]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:840
		_go_fuzz_dep_.CoverTab[108744]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:840
		// _ = "end of CoverTab[108744]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:840
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:840
	// _ = "end of CoverTab[108723]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:840
	_go_fuzz_dep_.CoverTab[108724]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:843
	switch ft.Elem().Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:844
		_go_fuzz_dep_.CoverTab[108745]++
												x, err := strconv.ParseBool(prop.Default)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:846
			_go_fuzz_dep_.CoverTab[108762]++
													return nil, false, fmt.Errorf("proto: bad default bool %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:847
			// _ = "end of CoverTab[108762]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:848
			_go_fuzz_dep_.CoverTab[108763]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:848
			// _ = "end of CoverTab[108763]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:848
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:848
		// _ = "end of CoverTab[108745]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:848
		_go_fuzz_dep_.CoverTab[108746]++
												sf.value = x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:849
		// _ = "end of CoverTab[108746]"
	case reflect.Float32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:850
		_go_fuzz_dep_.CoverTab[108747]++
												x, err := strconv.ParseFloat(prop.Default, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:852
			_go_fuzz_dep_.CoverTab[108764]++
													return nil, false, fmt.Errorf("proto: bad default float32 %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:853
			// _ = "end of CoverTab[108764]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:854
			_go_fuzz_dep_.CoverTab[108765]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:854
			// _ = "end of CoverTab[108765]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:854
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:854
		// _ = "end of CoverTab[108747]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:854
		_go_fuzz_dep_.CoverTab[108748]++
												sf.value = float32(x)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:855
		// _ = "end of CoverTab[108748]"
	case reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:856
		_go_fuzz_dep_.CoverTab[108749]++
												x, err := strconv.ParseFloat(prop.Default, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:858
			_go_fuzz_dep_.CoverTab[108766]++
													return nil, false, fmt.Errorf("proto: bad default float64 %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:859
			// _ = "end of CoverTab[108766]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:860
			_go_fuzz_dep_.CoverTab[108767]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:860
			// _ = "end of CoverTab[108767]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:860
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:860
		// _ = "end of CoverTab[108749]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:860
		_go_fuzz_dep_.CoverTab[108750]++
												sf.value = x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:861
		// _ = "end of CoverTab[108750]"
	case reflect.Int32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:862
		_go_fuzz_dep_.CoverTab[108751]++
												x, err := strconv.ParseInt(prop.Default, 10, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:864
			_go_fuzz_dep_.CoverTab[108768]++
													return nil, false, fmt.Errorf("proto: bad default int32 %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:865
			// _ = "end of CoverTab[108768]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:866
			_go_fuzz_dep_.CoverTab[108769]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:866
			// _ = "end of CoverTab[108769]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:866
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:866
		// _ = "end of CoverTab[108751]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:866
		_go_fuzz_dep_.CoverTab[108752]++
												sf.value = int32(x)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:867
		// _ = "end of CoverTab[108752]"
	case reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:868
		_go_fuzz_dep_.CoverTab[108753]++
												x, err := strconv.ParseInt(prop.Default, 10, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:870
			_go_fuzz_dep_.CoverTab[108770]++
													return nil, false, fmt.Errorf("proto: bad default int64 %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:871
			// _ = "end of CoverTab[108770]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:872
			_go_fuzz_dep_.CoverTab[108771]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:872
			// _ = "end of CoverTab[108771]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:872
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:872
		// _ = "end of CoverTab[108753]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:872
		_go_fuzz_dep_.CoverTab[108754]++
												sf.value = x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:873
		// _ = "end of CoverTab[108754]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:874
		_go_fuzz_dep_.CoverTab[108755]++
												sf.value = prop.Default
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:875
		// _ = "end of CoverTab[108755]"
	case reflect.Uint8:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:876
		_go_fuzz_dep_.CoverTab[108756]++

												sf.value = []byte(prop.Default)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:878
		// _ = "end of CoverTab[108756]"
	case reflect.Uint32:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:879
		_go_fuzz_dep_.CoverTab[108757]++
												x, err := strconv.ParseUint(prop.Default, 10, 32)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:881
			_go_fuzz_dep_.CoverTab[108772]++
													return nil, false, fmt.Errorf("proto: bad default uint32 %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:882
			// _ = "end of CoverTab[108772]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:883
			_go_fuzz_dep_.CoverTab[108773]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:883
			// _ = "end of CoverTab[108773]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:883
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:883
		// _ = "end of CoverTab[108757]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:883
		_go_fuzz_dep_.CoverTab[108758]++
												sf.value = uint32(x)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:884
		// _ = "end of CoverTab[108758]"
	case reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:885
		_go_fuzz_dep_.CoverTab[108759]++
												x, err := strconv.ParseUint(prop.Default, 10, 64)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:887
			_go_fuzz_dep_.CoverTab[108774]++
													return nil, false, fmt.Errorf("proto: bad default uint64 %q: %v", prop.Default, err)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:888
			// _ = "end of CoverTab[108774]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:889
			_go_fuzz_dep_.CoverTab[108775]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:889
			// _ = "end of CoverTab[108775]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:889
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:889
		// _ = "end of CoverTab[108759]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:889
		_go_fuzz_dep_.CoverTab[108760]++
												sf.value = x
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:890
		// _ = "end of CoverTab[108760]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:891
		_go_fuzz_dep_.CoverTab[108761]++
												return nil, false, fmt.Errorf("proto: unhandled def kind %v", ft.Elem().Kind())
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:892
		// _ = "end of CoverTab[108761]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:893
	// _ = "end of CoverTab[108724]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:893
	_go_fuzz_dep_.CoverTab[108725]++

											return sf, false, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:895
	// _ = "end of CoverTab[108725]"
}

// mapKeys returns a sort.Interface to be used for sorting the map keys.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:898
// Map fields may have key types of non-float scalars, strings and enums.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:900
func mapKeys(vs []reflect.Value) sort.Interface {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:900
	_go_fuzz_dep_.CoverTab[108776]++
											s := mapKeySorter{vs: vs}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:904
	if len(vs) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:904
		_go_fuzz_dep_.CoverTab[108779]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:905
		// _ = "end of CoverTab[108779]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:906
		_go_fuzz_dep_.CoverTab[108780]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:906
		// _ = "end of CoverTab[108780]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:906
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:906
	// _ = "end of CoverTab[108776]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:906
	_go_fuzz_dep_.CoverTab[108777]++
											switch vs[0].Kind() {
	case reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:908
		_go_fuzz_dep_.CoverTab[108781]++
												s.less = func(a, b reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:909
			_go_fuzz_dep_.CoverTab[108786]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:909
			return a.Int() < b.Int()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:909
			// _ = "end of CoverTab[108786]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:909
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:909
		// _ = "end of CoverTab[108781]"
	case reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:910
		_go_fuzz_dep_.CoverTab[108782]++
												s.less = func(a, b reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:911
			_go_fuzz_dep_.CoverTab[108787]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:911
			return a.Uint() < b.Uint()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:911
			// _ = "end of CoverTab[108787]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:911
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:911
		// _ = "end of CoverTab[108782]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:912
		_go_fuzz_dep_.CoverTab[108783]++
												s.less = func(a, b reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
			_go_fuzz_dep_.CoverTab[108788]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
			return !a.Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
				_go_fuzz_dep_.CoverTab[108789]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
				return b.Bool()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
				// _ = "end of CoverTab[108789]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
			}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
			// _ = "end of CoverTab[108788]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:913
		// _ = "end of CoverTab[108783]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:914
		_go_fuzz_dep_.CoverTab[108784]++
												s.less = func(a, b reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:915
			_go_fuzz_dep_.CoverTab[108790]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:915
			return a.String() < b.String()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:915
			// _ = "end of CoverTab[108790]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:915
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:915
		// _ = "end of CoverTab[108784]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:916
		_go_fuzz_dep_.CoverTab[108785]++
												panic(fmt.Sprintf("unsupported map key type: %v", vs[0].Kind()))
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:917
		// _ = "end of CoverTab[108785]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:918
	// _ = "end of CoverTab[108777]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:918
	_go_fuzz_dep_.CoverTab[108778]++

											return s
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:920
	// _ = "end of CoverTab[108778]"
}

type mapKeySorter struct {
	vs	[]reflect.Value
	less	func(a, b reflect.Value) bool
}

func (s mapKeySorter) Len() int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:928
	_go_fuzz_dep_.CoverTab[108791]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:928
	return len(s.vs)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:928
	// _ = "end of CoverTab[108791]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:928
}
func (s mapKeySorter) Swap(i, j int) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:929
	_go_fuzz_dep_.CoverTab[108792]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:929
	s.vs[i], s.vs[j] = s.vs[j], s.vs[i]
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:929
	// _ = "end of CoverTab[108792]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:929
}
func (s mapKeySorter) Less(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:930
	_go_fuzz_dep_.CoverTab[108793]++
											return s.less(s.vs[i], s.vs[j])
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:931
	// _ = "end of CoverTab[108793]"
}

// isProto3Zero reports whether v is a zero proto3 value.
func isProto3Zero(v reflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:935
	_go_fuzz_dep_.CoverTab[108794]++
											switch v.Kind() {
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:937
		_go_fuzz_dep_.CoverTab[108796]++
												return !v.Bool()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:938
		// _ = "end of CoverTab[108796]"
	case reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:939
		_go_fuzz_dep_.CoverTab[108797]++
												return v.Int() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:940
		// _ = "end of CoverTab[108797]"
	case reflect.Uint32, reflect.Uint64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:941
		_go_fuzz_dep_.CoverTab[108798]++
												return v.Uint() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:942
		// _ = "end of CoverTab[108798]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:943
		_go_fuzz_dep_.CoverTab[108799]++
												return v.Float() == 0
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:944
		// _ = "end of CoverTab[108799]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:945
		_go_fuzz_dep_.CoverTab[108800]++
												return v.String() == ""
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:946
		// _ = "end of CoverTab[108800]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:946
	default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:946
		_go_fuzz_dep_.CoverTab[108801]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:946
		// _ = "end of CoverTab[108801]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:947
	// _ = "end of CoverTab[108794]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:947
	_go_fuzz_dep_.CoverTab[108795]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:948
	// _ = "end of CoverTab[108795]"
}

const (
	// ProtoPackageIsVersion3 is referenced from generated protocol buffer files
	// to assert that that code is compatible with this version of the proto package.
	GoGoProtoPackageIsVersion3	= true

	// ProtoPackageIsVersion2 is referenced from generated protocol buffer files
	// to assert that that code is compatible with this version of the proto package.
	GoGoProtoPackageIsVersion2	= true

	// ProtoPackageIsVersion1 is referenced from generated protocol buffer files
	// to assert that that code is compatible with this version of the proto package.
	GoGoProtoPackageIsVersion1	= true
)

// InternalMessageInfo is a type used internally by generated .pb.go files.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:965
// This type is not intended to be used by non-generated code.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:965
// This type is not subject to any compatibility guarantee.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:968
type InternalMessageInfo struct {
	marshal		*marshalInfo
	unmarshal	*unmarshalInfo
	merge		*mergeInfo
	discard		*discardInfo
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:973
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/lib.go:973
var _ = _go_fuzz_dep_.CoverTab
