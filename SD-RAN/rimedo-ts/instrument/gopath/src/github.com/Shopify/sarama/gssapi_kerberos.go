//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:1
)

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"time"

	"github.com/jcmturner/gofork/encoding/asn1"
	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/iana/chksumtype"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"
)

const (
	TOK_ID_KRB_AP_REQ	= 256
	GSS_API_GENERIC_TAG	= 0x60
	KRB5_USER_AUTH		= 1
	KRB5_KEYTAB_AUTH	= 2
	GSS_API_INITIAL		= 1
	GSS_API_VERIFY		= 2
	GSS_API_FINISH		= 3
)

type GSSAPIConfig struct {
	AuthType		int
	KeyTabPath		string
	KerberosConfigPath	string
	ServiceName		string
	Username		string
	Password		string
	Realm			string
	DisablePAFXFAST		bool
}

type GSSAPIKerberosAuth struct {
	Config			*GSSAPIConfig
	ticket			messages.Ticket
	encKey			types.EncryptionKey
	NewKerberosClientFunc	func(config *GSSAPIConfig) (KerberosClient, error)
	step			int
}

type KerberosClient interface {
	Login() error
	GetServiceTicket(spn string) (messages.Ticket, types.EncryptionKey, error)
	Domain() string
	CName() types.PrincipalName
	Destroy()
}

// writePackage appends length in big endian before the payload, and sends it to kafka
func (krbAuth *GSSAPIKerberosAuth) writePackage(broker *Broker, payload []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:59
	_go_fuzz_dep_.CoverTab[103323]++
												length := uint64(len(payload))
												size := length + 4
												if size > math.MaxInt32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:62
		_go_fuzz_dep_.CoverTab[103326]++
													return 0, errors.New("payload too large, will overflow int32")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:63
		// _ = "end of CoverTab[103326]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:64
		_go_fuzz_dep_.CoverTab[103327]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:64
		// _ = "end of CoverTab[103327]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:64
	// _ = "end of CoverTab[103323]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:64
	_go_fuzz_dep_.CoverTab[103324]++
												finalPackage := make([]byte, size)
												copy(finalPackage[4:], payload)
												binary.BigEndian.PutUint32(finalPackage, uint32(length))
												bytes, err := broker.conn.Write(finalPackage)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:69
		_go_fuzz_dep_.CoverTab[103328]++
													return bytes, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:70
		// _ = "end of CoverTab[103328]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:71
		_go_fuzz_dep_.CoverTab[103329]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:71
		// _ = "end of CoverTab[103329]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:71
	// _ = "end of CoverTab[103324]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:71
	_go_fuzz_dep_.CoverTab[103325]++
												return bytes, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:72
	// _ = "end of CoverTab[103325]"
}

// readPackage reads payload length (4 bytes) and then reads the payload into []byte
func (krbAuth *GSSAPIKerberosAuth) readPackage(broker *Broker) ([]byte, int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:76
	_go_fuzz_dep_.CoverTab[103330]++
												bytesRead := 0
												lengthInBytes := make([]byte, 4)
												bytes, err := io.ReadFull(broker.conn, lengthInBytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:80
		_go_fuzz_dep_.CoverTab[103333]++
													return nil, bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:81
		// _ = "end of CoverTab[103333]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:82
		_go_fuzz_dep_.CoverTab[103334]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:82
		// _ = "end of CoverTab[103334]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:82
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:82
	// _ = "end of CoverTab[103330]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:82
	_go_fuzz_dep_.CoverTab[103331]++
												bytesRead += bytes
												payloadLength := binary.BigEndian.Uint32(lengthInBytes)
												payloadBytes := make([]byte, payloadLength)
												bytes, err = io.ReadFull(broker.conn, payloadBytes)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:87
		_go_fuzz_dep_.CoverTab[103335]++
													return payloadBytes, bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:88
		// _ = "end of CoverTab[103335]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:89
		_go_fuzz_dep_.CoverTab[103336]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:89
		// _ = "end of CoverTab[103336]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:89
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:89
	// _ = "end of CoverTab[103331]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:89
	_go_fuzz_dep_.CoverTab[103332]++
												bytesRead += bytes
												return payloadBytes, bytesRead, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:91
	// _ = "end of CoverTab[103332]"
}

func (krbAuth *GSSAPIKerberosAuth) newAuthenticatorChecksum() []byte {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:94
	_go_fuzz_dep_.CoverTab[103337]++
												a := make([]byte, 24)
												flags := []int{gssapi.ContextFlagInteg, gssapi.ContextFlagConf}
												binary.LittleEndian.PutUint32(a[:4], 16)
												for _, i := range flags {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:98
		_go_fuzz_dep_.CoverTab[103339]++
													f := binary.LittleEndian.Uint32(a[20:24])
													f |= uint32(i)
													binary.LittleEndian.PutUint32(a[20:24], f)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:101
		// _ = "end of CoverTab[103339]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:102
	// _ = "end of CoverTab[103337]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:102
	_go_fuzz_dep_.CoverTab[103338]++
												return a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:103
	// _ = "end of CoverTab[103338]"
}

/*
*
* Construct Kerberos AP_REQ package, conforming to RFC-4120
* https://tools.ietf.org/html/rfc4120#page-84
*
 */
func (krbAuth *GSSAPIKerberosAuth) createKrb5Token(
	domain string, cname types.PrincipalName,
	ticket messages.Ticket,
	sessionKey types.EncryptionKey) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:115
	_go_fuzz_dep_.CoverTab[103340]++
												auth, err := types.NewAuthenticator(domain, cname)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:117
		_go_fuzz_dep_.CoverTab[103344]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:118
		// _ = "end of CoverTab[103344]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:119
		_go_fuzz_dep_.CoverTab[103345]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:119
		// _ = "end of CoverTab[103345]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:119
	// _ = "end of CoverTab[103340]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:119
	_go_fuzz_dep_.CoverTab[103341]++
												auth.Cksum = types.Checksum{
		CksumType:	chksumtype.GSSAPI,
		Checksum:	krbAuth.newAuthenticatorChecksum(),
	}
	APReq, err := messages.NewAPReq(
		ticket,
		sessionKey,
		auth,
	)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:129
		_go_fuzz_dep_.CoverTab[103346]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:130
		// _ = "end of CoverTab[103346]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:131
		_go_fuzz_dep_.CoverTab[103347]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:131
		// _ = "end of CoverTab[103347]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:131
	// _ = "end of CoverTab[103341]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:131
	_go_fuzz_dep_.CoverTab[103342]++
												aprBytes := make([]byte, 2)
												binary.BigEndian.PutUint16(aprBytes, TOK_ID_KRB_AP_REQ)
												tb, err := APReq.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:135
		_go_fuzz_dep_.CoverTab[103348]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:136
		// _ = "end of CoverTab[103348]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:137
		_go_fuzz_dep_.CoverTab[103349]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:137
		// _ = "end of CoverTab[103349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:137
	// _ = "end of CoverTab[103342]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:137
	_go_fuzz_dep_.CoverTab[103343]++
												aprBytes = append(aprBytes, tb...)
												return aprBytes, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:139
	// _ = "end of CoverTab[103343]"
}

/*
*
*	Append the GSS-API header to the payload, conforming to RFC-2743
*	Section 3.1, Mechanism-Independent Token Format
*
*	https://tools.ietf.org/html/rfc2743#page-81
*
*	GSSAPIHeader + <specific mechanism payload>
*
 */
func (krbAuth *GSSAPIKerberosAuth) appendGSSAPIHeader(payload []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:152
	_go_fuzz_dep_.CoverTab[103350]++
												oidBytes, err := asn1.Marshal(gssapi.OIDKRB5.OID())
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:154
		_go_fuzz_dep_.CoverTab[103352]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:155
		// _ = "end of CoverTab[103352]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:156
		_go_fuzz_dep_.CoverTab[103353]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:156
		// _ = "end of CoverTab[103353]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:156
	// _ = "end of CoverTab[103350]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:156
	_go_fuzz_dep_.CoverTab[103351]++
												tkoLengthBytes := asn1tools.MarshalLengthBytes(len(oidBytes) + len(payload))
												GSSHeader := append([]byte{GSS_API_GENERIC_TAG}, tkoLengthBytes...)
												GSSHeader = append(GSSHeader, oidBytes...)
												GSSPackage := append(GSSHeader, payload...)
												return GSSPackage, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:161
	// _ = "end of CoverTab[103351]"
}

func (krbAuth *GSSAPIKerberosAuth) initSecContext(bytes []byte, kerberosClient KerberosClient) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:164
	_go_fuzz_dep_.CoverTab[103354]++
												switch krbAuth.step {
	case GSS_API_INITIAL:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:166
		_go_fuzz_dep_.CoverTab[103356]++
													aprBytes, err := krbAuth.createKrb5Token(
			kerberosClient.Domain(),
			kerberosClient.CName(),
			krbAuth.ticket,
			krbAuth.encKey)
		if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:172
			_go_fuzz_dep_.CoverTab[103363]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:173
			// _ = "end of CoverTab[103363]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:174
			_go_fuzz_dep_.CoverTab[103364]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:174
			// _ = "end of CoverTab[103364]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:174
		// _ = "end of CoverTab[103356]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:174
		_go_fuzz_dep_.CoverTab[103357]++
													krbAuth.step = GSS_API_VERIFY
													return krbAuth.appendGSSAPIHeader(aprBytes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:176
		// _ = "end of CoverTab[103357]"
	case GSS_API_VERIFY:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:177
		_go_fuzz_dep_.CoverTab[103358]++
													wrapTokenReq := gssapi.WrapToken{}
													if err := wrapTokenReq.Unmarshal(bytes, true); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:179
			_go_fuzz_dep_.CoverTab[103365]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:180
			// _ = "end of CoverTab[103365]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:181
			_go_fuzz_dep_.CoverTab[103366]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:181
			// _ = "end of CoverTab[103366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:181
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:181
		// _ = "end of CoverTab[103358]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:181
		_go_fuzz_dep_.CoverTab[103359]++

													isValid, err := wrapTokenReq.Verify(krbAuth.encKey, keyusage.GSSAPI_ACCEPTOR_SEAL)
													if !isValid {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:184
			_go_fuzz_dep_.CoverTab[103367]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:185
			// _ = "end of CoverTab[103367]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:186
			_go_fuzz_dep_.CoverTab[103368]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:186
			// _ = "end of CoverTab[103368]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:186
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:186
		// _ = "end of CoverTab[103359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:186
		_go_fuzz_dep_.CoverTab[103360]++

													wrapTokenResponse, err := gssapi.NewInitiatorWrapToken(wrapTokenReq.Payload, krbAuth.encKey)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:189
			_go_fuzz_dep_.CoverTab[103369]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:190
			// _ = "end of CoverTab[103369]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:191
			_go_fuzz_dep_.CoverTab[103370]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:191
			// _ = "end of CoverTab[103370]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:191
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:191
		// _ = "end of CoverTab[103360]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:191
		_go_fuzz_dep_.CoverTab[103361]++
													krbAuth.step = GSS_API_FINISH
													return wrapTokenResponse.Marshal()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:193
		// _ = "end of CoverTab[103361]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:193
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:193
		_go_fuzz_dep_.CoverTab[103362]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:193
		// _ = "end of CoverTab[103362]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:194
	// _ = "end of CoverTab[103354]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:194
	_go_fuzz_dep_.CoverTab[103355]++
												return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:195
	// _ = "end of CoverTab[103355]"
}

/* This does the handshake for authorization */
func (krbAuth *GSSAPIKerberosAuth) Authorize(broker *Broker) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:199
	_go_fuzz_dep_.CoverTab[103371]++
												kerberosClient, err := krbAuth.NewKerberosClientFunc(krbAuth.Config)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:201
		_go_fuzz_dep_.CoverTab[103375]++
													Logger.Printf("Kerberos client error: %s", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:203
		// _ = "end of CoverTab[103375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:204
		_go_fuzz_dep_.CoverTab[103376]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:204
		// _ = "end of CoverTab[103376]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:204
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:204
	// _ = "end of CoverTab[103371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:204
	_go_fuzz_dep_.CoverTab[103372]++

												err = kerberosClient.Login()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:207
		_go_fuzz_dep_.CoverTab[103377]++
													Logger.Printf("Kerberos client error: %s", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:209
		// _ = "end of CoverTab[103377]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:210
		_go_fuzz_dep_.CoverTab[103378]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:210
		// _ = "end of CoverTab[103378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:210
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:210
	// _ = "end of CoverTab[103372]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:210
	_go_fuzz_dep_.CoverTab[103373]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:214
	host := strings.SplitN(broker.addr, ":", 2)[0]
	spn := fmt.Sprintf("%s/%s", broker.conf.Net.SASL.GSSAPI.ServiceName, host)

	ticket, encKey, err := kerberosClient.GetServiceTicket(spn)
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:218
		_go_fuzz_dep_.CoverTab[103379]++
													Logger.Printf("Error getting Kerberos service ticket : %s", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:220
		// _ = "end of CoverTab[103379]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:221
		_go_fuzz_dep_.CoverTab[103380]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:221
		// _ = "end of CoverTab[103380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:221
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:221
	// _ = "end of CoverTab[103373]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:221
	_go_fuzz_dep_.CoverTab[103374]++
												krbAuth.ticket = ticket
												krbAuth.encKey = encKey
												krbAuth.step = GSS_API_INITIAL
												var receivedBytes []byte = nil
												defer kerberosClient.Destroy()
												for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:227
		_go_fuzz_dep_.CoverTab[103381]++
													packBytes, err := krbAuth.initSecContext(receivedBytes, kerberosClient)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:229
			_go_fuzz_dep_.CoverTab[103384]++
														Logger.Printf("Error while performing GSSAPI Kerberos Authentication: %s\n", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:231
			// _ = "end of CoverTab[103384]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:232
			_go_fuzz_dep_.CoverTab[103385]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:232
			// _ = "end of CoverTab[103385]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:232
		// _ = "end of CoverTab[103381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:232
		_go_fuzz_dep_.CoverTab[103382]++
													requestTime := time.Now()
													bytesWritten, err := krbAuth.writePackage(broker, packBytes)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:235
			_go_fuzz_dep_.CoverTab[103386]++
														Logger.Printf("Error while performing GSSAPI Kerberos Authentication: %s\n", err)
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:237
			// _ = "end of CoverTab[103386]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:238
			_go_fuzz_dep_.CoverTab[103387]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:238
			// _ = "end of CoverTab[103387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:238
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:238
		// _ = "end of CoverTab[103382]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:238
		_go_fuzz_dep_.CoverTab[103383]++
													broker.updateOutgoingCommunicationMetrics(bytesWritten)
													if krbAuth.step == GSS_API_VERIFY {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:240
			_go_fuzz_dep_.CoverTab[103388]++
														bytesRead := 0
														receivedBytes, bytesRead, err = krbAuth.readPackage(broker)
														requestLatency := time.Since(requestTime)
														broker.updateIncomingCommunicationMetrics(bytesRead, requestLatency)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:245
				_go_fuzz_dep_.CoverTab[103389]++
															Logger.Printf("Error while performing GSSAPI Kerberos Authentication: %s\n", err)
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:247
				// _ = "end of CoverTab[103389]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:248
				_go_fuzz_dep_.CoverTab[103390]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:248
				// _ = "end of CoverTab[103390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:248
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:248
			// _ = "end of CoverTab[103388]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:249
			_go_fuzz_dep_.CoverTab[103391]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:249
			if krbAuth.step == GSS_API_FINISH {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:249
				_go_fuzz_dep_.CoverTab[103392]++
															return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:250
				// _ = "end of CoverTab[103392]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:251
				_go_fuzz_dep_.CoverTab[103393]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:251
				// _ = "end of CoverTab[103393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:251
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:251
			// _ = "end of CoverTab[103391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:251
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:251
		// _ = "end of CoverTab[103383]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:252
	// _ = "end of CoverTab[103374]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:253
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/gssapi_kerberos.go:253
var _ = _go_fuzz_dep_.CoverTab
