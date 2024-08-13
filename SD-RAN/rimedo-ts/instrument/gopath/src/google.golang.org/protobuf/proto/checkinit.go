// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:5
)

import (
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// CheckInitialized returns an error if any required fields in m are not set.
func CheckInitialized(m Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:14
	_go_fuzz_dep_.CoverTab[50596]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:17
	if m == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:17
		_go_fuzz_dep_.CoverTab[50598]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:18
		// _ = "end of CoverTab[50598]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:19
		_go_fuzz_dep_.CoverTab[50599]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:19
		// _ = "end of CoverTab[50599]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:19
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:19
	// _ = "end of CoverTab[50596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:19
	_go_fuzz_dep_.CoverTab[50597]++

												return checkInitialized(m.ProtoReflect())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:21
	// _ = "end of CoverTab[50597]"
}

// CheckInitialized returns an error if any required fields in m are not set.
func checkInitialized(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:25
	_go_fuzz_dep_.CoverTab[50600]++
												if methods := protoMethods(m); methods != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:26
		_go_fuzz_dep_.CoverTab[50602]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:26
		return methods.CheckInitialized != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:26
		// _ = "end of CoverTab[50602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:26
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:26
		_go_fuzz_dep_.CoverTab[50603]++
													_, err := methods.CheckInitialized(protoiface.CheckInitializedInput{
			Message: m,
		})
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:30
		// _ = "end of CoverTab[50603]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:31
		_go_fuzz_dep_.CoverTab[50604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:31
		// _ = "end of CoverTab[50604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:31
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:31
	// _ = "end of CoverTab[50600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:31
	_go_fuzz_dep_.CoverTab[50601]++
												return checkInitializedSlow(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:32
	// _ = "end of CoverTab[50601]"
}

func checkInitializedSlow(m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:35
	_go_fuzz_dep_.CoverTab[50605]++
												md := m.Descriptor()
												fds := md.Fields()
												for i, nums := 0, md.RequiredNumbers(); i < nums.Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:38
		_go_fuzz_dep_.CoverTab[50608]++
													fd := fds.ByNumber(nums.Get(i))
													if !m.Has(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:40
			_go_fuzz_dep_.CoverTab[50609]++
														return errors.RequiredNotSet(string(fd.FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:41
			// _ = "end of CoverTab[50609]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:42
			_go_fuzz_dep_.CoverTab[50610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:42
			// _ = "end of CoverTab[50610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:42
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:42
		// _ = "end of CoverTab[50608]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:43
	// _ = "end of CoverTab[50605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:43
	_go_fuzz_dep_.CoverTab[50606]++
												var err error
												m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:45
		_go_fuzz_dep_.CoverTab[50611]++
													switch {
		case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:47
			_go_fuzz_dep_.CoverTab[50613]++
														if fd.Message() == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:48
				_go_fuzz_dep_.CoverTab[50619]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:49
				// _ = "end of CoverTab[50619]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:50
				_go_fuzz_dep_.CoverTab[50620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:50
				// _ = "end of CoverTab[50620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:50
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:50
			// _ = "end of CoverTab[50613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:50
			_go_fuzz_dep_.CoverTab[50614]++
														for i, list := 0, v.List(); i < list.Len() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:51
				_go_fuzz_dep_.CoverTab[50621]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:51
				return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:51
				// _ = "end of CoverTab[50621]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:51
			}(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:51
				_go_fuzz_dep_.CoverTab[50622]++
															err = checkInitialized(list.Get(i).Message())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:52
				// _ = "end of CoverTab[50622]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:53
			// _ = "end of CoverTab[50614]"
		case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:54
			_go_fuzz_dep_.CoverTab[50615]++
														if fd.MapValue().Message() == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:55
				_go_fuzz_dep_.CoverTab[50623]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:56
				// _ = "end of CoverTab[50623]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:57
				_go_fuzz_dep_.CoverTab[50624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:57
				// _ = "end of CoverTab[50624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:57
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:57
			// _ = "end of CoverTab[50615]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:57
			_go_fuzz_dep_.CoverTab[50616]++
														v.Map().Range(func(key protoreflect.MapKey, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:58
				_go_fuzz_dep_.CoverTab[50625]++
															err = checkInitialized(v.Message())
															return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:60
				// _ = "end of CoverTab[50625]"
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:61
			// _ = "end of CoverTab[50616]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:62
			_go_fuzz_dep_.CoverTab[50617]++
														if fd.Message() == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:63
				_go_fuzz_dep_.CoverTab[50626]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:64
				// _ = "end of CoverTab[50626]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:65
				_go_fuzz_dep_.CoverTab[50627]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:65
				// _ = "end of CoverTab[50627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:65
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:65
			// _ = "end of CoverTab[50617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:65
			_go_fuzz_dep_.CoverTab[50618]++
														err = checkInitialized(v.Message())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:66
			// _ = "end of CoverTab[50618]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:67
		// _ = "end of CoverTab[50611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:67
		_go_fuzz_dep_.CoverTab[50612]++
													return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:68
		// _ = "end of CoverTab[50612]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:69
	// _ = "end of CoverTab[50606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:69
	_go_fuzz_dep_.CoverTab[50607]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:70
	// _ = "end of CoverTab[50607]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:71
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/checkinit.go:71
var _ = _go_fuzz_dep_.CoverTab
