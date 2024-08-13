//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
package internal

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:17
)

import (
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
)

// handshakeClusterNameKey is the type used as the key to store cluster name in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:24
// the Attributes field of resolver.Address.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:26
type handshakeClusterNameKey struct{}

// SetXDSHandshakeClusterName returns a copy of addr in which the Attributes field
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:28
// is updated with the cluster name.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:30
func SetXDSHandshakeClusterName(addr resolver.Address, clusterName string) resolver.Address {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:30
	_go_fuzz_dep_.CoverTab[67329]++
													addr.Attributes = addr.Attributes.WithValue(handshakeClusterNameKey{}, clusterName)
													return addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:32
	// _ = "end of CoverTab[67329]"
}

// GetXDSHandshakeClusterName returns cluster name stored in attr.
func GetXDSHandshakeClusterName(attr *attributes.Attributes) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:36
	_go_fuzz_dep_.CoverTab[67330]++
													v := attr.Value(handshakeClusterNameKey{})
													name, ok := v.(string)
													return name, ok
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:39
	// _ = "end of CoverTab[67330]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:40
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/xds_handshake_cluster.go:40
var _ = _go_fuzz_dep_.CoverTab
