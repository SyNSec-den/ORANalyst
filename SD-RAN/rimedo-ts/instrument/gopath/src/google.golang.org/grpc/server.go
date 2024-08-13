//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
package grpc

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:19
)

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/trace"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal"
	"google.golang.org/grpc/internal/binarylog"
	"google.golang.org/grpc/internal/channelz"
	"google.golang.org/grpc/internal/grpcrand"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/grpcutil"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

const (
	defaultServerMaxReceiveMessageSize	= 1024 * 1024 * 4
	defaultServerMaxSendMessageSize		= math.MaxInt32

	// Server transports are tracked in a map which is keyed on listener
	// address. For regular gRPC traffic, connections are accepted in Serve()
	// through a call to Accept(), and we use the actual listener address as key
	// when we add it to the map. But for connections received through
	// ServeHTTP(), we do not have a listener and hence use this dummy value.
	listenerAddressForServeHTTP	= "listenerAddressForServeHTTP"
)

func init() {
	internal.GetServerCredentials = func(srv *Server) credentials.TransportCredentials {
		return srv.opts.creds
	}
	internal.DrainServerTransports = func(srv *Server, addr string) {
		srv.drainServerTransports(addr)
	}
	internal.AddGlobalServerOptions = func(opt ...ServerOption) {
		globalServerOptions = append(globalServerOptions, opt...)
	}
	internal.ClearGlobalServerOptions = func() {
		globalServerOptions = nil
	}
	internal.BinaryLogger = binaryLogger
	internal.JoinServerOptions = newJoinServerOption
}

var statusOK = status.New(codes.OK, "")
var logger = grpclog.Component("core")

type methodHandler func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor UnaryServerInterceptor) (interface{}, error)

// MethodDesc represents an RPC service's method specification.
type MethodDesc struct {
	MethodName	string
	Handler		methodHandler
}

// ServiceDesc represents an RPC service's specification.
type ServiceDesc struct {
	ServiceName	string
	// The pointer to the service interface. Used to check whether the user
	// provided implementation satisfies the interface requirements.
	HandlerType	interface{}
	Methods		[]MethodDesc
	Streams		[]StreamDesc
	Metadata	interface{}
}

// serviceInfo wraps information about a service. It is very similar to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:109
// ServiceDesc and is constructed from it for internal purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:111
type serviceInfo struct {
	// Contains the implementation for the methods in this service.
	serviceImpl	interface{}
	methods		map[string]*MethodDesc
	streams		map[string]*StreamDesc
	mdata		interface{}
}

type serverWorkerData struct {
	st	transport.ServerTransport
	wg	*sync.WaitGroup
	stream	*transport.Stream
}

// Server is a gRPC server to serve RPC requests.
type Server struct {
	opts	serverOptions

	mu	sync.Mutex	// guards following
	lis	map[net.Listener]bool
	// conns contains all active server transports. It is a map keyed on a
	// listener address with the value being the set of active transports
	// belonging to that listener.
	conns		map[string]map[transport.ServerTransport]bool
	serve		bool
	drain		bool
	cv		*sync.Cond		// signaled when connections close for GracefulStop
	services	map[string]*serviceInfo	// service name -> service info
	events		trace.EventLog

	quit			*grpcsync.Event
	done			*grpcsync.Event
	channelzRemoveOnce	sync.Once
	serveWG			sync.WaitGroup	// counts active Serve goroutines for GracefulStop

	channelzID	*channelz.Identifier
	czData		*channelzData

	serverWorkerChannels	[]chan *serverWorkerData
}

type serverOptions struct {
	creds			credentials.TransportCredentials
	codec			baseCodec
	cp			Compressor
	dc			Decompressor
	unaryInt		UnaryServerInterceptor
	streamInt		StreamServerInterceptor
	chainUnaryInts		[]UnaryServerInterceptor
	chainStreamInts		[]StreamServerInterceptor
	binaryLogger		binarylog.Logger
	inTapHandle		tap.ServerInHandle
	statsHandlers		[]stats.Handler
	maxConcurrentStreams	uint32
	maxReceiveMessageSize	int
	maxSendMessageSize	int
	unknownStreamDesc	*StreamDesc
	keepaliveParams		keepalive.ServerParameters
	keepalivePolicy		keepalive.EnforcementPolicy
	initialWindowSize	int32
	initialConnWindowSize	int32
	writeBufferSize		int
	readBufferSize		int
	connectionTimeout	time.Duration
	maxHeaderListSize	*uint32
	headerTableSize		*uint32
	numServerWorkers	uint32
}

var defaultServerOptions = serverOptions{
	maxReceiveMessageSize:	defaultServerMaxReceiveMessageSize,
	maxSendMessageSize:	defaultServerMaxSendMessageSize,
	connectionTimeout:	120 * time.Second,
	writeBufferSize:	defaultWriteBufSize,
	readBufferSize:		defaultReadBufSize,
}
var globalServerOptions []ServerOption

// A ServerOption sets options such as credentials, codec and keepalive parameters, etc.
type ServerOption interface {
	apply(*serverOptions)
}

// EmptyServerOption does not alter the server configuration. It can be embedded
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:194
// in another structure to build custom server options.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:194
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:194
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:194
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:194
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:194
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:201
type EmptyServerOption struct{}

func (EmptyServerOption) apply(*serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:203
	_go_fuzz_dep_.CoverTab[79883]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:203
	// _ = "end of CoverTab[79883]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:203
}

// funcServerOption wraps a function that modifies serverOptions into an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:205
// implementation of the ServerOption interface.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:207
type funcServerOption struct {
	f func(*serverOptions)
}

func (fdo *funcServerOption) apply(do *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:211
	_go_fuzz_dep_.CoverTab[79884]++
										fdo.f(do)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:212
	// _ = "end of CoverTab[79884]"
}

func newFuncServerOption(f func(*serverOptions)) *funcServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:215
	_go_fuzz_dep_.CoverTab[79885]++
										return &funcServerOption{
		f: f,
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:218
	// _ = "end of CoverTab[79885]"
}

// joinServerOption provides a way to combine arbitrary number of server
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:221
// options into one.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:223
type joinServerOption struct {
	opts []ServerOption
}

func (mdo *joinServerOption) apply(do *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:227
	_go_fuzz_dep_.CoverTab[79886]++
										for _, opt := range mdo.opts {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:228
		_go_fuzz_dep_.CoverTab[79887]++
											opt.apply(do)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:229
		// _ = "end of CoverTab[79887]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:230
	// _ = "end of CoverTab[79886]"
}

func newJoinServerOption(opts ...ServerOption) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:233
	_go_fuzz_dep_.CoverTab[79888]++
										return &joinServerOption{opts: opts}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:234
	// _ = "end of CoverTab[79888]"
}

// WriteBufferSize determines how much data can be batched before doing a write
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:237
// on the wire. The corresponding memory allocation for this buffer will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:237
// twice the size to keep syscalls low. The default value for this buffer is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:237
// 32KB. Zero or negative values will disable the write buffer such that each
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:237
// write will be on underlying connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:237
// Note: A Send call may not directly translate to a write.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:243
func WriteBufferSize(s int) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:243
	_go_fuzz_dep_.CoverTab[79889]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:244
		_go_fuzz_dep_.CoverTab[79890]++
											o.writeBufferSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:245
		// _ = "end of CoverTab[79890]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:246
	// _ = "end of CoverTab[79889]"
}

// ReadBufferSize lets you set the size of read buffer, this determines how much
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:249
// data can be read at most for one read syscall. The default value for this
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:249
// buffer is 32KB. Zero or negative values will disable read buffer for a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:249
// connection so data framer can access the underlying conn directly.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:253
func ReadBufferSize(s int) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:253
	_go_fuzz_dep_.CoverTab[79891]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:254
		_go_fuzz_dep_.CoverTab[79892]++
											o.readBufferSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:255
		// _ = "end of CoverTab[79892]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:256
	// _ = "end of CoverTab[79891]"
}

// InitialWindowSize returns a ServerOption that sets window size for stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:259
// The lower bound for window size is 64K and any value smaller than that will be ignored.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:261
func InitialWindowSize(s int32) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:261
	_go_fuzz_dep_.CoverTab[79893]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:262
		_go_fuzz_dep_.CoverTab[79894]++
											o.initialWindowSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:263
		// _ = "end of CoverTab[79894]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:264
	// _ = "end of CoverTab[79893]"
}

// InitialConnWindowSize returns a ServerOption that sets window size for a connection.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:267
// The lower bound for window size is 64K and any value smaller than that will be ignored.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:269
func InitialConnWindowSize(s int32) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:269
	_go_fuzz_dep_.CoverTab[79895]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:270
		_go_fuzz_dep_.CoverTab[79896]++
											o.initialConnWindowSize = s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:271
		// _ = "end of CoverTab[79896]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:272
	// _ = "end of CoverTab[79895]"
}

// KeepaliveParams returns a ServerOption that sets keepalive and max-age parameters for the server.
func KeepaliveParams(kp keepalive.ServerParameters) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:276
	_go_fuzz_dep_.CoverTab[79897]++
										if kp.Time > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:277
		_go_fuzz_dep_.CoverTab[79899]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:277
		return kp.Time < time.Second
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:277
		// _ = "end of CoverTab[79899]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:277
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:277
		_go_fuzz_dep_.CoverTab[79900]++
											logger.Warning("Adjusting keepalive ping interval to minimum period of 1s")
											kp.Time = time.Second
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:279
		// _ = "end of CoverTab[79900]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:280
		_go_fuzz_dep_.CoverTab[79901]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:280
		// _ = "end of CoverTab[79901]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:280
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:280
	// _ = "end of CoverTab[79897]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:280
	_go_fuzz_dep_.CoverTab[79898]++

										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:282
		_go_fuzz_dep_.CoverTab[79902]++
											o.keepaliveParams = kp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:283
		// _ = "end of CoverTab[79902]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:284
	// _ = "end of CoverTab[79898]"
}

// KeepaliveEnforcementPolicy returns a ServerOption that sets keepalive enforcement policy for the server.
func KeepaliveEnforcementPolicy(kep keepalive.EnforcementPolicy) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:288
	_go_fuzz_dep_.CoverTab[79903]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:289
		_go_fuzz_dep_.CoverTab[79904]++
											o.keepalivePolicy = kep
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:290
		// _ = "end of CoverTab[79904]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:291
	// _ = "end of CoverTab[79903]"
}

// CustomCodec returns a ServerOption that sets a codec for message marshaling and unmarshaling.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
// This will override any lookups by content-subtype for Codecs registered with RegisterCodec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
// Deprecated: register codecs using encoding.RegisterCodec. The server will
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
// automatically use registered codecs based on the incoming requests' headers.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
// See also
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
// https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md#using-a-codec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:294
// Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:303
func CustomCodec(codec Codec) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:303
	_go_fuzz_dep_.CoverTab[79905]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:304
		_go_fuzz_dep_.CoverTab[79906]++
											o.codec = codec
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:305
		// _ = "end of CoverTab[79906]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:306
	// _ = "end of CoverTab[79905]"
}

// ForceServerCodec returns a ServerOption that sets a codec for message
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// marshaling and unmarshaling.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// This will override any lookups by content-subtype for Codecs registered
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// with RegisterCodec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// See Content-Type on
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// more details. Also see the documentation on RegisterCodec and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// CallContentSubtype for more details on the interaction between encoding.Codec
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// and content-subtype.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// This function is provided for advanced users; prefer to register codecs
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// using encoding.RegisterCodec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// The server will automatically use registered codecs based on the incoming
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// requests' headers. See also
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md#using-a-codec.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:309
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:332
func ForceServerCodec(codec encoding.Codec) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:332
	_go_fuzz_dep_.CoverTab[79907]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:333
		_go_fuzz_dep_.CoverTab[79908]++
											o.codec = codec
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:334
		// _ = "end of CoverTab[79908]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:335
	// _ = "end of CoverTab[79907]"
}

// RPCCompressor returns a ServerOption that sets a compressor for outbound
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
// messages.  For backward compatibility, all outbound messages will be sent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
// using this compressor, regardless of incoming message compression.  By
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
// default, server messages will be sent using the same compressor with which
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
// request messages were sent.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
// Deprecated: use encoding.RegisterCompressor instead. Will be supported
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:338
// throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:346
func RPCCompressor(cp Compressor) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:346
	_go_fuzz_dep_.CoverTab[79909]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:347
		_go_fuzz_dep_.CoverTab[79910]++
											o.cp = cp
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:348
		// _ = "end of CoverTab[79910]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:349
	// _ = "end of CoverTab[79909]"
}

// RPCDecompressor returns a ServerOption that sets a decompressor for inbound
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:352
// messages.  It has higher priority than decompressors registered via
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:352
// encoding.RegisterCompressor.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:352
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:352
// Deprecated: use encoding.RegisterCompressor instead. Will be supported
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:352
// throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:358
func RPCDecompressor(dc Decompressor) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:358
	_go_fuzz_dep_.CoverTab[79911]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:359
		_go_fuzz_dep_.CoverTab[79912]++
											o.dc = dc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:360
		// _ = "end of CoverTab[79912]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:361
	// _ = "end of CoverTab[79911]"
}

// MaxMsgSize returns a ServerOption to set the max message size in bytes the server can receive.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:364
// If this is not set, gRPC uses the default limit.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:364
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:364
// Deprecated: use MaxRecvMsgSize instead. Will be supported throughout 1.x.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:368
func MaxMsgSize(m int) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:368
	_go_fuzz_dep_.CoverTab[79913]++
										return MaxRecvMsgSize(m)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:369
	// _ = "end of CoverTab[79913]"
}

// MaxRecvMsgSize returns a ServerOption to set the max message size in bytes the server can receive.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:372
// If this is not set, gRPC uses the default 4MB.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:374
func MaxRecvMsgSize(m int) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:374
	_go_fuzz_dep_.CoverTab[79914]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:375
		_go_fuzz_dep_.CoverTab[79915]++
											o.maxReceiveMessageSize = m
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:376
		// _ = "end of CoverTab[79915]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:377
	// _ = "end of CoverTab[79914]"
}

// MaxSendMsgSize returns a ServerOption to set the max message size in bytes the server can send.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:380
// If this is not set, gRPC uses the default `math.MaxInt32`.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:382
func MaxSendMsgSize(m int) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:382
	_go_fuzz_dep_.CoverTab[79916]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:383
		_go_fuzz_dep_.CoverTab[79917]++
											o.maxSendMessageSize = m
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:384
		// _ = "end of CoverTab[79917]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:385
	// _ = "end of CoverTab[79916]"
}

// MaxConcurrentStreams returns a ServerOption that will apply a limit on the number
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:388
// of concurrent streams to each ServerTransport.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:390
func MaxConcurrentStreams(n uint32) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:390
	_go_fuzz_dep_.CoverTab[79918]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:391
		_go_fuzz_dep_.CoverTab[79919]++
											o.maxConcurrentStreams = n
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:392
		// _ = "end of CoverTab[79919]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:393
	// _ = "end of CoverTab[79918]"
}

// Creds returns a ServerOption that sets credentials for server connections.
func Creds(c credentials.TransportCredentials) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:397
	_go_fuzz_dep_.CoverTab[79920]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:398
		_go_fuzz_dep_.CoverTab[79921]++
											o.creds = c
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:399
		// _ = "end of CoverTab[79921]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:400
	// _ = "end of CoverTab[79920]"
}

// UnaryInterceptor returns a ServerOption that sets the UnaryServerInterceptor for the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:403
// server. Only one unary interceptor can be installed. The construction of multiple
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:403
// interceptors (e.g., chaining) can be implemented at the caller.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:406
func UnaryInterceptor(i UnaryServerInterceptor) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:406
	_go_fuzz_dep_.CoverTab[79922]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:407
		_go_fuzz_dep_.CoverTab[79923]++
											if o.unaryInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:408
			_go_fuzz_dep_.CoverTab[79925]++
												panic("The unary server interceptor was already set and may not be reset.")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:409
			// _ = "end of CoverTab[79925]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:410
			_go_fuzz_dep_.CoverTab[79926]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:410
			// _ = "end of CoverTab[79926]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:410
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:410
		// _ = "end of CoverTab[79923]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:410
		_go_fuzz_dep_.CoverTab[79924]++
											o.unaryInt = i
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:411
		// _ = "end of CoverTab[79924]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:412
	// _ = "end of CoverTab[79922]"
}

// ChainUnaryInterceptor returns a ServerOption that specifies the chained interceptor
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:415
// for unary RPCs. The first interceptor will be the outer most,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:415
// while the last interceptor will be the inner most wrapper around the real call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:415
// All unary interceptors added by this method will be chained.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:419
func ChainUnaryInterceptor(interceptors ...UnaryServerInterceptor) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:419
	_go_fuzz_dep_.CoverTab[79927]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:420
		_go_fuzz_dep_.CoverTab[79928]++
											o.chainUnaryInts = append(o.chainUnaryInts, interceptors...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:421
		// _ = "end of CoverTab[79928]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:422
	// _ = "end of CoverTab[79927]"
}

// StreamInterceptor returns a ServerOption that sets the StreamServerInterceptor for the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:425
// server. Only one stream interceptor can be installed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:427
func StreamInterceptor(i StreamServerInterceptor) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:427
	_go_fuzz_dep_.CoverTab[79929]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:428
		_go_fuzz_dep_.CoverTab[79930]++
											if o.streamInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:429
			_go_fuzz_dep_.CoverTab[79932]++
												panic("The stream server interceptor was already set and may not be reset.")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:430
			// _ = "end of CoverTab[79932]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:431
			_go_fuzz_dep_.CoverTab[79933]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:431
			// _ = "end of CoverTab[79933]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:431
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:431
		// _ = "end of CoverTab[79930]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:431
		_go_fuzz_dep_.CoverTab[79931]++
											o.streamInt = i
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:432
		// _ = "end of CoverTab[79931]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:433
	// _ = "end of CoverTab[79929]"
}

// ChainStreamInterceptor returns a ServerOption that specifies the chained interceptor
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:436
// for streaming RPCs. The first interceptor will be the outer most,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:436
// while the last interceptor will be the inner most wrapper around the real call.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:436
// All stream interceptors added by this method will be chained.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:440
func ChainStreamInterceptor(interceptors ...StreamServerInterceptor) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:440
	_go_fuzz_dep_.CoverTab[79934]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:441
		_go_fuzz_dep_.CoverTab[79935]++
											o.chainStreamInts = append(o.chainStreamInts, interceptors...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:442
		// _ = "end of CoverTab[79935]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:443
	// _ = "end of CoverTab[79934]"
}

// InTapHandle returns a ServerOption that sets the tap handle for all the server
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:446
// transport to be created. Only one can be installed.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:446
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:446
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:446
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:446
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:446
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:453
func InTapHandle(h tap.ServerInHandle) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:453
	_go_fuzz_dep_.CoverTab[79936]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:454
		_go_fuzz_dep_.CoverTab[79937]++
											if o.inTapHandle != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:455
			_go_fuzz_dep_.CoverTab[79939]++
												panic("The tap handle was already set and may not be reset.")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:456
			// _ = "end of CoverTab[79939]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:457
			_go_fuzz_dep_.CoverTab[79940]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:457
			// _ = "end of CoverTab[79940]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:457
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:457
		// _ = "end of CoverTab[79937]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:457
		_go_fuzz_dep_.CoverTab[79938]++
											o.inTapHandle = h
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:458
		// _ = "end of CoverTab[79938]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:459
	// _ = "end of CoverTab[79936]"
}

// StatsHandler returns a ServerOption that sets the stats handler for the server.
func StatsHandler(h stats.Handler) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:463
	_go_fuzz_dep_.CoverTab[79941]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:464
		_go_fuzz_dep_.CoverTab[79942]++
											if h == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:465
			_go_fuzz_dep_.CoverTab[79944]++
												logger.Error("ignoring nil parameter in grpc.StatsHandler ServerOption")

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:469
			return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:469
			// _ = "end of CoverTab[79944]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:470
			_go_fuzz_dep_.CoverTab[79945]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:470
			// _ = "end of CoverTab[79945]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:470
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:470
		// _ = "end of CoverTab[79942]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:470
		_go_fuzz_dep_.CoverTab[79943]++
											o.statsHandlers = append(o.statsHandlers, h)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:471
		// _ = "end of CoverTab[79943]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:472
	// _ = "end of CoverTab[79941]"
}

// binaryLogger returns a ServerOption that can set the binary logger for the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:475
// server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:477
func binaryLogger(bl binarylog.Logger) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:477
	_go_fuzz_dep_.CoverTab[79946]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:478
		_go_fuzz_dep_.CoverTab[79947]++
											o.binaryLogger = bl
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:479
		// _ = "end of CoverTab[79947]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:480
	// _ = "end of CoverTab[79946]"
}

// UnknownServiceHandler returns a ServerOption that allows for adding a custom
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:483
// unknown service handler. The provided method is a bidi-streaming RPC service
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:483
// handler that will be invoked instead of returning the "unimplemented" gRPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:483
// error whenever a request is received for an unregistered service or method.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:483
// The handling function and stream interceptor (if set) have full access to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:483
// the ServerStream, including its Context.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:489
func UnknownServiceHandler(streamHandler StreamHandler) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:489
	_go_fuzz_dep_.CoverTab[79948]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:490
		_go_fuzz_dep_.CoverTab[79949]++
											o.unknownStreamDesc = &StreamDesc{
			StreamName:	"unknown_service_handler",
			Handler:	streamHandler,

			ClientStreams:	true,
			ServerStreams:	true,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:497
		// _ = "end of CoverTab[79949]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:498
	// _ = "end of CoverTab[79948]"
}

// ConnectionTimeout returns a ServerOption that sets the timeout for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
// connection establishment (up to and including HTTP/2 handshaking) for all
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
// new connections.  If this is not set, the default is 120 seconds.  A zero or
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
// negative value will result in an immediate timeout.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:501
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:510
func ConnectionTimeout(d time.Duration) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:510
	_go_fuzz_dep_.CoverTab[79950]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:511
		_go_fuzz_dep_.CoverTab[79951]++
											o.connectionTimeout = d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:512
		// _ = "end of CoverTab[79951]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:513
	// _ = "end of CoverTab[79950]"
}

// MaxHeaderListSize returns a ServerOption that sets the max (uncompressed) size
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:516
// of header list that the server is prepared to accept.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:518
func MaxHeaderListSize(s uint32) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:518
	_go_fuzz_dep_.CoverTab[79952]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:519
		_go_fuzz_dep_.CoverTab[79953]++
											o.maxHeaderListSize = &s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:520
		// _ = "end of CoverTab[79953]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:521
	// _ = "end of CoverTab[79952]"
}

// HeaderTableSize returns a ServerOption that sets the size of dynamic
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:524
// header table for stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:524
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:524
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:524
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:524
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:524
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:531
func HeaderTableSize(s uint32) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:531
	_go_fuzz_dep_.CoverTab[79954]++
										return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:532
		_go_fuzz_dep_.CoverTab[79955]++
											o.headerTableSize = &s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:533
		// _ = "end of CoverTab[79955]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:534
	// _ = "end of CoverTab[79954]"
}

// NumStreamWorkers returns a ServerOption that sets the number of worker
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
// goroutines that should be used to process incoming streams. Setting this to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
// zero (default) will disable workers and spawn a new goroutine for each
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
// stream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:537
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:546
func NumStreamWorkers(numServerWorkers uint32) ServerOption {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:546
	_go_fuzz_dep_.CoverTab[79956]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:551
	return newFuncServerOption(func(o *serverOptions) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:551
		_go_fuzz_dep_.CoverTab[79957]++
											o.numServerWorkers = numServerWorkers
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:552
		// _ = "end of CoverTab[79957]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:553
	// _ = "end of CoverTab[79956]"
}

// serverWorkerResetThreshold defines how often the stack must be reset. Every
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:556
// N requests, by spawning a new goroutine in its place, a worker can reset its
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:556
// stack so that large stacks don't live in memory forever. 2^16 should allow
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:556
// each goroutine stack to live for at least a few seconds in a typical
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:556
// workload (assuming a QPS of a few thousand requests/sec).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:561
const serverWorkerResetThreshold = 1 << 16

// serverWorkers blocks on a *transport.Stream channel forever and waits for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:563
// data to be fed by serveStreams. This allows different requests to be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:563
// processed by the same goroutine, removing the need for expensive stack
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:563
// re-allocations (see the runtime.morestack problem [1]).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:563
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:563
// [1] https://github.com/golang/go/issues/18138
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:569
func (s *Server) serverWorker(ch chan *serverWorkerData) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:569
	_go_fuzz_dep_.CoverTab[79958]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:572
	threshold := serverWorkerResetThreshold + grpcrand.Intn(serverWorkerResetThreshold)
	for completed := 0; completed < threshold; completed++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:573
		_go_fuzz_dep_.CoverTab[79960]++
											data, ok := <-ch
											if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:575
			_go_fuzz_dep_.CoverTab[79962]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:576
			// _ = "end of CoverTab[79962]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:577
			_go_fuzz_dep_.CoverTab[79963]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:577
			// _ = "end of CoverTab[79963]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:577
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:577
		// _ = "end of CoverTab[79960]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:577
		_go_fuzz_dep_.CoverTab[79961]++
											s.handleStream(data.st, data.stream, s.traceInfo(data.st, data.stream))
											data.wg.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:579
		// _ = "end of CoverTab[79961]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:580
	// _ = "end of CoverTab[79958]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:580
	_go_fuzz_dep_.CoverTab[79959]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:580
	_curRoutineNum92_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:580
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum92_)
										go s.serverWorker(ch)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:581
	// _ = "end of CoverTab[79959]"
}

// initServerWorkers creates worker goroutines and channels to process incoming
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:584
// connections to reduce the time spent overall on runtime.morestack.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:586
func (s *Server) initServerWorkers() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:586
	_go_fuzz_dep_.CoverTab[79964]++
										s.serverWorkerChannels = make([]chan *serverWorkerData, s.opts.numServerWorkers)
										for i := uint32(0); i < s.opts.numServerWorkers; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:588
		_go_fuzz_dep_.CoverTab[79965]++
											s.serverWorkerChannels[i] = make(chan *serverWorkerData)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:589
		_curRoutineNum93_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:589
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum93_)
											go s.serverWorker(s.serverWorkerChannels[i])
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:590
		// _ = "end of CoverTab[79965]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:591
	// _ = "end of CoverTab[79964]"
}

func (s *Server) stopServerWorkers() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:594
	_go_fuzz_dep_.CoverTab[79966]++
										for i := uint32(0); i < s.opts.numServerWorkers; i++ {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:595
		_go_fuzz_dep_.CoverTab[79967]++
											close(s.serverWorkerChannels[i])
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:596
		// _ = "end of CoverTab[79967]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:597
	// _ = "end of CoverTab[79966]"
}

// NewServer creates a gRPC server which has no service registered and has not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:600
// started to accept requests yet.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:602
func NewServer(opt ...ServerOption) *Server {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:602
	_go_fuzz_dep_.CoverTab[79968]++
										opts := defaultServerOptions
										for _, o := range globalServerOptions {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:604
		_go_fuzz_dep_.CoverTab[79973]++
											o.apply(&opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:605
		// _ = "end of CoverTab[79973]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:606
	// _ = "end of CoverTab[79968]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:606
	_go_fuzz_dep_.CoverTab[79969]++
										for _, o := range opt {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:607
		_go_fuzz_dep_.CoverTab[79974]++
											o.apply(&opts)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:608
		// _ = "end of CoverTab[79974]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:609
	// _ = "end of CoverTab[79969]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:609
	_go_fuzz_dep_.CoverTab[79970]++
										s := &Server{
		lis:		make(map[net.Listener]bool),
		opts:		opts,
		conns:		make(map[string]map[transport.ServerTransport]bool),
		services:	make(map[string]*serviceInfo),
		quit:		grpcsync.NewEvent(),
		done:		grpcsync.NewEvent(),
		czData:		new(channelzData),
	}
	chainUnaryServerInterceptors(s)
	chainStreamServerInterceptors(s)
	s.cv = sync.NewCond(&s.mu)
	if EnableTracing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:622
		_go_fuzz_dep_.CoverTab[79975]++
											_, file, line, _ := runtime.Caller(1)
											s.events = trace.NewEventLog("grpc.Server", fmt.Sprintf("%s:%d", file, line))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:624
		// _ = "end of CoverTab[79975]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:625
		_go_fuzz_dep_.CoverTab[79976]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:625
		// _ = "end of CoverTab[79976]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:625
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:625
	// _ = "end of CoverTab[79970]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:625
	_go_fuzz_dep_.CoverTab[79971]++

										if s.opts.numServerWorkers > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:627
		_go_fuzz_dep_.CoverTab[79977]++
											s.initServerWorkers()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:628
		// _ = "end of CoverTab[79977]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:629
		_go_fuzz_dep_.CoverTab[79978]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:629
		// _ = "end of CoverTab[79978]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:629
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:629
	// _ = "end of CoverTab[79971]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:629
	_go_fuzz_dep_.CoverTab[79972]++

										s.channelzID = channelz.RegisterServer(&channelzServer{s}, "")
										channelz.Info(logger, s.channelzID, "Server created")
										return s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:633
	// _ = "end of CoverTab[79972]"
}

// printf records an event in s's event log, unless s has been stopped.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:636
// REQUIRES s.mu is held.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:638
func (s *Server) printf(format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:638
	_go_fuzz_dep_.CoverTab[79979]++
										if s.events != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:639
		_go_fuzz_dep_.CoverTab[79980]++
											s.events.Printf(format, a...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:640
		// _ = "end of CoverTab[79980]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:641
		_go_fuzz_dep_.CoverTab[79981]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:641
		// _ = "end of CoverTab[79981]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:641
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:641
	// _ = "end of CoverTab[79979]"
}

// errorf records an error in s's event log, unless s has been stopped.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:644
// REQUIRES s.mu is held.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:646
func (s *Server) errorf(format string, a ...interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:646
	_go_fuzz_dep_.CoverTab[79982]++
										if s.events != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:647
		_go_fuzz_dep_.CoverTab[79983]++
											s.events.Errorf(format, a...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:648
		// _ = "end of CoverTab[79983]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:649
		_go_fuzz_dep_.CoverTab[79984]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:649
		// _ = "end of CoverTab[79984]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:649
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:649
	// _ = "end of CoverTab[79982]"
}

// ServiceRegistrar wraps a single method that supports service registration. It
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:652
// enables users to pass concrete types other than grpc.Server to the service
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:652
// registration methods exported by the IDL generated code.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:655
type ServiceRegistrar interface {
	// RegisterService registers a service and its implementation to the
	// concrete type implementing this interface.  It may not be called
	// once the server has started serving.
	// desc describes the service and its methods and handlers. impl is the
	// service implementation which is passed to the method handlers.
	RegisterService(desc *ServiceDesc, impl interface{})
}

// RegisterService registers a service and its implementation to the gRPC
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:664
// server. It is called from the IDL generated code. This must be called before
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:664
// invoking Serve. If ss is non-nil (for legacy code), its type is checked to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:664
// ensure it implements sd.HandlerType.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:668
func (s *Server) RegisterService(sd *ServiceDesc, ss interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:668
	_go_fuzz_dep_.CoverTab[79985]++
										if ss != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:669
		_go_fuzz_dep_.CoverTab[79987]++
											ht := reflect.TypeOf(sd.HandlerType).Elem()
											st := reflect.TypeOf(ss)
											if !st.Implements(ht) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:672
			_go_fuzz_dep_.CoverTab[79988]++
												logger.Fatalf("grpc: Server.RegisterService found the handler of type %v that does not satisfy %v", st, ht)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:673
			// _ = "end of CoverTab[79988]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:674
			_go_fuzz_dep_.CoverTab[79989]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:674
			// _ = "end of CoverTab[79989]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:674
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:674
		// _ = "end of CoverTab[79987]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:675
		_go_fuzz_dep_.CoverTab[79990]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:675
		// _ = "end of CoverTab[79990]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:675
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:675
	// _ = "end of CoverTab[79985]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:675
	_go_fuzz_dep_.CoverTab[79986]++
										s.register(sd, ss)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:676
	// _ = "end of CoverTab[79986]"
}

func (s *Server) register(sd *ServiceDesc, ss interface{}) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:679
	_go_fuzz_dep_.CoverTab[79991]++
										s.mu.Lock()
										defer s.mu.Unlock()
										s.printf("RegisterService(%q)", sd.ServiceName)
										if s.serve {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:683
		_go_fuzz_dep_.CoverTab[79996]++
											logger.Fatalf("grpc: Server.RegisterService after Server.Serve for %q", sd.ServiceName)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:684
		// _ = "end of CoverTab[79996]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:685
		_go_fuzz_dep_.CoverTab[79997]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:685
		// _ = "end of CoverTab[79997]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:685
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:685
	// _ = "end of CoverTab[79991]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:685
	_go_fuzz_dep_.CoverTab[79992]++
										if _, ok := s.services[sd.ServiceName]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:686
		_go_fuzz_dep_.CoverTab[79998]++
											logger.Fatalf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:687
		// _ = "end of CoverTab[79998]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:688
		_go_fuzz_dep_.CoverTab[79999]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:688
		// _ = "end of CoverTab[79999]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:688
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:688
	// _ = "end of CoverTab[79992]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:688
	_go_fuzz_dep_.CoverTab[79993]++
										info := &serviceInfo{
		serviceImpl:	ss,
		methods:	make(map[string]*MethodDesc),
		streams:	make(map[string]*StreamDesc),
		mdata:		sd.Metadata,
	}
	for i := range sd.Methods {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:695
		_go_fuzz_dep_.CoverTab[80000]++
											d := &sd.Methods[i]
											info.methods[d.MethodName] = d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:697
		// _ = "end of CoverTab[80000]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:698
	// _ = "end of CoverTab[79993]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:698
	_go_fuzz_dep_.CoverTab[79994]++
										for i := range sd.Streams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:699
		_go_fuzz_dep_.CoverTab[80001]++
											d := &sd.Streams[i]
											info.streams[d.StreamName] = d
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:701
		// _ = "end of CoverTab[80001]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:702
	// _ = "end of CoverTab[79994]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:702
	_go_fuzz_dep_.CoverTab[79995]++
										s.services[sd.ServiceName] = info
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:703
	// _ = "end of CoverTab[79995]"
}

// MethodInfo contains the information of an RPC including its method name and type.
type MethodInfo struct {
	// Name is the method name only, without the service name or package name.
	Name	string
	// IsClientStream indicates whether the RPC is a client streaming RPC.
	IsClientStream	bool
	// IsServerStream indicates whether the RPC is a server streaming RPC.
	IsServerStream	bool
}

// ServiceInfo contains unary RPC method info, streaming RPC method info and metadata for a service.
type ServiceInfo struct {
	Methods	[]MethodInfo
	// Metadata is the metadata specified in ServiceDesc when registering service.
	Metadata	interface{}
}

// GetServiceInfo returns a map from service names to ServiceInfo.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:723
// Service names include the package names, in the form of <package>.<service>.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:725
func (s *Server) GetServiceInfo() map[string]ServiceInfo {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:725
	_go_fuzz_dep_.CoverTab[80002]++
										ret := make(map[string]ServiceInfo)
										for n, srv := range s.services {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:727
		_go_fuzz_dep_.CoverTab[80004]++
											methods := make([]MethodInfo, 0, len(srv.methods)+len(srv.streams))
											for m := range srv.methods {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:729
			_go_fuzz_dep_.CoverTab[80007]++
												methods = append(methods, MethodInfo{
				Name:		m,
				IsClientStream:	false,
				IsServerStream:	false,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:734
			// _ = "end of CoverTab[80007]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:735
		// _ = "end of CoverTab[80004]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:735
		_go_fuzz_dep_.CoverTab[80005]++
											for m, d := range srv.streams {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:736
			_go_fuzz_dep_.CoverTab[80008]++
												methods = append(methods, MethodInfo{
				Name:		m,
				IsClientStream:	d.ClientStreams,
				IsServerStream:	d.ServerStreams,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:741
			// _ = "end of CoverTab[80008]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:742
		// _ = "end of CoverTab[80005]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:742
		_go_fuzz_dep_.CoverTab[80006]++

											ret[n] = ServiceInfo{
			Methods:	methods,
			Metadata:	srv.mdata,
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:747
		// _ = "end of CoverTab[80006]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:748
	// _ = "end of CoverTab[80002]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:748
	_go_fuzz_dep_.CoverTab[80003]++
										return ret
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:749
	// _ = "end of CoverTab[80003]"
}

// ErrServerStopped indicates that the operation is now illegal because of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:752
// the server being stopped.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:754
var ErrServerStopped = errors.New("grpc: the server has been stopped")

type listenSocket struct {
	net.Listener
	channelzID	*channelz.Identifier
}

func (l *listenSocket) ChannelzMetric() *channelz.SocketInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:761
	_go_fuzz_dep_.CoverTab[80009]++
										return &channelz.SocketInternalMetric{
		SocketOptions:	channelz.GetSocketOption(l.Listener),
		LocalAddr:	l.Listener.Addr(),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:765
	// _ = "end of CoverTab[80009]"
}

func (l *listenSocket) Close() error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:768
	_go_fuzz_dep_.CoverTab[80010]++
										err := l.Listener.Close()
										channelz.RemoveEntry(l.channelzID)
										channelz.Info(logger, l.channelzID, "ListenSocket deleted")
										return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:772
	// _ = "end of CoverTab[80010]"
}

// Serve accepts incoming connections on the listener lis, creating a new
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:775
// ServerTransport and service goroutine for each. The service goroutines
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:775
// read gRPC requests and then call the registered handlers to reply to them.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:775
// Serve returns when lis.Accept fails with fatal errors.  lis will be closed when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:775
// this method returns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:775
// Serve will return a non-nil error unless Stop or GracefulStop is called.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:781
func (s *Server) Serve(lis net.Listener) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:781
	_go_fuzz_dep_.CoverTab[80011]++
										s.mu.Lock()
										s.printf("serving")
										s.serve = true
										if s.lis == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:785
		_go_fuzz_dep_.CoverTab[80016]++

											s.mu.Unlock()
											lis.Close()
											return ErrServerStopped
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:789
		// _ = "end of CoverTab[80016]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:790
		_go_fuzz_dep_.CoverTab[80017]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:790
		// _ = "end of CoverTab[80017]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:790
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:790
	// _ = "end of CoverTab[80011]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:790
	_go_fuzz_dep_.CoverTab[80012]++

										s.serveWG.Add(1)
										defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:793
		_go_fuzz_dep_.CoverTab[80018]++
											s.serveWG.Done()
											if s.quit.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:795
			_go_fuzz_dep_.CoverTab[80019]++

												<-s.done.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:797
			// _ = "end of CoverTab[80019]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:798
			_go_fuzz_dep_.CoverTab[80020]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:798
			// _ = "end of CoverTab[80020]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:798
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:798
		// _ = "end of CoverTab[80018]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:799
	// _ = "end of CoverTab[80012]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:799
	_go_fuzz_dep_.CoverTab[80013]++

										ls := &listenSocket{Listener: lis}
										s.lis[ls] = true

										defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:804
		_go_fuzz_dep_.CoverTab[80021]++
											s.mu.Lock()
											if s.lis != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:806
			_go_fuzz_dep_.CoverTab[80023]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:806
			return s.lis[ls]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:806
			// _ = "end of CoverTab[80023]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:806
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:806
			_go_fuzz_dep_.CoverTab[80024]++
												ls.Close()
												delete(s.lis, ls)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:808
			// _ = "end of CoverTab[80024]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:809
			_go_fuzz_dep_.CoverTab[80025]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:809
			// _ = "end of CoverTab[80025]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:809
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:809
		// _ = "end of CoverTab[80021]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:809
		_go_fuzz_dep_.CoverTab[80022]++
											s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:810
		// _ = "end of CoverTab[80022]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:811
	// _ = "end of CoverTab[80013]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:811
	_go_fuzz_dep_.CoverTab[80014]++

										var err error
										ls.channelzID, err = channelz.RegisterListenSocket(ls, s.channelzID, lis.Addr().String())
										if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:815
		_go_fuzz_dep_.CoverTab[80026]++
											s.mu.Unlock()
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:817
		// _ = "end of CoverTab[80026]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:818
		_go_fuzz_dep_.CoverTab[80027]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:818
		// _ = "end of CoverTab[80027]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:818
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:818
	// _ = "end of CoverTab[80014]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:818
	_go_fuzz_dep_.CoverTab[80015]++
										s.mu.Unlock()
										channelz.Info(logger, ls.channelzID, "ListenSocket created")

										var tempDelay time.Duration	// how long to sleep on accept failure
										for {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:823
		_go_fuzz_dep_.CoverTab[80028]++
											rawConn, err := lis.Accept()
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:825
			_go_fuzz_dep_.CoverTab[80030]++
												if ne, ok := err.(interface {
				Temporary() bool
			}); ok && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:828
				_go_fuzz_dep_.CoverTab[80033]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:828
				return ne.Temporary()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:828
				// _ = "end of CoverTab[80033]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:828
			}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:828
				_go_fuzz_dep_.CoverTab[80034]++
													if tempDelay == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:829
					_go_fuzz_dep_.CoverTab[80038]++
														tempDelay = 5 * time.Millisecond
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:830
					// _ = "end of CoverTab[80038]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:831
					_go_fuzz_dep_.CoverTab[80039]++
														tempDelay *= 2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:832
					// _ = "end of CoverTab[80039]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:833
				// _ = "end of CoverTab[80034]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:833
				_go_fuzz_dep_.CoverTab[80035]++
													if max := 1 * time.Second; tempDelay > max {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:834
					_go_fuzz_dep_.CoverTab[80040]++
														tempDelay = max
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:835
					// _ = "end of CoverTab[80040]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:836
					_go_fuzz_dep_.CoverTab[80041]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:836
					// _ = "end of CoverTab[80041]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:836
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:836
				// _ = "end of CoverTab[80035]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:836
				_go_fuzz_dep_.CoverTab[80036]++
													s.mu.Lock()
													s.printf("Accept error: %v; retrying in %v", err, tempDelay)
													s.mu.Unlock()
													timer := time.NewTimer(tempDelay)
													select {
				case <-timer.C:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:842
					_go_fuzz_dep_.CoverTab[80042]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:842
					// _ = "end of CoverTab[80042]"
				case <-s.quit.Done():
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:843
					_go_fuzz_dep_.CoverTab[80043]++
														timer.Stop()
														return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:845
					// _ = "end of CoverTab[80043]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:846
				// _ = "end of CoverTab[80036]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:846
				_go_fuzz_dep_.CoverTab[80037]++
													continue
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:847
				// _ = "end of CoverTab[80037]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:848
				_go_fuzz_dep_.CoverTab[80044]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:848
				// _ = "end of CoverTab[80044]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:848
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:848
			// _ = "end of CoverTab[80030]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:848
			_go_fuzz_dep_.CoverTab[80031]++
												s.mu.Lock()
												s.printf("done serving; Accept = %v", err)
												s.mu.Unlock()

												if s.quit.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:853
				_go_fuzz_dep_.CoverTab[80045]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:854
				// _ = "end of CoverTab[80045]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:855
				_go_fuzz_dep_.CoverTab[80046]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:855
				// _ = "end of CoverTab[80046]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:855
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:855
			// _ = "end of CoverTab[80031]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:855
			_go_fuzz_dep_.CoverTab[80032]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:856
			// _ = "end of CoverTab[80032]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:857
			_go_fuzz_dep_.CoverTab[80047]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:857
			// _ = "end of CoverTab[80047]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:857
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:857
		// _ = "end of CoverTab[80028]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:857
		_go_fuzz_dep_.CoverTab[80029]++
											tempDelay = 0

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:864
		s.serveWG.Add(1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:864
		_curRoutineNum94_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:864
		_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum94_)
											go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:865
			_go_fuzz_dep_.CoverTab[80048]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:865
			defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:865
				_go_fuzz_dep_.CoverTab[80049]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:865
				_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum94_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:865
				// _ = "end of CoverTab[80049]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:865
			}()
												s.handleRawConn(lis.Addr().String(), rawConn)
												s.serveWG.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:867
			// _ = "end of CoverTab[80048]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:868
		// _ = "end of CoverTab[80029]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:869
	// _ = "end of CoverTab[80015]"
}

// handleRawConn forks a goroutine to handle a just-accepted connection that
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:872
// has not had any I/O performed on it yet.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:874
func (s *Server) handleRawConn(lisAddr string, rawConn net.Conn) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:874
	_go_fuzz_dep_.CoverTab[80050]++
										if s.quit.HasFired() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:875
		_go_fuzz_dep_.CoverTab[80054]++
											rawConn.Close()
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:877
		// _ = "end of CoverTab[80054]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:878
		_go_fuzz_dep_.CoverTab[80055]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:878
		// _ = "end of CoverTab[80055]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:878
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:878
	// _ = "end of CoverTab[80050]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:878
	_go_fuzz_dep_.CoverTab[80051]++
										rawConn.SetDeadline(time.Now().Add(s.opts.connectionTimeout))

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:882
	st := s.newHTTP2Transport(rawConn)
	rawConn.SetDeadline(time.Time{})
	if st == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:884
		_go_fuzz_dep_.CoverTab[80056]++
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:885
		// _ = "end of CoverTab[80056]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:886
		_go_fuzz_dep_.CoverTab[80057]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:886
		// _ = "end of CoverTab[80057]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:886
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:886
	// _ = "end of CoverTab[80051]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:886
	_go_fuzz_dep_.CoverTab[80052]++

										if !s.addConn(lisAddr, st) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:888
		_go_fuzz_dep_.CoverTab[80058]++
											return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:889
		// _ = "end of CoverTab[80058]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
		_go_fuzz_dep_.CoverTab[80059]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
		// _ = "end of CoverTab[80059]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
	// _ = "end of CoverTab[80052]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
	_go_fuzz_dep_.CoverTab[80053]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
	_curRoutineNum95_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:890
	_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum95_)
										go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:891
		_go_fuzz_dep_.CoverTab[80060]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:891
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:891
			_go_fuzz_dep_.CoverTab[80061]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:891
			_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum95_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:891
			// _ = "end of CoverTab[80061]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:891
		}()
											s.serveStreams(st)
											s.removeConn(lisAddr, st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:893
		// _ = "end of CoverTab[80060]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:894
	// _ = "end of CoverTab[80053]"
}

func (s *Server) drainServerTransports(addr string) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:897
	_go_fuzz_dep_.CoverTab[80062]++
										s.mu.Lock()
										conns := s.conns[addr]
										for st := range conns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:900
		_go_fuzz_dep_.CoverTab[80064]++
											st.Drain()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:901
		// _ = "end of CoverTab[80064]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:902
	// _ = "end of CoverTab[80062]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:902
	_go_fuzz_dep_.CoverTab[80063]++
										s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:903
	// _ = "end of CoverTab[80063]"
}

// newHTTP2Transport sets up a http/2 transport (using the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:906
// gRPC http2 server transport in transport/http2_server.go).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:908
func (s *Server) newHTTP2Transport(c net.Conn) transport.ServerTransport {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:908
	_go_fuzz_dep_.CoverTab[80065]++
										config := &transport.ServerConfig{
		MaxStreams:		s.opts.maxConcurrentStreams,
		ConnectionTimeout:	s.opts.connectionTimeout,
		Credentials:		s.opts.creds,
		InTapHandle:		s.opts.inTapHandle,
		StatsHandlers:		s.opts.statsHandlers,
		KeepaliveParams:	s.opts.keepaliveParams,
		KeepalivePolicy:	s.opts.keepalivePolicy,
		InitialWindowSize:	s.opts.initialWindowSize,
		InitialConnWindowSize:	s.opts.initialConnWindowSize,
		WriteBufferSize:	s.opts.writeBufferSize,
		ReadBufferSize:		s.opts.readBufferSize,
		ChannelzParentID:	s.channelzID,
		MaxHeaderListSize:	s.opts.maxHeaderListSize,
		HeaderTableSize:	s.opts.headerTableSize,
	}
	st, err := transport.NewServerTransport(c, config)
	if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:926
		_go_fuzz_dep_.CoverTab[80067]++
											s.mu.Lock()
											s.errorf("NewServerTransport(%q) failed: %v", c.RemoteAddr(), err)
											s.mu.Unlock()

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:932
		if err != credentials.ErrConnDispatched {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:932
			_go_fuzz_dep_.CoverTab[80069]++

												if err != io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:934
				_go_fuzz_dep_.CoverTab[80071]++
													channelz.Info(logger, s.channelzID, "grpc: Server.Serve failed to create ServerTransport: ", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:935
				// _ = "end of CoverTab[80071]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:936
				_go_fuzz_dep_.CoverTab[80072]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:936
				// _ = "end of CoverTab[80072]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:936
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:936
			// _ = "end of CoverTab[80069]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:936
			_go_fuzz_dep_.CoverTab[80070]++
												c.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:937
			// _ = "end of CoverTab[80070]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:938
			_go_fuzz_dep_.CoverTab[80073]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:938
			// _ = "end of CoverTab[80073]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:938
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:938
		// _ = "end of CoverTab[80067]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:938
		_go_fuzz_dep_.CoverTab[80068]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:939
		// _ = "end of CoverTab[80068]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:940
		_go_fuzz_dep_.CoverTab[80074]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:940
		// _ = "end of CoverTab[80074]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:940
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:940
	// _ = "end of CoverTab[80065]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:940
	_go_fuzz_dep_.CoverTab[80066]++

										return st
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:942
	// _ = "end of CoverTab[80066]"
}

func (s *Server) serveStreams(st transport.ServerTransport) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:945
	_go_fuzz_dep_.CoverTab[80075]++
										defer st.Close(errors.New("finished serving streams for the server transport"))
										var wg sync.WaitGroup

										var roundRobinCounter uint32
										st.HandleStreams(func(stream *transport.Stream) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:950
		_go_fuzz_dep_.CoverTab[80077]++
											wg.Add(1)
											if s.opts.numServerWorkers > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:952
			_go_fuzz_dep_.CoverTab[80078]++
												data := &serverWorkerData{st: st, wg: &wg, stream: stream}
												select {
			case s.serverWorkerChannels[atomic.AddUint32(&roundRobinCounter, 1)%s.opts.numServerWorkers] <- data:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:955
				_go_fuzz_dep_.CoverTab[80079]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:955
				// _ = "end of CoverTab[80079]"
			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:956
				_go_fuzz_dep_.CoverTab[80080]++

													go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:958
					_go_fuzz_dep_.CoverTab[80081]++
														s.handleStream(st, stream, s.traceInfo(st, stream))
														wg.Done()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:960
					// _ = "end of CoverTab[80081]"
				}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:961
				// _ = "end of CoverTab[80080]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:962
			// _ = "end of CoverTab[80078]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:963
			_go_fuzz_dep_.CoverTab[80082]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:963
			_curRoutineNum96_ := _atomic_.AddUint32(&_go_fuzz_dep_.RoutineNum, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:963
			_go_fuzz_dep_.RoutineInfo.AddCreatedRoutineNum(_curRoutineNum96_)
												go func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:964
				_go_fuzz_dep_.CoverTab[80083]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:964
				defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:964
					_go_fuzz_dep_.CoverTab[80084]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:964
					_go_fuzz_dep_.RoutineInfo.AddTerminatedRoutineNum(_curRoutineNum96_)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:964
					// _ = "end of CoverTab[80084]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:964
				}()
													defer wg.Done()
													s.handleStream(st, stream, s.traceInfo(st, stream))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:966
				// _ = "end of CoverTab[80083]"
			}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:967
			// _ = "end of CoverTab[80082]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:968
		// _ = "end of CoverTab[80077]"
	}, func(ctx context.Context, method string) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:969
		_go_fuzz_dep_.CoverTab[80085]++
											if !EnableTracing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:970
			_go_fuzz_dep_.CoverTab[80087]++
												return ctx
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:971
			// _ = "end of CoverTab[80087]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:972
			_go_fuzz_dep_.CoverTab[80088]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:972
			// _ = "end of CoverTab[80088]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:972
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:972
		// _ = "end of CoverTab[80085]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:972
		_go_fuzz_dep_.CoverTab[80086]++
											tr := trace.New("grpc.Recv."+methodFamily(method), method)
											return trace.NewContext(ctx, tr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:974
		// _ = "end of CoverTab[80086]"
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:975
	// _ = "end of CoverTab[80075]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:975
	_go_fuzz_dep_.CoverTab[80076]++
										wg.Wait()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:976
	// _ = "end of CoverTab[80076]"
}

var _ http.Handler = (*Server)(nil)

// ServeHTTP implements the Go standard library's http.Handler
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// interface by responding to the gRPC request r, by looking up
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// the requested gRPC method in the gRPC server s.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// The provided HTTP request must have arrived on an HTTP/2
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// connection. When using the Go standard library's server,
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// practically this means that the Request must also have arrived
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// over TLS.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// To share one port (such as 443 for https) between gRPC and an
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// existing http.Handler, use a root http.Handler such as:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//	if r.ProtoMajor == 2 && strings.HasPrefix(
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//		r.Header.Get("Content-Type"), "application/grpc") {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//		grpcServer.ServeHTTP(w, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//		yourMux.ServeHTTP(w, r)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// Note that ServeHTTP uses Go's HTTP/2 server implementation which is totally
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// separate from grpc-go's HTTP/2 server. Performance and features may vary
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// between the two paths. ServeHTTP does not support some gRPC features
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// available through grpc-go's HTTP/2 server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:981
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1009
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1009
	_go_fuzz_dep_.CoverTab[80089]++
											st, err := transport.NewServerHandlerTransport(w, r, s.opts.statsHandlers)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1011
		_go_fuzz_dep_.CoverTab[80092]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1014
		return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1014
		// _ = "end of CoverTab[80092]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1015
		_go_fuzz_dep_.CoverTab[80093]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1015
		// _ = "end of CoverTab[80093]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1015
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1015
	// _ = "end of CoverTab[80089]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1015
	_go_fuzz_dep_.CoverTab[80090]++
											if !s.addConn(listenerAddressForServeHTTP, st) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1016
		_go_fuzz_dep_.CoverTab[80094]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1017
		// _ = "end of CoverTab[80094]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1018
		_go_fuzz_dep_.CoverTab[80095]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1018
		// _ = "end of CoverTab[80095]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1018
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1018
	// _ = "end of CoverTab[80090]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1018
	_go_fuzz_dep_.CoverTab[80091]++
											defer s.removeConn(listenerAddressForServeHTTP, st)
											s.serveStreams(st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1020
	// _ = "end of CoverTab[80091]"
}

// traceInfo returns a traceInfo and associates it with stream, if tracing is enabled.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1023
// If tracing is not enabled, it returns nil.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1025
func (s *Server) traceInfo(st transport.ServerTransport, stream *transport.Stream) (trInfo *traceInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1025
	_go_fuzz_dep_.CoverTab[80096]++
											if !EnableTracing {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1026
		_go_fuzz_dep_.CoverTab[80100]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1027
		// _ = "end of CoverTab[80100]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1028
		_go_fuzz_dep_.CoverTab[80101]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1028
		// _ = "end of CoverTab[80101]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1028
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1028
	// _ = "end of CoverTab[80096]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1028
	_go_fuzz_dep_.CoverTab[80097]++
											tr, ok := trace.FromContext(stream.Context())
											if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1030
		_go_fuzz_dep_.CoverTab[80102]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1031
		// _ = "end of CoverTab[80102]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1032
		_go_fuzz_dep_.CoverTab[80103]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1032
		// _ = "end of CoverTab[80103]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1032
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1032
	// _ = "end of CoverTab[80097]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1032
	_go_fuzz_dep_.CoverTab[80098]++

											trInfo = &traceInfo{
		tr:	tr,
		firstLine: firstLine{
			client:		false,
			remoteAddr:	st.RemoteAddr(),
		},
	}
	if dl, ok := stream.Context().Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1041
		_go_fuzz_dep_.CoverTab[80104]++
												trInfo.firstLine.deadline = time.Until(dl)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1042
		// _ = "end of CoverTab[80104]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1043
		_go_fuzz_dep_.CoverTab[80105]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1043
		// _ = "end of CoverTab[80105]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1043
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1043
	// _ = "end of CoverTab[80098]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1043
	_go_fuzz_dep_.CoverTab[80099]++
											return trInfo
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1044
	// _ = "end of CoverTab[80099]"
}

func (s *Server) addConn(addr string, st transport.ServerTransport) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1047
	_go_fuzz_dep_.CoverTab[80106]++
											s.mu.Lock()
											defer s.mu.Unlock()
											if s.conns == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1050
		_go_fuzz_dep_.CoverTab[80110]++
												st.Close(errors.New("Server.addConn called when server has already been stopped"))
												return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1052
		// _ = "end of CoverTab[80110]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1053
		_go_fuzz_dep_.CoverTab[80111]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1053
		// _ = "end of CoverTab[80111]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1053
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1053
	// _ = "end of CoverTab[80106]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1053
	_go_fuzz_dep_.CoverTab[80107]++
											if s.drain {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1054
		_go_fuzz_dep_.CoverTab[80112]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1057
		st.Drain()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1057
		// _ = "end of CoverTab[80112]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1058
		_go_fuzz_dep_.CoverTab[80113]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1058
		// _ = "end of CoverTab[80113]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1058
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1058
	// _ = "end of CoverTab[80107]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1058
	_go_fuzz_dep_.CoverTab[80108]++

											if s.conns[addr] == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1060
		_go_fuzz_dep_.CoverTab[80114]++

												s.conns[addr] = make(map[transport.ServerTransport]bool)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1062
		// _ = "end of CoverTab[80114]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1063
		_go_fuzz_dep_.CoverTab[80115]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1063
		// _ = "end of CoverTab[80115]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1063
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1063
	// _ = "end of CoverTab[80108]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1063
	_go_fuzz_dep_.CoverTab[80109]++
											s.conns[addr][st] = true
											return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1065
	// _ = "end of CoverTab[80109]"
}

func (s *Server) removeConn(addr string, st transport.ServerTransport) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1068
	_go_fuzz_dep_.CoverTab[80116]++
											s.mu.Lock()
											defer s.mu.Unlock()

											conns := s.conns[addr]
											if conns != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1073
		_go_fuzz_dep_.CoverTab[80117]++
												delete(conns, st)
												if len(conns) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1075
			_go_fuzz_dep_.CoverTab[80119]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1079
			delete(s.conns, addr)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1079
			// _ = "end of CoverTab[80119]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1080
			_go_fuzz_dep_.CoverTab[80120]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1080
			// _ = "end of CoverTab[80120]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1080
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1080
		// _ = "end of CoverTab[80117]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1080
		_go_fuzz_dep_.CoverTab[80118]++
												s.cv.Broadcast()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1081
		// _ = "end of CoverTab[80118]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1082
		_go_fuzz_dep_.CoverTab[80121]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1082
		// _ = "end of CoverTab[80121]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1082
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1082
	// _ = "end of CoverTab[80116]"
}

func (s *Server) channelzMetric() *channelz.ServerInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1085
	_go_fuzz_dep_.CoverTab[80122]++
											return &channelz.ServerInternalMetric{
		CallsStarted:			atomic.LoadInt64(&s.czData.callsStarted),
		CallsSucceeded:			atomic.LoadInt64(&s.czData.callsSucceeded),
		CallsFailed:			atomic.LoadInt64(&s.czData.callsFailed),
		LastCallStartedTimestamp:	time.Unix(0, atomic.LoadInt64(&s.czData.lastCallStartedTime)),
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1091
	// _ = "end of CoverTab[80122]"
}

func (s *Server) incrCallsStarted() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1094
	_go_fuzz_dep_.CoverTab[80123]++
											atomic.AddInt64(&s.czData.callsStarted, 1)
											atomic.StoreInt64(&s.czData.lastCallStartedTime, time.Now().UnixNano())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1096
	// _ = "end of CoverTab[80123]"
}

func (s *Server) incrCallsSucceeded() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1099
	_go_fuzz_dep_.CoverTab[80124]++
											atomic.AddInt64(&s.czData.callsSucceeded, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1100
	// _ = "end of CoverTab[80124]"
}

func (s *Server) incrCallsFailed() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1103
	_go_fuzz_dep_.CoverTab[80125]++
											atomic.AddInt64(&s.czData.callsFailed, 1)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1104
	// _ = "end of CoverTab[80125]"
}

func (s *Server) sendResponse(t transport.ServerTransport, stream *transport.Stream, msg interface{}, cp Compressor, opts *transport.Options, comp encoding.Compressor) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1107
	_go_fuzz_dep_.CoverTab[80126]++
											data, err := encode(s.getCodec(stream.ContentSubtype()), msg)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1109
		_go_fuzz_dep_.CoverTab[80131]++
												channelz.Error(logger, s.channelzID, "grpc: server failed to encode response: ", err)
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1111
		// _ = "end of CoverTab[80131]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1112
		_go_fuzz_dep_.CoverTab[80132]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1112
		// _ = "end of CoverTab[80132]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1112
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1112
	// _ = "end of CoverTab[80126]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1112
	_go_fuzz_dep_.CoverTab[80127]++
											compData, err := compress(data, cp, comp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1114
		_go_fuzz_dep_.CoverTab[80133]++
												channelz.Error(logger, s.channelzID, "grpc: server failed to compress response: ", err)
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1116
		// _ = "end of CoverTab[80133]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1117
		_go_fuzz_dep_.CoverTab[80134]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1117
		// _ = "end of CoverTab[80134]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1117
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1117
	// _ = "end of CoverTab[80127]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1117
	_go_fuzz_dep_.CoverTab[80128]++
											hdr, payload := msgHeader(data, compData)

											if len(payload) > s.opts.maxSendMessageSize {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1120
		_go_fuzz_dep_.CoverTab[80135]++
												return status.Errorf(codes.ResourceExhausted, "grpc: trying to send message larger than max (%d vs. %d)", len(payload), s.opts.maxSendMessageSize)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1121
		// _ = "end of CoverTab[80135]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1122
		_go_fuzz_dep_.CoverTab[80136]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1122
		// _ = "end of CoverTab[80136]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1122
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1122
	// _ = "end of CoverTab[80128]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1122
	_go_fuzz_dep_.CoverTab[80129]++
											err = t.Write(stream, hdr, payload, opts)
											if err == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1124
		_go_fuzz_dep_.CoverTab[80137]++
												for _, sh := range s.opts.statsHandlers {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1125
			_go_fuzz_dep_.CoverTab[80138]++
													sh.HandleRPC(stream.Context(), outPayload(false, msg, data, payload, time.Now()))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1126
			// _ = "end of CoverTab[80138]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1127
		// _ = "end of CoverTab[80137]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1128
		_go_fuzz_dep_.CoverTab[80139]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1128
		// _ = "end of CoverTab[80139]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1128
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1128
	// _ = "end of CoverTab[80129]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1128
	_go_fuzz_dep_.CoverTab[80130]++
											return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1129
	// _ = "end of CoverTab[80130]"
}

// chainUnaryServerInterceptors chains all unary server interceptors into one.
func chainUnaryServerInterceptors(s *Server) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1133
	_go_fuzz_dep_.CoverTab[80140]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1136
	interceptors := s.opts.chainUnaryInts
	if s.opts.unaryInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1137
		_go_fuzz_dep_.CoverTab[80143]++
												interceptors = append([]UnaryServerInterceptor{s.opts.unaryInt}, s.opts.chainUnaryInts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1138
		// _ = "end of CoverTab[80143]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1139
		_go_fuzz_dep_.CoverTab[80144]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1139
		// _ = "end of CoverTab[80144]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1139
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1139
	// _ = "end of CoverTab[80140]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1139
	_go_fuzz_dep_.CoverTab[80141]++

											var chainedInt UnaryServerInterceptor
											if len(interceptors) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1142
		_go_fuzz_dep_.CoverTab[80145]++
												chainedInt = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1143
		// _ = "end of CoverTab[80145]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1144
		_go_fuzz_dep_.CoverTab[80146]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1144
		if len(interceptors) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1144
			_go_fuzz_dep_.CoverTab[80147]++
													chainedInt = interceptors[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1145
			// _ = "end of CoverTab[80147]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1146
			_go_fuzz_dep_.CoverTab[80148]++
													chainedInt = chainUnaryInterceptors(interceptors)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1147
			// _ = "end of CoverTab[80148]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1148
		// _ = "end of CoverTab[80146]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1148
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1148
	// _ = "end of CoverTab[80141]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1148
	_go_fuzz_dep_.CoverTab[80142]++

											s.opts.unaryInt = chainedInt
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1150
	// _ = "end of CoverTab[80142]"
}

func chainUnaryInterceptors(interceptors []UnaryServerInterceptor) UnaryServerInterceptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1153
	_go_fuzz_dep_.CoverTab[80149]++
											return func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1154
		_go_fuzz_dep_.CoverTab[80150]++
												return interceptors[0](ctx, req, info, getChainUnaryHandler(interceptors, 0, info, handler))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1155
		// _ = "end of CoverTab[80150]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1156
	// _ = "end of CoverTab[80149]"
}

func getChainUnaryHandler(interceptors []UnaryServerInterceptor, curr int, info *UnaryServerInfo, finalHandler UnaryHandler) UnaryHandler {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1159
	_go_fuzz_dep_.CoverTab[80151]++
											if curr == len(interceptors)-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1160
		_go_fuzz_dep_.CoverTab[80153]++
												return finalHandler
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1161
		// _ = "end of CoverTab[80153]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1162
		_go_fuzz_dep_.CoverTab[80154]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1162
		// _ = "end of CoverTab[80154]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1162
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1162
	// _ = "end of CoverTab[80151]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1162
	_go_fuzz_dep_.CoverTab[80152]++
											return func(ctx context.Context, req interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1163
		_go_fuzz_dep_.CoverTab[80155]++
												return interceptors[curr+1](ctx, req, info, getChainUnaryHandler(interceptors, curr+1, info, finalHandler))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1164
		// _ = "end of CoverTab[80155]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1165
	// _ = "end of CoverTab[80152]"
}

func (s *Server) processUnaryRPC(t transport.ServerTransport, stream *transport.Stream, info *serviceInfo, md *MethodDesc, trInfo *traceInfo) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1168
	_go_fuzz_dep_.CoverTab[80156]++
											shs := s.opts.statsHandlers
											if len(shs) != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		_go_fuzz_dep_.CoverTab[80176]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		return trInfo != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		// _ = "end of CoverTab[80176]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		_go_fuzz_dep_.CoverTab[80177]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		return channelz.IsOn()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		// _ = "end of CoverTab[80177]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1170
		_go_fuzz_dep_.CoverTab[80178]++
												if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1171
			_go_fuzz_dep_.CoverTab[80182]++
													s.incrCallsStarted()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1172
			// _ = "end of CoverTab[80182]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1173
			_go_fuzz_dep_.CoverTab[80183]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1173
			// _ = "end of CoverTab[80183]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1173
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1173
		// _ = "end of CoverTab[80178]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1173
		_go_fuzz_dep_.CoverTab[80179]++
												var statsBegin *stats.Begin
												for _, sh := range shs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1175
			_go_fuzz_dep_.CoverTab[80184]++
													beginTime := time.Now()
													statsBegin = &stats.Begin{
				BeginTime:	beginTime,
				IsClientStream:	false,
				IsServerStream:	false,
			}
													sh.HandleRPC(stream.Context(), statsBegin)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1182
			// _ = "end of CoverTab[80184]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1183
		// _ = "end of CoverTab[80179]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1183
		_go_fuzz_dep_.CoverTab[80180]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1184
			_go_fuzz_dep_.CoverTab[80185]++
													trInfo.tr.LazyLog(&trInfo.firstLine, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1185
			// _ = "end of CoverTab[80185]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1186
			_go_fuzz_dep_.CoverTab[80186]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1186
			// _ = "end of CoverTab[80186]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1186
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1186
		// _ = "end of CoverTab[80180]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1186
		_go_fuzz_dep_.CoverTab[80181]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1197
		defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1197
			_go_fuzz_dep_.CoverTab[80187]++
													if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1198
				_go_fuzz_dep_.CoverTab[80190]++
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1199
					_go_fuzz_dep_.CoverTab[80192]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1199
					return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1199
					// _ = "end of CoverTab[80192]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1199
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1199
					_go_fuzz_dep_.CoverTab[80193]++
															trInfo.tr.LazyLog(&fmtStringer{"%v", []interface{}{err}}, true)
															trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1201
					// _ = "end of CoverTab[80193]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1202
					_go_fuzz_dep_.CoverTab[80194]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1202
					// _ = "end of CoverTab[80194]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1202
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1202
				// _ = "end of CoverTab[80190]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1202
				_go_fuzz_dep_.CoverTab[80191]++
														trInfo.tr.Finish()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1203
				// _ = "end of CoverTab[80191]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1204
				_go_fuzz_dep_.CoverTab[80195]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1204
				// _ = "end of CoverTab[80195]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1204
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1204
			// _ = "end of CoverTab[80187]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1204
			_go_fuzz_dep_.CoverTab[80188]++

													for _, sh := range shs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1206
				_go_fuzz_dep_.CoverTab[80196]++
														end := &stats.End{
					BeginTime:	statsBegin.BeginTime,
					EndTime:	time.Now(),
				}
				if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1211
					_go_fuzz_dep_.CoverTab[80198]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1211
					return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1211
					// _ = "end of CoverTab[80198]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1211
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1211
					_go_fuzz_dep_.CoverTab[80199]++
															end.Error = toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1212
					// _ = "end of CoverTab[80199]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1213
					_go_fuzz_dep_.CoverTab[80200]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1213
					// _ = "end of CoverTab[80200]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1213
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1213
				// _ = "end of CoverTab[80196]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1213
				_go_fuzz_dep_.CoverTab[80197]++
														sh.HandleRPC(stream.Context(), end)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1214
				// _ = "end of CoverTab[80197]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1215
			// _ = "end of CoverTab[80188]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1215
			_go_fuzz_dep_.CoverTab[80189]++

													if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1217
				_go_fuzz_dep_.CoverTab[80201]++
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1218
					_go_fuzz_dep_.CoverTab[80202]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1218
					return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1218
					// _ = "end of CoverTab[80202]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1218
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1218
					_go_fuzz_dep_.CoverTab[80203]++
															s.incrCallsFailed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1219
					// _ = "end of CoverTab[80203]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1220
					_go_fuzz_dep_.CoverTab[80204]++
															s.incrCallsSucceeded()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1221
					// _ = "end of CoverTab[80204]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1222
				// _ = "end of CoverTab[80201]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1223
				_go_fuzz_dep_.CoverTab[80205]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1223
				// _ = "end of CoverTab[80205]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1223
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1223
			// _ = "end of CoverTab[80189]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1224
		// _ = "end of CoverTab[80181]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1225
		_go_fuzz_dep_.CoverTab[80206]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1225
		// _ = "end of CoverTab[80206]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1225
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1225
	// _ = "end of CoverTab[80156]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1225
	_go_fuzz_dep_.CoverTab[80157]++
											var binlogs []binarylog.MethodLogger
											if ml := binarylog.GetMethodLogger(stream.Method()); ml != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1227
		_go_fuzz_dep_.CoverTab[80207]++
												binlogs = append(binlogs, ml)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1228
		// _ = "end of CoverTab[80207]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1229
		_go_fuzz_dep_.CoverTab[80208]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1229
		// _ = "end of CoverTab[80208]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1229
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1229
	// _ = "end of CoverTab[80157]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1229
	_go_fuzz_dep_.CoverTab[80158]++
											if s.opts.binaryLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1230
		_go_fuzz_dep_.CoverTab[80209]++
												if ml := s.opts.binaryLogger.GetMethodLogger(stream.Method()); ml != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1231
			_go_fuzz_dep_.CoverTab[80210]++
													binlogs = append(binlogs, ml)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1232
			// _ = "end of CoverTab[80210]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1233
			_go_fuzz_dep_.CoverTab[80211]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1233
			// _ = "end of CoverTab[80211]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1233
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1233
		// _ = "end of CoverTab[80209]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1234
		_go_fuzz_dep_.CoverTab[80212]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1234
		// _ = "end of CoverTab[80212]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1234
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1234
	// _ = "end of CoverTab[80158]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1234
	_go_fuzz_dep_.CoverTab[80159]++
											if len(binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1235
		_go_fuzz_dep_.CoverTab[80213]++
												ctx := stream.Context()
												md, _ := metadata.FromIncomingContext(ctx)
												logEntry := &binarylog.ClientHeader{
			Header:		md,
			MethodName:	stream.Method(),
			PeerAddr:	nil,
		}
		if deadline, ok := ctx.Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1243
			_go_fuzz_dep_.CoverTab[80217]++
													logEntry.Timeout = time.Until(deadline)
													if logEntry.Timeout < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1245
				_go_fuzz_dep_.CoverTab[80218]++
														logEntry.Timeout = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1246
				// _ = "end of CoverTab[80218]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1247
				_go_fuzz_dep_.CoverTab[80219]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1247
				// _ = "end of CoverTab[80219]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1247
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1247
			// _ = "end of CoverTab[80217]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1248
			_go_fuzz_dep_.CoverTab[80220]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1248
			// _ = "end of CoverTab[80220]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1248
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1248
		// _ = "end of CoverTab[80213]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1248
		_go_fuzz_dep_.CoverTab[80214]++
												if a := md[":authority"]; len(a) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1249
			_go_fuzz_dep_.CoverTab[80221]++
													logEntry.Authority = a[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1250
			// _ = "end of CoverTab[80221]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1251
			_go_fuzz_dep_.CoverTab[80222]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1251
			// _ = "end of CoverTab[80222]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1251
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1251
		// _ = "end of CoverTab[80214]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1251
		_go_fuzz_dep_.CoverTab[80215]++
												if peer, ok := peer.FromContext(ctx); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1252
			_go_fuzz_dep_.CoverTab[80223]++
													logEntry.PeerAddr = peer.Addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1253
			// _ = "end of CoverTab[80223]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1254
			_go_fuzz_dep_.CoverTab[80224]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1254
			// _ = "end of CoverTab[80224]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1254
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1254
		// _ = "end of CoverTab[80215]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1254
		_go_fuzz_dep_.CoverTab[80216]++
												for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1255
			_go_fuzz_dep_.CoverTab[80225]++
													binlog.Log(ctx, logEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1256
			// _ = "end of CoverTab[80225]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1257
		// _ = "end of CoverTab[80216]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1258
		_go_fuzz_dep_.CoverTab[80226]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1258
		// _ = "end of CoverTab[80226]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1258
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1258
	// _ = "end of CoverTab[80159]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1258
	_go_fuzz_dep_.CoverTab[80160]++

	// comp and cp are used for compression.  decomp and dc are used for
	// decompression.  If comp and decomp are both set, they are the same;
	// however they are kept separate to ensure that at most one of the
											// compressor/decompressor variable pairs are set for use later.
											var comp, decomp encoding.Compressor
											var cp Compressor
											var dc Decompressor
											var sendCompressorName string

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1271
	if rc := stream.RecvCompress(); s.opts.dc != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1271
		_go_fuzz_dep_.CoverTab[80227]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1271
		return s.opts.dc.Type() == rc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1271
		// _ = "end of CoverTab[80227]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1271
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1271
		_go_fuzz_dep_.CoverTab[80228]++
												dc = s.opts.dc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1272
		// _ = "end of CoverTab[80228]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
		_go_fuzz_dep_.CoverTab[80229]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
		if rc != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
			_go_fuzz_dep_.CoverTab[80230]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
			return rc != encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
			// _ = "end of CoverTab[80230]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1273
			_go_fuzz_dep_.CoverTab[80231]++
													decomp = encoding.GetCompressor(rc)
													if decomp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1275
				_go_fuzz_dep_.CoverTab[80232]++
														st := status.Newf(codes.Unimplemented, "grpc: Decompressor is not installed for grpc-encoding %q", rc)
														t.WriteStatus(stream, st)
														return st.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1278
				// _ = "end of CoverTab[80232]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1279
				_go_fuzz_dep_.CoverTab[80233]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1279
				// _ = "end of CoverTab[80233]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1279
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1279
			// _ = "end of CoverTab[80231]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
			_go_fuzz_dep_.CoverTab[80234]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
			// _ = "end of CoverTab[80234]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
		// _ = "end of CoverTab[80229]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
	// _ = "end of CoverTab[80160]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1280
	_go_fuzz_dep_.CoverTab[80161]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1286
	if s.opts.cp != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1286
		_go_fuzz_dep_.CoverTab[80235]++
												cp = s.opts.cp
												sendCompressorName = cp.Type()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1288
		// _ = "end of CoverTab[80235]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
		_go_fuzz_dep_.CoverTab[80236]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
		if rc := stream.RecvCompress(); rc != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
			_go_fuzz_dep_.CoverTab[80237]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
			return rc != encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
			// _ = "end of CoverTab[80237]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1289
			_go_fuzz_dep_.CoverTab[80238]++

													comp = encoding.GetCompressor(rc)
													if comp != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1292
				_go_fuzz_dep_.CoverTab[80239]++
														sendCompressorName = comp.Name()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1293
				// _ = "end of CoverTab[80239]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1294
				_go_fuzz_dep_.CoverTab[80240]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1294
				// _ = "end of CoverTab[80240]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1294
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1294
			// _ = "end of CoverTab[80238]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
			_go_fuzz_dep_.CoverTab[80241]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
			// _ = "end of CoverTab[80241]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
		// _ = "end of CoverTab[80236]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
	// _ = "end of CoverTab[80161]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1295
	_go_fuzz_dep_.CoverTab[80162]++

											if sendCompressorName != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1297
		_go_fuzz_dep_.CoverTab[80242]++
												if err := stream.SetSendCompress(sendCompressorName); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1298
			_go_fuzz_dep_.CoverTab[80243]++
													return status.Errorf(codes.Internal, "grpc: failed to set send compressor: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1299
			// _ = "end of CoverTab[80243]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1300
			_go_fuzz_dep_.CoverTab[80244]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1300
			// _ = "end of CoverTab[80244]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1300
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1300
		// _ = "end of CoverTab[80242]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1301
		_go_fuzz_dep_.CoverTab[80245]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1301
		// _ = "end of CoverTab[80245]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1301
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1301
	// _ = "end of CoverTab[80162]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1301
	_go_fuzz_dep_.CoverTab[80163]++

											var payInfo *payloadInfo
											if len(shs) != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1304
		_go_fuzz_dep_.CoverTab[80246]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1304
		return len(binlogs) != 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1304
		// _ = "end of CoverTab[80246]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1304
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1304
		_go_fuzz_dep_.CoverTab[80247]++
												payInfo = &payloadInfo{}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1305
		// _ = "end of CoverTab[80247]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1306
		_go_fuzz_dep_.CoverTab[80248]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1306
		// _ = "end of CoverTab[80248]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1306
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1306
	// _ = "end of CoverTab[80163]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1306
	_go_fuzz_dep_.CoverTab[80164]++
											d, err := recvAndDecompress(&parser{r: stream}, stream, dc, s.opts.maxReceiveMessageSize, payInfo, decomp)
											if err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1308
		_go_fuzz_dep_.CoverTab[80249]++
												if e := t.WriteStatus(stream, status.Convert(err)); e != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1309
			_go_fuzz_dep_.CoverTab[80251]++
													channelz.Warningf(logger, s.channelzID, "grpc: Server.processUnaryRPC failed to write status: %v", e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1310
			// _ = "end of CoverTab[80251]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1311
			_go_fuzz_dep_.CoverTab[80252]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1311
			// _ = "end of CoverTab[80252]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1311
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1311
		// _ = "end of CoverTab[80249]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1311
		_go_fuzz_dep_.CoverTab[80250]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1312
		// _ = "end of CoverTab[80250]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1313
		_go_fuzz_dep_.CoverTab[80253]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1313
		// _ = "end of CoverTab[80253]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1313
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1313
	// _ = "end of CoverTab[80164]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1313
	_go_fuzz_dep_.CoverTab[80165]++
											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1314
		_go_fuzz_dep_.CoverTab[80254]++
												t.IncrMsgRecv()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1315
		// _ = "end of CoverTab[80254]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1316
		_go_fuzz_dep_.CoverTab[80255]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1316
		// _ = "end of CoverTab[80255]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1316
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1316
	// _ = "end of CoverTab[80165]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1316
	_go_fuzz_dep_.CoverTab[80166]++
											df := func(v interface{}) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1317
		_go_fuzz_dep_.CoverTab[80256]++
												if err := s.getCodec(stream.ContentSubtype()).Unmarshal(d, v); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1318
			_go_fuzz_dep_.CoverTab[80261]++
													return status.Errorf(codes.Internal, "grpc: error unmarshalling request: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1319
			// _ = "end of CoverTab[80261]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1320
			_go_fuzz_dep_.CoverTab[80262]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1320
			// _ = "end of CoverTab[80262]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1320
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1320
		// _ = "end of CoverTab[80256]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1320
		_go_fuzz_dep_.CoverTab[80257]++
												for _, sh := range shs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1321
			_go_fuzz_dep_.CoverTab[80263]++
													sh.HandleRPC(stream.Context(), &stats.InPayload{
				RecvTime:		time.Now(),
				Payload:		v,
				Length:			len(d),
				WireLength:		payInfo.compressedLength + headerLen,
				CompressedLength:	payInfo.compressedLength,
				Data:			d,
			})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1329
			// _ = "end of CoverTab[80263]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1330
		// _ = "end of CoverTab[80257]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1330
		_go_fuzz_dep_.CoverTab[80258]++
												if len(binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1331
			_go_fuzz_dep_.CoverTab[80264]++
													cm := &binarylog.ClientMessage{
				Message: d,
			}
			for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1335
				_go_fuzz_dep_.CoverTab[80265]++
														binlog.Log(stream.Context(), cm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1336
				// _ = "end of CoverTab[80265]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1337
			// _ = "end of CoverTab[80264]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1338
			_go_fuzz_dep_.CoverTab[80266]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1338
			// _ = "end of CoverTab[80266]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1338
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1338
		// _ = "end of CoverTab[80258]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1338
		_go_fuzz_dep_.CoverTab[80259]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1339
			_go_fuzz_dep_.CoverTab[80267]++
													trInfo.tr.LazyLog(&payload{sent: false, msg: v}, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1340
			// _ = "end of CoverTab[80267]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1341
			_go_fuzz_dep_.CoverTab[80268]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1341
			// _ = "end of CoverTab[80268]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1341
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1341
		// _ = "end of CoverTab[80259]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1341
		_go_fuzz_dep_.CoverTab[80260]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1342
		// _ = "end of CoverTab[80260]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1343
	// _ = "end of CoverTab[80166]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1343
	_go_fuzz_dep_.CoverTab[80167]++
											ctx := NewContextWithServerTransportStream(stream.Context(), stream)
											reply, appErr := md.Handler(info.serviceImpl, ctx, df, s.opts.unaryInt)
											if appErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1346
		_go_fuzz_dep_.CoverTab[80269]++
												appStatus, ok := status.FromError(appErr)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1348
			_go_fuzz_dep_.CoverTab[80274]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1351
			appStatus = status.FromContextError(appErr)
													appErr = appStatus.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1352
			// _ = "end of CoverTab[80274]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1353
			_go_fuzz_dep_.CoverTab[80275]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1353
			// _ = "end of CoverTab[80275]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1353
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1353
		// _ = "end of CoverTab[80269]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1353
		_go_fuzz_dep_.CoverTab[80270]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1354
			_go_fuzz_dep_.CoverTab[80276]++
													trInfo.tr.LazyLog(stringer(appStatus.Message()), true)
													trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1356
			// _ = "end of CoverTab[80276]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1357
			_go_fuzz_dep_.CoverTab[80277]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1357
			// _ = "end of CoverTab[80277]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1357
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1357
		// _ = "end of CoverTab[80270]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1357
		_go_fuzz_dep_.CoverTab[80271]++
												if e := t.WriteStatus(stream, appStatus); e != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1358
			_go_fuzz_dep_.CoverTab[80278]++
													channelz.Warningf(logger, s.channelzID, "grpc: Server.processUnaryRPC failed to write status: %v", e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1359
			// _ = "end of CoverTab[80278]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1360
			_go_fuzz_dep_.CoverTab[80279]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1360
			// _ = "end of CoverTab[80279]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1360
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1360
		// _ = "end of CoverTab[80271]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1360
		_go_fuzz_dep_.CoverTab[80272]++
												if len(binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1361
			_go_fuzz_dep_.CoverTab[80280]++
													if h, _ := stream.Header(); h.Len() > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1362
				_go_fuzz_dep_.CoverTab[80282]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1365
				sh := &binarylog.ServerHeader{
					Header: h,
				}
				for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1368
					_go_fuzz_dep_.CoverTab[80283]++
															binlog.Log(stream.Context(), sh)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1369
					// _ = "end of CoverTab[80283]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1370
				// _ = "end of CoverTab[80282]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1371
				_go_fuzz_dep_.CoverTab[80284]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1371
				// _ = "end of CoverTab[80284]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1371
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1371
			// _ = "end of CoverTab[80280]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1371
			_go_fuzz_dep_.CoverTab[80281]++
													st := &binarylog.ServerTrailer{
				Trailer:	stream.Trailer(),
				Err:		appErr,
			}
			for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1376
				_go_fuzz_dep_.CoverTab[80285]++
														binlog.Log(stream.Context(), st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1377
				// _ = "end of CoverTab[80285]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1378
			// _ = "end of CoverTab[80281]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1379
			_go_fuzz_dep_.CoverTab[80286]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1379
			// _ = "end of CoverTab[80286]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1379
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1379
		// _ = "end of CoverTab[80272]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1379
		_go_fuzz_dep_.CoverTab[80273]++
												return appErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1380
		// _ = "end of CoverTab[80273]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1381
		_go_fuzz_dep_.CoverTab[80287]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1381
		// _ = "end of CoverTab[80287]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1381
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1381
	// _ = "end of CoverTab[80167]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1381
	_go_fuzz_dep_.CoverTab[80168]++
											if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1382
		_go_fuzz_dep_.CoverTab[80288]++
												trInfo.tr.LazyLog(stringer("OK"), false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1383
		// _ = "end of CoverTab[80288]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1384
		_go_fuzz_dep_.CoverTab[80289]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1384
		// _ = "end of CoverTab[80289]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1384
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1384
	// _ = "end of CoverTab[80168]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1384
	_go_fuzz_dep_.CoverTab[80169]++
											opts := &transport.Options{Last: true}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1389
	if stream.SendCompress() != sendCompressorName {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1389
		_go_fuzz_dep_.CoverTab[80290]++
												comp = encoding.GetCompressor(stream.SendCompress())
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1390
		// _ = "end of CoverTab[80290]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1391
		_go_fuzz_dep_.CoverTab[80291]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1391
		// _ = "end of CoverTab[80291]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1391
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1391
	// _ = "end of CoverTab[80169]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1391
	_go_fuzz_dep_.CoverTab[80170]++
											if err := s.sendResponse(t, stream, reply, cp, opts, comp); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1392
		_go_fuzz_dep_.CoverTab[80292]++
												if err == io.EOF {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1393
			_go_fuzz_dep_.CoverTab[80296]++

													return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1395
			// _ = "end of CoverTab[80296]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1396
			_go_fuzz_dep_.CoverTab[80297]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1396
			// _ = "end of CoverTab[80297]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1396
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1396
		// _ = "end of CoverTab[80292]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1396
		_go_fuzz_dep_.CoverTab[80293]++
												if sts, ok := status.FromError(err); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1397
			_go_fuzz_dep_.CoverTab[80298]++
													if e := t.WriteStatus(stream, sts); e != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1398
				_go_fuzz_dep_.CoverTab[80299]++
														channelz.Warningf(logger, s.channelzID, "grpc: Server.processUnaryRPC failed to write status: %v", e)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1399
				// _ = "end of CoverTab[80299]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1400
				_go_fuzz_dep_.CoverTab[80300]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1400
				// _ = "end of CoverTab[80300]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1400
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1400
			// _ = "end of CoverTab[80298]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1401
			_go_fuzz_dep_.CoverTab[80301]++
													switch st := err.(type) {
			case transport.ConnectionError:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1403
				_go_fuzz_dep_.CoverTab[80302]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1403
				// _ = "end of CoverTab[80302]"

			default:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1405
				_go_fuzz_dep_.CoverTab[80303]++
														panic(fmt.Sprintf("grpc: Unexpected error (%T) from sendResponse: %v", st, st))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1406
				// _ = "end of CoverTab[80303]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1407
			// _ = "end of CoverTab[80301]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1408
		// _ = "end of CoverTab[80293]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1408
		_go_fuzz_dep_.CoverTab[80294]++
												if len(binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1409
			_go_fuzz_dep_.CoverTab[80304]++
													h, _ := stream.Header()
													sh := &binarylog.ServerHeader{
				Header: h,
			}
			st := &binarylog.ServerTrailer{
				Trailer:	stream.Trailer(),
				Err:		appErr,
			}
			for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1418
				_go_fuzz_dep_.CoverTab[80305]++
														binlog.Log(stream.Context(), sh)
														binlog.Log(stream.Context(), st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1420
				// _ = "end of CoverTab[80305]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1421
			// _ = "end of CoverTab[80304]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1422
			_go_fuzz_dep_.CoverTab[80306]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1422
			// _ = "end of CoverTab[80306]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1422
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1422
		// _ = "end of CoverTab[80294]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1422
		_go_fuzz_dep_.CoverTab[80295]++
												return err
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1423
		// _ = "end of CoverTab[80295]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1424
		_go_fuzz_dep_.CoverTab[80307]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1424
		// _ = "end of CoverTab[80307]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1424
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1424
	// _ = "end of CoverTab[80170]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1424
	_go_fuzz_dep_.CoverTab[80171]++
											if len(binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1425
		_go_fuzz_dep_.CoverTab[80308]++
												h, _ := stream.Header()
												sh := &binarylog.ServerHeader{
			Header: h,
		}
		sm := &binarylog.ServerMessage{
			Message: reply,
		}
		for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1433
			_go_fuzz_dep_.CoverTab[80309]++
													binlog.Log(stream.Context(), sh)
													binlog.Log(stream.Context(), sm)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1435
			// _ = "end of CoverTab[80309]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1436
		// _ = "end of CoverTab[80308]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1437
		_go_fuzz_dep_.CoverTab[80310]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1437
		// _ = "end of CoverTab[80310]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1437
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1437
	// _ = "end of CoverTab[80171]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1437
	_go_fuzz_dep_.CoverTab[80172]++
											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1438
		_go_fuzz_dep_.CoverTab[80311]++
												t.IncrMsgSent()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1439
		// _ = "end of CoverTab[80311]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1440
		_go_fuzz_dep_.CoverTab[80312]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1440
		// _ = "end of CoverTab[80312]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1440
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1440
	// _ = "end of CoverTab[80172]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1440
	_go_fuzz_dep_.CoverTab[80173]++
											if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1441
		_go_fuzz_dep_.CoverTab[80313]++
												trInfo.tr.LazyLog(&payload{sent: true, msg: reply}, true)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1442
		// _ = "end of CoverTab[80313]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1443
		_go_fuzz_dep_.CoverTab[80314]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1443
		// _ = "end of CoverTab[80314]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1443
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1443
	// _ = "end of CoverTab[80173]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1443
	_go_fuzz_dep_.CoverTab[80174]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1447
	if len(binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1447
		_go_fuzz_dep_.CoverTab[80315]++
												st := &binarylog.ServerTrailer{
			Trailer:	stream.Trailer(),
			Err:		appErr,
		}
		for _, binlog := range binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1452
			_go_fuzz_dep_.CoverTab[80316]++
													binlog.Log(stream.Context(), st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1453
			// _ = "end of CoverTab[80316]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1454
		// _ = "end of CoverTab[80315]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1455
		_go_fuzz_dep_.CoverTab[80317]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1455
		// _ = "end of CoverTab[80317]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1455
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1455
	// _ = "end of CoverTab[80174]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1455
	_go_fuzz_dep_.CoverTab[80175]++
											return t.WriteStatus(stream, statusOK)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1456
	// _ = "end of CoverTab[80175]"
}

// chainStreamServerInterceptors chains all stream server interceptors into one.
func chainStreamServerInterceptors(s *Server) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1460
	_go_fuzz_dep_.CoverTab[80318]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1463
	interceptors := s.opts.chainStreamInts
	if s.opts.streamInt != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1464
		_go_fuzz_dep_.CoverTab[80321]++
												interceptors = append([]StreamServerInterceptor{s.opts.streamInt}, s.opts.chainStreamInts...)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1465
		// _ = "end of CoverTab[80321]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1466
		_go_fuzz_dep_.CoverTab[80322]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1466
		// _ = "end of CoverTab[80322]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1466
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1466
	// _ = "end of CoverTab[80318]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1466
	_go_fuzz_dep_.CoverTab[80319]++

											var chainedInt StreamServerInterceptor
											if len(interceptors) == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1469
		_go_fuzz_dep_.CoverTab[80323]++
												chainedInt = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1470
		// _ = "end of CoverTab[80323]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1471
		_go_fuzz_dep_.CoverTab[80324]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1471
		if len(interceptors) == 1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1471
			_go_fuzz_dep_.CoverTab[80325]++
													chainedInt = interceptors[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1472
			// _ = "end of CoverTab[80325]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1473
			_go_fuzz_dep_.CoverTab[80326]++
													chainedInt = chainStreamInterceptors(interceptors)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1474
			// _ = "end of CoverTab[80326]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1475
		// _ = "end of CoverTab[80324]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1475
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1475
	// _ = "end of CoverTab[80319]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1475
	_go_fuzz_dep_.CoverTab[80320]++

											s.opts.streamInt = chainedInt
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1477
	// _ = "end of CoverTab[80320]"
}

func chainStreamInterceptors(interceptors []StreamServerInterceptor) StreamServerInterceptor {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1480
	_go_fuzz_dep_.CoverTab[80327]++
											return func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1481
		_go_fuzz_dep_.CoverTab[80328]++
												return interceptors[0](srv, ss, info, getChainStreamHandler(interceptors, 0, info, handler))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1482
		// _ = "end of CoverTab[80328]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1483
	// _ = "end of CoverTab[80327]"
}

func getChainStreamHandler(interceptors []StreamServerInterceptor, curr int, info *StreamServerInfo, finalHandler StreamHandler) StreamHandler {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1486
	_go_fuzz_dep_.CoverTab[80329]++
											if curr == len(interceptors)-1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1487
		_go_fuzz_dep_.CoverTab[80331]++
												return finalHandler
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1488
		// _ = "end of CoverTab[80331]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1489
		_go_fuzz_dep_.CoverTab[80332]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1489
		// _ = "end of CoverTab[80332]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1489
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1489
	// _ = "end of CoverTab[80329]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1489
	_go_fuzz_dep_.CoverTab[80330]++
											return func(srv interface{}, stream ServerStream) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1490
		_go_fuzz_dep_.CoverTab[80333]++
												return interceptors[curr+1](srv, stream, info, getChainStreamHandler(interceptors, curr+1, info, finalHandler))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1491
		// _ = "end of CoverTab[80333]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1492
	// _ = "end of CoverTab[80330]"
}

func (s *Server) processStreamingRPC(t transport.ServerTransport, stream *transport.Stream, info *serviceInfo, sd *StreamDesc, trInfo *traceInfo) (err error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1495
	_go_fuzz_dep_.CoverTab[80334]++
											if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1496
		_go_fuzz_dep_.CoverTab[80350]++
												s.incrCallsStarted()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1497
		// _ = "end of CoverTab[80350]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1498
		_go_fuzz_dep_.CoverTab[80351]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1498
		// _ = "end of CoverTab[80351]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1498
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1498
	// _ = "end of CoverTab[80334]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1498
	_go_fuzz_dep_.CoverTab[80335]++
											shs := s.opts.statsHandlers
											var statsBegin *stats.Begin
											if len(shs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1501
		_go_fuzz_dep_.CoverTab[80352]++
												beginTime := time.Now()
												statsBegin = &stats.Begin{
			BeginTime:	beginTime,
			IsClientStream:	sd.ClientStreams,
			IsServerStream:	sd.ServerStreams,
		}
		for _, sh := range shs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1508
			_go_fuzz_dep_.CoverTab[80353]++
													sh.HandleRPC(stream.Context(), statsBegin)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1509
			// _ = "end of CoverTab[80353]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1510
		// _ = "end of CoverTab[80352]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1511
		_go_fuzz_dep_.CoverTab[80354]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1511
		// _ = "end of CoverTab[80354]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1511
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1511
	// _ = "end of CoverTab[80335]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1511
	_go_fuzz_dep_.CoverTab[80336]++
											ctx := NewContextWithServerTransportStream(stream.Context(), stream)
											ss := &serverStream{
		ctx:			ctx,
		t:			t,
		s:			stream,
		p:			&parser{r: stream},
		codec:			s.getCodec(stream.ContentSubtype()),
		maxReceiveMessageSize:	s.opts.maxReceiveMessageSize,
		maxSendMessageSize:	s.opts.maxSendMessageSize,
		trInfo:			trInfo,
		statsHandler:		shs,
	}

	if len(shs) != 0 || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		_go_fuzz_dep_.CoverTab[80355]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		return trInfo != nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		// _ = "end of CoverTab[80355]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		_go_fuzz_dep_.CoverTab[80356]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		return channelz.IsOn()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		// _ = "end of CoverTab[80356]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1525
		_go_fuzz_dep_.CoverTab[80357]++

												defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1527
			_go_fuzz_dep_.CoverTab[80358]++
													if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1528
				_go_fuzz_dep_.CoverTab[80361]++
														ss.mu.Lock()
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1530
					_go_fuzz_dep_.CoverTab[80363]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1530
					return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1530
					// _ = "end of CoverTab[80363]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1530
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1530
					_go_fuzz_dep_.CoverTab[80364]++
															ss.trInfo.tr.LazyLog(&fmtStringer{"%v", []interface{}{err}}, true)
															ss.trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1532
					// _ = "end of CoverTab[80364]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1533
					_go_fuzz_dep_.CoverTab[80365]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1533
					// _ = "end of CoverTab[80365]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1533
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1533
				// _ = "end of CoverTab[80361]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1533
				_go_fuzz_dep_.CoverTab[80362]++
														ss.trInfo.tr.Finish()
														ss.trInfo.tr = nil
														ss.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1536
				// _ = "end of CoverTab[80362]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1537
				_go_fuzz_dep_.CoverTab[80366]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1537
				// _ = "end of CoverTab[80366]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1537
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1537
			// _ = "end of CoverTab[80358]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1537
			_go_fuzz_dep_.CoverTab[80359]++

													if len(shs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1539
				_go_fuzz_dep_.CoverTab[80367]++
														end := &stats.End{
					BeginTime:	statsBegin.BeginTime,
					EndTime:	time.Now(),
				}
				if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1544
					_go_fuzz_dep_.CoverTab[80369]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1544
					return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1544
					// _ = "end of CoverTab[80369]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1544
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1544
					_go_fuzz_dep_.CoverTab[80370]++
															end.Error = toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1545
					// _ = "end of CoverTab[80370]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1546
					_go_fuzz_dep_.CoverTab[80371]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1546
					// _ = "end of CoverTab[80371]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1546
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1546
				// _ = "end of CoverTab[80367]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1546
				_go_fuzz_dep_.CoverTab[80368]++
														for _, sh := range shs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1547
					_go_fuzz_dep_.CoverTab[80372]++
															sh.HandleRPC(stream.Context(), end)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1548
					// _ = "end of CoverTab[80372]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1549
				// _ = "end of CoverTab[80368]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1550
				_go_fuzz_dep_.CoverTab[80373]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1550
				// _ = "end of CoverTab[80373]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1550
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1550
			// _ = "end of CoverTab[80359]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1550
			_go_fuzz_dep_.CoverTab[80360]++

													if channelz.IsOn() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1552
				_go_fuzz_dep_.CoverTab[80374]++
														if err != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1553
					_go_fuzz_dep_.CoverTab[80375]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1553
					return err != io.EOF
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1553
					// _ = "end of CoverTab[80375]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1553
				}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1553
					_go_fuzz_dep_.CoverTab[80376]++
															s.incrCallsFailed()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1554
					// _ = "end of CoverTab[80376]"
				} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1555
					_go_fuzz_dep_.CoverTab[80377]++
															s.incrCallsSucceeded()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1556
					// _ = "end of CoverTab[80377]"
				}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1557
				// _ = "end of CoverTab[80374]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1558
				_go_fuzz_dep_.CoverTab[80378]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1558
				// _ = "end of CoverTab[80378]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1558
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1558
			// _ = "end of CoverTab[80360]"
		}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1559
		// _ = "end of CoverTab[80357]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1560
		_go_fuzz_dep_.CoverTab[80379]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1560
		// _ = "end of CoverTab[80379]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1560
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1560
	// _ = "end of CoverTab[80336]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1560
	_go_fuzz_dep_.CoverTab[80337]++

											if ml := binarylog.GetMethodLogger(stream.Method()); ml != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1562
		_go_fuzz_dep_.CoverTab[80380]++
												ss.binlogs = append(ss.binlogs, ml)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1563
		// _ = "end of CoverTab[80380]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1564
		_go_fuzz_dep_.CoverTab[80381]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1564
		// _ = "end of CoverTab[80381]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1564
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1564
	// _ = "end of CoverTab[80337]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1564
	_go_fuzz_dep_.CoverTab[80338]++
											if s.opts.binaryLogger != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1565
		_go_fuzz_dep_.CoverTab[80382]++
												if ml := s.opts.binaryLogger.GetMethodLogger(stream.Method()); ml != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1566
			_go_fuzz_dep_.CoverTab[80383]++
													ss.binlogs = append(ss.binlogs, ml)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1567
			// _ = "end of CoverTab[80383]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1568
			_go_fuzz_dep_.CoverTab[80384]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1568
			// _ = "end of CoverTab[80384]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1568
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1568
		// _ = "end of CoverTab[80382]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1569
		_go_fuzz_dep_.CoverTab[80385]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1569
		// _ = "end of CoverTab[80385]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1569
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1569
	// _ = "end of CoverTab[80338]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1569
	_go_fuzz_dep_.CoverTab[80339]++
											if len(ss.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1570
		_go_fuzz_dep_.CoverTab[80386]++
												md, _ := metadata.FromIncomingContext(ctx)
												logEntry := &binarylog.ClientHeader{
			Header:		md,
			MethodName:	stream.Method(),
			PeerAddr:	nil,
		}
		if deadline, ok := ctx.Deadline(); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1577
			_go_fuzz_dep_.CoverTab[80390]++
													logEntry.Timeout = time.Until(deadline)
													if logEntry.Timeout < 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1579
				_go_fuzz_dep_.CoverTab[80391]++
														logEntry.Timeout = 0
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1580
				// _ = "end of CoverTab[80391]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1581
				_go_fuzz_dep_.CoverTab[80392]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1581
				// _ = "end of CoverTab[80392]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1581
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1581
			// _ = "end of CoverTab[80390]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1582
			_go_fuzz_dep_.CoverTab[80393]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1582
			// _ = "end of CoverTab[80393]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1582
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1582
		// _ = "end of CoverTab[80386]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1582
		_go_fuzz_dep_.CoverTab[80387]++
												if a := md[":authority"]; len(a) > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1583
			_go_fuzz_dep_.CoverTab[80394]++
													logEntry.Authority = a[0]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1584
			// _ = "end of CoverTab[80394]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1585
			_go_fuzz_dep_.CoverTab[80395]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1585
			// _ = "end of CoverTab[80395]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1585
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1585
		// _ = "end of CoverTab[80387]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1585
		_go_fuzz_dep_.CoverTab[80388]++
												if peer, ok := peer.FromContext(ss.Context()); ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1586
			_go_fuzz_dep_.CoverTab[80396]++
													logEntry.PeerAddr = peer.Addr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1587
			// _ = "end of CoverTab[80396]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1588
			_go_fuzz_dep_.CoverTab[80397]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1588
			// _ = "end of CoverTab[80397]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1588
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1588
		// _ = "end of CoverTab[80388]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1588
		_go_fuzz_dep_.CoverTab[80389]++
												for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1589
			_go_fuzz_dep_.CoverTab[80398]++
													binlog.Log(stream.Context(), logEntry)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1590
			// _ = "end of CoverTab[80398]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1591
		// _ = "end of CoverTab[80389]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1592
		_go_fuzz_dep_.CoverTab[80399]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1592
		// _ = "end of CoverTab[80399]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1592
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1592
	// _ = "end of CoverTab[80339]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1592
	_go_fuzz_dep_.CoverTab[80340]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1596
	if rc := stream.RecvCompress(); s.opts.dc != nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1596
		_go_fuzz_dep_.CoverTab[80400]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1596
		return s.opts.dc.Type() == rc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1596
		// _ = "end of CoverTab[80400]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1596
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1596
		_go_fuzz_dep_.CoverTab[80401]++
												ss.dc = s.opts.dc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1597
		// _ = "end of CoverTab[80401]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
		_go_fuzz_dep_.CoverTab[80402]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
		if rc != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
			_go_fuzz_dep_.CoverTab[80403]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
			return rc != encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
			// _ = "end of CoverTab[80403]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1598
			_go_fuzz_dep_.CoverTab[80404]++
													ss.decomp = encoding.GetCompressor(rc)
													if ss.decomp == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1600
				_go_fuzz_dep_.CoverTab[80405]++
														st := status.Newf(codes.Unimplemented, "grpc: Decompressor is not installed for grpc-encoding %q", rc)
														t.WriteStatus(ss.s, st)
														return st.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1603
				// _ = "end of CoverTab[80405]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1604
				_go_fuzz_dep_.CoverTab[80406]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1604
				// _ = "end of CoverTab[80406]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1604
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1604
			// _ = "end of CoverTab[80404]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
			_go_fuzz_dep_.CoverTab[80407]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
			// _ = "end of CoverTab[80407]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
		// _ = "end of CoverTab[80402]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
	// _ = "end of CoverTab[80340]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1605
	_go_fuzz_dep_.CoverTab[80341]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1611
	if s.opts.cp != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1611
		_go_fuzz_dep_.CoverTab[80408]++
												ss.cp = s.opts.cp
												ss.sendCompressorName = s.opts.cp.Type()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1613
		// _ = "end of CoverTab[80408]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
		_go_fuzz_dep_.CoverTab[80409]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
		if rc := stream.RecvCompress(); rc != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
			_go_fuzz_dep_.CoverTab[80410]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
			return rc != encoding.Identity
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
			// _ = "end of CoverTab[80410]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
		}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1614
			_go_fuzz_dep_.CoverTab[80411]++

													ss.comp = encoding.GetCompressor(rc)
													if ss.comp != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1617
				_go_fuzz_dep_.CoverTab[80412]++
														ss.sendCompressorName = rc
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1618
				// _ = "end of CoverTab[80412]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1619
				_go_fuzz_dep_.CoverTab[80413]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1619
				// _ = "end of CoverTab[80413]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1619
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1619
			// _ = "end of CoverTab[80411]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
			_go_fuzz_dep_.CoverTab[80414]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
			// _ = "end of CoverTab[80414]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
		// _ = "end of CoverTab[80409]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
	// _ = "end of CoverTab[80341]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1620
	_go_fuzz_dep_.CoverTab[80342]++

											if ss.sendCompressorName != "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1622
		_go_fuzz_dep_.CoverTab[80415]++
												if err := stream.SetSendCompress(ss.sendCompressorName); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1623
			_go_fuzz_dep_.CoverTab[80416]++
													return status.Errorf(codes.Internal, "grpc: failed to set send compressor: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1624
			// _ = "end of CoverTab[80416]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1625
			_go_fuzz_dep_.CoverTab[80417]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1625
			// _ = "end of CoverTab[80417]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1625
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1625
		// _ = "end of CoverTab[80415]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1626
		_go_fuzz_dep_.CoverTab[80418]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1626
		// _ = "end of CoverTab[80418]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1626
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1626
	// _ = "end of CoverTab[80342]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1626
	_go_fuzz_dep_.CoverTab[80343]++

											ss.ctx = newContextWithRPCInfo(ss.ctx, false, ss.codec, ss.cp, ss.comp)

											if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1630
		_go_fuzz_dep_.CoverTab[80419]++
												trInfo.tr.LazyLog(&trInfo.firstLine, false)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1631
		// _ = "end of CoverTab[80419]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1632
		_go_fuzz_dep_.CoverTab[80420]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1632
		// _ = "end of CoverTab[80420]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1632
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1632
	// _ = "end of CoverTab[80343]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1632
	_go_fuzz_dep_.CoverTab[80344]++
											var appErr error
											var server interface{}
											if info != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1635
		_go_fuzz_dep_.CoverTab[80421]++
												server = info.serviceImpl
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1636
		// _ = "end of CoverTab[80421]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1637
		_go_fuzz_dep_.CoverTab[80422]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1637
		// _ = "end of CoverTab[80422]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1637
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1637
	// _ = "end of CoverTab[80344]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1637
	_go_fuzz_dep_.CoverTab[80345]++
											if s.opts.streamInt == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1638
		_go_fuzz_dep_.CoverTab[80423]++
												appErr = sd.Handler(server, ss)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1639
		// _ = "end of CoverTab[80423]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1640
		_go_fuzz_dep_.CoverTab[80424]++
												info := &StreamServerInfo{
			FullMethod:	stream.Method(),
			IsClientStream:	sd.ClientStreams,
			IsServerStream:	sd.ServerStreams,
		}
												appErr = s.opts.streamInt(server, ss, info, sd.Handler)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1646
		// _ = "end of CoverTab[80424]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1647
	// _ = "end of CoverTab[80345]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1647
	_go_fuzz_dep_.CoverTab[80346]++
											if appErr != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1648
		_go_fuzz_dep_.CoverTab[80425]++
												appStatus, ok := status.FromError(appErr)
												if !ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1650
			_go_fuzz_dep_.CoverTab[80429]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1653
			appStatus = status.FromContextError(appErr)
													appErr = appStatus.Err()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1654
			// _ = "end of CoverTab[80429]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1655
			_go_fuzz_dep_.CoverTab[80430]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1655
			// _ = "end of CoverTab[80430]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1655
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1655
		// _ = "end of CoverTab[80425]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1655
		_go_fuzz_dep_.CoverTab[80426]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1656
			_go_fuzz_dep_.CoverTab[80431]++
													ss.mu.Lock()
													ss.trInfo.tr.LazyLog(stringer(appStatus.Message()), true)
													ss.trInfo.tr.SetError()
													ss.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1660
			// _ = "end of CoverTab[80431]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1661
			_go_fuzz_dep_.CoverTab[80432]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1661
			// _ = "end of CoverTab[80432]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1661
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1661
		// _ = "end of CoverTab[80426]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1661
		_go_fuzz_dep_.CoverTab[80427]++
												if len(ss.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1662
			_go_fuzz_dep_.CoverTab[80433]++
													st := &binarylog.ServerTrailer{
				Trailer:	ss.s.Trailer(),
				Err:		appErr,
			}
			for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1667
				_go_fuzz_dep_.CoverTab[80434]++
														binlog.Log(stream.Context(), st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1668
				// _ = "end of CoverTab[80434]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1669
			// _ = "end of CoverTab[80433]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1670
			_go_fuzz_dep_.CoverTab[80435]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1670
			// _ = "end of CoverTab[80435]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1670
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1670
		// _ = "end of CoverTab[80427]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1670
		_go_fuzz_dep_.CoverTab[80428]++
												t.WriteStatus(ss.s, appStatus)

												return appErr
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1673
		// _ = "end of CoverTab[80428]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1674
		_go_fuzz_dep_.CoverTab[80436]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1674
		// _ = "end of CoverTab[80436]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1674
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1674
	// _ = "end of CoverTab[80346]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1674
	_go_fuzz_dep_.CoverTab[80347]++
											if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1675
		_go_fuzz_dep_.CoverTab[80437]++
												ss.mu.Lock()
												ss.trInfo.tr.LazyLog(stringer("OK"), false)
												ss.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1678
		// _ = "end of CoverTab[80437]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1679
		_go_fuzz_dep_.CoverTab[80438]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1679
		// _ = "end of CoverTab[80438]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1679
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1679
	// _ = "end of CoverTab[80347]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1679
	_go_fuzz_dep_.CoverTab[80348]++
											if len(ss.binlogs) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1680
		_go_fuzz_dep_.CoverTab[80439]++
												st := &binarylog.ServerTrailer{
			Trailer:	ss.s.Trailer(),
			Err:		appErr,
		}
		for _, binlog := range ss.binlogs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1685
			_go_fuzz_dep_.CoverTab[80440]++
													binlog.Log(stream.Context(), st)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1686
			// _ = "end of CoverTab[80440]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1687
		// _ = "end of CoverTab[80439]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1688
		_go_fuzz_dep_.CoverTab[80441]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1688
		// _ = "end of CoverTab[80441]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1688
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1688
	// _ = "end of CoverTab[80348]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1688
	_go_fuzz_dep_.CoverTab[80349]++
											return t.WriteStatus(ss.s, statusOK)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1689
	// _ = "end of CoverTab[80349]"
}

func (s *Server) handleStream(t transport.ServerTransport, stream *transport.Stream, trInfo *traceInfo) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1692
	_go_fuzz_dep_.CoverTab[80442]++
											sm := stream.Method()
											if sm != "" && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1694
		_go_fuzz_dep_.CoverTab[80450]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1694
		return sm[0] == '/'
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1694
		// _ = "end of CoverTab[80450]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1694
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1694
		_go_fuzz_dep_.CoverTab[80451]++
												sm = sm[1:]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1695
		// _ = "end of CoverTab[80451]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1696
		_go_fuzz_dep_.CoverTab[80452]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1696
		// _ = "end of CoverTab[80452]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1696
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1696
	// _ = "end of CoverTab[80442]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1696
	_go_fuzz_dep_.CoverTab[80443]++
											pos := strings.LastIndex(sm, "/")
											if pos == -1 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1698
		_go_fuzz_dep_.CoverTab[80453]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1699
			_go_fuzz_dep_.CoverTab[80457]++
													trInfo.tr.LazyLog(&fmtStringer{"Malformed method name %q", []interface{}{sm}}, true)
													trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1701
			// _ = "end of CoverTab[80457]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1702
			_go_fuzz_dep_.CoverTab[80458]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1702
			// _ = "end of CoverTab[80458]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1702
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1702
		// _ = "end of CoverTab[80453]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1702
		_go_fuzz_dep_.CoverTab[80454]++
												errDesc := fmt.Sprintf("malformed method name: %q", stream.Method())
												if err := t.WriteStatus(stream, status.New(codes.Unimplemented, errDesc)); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1704
			_go_fuzz_dep_.CoverTab[80459]++
													if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1705
				_go_fuzz_dep_.CoverTab[80461]++
														trInfo.tr.LazyLog(&fmtStringer{"%v", []interface{}{err}}, true)
														trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1707
				// _ = "end of CoverTab[80461]"
			} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1708
				_go_fuzz_dep_.CoverTab[80462]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1708
				// _ = "end of CoverTab[80462]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1708
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1708
			// _ = "end of CoverTab[80459]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1708
			_go_fuzz_dep_.CoverTab[80460]++
													channelz.Warningf(logger, s.channelzID, "grpc: Server.handleStream failed to write status: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1709
			// _ = "end of CoverTab[80460]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1710
			_go_fuzz_dep_.CoverTab[80463]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1710
			// _ = "end of CoverTab[80463]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1710
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1710
		// _ = "end of CoverTab[80454]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1710
		_go_fuzz_dep_.CoverTab[80455]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1711
			_go_fuzz_dep_.CoverTab[80464]++
													trInfo.tr.Finish()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1712
			// _ = "end of CoverTab[80464]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1713
			_go_fuzz_dep_.CoverTab[80465]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1713
			// _ = "end of CoverTab[80465]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1713
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1713
		// _ = "end of CoverTab[80455]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1713
		_go_fuzz_dep_.CoverTab[80456]++
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1714
		// _ = "end of CoverTab[80456]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1715
		_go_fuzz_dep_.CoverTab[80466]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1715
		// _ = "end of CoverTab[80466]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1715
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1715
	// _ = "end of CoverTab[80443]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1715
	_go_fuzz_dep_.CoverTab[80444]++
											service := sm[:pos]
											method := sm[pos+1:]

											srv, knownService := s.services[service]
											if knownService {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1720
		_go_fuzz_dep_.CoverTab[80467]++
												if md, ok := srv.methods[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1721
			_go_fuzz_dep_.CoverTab[80469]++
													s.processUnaryRPC(t, stream, srv, md, trInfo)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1723
			// _ = "end of CoverTab[80469]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1724
			_go_fuzz_dep_.CoverTab[80470]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1724
			// _ = "end of CoverTab[80470]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1724
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1724
		// _ = "end of CoverTab[80467]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1724
		_go_fuzz_dep_.CoverTab[80468]++
												if sd, ok := srv.streams[method]; ok {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1725
			_go_fuzz_dep_.CoverTab[80471]++
													s.processStreamingRPC(t, stream, srv, sd, trInfo)
													return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1727
			// _ = "end of CoverTab[80471]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1728
			_go_fuzz_dep_.CoverTab[80472]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1728
			// _ = "end of CoverTab[80472]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1728
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1728
		// _ = "end of CoverTab[80468]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1729
		_go_fuzz_dep_.CoverTab[80473]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1729
		// _ = "end of CoverTab[80473]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1729
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1729
	// _ = "end of CoverTab[80444]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1729
	_go_fuzz_dep_.CoverTab[80445]++

											if unknownDesc := s.opts.unknownStreamDesc; unknownDesc != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1731
		_go_fuzz_dep_.CoverTab[80474]++
												s.processStreamingRPC(t, stream, nil, unknownDesc, trInfo)
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1733
		// _ = "end of CoverTab[80474]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1734
		_go_fuzz_dep_.CoverTab[80475]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1734
		// _ = "end of CoverTab[80475]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1734
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1734
	// _ = "end of CoverTab[80445]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1734
	_go_fuzz_dep_.CoverTab[80446]++
											var errDesc string
											if !knownService {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1736
		_go_fuzz_dep_.CoverTab[80476]++
												errDesc = fmt.Sprintf("unknown service %v", service)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1737
		// _ = "end of CoverTab[80476]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1738
		_go_fuzz_dep_.CoverTab[80477]++
												errDesc = fmt.Sprintf("unknown method %v for service %v", method, service)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1739
		// _ = "end of CoverTab[80477]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1740
	// _ = "end of CoverTab[80446]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1740
	_go_fuzz_dep_.CoverTab[80447]++
											if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1741
		_go_fuzz_dep_.CoverTab[80478]++
												trInfo.tr.LazyPrintf("%s", errDesc)
												trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1743
		// _ = "end of CoverTab[80478]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1744
		_go_fuzz_dep_.CoverTab[80479]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1744
		// _ = "end of CoverTab[80479]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1744
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1744
	// _ = "end of CoverTab[80447]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1744
	_go_fuzz_dep_.CoverTab[80448]++
											if err := t.WriteStatus(stream, status.New(codes.Unimplemented, errDesc)); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1745
		_go_fuzz_dep_.CoverTab[80480]++
												if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1746
			_go_fuzz_dep_.CoverTab[80482]++
													trInfo.tr.LazyLog(&fmtStringer{"%v", []interface{}{err}}, true)
													trInfo.tr.SetError()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1748
			// _ = "end of CoverTab[80482]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1749
			_go_fuzz_dep_.CoverTab[80483]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1749
			// _ = "end of CoverTab[80483]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1749
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1749
		// _ = "end of CoverTab[80480]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1749
		_go_fuzz_dep_.CoverTab[80481]++
												channelz.Warningf(logger, s.channelzID, "grpc: Server.handleStream failed to write status: %v", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1750
		// _ = "end of CoverTab[80481]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1751
		_go_fuzz_dep_.CoverTab[80484]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1751
		// _ = "end of CoverTab[80484]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1751
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1751
	// _ = "end of CoverTab[80448]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1751
	_go_fuzz_dep_.CoverTab[80449]++
											if trInfo != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1752
		_go_fuzz_dep_.CoverTab[80485]++
												trInfo.tr.Finish()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1753
		// _ = "end of CoverTab[80485]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1754
		_go_fuzz_dep_.CoverTab[80486]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1754
		// _ = "end of CoverTab[80486]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1754
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1754
	// _ = "end of CoverTab[80449]"
}

// The key to save ServerTransportStream in the context.
type streamKey struct{}

// NewContextWithServerTransportStream creates a new context from ctx and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1760
// attaches stream to it.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1760
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1760
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1760
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1760
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1760
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1767
func NewContextWithServerTransportStream(ctx context.Context, stream ServerTransportStream) context.Context {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1767
	_go_fuzz_dep_.CoverTab[80487]++
											return context.WithValue(ctx, streamKey{}, stream)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1768
	// _ = "end of CoverTab[80487]"
}

// ServerTransportStream is a minimal interface that a transport stream must
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// implement. This can be used to mock an actual transport stream for tests of
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// handler code that use, for example, grpc.SetHeader (which requires some
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// stream to be in context).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// See also NewContextWithServerTransportStream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1771
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1782
type ServerTransportStream interface {
	Method() string
	SetHeader(md metadata.MD) error
	SendHeader(md metadata.MD) error
	SetTrailer(md metadata.MD) error
}

// ServerTransportStreamFromContext returns the ServerTransportStream saved in
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
// ctx. Returns nil if the given context has no stream associated with it
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
// (which implies it is not an RPC invocation context).
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1789
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1797
func ServerTransportStreamFromContext(ctx context.Context) ServerTransportStream {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1797
	_go_fuzz_dep_.CoverTab[80488]++
											s, _ := ctx.Value(streamKey{}).(ServerTransportStream)
											return s
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1799
	// _ = "end of CoverTab[80488]"
}

// Stop stops the gRPC server. It immediately closes all open
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1802
// connections and listeners.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1802
// It cancels all active RPCs on the server side and the corresponding
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1802
// pending RPCs on the client side will get notified by connection
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1802
// errors.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1807
func (s *Server) Stop() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1807
	_go_fuzz_dep_.CoverTab[80489]++
											s.quit.Fire()

											defer func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1810
		_go_fuzz_dep_.CoverTab[80496]++
												s.serveWG.Wait()
												s.done.Fire()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1812
		// _ = "end of CoverTab[80496]"
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1813
	// _ = "end of CoverTab[80489]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1813
	_go_fuzz_dep_.CoverTab[80490]++

											s.channelzRemoveOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1815
		_go_fuzz_dep_.CoverTab[80497]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1815
		channelz.RemoveEntry(s.channelzID)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1815
		// _ = "end of CoverTab[80497]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1815
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1815
	// _ = "end of CoverTab[80490]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1815
	_go_fuzz_dep_.CoverTab[80491]++

											s.mu.Lock()
											listeners := s.lis
											s.lis = nil
											conns := s.conns
											s.conns = nil

											s.cv.Broadcast()
											s.mu.Unlock()

											for lis := range listeners {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1826
		_go_fuzz_dep_.CoverTab[80498]++
												lis.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1827
		// _ = "end of CoverTab[80498]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1828
	// _ = "end of CoverTab[80491]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1828
	_go_fuzz_dep_.CoverTab[80492]++
											for _, cs := range conns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1829
		_go_fuzz_dep_.CoverTab[80499]++
												for st := range cs {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1830
			_go_fuzz_dep_.CoverTab[80500]++
													st.Close(errors.New("Server.Stop called"))
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1831
			// _ = "end of CoverTab[80500]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1832
		// _ = "end of CoverTab[80499]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1833
	// _ = "end of CoverTab[80492]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1833
	_go_fuzz_dep_.CoverTab[80493]++
											if s.opts.numServerWorkers > 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1834
		_go_fuzz_dep_.CoverTab[80501]++
												s.stopServerWorkers()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1835
		// _ = "end of CoverTab[80501]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1836
		_go_fuzz_dep_.CoverTab[80502]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1836
		// _ = "end of CoverTab[80502]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1836
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1836
	// _ = "end of CoverTab[80493]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1836
	_go_fuzz_dep_.CoverTab[80494]++

											s.mu.Lock()
											if s.events != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1839
		_go_fuzz_dep_.CoverTab[80503]++
												s.events.Finish()
												s.events = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1841
		// _ = "end of CoverTab[80503]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1842
		_go_fuzz_dep_.CoverTab[80504]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1842
		// _ = "end of CoverTab[80504]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1842
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1842
	// _ = "end of CoverTab[80494]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1842
	_go_fuzz_dep_.CoverTab[80495]++
											s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1843
	// _ = "end of CoverTab[80495]"
}

// GracefulStop stops the gRPC server gracefully. It stops the server from
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1846
// accepting new connections and RPCs and blocks until all the pending RPCs are
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1846
// finished.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1849
func (s *Server) GracefulStop() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1849
	_go_fuzz_dep_.CoverTab[80505]++
											s.quit.Fire()
											defer s.done.Fire()

											s.channelzRemoveOnce.Do(func() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1853
		_go_fuzz_dep_.CoverTab[80512]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1853
		channelz.RemoveEntry(s.channelzID)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1853
		// _ = "end of CoverTab[80512]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1853
	})
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1853
	// _ = "end of CoverTab[80505]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1853
	_go_fuzz_dep_.CoverTab[80506]++
											s.mu.Lock()
											if s.conns == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1855
		_go_fuzz_dep_.CoverTab[80513]++
												s.mu.Unlock()
												return
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1857
		// _ = "end of CoverTab[80513]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1858
		_go_fuzz_dep_.CoverTab[80514]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1858
		// _ = "end of CoverTab[80514]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1858
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1858
	// _ = "end of CoverTab[80506]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1858
	_go_fuzz_dep_.CoverTab[80507]++

											for lis := range s.lis {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1860
		_go_fuzz_dep_.CoverTab[80515]++
												lis.Close()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1861
		// _ = "end of CoverTab[80515]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1862
	// _ = "end of CoverTab[80507]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1862
	_go_fuzz_dep_.CoverTab[80508]++
											s.lis = nil
											if !s.drain {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1864
		_go_fuzz_dep_.CoverTab[80516]++
												for _, conns := range s.conns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1865
			_go_fuzz_dep_.CoverTab[80518]++
													for st := range conns {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1866
				_go_fuzz_dep_.CoverTab[80519]++
														st.Drain()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1867
				// _ = "end of CoverTab[80519]"
			}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1868
			// _ = "end of CoverTab[80518]"
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1869
		// _ = "end of CoverTab[80516]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1869
		_go_fuzz_dep_.CoverTab[80517]++
												s.drain = true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1870
		// _ = "end of CoverTab[80517]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1871
		_go_fuzz_dep_.CoverTab[80520]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1871
		// _ = "end of CoverTab[80520]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1871
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1871
	// _ = "end of CoverTab[80508]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1871
	_go_fuzz_dep_.CoverTab[80509]++

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1875
	s.mu.Unlock()
	s.serveWG.Wait()
	s.mu.Lock()

	for len(s.conns) != 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1879
		_go_fuzz_dep_.CoverTab[80521]++
												s.cv.Wait()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1880
		// _ = "end of CoverTab[80521]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1881
	// _ = "end of CoverTab[80509]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1881
	_go_fuzz_dep_.CoverTab[80510]++
											s.conns = nil
											if s.events != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1883
		_go_fuzz_dep_.CoverTab[80522]++
												s.events.Finish()
												s.events = nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1885
		// _ = "end of CoverTab[80522]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1886
		_go_fuzz_dep_.CoverTab[80523]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1886
		// _ = "end of CoverTab[80523]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1886
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1886
	// _ = "end of CoverTab[80510]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1886
	_go_fuzz_dep_.CoverTab[80511]++
											s.mu.Unlock()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1887
	// _ = "end of CoverTab[80511]"
}

// contentSubtype must be lowercase
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1890
// cannot return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1892
func (s *Server) getCodec(contentSubtype string) baseCodec {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1892
	_go_fuzz_dep_.CoverTab[80524]++
											if s.opts.codec != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1893
		_go_fuzz_dep_.CoverTab[80528]++
												return s.opts.codec
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1894
		// _ = "end of CoverTab[80528]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1895
		_go_fuzz_dep_.CoverTab[80529]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1895
		// _ = "end of CoverTab[80529]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1895
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1895
	// _ = "end of CoverTab[80524]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1895
	_go_fuzz_dep_.CoverTab[80525]++
											if contentSubtype == "" {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1896
		_go_fuzz_dep_.CoverTab[80530]++
												return encoding.GetCodec(proto.Name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1897
		// _ = "end of CoverTab[80530]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1898
		_go_fuzz_dep_.CoverTab[80531]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1898
		// _ = "end of CoverTab[80531]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1898
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1898
	// _ = "end of CoverTab[80525]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1898
	_go_fuzz_dep_.CoverTab[80526]++
											codec := encoding.GetCodec(contentSubtype)
											if codec == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1900
		_go_fuzz_dep_.CoverTab[80532]++
												return encoding.GetCodec(proto.Name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1901
		// _ = "end of CoverTab[80532]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1902
		_go_fuzz_dep_.CoverTab[80533]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1902
		// _ = "end of CoverTab[80533]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1902
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1902
	// _ = "end of CoverTab[80526]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1902
	_go_fuzz_dep_.CoverTab[80527]++
											return codec
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1903
	// _ = "end of CoverTab[80527]"
}

// SetHeader sets the header metadata to be sent from the server to the client.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// The context provided must be the context passed to the server's handler.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// Streaming RPCs should prefer the SetHeader method of the ServerStream.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// When called multiple times, all the provided metadata will be merged.  All
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// the metadata will be sent out when one of the following happens:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//   - grpc.SendHeader is called, or for streaming handlers, stream.SendHeader.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//   - The first response message is sent.  For unary handlers, this occurs when
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//     the handler returns; for streaming handlers, this can happen when stream's
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//     SendMsg method is called.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//   - An RPC status is sent out (error or success).  This occurs when the handler
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//     returns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// SetHeader will fail if called after any of the events above.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// The error returned is compatible with the status package.  However, the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// status code will often not match the RPC status as seen by the client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1906
// application, and therefore, should not be relied upon for this purpose.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1926
func SetHeader(ctx context.Context, md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1926
	_go_fuzz_dep_.CoverTab[80534]++
											if md.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1927
		_go_fuzz_dep_.CoverTab[80537]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1928
		// _ = "end of CoverTab[80537]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1929
		_go_fuzz_dep_.CoverTab[80538]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1929
		// _ = "end of CoverTab[80538]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1929
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1929
	// _ = "end of CoverTab[80534]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1929
	_go_fuzz_dep_.CoverTab[80535]++
											stream := ServerTransportStreamFromContext(ctx)
											if stream == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1931
		_go_fuzz_dep_.CoverTab[80539]++
												return status.Errorf(codes.Internal, "grpc: failed to fetch the stream from the context %v", ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1932
		// _ = "end of CoverTab[80539]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1933
		_go_fuzz_dep_.CoverTab[80540]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1933
		// _ = "end of CoverTab[80540]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1933
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1933
	// _ = "end of CoverTab[80535]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1933
	_go_fuzz_dep_.CoverTab[80536]++
											return stream.SetHeader(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1934
	// _ = "end of CoverTab[80536]"
}

// SendHeader sends header metadata. It may be called at most once, and may not
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
// be called after any event that causes headers to be sent (see SetHeader for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
// a complete list).  The provided md and headers set by SetHeader() will be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
// sent.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
// The error returned is compatible with the status package.  However, the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
// status code will often not match the RPC status as seen by the client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1937
// application, and therefore, should not be relied upon for this purpose.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1945
func SendHeader(ctx context.Context, md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1945
	_go_fuzz_dep_.CoverTab[80541]++
											stream := ServerTransportStreamFromContext(ctx)
											if stream == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1947
		_go_fuzz_dep_.CoverTab[80544]++
												return status.Errorf(codes.Internal, "grpc: failed to fetch the stream from the context %v", ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1948
		// _ = "end of CoverTab[80544]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1949
		_go_fuzz_dep_.CoverTab[80545]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1949
		// _ = "end of CoverTab[80545]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1949
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1949
	// _ = "end of CoverTab[80541]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1949
	_go_fuzz_dep_.CoverTab[80542]++
											if err := stream.SendHeader(md); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1950
		_go_fuzz_dep_.CoverTab[80546]++
												return toRPCErr(err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1951
		// _ = "end of CoverTab[80546]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1952
		_go_fuzz_dep_.CoverTab[80547]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1952
		// _ = "end of CoverTab[80547]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1952
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1952
	// _ = "end of CoverTab[80542]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1952
	_go_fuzz_dep_.CoverTab[80543]++
											return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1953
	// _ = "end of CoverTab[80543]"
}

// SetSendCompressor sets a compressor for outbound messages from the server.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// It must not be called after any event that causes headers to be sent
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// (see ServerStream.SetHeader for the complete list). Provided compressor is
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// used when below conditions are met:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//   - compressor is registered via encoding.RegisterCompressor
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//   - compressor name must exist in the client advertised compressor names
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//     sent in grpc-accept-encoding header. Use ClientSupportedCompressors to
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//     get client supported compressor names.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// The context provided must be the context passed to the server's handler.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// It must be noted that compressor name encoding.Identity disables the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// outbound compression.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// By default, server messages will be sent using the same compressor with
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// which request messages were sent.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// It is not safe to call SetSendCompressor concurrently with SendHeader and
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// SendMsg.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// Notice: This function is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1956
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1979
func SetSendCompressor(ctx context.Context, name string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1979
	_go_fuzz_dep_.CoverTab[80548]++
											stream, ok := ServerTransportStreamFromContext(ctx).(*transport.Stream)
											if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1981
		_go_fuzz_dep_.CoverTab[80551]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1981
		return stream == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1981
		// _ = "end of CoverTab[80551]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1981
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1981
		_go_fuzz_dep_.CoverTab[80552]++
												return fmt.Errorf("failed to fetch the stream from the given context")
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1982
		// _ = "end of CoverTab[80552]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1983
		_go_fuzz_dep_.CoverTab[80553]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1983
		// _ = "end of CoverTab[80553]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1983
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1983
	// _ = "end of CoverTab[80548]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1983
	_go_fuzz_dep_.CoverTab[80549]++

											if err := validateSendCompressor(name, stream.ClientAdvertisedCompressors()); err != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1985
		_go_fuzz_dep_.CoverTab[80554]++
												return fmt.Errorf("unable to set send compressor: %w", err)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1986
		// _ = "end of CoverTab[80554]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1987
		_go_fuzz_dep_.CoverTab[80555]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1987
		// _ = "end of CoverTab[80555]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1987
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1987
	// _ = "end of CoverTab[80549]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1987
	_go_fuzz_dep_.CoverTab[80550]++

											return stream.SetSendCompress(name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1989
	// _ = "end of CoverTab[80550]"
}

// ClientSupportedCompressors returns compressor names advertised by the client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
// via grpc-accept-encoding header.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
// The context provided must be the context passed to the server's handler.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
// # Experimental
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
// Notice: This function is EXPERIMENTAL and may be changed or removed in a
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:1992
// later release.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2001
func ClientSupportedCompressors(ctx context.Context) ([]string, error) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2001
	_go_fuzz_dep_.CoverTab[80556]++
											stream, ok := ServerTransportStreamFromContext(ctx).(*transport.Stream)
											if !ok || func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2003
		_go_fuzz_dep_.CoverTab[80558]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2003
		return stream == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2003
		// _ = "end of CoverTab[80558]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2003
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2003
		_go_fuzz_dep_.CoverTab[80559]++
												return nil, fmt.Errorf("failed to fetch the stream from the given context %v", ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2004
		// _ = "end of CoverTab[80559]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2005
		_go_fuzz_dep_.CoverTab[80560]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2005
		// _ = "end of CoverTab[80560]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2005
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2005
	// _ = "end of CoverTab[80556]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2005
	_go_fuzz_dep_.CoverTab[80557]++

											return strings.Split(stream.ClientAdvertisedCompressors(), ","), nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2007
	// _ = "end of CoverTab[80557]"
}

// SetTrailer sets the trailer metadata that will be sent when an RPC returns.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2010
// When called more than once, all the provided metadata will be merged.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2010
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2010
// The error returned is compatible with the status package.  However, the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2010
// status code will often not match the RPC status as seen by the client
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2010
// application, and therefore, should not be relied upon for this purpose.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2016
func SetTrailer(ctx context.Context, md metadata.MD) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2016
	_go_fuzz_dep_.CoverTab[80561]++
											if md.Len() == 0 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2017
		_go_fuzz_dep_.CoverTab[80564]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2018
		// _ = "end of CoverTab[80564]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2019
		_go_fuzz_dep_.CoverTab[80565]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2019
		// _ = "end of CoverTab[80565]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2019
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2019
	// _ = "end of CoverTab[80561]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2019
	_go_fuzz_dep_.CoverTab[80562]++
											stream := ServerTransportStreamFromContext(ctx)
											if stream == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2021
		_go_fuzz_dep_.CoverTab[80566]++
												return status.Errorf(codes.Internal, "grpc: failed to fetch the stream from the context %v", ctx)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2022
		// _ = "end of CoverTab[80566]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2023
		_go_fuzz_dep_.CoverTab[80567]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2023
		// _ = "end of CoverTab[80567]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2023
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2023
	// _ = "end of CoverTab[80562]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2023
	_go_fuzz_dep_.CoverTab[80563]++
											return stream.SetTrailer(md)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2024
	// _ = "end of CoverTab[80563]"
}

// Method returns the method string for the server context.  The returned
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2027
// string is in the format of "/service/method".
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2029
func Method(ctx context.Context) (string, bool) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2029
	_go_fuzz_dep_.CoverTab[80568]++
											s := ServerTransportStreamFromContext(ctx)
											if s == nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2031
		_go_fuzz_dep_.CoverTab[80570]++
												return "", false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2032
		// _ = "end of CoverTab[80570]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2033
		_go_fuzz_dep_.CoverTab[80571]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2033
		// _ = "end of CoverTab[80571]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2033
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2033
	// _ = "end of CoverTab[80568]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2033
	_go_fuzz_dep_.CoverTab[80569]++
											return s.Method(), true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2034
	// _ = "end of CoverTab[80569]"
}

type channelzServer struct {
	s *Server
}

func (c *channelzServer) ChannelzMetric() *channelz.ServerInternalMetric {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2041
	_go_fuzz_dep_.CoverTab[80572]++
											return c.s.channelzMetric()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2042
	// _ = "end of CoverTab[80572]"
}

// validateSendCompressor returns an error when given compressor name cannot be
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2045
// handled by the server or the client based on the advertised compressors.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2047
func validateSendCompressor(name, clientCompressors string) error {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2047
	_go_fuzz_dep_.CoverTab[80573]++
											if name == encoding.Identity {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2048
		_go_fuzz_dep_.CoverTab[80577]++
												return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2049
		// _ = "end of CoverTab[80577]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2050
		_go_fuzz_dep_.CoverTab[80578]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2050
		// _ = "end of CoverTab[80578]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2050
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2050
	// _ = "end of CoverTab[80573]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2050
	_go_fuzz_dep_.CoverTab[80574]++

											if !grpcutil.IsCompressorNameRegistered(name) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2052
		_go_fuzz_dep_.CoverTab[80579]++
												return fmt.Errorf("compressor not registered %q", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2053
		// _ = "end of CoverTab[80579]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2054
		_go_fuzz_dep_.CoverTab[80580]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2054
		// _ = "end of CoverTab[80580]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2054
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2054
	// _ = "end of CoverTab[80574]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2054
	_go_fuzz_dep_.CoverTab[80575]++

											for _, c := range strings.Split(clientCompressors, ",") {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2056
		_go_fuzz_dep_.CoverTab[80581]++
												if c == name {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2057
			_go_fuzz_dep_.CoverTab[80582]++
													return nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2058
			// _ = "end of CoverTab[80582]"
		} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2059
			_go_fuzz_dep_.CoverTab[80583]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2059
			// _ = "end of CoverTab[80583]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2059
		}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2059
		// _ = "end of CoverTab[80581]"
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2060
	// _ = "end of CoverTab[80575]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2060
	_go_fuzz_dep_.CoverTab[80576]++
											return fmt.Errorf("client does not support compressor %q", name)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2061
	// _ = "end of CoverTab[80576]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2062
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/server.go:2062
var _ = _go_fuzz_dep_.CoverTab
