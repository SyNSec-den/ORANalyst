// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:5
// Package messageset encodes and decodes the obsolete MessageSet wire format.
package messageset

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:6
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:6
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:6
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:6
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:6
)

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// The MessageSet wire format is equivalent to a message defined as follows,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
// where each Item defines an extension field with a field number of 'type_id'
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
// and content of 'message'. MessageSet extensions must be non-repeated message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
// fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//	message MessageSet {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//		repeated group Item = 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//			required int32 type_id = 2;
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//			required string message = 3;
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:16
//	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:27
const (
	FieldItem	= protowire.Number(1)
	FieldTypeID	= protowire.Number(2)
	FieldMessage	= protowire.Number(3)
)

// ExtensionName is the field name for extensions of MessageSet.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
// A valid MessageSet extension must be of the form:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//	message MyMessage {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//		extend proto2.bridge.MessageSet {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//			optional MyMessage message_set_extension = 1234;
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//		...
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:33
//	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:43
const ExtensionName = "message_set_extension"

// IsMessageSet returns whether the message uses the MessageSet wire format.
func IsMessageSet(md protoreflect.MessageDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:46
	_go_fuzz_dep_.CoverTab[49012]++
															xmd, ok := md.(interface{ IsMessageSet() bool })
															return ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:48
		_go_fuzz_dep_.CoverTab[49013]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:48
		return xmd.IsMessageSet()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:48
		// _ = "end of CoverTab[49013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:48
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:48
	// _ = "end of CoverTab[49012]"
}

// IsMessageSetExtension reports this field properly extends a MessageSet.
func IsMessageSetExtension(fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:52
	_go_fuzz_dep_.CoverTab[49014]++
															switch {
	case fd.Name() != ExtensionName:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:54
		_go_fuzz_dep_.CoverTab[49016]++
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:55
		// _ = "end of CoverTab[49016]"
	case !IsMessageSet(fd.ContainingMessage()):
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:56
		_go_fuzz_dep_.CoverTab[49017]++
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:57
		// _ = "end of CoverTab[49017]"
	case fd.FullName().Parent() != fd.Message().FullName():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:58
		_go_fuzz_dep_.CoverTab[49018]++
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:59
		// _ = "end of CoverTab[49018]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:59
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:59
		_go_fuzz_dep_.CoverTab[49019]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:59
		// _ = "end of CoverTab[49019]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:60
	// _ = "end of CoverTab[49014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:60
	_go_fuzz_dep_.CoverTab[49015]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:61
	// _ = "end of CoverTab[49015]"
}

// SizeField returns the size of a MessageSet item field containing an extension
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:64
// with the given field number, not counting the contents of the message subfield.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:66
func SizeField(num protowire.Number) int {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:66
	_go_fuzz_dep_.CoverTab[49020]++
															return 2*protowire.SizeTag(FieldItem) + protowire.SizeTag(FieldTypeID) + protowire.SizeVarint(uint64(num))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:67
	// _ = "end of CoverTab[49020]"
}

// Unmarshal parses a MessageSet.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:70
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:70
// It calls fn with the type ID and value of each item in the MessageSet.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:70
// Unknown fields are discarded.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:70
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:70
// If wantLen is true, the item values include the varint length prefix.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:70
// This is ugly, but simplifies the fast-path decoder in internal/impl.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:77
func Unmarshal(b []byte, wantLen bool, fn func(typeID protowire.Number, value []byte) error) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:77
	_go_fuzz_dep_.CoverTab[49021]++
															for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:78
		_go_fuzz_dep_.CoverTab[49023]++
																num, wtyp, n := protowire.ConsumeTag(b)
																if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:80
			_go_fuzz_dep_.CoverTab[49028]++
																	return protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:81
			// _ = "end of CoverTab[49028]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:82
			_go_fuzz_dep_.CoverTab[49029]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:82
			// _ = "end of CoverTab[49029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:82
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:82
		// _ = "end of CoverTab[49023]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:82
		_go_fuzz_dep_.CoverTab[49024]++
																b = b[n:]
																if num != FieldItem || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:84
			_go_fuzz_dep_.CoverTab[49030]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:84
			return wtyp != protowire.StartGroupType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:84
			// _ = "end of CoverTab[49030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:84
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:84
			_go_fuzz_dep_.CoverTab[49031]++
																	n := protowire.ConsumeFieldValue(num, wtyp, b)
																	if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:86
				_go_fuzz_dep_.CoverTab[49033]++
																		return protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:87
				// _ = "end of CoverTab[49033]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:88
				_go_fuzz_dep_.CoverTab[49034]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:88
				// _ = "end of CoverTab[49034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:88
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:88
			// _ = "end of CoverTab[49031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:88
			_go_fuzz_dep_.CoverTab[49032]++
																	b = b[n:]
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:90
			// _ = "end of CoverTab[49032]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:91
			_go_fuzz_dep_.CoverTab[49035]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:91
			// _ = "end of CoverTab[49035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:91
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:91
		// _ = "end of CoverTab[49024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:91
		_go_fuzz_dep_.CoverTab[49025]++
																typeID, value, n, err := ConsumeFieldValue(b, wantLen)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:93
			_go_fuzz_dep_.CoverTab[49036]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:94
			// _ = "end of CoverTab[49036]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:95
			_go_fuzz_dep_.CoverTab[49037]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:95
			// _ = "end of CoverTab[49037]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:95
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:95
		// _ = "end of CoverTab[49025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:95
		_go_fuzz_dep_.CoverTab[49026]++
																b = b[n:]
																if typeID == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:97
			_go_fuzz_dep_.CoverTab[49038]++
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:98
			// _ = "end of CoverTab[49038]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:99
			_go_fuzz_dep_.CoverTab[49039]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:99
			// _ = "end of CoverTab[49039]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:99
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:99
		// _ = "end of CoverTab[49026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:99
		_go_fuzz_dep_.CoverTab[49027]++
																if err := fn(typeID, value); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:100
			_go_fuzz_dep_.CoverTab[49040]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:101
			// _ = "end of CoverTab[49040]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:102
			_go_fuzz_dep_.CoverTab[49041]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:102
			// _ = "end of CoverTab[49041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:102
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:102
		// _ = "end of CoverTab[49027]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:103
	// _ = "end of CoverTab[49021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:103
	_go_fuzz_dep_.CoverTab[49022]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:104
	// _ = "end of CoverTab[49022]"
}

// ConsumeFieldValue parses b as a MessageSet item field value until and including
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:107
// the trailing end group marker. It assumes the start group tag has already been parsed.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:107
// It returns the contents of the type_id and message subfields and the total
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:107
// item length.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:107
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:107
// If wantLen is true, the returned message value includes the length prefix.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:113
func ConsumeFieldValue(b []byte, wantLen bool) (typeid protowire.Number, message []byte, n int, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:113
	_go_fuzz_dep_.CoverTab[49042]++
															ilen := len(b)
															for {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:115
		_go_fuzz_dep_.CoverTab[49043]++
																num, wtyp, n := protowire.ConsumeTag(b)
																if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:117
			_go_fuzz_dep_.CoverTab[49045]++
																	return 0, nil, 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:118
			// _ = "end of CoverTab[49045]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:119
			_go_fuzz_dep_.CoverTab[49046]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:119
			// _ = "end of CoverTab[49046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:119
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:119
		// _ = "end of CoverTab[49043]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:119
		_go_fuzz_dep_.CoverTab[49044]++
																b = b[n:]
																switch {
		case num == FieldItem && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:122
			_go_fuzz_dep_.CoverTab[49057]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:122
			return wtyp == protowire.EndGroupType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:122
			// _ = "end of CoverTab[49057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:122
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:122
			_go_fuzz_dep_.CoverTab[49047]++
																	if wantLen && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:123
				_go_fuzz_dep_.CoverTab[49058]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:123
				return len(message) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:123
				// _ = "end of CoverTab[49058]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:123
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:123
				_go_fuzz_dep_.CoverTab[49059]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:126
				message = protowire.AppendVarint(message, 0)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:126
				// _ = "end of CoverTab[49059]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:127
				_go_fuzz_dep_.CoverTab[49060]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:127
				// _ = "end of CoverTab[49060]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:127
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:127
			// _ = "end of CoverTab[49047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:127
			_go_fuzz_dep_.CoverTab[49048]++
																	return typeid, message, ilen - len(b), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:128
			// _ = "end of CoverTab[49048]"
		case num == FieldTypeID && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:129
			_go_fuzz_dep_.CoverTab[49061]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:129
			return wtyp == protowire.VarintType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:129
			// _ = "end of CoverTab[49061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:129
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:129
			_go_fuzz_dep_.CoverTab[49049]++
																	v, n := protowire.ConsumeVarint(b)
																	if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:131
				_go_fuzz_dep_.CoverTab[49062]++
																		return 0, nil, 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:132
				// _ = "end of CoverTab[49062]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:133
				_go_fuzz_dep_.CoverTab[49063]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:133
				// _ = "end of CoverTab[49063]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:133
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:133
			// _ = "end of CoverTab[49049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:133
			_go_fuzz_dep_.CoverTab[49050]++
																	b = b[n:]
																	if v < 1 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:135
				_go_fuzz_dep_.CoverTab[49064]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:135
				return v > math.MaxInt32
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:135
				// _ = "end of CoverTab[49064]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:135
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:135
				_go_fuzz_dep_.CoverTab[49065]++
																		return 0, nil, 0, errors.New("invalid type_id in message set")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:136
				// _ = "end of CoverTab[49065]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:137
				_go_fuzz_dep_.CoverTab[49066]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:137
				// _ = "end of CoverTab[49066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:137
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:137
			// _ = "end of CoverTab[49050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:137
			_go_fuzz_dep_.CoverTab[49051]++
																	typeid = protowire.Number(v)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:138
			// _ = "end of CoverTab[49051]"
		case num == FieldMessage && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:139
			_go_fuzz_dep_.CoverTab[49067]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:139
			return wtyp == protowire.BytesType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:139
			// _ = "end of CoverTab[49067]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:139
		}():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:139
			_go_fuzz_dep_.CoverTab[49052]++
																	m, n := protowire.ConsumeBytes(b)
																	if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:141
				_go_fuzz_dep_.CoverTab[49068]++
																		return 0, nil, 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:142
				// _ = "end of CoverTab[49068]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:143
				_go_fuzz_dep_.CoverTab[49069]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:143
				// _ = "end of CoverTab[49069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:143
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:143
			// _ = "end of CoverTab[49052]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:143
			_go_fuzz_dep_.CoverTab[49053]++
																	if message == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:144
				_go_fuzz_dep_.CoverTab[49070]++
																		if wantLen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:145
					_go_fuzz_dep_.CoverTab[49071]++
																			message = b[:n:n]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:146
					// _ = "end of CoverTab[49071]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:147
					_go_fuzz_dep_.CoverTab[49072]++
																			message = m[:len(m):len(m)]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:148
					// _ = "end of CoverTab[49072]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:149
				// _ = "end of CoverTab[49070]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:150
				_go_fuzz_dep_.CoverTab[49073]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:158
				if wantLen {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:158
					_go_fuzz_dep_.CoverTab[49074]++
																			_, nn := protowire.ConsumeVarint(message)
																			m0 := message[nn:]
																			message = nil
																			message = protowire.AppendVarint(message, uint64(len(m0)+len(m)))
																			message = append(message, m0...)
																			message = append(message, m...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:164
					// _ = "end of CoverTab[49074]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:165
					_go_fuzz_dep_.CoverTab[49075]++
																			message = append(message, m...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:166
					// _ = "end of CoverTab[49075]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:167
				// _ = "end of CoverTab[49073]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:168
			// _ = "end of CoverTab[49053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:168
			_go_fuzz_dep_.CoverTab[49054]++
																	b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:169
			// _ = "end of CoverTab[49054]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:170
			_go_fuzz_dep_.CoverTab[49055]++

																	n := protowire.ConsumeFieldValue(num, wtyp, b)
																	if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:173
				_go_fuzz_dep_.CoverTab[49076]++
																		return 0, nil, 0, protowire.ParseError(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:174
				// _ = "end of CoverTab[49076]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:175
				_go_fuzz_dep_.CoverTab[49077]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:175
				// _ = "end of CoverTab[49077]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:175
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:175
			// _ = "end of CoverTab[49055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:175
			_go_fuzz_dep_.CoverTab[49056]++
																	b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:176
			// _ = "end of CoverTab[49056]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:177
		// _ = "end of CoverTab[49044]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:178
	// _ = "end of CoverTab[49042]"
}

// AppendFieldStart appends the start of a MessageSet item field containing
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:181
// an extension with the given number. The caller must add the message
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:181
// subfield (including the tag).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:184
func AppendFieldStart(b []byte, num protowire.Number) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:184
	_go_fuzz_dep_.CoverTab[49078]++
															b = protowire.AppendTag(b, FieldItem, protowire.StartGroupType)
															b = protowire.AppendTag(b, FieldTypeID, protowire.VarintType)
															b = protowire.AppendVarint(b, uint64(num))
															return b
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:188
	// _ = "end of CoverTab[49078]"
}

// AppendFieldEnd appends the trailing end group marker for a MessageSet item field.
func AppendFieldEnd(b []byte) []byte {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:192
	_go_fuzz_dep_.CoverTab[49079]++
															return protowire.AppendTag(b, FieldItem, protowire.EndGroupType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:193
	// _ = "end of CoverTab[49079]"
}

// SizeUnknown returns the size of an unknown fields section in MessageSet format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:196
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:196
// See AppendUnknown.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:199
func SizeUnknown(unknown []byte) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:199
	_go_fuzz_dep_.CoverTab[49080]++
															for len(unknown) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:200
		_go_fuzz_dep_.CoverTab[49082]++
																num, typ, n := protowire.ConsumeTag(unknown)
																if n < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:202
			_go_fuzz_dep_.CoverTab[49085]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:202
			return typ != protowire.BytesType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:202
			// _ = "end of CoverTab[49085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:202
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:202
			_go_fuzz_dep_.CoverTab[49086]++
																	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:203
			// _ = "end of CoverTab[49086]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:204
			_go_fuzz_dep_.CoverTab[49087]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:204
			// _ = "end of CoverTab[49087]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:204
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:204
		// _ = "end of CoverTab[49082]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:204
		_go_fuzz_dep_.CoverTab[49083]++
																unknown = unknown[n:]
																_, n = protowire.ConsumeBytes(unknown)
																if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:207
			_go_fuzz_dep_.CoverTab[49088]++
																	return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:208
			// _ = "end of CoverTab[49088]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:209
			_go_fuzz_dep_.CoverTab[49089]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:209
			// _ = "end of CoverTab[49089]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:209
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:209
		// _ = "end of CoverTab[49083]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:209
		_go_fuzz_dep_.CoverTab[49084]++
																unknown = unknown[n:]
																size += SizeField(num) + protowire.SizeTag(FieldMessage) + n
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:211
		// _ = "end of CoverTab[49084]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:212
	// _ = "end of CoverTab[49080]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:212
	_go_fuzz_dep_.CoverTab[49081]++
															return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:213
	// _ = "end of CoverTab[49081]"
}

// AppendUnknown appends unknown fields to b in MessageSet format.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
// For historic reasons, unresolved items in a MessageSet are stored in a
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
// message's unknown fields section in non-MessageSet format. That is, an
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
// unknown item with typeID T and value V appears in the unknown fields as
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
// a field with number T and value V.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:216
// This function converts the unknown fields back into MessageSet form.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:224
func AppendUnknown(b, unknown []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:224
	_go_fuzz_dep_.CoverTab[49090]++
															for len(unknown) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:225
		_go_fuzz_dep_.CoverTab[49092]++
																num, typ, n := protowire.ConsumeTag(unknown)
																if n < 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:227
			_go_fuzz_dep_.CoverTab[49095]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:227
			return typ != protowire.BytesType
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:227
			// _ = "end of CoverTab[49095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:227
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:227
			_go_fuzz_dep_.CoverTab[49096]++
																	return nil, errors.New("invalid data in message set unknown fields")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:228
			// _ = "end of CoverTab[49096]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:229
			_go_fuzz_dep_.CoverTab[49097]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:229
			// _ = "end of CoverTab[49097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:229
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:229
		// _ = "end of CoverTab[49092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:229
		_go_fuzz_dep_.CoverTab[49093]++
																unknown = unknown[n:]
																_, n = protowire.ConsumeBytes(unknown)
																if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:232
			_go_fuzz_dep_.CoverTab[49098]++
																	return nil, errors.New("invalid data in message set unknown fields")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:233
			// _ = "end of CoverTab[49098]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:234
			_go_fuzz_dep_.CoverTab[49099]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:234
			// _ = "end of CoverTab[49099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:234
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:234
		// _ = "end of CoverTab[49093]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:234
		_go_fuzz_dep_.CoverTab[49094]++
																b = AppendFieldStart(b, num)
																b = protowire.AppendTag(b, FieldMessage, protowire.BytesType)
																b = append(b, unknown[:n]...)
																b = AppendFieldEnd(b)
																unknown = unknown[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:239
		// _ = "end of CoverTab[49094]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:240
	// _ = "end of CoverTab[49090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:240
	_go_fuzz_dep_.CoverTab[49091]++
															return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:241
	// _ = "end of CoverTab[49091]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:242
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/encoding/messageset/messageset.go:242
var _ = _go_fuzz_dep_.CoverTab
