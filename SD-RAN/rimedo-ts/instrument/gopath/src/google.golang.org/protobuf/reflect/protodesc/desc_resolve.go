// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
package protodesc

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:5
)

import (
	"google.golang.org/protobuf/internal/encoding/defval"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	"google.golang.org/protobuf/types/descriptorpb"
)

// resolver is a wrapper around a local registry of declarations within the file
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:17
// and the remote resolver. The remote resolver is restricted to only return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:17
// descriptors that have been imported.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:20
type resolver struct {
	local	descsByName
	remote	Resolver
	imports	importSet

	allowUnresolvable	bool
}

func (r *resolver) resolveMessageDependencies(ms []filedesc.Message, mds []*descriptorpb.DescriptorProto) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:28
	_go_fuzz_dep_.CoverTab[60548]++
														for i, md := range mds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:29
		_go_fuzz_dep_.CoverTab[60550]++
															m := &ms[i]
															for j, fd := range md.GetField() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:31
			_go_fuzz_dep_.CoverTab[60553]++
																f := &m.L2.Fields.List[j]
																if f.L1.Cardinality == protoreflect.Required {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:33
				_go_fuzz_dep_.CoverTab[60557]++
																	m.L2.RequiredNumbers.List = append(m.L2.RequiredNumbers.List, f.L1.Number)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:34
				// _ = "end of CoverTab[60557]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:35
				_go_fuzz_dep_.CoverTab[60558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:35
				// _ = "end of CoverTab[60558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:35
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:35
			// _ = "end of CoverTab[60553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:35
			_go_fuzz_dep_.CoverTab[60554]++
																if fd.OneofIndex != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:36
				_go_fuzz_dep_.CoverTab[60559]++
																	k := int(fd.GetOneofIndex())
																	if !(0 <= k && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:38
					_go_fuzz_dep_.CoverTab[60561]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:38
					return k < len(md.GetOneofDecl())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:38
					// _ = "end of CoverTab[60561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:38
				}()) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:38
					_go_fuzz_dep_.CoverTab[60562]++
																		return errors.New("message field %q has an invalid oneof index: %d", f.FullName(), k)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:39
					// _ = "end of CoverTab[60562]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:40
					_go_fuzz_dep_.CoverTab[60563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:40
					// _ = "end of CoverTab[60563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:40
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:40
				// _ = "end of CoverTab[60559]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:40
				_go_fuzz_dep_.CoverTab[60560]++
																	o := &m.L2.Oneofs.List[k]
																	f.L1.ContainingOneof = o
																	o.L1.Fields.List = append(o.L1.Fields.List, f)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:43
				// _ = "end of CoverTab[60560]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:44
				_go_fuzz_dep_.CoverTab[60564]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:44
				// _ = "end of CoverTab[60564]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:44
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:44
			// _ = "end of CoverTab[60554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:44
			_go_fuzz_dep_.CoverTab[60555]++

																if f.L1.Kind, f.L1.Enum, f.L1.Message, err = r.findTarget(f.Kind(), f.Parent().FullName(), partialName(fd.GetTypeName()), f.IsWeak()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:46
				_go_fuzz_dep_.CoverTab[60565]++
																	return errors.New("message field %q cannot resolve type: %v", f.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:47
				// _ = "end of CoverTab[60565]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:48
				_go_fuzz_dep_.CoverTab[60566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:48
				// _ = "end of CoverTab[60566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:48
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:48
			// _ = "end of CoverTab[60555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:48
			_go_fuzz_dep_.CoverTab[60556]++
																if fd.DefaultValue != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:49
				_go_fuzz_dep_.CoverTab[60567]++
																	v, ev, err := unmarshalDefault(fd.GetDefaultValue(), f, r.allowUnresolvable)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:51
					_go_fuzz_dep_.CoverTab[60569]++
																		return errors.New("message field %q has invalid default: %v", f.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:52
					// _ = "end of CoverTab[60569]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:53
					_go_fuzz_dep_.CoverTab[60570]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:53
					// _ = "end of CoverTab[60570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:53
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:53
				// _ = "end of CoverTab[60567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:53
				_go_fuzz_dep_.CoverTab[60568]++
																	f.L1.Default = filedesc.DefaultValue(v, ev)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:54
				// _ = "end of CoverTab[60568]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:55
				_go_fuzz_dep_.CoverTab[60571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:55
				// _ = "end of CoverTab[60571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:55
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:55
			// _ = "end of CoverTab[60556]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:56
		// _ = "end of CoverTab[60550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:56
		_go_fuzz_dep_.CoverTab[60551]++

															if err := r.resolveMessageDependencies(m.L1.Messages.List, md.GetNestedType()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:58
			_go_fuzz_dep_.CoverTab[60572]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:59
			// _ = "end of CoverTab[60572]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:60
			_go_fuzz_dep_.CoverTab[60573]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:60
			// _ = "end of CoverTab[60573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:60
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:60
		// _ = "end of CoverTab[60551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:60
		_go_fuzz_dep_.CoverTab[60552]++
															if err := r.resolveExtensionDependencies(m.L1.Extensions.List, md.GetExtension()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:61
			_go_fuzz_dep_.CoverTab[60574]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:62
			// _ = "end of CoverTab[60574]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:63
			_go_fuzz_dep_.CoverTab[60575]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:63
			// _ = "end of CoverTab[60575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:63
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:63
		// _ = "end of CoverTab[60552]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:64
	// _ = "end of CoverTab[60548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:64
	_go_fuzz_dep_.CoverTab[60549]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:65
	// _ = "end of CoverTab[60549]"
}

func (r *resolver) resolveExtensionDependencies(xs []filedesc.Extension, xds []*descriptorpb.FieldDescriptorProto) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:68
	_go_fuzz_dep_.CoverTab[60576]++
														for i, xd := range xds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:69
		_go_fuzz_dep_.CoverTab[60578]++
															x := &xs[i]
															if x.L1.Extendee, err = r.findMessageDescriptor(x.Parent().FullName(), partialName(xd.GetExtendee()), false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:71
			_go_fuzz_dep_.CoverTab[60581]++
																return errors.New("extension field %q cannot resolve extendee: %v", x.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:72
			// _ = "end of CoverTab[60581]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:73
			_go_fuzz_dep_.CoverTab[60582]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:73
			// _ = "end of CoverTab[60582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:73
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:73
		// _ = "end of CoverTab[60578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:73
		_go_fuzz_dep_.CoverTab[60579]++
															if x.L1.Kind, x.L2.Enum, x.L2.Message, err = r.findTarget(x.Kind(), x.Parent().FullName(), partialName(xd.GetTypeName()), false); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:74
			_go_fuzz_dep_.CoverTab[60583]++
																return errors.New("extension field %q cannot resolve type: %v", x.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:75
			// _ = "end of CoverTab[60583]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:76
			_go_fuzz_dep_.CoverTab[60584]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:76
			// _ = "end of CoverTab[60584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:76
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:76
		// _ = "end of CoverTab[60579]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:76
		_go_fuzz_dep_.CoverTab[60580]++
															if xd.DefaultValue != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:77
			_go_fuzz_dep_.CoverTab[60585]++
																v, ev, err := unmarshalDefault(xd.GetDefaultValue(), x, r.allowUnresolvable)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:79
				_go_fuzz_dep_.CoverTab[60587]++
																	return errors.New("extension field %q has invalid default: %v", x.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:80
				// _ = "end of CoverTab[60587]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:81
				_go_fuzz_dep_.CoverTab[60588]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:81
				// _ = "end of CoverTab[60588]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:81
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:81
			// _ = "end of CoverTab[60585]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:81
			_go_fuzz_dep_.CoverTab[60586]++
																x.L2.Default = filedesc.DefaultValue(v, ev)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:82
			// _ = "end of CoverTab[60586]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:83
			_go_fuzz_dep_.CoverTab[60589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:83
			// _ = "end of CoverTab[60589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:83
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:83
		// _ = "end of CoverTab[60580]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:84
	// _ = "end of CoverTab[60576]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:84
	_go_fuzz_dep_.CoverTab[60577]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:85
	// _ = "end of CoverTab[60577]"
}

func (r *resolver) resolveServiceDependencies(ss []filedesc.Service, sds []*descriptorpb.ServiceDescriptorProto) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:88
	_go_fuzz_dep_.CoverTab[60590]++
														for i, sd := range sds {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:89
		_go_fuzz_dep_.CoverTab[60592]++
															s := &ss[i]
															for j, md := range sd.GetMethod() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:91
			_go_fuzz_dep_.CoverTab[60593]++
																m := &s.L2.Methods.List[j]
																m.L1.Input, err = r.findMessageDescriptor(m.Parent().FullName(), partialName(md.GetInputType()), false)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:94
				_go_fuzz_dep_.CoverTab[60595]++
																	return errors.New("service method %q cannot resolve input: %v", m.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:95
				// _ = "end of CoverTab[60595]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:96
				_go_fuzz_dep_.CoverTab[60596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:96
				// _ = "end of CoverTab[60596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:96
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:96
			// _ = "end of CoverTab[60593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:96
			_go_fuzz_dep_.CoverTab[60594]++
																m.L1.Output, err = r.findMessageDescriptor(s.FullName(), partialName(md.GetOutputType()), false)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:98
				_go_fuzz_dep_.CoverTab[60597]++
																	return errors.New("service method %q cannot resolve output: %v", m.FullName(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:99
				// _ = "end of CoverTab[60597]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:100
				_go_fuzz_dep_.CoverTab[60598]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:100
				// _ = "end of CoverTab[60598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:100
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:100
			// _ = "end of CoverTab[60594]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:101
		// _ = "end of CoverTab[60592]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:102
	// _ = "end of CoverTab[60590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:102
	_go_fuzz_dep_.CoverTab[60591]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:103
	// _ = "end of CoverTab[60591]"
}

// findTarget finds an enum or message descriptor if k is an enum, message,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:106
// group, or unknown. If unknown, and the name could be resolved, the kind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:106
// returned kind is set based on the type of the resolved descriptor.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:109
func (r *resolver) findTarget(k protoreflect.Kind, scope protoreflect.FullName, ref partialName, isWeak bool) (protoreflect.Kind, protoreflect.EnumDescriptor, protoreflect.MessageDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:109
	_go_fuzz_dep_.CoverTab[60599]++
														switch k {
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:111
		_go_fuzz_dep_.CoverTab[60600]++
															ed, err := r.findEnumDescriptor(scope, ref, isWeak)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:113
			_go_fuzz_dep_.CoverTab[60609]++
																return 0, nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:114
			// _ = "end of CoverTab[60609]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:115
			_go_fuzz_dep_.CoverTab[60610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:115
			// _ = "end of CoverTab[60610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:115
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:115
		// _ = "end of CoverTab[60600]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:115
		_go_fuzz_dep_.CoverTab[60601]++
															return k, ed, nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:116
		// _ = "end of CoverTab[60601]"
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:117
		_go_fuzz_dep_.CoverTab[60602]++
															md, err := r.findMessageDescriptor(scope, ref, isWeak)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:119
			_go_fuzz_dep_.CoverTab[60611]++
																return 0, nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:120
			// _ = "end of CoverTab[60611]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:121
			_go_fuzz_dep_.CoverTab[60612]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:121
			// _ = "end of CoverTab[60612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:121
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:121
		// _ = "end of CoverTab[60602]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:121
		_go_fuzz_dep_.CoverTab[60603]++
															return k, nil, md, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:122
		// _ = "end of CoverTab[60603]"
	case 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:123
		_go_fuzz_dep_.CoverTab[60604]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:126
		d, err := r.findDescriptor(scope, ref)
		if err == protoregistry.NotFound && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
			_go_fuzz_dep_.CoverTab[60613]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
			return (r.allowUnresolvable || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
				_go_fuzz_dep_.CoverTab[60614]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
				return isWeak
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
				// _ = "end of CoverTab[60614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
			// _ = "end of CoverTab[60613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:127
			_go_fuzz_dep_.CoverTab[60615]++
																return k, filedesc.PlaceholderEnum(ref.FullName()), filedesc.PlaceholderMessage(ref.FullName()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:128
			// _ = "end of CoverTab[60615]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:129
			_go_fuzz_dep_.CoverTab[60616]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:129
			if err == protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:129
				_go_fuzz_dep_.CoverTab[60617]++
																	return 0, nil, nil, errors.New("%q not found", ref.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:130
				// _ = "end of CoverTab[60617]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:131
				_go_fuzz_dep_.CoverTab[60618]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:131
				if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:131
					_go_fuzz_dep_.CoverTab[60619]++
																		return 0, nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:132
					// _ = "end of CoverTab[60619]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
					_go_fuzz_dep_.CoverTab[60620]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
					// _ = "end of CoverTab[60620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
				// _ = "end of CoverTab[60618]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
			// _ = "end of CoverTab[60616]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
		// _ = "end of CoverTab[60604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:133
		_go_fuzz_dep_.CoverTab[60605]++
															switch d := d.(type) {
		case protoreflect.EnumDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:135
			_go_fuzz_dep_.CoverTab[60621]++
																return protoreflect.EnumKind, d, nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:136
			// _ = "end of CoverTab[60621]"
		case protoreflect.MessageDescriptor:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:137
			_go_fuzz_dep_.CoverTab[60622]++
																return protoreflect.MessageKind, nil, d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:138
			// _ = "end of CoverTab[60622]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:139
			_go_fuzz_dep_.CoverTab[60623]++
																return 0, nil, nil, errors.New("unknown kind")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:140
			// _ = "end of CoverTab[60623]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:141
		// _ = "end of CoverTab[60605]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:142
		_go_fuzz_dep_.CoverTab[60606]++
															if ref != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:143
			_go_fuzz_dep_.CoverTab[60624]++
																return 0, nil, nil, errors.New("target name cannot be specified for %v", k)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:144
			// _ = "end of CoverTab[60624]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:145
			_go_fuzz_dep_.CoverTab[60625]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:145
			// _ = "end of CoverTab[60625]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:145
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:145
		// _ = "end of CoverTab[60606]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:145
		_go_fuzz_dep_.CoverTab[60607]++
															if !k.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:146
			_go_fuzz_dep_.CoverTab[60626]++
																return 0, nil, nil, errors.New("invalid kind: %d", k)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:147
			// _ = "end of CoverTab[60626]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:148
			_go_fuzz_dep_.CoverTab[60627]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:148
			// _ = "end of CoverTab[60627]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:148
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:148
		// _ = "end of CoverTab[60607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:148
		_go_fuzz_dep_.CoverTab[60608]++
															return k, nil, nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:149
		// _ = "end of CoverTab[60608]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:150
	// _ = "end of CoverTab[60599]"
}

// findDescriptor finds the descriptor by name,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
// which may be a relative name within some scope.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
// Suppose the scope was "fizz.buzz" and the reference was "Foo.Bar",
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
// then the following full names are searched:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
//   - fizz.buzz.Foo.Bar
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
//   - fizz.Foo.Bar
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:153
//   - Foo.Bar
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:161
func (r *resolver) findDescriptor(scope protoreflect.FullName, ref partialName) (protoreflect.Descriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:161
	_go_fuzz_dep_.CoverTab[60628]++
														if !ref.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:162
		_go_fuzz_dep_.CoverTab[60631]++
															return nil, errors.New("invalid name reference: %q", ref)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:163
		// _ = "end of CoverTab[60631]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:164
		_go_fuzz_dep_.CoverTab[60632]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:164
		// _ = "end of CoverTab[60632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:164
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:164
	// _ = "end of CoverTab[60628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:164
	_go_fuzz_dep_.CoverTab[60629]++
														if ref.IsFull() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:165
		_go_fuzz_dep_.CoverTab[60633]++
															scope, ref = "", ref[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:166
		// _ = "end of CoverTab[60633]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:167
		_go_fuzz_dep_.CoverTab[60634]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:167
		// _ = "end of CoverTab[60634]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:167
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:167
	// _ = "end of CoverTab[60629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:167
	_go_fuzz_dep_.CoverTab[60630]++
														var foundButNotImported protoreflect.Descriptor
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:169
		_go_fuzz_dep_.CoverTab[60635]++

															s := protoreflect.FullName(ref)
															if scope != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:172
			_go_fuzz_dep_.CoverTab[60640]++
																s = scope + "." + s
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:173
			// _ = "end of CoverTab[60640]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:174
			_go_fuzz_dep_.CoverTab[60641]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:174
			// _ = "end of CoverTab[60641]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:174
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:174
		// _ = "end of CoverTab[60635]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:174
		_go_fuzz_dep_.CoverTab[60636]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:177
		if d, ok := r.local[s]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:177
			_go_fuzz_dep_.CoverTab[60642]++
																return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:178
			// _ = "end of CoverTab[60642]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:179
			_go_fuzz_dep_.CoverTab[60643]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:179
			// _ = "end of CoverTab[60643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:179
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:179
		// _ = "end of CoverTab[60636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:179
		_go_fuzz_dep_.CoverTab[60637]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:182
		d, err := r.remote.FindDescriptorByName(s)
		if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:183
			_go_fuzz_dep_.CoverTab[60644]++

																if r.imports[d.ParentFile().Path()] {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:185
				_go_fuzz_dep_.CoverTab[60646]++
																	return d, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:186
				// _ = "end of CoverTab[60646]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:187
				_go_fuzz_dep_.CoverTab[60647]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:187
				// _ = "end of CoverTab[60647]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:187
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:187
			// _ = "end of CoverTab[60644]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:187
			_go_fuzz_dep_.CoverTab[60645]++
																foundButNotImported = d
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:188
			// _ = "end of CoverTab[60645]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:189
			_go_fuzz_dep_.CoverTab[60648]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:189
			if err != protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:189
				_go_fuzz_dep_.CoverTab[60649]++
																	return nil, errors.Wrap(err, "%q", s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:190
				// _ = "end of CoverTab[60649]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
				_go_fuzz_dep_.CoverTab[60650]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
				// _ = "end of CoverTab[60650]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
			// _ = "end of CoverTab[60648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
		// _ = "end of CoverTab[60637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:191
		_go_fuzz_dep_.CoverTab[60638]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:194
		if scope == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:194
			_go_fuzz_dep_.CoverTab[60651]++
																if d := foundButNotImported; d != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:195
				_go_fuzz_dep_.CoverTab[60653]++
																	return nil, errors.New("resolved %q, but %q is not imported", d.FullName(), d.ParentFile().Path())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:196
				// _ = "end of CoverTab[60653]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:197
				_go_fuzz_dep_.CoverTab[60654]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:197
				// _ = "end of CoverTab[60654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:197
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:197
			// _ = "end of CoverTab[60651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:197
			_go_fuzz_dep_.CoverTab[60652]++
																return nil, protoregistry.NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:198
			// _ = "end of CoverTab[60652]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:199
			_go_fuzz_dep_.CoverTab[60655]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:199
			// _ = "end of CoverTab[60655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:199
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:199
		// _ = "end of CoverTab[60638]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:199
		_go_fuzz_dep_.CoverTab[60639]++
															scope = scope.Parent()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:200
		// _ = "end of CoverTab[60639]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:201
	// _ = "end of CoverTab[60630]"
}

func (r *resolver) findEnumDescriptor(scope protoreflect.FullName, ref partialName, isWeak bool) (protoreflect.EnumDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:204
	_go_fuzz_dep_.CoverTab[60656]++
														d, err := r.findDescriptor(scope, ref)
														if err == protoregistry.NotFound && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
		_go_fuzz_dep_.CoverTab[60659]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
		return (r.allowUnresolvable || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
			_go_fuzz_dep_.CoverTab[60660]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
			return isWeak
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
			// _ = "end of CoverTab[60660]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
		// _ = "end of CoverTab[60659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:206
		_go_fuzz_dep_.CoverTab[60661]++
															return filedesc.PlaceholderEnum(ref.FullName()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:207
		// _ = "end of CoverTab[60661]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:208
		_go_fuzz_dep_.CoverTab[60662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:208
		if err == protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:208
			_go_fuzz_dep_.CoverTab[60663]++
																return nil, errors.New("%q not found", ref.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:209
			// _ = "end of CoverTab[60663]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:210
			_go_fuzz_dep_.CoverTab[60664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:210
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:210
				_go_fuzz_dep_.CoverTab[60665]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:211
				// _ = "end of CoverTab[60665]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
				_go_fuzz_dep_.CoverTab[60666]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
				// _ = "end of CoverTab[60666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
			// _ = "end of CoverTab[60664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
		// _ = "end of CoverTab[60662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
	// _ = "end of CoverTab[60656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:212
	_go_fuzz_dep_.CoverTab[60657]++
														ed, ok := d.(protoreflect.EnumDescriptor)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:214
		_go_fuzz_dep_.CoverTab[60667]++
															return nil, errors.New("resolved %q, but it is not an enum", d.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:215
		// _ = "end of CoverTab[60667]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:216
		_go_fuzz_dep_.CoverTab[60668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:216
		// _ = "end of CoverTab[60668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:216
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:216
	// _ = "end of CoverTab[60657]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:216
	_go_fuzz_dep_.CoverTab[60658]++
														return ed, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:217
	// _ = "end of CoverTab[60658]"
}

func (r *resolver) findMessageDescriptor(scope protoreflect.FullName, ref partialName, isWeak bool) (protoreflect.MessageDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:220
	_go_fuzz_dep_.CoverTab[60669]++
														d, err := r.findDescriptor(scope, ref)
														if err == protoregistry.NotFound && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
		_go_fuzz_dep_.CoverTab[60672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
		return (r.allowUnresolvable || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
			_go_fuzz_dep_.CoverTab[60673]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
			return isWeak
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
			// _ = "end of CoverTab[60673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
		}())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
		// _ = "end of CoverTab[60672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:222
		_go_fuzz_dep_.CoverTab[60674]++
															return filedesc.PlaceholderMessage(ref.FullName()), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:223
		// _ = "end of CoverTab[60674]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:224
		_go_fuzz_dep_.CoverTab[60675]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:224
		if err == protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:224
			_go_fuzz_dep_.CoverTab[60676]++
																return nil, errors.New("%q not found", ref.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:225
			// _ = "end of CoverTab[60676]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:226
			_go_fuzz_dep_.CoverTab[60677]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:226
			if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:226
				_go_fuzz_dep_.CoverTab[60678]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:227
				// _ = "end of CoverTab[60678]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
				_go_fuzz_dep_.CoverTab[60679]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
				// _ = "end of CoverTab[60679]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
			// _ = "end of CoverTab[60677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
		// _ = "end of CoverTab[60675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
	// _ = "end of CoverTab[60669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:228
	_go_fuzz_dep_.CoverTab[60670]++
														md, ok := d.(protoreflect.MessageDescriptor)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:230
		_go_fuzz_dep_.CoverTab[60680]++
															return nil, errors.New("resolved %q, but it is not an message", d.FullName())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:231
		// _ = "end of CoverTab[60680]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:232
		_go_fuzz_dep_.CoverTab[60681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:232
		// _ = "end of CoverTab[60681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:232
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:232
	// _ = "end of CoverTab[60670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:232
	_go_fuzz_dep_.CoverTab[60671]++
														return md, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:233
	// _ = "end of CoverTab[60671]"
}

// partialName is the partial name. A leading dot means that the name is full,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:236
// otherwise the name is relative to some current scope.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:236
// See google.protobuf.FieldDescriptorProto.type_name.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:239
type partialName string

func (s partialName) IsFull() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:241
	_go_fuzz_dep_.CoverTab[60682]++
														return len(s) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:242
		_go_fuzz_dep_.CoverTab[60683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:242
		return s[0] == '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:242
		// _ = "end of CoverTab[60683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:242
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:242
	// _ = "end of CoverTab[60682]"
}

func (s partialName) IsValid() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:245
	_go_fuzz_dep_.CoverTab[60684]++
														if s.IsFull() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:246
		_go_fuzz_dep_.CoverTab[60686]++
															return protoreflect.FullName(s[1:]).IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:247
		// _ = "end of CoverTab[60686]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:248
		_go_fuzz_dep_.CoverTab[60687]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:248
		// _ = "end of CoverTab[60687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:248
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:248
	// _ = "end of CoverTab[60684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:248
	_go_fuzz_dep_.CoverTab[60685]++
														return protoreflect.FullName(s).IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:249
	// _ = "end of CoverTab[60685]"
}

const unknownPrefix = "*."

// FullName converts the partial name to a full name on a best-effort basis.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:254
// If relative, it creates an invalid full name, using a "*." prefix
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:254
// to indicate that the start of the full name is unknown.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:257
func (s partialName) FullName() protoreflect.FullName {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:257
	_go_fuzz_dep_.CoverTab[60688]++
														if s.IsFull() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:258
		_go_fuzz_dep_.CoverTab[60690]++
															return protoreflect.FullName(s[1:])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:259
		// _ = "end of CoverTab[60690]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:260
		_go_fuzz_dep_.CoverTab[60691]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:260
		// _ = "end of CoverTab[60691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:260
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:260
	// _ = "end of CoverTab[60688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:260
	_go_fuzz_dep_.CoverTab[60689]++
														return protoreflect.FullName(unknownPrefix + s)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:261
	// _ = "end of CoverTab[60689]"
}

func unmarshalDefault(s string, fd protoreflect.FieldDescriptor, allowUnresolvable bool) (protoreflect.Value, protoreflect.EnumValueDescriptor, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:264
	_go_fuzz_dep_.CoverTab[60692]++
														var evs protoreflect.EnumValueDescriptors
														if fd.Enum() != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:266
		_go_fuzz_dep_.CoverTab[60697]++
															evs = fd.Enum().Values()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:267
		// _ = "end of CoverTab[60697]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:268
		_go_fuzz_dep_.CoverTab[60698]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:268
		// _ = "end of CoverTab[60698]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:268
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:268
	// _ = "end of CoverTab[60692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:268
	_go_fuzz_dep_.CoverTab[60693]++
														v, ev, err := defval.Unmarshal(s, fd.Kind(), evs, defval.Descriptor)
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		_go_fuzz_dep_.CoverTab[60699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		return allowUnresolvable
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		// _ = "end of CoverTab[60699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		_go_fuzz_dep_.CoverTab[60700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		return evs != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		// _ = "end of CoverTab[60700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		_go_fuzz_dep_.CoverTab[60701]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		return protoreflect.Name(s).IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		// _ = "end of CoverTab[60701]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:270
		_go_fuzz_dep_.CoverTab[60702]++
															v = protoreflect.ValueOfEnum(0)
															if evs.Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:272
			_go_fuzz_dep_.CoverTab[60704]++
																v = protoreflect.ValueOfEnum(evs.Get(0).Number())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:273
			// _ = "end of CoverTab[60704]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:274
			_go_fuzz_dep_.CoverTab[60705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:274
			// _ = "end of CoverTab[60705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:274
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:274
		// _ = "end of CoverTab[60702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:274
		_go_fuzz_dep_.CoverTab[60703]++
															ev = filedesc.PlaceholderEnumValue(fd.Enum().FullName().Parent().Append(protoreflect.Name(s)))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:275
		// _ = "end of CoverTab[60703]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:276
		_go_fuzz_dep_.CoverTab[60706]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:276
		if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:276
			_go_fuzz_dep_.CoverTab[60707]++
																return v, ev, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:277
			// _ = "end of CoverTab[60707]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
			_go_fuzz_dep_.CoverTab[60708]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
			// _ = "end of CoverTab[60708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
		// _ = "end of CoverTab[60706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
	// _ = "end of CoverTab[60693]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:278
	_go_fuzz_dep_.CoverTab[60694]++
														if fd.Syntax() == protoreflect.Proto3 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:279
		_go_fuzz_dep_.CoverTab[60709]++
															return v, ev, errors.New("cannot be specified under proto3 semantics")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:280
		// _ = "end of CoverTab[60709]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:281
		_go_fuzz_dep_.CoverTab[60710]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:281
		// _ = "end of CoverTab[60710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:281
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:281
	// _ = "end of CoverTab[60694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:281
	_go_fuzz_dep_.CoverTab[60695]++
														if fd.Kind() == protoreflect.MessageKind || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		_go_fuzz_dep_.CoverTab[60711]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		return fd.Kind() == protoreflect.GroupKind
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		// _ = "end of CoverTab[60711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		_go_fuzz_dep_.CoverTab[60712]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		return fd.Cardinality() == protoreflect.Repeated
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		// _ = "end of CoverTab[60712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:282
		_go_fuzz_dep_.CoverTab[60713]++
															return v, ev, errors.New("cannot be specified on composite types")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:283
		// _ = "end of CoverTab[60713]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:284
		_go_fuzz_dep_.CoverTab[60714]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:284
		// _ = "end of CoverTab[60714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:284
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:284
	// _ = "end of CoverTab[60695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:284
	_go_fuzz_dep_.CoverTab[60696]++
														return v, ev, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:285
	// _ = "end of CoverTab[60696]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:286
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/reflect/protodesc/desc_resolve.go:286
var _ = _go_fuzz_dep_.CoverTab
