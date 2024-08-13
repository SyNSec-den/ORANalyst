// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:5
)

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

type mergeOptions struct{}

func (o mergeOptions) Merge(dst, src proto.Message) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:18
	_go_fuzz_dep_.CoverTab[57875]++
												proto.Merge(dst, src)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:19
	// _ = "end of CoverTab[57875]"
}

// merge is protoreflect.Methods.Merge.
func (mi *MessageInfo) merge(in protoiface.MergeInput) protoiface.MergeOutput {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:23
	_go_fuzz_dep_.CoverTab[57876]++
												dp, ok := mi.getPointer(in.Destination)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:25
		_go_fuzz_dep_.CoverTab[57879]++
													return protoiface.MergeOutput{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:26
		// _ = "end of CoverTab[57879]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:27
		_go_fuzz_dep_.CoverTab[57880]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:27
		// _ = "end of CoverTab[57880]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:27
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:27
	// _ = "end of CoverTab[57876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:27
	_go_fuzz_dep_.CoverTab[57877]++
												sp, ok := mi.getPointer(in.Source)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:29
		_go_fuzz_dep_.CoverTab[57881]++
													return protoiface.MergeOutput{}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:30
		// _ = "end of CoverTab[57881]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:31
		_go_fuzz_dep_.CoverTab[57882]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:31
		// _ = "end of CoverTab[57882]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:31
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:31
	// _ = "end of CoverTab[57877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:31
	_go_fuzz_dep_.CoverTab[57878]++
												mi.mergePointer(dp, sp, mergeOptions{})
												return protoiface.MergeOutput{Flags: protoiface.MergeComplete}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:33
	// _ = "end of CoverTab[57878]"
}

func (mi *MessageInfo) mergePointer(dst, src pointer, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:36
	_go_fuzz_dep_.CoverTab[57883]++
												mi.init()
												if dst.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:38
		_go_fuzz_dep_.CoverTab[57888]++
													panic(fmt.Sprintf("invalid value: merging into nil message"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:39
		// _ = "end of CoverTab[57888]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:40
		_go_fuzz_dep_.CoverTab[57889]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:40
		// _ = "end of CoverTab[57889]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:40
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:40
	// _ = "end of CoverTab[57883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:40
	_go_fuzz_dep_.CoverTab[57884]++
												if src.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:41
		_go_fuzz_dep_.CoverTab[57890]++
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:42
		// _ = "end of CoverTab[57890]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:43
		_go_fuzz_dep_.CoverTab[57891]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:43
		// _ = "end of CoverTab[57891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:43
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:43
	// _ = "end of CoverTab[57884]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:43
	_go_fuzz_dep_.CoverTab[57885]++
												for _, f := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:44
		_go_fuzz_dep_.CoverTab[57892]++
													if f.funcs.merge == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:45
			_go_fuzz_dep_.CoverTab[57895]++
														continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:46
			// _ = "end of CoverTab[57895]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:47
			_go_fuzz_dep_.CoverTab[57896]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:47
			// _ = "end of CoverTab[57896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:47
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:47
		// _ = "end of CoverTab[57892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:47
		_go_fuzz_dep_.CoverTab[57893]++
													sfptr := src.Apply(f.offset)
													if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:49
			_go_fuzz_dep_.CoverTab[57897]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:49
			return sfptr.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:49
			// _ = "end of CoverTab[57897]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:49
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:49
			_go_fuzz_dep_.CoverTab[57898]++
														continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:50
			// _ = "end of CoverTab[57898]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:51
			_go_fuzz_dep_.CoverTab[57899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:51
			// _ = "end of CoverTab[57899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:51
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:51
		// _ = "end of CoverTab[57893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:51
		_go_fuzz_dep_.CoverTab[57894]++
													f.funcs.merge(dst.Apply(f.offset), sfptr, f, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:52
		// _ = "end of CoverTab[57894]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:53
	// _ = "end of CoverTab[57885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:53
	_go_fuzz_dep_.CoverTab[57886]++
												if mi.extensionOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:54
		_go_fuzz_dep_.CoverTab[57900]++
													sext := src.Apply(mi.extensionOffset).Extensions()
													dext := dst.Apply(mi.extensionOffset).Extensions()
													if *dext == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:57
			_go_fuzz_dep_.CoverTab[57902]++
														*dext = make(map[int32]ExtensionField)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:58
			// _ = "end of CoverTab[57902]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:59
			_go_fuzz_dep_.CoverTab[57903]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:59
			// _ = "end of CoverTab[57903]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:59
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:59
		// _ = "end of CoverTab[57900]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:59
		_go_fuzz_dep_.CoverTab[57901]++
													for num, sx := range *sext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:60
			_go_fuzz_dep_.CoverTab[57904]++
														xt := sx.Type()
														xi := getExtensionFieldInfo(xt)
														if xi.funcs.merge == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:63
				_go_fuzz_dep_.CoverTab[57908]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:64
				// _ = "end of CoverTab[57908]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:65
				_go_fuzz_dep_.CoverTab[57909]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:65
				// _ = "end of CoverTab[57909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:65
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:65
			// _ = "end of CoverTab[57904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:65
			_go_fuzz_dep_.CoverTab[57905]++
														dx := (*dext)[num]
														var dv protoreflect.Value
														if dx.Type() == sx.Type() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:68
				_go_fuzz_dep_.CoverTab[57910]++
															dv = dx.Value()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:69
				// _ = "end of CoverTab[57910]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:70
				_go_fuzz_dep_.CoverTab[57911]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:70
				// _ = "end of CoverTab[57911]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:70
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:70
			// _ = "end of CoverTab[57905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:70
			_go_fuzz_dep_.CoverTab[57906]++
														if !dv.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:71
				_go_fuzz_dep_.CoverTab[57912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:71
				return xi.unmarshalNeedsValue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:71
				// _ = "end of CoverTab[57912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:71
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:71
				_go_fuzz_dep_.CoverTab[57913]++
															dv = xt.New()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:72
				// _ = "end of CoverTab[57913]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:73
				_go_fuzz_dep_.CoverTab[57914]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:73
				// _ = "end of CoverTab[57914]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:73
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:73
			// _ = "end of CoverTab[57906]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:73
			_go_fuzz_dep_.CoverTab[57907]++
														dv = xi.funcs.merge(dv, sx.Value(), opts)
														dx.Set(sx.Type(), dv)
														(*dext)[num] = dx
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:76
			// _ = "end of CoverTab[57907]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:77
		// _ = "end of CoverTab[57901]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:78
		_go_fuzz_dep_.CoverTab[57915]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:78
		// _ = "end of CoverTab[57915]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:78
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:78
	// _ = "end of CoverTab[57886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:78
	_go_fuzz_dep_.CoverTab[57887]++
												if mi.unknownOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:79
		_go_fuzz_dep_.CoverTab[57916]++
													su := mi.getUnknownBytes(src)
													if su != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:81
			_go_fuzz_dep_.CoverTab[57917]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:81
			return len(*su) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:81
			// _ = "end of CoverTab[57917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:81
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:81
			_go_fuzz_dep_.CoverTab[57918]++
														du := mi.mutableUnknownBytes(dst)
														*du = append(*du, *su...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:83
			// _ = "end of CoverTab[57918]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:84
			_go_fuzz_dep_.CoverTab[57919]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:84
			// _ = "end of CoverTab[57919]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:84
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:84
		// _ = "end of CoverTab[57916]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:85
		_go_fuzz_dep_.CoverTab[57920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:85
		// _ = "end of CoverTab[57920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:85
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:85
	// _ = "end of CoverTab[57887]"
}

func mergeScalarValue(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:88
	_go_fuzz_dep_.CoverTab[57921]++
												return src
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:89
	// _ = "end of CoverTab[57921]"
}

func mergeBytesValue(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:92
	_go_fuzz_dep_.CoverTab[57922]++
												return protoreflect.ValueOfBytes(append(emptyBuf[:], src.Bytes()...))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:93
	// _ = "end of CoverTab[57922]"
}

func mergeListValue(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:96
	_go_fuzz_dep_.CoverTab[57923]++
												dstl := dst.List()
												srcl := src.List()
												for i, llen := 0, srcl.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:99
			_go_fuzz_dep_.CoverTab[57925]++
														dstl.Append(srcl.Get(i))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:100
		// _ = "end of CoverTab[57925]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:101
	// _ = "end of CoverTab[57923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:101
	_go_fuzz_dep_.CoverTab[57924]++
													return dst
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:102
	// _ = "end of CoverTab[57924]"
}

func mergeBytesListValue(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:105
	_go_fuzz_dep_.CoverTab[57926]++
													dstl := dst.List()
													srcl := src.List()
													for i, llen := 0, srcl.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:108
		_go_fuzz_dep_.CoverTab[57928]++
														sb := srcl.Get(i).Bytes()
														db := append(emptyBuf[:], sb...)
														dstl.Append(protoreflect.ValueOfBytes(db))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:111
		// _ = "end of CoverTab[57928]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:112
	// _ = "end of CoverTab[57926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:112
	_go_fuzz_dep_.CoverTab[57927]++
													return dst
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:113
	// _ = "end of CoverTab[57927]"
}

func mergeMessageListValue(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:116
	_go_fuzz_dep_.CoverTab[57929]++
													dstl := dst.List()
													srcl := src.List()
													for i, llen := 0, srcl.Len(); i < llen; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:119
		_go_fuzz_dep_.CoverTab[57931]++
														sm := srcl.Get(i).Message()
														dm := proto.Clone(sm.Interface()).ProtoReflect()
														dstl.Append(protoreflect.ValueOfMessage(dm))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:122
		// _ = "end of CoverTab[57931]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:123
	// _ = "end of CoverTab[57929]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:123
	_go_fuzz_dep_.CoverTab[57930]++
													return dst
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:124
	// _ = "end of CoverTab[57930]"
}

func mergeMessageValue(dst, src protoreflect.Value, opts mergeOptions) protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:127
	_go_fuzz_dep_.CoverTab[57932]++
													opts.Merge(dst.Message().Interface(), src.Message().Interface())
													return dst
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:129
	// _ = "end of CoverTab[57932]"
}

func mergeMessage(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:132
	_go_fuzz_dep_.CoverTab[57933]++
													if f.mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:133
		_go_fuzz_dep_.CoverTab[57934]++
														if dst.Elem().IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:134
			_go_fuzz_dep_.CoverTab[57936]++
															dst.SetPointer(pointerOfValue(reflect.New(f.mi.GoReflectType.Elem())))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:135
			// _ = "end of CoverTab[57936]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:136
			_go_fuzz_dep_.CoverTab[57937]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:136
			// _ = "end of CoverTab[57937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:136
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:136
		// _ = "end of CoverTab[57934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:136
		_go_fuzz_dep_.CoverTab[57935]++
														f.mi.mergePointer(dst.Elem(), src.Elem(), opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:137
		// _ = "end of CoverTab[57935]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:138
		_go_fuzz_dep_.CoverTab[57938]++
														dm := dst.AsValueOf(f.ft).Elem()
														sm := src.AsValueOf(f.ft).Elem()
														if dm.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:141
			_go_fuzz_dep_.CoverTab[57940]++
															dm.Set(reflect.New(f.ft.Elem()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:142
			// _ = "end of CoverTab[57940]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:143
			_go_fuzz_dep_.CoverTab[57941]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:143
			// _ = "end of CoverTab[57941]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:143
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:143
		// _ = "end of CoverTab[57938]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:143
		_go_fuzz_dep_.CoverTab[57939]++
														opts.Merge(asMessage(dm), asMessage(sm))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:144
		// _ = "end of CoverTab[57939]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:145
	// _ = "end of CoverTab[57933]"
}

func mergeMessageSlice(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:148
	_go_fuzz_dep_.CoverTab[57942]++
													for _, sp := range src.PointerSlice() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:149
		_go_fuzz_dep_.CoverTab[57943]++
														dm := reflect.New(f.ft.Elem().Elem())
														if f.mi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:151
			_go_fuzz_dep_.CoverTab[57945]++
															f.mi.mergePointer(pointerOfValue(dm), sp, opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:152
			// _ = "end of CoverTab[57945]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:153
			_go_fuzz_dep_.CoverTab[57946]++
															opts.Merge(asMessage(dm), asMessage(sp.AsValueOf(f.ft.Elem().Elem())))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:154
			// _ = "end of CoverTab[57946]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:155
		// _ = "end of CoverTab[57943]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:155
		_go_fuzz_dep_.CoverTab[57944]++
														dst.AppendPointerSlice(pointerOfValue(dm))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:156
		// _ = "end of CoverTab[57944]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:157
	// _ = "end of CoverTab[57942]"
}

func mergeBytes(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:160
	_go_fuzz_dep_.CoverTab[57947]++
													*dst.Bytes() = append(emptyBuf[:], *src.Bytes()...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:161
	// _ = "end of CoverTab[57947]"
}

func mergeBytesNoZero(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:164
	_go_fuzz_dep_.CoverTab[57948]++
													v := *src.Bytes()
													if len(v) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:166
		_go_fuzz_dep_.CoverTab[57949]++
														*dst.Bytes() = append(emptyBuf[:], v...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:167
		// _ = "end of CoverTab[57949]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:168
		_go_fuzz_dep_.CoverTab[57950]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:168
		// _ = "end of CoverTab[57950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:168
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:168
	// _ = "end of CoverTab[57948]"
}

func mergeBytesSlice(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:171
	_go_fuzz_dep_.CoverTab[57951]++
													ds := dst.BytesSlice()
													for _, v := range *src.BytesSlice() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:173
		_go_fuzz_dep_.CoverTab[57952]++
														*ds = append(*ds, append(emptyBuf[:], v...))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:174
		// _ = "end of CoverTab[57952]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:175
	// _ = "end of CoverTab[57951]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:176
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/merge.go:176
var _ = _go_fuzz_dep_.CoverTab
