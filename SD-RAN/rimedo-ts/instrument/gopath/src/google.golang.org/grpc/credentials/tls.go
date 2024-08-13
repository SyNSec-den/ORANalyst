//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
package credentials

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:19
)

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/url"
	"os"

	credinternal "google.golang.org/grpc/internal/credentials"
)

// TLSInfo contains the auth information for a TLS authenticated connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:33
// It implements the AuthInfo interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:35
type TLSInfo struct {
	State	tls.ConnectionState
	CommonAuthInfo
	// This API is experimental.
	SPIFFEID	*url.URL
}

// AuthType returns the type of TLSInfo as a string.
func (t TLSInfo) AuthType() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:43
	_go_fuzz_dep_.CoverTab[62527]++
											return "tls"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:44
	// _ = "end of CoverTab[62527]"
}

// GetSecurityValue returns security info requested by channelz.
func (t TLSInfo) GetSecurityValue() ChannelzSecurityValue {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:48
	_go_fuzz_dep_.CoverTab[62528]++
											v := &TLSChannelzSecurityValue{
		StandardName: cipherSuiteLookup[t.State.CipherSuite],
	}

	if len(t.State.PeerCertificates) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:53
		_go_fuzz_dep_.CoverTab[62530]++
												v.RemoteCertificate = t.State.PeerCertificates[0].Raw
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:54
		// _ = "end of CoverTab[62530]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:55
		_go_fuzz_dep_.CoverTab[62531]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:55
		// _ = "end of CoverTab[62531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:55
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:55
	// _ = "end of CoverTab[62528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:55
	_go_fuzz_dep_.CoverTab[62529]++
											return v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:56
	// _ = "end of CoverTab[62529]"
}

// tlsCreds is the credentials required for authenticating a connection using TLS.
type tlsCreds struct {
	// TLS configuration
	config *tls.Config
}

func (c tlsCreds) Info() ProtocolInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:65
	_go_fuzz_dep_.CoverTab[62532]++
											return ProtocolInfo{
		SecurityProtocol:	"tls",
		SecurityVersion:	"1.2",
		ServerName:		c.config.ServerName,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:70
	// _ = "end of CoverTab[62532]"
}

func (c *tlsCreds) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (_ net.Conn, _ AuthInfo, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:73
	_go_fuzz_dep_.CoverTab[62533]++

											cfg := credinternal.CloneTLSConfig(c.config)
											if cfg.ServerName == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:76
		_go_fuzz_dep_.CoverTab[62538]++
												serverName, _, err := net.SplitHostPort(authority)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:78
			_go_fuzz_dep_.CoverTab[62540]++

													serverName = authority
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:80
			// _ = "end of CoverTab[62540]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:81
			_go_fuzz_dep_.CoverTab[62541]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:81
			// _ = "end of CoverTab[62541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:81
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:81
		// _ = "end of CoverTab[62538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:81
		_go_fuzz_dep_.CoverTab[62539]++
												cfg.ServerName = serverName
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:82
		// _ = "end of CoverTab[62539]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:83
		_go_fuzz_dep_.CoverTab[62542]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:83
		// _ = "end of CoverTab[62542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:83
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:83
	// _ = "end of CoverTab[62533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:83
	_go_fuzz_dep_.CoverTab[62534]++
											conn := tls.Client(rawConn, cfg)
											errChannel := make(chan error, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:85
	_curRoutineNum49_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:85
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum49_)
											go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:86
		_go_fuzz_dep_.CoverTab[62543]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:86
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:86
			_go_fuzz_dep_.CoverTab[62544]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:86
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum49_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:86
			// _ = "end of CoverTab[62544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:86
		}()
												errChannel <- conn.Handshake()
												close(errChannel)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:88
		// _ = "end of CoverTab[62543]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:89
	// _ = "end of CoverTab[62534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:89
	_go_fuzz_dep_.CoverTab[62535]++
											select {
	case err := <-errChannel:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:91
		_go_fuzz_dep_.CoverTab[62545]++
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:92
			_go_fuzz_dep_.CoverTab[62547]++
													conn.Close()
													return nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:94
			// _ = "end of CoverTab[62547]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:95
			_go_fuzz_dep_.CoverTab[62548]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:95
			// _ = "end of CoverTab[62548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:95
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:95
		// _ = "end of CoverTab[62545]"
	case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:96
		_go_fuzz_dep_.CoverTab[62546]++
												conn.Close()
												return nil, nil, ctx.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:98
		// _ = "end of CoverTab[62546]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:99
	// _ = "end of CoverTab[62535]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:99
	_go_fuzz_dep_.CoverTab[62536]++
											tlsInfo := TLSInfo{
		State:	conn.ConnectionState(),
		CommonAuthInfo: CommonAuthInfo{
			SecurityLevel: PrivacyAndIntegrity,
		},
	}
	id := credinternal.SPIFFEIDFromState(conn.ConnectionState())
	if id != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:107
		_go_fuzz_dep_.CoverTab[62549]++
													tlsInfo.SPIFFEID = id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:108
		// _ = "end of CoverTab[62549]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:109
		_go_fuzz_dep_.CoverTab[62550]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:109
		// _ = "end of CoverTab[62550]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:109
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:109
	// _ = "end of CoverTab[62536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:109
	_go_fuzz_dep_.CoverTab[62537]++
												return credinternal.WrapSyscallConn(rawConn, conn), tlsInfo, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:110
	// _ = "end of CoverTab[62537]"
}

func (c *tlsCreds) ServerHandshake(rawConn net.Conn) (net.Conn, AuthInfo, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:113
	_go_fuzz_dep_.CoverTab[62551]++
												conn := tls.Server(rawConn, c.config)
												if err := conn.Handshake(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:115
		_go_fuzz_dep_.CoverTab[62554]++
													conn.Close()
													return nil, nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:117
		// _ = "end of CoverTab[62554]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:118
		_go_fuzz_dep_.CoverTab[62555]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:118
		// _ = "end of CoverTab[62555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:118
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:118
	// _ = "end of CoverTab[62551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:118
	_go_fuzz_dep_.CoverTab[62552]++
												tlsInfo := TLSInfo{
		State:	conn.ConnectionState(),
		CommonAuthInfo: CommonAuthInfo{
			SecurityLevel: PrivacyAndIntegrity,
		},
	}
	id := credinternal.SPIFFEIDFromState(conn.ConnectionState())
	if id != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:126
		_go_fuzz_dep_.CoverTab[62556]++
													tlsInfo.SPIFFEID = id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:127
		// _ = "end of CoverTab[62556]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:128
		_go_fuzz_dep_.CoverTab[62557]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:128
		// _ = "end of CoverTab[62557]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:128
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:128
	// _ = "end of CoverTab[62552]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:128
	_go_fuzz_dep_.CoverTab[62553]++
												return credinternal.WrapSyscallConn(rawConn, conn), tlsInfo, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:129
	// _ = "end of CoverTab[62553]"
}

func (c *tlsCreds) Clone() TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:132
	_go_fuzz_dep_.CoverTab[62558]++
												return NewTLS(c.config)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:133
	// _ = "end of CoverTab[62558]"
}

func (c *tlsCreds) OverrideServerName(serverNameOverride string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:136
	_go_fuzz_dep_.CoverTab[62559]++
												c.config.ServerName = serverNameOverride
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:138
	// _ = "end of CoverTab[62559]"
}

// NewTLS uses c to construct a TransportCredentials based on TLS.
func NewTLS(c *tls.Config) TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:142
	_go_fuzz_dep_.CoverTab[62560]++
												tc := &tlsCreds{credinternal.CloneTLSConfig(c)}
												tc.config.NextProtos = credinternal.AppendH2ToNextProtos(tc.config.NextProtos)
												return tc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:145
	// _ = "end of CoverTab[62560]"
}

// NewClientTLSFromCert constructs TLS credentials from the provided root
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// certificate authority certificate(s) to validate server connections. If
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// certificates to establish the identity of the client need to be included in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// the credentials (eg: for mTLS), use NewTLS instead, where a complete
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// tls.Config can be specified.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// serverNameOverride is for testing only. If set to a non empty string,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// it will override the virtual host name of authority (e.g. :authority header
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:148
// field) in requests.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:156
func NewClientTLSFromCert(cp *x509.CertPool, serverNameOverride string) TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:156
	_go_fuzz_dep_.CoverTab[62561]++
												return NewTLS(&tls.Config{ServerName: serverNameOverride, RootCAs: cp})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:157
	// _ = "end of CoverTab[62561]"
}

// NewClientTLSFromFile constructs TLS credentials from the provided root
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// certificate authority certificate file(s) to validate server connections. If
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// certificates to establish the identity of the client need to be included in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// the credentials (eg: for mTLS), use NewTLS instead, where a complete
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// tls.Config can be specified.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// serverNameOverride is for testing only. If set to a non empty string,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// it will override the virtual host name of authority (e.g. :authority header
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:160
// field) in requests.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:168
func NewClientTLSFromFile(certFile, serverNameOverride string) (TransportCredentials, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:168
	_go_fuzz_dep_.CoverTab[62562]++
												b, err := os.ReadFile(certFile)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:170
		_go_fuzz_dep_.CoverTab[62565]++
													return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:171
		// _ = "end of CoverTab[62565]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:172
		_go_fuzz_dep_.CoverTab[62566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:172
		// _ = "end of CoverTab[62566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:172
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:172
	// _ = "end of CoverTab[62562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:172
	_go_fuzz_dep_.CoverTab[62563]++
												cp := x509.NewCertPool()
												if !cp.AppendCertsFromPEM(b) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:174
		_go_fuzz_dep_.CoverTab[62567]++
													return nil, fmt.Errorf("credentials: failed to append certificates")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:175
		// _ = "end of CoverTab[62567]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:176
		_go_fuzz_dep_.CoverTab[62568]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:176
		// _ = "end of CoverTab[62568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:176
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:176
	// _ = "end of CoverTab[62563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:176
	_go_fuzz_dep_.CoverTab[62564]++
												return NewTLS(&tls.Config{ServerName: serverNameOverride, RootCAs: cp}), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:177
	// _ = "end of CoverTab[62564]"
}

// NewServerTLSFromCert constructs TLS credentials from the input certificate for server.
func NewServerTLSFromCert(cert *tls.Certificate) TransportCredentials {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:181
	_go_fuzz_dep_.CoverTab[62569]++
												return NewTLS(&tls.Config{Certificates: []tls.Certificate{*cert}})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:182
	// _ = "end of CoverTab[62569]"
}

// NewServerTLSFromFile constructs TLS credentials from the input certificate file and key
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:185
// file for server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:187
func NewServerTLSFromFile(certFile, keyFile string) (TransportCredentials, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:187
	_go_fuzz_dep_.CoverTab[62570]++
												cert, err := tls.LoadX509KeyPair(certFile, keyFile)
												if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:189
		_go_fuzz_dep_.CoverTab[62572]++
													return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:190
		// _ = "end of CoverTab[62572]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:191
		_go_fuzz_dep_.CoverTab[62573]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:191
		// _ = "end of CoverTab[62573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:191
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:191
	// _ = "end of CoverTab[62570]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:191
	_go_fuzz_dep_.CoverTab[62571]++
												return NewTLS(&tls.Config{Certificates: []tls.Certificate{cert}}), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:192
	// _ = "end of CoverTab[62571]"
}

// TLSChannelzSecurityValue defines the struct that TLS protocol should return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:195
// from GetSecurityValue(), containing security info like cipher and certificate used.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:195
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:195
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:195
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:195
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:195
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:202
type TLSChannelzSecurityValue struct {
	ChannelzSecurityValue
	StandardName		string
	LocalCertificate	[]byte
	RemoteCertificate	[]byte
}

var cipherSuiteLookup = map[uint16]string{
	tls.TLS_RSA_WITH_RC4_128_SHA:			"TLS_RSA_WITH_RC4_128_SHA",
	tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA:		"TLS_RSA_WITH_3DES_EDE_CBC_SHA",
	tls.TLS_RSA_WITH_AES_128_CBC_SHA:		"TLS_RSA_WITH_AES_128_CBC_SHA",
	tls.TLS_RSA_WITH_AES_256_CBC_SHA:		"TLS_RSA_WITH_AES_256_CBC_SHA",
	tls.TLS_RSA_WITH_AES_128_GCM_SHA256:		"TLS_RSA_WITH_AES_128_GCM_SHA256",
	tls.TLS_RSA_WITH_AES_256_GCM_SHA384:		"TLS_RSA_WITH_AES_256_GCM_SHA384",
	tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:		"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:	"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
	tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA:		"TLS_ECDHE_RSA_WITH_RC4_128_SHA",
	tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:	"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:		"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:	"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:	"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:	"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:	"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
	tls.TLS_FALLBACK_SCSV:				"TLS_FALLBACK_SCSV",
	tls.TLS_RSA_WITH_AES_128_CBC_SHA256:		"TLS_RSA_WITH_AES_128_CBC_SHA256",
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:	"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305:	"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305:	"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
	tls.TLS_AES_128_GCM_SHA256:			"TLS_AES_128_GCM_SHA256",
	tls.TLS_AES_256_GCM_SHA384:			"TLS_AES_256_GCM_SHA384",
	tls.TLS_CHACHA20_POLY1305_SHA256:		"TLS_CHACHA20_POLY1305_SHA256",
}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:236
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/credentials/tls.go:236
var _ = _go_fuzz_dep_.CoverTab
