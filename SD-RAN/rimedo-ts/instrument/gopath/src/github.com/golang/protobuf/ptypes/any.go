// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
package ptypes

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:5
)

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	anypb "github.com/golang/protobuf/ptypes/any"
)

const urlPrefix = "type.googleapis.com/"

// AnyMessageName returns the message name contained in an anypb.Any message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:20
// Most type assertions should use the Is function instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:20
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:20
// Deprecated: Call the any.MessageName method instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:24
func AnyMessageName(any *anypb.Any) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:24
	_go_fuzz_dep_.CoverTab[68016]++
											name, err := anyMessageName(any)
											return string(name), err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:26
	// _ = "end of CoverTab[68016]"
}
func anyMessageName(any *anypb.Any) (protoreflect.FullName, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:28
	_go_fuzz_dep_.CoverTab[68017]++
											if any == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:29
		_go_fuzz_dep_.CoverTab[68021]++
												return "", fmt.Errorf("message is nil")
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:30
		// _ = "end of CoverTab[68021]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:31
		_go_fuzz_dep_.CoverTab[68022]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:31
		// _ = "end of CoverTab[68022]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:31
	// _ = "end of CoverTab[68017]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:31
	_go_fuzz_dep_.CoverTab[68018]++
											name := protoreflect.FullName(any.TypeUrl)
											if i := strings.LastIndex(any.TypeUrl, "/"); i >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:33
		_go_fuzz_dep_.CoverTab[68023]++
												name = name[i+len("/"):]
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:34
		// _ = "end of CoverTab[68023]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:35
		_go_fuzz_dep_.CoverTab[68024]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:35
		// _ = "end of CoverTab[68024]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:35
	// _ = "end of CoverTab[68018]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:35
	_go_fuzz_dep_.CoverTab[68019]++
											if !name.IsValid() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:36
		_go_fuzz_dep_.CoverTab[68025]++
												return "", fmt.Errorf("message type url %q is invalid", any.TypeUrl)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:37
		// _ = "end of CoverTab[68025]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:38
		_go_fuzz_dep_.CoverTab[68026]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:38
		// _ = "end of CoverTab[68026]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:38
	// _ = "end of CoverTab[68019]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:38
	_go_fuzz_dep_.CoverTab[68020]++
											return name, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:39
	// _ = "end of CoverTab[68020]"
}

// MarshalAny marshals the given message m into an anypb.Any message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:42
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:42
// Deprecated: Call the anypb.New function instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:45
func MarshalAny(m proto.Message) (*anypb.Any, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:45
	_go_fuzz_dep_.CoverTab[68027]++
											switch dm := m.(type) {
	case DynamicAny:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:47
		_go_fuzz_dep_.CoverTab[68030]++
												m = dm.Message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:48
		// _ = "end of CoverTab[68030]"
	case *DynamicAny:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:49
		_go_fuzz_dep_.CoverTab[68031]++
												if dm == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:50
			_go_fuzz_dep_.CoverTab[68033]++
													return nil, proto.ErrNil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:51
			// _ = "end of CoverTab[68033]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:52
			_go_fuzz_dep_.CoverTab[68034]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:52
			// _ = "end of CoverTab[68034]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:52
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:52
		// _ = "end of CoverTab[68031]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:52
		_go_fuzz_dep_.CoverTab[68032]++
												m = dm.Message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:53
		// _ = "end of CoverTab[68032]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:54
	// _ = "end of CoverTab[68027]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:54
	_go_fuzz_dep_.CoverTab[68028]++
											b, err := proto.Marshal(m)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:56
		_go_fuzz_dep_.CoverTab[68035]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:57
		// _ = "end of CoverTab[68035]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:58
		_go_fuzz_dep_.CoverTab[68036]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:58
		// _ = "end of CoverTab[68036]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:58
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:58
	// _ = "end of CoverTab[68028]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:58
	_go_fuzz_dep_.CoverTab[68029]++
											return &anypb.Any{TypeUrl: urlPrefix + proto.MessageName(m), Value: b}, nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:59
	// _ = "end of CoverTab[68029]"
}

// Empty returns a new message of the type specified in an anypb.Any message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:62
// It returns protoregistry.NotFound if the corresponding message type could not
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:62
// be resolved in the global registry.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:62
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:62
// Deprecated: Use protoregistry.GlobalTypes.FindMessageByName instead
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:62
// to resolve the message name and create a new instance of it.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:68
func Empty(any *anypb.Any) (proto.Message, error) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:68
	_go_fuzz_dep_.CoverTab[68037]++
											name, err := anyMessageName(any)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:70
		_go_fuzz_dep_.CoverTab[68040]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:71
		// _ = "end of CoverTab[68040]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:72
		_go_fuzz_dep_.CoverTab[68041]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:72
		// _ = "end of CoverTab[68041]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:72
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:72
	// _ = "end of CoverTab[68037]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:72
	_go_fuzz_dep_.CoverTab[68038]++
											mt, err := protoregistry.GlobalTypes.FindMessageByName(name)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:74
		_go_fuzz_dep_.CoverTab[68042]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:75
		// _ = "end of CoverTab[68042]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:76
		_go_fuzz_dep_.CoverTab[68043]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:76
		// _ = "end of CoverTab[68043]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:76
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:76
	// _ = "end of CoverTab[68038]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:76
	_go_fuzz_dep_.CoverTab[68039]++
											return proto.MessageV1(mt.New().Interface()), nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:77
	// _ = "end of CoverTab[68039]"
}

// UnmarshalAny unmarshals the encoded value contained in the anypb.Any message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
// into the provided message m. It returns an error if the target message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
// does not match the type in the Any message or if an unmarshal error occurs.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
// The target message m may be a *DynamicAny message. If the underlying message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
// type could not be resolved, then this returns protoregistry.NotFound.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:80
// Deprecated: Call the any.UnmarshalTo method instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:88
func UnmarshalAny(any *anypb.Any, m proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:88
	_go_fuzz_dep_.CoverTab[68044]++
											if dm, ok := m.(*DynamicAny); ok {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:89
		_go_fuzz_dep_.CoverTab[68048]++
												if dm.Message == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:90
			_go_fuzz_dep_.CoverTab[68050]++
													var err error
													dm.Message, err = Empty(any)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:93
				_go_fuzz_dep_.CoverTab[68051]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:94
				// _ = "end of CoverTab[68051]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:95
				_go_fuzz_dep_.CoverTab[68052]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:95
				// _ = "end of CoverTab[68052]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:95
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:95
			// _ = "end of CoverTab[68050]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:96
			_go_fuzz_dep_.CoverTab[68053]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:96
			// _ = "end of CoverTab[68053]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:96
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:96
		// _ = "end of CoverTab[68048]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:96
		_go_fuzz_dep_.CoverTab[68049]++
												m = dm.Message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:97
		// _ = "end of CoverTab[68049]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:98
		_go_fuzz_dep_.CoverTab[68054]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:98
		// _ = "end of CoverTab[68054]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:98
	// _ = "end of CoverTab[68044]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:98
	_go_fuzz_dep_.CoverTab[68045]++

											anyName, err := AnyMessageName(any)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:101
		_go_fuzz_dep_.CoverTab[68055]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:102
		// _ = "end of CoverTab[68055]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:103
		_go_fuzz_dep_.CoverTab[68056]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:103
		// _ = "end of CoverTab[68056]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:103
	// _ = "end of CoverTab[68045]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:103
	_go_fuzz_dep_.CoverTab[68046]++
											msgName := proto.MessageName(m)
											if anyName != msgName {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:105
		_go_fuzz_dep_.CoverTab[68057]++
												return fmt.Errorf("mismatched message type: got %q want %q", anyName, msgName)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:106
		// _ = "end of CoverTab[68057]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:107
		_go_fuzz_dep_.CoverTab[68058]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:107
		// _ = "end of CoverTab[68058]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:107
	// _ = "end of CoverTab[68046]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:107
	_go_fuzz_dep_.CoverTab[68047]++
											return proto.Unmarshal(any.Value, m)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:108
	// _ = "end of CoverTab[68047]"
}

// Is reports whether the Any message contains a message of the specified type.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:111
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:111
// Deprecated: Call the any.MessageIs method instead.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:114
func Is(any *anypb.Any, m proto.Message) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:114
	_go_fuzz_dep_.CoverTab[68059]++
											if any == nil || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:115
		_go_fuzz_dep_.CoverTab[68062]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:115
		return m == nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:115
		// _ = "end of CoverTab[68062]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:115
	}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:115
		_go_fuzz_dep_.CoverTab[68063]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:116
		// _ = "end of CoverTab[68063]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:117
		_go_fuzz_dep_.CoverTab[68064]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:117
		// _ = "end of CoverTab[68064]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:117
	// _ = "end of CoverTab[68059]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:117
	_go_fuzz_dep_.CoverTab[68060]++
											name := proto.MessageName(m)
											if !strings.HasSuffix(any.TypeUrl, name) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:119
		_go_fuzz_dep_.CoverTab[68065]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:120
		// _ = "end of CoverTab[68065]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:121
		_go_fuzz_dep_.CoverTab[68066]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:121
		// _ = "end of CoverTab[68066]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:121
	// _ = "end of CoverTab[68060]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:121
	_go_fuzz_dep_.CoverTab[68061]++
											return len(any.TypeUrl) == len(name) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:122
		_go_fuzz_dep_.CoverTab[68067]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:122
		return any.TypeUrl[len(any.TypeUrl)-len(name)-1] == '/'
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:122
		// _ = "end of CoverTab[68067]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:122
	}()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:122
	// _ = "end of CoverTab[68061]"
}

// DynamicAny is a value that can be passed to UnmarshalAny to automatically
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
// allocate a proto.Message for the type specified in an anypb.Any message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
// The allocated message is stored in the embedded proto.Message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
// Example:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
//	var x ptypes.DynamicAny
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
//	if err := ptypes.UnmarshalAny(a, &x); err != nil { ... }
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
//	fmt.Printf("unmarshaled message: %v", x.Message)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
// Deprecated: Use the any.UnmarshalNew method instead to unmarshal
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:125
// the any message contents into a new instance of the underlying message.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:136
type DynamicAny struct{ proto.Message }

func (m DynamicAny) String() string {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:138
	_go_fuzz_dep_.CoverTab[68068]++
											if m.Message == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:139
		_go_fuzz_dep_.CoverTab[68070]++
												return "<nil>"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:140
		// _ = "end of CoverTab[68070]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:141
		_go_fuzz_dep_.CoverTab[68071]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:141
		// _ = "end of CoverTab[68071]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:141
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:141
	// _ = "end of CoverTab[68068]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:141
	_go_fuzz_dep_.CoverTab[68069]++
											return m.Message.String()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:142
	// _ = "end of CoverTab[68069]"
}
func (m DynamicAny) Reset() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:144
	_go_fuzz_dep_.CoverTab[68072]++
											if m.Message == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:145
		_go_fuzz_dep_.CoverTab[68074]++
												return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:146
		// _ = "end of CoverTab[68074]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:147
		_go_fuzz_dep_.CoverTab[68075]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:147
		// _ = "end of CoverTab[68075]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:147
	// _ = "end of CoverTab[68072]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:147
	_go_fuzz_dep_.CoverTab[68073]++
											m.Message.Reset()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:148
	// _ = "end of CoverTab[68073]"
}
func (m DynamicAny) ProtoMessage() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:150
	_go_fuzz_dep_.CoverTab[68076]++
											return
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:151
	// _ = "end of CoverTab[68076]"
}
func (m DynamicAny) ProtoReflect() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:153
	_go_fuzz_dep_.CoverTab[68077]++
											if m.Message == nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:154
		_go_fuzz_dep_.CoverTab[68079]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:155
		// _ = "end of CoverTab[68079]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:156
		_go_fuzz_dep_.CoverTab[68080]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:156
		// _ = "end of CoverTab[68080]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:156
	// _ = "end of CoverTab[68077]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:156
	_go_fuzz_dep_.CoverTab[68078]++
											return dynamicAny{proto.MessageReflect(m.Message)}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:157
	// _ = "end of CoverTab[68078]"
}

type dynamicAny struct{ protoreflect.Message }

func (m dynamicAny) Type() protoreflect.MessageType {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:162
	_go_fuzz_dep_.CoverTab[68081]++
											return dynamicAnyType{m.Message.Type()}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:163
	// _ = "end of CoverTab[68081]"
}
func (m dynamicAny) New() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:165
	_go_fuzz_dep_.CoverTab[68082]++
											return dynamicAnyType{m.Message.Type()}.New()
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:166
	// _ = "end of CoverTab[68082]"
}
func (m dynamicAny) Interface() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:168
	_go_fuzz_dep_.CoverTab[68083]++
											return DynamicAny{proto.MessageV1(m.Message.Interface())}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:169
	// _ = "end of CoverTab[68083]"
}

type dynamicAnyType struct{ protoreflect.MessageType }

func (t dynamicAnyType) New() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:174
	_go_fuzz_dep_.CoverTab[68084]++
											return dynamicAny{t.MessageType.New()}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:175
	// _ = "end of CoverTab[68084]"
}
func (t dynamicAnyType) Zero() protoreflect.Message {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:177
	_go_fuzz_dep_.CoverTab[68085]++
											return dynamicAny{t.MessageType.Zero()}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:178
	// _ = "end of CoverTab[68085]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:179
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/ptypes/any.go:179
var _ = _go_fuzz_dep_.CoverTab
