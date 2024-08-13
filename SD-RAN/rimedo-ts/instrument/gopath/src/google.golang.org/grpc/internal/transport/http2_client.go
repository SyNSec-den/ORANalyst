//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:19
)

import (
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	icredentials "google.golang.org/grpc/internal/credentials"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/grpcutil"
	imetadata "google.golang.org/grpc/internal/metadata"
	istatus "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/internal/syscall"
	"google.golang.org/grpc/internal/transport/networktype"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

// clientConnectionCounter counts the number of connections a client has
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:55
// initiated (equal to the number of http2Clients created). Must be accessed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:55
// atomically.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:58
var clientConnectionCounter uint64

// http2Client implements the ClientTransport interface with HTTP2.
type http2Client struct {
	lastRead	int64	// Keep this field 64-bit aligned. Accessed atomically.
	ctx		context.Context
	cancel		context.CancelFunc
	ctxDone		<-chan struct{}	// Cache the ctx.Done() chan.
	userAgent	string
	// address contains the resolver returned address for this transport.
	// If the `ServerName` field is set, it takes precedence over `CallHdr.Host`
	// passed to `NewStream`, when determining the :authority header.
	address		resolver.Address
	md		metadata.MD
	conn		net.Conn	// underlying communication channel
	loopy		*loopyWriter
	remoteAddr	net.Addr
	localAddr	net.Addr
	authInfo	credentials.AuthInfo	// auth info about the connection

	readerDone	chan struct{}	// sync point to enable testing.
	writerDone	chan struct{}	// sync point to enable testing.
	// goAway is closed to notify the upper layer (i.e., addrConn.transportMonitor)
	// that the server sent GoAway on this transport.
	goAway	chan struct{}

	framer	*framer
	// controlBuf delivers all the control related tasks (e.g., window
	// updates, reset streams, and various settings) to the controller.
	// Do not access controlBuf with mu held.
	controlBuf	*controlBuffer
	fc		*trInFlow
	// The scheme used: https if TLS is on, http otherwise.
	scheme	string

	isSecure	bool

	perRPCCreds	[]credentials.PerRPCCredentials

	kp			keepalive.ClientParameters
	keepaliveEnabled	bool

	statsHandlers	[]stats.Handler

	initialWindowSize	int32

	// configured by peer through SETTINGS_MAX_HEADER_LIST_SIZE
	maxSendHeaderListSize	*uint32

	bdpEst	*bdpEstimator

	maxConcurrentStreams	uint32
	streamQuota		int64
	streamsQuotaAvailable	chan struct{}
	waitingStreams		uint32
	nextID			uint32
	registeredCompressors	string

	// Do not access controlBuf with mu held.
	mu		sync.Mutex	// guard the following variables
	state		transportState
	activeStreams	map[uint32]*Stream
	// prevGoAway ID records the Last-Stream-ID in the previous GOAway frame.
	prevGoAwayID	uint32
	// goAwayReason records the http2.ErrCode and debug data received with the
	// GoAway frame.
	goAwayReason	GoAwayReason
	// goAwayDebugMessage contains a detailed human readable string about a
	// GoAway frame, useful for error messages.
	goAwayDebugMessage	string
	// A condition variable used to signal when the keepalive goroutine should
	// go dormant. The condition for dormancy is based on the number of active
	// streams and the `PermitWithoutStream` keepalive client parameter. And
	// since the number of active streams is guarded by the above mutex, we use
	// the same for this condition variable as well.
	kpDormancyCond	*sync.Cond
	// A boolean to track whether the keepalive goroutine is dormant or not.
	// This is checked before attempting to signal the above condition
	// variable.
	kpDormant	bool

	// Fields below are for channelz metric collection.
	channelzID	*channelz.Identifier
	czData		*channelzData

	onClose	func(GoAwayReason)

	bufferPool	*bufferPool

	connectionID	uint64
}

func dial(ctx context.Context, fn func(context.Context, string) (net.Conn, error), addr resolver.Address, useProxy bool, grpcUA string) (net.Conn, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:150
	_go_fuzz_dep_.CoverTab[77152]++
														address := addr.Addr
														networkType, ok := networktype.Get(addr)
														if fn != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:153
		_go_fuzz_dep_.CoverTab[77156]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:161
		if networkType == "unix" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:161
			_go_fuzz_dep_.CoverTab[77158]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:161
			return !strings.HasPrefix(address, "\x00")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:161
			// _ = "end of CoverTab[77158]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:161
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:161
			_go_fuzz_dep_.CoverTab[77159]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:164
			if filepath.IsAbs(address) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:164
				_go_fuzz_dep_.CoverTab[77161]++
																	return fn(ctx, "unix://"+address)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:165
				// _ = "end of CoverTab[77161]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:166
				_go_fuzz_dep_.CoverTab[77162]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:166
				// _ = "end of CoverTab[77162]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:166
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:166
			// _ = "end of CoverTab[77159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:166
			_go_fuzz_dep_.CoverTab[77160]++
																return fn(ctx, "unix:"+address)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:167
			// _ = "end of CoverTab[77160]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:168
			_go_fuzz_dep_.CoverTab[77163]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:168
			// _ = "end of CoverTab[77163]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:168
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:168
		// _ = "end of CoverTab[77156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:168
		_go_fuzz_dep_.CoverTab[77157]++
															return fn(ctx, address)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:169
		// _ = "end of CoverTab[77157]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:170
		_go_fuzz_dep_.CoverTab[77164]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:170
		// _ = "end of CoverTab[77164]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:170
	// _ = "end of CoverTab[77152]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:170
	_go_fuzz_dep_.CoverTab[77153]++
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:171
		_go_fuzz_dep_.CoverTab[77165]++
															networkType, address = parseDialTarget(address)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:172
		// _ = "end of CoverTab[77165]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:173
		_go_fuzz_dep_.CoverTab[77166]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:173
		// _ = "end of CoverTab[77166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:173
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:173
	// _ = "end of CoverTab[77153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:173
	_go_fuzz_dep_.CoverTab[77154]++
														if networkType == "tcp" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:174
		_go_fuzz_dep_.CoverTab[77167]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:174
		return useProxy
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:174
		// _ = "end of CoverTab[77167]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:174
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:174
		_go_fuzz_dep_.CoverTab[77168]++
															return proxyDial(ctx, address, grpcUA)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:175
		// _ = "end of CoverTab[77168]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:176
		_go_fuzz_dep_.CoverTab[77169]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:176
		// _ = "end of CoverTab[77169]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:176
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:176
	// _ = "end of CoverTab[77154]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:176
	_go_fuzz_dep_.CoverTab[77155]++
														return (&net.Dialer{}).DialContext(ctx, networkType, address)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:177
	// _ = "end of CoverTab[77155]"
}

func isTemporary(err error) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:180
	_go_fuzz_dep_.CoverTab[77170]++
														switch err := err.(type) {
	case interface {
		Temporary() bool
	}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:184
		_go_fuzz_dep_.CoverTab[77172]++
															return err.Temporary()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:185
		// _ = "end of CoverTab[77172]"
	case interface {
		Timeout() bool
	}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:188
		_go_fuzz_dep_.CoverTab[77173]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:191
		return err.Timeout()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:191
		// _ = "end of CoverTab[77173]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:192
	// _ = "end of CoverTab[77170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:192
	_go_fuzz_dep_.CoverTab[77171]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:193
	// _ = "end of CoverTab[77171]"
}

// newHTTP2Client constructs a connected ClientTransport to addr based on HTTP2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:196
// and starts to receive messages on it. Non-nil error returns if construction
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:196
// fails.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:199
func newHTTP2Client(connectCtx, ctx context.Context, addr resolver.Address, opts ConnectOptions, onClose func(GoAwayReason)) (_ *http2Client, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:199
	_go_fuzz_dep_.CoverTab[77174]++
														scheme := "http"
														ctx, cancel := context.WithCancel(ctx)
														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:202
		_go_fuzz_dep_.CoverTab[77202]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:203
			_go_fuzz_dep_.CoverTab[77203]++
																cancel()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:204
			// _ = "end of CoverTab[77203]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:205
			_go_fuzz_dep_.CoverTab[77204]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:205
			// _ = "end of CoverTab[77204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:205
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:205
		// _ = "end of CoverTab[77202]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:206
	// _ = "end of CoverTab[77174]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:206
	_go_fuzz_dep_.CoverTab[77175]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:212
	connectCtx = icredentials.NewClientHandshakeInfoContext(connectCtx, credentials.ClientHandshakeInfo{Attributes: addr.Attributes})

	conn, err := dial(connectCtx, opts.Dialer, addr, opts.UseProxy, opts.UserAgent)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:215
		_go_fuzz_dep_.CoverTab[77205]++
															if opts.FailOnNonTempDialError {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:216
			_go_fuzz_dep_.CoverTab[77207]++
																return nil, connectionErrorf(isTemporary(err), err, "transport: error while dialing: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:217
			// _ = "end of CoverTab[77207]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:218
			_go_fuzz_dep_.CoverTab[77208]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:218
			// _ = "end of CoverTab[77208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:218
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:218
		// _ = "end of CoverTab[77205]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:218
		_go_fuzz_dep_.CoverTab[77206]++
															return nil, connectionErrorf(true, err, "transport: Error while dialing: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:219
		// _ = "end of CoverTab[77206]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:220
		_go_fuzz_dep_.CoverTab[77209]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:220
		// _ = "end of CoverTab[77209]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:220
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:220
	// _ = "end of CoverTab[77175]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:220
	_go_fuzz_dep_.CoverTab[77176]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:223
	defer func(conn net.Conn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:223
		_go_fuzz_dep_.CoverTab[77210]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:224
			_go_fuzz_dep_.CoverTab[77211]++
																conn.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:225
			// _ = "end of CoverTab[77211]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:226
			_go_fuzz_dep_.CoverTab[77212]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:226
			// _ = "end of CoverTab[77212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:226
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:226
		// _ = "end of CoverTab[77210]"
	}(conn)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:227
	// _ = "end of CoverTab[77176]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:227
	_go_fuzz_dep_.CoverTab[77177]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:235
	ctxMonitorDone := grpcsync.NewEvent()
	newClientCtx, newClientDone := context.WithCancel(connectCtx)
	defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:237
		_go_fuzz_dep_.CoverTab[77213]++
															newClientDone()
															<-ctxMonitorDone.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:239
		// _ = "end of CoverTab[77213]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:240
	// _ = "end of CoverTab[77177]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:240
	_go_fuzz_dep_.CoverTab[77178]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:240
	_curRoutineNum78_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:240
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum78_)
														go func(conn net.Conn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:241
		_go_fuzz_dep_.CoverTab[77214]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:241
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:241
			_go_fuzz_dep_.CoverTab[77215]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:241
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum78_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:241
			// _ = "end of CoverTab[77215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:241
		}()
															defer ctxMonitorDone.Fire()
															<-newClientCtx.Done()
															if err := connectCtx.Err(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:244
			_go_fuzz_dep_.CoverTab[77216]++

																if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:246
				_go_fuzz_dep_.CoverTab[77218]++
																	logger.Infof("newClientTransport: aborting due to connectCtx: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:247
				// _ = "end of CoverTab[77218]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:248
				_go_fuzz_dep_.CoverTab[77219]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:248
				// _ = "end of CoverTab[77219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:248
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:248
			// _ = "end of CoverTab[77216]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:248
			_go_fuzz_dep_.CoverTab[77217]++
																conn.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:249
			// _ = "end of CoverTab[77217]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:250
			_go_fuzz_dep_.CoverTab[77220]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:250
			// _ = "end of CoverTab[77220]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:250
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:250
		// _ = "end of CoverTab[77214]"
	}(conn)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:251
	// _ = "end of CoverTab[77178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:251
	_go_fuzz_dep_.CoverTab[77179]++

														kp := opts.KeepaliveParams

														if kp.Time == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:255
		_go_fuzz_dep_.CoverTab[77221]++
															kp.Time = defaultClientKeepaliveTime
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:256
		// _ = "end of CoverTab[77221]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:257
		_go_fuzz_dep_.CoverTab[77222]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:257
		// _ = "end of CoverTab[77222]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:257
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:257
	// _ = "end of CoverTab[77179]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:257
	_go_fuzz_dep_.CoverTab[77180]++
														if kp.Timeout == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:258
		_go_fuzz_dep_.CoverTab[77223]++
															kp.Timeout = defaultClientKeepaliveTimeout
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:259
		// _ = "end of CoverTab[77223]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:260
		_go_fuzz_dep_.CoverTab[77224]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:260
		// _ = "end of CoverTab[77224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:260
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:260
	// _ = "end of CoverTab[77180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:260
	_go_fuzz_dep_.CoverTab[77181]++
														keepaliveEnabled := false
														if kp.Time != infinity {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:262
		_go_fuzz_dep_.CoverTab[77225]++
															if err = syscall.SetTCPUserTimeout(conn, kp.Timeout); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:263
			_go_fuzz_dep_.CoverTab[77227]++
																return nil, connectionErrorf(false, err, "transport: failed to set TCP_USER_TIMEOUT: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:264
			// _ = "end of CoverTab[77227]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:265
			_go_fuzz_dep_.CoverTab[77228]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:265
			// _ = "end of CoverTab[77228]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:265
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:265
		// _ = "end of CoverTab[77225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:265
		_go_fuzz_dep_.CoverTab[77226]++
															keepaliveEnabled = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:266
		// _ = "end of CoverTab[77226]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:267
		_go_fuzz_dep_.CoverTab[77229]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:267
		// _ = "end of CoverTab[77229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:267
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:267
	// _ = "end of CoverTab[77181]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:267
	_go_fuzz_dep_.CoverTab[77182]++
														var (
		isSecure	bool
		authInfo	credentials.AuthInfo
	)
	transportCreds := opts.TransportCredentials
	perRPCCreds := opts.PerRPCCredentials

	if b := opts.CredsBundle; b != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:275
		_go_fuzz_dep_.CoverTab[77230]++
															if t := b.TransportCredentials(); t != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:276
			_go_fuzz_dep_.CoverTab[77232]++
																transportCreds = t
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:277
			// _ = "end of CoverTab[77232]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:278
			_go_fuzz_dep_.CoverTab[77233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:278
			// _ = "end of CoverTab[77233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:278
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:278
		// _ = "end of CoverTab[77230]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:278
		_go_fuzz_dep_.CoverTab[77231]++
															if t := b.PerRPCCredentials(); t != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:279
			_go_fuzz_dep_.CoverTab[77234]++
																perRPCCreds = append(perRPCCreds, t)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:280
			// _ = "end of CoverTab[77234]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:281
			_go_fuzz_dep_.CoverTab[77235]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:281
			// _ = "end of CoverTab[77235]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:281
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:281
		// _ = "end of CoverTab[77231]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:282
		_go_fuzz_dep_.CoverTab[77236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:282
		// _ = "end of CoverTab[77236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:282
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:282
	// _ = "end of CoverTab[77182]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:282
	_go_fuzz_dep_.CoverTab[77183]++
														if transportCreds != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:283
		_go_fuzz_dep_.CoverTab[77237]++
															conn, authInfo, err = transportCreds.ClientHandshake(connectCtx, addr.ServerName, conn)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:285
			_go_fuzz_dep_.CoverTab[77240]++
																return nil, connectionErrorf(isTemporary(err), err, "transport: authentication handshake failed: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:286
			// _ = "end of CoverTab[77240]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:287
			_go_fuzz_dep_.CoverTab[77241]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:287
			// _ = "end of CoverTab[77241]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:287
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:287
		// _ = "end of CoverTab[77237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:287
		_go_fuzz_dep_.CoverTab[77238]++
															for _, cd := range perRPCCreds {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:288
			_go_fuzz_dep_.CoverTab[77242]++
																if cd.RequireTransportSecurity() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:289
				_go_fuzz_dep_.CoverTab[77243]++
																	if ci, ok := authInfo.(interface {
					GetCommonAuthInfo() credentials.CommonAuthInfo
				}); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:292
					_go_fuzz_dep_.CoverTab[77244]++
																		secLevel := ci.GetCommonAuthInfo().SecurityLevel
																		if secLevel != credentials.InvalidSecurityLevel && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:294
						_go_fuzz_dep_.CoverTab[77245]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:294
						return secLevel < credentials.PrivacyAndIntegrity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:294
						// _ = "end of CoverTab[77245]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:294
					}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:294
						_go_fuzz_dep_.CoverTab[77246]++
																			return nil, connectionErrorf(true, nil, "transport: cannot send secure credentials on an insecure connection")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:295
						// _ = "end of CoverTab[77246]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:296
						_go_fuzz_dep_.CoverTab[77247]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:296
						// _ = "end of CoverTab[77247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:296
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:296
					// _ = "end of CoverTab[77244]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:297
					_go_fuzz_dep_.CoverTab[77248]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:297
					// _ = "end of CoverTab[77248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:297
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:297
				// _ = "end of CoverTab[77243]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:298
				_go_fuzz_dep_.CoverTab[77249]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:298
				// _ = "end of CoverTab[77249]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:298
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:298
			// _ = "end of CoverTab[77242]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:299
		// _ = "end of CoverTab[77238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:299
		_go_fuzz_dep_.CoverTab[77239]++
															isSecure = true
															if transportCreds.Info().SecurityProtocol == "tls" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:301
			_go_fuzz_dep_.CoverTab[77250]++
																scheme = "https"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:302
			// _ = "end of CoverTab[77250]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:303
			_go_fuzz_dep_.CoverTab[77251]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:303
			// _ = "end of CoverTab[77251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:303
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:303
		// _ = "end of CoverTab[77239]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:304
		_go_fuzz_dep_.CoverTab[77252]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:304
		// _ = "end of CoverTab[77252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:304
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:304
	// _ = "end of CoverTab[77183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:304
	_go_fuzz_dep_.CoverTab[77184]++
														dynamicWindow := true
														icwz := int32(initialWindowSize)
														if opts.InitialConnWindowSize >= defaultWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:307
		_go_fuzz_dep_.CoverTab[77253]++
															icwz = opts.InitialConnWindowSize
															dynamicWindow = false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:309
		// _ = "end of CoverTab[77253]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:310
		_go_fuzz_dep_.CoverTab[77254]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:310
		// _ = "end of CoverTab[77254]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:310
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:310
	// _ = "end of CoverTab[77184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:310
	_go_fuzz_dep_.CoverTab[77185]++
														writeBufSize := opts.WriteBufferSize
														readBufSize := opts.ReadBufferSize
														maxHeaderListSize := defaultClientMaxHeaderListSize
														if opts.MaxHeaderListSize != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:314
		_go_fuzz_dep_.CoverTab[77255]++
															maxHeaderListSize = *opts.MaxHeaderListSize
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:315
		// _ = "end of CoverTab[77255]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:316
		_go_fuzz_dep_.CoverTab[77256]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:316
		// _ = "end of CoverTab[77256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:316
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:316
	// _ = "end of CoverTab[77185]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:316
	_go_fuzz_dep_.CoverTab[77186]++
														t := &http2Client{
		ctx:			ctx,
		ctxDone:		ctx.Done(),
		cancel:			cancel,
		userAgent:		opts.UserAgent,
		registeredCompressors:	grpcutil.RegisteredCompressors(),
		address:		addr,
		conn:			conn,
		remoteAddr:		conn.RemoteAddr(),
		localAddr:		conn.LocalAddr(),
		authInfo:		authInfo,
		readerDone:		make(chan struct{}),
		writerDone:		make(chan struct{}),
		goAway:			make(chan struct{}),
		framer:			newFramer(conn, writeBufSize, readBufSize, maxHeaderListSize),
		fc:			&trInFlow{limit: uint32(icwz)},
		scheme:			scheme,
		activeStreams:		make(map[uint32]*Stream),
		isSecure:		isSecure,
		perRPCCreds:		perRPCCreds,
		kp:			kp,
		statsHandlers:		opts.StatsHandlers,
		initialWindowSize:	initialWindowSize,
		nextID:			1,
		maxConcurrentStreams:	defaultMaxStreamsClient,
		streamQuota:		defaultMaxStreamsClient,
		streamsQuotaAvailable:	make(chan struct{}, 1),
		czData:			new(channelzData),
		keepaliveEnabled:	keepaliveEnabled,
		bufferPool:		newBufferPool(),
		onClose:		onClose,
	}

	t.ctx = peer.NewContext(t.ctx, t.getPeer())

	if md, ok := addr.Metadata.(*metadata.MD); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:352
		_go_fuzz_dep_.CoverTab[77257]++
															t.md = *md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:353
		// _ = "end of CoverTab[77257]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:354
		_go_fuzz_dep_.CoverTab[77258]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:354
		if md := imetadata.Get(addr); md != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:354
			_go_fuzz_dep_.CoverTab[77259]++
																t.md = md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:355
			// _ = "end of CoverTab[77259]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
			_go_fuzz_dep_.CoverTab[77260]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
			// _ = "end of CoverTab[77260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
		// _ = "end of CoverTab[77258]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
	// _ = "end of CoverTab[77186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:356
	_go_fuzz_dep_.CoverTab[77187]++
														t.controlBuf = newControlBuffer(t.ctxDone)
														if opts.InitialWindowSize >= defaultWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:358
		_go_fuzz_dep_.CoverTab[77261]++
															t.initialWindowSize = opts.InitialWindowSize
															dynamicWindow = false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:360
		// _ = "end of CoverTab[77261]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:361
		_go_fuzz_dep_.CoverTab[77262]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:361
		// _ = "end of CoverTab[77262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:361
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:361
	// _ = "end of CoverTab[77187]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:361
	_go_fuzz_dep_.CoverTab[77188]++
														if dynamicWindow {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:362
		_go_fuzz_dep_.CoverTab[77263]++
															t.bdpEst = &bdpEstimator{
			bdp:			initialWindowSize,
			updateFlowControl:	t.updateFlowControl,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:366
		// _ = "end of CoverTab[77263]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:367
		_go_fuzz_dep_.CoverTab[77264]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:367
		// _ = "end of CoverTab[77264]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:367
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:367
	// _ = "end of CoverTab[77188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:367
	_go_fuzz_dep_.CoverTab[77189]++
														for _, sh := range t.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:368
		_go_fuzz_dep_.CoverTab[77265]++
															t.ctx = sh.TagConn(t.ctx, &stats.ConnTagInfo{
			RemoteAddr:	t.remoteAddr,
			LocalAddr:	t.localAddr,
		})
		connBegin := &stats.ConnBegin{
			Client: true,
		}
															sh.HandleConn(t.ctx, connBegin)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:376
		// _ = "end of CoverTab[77265]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:377
	// _ = "end of CoverTab[77189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:377
	_go_fuzz_dep_.CoverTab[77190]++
														t.channelzID, err = channelz.RegisterNormalSocket(t, opts.ChannelzParentID, fmt.Sprintf("%s -> %s", t.localAddr, t.remoteAddr))
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:379
		_go_fuzz_dep_.CoverTab[77266]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:380
		// _ = "end of CoverTab[77266]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:381
		_go_fuzz_dep_.CoverTab[77267]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:381
		// _ = "end of CoverTab[77267]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:381
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:381
	// _ = "end of CoverTab[77190]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:381
	_go_fuzz_dep_.CoverTab[77191]++
														if t.keepaliveEnabled {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:382
		_go_fuzz_dep_.CoverTab[77268]++
															t.kpDormancyCond = sync.NewCond(&t.mu)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:383
		_curRoutineNum81_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:383
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum81_)
															go t.keepalive()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:384
		// _ = "end of CoverTab[77268]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:385
		_go_fuzz_dep_.CoverTab[77269]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:385
		// _ = "end of CoverTab[77269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:385
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:385
	// _ = "end of CoverTab[77191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:385
	_go_fuzz_dep_.CoverTab[77192]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:393
	readerErrCh := make(chan error, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:393
	_curRoutineNum79_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:393
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum79_)
														go t.reader(readerErrCh)
														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:395
		_go_fuzz_dep_.CoverTab[77270]++
															if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:396
			_go_fuzz_dep_.CoverTab[77272]++
																err = <-readerErrCh
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:397
			// _ = "end of CoverTab[77272]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:398
			_go_fuzz_dep_.CoverTab[77273]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:398
			// _ = "end of CoverTab[77273]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:398
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:398
		// _ = "end of CoverTab[77270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:398
		_go_fuzz_dep_.CoverTab[77271]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:399
			_go_fuzz_dep_.CoverTab[77274]++
																t.Close(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:400
			// _ = "end of CoverTab[77274]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:401
			_go_fuzz_dep_.CoverTab[77275]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:401
			// _ = "end of CoverTab[77275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:401
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:401
		// _ = "end of CoverTab[77271]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:402
	// _ = "end of CoverTab[77192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:402
	_go_fuzz_dep_.CoverTab[77193]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:405
	n, err := t.conn.Write(clientPreface)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:406
		_go_fuzz_dep_.CoverTab[77276]++
															err = connectionErrorf(true, err, "transport: failed to write client preface: %v", err)
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:408
		// _ = "end of CoverTab[77276]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:409
		_go_fuzz_dep_.CoverTab[77277]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:409
		// _ = "end of CoverTab[77277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:409
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:409
	// _ = "end of CoverTab[77193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:409
	_go_fuzz_dep_.CoverTab[77194]++
														if n != len(clientPreface) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:410
		_go_fuzz_dep_.CoverTab[77278]++
															err = connectionErrorf(true, nil, "transport: preface mismatch, wrote %d bytes; want %d", n, len(clientPreface))
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:412
		// _ = "end of CoverTab[77278]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:413
		_go_fuzz_dep_.CoverTab[77279]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:413
		// _ = "end of CoverTab[77279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:413
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:413
	// _ = "end of CoverTab[77194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:413
	_go_fuzz_dep_.CoverTab[77195]++
														var ss []http2.Setting

														if t.initialWindowSize != defaultWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:416
		_go_fuzz_dep_.CoverTab[77280]++
															ss = append(ss, http2.Setting{
			ID:	http2.SettingInitialWindowSize,
			Val:	uint32(t.initialWindowSize),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:420
		// _ = "end of CoverTab[77280]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:421
		_go_fuzz_dep_.CoverTab[77281]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:421
		// _ = "end of CoverTab[77281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:421
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:421
	// _ = "end of CoverTab[77195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:421
	_go_fuzz_dep_.CoverTab[77196]++
														if opts.MaxHeaderListSize != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:422
		_go_fuzz_dep_.CoverTab[77282]++
															ss = append(ss, http2.Setting{
			ID:	http2.SettingMaxHeaderListSize,
			Val:	*opts.MaxHeaderListSize,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:426
		// _ = "end of CoverTab[77282]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:427
		_go_fuzz_dep_.CoverTab[77283]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:427
		// _ = "end of CoverTab[77283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:427
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:427
	// _ = "end of CoverTab[77196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:427
	_go_fuzz_dep_.CoverTab[77197]++
														err = t.framer.fr.WriteSettings(ss...)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:429
		_go_fuzz_dep_.CoverTab[77284]++
															err = connectionErrorf(true, err, "transport: failed to write initial settings frame: %v", err)
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:431
		// _ = "end of CoverTab[77284]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:432
		_go_fuzz_dep_.CoverTab[77285]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:432
		// _ = "end of CoverTab[77285]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:432
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:432
	// _ = "end of CoverTab[77197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:432
	_go_fuzz_dep_.CoverTab[77198]++

														if delta := uint32(icwz - defaultWindowSize); delta > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:434
		_go_fuzz_dep_.CoverTab[77286]++
															if err := t.framer.fr.WriteWindowUpdate(0, delta); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:435
			_go_fuzz_dep_.CoverTab[77287]++
																err = connectionErrorf(true, err, "transport: failed to write window update: %v", err)
																return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:437
			// _ = "end of CoverTab[77287]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:438
			_go_fuzz_dep_.CoverTab[77288]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:438
			// _ = "end of CoverTab[77288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:438
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:438
		// _ = "end of CoverTab[77286]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:439
		_go_fuzz_dep_.CoverTab[77289]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:439
		// _ = "end of CoverTab[77289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:439
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:439
	// _ = "end of CoverTab[77198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:439
	_go_fuzz_dep_.CoverTab[77199]++

														t.connectionID = atomic.AddUint64(&clientConnectionCounter, 1)

														if err := t.framer.writer.Flush(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:443
		_go_fuzz_dep_.CoverTab[77290]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:444
		// _ = "end of CoverTab[77290]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
		_go_fuzz_dep_.CoverTab[77291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
		// _ = "end of CoverTab[77291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
	// _ = "end of CoverTab[77199]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
	_go_fuzz_dep_.CoverTab[77200]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
	_curRoutineNum80_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:445
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum80_)
														go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:446
		_go_fuzz_dep_.CoverTab[77292]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:446
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:446
			_go_fuzz_dep_.CoverTab[77293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:446
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum80_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:446
			// _ = "end of CoverTab[77293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:446
		}()
															t.loopy = newLoopyWriter(clientSide, t.framer, t.controlBuf, t.bdpEst, t.conn)
															t.loopy.run()
															close(t.writerDone)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:449
		// _ = "end of CoverTab[77292]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:450
	// _ = "end of CoverTab[77200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:450
	_go_fuzz_dep_.CoverTab[77201]++
														return t, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:451
	// _ = "end of CoverTab[77201]"
}

func (t *http2Client) newStream(ctx context.Context, callHdr *CallHdr) *Stream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:454
	_go_fuzz_dep_.CoverTab[77294]++

														s := &Stream{
		ct:		t,
		done:		make(chan struct{}),
		method:		callHdr.Method,
		sendCompress:	callHdr.SendCompress,
		buf:		newRecvBuffer(),
		headerChan:	make(chan struct{}),
		contentSubtype:	callHdr.ContentSubtype,
		doneFunc:	callHdr.DoneFunc,
	}
	s.wq = newWriteQuota(defaultWriteQuota, s.done)
	s.requestRead = func(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:467
		_go_fuzz_dep_.CoverTab[77297]++
															t.adjustWindow(s, uint32(n))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:468
		// _ = "end of CoverTab[77297]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:469
	// _ = "end of CoverTab[77294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:469
	_go_fuzz_dep_.CoverTab[77295]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:473
	s.ctx = ctx
	s.trReader = &transportReader{
		reader: &recvBufferReader{
			ctx:		s.ctx,
			ctxDone:	s.ctx.Done(),
			recv:		s.buf,
			closeStream: func(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:479
				_go_fuzz_dep_.CoverTab[77298]++
																	t.CloseStream(s, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:480
				// _ = "end of CoverTab[77298]"
			},
			freeBuffer:	t.bufferPool.put,
		},
		windowHandler: func(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:484
			_go_fuzz_dep_.CoverTab[77299]++
																t.updateWindow(s, uint32(n))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:485
			// _ = "end of CoverTab[77299]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:487
	// _ = "end of CoverTab[77295]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:487
	_go_fuzz_dep_.CoverTab[77296]++
														return s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:488
	// _ = "end of CoverTab[77296]"
}

func (t *http2Client) getPeer() *peer.Peer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:491
	_go_fuzz_dep_.CoverTab[77300]++
														return &peer.Peer{
		Addr:		t.remoteAddr,
		AuthInfo:	t.authInfo,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:495
	// _ = "end of CoverTab[77300]"
}

func (t *http2Client) createHeaderFields(ctx context.Context, callHdr *CallHdr) ([]hpack.HeaderField, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:498
	_go_fuzz_dep_.CoverTab[77301]++
														aud := t.createAudience(callHdr)
														ri := credentials.RequestInfo{
		Method:		callHdr.Method,
		AuthInfo:	t.authInfo,
	}
	ctxWithRequestInfo := icredentials.NewRequestInfoContext(ctx, ri)
	authData, err := t.getTrAuthData(ctxWithRequestInfo, aud)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:506
		_go_fuzz_dep_.CoverTab[77314]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:507
		// _ = "end of CoverTab[77314]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:508
		_go_fuzz_dep_.CoverTab[77315]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:508
		// _ = "end of CoverTab[77315]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:508
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:508
	// _ = "end of CoverTab[77301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:508
	_go_fuzz_dep_.CoverTab[77302]++
														callAuthData, err := t.getCallAuthData(ctxWithRequestInfo, aud, callHdr)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:510
		_go_fuzz_dep_.CoverTab[77316]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:511
		// _ = "end of CoverTab[77316]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:512
		_go_fuzz_dep_.CoverTab[77317]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:512
		// _ = "end of CoverTab[77317]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:512
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:512
	// _ = "end of CoverTab[77302]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:512
	_go_fuzz_dep_.CoverTab[77303]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:516
	hfLen := 7
	hfLen += len(authData) + len(callAuthData)
	headerFields := make([]hpack.HeaderField, 0, hfLen)
	headerFields = append(headerFields, hpack.HeaderField{Name: ":method", Value: "POST"})
	headerFields = append(headerFields, hpack.HeaderField{Name: ":scheme", Value: t.scheme})
	headerFields = append(headerFields, hpack.HeaderField{Name: ":path", Value: callHdr.Method})
	headerFields = append(headerFields, hpack.HeaderField{Name: ":authority", Value: callHdr.Host})
	headerFields = append(headerFields, hpack.HeaderField{Name: "content-type", Value: grpcutil.ContentType(callHdr.ContentSubtype)})
	headerFields = append(headerFields, hpack.HeaderField{Name: "user-agent", Value: t.userAgent})
	headerFields = append(headerFields, hpack.HeaderField{Name: "te", Value: "trailers"})
	if callHdr.PreviousAttempts > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:526
		_go_fuzz_dep_.CoverTab[77318]++
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-previous-rpc-attempts", Value: strconv.Itoa(callHdr.PreviousAttempts)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:527
		// _ = "end of CoverTab[77318]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:528
		_go_fuzz_dep_.CoverTab[77319]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:528
		// _ = "end of CoverTab[77319]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:528
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:528
	// _ = "end of CoverTab[77303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:528
	_go_fuzz_dep_.CoverTab[77304]++

														registeredCompressors := t.registeredCompressors
														if callHdr.SendCompress != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:531
		_go_fuzz_dep_.CoverTab[77320]++
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-encoding", Value: callHdr.SendCompress})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:536
		if !grpcutil.IsCompressorNameRegistered(callHdr.SendCompress) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:536
			_go_fuzz_dep_.CoverTab[77321]++
																if registeredCompressors != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:537
				_go_fuzz_dep_.CoverTab[77323]++
																	registeredCompressors += ","
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:538
				// _ = "end of CoverTab[77323]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:539
				_go_fuzz_dep_.CoverTab[77324]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:539
				// _ = "end of CoverTab[77324]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:539
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:539
			// _ = "end of CoverTab[77321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:539
			_go_fuzz_dep_.CoverTab[77322]++
																registeredCompressors += callHdr.SendCompress
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:540
			// _ = "end of CoverTab[77322]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:541
			_go_fuzz_dep_.CoverTab[77325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:541
			// _ = "end of CoverTab[77325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:541
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:541
		// _ = "end of CoverTab[77320]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:542
		_go_fuzz_dep_.CoverTab[77326]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:542
		// _ = "end of CoverTab[77326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:542
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:542
	// _ = "end of CoverTab[77304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:542
	_go_fuzz_dep_.CoverTab[77305]++

														if registeredCompressors != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:544
		_go_fuzz_dep_.CoverTab[77327]++
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-accept-encoding", Value: registeredCompressors})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:545
		// _ = "end of CoverTab[77327]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:546
		_go_fuzz_dep_.CoverTab[77328]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:546
		// _ = "end of CoverTab[77328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:546
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:546
	// _ = "end of CoverTab[77305]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:546
	_go_fuzz_dep_.CoverTab[77306]++
														if dl, ok := ctx.Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:547
		_go_fuzz_dep_.CoverTab[77329]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:550
		timeout := time.Until(dl)
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-timeout", Value: grpcutil.EncodeDuration(timeout)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:551
		// _ = "end of CoverTab[77329]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:552
		_go_fuzz_dep_.CoverTab[77330]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:552
		// _ = "end of CoverTab[77330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:552
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:552
	// _ = "end of CoverTab[77306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:552
	_go_fuzz_dep_.CoverTab[77307]++
														for k, v := range authData {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:553
		_go_fuzz_dep_.CoverTab[77331]++
															headerFields = append(headerFields, hpack.HeaderField{Name: k, Value: encodeMetadataHeader(k, v)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:554
		// _ = "end of CoverTab[77331]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:555
	// _ = "end of CoverTab[77307]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:555
	_go_fuzz_dep_.CoverTab[77308]++
														for k, v := range callAuthData {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:556
		_go_fuzz_dep_.CoverTab[77332]++
															headerFields = append(headerFields, hpack.HeaderField{Name: k, Value: encodeMetadataHeader(k, v)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:557
		// _ = "end of CoverTab[77332]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:558
	// _ = "end of CoverTab[77308]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:558
	_go_fuzz_dep_.CoverTab[77309]++
														if b := stats.OutgoingTags(ctx); b != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:559
		_go_fuzz_dep_.CoverTab[77333]++
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-tags-bin", Value: encodeBinHeader(b)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:560
		// _ = "end of CoverTab[77333]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:561
		_go_fuzz_dep_.CoverTab[77334]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:561
		// _ = "end of CoverTab[77334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:561
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:561
	// _ = "end of CoverTab[77309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:561
	_go_fuzz_dep_.CoverTab[77310]++
														if b := stats.OutgoingTrace(ctx); b != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:562
		_go_fuzz_dep_.CoverTab[77335]++
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-trace-bin", Value: encodeBinHeader(b)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:563
		// _ = "end of CoverTab[77335]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:564
		_go_fuzz_dep_.CoverTab[77336]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:564
		// _ = "end of CoverTab[77336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:564
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:564
	// _ = "end of CoverTab[77310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:564
	_go_fuzz_dep_.CoverTab[77311]++

														if md, added, ok := metadata.FromOutgoingContextRaw(ctx); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:566
		_go_fuzz_dep_.CoverTab[77337]++
															var k string
															for k, vv := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:568
			_go_fuzz_dep_.CoverTab[77339]++

																if isReservedHeader(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:570
				_go_fuzz_dep_.CoverTab[77341]++
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:571
				// _ = "end of CoverTab[77341]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:572
				_go_fuzz_dep_.CoverTab[77342]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:572
				// _ = "end of CoverTab[77342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:572
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:572
			// _ = "end of CoverTab[77339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:572
			_go_fuzz_dep_.CoverTab[77340]++
																for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:573
				_go_fuzz_dep_.CoverTab[77343]++
																	headerFields = append(headerFields, hpack.HeaderField{Name: k, Value: encodeMetadataHeader(k, v)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:574
				// _ = "end of CoverTab[77343]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:575
			// _ = "end of CoverTab[77340]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:576
		// _ = "end of CoverTab[77337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:576
		_go_fuzz_dep_.CoverTab[77338]++
															for _, vv := range added {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:577
			_go_fuzz_dep_.CoverTab[77344]++
																for i, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:578
				_go_fuzz_dep_.CoverTab[77345]++
																	if i%2 == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:579
					_go_fuzz_dep_.CoverTab[77348]++
																		k = strings.ToLower(v)
																		continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:581
					// _ = "end of CoverTab[77348]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:582
					_go_fuzz_dep_.CoverTab[77349]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:582
					// _ = "end of CoverTab[77349]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:582
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:582
				// _ = "end of CoverTab[77345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:582
				_go_fuzz_dep_.CoverTab[77346]++

																	if isReservedHeader(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:584
					_go_fuzz_dep_.CoverTab[77350]++
																		continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:585
					// _ = "end of CoverTab[77350]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:586
					_go_fuzz_dep_.CoverTab[77351]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:586
					// _ = "end of CoverTab[77351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:586
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:586
				// _ = "end of CoverTab[77346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:586
				_go_fuzz_dep_.CoverTab[77347]++
																	headerFields = append(headerFields, hpack.HeaderField{Name: k, Value: encodeMetadataHeader(k, v)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:587
				// _ = "end of CoverTab[77347]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:588
			// _ = "end of CoverTab[77344]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:589
		// _ = "end of CoverTab[77338]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:590
		_go_fuzz_dep_.CoverTab[77352]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:590
		// _ = "end of CoverTab[77352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:590
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:590
	// _ = "end of CoverTab[77311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:590
	_go_fuzz_dep_.CoverTab[77312]++
														for k, vv := range t.md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:591
		_go_fuzz_dep_.CoverTab[77353]++
															if isReservedHeader(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:592
			_go_fuzz_dep_.CoverTab[77355]++
																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:593
			// _ = "end of CoverTab[77355]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:594
			_go_fuzz_dep_.CoverTab[77356]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:594
			// _ = "end of CoverTab[77356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:594
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:594
		// _ = "end of CoverTab[77353]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:594
		_go_fuzz_dep_.CoverTab[77354]++
															for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:595
			_go_fuzz_dep_.CoverTab[77357]++
																headerFields = append(headerFields, hpack.HeaderField{Name: k, Value: encodeMetadataHeader(k, v)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:596
			// _ = "end of CoverTab[77357]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:597
		// _ = "end of CoverTab[77354]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:598
	// _ = "end of CoverTab[77312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:598
	_go_fuzz_dep_.CoverTab[77313]++
														return headerFields, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:599
	// _ = "end of CoverTab[77313]"
}

func (t *http2Client) createAudience(callHdr *CallHdr) string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:602
	_go_fuzz_dep_.CoverTab[77358]++

														if len(t.perRPCCreds) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:604
		_go_fuzz_dep_.CoverTab[77361]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:604
		return callHdr.Creds == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:604
		// _ = "end of CoverTab[77361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:604
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:604
		_go_fuzz_dep_.CoverTab[77362]++
															return ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:605
		// _ = "end of CoverTab[77362]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:606
		_go_fuzz_dep_.CoverTab[77363]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:606
		// _ = "end of CoverTab[77363]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:606
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:606
	// _ = "end of CoverTab[77358]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:606
	_go_fuzz_dep_.CoverTab[77359]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:609
	host := strings.TrimSuffix(callHdr.Host, ":443")
	pos := strings.LastIndex(callHdr.Method, "/")
	if pos == -1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:611
		_go_fuzz_dep_.CoverTab[77364]++
															pos = len(callHdr.Method)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:612
		// _ = "end of CoverTab[77364]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:613
		_go_fuzz_dep_.CoverTab[77365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:613
		// _ = "end of CoverTab[77365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:613
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:613
	// _ = "end of CoverTab[77359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:613
	_go_fuzz_dep_.CoverTab[77360]++
														return "https://" + host + callHdr.Method[:pos]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:614
	// _ = "end of CoverTab[77360]"
}

func (t *http2Client) getTrAuthData(ctx context.Context, audience string) (map[string]string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:617
	_go_fuzz_dep_.CoverTab[77366]++
														if len(t.perRPCCreds) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:618
		_go_fuzz_dep_.CoverTab[77369]++
															return nil, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:619
		// _ = "end of CoverTab[77369]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:620
		_go_fuzz_dep_.CoverTab[77370]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:620
		// _ = "end of CoverTab[77370]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:620
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:620
	// _ = "end of CoverTab[77366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:620
	_go_fuzz_dep_.CoverTab[77367]++
														authData := map[string]string{}
														for _, c := range t.perRPCCreds {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:622
		_go_fuzz_dep_.CoverTab[77371]++
															data, err := c.GetRequestMetadata(ctx, audience)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:624
			_go_fuzz_dep_.CoverTab[77373]++
																if st, ok := status.FromError(err); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:625
				_go_fuzz_dep_.CoverTab[77375]++

																	if istatus.IsRestrictedControlPlaneCode(st) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:627
					_go_fuzz_dep_.CoverTab[77377]++
																		err = status.Errorf(codes.Internal, "transport: received per-RPC creds error with illegal status: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:628
					// _ = "end of CoverTab[77377]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:629
					_go_fuzz_dep_.CoverTab[77378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:629
					// _ = "end of CoverTab[77378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:629
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:629
				// _ = "end of CoverTab[77375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:629
				_go_fuzz_dep_.CoverTab[77376]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:630
				// _ = "end of CoverTab[77376]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:631
				_go_fuzz_dep_.CoverTab[77379]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:631
				// _ = "end of CoverTab[77379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:631
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:631
			// _ = "end of CoverTab[77373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:631
			_go_fuzz_dep_.CoverTab[77374]++

																return nil, status.Errorf(codes.Unauthenticated, "transport: per-RPC creds failed due to error: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:633
			// _ = "end of CoverTab[77374]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:634
			_go_fuzz_dep_.CoverTab[77380]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:634
			// _ = "end of CoverTab[77380]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:634
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:634
		// _ = "end of CoverTab[77371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:634
		_go_fuzz_dep_.CoverTab[77372]++
															for k, v := range data {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:635
			_go_fuzz_dep_.CoverTab[77381]++

																k = strings.ToLower(k)
																authData[k] = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:638
			// _ = "end of CoverTab[77381]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:639
		// _ = "end of CoverTab[77372]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:640
	// _ = "end of CoverTab[77367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:640
	_go_fuzz_dep_.CoverTab[77368]++
														return authData, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:641
	// _ = "end of CoverTab[77368]"
}

func (t *http2Client) getCallAuthData(ctx context.Context, audience string, callHdr *CallHdr) (map[string]string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:644
	_go_fuzz_dep_.CoverTab[77382]++
														var callAuthData map[string]string

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:649
	if callCreds := callHdr.Creds; callCreds != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:649
		_go_fuzz_dep_.CoverTab[77384]++
															if callCreds.RequireTransportSecurity() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:650
			_go_fuzz_dep_.CoverTab[77387]++
																ri, _ := credentials.RequestInfoFromContext(ctx)
																if !t.isSecure || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:652
				_go_fuzz_dep_.CoverTab[77388]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:652
				return credentials.CheckSecurityLevel(ri.AuthInfo, credentials.PrivacyAndIntegrity) != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:652
				// _ = "end of CoverTab[77388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:652
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:652
				_go_fuzz_dep_.CoverTab[77389]++
																	return nil, status.Error(codes.Unauthenticated, "transport: cannot send secure credentials on an insecure connection")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:653
				// _ = "end of CoverTab[77389]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:654
				_go_fuzz_dep_.CoverTab[77390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:654
				// _ = "end of CoverTab[77390]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:654
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:654
			// _ = "end of CoverTab[77387]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:655
			_go_fuzz_dep_.CoverTab[77391]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:655
			// _ = "end of CoverTab[77391]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:655
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:655
		// _ = "end of CoverTab[77384]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:655
		_go_fuzz_dep_.CoverTab[77385]++
															data, err := callCreds.GetRequestMetadata(ctx, audience)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:657
			_go_fuzz_dep_.CoverTab[77392]++
																if st, ok := status.FromError(err); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:658
				_go_fuzz_dep_.CoverTab[77394]++

																	if istatus.IsRestrictedControlPlaneCode(st) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:660
					_go_fuzz_dep_.CoverTab[77396]++
																		err = status.Errorf(codes.Internal, "transport: received per-RPC creds error with illegal status: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:661
					// _ = "end of CoverTab[77396]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:662
					_go_fuzz_dep_.CoverTab[77397]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:662
					// _ = "end of CoverTab[77397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:662
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:662
				// _ = "end of CoverTab[77394]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:662
				_go_fuzz_dep_.CoverTab[77395]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:663
				// _ = "end of CoverTab[77395]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:664
				_go_fuzz_dep_.CoverTab[77398]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:664
				// _ = "end of CoverTab[77398]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:664
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:664
			// _ = "end of CoverTab[77392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:664
			_go_fuzz_dep_.CoverTab[77393]++
																return nil, status.Errorf(codes.Internal, "transport: per-RPC creds failed due to error: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:665
			// _ = "end of CoverTab[77393]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:666
			_go_fuzz_dep_.CoverTab[77399]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:666
			// _ = "end of CoverTab[77399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:666
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:666
		// _ = "end of CoverTab[77385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:666
		_go_fuzz_dep_.CoverTab[77386]++
															callAuthData = make(map[string]string, len(data))
															for k, v := range data {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:668
			_go_fuzz_dep_.CoverTab[77400]++

																k = strings.ToLower(k)
																callAuthData[k] = v
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:671
			// _ = "end of CoverTab[77400]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:672
		// _ = "end of CoverTab[77386]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:673
		_go_fuzz_dep_.CoverTab[77401]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:673
		// _ = "end of CoverTab[77401]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:673
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:673
	// _ = "end of CoverTab[77382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:673
	_go_fuzz_dep_.CoverTab[77383]++
														return callAuthData, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:674
	// _ = "end of CoverTab[77383]"
}

// NewStreamError wraps an error and reports additional information.  Typically
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
// NewStream errors result in transparent retry, as they mean nothing went onto
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
// the wire.  However, there are two notable exceptions:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//  1. If the stream headers violate the max header list size allowed by the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//     server.  It's possible this could succeed on another transport, even if
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//     it's unlikely, but do not transparently retry.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//  2. If the credentials errored when requesting their headers.  In this case,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//     it's possible a retry can fix the problem, but indefinitely transparently
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//     retrying is not appropriate as it is likely the credentials, if they can
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:677
//     eventually succeed, would need I/O to do so.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:688
type NewStreamError struct {
	Err	error

	AllowTransparentRetry	bool
}

func (e NewStreamError) Error() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:694
	_go_fuzz_dep_.CoverTab[77402]++
														return e.Err.Error()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:695
	// _ = "end of CoverTab[77402]"
}

// NewStream creates a stream and registers it into the transport as "active"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:698
// streams.  All non-nil errors returned will be *NewStreamError.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:700
func (t *http2Client) NewStream(ctx context.Context, callHdr *CallHdr) (*Stream, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:700
	_go_fuzz_dep_.CoverTab[77403]++
														ctx = peer.NewContext(ctx, t.getPeer())

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:708
	if t.address.ServerName != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:708
		_go_fuzz_dep_.CoverTab[77413]++
															newCallHdr := *callHdr
															newCallHdr.Host = t.address.ServerName
															callHdr = &newCallHdr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:711
		// _ = "end of CoverTab[77413]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:712
		_go_fuzz_dep_.CoverTab[77414]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:712
		// _ = "end of CoverTab[77414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:712
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:712
	// _ = "end of CoverTab[77403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:712
	_go_fuzz_dep_.CoverTab[77404]++

														headerFields, err := t.createHeaderFields(ctx, callHdr)
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:715
		_go_fuzz_dep_.CoverTab[77415]++
															return nil, &NewStreamError{Err: err, AllowTransparentRetry: false}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:716
		// _ = "end of CoverTab[77415]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:717
		_go_fuzz_dep_.CoverTab[77416]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:717
		// _ = "end of CoverTab[77416]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:717
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:717
	// _ = "end of CoverTab[77404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:717
	_go_fuzz_dep_.CoverTab[77405]++
														s := t.newStream(ctx, callHdr)
														cleanup := func(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:719
		_go_fuzz_dep_.CoverTab[77417]++
															if s.swapState(streamDone) == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:720
			_go_fuzz_dep_.CoverTab[77419]++

																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:722
			// _ = "end of CoverTab[77419]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:723
			_go_fuzz_dep_.CoverTab[77420]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:723
			// _ = "end of CoverTab[77420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:723
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:723
		// _ = "end of CoverTab[77417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:723
		_go_fuzz_dep_.CoverTab[77418]++

															atomic.StoreUint32(&s.unprocessed, 1)
															s.write(recvMsg{err: err})
															close(s.done)

															if atomic.CompareAndSwapUint32(&s.headerChanClosed, 0, 1) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:729
			_go_fuzz_dep_.CoverTab[77421]++
																close(s.headerChan)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:730
			// _ = "end of CoverTab[77421]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:731
			_go_fuzz_dep_.CoverTab[77422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:731
			// _ = "end of CoverTab[77422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:731
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:731
		// _ = "end of CoverTab[77418]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:732
	// _ = "end of CoverTab[77405]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:732
	_go_fuzz_dep_.CoverTab[77406]++
														hdr := &headerFrame{
		hf:		headerFields,
		endStream:	false,
		initStream: func(id uint32) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:736
			_go_fuzz_dep_.CoverTab[77423]++
																t.mu.Lock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:740
			if t.state == closing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:740
				_go_fuzz_dep_.CoverTab[77427]++
																	t.mu.Unlock()
																	cleanup(ErrConnClosing)
																	return ErrConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:743
				// _ = "end of CoverTab[77427]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:744
				_go_fuzz_dep_.CoverTab[77428]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:744
				// _ = "end of CoverTab[77428]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:744
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:744
			// _ = "end of CoverTab[77423]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:744
			_go_fuzz_dep_.CoverTab[77424]++
																if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:745
				_go_fuzz_dep_.CoverTab[77429]++
																	atomic.AddInt64(&t.czData.streamsStarted, 1)
																	atomic.StoreInt64(&t.czData.lastStreamCreatedTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:747
				// _ = "end of CoverTab[77429]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:748
				_go_fuzz_dep_.CoverTab[77430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:748
				// _ = "end of CoverTab[77430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:748
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:748
			// _ = "end of CoverTab[77424]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:748
			_go_fuzz_dep_.CoverTab[77425]++

																if t.kpDormant {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:750
				_go_fuzz_dep_.CoverTab[77431]++
																	t.kpDormancyCond.Signal()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:751
				// _ = "end of CoverTab[77431]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:752
				_go_fuzz_dep_.CoverTab[77432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:752
				// _ = "end of CoverTab[77432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:752
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:752
			// _ = "end of CoverTab[77425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:752
			_go_fuzz_dep_.CoverTab[77426]++
																t.mu.Unlock()
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:754
			// _ = "end of CoverTab[77426]"
		},
		onOrphaned:	cleanup,
		wq:		s.wq,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:758
	// _ = "end of CoverTab[77406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:758
	_go_fuzz_dep_.CoverTab[77407]++
														firstTry := true
														var ch chan struct{}
														transportDrainRequired := false
														checkForStreamQuota := func(it interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:762
		_go_fuzz_dep_.CoverTab[77433]++
															if t.streamQuota <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:763
			_go_fuzz_dep_.CoverTab[77438]++
																if firstTry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:764
				_go_fuzz_dep_.CoverTab[77440]++
																	t.waitingStreams++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:765
				// _ = "end of CoverTab[77440]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:766
				_go_fuzz_dep_.CoverTab[77441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:766
				// _ = "end of CoverTab[77441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:766
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:766
			// _ = "end of CoverTab[77438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:766
			_go_fuzz_dep_.CoverTab[77439]++
																ch = t.streamsQuotaAvailable
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:768
			// _ = "end of CoverTab[77439]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:769
			_go_fuzz_dep_.CoverTab[77442]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:769
			// _ = "end of CoverTab[77442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:769
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:769
		// _ = "end of CoverTab[77433]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:769
		_go_fuzz_dep_.CoverTab[77434]++
															if !firstTry {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:770
			_go_fuzz_dep_.CoverTab[77443]++
																t.waitingStreams--
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:771
			// _ = "end of CoverTab[77443]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:772
			_go_fuzz_dep_.CoverTab[77444]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:772
			// _ = "end of CoverTab[77444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:772
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:772
		// _ = "end of CoverTab[77434]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:772
		_go_fuzz_dep_.CoverTab[77435]++
															t.streamQuota--
															h := it.(*headerFrame)
															h.streamID = t.nextID
															t.nextID += 2

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:780
		transportDrainRequired = t.nextID > MaxStreamID

		s.id = h.streamID
		s.fc = &inFlow{limit: uint32(t.initialWindowSize)}
		t.mu.Lock()
		if t.activeStreams == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:785
			_go_fuzz_dep_.CoverTab[77445]++
																t.mu.Unlock()
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:787
			// _ = "end of CoverTab[77445]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:788
			_go_fuzz_dep_.CoverTab[77446]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:788
			// _ = "end of CoverTab[77446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:788
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:788
		// _ = "end of CoverTab[77435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:788
		_go_fuzz_dep_.CoverTab[77436]++
															t.activeStreams[s.id] = s
															t.mu.Unlock()
															if t.streamQuota > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:791
			_go_fuzz_dep_.CoverTab[77447]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:791
			return t.waitingStreams > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:791
			// _ = "end of CoverTab[77447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:791
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:791
			_go_fuzz_dep_.CoverTab[77448]++
																select {
			case t.streamsQuotaAvailable <- struct{}{}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:793
				_go_fuzz_dep_.CoverTab[77449]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:793
				// _ = "end of CoverTab[77449]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:794
				_go_fuzz_dep_.CoverTab[77450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:794
				// _ = "end of CoverTab[77450]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:795
			// _ = "end of CoverTab[77448]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:796
			_go_fuzz_dep_.CoverTab[77451]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:796
			// _ = "end of CoverTab[77451]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:796
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:796
		// _ = "end of CoverTab[77436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:796
		_go_fuzz_dep_.CoverTab[77437]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:797
		// _ = "end of CoverTab[77437]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:798
	// _ = "end of CoverTab[77407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:798
	_go_fuzz_dep_.CoverTab[77408]++
														var hdrListSizeErr error
														checkForHeaderListSize := func(it interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:800
		_go_fuzz_dep_.CoverTab[77452]++
															if t.maxSendHeaderListSize == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:801
			_go_fuzz_dep_.CoverTab[77455]++
																return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:802
			// _ = "end of CoverTab[77455]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:803
			_go_fuzz_dep_.CoverTab[77456]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:803
			// _ = "end of CoverTab[77456]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:803
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:803
		// _ = "end of CoverTab[77452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:803
		_go_fuzz_dep_.CoverTab[77453]++
															hdrFrame := it.(*headerFrame)
															var sz int64
															for _, f := range hdrFrame.hf {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:806
			_go_fuzz_dep_.CoverTab[77457]++
																if sz += int64(f.Size()); sz > int64(*t.maxSendHeaderListSize) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:807
				_go_fuzz_dep_.CoverTab[77458]++
																	hdrListSizeErr = status.Errorf(codes.Internal, "header list size to send violates the maximum size (%d bytes) set by server", *t.maxSendHeaderListSize)
																	return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:809
				// _ = "end of CoverTab[77458]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:810
				_go_fuzz_dep_.CoverTab[77459]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:810
				// _ = "end of CoverTab[77459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:810
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:810
			// _ = "end of CoverTab[77457]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:811
		// _ = "end of CoverTab[77453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:811
		_go_fuzz_dep_.CoverTab[77454]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:812
		// _ = "end of CoverTab[77454]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:813
	// _ = "end of CoverTab[77408]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:813
	_go_fuzz_dep_.CoverTab[77409]++
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:814
		_go_fuzz_dep_.CoverTab[77460]++
															success, err := t.controlBuf.executeAndPut(func(it interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:815
			_go_fuzz_dep_.CoverTab[77465]++
																return checkForHeaderListSize(it) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:816
				_go_fuzz_dep_.CoverTab[77466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:816
				return checkForStreamQuota(it)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:816
				// _ = "end of CoverTab[77466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:816
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:816
			// _ = "end of CoverTab[77465]"
		}, hdr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:817
		// _ = "end of CoverTab[77460]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:817
		_go_fuzz_dep_.CoverTab[77461]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:818
			_go_fuzz_dep_.CoverTab[77467]++

																return nil, &NewStreamError{Err: err, AllowTransparentRetry: true}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:820
			// _ = "end of CoverTab[77467]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:821
			_go_fuzz_dep_.CoverTab[77468]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:821
			// _ = "end of CoverTab[77468]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:821
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:821
		// _ = "end of CoverTab[77461]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:821
		_go_fuzz_dep_.CoverTab[77462]++
															if success {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:822
			_go_fuzz_dep_.CoverTab[77469]++
																break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:823
			// _ = "end of CoverTab[77469]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:824
			_go_fuzz_dep_.CoverTab[77470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:824
			// _ = "end of CoverTab[77470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:824
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:824
		// _ = "end of CoverTab[77462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:824
		_go_fuzz_dep_.CoverTab[77463]++
															if hdrListSizeErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:825
			_go_fuzz_dep_.CoverTab[77471]++
																return nil, &NewStreamError{Err: hdrListSizeErr}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:826
			// _ = "end of CoverTab[77471]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:827
			_go_fuzz_dep_.CoverTab[77472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:827
			// _ = "end of CoverTab[77472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:827
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:827
		// _ = "end of CoverTab[77463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:827
		_go_fuzz_dep_.CoverTab[77464]++
															firstTry = false
															select {
		case <-ch:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:830
			_go_fuzz_dep_.CoverTab[77473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:830
			// _ = "end of CoverTab[77473]"
		case <-ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:831
			_go_fuzz_dep_.CoverTab[77474]++
																return nil, &NewStreamError{Err: ContextErr(ctx.Err())}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:832
			// _ = "end of CoverTab[77474]"
		case <-t.goAway:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:833
			_go_fuzz_dep_.CoverTab[77475]++
																return nil, &NewStreamError{Err: errStreamDrain, AllowTransparentRetry: true}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:834
			// _ = "end of CoverTab[77475]"
		case <-t.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:835
			_go_fuzz_dep_.CoverTab[77476]++
																return nil, &NewStreamError{Err: ErrConnClosing, AllowTransparentRetry: true}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:836
			// _ = "end of CoverTab[77476]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:837
		// _ = "end of CoverTab[77464]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:838
	// _ = "end of CoverTab[77409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:838
	_go_fuzz_dep_.CoverTab[77410]++
														if len(t.statsHandlers) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:839
		_go_fuzz_dep_.CoverTab[77477]++
															header, ok := metadata.FromOutgoingContext(ctx)
															if ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:841
			_go_fuzz_dep_.CoverTab[77479]++
																header.Set("user-agent", t.userAgent)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:842
			// _ = "end of CoverTab[77479]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:843
			_go_fuzz_dep_.CoverTab[77480]++
																header = metadata.Pairs("user-agent", t.userAgent)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:844
			// _ = "end of CoverTab[77480]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:845
		// _ = "end of CoverTab[77477]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:845
		_go_fuzz_dep_.CoverTab[77478]++
															for _, sh := range t.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:846
			_go_fuzz_dep_.CoverTab[77481]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:850
			outHeader := &stats.OutHeader{
				Client:		true,
				FullMethod:	callHdr.Method,
				RemoteAddr:	t.remoteAddr,
				LocalAddr:	t.localAddr,
				Compression:	callHdr.SendCompress,
				Header:		header,
			}
																sh.HandleRPC(s.ctx, outHeader)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:858
			// _ = "end of CoverTab[77481]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:859
		// _ = "end of CoverTab[77478]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:860
		_go_fuzz_dep_.CoverTab[77482]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:860
		// _ = "end of CoverTab[77482]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:860
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:860
	// _ = "end of CoverTab[77410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:860
	_go_fuzz_dep_.CoverTab[77411]++
														if transportDrainRequired {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:861
		_go_fuzz_dep_.CoverTab[77483]++
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:862
			_go_fuzz_dep_.CoverTab[77485]++
																logger.Infof("transport: t.nextID > MaxStreamID. Draining")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:863
			// _ = "end of CoverTab[77485]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:864
			_go_fuzz_dep_.CoverTab[77486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:864
			// _ = "end of CoverTab[77486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:864
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:864
		// _ = "end of CoverTab[77483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:864
		_go_fuzz_dep_.CoverTab[77484]++
															t.GracefulClose()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:865
		// _ = "end of CoverTab[77484]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:866
		_go_fuzz_dep_.CoverTab[77487]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:866
		// _ = "end of CoverTab[77487]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:866
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:866
	// _ = "end of CoverTab[77411]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:866
	_go_fuzz_dep_.CoverTab[77412]++
														return s, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:867
	// _ = "end of CoverTab[77412]"
}

// CloseStream clears the footprint of a stream when the stream is not needed any more.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:870
// This must not be executed in reader's goroutine.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:872
func (t *http2Client) CloseStream(s *Stream, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:872
	_go_fuzz_dep_.CoverTab[77488]++
														var (
		rst	bool
		rstCode	http2.ErrCode
	)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:877
		_go_fuzz_dep_.CoverTab[77490]++
															rst = true
															rstCode = http2.ErrCodeCancel
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:879
		// _ = "end of CoverTab[77490]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:880
		_go_fuzz_dep_.CoverTab[77491]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:880
		// _ = "end of CoverTab[77491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:880
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:880
	// _ = "end of CoverTab[77488]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:880
	_go_fuzz_dep_.CoverTab[77489]++
														t.closeStream(s, err, rst, rstCode, status.Convert(err), nil, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:881
	// _ = "end of CoverTab[77489]"
}

func (t *http2Client) closeStream(s *Stream, err error, rst bool, rstCode http2.ErrCode, st *status.Status, mdata map[string][]string, eosReceived bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:884
	_go_fuzz_dep_.CoverTab[77492]++

														if s.swapState(streamDone) == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:886
		_go_fuzz_dep_.CoverTab[77499]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:889
		<-s.done
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:890
		// _ = "end of CoverTab[77499]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:891
		_go_fuzz_dep_.CoverTab[77500]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:891
		// _ = "end of CoverTab[77500]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:891
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:891
	// _ = "end of CoverTab[77492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:891
	_go_fuzz_dep_.CoverTab[77493]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:895
	s.status = st
	if len(mdata) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:896
		_go_fuzz_dep_.CoverTab[77501]++
															s.trailer = mdata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:897
		// _ = "end of CoverTab[77501]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:898
		_go_fuzz_dep_.CoverTab[77502]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:898
		// _ = "end of CoverTab[77502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:898
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:898
	// _ = "end of CoverTab[77493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:898
	_go_fuzz_dep_.CoverTab[77494]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:899
		_go_fuzz_dep_.CoverTab[77503]++

															s.write(recvMsg{err: err})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:901
		// _ = "end of CoverTab[77503]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:902
		_go_fuzz_dep_.CoverTab[77504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:902
		// _ = "end of CoverTab[77504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:902
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:902
	// _ = "end of CoverTab[77494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:902
	_go_fuzz_dep_.CoverTab[77495]++

														if atomic.CompareAndSwapUint32(&s.headerChanClosed, 0, 1) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:904
		_go_fuzz_dep_.CoverTab[77505]++
															s.noHeaders = true
															close(s.headerChan)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:906
		// _ = "end of CoverTab[77505]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:907
		_go_fuzz_dep_.CoverTab[77506]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:907
		// _ = "end of CoverTab[77506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:907
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:907
	// _ = "end of CoverTab[77495]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:907
	_go_fuzz_dep_.CoverTab[77496]++
														cleanup := &cleanupStream{
		streamID:	s.id,
		onWrite: func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:910
			_go_fuzz_dep_.CoverTab[77507]++
																t.mu.Lock()
																if t.activeStreams != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:912
				_go_fuzz_dep_.CoverTab[77509]++
																	delete(t.activeStreams, s.id)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:913
				// _ = "end of CoverTab[77509]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:914
				_go_fuzz_dep_.CoverTab[77510]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:914
				// _ = "end of CoverTab[77510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:914
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:914
			// _ = "end of CoverTab[77507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:914
			_go_fuzz_dep_.CoverTab[77508]++
																t.mu.Unlock()
																if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:916
				_go_fuzz_dep_.CoverTab[77511]++
																	if eosReceived {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:917
					_go_fuzz_dep_.CoverTab[77512]++
																		atomic.AddInt64(&t.czData.streamsSucceeded, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:918
					// _ = "end of CoverTab[77512]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:919
					_go_fuzz_dep_.CoverTab[77513]++
																		atomic.AddInt64(&t.czData.streamsFailed, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:920
					// _ = "end of CoverTab[77513]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:921
				// _ = "end of CoverTab[77511]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:922
				_go_fuzz_dep_.CoverTab[77514]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:922
				// _ = "end of CoverTab[77514]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:922
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:922
			// _ = "end of CoverTab[77508]"
		},
		rst:		rst,
		rstCode:	rstCode,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:926
	// _ = "end of CoverTab[77496]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:926
	_go_fuzz_dep_.CoverTab[77497]++
														addBackStreamQuota := func(interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:927
		_go_fuzz_dep_.CoverTab[77515]++
															t.streamQuota++
															if t.streamQuota > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:929
			_go_fuzz_dep_.CoverTab[77517]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:929
			return t.waitingStreams > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:929
			// _ = "end of CoverTab[77517]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:929
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:929
			_go_fuzz_dep_.CoverTab[77518]++
																select {
			case t.streamsQuotaAvailable <- struct{}{}:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:931
				_go_fuzz_dep_.CoverTab[77519]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:931
				// _ = "end of CoverTab[77519]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:932
				_go_fuzz_dep_.CoverTab[77520]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:932
				// _ = "end of CoverTab[77520]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:933
			// _ = "end of CoverTab[77518]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:934
			_go_fuzz_dep_.CoverTab[77521]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:934
			// _ = "end of CoverTab[77521]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:934
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:934
		// _ = "end of CoverTab[77515]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:934
		_go_fuzz_dep_.CoverTab[77516]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:935
		// _ = "end of CoverTab[77516]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:936
	// _ = "end of CoverTab[77497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:936
	_go_fuzz_dep_.CoverTab[77498]++
														t.controlBuf.executeAndPut(addBackStreamQuota, cleanup)

														close(s.done)
														if s.doneFunc != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:940
		_go_fuzz_dep_.CoverTab[77522]++
															s.doneFunc()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:941
		// _ = "end of CoverTab[77522]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:942
		_go_fuzz_dep_.CoverTab[77523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:942
		// _ = "end of CoverTab[77523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:942
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:942
	// _ = "end of CoverTab[77498]"
}

// Close kicks off the shutdown process of the transport. This should be called
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:945
// only once on a transport. Once it is called, the transport should not be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:945
// accessed any more.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:948
func (t *http2Client) Close(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:948
	_go_fuzz_dep_.CoverTab[77524]++
														t.mu.Lock()

														if t.state == closing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:951
		_go_fuzz_dep_.CoverTab[77531]++
															t.mu.Unlock()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:953
		// _ = "end of CoverTab[77531]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:954
		_go_fuzz_dep_.CoverTab[77532]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:954
		// _ = "end of CoverTab[77532]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:954
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:954
	// _ = "end of CoverTab[77524]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:954
	_go_fuzz_dep_.CoverTab[77525]++
														if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:955
		_go_fuzz_dep_.CoverTab[77533]++
															logger.Infof("transport: closing: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:956
		// _ = "end of CoverTab[77533]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:957
		_go_fuzz_dep_.CoverTab[77534]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:957
		// _ = "end of CoverTab[77534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:957
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:957
	// _ = "end of CoverTab[77525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:957
	_go_fuzz_dep_.CoverTab[77526]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:960
	if t.state != draining {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:960
		_go_fuzz_dep_.CoverTab[77535]++
															t.onClose(GoAwayInvalid)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:961
		// _ = "end of CoverTab[77535]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:962
		_go_fuzz_dep_.CoverTab[77536]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:962
		// _ = "end of CoverTab[77536]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:962
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:962
	// _ = "end of CoverTab[77526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:962
	_go_fuzz_dep_.CoverTab[77527]++
														t.state = closing
														streams := t.activeStreams
														t.activeStreams = nil
														if t.kpDormant {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:966
		_go_fuzz_dep_.CoverTab[77537]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:969
		t.kpDormancyCond.Signal()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:969
		// _ = "end of CoverTab[77537]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:970
		_go_fuzz_dep_.CoverTab[77538]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:970
		// _ = "end of CoverTab[77538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:970
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:970
	// _ = "end of CoverTab[77527]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:970
	_go_fuzz_dep_.CoverTab[77528]++
														t.mu.Unlock()
														t.controlBuf.finish()
														t.cancel()
														t.conn.Close()
														channelz.RemoveEntry(t.channelzID)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:978
	_, goAwayDebugMessage := t.GetGoAwayReason()

	var st *status.Status
	if len(goAwayDebugMessage) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:981
		_go_fuzz_dep_.CoverTab[77539]++
															st = status.Newf(codes.Unavailable, "closing transport due to: %v, received prior goaway: %v", err, goAwayDebugMessage)
															err = st.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:983
		// _ = "end of CoverTab[77539]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:984
		_go_fuzz_dep_.CoverTab[77540]++
															st = status.New(codes.Unavailable, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:985
		// _ = "end of CoverTab[77540]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:986
	// _ = "end of CoverTab[77528]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:986
	_go_fuzz_dep_.CoverTab[77529]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:989
	for _, s := range streams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:989
		_go_fuzz_dep_.CoverTab[77541]++
															t.closeStream(s, err, false, http2.ErrCodeNo, st, nil, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:990
		// _ = "end of CoverTab[77541]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:991
	// _ = "end of CoverTab[77529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:991
	_go_fuzz_dep_.CoverTab[77530]++
														for _, sh := range t.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:992
		_go_fuzz_dep_.CoverTab[77542]++
															connEnd := &stats.ConnEnd{
			Client: true,
		}
															sh.HandleConn(t.ctx, connEnd)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:996
		// _ = "end of CoverTab[77542]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:997
	// _ = "end of CoverTab[77530]"
}

// GracefulClose sets the state to draining, which prevents new streams from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1000
// being created and causes the transport to be closed when the last active
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1000
// stream is closed.  If there are no active streams, the transport is closed
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1000
// immediately.  This does nothing if the transport is already draining or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1000
// closing.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1005
func (t *http2Client) GracefulClose() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1005
	_go_fuzz_dep_.CoverTab[77543]++
														t.mu.Lock()

														if t.state == draining || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1008
		_go_fuzz_dep_.CoverTab[77547]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1008
		return t.state == closing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1008
		// _ = "end of CoverTab[77547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1008
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1008
		_go_fuzz_dep_.CoverTab[77548]++
															t.mu.Unlock()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1010
		// _ = "end of CoverTab[77548]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1011
		_go_fuzz_dep_.CoverTab[77549]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1011
		// _ = "end of CoverTab[77549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1011
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1011
	// _ = "end of CoverTab[77543]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1011
	_go_fuzz_dep_.CoverTab[77544]++
														if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1012
		_go_fuzz_dep_.CoverTab[77550]++
															logger.Infof("transport: GracefulClose called")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1013
		// _ = "end of CoverTab[77550]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1014
		_go_fuzz_dep_.CoverTab[77551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1014
		// _ = "end of CoverTab[77551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1014
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1014
	// _ = "end of CoverTab[77544]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1014
	_go_fuzz_dep_.CoverTab[77545]++
														t.onClose(GoAwayInvalid)
														t.state = draining
														active := len(t.activeStreams)
														t.mu.Unlock()
														if active == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1019
		_go_fuzz_dep_.CoverTab[77552]++
															t.Close(connectionErrorf(true, nil, "no active streams left to process while draining"))
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1021
		// _ = "end of CoverTab[77552]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1022
		_go_fuzz_dep_.CoverTab[77553]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1022
		// _ = "end of CoverTab[77553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1022
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1022
	// _ = "end of CoverTab[77545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1022
	_go_fuzz_dep_.CoverTab[77546]++
														t.controlBuf.put(&incomingGoAway{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1023
	// _ = "end of CoverTab[77546]"
}

// Write formats the data into HTTP2 data frame(s) and sends it out. The caller
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1026
// should proceed only if Write returns nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1028
func (t *http2Client) Write(s *Stream, hdr []byte, data []byte, opts *Options) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1028
	_go_fuzz_dep_.CoverTab[77554]++
														if opts.Last {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1029
		_go_fuzz_dep_.CoverTab[77557]++

															if !s.compareAndSwapState(streamActive, streamWriteDone) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1031
			_go_fuzz_dep_.CoverTab[77558]++
																return errStreamDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1032
			// _ = "end of CoverTab[77558]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1033
			_go_fuzz_dep_.CoverTab[77559]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1033
			// _ = "end of CoverTab[77559]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1033
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1033
		// _ = "end of CoverTab[77557]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1034
		_go_fuzz_dep_.CoverTab[77560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1034
		if s.getState() != streamActive {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1034
			_go_fuzz_dep_.CoverTab[77561]++
																return errStreamDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1035
			// _ = "end of CoverTab[77561]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
			_go_fuzz_dep_.CoverTab[77562]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
			// _ = "end of CoverTab[77562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
		// _ = "end of CoverTab[77560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
	// _ = "end of CoverTab[77554]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1036
	_go_fuzz_dep_.CoverTab[77555]++
														df := &dataFrame{
		streamID:	s.id,
		endStream:	opts.Last,
		h:		hdr,
		d:		data,
	}
	if hdr != nil || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1043
		_go_fuzz_dep_.CoverTab[77563]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1043
		return data != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1043
		// _ = "end of CoverTab[77563]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1043
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1043
		_go_fuzz_dep_.CoverTab[77564]++
															if err := s.wq.get(int32(len(hdr) + len(data))); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1044
			_go_fuzz_dep_.CoverTab[77565]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1045
			// _ = "end of CoverTab[77565]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1046
			_go_fuzz_dep_.CoverTab[77566]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1046
			// _ = "end of CoverTab[77566]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1046
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1046
		// _ = "end of CoverTab[77564]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1047
		_go_fuzz_dep_.CoverTab[77567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1047
		// _ = "end of CoverTab[77567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1047
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1047
	// _ = "end of CoverTab[77555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1047
	_go_fuzz_dep_.CoverTab[77556]++
														return t.controlBuf.put(df)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1048
	// _ = "end of CoverTab[77556]"
}

func (t *http2Client) getStream(f http2.Frame) *Stream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1051
	_go_fuzz_dep_.CoverTab[77568]++
														t.mu.Lock()
														s := t.activeStreams[f.Header().StreamID]
														t.mu.Unlock()
														return s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1055
	// _ = "end of CoverTab[77568]"
}

// adjustWindow sends out extra window update over the initial window size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1058
// of stream if the application is requesting data larger in size than
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1058
// the window.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1061
func (t *http2Client) adjustWindow(s *Stream, n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1061
	_go_fuzz_dep_.CoverTab[77569]++
														if w := s.fc.maybeAdjust(n); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1062
		_go_fuzz_dep_.CoverTab[77570]++
															t.controlBuf.put(&outgoingWindowUpdate{streamID: s.id, increment: w})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1063
		// _ = "end of CoverTab[77570]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1064
		_go_fuzz_dep_.CoverTab[77571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1064
		// _ = "end of CoverTab[77571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1064
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1064
	// _ = "end of CoverTab[77569]"
}

// updateWindow adjusts the inbound quota for the stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1067
// Window updates will be sent out when the cumulative quota
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1067
// exceeds the corresponding threshold.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1070
func (t *http2Client) updateWindow(s *Stream, n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1070
	_go_fuzz_dep_.CoverTab[77572]++
														if w := s.fc.onRead(n); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1071
		_go_fuzz_dep_.CoverTab[77573]++
															t.controlBuf.put(&outgoingWindowUpdate{streamID: s.id, increment: w})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1072
		// _ = "end of CoverTab[77573]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1073
		_go_fuzz_dep_.CoverTab[77574]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1073
		// _ = "end of CoverTab[77574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1073
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1073
	// _ = "end of CoverTab[77572]"
}

// updateFlowControl updates the incoming flow control windows
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1076
// for the transport and the stream based on the current bdp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1076
// estimation.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1079
func (t *http2Client) updateFlowControl(n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1079
	_go_fuzz_dep_.CoverTab[77575]++
														updateIWS := func(interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1080
		_go_fuzz_dep_.CoverTab[77577]++
															t.initialWindowSize = int32(n)
															t.mu.Lock()
															for _, s := range t.activeStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1083
			_go_fuzz_dep_.CoverTab[77579]++
																s.fc.newLimit(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1084
			// _ = "end of CoverTab[77579]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1085
		// _ = "end of CoverTab[77577]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1085
		_go_fuzz_dep_.CoverTab[77578]++
															t.mu.Unlock()
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1087
		// _ = "end of CoverTab[77578]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1088
	// _ = "end of CoverTab[77575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1088
	_go_fuzz_dep_.CoverTab[77576]++
														t.controlBuf.executeAndPut(updateIWS, &outgoingWindowUpdate{streamID: 0, increment: t.fc.newLimit(n)})
														t.controlBuf.put(&outgoingSettings{
		ss: []http2.Setting{
			{
				ID:	http2.SettingInitialWindowSize,
				Val:	n,
			},
		},
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1097
	// _ = "end of CoverTab[77576]"
}

func (t *http2Client) handleData(f *http2.DataFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1100
	_go_fuzz_dep_.CoverTab[77580]++
														size := f.Header().Length
														var sendBDPPing bool
														if t.bdpEst != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1103
		_go_fuzz_dep_.CoverTab[77586]++
															sendBDPPing = t.bdpEst.add(size)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1104
		// _ = "end of CoverTab[77586]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1105
		_go_fuzz_dep_.CoverTab[77587]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1105
		// _ = "end of CoverTab[77587]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1105
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1105
	// _ = "end of CoverTab[77580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1105
	_go_fuzz_dep_.CoverTab[77581]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1115
	if w := t.fc.onData(size); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1115
		_go_fuzz_dep_.CoverTab[77588]++
															t.controlBuf.put(&outgoingWindowUpdate{
			streamID:	0,
			increment:	w,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1119
		// _ = "end of CoverTab[77588]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1120
		_go_fuzz_dep_.CoverTab[77589]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1120
		// _ = "end of CoverTab[77589]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1120
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1120
	// _ = "end of CoverTab[77581]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1120
	_go_fuzz_dep_.CoverTab[77582]++
														if sendBDPPing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1121
		_go_fuzz_dep_.CoverTab[77590]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1125
		if w := t.fc.reset(); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1125
			_go_fuzz_dep_.CoverTab[77592]++
																t.controlBuf.put(&outgoingWindowUpdate{
				streamID:	0,
				increment:	w,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1129
			// _ = "end of CoverTab[77592]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1130
			_go_fuzz_dep_.CoverTab[77593]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1130
			// _ = "end of CoverTab[77593]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1130
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1130
		// _ = "end of CoverTab[77590]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1130
		_go_fuzz_dep_.CoverTab[77591]++

															t.controlBuf.put(bdpPing)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1132
		// _ = "end of CoverTab[77591]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1133
		_go_fuzz_dep_.CoverTab[77594]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1133
		// _ = "end of CoverTab[77594]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1133
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1133
	// _ = "end of CoverTab[77582]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1133
	_go_fuzz_dep_.CoverTab[77583]++

														s := t.getStream(f)
														if s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1136
		_go_fuzz_dep_.CoverTab[77595]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1137
		// _ = "end of CoverTab[77595]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1138
		_go_fuzz_dep_.CoverTab[77596]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1138
		// _ = "end of CoverTab[77596]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1138
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1138
	// _ = "end of CoverTab[77583]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1138
	_go_fuzz_dep_.CoverTab[77584]++
														if size > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1139
		_go_fuzz_dep_.CoverTab[77597]++
															if err := s.fc.onData(size); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1140
			_go_fuzz_dep_.CoverTab[77600]++
																t.closeStream(s, io.EOF, true, http2.ErrCodeFlowControl, status.New(codes.Internal, err.Error()), nil, false)
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1142
			// _ = "end of CoverTab[77600]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1143
			_go_fuzz_dep_.CoverTab[77601]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1143
			// _ = "end of CoverTab[77601]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1143
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1143
		// _ = "end of CoverTab[77597]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1143
		_go_fuzz_dep_.CoverTab[77598]++
															if f.Header().Flags.Has(http2.FlagDataPadded) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1144
			_go_fuzz_dep_.CoverTab[77602]++
																if w := s.fc.onRead(size - uint32(len(f.Data()))); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1145
				_go_fuzz_dep_.CoverTab[77603]++
																	t.controlBuf.put(&outgoingWindowUpdate{s.id, w})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1146
				// _ = "end of CoverTab[77603]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1147
				_go_fuzz_dep_.CoverTab[77604]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1147
				// _ = "end of CoverTab[77604]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1147
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1147
			// _ = "end of CoverTab[77602]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1148
			_go_fuzz_dep_.CoverTab[77605]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1148
			// _ = "end of CoverTab[77605]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1148
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1148
		// _ = "end of CoverTab[77598]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1148
		_go_fuzz_dep_.CoverTab[77599]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1152
		if len(f.Data()) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1152
			_go_fuzz_dep_.CoverTab[77606]++
																buffer := t.bufferPool.get()
																buffer.Reset()
																buffer.Write(f.Data())
																s.write(recvMsg{buffer: buffer})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1156
			// _ = "end of CoverTab[77606]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1157
			_go_fuzz_dep_.CoverTab[77607]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1157
			// _ = "end of CoverTab[77607]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1157
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1157
		// _ = "end of CoverTab[77599]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1158
		_go_fuzz_dep_.CoverTab[77608]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1158
		// _ = "end of CoverTab[77608]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1158
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1158
	// _ = "end of CoverTab[77584]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1158
	_go_fuzz_dep_.CoverTab[77585]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1161
	if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1161
		_go_fuzz_dep_.CoverTab[77609]++
															t.closeStream(s, io.EOF, false, http2.ErrCodeNo, status.New(codes.Internal, "server closed the stream without sending trailers"), nil, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1162
		// _ = "end of CoverTab[77609]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1163
		_go_fuzz_dep_.CoverTab[77610]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1163
		// _ = "end of CoverTab[77610]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1163
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1163
	// _ = "end of CoverTab[77585]"
}

func (t *http2Client) handleRSTStream(f *http2.RSTStreamFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1166
	_go_fuzz_dep_.CoverTab[77611]++
														s := t.getStream(f)
														if s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1168
		_go_fuzz_dep_.CoverTab[77616]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1169
		// _ = "end of CoverTab[77616]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1170
		_go_fuzz_dep_.CoverTab[77617]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1170
		// _ = "end of CoverTab[77617]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1170
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1170
	// _ = "end of CoverTab[77611]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1170
	_go_fuzz_dep_.CoverTab[77612]++
														if f.ErrCode == http2.ErrCodeRefusedStream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1171
		_go_fuzz_dep_.CoverTab[77618]++

															atomic.StoreUint32(&s.unprocessed, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1173
		// _ = "end of CoverTab[77618]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1174
		_go_fuzz_dep_.CoverTab[77619]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1174
		// _ = "end of CoverTab[77619]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1174
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1174
	// _ = "end of CoverTab[77612]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1174
	_go_fuzz_dep_.CoverTab[77613]++
														statusCode, ok := http2ErrConvTab[f.ErrCode]
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1176
		_go_fuzz_dep_.CoverTab[77620]++
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1177
			_go_fuzz_dep_.CoverTab[77622]++
																logger.Warningf("transport: http2Client.handleRSTStream found no mapped gRPC status for the received http2 error: %v", f.ErrCode)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1178
			// _ = "end of CoverTab[77622]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1179
			_go_fuzz_dep_.CoverTab[77623]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1179
			// _ = "end of CoverTab[77623]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1179
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1179
		// _ = "end of CoverTab[77620]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1179
		_go_fuzz_dep_.CoverTab[77621]++
															statusCode = codes.Unknown
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1180
		// _ = "end of CoverTab[77621]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1181
		_go_fuzz_dep_.CoverTab[77624]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1181
		// _ = "end of CoverTab[77624]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1181
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1181
	// _ = "end of CoverTab[77613]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1181
	_go_fuzz_dep_.CoverTab[77614]++
														if statusCode == codes.Canceled {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1182
		_go_fuzz_dep_.CoverTab[77625]++
															if d, ok := s.ctx.Deadline(); ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1183
			_go_fuzz_dep_.CoverTab[77626]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1183
			return !d.After(time.Now())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1183
			// _ = "end of CoverTab[77626]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1183
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1183
			_go_fuzz_dep_.CoverTab[77627]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1186
			statusCode = codes.DeadlineExceeded
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1186
			// _ = "end of CoverTab[77627]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1187
			_go_fuzz_dep_.CoverTab[77628]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1187
			// _ = "end of CoverTab[77628]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1187
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1187
		// _ = "end of CoverTab[77625]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1188
		_go_fuzz_dep_.CoverTab[77629]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1188
		// _ = "end of CoverTab[77629]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1188
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1188
	// _ = "end of CoverTab[77614]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1188
	_go_fuzz_dep_.CoverTab[77615]++
														t.closeStream(s, io.EOF, false, http2.ErrCodeNo, status.Newf(statusCode, "stream terminated by RST_STREAM with error code: %v", f.ErrCode), nil, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1189
	// _ = "end of CoverTab[77615]"
}

func (t *http2Client) handleSettings(f *http2.SettingsFrame, isFirst bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1192
	_go_fuzz_dep_.CoverTab[77630]++
														if f.IsAck() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1193
		_go_fuzz_dep_.CoverTab[77635]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1194
		// _ = "end of CoverTab[77635]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1195
		_go_fuzz_dep_.CoverTab[77636]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1195
		// _ = "end of CoverTab[77636]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1195
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1195
	// _ = "end of CoverTab[77630]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1195
	_go_fuzz_dep_.CoverTab[77631]++
														var maxStreams *uint32
														var ss []http2.Setting
														var updateFuncs []func()
														f.ForeachSetting(func(s http2.Setting) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1199
		_go_fuzz_dep_.CoverTab[77637]++
															switch s.ID {
		case http2.SettingMaxConcurrentStreams:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1201
			_go_fuzz_dep_.CoverTab[77639]++
																maxStreams = new(uint32)
																*maxStreams = s.Val
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1203
			// _ = "end of CoverTab[77639]"
		case http2.SettingMaxHeaderListSize:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1204
			_go_fuzz_dep_.CoverTab[77640]++
																updateFuncs = append(updateFuncs, func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1205
				_go_fuzz_dep_.CoverTab[77642]++
																	t.maxSendHeaderListSize = new(uint32)
																	*t.maxSendHeaderListSize = s.Val
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1207
				// _ = "end of CoverTab[77642]"
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1208
			// _ = "end of CoverTab[77640]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1209
			_go_fuzz_dep_.CoverTab[77641]++
																ss = append(ss, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1210
			// _ = "end of CoverTab[77641]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1211
		// _ = "end of CoverTab[77637]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1211
		_go_fuzz_dep_.CoverTab[77638]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1212
		// _ = "end of CoverTab[77638]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1213
	// _ = "end of CoverTab[77631]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1213
	_go_fuzz_dep_.CoverTab[77632]++
														if isFirst && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1214
		_go_fuzz_dep_.CoverTab[77643]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1214
		return maxStreams == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1214
		// _ = "end of CoverTab[77643]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1214
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1214
		_go_fuzz_dep_.CoverTab[77644]++
															maxStreams = new(uint32)
															*maxStreams = math.MaxUint32
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1216
		// _ = "end of CoverTab[77644]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1217
		_go_fuzz_dep_.CoverTab[77645]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1217
		// _ = "end of CoverTab[77645]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1217
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1217
	// _ = "end of CoverTab[77632]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1217
	_go_fuzz_dep_.CoverTab[77633]++
														sf := &incomingSettings{
		ss: ss,
	}
	if maxStreams != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1221
		_go_fuzz_dep_.CoverTab[77646]++
															updateStreamQuota := func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1222
			_go_fuzz_dep_.CoverTab[77648]++
																delta := int64(*maxStreams) - int64(t.maxConcurrentStreams)
																t.maxConcurrentStreams = *maxStreams
																t.streamQuota += delta
																if delta > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1226
				_go_fuzz_dep_.CoverTab[77649]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1226
				return t.waitingStreams > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1226
				// _ = "end of CoverTab[77649]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1226
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1226
				_go_fuzz_dep_.CoverTab[77650]++
																	close(t.streamsQuotaAvailable)
																	t.streamsQuotaAvailable = make(chan struct{}, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1228
				// _ = "end of CoverTab[77650]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1229
				_go_fuzz_dep_.CoverTab[77651]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1229
				// _ = "end of CoverTab[77651]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1229
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1229
			// _ = "end of CoverTab[77648]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1230
		// _ = "end of CoverTab[77646]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1230
		_go_fuzz_dep_.CoverTab[77647]++
															updateFuncs = append(updateFuncs, updateStreamQuota)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1231
		// _ = "end of CoverTab[77647]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1232
		_go_fuzz_dep_.CoverTab[77652]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1232
		// _ = "end of CoverTab[77652]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1232
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1232
	// _ = "end of CoverTab[77633]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1232
	_go_fuzz_dep_.CoverTab[77634]++
														t.controlBuf.executeAndPut(func(interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1233
		_go_fuzz_dep_.CoverTab[77653]++
															for _, f := range updateFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1234
			_go_fuzz_dep_.CoverTab[77655]++
																f()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1235
			// _ = "end of CoverTab[77655]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1236
		// _ = "end of CoverTab[77653]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1236
		_go_fuzz_dep_.CoverTab[77654]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1237
		// _ = "end of CoverTab[77654]"
	}, sf)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1238
	// _ = "end of CoverTab[77634]"
}

func (t *http2Client) handlePing(f *http2.PingFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1241
	_go_fuzz_dep_.CoverTab[77656]++
														if f.IsAck() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1242
		_go_fuzz_dep_.CoverTab[77658]++

															if t.bdpEst != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1244
			_go_fuzz_dep_.CoverTab[77660]++
																t.bdpEst.calculate(f.Data)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1245
			// _ = "end of CoverTab[77660]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1246
			_go_fuzz_dep_.CoverTab[77661]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1246
			// _ = "end of CoverTab[77661]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1246
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1246
		// _ = "end of CoverTab[77658]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1246
		_go_fuzz_dep_.CoverTab[77659]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1247
		// _ = "end of CoverTab[77659]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1248
		_go_fuzz_dep_.CoverTab[77662]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1248
		// _ = "end of CoverTab[77662]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1248
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1248
	// _ = "end of CoverTab[77656]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1248
	_go_fuzz_dep_.CoverTab[77657]++
														pingAck := &ping{ack: true}
														copy(pingAck.data[:], f.Data[:])
														t.controlBuf.put(pingAck)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1251
	// _ = "end of CoverTab[77657]"
}

func (t *http2Client) handleGoAway(f *http2.GoAwayFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1254
	_go_fuzz_dep_.CoverTab[77663]++
														t.mu.Lock()
														if t.state == closing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1256
		_go_fuzz_dep_.CoverTab[77671]++
															t.mu.Unlock()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1258
		// _ = "end of CoverTab[77671]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1259
		_go_fuzz_dep_.CoverTab[77672]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1259
		// _ = "end of CoverTab[77672]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1259
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1259
	// _ = "end of CoverTab[77663]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1259
	_go_fuzz_dep_.CoverTab[77664]++
														if f.ErrCode == http2.ErrCodeEnhanceYourCalm && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1260
		_go_fuzz_dep_.CoverTab[77673]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1260
		return string(f.DebugData()) == "too_many_pings"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1260
		// _ = "end of CoverTab[77673]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1260
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1260
		_go_fuzz_dep_.CoverTab[77674]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1265
		logger.Errorf("Client received GoAway with error code ENHANCE_YOUR_CALM and debug data equal to ASCII \"too_many_pings\".")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1265
		// _ = "end of CoverTab[77674]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1266
		_go_fuzz_dep_.CoverTab[77675]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1266
		// _ = "end of CoverTab[77675]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1266
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1266
	// _ = "end of CoverTab[77664]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1266
	_go_fuzz_dep_.CoverTab[77665]++
														id := f.LastStreamID
														if id > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1268
		_go_fuzz_dep_.CoverTab[77676]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1268
		return id%2 == 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1268
		// _ = "end of CoverTab[77676]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1268
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1268
		_go_fuzz_dep_.CoverTab[77677]++
															t.mu.Unlock()
															t.Close(connectionErrorf(true, nil, "received goaway with non-zero even-numbered numbered stream id: %v", id))
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1271
		// _ = "end of CoverTab[77677]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1272
		_go_fuzz_dep_.CoverTab[77678]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1272
		// _ = "end of CoverTab[77678]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1272
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1272
	// _ = "end of CoverTab[77665]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1272
	_go_fuzz_dep_.CoverTab[77666]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1283
	select {
	case <-t.goAway:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1284
		_go_fuzz_dep_.CoverTab[77679]++

															if id > t.prevGoAwayID {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1286
			_go_fuzz_dep_.CoverTab[77681]++
																t.mu.Unlock()
																t.Close(connectionErrorf(true, nil, "received goaway with stream id: %v, which exceeds stream id of previous goaway: %v", id, t.prevGoAwayID))
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1289
			// _ = "end of CoverTab[77681]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1290
			_go_fuzz_dep_.CoverTab[77682]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1290
			// _ = "end of CoverTab[77682]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1290
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1290
		// _ = "end of CoverTab[77679]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1291
		_go_fuzz_dep_.CoverTab[77680]++
															t.setGoAwayReason(f)
															close(t.goAway)
															defer t.controlBuf.put(&incomingGoAway{})

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1298
		if t.state != draining {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1298
			_go_fuzz_dep_.CoverTab[77683]++
																t.onClose(t.goAwayReason)
																t.state = draining
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1300
			// _ = "end of CoverTab[77683]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1301
			_go_fuzz_dep_.CoverTab[77684]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1301
			// _ = "end of CoverTab[77684]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1301
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1301
		// _ = "end of CoverTab[77680]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1302
	// _ = "end of CoverTab[77666]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1302
	_go_fuzz_dep_.CoverTab[77667]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1305
	upperLimit := t.prevGoAwayID
	if upperLimit == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1306
		_go_fuzz_dep_.CoverTab[77685]++
															upperLimit = math.MaxUint32
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1307
		// _ = "end of CoverTab[77685]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1308
		_go_fuzz_dep_.CoverTab[77686]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1308
		// _ = "end of CoverTab[77686]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1308
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1308
	// _ = "end of CoverTab[77667]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1308
	_go_fuzz_dep_.CoverTab[77668]++

														t.prevGoAwayID = id
														if len(t.activeStreams) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1311
		_go_fuzz_dep_.CoverTab[77687]++
															t.mu.Unlock()
															t.Close(connectionErrorf(true, nil, "received goaway and there are no active streams"))
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1314
		// _ = "end of CoverTab[77687]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1315
		_go_fuzz_dep_.CoverTab[77688]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1315
		// _ = "end of CoverTab[77688]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1315
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1315
	// _ = "end of CoverTab[77668]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1315
	_go_fuzz_dep_.CoverTab[77669]++

														streamsToClose := make([]*Stream, 0)
														for streamID, stream := range t.activeStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1318
		_go_fuzz_dep_.CoverTab[77689]++
															if streamID > id && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1319
			_go_fuzz_dep_.CoverTab[77690]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1319
			return streamID <= upperLimit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1319
			// _ = "end of CoverTab[77690]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1319
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1319
			_go_fuzz_dep_.CoverTab[77691]++

																if streamID > id && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1321
				_go_fuzz_dep_.CoverTab[77692]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1321
				return streamID <= upperLimit
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1321
				// _ = "end of CoverTab[77692]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1321
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1321
				_go_fuzz_dep_.CoverTab[77693]++
																	atomic.StoreUint32(&stream.unprocessed, 1)
																	streamsToClose = append(streamsToClose, stream)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1323
				// _ = "end of CoverTab[77693]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1324
				_go_fuzz_dep_.CoverTab[77694]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1324
				// _ = "end of CoverTab[77694]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1324
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1324
			// _ = "end of CoverTab[77691]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1325
			_go_fuzz_dep_.CoverTab[77695]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1325
			// _ = "end of CoverTab[77695]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1325
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1325
		// _ = "end of CoverTab[77689]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1326
	// _ = "end of CoverTab[77669]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1326
	_go_fuzz_dep_.CoverTab[77670]++
														t.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1330
	for _, stream := range streamsToClose {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1330
		_go_fuzz_dep_.CoverTab[77696]++
															t.closeStream(stream, errStreamDrain, false, http2.ErrCodeNo, statusGoAway, nil, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1331
		// _ = "end of CoverTab[77696]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1332
	// _ = "end of CoverTab[77670]"
}

// setGoAwayReason sets the value of t.goAwayReason based
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1335
// on the GoAway frame received.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1335
// It expects a lock on transport's mutext to be held by
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1335
// the caller.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1339
func (t *http2Client) setGoAwayReason(f *http2.GoAwayFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1339
	_go_fuzz_dep_.CoverTab[77697]++
														t.goAwayReason = GoAwayNoReason
														switch f.ErrCode {
	case http2.ErrCodeEnhanceYourCalm:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1342
		_go_fuzz_dep_.CoverTab[77699]++
															if string(f.DebugData()) == "too_many_pings" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1343
			_go_fuzz_dep_.CoverTab[77701]++
																t.goAwayReason = GoAwayTooManyPings
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1344
			// _ = "end of CoverTab[77701]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
			_go_fuzz_dep_.CoverTab[77702]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
			// _ = "end of CoverTab[77702]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
		// _ = "end of CoverTab[77699]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
		_go_fuzz_dep_.CoverTab[77700]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1345
		// _ = "end of CoverTab[77700]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1346
	// _ = "end of CoverTab[77697]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1346
	_go_fuzz_dep_.CoverTab[77698]++
														if len(f.DebugData()) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1347
		_go_fuzz_dep_.CoverTab[77703]++
															t.goAwayDebugMessage = fmt.Sprintf("code: %s", f.ErrCode)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1348
		// _ = "end of CoverTab[77703]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1349
		_go_fuzz_dep_.CoverTab[77704]++
															t.goAwayDebugMessage = fmt.Sprintf("code: %s, debug data: %q", f.ErrCode, string(f.DebugData()))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1350
		// _ = "end of CoverTab[77704]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1351
	// _ = "end of CoverTab[77698]"
}

func (t *http2Client) GetGoAwayReason() (GoAwayReason, string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1354
	_go_fuzz_dep_.CoverTab[77705]++
														t.mu.Lock()
														defer t.mu.Unlock()
														return t.goAwayReason, t.goAwayDebugMessage
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1357
	// _ = "end of CoverTab[77705]"
}

func (t *http2Client) handleWindowUpdate(f *http2.WindowUpdateFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1360
	_go_fuzz_dep_.CoverTab[77706]++
														t.controlBuf.put(&incomingWindowUpdate{
		streamID:	f.Header().StreamID,
		increment:	f.Increment,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1364
	// _ = "end of CoverTab[77706]"
}

// operateHeaders takes action on the decoded headers.
func (t *http2Client) operateHeaders(frame *http2.MetaHeadersFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1368
	_go_fuzz_dep_.CoverTab[77707]++
														s := t.getStream(frame)
														if s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1370
		_go_fuzz_dep_.CoverTab[77719]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1371
		// _ = "end of CoverTab[77719]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1372
		_go_fuzz_dep_.CoverTab[77720]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1372
		// _ = "end of CoverTab[77720]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1372
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1372
	// _ = "end of CoverTab[77707]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1372
	_go_fuzz_dep_.CoverTab[77708]++
														endStream := frame.StreamEnded()
														atomic.StoreUint32(&s.bytesReceived, 1)
														initialHeader := atomic.LoadUint32(&s.headerChanClosed) == 0

														if !initialHeader && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1377
		_go_fuzz_dep_.CoverTab[77721]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1377
		return !endStream
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1377
		// _ = "end of CoverTab[77721]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1377
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1377
		_go_fuzz_dep_.CoverTab[77722]++

															st := status.New(codes.Internal, "a HEADERS frame cannot appear in the middle of a stream")
															t.closeStream(s, st.Err(), true, http2.ErrCodeProtocol, st, nil, false)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1381
		// _ = "end of CoverTab[77722]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1382
		_go_fuzz_dep_.CoverTab[77723]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1382
		// _ = "end of CoverTab[77723]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1382
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1382
	// _ = "end of CoverTab[77708]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1382
	_go_fuzz_dep_.CoverTab[77709]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1386
	if frame.Truncated {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1386
		_go_fuzz_dep_.CoverTab[77724]++
															se := status.New(codes.Internal, "peer header list size exceeded limit")
															t.closeStream(s, se.Err(), true, http2.ErrCodeFrameSize, se, nil, endStream)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1389
		// _ = "end of CoverTab[77724]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1390
		_go_fuzz_dep_.CoverTab[77725]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1390
		// _ = "end of CoverTab[77725]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1390
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1390
	// _ = "end of CoverTab[77709]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1390
	_go_fuzz_dep_.CoverTab[77710]++

														var (
		// If a gRPC Response-Headers has already been received, then it means
		// that the peer is speaking gRPC and we are in gRPC mode.
		isGRPC		= !initialHeader
		mdata		= make(map[string][]string)
		contentTypeErr	= "malformed header: missing HTTP content-type"
		grpcMessage	string
		statusGen	*status.Status
		recvCompress	string
		httpStatusCode	*int
		httpStatusErr	string
		rawStatusCode	= codes.Unknown
		// headerError is set if an error is encountered while parsing the headers
		headerError	string
	)

	if initialHeader {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1408
		_go_fuzz_dep_.CoverTab[77726]++
															httpStatusErr = "malformed header: missing HTTP status"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1409
		// _ = "end of CoverTab[77726]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1410
		_go_fuzz_dep_.CoverTab[77727]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1410
		// _ = "end of CoverTab[77727]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1410
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1410
	// _ = "end of CoverTab[77710]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1410
	_go_fuzz_dep_.CoverTab[77711]++

														for _, hf := range frame.Fields {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1412
		_go_fuzz_dep_.CoverTab[77728]++
															switch hf.Name {
		case "content-type":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1414
			_go_fuzz_dep_.CoverTab[77729]++
																if _, validContentType := grpcutil.ContentSubtype(hf.Value); !validContentType {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1415
				_go_fuzz_dep_.CoverTab[77742]++
																	contentTypeErr = fmt.Sprintf("transport: received unexpected content-type %q", hf.Value)
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1417
				// _ = "end of CoverTab[77742]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1418
				_go_fuzz_dep_.CoverTab[77743]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1418
				// _ = "end of CoverTab[77743]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1418
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1418
			// _ = "end of CoverTab[77729]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1418
			_go_fuzz_dep_.CoverTab[77730]++
																contentTypeErr = ""
																mdata[hf.Name] = append(mdata[hf.Name], hf.Value)
																isGRPC = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1421
			// _ = "end of CoverTab[77730]"
		case "grpc-encoding":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1422
			_go_fuzz_dep_.CoverTab[77731]++
																recvCompress = hf.Value
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1423
			// _ = "end of CoverTab[77731]"
		case "grpc-status":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1424
			_go_fuzz_dep_.CoverTab[77732]++
																code, err := strconv.ParseInt(hf.Value, 10, 32)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1426
				_go_fuzz_dep_.CoverTab[77744]++
																	se := status.New(codes.Internal, fmt.Sprintf("transport: malformed grpc-status: %v", err))
																	t.closeStream(s, se.Err(), true, http2.ErrCodeProtocol, se, nil, endStream)
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1429
				// _ = "end of CoverTab[77744]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1430
				_go_fuzz_dep_.CoverTab[77745]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1430
				// _ = "end of CoverTab[77745]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1430
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1430
			// _ = "end of CoverTab[77732]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1430
			_go_fuzz_dep_.CoverTab[77733]++
																rawStatusCode = codes.Code(uint32(code))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1431
			// _ = "end of CoverTab[77733]"
		case "grpc-message":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1432
			_go_fuzz_dep_.CoverTab[77734]++
																grpcMessage = decodeGrpcMessage(hf.Value)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1433
			// _ = "end of CoverTab[77734]"
		case "grpc-status-details-bin":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1434
			_go_fuzz_dep_.CoverTab[77735]++
																var err error
																statusGen, err = decodeGRPCStatusDetails(hf.Value)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1437
				_go_fuzz_dep_.CoverTab[77746]++
																	headerError = fmt.Sprintf("transport: malformed grpc-status-details-bin: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1438
				// _ = "end of CoverTab[77746]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1439
				_go_fuzz_dep_.CoverTab[77747]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1439
				// _ = "end of CoverTab[77747]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1439
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1439
			// _ = "end of CoverTab[77735]"
		case ":status":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1440
			_go_fuzz_dep_.CoverTab[77736]++
																if hf.Value == "200" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1441
				_go_fuzz_dep_.CoverTab[77748]++
																	httpStatusErr = ""
																	statusCode := 200
																	httpStatusCode = &statusCode
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1445
				// _ = "end of CoverTab[77748]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1446
				_go_fuzz_dep_.CoverTab[77749]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1446
				// _ = "end of CoverTab[77749]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1446
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1446
			// _ = "end of CoverTab[77736]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1446
			_go_fuzz_dep_.CoverTab[77737]++

																c, err := strconv.ParseInt(hf.Value, 10, 32)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1449
				_go_fuzz_dep_.CoverTab[77750]++
																	se := status.New(codes.Internal, fmt.Sprintf("transport: malformed http-status: %v", err))
																	t.closeStream(s, se.Err(), true, http2.ErrCodeProtocol, se, nil, endStream)
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1452
				// _ = "end of CoverTab[77750]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1453
				_go_fuzz_dep_.CoverTab[77751]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1453
				// _ = "end of CoverTab[77751]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1453
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1453
			// _ = "end of CoverTab[77737]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1453
			_go_fuzz_dep_.CoverTab[77738]++
																statusCode := int(c)
																httpStatusCode = &statusCode

																httpStatusErr = fmt.Sprintf(
				"unexpected HTTP status code received from server: %d (%s)",
				statusCode,
				http.StatusText(statusCode),
			)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1461
			// _ = "end of CoverTab[77738]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1462
			_go_fuzz_dep_.CoverTab[77739]++
																if isReservedHeader(hf.Name) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1463
				_go_fuzz_dep_.CoverTab[77752]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1463
				return !isWhitelistedHeader(hf.Name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1463
				// _ = "end of CoverTab[77752]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1463
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1463
				_go_fuzz_dep_.CoverTab[77753]++
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1464
				// _ = "end of CoverTab[77753]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1465
				_go_fuzz_dep_.CoverTab[77754]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1465
				// _ = "end of CoverTab[77754]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1465
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1465
			// _ = "end of CoverTab[77739]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1465
			_go_fuzz_dep_.CoverTab[77740]++
																v, err := decodeMetadataHeader(hf.Name, hf.Value)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1467
				_go_fuzz_dep_.CoverTab[77755]++
																	headerError = fmt.Sprintf("transport: malformed %s: %v", hf.Name, err)
																	logger.Warningf("Failed to decode metadata header (%q, %q): %v", hf.Name, hf.Value, err)
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1470
				// _ = "end of CoverTab[77755]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1471
				_go_fuzz_dep_.CoverTab[77756]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1471
				// _ = "end of CoverTab[77756]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1471
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1471
			// _ = "end of CoverTab[77740]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1471
			_go_fuzz_dep_.CoverTab[77741]++
																mdata[hf.Name] = append(mdata[hf.Name], v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1472
			// _ = "end of CoverTab[77741]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1473
		// _ = "end of CoverTab[77728]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1474
	// _ = "end of CoverTab[77711]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1474
	_go_fuzz_dep_.CoverTab[77712]++

														if !isGRPC || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1476
		_go_fuzz_dep_.CoverTab[77757]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1476
		return httpStatusErr != ""
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1476
		// _ = "end of CoverTab[77757]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1476
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1476
		_go_fuzz_dep_.CoverTab[77758]++
															var code = codes.Internal	// when header does not include HTTP status, return INTERNAL

															if httpStatusCode != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1479
			_go_fuzz_dep_.CoverTab[77762]++
																var ok bool
																code, ok = HTTPStatusConvTab[*httpStatusCode]
																if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1482
				_go_fuzz_dep_.CoverTab[77763]++
																	code = codes.Unknown
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1483
				// _ = "end of CoverTab[77763]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1484
				_go_fuzz_dep_.CoverTab[77764]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1484
				// _ = "end of CoverTab[77764]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1484
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1484
			// _ = "end of CoverTab[77762]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1485
			_go_fuzz_dep_.CoverTab[77765]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1485
			// _ = "end of CoverTab[77765]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1485
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1485
		// _ = "end of CoverTab[77758]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1485
		_go_fuzz_dep_.CoverTab[77759]++
															var errs []string
															if httpStatusErr != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1487
			_go_fuzz_dep_.CoverTab[77766]++
																errs = append(errs, httpStatusErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1488
			// _ = "end of CoverTab[77766]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1489
			_go_fuzz_dep_.CoverTab[77767]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1489
			// _ = "end of CoverTab[77767]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1489
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1489
		// _ = "end of CoverTab[77759]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1489
		_go_fuzz_dep_.CoverTab[77760]++
															if contentTypeErr != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1490
			_go_fuzz_dep_.CoverTab[77768]++
																errs = append(errs, contentTypeErr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1491
			// _ = "end of CoverTab[77768]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1492
			_go_fuzz_dep_.CoverTab[77769]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1492
			// _ = "end of CoverTab[77769]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1492
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1492
		// _ = "end of CoverTab[77760]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1492
		_go_fuzz_dep_.CoverTab[77761]++

															se := status.New(code, strings.Join(errs, "; "))
															t.closeStream(s, se.Err(), true, http2.ErrCodeProtocol, se, nil, endStream)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1496
		// _ = "end of CoverTab[77761]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1497
		_go_fuzz_dep_.CoverTab[77770]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1497
		// _ = "end of CoverTab[77770]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1497
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1497
	// _ = "end of CoverTab[77712]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1497
	_go_fuzz_dep_.CoverTab[77713]++

														if headerError != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1499
		_go_fuzz_dep_.CoverTab[77771]++
															se := status.New(codes.Internal, headerError)
															t.closeStream(s, se.Err(), true, http2.ErrCodeProtocol, se, nil, endStream)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1502
		// _ = "end of CoverTab[77771]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1503
		_go_fuzz_dep_.CoverTab[77772]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1503
		// _ = "end of CoverTab[77772]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1503
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1503
	// _ = "end of CoverTab[77713]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1503
	_go_fuzz_dep_.CoverTab[77714]++

														isHeader := false

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1508
	if atomic.CompareAndSwapUint32(&s.headerChanClosed, 0, 1) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1508
		_go_fuzz_dep_.CoverTab[77773]++
															s.headerValid = true
															if !endStream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1510
			_go_fuzz_dep_.CoverTab[77775]++

																isHeader = true

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1516
			s.recvCompress = recvCompress
			if len(mdata) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1517
				_go_fuzz_dep_.CoverTab[77776]++
																	s.header = mdata
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1518
				// _ = "end of CoverTab[77776]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1519
				_go_fuzz_dep_.CoverTab[77777]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1519
				// _ = "end of CoverTab[77777]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1519
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1519
			// _ = "end of CoverTab[77775]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1520
			_go_fuzz_dep_.CoverTab[77778]++

																s.noHeaders = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1522
			// _ = "end of CoverTab[77778]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1523
		// _ = "end of CoverTab[77773]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1523
		_go_fuzz_dep_.CoverTab[77774]++
															close(s.headerChan)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1524
		// _ = "end of CoverTab[77774]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1525
		_go_fuzz_dep_.CoverTab[77779]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1525
		// _ = "end of CoverTab[77779]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1525
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1525
	// _ = "end of CoverTab[77714]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1525
	_go_fuzz_dep_.CoverTab[77715]++

														for _, sh := range t.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1527
		_go_fuzz_dep_.CoverTab[77780]++
															if isHeader {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1528
			_go_fuzz_dep_.CoverTab[77781]++
																inHeader := &stats.InHeader{
				Client:		true,
				WireLength:	int(frame.Header().Length),
				Header:		metadata.MD(mdata).Copy(),
				Compression:	s.recvCompress,
			}
																sh.HandleRPC(s.ctx, inHeader)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1535
			// _ = "end of CoverTab[77781]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1536
			_go_fuzz_dep_.CoverTab[77782]++
																inTrailer := &stats.InTrailer{
				Client:		true,
				WireLength:	int(frame.Header().Length),
				Trailer:	metadata.MD(mdata).Copy(),
			}
																sh.HandleRPC(s.ctx, inTrailer)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1542
			// _ = "end of CoverTab[77782]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1543
		// _ = "end of CoverTab[77780]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1544
	// _ = "end of CoverTab[77715]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1544
	_go_fuzz_dep_.CoverTab[77716]++

														if !endStream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1546
		_go_fuzz_dep_.CoverTab[77783]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1547
		// _ = "end of CoverTab[77783]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1548
		_go_fuzz_dep_.CoverTab[77784]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1548
		// _ = "end of CoverTab[77784]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1548
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1548
	// _ = "end of CoverTab[77716]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1548
	_go_fuzz_dep_.CoverTab[77717]++

														if statusGen == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1550
		_go_fuzz_dep_.CoverTab[77785]++
															statusGen = status.New(rawStatusCode, grpcMessage)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1551
		// _ = "end of CoverTab[77785]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1552
		_go_fuzz_dep_.CoverTab[77786]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1552
		// _ = "end of CoverTab[77786]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1552
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1552
	// _ = "end of CoverTab[77717]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1552
	_go_fuzz_dep_.CoverTab[77718]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1555
	rst := s.getState() == streamActive
														t.closeStream(s, io.EOF, rst, http2.ErrCodeNo, statusGen, mdata, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1556
	// _ = "end of CoverTab[77718]"
}

// readServerPreface reads and handles the initial settings frame from the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1559
// server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1561
func (t *http2Client) readServerPreface() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1561
	_go_fuzz_dep_.CoverTab[77787]++
														frame, err := t.framer.fr.ReadFrame()
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1563
		_go_fuzz_dep_.CoverTab[77790]++
															return connectionErrorf(true, err, "error reading server preface: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1564
		// _ = "end of CoverTab[77790]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1565
		_go_fuzz_dep_.CoverTab[77791]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1565
		// _ = "end of CoverTab[77791]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1565
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1565
	// _ = "end of CoverTab[77787]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1565
	_go_fuzz_dep_.CoverTab[77788]++
														sf, ok := frame.(*http2.SettingsFrame)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1567
		_go_fuzz_dep_.CoverTab[77792]++
															return connectionErrorf(true, nil, "initial http2 frame from server is not a settings frame: %T", frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1568
		// _ = "end of CoverTab[77792]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1569
		_go_fuzz_dep_.CoverTab[77793]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1569
		// _ = "end of CoverTab[77793]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1569
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1569
	// _ = "end of CoverTab[77788]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1569
	_go_fuzz_dep_.CoverTab[77789]++
														t.handleSettings(sf, true)
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1571
	// _ = "end of CoverTab[77789]"
}

// reader verifies the server preface and reads all subsequent data from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1574
// network connection.  If the server preface is not read successfully, an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1574
// error is pushed to errCh; otherwise errCh is closed with no error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1577
func (t *http2Client) reader(errCh chan<- error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1577
	_go_fuzz_dep_.CoverTab[77794]++
														defer close(t.readerDone)

														if err := t.readServerPreface(); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1580
		_go_fuzz_dep_.CoverTab[77797]++
															errCh <- err
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1582
		// _ = "end of CoverTab[77797]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1583
		_go_fuzz_dep_.CoverTab[77798]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1583
		// _ = "end of CoverTab[77798]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1583
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1583
	// _ = "end of CoverTab[77794]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1583
	_go_fuzz_dep_.CoverTab[77795]++
														close(errCh)
														if t.keepaliveEnabled {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1585
		_go_fuzz_dep_.CoverTab[77799]++
															atomic.StoreInt64(&t.lastRead, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1586
		// _ = "end of CoverTab[77799]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1587
		_go_fuzz_dep_.CoverTab[77800]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1587
		// _ = "end of CoverTab[77800]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1587
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1587
	// _ = "end of CoverTab[77795]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1587
	_go_fuzz_dep_.CoverTab[77796]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1590
	for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1590
		_go_fuzz_dep_.CoverTab[77801]++
															t.controlBuf.throttle()
															frame, err := t.framer.fr.ReadFrame()
															if t.keepaliveEnabled {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1593
			_go_fuzz_dep_.CoverTab[77804]++
																atomic.StoreInt64(&t.lastRead, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1594
			// _ = "end of CoverTab[77804]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1595
			_go_fuzz_dep_.CoverTab[77805]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1595
			// _ = "end of CoverTab[77805]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1595
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1595
		// _ = "end of CoverTab[77801]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1595
		_go_fuzz_dep_.CoverTab[77802]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1596
			_go_fuzz_dep_.CoverTab[77806]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1600
			if se, ok := err.(http2.StreamError); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1600
				_go_fuzz_dep_.CoverTab[77807]++
																	t.mu.Lock()
																	s := t.activeStreams[se.StreamID]
																	t.mu.Unlock()
																	if s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1604
					_go_fuzz_dep_.CoverTab[77809]++

																		code := http2ErrConvTab[se.Code]
																		errorDetail := t.framer.fr.ErrorDetail()
																		var msg string
																		if errorDetail != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1609
						_go_fuzz_dep_.CoverTab[77811]++
																			msg = errorDetail.Error()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1610
						// _ = "end of CoverTab[77811]"
					} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1611
						_go_fuzz_dep_.CoverTab[77812]++
																			msg = "received invalid frame"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1612
						// _ = "end of CoverTab[77812]"
					}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1613
					// _ = "end of CoverTab[77809]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1613
					_go_fuzz_dep_.CoverTab[77810]++
																		t.closeStream(s, status.Error(code, msg), true, http2.ErrCodeProtocol, status.New(code, msg), nil, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1614
					// _ = "end of CoverTab[77810]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1615
					_go_fuzz_dep_.CoverTab[77813]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1615
					// _ = "end of CoverTab[77813]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1615
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1615
				// _ = "end of CoverTab[77807]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1615
				_go_fuzz_dep_.CoverTab[77808]++
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1616
				// _ = "end of CoverTab[77808]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1617
				_go_fuzz_dep_.CoverTab[77814]++

																	t.Close(connectionErrorf(true, err, "error reading from server: %v", err))
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1620
				// _ = "end of CoverTab[77814]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1621
			// _ = "end of CoverTab[77806]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1622
			_go_fuzz_dep_.CoverTab[77815]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1622
			// _ = "end of CoverTab[77815]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1622
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1622
		// _ = "end of CoverTab[77802]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1622
		_go_fuzz_dep_.CoverTab[77803]++
															switch frame := frame.(type) {
		case *http2.MetaHeadersFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1624
			_go_fuzz_dep_.CoverTab[77816]++
																t.operateHeaders(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1625
			// _ = "end of CoverTab[77816]"
		case *http2.DataFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1626
			_go_fuzz_dep_.CoverTab[77817]++
																t.handleData(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1627
			// _ = "end of CoverTab[77817]"
		case *http2.RSTStreamFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1628
			_go_fuzz_dep_.CoverTab[77818]++
																t.handleRSTStream(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1629
			// _ = "end of CoverTab[77818]"
		case *http2.SettingsFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1630
			_go_fuzz_dep_.CoverTab[77819]++
																t.handleSettings(frame, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1631
			// _ = "end of CoverTab[77819]"
		case *http2.PingFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1632
			_go_fuzz_dep_.CoverTab[77820]++
																t.handlePing(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1633
			// _ = "end of CoverTab[77820]"
		case *http2.GoAwayFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1634
			_go_fuzz_dep_.CoverTab[77821]++
																t.handleGoAway(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1635
			// _ = "end of CoverTab[77821]"
		case *http2.WindowUpdateFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1636
			_go_fuzz_dep_.CoverTab[77822]++
																t.handleWindowUpdate(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1637
			// _ = "end of CoverTab[77822]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1638
			_go_fuzz_dep_.CoverTab[77823]++
																if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1639
				_go_fuzz_dep_.CoverTab[77824]++
																	logger.Errorf("transport: http2Client.reader got unhandled frame type %v.", frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1640
				// _ = "end of CoverTab[77824]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1641
				_go_fuzz_dep_.CoverTab[77825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1641
				// _ = "end of CoverTab[77825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1641
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1641
			// _ = "end of CoverTab[77823]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1642
		// _ = "end of CoverTab[77803]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1643
	// _ = "end of CoverTab[77796]"
}

func minTime(a, b time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1646
	_go_fuzz_dep_.CoverTab[77826]++
														if a < b {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1647
		_go_fuzz_dep_.CoverTab[77828]++
															return a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1648
		// _ = "end of CoverTab[77828]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1649
		_go_fuzz_dep_.CoverTab[77829]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1649
		// _ = "end of CoverTab[77829]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1649
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1649
	// _ = "end of CoverTab[77826]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1649
	_go_fuzz_dep_.CoverTab[77827]++
														return b
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1650
	// _ = "end of CoverTab[77827]"
}

// keepalive running in a separate goroutine makes sure the connection is alive by sending pings.
func (t *http2Client) keepalive() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1654
	_go_fuzz_dep_.CoverTab[77830]++
														p := &ping{data: [8]byte{}}

														outstandingPing := false

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1660
	timeoutLeft := time.Duration(0)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1663
	prevNano := time.Now().UnixNano()
	timer := time.NewTimer(t.kp.Time)
	for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1665
		_go_fuzz_dep_.CoverTab[77831]++
															select {
		case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1667
			_go_fuzz_dep_.CoverTab[77832]++
																lastRead := atomic.LoadInt64(&t.lastRead)
																if lastRead > prevNano {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1669
				_go_fuzz_dep_.CoverTab[77840]++

																	outstandingPing = false

																	timer.Reset(time.Duration(lastRead) + t.kp.Time - time.Duration(time.Now().UnixNano()))
																	prevNano = lastRead
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1675
				// _ = "end of CoverTab[77840]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1676
				_go_fuzz_dep_.CoverTab[77841]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1676
				// _ = "end of CoverTab[77841]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1676
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1676
			// _ = "end of CoverTab[77832]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1676
			_go_fuzz_dep_.CoverTab[77833]++
																if outstandingPing && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1677
				_go_fuzz_dep_.CoverTab[77842]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1677
				return timeoutLeft <= 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1677
				// _ = "end of CoverTab[77842]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1677
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1677
				_go_fuzz_dep_.CoverTab[77843]++
																	t.Close(connectionErrorf(true, nil, "keepalive ping failed to receive ACK within timeout"))
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1679
				// _ = "end of CoverTab[77843]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1680
				_go_fuzz_dep_.CoverTab[77844]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1680
				// _ = "end of CoverTab[77844]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1680
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1680
			// _ = "end of CoverTab[77833]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1680
			_go_fuzz_dep_.CoverTab[77834]++
																t.mu.Lock()
																if t.state == closing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1682
				_go_fuzz_dep_.CoverTab[77845]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1689
				t.mu.Unlock()
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1690
				// _ = "end of CoverTab[77845]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1691
				_go_fuzz_dep_.CoverTab[77846]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1691
				// _ = "end of CoverTab[77846]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1691
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1691
			// _ = "end of CoverTab[77834]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1691
			_go_fuzz_dep_.CoverTab[77835]++
																if len(t.activeStreams) < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1692
				_go_fuzz_dep_.CoverTab[77847]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1692
				return !t.kp.PermitWithoutStream
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1692
				// _ = "end of CoverTab[77847]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1692
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1692
				_go_fuzz_dep_.CoverTab[77848]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1698
				outstandingPing = false
																	t.kpDormant = true
																	t.kpDormancyCond.Wait()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1700
				// _ = "end of CoverTab[77848]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1701
				_go_fuzz_dep_.CoverTab[77849]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1701
				// _ = "end of CoverTab[77849]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1701
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1701
			// _ = "end of CoverTab[77835]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1701
			_go_fuzz_dep_.CoverTab[77836]++
																t.kpDormant = false
																t.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1708
			if !outstandingPing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1708
				_go_fuzz_dep_.CoverTab[77850]++
																	if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1709
					_go_fuzz_dep_.CoverTab[77852]++
																		atomic.AddInt64(&t.czData.kpCount, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1710
					// _ = "end of CoverTab[77852]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1711
					_go_fuzz_dep_.CoverTab[77853]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1711
					// _ = "end of CoverTab[77853]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1711
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1711
				// _ = "end of CoverTab[77850]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1711
				_go_fuzz_dep_.CoverTab[77851]++
																	t.controlBuf.put(p)
																	timeoutLeft = t.kp.Timeout
																	outstandingPing = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1714
				// _ = "end of CoverTab[77851]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1715
				_go_fuzz_dep_.CoverTab[77854]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1715
				// _ = "end of CoverTab[77854]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1715
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1715
			// _ = "end of CoverTab[77836]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1715
			_go_fuzz_dep_.CoverTab[77837]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1720
			sleepDuration := minTime(t.kp.Time, timeoutLeft)
																timeoutLeft -= sleepDuration
																timer.Reset(sleepDuration)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1722
			// _ = "end of CoverTab[77837]"
		case <-t.ctx.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1723
			_go_fuzz_dep_.CoverTab[77838]++
																if !timer.Stop() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1724
				_go_fuzz_dep_.CoverTab[77855]++
																	<-timer.C
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1725
				// _ = "end of CoverTab[77855]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1726
				_go_fuzz_dep_.CoverTab[77856]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1726
				// _ = "end of CoverTab[77856]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1726
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1726
			// _ = "end of CoverTab[77838]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1726
			_go_fuzz_dep_.CoverTab[77839]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1727
			// _ = "end of CoverTab[77839]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1728
		// _ = "end of CoverTab[77831]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1729
	// _ = "end of CoverTab[77830]"
}

func (t *http2Client) Error() <-chan struct{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1732
	_go_fuzz_dep_.CoverTab[77857]++
														return t.ctx.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1733
	// _ = "end of CoverTab[77857]"
}

func (t *http2Client) GoAway() <-chan struct{} {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1736
	_go_fuzz_dep_.CoverTab[77858]++
														return t.goAway
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1737
	// _ = "end of CoverTab[77858]"
}

func (t *http2Client) ChannelzMetric() *channelz.SocketInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1740
	_go_fuzz_dep_.CoverTab[77859]++
														s := channelz.SocketInternalMetric{
		StreamsStarted:				atomic.LoadInt64(&t.czData.streamsStarted),
		StreamsSucceeded:			atomic.LoadInt64(&t.czData.streamsSucceeded),
		StreamsFailed:				atomic.LoadInt64(&t.czData.streamsFailed),
		MessagesSent:				atomic.LoadInt64(&t.czData.msgSent),
		MessagesReceived:			atomic.LoadInt64(&t.czData.msgRecv),
		KeepAlivesSent:				atomic.LoadInt64(&t.czData.kpCount),
		LastLocalStreamCreatedTimestamp:	time.Unix(0, atomic.LoadInt64(&t.czData.lastStreamCreatedTime)),
		LastMessageSentTimestamp:		time.Unix(0, atomic.LoadInt64(&t.czData.lastMsgSentTime)),
		LastMessageReceivedTimestamp:		time.Unix(0, atomic.LoadInt64(&t.czData.lastMsgRecvTime)),
		LocalFlowControlWindow:			int64(t.fc.getSize()),
		SocketOptions:				channelz.GetSocketOption(t.conn),
		LocalAddr:				t.localAddr,
		RemoteAddr:				t.remoteAddr,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1756
	}
	if au, ok := t.authInfo.(credentials.ChannelzSecurityInfo); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1757
		_go_fuzz_dep_.CoverTab[77861]++
															s.Security = au.GetSecurityValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1758
		// _ = "end of CoverTab[77861]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1759
		_go_fuzz_dep_.CoverTab[77862]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1759
		// _ = "end of CoverTab[77862]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1759
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1759
	// _ = "end of CoverTab[77859]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1759
	_go_fuzz_dep_.CoverTab[77860]++
														s.RemoteFlowControlWindow = t.getOutFlowWindow()
														return &s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1761
	// _ = "end of CoverTab[77860]"
}

func (t *http2Client) RemoteAddr() net.Addr {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1764
	_go_fuzz_dep_.CoverTab[77863]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1764
	return t.remoteAddr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1764
	// _ = "end of CoverTab[77863]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1764
}

func (t *http2Client) IncrMsgSent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1766
	_go_fuzz_dep_.CoverTab[77864]++
														atomic.AddInt64(&t.czData.msgSent, 1)
														atomic.StoreInt64(&t.czData.lastMsgSentTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1768
	// _ = "end of CoverTab[77864]"
}

func (t *http2Client) IncrMsgRecv() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1771
	_go_fuzz_dep_.CoverTab[77865]++
														atomic.AddInt64(&t.czData.msgRecv, 1)
														atomic.StoreInt64(&t.czData.lastMsgRecvTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1773
	// _ = "end of CoverTab[77865]"
}

func (t *http2Client) getOutFlowWindow() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1776
	_go_fuzz_dep_.CoverTab[77866]++
														resp := make(chan uint32, 1)
														timer := time.NewTimer(time.Second)
														defer timer.Stop()
														t.controlBuf.put(&outFlowControlSizeRequest{resp})
														select {
	case sz := <-resp:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1782
		_go_fuzz_dep_.CoverTab[77867]++
															return int64(sz)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1783
		// _ = "end of CoverTab[77867]"
	case <-t.ctxDone:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1784
		_go_fuzz_dep_.CoverTab[77868]++
															return -1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1785
		// _ = "end of CoverTab[77868]"
	case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1786
		_go_fuzz_dep_.CoverTab[77869]++
															return -2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1787
		// _ = "end of CoverTab[77869]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1788
	// _ = "end of CoverTab[77866]"
}

func (t *http2Client) stateForTesting() transportState {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1791
	_go_fuzz_dep_.CoverTab[77870]++
														t.mu.Lock()
														defer t.mu.Unlock()
														return t.state
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1794
	// _ = "end of CoverTab[77870]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1795
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_client.go:1795
var _ = _go_fuzz_dep_.CoverTab
