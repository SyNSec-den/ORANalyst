// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:5
)

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// DiscardUnknown recursively discards all unknown fields from this message
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
// and all embedded messages.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
//
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
// When unmarshaling a message with unrecognized fields, the tags and values
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
// of such fields are preserved in the Message. This allows a later call to
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
// marshal to be able to produce a message that continues to have those
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
// unrecognized fields. To avoid this, DiscardUnknown is used to
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:11
// explicitly clear the unknown fields after unmarshaling.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:19
func DiscardUnknown(m Message) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:19
	_go_fuzz_dep_.CoverTab[61251]++
												if m != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:20
		_go_fuzz_dep_.CoverTab[61252]++
													discardUnknown(MessageReflect(m))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:21
		// _ = "end of CoverTab[61252]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:22
		_go_fuzz_dep_.CoverTab[61253]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:22
		// _ = "end of CoverTab[61253]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:22
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:22
	// _ = "end of CoverTab[61251]"
}

func discardUnknown(m protoreflect.Message) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:25
	_go_fuzz_dep_.CoverTab[61254]++
												m.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:26
		_go_fuzz_dep_.CoverTab[61256]++
													switch {

		case fd.Cardinality() != protoreflect.Repeated:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:29
			_go_fuzz_dep_.CoverTab[61258]++
														if fd.Message() != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:30
				_go_fuzz_dep_.CoverTab[61262]++
															discardUnknown(m.Get(fd).Message())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:31
				// _ = "end of CoverTab[61262]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:32
				_go_fuzz_dep_.CoverTab[61263]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:32
				// _ = "end of CoverTab[61263]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:32
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:32
			// _ = "end of CoverTab[61258]"

		case fd.IsList():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:34
			_go_fuzz_dep_.CoverTab[61259]++
														if fd.Message() != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:35
				_go_fuzz_dep_.CoverTab[61264]++
															ls := m.Get(fd).List()
															for i := 0; i < ls.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:37
					_go_fuzz_dep_.CoverTab[61265]++
																discardUnknown(ls.Get(i).Message())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:38
					// _ = "end of CoverTab[61265]"
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:39
				// _ = "end of CoverTab[61264]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:40
				_go_fuzz_dep_.CoverTab[61266]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:40
				// _ = "end of CoverTab[61266]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:40
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:40
			// _ = "end of CoverTab[61259]"

		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:42
			_go_fuzz_dep_.CoverTab[61260]++
														if fd.MapValue().Message() != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:43
				_go_fuzz_dep_.CoverTab[61267]++
															ms := m.Get(fd).Map()
															ms.Range(func(_ protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:45
					_go_fuzz_dep_.CoverTab[61268]++
																discardUnknown(v.Message())
																return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:47
					// _ = "end of CoverTab[61268]"
				})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:48
				// _ = "end of CoverTab[61267]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
				_go_fuzz_dep_.CoverTab[61269]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
				// _ = "end of CoverTab[61269]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
			// _ = "end of CoverTab[61260]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
			_go_fuzz_dep_.CoverTab[61261]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:49
			// _ = "end of CoverTab[61261]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:50
		// _ = "end of CoverTab[61256]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:50
		_go_fuzz_dep_.CoverTab[61257]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:51
		// _ = "end of CoverTab[61257]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:52
	// _ = "end of CoverTab[61254]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:52
	_go_fuzz_dep_.CoverTab[61255]++

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:55
	if len(m.GetUnknown()) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:55
		_go_fuzz_dep_.CoverTab[61270]++
													m.SetUnknown(nil)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:56
		// _ = "end of CoverTab[61270]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:57
		_go_fuzz_dep_.CoverTab[61271]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:57
		// _ = "end of CoverTab[61271]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:57
	// _ = "end of CoverTab[61255]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:58
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/discard.go:58
var _ = _go_fuzz_dep_.CoverTab
