// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:5
)

import (
	"fmt"
	"math"
	"math/bits"
	"reflect"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/internal/genid"
	"google.golang.org/protobuf/internal/strs"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)

// ValidationStatus is the result of validating the wire-format encoding of a message.
type ValidationStatus int

const (
	// ValidationUnknown indicates that unmarshaling the message might succeed or fail.
	// The validator was unable to render a judgement.
	//
	// The only causes of this status are an aberrant message type appearing somewhere
	// in the message or a failure in the extension resolver.
	ValidationUnknown	ValidationStatus	= iota + 1

	// ValidationInvalid indicates that unmarshaling the message will fail.
	ValidationInvalid

	// ValidationValid indicates that unmarshaling the message will succeed.
	ValidationValid
)

func (v ValidationStatus) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:42
	_go_fuzz_dep_.CoverTab[58720]++
													switch v {
	case ValidationUnknown:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:44
		_go_fuzz_dep_.CoverTab[58721]++
														return "ValidationUnknown"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:45
		// _ = "end of CoverTab[58721]"
	case ValidationInvalid:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:46
		_go_fuzz_dep_.CoverTab[58722]++
														return "ValidationInvalid"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:47
		// _ = "end of CoverTab[58722]"
	case ValidationValid:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:48
		_go_fuzz_dep_.CoverTab[58723]++
														return "ValidationValid"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:49
		// _ = "end of CoverTab[58723]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:50
		_go_fuzz_dep_.CoverTab[58724]++
														return fmt.Sprintf("ValidationStatus(%d)", int(v))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:51
		// _ = "end of CoverTab[58724]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:52
	// _ = "end of CoverTab[58720]"
}

// Validate determines whether the contents of the buffer are a valid wire encoding
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:55
// of the message type.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:55
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:55
// This function is exposed for testing.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:59
func Validate(mt protoreflect.MessageType, in protoiface.UnmarshalInput) (out protoiface.UnmarshalOutput, _ ValidationStatus) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:59
	_go_fuzz_dep_.CoverTab[58725]++
													mi, ok := mt.(*MessageInfo)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:61
		_go_fuzz_dep_.CoverTab[58729]++
														return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:62
		// _ = "end of CoverTab[58729]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:63
		_go_fuzz_dep_.CoverTab[58730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:63
		// _ = "end of CoverTab[58730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:63
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:63
	// _ = "end of CoverTab[58725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:63
	_go_fuzz_dep_.CoverTab[58726]++
													if in.Resolver == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:64
		_go_fuzz_dep_.CoverTab[58731]++
														in.Resolver = protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:65
		// _ = "end of CoverTab[58731]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:66
		_go_fuzz_dep_.CoverTab[58732]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:66
		// _ = "end of CoverTab[58732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:66
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:66
	// _ = "end of CoverTab[58726]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:66
	_go_fuzz_dep_.CoverTab[58727]++
													o, st := mi.validate(in.Buf, 0, unmarshalOptions{
		flags:		in.Flags,
		resolver:	in.Resolver,
	})
	if o.initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:71
		_go_fuzz_dep_.CoverTab[58733]++
														out.Flags |= protoiface.UnmarshalInitialized
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:72
		// _ = "end of CoverTab[58733]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:73
		_go_fuzz_dep_.CoverTab[58734]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:73
		// _ = "end of CoverTab[58734]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:73
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:73
	// _ = "end of CoverTab[58727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:73
	_go_fuzz_dep_.CoverTab[58728]++
													return out, st
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:74
	// _ = "end of CoverTab[58728]"
}

type validationInfo struct {
	mi			*MessageInfo
	typ			validationType
	keyType, valType	validationType

	// For non-required fields, requiredBit is 0.
	//
	// For required fields, requiredBit's nth bit is set, where n is a
	// unique index in the range [0, MessageInfo.numRequiredFields).
	//
	// If there are more than 64 required fields, requiredBit is 0.
	requiredBit	uint64
}

type validationType uint8

const (
	validationTypeOther	validationType	= iota
	validationTypeMessage
	validationTypeGroup
	validationTypeMap
	validationTypeRepeatedVarint
	validationTypeRepeatedFixed32
	validationTypeRepeatedFixed64
	validationTypeVarint
	validationTypeFixed32
	validationTypeFixed64
	validationTypeBytes
	validationTypeUTF8String
	validationTypeMessageSetItem
)

func newFieldValidationInfo(mi *MessageInfo, si structInfo, fd protoreflect.FieldDescriptor, ft reflect.Type) validationInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:109
	_go_fuzz_dep_.CoverTab[58735]++
													var vi validationInfo
													switch {
	case fd.ContainingOneof() != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:112
		_go_fuzz_dep_.CoverTab[58740]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:112
		return !fd.ContainingOneof().IsSynthetic()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:112
		// _ = "end of CoverTab[58740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:112
	}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:112
		_go_fuzz_dep_.CoverTab[58738]++
														switch fd.Kind() {
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:114
			_go_fuzz_dep_.CoverTab[58741]++
															vi.typ = validationTypeMessage
															if ot, ok := si.oneofWrappersByNumber[fd.Number()]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:116
				_go_fuzz_dep_.CoverTab[58745]++
																vi.mi = getMessageInfo(ot.Field(0).Type)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:117
				// _ = "end of CoverTab[58745]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:118
				_go_fuzz_dep_.CoverTab[58746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:118
				// _ = "end of CoverTab[58746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:118
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:118
			// _ = "end of CoverTab[58741]"
		case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:119
			_go_fuzz_dep_.CoverTab[58742]++
															vi.typ = validationTypeGroup
															if ot, ok := si.oneofWrappersByNumber[fd.Number()]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:121
				_go_fuzz_dep_.CoverTab[58747]++
																vi.mi = getMessageInfo(ot.Field(0).Type)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:122
				// _ = "end of CoverTab[58747]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:123
				_go_fuzz_dep_.CoverTab[58748]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:123
				// _ = "end of CoverTab[58748]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:123
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:123
			// _ = "end of CoverTab[58742]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:124
			_go_fuzz_dep_.CoverTab[58743]++
															if strs.EnforceUTF8(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:125
				_go_fuzz_dep_.CoverTab[58749]++
																vi.typ = validationTypeUTF8String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:126
				// _ = "end of CoverTab[58749]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
				_go_fuzz_dep_.CoverTab[58750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
				// _ = "end of CoverTab[58750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
			// _ = "end of CoverTab[58743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
			_go_fuzz_dep_.CoverTab[58744]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:127
			// _ = "end of CoverTab[58744]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:128
		// _ = "end of CoverTab[58738]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:129
		_go_fuzz_dep_.CoverTab[58739]++
														vi = newValidationInfo(fd, ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:130
		// _ = "end of CoverTab[58739]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:131
	// _ = "end of CoverTab[58735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:131
	_go_fuzz_dep_.CoverTab[58736]++
													if fd.Cardinality() == protoreflect.Required {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:132
		_go_fuzz_dep_.CoverTab[58751]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:137
		if mi.numRequiredFields < math.MaxUint8 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:137
			_go_fuzz_dep_.CoverTab[58752]++
															mi.numRequiredFields++
															vi.requiredBit = 1 << (mi.numRequiredFields - 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:139
			// _ = "end of CoverTab[58752]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:140
			_go_fuzz_dep_.CoverTab[58753]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:140
			// _ = "end of CoverTab[58753]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:140
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:140
		// _ = "end of CoverTab[58751]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:141
		_go_fuzz_dep_.CoverTab[58754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:141
		// _ = "end of CoverTab[58754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:141
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:141
	// _ = "end of CoverTab[58736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:141
	_go_fuzz_dep_.CoverTab[58737]++
													return vi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:142
	// _ = "end of CoverTab[58737]"
}

func newValidationInfo(fd protoreflect.FieldDescriptor, ft reflect.Type) validationInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:145
	_go_fuzz_dep_.CoverTab[58755]++
													var vi validationInfo
													switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:148
		_go_fuzz_dep_.CoverTab[58757]++
														switch fd.Kind() {
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:150
			_go_fuzz_dep_.CoverTab[58761]++
															vi.typ = validationTypeMessage
															if ft.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:152
				_go_fuzz_dep_.CoverTab[58765]++
																vi.mi = getMessageInfo(ft.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:153
				// _ = "end of CoverTab[58765]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:154
				_go_fuzz_dep_.CoverTab[58766]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:154
				// _ = "end of CoverTab[58766]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:154
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:154
			// _ = "end of CoverTab[58761]"
		case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:155
			_go_fuzz_dep_.CoverTab[58762]++
															vi.typ = validationTypeGroup
															if ft.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:157
				_go_fuzz_dep_.CoverTab[58767]++
																vi.mi = getMessageInfo(ft.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:158
				// _ = "end of CoverTab[58767]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:159
				_go_fuzz_dep_.CoverTab[58768]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:159
				// _ = "end of CoverTab[58768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:159
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:159
			// _ = "end of CoverTab[58762]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:160
			_go_fuzz_dep_.CoverTab[58763]++
															vi.typ = validationTypeBytes
															if strs.EnforceUTF8(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:162
				_go_fuzz_dep_.CoverTab[58769]++
																vi.typ = validationTypeUTF8String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:163
				// _ = "end of CoverTab[58769]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:164
				_go_fuzz_dep_.CoverTab[58770]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:164
				// _ = "end of CoverTab[58770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:164
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:164
			// _ = "end of CoverTab[58763]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:165
			_go_fuzz_dep_.CoverTab[58764]++
															switch wireTypes[fd.Kind()] {
			case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:167
				_go_fuzz_dep_.CoverTab[58771]++
																vi.typ = validationTypeRepeatedVarint
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:168
				// _ = "end of CoverTab[58771]"
			case protowire.Fixed32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:169
				_go_fuzz_dep_.CoverTab[58772]++
																vi.typ = validationTypeRepeatedFixed32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:170
				// _ = "end of CoverTab[58772]"
			case protowire.Fixed64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:171
				_go_fuzz_dep_.CoverTab[58773]++
																vi.typ = validationTypeRepeatedFixed64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:172
				// _ = "end of CoverTab[58773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:172
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:172
				_go_fuzz_dep_.CoverTab[58774]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:172
				// _ = "end of CoverTab[58774]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:173
			// _ = "end of CoverTab[58764]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:174
		// _ = "end of CoverTab[58757]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:175
		_go_fuzz_dep_.CoverTab[58758]++
														vi.typ = validationTypeMap
														switch fd.MapKey().Kind() {
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:178
			_go_fuzz_dep_.CoverTab[58775]++
															if strs.EnforceUTF8(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:179
				_go_fuzz_dep_.CoverTab[58777]++
																vi.keyType = validationTypeUTF8String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:180
				// _ = "end of CoverTab[58777]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
				_go_fuzz_dep_.CoverTab[58778]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
				// _ = "end of CoverTab[58778]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
			// _ = "end of CoverTab[58775]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
			_go_fuzz_dep_.CoverTab[58776]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:181
			// _ = "end of CoverTab[58776]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:182
		// _ = "end of CoverTab[58758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:182
		_go_fuzz_dep_.CoverTab[58759]++
														switch fd.MapValue().Kind() {
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:184
			_go_fuzz_dep_.CoverTab[58779]++
															vi.valType = validationTypeMessage
															if ft.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:186
				_go_fuzz_dep_.CoverTab[58782]++
																vi.mi = getMessageInfo(ft.Elem())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:187
				// _ = "end of CoverTab[58782]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:188
				_go_fuzz_dep_.CoverTab[58783]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:188
				// _ = "end of CoverTab[58783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:188
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:188
			// _ = "end of CoverTab[58779]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:189
			_go_fuzz_dep_.CoverTab[58780]++
															if strs.EnforceUTF8(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:190
				_go_fuzz_dep_.CoverTab[58784]++
																vi.valType = validationTypeUTF8String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:191
				// _ = "end of CoverTab[58784]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
				_go_fuzz_dep_.CoverTab[58785]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
				// _ = "end of CoverTab[58785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
			// _ = "end of CoverTab[58780]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
			_go_fuzz_dep_.CoverTab[58781]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:192
			// _ = "end of CoverTab[58781]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:193
		// _ = "end of CoverTab[58759]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:194
		_go_fuzz_dep_.CoverTab[58760]++
														switch fd.Kind() {
		case protoreflect.MessageKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:196
			_go_fuzz_dep_.CoverTab[58786]++
															vi.typ = validationTypeMessage
															if !fd.IsWeak() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:198
				_go_fuzz_dep_.CoverTab[58790]++
																vi.mi = getMessageInfo(ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:199
				// _ = "end of CoverTab[58790]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:200
				_go_fuzz_dep_.CoverTab[58791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:200
				// _ = "end of CoverTab[58791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:200
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:200
			// _ = "end of CoverTab[58786]"
		case protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:201
			_go_fuzz_dep_.CoverTab[58787]++
															vi.typ = validationTypeGroup
															vi.mi = getMessageInfo(ft)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:203
			// _ = "end of CoverTab[58787]"
		case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:204
			_go_fuzz_dep_.CoverTab[58788]++
															vi.typ = validationTypeBytes
															if strs.EnforceUTF8(fd) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:206
				_go_fuzz_dep_.CoverTab[58792]++
																vi.typ = validationTypeUTF8String
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:207
				// _ = "end of CoverTab[58792]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:208
				_go_fuzz_dep_.CoverTab[58793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:208
				// _ = "end of CoverTab[58793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:208
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:208
			// _ = "end of CoverTab[58788]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:209
			_go_fuzz_dep_.CoverTab[58789]++
															switch wireTypes[fd.Kind()] {
			case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:211
				_go_fuzz_dep_.CoverTab[58794]++
																vi.typ = validationTypeVarint
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:212
				// _ = "end of CoverTab[58794]"
			case protowire.Fixed32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:213
				_go_fuzz_dep_.CoverTab[58795]++
																vi.typ = validationTypeFixed32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:214
				// _ = "end of CoverTab[58795]"
			case protowire.Fixed64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:215
				_go_fuzz_dep_.CoverTab[58796]++
																vi.typ = validationTypeFixed64
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:216
				// _ = "end of CoverTab[58796]"
			case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:217
				_go_fuzz_dep_.CoverTab[58797]++
																vi.typ = validationTypeBytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:218
				// _ = "end of CoverTab[58797]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:218
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:218
				_go_fuzz_dep_.CoverTab[58798]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:218
				// _ = "end of CoverTab[58798]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:219
			// _ = "end of CoverTab[58789]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:220
		// _ = "end of CoverTab[58760]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:221
	// _ = "end of CoverTab[58755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:221
	_go_fuzz_dep_.CoverTab[58756]++
													return vi
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:222
	// _ = "end of CoverTab[58756]"
}

func (mi *MessageInfo) validate(b []byte, groupTag protowire.Number, opts unmarshalOptions) (out unmarshalOutput, result ValidationStatus) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:225
	_go_fuzz_dep_.CoverTab[58799]++
													mi.init()
													type validationState struct {
		typ			validationType
		keyType, valType	validationType
		endGroup		protowire.Number
		mi			*MessageInfo
		tail			[]byte
		requiredMask		uint64
	}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:237
	states := make([]validationState, 0, 16)
	states = append(states, validationState{
		typ:	validationTypeMessage,
		mi:	mi,
	})
	if groupTag > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:242
		_go_fuzz_dep_.CoverTab[58803]++
														states[0].typ = validationTypeGroup
														states[0].endGroup = groupTag
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:244
		// _ = "end of CoverTab[58803]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:245
		_go_fuzz_dep_.CoverTab[58804]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:245
		// _ = "end of CoverTab[58804]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:245
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:245
	// _ = "end of CoverTab[58799]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:245
	_go_fuzz_dep_.CoverTab[58800]++
													initialized := true
													start := len(b)
State:
	for len(states) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:249
		_go_fuzz_dep_.CoverTab[58805]++
														st := &states[len(states)-1]
														for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:251
			_go_fuzz_dep_.CoverTab[58811]++
			// Parse the tag (field number and wire type).
			var tag uint64
			if b[0] < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:254
				_go_fuzz_dep_.CoverTab[58817]++
																tag = uint64(b[0])
																b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:256
				// _ = "end of CoverTab[58817]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
				_go_fuzz_dep_.CoverTab[58818]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
				if len(b) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
					_go_fuzz_dep_.CoverTab[58819]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
					return b[1] < 128
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
					// _ = "end of CoverTab[58819]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:257
					_go_fuzz_dep_.CoverTab[58820]++
																	tag = uint64(b[0]&0x7f) + uint64(b[1])<<7
																	b = b[2:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:259
					// _ = "end of CoverTab[58820]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:260
					_go_fuzz_dep_.CoverTab[58821]++
																	var n int
																	tag, n = protowire.ConsumeVarint(b)
																	if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:263
						_go_fuzz_dep_.CoverTab[58823]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:264
						// _ = "end of CoverTab[58823]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:265
						_go_fuzz_dep_.CoverTab[58824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:265
						// _ = "end of CoverTab[58824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:265
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:265
					// _ = "end of CoverTab[58821]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:265
					_go_fuzz_dep_.CoverTab[58822]++
																	b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:266
					// _ = "end of CoverTab[58822]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:267
				// _ = "end of CoverTab[58818]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:267
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:267
			// _ = "end of CoverTab[58811]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:267
			_go_fuzz_dep_.CoverTab[58812]++
															var num protowire.Number
															if n := tag >> 3; n < uint64(protowire.MinValidNumber) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:269
				_go_fuzz_dep_.CoverTab[58825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:269
				return n > uint64(protowire.MaxValidNumber)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:269
				// _ = "end of CoverTab[58825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:269
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:269
				_go_fuzz_dep_.CoverTab[58826]++
																return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:270
				// _ = "end of CoverTab[58826]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:271
				_go_fuzz_dep_.CoverTab[58827]++
																num = protowire.Number(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:272
				// _ = "end of CoverTab[58827]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:273
			// _ = "end of CoverTab[58812]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:273
			_go_fuzz_dep_.CoverTab[58813]++
															wtyp := protowire.Type(tag & 7)

															if wtyp == protowire.EndGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:276
				_go_fuzz_dep_.CoverTab[58828]++
																if st.endGroup == num {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:277
					_go_fuzz_dep_.CoverTab[58830]++
																	goto PopState
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:278
					// _ = "end of CoverTab[58830]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:279
					_go_fuzz_dep_.CoverTab[58831]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:279
					// _ = "end of CoverTab[58831]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:279
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:279
				// _ = "end of CoverTab[58828]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:279
				_go_fuzz_dep_.CoverTab[58829]++
																return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:280
				// _ = "end of CoverTab[58829]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:281
				_go_fuzz_dep_.CoverTab[58832]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:281
				// _ = "end of CoverTab[58832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:281
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:281
			// _ = "end of CoverTab[58813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:281
			_go_fuzz_dep_.CoverTab[58814]++
															var vi validationInfo
															switch {
			case st.typ == validationTypeMap:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:284
				_go_fuzz_dep_.CoverTab[58833]++
																switch num {
				case genid.MapEntry_Key_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:286
					_go_fuzz_dep_.CoverTab[58839]++
																	vi.typ = st.keyType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:287
					// _ = "end of CoverTab[58839]"
				case genid.MapEntry_Value_field_number:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:288
					_go_fuzz_dep_.CoverTab[58840]++
																	vi.typ = st.valType
																	vi.mi = st.mi
																	vi.requiredBit = 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:291
					// _ = "end of CoverTab[58840]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:291
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:291
					_go_fuzz_dep_.CoverTab[58841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:291
					// _ = "end of CoverTab[58841]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:292
				// _ = "end of CoverTab[58833]"
			case flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:293
				_go_fuzz_dep_.CoverTab[58842]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:293
				return st.mi.isMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:293
				// _ = "end of CoverTab[58842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:293
			}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:293
				_go_fuzz_dep_.CoverTab[58834]++
																switch num {
				case messageset.FieldItem:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:295
					_go_fuzz_dep_.CoverTab[58843]++
																	vi.typ = validationTypeMessageSetItem
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:296
					// _ = "end of CoverTab[58843]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:296
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:296
					_go_fuzz_dep_.CoverTab[58844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:296
					// _ = "end of CoverTab[58844]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:297
				// _ = "end of CoverTab[58834]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:298
				_go_fuzz_dep_.CoverTab[58835]++
																var f *coderFieldInfo
																if int(num) < len(st.mi.denseCoderFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:300
					_go_fuzz_dep_.CoverTab[58845]++
																	f = st.mi.denseCoderFields[num]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:301
					// _ = "end of CoverTab[58845]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:302
					_go_fuzz_dep_.CoverTab[58846]++
																	f = st.mi.coderFields[num]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:303
					// _ = "end of CoverTab[58846]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:304
				// _ = "end of CoverTab[58835]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:304
				_go_fuzz_dep_.CoverTab[58836]++
																if f != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:305
					_go_fuzz_dep_.CoverTab[58847]++
																	vi = f.validation
																	if vi.typ == validationTypeMessage && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:307
						_go_fuzz_dep_.CoverTab[58849]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:307
						return vi.mi == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:307
						// _ = "end of CoverTab[58849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:307
					}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:307
						_go_fuzz_dep_.CoverTab[58850]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:312
						fd := st.mi.Desc.Fields().ByNumber(num)
						if fd == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:313
							_go_fuzz_dep_.CoverTab[58852]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:313
							return !fd.IsWeak()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:313
							// _ = "end of CoverTab[58852]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:313
						}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:313
							_go_fuzz_dep_.CoverTab[58853]++
																			break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:314
							// _ = "end of CoverTab[58853]"
						} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:315
							_go_fuzz_dep_.CoverTab[58854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:315
							// _ = "end of CoverTab[58854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:315
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:315
						// _ = "end of CoverTab[58850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:315
						_go_fuzz_dep_.CoverTab[58851]++
																		messageName := fd.Message().FullName()
																		messageType, err := protoregistry.GlobalTypes.FindMessageByName(messageName)
																		switch err {
						case nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:319
							_go_fuzz_dep_.CoverTab[58855]++
																			vi.mi, _ = messageType.(*MessageInfo)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:320
							// _ = "end of CoverTab[58855]"
						case protoregistry.NotFound:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:321
							_go_fuzz_dep_.CoverTab[58856]++
																			vi.typ = validationTypeBytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:322
							// _ = "end of CoverTab[58856]"
						default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:323
							_go_fuzz_dep_.CoverTab[58857]++
																			return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:324
							// _ = "end of CoverTab[58857]"
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:325
						// _ = "end of CoverTab[58851]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:326
						_go_fuzz_dep_.CoverTab[58858]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:326
						// _ = "end of CoverTab[58858]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:326
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:326
					// _ = "end of CoverTab[58847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:326
					_go_fuzz_dep_.CoverTab[58848]++
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:327
					// _ = "end of CoverTab[58848]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:328
					_go_fuzz_dep_.CoverTab[58859]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:328
					// _ = "end of CoverTab[58859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:328
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:328
				// _ = "end of CoverTab[58836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:328
				_go_fuzz_dep_.CoverTab[58837]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:337
				xt, err := opts.resolver.FindExtensionByNumber(st.mi.Desc.FullName(), num)
				if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:338
					_go_fuzz_dep_.CoverTab[58860]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:338
					return err != protoregistry.NotFound
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:338
					// _ = "end of CoverTab[58860]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:338
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:338
					_go_fuzz_dep_.CoverTab[58861]++
																	return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:339
					// _ = "end of CoverTab[58861]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:340
					_go_fuzz_dep_.CoverTab[58862]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:340
					// _ = "end of CoverTab[58862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:340
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:340
				// _ = "end of CoverTab[58837]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:340
				_go_fuzz_dep_.CoverTab[58838]++
																if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:341
					_go_fuzz_dep_.CoverTab[58863]++
																	vi = getExtensionFieldInfo(xt).validation
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:342
					// _ = "end of CoverTab[58863]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:343
					_go_fuzz_dep_.CoverTab[58864]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:343
					// _ = "end of CoverTab[58864]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:343
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:343
				// _ = "end of CoverTab[58838]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:344
			// _ = "end of CoverTab[58814]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:344
			_go_fuzz_dep_.CoverTab[58815]++
															if vi.requiredBit != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:345
				_go_fuzz_dep_.CoverTab[58865]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:349
				ok := false
				switch vi.typ {
				case validationTypeVarint:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:351
					_go_fuzz_dep_.CoverTab[58867]++
																	ok = wtyp == protowire.VarintType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:352
					// _ = "end of CoverTab[58867]"
				case validationTypeFixed32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:353
					_go_fuzz_dep_.CoverTab[58868]++
																	ok = wtyp == protowire.Fixed32Type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:354
					// _ = "end of CoverTab[58868]"
				case validationTypeFixed64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:355
					_go_fuzz_dep_.CoverTab[58869]++
																	ok = wtyp == protowire.Fixed64Type
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:356
					// _ = "end of CoverTab[58869]"
				case validationTypeBytes, validationTypeUTF8String, validationTypeMessage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:357
					_go_fuzz_dep_.CoverTab[58870]++
																	ok = wtyp == protowire.BytesType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:358
					// _ = "end of CoverTab[58870]"
				case validationTypeGroup:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:359
					_go_fuzz_dep_.CoverTab[58871]++
																	ok = wtyp == protowire.StartGroupType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:360
					// _ = "end of CoverTab[58871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:360
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:360
					_go_fuzz_dep_.CoverTab[58872]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:360
					// _ = "end of CoverTab[58872]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:361
				// _ = "end of CoverTab[58865]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:361
				_go_fuzz_dep_.CoverTab[58866]++
																if ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:362
					_go_fuzz_dep_.CoverTab[58873]++
																	st.requiredMask |= vi.requiredBit
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:363
					// _ = "end of CoverTab[58873]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:364
					_go_fuzz_dep_.CoverTab[58874]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:364
					// _ = "end of CoverTab[58874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:364
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:364
				// _ = "end of CoverTab[58866]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:365
				_go_fuzz_dep_.CoverTab[58875]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:365
				// _ = "end of CoverTab[58875]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:365
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:365
			// _ = "end of CoverTab[58815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:365
			_go_fuzz_dep_.CoverTab[58816]++

															switch wtyp {
			case protowire.VarintType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:368
				_go_fuzz_dep_.CoverTab[58876]++
																if len(b) >= 10 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:369
					_go_fuzz_dep_.CoverTab[58887]++
																	switch {
					case b[0] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:371
						_go_fuzz_dep_.CoverTab[58888]++
																		b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:372
						// _ = "end of CoverTab[58888]"
					case b[1] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:373
						_go_fuzz_dep_.CoverTab[58889]++
																		b = b[2:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:374
						// _ = "end of CoverTab[58889]"
					case b[2] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:375
						_go_fuzz_dep_.CoverTab[58890]++
																		b = b[3:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:376
						// _ = "end of CoverTab[58890]"
					case b[3] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:377
						_go_fuzz_dep_.CoverTab[58891]++
																		b = b[4:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:378
						// _ = "end of CoverTab[58891]"
					case b[4] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:379
						_go_fuzz_dep_.CoverTab[58892]++
																		b = b[5:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:380
						// _ = "end of CoverTab[58892]"
					case b[5] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:381
						_go_fuzz_dep_.CoverTab[58893]++
																		b = b[6:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:382
						// _ = "end of CoverTab[58893]"
					case b[6] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:383
						_go_fuzz_dep_.CoverTab[58894]++
																		b = b[7:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:384
						// _ = "end of CoverTab[58894]"
					case b[7] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:385
						_go_fuzz_dep_.CoverTab[58895]++
																		b = b[8:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:386
						// _ = "end of CoverTab[58895]"
					case b[8] < 0x80:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:387
						_go_fuzz_dep_.CoverTab[58896]++
																		b = b[9:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:388
						// _ = "end of CoverTab[58896]"
					case b[9] < 0x80 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:389
						_go_fuzz_dep_.CoverTab[58899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:389
						return b[9] < 2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:389
						// _ = "end of CoverTab[58899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:389
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:389
						_go_fuzz_dep_.CoverTab[58897]++
																		b = b[10:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:390
						// _ = "end of CoverTab[58897]"
					default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:391
						_go_fuzz_dep_.CoverTab[58898]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:392
						// _ = "end of CoverTab[58898]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:393
					// _ = "end of CoverTab[58887]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:394
					_go_fuzz_dep_.CoverTab[58900]++
																	switch {
					case len(b) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:396
						_go_fuzz_dep_.CoverTab[58912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:396
						return b[0] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:396
						// _ = "end of CoverTab[58912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:396
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:396
						_go_fuzz_dep_.CoverTab[58901]++
																		b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:397
						// _ = "end of CoverTab[58901]"
					case len(b) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:398
						_go_fuzz_dep_.CoverTab[58913]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:398
						return b[1] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:398
						// _ = "end of CoverTab[58913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:398
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:398
						_go_fuzz_dep_.CoverTab[58902]++
																		b = b[2:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:399
						// _ = "end of CoverTab[58902]"
					case len(b) > 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:400
						_go_fuzz_dep_.CoverTab[58914]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:400
						return b[2] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:400
						// _ = "end of CoverTab[58914]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:400
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:400
						_go_fuzz_dep_.CoverTab[58903]++
																		b = b[3:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:401
						// _ = "end of CoverTab[58903]"
					case len(b) > 3 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:402
						_go_fuzz_dep_.CoverTab[58915]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:402
						return b[3] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:402
						// _ = "end of CoverTab[58915]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:402
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:402
						_go_fuzz_dep_.CoverTab[58904]++
																		b = b[4:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:403
						// _ = "end of CoverTab[58904]"
					case len(b) > 4 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:404
						_go_fuzz_dep_.CoverTab[58916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:404
						return b[4] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:404
						// _ = "end of CoverTab[58916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:404
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:404
						_go_fuzz_dep_.CoverTab[58905]++
																		b = b[5:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:405
						// _ = "end of CoverTab[58905]"
					case len(b) > 5 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:406
						_go_fuzz_dep_.CoverTab[58917]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:406
						return b[5] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:406
						// _ = "end of CoverTab[58917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:406
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:406
						_go_fuzz_dep_.CoverTab[58906]++
																		b = b[6:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:407
						// _ = "end of CoverTab[58906]"
					case len(b) > 6 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:408
						_go_fuzz_dep_.CoverTab[58918]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:408
						return b[6] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:408
						// _ = "end of CoverTab[58918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:408
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:408
						_go_fuzz_dep_.CoverTab[58907]++
																		b = b[7:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:409
						// _ = "end of CoverTab[58907]"
					case len(b) > 7 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:410
						_go_fuzz_dep_.CoverTab[58919]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:410
						return b[7] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:410
						// _ = "end of CoverTab[58919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:410
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:410
						_go_fuzz_dep_.CoverTab[58908]++
																		b = b[8:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:411
						// _ = "end of CoverTab[58908]"
					case len(b) > 8 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:412
						_go_fuzz_dep_.CoverTab[58920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:412
						return b[8] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:412
						// _ = "end of CoverTab[58920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:412
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:412
						_go_fuzz_dep_.CoverTab[58909]++
																		b = b[9:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:413
						// _ = "end of CoverTab[58909]"
					case len(b) > 9 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:414
						_go_fuzz_dep_.CoverTab[58921]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:414
						return b[9] < 2
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:414
						// _ = "end of CoverTab[58921]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:414
					}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:414
						_go_fuzz_dep_.CoverTab[58910]++
																		b = b[10:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:415
						// _ = "end of CoverTab[58910]"
					default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:416
						_go_fuzz_dep_.CoverTab[58911]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:417
						// _ = "end of CoverTab[58911]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:418
					// _ = "end of CoverTab[58900]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:419
				// _ = "end of CoverTab[58876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:419
				_go_fuzz_dep_.CoverTab[58877]++
																continue State
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:420
				// _ = "end of CoverTab[58877]"
			case protowire.BytesType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:421
				_go_fuzz_dep_.CoverTab[58878]++
																var size uint64
																if len(b) >= 1 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:423
					_go_fuzz_dep_.CoverTab[58922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:423
					return b[0] < 0x80
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:423
					// _ = "end of CoverTab[58922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:423
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:423
					_go_fuzz_dep_.CoverTab[58923]++
																	size = uint64(b[0])
																	b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:425
					// _ = "end of CoverTab[58923]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
					_go_fuzz_dep_.CoverTab[58924]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
					if len(b) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
						_go_fuzz_dep_.CoverTab[58925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
						return b[1] < 128
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
						// _ = "end of CoverTab[58925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
					}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:426
						_go_fuzz_dep_.CoverTab[58926]++
																		size = uint64(b[0]&0x7f) + uint64(b[1])<<7
																		b = b[2:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:428
						// _ = "end of CoverTab[58926]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:429
						_go_fuzz_dep_.CoverTab[58927]++
																		var n int
																		size, n = protowire.ConsumeVarint(b)
																		if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:432
							_go_fuzz_dep_.CoverTab[58929]++
																			return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:433
							// _ = "end of CoverTab[58929]"
						} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:434
							_go_fuzz_dep_.CoverTab[58930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:434
							// _ = "end of CoverTab[58930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:434
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:434
						// _ = "end of CoverTab[58927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:434
						_go_fuzz_dep_.CoverTab[58928]++
																		b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:435
						// _ = "end of CoverTab[58928]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:436
					// _ = "end of CoverTab[58924]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:436
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:436
				// _ = "end of CoverTab[58878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:436
				_go_fuzz_dep_.CoverTab[58879]++
																if size > uint64(len(b)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:437
					_go_fuzz_dep_.CoverTab[58931]++
																	return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:438
					// _ = "end of CoverTab[58931]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:439
					_go_fuzz_dep_.CoverTab[58932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:439
					// _ = "end of CoverTab[58932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:439
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:439
				// _ = "end of CoverTab[58879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:439
				_go_fuzz_dep_.CoverTab[58880]++
																v := b[:size]
																b = b[size:]
																switch vi.typ {
				case validationTypeMessage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:443
					_go_fuzz_dep_.CoverTab[58933]++
																	if vi.mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:444
						_go_fuzz_dep_.CoverTab[58942]++
																		return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:445
						// _ = "end of CoverTab[58942]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:446
						_go_fuzz_dep_.CoverTab[58943]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:446
						// _ = "end of CoverTab[58943]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:446
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:446
					// _ = "end of CoverTab[58933]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:446
					_go_fuzz_dep_.CoverTab[58934]++
																	vi.mi.init()
																	fallthrough
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:448
					// _ = "end of CoverTab[58934]"
				case validationTypeMap:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:449
					_go_fuzz_dep_.CoverTab[58935]++
																	if vi.mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:450
						_go_fuzz_dep_.CoverTab[58944]++
																		vi.mi.init()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:451
						// _ = "end of CoverTab[58944]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:452
						_go_fuzz_dep_.CoverTab[58945]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:452
						// _ = "end of CoverTab[58945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:452
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:452
					// _ = "end of CoverTab[58935]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:452
					_go_fuzz_dep_.CoverTab[58936]++
																	states = append(states, validationState{
						typ:		vi.typ,
						keyType:	vi.keyType,
						valType:	vi.valType,
						mi:		vi.mi,
						tail:		b,
					})
																	b = v
																	continue State
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:461
					// _ = "end of CoverTab[58936]"
				case validationTypeRepeatedVarint:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:462
					_go_fuzz_dep_.CoverTab[58937]++

																	for len(v) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:464
						_go_fuzz_dep_.CoverTab[58946]++
																		_, n := protowire.ConsumeVarint(v)
																		if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:466
							_go_fuzz_dep_.CoverTab[58948]++
																			return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:467
							// _ = "end of CoverTab[58948]"
						} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:468
							_go_fuzz_dep_.CoverTab[58949]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:468
							// _ = "end of CoverTab[58949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:468
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:468
						// _ = "end of CoverTab[58946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:468
						_go_fuzz_dep_.CoverTab[58947]++
																		v = v[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:469
						// _ = "end of CoverTab[58947]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:470
					// _ = "end of CoverTab[58937]"
				case validationTypeRepeatedFixed32:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:471
					_go_fuzz_dep_.CoverTab[58938]++

																	if len(v)%4 != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:473
						_go_fuzz_dep_.CoverTab[58950]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:474
						// _ = "end of CoverTab[58950]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:475
						_go_fuzz_dep_.CoverTab[58951]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:475
						// _ = "end of CoverTab[58951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:475
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:475
					// _ = "end of CoverTab[58938]"
				case validationTypeRepeatedFixed64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:476
					_go_fuzz_dep_.CoverTab[58939]++

																	if len(v)%8 != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:478
						_go_fuzz_dep_.CoverTab[58952]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:479
						// _ = "end of CoverTab[58952]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:480
						_go_fuzz_dep_.CoverTab[58953]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:480
						// _ = "end of CoverTab[58953]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:480
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:480
					// _ = "end of CoverTab[58939]"
				case validationTypeUTF8String:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:481
					_go_fuzz_dep_.CoverTab[58940]++
																	if !utf8.Valid(v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:482
						_go_fuzz_dep_.CoverTab[58954]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:483
						// _ = "end of CoverTab[58954]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
						_go_fuzz_dep_.CoverTab[58955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
						// _ = "end of CoverTab[58955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
					// _ = "end of CoverTab[58940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
					_go_fuzz_dep_.CoverTab[58941]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:484
					// _ = "end of CoverTab[58941]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:485
				// _ = "end of CoverTab[58880]"
			case protowire.Fixed32Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:486
				_go_fuzz_dep_.CoverTab[58881]++
																if len(b) < 4 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:487
					_go_fuzz_dep_.CoverTab[58956]++
																	return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:488
					// _ = "end of CoverTab[58956]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:489
					_go_fuzz_dep_.CoverTab[58957]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:489
					// _ = "end of CoverTab[58957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:489
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:489
				// _ = "end of CoverTab[58881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:489
				_go_fuzz_dep_.CoverTab[58882]++
																b = b[4:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:490
				// _ = "end of CoverTab[58882]"
			case protowire.Fixed64Type:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:491
				_go_fuzz_dep_.CoverTab[58883]++
																if len(b) < 8 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:492
					_go_fuzz_dep_.CoverTab[58958]++
																	return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:493
					// _ = "end of CoverTab[58958]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:494
					_go_fuzz_dep_.CoverTab[58959]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:494
					// _ = "end of CoverTab[58959]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:494
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:494
				// _ = "end of CoverTab[58883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:494
				_go_fuzz_dep_.CoverTab[58884]++
																b = b[8:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:495
				// _ = "end of CoverTab[58884]"
			case protowire.StartGroupType:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:496
				_go_fuzz_dep_.CoverTab[58885]++
																switch {
				case vi.typ == validationTypeGroup:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:498
					_go_fuzz_dep_.CoverTab[58960]++
																	if vi.mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:499
						_go_fuzz_dep_.CoverTab[58966]++
																		return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:500
						// _ = "end of CoverTab[58966]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:501
						_go_fuzz_dep_.CoverTab[58967]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:501
						// _ = "end of CoverTab[58967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:501
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:501
					// _ = "end of CoverTab[58960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:501
					_go_fuzz_dep_.CoverTab[58961]++
																	vi.mi.init()
																	states = append(states, validationState{
						typ:		validationTypeGroup,
						mi:		vi.mi,
						endGroup:	num,
					})
																	continue State
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:508
					// _ = "end of CoverTab[58961]"
				case flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:509
					_go_fuzz_dep_.CoverTab[58968]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:509
					return vi.typ == validationTypeMessageSetItem
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:509
					// _ = "end of CoverTab[58968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:509
				}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:509
					_go_fuzz_dep_.CoverTab[58962]++
																	typeid, v, n, err := messageset.ConsumeFieldValue(b, false)
																	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:511
						_go_fuzz_dep_.CoverTab[58969]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:512
						// _ = "end of CoverTab[58969]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:513
						_go_fuzz_dep_.CoverTab[58970]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:513
						// _ = "end of CoverTab[58970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:513
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:513
					// _ = "end of CoverTab[58962]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:513
					_go_fuzz_dep_.CoverTab[58963]++
																	xt, err := opts.resolver.FindExtensionByNumber(st.mi.Desc.FullName(), typeid)
																	switch {
					case err == protoregistry.NotFound:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:516
						_go_fuzz_dep_.CoverTab[58971]++
																		b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:517
						// _ = "end of CoverTab[58971]"
					case err != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:518
						_go_fuzz_dep_.CoverTab[58972]++
																		return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:519
						// _ = "end of CoverTab[58972]"
					default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:520
						_go_fuzz_dep_.CoverTab[58973]++
																		xvi := getExtensionFieldInfo(xt).validation
																		if xvi.mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:522
							_go_fuzz_dep_.CoverTab[58975]++
																			xvi.mi.init()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:523
							// _ = "end of CoverTab[58975]"
						} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:524
							_go_fuzz_dep_.CoverTab[58976]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:524
							// _ = "end of CoverTab[58976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:524
						}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:524
						// _ = "end of CoverTab[58973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:524
						_go_fuzz_dep_.CoverTab[58974]++
																		states = append(states, validationState{
							typ:	xvi.typ,
							mi:	xvi.mi,
							tail:	b[n:],
						})
																		b = v
																		continue State
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:531
						// _ = "end of CoverTab[58974]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:532
					// _ = "end of CoverTab[58963]"
				default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:533
					_go_fuzz_dep_.CoverTab[58964]++
																	n := protowire.ConsumeFieldValue(num, wtyp, b)
																	if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:535
						_go_fuzz_dep_.CoverTab[58977]++
																		return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:536
						// _ = "end of CoverTab[58977]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:537
						_go_fuzz_dep_.CoverTab[58978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:537
						// _ = "end of CoverTab[58978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:537
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:537
					// _ = "end of CoverTab[58964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:537
					_go_fuzz_dep_.CoverTab[58965]++
																	b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:538
					// _ = "end of CoverTab[58965]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:539
				// _ = "end of CoverTab[58885]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:540
				_go_fuzz_dep_.CoverTab[58886]++
																return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:541
				// _ = "end of CoverTab[58886]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:542
			// _ = "end of CoverTab[58816]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:543
		// _ = "end of CoverTab[58805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:543
		_go_fuzz_dep_.CoverTab[58806]++
														if st.endGroup != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:544
			_go_fuzz_dep_.CoverTab[58979]++
															return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:545
			// _ = "end of CoverTab[58979]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:546
			_go_fuzz_dep_.CoverTab[58980]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:546
			// _ = "end of CoverTab[58980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:546
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:546
		// _ = "end of CoverTab[58806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:546
		_go_fuzz_dep_.CoverTab[58807]++
														if len(b) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:547
			_go_fuzz_dep_.CoverTab[58981]++
															return out, ValidationInvalid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:548
			// _ = "end of CoverTab[58981]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:549
			_go_fuzz_dep_.CoverTab[58982]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:549
			// _ = "end of CoverTab[58982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:549
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:549
		// _ = "end of CoverTab[58807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:549
		_go_fuzz_dep_.CoverTab[58808]++
														b = st.tail
	PopState:
		numRequiredFields := 0
		switch st.typ {
		case validationTypeMessage, validationTypeGroup:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:554
			_go_fuzz_dep_.CoverTab[58983]++
															numRequiredFields = int(st.mi.numRequiredFields)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:555
			// _ = "end of CoverTab[58983]"
		case validationTypeMap:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:556
			_go_fuzz_dep_.CoverTab[58984]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:559
			if st.mi != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:559
				_go_fuzz_dep_.CoverTab[58986]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:559
				return st.mi.numRequiredFields > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:559
				// _ = "end of CoverTab[58986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:559
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:559
				_go_fuzz_dep_.CoverTab[58987]++
																numRequiredFields = 1
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:560
				// _ = "end of CoverTab[58987]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
				_go_fuzz_dep_.CoverTab[58988]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
				// _ = "end of CoverTab[58988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
			// _ = "end of CoverTab[58984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
			_go_fuzz_dep_.CoverTab[58985]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:561
			// _ = "end of CoverTab[58985]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:562
		// _ = "end of CoverTab[58808]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:562
		_go_fuzz_dep_.CoverTab[58809]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:566
		if numRequiredFields > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:566
			_go_fuzz_dep_.CoverTab[58989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:566
			return bits.OnesCount64(st.requiredMask) != numRequiredFields
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:566
			// _ = "end of CoverTab[58989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:566
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:566
			_go_fuzz_dep_.CoverTab[58990]++
															initialized = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:567
			// _ = "end of CoverTab[58990]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:568
			_go_fuzz_dep_.CoverTab[58991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:568
			// _ = "end of CoverTab[58991]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:568
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:568
		// _ = "end of CoverTab[58809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:568
		_go_fuzz_dep_.CoverTab[58810]++
														states = states[:len(states)-1]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:569
		// _ = "end of CoverTab[58810]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:570
	// _ = "end of CoverTab[58800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:570
	_go_fuzz_dep_.CoverTab[58801]++
													out.n = start - len(b)
													if initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:572
		_go_fuzz_dep_.CoverTab[58992]++
														out.initialized = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:573
		// _ = "end of CoverTab[58992]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:574
		_go_fuzz_dep_.CoverTab[58993]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:574
		// _ = "end of CoverTab[58993]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:574
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:574
	// _ = "end of CoverTab[58801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:574
	_go_fuzz_dep_.CoverTab[58802]++
													return out, ValidationValid
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:575
	// _ = "end of CoverTab[58802]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:576
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/validate.go:576
var _ = _go_fuzz_dep_.CoverTab
