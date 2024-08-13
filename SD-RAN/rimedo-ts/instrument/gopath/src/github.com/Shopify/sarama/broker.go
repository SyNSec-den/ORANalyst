//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1
)

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rcrowley/go-metrics"
)

// Broker represents a single Kafka broker connection. All operations on this object are entirely concurrency-safe.
type Broker struct {
	conf	*Config
	rack	*string

	id		int32
	addr		string
	correlationID	int32
	conn		net.Conn
	connErr		error
	lock		sync.Mutex
	opened		int32
	responses	chan *responsePromise
	done		chan bool

	registeredMetrics	[]string

	incomingByteRate	metrics.Meter
	requestRate		metrics.Meter
	requestSize		metrics.Histogram
	requestLatency		metrics.Histogram
	outgoingByteRate	metrics.Meter
	responseRate		metrics.Meter
	responseSize		metrics.Histogram
	requestsInFlight	metrics.Counter
	brokerIncomingByteRate	metrics.Meter
	brokerRequestRate	metrics.Meter
	brokerRequestSize	metrics.Histogram
	brokerRequestLatency	metrics.Histogram
	brokerOutgoingByteRate	metrics.Meter
	brokerResponseRate	metrics.Meter
	brokerResponseSize	metrics.Histogram
	brokerRequestsInFlight	metrics.Counter
	brokerThrottleTime	metrics.Histogram

	kerberosAuthenticator	GSSAPIKerberosAuth
}

// SASLMechanism specifies the SASL mechanism the client uses to authenticate with the broker
type SASLMechanism string

const (
	// SASLTypeOAuth represents the SASL/OAUTHBEARER mechanism (Kafka 2.0.0+)
	SASLTypeOAuth	= "OAUTHBEARER"
	// SASLTypePlaintext represents the SASL/PLAIN mechanism
	SASLTypePlaintext	= "PLAIN"
	// SASLTypeSCRAMSHA256 represents the SCRAM-SHA-256 mechanism.
	SASLTypeSCRAMSHA256	= "SCRAM-SHA-256"
	// SASLTypeSCRAMSHA512 represents the SCRAM-SHA-512 mechanism.
	SASLTypeSCRAMSHA512	= "SCRAM-SHA-512"
	SASLTypeGSSAPI		= "GSSAPI"
	// SASLHandshakeV0 is v0 of the Kafka SASL handshake protocol. Client and
	// server negotiate SASL auth using opaque packets.
	SASLHandshakeV0	= int16(0)
	// SASLHandshakeV1 is v1 of the Kafka SASL handshake protocol. Client and
	// server negotiate SASL by wrapping tokens with Kafka protocol headers.
	SASLHandshakeV1	= int16(1)
	// SASLExtKeyAuth is the reserved extension key name sent as part of the
	// SASL/OAUTHBEARER initial client response
	SASLExtKeyAuth	= "auth"
)

// AccessToken contains an access token used to authenticate a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:81
// SASL/OAUTHBEARER client along with associated metadata.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:83
type AccessToken struct {
	// Token is the access token payload.
	Token	string
	// Extensions is a optional map of arbitrary key-value pairs that can be
	// sent with the SASL/OAUTHBEARER initial client response. These values are
	// ignored by the SASL server if they are unexpected. This feature is only
	// supported by Kafka >= 2.1.0.
	Extensions	map[string]string
}

// AccessTokenProvider is the interface that encapsulates how implementors
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:93
// can generate access tokens for Kafka broker authentication.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:95
type AccessTokenProvider interface {
	// Token returns an access token. The implementation should ensure token
	// reuse so that multiple calls at connect time do not create multiple
	// tokens. The implementation should also periodically refresh the token in
	// order to guarantee that each call returns an unexpired token.  This
	// method should not block indefinitely--a timeout error should be returned
	// after a short period of inactivity so that the broker connection logic
	// can log debugging information and retry.
	Token() (*AccessToken, error)
}

// SCRAMClient is a an interface to a SCRAM
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:106
// client implementation.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:108
type SCRAMClient interface {
	// Begin prepares the client for the SCRAM exchange
	// with the server with a user name and a password
	Begin(userName, password, authzID string) error
	// Step steps client through the SCRAM exchange. It is
	// called repeatedly until it errors or `Done` returns true.
	Step(challenge string) (response string, err error)
	// Done should return true when the SCRAM conversation
	// is over.
	Done() bool
}

type responsePromise struct {
	requestTime	time.Time
	correlationID	int32
	headerVersion	int16
	handler		func([]byte, error)
	packets		chan []byte
	errors		chan error
}

func (p *responsePromise) handle(packets []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:129
	_go_fuzz_dep_.CoverTab[99370]++

											if p.handler != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:131
		_go_fuzz_dep_.CoverTab[99373]++
												p.handler(packets, err)
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:133
		// _ = "end of CoverTab[99373]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:134
		_go_fuzz_dep_.CoverTab[99374]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:134
		// _ = "end of CoverTab[99374]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:134
	// _ = "end of CoverTab[99370]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:134
	_go_fuzz_dep_.CoverTab[99371]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:136
		_go_fuzz_dep_.CoverTab[99375]++
												p.errors <- err
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:138
		// _ = "end of CoverTab[99375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:139
		_go_fuzz_dep_.CoverTab[99376]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:139
		// _ = "end of CoverTab[99376]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:139
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:139
	// _ = "end of CoverTab[99371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:139
	_go_fuzz_dep_.CoverTab[99372]++
											p.packets <- packets
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:140
	// _ = "end of CoverTab[99372]"
}

// NewBroker creates and returns a Broker targeting the given host:port address.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:143
// This does not attempt to actually connect, you have to call Open() for that.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:145
func NewBroker(addr string) *Broker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:145
	_go_fuzz_dep_.CoverTab[99377]++
											return &Broker{id: -1, addr: addr}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:146
	// _ = "end of CoverTab[99377]"
}

// Open tries to connect to the Broker if it is not already connected or connecting, but does not block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:149
// waiting for the connection to complete. This means that any subsequent operations on the broker will
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:149
// block waiting for the connection to succeed or fail. To get the effect of a fully synchronous Open call,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:149
// follow it by a call to Connected(). The only errors Open will return directly are ConfigurationError or
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:149
// AlreadyConnected. If conf is nil, the result of NewConfig() is used.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:154
func (b *Broker) Open(conf *Config) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:154
	_go_fuzz_dep_.CoverTab[99378]++
											if !atomic.CompareAndSwapInt32(&b.opened, 0, 1) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:155
		_go_fuzz_dep_.CoverTab[99383]++
												return ErrAlreadyConnected
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:156
		// _ = "end of CoverTab[99383]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:157
		_go_fuzz_dep_.CoverTab[99384]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:157
		// _ = "end of CoverTab[99384]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:157
	// _ = "end of CoverTab[99378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:157
	_go_fuzz_dep_.CoverTab[99379]++

											if conf == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:159
		_go_fuzz_dep_.CoverTab[99385]++
												conf = NewConfig()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:160
		// _ = "end of CoverTab[99385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:161
		_go_fuzz_dep_.CoverTab[99386]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:161
		// _ = "end of CoverTab[99386]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:161
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:161
	// _ = "end of CoverTab[99379]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:161
	_go_fuzz_dep_.CoverTab[99380]++

											err := conf.Validate()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:164
		_go_fuzz_dep_.CoverTab[99387]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:165
		// _ = "end of CoverTab[99387]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:166
		_go_fuzz_dep_.CoverTab[99388]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:166
		// _ = "end of CoverTab[99388]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:166
	// _ = "end of CoverTab[99380]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:166
	_go_fuzz_dep_.CoverTab[99381]++

											usingApiVersionsRequests := conf.Version.IsAtLeast(V2_4_0_0) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:168
		_go_fuzz_dep_.CoverTab[99389]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:168
		return conf.ApiVersionsRequest
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:168
		// _ = "end of CoverTab[99389]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:168
	}()

											b.lock.Lock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:170
	_curRoutineNum125_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:170
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum125_)

											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
		_go_fuzz_dep_.CoverTab[99390]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
			_go_fuzz_dep_.CoverTab[99391]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum125_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
			// _ = "end of CoverTab[99391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
		}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
		withRecover(func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:172
			_go_fuzz_dep_.CoverTab[99392]++
													defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:173
				_go_fuzz_dep_.CoverTab[99399]++
														b.lock.Unlock()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:179
				if usingApiVersionsRequests {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:179
					_go_fuzz_dep_.CoverTab[99400]++
															_, err = b.ApiVersions(&ApiVersionsRequest{
						Version:		3,
						ClientSoftwareName:	defaultClientSoftwareName,
						ClientSoftwareVersion:	version(),
					})
					if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:185
						_go_fuzz_dep_.CoverTab[99401]++
																Logger.Printf("Error while sending ApiVersionsRequest to broker %s: %s\n", b.addr, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:186
						// _ = "end of CoverTab[99401]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:187
						_go_fuzz_dep_.CoverTab[99402]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:187
						// _ = "end of CoverTab[99402]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:187
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:187
					// _ = "end of CoverTab[99400]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:188
					_go_fuzz_dep_.CoverTab[99403]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:188
					// _ = "end of CoverTab[99403]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:188
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:188
				// _ = "end of CoverTab[99399]"
			}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:189
			// _ = "end of CoverTab[99392]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:189
			_go_fuzz_dep_.CoverTab[99393]++
													dialer := conf.getDialer()
													b.conn, b.connErr = dialer.Dial("tcp", b.addr)
													if b.connErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:192
				_go_fuzz_dep_.CoverTab[99404]++
														Logger.Printf("Failed to connect to broker %s: %s\n", b.addr, b.connErr)
														b.conn = nil
														atomic.StoreInt32(&b.opened, 0)
														return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:196
				// _ = "end of CoverTab[99404]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:197
				_go_fuzz_dep_.CoverTab[99405]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:197
				// _ = "end of CoverTab[99405]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:197
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:197
			// _ = "end of CoverTab[99393]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:197
			_go_fuzz_dep_.CoverTab[99394]++
													if conf.Net.TLS.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:198
				_go_fuzz_dep_.CoverTab[99406]++
														b.conn = tls.Client(b.conn, validServerNameTLS(b.addr, conf.Net.TLS.Config))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:199
				// _ = "end of CoverTab[99406]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:200
				_go_fuzz_dep_.CoverTab[99407]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:200
				// _ = "end of CoverTab[99407]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:200
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:200
			// _ = "end of CoverTab[99394]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:200
			_go_fuzz_dep_.CoverTab[99395]++

													b.conn = newBufConn(b.conn)
													b.conf = conf

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:206
			b.incomingByteRate = metrics.GetOrRegisterMeter("incoming-byte-rate", conf.MetricRegistry)
													b.requestRate = metrics.GetOrRegisterMeter("request-rate", conf.MetricRegistry)
													b.requestSize = getOrRegisterHistogram("request-size", conf.MetricRegistry)
													b.requestLatency = getOrRegisterHistogram("request-latency-in-ms", conf.MetricRegistry)
													b.outgoingByteRate = metrics.GetOrRegisterMeter("outgoing-byte-rate", conf.MetricRegistry)
													b.responseRate = metrics.GetOrRegisterMeter("response-rate", conf.MetricRegistry)
													b.responseSize = getOrRegisterHistogram("response-size", conf.MetricRegistry)
													b.requestsInFlight = metrics.GetOrRegisterCounter("requests-in-flight", conf.MetricRegistry)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:216
			if b.id >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:216
				_go_fuzz_dep_.CoverTab[99408]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:216
				return !metrics.UseNilMetrics
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:216
				// _ = "end of CoverTab[99408]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:216
			}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:216
				_go_fuzz_dep_.CoverTab[99409]++
														b.registerMetrics()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:217
				// _ = "end of CoverTab[99409]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:218
				_go_fuzz_dep_.CoverTab[99410]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:218
				// _ = "end of CoverTab[99410]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:218
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:218
			// _ = "end of CoverTab[99395]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:218
			_go_fuzz_dep_.CoverTab[99396]++

													if conf.Net.SASL.Enable {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:220
				_go_fuzz_dep_.CoverTab[99411]++
														b.connErr = b.authenticateViaSASL()

														if b.connErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:223
					_go_fuzz_dep_.CoverTab[99412]++
															err = b.conn.Close()
															if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:225
						_go_fuzz_dep_.CoverTab[99414]++
																DebugLogger.Printf("Closed connection to broker %s\n", b.addr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:226
						// _ = "end of CoverTab[99414]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:227
						_go_fuzz_dep_.CoverTab[99415]++
																Logger.Printf("Error while closing connection to broker %s: %s\n", b.addr, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:228
						// _ = "end of CoverTab[99415]"
					}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:229
					// _ = "end of CoverTab[99412]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:229
					_go_fuzz_dep_.CoverTab[99413]++
															b.conn = nil
															atomic.StoreInt32(&b.opened, 0)
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:232
					// _ = "end of CoverTab[99413]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:233
					_go_fuzz_dep_.CoverTab[99416]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:233
					// _ = "end of CoverTab[99416]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:233
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:233
				// _ = "end of CoverTab[99411]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:234
				_go_fuzz_dep_.CoverTab[99417]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:234
				// _ = "end of CoverTab[99417]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:234
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:234
			// _ = "end of CoverTab[99396]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:234
			_go_fuzz_dep_.CoverTab[99397]++

													b.done = make(chan bool)
													b.responses = make(chan *responsePromise, b.conf.Net.MaxOpenRequests-1)

													if b.id >= 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:239
				_go_fuzz_dep_.CoverTab[99418]++
														DebugLogger.Printf("Connected to broker at %s (registered as #%d)\n", b.addr, b.id)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:240
				// _ = "end of CoverTab[99418]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:241
				_go_fuzz_dep_.CoverTab[99419]++
														DebugLogger.Printf("Connected to broker at %s (unregistered)\n", b.addr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:242
				// _ = "end of CoverTab[99419]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:243
			// _ = "end of CoverTab[99397]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:243
			_go_fuzz_dep_.CoverTab[99398]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:243
			_curRoutineNum126_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:243
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum126_)
													go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
				_go_fuzz_dep_.CoverTab[99420]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
				defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
					_go_fuzz_dep_.CoverTab[99421]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum126_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
					// _ = "end of CoverTab[99421]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
				}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
				withRecover(b.responseReceiver)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
				// _ = "end of CoverTab[99420]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
			}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:244
			// _ = "end of CoverTab[99398]"
		})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:245
		// _ = "end of CoverTab[99390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:245
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:245
	// _ = "end of CoverTab[99381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:245
	_go_fuzz_dep_.CoverTab[99382]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:247
	// _ = "end of CoverTab[99382]"
}

// Connected returns true if the broker is connected and false otherwise. If the broker is not
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:250
// connected but it had tried to connect, the error from that connection attempt is also returned.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:252
func (b *Broker) Connected() (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:252
	_go_fuzz_dep_.CoverTab[99422]++
											b.lock.Lock()
											defer b.lock.Unlock()

											return b.conn != nil, b.connErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:256
	// _ = "end of CoverTab[99422]"
}

// TLSConnectionState returns the client's TLS connection state. The second return value is false if this is not a tls connection or the connection has not yet been established.
func (b *Broker) TLSConnectionState() (state tls.ConnectionState, ok bool) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:260
	_go_fuzz_dep_.CoverTab[99423]++
											b.lock.Lock()
											defer b.lock.Unlock()

											if b.conn == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:264
		_go_fuzz_dep_.CoverTab[99427]++
												return state, false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:265
		// _ = "end of CoverTab[99427]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:266
		_go_fuzz_dep_.CoverTab[99428]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:266
		// _ = "end of CoverTab[99428]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:266
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:266
	// _ = "end of CoverTab[99423]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:266
	_go_fuzz_dep_.CoverTab[99424]++
											conn := b.conn
											if bconn, ok := b.conn.(*bufConn); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:268
		_go_fuzz_dep_.CoverTab[99429]++
												conn = bconn.Conn
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:269
		// _ = "end of CoverTab[99429]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:270
		_go_fuzz_dep_.CoverTab[99430]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:270
		// _ = "end of CoverTab[99430]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:270
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:270
	// _ = "end of CoverTab[99424]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:270
	_go_fuzz_dep_.CoverTab[99425]++
											if tc, ok := conn.(*tls.Conn); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:271
		_go_fuzz_dep_.CoverTab[99431]++
												return tc.ConnectionState(), true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:272
		// _ = "end of CoverTab[99431]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:273
		_go_fuzz_dep_.CoverTab[99432]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:273
		// _ = "end of CoverTab[99432]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:273
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:273
	// _ = "end of CoverTab[99425]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:273
	_go_fuzz_dep_.CoverTab[99426]++
											return state, false
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:274
	// _ = "end of CoverTab[99426]"
}

// Close closes the broker resources
func (b *Broker) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:278
	_go_fuzz_dep_.CoverTab[99433]++
											b.lock.Lock()
											defer b.lock.Unlock()

											if b.conn == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:282
		_go_fuzz_dep_.CoverTab[99436]++
												return ErrNotConnected
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:283
		// _ = "end of CoverTab[99436]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:284
		_go_fuzz_dep_.CoverTab[99437]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:284
		// _ = "end of CoverTab[99437]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:284
	// _ = "end of CoverTab[99433]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:284
	_go_fuzz_dep_.CoverTab[99434]++

											close(b.responses)
											<-b.done

											err := b.conn.Close()

											b.conn = nil
											b.connErr = nil
											b.done = nil
											b.responses = nil

											b.unregisterMetrics()

											if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:298
		_go_fuzz_dep_.CoverTab[99438]++
												DebugLogger.Printf("Closed connection to broker %s\n", b.addr)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:299
		// _ = "end of CoverTab[99438]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:300
		_go_fuzz_dep_.CoverTab[99439]++
												Logger.Printf("Error while closing connection to broker %s: %s\n", b.addr, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:301
		// _ = "end of CoverTab[99439]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:302
	// _ = "end of CoverTab[99434]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:302
	_go_fuzz_dep_.CoverTab[99435]++

											atomic.StoreInt32(&b.opened, 0)

											return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:306
	// _ = "end of CoverTab[99435]"
}

// ID returns the broker ID retrieved from Kafka's metadata, or -1 if that is not known.
func (b *Broker) ID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:310
	_go_fuzz_dep_.CoverTab[99440]++
											return b.id
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:311
	// _ = "end of CoverTab[99440]"
}

// Addr returns the broker address as either retrieved from Kafka's metadata or passed to NewBroker.
func (b *Broker) Addr() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:315
	_go_fuzz_dep_.CoverTab[99441]++
											return b.addr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:316
	// _ = "end of CoverTab[99441]"
}

// Rack returns the broker's rack as retrieved from Kafka's metadata or the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:319
// empty string if it is not known.  The returned value corresponds to the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:319
// broker's broker.rack configuration setting.  Requires protocol version to be
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:319
// at least v0.10.0.0.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:323
func (b *Broker) Rack() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:323
	_go_fuzz_dep_.CoverTab[99442]++
											if b.rack == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:324
		_go_fuzz_dep_.CoverTab[99444]++
												return ""
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:325
		// _ = "end of CoverTab[99444]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:326
		_go_fuzz_dep_.CoverTab[99445]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:326
		// _ = "end of CoverTab[99445]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:326
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:326
	// _ = "end of CoverTab[99442]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:326
	_go_fuzz_dep_.CoverTab[99443]++
											return *b.rack
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:327
	// _ = "end of CoverTab[99443]"
}

// GetMetadata send a metadata request and returns a metadata response or error
func (b *Broker) GetMetadata(request *MetadataRequest) (*MetadataResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:331
	_go_fuzz_dep_.CoverTab[99446]++
											response := new(MetadataResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:335
		_go_fuzz_dep_.CoverTab[99448]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:336
		// _ = "end of CoverTab[99448]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:337
		_go_fuzz_dep_.CoverTab[99449]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:337
		// _ = "end of CoverTab[99449]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:337
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:337
	// _ = "end of CoverTab[99446]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:337
	_go_fuzz_dep_.CoverTab[99447]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:339
	// _ = "end of CoverTab[99447]"
}

// GetConsumerMetadata send a consumer metadata request and returns a consumer metadata response or error
func (b *Broker) GetConsumerMetadata(request *ConsumerMetadataRequest) (*ConsumerMetadataResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:343
	_go_fuzz_dep_.CoverTab[99450]++
											response := new(ConsumerMetadataResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:347
		_go_fuzz_dep_.CoverTab[99452]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:348
		// _ = "end of CoverTab[99452]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:349
		_go_fuzz_dep_.CoverTab[99453]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:349
		// _ = "end of CoverTab[99453]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:349
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:349
	// _ = "end of CoverTab[99450]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:349
	_go_fuzz_dep_.CoverTab[99451]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:351
	// _ = "end of CoverTab[99451]"
}

// FindCoordinator sends a find coordinate request and returns a response or error
func (b *Broker) FindCoordinator(request *FindCoordinatorRequest) (*FindCoordinatorResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:355
	_go_fuzz_dep_.CoverTab[99454]++
											response := new(FindCoordinatorResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:359
		_go_fuzz_dep_.CoverTab[99456]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:360
		// _ = "end of CoverTab[99456]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:361
		_go_fuzz_dep_.CoverTab[99457]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:361
		// _ = "end of CoverTab[99457]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:361
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:361
	// _ = "end of CoverTab[99454]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:361
	_go_fuzz_dep_.CoverTab[99455]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:363
	// _ = "end of CoverTab[99455]"
}

// GetAvailableOffsets return an offset response or error
func (b *Broker) GetAvailableOffsets(request *OffsetRequest) (*OffsetResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:367
	_go_fuzz_dep_.CoverTab[99458]++
											response := new(OffsetResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:371
		_go_fuzz_dep_.CoverTab[99460]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:372
		// _ = "end of CoverTab[99460]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:373
		_go_fuzz_dep_.CoverTab[99461]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:373
		// _ = "end of CoverTab[99461]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:373
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:373
	// _ = "end of CoverTab[99458]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:373
	_go_fuzz_dep_.CoverTab[99459]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:375
	// _ = "end of CoverTab[99459]"
}

// ProduceCallback function is called once the produce response has been parsed
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:378
// or could not be read.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:380
type ProduceCallback func(*ProduceResponse, error)

// AsyncProduce sends a produce request and eventually call the provided callback
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// with a produce response or an error.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// Waiting for the response is generally not blocking on the contrary to using Produce.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// If the maximum number of in flight request configured is reached then
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// the request will be blocked till a previous response is received.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// When configured with RequiredAcks == NoResponse, the callback will not be invoked.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// If an error is returned because the request could not be sent then the callback
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:382
// will not be invoked either.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:392
func (b *Broker) AsyncProduce(request *ProduceRequest, cb ProduceCallback) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:392
	_go_fuzz_dep_.CoverTab[99462]++
											needAcks := request.RequiredAcks != NoResponse
	// Use a nil promise when no acks is required
	var promise *responsePromise

	if needAcks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:397
		_go_fuzz_dep_.CoverTab[99464]++

												res := new(ProduceResponse)
												promise = &responsePromise{
			headerVersion:	res.headerVersion(),

			handler: func(packets []byte, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:403
				_go_fuzz_dep_.CoverTab[99465]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:404
					_go_fuzz_dep_.CoverTab[99468]++

															cb(nil, err)
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:407
					// _ = "end of CoverTab[99468]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:408
					_go_fuzz_dep_.CoverTab[99469]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:408
					// _ = "end of CoverTab[99469]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:408
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:408
				// _ = "end of CoverTab[99465]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:408
				_go_fuzz_dep_.CoverTab[99466]++

														if err := versionedDecode(packets, res, request.version()); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:410
					_go_fuzz_dep_.CoverTab[99470]++

															cb(nil, err)
															return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:413
					// _ = "end of CoverTab[99470]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:414
					_go_fuzz_dep_.CoverTab[99471]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:414
					// _ = "end of CoverTab[99471]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:414
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:414
				// _ = "end of CoverTab[99466]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:414
				_go_fuzz_dep_.CoverTab[99467]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:417
				b.updateThrottleMetric(res.ThrottleTime)
														cb(res, nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:418
				// _ = "end of CoverTab[99467]"
			},
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:420
		// _ = "end of CoverTab[99464]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:421
		_go_fuzz_dep_.CoverTab[99472]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:421
		// _ = "end of CoverTab[99472]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:421
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:421
	// _ = "end of CoverTab[99462]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:421
	_go_fuzz_dep_.CoverTab[99463]++

											return b.sendWithPromise(request, promise)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:423
	// _ = "end of CoverTab[99463]"
}

// Produce returns a produce response or error
func (b *Broker) Produce(request *ProduceRequest) (*ProduceResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:427
	_go_fuzz_dep_.CoverTab[99473]++
											var (
		response	*ProduceResponse
		err		error
	)

	if request.RequiredAcks == NoResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:433
		_go_fuzz_dep_.CoverTab[99476]++
												err = b.sendAndReceive(request, nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:434
		// _ = "end of CoverTab[99476]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:435
		_go_fuzz_dep_.CoverTab[99477]++
												response = new(ProduceResponse)
												err = b.sendAndReceive(request, response)
												b.updateThrottleMetric(response.ThrottleTime)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:438
		// _ = "end of CoverTab[99477]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:439
	// _ = "end of CoverTab[99473]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:439
	_go_fuzz_dep_.CoverTab[99474]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:441
		_go_fuzz_dep_.CoverTab[99478]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:442
		// _ = "end of CoverTab[99478]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:443
		_go_fuzz_dep_.CoverTab[99479]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:443
		// _ = "end of CoverTab[99479]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:443
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:443
	// _ = "end of CoverTab[99474]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:443
	_go_fuzz_dep_.CoverTab[99475]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:445
	// _ = "end of CoverTab[99475]"
}

// Fetch returns a FetchResponse or error
func (b *Broker) Fetch(request *FetchRequest) (*FetchResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:449
	_go_fuzz_dep_.CoverTab[99480]++
											response := new(FetchResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:453
		_go_fuzz_dep_.CoverTab[99482]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:454
		// _ = "end of CoverTab[99482]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:455
		_go_fuzz_dep_.CoverTab[99483]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:455
		// _ = "end of CoverTab[99483]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:455
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:455
	// _ = "end of CoverTab[99480]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:455
	_go_fuzz_dep_.CoverTab[99481]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:457
	// _ = "end of CoverTab[99481]"
}

// CommitOffset return an Offset commit response or error
func (b *Broker) CommitOffset(request *OffsetCommitRequest) (*OffsetCommitResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:461
	_go_fuzz_dep_.CoverTab[99484]++
											response := new(OffsetCommitResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:465
		_go_fuzz_dep_.CoverTab[99486]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:466
		// _ = "end of CoverTab[99486]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:467
		_go_fuzz_dep_.CoverTab[99487]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:467
		// _ = "end of CoverTab[99487]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:467
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:467
	// _ = "end of CoverTab[99484]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:467
	_go_fuzz_dep_.CoverTab[99485]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:469
	// _ = "end of CoverTab[99485]"
}

// FetchOffset returns an offset fetch response or error
func (b *Broker) FetchOffset(request *OffsetFetchRequest) (*OffsetFetchResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:473
	_go_fuzz_dep_.CoverTab[99488]++
											response := new(OffsetFetchResponse)
											response.Version = request.Version

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:478
		_go_fuzz_dep_.CoverTab[99490]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:479
		// _ = "end of CoverTab[99490]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:480
		_go_fuzz_dep_.CoverTab[99491]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:480
		// _ = "end of CoverTab[99491]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:480
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:480
	// _ = "end of CoverTab[99488]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:480
	_go_fuzz_dep_.CoverTab[99489]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:482
	// _ = "end of CoverTab[99489]"
}

// JoinGroup returns a join group response or error
func (b *Broker) JoinGroup(request *JoinGroupRequest) (*JoinGroupResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:486
	_go_fuzz_dep_.CoverTab[99492]++
											response := new(JoinGroupResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:490
		_go_fuzz_dep_.CoverTab[99494]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:491
		// _ = "end of CoverTab[99494]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:492
		_go_fuzz_dep_.CoverTab[99495]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:492
		// _ = "end of CoverTab[99495]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:492
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:492
	// _ = "end of CoverTab[99492]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:492
	_go_fuzz_dep_.CoverTab[99493]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:494
	// _ = "end of CoverTab[99493]"
}

// SyncGroup returns a sync group response or error
func (b *Broker) SyncGroup(request *SyncGroupRequest) (*SyncGroupResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:498
	_go_fuzz_dep_.CoverTab[99496]++
											response := new(SyncGroupResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:502
		_go_fuzz_dep_.CoverTab[99498]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:503
		// _ = "end of CoverTab[99498]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:504
		_go_fuzz_dep_.CoverTab[99499]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:504
		// _ = "end of CoverTab[99499]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:504
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:504
	// _ = "end of CoverTab[99496]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:504
	_go_fuzz_dep_.CoverTab[99497]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:506
	// _ = "end of CoverTab[99497]"
}

// LeaveGroup return a leave group response or error
func (b *Broker) LeaveGroup(request *LeaveGroupRequest) (*LeaveGroupResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:510
	_go_fuzz_dep_.CoverTab[99500]++
											response := new(LeaveGroupResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:514
		_go_fuzz_dep_.CoverTab[99502]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:515
		// _ = "end of CoverTab[99502]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:516
		_go_fuzz_dep_.CoverTab[99503]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:516
		// _ = "end of CoverTab[99503]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:516
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:516
	// _ = "end of CoverTab[99500]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:516
	_go_fuzz_dep_.CoverTab[99501]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:518
	// _ = "end of CoverTab[99501]"
}

// Heartbeat returns a heartbeat response or error
func (b *Broker) Heartbeat(request *HeartbeatRequest) (*HeartbeatResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:522
	_go_fuzz_dep_.CoverTab[99504]++
											response := new(HeartbeatResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:526
		_go_fuzz_dep_.CoverTab[99506]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:527
		// _ = "end of CoverTab[99506]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:528
		_go_fuzz_dep_.CoverTab[99507]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:528
		// _ = "end of CoverTab[99507]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:528
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:528
	// _ = "end of CoverTab[99504]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:528
	_go_fuzz_dep_.CoverTab[99505]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:530
	// _ = "end of CoverTab[99505]"
}

// ListGroups return a list group response or error
func (b *Broker) ListGroups(request *ListGroupsRequest) (*ListGroupsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:534
	_go_fuzz_dep_.CoverTab[99508]++
											response := new(ListGroupsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:538
		_go_fuzz_dep_.CoverTab[99510]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:539
		// _ = "end of CoverTab[99510]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:540
		_go_fuzz_dep_.CoverTab[99511]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:540
		// _ = "end of CoverTab[99511]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:540
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:540
	// _ = "end of CoverTab[99508]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:540
	_go_fuzz_dep_.CoverTab[99509]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:542
	// _ = "end of CoverTab[99509]"
}

// DescribeGroups return describe group response or error
func (b *Broker) DescribeGroups(request *DescribeGroupsRequest) (*DescribeGroupsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:546
	_go_fuzz_dep_.CoverTab[99512]++
											response := new(DescribeGroupsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:550
		_go_fuzz_dep_.CoverTab[99514]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:551
		// _ = "end of CoverTab[99514]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:552
		_go_fuzz_dep_.CoverTab[99515]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:552
		// _ = "end of CoverTab[99515]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:552
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:552
	// _ = "end of CoverTab[99512]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:552
	_go_fuzz_dep_.CoverTab[99513]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:554
	// _ = "end of CoverTab[99513]"
}

// ApiVersions return api version response or error
func (b *Broker) ApiVersions(request *ApiVersionsRequest) (*ApiVersionsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:558
	_go_fuzz_dep_.CoverTab[99516]++
											response := new(ApiVersionsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:562
		_go_fuzz_dep_.CoverTab[99518]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:563
		// _ = "end of CoverTab[99518]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:564
		_go_fuzz_dep_.CoverTab[99519]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:564
		// _ = "end of CoverTab[99519]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:564
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:564
	// _ = "end of CoverTab[99516]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:564
	_go_fuzz_dep_.CoverTab[99517]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:566
	// _ = "end of CoverTab[99517]"
}

// CreateTopics send a create topic request and returns create topic response
func (b *Broker) CreateTopics(request *CreateTopicsRequest) (*CreateTopicsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:570
	_go_fuzz_dep_.CoverTab[99520]++
											response := new(CreateTopicsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:574
		_go_fuzz_dep_.CoverTab[99522]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:575
		// _ = "end of CoverTab[99522]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:576
		_go_fuzz_dep_.CoverTab[99523]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:576
		// _ = "end of CoverTab[99523]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:576
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:576
	// _ = "end of CoverTab[99520]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:576
	_go_fuzz_dep_.CoverTab[99521]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:578
	// _ = "end of CoverTab[99521]"
}

// DeleteTopics sends a delete topic request and returns delete topic response
func (b *Broker) DeleteTopics(request *DeleteTopicsRequest) (*DeleteTopicsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:582
	_go_fuzz_dep_.CoverTab[99524]++
											response := new(DeleteTopicsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:586
		_go_fuzz_dep_.CoverTab[99526]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:587
		// _ = "end of CoverTab[99526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:588
		_go_fuzz_dep_.CoverTab[99527]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:588
		// _ = "end of CoverTab[99527]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:588
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:588
	// _ = "end of CoverTab[99524]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:588
	_go_fuzz_dep_.CoverTab[99525]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:590
	// _ = "end of CoverTab[99525]"
}

// CreatePartitions sends a create partition request and returns create
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:593
// partitions response or error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:595
func (b *Broker) CreatePartitions(request *CreatePartitionsRequest) (*CreatePartitionsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:595
	_go_fuzz_dep_.CoverTab[99528]++
											response := new(CreatePartitionsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:599
		_go_fuzz_dep_.CoverTab[99530]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:600
		// _ = "end of CoverTab[99530]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:601
		_go_fuzz_dep_.CoverTab[99531]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:601
		// _ = "end of CoverTab[99531]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:601
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:601
	// _ = "end of CoverTab[99528]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:601
	_go_fuzz_dep_.CoverTab[99529]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:603
	// _ = "end of CoverTab[99529]"
}

// AlterPartitionReassignments sends a alter partition reassignments request and
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:606
// returns alter partition reassignments response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:608
func (b *Broker) AlterPartitionReassignments(request *AlterPartitionReassignmentsRequest) (*AlterPartitionReassignmentsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:608
	_go_fuzz_dep_.CoverTab[99532]++
											response := new(AlterPartitionReassignmentsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:612
		_go_fuzz_dep_.CoverTab[99534]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:613
		// _ = "end of CoverTab[99534]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:614
		_go_fuzz_dep_.CoverTab[99535]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:614
		// _ = "end of CoverTab[99535]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:614
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:614
	// _ = "end of CoverTab[99532]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:614
	_go_fuzz_dep_.CoverTab[99533]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:616
	// _ = "end of CoverTab[99533]"
}

// ListPartitionReassignments sends a list partition reassignments request and
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:619
// returns list partition reassignments response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:621
func (b *Broker) ListPartitionReassignments(request *ListPartitionReassignmentsRequest) (*ListPartitionReassignmentsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:621
	_go_fuzz_dep_.CoverTab[99536]++
											response := new(ListPartitionReassignmentsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:625
		_go_fuzz_dep_.CoverTab[99538]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:626
		// _ = "end of CoverTab[99538]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:627
		_go_fuzz_dep_.CoverTab[99539]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:627
		// _ = "end of CoverTab[99539]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:627
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:627
	// _ = "end of CoverTab[99536]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:627
	_go_fuzz_dep_.CoverTab[99537]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:629
	// _ = "end of CoverTab[99537]"
}

// DeleteRecords send a request to delete records and return delete record
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:632
// response or error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:634
func (b *Broker) DeleteRecords(request *DeleteRecordsRequest) (*DeleteRecordsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:634
	_go_fuzz_dep_.CoverTab[99540]++
											response := new(DeleteRecordsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:638
		_go_fuzz_dep_.CoverTab[99542]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:639
		// _ = "end of CoverTab[99542]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:640
		_go_fuzz_dep_.CoverTab[99543]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:640
		// _ = "end of CoverTab[99543]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:640
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:640
	// _ = "end of CoverTab[99540]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:640
	_go_fuzz_dep_.CoverTab[99541]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:642
	// _ = "end of CoverTab[99541]"
}

// DescribeAcls sends a describe acl request and returns a response or error
func (b *Broker) DescribeAcls(request *DescribeAclsRequest) (*DescribeAclsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:646
	_go_fuzz_dep_.CoverTab[99544]++
											response := new(DescribeAclsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:650
		_go_fuzz_dep_.CoverTab[99546]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:651
		// _ = "end of CoverTab[99546]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:652
		_go_fuzz_dep_.CoverTab[99547]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:652
		// _ = "end of CoverTab[99547]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:652
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:652
	// _ = "end of CoverTab[99544]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:652
	_go_fuzz_dep_.CoverTab[99545]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:654
	// _ = "end of CoverTab[99545]"
}

// CreateAcls sends a create acl request and returns a response or error
func (b *Broker) CreateAcls(request *CreateAclsRequest) (*CreateAclsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:658
	_go_fuzz_dep_.CoverTab[99548]++
											response := new(CreateAclsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:662
		_go_fuzz_dep_.CoverTab[99550]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:663
		// _ = "end of CoverTab[99550]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:664
		_go_fuzz_dep_.CoverTab[99551]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:664
		// _ = "end of CoverTab[99551]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:664
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:664
	// _ = "end of CoverTab[99548]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:664
	_go_fuzz_dep_.CoverTab[99549]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:666
	// _ = "end of CoverTab[99549]"
}

// DeleteAcls sends a delete acl request and returns a response or error
func (b *Broker) DeleteAcls(request *DeleteAclsRequest) (*DeleteAclsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:670
	_go_fuzz_dep_.CoverTab[99552]++
											response := new(DeleteAclsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:674
		_go_fuzz_dep_.CoverTab[99554]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:675
		// _ = "end of CoverTab[99554]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:676
		_go_fuzz_dep_.CoverTab[99555]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:676
		// _ = "end of CoverTab[99555]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:676
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:676
	// _ = "end of CoverTab[99552]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:676
	_go_fuzz_dep_.CoverTab[99553]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:678
	// _ = "end of CoverTab[99553]"
}

// InitProducerID sends an init producer request and returns a response or error
func (b *Broker) InitProducerID(request *InitProducerIDRequest) (*InitProducerIDResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:682
	_go_fuzz_dep_.CoverTab[99556]++
											response := new(InitProducerIDResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:686
		_go_fuzz_dep_.CoverTab[99558]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:687
		// _ = "end of CoverTab[99558]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:688
		_go_fuzz_dep_.CoverTab[99559]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:688
		// _ = "end of CoverTab[99559]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:688
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:688
	// _ = "end of CoverTab[99556]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:688
	_go_fuzz_dep_.CoverTab[99557]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:690
	// _ = "end of CoverTab[99557]"
}

// AddPartitionsToTxn send a request to add partition to txn and returns
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:693
// a response or error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:695
func (b *Broker) AddPartitionsToTxn(request *AddPartitionsToTxnRequest) (*AddPartitionsToTxnResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:695
	_go_fuzz_dep_.CoverTab[99560]++
											response := new(AddPartitionsToTxnResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:699
		_go_fuzz_dep_.CoverTab[99562]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:700
		// _ = "end of CoverTab[99562]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:701
		_go_fuzz_dep_.CoverTab[99563]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:701
		// _ = "end of CoverTab[99563]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:701
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:701
	// _ = "end of CoverTab[99560]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:701
	_go_fuzz_dep_.CoverTab[99561]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:703
	// _ = "end of CoverTab[99561]"
}

// AddOffsetsToTxn sends a request to add offsets to txn and returns a response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:706
// or error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:708
func (b *Broker) AddOffsetsToTxn(request *AddOffsetsToTxnRequest) (*AddOffsetsToTxnResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:708
	_go_fuzz_dep_.CoverTab[99564]++
											response := new(AddOffsetsToTxnResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:712
		_go_fuzz_dep_.CoverTab[99566]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:713
		// _ = "end of CoverTab[99566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:714
		_go_fuzz_dep_.CoverTab[99567]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:714
		// _ = "end of CoverTab[99567]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:714
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:714
	// _ = "end of CoverTab[99564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:714
	_go_fuzz_dep_.CoverTab[99565]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:716
	// _ = "end of CoverTab[99565]"
}

// EndTxn sends a request to end txn and returns a response or error
func (b *Broker) EndTxn(request *EndTxnRequest) (*EndTxnResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:720
	_go_fuzz_dep_.CoverTab[99568]++
											response := new(EndTxnResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:724
		_go_fuzz_dep_.CoverTab[99570]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:725
		// _ = "end of CoverTab[99570]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:726
		_go_fuzz_dep_.CoverTab[99571]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:726
		// _ = "end of CoverTab[99571]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:726
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:726
	// _ = "end of CoverTab[99568]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:726
	_go_fuzz_dep_.CoverTab[99569]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:728
	// _ = "end of CoverTab[99569]"
}

// TxnOffsetCommit sends a request to commit transaction offsets and returns
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:731
// a response or error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:733
func (b *Broker) TxnOffsetCommit(request *TxnOffsetCommitRequest) (*TxnOffsetCommitResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:733
	_go_fuzz_dep_.CoverTab[99572]++
											response := new(TxnOffsetCommitResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:737
		_go_fuzz_dep_.CoverTab[99574]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:738
		// _ = "end of CoverTab[99574]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:739
		_go_fuzz_dep_.CoverTab[99575]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:739
		// _ = "end of CoverTab[99575]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:739
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:739
	// _ = "end of CoverTab[99572]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:739
	_go_fuzz_dep_.CoverTab[99573]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:741
	// _ = "end of CoverTab[99573]"
}

// DescribeConfigs sends a request to describe config and returns a response or
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:744
// error
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:746
func (b *Broker) DescribeConfigs(request *DescribeConfigsRequest) (*DescribeConfigsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:746
	_go_fuzz_dep_.CoverTab[99576]++
											response := new(DescribeConfigsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:750
		_go_fuzz_dep_.CoverTab[99578]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:751
		// _ = "end of CoverTab[99578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:752
		_go_fuzz_dep_.CoverTab[99579]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:752
		// _ = "end of CoverTab[99579]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:752
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:752
	// _ = "end of CoverTab[99576]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:752
	_go_fuzz_dep_.CoverTab[99577]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:754
	// _ = "end of CoverTab[99577]"
}

// AlterConfigs sends a request to alter config and return a response or error
func (b *Broker) AlterConfigs(request *AlterConfigsRequest) (*AlterConfigsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:758
	_go_fuzz_dep_.CoverTab[99580]++
											response := new(AlterConfigsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:762
		_go_fuzz_dep_.CoverTab[99582]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:763
		// _ = "end of CoverTab[99582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:764
		_go_fuzz_dep_.CoverTab[99583]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:764
		// _ = "end of CoverTab[99583]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:764
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:764
	// _ = "end of CoverTab[99580]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:764
	_go_fuzz_dep_.CoverTab[99581]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:766
	// _ = "end of CoverTab[99581]"
}

// IncrementalAlterConfigs sends a request to incremental alter config and return a response or error
func (b *Broker) IncrementalAlterConfigs(request *IncrementalAlterConfigsRequest) (*IncrementalAlterConfigsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:770
	_go_fuzz_dep_.CoverTab[99584]++
											response := new(IncrementalAlterConfigsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:774
		_go_fuzz_dep_.CoverTab[99586]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:775
		// _ = "end of CoverTab[99586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:776
		_go_fuzz_dep_.CoverTab[99587]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:776
		// _ = "end of CoverTab[99587]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:776
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:776
	// _ = "end of CoverTab[99584]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:776
	_go_fuzz_dep_.CoverTab[99585]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:778
	// _ = "end of CoverTab[99585]"
}

// DeleteGroups sends a request to delete groups and returns a response or error
func (b *Broker) DeleteGroups(request *DeleteGroupsRequest) (*DeleteGroupsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:782
	_go_fuzz_dep_.CoverTab[99588]++
											response := new(DeleteGroupsResponse)

											if err := b.sendAndReceive(request, response); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:785
		_go_fuzz_dep_.CoverTab[99590]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:786
		// _ = "end of CoverTab[99590]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:787
		_go_fuzz_dep_.CoverTab[99591]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:787
		// _ = "end of CoverTab[99591]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:787
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:787
	// _ = "end of CoverTab[99588]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:787
	_go_fuzz_dep_.CoverTab[99589]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:789
	// _ = "end of CoverTab[99589]"
}

// DeleteOffsets sends a request to delete group offsets and returns a response or error
func (b *Broker) DeleteOffsets(request *DeleteOffsetsRequest) (*DeleteOffsetsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:793
	_go_fuzz_dep_.CoverTab[99592]++
											response := new(DeleteOffsetsResponse)

											if err := b.sendAndReceive(request, response); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:796
		_go_fuzz_dep_.CoverTab[99594]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:797
		// _ = "end of CoverTab[99594]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:798
		_go_fuzz_dep_.CoverTab[99595]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:798
		// _ = "end of CoverTab[99595]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:798
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:798
	// _ = "end of CoverTab[99592]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:798
	_go_fuzz_dep_.CoverTab[99593]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:800
	// _ = "end of CoverTab[99593]"
}

// DescribeLogDirs sends a request to get the broker's log dir paths and sizes
func (b *Broker) DescribeLogDirs(request *DescribeLogDirsRequest) (*DescribeLogDirsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:804
	_go_fuzz_dep_.CoverTab[99596]++
											response := new(DescribeLogDirsResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:808
		_go_fuzz_dep_.CoverTab[99598]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:809
		// _ = "end of CoverTab[99598]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:810
		_go_fuzz_dep_.CoverTab[99599]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:810
		// _ = "end of CoverTab[99599]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:810
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:810
	// _ = "end of CoverTab[99596]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:810
	_go_fuzz_dep_.CoverTab[99597]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:812
	// _ = "end of CoverTab[99597]"
}

// DescribeUserScramCredentials sends a request to get SCRAM users
func (b *Broker) DescribeUserScramCredentials(req *DescribeUserScramCredentialsRequest) (*DescribeUserScramCredentialsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:816
	_go_fuzz_dep_.CoverTab[99600]++
											res := new(DescribeUserScramCredentialsResponse)

											err := b.sendAndReceive(req, res)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:820
		_go_fuzz_dep_.CoverTab[99602]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:821
		// _ = "end of CoverTab[99602]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:822
		_go_fuzz_dep_.CoverTab[99603]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:822
		// _ = "end of CoverTab[99603]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:822
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:822
	// _ = "end of CoverTab[99600]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:822
	_go_fuzz_dep_.CoverTab[99601]++

											return res, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:824
	// _ = "end of CoverTab[99601]"
}

func (b *Broker) AlterUserScramCredentials(req *AlterUserScramCredentialsRequest) (*AlterUserScramCredentialsResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:827
	_go_fuzz_dep_.CoverTab[99604]++
											res := new(AlterUserScramCredentialsResponse)

											err := b.sendAndReceive(req, res)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:831
		_go_fuzz_dep_.CoverTab[99606]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:832
		// _ = "end of CoverTab[99606]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:833
		_go_fuzz_dep_.CoverTab[99607]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:833
		// _ = "end of CoverTab[99607]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:833
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:833
	// _ = "end of CoverTab[99604]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:833
	_go_fuzz_dep_.CoverTab[99605]++

											return res, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:835
	// _ = "end of CoverTab[99605]"
}

// DescribeClientQuotas sends a request to get the broker's quotas
func (b *Broker) DescribeClientQuotas(request *DescribeClientQuotasRequest) (*DescribeClientQuotasResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:839
	_go_fuzz_dep_.CoverTab[99608]++
											response := new(DescribeClientQuotasResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:843
		_go_fuzz_dep_.CoverTab[99610]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:844
		// _ = "end of CoverTab[99610]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:845
		_go_fuzz_dep_.CoverTab[99611]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:845
		// _ = "end of CoverTab[99611]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:845
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:845
	// _ = "end of CoverTab[99608]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:845
	_go_fuzz_dep_.CoverTab[99609]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:847
	// _ = "end of CoverTab[99609]"
}

// AlterClientQuotas sends a request to alter the broker's quotas
func (b *Broker) AlterClientQuotas(request *AlterClientQuotasRequest) (*AlterClientQuotasResponse, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:851
	_go_fuzz_dep_.CoverTab[99612]++
											response := new(AlterClientQuotasResponse)

											err := b.sendAndReceive(request, response)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:855
		_go_fuzz_dep_.CoverTab[99614]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:856
		// _ = "end of CoverTab[99614]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:857
		_go_fuzz_dep_.CoverTab[99615]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:857
		// _ = "end of CoverTab[99615]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:857
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:857
	// _ = "end of CoverTab[99612]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:857
	_go_fuzz_dep_.CoverTab[99613]++

											return response, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:859
	// _ = "end of CoverTab[99613]"
}

// readFull ensures the conn ReadDeadline has been setup before making a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:862
// call to io.ReadFull
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:864
func (b *Broker) readFull(buf []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:864
	_go_fuzz_dep_.CoverTab[99616]++
											if err := b.conn.SetReadDeadline(time.Now().Add(b.conf.Net.ReadTimeout)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:865
		_go_fuzz_dep_.CoverTab[99618]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:866
		// _ = "end of CoverTab[99618]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:867
		_go_fuzz_dep_.CoverTab[99619]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:867
		// _ = "end of CoverTab[99619]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:867
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:867
	// _ = "end of CoverTab[99616]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:867
	_go_fuzz_dep_.CoverTab[99617]++

											return io.ReadFull(b.conn, buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:869
	// _ = "end of CoverTab[99617]"
}

// write  ensures the conn WriteDeadline has been setup before making a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:872
// call to conn.Write
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:874
func (b *Broker) write(buf []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:874
	_go_fuzz_dep_.CoverTab[99620]++
											if err := b.conn.SetWriteDeadline(time.Now().Add(b.conf.Net.WriteTimeout)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:875
		_go_fuzz_dep_.CoverTab[99622]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:876
		// _ = "end of CoverTab[99622]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:877
		_go_fuzz_dep_.CoverTab[99623]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:877
		// _ = "end of CoverTab[99623]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:877
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:877
	// _ = "end of CoverTab[99620]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:877
	_go_fuzz_dep_.CoverTab[99621]++

											return b.conn.Write(buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:879
	// _ = "end of CoverTab[99621]"
}

func (b *Broker) send(rb protocolBody, promiseResponse bool, responseHeaderVersion int16) (*responsePromise, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:882
	_go_fuzz_dep_.CoverTab[99624]++
											var promise *responsePromise
											if promiseResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:884
		_go_fuzz_dep_.CoverTab[99627]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:887
		promise = &responsePromise{
			headerVersion:	responseHeaderVersion,
			packets:	make(chan []byte),
			errors:		make(chan error),
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:891
		// _ = "end of CoverTab[99627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:892
		_go_fuzz_dep_.CoverTab[99628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:892
		// _ = "end of CoverTab[99628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:892
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:892
	// _ = "end of CoverTab[99624]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:892
	_go_fuzz_dep_.CoverTab[99625]++

											if err := b.sendWithPromise(rb, promise); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:894
		_go_fuzz_dep_.CoverTab[99629]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:895
		// _ = "end of CoverTab[99629]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:896
		_go_fuzz_dep_.CoverTab[99630]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:896
		// _ = "end of CoverTab[99630]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:896
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:896
	// _ = "end of CoverTab[99625]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:896
	_go_fuzz_dep_.CoverTab[99626]++

											return promise, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:898
	// _ = "end of CoverTab[99626]"
}

func (b *Broker) sendWithPromise(rb protocolBody, promise *responsePromise) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:901
	_go_fuzz_dep_.CoverTab[99631]++
											b.lock.Lock()
											defer b.lock.Unlock()

											if b.conn == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:905
		_go_fuzz_dep_.CoverTab[99637]++
												if b.connErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:906
			_go_fuzz_dep_.CoverTab[99639]++
													return b.connErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:907
			// _ = "end of CoverTab[99639]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:908
			_go_fuzz_dep_.CoverTab[99640]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:908
			// _ = "end of CoverTab[99640]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:908
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:908
		// _ = "end of CoverTab[99637]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:908
		_go_fuzz_dep_.CoverTab[99638]++
												return ErrNotConnected
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:909
		// _ = "end of CoverTab[99638]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:910
		_go_fuzz_dep_.CoverTab[99641]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:910
		// _ = "end of CoverTab[99641]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:910
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:910
	// _ = "end of CoverTab[99631]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:910
	_go_fuzz_dep_.CoverTab[99632]++

											if !b.conf.Version.IsAtLeast(rb.requiredVersion()) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:912
		_go_fuzz_dep_.CoverTab[99642]++
												return ErrUnsupportedVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:913
		// _ = "end of CoverTab[99642]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:914
		_go_fuzz_dep_.CoverTab[99643]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:914
		// _ = "end of CoverTab[99643]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:914
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:914
	// _ = "end of CoverTab[99632]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:914
	_go_fuzz_dep_.CoverTab[99633]++

											req := &request{correlationID: b.correlationID, clientID: b.conf.ClientID, body: rb}
											buf, err := encode(req, b.conf.MetricRegistry)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:918
		_go_fuzz_dep_.CoverTab[99644]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:919
		// _ = "end of CoverTab[99644]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:920
		_go_fuzz_dep_.CoverTab[99645]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:920
		// _ = "end of CoverTab[99645]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:920
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:920
	// _ = "end of CoverTab[99633]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:920
	_go_fuzz_dep_.CoverTab[99634]++

											requestTime := time.Now()

											b.addRequestInFlightMetrics(1)
											bytes, err := b.write(buf)
											b.updateOutgoingCommunicationMetrics(bytes)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:927
		_go_fuzz_dep_.CoverTab[99646]++
												b.addRequestInFlightMetrics(-1)
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:929
		// _ = "end of CoverTab[99646]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:930
		_go_fuzz_dep_.CoverTab[99647]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:930
		// _ = "end of CoverTab[99647]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:930
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:930
	// _ = "end of CoverTab[99634]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:930
	_go_fuzz_dep_.CoverTab[99635]++
											b.correlationID++

											if promise == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:933
		_go_fuzz_dep_.CoverTab[99648]++

												b.updateRequestLatencyAndInFlightMetrics(time.Since(requestTime))
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:936
		// _ = "end of CoverTab[99648]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:937
		_go_fuzz_dep_.CoverTab[99649]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:937
		// _ = "end of CoverTab[99649]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:937
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:937
	// _ = "end of CoverTab[99635]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:937
	_go_fuzz_dep_.CoverTab[99636]++

											promise.requestTime = requestTime
											promise.correlationID = req.correlationID
											b.responses <- promise

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:943
	// _ = "end of CoverTab[99636]"
}

func (b *Broker) sendAndReceive(req protocolBody, res protocolBody) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:946
	_go_fuzz_dep_.CoverTab[99650]++
											responseHeaderVersion := int16(-1)
											if res != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:948
		_go_fuzz_dep_.CoverTab[99654]++
												responseHeaderVersion = res.headerVersion()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:949
		// _ = "end of CoverTab[99654]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:950
		_go_fuzz_dep_.CoverTab[99655]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:950
		// _ = "end of CoverTab[99655]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:950
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:950
	// _ = "end of CoverTab[99650]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:950
	_go_fuzz_dep_.CoverTab[99651]++

											promise, err := b.send(req, res != nil, responseHeaderVersion)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:953
		_go_fuzz_dep_.CoverTab[99656]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:954
		// _ = "end of CoverTab[99656]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:955
		_go_fuzz_dep_.CoverTab[99657]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:955
		// _ = "end of CoverTab[99657]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:955
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:955
	// _ = "end of CoverTab[99651]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:955
	_go_fuzz_dep_.CoverTab[99652]++

											if promise == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:957
		_go_fuzz_dep_.CoverTab[99658]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:958
		// _ = "end of CoverTab[99658]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:959
		_go_fuzz_dep_.CoverTab[99659]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:959
		// _ = "end of CoverTab[99659]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:959
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:959
	// _ = "end of CoverTab[99652]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:959
	_go_fuzz_dep_.CoverTab[99653]++

											select {
	case buf := <-promise.packets:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:962
		_go_fuzz_dep_.CoverTab[99660]++
												return versionedDecode(buf, res, req.version())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:963
		// _ = "end of CoverTab[99660]"
	case err = <-promise.errors:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:964
		_go_fuzz_dep_.CoverTab[99661]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:965
		// _ = "end of CoverTab[99661]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:966
	// _ = "end of CoverTab[99653]"
}

func (b *Broker) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:969
	_go_fuzz_dep_.CoverTab[99662]++
											b.id, err = pd.getInt32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:971
		_go_fuzz_dep_.CoverTab[99668]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:972
		// _ = "end of CoverTab[99668]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:973
		_go_fuzz_dep_.CoverTab[99669]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:973
		// _ = "end of CoverTab[99669]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:973
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:973
	// _ = "end of CoverTab[99662]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:973
	_go_fuzz_dep_.CoverTab[99663]++

											host, err := pd.getString()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:976
		_go_fuzz_dep_.CoverTab[99670]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:977
		// _ = "end of CoverTab[99670]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:978
		_go_fuzz_dep_.CoverTab[99671]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:978
		// _ = "end of CoverTab[99671]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:978
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:978
	// _ = "end of CoverTab[99663]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:978
	_go_fuzz_dep_.CoverTab[99664]++

											port, err := pd.getInt32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:981
		_go_fuzz_dep_.CoverTab[99672]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:982
		// _ = "end of CoverTab[99672]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:983
		_go_fuzz_dep_.CoverTab[99673]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:983
		// _ = "end of CoverTab[99673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:983
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:983
	// _ = "end of CoverTab[99664]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:983
	_go_fuzz_dep_.CoverTab[99665]++

											if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:985
		_go_fuzz_dep_.CoverTab[99674]++
												b.rack, err = pd.getNullableString()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:987
			_go_fuzz_dep_.CoverTab[99675]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:988
			// _ = "end of CoverTab[99675]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:989
			_go_fuzz_dep_.CoverTab[99676]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:989
			// _ = "end of CoverTab[99676]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:989
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:989
		// _ = "end of CoverTab[99674]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:990
		_go_fuzz_dep_.CoverTab[99677]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:990
		// _ = "end of CoverTab[99677]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:990
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:990
	// _ = "end of CoverTab[99665]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:990
	_go_fuzz_dep_.CoverTab[99666]++

											b.addr = net.JoinHostPort(host, fmt.Sprint(port))
											if _, _, err := net.SplitHostPort(b.addr); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:993
		_go_fuzz_dep_.CoverTab[99678]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:994
		// _ = "end of CoverTab[99678]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:995
		_go_fuzz_dep_.CoverTab[99679]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:995
		// _ = "end of CoverTab[99679]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:995
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:995
	// _ = "end of CoverTab[99666]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:995
	_go_fuzz_dep_.CoverTab[99667]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:997
	// _ = "end of CoverTab[99667]"
}

func (b *Broker) encode(pe packetEncoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1000
	_go_fuzz_dep_.CoverTab[99680]++
											host, portstr, err := net.SplitHostPort(b.addr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1002
		_go_fuzz_dep_.CoverTab[99685]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1003
		// _ = "end of CoverTab[99685]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1004
		_go_fuzz_dep_.CoverTab[99686]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1004
		// _ = "end of CoverTab[99686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1004
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1004
	// _ = "end of CoverTab[99680]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1004
	_go_fuzz_dep_.CoverTab[99681]++

											port, err := strconv.ParseInt(portstr, 10, 32)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1007
		_go_fuzz_dep_.CoverTab[99687]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1008
		// _ = "end of CoverTab[99687]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1009
		_go_fuzz_dep_.CoverTab[99688]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1009
		// _ = "end of CoverTab[99688]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1009
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1009
	// _ = "end of CoverTab[99681]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1009
	_go_fuzz_dep_.CoverTab[99682]++

											pe.putInt32(b.id)

											err = pe.putString(host)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1014
		_go_fuzz_dep_.CoverTab[99689]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1015
		// _ = "end of CoverTab[99689]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1016
		_go_fuzz_dep_.CoverTab[99690]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1016
		// _ = "end of CoverTab[99690]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1016
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1016
	// _ = "end of CoverTab[99682]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1016
	_go_fuzz_dep_.CoverTab[99683]++

											pe.putInt32(int32(port))

											if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1020
		_go_fuzz_dep_.CoverTab[99691]++
												err = pe.putNullableString(b.rack)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1022
			_go_fuzz_dep_.CoverTab[99692]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1023
			// _ = "end of CoverTab[99692]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1024
			_go_fuzz_dep_.CoverTab[99693]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1024
			// _ = "end of CoverTab[99693]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1024
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1024
		// _ = "end of CoverTab[99691]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1025
		_go_fuzz_dep_.CoverTab[99694]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1025
		// _ = "end of CoverTab[99694]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1025
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1025
	// _ = "end of CoverTab[99683]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1025
	_go_fuzz_dep_.CoverTab[99684]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1027
	// _ = "end of CoverTab[99684]"
}

func (b *Broker) responseReceiver() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1030
	_go_fuzz_dep_.CoverTab[99695]++
											var dead error

											for response := range b.responses {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1033
		_go_fuzz_dep_.CoverTab[99697]++
												if dead != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1034
			_go_fuzz_dep_.CoverTab[99703]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1037
			b.addRequestInFlightMetrics(-1)
													response.handle(nil, dead)
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1039
			// _ = "end of CoverTab[99703]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1040
			_go_fuzz_dep_.CoverTab[99704]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1040
			// _ = "end of CoverTab[99704]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1040
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1040
		// _ = "end of CoverTab[99697]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1040
		_go_fuzz_dep_.CoverTab[99698]++

												headerLength := getHeaderLength(response.headerVersion)
												header := make([]byte, headerLength)

												bytesReadHeader, err := b.readFull(header)
												requestLatency := time.Since(response.requestTime)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1047
			_go_fuzz_dep_.CoverTab[99705]++
													b.updateIncomingCommunicationMetrics(bytesReadHeader, requestLatency)
													dead = err
													response.handle(nil, err)
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1051
			// _ = "end of CoverTab[99705]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1052
			_go_fuzz_dep_.CoverTab[99706]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1052
			// _ = "end of CoverTab[99706]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1052
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1052
		// _ = "end of CoverTab[99698]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1052
		_go_fuzz_dep_.CoverTab[99699]++

												decodedHeader := responseHeader{}
												err = versionedDecode(header, &decodedHeader, response.headerVersion)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1056
			_go_fuzz_dep_.CoverTab[99707]++
													b.updateIncomingCommunicationMetrics(bytesReadHeader, requestLatency)
													dead = err
													response.handle(nil, err)
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1060
			// _ = "end of CoverTab[99707]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1061
			_go_fuzz_dep_.CoverTab[99708]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1061
			// _ = "end of CoverTab[99708]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1061
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1061
		// _ = "end of CoverTab[99699]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1061
		_go_fuzz_dep_.CoverTab[99700]++
												if decodedHeader.correlationID != response.correlationID {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1062
			_go_fuzz_dep_.CoverTab[99709]++
													b.updateIncomingCommunicationMetrics(bytesReadHeader, requestLatency)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1066
			dead = PacketDecodingError{fmt.Sprintf("correlation ID didn't match, wanted %d, got %d", response.correlationID, decodedHeader.correlationID)}
													response.handle(nil, dead)
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1068
			// _ = "end of CoverTab[99709]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1069
			_go_fuzz_dep_.CoverTab[99710]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1069
			// _ = "end of CoverTab[99710]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1069
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1069
		// _ = "end of CoverTab[99700]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1069
		_go_fuzz_dep_.CoverTab[99701]++

												buf := make([]byte, decodedHeader.length-int32(headerLength)+4)
												bytesReadBody, err := b.readFull(buf)
												b.updateIncomingCommunicationMetrics(bytesReadHeader+bytesReadBody, requestLatency)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1074
			_go_fuzz_dep_.CoverTab[99711]++
													dead = err
													response.handle(nil, err)
													continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1077
			// _ = "end of CoverTab[99711]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1078
			_go_fuzz_dep_.CoverTab[99712]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1078
			// _ = "end of CoverTab[99712]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1078
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1078
		// _ = "end of CoverTab[99701]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1078
		_go_fuzz_dep_.CoverTab[99702]++

												response.handle(buf, nil)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1080
		// _ = "end of CoverTab[99702]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1081
	// _ = "end of CoverTab[99695]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1081
	_go_fuzz_dep_.CoverTab[99696]++
											close(b.done)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1082
	// _ = "end of CoverTab[99696]"
}

func getHeaderLength(headerVersion int16) int8 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1085
	_go_fuzz_dep_.CoverTab[99713]++
											if headerVersion < 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1086
		_go_fuzz_dep_.CoverTab[99714]++
												return 8
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1087
		// _ = "end of CoverTab[99714]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1088
		_go_fuzz_dep_.CoverTab[99715]++

												return 9
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1090
		// _ = "end of CoverTab[99715]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1091
	// _ = "end of CoverTab[99713]"
}

func (b *Broker) authenticateViaSASL() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1094
	_go_fuzz_dep_.CoverTab[99716]++
											switch b.conf.Net.SASL.Mechanism {
	case SASLTypeOAuth:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1096
		_go_fuzz_dep_.CoverTab[99717]++
												return b.sendAndReceiveSASLOAuth(b.conf.Net.SASL.TokenProvider)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1097
		// _ = "end of CoverTab[99717]"
	case SASLTypeSCRAMSHA256, SASLTypeSCRAMSHA512:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1098
		_go_fuzz_dep_.CoverTab[99718]++
												return b.sendAndReceiveSASLSCRAM()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1099
		// _ = "end of CoverTab[99718]"
	case SASLTypeGSSAPI:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1100
		_go_fuzz_dep_.CoverTab[99719]++
												return b.sendAndReceiveKerberos()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1101
		// _ = "end of CoverTab[99719]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1102
		_go_fuzz_dep_.CoverTab[99720]++
												return b.sendAndReceiveSASLPlainAuth()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1103
		// _ = "end of CoverTab[99720]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1104
	// _ = "end of CoverTab[99716]"
}

func (b *Broker) sendAndReceiveKerberos() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1107
	_go_fuzz_dep_.CoverTab[99721]++
											b.kerberosAuthenticator.Config = &b.conf.Net.SASL.GSSAPI
											if b.kerberosAuthenticator.NewKerberosClientFunc == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1109
		_go_fuzz_dep_.CoverTab[99723]++
												b.kerberosAuthenticator.NewKerberosClientFunc = NewKerberosClient
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1110
		// _ = "end of CoverTab[99723]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1111
		_go_fuzz_dep_.CoverTab[99724]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1111
		// _ = "end of CoverTab[99724]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1111
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1111
	// _ = "end of CoverTab[99721]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1111
	_go_fuzz_dep_.CoverTab[99722]++
											return b.kerberosAuthenticator.Authorize(b)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1112
	// _ = "end of CoverTab[99722]"
}

func (b *Broker) sendAndReceiveSASLHandshake(saslType SASLMechanism, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1115
	_go_fuzz_dep_.CoverTab[99725]++
											rb := &SaslHandshakeRequest{Mechanism: string(saslType), Version: version}

											req := &request{correlationID: b.correlationID, clientID: b.conf.ClientID, body: rb}
											buf, err := encode(req, b.conf.MetricRegistry)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1120
		_go_fuzz_dep_.CoverTab[99732]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1121
		// _ = "end of CoverTab[99732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1122
		_go_fuzz_dep_.CoverTab[99733]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1122
		// _ = "end of CoverTab[99733]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1122
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1122
	// _ = "end of CoverTab[99725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1122
	_go_fuzz_dep_.CoverTab[99726]++

											requestTime := time.Now()

											b.addRequestInFlightMetrics(1)
											bytes, err := b.write(buf)
											b.updateOutgoingCommunicationMetrics(bytes)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1129
		_go_fuzz_dep_.CoverTab[99734]++
												b.addRequestInFlightMetrics(-1)
												Logger.Printf("Failed to send SASL handshake %s: %s\n", b.addr, err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1132
		// _ = "end of CoverTab[99734]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1133
		_go_fuzz_dep_.CoverTab[99735]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1133
		// _ = "end of CoverTab[99735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1133
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1133
	// _ = "end of CoverTab[99726]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1133
	_go_fuzz_dep_.CoverTab[99727]++
											b.correlationID++

											header := make([]byte, 8)
											_, err = b.readFull(header)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1138
		_go_fuzz_dep_.CoverTab[99736]++
												b.addRequestInFlightMetrics(-1)
												Logger.Printf("Failed to read SASL handshake header : %s\n", err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1141
		// _ = "end of CoverTab[99736]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1142
		_go_fuzz_dep_.CoverTab[99737]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1142
		// _ = "end of CoverTab[99737]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1142
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1142
	// _ = "end of CoverTab[99727]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1142
	_go_fuzz_dep_.CoverTab[99728]++

											length := binary.BigEndian.Uint32(header[:4])
											payload := make([]byte, length-4)
											n, err := b.readFull(payload)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1147
		_go_fuzz_dep_.CoverTab[99738]++
												b.addRequestInFlightMetrics(-1)
												Logger.Printf("Failed to read SASL handshake payload : %s\n", err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1150
		// _ = "end of CoverTab[99738]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1151
		_go_fuzz_dep_.CoverTab[99739]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1151
		// _ = "end of CoverTab[99739]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1151
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1151
	// _ = "end of CoverTab[99728]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1151
	_go_fuzz_dep_.CoverTab[99729]++

											b.updateIncomingCommunicationMetrics(n+8, time.Since(requestTime))
											res := &SaslHandshakeResponse{}

											err = versionedDecode(payload, res, 0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1157
		_go_fuzz_dep_.CoverTab[99740]++
												Logger.Printf("Failed to parse SASL handshake : %s\n", err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1159
		// _ = "end of CoverTab[99740]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1160
		_go_fuzz_dep_.CoverTab[99741]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1160
		// _ = "end of CoverTab[99741]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1160
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1160
	// _ = "end of CoverTab[99729]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1160
	_go_fuzz_dep_.CoverTab[99730]++

											if res.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1162
		_go_fuzz_dep_.CoverTab[99742]++
												Logger.Printf("Invalid SASL Mechanism : %s\n", res.Err.Error())
												return res.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1164
		// _ = "end of CoverTab[99742]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1165
		_go_fuzz_dep_.CoverTab[99743]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1165
		// _ = "end of CoverTab[99743]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1165
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1165
	// _ = "end of CoverTab[99730]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1165
	_go_fuzz_dep_.CoverTab[99731]++

											DebugLogger.Print("Successful SASL handshake. Available mechanisms: ", res.EnabledMechanisms)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1168
	// _ = "end of CoverTab[99731]"
}

// Kafka 0.10.x supported SASL PLAIN/Kerberos via KAFKA-3149 (KIP-43).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// Kafka 1.x.x onward added a SaslAuthenticate request/response message which
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// wraps the SASL flow in the Kafka protocol, which allows for returning
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// meaningful errors on authentication failure.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// In SASL Plain, Kafka expects the auth header to be in the following format
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// Message format (from https://tools.ietf.org/html/rfc4616):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	message   = [authzid] UTF8NUL authcid UTF8NUL passwd
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	authcid   = 1*SAFE ; MUST accept up to 255 octets
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	authzid   = 1*SAFE ; MUST accept up to 255 octets
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	passwd    = 1*SAFE ; MUST accept up to 255 octets
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	UTF8NUL   = %x00 ; UTF-8 encoded NUL character
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	SAFE      = UTF1 / UTF2 / UTF3 / UTF4
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//	               ;; any UTF-8 encoded Unicode character except NUL
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// With SASL v0 handshake and auth then:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// When credentials are valid, Kafka returns a 4 byte array of null characters.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// When credentials are invalid, Kafka closes the connection.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// With SASL v1 handshake and auth then:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// When credentials are invalid, Kafka replies with a SaslAuthenticate response
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1171
// containing an error code and message detailing the authentication failure.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1195
func (b *Broker) sendAndReceiveSASLPlainAuth() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1195
	_go_fuzz_dep_.CoverTab[99744]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1198
	if b.conf.Net.SASL.Handshake {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1198
		_go_fuzz_dep_.CoverTab[99747]++
												handshakeErr := b.sendAndReceiveSASLHandshake(SASLTypePlaintext, b.conf.Net.SASL.Version)
												if handshakeErr != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1200
			_go_fuzz_dep_.CoverTab[99748]++
													Logger.Printf("Error while performing SASL handshake %s\n", b.addr)
													return handshakeErr
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1202
			// _ = "end of CoverTab[99748]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1203
			_go_fuzz_dep_.CoverTab[99749]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1203
			// _ = "end of CoverTab[99749]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1203
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1203
		// _ = "end of CoverTab[99747]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1204
		_go_fuzz_dep_.CoverTab[99750]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1204
		// _ = "end of CoverTab[99750]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1204
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1204
	// _ = "end of CoverTab[99744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1204
	_go_fuzz_dep_.CoverTab[99745]++

											if b.conf.Net.SASL.Version == SASLHandshakeV1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1206
		_go_fuzz_dep_.CoverTab[99751]++
												return b.sendAndReceiveV1SASLPlainAuth()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1207
		// _ = "end of CoverTab[99751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1208
		_go_fuzz_dep_.CoverTab[99752]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1208
		// _ = "end of CoverTab[99752]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1208
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1208
	// _ = "end of CoverTab[99745]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1208
	_go_fuzz_dep_.CoverTab[99746]++
											return b.sendAndReceiveV0SASLPlainAuth()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1209
	// _ = "end of CoverTab[99746]"
}

// sendAndReceiveV0SASLPlainAuth flows the v0 sasl auth NOT wrapped in the kafka protocol
func (b *Broker) sendAndReceiveV0SASLPlainAuth() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1213
	_go_fuzz_dep_.CoverTab[99753]++
											length := len(b.conf.Net.SASL.AuthIdentity) + 1 + len(b.conf.Net.SASL.User) + 1 + len(b.conf.Net.SASL.Password)
											authBytes := make([]byte, length+4)
											binary.BigEndian.PutUint32(authBytes, uint32(length))
											copy(authBytes[4:], b.conf.Net.SASL.AuthIdentity+"\x00"+b.conf.Net.SASL.User+"\x00"+b.conf.Net.SASL.Password)

											requestTime := time.Now()

											b.addRequestInFlightMetrics(1)
											bytesWritten, err := b.write(authBytes)
											b.updateOutgoingCommunicationMetrics(bytesWritten)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1224
		_go_fuzz_dep_.CoverTab[99756]++
												b.addRequestInFlightMetrics(-1)
												Logger.Printf("Failed to write SASL auth header to broker %s: %s\n", b.addr, err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1227
		// _ = "end of CoverTab[99756]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1228
		_go_fuzz_dep_.CoverTab[99757]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1228
		// _ = "end of CoverTab[99757]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1228
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1228
	// _ = "end of CoverTab[99753]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1228
	_go_fuzz_dep_.CoverTab[99754]++

											header := make([]byte, 4)
											n, err := b.readFull(header)
											b.updateIncomingCommunicationMetrics(n, time.Since(requestTime))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1235
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1235
		_go_fuzz_dep_.CoverTab[99758]++
												Logger.Printf("Failed to read response while authenticating with SASL to broker %s: %s\n", b.addr, err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1237
		// _ = "end of CoverTab[99758]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1238
		_go_fuzz_dep_.CoverTab[99759]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1238
		// _ = "end of CoverTab[99759]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1238
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1238
	// _ = "end of CoverTab[99754]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1238
	_go_fuzz_dep_.CoverTab[99755]++

											DebugLogger.Printf("SASL authentication successful with broker %s:%v - %v\n", b.addr, n, header)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1241
	// _ = "end of CoverTab[99755]"
}

// sendAndReceiveV1SASLPlainAuth flows the v1 sasl authentication using the kafka protocol
func (b *Broker) sendAndReceiveV1SASLPlainAuth() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1245
	_go_fuzz_dep_.CoverTab[99760]++
											correlationID := b.correlationID

											requestTime := time.Now()

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1251
	b.addRequestInFlightMetrics(1)
	bytesWritten, err := b.sendSASLPlainAuthClientResponse(correlationID)
	b.updateOutgoingCommunicationMetrics(bytesWritten)

	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1255
		_go_fuzz_dep_.CoverTab[99763]++
												b.addRequestInFlightMetrics(-1)
												Logger.Printf("Failed to write SASL auth header to broker %s: %s\n", b.addr, err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1258
		// _ = "end of CoverTab[99763]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1259
		_go_fuzz_dep_.CoverTab[99764]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1259
		// _ = "end of CoverTab[99764]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1259
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1259
	// _ = "end of CoverTab[99760]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1259
	_go_fuzz_dep_.CoverTab[99761]++

											b.correlationID++

											bytesRead, err := b.receiveSASLServerResponse(&SaslAuthenticateResponse{}, correlationID)
											b.updateIncomingCommunicationMetrics(bytesRead, time.Since(requestTime))

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1267
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1267
		_go_fuzz_dep_.CoverTab[99765]++
												Logger.Printf("Error returned from broker during SASL flow %s: %s\n", b.addr, err.Error())
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1269
		// _ = "end of CoverTab[99765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1270
		_go_fuzz_dep_.CoverTab[99766]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1270
		// _ = "end of CoverTab[99766]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1270
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1270
	// _ = "end of CoverTab[99761]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1270
	_go_fuzz_dep_.CoverTab[99762]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1272
	// _ = "end of CoverTab[99762]"
}

// sendAndReceiveSASLOAuth performs the authentication flow as described by KIP-255
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1275
// https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=75968876
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1277
func (b *Broker) sendAndReceiveSASLOAuth(provider AccessTokenProvider) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1277
	_go_fuzz_dep_.CoverTab[99767]++
											if err := b.sendAndReceiveSASLHandshake(SASLTypeOAuth, SASLHandshakeV1); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1278
		_go_fuzz_dep_.CoverTab[99773]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1279
		// _ = "end of CoverTab[99773]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1280
		_go_fuzz_dep_.CoverTab[99774]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1280
		// _ = "end of CoverTab[99774]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1280
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1280
	// _ = "end of CoverTab[99767]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1280
	_go_fuzz_dep_.CoverTab[99768]++

											token, err := provider.Token()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1283
		_go_fuzz_dep_.CoverTab[99775]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1284
		// _ = "end of CoverTab[99775]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1285
		_go_fuzz_dep_.CoverTab[99776]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1285
		// _ = "end of CoverTab[99776]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1285
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1285
	// _ = "end of CoverTab[99768]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1285
	_go_fuzz_dep_.CoverTab[99769]++

											message, err := buildClientFirstMessage(token)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1288
		_go_fuzz_dep_.CoverTab[99777]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1289
		// _ = "end of CoverTab[99777]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1290
		_go_fuzz_dep_.CoverTab[99778]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1290
		// _ = "end of CoverTab[99778]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1290
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1290
	// _ = "end of CoverTab[99769]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1290
	_go_fuzz_dep_.CoverTab[99770]++

											challenged, err := b.sendClientMessage(message)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1293
		_go_fuzz_dep_.CoverTab[99779]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1294
		// _ = "end of CoverTab[99779]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1295
		_go_fuzz_dep_.CoverTab[99780]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1295
		// _ = "end of CoverTab[99780]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1295
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1295
	// _ = "end of CoverTab[99770]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1295
	_go_fuzz_dep_.CoverTab[99771]++

											if challenged {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1297
		_go_fuzz_dep_.CoverTab[99781]++

												_, err = b.sendClientMessage([]byte(`\x01`))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1299
		// _ = "end of CoverTab[99781]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1300
		_go_fuzz_dep_.CoverTab[99782]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1300
		// _ = "end of CoverTab[99782]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1300
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1300
	// _ = "end of CoverTab[99771]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1300
	_go_fuzz_dep_.CoverTab[99772]++

											return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1302
	// _ = "end of CoverTab[99772]"
}

// sendClientMessage sends a SASL/OAUTHBEARER client message and returns true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1305
// if the broker responds with a challenge, in which case the token is
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1305
// rejected.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1308
func (b *Broker) sendClientMessage(message []byte) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1308
	_go_fuzz_dep_.CoverTab[99783]++
											requestTime := time.Now()

											b.addRequestInFlightMetrics(1)
											correlationID := b.correlationID

											bytesWritten, err := b.sendSASLOAuthBearerClientMessage(message, correlationID)
											b.updateOutgoingCommunicationMetrics(bytesWritten)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1316
		_go_fuzz_dep_.CoverTab[99786]++
												b.addRequestInFlightMetrics(-1)
												return false, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1318
		// _ = "end of CoverTab[99786]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1319
		_go_fuzz_dep_.CoverTab[99787]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1319
		// _ = "end of CoverTab[99787]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1319
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1319
	// _ = "end of CoverTab[99783]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1319
	_go_fuzz_dep_.CoverTab[99784]++

											b.correlationID++

											res := &SaslAuthenticateResponse{}
											bytesRead, err := b.receiveSASLServerResponse(res, correlationID)

											requestLatency := time.Since(requestTime)
											b.updateIncomingCommunicationMetrics(bytesRead, requestLatency)

											isChallenge := len(res.SaslAuthBytes) > 0

											if isChallenge && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1331
		_go_fuzz_dep_.CoverTab[99788]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1331
		return err != nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1331
		// _ = "end of CoverTab[99788]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1331
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1331
		_go_fuzz_dep_.CoverTab[99789]++
												Logger.Printf("Broker rejected authentication token: %s", res.SaslAuthBytes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1332
		// _ = "end of CoverTab[99789]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1333
		_go_fuzz_dep_.CoverTab[99790]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1333
		// _ = "end of CoverTab[99790]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1333
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1333
	// _ = "end of CoverTab[99784]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1333
	_go_fuzz_dep_.CoverTab[99785]++

											return isChallenge, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1335
	// _ = "end of CoverTab[99785]"
}

func (b *Broker) sendAndReceiveSASLSCRAM() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1338
	_go_fuzz_dep_.CoverTab[99791]++
											if b.conf.Net.SASL.Version == SASLHandshakeV0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1339
		_go_fuzz_dep_.CoverTab[99793]++
												return b.sendAndReceiveSASLSCRAMv0()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1340
		// _ = "end of CoverTab[99793]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1341
		_go_fuzz_dep_.CoverTab[99794]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1341
		// _ = "end of CoverTab[99794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1341
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1341
	// _ = "end of CoverTab[99791]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1341
	_go_fuzz_dep_.CoverTab[99792]++
											return b.sendAndReceiveSASLSCRAMv1()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1342
	// _ = "end of CoverTab[99792]"
}

func (b *Broker) sendAndReceiveSASLSCRAMv0() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1345
	_go_fuzz_dep_.CoverTab[99795]++
											if err := b.sendAndReceiveSASLHandshake(b.conf.Net.SASL.Mechanism, SASLHandshakeV0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1346
		_go_fuzz_dep_.CoverTab[99800]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1347
		// _ = "end of CoverTab[99800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1348
		_go_fuzz_dep_.CoverTab[99801]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1348
		// _ = "end of CoverTab[99801]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1348
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1348
	// _ = "end of CoverTab[99795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1348
	_go_fuzz_dep_.CoverTab[99796]++

											scramClient := b.conf.Net.SASL.SCRAMClientGeneratorFunc()
											if err := scramClient.Begin(b.conf.Net.SASL.User, b.conf.Net.SASL.Password, b.conf.Net.SASL.SCRAMAuthzID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1351
		_go_fuzz_dep_.CoverTab[99802]++
												return fmt.Errorf("failed to start SCRAM exchange with the server: %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1352
		// _ = "end of CoverTab[99802]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1353
		_go_fuzz_dep_.CoverTab[99803]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1353
		// _ = "end of CoverTab[99803]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1353
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1353
	// _ = "end of CoverTab[99796]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1353
	_go_fuzz_dep_.CoverTab[99797]++

											msg, err := scramClient.Step("")
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1356
		_go_fuzz_dep_.CoverTab[99804]++
												return fmt.Errorf("failed to advance the SCRAM exchange: %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1357
		// _ = "end of CoverTab[99804]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1358
		_go_fuzz_dep_.CoverTab[99805]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1358
		// _ = "end of CoverTab[99805]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1358
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1358
	// _ = "end of CoverTab[99797]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1358
	_go_fuzz_dep_.CoverTab[99798]++

											for !scramClient.Done() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1360
		_go_fuzz_dep_.CoverTab[99806]++
												requestTime := time.Now()

												b.addRequestInFlightMetrics(1)
												length := len(msg)
												authBytes := make([]byte, length+4)
												binary.BigEndian.PutUint32(authBytes, uint32(length))
												copy(authBytes[4:], []byte(msg))
												_, err := b.write(authBytes)
												b.updateOutgoingCommunicationMetrics(length + 4)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1370
			_go_fuzz_dep_.CoverTab[99810]++
													b.addRequestInFlightMetrics(-1)
													Logger.Printf("Failed to write SASL auth header to broker %s: %s\n", b.addr, err.Error())
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1373
			// _ = "end of CoverTab[99810]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1374
			_go_fuzz_dep_.CoverTab[99811]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1374
			// _ = "end of CoverTab[99811]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1374
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1374
		// _ = "end of CoverTab[99806]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1374
		_go_fuzz_dep_.CoverTab[99807]++
												b.correlationID++
												header := make([]byte, 4)
												_, err = b.readFull(header)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1378
			_go_fuzz_dep_.CoverTab[99812]++
													b.addRequestInFlightMetrics(-1)
													Logger.Printf("Failed to read response header while authenticating with SASL to broker %s: %s\n", b.addr, err.Error())
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1381
			// _ = "end of CoverTab[99812]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1382
			_go_fuzz_dep_.CoverTab[99813]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1382
			// _ = "end of CoverTab[99813]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1382
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1382
		// _ = "end of CoverTab[99807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1382
		_go_fuzz_dep_.CoverTab[99808]++
												payload := make([]byte, int32(binary.BigEndian.Uint32(header)))
												n, err := b.readFull(payload)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1385
			_go_fuzz_dep_.CoverTab[99814]++
													b.addRequestInFlightMetrics(-1)
													Logger.Printf("Failed to read response payload while authenticating with SASL to broker %s: %s\n", b.addr, err.Error())
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1388
			// _ = "end of CoverTab[99814]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1389
			_go_fuzz_dep_.CoverTab[99815]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1389
			// _ = "end of CoverTab[99815]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1389
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1389
		// _ = "end of CoverTab[99808]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1389
		_go_fuzz_dep_.CoverTab[99809]++
												b.updateIncomingCommunicationMetrics(n+4, time.Since(requestTime))
												msg, err = scramClient.Step(string(payload))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1392
			_go_fuzz_dep_.CoverTab[99816]++
													Logger.Println("SASL authentication failed", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1394
			// _ = "end of CoverTab[99816]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1395
			_go_fuzz_dep_.CoverTab[99817]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1395
			// _ = "end of CoverTab[99817]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1395
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1395
		// _ = "end of CoverTab[99809]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1396
	// _ = "end of CoverTab[99798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1396
	_go_fuzz_dep_.CoverTab[99799]++

											DebugLogger.Println("SASL authentication succeeded")
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1399
	// _ = "end of CoverTab[99799]"
}

func (b *Broker) sendAndReceiveSASLSCRAMv1() error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1402
	_go_fuzz_dep_.CoverTab[99818]++
											if err := b.sendAndReceiveSASLHandshake(b.conf.Net.SASL.Mechanism, SASLHandshakeV1); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1403
		_go_fuzz_dep_.CoverTab[99823]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1404
		// _ = "end of CoverTab[99823]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1405
		_go_fuzz_dep_.CoverTab[99824]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1405
		// _ = "end of CoverTab[99824]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1405
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1405
	// _ = "end of CoverTab[99818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1405
	_go_fuzz_dep_.CoverTab[99819]++

											scramClient := b.conf.Net.SASL.SCRAMClientGeneratorFunc()
											if err := scramClient.Begin(b.conf.Net.SASL.User, b.conf.Net.SASL.Password, b.conf.Net.SASL.SCRAMAuthzID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1408
		_go_fuzz_dep_.CoverTab[99825]++
												return fmt.Errorf("failed to start SCRAM exchange with the server: %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1409
		// _ = "end of CoverTab[99825]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1410
		_go_fuzz_dep_.CoverTab[99826]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1410
		// _ = "end of CoverTab[99826]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1410
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1410
	// _ = "end of CoverTab[99819]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1410
	_go_fuzz_dep_.CoverTab[99820]++

											msg, err := scramClient.Step("")
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1413
		_go_fuzz_dep_.CoverTab[99827]++
												return fmt.Errorf("failed to advance the SCRAM exchange: %s", err.Error())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1414
		// _ = "end of CoverTab[99827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1415
		_go_fuzz_dep_.CoverTab[99828]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1415
		// _ = "end of CoverTab[99828]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1415
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1415
	// _ = "end of CoverTab[99820]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1415
	_go_fuzz_dep_.CoverTab[99821]++

											for !scramClient.Done() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1417
		_go_fuzz_dep_.CoverTab[99829]++
												requestTime := time.Now()

												b.addRequestInFlightMetrics(1)
												correlationID := b.correlationID
												bytesWritten, err := b.sendSaslAuthenticateRequest(correlationID, []byte(msg))
												b.updateOutgoingCommunicationMetrics(bytesWritten)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1424
			_go_fuzz_dep_.CoverTab[99832]++
													b.addRequestInFlightMetrics(-1)
													Logger.Printf("Failed to write SASL auth header to broker %s: %s\n", b.addr, err.Error())
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1427
			// _ = "end of CoverTab[99832]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1428
			_go_fuzz_dep_.CoverTab[99833]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1428
			// _ = "end of CoverTab[99833]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1428
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1428
		// _ = "end of CoverTab[99829]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1428
		_go_fuzz_dep_.CoverTab[99830]++

												b.correlationID++
												challenge, err := b.receiveSaslAuthenticateResponse(correlationID)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1432
			_go_fuzz_dep_.CoverTab[99834]++
													b.addRequestInFlightMetrics(-1)
													Logger.Printf("Failed to read response while authenticating with SASL to broker %s: %s\n", b.addr, err.Error())
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1435
			// _ = "end of CoverTab[99834]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1436
			_go_fuzz_dep_.CoverTab[99835]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1436
			// _ = "end of CoverTab[99835]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1436
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1436
		// _ = "end of CoverTab[99830]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1436
		_go_fuzz_dep_.CoverTab[99831]++

												b.updateIncomingCommunicationMetrics(len(challenge), time.Since(requestTime))
												msg, err = scramClient.Step(string(challenge))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1440
			_go_fuzz_dep_.CoverTab[99836]++
													Logger.Println("SASL authentication failed", err)
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1442
			// _ = "end of CoverTab[99836]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1443
			_go_fuzz_dep_.CoverTab[99837]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1443
			// _ = "end of CoverTab[99837]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1443
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1443
		// _ = "end of CoverTab[99831]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1444
	// _ = "end of CoverTab[99821]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1444
	_go_fuzz_dep_.CoverTab[99822]++

											DebugLogger.Println("SASL authentication succeeded")
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1447
	// _ = "end of CoverTab[99822]"
}

func (b *Broker) sendSaslAuthenticateRequest(correlationID int32, msg []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1450
	_go_fuzz_dep_.CoverTab[99838]++
											rb := &SaslAuthenticateRequest{msg}
											req := &request{correlationID: correlationID, clientID: b.conf.ClientID, body: rb}
											buf, err := encode(req, b.conf.MetricRegistry)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1454
		_go_fuzz_dep_.CoverTab[99840]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1455
		// _ = "end of CoverTab[99840]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1456
		_go_fuzz_dep_.CoverTab[99841]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1456
		// _ = "end of CoverTab[99841]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1456
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1456
	// _ = "end of CoverTab[99838]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1456
	_go_fuzz_dep_.CoverTab[99839]++

											return b.write(buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1458
	// _ = "end of CoverTab[99839]"
}

func (b *Broker) receiveSaslAuthenticateResponse(correlationID int32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1461
	_go_fuzz_dep_.CoverTab[99842]++
											buf := make([]byte, responseLengthSize+correlationIDSize)
											_, err := b.readFull(buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1464
		_go_fuzz_dep_.CoverTab[99849]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1465
		// _ = "end of CoverTab[99849]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1466
		_go_fuzz_dep_.CoverTab[99850]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1466
		// _ = "end of CoverTab[99850]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1466
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1466
	// _ = "end of CoverTab[99842]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1466
	_go_fuzz_dep_.CoverTab[99843]++

											header := responseHeader{}
											err = versionedDecode(buf, &header, 0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1470
		_go_fuzz_dep_.CoverTab[99851]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1471
		// _ = "end of CoverTab[99851]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1472
		_go_fuzz_dep_.CoverTab[99852]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1472
		// _ = "end of CoverTab[99852]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1472
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1472
	// _ = "end of CoverTab[99843]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1472
	_go_fuzz_dep_.CoverTab[99844]++

											if header.correlationID != correlationID {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1474
		_go_fuzz_dep_.CoverTab[99853]++
												return nil, fmt.Errorf("correlation ID didn't match, wanted %d, got %d", b.correlationID, header.correlationID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1475
		// _ = "end of CoverTab[99853]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1476
		_go_fuzz_dep_.CoverTab[99854]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1476
		// _ = "end of CoverTab[99854]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1476
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1476
	// _ = "end of CoverTab[99844]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1476
	_go_fuzz_dep_.CoverTab[99845]++

											buf = make([]byte, header.length-correlationIDSize)
											_, err = b.readFull(buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1480
		_go_fuzz_dep_.CoverTab[99855]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1481
		// _ = "end of CoverTab[99855]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1482
		_go_fuzz_dep_.CoverTab[99856]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1482
		// _ = "end of CoverTab[99856]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1482
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1482
	// _ = "end of CoverTab[99845]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1482
	_go_fuzz_dep_.CoverTab[99846]++

											res := &SaslAuthenticateResponse{}
											if err := versionedDecode(buf, res, 0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1485
		_go_fuzz_dep_.CoverTab[99857]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1486
		// _ = "end of CoverTab[99857]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1487
		_go_fuzz_dep_.CoverTab[99858]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1487
		// _ = "end of CoverTab[99858]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1487
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1487
	// _ = "end of CoverTab[99846]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1487
	_go_fuzz_dep_.CoverTab[99847]++
											if res.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1488
		_go_fuzz_dep_.CoverTab[99859]++
												return nil, res.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1489
		// _ = "end of CoverTab[99859]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1490
		_go_fuzz_dep_.CoverTab[99860]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1490
		// _ = "end of CoverTab[99860]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1490
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1490
	// _ = "end of CoverTab[99847]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1490
	_go_fuzz_dep_.CoverTab[99848]++
											return res.SaslAuthBytes, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1491
	// _ = "end of CoverTab[99848]"
}

// Build SASL/OAUTHBEARER initial client response as described by RFC-7628
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1494
// https://tools.ietf.org/html/rfc7628
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1496
func buildClientFirstMessage(token *AccessToken) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1496
	_go_fuzz_dep_.CoverTab[99861]++
											var ext string

											if token.Extensions != nil && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1499
		_go_fuzz_dep_.CoverTab[99863]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1499
		return len(token.Extensions) > 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1499
		// _ = "end of CoverTab[99863]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1499
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1499
		_go_fuzz_dep_.CoverTab[99864]++
												if _, ok := token.Extensions[SASLExtKeyAuth]; ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1500
			_go_fuzz_dep_.CoverTab[99866]++
													return []byte{}, fmt.Errorf("the extension `%s` is invalid", SASLExtKeyAuth)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1501
			// _ = "end of CoverTab[99866]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1502
			_go_fuzz_dep_.CoverTab[99867]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1502
			// _ = "end of CoverTab[99867]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1502
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1502
		// _ = "end of CoverTab[99864]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1502
		_go_fuzz_dep_.CoverTab[99865]++
												ext = "\x01" + mapToString(token.Extensions, "=", "\x01")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1503
		// _ = "end of CoverTab[99865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1504
		_go_fuzz_dep_.CoverTab[99868]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1504
		// _ = "end of CoverTab[99868]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1504
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1504
	// _ = "end of CoverTab[99861]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1504
	_go_fuzz_dep_.CoverTab[99862]++

											resp := []byte(fmt.Sprintf("n,,\x01auth=Bearer %s%s\x01\x01", token.Token, ext))

											return resp, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1508
	// _ = "end of CoverTab[99862]"
}

// mapToString returns a list of key-value pairs ordered by key.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1511
// keyValSep separates the key from the value. elemSep separates each pair.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1513
func mapToString(extensions map[string]string, keyValSep string, elemSep string) string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1513
	_go_fuzz_dep_.CoverTab[99869]++
											buf := make([]string, 0, len(extensions))

											for k, v := range extensions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1516
		_go_fuzz_dep_.CoverTab[99871]++
												buf = append(buf, k+keyValSep+v)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1517
		// _ = "end of CoverTab[99871]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1518
	// _ = "end of CoverTab[99869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1518
	_go_fuzz_dep_.CoverTab[99870]++

											sort.Strings(buf)

											return strings.Join(buf, elemSep)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1522
	// _ = "end of CoverTab[99870]"
}

func (b *Broker) sendSASLPlainAuthClientResponse(correlationID int32) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1525
	_go_fuzz_dep_.CoverTab[99872]++
											authBytes := []byte(b.conf.Net.SASL.AuthIdentity + "\x00" + b.conf.Net.SASL.User + "\x00" + b.conf.Net.SASL.Password)
											rb := &SaslAuthenticateRequest{authBytes}
											req := &request{correlationID: correlationID, clientID: b.conf.ClientID, body: rb}
											buf, err := encode(req, b.conf.MetricRegistry)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1530
		_go_fuzz_dep_.CoverTab[99874]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1531
		// _ = "end of CoverTab[99874]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1532
		_go_fuzz_dep_.CoverTab[99875]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1532
		// _ = "end of CoverTab[99875]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1532
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1532
	// _ = "end of CoverTab[99872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1532
	_go_fuzz_dep_.CoverTab[99873]++

											return b.write(buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1534
	// _ = "end of CoverTab[99873]"
}

func (b *Broker) sendSASLOAuthBearerClientMessage(initialResp []byte, correlationID int32) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1537
	_go_fuzz_dep_.CoverTab[99876]++
											rb := &SaslAuthenticateRequest{initialResp}

											req := &request{correlationID: correlationID, clientID: b.conf.ClientID, body: rb}

											buf, err := encode(req, b.conf.MetricRegistry)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1543
		_go_fuzz_dep_.CoverTab[99878]++
												return 0, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1544
		// _ = "end of CoverTab[99878]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1545
		_go_fuzz_dep_.CoverTab[99879]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1545
		// _ = "end of CoverTab[99879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1545
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1545
	// _ = "end of CoverTab[99876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1545
	_go_fuzz_dep_.CoverTab[99877]++

											return b.write(buf)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1547
	// _ = "end of CoverTab[99877]"
}

func (b *Broker) receiveSASLServerResponse(res *SaslAuthenticateResponse, correlationID int32) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1550
	_go_fuzz_dep_.CoverTab[99880]++
											buf := make([]byte, responseLengthSize+correlationIDSize)
											bytesRead, err := b.readFull(buf)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1553
		_go_fuzz_dep_.CoverTab[99887]++
												return bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1554
		// _ = "end of CoverTab[99887]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1555
		_go_fuzz_dep_.CoverTab[99888]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1555
		// _ = "end of CoverTab[99888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1555
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1555
	// _ = "end of CoverTab[99880]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1555
	_go_fuzz_dep_.CoverTab[99881]++

											header := responseHeader{}
											err = versionedDecode(buf, &header, 0)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1559
		_go_fuzz_dep_.CoverTab[99889]++
												return bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1560
		// _ = "end of CoverTab[99889]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1561
		_go_fuzz_dep_.CoverTab[99890]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1561
		// _ = "end of CoverTab[99890]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1561
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1561
	// _ = "end of CoverTab[99881]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1561
	_go_fuzz_dep_.CoverTab[99882]++

											if header.correlationID != correlationID {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1563
		_go_fuzz_dep_.CoverTab[99891]++
												return bytesRead, fmt.Errorf("correlation ID didn't match, wanted %d, got %d", b.correlationID, header.correlationID)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1564
		// _ = "end of CoverTab[99891]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1565
		_go_fuzz_dep_.CoverTab[99892]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1565
		// _ = "end of CoverTab[99892]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1565
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1565
	// _ = "end of CoverTab[99882]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1565
	_go_fuzz_dep_.CoverTab[99883]++

											buf = make([]byte, header.length-correlationIDSize)
											c, err := b.readFull(buf)
											bytesRead += c
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1570
		_go_fuzz_dep_.CoverTab[99893]++
												return bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1571
		// _ = "end of CoverTab[99893]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1572
		_go_fuzz_dep_.CoverTab[99894]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1572
		// _ = "end of CoverTab[99894]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1572
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1572
	// _ = "end of CoverTab[99883]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1572
	_go_fuzz_dep_.CoverTab[99884]++

											if err := versionedDecode(buf, res, 0); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1574
		_go_fuzz_dep_.CoverTab[99895]++
												return bytesRead, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1575
		// _ = "end of CoverTab[99895]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1576
		_go_fuzz_dep_.CoverTab[99896]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1576
		// _ = "end of CoverTab[99896]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1576
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1576
	// _ = "end of CoverTab[99884]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1576
	_go_fuzz_dep_.CoverTab[99885]++

											if res.Err != ErrNoError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1578
		_go_fuzz_dep_.CoverTab[99897]++
												return bytesRead, res.Err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1579
		// _ = "end of CoverTab[99897]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1580
		_go_fuzz_dep_.CoverTab[99898]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1580
		// _ = "end of CoverTab[99898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1580
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1580
	// _ = "end of CoverTab[99885]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1580
	_go_fuzz_dep_.CoverTab[99886]++

											return bytesRead, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1582
	// _ = "end of CoverTab[99886]"
}

func (b *Broker) updateIncomingCommunicationMetrics(bytes int, requestLatency time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1585
	_go_fuzz_dep_.CoverTab[99899]++
											b.updateRequestLatencyAndInFlightMetrics(requestLatency)
											b.responseRate.Mark(1)

											if b.brokerResponseRate != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1589
		_go_fuzz_dep_.CoverTab[99902]++
												b.brokerResponseRate.Mark(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1590
		// _ = "end of CoverTab[99902]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1591
		_go_fuzz_dep_.CoverTab[99903]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1591
		// _ = "end of CoverTab[99903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1591
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1591
	// _ = "end of CoverTab[99899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1591
	_go_fuzz_dep_.CoverTab[99900]++

											responseSize := int64(bytes)
											b.incomingByteRate.Mark(responseSize)
											if b.brokerIncomingByteRate != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1595
		_go_fuzz_dep_.CoverTab[99904]++
												b.brokerIncomingByteRate.Mark(responseSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1596
		// _ = "end of CoverTab[99904]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1597
		_go_fuzz_dep_.CoverTab[99905]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1597
		// _ = "end of CoverTab[99905]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1597
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1597
	// _ = "end of CoverTab[99900]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1597
	_go_fuzz_dep_.CoverTab[99901]++

											b.responseSize.Update(responseSize)
											if b.brokerResponseSize != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1600
		_go_fuzz_dep_.CoverTab[99906]++
												b.brokerResponseSize.Update(responseSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1601
		// _ = "end of CoverTab[99906]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1602
		_go_fuzz_dep_.CoverTab[99907]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1602
		// _ = "end of CoverTab[99907]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1602
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1602
	// _ = "end of CoverTab[99901]"
}

func (b *Broker) updateRequestLatencyAndInFlightMetrics(requestLatency time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1605
	_go_fuzz_dep_.CoverTab[99908]++
											requestLatencyInMs := int64(requestLatency / time.Millisecond)
											b.requestLatency.Update(requestLatencyInMs)

											if b.brokerRequestLatency != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1609
		_go_fuzz_dep_.CoverTab[99910]++
												b.brokerRequestLatency.Update(requestLatencyInMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1610
		// _ = "end of CoverTab[99910]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1611
		_go_fuzz_dep_.CoverTab[99911]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1611
		// _ = "end of CoverTab[99911]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1611
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1611
	// _ = "end of CoverTab[99908]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1611
	_go_fuzz_dep_.CoverTab[99909]++

											b.addRequestInFlightMetrics(-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1613
	// _ = "end of CoverTab[99909]"
}

func (b *Broker) addRequestInFlightMetrics(i int64) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1616
	_go_fuzz_dep_.CoverTab[99912]++
											b.requestsInFlight.Inc(i)
											if b.brokerRequestsInFlight != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1618
		_go_fuzz_dep_.CoverTab[99913]++
												b.brokerRequestsInFlight.Inc(i)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1619
		// _ = "end of CoverTab[99913]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1620
		_go_fuzz_dep_.CoverTab[99914]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1620
		// _ = "end of CoverTab[99914]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1620
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1620
	// _ = "end of CoverTab[99912]"
}

func (b *Broker) updateOutgoingCommunicationMetrics(bytes int) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1623
	_go_fuzz_dep_.CoverTab[99915]++
											b.requestRate.Mark(1)
											if b.brokerRequestRate != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1625
		_go_fuzz_dep_.CoverTab[99918]++
												b.brokerRequestRate.Mark(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1626
		// _ = "end of CoverTab[99918]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1627
		_go_fuzz_dep_.CoverTab[99919]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1627
		// _ = "end of CoverTab[99919]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1627
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1627
	// _ = "end of CoverTab[99915]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1627
	_go_fuzz_dep_.CoverTab[99916]++

											requestSize := int64(bytes)
											b.outgoingByteRate.Mark(requestSize)
											if b.brokerOutgoingByteRate != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1631
		_go_fuzz_dep_.CoverTab[99920]++
												b.brokerOutgoingByteRate.Mark(requestSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1632
		// _ = "end of CoverTab[99920]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1633
		_go_fuzz_dep_.CoverTab[99921]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1633
		// _ = "end of CoverTab[99921]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1633
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1633
	// _ = "end of CoverTab[99916]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1633
	_go_fuzz_dep_.CoverTab[99917]++

											b.requestSize.Update(requestSize)
											if b.brokerRequestSize != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1636
		_go_fuzz_dep_.CoverTab[99922]++
												b.brokerRequestSize.Update(requestSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1637
		// _ = "end of CoverTab[99922]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1638
		_go_fuzz_dep_.CoverTab[99923]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1638
		// _ = "end of CoverTab[99923]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1638
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1638
	// _ = "end of CoverTab[99917]"
}

func (b *Broker) updateThrottleMetric(throttleTime time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1641
	_go_fuzz_dep_.CoverTab[99924]++
											if throttleTime != time.Duration(0) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1642
		_go_fuzz_dep_.CoverTab[99925]++
												DebugLogger.Printf(
			"producer/broker/%d ProduceResponse throttled %v\n",
			b.ID(), throttleTime)
		if b.brokerThrottleTime != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1646
			_go_fuzz_dep_.CoverTab[99926]++
													throttleTimeInMs := int64(throttleTime / time.Millisecond)
													b.brokerThrottleTime.Update(throttleTimeInMs)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1648
			// _ = "end of CoverTab[99926]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1649
			_go_fuzz_dep_.CoverTab[99927]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1649
			// _ = "end of CoverTab[99927]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1649
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1649
		// _ = "end of CoverTab[99925]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1650
		_go_fuzz_dep_.CoverTab[99928]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1650
		// _ = "end of CoverTab[99928]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1650
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1650
	// _ = "end of CoverTab[99924]"
}

func (b *Broker) registerMetrics() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1653
	_go_fuzz_dep_.CoverTab[99929]++
											b.brokerIncomingByteRate = b.registerMeter("incoming-byte-rate")
											b.brokerRequestRate = b.registerMeter("request-rate")
											b.brokerRequestSize = b.registerHistogram("request-size")
											b.brokerRequestLatency = b.registerHistogram("request-latency-in-ms")
											b.brokerOutgoingByteRate = b.registerMeter("outgoing-byte-rate")
											b.brokerResponseRate = b.registerMeter("response-rate")
											b.brokerResponseSize = b.registerHistogram("response-size")
											b.brokerRequestsInFlight = b.registerCounter("requests-in-flight")
											b.brokerThrottleTime = b.registerHistogram("throttle-time-in-ms")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1662
	// _ = "end of CoverTab[99929]"
}

func (b *Broker) unregisterMetrics() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1665
	_go_fuzz_dep_.CoverTab[99930]++
											for _, name := range b.registeredMetrics {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1666
		_go_fuzz_dep_.CoverTab[99932]++
												b.conf.MetricRegistry.Unregister(name)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1667
		// _ = "end of CoverTab[99932]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1668
	// _ = "end of CoverTab[99930]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1668
	_go_fuzz_dep_.CoverTab[99931]++
											b.registeredMetrics = nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1669
	// _ = "end of CoverTab[99931]"
}

func (b *Broker) registerMeter(name string) metrics.Meter {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1672
	_go_fuzz_dep_.CoverTab[99933]++
											nameForBroker := getMetricNameForBroker(name, b)
											b.registeredMetrics = append(b.registeredMetrics, nameForBroker)
											return metrics.GetOrRegisterMeter(nameForBroker, b.conf.MetricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1675
	// _ = "end of CoverTab[99933]"
}

func (b *Broker) registerHistogram(name string) metrics.Histogram {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1678
	_go_fuzz_dep_.CoverTab[99934]++
											nameForBroker := getMetricNameForBroker(name, b)
											b.registeredMetrics = append(b.registeredMetrics, nameForBroker)
											return getOrRegisterHistogram(nameForBroker, b.conf.MetricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1681
	// _ = "end of CoverTab[99934]"
}

func (b *Broker) registerCounter(name string) metrics.Counter {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1684
	_go_fuzz_dep_.CoverTab[99935]++
											nameForBroker := getMetricNameForBroker(name, b)
											b.registeredMetrics = append(b.registeredMetrics, nameForBroker)
											return metrics.GetOrRegisterCounter(nameForBroker, b.conf.MetricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1687
	// _ = "end of CoverTab[99935]"
}

func validServerNameTLS(addr string, cfg *tls.Config) *tls.Config {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1690
	_go_fuzz_dep_.CoverTab[99936]++
											if cfg == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1691
		_go_fuzz_dep_.CoverTab[99940]++
												cfg = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1694
		// _ = "end of CoverTab[99940]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1695
		_go_fuzz_dep_.CoverTab[99941]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1695
		// _ = "end of CoverTab[99941]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1695
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1695
	// _ = "end of CoverTab[99936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1695
	_go_fuzz_dep_.CoverTab[99937]++
											if cfg.ServerName != "" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1696
		_go_fuzz_dep_.CoverTab[99942]++
												return cfg
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1697
		// _ = "end of CoverTab[99942]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1698
		_go_fuzz_dep_.CoverTab[99943]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1698
		// _ = "end of CoverTab[99943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1698
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1698
	// _ = "end of CoverTab[99937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1698
	_go_fuzz_dep_.CoverTab[99938]++

											c := cfg.Clone()
											sn, _, err := net.SplitHostPort(addr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1702
		_go_fuzz_dep_.CoverTab[99944]++
												Logger.Println(fmt.Errorf("failed to get ServerName from addr %w", err))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1703
		// _ = "end of CoverTab[99944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1704
		_go_fuzz_dep_.CoverTab[99945]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1704
		// _ = "end of CoverTab[99945]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1704
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1704
	// _ = "end of CoverTab[99938]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1704
	_go_fuzz_dep_.CoverTab[99939]++
											c.ServerName = sn
											return c
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1706
	// _ = "end of CoverTab[99939]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1707
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/broker.go:1707
var _ = _go_fuzz_dep_.CoverTab
