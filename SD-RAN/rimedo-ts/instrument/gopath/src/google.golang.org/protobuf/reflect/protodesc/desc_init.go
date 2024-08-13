// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
package protodesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:5
)

import (
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/types/descriptorpb"
)

type descsByName map[protoreflect.FullName]protoreflect.Descriptor

func (r descsByName) initEnumDeclarations(eds []*descriptorpb.EnumDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (es []filedesc.Enum, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:19
	_go_fuzz_dep_.CoverTab[60414]++
													es = make([]filedesc.Enum, len(eds))
													for i, ed := range eds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:21
		_go_fuzz_dep_.CoverTab[60416]++
														e := &es[i]
														e.L2 = new(filedesc.EnumL2)
														if e.L0, err = r.makeBase(e, parent, ed.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:24
			_go_fuzz_dep_.CoverTab[60421]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:25
			// _ = "end of CoverTab[60421]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:26
			_go_fuzz_dep_.CoverTab[60422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:26
			// _ = "end of CoverTab[60422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:26
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:26
		// _ = "end of CoverTab[60416]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:26
		_go_fuzz_dep_.CoverTab[60417]++
														if opts := ed.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:27
			_go_fuzz_dep_.CoverTab[60423]++
															opts = proto.Clone(opts).(*descriptorpb.EnumOptions)
															e.L2.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:29
				_go_fuzz_dep_.CoverTab[60424]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:29
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:29
				// _ = "end of CoverTab[60424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:29
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:29
			// _ = "end of CoverTab[60423]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:30
			_go_fuzz_dep_.CoverTab[60425]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:30
			// _ = "end of CoverTab[60425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:30
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:30
		// _ = "end of CoverTab[60417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:30
		_go_fuzz_dep_.CoverTab[60418]++
														for _, s := range ed.GetReservedName() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:31
			_go_fuzz_dep_.CoverTab[60426]++
															e.L2.ReservedNames.List = append(e.L2.ReservedNames.List, protoreflect.Name(s))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:32
			// _ = "end of CoverTab[60426]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:33
		// _ = "end of CoverTab[60418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:33
		_go_fuzz_dep_.CoverTab[60419]++
														for _, rr := range ed.GetReservedRange() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:34
			_go_fuzz_dep_.CoverTab[60427]++
															e.L2.ReservedRanges.List = append(e.L2.ReservedRanges.List, [2]protoreflect.EnumNumber{
				protoreflect.EnumNumber(rr.GetStart()),
				protoreflect.EnumNumber(rr.GetEnd()),
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:38
			// _ = "end of CoverTab[60427]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:39
		// _ = "end of CoverTab[60419]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:39
		_go_fuzz_dep_.CoverTab[60420]++
														if e.L2.Values.List, err = r.initEnumValuesFromDescriptorProto(ed.GetValue(), e, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:40
			_go_fuzz_dep_.CoverTab[60428]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:41
			// _ = "end of CoverTab[60428]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:42
			_go_fuzz_dep_.CoverTab[60429]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:42
			// _ = "end of CoverTab[60429]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:42
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:42
		// _ = "end of CoverTab[60420]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:43
	// _ = "end of CoverTab[60414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:43
	_go_fuzz_dep_.CoverTab[60415]++
													return es, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:44
	// _ = "end of CoverTab[60415]"
}

func (r descsByName) initEnumValuesFromDescriptorProto(vds []*descriptorpb.EnumValueDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (vs []filedesc.EnumValue, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:47
	_go_fuzz_dep_.CoverTab[60430]++
													vs = make([]filedesc.EnumValue, len(vds))
													for i, vd := range vds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:49
		_go_fuzz_dep_.CoverTab[60432]++
														v := &vs[i]
														if v.L0, err = r.makeBase(v, parent, vd.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:51
			_go_fuzz_dep_.CoverTab[60435]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:52
			// _ = "end of CoverTab[60435]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:53
			_go_fuzz_dep_.CoverTab[60436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:53
			// _ = "end of CoverTab[60436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:53
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:53
		// _ = "end of CoverTab[60432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:53
		_go_fuzz_dep_.CoverTab[60433]++
														if opts := vd.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:54
			_go_fuzz_dep_.CoverTab[60437]++
															opts = proto.Clone(opts).(*descriptorpb.EnumValueOptions)
															v.L1.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:56
				_go_fuzz_dep_.CoverTab[60438]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:56
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:56
				// _ = "end of CoverTab[60438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:56
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:56
			// _ = "end of CoverTab[60437]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:57
			_go_fuzz_dep_.CoverTab[60439]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:57
			// _ = "end of CoverTab[60439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:57
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:57
		// _ = "end of CoverTab[60433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:57
		_go_fuzz_dep_.CoverTab[60434]++
														v.L1.Number = protoreflect.EnumNumber(vd.GetNumber())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:58
		// _ = "end of CoverTab[60434]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:59
	// _ = "end of CoverTab[60430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:59
	_go_fuzz_dep_.CoverTab[60431]++
													return vs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:60
	// _ = "end of CoverTab[60431]"
}

func (r descsByName) initMessagesDeclarations(mds []*descriptorpb.DescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (ms []filedesc.Message, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:63
	_go_fuzz_dep_.CoverTab[60440]++
													ms = make([]filedesc.Message, len(mds))
													for i, md := range mds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:65
		_go_fuzz_dep_.CoverTab[60442]++
														m := &ms[i]
														m.L2 = new(filedesc.MessageL2)
														if m.L0, err = r.makeBase(m, parent, md.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:68
			_go_fuzz_dep_.CoverTab[60452]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:69
			// _ = "end of CoverTab[60452]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:70
			_go_fuzz_dep_.CoverTab[60453]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:70
			// _ = "end of CoverTab[60453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:70
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:70
		// _ = "end of CoverTab[60442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:70
		_go_fuzz_dep_.CoverTab[60443]++
														if opts := md.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:71
			_go_fuzz_dep_.CoverTab[60454]++
															opts = proto.Clone(opts).(*descriptorpb.MessageOptions)
															m.L2.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:73
				_go_fuzz_dep_.CoverTab[60456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:73
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:73
				// _ = "end of CoverTab[60456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:73
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:73
			// _ = "end of CoverTab[60454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:73
			_go_fuzz_dep_.CoverTab[60455]++
															m.L1.IsMapEntry = opts.GetMapEntry()
															m.L1.IsMessageSet = opts.GetMessageSetWireFormat()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:75
			// _ = "end of CoverTab[60455]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:76
			_go_fuzz_dep_.CoverTab[60457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:76
			// _ = "end of CoverTab[60457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:76
		// _ = "end of CoverTab[60443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:76
		_go_fuzz_dep_.CoverTab[60444]++
														for _, s := range md.GetReservedName() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:77
			_go_fuzz_dep_.CoverTab[60458]++
															m.L2.ReservedNames.List = append(m.L2.ReservedNames.List, protoreflect.Name(s))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:78
			// _ = "end of CoverTab[60458]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:79
		// _ = "end of CoverTab[60444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:79
		_go_fuzz_dep_.CoverTab[60445]++
														for _, rr := range md.GetReservedRange() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:80
			_go_fuzz_dep_.CoverTab[60459]++
															m.L2.ReservedRanges.List = append(m.L2.ReservedRanges.List, [2]protoreflect.FieldNumber{
				protoreflect.FieldNumber(rr.GetStart()),
				protoreflect.FieldNumber(rr.GetEnd()),
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:84
			// _ = "end of CoverTab[60459]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:85
		// _ = "end of CoverTab[60445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:85
		_go_fuzz_dep_.CoverTab[60446]++
														for _, xr := range md.GetExtensionRange() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:86
			_go_fuzz_dep_.CoverTab[60460]++
															m.L2.ExtensionRanges.List = append(m.L2.ExtensionRanges.List, [2]protoreflect.FieldNumber{
				protoreflect.FieldNumber(xr.GetStart()),
				protoreflect.FieldNumber(xr.GetEnd()),
			})
			var optsFunc func() protoreflect.ProtoMessage
			if opts := xr.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:92
				_go_fuzz_dep_.CoverTab[60462]++
																opts = proto.Clone(opts).(*descriptorpb.ExtensionRangeOptions)
																optsFunc = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:94
					_go_fuzz_dep_.CoverTab[60463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:94
					return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:94
					// _ = "end of CoverTab[60463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:94
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:94
				// _ = "end of CoverTab[60462]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:95
				_go_fuzz_dep_.CoverTab[60464]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:95
				// _ = "end of CoverTab[60464]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:95
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:95
			// _ = "end of CoverTab[60460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:95
			_go_fuzz_dep_.CoverTab[60461]++
															m.L2.ExtensionRangeOptions = append(m.L2.ExtensionRangeOptions, optsFunc)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:96
			// _ = "end of CoverTab[60461]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:97
		// _ = "end of CoverTab[60446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:97
		_go_fuzz_dep_.CoverTab[60447]++
														if m.L2.Fields.List, err = r.initFieldsFromDescriptorProto(md.GetField(), m, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:98
			_go_fuzz_dep_.CoverTab[60465]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:99
			// _ = "end of CoverTab[60465]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:100
			_go_fuzz_dep_.CoverTab[60466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:100
			// _ = "end of CoverTab[60466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:100
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:100
		// _ = "end of CoverTab[60447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:100
		_go_fuzz_dep_.CoverTab[60448]++
															if m.L2.Oneofs.List, err = r.initOneofsFromDescriptorProto(md.GetOneofDecl(), m, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:101
			_go_fuzz_dep_.CoverTab[60467]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:102
			// _ = "end of CoverTab[60467]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:103
			_go_fuzz_dep_.CoverTab[60468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:103
			// _ = "end of CoverTab[60468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:103
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:103
		// _ = "end of CoverTab[60448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:103
		_go_fuzz_dep_.CoverTab[60449]++
															if m.L1.Enums.List, err = r.initEnumDeclarations(md.GetEnumType(), m, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:104
			_go_fuzz_dep_.CoverTab[60469]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:105
			// _ = "end of CoverTab[60469]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:106
			_go_fuzz_dep_.CoverTab[60470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:106
			// _ = "end of CoverTab[60470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:106
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:106
		// _ = "end of CoverTab[60449]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:106
		_go_fuzz_dep_.CoverTab[60450]++
															if m.L1.Messages.List, err = r.initMessagesDeclarations(md.GetNestedType(), m, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:107
			_go_fuzz_dep_.CoverTab[60471]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:108
			// _ = "end of CoverTab[60471]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:109
			_go_fuzz_dep_.CoverTab[60472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:109
			// _ = "end of CoverTab[60472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:109
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:109
		// _ = "end of CoverTab[60450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:109
		_go_fuzz_dep_.CoverTab[60451]++
															if m.L1.Extensions.List, err = r.initExtensionDeclarations(md.GetExtension(), m, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:110
			_go_fuzz_dep_.CoverTab[60473]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:111
			// _ = "end of CoverTab[60473]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:112
			_go_fuzz_dep_.CoverTab[60474]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:112
			// _ = "end of CoverTab[60474]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:112
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:112
		// _ = "end of CoverTab[60451]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:113
	// _ = "end of CoverTab[60440]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:113
	_go_fuzz_dep_.CoverTab[60441]++
														return ms, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:114
	// _ = "end of CoverTab[60441]"
}

func (r descsByName) initFieldsFromDescriptorProto(fds []*descriptorpb.FieldDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (fs []filedesc.Field, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:117
	_go_fuzz_dep_.CoverTab[60475]++
														fs = make([]filedesc.Field, len(fds))
														for i, fd := range fds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:119
		_go_fuzz_dep_.CoverTab[60477]++
															f := &fs[i]
															if f.L0, err = r.makeBase(f, parent, fd.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:121
			_go_fuzz_dep_.CoverTab[60481]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:122
			// _ = "end of CoverTab[60481]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:123
			_go_fuzz_dep_.CoverTab[60482]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:123
			// _ = "end of CoverTab[60482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:123
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:123
		// _ = "end of CoverTab[60477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:123
		_go_fuzz_dep_.CoverTab[60478]++
															f.L1.IsProto3Optional = fd.GetProto3Optional()
															if opts := fd.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:125
			_go_fuzz_dep_.CoverTab[60483]++
																opts = proto.Clone(opts).(*descriptorpb.FieldOptions)
																f.L1.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:127
				_go_fuzz_dep_.CoverTab[60485]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:127
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:127
				// _ = "end of CoverTab[60485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:127
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:127
			// _ = "end of CoverTab[60483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:127
			_go_fuzz_dep_.CoverTab[60484]++
																f.L1.IsWeak = opts.GetWeak()
																f.L1.HasPacked = opts.Packed != nil
																f.L1.IsPacked = opts.GetPacked()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:130
			// _ = "end of CoverTab[60484]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:131
			_go_fuzz_dep_.CoverTab[60486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:131
			// _ = "end of CoverTab[60486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:131
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:131
		// _ = "end of CoverTab[60478]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:131
		_go_fuzz_dep_.CoverTab[60479]++
															f.L1.Number = protoreflect.FieldNumber(fd.GetNumber())
															f.L1.Cardinality = protoreflect.Cardinality(fd.GetLabel())
															if fd.Type != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:134
			_go_fuzz_dep_.CoverTab[60487]++
																f.L1.Kind = protoreflect.Kind(fd.GetType())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:135
			// _ = "end of CoverTab[60487]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:136
			_go_fuzz_dep_.CoverTab[60488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:136
			// _ = "end of CoverTab[60488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:136
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:136
		// _ = "end of CoverTab[60479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:136
		_go_fuzz_dep_.CoverTab[60480]++
															if fd.JsonName != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:137
			_go_fuzz_dep_.CoverTab[60489]++
																f.L1.StringName.InitJSON(fd.GetJsonName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:138
			// _ = "end of CoverTab[60489]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:139
			_go_fuzz_dep_.CoverTab[60490]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:139
			// _ = "end of CoverTab[60490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:139
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:139
		// _ = "end of CoverTab[60480]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:140
	// _ = "end of CoverTab[60475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:140
	_go_fuzz_dep_.CoverTab[60476]++
														return fs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:141
	// _ = "end of CoverTab[60476]"
}

func (r descsByName) initOneofsFromDescriptorProto(ods []*descriptorpb.OneofDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (os []filedesc.Oneof, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:144
	_go_fuzz_dep_.CoverTab[60491]++
														os = make([]filedesc.Oneof, len(ods))
														for i, od := range ods {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:146
		_go_fuzz_dep_.CoverTab[60493]++
															o := &os[i]
															if o.L0, err = r.makeBase(o, parent, od.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:148
			_go_fuzz_dep_.CoverTab[60495]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:149
			// _ = "end of CoverTab[60495]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:150
			_go_fuzz_dep_.CoverTab[60496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:150
			// _ = "end of CoverTab[60496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:150
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:150
		// _ = "end of CoverTab[60493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:150
		_go_fuzz_dep_.CoverTab[60494]++
															if opts := od.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:151
			_go_fuzz_dep_.CoverTab[60497]++
																opts = proto.Clone(opts).(*descriptorpb.OneofOptions)
																o.L1.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:153
				_go_fuzz_dep_.CoverTab[60498]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:153
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:153
				// _ = "end of CoverTab[60498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:153
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:153
			// _ = "end of CoverTab[60497]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:154
			_go_fuzz_dep_.CoverTab[60499]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:154
			// _ = "end of CoverTab[60499]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:154
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:154
		// _ = "end of CoverTab[60494]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:155
	// _ = "end of CoverTab[60491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:155
	_go_fuzz_dep_.CoverTab[60492]++
														return os, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:156
	// _ = "end of CoverTab[60492]"
}

func (r descsByName) initExtensionDeclarations(xds []*descriptorpb.FieldDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (xs []filedesc.Extension, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:159
	_go_fuzz_dep_.CoverTab[60500]++
														xs = make([]filedesc.Extension, len(xds))
														for i, xd := range xds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:161
		_go_fuzz_dep_.CoverTab[60502]++
															x := &xs[i]
															x.L2 = new(filedesc.ExtensionL2)
															if x.L0, err = r.makeBase(x, parent, xd.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:164
			_go_fuzz_dep_.CoverTab[60506]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:165
			// _ = "end of CoverTab[60506]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:166
			_go_fuzz_dep_.CoverTab[60507]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:166
			// _ = "end of CoverTab[60507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:166
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:166
		// _ = "end of CoverTab[60502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:166
		_go_fuzz_dep_.CoverTab[60503]++
															if opts := xd.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:167
			_go_fuzz_dep_.CoverTab[60508]++
																opts = proto.Clone(opts).(*descriptorpb.FieldOptions)
																x.L2.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:169
				_go_fuzz_dep_.CoverTab[60510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:169
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:169
				// _ = "end of CoverTab[60510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:169
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:169
			// _ = "end of CoverTab[60508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:169
			_go_fuzz_dep_.CoverTab[60509]++
																x.L2.IsPacked = opts.GetPacked()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:170
			// _ = "end of CoverTab[60509]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:171
			_go_fuzz_dep_.CoverTab[60511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:171
			// _ = "end of CoverTab[60511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:171
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:171
		// _ = "end of CoverTab[60503]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:171
		_go_fuzz_dep_.CoverTab[60504]++
															x.L1.Number = protoreflect.FieldNumber(xd.GetNumber())
															x.L1.Cardinality = protoreflect.Cardinality(xd.GetLabel())
															if xd.Type != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:174
			_go_fuzz_dep_.CoverTab[60512]++
																x.L1.Kind = protoreflect.Kind(xd.GetType())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:175
			// _ = "end of CoverTab[60512]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:176
			_go_fuzz_dep_.CoverTab[60513]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:176
			// _ = "end of CoverTab[60513]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:176
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:176
		// _ = "end of CoverTab[60504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:176
		_go_fuzz_dep_.CoverTab[60505]++
															if xd.JsonName != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:177
			_go_fuzz_dep_.CoverTab[60514]++
																x.L2.StringName.InitJSON(xd.GetJsonName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:178
			// _ = "end of CoverTab[60514]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:179
			_go_fuzz_dep_.CoverTab[60515]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:179
			// _ = "end of CoverTab[60515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:179
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:179
		// _ = "end of CoverTab[60505]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:180
	// _ = "end of CoverTab[60500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:180
	_go_fuzz_dep_.CoverTab[60501]++
														return xs, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:181
	// _ = "end of CoverTab[60501]"
}

func (r descsByName) initServiceDeclarations(sds []*descriptorpb.ServiceDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (ss []filedesc.Service, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:184
	_go_fuzz_dep_.CoverTab[60516]++
														ss = make([]filedesc.Service, len(sds))
														for i, sd := range sds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:186
		_go_fuzz_dep_.CoverTab[60518]++
															s := &ss[i]
															s.L2 = new(filedesc.ServiceL2)
															if s.L0, err = r.makeBase(s, parent, sd.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:189
			_go_fuzz_dep_.CoverTab[60521]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:190
			// _ = "end of CoverTab[60521]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:191
			_go_fuzz_dep_.CoverTab[60522]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:191
			// _ = "end of CoverTab[60522]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:191
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:191
		// _ = "end of CoverTab[60518]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:191
		_go_fuzz_dep_.CoverTab[60519]++
															if opts := sd.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:192
			_go_fuzz_dep_.CoverTab[60523]++
																opts = proto.Clone(opts).(*descriptorpb.ServiceOptions)
																s.L2.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:194
				_go_fuzz_dep_.CoverTab[60524]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:194
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:194
				// _ = "end of CoverTab[60524]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:194
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:194
			// _ = "end of CoverTab[60523]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:195
			_go_fuzz_dep_.CoverTab[60525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:195
			// _ = "end of CoverTab[60525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:195
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:195
		// _ = "end of CoverTab[60519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:195
		_go_fuzz_dep_.CoverTab[60520]++
															if s.L2.Methods.List, err = r.initMethodsFromDescriptorProto(sd.GetMethod(), s, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:196
			_go_fuzz_dep_.CoverTab[60526]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:197
			// _ = "end of CoverTab[60526]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:198
			_go_fuzz_dep_.CoverTab[60527]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:198
			// _ = "end of CoverTab[60527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:198
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:198
		// _ = "end of CoverTab[60520]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:199
	// _ = "end of CoverTab[60516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:199
	_go_fuzz_dep_.CoverTab[60517]++
														return ss, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:200
	// _ = "end of CoverTab[60517]"
}

func (r descsByName) initMethodsFromDescriptorProto(mds []*descriptorpb.MethodDescriptorProto, parent protoreflect.Descriptor, sb *strs.Builder) (ms []filedesc.Method, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:203
	_go_fuzz_dep_.CoverTab[60528]++
														ms = make([]filedesc.Method, len(mds))
														for i, md := range mds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:205
		_go_fuzz_dep_.CoverTab[60530]++
															m := &ms[i]
															if m.L0, err = r.makeBase(m, parent, md.GetName(), i, sb); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:207
			_go_fuzz_dep_.CoverTab[60533]++
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:208
			// _ = "end of CoverTab[60533]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:209
			_go_fuzz_dep_.CoverTab[60534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:209
			// _ = "end of CoverTab[60534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:209
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:209
		// _ = "end of CoverTab[60530]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:209
		_go_fuzz_dep_.CoverTab[60531]++
															if opts := md.GetOptions(); opts != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:210
			_go_fuzz_dep_.CoverTab[60535]++
																opts = proto.Clone(opts).(*descriptorpb.MethodOptions)
																m.L1.Options = func() protoreflect.ProtoMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:212
				_go_fuzz_dep_.CoverTab[60536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:212
				return opts
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:212
				// _ = "end of CoverTab[60536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:212
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:212
			// _ = "end of CoverTab[60535]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:213
			_go_fuzz_dep_.CoverTab[60537]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:213
			// _ = "end of CoverTab[60537]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:213
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:213
		// _ = "end of CoverTab[60531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:213
		_go_fuzz_dep_.CoverTab[60532]++
															m.L1.IsStreamingClient = md.GetClientStreaming()
															m.L1.IsStreamingServer = md.GetServerStreaming()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:215
		// _ = "end of CoverTab[60532]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:216
	// _ = "end of CoverTab[60528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:216
	_go_fuzz_dep_.CoverTab[60529]++
														return ms, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:217
	// _ = "end of CoverTab[60529]"
}

func (r descsByName) makeBase(child, parent protoreflect.Descriptor, name string, idx int, sb *strs.Builder) (filedesc.BaseL0, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:220
	_go_fuzz_dep_.CoverTab[60538]++
														if !protoreflect.Name(name).IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:221
		_go_fuzz_dep_.CoverTab[60542]++
															return filedesc.BaseL0{}, errors.New("descriptor %q has an invalid nested name: %q", parent.FullName(), name)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:222
		// _ = "end of CoverTab[60542]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:223
		_go_fuzz_dep_.CoverTab[60543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:223
		// _ = "end of CoverTab[60543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:223
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:223
	// _ = "end of CoverTab[60538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:223
	_go_fuzz_dep_.CoverTab[60539]++

	// Derive the full name of the child.
	// Note that enum values are a sibling to the enum parent in the namespace.
	var fullName protoreflect.FullName
	if _, ok := parent.(protoreflect.EnumDescriptor); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:228
		_go_fuzz_dep_.CoverTab[60544]++
															fullName = sb.AppendFullName(parent.FullName().Parent(), protoreflect.Name(name))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:229
		// _ = "end of CoverTab[60544]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:230
		_go_fuzz_dep_.CoverTab[60545]++
															fullName = sb.AppendFullName(parent.FullName(), protoreflect.Name(name))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:231
		// _ = "end of CoverTab[60545]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:232
	// _ = "end of CoverTab[60539]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:232
	_go_fuzz_dep_.CoverTab[60540]++
														if _, ok := r[fullName]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:233
		_go_fuzz_dep_.CoverTab[60546]++
															return filedesc.BaseL0{}, errors.New("descriptor %q already declared", fullName)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:234
		// _ = "end of CoverTab[60546]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:235
		_go_fuzz_dep_.CoverTab[60547]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:235
		// _ = "end of CoverTab[60547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:235
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:235
	// _ = "end of CoverTab[60540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:235
	_go_fuzz_dep_.CoverTab[60541]++
														r[fullName] = child

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:242
	return filedesc.BaseL0{
		FullName:	fullName,
		ParentFile:	parent.ParentFile().(*filedesc.File),
		Parent:		parent,
		Index:		idx,
	}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:247
	// _ = "end of CoverTab[60541]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:248
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_init.go:248
var _ = _go_fuzz_dep_.CoverTab
