//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
// Package attributes defines a generic key/value store used in various gRPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
// components.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
// Notice: This package is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:19
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
package attributes

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:26
)

// Attributes is an immutable struct for storing and retrieving generic
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:28
// key/value pairs.  Keys must be hashable, and users should define their own
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:28
// types for keys.  Values should not be modified after they are added to an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:28
// Attributes or if they were received from one.  If values implement 'Equal(o
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:28
// interface{}) bool', it will be called by (*Attributes).Equal to determine
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:28
// whether two values with the same key should be considered equal.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:34
type Attributes struct {
	m map[interface{}]interface{}
}

// New returns a new Attributes containing the key/value pair.
func New(key, value interface{}) *Attributes {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:39
	_go_fuzz_dep_.CoverTab[62423]++
												return &Attributes{m: map[interface{}]interface{}{key: value}}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:40
	// _ = "end of CoverTab[62423]"
}

// WithValue returns a new Attributes containing the previous keys and values
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:43
// and the new key/value pair.  If the same key appears multiple times, the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:43
// last value overwrites all previous values for that key.  To remove an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:43
// existing key, use a nil value.  value should not be modified later.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:47
func (a *Attributes) WithValue(key, value interface{}) *Attributes {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:47
	_go_fuzz_dep_.CoverTab[62424]++
												if a == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:48
		_go_fuzz_dep_.CoverTab[62427]++
													return New(key, value)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:49
		// _ = "end of CoverTab[62427]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:50
		_go_fuzz_dep_.CoverTab[62428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:50
		// _ = "end of CoverTab[62428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:50
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:50
	// _ = "end of CoverTab[62424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:50
	_go_fuzz_dep_.CoverTab[62425]++
												n := &Attributes{m: make(map[interface{}]interface{}, len(a.m)+1)}
												for k, v := range a.m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:52
		_go_fuzz_dep_.CoverTab[62429]++
													n.m[k] = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:53
		// _ = "end of CoverTab[62429]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:54
	// _ = "end of CoverTab[62425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:54
	_go_fuzz_dep_.CoverTab[62426]++
												n.m[key] = value
												return n
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:56
	// _ = "end of CoverTab[62426]"
}

// Value returns the value associated with these attributes for key, or nil if
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:59
// no value is associated with key.  The returned value should not be modified.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:61
func (a *Attributes) Value(key interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:61
	_go_fuzz_dep_.CoverTab[62430]++
												if a == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:62
		_go_fuzz_dep_.CoverTab[62432]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:63
		// _ = "end of CoverTab[62432]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:64
		_go_fuzz_dep_.CoverTab[62433]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:64
		// _ = "end of CoverTab[62433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:64
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:64
	// _ = "end of CoverTab[62430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:64
	_go_fuzz_dep_.CoverTab[62431]++
												return a.m[key]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:65
	// _ = "end of CoverTab[62431]"
}

// Equal returns whether a and o are equivalent.  If 'Equal(o interface{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:68
// bool' is implemented for a value in the attributes, it is called to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:68
// determine if the value matches the one stored in the other attributes.  If
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:68
// Equal is not implemented, standard equality is used to determine if the two
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:68
// values are equal. Note that some types (e.g. maps) aren't comparable by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:68
// default, so they must be wrapped in a struct, or in an alias type, with Equal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:68
// defined.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:75
func (a *Attributes) Equal(o *Attributes) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:75
	_go_fuzz_dep_.CoverTab[62434]++
												if a == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:76
		_go_fuzz_dep_.CoverTab[62439]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:76
		return o == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:76
		// _ = "end of CoverTab[62439]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:76
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:76
		_go_fuzz_dep_.CoverTab[62440]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:77
		// _ = "end of CoverTab[62440]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:78
		_go_fuzz_dep_.CoverTab[62441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:78
		// _ = "end of CoverTab[62441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:78
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:78
	// _ = "end of CoverTab[62434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:78
	_go_fuzz_dep_.CoverTab[62435]++
												if a == nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:79
		_go_fuzz_dep_.CoverTab[62442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:79
		return o == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:79
		// _ = "end of CoverTab[62442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:79
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:79
		_go_fuzz_dep_.CoverTab[62443]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:80
		// _ = "end of CoverTab[62443]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:81
		_go_fuzz_dep_.CoverTab[62444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:81
		// _ = "end of CoverTab[62444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:81
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:81
	// _ = "end of CoverTab[62435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:81
	_go_fuzz_dep_.CoverTab[62436]++
												if len(a.m) != len(o.m) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:82
		_go_fuzz_dep_.CoverTab[62445]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:83
		// _ = "end of CoverTab[62445]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:84
		_go_fuzz_dep_.CoverTab[62446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:84
		// _ = "end of CoverTab[62446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:84
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:84
	// _ = "end of CoverTab[62436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:84
	_go_fuzz_dep_.CoverTab[62437]++
												for k, v := range a.m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:85
		_go_fuzz_dep_.CoverTab[62447]++
													ov, ok := o.m[k]
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:87
			_go_fuzz_dep_.CoverTab[62449]++

														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:89
			// _ = "end of CoverTab[62449]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:90
			_go_fuzz_dep_.CoverTab[62450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:90
			// _ = "end of CoverTab[62450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:90
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:90
		// _ = "end of CoverTab[62447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:90
		_go_fuzz_dep_.CoverTab[62448]++
													if eq, ok := v.(interface{ Equal(o interface{}) bool }); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:91
			_go_fuzz_dep_.CoverTab[62451]++
														if !eq.Equal(ov) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:92
				_go_fuzz_dep_.CoverTab[62452]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:93
				// _ = "end of CoverTab[62452]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:94
				_go_fuzz_dep_.CoverTab[62453]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:94
				// _ = "end of CoverTab[62453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:94
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:94
			// _ = "end of CoverTab[62451]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:95
			_go_fuzz_dep_.CoverTab[62454]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:95
			if v != ov {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:95
				_go_fuzz_dep_.CoverTab[62455]++

															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:97
				// _ = "end of CoverTab[62455]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:98
				_go_fuzz_dep_.CoverTab[62456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:98
				// _ = "end of CoverTab[62456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:98
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:98
			// _ = "end of CoverTab[62454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:98
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:98
		// _ = "end of CoverTab[62448]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:99
	// _ = "end of CoverTab[62437]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:99
	_go_fuzz_dep_.CoverTab[62438]++
												return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:100
	// _ = "end of CoverTab[62438]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:101
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/attributes/attributes.go:101
var _ = _go_fuzz_dep_.CoverTab
