// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
package impl

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:5
)

import (
	"sync"

	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

func (mi *MessageInfo) checkInitialized(in protoiface.CheckInitializedInput) (protoiface.CheckInitializedOutput, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:15
	_go_fuzz_dep_.CoverTab[53648]++
													var p pointer
													if ms, ok := in.Message.(*messageState); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:17
		_go_fuzz_dep_.CoverTab[53650]++
														p = ms.pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:18
		// _ = "end of CoverTab[53650]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:19
		_go_fuzz_dep_.CoverTab[53651]++
														p = in.Message.(*messageReflectWrapper).pointer()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:20
		// _ = "end of CoverTab[53651]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:21
	// _ = "end of CoverTab[53648]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:21
	_go_fuzz_dep_.CoverTab[53649]++
													return protoiface.CheckInitializedOutput{}, mi.checkInitializedPointer(p)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:22
	// _ = "end of CoverTab[53649]"
}

func (mi *MessageInfo) checkInitializedPointer(p pointer) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:25
	_go_fuzz_dep_.CoverTab[53652]++
													mi.init()
													if !mi.needsInitCheck {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:27
		_go_fuzz_dep_.CoverTab[53657]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:28
		// _ = "end of CoverTab[53657]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:29
		_go_fuzz_dep_.CoverTab[53658]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:29
		// _ = "end of CoverTab[53658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:29
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:29
	// _ = "end of CoverTab[53652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:29
	_go_fuzz_dep_.CoverTab[53653]++
													if p.IsNil() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:30
		_go_fuzz_dep_.CoverTab[53659]++
														for _, f := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:31
			_go_fuzz_dep_.CoverTab[53661]++
															if f.isRequired {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:32
				_go_fuzz_dep_.CoverTab[53662]++
																return errors.RequiredNotSet(string(mi.Desc.Fields().ByNumber(f.num).FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:33
				// _ = "end of CoverTab[53662]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:34
				_go_fuzz_dep_.CoverTab[53663]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:34
				// _ = "end of CoverTab[53663]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:34
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:34
			// _ = "end of CoverTab[53661]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:35
		// _ = "end of CoverTab[53659]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:35
		_go_fuzz_dep_.CoverTab[53660]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:36
		// _ = "end of CoverTab[53660]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:37
		_go_fuzz_dep_.CoverTab[53664]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:37
		// _ = "end of CoverTab[53664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:37
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:37
	// _ = "end of CoverTab[53653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:37
	_go_fuzz_dep_.CoverTab[53654]++
													if mi.extensionOffset.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:38
		_go_fuzz_dep_.CoverTab[53665]++
														e := p.Apply(mi.extensionOffset).Extensions()
														if err := mi.isInitExtensions(e); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:40
			_go_fuzz_dep_.CoverTab[53666]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:41
			// _ = "end of CoverTab[53666]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:42
			_go_fuzz_dep_.CoverTab[53667]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:42
			// _ = "end of CoverTab[53667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:42
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:42
		// _ = "end of CoverTab[53665]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:43
		_go_fuzz_dep_.CoverTab[53668]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:43
		// _ = "end of CoverTab[53668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:43
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:43
	// _ = "end of CoverTab[53654]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:43
	_go_fuzz_dep_.CoverTab[53655]++
													for _, f := range mi.orderedCoderFields {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:44
		_go_fuzz_dep_.CoverTab[53669]++
														if !f.isRequired && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:45
			_go_fuzz_dep_.CoverTab[53673]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:45
			return f.funcs.isInit == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:45
			// _ = "end of CoverTab[53673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:45
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:45
			_go_fuzz_dep_.CoverTab[53674]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:46
			// _ = "end of CoverTab[53674]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:47
			_go_fuzz_dep_.CoverTab[53675]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:47
			// _ = "end of CoverTab[53675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:47
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:47
		// _ = "end of CoverTab[53669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:47
		_go_fuzz_dep_.CoverTab[53670]++
														fptr := p.Apply(f.offset)
														if f.isPointer && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:49
			_go_fuzz_dep_.CoverTab[53676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:49
			return fptr.Elem().IsNil()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:49
			// _ = "end of CoverTab[53676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:49
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:49
			_go_fuzz_dep_.CoverTab[53677]++
															if f.isRequired {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:50
				_go_fuzz_dep_.CoverTab[53679]++
																return errors.RequiredNotSet(string(mi.Desc.Fields().ByNumber(f.num).FullName()))
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:51
				// _ = "end of CoverTab[53679]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:52
				_go_fuzz_dep_.CoverTab[53680]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:52
				// _ = "end of CoverTab[53680]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:52
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:52
			// _ = "end of CoverTab[53677]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:52
			_go_fuzz_dep_.CoverTab[53678]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:53
			// _ = "end of CoverTab[53678]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:54
			_go_fuzz_dep_.CoverTab[53681]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:54
			// _ = "end of CoverTab[53681]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:54
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:54
		// _ = "end of CoverTab[53670]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:54
		_go_fuzz_dep_.CoverTab[53671]++
														if f.funcs.isInit == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:55
			_go_fuzz_dep_.CoverTab[53682]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:56
			// _ = "end of CoverTab[53682]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:57
			_go_fuzz_dep_.CoverTab[53683]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:57
			// _ = "end of CoverTab[53683]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:57
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:57
		// _ = "end of CoverTab[53671]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:57
		_go_fuzz_dep_.CoverTab[53672]++
														if err := f.funcs.isInit(fptr, f); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:58
			_go_fuzz_dep_.CoverTab[53684]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:59
			// _ = "end of CoverTab[53684]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:60
			_go_fuzz_dep_.CoverTab[53685]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:60
			// _ = "end of CoverTab[53685]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:60
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:60
		// _ = "end of CoverTab[53672]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:61
	// _ = "end of CoverTab[53655]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:61
	_go_fuzz_dep_.CoverTab[53656]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:62
	// _ = "end of CoverTab[53656]"
}

func (mi *MessageInfo) isInitExtensions(ext *map[int32]ExtensionField) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:65
	_go_fuzz_dep_.CoverTab[53686]++
													if ext == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:66
		_go_fuzz_dep_.CoverTab[53689]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:67
		// _ = "end of CoverTab[53689]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:68
		_go_fuzz_dep_.CoverTab[53690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:68
		// _ = "end of CoverTab[53690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:68
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:68
	// _ = "end of CoverTab[53686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:68
	_go_fuzz_dep_.CoverTab[53687]++
													for _, x := range *ext {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:69
		_go_fuzz_dep_.CoverTab[53691]++
														ei := getExtensionFieldInfo(x.Type())
														if ei.funcs.isInit == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:71
			_go_fuzz_dep_.CoverTab[53694]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:72
			// _ = "end of CoverTab[53694]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:73
			_go_fuzz_dep_.CoverTab[53695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:73
			// _ = "end of CoverTab[53695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:73
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:73
		// _ = "end of CoverTab[53691]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:73
		_go_fuzz_dep_.CoverTab[53692]++
														v := x.Value()
														if !v.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:75
			_go_fuzz_dep_.CoverTab[53696]++
															continue
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:76
			// _ = "end of CoverTab[53696]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:77
			_go_fuzz_dep_.CoverTab[53697]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:77
			// _ = "end of CoverTab[53697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:77
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:77
		// _ = "end of CoverTab[53692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:77
		_go_fuzz_dep_.CoverTab[53693]++
														if err := ei.funcs.isInit(v); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:78
			_go_fuzz_dep_.CoverTab[53698]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:79
			// _ = "end of CoverTab[53698]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:80
			_go_fuzz_dep_.CoverTab[53699]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:80
			// _ = "end of CoverTab[53699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:80
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:80
		// _ = "end of CoverTab[53693]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:81
	// _ = "end of CoverTab[53687]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:81
	_go_fuzz_dep_.CoverTab[53688]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:82
	// _ = "end of CoverTab[53688]"
}

var (
	needsInitCheckMu	sync.Mutex
	needsInitCheckMap	sync.Map
)

// needsInitCheck reports whether a message needs to be checked for partial initialization.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:90
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:90
// It returns true if the message transitively includes any required or extension fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:93
func needsInitCheck(md protoreflect.MessageDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:93
	_go_fuzz_dep_.CoverTab[53700]++
													if v, ok := needsInitCheckMap.Load(md); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:94
		_go_fuzz_dep_.CoverTab[53702]++
														if has, ok := v.(bool); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:95
			_go_fuzz_dep_.CoverTab[53703]++
															return has
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:96
			// _ = "end of CoverTab[53703]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:97
			_go_fuzz_dep_.CoverTab[53704]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:97
			// _ = "end of CoverTab[53704]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:97
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:97
		// _ = "end of CoverTab[53702]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:98
		_go_fuzz_dep_.CoverTab[53705]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:98
		// _ = "end of CoverTab[53705]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:98
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:98
	// _ = "end of CoverTab[53700]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:98
	_go_fuzz_dep_.CoverTab[53701]++
													needsInitCheckMu.Lock()
													defer needsInitCheckMu.Unlock()
													return needsInitCheckLocked(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:101
	// _ = "end of CoverTab[53701]"
}

func needsInitCheckLocked(md protoreflect.MessageDescriptor) (has bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:104
	_go_fuzz_dep_.CoverTab[53706]++
													if v, ok := needsInitCheckMap.Load(md); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:105
		_go_fuzz_dep_.CoverTab[53712]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:116
		has, ok := v.(bool)
		return ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:117
			_go_fuzz_dep_.CoverTab[53713]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:117
			return has
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:117
			// _ = "end of CoverTab[53713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:117
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:117
		// _ = "end of CoverTab[53712]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:118
		_go_fuzz_dep_.CoverTab[53714]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:118
		// _ = "end of CoverTab[53714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:118
	// _ = "end of CoverTab[53706]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:118
	_go_fuzz_dep_.CoverTab[53707]++
													needsInitCheckMap.Store(md, struct{}{})
													defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:120
		_go_fuzz_dep_.CoverTab[53715]++
														needsInitCheckMap.Store(md, has)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:121
		// _ = "end of CoverTab[53715]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:122
	// _ = "end of CoverTab[53707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:122
	_go_fuzz_dep_.CoverTab[53708]++
													if md.RequiredNumbers().Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:123
		_go_fuzz_dep_.CoverTab[53716]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:124
		// _ = "end of CoverTab[53716]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:125
		_go_fuzz_dep_.CoverTab[53717]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:125
		// _ = "end of CoverTab[53717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:125
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:125
	// _ = "end of CoverTab[53708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:125
	_go_fuzz_dep_.CoverTab[53709]++
													if md.ExtensionRanges().Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:126
		_go_fuzz_dep_.CoverTab[53718]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:127
		// _ = "end of CoverTab[53718]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:128
		_go_fuzz_dep_.CoverTab[53719]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:128
		// _ = "end of CoverTab[53719]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:128
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:128
	// _ = "end of CoverTab[53709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:128
	_go_fuzz_dep_.CoverTab[53710]++
													for i := 0; i < md.Fields().Len(); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:129
		_go_fuzz_dep_.CoverTab[53720]++
														fd := md.Fields().Get(i)

														if fd.IsMap() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:132
			_go_fuzz_dep_.CoverTab[53722]++
															fd = fd.MapValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:133
			// _ = "end of CoverTab[53722]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:134
			_go_fuzz_dep_.CoverTab[53723]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:134
			// _ = "end of CoverTab[53723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:134
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:134
		// _ = "end of CoverTab[53720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:134
		_go_fuzz_dep_.CoverTab[53721]++
														fmd := fd.Message()
														if fmd != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:136
			_go_fuzz_dep_.CoverTab[53724]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:136
			return needsInitCheckLocked(fmd)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:136
			// _ = "end of CoverTab[53724]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:136
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:136
			_go_fuzz_dep_.CoverTab[53725]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:137
			// _ = "end of CoverTab[53725]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:138
			_go_fuzz_dep_.CoverTab[53726]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:138
			// _ = "end of CoverTab[53726]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:138
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:138
		// _ = "end of CoverTab[53721]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:139
	// _ = "end of CoverTab[53710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:139
	_go_fuzz_dep_.CoverTab[53711]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:140
	// _ = "end of CoverTab[53711]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:141
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/impl/checkinit.go:141
var _ = _go_fuzz_dep_.CoverTab
