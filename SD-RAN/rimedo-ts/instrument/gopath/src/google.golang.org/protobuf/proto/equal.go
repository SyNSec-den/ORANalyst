// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
package proto

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:5
)

import (
	"bytes"
	"math"
	"reflect"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Equal reports whether two messages are equal.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// If two messages marshal to the same bytes under deterministic serialization,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// then Equal is guaranteed to report true.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// Two messages are equal if they belong to the same message descriptor,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// have the same set of populated known and extension field values,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// and the same set of unknown fields values. If either of the top-level
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// messages are invalid, then Equal reports true only if both are invalid.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
//
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// Scalar values are compared with the equivalent of the == operator in Go,
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// except bytes values which are compared using bytes.Equal and
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// floating point values which specially treat NaNs as equal.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// Message values are compared by recursively calling Equal.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// Lists are equal if each element value is also equal.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// Maps are equal if they have the same set of keys, where the pair of values
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:16
// for each key is also equal.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:32
func Equal(x, y Message) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:32
	_go_fuzz_dep_.CoverTab[51333]++
											if x == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:33
		_go_fuzz_dep_.CoverTab[51337]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:33
		return y == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:33
		// _ = "end of CoverTab[51337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:33
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:33
		_go_fuzz_dep_.CoverTab[51338]++
												return x == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:34
			_go_fuzz_dep_.CoverTab[51339]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:34
			return y == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:34
			// _ = "end of CoverTab[51339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:34
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:34
		// _ = "end of CoverTab[51338]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:35
		_go_fuzz_dep_.CoverTab[51340]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:35
		// _ = "end of CoverTab[51340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:35
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:35
	// _ = "end of CoverTab[51333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:35
	_go_fuzz_dep_.CoverTab[51334]++
											if reflect.TypeOf(x).Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:36
		_go_fuzz_dep_.CoverTab[51341]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:36
		return x == y
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:36
		// _ = "end of CoverTab[51341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:36
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:36
		_go_fuzz_dep_.CoverTab[51342]++

												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:38
		// _ = "end of CoverTab[51342]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:39
		_go_fuzz_dep_.CoverTab[51343]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:39
		// _ = "end of CoverTab[51343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:39
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:39
	// _ = "end of CoverTab[51334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:39
	_go_fuzz_dep_.CoverTab[51335]++
											mx := x.ProtoReflect()
											my := y.ProtoReflect()
											if mx.IsValid() != my.IsValid() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:42
		_go_fuzz_dep_.CoverTab[51344]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:43
		// _ = "end of CoverTab[51344]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:44
		_go_fuzz_dep_.CoverTab[51345]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:44
		// _ = "end of CoverTab[51345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:44
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:44
	// _ = "end of CoverTab[51335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:44
	_go_fuzz_dep_.CoverTab[51336]++
											return equalMessage(mx, my)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:45
	// _ = "end of CoverTab[51336]"
}

// equalMessage compares two messages.
func equalMessage(mx, my protoreflect.Message) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:49
	_go_fuzz_dep_.CoverTab[51346]++
											if mx.Descriptor() != my.Descriptor() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:50
		_go_fuzz_dep_.CoverTab[51352]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:51
		// _ = "end of CoverTab[51352]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:52
		_go_fuzz_dep_.CoverTab[51353]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:52
		// _ = "end of CoverTab[51353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:52
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:52
	// _ = "end of CoverTab[51346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:52
	_go_fuzz_dep_.CoverTab[51347]++

											nx := 0
											equal := true
											mx.Range(func(fd protoreflect.FieldDescriptor, vx protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:56
		_go_fuzz_dep_.CoverTab[51354]++
												nx++
												vy := my.Get(fd)
												equal = my.Has(fd) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:59
			_go_fuzz_dep_.CoverTab[51355]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:59
			return equalField(fd, vx, vy)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:59
			// _ = "end of CoverTab[51355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:59
		}()
												return equal
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:60
		// _ = "end of CoverTab[51354]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:61
	// _ = "end of CoverTab[51347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:61
	_go_fuzz_dep_.CoverTab[51348]++
											if !equal {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:62
		_go_fuzz_dep_.CoverTab[51356]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:63
		// _ = "end of CoverTab[51356]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:64
		_go_fuzz_dep_.CoverTab[51357]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:64
		// _ = "end of CoverTab[51357]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:64
	// _ = "end of CoverTab[51348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:64
	_go_fuzz_dep_.CoverTab[51349]++
											ny := 0
											my.Range(func(fd protoreflect.FieldDescriptor, vx protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:66
		_go_fuzz_dep_.CoverTab[51358]++
												ny++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:68
		// _ = "end of CoverTab[51358]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:69
	// _ = "end of CoverTab[51349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:69
	_go_fuzz_dep_.CoverTab[51350]++
											if nx != ny {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:70
		_go_fuzz_dep_.CoverTab[51359]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:71
		// _ = "end of CoverTab[51359]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:72
		_go_fuzz_dep_.CoverTab[51360]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:72
		// _ = "end of CoverTab[51360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:72
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:72
	// _ = "end of CoverTab[51350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:72
	_go_fuzz_dep_.CoverTab[51351]++

											return equalUnknown(mx.GetUnknown(), my.GetUnknown())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:74
	// _ = "end of CoverTab[51351]"
}

// equalField compares two fields.
func equalField(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:78
	_go_fuzz_dep_.CoverTab[51361]++
											switch {
	case fd.IsList():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:80
		_go_fuzz_dep_.CoverTab[51362]++
												return equalList(fd, x.List(), y.List())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:81
		// _ = "end of CoverTab[51362]"
	case fd.IsMap():
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:82
		_go_fuzz_dep_.CoverTab[51363]++
												return equalMap(fd, x.Map(), y.Map())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:83
		// _ = "end of CoverTab[51363]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:84
		_go_fuzz_dep_.CoverTab[51364]++
												return equalValue(fd, x, y)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:85
		// _ = "end of CoverTab[51364]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:86
	// _ = "end of CoverTab[51361]"
}

// equalMap compares two maps.
func equalMap(fd protoreflect.FieldDescriptor, x, y protoreflect.Map) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:90
	_go_fuzz_dep_.CoverTab[51365]++
											if x.Len() != y.Len() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:91
		_go_fuzz_dep_.CoverTab[51368]++
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:92
		// _ = "end of CoverTab[51368]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:93
		_go_fuzz_dep_.CoverTab[51369]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:93
		// _ = "end of CoverTab[51369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:93
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:93
	// _ = "end of CoverTab[51365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:93
	_go_fuzz_dep_.CoverTab[51366]++
											equal := true
											x.Range(func(k protoreflect.MapKey, vx protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:95
		_go_fuzz_dep_.CoverTab[51370]++
												vy := y.Get(k)
												equal = y.Has(k) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:97
			_go_fuzz_dep_.CoverTab[51371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:97
			return equalValue(fd.MapValue(), vx, vy)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:97
			// _ = "end of CoverTab[51371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:97
		}()
												return equal
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:98
		// _ = "end of CoverTab[51370]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:99
		// _ = "end of CoverTab[51366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:99
		_go_fuzz_dep_.CoverTab[51367]++
												return equal
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:100
	// _ = "end of CoverTab[51367]"
}

// equalList compares two lists.
func equalList(fd protoreflect.FieldDescriptor, x, y protoreflect.List) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:104
	_go_fuzz_dep_.CoverTab[51372]++
												if x.Len() != y.Len() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:105
		_go_fuzz_dep_.CoverTab[51375]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:106
		// _ = "end of CoverTab[51375]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:107
		_go_fuzz_dep_.CoverTab[51376]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:107
		// _ = "end of CoverTab[51376]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:107
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:107
	// _ = "end of CoverTab[51372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:107
	_go_fuzz_dep_.CoverTab[51373]++
												for i := x.Len() - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:108
		_go_fuzz_dep_.CoverTab[51377]++
													if !equalValue(fd, x.Get(i), y.Get(i)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:109
			_go_fuzz_dep_.CoverTab[51378]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:110
			// _ = "end of CoverTab[51378]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:111
			_go_fuzz_dep_.CoverTab[51379]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:111
			// _ = "end of CoverTab[51379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:111
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:111
		// _ = "end of CoverTab[51377]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:112
	// _ = "end of CoverTab[51373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:112
	_go_fuzz_dep_.CoverTab[51374]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:113
	// _ = "end of CoverTab[51374]"
}

// equalValue compares two singular values.
func equalValue(fd protoreflect.FieldDescriptor, x, y protoreflect.Value) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:117
	_go_fuzz_dep_.CoverTab[51380]++
												switch fd.Kind() {
	case protoreflect.BoolKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:119
		_go_fuzz_dep_.CoverTab[51381]++
													return x.Bool() == y.Bool()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:120
		// _ = "end of CoverTab[51381]"
	case protoreflect.EnumKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:121
		_go_fuzz_dep_.CoverTab[51382]++
													return x.Enum() == y.Enum()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:122
		// _ = "end of CoverTab[51382]"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:125
		_go_fuzz_dep_.CoverTab[51383]++
													return x.Int() == y.Int()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:126
		// _ = "end of CoverTab[51383]"
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:128
		_go_fuzz_dep_.CoverTab[51384]++
													return x.Uint() == y.Uint()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:129
		// _ = "end of CoverTab[51384]"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:130
		_go_fuzz_dep_.CoverTab[51385]++
													fx := x.Float()
													fy := y.Float()
													if math.IsNaN(fx) || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:133
			_go_fuzz_dep_.CoverTab[51391]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:133
			return math.IsNaN(fy)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:133
			// _ = "end of CoverTab[51391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:133
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:133
			_go_fuzz_dep_.CoverTab[51392]++
														return math.IsNaN(fx) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:134
				_go_fuzz_dep_.CoverTab[51393]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:134
				return math.IsNaN(fy)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:134
				// _ = "end of CoverTab[51393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:134
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:134
			// _ = "end of CoverTab[51392]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:135
			_go_fuzz_dep_.CoverTab[51394]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:135
			// _ = "end of CoverTab[51394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:135
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:135
		// _ = "end of CoverTab[51385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:135
		_go_fuzz_dep_.CoverTab[51386]++
													return fx == fy
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:136
		// _ = "end of CoverTab[51386]"
	case protoreflect.StringKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:137
		_go_fuzz_dep_.CoverTab[51387]++
													return x.String() == y.String()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:138
		// _ = "end of CoverTab[51387]"
	case protoreflect.BytesKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:139
		_go_fuzz_dep_.CoverTab[51388]++
													return bytes.Equal(x.Bytes(), y.Bytes())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:140
		// _ = "end of CoverTab[51388]"
	case protoreflect.MessageKind, protoreflect.GroupKind:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:141
		_go_fuzz_dep_.CoverTab[51389]++
													return equalMessage(x.Message(), y.Message())
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:142
		// _ = "end of CoverTab[51389]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:143
		_go_fuzz_dep_.CoverTab[51390]++
													return x.Interface() == y.Interface()
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:144
		// _ = "end of CoverTab[51390]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:145
	// _ = "end of CoverTab[51380]"
}

// equalUnknown compares unknown fields by direct comparison on the raw bytes
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:148
// of each individual field number.
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:150
func equalUnknown(x, y protoreflect.RawFields) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:150
	_go_fuzz_dep_.CoverTab[51395]++
												if len(x) != len(y) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:151
		_go_fuzz_dep_.CoverTab[51400]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:152
		// _ = "end of CoverTab[51400]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:153
		_go_fuzz_dep_.CoverTab[51401]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:153
		// _ = "end of CoverTab[51401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:153
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:153
	// _ = "end of CoverTab[51395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:153
	_go_fuzz_dep_.CoverTab[51396]++
												if bytes.Equal([]byte(x), []byte(y)) {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:154
		_go_fuzz_dep_.CoverTab[51402]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:155
		// _ = "end of CoverTab[51402]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:156
		_go_fuzz_dep_.CoverTab[51403]++
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:156
		// _ = "end of CoverTab[51403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:156
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:156
	// _ = "end of CoverTab[51396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:156
	_go_fuzz_dep_.CoverTab[51397]++

												mx := make(map[protoreflect.FieldNumber]protoreflect.RawFields)
												my := make(map[protoreflect.FieldNumber]protoreflect.RawFields)
												for len(x) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:160
		_go_fuzz_dep_.CoverTab[51404]++
													fnum, _, n := protowire.ConsumeField(x)
													mx[fnum] = append(mx[fnum], x[:n]...)
													x = x[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:163
		// _ = "end of CoverTab[51404]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:164
	// _ = "end of CoverTab[51397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:164
	_go_fuzz_dep_.CoverTab[51398]++
												for len(y) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:165
		_go_fuzz_dep_.CoverTab[51405]++
													fnum, _, n := protowire.ConsumeField(y)
													my[fnum] = append(my[fnum], y[:n]...)
													y = y[n:]
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:168
		// _ = "end of CoverTab[51405]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:169
	// _ = "end of CoverTab[51398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:169
	_go_fuzz_dep_.CoverTab[51399]++
												return reflect.DeepEqual(mx, my)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:170
	// _ = "end of CoverTab[51399]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:171
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/protobuf@v1.28.1/proto/equal.go:171
var _ = _go_fuzz_dep_.CoverTab
