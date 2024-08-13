// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
package protodesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:5
)

import (
	"strings"
	"unicode"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/types/descriptorpb"
)

func validateEnumDeclarations(es []filedesc.Enum, eds []*descriptorpb.EnumDescriptorProto) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:22
	_go_fuzz_dep_.CoverTab[60715]++
														for i, ed := range eds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:23
		_go_fuzz_dep_.CoverTab[60717]++
															e := &es[i]
															if err := e.L2.ReservedNames.CheckValid(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:25
			_go_fuzz_dep_.CoverTab[60724]++
																return errors.New("enum %q reserved names has %v", e.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:26
			// _ = "end of CoverTab[60724]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:27
			_go_fuzz_dep_.CoverTab[60725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:27
			// _ = "end of CoverTab[60725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:27
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:27
		// _ = "end of CoverTab[60717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:27
		_go_fuzz_dep_.CoverTab[60718]++
															if err := e.L2.ReservedRanges.CheckValid(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:28
			_go_fuzz_dep_.CoverTab[60726]++
																return errors.New("enum %q reserved ranges has %v", e.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:29
			// _ = "end of CoverTab[60726]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:30
			_go_fuzz_dep_.CoverTab[60727]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:30
			// _ = "end of CoverTab[60727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:30
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:30
		// _ = "end of CoverTab[60718]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:30
		_go_fuzz_dep_.CoverTab[60719]++
															if len(ed.GetValue()) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:31
			_go_fuzz_dep_.CoverTab[60728]++
																return errors.New("enum %q must contain at least one value declaration", e.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:32
			// _ = "end of CoverTab[60728]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:33
			_go_fuzz_dep_.CoverTab[60729]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:33
			// _ = "end of CoverTab[60729]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:33
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:33
		// _ = "end of CoverTab[60719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:33
		_go_fuzz_dep_.CoverTab[60720]++
															allowAlias := ed.GetOptions().GetAllowAlias()
															foundAlias := false
															for i := 0; i < e.Values().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:36
			_go_fuzz_dep_.CoverTab[60730]++
																v1 := e.Values().Get(i)
																if v2 := e.Values().ByNumber(v1.Number()); v1 != v2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:38
				_go_fuzz_dep_.CoverTab[60731]++
																	foundAlias = true
																	if !allowAlias {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:40
					_go_fuzz_dep_.CoverTab[60732]++
																		return errors.New("enum %q has conflicting non-aliased values on number %d: %q with %q", e.FullName(), v1.Number(), v1.Name(), v2.Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:41
					// _ = "end of CoverTab[60732]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:42
					_go_fuzz_dep_.CoverTab[60733]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:42
					// _ = "end of CoverTab[60733]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:42
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:42
				// _ = "end of CoverTab[60731]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:43
				_go_fuzz_dep_.CoverTab[60734]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:43
				// _ = "end of CoverTab[60734]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:43
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:43
			// _ = "end of CoverTab[60730]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:44
		// _ = "end of CoverTab[60720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:44
		_go_fuzz_dep_.CoverTab[60721]++
															if allowAlias && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:45
			_go_fuzz_dep_.CoverTab[60735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:45
			return !foundAlias
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:45
			// _ = "end of CoverTab[60735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:45
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:45
			_go_fuzz_dep_.CoverTab[60736]++
																return errors.New("enum %q allows aliases, but none were found", e.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:46
			// _ = "end of CoverTab[60736]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:47
			_go_fuzz_dep_.CoverTab[60737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:47
			// _ = "end of CoverTab[60737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:47
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:47
		// _ = "end of CoverTab[60721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:47
		_go_fuzz_dep_.CoverTab[60722]++
															if e.Syntax() == protoreflect.Proto3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:48
			_go_fuzz_dep_.CoverTab[60738]++
																if v := e.Values().Get(0); v.Number() != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:49
				_go_fuzz_dep_.CoverTab[60740]++
																	return errors.New("enum %q using proto3 semantics must have zero number for the first value", v.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:50
				// _ = "end of CoverTab[60740]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:51
				_go_fuzz_dep_.CoverTab[60741]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:51
				// _ = "end of CoverTab[60741]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:51
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:51
			// _ = "end of CoverTab[60738]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:51
			_go_fuzz_dep_.CoverTab[60739]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:55
			names := map[string]protoreflect.EnumValueDescriptor{}
			prefix := strings.Replace(strings.ToLower(string(e.Name())), "_", "", -1)
			for i := 0; i < e.Values().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:57
				_go_fuzz_dep_.CoverTab[60742]++
																	v1 := e.Values().Get(i)
																	s := strs.EnumValueName(strs.TrimEnumPrefix(string(v1.Name()), prefix))
																	if v2, ok := names[s]; ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:60
					_go_fuzz_dep_.CoverTab[60744]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:60
					return v1.Number() != v2.Number()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:60
					// _ = "end of CoverTab[60744]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:60
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:60
					_go_fuzz_dep_.CoverTab[60745]++
																		return errors.New("enum %q using proto3 semantics has conflict: %q with %q", e.FullName(), v1.Name(), v2.Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:61
					// _ = "end of CoverTab[60745]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:62
					_go_fuzz_dep_.CoverTab[60746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:62
					// _ = "end of CoverTab[60746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:62
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:62
				// _ = "end of CoverTab[60742]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:62
				_go_fuzz_dep_.CoverTab[60743]++
																	names[s] = v1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:63
				// _ = "end of CoverTab[60743]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:64
			// _ = "end of CoverTab[60739]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:65
			_go_fuzz_dep_.CoverTab[60747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:65
			// _ = "end of CoverTab[60747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:65
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:65
		// _ = "end of CoverTab[60722]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:65
		_go_fuzz_dep_.CoverTab[60723]++

															for j, vd := range ed.GetValue() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:67
			_go_fuzz_dep_.CoverTab[60748]++
																v := &e.L2.Values.List[j]
																if vd.Number == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:69
				_go_fuzz_dep_.CoverTab[60751]++
																	return errors.New("enum value %q must have a specified number", v.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:70
				// _ = "end of CoverTab[60751]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:71
				_go_fuzz_dep_.CoverTab[60752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:71
				// _ = "end of CoverTab[60752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:71
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:71
			// _ = "end of CoverTab[60748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:71
			_go_fuzz_dep_.CoverTab[60749]++
																if e.L2.ReservedNames.Has(v.Name()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:72
				_go_fuzz_dep_.CoverTab[60753]++
																	return errors.New("enum value %q must not use reserved name", v.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:73
				// _ = "end of CoverTab[60753]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:74
				_go_fuzz_dep_.CoverTab[60754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:74
				// _ = "end of CoverTab[60754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:74
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:74
			// _ = "end of CoverTab[60749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:74
			_go_fuzz_dep_.CoverTab[60750]++
																if e.L2.ReservedRanges.Has(v.Number()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:75
				_go_fuzz_dep_.CoverTab[60755]++
																	return errors.New("enum value %q must not use reserved number %d", v.FullName(), v.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:76
				// _ = "end of CoverTab[60755]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:77
				_go_fuzz_dep_.CoverTab[60756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:77
				// _ = "end of CoverTab[60756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:77
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:77
			// _ = "end of CoverTab[60750]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:78
		// _ = "end of CoverTab[60723]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:79
	// _ = "end of CoverTab[60715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:79
	_go_fuzz_dep_.CoverTab[60716]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:80
	// _ = "end of CoverTab[60716]"
}

func validateMessageDeclarations(ms []filedesc.Message, mds []*descriptorpb.DescriptorProto) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:83
	_go_fuzz_dep_.CoverTab[60757]++
														for i, md := range mds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:84
		_go_fuzz_dep_.CoverTab[60759]++
															m := &ms[i]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:88
		isMessageSet := md.GetOptions().GetMessageSetWireFormat()
		if err := m.L2.ReservedNames.CheckValid(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:89
			_go_fuzz_dep_.CoverTab[60772]++
																return errors.New("message %q reserved names has %v", m.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:90
			// _ = "end of CoverTab[60772]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:91
			_go_fuzz_dep_.CoverTab[60773]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:91
			// _ = "end of CoverTab[60773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:91
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:91
		// _ = "end of CoverTab[60759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:91
		_go_fuzz_dep_.CoverTab[60760]++
															if err := m.L2.ReservedRanges.CheckValid(isMessageSet); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:92
			_go_fuzz_dep_.CoverTab[60774]++
																return errors.New("message %q reserved ranges has %v", m.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:93
			// _ = "end of CoverTab[60774]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:94
			_go_fuzz_dep_.CoverTab[60775]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:94
			// _ = "end of CoverTab[60775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:94
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:94
		// _ = "end of CoverTab[60760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:94
		_go_fuzz_dep_.CoverTab[60761]++
															if err := m.L2.ExtensionRanges.CheckValid(isMessageSet); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:95
			_go_fuzz_dep_.CoverTab[60776]++
																return errors.New("message %q extension ranges has %v", m.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:96
			// _ = "end of CoverTab[60776]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:97
			_go_fuzz_dep_.CoverTab[60777]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:97
			// _ = "end of CoverTab[60777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:97
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:97
		// _ = "end of CoverTab[60761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:97
		_go_fuzz_dep_.CoverTab[60762]++
															if err := (*filedesc.FieldRanges).CheckOverlap(&m.L2.ReservedRanges, &m.L2.ExtensionRanges); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:98
			_go_fuzz_dep_.CoverTab[60778]++
																return errors.New("message %q reserved and extension ranges has %v", m.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:99
			// _ = "end of CoverTab[60778]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:100
			_go_fuzz_dep_.CoverTab[60779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:100
			// _ = "end of CoverTab[60779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:100
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:100
		// _ = "end of CoverTab[60762]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:100
		_go_fuzz_dep_.CoverTab[60763]++
															for i := 0; i < m.Fields().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:101
			_go_fuzz_dep_.CoverTab[60780]++
																f1 := m.Fields().Get(i)
																if f2 := m.Fields().ByNumber(f1.Number()); f1 != f2 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:103
				_go_fuzz_dep_.CoverTab[60781]++
																	return errors.New("message %q has conflicting fields: %q with %q", m.FullName(), f1.Name(), f2.Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:104
				// _ = "end of CoverTab[60781]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:105
				_go_fuzz_dep_.CoverTab[60782]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:105
				// _ = "end of CoverTab[60782]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:105
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:105
			// _ = "end of CoverTab[60780]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:106
		// _ = "end of CoverTab[60763]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:106
		_go_fuzz_dep_.CoverTab[60764]++
															if isMessageSet && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:107
			_go_fuzz_dep_.CoverTab[60783]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:107
			return !flags.ProtoLegacy
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:107
			// _ = "end of CoverTab[60783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:107
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:107
			_go_fuzz_dep_.CoverTab[60784]++
																return errors.New("message %q is a MessageSet, which is a legacy proto1 feature that is no longer supported", m.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:108
			// _ = "end of CoverTab[60784]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:109
			_go_fuzz_dep_.CoverTab[60785]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:109
			// _ = "end of CoverTab[60785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:109
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:109
		// _ = "end of CoverTab[60764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:109
		_go_fuzz_dep_.CoverTab[60765]++
															if isMessageSet && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
			_go_fuzz_dep_.CoverTab[60786]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
			return (m.Syntax() != protoreflect.Proto2 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
				_go_fuzz_dep_.CoverTab[60787]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
				return m.Fields().Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
				// _ = "end of CoverTab[60787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
				_go_fuzz_dep_.CoverTab[60788]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
				return m.ExtensionRanges().Len() == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
				// _ = "end of CoverTab[60788]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
			// _ = "end of CoverTab[60786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:110
			_go_fuzz_dep_.CoverTab[60789]++
																return errors.New("message %q is an invalid proto1 MessageSet", m.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:111
			// _ = "end of CoverTab[60789]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:112
			_go_fuzz_dep_.CoverTab[60790]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:112
			// _ = "end of CoverTab[60790]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:112
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:112
		// _ = "end of CoverTab[60765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:112
		_go_fuzz_dep_.CoverTab[60766]++
															if m.Syntax() == protoreflect.Proto3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:113
			_go_fuzz_dep_.CoverTab[60791]++
																if m.ExtensionRanges().Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:114
				_go_fuzz_dep_.CoverTab[60793]++
																	return errors.New("message %q using proto3 semantics cannot have extension ranges", m.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:115
				// _ = "end of CoverTab[60793]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:116
				_go_fuzz_dep_.CoverTab[60794]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:116
				// _ = "end of CoverTab[60794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:116
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:116
			// _ = "end of CoverTab[60791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:116
			_go_fuzz_dep_.CoverTab[60792]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:120
			names := map[string]protoreflect.FieldDescriptor{}
			for i := 0; i < m.Fields().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:121
				_go_fuzz_dep_.CoverTab[60795]++
																	f1 := m.Fields().Get(i)
																	s := strings.Replace(strings.ToLower(string(f1.Name())), "_", "", -1)
																	if f2, ok := names[s]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:124
					_go_fuzz_dep_.CoverTab[60797]++
																		return errors.New("message %q using proto3 semantics has conflict: %q with %q", m.FullName(), f1.Name(), f2.Name())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:125
					// _ = "end of CoverTab[60797]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:126
					_go_fuzz_dep_.CoverTab[60798]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:126
					// _ = "end of CoverTab[60798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:126
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:126
				// _ = "end of CoverTab[60795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:126
				_go_fuzz_dep_.CoverTab[60796]++
																	names[s] = f1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:127
				// _ = "end of CoverTab[60796]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:128
			// _ = "end of CoverTab[60792]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:129
			_go_fuzz_dep_.CoverTab[60799]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:129
			// _ = "end of CoverTab[60799]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:129
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:129
		// _ = "end of CoverTab[60766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:129
		_go_fuzz_dep_.CoverTab[60767]++

															for j, fd := range md.GetField() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:131
			_go_fuzz_dep_.CoverTab[60800]++
																f := &m.L2.Fields.List[j]
																if m.L2.ReservedNames.Has(f.Name()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:133
				_go_fuzz_dep_.CoverTab[60813]++
																	return errors.New("message field %q must not use reserved name", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:134
				// _ = "end of CoverTab[60813]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:135
				_go_fuzz_dep_.CoverTab[60814]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:135
				// _ = "end of CoverTab[60814]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:135
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:135
			// _ = "end of CoverTab[60800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:135
			_go_fuzz_dep_.CoverTab[60801]++
																if !f.Number().IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:136
				_go_fuzz_dep_.CoverTab[60815]++
																	return errors.New("message field %q has an invalid number: %d", f.FullName(), f.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:137
				// _ = "end of CoverTab[60815]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:138
				_go_fuzz_dep_.CoverTab[60816]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:138
				// _ = "end of CoverTab[60816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:138
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:138
			// _ = "end of CoverTab[60801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:138
			_go_fuzz_dep_.CoverTab[60802]++
																if !f.Cardinality().IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:139
				_go_fuzz_dep_.CoverTab[60817]++
																	return errors.New("message field %q has an invalid cardinality: %d", f.FullName(), f.Cardinality())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:140
				// _ = "end of CoverTab[60817]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:141
				_go_fuzz_dep_.CoverTab[60818]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:141
				// _ = "end of CoverTab[60818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:141
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:141
			// _ = "end of CoverTab[60802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:141
			_go_fuzz_dep_.CoverTab[60803]++
																if m.L2.ReservedRanges.Has(f.Number()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:142
				_go_fuzz_dep_.CoverTab[60819]++
																	return errors.New("message field %q must not use reserved number %d", f.FullName(), f.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:143
				// _ = "end of CoverTab[60819]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:144
				_go_fuzz_dep_.CoverTab[60820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:144
				// _ = "end of CoverTab[60820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:144
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:144
			// _ = "end of CoverTab[60803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:144
			_go_fuzz_dep_.CoverTab[60804]++
																if m.L2.ExtensionRanges.Has(f.Number()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:145
				_go_fuzz_dep_.CoverTab[60821]++
																	return errors.New("message field %q with number %d in extension range", f.FullName(), f.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:146
				// _ = "end of CoverTab[60821]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:147
				_go_fuzz_dep_.CoverTab[60822]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:147
				// _ = "end of CoverTab[60822]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:147
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:147
			// _ = "end of CoverTab[60804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:147
			_go_fuzz_dep_.CoverTab[60805]++
																if fd.Extendee != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:148
				_go_fuzz_dep_.CoverTab[60823]++
																	return errors.New("message field %q may not have extendee: %q", f.FullName(), fd.GetExtendee())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:149
				// _ = "end of CoverTab[60823]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:150
				_go_fuzz_dep_.CoverTab[60824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:150
				// _ = "end of CoverTab[60824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:150
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:150
			// _ = "end of CoverTab[60805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:150
			_go_fuzz_dep_.CoverTab[60806]++
																if f.L1.IsProto3Optional {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:151
				_go_fuzz_dep_.CoverTab[60825]++
																	if f.Syntax() != protoreflect.Proto3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:152
					_go_fuzz_dep_.CoverTab[60828]++
																		return errors.New("message field %q under proto3 optional semantics must be specified in the proto3 syntax", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:153
					// _ = "end of CoverTab[60828]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:154
					_go_fuzz_dep_.CoverTab[60829]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:154
					// _ = "end of CoverTab[60829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:154
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:154
				// _ = "end of CoverTab[60825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:154
				_go_fuzz_dep_.CoverTab[60826]++
																	if f.Cardinality() != protoreflect.Optional {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:155
					_go_fuzz_dep_.CoverTab[60830]++
																		return errors.New("message field %q under proto3 optional semantics must have optional cardinality", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:156
					// _ = "end of CoverTab[60830]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:157
					_go_fuzz_dep_.CoverTab[60831]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:157
					// _ = "end of CoverTab[60831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:157
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:157
				// _ = "end of CoverTab[60826]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:157
				_go_fuzz_dep_.CoverTab[60827]++
																	if f.ContainingOneof() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:158
					_go_fuzz_dep_.CoverTab[60832]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:158
					return f.ContainingOneof().Fields().Len() != 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:158
					// _ = "end of CoverTab[60832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:158
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:158
					_go_fuzz_dep_.CoverTab[60833]++
																		return errors.New("message field %q under proto3 optional semantics must be within a single element oneof", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:159
					// _ = "end of CoverTab[60833]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:160
					_go_fuzz_dep_.CoverTab[60834]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:160
					// _ = "end of CoverTab[60834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:160
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:160
				// _ = "end of CoverTab[60827]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:161
				_go_fuzz_dep_.CoverTab[60835]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:161
				// _ = "end of CoverTab[60835]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:161
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:161
			// _ = "end of CoverTab[60806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:161
			_go_fuzz_dep_.CoverTab[60807]++
																if f.IsWeak() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:162
				_go_fuzz_dep_.CoverTab[60836]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:162
				return !flags.ProtoLegacy
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:162
				// _ = "end of CoverTab[60836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:162
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:162
				_go_fuzz_dep_.CoverTab[60837]++
																	return errors.New("message field %q is a weak field, which is a legacy proto1 feature that is no longer supported", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:163
				// _ = "end of CoverTab[60837]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:164
				_go_fuzz_dep_.CoverTab[60838]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:164
				// _ = "end of CoverTab[60838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:164
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:164
			// _ = "end of CoverTab[60807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:164
			_go_fuzz_dep_.CoverTab[60808]++
																if f.IsWeak() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
				_go_fuzz_dep_.CoverTab[60839]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
				return (f.Syntax() != protoreflect.Proto2 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
					_go_fuzz_dep_.CoverTab[60840]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
					return !isOptionalMessage(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
					// _ = "end of CoverTab[60840]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
					_go_fuzz_dep_.CoverTab[60841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
					return f.ContainingOneof() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
					// _ = "end of CoverTab[60841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
				}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
				// _ = "end of CoverTab[60839]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:165
				_go_fuzz_dep_.CoverTab[60842]++
																	return errors.New("message field %q may only be weak for an optional message", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:166
				// _ = "end of CoverTab[60842]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:167
				_go_fuzz_dep_.CoverTab[60843]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:167
				// _ = "end of CoverTab[60843]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:167
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:167
			// _ = "end of CoverTab[60808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:167
			_go_fuzz_dep_.CoverTab[60809]++
																if f.IsPacked() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:168
				_go_fuzz_dep_.CoverTab[60844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:168
				return !isPackable(f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:168
				// _ = "end of CoverTab[60844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:168
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:168
				_go_fuzz_dep_.CoverTab[60845]++
																	return errors.New("message field %q is not packable", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:169
				// _ = "end of CoverTab[60845]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:170
				_go_fuzz_dep_.CoverTab[60846]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:170
				// _ = "end of CoverTab[60846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:170
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:170
			// _ = "end of CoverTab[60809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:170
			_go_fuzz_dep_.CoverTab[60810]++
																if err := checkValidGroup(f); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:171
				_go_fuzz_dep_.CoverTab[60847]++
																	return errors.New("message field %q is an invalid group: %v", f.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:172
				// _ = "end of CoverTab[60847]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:173
				_go_fuzz_dep_.CoverTab[60848]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:173
				// _ = "end of CoverTab[60848]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:173
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:173
			// _ = "end of CoverTab[60810]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:173
			_go_fuzz_dep_.CoverTab[60811]++
																if err := checkValidMap(f); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:174
				_go_fuzz_dep_.CoverTab[60849]++
																	return errors.New("message field %q is an invalid map: %v", f.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:175
				// _ = "end of CoverTab[60849]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:176
				_go_fuzz_dep_.CoverTab[60850]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:176
				// _ = "end of CoverTab[60850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:176
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:176
			// _ = "end of CoverTab[60811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:176
			_go_fuzz_dep_.CoverTab[60812]++
																if f.Syntax() == protoreflect.Proto3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:177
				_go_fuzz_dep_.CoverTab[60851]++
																	if f.Cardinality() == protoreflect.Required {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:178
					_go_fuzz_dep_.CoverTab[60853]++
																		return errors.New("message field %q using proto3 semantics cannot be required", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:179
					// _ = "end of CoverTab[60853]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:180
					_go_fuzz_dep_.CoverTab[60854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:180
					// _ = "end of CoverTab[60854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:180
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:180
				// _ = "end of CoverTab[60851]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:180
				_go_fuzz_dep_.CoverTab[60852]++
																	if f.Enum() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					_go_fuzz_dep_.CoverTab[60855]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					return !f.Enum().IsPlaceholder()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					// _ = "end of CoverTab[60855]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					_go_fuzz_dep_.CoverTab[60856]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					return f.Enum().Syntax() != protoreflect.Proto3
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					// _ = "end of CoverTab[60856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:181
					_go_fuzz_dep_.CoverTab[60857]++
																		return errors.New("message field %q using proto3 semantics may only depend on a proto3 enum", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:182
					// _ = "end of CoverTab[60857]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:183
					_go_fuzz_dep_.CoverTab[60858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:183
					// _ = "end of CoverTab[60858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:183
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:183
				// _ = "end of CoverTab[60852]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:184
				_go_fuzz_dep_.CoverTab[60859]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:184
				// _ = "end of CoverTab[60859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:184
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:184
			// _ = "end of CoverTab[60812]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:185
		// _ = "end of CoverTab[60767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:185
		_go_fuzz_dep_.CoverTab[60768]++
															seenSynthetic := false
															for j := range md.GetOneofDecl() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:187
			_go_fuzz_dep_.CoverTab[60860]++
																o := &m.L2.Oneofs.List[j]
																if o.Fields().Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:189
				_go_fuzz_dep_.CoverTab[60865]++
																	return errors.New("message oneof %q must contain at least one field declaration", o.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:190
				// _ = "end of CoverTab[60865]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:191
				_go_fuzz_dep_.CoverTab[60866]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:191
				// _ = "end of CoverTab[60866]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:191
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:191
			// _ = "end of CoverTab[60860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:191
			_go_fuzz_dep_.CoverTab[60861]++
																if n := o.Fields().Len(); n-1 != (o.Fields().Get(n-1).Index() - o.Fields().Get(0).Index()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:192
				_go_fuzz_dep_.CoverTab[60867]++
																	return errors.New("message oneof %q must have consecutively declared fields", o.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:193
				// _ = "end of CoverTab[60867]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:194
				_go_fuzz_dep_.CoverTab[60868]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:194
				// _ = "end of CoverTab[60868]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:194
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:194
			// _ = "end of CoverTab[60861]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:194
			_go_fuzz_dep_.CoverTab[60862]++

																if o.IsSynthetic() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:196
				_go_fuzz_dep_.CoverTab[60869]++
																	seenSynthetic = true
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:198
				// _ = "end of CoverTab[60869]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:199
				_go_fuzz_dep_.CoverTab[60870]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:199
				// _ = "end of CoverTab[60870]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:199
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:199
			// _ = "end of CoverTab[60862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:199
			_go_fuzz_dep_.CoverTab[60863]++
																if !o.IsSynthetic() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:200
				_go_fuzz_dep_.CoverTab[60871]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:200
				return seenSynthetic
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:200
				// _ = "end of CoverTab[60871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:200
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:200
				_go_fuzz_dep_.CoverTab[60872]++
																	return errors.New("message oneof %q must be declared before synthetic oneofs", o.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:201
				// _ = "end of CoverTab[60872]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:202
				_go_fuzz_dep_.CoverTab[60873]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:202
				// _ = "end of CoverTab[60873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:202
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:202
			// _ = "end of CoverTab[60863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:202
			_go_fuzz_dep_.CoverTab[60864]++

																for i := 0; i < o.Fields().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:204
				_go_fuzz_dep_.CoverTab[60874]++
																	f := o.Fields().Get(i)
																	if f.Cardinality() != protoreflect.Optional {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:206
					_go_fuzz_dep_.CoverTab[60876]++
																		return errors.New("message field %q belongs in a oneof and must be optional", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:207
					// _ = "end of CoverTab[60876]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:208
					_go_fuzz_dep_.CoverTab[60877]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:208
					// _ = "end of CoverTab[60877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:208
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:208
				// _ = "end of CoverTab[60874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:208
				_go_fuzz_dep_.CoverTab[60875]++
																	if f.IsWeak() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:209
					_go_fuzz_dep_.CoverTab[60878]++
																		return errors.New("message field %q belongs in a oneof and must not be a weak reference", f.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:210
					// _ = "end of CoverTab[60878]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:211
					_go_fuzz_dep_.CoverTab[60879]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:211
					// _ = "end of CoverTab[60879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:211
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:211
				// _ = "end of CoverTab[60875]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:212
			// _ = "end of CoverTab[60864]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:213
		// _ = "end of CoverTab[60768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:213
		_go_fuzz_dep_.CoverTab[60769]++

															if err := validateEnumDeclarations(m.L1.Enums.List, md.GetEnumType()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:215
			_go_fuzz_dep_.CoverTab[60880]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:216
			// _ = "end of CoverTab[60880]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:217
			_go_fuzz_dep_.CoverTab[60881]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:217
			// _ = "end of CoverTab[60881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:217
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:217
		// _ = "end of CoverTab[60769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:217
		_go_fuzz_dep_.CoverTab[60770]++
															if err := validateMessageDeclarations(m.L1.Messages.List, md.GetNestedType()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:218
			_go_fuzz_dep_.CoverTab[60882]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:219
			// _ = "end of CoverTab[60882]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:220
			_go_fuzz_dep_.CoverTab[60883]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:220
			// _ = "end of CoverTab[60883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:220
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:220
		// _ = "end of CoverTab[60770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:220
		_go_fuzz_dep_.CoverTab[60771]++
															if err := validateExtensionDeclarations(m.L1.Extensions.List, md.GetExtension()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:221
			_go_fuzz_dep_.CoverTab[60884]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:222
			// _ = "end of CoverTab[60884]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:223
			_go_fuzz_dep_.CoverTab[60885]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:223
			// _ = "end of CoverTab[60885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:223
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:223
		// _ = "end of CoverTab[60771]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:224
	// _ = "end of CoverTab[60757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:224
	_go_fuzz_dep_.CoverTab[60758]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:225
	// _ = "end of CoverTab[60758]"
}

func validateExtensionDeclarations(xs []filedesc.Extension, xds []*descriptorpb.FieldDescriptorProto) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:228
	_go_fuzz_dep_.CoverTab[60886]++
														for i, xd := range xds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:229
		_go_fuzz_dep_.CoverTab[60888]++
															x := &xs[i]

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
		if n := x.Number(); n < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
			_go_fuzz_dep_.CoverTab[60898]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
			return (protowire.FirstReservedNumber <= n && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
				_go_fuzz_dep_.CoverTab[60899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
				return n <= protowire.LastReservedNumber
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
				// _ = "end of CoverTab[60899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
			// _ = "end of CoverTab[60898]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:235
			_go_fuzz_dep_.CoverTab[60900]++
																return errors.New("extension field %q has an invalid number: %d", x.FullName(), x.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:236
			// _ = "end of CoverTab[60900]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:237
			_go_fuzz_dep_.CoverTab[60901]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:237
			// _ = "end of CoverTab[60901]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:237
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:237
		// _ = "end of CoverTab[60888]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:237
		_go_fuzz_dep_.CoverTab[60889]++
															if !x.Cardinality().IsValid() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:238
			_go_fuzz_dep_.CoverTab[60902]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:238
			return x.Cardinality() == protoreflect.Required
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:238
			// _ = "end of CoverTab[60902]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:238
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:238
			_go_fuzz_dep_.CoverTab[60903]++
																return errors.New("extension field %q has an invalid cardinality: %d", x.FullName(), x.Cardinality())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:239
			// _ = "end of CoverTab[60903]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:240
			_go_fuzz_dep_.CoverTab[60904]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:240
			// _ = "end of CoverTab[60904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:240
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:240
		// _ = "end of CoverTab[60889]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:240
		_go_fuzz_dep_.CoverTab[60890]++
															if xd.JsonName != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:241
			_go_fuzz_dep_.CoverTab[60905]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:245
			if xd.GetJsonName() != strs.JSONCamelCase(string(x.Name())) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:245
				_go_fuzz_dep_.CoverTab[60906]++
																	return errors.New("extension field %q may not have an explicitly set JSON name: %q", x.FullName(), xd.GetJsonName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:246
				// _ = "end of CoverTab[60906]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:247
				_go_fuzz_dep_.CoverTab[60907]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:247
				// _ = "end of CoverTab[60907]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:247
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:247
			// _ = "end of CoverTab[60905]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:248
			_go_fuzz_dep_.CoverTab[60908]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:248
			// _ = "end of CoverTab[60908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:248
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:248
		// _ = "end of CoverTab[60890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:248
		_go_fuzz_dep_.CoverTab[60891]++
															if xd.OneofIndex != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:249
			_go_fuzz_dep_.CoverTab[60909]++
																return errors.New("extension field %q may not be part of a oneof", x.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:250
			// _ = "end of CoverTab[60909]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:251
			_go_fuzz_dep_.CoverTab[60910]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:251
			// _ = "end of CoverTab[60910]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:251
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:251
		// _ = "end of CoverTab[60891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:251
		_go_fuzz_dep_.CoverTab[60892]++
															if md := x.ContainingMessage(); !md.IsPlaceholder() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:252
			_go_fuzz_dep_.CoverTab[60911]++
																if !md.ExtensionRanges().Has(x.Number()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:253
				_go_fuzz_dep_.CoverTab[60914]++
																	return errors.New("extension field %q extends %q with non-extension field number: %d", x.FullName(), md.FullName(), x.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:254
				// _ = "end of CoverTab[60914]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:255
				_go_fuzz_dep_.CoverTab[60915]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:255
				// _ = "end of CoverTab[60915]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:255
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:255
			// _ = "end of CoverTab[60911]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:255
			_go_fuzz_dep_.CoverTab[60912]++
																isMessageSet := md.Options().(*descriptorpb.MessageOptions).GetMessageSetWireFormat()
																if isMessageSet && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:257
				_go_fuzz_dep_.CoverTab[60916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:257
				return !isOptionalMessage(x)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:257
				// _ = "end of CoverTab[60916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:257
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:257
				_go_fuzz_dep_.CoverTab[60917]++
																	return errors.New("extension field %q extends MessageSet and must be an optional message", x.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:258
				// _ = "end of CoverTab[60917]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:259
				_go_fuzz_dep_.CoverTab[60918]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:259
				// _ = "end of CoverTab[60918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:259
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:259
			// _ = "end of CoverTab[60912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:259
			_go_fuzz_dep_.CoverTab[60913]++
																if !isMessageSet && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:260
				_go_fuzz_dep_.CoverTab[60919]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:260
				return !x.Number().IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:260
				// _ = "end of CoverTab[60919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:260
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:260
				_go_fuzz_dep_.CoverTab[60920]++
																	return errors.New("extension field %q has an invalid number: %d", x.FullName(), x.Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:261
				// _ = "end of CoverTab[60920]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:262
				_go_fuzz_dep_.CoverTab[60921]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:262
				// _ = "end of CoverTab[60921]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:262
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:262
			// _ = "end of CoverTab[60913]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:263
			_go_fuzz_dep_.CoverTab[60922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:263
			// _ = "end of CoverTab[60922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:263
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:263
		// _ = "end of CoverTab[60892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:263
		_go_fuzz_dep_.CoverTab[60893]++
															if xd.GetOptions().GetWeak() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:264
			_go_fuzz_dep_.CoverTab[60923]++
																return errors.New("extension field %q cannot be a weak reference", x.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:265
			// _ = "end of CoverTab[60923]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:266
			_go_fuzz_dep_.CoverTab[60924]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:266
			// _ = "end of CoverTab[60924]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:266
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:266
		// _ = "end of CoverTab[60893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:266
		_go_fuzz_dep_.CoverTab[60894]++
															if x.IsPacked() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:267
			_go_fuzz_dep_.CoverTab[60925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:267
			return !isPackable(x)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:267
			// _ = "end of CoverTab[60925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:267
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:267
			_go_fuzz_dep_.CoverTab[60926]++
																return errors.New("extension field %q is not packable", x.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:268
			// _ = "end of CoverTab[60926]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:269
			_go_fuzz_dep_.CoverTab[60927]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:269
			// _ = "end of CoverTab[60927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:269
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:269
		// _ = "end of CoverTab[60894]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:269
		_go_fuzz_dep_.CoverTab[60895]++
															if err := checkValidGroup(x); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:270
			_go_fuzz_dep_.CoverTab[60928]++
																return errors.New("extension field %q is an invalid group: %v", x.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:271
			// _ = "end of CoverTab[60928]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:272
			_go_fuzz_dep_.CoverTab[60929]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:272
			// _ = "end of CoverTab[60929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:272
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:272
		// _ = "end of CoverTab[60895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:272
		_go_fuzz_dep_.CoverTab[60896]++
															if md := x.Message(); md != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:273
			_go_fuzz_dep_.CoverTab[60930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:273
			return md.IsMapEntry()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:273
			// _ = "end of CoverTab[60930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:273
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:273
			_go_fuzz_dep_.CoverTab[60931]++
																return errors.New("extension field %q cannot be a map entry", x.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:274
			// _ = "end of CoverTab[60931]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:275
			_go_fuzz_dep_.CoverTab[60932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:275
			// _ = "end of CoverTab[60932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:275
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:275
		// _ = "end of CoverTab[60896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:275
		_go_fuzz_dep_.CoverTab[60897]++
															if x.Syntax() == protoreflect.Proto3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:276
			_go_fuzz_dep_.CoverTab[60933]++
																switch x.ContainingMessage().FullName() {
			case (*descriptorpb.FileOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:278
				_go_fuzz_dep_.CoverTab[60934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:278
				// _ = "end of CoverTab[60934]"
			case (*descriptorpb.EnumOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:279
				_go_fuzz_dep_.CoverTab[60935]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:279
				// _ = "end of CoverTab[60935]"
			case (*descriptorpb.EnumValueOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:280
				_go_fuzz_dep_.CoverTab[60936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:280
				// _ = "end of CoverTab[60936]"
			case (*descriptorpb.MessageOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:281
				_go_fuzz_dep_.CoverTab[60937]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:281
				// _ = "end of CoverTab[60937]"
			case (*descriptorpb.FieldOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:282
				_go_fuzz_dep_.CoverTab[60938]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:282
				// _ = "end of CoverTab[60938]"
			case (*descriptorpb.OneofOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:283
				_go_fuzz_dep_.CoverTab[60939]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:283
				// _ = "end of CoverTab[60939]"
			case (*descriptorpb.ExtensionRangeOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:284
				_go_fuzz_dep_.CoverTab[60940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:284
				// _ = "end of CoverTab[60940]"
			case (*descriptorpb.ServiceOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:285
				_go_fuzz_dep_.CoverTab[60941]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:285
				// _ = "end of CoverTab[60941]"
			case (*descriptorpb.MethodOptions)(nil).ProtoReflect().Descriptor().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:286
				_go_fuzz_dep_.CoverTab[60942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:286
				// _ = "end of CoverTab[60942]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:287
				_go_fuzz_dep_.CoverTab[60943]++
																	return errors.New("extension field %q cannot be declared in proto3 unless extended descriptor options", x.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:288
				// _ = "end of CoverTab[60943]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:289
			// _ = "end of CoverTab[60933]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:290
			_go_fuzz_dep_.CoverTab[60944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:290
			// _ = "end of CoverTab[60944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:290
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:290
		// _ = "end of CoverTab[60897]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:291
	// _ = "end of CoverTab[60886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:291
	_go_fuzz_dep_.CoverTab[60887]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:292
	// _ = "end of CoverTab[60887]"
}

// isOptionalMessage reports whether this is an optional message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:295
// If the kind is unknown, it is assumed to be a message.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:297
func isOptionalMessage(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:297
	_go_fuzz_dep_.CoverTab[60945]++
														return (fd.Kind() == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
		_go_fuzz_dep_.CoverTab[60946]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
		return fd.Kind() == protoreflect.MessageKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
		// _ = "end of CoverTab[60946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
	}()) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
		_go_fuzz_dep_.CoverTab[60947]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
		return fd.Cardinality() == protoreflect.Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
		// _ = "end of CoverTab[60947]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:298
	// _ = "end of CoverTab[60945]"
}

// isPackable checks whether the pack option can be specified.
func isPackable(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:302
	_go_fuzz_dep_.CoverTab[60948]++
														switch fd.Kind() {
	case protoreflect.StringKind, protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:304
		_go_fuzz_dep_.CoverTab[60950]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:305
		// _ = "end of CoverTab[60950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:305
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:305
		_go_fuzz_dep_.CoverTab[60951]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:305
		// _ = "end of CoverTab[60951]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:306
	// _ = "end of CoverTab[60948]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:306
	_go_fuzz_dep_.CoverTab[60949]++
														return fd.IsList()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:307
	// _ = "end of CoverTab[60949]"
}

// checkValidGroup reports whether fd is a valid group according to the same
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:310
// rules that protoc imposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:312
func checkValidGroup(fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:312
	_go_fuzz_dep_.CoverTab[60952]++
														md := fd.Message()
														switch {
	case fd.Kind() != protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:315
		_go_fuzz_dep_.CoverTab[60954]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:316
		// _ = "end of CoverTab[60954]"
	case fd.Syntax() != protoreflect.Proto2:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:317
		_go_fuzz_dep_.CoverTab[60955]++
															return errors.New("invalid under proto2 semantics")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:318
		// _ = "end of CoverTab[60955]"
	case md == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:319
		_go_fuzz_dep_.CoverTab[60961]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:319
		return md.IsPlaceholder()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:319
		// _ = "end of CoverTab[60961]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:319
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:319
		_go_fuzz_dep_.CoverTab[60956]++
															return errors.New("message must be resolvable")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:320
		// _ = "end of CoverTab[60956]"
	case fd.FullName().Parent() != md.FullName().Parent():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:321
		_go_fuzz_dep_.CoverTab[60957]++
															return errors.New("message and field must be declared in the same scope")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:322
		// _ = "end of CoverTab[60957]"
	case !unicode.IsUpper(rune(md.Name()[0])):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:323
		_go_fuzz_dep_.CoverTab[60958]++
															return errors.New("message name must start with an uppercase")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:324
		// _ = "end of CoverTab[60958]"
	case fd.Name() != protoreflect.Name(strings.ToLower(string(md.Name()))):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:325
		_go_fuzz_dep_.CoverTab[60959]++
															return errors.New("field name must be lowercased form of the message name")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:326
		// _ = "end of CoverTab[60959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:326
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:326
		_go_fuzz_dep_.CoverTab[60960]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:326
		// _ = "end of CoverTab[60960]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:327
	// _ = "end of CoverTab[60952]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:327
	_go_fuzz_dep_.CoverTab[60953]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:328
	// _ = "end of CoverTab[60953]"
}

// checkValidMap checks whether the field is a valid map according to the same
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:331
// rules that protoc imposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:331
// See protoc v3.8.0: src/google/protobuf/descriptor.cc:6045-6115
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:334
func checkValidMap(fd protoreflect.FieldDescriptor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:334
	_go_fuzz_dep_.CoverTab[60962]++
														md := fd.Message()
														switch {
	case md == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:337
		_go_fuzz_dep_.CoverTab[60975]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:337
		return !md.IsMapEntry()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:337
		// _ = "end of CoverTab[60975]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:337
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:337
		_go_fuzz_dep_.CoverTab[60967]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:338
		// _ = "end of CoverTab[60967]"
	case fd.FullName().Parent() != md.FullName().Parent():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:339
		_go_fuzz_dep_.CoverTab[60968]++
															return errors.New("message and field must be declared in the same scope")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:340
		// _ = "end of CoverTab[60968]"
	case md.Name() != protoreflect.Name(strs.MapEntryName(string(fd.Name()))):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:341
		_go_fuzz_dep_.CoverTab[60969]++
															return errors.New("incorrect implicit map entry name")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:342
		// _ = "end of CoverTab[60969]"
	case fd.Cardinality() != protoreflect.Repeated:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:343
		_go_fuzz_dep_.CoverTab[60970]++
															return errors.New("field must be repeated")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:344
		// _ = "end of CoverTab[60970]"
	case md.Fields().Len() != 2:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:345
		_go_fuzz_dep_.CoverTab[60971]++
															return errors.New("message must have exactly two fields")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:346
		// _ = "end of CoverTab[60971]"
	case md.ExtensionRanges().Len() > 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:347
		_go_fuzz_dep_.CoverTab[60972]++
															return errors.New("message must not have any extension ranges")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:348
		// _ = "end of CoverTab[60972]"
	case md.Enums().Len()+md.Messages().Len()+md.Extensions().Len() > 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:349
		_go_fuzz_dep_.CoverTab[60973]++
															return errors.New("message must not have any nested declarations")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:350
		// _ = "end of CoverTab[60973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:350
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:350
		_go_fuzz_dep_.CoverTab[60974]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:350
		// _ = "end of CoverTab[60974]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:351
	// _ = "end of CoverTab[60962]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:351
	_go_fuzz_dep_.CoverTab[60963]++
														kf := md.Fields().Get(0)
														vf := md.Fields().Get(1)
														switch {
	case kf.Name() != genid.MapEntry_Key_field_name || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		_go_fuzz_dep_.CoverTab[60979]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		return kf.Number() != genid.MapEntry_Key_field_number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		// _ = "end of CoverTab[60979]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		_go_fuzz_dep_.CoverTab[60980]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		return kf.Cardinality() != protoreflect.Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		// _ = "end of CoverTab[60980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		_go_fuzz_dep_.CoverTab[60981]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		return kf.ContainingOneof() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		// _ = "end of CoverTab[60981]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		_go_fuzz_dep_.CoverTab[60982]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		return kf.HasDefault()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		// _ = "end of CoverTab[60982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:355
		_go_fuzz_dep_.CoverTab[60976]++
															return errors.New("invalid key field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:356
		// _ = "end of CoverTab[60976]"
	case vf.Name() != genid.MapEntry_Value_field_name || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		_go_fuzz_dep_.CoverTab[60983]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		return vf.Number() != genid.MapEntry_Value_field_number
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		// _ = "end of CoverTab[60983]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		_go_fuzz_dep_.CoverTab[60984]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		return vf.Cardinality() != protoreflect.Optional
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		// _ = "end of CoverTab[60984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		_go_fuzz_dep_.CoverTab[60985]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		return vf.ContainingOneof() != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		// _ = "end of CoverTab[60985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		_go_fuzz_dep_.CoverTab[60986]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		return vf.HasDefault()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		// _ = "end of CoverTab[60986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:357
		_go_fuzz_dep_.CoverTab[60977]++
															return errors.New("invalid value field")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:358
		// _ = "end of CoverTab[60977]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:358
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:358
		_go_fuzz_dep_.CoverTab[60978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:358
		// _ = "end of CoverTab[60978]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:359
	// _ = "end of CoverTab[60963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:359
	_go_fuzz_dep_.CoverTab[60964]++
														switch kf.Kind() {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:361
		_go_fuzz_dep_.CoverTab[60987]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:361
		// _ = "end of CoverTab[60987]"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:362
		_go_fuzz_dep_.CoverTab[60988]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:362
		// _ = "end of CoverTab[60988]"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:363
		_go_fuzz_dep_.CoverTab[60989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:363
		// _ = "end of CoverTab[60989]"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:364
		_go_fuzz_dep_.CoverTab[60990]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:364
		// _ = "end of CoverTab[60990]"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:365
		_go_fuzz_dep_.CoverTab[60991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:365
		// _ = "end of CoverTab[60991]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:366
		_go_fuzz_dep_.CoverTab[60992]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:366
		// _ = "end of CoverTab[60992]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:367
		_go_fuzz_dep_.CoverTab[60993]++
															return errors.New("invalid key kind: %v", kf.Kind())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:368
		// _ = "end of CoverTab[60993]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:369
	// _ = "end of CoverTab[60964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:369
	_go_fuzz_dep_.CoverTab[60965]++
														if e := vf.Enum(); e != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		_go_fuzz_dep_.CoverTab[60994]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		return e.Values().Len() > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		// _ = "end of CoverTab[60994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		_go_fuzz_dep_.CoverTab[60995]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		return e.Values().Get(0).Number() != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		// _ = "end of CoverTab[60995]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:370
		_go_fuzz_dep_.CoverTab[60996]++
															return errors.New("map enum value must have zero number for the first value")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:371
		// _ = "end of CoverTab[60996]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:372
		_go_fuzz_dep_.CoverTab[60997]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:372
		// _ = "end of CoverTab[60997]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:372
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:372
	// _ = "end of CoverTab[60965]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:372
	_go_fuzz_dep_.CoverTab[60966]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:373
	// _ = "end of CoverTab[60966]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:374
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_validate.go:374
var _ = _go_fuzz_dep_.CoverTab
