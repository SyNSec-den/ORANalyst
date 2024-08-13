// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
package proto

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:5
)

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// SetDefaults sets unpopulated scalar fields to their default values.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:11
// Fields within a oneof are not set even if they have a default value.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:11
// SetDefaults is recursively called upon any populated message fields.
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:14
func SetDefaults(m Message) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:14
	_go_fuzz_dep_.CoverTab[61194]++
												if m != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:15
		_go_fuzz_dep_.CoverTab[61195]++
													setDefaults(MessageReflect(m))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:16
		// _ = "end of CoverTab[61195]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:17
		_go_fuzz_dep_.CoverTab[61196]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:17
		// _ = "end of CoverTab[61196]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:17
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:17
	// _ = "end of CoverTab[61194]"
}

func setDefaults(m protoreflect.Message) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:20
	_go_fuzz_dep_.CoverTab[61197]++
												fds := m.Descriptor().Fields()
												for i := 0; i < fds.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:22
		_go_fuzz_dep_.CoverTab[61199]++
													fd := fds.Get(i)
													if !m.Has(fd) {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:24
			_go_fuzz_dep_.CoverTab[61200]++
														if fd.HasDefault() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:25
				_go_fuzz_dep_.CoverTab[61202]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:25
				return fd.ContainingOneof() == nil
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:25
				// _ = "end of CoverTab[61202]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:25
			}() {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:25
				_go_fuzz_dep_.CoverTab[61203]++
															v := fd.Default()
															if fd.Kind() == protoreflect.BytesKind {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:27
					_go_fuzz_dep_.CoverTab[61205]++
																v = protoreflect.ValueOf(append([]byte(nil), v.Bytes()...))
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:28
					// _ = "end of CoverTab[61205]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:29
					_go_fuzz_dep_.CoverTab[61206]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:29
					// _ = "end of CoverTab[61206]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:29
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:29
				// _ = "end of CoverTab[61203]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:29
				_go_fuzz_dep_.CoverTab[61204]++
															m.Set(fd, v)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:30
				// _ = "end of CoverTab[61204]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:31
				_go_fuzz_dep_.CoverTab[61207]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:31
				// _ = "end of CoverTab[61207]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:31
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:31
			// _ = "end of CoverTab[61200]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:31
			_go_fuzz_dep_.CoverTab[61201]++
														continue
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:32
			// _ = "end of CoverTab[61201]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:33
			_go_fuzz_dep_.CoverTab[61208]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:33
			// _ = "end of CoverTab[61208]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:33
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:33
		// _ = "end of CoverTab[61199]"
	}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:34
	// _ = "end of CoverTab[61197]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:34
	_go_fuzz_dep_.CoverTab[61198]++

												m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:36
		_go_fuzz_dep_.CoverTab[61209]++
													switch {

		case fd.Cardinality() != protoreflect.Repeated:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:39
			_go_fuzz_dep_.CoverTab[61211]++
														if fd.Message() != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:40
				_go_fuzz_dep_.CoverTab[61215]++
															setDefaults(m.Get(fd).Message())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:41
				// _ = "end of CoverTab[61215]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:42
				_go_fuzz_dep_.CoverTab[61216]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:42
				// _ = "end of CoverTab[61216]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:42
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:42
			// _ = "end of CoverTab[61211]"

		case fd.IsList():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:44
			_go_fuzz_dep_.CoverTab[61212]++
														if fd.Message() != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:45
				_go_fuzz_dep_.CoverTab[61217]++
															ls := m.Get(fd).List()
															for i := 0; i < ls.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:47
					_go_fuzz_dep_.CoverTab[61218]++
																setDefaults(ls.Get(i).Message())
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:48
					// _ = "end of CoverTab[61218]"
				}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:49
				// _ = "end of CoverTab[61217]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:50
				_go_fuzz_dep_.CoverTab[61219]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:50
				// _ = "end of CoverTab[61219]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:50
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:50
			// _ = "end of CoverTab[61212]"

		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:52
			_go_fuzz_dep_.CoverTab[61213]++
														if fd.MapValue().Message() != nil {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:53
				_go_fuzz_dep_.CoverTab[61220]++
															ms := m.Get(fd).Map()
															ms.Range(func(_ protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:55
					_go_fuzz_dep_.CoverTab[61221]++
																setDefaults(v.Message())
																return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:57
					// _ = "end of CoverTab[61221]"
				})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:58
				// _ = "end of CoverTab[61220]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
				_go_fuzz_dep_.CoverTab[61222]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
				// _ = "end of CoverTab[61222]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
			}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
			// _ = "end of CoverTab[61213]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
		default:
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
			_go_fuzz_dep_.CoverTab[61214]++
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:59
			// _ = "end of CoverTab[61214]"
		}
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:60
		// _ = "end of CoverTab[61209]"
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:60
		_go_fuzz_dep_.CoverTab[61210]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:61
		// _ = "end of CoverTab[61210]"
	})
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:62
	// _ = "end of CoverTab[61198]"
}

//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/golang/protobuf@v1.5.3/proto/defaults.go:63
var _ = _go_fuzz_dep_.CoverTab
