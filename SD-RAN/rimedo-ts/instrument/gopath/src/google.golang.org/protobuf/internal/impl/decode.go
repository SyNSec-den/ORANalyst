// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:5
)

import (
	"math/bits"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)

var errDecode = errors.New("cannot parse invalid wire-format data")
var errRecursionDepth = errors.New("exceeded maximum recursion depth")

type unmarshalOptions struct {
	flags		protoiface.UnmarshalInputFlags
	resolver	interface {
		FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)
		FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
	}
	depth	int
}

func (o unmarshalOptions) Options() proto.UnmarshalOptions {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:31
	_go_fuzz_dep_.CoverTab[57174]++
													return proto.UnmarshalOptions{
		Merge:		true,
		AllowPartial:	true,
		DiscardUnknown:	o.DiscardUnknown(),
		Resolver:	o.resolver,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:37
	// _ = "end of CoverTab[57174]"
}

func (o unmarshalOptions) DiscardUnknown() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:40
	_go_fuzz_dep_.CoverTab[57175]++
													return o.flags&protoiface.UnmarshalDiscardUnknown != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:41
	// _ = "end of CoverTab[57175]"
}

func (o unmarshalOptions) IsDefault() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:44
	_go_fuzz_dep_.CoverTab[57176]++
													return o.flags == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:45
		_go_fuzz_dep_.CoverTab[57177]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:45
		return o.resolver == protoregistry.GlobalTypes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:45
		// _ = "end of CoverTab[57177]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:45
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:45
	// _ = "end of CoverTab[57176]"
}

var lazyUnmarshalOptions = unmarshalOptions{
	resolver:	protoregistry.GlobalTypes,
	depth:		protowire.DefaultRecursionLimit,
}

type unmarshalOutput struct {
	n		int	// number of bytes consumed
	initialized	bool
}

// unmarshal is protoreflect.Methods.Unmarshal.
func (mi *MessageInfo) unmarshal(in protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:59
	_go_fuzz_dep_.CoverTab[57178]++
													var p pointer
													if ms, ok := in.Message.(*messageState); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:61
		_go_fuzz_dep_.CoverTab[57181]++
														p = ms.pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:62
		// _ = "end of CoverTab[57181]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:63
		_go_fuzz_dep_.CoverTab[57182]++
														p = in.Message.(*messageReflectWrapper).pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:64
		// _ = "end of CoverTab[57182]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:65
	// _ = "end of CoverTab[57178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:65
	_go_fuzz_dep_.CoverTab[57179]++
													out, err := mi.unmarshalPointer(in.Buf, p, 0, unmarshalOptions{
		flags:		in.Flags,
		resolver:	in.Resolver,
		depth:		in.Depth,
	})
	var flags protoiface.UnmarshalOutputFlags
	if out.initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:72
		_go_fuzz_dep_.CoverTab[57183]++
														flags |= protoiface.UnmarshalInitialized
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:73
		// _ = "end of CoverTab[57183]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:74
		_go_fuzz_dep_.CoverTab[57184]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:74
		// _ = "end of CoverTab[57184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:74
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:74
	// _ = "end of CoverTab[57179]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:74
	_go_fuzz_dep_.CoverTab[57180]++
													return protoiface.UnmarshalOutput{
		Flags: flags,
	}, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:77
	// _ = "end of CoverTab[57180]"
}

// errUnknown is returned during unmarshaling to indicate a parse error that
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:80
// should result in a field being placed in the unknown fields section (for example,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:80
// when the wire type doesn't match) as opposed to the entire unmarshal operation
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:80
// failing (for example, when a field extends past the available input).
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:80
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:80
// This is a sentinel error which should never be visible to the user.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:86
var errUnknown = errors.New("unknown")

func (mi *MessageInfo) unmarshalPointer(b []byte, p pointer, groupTag protowire.Number, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:88
	_go_fuzz_dep_.CoverTab[57185]++
													mi.init()
													opts.depth--
													if opts.depth < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:91
		_go_fuzz_dep_.CoverTab[57192]++
														return out, errRecursionDepth
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:92
		// _ = "end of CoverTab[57192]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:93
		_go_fuzz_dep_.CoverTab[57193]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:93
		// _ = "end of CoverTab[57193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:93
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:93
	// _ = "end of CoverTab[57185]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:93
	_go_fuzz_dep_.CoverTab[57186]++
													if flags.ProtoLegacy && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:94
		_go_fuzz_dep_.CoverTab[57194]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:94
		return mi.isMessageSet
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:94
		// _ = "end of CoverTab[57194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:94
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:94
		_go_fuzz_dep_.CoverTab[57195]++
														return unmarshalMessageSet(mi, b, p, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:95
		// _ = "end of CoverTab[57195]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:96
		_go_fuzz_dep_.CoverTab[57196]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:96
		// _ = "end of CoverTab[57196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:96
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:96
	// _ = "end of CoverTab[57186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:96
	_go_fuzz_dep_.CoverTab[57187]++
													initialized := true
													var requiredMask uint64
													var exts *map[int32]ExtensionField
													start := len(b)
													for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:101
		_go_fuzz_dep_.CoverTab[57197]++
		// Parse the tag (field number and wire type).
		var tag uint64
		if b[0] < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:104
			_go_fuzz_dep_.CoverTab[57204]++
															tag = uint64(b[0])
															b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:106
			// _ = "end of CoverTab[57204]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
			_go_fuzz_dep_.CoverTab[57205]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
			if len(b) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
				_go_fuzz_dep_.CoverTab[57206]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
				return b[1] < 128
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
				// _ = "end of CoverTab[57206]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:107
				_go_fuzz_dep_.CoverTab[57207]++
																tag = uint64(b[0]&0x7f) + uint64(b[1])<<7
																b = b[2:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:109
				// _ = "end of CoverTab[57207]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:110
				_go_fuzz_dep_.CoverTab[57208]++
																var n int
																tag, n = protowire.ConsumeVarint(b)
																if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:113
					_go_fuzz_dep_.CoverTab[57210]++
																	return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:114
					// _ = "end of CoverTab[57210]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:115
					_go_fuzz_dep_.CoverTab[57211]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:115
					// _ = "end of CoverTab[57211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:115
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:115
				// _ = "end of CoverTab[57208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:115
				_go_fuzz_dep_.CoverTab[57209]++
																b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:116
				// _ = "end of CoverTab[57209]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:117
			// _ = "end of CoverTab[57205]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:117
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:117
		// _ = "end of CoverTab[57197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:117
		_go_fuzz_dep_.CoverTab[57198]++
														var num protowire.Number
														if n := tag >> 3; n < uint64(protowire.MinValidNumber) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:119
			_go_fuzz_dep_.CoverTab[57212]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:119
			return n > uint64(protowire.MaxValidNumber)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:119
			// _ = "end of CoverTab[57212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:119
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:119
			_go_fuzz_dep_.CoverTab[57213]++
															return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:120
			// _ = "end of CoverTab[57213]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:121
			_go_fuzz_dep_.CoverTab[57214]++
															num = protowire.Number(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:122
			// _ = "end of CoverTab[57214]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:123
		// _ = "end of CoverTab[57198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:123
		_go_fuzz_dep_.CoverTab[57199]++
														wtyp := protowire.Type(tag & 7)

														if wtyp == protowire.EndGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:126
			_go_fuzz_dep_.CoverTab[57215]++
															if num != groupTag {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:127
				_go_fuzz_dep_.CoverTab[57217]++
																return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:128
				// _ = "end of CoverTab[57217]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:129
				_go_fuzz_dep_.CoverTab[57218]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:129
				// _ = "end of CoverTab[57218]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:129
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:129
			// _ = "end of CoverTab[57215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:129
			_go_fuzz_dep_.CoverTab[57216]++
															groupTag = 0
															break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:131
			// _ = "end of CoverTab[57216]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:132
			_go_fuzz_dep_.CoverTab[57219]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:132
			// _ = "end of CoverTab[57219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:132
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:132
		// _ = "end of CoverTab[57199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:132
		_go_fuzz_dep_.CoverTab[57200]++

														var f *coderFieldInfo
														if int(num) < len(mi.denseCoderFields) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:135
			_go_fuzz_dep_.CoverTab[57220]++
															f = mi.denseCoderFields[num]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:136
			// _ = "end of CoverTab[57220]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:137
			_go_fuzz_dep_.CoverTab[57221]++
															f = mi.coderFields[num]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:138
			// _ = "end of CoverTab[57221]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:139
		// _ = "end of CoverTab[57200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:139
		_go_fuzz_dep_.CoverTab[57201]++
														var n int
														err := errUnknown
														switch {
		case f != nil:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:143
			_go_fuzz_dep_.CoverTab[57222]++
															if f.funcs.unmarshal == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:144
				_go_fuzz_dep_.CoverTab[57229]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:145
				// _ = "end of CoverTab[57229]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:146
				_go_fuzz_dep_.CoverTab[57230]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:146
				// _ = "end of CoverTab[57230]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:146
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:146
			// _ = "end of CoverTab[57222]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:146
			_go_fuzz_dep_.CoverTab[57223]++
															var o unmarshalOutput
															o, err = f.funcs.unmarshal(b, p.Apply(f.offset), wtyp, f, opts)
															n = o.n
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:150
				_go_fuzz_dep_.CoverTab[57231]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:151
				// _ = "end of CoverTab[57231]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:152
				_go_fuzz_dep_.CoverTab[57232]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:152
				// _ = "end of CoverTab[57232]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:152
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:152
			// _ = "end of CoverTab[57223]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:152
			_go_fuzz_dep_.CoverTab[57224]++
															requiredMask |= f.validation.requiredBit
															if f.funcs.isInit != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:154
				_go_fuzz_dep_.CoverTab[57233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:154
				return !o.initialized
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:154
				// _ = "end of CoverTab[57233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:154
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:154
				_go_fuzz_dep_.CoverTab[57234]++
																initialized = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:155
				// _ = "end of CoverTab[57234]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:156
				_go_fuzz_dep_.CoverTab[57235]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:156
				// _ = "end of CoverTab[57235]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:156
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:156
			// _ = "end of CoverTab[57224]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:157
			_go_fuzz_dep_.CoverTab[57225]++

															if exts == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:159
				_go_fuzz_dep_.CoverTab[57236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:159
				return mi.extensionOffset.IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:159
				// _ = "end of CoverTab[57236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:159
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:159
				_go_fuzz_dep_.CoverTab[57237]++
																exts = p.Apply(mi.extensionOffset).Extensions()
																if *exts == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:161
					_go_fuzz_dep_.CoverTab[57238]++
																	*exts = make(map[int32]ExtensionField)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:162
					// _ = "end of CoverTab[57238]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:163
					_go_fuzz_dep_.CoverTab[57239]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:163
					// _ = "end of CoverTab[57239]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:163
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:163
				// _ = "end of CoverTab[57237]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:164
				_go_fuzz_dep_.CoverTab[57240]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:164
				// _ = "end of CoverTab[57240]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:164
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:164
			// _ = "end of CoverTab[57225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:164
			_go_fuzz_dep_.CoverTab[57226]++
															if exts == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:165
				_go_fuzz_dep_.CoverTab[57241]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:166
				// _ = "end of CoverTab[57241]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:167
				_go_fuzz_dep_.CoverTab[57242]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:167
				// _ = "end of CoverTab[57242]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:167
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:167
			// _ = "end of CoverTab[57226]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:167
			_go_fuzz_dep_.CoverTab[57227]++
															var o unmarshalOutput
															o, err = mi.unmarshalExtension(b, num, wtyp, *exts, opts)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:170
				_go_fuzz_dep_.CoverTab[57243]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:171
				// _ = "end of CoverTab[57243]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:172
				_go_fuzz_dep_.CoverTab[57244]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:172
				// _ = "end of CoverTab[57244]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:172
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:172
			// _ = "end of CoverTab[57227]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:172
			_go_fuzz_dep_.CoverTab[57228]++
															n = o.n
															if !o.initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:174
				_go_fuzz_dep_.CoverTab[57245]++
																initialized = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:175
				// _ = "end of CoverTab[57245]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:176
				_go_fuzz_dep_.CoverTab[57246]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:176
				// _ = "end of CoverTab[57246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:176
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:176
			// _ = "end of CoverTab[57228]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:177
		// _ = "end of CoverTab[57201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:177
		_go_fuzz_dep_.CoverTab[57202]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:178
			_go_fuzz_dep_.CoverTab[57247]++
															if err != errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:179
				_go_fuzz_dep_.CoverTab[57250]++
																return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:180
				// _ = "end of CoverTab[57250]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:181
				_go_fuzz_dep_.CoverTab[57251]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:181
				// _ = "end of CoverTab[57251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:181
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:181
			// _ = "end of CoverTab[57247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:181
			_go_fuzz_dep_.CoverTab[57248]++
															n = protowire.ConsumeFieldValue(num, wtyp, b)
															if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:183
				_go_fuzz_dep_.CoverTab[57252]++
																return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:184
				// _ = "end of CoverTab[57252]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:185
				_go_fuzz_dep_.CoverTab[57253]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:185
				// _ = "end of CoverTab[57253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:185
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:185
			// _ = "end of CoverTab[57248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:185
			_go_fuzz_dep_.CoverTab[57249]++
															if !opts.DiscardUnknown() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:186
				_go_fuzz_dep_.CoverTab[57254]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:186
				return mi.unknownOffset.IsValid()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:186
				// _ = "end of CoverTab[57254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:186
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:186
				_go_fuzz_dep_.CoverTab[57255]++
																u := mi.mutableUnknownBytes(p)
																*u = protowire.AppendTag(*u, num, wtyp)
																*u = append(*u, b[:n]...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:189
				// _ = "end of CoverTab[57255]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:190
				_go_fuzz_dep_.CoverTab[57256]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:190
				// _ = "end of CoverTab[57256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:190
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:190
			// _ = "end of CoverTab[57249]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:191
			_go_fuzz_dep_.CoverTab[57257]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:191
			// _ = "end of CoverTab[57257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:191
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:191
		// _ = "end of CoverTab[57202]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:191
		_go_fuzz_dep_.CoverTab[57203]++
														b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:192
		// _ = "end of CoverTab[57203]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:193
	// _ = "end of CoverTab[57187]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:193
	_go_fuzz_dep_.CoverTab[57188]++
													if groupTag != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:194
		_go_fuzz_dep_.CoverTab[57258]++
														return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:195
		// _ = "end of CoverTab[57258]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:196
		_go_fuzz_dep_.CoverTab[57259]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:196
		// _ = "end of CoverTab[57259]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:196
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:196
	// _ = "end of CoverTab[57188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:196
	_go_fuzz_dep_.CoverTab[57189]++
													if mi.numRequiredFields > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:197
		_go_fuzz_dep_.CoverTab[57260]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:197
		return bits.OnesCount64(requiredMask) != int(mi.numRequiredFields)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:197
		// _ = "end of CoverTab[57260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:197
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:197
		_go_fuzz_dep_.CoverTab[57261]++
														initialized = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:198
		// _ = "end of CoverTab[57261]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:199
		_go_fuzz_dep_.CoverTab[57262]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:199
		// _ = "end of CoverTab[57262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:199
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:199
	// _ = "end of CoverTab[57189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:199
	_go_fuzz_dep_.CoverTab[57190]++
													if initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:200
		_go_fuzz_dep_.CoverTab[57263]++
														out.initialized = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:201
		// _ = "end of CoverTab[57263]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:202
		_go_fuzz_dep_.CoverTab[57264]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:202
		// _ = "end of CoverTab[57264]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:202
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:202
	// _ = "end of CoverTab[57190]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:202
	_go_fuzz_dep_.CoverTab[57191]++
													out.n = start - len(b)
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:204
	// _ = "end of CoverTab[57191]"
}

func (mi *MessageInfo) unmarshalExtension(b []byte, num protowire.Number, wtyp protowire.Type, exts map[int32]ExtensionField, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:207
	_go_fuzz_dep_.CoverTab[57265]++
													x := exts[int32(num)]
													xt := x.Type()
													if xt == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:210
		_go_fuzz_dep_.CoverTab[57272]++
														var err error
														xt, err = opts.resolver.FindExtensionByNumber(mi.Desc.FullName(), num)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:213
			_go_fuzz_dep_.CoverTab[57273]++
															if err == protoregistry.NotFound {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:214
				_go_fuzz_dep_.CoverTab[57275]++
																return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:215
				// _ = "end of CoverTab[57275]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:216
				_go_fuzz_dep_.CoverTab[57276]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:216
				// _ = "end of CoverTab[57276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:216
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:216
			// _ = "end of CoverTab[57273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:216
			_go_fuzz_dep_.CoverTab[57274]++
															return out, errors.New("%v: unable to resolve extension %v: %v", mi.Desc.FullName(), num, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:217
			// _ = "end of CoverTab[57274]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:218
			_go_fuzz_dep_.CoverTab[57277]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:218
			// _ = "end of CoverTab[57277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:218
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:218
		// _ = "end of CoverTab[57272]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:219
		_go_fuzz_dep_.CoverTab[57278]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:219
		// _ = "end of CoverTab[57278]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:219
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:219
	// _ = "end of CoverTab[57265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:219
	_go_fuzz_dep_.CoverTab[57266]++
													xi := getExtensionFieldInfo(xt)
													if xi.funcs.unmarshal == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:221
		_go_fuzz_dep_.CoverTab[57279]++
														return out, errUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:222
		// _ = "end of CoverTab[57279]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:223
		_go_fuzz_dep_.CoverTab[57280]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:223
		// _ = "end of CoverTab[57280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:223
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:223
	// _ = "end of CoverTab[57266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:223
	_go_fuzz_dep_.CoverTab[57267]++
													if flags.LazyUnmarshalExtensions {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:224
		_go_fuzz_dep_.CoverTab[57281]++
														if opts.IsDefault() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:225
			_go_fuzz_dep_.CoverTab[57282]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:225
			return x.canLazy(xt)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:225
			// _ = "end of CoverTab[57282]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:225
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:225
			_go_fuzz_dep_.CoverTab[57283]++
															out, valid := skipExtension(b, xi, num, wtyp, opts)
															switch valid {
			case ValidationValid:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:228
				_go_fuzz_dep_.CoverTab[57284]++
																if out.initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:229
					_go_fuzz_dep_.CoverTab[57288]++
																	x.appendLazyBytes(xt, xi, num, wtyp, b[:out.n])
																	exts[int32(num)] = x
																	return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:232
					// _ = "end of CoverTab[57288]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:233
					_go_fuzz_dep_.CoverTab[57289]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:233
					// _ = "end of CoverTab[57289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:233
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:233
				// _ = "end of CoverTab[57284]"
			case ValidationInvalid:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:234
				_go_fuzz_dep_.CoverTab[57285]++
																return out, errDecode
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:235
				// _ = "end of CoverTab[57285]"
			case ValidationUnknown:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:236
				_go_fuzz_dep_.CoverTab[57286]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:236
				// _ = "end of CoverTab[57286]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:236
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:236
				_go_fuzz_dep_.CoverTab[57287]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:236
				// _ = "end of CoverTab[57287]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:237
			// _ = "end of CoverTab[57283]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:238
			_go_fuzz_dep_.CoverTab[57290]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:238
			// _ = "end of CoverTab[57290]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:238
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:238
		// _ = "end of CoverTab[57281]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:239
		_go_fuzz_dep_.CoverTab[57291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:239
		// _ = "end of CoverTab[57291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:239
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:239
	// _ = "end of CoverTab[57267]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:239
	_go_fuzz_dep_.CoverTab[57268]++
													ival := x.Value()
													if !ival.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:241
		_go_fuzz_dep_.CoverTab[57292]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:241
		return xi.unmarshalNeedsValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:241
		// _ = "end of CoverTab[57292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:241
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:241
		_go_fuzz_dep_.CoverTab[57293]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:245
		ival = xt.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:245
		// _ = "end of CoverTab[57293]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:246
		_go_fuzz_dep_.CoverTab[57294]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:246
		// _ = "end of CoverTab[57294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:246
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:246
	// _ = "end of CoverTab[57268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:246
	_go_fuzz_dep_.CoverTab[57269]++
													v, out, err := xi.funcs.unmarshal(b, ival, num, wtyp, opts)
													if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:248
		_go_fuzz_dep_.CoverTab[57295]++
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:249
		// _ = "end of CoverTab[57295]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:250
		_go_fuzz_dep_.CoverTab[57296]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:250
		// _ = "end of CoverTab[57296]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:250
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:250
	// _ = "end of CoverTab[57269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:250
	_go_fuzz_dep_.CoverTab[57270]++
													if xi.funcs.isInit == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:251
		_go_fuzz_dep_.CoverTab[57297]++
														out.initialized = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:252
		// _ = "end of CoverTab[57297]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:253
		_go_fuzz_dep_.CoverTab[57298]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:253
		// _ = "end of CoverTab[57298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:253
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:253
	// _ = "end of CoverTab[57270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:253
	_go_fuzz_dep_.CoverTab[57271]++
													x.Set(xt, v)
													exts[int32(num)] = x
													return out, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:256
	// _ = "end of CoverTab[57271]"
}

func skipExtension(b []byte, xi *extensionFieldInfo, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, _ ValidationStatus) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:259
	_go_fuzz_dep_.CoverTab[57299]++
													if xi.validation.mi == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:260
		_go_fuzz_dep_.CoverTab[57301]++
														return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:261
		// _ = "end of CoverTab[57301]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:262
		_go_fuzz_dep_.CoverTab[57302]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:262
		// _ = "end of CoverTab[57302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:262
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:262
	// _ = "end of CoverTab[57299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:262
	_go_fuzz_dep_.CoverTab[57300]++
													xi.validation.mi.init()
													switch xi.validation.typ {
	case validationTypeMessage:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:265
		_go_fuzz_dep_.CoverTab[57303]++
														if wtyp != protowire.BytesType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:266
			_go_fuzz_dep_.CoverTab[57309]++
															return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:267
			// _ = "end of CoverTab[57309]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:268
			_go_fuzz_dep_.CoverTab[57310]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:268
			// _ = "end of CoverTab[57310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:268
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:268
		// _ = "end of CoverTab[57303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:268
		_go_fuzz_dep_.CoverTab[57304]++
														v, n := protowire.ConsumeBytes(b)
														if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:270
			_go_fuzz_dep_.CoverTab[57311]++
															return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:271
			// _ = "end of CoverTab[57311]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:272
			_go_fuzz_dep_.CoverTab[57312]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:272
			// _ = "end of CoverTab[57312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:272
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:272
		// _ = "end of CoverTab[57304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:272
		_go_fuzz_dep_.CoverTab[57305]++
														out, st := xi.validation.mi.validate(v, 0, opts)
														out.n = n
														return out, st
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:275
		// _ = "end of CoverTab[57305]"
	case validationTypeGroup:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:276
		_go_fuzz_dep_.CoverTab[57306]++
														if wtyp != protowire.StartGroupType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:277
			_go_fuzz_dep_.CoverTab[57313]++
															return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:278
			// _ = "end of CoverTab[57313]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:279
			_go_fuzz_dep_.CoverTab[57314]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:279
			// _ = "end of CoverTab[57314]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:279
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:279
		// _ = "end of CoverTab[57306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:279
		_go_fuzz_dep_.CoverTab[57307]++
														out, st := xi.validation.mi.validate(b, num, opts)
														return out, st
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:281
		// _ = "end of CoverTab[57307]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:282
		_go_fuzz_dep_.CoverTab[57308]++
														return out, ValidationUnknown
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:283
		// _ = "end of CoverTab[57308]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:284
	// _ = "end of CoverTab[57300]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:285
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/decode.go:285
var _ = _go_fuzz_dep_.CoverTab
