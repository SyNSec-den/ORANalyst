//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:19
// Package peer defines various peer information associated with RPCs and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:19
// corresponding utils.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
package peer

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:21
)

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

// Peer contains the information of the peer for an RPC, such as the address
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:30
// and authentication information.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:32
type Peer struct {
	// Addr is the peer address.
	Addr	net.Addr
	// AuthInfo is the authentication information of the transport.
	// It is nil if there is no transport security being used.
	AuthInfo	credentials.AuthInfo
}

type peerKey struct{}

// NewContext creates a new context with peer information attached.
func NewContext(ctx context.Context, p *Peer) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:43
	_go_fuzz_dep_.CoverTab[76018]++
											return context.WithValue(ctx, peerKey{}, p)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:44
	// _ = "end of CoverTab[76018]"
}

// FromContext returns the peer information in ctx if it exists.
func FromContext(ctx context.Context) (p *Peer, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:48
	_go_fuzz_dep_.CoverTab[76019]++
											p, ok = ctx.Value(peerKey{}).(*Peer)
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:50
	// _ = "end of CoverTab[76019]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:51
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/peer/peer.go:51
var _ = _go_fuzz_dep_.CoverTab
