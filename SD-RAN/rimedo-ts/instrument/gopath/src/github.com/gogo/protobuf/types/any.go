// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2016 The Go Authors.  All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
package types

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:32
)

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:37
import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gogo/protobuf/proto"
)

const googleApis = "type.googleapis.com/"

// AnyMessageName returns the name of the message contained in a google.protobuf.Any message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:47
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:47
// Note that regular type assertions should be done using the Is
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:47
// function. AnyMessageName is provided for less common use cases like filtering a
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:47
// sequence of Any messages based on a set of allowed message type names.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:52
func AnyMessageName(any *Any) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:52
	_go_fuzz_dep_.CoverTab[134219]++
											if any == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:53
		_go_fuzz_dep_.CoverTab[134222]++
												return "", fmt.Errorf("message is nil")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:54
		// _ = "end of CoverTab[134222]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:55
		_go_fuzz_dep_.CoverTab[134223]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:55
		// _ = "end of CoverTab[134223]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:55
	// _ = "end of CoverTab[134219]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:55
	_go_fuzz_dep_.CoverTab[134220]++
											slash := strings.LastIndex(any.TypeUrl, "/")
											if slash < 0 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:57
		_go_fuzz_dep_.CoverTab[134224]++
												return "", fmt.Errorf("message type url %q is invalid", any.TypeUrl)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:58
		// _ = "end of CoverTab[134224]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:59
		_go_fuzz_dep_.CoverTab[134225]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:59
		// _ = "end of CoverTab[134225]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:59
	// _ = "end of CoverTab[134220]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:59
	_go_fuzz_dep_.CoverTab[134221]++
											return any.TypeUrl[slash+1:], nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:60
	// _ = "end of CoverTab[134221]"
}

// MarshalAny takes the protocol buffer and encodes it into google.protobuf.Any.
func MarshalAny(pb proto.Message) (*Any, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:64
	_go_fuzz_dep_.CoverTab[134226]++
											value, err := proto.Marshal(pb)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:66
		_go_fuzz_dep_.CoverTab[134228]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:67
		// _ = "end of CoverTab[134228]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:68
		_go_fuzz_dep_.CoverTab[134229]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:68
		// _ = "end of CoverTab[134229]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:68
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:68
	// _ = "end of CoverTab[134226]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:68
	_go_fuzz_dep_.CoverTab[134227]++
											return &Any{TypeUrl: googleApis + proto.MessageName(pb), Value: value}, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:69
	// _ = "end of CoverTab[134227]"
}

// DynamicAny is a value that can be passed to UnmarshalAny to automatically
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
// allocate a proto.Message for the type specified in a google.protobuf.Any
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
// message. The allocated message is stored in the embedded proto.Message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
// Example:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
//	var x ptypes.DynamicAny
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
//	if err := ptypes.UnmarshalAny(a, &x); err != nil { ... }
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:72
//	fmt.Printf("unmarshaled message: %v", x.Message)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:81
type DynamicAny struct {
	proto.Message
}

// Empty returns a new proto.Message of the type specified in a
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:85
// google.protobuf.Any message. It returns an error if corresponding message
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:85
// type isn't linked in.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:88
func EmptyAny(any *Any) (proto.Message, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:88
	_go_fuzz_dep_.CoverTab[134230]++
											aname, err := AnyMessageName(any)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:90
		_go_fuzz_dep_.CoverTab[134233]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:91
		// _ = "end of CoverTab[134233]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:92
		_go_fuzz_dep_.CoverTab[134234]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:92
		// _ = "end of CoverTab[134234]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:92
	// _ = "end of CoverTab[134230]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:92
	_go_fuzz_dep_.CoverTab[134231]++

											t := proto.MessageType(aname)
											if t == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:95
		_go_fuzz_dep_.CoverTab[134235]++
												return nil, fmt.Errorf("any: message type %q isn't linked in", aname)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:96
		// _ = "end of CoverTab[134235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:97
		_go_fuzz_dep_.CoverTab[134236]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:97
		// _ = "end of CoverTab[134236]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:97
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:97
	// _ = "end of CoverTab[134231]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:97
	_go_fuzz_dep_.CoverTab[134232]++
											return reflect.New(t.Elem()).Interface().(proto.Message), nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:98
	// _ = "end of CoverTab[134232]"
}

// UnmarshalAny parses the protocol buffer representation in a google.protobuf.Any
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:101
// message and places the decoded result in pb. It returns an error if type of
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:101
// contents of Any message does not match type of pb message.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:101
//
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:101
// pb can be a proto.Message, or a *DynamicAny.
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:106
func UnmarshalAny(any *Any, pb proto.Message) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:106
	_go_fuzz_dep_.CoverTab[134237]++
											if d, ok := pb.(*DynamicAny); ok {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:107
		_go_fuzz_dep_.CoverTab[134241]++
												if d.Message == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:108
			_go_fuzz_dep_.CoverTab[134243]++
													var err error
													d.Message, err = EmptyAny(any)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:111
				_go_fuzz_dep_.CoverTab[134244]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:112
				// _ = "end of CoverTab[134244]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:113
				_go_fuzz_dep_.CoverTab[134245]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:113
				// _ = "end of CoverTab[134245]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:113
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:113
			// _ = "end of CoverTab[134243]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:114
			_go_fuzz_dep_.CoverTab[134246]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:114
			// _ = "end of CoverTab[134246]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:114
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:114
		// _ = "end of CoverTab[134241]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:114
		_go_fuzz_dep_.CoverTab[134242]++
												return UnmarshalAny(any, d.Message)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:115
		// _ = "end of CoverTab[134242]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:116
		_go_fuzz_dep_.CoverTab[134247]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:116
		// _ = "end of CoverTab[134247]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:116
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:116
	// _ = "end of CoverTab[134237]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:116
	_go_fuzz_dep_.CoverTab[134238]++

											aname, err := AnyMessageName(any)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:119
		_go_fuzz_dep_.CoverTab[134248]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:120
		// _ = "end of CoverTab[134248]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:121
		_go_fuzz_dep_.CoverTab[134249]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:121
		// _ = "end of CoverTab[134249]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:121
	// _ = "end of CoverTab[134238]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:121
	_go_fuzz_dep_.CoverTab[134239]++

											mname := proto.MessageName(pb)
											if aname != mname {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:124
		_go_fuzz_dep_.CoverTab[134250]++
												return fmt.Errorf("mismatched message type: got %q want %q", aname, mname)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:125
		// _ = "end of CoverTab[134250]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:126
		_go_fuzz_dep_.CoverTab[134251]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:126
		// _ = "end of CoverTab[134251]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:126
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:126
	// _ = "end of CoverTab[134239]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:126
	_go_fuzz_dep_.CoverTab[134240]++
											return proto.Unmarshal(any.Value, pb)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:127
	// _ = "end of CoverTab[134240]"
}

// Is returns true if any value contains a given message type.
func Is(any *Any, pb proto.Message) bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:131
	_go_fuzz_dep_.CoverTab[134252]++

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:134
	if any == nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:134
		_go_fuzz_dep_.CoverTab[134254]++
												return false
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:135
		// _ = "end of CoverTab[134254]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:136
		_go_fuzz_dep_.CoverTab[134255]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:136
		// _ = "end of CoverTab[134255]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:136
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:136
	// _ = "end of CoverTab[134252]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:136
	_go_fuzz_dep_.CoverTab[134253]++
											name := proto.MessageName(pb)
											prefix := len(any.TypeUrl) - len(name)
											return prefix >= 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
		_go_fuzz_dep_.CoverTab[134256]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
		return any.TypeUrl[prefix-1] == '/'
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
		// _ = "end of CoverTab[134256]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
		_go_fuzz_dep_.CoverTab[134257]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
		return any.TypeUrl[prefix:] == name
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
		// _ = "end of CoverTab[134257]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
	}()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:139
	// _ = "end of CoverTab[134253]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:140
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/any.go:140
var _ = _go_fuzz_dep_.CoverTab
