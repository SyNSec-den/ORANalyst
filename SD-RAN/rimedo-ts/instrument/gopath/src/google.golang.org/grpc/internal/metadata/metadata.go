//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:19
// Package metadata contains functions to set and get metadata from addresses.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:19
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:19
// This package is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
package metadata

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:22
)

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
)

type mdKeyType string

const mdKey = mdKeyType("grpc.internal.address.metadata")

type mdValue metadata.MD

func (m mdValue) Equal(o interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:38
	_go_fuzz_dep_.CoverTab[68902]++
													om, ok := o.(mdValue)
													if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:40
		_go_fuzz_dep_.CoverTab[68906]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:41
		// _ = "end of CoverTab[68906]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:42
		_go_fuzz_dep_.CoverTab[68907]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:42
		// _ = "end of CoverTab[68907]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:42
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:42
	// _ = "end of CoverTab[68902]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:42
	_go_fuzz_dep_.CoverTab[68903]++
													if len(m) != len(om) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:43
		_go_fuzz_dep_.CoverTab[68908]++
														return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:44
		// _ = "end of CoverTab[68908]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:45
		_go_fuzz_dep_.CoverTab[68909]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:45
		// _ = "end of CoverTab[68909]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:45
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:45
	// _ = "end of CoverTab[68903]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:45
	_go_fuzz_dep_.CoverTab[68904]++
													for k, v := range m {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:46
		_go_fuzz_dep_.CoverTab[68910]++
														ov := om[k]
														if len(ov) != len(v) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:48
			_go_fuzz_dep_.CoverTab[68912]++
															return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:49
			// _ = "end of CoverTab[68912]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:50
			_go_fuzz_dep_.CoverTab[68913]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:50
			// _ = "end of CoverTab[68913]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:50
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:50
		// _ = "end of CoverTab[68910]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:50
		_go_fuzz_dep_.CoverTab[68911]++
														for i, ve := range v {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:51
			_go_fuzz_dep_.CoverTab[68914]++
															if ov[i] != ve {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:52
				_go_fuzz_dep_.CoverTab[68915]++
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:53
				// _ = "end of CoverTab[68915]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:54
				_go_fuzz_dep_.CoverTab[68916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:54
				// _ = "end of CoverTab[68916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:54
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:54
			// _ = "end of CoverTab[68914]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:55
		// _ = "end of CoverTab[68911]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:56
	// _ = "end of CoverTab[68904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:56
	_go_fuzz_dep_.CoverTab[68905]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:57
	// _ = "end of CoverTab[68905]"
}

// Get returns the metadata of addr.
func Get(addr resolver.Address) metadata.MD {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:61
	_go_fuzz_dep_.CoverTab[68917]++
													attrs := addr.Attributes
													if attrs == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:63
		_go_fuzz_dep_.CoverTab[68919]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:64
		// _ = "end of CoverTab[68919]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:65
		_go_fuzz_dep_.CoverTab[68920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:65
		// _ = "end of CoverTab[68920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:65
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:65
	// _ = "end of CoverTab[68917]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:65
	_go_fuzz_dep_.CoverTab[68918]++
													md, _ := attrs.Value(mdKey).(mdValue)
													return metadata.MD(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:67
	// _ = "end of CoverTab[68918]"
}

// Set sets (overrides) the metadata in addr.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:70
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:70
// When a SubConn is created with this address, the RPCs sent on it will all
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:70
// have this metadata.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:74
func Set(addr resolver.Address, md metadata.MD) resolver.Address {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:74
	_go_fuzz_dep_.CoverTab[68921]++
													addr.Attributes = addr.Attributes.WithValue(mdKey, mdValue(md))
													return addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:76
	// _ = "end of CoverTab[68921]"
}

// Validate validates every pair in md with ValidatePair.
func Validate(md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:80
	_go_fuzz_dep_.CoverTab[68922]++
													for k, vals := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:81
		_go_fuzz_dep_.CoverTab[68924]++
														if err := ValidatePair(k, vals...); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:82
			_go_fuzz_dep_.CoverTab[68925]++
															return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:83
			// _ = "end of CoverTab[68925]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:84
			_go_fuzz_dep_.CoverTab[68926]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:84
			// _ = "end of CoverTab[68926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:84
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:84
		// _ = "end of CoverTab[68924]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:85
	// _ = "end of CoverTab[68922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:85
	_go_fuzz_dep_.CoverTab[68923]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:86
	// _ = "end of CoverTab[68923]"
}

// hasNotPrintable return true if msg contains any characters which are not in %x20-%x7E
func hasNotPrintable(msg string) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:90
	_go_fuzz_dep_.CoverTab[68927]++

													for i := 0; i < len(msg); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:92
		_go_fuzz_dep_.CoverTab[68929]++
														if msg[i] < 0x20 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:93
			_go_fuzz_dep_.CoverTab[68930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:93
			return msg[i] > 0x7E
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:93
			// _ = "end of CoverTab[68930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:93
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:93
			_go_fuzz_dep_.CoverTab[68931]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:94
			// _ = "end of CoverTab[68931]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:95
			_go_fuzz_dep_.CoverTab[68932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:95
			// _ = "end of CoverTab[68932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:95
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:95
		// _ = "end of CoverTab[68929]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:96
	// _ = "end of CoverTab[68927]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:96
	_go_fuzz_dep_.CoverTab[68928]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:97
	// _ = "end of CoverTab[68928]"
}

// ValidatePair validate a key-value pair with the following rules (the pseudo-header will be skipped) :
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:100
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:100
// - key must contain one or more characters.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:100
// - the characters in the key must be contained in [0-9 a-z _ - .].
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:100
// - if the key ends with a "-bin" suffix, no validation of the corresponding value is performed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:100
// - the characters in the every value must be printable (in [%x20-%x7E]).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:106
func ValidatePair(key string, vals ...string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:106
	_go_fuzz_dep_.CoverTab[68933]++

													if key == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:108
		_go_fuzz_dep_.CoverTab[68939]++
														return fmt.Errorf("there is an empty key in the header")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:109
		// _ = "end of CoverTab[68939]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:110
		_go_fuzz_dep_.CoverTab[68940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:110
		// _ = "end of CoverTab[68940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:110
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:110
	// _ = "end of CoverTab[68933]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:110
	_go_fuzz_dep_.CoverTab[68934]++

													if key[0] == ':' {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:112
		_go_fuzz_dep_.CoverTab[68941]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:113
		// _ = "end of CoverTab[68941]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:114
		_go_fuzz_dep_.CoverTab[68942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:114
		// _ = "end of CoverTab[68942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:114
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:114
	// _ = "end of CoverTab[68934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:114
	_go_fuzz_dep_.CoverTab[68935]++

													for i := 0; i < len(key); i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:116
		_go_fuzz_dep_.CoverTab[68943]++
														r := key[i]
														if !(r >= 'a' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			_go_fuzz_dep_.CoverTab[68944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			return r <= 'z'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			// _ = "end of CoverTab[68944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
		}()) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			_go_fuzz_dep_.CoverTab[68945]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			return !(r >= '0' && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
				_go_fuzz_dep_.CoverTab[68946]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
				return r <= '9'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
				// _ = "end of CoverTab[68946]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			}())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			// _ = "end of CoverTab[68945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			_go_fuzz_dep_.CoverTab[68947]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			return r != '.'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			// _ = "end of CoverTab[68947]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			_go_fuzz_dep_.CoverTab[68948]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			return r != '-'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			// _ = "end of CoverTab[68948]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			_go_fuzz_dep_.CoverTab[68949]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			return r != '_'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			// _ = "end of CoverTab[68949]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:118
			_go_fuzz_dep_.CoverTab[68950]++
															return fmt.Errorf("header key %q contains illegal characters not in [0-9a-z-_.]", key)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:119
			// _ = "end of CoverTab[68950]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:120
			_go_fuzz_dep_.CoverTab[68951]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:120
			// _ = "end of CoverTab[68951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:120
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:120
		// _ = "end of CoverTab[68943]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:121
	// _ = "end of CoverTab[68935]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:121
	_go_fuzz_dep_.CoverTab[68936]++
													if strings.HasSuffix(key, "-bin") {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:122
		_go_fuzz_dep_.CoverTab[68952]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:123
		// _ = "end of CoverTab[68952]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:124
		_go_fuzz_dep_.CoverTab[68953]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:124
		// _ = "end of CoverTab[68953]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:124
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:124
	// _ = "end of CoverTab[68936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:124
	_go_fuzz_dep_.CoverTab[68937]++

													for _, val := range vals {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:126
		_go_fuzz_dep_.CoverTab[68954]++
														if hasNotPrintable(val) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:127
			_go_fuzz_dep_.CoverTab[68955]++
															return fmt.Errorf("header key %q contains value with non-printable ASCII characters", key)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:128
			// _ = "end of CoverTab[68955]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:129
			_go_fuzz_dep_.CoverTab[68956]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:129
			// _ = "end of CoverTab[68956]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:129
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:129
		// _ = "end of CoverTab[68954]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:130
	// _ = "end of CoverTab[68937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:130
	_go_fuzz_dep_.CoverTab[68938]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:131
	// _ = "end of CoverTab[68938]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:132
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/metadata/metadata.go:132
var _ = _go_fuzz_dep_.CoverTab
