//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:19
// Package insecure provides an implementation of the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:19
// credentials.TransportCredentials interface which disables transport security.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
package insecure

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:21
)

import (
	"context"
	"net"

	"google.golang.org/grpc/credentials"
)

// NewCredentials returns a credentials which disables transport security.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:30
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:30
// Note that using this credentials with per-RPC credentials which require
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:30
// transport security is incompatible and will cause grpc.Dial() to fail.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:34
func NewCredentials() credentials.TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:34
	_go_fuzz_dep_.CoverTab[67571]++
													return insecureTC{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:35
	// _ = "end of CoverTab[67571]"
}

// insecureTC implements the insecure transport credentials. The handshake
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:38
// methods simply return the passed in net.Conn and set the security level to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:38
// NoSecurity.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:41
type insecureTC struct{}

func (insecureTC) ClientHandshake(ctx context.Context, _ string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:43
	_go_fuzz_dep_.CoverTab[67572]++
													return conn, info{credentials.CommonAuthInfo{SecurityLevel: credentials.NoSecurity}}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:44
	// _ = "end of CoverTab[67572]"
}

func (insecureTC) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:47
	_go_fuzz_dep_.CoverTab[67573]++
													return conn, info{credentials.CommonAuthInfo{SecurityLevel: credentials.NoSecurity}}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:48
	// _ = "end of CoverTab[67573]"
}

func (insecureTC) Info() credentials.ProtocolInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:51
	_go_fuzz_dep_.CoverTab[67574]++
													return credentials.ProtocolInfo{SecurityProtocol: "insecure"}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:52
	// _ = "end of CoverTab[67574]"
}

func (insecureTC) Clone() credentials.TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:55
	_go_fuzz_dep_.CoverTab[67575]++
													return insecureTC{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:56
	// _ = "end of CoverTab[67575]"
}

func (insecureTC) OverrideServerName(string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:59
	_go_fuzz_dep_.CoverTab[67576]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:60
	// _ = "end of CoverTab[67576]"
}

// info contains the auth information for an insecure connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:63
// It implements the AuthInfo interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:65
type info struct {
	credentials.CommonAuthInfo
}

// AuthType returns the type of info as a string.
func (info) AuthType() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:70
	_go_fuzz_dep_.CoverTab[67577]++
													return "insecure"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:71
	// _ = "end of CoverTab[67577]"
}

// insecureBundle implements an insecure bundle.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:74
// An insecure bundle provides a thin wrapper around insecureTC to support
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:74
// the credentials.Bundle interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:77
type insecureBundle struct{}

// NewBundle returns a bundle with disabled transport security and no per rpc credential.
func NewBundle() credentials.Bundle {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:80
	_go_fuzz_dep_.CoverTab[67578]++
													return insecureBundle{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:81
	// _ = "end of CoverTab[67578]"
}

// NewWithMode returns a new insecure Bundle. The mode is ignored.
func (insecureBundle) NewWithMode(string) (credentials.Bundle, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:85
	_go_fuzz_dep_.CoverTab[67579]++
													return insecureBundle{}, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:86
	// _ = "end of CoverTab[67579]"
}

// PerRPCCredentials returns an nil implementation as insecure
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:89
// bundle does not support a per rpc credential.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:91
func (insecureBundle) PerRPCCredentials() credentials.PerRPCCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:91
	_go_fuzz_dep_.CoverTab[67580]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:92
	// _ = "end of CoverTab[67580]"
}

// TransportCredentials returns the underlying insecure transport credential.
func (insecureBundle) TransportCredentials() credentials.TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:96
	_go_fuzz_dep_.CoverTab[67581]++
													return NewCredentials()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:97
	// _ = "end of CoverTab[67581]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/insecure/insecure.go:98
var _ = _go_fuzz_dep_.CoverTab
