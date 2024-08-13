//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:19
// Package credentials implements various credentials supported by gRPC library,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:19
// which encapsulate all the state needed by a client to authenticate with a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:19
// server and make various assertions, e.g., about the client's identity, role,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:19
// or whether it is authorized to make a particular call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
package credentials

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:23
)

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/attributes"
	icredentials "google.golang.org/grpc/internal/credentials"
)

// PerRPCCredentials defines the common interface for the credentials which need to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:36
// attach security information to every RPC (e.g., oauth2).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:38
type PerRPCCredentials interface {
	// GetRequestMetadata gets the current request metadata, refreshing tokens
	// if required. This should be called by the transport layer on each
	// request, and the data should be populated in headers or other
	// context. If a status code is returned, it will be used as the status for
	// the RPC (restricted to an allowable set of codes as defined by gRFC
	// A54). uri is the URI of the entry point for the request.  When supported
	// by the underlying implementation, ctx can be used for timeout and
	// cancellation. Additionally, RequestInfo data will be available via ctx
	// to this call.  TODO(zhaoq): Define the set of the qualified keys instead
	// of leaving it as an arbitrary string.
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	// RequireTransportSecurity indicates whether the credentials requires
	// transport security.
	RequireTransportSecurity() bool
}

// SecurityLevel defines the protection level on an established connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:55
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:55
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:58
type SecurityLevel int

const (
	// InvalidSecurityLevel indicates an invalid security level.
	// The zero SecurityLevel value is invalid for backward compatibility.
	InvalidSecurityLevel	SecurityLevel	= iota
	// NoSecurity indicates a connection is insecure.
	NoSecurity
	// IntegrityOnly indicates a connection only provides integrity protection.
	IntegrityOnly
	// PrivacyAndIntegrity indicates a connection provides both privacy and integrity protection.
	PrivacyAndIntegrity
)

// String returns SecurityLevel in a string format.
func (s SecurityLevel) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:73
	_go_fuzz_dep_.CoverTab[62506]++
												switch s {
	case NoSecurity:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:75
		_go_fuzz_dep_.CoverTab[62508]++
													return "NoSecurity"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:76
		// _ = "end of CoverTab[62508]"
	case IntegrityOnly:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:77
		_go_fuzz_dep_.CoverTab[62509]++
													return "IntegrityOnly"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:78
		// _ = "end of CoverTab[62509]"
	case PrivacyAndIntegrity:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:79
		_go_fuzz_dep_.CoverTab[62510]++
													return "PrivacyAndIntegrity"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:80
		// _ = "end of CoverTab[62510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:80
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:80
		_go_fuzz_dep_.CoverTab[62511]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:80
		// _ = "end of CoverTab[62511]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:81
	// _ = "end of CoverTab[62506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:81
	_go_fuzz_dep_.CoverTab[62507]++
												return fmt.Sprintf("invalid SecurityLevel: %v", int(s))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:82
	// _ = "end of CoverTab[62507]"
}

// CommonAuthInfo contains authenticated information common to AuthInfo implementations.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:85
// It should be embedded in a struct implementing AuthInfo to provide additional information
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:85
// about the credentials.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:85
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:85
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:90
type CommonAuthInfo struct {
	SecurityLevel SecurityLevel
}

// GetCommonAuthInfo returns the pointer to CommonAuthInfo struct.
func (c CommonAuthInfo) GetCommonAuthInfo() CommonAuthInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:95
	_go_fuzz_dep_.CoverTab[62512]++
												return c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:96
	// _ = "end of CoverTab[62512]"
}

// ProtocolInfo provides information regarding the gRPC wire protocol version,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:99
// security protocol, security protocol version in use, server name, etc.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:101
type ProtocolInfo struct {
	// ProtocolVersion is the gRPC wire protocol version.
	ProtocolVersion	string
	// SecurityProtocol is the security protocol in use.
	SecurityProtocol	string
	// SecurityVersion is the security protocol version.  It is a static version string from the
	// credentials, not a value that reflects per-connection protocol negotiation.  To retrieve
	// details about the credentials used for a connection, use the Peer's AuthInfo field instead.
	//
	// Deprecated: please use Peer.AuthInfo.
	SecurityVersion	string
	// ServerName is the user-configured server name.
	ServerName	string
}

// AuthInfo defines the common interface for the auth information the users are interested in.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:116
// A struct that implements AuthInfo should embed CommonAuthInfo by including additional
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:116
// information about the credentials in it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:119
type AuthInfo interface {
	AuthType() string
}

// ErrConnDispatched indicates that rawConn has been dispatched out of gRPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:123
// and the caller should not close rawConn.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:125
var ErrConnDispatched = errors.New("credentials: rawConn is dispatched out of gRPC")

// TransportCredentials defines the common interface for all the live gRPC wire
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:127
// protocols and supported transport security protocols (e.g., TLS, SSL).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:129
type TransportCredentials interface {
	// ClientHandshake does the authentication handshake specified by the
	// corresponding authentication protocol on rawConn for clients. It returns
	// the authenticated connection and the corresponding auth information
	// about the connection.  The auth information should embed CommonAuthInfo
	// to return additional information about the credentials. Implementations
	// must use the provided context to implement timely cancellation.  gRPC
	// will try to reconnect if the error returned is a temporary error
	// (io.EOF, context.DeadlineExceeded or err.Temporary() == true).  If the
	// returned error is a wrapper error, implementations should make sure that
	// the error implements Temporary() to have the correct retry behaviors.
	// Additionally, ClientHandshakeInfo data will be available via the context
	// passed to this call.
	//
	// The second argument to this method is the `:authority` header value used
	// while creating new streams on this connection after authentication
	// succeeds. Implementations must use this as the server name during the
	// authentication handshake.
	//
	// If the returned net.Conn is closed, it MUST close the net.Conn provided.
	ClientHandshake(context.Context, string, net.Conn) (net.Conn, AuthInfo, error)
	// ServerHandshake does the authentication handshake for servers. It returns
	// the authenticated connection and the corresponding auth information about
	// the connection. The auth information should embed CommonAuthInfo to return additional information
	// about the credentials.
	//
	// If the returned net.Conn is closed, it MUST close the net.Conn provided.
	ServerHandshake(net.Conn) (net.Conn, AuthInfo, error)
	// Info provides the ProtocolInfo of this TransportCredentials.
	Info() ProtocolInfo
	// Clone makes a copy of this TransportCredentials.
	Clone() TransportCredentials
	// OverrideServerName specifies the value used for the following:
	// - verifying the hostname on the returned certificates
	// - as SNI in the client's handshake to support virtual hosting
	// - as the value for `:authority` header at stream creation time
	//
	// Deprecated: use grpc.WithAuthority instead. Will be supported
	// throughout 1.x.
	OverrideServerName(string) error
}

// Bundle is a combination of TransportCredentials and PerRPCCredentials.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
// It also contains a mode switching method, so it can be used as a combination
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
// of different credential policies.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
// Bundle cannot be used together with individual TransportCredentials.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
// PerRPCCredentials from Bundle will be appended to other PerRPCCredentials.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:171
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:180
type Bundle interface {
	// TransportCredentials returns the transport credentials from the Bundle.
	//
	// Implementations must return non-nil transport credentials. If transport
	// security is not needed by the Bundle, implementations may choose to
	// return insecure.NewCredentials().
	TransportCredentials() TransportCredentials

	// PerRPCCredentials returns the per-RPC credentials from the Bundle.
	//
	// May be nil if per-RPC credentials are not needed.
	PerRPCCredentials() PerRPCCredentials

	// NewWithMode should make a copy of Bundle, and switch mode. Modifying the
	// existing Bundle may cause races.
	//
	// NewWithMode returns nil if the requested mode is not supported.
	NewWithMode(mode string) (Bundle, error)
}

// RequestInfo contains request data attached to the context passed to GetRequestMetadata calls.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:200
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:200
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:203
type RequestInfo struct {
	// The method passed to Invoke or NewStream for this RPC. (For proto methods, this has the format "/some.Service/Method")
	Method	string
	// AuthInfo contains the information from a security handshake (TransportCredentials.ClientHandshake, TransportCredentials.ServerHandshake)
	AuthInfo	AuthInfo
}

// RequestInfoFromContext extracts the RequestInfo from the context if it exists.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:210
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:210
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:213
func RequestInfoFromContext(ctx context.Context) (ri RequestInfo, ok bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:213
	_go_fuzz_dep_.CoverTab[62513]++
													ri, ok = icredentials.RequestInfoFromContext(ctx).(RequestInfo)
													return ri, ok
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:215
	// _ = "end of CoverTab[62513]"
}

// ClientHandshakeInfo holds data to be passed to ClientHandshake. This makes
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:218
// it possible to pass arbitrary data to the handshaker from gRPC, resolver,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:218
// balancer etc. Individual credential implementations control the actual
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:218
// format of the data that they are willing to receive.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:218
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:218
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:224
type ClientHandshakeInfo struct {
	// Attributes contains the attributes for the address. It could be provided
	// by the gRPC, resolver, balancer etc.
	Attributes *attributes.Attributes
}

// ClientHandshakeInfoFromContext returns the ClientHandshakeInfo struct stored
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:230
// in ctx.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:230
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:230
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:234
func ClientHandshakeInfoFromContext(ctx context.Context) ClientHandshakeInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:234
	_go_fuzz_dep_.CoverTab[62514]++
													chi, _ := icredentials.ClientHandshakeInfoFromContext(ctx).(ClientHandshakeInfo)
													return chi
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:236
	// _ = "end of CoverTab[62514]"
}

// CheckSecurityLevel checks if a connection's security level is greater than or equal to the specified one.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:239
// It returns success if 1) the condition is satisified or 2) AuthInfo struct does not implement GetCommonAuthInfo() method
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:239
// or 3) CommonAuthInfo.SecurityLevel has an invalid zero value. For 2) and 3), it is for the purpose of backward-compatibility.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:239
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:239
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:244
func CheckSecurityLevel(ai AuthInfo, level SecurityLevel) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:244
	_go_fuzz_dep_.CoverTab[62515]++
													type internalInfo interface {
		GetCommonAuthInfo() CommonAuthInfo
	}
	if ai == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:248
		_go_fuzz_dep_.CoverTab[62518]++
														return errors.New("AuthInfo is nil")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:249
		// _ = "end of CoverTab[62518]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:250
		_go_fuzz_dep_.CoverTab[62519]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:250
		// _ = "end of CoverTab[62519]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:250
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:250
	// _ = "end of CoverTab[62515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:250
	_go_fuzz_dep_.CoverTab[62516]++
													if ci, ok := ai.(internalInfo); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:251
		_go_fuzz_dep_.CoverTab[62520]++

														if ci.GetCommonAuthInfo().SecurityLevel == InvalidSecurityLevel {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:253
			_go_fuzz_dep_.CoverTab[62522]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:254
			// _ = "end of CoverTab[62522]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:255
			_go_fuzz_dep_.CoverTab[62523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:255
			// _ = "end of CoverTab[62523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:255
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:255
		// _ = "end of CoverTab[62520]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:255
		_go_fuzz_dep_.CoverTab[62521]++
														if ci.GetCommonAuthInfo().SecurityLevel < level {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:256
			_go_fuzz_dep_.CoverTab[62524]++
															return fmt.Errorf("requires SecurityLevel %v; connection has %v", level, ci.GetCommonAuthInfo().SecurityLevel)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:257
			// _ = "end of CoverTab[62524]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:258
			_go_fuzz_dep_.CoverTab[62525]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:258
			// _ = "end of CoverTab[62525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:258
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:258
		// _ = "end of CoverTab[62521]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:259
		_go_fuzz_dep_.CoverTab[62526]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:259
		// _ = "end of CoverTab[62526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:259
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:259
	// _ = "end of CoverTab[62516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:259
	_go_fuzz_dep_.CoverTab[62517]++

													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:261
	// _ = "end of CoverTab[62517]"
}

// ChannelzSecurityInfo defines the interface that security protocols should implement
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:264
// in order to provide security info to channelz.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:264
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:264
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:268
type ChannelzSecurityInfo interface {
	GetSecurityValue() ChannelzSecurityValue
}

// ChannelzSecurityValue defines the interface that GetSecurityValue() return value
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:272
// should satisfy. This interface should only be satisfied by *TLSChannelzSecurityValue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:272
// and *OtherChannelzSecurityValue.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:272
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:272
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:277
type ChannelzSecurityValue interface {
	isChannelzSecurityValue()
}

// OtherChannelzSecurityValue defines the struct that non-TLS protocol should return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:281
// from GetSecurityValue(), which contains protocol specific security info. Note
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:281
// the Value field will be sent to users of channelz requesting channel info, and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:281
// thus sensitive info should better be avoided.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:281
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:281
// This API is experimental.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:287
type OtherChannelzSecurityValue struct {
	ChannelzSecurityValue
	Name	string
	Value	proto.Message
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:291
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/credentials.go:291
var _ = _go_fuzz_dep_.CoverTab
