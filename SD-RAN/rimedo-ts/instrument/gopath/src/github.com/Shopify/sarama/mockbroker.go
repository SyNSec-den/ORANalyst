//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:1
)

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	expectationTimeout = 500 * time.Millisecond
)

type GSSApiHandlerFunc func([]byte) []byte

type requestHandlerFunc func(req *request) (res encoderWithHeader)

// RequestNotifierFunc is invoked when a mock broker processes a request successfully
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:25
// and will provides the number of bytes read and written.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:27
type RequestNotifierFunc func(bytesRead, bytesWritten int)

// MockBroker is a mock Kafka broker that is used in unit tests. It is exposed
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// to facilitate testing of higher level or specialized consumers and producers
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// built on top of Sarama. Note that it does not 'mimic' the Kafka API protocol,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// but rather provides a facility to do that. It takes care of the TCP
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// transport, request unmarshalling, response marshaling, and makes it the test
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// writer responsibility to program correct according to the Kafka API protocol
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// MockBroker behavior.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// MockBroker is implemented as a TCP server listening on a kernel-selected
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// localhost port that can accept many connections. It reads Kafka requests
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// from that connection and returns responses programmed by the SetHandlerByMap
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// function. If a MockBroker receives a request that it has no programmed
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// response for, then it returns nothing and the request times out.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// A set of MockRequest builders to define mappings used by MockBroker is
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// provided by Sarama. But users can develop MockRequests of their own and use
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// them along with or instead of the standard ones.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// When running tests with MockBroker it is strongly recommended to specify
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// a timeout to `go test` so that if the broker hangs waiting for a response,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// the test panics.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
//
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// It is not necessary to prefix message length or correlation ID to your
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:29
// response bytes, the server does that automatically as a convenience.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:53
type MockBroker struct {
	brokerID	int32
	port		int32
	closing		chan none
	stopper		chan none
	expectations	chan encoderWithHeader
	listener	net.Listener
	t		TestReporter
	latency		time.Duration
	handler		requestHandlerFunc
	notifier	RequestNotifierFunc
	history		[]RequestResponse
	lock		sync.Mutex
	gssApiHandler	GSSApiHandlerFunc
}

// RequestResponse represents a Request/Response pair processed by MockBroker.
type RequestResponse struct {
	Request		protocolBody
	Response	encoder
}

// SetLatency makes broker pause for the specified period every time before
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:75
// replying.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:77
func (b *MockBroker) SetLatency(latency time.Duration) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:77
	_go_fuzz_dep_.CoverTab[104245]++
											b.latency = latency
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:78
	// _ = "end of CoverTab[104245]"
}

// SetHandlerByMap defines mapping of Request types to MockResponses. When a
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:81
// request is received by the broker, it looks up the request type in the map
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:81
// and uses the found MockResponse instance to generate an appropriate reply.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:81
// If the request type is not found in the map then nothing is sent.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:85
func (b *MockBroker) SetHandlerByMap(handlerMap map[string]MockResponse) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:85
	_go_fuzz_dep_.CoverTab[104246]++
											fnMap := make(map[string]MockResponse)
											for k, v := range handlerMap {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:87
		_go_fuzz_dep_.CoverTab[104248]++
												fnMap[k] = v
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:88
		// _ = "end of CoverTab[104248]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:89
	// _ = "end of CoverTab[104246]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:89
	_go_fuzz_dep_.CoverTab[104247]++
											b.setHandler(func(req *request) (res encoderWithHeader) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:90
		_go_fuzz_dep_.CoverTab[104249]++
												reqTypeName := reflect.TypeOf(req.body).Elem().Name()
												mockResponse := fnMap[reqTypeName]
												if mockResponse == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:93
			_go_fuzz_dep_.CoverTab[104251]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:94
			// _ = "end of CoverTab[104251]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:95
			_go_fuzz_dep_.CoverTab[104252]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:95
			// _ = "end of CoverTab[104252]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:95
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:95
		// _ = "end of CoverTab[104249]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:95
		_go_fuzz_dep_.CoverTab[104250]++
												return mockResponse.For(req.body)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:96
		// _ = "end of CoverTab[104250]"
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:97
	// _ = "end of CoverTab[104247]"
}

// SetNotifier set a function that will get invoked whenever a request has been
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:100
// processed successfully and will provide the number of bytes read and written
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:102
func (b *MockBroker) SetNotifier(notifier RequestNotifierFunc) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:102
	_go_fuzz_dep_.CoverTab[104253]++
											b.lock.Lock()
											b.notifier = notifier
											b.lock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:105
	// _ = "end of CoverTab[104253]"
}

// BrokerID returns broker ID assigned to the broker.
func (b *MockBroker) BrokerID() int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:109
	_go_fuzz_dep_.CoverTab[104254]++
											return b.brokerID
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:110
	// _ = "end of CoverTab[104254]"
}

// History returns a slice of RequestResponse pairs in the order they were
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:113
// processed by the broker. Note that in case of multiple connections to the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:113
// broker the order expected by a test can be different from the order recorded
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:113
// in the history, unless some synchronization is implemented in the test.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:117
func (b *MockBroker) History() []RequestResponse {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:117
	_go_fuzz_dep_.CoverTab[104255]++
											b.lock.Lock()
											history := make([]RequestResponse, len(b.history))
											copy(history, b.history)
											b.lock.Unlock()
											return history
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:122
	// _ = "end of CoverTab[104255]"
}

// Port returns the TCP port number the broker is listening for requests on.
func (b *MockBroker) Port() int32 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:126
	_go_fuzz_dep_.CoverTab[104256]++
											return b.port
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:127
	// _ = "end of CoverTab[104256]"
}

// Addr returns the broker connection string in the form "<address>:<port>".
func (b *MockBroker) Addr() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:131
	_go_fuzz_dep_.CoverTab[104257]++
											return b.listener.Addr().String()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:132
	// _ = "end of CoverTab[104257]"
}

// Close terminates the broker blocking until it stops internal goroutines and
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:135
// releases all resources.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:137
func (b *MockBroker) Close() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:137
	_go_fuzz_dep_.CoverTab[104258]++
											close(b.expectations)
											if len(b.expectations) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:139
		_go_fuzz_dep_.CoverTab[104260]++
												buf := bytes.NewBufferString(fmt.Sprintf("mockbroker/%d: not all expectations were satisfied! Still waiting on:\n", b.BrokerID()))
												for e := range b.expectations {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:141
			_go_fuzz_dep_.CoverTab[104262]++
													_, _ = buf.WriteString(spew.Sdump(e))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:142
			// _ = "end of CoverTab[104262]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:143
		// _ = "end of CoverTab[104260]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:143
		_go_fuzz_dep_.CoverTab[104261]++
												b.t.Error(buf.String())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:144
		// _ = "end of CoverTab[104261]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:145
		_go_fuzz_dep_.CoverTab[104263]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:145
		// _ = "end of CoverTab[104263]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:145
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:145
	// _ = "end of CoverTab[104258]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:145
	_go_fuzz_dep_.CoverTab[104259]++
											close(b.closing)
											<-b.stopper
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:147
	// _ = "end of CoverTab[104259]"
}

// setHandler sets the specified function as the request handler. Whenever
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:150
// a mock broker reads a request from the wire it passes the request to the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:150
// function and sends back whatever the handler function returns.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:153
func (b *MockBroker) setHandler(handler requestHandlerFunc) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:153
	_go_fuzz_dep_.CoverTab[104264]++
											b.lock.Lock()
											b.handler = handler
											b.lock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:156
	// _ = "end of CoverTab[104264]"
}

func (b *MockBroker) serverLoop() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:159
	_go_fuzz_dep_.CoverTab[104265]++
											defer close(b.stopper)
											var err error
											var conn net.Conn
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:162
	_curRoutineNum141_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:162
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum141_)

											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:164
		_go_fuzz_dep_.CoverTab[104268]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:164
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:164
			_go_fuzz_dep_.CoverTab[104269]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:164
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum141_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:164
			// _ = "end of CoverTab[104269]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:164
		}()
												<-b.closing
												err := b.listener.Close()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:167
			_go_fuzz_dep_.CoverTab[104270]++
													b.t.Error(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:168
			// _ = "end of CoverTab[104270]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:169
			_go_fuzz_dep_.CoverTab[104271]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:169
			// _ = "end of CoverTab[104271]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:169
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:169
		// _ = "end of CoverTab[104268]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:170
	// _ = "end of CoverTab[104265]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:170
	_go_fuzz_dep_.CoverTab[104266]++

											wg := &sync.WaitGroup{}
											i := 0
											for conn, err = b.listener.Accept(); err == nil; conn, err = b.listener.Accept() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:174
		_go_fuzz_dep_.CoverTab[104272]++
												wg.Add(1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:175
		_curRoutineNum142_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:175
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum142_)
												go b.handleRequests(conn, i, wg)
												i++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:177
		// _ = "end of CoverTab[104272]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:178
	// _ = "end of CoverTab[104266]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:178
	_go_fuzz_dep_.CoverTab[104267]++
											wg.Wait()
											Logger.Printf("*** mockbroker/%d: listener closed, err=%v", b.BrokerID(), err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:180
	// _ = "end of CoverTab[104267]"
}

func (b *MockBroker) SetGSSAPIHandler(handler GSSApiHandlerFunc) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:183
	_go_fuzz_dep_.CoverTab[104273]++
											b.gssApiHandler = handler
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:184
	// _ = "end of CoverTab[104273]"
}

func (b *MockBroker) readToBytes(r io.Reader) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:187
	_go_fuzz_dep_.CoverTab[104274]++
											var (
		bytesRead	int
		lengthBytes	= make([]byte, 4)
	)

	if _, err := io.ReadFull(r, lengthBytes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:193
		_go_fuzz_dep_.CoverTab[104278]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:194
		// _ = "end of CoverTab[104278]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:195
		_go_fuzz_dep_.CoverTab[104279]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:195
		// _ = "end of CoverTab[104279]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:195
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:195
	// _ = "end of CoverTab[104274]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:195
	_go_fuzz_dep_.CoverTab[104275]++

											bytesRead += len(lengthBytes)
											length := int32(binary.BigEndian.Uint32(lengthBytes))

											if length <= 4 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:200
		_go_fuzz_dep_.CoverTab[104280]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:200
		return length > MaxRequestSize
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:200
		// _ = "end of CoverTab[104280]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:200
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:200
		_go_fuzz_dep_.CoverTab[104281]++
												return nil, PacketDecodingError{fmt.Sprintf("message of length %d too large or too small", length)}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:201
		// _ = "end of CoverTab[104281]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:202
		_go_fuzz_dep_.CoverTab[104282]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:202
		// _ = "end of CoverTab[104282]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:202
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:202
	// _ = "end of CoverTab[104275]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:202
	_go_fuzz_dep_.CoverTab[104276]++

											encodedReq := make([]byte, length)
											if _, err := io.ReadFull(r, encodedReq); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:205
		_go_fuzz_dep_.CoverTab[104283]++
												return nil, err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:206
		// _ = "end of CoverTab[104283]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:207
		_go_fuzz_dep_.CoverTab[104284]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:207
		// _ = "end of CoverTab[104284]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:207
	// _ = "end of CoverTab[104276]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:207
	_go_fuzz_dep_.CoverTab[104277]++

											bytesRead += len(encodedReq)

											fullBytes := append(lengthBytes, encodedReq...)

											return fullBytes, nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:213
	// _ = "end of CoverTab[104277]"
}

func (b *MockBroker) isGSSAPI(buffer []byte) bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:216
	_go_fuzz_dep_.CoverTab[104285]++
											return buffer[4] == 0x60 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:217
		_go_fuzz_dep_.CoverTab[104286]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:217
		return bytes.Equal(buffer[4:6], []byte{0x05, 0x04})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:217
		// _ = "end of CoverTab[104286]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:217
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:217
	// _ = "end of CoverTab[104285]"
}

func (b *MockBroker) handleRequests(conn io.ReadWriteCloser, idx int, wg *sync.WaitGroup) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:220
	_go_fuzz_dep_.CoverTab[104287]++
											defer wg.Done()
											defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:222
		_go_fuzz_dep_.CoverTab[104291]++
												_ = conn.Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:223
		// _ = "end of CoverTab[104291]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:224
	// _ = "end of CoverTab[104287]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:224
	_go_fuzz_dep_.CoverTab[104288]++
											s := spew.NewDefaultConfig()
											s.MaxDepth = 1
											Logger.Printf("*** mockbroker/%d/%d: connection opened", b.BrokerID(), idx)
											var err error

											abort := make(chan none)
											defer close(abort)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:231
	_curRoutineNum143_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:231
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum143_)
											go func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:232
		_go_fuzz_dep_.CoverTab[104292]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:232
		defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:232
			_go_fuzz_dep_.CoverTab[104293]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:232
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum143_)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:232
			// _ = "end of CoverTab[104293]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:232
		}()
												select {
		case <-b.closing:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:234
			_go_fuzz_dep_.CoverTab[104294]++
													_ = conn.Close()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:235
			// _ = "end of CoverTab[104294]"
		case <-abort:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:236
			_go_fuzz_dep_.CoverTab[104295]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:236
			// _ = "end of CoverTab[104295]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:237
		// _ = "end of CoverTab[104292]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:238
	// _ = "end of CoverTab[104288]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:238
	_go_fuzz_dep_.CoverTab[104289]++

											var bytesWritten int
											var bytesRead int
											for {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:242
		_go_fuzz_dep_.CoverTab[104296]++
												buffer, err := b.readToBytes(conn)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:244
			_go_fuzz_dep_.CoverTab[104300]++
													Logger.Printf("*** mockbroker/%d/%d: invalid request: err=%+v, %+v", b.brokerID, idx, err, spew.Sdump(buffer))
													b.serverError(err)
													break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:247
			// _ = "end of CoverTab[104300]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:248
			_go_fuzz_dep_.CoverTab[104301]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:248
			// _ = "end of CoverTab[104301]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:248
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:248
		// _ = "end of CoverTab[104296]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:248
		_go_fuzz_dep_.CoverTab[104297]++

												bytesWritten = 0
												if !b.isGSSAPI(buffer) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:251
			_go_fuzz_dep_.CoverTab[104302]++
													req, br, err := decodeRequest(bytes.NewReader(buffer))
													bytesRead = br
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:254
				_go_fuzz_dep_.CoverTab[104310]++
														Logger.Printf("*** mockbroker/%d/%d: invalid request: err=%+v, %+v", b.brokerID, idx, err, spew.Sdump(req))
														b.serverError(err)
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:257
				// _ = "end of CoverTab[104310]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:258
				_go_fuzz_dep_.CoverTab[104311]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:258
				// _ = "end of CoverTab[104311]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:258
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:258
			// _ = "end of CoverTab[104302]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:258
			_go_fuzz_dep_.CoverTab[104303]++

													if b.latency > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:260
				_go_fuzz_dep_.CoverTab[104312]++
														time.Sleep(b.latency)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:261
				// _ = "end of CoverTab[104312]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:262
				_go_fuzz_dep_.CoverTab[104313]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:262
				// _ = "end of CoverTab[104313]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:262
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:262
			// _ = "end of CoverTab[104303]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:262
			_go_fuzz_dep_.CoverTab[104304]++

													b.lock.Lock()
													res := b.handler(req)
													b.history = append(b.history, RequestResponse{req.body, res})
													b.lock.Unlock()

													if res == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:269
				_go_fuzz_dep_.CoverTab[104314]++
														Logger.Printf("*** mockbroker/%d/%d: ignored %v", b.brokerID, idx, spew.Sdump(req))
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:271
				// _ = "end of CoverTab[104314]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:272
				_go_fuzz_dep_.CoverTab[104315]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:272
				// _ = "end of CoverTab[104315]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:272
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:272
			// _ = "end of CoverTab[104304]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:272
			_go_fuzz_dep_.CoverTab[104305]++
													Logger.Printf(
				"*** mockbroker/%d/%d: replied to %T with %T\n-> %s\n-> %s",
				b.brokerID, idx, req.body, res,
				s.Sprintf("%#v", req.body),
				s.Sprintf("%#v", res),
			)

			encodedRes, err := encode(res, nil)
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:281
				_go_fuzz_dep_.CoverTab[104316]++
														b.serverError(err)
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:283
				// _ = "end of CoverTab[104316]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:284
				_go_fuzz_dep_.CoverTab[104317]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:284
				// _ = "end of CoverTab[104317]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:284
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:284
			// _ = "end of CoverTab[104305]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:284
			_go_fuzz_dep_.CoverTab[104306]++
													if len(encodedRes) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:285
				_go_fuzz_dep_.CoverTab[104318]++
														b.lock.Lock()
														if b.notifier != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:287
					_go_fuzz_dep_.CoverTab[104320]++
															b.notifier(bytesRead, 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:288
					// _ = "end of CoverTab[104320]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:289
					_go_fuzz_dep_.CoverTab[104321]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:289
					// _ = "end of CoverTab[104321]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:289
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:289
				// _ = "end of CoverTab[104318]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:289
				_go_fuzz_dep_.CoverTab[104319]++
														b.lock.Unlock()
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:291
				// _ = "end of CoverTab[104319]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:292
				_go_fuzz_dep_.CoverTab[104322]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:292
				// _ = "end of CoverTab[104322]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:292
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:292
			// _ = "end of CoverTab[104306]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:292
			_go_fuzz_dep_.CoverTab[104307]++

													resHeader := b.encodeHeader(res.headerVersion(), req.correlationID, uint32(len(encodedRes)))
													if _, err = conn.Write(resHeader); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:295
				_go_fuzz_dep_.CoverTab[104323]++
														b.serverError(err)
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:297
				// _ = "end of CoverTab[104323]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:298
				_go_fuzz_dep_.CoverTab[104324]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:298
				// _ = "end of CoverTab[104324]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:298
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:298
			// _ = "end of CoverTab[104307]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:298
			_go_fuzz_dep_.CoverTab[104308]++
													if _, err = conn.Write(encodedRes); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:299
				_go_fuzz_dep_.CoverTab[104325]++
														b.serverError(err)
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:301
				// _ = "end of CoverTab[104325]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:302
				_go_fuzz_dep_.CoverTab[104326]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:302
				// _ = "end of CoverTab[104326]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:302
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:302
			// _ = "end of CoverTab[104308]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:302
			_go_fuzz_dep_.CoverTab[104309]++
													bytesWritten = len(resHeader) + len(encodedRes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:303
			// _ = "end of CoverTab[104309]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:304
			_go_fuzz_dep_.CoverTab[104327]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:307
			b.lock.Lock()
			res := b.gssApiHandler(buffer)
			b.lock.Unlock()
			if res == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:310
				_go_fuzz_dep_.CoverTab[104330]++
														Logger.Printf("*** mockbroker/%d/%d: ignored %v", b.brokerID, idx, spew.Sdump(buffer))
														continue
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:312
				// _ = "end of CoverTab[104330]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:313
				_go_fuzz_dep_.CoverTab[104331]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:313
				// _ = "end of CoverTab[104331]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:313
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:313
			// _ = "end of CoverTab[104327]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:313
			_go_fuzz_dep_.CoverTab[104328]++
													if _, err = conn.Write(res); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:314
				_go_fuzz_dep_.CoverTab[104332]++
														b.serverError(err)
														break
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:316
				// _ = "end of CoverTab[104332]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:317
				_go_fuzz_dep_.CoverTab[104333]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:317
				// _ = "end of CoverTab[104333]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:317
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:317
			// _ = "end of CoverTab[104328]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:317
			_go_fuzz_dep_.CoverTab[104329]++
													bytesWritten = len(res)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:318
			// _ = "end of CoverTab[104329]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:319
		// _ = "end of CoverTab[104297]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:319
		_go_fuzz_dep_.CoverTab[104298]++

												b.lock.Lock()
												if b.notifier != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:322
			_go_fuzz_dep_.CoverTab[104334]++
													b.notifier(bytesRead, bytesWritten)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:323
			// _ = "end of CoverTab[104334]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:324
			_go_fuzz_dep_.CoverTab[104335]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:324
			// _ = "end of CoverTab[104335]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:324
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:324
		// _ = "end of CoverTab[104298]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:324
		_go_fuzz_dep_.CoverTab[104299]++
												b.lock.Unlock()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:325
		// _ = "end of CoverTab[104299]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:326
	// _ = "end of CoverTab[104289]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:326
	_go_fuzz_dep_.CoverTab[104290]++
											Logger.Printf("*** mockbroker/%d/%d: connection closed, err=%v", b.BrokerID(), idx, err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:327
	// _ = "end of CoverTab[104290]"
}

func (b *MockBroker) encodeHeader(headerVersion int16, correlationId int32, payloadLength uint32) []byte {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:330
	_go_fuzz_dep_.CoverTab[104336]++
											headerLength := uint32(8)

											if headerVersion >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:333
		_go_fuzz_dep_.CoverTab[104339]++
												headerLength = 9
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:334
		// _ = "end of CoverTab[104339]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:335
		_go_fuzz_dep_.CoverTab[104340]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:335
		// _ = "end of CoverTab[104340]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:335
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:335
	// _ = "end of CoverTab[104336]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:335
	_go_fuzz_dep_.CoverTab[104337]++

											resHeader := make([]byte, headerLength)
											binary.BigEndian.PutUint32(resHeader, payloadLength+headerLength-4)
											binary.BigEndian.PutUint32(resHeader[4:], uint32(correlationId))

											if headerVersion >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:341
		_go_fuzz_dep_.CoverTab[104341]++
												binary.PutUvarint(resHeader[8:], 0)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:342
		// _ = "end of CoverTab[104341]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:343
		_go_fuzz_dep_.CoverTab[104342]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:343
		// _ = "end of CoverTab[104342]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:343
	// _ = "end of CoverTab[104337]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:343
	_go_fuzz_dep_.CoverTab[104338]++

											return resHeader
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:345
	// _ = "end of CoverTab[104338]"
}

func (b *MockBroker) defaultRequestHandler(req *request) (res encoderWithHeader) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:348
	_go_fuzz_dep_.CoverTab[104343]++
											select {
	case res, ok := <-b.expectations:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:350
		_go_fuzz_dep_.CoverTab[104344]++
												if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:351
			_go_fuzz_dep_.CoverTab[104347]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:352
			// _ = "end of CoverTab[104347]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:353
			_go_fuzz_dep_.CoverTab[104348]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:353
			// _ = "end of CoverTab[104348]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:353
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:353
		// _ = "end of CoverTab[104344]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:353
		_go_fuzz_dep_.CoverTab[104345]++
												return res
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:354
		// _ = "end of CoverTab[104345]"
	case <-time.After(expectationTimeout):
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:355
		_go_fuzz_dep_.CoverTab[104346]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:356
		// _ = "end of CoverTab[104346]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:357
	// _ = "end of CoverTab[104343]"
}

func (b *MockBroker) serverError(err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:360
	_go_fuzz_dep_.CoverTab[104349]++
											isConnectionClosedError := false
											if _, ok := err.(*net.OpError); ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:362
		_go_fuzz_dep_.CoverTab[104352]++
												isConnectionClosedError = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:363
		// _ = "end of CoverTab[104352]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:364
		_go_fuzz_dep_.CoverTab[104353]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:364
		if err == io.EOF {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:364
			_go_fuzz_dep_.CoverTab[104354]++
													isConnectionClosedError = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:365
			// _ = "end of CoverTab[104354]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:366
			_go_fuzz_dep_.CoverTab[104355]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:366
			if err.Error() == "use of closed network connection" {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:366
				_go_fuzz_dep_.CoverTab[104356]++
														isConnectionClosedError = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:367
				// _ = "end of CoverTab[104356]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
				_go_fuzz_dep_.CoverTab[104357]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
				// _ = "end of CoverTab[104357]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
			// _ = "end of CoverTab[104355]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
		// _ = "end of CoverTab[104353]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
	// _ = "end of CoverTab[104349]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:368
	_go_fuzz_dep_.CoverTab[104350]++

											if isConnectionClosedError {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:370
		_go_fuzz_dep_.CoverTab[104358]++
												return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:371
		// _ = "end of CoverTab[104358]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:372
		_go_fuzz_dep_.CoverTab[104359]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:372
		// _ = "end of CoverTab[104359]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:372
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:372
	// _ = "end of CoverTab[104350]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:372
	_go_fuzz_dep_.CoverTab[104351]++

											b.t.Errorf(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:374
	// _ = "end of CoverTab[104351]"
}

// NewMockBroker launches a fake Kafka broker. It takes a TestReporter as provided by the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:377
// test framework and a channel of responses to use.  If an error occurs it is
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:377
// simply logged to the TestReporter and the broker exits.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:380
func NewMockBroker(t TestReporter, brokerID int32) *MockBroker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:380
	_go_fuzz_dep_.CoverTab[104360]++
											return NewMockBrokerAddr(t, brokerID, "localhost:0")
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:381
	// _ = "end of CoverTab[104360]"
}

// NewMockBrokerAddr behaves like newMockBroker but listens on the address you give
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:384
// it rather than just some ephemeral port.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:386
func NewMockBrokerAddr(t TestReporter, brokerID int32, addr string) *MockBroker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:386
	_go_fuzz_dep_.CoverTab[104361]++
											listener, err := net.Listen("tcp", addr)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:388
		_go_fuzz_dep_.CoverTab[104363]++
												t.Fatal(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:389
		// _ = "end of CoverTab[104363]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:390
		_go_fuzz_dep_.CoverTab[104364]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:390
		// _ = "end of CoverTab[104364]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:390
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:390
	// _ = "end of CoverTab[104361]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:390
	_go_fuzz_dep_.CoverTab[104362]++
											return NewMockBrokerListener(t, brokerID, listener)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:391
	// _ = "end of CoverTab[104362]"
}

// NewMockBrokerListener behaves like newMockBrokerAddr but accepts connections on the listener specified.
func NewMockBrokerListener(t TestReporter, brokerID int32, listener net.Listener) *MockBroker {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:395
	_go_fuzz_dep_.CoverTab[104365]++
											var err error

											broker := &MockBroker{
		closing:	make(chan none),
		stopper:	make(chan none),
		t:		t,
		brokerID:	brokerID,
		expectations:	make(chan encoderWithHeader, 512),
		listener:	listener,
	}
	broker.handler = broker.defaultRequestHandler

	Logger.Printf("*** mockbroker/%d listening on %s\n", brokerID, broker.listener.Addr().String())
	_, portStr, err := net.SplitHostPort(broker.listener.Addr().String())
	if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:410
		_go_fuzz_dep_.CoverTab[104368]++
												t.Fatal(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:411
		// _ = "end of CoverTab[104368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:412
		_go_fuzz_dep_.CoverTab[104369]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:412
		// _ = "end of CoverTab[104369]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:412
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:412
	// _ = "end of CoverTab[104365]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:412
	_go_fuzz_dep_.CoverTab[104366]++
											tmp, err := strconv.ParseInt(portStr, 10, 32)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:414
		_go_fuzz_dep_.CoverTab[104370]++
												t.Fatal(err)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:415
		// _ = "end of CoverTab[104370]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:416
		_go_fuzz_dep_.CoverTab[104371]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:416
		// _ = "end of CoverTab[104371]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:416
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:416
	// _ = "end of CoverTab[104366]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:416
	_go_fuzz_dep_.CoverTab[104367]++
											broker.port = int32(tmp)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:417
	_curRoutineNum144_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:417
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum144_)

											go broker.serverLoop()

											return broker
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:421
	// _ = "end of CoverTab[104367]"
}

func (b *MockBroker) Returns(e encoderWithHeader) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:424
	_go_fuzz_dep_.CoverTab[104372]++
											b.expectations <- e
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:425
	// _ = "end of CoverTab[104372]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:426
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/mockbroker.go:426
var _ = _go_fuzz_dep_.CoverTab
