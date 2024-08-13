// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:5
)

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/order"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func (o MarshalOptions) sizeMessageSet(m protoreflect.Message) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:17
	_go_fuzz_dep_.CoverTab[51482]++
												m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:18
		_go_fuzz_dep_.CoverTab[51484]++
													size += messageset.SizeField(fd.Number())
													size += protowire.SizeTag(messageset.FieldMessage)
													size += protowire.SizeBytes(o.size(v.Message()))
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:22
		// _ = "end of CoverTab[51484]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:23
	// _ = "end of CoverTab[51482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:23
	_go_fuzz_dep_.CoverTab[51483]++
												size += messageset.SizeUnknown(m.GetUnknown())
												return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:25
	// _ = "end of CoverTab[51483]"
}

func (o MarshalOptions) marshalMessageSet(b []byte, m protoreflect.Message) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:28
	_go_fuzz_dep_.CoverTab[51485]++
												if !flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:29
		_go_fuzz_dep_.CoverTab[51490]++
													return b, errors.New("no support for message_set_wire_format")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:30
		// _ = "end of CoverTab[51490]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:31
		_go_fuzz_dep_.CoverTab[51491]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:31
		// _ = "end of CoverTab[51491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:31
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:31
	// _ = "end of CoverTab[51485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:31
	_go_fuzz_dep_.CoverTab[51486]++
												fieldOrder := order.AnyFieldOrder
												if o.Deterministic {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:33
		_go_fuzz_dep_.CoverTab[51492]++
													fieldOrder = order.NumberFieldOrder
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:34
		// _ = "end of CoverTab[51492]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:35
		_go_fuzz_dep_.CoverTab[51493]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:35
		// _ = "end of CoverTab[51493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:35
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:35
	// _ = "end of CoverTab[51486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:35
	_go_fuzz_dep_.CoverTab[51487]++
												var err error
												order.RangeFields(m, fieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:37
		_go_fuzz_dep_.CoverTab[51494]++
													b, err = o.marshalMessageSetField(b, fd, v)
													return err == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:39
		// _ = "end of CoverTab[51494]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:40
	// _ = "end of CoverTab[51487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:40
	_go_fuzz_dep_.CoverTab[51488]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:41
		_go_fuzz_dep_.CoverTab[51495]++
													return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:42
		// _ = "end of CoverTab[51495]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:43
		_go_fuzz_dep_.CoverTab[51496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:43
		// _ = "end of CoverTab[51496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:43
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:43
	// _ = "end of CoverTab[51488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:43
	_go_fuzz_dep_.CoverTab[51489]++
												return messageset.AppendUnknown(b, m.GetUnknown())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:44
	// _ = "end of CoverTab[51489]"
}

func (o MarshalOptions) marshalMessageSetField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:47
	_go_fuzz_dep_.CoverTab[51497]++
												b = messageset.AppendFieldStart(b, fd.Number())
												b = protowire.AppendTag(b, messageset.FieldMessage, protowire.BytesType)
												b = protowire.AppendVarint(b, uint64(o.Size(value.Message().Interface())))
												b, err := o.marshalMessage(b, value.Message())
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:52
		_go_fuzz_dep_.CoverTab[51499]++
													return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:53
		// _ = "end of CoverTab[51499]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:54
		_go_fuzz_dep_.CoverTab[51500]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:54
		// _ = "end of CoverTab[51500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:54
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:54
	// _ = "end of CoverTab[51497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:54
	_go_fuzz_dep_.CoverTab[51498]++
												b = messageset.AppendFieldEnd(b)
												return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:56
	// _ = "end of CoverTab[51498]"
}

func (o UnmarshalOptions) unmarshalMessageSet(b []byte, m protoreflect.Message) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:59
	_go_fuzz_dep_.CoverTab[51501]++
												if !flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:60
		_go_fuzz_dep_.CoverTab[51503]++
													return errors.New("no support for message_set_wire_format")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:61
		// _ = "end of CoverTab[51503]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:62
		_go_fuzz_dep_.CoverTab[51504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:62
		// _ = "end of CoverTab[51504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:62
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:62
	// _ = "end of CoverTab[51501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:62
	_go_fuzz_dep_.CoverTab[51502]++
												return messageset.Unmarshal(b, false, func(num protowire.Number, v []byte) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:63
		_go_fuzz_dep_.CoverTab[51505]++
													err := o.unmarshalMessageSetField(m, num, v)
													if err == errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:65
			_go_fuzz_dep_.CoverTab[51507]++
														unknown := m.GetUnknown()
														unknown = protowire.AppendTag(unknown, num, protowire.BytesType)
														unknown = protowire.AppendBytes(unknown, v)
														m.SetUnknown(unknown)
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:70
			// _ = "end of CoverTab[51507]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:71
			_go_fuzz_dep_.CoverTab[51508]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:71
			// _ = "end of CoverTab[51508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:71
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:71
		// _ = "end of CoverTab[51505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:71
		_go_fuzz_dep_.CoverTab[51506]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:72
		// _ = "end of CoverTab[51506]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:73
	// _ = "end of CoverTab[51502]"
}

func (o UnmarshalOptions) unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:76
	_go_fuzz_dep_.CoverTab[51509]++
												md := m.Descriptor()
												if !md.ExtensionRanges().Has(num) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:78
		_go_fuzz_dep_.CoverTab[51514]++
													return errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:79
		// _ = "end of CoverTab[51514]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:80
		_go_fuzz_dep_.CoverTab[51515]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:80
		// _ = "end of CoverTab[51515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:80
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:80
	// _ = "end of CoverTab[51509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:80
	_go_fuzz_dep_.CoverTab[51510]++
												xt, err := o.Resolver.FindExtensionByNumber(md.FullName(), num)
												if err == protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:82
		_go_fuzz_dep_.CoverTab[51516]++
													return errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:83
		// _ = "end of CoverTab[51516]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:84
		_go_fuzz_dep_.CoverTab[51517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:84
		// _ = "end of CoverTab[51517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:84
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:84
	// _ = "end of CoverTab[51510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:84
	_go_fuzz_dep_.CoverTab[51511]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:85
		_go_fuzz_dep_.CoverTab[51518]++
													return errors.New("%v: unable to resolve extension %v: %v", md.FullName(), num, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:86
		// _ = "end of CoverTab[51518]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:87
		_go_fuzz_dep_.CoverTab[51519]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:87
		// _ = "end of CoverTab[51519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:87
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:87
	// _ = "end of CoverTab[51511]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:87
	_go_fuzz_dep_.CoverTab[51512]++
												xd := xt.TypeDescriptor()
												if err := o.unmarshalMessage(v, m.Mutable(xd).Message()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:89
		_go_fuzz_dep_.CoverTab[51520]++
													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:90
		// _ = "end of CoverTab[51520]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:91
		_go_fuzz_dep_.CoverTab[51521]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:91
		// _ = "end of CoverTab[51521]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:91
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:91
	// _ = "end of CoverTab[51512]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:91
	_go_fuzz_dep_.CoverTab[51513]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:92
	// _ = "end of CoverTab[51513]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:93
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/messageset.go:93
var _ = _go_fuzz_dep_.CoverTab
