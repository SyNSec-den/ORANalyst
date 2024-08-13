// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
package protojson

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:5
)

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/internal/encoding/json"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type marshalFunc func(encoder, protoreflect.Message) error

// wellKnownTypeMarshaler returns a marshal function if the message type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:25
// has specialized serialization behavior. It returns nil otherwise.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:27
func wellKnownTypeMarshaler(name protoreflect.FullName) marshalFunc {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:27
	_go_fuzz_dep_.CoverTab[66350]++
														if name.Parent() == genid.GoogleProtobuf_package {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:28
		_go_fuzz_dep_.CoverTab[66352]++
															switch name.Name() {
		case genid.Any_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:30
			_go_fuzz_dep_.CoverTab[66353]++
																return encoder.marshalAny
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:31
			// _ = "end of CoverTab[66353]"
		case genid.Timestamp_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:32
			_go_fuzz_dep_.CoverTab[66354]++
																return encoder.marshalTimestamp
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:33
			// _ = "end of CoverTab[66354]"
		case genid.Duration_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:34
			_go_fuzz_dep_.CoverTab[66355]++
																return encoder.marshalDuration
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:35
			// _ = "end of CoverTab[66355]"
		case genid.BoolValue_message_name,
			genid.Int32Value_message_name,
			genid.Int64Value_message_name,
			genid.UInt32Value_message_name,
			genid.UInt64Value_message_name,
			genid.FloatValue_message_name,
			genid.DoubleValue_message_name,
			genid.StringValue_message_name,
			genid.BytesValue_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:44
			_go_fuzz_dep_.CoverTab[66356]++
																return encoder.marshalWrapperType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:45
			// _ = "end of CoverTab[66356]"
		case genid.Struct_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:46
			_go_fuzz_dep_.CoverTab[66357]++
																return encoder.marshalStruct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:47
			// _ = "end of CoverTab[66357]"
		case genid.ListValue_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:48
			_go_fuzz_dep_.CoverTab[66358]++
																return encoder.marshalListValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:49
			// _ = "end of CoverTab[66358]"
		case genid.Value_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:50
			_go_fuzz_dep_.CoverTab[66359]++
																return encoder.marshalKnownValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:51
			// _ = "end of CoverTab[66359]"
		case genid.FieldMask_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:52
			_go_fuzz_dep_.CoverTab[66360]++
																return encoder.marshalFieldMask
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:53
			// _ = "end of CoverTab[66360]"
		case genid.Empty_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:54
			_go_fuzz_dep_.CoverTab[66361]++
																return encoder.marshalEmpty
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:55
			// _ = "end of CoverTab[66361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:55
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:55
			_go_fuzz_dep_.CoverTab[66362]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:55
			// _ = "end of CoverTab[66362]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:56
		// _ = "end of CoverTab[66352]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:57
		_go_fuzz_dep_.CoverTab[66363]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:57
		// _ = "end of CoverTab[66363]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:57
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:57
	// _ = "end of CoverTab[66350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:57
	_go_fuzz_dep_.CoverTab[66351]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:58
	// _ = "end of CoverTab[66351]"
}

type unmarshalFunc func(decoder, protoreflect.Message) error

// wellKnownTypeUnmarshaler returns a unmarshal function if the message type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:63
// has specialized serialization behavior. It returns nil otherwise.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:65
func wellKnownTypeUnmarshaler(name protoreflect.FullName) unmarshalFunc {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:65
	_go_fuzz_dep_.CoverTab[66364]++
														if name.Parent() == genid.GoogleProtobuf_package {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:66
		_go_fuzz_dep_.CoverTab[66366]++
															switch name.Name() {
		case genid.Any_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:68
			_go_fuzz_dep_.CoverTab[66367]++
																return decoder.unmarshalAny
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:69
			// _ = "end of CoverTab[66367]"
		case genid.Timestamp_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:70
			_go_fuzz_dep_.CoverTab[66368]++
																return decoder.unmarshalTimestamp
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:71
			// _ = "end of CoverTab[66368]"
		case genid.Duration_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:72
			_go_fuzz_dep_.CoverTab[66369]++
																return decoder.unmarshalDuration
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:73
			// _ = "end of CoverTab[66369]"
		case genid.BoolValue_message_name,
			genid.Int32Value_message_name,
			genid.Int64Value_message_name,
			genid.UInt32Value_message_name,
			genid.UInt64Value_message_name,
			genid.FloatValue_message_name,
			genid.DoubleValue_message_name,
			genid.StringValue_message_name,
			genid.BytesValue_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:82
			_go_fuzz_dep_.CoverTab[66370]++
																return decoder.unmarshalWrapperType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:83
			// _ = "end of CoverTab[66370]"
		case genid.Struct_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:84
			_go_fuzz_dep_.CoverTab[66371]++
																return decoder.unmarshalStruct
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:85
			// _ = "end of CoverTab[66371]"
		case genid.ListValue_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:86
			_go_fuzz_dep_.CoverTab[66372]++
																return decoder.unmarshalListValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:87
			// _ = "end of CoverTab[66372]"
		case genid.Value_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:88
			_go_fuzz_dep_.CoverTab[66373]++
																return decoder.unmarshalKnownValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:89
			// _ = "end of CoverTab[66373]"
		case genid.FieldMask_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:90
			_go_fuzz_dep_.CoverTab[66374]++
																return decoder.unmarshalFieldMask
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:91
			// _ = "end of CoverTab[66374]"
		case genid.Empty_message_name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:92
			_go_fuzz_dep_.CoverTab[66375]++
																return decoder.unmarshalEmpty
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:93
			// _ = "end of CoverTab[66375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:93
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:93
			_go_fuzz_dep_.CoverTab[66376]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:93
			// _ = "end of CoverTab[66376]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:94
		// _ = "end of CoverTab[66366]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:95
		_go_fuzz_dep_.CoverTab[66377]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:95
		// _ = "end of CoverTab[66377]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:95
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:95
	// _ = "end of CoverTab[66364]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:95
	_go_fuzz_dep_.CoverTab[66365]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:96
	// _ = "end of CoverTab[66365]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:105
func (e encoder) marshalAny(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:105
	_go_fuzz_dep_.CoverTab[66378]++
															fds := m.Descriptor().Fields()
															fdType := fds.ByNumber(genid.Any_TypeUrl_field_number)
															fdValue := fds.ByNumber(genid.Any_Value_field_number)

															if !m.Has(fdType) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:110
		_go_fuzz_dep_.CoverTab[66384]++
																if !m.Has(fdValue) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:111
			_go_fuzz_dep_.CoverTab[66385]++

																	e.StartObject()
																	e.EndObject()
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:115
			// _ = "end of CoverTab[66385]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:116
			_go_fuzz_dep_.CoverTab[66386]++

																	return errors.New("%s: %v is not set", genid.Any_message_fullname, genid.Any_TypeUrl_field_name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:118
			// _ = "end of CoverTab[66386]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:119
		// _ = "end of CoverTab[66384]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:120
		_go_fuzz_dep_.CoverTab[66387]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:120
		// _ = "end of CoverTab[66387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:120
	// _ = "end of CoverTab[66378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:120
	_go_fuzz_dep_.CoverTab[66379]++

															typeVal := m.Get(fdType)
															valueVal := m.Get(fdValue)

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:126
	typeURL := typeVal.String()
	emt, err := e.opts.Resolver.FindMessageByURL(typeURL)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:128
		_go_fuzz_dep_.CoverTab[66388]++
																return errors.New("%s: unable to resolve %q: %v", genid.Any_message_fullname, typeURL, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:129
		// _ = "end of CoverTab[66388]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:130
		_go_fuzz_dep_.CoverTab[66389]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:130
		// _ = "end of CoverTab[66389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:130
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:130
	// _ = "end of CoverTab[66379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:130
	_go_fuzz_dep_.CoverTab[66380]++

															em := emt.New()
															err = proto.UnmarshalOptions{
		AllowPartial:	true,
		Resolver:	e.opts.Resolver,
	}.Unmarshal(valueVal.Bytes(), em.Interface())
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:137
		_go_fuzz_dep_.CoverTab[66390]++
																return errors.New("%s: unable to unmarshal %q: %v", genid.Any_message_fullname, typeURL, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:138
		// _ = "end of CoverTab[66390]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:139
		_go_fuzz_dep_.CoverTab[66391]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:139
		// _ = "end of CoverTab[66391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:139
	// _ = "end of CoverTab[66380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:139
	_go_fuzz_dep_.CoverTab[66381]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:144
	if marshal := wellKnownTypeMarshaler(emt.Descriptor().FullName()); marshal != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:144
		_go_fuzz_dep_.CoverTab[66392]++
																e.StartObject()
																defer e.EndObject()

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:149
		e.WriteName("@type")
		if err := e.WriteString(typeURL); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:150
			_go_fuzz_dep_.CoverTab[66394]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:151
			// _ = "end of CoverTab[66394]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:152
			_go_fuzz_dep_.CoverTab[66395]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:152
			// _ = "end of CoverTab[66395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:152
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:152
		// _ = "end of CoverTab[66392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:152
		_go_fuzz_dep_.CoverTab[66393]++

																e.WriteName("value")
																return marshal(e, em)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:155
		// _ = "end of CoverTab[66393]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:156
		_go_fuzz_dep_.CoverTab[66396]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:156
		// _ = "end of CoverTab[66396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:156
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:156
	// _ = "end of CoverTab[66381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:156
	_go_fuzz_dep_.CoverTab[66382]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:159
	if err := e.marshalMessage(em, typeURL); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:159
		_go_fuzz_dep_.CoverTab[66397]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:160
		// _ = "end of CoverTab[66397]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:161
		_go_fuzz_dep_.CoverTab[66398]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:161
		// _ = "end of CoverTab[66398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:161
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:161
	// _ = "end of CoverTab[66382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:161
	_go_fuzz_dep_.CoverTab[66383]++

															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:163
	// _ = "end of CoverTab[66383]"
}

func (d decoder) unmarshalAny(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:166
	_go_fuzz_dep_.CoverTab[66399]++

															start, err := d.Peek()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:169
		_go_fuzz_dep_.CoverTab[66406]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:170
		// _ = "end of CoverTab[66406]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:171
		_go_fuzz_dep_.CoverTab[66407]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:171
		// _ = "end of CoverTab[66407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:171
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:171
	// _ = "end of CoverTab[66399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:171
	_go_fuzz_dep_.CoverTab[66400]++
															if start.Kind() != json.ObjectOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:172
		_go_fuzz_dep_.CoverTab[66408]++
																return d.unexpectedTokenError(start)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:173
		// _ = "end of CoverTab[66408]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:174
		_go_fuzz_dep_.CoverTab[66409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:174
		// _ = "end of CoverTab[66409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:174
	// _ = "end of CoverTab[66400]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:174
	_go_fuzz_dep_.CoverTab[66401]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:179
	dec := decoder{d.Clone(), UnmarshalOptions{}}
	tok, err := findTypeURL(dec)
	switch err {
	case errEmptyObject:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:182
		_go_fuzz_dep_.CoverTab[66410]++

																d.Read()
																d.Read()
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:186
		// _ = "end of CoverTab[66410]"

	case errMissingType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:188
		_go_fuzz_dep_.CoverTab[66411]++
																if d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:189
			_go_fuzz_dep_.CoverTab[66414]++

																	return d.skipJSONValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:191
			// _ = "end of CoverTab[66414]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:192
			_go_fuzz_dep_.CoverTab[66415]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:192
			// _ = "end of CoverTab[66415]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:192
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:192
		// _ = "end of CoverTab[66411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:192
		_go_fuzz_dep_.CoverTab[66412]++

																return d.newError(start.Pos(), err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:194
		// _ = "end of CoverTab[66412]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:196
		_go_fuzz_dep_.CoverTab[66413]++
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:197
			_go_fuzz_dep_.CoverTab[66416]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:198
			// _ = "end of CoverTab[66416]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:199
			_go_fuzz_dep_.CoverTab[66417]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:199
			// _ = "end of CoverTab[66417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:199
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:199
		// _ = "end of CoverTab[66413]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:200
	// _ = "end of CoverTab[66401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:200
	_go_fuzz_dep_.CoverTab[66402]++

															typeURL := tok.ParsedString()
															emt, err := d.opts.Resolver.FindMessageByURL(typeURL)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:204
		_go_fuzz_dep_.CoverTab[66418]++
																return d.newError(tok.Pos(), "unable to resolve %v: %q", tok.RawString(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:205
		// _ = "end of CoverTab[66418]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:206
		_go_fuzz_dep_.CoverTab[66419]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:206
		// _ = "end of CoverTab[66419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:206
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:206
	// _ = "end of CoverTab[66402]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:206
	_go_fuzz_dep_.CoverTab[66403]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:209
	em := emt.New()
	if unmarshal := wellKnownTypeUnmarshaler(emt.Descriptor().FullName()); unmarshal != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:210
		_go_fuzz_dep_.CoverTab[66420]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:213
		if err := d.unmarshalAnyValue(unmarshal, em); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:213
			_go_fuzz_dep_.CoverTab[66421]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:214
			// _ = "end of CoverTab[66421]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:215
			_go_fuzz_dep_.CoverTab[66422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:215
			// _ = "end of CoverTab[66422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:215
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:215
		// _ = "end of CoverTab[66420]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:216
		_go_fuzz_dep_.CoverTab[66423]++

																if err := d.unmarshalMessage(em, true); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:218
			_go_fuzz_dep_.CoverTab[66424]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:219
			// _ = "end of CoverTab[66424]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:220
			_go_fuzz_dep_.CoverTab[66425]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:220
			// _ = "end of CoverTab[66425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:220
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:220
		// _ = "end of CoverTab[66423]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:221
	// _ = "end of CoverTab[66403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:221
	_go_fuzz_dep_.CoverTab[66404]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:224
	b, err := proto.MarshalOptions{
		AllowPartial:	true,
		Deterministic:	true,
	}.Marshal(em.Interface())
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:228
		_go_fuzz_dep_.CoverTab[66426]++
																return d.newError(start.Pos(), "error in marshaling Any.value field: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:229
		// _ = "end of CoverTab[66426]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:230
		_go_fuzz_dep_.CoverTab[66427]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:230
		// _ = "end of CoverTab[66427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:230
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:230
	// _ = "end of CoverTab[66404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:230
	_go_fuzz_dep_.CoverTab[66405]++

															fds := m.Descriptor().Fields()
															fdType := fds.ByNumber(genid.Any_TypeUrl_field_number)
															fdValue := fds.ByNumber(genid.Any_Value_field_number)

															m.Set(fdType, protoreflect.ValueOfString(typeURL))
															m.Set(fdValue, protoreflect.ValueOfBytes(b))
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:238
	// _ = "end of CoverTab[66405]"
}

var errEmptyObject = fmt.Errorf(`empty object`)
var errMissingType = fmt.Errorf(`missing "@type" field`)

// findTypeURL returns the token for the "@type" field value from the given
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:244
// JSON bytes. It is expected that the given bytes start with json.ObjectOpen.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:244
// It returns errEmptyObject if the JSON object is empty or errMissingType if
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:244
// @type field does not exist. It returns other error if the @type field is not
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:244
// valid or other decoding issues.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:249
func findTypeURL(d decoder) (json.Token, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:249
	_go_fuzz_dep_.CoverTab[66428]++
															var typeURL string
															var typeTok json.Token
															numFields := 0

															d.Read()

Loop:
	for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:257
		_go_fuzz_dep_.CoverTab[66430]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:259
			_go_fuzz_dep_.CoverTab[66432]++
																	return json.Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:260
			// _ = "end of CoverTab[66432]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:261
			_go_fuzz_dep_.CoverTab[66433]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:261
			// _ = "end of CoverTab[66433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:261
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:261
		// _ = "end of CoverTab[66430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:261
		_go_fuzz_dep_.CoverTab[66431]++

																switch tok.Kind() {
		case json.ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:264
			_go_fuzz_dep_.CoverTab[66434]++
																	if typeURL == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:265
				_go_fuzz_dep_.CoverTab[66443]++

																		if numFields > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:267
					_go_fuzz_dep_.CoverTab[66445]++
																			return json.Token{}, errMissingType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:268
					// _ = "end of CoverTab[66445]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:269
					_go_fuzz_dep_.CoverTab[66446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:269
					// _ = "end of CoverTab[66446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:269
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:269
				// _ = "end of CoverTab[66443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:269
				_go_fuzz_dep_.CoverTab[66444]++
																		return json.Token{}, errEmptyObject
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:270
				// _ = "end of CoverTab[66444]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:271
				_go_fuzz_dep_.CoverTab[66447]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:271
				// _ = "end of CoverTab[66447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:271
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:271
			// _ = "end of CoverTab[66434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:271
			_go_fuzz_dep_.CoverTab[66435]++
																	break Loop
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:272
			// _ = "end of CoverTab[66435]"

		case json.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:274
			_go_fuzz_dep_.CoverTab[66436]++
																	numFields++
																	if tok.Name() != "@type" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:276
				_go_fuzz_dep_.CoverTab[66448]++

																		if err := d.skipJSONValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:278
					_go_fuzz_dep_.CoverTab[66450]++
																			return json.Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:279
					// _ = "end of CoverTab[66450]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:280
					_go_fuzz_dep_.CoverTab[66451]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:280
					// _ = "end of CoverTab[66451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:280
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:280
				// _ = "end of CoverTab[66448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:280
				_go_fuzz_dep_.CoverTab[66449]++
																		continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:281
				// _ = "end of CoverTab[66449]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:282
				_go_fuzz_dep_.CoverTab[66452]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:282
				// _ = "end of CoverTab[66452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:282
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:282
			// _ = "end of CoverTab[66436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:282
			_go_fuzz_dep_.CoverTab[66437]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:285
			if typeURL != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:285
				_go_fuzz_dep_.CoverTab[66453]++
																		return json.Token{}, d.newError(tok.Pos(), `duplicate "@type" field`)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:286
				// _ = "end of CoverTab[66453]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:287
				_go_fuzz_dep_.CoverTab[66454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:287
				// _ = "end of CoverTab[66454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:287
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:287
			// _ = "end of CoverTab[66437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:287
			_go_fuzz_dep_.CoverTab[66438]++

																	tok, err := d.Read()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:290
				_go_fuzz_dep_.CoverTab[66455]++
																		return json.Token{}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:291
				// _ = "end of CoverTab[66455]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:292
				_go_fuzz_dep_.CoverTab[66456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:292
				// _ = "end of CoverTab[66456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:292
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:292
			// _ = "end of CoverTab[66438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:292
			_go_fuzz_dep_.CoverTab[66439]++
																	if tok.Kind() != json.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:293
				_go_fuzz_dep_.CoverTab[66457]++
																		return json.Token{}, d.newError(tok.Pos(), `@type field value is not a string: %v`, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:294
				// _ = "end of CoverTab[66457]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:295
				_go_fuzz_dep_.CoverTab[66458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:295
				// _ = "end of CoverTab[66458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:295
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:295
			// _ = "end of CoverTab[66439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:295
			_go_fuzz_dep_.CoverTab[66440]++
																	typeURL = tok.ParsedString()
																	if typeURL == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:297
				_go_fuzz_dep_.CoverTab[66459]++
																		return json.Token{}, d.newError(tok.Pos(), `@type field contains empty value`)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:298
				// _ = "end of CoverTab[66459]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:299
				_go_fuzz_dep_.CoverTab[66460]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:299
				// _ = "end of CoverTab[66460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:299
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:299
			// _ = "end of CoverTab[66440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:299
			_go_fuzz_dep_.CoverTab[66441]++
																	typeTok = tok
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:300
			// _ = "end of CoverTab[66441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:300
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:300
			_go_fuzz_dep_.CoverTab[66442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:300
			// _ = "end of CoverTab[66442]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:301
		// _ = "end of CoverTab[66431]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:302
	// _ = "end of CoverTab[66428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:302
	_go_fuzz_dep_.CoverTab[66429]++

															return typeTok, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:304
	// _ = "end of CoverTab[66429]"
}

// skipJSONValue parses a JSON value (null, boolean, string, number, object and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:307
// array) in order to advance the read to the next JSON value. It relies on
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:307
// the decoder returning an error if the types are not in valid sequence.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:310
func (d decoder) skipJSONValue() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:310
	_go_fuzz_dep_.CoverTab[66461]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:312
		_go_fuzz_dep_.CoverTab[66464]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:313
		// _ = "end of CoverTab[66464]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:314
		_go_fuzz_dep_.CoverTab[66465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:314
		// _ = "end of CoverTab[66465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:314
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:314
	// _ = "end of CoverTab[66461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:314
	_go_fuzz_dep_.CoverTab[66462]++

															switch tok.Kind() {
	case json.ObjectOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:317
		_go_fuzz_dep_.CoverTab[66466]++
																for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:318
			_go_fuzz_dep_.CoverTab[66469]++
																	tok, err := d.Read()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:320
				_go_fuzz_dep_.CoverTab[66471]++
																		return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:321
				// _ = "end of CoverTab[66471]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:322
				_go_fuzz_dep_.CoverTab[66472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:322
				// _ = "end of CoverTab[66472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:322
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:322
			// _ = "end of CoverTab[66469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:322
			_go_fuzz_dep_.CoverTab[66470]++
																	switch tok.Kind() {
			case json.ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:324
				_go_fuzz_dep_.CoverTab[66473]++
																		return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:325
				// _ = "end of CoverTab[66473]"
			case json.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:326
				_go_fuzz_dep_.CoverTab[66474]++

																		if err := d.skipJSONValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:328
					_go_fuzz_dep_.CoverTab[66476]++
																			return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:329
					// _ = "end of CoverTab[66476]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
					_go_fuzz_dep_.CoverTab[66477]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
					// _ = "end of CoverTab[66477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
				// _ = "end of CoverTab[66474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
				_go_fuzz_dep_.CoverTab[66475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:330
				// _ = "end of CoverTab[66475]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:331
			// _ = "end of CoverTab[66470]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:332
		// _ = "end of CoverTab[66466]"

	case json.ArrayOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:334
		_go_fuzz_dep_.CoverTab[66467]++
																for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:335
			_go_fuzz_dep_.CoverTab[66478]++
																	tok, err := d.Peek()
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:337
				_go_fuzz_dep_.CoverTab[66480]++
																		return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:338
				// _ = "end of CoverTab[66480]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:339
				_go_fuzz_dep_.CoverTab[66481]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:339
				// _ = "end of CoverTab[66481]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:339
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:339
			// _ = "end of CoverTab[66478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:339
			_go_fuzz_dep_.CoverTab[66479]++
																	switch tok.Kind() {
			case json.ArrayClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:341
				_go_fuzz_dep_.CoverTab[66482]++
																		d.Read()
																		return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:343
				// _ = "end of CoverTab[66482]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:344
				_go_fuzz_dep_.CoverTab[66483]++

																		if err := d.skipJSONValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:346
					_go_fuzz_dep_.CoverTab[66484]++
																			return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:347
					// _ = "end of CoverTab[66484]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:348
					_go_fuzz_dep_.CoverTab[66485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:348
					// _ = "end of CoverTab[66485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:348
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:348
				// _ = "end of CoverTab[66483]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:349
			// _ = "end of CoverTab[66479]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:350
		// _ = "end of CoverTab[66467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:350
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:350
		_go_fuzz_dep_.CoverTab[66468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:350
		// _ = "end of CoverTab[66468]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:351
	// _ = "end of CoverTab[66462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:351
	_go_fuzz_dep_.CoverTab[66463]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:352
	// _ = "end of CoverTab[66463]"
}

// unmarshalAnyValue unmarshals the given custom-type message from the JSON
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:355
// object's "value" field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:357
func (d decoder) unmarshalAnyValue(unmarshal unmarshalFunc, m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:357
	_go_fuzz_dep_.CoverTab[66486]++

															d.Read()

															var found bool	// Used for detecting duplicate "value".
															for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:362
		_go_fuzz_dep_.CoverTab[66487]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:364
			_go_fuzz_dep_.CoverTab[66489]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:365
			// _ = "end of CoverTab[66489]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:366
			_go_fuzz_dep_.CoverTab[66490]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:366
			// _ = "end of CoverTab[66490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:366
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:366
		// _ = "end of CoverTab[66487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:366
		_go_fuzz_dep_.CoverTab[66488]++
																switch tok.Kind() {
		case json.ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:368
			_go_fuzz_dep_.CoverTab[66491]++
																	if !found {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:369
				_go_fuzz_dep_.CoverTab[66495]++
																		return d.newError(tok.Pos(), `missing "value" field`)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:370
				// _ = "end of CoverTab[66495]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:371
				_go_fuzz_dep_.CoverTab[66496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:371
				// _ = "end of CoverTab[66496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:371
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:371
			// _ = "end of CoverTab[66491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:371
			_go_fuzz_dep_.CoverTab[66492]++
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:372
			// _ = "end of CoverTab[66492]"

		case json.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:374
			_go_fuzz_dep_.CoverTab[66493]++
																	switch tok.Name() {
			case "@type":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:376
				_go_fuzz_dep_.CoverTab[66497]++

																		d.Read()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:378
				// _ = "end of CoverTab[66497]"

			case "value":
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:380
				_go_fuzz_dep_.CoverTab[66498]++
																		if found {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:381
					_go_fuzz_dep_.CoverTab[66503]++
																			return d.newError(tok.Pos(), `duplicate "value" field`)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:382
					// _ = "end of CoverTab[66503]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:383
					_go_fuzz_dep_.CoverTab[66504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:383
					// _ = "end of CoverTab[66504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:383
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:383
				// _ = "end of CoverTab[66498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:383
				_go_fuzz_dep_.CoverTab[66499]++

																		if err := unmarshal(d, m); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:385
					_go_fuzz_dep_.CoverTab[66505]++
																			return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:386
					// _ = "end of CoverTab[66505]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:387
					_go_fuzz_dep_.CoverTab[66506]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:387
					// _ = "end of CoverTab[66506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:387
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:387
				// _ = "end of CoverTab[66499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:387
				_go_fuzz_dep_.CoverTab[66500]++
																		found = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:388
				// _ = "end of CoverTab[66500]"

			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:390
				_go_fuzz_dep_.CoverTab[66501]++
																		if d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:391
					_go_fuzz_dep_.CoverTab[66507]++
																			if err := d.skipJSONValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:392
						_go_fuzz_dep_.CoverTab[66509]++
																				return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:393
						// _ = "end of CoverTab[66509]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:394
						_go_fuzz_dep_.CoverTab[66510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:394
						// _ = "end of CoverTab[66510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:394
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:394
					// _ = "end of CoverTab[66507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:394
					_go_fuzz_dep_.CoverTab[66508]++
																			continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:395
					// _ = "end of CoverTab[66508]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:396
					_go_fuzz_dep_.CoverTab[66511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:396
					// _ = "end of CoverTab[66511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:396
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:396
				// _ = "end of CoverTab[66501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:396
				_go_fuzz_dep_.CoverTab[66502]++
																		return d.newError(tok.Pos(), "unknown field %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:397
				// _ = "end of CoverTab[66502]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:398
			// _ = "end of CoverTab[66493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:398
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:398
			_go_fuzz_dep_.CoverTab[66494]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:398
			// _ = "end of CoverTab[66494]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:399
		// _ = "end of CoverTab[66488]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:400
	// _ = "end of CoverTab[66486]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:405
func (e encoder) marshalWrapperType(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:405
	_go_fuzz_dep_.CoverTab[66512]++
															fd := m.Descriptor().Fields().ByNumber(genid.WrapperValue_Value_field_number)
															val := m.Get(fd)
															return e.marshalSingular(val, fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:408
	// _ = "end of CoverTab[66512]"
}

func (d decoder) unmarshalWrapperType(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:411
	_go_fuzz_dep_.CoverTab[66513]++
															fd := m.Descriptor().Fields().ByNumber(genid.WrapperValue_Value_field_number)
															val, err := d.unmarshalScalar(fd)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:414
		_go_fuzz_dep_.CoverTab[66515]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:415
		// _ = "end of CoverTab[66515]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:416
		_go_fuzz_dep_.CoverTab[66516]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:416
		// _ = "end of CoverTab[66516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:416
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:416
	// _ = "end of CoverTab[66513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:416
	_go_fuzz_dep_.CoverTab[66514]++
															m.Set(fd, val)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:418
	// _ = "end of CoverTab[66514]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:423
func (e encoder) marshalEmpty(protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:423
	_go_fuzz_dep_.CoverTab[66517]++
															e.StartObject()
															e.EndObject()
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:426
	// _ = "end of CoverTab[66517]"
}

func (d decoder) unmarshalEmpty(protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:429
	_go_fuzz_dep_.CoverTab[66518]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:431
		_go_fuzz_dep_.CoverTab[66521]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:432
		// _ = "end of CoverTab[66521]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:433
		_go_fuzz_dep_.CoverTab[66522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:433
		// _ = "end of CoverTab[66522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:433
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:433
	// _ = "end of CoverTab[66518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:433
	_go_fuzz_dep_.CoverTab[66519]++
															if tok.Kind() != json.ObjectOpen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:434
		_go_fuzz_dep_.CoverTab[66523]++
																return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:435
		// _ = "end of CoverTab[66523]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:436
		_go_fuzz_dep_.CoverTab[66524]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:436
		// _ = "end of CoverTab[66524]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:436
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:436
	// _ = "end of CoverTab[66519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:436
	_go_fuzz_dep_.CoverTab[66520]++

															for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:438
		_go_fuzz_dep_.CoverTab[66525]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:440
			_go_fuzz_dep_.CoverTab[66527]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:441
			// _ = "end of CoverTab[66527]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:442
			_go_fuzz_dep_.CoverTab[66528]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:442
			// _ = "end of CoverTab[66528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:442
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:442
		// _ = "end of CoverTab[66525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:442
		_go_fuzz_dep_.CoverTab[66526]++
																switch tok.Kind() {
		case json.ObjectClose:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:444
			_go_fuzz_dep_.CoverTab[66529]++
																	return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:445
			// _ = "end of CoverTab[66529]"

		case json.Name:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:447
			_go_fuzz_dep_.CoverTab[66530]++
																	if d.opts.DiscardUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:448
				_go_fuzz_dep_.CoverTab[66533]++
																		if err := d.skipJSONValue(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:449
					_go_fuzz_dep_.CoverTab[66535]++
																			return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:450
					// _ = "end of CoverTab[66535]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:451
					_go_fuzz_dep_.CoverTab[66536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:451
					// _ = "end of CoverTab[66536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:451
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:451
				// _ = "end of CoverTab[66533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:451
				_go_fuzz_dep_.CoverTab[66534]++
																		continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:452
				// _ = "end of CoverTab[66534]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:453
				_go_fuzz_dep_.CoverTab[66537]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:453
				// _ = "end of CoverTab[66537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:453
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:453
			// _ = "end of CoverTab[66530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:453
			_go_fuzz_dep_.CoverTab[66531]++
																	return d.newError(tok.Pos(), "unknown field %v", tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:454
			// _ = "end of CoverTab[66531]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:456
			_go_fuzz_dep_.CoverTab[66532]++
																	return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:457
			// _ = "end of CoverTab[66532]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:458
		// _ = "end of CoverTab[66526]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:459
	// _ = "end of CoverTab[66520]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:465
func (e encoder) marshalStruct(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:465
	_go_fuzz_dep_.CoverTab[66538]++
															fd := m.Descriptor().Fields().ByNumber(genid.Struct_Fields_field_number)
															return e.marshalMap(m.Get(fd).Map(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:467
	// _ = "end of CoverTab[66538]"
}

func (d decoder) unmarshalStruct(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:470
	_go_fuzz_dep_.CoverTab[66539]++
															fd := m.Descriptor().Fields().ByNumber(genid.Struct_Fields_field_number)
															return d.unmarshalMap(m.Mutable(fd).Map(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:472
	// _ = "end of CoverTab[66539]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:479
func (e encoder) marshalListValue(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:479
	_go_fuzz_dep_.CoverTab[66540]++
															fd := m.Descriptor().Fields().ByNumber(genid.ListValue_Values_field_number)
															return e.marshalList(m.Get(fd).List(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:481
	// _ = "end of CoverTab[66540]"
}

func (d decoder) unmarshalListValue(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:484
	_go_fuzz_dep_.CoverTab[66541]++
															fd := m.Descriptor().Fields().ByNumber(genid.ListValue_Values_field_number)
															return d.unmarshalList(m.Mutable(fd).List(), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:486
	// _ = "end of CoverTab[66541]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:493
func (e encoder) marshalKnownValue(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:493
	_go_fuzz_dep_.CoverTab[66542]++
															od := m.Descriptor().Oneofs().ByName(genid.Value_Kind_oneof_name)
															fd := m.WhichOneof(od)
															if fd == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:496
		_go_fuzz_dep_.CoverTab[66545]++
																return errors.New("%s: none of the oneof fields is set", genid.Value_message_fullname)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:497
		// _ = "end of CoverTab[66545]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:498
		_go_fuzz_dep_.CoverTab[66546]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:498
		// _ = "end of CoverTab[66546]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:498
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:498
	// _ = "end of CoverTab[66542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:498
	_go_fuzz_dep_.CoverTab[66543]++
															if fd.Number() == genid.Value_NumberValue_field_number {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:499
		_go_fuzz_dep_.CoverTab[66547]++
																if v := m.Get(fd).Float(); math.IsNaN(v) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:500
			_go_fuzz_dep_.CoverTab[66548]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:500
			return math.IsInf(v, 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:500
			// _ = "end of CoverTab[66548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:500
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:500
			_go_fuzz_dep_.CoverTab[66549]++
																	return errors.New("%s: invalid %v value", genid.Value_NumberValue_field_fullname, v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:501
			// _ = "end of CoverTab[66549]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:502
			_go_fuzz_dep_.CoverTab[66550]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:502
			// _ = "end of CoverTab[66550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:502
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:502
		// _ = "end of CoverTab[66547]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:503
		_go_fuzz_dep_.CoverTab[66551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:503
		// _ = "end of CoverTab[66551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:503
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:503
	// _ = "end of CoverTab[66543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:503
	_go_fuzz_dep_.CoverTab[66544]++
															return e.marshalSingular(m.Get(fd), fd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:504
	// _ = "end of CoverTab[66544]"
}

func (d decoder) unmarshalKnownValue(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:507
	_go_fuzz_dep_.CoverTab[66552]++
															tok, err := d.Peek()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:509
		_go_fuzz_dep_.CoverTab[66555]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:510
		// _ = "end of CoverTab[66555]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:511
		_go_fuzz_dep_.CoverTab[66556]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:511
		// _ = "end of CoverTab[66556]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:511
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:511
	// _ = "end of CoverTab[66552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:511
	_go_fuzz_dep_.CoverTab[66553]++

															var fd protoreflect.FieldDescriptor
															var val protoreflect.Value
															switch tok.Kind() {
	case json.Null:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:516
		_go_fuzz_dep_.CoverTab[66557]++
																d.Read()
																fd = m.Descriptor().Fields().ByNumber(genid.Value_NullValue_field_number)
																val = protoreflect.ValueOfEnum(0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:519
		// _ = "end of CoverTab[66557]"

	case json.Bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:521
		_go_fuzz_dep_.CoverTab[66558]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:523
			_go_fuzz_dep_.CoverTab[66567]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:524
			// _ = "end of CoverTab[66567]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:525
			_go_fuzz_dep_.CoverTab[66568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:525
			// _ = "end of CoverTab[66568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:525
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:525
		// _ = "end of CoverTab[66558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:525
		_go_fuzz_dep_.CoverTab[66559]++
																fd = m.Descriptor().Fields().ByNumber(genid.Value_BoolValue_field_number)
																val = protoreflect.ValueOfBool(tok.Bool())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:527
		// _ = "end of CoverTab[66559]"

	case json.Number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:529
		_go_fuzz_dep_.CoverTab[66560]++
																tok, err := d.Read()
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:531
			_go_fuzz_dep_.CoverTab[66569]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:532
			// _ = "end of CoverTab[66569]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:533
			_go_fuzz_dep_.CoverTab[66570]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:533
			// _ = "end of CoverTab[66570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:533
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:533
		// _ = "end of CoverTab[66560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:533
		_go_fuzz_dep_.CoverTab[66561]++
																fd = m.Descriptor().Fields().ByNumber(genid.Value_NumberValue_field_number)
																var ok bool
																val, ok = unmarshalFloat(tok, 64)
																if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:537
			_go_fuzz_dep_.CoverTab[66571]++
																	return d.newError(tok.Pos(), "invalid %v: %v", genid.Value_message_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:538
			// _ = "end of CoverTab[66571]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:539
			_go_fuzz_dep_.CoverTab[66572]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:539
			// _ = "end of CoverTab[66572]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:539
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:539
		// _ = "end of CoverTab[66561]"

	case json.String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:541
		_go_fuzz_dep_.CoverTab[66562]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:548
		tok, err := d.Read()
		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:549
			_go_fuzz_dep_.CoverTab[66573]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:550
			// _ = "end of CoverTab[66573]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:551
			_go_fuzz_dep_.CoverTab[66574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:551
			// _ = "end of CoverTab[66574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:551
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:551
		// _ = "end of CoverTab[66562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:551
		_go_fuzz_dep_.CoverTab[66563]++
																fd = m.Descriptor().Fields().ByNumber(genid.Value_StringValue_field_number)
																val = protoreflect.ValueOfString(tok.ParsedString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:553
		// _ = "end of CoverTab[66563]"

	case json.ObjectOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:555
		_go_fuzz_dep_.CoverTab[66564]++
																fd = m.Descriptor().Fields().ByNumber(genid.Value_StructValue_field_number)
																val = m.NewField(fd)
																if err := d.unmarshalStruct(val.Message()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:558
			_go_fuzz_dep_.CoverTab[66575]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:559
			// _ = "end of CoverTab[66575]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:560
			_go_fuzz_dep_.CoverTab[66576]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:560
			// _ = "end of CoverTab[66576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:560
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:560
		// _ = "end of CoverTab[66564]"

	case json.ArrayOpen:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:562
		_go_fuzz_dep_.CoverTab[66565]++
																fd = m.Descriptor().Fields().ByNumber(genid.Value_ListValue_field_number)
																val = m.NewField(fd)
																if err := d.unmarshalListValue(val.Message()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:565
			_go_fuzz_dep_.CoverTab[66577]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:566
			// _ = "end of CoverTab[66577]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:567
			_go_fuzz_dep_.CoverTab[66578]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:567
			// _ = "end of CoverTab[66578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:567
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:567
		// _ = "end of CoverTab[66565]"

	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:569
		_go_fuzz_dep_.CoverTab[66566]++
																return d.newError(tok.Pos(), "invalid %v: %v", genid.Value_message_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:570
		// _ = "end of CoverTab[66566]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:571
	// _ = "end of CoverTab[66553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:571
	_go_fuzz_dep_.CoverTab[66554]++

															m.Set(fd, val)
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:574
	// _ = "end of CoverTab[66554]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:589
const (
	secondsInNanos		= 999999999
	maxSecondsInDuration	= 315576000000
)

func (e encoder) marshalDuration(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:594
	_go_fuzz_dep_.CoverTab[66579]++
															fds := m.Descriptor().Fields()
															fdSeconds := fds.ByNumber(genid.Duration_Seconds_field_number)
															fdNanos := fds.ByNumber(genid.Duration_Nanos_field_number)

															secsVal := m.Get(fdSeconds)
															nanosVal := m.Get(fdNanos)
															secs := secsVal.Int()
															nanos := nanosVal.Int()
															if secs < -maxSecondsInDuration || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:603
		_go_fuzz_dep_.CoverTab[66584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:603
		return secs > maxSecondsInDuration
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:603
		// _ = "end of CoverTab[66584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:603
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:603
		_go_fuzz_dep_.CoverTab[66585]++
																return errors.New("%s: seconds out of range %v", genid.Duration_message_fullname, secs)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:604
		// _ = "end of CoverTab[66585]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:605
		_go_fuzz_dep_.CoverTab[66586]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:605
		// _ = "end of CoverTab[66586]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:605
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:605
	// _ = "end of CoverTab[66579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:605
	_go_fuzz_dep_.CoverTab[66580]++
															if nanos < -secondsInNanos || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:606
		_go_fuzz_dep_.CoverTab[66587]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:606
		return nanos > secondsInNanos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:606
		// _ = "end of CoverTab[66587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:606
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:606
		_go_fuzz_dep_.CoverTab[66588]++
																return errors.New("%s: nanos out of range %v", genid.Duration_message_fullname, nanos)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:607
		// _ = "end of CoverTab[66588]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:608
		_go_fuzz_dep_.CoverTab[66589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:608
		// _ = "end of CoverTab[66589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:608
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:608
	// _ = "end of CoverTab[66580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:608
	_go_fuzz_dep_.CoverTab[66581]++
															if (secs > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		_go_fuzz_dep_.CoverTab[66590]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		return nanos < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		// _ = "end of CoverTab[66590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		_go_fuzz_dep_.CoverTab[66591]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		return (secs < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
			_go_fuzz_dep_.CoverTab[66592]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
			return nanos > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
			// _ = "end of CoverTab[66592]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		// _ = "end of CoverTab[66591]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:609
		_go_fuzz_dep_.CoverTab[66593]++
																return errors.New("%s: signs of seconds and nanos do not match", genid.Duration_message_fullname)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:610
		// _ = "end of CoverTab[66593]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:611
		_go_fuzz_dep_.CoverTab[66594]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:611
		// _ = "end of CoverTab[66594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:611
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:611
	// _ = "end of CoverTab[66581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:611
	_go_fuzz_dep_.CoverTab[66582]++
	// Generated output always contains 0, 3, 6, or 9 fractional digits,
	// depending on required precision, followed by the suffix "s".
	var sign string
	if secs < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:615
		_go_fuzz_dep_.CoverTab[66595]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:615
		return nanos < 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:615
		// _ = "end of CoverTab[66595]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:615
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:615
		_go_fuzz_dep_.CoverTab[66596]++
																sign, secs, nanos = "-", -1*secs, -1*nanos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:616
		// _ = "end of CoverTab[66596]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:617
		_go_fuzz_dep_.CoverTab[66597]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:617
		// _ = "end of CoverTab[66597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:617
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:617
	// _ = "end of CoverTab[66582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:617
	_go_fuzz_dep_.CoverTab[66583]++
															x := fmt.Sprintf("%s%d.%09d", sign, secs, nanos)
															x = strings.TrimSuffix(x, "000")
															x = strings.TrimSuffix(x, "000")
															x = strings.TrimSuffix(x, ".000")
															e.WriteString(x + "s")
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:623
	// _ = "end of CoverTab[66583]"
}

func (d decoder) unmarshalDuration(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:626
	_go_fuzz_dep_.CoverTab[66598]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:628
		_go_fuzz_dep_.CoverTab[66603]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:629
		// _ = "end of CoverTab[66603]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:630
		_go_fuzz_dep_.CoverTab[66604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:630
		// _ = "end of CoverTab[66604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:630
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:630
	// _ = "end of CoverTab[66598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:630
	_go_fuzz_dep_.CoverTab[66599]++
															if tok.Kind() != json.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:631
		_go_fuzz_dep_.CoverTab[66605]++
																return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:632
		// _ = "end of CoverTab[66605]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:633
		_go_fuzz_dep_.CoverTab[66606]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:633
		// _ = "end of CoverTab[66606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:633
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:633
	// _ = "end of CoverTab[66599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:633
	_go_fuzz_dep_.CoverTab[66600]++

															secs, nanos, ok := parseDuration(tok.ParsedString())
															if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:636
		_go_fuzz_dep_.CoverTab[66607]++
																return d.newError(tok.Pos(), "invalid %v value %v", genid.Duration_message_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:637
		// _ = "end of CoverTab[66607]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:638
		_go_fuzz_dep_.CoverTab[66608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:638
		// _ = "end of CoverTab[66608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:638
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:638
	// _ = "end of CoverTab[66600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:638
	_go_fuzz_dep_.CoverTab[66601]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:641
	if secs < -maxSecondsInDuration || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:641
		_go_fuzz_dep_.CoverTab[66609]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:641
		return secs > maxSecondsInDuration
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:641
		// _ = "end of CoverTab[66609]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:641
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:641
		_go_fuzz_dep_.CoverTab[66610]++
																return d.newError(tok.Pos(), "%v value out of range: %v", genid.Duration_message_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:642
		// _ = "end of CoverTab[66610]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:643
		_go_fuzz_dep_.CoverTab[66611]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:643
		// _ = "end of CoverTab[66611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:643
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:643
	// _ = "end of CoverTab[66601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:643
	_go_fuzz_dep_.CoverTab[66602]++

															fds := m.Descriptor().Fields()
															fdSeconds := fds.ByNumber(genid.Duration_Seconds_field_number)
															fdNanos := fds.ByNumber(genid.Duration_Nanos_field_number)

															m.Set(fdSeconds, protoreflect.ValueOfInt64(secs))
															m.Set(fdNanos, protoreflect.ValueOfInt32(nanos))
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:651
	// _ = "end of CoverTab[66602]"
}

// parseDuration parses the given input string for seconds and nanoseconds value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:654
// for the Duration JSON format. The format is a decimal number with a suffix
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:654
// 's'. It can have optional plus/minus sign. There needs to be at least an
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:654
// integer or fractional part. Fractional part is limited to 9 digits only for
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:654
// nanoseconds precision, regardless of whether there are trailing zero digits.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:654
// Example values are 1s, 0.1s, 1.s, .1s, +1s, -1s, -.1s.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:660
func parseDuration(input string) (int64, int32, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:660
	_go_fuzz_dep_.CoverTab[66612]++
															b := []byte(input)
															size := len(b)
															if size < 2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:663
		_go_fuzz_dep_.CoverTab[66622]++
																return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:664
		// _ = "end of CoverTab[66622]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:665
		_go_fuzz_dep_.CoverTab[66623]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:665
		// _ = "end of CoverTab[66623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:665
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:665
	// _ = "end of CoverTab[66612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:665
	_go_fuzz_dep_.CoverTab[66613]++
															if b[size-1] != 's' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:666
		_go_fuzz_dep_.CoverTab[66624]++
																return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:667
		// _ = "end of CoverTab[66624]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:668
		_go_fuzz_dep_.CoverTab[66625]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:668
		// _ = "end of CoverTab[66625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:668
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:668
	// _ = "end of CoverTab[66613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:668
	_go_fuzz_dep_.CoverTab[66614]++
															b = b[:size-1]

	// Read optional plus/minus symbol.
	var neg bool
	switch b[0] {
	case '-':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:674
		_go_fuzz_dep_.CoverTab[66626]++
																neg = true
																b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:676
		// _ = "end of CoverTab[66626]"
	case '+':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:677
		_go_fuzz_dep_.CoverTab[66627]++
																b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:678
		// _ = "end of CoverTab[66627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:678
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:678
		_go_fuzz_dep_.CoverTab[66628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:678
		// _ = "end of CoverTab[66628]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:679
	// _ = "end of CoverTab[66614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:679
	_go_fuzz_dep_.CoverTab[66615]++
															if len(b) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:680
		_go_fuzz_dep_.CoverTab[66629]++
																return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:681
		// _ = "end of CoverTab[66629]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:682
		_go_fuzz_dep_.CoverTab[66630]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:682
		// _ = "end of CoverTab[66630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:682
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:682
	// _ = "end of CoverTab[66615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:682
	_go_fuzz_dep_.CoverTab[66616]++

	// Read the integer part.
	var intp []byte
	switch {
	case b[0] == '0':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:687
		_go_fuzz_dep_.CoverTab[66631]++
																b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:688
		// _ = "end of CoverTab[66631]"

	case '1' <= b[0] && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:690
		_go_fuzz_dep_.CoverTab[66636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:690
		return b[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:690
		// _ = "end of CoverTab[66636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:690
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:690
		_go_fuzz_dep_.CoverTab[66632]++
																intp = b[0:]
																b = b[1:]
																n := 1
																for len(b) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			_go_fuzz_dep_.CoverTab[66637]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			return '0' <= b[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			// _ = "end of CoverTab[66637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			_go_fuzz_dep_.CoverTab[66638]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			return b[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			// _ = "end of CoverTab[66638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:694
			_go_fuzz_dep_.CoverTab[66639]++
																	n++
																	b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:696
			// _ = "end of CoverTab[66639]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:697
		// _ = "end of CoverTab[66632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:697
		_go_fuzz_dep_.CoverTab[66633]++
																intp = intp[:n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:698
		// _ = "end of CoverTab[66633]"

	case b[0] == '.':
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:700
		_go_fuzz_dep_.CoverTab[66634]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:700
		// _ = "end of CoverTab[66634]"

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:703
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:703
		_go_fuzz_dep_.CoverTab[66635]++
																return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:704
		// _ = "end of CoverTab[66635]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:705
	// _ = "end of CoverTab[66616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:705
	_go_fuzz_dep_.CoverTab[66617]++

															hasFrac := false
															var frac [9]byte
															if len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:709
		_go_fuzz_dep_.CoverTab[66640]++
																if b[0] != '.' {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:710
			_go_fuzz_dep_.CoverTab[66645]++
																	return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:711
			// _ = "end of CoverTab[66645]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:712
			_go_fuzz_dep_.CoverTab[66646]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:712
			// _ = "end of CoverTab[66646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:712
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:712
		// _ = "end of CoverTab[66640]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:712
		_go_fuzz_dep_.CoverTab[66641]++

																b = b[1:]
																n := 0
																for len(b) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			_go_fuzz_dep_.CoverTab[66647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			return n < 9
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			// _ = "end of CoverTab[66647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			_go_fuzz_dep_.CoverTab[66648]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			return '0' <= b[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			// _ = "end of CoverTab[66648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			_go_fuzz_dep_.CoverTab[66649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			return b[0] <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			// _ = "end of CoverTab[66649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:716
			_go_fuzz_dep_.CoverTab[66650]++
																	frac[n] = b[0]
																	n++
																	b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:719
			// _ = "end of CoverTab[66650]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:720
		// _ = "end of CoverTab[66641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:720
		_go_fuzz_dep_.CoverTab[66642]++

																if len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:722
			_go_fuzz_dep_.CoverTab[66651]++
																	return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:723
			// _ = "end of CoverTab[66651]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:724
			_go_fuzz_dep_.CoverTab[66652]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:724
			// _ = "end of CoverTab[66652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:724
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:724
		// _ = "end of CoverTab[66642]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:724
		_go_fuzz_dep_.CoverTab[66643]++

																for i := n; i < 9; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:726
			_go_fuzz_dep_.CoverTab[66653]++
																	frac[i] = '0'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:727
			// _ = "end of CoverTab[66653]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:728
		// _ = "end of CoverTab[66643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:728
		_go_fuzz_dep_.CoverTab[66644]++
																hasFrac = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:729
		// _ = "end of CoverTab[66644]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:730
		_go_fuzz_dep_.CoverTab[66654]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:730
		// _ = "end of CoverTab[66654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:730
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:730
	// _ = "end of CoverTab[66617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:730
	_go_fuzz_dep_.CoverTab[66618]++

															var secs int64
															if len(intp) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:733
		_go_fuzz_dep_.CoverTab[66655]++
																var err error
																secs, err = strconv.ParseInt(string(intp), 10, 64)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:736
			_go_fuzz_dep_.CoverTab[66656]++
																	return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:737
			// _ = "end of CoverTab[66656]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:738
			_go_fuzz_dep_.CoverTab[66657]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:738
			// _ = "end of CoverTab[66657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:738
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:738
		// _ = "end of CoverTab[66655]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:739
		_go_fuzz_dep_.CoverTab[66658]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:739
		// _ = "end of CoverTab[66658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:739
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:739
	// _ = "end of CoverTab[66618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:739
	_go_fuzz_dep_.CoverTab[66619]++

															var nanos int64
															if hasFrac {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:742
		_go_fuzz_dep_.CoverTab[66659]++
																nanob := bytes.TrimLeft(frac[:], "0")
																if len(nanob) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:744
			_go_fuzz_dep_.CoverTab[66660]++
																	var err error
																	nanos, err = strconv.ParseInt(string(nanob), 10, 32)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:747
				_go_fuzz_dep_.CoverTab[66661]++
																		return 0, 0, false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:748
				// _ = "end of CoverTab[66661]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:749
				_go_fuzz_dep_.CoverTab[66662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:749
				// _ = "end of CoverTab[66662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:749
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:749
			// _ = "end of CoverTab[66660]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:750
			_go_fuzz_dep_.CoverTab[66663]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:750
			// _ = "end of CoverTab[66663]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:750
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:750
		// _ = "end of CoverTab[66659]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:751
		_go_fuzz_dep_.CoverTab[66664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:751
		// _ = "end of CoverTab[66664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:751
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:751
	// _ = "end of CoverTab[66619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:751
	_go_fuzz_dep_.CoverTab[66620]++

															if neg {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:753
		_go_fuzz_dep_.CoverTab[66665]++
																if secs > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:754
			_go_fuzz_dep_.CoverTab[66667]++
																	secs = -secs
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:755
			// _ = "end of CoverTab[66667]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:756
			_go_fuzz_dep_.CoverTab[66668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:756
			// _ = "end of CoverTab[66668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:756
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:756
		// _ = "end of CoverTab[66665]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:756
		_go_fuzz_dep_.CoverTab[66666]++
																if nanos > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:757
			_go_fuzz_dep_.CoverTab[66669]++
																	nanos = -nanos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:758
			// _ = "end of CoverTab[66669]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:759
			_go_fuzz_dep_.CoverTab[66670]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:759
			// _ = "end of CoverTab[66670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:759
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:759
		// _ = "end of CoverTab[66666]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:760
		_go_fuzz_dep_.CoverTab[66671]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:760
		// _ = "end of CoverTab[66671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:760
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:760
	// _ = "end of CoverTab[66620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:760
	_go_fuzz_dep_.CoverTab[66621]++
															return secs, int32(nanos), true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:761
	// _ = "end of CoverTab[66621]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:777
const (
	maxTimestampSeconds	= 253402300799
	minTimestampSeconds	= -62135596800
)

func (e encoder) marshalTimestamp(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:782
	_go_fuzz_dep_.CoverTab[66672]++
															fds := m.Descriptor().Fields()
															fdSeconds := fds.ByNumber(genid.Timestamp_Seconds_field_number)
															fdNanos := fds.ByNumber(genid.Timestamp_Nanos_field_number)

															secsVal := m.Get(fdSeconds)
															nanosVal := m.Get(fdNanos)
															secs := secsVal.Int()
															nanos := nanosVal.Int()
															if secs < minTimestampSeconds || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:791
		_go_fuzz_dep_.CoverTab[66675]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:791
		return secs > maxTimestampSeconds
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:791
		// _ = "end of CoverTab[66675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:791
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:791
		_go_fuzz_dep_.CoverTab[66676]++
																return errors.New("%s: seconds out of range %v", genid.Timestamp_message_fullname, secs)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:792
		// _ = "end of CoverTab[66676]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:793
		_go_fuzz_dep_.CoverTab[66677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:793
		// _ = "end of CoverTab[66677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:793
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:793
	// _ = "end of CoverTab[66672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:793
	_go_fuzz_dep_.CoverTab[66673]++
															if nanos < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:794
		_go_fuzz_dep_.CoverTab[66678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:794
		return nanos > secondsInNanos
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:794
		// _ = "end of CoverTab[66678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:794
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:794
		_go_fuzz_dep_.CoverTab[66679]++
																return errors.New("%s: nanos out of range %v", genid.Timestamp_message_fullname, nanos)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:795
		// _ = "end of CoverTab[66679]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:796
		_go_fuzz_dep_.CoverTab[66680]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:796
		// _ = "end of CoverTab[66680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:796
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:796
	// _ = "end of CoverTab[66673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:796
	_go_fuzz_dep_.CoverTab[66674]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:799
	t := time.Unix(secs, nanos).UTC()
															x := t.Format("2006-01-02T15:04:05.000000000")
															x = strings.TrimSuffix(x, "000")
															x = strings.TrimSuffix(x, "000")
															x = strings.TrimSuffix(x, ".000")
															e.WriteString(x + "Z")
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:805
	// _ = "end of CoverTab[66674]"
}

func (d decoder) unmarshalTimestamp(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:808
	_go_fuzz_dep_.CoverTab[66681]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:810
		_go_fuzz_dep_.CoverTab[66686]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:811
		// _ = "end of CoverTab[66686]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:812
		_go_fuzz_dep_.CoverTab[66687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:812
		// _ = "end of CoverTab[66687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:812
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:812
	// _ = "end of CoverTab[66681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:812
	_go_fuzz_dep_.CoverTab[66682]++
															if tok.Kind() != json.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:813
		_go_fuzz_dep_.CoverTab[66688]++
																return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:814
		// _ = "end of CoverTab[66688]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:815
		_go_fuzz_dep_.CoverTab[66689]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:815
		// _ = "end of CoverTab[66689]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:815
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:815
	// _ = "end of CoverTab[66682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:815
	_go_fuzz_dep_.CoverTab[66683]++

															t, err := time.Parse(time.RFC3339Nano, tok.ParsedString())
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:818
		_go_fuzz_dep_.CoverTab[66690]++
																return d.newError(tok.Pos(), "invalid %v value %v", genid.Timestamp_message_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:819
		// _ = "end of CoverTab[66690]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:820
		_go_fuzz_dep_.CoverTab[66691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:820
		// _ = "end of CoverTab[66691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:820
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:820
	// _ = "end of CoverTab[66683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:820
	_go_fuzz_dep_.CoverTab[66684]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:823
	secs := t.Unix()
	if secs < minTimestampSeconds || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:824
		_go_fuzz_dep_.CoverTab[66692]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:824
		return secs > maxTimestampSeconds
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:824
		// _ = "end of CoverTab[66692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:824
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:824
		_go_fuzz_dep_.CoverTab[66693]++
																return d.newError(tok.Pos(), "%v value out of range: %v", genid.Timestamp_message_fullname, tok.RawString())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:825
		// _ = "end of CoverTab[66693]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:826
		_go_fuzz_dep_.CoverTab[66694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:826
		// _ = "end of CoverTab[66694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:826
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:826
	// _ = "end of CoverTab[66684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:826
	_go_fuzz_dep_.CoverTab[66685]++

															fds := m.Descriptor().Fields()
															fdSeconds := fds.ByNumber(genid.Timestamp_Seconds_field_number)
															fdNanos := fds.ByNumber(genid.Timestamp_Nanos_field_number)

															m.Set(fdSeconds, protoreflect.ValueOfInt64(secs))
															m.Set(fdNanos, protoreflect.ValueOfInt32(int32(t.Nanosecond())))
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:834
	// _ = "end of CoverTab[66685]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:842
func (e encoder) marshalFieldMask(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:842
	_go_fuzz_dep_.CoverTab[66695]++
															fd := m.Descriptor().Fields().ByNumber(genid.FieldMask_Paths_field_number)
															list := m.Get(fd).List()
															paths := make([]string, 0, list.Len())

															for i := 0; i < list.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:847
		_go_fuzz_dep_.CoverTab[66697]++
																s := list.Get(i).String()
																if !protoreflect.FullName(s).IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:849
			_go_fuzz_dep_.CoverTab[66700]++
																	return errors.New("%s contains invalid path: %q", genid.FieldMask_Paths_field_fullname, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:850
			// _ = "end of CoverTab[66700]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:851
			_go_fuzz_dep_.CoverTab[66701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:851
			// _ = "end of CoverTab[66701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:851
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:851
		// _ = "end of CoverTab[66697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:851
		_go_fuzz_dep_.CoverTab[66698]++

																cc := strs.JSONCamelCase(s)
																if s != strs.JSONSnakeCase(cc) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:854
			_go_fuzz_dep_.CoverTab[66702]++
																	return errors.New("%s contains irreversible value %q", genid.FieldMask_Paths_field_fullname, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:855
			// _ = "end of CoverTab[66702]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:856
			_go_fuzz_dep_.CoverTab[66703]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:856
			// _ = "end of CoverTab[66703]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:856
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:856
		// _ = "end of CoverTab[66698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:856
		_go_fuzz_dep_.CoverTab[66699]++
																paths = append(paths, cc)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:857
		// _ = "end of CoverTab[66699]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:858
	// _ = "end of CoverTab[66695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:858
	_go_fuzz_dep_.CoverTab[66696]++

															e.WriteString(strings.Join(paths, ","))
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:861
	// _ = "end of CoverTab[66696]"
}

func (d decoder) unmarshalFieldMask(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:864
	_go_fuzz_dep_.CoverTab[66704]++
															tok, err := d.Read()
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:866
		_go_fuzz_dep_.CoverTab[66709]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:867
		// _ = "end of CoverTab[66709]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:868
		_go_fuzz_dep_.CoverTab[66710]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:868
		// _ = "end of CoverTab[66710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:868
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:868
	// _ = "end of CoverTab[66704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:868
	_go_fuzz_dep_.CoverTab[66705]++
															if tok.Kind() != json.String {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:869
		_go_fuzz_dep_.CoverTab[66711]++
																return d.unexpectedTokenError(tok)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:870
		// _ = "end of CoverTab[66711]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:871
		_go_fuzz_dep_.CoverTab[66712]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:871
		// _ = "end of CoverTab[66712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:871
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:871
	// _ = "end of CoverTab[66705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:871
	_go_fuzz_dep_.CoverTab[66706]++
															str := strings.TrimSpace(tok.ParsedString())
															if str == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:873
		_go_fuzz_dep_.CoverTab[66713]++
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:874
		// _ = "end of CoverTab[66713]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:875
		_go_fuzz_dep_.CoverTab[66714]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:875
		// _ = "end of CoverTab[66714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:875
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:875
	// _ = "end of CoverTab[66706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:875
	_go_fuzz_dep_.CoverTab[66707]++
															paths := strings.Split(str, ",")

															fd := m.Descriptor().Fields().ByNumber(genid.FieldMask_Paths_field_number)
															list := m.Mutable(fd).List()

															for _, s0 := range paths {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:881
		_go_fuzz_dep_.CoverTab[66715]++
																s := strs.JSONSnakeCase(s0)
																if strings.Contains(s0, "_") || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:883
			_go_fuzz_dep_.CoverTab[66717]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:883
			return !protoreflect.FullName(s).IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:883
			// _ = "end of CoverTab[66717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:883
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:883
			_go_fuzz_dep_.CoverTab[66718]++
																	return d.newError(tok.Pos(), "%v contains invalid path: %q", genid.FieldMask_Paths_field_fullname, s0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:884
			// _ = "end of CoverTab[66718]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:885
			_go_fuzz_dep_.CoverTab[66719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:885
			// _ = "end of CoverTab[66719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:885
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:885
		// _ = "end of CoverTab[66715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:885
		_go_fuzz_dep_.CoverTab[66716]++
																list.Append(protoreflect.ValueOfString(s))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:886
		// _ = "end of CoverTab[66716]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:887
	// _ = "end of CoverTab[66707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:887
	_go_fuzz_dep_.CoverTab[66708]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:888
	// _ = "end of CoverTab[66708]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:889
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/encoding/protojson/well_known_types.go:889
var _ = _go_fuzz_dep_.CoverTab
