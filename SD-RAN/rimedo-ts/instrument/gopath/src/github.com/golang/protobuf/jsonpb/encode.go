// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
package jsonpb

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:5
)

import (
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

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	protoV2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const wrapJSONMarshalV2 = false

// Marshaler is a configurable object for marshaling protocol buffer messages
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:28
// to the specified JSON representation.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:30
type Marshaler struct {
	// OrigName specifies whether to use the original protobuf name for fields.
	OrigName	bool

	// EnumsAsInts specifies whether to render enum values as integers,
	// as opposed to string values.
	EnumsAsInts	bool

	// EmitDefaults specifies whether to render fields with zero values.
	EmitDefaults	bool

	// Indent controls whether the output is compact or not.
	// If empty, the output is compact JSON. Otherwise, every JSON object
	// entry and JSON array value will be on its own line.
	// Each line will be preceded by repeated copies of Indent, where the
	// number of copies is the current indentation depth.
	Indent	string

	// AnyResolver is used to resolve the google.protobuf.Any well-known type.
	// If unset, the global registry is used by default.
	AnyResolver	AnyResolver
}

// JSONPBMarshaler is implemented by protobuf messages that customize the
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
// way they are marshaled to JSON. Messages that implement this should also
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
// implement JSONPBUnmarshaler so that the custom format can be parsed.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
// The JSON marshaling must follow the proto to JSON specification:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
//	https://developers.google.com/protocol-buffers/docs/proto3#json
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:53
// Deprecated: Custom types should implement protobuf reflection instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:61
type JSONPBMarshaler interface {
	MarshalJSONPB(*Marshaler) ([]byte, error)
}

// Marshal serializes a protobuf message as JSON into w.
func (jm *Marshaler) Marshal(w io.Writer, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:66
	_go_fuzz_dep_.CoverTab[66988]++
												b, err := jm.marshal(m)
												if len(b) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:68
		_go_fuzz_dep_.CoverTab[66990]++
													if _, err := w.Write(b); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:69
			_go_fuzz_dep_.CoverTab[66991]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:70
			// _ = "end of CoverTab[66991]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:71
			_go_fuzz_dep_.CoverTab[66992]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:71
			// _ = "end of CoverTab[66992]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:71
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:71
		// _ = "end of CoverTab[66990]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:72
		_go_fuzz_dep_.CoverTab[66993]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:72
		// _ = "end of CoverTab[66993]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:72
	// _ = "end of CoverTab[66988]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:72
	_go_fuzz_dep_.CoverTab[66989]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:73
	// _ = "end of CoverTab[66989]"
}

// MarshalToString serializes a protobuf message as JSON in string form.
func (jm *Marshaler) MarshalToString(m proto.Message) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:77
	_go_fuzz_dep_.CoverTab[66994]++
												b, err := jm.marshal(m)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:79
		_go_fuzz_dep_.CoverTab[66996]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:80
		// _ = "end of CoverTab[66996]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:81
		_go_fuzz_dep_.CoverTab[66997]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:81
		// _ = "end of CoverTab[66997]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:81
	// _ = "end of CoverTab[66994]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:81
	_go_fuzz_dep_.CoverTab[66995]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:82
	// _ = "end of CoverTab[66995]"
}

func (jm *Marshaler) marshal(m proto.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:85
	_go_fuzz_dep_.CoverTab[66998]++
												v := reflect.ValueOf(m)
												if m == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
		_go_fuzz_dep_.CoverTab[67001]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
		return (v.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
			_go_fuzz_dep_.CoverTab[67002]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
			return v.IsNil()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
			// _ = "end of CoverTab[67002]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
		}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
		// _ = "end of CoverTab[67001]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:87
		_go_fuzz_dep_.CoverTab[67003]++
													return nil, errors.New("Marshal called with nil")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:88
		// _ = "end of CoverTab[67003]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:89
		_go_fuzz_dep_.CoverTab[67004]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:89
		// _ = "end of CoverTab[67004]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:89
	// _ = "end of CoverTab[66998]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:89
	_go_fuzz_dep_.CoverTab[66999]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:93
	if jsm, ok := m.(JSONPBMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:93
		_go_fuzz_dep_.CoverTab[67005]++
													return jsm.MarshalJSONPB(jm)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:94
		// _ = "end of CoverTab[67005]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:95
		_go_fuzz_dep_.CoverTab[67006]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:95
		// _ = "end of CoverTab[67006]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:95
	// _ = "end of CoverTab[66999]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:95
	_go_fuzz_dep_.CoverTab[67000]++

												if wrapJSONMarshalV2 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:97
		_go_fuzz_dep_.CoverTab[67007]++
													opts := protojson.MarshalOptions{
			UseProtoNames:		jm.OrigName,
			UseEnumNumbers:		jm.EnumsAsInts,
			EmitUnpopulated:	jm.EmitDefaults,
			Indent:			jm.Indent,
		}
		if jm.AnyResolver != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:104
			_go_fuzz_dep_.CoverTab[67009]++
														opts.Resolver = anyResolver{jm.AnyResolver}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:105
			// _ = "end of CoverTab[67009]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:106
			_go_fuzz_dep_.CoverTab[67010]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:106
			// _ = "end of CoverTab[67010]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:106
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:106
		// _ = "end of CoverTab[67007]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:106
		_go_fuzz_dep_.CoverTab[67008]++
													return opts.Marshal(proto.MessageReflect(m).Interface())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:107
		// _ = "end of CoverTab[67008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:108
		_go_fuzz_dep_.CoverTab[67011]++

													m2 := proto.MessageReflect(m)
													if err := protoV2.CheckInitialized(m2.Interface()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:111
			_go_fuzz_dep_.CoverTab[67013]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:112
			// _ = "end of CoverTab[67013]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:113
			_go_fuzz_dep_.CoverTab[67014]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:113
			// _ = "end of CoverTab[67014]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:113
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:113
		// _ = "end of CoverTab[67011]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:113
		_go_fuzz_dep_.CoverTab[67012]++

													w := jsonWriter{Marshaler: jm}
													err := w.marshalMessage(m2, "", "")
													return w.buf, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:117
		// _ = "end of CoverTab[67012]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:118
	// _ = "end of CoverTab[67000]"
}

type jsonWriter struct {
	*Marshaler
	buf	[]byte
}

func (w *jsonWriter) write(s string) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:126
	_go_fuzz_dep_.CoverTab[67015]++
												w.buf = append(w.buf, s...)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:127
	// _ = "end of CoverTab[67015]"
}

func (w *jsonWriter) marshalMessage(m protoreflect.Message, indent, typeURL string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:130
	_go_fuzz_dep_.CoverTab[67016]++
												if jsm, ok := proto.MessageV1(m.Interface()).(JSONPBMarshaler); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:131
		_go_fuzz_dep_.CoverTab[67024]++
													b, err := jsm.MarshalJSONPB(w.Marshaler)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:133
			_go_fuzz_dep_.CoverTab[67027]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:134
			// _ = "end of CoverTab[67027]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:135
			_go_fuzz_dep_.CoverTab[67028]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:135
			// _ = "end of CoverTab[67028]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:135
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:135
		// _ = "end of CoverTab[67024]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:135
		_go_fuzz_dep_.CoverTab[67025]++
													if typeURL != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:136
			_go_fuzz_dep_.CoverTab[67029]++
			// we are marshaling this object to an Any type
			var js map[string]*json.RawMessage
			if err = json.Unmarshal(b, &js); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:139
				_go_fuzz_dep_.CoverTab[67032]++
															return fmt.Errorf("type %T produced invalid JSON: %v", m.Interface(), err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:140
				// _ = "end of CoverTab[67032]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:141
				_go_fuzz_dep_.CoverTab[67033]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:141
				// _ = "end of CoverTab[67033]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:141
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:141
			// _ = "end of CoverTab[67029]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:141
			_go_fuzz_dep_.CoverTab[67030]++
														turl, err := json.Marshal(typeURL)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:143
				_go_fuzz_dep_.CoverTab[67034]++
															return fmt.Errorf("failed to marshal type URL %q to JSON: %v", typeURL, err)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:144
				// _ = "end of CoverTab[67034]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:145
				_go_fuzz_dep_.CoverTab[67035]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:145
				// _ = "end of CoverTab[67035]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:145
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:145
			// _ = "end of CoverTab[67030]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:145
			_go_fuzz_dep_.CoverTab[67031]++
														js["@type"] = (*json.RawMessage)(&turl)
														if b, err = json.Marshal(js); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:147
				_go_fuzz_dep_.CoverTab[67036]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:148
				// _ = "end of CoverTab[67036]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:149
				_go_fuzz_dep_.CoverTab[67037]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:149
				// _ = "end of CoverTab[67037]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:149
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:149
			// _ = "end of CoverTab[67031]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:150
			_go_fuzz_dep_.CoverTab[67038]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:150
			// _ = "end of CoverTab[67038]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:150
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:150
		// _ = "end of CoverTab[67025]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:150
		_go_fuzz_dep_.CoverTab[67026]++
													w.write(string(b))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:152
		// _ = "end of CoverTab[67026]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:153
		_go_fuzz_dep_.CoverTab[67039]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:153
		// _ = "end of CoverTab[67039]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:153
	// _ = "end of CoverTab[67016]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:153
	_go_fuzz_dep_.CoverTab[67017]++

												md := m.Descriptor()
												fds := md.Fields()

	// Handle well-known types.
	const secondInNanos = int64(time.Second / time.Nanosecond)
	switch wellKnownType(md.FullName()) {
	case "Any":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:161
		_go_fuzz_dep_.CoverTab[67040]++
													return w.marshalAny(m, indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:162
		// _ = "end of CoverTab[67040]"
	case "BoolValue", "BytesValue", "StringValue",
		"Int32Value", "UInt32Value", "FloatValue",
		"Int64Value", "UInt64Value", "DoubleValue":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:165
		_go_fuzz_dep_.CoverTab[67041]++
													fd := fds.ByNumber(1)
													return w.marshalValue(fd, m.Get(fd), indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:167
		// _ = "end of CoverTab[67041]"
	case "Duration":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:168
		_go_fuzz_dep_.CoverTab[67042]++
													const maxSecondsInDuration = 315576000000

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:172
		s := m.Get(fds.ByNumber(1)).Int()
		ns := m.Get(fds.ByNumber(2)).Int()
		if s < -maxSecondsInDuration || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:174
			_go_fuzz_dep_.CoverTab[67053]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:174
			return s > maxSecondsInDuration
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:174
			// _ = "end of CoverTab[67053]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:174
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:174
			_go_fuzz_dep_.CoverTab[67054]++
														return fmt.Errorf("seconds out of range %v", s)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:175
			// _ = "end of CoverTab[67054]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:176
			_go_fuzz_dep_.CoverTab[67055]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:176
			// _ = "end of CoverTab[67055]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:176
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:176
		// _ = "end of CoverTab[67042]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:176
		_go_fuzz_dep_.CoverTab[67043]++
													if ns <= -secondInNanos || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:177
			_go_fuzz_dep_.CoverTab[67056]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:177
			return ns >= secondInNanos
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:177
			// _ = "end of CoverTab[67056]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:177
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:177
			_go_fuzz_dep_.CoverTab[67057]++
														return fmt.Errorf("ns out of range (%v, %v)", -secondInNanos, secondInNanos)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:178
			// _ = "end of CoverTab[67057]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:179
			_go_fuzz_dep_.CoverTab[67058]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:179
			// _ = "end of CoverTab[67058]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:179
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:179
		// _ = "end of CoverTab[67043]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:179
		_go_fuzz_dep_.CoverTab[67044]++
													if (s > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			_go_fuzz_dep_.CoverTab[67059]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			return ns < 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			// _ = "end of CoverTab[67059]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			_go_fuzz_dep_.CoverTab[67060]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			return (s < 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
				_go_fuzz_dep_.CoverTab[67061]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
				return ns > 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
				// _ = "end of CoverTab[67061]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			// _ = "end of CoverTab[67060]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:180
			_go_fuzz_dep_.CoverTab[67062]++
														return errors.New("signs of seconds and nanos do not match")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:181
			// _ = "end of CoverTab[67062]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:182
			_go_fuzz_dep_.CoverTab[67063]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:182
			// _ = "end of CoverTab[67063]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:182
		// _ = "end of CoverTab[67044]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:182
		_go_fuzz_dep_.CoverTab[67045]++
													var sign string
													if s < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:184
			_go_fuzz_dep_.CoverTab[67064]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:184
			return ns < 0
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:184
			// _ = "end of CoverTab[67064]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:184
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:184
			_go_fuzz_dep_.CoverTab[67065]++
														sign, s, ns = "-", -1*s, -1*ns
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:185
			// _ = "end of CoverTab[67065]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:186
			_go_fuzz_dep_.CoverTab[67066]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:186
			// _ = "end of CoverTab[67066]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:186
		// _ = "end of CoverTab[67045]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:186
		_go_fuzz_dep_.CoverTab[67046]++
													x := fmt.Sprintf("%s%d.%09d", sign, s, ns)
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, ".000")
													w.write(fmt.Sprintf(`"%vs"`, x))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:192
		// _ = "end of CoverTab[67046]"
	case "Timestamp":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:193
		_go_fuzz_dep_.CoverTab[67047]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:196
		s := m.Get(fds.ByNumber(1)).Int()
		ns := m.Get(fds.ByNumber(2)).Int()
		if ns < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:198
			_go_fuzz_dep_.CoverTab[67067]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:198
			return ns >= secondInNanos
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:198
			// _ = "end of CoverTab[67067]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:198
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:198
			_go_fuzz_dep_.CoverTab[67068]++
														return fmt.Errorf("ns out of range [0, %v)", secondInNanos)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:199
			// _ = "end of CoverTab[67068]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:200
			_go_fuzz_dep_.CoverTab[67069]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:200
			// _ = "end of CoverTab[67069]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:200
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:200
		// _ = "end of CoverTab[67047]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:200
		_go_fuzz_dep_.CoverTab[67048]++
													t := time.Unix(s, ns).UTC()

													x := t.Format("2006-01-02T15:04:05.000000000")
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, "000")
													x = strings.TrimSuffix(x, ".000")
													w.write(fmt.Sprintf(`"%vZ"`, x))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:208
		// _ = "end of CoverTab[67048]"
	case "Value":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:209
		_go_fuzz_dep_.CoverTab[67049]++

													od := md.Oneofs().Get(0)
													fd := m.WhichOneof(od)
													if fd == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:213
			_go_fuzz_dep_.CoverTab[67070]++
														return errors.New("nil Value")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:214
			// _ = "end of CoverTab[67070]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:215
			_go_fuzz_dep_.CoverTab[67071]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:215
			// _ = "end of CoverTab[67071]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:215
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:215
		// _ = "end of CoverTab[67049]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:215
		_go_fuzz_dep_.CoverTab[67050]++
													return w.marshalValue(fd, m.Get(fd), indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:216
		// _ = "end of CoverTab[67050]"
	case "Struct", "ListValue":
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:217
		_go_fuzz_dep_.CoverTab[67051]++

													fd := fds.ByNumber(1)
													return w.marshalValue(fd, m.Get(fd), indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:220
		// _ = "end of CoverTab[67051]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:220
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:220
		_go_fuzz_dep_.CoverTab[67052]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:220
		// _ = "end of CoverTab[67052]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:221
	// _ = "end of CoverTab[67017]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:221
	_go_fuzz_dep_.CoverTab[67018]++

												w.write("{")
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:224
		_go_fuzz_dep_.CoverTab[67072]++
													w.write("\n")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:225
		// _ = "end of CoverTab[67072]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:226
		_go_fuzz_dep_.CoverTab[67073]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:226
		// _ = "end of CoverTab[67073]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:226
	// _ = "end of CoverTab[67018]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:226
	_go_fuzz_dep_.CoverTab[67019]++

												firstField := true
												if typeURL != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:229
		_go_fuzz_dep_.CoverTab[67074]++
													if err := w.marshalTypeURL(indent, typeURL); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:230
			_go_fuzz_dep_.CoverTab[67076]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:231
			// _ = "end of CoverTab[67076]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:232
			_go_fuzz_dep_.CoverTab[67077]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:232
			// _ = "end of CoverTab[67077]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:232
		// _ = "end of CoverTab[67074]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:232
		_go_fuzz_dep_.CoverTab[67075]++
													firstField = false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:233
		// _ = "end of CoverTab[67075]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:234
		_go_fuzz_dep_.CoverTab[67078]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:234
		// _ = "end of CoverTab[67078]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:234
	// _ = "end of CoverTab[67019]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:234
	_go_fuzz_dep_.CoverTab[67020]++

												for i := 0; i < fds.Len(); {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:236
		_go_fuzz_dep_.CoverTab[67079]++
													fd := fds.Get(i)
													if od := fd.ContainingOneof(); od != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:238
			_go_fuzz_dep_.CoverTab[67084]++
														fd = m.WhichOneof(od)
														i += od.Fields().Len()
														if fd == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:241
				_go_fuzz_dep_.CoverTab[67085]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:242
				// _ = "end of CoverTab[67085]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:243
				_go_fuzz_dep_.CoverTab[67086]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:243
				// _ = "end of CoverTab[67086]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:243
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:243
			// _ = "end of CoverTab[67084]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:244
			_go_fuzz_dep_.CoverTab[67087]++
														i++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:245
			// _ = "end of CoverTab[67087]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:246
		// _ = "end of CoverTab[67079]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:246
		_go_fuzz_dep_.CoverTab[67080]++

													v := m.Get(fd)

													if !m.Has(fd) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:250
			_go_fuzz_dep_.CoverTab[67088]++
														if !w.EmitDefaults || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:251
				_go_fuzz_dep_.CoverTab[67090]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:251
				return fd.ContainingOneof() != nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:251
				// _ = "end of CoverTab[67090]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:251
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:251
				_go_fuzz_dep_.CoverTab[67091]++
															continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:252
				// _ = "end of CoverTab[67091]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:253
				_go_fuzz_dep_.CoverTab[67092]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:253
				// _ = "end of CoverTab[67092]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:253
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:253
			// _ = "end of CoverTab[67088]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:253
			_go_fuzz_dep_.CoverTab[67089]++
														if fd.Cardinality() != protoreflect.Repeated && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
				_go_fuzz_dep_.CoverTab[67093]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
				return (fd.Message() != nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
					_go_fuzz_dep_.CoverTab[67094]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
					return fd.Syntax() == protoreflect.Proto2
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
					// _ = "end of CoverTab[67094]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
				}())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
				// _ = "end of CoverTab[67093]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:254
				_go_fuzz_dep_.CoverTab[67095]++
															v = protoreflect.Value{}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:255
				// _ = "end of CoverTab[67095]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:256
				_go_fuzz_dep_.CoverTab[67096]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:256
				// _ = "end of CoverTab[67096]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:256
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:256
			// _ = "end of CoverTab[67089]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:257
			_go_fuzz_dep_.CoverTab[67097]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:257
			// _ = "end of CoverTab[67097]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:257
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:257
		// _ = "end of CoverTab[67080]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:257
		_go_fuzz_dep_.CoverTab[67081]++

													if !firstField {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:259
			_go_fuzz_dep_.CoverTab[67098]++
														w.writeComma()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:260
			// _ = "end of CoverTab[67098]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:261
			_go_fuzz_dep_.CoverTab[67099]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:261
			// _ = "end of CoverTab[67099]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:261
		// _ = "end of CoverTab[67081]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:261
		_go_fuzz_dep_.CoverTab[67082]++
													if err := w.marshalField(fd, v, indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:262
			_go_fuzz_dep_.CoverTab[67100]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:263
			// _ = "end of CoverTab[67100]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:264
			_go_fuzz_dep_.CoverTab[67101]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:264
			// _ = "end of CoverTab[67101]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:264
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:264
		// _ = "end of CoverTab[67082]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:264
		_go_fuzz_dep_.CoverTab[67083]++
													firstField = false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:265
		// _ = "end of CoverTab[67083]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:266
	// _ = "end of CoverTab[67020]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:266
	_go_fuzz_dep_.CoverTab[67021]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:269
	if md.ExtensionRanges().Len() > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:269
		_go_fuzz_dep_.CoverTab[67102]++
		// Collect a sorted list of all extension descriptor and values.
		type ext struct {
			desc	protoreflect.FieldDescriptor
			val	protoreflect.Value
		}
		var exts []ext
		m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:276
			_go_fuzz_dep_.CoverTab[67105]++
														if fd.IsExtension() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:277
				_go_fuzz_dep_.CoverTab[67107]++
															exts = append(exts, ext{fd, v})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:278
				// _ = "end of CoverTab[67107]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:279
				_go_fuzz_dep_.CoverTab[67108]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:279
				// _ = "end of CoverTab[67108]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:279
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:279
			// _ = "end of CoverTab[67105]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:279
			_go_fuzz_dep_.CoverTab[67106]++
														return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:280
			// _ = "end of CoverTab[67106]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:281
		// _ = "end of CoverTab[67102]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:281
		_go_fuzz_dep_.CoverTab[67103]++
													sort.Slice(exts, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:282
			_go_fuzz_dep_.CoverTab[67109]++
														return exts[i].desc.Number() < exts[j].desc.Number()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:283
			// _ = "end of CoverTab[67109]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:284
		// _ = "end of CoverTab[67103]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:284
		_go_fuzz_dep_.CoverTab[67104]++

													for _, ext := range exts {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:286
			_go_fuzz_dep_.CoverTab[67110]++
														if !firstField {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:287
				_go_fuzz_dep_.CoverTab[67113]++
															w.writeComma()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:288
				// _ = "end of CoverTab[67113]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:289
				_go_fuzz_dep_.CoverTab[67114]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:289
				// _ = "end of CoverTab[67114]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:289
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:289
			// _ = "end of CoverTab[67110]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:289
			_go_fuzz_dep_.CoverTab[67111]++
														if err := w.marshalField(ext.desc, ext.val, indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:290
				_go_fuzz_dep_.CoverTab[67115]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:291
				// _ = "end of CoverTab[67115]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:292
				_go_fuzz_dep_.CoverTab[67116]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:292
				// _ = "end of CoverTab[67116]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:292
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:292
			// _ = "end of CoverTab[67111]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:292
			_go_fuzz_dep_.CoverTab[67112]++
														firstField = false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:293
			// _ = "end of CoverTab[67112]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:294
		// _ = "end of CoverTab[67104]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:295
		_go_fuzz_dep_.CoverTab[67117]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:295
		// _ = "end of CoverTab[67117]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:295
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:295
	// _ = "end of CoverTab[67021]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:295
	_go_fuzz_dep_.CoverTab[67022]++

												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:297
		_go_fuzz_dep_.CoverTab[67118]++
													w.write("\n")
													w.write(indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:299
		// _ = "end of CoverTab[67118]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:300
		_go_fuzz_dep_.CoverTab[67119]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:300
		// _ = "end of CoverTab[67119]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:300
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:300
	// _ = "end of CoverTab[67022]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:300
	_go_fuzz_dep_.CoverTab[67023]++
												w.write("}")
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:302
	// _ = "end of CoverTab[67023]"
}

func (w *jsonWriter) writeComma() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:305
	_go_fuzz_dep_.CoverTab[67120]++
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:306
		_go_fuzz_dep_.CoverTab[67121]++
													w.write(",\n")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:307
		// _ = "end of CoverTab[67121]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:308
		_go_fuzz_dep_.CoverTab[67122]++
													w.write(",")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:309
		// _ = "end of CoverTab[67122]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:310
	// _ = "end of CoverTab[67120]"
}

func (w *jsonWriter) marshalAny(m protoreflect.Message, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:313
	_go_fuzz_dep_.CoverTab[67123]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:318
	md := m.Descriptor()
	typeURL := m.Get(md.Fields().ByNumber(1)).String()
	rawVal := m.Get(md.Fields().ByNumber(2)).Bytes()

	var m2 protoreflect.Message
	if w.AnyResolver != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:323
		_go_fuzz_dep_.CoverTab[67132]++
													mi, err := w.AnyResolver.Resolve(typeURL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:325
			_go_fuzz_dep_.CoverTab[67134]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:326
			// _ = "end of CoverTab[67134]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:327
			_go_fuzz_dep_.CoverTab[67135]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:327
			// _ = "end of CoverTab[67135]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:327
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:327
		// _ = "end of CoverTab[67132]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:327
		_go_fuzz_dep_.CoverTab[67133]++
													m2 = proto.MessageReflect(mi)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:328
		// _ = "end of CoverTab[67133]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:329
		_go_fuzz_dep_.CoverTab[67136]++
													mt, err := protoregistry.GlobalTypes.FindMessageByURL(typeURL)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:331
			_go_fuzz_dep_.CoverTab[67138]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:332
			// _ = "end of CoverTab[67138]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:333
			_go_fuzz_dep_.CoverTab[67139]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:333
			// _ = "end of CoverTab[67139]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:333
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:333
		// _ = "end of CoverTab[67136]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:333
		_go_fuzz_dep_.CoverTab[67137]++
													m2 = mt.New()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:334
		// _ = "end of CoverTab[67137]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:335
	// _ = "end of CoverTab[67123]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:335
	_go_fuzz_dep_.CoverTab[67124]++

												if err := protoV2.Unmarshal(rawVal, m2.Interface()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:337
		_go_fuzz_dep_.CoverTab[67140]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:338
		// _ = "end of CoverTab[67140]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:339
		_go_fuzz_dep_.CoverTab[67141]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:339
		// _ = "end of CoverTab[67141]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:339
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:339
	// _ = "end of CoverTab[67124]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:339
	_go_fuzz_dep_.CoverTab[67125]++

												if wellKnownType(m2.Descriptor().FullName()) == "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:341
		_go_fuzz_dep_.CoverTab[67142]++
													return w.marshalMessage(m2, indent, typeURL)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:342
		// _ = "end of CoverTab[67142]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:343
		_go_fuzz_dep_.CoverTab[67143]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:343
		// _ = "end of CoverTab[67143]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:343
	// _ = "end of CoverTab[67125]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:343
	_go_fuzz_dep_.CoverTab[67126]++

												w.write("{")
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:346
		_go_fuzz_dep_.CoverTab[67144]++
													w.write("\n")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:347
		// _ = "end of CoverTab[67144]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:348
		_go_fuzz_dep_.CoverTab[67145]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:348
		// _ = "end of CoverTab[67145]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:348
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:348
	// _ = "end of CoverTab[67126]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:348
	_go_fuzz_dep_.CoverTab[67127]++
												if err := w.marshalTypeURL(indent, typeURL); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:349
		_go_fuzz_dep_.CoverTab[67146]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:350
		// _ = "end of CoverTab[67146]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:351
		_go_fuzz_dep_.CoverTab[67147]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:351
		// _ = "end of CoverTab[67147]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:351
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:351
	// _ = "end of CoverTab[67127]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:351
	_go_fuzz_dep_.CoverTab[67128]++
												w.writeComma()
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:353
		_go_fuzz_dep_.CoverTab[67148]++
													w.write(indent)
													w.write(w.Indent)
													w.write(`"value": `)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:356
		// _ = "end of CoverTab[67148]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:357
		_go_fuzz_dep_.CoverTab[67149]++
													w.write(`"value":`)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:358
		// _ = "end of CoverTab[67149]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:359
	// _ = "end of CoverTab[67128]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:359
	_go_fuzz_dep_.CoverTab[67129]++
												if err := w.marshalMessage(m2, indent+w.Indent, ""); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:360
		_go_fuzz_dep_.CoverTab[67150]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:361
		// _ = "end of CoverTab[67150]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:362
		_go_fuzz_dep_.CoverTab[67151]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:362
		// _ = "end of CoverTab[67151]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:362
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:362
	// _ = "end of CoverTab[67129]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:362
	_go_fuzz_dep_.CoverTab[67130]++
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:363
		_go_fuzz_dep_.CoverTab[67152]++
													w.write("\n")
													w.write(indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:365
		// _ = "end of CoverTab[67152]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:366
		_go_fuzz_dep_.CoverTab[67153]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:366
		// _ = "end of CoverTab[67153]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:366
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:366
	// _ = "end of CoverTab[67130]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:366
	_go_fuzz_dep_.CoverTab[67131]++
												w.write("}")
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:368
	// _ = "end of CoverTab[67131]"
}

func (w *jsonWriter) marshalTypeURL(indent, typeURL string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:371
	_go_fuzz_dep_.CoverTab[67154]++
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:372
		_go_fuzz_dep_.CoverTab[67158]++
													w.write(indent)
													w.write(w.Indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:374
		// _ = "end of CoverTab[67158]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:375
		_go_fuzz_dep_.CoverTab[67159]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:375
		// _ = "end of CoverTab[67159]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:375
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:375
	// _ = "end of CoverTab[67154]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:375
	_go_fuzz_dep_.CoverTab[67155]++
												w.write(`"@type":`)
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:377
		_go_fuzz_dep_.CoverTab[67160]++
													w.write(" ")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:378
		// _ = "end of CoverTab[67160]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:379
		_go_fuzz_dep_.CoverTab[67161]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:379
		// _ = "end of CoverTab[67161]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:379
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:379
	// _ = "end of CoverTab[67155]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:379
	_go_fuzz_dep_.CoverTab[67156]++
												b, err := json.Marshal(typeURL)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:381
		_go_fuzz_dep_.CoverTab[67162]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:382
		// _ = "end of CoverTab[67162]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:383
		_go_fuzz_dep_.CoverTab[67163]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:383
		// _ = "end of CoverTab[67163]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:383
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:383
	// _ = "end of CoverTab[67156]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:383
	_go_fuzz_dep_.CoverTab[67157]++
												w.write(string(b))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:385
	// _ = "end of CoverTab[67157]"
}

// marshalField writes field description and value to the Writer.
func (w *jsonWriter) marshalField(fd protoreflect.FieldDescriptor, v protoreflect.Value, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:389
	_go_fuzz_dep_.CoverTab[67164]++
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:390
		_go_fuzz_dep_.CoverTab[67168]++
													w.write(indent)
													w.write(w.Indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:392
		// _ = "end of CoverTab[67168]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:393
		_go_fuzz_dep_.CoverTab[67169]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:393
		// _ = "end of CoverTab[67169]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:393
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:393
	// _ = "end of CoverTab[67164]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:393
	_go_fuzz_dep_.CoverTab[67165]++
												w.write(`"`)
												switch {
	case fd.IsExtension():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:396
		_go_fuzz_dep_.CoverTab[67170]++

													name := string(fd.FullName())
													if isMessageSet(fd.ContainingMessage()) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:399
			_go_fuzz_dep_.CoverTab[67175]++
														name = strings.TrimSuffix(name, ".message_set_extension")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:400
			// _ = "end of CoverTab[67175]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:401
			_go_fuzz_dep_.CoverTab[67176]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:401
			// _ = "end of CoverTab[67176]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:401
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:401
		// _ = "end of CoverTab[67170]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:401
		_go_fuzz_dep_.CoverTab[67171]++

													w.write("[" + name + "]")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:403
		// _ = "end of CoverTab[67171]"
	case w.OrigName:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:404
		_go_fuzz_dep_.CoverTab[67172]++
													name := string(fd.Name())
													if fd.Kind() == protoreflect.GroupKind {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:406
			_go_fuzz_dep_.CoverTab[67177]++
														name = string(fd.Message().Name())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:407
			// _ = "end of CoverTab[67177]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:408
			_go_fuzz_dep_.CoverTab[67178]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:408
			// _ = "end of CoverTab[67178]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:408
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:408
		// _ = "end of CoverTab[67172]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:408
		_go_fuzz_dep_.CoverTab[67173]++
													w.write(name)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:409
		// _ = "end of CoverTab[67173]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:410
		_go_fuzz_dep_.CoverTab[67174]++
													w.write(string(fd.JSONName()))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:411
		// _ = "end of CoverTab[67174]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:412
	// _ = "end of CoverTab[67165]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:412
	_go_fuzz_dep_.CoverTab[67166]++
												w.write(`":`)
												if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:414
		_go_fuzz_dep_.CoverTab[67179]++
													w.write(" ")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:415
		// _ = "end of CoverTab[67179]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:416
		_go_fuzz_dep_.CoverTab[67180]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:416
		// _ = "end of CoverTab[67180]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:416
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:416
	// _ = "end of CoverTab[67166]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:416
	_go_fuzz_dep_.CoverTab[67167]++
												return w.marshalValue(fd, v, indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:417
	// _ = "end of CoverTab[67167]"
}

func (w *jsonWriter) marshalValue(fd protoreflect.FieldDescriptor, v protoreflect.Value, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:420
	_go_fuzz_dep_.CoverTab[67181]++
												switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:422
		_go_fuzz_dep_.CoverTab[67182]++
													w.write("[")
													comma := ""
													lv := v.List()
													for i := 0; i < lv.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:426
			_go_fuzz_dep_.CoverTab[67191]++
														w.write(comma)
														if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:428
				_go_fuzz_dep_.CoverTab[67194]++
															w.write("\n")
															w.write(indent)
															w.write(w.Indent)
															w.write(w.Indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:432
				// _ = "end of CoverTab[67194]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:433
				_go_fuzz_dep_.CoverTab[67195]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:433
				// _ = "end of CoverTab[67195]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:433
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:433
			// _ = "end of CoverTab[67191]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:433
			_go_fuzz_dep_.CoverTab[67192]++
														if err := w.marshalSingularValue(fd, lv.Get(i), indent+w.Indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:434
				_go_fuzz_dep_.CoverTab[67196]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:435
				// _ = "end of CoverTab[67196]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:436
				_go_fuzz_dep_.CoverTab[67197]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:436
				// _ = "end of CoverTab[67197]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:436
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:436
			// _ = "end of CoverTab[67192]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:436
			_go_fuzz_dep_.CoverTab[67193]++
														comma = ","
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:437
			// _ = "end of CoverTab[67193]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:438
		// _ = "end of CoverTab[67182]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:438
		_go_fuzz_dep_.CoverTab[67183]++
													if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:439
			_go_fuzz_dep_.CoverTab[67198]++
														w.write("\n")
														w.write(indent)
														w.write(w.Indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:442
			// _ = "end of CoverTab[67198]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:443
			_go_fuzz_dep_.CoverTab[67199]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:443
			// _ = "end of CoverTab[67199]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:443
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:443
		// _ = "end of CoverTab[67183]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:443
		_go_fuzz_dep_.CoverTab[67184]++
													w.write("]")
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:445
		// _ = "end of CoverTab[67184]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:446
		_go_fuzz_dep_.CoverTab[67185]++
													kfd := fd.MapKey()
													vfd := fd.MapValue()
													mv := v.Map()

		// Collect a sorted list of all map keys and values.
		type entry struct{ key, val protoreflect.Value }
		var entries []entry
		mv.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:454
			_go_fuzz_dep_.CoverTab[67200]++
														entries = append(entries, entry{k.Value(), v})
														return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:456
			// _ = "end of CoverTab[67200]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:457
		// _ = "end of CoverTab[67185]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:457
		_go_fuzz_dep_.CoverTab[67186]++
													sort.Slice(entries, func(i, j int) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:458
			_go_fuzz_dep_.CoverTab[67201]++
														switch kfd.Kind() {
			case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:460
				_go_fuzz_dep_.CoverTab[67202]++
															return !entries[i].key.Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:461
					_go_fuzz_dep_.CoverTab[67207]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:461
					return entries[j].key.Bool()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:461
					// _ = "end of CoverTab[67207]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:461
				}()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:461
				// _ = "end of CoverTab[67202]"
			case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:462
				_go_fuzz_dep_.CoverTab[67203]++
															return entries[i].key.Int() < entries[j].key.Int()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:463
				// _ = "end of CoverTab[67203]"
			case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:464
				_go_fuzz_dep_.CoverTab[67204]++
															return entries[i].key.Uint() < entries[j].key.Uint()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:465
				// _ = "end of CoverTab[67204]"
			case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:466
				_go_fuzz_dep_.CoverTab[67205]++
															return entries[i].key.String() < entries[j].key.String()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:467
				// _ = "end of CoverTab[67205]"
			default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:468
				_go_fuzz_dep_.CoverTab[67206]++
															panic("invalid kind")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:469
				// _ = "end of CoverTab[67206]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:470
			// _ = "end of CoverTab[67201]"
		})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:471
		// _ = "end of CoverTab[67186]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:471
		_go_fuzz_dep_.CoverTab[67187]++

													w.write(`{`)
													comma := ""
													for _, entry := range entries {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:475
			_go_fuzz_dep_.CoverTab[67208]++
														w.write(comma)
														if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:477
				_go_fuzz_dep_.CoverTab[67213]++
															w.write("\n")
															w.write(indent)
															w.write(w.Indent)
															w.write(w.Indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:481
				// _ = "end of CoverTab[67213]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:482
				_go_fuzz_dep_.CoverTab[67214]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:482
				// _ = "end of CoverTab[67214]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:482
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:482
			// _ = "end of CoverTab[67208]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:482
			_go_fuzz_dep_.CoverTab[67209]++

														s := fmt.Sprint(entry.key.Interface())
														b, err := json.Marshal(s)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:486
				_go_fuzz_dep_.CoverTab[67215]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:487
				// _ = "end of CoverTab[67215]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:488
				_go_fuzz_dep_.CoverTab[67216]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:488
				// _ = "end of CoverTab[67216]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:488
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:488
			// _ = "end of CoverTab[67209]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:488
			_go_fuzz_dep_.CoverTab[67210]++
														w.write(string(b))

														w.write(`:`)
														if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:492
				_go_fuzz_dep_.CoverTab[67217]++
															w.write(` `)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:493
				// _ = "end of CoverTab[67217]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:494
				_go_fuzz_dep_.CoverTab[67218]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:494
				// _ = "end of CoverTab[67218]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:494
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:494
			// _ = "end of CoverTab[67210]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:494
			_go_fuzz_dep_.CoverTab[67211]++

														if err := w.marshalSingularValue(vfd, entry.val, indent+w.Indent); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:496
				_go_fuzz_dep_.CoverTab[67219]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:497
				// _ = "end of CoverTab[67219]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:498
				_go_fuzz_dep_.CoverTab[67220]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:498
				// _ = "end of CoverTab[67220]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:498
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:498
			// _ = "end of CoverTab[67211]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:498
			_go_fuzz_dep_.CoverTab[67212]++
														comma = ","
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:499
			// _ = "end of CoverTab[67212]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:500
		// _ = "end of CoverTab[67187]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:500
		_go_fuzz_dep_.CoverTab[67188]++
													if w.Indent != "" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:501
			_go_fuzz_dep_.CoverTab[67221]++
														w.write("\n")
														w.write(indent)
														w.write(w.Indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:504
			// _ = "end of CoverTab[67221]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:505
			_go_fuzz_dep_.CoverTab[67222]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:505
			// _ = "end of CoverTab[67222]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:505
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:505
		// _ = "end of CoverTab[67188]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:505
		_go_fuzz_dep_.CoverTab[67189]++
													w.write(`}`)
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:507
		// _ = "end of CoverTab[67189]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:508
		_go_fuzz_dep_.CoverTab[67190]++
													return w.marshalSingularValue(fd, v, indent)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:509
		// _ = "end of CoverTab[67190]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:510
	// _ = "end of CoverTab[67181]"
}

func (w *jsonWriter) marshalSingularValue(fd protoreflect.FieldDescriptor, v protoreflect.Value, indent string) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:513
	_go_fuzz_dep_.CoverTab[67223]++
												switch {
	case !v.IsValid():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:515
		_go_fuzz_dep_.CoverTab[67224]++
													w.write("null")
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:517
		// _ = "end of CoverTab[67224]"
	case fd.Message() != nil:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:518
		_go_fuzz_dep_.CoverTab[67225]++
													return w.marshalMessage(v.Message(), indent+w.Indent, "")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:519
		// _ = "end of CoverTab[67225]"
	case fd.Enum() != nil:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:520
		_go_fuzz_dep_.CoverTab[67226]++
													if fd.Enum().FullName() == "google.protobuf.NullValue" {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:521
			_go_fuzz_dep_.CoverTab[67232]++
														w.write("null")
														return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:523
			// _ = "end of CoverTab[67232]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:524
			_go_fuzz_dep_.CoverTab[67233]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:524
			// _ = "end of CoverTab[67233]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:524
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:524
		// _ = "end of CoverTab[67226]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:524
		_go_fuzz_dep_.CoverTab[67227]++

													vd := fd.Enum().Values().ByNumber(v.Enum())
													if vd == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:527
			_go_fuzz_dep_.CoverTab[67234]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:527
			return w.EnumsAsInts
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:527
			// _ = "end of CoverTab[67234]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:527
		}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:527
			_go_fuzz_dep_.CoverTab[67235]++
														w.write(strconv.Itoa(int(v.Enum())))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:528
			// _ = "end of CoverTab[67235]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:529
			_go_fuzz_dep_.CoverTab[67236]++
														w.write(`"` + string(vd.Name()) + `"`)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:530
			// _ = "end of CoverTab[67236]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:531
		// _ = "end of CoverTab[67227]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:531
		_go_fuzz_dep_.CoverTab[67228]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:532
		// _ = "end of CoverTab[67228]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:533
		_go_fuzz_dep_.CoverTab[67229]++
													switch v.Interface().(type) {
		case float32, float64:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:535
			_go_fuzz_dep_.CoverTab[67237]++
														switch {
			case math.IsInf(v.Float(), +1):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:537
				_go_fuzz_dep_.CoverTab[67239]++
															w.write(`"Infinity"`)
															return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:539
				// _ = "end of CoverTab[67239]"
			case math.IsInf(v.Float(), -1):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:540
				_go_fuzz_dep_.CoverTab[67240]++
															w.write(`"-Infinity"`)
															return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:542
				// _ = "end of CoverTab[67240]"
			case math.IsNaN(v.Float()):
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:543
				_go_fuzz_dep_.CoverTab[67241]++
															w.write(`"NaN"`)
															return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:545
				// _ = "end of CoverTab[67241]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:545
			default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:545
				_go_fuzz_dep_.CoverTab[67242]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:545
				// _ = "end of CoverTab[67242]"
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:546
			// _ = "end of CoverTab[67237]"
		case int64, uint64:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:547
			_go_fuzz_dep_.CoverTab[67238]++
														w.write(fmt.Sprintf(`"%d"`, v.Interface()))
														return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:549
			// _ = "end of CoverTab[67238]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:550
		// _ = "end of CoverTab[67229]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:550
		_go_fuzz_dep_.CoverTab[67230]++

													b, err := json.Marshal(v.Interface())
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:553
			_go_fuzz_dep_.CoverTab[67243]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:554
			// _ = "end of CoverTab[67243]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:555
			_go_fuzz_dep_.CoverTab[67244]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:555
			// _ = "end of CoverTab[67244]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:555
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:555
		// _ = "end of CoverTab[67230]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:555
		_go_fuzz_dep_.CoverTab[67231]++
													w.write(string(b))
													return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:557
		// _ = "end of CoverTab[67231]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:558
	// _ = "end of CoverTab[67223]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:559
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/jsonpb/encode.go:559
var _ = _go_fuzz_dep_.CoverTab
