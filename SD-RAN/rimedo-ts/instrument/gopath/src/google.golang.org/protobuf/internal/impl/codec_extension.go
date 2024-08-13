// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:5
)

import (
	"sync"
	"sync/atomic"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type extensionFieldInfo struct {
	wiretag			uint64
	tagsize			int
	unmarshalNeedsValue	bool
	funcs			valueCoderFuncs
	validation		validationInfo
}

var legacyExtensionFieldInfoCache sync.Map	// map[protoreflect.ExtensionType]*extensionFieldInfo

func getExtensionFieldInfo(xt protoreflect.ExtensionType) *extensionFieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:26
	_go_fuzz_dep_.CoverTab[53727]++
														if xi, ok := xt.(*ExtensionInfo); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:27
		_go_fuzz_dep_.CoverTab[53729]++
															xi.lazyInit()
															return xi.info
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:29
		// _ = "end of CoverTab[53729]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:30
		_go_fuzz_dep_.CoverTab[53730]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:30
		// _ = "end of CoverTab[53730]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:30
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:30
	// _ = "end of CoverTab[53727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:30
	_go_fuzz_dep_.CoverTab[53728]++
														return legacyLoadExtensionFieldInfo(xt)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:31
	// _ = "end of CoverTab[53728]"
}

// legacyLoadExtensionFieldInfo dynamically loads a *ExtensionInfo for xt.
func legacyLoadExtensionFieldInfo(xt protoreflect.ExtensionType) *extensionFieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:35
	_go_fuzz_dep_.CoverTab[53731]++
														if xi, ok := legacyExtensionFieldInfoCache.Load(xt); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:36
		_go_fuzz_dep_.CoverTab[53734]++
															return xi.(*extensionFieldInfo)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:37
		// _ = "end of CoverTab[53734]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:38
		_go_fuzz_dep_.CoverTab[53735]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:38
		// _ = "end of CoverTab[53735]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:38
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:38
	// _ = "end of CoverTab[53731]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:38
	_go_fuzz_dep_.CoverTab[53732]++
														e := makeExtensionFieldInfo(xt.TypeDescriptor())
														if e, ok := legacyMessageTypeCache.LoadOrStore(xt, e); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:40
		_go_fuzz_dep_.CoverTab[53736]++
															return e.(*extensionFieldInfo)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:41
		// _ = "end of CoverTab[53736]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:42
		_go_fuzz_dep_.CoverTab[53737]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:42
		// _ = "end of CoverTab[53737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:42
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:42
	// _ = "end of CoverTab[53732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:42
	_go_fuzz_dep_.CoverTab[53733]++
														return e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:43
	// _ = "end of CoverTab[53733]"
}

func makeExtensionFieldInfo(xd protoreflect.ExtensionDescriptor) *extensionFieldInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:46
	_go_fuzz_dep_.CoverTab[53738]++
														var wiretag uint64
														if !xd.IsPacked() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:48
		_go_fuzz_dep_.CoverTab[53741]++
															wiretag = protowire.EncodeTag(xd.Number(), wireTypes[xd.Kind()])
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:49
		// _ = "end of CoverTab[53741]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:50
		_go_fuzz_dep_.CoverTab[53742]++
															wiretag = protowire.EncodeTag(xd.Number(), protowire.BytesType)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:51
		// _ = "end of CoverTab[53742]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:52
	// _ = "end of CoverTab[53738]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:52
	_go_fuzz_dep_.CoverTab[53739]++
														e := &extensionFieldInfo{
		wiretag:	wiretag,
		tagsize:	protowire.SizeVarint(wiretag),
		funcs:		encoderFuncsForValue(xd),
	}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:61
	switch xd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind, protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:62
		_go_fuzz_dep_.CoverTab[53743]++
															e.unmarshalNeedsValue = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:63
		// _ = "end of CoverTab[53743]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:64
		_go_fuzz_dep_.CoverTab[53744]++
															if xd.Cardinality() == protoreflect.Repeated {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:65
			_go_fuzz_dep_.CoverTab[53745]++
																e.unmarshalNeedsValue = true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:66
			// _ = "end of CoverTab[53745]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:67
			_go_fuzz_dep_.CoverTab[53746]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:67
			// _ = "end of CoverTab[53746]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:67
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:67
		// _ = "end of CoverTab[53744]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:68
	// _ = "end of CoverTab[53739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:68
	_go_fuzz_dep_.CoverTab[53740]++
														return e
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:69
	// _ = "end of CoverTab[53740]"
}

type lazyExtensionValue struct {
	atomicOnce	uint32	// atomically set if value is valid
	mu		sync.Mutex
	xi		*extensionFieldInfo
	value		protoreflect.Value
	b		[]byte
	fn		func() protoreflect.Value
}

type ExtensionField struct {
	typ	protoreflect.ExtensionType

	// value is either the value of GetValue,
	// or a *lazyExtensionValue that then returns the value of GetValue.
	value	protoreflect.Value
	lazy	*lazyExtensionValue
}

func (f *ExtensionField) appendLazyBytes(xt protoreflect.ExtensionType, xi *extensionFieldInfo, num protowire.Number, wtyp protowire.Type, b []byte) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:90
	_go_fuzz_dep_.CoverTab[53747]++
														if f.lazy == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:91
		_go_fuzz_dep_.CoverTab[53749]++
															f.lazy = &lazyExtensionValue{xi: xi}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:92
		// _ = "end of CoverTab[53749]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:93
		_go_fuzz_dep_.CoverTab[53750]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:93
		// _ = "end of CoverTab[53750]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:93
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:93
	// _ = "end of CoverTab[53747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:93
	_go_fuzz_dep_.CoverTab[53748]++
														f.typ = xt
														f.lazy.xi = xi
														f.lazy.b = protowire.AppendTag(f.lazy.b, num, wtyp)
														f.lazy.b = append(f.lazy.b, b...)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:97
	// _ = "end of CoverTab[53748]"
}

func (f *ExtensionField) canLazy(xt protoreflect.ExtensionType) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:100
	_go_fuzz_dep_.CoverTab[53751]++
														if f.typ == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:101
		_go_fuzz_dep_.CoverTab[53754]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:102
		// _ = "end of CoverTab[53754]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:103
		_go_fuzz_dep_.CoverTab[53755]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:103
		// _ = "end of CoverTab[53755]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:103
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:103
	// _ = "end of CoverTab[53751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:103
	_go_fuzz_dep_.CoverTab[53752]++
														if f.typ == xt && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		_go_fuzz_dep_.CoverTab[53756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		return f.lazy != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		// _ = "end of CoverTab[53756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		_go_fuzz_dep_.CoverTab[53757]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		return atomic.LoadUint32(&f.lazy.atomicOnce) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		// _ = "end of CoverTab[53757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:104
		_go_fuzz_dep_.CoverTab[53758]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:105
		// _ = "end of CoverTab[53758]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:106
		_go_fuzz_dep_.CoverTab[53759]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:106
		// _ = "end of CoverTab[53759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:106
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:106
	// _ = "end of CoverTab[53752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:106
	_go_fuzz_dep_.CoverTab[53753]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:107
	// _ = "end of CoverTab[53753]"
}

func (f *ExtensionField) lazyInit() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:110
	_go_fuzz_dep_.CoverTab[53760]++
														f.lazy.mu.Lock()
														defer f.lazy.mu.Unlock()
														if atomic.LoadUint32(&f.lazy.atomicOnce) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:113
		_go_fuzz_dep_.CoverTab[53763]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:114
		// _ = "end of CoverTab[53763]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:115
		_go_fuzz_dep_.CoverTab[53764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:115
		// _ = "end of CoverTab[53764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:115
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:115
	// _ = "end of CoverTab[53760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:115
	_go_fuzz_dep_.CoverTab[53761]++
														if f.lazy.xi != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:116
		_go_fuzz_dep_.CoverTab[53765]++
															b := f.lazy.b
															val := f.typ.New()
															for len(b) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:119
			_go_fuzz_dep_.CoverTab[53767]++
																var tag uint64
																if b[0] < 0x80 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:121
				_go_fuzz_dep_.CoverTab[53770]++
																	tag = uint64(b[0])
																	b = b[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:123
				// _ = "end of CoverTab[53770]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
				_go_fuzz_dep_.CoverTab[53771]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
				if len(b) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
					_go_fuzz_dep_.CoverTab[53772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
					return b[1] < 128
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
					// _ = "end of CoverTab[53772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:124
					_go_fuzz_dep_.CoverTab[53773]++
																		tag = uint64(b[0]&0x7f) + uint64(b[1])<<7
																		b = b[2:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:126
					// _ = "end of CoverTab[53773]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:127
					_go_fuzz_dep_.CoverTab[53774]++
																		var n int
																		tag, n = protowire.ConsumeVarint(b)
																		if n < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:130
						_go_fuzz_dep_.CoverTab[53776]++
																			panic(errors.New("bad tag in lazy extension decoding"))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:131
						// _ = "end of CoverTab[53776]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:132
						_go_fuzz_dep_.CoverTab[53777]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:132
						// _ = "end of CoverTab[53777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:132
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:132
					// _ = "end of CoverTab[53774]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:132
					_go_fuzz_dep_.CoverTab[53775]++
																		b = b[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:133
					// _ = "end of CoverTab[53775]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:134
				// _ = "end of CoverTab[53771]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:134
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:134
			// _ = "end of CoverTab[53767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:134
			_go_fuzz_dep_.CoverTab[53768]++
																num := protowire.Number(tag >> 3)
																wtyp := protowire.Type(tag & 7)
																var out unmarshalOutput
																var err error
																val, out, err = f.lazy.xi.funcs.unmarshal(b, val, num, wtyp, lazyUnmarshalOptions)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:140
				_go_fuzz_dep_.CoverTab[53778]++
																	panic(errors.New("decode failure in lazy extension decoding: %v", err))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:141
				// _ = "end of CoverTab[53778]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:142
				_go_fuzz_dep_.CoverTab[53779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:142
				// _ = "end of CoverTab[53779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:142
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:142
			// _ = "end of CoverTab[53768]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:142
			_go_fuzz_dep_.CoverTab[53769]++
																b = b[out.n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:143
			// _ = "end of CoverTab[53769]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:144
		// _ = "end of CoverTab[53765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:144
		_go_fuzz_dep_.CoverTab[53766]++
															f.lazy.value = val
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:145
		// _ = "end of CoverTab[53766]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:146
		_go_fuzz_dep_.CoverTab[53780]++
															f.lazy.value = f.lazy.fn()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:147
		// _ = "end of CoverTab[53780]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:148
	// _ = "end of CoverTab[53761]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:148
	_go_fuzz_dep_.CoverTab[53762]++
														f.lazy.xi = nil
														f.lazy.fn = nil
														f.lazy.b = nil
														atomic.StoreUint32(&f.lazy.atomicOnce, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:152
	// _ = "end of CoverTab[53762]"
}

// Set sets the type and value of the extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:155
// This must not be called concurrently.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:157
func (f *ExtensionField) Set(t protoreflect.ExtensionType, v protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:157
	_go_fuzz_dep_.CoverTab[53781]++
														f.typ = t
														f.value = v
														f.lazy = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:160
	// _ = "end of CoverTab[53781]"
}

// SetLazy sets the type and a value that is to be lazily evaluated upon first use.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:163
// This must not be called concurrently.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:165
func (f *ExtensionField) SetLazy(t protoreflect.ExtensionType, fn func() protoreflect.Value) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:165
	_go_fuzz_dep_.CoverTab[53782]++
														f.typ = t
														f.lazy = &lazyExtensionValue{fn: fn}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:167
	// _ = "end of CoverTab[53782]"
}

// Value returns the value of the extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:170
// This may be called concurrently.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:172
func (f *ExtensionField) Value() protoreflect.Value {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:172
	_go_fuzz_dep_.CoverTab[53783]++
														if f.lazy != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:173
		_go_fuzz_dep_.CoverTab[53785]++
															if atomic.LoadUint32(&f.lazy.atomicOnce) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:174
			_go_fuzz_dep_.CoverTab[53787]++
																f.lazyInit()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:175
			// _ = "end of CoverTab[53787]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:176
			_go_fuzz_dep_.CoverTab[53788]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:176
			// _ = "end of CoverTab[53788]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:176
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:176
		// _ = "end of CoverTab[53785]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:176
		_go_fuzz_dep_.CoverTab[53786]++
															return f.lazy.value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:177
		// _ = "end of CoverTab[53786]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:178
		_go_fuzz_dep_.CoverTab[53789]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:178
		// _ = "end of CoverTab[53789]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:178
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:178
	// _ = "end of CoverTab[53783]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:178
	_go_fuzz_dep_.CoverTab[53784]++
														return f.value
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:179
	// _ = "end of CoverTab[53784]"
}

// Type returns the type of the extension field.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:182
// This may be called concurrently.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:184
func (f ExtensionField) Type() protoreflect.ExtensionType {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:184
	_go_fuzz_dep_.CoverTab[53790]++
														return f.typ
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:185
	// _ = "end of CoverTab[53790]"
}

// IsSet returns whether the extension field is set.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:188
// This may be called concurrently.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:190
func (f ExtensionField) IsSet() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:190
	_go_fuzz_dep_.CoverTab[53791]++
														return f.typ != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:191
	// _ = "end of CoverTab[53791]"
}

// IsLazy reports whether a field is lazily encoded.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:194
// It is exported for testing.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:196
func IsLazy(m protoreflect.Message, fd protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:196
	_go_fuzz_dep_.CoverTab[53792]++
														var mi *MessageInfo
														var p pointer
														switch m := m.(type) {
	case *messageState:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:200
		_go_fuzz_dep_.CoverTab[53797]++
															mi = m.messageInfo()
															p = m.pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:202
		// _ = "end of CoverTab[53797]"
	case *messageReflectWrapper:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:203
		_go_fuzz_dep_.CoverTab[53798]++
															mi = m.messageInfo()
															p = m.pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:205
		// _ = "end of CoverTab[53798]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:206
		_go_fuzz_dep_.CoverTab[53799]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:207
		// _ = "end of CoverTab[53799]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:208
	// _ = "end of CoverTab[53792]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:208
	_go_fuzz_dep_.CoverTab[53793]++
														xd, ok := fd.(protoreflect.ExtensionTypeDescriptor)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:210
		_go_fuzz_dep_.CoverTab[53800]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:211
		// _ = "end of CoverTab[53800]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:212
		_go_fuzz_dep_.CoverTab[53801]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:212
		// _ = "end of CoverTab[53801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:212
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:212
	// _ = "end of CoverTab[53793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:212
	_go_fuzz_dep_.CoverTab[53794]++
														xt := xd.Type()
														ext := mi.extensionMap(p)
														if ext == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:215
		_go_fuzz_dep_.CoverTab[53802]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:216
		// _ = "end of CoverTab[53802]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:217
		_go_fuzz_dep_.CoverTab[53803]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:217
		// _ = "end of CoverTab[53803]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:217
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:217
	// _ = "end of CoverTab[53794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:217
	_go_fuzz_dep_.CoverTab[53795]++
														f, ok := (*ext)[int32(fd.Number())]
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:219
		_go_fuzz_dep_.CoverTab[53804]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:220
		// _ = "end of CoverTab[53804]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:221
		_go_fuzz_dep_.CoverTab[53805]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:221
		// _ = "end of CoverTab[53805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:221
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:221
	// _ = "end of CoverTab[53795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:221
	_go_fuzz_dep_.CoverTab[53796]++
														return f.typ == xt && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
		_go_fuzz_dep_.CoverTab[53806]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
		return f.lazy != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
		// _ = "end of CoverTab[53806]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
		_go_fuzz_dep_.CoverTab[53807]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
		return atomic.LoadUint32(&f.lazy.atomicOnce) == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
		// _ = "end of CoverTab[53807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:222
	// _ = "end of CoverTab[53796]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:223
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/codec_extension.go:223
var _ = _go_fuzz_dep_.CoverTab
