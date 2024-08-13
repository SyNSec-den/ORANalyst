// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:5
)

import (
	"sort"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
)

func sizeMessageSet(mi *MessageInfo, p pointer, opts marshalOptions) (size int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:16
	_go_fuzz_dep_.CoverTab[56452]++
														if !flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:17
		_go_fuzz_dep_.CoverTab[56456]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:18
		// _ = "end of CoverTab[56456]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:19
		_go_fuzz_dep_.CoverTab[56457]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:19
		// _ = "end of CoverTab[56457]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:19
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:19
	// _ = "end of CoverTab[56452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:19
	_go_fuzz_dep_.CoverTab[56453]++

														ext := *p.Apply(mi.extensionOffset).Extensions()
														for _, x := range ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:22
		_go_fuzz_dep_.CoverTab[56458]++
															xi := getExtensionFieldInfo(x.Type())
															if xi.funcs.size == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:24
			_go_fuzz_dep_.CoverTab[56460]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:25
			// _ = "end of CoverTab[56460]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:26
			_go_fuzz_dep_.CoverTab[56461]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:26
			// _ = "end of CoverTab[56461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:26
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:26
		// _ = "end of CoverTab[56458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:26
		_go_fuzz_dep_.CoverTab[56459]++
															num, _ := protowire.DecodeTag(xi.wiretag)
															size += messageset.SizeField(num)
															size += xi.funcs.size(x.Value(), protowire.SizeTag(messageset.FieldMessage), opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:29
		// _ = "end of CoverTab[56459]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:30
	// _ = "end of CoverTab[56453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:30
	_go_fuzz_dep_.CoverTab[56454]++

														if u := mi.getUnknownBytes(p); u != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:32
		_go_fuzz_dep_.CoverTab[56462]++
															size += messageset.SizeUnknown(*u)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:33
		// _ = "end of CoverTab[56462]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:34
		_go_fuzz_dep_.CoverTab[56463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:34
		// _ = "end of CoverTab[56463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:34
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:34
	// _ = "end of CoverTab[56454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:34
	_go_fuzz_dep_.CoverTab[56455]++

														return size
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:36
	// _ = "end of CoverTab[56455]"
}

func marshalMessageSet(mi *MessageInfo, b []byte, p pointer, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:39
	_go_fuzz_dep_.CoverTab[56464]++
														if !flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:40
		_go_fuzz_dep_.CoverTab[56468]++
															return b, errors.New("no support for message_set_wire_format")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:41
		// _ = "end of CoverTab[56468]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:42
		_go_fuzz_dep_.CoverTab[56469]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:42
		// _ = "end of CoverTab[56469]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:42
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:42
	// _ = "end of CoverTab[56464]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:42
	_go_fuzz_dep_.CoverTab[56465]++

														ext := *p.Apply(mi.extensionOffset).Extensions()
														switch len(ext) {
	case 0:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:46
		_go_fuzz_dep_.CoverTab[56470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:46
		// _ = "end of CoverTab[56470]"
	case 1:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:47
		_go_fuzz_dep_.CoverTab[56471]++

															for _, x := range ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:49
			_go_fuzz_dep_.CoverTab[56474]++
																var err error
																b, err = marshalMessageSetField(mi, b, x, opts)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:52
				_go_fuzz_dep_.CoverTab[56475]++
																	return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:53
				// _ = "end of CoverTab[56475]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:54
				_go_fuzz_dep_.CoverTab[56476]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:54
				// _ = "end of CoverTab[56476]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:54
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:54
			// _ = "end of CoverTab[56474]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:55
		// _ = "end of CoverTab[56471]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:56
		_go_fuzz_dep_.CoverTab[56472]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:59
		keys := make([]int, 0, len(ext))
		for k := range ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:60
			_go_fuzz_dep_.CoverTab[56477]++
																keys = append(keys, int(k))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:61
			// _ = "end of CoverTab[56477]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:62
		// _ = "end of CoverTab[56472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:62
		_go_fuzz_dep_.CoverTab[56473]++
															sort.Ints(keys)
															for _, k := range keys {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:64
			_go_fuzz_dep_.CoverTab[56478]++
																var err error
																b, err = marshalMessageSetField(mi, b, ext[int32(k)], opts)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:67
				_go_fuzz_dep_.CoverTab[56479]++
																	return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:68
				// _ = "end of CoverTab[56479]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:69
				_go_fuzz_dep_.CoverTab[56480]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:69
				// _ = "end of CoverTab[56480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:69
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:69
			// _ = "end of CoverTab[56478]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:70
		// _ = "end of CoverTab[56473]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:71
	// _ = "end of CoverTab[56465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:71
	_go_fuzz_dep_.CoverTab[56466]++

														if u := mi.getUnknownBytes(p); u != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:73
		_go_fuzz_dep_.CoverTab[56481]++
															var err error
															b, err = messageset.AppendUnknown(b, *u)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:76
			_go_fuzz_dep_.CoverTab[56482]++
																return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:77
			// _ = "end of CoverTab[56482]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:78
			_go_fuzz_dep_.CoverTab[56483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:78
			// _ = "end of CoverTab[56483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:78
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:78
		// _ = "end of CoverTab[56481]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:79
		_go_fuzz_dep_.CoverTab[56484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:79
		// _ = "end of CoverTab[56484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:79
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:79
	// _ = "end of CoverTab[56466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:79
	_go_fuzz_dep_.CoverTab[56467]++

														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:81
	// _ = "end of CoverTab[56467]"
}

func marshalMessageSetField(mi *MessageInfo, b []byte, x ExtensionField, opts marshalOptions) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:84
	_go_fuzz_dep_.CoverTab[56485]++
														xi := getExtensionFieldInfo(x.Type())
														num, _ := protowire.DecodeTag(xi.wiretag)
														b = messageset.AppendFieldStart(b, num)
														b, err := xi.funcs.marshal(b, x.Value(), protowire.EncodeTag(messageset.FieldMessage, protowire.BytesType), opts)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:89
		_go_fuzz_dep_.CoverTab[56487]++
															return b, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:90
		// _ = "end of CoverTab[56487]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:91
		_go_fuzz_dep_.CoverTab[56488]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:91
		// _ = "end of CoverTab[56488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:91
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:91
	// _ = "end of CoverTab[56485]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:91
	_go_fuzz_dep_.CoverTab[56486]++
														b = messageset.AppendFieldEnd(b)
														return b, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:93
	// _ = "end of CoverTab[56486]"
}

func unmarshalMessageSet(mi *MessageInfo, b []byte, p pointer, opts unmarshalOptions) (out unmarshalOutput, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:96
	_go_fuzz_dep_.CoverTab[56489]++
														if !flags.ProtoLegacy {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:97
		_go_fuzz_dep_.CoverTab[56493]++
															return out, errors.New("no support for message_set_wire_format")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:98
		// _ = "end of CoverTab[56493]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:99
		_go_fuzz_dep_.CoverTab[56494]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:99
		// _ = "end of CoverTab[56494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:99
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:99
	// _ = "end of CoverTab[56489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:99
	_go_fuzz_dep_.CoverTab[56490]++

														ep := p.Apply(mi.extensionOffset).Extensions()
														if *ep == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:102
		_go_fuzz_dep_.CoverTab[56495]++
															*ep = make(map[int32]ExtensionField)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:103
		// _ = "end of CoverTab[56495]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:104
		_go_fuzz_dep_.CoverTab[56496]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:104
		// _ = "end of CoverTab[56496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:104
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:104
	// _ = "end of CoverTab[56490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:104
	_go_fuzz_dep_.CoverTab[56491]++
														ext := *ep
														initialized := true
														err = messageset.Unmarshal(b, true, func(num protowire.Number, v []byte) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:107
		_go_fuzz_dep_.CoverTab[56497]++
															o, err := mi.unmarshalExtension(v, num, protowire.BytesType, ext, opts)
															if err == errUnknown {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:109
			_go_fuzz_dep_.CoverTab[56500]++
																u := mi.mutableUnknownBytes(p)
																*u = protowire.AppendTag(*u, num, protowire.BytesType)
																*u = append(*u, v...)
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:113
			// _ = "end of CoverTab[56500]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:114
			_go_fuzz_dep_.CoverTab[56501]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:114
			// _ = "end of CoverTab[56501]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:114
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:114
		// _ = "end of CoverTab[56497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:114
		_go_fuzz_dep_.CoverTab[56498]++
															if !o.initialized {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:115
			_go_fuzz_dep_.CoverTab[56502]++
																initialized = false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:116
			// _ = "end of CoverTab[56502]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:117
			_go_fuzz_dep_.CoverTab[56503]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:117
			// _ = "end of CoverTab[56503]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:117
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:117
		// _ = "end of CoverTab[56498]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:117
		_go_fuzz_dep_.CoverTab[56499]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:118
		// _ = "end of CoverTab[56499]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:119
	// _ = "end of CoverTab[56491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:119
	_go_fuzz_dep_.CoverTab[56492]++
														out.n = len(b)
														out.initialized = initialized
														return out, err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:122
	// _ = "end of CoverTab[56492]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:123
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_messageset.go:123
var _ = _go_fuzz_dep_.CoverTab
