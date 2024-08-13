//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:19
// Package networktype declares the network type to be used in the default
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:19
// dialer. Attribute of a resolver.Address.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
package networktype

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:21
)

import (
	"google.golang.org/grpc/resolver"
)

// keyType is the key to use for storing State in Attributes.
type keyType string

const key = keyType("grpc.internal.transport.networktype")

// Set returns a copy of the provided address with attributes containing networkType.
func Set(address resolver.Address, networkType string) resolver.Address {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:33
	_go_fuzz_dep_.CoverTab[69160]++
															address.Attributes = address.Attributes.WithValue(key, networkType)
															return address
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:35
	// _ = "end of CoverTab[69160]"
}

// Get returns the network type in the resolver.Address and true, or "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:38
// if not present.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:40
func Get(address resolver.Address) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:40
	_go_fuzz_dep_.CoverTab[69161]++
															v := address.Attributes.Value(key)
															if v == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:42
		_go_fuzz_dep_.CoverTab[69163]++
																return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:43
		// _ = "end of CoverTab[69163]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:44
		_go_fuzz_dep_.CoverTab[69164]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:44
		// _ = "end of CoverTab[69164]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:44
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:44
	// _ = "end of CoverTab[69161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:44
	_go_fuzz_dep_.CoverTab[69162]++
															return v.(string), true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:45
	// _ = "end of CoverTab[69162]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:46
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/networktype/networktype.go:46
var _ = _go_fuzz_dep_.CoverTab
