//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
package transport

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:19
)

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
	"google.golang.org/grpc/internal/grpcutil"
	"google.golang.org/grpc/internal/syscall"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcrand"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

var (
	// ErrIllegalHeaderWrite indicates that setting header is illegal because of
	// the stream's state.
	ErrIllegalHeaderWrite	= status.Error(codes.Internal, "transport: SendHeader called multiple times")
	// ErrHeaderListSizeLimitViolation indicates that the header list size is larger
	// than the limit set by peer.
	ErrHeaderListSizeLimitViolation	= status.Error(codes.Internal, "transport: trying to send header list size larger than the limit set by peer")
)

// serverConnectionCounter counts the number of connections a server has seen
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:63
// (equal to the number of http2Servers created). Must be accessed atomically.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:65
var serverConnectionCounter uint64

// http2Server implements the ServerTransport interface with HTTP2.
type http2Server struct {
	lastRead	int64	// Keep this field 64-bit aligned. Accessed atomically.
	ctx		context.Context
	done		chan struct{}
	conn		net.Conn
	loopy		*loopyWriter
	readerDone	chan struct{}	// sync point to enable testing.
	writerDone	chan struct{}	// sync point to enable testing.
	remoteAddr	net.Addr
	localAddr	net.Addr
	authInfo	credentials.AuthInfo	// auth info about the connection
	inTapHandle	tap.ServerInHandle
	framer		*framer
	// The max number of concurrent streams.
	maxStreams	uint32
	// controlBuf delivers all the control related tasks (e.g., window
	// updates, reset streams, and various settings) to the controller.
	controlBuf	*controlBuffer
	fc		*trInFlow
	stats		[]stats.Handler
	// Keepalive and max-age parameters for the server.
	kp	keepalive.ServerParameters
	// Keepalive enforcement policy.
	kep	keepalive.EnforcementPolicy
	// The time instance last ping was received.
	lastPingAt	time.Time
	// Number of times the client has violated keepalive ping policy so far.
	pingStrikes	uint8
	// Flag to signify that number of ping strikes should be reset to 0.
	// This is set whenever data or header frames are sent.
	// 1 means yes.
	resetPingStrikes	uint32	// Accessed atomically.
	initialWindowSize	int32
	bdpEst			*bdpEstimator
	maxSendHeaderListSize	*uint32

	mu	sync.Mutex	// guard the following

	// drainEvent is initialized when Drain() is called the first time. After
	// which the server writes out the first GoAway(with ID 2^31-1) frame. Then
	// an independent goroutine will be launched to later send the second
	// GoAway. During this time we don't want to write another first GoAway(with
	// ID 2^31 -1) frame. Thus call to Drain() will be a no-op if drainEvent is
	// already initialized since draining is already underway.
	drainEvent	*grpcsync.Event
	state		transportState
	activeStreams	map[uint32]*Stream
	// idle is the time instant when the connection went idle.
	// This is either the beginning of the connection or when the number of
	// RPCs go down to 0.
	// When the connection is busy, this value is set to 0.
	idle	time.Time

	// Fields below are for channelz metric collection.
	channelzID	*channelz.Identifier
	czData		*channelzData
	bufferPool	*bufferPool

	connectionID	uint64

	// maxStreamMu guards the maximum stream ID
	// This lock may not be taken if mu is already held.
	maxStreamMu	sync.Mutex
	maxStreamID	uint32	// max stream ID ever seen
}

// NewServerTransport creates a http2 transport with conn and configuration
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:134
// options from config.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:134
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:134
// It returns a non-nil transport and a nil error on success. On failure, it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:134
// returns a nil transport and a non-nil error. For a special case where the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:134
// underlying conn gets closed before the client preface could be read, it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:134
// returns a nil transport and a nil error.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:141
func NewServerTransport(conn net.Conn, config *ServerConfig) (_ ServerTransport, err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:141
	_go_fuzz_dep_.CoverTab[77871]++
														var authInfo credentials.AuthInfo
														rawConn := conn
														if config.Credentials != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:144
		_go_fuzz_dep_.CoverTab[77899]++
															var err error
															conn, authInfo, err = config.Credentials.ServerHandshake(rawConn)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:147
			_go_fuzz_dep_.CoverTab[77900]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:152
			if err == credentials.ErrConnDispatched || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:152
				_go_fuzz_dep_.CoverTab[77902]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:152
				return err == io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:152
				// _ = "end of CoverTab[77902]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:152
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:152
				_go_fuzz_dep_.CoverTab[77903]++
																	return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:153
				// _ = "end of CoverTab[77903]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:154
				_go_fuzz_dep_.CoverTab[77904]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:154
				// _ = "end of CoverTab[77904]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:154
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:154
			// _ = "end of CoverTab[77900]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:154
			_go_fuzz_dep_.CoverTab[77901]++
																return nil, connectionErrorf(false, err, "ServerHandshake(%q) failed: %v", rawConn.RemoteAddr(), err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:155
			// _ = "end of CoverTab[77901]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:156
			_go_fuzz_dep_.CoverTab[77905]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:156
			// _ = "end of CoverTab[77905]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:156
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:156
		// _ = "end of CoverTab[77899]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:157
		_go_fuzz_dep_.CoverTab[77906]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:157
		// _ = "end of CoverTab[77906]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:157
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:157
	// _ = "end of CoverTab[77871]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:157
	_go_fuzz_dep_.CoverTab[77872]++
														writeBufSize := config.WriteBufferSize
														readBufSize := config.ReadBufferSize
														maxHeaderListSize := defaultServerMaxHeaderListSize
														if config.MaxHeaderListSize != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:161
		_go_fuzz_dep_.CoverTab[77907]++
															maxHeaderListSize = *config.MaxHeaderListSize
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:162
		// _ = "end of CoverTab[77907]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:163
		_go_fuzz_dep_.CoverTab[77908]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:163
		// _ = "end of CoverTab[77908]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:163
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:163
	// _ = "end of CoverTab[77872]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:163
	_go_fuzz_dep_.CoverTab[77873]++
														framer := newFramer(conn, writeBufSize, readBufSize, maxHeaderListSize)

														isettings := []http2.Setting{{
		ID:	http2.SettingMaxFrameSize,
		Val:	http2MaxFrameLen,
	}}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:172
	maxStreams := config.MaxStreams
	if maxStreams == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:173
		_go_fuzz_dep_.CoverTab[77909]++
															maxStreams = math.MaxUint32
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:174
		// _ = "end of CoverTab[77909]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:175
		_go_fuzz_dep_.CoverTab[77910]++
															isettings = append(isettings, http2.Setting{
			ID:	http2.SettingMaxConcurrentStreams,
			Val:	maxStreams,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:179
		// _ = "end of CoverTab[77910]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:180
	// _ = "end of CoverTab[77873]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:180
	_go_fuzz_dep_.CoverTab[77874]++
														dynamicWindow := true
														iwz := int32(initialWindowSize)
														if config.InitialWindowSize >= defaultWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:183
		_go_fuzz_dep_.CoverTab[77911]++
															iwz = config.InitialWindowSize
															dynamicWindow = false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:185
		// _ = "end of CoverTab[77911]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:186
		_go_fuzz_dep_.CoverTab[77912]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:186
		// _ = "end of CoverTab[77912]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:186
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:186
	// _ = "end of CoverTab[77874]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:186
	_go_fuzz_dep_.CoverTab[77875]++
														icwz := int32(initialWindowSize)
														if config.InitialConnWindowSize >= defaultWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:188
		_go_fuzz_dep_.CoverTab[77913]++
															icwz = config.InitialConnWindowSize
															dynamicWindow = false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:190
		// _ = "end of CoverTab[77913]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:191
		_go_fuzz_dep_.CoverTab[77914]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:191
		// _ = "end of CoverTab[77914]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:191
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:191
	// _ = "end of CoverTab[77875]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:191
	_go_fuzz_dep_.CoverTab[77876]++
														if iwz != defaultWindowSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:192
		_go_fuzz_dep_.CoverTab[77915]++
															isettings = append(isettings, http2.Setting{
			ID:	http2.SettingInitialWindowSize,
			Val:	uint32(iwz)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:195
		// _ = "end of CoverTab[77915]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:196
		_go_fuzz_dep_.CoverTab[77916]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:196
		// _ = "end of CoverTab[77916]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:196
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:196
	// _ = "end of CoverTab[77876]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:196
	_go_fuzz_dep_.CoverTab[77877]++
														if config.MaxHeaderListSize != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:197
		_go_fuzz_dep_.CoverTab[77917]++
															isettings = append(isettings, http2.Setting{
			ID:	http2.SettingMaxHeaderListSize,
			Val:	*config.MaxHeaderListSize,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:201
		// _ = "end of CoverTab[77917]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:202
		_go_fuzz_dep_.CoverTab[77918]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:202
		// _ = "end of CoverTab[77918]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:202
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:202
	// _ = "end of CoverTab[77877]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:202
	_go_fuzz_dep_.CoverTab[77878]++
														if config.HeaderTableSize != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:203
		_go_fuzz_dep_.CoverTab[77919]++
															isettings = append(isettings, http2.Setting{
			ID:	http2.SettingHeaderTableSize,
			Val:	*config.HeaderTableSize,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:207
		// _ = "end of CoverTab[77919]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:208
		_go_fuzz_dep_.CoverTab[77920]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:208
		// _ = "end of CoverTab[77920]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:208
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:208
	// _ = "end of CoverTab[77878]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:208
	_go_fuzz_dep_.CoverTab[77879]++
														if err := framer.fr.WriteSettings(isettings...); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:209
		_go_fuzz_dep_.CoverTab[77921]++
															return nil, connectionErrorf(false, err, "transport: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:210
		// _ = "end of CoverTab[77921]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:211
		_go_fuzz_dep_.CoverTab[77922]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:211
		// _ = "end of CoverTab[77922]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:211
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:211
	// _ = "end of CoverTab[77879]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:211
	_go_fuzz_dep_.CoverTab[77880]++

														if delta := uint32(icwz - defaultWindowSize); delta > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:213
		_go_fuzz_dep_.CoverTab[77923]++
															if err := framer.fr.WriteWindowUpdate(0, delta); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:214
			_go_fuzz_dep_.CoverTab[77924]++
																return nil, connectionErrorf(false, err, "transport: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:215
			// _ = "end of CoverTab[77924]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:216
			_go_fuzz_dep_.CoverTab[77925]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:216
			// _ = "end of CoverTab[77925]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:216
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:216
		// _ = "end of CoverTab[77923]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:217
		_go_fuzz_dep_.CoverTab[77926]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:217
		// _ = "end of CoverTab[77926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:217
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:217
	// _ = "end of CoverTab[77880]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:217
	_go_fuzz_dep_.CoverTab[77881]++
														kp := config.KeepaliveParams
														if kp.MaxConnectionIdle == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:219
		_go_fuzz_dep_.CoverTab[77927]++
															kp.MaxConnectionIdle = defaultMaxConnectionIdle
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:220
		// _ = "end of CoverTab[77927]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:221
		_go_fuzz_dep_.CoverTab[77928]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:221
		// _ = "end of CoverTab[77928]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:221
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:221
	// _ = "end of CoverTab[77881]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:221
	_go_fuzz_dep_.CoverTab[77882]++
														if kp.MaxConnectionAge == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:222
		_go_fuzz_dep_.CoverTab[77929]++
															kp.MaxConnectionAge = defaultMaxConnectionAge
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:223
		// _ = "end of CoverTab[77929]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:224
		_go_fuzz_dep_.CoverTab[77930]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:224
		// _ = "end of CoverTab[77930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:224
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:224
	// _ = "end of CoverTab[77882]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:224
	_go_fuzz_dep_.CoverTab[77883]++

														kp.MaxConnectionAge += getJitter(kp.MaxConnectionAge)
														if kp.MaxConnectionAgeGrace == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:227
		_go_fuzz_dep_.CoverTab[77931]++
															kp.MaxConnectionAgeGrace = defaultMaxConnectionAgeGrace
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:228
		// _ = "end of CoverTab[77931]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:229
		_go_fuzz_dep_.CoverTab[77932]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:229
		// _ = "end of CoverTab[77932]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:229
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:229
	// _ = "end of CoverTab[77883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:229
	_go_fuzz_dep_.CoverTab[77884]++
														if kp.Time == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:230
		_go_fuzz_dep_.CoverTab[77933]++
															kp.Time = defaultServerKeepaliveTime
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:231
		// _ = "end of CoverTab[77933]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:232
		_go_fuzz_dep_.CoverTab[77934]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:232
		// _ = "end of CoverTab[77934]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:232
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:232
	// _ = "end of CoverTab[77884]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:232
	_go_fuzz_dep_.CoverTab[77885]++
														if kp.Timeout == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:233
		_go_fuzz_dep_.CoverTab[77935]++
															kp.Timeout = defaultServerKeepaliveTimeout
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:234
		// _ = "end of CoverTab[77935]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:235
		_go_fuzz_dep_.CoverTab[77936]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:235
		// _ = "end of CoverTab[77936]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:235
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:235
	// _ = "end of CoverTab[77885]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:235
	_go_fuzz_dep_.CoverTab[77886]++
														if kp.Time != infinity {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:236
		_go_fuzz_dep_.CoverTab[77937]++
															if err = syscall.SetTCPUserTimeout(conn, kp.Timeout); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:237
			_go_fuzz_dep_.CoverTab[77938]++
																return nil, connectionErrorf(false, err, "transport: failed to set TCP_USER_TIMEOUT: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:238
			// _ = "end of CoverTab[77938]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:239
			_go_fuzz_dep_.CoverTab[77939]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:239
			// _ = "end of CoverTab[77939]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:239
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:239
		// _ = "end of CoverTab[77937]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:240
		_go_fuzz_dep_.CoverTab[77940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:240
		// _ = "end of CoverTab[77940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:240
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:240
	// _ = "end of CoverTab[77886]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:240
	_go_fuzz_dep_.CoverTab[77887]++
														kep := config.KeepalivePolicy
														if kep.MinTime == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:242
		_go_fuzz_dep_.CoverTab[77941]++
															kep.MinTime = defaultKeepalivePolicyMinTime
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:243
		// _ = "end of CoverTab[77941]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:244
		_go_fuzz_dep_.CoverTab[77942]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:244
		// _ = "end of CoverTab[77942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:244
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:244
	// _ = "end of CoverTab[77887]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:244
	_go_fuzz_dep_.CoverTab[77888]++

														done := make(chan struct{})
														t := &http2Server{
		ctx:			setConnection(context.Background(), rawConn),
		done:			done,
		conn:			conn,
		remoteAddr:		conn.RemoteAddr(),
		localAddr:		conn.LocalAddr(),
		authInfo:		authInfo,
		framer:			framer,
		readerDone:		make(chan struct{}),
		writerDone:		make(chan struct{}),
		maxStreams:		maxStreams,
		inTapHandle:		config.InTapHandle,
		fc:			&trInFlow{limit: uint32(icwz)},
		state:			reachable,
		activeStreams:		make(map[uint32]*Stream),
		stats:			config.StatsHandlers,
		kp:			kp,
		idle:			time.Now(),
		kep:			kep,
		initialWindowSize:	iwz,
		czData:			new(channelzData),
		bufferPool:		newBufferPool(),
	}

	t.ctx = peer.NewContext(t.ctx, t.getPeer())

	t.controlBuf = newControlBuffer(t.done)
	if dynamicWindow {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:274
		_go_fuzz_dep_.CoverTab[77943]++
															t.bdpEst = &bdpEstimator{
			bdp:			initialWindowSize,
			updateFlowControl:	t.updateFlowControl,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:278
		// _ = "end of CoverTab[77943]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:279
		_go_fuzz_dep_.CoverTab[77944]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:279
		// _ = "end of CoverTab[77944]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:279
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:279
	// _ = "end of CoverTab[77888]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:279
	_go_fuzz_dep_.CoverTab[77889]++
														for _, sh := range t.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:280
		_go_fuzz_dep_.CoverTab[77945]++
															t.ctx = sh.TagConn(t.ctx, &stats.ConnTagInfo{
			RemoteAddr:	t.remoteAddr,
			LocalAddr:	t.localAddr,
		})
															connBegin := &stats.ConnBegin{}
															sh.HandleConn(t.ctx, connBegin)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:286
		// _ = "end of CoverTab[77945]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:287
	// _ = "end of CoverTab[77889]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:287
	_go_fuzz_dep_.CoverTab[77890]++
														t.channelzID, err = channelz.RegisterNormalSocket(t, config.ChannelzParentID, fmt.Sprintf("%s -> %s", t.remoteAddr, t.localAddr))
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:289
		_go_fuzz_dep_.CoverTab[77946]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:290
		// _ = "end of CoverTab[77946]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:291
		_go_fuzz_dep_.CoverTab[77947]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:291
		// _ = "end of CoverTab[77947]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:291
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:291
	// _ = "end of CoverTab[77890]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:291
	_go_fuzz_dep_.CoverTab[77891]++

														t.connectionID = atomic.AddUint64(&serverConnectionCounter, 1)
														t.framer.writer.Flush()

														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:296
		_go_fuzz_dep_.CoverTab[77948]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:297
			_go_fuzz_dep_.CoverTab[77949]++
																t.Close(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:298
			// _ = "end of CoverTab[77949]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:299
			_go_fuzz_dep_.CoverTab[77950]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:299
			// _ = "end of CoverTab[77950]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:299
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:299
		// _ = "end of CoverTab[77948]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:300
	// _ = "end of CoverTab[77891]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:300
	_go_fuzz_dep_.CoverTab[77892]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:303
	preface := make([]byte, len(clientPreface))
	if _, err := io.ReadFull(t.conn, preface); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:304
		_go_fuzz_dep_.CoverTab[77951]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:310
		if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:310
			_go_fuzz_dep_.CoverTab[77953]++
																return nil, io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:311
			// _ = "end of CoverTab[77953]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:312
			_go_fuzz_dep_.CoverTab[77954]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:312
			// _ = "end of CoverTab[77954]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:312
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:312
		// _ = "end of CoverTab[77951]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:312
		_go_fuzz_dep_.CoverTab[77952]++
															return nil, connectionErrorf(false, err, "transport: http2Server.HandleStreams failed to receive the preface from client: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:313
		// _ = "end of CoverTab[77952]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:314
		_go_fuzz_dep_.CoverTab[77955]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:314
		// _ = "end of CoverTab[77955]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:314
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:314
	// _ = "end of CoverTab[77892]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:314
	_go_fuzz_dep_.CoverTab[77893]++
														if !bytes.Equal(preface, clientPreface) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:315
		_go_fuzz_dep_.CoverTab[77956]++
															return nil, connectionErrorf(false, nil, "transport: http2Server.HandleStreams received bogus greeting from client: %q", preface)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:316
		// _ = "end of CoverTab[77956]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:317
		_go_fuzz_dep_.CoverTab[77957]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:317
		// _ = "end of CoverTab[77957]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:317
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:317
	// _ = "end of CoverTab[77893]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:317
	_go_fuzz_dep_.CoverTab[77894]++

														frame, err := t.framer.fr.ReadFrame()
														if err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:320
		_go_fuzz_dep_.CoverTab[77958]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:320
		return err == io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:320
		// _ = "end of CoverTab[77958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:320
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:320
		_go_fuzz_dep_.CoverTab[77959]++
															return nil, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:321
		// _ = "end of CoverTab[77959]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:322
		_go_fuzz_dep_.CoverTab[77960]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:322
		// _ = "end of CoverTab[77960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:322
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:322
	// _ = "end of CoverTab[77894]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:322
	_go_fuzz_dep_.CoverTab[77895]++
														if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:323
		_go_fuzz_dep_.CoverTab[77961]++
															return nil, connectionErrorf(false, err, "transport: http2Server.HandleStreams failed to read initial settings frame: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:324
		// _ = "end of CoverTab[77961]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:325
		_go_fuzz_dep_.CoverTab[77962]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:325
		// _ = "end of CoverTab[77962]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:325
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:325
	// _ = "end of CoverTab[77895]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:325
	_go_fuzz_dep_.CoverTab[77896]++
														atomic.StoreInt64(&t.lastRead, time.Now().UnixNano())
														sf, ok := frame.(*http2.SettingsFrame)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:328
		_go_fuzz_dep_.CoverTab[77963]++
															return nil, connectionErrorf(false, nil, "transport: http2Server.HandleStreams saw invalid preface type %T from client", frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:329
		// _ = "end of CoverTab[77963]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:330
		_go_fuzz_dep_.CoverTab[77964]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:330
		// _ = "end of CoverTab[77964]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:330
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:330
	// _ = "end of CoverTab[77896]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:330
	_go_fuzz_dep_.CoverTab[77897]++
														t.handleSettings(sf)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:331
	_curRoutineNum82_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:331
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum82_)

														go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:333
		_go_fuzz_dep_.CoverTab[77965]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:333
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:333
			_go_fuzz_dep_.CoverTab[77966]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:333
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum82_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:333
			// _ = "end of CoverTab[77966]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:333
		}()
															t.loopy = newLoopyWriter(serverSide, t.framer, t.controlBuf, t.bdpEst, t.conn)
															t.loopy.ssGoAwayHandler = t.outgoingGoAwayHandler
															t.loopy.run()
															close(t.writerDone)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:337
		// _ = "end of CoverTab[77965]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:338
	// _ = "end of CoverTab[77897]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:338
	_go_fuzz_dep_.CoverTab[77898]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:338
	_curRoutineNum83_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:338
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum83_)
														go t.keepalive()
														return t, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:340
	// _ = "end of CoverTab[77898]"
}

// operateHeaders takes action on the decoded headers. Returns an error if fatal
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:343
// error encountered and transport needs to close, otherwise returns nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:345
func (t *http2Server) operateHeaders(frame *http2.MetaHeadersFrame, handle func(*Stream), traceCtx func(context.Context, string) context.Context) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:345
	_go_fuzz_dep_.CoverTab[77967]++

														t.maxStreamMu.Lock()
														defer t.maxStreamMu.Unlock()

														streamID := frame.Header().StreamID

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:354
	if frame.Truncated {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:354
		_go_fuzz_dep_.CoverTab[77988]++
															t.controlBuf.put(&cleanupStream{
			streamID:	streamID,
			rst:		true,
			rstCode:	http2.ErrCodeFrameSize,
			onWrite:	func() { _go_fuzz_dep_.CoverTab[77990]++; // _ = "end of CoverTab[77990]" },
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:360
		// _ = "end of CoverTab[77988]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:360
		_go_fuzz_dep_.CoverTab[77989]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:361
		// _ = "end of CoverTab[77989]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:362
		_go_fuzz_dep_.CoverTab[77991]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:362
		// _ = "end of CoverTab[77991]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:362
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:362
	// _ = "end of CoverTab[77967]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:362
	_go_fuzz_dep_.CoverTab[77968]++

														if streamID%2 != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:364
		_go_fuzz_dep_.CoverTab[77992]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:364
		return streamID <= t.maxStreamID
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:364
		// _ = "end of CoverTab[77992]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:364
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:364
		_go_fuzz_dep_.CoverTab[77993]++

															return fmt.Errorf("received an illegal stream id: %v. headers frame: %+v", streamID, frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:366
		// _ = "end of CoverTab[77993]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:367
		_go_fuzz_dep_.CoverTab[77994]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:367
		// _ = "end of CoverTab[77994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:367
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:367
	// _ = "end of CoverTab[77968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:367
	_go_fuzz_dep_.CoverTab[77969]++
														t.maxStreamID = streamID

														buf := newRecvBuffer()
														s := &Stream{
		id:	streamID,
		st:	t,
		buf:	buf,
		fc:	&inFlow{limit: uint32(t.initialWindowSize)},
	}
	var (
		// if false, content-type was missing or invalid
		isGRPC		= false
		contentType	= ""
		mdata		= make(metadata.MD, len(frame.Fields))
		httpMethod	string
		// these are set if an error is encountered while parsing the headers
		protocolError	bool
		headerError	*status.Status

		timeoutSet	bool
		timeout		time.Duration
	)

	for _, hf := range frame.Fields {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:391
		_go_fuzz_dep_.CoverTab[77995]++
															switch hf.Name {
		case "content-type":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:393
			_go_fuzz_dep_.CoverTab[77996]++
																contentSubtype, validContentType := grpcutil.ContentSubtype(hf.Value)
																if !validContentType {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:395
				_go_fuzz_dep_.CoverTab[78010]++
																	contentType = hf.Value
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:397
				// _ = "end of CoverTab[78010]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:398
				_go_fuzz_dep_.CoverTab[78011]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:398
				// _ = "end of CoverTab[78011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:398
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:398
			// _ = "end of CoverTab[77996]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:398
			_go_fuzz_dep_.CoverTab[77997]++
																mdata[hf.Name] = append(mdata[hf.Name], hf.Value)
																s.contentSubtype = contentSubtype
																isGRPC = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:401
			// _ = "end of CoverTab[77997]"

		case "grpc-accept-encoding":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:403
			_go_fuzz_dep_.CoverTab[77998]++
																mdata[hf.Name] = append(mdata[hf.Name], hf.Value)
																if hf.Value == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:405
				_go_fuzz_dep_.CoverTab[78012]++
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:406
				// _ = "end of CoverTab[78012]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:407
				_go_fuzz_dep_.CoverTab[78013]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:407
				// _ = "end of CoverTab[78013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:407
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:407
			// _ = "end of CoverTab[77998]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:407
			_go_fuzz_dep_.CoverTab[77999]++
																compressors := hf.Value
																if s.clientAdvertisedCompressors != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:409
				_go_fuzz_dep_.CoverTab[78014]++
																	compressors = s.clientAdvertisedCompressors + "," + compressors
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:410
				// _ = "end of CoverTab[78014]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:411
				_go_fuzz_dep_.CoverTab[78015]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:411
				// _ = "end of CoverTab[78015]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:411
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:411
			// _ = "end of CoverTab[77999]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:411
			_go_fuzz_dep_.CoverTab[78000]++
																s.clientAdvertisedCompressors = compressors
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:412
			// _ = "end of CoverTab[78000]"
		case "grpc-encoding":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:413
			_go_fuzz_dep_.CoverTab[78001]++
																s.recvCompress = hf.Value
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:414
			// _ = "end of CoverTab[78001]"
		case ":method":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:415
			_go_fuzz_dep_.CoverTab[78002]++
																httpMethod = hf.Value
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:416
			// _ = "end of CoverTab[78002]"
		case ":path":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:417
			_go_fuzz_dep_.CoverTab[78003]++
																s.method = hf.Value
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:418
			// _ = "end of CoverTab[78003]"
		case "grpc-timeout":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:419
			_go_fuzz_dep_.CoverTab[78004]++
																timeoutSet = true
																var err error
																if timeout, err = decodeTimeout(hf.Value); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:422
				_go_fuzz_dep_.CoverTab[78016]++
																	headerError = status.Newf(codes.Internal, "malformed grpc-timeout: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:423
				// _ = "end of CoverTab[78016]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:424
				_go_fuzz_dep_.CoverTab[78017]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:424
				// _ = "end of CoverTab[78017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:424
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:424
			// _ = "end of CoverTab[78004]"

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:427
		case "connection":
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:427
			_go_fuzz_dep_.CoverTab[78005]++
																if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:428
				_go_fuzz_dep_.CoverTab[78018]++
																	logger.Errorf("transport: http2Server.operateHeaders parsed a :connection header which makes a request malformed as per the HTTP/2 spec")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:429
				// _ = "end of CoverTab[78018]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:430
				_go_fuzz_dep_.CoverTab[78019]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:430
				// _ = "end of CoverTab[78019]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:430
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:430
			// _ = "end of CoverTab[78005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:430
			_go_fuzz_dep_.CoverTab[78006]++
																protocolError = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:431
			// _ = "end of CoverTab[78006]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:432
			_go_fuzz_dep_.CoverTab[78007]++
																if isReservedHeader(hf.Name) && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:433
				_go_fuzz_dep_.CoverTab[78020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:433
				return !isWhitelistedHeader(hf.Name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:433
				// _ = "end of CoverTab[78020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:433
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:433
				_go_fuzz_dep_.CoverTab[78021]++
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:434
				// _ = "end of CoverTab[78021]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:435
				_go_fuzz_dep_.CoverTab[78022]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:435
				// _ = "end of CoverTab[78022]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:435
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:435
			// _ = "end of CoverTab[78007]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:435
			_go_fuzz_dep_.CoverTab[78008]++
																v, err := decodeMetadataHeader(hf.Name, hf.Value)
																if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:437
				_go_fuzz_dep_.CoverTab[78023]++
																	headerError = status.Newf(codes.Internal, "malformed binary metadata %q in header %q: %v", hf.Value, hf.Name, err)
																	logger.Warningf("Failed to decode metadata header (%q, %q): %v", hf.Name, hf.Value, err)
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:440
				// _ = "end of CoverTab[78023]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:441
				_go_fuzz_dep_.CoverTab[78024]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:441
				// _ = "end of CoverTab[78024]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:441
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:441
			// _ = "end of CoverTab[78008]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:441
			_go_fuzz_dep_.CoverTab[78009]++
																mdata[hf.Name] = append(mdata[hf.Name], v)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:442
			// _ = "end of CoverTab[78009]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:443
		// _ = "end of CoverTab[77995]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:444
	// _ = "end of CoverTab[77969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:444
	_go_fuzz_dep_.CoverTab[77970]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:451
	if len(mdata[":authority"]) > 1 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:451
		_go_fuzz_dep_.CoverTab[78025]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:451
		return len(mdata["host"]) > 1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:451
		// _ = "end of CoverTab[78025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:451
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:451
		_go_fuzz_dep_.CoverTab[78026]++
															errMsg := fmt.Sprintf("num values of :authority: %v, num values of host: %v, both must only have 1 value as per HTTP/2 spec", len(mdata[":authority"]), len(mdata["host"]))
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:453
			_go_fuzz_dep_.CoverTab[78028]++
																logger.Errorf("transport: %v", errMsg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:454
			// _ = "end of CoverTab[78028]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:455
			_go_fuzz_dep_.CoverTab[78029]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:455
			// _ = "end of CoverTab[78029]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:455
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:455
		// _ = "end of CoverTab[78026]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:455
		_go_fuzz_dep_.CoverTab[78027]++
															t.controlBuf.put(&earlyAbortStream{
			httpStatus:	http.StatusBadRequest,
			streamID:	streamID,
			contentSubtype:	s.contentSubtype,
			status:		status.New(codes.Internal, errMsg),
			rst:		!frame.StreamEnded(),
		})
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:463
		// _ = "end of CoverTab[78027]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:464
		_go_fuzz_dep_.CoverTab[78030]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:464
		// _ = "end of CoverTab[78030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:464
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:464
	// _ = "end of CoverTab[77970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:464
	_go_fuzz_dep_.CoverTab[77971]++

														if protocolError {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:466
		_go_fuzz_dep_.CoverTab[78031]++
															t.controlBuf.put(&cleanupStream{
			streamID:	streamID,
			rst:		true,
			rstCode:	http2.ErrCodeProtocol,
			onWrite:	func() { _go_fuzz_dep_.CoverTab[78033]++; // _ = "end of CoverTab[78033]" },
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:472
		// _ = "end of CoverTab[78031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:472
		_go_fuzz_dep_.CoverTab[78032]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:473
		// _ = "end of CoverTab[78032]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:474
		_go_fuzz_dep_.CoverTab[78034]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:474
		// _ = "end of CoverTab[78034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:474
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:474
	// _ = "end of CoverTab[77971]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:474
	_go_fuzz_dep_.CoverTab[77972]++
														if !isGRPC {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:475
		_go_fuzz_dep_.CoverTab[78035]++
															t.controlBuf.put(&earlyAbortStream{
			httpStatus:	http.StatusUnsupportedMediaType,
			streamID:	streamID,
			contentSubtype:	s.contentSubtype,
			status:		status.Newf(codes.InvalidArgument, "invalid gRPC request content-type %q", contentType),
			rst:		!frame.StreamEnded(),
		})
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:483
		// _ = "end of CoverTab[78035]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:484
		_go_fuzz_dep_.CoverTab[78036]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:484
		// _ = "end of CoverTab[78036]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:484
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:484
	// _ = "end of CoverTab[77972]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:484
	_go_fuzz_dep_.CoverTab[77973]++
														if headerError != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:485
		_go_fuzz_dep_.CoverTab[78037]++
															t.controlBuf.put(&earlyAbortStream{
			httpStatus:	http.StatusBadRequest,
			streamID:	streamID,
			contentSubtype:	s.contentSubtype,
			status:		headerError,
			rst:		!frame.StreamEnded(),
		})
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:493
		// _ = "end of CoverTab[78037]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:494
		_go_fuzz_dep_.CoverTab[78038]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:494
		// _ = "end of CoverTab[78038]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:494
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:494
	// _ = "end of CoverTab[77973]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:494
	_go_fuzz_dep_.CoverTab[77974]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:497
	if len(mdata[":authority"]) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:497
		_go_fuzz_dep_.CoverTab[78039]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:500
		if host, ok := mdata["host"]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:500
			_go_fuzz_dep_.CoverTab[78040]++
																mdata[":authority"] = host
																delete(mdata, "host")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:502
			// _ = "end of CoverTab[78040]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:503
			_go_fuzz_dep_.CoverTab[78041]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:503
			// _ = "end of CoverTab[78041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:503
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:503
		// _ = "end of CoverTab[78039]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:504
		_go_fuzz_dep_.CoverTab[78042]++

															delete(mdata, "host")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:506
		// _ = "end of CoverTab[78042]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:507
	// _ = "end of CoverTab[77974]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:507
	_go_fuzz_dep_.CoverTab[77975]++

														if frame.StreamEnded() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:509
		_go_fuzz_dep_.CoverTab[78043]++

															s.state = streamReadDone
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:511
		// _ = "end of CoverTab[78043]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:512
		_go_fuzz_dep_.CoverTab[78044]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:512
		// _ = "end of CoverTab[78044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:512
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:512
	// _ = "end of CoverTab[77975]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:512
	_go_fuzz_dep_.CoverTab[77976]++
														if timeoutSet {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:513
		_go_fuzz_dep_.CoverTab[78045]++
															s.ctx, s.cancel = context.WithTimeout(t.ctx, timeout)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:514
		// _ = "end of CoverTab[78045]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:515
		_go_fuzz_dep_.CoverTab[78046]++
															s.ctx, s.cancel = context.WithCancel(t.ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:516
		// _ = "end of CoverTab[78046]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:517
	// _ = "end of CoverTab[77976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:517
	_go_fuzz_dep_.CoverTab[77977]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:520
	if len(mdata) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:520
		_go_fuzz_dep_.CoverTab[78047]++
															s.ctx = metadata.NewIncomingContext(s.ctx, mdata)
															if statsTags := mdata["grpc-tags-bin"]; len(statsTags) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:522
			_go_fuzz_dep_.CoverTab[78049]++
																s.ctx = stats.SetIncomingTags(s.ctx, []byte(statsTags[len(statsTags)-1]))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:523
			// _ = "end of CoverTab[78049]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:524
			_go_fuzz_dep_.CoverTab[78050]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:524
			// _ = "end of CoverTab[78050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:524
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:524
		// _ = "end of CoverTab[78047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:524
		_go_fuzz_dep_.CoverTab[78048]++
															if statsTrace := mdata["grpc-trace-bin"]; len(statsTrace) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:525
			_go_fuzz_dep_.CoverTab[78051]++
																s.ctx = stats.SetIncomingTrace(s.ctx, []byte(statsTrace[len(statsTrace)-1]))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:526
			// _ = "end of CoverTab[78051]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:527
			_go_fuzz_dep_.CoverTab[78052]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:527
			// _ = "end of CoverTab[78052]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:527
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:527
		// _ = "end of CoverTab[78048]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:528
		_go_fuzz_dep_.CoverTab[78053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:528
		// _ = "end of CoverTab[78053]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:528
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:528
	// _ = "end of CoverTab[77977]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:528
	_go_fuzz_dep_.CoverTab[77978]++
														t.mu.Lock()
														if t.state != reachable {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:530
		_go_fuzz_dep_.CoverTab[78054]++
															t.mu.Unlock()
															s.cancel()
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:533
		// _ = "end of CoverTab[78054]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:534
		_go_fuzz_dep_.CoverTab[78055]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:534
		// _ = "end of CoverTab[78055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:534
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:534
	// _ = "end of CoverTab[77978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:534
	_go_fuzz_dep_.CoverTab[77979]++
														if uint32(len(t.activeStreams)) >= t.maxStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:535
		_go_fuzz_dep_.CoverTab[78056]++
															t.mu.Unlock()
															t.controlBuf.put(&cleanupStream{
			streamID:	streamID,
			rst:		true,
			rstCode:	http2.ErrCodeRefusedStream,
			onWrite:	func() { _go_fuzz_dep_.CoverTab[78058]++; // _ = "end of CoverTab[78058]" },
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:542
		// _ = "end of CoverTab[78056]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:542
		_go_fuzz_dep_.CoverTab[78057]++
															s.cancel()
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:544
		// _ = "end of CoverTab[78057]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:545
		_go_fuzz_dep_.CoverTab[78059]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:545
		// _ = "end of CoverTab[78059]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:545
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:545
	// _ = "end of CoverTab[77979]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:545
	_go_fuzz_dep_.CoverTab[77980]++
														if httpMethod != http.MethodPost {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:546
		_go_fuzz_dep_.CoverTab[78060]++
															t.mu.Unlock()
															errMsg := fmt.Sprintf("http2Server.operateHeaders parsed a :method field: %v which should be POST", httpMethod)
															if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:549
			_go_fuzz_dep_.CoverTab[78062]++
																logger.Infof("transport: %v", errMsg)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:550
			// _ = "end of CoverTab[78062]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:551
			_go_fuzz_dep_.CoverTab[78063]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:551
			// _ = "end of CoverTab[78063]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:551
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:551
		// _ = "end of CoverTab[78060]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:551
		_go_fuzz_dep_.CoverTab[78061]++
															t.controlBuf.put(&earlyAbortStream{
			httpStatus:	405,
			streamID:	streamID,
			contentSubtype:	s.contentSubtype,
			status:		status.New(codes.Internal, errMsg),
			rst:		!frame.StreamEnded(),
		})
															s.cancel()
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:560
		// _ = "end of CoverTab[78061]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:561
		_go_fuzz_dep_.CoverTab[78064]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:561
		// _ = "end of CoverTab[78064]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:561
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:561
	// _ = "end of CoverTab[77980]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:561
	_go_fuzz_dep_.CoverTab[77981]++
														if t.inTapHandle != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:562
		_go_fuzz_dep_.CoverTab[78065]++
															var err error
															if s.ctx, err = t.inTapHandle(s.ctx, &tap.Info{FullMethodName: s.method}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:564
			_go_fuzz_dep_.CoverTab[78066]++
																t.mu.Unlock()
																if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:566
				_go_fuzz_dep_.CoverTab[78069]++
																	logger.Infof("transport: http2Server.operateHeaders got an error from InTapHandle: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:567
				// _ = "end of CoverTab[78069]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:568
				_go_fuzz_dep_.CoverTab[78070]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:568
				// _ = "end of CoverTab[78070]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:568
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:568
			// _ = "end of CoverTab[78066]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:568
			_go_fuzz_dep_.CoverTab[78067]++
																stat, ok := status.FromError(err)
																if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:570
				_go_fuzz_dep_.CoverTab[78071]++
																	stat = status.New(codes.PermissionDenied, err.Error())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:571
				// _ = "end of CoverTab[78071]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:572
				_go_fuzz_dep_.CoverTab[78072]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:572
				// _ = "end of CoverTab[78072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:572
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:572
			// _ = "end of CoverTab[78067]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:572
			_go_fuzz_dep_.CoverTab[78068]++
																t.controlBuf.put(&earlyAbortStream{
				httpStatus:	200,
				streamID:	s.id,
				contentSubtype:	s.contentSubtype,
				status:		stat,
				rst:		!frame.StreamEnded(),
			})
																return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:580
			// _ = "end of CoverTab[78068]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:581
			_go_fuzz_dep_.CoverTab[78073]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:581
			// _ = "end of CoverTab[78073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:581
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:581
		// _ = "end of CoverTab[78065]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:582
		_go_fuzz_dep_.CoverTab[78074]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:582
		// _ = "end of CoverTab[78074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:582
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:582
	// _ = "end of CoverTab[77981]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:582
	_go_fuzz_dep_.CoverTab[77982]++
														t.activeStreams[streamID] = s
														if len(t.activeStreams) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:584
		_go_fuzz_dep_.CoverTab[78075]++
															t.idle = time.Time{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:585
		// _ = "end of CoverTab[78075]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:586
		_go_fuzz_dep_.CoverTab[78076]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:586
		// _ = "end of CoverTab[78076]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:586
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:586
	// _ = "end of CoverTab[77982]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:586
	_go_fuzz_dep_.CoverTab[77983]++
														t.mu.Unlock()
														if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:588
		_go_fuzz_dep_.CoverTab[78077]++
															atomic.AddInt64(&t.czData.streamsStarted, 1)
															atomic.StoreInt64(&t.czData.lastStreamCreatedTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:590
		// _ = "end of CoverTab[78077]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:591
		_go_fuzz_dep_.CoverTab[78078]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:591
		// _ = "end of CoverTab[78078]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:591
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:591
	// _ = "end of CoverTab[77983]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:591
	_go_fuzz_dep_.CoverTab[77984]++
														s.requestRead = func(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:592
		_go_fuzz_dep_.CoverTab[78079]++
															t.adjustWindow(s, uint32(n))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:593
		// _ = "end of CoverTab[78079]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:594
	// _ = "end of CoverTab[77984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:594
	_go_fuzz_dep_.CoverTab[77985]++
														s.ctx = traceCtx(s.ctx, s.method)
														for _, sh := range t.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:596
		_go_fuzz_dep_.CoverTab[78080]++
															s.ctx = sh.TagRPC(s.ctx, &stats.RPCTagInfo{FullMethodName: s.method})
															inHeader := &stats.InHeader{
			FullMethod:	s.method,
			RemoteAddr:	t.remoteAddr,
			LocalAddr:	t.localAddr,
			Compression:	s.recvCompress,
			WireLength:	int(frame.Header().Length),
			Header:		mdata.Copy(),
		}
															sh.HandleRPC(s.ctx, inHeader)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:606
		// _ = "end of CoverTab[78080]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:607
	// _ = "end of CoverTab[77985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:607
	_go_fuzz_dep_.CoverTab[77986]++
														s.ctxDone = s.ctx.Done()
														s.wq = newWriteQuota(defaultWriteQuota, s.ctxDone)
														s.trReader = &transportReader{
		reader: &recvBufferReader{
			ctx:		s.ctx,
			ctxDone:	s.ctxDone,
			recv:		s.buf,
			freeBuffer:	t.bufferPool.put,
		},
		windowHandler: func(n int) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:617
			_go_fuzz_dep_.CoverTab[78081]++
																t.updateWindow(s, uint32(n))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:618
			// _ = "end of CoverTab[78081]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:620
	// _ = "end of CoverTab[77986]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:620
	_go_fuzz_dep_.CoverTab[77987]++

														t.controlBuf.put(&registerStream{
		streamID:	s.id,
		wq:		s.wq,
	})
														handle(s)
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:627
	// _ = "end of CoverTab[77987]"
}

// HandleStreams receives incoming streams using the given handler. This is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:630
// typically run in a separate goroutine.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:630
// traceCtx attaches trace to ctx and returns the new context.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:633
func (t *http2Server) HandleStreams(handle func(*Stream), traceCtx func(context.Context, string) context.Context) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:633
	_go_fuzz_dep_.CoverTab[78082]++
														defer close(t.readerDone)
														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:635
		_go_fuzz_dep_.CoverTab[78083]++
															t.controlBuf.throttle()
															frame, err := t.framer.fr.ReadFrame()
															atomic.StoreInt64(&t.lastRead, time.Now().UnixNano())
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:639
			_go_fuzz_dep_.CoverTab[78085]++
																if se, ok := err.(http2.StreamError); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:640
				_go_fuzz_dep_.CoverTab[78088]++
																	if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:641
					_go_fuzz_dep_.CoverTab[78091]++
																		logger.Warningf("transport: http2Server.HandleStreams encountered http2.StreamError: %v", se)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:642
					// _ = "end of CoverTab[78091]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:643
					_go_fuzz_dep_.CoverTab[78092]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:643
					// _ = "end of CoverTab[78092]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:643
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:643
				// _ = "end of CoverTab[78088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:643
				_go_fuzz_dep_.CoverTab[78089]++
																	t.mu.Lock()
																	s := t.activeStreams[se.StreamID]
																	t.mu.Unlock()
																	if s != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:647
					_go_fuzz_dep_.CoverTab[78093]++
																		t.closeStream(s, true, se.Code, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:648
					// _ = "end of CoverTab[78093]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:649
					_go_fuzz_dep_.CoverTab[78094]++
																		t.controlBuf.put(&cleanupStream{
						streamID:	se.StreamID,
						rst:		true,
						rstCode:	se.Code,
						onWrite:	func() { _go_fuzz_dep_.CoverTab[78095]++; // _ = "end of CoverTab[78095]" },
					})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:655
					// _ = "end of CoverTab[78094]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:656
				// _ = "end of CoverTab[78089]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:656
				_go_fuzz_dep_.CoverTab[78090]++
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:657
				// _ = "end of CoverTab[78090]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:658
				_go_fuzz_dep_.CoverTab[78096]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:658
				// _ = "end of CoverTab[78096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:658
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:658
			// _ = "end of CoverTab[78085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:658
			_go_fuzz_dep_.CoverTab[78086]++
																if err == io.EOF || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:659
				_go_fuzz_dep_.CoverTab[78097]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:659
				return err == io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:659
				// _ = "end of CoverTab[78097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:659
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:659
				_go_fuzz_dep_.CoverTab[78098]++
																	t.Close(err)
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:661
				// _ = "end of CoverTab[78098]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:662
				_go_fuzz_dep_.CoverTab[78099]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:662
				// _ = "end of CoverTab[78099]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:662
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:662
			// _ = "end of CoverTab[78086]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:662
			_go_fuzz_dep_.CoverTab[78087]++
																t.Close(err)
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:664
			// _ = "end of CoverTab[78087]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:665
			_go_fuzz_dep_.CoverTab[78100]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:665
			// _ = "end of CoverTab[78100]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:665
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:665
		// _ = "end of CoverTab[78083]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:665
		_go_fuzz_dep_.CoverTab[78084]++
															switch frame := frame.(type) {
		case *http2.MetaHeadersFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:667
			_go_fuzz_dep_.CoverTab[78101]++
																if err := t.operateHeaders(frame, handle, traceCtx); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:668
				_go_fuzz_dep_.CoverTab[78109]++
																	t.Close(err)
																	break
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:670
				// _ = "end of CoverTab[78109]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:671
				_go_fuzz_dep_.CoverTab[78110]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:671
				// _ = "end of CoverTab[78110]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:671
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:671
			// _ = "end of CoverTab[78101]"
		case *http2.DataFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:672
			_go_fuzz_dep_.CoverTab[78102]++
																t.handleData(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:673
			// _ = "end of CoverTab[78102]"
		case *http2.RSTStreamFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:674
			_go_fuzz_dep_.CoverTab[78103]++
																t.handleRSTStream(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:675
			// _ = "end of CoverTab[78103]"
		case *http2.SettingsFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:676
			_go_fuzz_dep_.CoverTab[78104]++
																t.handleSettings(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:677
			// _ = "end of CoverTab[78104]"
		case *http2.PingFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:678
			_go_fuzz_dep_.CoverTab[78105]++
																t.handlePing(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:679
			// _ = "end of CoverTab[78105]"
		case *http2.WindowUpdateFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:680
			_go_fuzz_dep_.CoverTab[78106]++
																t.handleWindowUpdate(frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:681
			// _ = "end of CoverTab[78106]"
		case *http2.GoAwayFrame:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:682
			_go_fuzz_dep_.CoverTab[78107]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:682
			// _ = "end of CoverTab[78107]"

		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:684
			_go_fuzz_dep_.CoverTab[78108]++
																if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:685
				_go_fuzz_dep_.CoverTab[78111]++
																	logger.Errorf("transport: http2Server.HandleStreams found unhandled frame type %v.", frame)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:686
				// _ = "end of CoverTab[78111]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:687
				_go_fuzz_dep_.CoverTab[78112]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:687
				// _ = "end of CoverTab[78112]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:687
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:687
			// _ = "end of CoverTab[78108]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:688
		// _ = "end of CoverTab[78084]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:689
	// _ = "end of CoverTab[78082]"
}

func (t *http2Server) getStream(f http2.Frame) (*Stream, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:692
	_go_fuzz_dep_.CoverTab[78113]++
														t.mu.Lock()
														defer t.mu.Unlock()
														if t.activeStreams == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:695
		_go_fuzz_dep_.CoverTab[78116]++

															return nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:697
		// _ = "end of CoverTab[78116]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:698
		_go_fuzz_dep_.CoverTab[78117]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:698
		// _ = "end of CoverTab[78117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:698
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:698
	// _ = "end of CoverTab[78113]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:698
	_go_fuzz_dep_.CoverTab[78114]++
														s, ok := t.activeStreams[f.Header().StreamID]
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:700
		_go_fuzz_dep_.CoverTab[78118]++

															return nil, false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:702
		// _ = "end of CoverTab[78118]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:703
		_go_fuzz_dep_.CoverTab[78119]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:703
		// _ = "end of CoverTab[78119]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:703
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:703
	// _ = "end of CoverTab[78114]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:703
	_go_fuzz_dep_.CoverTab[78115]++
														return s, true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:704
	// _ = "end of CoverTab[78115]"
}

// adjustWindow sends out extra window update over the initial window size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:707
// of stream if the application is requesting data larger in size than
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:707
// the window.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:710
func (t *http2Server) adjustWindow(s *Stream, n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:710
	_go_fuzz_dep_.CoverTab[78120]++
														if w := s.fc.maybeAdjust(n); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:711
		_go_fuzz_dep_.CoverTab[78121]++
															t.controlBuf.put(&outgoingWindowUpdate{streamID: s.id, increment: w})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:712
		// _ = "end of CoverTab[78121]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:713
		_go_fuzz_dep_.CoverTab[78122]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:713
		// _ = "end of CoverTab[78122]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:713
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:713
	// _ = "end of CoverTab[78120]"

}

// updateWindow adjusts the inbound quota for the stream and the transport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:717
// Window updates will deliver to the controller for sending when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:717
// the cumulative quota exceeds the corresponding threshold.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:720
func (t *http2Server) updateWindow(s *Stream, n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:720
	_go_fuzz_dep_.CoverTab[78123]++
														if w := s.fc.onRead(n); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:721
		_go_fuzz_dep_.CoverTab[78124]++
															t.controlBuf.put(&outgoingWindowUpdate{streamID: s.id,
			increment:	w,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:724
		// _ = "end of CoverTab[78124]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:725
		_go_fuzz_dep_.CoverTab[78125]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:725
		// _ = "end of CoverTab[78125]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:725
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:725
	// _ = "end of CoverTab[78123]"
}

// updateFlowControl updates the incoming flow control windows
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:728
// for the transport and the stream based on the current bdp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:728
// estimation.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:731
func (t *http2Server) updateFlowControl(n uint32) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:731
	_go_fuzz_dep_.CoverTab[78126]++
														t.mu.Lock()
														for _, s := range t.activeStreams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:733
		_go_fuzz_dep_.CoverTab[78128]++
															s.fc.newLimit(n)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:734
		// _ = "end of CoverTab[78128]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:735
	// _ = "end of CoverTab[78126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:735
	_go_fuzz_dep_.CoverTab[78127]++
														t.initialWindowSize = int32(n)
														t.mu.Unlock()
														t.controlBuf.put(&outgoingWindowUpdate{
		streamID:	0,
		increment:	t.fc.newLimit(n),
	})
	t.controlBuf.put(&outgoingSettings{
		ss: []http2.Setting{
			{
				ID:	http2.SettingInitialWindowSize,
				Val:	n,
			},
		},
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:749
	// _ = "end of CoverTab[78127]"

}

func (t *http2Server) handleData(f *http2.DataFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:753
	_go_fuzz_dep_.CoverTab[78129]++
														size := f.Header().Length
														var sendBDPPing bool
														if t.bdpEst != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:756
		_go_fuzz_dep_.CoverTab[78136]++
															sendBDPPing = t.bdpEst.add(size)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:757
		// _ = "end of CoverTab[78136]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:758
		_go_fuzz_dep_.CoverTab[78137]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:758
		// _ = "end of CoverTab[78137]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:758
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:758
	// _ = "end of CoverTab[78129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:758
	_go_fuzz_dep_.CoverTab[78130]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:767
	if w := t.fc.onData(size); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:767
		_go_fuzz_dep_.CoverTab[78138]++
															t.controlBuf.put(&outgoingWindowUpdate{
			streamID:	0,
			increment:	w,
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:771
		// _ = "end of CoverTab[78138]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:772
		_go_fuzz_dep_.CoverTab[78139]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:772
		// _ = "end of CoverTab[78139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:772
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:772
	// _ = "end of CoverTab[78130]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:772
	_go_fuzz_dep_.CoverTab[78131]++
														if sendBDPPing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:773
		_go_fuzz_dep_.CoverTab[78140]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:776
		if w := t.fc.reset(); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:776
			_go_fuzz_dep_.CoverTab[78142]++
																t.controlBuf.put(&outgoingWindowUpdate{
				streamID:	0,
				increment:	w,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:780
			// _ = "end of CoverTab[78142]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:781
			_go_fuzz_dep_.CoverTab[78143]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:781
			// _ = "end of CoverTab[78143]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:781
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:781
		// _ = "end of CoverTab[78140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:781
		_go_fuzz_dep_.CoverTab[78141]++
															t.controlBuf.put(bdpPing)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:782
		// _ = "end of CoverTab[78141]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:783
		_go_fuzz_dep_.CoverTab[78144]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:783
		// _ = "end of CoverTab[78144]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:783
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:783
	// _ = "end of CoverTab[78131]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:783
	_go_fuzz_dep_.CoverTab[78132]++

														s, ok := t.getStream(f)
														if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:786
		_go_fuzz_dep_.CoverTab[78145]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:787
		// _ = "end of CoverTab[78145]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:788
		_go_fuzz_dep_.CoverTab[78146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:788
		// _ = "end of CoverTab[78146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:788
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:788
	// _ = "end of CoverTab[78132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:788
	_go_fuzz_dep_.CoverTab[78133]++
														if s.getState() == streamReadDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:789
		_go_fuzz_dep_.CoverTab[78147]++
															t.closeStream(s, true, http2.ErrCodeStreamClosed, false)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:791
		// _ = "end of CoverTab[78147]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:792
		_go_fuzz_dep_.CoverTab[78148]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:792
		// _ = "end of CoverTab[78148]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:792
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:792
	// _ = "end of CoverTab[78133]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:792
	_go_fuzz_dep_.CoverTab[78134]++
														if size > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:793
		_go_fuzz_dep_.CoverTab[78149]++
															if err := s.fc.onData(size); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:794
			_go_fuzz_dep_.CoverTab[78152]++
																t.closeStream(s, true, http2.ErrCodeFlowControl, false)
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:796
			// _ = "end of CoverTab[78152]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:797
			_go_fuzz_dep_.CoverTab[78153]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:797
			// _ = "end of CoverTab[78153]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:797
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:797
		// _ = "end of CoverTab[78149]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:797
		_go_fuzz_dep_.CoverTab[78150]++
															if f.Header().Flags.Has(http2.FlagDataPadded) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:798
			_go_fuzz_dep_.CoverTab[78154]++
																if w := s.fc.onRead(size - uint32(len(f.Data()))); w > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:799
				_go_fuzz_dep_.CoverTab[78155]++
																	t.controlBuf.put(&outgoingWindowUpdate{s.id, w})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:800
				// _ = "end of CoverTab[78155]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:801
				_go_fuzz_dep_.CoverTab[78156]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:801
				// _ = "end of CoverTab[78156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:801
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:801
			// _ = "end of CoverTab[78154]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:802
			_go_fuzz_dep_.CoverTab[78157]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:802
			// _ = "end of CoverTab[78157]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:802
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:802
		// _ = "end of CoverTab[78150]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:802
		_go_fuzz_dep_.CoverTab[78151]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:806
		if len(f.Data()) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:806
			_go_fuzz_dep_.CoverTab[78158]++
																buffer := t.bufferPool.get()
																buffer.Reset()
																buffer.Write(f.Data())
																s.write(recvMsg{buffer: buffer})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:810
			// _ = "end of CoverTab[78158]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:811
			_go_fuzz_dep_.CoverTab[78159]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:811
			// _ = "end of CoverTab[78159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:811
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:811
		// _ = "end of CoverTab[78151]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:812
		_go_fuzz_dep_.CoverTab[78160]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:812
		// _ = "end of CoverTab[78160]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:812
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:812
	// _ = "end of CoverTab[78134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:812
	_go_fuzz_dep_.CoverTab[78135]++
														if f.StreamEnded() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:813
		_go_fuzz_dep_.CoverTab[78161]++

															s.compareAndSwapState(streamActive, streamReadDone)
															s.write(recvMsg{err: io.EOF})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:816
		// _ = "end of CoverTab[78161]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:817
		_go_fuzz_dep_.CoverTab[78162]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:817
		// _ = "end of CoverTab[78162]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:817
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:817
	// _ = "end of CoverTab[78135]"
}

func (t *http2Server) handleRSTStream(f *http2.RSTStreamFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:820
	_go_fuzz_dep_.CoverTab[78163]++

														if s, ok := t.getStream(f); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:822
		_go_fuzz_dep_.CoverTab[78165]++
															t.closeStream(s, false, 0, false)
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:824
		// _ = "end of CoverTab[78165]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:825
		_go_fuzz_dep_.CoverTab[78166]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:825
		// _ = "end of CoverTab[78166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:825
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:825
	// _ = "end of CoverTab[78163]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:825
	_go_fuzz_dep_.CoverTab[78164]++

														t.controlBuf.put(&cleanupStream{
		streamID:	f.Header().StreamID,
		rst:		false,
		rstCode:	0,
		onWrite:	func() { _go_fuzz_dep_.CoverTab[78167]++; // _ = "end of CoverTab[78167]" },
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:832
	// _ = "end of CoverTab[78164]"
}

func (t *http2Server) handleSettings(f *http2.SettingsFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:835
	_go_fuzz_dep_.CoverTab[78168]++
														if f.IsAck() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:836
		_go_fuzz_dep_.CoverTab[78171]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:837
		// _ = "end of CoverTab[78171]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:838
		_go_fuzz_dep_.CoverTab[78172]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:838
		// _ = "end of CoverTab[78172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:838
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:838
	// _ = "end of CoverTab[78168]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:838
	_go_fuzz_dep_.CoverTab[78169]++
														var ss []http2.Setting
														var updateFuncs []func()
														f.ForeachSetting(func(s http2.Setting) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:841
		_go_fuzz_dep_.CoverTab[78173]++
															switch s.ID {
		case http2.SettingMaxHeaderListSize:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:843
			_go_fuzz_dep_.CoverTab[78175]++
																updateFuncs = append(updateFuncs, func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:844
				_go_fuzz_dep_.CoverTab[78177]++
																	t.maxSendHeaderListSize = new(uint32)
																	*t.maxSendHeaderListSize = s.Val
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:846
				// _ = "end of CoverTab[78177]"
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:847
			// _ = "end of CoverTab[78175]"
		default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:848
			_go_fuzz_dep_.CoverTab[78176]++
																ss = append(ss, s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:849
			// _ = "end of CoverTab[78176]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:850
		// _ = "end of CoverTab[78173]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:850
		_go_fuzz_dep_.CoverTab[78174]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:851
		// _ = "end of CoverTab[78174]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:852
	// _ = "end of CoverTab[78169]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:852
	_go_fuzz_dep_.CoverTab[78170]++
														t.controlBuf.executeAndPut(func(interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:853
		_go_fuzz_dep_.CoverTab[78178]++
															for _, f := range updateFuncs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:854
			_go_fuzz_dep_.CoverTab[78180]++
																f()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:855
			// _ = "end of CoverTab[78180]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:856
		// _ = "end of CoverTab[78178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:856
		_go_fuzz_dep_.CoverTab[78179]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:857
		// _ = "end of CoverTab[78179]"
	}, &incomingSettings{
		ss: ss,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:860
	// _ = "end of CoverTab[78170]"
}

const (
	maxPingStrikes		= 2
	defaultPingTimeout	= 2 * time.Hour
)

func (t *http2Server) handlePing(f *http2.PingFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:868
	_go_fuzz_dep_.CoverTab[78181]++
														if f.IsAck() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:869
		_go_fuzz_dep_.CoverTab[78186]++
															if f.Data == goAwayPing.data && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:870
			_go_fuzz_dep_.CoverTab[78189]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:870
			return t.drainEvent != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:870
			// _ = "end of CoverTab[78189]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:870
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:870
			_go_fuzz_dep_.CoverTab[78190]++
																t.drainEvent.Fire()
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:872
			// _ = "end of CoverTab[78190]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:873
			_go_fuzz_dep_.CoverTab[78191]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:873
			// _ = "end of CoverTab[78191]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:873
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:873
		// _ = "end of CoverTab[78186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:873
		_go_fuzz_dep_.CoverTab[78187]++

															if t.bdpEst != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:875
			_go_fuzz_dep_.CoverTab[78192]++
																t.bdpEst.calculate(f.Data)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:876
			// _ = "end of CoverTab[78192]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:877
			_go_fuzz_dep_.CoverTab[78193]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:877
			// _ = "end of CoverTab[78193]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:877
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:877
		// _ = "end of CoverTab[78187]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:877
		_go_fuzz_dep_.CoverTab[78188]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:878
		// _ = "end of CoverTab[78188]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:879
		_go_fuzz_dep_.CoverTab[78194]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:879
		// _ = "end of CoverTab[78194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:879
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:879
	// _ = "end of CoverTab[78181]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:879
	_go_fuzz_dep_.CoverTab[78182]++
														pingAck := &ping{ack: true}
														copy(pingAck.data[:], f.Data[:])
														t.controlBuf.put(pingAck)

														now := time.Now()
														defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:885
		_go_fuzz_dep_.CoverTab[78195]++
															t.lastPingAt = now
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:886
		// _ = "end of CoverTab[78195]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:887
	// _ = "end of CoverTab[78182]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:887
	_go_fuzz_dep_.CoverTab[78183]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:891
	if atomic.CompareAndSwapUint32(&t.resetPingStrikes, 1, 0) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:891
		_go_fuzz_dep_.CoverTab[78196]++
															t.pingStrikes = 0
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:893
		// _ = "end of CoverTab[78196]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:894
		_go_fuzz_dep_.CoverTab[78197]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:894
		// _ = "end of CoverTab[78197]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:894
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:894
	// _ = "end of CoverTab[78183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:894
	_go_fuzz_dep_.CoverTab[78184]++
														t.mu.Lock()
														ns := len(t.activeStreams)
														t.mu.Unlock()
														if ns < 1 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:898
		_go_fuzz_dep_.CoverTab[78198]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:898
		return !t.kep.PermitWithoutStream
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:898
		// _ = "end of CoverTab[78198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:898
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:898
		_go_fuzz_dep_.CoverTab[78199]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:901
		if t.lastPingAt.Add(defaultPingTimeout).After(now) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:901
			_go_fuzz_dep_.CoverTab[78200]++
																t.pingStrikes++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:902
			// _ = "end of CoverTab[78200]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:903
			_go_fuzz_dep_.CoverTab[78201]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:903
			// _ = "end of CoverTab[78201]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:903
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:903
		// _ = "end of CoverTab[78199]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:904
		_go_fuzz_dep_.CoverTab[78202]++

															if t.lastPingAt.Add(t.kep.MinTime).After(now) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:906
			_go_fuzz_dep_.CoverTab[78203]++
																t.pingStrikes++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:907
			// _ = "end of CoverTab[78203]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:908
			_go_fuzz_dep_.CoverTab[78204]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:908
			// _ = "end of CoverTab[78204]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:908
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:908
		// _ = "end of CoverTab[78202]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:909
	// _ = "end of CoverTab[78184]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:909
	_go_fuzz_dep_.CoverTab[78185]++

														if t.pingStrikes > maxPingStrikes {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:911
		_go_fuzz_dep_.CoverTab[78205]++

															t.controlBuf.put(&goAway{code: http2.ErrCodeEnhanceYourCalm, debugData: []byte("too_many_pings"), closeConn: errors.New("got too many pings from the client")})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:913
		// _ = "end of CoverTab[78205]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:914
		_go_fuzz_dep_.CoverTab[78206]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:914
		// _ = "end of CoverTab[78206]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:914
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:914
	// _ = "end of CoverTab[78185]"
}

func (t *http2Server) handleWindowUpdate(f *http2.WindowUpdateFrame) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:917
	_go_fuzz_dep_.CoverTab[78207]++
														t.controlBuf.put(&incomingWindowUpdate{
		streamID:	f.Header().StreamID,
		increment:	f.Increment,
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:921
	// _ = "end of CoverTab[78207]"
}

func appendHeaderFieldsFromMD(headerFields []hpack.HeaderField, md metadata.MD) []hpack.HeaderField {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:924
	_go_fuzz_dep_.CoverTab[78208]++
														for k, vv := range md {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:925
		_go_fuzz_dep_.CoverTab[78210]++
															if isReservedHeader(k) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:926
			_go_fuzz_dep_.CoverTab[78212]++

																continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:928
			// _ = "end of CoverTab[78212]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:929
			_go_fuzz_dep_.CoverTab[78213]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:929
			// _ = "end of CoverTab[78213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:929
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:929
		// _ = "end of CoverTab[78210]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:929
		_go_fuzz_dep_.CoverTab[78211]++
															for _, v := range vv {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:930
			_go_fuzz_dep_.CoverTab[78214]++
																headerFields = append(headerFields, hpack.HeaderField{Name: k, Value: encodeMetadataHeader(k, v)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:931
			// _ = "end of CoverTab[78214]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:932
		// _ = "end of CoverTab[78211]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:933
	// _ = "end of CoverTab[78208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:933
	_go_fuzz_dep_.CoverTab[78209]++
														return headerFields
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:934
	// _ = "end of CoverTab[78209]"
}

func (t *http2Server) checkForHeaderListSize(it interface{}) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:937
	_go_fuzz_dep_.CoverTab[78215]++
														if t.maxSendHeaderListSize == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:938
		_go_fuzz_dep_.CoverTab[78218]++
															return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:939
		// _ = "end of CoverTab[78218]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:940
		_go_fuzz_dep_.CoverTab[78219]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:940
		// _ = "end of CoverTab[78219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:940
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:940
	// _ = "end of CoverTab[78215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:940
	_go_fuzz_dep_.CoverTab[78216]++
														hdrFrame := it.(*headerFrame)
														var sz int64
														for _, f := range hdrFrame.hf {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:943
		_go_fuzz_dep_.CoverTab[78220]++
															if sz += int64(f.Size()); sz > int64(*t.maxSendHeaderListSize) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:944
			_go_fuzz_dep_.CoverTab[78221]++
																if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:945
				_go_fuzz_dep_.CoverTab[78223]++
																	logger.Errorf("header list size to send violates the maximum size (%d bytes) set by client", *t.maxSendHeaderListSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:946
				// _ = "end of CoverTab[78223]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:947
				_go_fuzz_dep_.CoverTab[78224]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:947
				// _ = "end of CoverTab[78224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:947
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:947
			// _ = "end of CoverTab[78221]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:947
			_go_fuzz_dep_.CoverTab[78222]++
																return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:948
			// _ = "end of CoverTab[78222]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:949
			_go_fuzz_dep_.CoverTab[78225]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:949
			// _ = "end of CoverTab[78225]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:949
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:949
		// _ = "end of CoverTab[78220]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:950
	// _ = "end of CoverTab[78216]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:950
	_go_fuzz_dep_.CoverTab[78217]++
														return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:951
	// _ = "end of CoverTab[78217]"
}

func (t *http2Server) streamContextErr(s *Stream) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:954
	_go_fuzz_dep_.CoverTab[78226]++
														select {
	case <-t.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:956
		_go_fuzz_dep_.CoverTab[78228]++
															return ErrConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:957
		// _ = "end of CoverTab[78228]"
	default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:958
		_go_fuzz_dep_.CoverTab[78229]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:958
		// _ = "end of CoverTab[78229]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:959
	// _ = "end of CoverTab[78226]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:959
	_go_fuzz_dep_.CoverTab[78227]++
														return ContextErr(s.ctx.Err())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:960
	// _ = "end of CoverTab[78227]"
}

// WriteHeader sends the header metadata md back to the client.
func (t *http2Server) WriteHeader(s *Stream, md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:964
	_go_fuzz_dep_.CoverTab[78230]++
														s.hdrMu.Lock()
														defer s.hdrMu.Unlock()
														if s.getState() == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:967
		_go_fuzz_dep_.CoverTab[78235]++
															return t.streamContextErr(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:968
		// _ = "end of CoverTab[78235]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:969
		_go_fuzz_dep_.CoverTab[78236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:969
		// _ = "end of CoverTab[78236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:969
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:969
	// _ = "end of CoverTab[78230]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:969
	_go_fuzz_dep_.CoverTab[78231]++

														if s.updateHeaderSent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:971
		_go_fuzz_dep_.CoverTab[78237]++
															return ErrIllegalHeaderWrite
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:972
		// _ = "end of CoverTab[78237]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:973
		_go_fuzz_dep_.CoverTab[78238]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:973
		// _ = "end of CoverTab[78238]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:973
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:973
	// _ = "end of CoverTab[78231]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:973
	_go_fuzz_dep_.CoverTab[78232]++

														if md.Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:975
		_go_fuzz_dep_.CoverTab[78239]++
															if s.header.Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:976
			_go_fuzz_dep_.CoverTab[78240]++
																s.header = metadata.Join(s.header, md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:977
			// _ = "end of CoverTab[78240]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:978
			_go_fuzz_dep_.CoverTab[78241]++
																s.header = md
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:979
			// _ = "end of CoverTab[78241]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:980
		// _ = "end of CoverTab[78239]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:981
		_go_fuzz_dep_.CoverTab[78242]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:981
		// _ = "end of CoverTab[78242]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:981
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:981
	// _ = "end of CoverTab[78232]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:981
	_go_fuzz_dep_.CoverTab[78233]++
														if err := t.writeHeaderLocked(s); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:982
		_go_fuzz_dep_.CoverTab[78243]++
															return status.Convert(err).Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:983
		// _ = "end of CoverTab[78243]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:984
		_go_fuzz_dep_.CoverTab[78244]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:984
		// _ = "end of CoverTab[78244]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:984
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:984
	// _ = "end of CoverTab[78233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:984
	_go_fuzz_dep_.CoverTab[78234]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:985
	// _ = "end of CoverTab[78234]"
}

func (t *http2Server) setResetPingStrikes() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:988
	_go_fuzz_dep_.CoverTab[78245]++
														atomic.StoreUint32(&t.resetPingStrikes, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:989
	// _ = "end of CoverTab[78245]"
}

func (t *http2Server) writeHeaderLocked(s *Stream) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:992
	_go_fuzz_dep_.CoverTab[78246]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:995
	headerFields := make([]hpack.HeaderField, 0, 2)
	headerFields = append(headerFields, hpack.HeaderField{Name: ":status", Value: "200"})
	headerFields = append(headerFields, hpack.HeaderField{Name: "content-type", Value: grpcutil.ContentType(s.contentSubtype)})
	if s.sendCompress != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:998
		_go_fuzz_dep_.CoverTab[78250]++
															headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-encoding", Value: s.sendCompress})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:999
		// _ = "end of CoverTab[78250]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1000
		_go_fuzz_dep_.CoverTab[78251]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1000
		// _ = "end of CoverTab[78251]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1000
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1000
	// _ = "end of CoverTab[78246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1000
	_go_fuzz_dep_.CoverTab[78247]++
														headerFields = appendHeaderFieldsFromMD(headerFields, s.header)
														success, err := t.controlBuf.executeAndPut(t.checkForHeaderListSize, &headerFrame{
		streamID:	s.id,
		hf:		headerFields,
		endStream:	false,
		onWrite:	t.setResetPingStrikes,
	})
	if !success {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1008
		_go_fuzz_dep_.CoverTab[78252]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1009
			_go_fuzz_dep_.CoverTab[78254]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1010
			// _ = "end of CoverTab[78254]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1011
			_go_fuzz_dep_.CoverTab[78255]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1011
			// _ = "end of CoverTab[78255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1011
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1011
		// _ = "end of CoverTab[78252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1011
		_go_fuzz_dep_.CoverTab[78253]++
															t.closeStream(s, true, http2.ErrCodeInternal, false)
															return ErrHeaderListSizeLimitViolation
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1013
		// _ = "end of CoverTab[78253]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1014
		_go_fuzz_dep_.CoverTab[78256]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1014
		// _ = "end of CoverTab[78256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1014
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1014
	// _ = "end of CoverTab[78247]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1014
	_go_fuzz_dep_.CoverTab[78248]++
														for _, sh := range t.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1015
		_go_fuzz_dep_.CoverTab[78257]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1018
		outHeader := &stats.OutHeader{
			Header:		s.header.Copy(),
			Compression:	s.sendCompress,
		}
															sh.HandleRPC(s.Context(), outHeader)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1022
		// _ = "end of CoverTab[78257]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1023
	// _ = "end of CoverTab[78248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1023
	_go_fuzz_dep_.CoverTab[78249]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1024
	// _ = "end of CoverTab[78249]"
}

// WriteStatus sends stream status to the client and terminates the stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1027
// There is no further I/O operations being able to perform on this stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1027
// TODO(zhaoq): Now it indicates the end of entire stream. Revisit if early
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1027
// OK is adopted.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1031
func (t *http2Server) WriteStatus(s *Stream, st *status.Status) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1031
	_go_fuzz_dep_.CoverTab[78258]++
														s.hdrMu.Lock()
														defer s.hdrMu.Unlock()

														if s.getState() == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1035
		_go_fuzz_dep_.CoverTab[78264]++
															return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1036
		// _ = "end of CoverTab[78264]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1037
		_go_fuzz_dep_.CoverTab[78265]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1037
		// _ = "end of CoverTab[78265]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1037
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1037
	// _ = "end of CoverTab[78258]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1037
	_go_fuzz_dep_.CoverTab[78259]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1041
	headerFields := make([]hpack.HeaderField, 0, 2)
	if !s.updateHeaderSent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1042
		_go_fuzz_dep_.CoverTab[78266]++
															if len(s.header) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1043
			_go_fuzz_dep_.CoverTab[78267]++
																if err := t.writeHeaderLocked(s); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1044
				_go_fuzz_dep_.CoverTab[78268]++
																	return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1045
				// _ = "end of CoverTab[78268]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1046
				_go_fuzz_dep_.CoverTab[78269]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1046
				// _ = "end of CoverTab[78269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1046
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1046
			// _ = "end of CoverTab[78267]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1047
			_go_fuzz_dep_.CoverTab[78270]++
																headerFields = append(headerFields, hpack.HeaderField{Name: ":status", Value: "200"})
																headerFields = append(headerFields, hpack.HeaderField{Name: "content-type", Value: grpcutil.ContentType(s.contentSubtype)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1049
			// _ = "end of CoverTab[78270]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1050
		// _ = "end of CoverTab[78266]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1051
		_go_fuzz_dep_.CoverTab[78271]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1051
		// _ = "end of CoverTab[78271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1051
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1051
	// _ = "end of CoverTab[78259]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1051
	_go_fuzz_dep_.CoverTab[78260]++
														headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-status", Value: strconv.Itoa(int(st.Code()))})
														headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-message", Value: encodeGrpcMessage(st.Message())})

														if p := st.Proto(); p != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1055
		_go_fuzz_dep_.CoverTab[78272]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1055
		return len(p.Details) > 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1055
		// _ = "end of CoverTab[78272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1055
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1055
		_go_fuzz_dep_.CoverTab[78273]++
															stBytes, err := proto.Marshal(p)
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1057
			_go_fuzz_dep_.CoverTab[78274]++

																logger.Errorf("transport: failed to marshal rpc status: %v, error: %v", p, err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1059
			// _ = "end of CoverTab[78274]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1060
			_go_fuzz_dep_.CoverTab[78275]++
																headerFields = append(headerFields, hpack.HeaderField{Name: "grpc-status-details-bin", Value: encodeBinHeader(stBytes)})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1061
			// _ = "end of CoverTab[78275]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1062
		// _ = "end of CoverTab[78273]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1063
		_go_fuzz_dep_.CoverTab[78276]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1063
		// _ = "end of CoverTab[78276]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1063
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1063
	// _ = "end of CoverTab[78260]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1063
	_go_fuzz_dep_.CoverTab[78261]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1066
	headerFields = appendHeaderFieldsFromMD(headerFields, s.trailer)
	trailingHeader := &headerFrame{
		streamID:	s.id,
		hf:		headerFields,
		endStream:	true,
		onWrite:	t.setResetPingStrikes,
	}

	success, err := t.controlBuf.execute(t.checkForHeaderListSize, trailingHeader)
	if !success {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1075
		_go_fuzz_dep_.CoverTab[78277]++
															if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1076
			_go_fuzz_dep_.CoverTab[78279]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1077
			// _ = "end of CoverTab[78279]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1078
			_go_fuzz_dep_.CoverTab[78280]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1078
			// _ = "end of CoverTab[78280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1078
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1078
		// _ = "end of CoverTab[78277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1078
		_go_fuzz_dep_.CoverTab[78278]++
															t.closeStream(s, true, http2.ErrCodeInternal, false)
															return ErrHeaderListSizeLimitViolation
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1080
		// _ = "end of CoverTab[78278]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1081
		_go_fuzz_dep_.CoverTab[78281]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1081
		// _ = "end of CoverTab[78281]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1081
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1081
	// _ = "end of CoverTab[78261]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1081
	_go_fuzz_dep_.CoverTab[78262]++

														rst := s.getState() == streamActive
														t.finishStream(s, rst, http2.ErrCodeNo, trailingHeader, true)
														for _, sh := range t.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1085
		_go_fuzz_dep_.CoverTab[78282]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1088
		sh.HandleRPC(s.Context(), &stats.OutTrailer{
			Trailer: s.trailer.Copy(),
		})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1090
		// _ = "end of CoverTab[78282]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1091
	// _ = "end of CoverTab[78262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1091
	_go_fuzz_dep_.CoverTab[78263]++
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1092
	// _ = "end of CoverTab[78263]"
}

// Write converts the data into HTTP2 data frame and sends it out. Non-nil error
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1095
// is returns if it fails (e.g., framing error, transport error).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1097
func (t *http2Server) Write(s *Stream, hdr []byte, data []byte, opts *Options) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1097
	_go_fuzz_dep_.CoverTab[78283]++
														if !s.isHeaderSent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1098
		_go_fuzz_dep_.CoverTab[78286]++
															if err := t.WriteHeader(s, nil); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1099
			_go_fuzz_dep_.CoverTab[78287]++
																return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1100
			// _ = "end of CoverTab[78287]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1101
			_go_fuzz_dep_.CoverTab[78288]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1101
			// _ = "end of CoverTab[78288]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1101
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1101
		// _ = "end of CoverTab[78286]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1102
		_go_fuzz_dep_.CoverTab[78289]++

															if s.getState() == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1104
			_go_fuzz_dep_.CoverTab[78290]++
																return t.streamContextErr(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1105
			// _ = "end of CoverTab[78290]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1106
			_go_fuzz_dep_.CoverTab[78291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1106
			// _ = "end of CoverTab[78291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1106
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1106
		// _ = "end of CoverTab[78289]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1107
	// _ = "end of CoverTab[78283]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1107
	_go_fuzz_dep_.CoverTab[78284]++
														df := &dataFrame{
		streamID:	s.id,
		h:		hdr,
		d:		data,
		onEachWrite:	t.setResetPingStrikes,
	}
	if err := s.wq.get(int32(len(hdr) + len(data))); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1114
		_go_fuzz_dep_.CoverTab[78292]++
															return t.streamContextErr(s)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1115
		// _ = "end of CoverTab[78292]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1116
		_go_fuzz_dep_.CoverTab[78293]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1116
		// _ = "end of CoverTab[78293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1116
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1116
	// _ = "end of CoverTab[78284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1116
	_go_fuzz_dep_.CoverTab[78285]++
														return t.controlBuf.put(df)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1117
	// _ = "end of CoverTab[78285]"
}

// keepalive running in a separate goroutine does the following:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1120
// 1. Gracefully closes an idle connection after a duration of keepalive.MaxConnectionIdle.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1120
// 2. Gracefully closes any connection after a duration of keepalive.MaxConnectionAge.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1120
// 3. Forcibly closes a connection after an additive period of keepalive.MaxConnectionAgeGrace over keepalive.MaxConnectionAge.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1120
// 4. Makes sure a connection is alive by sending pings with a frequency of keepalive.Time and closes a non-responsive connection
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1120
// after an additional duration of keepalive.Timeout.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1126
func (t *http2Server) keepalive() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1126
	_go_fuzz_dep_.CoverTab[78294]++
														p := &ping{}

														outstandingPing := false

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1132
	kpTimeoutLeft := time.Duration(0)

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1135
	prevNano := time.Now().UnixNano()

	idleTimer := time.NewTimer(t.kp.MaxConnectionIdle)
	ageTimer := time.NewTimer(t.kp.MaxConnectionAge)
	kpTimer := time.NewTimer(t.kp.Time)
	defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1140
		_go_fuzz_dep_.CoverTab[78296]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1144
		idleTimer.Stop()
															ageTimer.Stop()
															kpTimer.Stop()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1146
		// _ = "end of CoverTab[78296]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1147
	// _ = "end of CoverTab[78294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1147
	_go_fuzz_dep_.CoverTab[78295]++

														for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1149
		_go_fuzz_dep_.CoverTab[78297]++
															select {
		case <-idleTimer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1151
			_go_fuzz_dep_.CoverTab[78298]++
																t.mu.Lock()
																idle := t.idle
																if idle.IsZero() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1154
				_go_fuzz_dep_.CoverTab[78308]++
																	t.mu.Unlock()
																	idleTimer.Reset(t.kp.MaxConnectionIdle)
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1157
				// _ = "end of CoverTab[78308]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1158
				_go_fuzz_dep_.CoverTab[78309]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1158
				// _ = "end of CoverTab[78309]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1158
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1158
			// _ = "end of CoverTab[78298]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1158
			_go_fuzz_dep_.CoverTab[78299]++
																val := t.kp.MaxConnectionIdle - time.Since(idle)
																t.mu.Unlock()
																if val <= 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1161
				_go_fuzz_dep_.CoverTab[78310]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1164
				t.Drain()
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1165
				// _ = "end of CoverTab[78310]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1166
				_go_fuzz_dep_.CoverTab[78311]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1166
				// _ = "end of CoverTab[78311]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1166
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1166
			// _ = "end of CoverTab[78299]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1166
			_go_fuzz_dep_.CoverTab[78300]++
																idleTimer.Reset(val)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1167
			// _ = "end of CoverTab[78300]"
		case <-ageTimer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1168
			_go_fuzz_dep_.CoverTab[78301]++
																t.Drain()
																ageTimer.Reset(t.kp.MaxConnectionAgeGrace)
																select {
			case <-ageTimer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1172
				_go_fuzz_dep_.CoverTab[78312]++

																	if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1174
					_go_fuzz_dep_.CoverTab[78315]++
																		logger.Infof("transport: closing server transport due to maximum connection age.")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1175
					// _ = "end of CoverTab[78315]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1176
					_go_fuzz_dep_.CoverTab[78316]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1176
					// _ = "end of CoverTab[78316]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1176
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1176
				// _ = "end of CoverTab[78312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1176
				_go_fuzz_dep_.CoverTab[78313]++
																	t.controlBuf.put(closeConnection{})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1177
				// _ = "end of CoverTab[78313]"
			case <-t.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1178
				_go_fuzz_dep_.CoverTab[78314]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1178
				// _ = "end of CoverTab[78314]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1179
			// _ = "end of CoverTab[78301]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1179
			_go_fuzz_dep_.CoverTab[78302]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1180
			// _ = "end of CoverTab[78302]"
		case <-kpTimer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1181
			_go_fuzz_dep_.CoverTab[78303]++
																lastRead := atomic.LoadInt64(&t.lastRead)
																if lastRead > prevNano {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1183
				_go_fuzz_dep_.CoverTab[78317]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1187
				outstandingPing = false
																	kpTimer.Reset(time.Duration(lastRead) + t.kp.Time - time.Duration(time.Now().UnixNano()))
																	prevNano = lastRead
																	continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1190
				// _ = "end of CoverTab[78317]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1191
				_go_fuzz_dep_.CoverTab[78318]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1191
				// _ = "end of CoverTab[78318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1191
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1191
			// _ = "end of CoverTab[78303]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1191
			_go_fuzz_dep_.CoverTab[78304]++
																if outstandingPing && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1192
				_go_fuzz_dep_.CoverTab[78319]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1192
				return kpTimeoutLeft <= 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1192
				// _ = "end of CoverTab[78319]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1192
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1192
				_go_fuzz_dep_.CoverTab[78320]++
																	t.Close(fmt.Errorf("keepalive ping not acked within timeout %s", t.kp.Time))
																	return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1194
				// _ = "end of CoverTab[78320]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1195
				_go_fuzz_dep_.CoverTab[78321]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1195
				// _ = "end of CoverTab[78321]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1195
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1195
			// _ = "end of CoverTab[78304]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1195
			_go_fuzz_dep_.CoverTab[78305]++
																if !outstandingPing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1196
				_go_fuzz_dep_.CoverTab[78322]++
																	if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1197
					_go_fuzz_dep_.CoverTab[78324]++
																		atomic.AddInt64(&t.czData.kpCount, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1198
					// _ = "end of CoverTab[78324]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1199
					_go_fuzz_dep_.CoverTab[78325]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1199
					// _ = "end of CoverTab[78325]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1199
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1199
				// _ = "end of CoverTab[78322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1199
				_go_fuzz_dep_.CoverTab[78323]++
																	t.controlBuf.put(p)
																	kpTimeoutLeft = t.kp.Timeout
																	outstandingPing = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1202
				// _ = "end of CoverTab[78323]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1203
				_go_fuzz_dep_.CoverTab[78326]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1203
				// _ = "end of CoverTab[78326]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1203
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1203
			// _ = "end of CoverTab[78305]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1203
			_go_fuzz_dep_.CoverTab[78306]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1208
			sleepDuration := minTime(t.kp.Time, kpTimeoutLeft)
																kpTimeoutLeft -= sleepDuration
																kpTimer.Reset(sleepDuration)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1210
			// _ = "end of CoverTab[78306]"
		case <-t.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1211
			_go_fuzz_dep_.CoverTab[78307]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1212
			// _ = "end of CoverTab[78307]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1213
		// _ = "end of CoverTab[78297]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1214
	// _ = "end of CoverTab[78295]"
}

// Close starts shutting down the http2Server transport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1217
// TODO(zhaoq): Now the destruction is not blocked on any pending streams. This
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1217
// could cause some resource issue. Revisit this later.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1220
func (t *http2Server) Close(err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1220
	_go_fuzz_dep_.CoverTab[78327]++
														t.mu.Lock()
														if t.state == closing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1222
		_go_fuzz_dep_.CoverTab[78332]++
															t.mu.Unlock()
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1224
		// _ = "end of CoverTab[78332]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1225
		_go_fuzz_dep_.CoverTab[78333]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1225
		// _ = "end of CoverTab[78333]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1225
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1225
	// _ = "end of CoverTab[78327]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1225
	_go_fuzz_dep_.CoverTab[78328]++
														if logger.V(logLevel) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1226
		_go_fuzz_dep_.CoverTab[78334]++
															logger.Infof("transport: closing: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1227
		// _ = "end of CoverTab[78334]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1228
		_go_fuzz_dep_.CoverTab[78335]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1228
		// _ = "end of CoverTab[78335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1228
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1228
	// _ = "end of CoverTab[78328]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1228
	_go_fuzz_dep_.CoverTab[78329]++
														t.state = closing
														streams := t.activeStreams
														t.activeStreams = nil
														t.mu.Unlock()
														t.controlBuf.finish()
														close(t.done)
														if err := t.conn.Close(); err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1235
		_go_fuzz_dep_.CoverTab[78336]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1235
		return logger.V(logLevel)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1235
		// _ = "end of CoverTab[78336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1235
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1235
		_go_fuzz_dep_.CoverTab[78337]++
															logger.Infof("transport: error closing conn during Close: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1236
		// _ = "end of CoverTab[78337]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1237
		_go_fuzz_dep_.CoverTab[78338]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1237
		// _ = "end of CoverTab[78338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1237
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1237
	// _ = "end of CoverTab[78329]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1237
	_go_fuzz_dep_.CoverTab[78330]++
														channelz.RemoveEntry(t.channelzID)

														for _, s := range streams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1240
		_go_fuzz_dep_.CoverTab[78339]++
															s.cancel()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1241
		// _ = "end of CoverTab[78339]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1242
	// _ = "end of CoverTab[78330]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1242
	_go_fuzz_dep_.CoverTab[78331]++
														for _, sh := range t.stats {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1243
		_go_fuzz_dep_.CoverTab[78340]++
															connEnd := &stats.ConnEnd{}
															sh.HandleConn(t.ctx, connEnd)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1245
		// _ = "end of CoverTab[78340]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1246
	// _ = "end of CoverTab[78331]"
}

// deleteStream deletes the stream s from transport's active streams.
func (t *http2Server) deleteStream(s *Stream, eosReceived bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1250
	_go_fuzz_dep_.CoverTab[78341]++

														t.mu.Lock()
														if _, ok := t.activeStreams[s.id]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1253
		_go_fuzz_dep_.CoverTab[78343]++
															delete(t.activeStreams, s.id)
															if len(t.activeStreams) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1255
			_go_fuzz_dep_.CoverTab[78344]++
																t.idle = time.Now()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1256
			// _ = "end of CoverTab[78344]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1257
			_go_fuzz_dep_.CoverTab[78345]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1257
			// _ = "end of CoverTab[78345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1257
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1257
		// _ = "end of CoverTab[78343]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1258
		_go_fuzz_dep_.CoverTab[78346]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1258
		// _ = "end of CoverTab[78346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1258
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1258
	// _ = "end of CoverTab[78341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1258
	_go_fuzz_dep_.CoverTab[78342]++
														t.mu.Unlock()

														if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1261
		_go_fuzz_dep_.CoverTab[78347]++
															if eosReceived {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1262
			_go_fuzz_dep_.CoverTab[78348]++
																atomic.AddInt64(&t.czData.streamsSucceeded, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1263
			// _ = "end of CoverTab[78348]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1264
			_go_fuzz_dep_.CoverTab[78349]++
																atomic.AddInt64(&t.czData.streamsFailed, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1265
			// _ = "end of CoverTab[78349]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1266
		// _ = "end of CoverTab[78347]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1267
		_go_fuzz_dep_.CoverTab[78350]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1267
		// _ = "end of CoverTab[78350]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1267
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1267
	// _ = "end of CoverTab[78342]"
}

// finishStream closes the stream and puts the trailing headerFrame into controlbuf.
func (t *http2Server) finishStream(s *Stream, rst bool, rstCode http2.ErrCode, hdr *headerFrame, eosReceived bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1271
	_go_fuzz_dep_.CoverTab[78351]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1275
	s.cancel()

	oldState := s.swapState(streamDone)
	if oldState == streamDone {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1278
		_go_fuzz_dep_.CoverTab[78354]++

															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1280
		// _ = "end of CoverTab[78354]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1281
		_go_fuzz_dep_.CoverTab[78355]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1281
		// _ = "end of CoverTab[78355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1281
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1281
	// _ = "end of CoverTab[78351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1281
	_go_fuzz_dep_.CoverTab[78352]++

														hdr.cleanup = &cleanupStream{
		streamID:	s.id,
		rst:		rst,
		rstCode:	rstCode,
		onWrite: func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1287
			_go_fuzz_dep_.CoverTab[78356]++
																t.deleteStream(s, eosReceived)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1288
			// _ = "end of CoverTab[78356]"
		},
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1290
	// _ = "end of CoverTab[78352]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1290
	_go_fuzz_dep_.CoverTab[78353]++
														t.controlBuf.put(hdr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1291
	// _ = "end of CoverTab[78353]"
}

// closeStream clears the footprint of a stream when the stream is not needed any more.
func (t *http2Server) closeStream(s *Stream, rst bool, rstCode http2.ErrCode, eosReceived bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1295
	_go_fuzz_dep_.CoverTab[78357]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1299
	s.cancel()

	s.swapState(streamDone)
	t.deleteStream(s, eosReceived)

	t.controlBuf.put(&cleanupStream{
		streamID:	s.id,
		rst:		rst,
		rstCode:	rstCode,
		onWrite:	func() { _go_fuzz_dep_.CoverTab[78358]++; // _ = "end of CoverTab[78358]" },
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1309
	// _ = "end of CoverTab[78357]"
}

func (t *http2Server) RemoteAddr() net.Addr {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1312
	_go_fuzz_dep_.CoverTab[78359]++
														return t.remoteAddr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1313
	// _ = "end of CoverTab[78359]"
}

func (t *http2Server) Drain() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1316
	_go_fuzz_dep_.CoverTab[78360]++
														t.mu.Lock()
														defer t.mu.Unlock()
														if t.drainEvent != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1319
		_go_fuzz_dep_.CoverTab[78362]++
															return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1320
		// _ = "end of CoverTab[78362]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1321
		_go_fuzz_dep_.CoverTab[78363]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1321
		// _ = "end of CoverTab[78363]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1321
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1321
	// _ = "end of CoverTab[78360]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1321
	_go_fuzz_dep_.CoverTab[78361]++
														t.drainEvent = grpcsync.NewEvent()
														t.controlBuf.put(&goAway{code: http2.ErrCodeNo, debugData: []byte{}, headsUp: true})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1323
	// _ = "end of CoverTab[78361]"
}

var goAwayPing = &ping{data: [8]byte{1, 6, 1, 8, 0, 3, 3, 9}}

// Handles outgoing GoAway and returns true if loopy needs to put itself
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1328
// in draining mode.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1330
func (t *http2Server) outgoingGoAwayHandler(g *goAway) (bool, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1330
	_go_fuzz_dep_.CoverTab[78364]++
														t.maxStreamMu.Lock()
														t.mu.Lock()
														if t.state == closing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1333
		_go_fuzz_dep_.CoverTab[78370]++
															t.mu.Unlock()
															t.maxStreamMu.Unlock()

															return false, ErrConnClosing
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1337
		// _ = "end of CoverTab[78370]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1338
		_go_fuzz_dep_.CoverTab[78371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1338
		// _ = "end of CoverTab[78371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1338
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1338
	// _ = "end of CoverTab[78364]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1338
	_go_fuzz_dep_.CoverTab[78365]++
														if !g.headsUp {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1339
		_go_fuzz_dep_.CoverTab[78372]++

															t.state = draining
															sid := t.maxStreamID
															retErr := g.closeConn
															if len(t.activeStreams) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1344
			_go_fuzz_dep_.CoverTab[78376]++
																retErr = errors.New("second GOAWAY written and no active streams left to process")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1345
			// _ = "end of CoverTab[78376]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1346
			_go_fuzz_dep_.CoverTab[78377]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1346
			// _ = "end of CoverTab[78377]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1346
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1346
		// _ = "end of CoverTab[78372]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1346
		_go_fuzz_dep_.CoverTab[78373]++
															t.mu.Unlock()
															t.maxStreamMu.Unlock()
															if err := t.framer.fr.WriteGoAway(sid, g.code, g.debugData); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1349
			_go_fuzz_dep_.CoverTab[78378]++
																return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1350
			// _ = "end of CoverTab[78378]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1351
			_go_fuzz_dep_.CoverTab[78379]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1351
			// _ = "end of CoverTab[78379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1351
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1351
		// _ = "end of CoverTab[78373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1351
		_go_fuzz_dep_.CoverTab[78374]++
															if retErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1352
			_go_fuzz_dep_.CoverTab[78380]++
																return false, retErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1353
			// _ = "end of CoverTab[78380]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1354
			_go_fuzz_dep_.CoverTab[78381]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1354
			// _ = "end of CoverTab[78381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1354
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1354
		// _ = "end of CoverTab[78374]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1354
		_go_fuzz_dep_.CoverTab[78375]++
															return true, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1355
		// _ = "end of CoverTab[78375]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1356
		_go_fuzz_dep_.CoverTab[78382]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1356
		// _ = "end of CoverTab[78382]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1356
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1356
	// _ = "end of CoverTab[78365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1356
	_go_fuzz_dep_.CoverTab[78366]++
														t.mu.Unlock()
														t.maxStreamMu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1365
	if err := t.framer.fr.WriteGoAway(math.MaxUint32, http2.ErrCodeNo, []byte{}); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1365
		_go_fuzz_dep_.CoverTab[78383]++
															return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1366
		// _ = "end of CoverTab[78383]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1367
		_go_fuzz_dep_.CoverTab[78384]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1367
		// _ = "end of CoverTab[78384]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1367
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1367
	// _ = "end of CoverTab[78366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1367
	_go_fuzz_dep_.CoverTab[78367]++
														if err := t.framer.fr.WritePing(false, goAwayPing.data); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1368
		_go_fuzz_dep_.CoverTab[78385]++
															return false, err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1369
		// _ = "end of CoverTab[78385]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
		_go_fuzz_dep_.CoverTab[78386]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
		// _ = "end of CoverTab[78386]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
	// _ = "end of CoverTab[78367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
	_go_fuzz_dep_.CoverTab[78368]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
	_curRoutineNum84_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1370
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum84_)
														go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1371
		_go_fuzz_dep_.CoverTab[78387]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1371
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1371
			_go_fuzz_dep_.CoverTab[78389]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1371
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum84_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1371
			// _ = "end of CoverTab[78389]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1371
		}()
															timer := time.NewTimer(time.Minute)
															defer timer.Stop()
															select {
		case <-t.drainEvent.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1375
			_go_fuzz_dep_.CoverTab[78390]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1375
			// _ = "end of CoverTab[78390]"
		case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1376
			_go_fuzz_dep_.CoverTab[78391]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1376
			// _ = "end of CoverTab[78391]"
		case <-t.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1377
			_go_fuzz_dep_.CoverTab[78392]++
																return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1378
			// _ = "end of CoverTab[78392]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1379
		// _ = "end of CoverTab[78387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1379
		_go_fuzz_dep_.CoverTab[78388]++
															t.controlBuf.put(&goAway{code: g.code, debugData: g.debugData})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1380
		// _ = "end of CoverTab[78388]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1381
	// _ = "end of CoverTab[78368]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1381
	_go_fuzz_dep_.CoverTab[78369]++
														return false, nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1382
	// _ = "end of CoverTab[78369]"
}

func (t *http2Server) ChannelzMetric() *channelz.SocketInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1385
	_go_fuzz_dep_.CoverTab[78393]++
														s := channelz.SocketInternalMetric{
		StreamsStarted:				atomic.LoadInt64(&t.czData.streamsStarted),
		StreamsSucceeded:			atomic.LoadInt64(&t.czData.streamsSucceeded),
		StreamsFailed:				atomic.LoadInt64(&t.czData.streamsFailed),
		MessagesSent:				atomic.LoadInt64(&t.czData.msgSent),
		MessagesReceived:			atomic.LoadInt64(&t.czData.msgRecv),
		KeepAlivesSent:				atomic.LoadInt64(&t.czData.kpCount),
		LastRemoteStreamCreatedTimestamp:	time.Unix(0, atomic.LoadInt64(&t.czData.lastStreamCreatedTime)),
		LastMessageSentTimestamp:		time.Unix(0, atomic.LoadInt64(&t.czData.lastMsgSentTime)),
		LastMessageReceivedTimestamp:		time.Unix(0, atomic.LoadInt64(&t.czData.lastMsgRecvTime)),
		LocalFlowControlWindow:			int64(t.fc.getSize()),
		SocketOptions:				channelz.GetSocketOption(t.conn),
		LocalAddr:				t.localAddr,
		RemoteAddr:				t.remoteAddr,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1401
	}
	if au, ok := t.authInfo.(credentials.ChannelzSecurityInfo); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1402
		_go_fuzz_dep_.CoverTab[78395]++
															s.Security = au.GetSecurityValue()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1403
		// _ = "end of CoverTab[78395]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1404
		_go_fuzz_dep_.CoverTab[78396]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1404
		// _ = "end of CoverTab[78396]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1404
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1404
	// _ = "end of CoverTab[78393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1404
	_go_fuzz_dep_.CoverTab[78394]++
														s.RemoteFlowControlWindow = t.getOutFlowWindow()
														return &s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1406
	// _ = "end of CoverTab[78394]"
}

func (t *http2Server) IncrMsgSent() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1409
	_go_fuzz_dep_.CoverTab[78397]++
														atomic.AddInt64(&t.czData.msgSent, 1)
														atomic.StoreInt64(&t.czData.lastMsgSentTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1411
	// _ = "end of CoverTab[78397]"
}

func (t *http2Server) IncrMsgRecv() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1414
	_go_fuzz_dep_.CoverTab[78398]++
														atomic.AddInt64(&t.czData.msgRecv, 1)
														atomic.StoreInt64(&t.czData.lastMsgRecvTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1416
	// _ = "end of CoverTab[78398]"
}

func (t *http2Server) getOutFlowWindow() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1419
	_go_fuzz_dep_.CoverTab[78399]++
														resp := make(chan uint32, 1)
														timer := time.NewTimer(time.Second)
														defer timer.Stop()
														t.controlBuf.put(&outFlowControlSizeRequest{resp})
														select {
	case sz := <-resp:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1425
		_go_fuzz_dep_.CoverTab[78400]++
															return int64(sz)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1426
		// _ = "end of CoverTab[78400]"
	case <-t.done:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1427
		_go_fuzz_dep_.CoverTab[78401]++
															return -1
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1428
		// _ = "end of CoverTab[78401]"
	case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1429
		_go_fuzz_dep_.CoverTab[78402]++
															return -2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1430
		// _ = "end of CoverTab[78402]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1431
	// _ = "end of CoverTab[78399]"
}

func (t *http2Server) getPeer() *peer.Peer {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1434
	_go_fuzz_dep_.CoverTab[78403]++
														return &peer.Peer{
		Addr:		t.remoteAddr,
		AuthInfo:	t.authInfo,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1438
	// _ = "end of CoverTab[78403]"
}

func getJitter(v time.Duration) time.Duration {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1441
	_go_fuzz_dep_.CoverTab[78404]++
														if v == infinity {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1442
		_go_fuzz_dep_.CoverTab[78406]++
															return 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1443
		// _ = "end of CoverTab[78406]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1444
		_go_fuzz_dep_.CoverTab[78407]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1444
		// _ = "end of CoverTab[78407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1444
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1444
	// _ = "end of CoverTab[78404]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1444
	_go_fuzz_dep_.CoverTab[78405]++

														r := int64(v / 10)
														j := grpcrand.Int63n(2*r) - r
														return time.Duration(j)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1448
	// _ = "end of CoverTab[78405]"
}

type connectionKey struct{}

// GetConnection gets the connection from the context.
func GetConnection(ctx context.Context) net.Conn {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1454
	_go_fuzz_dep_.CoverTab[78408]++
														conn, _ := ctx.Value(connectionKey{}).(net.Conn)
														return conn
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1456
	// _ = "end of CoverTab[78408]"
}

// SetConnection adds the connection to the context to be able to get
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1459
// information about the destination ip and port for an incoming RPC. This also
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1459
// allows any unary or streaming interceptors to see the connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1462
func setConnection(ctx context.Context, conn net.Conn) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1462
	_go_fuzz_dep_.CoverTab[78409]++
														return context.WithValue(ctx, connectionKey{}, conn)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1463
	// _ = "end of CoverTab[78409]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1464
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/transport/http2_server.go:1464
var _ = _go_fuzz_dep_.CoverTab
