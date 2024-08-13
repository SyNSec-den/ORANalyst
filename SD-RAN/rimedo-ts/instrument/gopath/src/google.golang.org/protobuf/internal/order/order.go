// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
package order

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:5
)

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// FieldOrder specifies the ordering to visit message fields.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:11
// It is a function that reports whether x is ordered before y.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:13
type FieldOrder func(x, y protoreflect.FieldDescriptor) bool

var (
	// AnyFieldOrder specifies no specific field ordering.
	AnyFieldOrder	FieldOrder	= nil

	// LegacyFieldOrder sorts fields in the same ordering as emitted by
	// wire serialization in the github.com/golang/protobuf implementation.
	LegacyFieldOrder	FieldOrder	= func(x, y protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:21
		_go_fuzz_dep_.CoverTab[50115]++
														ox, oy := x.ContainingOneof(), y.ContainingOneof()
														inOneof := func(od protoreflect.OneofDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:23
			_go_fuzz_dep_.CoverTab[50120]++
															return od != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:24
				_go_fuzz_dep_.CoverTab[50121]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:24
				return !od.IsSynthetic()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:24
				// _ = "end of CoverTab[50121]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:24
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:24
			// _ = "end of CoverTab[50120]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:25
		// _ = "end of CoverTab[50115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:25
		_go_fuzz_dep_.CoverTab[50116]++

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:28
		if x.IsExtension() != y.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:28
			_go_fuzz_dep_.CoverTab[50122]++
															return x.IsExtension() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:29
				_go_fuzz_dep_.CoverTab[50123]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:29
				return !y.IsExtension()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:29
				// _ = "end of CoverTab[50123]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:29
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:29
			// _ = "end of CoverTab[50122]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:30
			_go_fuzz_dep_.CoverTab[50124]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:30
			// _ = "end of CoverTab[50124]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:30
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:30
		// _ = "end of CoverTab[50116]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:30
		_go_fuzz_dep_.CoverTab[50117]++

														if inOneof(ox) != inOneof(oy) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:32
			_go_fuzz_dep_.CoverTab[50125]++
															return !inOneof(ox) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:33
				_go_fuzz_dep_.CoverTab[50126]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:33
				return inOneof(oy)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:33
				// _ = "end of CoverTab[50126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:33
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:33
			// _ = "end of CoverTab[50125]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:34
			_go_fuzz_dep_.CoverTab[50127]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:34
			// _ = "end of CoverTab[50127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:34
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:34
		// _ = "end of CoverTab[50117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:34
		_go_fuzz_dep_.CoverTab[50118]++

														if ox != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			_go_fuzz_dep_.CoverTab[50128]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			return oy != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			// _ = "end of CoverTab[50128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			_go_fuzz_dep_.CoverTab[50129]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			return ox != oy
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			// _ = "end of CoverTab[50129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:36
			_go_fuzz_dep_.CoverTab[50130]++
															return ox.Index() < oy.Index()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:37
			// _ = "end of CoverTab[50130]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:38
			_go_fuzz_dep_.CoverTab[50131]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:38
			// _ = "end of CoverTab[50131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:38
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:38
		// _ = "end of CoverTab[50118]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:38
		_go_fuzz_dep_.CoverTab[50119]++

														return x.Number() < y.Number()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:40
		// _ = "end of CoverTab[50119]"
	}

	// NumberFieldOrder sorts fields by their field number.
	NumberFieldOrder	FieldOrder	= func(x, y protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:44
		_go_fuzz_dep_.CoverTab[50132]++
														return x.Number() < y.Number()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:45
		// _ = "end of CoverTab[50132]"
	}

	// IndexNameFieldOrder sorts non-extension fields before extension fields.
	// Non-extensions are sorted according to their declaration index.
	// Extensions are sorted according to their full name.
	IndexNameFieldOrder	FieldOrder	= func(x, y protoreflect.FieldDescriptor) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:51
		_go_fuzz_dep_.CoverTab[50133]++

														if x.IsExtension() != y.IsExtension() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:53
			_go_fuzz_dep_.CoverTab[50136]++
															return !x.IsExtension() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:54
				_go_fuzz_dep_.CoverTab[50137]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:54
				return y.IsExtension()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:54
				// _ = "end of CoverTab[50137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:54
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:54
			// _ = "end of CoverTab[50136]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:55
			_go_fuzz_dep_.CoverTab[50138]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:55
			// _ = "end of CoverTab[50138]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:55
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:55
		// _ = "end of CoverTab[50133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:55
		_go_fuzz_dep_.CoverTab[50134]++

														if x.IsExtension() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:57
			_go_fuzz_dep_.CoverTab[50139]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:57
			return y.IsExtension()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:57
			// _ = "end of CoverTab[50139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:57
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:57
			_go_fuzz_dep_.CoverTab[50140]++
															return x.FullName() < y.FullName()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:58
			// _ = "end of CoverTab[50140]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:59
			_go_fuzz_dep_.CoverTab[50141]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:59
			// _ = "end of CoverTab[50141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:59
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:59
		// _ = "end of CoverTab[50134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:59
		_go_fuzz_dep_.CoverTab[50135]++

														return x.Index() < y.Index()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:61
		// _ = "end of CoverTab[50135]"
	}
)

// KeyOrder specifies the ordering to visit map entries.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:65
// It is a function that reports whether x is ordered before y.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:67
type KeyOrder func(x, y protoreflect.MapKey) bool

var (
	// AnyKeyOrder specifies no specific key ordering.
	AnyKeyOrder	KeyOrder	= nil

	// GenericKeyOrder sorts false before true, numeric keys in ascending order,
	// and strings in lexicographical ordering according to UTF-8 codepoints.
	GenericKeyOrder	KeyOrder	= func(x, y protoreflect.MapKey) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:75
		_go_fuzz_dep_.CoverTab[50142]++
														switch x.Interface().(type) {
		case bool:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:77
			_go_fuzz_dep_.CoverTab[50143]++
															return !x.Bool() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:78
				_go_fuzz_dep_.CoverTab[50148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:78
				return y.Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:78
				// _ = "end of CoverTab[50148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:78
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:78
			// _ = "end of CoverTab[50143]"
		case int32, int64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:79
			_go_fuzz_dep_.CoverTab[50144]++
															return x.Int() < y.Int()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:80
			// _ = "end of CoverTab[50144]"
		case uint32, uint64:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:81
			_go_fuzz_dep_.CoverTab[50145]++
															return x.Uint() < y.Uint()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:82
			// _ = "end of CoverTab[50145]"
		case string:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:83
			_go_fuzz_dep_.CoverTab[50146]++
															return x.String() < y.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:84
			// _ = "end of CoverTab[50146]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:85
			_go_fuzz_dep_.CoverTab[50147]++
															panic("invalid map key type")
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:86
			// _ = "end of CoverTab[50147]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:87
		// _ = "end of CoverTab[50142]"
	}
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:89
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/internal/order/order.go:89
var _ = _go_fuzz_dep_.CoverTab
